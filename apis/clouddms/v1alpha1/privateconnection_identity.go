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

	"github.com/GoogleCloudPlatform/k8s-config-connector/apis/common"
	"github.com/GoogleCloudPlatform/k8s-config-connector/apis/common/identity"
	refs "github.com/GoogleCloudPlatform/k8s-config-connector/apis/refs/v1beta1"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/gcpurls"
	"k8s.io/apimachinery/pkg/apis/meta/v1/unstructured"
	"sigs.k8s.io/controller-runtime/pkg/client"
)

var (
	_ identity.IdentityV2 = &CloudDMSPrivateConnectionIdentity{}
	_ identity.Resource   = &CloudDMSPrivateConnection{}
)

var CloudDMSPrivateConnectionIdentityFormat = gcpurls.Template[CloudDMSPrivateConnectionIdentity]("datamigration.googleapis.com", "projects/{project}/locations/{location}/privateConnections/{privateconnection}")

// +k8s:deepcopy-gen=false
type CloudDMSPrivateConnectionIdentity struct {
	Project           string
	Location          string
	PrivateConnection string
}

func (i *CloudDMSPrivateConnectionIdentity) String() string {
	return CloudDMSPrivateConnectionIdentityFormat.ToString(*i)
}

func (i *CloudDMSPrivateConnectionIdentity) FromExternal(ref string) error {
	parsed, match, err := CloudDMSPrivateConnectionIdentityFormat.Parse(ref)
	if err != nil {
		return fmt.Errorf("format of CloudDMSPrivateConnection external=%q was not known (use %s): %w", ref, CloudDMSPrivateConnectionIdentityFormat.CanonicalForm(), err)
	}
	if !match {
		return fmt.Errorf("format of CloudDMSPrivateConnection external=%q was not known (use %s)", ref, CloudDMSPrivateConnectionIdentityFormat.CanonicalForm())
	}

	*i = *parsed
	return nil
}

func (i *CloudDMSPrivateConnectionIdentity) Host() string {
	return CloudDMSPrivateConnectionIdentityFormat.Host()
}

func getIdentityFromCloudDMSPrivateConnectionSpec(ctx context.Context, reader client.Reader, obj client.Object) (*CloudDMSPrivateConnectionIdentity, error) {
	var resourceID, location, projectID string

	u, ok := obj.(*unstructured.Unstructured)
	if ok {
		resourceID, _, _ = unstructured.NestedString(u.Object, "spec", "resourceID")
		if resourceID == "" {
			resourceID = u.GetName()
		}
		location, _, _ = unstructured.NestedString(u.Object, "spec", "location")
		projectRefName, _, _ := unstructured.NestedString(u.Object, "spec", "projectRef", "name")
		projectRefNamespace, _, _ := unstructured.NestedString(u.Object, "spec", "projectRef", "namespace")
		projectRefExternal, _, _ := unstructured.NestedString(u.Object, "spec", "projectRef", "external")
		projectRef := &refs.ProjectRef{
			Name:      projectRefName,
			Namespace: projectRefNamespace,
			External:  projectRefExternal,
		}
		resolvedProject, err := refs.ResolveProject(ctx, reader, u.GetNamespace(), projectRef)
		if err != nil {
			return nil, err
		}
		projectID = resolvedProject.ProjectID
	} else {
		typedObj, ok := obj.(*CloudDMSPrivateConnection)
		if !ok {
			return nil, fmt.Errorf("unexpected type %T", obj)
		}
		resourceID = common.ValueOf(typedObj.Spec.ResourceID)
		if resourceID == "" {
			resourceID = typedObj.GetName()
		}
		if typedObj.Spec.Parent != nil {
			location = typedObj.Spec.Parent.Location
			resolvedProject, err := refs.ResolveProject(ctx, reader, typedObj.GetNamespace(), typedObj.Spec.Parent.ProjectRef)
			if err != nil {
				return nil, err
			}
			projectID = resolvedProject.ProjectID
		}
	}

	if resourceID == "" {
		return nil, fmt.Errorf("cannot resolve resource ID")
	}
	if location == "" {
		return nil, fmt.Errorf("cannot resolve location")
	}
	if projectID == "" {
		return nil, fmt.Errorf("cannot resolve project")
	}

	return &CloudDMSPrivateConnectionIdentity{
		Project:           projectID,
		Location:          location,
		PrivateConnection: resourceID,
	}, nil
}

func (obj *CloudDMSPrivateConnection) GetIdentity(ctx context.Context, reader client.Reader) (identity.Identity, error) {
	specIdentity, err := getIdentityFromCloudDMSPrivateConnectionSpec(ctx, reader, obj)
	if err != nil {
		return nil, err
	}

	// Cross-check the identity against the status value, if present.
	externalRef := common.ValueOf(obj.Status.ExternalRef)
	if externalRef != "" {
		// Validate desired with actual
		statusIdentity := &CloudDMSPrivateConnectionIdentity{}
		if err := statusIdentity.FromExternal(externalRef); err != nil {
			return nil, err
		}

		if statusIdentity.String() != specIdentity.String() {
			return nil, fmt.Errorf("cannot change CloudDMSPrivateConnection identity (old=%q, new=%q)", statusIdentity.String(), specIdentity.String())
		}
	}

	return specIdentity, nil
}
