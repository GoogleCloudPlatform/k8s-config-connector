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
	"k8s.io/apimachinery/pkg/apis/meta/v1/unstructured"
	"sigs.k8s.io/controller-runtime/pkg/client"
)

var (
	_ identity.IdentityV2 = &ComputeResourcePolicyIdentity{}
	_ identity.Resource   = &ComputeResourcePolicy{}
)

var ComputeResourcePolicyIdentityFormat = gcpurls.Template[ComputeResourcePolicyIdentity](
	"compute.googleapis.com",
	"projects/{Project}/regions/{Region}/resourcePolicies/{ResourcePolicy}",
)

// +k8s:deepcopy-gen=false
type ComputeResourcePolicyIdentity struct {
	Project        string
	Region         string
	ResourcePolicy string
}

func (i *ComputeResourcePolicyIdentity) String() string {
	return ComputeResourcePolicyIdentityFormat.ToString(*i)
}

func (i *ComputeResourcePolicyIdentity) FromExternal(ref string) error {
	parsed, match, err := ComputeResourcePolicyIdentityFormat.Parse(ref)
	if err != nil {
		return fmt.Errorf("format of ComputeResourcePolicy external=%q was not known (use %s): %w", ref, ComputeResourcePolicyIdentityFormat.CanonicalForm(), err)
	}
	if !match {
		return fmt.Errorf("format of ComputeResourcePolicy external=%q was not known (use %s)", ref, ComputeResourcePolicyIdentityFormat.CanonicalForm())
	}

	*i = *parsed
	return nil
}

func (i *ComputeResourcePolicyIdentity) Host() string {
	return ComputeResourcePolicyIdentityFormat.Host()
}

func getIdentityFromComputeResourcePolicySpec(ctx context.Context, reader client.Reader, obj client.Object) (*ComputeResourcePolicyIdentity, error) {
	resourceID, err := refs.GetResourceID(obj)
	if err != nil {
		return nil, fmt.Errorf("cannot resolve resource ID: %w", err)
	}

	u, ok := obj.(*unstructured.Unstructured)
	if !ok {
		return nil, fmt.Errorf("expected *unstructured.Unstructured, got %T", obj)
	}

	region, _, err := unstructured.NestedString(u.Object, "spec", "region")
	if err != nil {
		return nil, fmt.Errorf("cannot resolve region: %w", err)
	}
	if region == "" {
		return nil, fmt.Errorf("region is required but not found in spec")
	}

	projectID, err := refs.ResolveProjectID(ctx, reader, obj)
	if err != nil {
		return nil, fmt.Errorf("cannot resolve project: %w", err)
	}

	identity := &ComputeResourcePolicyIdentity{
		Project:        projectID,
		Region:         region,
		ResourcePolicy: resourceID,
	}
	return identity, nil
}

func (obj *ComputeResourcePolicy) GetIdentity(ctx context.Context, reader client.Reader) (identity.Identity, error) {
	resourceID := common.ValueOf(obj.Spec.ResourceID)
	if resourceID == "" {
		resourceID = obj.GetName()
	}

	if resourceID == "" {
		return nil, fmt.Errorf("cannot resolve resource ID")
	}

	projectID, err := refs.ResolveProjectID(ctx, reader, obj)
	if err != nil {
		return nil, err
	}

	region := common.ValueOf(obj.Spec.Region)
	if region == "" {
		return nil, fmt.Errorf("cannot resolve region")
	}

	specIdentity := &ComputeResourcePolicyIdentity{
		Project:        projectID,
		Region:         region,
		ResourcePolicy: resourceID,
	}

	externalRef := common.ValueOf(obj.Status.ExternalRef)
	if externalRef != "" {
		statusIdentity := &ComputeResourcePolicyIdentity{}
		if err := statusIdentity.FromExternal(externalRef); err != nil {
			return nil, err
		}

		if statusIdentity.String() != specIdentity.String() {
			return nil, fmt.Errorf("cannot change ComputeResourcePolicy identity (old=%q, new=%q)", statusIdentity.String(), specIdentity.String())
		}
	}

	return specIdentity, nil
}
