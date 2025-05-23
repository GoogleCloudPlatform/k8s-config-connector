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

package logging

import (
	"fmt"
	"reflect"
	"sort"

	krm "github.com/GoogleCloudPlatform/k8s-config-connector/apis/logging/v1beta1"
	loggingv1beta1 "github.com/GoogleCloudPlatform/k8s-config-connector/apis/logging/v1beta1"
	refs "github.com/GoogleCloudPlatform/k8s-config-connector/apis/refs/v1beta1"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/util"
	api "google.golang.org/api/logging/v2"
	"k8s.io/apimachinery/pkg/apis/meta/v1/unstructured"
	"k8s.io/apimachinery/pkg/runtime"
)

func setStatus(u *unstructured.Unstructured, typedStatus any) error {
	status, err := runtime.DefaultUnstructuredConverter.ToUnstructured(typedStatus)
	if err != nil {
		return fmt.Errorf("error converting status to unstructured: %w", err)
	}

	// Use existing values for conditions/observedGeneration; they are managed in k8s not the GCP API
	old, _, _ := unstructured.NestedMap(u.Object, "status")
	if old != nil {
		status["conditions"] = old["conditions"]
		status["observedGeneration"] = old["observedGeneration"]
	}

	u.Object["status"] = status

	return nil
}

func getObservedGeneration(u *unstructured.Unstructured) int64 {
	v, _, _ := unstructured.NestedInt64(u.Object, "status", "observedGeneration")
	return v
}

// todo acpana: house these somewhere else

func compareMetricDescriptors(kccObj *krm.LogmetricMetricDescriptor, apiObj *api.MetricDescriptor) bool {
	return reflect.DeepEqual(kccObj, convertAPItoKRM_MetricDescriptor(apiObj))
}

func validateImmutableFieldsUpdated(kccObj *krm.LogmetricMetricDescriptor, apiObj *api.MetricDescriptor) error {
	actualMetricDescriptor := convertAPItoKRM_MetricDescriptor(apiObj)

	modified := []string{}
	if kccObj == nil && apiObj == nil {
		return nil
	}
	if kccObj == nil || apiObj == nil {
		return fmt.Errorf("cannot make changes to immutable field: metricDescriptor")
	}
	if !reflect.DeepEqual(kccObj.MetricKind, actualMetricDescriptor.MetricKind) {
		modified = append(modified, "metricDescriptor.MetricKind")
	}
	if !reflect.DeepEqual(kccObj.ValueType, actualMetricDescriptor.ValueType) {
		modified = append(modified, "metricDescriptor.ValueType")
	}
	if len(modified) != 0 {
		return fmt.Errorf("cannot make changes to immutable field(s): %v", modified)
	}
	return nil
}

func convertAPItoKRM_LoggingLogMetric(projectID string, in *api.LogMetric) (*unstructured.Unstructured, error) {
	if in == nil {
		return nil, fmt.Errorf("api logMetric is nil")
	}

	lm := &krm.LoggingLogMetric{}
	lm.SetGroupVersionKind(krm.LoggingLogMetricGVK)
	lm.SetName(in.Name)
	// lm.SetNamespace(in.Namespace) // todo acpana figure out namespace setting

	lm.Spec.Description = &in.Description
	lm.Spec.Disabled = &in.Disabled
	lm.Spec.Filter = in.Filter
	lm.Spec.MetricDescriptor = convertAPItoKRM_MetricDescriptor(in.MetricDescriptor)
	lm.Spec.LabelExtractors = in.LabelExtractors
	lm.Spec.BucketOptions = convertAPItoKRM_BucketOptions(in.BucketOptions)
	lm.Spec.ValueExtractor = &in.ValueExtractor
	if in.BucketName != "" {
		lm.Spec.LoggingLogBucketRef = &loggingv1beta1.LoggingLogBucketRef{
			External: in.BucketName,
		}
	}

	lm.Spec.ProjectRef = refs.ProjectRef{
		External: projectID,
	}

	u := &unstructured.Unstructured{}
	if err := util.Marshal(lm, u); err != nil {
		return nil, fmt.Errorf("error marshing logMetric to unstructured %w", err)
	}

	return u, nil
}

func convertAPItoKRM_BucketOptions(in *api.BucketOptions) *krm.LogmetricBucketOptions {
	if in == nil {
		return nil
	}

	options := &krm.LogmetricBucketOptions{}

	if in.ExplicitBuckets != nil {
		options.ExplicitBuckets = &krm.LogmetricExplicitBuckets{
			Bounds: in.ExplicitBuckets.Bounds,
		}
	}
	if in.ExponentialBuckets != nil {
		options.ExponentialBuckets = &krm.LogmetricExponentialBuckets{
			GrowthFactor:     &in.ExponentialBuckets.GrowthFactor,
			NumFiniteBuckets: &in.ExponentialBuckets.NumFiniteBuckets,
			Scale:            &in.ExponentialBuckets.Scale,
		}
	}
	if in.LinearBuckets != nil {
		options.LinearBuckets = &krm.LogmetricLinearBuckets{
			NumFiniteBuckets: &in.LinearBuckets.NumFiniteBuckets,
			Offset:           &in.LinearBuckets.Offset,
			Width:            &in.LinearBuckets.Width,
		}
	}

	return options
}

