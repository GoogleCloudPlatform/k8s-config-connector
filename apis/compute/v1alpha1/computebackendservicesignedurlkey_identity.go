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
	_ identity.IdentityV2 = &ComputeBackendServiceSignedURLKeyIdentity{}
	_ identity.Resource   = &ComputeBackendServiceSignedURLKey{}
)

var ComputeBackendServiceSignedURLKeyIdentityFormat = gcpurls.Template[ComputeBackendServiceSignedURLKeyIdentity](
	"compute.googleapis.com",
	"projects/{project}/global/backendServices/{backendService}/signedUrlKeys/{signedUrlKey}",
)

// ComputeBackendServiceSignedURLKeyIdentity is the identity of a GCP ComputeBackendServiceSignedURLKey resource.
// +k8s:deepcopy-gen=false
type ComputeBackendServiceSignedURLKeyIdentity struct {
	Project        string
	BackendService string
	SignedUrlKey   string
}

func (i *ComputeBackendServiceSignedURLKeyIdentity) String() string {
	return ComputeBackendServiceSignedURLKeyIdentityFormat.ToString(*i)
}

func (i *ComputeBackendServiceSignedURLKeyIdentity) FromExternal(ref string) error {
	parsed, match, err := ComputeBackendServiceSignedURLKeyIdentityFormat.Parse(ref)
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

func parseBackendServiceID(external string) (string, error) {
	if strings.Contains(external, "/backendServices/") {
		parts := strings.Split(external, "/backendServices/")
		if len(parts) == 2 && parts[1] != "" {
			subparts := strings.Split(parts[1], "/")
			return subparts[0], nil
		}
	}
	if strings.Contains(external, "/") {
		return "", fmt.Errorf("invalid ComputeBackendService external reference format: %q", external)
	}
	if external == "" {
		return "", fmt.Errorf("ComputeBackendService external reference is empty")
	}
	return external, nil
}

func getIdentityFromComputeBackendServiceSignedURLKeySpec(ctx context.Context, reader client.Reader, obj *ComputeBackendServiceSignedURLKey) (*ComputeBackendServiceSignedURLKeyIdentity, error) {
	resourceID, err := refs.GetResourceID(obj)
	if err != nil {
		return nil, fmt.Errorf("cannot resolve resource ID")
	}

	projectID, err := refs.ResolveProjectID(ctx, reader, obj)
	if err != nil {
		return nil, fmt.Errorf("cannot resolve project")
	}

	backendServiceExternal, err := obj.Spec.BackendServiceRef.NormalizedExternal(ctx, reader, obj.Namespace)
	if err != nil {
		return nil, fmt.Errorf("cannot resolve backend service: %w", err)
	}

	backendService, err := parseBackendServiceID(backendServiceExternal)
	if err != nil {
		return nil, fmt.Errorf("cannot parse backend service reference: %w", err)
	}

	identity := &ComputeBackendServiceSignedURLKeyIdentity{
		Project:        projectID,
		BackendService: backendService,
		SignedUrlKey:   resourceID,
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
