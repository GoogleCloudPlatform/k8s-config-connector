// Copyright 2022 Google LLC
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

package main

import (
	"context"
	goflag "flag"
	"fmt"
	"net/http"
	_ "net/http/pprof" // Needed to allow pprof server to accept requests
	"os"
	"strings"
	"time"

	"github.com/GoogleCloudPlatform/k8s-config-connector/cmd/recorder/kube"
	opk8s "github.com/GoogleCloudPlatform/k8s-config-connector/operator/pkg/k8s"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/gcp/profiler"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/k8s"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/logging"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/metrics"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/ready"

	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	flag "github.com/spf13/pflag"
	"k8s.io/apimachinery/pkg/apis/meta/v1/unstructured"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/apimachinery/pkg/runtime/schema"
	_ "k8s.io/client-go/plugin/pkg/client/auth/gcp"
	"k8s.io/client-go/rest"
	"k8s.io/klog/v2"
	"sigs.k8s.io/controller-runtime/pkg/client/config"
	crlog "sigs.k8s.io/controller-runtime/pkg/log"
)

var (
	logger           = crlog.Log.WithName("setup")
	appliedResources = metrics.NewAppliedResourcesCollector()
)

func main() {
	if err := run(context.Background()); err != nil {
		fmt.Fprintf(os.Stderr, "%v\n", err)
		os.Exit(1)
	}
}

func run(ctx context.Context) error {
	var (
		prometheusScrapeEndpoint string
		metricInterval           int
		enablePprof              bool
		pprofPort                int
	)
	flag.StringVar(&prometheusScrapeEndpoint, "prometheus-scrape-endpoint", ":8888", "configure the Prometheus scrape endpoint; :8888 as default")
	flag.IntVar(&metricInterval, "metric-interval", 60, "the time interval of each recording in seconds")
	flag.BoolVar(&enablePprof, "enable-pprof", false, "Enable the pprof server.")
	flag.IntVar(&pprofPort, "pprof-port", 6060, "The port that the pprof server binds to if enabled.")
	profiler.AddFlag(flag.CommandLine)

	// Support default klog verbosity (so that we can see client-go traffic)
	klogFlagSet := goflag.NewFlagSet("klog", goflag.ExitOnError)
	klog.InitFlags(klogFlagSet)
	flag.CommandLine.AddGoFlag(klogFlagSet.Lookup("v"))

	flag.CommandLine.AddGoFlagSet(goflag.CommandLine)
	flag.Parse()

	kccVersion := os.Getenv("CONFIG_CONNECTOR_VERSION")

	logger := klog.NewKlogr()
	ctx = klog.NewContext(ctx, logger)
	crlog.SetLogger(logger)

	logger.Info("Recording the stats of Config Connector resources")

	// Start pprof server if enabled
	if enablePprof {
		go func() {
			if err := http.ListenAndServe(fmt.Sprintf(":%d", pprofPort), nil); err != nil {
				logger.Error(err, "error while running pprof server")
			}
		}()
	}

	// Start Cloud Profiler agent if enabled
	if err := profiler.StartIfEnabled(); err != nil {
		logging.Fatal(err, "error starting Cloud Profiler agent")
	}

	// Register the Prometheus metrics
	prometheus.MustRegister(appliedResources)
	prometheus.MustRegister(metrics.NewBuildInfoCollector(kccVersion))

	// Expose the registered metrics via HTTP.
	go func() {
		http.Handle("/metrics", promhttp.Handler())
		logging.Fatal(http.ListenAndServe(prometheusScrapeEndpoint, nil), "error registering the Prometheus HTTP handler")
	}()

	// Set up the HTTP server for the readiness probe
	logger.Info("Setting container as ready...")
	ready.SetContainerAsReady()
	logger.Info("Container is ready.")

	// Get a config to talk to the apiserver
	restConfig, err := config.GetConfig()
	if err != nil {
		return fmt.Errorf("error getting kubernetes configuration: %w", err)
	}

	restHTTPClient, err := rest.HTTPClientFor(restConfig)
	if err != nil {
		return fmt.Errorf("building kubernetes http client: %w", err)
	}

	kubeTarget, err := kube.NewTarget(restConfig, restHTTPClient)
	if err != nil {
		return fmt.Errorf("building kubernetes target: %w", err)
	}

	crdGVR := schema.GroupVersionResource{Group: "apiextensions.k8s.io", Version: "v1", Resource: "customresourcedefinitions"}
	crdInfos := kube.WatchKube(ctx, kubeTarget, crdGVR, buildCRDInfo)

	statViews := make(map[CRDInfo]*kube.KubeView[ResourceStats])

	for {
		time.Sleep(time.Duration(metricInterval) * time.Second)

		// Reset all metrics before updating.
		appliedResources.Reset()

		// Skip reporting if CRDs aren't synced up.
		if !crdInfos.HasSyncedOnce() {
			logger.Info("CRDs have not yet synced, skipping metric reporting")
			continue
		}

		seenCRDs := make(map[CRDInfo]bool)
		for _, crdInfo := range crdInfos.Snapshot() {
			// Skip non-KCC resources.
			if !strings.HasSuffix(crdInfo.GVK.Group, ".cnrm.cloud.google.com") {
				continue
			}

			// Skip ignored CRDs.
			crdName := gvkToCRDName(crdInfo.GVK)
			if _, ok := opk8s.IgnoredCRDList[crdName]; ok {
				logger.Error(fmt.Errorf("unexpected CRD %s", crdName),
					fmt.Sprintf("please run `kubectl delete crd %s` to "+
						"delete the orphaned CRD", crdName),
					"crd", crdName)
				continue
			}

			// Record all KCC CRDs we see, so we can clean up unused watches.
			seenCRDs[crdInfo] = true

			// Register watch for this resource if we haven't already.
			if _, ok := statViews[crdInfo]; !ok {
				statView := kube.WatchKube(ctx, kubeTarget, crdInfo.GVR, gatherResourceStats)
				statViews[crdInfo] = statView
			}

			// Skip reporting for this resource if we aren't synced up.
			if !statViews[crdInfo].HasSyncedOnce() {
				logger.Info("CRs have not yet synced, skipping metric reporting", "gvk", crdInfo.GVK)
				continue
			}

			// Aggregate stats for each namespace.
			nsAggStats := make(map[string]*AggregatedResourceStats)
			for i, s := range statViews[crdInfo].Snapshot() {
				ns := i.Namespace
				nsStats, ok := nsAggStats[ns]
				if !ok {
					nsStats = NewAggregatedResourceStats()
					nsAggStats[ns] = nsStats
				}
				nsStats.lastConditionCounts[s.lastCondition]++
			}

			// Record stats.
			for ns, stats := range nsAggStats {
				for condition, count := range stats.lastConditionCounts {
					logger.V(2).Info("posting metrics", "namespace", ns, "gvk", crdInfo.GVK.String(), "status", condition, "count", count)
					appliedResources.WithLabelValues(ns, crdInfo.GVK.GroupKind().String(), condition).Set(float64(count))
				}
			}
		}

		// Cleanup stale watches.
		for crdInfo, view := range statViews {
			if _, ok := seenCRDs[crdInfo]; !ok {
				logger.Info("removing stale watch for resource", "gvk", crdInfo.GVK.String())
				view.Close()
				delete(statViews, crdInfo)
			}
		}
	}
}

