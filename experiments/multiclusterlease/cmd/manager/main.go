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

/*


Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package main

import (
	"context"
	"flag"
	"fmt"
	"os"

	"cloud.google.com/go/storage"
	"github.com/GoogleCloudPlatform/k8s-config-connector/experiments/multiclusterlease/controllers"
	"github.com/google/uuid"
	"go.uber.org/zap/zapcore"
	_ "k8s.io/client-go/plugin/pkg/client/auth/gcp"
	ctrl "sigs.k8s.io/controller-runtime"
	"sigs.k8s.io/controller-runtime/pkg/log/zap"
	// +kubebuilder:scaffold:imports
)

var (
	setupLog = ctrl.Log.WithName("setup")
)

func main() {
	var metricsAddr string
	var gcsBucketName string
	var clusterIdentity string
	var verbose bool

	flag.StringVar(&metricsAddr, "metrics-addr", ":8080", "The address the metric endpoint binds to.")
	flag.StringVar(&gcsBucketName, "gcs-bucket", "", "The GCS bucket to use for multi-cluster leader election.")
	flag.StringVar(&clusterIdentity, "cluster-identity", "", "The identity of this cluster (defaults to a generated UUID).")
	flag.BoolVar(&verbose, "verbose", false, "Enable verbose logging")
	flag.Parse()

	// Configure logging
	opts := []zap.Opts{zap.UseDevMode(true)}
	if verbose {
		opts = append(opts, zap.Level(zapcore.DebugLevel))
	}
	ctrl.SetLogger(zap.New(opts...))

	// Set up a unique identity for this cluster if not provided
	if clusterIdentity == "" {
		var err error
		clusterIdentity, err = os.Hostname()
		if err != nil {
			clusterIdentity = uuid.New().String()
		}
		setupLog.Info("Generated cluster identity", "identity", clusterIdentity)
	}

	// Validate required flags
	if gcsBucketName == "" {
		setupLog.Error(fmt.Errorf("gcs-bucket flag is required"), "missing required flag")
		os.Exit(1)
	}

	// Create manager
	setupLog.Info("Creating manager")
	mgr, err := ctrl.NewManager(ctrl.GetConfigOrDie(), ctrl.Options{
		Scheme:             controllers.BuildScheme(),
		MetricsBindAddress: metricsAddr,
		Port:               9443,
		LeaderElection:     false,
	})
	if err != nil {
		setupLog.Error(err, "unable to start manager")
		os.Exit(1)
	}

	// Create GCS client
	setupLog.Info("Creating GCS client", "bucket", gcsBucketName)
	ctx := context.Background()
	gcsClient, err := storage.NewClient(ctx)
	if err != nil {
		setupLog.Error(err, "unable to create GCS client")
		os.Exit(1)
	}
	defer gcsClient.Close()

	// Verify bucket exists
	setupLog.Info("Verifying GCS bucket exists")
	bucket := gcsClient.Bucket(gcsBucketName)
	_, err = bucket.Attrs(ctx)
	if err != nil {
		setupLog.Error(err, "unable to access GCS bucket", "bucket", gcsBucketName)
		os.Exit(1)
	}

	// Create and set up the MultiClusterLeaseReconciler
	setupLog.Info("Creating MultiClusterLeaseReconciler")
	reconciler := controllers.NewMultiClusterLeaseReconciler(
		mgr.GetClient(),
		ctrl.Log.WithName("controllers").WithName("MultiClusterLease"),
		gcsClient,
		gcsBucketName,
		clusterIdentity,
	)

	if err = reconciler.SetupWithManager(mgr); err != nil {
		setupLog.Error(err, "unable to create controller", "controller", "MultiClusterLease")
		os.Exit(1)
	}
	// +kubebuilder:scaffold:builder

	setupLog.Info("Starting manager", "clusterIdentity", clusterIdentity, "gcsBucket", gcsBucketName)
	if err := mgr.Start(ctrl.SetupSignalHandler()); err != nil {
		setupLog.Error(err, "problem running manager")
		os.Exit(1)
	}
}
