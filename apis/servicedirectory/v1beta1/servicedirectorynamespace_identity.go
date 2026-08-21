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
	_ identity.IdentityV2 = &ServiceDirectoryNamespaceIdentity{}
	_ identity.Resource   = &ServiceDirectoryNamespace{}
)

var ServiceDirectoryNamespaceIdentityFormat = gcpurls.Template[ServiceDirectoryNamespaceIdentity]("servicedirectory.googleapis.com", "projects/{project}/locations/{location}/namespaces/{namespace}")

// +k8s:deepcopy-gen=false

// ServiceDirectoryNamespaceIdentity is the identity of a GCP ServiceDirectoryNamespace resource.
type ServiceDirectoryNamespaceIdentity struct {
	Project   string
	Location  string
	Namespace string
}

func (i *ServiceDirectoryNamespaceIdentity) String() string {
	return ServiceDirectoryNamespaceIdentityFormat.ToString(*i)
}

func (i *ServiceDirectoryNamespaceIdentity) ParentString() string {
	return fmt.Sprintf("projects/%s/locations/%s", i.Project, i.Location)
}

func (i *ServiceDirectoryNamespaceIdentity) FromExternal(ref string) error {
	ref = strings.TrimPrefix(ref, "https://servicedirectory.googleapis.com/v1/")
	ref = strings.TrimPrefix(ref, "https://servicedirectory.googleapis.com/v1beta1/")
	ref = strings.TrimPrefix(ref, "https://servicedirectory.googleapis.com/")

	parsed, match, err := ServiceDirectoryNamespaceIdentityFormat.Parse(ref)
	if err != nil {
		return fmt.Errorf("format of ServiceDirectoryNamespace external=%q was not known (use %s): %w", ref, ServiceDirectoryNamespaceIdentityFormat.CanonicalForm(), err)
	}
	if !match {
		return fmt.Errorf("format of ServiceDirectoryNamespace external=%q was not known (use %s)", ref, ServiceDirectoryNamespaceIdentityFormat.CanonicalForm())
	}

	*i = *parsed
	return nil
}

func (i *ServiceDirectoryNamespaceIdentity) Host() string {
	return ServiceDirectoryNamespaceIdentityFormat.Host()
}

func getIdentityFromServiceDirectoryNamespaceSpec(ctx context.Context, reader client.Reader, obj *ServiceDirectoryNamespace) (*ServiceDirectoryNamespaceIdentity, error) {
	resourceID, err := refs.GetResourceID(obj)
	if err != nil {
		return nil, fmt.Errorf("cannot resolve resource ID: %w", err)
	}

	location, err := refs.GetLocation(obj)
	if err != nil {
		return nil, fmt.Errorf("cannot resolve location: %w", err)
	}

	projectID, err := refs.ResolveProjectID(ctx, reader, obj)
	if err != nil {
		return nil, fmt.Errorf("cannot resolve project: %w", err)
	}

	identity := &ServiceDirectoryNamespaceIdentity{
		Project:   projectID,
		Location:  location,
		Namespace: resourceID,
	}
	return identity, nil
}

func (obj *ServiceDirectoryNamespace) GetIdentity(ctx context.Context, reader client.Reader) (identity.Identity, error) {
	specIdentity, err := getIdentityFromServiceDirectoryNamespaceSpec(ctx, reader, obj)
	if err != nil {
		return nil, err
	}

	// Cross-check the identity against the status value, if present.
	externalRef := common.ValueOf(obj.Status.Name)
	if externalRef != "" {
		// Validate desired with actual
		statusIdentity := &ServiceDirectoryNamespaceIdentity{}
		if err := statusIdentity.FromExternal(externalRef); err != nil {
			return nil, err
		}

		if statusIdentity.String() != specIdentity.String() {
			return nil, fmt.Errorf("cannot change ServiceDirectoryNamespace identity (old=%q, new=%q)", statusIdentity.String(), specIdentity.String())
		}
	}

	return specIdentity, nil
}
