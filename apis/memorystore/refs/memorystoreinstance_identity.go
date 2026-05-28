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

package refs

import (
	"fmt"

	"github.com/GoogleCloudPlatform/k8s-config-connector/apis/common/identity"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/gcpurls"
)

var (
	_ identity.IdentityV2 = &MemorystoreInstanceIdentity{}
)

var MemorystoreInstanceIdentityFormat = gcpurls.Template[MemorystoreInstanceIdentity]("memorystore.googleapis.com", "projects/{project}/locations/{location}/instances/{instance}")

// +k8s:deepcopy-gen=false
type MemorystoreInstanceIdentity struct {
	Project  string
	Location string
	Instance string
}

func (i *MemorystoreInstanceIdentity) String() string {
	return MemorystoreInstanceIdentityFormat.ToString(*i)
}

func (i *MemorystoreInstanceIdentity) FromExternal(ref string) error {
	parsed, match, err := MemorystoreInstanceIdentityFormat.Parse(ref)
	if err != nil {
		return fmt.Errorf("format of MemorystoreInstance external=%q was not known (use %s): %w", ref, MemorystoreInstanceIdentityFormat.CanonicalForm(), err)
	}
	if !match {
		return fmt.Errorf("format of MemorystoreInstance external=%q was not known (use %s)", ref, MemorystoreInstanceIdentityFormat.CanonicalForm())
	}

	*i = *parsed
	return nil
}

func (i *MemorystoreInstanceIdentity) Host() string {
	return MemorystoreInstanceIdentityFormat.Host()
}
