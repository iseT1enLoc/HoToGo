package services

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"
	"path/filepath"

	"github.com/cloudinary/cloudinary-go"
	"github.com/cloudinary/cloudinary-go/api/uploader"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

/*
	func UploadCloud(w http.ResponseWriter, r *http.Request) {
		r.ParseMultipartForm(10 << 20) // Max upload size

		file, handler, err := r.FormFile("file")
		if err != nil {
			fmt.Println("Error Retrieving the File")
			fmt.Println(err)
			return
		}
		defer file.Close()

		fmt.Printf("Uploaded File: %+v\n", handler.Filename)
		fmt.Printf("File Size: %+v\n", handler.Size)
		fmt.Printf("MIME Header: %+v\n", handler.Header)

		// Initialize Cloudinary configuration
		// load .env file
		errEnv := godotenv.Load()

		if errEnv != nil {
			log.Fatalf("Error loading .env file")
		}

		// Create Cloudinary uploader
		cld, err := cloudinary.NewFromParams(os.Getenv("CLOUD_NAME"), os.Getenv("CLOUD_API_KEY"), os.Getenv("CLOUD_API_SECRET"))
		if err != nil {
			fmt.Println("Error Initializing Cloudinary Configuration")
			fmt.Println(err)
			return
		}
		uploadApi := cld.Upload

		// get file name
		var publicID = handler.Filename[:len(handler.Filename)-4]

		// Upload file to Cloudinary
		result, err := uploadApi.Upload(context.Background(), file, uploader.UploadParams{PublicID: publicID})
		if err != nil {
			fmt.Println("Error Uploading File to Cloudinary")
			fmt.Println(err)
			return
		}

		fmt.Println("File Successfully Uploaded to Cloudinary")
		fmt.Println("Public ID:", result.PublicID)
		fmt.Println("URL:", result.SecureURL)

		// Return response to client
		w.WriteHeader(http.StatusOK)
		// format a response object
		res := utils.UploadResponse{
			Message: "File Successfully Uploaded to Cloudinary",
			URL:     result.SecureURL,
		}

		// send the response
		json.NewEncoder(w).Encode(res)
	}
*/
func UploadCloud(c *gin.Context) {
	fmt.Println("Enter upload")
	// Parse the multipart form with a max upload size
	err := c.Request.ParseMultipartForm(10 << 20)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Failed to parse multipart form"})
		fmt.Println("line 82")
		return
	}

	// Retrieve the file from the form
	file, handler, err := c.Request.FormFile("file")
	if err != nil {
		fmt.Println("line 88")
		c.JSON(http.StatusBadRequest, gin.H{"error": "Error retrieving the file"})
		return
	}
	defer file.Close()

	fmt.Printf("Uploaded File: %+v\n", handler.Filename)
	fmt.Printf("File Size: %+v\n", handler.Size)
	fmt.Printf("MIME Header: %+v\n", handler.Header)

	// Initialize Cloudinary configuration
	errEnv := godotenv.Load()
	if errEnv != nil {
		log.Fatalf("Error loading .env file")
	}

	// Create Cloudinary uploader
	cld, err := cloudinary.NewFromParams(os.Getenv("CLOUD_NAME"), os.Getenv("CLOUD_API_KEY"), os.Getenv("CLOUD_API_SECRET"))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error initializing Cloudinary configuration"})
		return
	}
	uploadApi := cld.Upload

	// Get file name without extension for the public ID
	publicID := handler.Filename[:len(handler.Filename)-len(filepath.Ext(handler.Filename))]

	// Upload file to Cloudinary
	result, err := uploadApi.Upload(context.Background(), file, uploader.UploadParams{PublicID: publicID})
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error uploading file to Cloudinary"})
		return
	}

	fmt.Println("File Successfully Uploaded to Cloudinary")
	fmt.Println("Public ID:", result.PublicID)
	fmt.Println("URL:", result.SecureURL)

	// Return response to client
	c.JSON(http.StatusOK, gin.H{
		"message": "File successfully uploaded to Cloudinary",
		"url":     result.SecureURL,
	})
}
