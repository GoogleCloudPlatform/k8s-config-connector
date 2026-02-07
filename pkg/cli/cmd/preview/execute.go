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

package preview

import (
	"context"
	"flag"
	"fmt"
	"time"

	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/cli/cmd/preview/parameters"
	corepreview "github.com/GoogleCloudPlatform/k8s-config-connector/pkg/cli/preview"
	"golang.org/x/oauth2"
	"golang.org/x/oauth2/google"
	"k8s.io/client-go/rest"
	"k8s.io/client-go/tools/clientcmd"
	"k8s.io/klog/v2"
	"k8s.io/klog/v2/klogr"
	"sigs.k8s.io/controller-runtime/pkg/log"
)

const (
	defaultScope   = "https://www.googleapis.com/auth/cloud-platform"
	defaultTimeout = 15
)

func Execute(ctx context.Context, opts *parameters.Parameters) error {
	// Use a custom FlagSet for klog to avoid conflicts with global flag.CommandLine
	klogFlags := flag.NewFlagSet("klog", flag.ExitOnError)
	klog.InitFlags(klogFlags)
	if err := klogFlags.Set("v", fmt.Sprintf("%d", opts.Verbose)); err != nil {
		fmt.Printf("Failed to set -v flag: %v\n", err)
	}
	if err := klogFlags.Set("alsologtostderr", "true"); err != nil {
		fmt.Printf("Failed to set -alsologtostderr flag: %v\n", err)
	}
	log.SetLogger(klogr.New())
	defer klog.Flush()

	klog.V(0).Info("Starting preview tool.")
	upstreamRESTConfig, err := getRESTConfig(ctx, opts)
	if err != nil {
		return fmt.Errorf("error building kubeconfig: %w", err)
	}
	recorder := corepreview.NewRecorder()
	if err := recorder.PreloadGKNN(ctx, upstreamRESTConfig, opts.Namespace); err != nil {
		return fmt.Errorf("error preloading the list of resources to reconcile: %w", err)
	}

	authorization, err := getGCPAuthorization(ctx, opts)
	if err != nil {
		return fmt.Errorf("error building GCP authorization: %w", err)
	}
	previewInstance, err := corepreview.NewPreviewInstance(recorder, corepreview.PreviewInstanceOptions{
		UpstreamRESTConfig:       upstreamRESTConfig,
		UpstreamGCPAuthorization: authorization,
		UpstreamGCPHTTPClient:    nil,
		UpstreamGCPQPS:           opts.GCPQPS,
		UpstreamGCPBurst:         opts.GCPBurst,
		Namespace:                opts.Namespace,
	})
	if err != nil {
		return fmt.Errorf("building preview instance: %w", err)
	}
	// TODO: Consider if 0 means no timeout.
	if opts.Timeout == 0 {
		opts.Timeout = defaultTimeout
	}
	ctx, cancel := context.WithTimeout(ctx, time.Duration(opts.Timeout)*time.Minute)
	defer cancel()
	defer func() {
		if err := printCapturedObjects(recorder, opts); err != nil {
			klog.Error(err, "error printing captured objects")
		}
	}()

	errChan := make(chan error, 1)
	go func() {
		errChan <- previewInstance.Start(ctx)
	}()

	select {
	case err := <-errChan:
		if err != nil {
			return fmt.Errorf("error starting preview: %w", err)
		}
	case <-ctx.Done():
		return fmt.Errorf("timeout reached: %w", ctx.Err())
	}
	klog.V(0).Info("Preview finished successfully")
	return nil
}

func getRESTConfig(ctx context.Context, opts *parameters.Parameters) (*rest.Config, error) {
	// TODO: Add rate limiting to Kube client.
	if opts.InCluster {
		return rest.InClusterConfig()
	}
	var loadingRules clientcmd.ClientConfigLoader
	if opts.Kubeconfig != "" {
		loadingRules = &clientcmd.ClientConfigLoadingRules{ExplicitPath: opts.Kubeconfig}
	} else {
		loadingRules = clientcmd.NewDefaultClientConfigLoadingRules()
	}
	return clientcmd.NewNonInteractiveDeferredLoadingClientConfig(
		loadingRules,
		&clientcmd.ConfigOverrides{}).ClientConfig()
}

func getGCPAuthorization(ctx context.Context, opts *parameters.Parameters) (oauth2.TokenSource, error) {
	// TODO: Add scope
	scopes := []string{defaultScope}
	ts, err := google.DefaultTokenSource(ctx, scopes...)
	if err != nil {
		return nil, err
	}
	return ts, nil
}

func printCapturedObjects(recorder *corepreview.Recorder, opts *parameters.Parameters) error {
	now := time.Now()
	timestamp := now.Format("20060102-150405.000")
	summaryFile := fmt.Sprintf("%s-%s", opts.ReportNamePrefix, timestamp)
	if err := recorder.SummaryReport(summaryFile); err != nil {
		return fmt.Errorf("error writing summary: %w", err)
	}

	if opts.FullReport {
		outputFile := fmt.Sprintf("%s-%s-full", opts.ReportNamePrefix, timestamp)
		if err := recorder.ExportDetailObjectsEvent(outputFile); err != nil {
			return fmt.Errorf("error writing events: %w", err)
		}

	}
	return nil
}
