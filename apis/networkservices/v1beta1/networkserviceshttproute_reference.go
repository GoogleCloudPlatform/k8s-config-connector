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

	"github.com/GoogleCloudPlatform/k8s-config-connector/apis/common"
	"github.com/GoogleCloudPlatform/k8s-config-connector/apis/common/identity"
	refsv1beta1 "github.com/GoogleCloudPlatform/k8s-config-connector/apis/refs/v1beta1"
	"k8s.io/apimachinery/pkg/apis/meta/v1/unstructured"
	"k8s.io/apimachinery/pkg/runtime/schema"
	"k8s.io/apimachinery/pkg/types"
	"sigs.k8s.io/controller-runtime/pkg/client"
)

var _ refsv1beta1.Ref = &NetworkServicesHTTPRouteRef{}

// NetworkServicesHTTPRouteRef is a reference to a GCP NetworkServicesHTTPRoute.
type NetworkServicesHTTPRouteRef struct {
	// A reference to an externally managed NetworkServicesHTTPRoute resource.
	// Should be in the format "projects/{{projectID}}/locations/{{location}}/httpRoutes/{{httpRouteID}}".
	External string `json:"external,omitempty"`

	// The name of a NetworkServicesHTTPRoute resource.
	Name string `json:"name,omitempty"`

	// The namespace of a NetworkServicesHTTPRoute resource.
	Namespace string `json:"namespace,omitempty"`
}

func init() {
	refsv1beta1.Register(&NetworkServicesHTTPRouteRef{}, &NetworkServicesHTTPRoute{})
}

func (r *NetworkServicesHTTPRouteRef) GetGVK() schema.GroupVersionKind {
	return NetworkServicesHTTPRouteGVK
}

func (r *NetworkServicesHTTPRouteRef) GetNamespacedName() types.NamespacedName {
	return types.NamespacedName{
		Name:      r.Name,
		Namespace: r.Namespace,
	}
}

func (r *NetworkServicesHTTPRouteRef) GetExternal() string {
	return r.External
}

func (r *NetworkServicesHTTPRouteRef) SetExternal(ref string) {
	r.External = ref
	r.Name = ""
	r.Namespace = ""
}

func (r *NetworkServicesHTTPRouteRef) ValidateExternal(ref string) error {
	id := &NetworkServicesHTTPRouteIdentity{}
	if err := id.FromExternal(ref); err != nil {
		return err
	}
	return nil
}

func (r *NetworkServicesHTTPRouteRef) ParseExternalToIdentity() (identity.Identity, error) {
	id := &NetworkServicesHTTPRouteIdentity{}
	if err := id.FromExternal(r.External); err != nil {
		return nil, err
	}
	return id, nil
}

func (r *NetworkServicesHTTPRouteRef) Normalize(ctx context.Context, reader client.Reader, defaultNamespace string) error {
	if r.External != "" && r.Name != "" {
		return fmt.Errorf("cannot specify both name and external on %s reference", r.GetGVK().Kind)
	}
	fallback := func(u *unstructured.Unstructured) string {
		// DCL / Legacy fallback checks if referenced resource is ready in GCP
		ready := false
		conditions, found, err := unstructured.NestedSlice(u.Object, "status", "conditions")
		if err == nil && found {
			for _, condVal := range conditions {
				cond, ok := condVal.(map[string]interface{})
				if !ok {
					continue
				}
				if cond["type"] == "Ready" && cond["status"] == "True" {
					ready = true
					break
				}
			}
		}
		if !ready {
			return ""
		}

		obj, err := common.ToStructuredType[*NetworkServicesHTTPRoute](u)
		if err != nil {
			return ""
		}
		identity, err := NewNetworkServicesHTTPRouteIdentity(ctx, reader, obj)
		if err != nil {
			return ""
		}
		return identity.String()
	}
	return refsv1beta1.NormalizeWithFallback(ctx, reader, r, defaultNamespace, fallback)
}

var _ refsv1beta1.Ref = &BackendServiceRef{}

func (r *BackendServiceRef) GetGVK() schema.GroupVersionKind {
	return schema.GroupVersionKind{
		Group:   "compute.cnrm.cloud.google.com",
		Version: "v1beta1",
		Kind:    "ComputeBackendService",
	}
}

func (r *BackendServiceRef) GetNamespacedName() types.NamespacedName {
	return types.NamespacedName{
		Name:      r.Name,
		Namespace: r.Namespace,
	}
}

func (r *BackendServiceRef) GetExternal() string {
	return r.External
}

func (r *BackendServiceRef) SetExternal(ref string) {
	r.External = ref
	r.Name = ""
	r.Namespace = ""
}

func (r *BackendServiceRef) ValidateExternal(ref string) error {
	// Accept any external format for backward compatibility, similar to ComputeBackendServiceRef
	return nil
}

func (r *BackendServiceRef) Normalize(ctx context.Context, reader client.Reader, defaultNamespace string) error {
	if r.External != "" && r.Name != "" {
		return fmt.Errorf("cannot specify both name and external on %s reference", r.GetGVK().Kind)
	}
	fallback := func(u *unstructured.Unstructured) string {
		selfLink, _, _ := unstructured.NestedString(u.Object, "status", "selfLink")
		return selfLink
	}
	return refsv1beta1.NormalizeWithFallback(ctx, reader, r, defaultNamespace, fallback)
}
