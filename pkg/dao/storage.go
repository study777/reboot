package dao

import (
	"context"
	"reboot/pkg/dao/mysql/types"
)

type Storage interface {
	CreateTask(ctx context.Context, task *types.Task) (int64, error)

	GetTask(ctx context.Context)
}
