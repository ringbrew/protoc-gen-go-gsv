package domain

const serviceGenImpl = `package [[.packageName]]

import (
	"github.com/ringbrew/gsv/service"
	"[[.module]]/export/example"
	"google.golang.org/grpc"
)

type [[.serviceName]] struct {
	[[.packageName]].Unimplemented[[.serviceName]]Server
}

func New[[.serviceName]]() service.Service {
	return &[[.serviceName]]{}
}

func (s *[[.serviceName]]) Name() string {
	return "[[.packageName]].[[.serviceName]]"
}

func (s *[[.serviceName]]) Remark() string {
	return ""
}

func (s *[[.serviceName]]) Description() service.Description {
	return service.Description{
		Valid:           true,
		GrpcServiceDesc: []grpc.ServiceDesc{[[.packageName]].[[.serviceName]]_ServiceDesc},
		GrpcGateway:     nil,
	}
}`
