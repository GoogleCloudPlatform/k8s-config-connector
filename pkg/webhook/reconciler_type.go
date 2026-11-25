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

package webhook

import (
	"context"
	"fmt"

	operatorv1beta1 "github.com/GoogleCloudPlatform/k8s-config-connector/operator/pkg/apis/core/v1beta1"
	operatork8s "github.com/GoogleCloudPlatform/k8s-config-connector/operator/pkg/k8s"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/parent/controllertype"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/resourceconfig"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/k8s"

	apierrors "k8s.io/apimachinery/pkg/api/errors"
	"k8s.io/apimachinery/pkg/api/meta"
	"k8s.io/apimachinery/pkg/apis/meta/v1/unstructured"
	"k8s.io/apimachinery/pkg/types"
	"k8s.io/klog/v2"
	"sigs.k8s.io/controller-runtime/pkg/client"
)

// reconcilerTypeForObject returns the controller that will reconcile obj. It makes the same
// decision as the parent reconciler (see controllertype.ForObject), including the fallback to the
// kind's default controller when the chosen controller isn't supported. The boolean result is false
// when the decision can't be made, for example when the ConfigConnectorContext can't be read;
// callers then keep the kind-based behavior.
func reconcilerTypeForObject(ctx context.Context, reader client.Reader, obj *unstructured.Unstructured) (k8s.ReconcilerType, bool) {
	gvk := obj.GroupVersionKind()
	reconcilerType, err := controllertype.ForObject(ctx, gvk, obj, func(ctx context.Context) (map[string]k8s.ReconcilerType, error) {
		return controllerOverrides(ctx, reader, obj.GetNamespace())
	})
	if err != nil {
		klog.Warningf("cannot determine the controller for %v %v/%v, using kind-based rules: %v", gvk.Kind, obj.GetNamespace(), obj.GetName(), err)
		return "", false
	}
	if resourceconfig.IsControllerSupported(gvk, reconcilerType) {
		return reconcilerType, true
	}
	config, err := resourceconfig.LoadConfig().GetControllersForGVK(gvk)
	if err != nil {
		klog.Warningf("cannot determine the controller for %v %v/%v, using kind-based rules: %v", gvk.Kind, obj.GetNamespace(), obj.GetName(), err)
		return "", false
	}
	return config.DefaultController, true
}

// controllerOverrides returns the controllerOverrides of the ConfigConnectorContext in namespace.
// It returns nil when Config Connector runs in cluster mode, or when there is no ConfigConnector or
// ConfigConnectorContext object. Unlike kccstate.FetchLiveKCCState, a missing
// ConfigConnectorContext isn't an error, because the webhooks also admit objects in namespaces that
// no controller manages yet.
func controllerOverrides(ctx context.Context, reader client.Reader, namespace string) (map[string]k8s.ReconcilerType, error) {
	if reader == nil {
		return nil, nil
	}

	cc := &unstructured.Unstructured{}
	cc.SetGroupVersionKind(operatorv1beta1.ConfigConnectorGroupVersionKind)
	if err := reader.Get(ctx, types.NamespacedName{Name: operatorv1beta1.ConfigConnectorAllowedName}, cc); err != nil {
		if apierrors.IsNotFound(err) || meta.IsNoMatchError(err) {
			return nil, nil
		}
		return nil, fmt.Errorf("getting ConfigConnector: %w", err)
	}
	mode, _, err := unstructured.NestedString(cc.Object, "spec", "mode")
	if err != nil {
		return nil, fmt.Errorf("reading spec.mode of ConfigConnector: %w", err)
	}
	if mode != "" && mode != operatork8s.NamespacedMode {
		return nil, nil
	}

	ccc := &unstructured.Unstructured{}
	ccc.SetGroupVersionKind(operatorv1beta1.ConfigConnectorContextGroupVersionKind)
	key := types.NamespacedName{Namespace: namespace, Name: operatorv1beta1.ConfigConnectorContextAllowedName}
	if err := reader.Get(ctx, key, ccc); err != nil {
		if apierrors.IsNotFound(err) || meta.IsNoMatchError(err) {
			return nil, nil
		}
		return nil, fmt.Errorf("getting ConfigConnectorContext in namespace %q: %w", namespace, err)
	}
	raw, _, err := unstructured.NestedStringMap(ccc.Object, "spec", "experiments", "controllerOverrides")
	if err != nil {
		return nil, fmt.Errorf("reading controllerOverrides of ConfigConnectorContext in namespace %q: %w", namespace, err)
	}
	overrides := make(map[string]k8s.ReconcilerType, len(raw))
	for kind, reconcilerType := range raw {
		overrides[kind] = k8s.ReconcilerType(reconcilerType)
	}
	return overrides, nil
}
