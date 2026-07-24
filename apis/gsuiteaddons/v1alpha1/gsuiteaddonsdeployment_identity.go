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
	"context"
	"fmt"

	"github.com/GoogleCloudPlatform/k8s-config-connector/apis/common/identity"
	refs "github.com/GoogleCloudPlatform/k8s-config-connector/apis/refs/v1beta1"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/gcpurls"
	"sigs.k8s.io/controller-runtime/pkg/client"
)

var (
	_ identity.IdentityV2 = &GSuiteAddonsDeploymentIdentity{}
	_ identity.Resource   = &GSuiteAddonsDeployment{}
)

var GSuiteAddonsDeploymentIdentityFormat = gcpurls.Template[GSuiteAddonsDeploymentIdentity]("gsuiteaddons.googleapis.com", "projects/{project}/deployments/{deployment}")

// +k8s:deepcopy-gen=false

// GSuiteAddonsDeploymentIdentity is the identity of a GCP GSuiteAddonsDeployment resource.
type GSuiteAddonsDeploymentIdentity struct {
	Project    string
	Deployment string
}

func (i *GSuiteAddonsDeploymentIdentity) String() string {
	return GSuiteAddonsDeploymentIdentityFormat.ToString(*i)
}

func (i *GSuiteAddonsDeploymentIdentity) FromExternal(ref string) error {
	parsed, match, err := GSuiteAddonsDeploymentIdentityFormat.Parse(ref)
	if err != nil {
		return fmt.Errorf("format of GSuiteAddonsDeployment external=%q was not known (use %s): %w", ref, GSuiteAddonsDeploymentIdentityFormat.CanonicalForm(), err)
	}
	if !match {
		return fmt.Errorf("format of GSuiteAddonsDeployment external=%q was not known (use %s)", ref, GSuiteAddonsDeploymentIdentityFormat.CanonicalForm())
	}

	*i = *parsed
	return nil
}

func (i *GSuiteAddonsDeploymentIdentity) Host() string {
	return GSuiteAddonsDeploymentIdentityFormat.Host()
}

func (i *GSuiteAddonsDeploymentIdentity) ParentString() string {
	return "projects/" + i.Project
}

func getIdentityFromGSuiteAddonsDeploymentSpec(ctx context.Context, reader client.Reader, obj *GSuiteAddonsDeployment) (*GSuiteAddonsDeploymentIdentity, error) {
	resourceID, err := refs.GetResourceID(obj)
	if err != nil {
		return nil, fmt.Errorf("cannot resolve resource ID: %w", err)
	}

	projectID, err := refs.ResolveProjectID(ctx, reader, obj)
	if err != nil {
		return nil, fmt.Errorf("cannot resolve project: %w", err)
	}

	identity := &GSuiteAddonsDeploymentIdentity{
		Project:    projectID,
		Deployment: resourceID,
	}
	return identity, nil
}

func (obj *GSuiteAddonsDeployment) GetIdentity(ctx context.Context, reader client.Reader) (identity.Identity, error) {
	specIdentity, err := getIdentityFromGSuiteAddonsDeploymentSpec(ctx, reader, obj)
	if err != nil {
		return nil, err
	}

	return specIdentity, nil
}
