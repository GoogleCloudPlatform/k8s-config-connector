// Copyright 2023 Google LLC
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

package commonoperator

import (
	"context"
	"flag"
	"fmt"
	"os"

	// Import all Kubernetes client auth plugins (e.g. Azure, GCP, OIDC, etc.)
	// to ensure that exec-entrypoint and run can make use of them.
	_ "k8s.io/client-go/plugin/pkg/client/auth"
	"k8s.io/klog/v2"
	"k8s.io/klog/v2/klogr"

	"k8s.io/apimachinery/pkg/runtime"
	clientgoscheme "k8s.io/client-go/kubernetes/scheme"
	ctrl "sigs.k8s.io/controller-runtime"
	"sigs.k8s.io/controller-runtime/pkg/healthz"
	"sigs.k8s.io/controller-runtime/pkg/manager"
	metricsserver "sigs.k8s.io/controller-runtime/pkg/metrics/server"
	//+kubebuilder:scaffold:imports
)

type Options struct {
	LeaderElectionID string
}
type Operator struct {
	schemas     []AddToSchemaFunc
	reconcilers []Reconciler
}

type Reconciler interface {
	SetupWithManager(mgr manager.Manager) error
}

type AddToSchemaFunc func(schema *runtime.Scheme) error

func (o *Operator) RegisterSchema(schema AddToSchemaFunc) {
	o.schemas = append(o.schemas, schema)
}

func (o *Operator) RegisterReconciler(reconciler Reconciler) {
	o.reconcilers = append(o.reconcilers, reconciler)
}

func (o *Operator) RunAsMain(opt Options) {
	ctx := ctrl.SetupSignalHandler()
	if err := o.Run(ctx, opt); err != nil {
		fmt.Fprintf(os.Stderr, "%v\n", err)
		os.Exit(1)
	}
}

func (o *Operator) Run(ctx context.Context, opt Options) error {
	metricsAddr := ":8080"
	enableLeaderElection := true
	if opt.LeaderElectionID == "" {
		enableLeaderElection = false
		klog.Warningf("running without leader election")
	} else {
		flag.BoolVar(&enableLeaderElection, "leader-elect", enableLeaderElection,
			"Enable leader election for controller manager. "+
				"Enabling this will ensure there is only one active controller manager.")
	}
	probeAddr := ":8081"
	flag.StringVar(&metricsAddr, "metrics-bind-address", metricsAddr, "The address the metric endpoint binds to.")
	flag.StringVar(&probeAddr, "health-probe-bind-address", probeAddr, "The address the probe endpoint binds to.")
	klog.InitFlags(nil)
	flag.Parse()

	ctrl.SetLogger(klogr.New())

	scheme := runtime.NewScheme()

	log := ctrl.Log.WithName("setup")

	if err := clientgoscheme.AddToScheme(scheme); err != nil {
		return fmt.Errorf("error initializing client-go scheme: %w", err)
	}

	for _, schemaFunc := range o.schemas {
		if err := schemaFunc(scheme); err != nil {
			return fmt.Errorf("error initializing schema: %w", err)
		}
	}

	mgr, err := ctrl.NewManager(ctrl.GetConfigOrDie(), ctrl.Options{
		Scheme: scheme,
		Metrics: metricsserver.Options{
			BindAddress: metricsAddr,
		},
		HealthProbeBindAddress: probeAddr,
		LeaderElection:         enableLeaderElection,
		LeaderElectionID:       opt.LeaderElectionID,

		// // MapperProvider provides the rest mapper used to map go types to Kubernetes APIs
		// MapperProvider: restmapper.NewControllerRESTMapper,
	})
	if err != nil {
		return fmt.Errorf("error creating manager: %w", err)
	}

	for _, reconciler := range o.reconcilers {
		if err = reconciler.SetupWithManager(mgr); err != nil {
			return fmt.Errorf("error creating controller %T: %w", reconciler, err)
		}
	}

	//+kubebuilder:scaffold:builder

	if err := mgr.AddHealthzCheck("healthz", healthz.Ping); err != nil {
		return fmt.Errorf("error setting up health check: %w", err)
	}
	if err := mgr.AddReadyzCheck("readyz", healthz.Ping); err != nil {
		return fmt.Errorf("error setting up ready check: %w", err)
	}

	log.Info("starting manager")
	if err := mgr.Start(ctx); err != nil {
		return fmt.Errorf("error running manager: %w", err)
	}
	return nil
}
