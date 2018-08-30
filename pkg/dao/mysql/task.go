package mysql

import (
	"context"
	"fmt"
	"reboot/pkg/dao/mysql/types"
)

//CreateTask  insert db
func (m *mysql) CreateTask(ctx context.Context, task *types.Task) (int64, error) {
	return 0, nil
}

func (m *mysql) GetTask(ctx context.Context) {
	fmt.Println("GetTask")
}
