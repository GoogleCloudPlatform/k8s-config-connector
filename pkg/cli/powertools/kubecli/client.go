// Copyright 2024 Google LLC
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

package kubecli

import (
	"context"
	"fmt"
	"strings"

	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/apis/meta/v1/unstructured"
	"k8s.io/apimachinery/pkg/runtime/schema"
	"k8s.io/apimachinery/pkg/types"
	"k8s.io/client-go/discovery"
	"k8s.io/client-go/rest"
	"sigs.k8s.io/controller-runtime/pkg/client"
	"sigs.k8s.io/controller-runtime/pkg/client/config"
)

type Client struct {
	client.Client
	DiscoveryClient discovery.DiscoveryInterface
}

func NewClient(ctx context.Context, options ClusterOptions) (*Client, error) {
	restConfig, err := config.GetConfig()
	if err != nil {
		return nil, fmt.Errorf("getting kubernetes configuration: %w", err)
	}

	if options.Impersonate != nil {
		restConfig.Impersonate = *options.Impersonate
	}

	httpClient, err := rest.HTTPClientFor(restConfig)
	if err != nil {
		return nil, fmt.Errorf("building kubernetes http client: %w", err)
	}

	kubeClient, err := client.New(restConfig, client.Options{
		HTTPClient: httpClient,
	})
	if err != nil {
		return nil, fmt.Errorf("building kubernetes client: %w", err)
	}

	discoveryClient, err := discovery.NewDiscoveryClientForConfigAndClient(restConfig, httpClient)
	if err != nil {
		return nil, fmt.Errorf("building discovery client: %w", err)
	}

	return &Client{
		DiscoveryClient: discoveryClient,
		Client:          kubeClient,
	}, nil
}

func (c *Client) GetObject(ctx context.Context, options ObjectOptions) (*unstructured.Unstructured, error) {
	if options.Kind == "" {
		return nil, fmt.Errorf("must specify object kind to target")
	}

	if options.Name == "" {
		return nil, fmt.Errorf("must specify object name to target")
	}

	if options.Namespace == "" {
		return nil, fmt.Errorf("must specify object namespace to target")
	}

	resources, err := c.DiscoveryClient.ServerPreferredResources()
	if err != nil {
		return nil, fmt.Errorf("discovering server resources: %w", err)
	}

	var matches []metav1.APIResource
	for _, group := range resources {
		for _, resource := range group.APIResources {
			match := false
			if strings.EqualFold(resource.Kind, options.Kind) {
				match = true
			}
			if strings.EqualFold(resource.Name, options.Kind) {
				match = true
			}
			if strings.EqualFold(resource.SingularName, options.Kind) {
				match = true
			}
			for _, shortName := range resource.ShortNames {
				if strings.EqualFold(shortName, options.Kind) {
					match = true
				}
			}
			if match {
				gv, err := schema.ParseGroupVersion(group.GroupVersion)
				if err != nil {
					return nil, fmt.Errorf("parsing group version %q: %w", group.GroupVersion, err)
				}

				// populate the group and version
				r := resource
				r.Group = gv.Group
				r.Version = gv.Version

				matches = append(matches, r)
			}
		}
	}
	if len(matches) == 0 {
		return nil, fmt.Errorf("did not find any kubernetes kinds for %q", options.Kind)
	}
	if len(matches) > 1 {
		// TODO: Print fully-qualified names
		return nil, fmt.Errorf("found multiple kubernetes kind for %q", options.Kind)
	}
	resource := matches[0]

	gvk := schema.GroupVersionKind{Group: resource.Group, Version: resource.Version, Kind: resource.Kind}

	key := types.NamespacedName{
		Name:      options.Name,
		Namespace: options.Namespace,
	}

	u := &unstructured.Unstructured{}
	u.SetGroupVersionKind(gvk)

	if err := c.Client.Get(ctx, key, u); err != nil {
		return nil, fmt.Errorf("getting object %v: %w", key, err)
	}
	return u, nil
}
