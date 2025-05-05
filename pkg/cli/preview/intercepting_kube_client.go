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

package preview

import (
	"context"
	"fmt"
	"maps"
	"net/http"
	"slices"
	"sync"
	"sync/atomic"

	"k8s.io/apimachinery/pkg/api/meta"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/apimachinery/pkg/runtime/schema"
	"k8s.io/client-go/rest"
	"k8s.io/klog/v2"
	"sigs.k8s.io/controller-runtime/pkg/cache"
	"sigs.k8s.io/controller-runtime/pkg/client"
	"sigs.k8s.io/controller-runtime/pkg/client/apiutil"
)

// interceptingKubeClient is a Kubernetes client that intercepts Kubernetes API calls.
// It forwards read-only operations "upstream" to real Kubernetes.
// It returns an error on any write operations.
type interceptingKubeClient struct {
	upstreamKubeClient KubeClient
	upstreamRestMapper meta.RESTMapper

	recorder *Recorder
}

func BuildKubeClient(upstreamRestConfig *rest.Config) (KubeClient, meta.RESTMapper, error) {
	httpClient, err := rest.HTTPClientFor(upstreamRestConfig)
	if err != nil {
		return nil, nil, fmt.Errorf("building http client: %w", err)
	}

	upstreamRestMapper, err := apiutil.NewDiscoveryRESTMapper(upstreamRestConfig, httpClient)
	if err != nil {
		return nil, nil, fmt.Errorf("creating REST mapper: %w", err)
	}

	// TODO: Replace with rest.DefaultServerUrlFor
	baseURL, _, err := DefaultServerUrlFor(upstreamRestConfig)
	if err != nil {
		return nil, nil, fmt.Errorf("getting base url: %w", err)
	}
	clientOptions := ClientOptions{
		BaseURL:    baseURL,
		HTTPClient: httpClient,
	}
	upstreamClient := NewStreamingClient(clientOptions)

	return upstreamClient, upstreamRestMapper, nil
}

// newInterceptingKubeClient creates a new interceptingKubeClient.
func newInterceptingKubeClient(recorder *Recorder, upstreamKubeClient KubeClient, upstreamRestMapper meta.RESTMapper) (*interceptingKubeClient, error) {
	return &interceptingKubeClient{
		upstreamKubeClient: upstreamKubeClient,
		upstreamRestMapper: upstreamRestMapper,
		recorder:           recorder,
	}, nil
}

// MapperProvider returns the REST mapper for the interceptingKubeClient.
func (c *interceptingKubeClient) MapperProvider(*rest.Config, *http.Client) (meta.RESTMapper, error) {
	// TODO: Verify restConfig or httpClient?
	return c.upstreamRestMapper, nil
}

// NewClient creates a controller-runtime client that intercepts Kubernetes API calls.
func (c *interceptingKubeClient) NewClient(config *rest.Config, options client.Options) (client.Client, error) {
	typeStore := &typeStore{
		restMapper: c.upstreamRestMapper,
		scheme:     options.Scheme,
	}

	client := &interceptingControllerRuntimeClient{parent: c, typeStore: typeStore}
	if options.Cache == nil || options.Cache.Reader == nil {
		return client, nil
	}
	client.cache = options.Cache.Reader
	client.cacheUnstructured = options.Cache.Unstructured
	client.uncachedGVKs = make(map[schema.GroupVersionKind]struct{})
	for _, obj := range options.Cache.DisableFor {
		gvk, err := apiutil.GVKForObject(obj, options.Scheme)
		if err != nil {
			return nil, err
		}
		client.uncachedGVKs[gvk] = struct{}{}
	}
	return client, nil
}

// NewCache creates a controller-runtime cache that intercepts Kubernetes API calls.
func (c *interceptingKubeClient) NewCache(restConfig *rest.Config, opts cache.Options) (cache.Cache, error) {
	typeStore := &typeStore{
		restMapper: c.upstreamRestMapper,
		scheme:     opts.Scheme,
	}

	return newInterceptingControllerRuntimeCache(c.upstreamKubeClient, typeStore)
}

