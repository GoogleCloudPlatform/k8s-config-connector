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

package testcontroller

import (
	"context"
	"fmt"
	"log"
	"math/rand"
	"regexp"
	"sync"
	"testing"
	"time"

	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/kccmanager/nocache"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/crd/crdloader"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/k8s"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/test"
	testgcp "github.com/GoogleCloudPlatform/k8s-config-connector/pkg/test/gcp"
	cnrmwebhook "github.com/GoogleCloudPlatform/k8s-config-connector/pkg/webhook"

	corev1 "k8s.io/api/core/v1"
	apiextensions "k8s.io/apiextensions-apiserver/pkg/apis/apiextensions/v1"
	"k8s.io/apimachinery/pkg/api/errors"
	apierrors "k8s.io/apimachinery/pkg/api/errors"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/types"
	"k8s.io/klog/v2"
	"sigs.k8s.io/controller-runtime/pkg/client"
	"sigs.k8s.io/controller-runtime/pkg/envtest"
	"sigs.k8s.io/controller-runtime/pkg/manager"
	"sigs.k8s.io/controller-runtime/pkg/reconcile"
	"sigs.k8s.io/controller-runtime/pkg/webhook"
)

var (
	ExpectedSuccessfulReconcileResult   = reconcile.Result{RequeueAfter: k8s.MeanReconcileReenqueuePeriod}
	ExpectedUnsuccessfulReconcileResult = reconcile.Result{Requeue: false, RequeueAfter: 0 * time.Minute}
	ExpectedRequeueReconcileStruct      = reconcile.Result{Requeue: true}
)

// StartTestManager begins a new test manager, and returns a function
// to gracefully shutdown.
func StartTestManagerInstance(env *envtest.Environment, testType test.TestType, whCfgs []cnrmwebhook.WebhookConfig) (manager.Manager, func()) {
	mgr, stopFunc, err := startTestManager(env, testType, whCfgs)
	if err != nil {
		log.Fatal(err)
	}
	return mgr, stopFunc
}

func startTestManager(env *envtest.Environment, testType test.TestType, whCfgs []cnrmwebhook.WebhookConfig) (manager.Manager, func(), error) {
	mgr, err := manager.New(env.Config, manager.Options{
		Port:    env.WebhookInstallOptions.LocalServingPort,
		Host:    env.WebhookInstallOptions.LocalServingHost,
		CertDir: env.WebhookInstallOptions.LocalServingCertDir,
		// supply a concrete client to disable the default behavior of caching
		NewClient: nocache.NoCacheClientFunc,
		// Disable metrics server for testing
		MetricsBindAddress: "0",
	})
	if err != nil {
		return nil, nil, fmt.Errorf("error creating manager: %v", err)
	}
	if testType == test.IntegrationTestType {
		server := mgr.GetWebhookServer()
		for _, cfg := range whCfgs {
			server.Register(cfg.Path, &webhook.Admission{Handler: cfg.Handler})
		}
	}
	stop := startMgr(mgr, log.Fatalf)
	return mgr, stop, nil
}

