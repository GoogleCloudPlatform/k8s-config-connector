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

//go:build integration
// +build integration

package gsakeysecretgenerator

import (
	"context"
	"fmt"
	"log"
	"os"
	"path/filepath"
	"strings"
	"testing"
	"time"

	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/dynamic"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/jitter"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/tf"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/gcp"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/k8s"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/stateintospec"
	testcontroller "github.com/GoogleCloudPlatform/k8s-config-connector/pkg/test/controller"
	testgcp "github.com/GoogleCloudPlatform/k8s-config-connector/pkg/test/gcp"
	testjitter "github.com/GoogleCloudPlatform/k8s-config-connector/pkg/test/jitter"
	testk8s "github.com/GoogleCloudPlatform/k8s-config-connector/pkg/test/k8s"
	testmain "github.com/GoogleCloudPlatform/k8s-config-connector/pkg/test/main"
	testvariable "github.com/GoogleCloudPlatform/k8s-config-connector/pkg/test/resourcefixture/variable"
	testservicemappingloader "github.com/GoogleCloudPlatform/k8s-config-connector/pkg/test/servicemappingloader"
	tfprovider "github.com/GoogleCloudPlatform/k8s-config-connector/pkg/tf/provider"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/util/repo"

	"github.com/ghodss/yaml"
	tfschema "github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"golang.org/x/sync/semaphore"
	"google.golang.org/api/iam/v1"
	v1 "k8s.io/api/core/v1"
	"k8s.io/apimachinery/pkg/apis/meta/v1/unstructured"
	"k8s.io/apimachinery/pkg/util/wait"
	_ "k8s.io/client-go/plugin/pkg/client/auth/gcp"
	"sigs.k8s.io/controller-runtime/pkg/client"
	"sigs.k8s.io/controller-runtime/pkg/event"
	"sigs.k8s.io/controller-runtime/pkg/manager"
	"sigs.k8s.io/controller-runtime/pkg/reconcile"
)

var (
	mgr            manager.Manager
	expectedResult = reconcile.Result{}
)

func TestServiceAccountKey(t *testing.T) {
	t.Helper()
	kubeClient := mgr.GetClient()
	provider := tfprovider.NewOrLogFatal(tfprovider.DefaultConfig)
	ctx := context.TODO()
	uuid := testvariable.NewUniqueID()
	project := testgcp.GetDefaultProjectID(t)
	iamClient := testgcp.NewIAMClient(t)
	testcontroller.SetupNamespaceForDefaultProject(t, kubeClient, project)

	saReconciler := newTestReconciler(t, mgr, repo.GetServiceAccountCRDPath(), provider)
	sakReconciler := newTestReconciler(t, mgr, repo.GetServiceAccountKeyCRDPath(), provider)
	generator := newSecretGenerator(t, mgr, repo.GetServiceAccountKeyCRDPath())
	// create the dependent service account
	sa := convertToUnstructAndReplaceName(t, uuid, project, "", "testdata/gsa.yaml")
	if err := kubeClient.Create(ctx, sa); err != nil {
		t.Fatalf("couldn't create google service account %v: %v", sa.GetName(), err)
	}
	saName := k8s.GetNamespacedName(sa)
	saRequest := reconcile.Request{NamespacedName: saName}
	if _, err := saReconciler.Reconcile(context.TODO(), saRequest); err != nil {
		t.Fatalf("error reconciling iamserviceaccount %v: %v", saName, err)
	}
	if err := kubeClient.Get(ctx, saName, sa); err != nil {
		t.Fatalf("error getting iamserviceaccount %v: %v", saName, err)
	}
	// create a service account key sample
	gsakey := convertToUnstructAndReplaceName(t, uuid, project, sa.GetName(), "testdata/gsakey.yaml")
	if err := kubeClient.Create(ctx, gsakey); err != nil {
		t.Fatalf("couldn't create google service account key %v: %v", sa.GetName(), err)
	}
	sakName := k8s.GetNamespacedName(gsakey)
	sakRequest := reconcile.Request{NamespacedName: sakName}
	if _, err := sakReconciler.Reconcile(context.TODO(), sakRequest); err != nil {
		t.Fatalf("error reconciling iamserviceaccountkey %v: %v", sakName, err)
	}
	if err := kubeClient.Get(ctx, sakName, gsakey); err != nil {
		t.Fatalf("error getting iamserviceaccountkey %v: %v", sakName, err)
	}
	// check the status and make sure the underlying resource is created
	keyName, found, err := unstructured.NestedString(gsakey.Object, "status", "name")
	if !found || err != nil {
		t.Fatalf("couldn't find name from %v status: %v", gsakey.GetName(), err)
	}

	// Wait for key to propagate
	// Per https://cloud.google.com/iam/docs/creating-managing-service-account-keys:
	// "After you create a key, you might need to wait for 60 seconds or more before you perform another operation with the key."
	var lastErr error
	if err := wait.PollImmediate(10*time.Second, 2*time.Minute, func() (bool, error) {
		if _, err := iamClient.Projects.ServiceAccounts.Keys.Get(keyName).Do(); err != nil {
			lastErr = err
			return false, nil
		} else {
			return true, nil
		}
	}); err != nil {
		t.Fatalf("error calling iam service to get the service account key %v (despite polling, lastErr=%v): %v", keyName, lastErr, err)
	}

	// invoke the secret generator
	if _, err := generator.Reconcile(context.TODO(), sakRequest); err != nil {
		t.Fatalf("error reconciling iamserviceaccountkey %v to create a secret: %v", sakName, err)
	}
	// check the event about the service account key
	secretCreated := false
	eventList := &v1.EventList{}
	if err := kubeClient.List(context.TODO(), eventList, &client.ListOptions{Namespace: gsakey.GetNamespace()}); err != nil {
		t.Fatalf("unable to list objects: %v", err)
	}
	events := testcontroller.CollectEvents(t, mgr.GetConfig(), gsakey.GetNamespace(), 5, 5*time.Second)
	for _, e := range events {
		obj := &e.InvolvedObject
		if (obj.Kind == gsakey.GetKind()) && (obj.Namespace == gsakey.GetNamespace()) && (obj.Name == gsakey.GetName()) {
			if e.Reason == "Created" && e.Message == fmt.Sprintf("secret %v in namespace %v Successfully created", gsakey.GetName(), gsakey.GetNamespace()) {
				secretCreated = true
				break
			}
		}
	}
	if !secretCreated {
		t.Fatalf("no event found to show the secret %v is created", gsakey.GetName())
	}
	// delete the service account key object
	testk8s.RemoveDeletionDefenderFinalizerForUnstructured(t, gsakey, kubeClient)
	if err := kubeClient.Delete(ctx, gsakey); err != nil {
		t.Fatalf("error deleting iamserviceaccountkey %v: %v", gsakey.GetName(), err)
	}
	if _, err := sakReconciler.Reconcile(context.TODO(), sakRequest); err != nil {
		t.Fatalf("error reconciling iamserviceaccountkey %v: %v", sakName, err)
	}
	verifyGSAKeyRemoved(t, iamClient, keyName)

	// clean up the dependent service account
	accountName := fmt.Sprintf("projects/%v/serviceAccounts/%v@%v.iam.gserviceaccount.com", project, sa.GetName(), project)
	iamClient.Projects.ServiceAccounts.Delete(accountName).Do()
}

