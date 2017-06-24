package main

import (
	"github.com/docktermj/go-hello-plugins/interfaces/hello"
	"github.com/hashicorp/go-plugin"
)

// ----------------------------------------------------------------------------
// Implementation of interface.
// ----------------------------------------------------------------------------

type Hello struct{}

func (Hello) Speak() string {
	return "Hello Universe!"
}

// ----------------------------------------------------------------------------
// Install and run plugin.
// ----------------------------------------------------------------------------

// Information to verify correct plugin.
var handshakeConfig = plugin.HandshakeConfig{
	ProtocolVersion:  1,
	MagicCookieKey:   "hello-english-cookie-key",
	MagicCookieValue: "hello-english-cookie-value",
}

// pluginMap is the map of plugins we can dispense.
var pluginMap = map[string]plugin.Plugin{
	"hello-english": &hello.Plugin{Impl: new(Hello)},
}

// Run the plugin.
func main() {
	plugin.Serve(&plugin.ServeConfig{
		HandshakeConfig: handshakeConfig,
		Plugins:         pluginMap,
	})
}
