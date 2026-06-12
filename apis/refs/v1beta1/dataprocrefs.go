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
	"fmt"
	"strings"

	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/k8s"
	apierrors "k8s.io/apimachinery/pkg/api/errors"
	"k8s.io/apimachinery/pkg/apis/meta/v1/unstructured"
	"k8s.io/apimachinery/pkg/runtime/schema"
	"k8s.io/apimachinery/pkg/types"
	"sigs.k8s.io/controller-runtime/pkg/client"
)

var DataprocClusterGVK = schema.GroupVersionKind{
	Group:   "dataproc.cnrm.cloud.google.com",
	Version: "v1beta1",
	Kind:    "DataprocCluster",
}

type DataprocClusterRef struct {
	// A reference to an externally managed DataprocCluster resource.
	// Should be in the format "projects/{{projectID}}/regions/{{region}}/clusters/{{clusterName}}".
	External string `json:"external,omitempty"`

	// The name of a DataprocCluster resource.
	Name string `json:"name,omitempty"`

	// The namespace of a DataprocCluster resource.
	Namespace string `json:"namespace,omitempty"`
}

// NormalizedExternal returns the external identity of the referenced DataprocCluster.
func (r *DataprocClusterRef) NormalizedExternal(ctx context.Context, reader client.Reader, otherNamespace string) (string, error) {
	if r.External != "" && r.Name != "" {
		return "", fmt.Errorf("cannot specify both name and external on DataprocCluster reference")
	}
	// From given External
	if r.External != "" {
		return r.External, nil
	}

	// From the Config Connector object
	ns := r.Namespace
	if ns == "" {
		ns = otherNamespace
	}
	key := types.NamespacedName{Name: r.Name, Namespace: ns}
	u := &unstructured.Unstructured{}
	u.SetGroupVersionKind(DataprocClusterGVK)
	if err := reader.Get(ctx, key, u); err != nil {
		if apierrors.IsNotFound(err) {
			return "", k8s.NewReferenceNotFoundError(u.GroupVersionKind(), key)
		}
		return "", fmt.Errorf("reading referenced DataprocCluster %s: %w", key, err)
	}

	// Get external from status.externalRef. This is the most trustworthy place.
	actualExternalRef, _, err := unstructured.NestedString(u.Object, "status", "externalRef")
	if err != nil {
		return "", fmt.Errorf("reading status.externalRef: %w", err)
	}
	if actualExternalRef != "" {
		return actualExternalRef, nil
	}

	// Fallback to building external from parts (this should ideally be handled by IdentityV2 once DataprocCluster is migrated).
	projectID, err := ResolveProjectID(ctx, reader, u)
	if err != nil {
		return "", err
	}
	region, _, err := unstructured.NestedString(u.Object, "spec", "region")
	if err != nil || region == "" {
		// Fallback for location if region isn't set, though it should be.
		region, _, _ = unstructured.NestedString(u.Object, "spec", "location")
	}

	resourceID, err := GetResourceID(u)
	if err != nil {
		return "", err
	}

	return fmt.Sprintf("projects/%s/regions/%s/clusters/%s", projectID, region, resourceID), nil
}

func ParseDataprocClusterExternal(external string) (project, region, cluster string, err error) {
	tokens := strings.Split(external, "/")
	if len(tokens) != 6 || tokens[0] != "projects" || tokens[2] != "regions" || tokens[4] != "clusters" {
		return "", "", "", fmt.Errorf("format of DataprocCluster external=%q was not known (use projects/{{projectID}}/regions/{{region}}/clusters/{{clusterName}})", external)
	}
	return tokens[1], tokens[3], tokens[5], nil
}

func (r *DataprocClusterRef) GetGVK() schema.GroupVersionKind {
	return DataprocClusterGVK
}

func (r *DataprocClusterRef) GetNamespacedName() types.NamespacedName {
	return types.NamespacedName{
		Name:      r.Name,
		Namespace: r.Namespace,
	}
}

func (r *DataprocClusterRef) GetExternal() string {
	return r.External
}

func (r *DataprocClusterRef) SetExternal(ref string) {
	r.External = ref
	r.Name = ""
	r.Namespace = ""
}

func (r *DataprocClusterRef) ValidateExternal(ref string) error {
	_, _, _, err := ParseDataprocClusterExternal(ref)
	return err
}

func (r *DataprocClusterRef) Normalize(ctx context.Context, reader client.Reader, defaultNamespace string) error {
	external, err := r.NormalizedExternal(ctx, reader, defaultNamespace)
	if err != nil {
		return err
	}
	r.SetExternal(external)
	return nil
}
