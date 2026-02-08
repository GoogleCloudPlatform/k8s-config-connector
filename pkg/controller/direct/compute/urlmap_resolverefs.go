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

package compute

import (
	"context"
	"fmt"

	krm "github.com/GoogleCloudPlatform/k8s-config-connector/apis/compute/v1beta1"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/clients/generated/apis/k8s/v1alpha1"
	"k8s.io/apimachinery/pkg/apis/meta/v1/unstructured"
	"k8s.io/apimachinery/pkg/runtime/schema"
	"k8s.io/apimachinery/pkg/types"
	"sigs.k8s.io/controller-runtime/pkg/client"
)

func resolveComputeURLMapRefs(ctx context.Context, reader client.Reader, obj *krm.ComputeURLMap) error {
	// Resolve DefaultService
	if obj.Spec.DefaultService != nil {
		if err := resolveUrlmapDefaultServiceRefs(ctx, reader, obj, obj.Spec.DefaultService); err != nil {
			return err
		}
	}

	// Resolve DefaultRouteAction
	if obj.Spec.DefaultRouteAction != nil {
		if err := resolveUrlmapDefaultRouteActionRefs(ctx, reader, obj, obj.Spec.DefaultRouteAction); err != nil {
			return err
		}
	}

	// Resolve PathMatcher
	for i := range obj.Spec.PathMatcher {
		pm := &obj.Spec.PathMatcher[i]
		if pm.DefaultService != nil {
			if err := resolveUrlmapDefaultServiceRefs(ctx, reader, obj, pm.DefaultService); err != nil {
				return err
			}
		}
		if pm.DefaultRouteAction != nil {
			if err := resolveUrlmapDefaultRouteActionRefs(ctx, reader, obj, pm.DefaultRouteAction); err != nil {
				return err
			}
		}
		for j := range pm.PathRule {
			pr := &pm.PathRule[j]
			if pr.Service != nil {
				if err := resolveUrlmapServiceRefs(ctx, reader, obj, pr.Service); err != nil {
					return err
				}
			}
			if pr.RouteAction != nil {
				if err := resolveUrlmapRouteActionRefs(ctx, reader, obj, pr.RouteAction); err != nil {
					return err
				}
			}
		}
		for j := range pm.RouteRules {
			rr := &pm.RouteRules[j]
			if rr.RouteAction != nil {
				if err := resolveUrlmapRouteActionRefs(ctx, reader, obj, rr.RouteAction); err != nil {
					return err
				}
			}
		}
	}

	// Resolve Test
	for i := range obj.Spec.Test {
		t := &obj.Spec.Test[i]
		if err := resolveUrlmapServiceRefs(ctx, reader, obj, &t.Service); err != nil {
			return err
		}
	}

	return nil
}

func resolveResourceRef(ctx context.Context, reader client.Reader, obj client.Object, ref *v1alpha1.ResourceRef, gvk schema.GroupVersionKind, targetField string) error {
	if ref == nil {
		return nil
	}

	if ref.External != "" {
		if ref.Name != "" {
			return fmt.Errorf("cannot specify both name and external on reference")
		}
		return nil
	}

	if ref.Name == "" {
		return fmt.Errorf("must specify either name or external on reference")
	}

	key := types.NamespacedName{
		Namespace: ref.Namespace,
		Name:      ref.Name,
	}
	if key.Namespace == "" {
		key.Namespace = obj.GetNamespace()
	}

	resource, err := resolveResourceName(ctx, reader, key, gvk)
	if err != nil {
		return err
	}

	val, _, err := unstructured.NestedString(resource.Object, "status", targetField)
	if err != nil || val == "" {
		return fmt.Errorf("cannot get %s for referenced %s %v (status.%s is empty)", targetField, resource.GetKind(), resource.GetNamespace(), targetField)
	}
	ref.External = val
	return nil
}

