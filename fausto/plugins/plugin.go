package plugins

import (
	"sync"
)

type Plugin interface {
    Init(interface{}) 
    Execute()
}

func DiscoverPlugins() []Plugin {
    plugins := []Plugin{}

    // TODO: Add the plugns to the list

    return plugins
}

func InitializePlugins(pluginList []Plugin) {
    var wg sync.WaitGroup
    for _, p := range(pluginList) {
        wg.Add(1)

        go func(pl Plugin) {
            defer wg.Done()
            pl.Execute()
        }(p)
    }

    wg.Wait()
}

