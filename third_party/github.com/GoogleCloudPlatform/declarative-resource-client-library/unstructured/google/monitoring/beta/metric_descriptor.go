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
package monitoring

import (
	"context"
	"fmt"
	"github.com/GoogleCloudPlatform/declarative-resource-client-library/dcl"
	dclService "github.com/GoogleCloudPlatform/declarative-resource-client-library/services/google/monitoring/beta"
	"github.com/GoogleCloudPlatform/declarative-resource-client-library/unstructured"
)

type MetricDescriptor struct{}

func MetricDescriptorToUnstructured(r *dclService.MetricDescriptor) *unstructured.Resource {
	u := &unstructured.Resource{
		STV: unstructured.ServiceTypeVersion{
			Service: "monitoring",
			Version: "beta",
			Type:    "MetricDescriptor",
		},
		Object: make(map[string]interface{}),
	}
	if r.Description != nil {
		u.Object["description"] = *r.Description
	}
	if r.DisplayName != nil {
		u.Object["displayName"] = *r.DisplayName
	}
	var rLabels []interface{}
	for _, rLabelsVal := range r.Labels {
		rLabelsObject := make(map[string]interface{})
		if rLabelsVal.Description != nil {
			rLabelsObject["description"] = *rLabelsVal.Description
		}
		if rLabelsVal.Key != nil {
			rLabelsObject["key"] = *rLabelsVal.Key
		}
		if rLabelsVal.ValueType != nil {
			rLabelsObject["valueType"] = string(*rLabelsVal.ValueType)
		}
		rLabels = append(rLabels, rLabelsObject)
	}
	u.Object["labels"] = rLabels
	if r.LaunchStage != nil {
		u.Object["launchStage"] = string(*r.LaunchStage)
	}
	if r.Metadata != nil && r.Metadata != dclService.EmptyMetricDescriptorMetadata {
		rMetadata := make(map[string]interface{})
		if r.Metadata.IngestDelay != nil {
			rMetadata["ingestDelay"] = *r.Metadata.IngestDelay
		}
		if r.Metadata.LaunchStage != nil {
			rMetadata["launchStage"] = string(*r.Metadata.LaunchStage)
		}
		if r.Metadata.SamplePeriod != nil {
			rMetadata["samplePeriod"] = *r.Metadata.SamplePeriod
		}
		u.Object["metadata"] = rMetadata
	}
	if r.MetricKind != nil {
		u.Object["metricKind"] = string(*r.MetricKind)
	}
	var rMonitoredResourceTypes []interface{}
	for _, rMonitoredResourceTypesVal := range r.MonitoredResourceTypes {
		rMonitoredResourceTypes = append(rMonitoredResourceTypes, rMonitoredResourceTypesVal)
	}
	u.Object["monitoredResourceTypes"] = rMonitoredResourceTypes
	if r.Project != nil {
		u.Object["project"] = *r.Project
	}
	if r.SelfLink != nil {
		u.Object["selfLink"] = *r.SelfLink
	}
	if r.Type != nil {
		u.Object["type"] = *r.Type
	}
	if r.Unit != nil {
		u.Object["unit"] = *r.Unit
	}
	if r.ValueType != nil {
		u.Object["valueType"] = string(*r.ValueType)
	}
	return u
}

