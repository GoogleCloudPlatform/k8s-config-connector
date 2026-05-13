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
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/config"
	pb "github.com/GoogleCloudPlatform/k8s-config-connector/pkg/gcpclients/generated/google/firestore/admin/v1"
	"google.golang.org/api/transport/grpc"
	"google.golang.org/genproto/googleapis/longrunning"
)

func newFirestoreAdminClient(ctx context.Context, config *config.ControllerConfig) (pb.FirestoreAdminClient, error) {
	opts, err := config.GRPCClientOptions()
	if err != nil {
		return nil, err
	}
	conn, err := grpc.Dial(ctx, opts...)
	if err != nil {
		return nil, err
	}
	return pb.NewFirestoreAdminClient(conn), nil
}

func newOperationsClient(ctx context.Context, config *config.ControllerConfig) (longrunning.OperationsClient, error) {
	opts, err := config.GRPCClientOptions()
	if err != nil {
		return nil, err
	}
	conn, err := grpc.Dial(ctx, opts...)
	if err != nil {
		return nil, err
	}
	return longrunning.NewOperationsClient(conn), nil
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
