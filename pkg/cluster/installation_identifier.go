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

package cluster

import (
	"context"
	"fmt"
	"time"

	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/randomid"
	corev1 "k8s.io/api/core/v1"
	"k8s.io/apimachinery/pkg/api/errors"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/types"
	"k8s.io/apimachinery/pkg/util/wait"
	"sigs.k8s.io/controller-runtime/pkg/client"
	"sigs.k8s.io/controller-runtime/pkg/log"
)

// Get an identifier for the given namespace name.
//
// The first time GetNamespaceID(...) is called, an ID is generated and persisted to the APIServer in a ConfigMap.
// The xid library is used to generate an ID that is unique across all KCC installations and namespaces.
func GetNamespaceID(namespaceIDConfigMapNN types.NamespacedName, kubeClient client.Client, ctx context.Context, namespace string) (string, error) {
	return getOrSetNamespaceId(ctx, namespaceIDConfigMapNN, kubeClient, namespace, nil)
}

// Set the namespace ID value. This is useful for scenarios where ID uniqueness is not desired, for example, while
// running integration tests, we use a single project in which to run all the tests. Each of the tests runs in a
// different namespace. If each namespace has a unique id then resource contention prevention will only allow a single
// test to succeed at a time. To enable parallel testing, we have all tests running against the main test project use
// the same ID for their namespace.
func SetNamespaceID(namespaceIDConfigMapNN types.NamespacedName, kubeClient client.Client, ctx context.Context, namespace, uniqueID string) error {
	_, err := getOrSetNamespaceId(ctx, namespaceIDConfigMapNN, kubeClient, namespace, &uniqueID)
	return err
}

// Delete the namespace and its ID from configMap. This prevents us from hitting config map size limit. (The data stored in a
// ConfigMap cannot exceed 1 MiB.)
func DeleteNamespaceID(ctx context.Context, namespaceIDConfigMapNN types.NamespacedName, kubeClient client.Client, namespace string) error {
	log := log.FromContext(ctx)
	backoff := wait.Backoff{Steps: 5, Duration: 500 * time.Millisecond, Factor: 1.5}
	deleteNamespaceIDFunc := func() (bool, error) {
		configMap, err := createOrGetNamespaceIDConfigMap(ctx, namespaceIDConfigMapNN, kubeClient)
		if err != nil {
			return false, err
		}
		if configMap.Data == nil {
			return true, nil
		}
		delete(configMap.Data, namespace)
		err = kubeClient.Update(ctx, configMap)
		if err == nil {
			return true, nil
		}
		if errors.IsConflict(err) {
			log.Info("conflict while updating configmap (may retry)", "error", err)
			return false, nil
		}
		return false, err
	}
	if err := wait.ExponentialBackoff(backoff, deleteNamespaceIDFunc); err != nil {
		return err
	}
	return nil
}

func getOrSetNamespaceId(ctx context.Context, namespaceIDConfigMapNN types.NamespacedName, kubeClient client.Client, namespace string, idToSet *string) (string, error) {
	log := log.FromContext(ctx)
	backoff := wait.Backoff{Steps: 5, Duration: 500 * time.Millisecond, Factor: 1.5}

	var id string
	getOrUpdateConfigMapFunc := func() (bool, error) {
		configMap, err := createOrGetNamespaceIDConfigMap(ctx, namespaceIDConfigMapNN, kubeClient)
		if err != nil {
			return false, err
		}
		if s, ok := configMap.Data[namespace]; ok && idToSet == nil {
			id = s
			return true, nil
		}
		if configMap.Data == nil {
			configMap.Data = make(map[string]string)
		}
		if idToSet == nil {
			id = generateID()
		} else {
			id = *idToSet
		}
		configMap.Data[namespace] = id
		err = kubeClient.Update(ctx, configMap)
		if err == nil {
			return true, nil
		}
		if errors.IsConflict(err) {
			log.Info("conflict while updating configmap (may retry)", "error", err)
			return false, nil
		}
		return false, fmt.Errorf("error updating config map '%v': %w", namespaceIDConfigMapNN, err)
	}
	if err := wait.ExponentialBackoff(backoff, getOrUpdateConfigMapFunc); err != nil {
		return "", err
	}
	return id, nil
}

func createOrGetNamespaceIDConfigMap(ctx context.Context, namespaceIDConfigMapNN types.NamespacedName, kubeClient client.Client) (*corev1.ConfigMap, error) {
	configMap := newConfigMap(namespaceIDConfigMapNN)
	if err := kubeClient.Create(ctx, &configMap); err == nil {
		return &configMap, nil
	} else if !errors.IsAlreadyExists(err) {
		return nil, fmt.Errorf("error creating configmap '%v': %v", namespaceIDConfigMapNN, err)
	}
	if err := kubeClient.Get(ctx, namespaceIDConfigMapNN, &configMap); err != nil {
		return nil, fmt.Errorf("error getting configmap '%v': %v", namespaceIDConfigMapNN, err)
	}
	return &configMap, nil
}

func generateID() string {
	return randomid.New().String()
}

func newConfigMap(namespaceIDConfigMapNN types.NamespacedName) corev1.ConfigMap {
	return corev1.ConfigMap{
		ObjectMeta: metav1.ObjectMeta{
			Name:      namespaceIDConfigMapNN.Name,
			Namespace: namespaceIDConfigMapNN.Namespace,
		},
	}
}
