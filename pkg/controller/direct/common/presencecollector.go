// Copyright 2026 Google LLC
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

package common

import (
	"strings"

	"k8s.io/apimachinery/pkg/util/sets"
)

// PresenceCollector collects all populated field paths during a single DFS traversal of an object.
type PresenceCollector struct {
	Populated sets.Set[string]
}

// NewPresenceCollector creates a new PresenceCollector instance.
func NewPresenceCollector() *PresenceCollector {
	return &PresenceCollector{
		Populated: sets.New[string](),
	}
}

// VisitField implements the Visitor interface, recording all visited paths into the Populated set.
func (c *PresenceCollector) VisitField(path string, value any) error {
	cleanPath := strings.TrimPrefix(path, ".")
	if cleanPath != "" {
		c.Populated.Insert(cleanPath)
	}
	return nil
}

// CollectPresentFields traverses obj using VisitFields and returns a set of all non-nil dot-separated field paths.
func CollectPresentFields(obj any) sets.Set[string] {
	collector := NewPresenceCollector()
	if obj != nil {
		_ = VisitFields(obj, collector)
	}
	return collector.Populated
}
