package service

import (
	"context"
)

type Options struct {
	//DB
}

type Operation interface {
	GetTask(ctx context.Context) error
	ListTask(ctx context.Context) error
	DeleteTask(ctx context.Context) error
	CreateTask(ctx context.Context) error
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
