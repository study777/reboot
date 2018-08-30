package service

import (
	"context"
	"reboot/pkg/dao"
)

type Options struct {
	DB dao.Storage
}

type Operation interface {
	GetTask(ctx context.Context) error
	ListTask(ctx context.Context) error
	DeleteTask(ctx context.Context) error
	//CreateTask(ctx context.Context) error
	CreateTask(ctx context.Context, namespace string, resource string) error
	UpdateTask(ctx context.Context) error
}

type service struct {
	opt *Options
}

//定义一个工厂函数
func New(opt *Options) Operation {
	return &service{
		opt: opt,
	}
}

type Task struct {
	Resource string
}
