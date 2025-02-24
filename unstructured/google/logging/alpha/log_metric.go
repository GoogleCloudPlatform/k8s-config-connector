// Copyright 2024 Google LLC. All Rights Reserved.
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
package logging

import (
	"context"
	"fmt"
	"github.com/GoogleCloudPlatform/declarative-resource-client-library/dcl"
	dclService "github.com/GoogleCloudPlatform/declarative-resource-client-library/services/google/logging/alpha"
	"github.com/GoogleCloudPlatform/declarative-resource-client-library/unstructured"
)

type LogMetric struct{}

func LogMetricToUnstructured(r *dclService.LogMetric) *unstructured.Resource {
	u := &unstructured.Resource{
		STV: unstructured.ServiceTypeVersion{
			Service: "logging",
			Version: "alpha",
			Type:    "LogMetric",
		},
		Object: make(map[string]interface{}),
	}
	if r.BucketOptions != nil && r.BucketOptions != dclService.EmptyLogMetricBucketOptions {
		rBucketOptions := make(map[string]interface{})
		if r.BucketOptions.ExplicitBuckets != nil && r.BucketOptions.ExplicitBuckets != dclService.EmptyLogMetricBucketOptionsExplicitBuckets {
			rBucketOptionsExplicitBuckets := make(map[string]interface{})
			var rBucketOptionsExplicitBucketsBounds []interface{}
			for _, rBucketOptionsExplicitBucketsBoundsVal := range r.BucketOptions.ExplicitBuckets.Bounds {
				rBucketOptionsExplicitBucketsBounds = append(rBucketOptionsExplicitBucketsBounds, rBucketOptionsExplicitBucketsBoundsVal)
			}
			rBucketOptionsExplicitBuckets["bounds"] = rBucketOptionsExplicitBucketsBounds
			rBucketOptions["explicitBuckets"] = rBucketOptionsExplicitBuckets
		}
		if r.BucketOptions.ExponentialBuckets != nil && r.BucketOptions.ExponentialBuckets != dclService.EmptyLogMetricBucketOptionsExponentialBuckets {
			rBucketOptionsExponentialBuckets := make(map[string]interface{})
			if r.BucketOptions.ExponentialBuckets.GrowthFactor != nil {
				rBucketOptionsExponentialBuckets["growthFactor"] = *r.BucketOptions.ExponentialBuckets.GrowthFactor
			}
			if r.BucketOptions.ExponentialBuckets.NumFiniteBuckets != nil {
				rBucketOptionsExponentialBuckets["numFiniteBuckets"] = *r.BucketOptions.ExponentialBuckets.NumFiniteBuckets
			}
			if r.BucketOptions.ExponentialBuckets.Scale != nil {
				rBucketOptionsExponentialBuckets["scale"] = *r.BucketOptions.ExponentialBuckets.Scale
			}
			rBucketOptions["exponentialBuckets"] = rBucketOptionsExponentialBuckets
		}
		if r.BucketOptions.LinearBuckets != nil && r.BucketOptions.LinearBuckets != dclService.EmptyLogMetricBucketOptionsLinearBuckets {
			rBucketOptionsLinearBuckets := make(map[string]interface{})
			if r.BucketOptions.LinearBuckets.NumFiniteBuckets != nil {
				rBucketOptionsLinearBuckets["numFiniteBuckets"] = *r.BucketOptions.LinearBuckets.NumFiniteBuckets
			}
			if r.BucketOptions.LinearBuckets.Offset != nil {
				rBucketOptionsLinearBuckets["offset"] = *r.BucketOptions.LinearBuckets.Offset
			}
			if r.BucketOptions.LinearBuckets.Width != nil {
				rBucketOptionsLinearBuckets["width"] = *r.BucketOptions.LinearBuckets.Width
			}
			rBucketOptions["linearBuckets"] = rBucketOptionsLinearBuckets
		}
		u.Object["bucketOptions"] = rBucketOptions
	}
	if r.CreateTime != nil {
		u.Object["createTime"] = *r.CreateTime
	}
	if r.Description != nil {
		u.Object["description"] = *r.Description
	}
	if r.Disabled != nil {
		u.Object["disabled"] = *r.Disabled
	}
	if r.Filter != nil {
		u.Object["filter"] = *r.Filter
	}
	if r.LabelExtractors != nil {
		rLabelExtractors := make(map[string]interface{})
		for k, v := range r.LabelExtractors {
			rLabelExtractors[k] = v
		}
		u.Object["labelExtractors"] = rLabelExtractors
	}
	if r.MetricDescriptor != nil && r.MetricDescriptor != dclService.EmptyLogMetricMetricDescriptor {
		rMetricDescriptor := make(map[string]interface{})
		if r.MetricDescriptor.Description != nil {
			rMetricDescriptor["description"] = *r.MetricDescriptor.Description
		}
		if r.MetricDescriptor.DisplayName != nil {
			rMetricDescriptor["displayName"] = *r.MetricDescriptor.DisplayName
		}
		var rMetricDescriptorLabels []interface{}
		for _, rMetricDescriptorLabelsVal := range r.MetricDescriptor.Labels {
			rMetricDescriptorLabelsObject := make(map[string]interface{})
			if rMetricDescriptorLabelsVal.Description != nil {
				rMetricDescriptorLabelsObject["description"] = *rMetricDescriptorLabelsVal.Description
			}
			if rMetricDescriptorLabelsVal.Key != nil {
				rMetricDescriptorLabelsObject["key"] = *rMetricDescriptorLabelsVal.Key
			}
			if rMetricDescriptorLabelsVal.ValueType != nil {
				rMetricDescriptorLabelsObject["valueType"] = string(*rMetricDescriptorLabelsVal.ValueType)
			}
			rMetricDescriptorLabels = append(rMetricDescriptorLabels, rMetricDescriptorLabelsObject)
		}
		rMetricDescriptor["labels"] = rMetricDescriptorLabels
		if r.MetricDescriptor.LaunchStage != nil {
			rMetricDescriptor["launchStage"] = string(*r.MetricDescriptor.LaunchStage)
		}
		if r.MetricDescriptor.Metadata != nil && r.MetricDescriptor.Metadata != dclService.EmptyLogMetricMetricDescriptorMetadata {
			rMetricDescriptorMetadata := make(map[string]interface{})
			if r.MetricDescriptor.Metadata.IngestDelay != nil {
				rMetricDescriptorMetadata["ingestDelay"] = *r.MetricDescriptor.Metadata.IngestDelay
			}
			if r.MetricDescriptor.Metadata.SamplePeriod != nil {
				rMetricDescriptorMetadata["samplePeriod"] = *r.MetricDescriptor.Metadata.SamplePeriod
			}
			rMetricDescriptor["metadata"] = rMetricDescriptorMetadata
		}
		if r.MetricDescriptor.MetricKind != nil {
			rMetricDescriptor["metricKind"] = string(*r.MetricDescriptor.MetricKind)
		}
		var rMetricDescriptorMonitoredResourceTypes []interface{}
		for _, rMetricDescriptorMonitoredResourceTypesVal := range r.MetricDescriptor.MonitoredResourceTypes {
			rMetricDescriptorMonitoredResourceTypes = append(rMetricDescriptorMonitoredResourceTypes, rMetricDescriptorMonitoredResourceTypesVal)
		}
		rMetricDescriptor["monitoredResourceTypes"] = rMetricDescriptorMonitoredResourceTypes
		if r.MetricDescriptor.Name != nil {
			rMetricDescriptor["name"] = *r.MetricDescriptor.Name
		}
		if r.MetricDescriptor.Type != nil {
			rMetricDescriptor["type"] = *r.MetricDescriptor.Type
		}
		if r.MetricDescriptor.Unit != nil {
			rMetricDescriptor["unit"] = *r.MetricDescriptor.Unit
		}
		if r.MetricDescriptor.ValueType != nil {
			rMetricDescriptor["valueType"] = string(*r.MetricDescriptor.ValueType)
		}
		u.Object["metricDescriptor"] = rMetricDescriptor
	}
	if r.Name != nil {
		u.Object["name"] = *r.Name
	}
	if r.Project != nil {
		u.Object["project"] = *r.Project
	}
	if r.UpdateTime != nil {
		u.Object["updateTime"] = *r.UpdateTime
	}
	if r.ValueExtractor != nil {
		u.Object["valueExtractor"] = *r.ValueExtractor
	}
	return u
}

