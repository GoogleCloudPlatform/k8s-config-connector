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
	"strings"

	"github.com/GoogleCloudPlatform/k8s-config-connector/apis/common"
	"github.com/GoogleCloudPlatform/k8s-config-connector/apis/common/identity"
	"github.com/GoogleCloudPlatform/k8s-config-connector/apis/common/parent"
	refsv1beta1 "github.com/GoogleCloudPlatform/k8s-config-connector/apis/refs/v1beta1"
	"sigs.k8s.io/controller-runtime/pkg/client"
)

var _ identity.Identity = &ClusterIdentity{}

// ClusterIdentity is the identity of a HypercomputeClusterCluster.
type ClusterIdentity struct {
	parent *parent.ProjectAndLocationParent
	id     string
}

func (i *ClusterIdentity) String() string {
	return i.parent.String() + "/clusters/" + i.id
}

func (i *ClusterIdentity) ID() string {
	return i.id
}

func (i *ClusterIdentity) Parent() *parent.ProjectAndLocationParent {
	return i.parent
}

func (i *ClusterIdentity) FromExternal(ref string) error {
	tokens := strings.Split(ref, "/clusters/")
	if len(tokens) != 2 {
		return fmt.Errorf("format of clusters external=%q was not known (use projects/{{projectID}}/locations/{{location}}/clusters/{{clusterID}})", ref)
	}
	i.parent = &parent.ProjectAndLocationParent{}
	if err := i.parent.FromExternal(tokens[0]); err != nil {
		return err
	}
	i.id = tokens[1]
	if i.id == "" {
		return fmt.Errorf("clusterID was empty in external=%q", ref)
	}
	return nil
}

var _ identity.Resource = &HypercomputeClusterCluster{}

func (obj *HypercomputeClusterCluster) GetIdentity(ctx context.Context, reader client.Reader) (identity.Identity, error) {
	cluster := &ClusterIdentity{}
	cluster.parent = &parent.ProjectAndLocationParent{}

	if obj.Spec.ProjectRef == nil {
		return nil, fmt.Errorf("spec.projectRef cannot be nil")
	}

	// Resolve user-configured Parent
	project, err := refsv1beta1.ResolveProject(ctx, reader, obj.GetNamespace(), obj.Spec.ProjectRef)
	if err != nil {
		return nil, err
	}
	cluster.parent.ProjectID = project.ProjectID
	cluster.parent.Location = obj.Spec.Location

	// Get user-configured ID
	cluster.id = common.ValueOf(obj.Spec.ResourceID)
	if cluster.id == "" {
		cluster.id = obj.GetName()
	}
	if cluster.id == "" {
		return nil, fmt.Errorf("cannot resolve resource ID")
	}

	// Validate against the ID stored in status.externalRef, if any
	externalRef := common.ValueOf(obj.Status.ExternalRef)
	if externalRef != "" {
		statusIdentity := &ClusterIdentity{}
		if err := statusIdentity.FromExternal(externalRef); err != nil {
			return nil, fmt.Errorf("cannot parse existing externalRef=%q: %w", externalRef, err)
		}
		if statusIdentity.String() != cluster.String() {
			return nil, fmt.Errorf("existing externalRef=%q does not match the identity resolved from spec: %q", externalRef, cluster.String())
		}
	}
	return cluster, nil
}
