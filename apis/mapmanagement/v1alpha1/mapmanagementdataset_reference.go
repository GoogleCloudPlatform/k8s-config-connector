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

package v1alpha1

import (
	"context"
	"fmt"
	"strings"

	refsv1beta1 "github.com/GoogleCloudPlatform/k8s-config-connector/apis/refs/v1beta1"
	"k8s.io/apimachinery/pkg/runtime/schema"
	"k8s.io/apimachinery/pkg/types"
	"sigs.k8s.io/controller-runtime/pkg/client"
)

var _ refsv1beta1.Ref = &MapManagementDatasetRef{}

var MapManagementDatasetGVK = GroupVersion.WithKind("MapManagementDataset")

// MapManagementDatasetRef is a reference to a GCP MapManagementDataset resource.
type MapManagementDatasetRef struct {
	// A reference to an externally managed MapManagementDataset resource.
	// Should be in the format "projects/{{projectID}}/datasets/{{datasetID}}".
	External string `json:"external,omitempty"`
}

func (r *MapManagementDatasetRef) GetGVK() schema.GroupVersionKind {
	return MapManagementDatasetGVK
}

func (r *MapManagementDatasetRef) GetNamespacedName() types.NamespacedName {
	return types.NamespacedName{}
}

func (r *MapManagementDatasetRef) GetExternal() string {
	return r.External
}

func (r *MapManagementDatasetRef) SetExternal(ref string) {
	r.External = ref
}

func (r *MapManagementDatasetRef) ValidateExternal(ref string) error {
	if _, err := ParseMapManagementDatasetExternal(ref); err != nil {
		return err
	}
	return nil
}

func (r *MapManagementDatasetRef) Normalize(ctx context.Context, reader client.Reader, defaultNamespace string) error {
	if r.External == "" {
		return fmt.Errorf("external reference must be specified for %s", MapManagementDatasetGVK.Kind)
	}
	return r.ValidateExternal(r.External)
}

func ParseMapManagementDatasetExternal(external string) (string, error) {
	tokens := strings.Split(external, "/")
	if len(tokens) != 4 || tokens[0] != "projects" || tokens[2] != "datasets" || tokens[1] == "" || tokens[3] == "" {
		return "", fmt.Errorf("format of MapManagementDataset external=%q was not known (use projects/{{projectID}}/datasets/{{datasetID}})", external)
	}
	return external, nil
}
