package router

import (
	"golang-rest-api-template/env"
	"golang-rest-api-template/pkg/auth"
	"golang-rest-api-template/pkg/database"

	"github.com/gin-gonic/gin"
)

func RegisterRoutes(r *gin.Engine, db database.Database, redis database.Redis, env *env.Env) {
	r.Use(gin.Logger())
	r.Use(gin.Recovery())

	jwt := &auth.JWT{
		SecretKey: []byte(env.SecretKey),
		Duration:  env.TokenTTL,
	}

	apiRouter := r.Group("/api")
	RegisterUserRoutes(apiRouter, db, redis, jwt)
}
