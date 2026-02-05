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

	"github.com/GoogleCloudPlatform/k8s-config-connector/apis/common"
	"github.com/GoogleCloudPlatform/k8s-config-connector/apis/common/identity"
	refs "github.com/GoogleCloudPlatform/k8s-config-connector/apis/refs/v1beta1"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/gcpurls"
	"sigs.k8s.io/controller-runtime/pkg/client"
)

var (
	_ identity.IdentityV2 = &AlloyDBInstanceIdentity{}
	_ identity.Resource   = &AlloyDBInstance{}
)

var AlloyDBInstanceIdentityFormat = gcpurls.Template[AlloyDBInstanceIdentity]("alloydb.googleapis.com", "projects/{project}/locations/{location}/clusters/{cluster}/instances/{instance}")

// +k8s:deepcopy-gen=false
type AlloyDBInstanceIdentity struct {
	Project  string
	Location string
	Cluster  string
	Instance string
}

func (i *AlloyDBInstanceIdentity) String() string {
	return AlloyDBInstanceIdentityFormat.ToString(*i)
}

func (i *AlloyDBInstanceIdentity) FromExternal(ref string) error {
	parsed, match, err := AlloyDBInstanceIdentityFormat.Parse(ref)
	if err != nil {
		return fmt.Errorf("format of AlloyDBInstance external=%q was not known (use %s): %w", ref, AlloyDBInstanceIdentityFormat.CanonicalForm(), err)
	}
	if !match {
		return fmt.Errorf("format of AlloyDBInstance external=%q was not known (use %s)", ref, AlloyDBInstanceIdentityFormat.CanonicalForm())
	}

	*i = *parsed
	return nil
}

func (i *AlloyDBInstanceIdentity) Host() string {
	return AlloyDBInstanceIdentityFormat.Host()
}

func (i *AlloyDBInstanceIdentity) ID() string {
	return i.Instance
}

func (i *AlloyDBInstanceIdentity) ParentString() string {
	return fmt.Sprintf("projects/%s/locations/%s/clusters/%s", i.Project, i.Location, i.Cluster)
}

func getIdentityFromAlloyDBInstanceSpec(ctx context.Context, reader client.Reader, obj *AlloyDBInstance) (*AlloyDBInstanceIdentity, error) {
	clusterRef, err := refs.ResolveAlloyDBCluster(ctx, reader, obj, obj.Spec.ClusterRef)
	if err != nil {
		return nil, err
	}

	resourceID, err := refs.GetResourceID(obj)
	if err != nil {
		return nil, err
	}

	return &AlloyDBInstanceIdentity{
		Project:  clusterRef.ProjectID,
		Location: clusterRef.Location,
		Cluster:  clusterRef.ClusterID,
		Instance: resourceID,
	}, nil
}

func (obj *AlloyDBInstance) GetIdentity(ctx context.Context, reader client.Reader) (identity.Identity, error) {
	specIdentity, err := getIdentityFromAlloyDBInstanceSpec(ctx, reader, obj)
	if err != nil {
		return nil, err
	}

	// Cross-check the identity against the status value, if present.
	externalRef := common.ValueOf(obj.Status.ExternalRef)
	if externalRef != "" {
		// Validate desired with actual
		statusIdentity := &AlloyDBInstanceIdentity{}
		if err := statusIdentity.FromExternal(externalRef); err != nil {
			return nil, err
		}

		if statusIdentity.String() != specIdentity.String() {
			return nil, fmt.Errorf("cannot change AlloyDBInstance identity (old=%q, new=%q)", statusIdentity.String(), specIdentity.String())
		}
	}

	return specIdentity, nil
}
