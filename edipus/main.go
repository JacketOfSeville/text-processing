package main

import (
	"fmt"
	"html/template"

	"github.com/Gustrb/text-processing/edipus/handlers"
	"github.com/Gustrb/text-processing/edipus/utils"
	"github.com/gin-gonic/gin"
)

const (
	Host           = "0.0.0.0"
	Port           = "8081"
	FileUploadPath = "public/file-upload.html"
)

func main() {
	r := gin.Default()
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "Pong :)",
		})
	})

	fileUploadContent, err := utils.ReadFileFromPath(FileUploadPath)
	if err != nil {
		fmt.Printf("Failed to load template file: %s, because: %v", FileUploadPath, err)
		return
	}

	tmpl, err := template.New("file-upload").Parse(string(fileUploadContent))
	if err != nil {
		fmt.Printf("Failed to parse template content: %s, because: %v", string(fileUploadContent), err)
		return
	}

	r.GET("/", func(c *gin.Context) {
		// render the template
		tmpl.Execute(c.Writer, nil)
	})

	r.POST("files/upload", handlers.UploadFileHandler)

	r.Run(fmt.Sprintf("%s:%s", Host, Port))
}
