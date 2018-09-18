package taskscheduler

import (
	"context"
	"errors"
	"github.com/leopoldxx/go-utils/trace"
	"reboot/pkg/dao"
	"reboot/pkg/enum"
	"reboot/pkg/task/scheduler"
	taskTypes "reboot/pkg/task/taskscheduler/types"
	schedTypes "reboot/pkg/task/types"
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

func (sched *taskScheduler) Init(cfg schedTypes.InitConfigs) error {
	if sched == nil {
		return errors.New("task sched is nil")
	}
	sched.dao = cfg.Dao
	return nil
}

func (sched *taskScheduler) Schedule(ctx context.Context, task *schedTypes.Task) error {
	tracer := trace.GetTraceFromContext(ctx)
	tracer.Infof("Get a task,%v\n", *task)
	t, err := convertSchedTaskToTask(task)
	if err != nil {
		tracer.Error("convert task failed %v\n", err)
		return err
	}
	if len(t.Status.State) == 0 {
		t.Status.State = enum.TaskPending
		t.Status.TryTimes = 0

	}
	tracer.Infof("current state is %v\n", t.Status.State)

	switch t.Status.State {
	case enum.TaskPending:
		return sched.TaskPending(ctx, t)
	case enum.TaskDoing:
		return sched.TaskDoing(ctx, t)
	case enum.TaskDone:
		return sched.TaskDone(ctx, t)
	default:
		tracer.Errorf("unknown status %s of task %v\n", t.Status.State.String(), *task)
		return errors.New("unknown status of the task")
	}

}

func (sched *taskScheduler) TaskPending(ctx context.Context, t *taskTypes.Task) error {
	tracer := trace.GetTraceFromContext(ctx)
	tracer.Info("call TaskPending func")
	/*
	   Todo: xxx
	*/
	t.Status.State = enum.TaskDoing
	return sched.updateTaskStatus(ctx, t)
}

func (sched *taskScheduler) TaskDoing(ctx context.Context, t *taskTypes.Task) error {
	tracer := trace.GetTraceFromContext(ctx)
	tracer.Info("call TaskDoing func")
	/*
	   Todo: xxx
	*/
	t.Status.State = enum.TaskDone
	return sched.updateTaskStatus(ctx, t)
}

func (sched *taskScheduler) TaskDone(ctx context.Context, t *taskTypes.Task) error {
	tracer := trace.GetTraceFromContext(ctx)
	tracer.Info("call TaskDone func")
	/*
	   Todo: xxx
	*/
	t.Common.Close = true
	return sched.updateTaskStatus(ctx, t)
}

func (sched *taskScheduler) updateTaskStatus(ctx context.Context, task *taskTypes.Task) error {
	tracer := trace.GetTraceFromContext(ctx)
	dbtask := convertTaskToDBTask(task)
	err := sched.dao.UpdateTask(ctx, dbtask)
	if err != nil {
		tracer.Error(err)
		return err
	}
	return nil
}
