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

package configconnector

import (
	"context"
	"fmt"

	"github.com/GoogleCloudPlatform/k8s-config-connector/operator/pkg/controllers"
	"github.com/GoogleCloudPlatform/k8s-config-connector/operator/pkg/environment"
	apiextensionsv1 "k8s.io/apiextensions-apiserver/pkg/apis/apiextensions/v1"
	"k8s.io/apimachinery/pkg/apis/meta/v1/unstructured"
	"k8s.io/apimachinery/pkg/runtime/schema"
	"sigs.k8s.io/controller-runtime/pkg/client"
	"sigs.k8s.io/kubebuilder-declarative-pattern/pkg/patterns/declarative"
	"sigs.k8s.io/kubebuilder-declarative-pattern/pkg/patterns/declarative/pkg/manifest"
)

func (r *Reconciler) transformForSovereignEnvironment() declarative.ObjectTransform {
	return func(ctx context.Context, o declarative.DeclarativeObject, m *manifest.Objects) error {
		env := environment.Detect()
		if !env.IsCustomUniverse() && len(env.SupportedServices) == 0 {
			return nil
		}
		if err := controllers.ApplyUniverseDomain(m, env.UniverseDomain); err != nil {
			return fmt.Errorf("error applying universe domain to operator workloads: %w", err)
		}
		FilterManifestCRDs(m, env)
		return nil
	}
}

// FilterManifestCRDs removes CRDs from the installation manifest whose service group
// is not supported in the detected cloud environment.
func FilterManifestCRDs(m *manifest.Objects, env environment.Config) {
	if len(env.SupportedServices) == 0 {
		return
	}
	filtered := make([]*manifest.Object, 0, len(m.Items))
	for _, obj := range m.Items {
		if obj.Kind == "CustomResourceDefinition" && isKCCCRD(obj) {
			if !env.IsCRDGroupSupported(obj.UnstructuredObject().GetName()) {
				continue
			}
		}
		filtered = append(filtered, obj)
	}
	m.Items = filtered
}

// PruneUnsupportedCRDs deletes registered KCC CRDs whose service group is not supported in
// the detected environment, provided zero active custom resource instances exist in the cluster.
func PruneUnsupportedCRDs(ctx context.Context, c client.Client, env environment.Config) ([]string, error) {
	if len(env.SupportedServices) == 0 {
		return nil, nil
	}
	crdList := &apiextensionsv1.CustomResourceDefinitionList{}
	if err := c.List(ctx, crdList); err != nil {
		return nil, fmt.Errorf("listing CRDs for sovereign pruning: %w", err)
	}

	var pruned []string
	for i := range crdList.Items {
		crd := &crdList.Items[i]
		if env.IsCRDGroupSupported(crd.Name) {
			continue
		}
		hasInstances, err := crdHasActiveInstances(ctx, c, crd)
		if err != nil || hasInstances {
			continue
		}
		if err := c.Delete(ctx, crd); err != nil {
			return pruned, fmt.Errorf("deleting unsupported CRD %s: %w", crd.Name, err)
		}
		pruned = append(pruned, crd.Name)
	}
	return pruned, nil
}

func crdHasActiveInstances(ctx context.Context, c client.Client, crd *apiextensionsv1.CustomResourceDefinition) (bool, error) {
	var servedVersion string
	for _, v := range crd.Spec.Versions {
		if v.Served {
			servedVersion = v.Name
			if v.Storage {
				break
			}
		}
	}
	if servedVersion == "" || crd.Spec.Names.ListKind == "" {
		return false, nil
	}

	list := &unstructured.UnstructuredList{}
	list.SetGroupVersionKind(schema.GroupVersionKind{
		Group:   crd.Spec.Group,
		Version: servedVersion,
		Kind:    crd.Spec.Names.ListKind,
	})
	if err := c.List(ctx, list, client.Limit(1)); err != nil {
		return false, err
	}
	return len(list.Items) > 0, nil
}
