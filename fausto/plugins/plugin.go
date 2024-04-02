package plugins

import (
	"sync"
)

type PluginInputData struct {
	Content string
}

type Plugin interface {
	Execute(PluginInputData)
}

func DiscoverPlugins() []Plugin {
	plugins := []Plugin{}

	return plugins
}

func RunPlugins(pluginList []Plugin, data PluginInputData) {
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
