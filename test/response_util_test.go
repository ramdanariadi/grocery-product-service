package test

import (
	"fmt"
	"github.com/ramdanariadi/grocery-product-service/main/setup"
	"testing"
)

type foo struct {
	a string
}

func TestResponseUtil(t *testing.T) {
	var foos []foo
	foos = append(foos, foo{
		a: "a",
	})
	status, message := setup.ResponseForQuerying(len(foos) > 0)
	fmt.Println(status)
	fmt.Println(message)
}
