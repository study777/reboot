package service

import (
	"context"
	"errors"
	"github.com/leopoldxx/go-utils/trace"
	"reboot/pkg/dao/mysql/types"
)

func (s *service) GetTask(ctx context.Context) error {
	tracer := trace.GetTraceFromContext(ctx)
	tracer.Info("s.GetTask")
	return nil
}

func (s *service) ListTask(ctx context.Context) error {
	tracer := trace.GetTraceFromContext(ctx)
	tracer.Info("s.ListTask")
	return nil
}

func (s *service) CreateTask(ctx context.Context, namespace string, resource string) (*types.Task, error) {
	tracer := trace.GetTraceFromContext(ctx)
	tracer.Info("s.CreateTask")

	task := &types.Task{
		NameSpace: namespace,
		Resource:  resource,
		Type:      "create",
	}
	if s.opt.DB == nil {
		tracer.Error("db is nil")
		return task, errors.New("db is nil")
	}

	id, err := s.opt.DB.CreateTask(ctx, task)
	if err != nil {
		tracer.Error(err)
		return task, err
	}
	task.ID = id
	return task, nil
}

func (s *service) UpdateTask(ctx context.Context) error {
	tracer := trace.GetTraceFromContext(ctx)
	tracer.Info("s.UpdateTask")
	return nil
}

func (s *service) DeleteTask(ctx context.Context) error {
	tracer := trace.GetTraceFromContext(ctx)
	tracer.Info("s.UpdateTask")
	return nil
}
