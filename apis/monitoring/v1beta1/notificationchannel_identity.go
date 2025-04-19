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

// +tool:krm-identity
// proto.service: google.monitoring.v3.NotificationChannelService
// proto.message: google.monitoring.v3.NotificationChannel
// crd.type: MonitoringNotificationChannel
// crd.version: v1beta1

import (
	"fmt"
	"strings"

	"github.com/GoogleCloudPlatform/k8s-config-connector/apis/common/identity"
	"github.com/GoogleCloudPlatform/k8s-config-connector/apis/common/parent"
)

var _ identity.Identity = &MonitoringNotificationChannelIdentity{}

// +k8s:deepcopy-gen=false
type MonitoringNotificationChannelIdentity struct {
	ParentID *parent.ProjectParent
	Node     string
}

func (i *MonitoringNotificationChannelIdentity) String() string {
	return i.ParentID.String() + "/nodes/" + i.Node
}

func (i *MonitoringNotificationChannelIdentity) FromExternal(external string) error {
	tokens := strings.Split(external, "/")
	if len(tokens) == 4 && tokens[0] == "projects" && tokens[4] == "notificationChannels" {
		i.ParentID = &parent.ProjectParent{
			ProjectID: tokens[1],
		}
		i.Node = tokens[5]
		return nil
	}
	return fmt.Errorf("format of MonitoringNotificationChannel external=%q was not known (use projects/{{projectID}}/notificationChannels/{{notificationChannelID}})", external)
}

// var _ identity.Resource = &MonitoringNotificationChannel{}

// func (obj *MonitoringNotificationChannel) GetIdentity(ctx context.Context, reader client.Reader) (identity.Identity, error) {
// 	// Get parent ID
// 	parentID, err := obj.GetParentIdentity(ctx, reader)
// 	if err != nil {
// 		return nil, err
// 	}

// 	// Get desired ID
// 	resourceID := common.ValueOf(obj.Spec.ResourceID)
// 	if resourceID == "" {
// 		resourceID = obj.GetName()
// 	}
// 	if resourceID == "" {
// 		return nil, fmt.Errorf("cannot resolve resource ID")
// 	}

// 	id := &MonitoringNotificationChannelIdentity{
// 		ParentID: parentID.(*parent.ProjectParent),
// 		Node:     resourceID,
// 	}

// 	// Attempt to ensure ID is immutable, by verifying against previously-set `status.externalRef`.
// 	externalRef := common.ValueOf(obj.Status.ExternalRef)
// 	if externalRef != "" {
// 		previousID := &MonitoringNotificationChannelIdentity{}
// 		if err := previousID.FromExternal(externalRef); err != nil {
// 			return nil, err
// 		}
// 		if id.String() != previousID.String() {
// 			return nil, fmt.Errorf("cannot update MonitoringNotificationChannel identity (old=%q, new=%q): identity is immutable", previousID.String(), id.String())
// 		}
// 	}

// 	return id, nil
// }

// func (obj *MonitoringNotificationChannel) GetParentIdentity(ctx context.Context, reader client.Reader) (identity.Identity, error) {
// 	// TODO: Can we extract helper?

// 	// Normalize projectRef
// 	if err := obj.Spec.ProjectRef.Normalize(ctx, reader, obj.GetNamespace()); err != nil {
// 		return nil, err
// 	}

// 	location := obj.Spec.Location

// 	external := obj.Spec.ProjectRef.External + "/locations/" + location

// 	// Get parent identity
// 	parentID := &parent.ProjectAndLocationParent{}
// 	if err := parentID.FromExternal(external); err != nil {
// 		return nil, err
// 	}
// 	return parentID, nil
// }
