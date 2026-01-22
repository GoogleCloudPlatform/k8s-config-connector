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

package kccmanager_test

import (
	"context"
	"fmt"
	"math/rand"
	"testing"
	"time"

	operatorv1beta1 "github.com/GoogleCloudPlatform/k8s-config-connector/operator/pkg/apis/core/v1beta1"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/kccmanager"
	corev1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/apis/meta/v1/unstructured"
	"k8s.io/apimachinery/pkg/types"
	"k8s.io/apimachinery/pkg/util/wait"
	"sigs.k8s.io/controller-runtime/pkg/cache"
	"sigs.k8s.io/controller-runtime/pkg/client"
	"sigs.k8s.io/controller-runtime/pkg/config"
	"sigs.k8s.io/controller-runtime/pkg/manager"
	"sigs.k8s.io/controller-runtime/pkg/metrics/server"
)

func TestResourceExclusion(t *testing.T) {
	ctx := context.Background()
	// Use the config from the existing manager started in TestMain
	restConfig := clusterModeManager.GetConfig()
	c := clusterModeManager.GetClient()
	skipNameValidation := true

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
			c.Delete(ctx, cc)
		}()

		if err := c.Create(ctx, cc); err != nil {
			t.Fatalf("error creating ConfigConnector: %v", err)
		}

		// Start a new KCC manager in cluster mode
		cfg := kccmanager.Config{
			ManagerOptions: manager.Options{
				Controller: config.Controller{
					SkipNameValidation: &skipNameValidation,
				},
				Metrics: server.Options{
					BindAddress: "0",
				},
				HealthProbeBindAddress: "0",
			},
			ScopedNamespace: "", // Cluster mode
		}

		kccMgr, err := kccmanager.New(ctx, restConfig, cfg)
		if err != nil {
			t.Fatalf("error creating kcc manager: %v", err)
		}

		mgrCtx, cancel := context.WithCancel(ctx)
		defer cancel()
		go func() {
			kccMgr.Start(mgrCtx)
		}()

		// Create a PubSubTopic
		topicName := "test-topic-exclusion-" + randomString()
		topic := &unstructured.Unstructured{
			Object: map[string]interface{}{
				"apiVersion": "pubsub.cnrm.cloud.google.com/v1beta1",
				"kind":       "PubSubTopic",
				"metadata": map[string]interface{}{
					"name":      topicName,
					"namespace": "default",
				},
			},
		}

		if err := c.Create(ctx, topic); err != nil {
			t.Fatalf("error creating PubSubTopic: %v", err)
		}
		defer c.Delete(context.Background(), topic)

		// Assert status is NOT updated (controller not managing it)
		// We wait a bit to be sure it doesn't update.
		// If it updates, it means the controller picked it up.
		err = wait.PollImmediate(1*time.Second, 10*time.Second, func() (bool, error) {
			u := &unstructured.Unstructured{}
			u.SetGroupVersionKind(topic.GroupVersionKind())
			if err := c.Get(ctx, types.NamespacedName{Name: topicName, Namespace: "default"}, u); err != nil {
				return false, err
			}
			// Check if status.conditions contains Ready
			status, ok := u.Object["status"].(map[string]interface{})
			if !ok {
				return false, nil // No status, good
			}
			conditions, ok := status["conditions"].([]interface{})
			if !ok {
				return false, nil // No conditions, good
			}
			if len(conditions) > 0 {
				return true, nil // Found conditions, fail
			}
			return false, nil
		})

		if err == nil {
			// PollImmediate returns nil if condition returns true (meaning we FOUND conditions)
			t.Fatalf("PubSubTopic was reconciled but expected it to be excluded")
		}
		// If err != nil (timeout), it means we never found conditions -> Success
	})

	// 2. Namespace Mode Exclusion
	t.Run("NamespaceModeExclusion", func(t *testing.T) {
		ns := "test-ns-" + randomString()
		ensureNamespace(ctx, t, c, ns)

		// Create ConfigConnectorContext in ns disabling PubSubTOPIC
		ccc := &operatorv1beta1.ConfigConnectorContext{
			ObjectMeta: metav1.ObjectMeta{
				Name:      operatorv1beta1.ConfigConnectorContextAllowedName,
				Namespace: ns,
			},
			Spec: operatorv1beta1.ConfigConnectorContextSpec{
				Experiments: &operatorv1beta1.Experiments{
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
		if err := c.Create(ctx, ccc); err != nil {
			t.Fatalf("error creating ConfigConnectorContext: %v", err)
		}

		// Start Manager in Namespace Mode
		cfg := kccmanager.Config{
			ManagerOptions: manager.Options{
				Controller: config.Controller{
					SkipNameValidation: &skipNameValidation,
				},
				Metrics: server.Options{
					BindAddress: "0",
				},
				HealthProbeBindAddress: "0",
				Cache: cache.Options{
					DefaultNamespaces: map[string]cache.Config{
						ns: {},
					},
				},
			},
			ScopedNamespace: ns,
		}
		kccMgr, err := kccmanager.New(ctx, restConfig, cfg)
		if err != nil {
			t.Fatalf("error creating kcc manager: %v", err)
		}

		mgrCtx, cancel := context.WithCancel(ctx)
		defer cancel()
		go func() {
			kccMgr.Start(mgrCtx)
		}()

		// Create PubSubTopic in ns
		topicName := "test-topic-ns-exclusion-" + randomString()
		topic := &unstructured.Unstructured{
			Object: map[string]interface{}{
				"apiVersion": "pubsub.cnrm.cloud.google.com/v1beta1",
				"kind":       "PubSubTopic",
				"metadata": map[string]interface{}{
					"name":      topicName,
					"namespace": ns,
				},
			},
		}
		if err := c.Create(ctx, topic); err != nil {
			t.Fatalf("error creating PubSubTopic: %v", err)
		}
		defer c.Delete(context.Background(), topic)

		// Assert NO reconciliation
		err = wait.PollImmediate(1*time.Second, 10*time.Second, func() (bool, error) {
			u := &unstructured.Unstructured{}
			u.SetGroupVersionKind(topic.GroupVersionKind())
			if err := c.Get(ctx, types.NamespacedName{Name: topicName, Namespace: ns}, u); err != nil {
				return false, err
			}
			status, ok := u.Object["status"].(map[string]interface{})
			if !ok {
				return false, nil
			}
			conditions, ok := status["conditions"].([]interface{})
			if !ok {
				return false, nil
			}
			if len(conditions) > 0 {
				return true, nil
			}
			return false, nil
		})
		if err == nil {
			t.Fatalf("PubSubTopic in excluded namespace was reconciled")
		}
	})
}

func ensureNamespace(ctx context.Context, t *testing.T, c client.Client, name string) {
	t.Helper()
	ns := &corev1.Namespace{
		ObjectMeta: metav1.ObjectMeta{
			Name: name,
		},
	}
	if err := c.Create(ctx, ns); err != nil {
		t.Logf("error creating namespace %v: %v", name, err)
	}
}

func randomString() string {
	rand.Seed(time.Now().UnixNano())
	return fmt.Sprintf("%d", rand.Intn(100000))
}
