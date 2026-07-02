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
	"sigs.k8s.io/controller-runtime/pkg/client"
)

var (
	_ identity.IdentityV2 = &TestingDeviceSessionIdentity{}
	_ identity.Resource   = &TestingDeviceSession{}
)

var TestingDeviceSessionIdentityFormat = gcpurls.Template[TestingDeviceSessionIdentity]("testing.googleapis.com", "projects/{project}/deviceSessions/{session}")

// +k8s:deepcopy-gen=false
type TestingDeviceSessionIdentity struct {
	Project string
	Session string
}

func (i *TestingDeviceSessionIdentity) String() string {
	return TestingDeviceSessionIdentityFormat.ToString(*i)
}

func (i *TestingDeviceSessionIdentity) ParentString() string {
	return "projects/" + i.Project
}

func (i *TestingDeviceSessionIdentity) FromExternal(ref string) error {
	parsed, match, err := TestingDeviceSessionIdentityFormat.Parse(ref)
	if err != nil {
		return fmt.Errorf("format of TestingDeviceSession external=%q was not known (use %s): %w", ref, TestingDeviceSessionIdentityFormat.CanonicalForm(), err)
	}
	if !match {
		return fmt.Errorf("format of TestingDeviceSession external=%q was not known (use %s)", ref, TestingDeviceSessionIdentityFormat.CanonicalForm())
	}

	*i = *parsed
	return nil
}

func (i *TestingDeviceSessionIdentity) Host() string {
	return TestingDeviceSessionIdentityFormat.Host()
}

func getIdentityFromTestingDeviceSessionSpec(ctx context.Context, reader client.Reader, obj client.Object) (*TestingDeviceSessionIdentity, error) {
	_, ok := obj.(*TestingDeviceSession)
	if !ok {
		return nil, fmt.Errorf("object is not a TestingDeviceSession")
	}
	resourceID, err := refs.GetResourceID(obj)
	if err != nil {
		return nil, fmt.Errorf("cannot resolve resource ID: %w", err)
	}

	projectID, err := refs.ResolveProjectID(ctx, reader, obj)
	if err != nil {
		return nil, fmt.Errorf("cannot resolve project: %w", err)
	}

	identity := &TestingDeviceSessionIdentity{
		Project: projectID,
		Session: resourceID,
	}
	return identity, nil
}

func (obj *TestingDeviceSession) GetIdentity(ctx context.Context, reader client.Reader) (identity.Identity, error) {
	specIdentity, err := getIdentityFromTestingDeviceSessionSpec(ctx, reader, obj)
	if err != nil {
		return nil, err
	}

	// Cross-check the identity against the status value, if present.
	externalRef := common.ValueOf(obj.Status.ExternalRef)
	if externalRef != "" {
		// Validate desired with actual
		statusIdentity := &TestingDeviceSessionIdentity{}
		if err := statusIdentity.FromExternal(externalRef); err != nil {
			return nil, err
		}

		if statusIdentity.String() != specIdentity.String() {
			return nil, fmt.Errorf("cannot change TestingDeviceSession identity (old=%q, new=%q)", statusIdentity.String(), specIdentity.String())
		}
	}

	return specIdentity, nil
}

// ExternalIdentifier implements the identity.ExternalIdentifier interface.
func (c *TestingDeviceSession) ExternalIdentifier() *string {
	if c.Status.ExternalRef != nil {
		return c.Status.ExternalRef
	}
	id := c.GetName()
	if c.Spec.ResourceID != nil {
		id = *c.Spec.ResourceID
	}
	if c.Spec.ProjectRef != nil && c.Spec.ProjectRef.External != "" {
		s := "projects/" + c.Spec.ProjectRef.External + "/deviceSessions/" + id
		return &s
	}
	return nil
}
