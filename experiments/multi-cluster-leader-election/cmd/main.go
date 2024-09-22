/*
Copyright 2024.

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
	"flag"
	"os"
	"time"

	// Import all Kubernetes client auth plugins (e.g. Azure, GCP, OIDC, etc.)
	// to ensure that exec-entrypoint and run can make use of them.
	_ "k8s.io/client-go/plugin/pkg/client/auth"

	"cloud.google.com/go/storage"
	"k8s.io/apimachinery/pkg/runtime"
	utilruntime "k8s.io/apimachinery/pkg/util/runtime"
	clientgoscheme "k8s.io/client-go/kubernetes/scheme"
	"k8s.io/klog/v2"
	ctrl "sigs.k8s.io/controller-runtime"
	"sigs.k8s.io/controller-runtime/pkg/healthz"
	"sigs.k8s.io/controller-runtime/pkg/log/zap"
	metricsserver "sigs.k8s.io/controller-runtime/pkg/metrics/server"

	leaderelectionv1 "github.com/600lyy/multi-cluster-leader-election/api/v1"
	"github.com/600lyy/multi-cluster-leader-election/internal/controller"
	"github.com/600lyy/multi-cluster-leader-election/pkg/leaderelection"
	"github.com/600lyy/multi-cluster-leader-election/pkg/resourcelock"
	//+kubebuilder:scaffold:imports
)

var (
	scheme   = runtime.NewScheme()
	setupLog = ctrl.Log.WithName("setup")
)

func init() {
	utilruntime.Must(clientgoscheme.AddToScheme(scheme))

	utilruntime.Must(leaderelectionv1.AddToScheme(scheme))
	//+kubebuilder:scaffold:scheme
}

func main() {
	var metricsAddr string
	var enableLeaderElection bool
	var probeAddr string
	var leaseDuration int64
	var renewDeadline int64
	var retryPeriod int64
	var projectId string
	var leaseBucket string
	var leaseFile string
	flag.StringVar(&metricsAddr, "metrics-bind-address", ":8080", "The address the metric endpoint binds to.")
	flag.StringVar(&probeAddr, "health-probe-bind-address", ":8081", "The address the probe endpoint binds to.")
	flag.BoolVar(&enableLeaderElection, "leader-elect", false,
		"Enable leader election for controller manager. "+
			"Enabling this will ensure there is only one active controller manager.")
	flag.Int64Var(&leaseDuration, "lease-duration", 15,
		"LeaseDuration is the duration that non-leader candidates will wait to force acquire leadership. ")
	flag.Int64Var(&renewDeadline, "renew-deadline", 10,
		"RenewDeadline is the duration that the acting controlplane will retry refreshing leadership before giving up. ")
	flag.Int64Var(&retryPeriod, "retry-period", 4,
		"RetryPeriod is the duration the LeaderElector clients should wait between tries of actions. ")
	flag.StringVar(&projectId, "project-id", "", "Google project ID to which the bucke belong")
	flag.StringVar(&leaseBucket, "lease-storage-bucket", "", "The Google Clous Storage bucket that holds the lease.")
	flag.StringVar(&leaseFile, "lease-file", "", "The file name of the lease in the storage bucket")
	opts := zap.Options{
		Development: true,
	}
	klog.InitFlags(nil)
	opts.BindFlags(flag.CommandLine)
	flag.Parse()

	ctrl.SetLogger(zap.New(zap.UseFlagOptions(&opts)))

	duration := time.Duration(leaseDuration) * time.Second
	retry := time.Duration(retryPeriod) * time.Second
	renew := time.Duration(renewDeadline) * time.Second

	mgr, err := ctrl.NewManager(ctrl.GetConfigOrDie(), ctrl.Options{
		Scheme:                 scheme,
		Metrics:                metricsserver.Options{BindAddress: metricsAddr},
		HealthProbeBindAddress: probeAddr,
		LeaderElection:         enableLeaderElection,
		LeaderElectionID:       "b4973361.600lyy.io",
		// LeaderElectionReleaseOnCancel defines if the leader should step down voluntarily
		// when the Manager ends. This requires the binary to immediately end when the
		// Manager is stopped, otherwise, this setting is unsafe. Setting this significantly
		// speeds up voluntary leader transitions as the new leader don't have to wait
		// LeaseDuration time first.
		//
		// In the default scaffold provided, the program ends immediately after
		// the manager stops, so would be fine to enable this option. However,
		// if you are doing or is intended to do any operation such as perform cleanups
		// after the manager stops then its usage might be unsafe.
		// LeaderElectionReleaseOnCancel: true,
	})
	if err != nil {
		setupLog.Error(err, "unable to start manager")
		os.Exit(1)
	}

	ctx := ctrl.SetupSignalHandler()

	setupLog.Info("setting up the storage client")
	gscClient, err := storage.NewClient(ctx)
	if err != nil {
		setupLog.Error(err, "unable to start the Google cloud storage client")
	}

	setupLog.Info("setting up leader election")
	resourceLock := &resourcelock.LeaseLock{
		Client:     gscClient,
		BucketName: leaseBucket,
		ProjectId:  projectId,
	}
	l, err := leaderelection.NewLeaderElector(leaderelection.LeaderElectionConfig{
		Lock:          resourceLock,
		LeaseDuration: duration,
		RenewDeadline: renew,
		RetryPeriod:   retry,
	})

	if err != nil {
		setupLog.Error(err, "failed to enable leader election")
		os.Exit(1)
	}

	if err = (&controller.LeaseReconciler{
		Client:        mgr.GetClient(),
		Scheme:        mgr.GetScheme(),
		Identify:      mgr.GetConfig().Host,
		LeaderElector: l,
	}).SetupWithManager(mgr); err != nil {
		setupLog.Error(err, "unable to create controller", "controller", "Lease")
		os.Exit(1)
	}
	//+kubebuilder:scaffold:builder

	if err := mgr.AddHealthzCheck("healthz", healthz.Ping); err != nil {
		setupLog.Error(err, "unable to set up health check")
		os.Exit(1)
	}
	if err := mgr.AddReadyzCheck("readyz", healthz.Ping); err != nil {
		setupLog.Error(err, "unable to set up ready check")
		os.Exit(1)
	}

	setupLog.Info("starting manager")
	if err := mgr.Start(ctx); err != nil {
		setupLog.Error(err, "problem running manager")
		os.Exit(1)
	}
}
