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

package mustache

import (
	"context"
	"fmt"

	"github.com/google/cel-go/cel"
	"k8s.io/klog/v2"
)

type MustacheExpression struct {
	Expression string
}

func (l *MustacheExpression) Eval(ctx context.Context, evaluationActivation *Activation) (string, error) {
	// TODO: Pre-parse

	var envOptions []cel.EnvOption
	envOptions = append(envOptions, cel.Variable("object", cel.DynType))
	for k := range evaluationActivation.Objects {
		envOptions = append(envOptions, cel.Variable(k, cel.DynType))
	}

	env, err := cel.NewEnv(envOptions...)
	if err != nil {
		return "", fmt.Errorf("building CEL env: %w", err)
	}

	ast, issues := env.Compile(l.Expression)
	if issues != nil && issues.Err() != nil {
		err := issues.Err()
		return "", fmt.Errorf("compiling CEL expression %q: %w", l.Expression, err)
	}
	prg, err := env.Program(ast)
	if err != nil {
		return "", fmt.Errorf("building CEL program for %q: %w", l.Expression, err)
	}

	scope := map[string]any{}
	scope["object"] = evaluationActivation.Object
	for k, v := range evaluationActivation.Objects {
		scope[k] = v
	}
	out, details, err := prg.Eval(scope)
	if err != nil {
		return "", fmt.Errorf("evaluating CEL expression %q: %w", l.Expression, err)
	}
	klog.Infof("out is %v", out)

	klog.Infof("details is %v", details)

	return fmt.Sprintf("%v", out), nil
}
