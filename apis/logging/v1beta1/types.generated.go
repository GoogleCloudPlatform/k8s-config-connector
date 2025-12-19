// Copyright 2025 Google LLC
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

// +generated:types
// krm.group: logging.cnrm.cloud.google.com
// krm.version: v1beta1
// proto.service: google.logging.v2
// resource: LoggingLogMetric:LogMetric

package v1beta1

// +kcc:proto=google.api.Distribution.BucketOptions
type Distribution_BucketOptions struct {
	// The linear bucket.
	// +kcc:proto:field=google.api.Distribution.BucketOptions.linear_buckets
	LinearBuckets *Distribution_BucketOptions_Linear `json:"linearBuckets,omitempty"`

	// The exponential buckets.
	// +kcc:proto:field=google.api.Distribution.BucketOptions.exponential_buckets
	ExponentialBuckets *Distribution_BucketOptions_Exponential `json:"exponentialBuckets,omitempty"`

	// The explicit buckets.
	// +kcc:proto:field=google.api.Distribution.BucketOptions.explicit_buckets
	ExplicitBuckets *Distribution_BucketOptions_Explicit `json:"explicitBuckets,omitempty"`
}

// +kcc:proto=google.api.Distribution.BucketOptions.Explicit
type Distribution_BucketOptions_Explicit struct {
	// The values must be monotonically increasing.
	// +kcc:proto:field=google.api.Distribution.BucketOptions.Explicit.bounds
	Bounds []float64 `json:"bounds,omitempty"`
}

// +kcc:proto=google.api.Distribution.BucketOptions.Exponential
type Distribution_BucketOptions_Exponential struct {
	// Must be greater than 0.
	// +kcc:proto:field=google.api.Distribution.BucketOptions.Exponential.num_finite_buckets
	NumFiniteBuckets *int32 `json:"numFiniteBuckets,omitempty"`

	// Must be greater than 1.
	// +kcc:proto:field=google.api.Distribution.BucketOptions.Exponential.growth_factor
	GrowthFactor *float64 `json:"growthFactor,omitempty"`

	// Must be greater than 0.
	// +kcc:proto:field=google.api.Distribution.BucketOptions.Exponential.scale
	Scale *float64 `json:"scale,omitempty"`
}

// +kcc:proto=google.api.Distribution.BucketOptions.Linear
type Distribution_BucketOptions_Linear struct {
	// Must be greater than 0.
	// +kcc:proto:field=google.api.Distribution.BucketOptions.Linear.num_finite_buckets
	NumFiniteBuckets *int32 `json:"numFiniteBuckets,omitempty"`

	// Must be greater than 0.
	// +kcc:proto:field=google.api.Distribution.BucketOptions.Linear.width
	Width *float64 `json:"width,omitempty"`

	// Lower bound of the first bucket.
	// +kcc:proto:field=google.api.Distribution.BucketOptions.Linear.offset
	Offset *float64 `json:"offset,omitempty"`
}

// +kcc:proto=google.api.LabelDescriptor
type LabelDescriptor struct {
	// The label key.
	// +kcc:proto:field=google.api.LabelDescriptor.key
	Key *string `json:"key,omitempty"`

	// The type of data that can be assigned to the label.
	// +kcc:proto:field=google.api.LabelDescriptor.value_type
	ValueType *string `json:"valueType,omitempty"`

	// A human-readable description for the label.
	// +kcc:proto:field=google.api.LabelDescriptor.description
	Description *string `json:"description,omitempty"`
}