func UnstructuredToLogMetric(u *unstructured.Resource) (*dclService.LogMetric, error) {
	r := &dclService.LogMetric{}
	if _, ok := u.Object["bucketOptions"]; ok {
		if rBucketOptions, ok := u.Object["bucketOptions"].(map[string]interface{}); ok {
			r.BucketOptions = &dclService.LogMetricBucketOptions{}
			if _, ok := rBucketOptions["explicitBuckets"]; ok {
				if rBucketOptionsExplicitBuckets, ok := rBucketOptions["explicitBuckets"].(map[string]interface{}); ok {
					r.BucketOptions.ExplicitBuckets = &dclService.LogMetricBucketOptionsExplicitBuckets{}
					if _, ok := rBucketOptionsExplicitBuckets["bounds"]; ok {
						if s, ok := rBucketOptionsExplicitBuckets["bounds"].([]interface{}); ok {
							for _, ss := range s {
								if floatVal, ok := ss.(float64); ok {
									r.BucketOptions.ExplicitBuckets.Bounds = append(r.BucketOptions.ExplicitBuckets.Bounds, floatVal)
								}
							}
						} else {
							return nil, fmt.Errorf("r.BucketOptions.ExplicitBuckets.Bounds: expected []interface{}")
						}
					}
				} else {
					return nil, fmt.Errorf("r.BucketOptions.ExplicitBuckets: expected map[string]interface{}")
				}
			}
			if _, ok := rBucketOptions["exponentialBuckets"]; ok {
				if rBucketOptionsExponentialBuckets, ok := rBucketOptions["exponentialBuckets"].(map[string]interface{}); ok {
					r.BucketOptions.ExponentialBuckets = &dclService.LogMetricBucketOptionsExponentialBuckets{}
					if _, ok := rBucketOptionsExponentialBuckets["growthFactor"]; ok {
						if f, ok := rBucketOptionsExponentialBuckets["growthFactor"].(float64); ok {
							r.BucketOptions.ExponentialBuckets.GrowthFactor = dcl.Float64(f)
						} else {
							return nil, fmt.Errorf("r.BucketOptions.ExponentialBuckets.GrowthFactor: expected float64")
						}
					}
					if _, ok := rBucketOptionsExponentialBuckets["numFiniteBuckets"]; ok {
						if i, ok := rBucketOptionsExponentialBuckets["numFiniteBuckets"].(int64); ok {
							r.BucketOptions.ExponentialBuckets.NumFiniteBuckets = dcl.Int64(i)
						} else {
							return nil, fmt.Errorf("r.BucketOptions.ExponentialBuckets.NumFiniteBuckets: expected int64")
						}
					}
					if _, ok := rBucketOptionsExponentialBuckets["scale"]; ok {
						if f, ok := rBucketOptionsExponentialBuckets["scale"].(float64); ok {
							r.BucketOptions.ExponentialBuckets.Scale = dcl.Float64(f)
						} else {
							return nil, fmt.Errorf("r.BucketOptions.ExponentialBuckets.Scale: expected float64")
						}
					}
				} else {
					return nil, fmt.Errorf("r.BucketOptions.ExponentialBuckets: expected map[string]interface{}")
				}
			}
			if _, ok := rBucketOptions["linearBuckets"]; ok {
				if rBucketOptionsLinearBuckets, ok := rBucketOptions["linearBuckets"].(map[string]interface{}); ok {
					r.BucketOptions.LinearBuckets = &dclService.LogMetricBucketOptionsLinearBuckets{}
					if _, ok := rBucketOptionsLinearBuckets["numFiniteBuckets"]; ok {
						if i, ok := rBucketOptionsLinearBuckets["numFiniteBuckets"].(int64); ok {
							r.BucketOptions.LinearBuckets.NumFiniteBuckets = dcl.Int64(i)
						} else {
							return nil, fmt.Errorf("r.BucketOptions.LinearBuckets.NumFiniteBuckets: expected int64")
						}
					}
					if _, ok := rBucketOptionsLinearBuckets["offset"]; ok {
						if f, ok := rBucketOptionsLinearBuckets["offset"].(float64); ok {
							r.BucketOptions.LinearBuckets.Offset = dcl.Float64(f)
						} else {
							return nil, fmt.Errorf("r.BucketOptions.LinearBuckets.Offset: expected float64")
						}
					}
					if _, ok := rBucketOptionsLinearBuckets["width"]; ok {
						if f, ok := rBucketOptionsLinearBuckets["width"].(float64); ok {
							r.BucketOptions.LinearBuckets.Width = dcl.Float64(f)
						} else {
							return nil, fmt.Errorf("r.BucketOptions.LinearBuckets.Width: expected float64")
						}
					}
				} else {
					return nil, fmt.Errorf("r.BucketOptions.LinearBuckets: expected map[string]interface{}")
				}
			}
		} else {
			return nil, fmt.Errorf("r.BucketOptions: expected map[string]interface{}")
		}
	}
	if _, ok := u.Object["createTime"]; ok {
		if s, ok := u.Object["createTime"].(string); ok {
			r.CreateTime = dcl.String(s)
		} else {
			return nil, fmt.Errorf("r.CreateTime: expected string")
		}
	}
	if _, ok := u.Object["description"]; ok {
		if s, ok := u.Object["description"].(string); ok {
			r.Description = dcl.String(s)
		} else {
			return nil, fmt.Errorf("r.Description: expected string")
		}
	}
	if _, ok := u.Object["disabled"]; ok {
		if b, ok := u.Object["disabled"].(bool); ok {
			r.Disabled = dcl.Bool(b)
		} else {
			return nil, fmt.Errorf("r.Disabled: expected bool")
		}
	}
	if _, ok := u.Object["filter"]; ok {
		if s, ok := u.Object["filter"].(string); ok {
			r.Filter = dcl.String(s)
		} else {
			return nil, fmt.Errorf("r.Filter: expected string")
		}
	}
	if _, ok := u.Object["labelExtractors"]; ok {
		if rLabelExtractors, ok := u.Object["labelExtractors"].(map[string]interface{}); ok {
			m := make(map[string]string)
			for k, v := range rLabelExtractors {
				if s, ok := v.(string); ok {
					m[k] = s
				}
			}
			r.LabelExtractors = m
		} else {
			return nil, fmt.Errorf("r.LabelExtractors: expected map[string]interface{}")
		}
	}
	if _, ok := u.Object["metricDescriptor"]; ok {
		if rMetricDescriptor, ok := u.Object["metricDescriptor"].(map[string]interface{}); ok {
			r.MetricDescriptor = &dclService.LogMetricMetricDescriptor{}
			if _, ok := rMetricDescriptor["description"]; ok {
				if s, ok := rMetricDescriptor["description"].(string); ok {
					r.MetricDescriptor.Description = dcl.String(s)
				} else {
					return nil, fmt.Errorf("r.MetricDescriptor.Description: expected string")
				}
			}
			if _, ok := rMetricDescriptor["displayName"]; ok {
				if s, ok := rMetricDescriptor["displayName"].(string); ok {
					r.MetricDescriptor.DisplayName = dcl.String(s)
				} else {
					return nil, fmt.Errorf("r.MetricDescriptor.DisplayName: expected string")
				}
			}
			if _, ok := rMetricDescriptor["labels"]; ok {
				if s, ok := rMetricDescriptor["labels"].([]interface{}); ok {
					for _, o := range s {
						if objval, ok := o.(map[string]interface{}); ok {
							var rMetricDescriptorLabels dclService.LogMetricMetricDescriptorLabels
							if _, ok := objval["description"]; ok {
								if s, ok := objval["description"].(string); ok {
									rMetricDescriptorLabels.Description = dcl.String(s)
								} else {
									return nil, fmt.Errorf("rMetricDescriptorLabels.Description: expected string")
								}
							}
							if _, ok := objval["key"]; ok {
								if s, ok := objval["key"].(string); ok {
									rMetricDescriptorLabels.Key = dcl.String(s)
								} else {
									return nil, fmt.Errorf("rMetricDescriptorLabels.Key: expected string")
								}
							}
							if _, ok := objval["valueType"]; ok {
								if s, ok := objval["valueType"].(string); ok {
									rMetricDescriptorLabels.ValueType = dclService.LogMetricMetricDescriptorLabelsValueTypeEnumRef(s)
								} else {
									return nil, fmt.Errorf("rMetricDescriptorLabels.ValueType: expected string")
								}
							}
							r.MetricDescriptor.Labels = append(r.MetricDescriptor.Labels, rMetricDescriptorLabels)
						}
					}
				} else {
					return nil, fmt.Errorf("r.MetricDescriptor.Labels: expected []interface{}")
				}
			}
			if _, ok := rMetricDescriptor["launchStage"]; ok {
				if s, ok := rMetricDescriptor["launchStage"].(string); ok {
					r.MetricDescriptor.LaunchStage = dclService.LogMetricMetricDescriptorLaunchStageEnumRef(s)
				} else {
					return nil, fmt.Errorf("r.MetricDescriptor.LaunchStage: expected string")
				}
			}
			if _, ok := rMetricDescriptor["metadata"]; ok {
				if rMetricDescriptorMetadata, ok := rMetricDescriptor["metadata"].(map[string]interface{}); ok {
					r.MetricDescriptor.Metadata = &dclService.LogMetricMetricDescriptorMetadata{}
					if _, ok := rMetricDescriptorMetadata["ingestDelay"]; ok {
						if s, ok := rMetricDescriptorMetadata["ingestDelay"].(string); ok {
							r.MetricDescriptor.Metadata.IngestDelay = dcl.String(s)
						} else {
							return nil, fmt.Errorf("r.MetricDescriptor.Metadata.IngestDelay: expected string")
						}
					}
					if _, ok := rMetricDescriptorMetadata["samplePeriod"]; ok {
						if s, ok := rMetricDescriptorMetadata["samplePeriod"].(string); ok {
							r.MetricDescriptor.Metadata.SamplePeriod = dcl.String(s)
						} else {
							return nil, fmt.Errorf("r.MetricDescriptor.Metadata.SamplePeriod: expected string")
						}
					}
				} else {
					return nil, fmt.Errorf("r.MetricDescriptor.Metadata: expected map[string]interface{}")
				}
			}
			if _, ok := rMetricDescriptor["metricKind"]; ok {
				if s, ok := rMetricDescriptor["metricKind"].(string); ok {
					r.MetricDescriptor.MetricKind = dclService.LogMetricMetricDescriptorMetricKindEnumRef(s)
				} else {
					return nil, fmt.Errorf("r.MetricDescriptor.MetricKind: expected string")
				}
			}
			if _, ok := rMetricDescriptor["monitoredResourceTypes"]; ok {
				if s, ok := rMetricDescriptor["monitoredResourceTypes"].([]interface{}); ok {
					for _, ss := range s {
						if strval, ok := ss.(string); ok {
							r.MetricDescriptor.MonitoredResourceTypes = append(r.MetricDescriptor.MonitoredResourceTypes, strval)
						}
					}
				} else {
					return nil, fmt.Errorf("r.MetricDescriptor.MonitoredResourceTypes: expected []interface{}")
				}
			}
			if _, ok := rMetricDescriptor["name"]; ok {
				if s, ok := rMetricDescriptor["name"].(string); ok {
					r.MetricDescriptor.Name = dcl.String(s)
				} else {
					return nil, fmt.Errorf("r.MetricDescriptor.Name: expected string")
				}
			}
			if _, ok := rMetricDescriptor["type"]; ok {
				if s, ok := rMetricDescriptor["type"].(string); ok {
					r.MetricDescriptor.Type = dcl.String(s)
				} else {
					return nil, fmt.Errorf("r.MetricDescriptor.Type: expected string")
				}
			}
			if _, ok := rMetricDescriptor["unit"]; ok {
				if s, ok := rMetricDescriptor["unit"].(string); ok {
					r.MetricDescriptor.Unit = dcl.String(s)
				} else {
					return nil, fmt.Errorf("r.MetricDescriptor.Unit: expected string")
				}
			}
			if _, ok := rMetricDescriptor["valueType"]; ok {
				if s, ok := rMetricDescriptor["valueType"].(string); ok {
					r.MetricDescriptor.ValueType = dclService.LogMetricMetricDescriptorValueTypeEnumRef(s)
				} else {
					return nil, fmt.Errorf("r.MetricDescriptor.ValueType: expected string")
				}
			}
		} else {
			return nil, fmt.Errorf("r.MetricDescriptor: expected map[string]interface{}")
		}
	}
	if _, ok := u.Object["name"]; ok {
		if s, ok := u.Object["name"].(string); ok {
			r.Name = dcl.String(s)
		} else {
			return nil, fmt.Errorf("r.Name: expected string")
		}
	}
	if _, ok := u.Object["project"]; ok {
		if s, ok := u.Object["project"].(string); ok {
			r.Project = dcl.String(s)
		} else {
			return nil, fmt.Errorf("r.Project: expected string")
		}
	}
	if _, ok := u.Object["updateTime"]; ok {
		if s, ok := u.Object["updateTime"].(string); ok {
			r.UpdateTime = dcl.String(s)
		} else {
			return nil, fmt.Errorf("r.UpdateTime: expected string")
		}
	}
	if _, ok := u.Object["valueExtractor"]; ok {
		if s, ok := u.Object["valueExtractor"].(string); ok {
			r.ValueExtractor = dcl.String(s)
		} else {
			return nil, fmt.Errorf("r.ValueExtractor: expected string")
		}
	}
	return r, nil
}

