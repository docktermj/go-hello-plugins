package hello

// Reference: https://github.com/hashicorp/go-plugin/blob/master/examples/basic/commons/greeter_interface.go

import (
	"net/rpc"

	"github.com/hashicorp/go-plugin"
)


type RpcServer struct {
	Impl Hello
}

func (thisType *RpcServer) Greet(args interface{}, resp *string) error {
	*resp = thisType.Impl.Greet()
	return nil
}
