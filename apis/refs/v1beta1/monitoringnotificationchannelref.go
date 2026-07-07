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

	apierrors "k8s.io/apimachinery/pkg/api/errors"
	"k8s.io/apimachinery/pkg/apis/meta/v1/unstructured"
	"k8s.io/apimachinery/pkg/runtime/schema"
	"k8s.io/apimachinery/pkg/types"
	"sigs.k8s.io/controller-runtime/pkg/client"
)

type MonitoringNotificationChannelRef struct {
	// Name of the referenced object.
	// +optional
	Name string `json:"name,omitempty"`

	// Namespace of the referenced object.
	// +optional
	Namespace string `json:"namespace,omitempty"`

	// The MonitoringNotificationChannel resource name, when not managed by Config Connector.
	// For example: projects/{project}/notificationChannels/{channel_id}
	// +optional
	External string `json:"external,omitempty"`
}

type MonitoringNotificationChannel struct {
	Project   string
	ChannelID string
}

func (c *MonitoringNotificationChannel) String() string {
	return fmt.Sprintf("projects/%s/notificationChannels/%s", c.Project, c.ChannelID)
}

func ResolveMonitoringNotificationChannelRef(ctx context.Context, reader client.Reader, obj client.Object, ref *MonitoringNotificationChannelRef) (*MonitoringNotificationChannel, error) {
	if ref == nil {
		return nil, nil
	}

	if ref.Name == "" && ref.External == "" {
		return nil, fmt.Errorf("must specify either name or external on notificationChannelRef")
	}
	if ref.External != "" && ref.Name != "" {
		return nil, fmt.Errorf("cannot specify both name and external")
	}

	if ref.External != "" {
		// External is already the full projects/{project}/notificationChannels/{channel_id} path
		// We can parse it or just return it.
		// Wait, let's parse it to ensure it is in correct format.
		var project, channelID string
		if n, err := fmt.Sscanf(ref.External, "projects/%s/notificationChannels/%s", &project, &channelID); err != nil || n != 2 {
			// Sscanf might fail because %s matches greedily. Let's split using strings.Split.
			tokens := strings.Split(ref.External, "/")
			if len(tokens) == 4 && tokens[0] == "projects" && tokens[2] == "notificationChannels" {
				project = tokens[1]
				channelID = tokens[3]
			} else {
				return nil, fmt.Errorf("format of external notification channel %q was not known", ref.External)
			}
		} else {
			project = strings.TrimSuffix(project, "/notificationChannels")
		}
		return &MonitoringNotificationChannel{
			Project:   project,
			ChannelID: channelID,
		}, nil
	}

	key := types.NamespacedName{
		Namespace: ref.Namespace,
		Name:      ref.Name,
	}
	if key.Namespace == "" {
		key.Namespace = obj.GetNamespace()
	}

	channel := &unstructured.Unstructured{}
	channel.SetGroupVersionKind(schema.GroupVersionKind{
		Group:   "monitoring.cnrm.cloud.google.com",
		Version: "v1beta1",
		Kind:    "MonitoringNotificationChannel",
	})
	if err := reader.Get(ctx, key, channel); err != nil {
		if apierrors.IsNotFound(err) {
			return nil, fmt.Errorf("referenced MonitoringNotificationChannel %v not found", key)
		}
		return nil, fmt.Errorf("error reading referenced MonitoringNotificationChannel %v: %w", key, err)
	}

	resourceID, _, err := unstructured.NestedString(channel.Object, "spec", "resourceID")
	if err != nil {
		return nil, fmt.Errorf("reading spec.resourceID from MonitoringNotificationChannel %v: %w", key, err)
	}
	if resourceID == "" {
		resourceID = channel.GetName()
	}

	projectID, err := ResolveProjectID(ctx, reader, channel)
	if err != nil {
		return nil, err
	}

	return &MonitoringNotificationChannel{
		Project:   projectID,
		ChannelID: resourceID,
	}, nil
}
