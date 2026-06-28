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
	apirefs "github.com/GoogleCloudPlatform/k8s-config-connector/apis/refs"
	refs "github.com/GoogleCloudPlatform/k8s-config-connector/apis/refs/v1beta1"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/gcpurls"
	"sigs.k8s.io/controller-runtime/pkg/client"
)

var (
	_ identity.IdentityV2 = &ComputeHTTPSHealthCheckIdentity{}
	_ identity.Resource   = &ComputeHTTPSHealthCheck{}
)

var ComputeHTTPSHealthCheckIdentityFormat = gcpurls.Template[ComputeHTTPSHealthCheckIdentity]("compute.googleapis.com", "projects/{project}/global/httpsHealthChecks/{httpshealthcheck}")

// ComputeHTTPSHealthCheckIdentity is the identity of a GCP ComputeHTTPSHealthCheck resource.
// +k8s:deepcopy-gen=false
type ComputeHTTPSHealthCheckIdentity struct {
	Project          string
	HttpsHealthCheck string
}

func (i *ComputeHTTPSHealthCheckIdentity) String() string {
	return ComputeHTTPSHealthCheckIdentityFormat.ToString(*i)
}

func (i *ComputeHTTPSHealthCheckIdentity) FromExternal(ref string) error {
	trimmedRef := apirefs.TrimComputeURIPrefix(ref)
	parsed, match, err := ComputeHTTPSHealthCheckIdentityFormat.Parse(trimmedRef)
	if err != nil {
		return fmt.Errorf("format of ComputeHTTPSHealthCheck external=%q was not known (use %s): %w", ref, ComputeHTTPSHealthCheckIdentityFormat.CanonicalForm(), err)
	}
	if !match {
		return fmt.Errorf("format of ComputeHTTPSHealthCheck external=%q was not known (use %s)", ref, ComputeHTTPSHealthCheckIdentityFormat.CanonicalForm())
	}

	*i = *parsed
	return nil
}

func (i *ComputeHTTPSHealthCheckIdentity) Host() string {
	return ComputeHTTPSHealthCheckIdentityFormat.Host()
}

func (i *ComputeHTTPSHealthCheckIdentity) ParentString() string {
	return fmt.Sprintf("projects/%s/global", i.Project)
}

func ParseComputeHTTPSHealthCheckExternal(external string) (*ComputeHTTPSHealthCheckIdentity, error) {
	if external == "" {
		return nil, fmt.Errorf("empty ComputeHTTPSHealthCheck external value")
	}
	id := &ComputeHTTPSHealthCheckIdentity{}
	if err := id.FromExternal(external); err != nil {
		return nil, err
	}
	return id, nil
}

func getIdentityFromComputeHTTPSHealthCheckSpec(ctx context.Context, reader client.Reader, obj *ComputeHTTPSHealthCheck) (*ComputeHTTPSHealthCheckIdentity, error) {
	resourceID, err := refs.GetResourceID(obj)
	if err != nil {
		return nil, fmt.Errorf("cannot resolve resource ID")
	}

	projectID, err := refs.ResolveProjectID(ctx, reader, obj)
	if err != nil {
		return nil, fmt.Errorf("cannot resolve project")
	}

	identity := &ComputeHTTPSHealthCheckIdentity{
		Project:          projectID,
		HttpsHealthCheck: resourceID,
	}
	return identity, nil
}

func (obj *ComputeHTTPSHealthCheck) GetIdentity(ctx context.Context, reader client.Reader) (identity.Identity, error) {
	specIdentity, err := getIdentityFromComputeHTTPSHealthCheckSpec(ctx, reader, obj)
	if err != nil {
		return nil, err
	}

	// Cross-check the identity against status.selfLink, if present.
	selfLink := common.ValueOf(obj.Status.SelfLink)
	if selfLink != "" {
		statusIdentity := &ComputeHTTPSHealthCheckIdentity{}
		if err := statusIdentity.FromExternal(selfLink); err != nil {
			return nil, err
		}

		if statusIdentity.String() != specIdentity.String() {
			return nil, fmt.Errorf("cannot change ComputeHTTPSHealthCheck identity (old=%q, new=%q)", statusIdentity.String(), specIdentity.String())
		}
	}

	return specIdentity, nil
}
