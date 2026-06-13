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

	apirefs "github.com/GoogleCloudPlatform/k8s-config-connector/apis/refs"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/gcpurls"
)

var ComputeRouterIdentityFormat = gcpurls.Template[ComputeRouterIdentity]("compute.googleapis.com", "projects/{project}/regions/{region}/routers/{router}")

// ComputeRouterIdentity is the identity of a GCP ComputeRouter resource.
// +k8s:deepcopy-gen=false
type ComputeRouterIdentity struct {
	Project string
	Region  string
	Router  string
}

func (i *ComputeRouterIdentity) String() string {
	return ComputeRouterIdentityFormat.ToString(*i)
}

func (i *ComputeRouterIdentity) FromExternal(ref string) error {
	trimmedRef := apirefs.TrimComputeURIPrefix(ref)
	parsed, match, err := ComputeRouterIdentityFormat.Parse(trimmedRef)
	if err != nil {
		return fmt.Errorf("format of ComputeRouter external=%q was not known (use %s): %w", ref, ComputeRouterIdentityFormat.CanonicalForm(), err)
	}
	if !match {
		return fmt.Errorf("format of ComputeRouter external=%q was not known (use %s)", ref, ComputeRouterIdentityFormat.CanonicalForm())
	}

	*i = *parsed
	return nil
}

func (i *ComputeRouterIdentity) Host() string {
	return ComputeRouterIdentityFormat.Host()
}

func (i *ComputeRouterIdentity) ParentString() string {
	return fmt.Sprintf("projects/%s/regions/%s", i.Project, i.Region)
}

func ParseComputeRouterExternal(external string) (*ComputeRouterIdentity, error) {
	if external == "" {
		return nil, fmt.Errorf("empty ComputeRouter external value")
	}
	id := &ComputeRouterIdentity{}
	if err := id.FromExternal(external); err != nil {
		return nil, err
	}
	return id, nil
}
