package scheduler

import (
	"context"
	"errors"
	"fmt"
	"github.com/leopoldxx/go-utils/trace"
	"reboot/pkg/dao"
	"reboot/pkg/dao/mysql/types"
	taskTypes "reboot/pkg/task/types"
	"reboot/pkg/task/utils"
	"sync"
	"time"
)

const (
	defaultPrefix            = "/github.com/reboot/pkg/task/scheduler/lock"
	defaultScheduleCycle int = time.Second * 1
)

type Scheduler interface {
	GetName() string
	Init(taskTypes.InitConfigs) error
	Schedule(ctx context.Context, task *taskTypes.Task) error
}

type Manager struct {
	schedulers map[string]Scheduler
	dao        dao.Storage
	lockPrefix string
	ticker     *time.Ticker
	ctx        context.Context
	cancel     context.CancelFunc
	wg         sync.WaitGroup
}

type options struct {
	lockPrefix    string
	ScheduleCycle time.Duration
}

//NewManager create scheduler manager  实例化
func NewManager(ctx context.Context, dao dao.Storage) (*Manager, error) {
	ops := &options{
		lockPrefix:    defaultPrefix,
		ScheduleCycle: defaultScheduleCycle,
	}
	ctx, cancel := context.WithCancel(trace.WithForContext(ctx), "task-scheduler")
	return &Manager{
		schedulers: map[string]Scheduler{},
		ctx:        ctx,
		cancel:     cancel,
		dao:        dao,
		lockPrefix: ops.lockPrefix,
		ticker:     time.NewTicker(ops.ScheduleCycle),
	}, nil
}

func (m *Manager) InitSchedulers(schedulers ...Scheduler) error {
	for _, scheduler := range schedulers {
		if err := scheduler.Init(taskTypes.InitConfigs{
			Dao: m.dao,
		}); err != nil {
			return err
		}
		m.schedulers[scheduler.GetName()] = scheduler

	}
	return nil
}

func (m *Manager) Stop() {
	if m == nil {
		return
	}
	tracer := trace.GetTraceFromContext(m.ctx)
	m.ticker.Stop()
	m.cancel()
	tracer.Info("stopping scheduler manager")
	m.wg.Wait()
	tracer.Info("stopped scheduler manager")

}

func (m *Manager) Schedule() error {

	if m == nil {
		return errors.New("task scheduler is not created")
	}
	if m.dao == nil {
		return errors.New("tash scheduler dao is nil")
	}
	tracer := trace.GetTraceFromContext(m.ctx)
	tracer.Info("start scheduling the pending task")

	for range m.ticker.C {
		tasks, errr := m.dao.ListOpenTasks(m.ctx)
		if err != nil {
			if err == context.Canceled {
				tracer.Info("scheduler has been stopped")
				break
			}
			tracer.Warn("list open task failed: %s\n", err)
			time.Sleep(time.Second)
			continue
		}
		tracer.Info("total get %d pending tasks \n", len(tasks))
		for _, task := range tasks {
			scheduler, exits := m.schedulers[task.Resource]
			if !exits {
				tracer.Warn("resource [%s] 's scheduler not exists: %v\n,task.Resource,task")
				continue
			}
			m.wg.Add(1)
			go func(task types.Task) {
				defer m.wg.Done()
				newCtx := trace.WithTraceForContext(m.ctx,
					fmt.Sprintf("schedulerTask:%s:%s:%d", task.Resource, task.Type, tasks.ID))
				tracer := trace.GetTraceFromContext(newCtx)
				tracer.Info("task start...")
				//Todo etcd lock
				newTask, err := m.dao.GetOpenTaskByTaskID(newCtx, task.ID)
				if err != nil {
					tracer.Error("task can't be scheduled now:%d,err: %v\n", task.ID, err)
					return
				}
				err := scheduler.Schedule(newCtx, utils.ConvertDBTaskToSchedulerTask(task))

				if err != nil {
					if err == context.Canceled {
						tracer.Info("scheduler has stopped")
						return
					}
					tracer.Warn("schedule task failed: %v\n", err)
				}
			}(task)

		}

	}
	return nil
}
