package router

import (
	"experiment/handler/class"
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
		p.GET("/:id", problem.Detail)
	}

	e := g.Group("/api/experiment")
	{
		e.POST("", middleware.TeacherAuthMiddleware(), experiment.Create)
		e.GET("/:id", middleware.AuthMiddleware(), experiment.ProblemList)
	}

	c := g.Group("api/class")
	{
		// 老师管理的班级列表
		c.GET("/:id", middleware.TeacherAuthMiddleware(), class.GetClassByTid)
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
