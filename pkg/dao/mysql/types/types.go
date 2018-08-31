package types

import (
	"time"
)

type Task struct {
	ID               int64     `db:"id"`
	NameSpace        string    `db:"namespace"`
	Resource         string    `db:"resource"`
	Type             string    `db:"task_type"`
	Spec             string    `db:"spec"`
	Status           string    `db:"status"`
	IsCanceled       bool      `db:"is_canceled"`
	IsSkipPaused     bool      `db:"is_skip_paused"`
	IsUrgentSkipped  bool      `db:"is_urgent_skipped"`
	IsClosed         bool      `db:"is_closed"`
	IsPaused         bool      `db:"is_paused"`
	IsClosedManually bool      `db:"is_closed_manually"`
	OpUser           string    `db:"op_user"`
	CreateTime       time.Time `db:"create_time"`
	LastUpdateTime   time.Time `db:"last_update_time"`
}

type Field string

const (
	FieldID        = Field("id")
	FieldNameSpace = Field("namespace")
	FieldResource  = Field("resource")
	FieldType      = Field("task_type")
	FieldStatus    = Field("status")
	FieldIsPaused  = Field("is_paused")
	FieldIsClosed  = Field("is_closed")
)

type Value interface{}
