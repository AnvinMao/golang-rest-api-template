package router

import (
	"context"
	"golang-rest-api-template/database"
	"golang-rest-api-template/internal/handler"
	"golang-rest-api-template/internal/middleware"
	"golang-rest-api-template/internal/service"
	"golang-rest-api-template/pkg/auth"

	"github.com/gin-gonic/gin"
)

func RegisterUserRoutes(r *gin.RouterGroup, db database.Database, redis database.Redis, ctx *context.Context, jwt *auth.JWT) {
	c := &handler.UserHandler{
		UserService: service.NewUserService(db, redis, ctx),
		Jwt:         jwt,
	}

	auth := r.Group("/user", middleware.AuthUser(jwt))
	{
		auth.GET("list", c.FetchUsers)
		auth.GET(":id", c.GetUser)
		auth.POST("create", c.CreateUser)
	}
}
