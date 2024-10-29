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

package cluster_test

import (
	"context"
	"testing"

	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/cluster"
	testcontroller "github.com/GoogleCloudPlatform/k8s-config-connector/pkg/test/controller"
	testmain "github.com/GoogleCloudPlatform/k8s-config-connector/pkg/test/main"
	corev1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/types"

	"sigs.k8s.io/controller-runtime/pkg/manager"
)

const (
	namespaceName = "my-namespace"
	setIDValue    = "my-specified-id"
)

var (
	mgr manager.Manager

	namespaceIDConfigMapNN = types.NamespacedName{
		Namespace: "system-namespace-name",
		Name:      "namespace-id",
	}
)

func TestGetNamespaceId(t *testing.T) {
	if err := testcontroller.EnsureNamespaceExists(mgr.GetClient(), namespaceIDConfigMapNN.Namespace); err != nil {
		t.Fatal(err)
	}
	testGetNamespaceID(t, mgr)
	testSetNamespaceID(t, mgr)
	testDeleteNamespaceID(t, mgr)
}

func testGetNamespaceID(t *testing.T, mgr manager.Manager) {
	id := getNamespaceID(t, mgr)
	if id == setIDValue {
		t.Fatalf("expected a generated value for id, instead got '%v'", id)
	}
}

func testSetNamespaceID(t *testing.T, mgr manager.Manager) {
	err := cluster.SetNamespaceID(context.TODO(), namespaceIDConfigMapNN, mgr.GetClient(), namespaceName, setIDValue)
	if err != nil {
		t.Fatalf("unexpected error when setting namespace id: %v", err)
	}
	id := getNamespaceID(t, mgr)
	if id != setIDValue {
		t.Fatalf("unexpected id value: got '%v', want '%v'", id, setIDValue)
	}
}

func getNamespaceID(t *testing.T, mgr manager.Manager) string {
	t.Helper()
	id, err := cluster.GetNamespaceID(context.TODO(), namespaceIDConfigMapNN, mgr.GetClient(), namespaceName)
	if err != nil {
		t.Fatalf("unexpected error when getting namespace id: %v", err)
	}
	if id == "" {
		t.Fatalf("expected valid id value, instead got empty string")
	}
	return id
}

func testDeleteNamespaceID(t *testing.T, mgr manager.Manager) {
	if err := cluster.DeleteNamespaceID(context.TODO(), namespaceIDConfigMapNN, mgr.GetClient(), namespaceName); err != nil {
		t.Fatalf("unexpected error when deleting namespace id: %v", err)
	}

	var configMap = corev1.ConfigMap{
		ObjectMeta: metav1.ObjectMeta{
			Name:      namespaceIDConfigMapNN.Name,
			Namespace: namespaceIDConfigMapNN.Namespace,
		},
	}

	if err := mgr.GetClient().Get(context.TODO(), namespaceIDConfigMapNN, &configMap); err != nil {
		t.Fatalf("error getting configmap '%v': %v", namespaceIDConfigMapNN, err)
	}

	if val, ok := configMap.Data[namespaceName]; ok {
		t.Fatalf("error checking deleted namespace, unexpected value from configmap '%v': %v", namespaceName, val)
	}
}

func TestMain(m *testing.M) {
	testmain.ForUnitTests(m, &mgr)
}
