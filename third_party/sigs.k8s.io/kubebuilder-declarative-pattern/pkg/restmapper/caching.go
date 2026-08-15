package restmapper

import (
	"context"
	"fmt"
	"strings"
	"sync"

	apierrors "k8s.io/apimachinery/pkg/api/errors"
	"k8s.io/apimachinery/pkg/api/meta"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime/schema"
	"k8s.io/client-go/discovery"
	"k8s.io/klog/v2"
	"sigs.k8s.io/controller-runtime/pkg/log"
)

// cache is our cache of schema information.
type cache struct {
	mutex               sync.Mutex
	cachedAllGroups     map[string]metav1.APIGroup
	cachedGroupVersions map[schema.GroupVersion]*cachedGroupVersion
}

// newCache is the constructor for a cache.
func newCache() *cache {
	return &cache{
		cachedGroupVersions: make(map[schema.GroupVersion]*cachedGroupVersion),
	}
}

// fetchAllGroups returns the APIGroup for the specified group, querying discovery if not cached.
// If not found, returns APIGroup{}, false, nil
func (c *cache) fetchAllGroups(ctx context.Context, discovery discovery.DiscoveryInterface) (map[string]metav1.APIGroup, error) {
	log := log.FromContext(ctx)

	c.mutex.Lock()
	defer c.mutex.Unlock()

	if c.cachedAllGroups == nil {
		log.Info("discovering server groups")
		serverGroups, err := discovery.ServerGroups()
		if err != nil {
			klog.Infof("unexpected error from ServerGroups: %v", err)
			return nil, fmt.Errorf("error from ServerGroups: %w", err)
		}

		groups := make(map[string]metav1.APIGroup)
		for i := range serverGroups.Groups {
			group := &serverGroups.Groups[i]
			groups[group.Name] = *group
		}
		c.cachedAllGroups = groups
	}

	return c.cachedAllGroups, nil
}

// cachedGroupVersion caches (all) the resource information for a particular groupversion.
type cachedGroupVersion struct {
	gv                    schema.GroupVersion
	mutex                 sync.Mutex
	cachedServerResources map[string]cachedResource
}

// cachedResource caches the information for a particular resource.
type cachedResource struct {
	resource string
	scope    meta.RESTScope
	gvk      schema.GroupVersionKind
}

func (r *cachedResource) GVR() schema.GroupVersionResource {
	return r.gvk.GroupVersion().WithResource(r.resource)
}

func (r *cachedResource) GVK() schema.GroupVersionKind {
	return r.gvk
}

func (r *cachedResource) RESTMapping() *meta.RESTMapping {
	return &meta.RESTMapping{
		Resource:         r.GVR(),
		GroupVersionKind: r.GVK(),
		Scope:            r.scope,
	}
}

// KindsFor finds out the Kind from the GVR in the cache. If the GVR version is not given, we will iterate all the matching
// GR in the cache and return the first matching one.
func (c *cache) KindsFor(ctx context.Context, gvr schema.GroupVersionResource, discovery discovery.DiscoveryInterface) ([]schema.GroupVersionKind, error) {
	var matches []schema.GroupVersionKind
	if gvr.Version != "" {
		gv := gvr.GroupVersion()
		cachedGV := c.cacheForGroupVersion(gv)

		all, err := cachedGV.KindsFor(ctx, gvr, discovery)
		if err != nil {
			return nil, err
		}

		matches = append(matches, all...)

		return matches, nil
	}

	allGroups, err := c.fetchAllGroups(ctx, discovery)
	if err != nil {
		return nil, err
	}
	for groupName, group := range allGroups {
		if groupName != gvr.Group {
			continue
		}
		for _, version := range group.Versions {
			gv := schema.GroupVersion{Group: groupName, Version: version.Version}
			cachedGV := c.cacheForGroupVersion(gv)

			all, err := cachedGV.KindsFor(ctx, gvr, discovery)
			if err != nil {
				return nil, err
			}

			matches = append(matches, all...)
		}
	}
	return matches, nil
}

func (c *cache) cacheForGroupVersion(gv schema.GroupVersion) *cachedGroupVersion {
	c.mutex.Lock()
	cached := c.cachedGroupVersions[gv]
	if cached == nil {
		cached = &cachedGroupVersion{gv: gv}
		c.cachedGroupVersions[gv] = cached
	}
	c.mutex.Unlock()
	return cached
}

// findRESTMapping returns the RESTMapping for the specified GVK, querying discovery if not cached.
func (c *cache) findRESTMapping(ctx context.Context, discovery discovery.DiscoveryInterface, gv schema.GroupVersion, kind string) (*meta.RESTMapping, error) {
	cachedGV := c.cacheForGroupVersion(gv)
	return cachedGV.findRESTMapping(ctx, discovery, kind)
}

// findRESTMapping returns the RESTMapping for the specified GVK, querying discovery if not cached.
func (c *cachedGroupVersion) findRESTMapping(ctx context.Context, discovery discovery.DiscoveryInterface, kind string) (*meta.RESTMapping, error) {
	resources, err := c.fetchServerResources(ctx, discovery)
	if err != nil {
		return nil, err
	}

	for _, resource := range resources {
		if resource.GVK().Kind == kind {
			return resource.RESTMapping(), nil
		}
	}
	return nil, nil
}

// fetch returns the metadata, fetching it if not cached.
func (c *cachedGroupVersion) fetchServerResources(ctx context.Context, discovery discovery.DiscoveryInterface) (map[string]cachedResource, error) {
	log := log.FromContext(ctx)

	c.mutex.Lock()
	defer c.mutex.Unlock()

	if c.cachedServerResources != nil {
		return c.cachedServerResources, nil
	}

	log.Info("discovering server resources for group/version", "gv", c.gv.String())
	resourceList, err := discovery.ServerResourcesForGroupVersion(c.gv.String())
	if err != nil {
		// We treat "no match" as an empty result, but any other error percolates back up
		if meta.IsNoMatchError(err) || apierrors.IsNotFound(err) {
			return nil, nil
		} else {
			klog.Infof("unexpected error from ServerResourcesForGroupVersion(%v): %v", c.gv, err)
			return nil, fmt.Errorf("error from ServerResourcesForGroupVersion(%v): %w", c.gv, err)
		}
	}

	result := make(map[string]cachedResource)
	for i := range resourceList.APIResources {
		resource := resourceList.APIResources[i]

		// if we have a slash, then this is a subresource and we shouldn't create mappings for those.
		if strings.Contains(resource.Name, "/") {
			continue
		}

		scope := meta.RESTScopeRoot
		if resource.Namespaced {
			scope = meta.RESTScopeNamespace
		}
		result[resource.Name] = cachedResource{
			resource: resource.Name,
			scope:    scope,
			gvk:      c.gv.WithKind(resource.Kind),
		}
	}
	c.cachedServerResources = result
	return result, nil
}

func (c *cachedGroupVersion) KindsFor(ctx context.Context, filterGVR schema.GroupVersionResource, discovery discovery.DiscoveryInterface) ([]schema.GroupVersionKind, error) {
	serverResources, err := c.fetchServerResources(ctx, discovery)
	if err != nil {
		return nil, err
	}
	var matches []schema.GroupVersionKind
	for _, resource := range serverResources {
		resourceGVR := resource.GVR()

		if resourceGVR.Group != filterGVR.Group {
			continue
		}
		if filterGVR.Version != "" && resourceGVR.Version != filterGVR.Version {
			continue
		}
		if filterGVR.Resource != "" && resourceGVR.Resource != filterGVR.Resource {
			continue
		}
		matches = append(matches, resource.GVK())
	}

	return matches, nil
}
