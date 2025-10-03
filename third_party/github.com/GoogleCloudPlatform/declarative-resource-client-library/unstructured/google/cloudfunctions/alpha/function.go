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
package cloudfunctions

import (
	"context"
	"fmt"
	"github.com/GoogleCloudPlatform/declarative-resource-client-library/dcl"
	dclService "github.com/GoogleCloudPlatform/declarative-resource-client-library/services/google/cloudfunctions/alpha"
	"github.com/GoogleCloudPlatform/declarative-resource-client-library/services/google/iam"
	"github.com/GoogleCloudPlatform/declarative-resource-client-library/unstructured"
	iamUnstruct "github.com/GoogleCloudPlatform/declarative-resource-client-library/unstructured/google/iam"
)

type Function struct{}

func FunctionToUnstructured(r *dclService.Function) *unstructured.Resource {
	u := &unstructured.Resource{
		STV: unstructured.ServiceTypeVersion{
			Service: "cloudfunctions",
			Version: "alpha",
			Type:    "Function",
		},
		Object: make(map[string]interface{}),
	}
	if r.AvailableMemoryMb != nil {
		u.Object["availableMemoryMb"] = *r.AvailableMemoryMb
	}
	if r.Description != nil {
		u.Object["description"] = *r.Description
	}
	if r.EntryPoint != nil {
		u.Object["entryPoint"] = *r.EntryPoint
	}
	if r.EnvironmentVariables != nil {
		rEnvironmentVariables := make(map[string]interface{})
		for k, v := range r.EnvironmentVariables {
			rEnvironmentVariables[k] = v
		}
		u.Object["environmentVariables"] = rEnvironmentVariables
	}
	if r.EventTrigger != nil && r.EventTrigger != dclService.EmptyFunctionEventTrigger {
		rEventTrigger := make(map[string]interface{})
		if r.EventTrigger.EventType != nil {
			rEventTrigger["eventType"] = *r.EventTrigger.EventType
		}
		if r.EventTrigger.FailurePolicy != nil {
			rEventTrigger["failurePolicy"] = *r.EventTrigger.FailurePolicy
		}
		if r.EventTrigger.Resource != nil {
			rEventTrigger["resource"] = *r.EventTrigger.Resource
		}
		if r.EventTrigger.Service != nil {
			rEventTrigger["service"] = *r.EventTrigger.Service
		}
		u.Object["eventTrigger"] = rEventTrigger
	}
	if r.HttpsTrigger != nil && r.HttpsTrigger != dclService.EmptyFunctionHttpsTrigger {
		rHttpsTrigger := make(map[string]interface{})
		if r.HttpsTrigger.SecurityLevel != nil {
			rHttpsTrigger["securityLevel"] = string(*r.HttpsTrigger.SecurityLevel)
		}
		if r.HttpsTrigger.Url != nil {
			rHttpsTrigger["url"] = *r.HttpsTrigger.Url
		}
		u.Object["httpsTrigger"] = rHttpsTrigger
	}
	if r.IngressSettings != nil {
		u.Object["ingressSettings"] = string(*r.IngressSettings)
	}
	if r.Labels != nil {
		rLabels := make(map[string]interface{})
		for k, v := range r.Labels {
			rLabels[k] = v
		}
		u.Object["labels"] = rLabels
	}
	if r.MaxInstances != nil {
		u.Object["maxInstances"] = *r.MaxInstances
	}
	if r.Name != nil {
		u.Object["name"] = *r.Name
	}
	if r.Project != nil {
		u.Object["project"] = *r.Project
	}
	if r.Region != nil {
		u.Object["region"] = *r.Region
	}
	if r.Runtime != nil {
		u.Object["runtime"] = *r.Runtime
	}
	if r.ServiceAccountEmail != nil {
		u.Object["serviceAccountEmail"] = *r.ServiceAccountEmail
	}
	if r.SourceArchiveUrl != nil {
		u.Object["sourceArchiveUrl"] = *r.SourceArchiveUrl
	}
	if r.SourceRepository != nil && r.SourceRepository != dclService.EmptyFunctionSourceRepository {
		rSourceRepository := make(map[string]interface{})
		if r.SourceRepository.DeployedUrl != nil {
			rSourceRepository["deployedUrl"] = *r.SourceRepository.DeployedUrl
		}
		if r.SourceRepository.Url != nil {
			rSourceRepository["url"] = *r.SourceRepository.Url
		}
		u.Object["sourceRepository"] = rSourceRepository
	}
	if r.Status != nil {
		u.Object["status"] = string(*r.Status)
	}
	if r.Timeout != nil {
		u.Object["timeout"] = *r.Timeout
	}
	if r.UpdateTime != nil {
		u.Object["updateTime"] = *r.UpdateTime
	}
	if r.VersionId != nil {
		u.Object["versionId"] = *r.VersionId
	}
	if r.VPCConnector != nil {
		u.Object["vpcConnector"] = *r.VPCConnector
	}
	if r.VPCConnectorEgressSettings != nil {
		u.Object["vpcConnectorEgressSettings"] = string(*r.VPCConnectorEgressSettings)
	}
	return u
}

