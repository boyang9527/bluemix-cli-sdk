package main

import (
	"github.com/IBM-Cloud/ibm-cloud-cli-sdk/plugin"

	"fmt"
)

type DemoPlugin struct{}

func main() {
	plugin.Start(new(DemoPlugin))
}

func (n *DemoPlugin) Run(context plugin.PluginContext, args []string) {
	switch args[0] {
	case "list":
		fmt.Println("Running command 'list'.")
	case "show":
		fmt.Println("Running command 'show'.")
	}
}

func (n *DemoPlugin) GetMetadata() plugin.PluginMetadata {
	return plugin.PluginMetadata{
		Name: "demo-plugin",
		Version: plugin.VersionType{
			Major: 0,
			Minor: 0,
			Build: 1,
		},
		Namespaces: []plugin.Namespace{
			plugin.Namespace{
				Name:        "demo",
				Description: "Plugin demonstration",
			},
		},
		Commands: []plugin.Command{
			{
				Namespace:   "demo",
				Name:        "list",
				Description: "List resources",
				Usage:       "ibmcloud demo list",
			},
			{
				Namespace:   "demo",
				Name:        "show",
				Description: "Show the details of a resource",
				Usage:       "ibmcloud demo show",
			},
		},
	}
}
