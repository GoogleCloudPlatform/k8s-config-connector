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
	"k8s.io/apimachinery/pkg/runtime/schema"
	"k8s.io/apimachinery/pkg/types"
	"sigs.k8s.io/controller-runtime/pkg/client"
)

var AppEngineServiceGVK = GroupVersion.WithKind("AppEngineService")

var _ refs.Ref = &AppEngineServiceRef{}
var _ refs.ExternalRef = &AppEngineServiceRef{}

type AppEngineServiceRef struct {
	// Allowed value: The `name` field of an `AppEngineService` resource.
	External string `json:"external,omitempty"`

	// Name of the referent. More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#names
	Name string `json:"name,omitempty"`

	// Namespace of the referent. More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/namespaces/
	Namespace string `json:"namespace,omitempty"`
}

func init() {
	refs.Register(&AppEngineServiceRef{})
}

func (r *AppEngineServiceRef) GetGVK() schema.GroupVersionKind {
	return AppEngineServiceGVK
}

func (r *AppEngineServiceRef) GetNamespacedName() types.NamespacedName {
	return types.NamespacedName{
		Name:      r.Name,
		Namespace: r.Namespace,
	}
}

func (r *AppEngineServiceRef) GetExternal() string {
	return r.External
}

func (r *AppEngineServiceRef) SetExternal(ref string) {
	r.External = ref
	r.Name = ""
	r.Namespace = ""
}

func (r *AppEngineServiceRef) Normalize(ctx context.Context, reader client.Reader, defaultNamespace string) error {
	return refs.Normalize(ctx, reader, r, defaultNamespace)
}

var AppEngineServiceIdentityFormat = gcpurls.Template[AppEngineServiceIdentity]("appengine.googleapis.com", "apps/{app}/services/{service}")

var _ identity.Identity = &AppEngineServiceIdentity{}
var _ identity.IdentityV2 = &AppEngineServiceIdentity{}

// AppEngineServiceIdentity is the identity of a GCP AppEngineService resource.
// +k8s:deepcopy-gen=false
type AppEngineServiceIdentity struct {
	App     string
	Service string
}

func (i *AppEngineServiceIdentity) String() string {
	return AppEngineServiceIdentityFormat.ToString(*i)
}

func (i *AppEngineServiceIdentity) Host() string {
	return AppEngineServiceIdentityFormat.Host()
}

func (i *AppEngineServiceIdentity) ParentString() string {
	return "apps/" + i.App
}

func (i *AppEngineServiceIdentity) FromExternal(ref string) error {
	parsed, match, err := AppEngineServiceIdentityFormat.Parse(ref)
	if err != nil {
		return fmt.Errorf("format of AppEngineService external=%q was not known (use %s): %w", ref, AppEngineServiceIdentityFormat.CanonicalForm(), err)
	}
	if !match {
		return fmt.Errorf("format of AppEngineService external=%q was not known (use %s)", ref, AppEngineServiceIdentityFormat.CanonicalForm())
	}

	*i = *parsed
	return nil
}

func (r *AppEngineServiceRef) ValidateExternal(ref string) error {
	id := &AppEngineServiceIdentity{}
	if err := id.FromExternal(ref); err != nil {
		return err
	}
	return nil
}

func (r *AppEngineServiceRef) ParseExternalToIdentity() (identity.Identity, error) {
	id := &AppEngineServiceIdentity{}
	if err := id.FromExternal(r.External); err != nil {
		return nil, err
	}
	return id, nil
}
