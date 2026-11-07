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

package krmtotf

import (
	"testing"

	corekccv1alpha1 "github.com/GoogleCloudPlatform/k8s-config-connector/pkg/apis/core/v1alpha1"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/k8s"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime/schema"
)

func TestLoggingLogSinkParentContainerAnnotationFallback(t *testing.T) {
	rc := corekccv1alpha1.ResourceConfig{
		Name:       "google_logging_log_sink",
		Kind:       "LoggingLogSink",
		IDTemplate: "{{project?}}#{{folder?}}#{{org_id?}}#{{name}}",
		MetadataMapping: corekccv1alpha1.MetadataMapping{
			Name: "name",
		},
		ResourceID: corekccv1alpha1.ResourceID{
			TargetField: "name",
		},
		ResourceReferences: []corekccv1alpha1.ReferenceConfig{
			{
				TFField: "project",
				TypeConfig: corekccv1alpha1.TypeConfig{
					Key:    "projectRef",
					Parent: true,
					GVK: schema.GroupVersionKind{
						Group:   "resourcemanager.cnrm.cloud.google.com",
						Version: "v1beta1",
						Kind:    "Project",
					},
				},
			},
			{
				TFField: "folder",
				TypeConfig: corekccv1alpha1.TypeConfig{
					Key:    "folderRef",
					Parent: true,
					GVK: schema.GroupVersionKind{
						Group:   "resourcemanager.cnrm.cloud.google.com",
						Version: "v1beta1",
						Kind:    "Folder",
					},
				},
			},
			{
				TFField: "org_id",
				TypeConfig: corekccv1alpha1.TypeConfig{
					Key:    "organizationRef",
					Parent: true,
					GVK: schema.GroupVersionKind{
						Group:   "resourcemanager.cnrm.cloud.google.com",
						Version: "v1beta1",
						Kind:    "Organization",
					},
				},
			},
		},
	}

	t.Run("resolves project from container annotation when parent refs omitted", func(t *testing.T) {
		r := &Resource{
			Resource: k8s.Resource{
				ObjectMeta: metav1.ObjectMeta{
					Name: "my-sink",
					Annotations: map[string]string{
						k8s.ProjectIDAnnotation: "my-annotated-project",
					},
				},
				Spec: map[string]interface{}{},
			},
			ResourceConfig: rc,
		}

		gotID, err := expandTemplate(rc.IDTemplate, r, nil, nil)
		if err != nil {
			t.Fatalf("expandTemplate unexpected error: %v", err)
		}
		wantID := "my-annotated-project###my-sink"
		if gotID != wantID {
			t.Errorf("expandTemplate() = %q, want %q", gotID, wantID)
		}

		cfg := map[string]interface{}{}
		if err := resolveContainerValue(cfg, r, nil, nil); err != nil {
			t.Fatalf("resolveContainerValue unexpected error: %v", err)
		}
		if cfg["project"] != "my-annotated-project" {
			t.Errorf("cfg[\"project\"] = %v, want %q", cfg["project"], "my-annotated-project")
		}
	})

	t.Run("does not use project annotation when folderRef is explicitly specified in spec", func(t *testing.T) {
		r := &Resource{
			Resource: k8s.Resource{
				ObjectMeta: metav1.ObjectMeta{
					Name: "my-folder-sink",
					Annotations: map[string]string{
						k8s.ProjectIDAnnotation: "my-annotated-project",
					},
				},
				Spec: map[string]interface{}{
					"folderRef": map[string]interface{}{
						"external": "123456789",
					},
				},
			},
			ResourceConfig: rc,
		}

		gotID, err := expandTemplate(rc.IDTemplate, r, nil, nil)
		if err != nil {
			t.Fatalf("expandTemplate unexpected error: %v", err)
		}
		wantID := "#123456789##my-folder-sink"
		if gotID != wantID {
			t.Errorf("expandTemplate() = %q, want %q", gotID, wantID)
		}
	})
}
