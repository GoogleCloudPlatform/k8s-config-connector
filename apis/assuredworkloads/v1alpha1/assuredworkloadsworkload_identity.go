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
	"k8s.io/apimachinery/pkg/apis/meta/v1/unstructured"
	"k8s.io/apimachinery/pkg/runtime"
	"sigs.k8s.io/controller-runtime/pkg/client"
)

var (
	_ identity.IdentityV2 = &AssuredWorkloadsWorkloadIdentity{}
	_ identity.Resource   = &AssuredWorkloadsWorkload{}
)

var AssuredWorkloadsWorkloadIdentityFormat = gcpurls.Template[AssuredWorkloadsWorkloadIdentity]("assuredworkloads.googleapis.com", "organizations/{organization}/locations/{location}/workloads/{workload}")

// AssuredWorkloadsWorkloadIdentity is the identity of a GCP AssuredWorkloadsWorkload resource.
// +k8s:deepcopy-gen=false
type AssuredWorkloadsWorkloadIdentity struct {
	Organization string
	Location     string
	Workload     string
}

func (i *AssuredWorkloadsWorkloadIdentity) String() string {
	return AssuredWorkloadsWorkloadIdentityFormat.ToString(*i)
}

func (i *AssuredWorkloadsWorkloadIdentity) FromExternal(ref string) error {
	parsed, match, err := AssuredWorkloadsWorkloadIdentityFormat.Parse(ref)
	if err != nil {
		return fmt.Errorf("format of AssuredWorkloadsWorkload external=%q was not known (use %s): %w", ref, AssuredWorkloadsWorkloadIdentityFormat.CanonicalForm(), err)
	}
	if !match {
		return fmt.Errorf("format of AssuredWorkloadsWorkload external=%q was not known (use %s)", ref, AssuredWorkloadsWorkloadIdentityFormat.CanonicalForm())
	}

	*i = *parsed
	return nil
}

func (i *AssuredWorkloadsWorkloadIdentity) Host() string {
	return AssuredWorkloadsWorkloadIdentityFormat.Host()
}

func (i *AssuredWorkloadsWorkloadIdentity) ParentString() string {
	return "organizations/" + i.Organization + "/locations/" + i.Location
}

func getIdentityFromAssuredWorkloadsWorkloadSpec(ctx context.Context, reader client.Reader, obj client.Object) (*AssuredWorkloadsWorkloadIdentity, error) {
	resourceID, err := refs.GetResourceID(obj)
	if err != nil {
		return nil, fmt.Errorf("cannot resolve resource ID")
	}

	location, err := refs.GetLocation(obj)
	if err != nil {
		return nil, fmt.Errorf("cannot resolve location: %w", err)
	}

	organizationID, err := resolveWorkloadOrganizationID(ctx, reader, obj)
	if err != nil {
		return nil, fmt.Errorf("cannot resolve organization: %w", err)
	}

	identity := &AssuredWorkloadsWorkloadIdentity{
		Organization: organizationID,
		Location:     location,
		Workload:     resourceID,
	}
	return identity, nil
}

func resolveWorkloadOrganizationID(ctx context.Context, reader client.Reader, obj client.Object) (string, error) {
	u, ok := obj.(*unstructured.Unstructured)
	if !ok {
		m, err := runtime.DefaultUnstructuredConverter.ToUnstructured(obj)
		if err != nil {
			return "", fmt.Errorf("expected an Unstructured object but got %T; additionally, failed to convert to unstructured: %w", obj, err)
		}
		u = &unstructured.Unstructured{Object: m}
	}
	return refs.ResolveOrganizationID(ctx, reader, u)
}

func (obj *AssuredWorkloadsWorkload) GetIdentity(ctx context.Context, reader client.Reader) (identity.Identity, error) {
	specIdentity, err := getIdentityFromAssuredWorkloadsWorkloadSpec(ctx, reader, obj)
	if err != nil {
		return nil, err
	}

	// Cross-check the identity against the status value, if present.
	externalRef := common.ValueOf(obj.Status.ExternalRef)
	if externalRef != "" {
		// Validate desired with actual
		statusIdentity := &AssuredWorkloadsWorkloadIdentity{}
		if err := statusIdentity.FromExternal(externalRef); err != nil {
			return nil, err
		}

		if statusIdentity.String() != specIdentity.String() {
			return nil, fmt.Errorf("cannot change AssuredWorkloadsWorkload identity (old=%q, new=%q)", statusIdentity.String(), specIdentity.String())
		}
	}

	return specIdentity, nil
}
