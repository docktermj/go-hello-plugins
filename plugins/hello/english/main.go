package main

import (
	"fmt"

	"github.com/docktermj/go-hello-plugins/plugins/hello"
	"github.com/hashicorp/go-plugin"
)

// Implementation of interface.

type Hello struct{}

func (Hello) Speak() string {
	fmt.Println("Hello, Universe!")
	return "Hello English return!"
}

// Install plugin.

var handshakeConfig = plugin.HandshakeConfig{
	ProtocolVersion:  1,
	MagicCookieKey:   "BASIC_PLUGIN_ENGLISH",
	MagicCookieValue: "helloEnglish",
}

var pluginMap = map[string]plugin.Plugin{
	"helloEnglish": &hello.MyPlugin{Impl: new(Hello)},
}

func main() {
	fmt.Println(">>>> plugin start")	
	plugin.Serve(&plugin.ServeConfig{
		HandshakeConfig: handshakeConfig,
		Plugins:         pluginMap,
	})
	fmt.Println(">>>> plugin")
}
