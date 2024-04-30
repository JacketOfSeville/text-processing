package handlers

import (
	"bytes"
	"encoding/json"
	"io"
	"net/http"

	"github.com/Gustrb/text-processing/edipus/extractors"
	"github.com/Gustrb/text-processing/edipus/models"
	"github.com/Gustrb/text-processing/edipus/utils"

	"github.com/gin-gonic/gin"
)

const (
	FaustoURL = "http://fausto:8080"
)

type FaustoRequest struct {
	Content string `json:"content"`
	Name    string `json:"name"`
}

func sendToFausto(content string, fileToProcess *models.FileToProcess, c *gin.Context) {
	// Perform a POST request to the /file endpoint of the Fausto service
	// with the content of the file
	faustoRequest := FaustoRequest{
		Content: content,
		Name:    fileToProcess.Name,
	}

	// Serialize the FaustoRequest struct to JSON
	data, err := json.Marshal(faustoRequest)
	if err != nil {
		c.JSON(500, gin.H{"error": "Failed to serialize Fausto request"})
		return
	}

	// Send the request to Fausto
	httpResponse, err := http.Post(FaustoURL+"/file", "application/json", bytes.NewReader(data))
	if err != nil {
		c.JSON(500, gin.H{"error": "Failed to send request to Fausto"})
		return
	}

	defer httpResponse.Body.Close()
}

// UploadFileHandler receives a file through multipart form and processes it
// to extract the text content, and pass it to the text processing service.
//
// This function should be extension-agnostic, meaning that it should be able
// to process .txt, .docx, etc
func UploadFileHandler(c *gin.Context) {
	// first we need to get the file from the request
	fileHeader, err := c.FormFile("file")
	if err != nil {
		c.JSON(400, gin.H{"error": "No file uploaded"})
		return
	}

	file, err := fileHeader.Open()
	defer file.Close()

	if err != nil {
		c.JSON(500, gin.H{"error": "Failed to open file"})
		return
	}

	// Read the file content
	fileContent, err := io.ReadAll(file)

	if err != nil {
		c.JSON(500, gin.H{"error": "Failed to read file content"})
		return
	}

	// extract the file extension
	extension, err := utils.ExtractExtension(fileHeader.Filename)
	if err != nil {
		c.JSON(500, gin.H{"error": "Failed to extract file extension"})
		return
	}

	fileToProcess := models.FileToProcess{
		Content: fileContent,
		Name:    fileHeader.Filename,
	}

	var content string

	switch extension {
	case ".txt":
		content, err = extractors.ExtractTextFromTxt(fileToProcess)

	case ".docx":
		content, err = extractors.ExtractTextFromDocx(fileToProcess)
	}

	if err != nil {
		c.JSON(500, gin.H{"error": "Failed to extract text from file"})
		return
	}

	sendToFausto(content, &fileToProcess, c)

	c.JSON(200, gin.H{"message": "File uploaded"})
}
