package taskscheduler

import (
	"encoding/json"
	dbtypes "reboot/pkg/dao/mysql/types"
	tasktypes "reboot/pkg/task/taskscheduler/types"
	schedtypes "reboot/pkg/task/types"
)

func convertSchedTaskToTask(schedtask *schedtypes.Task) (*tasktypes.Task, error) {
	t := &tasktypes.Task{
		Common: &schedtask.Common,
	}
	// t.Spec = &tasktypes.Spec{}
	// err := json.Unmarshal([]byte(schedtask.Spec), t.Spec)
	// if err != nil {
	// 	return nil, err
	// }
	t.Status = &tasktypes.Status{}
	if len(schedtask.Status) == 0 {
		schedtask.Status = "{}"
	}
	err := json.Unmarshal([]byte(schedtask.Status), t.Status)
	if err != nil {
		return nil, err
	}

	return t, nil
}

func convertTaskToDBTask(grouptask *tasktypes.Task) *dbtypes.Task {
	dbtask := &dbtypes.Task{}

	dbtask.ID = grouptask.Common.ID
	dbtask.Resource = string(grouptask.Common.Resource)
	dbtask.Type = string(grouptask.Common.Type)
	dbtask.IsClosed = grouptask.Common.Close
	dbtask.IsPaused = grouptask.Common.Pause

	// data, _ := json.Marshal(grouptask.Spec)
	// dbtask.Spec = string(data)
	data, _ := json.Marshal(grouptask.Status)
	dbtask.Status = string(data)
	return dbtask
}
