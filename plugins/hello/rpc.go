package hello

// Reference: https://github.com/hashicorp/go-plugin/blob/master/examples/basic/commons/greeter_interface.go

import (
	"fmt"
	"net/rpc"
)

type Rpc struct{ client *rpc.Client }

func (thisType *Rpc) Speak() string {
	fmt.Println("In Rpc.Speak")
	var resp string
	err := thisType.client.Call("Plugin.Speak", new(interface{}), &resp)
	if err != nil {
		panic(err)
	}
	return resp
}
