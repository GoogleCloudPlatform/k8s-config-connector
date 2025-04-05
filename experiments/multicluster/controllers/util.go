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

package controllers

import (
	"context"
	"fmt"

	multiclusterv1alpha1 "github.com/GoogleCloudPlatform/k8s-config-connector/experiments/multicluster/api/v1alpha1"
)

const (
	multiClusterLeaseFinalizer = "multiclusterlease.cnrm.cloud.google.com/finalizer"
)

func mapKey(mcl *multiclusterv1alpha1.MultiClusterLease) string {
	return fmt.Sprintf("%s/%s", mcl.Spec.MultiClusterUID, mcl.Spec.Identity)
}

func (r *MultiClusterLeaseReconciler) hasFinalizer(mcl *multiclusterv1alpha1.MultiClusterLease) bool {
	for _, f := range mcl.GetFinalizers() {
		if f == multiClusterLeaseFinalizer {
			return true
		}
	}
	return false
}

func (r *MultiClusterLeaseReconciler) ensureFinalizer(ctx context.Context, mcl *multiclusterv1alpha1.MultiClusterLease) error {
	for _, f := range mcl.GetFinalizers() {
		if f == multiClusterLeaseFinalizer {
			// finalizer already exists
			return nil
		}
	}
	mcl.SetFinalizers(append(mcl.GetFinalizers(), multiClusterLeaseFinalizer))
	return r.client.Update(ctx, mcl)
}

func (r *MultiClusterLeaseReconciler) removeFinalizer(ctx context.Context, mcl *multiclusterv1alpha1.MultiClusterLease) error {
	found := false
	var finalizers []string
	for _, f := range mcl.GetFinalizers() {
		if f != multiClusterLeaseFinalizer {
			finalizers = append(finalizers, f)
		} else {
			found = true
		}
	}
	if found {
		mcl.SetFinalizers(finalizers)
		return r.client.Update(ctx, mcl)
	}
	return nil
}
