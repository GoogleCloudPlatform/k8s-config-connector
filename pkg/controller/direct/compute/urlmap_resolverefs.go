// Copyright 2026 Google LLC
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
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/k8s"
	apierrors "k8s.io/apimachinery/pkg/api/errors"
	"k8s.io/apimachinery/pkg/apis/meta/v1/unstructured"
	"k8s.io/apimachinery/pkg/runtime/schema"
	"k8s.io/apimachinery/pkg/types"
	"sigs.k8s.io/controller-runtime/pkg/client"
)

func resolveResourceRef(ctx context.Context, reader client.Reader, src client.Object, ref *v1alpha1.ResourceRef, gvk schema.GroupVersionKind) error {
	if ref == nil {
		return nil
	}

	if ref.External != "" && ref.Name != "" {
		return fmt.Errorf("cannot specify both name and external on %s reference", gvk.Kind)
	}

	if ref.External != "" {
		return nil
	}

	if ref.Name == "" {
		return nil
	}

	key := types.NamespacedName{
		Namespace: ref.Namespace,
		Name:      ref.Name,
	}
	if key.Namespace == "" {
		key.Namespace = src.GetNamespace()
	}

	u := &unstructured.Unstructured{}
	u.SetGroupVersionKind(gvk)
	if err := reader.Get(ctx, key, u); err != nil {
		if apierrors.IsNotFound(err) {
			return k8s.NewReferenceNotFoundError(u.GroupVersionKind(), key)
		}
		return fmt.Errorf("reading referenced %s %s: %w", gvk, key, err)
	}

	// targetField: self_link
	selfLink, _, err := unstructured.NestedString(u.Object, "status", "selfLink")
	if err != nil || selfLink == "" {
		return fmt.Errorf("cannot get selfLink for referenced %s %v (status.selfLink is empty)", u.GetKind(), u.GetNamespace())
	}
	ref.External = selfLink
	return nil
}

func resolveURLMapRefs(ctx context.Context, reader client.Reader, obj *krm.ComputeURLMap) error {
	backendServiceGVK := schema.GroupVersionKind{Group: "compute.cnrm.cloud.google.com", Version: "v1beta1", Kind: "ComputeBackendService"}
	backendBucketGVK := schema.GroupVersionKind{Group: "compute.cnrm.cloud.google.com", Version: "v1beta1", Kind: "ComputeBackendBucket"}

	resolveDefaultService := func(s *krm.ComputeURLMapDefaultService) error {
		if s == nil {
			return nil
		}
		if err := resolveResourceRef(ctx, reader, obj, s.BackendServiceRef, backendServiceGVK); err != nil {
			return err
		}
		if err := resolveResourceRef(ctx, reader, obj, s.BackendBucketRef, backendBucketGVK); err != nil {
			return err
		}
		return nil
	}

	resolveRouteAction := func(ra *krm.ComputeURLMapHTTPRouteAction) error {
		if ra == nil {
			return nil
		}
		if ra.RequestMirrorPolicy != nil && ra.RequestMirrorPolicy.BackendServiceRef != nil {
			if _, err := ra.RequestMirrorPolicy.BackendServiceRef.NormalizedExternal(ctx, reader, obj.Namespace); err != nil {
				return err
			}
		}
		for i := range ra.WeightedBackendServices {
			wbs := &ra.WeightedBackendServices[i]
			if wbs.BackendServiceRef != nil {
				if _, err := wbs.BackendServiceRef.NormalizedExternal(ctx, reader, obj.Namespace); err != nil {
					return err
				}
			}
		}
		return nil
	}

	if err := resolveDefaultService(obj.Spec.DefaultService); err != nil {
		return err
	}
	if err := resolveRouteAction(obj.Spec.DefaultRouteAction); err != nil {
		return err
	}

	for i := range obj.Spec.PathMatcher {
		pm := &obj.Spec.PathMatcher[i]
		if err := resolveDefaultService(pm.DefaultService); err != nil {
			return err
		}
		if err := resolveRouteAction(pm.DefaultRouteAction); err != nil {
			return err
		}
		for j := range pm.PathRule {
			pr := &pm.PathRule[j]
			if err := resolveDefaultService(pr.Service); err != nil {
				return err
			}
			if err := resolveRouteAction(pr.RouteAction); err != nil {
				return err
			}
		}
		for j := range pm.RouteRules {
			rr := &pm.RouteRules[j]
			if err := resolveDefaultService(rr.Service); err != nil {
				return err
			}
			if err := resolveRouteAction(rr.RouteAction); err != nil {
				return err
			}
		}
	}

	for i := range obj.Spec.Test {
		t := &obj.Spec.Test[i]
		if err := resolveDefaultService(t.Service); err != nil {
			return err
		}
	}

	return nil
}
