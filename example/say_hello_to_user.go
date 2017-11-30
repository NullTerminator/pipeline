package main

import (
	"fmt"
	"pipeline"
)

type (
	SayHelloToUserTask struct {
	}

	HavingName interface {
		GetName() string
	}
)

var _ HavingName = (*User)(nil)

func NewSayHelloToUserTask() *SayHelloToUserTask {
	return &SayHelloToUserTask{}
}

func (t *SayHelloToUserTask) Requires() pipeline.TaskRequirements {
	requirements := pipeline.NewTaskRequirements()
	var havingName *HavingName
	requirements.Add("user", havingName)
	return requirements
}

func (t *SayHelloToUserTask) Provides() pipeline.TaskRequirements {
	return pipeline.NewTaskRequirements()
}

func (t *SayHelloToUserTask) Run(ctx pipeline.Context) error {
	maybeName, _ := ctx.Get("user")
	withName := maybeName.(HavingName)

	fmt.Println(fmt.Sprintf("Hello %s!!!", withName.GetName()))

	return nil
}
