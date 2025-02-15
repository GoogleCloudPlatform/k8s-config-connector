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
	"errors"
	"fmt"
	"log"
	"math/rand"
	"regexp"
	"strings"
	"sync"
	"testing"
	"time"

	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/kccmanager/nocache"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/k8s"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/test"
	testgcp "github.com/GoogleCloudPlatform/k8s-config-connector/pkg/test/gcp"
	cnrmwebhook "github.com/GoogleCloudPlatform/k8s-config-connector/pkg/webhook"

	corev1 "k8s.io/api/core/v1"
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

const (
	// transientErrorsMaxRetries sets the max number of retries on a transient error
	transientErrorsMaxRetries = 5
	// transientErrorsRetryInterval sets the interval between retries on a transient error
	transientErrorsRetryInterval = 5 * time.Second
)

// StartTestManager begins a new test manager, and returns a function
// to gracefully shutdown.
func StartTestManagerInstance(env *envtest.Environment, testType test.Type, whCfgs []cnrmwebhook.Config) (manager.Manager, func()) {
	mgr, stopFunc, err := startTestManager(env, testType, whCfgs)
	if err != nil {
		log.Fatal(err)
	}
	return mgr, stopFunc
}

func startTestManager(env *envtest.Environment, testType test.Type, whCfgs []cnrmwebhook.Config) (manager.Manager, func(), error) {
	opts := manager.Options{
		Port:    env.WebhookInstallOptions.LocalServingPort,
		Host:    env.WebhookInstallOptions.LocalServingHost,
		CertDir: env.WebhookInstallOptions.LocalServingCertDir,
		// Disable metrics server for testing
		MetricsBindAddress: "0",
	}
	// supply a concrete client to disable the default behavior of caching
	nocache.OnlyCacheCCAndCCC(&opts)

	mgr, err := manager.New(env.Config, opts)
	if err != nil {
		return nil, nil, fmt.Errorf("error creating manager: %w", err)
	}
	if testType == test.IntegrationTestType {
		server := mgr.GetWebhookServer()
		for _, cfg := range whCfgs {
			handler := cfg.HandlerFunc(mgr)
			server.Register(cfg.Path, &webhook.Admission{Handler: handler})
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

// isTransientError reports whether the reconciler error is a random "flake" and we should retry.
func isTransientError(t *testing.T, err error) bool {
	if err == nil {
		return false
	}

	// Print the chain so we don't have to use string matching for future errors
	var chain []string
	current := err
	for {
		chain = append(chain, fmt.Sprintf("[%T: %+v]", current, current))
		current = errors.Unwrap(current)
		if current == nil {
			break
		}
	}

	errorMessage := err.Error()

	// Permission denied errors are considered transient
	// We don't know the exact error currently, use string matching for now...
	//
	// Example error:
	// {"severity":"info","timestamp":"2022-12-06T20:27:32.799Z","logger":"iapidentityawareproxyclient-controller","msg":"creating/updating underlying resource","resource":{"namespace":"jcjjsgqldlbw7hcvseiq","name":"iapidentityawareproxyclient-jcjjsgqldlbw7hcvseiq"}}
	// W1206 20:27:35.461665  113200 logger.go:58] [DCL WARNING] [RequestID:km5nd0fv]  get returned error: googleapi: Error 403: The caller does not have permission
	// {"severity":"error","timestamp":"2022-12-06T20:27:35.461Z","logger":"iapidentityawareproxyclient-controller","msg":"error applying desired state","resource":{"namespace":"jcjjsgqldlbw7hcvseiq","name":"iapidentityawareproxyclient-jcjjsgqldlbw7hcvseiq"},"error":"googleapi: Error 403: The caller does not have permission"}
	// dynamic_controller_integration_test.go:190: reconcile returned unexpected error: Update call failed: error applying desired state: googleapi: Error 403: The caller does not have permission
	if strings.Contains(errorMessage, "The caller does not have permission") {
		t.Logf("permission error found; considered transient; chain is %v", chain)
		return true
	}

	// Internal errors are considered transient
	// We don't know the exact error currently, use string matching for now...
	//
	// Example error:
	// chain is [[*errors.errorString: Delete call failed: error deleting resource: [{0 Error when reading or editing Project Service projects/clienttls-244gvcgzxgegwhmfvqgq/services/: Error disabling service "networksecurity.googleapis.com" for project "clienttls-244gvcgzxgegwhmfvqgq": Error waiting for api to disable: Error code 13, message: [An internal exception occurred.  Help Token: AZWD64pDMtDdLt4XOiuQgfBiJS-s2K6hSHg4cKv6GBl2Wibfb_wEnkl8HZjT7unqZSibwlNEmXpHwJ3AFbmfidKSWtWc9CtNL15HcR53H0ETgtB8] with failed services [networksecurity.googleapis.com]  []}]]]
	//testreconciler.go:96: reconcile returned unexpected error: Delete call failed: error deleting resource: [{0 Error when reading or editing Project Service projects/clienttls-244gvcgzxgegwhmfvqgq/services/: Error disabling service "networksecurity.googleapis.com" for project "clienttls-244gvcgzxgegwhmfvqgq": Error waiting for api to disable: Error code 13, message: [An internal exception occurred.
	if strings.Contains(errorMessage, "An internal exception occurred") {
		t.Logf("internal error found; considered transient; chain is %v", chain)
		return true
	}

	// "is not ready" errors are considered transient
	// We don't know the exact error currently, use string matching for now...
	//
	// Example error:
	// reconcile.go:164: error was not considered transient; chain is [[*errors.errorString: Update call failed: error applying desired state: operation received error: error code "3", message: The resource 'projects/cnrm-test-mqtuo70y3lg3w1m7/regions/us-central1/subnetworks/default' is not ready, details: []
	// details: map[]]]
	if strings.Contains(errorMessage, "is not ready") {
		t.Logf("internal error found; considered transient; chain is %v", chain)
		return true
	}

	// "missing permission on" errors are considered transient
	// We don't know the exact error currently, use string matching for now...
	//
	// Example error:
	// reconcile.go:175: error was not considered transient; chain is [[*errors.errorString: Update call failed: error applying desired state: summary: failed prerequisites: missing permission on "billingAccounts/0162D7-7B0CB6-ED962E": billing.resourceAssociations.create]]
	if strings.Contains(errorMessage, "missing permission on") {
		t.Logf("internal error found; considered transient; chain is %v", chain)
		return true
	}

	// "Hook call/poll failed for service" errors are considered transient
	// We don't know the exact error currently, use string matching for now...
	//
	// Example error:
	// testreconciler.go:96: reconcile returned unexpected error: Delete call failed: error deleting resource: [{0 Error when reading or editing Project Service projects/clienttls-aaoksjdrfqbos22kkhaa/services/: Error disabling service "networksecurity.googleapis.com" for project "clienttls-aaoksjdrfqbos22kkhaa": Error waiting for api to disable: Error code 8, message: [Hook call/poll failed for service "networksecurity.googleapis.com".
	// Help Token: AZWD64q7zyHTI4hHRS7MG0gHM4T8nMAXsiKCMAohWDFWVzK5BIZes3oQScpmnmkpTBlr0T9zldAZZuOWsjgv7BdRwGCGoOFdr2KqNqOarqlffbV3] with failed services [networksecurity.googleapis.com]  []}]
	if strings.Contains(errorMessage, "Hook call/poll failed for service") {
		t.Logf("internal error found; considered transient; chain is %v", chain)
		return true
	}

	t.Logf("error was not considered transient; chain is %v", chain)
	return false
}

// RunReconcilerAssertResults asserts the expected state of the reconciler run.
func RunReconcilerAssertResults(ctx context.Context, t *testing.T, reconciler reconcile.Reconciler,
	kind string, objectMeta v1.ObjectMeta,
	expectedResult reconcile.Result, expectedErrorRegex *regexp.Regexp) {
	attempt := 0
tryAgain:
	attempt++
	t.Helper()
	startTime := time.Now()
	t.Logf("starting reconcile for %v:%v/%v", kind, objectMeta.Namespace, objectMeta.Name)
	reconcileRequest := reconcile.Request{NamespacedName: k8s.GetNamespacedName(objectMeta.GetObjectMeta())}
	result, err := reconciler.Reconcile(ctx, reconcileRequest)
	t.Logf("reconcile for %v:%v/%v took %v, result was (%v, %v)",
		kind, objectMeta.Namespace, objectMeta.Name, time.Since(startTime), result, err)
	// Retry if we see a "transient" error (up to our retry limit)
	if err != nil {
		if isTransientError(t, err) {
			if attempt < transientErrorsMaxRetries {
				t.Logf("detected transient error, will retry: %v", err)
				time.Sleep(transientErrorsRetryInterval)
				goto tryAgain
			}

			t.Logf("detected transient error, but maximum number of retries reached: %v", err)
		}
	}

	if expectedErrorRegex == nil {
		if err != nil {
			t.Fatal(fmt.Errorf("reconcile returned unexpected error: %w", err))
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
		if !apierrors.IsAlreadyExists(err) {
			return fmt.Errorf("error creating namespace %v: %w", name, err)
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

func EnsureNamespaceHasProjectIDAnnotation(t *testing.T, c client.Client, namespaceName, projectID string) {
	t.Helper()
	err := createNamespaceprojectIDAnnotation(context.TODO(), c, namespaceName, projectID)
	if err != nil {
		t.Fatal(err)
	}
}

func createNamespaceprojectIDAnnotation(ctx context.Context, c client.Client, namespaceName, projectID string) error {
tryAgain:
	attempt := 0
	var ns corev1.Namespace
	if err := c.Get(ctx, types.NamespacedName{Name: namespaceName}, &ns); err != nil {
		return fmt.Errorf("error getting namespace %q: %w", namespaceName, err)
	}
	if val, ok := k8s.GetAnnotation(k8s.ProjectIDAnnotation, &ns); ok {
		if val == projectID {
			klog.Infof("namespace %q already has project id annotation value %q", namespaceName, projectID)
			return nil
		}

		return fmt.Errorf("cannot set project id annotatation value to %q: the annotation already contained a value of %q", projectID, val)
	}
	k8s.SetAnnotation(k8s.ProjectIDAnnotation, projectID, &ns)
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
	lowerBound := expectedResult.RequeueAfter / 2
	upperBound := expectedResult.RequeueAfter / 2 * 3
	return requeueEqual && (result.RequeueAfter >= lowerBound && result.RequeueAfter < upperBound || result.RequeueAfter == 0 && expectedResult.RequeueAfter == 0)
}
