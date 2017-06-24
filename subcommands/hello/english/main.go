package english

// Reference: https://github.com/hashicorp/go-plugin/blob/master/examples/basic/main.go

import (
	"fmt"
	"github.com/docktermj/go-hello-plugins/plugins/hello"
	"log"
	"os/exec"

	"github.com/hashicorp/go-plugin"
)

var handshakeConfig = plugin.HandshakeConfig{
	ProtocolVersion:  1,
	MagicCookieKey:   "BASIC_PLUGIN_ENGLISH",
	MagicCookieValue: "helloEnglish",
}

// pluginMap is the map of plugins we can dispense.
var pluginMap = map[string]plugin.Plugin{
	"helloEnglish": &hello.MyPlugin{},
}

func Command(argv []string) {

	// We're a host! Start by launching the plugin process.
	client := plugin.NewClient(&plugin.ClientConfig{
		HandshakeConfig: handshakeConfig,
		Plugins:         pluginMap,
		Cmd:             exec.Command("./plugins/hello/english/english", "plugin"),
	})
	defer client.Kill()

	// Connect via RPC
	rpcClient, err := client.Client()
	if err != nil {
		log.Fatal(err)
	}

	// Request the plugin
	raw, err := rpcClient.Dispense("englishHello")
	if err != nil {
		log.Fatal(err)
	}

	// We should have a Greeter now! This feels like a normal interface
	// implementation but is in fact over an RPC connection.
	myHello := raw.(hello.Hello)
	fmt.Println(myHello.Speak())
}
