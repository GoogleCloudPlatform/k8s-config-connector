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
	_ identity.IdentityV2 = &ObservabilityDatasetIdentity{}
)

var ObservabilityDatasetIdentityFormat = gcpurls.Template[ObservabilityDatasetIdentity]("observability.googleapis.com", "projects/{project}/locations/{location}/buckets/{bucket}/datasets/{dataset}")

// ObservabilityDatasetIdentity is the identity of a GCP ObservabilityDataset resource.
// +k8s:deepcopy-gen=false
type ObservabilityDatasetIdentity struct {
	Project  string
	Location string
	Bucket   string
	Dataset  string
}

func (i *ObservabilityDatasetIdentity) String() string {
	return ObservabilityDatasetIdentityFormat.ToString(*i)
}

func (i *ObservabilityDatasetIdentity) FromExternal(ref string) error {
	parsed, match, err := ObservabilityDatasetIdentityFormat.Parse(ref)
	if err != nil {
		return fmt.Errorf("format of ObservabilityDataset external=%q was not known (use %s): %w", ref, ObservabilityDatasetIdentityFormat.CanonicalForm(), err)
	}
	if !match {
		return fmt.Errorf("format of ObservabilityDataset external=%q was not known (use %s)", ref, ObservabilityDatasetIdentityFormat.CanonicalForm())
	}

	*i = *parsed
	return nil
}

func (i *ObservabilityDatasetIdentity) Host() string {
	return ObservabilityDatasetIdentityFormat.Host()
}

func (i *ObservabilityDatasetIdentity) ParentString() string {
	return fmt.Sprintf("projects/%s/locations/%s/buckets/%s", i.Project, i.Location, i.Bucket)
}
