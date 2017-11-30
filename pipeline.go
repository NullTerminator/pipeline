package pipeline

import (
	"errors"
	"fmt"
	"reflect"
)

type (
	Pipeline struct {
		ctx   Context
		tasks []Task
	}
)

func NewPipeline(ctx Context) *Pipeline {
	return &Pipeline{
		ctx: ctx,
	}
}

func (p *Pipeline) Run() error {
	var err error
	for _, task := range p.tasks {
		err = task.Run(p.ctx)
		if err != nil {
			return err
		}
	}
	return nil
}

func (p *Pipeline) Add(t Task) error {
	fmt.Println(fmt.Sprintf("Adding task: %#v", t))
	err := p.checkRequirementsForTask(t)
	if err != nil {
		return errors.New(fmt.Sprintf("Missing requirements for task: %#v - %s", t, err))
	}

	p.tasks = append(p.tasks, t)
	return nil
}

func (p *Pipeline) checkRequirementsForTask(t Task) error {
	for k, v := range t.Requires().All() {
		err := p.provides(k, v)
		if err != nil {
			return err
		}
	}

	return nil
}

func (p *Pipeline) provides(key string, val interface{}) error {
	v, ok := p.ctx.Get(key)
	if ok {
		if match(v, val) {
			fmt.Println("Context provides ", key)
			return nil
		} else {
			return errors.New(fmt.Sprintf("Wrong value set for key: %s", key))
		}
	}

	fmt.Println(fmt.Sprintf("Check %d tasks for key %s", len(p.tasks), key))
	for i := len(p.tasks) - 1; i >= 0; i-- {
		for providedKey, providedVal := range p.tasks[i].Provides().All() {
			if providedKey == key {
				if match(providedVal, val) {
					return nil
				} else {
					return errors.New(fmt.Sprintf("Wrong value available for key: %s", key))
				}
			}
		}
	}

	return errors.New(fmt.Sprintf("No value available for key: %s", key))
}

func match(have interface{}, want interface{}) bool {
	have_t := reflect.TypeOf(have)
	have_k := have_t.Kind()
	if have_k == reflect.Ptr {
		have_t = have_t.Elem()
		have_k = have_t.Kind()
	}

	want_t := reflect.TypeOf(want)
	want_k := want_t.Kind()
	if want_k == reflect.Ptr {
		want_t = want_t.Elem()
		want_k = want_t.Kind()
	}

	if have_k == reflect.Struct && want_k == reflect.Interface {
		return have_t.Implements(want_t)
	}

	return have_t == want_t
}
