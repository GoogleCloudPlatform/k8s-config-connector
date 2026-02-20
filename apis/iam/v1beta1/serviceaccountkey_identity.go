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
	"fmt"

	"github.com/GoogleCloudPlatform/k8s-config-connector/apis/common/identity"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/gcpurls"
)

var (
	_ identity.IdentityV2 = &ServiceAccountIdentity{}
	_ identity.IdentityV2 = &ServiceAccountKeyIdentity{}

	serviceAccountFormatTemplate    = "projects/{project}/serviceAccounts/{account}"
	serviceAccountKeyFormatTemplate = serviceAccountFormatTemplate + "/keys/{id}"

	serviceAccountFormat = gcpurls.Template[ServiceAccountIdentity](
		"iam.googleapis.com",
		serviceAccountFormatTemplate,
	)
	serviceAccountKeyFormat = gcpurls.Template[ServiceAccountKeyIdentity](
		"iam.googleapis.com",
		serviceAccountKeyFormatTemplate,
	)
)

type ServiceAccountKeyIdentity struct {
	ServiceAccountIdentity
	Id string
}

func (i *ServiceAccountKeyIdentity) String() string {
	return serviceAccountKeyFormat.ToString(*i)
}

func (i *ServiceAccountKeyIdentity) FromExternal(ref string) error {
	parsed, match, err := serviceAccountKeyFormat.Parse(ref)
	if err != nil {
		return fmt.Errorf("format of ServiceAccountKeyIdentity external=%q was not known (use %s): %w", ref, serviceAccountKeyFormat.CanonicalForm(), err)
	}
	if !match {
		return fmt.Errorf("format of ServiceAccountKeyIdentity external=%q was not known (use %s)", ref, serviceAccountKeyFormat.CanonicalForm())
	}
	*i = *parsed
	return nil
}

func (i *ServiceAccountKeyIdentity) Host() string {
	return serviceAccountKeyFormat.Host()
}

func (i *ServiceAccountKeyIdentity) Parent() *ServiceAccountIdentity {
	return &i.ServiceAccountIdentity
}

func (i *ServiceAccountKeyIdentity) ID() string {
	return i.Id
}

type ServiceAccountIdentity struct {
	Project string
	Account string
}

func (p *ServiceAccountIdentity) String() string {
	return serviceAccountFormat.ToString(*p)
}

func (p *ServiceAccountIdentity) FromExternal(ref string) error {
	parsed, match, err := serviceAccountFormat.Parse(ref)
	if err != nil {
		return fmt.Errorf("format of ServiceAccountIdentity external=%q was not known (use %s): %w", ref, serviceAccountFormat.CanonicalForm(), err)
	}
	if !match {
		return fmt.Errorf("format of ServiceAccountIdentity external=%q was not known (use %s)", ref, serviceAccountFormat.CanonicalForm())
	}
	*p = *parsed
	return nil
}

func (p *ServiceAccountIdentity) Host() string {
	return serviceAccountFormat.Host()
}
