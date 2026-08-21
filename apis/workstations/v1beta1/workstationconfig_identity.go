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
	_ identity.IdentityV2 = &WorkstationConfigIdentity{}
	_ identity.Resource   = &WorkstationConfig{}
)

var WorkstationConfigIdentityFormat = gcpurls.Template[WorkstationConfigIdentity]("workstations.googleapis.com", "projects/{project}/locations/{location}/workstationClusters/{workstationcluster}/workstationConfigs/{workstationconfig}")

// +k8s:deepcopy-gen=false
type WorkstationConfigIdentity struct {
	Project            string
	Location           string
	WorkstationCluster string
	WorkstationConfig  string
}

func (i *WorkstationConfigIdentity) String() string {
	return WorkstationConfigIdentityFormat.ToString(*i)
}

func (i *WorkstationConfigIdentity) FromExternal(ref string) error {
	parsed, match, err := WorkstationConfigIdentityFormat.Parse(ref)
	if err != nil {
		return fmt.Errorf("format of WorkstationConfig external=%q was not known (use %s): %w", ref, WorkstationConfigIdentityFormat.CanonicalForm(), err)
	}
	if !match {
		return fmt.Errorf("format of WorkstationConfig external=%q was not known (use %s)", ref, WorkstationConfigIdentityFormat.CanonicalForm())
	}

	*i = *parsed
	return nil
}

func (i *WorkstationConfigIdentity) Host() string {
	return WorkstationConfigIdentityFormat.Host()
}

func getIdentityFromWorkstationConfigSpec(ctx context.Context, reader client.Reader, obj client.Object) (*WorkstationConfigIdentity, error) {
	workstationConfig, ok := obj.(*WorkstationConfig)
	if !ok {
		return nil, fmt.Errorf("object is not a WorkstationConfig")
	}

	resourceID, err := refs.GetResourceID(obj)
	if err != nil {
		return nil, fmt.Errorf("cannot resolve resource ID")
	}

	// Get Parent
	clusterRef := workstationConfig.Spec.Parent
	if clusterRef == nil {
		return nil, fmt.Errorf("no parent cluster")
	}
	if err := clusterRef.Normalize(ctx, reader, workstationConfig.Namespace); err != nil {
		return nil, fmt.Errorf("cannot resolve cluster: %w", err)
	}
	clusterExternal := clusterRef.External
	clusterIdentity := &WorkstationClusterIdentity{}
	if err := clusterIdentity.FromExternal(clusterExternal); err != nil {
		return nil, fmt.Errorf("cannot parse external cluster: %w", err)
	}

	projectID := clusterIdentity.Project
	location := clusterIdentity.Location
	cluster := clusterIdentity.WorkstationCluster

	identity := &WorkstationConfigIdentity{
		Project:            projectID,
		Location:           location,
		WorkstationCluster: cluster,
		WorkstationConfig:  resourceID,
	}
	return identity, nil
}

func (obj *WorkstationConfig) GetIdentity(ctx context.Context, reader client.Reader) (identity.Identity, error) {
	specIdentity, err := getIdentityFromWorkstationConfigSpec(ctx, reader, obj)
	if err != nil {
		return nil, err
	}

	// Cross-check the identity against the status value, if present.
	externalRef := common.ValueOf(obj.Status.ExternalRef)
	if externalRef != "" {
		// Validate desired with actual
		statusIdentity := &WorkstationConfigIdentity{}
		if err := statusIdentity.FromExternal(externalRef); err != nil {
			return nil, err
		}

		if statusIdentity.String() != specIdentity.String() {
			return nil, fmt.Errorf("cannot change WorkstationConfig identity (old=%q, new=%q)", statusIdentity.String(), specIdentity.String())
		}
	}

	return specIdentity, nil
}
