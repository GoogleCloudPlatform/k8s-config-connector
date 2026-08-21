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
	_ identity.ServerGeneratedIdentity = &IAPBrandIdentity{}
	_ identity.Resource                = &IAPBrand{}
)

var IAPBrandIdentityFormat = gcpurls.Template[IAPBrandIdentity]("iap.googleapis.com", "projects/{project}/brands/{brand}")

// IAPBrandIdentity is the identity of a GCP IAPBrand resource.
// +k8s:deepcopy-gen=false
type IAPBrandIdentity struct {
	Project string
	Brand   string
}

func (i *IAPBrandIdentity) HasIdentitySpecified() bool {
	return i.Brand != ""
}

func (i *IAPBrandIdentity) String() string {
	return IAPBrandIdentityFormat.ToString(*i)
}

func (i *IAPBrandIdentity) ParentString() string {
	return "projects/" + i.Project
}

func (i *IAPBrandIdentity) FromExternal(ref string) error {
	parsed, match, err := IAPBrandIdentityFormat.Parse(ref)
	if err != nil {
		return fmt.Errorf("format of IAPBrand external=%q was not known (use %s): %w", ref, IAPBrandIdentityFormat.CanonicalForm(), err)
	}
	if !match {
		return fmt.Errorf("format of IAPBrand external=%q was not known (use %s)", ref, IAPBrandIdentityFormat.CanonicalForm())
	}

	*i = *parsed
	return nil
}

func (i *IAPBrandIdentity) Host() string {
	return IAPBrandIdentityFormat.Host()
}

func getIdentityFromIAPBrandSpec(ctx context.Context, reader client.Reader, obj *IAPBrand) (*IAPBrandIdentity, error) {
	// For IAPBrand, resourceID is optional and can be empty.
	resourceID := common.ValueOf(obj.Spec.ResourceID)

	projectID, err := refs.ResolveProjectID(ctx, reader, obj)
	if err != nil {
		return nil, fmt.Errorf("cannot resolve project: %w", err)
	}

	identity := &IAPBrandIdentity{
		Project: projectID,
		Brand:   resourceID,
	}
	return identity, nil
}

func (obj *IAPBrand) GetIdentity(ctx context.Context, reader client.Reader) (identity.Identity, error) {
	specIdentity, err := getIdentityFromIAPBrandSpec(ctx, reader, obj)
	if err != nil {
		return nil, err
	}

	// Note: We do NOT have Status.ExternalRef on IAPBrand, so we do NOT cross-check.
	return specIdentity, nil
}