func UnstructuredToFunction(u *unstructured.Resource) (*dclService.Function, error) {
	r := &dclService.Function{}
	if _, ok := u.Object["availableMemoryMb"]; ok {
		if i, ok := u.Object["availableMemoryMb"].(int64); ok {
			r.AvailableMemoryMb = dcl.Int64(i)
		} else {
			return nil, fmt.Errorf("r.AvailableMemoryMb: expected int64")
		}
	}
	if _, ok := u.Object["description"]; ok {
		if s, ok := u.Object["description"].(string); ok {
			r.Description = dcl.String(s)
		} else {
			return nil, fmt.Errorf("r.Description: expected string")
		}
	}
	if _, ok := u.Object["entryPoint"]; ok {
		if s, ok := u.Object["entryPoint"].(string); ok {
			r.EntryPoint = dcl.String(s)
		} else {
			return nil, fmt.Errorf("r.EntryPoint: expected string")
		}
	}
	if _, ok := u.Object["environmentVariables"]; ok {
		if rEnvironmentVariables, ok := u.Object["environmentVariables"].(map[string]interface{}); ok {
			m := make(map[string]string)
			for k, v := range rEnvironmentVariables {
				if s, ok := v.(string); ok {
					m[k] = s
				}
			}
			r.EnvironmentVariables = m
		} else {
			return nil, fmt.Errorf("r.EnvironmentVariables: expected map[string]interface{}")
		}
	}
	if _, ok := u.Object["eventTrigger"]; ok {
		if rEventTrigger, ok := u.Object["eventTrigger"].(map[string]interface{}); ok {
			r.EventTrigger = &dclService.FunctionEventTrigger{}
			if _, ok := rEventTrigger["eventType"]; ok {
				if s, ok := rEventTrigger["eventType"].(string); ok {
					r.EventTrigger.EventType = dcl.String(s)
				} else {
					return nil, fmt.Errorf("r.EventTrigger.EventType: expected string")
				}
			}
			if _, ok := rEventTrigger["failurePolicy"]; ok {
				if b, ok := rEventTrigger["failurePolicy"].(bool); ok {
					r.EventTrigger.FailurePolicy = dcl.Bool(b)
				} else {
					return nil, fmt.Errorf("r.EventTrigger.FailurePolicy: expected bool")
				}
			}
			if _, ok := rEventTrigger["resource"]; ok {
				if s, ok := rEventTrigger["resource"].(string); ok {
					r.EventTrigger.Resource = dcl.String(s)
				} else {
					return nil, fmt.Errorf("r.EventTrigger.Resource: expected string")
				}
			}
			if _, ok := rEventTrigger["service"]; ok {
				if s, ok := rEventTrigger["service"].(string); ok {
					r.EventTrigger.Service = dcl.String(s)
				} else {
					return nil, fmt.Errorf("r.EventTrigger.Service: expected string")
				}
			}
		} else {
			return nil, fmt.Errorf("r.EventTrigger: expected map[string]interface{}")
		}
	}
	if _, ok := u.Object["httpsTrigger"]; ok {
		if rHttpsTrigger, ok := u.Object["httpsTrigger"].(map[string]interface{}); ok {
			r.HttpsTrigger = &dclService.FunctionHttpsTrigger{}
			if _, ok := rHttpsTrigger["securityLevel"]; ok {
				if s, ok := rHttpsTrigger["securityLevel"].(string); ok {
					r.HttpsTrigger.SecurityLevel = dclService.FunctionHttpsTriggerSecurityLevelEnumRef(s)
				} else {
					return nil, fmt.Errorf("r.HttpsTrigger.SecurityLevel: expected string")
				}
			}
			if _, ok := rHttpsTrigger["url"]; ok {
				if s, ok := rHttpsTrigger["url"].(string); ok {
					r.HttpsTrigger.Url = dcl.String(s)
				} else {
					return nil, fmt.Errorf("r.HttpsTrigger.Url: expected string")
				}
			}
		} else {
			return nil, fmt.Errorf("r.HttpsTrigger: expected map[string]interface{}")
		}
	}
	if _, ok := u.Object["ingressSettings"]; ok {
		if s, ok := u.Object["ingressSettings"].(string); ok {
			r.IngressSettings = dclService.FunctionIngressSettingsEnumRef(s)
		} else {
			return nil, fmt.Errorf("r.IngressSettings: expected string")
		}
	}
	if _, ok := u.Object["labels"]; ok {
		if rLabels, ok := u.Object["labels"].(map[string]interface{}); ok {
			m := make(map[string]string)
			for k, v := range rLabels {
				if s, ok := v.(string); ok {
					m[k] = s
				}
			}
			r.Labels = m
		} else {
			return nil, fmt.Errorf("r.Labels: expected map[string]interface{}")
		}
	}
	if _, ok := u.Object["maxInstances"]; ok {
		if i, ok := u.Object["maxInstances"].(int64); ok {
			r.MaxInstances = dcl.Int64(i)
		} else {
			return nil, fmt.Errorf("r.MaxInstances: expected int64")
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
	if _, ok := u.Object["region"]; ok {
		if s, ok := u.Object["region"].(string); ok {
			r.Region = dcl.String(s)
		} else {
			return nil, fmt.Errorf("r.Region: expected string")
		}
	}
	if _, ok := u.Object["runtime"]; ok {
		if s, ok := u.Object["runtime"].(string); ok {
			r.Runtime = dcl.String(s)
		} else {
			return nil, fmt.Errorf("r.Runtime: expected string")
		}
	}
	if _, ok := u.Object["serviceAccountEmail"]; ok {
		if s, ok := u.Object["serviceAccountEmail"].(string); ok {
			r.ServiceAccountEmail = dcl.String(s)
		} else {
			return nil, fmt.Errorf("r.ServiceAccountEmail: expected string")
		}
	}
	if _, ok := u.Object["sourceArchiveUrl"]; ok {
		if s, ok := u.Object["sourceArchiveUrl"].(string); ok {
			r.SourceArchiveUrl = dcl.String(s)
		} else {
			return nil, fmt.Errorf("r.SourceArchiveUrl: expected string")
		}
	}
	if _, ok := u.Object["sourceRepository"]; ok {
		if rSourceRepository, ok := u.Object["sourceRepository"].(map[string]interface{}); ok {
			r.SourceRepository = &dclService.FunctionSourceRepository{}
			if _, ok := rSourceRepository["deployedUrl"]; ok {
				if s, ok := rSourceRepository["deployedUrl"].(string); ok {
					r.SourceRepository.DeployedUrl = dcl.String(s)
				} else {
					return nil, fmt.Errorf("r.SourceRepository.DeployedUrl: expected string")
				}
			}
			if _, ok := rSourceRepository["url"]; ok {
				if s, ok := rSourceRepository["url"].(string); ok {
					r.SourceRepository.Url = dcl.String(s)
				} else {
					return nil, fmt.Errorf("r.SourceRepository.Url: expected string")
				}
			}
		} else {
			return nil, fmt.Errorf("r.SourceRepository: expected map[string]interface{}")
		}
	}
	if _, ok := u.Object["status"]; ok {
		if s, ok := u.Object["status"].(string); ok {
			r.Status = dclService.FunctionStatusEnumRef(s)
		} else {
			return nil, fmt.Errorf("r.Status: expected string")
		}
	}
	if _, ok := u.Object["timeout"]; ok {
		if s, ok := u.Object["timeout"].(string); ok {
			r.Timeout = dcl.String(s)
		} else {
			return nil, fmt.Errorf("r.Timeout: expected string")
		}
	}
	if _, ok := u.Object["updateTime"]; ok {
		if s, ok := u.Object["updateTime"].(string); ok {
			r.UpdateTime = dcl.String(s)
		} else {
			return nil, fmt.Errorf("r.UpdateTime: expected string")
		}
	}
	if _, ok := u.Object["versionId"]; ok {
		if i, ok := u.Object["versionId"].(int64); ok {
			r.VersionId = dcl.Int64(i)
		} else {
			return nil, fmt.Errorf("r.VersionId: expected int64")
		}
	}
	if _, ok := u.Object["vpcConnector"]; ok {
		if s, ok := u.Object["vpcConnector"].(string); ok {
			r.VPCConnector = dcl.String(s)
		} else {
			return nil, fmt.Errorf("r.VPCConnector: expected string")
		}
	}
	if _, ok := u.Object["vpcConnectorEgressSettings"]; ok {
		if s, ok := u.Object["vpcConnectorEgressSettings"].(string); ok {
			r.VPCConnectorEgressSettings = dclService.FunctionVPCConnectorEgressSettingsEnumRef(s)
		} else {
			return nil, fmt.Errorf("r.VPCConnectorEgressSettings: expected string")
		}
	}
	return r, nil
}

func GetFunction(ctx context.Context, config *dcl.Config, u *unstructured.Resource) (*unstructured.Resource, error) {
	c := dclService.NewClient(config)
	r, err := UnstructuredToFunction(u)
	if err != nil {
		return nil, err
	}
	r, err = c.GetFunction(ctx, r)
	if err != nil {
		return nil, err
	}
	return FunctionToUnstructured(r), nil
}

func ListFunction(ctx context.Context, config *dcl.Config, project string, region string) ([]*unstructured.Resource, error) {
	c := dclService.NewClient(config)
	l, err := c.ListFunction(ctx, project, region)
	if err != nil {
		return nil, err
	}
	var resources []*unstructured.Resource
	for {
		for _, r := range l.Items {
			resources = append(resources, FunctionToUnstructured(r))
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

func ApplyFunction(ctx context.Context, config *dcl.Config, u *unstructured.Resource, opts ...dcl.ApplyOption) (*unstructured.Resource, error) {
	c := dclService.NewClient(config)
	r, err := UnstructuredToFunction(u)
	if err != nil {
		return nil, err
	}
	if ush := unstructured.FetchStateHint(opts); ush != nil {
		sh, err := UnstructuredToFunction(ush)
		if err != nil {
			return nil, err
		}
		opts = append(opts, dcl.WithStateHint(sh))
	}
	r, err = c.ApplyFunction(ctx, r, opts...)
	if err != nil {
		return nil, err
	}
	return FunctionToUnstructured(r), nil
}

func FunctionHasDiff(ctx context.Context, config *dcl.Config, u *unstructured.Resource, opts ...dcl.ApplyOption) (bool, error) {
	c := dclService.NewClient(config)
	r, err := UnstructuredToFunction(u)
	if err != nil {
		return false, err
	}
	if ush := unstructured.FetchStateHint(opts); ush != nil {
		sh, err := UnstructuredToFunction(ush)
		if err != nil {
			return false, err
		}
		opts = append(opts, dcl.WithStateHint(sh))
	}
	opts = append(opts, dcl.WithLifecycleParam(dcl.BlockDestruction), dcl.WithLifecycleParam(dcl.BlockCreation), dcl.WithLifecycleParam(dcl.BlockModification))
	_, err = c.ApplyFunction(ctx, r, opts...)
	if err != nil {
		if _, ok := err.(dcl.ApplyInfeasibleError); ok {
			return true, nil
		}
		return false, err
	}
	return false, nil
}

func DeleteFunction(ctx context.Context, config *dcl.Config, u *unstructured.Resource) error {
	c := dclService.NewClient(config)
	r, err := UnstructuredToFunction(u)
	if err != nil {
		return err
	}
	return c.DeleteFunction(ctx, r)
}

func FunctionID(u *unstructured.Resource) (string, error) {
	r, err := UnstructuredToFunction(u)
	if err != nil {
		return "", err
	}
	return r.ID()
}

func (r *Function) STV() unstructured.ServiceTypeVersion {
	return unstructured.ServiceTypeVersion{
		"cloudfunctions",
		"Function",
		"alpha",
	}
}

func SetPolicyFunction(ctx context.Context, config *dcl.Config, u *unstructured.Resource, p *unstructured.Resource) (*unstructured.Resource, error) {
	r, err := UnstructuredToFunction(u)
	if err != nil {
		return nil, err
	}
	policy, err := iamUnstruct.UnstructuredToPolicy(p)
	if err != nil {
		return nil, err
	}
	policy.Resource = r
	iamClient := iam.NewClient(config)
	newPolicy, err := iamClient.SetPolicy(ctx, policy)
	if err != nil {
		return nil, err
	}
	return iamUnstruct.PolicyToUnstructured(newPolicy), nil
}

func SetPolicyWithEtagFunction(ctx context.Context, config *dcl.Config, u *unstructured.Resource, p *unstructured.Resource) (*unstructured.Resource, error) {
	r, err := UnstructuredToFunction(u)
	if err != nil {
		return nil, err
	}
	policy, err := iamUnstruct.UnstructuredToPolicy(p)
	if err != nil {
		return nil, err
	}
	policy.Resource = r
	iamClient := iam.NewClient(config)
	newPolicy, err := iamClient.SetPolicyWithEtag(ctx, policy)
	if err != nil {
		return nil, err
	}
	return iamUnstruct.PolicyToUnstructured(newPolicy), nil
}

func GetPolicyFunction(ctx context.Context, config *dcl.Config, u *unstructured.Resource) (*unstructured.Resource, error) {
	r, err := UnstructuredToFunction(u)
	if err != nil {
		return nil, err
	}
	iamClient := iam.NewClient(config)
	policy, err := iamClient.GetPolicy(ctx, r)
	if err != nil {
		return nil, err
	}
	return iamUnstruct.PolicyToUnstructured(policy), nil
}

func SetPolicyMemberFunction(ctx context.Context, config *dcl.Config, u *unstructured.Resource, m *unstructured.Resource) (*unstructured.Resource, error) {
	r, err := UnstructuredToFunction(u)
	if err != nil {
		return nil, err
	}
	member, err := iamUnstruct.UnstructuredToMember(m)
	if err != nil {
		return nil, err
	}
	member.Resource = r
	iamClient := iam.NewClient(config)
	policy, err := iamClient.SetMember(ctx, member)
	if err != nil {
		return nil, err
	}
	return iamUnstruct.PolicyToUnstructured(policy), nil
}

func GetPolicyMemberFunction(ctx context.Context, config *dcl.Config, u *unstructured.Resource, role, member string) (*unstructured.Resource, error) {
	r, err := UnstructuredToFunction(u)
	if err != nil {
		return nil, err
	}
	iamClient := iam.NewClient(config)
	policyMember, err := iamClient.GetMember(ctx, r, role, member)
	if err != nil {
		return nil, err
	}
	return iamUnstruct.MemberToUnstructured(policyMember), nil
}

func DeletePolicyMemberFunction(ctx context.Context, config *dcl.Config, u *unstructured.Resource, m *unstructured.Resource) error {
	r, err := UnstructuredToFunction(u)
	if err != nil {
		return err
	}
	member, err := iamUnstruct.UnstructuredToMember(m)
	if err != nil {
		return err
	}
	member.Resource = r
	iamClient := iam.NewClient(config)
	if err := iamClient.DeleteMember(ctx, member); err != nil {
		return err
	}
	return nil
}

func (r *Function) SetPolicyMember(ctx context.Context, config *dcl.Config, resource *unstructured.Resource, member *unstructured.Resource) (*unstructured.Resource, error) {
	return SetPolicyMemberFunction(ctx, config, resource, member)
}

func (r *Function) GetPolicyMember(ctx context.Context, config *dcl.Config, resource *unstructured.Resource, role, member string) (*unstructured.Resource, error) {
	return GetPolicyMemberFunction(ctx, config, resource, role, member)
}

func (r *Function) DeletePolicyMember(ctx context.Context, config *dcl.Config, resource *unstructured.Resource, member *unstructured.Resource) error {
	return DeletePolicyMemberFunction(ctx, config, resource, member)
}

func (r *Function) SetPolicy(ctx context.Context, config *dcl.Config, resource *unstructured.Resource, policy *unstructured.Resource) (*unstructured.Resource, error) {
	return SetPolicyFunction(ctx, config, resource, policy)
}

func (r *Function) SetPolicyWithEtag(ctx context.Context, config *dcl.Config, resource *unstructured.Resource, policy *unstructured.Resource) (*unstructured.Resource, error) {
	return SetPolicyWithEtagFunction(ctx, config, resource, policy)
}

func (r *Function) GetPolicy(ctx context.Context, config *dcl.Config, resource *unstructured.Resource) (*unstructured.Resource, error) {
	return GetPolicyFunction(ctx, config, resource)
}

func (r *Function) Get(ctx context.Context, config *dcl.Config, resource *unstructured.Resource) (*unstructured.Resource, error) {
	return GetFunction(ctx, config, resource)
}

func (r *Function) Apply(ctx context.Context, config *dcl.Config, resource *unstructured.Resource, opts ...dcl.ApplyOption) (*unstructured.Resource, error) {
	return ApplyFunction(ctx, config, resource, opts...)
}

func (r *Function) HasDiff(ctx context.Context, config *dcl.Config, resource *unstructured.Resource, opts ...dcl.ApplyOption) (bool, error) {
	return FunctionHasDiff(ctx, config, resource, opts...)
}

func (r *Function) Delete(ctx context.Context, config *dcl.Config, resource *unstructured.Resource) error {
	return DeleteFunction(ctx, config, resource)
}

func (r *Function) ID(resource *unstructured.Resource) (string, error) {
	return FunctionID(resource)
}

func init() {
	unstructured.Register(&Function{})
}
