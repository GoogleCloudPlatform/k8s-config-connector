// Copyright 2022 Google LLC
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

package preflight

import (
	"context"

	"sigs.k8s.io/kubebuilder-declarative-pattern/pkg/patterns/declarative"
)

type CompositePreflight struct {
	checkers []declarative.Preflight
}

func NewCompositePreflight(checkers []declarative.Preflight) *CompositePreflight {
	return &CompositePreflight{checkers: checkers}
}

func (c *CompositePreflight) Preflight(ctx context.Context, o declarative.DeclarativeObject) error {
	for _, checker := range c.checkers {
		if err := checker.Preflight(ctx, o); err != nil {
			return err
		}
	}
	return nil
}
