package main

import (
	"github.com/gin-gonic/gin"
	cors "github.com/rs/cors/wrapper/gin"
	"schedule/controllers"
	"schedule/middlewares"
	"schedule/models"
)

func main() {

	models.ConnectDataBase()

	r := gin.Default()
	r.Use(cors.AllowAll())
	r.Static("/static", "./static")

	public := r.Group("/")
	public.POST("/register", controllers.Register)
	public.POST("/login", controllers.Login)

	protected := r.Group("/admin")
	protected.Use(middlewares.JwtAuthMiddleware())
	protected.GET("/user", controllers.CurrentUser)

	projects := r.Group("/projects")
	projects.POST("/create", controllers.CreateProject)
	projects.GET("/getUserProjects/:user_id", controllers.GetUserProjects)
	projects.GET("/getProjectById/:project_id", controllers.GetProjectById)
	projects.PATCH("/updateInformation", controllers.UpdateInformation)

	columns := r.Group("/columns")
	columns.POST("/create", controllers.CreateColumn)
	columns.PATCH("/update", controllers.UpdateColumn)

	sections := r.Group("/sections")
	sections.PATCH("/update", controllers.UpdateSection)

	textPoints := r.Group("/textPoints")
	textPoints.PATCH("/update", controllers.UpdateTextPoint)

	r.Run(":8080")

}