func resolveUrlmapDefaultServiceRefs(ctx context.Context, reader client.Reader, obj *krm.ComputeURLMap, ds *krm.UrlmapDefaultService) error {
	if ds.BackendBucketRef != nil {
		if err := resolveResourceRef(ctx, reader, obj, ds.BackendBucketRef, schema.GroupVersionKind{
			Group:   "compute.cnrm.cloud.google.com",
			Version: "v1beta1",
			Kind:    "ComputeBackendBucket",
		}, "selfLink"); err != nil {
			return err
		}
	}
	if ds.BackendServiceRef != nil {
		if err := resolveResourceRef(ctx, reader, obj, ds.BackendServiceRef, schema.GroupVersionKind{
			Group:   "compute.cnrm.cloud.google.com",
			Version: "v1beta1",
			Kind:    "ComputeBackendService",
		}, "selfLink"); err != nil {
			return err
		}
	}
	return nil
}

func resolveUrlmapServiceRefs(ctx context.Context, reader client.Reader, obj *krm.ComputeURLMap, s *krm.UrlmapService) error {
	if s.BackendBucketRef != nil {
		if err := resolveResourceRef(ctx, reader, obj, s.BackendBucketRef, schema.GroupVersionKind{
			Group:   "compute.cnrm.cloud.google.com",
			Version: "v1beta1",
			Kind:    "ComputeBackendBucket",
		}, "selfLink"); err != nil {
			return err
		}
	}
	if s.BackendServiceRef != nil {
		if err := resolveResourceRef(ctx, reader, obj, s.BackendServiceRef, schema.GroupVersionKind{
			Group:   "compute.cnrm.cloud.google.com",
			Version: "v1beta1",
			Kind:    "ComputeBackendService",
		}, "selfLink"); err != nil {
			return err
		}
	}
	return nil
}

func resolveUrlmapDefaultRouteActionRefs(ctx context.Context, reader client.Reader, obj *krm.ComputeURLMap, dra *krm.UrlmapDefaultRouteAction) error {
	if dra.RequestMirrorPolicy != nil {
		// BackendServiceRef is value type in RequestMirrorPolicy
		// But in krm.UrlmapRequestMirrorPolicy it is v1alpha1.ResourceRef (value)
		// We need to pass address of it
		if err := resolveResourceRef(ctx, reader, obj, &dra.RequestMirrorPolicy.BackendServiceRef, schema.GroupVersionKind{
			Group:   "compute.cnrm.cloud.google.com",
			Version: "v1beta1",
			Kind:    "ComputeBackendService",
		}, "selfLink"); err != nil {
			return err
		}
	}
	for i := range dra.WeightedBackendServices {
		wbs := &dra.WeightedBackendServices[i]
		if err := resolveResourceRef(ctx, reader, obj, &wbs.BackendServiceRef, schema.GroupVersionKind{
			Group:   "compute.cnrm.cloud.google.com",
			Version: "v1beta1",
			Kind:    "ComputeBackendService",
		}, "selfLink"); err != nil {
			return err
		}
	}
	return nil
}

func resolveUrlmapRouteActionRefs(ctx context.Context, reader client.Reader, obj *krm.ComputeURLMap, ra *krm.UrlmapRouteAction) error {
	if ra.RequestMirrorPolicy != nil {
		if err := resolveResourceRef(ctx, reader, obj, &ra.RequestMirrorPolicy.BackendServiceRef, schema.GroupVersionKind{
			Group:   "compute.cnrm.cloud.google.com",
			Version: "v1beta1",
			Kind:    "ComputeBackendService",
		}, "selfLink"); err != nil {
			return err
		}
	}
	for i := range ra.WeightedBackendServices {
		wbs := &ra.WeightedBackendServices[i]
		if err := resolveResourceRef(ctx, reader, obj, &wbs.BackendServiceRef, schema.GroupVersionKind{
			Group:   "compute.cnrm.cloud.google.com",
			Version: "v1beta1",
			Kind:    "ComputeBackendService",
		}, "selfLink"); err != nil {
			return err
		}
	}
	return nil
}