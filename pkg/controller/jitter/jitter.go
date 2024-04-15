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
	"fmt"
	"time"

	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/reconciliationinterval"
	dclmetadata "github.com/GoogleCloudPlatform/k8s-config-connector/pkg/dcl/metadata"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/k8s"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/servicemapping/servicemappingloader"

	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime/schema"
	"k8s.io/apimachinery/pkg/util/wait"
)

type Generator interface {
	// WatchJitteredTimeout returns a wait duration to reenqueue the request between
	// 1/2 * MeanReconcileReenqueuePeriod and 3/2 * MeanReconcileReenqueuePeriod (not inclusive of
	// upper bound). The mean duration to reenqueue is MeanReconcileReenqueuePeriod.
	//
	// Use WatchJitteredTimeout whenever we need to wait for a resource to be ready.
	WatchJitteredTimeout() time.Duration

	// JitteredReenqueue returns a wait duration to reenqueue the request based
	// on configured reconcile interval in TF servicemapping, DCL metadata, IAM resource config.
	// The wait duration can be overridden with the reconcile interval configured as the object's annotation.
	//
	// Use JitteredReenqueue whenever we need to reenqueue a reconciliation.
	JitteredReenqueue(gvk schema.GroupVersionKind, obj metav1.Object) (time.Duration, error)
}

type DefaultJitterGenerator struct {
	tfLoader  *servicemappingloader.ServiceMappingLoader
	dclLoader dclmetadata.ServiceMetadataLoader
}

var _ Generator = &DefaultJitterGenerator{}

func NewDefaultGenerator(tfLoader *servicemappingloader.ServiceMappingLoader, dclLoader dclmetadata.ServiceMetadataLoader) *DefaultJitterGenerator {
	return &DefaultJitterGenerator{
		tfLoader:  tfLoader,
		dclLoader: dclLoader,
	}
}

// JitteredReenqueue returns a wait duration to reenqueue the request based
// on configured reconcile interval in TF servicemapping, DCL metadata, IAM resource config.
// The wait duration can be overridden with the reconcile interval configured as the object's annotation.
//
// Use JitteredReenqueue whenever we need to reenqueue a reconciliation.
func (l *DefaultJitterGenerator) JitteredReenqueue(gvk schema.GroupVersionKind, obj metav1.Object) (time.Duration, error) {
	if val, ok := k8s.GetAnnotation(k8s.ReconcileIntervalInSecondsAnnotation, obj); ok {
		reconcileInterval, err := reconciliationinterval.MeanReconcileReenqueuePeriodFromAnnotation(val)
		if err != nil {
			return 0, err
		}
		return wait.Jitter(reconcileInterval/2, k8s.JitterFactor), nil
	}
	return wait.Jitter(reconciliationinterval.MeanReconcileReenqueuePeriod(gvk, l.tfLoader, l.dclLoader)/2, k8s.JitterFactor), nil
}

// WatchJitteredTimeout returns a wait duration to reenqueue the request between
// 1/2 * MeanReconcileReenqueuePeriod and 3/2 * MeanReconcileReenqueuePeriod (not inclusive of
// upper bound). The mean duration to reenqueue is MeanReconcileReenqueuePeriod.
//
// Use WatchJitteredTimeout whenever we need to wait for a resource to be ready.
func (l *DefaultJitterGenerator) WatchJitteredTimeout() time.Duration {
	return wait.Jitter(k8s.MeanReconcileReenqueuePeriod/2, k8s.JitterFactor)
}

// SimpleJitterGenerator does not have any service mapping knowledge.
type SimpleJitterGenerator struct {
}

// JitteredReenqueue not implemented as it relies on service mapping knowledge.
func (*SimpleJitterGenerator) JitteredReenqueue(gvk schema.GroupVersionKind, obj metav1.Object) (time.Duration, error) {
	return 0, fmt.Errorf("no service mapping knowledge")
}

// WatchJitteredTimeout implements Generator.
func (*SimpleJitterGenerator) WatchJitteredTimeout() time.Duration {
	return wait.Jitter(k8s.MeanReconcileReenqueuePeriod/2, k8s.JitterFactor)
}

var _ Generator = &SimpleJitterGenerator{}
