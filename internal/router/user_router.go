package router

import (
	"golang-rest-api-template/internal/handler"
	"golang-rest-api-template/internal/middleware"
	"golang-rest-api-template/internal/service"
	"golang-rest-api-template/pkg/auth"
	"golang-rest-api-template/pkg/database"

	"github.com/gin-gonic/gin"
)

func RegisterUserRoutes(r *gin.RouterGroup, db database.Database, redis database.Redis, jwt *auth.JWT) {
	user := &handler.UserHandler{
		UserService: service.NewUserService(db, redis),
		Jwt:         jwt,
	}

	auth := r.Group("user", middleware.AuthUser(jwt))
	{
		auth.GET("list", user.FetchUsers)
		auth.GET(":id", user.GetUser)
		auth.POST("create", user.CreateUser)
	}
}
