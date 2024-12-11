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

	apierrors "k8s.io/apimachinery/pkg/api/errors"
	"k8s.io/apimachinery/pkg/apis/meta/v1/unstructured"
	"k8s.io/apimachinery/pkg/runtime/schema"
	"k8s.io/apimachinery/pkg/types"
	"sigs.k8s.io/controller-runtime/pkg/client"
)

type AlloyDBClusterRef struct {
	// If provided must be in the format `projects/[projectId]/locations/[location]/clusters/[clusterId]`.
	External string `json:"external,omitempty"`
	// The `metadata.name` field of a `AlloyDBCluster` resource.
	Name string `json:"name,omitempty"`
	// The `metadata.namespace` field of a `AlloyDBCluster` resource.
	Namespace string `json:"namespace,omitempty"`
}

type AlloyDBCluster struct {
	projectID string
	location  string
	clusterID string
}

func ResolveAlloyDBCluster(ctx context.Context, reader client.Reader, src client.Object, ref *AlloyDBClusterRef) (*AlloyDBCluster, error) {
	if ref == nil {
		return nil, nil
	}

	if ref.Name == "" && ref.External == "" {
		return nil, fmt.Errorf("must specify either name or external on AlloyDBClusterRef")
	}
	if ref.Name != "" && ref.External != "" {
		return nil, fmt.Errorf("cannot specify both name and external on AlloyDBClusterRef")
	}

	// External is provided.
	if ref.External != "" {
		// External should be in the `projects/[projectId]/locations/[location]/clusters/[clusterId]` format.
		tokens := strings.Split(ref.External, "/")
		if len(tokens) == 6 && tokens[0] == "projects" && tokens[2] == "locations" && tokens[4] == "clusters" {
			return &AlloyDBCluster{
				projectID: tokens[1],
				location:  tokens[3],
				clusterID: tokens[5],
			}, nil
		}
		return nil, fmt.Errorf("format of AlloyDBClusterRef external=%q was not known (use projects/[projectId]/locations/[location]/clusters/[clusterId])", ref.External)

	}

	// Fetch AlloyDBCluster object to construct the external form.
	cluster := &unstructured.Unstructured{}
	cluster.SetGroupVersionKind(schema.GroupVersionKind{
		Group:   "alloydb.cnrm.cloud.google.com",
		Version: "v1beta1",
		Kind:    "AlloyDBCluster",
	})
	nn := types.NamespacedName{
		Namespace: ref.Namespace,
		Name:      ref.Name,
	}
	if nn.Namespace == "" {
		nn.Namespace = src.GetNamespace()
	}
	if err := reader.Get(ctx, nn, cluster); err != nil {
		if apierrors.IsNotFound(err) {
			return nil, fmt.Errorf("referenced AlloyDBCluster %v not found", nn)
		}
		return nil, fmt.Errorf("error reading referenced AlloyDBCluster %v: %w", nn, err)
	}
	projectID, err := ResolveProjectID(ctx, reader, cluster)
	if err != nil {
		return nil, err
	}
	location, err := GetLocation(cluster)
	if err != nil {
		return nil, err
	}
	clusterID, err := GetResourceID(cluster)
	if err != nil {
		return nil, err
	}
	return &AlloyDBCluster{
		projectID: projectID,
		location:  location,
		clusterID: clusterID,
	}, nil
}

type AlloyDBClusterTypeRef struct {
	// The type of instance. Possible values: ["PRIMARY", "READ_POOL", "SECONDARY"]
	External string `json:"external,omitempty"`
	// The `metadata.name` field of a `AlloyDBCluster` resource.
	Name string `json:"name,omitempty"`
	// The `metadata.namespace` field of a `AlloyDBCluster` resource.
	Namespace string `json:"namespace,omitempty"`
}

func ResolveAlloyDBClusterType(ctx context.Context, reader client.Reader, src client.Object, ref *AlloyDBClusterTypeRef) (*string, error) {
	if ref == nil {
		return nil, nil
	}

	if ref.Name == "" && ref.External == "" {
		return nil, fmt.Errorf("must specify either name or external on AlloyDBClusterRef")
	}
	if ref.Name != "" && ref.External != "" {
		return nil, fmt.Errorf("cannot specify both name and external on AlloyDBClusterRef")
	}

	// External is provided.
	if ref.External != "" {
		return lazyPtr(ref.External), nil
	}

	// Fetch AlloyDBCluster object to construct the external form.
	cluster := &unstructured.Unstructured{}
	cluster.SetGroupVersionKind(schema.GroupVersionKind{
		Group:   "alloydb.cnrm.cloud.google.com",
		Version: "v1beta1",
		Kind:    "AlloyDBCluster",
	})
	nn := types.NamespacedName{
		Namespace: ref.Namespace,
		Name:      ref.Name,
	}
	if nn.Namespace == "" {
		nn.Namespace = src.GetNamespace()
	}
	if err := reader.Get(ctx, nn, cluster); err != nil {
		if apierrors.IsNotFound(err) {
			return nil, fmt.Errorf("referenced AlloyDBCluster %v not found", nn)
		}
		return nil, fmt.Errorf("error reading referenced AlloyDBCluster %v: %w", nn, err)
	}
	clusterType, _, err := unstructured.NestedString(cluster.Object, "spec", "clusterType")
	if err != nil {
		return nil, fmt.Errorf("reading spec.clusterType from %v %v/%v: %w", cluster.GroupVersionKind().Kind, cluster.GetNamespace(), cluster.GetName(), err)
	}
	if clusterType == "" {
		// clusterType is defaulted to "PRIMARY" when not set.
		clusterType = "PRIMARY"
	}
	return lazyPtr(clusterType), nil
}