// interceptingControllerRuntimeClient is a controller-runtime client that intercepts Kubernetes API calls.
// It forwards read-only operations "upstream" to real Kubernetes.
// It returns a BlockedKubeError on any write operations.
type interceptingControllerRuntimeClient struct {
	parent    *interceptingKubeClient
	typeStore *typeStore

	cache             client.Reader
	uncachedGVKs      map[schema.GroupVersionKind]struct{}
	cacheUnstructured bool
}

// blockedMethod is called when a write operation is attempted.
// It returns an error, so that the write operation is not forwarded upstream.
func (c *interceptingControllerRuntimeClient) blockedMethod(ctx context.Context, method string, args ...any) error {
	c.parent.recorder.RecordBlockedKubeMethod(ctx, method, args...)
	return fmt.Errorf("%q blocked in preview mode", method)
}

// ignoredMethod is called when a read operation is attempted.
// It returns nil, but does not actually forward the operation upstream.
// This is useful for status updates, where we want to record the GCP operation,
// which typically happens after the status update is made.
func (c *interceptingControllerRuntimeClient) ignoredMethod(ctx context.Context, method string, args ...any) error {
	c.parent.recorder.RecordIgnoredKubeMethod(ctx, method, args...)
	return nil
}

var _ client.Client = &interceptingControllerRuntimeClient{}

func (c *interceptingControllerRuntimeClient) shouldBypassCache(isUnstructured bool, typeInfo *typeInfo) bool {
	if c.cache == nil {
		return true
	}

	gvk := typeInfo.gvk
	// if meta.IsListType(obj) {
	// 	gvk.Kind = strings.TrimSuffix(gvk.Kind, "List")
	// }
	if _, isUncached := c.uncachedGVKs[gvk]; isUncached {
		return true
	}
	if !c.cacheUnstructured {
		return isUnstructured
	}
	return false
}

// Get retrieves an obj for the given object key from the Kubernetes Cluster.
// obj must be a struct pointer so that obj can be updated with the response
// returned by the Server.
func (c *interceptingControllerRuntimeClient) Get(ctx context.Context, key client.ObjectKey, obj client.Object, opts ...client.GetOption) error {
	_, isUnstructured := obj.(runtime.Unstructured)

	typeInfo, err := c.typeStore.getTypeInfo(obj)
	if err != nil {
		return err
	}

	shouldBypassCache := c.shouldBypassCache(isUnstructured, typeInfo)
	if !shouldBypassCache {
		return c.cache.Get(ctx, key, obj, opts...)
	}

	client := c.parent.upstreamKubeClient
	namespace := key.Namespace
	name := key.Name
	if len(opts) > 0 {
		klog.Fatalf("interceptingControllerRuntimeClient: GET %v into %T with opts %v (cache=%+v)", key, obj, opts, c.cache)
	}

	return client.Get(ctx, typeInfo, namespace, name, obj)
}

// List retrieves list of objects for a given namespace and list options. On a
// successful call, Items field in the list will be populated with the
// result returned from the server.
func (c *interceptingControllerRuntimeClient) List(ctx context.Context, list client.ObjectList, opts ...client.ListOption) error {
	panic("not implemented")
}

// Create saves the object obj in the Kubernetes cluster. obj must be a
// struct pointer so that obj can be updated with the content returned by the Server.
func (c *interceptingControllerRuntimeClient) Create(ctx context.Context, obj client.Object, opts ...client.CreateOption) error {
	return c.blockedMethod(ctx, "create", obj, opts)
}

// Delete deletes the given obj from Kubernetes cluster.
func (c *interceptingControllerRuntimeClient) Delete(ctx context.Context, obj client.Object, opts ...client.DeleteOption) error {
	return c.blockedMethod(ctx, "delete", obj, opts)
}

// Update updates the given obj in the Kubernetes cluster. obj must be a
// struct pointer so that obj can be updated with the content returned by the Server.
func (c *interceptingControllerRuntimeClient) Update(ctx context.Context, obj client.Object, opts ...client.UpdateOption) error {
	return c.blockedMethod(ctx, "update", obj, opts)
}

