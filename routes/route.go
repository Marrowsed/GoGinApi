package routes

import (
	"gin_api/controllers"

	"github.com/gin-gonic/gin"
)

func HandleRequests() {
	r := gin.Default()
	r.GET("/students", controllers.ReturnStudents)
	r.GET("/students/:id", controllers.ReturnStudentById)
	r.GET("/:name", controllers.Greetings)
	r.POST("/students", controllers.CreateStudent)
	r.DELETE("/students/:id", controllers.DeleteById)
	r.PATCH("/students/:id", controllers.UpdateById)
	r.GET("/students/document/:document", controllers.ReturnStudentByDocument)
	r.Run()
}
