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

package preview

import (
	"context"
	"fmt"
	"os/signal"
	"syscall"
	"time"

	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/cli/preview"
	"github.com/spf13/cobra"
	"golang.org/x/oauth2"
	"golang.org/x/oauth2/google"
	"k8s.io/client-go/rest"
	"k8s.io/client-go/tools/clientcmd"
	"k8s.io/klog/v2"
	"k8s.io/klog/v2/klogr"
	"sigs.k8s.io/controller-runtime/pkg/log"
)

const (
	kubeconfigFlag     = "kubeconfig"
	timeoutFlag        = "timeout"

	reportNamePrefixFlag = "report-prefix"
	defaultScope         = "https://www.googleapis.com/auth/cloud-platform"
)

const (
	examples = `
	# preview Config Connector resources
	kompanion preview
	`
)

type PreviewOptions struct {
	kubeconfig       string
	timeout          int
	reportNamePrefix string
	fullReport       bool
	qps              float64
	burst            int
}

func BuildPreviewCmd() *cobra.Command {
	var opts PreviewOptions
	cmd := &cobra.Command{
		Use:     "preview",
		Short:   "preview Config Connector resources",
		Example: examples,
		RunE: func(cmd *cobra.Command, args []string) error {
			return RunPreview(cmd.Context(), &opts)
		},
		Args: cobra.ExactArgs(0),
	}
	// TODO: Add scope
	cmd.Flags().StringVarP(&opts.kubeconfig, kubeconfigFlag, "", opts.kubeconfig, "path to the kubeconfig file.")
	cmd.Flags().IntVarP(&opts.timeout, timeoutFlag, "", 15, "timeout in minutes. Default to 15 minutes.")
	cmd.Flags().StringVarP(&opts.reportNamePrefix, reportNamePrefixFlag, "", "preview-report", "Prefix for the report name. The tool appends a timestamp to this in the format \"YYYYMMDD-HHMMSS.milliseconds\".")
	cmd.Flags().BoolVarP(&opts.fullReport, "full-report", "f", false, "Enable verbose logging.")
	cmd.Flags().Float64Var(&opts.qps, "qps", 5.0, "QPS to use while talking with the GCP servers per API. Default to 5.0.")
	cmd.Flags().IntVar(&opts.burst, "burst", 5, "Burst to use while talking with the GCP servers per API. Default to 5.")

	return cmd
}

func getRESTConfig(ctx context.Context, opts *PreviewOptions) (*rest.Config, error) {
	// TODO: Add rate limiting.
	var loadingRules clientcmd.ClientConfigLoader
	if opts.kubeconfig != "" {
		loadingRules = &clientcmd.ClientConfigLoadingRules{ExplicitPath: opts.kubeconfig}
	} else {
		loadingRules = clientcmd.NewDefaultClientConfigLoadingRules()
	}
	return clientcmd.NewNonInteractiveDeferredLoadingClientConfig(
		loadingRules,
		&clientcmd.ConfigOverrides{}).ClientConfig()
}

func getGCPAuthorization(ctx context.Context, opts *PreviewOptions) (oauth2.TokenSource, error) {
	// TODO: Add scope
	scopes := []string{defaultScope}
	ts, err := google.DefaultTokenSource(ctx, scopes...)
	if err != nil {
		return nil, err
	}
	return ts, nil
}

func RunPreview(ctx context.Context, opts *PreviewOptions) error {
	log.SetLogger(klogr.New())
	klog.Info("Starting preview tool.")
	upstreamRESTConfig, err := getRESTConfig(ctx, opts)
	if err != nil {
		return fmt.Errorf("error building kubeconfig: %w", err)
	}
	recorder := preview.NewRecorder()
	if err := recorder.PreloadGKNN(ctx, upstreamRESTConfig); err != nil {
		return fmt.Errorf("error preload the list of resources to reconcile: %w", err)
	}
	klog.Info("Successfully preload the list of resources to reconcile.")
	authorization, err := getGCPAuthorization(ctx, opts)
	if err != nil {
		return fmt.Errorf("error building GCP authorization: %w", err)
	}
	preview, err := preview.NewPreviewInstance(recorder, preview.PreviewInstanceOptions{
		UpstreamRESTConfig:       upstreamRESTConfig,
		UpstreamGCPAuthorization: authorization,
		UpstreamGCPHTTPClient:    nil,
		QPS:                      opts.qps,
		Burst:                    opts.burst,
	})
	if err != nil {
		return fmt.Errorf("building preview instance: %v", err)
	}
	defer func() {
		if err := printCapturedObjects(recorder, opts.reportNamePrefix, opts.fullReport); err != nil {
			klog.Errorf("error printing captured objects: %v", err)
		}
	}()
	signalCtx, stop := signal.NotifyContext(ctx, syscall.SIGINT, syscall.SIGTERM)
	defer stop()

	runCtx, cancel := context.WithTimeout(signalCtx, time.Duration(opts.timeout)*time.Minute)
	defer cancel()

	if err := preview.Start(runCtx); err != nil {
		return fmt.Errorf("error running preview: %w", err)
	}
	klog.Info("Preview finished successfully")
	return nil
}

func printCapturedObjects(recorder *preview.Recorder, prefix string, full bool) error {
	now := time.Now()
	timestamp := now.Format("20060102-150405.000")
	summaryFile := fmt.Sprintf("%s-%s", prefix, timestamp)
	if err := recorder.SummaryReport(summaryFile); err != nil {
		return fmt.Errorf("error writing summary: %w", err)
	}

	if full {
		outputFile := fmt.Sprintf("%s-%s-full", prefix, timestamp)
		if err := recorder.ExportDetailObjectsEvent(outputFile); err != nil {
			return fmt.Errorf("error writing events: %w", err)
		}
	}
	return nil
}
