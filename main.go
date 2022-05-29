package main

import (
	"flag"
	"github.com/ringbrew/protoc-gen-go-gsv/domain"
	"google.golang.org/protobuf/compiler/protogen"
)

func main() {
	sg := domain.NewServiceGen()
	var flags flag.FlagSet
	flags.String("module", "", "")

	protogen.Options{
		ParamFunc: flags.Set,
	}.Run(sg.Generate)
}
