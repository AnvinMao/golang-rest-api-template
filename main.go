package main

import (
	"context"
	"golang-rest-api-template/database"
	"golang-rest-api-template/env"
	"golang-rest-api-template/internal/router"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	env := env.NewEnv()
	db := &database.GormDatabase{DB: database.NewDatabase(env)}
	redis := database.NewRedisClient(env)

	ctx := context.Background()

	gin.SetMode(env.AppEnv)
	r := gin.New()
	router.RegisterRoutes(r, db, redis, &ctx, env)

	server := &http.Server{
		Addr:           env.ServerAddress,
		Handler:        r,
		ReadTimeout:    env.ReadTimeout,
		WriteTimeout:   env.WriteTimeout,
		MaxHeaderBytes: 1 << 20,
	}

	log.Printf("[info] http server listening %s", env.ServerAddress)

	server.ListenAndServe()
}
