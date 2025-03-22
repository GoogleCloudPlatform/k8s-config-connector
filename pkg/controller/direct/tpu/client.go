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

// +tool:controller-client
// proto.service: google.cloud.tpu.v2.Tpu

package tpu

import (
	"context"
	"fmt"
	"math"

	"cloud.google.com/go/longrunning"
	lroauto "cloud.google.com/go/longrunning/autogen"
	"cloud.google.com/go/longrunning/autogen/longrunningpb"
	tpuv2 "cloud.google.com/go/tpu/apiv2/tpupb"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/config"
	"google.golang.org/api/option"
	"google.golang.org/api/option/internaloption"
	gtransport "google.golang.org/api/transport/grpc"
	"google.golang.org/grpc"
)

type gcpClient struct {
	config config.ControllerConfig
}

func newGCPClient(ctx context.Context, config *config.ControllerConfig) (*gcpClient, error) {
	gcpClient := &gcpClient{
		config: *config,
	}
	return gcpClient, nil
}

func (m *gcpClient) newClient(ctx context.Context) (*TPUV2Client, error) {
	opts, err := m.config.GRPCClientOptions()
	if err != nil {
		return nil, err
	}
	client, err := NewTPUV2ClientGRPC(ctx, opts...)
	if err != nil {
		return nil, fmt.Errorf("building tpu client: %w", err)
	}
	return client, err
}

// When we publish the v2 API, we should switch to using it

type TPUV2Client struct {
	tpuv2.TpuClient
	lroClient *lroauto.OperationsClient
}

func NewTPUV2ClientGRPC(ctx context.Context, opts ...option.ClientOption) (*TPUV2Client, error) {
	clientOpts := defaultGRPCClientOptions()

	connPool, err := gtransport.DialPool(ctx, append(clientOpts, opts...)...)
	if err != nil {
		return nil, fmt.Errorf("building grpc connection pool: %w", err)
	}

	lroClient, err := lroauto.NewOperationsClient(ctx, gtransport.WithConnPool(connPool))
	if err != nil {
		// This error "should not happen", since we are just reusing old connection pool
		// and never actually need to dial.
		// If this does happen, we could leak connp. However, we cannot close conn:
		// If the user invoked the constructor with option.WithGRPCConn,
		// we would close a connection that's still in use.
		// TODO: investigate error conditions.
		return nil, fmt.Errorf("building lro client: %w", err)
	}

	return &TPUV2Client{
		TpuClient: tpuv2.NewTpuClient(connPool),
		lroClient: lroClient,
	}, nil
}

func (c *TPUV2Client) WaitForLRO(ctx context.Context, op *longrunningpb.Operation) error {
	lro := longrunning.InternalNewOperation(c.lroClient, op)
	if err := lro.Wait(ctx, nil); err != nil {
		return err
	}
	return nil
}

func defaultGRPCClientOptions() []option.ClientOption {
	return []option.ClientOption{
		internaloption.WithDefaultEndpoint("tpu.googleapis.com:443"),
		internaloption.WithDefaultEndpointTemplate("tpu.UNIVERSE_DOMAIN:443"),
		internaloption.WithDefaultMTLSEndpoint("tpu.mtls.googleapis.com:443"),
		internaloption.WithDefaultUniverseDomain("googleapis.com"),
		internaloption.WithDefaultAudience("https://tpu.googleapis.com/"),
		internaloption.WithDefaultScopes(DefaultAuthScopes()...),
		internaloption.EnableJwtWithScope(),
		internaloption.EnableNewAuthLibrary(),
		option.WithGRPCDialOption(grpc.WithDefaultCallOptions(
			grpc.MaxCallRecvMsgSize(math.MaxInt32))),
	}
}

// DefaultAuthScopes reports the default set of authentication scopes to use with this package.
func DefaultAuthScopes() []string {
	return []string{
		"https://www.googleapis.com/auth/cloud-platform",
	}
}
