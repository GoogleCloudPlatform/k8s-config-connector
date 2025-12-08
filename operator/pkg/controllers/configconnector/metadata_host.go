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

package configconnector

import (
	"context"
	"fmt"

	"k8s.io/apimachinery/pkg/apis/meta/v1/unstructured"
	"sigs.k8s.io/controller-runtime/pkg/log"
	"sigs.k8s.io/kubebuilder-declarative-pattern/pkg/patterns/declarative"
	"sigs.k8s.io/kubebuilder-declarative-pattern/pkg/patterns/declarative/pkg/manifest"

	corev1beta1 "github.com/GoogleCloudPlatform/k8s-config-connector/operator/pkg/apis/core/v1beta1"
)

// transformForMetadataHost returns a transform function that injects the GCE_METADATA_HOST
// environment variable into controller containers when spec.metadataHost is set.
// This enables Config Connector to work in IPv6-only GKE clusters where the default
// metadata IP (169.254.169.254) is not reachable.
func (r *Reconciler) transformForMetadataHost() declarative.ObjectTransform {
	return func(ctx context.Context, o declarative.DeclarativeObject, m *manifest.Objects) error {
		cc, ok := o.(*corev1beta1.ConfigConnector)
		if !ok {
			return fmt.Errorf("expected the resource to be a ConfigConnector, but it was of type %T", o)
		}

		if err := r.applyMetadataHost(ctx, cc, m); err != nil {
			return fmt.Errorf("error applying metadata host transform: %w", err)
		}
		return nil
	}
}

func (r *Reconciler) applyMetadataHost(ctx context.Context, cc *corev1beta1.ConfigConnector, m *manifest.Objects) error {
	if cc.Spec.MetadataHost == "" {
		return nil
	}

	log := log.FromContext(ctx)
	log.Info("applying GCE_METADATA_HOST environment variable", "metadataHost", cc.Spec.MetadataHost)

	for _, item := range m.Items {
		// Apply to workloads that access GCP metadata (StatefulSets, Deployments, DaemonSets)
		if item.Kind != "StatefulSet" && item.Kind != "Deployment" && item.Kind != "DaemonSet" {
			continue
		}

		if err := item.MutateContainers(func(container map[string]interface{}) error {
			return addMetadataHostEnvVar(container, cc.Spec.MetadataHost)
		}); err != nil {
			return fmt.Errorf("failed to apply metadata host to %s/%s: %w", item.Kind, item.GetName(), err)
		}
		log.V(1).Info("injected GCE_METADATA_HOST into containers", "kind", item.Kind, "name", item.GetName())
	}

	return nil
}

// addMetadataHostEnvVar adds the GCE_METADATA_HOST environment variable to a container spec.
func addMetadataHostEnvVar(container map[string]interface{}, metadataHost string) error {
	// Get existing env vars or create empty slice
	existingEnv, _, _ := unstructured.NestedSlice(container, "env")

	// Check if GCE_METADATA_HOST is already set - if so, don't override it
	for _, e := range existingEnv {
		envMap, ok := e.(map[string]interface{})
		if !ok {
			continue
		}
		name, _, _ := unstructured.NestedString(envMap, "name")
		if name == "GCE_METADATA_HOST" {
			// Already set, preserve existing value
			return nil
		}
	}

	// Add new env var
	newEnvVar := map[string]interface{}{
		"name":  "GCE_METADATA_HOST",
		"value": metadataHost,
	}
	existingEnv = append(existingEnv, newEnvVar)

	if err := unstructured.SetNestedSlice(container, existingEnv, "env"); err != nil {
		return fmt.Errorf("failed to set env vars: %w", err)
	}

	return nil
}
