// Copyright 2026 Google LLC
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
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
	krmv1alpha1 "github.com/gke-labs/kube-etl/syncer/api/v1alpha1"
	"os"
	"time"

	"k8s.io/klog/v2"

	flag "github.com/spf13/pflag"
	_ "k8s.io/client-go/plugin/pkg/client/auth"

	"k8s.io/apimachinery/pkg/runtime"
	utilruntime "k8s.io/apimachinery/pkg/util/runtime"
	clientgoscheme "k8s.io/client-go/kubernetes/scheme"
	ctrl "sigs.k8s.io/controller-runtime"
	"sigs.k8s.io/controller-runtime/pkg/cache"

	krmSyncer "github.com/gke-labs/kube-etl/syncer/controllers"
)

var (
	scheme   = runtime.NewScheme()
	setupLog = ctrl.Log.WithName("setup")
)

func init() {
	utilruntime.Must(clientgoscheme.AddToScheme(scheme))

	utilruntime.Must(krmv1alpha1.AddToScheme(scheme))
	//+kubebuilder:scaffold:scheme
}

func main() {
	if err := run(context.Background()); err != nil {
		fmt.Fprintf(os.Stderr, "%v\n", err)
		os.Exit(1)
	}
}

func run(ctx context.Context) error {
	var syncPeriod time.Duration
	flag.DurationVar(&syncPeriod, "sync-period", 0, "The minimum frequency at which watched resources are reconciled (0 or negative disables periodic sync)")
	var maxConcurrentReconciles int
	flag.IntVar(&maxConcurrentReconciles, "max-concurrent-reconciles", krmSyncer.DefaultMaxConcurrentReconciles, "Maximum number of concurrent reconciles for dynamic controllers.")

	// Configure logging
	klogFlagSet := goflag.NewFlagSet("klog", goflag.ExitOnError)
	klog.InitFlags(klogFlagSet)
	// Support default klog verbosity `-v`
	flag.CommandLine.AddGoFlag(klogFlagSet.Lookup("v"))
	flag.CommandLine.AddGoFlagSet(goflag.CommandLine)
	flag.Parse()

	ctx = klog.NewContext(ctx, setupLog)
	ctrl.SetLogger(klog.NewKlogr())

	if maxConcurrentReconciles < 1 {
		return fmt.Errorf("--max-concurrent-reconciles must be at least 1, got %d", maxConcurrentReconciles)
	}

	var syncPeriodOpt *time.Duration
	if syncPeriod > 0 {
		syncPeriodOpt = &syncPeriod
	}

	mgrOpts := ctrl.Options{
		Scheme: scheme,
	}
	if syncPeriodOpt != nil {
		mgrOpts.Cache = cache.Options{
			SyncPeriod: syncPeriodOpt,
		}
	}

	setupLog.Info("Creating manager")
	mgr, err := ctrl.NewManager(ctrl.GetConfigOrDie(), mgrOpts)
	if err != nil {
		setupLog.Error(err, "unable to start manager")
		return err
	}

	// Create and set up the SyncerReconciler
	setupLog.Info("Creating SyncerReconciler")
	if err = (&krmSyncer.KRMSyncerReconciler{
		Client:                  mgr.GetClient(),
		Scheme:                  mgr.GetScheme(),
		Manager:                 mgr,
		SyncPeriod:              syncPeriodOpt,
		MaxConcurrentReconciles: maxConcurrentReconciles,
	}).SetupWithManager(mgr); err != nil {
		setupLog.Error(err, "unable to create controller", "controller", "Syncer")
		return err
	}

	setupLog.Info("starting manager")
	if err = mgr.Start(ctrl.SetupSignalHandler()); err != nil {
		setupLog.Error(err, "problem running manager")
		return err
	}
	return nil
}
