package main

import (
	"fmt"
	"log"
	"net/http"
	"path/filepath"

	"github.com/gin-gonic/gin"
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

    // Retrieve file information
    extension := filepath.Ext(file.Filename)
    // Generate random file name for the new uploaded file so it doesn't override the old file with same name
    newFileName := uuid.New().String() + extension

    // The file is received, so let's save it
    if err := c.SaveUploadedFile(file, "/some/path/on/server/" + newFileName); err != nil {
        c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
            "message": "Unable to save the file",
        })
        return
    }

    // File saved successfully. Return proper result
    c.JSON(http.StatusOK, gin.H{
        "message": "Your file has been successfully uploaded."
    })
}

func main() {
	r := gin.Default
	r.POST("/uploads/image",saveFileHandler())
	
	//http.ListenAndServe(":9000",r)
	
	r.Run()
}
