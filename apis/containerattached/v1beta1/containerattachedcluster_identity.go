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
	_ identity.IdentityV2 = &ContainerAttachedClusterIdentity{}
	_ identity.Resource   = &ContainerAttachedCluster{}
)

var ContainerAttachedClusterIdentityFormat = gcpurls.Template[ContainerAttachedClusterIdentity]("gkemulticloud.googleapis.com", "projects/{project}/locations/{location}/attachedClusters/{containerattachedcluster}")

// +k8s:deepcopy-gen=false
type ContainerAttachedClusterIdentity struct {
	Project                  string
	Location                 string
	ContainerAttachedCluster string
}

func (i *ContainerAttachedClusterIdentity) String() string {
	return ContainerAttachedClusterIdentityFormat.ToString(*i)
}

func (i *ContainerAttachedClusterIdentity) FromExternal(ref string) error {
	parsed, match, err := ContainerAttachedClusterIdentityFormat.Parse(ref)
	if err != nil {
		return fmt.Errorf("format of ContainerAttachedCluster external=%q was not known (use %s): %w", ref, ContainerAttachedClusterIdentityFormat.CanonicalForm(), err)
	}
	if !match {
		return fmt.Errorf("format of ContainerAttachedCluster external=%q was not known (use %s)", ref, ContainerAttachedClusterIdentityFormat.CanonicalForm())
	}

	*i = *parsed
	return nil
}

func (i *ContainerAttachedClusterIdentity) Host() string {
	return ContainerAttachedClusterIdentityFormat.Host()
}

func getIdentityFromContainerAttachedClusterSpec(ctx context.Context, reader client.Reader, obj client.Object) (*ContainerAttachedClusterIdentity, error) {
	resourceID, err := refs.GetResourceID(obj)
	if err != nil {
		return nil, fmt.Errorf("cannot resolve resource ID")
	}

	location, err := refs.GetLocation(obj)
	if err != nil {
		return nil, fmt.Errorf("cannot resolve resource ID")
	}

	projectID, err := refs.ResolveProjectID(ctx, reader, obj)
	if err != nil {
		return nil, fmt.Errorf("cannot resolve project")
	}

	identity := &ContainerAttachedClusterIdentity{
		Project:                  projectID,
		Location:                 location,
		ContainerAttachedCluster: resourceID,
	}
	return identity, nil
}

func (obj *ContainerAttachedCluster) GetIdentity(ctx context.Context, reader client.Reader) (identity.Identity, error) {
	specIdentity, err := getIdentityFromContainerAttachedClusterSpec(ctx, reader, obj)
	if err != nil {
		return nil, err
	}

	return specIdentity, nil
}
