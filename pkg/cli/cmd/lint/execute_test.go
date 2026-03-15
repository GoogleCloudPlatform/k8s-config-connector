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

package lint

import (
	"context"
	"testing"
)

func TestShouldExclude(t *testing.T) {
	tests := []struct {
		name     string
		excludes []string
		includes []string
		want     bool
	}{
		{
			name:     "kube-system",
			excludes: []string{"kube"},
			includes: []string{},
			want:     true,
		},
		{
			name:     "my-namespace",
			excludes: []string{"kube"},
			includes: []string{},
			want:     false,
		},
		{
			name:     "my-namespace",
			excludes: []string{"kube"},
			includes: []string{"my"},
			want:     false,
		},
		{
			name:     "other-namespace",
			excludes: []string{"kube"},
			includes: []string{"my"},
			want:     true,
		},
		{
			name:     "my-prefix-namespace",
			excludes: []string{},
			includes: []string{"my-prefix"},
			want:     false,
		},
		{
			name:     "prefix-match",
			excludes: []string{"pre"},
			includes: []string{},
			want:     true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := shouldExclude(tt.name, tt.excludes, tt.includes, 0, nil); got != tt.want {
				t.Errorf("shouldExclude(%q, %v, %v) = %v, want %v", tt.name, tt.excludes, tt.includes, got, tt.want)
			}
		})
	}
}

func TestTaskQueue(t *testing.T) {
	q := &taskQueue{}
	if q.GetWork() != nil {
		t.Error("expected nil work from empty queue")
	}

	task1 := &mockTask{id: 1}
	task2 := &mockTask{id: 2}

	q.AddTask(task1)
	q.AddTask(task2)

	if got := q.GetWork().(*mockTask).id; got != 1 {
		t.Errorf("expected task 1, got %v", got)
	}
	if got := q.GetWork().(*mockTask).id; got != 2 {
		t.Errorf("expected task 2, got %v", got)
	}
	if q.GetWork() != nil {
		t.Error("expected nil work after consuming all tasks")
	}
}

type mockTask struct {
	id int
}

func (m *mockTask) Run(ctx context.Context) error {
	return nil
}

func TestResultAggregation(t *testing.T) {
	r := &Result{}
	r.addNewResource("ns1", "Kind1", "res1")
	r.addNewResource("ns1", "Kind1", "res2")
	r.addNewResource("ns1", "Kind2", "res3")
	r.addNewResource("ns2", "Kind1", "res4")

	r.lock.Lock()
	defer r.lock.Unlock()

	if len(r.resources) != 2 {
		t.Errorf("expected 2 namespaces, got %d", len(r.resources))
	}
	if len(r.resources["ns1"]) != 2 {
		t.Errorf("expected 2 kinds in ns1, got %d", len(r.resources["ns1"]))
	}
	if len(r.resources["ns1"]["Kind1"]) != 2 {
		t.Errorf("expected 2 resources in ns1/Kind1, got %d", len(r.resources["ns1"]["Kind1"]))
	}
}
