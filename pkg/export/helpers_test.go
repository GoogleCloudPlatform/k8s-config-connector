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

package export_test

import (
	"testing"

	"k8s.io/apimachinery/pkg/apis/meta/v1/unstructured"

	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/export"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/k8s"
)

func TestSetProjectID(t *testing.T) {
	u := &unstructured.Unstructured{}
	export.SetProjectID(u, "my-project-123")

	ann := u.GetAnnotations()
	if ann == nil {
		t.Fatalf("expected annotations to be set")
	}

	val := ann[k8s.ProjectIDAnnotation]
	if val != "my-project-123" {
		t.Errorf("expected annotation %q to be %q, got %q", k8s.ProjectIDAnnotation, "my-project-123", val)
	}
}

func TestSetLabels(t *testing.T) {
	u := &unstructured.Unstructured{}
	labels := map[string]string{
		"foo": "bar",
		"baz": "qux",
	}
	export.SetLabels(u, labels)

	gotLabels := u.GetLabels()
	if len(gotLabels) != 2 {
		t.Fatalf("expected 2 labels, got %d", len(gotLabels))
	}

	if gotLabels["foo"] != "bar" || gotLabels["baz"] != "qux" {
		t.Errorf("expected labels map to match %+v, got %+v", labels, gotLabels)
	}
}
