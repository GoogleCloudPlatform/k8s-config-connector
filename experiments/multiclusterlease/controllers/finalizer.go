// Copyright 2025 Google LLC
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package controllers

import (
	"context"
	"slices"

	v1alpha1 "github.com/GoogleCloudPlatform/k8s-config-connector/experiments/multiclusterlease/api/v1alpha1"
)

const finalizerName = "multiclusterlease.core.cnrm.cloud.google.com/finalizer"

func (r *MultiClusterLeaseReconciler) hasFinalizer(mcl *v1alpha1.MultiClusterLease) bool {
	return slices.Contains(mcl.GetFinalizers(), finalizerName)
}

func (r *MultiClusterLeaseReconciler) ensureFinalizer(ctx context.Context, mcl *v1alpha1.MultiClusterLease) error {
	if !r.hasFinalizer(mcl) {
		mcl.SetFinalizers(append(mcl.GetFinalizers(), finalizerName))
		return r.Update(ctx, mcl)
	}
	return nil
}

func (r *MultiClusterLeaseReconciler) removeFinalizer(ctx context.Context, mcl *v1alpha1.MultiClusterLease) error {
	mcl.SetFinalizers(removeString(mcl.GetFinalizers(), finalizerName))
	return r.Update(ctx, mcl)
}

func removeString(slice []string, s string) []string {
	result := make([]string, 0, len(slice))
	for _, item := range slice {
		if item != s {
			result = append(result, item)
		}
	}
	return result
}
