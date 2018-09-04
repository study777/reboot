package pkgs

import (
	"context"
	"github.com/leopoldxx/go-utils/trace"
	"reboot/pkg/dao"
	"reboot/pkg/task/scheduler"
	"reboot/pkg/task/taskscheduler"
	"sync"
)

var (
	manager     *scheduler.Manager
	managerOnce sync.Once
)

func GetScheduler(sto dao.Storage) *scheduler.Manager {
	managerOnce.Do(func() {
		var err error
		ctx := trace.WithTraceForContext(context.TODO(), "task-scheduler")
		manager, err = scheduler.NewManager(ctx, sto)
		if err != nil {
			panic(err)
		}
		if err = manager.InitSchedulers(taskscheduler.Scheduler()); err != nil {
			panic(err)
		}
	})
	return manager
}
