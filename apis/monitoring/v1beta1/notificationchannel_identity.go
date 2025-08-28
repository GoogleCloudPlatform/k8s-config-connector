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
	refs "github.com/GoogleCloudPlatform/k8s-config-connector/apis/refs/v1beta1"
	"sigs.k8s.io/controller-runtime/pkg/client"
)

var _ identity.Identity = &NotificationChannelIdentity{}

type NotificationChannelIdentity struct {
	ParentID   *refs.Project
	ResourceID string
}

func (i *NotificationChannelIdentity) String() string {
	return i.ParentID.String() + "/notificationChannels/" + i.ResourceID
}

func (i *NotificationChannelIdentity) FromExternal(ref string) error {
	tokens := strings.Split(ref, "/")
	if len(tokens) == 4 && tokens[3] == "notificationChannels" {

		parentID := &refs.Project{}
		if err := parentID.FromExternal(strings.Join(tokens[:2], "/")); err != nil {
			return fmt.Errorf("format of MonitoringNotificationChannel ref=%q was not known (use %q)", ref, "projects/{projectId}/notificationChannels/{channelID}")
		}

		resourceID := tokens[3]

		i.ParentID = parentID
		i.ResourceID = resourceID

		return nil

	}

	return fmt.Errorf("format of MonitoringNotificationChannel ref=%q was not known (use %q)", ref, "projects/{projectId}/notificationChannels/{channelID}")
}

var _ identity.Resource = &MonitoringNotificationChannel{}

func (obj *MonitoringNotificationChannel) GetIdentity(ctx context.Context, reader client.Reader) (identity.Identity, error) {
	// Get parent ID
	parentID, err := obj.GetParentIdentity(ctx, reader)
	if err != nil {
		return nil, err
	}

	// Get resource ID
	resourceID := common.ValueOf(obj.Spec.ResourceID)
	if resourceID == "" {
		resourceID = obj.GetName()
	}
	if resourceID == "" {
		return nil, fmt.Errorf("cannot resolve resource ID")
	}

	id := &NotificationChannelIdentity{
		ParentID:   parentID.(*refs.Project),
		ResourceID: resourceID,
	}

	// // Attempt to ensure ID is immutable, by verifying against previously-set `status.externalRef`.
	// externalRef := common.ValueOf(obj.Status.ExternalRef)
	// if externalRef != "" {
	// 	previousID := &NotificationChannelIdentity{}
	// 	if err := previousID.FromExternal(externalRef); err != nil {
	// 		return nil, err
	// 	}
	// 	if id.String() != previousID.String() {
	// 		return nil, fmt.Errorf("cannot update MonitoringNotificationChannel identity (old=%q, new=%q): identity is immutable", previousID.String(), id.String())
	// 	}
	// }

	return id, nil
}

func (obj *MonitoringNotificationChannel) GetParentIdentity(ctx context.Context, reader client.Reader) (identity.Identity, error) {
	return refs.ResolveProjectIdentity(ctx, reader, obj)
}
