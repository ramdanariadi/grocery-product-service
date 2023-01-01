package utils

import (
	"fmt"
	"reflect"
)

func PanicIfError(err error) {
	if err != nil {
		fmt.Println("err ", err.Error())
		fmt.Println(reflect.TypeOf(err))
		panic(err)
	}
}
