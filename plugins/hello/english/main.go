package main

import (
	"github.com/hashicorp/go-plugin"
	"github.com/docktermj/go-hello-plugins/interfaces/hello"	
)

// Implementation of interface.

type Hello struct{}

func (Hello) Speak() string {
	return "Hello Universe!"
}

// Install plugin.

var handshakeConfig = plugin.HandshakeConfig{
	ProtocolVersion:  1,
	MagicCookieKey:   "BASIC_PLUGIN_ENGLISH",
	MagicCookieValue: "hello-english",
}

var pluginMap = map[string]plugin.Plugin{
	"hello-english": &hello.HelloPlugin{Impl: new(Hello)},
}

func main() {
	plugin.Serve(&plugin.ServeConfig{
		HandshakeConfig: handshakeConfig,
		Plugins:         pluginMap,
	})
}
