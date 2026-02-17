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
	"sigs.k8s.io/controller-runtime/pkg/client"
)

func resolveURLMapRefs(ctx context.Context, reader client.Reader, obj *krm.ComputeURLMap) error {
	if obj.Spec.ProjectRef != nil {
		if err := obj.Spec.ProjectRef.Normalize(ctx, reader, obj.Namespace); err != nil {
			return err
		}
	}

	if obj.Spec.DefaultService != nil {
		if err := resolveURLMapServiceRef(ctx, reader, obj.Namespace, obj.Spec.DefaultService); err != nil {
			return err
		}
	}

	for i := range obj.Spec.PathMatchers {
		pm := &obj.Spec.PathMatchers[i]
		if pm.DefaultService != nil {
			if err := resolveURLMapServiceRef(ctx, reader, obj.Namespace, pm.DefaultService); err != nil {
				return err
			}
		}
		for j := range pm.PathRules {
			pr := &pm.PathRules[j]
			if pr.Service != nil {
				if err := resolveURLMapServiceRef(ctx, reader, obj.Namespace, pr.Service); err != nil {
					return err
				}
			}
		}
		for j := range pm.RouteRules {
			rr := &pm.RouteRules[j]
			if rr.Service != nil {
				if err := resolveURLMapServiceRef(ctx, reader, obj.Namespace, rr.Service); err != nil {
					return err
				}
			}
			if rr.RouteAction != nil {
				for k := range rr.RouteAction.WeightedBackendServices {
					wbs := &rr.RouteAction.WeightedBackendServices[k]
					if wbs.BackendService != nil {
						external, err := wbs.BackendService.NormalizedExternal(ctx, reader, obj.Namespace)
						if err != nil {
							return err
						}
						wbs.BackendService.External = external
					}
				}
			}
		}
	}

	for i := range obj.Spec.Tests {
		t := &obj.Spec.Tests[i]
		if t.Service != nil {
			if err := resolveURLMapServiceRef(ctx, reader, obj.Namespace, t.Service); err != nil {
				return err
			}
		}
	}

	return nil
}

func resolveURLMapServiceRef(ctx context.Context, reader client.Reader, namespace string, ref *krm.ComputeURLMapServiceRef) error {
	if ref == nil {
		return nil
	}
	if ref.BackendBucketRef != nil {
		external, err := ref.BackendBucketRef.NormalizedExternal(ctx, reader, namespace)
		if err != nil {
			return err
		}
		ref.BackendBucketRef.External = external
	}
	if ref.BackendServiceRef != nil {
		external, err := ref.BackendServiceRef.NormalizedExternal(ctx, reader, namespace)
		if err != nil {
			return err
		}
		ref.BackendServiceRef.External = external
	}
	if ref.BackendBucketRef != nil && ref.BackendServiceRef != nil {
		return fmt.Errorf("cannot specify both backendBucketRef and backendServiceRef")
	}
	return nil
}
