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
	_ identity.IdentityV2 = &ContainerClusterIdentity{}
	_ identity.Resource   = &ContainerCluster{}
)

var ContainerClusterIdentityFormat = gcpurls.Template[ContainerClusterIdentity]("container.googleapis.com", "projects/{project}/locations/{location}/clusters/{cluster}")

// +k8s:deepcopy-gen=false
type ContainerClusterIdentity struct {
	Project  string
	Location string
	Cluster  string
}

func (i *ContainerClusterIdentity) String() string {
	return ContainerClusterIdentityFormat.ToString(*i)
}

func (i *ContainerClusterIdentity) FromExternal(ref string) error {
	if parsed, match, _ := ContainerClusterIdentityFormat.Parse(ref); match {
		*i = *parsed
		return nil
	}
	// Try zonal format
	zonalFormat := gcpurls.Template[ContainerClusterIdentity]("container.googleapis.com", "projects/{project}/zones/{location}/clusters/{cluster}")
	if parsed, match, _ := zonalFormat.Parse(ref); match {
		*i = *parsed
		return nil
	}
	return fmt.Errorf("format of ContainerCluster external=%q was not known (use %s)", ref, ContainerClusterIdentityFormat.CanonicalForm())
}

func (i *ContainerClusterIdentity) Host() string {
	return ContainerClusterIdentityFormat.Host()
}

func getIdentityFromContainerClusterSpec(ctx context.Context, reader client.Reader, obj client.Object) (*ContainerClusterIdentity, error) {
	cluster, ok := obj.(*ContainerCluster)
	if !ok {
		return nil, fmt.Errorf("object is not a ContainerCluster")
	}

	resourceID := cluster.Spec.ResourceID
	if resourceID == nil || *resourceID == "" {
		resourceID = &cluster.Name
	}

	location := cluster.Spec.Location
	if location == nil || *location == "" {
		return nil, fmt.Errorf("cannot resolve location")
	}

	projectID, err := refs.ResolveProjectID(ctx, reader, obj)
	if err != nil {
		return nil, fmt.Errorf("cannot resolve project")
	}

	identity := &ContainerClusterIdentity{
		Project:  projectID,
		Location: *location,
		Cluster:  *resourceID,
	}
	return identity, nil
}

func (obj *ContainerCluster) GetIdentity(ctx context.Context, reader client.Reader) (identity.Identity, error) {
	return getIdentityFromContainerClusterSpec(ctx, reader, obj)
}
