package env

import (
	"log"
	"os"
	"strconv"
	"time"

	"github.com/joho/godotenv"
)

type Env struct {
	AppEnv           string
	ServerAddress    string
	ReadTimeout      time.Duration
	WriteTimeout     time.Duration
	SecretKey        string
	TokenTTL         time.Duration
	DBHost           string
	DBPort           string
	DBName           string
	DBUser           string
	DBPassword       string
	DBMaxConn        int
	DBMaxIdle        int
	DBIdleTimeout    time.Duration
	RedisAddress     string
	RedisPassword    string
	RedisMaxIdle     int
	RedisMaxActive   int
	RedisIdleTimeout time.Duration
}

func NewEnv() *Env {
	err := godotenv.Load()
	if err != nil {
		log.Println("Warning: Could not load .env file, using system environment variables.")
	}

	return &Env{
		AppEnv:           getEnv("APP_ENV", "debug"),
		ServerAddress:    getEnv("SERVER_ADDRESS", ":8080"),
		ReadTimeout:      getDuration("READ_TIME_OUT", 20*time.Second),
		WriteTimeout:     getDuration("WRITE_TIME_OUT", 20*time.Second),
		SecretKey:        getEnv("SECRET_KEY", "secretKey"),
		TokenTTL:         getDuration("TOKEN_TTL", 24*time.Hour),
		DBHost:           getEnv("DB_HOST", "127.0.0.1"),
		DBPort:           getEnv("DB_PORT", "5432"),
		DBName:           getEnv("DB_NAME", ""),
		DBUser:           getEnv("DB_USER", ""),
		DBPassword:       getEnv("DB_PASSWORD", ""),
		DBMaxConn:        getInt("DB_MAX_CONN", 64),
		DBMaxIdle:        getInt("DB_MAX_IDLE", 10),
		DBIdleTimeout:    getDuration("DB_IDLE_TIME_OUT", 30*time.Minute),
		RedisAddress:     getEnv("REDIS_ADDRESS", "127.0.0.1:6379"),
		RedisPassword:    getEnv("REDIS_PASSWORD", ""),
		RedisMaxIdle:     getInt("REDIS_MAX_IDLE", 10),
		RedisMaxActive:   getInt("REDIS_MAX_ACTIVE", 32),
		RedisIdleTimeout: getDuration("REDIS_IDLE_TIME_OUT", 60*time.Second),
	}
}

func getEnv(key, fallback string) string {
	if value, exists := os.LookupEnv(key); exists {
		return value
	}

	return fallback
}

func getInt(key string, fallback int) int {
	if v, exists := os.LookupEnv(key); exists {
		value, err := strconv.Atoi(v)
		if err == nil {
			return value
		}
	}

	return fallback
}

func getDuration(key string, fallback time.Duration) time.Duration {
	if v, exists := os.LookupEnv(key); exists {
		value, err := time.ParseDuration(v)
		if err == nil {
			return value
		}
	}

	return fallback
}
