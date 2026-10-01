// Copyright 2026 Google LLC
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

package v1beta1

import (
	"context"
	"fmt"

	"github.com/GoogleCloudPlatform/k8s-config-connector/apis/common/identity"
	refs "github.com/GoogleCloudPlatform/k8s-config-connector/apis/refs/v1beta1"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/gcpurls"
	"sigs.k8s.io/controller-runtime/pkg/client"
)

var (
	_ identity.IdentityV2 = &ContainerNodePoolIdentity{}
	_ identity.Resource   = &ContainerNodePool{}
)

var RegionalContainerNodePoolIdentityFormat = gcpurls.Template[ContainerNodePoolIdentity]("container.googleapis.com", "projects/{project}/locations/{location}/clusters/{cluster}/nodePools/{nodePool}")
var ZonalContainerNodePoolIdentityFormat = gcpurls.Template[ContainerNodePoolIdentity]("container.googleapis.com", "projects/{project}/zones/{location}/clusters/{cluster}/nodePools/{nodePool}")

// +k8s:deepcopy-gen=false
// ContainerNodePoolIdentity is the identity of a GCP ContainerNodePool resource.
//
// Note on migration diff: In the legacy (Terraform) controller, the resource paths for
// Create and Update were inconsistent: Create used the legacy zonal path (.../zones/{zone}/clusters/{cluster}/nodePools)
// and populated status.externalRef with "zones/{zone}", while Update used the canonical
// location-based path (.../locations/{location}/clusters/{cluster}/nodePools/{nodepool}).
// In the direct controller, we make the identity and paths consistent across the entire
// resource lifecycle by always using the canonical "locations/{location}" format for both
// regional and zonal node pools. This intentional alignment causes a diff in status.externalRef
// and HTTP traffic logs when migrating from the legacy controller.
type ContainerNodePoolIdentity struct {
	Project  string
	Location string
	Cluster  string
	NodePool string
}

// String returns the canonical GCP resource name in the format
// "projects/{project}/locations/{location}/clusters/{cluster}/nodePools/{nodePool}".
func (i *ContainerNodePoolIdentity) String() string {
	return RegionalContainerNodePoolIdentityFormat.ToString(*i)
}

func (i *ContainerNodePoolIdentity) FromExternal(ref string) error {
	ref = identity.StripReferencePrefixes(ref, "container.googleapis.com")
	if parsed, match, _ := RegionalContainerNodePoolIdentityFormat.Parse(ref); match {
		*i = *parsed
		return nil
	}
	if parsed, match, _ := ZonalContainerNodePoolIdentityFormat.Parse(ref); match {
		*i = *parsed
		return nil
	}
	return fmt.Errorf("format of ContainerNodePool external=%q was not known (use %s or %s)", ref, RegionalContainerNodePoolIdentityFormat.CanonicalForm(), ZonalContainerNodePoolIdentityFormat.CanonicalForm())
}

func (i *ContainerNodePoolIdentity) Host() string {
	return RegionalContainerNodePoolIdentityFormat.Host()
}

// ParentString returns the parent ContainerCluster GCP URI.
func (i *ContainerNodePoolIdentity) ParentString() string {
	return fmt.Sprintf("projects/%s/locations/%s/clusters/%s", i.Project, i.Location, i.Cluster)
}

// LocationValue returns the location (zone or region) of the node pool.
func (i *ContainerNodePoolIdentity) LocationValue() string {
	return i.Location
}

func getIdentityFromContainerNodePoolSpec(ctx context.Context, reader client.Reader, obj *ContainerNodePool) (*ContainerNodePoolIdentity, error) {
	clusterRef := obj.Spec.ClusterRef.DeepCopy()
	if err := clusterRef.Normalize(ctx, reader, obj.GetNamespace()); err != nil {
		return nil, err
	}
	clusterIdentity, err := clusterRef.ParseExternalToIdentity()
	if err != nil {
		return nil, err
	}
	clusterId := clusterIdentity.(*ContainerClusterIdentity)

	resourceID, err := refs.GetResourceID(obj)
	if err != nil {
		return nil, err
	}

	location := obj.Spec.Location
	if location == "" {
		location = clusterId.Location
	}

	identity := &ContainerNodePoolIdentity{
		Project:  clusterId.Project,
		Location: location,
		Cluster:  clusterId.Cluster,
		NodePool: resourceID,
	}

	return identity, nil
}

func (obj *ContainerNodePool) GetIdentity(ctx context.Context, reader client.Reader) (identity.Identity, error) {
	specIdentity, err := getIdentityFromContainerNodePoolSpec(ctx, reader, obj)
	if err != nil {
		return nil, err
	}

	if obj.Status.ExternalRef != nil && *obj.Status.ExternalRef != "" {
		externalRef := *obj.Status.ExternalRef
		statusIdentity := &ContainerNodePoolIdentity{}
		if err := statusIdentity.FromExternal(externalRef); err != nil {
			return nil, fmt.Errorf("cannot parse existing externalRef=%q: %w", externalRef, err)
		}
		if statusIdentity.String() != specIdentity.String() {
			return nil, fmt.Errorf("existing externalRef=%q does not match the identity resolved from spec: %q", externalRef, specIdentity.String())
		}
	}

	return specIdentity, nil
}

func (obj *ContainerNodePool) ExternalIdentifier() *string {
	return obj.Status.ExternalRef
}
