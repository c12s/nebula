package etcd

import (
	"context"
	"github.com/c12s/nebula/internal/storage"
	"github.com/c12s/nebula/pkg/api"
)

type Etcd struct {
}

func NewEtcdStorage() (error, storage.Storage) {
	return nil, &Etcd{}
}

func (etcd *Etcd) Create(ctx context.Context, req *api.AddDeploymentReq) error {
	return nil
}

func (etcd *Etcd) Init() {}

func (etcd *Etcd) Close() {}
