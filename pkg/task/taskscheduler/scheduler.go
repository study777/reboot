package taskscheduler

import (
	"reboot/pkg/dao"
	"reboot/pkg/task/scheduler"
	schedTypes "reboot/pkg/task/scheduler/types"
)

type taskScheduler struct {
	dao dao.Storage
}

var (
	task = taskScheduler{}
)

func Scheduler() scheduler.Scheduler {
	return &task
}

func (sched *taskScheduler) GetName() string {
	return string("task")
}

func (sched *taskScheduler) Init(cfg, schedTypes.InitConfigs) error {
	if sched == nil {
		return errors.New("task sched is nil")
	}
	sched.dao = cfg.Dao
	return nil
}

func (sched *taskScheduler) Schedule(ctx context.Context, task *schedTypes.Task) error {
	tracer := trace.GetTraceFromContext(ctx)
	tracer.Info("Get a task,%v\n", *task)
	return nil

}
