// Copyright 2025 Google LLC
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
	"context"
	"fmt"

	"github.com/GoogleCloudPlatform/k8s-config-connector/apis/common"
	"github.com/GoogleCloudPlatform/k8s-config-connector/apis/common/identity"
	"github.com/GoogleCloudPlatform/k8s-config-connector/apis/common/parent"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/gcpurls"
	"sigs.k8s.io/controller-runtime/pkg/client"
)

var _ identity.Identity = &EnvironmentIdentity{}

// EnvironmentIdentity defines the resource reference to ComposerEnvironment, which "External" field
// holds the GCP identifier for the KRM object.
type EnvironmentIdentity struct {
	parent *parent.ProjectAndLocationParent
	id     string
}

func (i *EnvironmentIdentity) String() string {
	return i.parent.String() + "/environments/" + i.id
}

func (i *EnvironmentIdentity) ID() string {
	return i.id
}

func (i *EnvironmentIdentity) Parent() *parent.ProjectAndLocationParent {
	return i.parent
}

type environmentExternal struct {
	Project     string
	Location    string
	Environment string
}

var environmentURLTemplate = gcpurls.Template[environmentExternal](
	"composer.googleapis.com",
	"projects/{project}/locations/{location}/environments/{environment}",
)

func (i *EnvironmentIdentity) FromExternal(ref string) error {
	external, matches, err := environmentURLTemplate.Parse(ref)
	if err != nil {
		return err
	}
	if !matches {
		return fmt.Errorf("external %q does not match format %q", ref, environmentURLTemplate.CanonicalForm())
	}
	i.parent = &parent.ProjectAndLocationParent{
		ProjectID: external.Project,
		Location:  external.Location,
	}
	i.id = external.Environment
	return nil
}

var _ identity.Resource = &ComposerEnvironment{}

func (obj *ComposerEnvironment) GetIdentity(ctx context.Context, reader client.Reader) (identity.Identity, error) {
	environment := &EnvironmentIdentity{
		parent: &parent.ProjectAndLocationParent{},
	}

	// Resolve user-configured Parent
	if obj.Spec.ParentRef == nil {
		return nil, fmt.Errorf("spec.parentRef is required")
	}
	if err := obj.Spec.ParentRef.Build(ctx, reader, obj.GetNamespace(), environment.parent); err != nil {
		return nil, err
	}

	// Get user-configured ID
	environment.id = common.ValueOf(obj.Spec.ResourceID)
	if environment.id == "" {
		environment.id = obj.GetName()
	}
	if environment.id == "" {
		return nil, fmt.Errorf("cannot resolve resource ID")
	}

	// Validate against the ID stored in status.externalRef, if any
	externalRef := common.ValueOf(obj.Status.ExternalRef)
	if externalRef != "" {
		statusIdentity := &EnvironmentIdentity{}
		if err := statusIdentity.FromExternal(externalRef); err != nil {
			return nil, fmt.Errorf("cannot parse existing status.externalRef %q: %w", externalRef, err)
		}
		if statusIdentity.String() != environment.String() {
			return nil, fmt.Errorf("existing status.externalRef %q does not match the identity resolved from spec: %q. "+
				"The resource might have been moved or renamed in GCP, or the spec might have been changed to point to a different resource. "+
				"Please verify the spec and the actual resource in GCP.", externalRef, environment.String())
		}
	}
	return environment, nil
}
