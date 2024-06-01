package handlers

import (
	"fmt"
	"log"
	"net/http"
	"repository"

	utils "example.com/response"
	"github.com/gin-gonic/gin"
	models "model.go"
)

func GetAllCustomer(c *gin.Context) {
	// get all the customer in the db
	fmt.Println("Entered get all customer function")
	customers, err := repository.GetAllCustomer()

	if err != nil {
		log.Fatalf("Unable to get all topic. %v", err)
	}

	res := utils.Response{Message: "ok", Data: customers}
	c.JSON(http.StatusOK, res)
}

func GetCustomerById(c *gin.Context) {
	fmt.Println("Entered get customer by id")
	id := c.Param("id")
	customer, err := repository.GetCustomerById(id)

	if err != nil {
		log.Fatalf("Unable to get customer. %v", err)
	}
	res := utils.Response{Message: "Successfully get Customer by id", Data: customer}
	c.JSON(http.StatusOK, res)
}
func InsertNewCustomer(c *gin.Context) {
	fmt.Println("Entered insert new customer")
	//create an empty customer
	// Create an empty topic of type models.topic
	var customer models.Customer

	// Bind the JSON request body to the topic object
	if err := c.ShouldBindJSON(&customer); err != nil {
		fmt.Println(customer)
		c.JSON(http.StatusBadRequest, utils.Response{Message: "Error happend", Data: err})
		//c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": "Unable to decode request body"})
		return
	}

	// Call insert topic function and pass the topic
	insertedCustomer, err := repository.InsertNewCustomer(customer)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	// Format a response object with data
	res := utils.Response{
		Message: "Topic created successfully",
		Data:    insertedCustomer}

	// Send the response with status created
	c.JSON(http.StatusCreated, res)
}
func UpdateCustomerInfo(c *gin.Context) {
	fmt.Println("Enter updated handlers")
	//first we should get the customer id that's going to be updated
	cus_id := c.Param("id")

	// Create an empty customer object
	var updatedCustomer models.Customer

	// Bind the JSON request body to the updatedProduct object
	if err := c.BindJSON(&updatedCustomer); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": "Invalid request body"})
		return
	}
	// Update the customer in the database
	updatedRows := repository.UpdateCustomer(cus_id, updatedCustomer)
	// format the message string
	msg := fmt.Sprintf("customer updated successfully. Total rows/record affected %v", updatedRows)

	updatedCustomer.CustomerId = cus_id

	// format the response message
	res := utils.Response{
		Message: msg,
		Data:    updatedCustomer,
	}

	c.JSON(http.StatusAccepted, res)
}
func DeleteCustomer(c *gin.Context) {
	fmt.Println("Enter delete handlers")
	//first we should get the customer id that's going to be updated
	cus_id := c.Param("id")

	// Update the customer in the database
	deletedRows := repository.DeleteCustomer(cus_id)
	// format the message string
	msg := fmt.Sprintf("customer deleted successfully. Total rows/record affected %v", deletedRows)

	// format the response message
	res := utils.Response{
		Message: msg,
		Data:    deletedRows,
	}

	c.JSON(http.StatusAccepted, res)
}
