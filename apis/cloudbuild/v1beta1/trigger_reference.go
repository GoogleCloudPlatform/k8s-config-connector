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

	"github.com/GoogleCloudPlatform/k8s-config-connector/apis/common/identity"
	refs "github.com/GoogleCloudPlatform/k8s-config-connector/apis/refs/v1beta1"
	"k8s.io/apimachinery/pkg/apis/meta/v1/unstructured"
	"k8s.io/apimachinery/pkg/runtime/schema"
	"k8s.io/apimachinery/pkg/types"
	"sigs.k8s.io/controller-runtime/pkg/client"
)

var _ refs.Ref = &CloudBuildTriggerRef{}

// CloudBuildTriggerRef is a reference to a CloudBuildTrigger resource.
type CloudBuildTriggerRef struct {
	// A reference to an externally managed CloudBuildTrigger resource.
	// Should be in the format "projects/{project}/locations/{location}/triggers/{trigger}".
	External string `json:"external,omitempty"`

	// The name of a CloudBuildTrigger resource.
	Name string `json:"name,omitempty"`

	// The namespace of a CloudBuildTrigger resource.
	Namespace string `json:"namespace,omitempty"`
}

func init() {
	refs.Register(&CloudBuildTriggerRef{})
}

func (r *CloudBuildTriggerRef) GetGVK() schema.GroupVersionKind {
	return CloudBuildTriggerGVK
}

func (r *CloudBuildTriggerRef) GetNamespacedName() types.NamespacedName {
	return types.NamespacedName{
		Name:      r.Name,
		Namespace: r.Namespace,
	}
}

func (r *CloudBuildTriggerRef) GetExternal() string {
	return r.External
}

func (r *CloudBuildTriggerRef) SetExternal(ref string) {
	r.External = ref
}

func (r *CloudBuildTriggerRef) ValidateExternal(ref string) error {
	id := &CloudBuildTriggerIdentity{}
	if err := id.FromExternal(ref); err != nil {
		return err
	}
	return nil
}

func (r *CloudBuildTriggerRef) ParseExternalToIdentity() (identity.Identity, error) {
	id := &CloudBuildTriggerIdentity{}
	if err := id.FromExternal(r.External); err != nil {
		return nil, err
	}
	return id, nil
}

func (r *CloudBuildTriggerRef) Normalize(ctx context.Context, reader client.Reader, defaultNamespace string) error {
	fallback := func(u *unstructured.Unstructured) string {
		// Trigger ID is server-generated, so we need it from status.
		triggerID, _, _ := unstructured.NestedString(u.Object, "status", "triggerId")
		if triggerID == "" {
			return ""
		}

		location, _, _ := unstructured.NestedString(u.Object, "spec", "location")
		if location == "" {
			return ""
		}

		projectID, err := refs.ResolveProjectID(ctx, reader, u)
		if err != nil {
			return ""
		}

		return fmt.Sprintf("projects/%s/locations/%s/triggers/%s", projectID, location, triggerID)
	}
	return refs.NormalizeWithFallback(ctx, reader, r, defaultNamespace, fallback)
}
