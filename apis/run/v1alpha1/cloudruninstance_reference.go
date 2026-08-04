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

	"github.com/GoogleCloudPlatform/k8s-config-connector/apis/common/identity"
	refsv1beta1 "github.com/GoogleCloudPlatform/k8s-config-connector/apis/refs/v1beta1"
	"k8s.io/apimachinery/pkg/runtime/schema"
	"k8s.io/apimachinery/pkg/types"
	"sigs.k8s.io/controller-runtime/pkg/client"
)

var _ refsv1beta1.Ref = &CloudRunInstanceRef{}

func init() {
	refsv1beta1.Register(&CloudRunInstanceRef{}, &CloudRunInstance{})
}

// CloudRunInstanceRef is a reference to a GCP CloudRunInstance.
type CloudRunInstanceRef struct {
	// A reference to an externally managed CloudRunInstance resource. Should be in the format "projects/{{projectID}}/locations/{{location}}/instances/{{instanceID}}".
	External string `json:"external,omitempty"`

	// The name of a CloudRunInstance resource.
	Name string `json:"name,omitempty"`

	// The namespace of a CloudRunInstance resource.
	Namespace string `json:"namespace,omitempty"`
}

func (r *CloudRunInstanceRef) GetGVK() schema.GroupVersionKind {
	return CloudRunInstanceGVK
}

func (r *CloudRunInstanceRef) GetNamespacedName() types.NamespacedName {
	return types.NamespacedName{
		Name:      r.Name,
		Namespace: r.Namespace,
	}
}

func (r *CloudRunInstanceRef) GetExternal() string {
	return r.External
}

func (r *CloudRunInstanceRef) SetExternal(ref string) {
	r.External = ref
	r.Name = ""
	r.Namespace = ""
}

func (r *CloudRunInstanceRef) ValidateExternal(ref string) error {
	id := &CloudRunInstanceIdentity{}
	if err := id.FromExternal(ref); err != nil {
		return err
	}
	return nil
}

func (r *CloudRunInstanceRef) ParseExternalToIdentity() (identity.Identity, error) {
	id := &CloudRunInstanceIdentity{}
	if err := id.FromExternal(r.External); err != nil {
		return nil, err
	}
	return id, nil
}

func (r *CloudRunInstanceRef) Normalize(ctx context.Context, reader client.Reader, defaultNamespace string) error {
	return refsv1beta1.Normalize(ctx, reader, r, defaultNamespace)
}
