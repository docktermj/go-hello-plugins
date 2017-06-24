package hello

// Reference: https://github.com/hashicorp/go-plugin/blob/master/examples/basic/commons/greeter_interface.go

import (
	"net/rpc"

	"github.com/hashicorp/go-plugin"
)

type Plugin struct {
	Impl Hello
}

func (thisType *Plugin) Server(*plugin.MuxBroker) (interface{}, error) {
	return &RpcServer{Impl: thisType.Impl}, nil
}

func (Plugin) Client(b *plugin.MuxBroker, c *rpc.Client) (interface{}, error) {
	return &Rpc{client: c}, nil
}
