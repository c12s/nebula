package handlers

import (
	"context"
	"github.com/c12s/nebula/pkg/api"
)


type NebulaGrpcHandler struct {
	api.UnimplementedNebulaServer
}

func NewNebulaGrpcHandler() api.NebulaServer {
	return NebulaGrpcHandler{}
}


func (m NebulaGrpcHandler) AddDeployment(ctx context.Context, req *api.AddDeploymentReq) (*api.AddDeploymentResp, error) {
	return &api.AddDeploymentResp{}, nil
}
