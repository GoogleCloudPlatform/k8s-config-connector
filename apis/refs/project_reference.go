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

package refs

import (
	"context"

	"github.com/GoogleCloudPlatform/k8s-config-connector/apis/common/identity"
	refsv1beta1 "github.com/GoogleCloudPlatform/k8s-config-connector/apis/refs/v1beta1"
	"k8s.io/apimachinery/pkg/runtime/schema"
	"k8s.io/apimachinery/pkg/types"
	"sigs.k8s.io/controller-runtime/pkg/client"
)

var ProjectGVK = schema.GroupVersionKind{
	Group:   "resourcemanager.cnrm.cloud.google.com",
	Version: "v1beta1",
	Kind:    "Project",
}

// ProjectRef is a reference to a Google Cloud Project.
type ProjectRef struct {
	/* The `projectID` field of a project, when not managed by Config Connector. */
	External string `json:"external,omitempty"`
	/* The `name` field of a `Project` resource. */
	Name string `json:"name,omitempty"`
	/* The `namespace` field of a `Project` resource. */
	Namespace string `json:"namespace,omitempty"`
}

func (r *ProjectRef) GetGVK() schema.GroupVersionKind {
	return ProjectGVK
}

func (r *ProjectRef) GetNamespacedName() types.NamespacedName {
	return types.NamespacedName{
		Name:      r.Name,
		Namespace: r.Namespace,
	}
}

func (r *ProjectRef) GetExternal() string {
	return r.External
}

func (r *ProjectRef) SetExternal(ref string) {
	r.External = ref
	r.Name = ""
	r.Namespace = ""
}

func (r *ProjectRef) Normalize(ctx context.Context, reader client.Reader, defaultNamespace string) error {
	v1beta1Ref := &refsv1beta1.ProjectRef{
		External:  r.External,
		Name:      r.Name,
		Namespace: r.Namespace,
	}
	if err := v1beta1Ref.Normalize(ctx, reader, defaultNamespace); err != nil {
		return err
	}
	r.External = v1beta1Ref.External
	return nil
}

func (r *ProjectRef) ValidateExternal(ref string) error {
	id := &refsv1beta1.ProjectIdentity{}
	if err := id.FromExternal(ref); err != nil {
		return err
	}
	return nil
}

func (r *ProjectRef) ParseExternalToIdentity() (identity.Identity, error) {
	id := &refsv1beta1.ProjectIdentity{}
	if err := id.FromExternal(r.External); err != nil {
		return nil, err
	}
	return id, nil
}
