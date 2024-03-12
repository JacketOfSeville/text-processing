package main

import (
	"github.com/Gustrb/text-processing/fausto/plugins"
)

func main() {
    pluginList := plugins.DiscoverPlugins()

    plugins.InitializePlugins(pluginList)

    // TODO: Start the router
}

