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

	bigtablev1beta1 "github.com/GoogleCloudPlatform/k8s-config-connector/apis/bigtable/v1beta1"
	"github.com/GoogleCloudPlatform/k8s-config-connector/apis/common"
	"github.com/GoogleCloudPlatform/k8s-config-connector/apis/common/identity"
	"github.com/GoogleCloudPlatform/k8s-config-connector/apis/common/parent"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/gcpurls"
	"sigs.k8s.io/controller-runtime/pkg/client"
)

var (
	_ identity.IdentityV2 = &BigtableMemoryLayerIdentity{}
	_ identity.Resource   = &BigtableMemoryLayer{}
)

var BigtableMemoryLayerIdentityFormat = gcpurls.Template[BigtableMemoryLayerIdentity]("bigtable.googleapis.com", "projects/{project}/instances/{instance}/clusters/{cluster}/memoryLayer")

// BigtableMemoryLayerIdentity is the identity of a GCP BigtableMemoryLayer resource.
// +k8s:deepcopy-gen=false
type BigtableMemoryLayerIdentity struct {
	Project  string
	Instance string
	Cluster  string
}

func (i *BigtableMemoryLayerIdentity) String() string {
	return BigtableMemoryLayerIdentityFormat.ToString(*i)
}

func (i *BigtableMemoryLayerIdentity) FromExternal(ref string) error {
	parsed, match, err := BigtableMemoryLayerIdentityFormat.Parse(ref)
	if err != nil {
		return fmt.Errorf("format of BigtableMemoryLayer external=%q was not known (use %s): %w", ref, BigtableMemoryLayerIdentityFormat.CanonicalForm(), err)
	}
	if !match {
		return fmt.Errorf("format of BigtableMemoryLayer external=%q was not known (use %s)", ref, BigtableMemoryLayerIdentityFormat.CanonicalForm())
	}

	*i = *parsed
	return nil
}

func (i *BigtableMemoryLayerIdentity) Host() string {
	return BigtableMemoryLayerIdentityFormat.Host()
}

func (i *BigtableMemoryLayerIdentity) ParentString() string {
	return fmt.Sprintf("projects/%s/instances/%s/clusters/%s", i.Project, i.Instance, i.Cluster)
}

func (i *BigtableMemoryLayerIdentity) ClusterID() string {
	return i.Cluster
}

func (i *BigtableMemoryLayerIdentity) InstanceID() string {
	return i.Instance
}

func (i *BigtableMemoryLayerIdentity) ProjectID() string {
	return i.Project
}

func getIdentityFromBigtableMemoryLayerSpec(ctx context.Context, reader client.Reader, obj *BigtableMemoryLayer) (*BigtableMemoryLayerIdentity, error) {
	// Resolve clusterRef
	clusterRef := obj.Spec.ClusterRef
	clusterExternal, err := clusterRef.NormalizedExternal(ctx, reader, obj.GetNamespace())
	if err != nil {
		return nil, err
	}
	instanceParent, clusterID, err := ParseClusterExternal(clusterExternal)
	if err != nil {
		return nil, err
	}

	return &BigtableMemoryLayerIdentity{
		Project:  instanceParent.Parent.ProjectID,
		Instance: instanceParent.Id,
		Cluster:  clusterID,
	}, nil
}

// GetIdentity extracts the identity of the BigtableMemoryLayer resource.
func (obj *BigtableMemoryLayer) GetIdentity(ctx context.Context, reader client.Reader) (identity.Identity, error) {
	specIdentity, err := getIdentityFromBigtableMemoryLayerSpec(ctx, reader, obj)
	if err != nil {
		return nil, err
	}

	externalRef := common.ValueOf(obj.Status.ExternalRef)
	if externalRef != "" {
		// Validate desired with actual
		statusIdentity := &BigtableMemoryLayerIdentity{}
		if err := statusIdentity.FromExternal(externalRef); err != nil {
			return nil, err
		}

		if statusIdentity.Project != specIdentity.Project {
			return nil, fmt.Errorf("spec.clusterRef ProjectID changed, expect %s, got %s", statusIdentity.Project, specIdentity.Project)
		}
		if statusIdentity.Instance != specIdentity.Instance {
			return nil, fmt.Errorf("spec.clusterRef InstanceID changed, expect %s, got %s", statusIdentity.Instance, specIdentity.Instance)
		}
		if statusIdentity.Cluster != specIdentity.Cluster {
			return nil, fmt.Errorf("spec.clusterRef ClusterID changed, expect %s, got %s", statusIdentity.Cluster, specIdentity.Cluster)
		}
	}

	return specIdentity, nil
}

// NewBigtableMemoryLayerIdentity builds a BigtableMemoryLayerIdentity from the Config Connector BigtableMemoryLayer object.
func NewBigtableMemoryLayerIdentity(ctx context.Context, reader client.Reader, obj *BigtableMemoryLayer) (*BigtableMemoryLayerIdentity, error) {
	id, err := obj.GetIdentity(ctx, reader)
	if err != nil {
		return nil, err
	}
	return id.(*BigtableMemoryLayerIdentity), nil
}

func ParseBigtableMemoryLayerExternal(external string) (*ClusterIdentity, error) {
	id := &BigtableMemoryLayerIdentity{}
	if err := id.FromExternal(external); err != nil {
		return nil, err
	}
	p := &ClusterIdentity{
		parent: &bigtablev1beta1.InstanceIdentity{
			Parent: &parent.ProjectParent{
				ProjectID: id.Project,
			},
			Id: id.Instance,
		},
		id: id.Cluster,
	}
	return p, nil
}
