package plugins

import (
	"sync"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

var pluginList = []Plugin{
	ProfanityDetectorNew(),
	SpellCheckerNew(),
	WordCountNew(),
}

var pluginsInitialized bool = false

type PluginInputData struct {
	Content string
	Id      primitive.ObjectID
}

type Plugin interface {
	Execute(PluginInputData)
	Init()
}

func InitPlugins() {
	var wg sync.WaitGroup
	for _, p := range pluginList {
		wg.Add(1)
		go func(pl Plugin) {
			defer wg.Done()
			pl.Init()
		}(p)
	}

	wg.Wait()
	pluginsInitialized = true
}

func RunPlugins(data PluginInputData) {
	if !pluginsInitialized {
		InitPlugins()
	}

	var wg sync.WaitGroup
	for _, p := range pluginList {
		wg.Add(1)

		go func(pl Plugin) {
			defer wg.Done()
			pl.Execute(data)
		}(p)
	}

	wg.Wait()
}
