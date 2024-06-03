package main

import (
	"example.com/handlers"
	"example.com/services"
	"github.com/gin-gonic/gin"
	_ "github.com/lib/pq"
)

func main() {
	router := gin.Default()
	v1 := router.Group("/v1")
	{
		v1.GET("/customers/", handlers.GetAllCustomer)
		v1.GET("/customers/:id", handlers.GetCustomerById)
		v1.POST("/customers", handlers.InsertNewCustomer)
		v1.PUT("/customers/:id", handlers.UpdateCustomerInfo)
		v1.DELETE("/customers/:id", handlers.DeleteCustomer)

		v1.GET("/field_types", handlers.GetAllFieldType)

		v1.POST("/upload/", services.UploadCloud)
	}

	router.Run("localhost:8080")
}
