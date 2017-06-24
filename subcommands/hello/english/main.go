package english

// Reference: https://github.com/hashicorp/go-plugin/blob/master/examples/basic/main.go

import (
	"fmt"
	"io/ioutil"
	"log"
	"os/exec"

	"github.com/docktermj/go-hello-plugins/interfaces/hello"
	"github.com/hashicorp/go-plugin"
)

// Information to verify correct plugin.
var handshakeConfig = plugin.HandshakeConfig{
	ProtocolVersion:  1,
	MagicCookieKey:   "hello-english-cookie-key",
	MagicCookieValue: "hello-english-cookie-value",
}

// pluginMap is the map of plugins we can dispense.
var pluginMap = map[string]plugin.Plugin{
	"hello-english": &hello.Plugin{},
}

// The command.
func Command(argv []string) {

	// We don't want to see the plugin logs.
	log.SetOutput(ioutil.Discard)

	// We're a host! Start by launching the plugin process.
	client := plugin.NewClient(&plugin.ClientConfig{
		HandshakeConfig: handshakeConfig,
		Plugins:         pluginMap,
		Cmd:             exec.Command("./plugins/hello/english/english", "plugin"),
	})
	defer client.Kill()

	// Connect via RPC.
	rpcClient, err := client.Client()
	if err != nil {
		log.Fatal(err)
	}

	// Request the plugin.
	raw, err := rpcClient.Dispense("hello-english")
	if err != nil {
		log.Fatal(err)
	}

	// Call the plugin.
	myHello := raw.(hello.Hello)
	fmt.Println(myHello.Speak())
}
