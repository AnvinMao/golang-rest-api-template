package router

import (
	"context"
	"golang-rest-api-template/database"
	"golang-rest-api-template/env"
	"golang-rest-api-template/pkg/auth"

	"github.com/gin-gonic/gin"
)

func RegisterRoutes(r *gin.Engine, db database.Database, redis database.Redis, ctx *context.Context, env *env.Env) {
	r.Use(gin.Logger())
	r.Use(gin.Recovery())

	jwt := &auth.JWT{
		SecretKey: []byte(env.SecretKey),
		Duration:  env.TokenTTL,
	}

	apiRouter := r.Group("/api")
	RegisterUserRoutes(apiRouter, db, redis, ctx, jwt)
}
