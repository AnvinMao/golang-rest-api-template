package app

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type Response struct {
	C *gin.Context
}

type ErrorResponse struct {
	Message any `json:"message"`
}

func (r *Response) Success(data any) {
	r.C.JSON(http.StatusOK, data)
}

func (r *Response) Error(msg string) {
	r.C.AbortWithStatusJSON(http.StatusInternalServerError, ErrorResponse{
		Message: msg,
	})
}

func (r *Response) BadRequest(msg any) {
	r.C.AbortWithStatusJSON(http.StatusBadRequest, ErrorResponse{
		Message: msg,
	})
}

func (r *Response) NotFound(msg string) {
	r.C.AbortWithStatusJSON(http.StatusNotFound, ErrorResponse{
		Message: msg,
	})
}

func (r *Response) Unauthorized(msg string) {
	r.C.AbortWithStatusJSON(http.StatusUnauthorized, ErrorResponse{
		Message: msg,
	})
}
