package pipeline

type (
	resource struct {
		Key   string
		Value interface{}
	}

	TaskRequirements interface {
		Add(string, interface{})
		All() map[string]interface{}
	}

	taskRequirements struct {
		requirements map[string]interface{}
	}

	Task interface {
		Run(ctx Context) error
		Requires() TaskRequirements
		Provides() TaskRequirements
	}
)

func NewTaskRequirements() TaskRequirements {
	return &taskRequirements{
		requirements: make(map[string]interface{}),
	}
}

func (tr *taskRequirements) Add(key string, val interface{}) {
	tr.requirements[key] = val
}

func (tr *taskRequirements) All() map[string]interface{} {
	return tr.requirements
}
