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

	"sigs.k8s.io/kubebuilder-declarative-pattern/pkg/restmapper"
	//+kubebuilder:scaffold:imports
)

type Operator struct {
	LeaderElectionID string

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

func (o *Operator) RunMain() {
	ctx := ctrl.SetupSignalHandler()
	if err := o.run(ctx); err != nil {
		fmt.Fprintf(os.Stderr, "%v\n", err)
		os.Exit(1)
	}
}

func (o *Operator) run(ctx context.Context) error {
	var metricsAddr string
	var enableLeaderElection bool
	var probeAddr string
	flag.StringVar(&metricsAddr, "metrics-bind-address", ":8080", "The address the metric endpoint binds to.")
	flag.StringVar(&probeAddr, "health-probe-bind-address", ":8081", "The address the probe endpoint binds to.")
	flag.BoolVar(&enableLeaderElection, "leader-elect", false,
		"Enable leader election for controller manager. "+
			"Enabling this will ensure there is only one active controller manager.")
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
		Scheme:             scheme,
		MetricsBindAddress: metricsAddr,
		// Metrics: server.Options{
		//     	BindAddress:     metricsAddr,
		// }
		Port:                   9443,
		HealthProbeBindAddress: probeAddr,
		LeaderElection:         enableLeaderElection,
		LeaderElectionID:       o.LeaderElectionID,

		// MapperProvider provides the rest mapper used to map go types to Kubernetes APIs
		MapperProvider: restmapper.NewControllerRESTMapper,
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
