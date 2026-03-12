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

package firestore

import (
	"context"
	"fmt"

	api "cloud.google.com/go/firestore/apiv1"
	adminapi "cloud.google.com/go/firestore/apiv1/admin"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/config"
)

func newFirestoreAdminClient(ctx context.Context, config *config.ControllerConfig) (*adminapi.FirestoreAdminClient, error) {
	opts, err := config.RESTClientOptions()
	if err != nil {
		return nil, err
	}
	client, err := adminapi.NewFirestoreAdminRESTClient(ctx, opts...)
	if err != nil {
		return nil, fmt.Errorf("building firestore admin client: %w", err)
	}
	return client, err
}

func newFirestoreClient(ctx context.Context, config *config.ControllerConfig) (*api.Client, error) {
	opts, err := config.RESTClientOptions()
	if err != nil {
		return nil, err
	}
	client, err := api.NewRESTClient(ctx, opts...)
	if err != nil {
		return nil, fmt.Errorf("building firestore client: %w", err)
	}
	return client, err
}
