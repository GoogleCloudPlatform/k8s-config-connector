// Copyright 2025 Google LLC
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

	"k8s.io/apimachinery/pkg/runtime/schema"
	"k8s.io/apimachinery/pkg/types"
	"sigs.k8s.io/controller-runtime/pkg/client"
)

var _ Ref = &MemorystoreInstanceRef{}

// MemorystoreInstanceRef defines the resource reference to MemorystoreInstance, which "External" field
// holds the GCP identifier for the KRM object.
type MemorystoreInstanceRef struct {
	// A reference to an externally managed MemorystoreInstance resource.
	// Should be in the format "projects/{{projectID}}/locations/{{location}}/instances/{{instanceID}}".
	External string `json:"external,omitempty"`

	// The name of a MemorystoreInstance resource.
	Name string `json:"name,omitempty"`

	// The namespace of a MemorystoreInstance resource.
	Namespace string `json:"namespace,omitempty"`
}

// GetGVK returns the schema.GroupVersionKind of the reference type
func (r *MemorystoreInstanceRef) GetGVK() schema.GroupVersionKind {
	return schema.GroupVersionKind{
		Group:   "memorystore.cnrm.cloud.google.com",
		Version: "v1beta1",
		Kind:    "MemorystoreInstance",
	}
}

// GetNamespacedName returns the types.NamespacedName of a given reference
func (r *MemorystoreInstanceRef) GetNamespacedName() types.NamespacedName {
	return types.NamespacedName{Name: r.Name, Namespace: r.Namespace}
}

// GetExternal returns the external reference string (if set) of the reference.
func (r *MemorystoreInstanceRef) GetExternal() string {
	return r.External
}

// SetExternal sets the external reference string for a reference.
func (r *MemorystoreInstanceRef) SetExternal(ref string) {
	r.External = ref
}

// ValidateExternal returns nil if the given external reference string has a valid format for the reference.
// Otherwise, it returns an error.
// Format: projects/{{project_id}}/locations/{{location}}/instances/{{instance_id}}
func (r *MemorystoreInstanceRef) ValidateExternal(ref string) error {
	parts := strings.Split(ref, "/")
	if len(parts) != 6 {
		return fmt.Errorf("invalid external reference format: %s", ref)
	}
	if parts[0] != "projects" || parts[2] != "locations" || parts[4] != "instances" {
		return fmt.Errorf("invalid external reference format: %s", ref)
	}
	return nil
}

// Normalize ensures the "External" reference (in string format) is
// set for a given Ref, and that it has the correct format.
//
// If "External" is already set, the format will be validated.
//
// If "External" is not set, the NamespacedName will be used to query the
// referenced object from the K8s API and fetch it's external reference
// value. If "Namespace" is not specified in the reference, the
// `defaultNamespace“ will be used instead.
func (r *MemorystoreInstanceRef) Normalize(ctx context.Context, reader client.Reader, defaultNamespace string) error {
	return Normalize(ctx, reader, r, defaultNamespace)
}
