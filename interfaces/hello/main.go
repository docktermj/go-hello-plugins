package hello

import (
	"net/rpc"

	"github.com/hashicorp/go-plugin"
)

// ----------------------------------------------------------------------------
// Interface
// ----------------------------------------------------------------------------

type Hello interface {
	Speak() string
}

// ----------------------------------------------------------------------------
// RPC
// ----------------------------------------------------------------------------

type HelloRPC struct{ client *rpc.Client }

func (g *HelloRPC) Speak() string {
	var resp string
	err := g.client.Call("Plugin.Speak", new(interface{}), &resp)
	if err != nil {
		panic(err)
	}

	return resp
}

// ----------------------------------------------------------------------------
// RPC
// ----------------------------------------------------------------------------

type HelloRPCServer struct {
	Impl Hello
}

func (s *HelloRPCServer) Speak(args interface{}, resp *string) error {
	*resp = s.Impl.Speak()
	return nil
}

// ----------------------------------------------------------------------------
// Plugin
// ----------------------------------------------------------------------------

type HelloPlugin struct {
	Impl Hello
}

func (p *HelloPlugin) Server(*plugin.MuxBroker) (interface{}, error) {
	return &HelloRPCServer{Impl: p.Impl}, nil
}

func (HelloPlugin) Client(b *plugin.MuxBroker, c *rpc.Client) (interface{}, error) {
	return &HelloRPC{client: c}, nil
}
