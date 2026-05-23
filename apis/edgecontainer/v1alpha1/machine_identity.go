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

package v1alpha1

import (
	"context"
	"fmt"

	"github.com/GoogleCloudPlatform/k8s-config-connector/apis/common"
	"github.com/GoogleCloudPlatform/k8s-config-connector/apis/common/identity"
	refs "github.com/GoogleCloudPlatform/k8s-config-connector/apis/refs/v1beta1"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/gcpurls"
	"k8s.io/apimachinery/pkg/apis/meta/v1/unstructured"
	"sigs.k8s.io/controller-runtime/pkg/client"
)

var (
	_ identity.IdentityV2 = &EdgeContainerMachineIdentity{}
	_ identity.Resource   = &EdgeContainerMachine{}
)

var EdgeContainerMachineIdentityFormat = gcpurls.Template[EdgeContainerMachineIdentity]("edgecontainer.googleapis.com", "projects/{project}/locations/{location}/machines/{machine}")

// +k8s:deepcopy-gen=false
type EdgeContainerMachineIdentity struct {
	Project  string
	Location string
	Machine  string
}

func (i *EdgeContainerMachineIdentity) String() string {
	return EdgeContainerMachineIdentityFormat.ToString(*i)
}

func (i *EdgeContainerMachineIdentity) FromExternal(ref string) error {
	parsed, match, err := EdgeContainerMachineIdentityFormat.Parse(ref)
	if err != nil {
		return fmt.Errorf("format of EdgeContainerMachine external=%q was not known (use %s): %w", ref, EdgeContainerMachineIdentityFormat.CanonicalForm(), err)
	}
	if !match {
		return fmt.Errorf("format of EdgeContainerMachine external=%q was not known (use %s)", ref, EdgeContainerMachineIdentityFormat.CanonicalForm())
	}

	*i = *parsed
	return nil
}

func (i *EdgeContainerMachineIdentity) Host() string {
	return EdgeContainerMachineIdentityFormat.Host()
}

func getIdentityFromEdgeContainerMachineSpec(ctx context.Context, reader client.Reader, obj client.Object) (*EdgeContainerMachineIdentity, error) {
	u, ok := obj.(*EdgeContainerMachine)
	var resourceID string
	var location string
	var projectRef *refs.ProjectRef
	var err error

	if ok {
		resourceID = common.ValueOf(u.Spec.ResourceID)
		if resourceID == "" {
			resourceID = u.GetName()
		}
		location = u.Spec.Parent.Location
		projectRef = u.Spec.Parent.ProjectRef
	} else {
		// handle unstructured
		unstruct, ok := obj.(*unstructured.Unstructured)
		if !ok {
			return nil, fmt.Errorf("object is not a EdgeContainerMachine or Unstructured")
		}
		resourceID, _, err = unstructured.NestedString(unstruct.Object, "spec", "resourceID")
		if err != nil {
			return nil, err
		}
		if resourceID == "" {
			resourceID = unstruct.GetName()
		}
		location, _, err = unstructured.NestedString(unstruct.Object, "spec", "location")
		if err != nil {
			return nil, err
		}
		// Try resolving projectRef
		projRefMap, found, err := unstructured.NestedMap(unstruct.Object, "spec", "projectRef")
		if err != nil {
			return nil, err
		}
		if found {
			projectRef = &refs.ProjectRef{}
			if name, ok := projRefMap["name"].(string); ok {
				projectRef.Name = name
			}
			if namespace, ok := projRefMap["namespace"].(string); ok {
				projectRef.Namespace = namespace
			}
			if external, ok := projRefMap["external"].(string); ok {
				projectRef.External = external
			}
		}
	}

	if resourceID == "" {
		return nil, fmt.Errorf("cannot resolve resource ID")
	}
	if location == "" {
		return nil, fmt.Errorf("cannot resolve location")
	}

	projectRefResolved, err := refs.ResolveProject(ctx, reader, obj.GetNamespace(), projectRef)
	if err != nil {
		return nil, err
	}
	projectID := projectRefResolved.ProjectID
	if projectID == "" {
		return nil, fmt.Errorf("cannot resolve project")
	}

	identity := &EdgeContainerMachineIdentity{
		Project:  projectID,
		Location: location,
		Machine:  resourceID,
	}
	return identity, nil
}

func (obj *EdgeContainerMachine) GetIdentity(ctx context.Context, reader client.Reader) (identity.Identity, error) {
	specIdentity, err := getIdentityFromEdgeContainerMachineSpec(ctx, reader, obj)
	if err != nil {
		return nil, err
	}

	externalRef := common.ValueOf(obj.Status.ExternalRef)
	if externalRef != "" {
		statusIdentity := &EdgeContainerMachineIdentity{}
		if err := statusIdentity.FromExternal(externalRef); err != nil {
			return nil, err
		}

		if statusIdentity.String() != specIdentity.String() {
			return nil, fmt.Errorf("cannot change EdgeContainerMachine identity (old=%q, new=%q)", statusIdentity.String(), specIdentity.String())
		}
	}

	return specIdentity, nil
}
