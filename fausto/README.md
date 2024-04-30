## Fausto

Fausto is the main component for the architecture, it is a mikrokernel that can have many plugins, you can find all of them inside the 'plugins' folder

### Service arch

The main idea is, to have an HTTP server listening on a port (specified by `config.yml`), and have a route that stores .txt files into the database (we are using a mongodb one).

Whenever we store said file, we are going to call all the registered plugins (you can check in the `DiscoverPlugins` inside `plugins/plugin.go`) with the new data, so all the plugins can do whatever they want with the data.

### How to add new plugins

Let's say we want to add a plugin that whenever the text gets inserted into the database, we want to log it into the console.

So first, we would add a new file into the `plugins` directory, let's say, log_file_content_plugin.go is the name of the file.

Inside it you can just add the following code:

```go
package plugins

type LogFileContentPlugin struct {}

func (*LogFileContentPlugin) Execute(data PluginInputData) {
    print(data.Content)
}

```

And then, inside the `plugins/plugin.go` file you can change the `DiscoverPlugins` function to be like this:

```go
func DiscoverPlugins() []Plugin {
	plugins := []Plugin{
		&LogFileContentPlugin{},
	}

	return plugins
}
```

Now, when you update a file, using the `/file` route, the file will be logged to stdout
