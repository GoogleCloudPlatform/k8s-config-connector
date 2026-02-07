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

	"github.com/GoogleCloudPlatform/k8s-config-connector/apis/common/identity"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/gcpurls"
)

var (
	_ identity.IdentityV2 = &ServiceAccountKeyIdentity{}
	_ identity.IdentityV2 = &ServiceAccountKeyParent{}

	serviceAccountKeyFormat = gcpurls.Template[ServiceAccountKeyIdentity](
		"iam.googleapis.com",
		"projects/{ServiceAccountKeyParent.project}/serviceAccounts/{ServiceAccountKeyParent.account}/keys/{id}",
	)
	serviceAccountFormat = gcpurls.Template[ServiceAccountKeyParent](
		"iam.googleapis.com",
		"projects/{project}/serviceAccounts/{account}",
	)
)

type ServiceAccountKeyIdentity struct {
	ServiceAccountKeyParent
	Id string
}

func (i *ServiceAccountKeyIdentity) String() string {
	return serviceAccountKeyFormat.ToString(*i)
}

func (i *ServiceAccountKeyIdentity) FromExternal(ref string) error {
	parsed, match, err := serviceAccountKeyFormat.Parse(ref)
	if err != nil {
		return fmt.Errorf("format of ServiceAccountKeyIdentity external=%q was not known (use %s): %w", ref, serviceAccountKeyFormat.CanonicalForm(), err)
	}
	if !match {
		return fmt.Errorf("format of ServiceAccountKeyIdentity external=%q was not known (use %s)", ref, serviceAccountKeyFormat.CanonicalForm())
	}
	*i = *parsed
	return nil
}

func (i *ServiceAccountKeyIdentity) Host() string {
	return serviceAccountKeyFormat.Host()
}

func (i *ServiceAccountKeyIdentity) Parent() *ServiceAccountKeyParent {
	return &i.ServiceAccountKeyParent
}

func (i *ServiceAccountKeyIdentity) ID() string {
	return i.Id
}

/*
	func getIdentityFromServiceAccountKeySpec(ctx context.Context, reader client.Reader, obj client.Object) (*ServiceAccountKeyIdentity, error) {
		resourceID, err := refsv1beta1.GetResourceID(obj)
		if err != nil {
			return nil, fmt.Errorf("cannot resolve resource ID")
		}

		location, err := refsv1beta1.GetLocation(obj)
		if err != nil {
			return nil, fmt.Errorf("cannot resolve resource ID")
		}

		projectID, err := refsv1beta1.ResolveProjectID(ctx, reader, obj)
		if err != nil {
			return nil, fmt.Errorf("cannot resolve project")
		}

		identity := &ArtifactRegistryRepositoryIdentity{
			Project:    projectID,
			Location:   location,
			Repository: resourceID,
		}
		return identity, nil
	}
*/
type ServiceAccountKeyParent struct {
	Project string
	Account string
}

func (p *ServiceAccountKeyParent) String() string {
	return serviceAccountFormat.ToString(*p)
}

func (p *ServiceAccountKeyParent) FromExternal(ref string) error {
	parsed, match, err := serviceAccountFormat.Parse(ref)
	if err != nil {
		return fmt.Errorf("format of ServiceAccountKeyParent external=%q was not known (use %s): %w", ref, serviceAccountFormat.CanonicalForm(), err)
	}
	if !match {
		return fmt.Errorf("format of ServiceAccountKeyParent external=%q was not known (use %s)", ref, serviceAccountFormat.CanonicalForm())
	}
	*p = *parsed
	return nil
}

func (p *ServiceAccountKeyParent) Host() string {
	return serviceAccountFormat.Host()
}

// NewServiceAccountKeyIdentity builds a ServiceAccountKeyIdentity from the resource spec.
/*
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
*/

// ParseServiceAccountKeyExternal parses a service account key external reference.
// Supports formats:
// - projects/{project}/serviceAccounts/{account}/keys/{key}
// - //iam.googleapis.com/projects/{project}/serviceAccounts/{account}/keys/{key}
/*
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
*/
