package handlers

import (
	"fmt"
	"log"
	"net/http"
	"repository"

	utils "example.com/response"
	"github.com/gin-gonic/gin"
)

func GetAllFieldType(c *gin.Context) {
	// get all the customer in the db
	fmt.Println("Entered get all customer function")
	field_types, err := repository.GetAllFieldType()

	if err != nil {
		log.Fatalf("Unable to get all topic. %v", err)
	}

	res := utils.Response{Message: "ok", Data: field_types}
	c.JSON(http.StatusOK, res)
}
