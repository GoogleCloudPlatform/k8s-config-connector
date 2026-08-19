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

	"github.com/GoogleCloudPlatform/k8s-config-connector/apis/common/identity"
	refs "github.com/GoogleCloudPlatform/k8s-config-connector/apis/refs/v1beta1"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/gcpurls"
	"sigs.k8s.io/controller-runtime/pkg/client"
)

var (
	_ identity.IdentityV2 = &AssuredWorkloadsWorkloadIdentity{}
)

var (
	WorkloadIdentityFormat = gcpurls.Template[AssuredWorkloadsWorkloadIdentity]("assuredworkloads.googleapis.com", "organizations/{Organization}/locations/{Location}/workloads/{Workload}")
)

// AssuredWorkloadsWorkloadIdentity defines the resource reference to AssuredWorkloadsWorkload, which "External" field
// holds the GCP identifier for the KRM object.
// +k8s:deepcopy-gen=false
type AssuredWorkloadsWorkloadIdentity struct {
	Organization string
	Location     string
	Workload     string
}

func (i *AssuredWorkloadsWorkloadIdentity) String() string {
	return WorkloadIdentityFormat.ToString(*i)
}

func (i *AssuredWorkloadsWorkloadIdentity) ParentString() string {
	return fmt.Sprintf("organizations/%s/locations/%s", i.Organization, i.Location)
}

func (i *AssuredWorkloadsWorkloadIdentity) FromExternal(ref string) error {
	parsed, match, err := WorkloadIdentityFormat.Parse(ref)
	if err != nil {
		return fmt.Errorf("format of AssuredWorkloadsWorkload external=%q was not known (use %s): %w", ref, WorkloadIdentityFormat.CanonicalForm(), err)
	}
	if !match {
		return fmt.Errorf("format of AssuredWorkloadsWorkload external=%q was not known (use %s)", ref, WorkloadIdentityFormat.CanonicalForm())
	}

	*i = *parsed
	return nil
}

func (i *AssuredWorkloadsWorkloadIdentity) Host() string {
	return WorkloadIdentityFormat.Host()
}

func getIdentityFromAssuredWorkloadsWorkloadSpec(ctx context.Context, reader client.Reader, obj *AssuredWorkloadsWorkload) (*AssuredWorkloadsWorkloadIdentity, error) {
	if obj.Spec.OrganizationRef == nil {
		return nil, fmt.Errorf("organizationRef must be set")
	}
	orgID, err := refs.ResolveOrganization(ctx, reader, obj, obj.Spec.OrganizationRef)
	if err != nil {
		return nil, fmt.Errorf("cannot resolve organizationRef: %w", err)
	}

	resourceID, err := refs.GetResourceID(obj)
	if err != nil {
		return nil, fmt.Errorf("cannot resolve resource ID: %w", err)
	}

	return &AssuredWorkloadsWorkloadIdentity{
		Organization: orgID.OrganizationID,
		Location:     obj.Spec.Location,
		Workload:     resourceID,
	}, nil
}

func ParseWorkloadExternal(external string) (*AssuredWorkloadsWorkloadIdentity, error) {
	// formats:
	// organizations/{organization}/locations/{location}/workloads/{workload}
	tokens := strings.Split(external, "/")
	if len(tokens) == 6 && tokens[0] == "organizations" && tokens[2] == "locations" && tokens[4] == "workloads" {
		return &AssuredWorkloadsWorkloadIdentity{
			Organization: tokens[1],
			Location:     tokens[3],
			Workload:     tokens[5],
		}, nil
	}
	return nil, fmt.Errorf("invalid external reference: %s", external)
}
