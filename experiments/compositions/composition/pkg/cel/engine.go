// Copyright 2024 Google LLC
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//      http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package cel

import (
	"github.com/google/cel-go/cel"
	"github.com/google/cel-go/common/types/ref"
	"github.com/wzshiming/easycel"
	"k8s.io/apimachinery/pkg/apis/meta/v1/unstructured"
)

type Engine struct {
	u        *unstructured.Unstructured
	registry *easycel.Registry
	env      *easycel.Environment
}

func NewEngine(u *unstructured.Unstructured) (*Engine, error) {
	// TODO: what withtagname ? can we pass yaml ?
	registry := easycel.NewRegistry("cel-engine", easycel.WithTagName("json"))

	// Register variables
	err := registry.RegisterVariable("spec", map[string]interface{}{})
	if err != nil {
		return nil, err
	}
	err = registry.RegisterVariable("status", map[string]interface{}{})
	if err != nil {
		return nil, err
	}
	err = registry.RegisterVariable("metadata", map[string]interface{}{})
	if err != nil {
		return nil, err
	}

	// Environment
	env, err := easycel.NewEnvironment(cel.Lib(registry))
	if err != nil {
		return nil, err
	}

	return &Engine{
		u:        u,
		registry: registry,
		env:      env,
	}, nil
}

func (e *Engine) Eval(expression string) (ref.Val, error) {
	prog, err := e.env.Program(expression)
	if err != nil {
		return nil, err
	}
	val, _, err := prog.Eval(e.u.Object)
	if err != nil {
		return nil, err
	}

	return val, nil
}
