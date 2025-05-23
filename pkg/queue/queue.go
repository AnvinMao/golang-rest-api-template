package queue

import (
	"encoding/json"
	"fmt"
	"golang-rest-api-template/pkg/database"
	"strings"
	"time"

	"github.com/redis/go-redis/v9"
)

type Task struct {
	Type    string                 `json:"type"`
	Payload map[string]interface{} `json:"payload"`
}

type RedisQueue struct {
	redis      database.Redis
	streamName string
	groupName  string
}

func NewRedisQueue(redis database.Redis, streamName, groupName string) (*RedisQueue, error) {
	err := redis.XGroupCreateMkStream(streamName, groupName, "0").Err()
	if err != nil && !strings.Contains(err.Error(), "BUSYGROUP") {
		return nil, err
	}

	return &RedisQueue{
		redis:      redis,
		streamName: streamName,
		groupName:  groupName,
	}, nil
}

func (q *RedisQueue) Produce(task Task) error {
	data, err := json.Marshal(task)
	if err != nil {
		return err
	}

	return q.redis.XAdd(&redis.XAddArgs{
		Stream: q.streamName,
		Values: map[string]interface{}{"data": data},
	}).Err()
}

func (q *RedisQueue) Consume(consumer string, handler func(Task) error) error {
	for {
		streams, err := q.redis.XReadGroup(&redis.XReadGroupArgs{
			Group:    q.groupName,
			Consumer: consumer,
			Streams:  []string{q.streamName, ">"},
			Count:    1,
			Block:    5 * time.Second,
		}).Result()

		if err != nil && err != redis.Nil {
			fmt.Printf("Consume error: %v\n", err)
			time.Sleep(time.Second)
			continue
		}

		for _, stream := range streams {
			for _, message := range stream.Messages {
				var task Task
				data, ok := message.Values["data"].(string)
				if !ok {
					continue
				}

				if err := json.Unmarshal([]byte(data), &task); err != nil {
					continue
				}

				if err := handler(task); err != nil {
					continue
				}

				q.redis.XAck(q.streamName, q.groupName, message.ID)
			}
		}
	}
}
