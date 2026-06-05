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
	_ identity.IdentityV2 = &DeviceStreamingDeviceSessionIdentity{}
	_ identity.Resource   = &DeviceStreamingDeviceSession{}
)

var DeviceStreamingDeviceSessionIdentityFormat = gcpurls.Template[DeviceStreamingDeviceSessionIdentity]("devicestreaming.googleapis.com", "projects/{project}/deviceSessions/{devicesession}")

// +k8s:deepcopy-gen=false
type DeviceStreamingDeviceSessionIdentity struct {
	Project       string
	DeviceSession string
}

func (i *DeviceStreamingDeviceSessionIdentity) String() string {
	return DeviceStreamingDeviceSessionIdentityFormat.ToString(*i)
}

func (i *DeviceStreamingDeviceSessionIdentity) FromExternal(ref string) error {
	parsed, match, err := DeviceStreamingDeviceSessionIdentityFormat.Parse(ref)
	if err != nil {
		return fmt.Errorf("format of DeviceStreamingDeviceSession external=%q was not known (use %s): %w", ref, DeviceStreamingDeviceSessionIdentityFormat.CanonicalForm(), err)
	}
	if !match {
		return fmt.Errorf("format of DeviceStreamingDeviceSession external=%q was not known (use %s)", ref, DeviceStreamingDeviceSessionIdentityFormat.CanonicalForm())
	}

	*i = *parsed
	return nil
}

func (i *DeviceStreamingDeviceSessionIdentity) Host() string {
	return DeviceStreamingDeviceSessionIdentityFormat.Host()
}

func getIdentityFromDeviceStreamingDeviceSessionSpec(ctx context.Context, reader client.Reader, obj client.Object) (*DeviceStreamingDeviceSessionIdentity, error) {
	_, ok := obj.(*DeviceStreamingDeviceSession)
	if !ok {
		return nil, fmt.Errorf("object is not a DeviceStreamingDeviceSession")
	}
	resourceID, err := refs.GetResourceID(obj)
	if err != nil {
		return nil, fmt.Errorf("cannot resolve resource ID: %w", err)
	}

	projectID, err := refs.ResolveProjectID(ctx, reader, obj)
	if err != nil {
		return nil, fmt.Errorf("cannot resolve project: %w", err)
	}

	identity := &DeviceStreamingDeviceSessionIdentity{
		Project:       projectID,
		DeviceSession: resourceID,
	}
	return identity, nil
}

func (obj *DeviceStreamingDeviceSession) GetIdentity(ctx context.Context, reader client.Reader) (identity.Identity, error) {
	specIdentity, err := getIdentityFromDeviceStreamingDeviceSessionSpec(ctx, reader, obj)
	if err != nil {
		return nil, err
	}

	// Cross-check the identity against the status value, if present.
	externalRef := common.ValueOf(obj.Status.ExternalRef)
	if externalRef != "" {
		// Validate desired with actual
		statusIdentity := &DeviceStreamingDeviceSessionIdentity{}
		if err := statusIdentity.FromExternal(externalRef); err != nil {
			return nil, err
		}

		if statusIdentity.String() != specIdentity.String() {
			return nil, fmt.Errorf("cannot change DeviceStreamingDeviceSession identity (old=%q, new=%q)", statusIdentity.String(), specIdentity.String())
		}
	}

	return specIdentity, nil
}

// ExternalIdentifier implements the identity.ExternalIdentifier interface.
func (c *DeviceStreamingDeviceSession) ExternalIdentifier() *string {
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
