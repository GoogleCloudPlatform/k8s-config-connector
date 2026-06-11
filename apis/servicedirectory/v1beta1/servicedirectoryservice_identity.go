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

	"github.com/GoogleCloudPlatform/k8s-config-connector/apis/common"
	"github.com/GoogleCloudPlatform/k8s-config-connector/apis/common/identity"
	refs "github.com/GoogleCloudPlatform/k8s-config-connector/apis/refs/v1beta1"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/gcpurls"
	"sigs.k8s.io/controller-runtime/pkg/client"
)

var (
	_ identity.IdentityV2 = &ServiceDirectoryServiceIdentity{}
	_ identity.Resource   = &ServiceDirectoryService{}
)

var ServiceDirectoryServiceIdentityFormat = gcpurls.Template[ServiceDirectoryServiceIdentity]("servicedirectory.googleapis.com", "projects/{project}/locations/{location}/namespaces/{namespace}/services/{service}")

// +k8s:deepcopy-gen=false

// ServiceDirectoryServiceIdentity is the identity of a GCP ServiceDirectoryService resource.
type ServiceDirectoryServiceIdentity struct {
	Project   string
	Location  string
	Namespace string
	Service   string
}

func (i *ServiceDirectoryServiceIdentity) String() string {
	return ServiceDirectoryServiceIdentityFormat.ToString(*i)
}

func (i *ServiceDirectoryServiceIdentity) ParentString() string {
	return fmt.Sprintf("projects/%s/locations/%s/namespaces/%s", i.Project, i.Location, i.Namespace)
}

func (i *ServiceDirectoryServiceIdentity) FromExternal(ref string) error {
	ref = strings.TrimPrefix(ref, "https://servicedirectory.googleapis.com/v1/")
	ref = strings.TrimPrefix(ref, "https://servicedirectory.googleapis.com/v1beta1/")
	ref = strings.TrimPrefix(ref, "https://servicedirectory.googleapis.com/")

	parsed, match, err := ServiceDirectoryServiceIdentityFormat.Parse(ref)
	if err != nil {
		return fmt.Errorf("format of ServiceDirectoryService external=%q was not known (use %s): %w", ref, ServiceDirectoryServiceIdentityFormat.CanonicalForm(), err)
	}
	if !match {
		return fmt.Errorf("format of ServiceDirectoryService external=%q was not known (use %s)", ref, ServiceDirectoryServiceIdentityFormat.CanonicalForm())
	}

	*i = *parsed
	return nil
}

func (i *ServiceDirectoryServiceIdentity) Host() string {
	return ServiceDirectoryServiceIdentityFormat.Host()
}

func getIdentityFromServiceDirectoryServiceSpec(ctx context.Context, reader client.Reader, obj *ServiceDirectoryService) (*ServiceDirectoryServiceIdentity, error) {
	namespaceRef := obj.Spec.NamespaceRef
	if err := (&namespaceRef).Normalize(ctx, reader, obj.GetNamespace()); err != nil {
		return nil, fmt.Errorf("resolving spec.namespaceRef: %w", err)
	}

	namespaceIdentity := &ServiceDirectoryNamespaceIdentity{}
	if err := namespaceIdentity.FromExternal(namespaceRef.GetExternal()); err != nil {
		return nil, fmt.Errorf("parsing namespaceRef.external=%q: %w", namespaceRef.GetExternal(), err)
	}

	resourceID, err := refs.GetResourceID(obj)
	if err != nil {
		return nil, fmt.Errorf("cannot resolve resource ID: %w", err)
	}

	identity := &ServiceDirectoryServiceIdentity{
		Project:   namespaceIdentity.Project,
		Location:  namespaceIdentity.Location,
		Namespace: namespaceIdentity.Namespace,
		Service:   resourceID,
	}
	return identity, nil
}

func (obj *ServiceDirectoryService) GetIdentity(ctx context.Context, reader client.Reader) (identity.Identity, error) {
	specIdentity, err := getIdentityFromServiceDirectoryServiceSpec(ctx, reader, obj)
	if err != nil {
		return nil, err
	}

	// Cross-check the identity against the status value, if present.
	externalRef := common.ValueOf(obj.Status.Name)
	if externalRef != "" {
		// Validate desired with actual
		statusIdentity := &ServiceDirectoryServiceIdentity{}
		if err := statusIdentity.FromExternal(externalRef); err != nil {
			return nil, err
		}

		if statusIdentity.String() != specIdentity.String() {
			return nil, fmt.Errorf("cannot change ServiceDirectoryService identity (old=%q, new=%q)", statusIdentity.String(), specIdentity.String())
		}
	}

	return specIdentity, nil
}
