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
	"errors"
	"fmt"
	"reflect"

	krm "github.com/GoogleCloudPlatform/k8s-config-connector/apis/resources/logging/v1beta1"
	"github.com/googleapis/gax-go/v2/apierror"
	api "google.golang.org/api/logging/v2"
	"k8s.io/apimachinery/pkg/apis/meta/v1/unstructured"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/klog/v2"
)

// todo acpana: add to factor out to top level package
// todo acpana: begin
func ValueOf[T any](p *T) T {
	var v T
	if p != nil {
		v = *p
	}
	return v
}

// IsNotFound returns true if the given error is an HTTP 404.
func IsNotFound(err error) bool {
	return HasHTTPCode(err, 404)
}

// HasHTTPCode returns true if the given error is an HTTP response with the given code.
func HasHTTPCode(err error, code int) bool {
	if err == nil {
		return false
	}
	apiError := &apierror.APIError{}
	if errors.As(err, &apiError) {
		if apiError.HTTPCode() == code {
			return true
		}
	} else {
		klog.Warningf("unexpected error type %T", err)
	}
	return false
}

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

// todo acpana: end common things

// todo acpana: house these somewhere else

func compareMetricDescriptors(kccObj *krm.LogmetricMetricDescriptor, apiObj *api.MetricDescriptor) bool {
	return reflect.DeepEqual(kccObj, convertAPItoKRM_MetricDescriptor(apiObj))
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
			Description: &apiLabel.Description,
			Key:         &apiLabel.Key,
			ValueType:   &apiLabel.ValueType,
		}
	}
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
		apiObj.ExponentialBuckets.NumFiniteBuckets = ValueOf(kccObj.ExponentialBuckets.NumFiniteBuckets)
		apiObj.ExponentialBuckets.GrowthFactor = ValueOf(kccObj.ExponentialBuckets.GrowthFactor)
		apiObj.ExponentialBuckets.Scale = ValueOf(kccObj.ExponentialBuckets.Scale)
	}
	if kccObj.LinearBuckets != nil {
		apiObj.LinearBuckets = &api.Linear{}
		apiObj.LinearBuckets.NumFiniteBuckets = ValueOf(kccObj.LinearBuckets.NumFiniteBuckets)
		apiObj.LinearBuckets.Offset = ValueOf(kccObj.LinearBuckets.Offset)
		apiObj.LinearBuckets.Width = ValueOf(kccObj.LinearBuckets.Width)
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
	logMetric.Description = ValueOf(kccObjSpec.Description)
	logMetric.Disabled = ValueOf(kccObjSpec.Disabled)
	logMetric.Filter = kccObjSpec.Filter
	logMetric.LabelExtractors = kccObjSpec.LabelExtractors
	if kccObjSpec.MetricDescriptor != nil {
		logMetric.MetricDescriptor = convertKCCtoAPIForMetricDescriptor(kccObjSpec.MetricDescriptor)
	}
	logMetric.ValueExtractor = ValueOf(kccObjSpec.ValueExtractor)
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
		metricDescriptor.DisplayName = ValueOf(kccObj.DisplayName)
		// TODO: Why the same?
		metricDescriptor.Name = ValueOf(kccObj.DisplayName)
	}
	if kccObj.Labels != nil {
		metricDescriptor.Labels = convertKCCtoAPIForLogMetricLabels(kccObj.Labels)
	}
	metricDescriptor.LaunchStage = ValueOf(kccObj.LaunchStage)
	if kccObj.Metadata != nil {
		metricDescriptor.Metadata = convertKCCtoAPIForLogMetricMetadata(kccObj.Metadata)
	}
	if kccObj.MetricKind != nil {
		metricDescriptor.MetricKind = ValueOf(kccObj.MetricKind)
	}
	if kccObj.Unit != nil {
		metricDescriptor.Unit = ValueOf(kccObj.Unit)
	}
	// immutable
	if kccObj.ValueType != nil {
		metricDescriptor.ValueType = ValueOf(kccObj.ValueType)
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

		apiLabels[i].Description = ValueOf(kccLabel.Description)
		apiLabels[i].Key = ValueOf(kccLabel.Key)
		apiLabels[i].ValueType = ValueOf(kccLabel.ValueType)
	}

	return apiLabels
}

func convertKCCtoAPIForLogMetricMetadata(kccMetadata *krm.LogmetricMetadata) *api.MetricDescriptorMetadata {
	if kccMetadata == nil {
		return nil
	}

	metadata := &api.MetricDescriptorMetadata{}
	metadata.IngestDelay = ValueOf(kccMetadata.IngestDelay)
	metadata.SamplePeriod = ValueOf(kccMetadata.SamplePeriod)

	return metadata
}
