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

package testmain_test

import (
	"testing"

	testmain "github.com/GoogleCloudPlatform/k8s-config-connector/pkg/test/main"
	"k8s.io/apimachinery/pkg/version"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/rest"
	"sigs.k8s.io/controller-runtime/pkg/manager"
)

var (
	mgr manager.Manager
)

// This test ensures that all test cases are running against the desired version
// of kube-apiserver. The desired version is specified in scripts/shared-vars.sh.
// The `expectedVersion` variable in the test will need to be updated when we upgrade the kube-apiserver binary.
func TestKubeAPIServerVersion(t *testing.T) {
	expectedVersion := "v1.31.0"
	v, err := getKubernetesVersion(mgr.GetConfig())
	if err != nil {
		t.Fatalf("error retrieving the Kubernetes API Server version")
	}
	if v.String() != expectedVersion {
		t.Fatalf("got the Kubernetes API Server version %v, expect to have %v", v, expectedVersion)
	}
}

// getKubernetesVersion retrieves the Kubernetes API Server version,
// and returns the major and minor version.
func getKubernetesVersion(clientConfig *rest.Config) (*version.Info, error) {
	clientGoClient := kubernetes.NewForConfigOrDie(clientConfig)
	version, err := clientGoClient.ServerVersion()
	return version, err
}

func TestMain(m *testing.M) {
	testmain.ForUnitTests(m, &mgr)
}
