package handlers

import (
	"fmt"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
)

func AddSong(c *gin.Context) {
	file, err := c.FormFile("file")
	if err != nil {
		c.JSON(400, gin.H{"error": "File not found"})
		return
	}
	f, err := os.CreateTemp("./uploads", "audio-*.wav")
	if err != nil {
		c.JSON(400, gin.H{"error": "Failed to save file"})
		return
	}

	fmt.Println("File name :", f.Name())
	c.SaveUploadedFile(file, f.Name())
	c.JSON(http.StatusOK, gin.H{"message": "Added the song"})
}
