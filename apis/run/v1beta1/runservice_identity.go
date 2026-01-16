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

	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/gcpurls"
)

var RunServiceIdentityFormat = gcpurls.Template[RunServiceIdentity]("run.googleapis.com", "projects/{project}/locations/{location}/services/{service}")

type RunServiceIdentity struct {
	Project  string
	Location string
	Service  string
}

func (i *RunServiceIdentity) FromExternal(ref string) error {
	parsed, match, err := RunServiceIdentityFormat.Parse(ref)
	if err != nil {
		return fmt.Errorf("format of RunService external=%q was not known (use %s): %w", ref, RunServiceIdentityFormat.CanonicalForm(), err)
	}
	if !match {
		return fmt.Errorf("format of RunService external=%q was not known (use %s)", ref, RunServiceIdentityFormat.CanonicalForm())
	}

	*i = *parsed
	return nil
}
