package middleware

import (
	"golang-rest-api-template/pkg/app"
	"golang-rest-api-template/pkg/auth"
	"strings"

	"github.com/gin-gonic/gin"
)

func AuthUser(jwt *auth.JWT) gin.HandlerFunc {

	return func(c *gin.Context) {
		r := app.Response{C: c}
		header := c.GetHeader("Authorization")
		if header == "" {
			r.Unauthorized("token lost")
			return
		}

		const BearerSchema = "Bearer "
		if !strings.HasPrefix(header, BearerSchema) {
			r.Unauthorized("invalid token")
			return
		}

		token := header[len(BearerSchema):]

		claims, err := jwt.ParseToken(token)
		if err != nil {
			r.Unauthorized("invalid token")
			return
		}

		c.Set("userId", claims.UserID)
		c.Next()
	}
}
