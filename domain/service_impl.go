package domain

const serviceGenImpl = `package [[.packageName]]

import (
	"github.com/ringbrew/gsv/service"
	"[[.module]]/export/[[.packageName]]"
	"[[.module]]/internal/domain"
	"google.golang.org/grpc"
)

type [[.serviceName]] struct {
	ctx *domain.UseCaseContext
	[[.packageName]].Unimplemented[[.protoServiceName]]Server
}

func New[[.serviceName]](ctx *domain.UseCaseContext) service.Service {
	return &[[.serviceName]]{
		ctx: ctx,
	}
}

func (s *[[.serviceName]]) Name() string {
	return "[[.packageName]].[[.protoServiceName]]"
}

func (s *[[.serviceName]]) Remark() string {
	return ""
}

func (s *[[.serviceName]]) Description() service.Description {
	return service.Description{
		Valid:           true,
		GrpcServiceDesc: []grpc.ServiceDesc{[[.packageName]].[[.protoServiceName]]_ServiceDesc},
		GrpcGateway:     nil,
	}
}`

const serviceDefineImpl = `// Code generated by protoc-gen-go-gsv. DO NOT EDIT.
package [[.packageName]]

const (
	[[.serviceName]]Name = "[[.packageName]].[[.protoServiceName]]"
)
`
