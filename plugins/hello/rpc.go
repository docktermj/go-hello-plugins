package hello

// Reference: https://github.com/hashicorp/go-plugin/blob/master/examples/basic/commons/greeter_interface.go

import (
	"net/rpc"

	"github.com/hashicorp/go-plugin"
)

type Rpc struct{ client *rpc.Client }

func (thisType *Rpc) Speek() string {
	var resp string
	err := thisType.client.Call("Plugin.Speak", new(interface{}), &resp)
	if err != nil {
		panic(err)
	}
	return resp
}