func convertAPItoKRM_MetricDescriptorStatus(apiObj *api.MetricDescriptor) *krm.LogmetricMetricDescriptorStatus {
	if apiObj == nil {
		return nil
	}

	ret := &krm.LogmetricMetricDescriptorStatus{
		MonitoredResourceTypes: apiObj.MonitoredResourceTypes,
		Name:                   &apiObj.Name,
		Type:                   &apiObj.Type,
	}

	// for backwards compatibility we don't publish the description in the status unless it is set.
	if apiObj.Description != "" {
		ret.Description = &apiObj.Description
	}

	return ret
}

func convertAPItoKRM_MetricDescriptor(apiObj *api.MetricDescriptor) *krm.LogmetricMetricDescriptor {
	if apiObj == nil {
		return nil
	}

	return &krm.LogmetricMetricDescriptor{
		DisplayName: &apiObj.DisplayName,
		Labels:      convertAPItoKRM_LogMetricLabels(apiObj.Labels),
		LaunchStage: &apiObj.LaunchStage,
		Metadata:    convertAPItoKRM_LogMetricMetadata(apiObj.Metadata),
		MetricKind:  &apiObj.MetricKind, // immutable
		Unit:        &apiObj.Unit,
		ValueType:   &apiObj.ValueType, // immutable
	}
}

func convertAPItoKRM_LogMetricLabels(apiLabels []*api.LabelDescriptor) []krm.LogmetricLabels {
	if len(apiLabels) == 0 {
		return nil
	}
	kccLabels := make([]krm.LogmetricLabels, len(apiLabels))
	for i, apiLabel := range apiLabels {
		kccLabels[i] = krm.LogmetricLabels{
			Description: &apiLabel.Description, // immutable
			Key:         &apiLabel.Key,         // immutable
			ValueType:   &apiLabel.ValueType,   // immutable
		}

		// this is a quirk of the API where the "STRING" default value gets returned as "".
		if direct.ValueOf(kccLabels[i].ValueType) == "" {
			*kccLabels[i].ValueType = "STRING" // "" defaults to "STRING"
		}
	}

	sort.Slice(kccLabels, func(i, j int) bool {
		return *kccLabels[i].Key < *kccLabels[j].Key
	})

	return kccLabels
}

func convertAPItoKRM_LogMetricMetadata(apiMetadata *api.MetricDescriptorMetadata) *krm.LogmetricMetadata {
	if apiMetadata == nil {
		return nil
	}
	return &krm.LogmetricMetadata{
		IngestDelay:  &apiMetadata.IngestDelay,
		SamplePeriod: &apiMetadata.SamplePeriod,
	}
}

// compareBucketOptions return true if the bucket options are the same, false otherwise.
func compareBucketOptions(kccObj *krm.LogmetricBucketOptions, apiObj *api.BucketOptions) bool {
	if kccObj == nil && apiObj == nil {
		return true
	}
	if kccObj == nil || apiObj == nil {
		return false
	}

	if equal := compareExplicitBuckets(kccObj.ExplicitBuckets, apiObj.ExplicitBuckets); !equal {
		return false
	}
	if equal := compareExponentialBuckets(kccObj.ExponentialBuckets, apiObj.ExponentialBuckets); !equal {
		return false
	}
	if equal := compareLinearBuckets(kccObj.LinearBuckets, apiObj.LinearBuckets); !equal {
		return false
	}

	return true
}

func compareExplicitBuckets(kccObj *krm.LogmetricExplicitBuckets, apiObj *api.Explicit) bool {
	if kccObj == nil && apiObj == nil {
		return true
	} else if kccObj == nil || apiObj == nil {
		return false
	}

	return reflect.DeepEqual(kccObj.Bounds, apiObj.Bounds)
}

func compareExponentialBuckets(kccObj *krm.LogmetricExponentialBuckets, apiObj *api.Exponential) bool {
	if kccObj == nil && apiObj == nil {
		return true
	} else if kccObj == nil || apiObj == nil {
		return false
	}

	apiExponentialBuckets := struct {
		growthFactor     float64
		numFiniteBuckets int64
		scale            float64
	}{
		growthFactor:     apiObj.GrowthFactor,
		numFiniteBuckets: apiObj.NumFiniteBuckets,
		scale:            apiObj.Scale,
	}

	return reflect.DeepEqual(kccObj, apiExponentialBuckets)
}

func compareLinearBuckets(kccObj *krm.LogmetricLinearBuckets, apiObj *api.Linear) bool {
	if kccObj == nil && apiObj == nil {
		return true
	} else if kccObj == nil || apiObj == nil {
		return false
	}
	apiLinearBuckets := struct {
		numFiniteBuckets int64
		offset           float64
		witdh            float64
	}{
		numFiniteBuckets: apiObj.NumFiniteBuckets,
		offset:           apiObj.Offset,
		witdh:            apiObj.Width,
	}

	return reflect.DeepEqual(kccObj, apiLinearBuckets)
}

