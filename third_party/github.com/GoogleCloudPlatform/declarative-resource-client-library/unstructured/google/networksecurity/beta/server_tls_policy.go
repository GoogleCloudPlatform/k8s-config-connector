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
package networksecurity

import (
	"context"
	"fmt"
	"github.com/GoogleCloudPlatform/declarative-resource-client-library/dcl"
	"github.com/GoogleCloudPlatform/declarative-resource-client-library/services/google/iam"
	dclService "github.com/GoogleCloudPlatform/declarative-resource-client-library/services/google/networksecurity/beta"
	"github.com/GoogleCloudPlatform/declarative-resource-client-library/unstructured"
	iamUnstruct "github.com/GoogleCloudPlatform/declarative-resource-client-library/unstructured/google/iam"
)

type ServerTlsPolicy struct{}

func ServerTlsPolicyToUnstructured(r *dclService.ServerTlsPolicy) *unstructured.Resource {
	u := &unstructured.Resource{
		STV: unstructured.ServiceTypeVersion{
			Service: "networksecurity",
			Version: "beta",
			Type:    "ServerTlsPolicy",
		},
		Object: make(map[string]interface{}),
	}
	if r.AllowOpen != nil {
		u.Object["allowOpen"] = *r.AllowOpen
	}
	if r.CreateTime != nil {
		u.Object["createTime"] = *r.CreateTime
	}
	if r.Description != nil {
		u.Object["description"] = *r.Description
	}
	if r.Labels != nil {
		rLabels := make(map[string]interface{})
		for k, v := range r.Labels {
			rLabels[k] = v
		}
		u.Object["labels"] = rLabels
	}
	if r.Location != nil {
		u.Object["location"] = *r.Location
	}
	if r.MtlsPolicy != nil && r.MtlsPolicy != dclService.EmptyServerTlsPolicyMtlsPolicy {
		rMtlsPolicy := make(map[string]interface{})
		var rMtlsPolicyClientValidationCa []interface{}
		for _, rMtlsPolicyClientValidationCaVal := range r.MtlsPolicy.ClientValidationCa {
			rMtlsPolicyClientValidationCaObject := make(map[string]interface{})
			if rMtlsPolicyClientValidationCaVal.CertificateProviderInstance != nil && rMtlsPolicyClientValidationCaVal.CertificateProviderInstance != dclService.EmptyServerTlsPolicyMtlsPolicyClientValidationCaCertificateProviderInstance {
				rMtlsPolicyClientValidationCaValCertificateProviderInstance := make(map[string]interface{})
				if rMtlsPolicyClientValidationCaVal.CertificateProviderInstance.PluginInstance != nil {
					rMtlsPolicyClientValidationCaValCertificateProviderInstance["pluginInstance"] = *rMtlsPolicyClientValidationCaVal.CertificateProviderInstance.PluginInstance
				}
				rMtlsPolicyClientValidationCaObject["certificateProviderInstance"] = rMtlsPolicyClientValidationCaValCertificateProviderInstance
			}
			if rMtlsPolicyClientValidationCaVal.GrpcEndpoint != nil && rMtlsPolicyClientValidationCaVal.GrpcEndpoint != dclService.EmptyServerTlsPolicyMtlsPolicyClientValidationCaGrpcEndpoint {
				rMtlsPolicyClientValidationCaValGrpcEndpoint := make(map[string]interface{})
				if rMtlsPolicyClientValidationCaVal.GrpcEndpoint.TargetUri != nil {
					rMtlsPolicyClientValidationCaValGrpcEndpoint["targetUri"] = *rMtlsPolicyClientValidationCaVal.GrpcEndpoint.TargetUri
				}
				rMtlsPolicyClientValidationCaObject["grpcEndpoint"] = rMtlsPolicyClientValidationCaValGrpcEndpoint
			}
			rMtlsPolicyClientValidationCa = append(rMtlsPolicyClientValidationCa, rMtlsPolicyClientValidationCaObject)
		}
		rMtlsPolicy["clientValidationCa"] = rMtlsPolicyClientValidationCa
		u.Object["mtlsPolicy"] = rMtlsPolicy
	}
	if r.Name != nil {
		u.Object["name"] = *r.Name
	}
	if r.Project != nil {
		u.Object["project"] = *r.Project
	}
	if r.ServerCertificate != nil && r.ServerCertificate != dclService.EmptyServerTlsPolicyServerCertificate {
		rServerCertificate := make(map[string]interface{})
		if r.ServerCertificate.CertificateProviderInstance != nil && r.ServerCertificate.CertificateProviderInstance != dclService.EmptyServerTlsPolicyServerCertificateCertificateProviderInstance {
			rServerCertificateCertificateProviderInstance := make(map[string]interface{})
			if r.ServerCertificate.CertificateProviderInstance.PluginInstance != nil {
				rServerCertificateCertificateProviderInstance["pluginInstance"] = *r.ServerCertificate.CertificateProviderInstance.PluginInstance
			}
			rServerCertificate["certificateProviderInstance"] = rServerCertificateCertificateProviderInstance
		}
		if r.ServerCertificate.GrpcEndpoint != nil && r.ServerCertificate.GrpcEndpoint != dclService.EmptyServerTlsPolicyServerCertificateGrpcEndpoint {
			rServerCertificateGrpcEndpoint := make(map[string]interface{})
			if r.ServerCertificate.GrpcEndpoint.TargetUri != nil {
				rServerCertificateGrpcEndpoint["targetUri"] = *r.ServerCertificate.GrpcEndpoint.TargetUri
			}
			rServerCertificate["grpcEndpoint"] = rServerCertificateGrpcEndpoint
		}
		u.Object["serverCertificate"] = rServerCertificate
	}
	if r.UpdateTime != nil {
		u.Object["updateTime"] = *r.UpdateTime
	}
	return u
}

