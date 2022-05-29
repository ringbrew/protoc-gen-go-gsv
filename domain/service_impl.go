package domain

const serviceGenImpl = `package [[.packageName]]

import (
	"github.com/ringbrew/gsv/service"
	"github.com/ringbrew/gsv/service/example/export/example"
	"google.golang.org/grpc"
)

type [[.serviceName]] struct {
	[[.packageName]].Unimplemented[[.serviceName]]Server
}

func NewService() service.Service {
	return &Service{}
}

func (s *Service) Name() string {
	return "[[.serviceName]]"
}

func (s *Service) Remark() string {
	return ""
}

func (s *Service) Description() service.Description {
	return service.Description{
		Valid:           true,
		GrpcServiceDesc: []grpc.ServiceDesc{[[.packageName]].[[.serviceName]]_ServiceDesc},
		GrpcGateway:     nil,
	}
}`