func UnstructuredToMetricDescriptor(u *unstructured.Resource) (*dclService.MetricDescriptor, error) {
	r := &dclService.MetricDescriptor{}
	if _, ok := u.Object["description"]; ok {
		if s, ok := u.Object["description"].(string); ok {
			r.Description = dcl.String(s)
		} else {
			return nil, fmt.Errorf("r.Description: expected string")
		}
	}
	if _, ok := u.Object["displayName"]; ok {
		if s, ok := u.Object["displayName"].(string); ok {
			r.DisplayName = dcl.String(s)
		} else {
			return nil, fmt.Errorf("r.DisplayName: expected string")
		}
	}
	if _, ok := u.Object["labels"]; ok {
		if s, ok := u.Object["labels"].([]interface{}); ok {
			for _, o := range s {
				if objval, ok := o.(map[string]interface{}); ok {
					var rLabels dclService.MetricDescriptorLabels
					if _, ok := objval["description"]; ok {
						if s, ok := objval["description"].(string); ok {
							rLabels.Description = dcl.String(s)
						} else {
							return nil, fmt.Errorf("rLabels.Description: expected string")
						}
					}
					if _, ok := objval["key"]; ok {
						if s, ok := objval["key"].(string); ok {
							rLabels.Key = dcl.String(s)
						} else {
							return nil, fmt.Errorf("rLabels.Key: expected string")
						}
					}
					if _, ok := objval["valueType"]; ok {
						if s, ok := objval["valueType"].(string); ok {
							rLabels.ValueType = dclService.MetricDescriptorLabelsValueTypeEnumRef(s)
						} else {
							return nil, fmt.Errorf("rLabels.ValueType: expected string")
						}
					}
					r.Labels = append(r.Labels, rLabels)
				}
			}
		} else {
			return nil, fmt.Errorf("r.Labels: expected []interface{}")
		}
	}
	if _, ok := u.Object["launchStage"]; ok {
		if s, ok := u.Object["launchStage"].(string); ok {
			r.LaunchStage = dclService.MetricDescriptorLaunchStageEnumRef(s)
		} else {
			return nil, fmt.Errorf("r.LaunchStage: expected string")
		}
	}
	if _, ok := u.Object["metadata"]; ok {
		if rMetadata, ok := u.Object["metadata"].(map[string]interface{}); ok {
			r.Metadata = &dclService.MetricDescriptorMetadata{}
			if _, ok := rMetadata["ingestDelay"]; ok {
				if s, ok := rMetadata["ingestDelay"].(string); ok {
					r.Metadata.IngestDelay = dcl.String(s)
				} else {
					return nil, fmt.Errorf("r.Metadata.IngestDelay: expected string")
				}
			}
			if _, ok := rMetadata["launchStage"]; ok {
				if s, ok := rMetadata["launchStage"].(string); ok {
					r.Metadata.LaunchStage = dclService.MetricDescriptorMetadataLaunchStageEnumRef(s)
				} else {
					return nil, fmt.Errorf("r.Metadata.LaunchStage: expected string")
				}
			}
			if _, ok := rMetadata["samplePeriod"]; ok {
				if s, ok := rMetadata["samplePeriod"].(string); ok {
					r.Metadata.SamplePeriod = dcl.String(s)
				} else {
					return nil, fmt.Errorf("r.Metadata.SamplePeriod: expected string")
				}
			}
		} else {
			return nil, fmt.Errorf("r.Metadata: expected map[string]interface{}")
		}
	}
	if _, ok := u.Object["metricKind"]; ok {
		if s, ok := u.Object["metricKind"].(string); ok {
			r.MetricKind = dclService.MetricDescriptorMetricKindEnumRef(s)
		} else {
			return nil, fmt.Errorf("r.MetricKind: expected string")
		}
	}
	if _, ok := u.Object["monitoredResourceTypes"]; ok {
		if s, ok := u.Object["monitoredResourceTypes"].([]interface{}); ok {
			for _, ss := range s {
				if strval, ok := ss.(string); ok {
					r.MonitoredResourceTypes = append(r.MonitoredResourceTypes, strval)
				}
			}
		} else {
			return nil, fmt.Errorf("r.MonitoredResourceTypes: expected []interface{}")
		}
	}
	if _, ok := u.Object["project"]; ok {
		if s, ok := u.Object["project"].(string); ok {
			r.Project = dcl.String(s)
		} else {
			return nil, fmt.Errorf("r.Project: expected string")
		}
	}
	if _, ok := u.Object["selfLink"]; ok {
		if s, ok := u.Object["selfLink"].(string); ok {
			r.SelfLink = dcl.String(s)
		} else {
			return nil, fmt.Errorf("r.SelfLink: expected string")
		}
	}
	if _, ok := u.Object["type"]; ok {
		if s, ok := u.Object["type"].(string); ok {
			r.Type = dcl.String(s)
		} else {
			return nil, fmt.Errorf("r.Type: expected string")
		}
	}
	if _, ok := u.Object["unit"]; ok {
		if s, ok := u.Object["unit"].(string); ok {
			r.Unit = dcl.String(s)
		} else {
			return nil, fmt.Errorf("r.Unit: expected string")
		}
	}
	if _, ok := u.Object["valueType"]; ok {
		if s, ok := u.Object["valueType"].(string); ok {
			r.ValueType = dclService.MetricDescriptorValueTypeEnumRef(s)
		} else {
			return nil, fmt.Errorf("r.ValueType: expected string")
		}
	}
	return r, nil
}

func GetMetricDescriptor(ctx context.Context, config *dcl.Config, u *unstructured.Resource) (*unstructured.Resource, error) {
	c := dclService.NewClient(config)
	r, err := UnstructuredToMetricDescriptor(u)
	if err != nil {
		return nil, err
	}
	r, err = c.GetMetricDescriptor(ctx, r)
	if err != nil {
		return nil, err
	}
	return MetricDescriptorToUnstructured(r), nil
}

