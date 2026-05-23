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

	"github.com/GoogleCloudPlatform/k8s-config-connector/apis/common/identity"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/gcpurls"
	apierrors "k8s.io/apimachinery/pkg/api/errors"
	"k8s.io/apimachinery/pkg/apis/meta/v1/unstructured"
	"k8s.io/apimachinery/pkg/runtime/schema"
	"k8s.io/apimachinery/pkg/types"
	"sigs.k8s.io/controller-runtime/pkg/client"
)

var PrivateCACAPoolGVK = schema.GroupVersionKind{
	Group:   "privateca.cnrm.cloud.google.com",
	Version: "v1beta1",
	Kind:    "PrivateCACAPool",
}

var (
	_ identity.IdentityV2 = &PrivateCACAPoolIdentity{}
	_ Ref                 = &PrivateCACAPoolRef{}
	_ ExternalRef         = &PrivateCACAPoolRef{}
)

var PrivateCACAPoolIdentityFormat = gcpurls.Template[PrivateCACAPoolIdentity]("privateca.googleapis.com", "projects/{project}/locations/{location}/caPools/{caPool}")

// +k8s:deepcopy-gen=false
type PrivateCACAPoolIdentity struct {
	Project  string
	Location string
	CAPool   string
}

func (i *PrivateCACAPoolIdentity) String() string {
	return PrivateCACAPoolIdentityFormat.ToString(*i)
}

func (i *PrivateCACAPoolIdentity) FromExternal(ref string) error {
	parsed, match, err := PrivateCACAPoolIdentityFormat.Parse(ref)
	if err != nil {
		return fmt.Errorf("format of PrivateCACAPool external=%q was not known (use %s): %w", ref, PrivateCACAPoolIdentityFormat.CanonicalForm(), err)
	}
	if !match {
		return fmt.Errorf("format of PrivateCACAPool external=%q was not known (use %s)", ref, PrivateCACAPoolIdentityFormat.CanonicalForm())
	}

	*i = *parsed
	return nil
}

func (i *PrivateCACAPoolIdentity) Host() string {
	return PrivateCACAPoolIdentityFormat.Host()
}

type PrivateCACAPoolRef struct {
	// A reference to an externally managed PrivateCACAPool resource.
	// Should be in the format "projects/{{projectID}}/locations/{{location}}/caPools/{{caPoolID}}".
	External string `json:"external,omitempty"`

	// The name of a PrivateCACAPool resource.
	Name string `json:"name,omitempty"`

	// The namespace of a PrivateCACAPool resource.
	Namespace string `json:"namespace,omitempty"`
}

type PrivateCACAPool struct {
	Ref        *PrivateCACAPoolRef
	ResourceID string
}

func init() {
	Register(&PrivateCACAPoolRef{})
}

func (r *PrivateCACAPoolRef) GetGVK() schema.GroupVersionKind {
	return PrivateCACAPoolGVK
}

func (r *PrivateCACAPoolRef) GetNamespacedName() types.NamespacedName {
	return types.NamespacedName{
		Name:      r.Name,
		Namespace: r.Namespace,
	}
}

func (r *PrivateCACAPoolRef) GetExternal() string {
	return r.External
}

func (r *PrivateCACAPoolRef) SetExternal(ref string) {
	r.External = ref
}

func (r *PrivateCACAPoolRef) ValidateExternal(ref string) error {
	id := &PrivateCACAPoolIdentity{}
	if err := id.FromExternal(ref); err != nil {
		return err
	}
	return nil
}

func (r *PrivateCACAPoolRef) ParseExternalToIdentity() (identity.Identity, error) {
	id := &PrivateCACAPoolIdentity{}
	if err := id.FromExternal(r.External); err != nil {
		return nil, err
	}
	return id, nil
}

func (r *PrivateCACAPoolRef) Normalize(ctx context.Context, reader client.Reader, defaultNamespace string) error {
	fallback := func(u *unstructured.Unstructured) string {
		identity, err := GetIdentityFromPrivateCACAPoolSpec(ctx, reader, u)
		if err != nil {
			return ""
		}
		return identity.String()
	}
	return NormalizeWithFallback(ctx, reader, r, defaultNamespace, fallback)
}

func GetIdentityFromPrivateCACAPoolSpec(ctx context.Context, reader client.Reader, obj client.Object) (*PrivateCACAPoolIdentity, error) {
	resourceID, err := GetResourceID(obj)
	if err != nil {
		return nil, fmt.Errorf("cannot resolve resource ID")
	}

	location, err := GetLocation(obj)
	if err != nil {
		return nil, fmt.Errorf("cannot resolve location")
	}

	projectID, err := ResolveProjectID(ctx, reader, obj)
	if err != nil {
		return nil, fmt.Errorf("cannot resolve project")
	}

	identity := &PrivateCACAPoolIdentity{
		Project:  projectID,
		Location: location,
		CAPool:   resourceID,
	}
	return identity, nil
}

// ResolvePrivateCACAPoolRef will resolve a PrivateCACAPoolRef to a PrivateCACAPool.
// Deprecated: Use Normalize or generic reference resolution instead.
func ResolvePrivateCACAPoolRef(ctx context.Context, reader client.Reader, src client.Object, ref *PrivateCACAPoolRef) (*PrivateCACAPoolRef, error) {
	if ref == nil {
		return nil, nil
	}

	if ref.Name == "" && ref.External == "" {
		return nil, fmt.Errorf("must specify either name or external on PrivateCACAPoolRef")
	}
	if ref.Name != "" && ref.External != "" {
		return nil, fmt.Errorf("cannot specify both name and external on PrivateCACAPoolRef")
	}

	// External should be in the `projects/{project_id}/locations/{region}/caPools/{caPool}` format
	if ref.External != "" {
		tokens := strings.Split(ref.External, "/")
		if len(tokens) == 6 && tokens[0] == "projects" && tokens[2] == "locations" && tokens[4] == "caPools" {
			ref = &PrivateCACAPoolRef{
				External: fmt.Sprintf("projects/%s/locations/%s/caPools/%s", tokens[1], tokens[3], tokens[5]),
			}
			return ref, nil
		}
		return nil, fmt.Errorf("format of PrivateCACAPoolRef external=%q was not known (use projects/{project_id}/locations/{region}/caPools/{caPool})", ref.External)
	}

	key := types.NamespacedName{
		Namespace: ref.Namespace,
		Name:      ref.Name,
	}
	if key.Namespace == "" {
		key.Namespace = src.GetNamespace()
	}

	// Fetch object from k8s cluster to construct the external form
	caPool := &unstructured.Unstructured{}
	caPool.SetGroupVersionKind(PrivateCACAPoolGVK)
	if err := reader.Get(ctx, key, caPool); err != nil {
		if apierrors.IsNotFound(err) {
			return nil, fmt.Errorf("referenced PrivateCACAPool %v not found", key)
		}
		return nil, fmt.Errorf("error reading referenced PrivateCACAPool %v: %w", key, err)
	}

	caPoolResourceID, err := GetResourceID(caPool)
	if err != nil {
		return nil, err
	}

	projectID, err := ResolveProjectID(ctx, reader, caPool)
	if err != nil {
		return nil, err
	}

	location, err := GetLocation(caPool)
	if err != nil {
		return nil, err
	}

	ref = &PrivateCACAPoolRef{
		External: fmt.Sprintf("projects/%s/locations/%s/caPools/%s", projectID, location, caPoolResourceID),
	}

	return ref, nil
}
