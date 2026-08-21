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
	"errors"
	"flag"
	"fmt"
	"time"

	corepreview "github.com/GoogleCloudPlatform/k8s-config-connector/pkg/cli/preview"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/resourceconfig"
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

func Execute(ctx context.Context, opts *Options) error {
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

	if recorder.RemainResourcesCount == 0 {
		klog.V(0).Info("No resources found to reconcile")
		return nil
	}

	authorization, err := getGCPAuthorization(ctx, opts)
	if err != nil {
		return fmt.Errorf("error building GCP authorization: %w", err)
	}
	instanceOptions := corepreview.PreviewInstanceOptions{
		UpstreamRESTConfig:       upstreamRESTConfig,
		UpstreamGCPAuthorization: authorization,
		UpstreamGCPHTTPClient:    nil,
		UpstreamGCPQPS:           opts.GCPQPS,
		UpstreamGCPBurst:         opts.GCPBurst,
		Namespace:                opts.Namespace,
	}
	if opts.Timeout == 0 {
		opts.Timeout = defaultTimeout
	}

	// First run: execute the KCC Manager configuring resources strictly under their default controller type
	// This generates baseline preview results reflecting standard reconciliation behavior.
	defaultRunRecorder, err := runKCCManagerPreview(ctx, recorder, instanceOptions, opts.Timeout)
	if err != nil {
		var timeoutErr *TimeoutError
		if errors.As(err, &timeoutErr) {
			// Log out the timeout error and the number of resources not fully reconciled, then continue the run.
			klog.Errorf("Timeout reached during default controller preview run. Number of resources not fully reconciled: %d. Error: %v", defaultRunRecorder.GetRemainResourcesCount(), err)
		} else {
			return fmt.Errorf("error running KCC manager preview with default controller type: %w", err)
		}
	}
	defaultRunResult := defaultRunRecorder.GetOrCreateReconciledResults()

	// Second run: inject an ObjectTransformer into the instance options that explicitly overrides the
	// target controller execution flag for the collected resources, then launch the manager again.
	// This captures and isolates the results from alternative controller reconciliation pathways.
	instanceOptions.ObjectTransformers = []corepreview.ObjectTransformer{corepreview.NewSetAlternativeOverrideTransformer(opts.Namespace)}
	alternativeRunRecorder, err := runKCCManagerPreview(ctx, recorder, instanceOptions, opts.Timeout)
	if err != nil {
		var timeoutErr *TimeoutError
		if errors.As(err, &timeoutErr) {
			// Log out the timeout error and the number of resources not fully reconciled, then continue the run.
			klog.Errorf("Timeout reached during alternative controller preview run. Number of resources not fully reconciled: %d. Error: %v", alternativeRunRecorder.GetRemainResourcesCount(), err)
		} else {
			return fmt.Errorf("error running KCC manager preview with alternative controller type: %w", err)
		}
	}
	alternativeRunResult := alternativeRunRecorder.GetOrCreateReconciledResults()

	return printCapturedObjects(defaultRunRecorder, alternativeRunRecorder, defaultRunResult, alternativeRunResult, opts)
}

func getRESTConfig(ctx context.Context, opts *Options) (*rest.Config, error) {
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

func getGCPAuthorization(ctx context.Context, opts *Options) (oauth2.TokenSource, error) {
	// TODO: Add scope
	scopes := []string{defaultScope}
	ts, err := google.DefaultTokenSource(ctx, scopes...)
	if err != nil {
		return nil, err
	}
	return ts, nil
}

func printCapturedObjects(defaultRunRecorder, alternativeRunRecorder *corepreview.Recorder, defaultRunResult, alternativeRunResult *corepreview.RecorderReconciledResults, opts *Options) error {
	now := time.Now()
	timestamp := now.Format("20060102-150405.000")
	summaryFile := fmt.Sprintf("%s-%s", opts.ReportNamePrefix, timestamp)
	altExpectedMap := corepreview.GetAlternativeControllerExpectedMap(resourceconfig.ControllerConfigStatic)
	if err := defaultRunResult.CombinedSummaryReport(summaryFile, alternativeRunResult, altExpectedMap); err != nil {
		return fmt.Errorf("error writing summary: %w", err)
	}

	if opts.FullReport {
		outputFile := fmt.Sprintf("%s-%s-full", opts.ReportNamePrefix, timestamp)
		if err := defaultRunRecorder.ExportDetailObjectsEvent(outputFile + "-default"); err != nil {
			return fmt.Errorf("error writing events: %w", err)
		}
		if err := alternativeRunRecorder.ExportDetailObjectsEvent(outputFile + "-alternative"); err != nil {
			return fmt.Errorf("error writing events: %w", err)
		}
	}
	return nil
}

func runKCCManagerPreview(ctx context.Context, recorder *corepreview.Recorder, instanceOptions corepreview.PreviewInstanceOptions, timeout int) (*corepreview.Recorder, error) {
	// Make a copy of the recorder to avoid modifying the original recorder.
	recorder = recorder.DeepCopy()

	instance, err := corepreview.NewPreviewInstance(recorder, instanceOptions)
	if err != nil {
		return nil, fmt.Errorf("building preview instance: %w", err)
	}
	ctx, cancel := context.WithTimeout(ctx, time.Duration(timeout)*time.Minute)
	defer cancel()

	errChan := make(chan error, 1)
	go func() {
		// The manager runs in a background goroutine and continuously mutates the recorder.
		errChan <- instance.Start(ctx)
	}()

	select {
	case err := <-errChan:
		if err != nil {
			return nil, fmt.Errorf("error starting preview: %w", err)
		}
	case <-ctx.Done():
		// When a timeout occurs, the context is cancelled, signaling the manager to stop.
		// However, we must explicitly wait for the manager's goroutine to finish its graceful
		// shutdown before returning. This prevents a data race where the main thread reads
		// from the recorder (e.g., to generate reports) while the background manager is
		// still performing its final mutations.
		<-errChan
		return recorder, &TimeoutError{Err: ctx.Err()}
	}
	return recorder, nil
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
