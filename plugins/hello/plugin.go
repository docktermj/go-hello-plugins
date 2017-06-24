package hello

// Reference: https://github.com/hashicorp/go-plugin/blob/master/examples/basic/commons/greeter_interface.go

import (
	"net/rpc"
	"fmt"

	"github.com/hashicorp/go-plugin"
)

type MyPlugin struct {
	Impl Hello
}

func (thisType *MyPlugin) Server(*plugin.MuxBroker) (interface{}, error) {
	fmt.Println("In MyPlugin.Server")	
	return &RpcServer{Impl: thisType.Impl}, nil
}

func (MyPlugin) Client(broker *plugin.MuxBroker, client *rpc.Client) (interface{}, error) {
	fmt.Println("In MyPlugin.Client")	
	return &Rpc{client: client}, nil
}
