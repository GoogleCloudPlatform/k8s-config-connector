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
	"reflect"
	"strconv"
	"strings"
	"testing"
	"time"

	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/k8s"
	testgcp "github.com/GoogleCloudPlatform/k8s-config-connector/pkg/test/gcp"

	"errors"

	v1 "k8s.io/api/core/v1"
	apierrors "k8s.io/apimachinery/pkg/api/errors"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/apis/meta/v1/unstructured"
	"k8s.io/apimachinery/pkg/util/wait"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/rest"
	"sigs.k8s.io/controller-runtime/pkg/client"
)

func DeleteAllEventsForUnstruct(t *testing.T, c client.Client, unstruct *unstructured.Unstructured) {
	for _, e := range getEventsForObject(t, c, unstruct.GetKind(), unstruct.GetName(), unstruct.GetNamespace()) {
		if err := c.Delete(context.TODO(), &e); err != nil {
			t.Fatalf("unable to delete event for %v %v/%v: %v", unstruct.GetKind(), unstruct.GetNamespace(), unstruct.GetName(), err)
		}
	}
}

func AssertEventRecordedForObjectMetaAndKind(t *testing.T, c client.Client, kind string, om *metav1.ObjectMeta, reason string) {
	assertEventRecorded(t, c, kind, om.Name, om.Namespace, reason)
}

func AssertEventRecordedforUnstruct(t *testing.T, c client.Client, unstruct *unstructured.Unstructured, reason string) {
	assertEventRecorded(t, c, unstruct.GetKind(), unstruct.GetName(), unstruct.GetNamespace(), reason)
}

func AssertEventNotRecordedforUnstruct(t *testing.T, c client.Client, unstruct *unstructured.Unstructured, reason string) {
	assertEventNotRecorded(t, c, unstruct.GetKind(), unstruct.GetName(), unstruct.GetNamespace(), reason)
}

func AssertObservedGenerationEquals(t *testing.T, unstruct *unstructured.Unstructured, preReconcileGeneration int64) {
	observedGeneration, found, err := unstructured.NestedInt64(unstruct.Object, "status", "observedGeneration")
	if err != nil {
		t.Errorf("error getting the value for 'status.observedGeneration': %v", err)
	}
	if !found {
		t.Errorf("'status.observedGeneration' is not found")
	}
	if observedGeneration != preReconcileGeneration {
		t.Errorf("observedGeneration %v doesn't match with the pre-reconcile generation %v", observedGeneration, preReconcileGeneration)
	}
}

func assertEventRecorded(t *testing.T, c client.Client, kind, name, namespace, reason string) {
	err := waitUntilEventRecorded(t, c, kind, name, namespace, reason)
	if err != nil {
		t.Errorf("event with reason '%v' not recorded for %v %v/%v", reason, kind, namespace, name)
	}
}

func assertEventNotRecorded(t *testing.T, c client.Client, kind, name, namespace, reason string) {
	err := waitUntilEventRecorded(t, c, kind, name, namespace, reason)
	if err == nil {
		t.Errorf("expected event with reason '%v' to not be recorded for %v %v/%v, but it was", reason, kind, namespace, name)
	} else if !errors.Is(err, wait.ErrWaitTimeout) {
		t.Errorf("error waiting for event with reason '%v' to be recorded for %v %v/%v: %v", reason, kind, namespace, name, err)
	}
}

func waitUntilEventRecorded(t *testing.T, c client.Client, kind, name, namespace, reason string) error {
	// Event firing is asynchronous, so we need to poll for whether it occurs
	interval := 10 * time.Second
	timeout := 1 * time.Minute
	return wait.PollImmediate(interval, timeout, func() (done bool, err error) {
		return eventRecorded(t, c, kind, name, namespace, reason), nil
	})
}

func eventRecorded(t *testing.T, c client.Client, kind, name, namespace, reason string) bool {
	for _, e := range getEventsForObject(t, c, kind, name, namespace) {
		if e.Reason == reason {
			return true
		}
	}
	return false
}

func getEventsForObject(t *testing.T, c client.Client, kind, name, namespace string) []v1.Event {
	listOptions := client.ListOptions{
		Namespace: namespace,
	}
	events := make([]v1.Event, 0)
	for ok := true; ok; ok = listOptions.Continue != "" {
		var eventList v1.EventList
		if err := c.List(context.TODO(), &eventList, &listOptions); err != nil {
			t.Fatalf("error listing events for %v %v/%v: %v", kind, namespace, name, err)
		}
		for _, e := range eventList.Items {
			obj := &e.InvolvedObject
			if (obj.Kind == kind) && (obj.Namespace == namespace) && (obj.Name == name) {
				events = append(events, e)
			}
		}
		listOptions.Continue = eventList.Continue
	}
	return events
}

