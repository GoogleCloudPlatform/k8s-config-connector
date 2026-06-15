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
	refs "github.com/GoogleCloudPlatform/k8s-config-connector/apis/refs/v1beta1"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/gcpurls"
	"sigs.k8s.io/controller-runtime/pkg/client"
)

var (
	_ identity.IdentityV2 = &GKEMulticloudAttachedClusterIdentity{}
	_ identity.Resource   = &GKEMulticloudAttachedCluster{}
)

var GKEMulticloudAttachedClusterIdentityFormat = gcpurls.Template[GKEMulticloudAttachedClusterIdentity]("gkemulticloud.googleapis.com", "projects/{project}/locations/{location}/attachedClusters/{attached_cluster}")

// +k8s:deepcopy-gen=false
type GKEMulticloudAttachedClusterIdentity struct {
	Project          string
	Location         string
	Attached_cluster string
}

func (i *GKEMulticloudAttachedClusterIdentity) String() string {
	return GKEMulticloudAttachedClusterIdentityFormat.ToString(*i)
}

func (i *GKEMulticloudAttachedClusterIdentity) FromExternal(ref string) error {
	parsed, match, err := GKEMulticloudAttachedClusterIdentityFormat.Parse(ref)
	if err != nil {
		return fmt.Errorf("format of GKEMulticloudAttachedCluster external=%q was not known (use %s): %w", ref, GKEMulticloudAttachedClusterIdentityFormat.CanonicalForm(), err)
	}
	if !match {
		return fmt.Errorf("format of GKEMulticloudAttachedCluster external=%q was not known (use %s)", ref, GKEMulticloudAttachedClusterIdentityFormat.CanonicalForm())
	}

	*i = *parsed
	return nil
}

func (i *GKEMulticloudAttachedClusterIdentity) Host() string {
	return GKEMulticloudAttachedClusterIdentityFormat.Host()
}

func (i *GKEMulticloudAttachedClusterIdentity) ID() string {
	return i.Attached_cluster
}

func (i *GKEMulticloudAttachedClusterIdentity) ParentString() string {
	return fmt.Sprintf("projects/%s/locations/%s", i.Project, i.Location)
}

func getIdentityFromGKEMulticloudAttachedClusterSpec(ctx context.Context, reader client.Reader, obj client.Object) (*GKEMulticloudAttachedClusterIdentity, error) {
	resourceID, err := refs.GetResourceID(obj)
	if err != nil {
		return nil, fmt.Errorf("cannot resolve resource ID")
	}

	location, err := refs.GetLocation(obj)
	if err != nil {
		return nil, fmt.Errorf("cannot resolve location")
	}

	projectID, err := refs.ResolveProjectID(ctx, reader, obj)
	if err != nil {
		return nil, fmt.Errorf("cannot resolve project")
	}

	identity := &GKEMulticloudAttachedClusterIdentity{
		Project:          projectID,
		Location:         location,
		Attached_cluster: resourceID,
	}
	return identity, nil
}

func (obj *GKEMulticloudAttachedCluster) GetIdentity(ctx context.Context, reader client.Reader) (identity.Identity, error) {
	specIdentity, err := getIdentityFromGKEMulticloudAttachedClusterSpec(ctx, reader, obj)
	if err != nil {
		return nil, err
	}

	// Cross-check the identity against the status value, if present.
	externalRef := common.ValueOf(obj.Status.ExternalRef)
	if externalRef != "" {
		// Validate desired with actual
		statusIdentity := &GKEMulticloudAttachedClusterIdentity{}
		if err := statusIdentity.FromExternal(externalRef); err != nil {
			return nil, err
		}

		if statusIdentity.String() != specIdentity.String() {
			return nil, fmt.Errorf("cannot change GKEMulticloudAttachedCluster identity (old=%q, new=%q)", statusIdentity.String(), specIdentity.String())
		}
	}

	return specIdentity, nil
}

func (obj *GKEMulticloudAttachedCluster) ExternalIdentifier() *string {
	return obj.Status.ExternalRef
}
