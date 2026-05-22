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
	_ identity.IdentityV2 = &ContainerClusterIdentity{}
	_ identity.Resource   = &ContainerCluster{}
)

var RegionalContainerClusterIdentityFormat = gcpurls.Template[ContainerClusterIdentity]("container.googleapis.com", "projects/{project}/locations/{location}/clusters/{cluster}")
var ZonalContainerClusterIdentityFormat = gcpurls.Template[ContainerClusterIdentity]("container.googleapis.com", "projects/{project}/zones/{location}/clusters/{cluster}")

// +k8s:deepcopy-gen=false
type ContainerClusterIdentity struct {
	Project  string
	Location string
	Cluster  string
}

func (i *ContainerClusterIdentity) String() string {
	return RegionalContainerClusterIdentityFormat.ToString(*i)
}

func (i *ContainerClusterIdentity) FromExternal(ref string) error {
	if idx := strings.Index(ref, "projects/"); idx != -1 {
		ref = ref[idx:]
	}
	if parsed, match, _ := RegionalContainerClusterIdentityFormat.Parse(ref); match {
		*i = *parsed
		return nil
	}
	if parsed, match, _ := ZonalContainerClusterIdentityFormat.Parse(ref); match {
		*i = *parsed
		return nil
	}
	return fmt.Errorf("format of ContainerCluster external=%q was not known (use %s or %s)", ref, RegionalContainerClusterIdentityFormat.CanonicalForm(), ZonalContainerClusterIdentityFormat.CanonicalForm())
}

func (i *ContainerClusterIdentity) Host() string {
	return RegionalContainerClusterIdentityFormat.Host()
}

func getIdentityFromContainerClusterSpec(ctx context.Context, reader client.Reader, obj client.Object) (*ContainerClusterIdentity, error) {
	resourceID, err := refs.GetResourceID(obj)
	if err != nil {
		return nil, err
	}

	location, err := refs.GetLocation(obj)
	if err != nil {
		return nil, err
	}

	projectID, err := refs.ResolveProjectID(ctx, reader, obj)
	if err != nil {
		return nil, err
	}

	identity := &ContainerClusterIdentity{
		Project:  projectID,
		Location: location,
		Cluster:  resourceID,
	}
	return identity, nil
}

func (obj *ContainerCluster) GetIdentity(ctx context.Context, reader client.Reader) (identity.Identity, error) {
	specIdentity, err := getIdentityFromContainerClusterSpec(ctx, reader, obj)
	if err != nil {
		return nil, err
	}

	// Cross-check the identity against the status value, if present.
	// NOTE: ContainerCluster does not yet support status.externalRef, but we parse `selfLink` instead
	if obj.Status.SelfLink != nil && *obj.Status.SelfLink != "" {
		selfLink := *obj.Status.SelfLink
		// Validate desired with actual
		statusIdentity := &ContainerClusterIdentity{}
		if err := statusIdentity.FromExternal(selfLink); err != nil {
			return nil, err
		}

		if statusIdentity.String() != specIdentity.String() {
			return nil, fmt.Errorf("cannot change ContainerCluster identity (old=%q, new=%q)", statusIdentity.String(), specIdentity.String())
		}
	}

	return specIdentity, nil
}
