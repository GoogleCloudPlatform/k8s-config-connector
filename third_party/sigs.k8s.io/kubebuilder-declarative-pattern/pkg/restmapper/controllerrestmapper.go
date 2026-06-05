package restmapper

import (
	"context"
	"fmt"

	"k8s.io/apimachinery/pkg/api/meta"
	"k8s.io/apimachinery/pkg/runtime/schema"
	"k8s.io/client-go/discovery"
)

// ControllerRESTMapper is a meta.RESTMapper that is optimized for controllers.
// It caches results in memory, and minimizes discovery because we don't need shortnames etc in controllers.
// Controllers primarily need to map from GVK -> GVR.
type ControllerRESTMapper struct {
	uncached discovery.DiscoveryInterface
	cache    *cache
}

var _ meta.RESTMapper = &ControllerRESTMapper{}

// KindFor takes a partial resource and returns the single match.  Returns an error if there are multiple matches
func (m *ControllerRESTMapper) KindFor(resource schema.GroupVersionResource) (schema.GroupVersionKind, error) {
	ctx := context.TODO()
	kinds, err := m.cache.KindsFor(ctx, resource, m.uncached)
	if err != nil {
		return schema.GroupVersionKind{}, err
	}
	if len(kinds) == 0 {
		return schema.GroupVersionKind{}, fmt.Errorf("found no matching kinds for %v", resource.String())
	}
	if len(kinds) > 1 {
		return schema.GroupVersionKind{}, fmt.Errorf("found multiple kinds for %v: %v", resource.String(), kinds)
	}
	return kinds[0], nil
}

// KindsFor takes a partial resource and returns the list of potential kinds in priority order
func (m *ControllerRESTMapper) KindsFor(resource schema.GroupVersionResource) ([]schema.GroupVersionKind, error) {
	return nil, fmt.Errorf("ControllerRESTMapper does not support KindsFor operation")
}

// ResourceFor takes a partial resource and returns the single match.  Returns an error if there are multiple matches
func (m *ControllerRESTMapper) ResourceFor(input schema.GroupVersionResource) (schema.GroupVersionResource, error) {
	return schema.GroupVersionResource{}, fmt.Errorf("ControllerRESTMapper does not support ResourceFor operation")
}

// ResourcesFor takes a partial resource and returns the list of potential resource in priority order
func (m *ControllerRESTMapper) ResourcesFor(input schema.GroupVersionResource) ([]schema.GroupVersionResource, error) {
	return nil, fmt.Errorf("ControllerRESTMapper does not support ResourcesFor operation")
}

// RESTMapping identifies a preferred resource mapping for the provided group kind.
func (m *ControllerRESTMapper) RESTMapping(gk schema.GroupKind, versions ...string) (*meta.RESTMapping, error) {
	restMappings, err := m.RESTMappings(gk, versions...)
	if err != nil {
		return nil, err
	}

	if len(restMappings) >= 1 {
		// RESTMappings should return preferred version first.
		return restMappings[0], nil
	}

	return nil, &meta.NoKindMatchError{GroupKind: gk, SearchedVersions: versions}
}

// RESTMappings returns all resource mappings for the provided group kind if no
// version search is provided. Otherwise identifies a preferred resource mapping for
// the provided version(s).
func (m *ControllerRESTMapper) RESTMappings(gk schema.GroupKind, versions ...string) ([]*meta.RESTMapping, error) {
	ctx := context.TODO()

	var mappings []*meta.RESTMapping

	if len(versions) == 0 {
		allGroups, err := m.cache.fetchAllGroups(ctx, m.uncached)
		if err != nil {
			return nil, err
		}
		group, found := allGroups[gk.Group]
		if !found {
			return nil, &meta.NoResourceMatchError{PartialResource: schema.GroupVersionResource{Group: gk.Group, Resource: gk.Kind}}
		}

		if group.PreferredVersion.Version != "" {
			gv := schema.GroupVersion{Group: gk.Group, Version: group.PreferredVersion.Version}
			mapping, err := m.cache.findRESTMapping(ctx, m.uncached, gv, gk.Kind)
			if err != nil {
				return nil, err
			}
			if mapping != nil {
				mappings = append(mappings, mapping)
			}
		}

		for i := range group.Versions {
			gv := schema.GroupVersion{Group: gk.Group, Version: group.Versions[i].Version}
			if gv.Version == group.PreferredVersion.Version {
				continue
			}
			mapping, err := m.cache.findRESTMapping(ctx, m.uncached, gv, gk.Kind)
			if err != nil {
				return nil, err
			}
			if mapping != nil {
				mappings = append(mappings, mapping)
			}
		}
	} else {
		for _, version := range versions {
			gv := schema.GroupVersion{Group: gk.Group, Version: version}
			mapping, err := m.cache.findRESTMapping(ctx, m.uncached, gv, gk.Kind)
			if err != nil {
				return nil, err
			}
			if mapping != nil {
				mappings = append(mappings, mapping)
			}
		}
	}

	if len(mappings) == 0 {
		return nil, &meta.NoResourceMatchError{PartialResource: schema.GroupVersionResource{Group: gk.Group, Resource: gk.Kind}}
	}
	return mappings, nil
}

func (m *ControllerRESTMapper) ResourceSingularizer(resource string) (singular string, err error) {
	return "", fmt.Errorf("ControllerRESTMapper does not support ResourceSingularizer operation")
}
