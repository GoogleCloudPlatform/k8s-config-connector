// Copyright 2025 Google LLC
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//    http://www.apache.org/licenses/LICENSE-2.0
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
	"strings"

	refsv1beta1 "github.com/GoogleCloudPlatform/k8s-config-connector/apis/refs/v1beta1"
	apierrors "k8s.io/apimachinery/pkg/api/errors"
	"k8s.io/apimachinery/pkg/apis/meta/v1/unstructured"
	"k8s.io/apimachinery/pkg/runtime/schema"
	"k8s.io/apimachinery/pkg/types"
	"sigs.k8s.io/controller-runtime/pkg/client"
)

// BackendBucketSignedURLKeyIdentity defines the resource reference to ComputeBackendBucketSignedURLKey.
// The identity is (project, backendBucket, keyName).
type BackendBucketSignedURLKeyIdentity struct {
	parent  *BackendBucketSignedURLKeyParent
	keyName string
}

// BackendBucketSignedURLKeyParent holds the parent fields for a BackendBucketSignedURLKey.
type BackendBucketSignedURLKeyParent struct {
	ProjectID    string
	BackendBucket string
}

func (p *BackendBucketSignedURLKeyParent) String() string {
	return "projects/" + p.ProjectID + "/global/backendBuckets/" + p.BackendBucket
}

// String returns the canonical external reference string.
// Format: projects/{project}/global/backendBuckets/{bucket}/signedUrlKeys/{keyName}
func (i *BackendBucketSignedURLKeyIdentity) String() string {
	return i.parent.String() + "/signedUrlKeys/" + i.keyName
}

func (i *BackendBucketSignedURLKeyIdentity) Parent() *BackendBucketSignedURLKeyParent {
	return i.parent
}

func (i *BackendBucketSignedURLKeyIdentity) KeyName() string {
	return i.keyName
}

// NewBackendBucketSignedURLKeyIdentity builds an identity from the Config Connector object.
func NewBackendBucketSignedURLKeyIdentity(ctx context.Context, reader client.Reader, obj *ComputeBackendBucketSignedURLKey) (*BackendBucketSignedURLKeyIdentity, error) {
	projectRef, err := refsv1beta1.ResolveProject(ctx, reader, obj.GetNamespace(), obj.Spec.ProjectRef)
	if err != nil {
		return nil, err
	}
	projectID := projectRef.ProjectID
	if projectID == "" {
		return nil, fmt.Errorf("cannot resolve project")
	}

	bucketName, err := resolveBackendBucketName(ctx, reader, obj)
	if err != nil {
		return nil, err
	}

	// keyName is spec.resourceID if set, otherwise metadata.name.
	keyName := obj.GetName()
	if obj.Spec.ResourceID != nil && *obj.Spec.ResourceID != "" {
		keyName = *obj.Spec.ResourceID
	}
	if keyName == "" {
		return nil, fmt.Errorf("cannot determine key name: metadata.name and spec.resourceID are both empty")
	}

	return &BackendBucketSignedURLKeyIdentity{
		parent: &BackendBucketSignedURLKeyParent{
			ProjectID:    projectID,
			BackendBucket: bucketName,
		},
		keyName: keyName,
	}, nil
}

// resolveBackendBucketName resolves the BackendBucketRef to a bucket name.
// If external is set it is used directly. If name is set, the ComputeBackendBucket
// resource is looked up and its resourceID (or metadata.name) is used.
func resolveBackendBucketName(ctx context.Context, reader client.Reader, obj *ComputeBackendBucketSignedURLKey) (string, error) {
	ref := obj.Spec.BackendBucketRef
	if ref.External != "" {
		if ref.Name != "" {
			return "", fmt.Errorf("cannot specify both name and external on backendBucketRef")
		}
		// External may be either just the name or a full resource path.
		// Normalize: if it contains slashes, extract the name segment.
		parts := strings.Split(ref.External, "/")
		return parts[len(parts)-1], nil
	}

	if ref.Name == "" {
		return "", fmt.Errorf("must specify either name or external on backendBucketRef")
	}

	ns := ref.Namespace
	if ns == "" {
		ns = obj.GetNamespace()
	}

	bucketObj := &unstructured.Unstructured{}
	bucketObj.SetGroupVersionKind(schema.GroupVersionKind{
		Group:   "compute.cnrm.cloud.google.com",
		Version: "v1beta1",
		Kind:    "ComputeBackendBucket",
	})
	key := types.NamespacedName{Namespace: ns, Name: ref.Name}
	if err := reader.Get(ctx, key, bucketObj); err != nil {
		if apierrors.IsNotFound(err) {
			return "", fmt.Errorf("referenced ComputeBackendBucket %v not found", key)
		}
		return "", fmt.Errorf("error reading referenced ComputeBackendBucket %v: %w", key, err)
	}

	// Use spec.resourceID if set, otherwise metadata.name.
	resourceID, _, _ := unstructured.NestedString(bucketObj.Object, "spec", "resourceID")
	if resourceID != "" {
		return resourceID, nil
	}
	return bucketObj.GetName(), nil
}

// ParseBackendBucketSignedURLKeyExternal parses an external reference string.
// Expected format:
//
//	projects/{project}/global/backendBuckets/{bucket}/signedUrlKeys/{keyName}
//
// Token indices (0-based): 0=projects 1={project} 2=global 3=backendBuckets 4={bucket} 5=signedUrlKeys 6={keyName}
func ParseBackendBucketSignedURLKeyExternal(external string) (*BackendBucketSignedURLKeyIdentity, error) {
	tokens := strings.Split(external, "/")
	// 7 tokens: projects / {project} / global / backendBuckets / {bucket} / signedUrlKeys / {keyName}
	if len(tokens) != 7 {
		return nil, fmt.Errorf("format of ComputeBackendBucketSignedURLKey external=%q was not known (expected 7 tokens, got %d)", external, len(tokens))
	}
	if tokens[0] != "projects" || tokens[2] != "global" || tokens[3] != "backendBuckets" || tokens[5] != "signedUrlKeys" {
		return nil, fmt.Errorf("format of ComputeBackendBucketSignedURLKey external=%q was not known", external)
	}
	if tokens[1] == "" || tokens[4] == "" || tokens[6] == "" {
		return nil, fmt.Errorf("format of ComputeBackendBucketSignedURLKey external=%q was not known (empty segment)", external)
	}

	return &BackendBucketSignedURLKeyIdentity{
		parent: &BackendBucketSignedURLKeyParent{
			ProjectID:    tokens[1],
			BackendBucket: tokens[4],
		},
		keyName: tokens[6],
	}, nil
}
