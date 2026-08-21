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
	"os"
	"path/filepath"
	"regexp"
	"strings"
	"time"

	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/apis/meta/v1/unstructured"
	"k8s.io/apimachinery/pkg/runtime/schema"
	"k8s.io/apimachinery/pkg/types"
	"k8s.io/client-go/discovery"
	diskcached "k8s.io/client-go/discovery/cached/disk"
	"k8s.io/client-go/rest"
	"k8s.io/client-go/tools/clientcmd"
	"k8s.io/client-go/util/homedir"
	"sigs.k8s.io/controller-runtime/pkg/client"
	"sigs.k8s.io/controller-runtime/pkg/client/config"
)

type Client struct {
	client.Client
	DiscoveryClient discovery.DiscoveryInterface
}

func NewClient(ctx context.Context, options ClusterOptions) (*Client, error) {
	var restConfig *rest.Config
	if options.Kubeconfig != "" {
		rc, err := clientcmd.BuildConfigFromFlags("", options.Kubeconfig)
		if err != nil {
			return nil, fmt.Errorf("loading kubernetes configuration from %q: %w", options.Kubeconfig, err)
		}
		restConfig = rc
	} else {
		rc, err := config.GetConfig()
		if err != nil {
			return nil, fmt.Errorf("getting kubernetes configuration: %w", err)
		}
		restConfig = rc
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

	discoveryClient, err := buildDiscoveryClient(ctx, restConfig)
	// discoveryClient, err := discovery.NewDiscoveryClientForConfigAndClient(restConfig, httpClient)
	if err != nil {
		return nil, fmt.Errorf("building discovery client: %w", err)
	}

	return &Client{
		DiscoveryClient: discoveryClient,
		Client:          kubeClient,
	}, nil
}

func buildDiscoveryClient(ctx context.Context, restConfig *rest.Config) (discovery.DiscoveryInterface, error) {
	// Based on toDiscoveryClient in https://github.com/kubernetes/kubernetes/blob/v1.30.0-alpha.0/staging/src/k8s.io/cli-runtime/pkg/genericclioptions/config_flags.go

	config := *restConfig

	// config.Burst = f.discoveryBurst
	// config.QPS = f.discoveryQPS

	cacheDir := getDefaultCacheDir()

	// // retrieve a user-provided value for the "cache-dir"
	// // override httpCacheDir and discoveryCacheDir if user-value is given.
	// // user-provided value has higher precedence than default
	// // and KUBECACHEDIR environment variable.
	// if f.CacheDir != nil && *f.CacheDir != "" && *f.CacheDir != getDefaultCacheDir() {
	// 	cacheDir = *f.CacheDir
	// }

	httpCacheDir := filepath.Join(cacheDir, "http")
	discoveryCacheDir := computeDiscoverCacheDir(filepath.Join(cacheDir, "discovery"), config.Host)

	return diskcached.NewCachedDiscoveryClientForConfig(&config, discoveryCacheDir, httpCacheDir, time.Duration(6*time.Hour))
}

// overlyCautiousIllegalFileCharacters matches characters that *might* not be supported.  Windows is really restrictive, so this is really restrictive
var overlyCautiousIllegalFileCharacters = regexp.MustCompile(`[^(\w/.)]`)

// computeDiscoverCacheDir takes the parentDir and the host and comes up with a "usually non-colliding" name.
func computeDiscoverCacheDir(parentDir, host string) string {
	// strip the optional scheme from host if its there:
	schemelessHost := strings.Replace(strings.Replace(host, "https://", "", 1), "http://", "", 1)
	// now do a simple collapse of non-AZ09 characters.  Collisions are possible but unlikely.  Even if we do collide the problem is short lived
	safeHost := overlyCautiousIllegalFileCharacters.ReplaceAllString(schemelessHost, "_")
	return filepath.Join(parentDir, safeHost)
}

// getDefaultCacheDir returns default caching directory path.
// it first looks at KUBECACHEDIR env var if it is set, otherwise
// it returns standard kube cache dir.
func getDefaultCacheDir() string {
	if kcd := os.Getenv("KUBECACHEDIR"); kcd != "" {
		return kcd
	}

	return filepath.Join(homedir.HomeDir(), ".kube", "cache")
}

func (c *Client) GetObject(ctx context.Context, options ObjectOptions) (*unstructured.Unstructured, error) {
	if options.Kind == "" {
		return nil, fmt.Errorf("must specify object kind to target (use --kind flag)")
	}

	if options.Name == "" {
		return nil, fmt.Errorf("must specify object name to target (use --name flag)")
	}

	if options.Namespace == "" {
		return nil, fmt.Errorf("must specify object namespace to target (use --namespace flag)")
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