func ListMetricDescriptor(ctx context.Context, config *dcl.Config, project string) ([]*unstructured.Resource, error) {
	c := dclService.NewClient(config)
	l, err := c.ListMetricDescriptor(ctx, project)
	if err != nil {
		return nil, err
	}
	var resources []*unstructured.Resource
	for {
		for _, r := range l.Items {
			resources = append(resources, MetricDescriptorToUnstructured(r))
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

func ApplyMetricDescriptor(ctx context.Context, config *dcl.Config, u *unstructured.Resource, opts ...dcl.ApplyOption) (*unstructured.Resource, error) {
	c := dclService.NewClient(config)
	r, err := UnstructuredToMetricDescriptor(u)
	if err != nil {
		return nil, err
	}
	if ush := unstructured.FetchStateHint(opts); ush != nil {
		sh, err := UnstructuredToMetricDescriptor(ush)
		if err != nil {
			return nil, err
		}
		opts = append(opts, dcl.WithStateHint(sh))
	}
	r, err = c.ApplyMetricDescriptor(ctx, r, opts...)
	if err != nil {
		return nil, err
	}
	return MetricDescriptorToUnstructured(r), nil
}

func MetricDescriptorHasDiff(ctx context.Context, config *dcl.Config, u *unstructured.Resource, opts ...dcl.ApplyOption) (bool, error) {
	c := dclService.NewClient(config)
	r, err := UnstructuredToMetricDescriptor(u)
	if err != nil {
		return false, err
	}
	if ush := unstructured.FetchStateHint(opts); ush != nil {
		sh, err := UnstructuredToMetricDescriptor(ush)
		if err != nil {
			return false, err
		}
		opts = append(opts, dcl.WithStateHint(sh))
	}
	opts = append(opts, dcl.WithLifecycleParam(dcl.BlockDestruction), dcl.WithLifecycleParam(dcl.BlockCreation), dcl.WithLifecycleParam(dcl.BlockModification))
	_, err = c.ApplyMetricDescriptor(ctx, r, opts...)
	if err != nil {
		if _, ok := err.(dcl.ApplyInfeasibleError); ok {
			return true, nil
		}
		return false, err
	}
	return false, nil
}

func DeleteMetricDescriptor(ctx context.Context, config *dcl.Config, u *unstructured.Resource) error {
	c := dclService.NewClient(config)
	r, err := UnstructuredToMetricDescriptor(u)
	if err != nil {
		return err
	}
	return c.DeleteMetricDescriptor(ctx, r)
}

func MetricDescriptorID(u *unstructured.Resource) (string, error) {
	r, err := UnstructuredToMetricDescriptor(u)
	if err != nil {
		return "", err
	}
	return r.ID()
}

func (r *MetricDescriptor) STV() unstructured.ServiceTypeVersion {
	return unstructured.ServiceTypeVersion{
		"monitoring",
		"MetricDescriptor",
		"beta",
	}
}

func (r *MetricDescriptor) SetPolicyMember(ctx context.Context, config *dcl.Config, resource *unstructured.Resource, member *unstructured.Resource) (*unstructured.Resource, error) {
	return nil, unstructured.ErrNoSuchMethod
}

func (r *MetricDescriptor) GetPolicyMember(ctx context.Context, config *dcl.Config, resource *unstructured.Resource, role, member string) (*unstructured.Resource, error) {
	return nil, unstructured.ErrNoSuchMethod
}

func (r *MetricDescriptor) DeletePolicyMember(ctx context.Context, config *dcl.Config, resource *unstructured.Resource, member *unstructured.Resource) error {
	return unstructured.ErrNoSuchMethod
}

func (r *MetricDescriptor) SetPolicy(ctx context.Context, config *dcl.Config, resource *unstructured.Resource, policy *unstructured.Resource) (*unstructured.Resource, error) {
	return nil, unstructured.ErrNoSuchMethod
}

func (r *MetricDescriptor) SetPolicyWithEtag(ctx context.Context, config *dcl.Config, resource *unstructured.Resource, policy *unstructured.Resource) (*unstructured.Resource, error) {
	return nil, unstructured.ErrNoSuchMethod
}

func (r *MetricDescriptor) GetPolicy(ctx context.Context, config *dcl.Config, resource *unstructured.Resource) (*unstructured.Resource, error) {
	return nil, unstructured.ErrNoSuchMethod
}

func (r *MetricDescriptor) Get(ctx context.Context, config *dcl.Config, resource *unstructured.Resource) (*unstructured.Resource, error) {
	return GetMetricDescriptor(ctx, config, resource)
}

func (r *MetricDescriptor) Apply(ctx context.Context, config *dcl.Config, resource *unstructured.Resource, opts ...dcl.ApplyOption) (*unstructured.Resource, error) {
	return ApplyMetricDescriptor(ctx, config, resource, opts...)
}

func (r *MetricDescriptor) HasDiff(ctx context.Context, config *dcl.Config, resource *unstructured.Resource, opts ...dcl.ApplyOption) (bool, error) {
	return MetricDescriptorHasDiff(ctx, config, resource, opts...)
}

func (r *MetricDescriptor) Delete(ctx context.Context, config *dcl.Config, resource *unstructured.Resource) error {
	return DeleteMetricDescriptor(ctx, config, resource)
}

func (r *MetricDescriptor) ID(resource *unstructured.Resource) (string, error) {
	return MetricDescriptorID(resource)
}

func init() {
	unstructured.Register(&MetricDescriptor{})
}