func WaitForUnstructDeleteToFinish(t *testing.T, kubeClient client.Client, origUnstruct *unstructured.Unstructured) {
	unstruct := origUnstruct.DeepCopy()
	err := wait.PollImmediate(1*time.Second, 30*time.Second, func() (done bool, err error) {
		err = kubeClient.Get(context.TODO(), k8s.GetNamespacedName(unstruct), unstruct)
		if err == nil {
			return false, nil
		}
		if apierrors.IsNotFound(err) {
			return true, nil
		}
		return true, err
	})
	if err != nil {
		t.Fatalf("error waiting for %v %v/%v to be deleted: %v", unstruct.GetKind(), unstruct.GetNamespace(), unstruct.GetName(), err)
	}
}

// ReplaceTestVars replaces all occurrences of placeholder strings e.g. ${uniqueId} in a given byte slice.
func ReplaceTestVars(t *testing.T, b []byte, uniqueID string, project testgcp.GCPProject) []byte {
	s := string(b)
	s = strings.Replace(s, "${uniqueId}", uniqueID, -1)
	s = strings.Replace(s, "${projectId}", project.ProjectID, -1)
	if strings.Contains(s, "${projectNumber}") {
		projectNumber := strconv.FormatInt(project.ProjectNumber, 10)
		s = strings.Replace(s, "${projectNumber}", projectNumber, -1)
	}
	// Handle placeholder strings for folder id and org id specially because they are pure numbers while yaml marshalling expects strings.
	s = strings.Replace(s, fmt.Sprintf("folders/${%s}", testgcp.TestFolderID.Key), fmt.Sprintf("folders/%s", testgcp.TestFolderID.Get()), -1)
	s = strings.Replace(s, fmt.Sprintf("${%s}", testgcp.TestFolderID.Key), fmt.Sprintf("\"%s\"", testgcp.TestFolderID.Get()), -1)
	s = strings.Replace(s, fmt.Sprintf("folders/${%s}", testgcp.TestFolder2ID.Key), fmt.Sprintf("folders/%s", testgcp.TestFolder2ID.Get()), -1)
	s = strings.Replace(s, fmt.Sprintf("${%s}", testgcp.TestFolder2ID.Key), fmt.Sprintf("\"%s\"", testgcp.TestFolder2ID.Get()), -1)
	s = strings.Replace(s, fmt.Sprintf("organizations/${%s}", testgcp.TestOrgID.Key), fmt.Sprintf("organizations/%s", testgcp.TestOrgID.Get()), -1)
	s = strings.Replace(s, fmt.Sprintf("${%s}", testgcp.TestOrgID.Key), fmt.Sprintf("\"%s\"", testgcp.TestOrgID.Get()), -1)
	s = strings.Replace(s, fmt.Sprintf("${%s}", testgcp.TestDependentOrgProjectID.Key), fmt.Sprintf("\"%s\"", testgcp.TestDependentOrgProjectID.Get()), -1)
	s = strings.Replace(s, fmt.Sprintf("projects/${%s}", testgcp.TestDependentOrgProjectID.Key), fmt.Sprintf("projects/%s", testgcp.TestDependentOrgProjectID.Get()), -1)
	s = strings.Replace(s, fmt.Sprintf("projects/${%s}", testgcp.TestDependentFolderProjectID.Key), fmt.Sprintf("projects/%s", testgcp.TestDependentFolderProjectID.Get()), -1)
	s = strings.Replace(s, fmt.Sprintf("${%s}", testgcp.TestDependentOrgProjectIDWithoutQuotation), fmt.Sprintf("%s", testgcp.TestDependentOrgProjectID.Get()), -1)
	s = strings.Replace(s, fmt.Sprintf("${%s}", testgcp.TestDependentFolderProjectID.Key), fmt.Sprintf("\"%s\"", testgcp.TestDependentFolderProjectID.Get()), -1)
	s = strings.Replace(s, fmt.Sprintf("projects/${%s}", testgcp.TestDependentFolder2ProjectID), fmt.Sprintf("projects/%s", testgcp.GetDependentFolder2ProjectID(t)), -1)
	s = strings.Replace(s, fmt.Sprintf("projects/${%s}", testgcp.TestDependentNoNetworkProjectID.Key), fmt.Sprintf("projects/%s", testgcp.TestDependentNoNetworkProjectID.Get()), -1)
	s = strings.Replace(s, fmt.Sprintf("${%s}", testgcp.TestDependentNoNetworkProjectID.Key), fmt.Sprintf("\"%s\"", testgcp.TestDependentNoNetworkProjectID.Get()), -1)
	s = strings.Replace(s, fmt.Sprintf("organizations/${%s}", testgcp.IAMIntegrationTestsOrganizationID.Key), fmt.Sprintf("organizations/%s", testgcp.IAMIntegrationTestsOrganizationID.Get()), -1)
	s = strings.Replace(s, fmt.Sprintf("${%s}", testgcp.IAMIntegrationTestsOrganizationID.Key), fmt.Sprintf("\"%s\"", testgcp.IAMIntegrationTestsOrganizationID.Get()), -1)
	s = strings.Replace(s, fmt.Sprintf("${%s}", testgcp.IsolatedTestOrgName.Key), testgcp.IsolatedTestOrgName.Get(), -1)
	s = strings.Replace(s, fmt.Sprintf("${%s}", testgcp.TestBillingAccountID.Key), testgcp.TestBillingAccountID.Get(), -1)
	s = strings.Replace(s, fmt.Sprintf("${%s}", testgcp.TestBillingAccountIDForBillingResources.Key), testgcp.TestBillingAccountIDForBillingResources.Get(), -1)
	s = strings.Replace(s, fmt.Sprintf("${%s}", testgcp.FirestoreTestProject.Key), testgcp.FirestoreTestProject.Get(), -1)
	s = strings.Replace(s, fmt.Sprintf("${%s}", testgcp.IAMIntegrationTestsBillingAccountID.Key), testgcp.IAMIntegrationTestsBillingAccountID.Get(), -1)
	s = strings.Replace(s, fmt.Sprintf("${%s}", testgcp.CloudFunctionsTestProject), testgcp.GetCloudFunctionsTestProject(t), -1)
	s = strings.Replace(s, fmt.Sprintf("${%s}", testgcp.IdentityPlatformTestProject.Key), testgcp.IdentityPlatformTestProject.Get(), -1)
	s = strings.Replace(s, fmt.Sprintf("${%s}", testgcp.InterconnectTestProject), testgcp.GetInterconnectTestProject(t), -1)
	s = strings.Replace(s, fmt.Sprintf("${%s}", testgcp.HighCPUQuotaTestProject), testgcp.GetHighCPUQuotaTestProject(t), -1)
	s = strings.Replace(s, fmt.Sprintf("${%s}", testgcp.RecaptchaEnterpriseTestProject.Key), testgcp.RecaptchaEnterpriseTestProject.Get(), -1)
	s = strings.Replace(s, fmt.Sprintf("${%s}", testgcp.TestKCCAttachedClusterProject.Key), testgcp.TestKCCAttachedClusterProject.Get(), -1)
	s = strings.Replace(s, fmt.Sprintf("${%s}", testgcp.TestKCCAttachedClusterPlatformVersion.Key), testgcp.TestKCCAttachedClusterPlatformVersion.Get(), -1)
	s = strings.Replace(s, fmt.Sprintf("${%s}", testgcp.TestKCCVertexAIIndexBucket.Key), testgcp.TestKCCVertexAIIndexBucket.Get(), -1)
	s = strings.Replace(s, fmt.Sprintf("${%s}", testgcp.TestKCCVertexAIIndexDataURI.Key), testgcp.TestKCCVertexAIIndexDataURI.Get(), -1)
	s = strings.Replace(s, fmt.Sprintf("${%s}", testgcp.TestInterconnectID.Key), testgcp.TestInterconnectID.Get(), -1)
	return []byte(s)
}