// Patch patches the given obj in the Kubernetes cluster. obj must be a
// struct pointer so that obj can be updated with the content returned by the Server.
func (c *interceptingControllerRuntimeClient) Patch(ctx context.Context, obj client.Object, patch client.Patch, opts ...client.PatchOption) error {
	return c.blockedMethod(ctx, "patch", obj, opts)
}

// DeleteAllOf deletes all objects of the given type matching the given options.
func (c *interceptingControllerRuntimeClient) DeleteAllOf(ctx context.Context, obj client.Object, opts ...client.DeleteAllOfOption) error {
	return c.blockedMethod(ctx, "deleteAllOf", obj, opts)
}

// Create a client which can update status subresource for kubernetes objects.
func (c *interceptingControllerRuntimeClient) Status() client.SubResourceWriter {
	return &interceptingControllerRuntimeClientSubResourceWriter{client: c, Subresource: "status"}
}

func (c *interceptingControllerRuntimeClient) SubResource(subResource string) client.SubResourceClient {
	return &interceptingControllerRuntimeClientSubResourceWriter{client: c, Subresource: subResource}
}

// Scheme returns the scheme this client is using.
func (c *interceptingControllerRuntimeClient) Scheme() *runtime.Scheme {
	panic("not implemented")
}

// RESTMapper returns the rest this client is using.
func (c *interceptingControllerRuntimeClient) RESTMapper() meta.RESTMapper {
	panic("not implemented")
}

// GroupVersionKindFor returns the GroupVersionKind for the given object.
func (c *interceptingControllerRuntimeClient) GroupVersionKindFor(obj runtime.Object) (schema.GroupVersionKind, error) {
	panic("not implemented")
}

// IsObjectNamespaced returns true if the GroupVersionKind of the object is namespaced.
func (c *interceptingControllerRuntimeClient) IsObjectNamespaced(obj runtime.Object) (bool, error) {
	panic("not implemented")
}

type interceptingControllerRuntimeClientSubResourceWriter struct {
	client *interceptingControllerRuntimeClient

	Subresource string
}

var _ client.SubResourceWriter = &interceptingControllerRuntimeClientSubResourceWriter{}

func (c *interceptingControllerRuntimeClientSubResourceWriter) Get(ctx context.Context, obj client.Object, subResource client.Object, opts ...client.SubResourceGetOption) error {
	return c.client.blockedMethod(ctx, "status.get", obj, subResource, opts)
}

func (c *interceptingControllerRuntimeClientSubResourceWriter) Create(ctx context.Context, obj client.Object, subResource client.Object, opts ...client.SubResourceCreateOption) error {
	return c.client.blockedMethod(ctx, "status.create", obj, subResource, opts)
}

func (c *interceptingControllerRuntimeClientSubResourceWriter) Update(ctx context.Context, obj client.Object, opts ...client.SubResourceUpdateOption) error {
	return c.client.ignoredMethod(ctx, "status.update", obj, opts)
}

func (c *interceptingControllerRuntimeClientSubResourceWriter) Patch(ctx context.Context, obj client.Object, patch client.Patch, opts ...client.SubResourcePatchOption) error {
	return c.client.blockedMethod(ctx, "status.patch", obj, patch, opts)
}

// interceptingControllerRuntimeCache is a controller-runtime cache that intercepts Kubernetes API calls.
// Write operations would be blocked, but generally the cache is only used for read operations.
type interceptingControllerRuntimeCache struct {
	streamingClient KubeClient
	typeStore       *typeStore

	mutex sync.Mutex

	started atomic.Bool

	informers map[schema.GroupVersionKind]*streamingInformer
}

// newInterceptingControllerRuntimeCache creates a new interceptingControllerRuntimeCache.
func newInterceptingControllerRuntimeCache(streamingClient KubeClient, typeStore *typeStore) (*interceptingControllerRuntimeCache, error) {
	return &interceptingControllerRuntimeCache{
		streamingClient: streamingClient,
		informers:       make(map[schema.GroupVersionKind]*streamingInformer),
		typeStore:       typeStore,
	}, nil
}

var _ cache.Cache = &interceptingControllerRuntimeCache{}

