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

	"github.com/GoogleCloudPlatform/k8s-config-connector/apis/common/identity"
	refs "github.com/GoogleCloudPlatform/k8s-config-connector/apis/refs/v1beta1"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/gcpurls"
	"k8s.io/apimachinery/pkg/apis/meta/v1/unstructured"
	"k8s.io/apimachinery/pkg/runtime/schema"
	"k8s.io/apimachinery/pkg/types"
	"sigs.k8s.io/controller-runtime/pkg/client"
)

var _ refs.Ref = &ComputeHTTPHealthCheckRef{}

var ComputeHTTPHealthCheckGVK = schema.GroupVersionKind{
	Group:   "compute.cnrm.cloud.google.com",
	Version: "v1beta1",
	Kind:    "ComputeHTTPHealthCheck",
}

// ComputeHTTPHealthCheckRef is a reference to a ComputeHTTPHealthCheck.
type ComputeHTTPHealthCheckRef struct {
	// A reference to an externally managed ComputeHTTPHealthCheck resource.
	// Should be in the format "projects/{{projectID}}/global/httpHealthChecks/{{httpHealthCheck}}".
	External string `json:"external,omitempty"`

	// The name of a ComputeHTTPHealthCheck resource.
	Name string `json:"name,omitempty"`

	// The namespace of a ComputeHTTPHealthCheck resource.
	Namespace string `json:"namespace,omitempty"`
}

func init() {
	refs.Register(&ComputeHTTPHealthCheckRef{})
}

func (r *ComputeHTTPHealthCheckRef) GetGVK() schema.GroupVersionKind {
	return ComputeHTTPHealthCheckGVK
}

func (r *ComputeHTTPHealthCheckRef) GetNamespacedName() types.NamespacedName {
	return types.NamespacedName{
		Name:      r.Name,
		Namespace: r.Namespace,
	}
}

func (r *ComputeHTTPHealthCheckRef) GetExternal() string {
	return r.External
}

func (r *ComputeHTTPHealthCheckRef) SetExternal(ref string) {
	r.External = ref
	r.Name = ""
	r.Namespace = ""
}

var ComputeHTTPHealthCheckIdentityFormat = gcpurls.Template[ComputeHTTPHealthCheckIdentity](
	"compute.googleapis.com",
	"projects/{project}/global/httpHealthChecks/{httphealthcheck}",
)

// +k8s:deepcopy-gen=false
type ComputeHTTPHealthCheckIdentity struct {
	Project         string
	HttpHealthCheck string
}

func (i *ComputeHTTPHealthCheckIdentity) String() string {
	return ComputeHTTPHealthCheckIdentityFormat.ToString(*i)
}

func (i *ComputeHTTPHealthCheckIdentity) FromExternal(ref string) error {
	trimmedRef := refs.TrimComputeURIPrefix(ref)
	parsed, match, err := ComputeHTTPHealthCheckIdentityFormat.Parse(trimmedRef)
	if err != nil {
		return fmt.Errorf("format of ComputeHTTPHealthCheck external=%q was not known (use %s): %w", ref, ComputeHTTPHealthCheckIdentityFormat.CanonicalForm(), err)
	}
	if !match {
		return fmt.Errorf("format of ComputeHTTPHealthCheck external=%q was not known (use %s)", ref, ComputeHTTPHealthCheckIdentityFormat.CanonicalForm())
	}

	*i = *parsed
	return nil
}

func (i *ComputeHTTPHealthCheckIdentity) Host() string {
	return ComputeHTTPHealthCheckIdentityFormat.Host()
}

func (r *ComputeHTTPHealthCheckRef) ValidateExternal(ref string) error {
	id := &ComputeHTTPHealthCheckIdentity{}
	if err := id.FromExternal(ref); err != nil {
		return err
	}
	return nil
}

func (r *ComputeHTTPHealthCheckRef) ParseExternalToIdentity() (identity.Identity, error) {
	id := &ComputeHTTPHealthCheckIdentity{}
	if err := id.FromExternal(r.External); err != nil {
		return nil, err
	}
	return id, nil
}

func (r *ComputeHTTPHealthCheckRef) Normalize(ctx context.Context, reader client.Reader, defaultNamespace string) error {
	if r.External != "" {
		r.External = refs.TrimComputeURIPrefix(r.External)
	}

	fallback := func(u *unstructured.Unstructured) string {
		selfLink, _, _ := unstructured.NestedString(u.Object, "status", "selfLink")
		if selfLink != "" {
			trimmed := refs.TrimComputeURIPrefix(selfLink)
			id := &ComputeHTTPHealthCheckIdentity{}
			if err := id.FromExternal(trimmed); err == nil {
				return trimmed
			}
		}
		return ""
	}

	return refs.NormalizeWithFallback(ctx, reader, r, defaultNamespace, fallback)
}
