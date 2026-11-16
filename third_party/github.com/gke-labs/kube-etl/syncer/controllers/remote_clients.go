// Copyright 2026 Google LLC
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package controllers

import (
	"context"
	"sync"

	krmv1alpha1 "github.com/gke-labs/kube-etl/syncer/api/v1alpha1"
	"sigs.k8s.io/controller-runtime/pkg/client"
)

// remoteClientCache caches direct (uncached-read) clients to remote clusters,
// keyed by RemoteConfigProvider.Key, so Push mode reuses one client (with its
// RESTMapper, discovery results and HTTP connections) per remote instead of
// building a new one on every event.
type remoteClientCache struct {
	provider RemoteConfigProvider

	mu      sync.Mutex
	clients map[string]client.Client
}

func newRemoteClientCache(provider RemoteConfigProvider) *remoteClientCache {
	return &remoteClientCache{provider: provider, clients: make(map[string]client.Client)}
}

// get returns the client for the remote of a KRMSyncer, creating it on first use.
func (c *remoteClientCache) get(ctx context.Context, namespace string, remote *krmv1alpha1.RemoteConfig) (client.Client, error) {
	key, err := c.provider.Key(namespace, remote)
	if err != nil {
		return nil, err
	}

	c.mu.Lock()
	defer c.mu.Unlock()
	if cl, ok := c.clients[key]; ok {
		return cl, nil
	}

	restConfig, err := c.provider.RESTConfig(ctx, namespace, remote)
	if err != nil {
		return nil, err
	}
	cl, err := client.New(restConfig, client.Options{})
	if err != nil {
		return nil, err
	}
	c.clients[key] = cl
	return cl, nil
}
