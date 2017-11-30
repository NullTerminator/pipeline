package main

import (
	"pipeline"
)

func main() {
	loadUser := NewLoadUserTask()
	sayHi := NewSayHelloToUserTask()

	ctx := pipeline.NewContext()

	ctx.Set("user_id", 42)

	pipe := pipeline.NewPipeline(ctx)
	err := pipe.Add(loadUser)
	if err != nil {
		panic(err)
	}
	err = pipe.Add(sayHi)
	if err != nil {
		panic(err)
	}

	err = pipe.Run()
	if err != nil {
		panic(err)
	}
}
