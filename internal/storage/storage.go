package storage

import (
	"context"
	"github.com/c12s/nebula/pkg/api"
)


type Storage interface {
	Create(ctx context.Context, req *api.AddDeploymentReq) error
	Init()
	Close()
}
