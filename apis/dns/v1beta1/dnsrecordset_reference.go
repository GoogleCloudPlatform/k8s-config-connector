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

	"github.com/GoogleCloudPlatform/k8s-config-connector/apis/common"
	"github.com/GoogleCloudPlatform/k8s-config-connector/apis/common/identity"
	computerefs "github.com/GoogleCloudPlatform/k8s-config-connector/apis/compute/computerefs"
	computev1beta1 "github.com/GoogleCloudPlatform/k8s-config-connector/apis/compute/v1beta1"
	refs "github.com/GoogleCloudPlatform/k8s-config-connector/apis/refs/v1beta1"
	"k8s.io/apimachinery/pkg/apis/meta/v1/unstructured"
	"k8s.io/apimachinery/pkg/runtime/schema"
	"k8s.io/apimachinery/pkg/types"
	"sigs.k8s.io/controller-runtime/pkg/client"
)

var _ refs.Ref = &DNSRecordSetRef{}

// DNSRecordSetRef is a reference to a DNSRecordSet.
type DNSRecordSetRef struct {
	// A reference to an externally managed DNSRecordSet resource.
	// Should be in the format "projects/{project}/locations/{location}/managedZones/{managedZone}/rrsets/{name} or projects/{project}/managedZones/{managedZone}/rrsets/{name}".
	External string `json:"external,omitempty"`

	// The name of a DNSRecordSet resource.
	Name string `json:"name,omitempty"`

	// The namespace of a DNSRecordSet resource.
	Namespace string `json:"namespace,omitempty"`
}

func init() {
	refs.Register(&DNSRecordSetRef{}, &DNSRecordSet{})
	refs.Register(&RecordsetRrdatasRefs{})
}

func (r *DNSRecordSetRef) GetGVK() schema.GroupVersionKind {
	return DNSRecordSetGVK
}

func (r *DNSRecordSetRef) GetNamespacedName() types.NamespacedName {
	return types.NamespacedName{
		Name:      r.Name,
		Namespace: r.Namespace,
	}
}

func (r *DNSRecordSetRef) GetExternal() string {
	return r.External
}

func (r *DNSRecordSetRef) SetExternal(ref string) {
	r.External = ref
	r.Name = ""
	r.Namespace = ""
}

func (r *DNSRecordSetRef) ValidateExternal(ref string) error {
	id := &DNSRecordSetIdentity{}
	if err := id.FromExternal(ref); err != nil {
		return err
	}
	return nil
}

func (r *DNSRecordSetRef) ParseExternalToIdentity() (identity.Identity, error) {
	id := &DNSRecordSetIdentity{}
	if err := id.FromExternal(r.External); err != nil {
		return nil, err
	}
	return id, nil
}

func (r *DNSRecordSetRef) Normalize(ctx context.Context, reader client.Reader, defaultNamespace string) error {
	fallback := func(u *unstructured.Unstructured) string {
		typed, err := common.ToStructuredType[*DNSRecordSet](u)
		if err != nil {
			return ""
		}
		identity, err := getIdentityFromDNSRecordSetSpec(ctx, reader, typed)
		if err != nil {
			return ""
		}
		return identity.String()
	}
	return refs.NormalizeWithFallback(ctx, reader, r, defaultNamespace, fallback)
}

var _ refs.Ref = &RecordsetRrdatasRefs{}

func (r *RecordsetRrdatasRefs) GetGVK() schema.GroupVersionKind {
	return computev1beta1.ComputeAddressGVK
}

func (r *RecordsetRrdatasRefs) GetNamespacedName() types.NamespacedName {
	var name, namespace string
	if r.Name != nil {
		name = *r.Name
	}
	if r.Namespace != nil {
		namespace = *r.Namespace
	}
	return types.NamespacedName{
		Name:      name,
		Namespace: namespace,
	}
}

func (r *RecordsetRrdatasRefs) GetExternal() string {
	if r.External != nil {
		return *r.External
	}
	return ""
}

func (r *RecordsetRrdatasRefs) SetExternal(ref string) {
	r.External = &ref
	r.Name = nil
	r.Namespace = nil
}

func (r *RecordsetRrdatasRefs) ValidateExternal(ref string) error {
	return nil
}

func (r *RecordsetRrdatasRefs) Normalize(ctx context.Context, reader client.Reader, defaultNamespace string) error {
	if r.External != nil && *r.External != "" {
		return nil
	}

	key := r.GetNamespacedName()
	if key.Name == "" {
		return nil
	}
	if key.Namespace == "" {
		key.Namespace = defaultNamespace
	}

	ref := &computev1beta1.ComputeAddressRef{
		Name:      key.Name,
		Namespace: key.Namespace,
	}

	// Create a non-gcp resolver (config: nil) since we are resolving a Kube reference
	resolver := computerefs.NewComputeAddressResolver(nil)

	// Create a dummy source object just to pass the defaultNamespace
	src := &unstructured.Unstructured{}
	src.SetNamespace(defaultNamespace)

	ip, err := resolver.ResolveComputeAddressIP(ctx, reader, src, ref)
	if err != nil {
		return err
	}

	r.External = &ip
	return nil
}
