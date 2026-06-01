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

package v1alpha1

import (
	"context"
	"fmt"

	"github.com/GoogleCloudPlatform/k8s-config-connector/apis/common"
	"github.com/GoogleCloudPlatform/k8s-config-connector/apis/common/identity"
	"github.com/GoogleCloudPlatform/k8s-config-connector/apis/refs/v1beta1"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/gcpurls"
	"k8s.io/apimachinery/pkg/apis/meta/v1/unstructured"
	"k8s.io/apimachinery/pkg/runtime"
	"sigs.k8s.io/controller-runtime/pkg/client"
)

var _ identity.IdentityV2 = &NodeGroupIdentity{}
var _ identity.Resource = &DataprocNodeGroup{}

func NewNodeGroupIdentity(ctx context.Context, reader client.Reader, obj *DataprocNodeGroup) (*NodeGroupIdentity, error) {
	id, err := getNodeGroupIdentityFromSpec(ctx, reader, obj)
	if err != nil {
		return nil, err
	}

	externalRef := common.ValueOf(obj.Status.ExternalRef)
	if externalRef != "" {
		actual := &NodeGroupIdentity{}
		if err := actual.FromExternal(externalRef); err != nil {
			return nil, err
		}
		if actual.Project != id.Project {
			return nil, fmt.Errorf("spec.projectRef changed, expect %s, got %s", actual.Project, id.Project)
		}
		if actual.Region != id.Region {
			return nil, fmt.Errorf("spec.location changed, expect %s, got %s", actual.Region, id.Region)
		}
		if actual.Cluster != id.Cluster {
			return nil, fmt.Errorf("spec.clusterRef changed, expect %s, got %s", actual.Cluster, id.Cluster)
		}
		if actual.NodeGroup != id.NodeGroup {
			return nil, fmt.Errorf("cannot reset `metadata.name` or `spec.resourceID` to %s, since it has already assigned to %s", id.NodeGroup, actual.NodeGroup)
		}
	}

	return id, nil
}

func (s *DataprocNodeGroup) GetIdentity(ctx context.Context, reader client.Reader) (identity.Identity, error) {
	return NewNodeGroupIdentity(ctx, reader, s)
}

// NodeGroupIdentity is the identity of a DataprocNodeGroup.
// +k8s:deepcopy-gen=false
type NodeGroupIdentity struct {
	Project   string
	Region    string
	Cluster   string
	NodeGroup string
}

var NodeGroupIdentityFormat = gcpurls.Template[NodeGroupIdentity]("dataproc.googleapis.com", "projects/{project}/regions/{region}/clusters/{cluster}/nodeGroups/{nodeGroup}")

func (i *NodeGroupIdentity) String() string {
	return NodeGroupIdentityFormat.ToString(*i)
}

func (i *NodeGroupIdentity) FromExternal(external string) error {
	parsed, match, err := NodeGroupIdentityFormat.Parse(external)
	if err != nil {
		return fmt.Errorf("format of DataprocNodeGroup external=%q was not known (use %s): %w", external, NodeGroupIdentityFormat.CanonicalForm(), err)
	}
	if !match {
		return fmt.Errorf("format of DataprocNodeGroup external=%q was not known (use %s)", external, NodeGroupIdentityFormat.CanonicalForm())
	}

	*i = *parsed
	return nil
}

func (i *NodeGroupIdentity) Host() string {
	return NodeGroupIdentityFormat.Host()
}

func getNodeGroupIdentityFromSpec(ctx context.Context, reader client.Reader, obj client.Object) (*NodeGroupIdentity, error) {
	u, ok := obj.(*unstructured.Unstructured)
	if !ok {
		m, err := runtime.DefaultUnstructuredConverter.ToUnstructured(obj)
		if err != nil {
			return nil, err
		}
		u = &unstructured.Unstructured{Object: m}
	}

	projectID, err := v1beta1.ResolveProjectID(ctx, reader, u)
	if err != nil {
		return nil, err
	}

	region, _, err := unstructured.NestedString(u.Object, "spec", "location")
	if err != nil {
		return nil, fmt.Errorf("reading spec.location: %w", err)
	}
	if region == "" {
		return nil, fmt.Errorf("location must be specified")
	}

	clusterRefMap, found, err := unstructured.NestedStringMap(u.Object, "spec", "clusterRef")
	if err != nil {
		return nil, fmt.Errorf("reading spec.clusterRef: %w", err)
	}
	if !found {
		return nil, fmt.Errorf("spec.clusterRef is required")
	}

	clusterRef := &v1beta1.DataprocClusterRef{
		External:  clusterRefMap["external"],
		Name:      clusterRefMap["name"],
		Namespace: clusterRefMap["namespace"],
	}
	clusterExternal, err := clusterRef.NormalizedExternal(ctx, reader, u.GetNamespace())
	if err != nil {
		return nil, fmt.Errorf("normalizing clusterRef: %w", err)
	}

	_, _, cluster, err := v1beta1.ParseDataprocClusterExternal(clusterExternal)
	if err != nil {
		return nil, err
	}

	resourceID, err := v1beta1.GetResourceID(u)
	if err != nil {
		return nil, err
	}
	if resourceID == "" {
		return nil, fmt.Errorf("cannot resolve resource ID")
	}

	return &NodeGroupIdentity{
		Project:   projectID,
		Region:    region,
		Cluster:   cluster,
		NodeGroup: resourceID,
	}, nil
}