func convertKCCtoAPIForBucketOptions(kccObj *krm.LogmetricBucketOptions) *api.BucketOptions {
	if kccObj == nil {
		return nil
	}

	apiObj := &api.BucketOptions{}
	if kccObj.ExplicitBuckets != nil {
		apiObj.ExplicitBuckets = &api.Explicit{}
		apiObj.ExplicitBuckets.Bounds = kccObj.ExplicitBuckets.Bounds
	}
	if kccObj.ExponentialBuckets != nil {
		apiObj.ExponentialBuckets = &api.Exponential{}
		apiObj.ExponentialBuckets.NumFiniteBuckets = direct.ValueOf(kccObj.ExponentialBuckets.NumFiniteBuckets)
		apiObj.ExponentialBuckets.GrowthFactor = direct.ValueOf(kccObj.ExponentialBuckets.GrowthFactor)
		apiObj.ExponentialBuckets.Scale = direct.ValueOf(kccObj.ExponentialBuckets.Scale)
	}
	if kccObj.LinearBuckets != nil {
		apiObj.LinearBuckets = &api.Linear{}
		apiObj.LinearBuckets.NumFiniteBuckets = direct.ValueOf(kccObj.LinearBuckets.NumFiniteBuckets)
		apiObj.LinearBuckets.Offset = direct.ValueOf(kccObj.LinearBuckets.Offset)
		apiObj.LinearBuckets.Width = direct.ValueOf(kccObj.LinearBuckets.Width)
	}

	return apiObj
}

func convertKCCtoAPI(kccObjSpec *krm.LoggingLogMetricSpec) *api.LogMetric {
	if kccObjSpec == nil {
		return nil
	}
	logMetric := &api.LogMetric{}

	if kccObjSpec.BucketOptions != nil {
		logMetric.BucketOptions = convertKCCtoAPIForBucketOptions(kccObjSpec.BucketOptions)
	}
	logMetric.Description = direct.ValueOf(kccObjSpec.Description)
	logMetric.Disabled = direct.ValueOf(kccObjSpec.Disabled)
	logMetric.Filter = kccObjSpec.Filter
	logMetric.LabelExtractors = kccObjSpec.LabelExtractors
	if kccObjSpec.MetricDescriptor != nil {
		logMetric.MetricDescriptor = convertKCCtoAPIForMetricDescriptor(kccObjSpec.MetricDescriptor)
	}
	logMetric.ValueExtractor = direct.ValueOf(kccObjSpec.ValueExtractor)
	if kccObjSpec.LoggingLogBucketRef != nil {
		// assumes kccObjSpec.LoggingLogBucketRef.External is normalized by LogBucketRef_ConvertToExternal
		logMetric.BucketName = kccObjSpec.LoggingLogBucketRef.External
	}

	return logMetric

}

func convertKCCtoAPIForMetricDescriptor(kccObj *krm.LogmetricMetricDescriptor) *api.MetricDescriptor {
	if kccObj == nil {
		return nil
	}

	metricDescriptor := &api.MetricDescriptor{}
	if kccObj.DisplayName != nil {
		metricDescriptor.DisplayName = direct.ValueOf(kccObj.DisplayName)
	}
	if kccObj.Labels != nil {
		metricDescriptor.Labels = convertKCCtoAPIForLogMetricLabels(kccObj.Labels)
	}
	metricDescriptor.LaunchStage = direct.ValueOf(kccObj.LaunchStage)
	if kccObj.Metadata != nil {
		metricDescriptor.Metadata = convertKCCtoAPIForLogMetricMetadata(kccObj.Metadata)
	}
	// immutable
	if kccObj.MetricKind != nil {
		metricDescriptor.MetricKind = direct.ValueOf(kccObj.MetricKind)
	}
	if kccObj.Unit != nil {
		metricDescriptor.Unit = direct.ValueOf(kccObj.Unit)
	}
	// immutable
	if kccObj.ValueType != nil {
		metricDescriptor.ValueType = direct.ValueOf(kccObj.ValueType)
	}
	return metricDescriptor
}

func convertKCCtoAPIForLogMetricLabels(kccLabels []krm.LogmetricLabels) []*api.LabelDescriptor {
	if len(kccLabels) == 0 {
		return nil
	}
	apiLabels := make([]*api.LabelDescriptor, len(kccLabels))
	for i, kccLabel := range kccLabels {
		apiLabels[i] = &api.LabelDescriptor{}

		apiLabels[i].Description = direct.ValueOf(kccLabel.Description)
		apiLabels[i].Key = direct.ValueOf(kccLabel.Key)
		apiLabels[i].ValueType = direct.ValueOf(kccLabel.ValueType)
	}

	return apiLabels
}

func convertKCCtoAPIForLogMetricMetadata(kccMetadata *krm.LogmetricMetadata) *api.MetricDescriptorMetadata {
	if kccMetadata == nil {
		return nil
	}

	metadata := &api.MetricDescriptorMetadata{}
	metadata.IngestDelay = direct.ValueOf(kccMetadata.IngestDelay)
	metadata.SamplePeriod = direct.ValueOf(kccMetadata.SamplePeriod)

	return metadata
}