func GetLogMetric(ctx context.Context, config *dcl.Config, u *unstructured.Resource) (*unstructured.Resource, error) {
	c := dclService.NewClient(config)
	r, err := UnstructuredToLogMetric(u)
	if err != nil {
		return nil, err
	}
	r, err = c.GetLogMetric(ctx, r)
	if err != nil {
		return nil, err
	}
	return LogMetricToUnstructured(r), nil
}

func ListLogMetric(ctx context.Context, config *dcl.Config, project string) ([]*unstructured.Resource, error) {
	c := dclService.NewClient(config)
	l, err := c.ListLogMetric(ctx, project)
	if err != nil {
		return nil, err
	}
	var resources []*unstructured.Resource
	for {
		for _, r := range l.Items {
			resources = append(resources, LogMetricToUnstructured(r))
		}
		if !l.HasNext() {
			break
		}
		if err := l.Next(ctx, c); err != nil {
			return nil, err
		}
	}
	return resources, nil
}

func ApplyLogMetric(ctx context.Context, config *dcl.Config, u *unstructured.Resource, opts ...dcl.ApplyOption) (*unstructured.Resource, error) {
	c := dclService.NewClient(config)
	r, err := UnstructuredToLogMetric(u)
	if err != nil {
		return nil, err
	}
	if ush := unstructured.FetchStateHint(opts); ush != nil {
		sh, err := UnstructuredToLogMetric(ush)
		if err != nil {
			return nil, err
		}
		opts = append(opts, dcl.WithStateHint(sh))
	}
	r, err = c.ApplyLogMetric(ctx, r, opts...)
	if err != nil {
		return nil, err
	}
	return LogMetricToUnstructured(r), nil
}

