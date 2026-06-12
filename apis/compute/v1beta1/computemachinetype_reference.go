// Copyright 2026 Google LLC
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not not use this file except in compliance with the License.
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

	refsv1beta1 "github.com/GoogleCloudPlatform/k8s-config-connector/apis/refs/v1beta1"
	"k8s.io/apimachinery/pkg/runtime/schema"
	"k8s.io/apimachinery/pkg/types"
	"sigs.k8s.io/controller-runtime/pkg/client"
)

var (
	ComputeMachineTypeGVK = GroupVersion.WithKind("ComputeMachineType")
)
var _ refsv1beta1.Ref = &ComputeMachineTypeRef{}

// ComputeMachineTypeRef is a reference to a ComputeMachineType resource.
// Note: ComputeMachineType is not yet managed by Config Connector.
type ComputeMachineTypeRef struct {
	// A reference to an externally managed ComputeMachineType resource.
	// Should be in the format "projects/{{project}}/zones/{{zone}}/machineTypes/{{name}}".
	External string `json:"external,omitempty"`

	// The name of a ComputeMachineType resource.
	Name string `json:"name,omitempty"`

	// The namespace of a ComputeMachineType resource.
	Namespace string `json:"namespace,omitempty"`
}

func (r *ComputeMachineTypeRef) GetGVK() schema.GroupVersionKind {
	return ComputeMachineTypeGVK
}

func (r *ComputeMachineTypeRef) GetNamespacedName() types.NamespacedName {
	return types.NamespacedName{
		Name:      r.Name,
		Namespace: r.Namespace,
	}
}

func (r *ComputeMachineTypeRef) GetExternal() string {
	return r.External
}

func (r *ComputeMachineTypeRef) SetExternal(ref string) {
	r.External = ref
}

func (r *ComputeMachineTypeRef) ValidateExternal(ref string) error {
	id := &ComputeMachineTypeIdentity{}
	if err := id.FromExternal(r.GetExternal()); err != nil {
		return err
	}
	return nil
}

func (r *ComputeMachineTypeRef) Normalize(ctx context.Context, reader client.Reader, defaultNamespace string) error {
	if r.GetExternal() != "" {
		return nil
	}
	return fmt.Errorf("ComputeMachineType is not managed by KCC, so only the `external` field is supported")
}
