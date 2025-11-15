// Copyright 2024 Google LLC
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

// +tool:controller-client
// proto.service: google.monitoring.dashboard.v1.DashboardsService

package monitoring

import (
	"context"
	"fmt"

	api "cloud.google.com/go/monitoring/apiv3/v2"
	dashboardapi "cloud.google.com/go/monitoring/dashboard/apiv1"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/config"
)

func newDashboardsClient(ctx context.Context, config *config.ControllerConfig) (*dashboardapi.DashboardsClient, error) {
	opts, err := config.RESTClientOptions()
	if err != nil {
		return nil, err
	}
	client, err := dashboardapi.NewDashboardsRESTClient(ctx, opts...)
	if err != nil {
		return nil, fmt.Errorf("building dashboard client: %w", err)
	}
	return client, err
}

func newNotificationChannelsClient(ctx context.Context, config *config.ControllerConfig) (*api.NotificationChannelClient, error) {
	opts, err := config.RESTClientOptions()
	if err != nil {
		return nil, err
	}
	client, err := api.NewNotificationChannelClient(ctx, opts...)
	if err != nil {
		return nil, fmt.Errorf("building notification channel client: %w", err)
	}
	return client, err
}
