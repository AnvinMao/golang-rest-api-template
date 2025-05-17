package service

import (
	"context"
	"encoding/json"
	"fmt"
	"golang-rest-api-template/database"
	"golang-rest-api-template/internal/model"
	"golang-rest-api-template/pkg/auth"
	"strings"
	"time"
)

type UserService interface {
	Create(request *model.RegisterRequest) (*model.User, error)
	Fetch(query *model.UserQuery) ([]model.User, error)
	GetUser(id int64) (*model.User, error)
}

type userService struct {
	DB    database.Database
	Redis database.Redis
	Ctx   *context.Context
}

func NewUserService(db database.Database, redis database.Redis, ctx *context.Context) UserService {
	return &userService{
		DB:    db,
		Redis: redis,
		Ctx:   ctx,
	}
}

func (serv *userService) Create(request *model.RegisterRequest) (*model.User, error) {
	user := &model.User{
		Name:     request.Email,
		Password: auth.HashPassword(request.Password),
		Email:    request.Email,
		Username: strings.Split(request.Email, "@")[0],
		Status:   1,
	}

	result := serv.DB.Create(user)
	return user, result.Error
}

func (serv *userService) Fetch(query *model.UserQuery) ([]model.User, error) {
	var users []model.User

	q := serv.DB.Select("id", "name", "username")
	if query.Email != "" {
		q.Where("email = ?", query.Email)
	}

	page := 1
	if query.Page > 0 {
		page = query.Page
	}

	pageSize := 10
	offset := (page - 1) * pageSize

	err := q.Where("status = ?", 1).Order("id DESC").Limit(pageSize).Offset(offset).Find(&users).Error
	if err != nil {
		return nil, err
	}

	return users, nil
}

func (serv *userService) GetUser(id int64) (*model.User, error) {
	var user model.User
	cacheKey := fmt.Sprintf("user.%d", id)
	cacheUser, err := serv.Redis.Get(*serv.Ctx, cacheKey).Result()
	if err == nil {
		err := json.Unmarshal([]byte(cacheUser), &user)
		if err == nil {
			return &user, nil
		}
	}

	if err := serv.DB.First(&user, id).Error(); err != nil {
		return nil, err
	}

	if serialized, err := json.Marshal(user); err == nil {
		serv.Redis.Set(*serv.Ctx, cacheKey, serialized, time.Minute)
	}

	return &user, nil
}
