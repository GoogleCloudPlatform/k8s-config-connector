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

package privatecarefs

import (
	"fmt"

	"github.com/GoogleCloudPlatform/k8s-config-connector/apis/common/identity"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/gcpurls"
)

var (
	_ identity.IdentityV2 = &PrivateCACAPoolIdentity{}
)

var PrivateCACAPoolIdentityFormat = gcpurls.Template[PrivateCACAPoolIdentity]("privateca.googleapis.com", "projects/{project}/locations/{location}/caPools/{caPool}")

type PrivateCACAPoolIdentity struct {
	Project  string
	Location string
	CAPool   string
}

func (i *PrivateCACAPoolIdentity) String() string {
	return PrivateCACAPoolIdentityFormat.ToString(*i)
}

func (i *PrivateCACAPoolIdentity) FromExternal(ref string) error {
	parsed, match, err := PrivateCACAPoolIdentityFormat.Parse(ref)
	if err != nil {
		return fmt.Errorf("format of PrivateCACAPool external=%q was not known (use %s): %w", ref, PrivateCACAPoolIdentityFormat.CanonicalForm(), err)
	}
	if !match {
		return fmt.Errorf("format of PrivateCACAPool external=%q was not known (use %s)", ref, PrivateCACAPoolIdentityFormat.CanonicalForm())
	}

	*i = *parsed
	return nil
}

func (i *PrivateCACAPoolIdentity) Host() string {
	return PrivateCACAPoolIdentityFormat.Host()
}
