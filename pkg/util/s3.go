package util

import (
	"context"
	"fmt"
	"log"
	"mime/multipart"
	"os"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/feature/s3/manager"
	"github.com/aws/aws-sdk-go-v2/service/s3"
	"github.com/joho/godotenv"
)

// UploadImage uploads a file to an AWS S3 bucket.
func UploadImage(file *multipart.FileHeader) (string, error) {

	err := godotenv.Load()
	if err != nil {
		log.Fatalf("Error loading .env file")
	}

	awsRegion := os.Getenv("AWS_REGION")
	if awsRegion == "" {
		log.Fatalf("AWS_REGION is not set in the environment")
	}

	// Load the AWS config with the specified region
	cfg, err := config.LoadDefaultConfig(context.TODO(), config.WithRegion(awsRegion))
	if err != nil {
		log.Printf("error loading AWS config: %v", err)
		return "", fmt.Errorf("error loading AWS config: %w", err)
	}

	// Create a new S3 client
	client := s3.NewFromConfig(cfg)

	// Open the file
	f, err := file.Open()
	if err != nil {
		log.Printf("error opening file: %v", err)
		return "", fmt.Errorf("error opening file: %w", err)
	}
	defer f.Close()

	// Create an uploader with the S3 client
	uploader := manager.NewUploader(client)

	bucketName := "gbubemi"
	key := file.Filename
	_, err = uploader.Upload(context.TODO(), &s3.PutObjectInput{
		Bucket: aws.String(bucketName),
		Key:    aws.String(key),
		Body:   f,
		ACL:    "public-read",
	})
	if err != nil {
		log.Printf("error uploading file to S3: %v", err)
		return "", fmt.Errorf("error uploading file to S3: %w", err)
	}

	// Return the URL of the uploaded file
	region := cfg.Region // Get the region from the config
	url := fmt.Sprintf("https://%s.s3.%s.amazonaws.com/%s", bucketName, region, key)
	return url, nil

	// // Return the URL of the uploaded file
	// url := fmt.Sprintf("https://%s.s3.%s.amazonaws.com/%s", "gbubemi", awsRegion, file.Filename)
	// return url, nil
}

// func (controller *UserController) UploadImageHandler(c *gin.Context) {
// 	// Get the user ID from the context
// 	userID, exists := c.Get("userID")
// 	if !exists {
// 		c.JSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized"})
// 		return
// 	}

// 	userIDStr, ok := userID.(string)
// 	if !ok {
// 		c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid user ID"})
// 		return
// 	}

// 	// Get the image from form data
// 	file, err := c.FormFile("image")
// 	if err != nil {
// 		c.JSON(http.StatusBadRequest, gin.H{"error": "Image file is required"})
// 		return
// 	}

// 	// Upload the image directly using utils.UploadImage
// 	imageURL, err := utils.UploadImage(file)
// 	if err != nil {
// 		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to upload image"})
// 		return
// 	}

// 	// Update the user with the image URL
// 	if err := controller.userService.UploadImage(userIDStr, imageURL); err != nil {
// 		if err.Error() == "user not found" {
// 			c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
// 			return
// 		}
// 		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
// 		return
// 	}

// 	c.JSON(http.StatusOK, gin.H{"message": "Image uploaded successfully"})
// }
