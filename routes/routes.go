package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/ucheonukaji/sms/controller"
)

func AllRoutes(router *gin.Engine) {
	rc := router.Group("/api/courses")
	{
		rc.POST("/", controller.CreateCourse)
		rc.GET("/", controller.RetrieveCourses)
		rc.GET("/:id", controller.RetrieveCourseById)
	}

	rs := router.Group("/api/students")

	{
		rs.POST("/", controller.CreateStudents)
		rs.GET("/", controller.RetrieveStudents)
		rs.GET("/:id", controller.RetrieveStudentById)
	}

}
