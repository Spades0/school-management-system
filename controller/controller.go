package controller

import (
	"github.com/gin-gonic/gin"
	"github.com/ucheonukaji/sms/models"
	"github.com/ucheonukaji/sms/setup"
	"net/http"
)

var DB = setup.GetDB()

func CreateCourse(c *gin.Context) {
	var course models.Course
	if err := c.ShouldBind(&course); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	}
	result := DB.Create(&course)
	if result.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": result.Error.Error()})
		return
	}  
	c.JSON(http.StatusOK, gin.H{"message": course})
}

func RetrieveCourses(c *gin.Context) {
	var courses []models.Course
	if err := DB.Find(&courses).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, courses)
}

func RetrieveCourseById(c *gin.Context) {
	var course models.Course
	id := c.Param("id")
	if err := DB.First(&course, "course_id = ?", id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, course)
}

func CreateStudents(c *gin.Context) {
	var student models.Student
	if err := c.ShouldBind(&student); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	}
	result := DB.Create(&student)
	if result.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": result.Error.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": student})
}

func RetrieveStudents(c *gin.Context) {
	var students []models.Student
	if err := DB.Find(&students).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, students)
}

func RetrieveStudentById(c *gin.Context) {
	var student models.Student
	id := c.Param("id")
	if err := DB.First(&student, "course_id = ?", id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, student)
}
