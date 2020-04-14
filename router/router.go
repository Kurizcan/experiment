package router

import (
	"experiment/handler/class"
	"experiment/handler/experiment"
	"experiment/handler/problem"
	"experiment/handler/student"
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
		e.GET("/list/:id", middleware.AuthMiddleware(), experiment.ProblemList)
		e.GET("/class/detail", middleware.AuthMiddleware(), experiment.ClassDetail)
		e.POST("/class/distributed", middleware.TeacherAuthMiddleware(), experiment.Distributed)
	}

	c := g.Group("api/class")
	{
		// 老师管理的班级列表
		c.GET("/:id", middleware.TeacherAuthMiddleware(), class.GetClassByTid)
		// 查看班级详情，包括各项试验数据
		c.GET("/:id/detail", middleware.TeacherAuthMiddleware(), class.GetClassDetail)
	}

	s := g.Group("api/student")
	s.Use(middleware.StudentAuthMiddleware())
	{
		s.GET("/experiments", student.MyExperiments)
		s.POST("/submit", student.ProblemSubmit)
		s.GET("/submit/:id", student.GetStatus)
		s.GET("/problem/detail", student.GetProblemDetail)
		s.POST("experiment/submit/:id", student.ExperimentSubmit)
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
