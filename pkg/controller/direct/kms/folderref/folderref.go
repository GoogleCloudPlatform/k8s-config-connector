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

package folderref

import (
	"context"
	"fmt"
	"strings"

	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/clients/generated/apis/k8s/v1alpha1"
	apierrors "k8s.io/apimachinery/pkg/api/errors"
	refs "github.com/GoogleCloudPlatform/k8s-config-connector/apis/refs/v1beta1"

	"k8s.io/apimachinery/pkg/apis/meta/v1/unstructured"
	"k8s.io/apimachinery/pkg/runtime/schema"
	"k8s.io/apimachinery/pkg/types"
	"sigs.k8s.io/controller-runtime/pkg/client"
)

// The Folder that this resource belongs to.
type FolderRef struct {
	/* The `folderID` field of a folder, when not managed by Config Connector. */
	External string `json:"external,omitempty"`
	/* The `name` field of a `Folder` resource. */
	Name string `json:"name,omitempty"`
	/* The `namespace` field of a `Folder` resource. */
	Namespace string `json:"namespace,omitempty"`
	// The kind of the Folder resource; optional but must be `Folder` if provided.
	// +optional
	Kind string `json:"kind,omitempty"`
}

// AsFolderRef converts a generic ResourceRef into a FolderRef
func AsFolderRef(in *v1alpha1.ResourceRef) *FolderRef {
	if in == nil {
		return nil
	}
	return &FolderRef{
		Namespace: in.Namespace,
		Name:      in.Name,
		External:  in.External,
		Kind:      in.Kind,
	}
}

type Folder struct {
	FolderID string
}

// ResolveFolderFromAnnotation resolves the folderID to use for a resource,
// it should be used for resources which do not have a folderRef
func ResolveFolderFromAnnotation(ctx context.Context, reader client.Reader, src client.Object) (*Folder, error) {
	if folderID := src.GetAnnotations()["cnrm.cloud.google.com/folder-id"]; folderID != "" {
		return &Folder{FolderID: folderID}, nil
	}

	return nil, fmt.Errorf("folder-id annotation not set on resource")
}

// ResolveFolder will resolve a FolderRef to a Folder, with the FolderID
func ResolveFolder(ctx context.Context, reader client.Reader, src client.Object, ref *FolderRef) (*Folder, error) {
	if ref == nil {
		return nil, nil
	}

	if ref.Kind != "" {
		if ref.Kind != "Folder" {
			return nil, fmt.Errorf("kind is optional on folder reference, but must be \"Folder\" if provided")
		}
	}

	if ref.External != "" {
		if ref.Name != "" {
			return nil, fmt.Errorf("cannot specify both name and external on folder reference")
		}

		tokens := strings.Split(ref.External, "/")
		if len(tokens) == 1 {
			return &Folder{FolderID: tokens[0]}, nil
		}
		if len(tokens) == 2 && tokens[0] == "folders" {
			return &Folder{FolderID: tokens[1]}, nil
		}
		return nil, fmt.Errorf("format of folder external=%q was not known (use folders/<folderId> or <folderId>)", ref.External)
	}

	if ref.Name == "" {
		return nil, fmt.Errorf("must specify either name or external on folder reference")
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

	folderID, err := refs.GetResourceID(folder)
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