func verifyGSAKeyRemoved(t *testing.T, iamClient *iam.Service, keyName string) {
	// iam is eventually consistent so poll until the key is not found
	err := wait.PollImmediate(10*time.Second, 2*time.Minute, func() (done bool, err error) {
		if _, err := iamClient.Projects.ServiceAccounts.Keys.Get(keyName).Do(); err == nil || !gcp.IsNotFoundError(err) {
			return false, nil
		}
		return true, nil
	})
	if err != nil {
		t.Fatalf("the underlying gsa key %v doesn't get removed: %v", keyName, err)
	}
}

func newTestReconciler(t *testing.T, mgr manager.Manager, crdPath string, provider *tfschema.Provider) reconcile.Reconciler {
	crdPath, err := filepath.Abs(crdPath)
	if err != nil {
		t.Fatalf("error getting path to CRD: %v", err)
	}
	crd := dynamic.UnmarshalFileToCRD(t, crdPath)
	smLoader := testservicemappingloader.New(t)
	// Set 'immediateReconcileRequests' and 'resourceWatcherRoutines'
	// to nil to disable reconciler's ability to create asynchronous
	// watches on unready dependencies. This feature of the reconciler
	// is unnecessary for our integration tests since we reconcile
	// each dependency first before the resource under test is
	// reconciled. Overall, the feature adds risk of complications
	// due to it's multi-threaded nature.
	var immediateReconcileRequests chan event.GenericEvent = nil
	var resourceWatcherRoutines *semaphore.Weighted = nil

	stateIntoSpecDefaulter := stateintospec.NewStateIntoSpecDefaulter(mgr.GetClient())
	reconciler, err := tf.NewReconciler(mgr, crd, provider, smLoader, immediateReconcileRequests, resourceWatcherRoutines, []k8s.Defaulter{stateIntoSpecDefaulter}, &testjitter.TestJitterGenerator{})
	if err != nil {
		t.Fatalf("error creating reconciler: %v", err)
	}
	return reconciler
}

func newSecretGenerator(t *testing.T, mgr manager.Manager, crdPath string) reconcile.Reconciler {
	crdPath, err := filepath.Abs(crdPath)
	if err != nil {
		t.Fatalf("error getting path to CRD: %v", err)
	}
	crd := dynamic.UnmarshalFileToCRD(t, crdPath)

	reconciler := newReconciler(mgr, crd, &jitter.SimpleJitterGenerator{})
	return reconciler
}

func convertToUnstructAndReplaceName(t *testing.T, testID, testNamespace, sa string, fileName string) *unstructured.Unstructured {
	b, err := os.ReadFile(fileName)
	if err != nil {
		log.Fatalf("error reading file '%v': %v", fileName, err)
	}
	s := string(b)
	s = strings.Replace(s, "${uniqueId}", testID, -1)
	s = strings.Replace(s, "${projectId}", testNamespace, -1)
	s = strings.Replace(s, "${IAMServiceAccount}", sa, -1)
	b = []byte(s)

	// Convert new bytes to unstructured object
	u := &unstructured.Unstructured{}
	err = yaml.Unmarshal(b, u)
	if err != nil {
		t.Fatalf("error unmarshalling bytes to CRD: %v", err)
	}
	return u
}

func TestMain(m *testing.M) {
	testmain.ForIntegrationTests(m, &mgr)
}
