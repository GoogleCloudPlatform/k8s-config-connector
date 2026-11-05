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
	dataformv1beta1 "github.com/GoogleCloudPlatform/k8s-config-connector/apis/dataform/v1beta1"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/gcpurls"
)

var (
	_ identity.IdentityV2 = &DataformCompilationResultIdentity{}
)

var DataformCompilationResultIdentityFormat = gcpurls.Template[DataformCompilationResultIdentity]("dataform.googleapis.com", "projects/{project}/locations/{location}/repositories/{repository}/compilationResults/{compilationResult}")

// DataformCompilationResultIdentity is the identity of a GCP DataformCompilationResult resource.
// +k8s:deepcopy-gen=false
type DataformCompilationResultIdentity struct {
	Project           string
	Location          string
	Repository        string
	CompilationResult string
}

func (i *DataformCompilationResultIdentity) String() string {
	return DataformCompilationResultIdentityFormat.ToString(*i)
}

func (i *DataformCompilationResultIdentity) FromExternal(ref string) error {
	parsed, match, err := DataformCompilationResultIdentityFormat.Parse(ref)
	if err != nil {
		return fmt.Errorf("format of DataformCompilationResult external=%q was not known (use %s): %w", ref, DataformCompilationResultIdentityFormat.CanonicalForm(), err)
	}
	if !match {
		return fmt.Errorf("format of DataformCompilationResult external=%q was not known (use %s)", ref, DataformCompilationResultIdentityFormat.CanonicalForm())
	}

	*i = *parsed
	return nil
}

func (i *DataformCompilationResultIdentity) Host() string {
	return DataformCompilationResultIdentityFormat.Host()
}

func (i *DataformCompilationResultIdentity) ParentString() string {
	parent := &dataformv1beta1.DataformRepositoryIdentity{
		Project:    i.Project,
		Location:   i.Location,
		Repository: i.Repository,
	}
	return parent.String()
}
