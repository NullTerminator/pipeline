package main

import (
	"fmt"
	"pipeline"
)

const USER_ID_K = "user_id"

type (
	LoadUserTask struct {
	}
)

func NewLoadUserTask() *LoadUserTask {
	return &LoadUserTask{}
}

func (t *LoadUserTask) Requires() pipeline.TaskRequirements {
	requirements := pipeline.NewTaskRequirements()
	requirements.Add(USER_ID_K, 0)
	return requirements
}

func (t *LoadUserTask) Provides() pipeline.TaskRequirements {
	requirements := pipeline.NewTaskRequirements()
	requirements.Add("user", User{})
	return requirements
}

func (t *LoadUserTask) Run(ctx pipeline.Context) error {
	maybeUserId, _ := ctx.Get(USER_ID_K)
	user_id := maybeUserId.(int)

	fmt.Println("Loading user with id: ", user_id)
	user := User{
		Name: "Bob",
	}
	return ctx.Set("user", user)
}
