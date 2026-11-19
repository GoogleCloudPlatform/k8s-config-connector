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
	"fmt"

	"github.com/GoogleCloudPlatform/k8s-config-connector/apis/common/identity"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/gcpurls"
)

var (
	_ identity.IdentityV2 = &CesAppVersionIdentity{}
)

var CesAppVersionIdentityFormat = gcpurls.Template[CesAppVersionIdentity]("ces.googleapis.com", "projects/{project}/locations/{location}/apps/{app}/versions/{version}")

// CesAppVersionIdentity is the identity of a GCP CesAppVersion resource.
// +k8s:deepcopy-gen=false
type CesAppVersionIdentity struct {
	Project  string
	Location string
	App      string
	Version  string
}

func (i *CesAppVersionIdentity) String() string {
	return CesAppVersionIdentityFormat.ToString(*i)
}

func (i *CesAppVersionIdentity) FromExternal(ref string) error {
	parsed, match, err := CesAppVersionIdentityFormat.Parse(ref)
	if err != nil {
		return fmt.Errorf("format of CesAppVersion external=%q was not known (use %s): %w", ref, CesAppVersionIdentityFormat.CanonicalForm(), err)
	}
	if !match {
		return fmt.Errorf("format of CesAppVersion external=%q was not known (use %s)", ref, CesAppVersionIdentityFormat.CanonicalForm())
	}

	*i = *parsed
	return nil
}

func (i *CesAppVersionIdentity) Host() string {
	return CesAppVersionIdentityFormat.Host()
}

func (i *CesAppVersionIdentity) ParentString() string {
	parent := &CESAppIdentity{
		Project:  i.Project,
		Location: i.Location,
		App:      i.App,
	}
	return parent.String()
}
