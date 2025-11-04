package main

import (
	"fmt"

	"github.com/protocol-buffers/go-proto-code/proto-version-go-code-custom/proto_version_go_code_custom"
	"google.golang.org/protobuf/proto"
)

func main(){

	msg:=&proto_version_go_code_custom.Info{
		Name: "Hello",
		Age: proto.Uint64(0),
		Location: "Nepal",
	}

	// data,_:=proto.Marshal(msg)


	// fmt.Println(msg.FavouriteFood) //Unset field and non optional field has zero value in proto 3
	// fmt.Println(msg.Age) //Unset field and non optional field has zero value in proto 3

		fmt.Println(msg.FavouriteFood) //When optional then FavouriteFood is a pointer
	fmt.Println(*(msg.Age)) //Age is optional so is a pointer so need to deference it.
	
}