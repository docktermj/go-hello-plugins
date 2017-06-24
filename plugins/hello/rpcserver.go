package hello

// Reference: https://github.com/hashicorp/go-plugin/blob/master/examples/basic/commons/greeter_interface.go

import (
	"fmt"
)

type RpcServer struct {
	Impl Hello
}

func (thisType *RpcServer) Speak(args interface{}, resp *string) error {
	fmt.Println("In rpcserver.Speak")
	*resp = thisType.Impl.Speak()
	return nil
}
