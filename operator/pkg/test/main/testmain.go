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

package testmain

import (
	"context"
	"fmt"
	"log"
	"sync"
	"time"

	appsv1 "k8s.io/api/apps/v1"
	corev1 "k8s.io/api/core/v1"
	apiextensions "k8s.io/apiextensions-apiserver/pkg/apis/apiextensions/v1"
	"k8s.io/client-go/kubernetes/scheme"
	"k8s.io/client-go/rest"
	"sigs.k8s.io/controller-runtime/pkg/envtest"
	"sigs.k8s.io/controller-runtime/pkg/manager"
	metricsserver "sigs.k8s.io/controller-runtime/pkg/metrics/server"

	customizev1beta1 "github.com/GoogleCloudPlatform/k8s-config-connector/operator/pkg/apis/core/customize/v1beta1"
	corev1beta1 "github.com/GoogleCloudPlatform/k8s-config-connector/operator/pkg/apis/core/v1beta1"
	"github.com/GoogleCloudPlatform/k8s-config-connector/operator/pkg/controllers"
	"github.com/GoogleCloudPlatform/k8s-config-connector/operator/pkg/test/util/paths"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/kccmanager/nocache"
)

func init() {
	s := scheme.Scheme
	if err := apiextensions.SchemeBuilder.AddToScheme(s); err != nil {
		log.Fatalf("error registering apiextensions v1beta1 scheme: %v", err)
	}
	if err := corev1beta1.SchemeBuilder.AddToScheme(s); err != nil {
		log.Fatalf("error registering core kcc operator scheme: %v", err)
	}
	if err := corev1.SchemeBuilder.AddToScheme(s); err != nil {
		log.Fatalf("error registering core v1 scheme: %v", err)
	}
	if err := appsv1.SchemeBuilder.AddToScheme(s); err != nil {
		log.Fatalf("error registering apps v1 scheme: %v", err)
	}
	if err := customizev1beta1.SchemeBuilder.AddToScheme(s); err != nil {
		log.Fatalf("error registering kcc customization scheme: %v", err)
	}
}

// startTestEnv starts a local K8S API server to run unit tests. Tests using
// this function do not require an external API server to execute.
func startTestEnv() (*rest.Config, func()) {
	testEnv := &envtest.Environment{
		CRDDirectoryPaths:        []string{paths.GetOperatorCRDsPath()},
		ControlPlaneStartTimeout: time.Minute,
		ControlPlaneStopTimeout:  time.Minute,
	}
	var err error
	cfg, err := testEnv.Start()
	if err != nil {
		log.Fatal(err)
	}
	stop := func() {
		if err := testEnv.Stop(); err != nil {
			log.Printf("unable to stop the test environment: %v", err)
		}
	}
	return cfg, stop
}

func StartTestManager(cfg *rest.Config) (manager.Manager, func(), error) {
	scheme := controllers.BuildScheme()

	opts := manager.Options{
		Metrics: metricsserver.Options{
			// Prevent manager from binding to a port to serve prometheus metrics
			// since creating multiple managers for tests will fail if more than
			// one manager tries to bind to the same port.
			BindAddress: "0",
		},
		// Prevent manager from binding to a port to serve health probes since
		// creating multiple managers for tests will fail if more than one
		// manager tries to bind to the same port.
		HealthProbeBindAddress: "0",
		Scheme:                 scheme,
	}
	// Supply a concrete client to disable the default behavior of caching
	nocache.TurnOffAllCaching(&opts)

	mgr, err := manager.New(cfg, opts)
	if err != nil {
		return nil, nil, fmt.Errorf("error creating manager: %w", err)
	}
	stopFunc := startMgr(mgr, log.Fatalf)
	return mgr, stopFunc, nil
}

func startMgr(mgr manager.Manager, mgrStartErrHandler func(string, ...interface{})) func() {
	ctx, cancel := context.WithCancel(context.TODO())
	// it is important to wait for the below goroutine to terminate before attempting to exit the application because
	// otherwise there can be panics and unexpected behavior while the manager is shutting down
	wg := sync.WaitGroup{}
	wg.Add(1)
	go func() {
		defer wg.Done()
		if err := mgr.Start(ctx); err != nil {
			mgrStartErrHandler("unable to start manager: %w", err)
		}
	}()
	stop := func() {
		// calling cancel() will cancel the context 'ctx', the mgr will stop all runnables and Start() will return and
		// the above goroutine will exit
		cancel()
		// wait for the goroutine above to exit (it has a deferred wg.Done())
		wg.Wait()
	}
	return stop
}

func StartTestManagerFromNewTestEnv() (manager.Manager, func()) {
	cfg, stopEnv := startTestEnv()
	mgr, stopMgr, err := StartTestManager(cfg)
	if err != nil {
		log.Fatal(err)
	}
	stop := func() {
		stopMgr()
		stopEnv()
	}
	return mgr, stop
}