func StartMgr(t *testing.T, mgr manager.Manager) func() {
	return startMgr(mgr, t.Fatalf)
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
			mgrStartErrHandler("unable to start manager: %v", err)
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

// RunReconcilerAssertResults asserts the expected state of the reconciler run.
func RunReconcilerAssertResults(t *testing.T, reconciler reconcile.Reconciler, objectMeta v1.ObjectMeta,
	expectedResult reconcile.Result, expectedErrorRegex *regexp.Regexp) {
	t.Helper()
	reconcileRequest := reconcile.Request{NamespacedName: k8s.GetNamespacedName(objectMeta.GetObjectMeta())}
	result, err := reconciler.Reconcile(context.Background(), reconcileRequest)
	if expectedErrorRegex == nil {
		if err != nil {
			t.Fatalf("reconcile returned unexpected error: %v", err)
		}
	} else {
		if err == nil || !expectedErrorRegex.MatchString(err.Error()) {
			t.Fatalf("error '%v' does not match regex '%v'", err, expectedErrorRegex)
		}
	}
	if !(requeueEqualAndRequeueAfterWithinBoundsOfMean(result, expectedResult)) {
		t.Fatalf("reconcile result mismatch: got '%v', want within %v of '%v'", result, k8s.MeanReconcileReenqueuePeriod/2, expectedResult)
	}
}

func GetCRDForKind(t *testing.T, kubeClient client.Client, kind string) *apiextensions.CustomResourceDefinition {
	t.Helper()
	c, err := crdloader.GetCRDForKind(kind)
	if err != nil {
		t.Fatal(err)
	}
	return c
}

func SetupNamespaceForDefaultProject(t *testing.T, c client.Client, name string) {
	projectID := testgcp.GetDefaultProjectID(t)
	SetupNamespaceForProject(t, c, name, projectID)
}

func SetupNamespaceForProject(t *testing.T, c client.Client, name, projectID string) {
	EnsureNamespaceExistsT(t, c, name)
	EnsureNamespaceHasProjectIDAnnotation(t, c, name, projectID)
}

func EnsureNamespaceExists(c client.Client, name string) error {
	ns := &corev1.Namespace{}
	ns.SetName(name)
	if err := c.Create(context.Background(), ns); err != nil {
		if !errors.IsAlreadyExists(err) {
			return fmt.Errorf("error creating namespace %v: %v", name, err)
		}
	}
	return nil
}

func EnsureNamespaceExistsT(t *testing.T, c client.Client, name string) {
	t.Helper()
	if err := EnsureNamespaceExists(c, name); err != nil {
		t.Fatal(err)
	}
}

func EnsureNamespaceHasProjectIDAnnotation(t *testing.T, c client.Client, namespaceName, projectId string) {
	t.Helper()
	err := createNamespaceProjectIdAnnotation(context.TODO(), c, namespaceName, projectId)
	if err != nil {
		t.Fatal(err)
	}
}

func createNamespaceProjectIdAnnotation(ctx context.Context, c client.Client, namespaceName, projectId string) error {
tryAgain:
	attempt := 0
	var ns corev1.Namespace
	if err := c.Get(ctx, types.NamespacedName{Name: namespaceName}, &ns); err != nil {
		return fmt.Errorf("error getting namespace %q: %w", namespaceName, err)
	}
	if val, ok := k8s.GetAnnotation(k8s.ProjectIDAnnotation, &ns); ok {
		if val == projectId {
			klog.Infof("namespace %q already has project id annotation value %q", namespaceName, projectId)
			return nil
		} else {
			return fmt.Errorf("cannot set project id annotatation value to %q: the annotation already contained a value of %q",
				projectId, val)
		}
	}
	k8s.SetAnnotation(k8s.ProjectIDAnnotation, projectId, &ns)
	err := c.Update(ctx, &ns)
	if err != nil {
		if apierrors.IsConflict(err) {
			attempt++
			if attempt < 10 {
				klog.Warningf("detected concurrent modification error updating namespace %q, will retry", namespaceName)
				time.Sleep(time.Millisecond * time.Duration(rand.Intn(1000)))
				goto tryAgain
			}
		}
		return fmt.Errorf("error setting project id on namespace %q: %w", namespaceName, err)
	}
	return nil
}

func requeueEqualAndRequeueAfterWithinBoundsOfMean(result reconcile.Result, expectedResult reconcile.Result) bool {
	requeueEqual := result.Requeue == expectedResult.Requeue
	lowerBound := expectedResult.RequeueAfter - k8s.MeanReconcileReenqueuePeriod/2
	upperBound := expectedResult.RequeueAfter + k8s.MeanReconcileReenqueuePeriod/2
	return requeueEqual && result.RequeueAfter >= lowerBound && result.RequeueAfter < upperBound
}
