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

package configconnector

import (
	"context"
	"fmt"

	"github.com/GoogleCloudPlatform/k8s-config-connector/operator/pkg/discovery"
	"github.com/GoogleCloudPlatform/k8s-config-connector/operator/pkg/k8s"
	"github.com/go-logr/logr"
	apiextensionsv1 "k8s.io/apiextensions-apiserver/pkg/apis/apiextensions/v1"
	apierrors "k8s.io/apimachinery/pkg/api/errors"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/apis/meta/v1/unstructured"
	"k8s.io/apimachinery/pkg/runtime/schema"
	"sigs.k8s.io/controller-runtime/pkg/client"
)

// PruneUnsupportedSovereignCRDs inspects registered KCC CRDs on the Kubernetes cluster.
// For any CRD whose API group is unsupported in the detected sovereign cloud:
// 1. It checks if any custom resources of that kind exist.
// 2. If 0 instances exist, it safely deletes the CRD from the cluster.
// 3. If instances exist, it skips deletion to prevent silent user data loss and logs a warning.
func PruneUnsupportedSovereignCRDs(ctx context.Context, kubeClient client.Client, allowedGroups map[string]bool, logger logr.Logger) error {
	var pageToken string
	for {
		crdList, nextToken, err := k8s.ListCRDs(ctx, kubeClient, pageToken)
		if err != nil {
			return fmt.Errorf("error listing cluster CRDs for sovereign pruning: %w", err)
		}

		for _, crd := range crdList {
			if discovery.IsCRDSupported(crd.Name, allowedGroups) {
				continue
			}

			// Unsupported CRD found. Check if instances exist.
			hasInstances, err := hasActiveInstances(ctx, kubeClient, &crd)
			if err != nil {
				logger.Error(err, "error checking active instances for unsupported CRD", "crd", crd.Name)
				continue
			}

			if hasInstances {
				logger.Info("skipping deletion of unsupported sovereign CRD because active resources exist", "crd", crd.Name)
				continue
			}

			logger.Info("safely pruning unsupported sovereign CRD from cluster", "crd", crd.Name)
			if err := kubeClient.Delete(ctx, &crd); err != nil && !apierrors.IsNotFound(err) {
				logger.Error(err, "failed to delete unsupported sovereign CRD", "crd", crd.Name)
			}
		}

		if nextToken == "" {
			break
		}
		pageToken = nextToken
	}

	return nil
}

func hasActiveInstances(ctx context.Context, kubeClient client.Client, crd *apiextensionsv1.CustomResourceDefinition) (bool, error) {
	group := crd.Spec.Group
	kind := crd.Spec.Names.Kind

	// Check each served version
	for _, version := range crd.Spec.Versions {
		if !version.Served {
			continue
		}

		gvk := schema.GroupVersionKind{
			Group:   group,
			Version: version.Name,
			Kind:    kind,
		}

		list := &unstructured.UnstructuredList{}
		list.SetGroupVersionKind(gvk)

		opts := &client.ListOptions{
			Limit: 1,
			Raw:   &metav1.ListOptions{},
		}

		if err := kubeClient.List(ctx, list, opts); err != nil {
			if apierrors.IsNotFound(err) {
				continue
			}
			return false, err
		}

		if len(list.Items) > 0 {
			return true, nil
		}
	}

	return false, nil
}
