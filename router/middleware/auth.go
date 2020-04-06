package middleware

import (
	"experiment/handler"
	"experiment/pkg/constvar"
	"experiment/pkg/errno"
	"experiment/pkg/token"
	"github.com/gin-gonic/gin"
)

func AuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		// Parse the json web token.
		ctx, err := token.ParseRequest(c)
		if err != nil {
			handler.SendResponse(c, errno.ErrTokenInvalid, nil)
			c.Abort()
			return
		}
		c.Set("number", ctx.Number)
		c.Set("type", ctx.Type)
		c.Set("username", ctx.Username)
		c.Next()
	}
}

func StudentAuthMiddleware() gin.HandlerFunc {
	return func(context *gin.Context) {
		// Parse the json web token.
		ctx, err := token.ParseRequest(context)
		if err != nil {
			handler.SendResponse(context, errno.ErrTokenInvalid, nil)
			context.Abort()
			return
		}
		if ctx.Type != constvar.Student {
			handler.SendResponse(context, errno.ErrAuthority, nil)
			context.Abort()
			return
		}
		context.Next()
	}
}

func TeacherAuthMiddleware() gin.HandlerFunc {
	return func(context *gin.Context) {
		// Parse the json web token.
		ctx, err := token.ParseRequest(context)
		if err != nil {
			handler.SendResponse(context, errno.ErrTokenInvalid, nil)
			context.Abort()
			return
		}
		if ctx.Type != constvar.Teacher {
			handler.SendResponse(context, errno.ErrAuthority, nil)
			context.Abort()
			return
		}
		context.Next()
	}
}
