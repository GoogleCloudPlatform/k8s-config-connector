// Copyright 2022 Google LLC
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

package dcl

import (
	"fmt"

	"k8s.io/apimachinery/pkg/runtime"

	dclextension "github.com/GoogleCloudPlatform/k8s-config-connector/pkg/dcl/extension"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/k8s"

	"github.com/nasa9084/go-openapi"
	"k8s.io/apimachinery/pkg/apis/meta/v1/unstructured"
)

// Resource is a wrapper around k8s.Resource and adds information regarding its
// corresponding DCL resource and maintains an original copy of the
// k8s.Resource.
type Resource struct {
	k8s.Resource `json:",inline"`

	Original *k8s.Resource `json:"-"`

	// Fields related to DCL processing
	Schema *openapi.Schema `json:"-"`
}

func NewResource(u *unstructured.Unstructured, schema *openapi.Schema) (*Resource, error) {
	resource := &Resource{
		Schema: schema,
	}

	r, err := k8s.NewResource(u)
	if err != nil {
		return nil, err
	}
	resource.Resource = *r

	// Intentionally re-create the K8s resource to create a separate copy.
	resource.Original, err = k8s.NewResource(u)
	if err != nil {
		return nil, err
	}

	if err := resource.ValidateResourceIDIfSupported(); err != nil {
		return nil, err
	}
	return resource, nil
}

// DeepCopyObject is needed to implement the interface of client.Object.
func (r *Resource) DeepCopyObject() runtime.Object {
	panic("unexpected call to resource.DeepCopyObject(...)")
}

func (r *Resource) ValidateResourceIDIfSupported() error {
	// The resource ID is represented by the top level 'name' field consistently in DCL.
	_, found := dclextension.GetNameFieldSchema(r.Schema)
	if !found {
		// The resource doesn't have a 'resourceID' field.
		return nil
	}

	_, err := r.IsResourceIDConfigured()
	if err != nil {
		return fmt.Errorf("error validating '%s' field: %w", k8s.ResourceIDFieldPath, err)
	}
	return nil
}

func (r *Resource) HasServerGeneratedIDButNotConfigured() (bool, error) {
	s, found := dclextension.GetNameFieldSchema(r.Schema)
	if !found {
		// The resource doesn't have a 'resourceID' field.
		return false, nil
	}

	isServerGenerated, isConfigured, err := r.getResourceIDMetadata(s)
	if err != nil {
		return false, err
	}
	return isServerGenerated && !isConfigured, nil
}

func (r *Resource) HasServerGeneratedIDAndConfigured() (bool, error) {
	s, found := dclextension.GetNameFieldSchema(r.Schema)
	if !found {
		// The resource doesn't have a 'resourceID' field.
		return false, nil
	}
	isServerGenerated, isConfigured, err := r.getResourceIDMetadata(s)
	if err != nil {
		return false, err
	}
	return isServerGenerated && isConfigured, nil
}

func (r *Resource) getResourceIDMetadata(s *openapi.Schema) (isServerGenerated bool, isConfigured bool, err error) {
	isConfigured, err = r.IsResourceIDConfigured()
	if err != nil {
		return false, false, fmt.Errorf("error checking if '%s' field is configured: %w", k8s.ResourceIDFieldPath, err)
	}
	isServerGenerated, err = dclextension.IsResourceIDFieldServerGenerated(s)
	if err != nil {
		return false, false, fmt.Errorf("error parsing `resourceID` field schema: %w", err)
	}
	return isServerGenerated, isConfigured, nil
}

func (r *Resource) HasMutableButUnreadableFields() (bool, error) {
	// The resource uses the state-hint directive if it contains any mutable but
	// unreadable fields. The state-hint directive is only used by mutable but
	// unreadable fields.
	hasStateHint, err := dclextension.HasStateHint(r.Schema)
	if err != nil {
		return false, err
	}
	return hasStateHint, nil
}
