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
	"time"

	dclmetadata "github.com/GoogleCloudPlatform/k8s-config-connector/pkg/dcl/metadata"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/gcp/profiler"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/gvks/supportedgvks"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/k8s"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/logging"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/metrics"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/ready"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/servicemapping/servicemappingloader"

	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	flag "github.com/spf13/pflag"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/apis/meta/v1/unstructured"
	"k8s.io/apimachinery/pkg/runtime/schema"
	_ "k8s.io/client-go/plugin/pkg/client/auth/gcp"
	"k8s.io/klog/v2/klogr"
	"sigs.k8s.io/controller-runtime/pkg/client"
	"sigs.k8s.io/controller-runtime/pkg/client/config"
	klog "sigs.k8s.io/controller-runtime/pkg/log"
)

const (
	NumberOfWorkers    = 20
	MaximumListResults = 50
)

var (
	logger           = klog.Log.WithName("setup")
	appliedResources = metrics.NewAppliedResourcesCollector()
)

func main() {

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
	flag.CommandLine.AddGoFlagSet(goflag.CommandLine)
	flag.Parse()
	kccVersion := os.Getenv("CONFIG_CONNECTOR_VERSION")

	klog.SetLogger(klogr.New())

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
	cfg, err := config.GetConfig()
	if err != nil {
		logging.Fatal(err, "error getting configuration from APIServer.")
	}

	// Get a client to talk to the APIServer
	c, err := client.New(cfg, client.Options{})
	if err != nil {
		logging.Fatal(err, "error getting client.")
	}

	smLoader, err := servicemappingloader.New()
	if err != nil {
		logging.Fatal(err, "error getting new service mapping loader")
	}

	supportedGVKs := supportedgvks.All(smLoader, dclmetadata.New())
	for {
		err := doRecord(c, supportedGVKs)
		if err != nil {
			logger.Error(err, "error recording metrics.")
		}
		time.Sleep(time.Duration(metricInterval) * time.Second)
	}
}

func doRecord(c client.Client, gvks []schema.GroupVersionKind) error {
	logger.Info("listing all CRDs managed by Config Connector.")

	// reset all metrics in this vector before the new run of recording
	appliedResources.Reset()
	// worker pool with semaphore
	sem := make(chan struct{}, NumberOfWorkers)
	for _, gvk := range gvks {
		gvk := gvk
		sem <- struct{}{}
		go func() {
			defer func() { <-sem }()
			err := recordMetricsForGVK(c, gvk)
			if err != nil {
				logger.Error(err, "error recording metrics for CRD %v: %v", gvk.String())
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

func recordMetricsForGVK(c client.Client, gvk schema.GroupVersionKind) error {
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
			logger.Error(err, "error getting the last condition for metrics for %v: %v", gvk.String())
			return nil
		}
		s.countByStatus[lastCondition]++
		return nil
	}); err != nil {
		return fmt.Errorf("error listing objects for %v: %w", gvk.String(), err)
	}
	for ns, stats := range statsNamespaceMap {
		for status, count := range stats.countByStatus {
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
