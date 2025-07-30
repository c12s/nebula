package service

import (
	"context"
	"github.com/c12s/nebula/pkg/api"
)


type NebulaGrpcHandler struct {
	api.UnimplementedNebulaServer
}

func NewMeridianGrpcHandler() api.MeridianServer {
	return MeridianGrpcHandler{}
}


func (m NebulaGrpcHandler) AddDeployment(ctx context.Context, req *api.AddDeploymentReq) (*api.AddDeploymentResp, error) {
	return &api.AddDeploymentResp{}, nil
}
