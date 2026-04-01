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
	"strings"
	"testing"
	"time"

	operatorv1beta1 "github.com/GoogleCloudPlatform/k8s-config-connector/operator/pkg/apis/core/v1beta1"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/kccmanager"
	corev1 "k8s.io/api/core/v1"
	"k8s.io/apimachinery/pkg/api/errors"
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
					ResourceSettings: &operatorv1beta1.ResourceSettings{
						Mode: operatorv1beta1.ResourceSettingsModeExclude,
						Resources: []operatorv1beta1.ResourceFilter{
							{
								Group: "pubsub.cnrm.cloud.google.com",
								Kind:  "PubSubTopic",
							},
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
			ScopedNamespace:    "", // Cluster mode
			SkipNameValidation: true,
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
					ResourceSettings: &operatorv1beta1.ResourceSettings{
						Mode: operatorv1beta1.ResourceSettingsModeExclude,
						Resources: []operatorv1beta1.ResourceFilter{
							{
								Group: "pubsub.cnrm.cloud.google.com",
								Kind:  "PubSubTopic",
							},
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
			ScopedNamespace:    ns,
			SkipNameValidation: true,
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

	// 3. Namespace Mode Precedence
	t.Run("AdditiveExclusion", func(t *testing.T) {
		ns := "test-ns-additive-" + randomString()
		ensureNamespace(ctx, t, c, ns)

		// 1. Create Global ConfigConnector disabling PubSubTopic
		cc := &operatorv1beta1.ConfigConnector{
			ObjectMeta: metav1.ObjectMeta{
				Name: operatorv1beta1.ConfigConnectorAllowedName,
			},
			Spec: operatorv1beta1.ConfigConnectorSpec{
				Mode: "namespaced",
				Experiments: &operatorv1beta1.CCExperiments{
					ResourceSettings: &operatorv1beta1.ResourceSettings{
						Mode: operatorv1beta1.ResourceSettingsModeExclude,
						Resources: []operatorv1beta1.ResourceFilter{
							{
								Group: "pubsub.cnrm.cloud.google.com",
								Kind:  "PubSubTopic",
							},
						},
					},
				},
			},
		}
		if err := c.Create(ctx, cc); err != nil {
			// If already exists (from previous test), update it
			if errors.IsAlreadyExists(err) {
				existingCC := &operatorv1beta1.ConfigConnector{}
				if err := c.Get(ctx, types.NamespacedName{Name: operatorv1beta1.ConfigConnectorAllowedName}, existingCC); err != nil {
					t.Fatalf("error getting existing ConfigConnector: %v", err)
				}
				existingCC.Spec.Experiments = cc.Spec.Experiments
				if err := c.Update(ctx, existingCC); err != nil {
					t.Fatalf("error updating ConfigConnector: %v", err)
				}
			} else {
				t.Fatalf("error creating ConfigConnector: %v", err)
			}
		}
		defer func() {
			c.Delete(ctx, cc)
		}()

		// 2. Create Local ConfigConnectorContext disabling StorageBucket
		ccc := &operatorv1beta1.ConfigConnectorContext{
			ObjectMeta: metav1.ObjectMeta{
				Name:      operatorv1beta1.ConfigConnectorContextAllowedName,
				Namespace: ns,
			},
			Spec: operatorv1beta1.ConfigConnectorContextSpec{
				Experiments: &operatorv1beta1.Experiments{
					ResourceSettings: &operatorv1beta1.ResourceSettings{
						Mode: operatorv1beta1.ResourceSettingsModeExclude,
						Resources: []operatorv1beta1.ResourceFilter{
							{
								Group: "storage.cnrm.cloud.google.com",
								Kind:  "StorageBucket",
							},
						},
					},
				},
			},
		}
		if err := c.Create(ctx, ccc); err != nil {
			t.Fatalf("error creating ConfigConnectorContext: %v", err)
		}

		// 3. Start Manager in Namespace Mode
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
			ScopedNamespace:    ns,
			SkipNameValidation: true,
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

		// 4. Create PubSubTopic in ns -> Should NOT be Reconciled (Disabled by Global)
		topicName := "test-topic-additive-" + randomString()
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
			t.Fatalf("PubSubTopic should NOT have been reconciled (Global Disable active), but obtained conditions")
		}

		// 5. Create StorageBucket in ns -> Should NOT be Reconciled (Disabled by Local)
		bucketName := "test-bucket-additive-" + randomString()
		bucket := &unstructured.Unstructured{
			Object: map[string]interface{}{
				"apiVersion": "storage.cnrm.cloud.google.com/v1beta1",
				"kind":       "StorageBucket",
				"metadata": map[string]interface{}{
					"name":      bucketName,
					"namespace": ns,
				},
			},
		}
		if err := c.Create(ctx, bucket); err != nil {
			t.Fatalf("error creating StorageBucket: %v", err)
		}
		defer c.Delete(context.Background(), bucket)

		// Assert NO reconciliation
		err = wait.PollImmediate(1*time.Second, 10*time.Second, func() (bool, error) {
			u := &unstructured.Unstructured{}
			u.SetGroupVersionKind(bucket.GroupVersionKind())
			if err := c.Get(ctx, types.NamespacedName{Name: bucketName, Namespace: ns}, u); err != nil {
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
			t.Fatalf("StorageBucket should NOT have been reconciled (Local Disable active), but obtained conditions")
		}
	})

	// 4. Namespace Mode Inclusion (Allow-list)
	t.Run("InclusiveMode", func(t *testing.T) {
		ns := "test-ns-inclusive-" + randomString()
		ensureNamespace(ctx, t, c, ns)

		// Create ConfigConnectorContext in ns with Inclusive Mode enabling only PubSubTopic
		ccc := &operatorv1beta1.ConfigConnectorContext{
			ObjectMeta: metav1.ObjectMeta{
				Name:      operatorv1beta1.ConfigConnectorContextAllowedName,
				Namespace: ns,
			},
			Spec: operatorv1beta1.ConfigConnectorContextSpec{
				Experiments: &operatorv1beta1.Experiments{
					ResourceSettings: &operatorv1beta1.ResourceSettings{
						Mode: operatorv1beta1.ResourceSettingsModeInclude,
						Resources: []operatorv1beta1.ResourceFilter{
							{
								Group: "pubsub.cnrm.cloud.google.com",
								Kind:  "PubSubTopic",
							},
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
			ScopedNamespace:    ns,
			SkipNameValidation: true,
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

		// A. Create PubSubTopic in ns -> Should be Reconciled (Enabled by Allow-list)
		topicName := "test-topic-inclusive-" + randomString()
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

		// Assert reconciliation (Check for Ready condition)
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
		if err != nil {
			t.Fatalf("PubSubTopic in inclusive mode was NOT reconciled within timeout: %v", err)
		}

		// B. Create StorageBucket in ns -> Should NOT be Reconciled (Not in Allow-list)
		bucketName := "test-bucket-inclusive-" + randomString()
		bucket := &unstructured.Unstructured{
			Object: map[string]interface{}{
				"apiVersion": "storage.cnrm.cloud.google.com/v1beta1",
				"kind":       "StorageBucket",
				"metadata": map[string]interface{}{
					"name":      bucketName,
					"namespace": ns,
				},
			},
		}
		if err := c.Create(ctx, bucket); err != nil {
			t.Fatalf("error creating StorageBucket: %v", err)
		}
		defer c.Delete(context.Background(), bucket)

		// Assert NO reconciliation
		err = wait.PollImmediate(1*time.Second, 10*time.Second, func() (bool, error) {
			u := &unstructured.Unstructured{}
			u.SetGroupVersionKind(bucket.GroupVersionKind())
			if err := c.Get(ctx, types.NamespacedName{Name: bucketName, Namespace: ns}, u); err != nil {
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
			t.Fatalf("StorageBucket should NOT have been reconciled (not in allow-list), but obtained conditions")
		}
	})

	t.Run("ConflictMode", func(t *testing.T) {
		ns := "test-ns-conflict-" + randomString()
		ensureNamespace(ctx, t, c, ns)

		// CC = Exclusive (false)
		cc := &operatorv1beta1.ConfigConnector{
			ObjectMeta: metav1.ObjectMeta{
				Name: operatorv1beta1.ConfigConnectorAllowedName,
			},
			Spec: operatorv1beta1.ConfigConnectorSpec{
				Experiments: &operatorv1beta1.CCExperiments{
					ResourceSettings: &operatorv1beta1.ResourceSettings{
						Mode: operatorv1beta1.ResourceSettingsModeExclude,
					},
				},
			},
		}
		if err := c.Create(ctx, cc); err != nil {
			t.Fatalf("error creating ConfigConnector: %v", err)
		}
		defer c.Delete(context.Background(), cc)

		// CCC = Inclusive (true)
		ccc := &operatorv1beta1.ConfigConnectorContext{
			ObjectMeta: metav1.ObjectMeta{
				Name:      operatorv1beta1.ConfigConnectorContextAllowedName,
				Namespace: ns,
			},
			Spec: operatorv1beta1.ConfigConnectorContextSpec{
				Experiments: &operatorv1beta1.Experiments{
					ResourceSettings: &operatorv1beta1.ResourceSettings{
						Mode: operatorv1beta1.ResourceSettingsModeInclude,
					},
				},
			},
		}
		if err := c.Create(ctx, ccc); err != nil {
			t.Fatalf("error creating ConfigConnectorContext: %v", err)
		}
		defer c.Delete(context.Background(), ccc)

		cfg := kccmanager.Config{
			ScopedNamespace: ns,
		}
		_, err := kccmanager.New(ctx, restConfig, cfg)
		if err == nil {
			t.Errorf("expected conflict error, but got nil")
		} else if !strings.Contains(err.Error(), "conflict") {
			t.Errorf("expected conflict error containing 'conflict', but got: %v", err)
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
