package main

import proto "github.com/protocol-buffers/go-proto-code/proto-version-go-code/proto-version-go-code"


func main() {
id := proto.Id{Value: 1}
println(id.String())
}