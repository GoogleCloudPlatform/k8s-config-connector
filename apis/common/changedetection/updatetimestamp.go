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

package changedetection

import (
	"context"
	"fmt"
	"strings"
	"time"

	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/apis/k8s/v1alpha1"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct"
	"google.golang.org/protobuf/types/known/timestamppb"
	corev1 "k8s.io/api/core/v1"
	"k8s.io/apimachinery/pkg/apis/meta/v1/unstructured"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/klog/v2"
)

// objectCore holds the fields that are relevant to an change detection.
type objectCore struct {
	Status objectCoreStatus `json:"status"`
}

type objectCoreStatus struct {
	Conditions []v1alpha1.Condition `json:"conditions,omitempty"`

	// Compared to the object's generation to detect spec changes
	ObservedGeneration *int64 `json:"observedGeneration,omitempty"`
}

// ShouldReconcileBasedOnEtag checks if we should reconcile based on the GCP etag matching the KRM etag.
// If the etag in KRM status is the same as the GCP etag, we consider the GCP object not to have changed.
// We also consider the object to have changes if the KRM object generation != observedGeneration (spec changes),
// and we also reconcile again if the object is not healthy (based on status.conditions).
//
// A few problems with the approach:
// * We miss changes due to labels or annotations.
// * If there's a change in the GCP object that isn't reflected in etag, we miss that (seems unlikely)
// * Because we set spec.resourceID, we do an extra reconciliation after first creation (because we bump generation).
func ShouldReconcileBasedOnUpdatedAt(ctx context.Context, u *unstructured.Unstructured, fieldPath string, gcpUpdateTimestamp *timestamppb.Timestamp) bool {
	log := klog.FromContext(ctx)

	if gcpUpdateTimestamp == nil {
		log.V(2).Info("updateTime not set in GCP")
		return true
	}

	gcpUpdateTime := gcpUpdateTimestamp.AsTime().Format(time.RFC3339Nano)

	obj := &objectCore{}
	if err := runtime.DefaultUnstructuredConverter.FromUnstructured(u.Object, &obj); err != nil {
		log.Error(err, "error converting from unstructured")
		return true
	}

	fieldPath = strings.TrimPrefix(fieldPath, ".")
	fields := strings.Split(fieldPath, ".")

	objUpdateTime, _, err := unstructured.NestedString(u.Object, fields...)
	if err != nil {
		log.Error(fmt.Errorf("getting field %q: %w", fieldPath, err), "error getting from unstructured")
		return true
	}

	if u.GetGeneration() != direct.ValueOf(obj.Status.ObservedGeneration) {
		log.V(2).Info("generation does not match", "generation", u.GetGeneration(), "observedGeneration", direct.ValueOf(obj.Status.ObservedGeneration))
		return true
	}

	if gcpUpdateTime == "" {
		log.V(2).Info("updateTime not set in GCP")
		return true
	}

	if objUpdateTime == "" {
		log.V(2).Info("updateTime not set in KRM object")
		return true
	}

	if gcpUpdateTime != objUpdateTime {
		log.V(2).Info("object status updateTime does not match gcp updateTime", "objUpdateTime", objUpdateTime, "gcpUpdateTime", gcpUpdateTime)
		return true
	}

	if obj.Status.Conditions != nil {
		// if there was a previously failing update let's make sure we give
		// the update a chance to heal or keep marking it as failed

		ready := false
		for _, condition := range obj.Status.Conditions {
			if condition.Type == v1alpha1.ReadyConditionType {
				if condition.Status == corev1.ConditionTrue {
					ready = true
				}
			}
		}

		if !ready {
			log.V(2).Info("status.conditions indicates object is not ready yet")
			return true
		}
	}

	log.V(2).Info("object etag matches gcp etag", "objUpdateTime", objUpdateTime, "gcpUpdateTime", gcpUpdateTime)
	return false
}
