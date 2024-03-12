package router

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type CreateFileRequest struct {
    Content string `json:"content"`
}

func HandleCreateFile(c *gin.Context) {
    var requestData CreateFileRequest
    
    if err := c.BindJSON(&requestData); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }

    c.JSON(200, gin.H{ "message": requestData.Content })
}
