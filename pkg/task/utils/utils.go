package utils

import (
	dbtypes "reboot/pkg/dao/mysql/types"
	schedtypes "reboot/pkg/task/types"
)

// ConvertDBTaskToSchedulerTask xxx
func ConvertDBTaskToSchedulerTask(dbtask *dbtypes.Task) *schedtypes.Task {
	task := &schedtypes.Task{}
	task.ID = dbtask.ID
	task.Pause = dbtask.IsPaused
	task.Close = dbtask.IsClosed
	task.IsClosedManually = dbtask.IsClosedManually
	task.Resource = dbtask.Resource
	task.Type = dbtask.Type
	task.LastUpdateTime = dbtask.LastUpdateTime

	task.Spec = dbtask.Spec
	task.Status = dbtask.Status
	return task
}

// ConvertSchedulerTaskToDBTask xxx
func ConvertSchedulerTaskToDBTask(task *schedtypes.Task) *dbtypes.Task {
	dbtask := &dbtypes.Task{}
	dbtask.ID = task.ID
	dbtask.Resource = task.Resource
	dbtask.Type = task.Type
	dbtask.LastUpdateTime = task.LastUpdateTime
	dbtask.Spec = task.Spec
	dbtask.Status = task.Status
	return dbtask
}