func ResolveAlloyDBClusterName(ctx context.Context, reader client.Reader, obj *unstructured.Unstructured) (string, error) {
	clusterRefExternal, _, _ := unstructured.NestedString(obj.Object, "spec", "clusterRef", "external")
	if clusterRefExternal != "" {
		clusterRef := AlloyDBClusterRef{
			External: clusterRefExternal,
		}

		cluster, err := ResolveAlloyDBCluster(ctx, reader, obj, &clusterRef)
		if err != nil {
			return "", fmt.Errorf("cannot parse clusterRef.external %q in %v %v/%v: %w", clusterRefExternal, obj.GetKind(), obj.GetNamespace(), obj.GetName(), err)
		}
		return cluster.String(), nil
	}

	clusterRefName, _, _ := unstructured.NestedString(obj.Object, "spec", "clusterRef", "name")
	if clusterRefName != "" {
		clusterRefNamespace, _, _ := unstructured.NestedString(obj.Object, "spec", "clusterRef", "namespace")

		clusterRef := AlloyDBClusterRef{
			Name:      clusterRefName,
			Namespace: clusterRefNamespace,
		}
		if clusterRef.Namespace == "" {
			clusterRef.Namespace = obj.GetNamespace()
		}

		cluster, err := ResolveAlloyDBCluster(ctx, reader, obj, &clusterRef)
		if err != nil {
			return "", fmt.Errorf("cannot parse clusterRef in %v %v/%v: %w", obj.GetKind(), obj.GetNamespace(), obj.GetName(), err)
		}
		return cluster.String(), nil
	}

	return "", fmt.Errorf("cannot find AlloyDB cluster name for %v %v/%v", obj.GetKind(), obj.GetNamespace(), obj.GetName())
}

func (c *AlloyDBCluster) String() string {
	return fmt.Sprintf("projects/%s/locations/%s/clusters/%s", c.projectID, c.location, c.clusterID)
}

type AlloyDBInstanceRef struct {
	// If provided must be in the format `projects/[projectId]/locations/[location]/clusters/[clusterId]/instances/[instanceId]`.
	External string `json:"external,omitempty"`
	// The `metadata.name` field of a `AlloyDBInstance` resource.
	Name string `json:"name,omitempty"`
	// The `metadata.namespace` field of a `AlloyDBInstance` resource.
	Namespace string `json:"namespace,omitempty"`
}

type AlloyDBInstance struct {
	clusterName string
	instanceID  string
}

func ResolveAlloyDBInstance(ctx context.Context, reader client.Reader, src client.Object, ref *AlloyDBInstanceRef) (*AlloyDBInstance, error) {
	if ref == nil {
		return nil, nil
	}

	if ref.Name == "" && ref.External == "" {
		return nil, fmt.Errorf("must specify either name or external on AlloyDBInstanceRef")
	}
	if ref.Name != "" && ref.External != "" {
		return nil, fmt.Errorf("cannot specify both name and external on AlloyDBInstanceRef")
	}

	// External is provided.
	if ref.External != "" {
		// External should be in the `projects/[projectId]/locations/[location]/clusters/[clusterId]/instances/[instanceId]` format.
		tokens := strings.Split(ref.External, "/")
		if len(tokens) == 6 && tokens[0] == "projects" && tokens[2] == "locations" && tokens[4] == "clusters" && tokens[6] == "instances" {
			return &AlloyDBInstance{
				clusterName: fmt.Sprintf("%s/%s/%s/%s/%s/%s", tokens[0], tokens[1], tokens[2], tokens[3], tokens[4], tokens[5]),
				instanceID:  tokens[7],
			}, nil
		}
		return nil, fmt.Errorf("format of AlloyDBInstanceRef external=%q was not known (use projects/[projectId]/locations/[location]/clusters/[clusterId]/instances/[instanceId])", ref.External)

	}

	// Fetch AlloyDBInstance object to construct the external form.
	instance := &unstructured.Unstructured{}
	instance.SetGroupVersionKind(schema.GroupVersionKind{
		Group:   "alloydb.cnrm.cloud.google.com",
		Version: "v1beta1",
		Kind:    "AlloyDBInstance",
	})
	nn := types.NamespacedName{
		Namespace: ref.Namespace,
		Name:      ref.Name,
	}
	if nn.Namespace == "" {
		nn.Namespace = src.GetNamespace()
	}
	if err := reader.Get(ctx, nn, instance); err != nil {
		if apierrors.IsNotFound(err) {
			return nil, fmt.Errorf("referenced AlloyDBInstance %v not found", nn)
		}
		return nil, fmt.Errorf("error reading referenced AlloyDBInstance %v: %w", nn, err)
	}
	clusterName, err := ResolveAlloyDBClusterName(ctx, reader, instance)
	if err != nil {
		return nil, err
	}
	instanceID, err := GetResourceID(instance)
	if err != nil {
		return nil, err
	}
	return &AlloyDBInstance{
		clusterName: clusterName,
		instanceID:  instanceID,
	}, nil
}

func (i *AlloyDBInstance) String() string {
	return fmt.Sprintf("%s/instances/%s", i.clusterName, i.instanceID)
}
