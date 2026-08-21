// Copyright 2024 Google LLC
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
	refsv1beta1 "github.com/GoogleCloudPlatform/k8s-config-connector/apis/refs/v1beta1"
	"k8s.io/apimachinery/pkg/apis/meta/v1/unstructured"
	"sigs.k8s.io/controller-runtime/pkg/client"
)

type SecretVersionIdentity struct {
	id                      string
	parent                  *SecretVersionParent
	serviceGeneratedIDKnown *bool
}

// HasKnownID tells whether Config Connector knows the resource identity.
// If not, Config Connector saves one GCP GET call, and starts the CREATE call directly.
// This is mostly for GCP services that do not allow user to specify ID, but assign an ID when creating the object.
func (i *SecretVersionIdentity) HasKnownID() bool {
	return *i.serviceGeneratedIDKnown
}

func (i *SecretVersionIdentity) String() string {
	return i.parent.String() + "/versions/" + i.id
}

func (r *SecretVersionIdentity) Parent() *SecretVersionParent {
	return r.parent
}

func (r *SecretVersionIdentity) ID() string {
	return r.id
}

type SecretVersionParent struct {
	ProjectID string
	SecretID  string
}

func (p *SecretVersionParent) String() string {
	return "projects/" + p.ProjectID + "/secrets/" + p.SecretID
}

func NewSecretVersionIdentity(ctx context.Context, reader client.Reader, obj *SecretManagerSecretVersion, u *unstructured.Unstructured) (*SecretVersionIdentity, error) {
	// Get Parent
	projectID, err := refsv1beta1.ResolveProjectID(ctx, reader, u)
	if err != nil {
		return nil, err
	}
	secretExternal, err := obj.Spec.SecretRef.NormalizedExternal(ctx, reader, obj.GetNamespace())
	if err != nil {
		return nil, err
	}
	secretIdentity, err := ParseSecretExternal(secretExternal)
	if err != nil {
		return nil, err
	}
	secretID := secretIdentity.ID()

	// If `spec.resourceID` is not empty, it means user wants to acquire the object.
	desiredVersionID := common.ValueOf(obj.Spec.ResourceID)

	externalRef := common.ValueOf(obj.Status.ExternalRef)
	if externalRef != "" {
		actualIdentity, err := ParseSecretVersionExternal(externalRef)
		if err != nil {
			return nil, err
		}
		if actualIdentity.parent.ProjectID != projectID {
			return nil, fmt.Errorf("spec.projectRef changed, expect %s, got %s", actualIdentity.parent.ProjectID, projectID)
		}
		if actualIdentity.parent.SecretID != secretID {
			return nil, fmt.Errorf("spec.projectRef changed, expect %s, got %s", actualIdentity.parent.ProjectID, projectID)
		}
		if desiredVersionID != "" && actualIdentity.id != desiredVersionID {
			return nil, fmt.Errorf("cannot reset `metadata.name` or `spec.resourceID` to %s, since it has already assigned to %s",
				desiredVersionID, actualIdentity.id)
		}
		desiredVersionID = actualIdentity.id
	}

	known := false
	if externalRef != "" {
		known = true
	}
	if desiredVersionID != "" {
		known = true
	}
	return &SecretVersionIdentity{
		parent: &SecretVersionParent{
			ProjectID: projectID,
			SecretID:  secretID,
		},
		id:                      desiredVersionID,
		serviceGeneratedIDKnown: &known,
	}, nil
}
