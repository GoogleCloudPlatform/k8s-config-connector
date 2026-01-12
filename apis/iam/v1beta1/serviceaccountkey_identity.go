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
	"context"
	"fmt"
	"strings"

	refsv1beta1 "github.com/GoogleCloudPlatform/k8s-config-connector/apis/refs/v1beta1"
	"k8s.io/apimachinery/pkg/apis/meta/v1/unstructured"
	"sigs.k8s.io/controller-runtime/pkg/client"
)

type ServiceAccountKeyIdentity struct {
	parent *ServiceAccountKeyParent
	id     string
}

func (i *ServiceAccountKeyIdentity) String() string {
	return i.parent.String() + "/keys/" + i.id
}

func (i *ServiceAccountKeyIdentity) Parent() *ServiceAccountKeyParent {
	return i.parent
}

func (i *ServiceAccountKeyIdentity) ID() string {
	return i.id
}

type ServiceAccountKeyParent struct {
	ProjectID        string
	ServiceAccountID string
}

func (p *ServiceAccountKeyParent) String() string {
	return "projects/" + p.ProjectID + "/serviceAccounts/" + p.ServiceAccountID
}

// NewServiceAccountKeyIdentity builds a ServiceAccountKeyIdentity from the resource spec.
func NewServiceAccountKeyIdentity(ctx context.Context, reader client.Reader, obj *IAMServiceAccountKey, u *unstructured.Unstructured) (*ServiceAccountKeyIdentity, error) {
	// Get ProjectID
	projectID, err := refsv1beta1.ResolveProjectID(ctx, reader, u)
	if err != nil {
		return nil, err
	}

	// Get ServiceAccountID from the reference
	if obj.Spec.ServiceAccountRef == nil {
		return nil, fmt.Errorf("spec.serviceAccountRef is required")
	}

	// Resolve the service account reference
	serviceAccountRef := *obj.Spec.ServiceAccountRef
	if err := serviceAccountRef.Resolve(ctx, reader, obj); err != nil {
		return nil, fmt.Errorf("cannot resolve spec.serviceAccountRef: %w", err)
	}

	// The resolved External should be an email address
	serviceAccountID := serviceAccountRef.External
	if serviceAccountID == "" {
		return nil, fmt.Errorf("serviceAccountRef resolved to empty external")
	}

	parent := &ServiceAccountKeyParent{
		ProjectID:        projectID,
		ServiceAccountID: serviceAccountID,
	}

	// The key ID comes from status.name after creation
	// For now, we cannot determine the ID until the resource is created
	keyID := ""
	if obj.Status.Name != nil && *obj.Status.Name != "" {
		// Extract key ID from the full name
		// status.name format: projects/{project}/serviceAccounts/{account}/keys/{key}
		if parsed, err := ParseServiceAccountKeyExternal(*obj.Status.Name); err == nil {
			keyID = parsed.id
		}
	}

	return &ServiceAccountKeyIdentity{
		parent: parent,
		id:     keyID,
	}, nil
}

// ParseServiceAccountKeyExternal parses a service account key external reference.
// Supports formats:
// - projects/{project}/serviceAccounts/{account}/keys/{key}
// - //iam.googleapis.com/projects/{project}/serviceAccounts/{account}/keys/{key}
func ParseServiceAccountKeyExternal(external string) (*ServiceAccountKeyIdentity, error) {
	if external == "" {
		return nil, fmt.Errorf("missing external value")
	}

	external = strings.TrimPrefix(external, "//iam.googleapis.com/")
	external = strings.TrimPrefix(external, "/")
	tokens := strings.Split(external, "/")

	if len(tokens) != 6 || tokens[0] != "projects" || tokens[2] != "serviceAccounts" || tokens[4] != "keys" {
		return nil, fmt.Errorf("format of IAMServiceAccountKey external=%q was not known (use projects/{{projectID}}/serviceAccounts/{{serviceAccountID}}/keys/{{keyID}})", external)
	}

	return &ServiceAccountKeyIdentity{
		parent: &ServiceAccountKeyParent{
			ProjectID:        tokens[1],
			ServiceAccountID: tokens[3],
		},
		id: tokens[5],
	}, nil
}
