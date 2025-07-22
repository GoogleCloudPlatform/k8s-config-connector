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

package v1beta1

import (
	"sigs.k8s.io/controller-runtime/pkg/conversion"
)

// ConvertTo converts this LoggingLink to the Hub version.
func (src *LoggingLink) ConvertTo(dstRaw conversion.Hub) error {
	dst := dstRaw.(*LoggingLink)
	dst.ObjectMeta = src.ObjectMeta
	dst.Spec.ResourceID = src.Spec.ResourceID
	dst.Spec.Description = src.Spec.Description
	if src.Spec.LoggingLogBucketRef != nil {
		dst.Spec.LoggingLogBucketRef = &LoggingLogBucketRef{
			Name:      src.Spec.LoggingLogBucketRef.Name,
			Namespace: src.Spec.LoggingLogBucketRef.Namespace,
			External:  src.Spec.LoggingLogBucketRef.External,
		}
	}
	dst.Status.Conditions = src.Status.Conditions
	dst.Status.ObservedGeneration = src.Status.ObservedGeneration
	dst.Status.ExternalRef = src.Status.ExternalRef
	if src.Status.ObservedState != nil {
		dst.Status.ObservedState = &LoggingLinkObservedState{
			CreateTime:     src.Status.ObservedState.CreateTime,
			LifecycleState: src.Status.ObservedState.LifecycleState,
		}
		if src.Status.ObservedState.BigQueryDataset != nil {
			dst.Status.ObservedState.BigQueryDataset = &BigQueryDatasetObservedState{
				DatasetId: src.Status.ObservedState.BigQueryDataset.DatasetId,
			}
		}
	}
	return nil
}

// ConvertFrom converts from the Hub version to this version.
func (dst *LoggingLink) ConvertFrom(srcRaw conversion.Hub) error {
	src := srcRaw.(*LoggingLink)
	dst.ObjectMeta = src.ObjectMeta
	dst.Spec.ResourceID = src.Spec.ResourceID
	dst.Spec.Description = src.Spec.Description
	if src.Spec.LoggingLogBucketRef != nil {
		dst.Spec.LoggingLogBucketRef = &LoggingLogBucketRef{
			Name:      src.Spec.LoggingLogBucketRef.Name,
			Namespace: src.Spec.LoggingLogBucketRef.Namespace,
			External:  src.Spec.LoggingLogBucketRef.External,
		}
	}
	dst.Status.Conditions = src.Status.Conditions
	dst.Status.ObservedGeneration = src.Status.ObservedGeneration
	dst.Status.ExternalRef = src.Status.ExternalRef
	if src.Status.ObservedState != nil {
		dst.Status.ObservedState = &LoggingLinkObservedState{
			CreateTime:     src.Status.ObservedState.CreateTime,
			LifecycleState: src.Status.ObservedState.LifecycleState,
		}
		if src.Status.ObservedState.BigQueryDataset != nil {
			dst.Status.ObservedState.BigQueryDataset = &BigQueryDatasetObservedState{
				DatasetId: src.Status.ObservedState.BigQueryDataset.DatasetId,
			}
		}
	}
	return nil
}
