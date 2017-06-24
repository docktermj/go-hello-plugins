package main

import (
	"github.com/docktermj/go-hello-plugins/interfaces/hello"
	"github.com/hashicorp/go-plugin"
)

// Here is a real implementation of Greeter
type GreeterHello struct{}

func (GreeterHello) Speak() string {
	return "Hi from Italy!"
}

// handshakeConfigs are used to just do a basic handshake between
// a plugin and host. If the handshake fails, a user friendly error is shown.
// This prevents users from executing bad plugins or executing a plugin
// directory. It is a UX feature, not a security feature.
var handshakeConfig = plugin.HandshakeConfig{
	ProtocolVersion:  1,
	MagicCookieKey:   "BASIC_PLUGIN_ITALY",
	MagicCookieValue: "hello-italy",
}

// pluginMap is the map of plugins we can dispense.
var pluginMap = map[string]plugin.Plugin{
	"greeter": &hello.HelloPlugin{Impl: new(GreeterHello)},
}

func main() {
	plugin.Serve(&plugin.ServeConfig{
		HandshakeConfig: handshakeConfig,
		Plugins:         pluginMap,
	})
}
