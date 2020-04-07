package router

import (
	"experiment/handler/experiment"
	"experiment/handler/problem"
	"net/http"

	"experiment/handler/sd"
	"experiment/handler/user"
	"experiment/router/middleware"

	"github.com/gin-gonic/gin"
)

// Load loads the middlewares, routes, handlers.
func Load(g *gin.Engine, mw ...gin.HandlerFunc) *gin.Engine {
	// Middlewares.
	g.Use(gin.Recovery())
	g.Use(middleware.NoCache)
	g.Use(middleware.Options)
	g.Use(middleware.Secure)
	g.Use(mw...)
	// 404 Handler.
	g.NoRoute(func(c *gin.Context) {
		c.String(http.StatusNotFound, "The incorrect API route.")
	})

	g.POST("/api/login", user.Login)

	u := g.Group("/api/user")
	{
		u.POST("", user.Create)
		u.GET("/:id", middleware.AuthMiddleware(), user.Get)
	}

	p := g.Group("/api/problem")
	{
		p.POST("", middleware.TeacherAuthMiddleware(), problem.Create)
		p.PUT("/:id", middleware.TeacherAuthMiddleware(), problem.UploadData)
	}

	e := g.Group("/api/experiment")
	{
		e.POST("", middleware.TeacherAuthMiddleware(), experiment.Create)
	}

	// The health check handlers
	svcd := g.Group("/sd")
	{
		svcd.GET("/health", sd.HealthCheck)
		svcd.GET("/disk", sd.DiskCheck)
		svcd.GET("/cpu", sd.CPUCheck)
		svcd.GET("/ram", sd.RAMCheck)
	}

	return g
}
