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

package iamrefs

import (
	"fmt"
	"strings"

	"github.com/GoogleCloudPlatform/k8s-config-connector/apis/common/identity"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/gcpurls"
)

var (
	_ identity.IdentityV2 = &IAMServiceAccountIdentity{}
)

var IAMServiceAccountIdentityFormat = gcpurls.Template[IAMServiceAccountIdentity](
	"iam.googleapis.com",
	"projects/{project}/serviceAccounts/{account}",
)

// IAMServiceAccountIdentity is the identity of a GCP IAMServiceAccount resource.
// +k8s:deepcopy-gen=true
type IAMServiceAccountIdentity struct {
	Project string
	Account string
}

func (i *IAMServiceAccountIdentity) String() string {
	return IAMServiceAccountIdentityFormat.ToString(*i)
}

func (i *IAMServiceAccountIdentity) FromExternal(ref string) error {
	parsed, match, err := IAMServiceAccountIdentityFormat.Parse(ref)
	if err != nil {
		return fmt.Errorf("format of IAMServiceAccount external=%q was not known (use %s): %w", ref, IAMServiceAccountIdentityFormat.CanonicalForm(), err)
	}
	if !match {
		return fmt.Errorf("format of IAMServiceAccount external=%q was not known (use %s)", ref, IAMServiceAccountIdentityFormat.CanonicalForm())
	}

	if strings.Contains(parsed.Account, "@") {
		return fmt.Errorf("format of IAMServiceAccount external=%q was not known (use %s): email format is not allowed in identity", ref, IAMServiceAccountIdentityFormat.CanonicalForm())
	}

	*i = *parsed
	return nil
}

func (i *IAMServiceAccountIdentity) Host() string {
	return IAMServiceAccountIdentityFormat.Host()
}

func (i *IAMServiceAccountIdentity) ParentString() string {
	return "projects/" + i.Project
}
