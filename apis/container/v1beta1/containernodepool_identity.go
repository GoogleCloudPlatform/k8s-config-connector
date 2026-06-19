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
	"strings"

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
var ZonalContainerNodePoolIdentityFormat = gcpurls.Template[ContainerNodePoolIdentity]("container.googleapis.com", "projects/{project}/zones/{zone}/clusters/{cluster}/nodePools/{nodePool}")

// +k8s:deepcopy-gen=false
// ContainerNodePoolIdentity is the identity of a GCP ContainerNodePool resource.
type ContainerNodePoolIdentity struct {
	Project  string
	Location string
	Zone     string
	Cluster  string
	NodePool string
}

func (i *ContainerNodePoolIdentity) String() string {
	if i.Zone != "" {
		return ZonalContainerNodePoolIdentityFormat.ToString(*i)
	}
	return RegionalContainerNodePoolIdentityFormat.ToString(*i)
}

func (i *ContainerNodePoolIdentity) FromExternal(ref string) error {
	if idx := strings.Index(ref, "projects/"); idx != -1 {
		ref = ref[idx:]
	}
	if parsed, match, _ := ZonalContainerNodePoolIdentityFormat.Parse(ref); match {
		*i = *parsed
		return nil
	}
	if parsed, match, _ := RegionalContainerNodePoolIdentityFormat.Parse(ref); match {
		*i = *parsed
		return nil
	}
	return fmt.Errorf("format of ContainerNodePool external=%q was not known (use %s or %s)", ref, ZonalContainerNodePoolIdentityFormat.CanonicalForm(), RegionalContainerNodePoolIdentityFormat.CanonicalForm())
}

func (i *ContainerNodePoolIdentity) Host() string {
	return RegionalContainerNodePoolIdentityFormat.Host()
}

// ParentString returns the parent ContainerCluster GCP URI.
func (i *ContainerNodePoolIdentity) ParentString() string {
	if i.Zone != "" {
		return fmt.Sprintf("projects/%s/zones/%s/clusters/%s", i.Project, i.Zone, i.Cluster)
	}
	return fmt.Sprintf("projects/%s/locations/%s/clusters/%s", i.Project, i.Location, i.Cluster)
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
		Cluster:  clusterId.Cluster,
		NodePool: resourceID,
	}

	if len(strings.Split(location, "-")) == 3 {
		identity.Zone = location
	} else {
		identity.Location = location
	}

	return identity, nil
}

func (obj *ContainerNodePool) GetIdentity(ctx context.Context, reader client.Reader) (identity.Identity, error) {
	return getIdentityFromContainerNodePoolSpec(ctx, reader, obj)
}
