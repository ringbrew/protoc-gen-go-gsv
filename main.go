package main

import (
	"github.com/ringbrew/protoc-gen-go-gsv/domain"
	"google.golang.org/grpc/benchmark/flags"
	"google.golang.org/protobuf/compiler/protogen"
)

func main() {
	var ps = domain.GetParamSet()
	protogen.Options{
		ParamFunc: ps.Set,
	}.Run(domain.NewServiceGen().Generate)
}