func UnstructuredToServerTlsPolicy(u *unstructured.Resource) (*dclService.ServerTlsPolicy, error) {
	r := &dclService.ServerTlsPolicy{}
	if _, ok := u.Object["allowOpen"]; ok {
		if b, ok := u.Object["allowOpen"].(bool); ok {
			r.AllowOpen = dcl.Bool(b)
		} else {
			return nil, fmt.Errorf("r.AllowOpen: expected bool")
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
	if _, ok := u.Object["location"]; ok {
		if s, ok := u.Object["location"].(string); ok {
			r.Location = dcl.String(s)
		} else {
			return nil, fmt.Errorf("r.Location: expected string")
		}
	}
	if _, ok := u.Object["mtlsPolicy"]; ok {
		if rMtlsPolicy, ok := u.Object["mtlsPolicy"].(map[string]interface{}); ok {
			r.MtlsPolicy = &dclService.ServerTlsPolicyMtlsPolicy{}
			if _, ok := rMtlsPolicy["clientValidationCa"]; ok {
				if s, ok := rMtlsPolicy["clientValidationCa"].([]interface{}); ok {
					for _, o := range s {
						if objval, ok := o.(map[string]interface{}); ok {
							var rMtlsPolicyClientValidationCa dclService.ServerTlsPolicyMtlsPolicyClientValidationCa
							if _, ok := objval["certificateProviderInstance"]; ok {
								if rMtlsPolicyClientValidationCaCertificateProviderInstance, ok := objval["certificateProviderInstance"].(map[string]interface{}); ok {
									rMtlsPolicyClientValidationCa.CertificateProviderInstance = &dclService.ServerTlsPolicyMtlsPolicyClientValidationCaCertificateProviderInstance{}
									if _, ok := rMtlsPolicyClientValidationCaCertificateProviderInstance["pluginInstance"]; ok {
										if s, ok := rMtlsPolicyClientValidationCaCertificateProviderInstance["pluginInstance"].(string); ok {
											rMtlsPolicyClientValidationCa.CertificateProviderInstance.PluginInstance = dcl.String(s)
										} else {
											return nil, fmt.Errorf("rMtlsPolicyClientValidationCa.CertificateProviderInstance.PluginInstance: expected string")
										}
									}
								} else {
									return nil, fmt.Errorf("rMtlsPolicyClientValidationCa.CertificateProviderInstance: expected map[string]interface{}")
								}
							}
							if _, ok := objval["grpcEndpoint"]; ok {
								if rMtlsPolicyClientValidationCaGrpcEndpoint, ok := objval["grpcEndpoint"].(map[string]interface{}); ok {
									rMtlsPolicyClientValidationCa.GrpcEndpoint = &dclService.ServerTlsPolicyMtlsPolicyClientValidationCaGrpcEndpoint{}
									if _, ok := rMtlsPolicyClientValidationCaGrpcEndpoint["targetUri"]; ok {
										if s, ok := rMtlsPolicyClientValidationCaGrpcEndpoint["targetUri"].(string); ok {
											rMtlsPolicyClientValidationCa.GrpcEndpoint.TargetUri = dcl.String(s)
										} else {
											return nil, fmt.Errorf("rMtlsPolicyClientValidationCa.GrpcEndpoint.TargetUri: expected string")
										}
									}
								} else {
									return nil, fmt.Errorf("rMtlsPolicyClientValidationCa.GrpcEndpoint: expected map[string]interface{}")
								}
							}
							r.MtlsPolicy.ClientValidationCa = append(r.MtlsPolicy.ClientValidationCa, rMtlsPolicyClientValidationCa)
						}
					}
				} else {
					return nil, fmt.Errorf("r.MtlsPolicy.ClientValidationCa: expected []interface{}")
				}
			}
		} else {
			return nil, fmt.Errorf("r.MtlsPolicy: expected map[string]interface{}")
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
	if _, ok := u.Object["serverCertificate"]; ok {
		if rServerCertificate, ok := u.Object["serverCertificate"].(map[string]interface{}); ok {
			r.ServerCertificate = &dclService.ServerTlsPolicyServerCertificate{}
			if _, ok := rServerCertificate["certificateProviderInstance"]; ok {
				if rServerCertificateCertificateProviderInstance, ok := rServerCertificate["certificateProviderInstance"].(map[string]interface{}); ok {
					r.ServerCertificate.CertificateProviderInstance = &dclService.ServerTlsPolicyServerCertificateCertificateProviderInstance{}
					if _, ok := rServerCertificateCertificateProviderInstance["pluginInstance"]; ok {
						if s, ok := rServerCertificateCertificateProviderInstance["pluginInstance"].(string); ok {
							r.ServerCertificate.CertificateProviderInstance.PluginInstance = dcl.String(s)
						} else {
							return nil, fmt.Errorf("r.ServerCertificate.CertificateProviderInstance.PluginInstance: expected string")
						}
					}
				} else {
					return nil, fmt.Errorf("r.ServerCertificate.CertificateProviderInstance: expected map[string]interface{}")
				}
			}
			if _, ok := rServerCertificate["grpcEndpoint"]; ok {
				if rServerCertificateGrpcEndpoint, ok := rServerCertificate["grpcEndpoint"].(map[string]interface{}); ok {
					r.ServerCertificate.GrpcEndpoint = &dclService.ServerTlsPolicyServerCertificateGrpcEndpoint{}
					if _, ok := rServerCertificateGrpcEndpoint["targetUri"]; ok {
						if s, ok := rServerCertificateGrpcEndpoint["targetUri"].(string); ok {
							r.ServerCertificate.GrpcEndpoint.TargetUri = dcl.String(s)
						} else {
							return nil, fmt.Errorf("r.ServerCertificate.GrpcEndpoint.TargetUri: expected string")
						}
					}
				} else {
					return nil, fmt.Errorf("r.ServerCertificate.GrpcEndpoint: expected map[string]interface{}")
				}
			}
		} else {
			return nil, fmt.Errorf("r.ServerCertificate: expected map[string]interface{}")
		}
	}
	if _, ok := u.Object["updateTime"]; ok {
		if s, ok := u.Object["updateTime"].(string); ok {
			r.UpdateTime = dcl.String(s)
		} else {
			return nil, fmt.Errorf("r.UpdateTime: expected string")
		}
	}
	return r, nil
}

func GetServerTlsPolicy(ctx context.Context, config *dcl.Config, u *unstructured.Resource) (*unstructured.Resource, error) {
	c := dclService.NewClient(config)
	r, err := UnstructuredToServerTlsPolicy(u)
	if err != nil {
		return nil, err
	}
	r, err = c.GetServerTlsPolicy(ctx, r)
	if err != nil {
		return nil, err
	}
	return ServerTlsPolicyToUnstructured(r), nil
}

func ListServerTlsPolicy(ctx context.Context, config *dcl.Config, project string, location string) ([]*unstructured.Resource, error) {
	c := dclService.NewClient(config)
	l, err := c.ListServerTlsPolicy(ctx, project, location)
	if err != nil {
		return nil, err
	}
	var resources []*unstructured.Resource
	for {
		for _, r := range l.Items {
			resources = append(resources, ServerTlsPolicyToUnstructured(r))
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

func ApplyServerTlsPolicy(ctx context.Context, config *dcl.Config, u *unstructured.Resource, opts ...dcl.ApplyOption) (*unstructured.Resource, error) {
	c := dclService.NewClient(config)
	r, err := UnstructuredToServerTlsPolicy(u)
	if err != nil {
		return nil, err
	}
	if ush := unstructured.FetchStateHint(opts); ush != nil {
		sh, err := UnstructuredToServerTlsPolicy(ush)
		if err != nil {
			return nil, err
		}
		opts = append(opts, dcl.WithStateHint(sh))
	}
	r, err = c.ApplyServerTlsPolicy(ctx, r, opts...)
	if err != nil {
		return nil, err
	}
	return ServerTlsPolicyToUnstructured(r), nil
}

func ServerTlsPolicyHasDiff(ctx context.Context, config *dcl.Config, u *unstructured.Resource, opts ...dcl.ApplyOption) (bool, error) {
	c := dclService.NewClient(config)
	r, err := UnstructuredToServerTlsPolicy(u)
	if err != nil {
		return false, err
	}
	if ush := unstructured.FetchStateHint(opts); ush != nil {
		sh, err := UnstructuredToServerTlsPolicy(ush)
		if err != nil {
			return false, err
		}
		opts = append(opts, dcl.WithStateHint(sh))
	}
	opts = append(opts, dcl.WithLifecycleParam(dcl.BlockDestruction), dcl.WithLifecycleParam(dcl.BlockCreation), dcl.WithLifecycleParam(dcl.BlockModification))
	_, err = c.ApplyServerTlsPolicy(ctx, r, opts...)
	if err != nil {
		if _, ok := err.(dcl.ApplyInfeasibleError); ok {
			return true, nil
		}
		return false, err
	}
	return false, nil
}

func DeleteServerTlsPolicy(ctx context.Context, config *dcl.Config, u *unstructured.Resource) error {
	c := dclService.NewClient(config)
	r, err := UnstructuredToServerTlsPolicy(u)
	if err != nil {
		return err
	}
	return c.DeleteServerTlsPolicy(ctx, r)
}

func ServerTlsPolicyID(u *unstructured.Resource) (string, error) {
	r, err := UnstructuredToServerTlsPolicy(u)
	if err != nil {
		return "", err
	}
	return r.ID()
}

func (r *ServerTlsPolicy) STV() unstructured.ServiceTypeVersion {
	return unstructured.ServiceTypeVersion{
		"networksecurity",
		"ServerTlsPolicy",
		"beta",
	}
}

func SetPolicyServerTlsPolicy(ctx context.Context, config *dcl.Config, u *unstructured.Resource, p *unstructured.Resource) (*unstructured.Resource, error) {
	r, err := UnstructuredToServerTlsPolicy(u)
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

func SetPolicyWithEtagServerTlsPolicy(ctx context.Context, config *dcl.Config, u *unstructured.Resource, p *unstructured.Resource) (*unstructured.Resource, error) {
	r, err := UnstructuredToServerTlsPolicy(u)
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

func GetPolicyServerTlsPolicy(ctx context.Context, config *dcl.Config, u *unstructured.Resource) (*unstructured.Resource, error) {
	r, err := UnstructuredToServerTlsPolicy(u)
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

func SetPolicyMemberServerTlsPolicy(ctx context.Context, config *dcl.Config, u *unstructured.Resource, m *unstructured.Resource) (*unstructured.Resource, error) {
	r, err := UnstructuredToServerTlsPolicy(u)
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

func GetPolicyMemberServerTlsPolicy(ctx context.Context, config *dcl.Config, u *unstructured.Resource, role, member string) (*unstructured.Resource, error) {
	r, err := UnstructuredToServerTlsPolicy(u)
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

func DeletePolicyMemberServerTlsPolicy(ctx context.Context, config *dcl.Config, u *unstructured.Resource, m *unstructured.Resource) error {
	r, err := UnstructuredToServerTlsPolicy(u)
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

func (r *ServerTlsPolicy) SetPolicyMember(ctx context.Context, config *dcl.Config, resource *unstructured.Resource, member *unstructured.Resource) (*unstructured.Resource, error) {
	return SetPolicyMemberServerTlsPolicy(ctx, config, resource, member)
}

func (r *ServerTlsPolicy) GetPolicyMember(ctx context.Context, config *dcl.Config, resource *unstructured.Resource, role, member string) (*unstructured.Resource, error) {
	return GetPolicyMemberServerTlsPolicy(ctx, config, resource, role, member)
}

func (r *ServerTlsPolicy) DeletePolicyMember(ctx context.Context, config *dcl.Config, resource *unstructured.Resource, member *unstructured.Resource) error {
	return DeletePolicyMemberServerTlsPolicy(ctx, config, resource, member)
}

func (r *ServerTlsPolicy) SetPolicy(ctx context.Context, config *dcl.Config, resource *unstructured.Resource, policy *unstructured.Resource) (*unstructured.Resource, error) {
	return SetPolicyServerTlsPolicy(ctx, config, resource, policy)
}

func (r *ServerTlsPolicy) SetPolicyWithEtag(ctx context.Context, config *dcl.Config, resource *unstructured.Resource, policy *unstructured.Resource) (*unstructured.Resource, error) {
	return SetPolicyWithEtagServerTlsPolicy(ctx, config, resource, policy)
}

func (r *ServerTlsPolicy) GetPolicy(ctx context.Context, config *dcl.Config, resource *unstructured.Resource) (*unstructured.Resource, error) {
	return GetPolicyServerTlsPolicy(ctx, config, resource)
}

func (r *ServerTlsPolicy) Get(ctx context.Context, config *dcl.Config, resource *unstructured.Resource) (*unstructured.Resource, error) {
	return GetServerTlsPolicy(ctx, config, resource)
}

func (r *ServerTlsPolicy) Apply(ctx context.Context, config *dcl.Config, resource *unstructured.Resource, opts ...dcl.ApplyOption) (*unstructured.Resource, error) {
	return ApplyServerTlsPolicy(ctx, config, resource, opts...)
}

func (r *ServerTlsPolicy) HasDiff(ctx context.Context, config *dcl.Config, resource *unstructured.Resource, opts ...dcl.ApplyOption) (bool, error) {
	return ServerTlsPolicyHasDiff(ctx, config, resource, opts...)
}

func (r *ServerTlsPolicy) Delete(ctx context.Context, config *dcl.Config, resource *unstructured.Resource) error {
	return DeleteServerTlsPolicy(ctx, config, resource)
}

func (r *ServerTlsPolicy) ID(resource *unstructured.Resource) (string, error) {
	return ServerTlsPolicyID(resource)
}

func init() {
	unstructured.Register(&ServerTlsPolicy{})
}