func LogMetricHasDiff(ctx context.Context, config *dcl.Config, u *unstructured.Resource, opts ...dcl.ApplyOption) (bool, error) {
	c := dclService.NewClient(config)
	r, err := UnstructuredToLogMetric(u)
	if err != nil {
		return false, err
	}
	if ush := unstructured.FetchStateHint(opts); ush != nil {
		sh, err := UnstructuredToLogMetric(ush)
		if err != nil {
			return false, err
		}
		opts = append(opts, dcl.WithStateHint(sh))
	}
	opts = append(opts, dcl.WithLifecycleParam(dcl.BlockDestruction), dcl.WithLifecycleParam(dcl.BlockCreation), dcl.WithLifecycleParam(dcl.BlockModification))
	_, err = c.ApplyLogMetric(ctx, r, opts...)
	if err != nil {
		if _, ok := err.(dcl.ApplyInfeasibleError); ok {
			return true, nil
		}
		return false, err
	}
	return false, nil
}

func DeleteLogMetric(ctx context.Context, config *dcl.Config, u *unstructured.Resource) error {
	c := dclService.NewClient(config)
	r, err := UnstructuredToLogMetric(u)
	if err != nil {
		return err
	}
	return c.DeleteLogMetric(ctx, r)
}

func LogMetricID(u *unstructured.Resource) (string, error) {
	r, err := UnstructuredToLogMetric(u)
	if err != nil {
		return "", err
	}
	return r.ID()
}

