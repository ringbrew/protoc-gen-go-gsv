package main

import (
	"flag"
	"github.com/ringbrew/protoc-gen-go-gsv/domain"
	"google.golang.org/protobuf/compiler/protogen"
)

func main() {
	var flags flag.FlagSet

	protogen.Options{
		ParamFunc: flags.Set,
	}.Run(domain.NewServiceGen().Generate)
}
