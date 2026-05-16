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

	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/gcpurls"
)

var InstanceIdentityFormat = gcpurls.Template[InstanceIdentity]("memorystore.googleapis.com", "projects/{project}/locations/{location}/instances/{instance}")

type InstanceIdentity struct {
	Project  string
	Location string
	Instance string
}

func (i *InstanceIdentity) FromExternal(ref string) error {
	parsed, match, err := InstanceIdentityFormat.Parse(ref)
	if err != nil {
		return fmt.Errorf("format of Instance external=%q was not known (use %s): %w", ref, InstanceIdentityFormat.CanonicalForm(), err)
	}
	if !match {
		return fmt.Errorf("format of Instance external=%q was not known (use %s)", ref, InstanceIdentityFormat.CanonicalForm())
	}

	*i = *parsed
	return nil
}
