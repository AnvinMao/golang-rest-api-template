package handler

import (
	"golang-rest-api-template/internal/model"
	"golang-rest-api-template/internal/service"
	"golang-rest-api-template/pkg/app"
	"golang-rest-api-template/pkg/auth"
	"golang-rest-api-template/pkg/utils"

	"github.com/gin-gonic/gin"
)

type UserHandler struct {
	UserService service.UserService
	Jwt         *auth.JWT
}

func (h *UserHandler) FetchUsers(c *gin.Context) {
	r := app.Response{C: c}
	var query model.UserQuery
	if errors := utils.BindAndValidate(c, &query); errors != nil {
		r.BadRequest(errors)
		return
	}

	users, err := h.UserService.Fetch(&query)
	if err != nil {
		r.Error("failed to get data")
		return
	}

	r.Success(users)
}

func (h *UserHandler) GetUser(c *gin.Context) {
	r := app.Response{C: c}
	uid := utils.StrTo(c.Param("id")).MustInt64()
	user, err := h.UserService.GetUser(uid)
	if err != nil {
		r.NotFound("user not founded")
		return
	}

	r.Success(user)
}

func (h *UserHandler) CreateUser(c *gin.Context) {
	r := app.Response{C: c}
	var request model.RegisterRequest
	if errors := utils.BindAndValidate(c, &request); errors != nil {
		r.BadRequest(errors)
		return
	}

	user, err := h.UserService.Create(&request)
	if err != nil {
		r.Error("failed to create user")
		return
	}

	r.Success(user)
}
