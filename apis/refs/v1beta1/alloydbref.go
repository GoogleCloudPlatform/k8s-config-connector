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

	"github.com/GoogleCloudPlatform/k8s-config-connector/apis/common"

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
	ProjectID string
	Location  string
	ClusterID string
}

// TODO: Remove after AlloyDBCluster is migrated to direct controller.
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
				ProjectID: tokens[1],
				Location:  tokens[3],
				ClusterID: tokens[5],
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
		ProjectID: projectID,
		Location:  location,
		ClusterID: clusterID,
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
		return common.LazyPtr(ref.External), nil
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
	return common.LazyPtr(clusterType), nil
}

func (c *AlloyDBCluster) String() string {
	return fmt.Sprintf("projects/%s/locations/%s/clusters/%s", c.ProjectID, c.Location, c.ClusterID)
}

type AlloyDBBackupRef struct {
	// If provided must be in the format `projects/[projectId]/locations/[location]/backups/[backupId]`.
	External string `json:"external,omitempty"`
	// The `metadata.name` field of a `AlloyDBBackup` resource.
	Name string `json:"name,omitempty"`
	// The `metadata.namespace` field of a `AlloyDBBackup` resource.
	Namespace string `json:"namespace,omitempty"`
}

func ResolveAlloyDBBackupRef(ctx context.Context, reader client.Reader, src client.Object, ref *AlloyDBBackupRef) (*AlloyDBBackupRef, error) {
	if ref == nil {
		return nil, nil
	}

	if ref.Name == "" && ref.External == "" {
		return nil, fmt.Errorf("must specify either name or external on AlloyDBBackupRef")
	}
	if ref.Name != "" && ref.External != "" {
		return nil, fmt.Errorf("cannot specify both name and external on AlloyDBBackupRef")
	}

	// External should be in the `projects/[projectId]/locations/[location]/backups/[backupId]` format
	if ref.External != "" {
		tokens := strings.Split(ref.External, "/")
		if len(tokens) == 6 && tokens[0] == "projects" && tokens[2] == "locations" && tokens[4] == "backups" {
			ref = &AlloyDBBackupRef{
				External: fmt.Sprintf("projects/%s/locations/%s/backups/%s", tokens[1], tokens[3], tokens[5]),
			}
			return ref, nil
		}
		return nil, fmt.Errorf("format of AlloyDBBackupRef external=%q was not known (use projects/[projectId]/locations/[location]/backups/[backupId])", ref.External)
	}

	backupNN := types.NamespacedName{
		Namespace: ref.Namespace,
		Name:      ref.Name,
	}
	if backupNN.Namespace == "" {
		backupNN.Namespace = src.GetNamespace()
	}

	// Fetch object from k8s cluster to construct the external form
	backup := &unstructured.Unstructured{}
	backup.SetGroupVersionKind(schema.GroupVersionKind{
		Group:   "alloydb.cnrm.cloud.google.com",
		Version: "v1beta1",
		Kind:    "AlloyDBBackup",
	})
	if err := reader.Get(ctx, backupNN, backup); err != nil {
		if apierrors.IsNotFound(err) {
			return nil, fmt.Errorf("referenced AlloyDBBackup %v not found", backupNN)
		}
		return nil, fmt.Errorf("error reading referenced AlloyDBBackup %v: %w", backupNN, err)
	}

	resourceID, err := GetResourceID(backup)
	if err != nil {
		return nil, err
	}

	projectID, err := ResolveProjectID(ctx, reader, backup)
	if err != nil {
		return nil, err
	}

	location, err := GetLocation(backup)
	if err != nil {
		return nil, err
	}

	ref = &AlloyDBBackupRef{
		External: fmt.Sprintf("projects/%s/locations/%s/backups/%s", projectID, location, resourceID),
	}
	return ref, nil
}
