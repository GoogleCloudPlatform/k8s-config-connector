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

package main

import (
	"context"
	"flag"
	"fmt"
	"log"
	"time"

	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/apis/meta/v1/unstructured"
	"k8s.io/apimachinery/pkg/runtime/schema"
	"k8s.io/apimachinery/pkg/util/wait"
	"k8s.io/client-go/dynamic"
	"k8s.io/client-go/util/retry"
	ctrl "sigs.k8s.io/controller-runtime"
	"sigs.k8s.io/yaml"
)

const defaultConfigConnector = `
apiVersion: core.cnrm.cloud.google.com/v1beta1
kind: ConfigConnector
metadata:
  name: configconnector.core.cnrm.cloud.google.com
spec:
  mode: namespaced
`

const defaultConfigConnectorNamespaceSeparated = `
apiVersion: core.cnrm.cloud.google.com/v1beta1
kind: ConfigConnector
metadata:
  name: configconnector.core.cnrm.cloud.google.com
  labels:
    tenancy.gke.io/access-level: supervisor
    tenancy.gke.io/project: no-project
    tenancy.gke.io/tenant: no-tenant
    cnrm.cloud.google.com/manager-namespace-suffix: %s  
spec:
  mode: namespaced
`

var configConnectorResource = schema.GroupVersionResource{
	Group:    "core.cnrm.cloud.google.com",
	Version:  "v1beta1",
	Resource: "configconnectors",
}

// This program exists to do the work that needs to be done for the GKE add-on
// after the manager container is created. This is meant to be used for the GKE
// add-on only, not the standalone operator.
func main() {
	var managerNamespaceSuffix string
	flag.StringVar(&managerNamespaceSuffix, "manager-namespace-suffix", "", "Create controller manager pod/SA in a separate namespace, replacing suffix of watched namespace with the specified suffix.")
	flag.Parse()

	ctx := context.Background()
	dynamicClient, err := dynamic.NewForConfig(ctrl.GetConfigOrDie())
	if err != nil {
		log.Fatalf("error creating dynamic client: %v", err)
	}
	if err := createDefaultConfigConnector(ctx, dynamicClient, managerNamespaceSuffix); err != nil {
		log.Fatalf("error creating default ConfigConnector object: %v", err)
	}
}

// createDefaultConfigConnector creates a ConfigConnector object on the K8s API
// server. This is done for users who want to have a default ConfigConnector
// object created for them upon enabling the GKE add-on.
func createDefaultConfigConnector(ctx context.Context, dynamicClient dynamic.Interface, managerNamespaceSuffix string) error {
	u := &unstructured.Unstructured{}
	var b []byte
	if managerNamespaceSuffix == "" {
		b = []byte(defaultConfigConnector)
	} else {
		b = []byte(fmt.Sprintf(defaultConfigConnectorNamespaceSeparated, managerNamespaceSuffix))
	}
	if err := yaml.Unmarshal(b, u); err != nil {
		return fmt.Errorf("error unmarshalling bytes to unstruct: %w", err)
	}

	// Create the ConfigConnector object. Retry on error just in case the
	// ConfigConnector CRD does not exist yet.
	backoff := wait.Backoff{
		// Make 10 attempts with 1s delay between each attempt.
		Steps:    10,
		Duration: 1 * time.Second,
		Factor:   1.0,
	}
	retriable := func(_ error) bool {
		// Retry on all errors.
		return true
	}
	return retry.OnError(backoff, retriable, func() error {
		// Terminate in case a ConfigConnector object already exists.
		_, err := dynamicClient.Resource(configConnectorResource).Get(ctx, u.GetName(), metav1.GetOptions{})
		if err == nil {
			return nil
		}

		_, err = dynamicClient.Resource(configConnectorResource).Create(ctx, u, metav1.CreateOptions{})
		return err
	})
}
