// Copyright 2025 Google LLC
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

package kccmanager

import (
	"context"
	"testing"

	operatorv1beta1 "github.com/GoogleCloudPlatform/k8s-config-connector/operator/pkg/apis/core/v1beta1"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/kccmanager"
	testk8s "github.com/GoogleCloudPlatform/k8s-config-connector/pkg/test/k8s"
	testmain "github.com/GoogleCloudPlatform/k8s-config-connector/pkg/test/main"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"sigs.k8s.io/controller-runtime/pkg/manager"
)

func TestResourceExclusion(t *testing.T) {
	ctx := context.Background()
	mgr, stop := testmain.StartTestManager(t, ctx)
	defer stop()

	c := mgr.GetClient()

	// 1. Cluster Mode Exclusion
	t.Run("ClusterModeExclusion", func(t *testing.T) {
		// Disable PubSubTopic in ConfigConnector
		cc := &operatorv1beta1.ConfigConnector{
			ObjectMeta: metav1.ObjectMeta{
				Name: operatorv1beta1.ConfigConnectorAllowedName,
			},
			Spec: operatorv1beta1.ConfigConnectorSpec{
				Experiments: &operatorv1beta1.CCExperiments{
					ResourceSettings: []operatorv1beta1.ResourceSettings{
						{
							Group:   "pubsub.cnrm.cloud.google.com",
							Kind:    "PubSubTopic",
							Enabled: false,
						},
					},
				},
			},
		}
		// Ensure cleanup
		defer func() {
			testk8s.RemoveConfigConnector(ctx, t, c)
		}()

		if err := c.Create(ctx, cc); err != nil {
			t.Fatalf("error creating ConfigConnector: %v", err)
		}

		// Start a new KCC manager in cluster mode
		// We need to run it in a separate goroutine or just create it and verify controllers
		// Since we cannot easily inspect internal state of a started manager, we can check if the controller is registered?
		// Or easier: check if it reconciles.
		// Construct a separate manager (simulating the kcc-manager binary)
		cfg := kccmanager.Config{
			ManagerOptions: manager.Options{
				MetricsBindAddress: "0", // Disable metrics serving to avoid conflict
			},
			ScopedNamespace: "", // Cluster mode
		}
		// Pass the shared test env config
		kccMgr, err := kccmanager.New(ctx, mgr.GetConfig(), cfg)
		if err != nil {
			t.Fatalf("error creating kcc manager: %v", err)
		}

		// Check if controller is registered?
		// Unfortuantely AddDefaultControllers registers them directly.
		// We can try to use a mock or check side effects.
		// One side effect is: if we create a PubSubTopic, it stays in Unmanaged?
		// But in integration test, if we don't start the manager, it won't reconcile anyway.
		// If we DO start the manager, it should reconcile if enabled.

		// Let's start the manager in a go routine
		mgrCtx, cancel := context.WithCancel(ctx)
		defer cancel()
		go func() {
			if err := kccMgr.Start(mgrCtx); err != nil {
				// It might fail if ports conflict, but we disabled metrics.
				// Health probe might conflict.
				t.Logf("kcc manager stopped: %v", err)
			}
		}()

		// Verify PubSubTopic is NOT reconciled
		// Create a PubSubTopic
		// ...
		// Assert status is empty or condition NotReady?
		// If controller is missing, it will do nothing. status will be empty.

		// TODO: Implement actual resource creation and check
	})

	// 2. Namespace Mode Exclusion
	// ...
}
