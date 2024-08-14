package handlers

// import (
// 	"context"
// 	"fmt"
// 	"go_tg_api_gateway/genproto/coins_service"
// 	"os"
// 	"path/filepath"
// 	"strings"

// 	"github.com/gin-gonic/gin"
// 	"github.com/google/uuid"
// 	"github.com/spf13/cast"
// )

// // UploadUserFile godoc
// // @ID UploadUserFile
// // @Router /user/image [POST]
// // @Summary UploadUserFile
// // @Description  UploadUserFile
// // @Tags User
// // @Accept multipart/form-data
// // @Produce json
// // @Param file formData file true "Upload file"
// // @Success 200 {object} http.Response{data=string} "Success message"
// // @Response 400 {object} http.Response{data=string} "Invalid Argument"
// // @Failure 500 {object} http.Response{data=string} "Server Error"
// func (h *Handler) HandleUserFileupload(c *gin.Context) {

// 	// Get the uploaded file from the form:
// 	file, err := c.FormFile("file")
// 	if err != nil {
// 		c.JSON(400, gin.H{"error": "Unable to get file"})
// 		return
// 	}
// 	// Check file format
// 	ext := filepath.Ext(file.Filename)
// 	validFormats := map[string]bool{
// 		".jpg":  true,
// 		".jpeg": true,
// 		".png":  true,
// 		".gif":  true,
// 	}

// 	if !validFormats[ext] {
// 		c.JSON(400, gin.H{"error": "Invalid file format"})
// 		return
// 	}
// 	// Create the uploads directory if it doesn't exist:
// 	err = os.MkdirAll("images/user", os.ModePerm)
// 	if err != nil {
// 		c.JSON(500, gin.H{"error": "Unable to create uploads directory"})
// 		return
// 	}

// 	// Generate a unique filename using UUID:
// 	uniqueId := uuid.New()
// 	filename := strings.Replace(uniqueId.String(), "-", "", -1) + filepath.Ext(file.Filename)
// 	filePath := filepath.Join("images/user", filename)
// 	// Save the file to the disk:
// 	fmt.Println(filePath)
// 	if err := c.SaveUploadedFile(file, filePath); err != nil {
// 		c.JSON(500, gin.H{"error": "Unable to save file"})
// 		return
// 	}
// 	imageLink := filepath.Join("images/user", strings.Replace(uniqueId.String(), "-", "", -1))
// 	_, err = h.services.FileImage().ImageUpload(
// 		context.Background(),
// 		&coins_service.ImageData{
// 			Id:        cast.ToString(uniqueId),
// 			ImageLink: imageLink,
// 		},
// 	)
// 	if err != nil {
// 		c.JSON(500, gin.H{"error": "Unable to save file info to database"})
// 		return
// 	}

// 	c.JSON(200, gin.H{"message": "File uploaded successfully"})
// }

// // DeleteUser godoc
// // @ID delete_image
// // @Router /user/image/{id} [DELETE]
// // @Summary Delete User Image
// // @Description Delete User Image
// // @Tags User
// // @Accept json
// // @Produce json
// // @Param id path string true "id"
// // @Success 200 {object} http.Response{data=object{}} "Success"
// // @Response 400 {object} http.Response{data=string} "Bad Request"
// // @Failure 500 {object} http.Response{data=string} "Server Error"
// func (h *Handler) HandleUserDeleteImage(c *gin.Context) {

// 	// Get the image name from the URL path
// 	imageName := c.Param("id")
// 	name := strings.Replace(imageName, "-", "", -1)
// 	// Define possible file extensions
// 	extensions := []string{".png", ".gif", ".jpg", ".jpeg"}

// 	var filePath string
// 	var found bool

// 	// Iterate through possible extensions
// 	for _, ext := range extensions {
// 		potentialPath := fmt.Sprintf("./images/user/%s%s", name, ext)
// 		if _, err := os.Stat(potentialPath); err == nil {
// 			filePath = potentialPath
// 			found = true
// 			break
// 		}
// 	}

// 	// If no file was found, return an error
// 	if !found {
// 		c.JSON(400, gin.H{"error": "File not found"})
// 		return
// 	}

// 	// Attempt to remove the file
// 	err := os.Remove(filePath)
// 	if err != nil {
// 		c.JSON(500, gin.H{"error": "Server Error", "details": err.Error()})
// 		return
// 	}
// 	_, err = h.services.FileImage().ImageDelete(context.Background(), &coins_service.ImagePrimaryKey{Id: imageName})
// 	if err != nil {
// 		c.JSON(500, gin.H{"error": "Server Error", "details": err.Error()})
// 		return
// 	}
// 	// Return success response
// 	c.JSON(200, gin.H{"message": "Image deleted successfully"})
// }

// // GetFileURL godoc
// // @ID get_user_file_url
// // @Router /user/image/{id} [GET]
// // @Summary Get User Image URL
// // @Description Retrieve the URL for a specific user image
// // @Tags User
// // @Accept json
// // @Produce json
// // @Param id path string true "Image ID"
// // @Success 200 {object} http.Response{data=string} "Success"
// // @Response 400 {object} http.Response{data=string} "Bad Request"
// // @Response 404 {object} http.Response{data=string} "File not found"
// // @Failure 500 {object} http.Response{data=string} "Server Error"
// func (h *Handler) GetUserFileURL(c *gin.Context) {
// 	fileName := c.Param("id")
// 	fileName = strings.Replace(fileName, "-", "", -1)
// 	extensions := []string{".png", ".gif", ".jpg", ".jpeg"}

// 	var filePath, imageUrl string
// 	var found bool

// 	// Iterate through possible extensions
// 	for _, ext := range extensions {
// 		potentialPath := fmt.Sprintf("./images/user/%s%s", fileName, ext)
// 		link := fmt.Sprintf("user/image/%s%s", fileName, ext)
// 		if _, err := os.Stat(potentialPath); err == nil {
// 			filePath = potentialPath
// 			imageUrl = link
// 			found = true
// 			break
// 		}
// 	}
// 	// If no file was found, return an error
// 	if !found {
// 		c.JSON(400, gin.H{"error": "File type not found"})
// 		return
// 	}
// 	fmt.Println(filePath)

// 	// Faylning mavjudligini tekshirish
// 	if _, err := os.Stat(filePath); os.IsNotExist(err) {
// 		c.JSON(404, gin.H{"error": "File not found"})
// 		return
// 	}

// 	// Fayl URL'sini yaratish
// 	fileURL := fmt.Sprintf("http://localhost:8080/%s", imageUrl)

// 	// URL'ni frontendga qaytarish
// 	c.JSON(200, gin.H{"file_url": fileURL})
// }