func (r *LogMetric) STV() unstructured.ServiceTypeVersion {
	return unstructured.ServiceTypeVersion{
		"logging",
		"LogMetric",
		"alpha",
	}
}

func (r *LogMetric) SetPolicyMember(ctx context.Context, config *dcl.Config, resource *unstructured.Resource, member *unstructured.Resource) (*unstructured.Resource, error) {
	return nil, unstructured.ErrNoSuchMethod
}

func (r *LogMetric) GetPolicyMember(ctx context.Context, config *dcl.Config, resource *unstructured.Resource, role, member string) (*unstructured.Resource, error) {
	return nil, unstructured.ErrNoSuchMethod
}

func (r *LogMetric) DeletePolicyMember(ctx context.Context, config *dcl.Config, resource *unstructured.Resource, member *unstructured.Resource) error {
	return unstructured.ErrNoSuchMethod
}

func (r *LogMetric) SetPolicy(ctx context.Context, config *dcl.Config, resource *unstructured.Resource, policy *unstructured.Resource) (*unstructured.Resource, error) {
	return nil, unstructured.ErrNoSuchMethod
}

func (r *LogMetric) SetPolicyWithEtag(ctx context.Context, config *dcl.Config, resource *unstructured.Resource, policy *unstructured.Resource) (*unstructured.Resource, error) {
	return nil, unstructured.ErrNoSuchMethod
}

func (r *LogMetric) GetPolicy(ctx context.Context, config *dcl.Config, resource *unstructured.Resource) (*unstructured.Resource, error) {
	return nil, unstructured.ErrNoSuchMethod
}

func (r *LogMetric) Get(ctx context.Context, config *dcl.Config, resource *unstructured.Resource) (*unstructured.Resource, error) {
	return GetLogMetric(ctx, config, resource)
}

func (r *LogMetric) Apply(ctx context.Context, config *dcl.Config, resource *unstructured.Resource, opts ...dcl.ApplyOption) (*unstructured.Resource, error) {
	return ApplyLogMetric(ctx, config, resource, opts...)
}

func (r *LogMetric) HasDiff(ctx context.Context, config *dcl.Config, resource *unstructured.Resource, opts ...dcl.ApplyOption) (bool, error) {
	return LogMetricHasDiff(ctx, config, resource, opts...)
}

func (r *LogMetric) Delete(ctx context.Context, config *dcl.Config, resource *unstructured.Resource) error {
	return DeleteLogMetric(ctx, config, resource)
}

func (r *LogMetric) ID(resource *unstructured.Resource) (string, error) {
	return LogMetricID(resource)
}

func init() {
	unstructured.Register(&LogMetric{})
}
