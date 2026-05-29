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

	"github.com/GoogleCloudPlatform/k8s-config-connector/apis/common"
	"github.com/GoogleCloudPlatform/k8s-config-connector/apis/common/identity"
	refs "github.com/GoogleCloudPlatform/k8s-config-connector/apis/refs/v1beta1"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/gcpurls"
	"sigs.k8s.io/controller-runtime/pkg/client"
)

var (
	_ identity.IdentityV2 = &MonitoringNotificationChannelIdentity{}
	_ identity.Resource   = &MonitoringNotificationChannel{}
)

var MonitoringNotificationChannelIdentityFormat = gcpurls.Template[MonitoringNotificationChannelIdentity]("monitoring.googleapis.com", "projects/{project}/notificationChannels/{notificationchannel}")

// +k8s:deepcopy-gen=false
type MonitoringNotificationChannelIdentity struct {
	Project             string
	NotificationChannel string
}

func (i *MonitoringNotificationChannelIdentity) String() string {
	return MonitoringNotificationChannelIdentityFormat.ToString(*i)
}

func (i *MonitoringNotificationChannelIdentity) FromExternal(ref string) error {
	parsed, match, err := MonitoringNotificationChannelIdentityFormat.Parse(ref)
	if err != nil {
		return fmt.Errorf("format of MonitoringNotificationChannel external=%q was not known (use %s): %w", ref, MonitoringNotificationChannelIdentityFormat.CanonicalForm(), err)
	}
	if !match {
		return fmt.Errorf("format of MonitoringNotificationChannel external=%q was not known (use %s)", ref, MonitoringNotificationChannelIdentityFormat.CanonicalForm())
	}

	*i = *parsed
	return nil
}

func (i *MonitoringNotificationChannelIdentity) Host() string {
	return MonitoringNotificationChannelIdentityFormat.Host()
}

func getIdentityFromMonitoringNotificationChannelSpec(ctx context.Context, reader client.Reader, obj client.Object) (*MonitoringNotificationChannelIdentity, error) {
	resourceID, err := refs.GetResourceID(obj)
	if err != nil {
		return nil, fmt.Errorf("cannot resolve resource ID")
	}

	projectID, err := refs.ResolveProjectID(ctx, reader, obj)
	if err != nil {
		return nil, fmt.Errorf("cannot resolve project")
	}

	identity := &MonitoringNotificationChannelIdentity{
		Project:             projectID,
		NotificationChannel: resourceID,
	}
	return identity, nil
}

func (obj *MonitoringNotificationChannel) GetIdentity(ctx context.Context, reader client.Reader) (identity.Identity, error) {
	specIdentity, err := getIdentityFromMonitoringNotificationChannelSpec(ctx, reader, obj)
	if err != nil {
		return nil, err
	}

	// Cross-check the identity against the status value, if present.
	externalRef := common.ValueOf(obj.Status.Name)
	if externalRef != "" {
		// Validate desired with actual
		statusIdentity := &MonitoringNotificationChannelIdentity{}
		if err := statusIdentity.FromExternal(externalRef); err != nil {
			return nil, err
		}

		if statusIdentity.String() != specIdentity.String() {
			return nil, fmt.Errorf("cannot change MonitoringNotificationChannel identity (old=%q, new=%q)", statusIdentity.String(), specIdentity.String())
		}
	}

	return specIdentity, nil
}
