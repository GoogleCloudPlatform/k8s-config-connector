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
	"time"

	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/cli/preview"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/resourceconfig"
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
	kubeconfigFlag = "kubeconfig"
	timeoutFlag    = "timeout"

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
	gcpQPS           float64
	gcpBurst         int
	namespace        string
	inCluster        bool
	reconcilerTypeOverride []string
}

type TimeoutError struct {
	Err error
}

func (e *TimeoutError) Error() string {
	return fmt.Sprintf("timeout reached: %v", e.Err)
}

func (e *TimeoutError) Unwrap() error {
	return e.Err
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
	cmd.Flags().Float64VarP(&opts.gcpQPS, "gcpQPS", "q", 5.0, "Maximum qps for GCP API requests, per service. Default to 5.0. Set gcpQPS to 0 to disable rate limiting.")
	cmd.Flags().IntVarP(&opts.gcpBurst, "gcpBurst", "b", 5, "Maximum burst for GCP API requests, per service. Default to 5. Set gcpQPS to 0 to disable rate limiting.")
	cmd.Flags().StringVarP(&opts.namespace, "namespace", "n", "", "Namespace to preview. If not specified, all namespaces will be previewed.")
	cmd.Flags().BoolVarP(&opts.inCluster, "in-cluster", "", false, "Run in GKE cluster.")
	cmd.Flags().StringArrayVarP(&opts.reconcilerTypeOverride, "reconciler-type-override", "r", []string{}, "Reconciler type to override in form Kind.Group:ReconcilerType.")
	return cmd
}

func getRESTConfig(ctx context.Context, opts *PreviewOptions) (*rest.Config, error) {
	// TODO: Add rate limiting.
	if opts.inCluster {
		return rest.InClusterConfig()
	}
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
	defer klog.Flush()
	defaultRecorder, err := runKCCManagerPreview(ctx, map[string]string{}, opts)
	if err != nil {
		if _, ok := err.(*TimeoutError); ok {
			klog.Info("Timeout reached when running preview with default reconcilers, exporting partial report")
			return nil
		}
		return fmt.Errorf("error running preview with default reconcilers: %w", err)
	}

	// Load the static config to generation overrides.
	config := resourceconfig.LoadConfig()
	override, err := generateAlternativeReconcilerOverride(config)
	if err != nil {
		return fmt.Errorf("failed to generate reconciler config: %w", err)
	}
	alternativeRecorder, err := runKCCManagerPreview(ctx, override, opts)
	if err != nil {
		if _, ok := err.(*TimeoutError); ok {
			klog.Info("Timeout reached when running preview with alternative reconcilers, exporting partial report")
			return nil
		}
		return fmt.Errorf("error running preview with alternative reconcilers: %w", err)
	}

	report := preview.NewPreviewReport(defaultRecorder, alternativeRecorder)
	// If running as GKE job, just print the failure details.
	if opts.inCluster {
		for _, failure := range report.GetFailures(){
			klog.InfoS("PreviewResult", "ns", failure.GKNN.Namespace, "name", failure.GKNN.Name, "group", failure.GKNN.Group, "kind", failure.GKNN.Kind, "controller_type", failure.ControllerType, "reconcile_status", failure.ReconcileStatus.String(), "fields", preview.FormatFieldIDs(failure.Diffs))
		}
		return nil
	}
	if opts.fullReport {
		defaultRecorder.ExportDetailObjectsEvent(generateFilename(opts.reportNamePrefix) + "default-full-events")
		alternativeRecorder.ExportDetailObjectsEvent(generateFilename(opts.reportNamePrefix) + "alternative-full-events")
	}
	report.ExportSummary(generateFilename(opts.reportNamePrefix) + "-summary")
	report.ExportFailedResults(generateFilename(opts.reportNamePrefix) + "-failure-details")
	return nil
}

func runKCCManagerPreview(ctx context.Context, reconcilerOverride map[string]string, opts *PreviewOptions) (*preview.Recorder, error) {
	upstreamRESTConfig, err := getRESTConfig(ctx, opts)
	if err != nil {
		return nil, fmt.Errorf("error building kubeconfig: %w", err)
	}
	authorization, err := getGCPAuthorization(ctx, opts)
	if err != nil {
		return nil, fmt.Errorf("error building GCP authorization: %w", err)
	}
	recorder := preview.NewRecorder()
	klog.Info("Preloading the list of resources to reconcile")
	if err := recorder.PreloadGKNN(ctx, upstreamRESTConfig, opts.namespace); err != nil {
		return nil, fmt.Errorf("error preload the list of resources to reconcile: %w", err)
	}
	klog.Info("Successfully preload the list of resources to reconcile.")

	preview, err := preview.NewPreviewInstance(recorder, preview.PreviewInstanceOptions{
		UpstreamRESTConfig:       upstreamRESTConfig,
		UpstreamGCPAuthorization: authorization,
		UpstreamGCPHTTPClient:    nil,
		UpstreamGCPQPS:           opts.gcpQPS,
		UpstreamGCPBurst:         opts.gcpBurst,
		Namespace:                opts.namespace,
		ReconcilerOverride:       reconcilerOverride,
	})
	if err != nil {
		return nil, fmt.Errorf("building preview instance: %v", err)
	}
	ctx, cancel := context.WithTimeout(ctx, time.Duration(opts.timeout)*time.Minute)
	defer cancel()

	errChan := make(chan error, 1)
	go func() {
		klog.Info("Starting preview")
		errChan <- preview.Start(ctx)
	}()

	select {
	case err := <-errChan:
		if err != nil {
			return nil, fmt.Errorf("error starting preview: %w", err)
		}
	case <-ctx.Done():
		return recorder, &TimeoutError{Err: ctx.Err()}
	}
	klog.Info("Preview finished successfully")
	return recorder, nil
}

func generateFilename(prefix string) string {
	now := time.Now()
	timestamp := now.Format("20060102-150405.000")
	return fmt.Sprintf("%s-%s", prefix, timestamp)
}

// generateAlternativeReconcilerOverride generates a map of resource to alternative reconciler type.
func generateAlternativeReconcilerOverride(config resourceconfig.ResourcesControllerMap) (map[string]string, error) {
	overrideMap := make(map[string]string)
	for gk, controllerConfig := range config {
		for _, controller := range controllerConfig.SupportedControllers {
			// Assuming there are maxium 2 supported controller types.
			// If there is alternative controller type, we will use it.
			if controller != controllerConfig.DefaultController {
				overrideMap[gk.String()] = string(controller)
				break
			}
		}
	}
	return overrideMap, nil
}