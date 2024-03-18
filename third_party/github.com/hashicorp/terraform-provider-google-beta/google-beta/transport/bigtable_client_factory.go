// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0
package transport

import (
	"context"
	"os"

	"cloud.google.com/go/bigtable"
	"golang.org/x/oauth2"
	"google.golang.org/api/option"
)

type BigtableClientFactory struct {
	gRPCLoggingOptions  []option.ClientOption
	UserAgent           string
	TokenSource         oauth2.TokenSource
	BillingProject      string
	UserProjectOverride bool
}

func (s BigtableClientFactory) NewInstanceAdminClient(project string) (*bigtable.InstanceAdminClient, error) {
	var opts []option.ClientOption
	if requestReason := os.Getenv("CLOUDSDK_CORE_REQUEST_REASON"); requestReason != "" {
		opts = append(opts, option.WithRequestReason(requestReason))
	}

	if s.UserProjectOverride && s.BillingProject != "" {
		opts = append(opts, option.WithQuotaProject(s.BillingProject))
	}

	opts = append(opts, option.WithTokenSource(s.TokenSource), option.WithUserAgent(s.UserAgent))
	opts = append(opts, s.gRPCLoggingOptions...)

	// //

	// opts = append(opts, option.WithHTTPClient())

	opts = append(opts, s.gRPCLoggingOptions...)
	if GRPCConnectionFactory != nil {
		conn := GRPCConnectionFactory(ctx)
		opts = append(opts, option.WithGRPCConn(conn))
	}
	return bigtable.NewInstanceAdminClient(ctx, project, opts...)
}

func (s BigtableClientFactory) NewAdminClient(project, instance string) (*bigtable.AdminClient, error) {
	ctx := context.TODO()

	var opts []option.ClientOption
	if requestReason := os.Getenv("CLOUDSDK_CORE_REQUEST_REASON"); requestReason != "" {
		opts = append(opts, option.WithRequestReason(requestReason))
	}

	if s.UserProjectOverride && s.BillingProject != "" {
		opts = append(opts, option.WithQuotaProject(s.BillingProject))
	}

	opts = append(opts, option.WithTokenSource(s.TokenSource), option.WithUserAgent(s.UserAgent))
	// if s.HTTPClient != nil {
	// 	// opts = append(opts, option.WithHTTPClient(s.HTTPClient))
	// } else {
	// 	opts = append(opts, s.gRPCLoggingOptions...)
	// }
	opts = append(opts, s.gRPCLoggingOptions...)
	if GRPCConnectionFactory != nil {
		conn := GRPCConnectionFactory(ctx)
		opts = append(opts, option.WithGRPCConn(conn))
	}
	return bigtable.NewAdminClient(ctx, project, instance, opts...)
}

func (s BigtableClientFactory) NewClient(project, instance string) (*bigtable.Client, error) {
	ctx := context.TODO()

	var opts []option.ClientOption
	if requestReason := os.Getenv("CLOUDSDK_CORE_REQUEST_REASON"); requestReason != "" {
		opts = append(opts, option.WithRequestReason(requestReason))
	}

	if s.UserProjectOverride && s.BillingProject != "" {
		opts = append(opts, option.WithQuotaProject(s.BillingProject))
	}

	opts = append(opts, option.WithTokenSource(s.TokenSource), option.WithUserAgent(s.UserAgent))
	opts = append(opts, s.gRPCLoggingOptions...)
	if GRPCConnectionFactory != nil {
		conn := GRPCConnectionFactory(ctx)
		opts = append(opts, option.WithGRPCConn(conn))
	}
	return bigtable.NewClient(ctx, project, instance, opts...)
}