func gvkToCRDName(gvk schema.GroupVersionKind) string {
	pluralLowercaseKind := strings.ToLower(gvk.Kind) + "s"
	return pluralLowercaseKind + "." + gvk.Group
}

type ResourceStats struct {
	lastCondition string
}

type AggregatedResourceStats struct {
	lastConditionCounts map[string]int64
}

func NewAggregatedResourceStats() *AggregatedResourceStats {
	return &AggregatedResourceStats{
		lastConditionCounts: make(map[string]int64),
	}
}

func gatherResourceStats(u *unstructured.Unstructured) ResourceStats {
	lastCondition, err := getLastCondition(u)
	if err != nil {
		logger.Error(err, "error getting last condition for CR", "gvk", u.GroupVersionKind(), "name", u.GetName(), "namespace", u.GetNamespace())
	}
	return ResourceStats{
		lastCondition: lastCondition,
	}
}

// TODO: consolidate the logic with krmtotf.GetReadyCondition
func getLastCondition(u *unstructured.Unstructured) (string, error) {
	currConditionsRaw, found, err := unstructured.NestedSlice(u.Object, "status", "conditions")
	if err != nil {
		return "", err
	}
	if !found || len(currConditionsRaw) == 0 {
		return "NoCondition", nil
	}

	currConditions, err := k8s.MarshalAsConditionsSlice(currConditionsRaw)
	if err != nil {
		return "", err
	}
	if currConditions[0].Reason == "" {
		return k8s.NoCondition, nil
	}
	return currConditions[0].Reason, nil
}

type CRDInfo struct {
	GVR schema.GroupVersionResource
	GVK schema.GroupVersionKind
}

// customResourceDefinition is a cut-down version of the CRD resource, so we can easily extract the GVK/GVR
type customResourceDefinition struct {
	Spec customResourceDefinitionSpec `json:"spec"`
}

type customResourceDefinitionSpec struct {
	Names    customResourceDefinitionNames     `json:"names"`
	Versions []customResourceDefinitionVersion `json:"versions"`
}

type customResourceDefinitionNames struct {
	Kind string `json:"kind"`
}

type customResourceDefinitionVersion struct {
	Name string `json:"name"`
}

// buildCRDInfo extracts the GVK/GVR from a CRD.
func buildCRDInfo(u *unstructured.Unstructured) CRDInfo {
	if _, found := u.GetLabels()["cnrm.cloud.google.com/managed-by-kcc"]; !found {
		return CRDInfo{}
	}

	tokens := strings.SplitN(u.GetName(), ".", 2)
	if len(tokens) != 2 {
		logger.Info("cannot parse crd name", "name", u.GetName())
		return CRDInfo{}
	}

	crd := &customResourceDefinition{}
	if err := runtime.DefaultUnstructuredConverter.FromUnstructured(u.Object, crd); err != nil {
		logger.Error(err, "parsing CRD")
		return CRDInfo{}
	}

	kind := crd.Spec.Names.Kind
	if kind == "" {
		logger.Info("cannot extract crd kind", "name", u.GetName())
		return CRDInfo{}
	}

	version := ""
	for _, versionObj := range crd.Spec.Versions {
		if versionObj.Name != "" {
			version = versionObj.Name
			break
		}
	}

	if version == "" {
		logger.Info("cannot extract crd version", "name", u.GetName())
		return CRDInfo{}
	}

	var info CRDInfo
	info.GVR.Resource = tokens[0]
	info.GVR.Version = version
	info.GVR.Group = tokens[1]

	info.GVK.Group = tokens[1]
	info.GVK.Version = version
	info.GVK.Kind = kind

	return info
}
