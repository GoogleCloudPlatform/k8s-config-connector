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

	"github.com/GoogleCloudPlatform/k8s-config-connector/apis/common/identity"
	refs "github.com/GoogleCloudPlatform/k8s-config-connector/apis/refs/v1beta1"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/gcpurls"
	"sigs.k8s.io/controller-runtime/pkg/client"
)

var (
	_ identity.IdentityV2 = &CloudTalentSolutionCompanyIdentity{}
	_ identity.Resource   = &CloudTalentSolutionCompany{}
)

var CloudTalentSolutionCompanyIdentityFormat = gcpurls.Template[CloudTalentSolutionCompanyIdentity]("jobs.googleapis.com", "projects/{project}/tenants/{tenant}/companies/{company}")

// CloudTalentSolutionCompanyIdentity is the identity of a GCP CloudTalentSolutionCompany resource.
// +k8s:deepcopy-gen=false
type CloudTalentSolutionCompanyIdentity struct {
	Project string
	Tenant  string
	Company string
}

func (i *CloudTalentSolutionCompanyIdentity) String() string {
	return CloudTalentSolutionCompanyIdentityFormat.ToString(*i)
}

func (i *CloudTalentSolutionCompanyIdentity) FromExternal(ref string) error {
	parsed, match, err := CloudTalentSolutionCompanyIdentityFormat.Parse(ref)
	if err != nil {
		return fmt.Errorf("format of CloudTalentSolutionCompany external=%q was not known (use %s): %w", ref, CloudTalentSolutionCompanyIdentityFormat.CanonicalForm(), err)
	}
	if !match {
		return fmt.Errorf("format of CloudTalentSolutionCompany external=%q was not known (use %s)", ref, CloudTalentSolutionCompanyIdentityFormat.CanonicalForm())
	}

	*i = *parsed
	return nil
}

func (i *CloudTalentSolutionCompanyIdentity) Host() string {
	return CloudTalentSolutionCompanyIdentityFormat.Host()
}

func (i *CloudTalentSolutionCompanyIdentity) ParentString() string {
	return fmt.Sprintf("projects/%s/tenants/%s", i.Project, i.Tenant)
}

func getIdentityFromCloudTalentSolutionCompanySpec(ctx context.Context, reader client.Reader, obj *CloudTalentSolutionCompany) (*CloudTalentSolutionCompanyIdentity, error) {
	resourceID, err := refs.GetResourceID(obj)
	if err != nil {
		return nil, fmt.Errorf("cannot resolve resource ID: %w", err)
	}

	projectID, err := refs.ResolveProjectID(ctx, reader, obj)
	if err != nil {
		return nil, fmt.Errorf("cannot resolve project: %w", err)
	}

	tenant := ""
	if obj.Spec.Tenant != nil {
		tenant = *obj.Spec.Tenant
	}

	identity := &CloudTalentSolutionCompanyIdentity{
		Project: projectID,
		Tenant:  tenant,
		Company: resourceID,
	}
	return identity, nil
}

func (obj *CloudTalentSolutionCompany) GetIdentity(ctx context.Context, reader client.Reader) (identity.Identity, error) {
	specIdentity, err := getIdentityFromCloudTalentSolutionCompanySpec(ctx, reader, obj)
	if err != nil {
		return nil, err
	}

	if obj.Status.ExternalRef != nil {
		statusIdentity := &CloudTalentSolutionCompanyIdentity{}
		if err := statusIdentity.FromExternal(*obj.Status.ExternalRef); err != nil {
			return nil, err
		}

		if specIdentity.Tenant == "default" {
			specIdentity.Tenant = statusIdentity.Tenant
		}

		if statusIdentity.String() != specIdentity.String() {
			return nil, fmt.Errorf("cannot change CloudTalentSolutionCompany identity (old=%q, new=%q)", statusIdentity.String(), specIdentity.String())
		}
	}

	return specIdentity, nil
}