// Get retrieves an obj for the given object key from the Kubernetes Cluster.
// obj must be a struct pointer so that obj can be updated with the response
// returned by the Server.
func (c *interceptingControllerRuntimeCache) Get(ctx context.Context, key client.ObjectKey, obj client.Object, opts ...client.GetOption) error {
	if len(opts) != 0 {
		klog.Fatalf("interceptingControllerRuntimeCache: GET %v into %T with opts %v", key, obj, opts)
		panic("not implemented")
	}

	typeInfo, err := c.typeStore.getTypeInfo(obj)
	if err != nil {
		return err
	}
	informer, err := c.getOrCreateInformer(ctx, typeInfo)
	if err != nil {
		return err
	}
	return informer.Get(ctx, key, obj)
}

// List retrieves list of objects for a given namespace and list options. On a
// successful call, Items field in the list will be populated with the
// result returned from the server.
func (c *interceptingControllerRuntimeCache) List(ctx context.Context, list client.ObjectList, opts ...client.ListOption) error {
	panic("not implemented")
}

// GetInformer fetches or constructs an informer for the given object that corresponds to a single
// API kind and resource.
func (c *interceptingControllerRuntimeCache) GetInformer(ctx context.Context, obj client.Object) (cache.Informer, error) {
	typeInfo, err := c.typeStore.getTypeInfo(obj)
	if err != nil {
		return nil, err
	}

	return c.getOrCreateInformer(ctx, typeInfo)
}

func (c *interceptingControllerRuntimeCache) getOrCreateInformer(ctx context.Context, typeInfo *typeInfo) (*streamingInformer, error) {
	gvk := typeInfo.gvk

	c.mutex.Lock()
	defer c.mutex.Unlock()

	existing := c.informers[gvk]
	if existing != nil {
		return existing, nil
	}

	informer, err := newStreamingInformer(c.streamingClient, typeInfo)
	if err != nil {
		return nil, err
	}
	c.informers[gvk] = informer

	if c.started.Load() {
		if err := informer.Start(ctx); err != nil {
			return nil, fmt.Errorf("error starting informer: %w", err)
		}
	}

	return informer, nil
}

// GetInformerForKind is similar to GetInformer, except that it takes a group-version-kind, instead
// of the underlying object.
func (c *interceptingControllerRuntimeCache) GetInformerForKind(ctx context.Context, gvk schema.GroupVersionKind) (cache.Informer, error) {
	typeInfo, err := c.typeStore.getTypeInfoForGVK(gvk)
	if err != nil {
		return nil, err
	}

	return c.getOrCreateInformer(ctx, typeInfo)
}

// Start runs all the informers known to this cache until the context is closed.
// It blocks.
func (c *interceptingControllerRuntimeCache) Start(ctx context.Context) error {
	informers := c.snapshotInformers()

	for _, informer := range informers {
		if err := informer.Start(ctx); err != nil {
			return err
		}
	}
	c.started.Store(true)
	return nil
}

// WaitForCacheSync waits for all the caches to sync.  Returns false if it could not sync a cache.
func (c *interceptingControllerRuntimeCache) WaitForCacheSync(ctx context.Context) bool {
	informers := c.snapshotInformers()

	for _, informer := range informers {
		if !informer.WaitForCacheSync(ctx) {
			return false
		}
	}
	return true
}

// snapshotInformers is a helper function that returns a snapshot of the informers.
// It is used to avoid race conditions when iterating over the informers.
func (c *interceptingControllerRuntimeCache) snapshotInformers() []*streamingInformer {
	c.mutex.Lock()
	defer c.mutex.Unlock()
	return slices.Collect(maps.Values(c.informers))
}

// IndexFields adds an index with the given field name on the given object type
// by using the given function to extract the value for that field.  If you want
// compatibility with the Kubernetes API server, only return one key, and only use
// fields that the API server supports.  Otherwise, you can return multiple keys,
// and "equality" in the field selector means that at least one key matches the value.
// The FieldIndexer will automatically take care of indexing over namespace
// and supporting efficient all-namespace queries.
func (c *interceptingControllerRuntimeCache) IndexField(ctx context.Context, obj client.Object, field string, extractValue client.IndexerFunc) error {
	panic("not implemented")
}
