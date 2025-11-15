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

package v1beta1

import (
	"context"
	"fmt"
	"strings"

	"github.com/GoogleCloudPlatform/k8s-config-connector/apis/common"
	"github.com/GoogleCloudPlatform/k8s-config-connector/apis/common/identity"
	"github.com/GoogleCloudPlatform/k8s-config-connector/apis/common/parent"
	refsv1beta1 "github.com/GoogleCloudPlatform/k8s-config-connector/apis/refs/v1beta1"
	"sigs.k8s.io/controller-runtime/pkg/client"
)

const (
	// MonitoringNotificationChannelIdentityURL is the format for the externalRef of a MonitoringNotificationChannel.
	MonitoringNotificationChannelIdentityURL = "projects/{{project}}/notificationChannels/{{notificationChannel}}"

	ServicePrefix = "//monitoring.googleapis.com/"
)

var _ identity.Identity = &MonitoringNotificationChannelIdentity{}

// MonitoringNotificationChannelIdentity represents the identity of a MonitoringNotificationChannel.
// +k8s:deepcopy-gen=false
type MonitoringNotificationChannelIdentity struct {
	Parent              *parent.ProjectParent
	NotificationChannel string
}

func (i *MonitoringNotificationChannelIdentity) String() string {
	return "notificationChannels/" + i.NotificationChannel
}

func (i *MonitoringNotificationChannelIdentity) FromExternal(ref string) error {
	// Should be able to parse https://docs.cloud.google.com/asset-inventory/docs/asset-names
	ref = strings.TrimPrefix(ref, ServicePrefix)

	tokens := strings.Split(ref, "/")
	if len(tokens) == 4 && tokens[0] == "projects" && tokens[2] == "notificationChannels" {
		i.Parent = &parent.ProjectParent{}
		if err := i.Parent.FromExternal(strings.Join(tokens[0:2], "/")); err != nil {
			return fmt.Errorf("cannot parse project from external=%q: %w", ref, err)
		}
		i.NotificationChannel = tokens[3]
		if i.NotificationChannel == "" {
			return fmt.Errorf("notificationChannel was empty in external=%q", ref)
		}
		return nil
	}

	return fmt.Errorf("format of NotificationChannel external=%q was not known (use %s)", ref, MonitoringNotificationChannelIdentityURL)
}

var _ identity.Resource = &MonitoringNotificationChannel{}

func (obj *MonitoringNotificationChannel) GetIdentity(ctx context.Context, reader client.Reader) (identity.Identity, error) {
	// Get desired resource ID
	resourceID := common.ValueOf(obj.Spec.ResourceID)

	// Server-generated ID; do not fallback to name
	// if resourceID == "" {
	// 	resourceID = obj.GetName()
	// }

	if resourceID == "" {
		return nil, fmt.Errorf("cannot resolve resource ID")
	}

	projectID, err := refsv1beta1.ResolveProjectID(ctx, reader, obj)
	if err != nil {
		return nil, fmt.Errorf("cannot resolve project: %w", err)
	}

	newIdentity := &MonitoringNotificationChannelIdentity{
		Parent:              &parent.ProjectParent{ProjectID: projectID},
		NotificationChannel: resourceID,
	}

	// Validate against the ID stored in status.name
	externalRef := common.ValueOf(obj.Status.Name)
	if externalRef != "" {
		statusIdentity := &MonitoringNotificationChannelIdentity{}
		if err := statusIdentity.FromExternal(externalRef); err != nil {
			return nil, fmt.Errorf("cannot parse existing status.name=%q: %w", externalRef, err)
		}
		if statusIdentity.String() != newIdentity.String() {
			return nil, fmt.Errorf("existing status.name=%q does not match the identity resolved from spec: %q", externalRef, newIdentity.String())
		}
	}

	return newIdentity, nil
}
