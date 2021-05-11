package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"path"
	"path/filepath"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

const (
	UPLOAD_DIR = "uploads"
)

func saveFileHandler(c *gin.Context) {
	file, err := c.FormFile("file")

	// The file cannot be received.
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"message": "No file is received",
		})
		return
	}

	//generating saving path
	now := time.Now().UTC()
	folderHierarchy := fmt.Sprintf("%d/%02d/%02d", now.Year(), now.Month(), now.Day()) // splitted dateStr
	storingPath := path.Join(UPLOAD_DIR, folderHierarchy)
	os.MkdirAll(storingPath, os.ModePerm)

	// Retrieve file information
	extension := filepath.Ext(file.Filename)
	// Generate random file name for the new uploaded file so it doesn't override the old file with same name
	newFileName := uuid.New().String() + extension

	destination := path.Join(storingPath, newFileName)

	// The file is received, so let's save it
	if err := c.SaveUploadedFile(file, destination); err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
			"message": "Unable to save the file",
		})
		return
	}

	// formatting file hierarchy
	fileEndpointPath := fmt.Sprintf("%d-%02d-%02d-%s", now.Year(), now.Month(), now.Day(), newFileName)

	// File saved successfully. Return proper result
	c.JSON(http.StatusOK, gin.H{
		"message": fileEndpointPath,
	})
}

// This function is stand for doing some pre-required process
func seed() {
	workingDir, err := os.Getwd()
	if err != nil {
		log.Fatal(err)
	}
	uploadDir := path.Join(workingDir, UPLOAD_DIR)
	os.MkdirAll(uploadDir, os.ModePerm)
}

func main() {
	r := gin.Default()
	r.POST("/uploads/image", saveFileHandler)

	seed()

	r.Run()
}
