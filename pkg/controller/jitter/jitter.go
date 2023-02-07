// Copyright 2022 Google LLC
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

package jitter

import (
	"time"

	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/reconciliationinterval"
	dclmetadata "github.com/GoogleCloudPlatform/k8s-config-connector/pkg/dcl/metadata"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/k8s"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/servicemapping/servicemappingloader"

	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime/schema"
	"k8s.io/apimachinery/pkg/util/wait"
)

// GenerateWatchJitteredTimeoutPeriod returns a wait duration to reenqueue the request between
// 1/2 * MeanReconcileReenqueuePeriod and 3/2 * MeanReconcileReenqueuePeriod (not inclusive of
// upper bound). The mean duration to reenqueue is MeanReconcileReenqueuePeriod.
func GenerateWatchJitteredTimeoutPeriod() time.Duration {
	return wait.Jitter(k8s.MeanReconcileReenqueuePeriod/2, k8s.JitterFactor)
}

// GenerateJitteredReenqueuePeriod returns a wait duration to reenqueue the request based
// on configured reconcile interval in TF servicemapping, DCL metadata, IAM resource config.
// The wait duration can be overridden with the reconcile interval configured as the object's annotation.
func GenerateJitteredReenqueuePeriod(gvk schema.GroupVersionKind,
	smLoader *servicemappingloader.ServiceMappingLoader,
	serviceMetadataLoader dclmetadata.ServiceMetadataLoader, obj metav1.Object) (time.Duration, error) {
	if val, ok := k8s.GetAnnotation(k8s.ReconcileIntervalInSecondsAnnotation, obj); ok {
		reconcileInterval, err := reconciliationinterval.MeanReconcileReenqueuePeriodFromAnnotation(val)
		if err != nil {
			return 0, err
		}
		return wait.Jitter(reconcileInterval/2, k8s.JitterFactor), nil
	}
	return wait.Jitter(reconciliationinterval.MeanReconcileReenqueuePeriod(gvk, smLoader, serviceMetadataLoader)/2, k8s.JitterFactor), nil
}