// Collects an expected number of events from the API server. The timeout is applied on a per-event basis, so it is possible this function
// takes upwards of expectedCount * timeoutSeconds duration to 'timeout'. When a timeout does occur, t.Fatal(...) is invoked.
func CollectEvents(t *testing.T, config *rest.Config, namespace string, expectedCount int, timeout time.Duration) []v1.Event {
	t.Helper()
	clientSet, err := kubernetes.NewForConfig(config)
	if err != nil {
		t.Fatalf("error creating k8s client: %v", err)
	}
	listOptions := metav1.ListOptions{}
	watcher, err := clientSet.CoreV1().Events(namespace).Watch(context.Background(), listOptions)
	if err != nil {
		t.Fatalf("error creating event watch: %v", err)
	}
	defer watcher.Stop()
	results := make([]v1.Event, 0)
	ch := watcher.ResultChan()
	for i := 0; i < expectedCount; i++ {
		select {
		case res := <-ch:
			event, ok := res.Object.(*v1.Event)
			if !ok {
				t.Fatalf("unexpected type returned in channel: got '%v', watch '%v'", reflect.TypeOf(res), reflect.TypeOf(v1.Event{}))
			}
			results = append(results, *event)
		case <-time.After(timeout):
			t.Fatalf("expected '%v' event(s), collected '%v' event(s), timed out waiting for the last '%v' event(s)t'",
				expectedCount, len(results), expectedCount-len(results))
		}
	}
	return results
}
