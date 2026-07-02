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
	_ identity.IdentityV2 = &ComputeHealthCheckIdentity{}
	_ identity.Resource   = &ComputeHealthCheck{}
)

var GlobalComputeHealthCheckIdentityFormat = gcpurls.Template[ComputeHealthCheckIdentity]("compute.googleapis.com", "projects/{project}/global/healthChecks/{healthCheck}")
var RegionalComputeHealthCheckIdentityFormat = gcpurls.Template[ComputeHealthCheckIdentity]("compute.googleapis.com", "projects/{project}/regions/{region}/healthChecks/{healthCheck}")

// +k8s:deepcopy-gen=false
type ComputeHealthCheckIdentity struct {
	Project     string
	Region      string
	HealthCheck string
}

func (i *ComputeHealthCheckIdentity) String() string {
	if i.Region == "global" || i.Region == "" {
		return GlobalComputeHealthCheckIdentityFormat.ToString(*i)
	}
	return RegionalComputeHealthCheckIdentityFormat.ToString(*i)
}

func (i *ComputeHealthCheckIdentity) FromExternal(ref string) error {
	if parsed, match, err := RegionalComputeHealthCheckIdentityFormat.Parse(ref); match && err == nil {
		*i = *parsed
		return nil
	} else if err != nil {
		return err
	}
	if parsed, match, err := GlobalComputeHealthCheckIdentityFormat.Parse(ref); match && err == nil {
		*i = *parsed
		i.Region = "global"
		return nil
	} else if err != nil {
		return err
	}
	return fmt.Errorf("format of ComputeHealthCheck external=%q was not known (use %s or %s)", ref, RegionalComputeHealthCheckIdentityFormat.CanonicalForm(), GlobalComputeHealthCheckIdentityFormat.CanonicalForm())
}

func (i *ComputeHealthCheckIdentity) Host() string {
	return GlobalComputeHealthCheckIdentityFormat.Host()
}

func getIdentityFromComputeHealthCheckSpec(ctx context.Context, reader client.Reader, obj client.Object) (*ComputeHealthCheckIdentity, error) {
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

	identity := &ComputeHealthCheckIdentity{
		Project:     projectID,
		Region:      location,
		HealthCheck: resourceID,
	}
	return identity, nil
}

func (obj *ComputeHealthCheck) GetIdentity(ctx context.Context, reader client.Reader) (identity.Identity, error) {
	specIdentity, err := getIdentityFromComputeHealthCheckSpec(ctx, reader, obj)
	if err != nil {
		return nil, err
	}

	// Cross-check the identity against the status value, if present.
	if obj.Status.SelfLink != nil && *obj.Status.SelfLink != "" {
		selfLink := *obj.Status.SelfLink
		// Validate desired with actual
		statusIdentity := &ComputeHealthCheckIdentity{}
		if err := statusIdentity.FromExternal(selfLink); err != nil {
			return nil, err
		}

		if statusIdentity.String() != specIdentity.String() {
			return nil, fmt.Errorf("cannot change ComputeHealthCheck identity (old=%q, new=%q)", statusIdentity.String(), specIdentity.String())
		}
	}

	return specIdentity, nil
}
