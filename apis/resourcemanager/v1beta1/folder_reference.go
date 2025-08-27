// Copyright 2025 Google LLC
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

	refsv1beta1 "github.com/GoogleCloudPlatform/k8s-config-connector/apis/refs/v1beta1"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/k8s"
	apierrors "k8s.io/apimachinery/pkg/api/errors"
	"k8s.io/apimachinery/pkg/apis/meta/v1/unstructured"
	"k8s.io/apimachinery/pkg/runtime/schema"
	"k8s.io/apimachinery/pkg/types"
	"sigs.k8s.io/controller-runtime/pkg/client"
)

var _ refsv1beta1.Ref = &FolderRef{}
var FolderGVK = GroupVersion.WithKind("Folder")

// FolderRef represents the Folder that this resource belongs to.
type FolderRef struct {
	// A reference to an externally managed Folder resource.
	// Should be in the format "folders/{{folderID}}" or "{{folderID}}".
	External string `json:"external,omitempty"`
	// The name of a Folder resource.
	Name string `json:"name,omitempty"`
	// The namespace of a Folder resource.
	Namespace string `json:"namespace,omitempty"`
}

func (r *FolderRef) GetGVK() schema.GroupVersionKind {
	return FolderGVK
}

func (r *FolderRef) GetNamespacedName() types.NamespacedName {
	return types.NamespacedName{}
}

func (r *FolderRef) GetExternal() string {
	return r.External
}

func (r *FolderRef) SetExternal(external string) {
	r.External = external
}

func (r *FolderRef) ValidateExternal(external string) error {
	if _, err := ParseFolderExternal(external); err != nil {
		return err
	}
	return nil
}

func (r *FolderRef) Normalize(ctx context.Context, reader client.Reader, defaultNamespace string) error {
	if r.External == "" {
		key := r.GetNamespacedName()
		if key.Namespace == "" {
			key.Namespace = defaultNamespace
		}
		u := &unstructured.Unstructured{}
		u.SetGroupVersionKind(r.GetGVK())
		if err := reader.Get(ctx, key, u); err != nil {
			if apierrors.IsNotFound(err) {
				return k8s.NewReferenceNotFoundError(u.GroupVersionKind(), key)
			}
			return fmt.Errorf("reading referenced %s %s: %w", r.GetGVK(), key, err)
		}

		folderID, err := refsv1beta1.GetResourceID(u)
		if err != nil {
			return err
		}
		r.SetExternal("folders/" + folderID)
		return nil
	}
	return r.ValidateExternal(r.External)
}

// ResolveFolderFromAnnotation resolves the FolderID to use for a resource,
// it should be used for resources which do not have 'spec.folderRef'.
// todo: Identify the resources or use cases where this function is necessary.
func ResolveFolderFromAnnotation(ctx context.Context, reader client.Reader, src client.Object) (string, error) {
	if folderID := src.GetAnnotations()["cnrm.cloud.google.com/folder-id"]; folderID != "" {
		return folderID, nil
	}
	return "", fmt.Errorf("folder-id annotation not set on resource")
}
