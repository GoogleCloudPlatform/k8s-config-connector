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
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/gcp/profiler"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/k8s"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/logging"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/metrics"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/ready"

	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	flag "github.com/spf13/pflag"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/apis/meta/v1/unstructured"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/apimachinery/pkg/runtime/schema"
	_ "k8s.io/client-go/plugin/pkg/client/auth/gcp"
	"k8s.io/client-go/rest"
	"k8s.io/klog/v2"
	"sigs.k8s.io/controller-runtime/pkg/client"
	"sigs.k8s.io/controller-runtime/pkg/client/config"
	crlog "sigs.k8s.io/controller-runtime/pkg/log"
)

const (
	NumberOfWorkers    = 20
	MaximumListResults = 50
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

	// Get a client to talk to the APIServer
	kubeClient, err := client.New(restConfig, client.Options{})
	if err != nil {
		return fmt.Errorf("building kubernetes client: %w", err)
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

	for {
		time.Sleep(time.Duration(metricInterval) * time.Second)
		if !crdInfos.HasSyncedOnce() {
			logger.Info("CRDs have not yet synced, skipping metric collection")
			continue
		}
		var gvks []schema.GroupVersionKind
		for _, crdInfo := range crdInfos.Snapshot() {
			if !strings.HasSuffix(crdInfo.GVK.Group, ".cnrm.cloud.google.com") {
				continue
			}
			gvks = append(gvks, crdInfo.GVK)
		}

		if err := doRecord(ctx, kubeClient, gvks); err != nil {
			logger.Error(err, "error recording metrics")
		}
	}
}

func doRecord(ctx context.Context, c client.Client, gvks []schema.GroupVersionKind) error {
	logger := klog.FromContext(ctx)

	// reset all metrics in this vector before the new run of recording
	appliedResources.Reset()
	// worker pool with semaphore
	sem := make(chan struct{}, NumberOfWorkers)
	for _, gvk := range gvks {
		gvk := gvk
		sem <- struct{}{}
		go func() {
			defer func() { <-sem }()
			err := recordMetricsForGVK(ctx, c, gvk)
			if err != nil {
				logger.Error(err, "error recording metrics for CRD", "gvk", gvk.String())
			}
		}()
	}
	for i := 0; i < NumberOfWorkers; i++ {
		sem <- struct{}{}
	}
	logger.Info("finished one run of recording resource metrics.")
	return nil
}

func forEach(c client.Client, gvk schema.GroupVersionKind, listOptions *client.ListOptions, fn func(unstructured.Unstructured) error) error {
	for ok := true; ok; ok = listOptions.Continue != "" {
		list := unstructured.UnstructuredList{}
		list.SetGroupVersionKind(gvk)
		err := c.List(context.Background(), &list, listOptions)
		if err != nil {
			return fmt.Errorf("error listing objects: %w", err)
		}
		for _, item := range list.Items {
			if err := fn(item); err != nil {
				return err
			}
		}
		listOptions.Continue = list.GetContinue()
	}
	return nil
}

func recordMetricsForGVK(ctx context.Context, c client.Client, gvk schema.GroupVersionKind) error {
	logger := klog.FromContext(ctx)

	opts := &client.ListOptions{
		Limit: MaximumListResults,
		Raw:   &v1.ListOptions{},
	}
	statsNamespaceMap := make(map[string]*Stats)
	if err := forEach(c, gvk, opts, func(obj unstructured.Unstructured) error {
		namespace := obj.GetNamespace()
		s := statsNamespaceMap[namespace]
		if s == nil {
			s = &Stats{make(map[string]int64)}
			statsNamespaceMap[namespace] = s
		}
		lastCondition, err := getTheLastCondition(obj)
		if err != nil {
			logger.Error(err, "error getting the last condition for metrics", "gvk", gvk.String())
			return nil
		}
		s.countByStatus[lastCondition]++
		return nil
	}); err != nil {
		return fmt.Errorf("error listing objects for %v: %w", gvk.String(), err)
	}
	for ns, stats := range statsNamespaceMap {
		for status, count := range stats.countByStatus {
			logger.V(2).Info("posting metrics", "namespace", ns, "gvk", gvk.String(), "status", status, "count", count)
			appliedResources.WithLabelValues(ns, gvk.GroupKind().String(), status).Set(float64(count))
		}
	}
	return nil
}

type Stats struct {
	countByStatus map[string]int64
}

// TODO: consolidate the logic with krmtotf.GetReadyCondition
func getTheLastCondition(obj unstructured.Unstructured) (string, error) {
	currConditionsRaw, found, err := unstructured.NestedSlice(obj.Object, "status", "conditions")
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
