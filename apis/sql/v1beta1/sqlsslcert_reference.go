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
	refs "github.com/GoogleCloudPlatform/k8s-config-connector/apis/refs/v1beta1"
	"k8s.io/apimachinery/pkg/apis/meta/v1/unstructured"
	"k8s.io/apimachinery/pkg/runtime/schema"
	"k8s.io/apimachinery/pkg/types"
	"sigs.k8s.io/controller-runtime/pkg/client"
)

var SQLSSLCertGVK = schema.GroupVersionKind{
	Group:   "sql.cnrm.cloud.google.com",
	Version: "v1beta1",
	Kind:    "SQLSSLCert",
}

var _ refs.Ref = &SQLSSLCertRef{}

// SQLSSLCertRef is a reference to a GCP SQLSSLCert.
type SQLSSLCertRef struct {
	// A reference to an externally managed SQLSSLCert resource.
	// Should be in the format "projects/{{projectID}}/instances/{{instance}}/sslCerts/{{sha1Fingerprint}}".
	External string `json:"external,omitempty"`

	// The name of a SQLSSLCert resource.
	Name string `json:"name,omitempty"`

	// The namespace of a SQLSSLCert resource.
	Namespace string `json:"namespace,omitempty"`
}

func init() {
	refs.Register(&SQLSSLCertRef{}, &SQLSSLCert{})
}

func (r *SQLSSLCertRef) GetGVK() schema.GroupVersionKind {
	return SQLSSLCertGVK
}

func (r *SQLSSLCertRef) GetNamespacedName() types.NamespacedName {
	return types.NamespacedName{
		Name:      r.Name,
		Namespace: r.Namespace,
	}
}

func (r *SQLSSLCertRef) GetExternal() string {
	return r.External
}

func (r *SQLSSLCertRef) SetExternal(ref string) {
	r.External = ref
	r.Name = ""
	r.Namespace = ""
}

func (r *SQLSSLCertRef) ValidateExternal(ref string) error {
	id := &SQLSSLCertIdentity{}
	return id.FromExternal(ref)
}

func (r *SQLSSLCertRef) ParseExternalToIdentity() (identity.Identity, error) {
	id := &SQLSSLCertIdentity{}
	if err := id.FromExternal(r.External); err != nil {
		return nil, err
	}
	return id, nil
}

func (r *SQLSSLCertRef) Normalize(ctx context.Context, reader client.Reader, defaultNamespace string) error {
	fallback := func(u *unstructured.Unstructured) string {
		var ready bool
		conditions, _, _ := unstructured.NestedSlice(u.Object, "status", "conditions")
		for _, condition := range conditions {
			if cond, ok := condition.(map[string]interface{}); ok {
				if cond["type"] == "Ready" && cond["status"] == "True" {
					ready = true
					break
				}
			}
		}
		if !ready {
			return ""
		}

		obj, err := common.ToStructuredType[*SQLSSLCert](u)
		if err != nil {
			return ""
		}
		id, err := obj.GetIdentity(ctx, reader)
		if err != nil {
			return ""
		}
		serverGen, ok := id.(identity.ServerGeneratedIdentity)
		if ok && !serverGen.HasIdentitySpecified() {
			return ""
		}
		return id.String()
	}
	return refs.NormalizeWithFallback(ctx, reader, r, defaultNamespace, fallback)
}
