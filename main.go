package main

import (
	"gin_api/database"
	"gin_api/routes"
)

func main() {
	/* mock Data
	models.Students = []models.Student{
		{Name: "John Doe", Document: "000000000", Age: 1},
		{Name: "Lorem Ipsun", Document: "111111111", Age: 2},
	} */
	database.DbConnect()
	routes.HandleRequests()
}
