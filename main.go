package main

import (
	"log"
	"net/http"
	"os"
	"path/filepath"
	"strings"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	r.POST("/upload", func(c *gin.Context) {
		// Retrieve the file from the form-data
		file, err := c.FormFile("file")
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		// Check if the file has the .tar.gz extension
		ext := strings.ToLower(filepath.Ext(file.Filename))
		if ext != ".gz" {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid file format. Only .tar.gz files are allowed."})
			return
		}

		// Generate a unique file name
		filename := filepath.Base(file.Filename)

		// Save the file to disk
		if err := c.SaveUploadedFile(file, "./uploads/"+filename); err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}

		c.JSON(http.StatusOK,
			gin.H{
				"message":    "File uploaded successfully!",
				"fileName":   filename,
				"fileSize":   file.Size,
				"filePath":   "/uploads/" + filename,
				"directLink": "http://localhost:8080/files/" + filename,
			})
	})

	r.GET("/files/:filename", func(c *gin.Context) {
		filename := c.Param("filename")
		filepath := "./uploads/" + filename

		// Check if the file exists
		_, err := os.Stat(filepath)
		if os.IsNotExist(err) {
			c.JSON(http.StatusNotFound, gin.H{"error": "File not found."})
			return
		}

		// Serve the file for download
		c.File(filepath)
	})

	if err := r.Run(":8080"); err != nil {
		log.Fatal(err)
	}
}
