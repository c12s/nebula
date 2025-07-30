package handlers

import (
	"context"
	"github.com/c12s/nebula/pkg/api"
	"github.com/c12s/nebula/internal/storage"
)


type NebulaGrpcHandler struct {
	api.UnimplementedNebulaServer
	Storage storage.Storage
}

func NewNebulaGrpcHandler(storage storage.Storage) api.NebulaServer {
	return NebulaGrpcHandler{}
}


func (m NebulaGrpcHandler) AddDeployment(ctx context.Context, req *api.AddDeploymentReq) (*api.AddDeploymentResp, error) {
	return &api.AddDeploymentResp{}, nil
}
