package types

import (
	"reboot/pkg/enum"
	"reboot/pkg/task/types"
)

// Task represent an PodVMGroup task
type Task struct {
	Common *types.Common
	//*types.Task
	Spec   *Spec
	Status *Status
}

// Spec is the task-related task's specification
type Spec struct {
	OpUser string
}

// Status is the task-related task's status
type Status struct {
	//ClusterIP string
	State    enum.State
	TryTimes int
}
