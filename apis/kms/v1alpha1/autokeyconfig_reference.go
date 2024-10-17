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

package v1alpha1

import (
	"context"
	"fmt"
	"strings"

	refsv1beta1 "github.com/GoogleCloudPlatform/k8s-config-connector/apis/refs/v1beta1"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/k8s"
	apierrors "k8s.io/apimachinery/pkg/api/errors"
	"k8s.io/apimachinery/pkg/apis/meta/v1/unstructured"
	"k8s.io/apimachinery/pkg/types"
	"sigs.k8s.io/controller-runtime/pkg/client"
)

var _ refsv1beta1.ExternalNormalizer = &KMSAutokeyConfigRef{}

// KMSAutokeyConfigRef defines the resource reference to KMSAutokeyConfig, which "External" field
// holds the GCP identifier for the KRM object.
type KMSAutokeyConfigRef struct {
	// A reference to an externally managed KMSAutokeyConfig resource.
	// Should be in the format "projects/<projectID>/locations/<location>/autokeyconfigs/<autokeyconfigID>".
	External string `json:"external,omitempty"`

	// The name of a KMSAutokeyConfig resource.
	Name string `json:"name,omitempty"`

	// The namespace of a KMSAutokeyConfig resource.
	Namespace string `json:"namespace,omitempty"`

	parent *KMSAutokeyConfigParent
}

// NormalizedExternal provision the "External" value for other resource that depends on KMSAutokeyConfig.
// If the "External" is given in the other resource's spec.KMSAutokeyConfigRef, the given value will be used.
// Otherwise, the "Name" and "Namespace" will be used to query the actual KMSAutokeyConfig object from the cluster.
func (r *KMSAutokeyConfigRef) NormalizedExternal(ctx context.Context, reader client.Reader, otherNamespace string) (string, error) {
	if r.External != "" && r.Name != "" {
		return "", fmt.Errorf("cannot specify both name and external on %s reference", KMSAutokeyConfigGVK.Kind)
	}
	// From given External
	if r.External != "" {
		if _, err := ParseKMSAutokeyConfigExternal(r.External); err != nil {
			return "", err
		}
		return r.External, nil
	}

	// From the Config Connector object
	if r.Namespace == "" {
		r.Namespace = otherNamespace
	}
	key := types.NamespacedName{Name: r.Name, Namespace: r.Namespace}
	u := &unstructured.Unstructured{}
	u.SetGroupVersionKind(KMSAutokeyConfigGVK)
	if err := reader.Get(ctx, key, u); err != nil {
		if apierrors.IsNotFound(err) {
			return "", k8s.NewReferenceNotFoundError(u.GroupVersionKind(), key)
		}
		return "", fmt.Errorf("reading referenced %s %s: %w", KMSAutokeyConfigGVK, key, err)
	}
	// Get external from status.externalRef. This is the most trustworthy place.
	actualExternalRef, _, err := unstructured.NestedString(u.Object, "status", "externalRef")
	if err != nil {
		return "", fmt.Errorf("reading status.externalRef: %w", err)
	}
	if actualExternalRef == "" {
		return "", k8s.NewReferenceNotReadyError(u.GroupVersionKind(), key)
	}
	r.External = actualExternalRef
	return r.External, nil
}

// New builds a KMSAutokeyConfigRef from the Config Connector KMSAutokeyConfig object.
func NewKMSAutokeyConfigRef(ctx context.Context, reader client.Reader, obj *KMSAutokeyConfig) (*KMSAutokeyConfigRef, error) {
	id := &KMSAutokeyConfigRef{}

	// Get Parent
	folderRef, err := refsv1beta1.ResolveFolder(ctx, reader, obj, obj.Spec.FolderRef)
	if err != nil {
		return nil, err
	}
	folderID := folderRef.FolderID
	if folderID == "" {
		return nil, fmt.Errorf("cannot resolve project")
	}
	id.parent = &KMSAutokeyConfigParent{FolderID: folderID}

	// Use approved External
	externalRef := valueOf(obj.Status.ExternalRef)
	if externalRef == "" {
		id.External = asKMSAutokeyConfigExternal(id.parent)
		return id, nil
	}

	// Validate desired with actual
	actualParent, err := ParseKMSAutokeyConfigExternal(externalRef)
	if err != nil {
		return nil, err
	}
	if actualParent.FolderID != folderID {
		return nil, fmt.Errorf("spec.folderRef changed, expect %s, got %s", actualParent.FolderID, folderID)
	}
	id.External = externalRef
	id.parent = &KMSAutokeyConfigParent{FolderID: folderID}
	return id, nil
}

func (r *KMSAutokeyConfigRef) Parent() (*KMSAutokeyConfigParent, error) {
	if r.parent != nil {
		return r.parent, nil
	}
	if r.External != "" {
		parent, err := ParseKMSAutokeyConfigExternal(r.External)
		if err != nil {
			return nil, err
		}
		return parent, nil
	}
	return nil, fmt.Errorf("KMSAutokeyConfigRef not initialized from `NewKMSAutokeyConfigRef` or `NormalizedExternal`")
}

type KMSAutokeyConfigParent struct {
	FolderID string
	Location string
}

func (p *KMSAutokeyConfigParent) String() string {
	return "folders/" + p.FolderID
}

func asKMSAutokeyConfigExternal(parent *KMSAutokeyConfigParent) (external string) {
	return parent.String() + "/autokeyConfig"
}

func ParseKMSAutokeyConfigExternal(external string) (parent *KMSAutokeyConfigParent, err error) {
	external = strings.TrimPrefix(external, "/")
	tokens := strings.Split(external, "/")
	if len(tokens) != 3 || tokens[0] != "folders" || tokens[2] != "autokeyConfig" {
		return nil, fmt.Errorf("format of KMSAutokeyConfig external=%q was not known (use projects/<projectId>/locations/<location>/autokeyconfigs/<autokeyconfigID>)", external)
	}
	parent = &KMSAutokeyConfigParent{
		FolderID: tokens[1],
	}
	return parent, nil
}

func valueOf[T any](t *T) T {
	var zeroVal T
	if t == nil {
		return zeroVal
	}
	return *t
}
