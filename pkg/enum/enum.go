package enum

type State string

func (s State) String() string {
	return string(s)
}

const (
	TaskPending = State("task-pending")
	TaskDoing   = State("task-doing")
	TaskDone    = State("task-done")
)
