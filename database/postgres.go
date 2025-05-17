package database

import (
	"fmt"
	"golang-rest-api-template/env"
	"log"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

type Database interface {
	Select(query any, args ...any) *gorm.DB
	Offset(offset int) *gorm.DB
	Limit(limit int) *gorm.DB
	Find(any, ...any) *gorm.DB
	Create(value any) *gorm.DB
	Where(query any, args ...any) Database
	Delete(any, ...any) *gorm.DB
	Model(model any) *gorm.DB
	First(dest any, conds ...any) Database
	Updates(any) *gorm.DB
	Order(value any) *gorm.DB
	Error() error
}

type GormDatabase struct {
	*gorm.DB
}

func (db *GormDatabase) Where(query any, args ...any) Database {
	return &GormDatabase{db.DB.Where(query, args...)}
}

func (db *GormDatabase) First(dest any, conds ...any) Database {
	return &GormDatabase{db.DB.First(dest, conds...)}
}

func (db *GormDatabase) Error() error {
	return db.DB.Error
}

func NewDatabase(env *env.Env) *gorm.DB {
	dsn := fmt.Sprintf(
		"host=%s user=%s password=%s dbname=%s port=%s sslmode=disable",
		env.DBHost, env.DBUser, env.DBPassword, env.DBName, env.DBPort,
	)

	database, err := gorm.Open(postgres.Open(dsn), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Info),
	})

	if err != nil {
		log.Fatalf("Failed to connect to database: %v", err)
	}

	db, err := database.DB()
	if err != nil {
		log.Fatalf("Failed to get raw DB: %v", err)
	}

	db.SetMaxOpenConns(env.DBMaxConn)
	db.SetMaxIdleConns(env.DBMaxIdle)
	db.SetConnMaxIdleTime(env.DBIdleTimeout)

	fmt.Println("Successfully connected to PostgreSQL!")

	return database
}
