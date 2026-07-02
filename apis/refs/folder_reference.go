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
	"k8s.io/apimachinery/pkg/apis/meta/v1/unstructured"
	"k8s.io/apimachinery/pkg/runtime/schema"
	"k8s.io/apimachinery/pkg/types"
	"sigs.k8s.io/controller-runtime/pkg/client"
)

var folderGVK = schema.GroupVersionKind{
	Group:   "resourcemanager.cnrm.cloud.google.com",
	Version: "v1beta1",
	Kind:    "Folder",
}

var _ refsv1beta1.Ref = &FolderRef{}

// FolderRef is a reference to a GCP Folder.
type FolderRef struct {
	/* The 'name' field of a folder, when not managed by Config Connector. */
	// +optional
	External string `json:"external,omitempty"`
	/* The 'name' field of a 'Folder' resource. */
	// +optional
	Name string `json:"name,omitempty"`
	/* The 'namespace' field of a 'Folder' resource. */
	// +optional
	Namespace string `json:"namespace,omitempty"`
}

func init() {
	refsv1beta1.Register(&FolderRef{}, nil)
}

func (r *FolderRef) GetGVK() schema.GroupVersionKind {
	return folderGVK
}

func (r *FolderRef) GetNamespacedName() types.NamespacedName {
	return types.NamespacedName{
		Name:      r.Name,
		Namespace: r.Namespace,
	}
}

func (r *FolderRef) GetExternal() string {
	return r.External
}

func (r *FolderRef) SetExternal(ref string) {
	r.External = ref
	r.Name = ""
	r.Namespace = ""
}

func (r *FolderRef) ValidateExternal(ref string) error {
	id := &FolderIdentity{}
	if err := id.FromExternal(ref); err != nil {
		return err
	}
	return nil
}

func (r *FolderRef) ParseExternalToIdentity() (identity.Identity, error) {
	id := &FolderIdentity{}
	if err := id.FromExternal(r.External); err != nil {
		return nil, err
	}
	return id, nil
}

func (r *FolderRef) Normalize(ctx context.Context, reader client.Reader, defaultNamespace string) error {
	fallback := func(u *unstructured.Unstructured) string {
		identity, err := Folder_IdentityFromSpec(ctx, reader, u)
		if err != nil {
			return ""
		}
		return identity.String()
	}
	return refsv1beta1.NormalizeWithFallback(ctx, reader, r, defaultNamespace, fallback)
}
