package router

import "github.com/gin-gonic/gin"

func HandleCreateFile(c *gin.Context) {
    c.JSON(200, gin.H{ "message": "Hello world" })
}
