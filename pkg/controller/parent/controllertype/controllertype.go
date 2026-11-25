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

// Package controllertype decides which controller (Terraform, DCL, Direct, ...) reconciles a
// Config Connector object. The parent reconciler and the admission webhooks both use it, so they
// always agree on which controller owns an object.
package controllertype

import (
	"context"
	"fmt"

	bigquerykrm "github.com/GoogleCloudPlatform/k8s-config-connector/apis/bigquery/v1beta1"
	computekrm "github.com/GoogleCloudPlatform/k8s-config-connector/apis/compute/v1beta1"
	secretkrm "github.com/GoogleCloudPlatform/k8s-config-connector/apis/secretmanager/v1beta1"
	spannerkrm "github.com/GoogleCloudPlatform/k8s-config-connector/apis/spanner/v1beta1"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct"
	kccpredicate "github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/predicate"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/resourceconfig"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/k8s"

	"k8s.io/apimachinery/pkg/apis/meta/v1/unstructured"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/apimachinery/pkg/runtime/schema"
)

// OverridesFunc returns the controllerOverrides of the ConfigConnectorContext in the object's
// namespace, keyed by "Kind.group". ForObject calls it only when no earlier rule has decided.
type OverridesFunc func(ctx context.Context) (map[string]k8s.ReconcilerType, error)

// ForObject returns the controller type for u. Rules are applied in this order:
//  1. the alpha.cnrm.cloud.google.com/reconciler: direct annotation,
//  2. per-kind special cases that route some objects to Direct based on their spec,
//  3. the controllerOverrides returned by overrides (may be nil),
//  4. the kind's default controller.
//
// The result is not checked against the kind's supported controllers.
func ForObject(ctx context.Context, gvk schema.GroupVersionKind, u *unstructured.Unstructured, overrides OverridesFunc) (k8s.ReconcilerType, error) {
	// Check for resource annotation
	annotations := u.GetAnnotations()
	if annotations[k8s.AlphaReconcilerAnnotation] == "direct" {
		return k8s.ReconcilerTypeDirect, nil
	}

	// Special case handling. Will be removed after the resources have turned on direct as default.
	if gvk.Kind == "BigQueryTable" {
		obj := &bigquerykrm.BigQueryTable{}
		if _, ok := annotations[kccpredicate.AnnotationUnmanaged]; ok {
			return k8s.ReconcilerTypeDirect, nil
		}
		if err := runtime.DefaultUnstructuredConverter.FromUnstructured(u.Object, &obj); err != nil {
			return "", fmt.Errorf("error converting to %T: %w", obj, err)
		}
		if obj.Spec.Labels != nil {
			return k8s.ReconcilerTypeDirect, nil
		}
	}
	if gvk.Kind == "ComputeForwardingRule" {
		obj := &computekrm.ComputeForwardingRule{}
		if err := runtime.DefaultUnstructuredConverter.FromUnstructured(u.Object, &obj); err != nil {
			return "", fmt.Errorf("error converting to %T: %w", obj, err)
		}
		if obj.Spec.Target != nil && obj.Spec.Target.GoogleAPIsBundle != nil {
			return k8s.ReconcilerTypeDirect, nil
		}
	}
	if gvk.Kind == "ComputeTargetTCPProxy" {
		obj := &computekrm.ComputeTargetTCPProxy{}
		if err := runtime.DefaultUnstructuredConverter.FromUnstructured(u.Object, &obj); err != nil {
			return "", fmt.Errorf("error converting to %T: %w", obj, err)
		}
		if obj.Spec.Location != nil && obj.Spec.Location != direct.PtrTo("global") {
			return k8s.ReconcilerTypeDirect, nil
		}
	}
	if gvk.Kind == "SecretManagerSecret" {
		obj := &secretkrm.SecretManagerSecret{}
		if err := runtime.DefaultUnstructuredConverter.FromUnstructured(u.Object, &obj); err != nil {
			return "", fmt.Errorf("error converting to %T: %w", obj, err)
		}
		if obj.Spec.Labels != nil {
			return k8s.ReconcilerTypeDirect, nil
		}
	}
	if gvk.Kind == "SpannerInstance" {
		obj := &spannerkrm.SpannerInstance{}
		if err := runtime.DefaultUnstructuredConverter.FromUnstructured(u.Object, &obj); err != nil {
			return "", fmt.Errorf("error converting to %T: %w", obj, err)
		}
		if obj.Spec.DefaultBackupScheduleType != nil || obj.Spec.Labels != nil || obj.Spec.Edition != nil || obj.Spec.AutoscalingConfig != nil {
			return k8s.ReconcilerTypeDirect, nil
		}
	}

	// Check for CCC setting
	if overrides != nil {
		controllerOverrides, err := overrides(ctx)
		if err != nil {
			return "", err
		}
		if v, ok := controllerOverrides[gvk.GroupKind().String()]; ok {
			return v, nil
		}
	}

	// Fallback to static config
	gk := gvk.GroupKind()
	resourcesControllersConfig := resourceconfig.LoadConfig()
	config, err := resourcesControllersConfig.GetControllersForGVK(gvk)
	if err != nil {
		return "", fmt.Errorf("error getting controller config found for GroupKind %v", gk)
	}
	return config.DefaultController, nil
}
