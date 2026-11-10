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

package main

import (
	"context"
	goflag "flag"
	"fmt"
	"os"

	gcs "cloud.google.com/go/storage"
	"github.com/gke-labs/multicluster-leader-election/controllers"
	"github.com/gke-labs/multicluster-leader-election/pkg/storage"
	flag "github.com/spf13/pflag"
	_ "k8s.io/client-go/plugin/pkg/client/auth/gcp"
	"k8s.io/klog/v2"
	ctrl "sigs.k8s.io/controller-runtime"
	metricsserver "sigs.k8s.io/controller-runtime/pkg/metrics/server"
	// +kubebuilder:scaffold:imports
)

var setupLog = ctrl.Log.WithName("setup")

func main() {
	if err := run(context.Background()); err != nil {
		fmt.Fprintf(os.Stderr, "%v\n", err)
		os.Exit(1)
	}
}

func run(ctx context.Context) error {
	var metricsAddr string
	var gcsBucketName string

	flag.StringVar(&metricsAddr, "metrics-addr", ":8080", "The address the metric endpoint binds to.")
	flag.StringVar(&gcsBucketName, "gcs-bucket", "", "The GCS bucket to use for multi-cluster leader election.")

	// Configure logging
	klogFlagSet := goflag.NewFlagSet("klog", goflag.ExitOnError)
	klog.InitFlags(klogFlagSet)
	// Support default klog verbosity `-v`
	flag.CommandLine.AddGoFlag(klogFlagSet.Lookup("v"))
	flag.CommandLine.AddGoFlagSet(goflag.CommandLine)
	flag.Parse()

	ctx = klog.NewContext(ctx, setupLog)
	ctrl.SetLogger(klog.NewKlogr())

	// Validate required flags
	if gcsBucketName == "" {
		err := fmt.Errorf("gcs-bucket flag is required")
		setupLog.Error(err, "missing required flag")
		return err
	}

	// Create manager
	setupLog.Info("Creating manager")
	mgr, err := ctrl.NewManager(ctrl.GetConfigOrDie(), ctrl.Options{
		Scheme: controllers.BuildScheme(),
		Metrics: metricsserver.Options{
			BindAddress: metricsAddr,
		},
		LeaderElection: false,
	})
	if err != nil {
		setupLog.Error(err, "unable to start manager")
		return err
	}

	// Create GCS client
	setupLog.Info("Creating GCS client", "bucket", gcsBucketName)
	gcsClient, err := gcs.NewClient(ctx)
	if err != nil {
		setupLog.Error(err, "unable to create GCS client")
		return err
	}
	defer gcsClient.Close()

	// Verify bucket exists
	setupLog.Info("Verifying GCS bucket exists")
	bucket := gcsClient.Bucket(gcsBucketName)
	_, err = bucket.Attrs(ctx)
	if err != nil {
		setupLog.Error(err, "unable to access GCS bucket", "bucket", gcsBucketName)
		return err
	}

	// Create storage backend
	setupLog.Info("Creating GCS storage backend", "bucket", gcsBucketName)
	gcsStorage := storage.NewGCSStorage(gcsClient, gcsBucketName)

	// Create and set up the MultiClusterLeaseReconciler
	setupLog.Info("Creating MultiClusterLeaseReconciler")
	reconciler := controllers.NewMultiClusterLeaseReconciler(
		mgr.GetClient(),
		ctrl.Log.WithName("controllers").WithName("MultiClusterLease"),
		gcsStorage,
	)

	if err = reconciler.SetupWithManager(mgr); err != nil {
		setupLog.Error(err, "unable to create controller", "controller", "MultiClusterLease")
		return err
	}
	// +kubebuilder:scaffold:builder

	setupLog.Info("Starting manager", "gcsBucket", gcsBucketName)
	if err := mgr.Start(ctrl.SetupSignalHandler()); err != nil {
		setupLog.Error(err, "problem running manager")
		return err
	}
	return nil
}
