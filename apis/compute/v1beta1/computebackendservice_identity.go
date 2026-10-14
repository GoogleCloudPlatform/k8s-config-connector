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
	_ identity.IdentityV2 = &ComputeBackendServiceIdentity{}
	_ identity.Resource   = &ComputeBackendService{}
)

var ComputeGlobalBackendServiceIdentityFormat = gcpurls.Template[ComputeBackendServiceIdentity]("compute.googleapis.com", "projects/{project}/global/backendServices/{backendservice}")
var ComputeRegionalBackendServiceIdentityFormat = gcpurls.Template[ComputeBackendServiceIdentity]("compute.googleapis.com", "projects/{project}/regions/{location}/backendServices/{backendservice}")

// ComputeBackendServiceIdentity is the identity of a GCP ComputeBackendService resource.
// +k8s:deepcopy-gen=false
type ComputeBackendServiceIdentity struct {
	Project        string
	Location       string
	BackendService string
}

func (i *ComputeBackendServiceIdentity) IsGlobal() bool {
	return i.Location == "" || i.Location == "global"
}

func (i *ComputeBackendServiceIdentity) String() string {
	if !i.IsGlobal() {
		return ComputeRegionalBackendServiceIdentityFormat.ToString(*i)
	}
	return ComputeGlobalBackendServiceIdentityFormat.ToString(*i)
}

func (i *ComputeBackendServiceIdentity) FromExternal(ref string) error {
	ref = refs.TrimComputeURIPrefix(ref)

	if parsed, match, _ := ComputeGlobalBackendServiceIdentityFormat.Parse(ref); match {
		*i = *parsed
		return nil
	}
	if parsed, match, _ := ComputeRegionalBackendServiceIdentityFormat.Parse(ref); match {
		*i = *parsed
		return nil
	}
	return fmt.Errorf("format of ComputeBackendService external=%q was not known (use %s or %s)", ref, ComputeGlobalBackendServiceIdentityFormat.CanonicalForm(), ComputeRegionalBackendServiceIdentityFormat.CanonicalForm())
}

func (i *ComputeBackendServiceIdentity) Host() string {
	return ComputeGlobalBackendServiceIdentityFormat.Host()
}

func getIdentityFromComputeBackendServiceSpec(ctx context.Context, reader client.Reader, obj *ComputeBackendService) (*ComputeBackendServiceIdentity, error) {
	resourceID, err := refs.GetResourceID(obj)
	if err != nil {
		return nil, fmt.Errorf("cannot resolve resource ID: %w", err)
	}

	projectID, err := refs.ResolveProjectID(ctx, reader, obj)
	if err != nil {
		return nil, fmt.Errorf("cannot resolve project: %w", err)
	}

	location := obj.Spec.Location
	if location == "" {
		location = "global"
	}

	identity := &ComputeBackendServiceIdentity{
		Project:        projectID,
		Location:       location,
		BackendService: resourceID,
	}

	return identity, nil
}

func (obj *ComputeBackendService) GetIdentity(ctx context.Context, reader client.Reader) (identity.Identity, error) {
	specIdentity, err := getIdentityFromComputeBackendServiceSpec(ctx, reader, obj)
	if err != nil {
		return nil, err
	}

	// Cross-check the identity against the status value, if present.
	externalRef := common.ValueOf(obj.Status.SelfLink)
	if externalRef != "" {
		statusIdentity := &ComputeBackendServiceIdentity{}
		if err := statusIdentity.FromExternal(externalRef); err != nil {
			return nil, err
		}

		if statusIdentity.String() != specIdentity.String() {
			return nil, fmt.Errorf("cannot change ComputeBackendService identity (old=%q, new=%q)", statusIdentity.String(), specIdentity.String())
		}
	}

	return specIdentity, nil
}
