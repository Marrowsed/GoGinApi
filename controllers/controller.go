package controllers

import (
	"gin_api/database"
	"gin_api/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

/*func ReturnMockData(c *gin.Context) {
	c.JSON(200, gin.H{
		"id":   1,
		"name": "John Doe",
	})
	c.JSON(200, models.Students)
}
*/

func Greetings(c *gin.Context) {
	name := c.Params.ByName("name")
	c.JSON(200, gin.H{
		"Greetings": "Hello, " + name + " Welcome !",
	})
}

func ReturnStudents(c *gin.Context) {
	var allStudents []models.Student
	database.DB.Find(&allStudents)
	if len(allStudents) == 0 {
		c.JSON(http.StatusNotFound, gin.H{
			"Error": "Students not found !",
		})
		return
	}
	c.JSON(http.StatusOK, allStudents)
}

func ReturnStudentById(c *gin.Context) {
	var student models.Student

	id := c.Params.ByName("id")

	database.DB.First(&student, id)

	if student.ID == 0 {
		c.JSON(http.StatusNotFound, gin.H{
			"Error": "Student not found !",
		})
		return
	}
	c.JSON(http.StatusOK, student)
}

func ReturnStudentByDocument(c *gin.Context) {
	var student models.Student
	document := c.Param("cpf")
	database.DB.Where(&models.Student{Document: document}).First(&student)

	if student.ID == 0 {
		c.JSON(http.StatusNotFound, gin.H{
			"Error": "Student not found !",
		})
		return
	}

	c.JSON(http.StatusOK, student)
}

func CreateStudent(c *gin.Context) {
	var newStudent models.Student
	if err := c.ShouldBindJSON(&newStudent); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	database.DB.Create(&newStudent)
	c.JSON(http.StatusCreated, newStudent)
}

func DeleteById(c *gin.Context) {
	var student models.Student

	id := c.Params.ByName("id")

	database.DB.Delete(&student, id)

	c.JSON(204, gin.H{
		"Message": "Student deleted !",
	})
}

func UpdateById(c *gin.Context) {
	var student models.Student
	id := c.Params.ByName("id")
	database.DB.First(&student, id)

	if err := c.ShouldBindJSON(&student); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error()})
		return
	}

	database.DB.Model(&student).UpdateColumns(student)
	c.JSON(http.StatusOK, &student)

}
