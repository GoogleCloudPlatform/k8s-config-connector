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

	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/k8s"
	apierrors "k8s.io/apimachinery/pkg/api/errors"
	"k8s.io/apimachinery/pkg/apis/meta/v1/unstructured"
	"k8s.io/apimachinery/pkg/runtime/schema"
	"k8s.io/apimachinery/pkg/types"
	"sigs.k8s.io/controller-runtime/pkg/client"
)

type DataprocClusterRef struct {
	/* The self-link of an existing Dataproc Cluster to act as a Spark History Server for the connection , when not managed by Config Connector. */
	External string `json:"external,omitempty"`
	/* The `name` field of a Dataproc Cluster. */
	Name string `json:"name,omitempty"`
	/* The `namespace` field of a Dataproc Cluster. */
	Namespace string `json:"namespace,omitempty"`
}

type DataprocCluster struct {
	ProjectID string
	Region    string
	ClusterID string
}

func (s *DataprocCluster) String() string {
	return "projects/" + s.ProjectID + "/regions/" + s.Region + "/clusters/" + s.ClusterID
}

func ResolveDataprocClusterRef(ctx context.Context, reader client.Reader, obj client.Object, ref *DataprocClusterRef) (*DataprocCluster, error) {
	if ref == nil {
		return nil, nil
	}

	if ref.Name == "" && ref.External == "" {
		return nil, fmt.Errorf("must specify either name or external on dataprocClusterRef")
	}
	if ref.External != "" && ref.Name != "" {
		return nil, fmt.Errorf("cannot specify both spec.dataprocClusterRef.name and spec.dataprocClusterRef.external")
	}

	if ref.External != "" {
		// External must be in form `projects/<projectID>/regions/<region>/clusters/<clusterName>`.
		// see https://cloud.google.com/dataproc/docs/reference/rest/v1/projects.regions.clusters/create
		tokens := strings.Split(ref.External, "/")
		if len(tokens) == 6 && tokens[0] == "projects" && tokens[2] == "regions" && tokens[4] == "clusters" {
			return &DataprocCluster{
				ProjectID: tokens[1],
				Region:    tokens[3],
				ClusterID: tokens[5],
			}, nil
		}
		return nil, fmt.Errorf("format of DataprocCluster external=%q was not known (use projects/<projectID>/regions/<region>/clusters/<clusterName>)", ref.External)
	}

	key := types.NamespacedName{
		Namespace: ref.Namespace,
		Name:      ref.Name,
	}
	if key.Namespace == "" {
		key.Namespace = obj.GetNamespace()
	}

	cluster := &unstructured.Unstructured{}
	cluster.SetGroupVersionKind(schema.GroupVersionKind{
		Group:   "dataproc.cnrm.cloud.google.com",
		Version: "v1beta1",
		Kind:    "DataprocCluster",
	})
	if err := reader.Get(ctx, key, cluster); err != nil {
		if apierrors.IsNotFound(err) {
			return nil, fmt.Errorf("referenced DataprocCluster %v not found", key)
		}
		return nil, fmt.Errorf("error reading referenced DataprocCluster %v: %w", key, err)
	}
	resource, err := k8s.NewResource(cluster)
	if err != nil {
		return nil, fmt.Errorf("error converting unstructured to resource: %w", err)
	}
	if !k8s.IsResourceReady(resource) {
		return nil, k8s.NewReferenceNotReadyError(cluster.GroupVersionKind(), key)
	}

	resourceID, _, err := unstructured.NestedString(cluster.Object, "spec", "resourceID")
	if err != nil {
		return nil, fmt.Errorf("reading spec.resourceID from DataprocCluster %s/%s: %w", cluster.GetNamespace(), cluster.GetName(), err)
	}
	if resourceID == "" {
		resourceID = cluster.GetName()
	}

	projectID, err := ResolveProjectID(ctx, reader, cluster)
	if err != nil {
		return nil, err
	}

	region, _, err := unstructured.NestedString(cluster.Object, "spec", "region")
	if err != nil {
		return nil, fmt.Errorf("reading spec.region from DataprocCluster %s/%s: %w", cluster.GetNamespace(), cluster.GetName(), err)
	}

	return &DataprocCluster{
		ProjectID: projectID,
		Region:    region,
		ClusterID: resourceID,
	}, nil
}
