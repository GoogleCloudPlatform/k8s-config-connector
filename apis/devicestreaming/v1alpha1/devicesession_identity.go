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
	_ identity.IdentityV2 = &DeviceStreamingSessionIdentity{}
	_ identity.Resource   = &DeviceStreamingSession{}
)

var DeviceStreamingSessionIdentityFormat = gcpurls.Template[DeviceStreamingSessionIdentity]("devicestreaming.googleapis.com", "projects/{project}/deviceSessions/{device_session}")

// +k8s:deepcopy-gen=false
type DeviceStreamingSessionIdentity struct {
	Project       string
	DeviceSession string
}

func (i *DeviceStreamingSessionIdentity) String() string {
	return DeviceStreamingSessionIdentityFormat.ToString(*i)
}

func (i *DeviceStreamingSessionIdentity) FromExternal(ref string) error {
	parsed, match, err := DeviceStreamingSessionIdentityFormat.Parse(ref)
	if err != nil {
		return fmt.Errorf("format of DeviceStreamingSession external=%q was not known (use %s): %w", ref, DeviceStreamingSessionIdentityFormat.CanonicalForm(), err)
	}
	if !match {
		return fmt.Errorf("format of DeviceStreamingSession external=%q was not known (use %s)", ref, DeviceStreamingSessionIdentityFormat.CanonicalForm())
	}

	*i = *parsed
	return nil
}

func (i *DeviceStreamingSessionIdentity) Host() string {
	return DeviceStreamingSessionIdentityFormat.Host()
}

func getIdentityFromDeviceStreamingSessionSpec(ctx context.Context, reader client.Reader, obj client.Object) (*DeviceStreamingSessionIdentity, error) {
	_, ok := obj.(*DeviceStreamingSession)
	if !ok {
		return nil, fmt.Errorf("object is not a DeviceStreamingSession")
	}
	resourceID, err := refs.GetResourceID(obj)
	if err != nil {
		return nil, fmt.Errorf("cannot resolve resource ID: %w", err)
	}

	projectID, err := refs.ResolveProjectID(ctx, reader, obj)
	if err != nil {
		return nil, fmt.Errorf("cannot resolve project: %w", err)
	}

	identity := &DeviceStreamingSessionIdentity{
		Project:       projectID,
		DeviceSession: resourceID,
	}
	return identity, nil
}

func (obj *DeviceStreamingSession) GetIdentity(ctx context.Context, reader client.Reader) (identity.Identity, error) {
	specIdentity, err := getIdentityFromDeviceStreamingSessionSpec(ctx, reader, obj)
	if err != nil {
		return nil, err
	}

	// Cross-check the identity against the status value, if present.
	externalRef := common.ValueOf(obj.Status.ExternalRef)
	if externalRef != "" {
		// Validate desired with actual
		statusIdentity := &DeviceStreamingSessionIdentity{}
		if err := statusIdentity.FromExternal(externalRef); err != nil {
			return nil, err
		}

		if statusIdentity.String() != specIdentity.String() {
			return nil, fmt.Errorf("cannot change DeviceStreamingSession identity (old=%q, new=%q)", statusIdentity.String(), specIdentity.String())
		}
	}

	return specIdentity, nil
}

// ExternalIdentifier implements the identity.ExternalIdentifier interface.
func (c *DeviceStreamingSession) ExternalIdentifier() *string {
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
