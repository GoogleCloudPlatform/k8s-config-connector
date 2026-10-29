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

package v1beta1

import (
	"context"
	"fmt"
	"strings"

	"k8s.io/apimachinery/pkg/apis/meta/v1/unstructured"
	"k8s.io/apimachinery/pkg/runtime/schema"
	"k8s.io/apimachinery/pkg/types"
	"sigs.k8s.io/controller-runtime/pkg/client"
)

var (
	BigQueryDatasetGVK = schema.GroupVersionKind{
		Group:   "bigquery.cnrm.cloud.google.com",
		Version: "v1beta1",
		Kind:    "BigQueryDataset",
	}

	_ Ref = &BigQueryDatasetRef{}
)

// BigQueryDatasetRef is a reference to a BigQueryDataset resource.
type BigQueryDatasetRef struct {
	// Name of the referenced object.
	// +optional
	Name string `json:"name,omitempty"`

	// Namespace of the referenced object.
	// +optional
	Namespace string `json:"namespace,omitempty"`

	// The BigQueryDataset selfLink, when not managed by Config Connector.
	// +optional
	External string `json:"external,omitempty"`
}

func (r *BigQueryDatasetRef) GetGVK() schema.GroupVersionKind {
	return BigQueryDatasetGVK
}

func (r *BigQueryDatasetRef) GetNamespacedName() types.NamespacedName {
	return types.NamespacedName{
		Name:      r.Name,
		Namespace: r.Namespace,
	}
}

func (r *BigQueryDatasetRef) GetExternal() string {
	return r.External
}

func (r *BigQueryDatasetRef) SetExternal(ref string) {
	r.External = ref
	r.Name = ""
	r.Namespace = ""
}

func (r *BigQueryDatasetRef) ValidateExternal(ref string) error {
	tokens := strings.Split(strings.TrimPrefix(ref, "/"), "/")
	if len(tokens) == 4 && tokens[0] == "projects" && tokens[2] == "datasets" && tokens[1] != "" && tokens[3] != "" {
		return nil
	}
	return fmt.Errorf("format of BigQueryDataset external=%q was not known (use projects/<projectId>/datasets/<datasetId>)", ref)
}

func (r *BigQueryDatasetRef) Normalize(ctx context.Context, reader client.Reader, defaultNamespace string) error {
	fallback := func(u *unstructured.Unstructured) string {
		projectID, err := ResolveProjectID(ctx, reader, u)
		if err != nil || projectID == "" {
			return ""
		}
		resourceID, err := GetResourceID(u)
		if err != nil || resourceID == "" {
			return ""
		}
		return fmt.Sprintf("projects/%s/datasets/%s", projectID, resourceID)
	}
	return NormalizeWithFallback(ctx, reader, r, defaultNamespace, fallback)
}
