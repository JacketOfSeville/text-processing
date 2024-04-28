package main

import (
	"fmt"
	"log"

	"github.com/Gustrb/text-processing/fausto/config"
	"github.com/Gustrb/text-processing/fausto/plugins"
	"github.com/Gustrb/text-processing/fausto/router"
	"github.com/Gustrb/text-processing/fausto/store"
	"github.com/gin-gonic/gin"
)

func main() {
	config, err := config.ConfigFromFile("config.yml")
	if err != nil {
		panic(err)
	}

	err = store.Initialize(config.Database.Uri)
	if err != nil {
		panic(err)
	}

	log.Println("Connected to the database")

	defer store.Disconnect()

	plugins.InitPlugins()

	r := gin.Default()
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "Pong :)",
		})
	})

	r.POST("/file", router.HandleCreateFile)

	r.Run(fmt.Sprintf("%s:%s", config.Server.Host, config.Server.Port))
}
