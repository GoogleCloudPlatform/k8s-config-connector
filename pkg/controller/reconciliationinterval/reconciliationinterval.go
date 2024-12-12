// Copyright 2023 Google LLC
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

package reconciliationinterval

import (
	"fmt"
	"strconv"
	"time"

	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/apis/iam/v1beta1"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/dcl/metadata"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/k8s"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/servicemapping/servicemappingloader"

	"k8s.io/apimachinery/pkg/runtime/schema"
)

// MeanReconcileReenqueuePeriod returns the mean reconciliation interval for a resource based
// on configured reconcile interval in TF servicemapping, DCL metadata, IAM resource config.
func MeanReconcileReenqueuePeriod(gvk schema.GroupVersionKind,
	smLoader *servicemappingloader.ServiceMappingLoader,
	serviceMetadataLoader metadata.ServiceMetadataLoader) time.Duration {
	// Check if the resource has configured reconcile interval in service mapping.
	if smLoader != nil {
		rcs, err := smLoader.GetResourceConfigs(gvk)
		if err == nil && len(rcs) > 0 {
			// One GVK can map to multiple ResourceConfigs, however these ResournceConfigs will share the same reconcile interval.
			if rcs[0].ReconciliationIntervalInSeconds != nil {
				return time.Duration(*rcs[0].ReconciliationIntervalInSeconds) * time.Second
			}
		}
	}
	// Check if the resource has configured reconcile interval in service metadata.
	if serviceMetadataLoader != nil {
		resourceMetadata, found := serviceMetadataLoader.GetResourceWithGVK(gvk)
		if found && resourceMetadata.ReconciliationIntervalInSeconds != nil {
			return time.Duration(*resourceMetadata.ReconciliationIntervalInSeconds) * time.Second
		}
	}
	// Check if the resource belongs to IAMPolicy/IAMPartialPolicy/IAMPolicyMember/IAMAuditConfig.
	switch gvk.Kind {
	case "IAMPolicy":
		return v1beta1.IAMPolicyReconcileInterval
	case "IAMPartialPolicy":
		return v1beta1.IAMPartialPolicyReconcileInterval
	case "IAMPolicyMember":
		return v1beta1.IAMPolicyMemberReconcileInterval
	case "IAMAuditConfig":
		return v1beta1.IAMAuditConfigReconcileInterval
	}
	// If no GVK specific reconcile interval configured, return default value.
	return k8s.MeanReconcileReenqueuePeriod
}

func MeanReconcileReenqueuePeriodFromAnnotation(val string) (time.Duration, error) {
	reconcileIntervalInSeconds, err := strconv.ParseInt(val, 10, 32)
	if err != nil {
		return 0, fmt.Errorf("error converting the annotation %s's value %s to int32", k8s.ReconcileIntervalInSecondsAnnotation, val)
	}
	if reconcileIntervalInSeconds < 0 {
		return 0, fmt.Errorf("a negative val %d is set in the annotation %s", reconcileIntervalInSeconds, k8s.ReconcileIntervalInSecondsAnnotation)
	}
	return time.Duration(reconcileIntervalInSeconds) * time.Second, nil
}
