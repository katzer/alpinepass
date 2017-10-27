package main

import "github.com/appPlant/alpinepass/src/app"
import "github.com/appPlant/alpinepass/src/plugin"

func main() {
	plugin.ExecutePlugin() //TODO Remove this temporary call.
	app.RunApp()
}
