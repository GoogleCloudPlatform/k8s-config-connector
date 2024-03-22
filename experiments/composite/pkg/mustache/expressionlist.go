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
	"strings"
)

type Expression interface {
	Eval(ctx context.Context, evaluationActivation *Activation) (string, error)
}

type ExpressionList struct {
	Expressions []Expression
}

func (l *ExpressionList) Eval(ctx context.Context, activation *Activation) (string, error) {
	var values []string
	for _, e := range l.Expressions {
		v, err := e.Eval(ctx, activation)
		if err != nil {
			return "", err
		}
		values = append(values, v)
	}
	return strings.Join(values, ""), nil
}
