This document provides instructions on how to handle API resource references in KCC.

## Background

In KCC, a resource can reference another resource. For example, a `StorageBucket` might have a `kmsKeyName` field that references a `KMSKey` resource. When KCC sees such a reference, it needs to resolve the reference to the fully-qualified GCP resource name.

The way KCC handles this is by defining a "reference" object. This object is a struct that implements the `refsv1beta1.ExternalNormalizer` interface. This interface has a single method, `NormalizedExternal`, which takes a `client.Reader` and returns the fully-qualified GCP resource name.

The reference object is typically placed in its own file, `apis/<service>/<version>/*_reference.go`. For example, the reference object for `AlloyDBCluster` is in `apis/alloydb/v1beta1/cluster_reference.go`.

## Task

Your task is to identify fields in KCC's API that are references to other GCP resources, and change them to use the KCC reference object.

### Identifying Reference Fields

There are two cases to consider:

1.  **Greenfield resources:** These are resources that are being added to KCC for the first time. In this case, you should look at the comments for each field in the API. If a field's comment indicates that it is a reference to another GCP resource, then you should mark it as a reference by adding the `//+kcc:ref={Your guess of its KCC kind}` annotation.

2.  **Terraform/DCL-migrated resources:** These are resources that are being migrated from Terraform or DCL to KCC. In this case, you should look at the `config/crds` differences. The field may be different by ending with `*Ref`, and is a structure with fields like "name", "namespace", and/or "external", rather than a string.

### Changing a Field to a Reference

Once you have identified a reference field, you need to change its type to be a pointer to the reference object. For example, if you have a field `KmsKeyName *string`, and you have determined that it is a reference to a `KMSKey` resource, you should change it to `KmsKeyRef *KMSKeyRef`.

If the reference object does not exist, you should create one. You can use `apis/alloydb/v1beta1/cluster_reference.go` as a template.

Here is an example of a reference object:

```go
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

	refsv1beta1 "github.com/GoogleCloudPlatform/k8s-config-connector/apis/refs/v1beta1"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/k8s"
	apierrors "k8s.io/apimachinery/pkg/api/errors"
	"k8s.io/apimachinery/pkg/apis/meta/v1/unstructured"
	"k8s.io/apimachinery/pkg/types"
	"sigs.k8s.io/controller-runtime/pkg/client"
)

var _ refsv1beta1.ExternalNormalizer = &ClusterRef{}

// ClusterRef defines the resource reference to AlloyDBCluster, which "External" field
// holds the GCP identifier for the KRM object.
type ClusterRef struct {
	// A reference to an externally managed AlloyDBCluster resource.
	// Should be in the format "projects/{{projectID}}/locations/{{location}}/clusters/{{clusterID}}".
	External string `json:"external,omitempty"`

	// The name of a AlloyDBCluster resource.
	Name string `json:"name,omitempty"`

	// The namespace of a AlloyDBCluster resource.
	Namespace string `json:"namespace,omitempty"`
}

// NormalizedExternal provision the "External" value for other resource that depends on AlloyDBCluster.
// If the "External" is given in the other resource's spec.AlloyDBClusterRef, the given value will be used.
// Otherwise, the "Name" and "Namespace" will be used to query the actual AlloyDBCluster object from the cluster.
func (r *ClusterRef) NormalizedExternal(ctx context.Context, reader client.Reader, otherNamespace string) (string, error) {
	if r.External != "" && r.Name != "" {
		return "", fmt.Errorf("cannot specify both name and external on %s reference", AlloyDBClusterGVK.Kind)
	}
	// From given External
	if r.External != "" {
		if _, _, err := ParseClusterExternal(r.External); err != nil {
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
	u.SetGroupVersionKind(AlloyDBClusterGVK)
	if err := reader.Get(ctx, key, u); err != nil {
		if apierrors.IsNotFound(err) {
			return "", k8s.NewReferenceNotFoundError(u.GroupVersionKind(), key)
		}
		return "", fmt.Errorf("reading referenced %s %s: %w", AlloyDBClusterGVK, key, err)
	}
	// Get external from status.externalRef. This is the most trustworthy place.
	actualExternalRef, _, err := unstructured.NestedString(u.Object, "status", "externalRef")
	if err != nil {
		return "", fmt.Errorf("reading status.externalRef: %w", err)
	}
	if actualExternalRef == "" {
		// It's possible the referenced AlloyDBCluster is a legacy one and doesn't
		// have `status.externalRef`.
		ready, err := isResourceReady(u)
		if err != nil {
			return "", fmt.Errorf("checking if referenced %s %s is ready: %w", AlloyDBClusterGVK, key, err)
		}
		if !ready {
			return "", k8s.NewReferenceNotReadyError(u.GroupVersionKind(), key)
		}
		projectID, err := refsv1beta1.ResolveProjectID(ctx, reader, u)
		if err != nil {
			return "", err
		}
		location, err := refsv1beta1.GetLocation(u)
		if err != nil {
			return "", err
		}
		clusterID, err := refsv1beta1.GetResourceID(u)
		if err != nil {
			return "", err
		}
		return fmt.Sprintf("projects/%s/locations/%s/clusters/%s", projectID, location, clusterID), nil
	}

	r.External = actualExternalRef
	return r.External, nil
}
```
