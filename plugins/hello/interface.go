package hello

// Reference: https://github.com/hashicorp/go-plugin/blob/master/examples/basic/commons/greeter_interface.go

// List of methods that a "Hello" plugin needs to implement.
type Hello interface {
	Speak() string
}
