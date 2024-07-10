package cel

import (
	"github.com/google/cel-go/cel"
	"github.com/google/cel-go/common/types/ref"
	"github.com/wzshiming/easycel"
)

type Engine struct {
	resource string
	values   map[string]interface{}
	registry *easycel.Registry
	env      *easycel.Environment
}

func NewEngine(resource string, values map[string]interface{}) (*Engine, error) {
	// TODO: what withtagname ? can we pass yaml ?
	registry := easycel.NewRegistry("cel-engine", easycel.WithTagName("json"))

	// Register variables
	err := registry.RegisterVariable(resource, map[string]interface{}{})
	if err != nil {
		return nil, err
	}
	err = registry.RegisterVariable("fetched", map[string]interface{}{})
	if err != nil {
		return nil, err
	}
	err = registry.RegisterVariable("context", map[string]interface{}{})
	if err != nil {
		return nil, err
	}

	// Environment
	env, err := easycel.NewEnvironment(cel.Lib(registry))
	if err != nil {
		return nil, err
	}

	return &Engine{
		resource: resource,
		values:   values,
		registry: registry,
		env:      env,
	}, nil
}

func (e *Engine) Eval(expression string) (ref.Val, error) {
	prog, err := e.env.Program(expression)
	if err != nil {
		return nil, err
	}
	val, _, err := prog.Eval(e.values)
	if err != nil {
		return nil, err
	}

	return val, nil
}
