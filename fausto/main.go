package main

import (
	"github.com/Gustrb/text-processing/fausto/plugins"
    "github.com/gin-gonic/gin"
)

func main() {
    pluginList := plugins.DiscoverPlugins()

    plugins.InitializePlugins(pluginList)

    r := gin.Default()
    r.GET("/ping", func (c *gin.Context) {
        c.JSON(200, gin.H{
            "message": "Pong :)",
        })
    })

    r.Run()
}

