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
	apirefs "github.com/GoogleCloudPlatform/k8s-config-connector/apis/refs"
	refs "github.com/GoogleCloudPlatform/k8s-config-connector/apis/refs/v1beta1"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/gcpurls"
	"sigs.k8s.io/controller-runtime/pkg/client"
)

var (
	_ identity.IdentityV2 = &ComputeBackendServiceSignedURLKeyIdentity{}
	_ identity.Resource   = &ComputeBackendServiceSignedURLKey{}
)

var ComputeBackendServiceSignedURLKeyIdentityFormat = gcpurls.Template[ComputeBackendServiceSignedURLKeyIdentity](
	"compute.googleapis.com",
	"projects/{project}/global/backendServices/{backendservice}/signedUrlKeys/{name}",
)

// ComputeBackendServiceSignedURLKeyIdentity is the identity of a GCP ComputeBackendServiceSignedURLKey resource.
// +k8s:deepcopy-gen=false
type ComputeBackendServiceSignedURLKeyIdentity struct {
	Project        string
	BackendService string
	Name           string
}

func (i *ComputeBackendServiceSignedURLKeyIdentity) String() string {
	return ComputeBackendServiceSignedURLKeyIdentityFormat.ToString(*i)
}

func (i *ComputeBackendServiceSignedURLKeyIdentity) FromExternal(ref string) error {
	trimmedRef := apirefs.TrimComputeURIPrefix(ref)
	parsed, match, err := ComputeBackendServiceSignedURLKeyIdentityFormat.Parse(trimmedRef)
	if err != nil {
		return fmt.Errorf("format of ComputeBackendServiceSignedURLKey external=%q was not known (use %s): %w", ref, ComputeBackendServiceSignedURLKeyIdentityFormat.CanonicalForm(), err)
	}
	if !match {
		return fmt.Errorf("format of ComputeBackendServiceSignedURLKey external=%q was not known (use %s)", ref, ComputeBackendServiceSignedURLKeyIdentityFormat.CanonicalForm())
	}

	*i = *parsed
	return nil
}

func (i *ComputeBackendServiceSignedURLKeyIdentity) Host() string {
	return ComputeBackendServiceSignedURLKeyIdentityFormat.Host()
}

func (i *ComputeBackendServiceSignedURLKeyIdentity) ParentString() string {
	return fmt.Sprintf("projects/%s/global/backendServices/%s", i.Project, i.BackendService)
}

func ParseComputeBackendServiceSignedURLKeyExternal(external string) (*ComputeBackendServiceSignedURLKeyIdentity, error) {
	if external == "" {
		return nil, fmt.Errorf("empty ComputeBackendServiceSignedURLKey external value")
	}
	id := &ComputeBackendServiceSignedURLKeyIdentity{}
	if err := id.FromExternal(external); err != nil {
		return nil, err
	}
	return id, nil
}

func getIdentityFromComputeBackendServiceSignedURLKeySpec(ctx context.Context, reader client.Reader, obj *ComputeBackendServiceSignedURLKey) (*ComputeBackendServiceSignedURLKeyIdentity, error) {
	resourceID, err := refs.GetResourceID(obj)
	if err != nil {
		return nil, fmt.Errorf("cannot resolve resource ID")
	}

	projectID, err := refs.ResolveProjectID(ctx, reader, obj)
	if err != nil {
		return nil, fmt.Errorf("cannot resolve project: %w", err)
	}

	backendServiceExternal, err := obj.Spec.BackendServiceRef.NormalizedExternal(ctx, reader, obj.Namespace)
	if err != nil {
		return nil, fmt.Errorf("cannot resolve backendServiceRef: %w", err)
	}

	// Trim prefix and extract backend service name from the selfLink url.
	// Format is: projects/{project}/global/backendServices/{backendService}
	backendServiceRefTrimmed := apirefs.TrimComputeURIPrefix(backendServiceExternal)
	parts := strings.Split(backendServiceRefTrimmed, "/")
	if len(parts) < 5 || parts[0] != "projects" || parts[2] != "global" || parts[3] != "backendServices" {
		return nil, fmt.Errorf("invalid resolved backendServiceRef URL format: %q", backendServiceExternal)
	}
	backendServiceName := parts[4]

	identity := &ComputeBackendServiceSignedURLKeyIdentity{
		Project:        projectID,
		BackendService: backendServiceName,
		Name:           resourceID,
	}
	return identity, nil
}

func (obj *ComputeBackendServiceSignedURLKey) GetIdentity(ctx context.Context, reader client.Reader) (identity.Identity, error) {
	specIdentity, err := getIdentityFromComputeBackendServiceSignedURLKeySpec(ctx, reader, obj)
	if err != nil {
		return nil, err
	}

	return specIdentity, nil
}
