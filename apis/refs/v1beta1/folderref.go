// Copyright 2024 Google LLC
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

	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/clients/generated/apis/k8s/v1alpha1"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/k8s"

	apierrors "k8s.io/apimachinery/pkg/api/errors"
	"k8s.io/apimachinery/pkg/apis/meta/v1/unstructured"
	"k8s.io/apimachinery/pkg/runtime/schema"
	"k8s.io/apimachinery/pkg/types"
	"sigs.k8s.io/controller-runtime/pkg/client"
)

// FolderRef represents the Folder that this resource belongs to.
type FolderRef struct {
	// The 'name' field of a folder, when not managed by Config Connector.
	// This field must be set when 'name' field is not set.
	// +optional
	External string `json:"external,omitempty"`
	// The 'name' field of a 'Folder' resource.
	// This field must be set when 'external' field is not set.
	// +optional
	Name string `json:"name,omitempty"`
	// The 'namespace' field of a 'Folder' resource.
	// If unset, the namespace is defaulted to the namespace of the referencer
	// resource.
	// +optional
	Namespace string `json:"namespace,omitempty"`
}

// AsFolderRef converts a generic ResourceRef into a FolderRef.
func AsFolderRef(in *v1alpha1.ResourceRef) *FolderRef {
	if in == nil {
		return nil
	}
	return &FolderRef{
		Namespace: in.Namespace,
		Name:      in.Name,
		External:  in.External,
	}
}

type Folder struct {
	FolderID string
}

// ResolveFolderFromAnnotation resolves the FolderID to use for a resource,
// it should be used for resources which do not have 'spec.folderRef'.
func ResolveFolderFromAnnotation(ctx context.Context, reader client.Reader, src client.Object) (*Folder, error) {
	if folderID := src.GetAnnotations()["cnrm.cloud.google.com/folder-id"]; folderID != "" {
		return &Folder{FolderID: folderID}, nil
	}

	return nil, fmt.Errorf("folder-id annotation not set on resource")
}

// ResolveFolder will resolve a FolderRef to a Folder, with the FolderID.
func ResolveFolder(ctx context.Context, reader client.Reader, src client.Object, ref *FolderRef) (*Folder, error) {
	if ref == nil {
		return nil, nil
	}

	if ref.External != "" {
		if ref.Name != "" {
			return nil, fmt.Errorf("cannot specify both 'name' and 'external' in 'folderRef'")
		}

		tokens := strings.Split(ref.External, "/")
		if len(tokens) == 2 && tokens[0] == "folders" {
			return &Folder{FolderID: tokens[1]}, nil
		}
		return nil, fmt.Errorf("format of 'folderRef.external'=%q was not known (use folders/<folderID>)", ref.External)
	}

	if ref.Name == "" {
		return nil, fmt.Errorf("must specify either 'name' and 'external' in 'folderRef'")
	}

	key := types.NamespacedName{
		Namespace: ref.Namespace,
		Name:      ref.Name,
	}
	if key.Namespace == "" {
		key.Namespace = src.GetNamespace()
	}

	folder := &unstructured.Unstructured{}
	folder.SetGroupVersionKind(schema.GroupVersionKind{
		Group:   "resourcemanager.cnrm.cloud.google.com",
		Version: "v1beta1",
		Kind:    "Folder",
	})
	if err := reader.Get(ctx, key, folder); err != nil {
		if apierrors.IsNotFound(err) {
			return nil, fmt.Errorf("referenced Folder %v not found", key)
		}
		return nil, fmt.Errorf("error reading referenced Folder %v: %w", key, err)
	}
	resource, err := k8s.NewResource(folder)
	if err != nil {
		return nil, fmt.Errorf("error converting unstructured to resource: %w", err)
	}
	if !k8s.IsResourceReady(resource) {
		return nil, k8s.NewReferenceNotReadyError(folder.GroupVersionKind(), key)
	}

	folderID, err := GetResourceID(folder)
	if err != nil {
		return nil, err
	}

	return &Folder{
		FolderID: folderID,
	}, nil
}

func ResolveFolderID(ctx context.Context, reader client.Reader, obj *unstructured.Unstructured) (string, error) {
	folderRefExternal, _, _ := unstructured.NestedString(obj.Object, "spec", "folderRef", "external")
	if folderRefExternal != "" {
		folderRef := FolderRef{
			External: folderRefExternal,
		}

		folder, err := ResolveFolder(ctx, reader, obj, &folderRef)
		if err != nil {
			return "", fmt.Errorf("cannot parse folderRef.external %q in %v %v/%v: %w", folderRefExternal, obj.GetKind(), obj.GetNamespace(), obj.GetName(), err)
		}
		return folder.FolderID, nil
	}

	folderRefName, _, _ := unstructured.NestedString(obj.Object, "spec", "folderRef", "name")
	if folderRefName != "" {
		folderRefNamespace, _, _ := unstructured.NestedString(obj.Object, "spec", "folderRef", "namespace")

		folderRef := FolderRef{
			Name:      folderRefName,
			Namespace: folderRefNamespace,
		}
		if folderRef.Namespace == "" {
			folderRef.Namespace = obj.GetNamespace()
		}

		folder, err := ResolveFolder(ctx, reader, obj, &folderRef)
		if err != nil {
			return "", fmt.Errorf("cannot parse folderRef in %v %v/%v: %w", obj.GetKind(), obj.GetNamespace(), obj.GetName(), err)
		}
		return folder.FolderID, nil
	}

	if folderID := obj.GetAnnotations()["cnrm.cloud.google.com/folder-id"]; folderID != "" {
		return folderID, nil
	}

	return "", fmt.Errorf("cannot find folder id for %v %v/%v", obj.GetKind(), obj.GetNamespace(), obj.GetName())
}
