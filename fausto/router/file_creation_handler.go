package router

import (
	"net/http"

	"github.com/Gustrb/text-processing/fausto/plugins"
	"github.com/Gustrb/text-processing/fausto/store"
	"github.com/gin-gonic/gin"
)

type CreateFileRequest struct {
	Content string `json:"content"`
}

func HandleCreateFile(c *gin.Context) {
	var requestData CreateFileRequest

	pluginList := plugins.DiscoverPlugins()

	if err := c.BindJSON(&requestData); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	data := &store.CreateFileDTO{Content: requestData.Content}

	if err := store.GetStore().FileStore().CreateFile(data); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	plugins.RunPlugins(pluginList, plugins.PluginInputData{Content: requestData.Content, Id: data.Id})

	c.JSON(200, gin.H{"file": data})
}
