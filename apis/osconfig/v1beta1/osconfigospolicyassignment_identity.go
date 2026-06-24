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
	_ identity.IdentityV2 = &OSConfigOSPolicyAssignmentIdentity{}
	_ identity.Resource   = &OSConfigOSPolicyAssignment{}
)

var OSConfigOSPolicyAssignmentIdentityFormat = gcpurls.Template[OSConfigOSPolicyAssignmentIdentity]("osconfig.googleapis.com", "projects/{project}/locations/{location}/osPolicyAssignments/{osConfigOSPolicyAssignment}")

// +k8s:deepcopy-gen=false
type OSConfigOSPolicyAssignmentIdentity struct {
	Project                    string
	Location                   string
	OSConfigOSPolicyAssignment string
}

func (i *OSConfigOSPolicyAssignmentIdentity) String() string {
	return OSConfigOSPolicyAssignmentIdentityFormat.ToString(*i)
}

func (i *OSConfigOSPolicyAssignmentIdentity) FromExternal(ref string) error {
	parsed, match, err := OSConfigOSPolicyAssignmentIdentityFormat.Parse(ref)
	if err != nil {
		return fmt.Errorf("format of OSConfigOSPolicyAssignment external=%q was not known (use %s): %w", ref, OSConfigOSPolicyAssignmentIdentityFormat.CanonicalForm(), err)
	}
	if !match {
		return fmt.Errorf("format of OSConfigOSPolicyAssignment external=%q was not known (use %s)", ref, OSConfigOSPolicyAssignmentIdentityFormat.CanonicalForm())
	}

	*i = *parsed
	return nil
}

func (i *OSConfigOSPolicyAssignmentIdentity) Host() string {
	return OSConfigOSPolicyAssignmentIdentityFormat.Host()
}

func getIdentityFromOSConfigOSPolicyAssignmentSpec(ctx context.Context, reader client.Reader, obj client.Object) (*OSConfigOSPolicyAssignmentIdentity, error) {
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

	identity := &OSConfigOSPolicyAssignmentIdentity{
		Project:                    projectID,
		Location:                   location,
		OSConfigOSPolicyAssignment: resourceID,
	}
	return identity, nil
}

func (obj *OSConfigOSPolicyAssignment) GetIdentity(ctx context.Context, reader client.Reader) (identity.Identity, error) {
	specIdentity, err := getIdentityFromOSConfigOSPolicyAssignmentSpec(ctx, reader, obj)
	if err != nil {
		return nil, err
	}

	return specIdentity, nil
}
