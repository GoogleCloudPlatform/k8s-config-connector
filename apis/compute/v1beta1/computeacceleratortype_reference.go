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
	ComputeAcceleratorTypeGVK = GroupVersion.WithKind("ComputeAcceleratorType")
)
var _ refsv1beta1.Ref = &ComputeAcceleratorTypeRef{}

// ComputeAcceleratorTypeRef is a reference to a ComputeAcceleratorType resource.
// Note: ComputeAcceleratorType is not yet managed by Config Connector.
type ComputeAcceleratorTypeRef struct {
	// A reference to an externally managed ComputeAcceleratorType resource.
	// Should be in the format "projects/{{project}}/zones/{{zone}}/acceleratorTypes/{{name}}".
	External string `json:"external,omitempty"`

	// The name of a ComputeAcceleratorType resource.
	Name string `json:"name,omitempty"`

	// The namespace of a ComputeAcceleratorType resource.
	Namespace string `json:"namespace,omitempty"`
}

func (r *ComputeAcceleratorTypeRef) GetGVK() schema.GroupVersionKind {
	return ComputeAcceleratorTypeGVK
}

func (r *ComputeAcceleratorTypeRef) GetNamespacedName() types.NamespacedName {
	return types.NamespacedName{
		Name:      r.Name,
		Namespace: r.Namespace,
	}
}

func (r *ComputeAcceleratorTypeRef) GetExternal() string {
	return r.External
}

func (r *ComputeAcceleratorTypeRef) SetExternal(ref string) {
	r.External = ref
}

func (r *ComputeAcceleratorTypeRef) ValidateExternal(ref string) error {
	id := &ComputeAcceleratorTypeIdentity{}
	if err := id.FromExternal(r.GetExternal()); err != nil {
		return err
	}
	return nil
}

func (r *ComputeAcceleratorTypeRef) Normalize(ctx context.Context, reader client.Reader, defaultNamespace string) error {
	if r.GetExternal() != "" {
		return nil
	}
	return fmt.Errorf("ComputeAcceleratorType is not managed by KCC, so only the `external` field is supported")
}
