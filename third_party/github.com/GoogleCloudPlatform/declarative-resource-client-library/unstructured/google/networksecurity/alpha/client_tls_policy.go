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
	dclService "github.com/GoogleCloudPlatform/declarative-resource-client-library/services/google/networksecurity/alpha"
	"github.com/GoogleCloudPlatform/declarative-resource-client-library/unstructured"
	iamUnstruct "github.com/GoogleCloudPlatform/declarative-resource-client-library/unstructured/google/iam"
)

type ClientTlsPolicy struct{}

func ClientTlsPolicyToUnstructured(r *dclService.ClientTlsPolicy) *unstructured.Resource {
	u := &unstructured.Resource{
		STV: unstructured.ServiceTypeVersion{
			Service: "networksecurity",
			Version: "alpha",
			Type:    "ClientTlsPolicy",
		},
		Object: make(map[string]interface{}),
	}
	if r.ClientCertificate != nil && r.ClientCertificate != dclService.EmptyClientTlsPolicyClientCertificate {
		rClientCertificate := make(map[string]interface{})
		if r.ClientCertificate.CertificateProviderInstance != nil && r.ClientCertificate.CertificateProviderInstance != dclService.EmptyClientTlsPolicyClientCertificateCertificateProviderInstance {
			rClientCertificateCertificateProviderInstance := make(map[string]interface{})
			if r.ClientCertificate.CertificateProviderInstance.PluginInstance != nil {
				rClientCertificateCertificateProviderInstance["pluginInstance"] = *r.ClientCertificate.CertificateProviderInstance.PluginInstance
			}
			rClientCertificate["certificateProviderInstance"] = rClientCertificateCertificateProviderInstance
		}
		if r.ClientCertificate.GrpcEndpoint != nil && r.ClientCertificate.GrpcEndpoint != dclService.EmptyClientTlsPolicyClientCertificateGrpcEndpoint {
			rClientCertificateGrpcEndpoint := make(map[string]interface{})
			if r.ClientCertificate.GrpcEndpoint.TargetUri != nil {
				rClientCertificateGrpcEndpoint["targetUri"] = *r.ClientCertificate.GrpcEndpoint.TargetUri
			}
			rClientCertificate["grpcEndpoint"] = rClientCertificateGrpcEndpoint
		}
		if r.ClientCertificate.LocalFilepath != nil && r.ClientCertificate.LocalFilepath != dclService.EmptyClientTlsPolicyClientCertificateLocalFilepath {
			rClientCertificateLocalFilepath := make(map[string]interface{})
			if r.ClientCertificate.LocalFilepath.CertificatePath != nil {
				rClientCertificateLocalFilepath["certificatePath"] = *r.ClientCertificate.LocalFilepath.CertificatePath
			}
			if r.ClientCertificate.LocalFilepath.PrivateKeyPath != nil {
				rClientCertificateLocalFilepath["privateKeyPath"] = *r.ClientCertificate.LocalFilepath.PrivateKeyPath
			}
			rClientCertificate["localFilepath"] = rClientCertificateLocalFilepath
		}
		u.Object["clientCertificate"] = rClientCertificate
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
	if r.Name != nil {
		u.Object["name"] = *r.Name
	}
	if r.Project != nil {
		u.Object["project"] = *r.Project
	}
	var rServerValidationCa []interface{}
	for _, rServerValidationCaVal := range r.ServerValidationCa {
		rServerValidationCaObject := make(map[string]interface{})
		if rServerValidationCaVal.CaCertPath != nil {
			rServerValidationCaObject["caCertPath"] = *rServerValidationCaVal.CaCertPath
		}
		if rServerValidationCaVal.CertificateProviderInstance != nil && rServerValidationCaVal.CertificateProviderInstance != dclService.EmptyClientTlsPolicyServerValidationCaCertificateProviderInstance {
			rServerValidationCaValCertificateProviderInstance := make(map[string]interface{})
			if rServerValidationCaVal.CertificateProviderInstance.PluginInstance != nil {
				rServerValidationCaValCertificateProviderInstance["pluginInstance"] = *rServerValidationCaVal.CertificateProviderInstance.PluginInstance
			}
			rServerValidationCaObject["certificateProviderInstance"] = rServerValidationCaValCertificateProviderInstance
		}
		if rServerValidationCaVal.GrpcEndpoint != nil && rServerValidationCaVal.GrpcEndpoint != dclService.EmptyClientTlsPolicyServerValidationCaGrpcEndpoint {
			rServerValidationCaValGrpcEndpoint := make(map[string]interface{})
			if rServerValidationCaVal.GrpcEndpoint.TargetUri != nil {
				rServerValidationCaValGrpcEndpoint["targetUri"] = *rServerValidationCaVal.GrpcEndpoint.TargetUri
			}
			rServerValidationCaObject["grpcEndpoint"] = rServerValidationCaValGrpcEndpoint
		}
		rServerValidationCa = append(rServerValidationCa, rServerValidationCaObject)
	}
	u.Object["serverValidationCa"] = rServerValidationCa
	if r.Sni != nil {
		u.Object["sni"] = *r.Sni
	}
	if r.UpdateTime != nil {
		u.Object["updateTime"] = *r.UpdateTime
	}
	return u
}

func UnstructuredToClientTlsPolicy(u *unstructured.Resource) (*dclService.ClientTlsPolicy, error) {
	r := &dclService.ClientTlsPolicy{}
	if _, ok := u.Object["clientCertificate"]; ok {
		if rClientCertificate, ok := u.Object["clientCertificate"].(map[string]interface{}); ok {
			r.ClientCertificate = &dclService.ClientTlsPolicyClientCertificate{}
			if _, ok := rClientCertificate["certificateProviderInstance"]; ok {
				if rClientCertificateCertificateProviderInstance, ok := rClientCertificate["certificateProviderInstance"].(map[string]interface{}); ok {
					r.ClientCertificate.CertificateProviderInstance = &dclService.ClientTlsPolicyClientCertificateCertificateProviderInstance{}
					if _, ok := rClientCertificateCertificateProviderInstance["pluginInstance"]; ok {
						if s, ok := rClientCertificateCertificateProviderInstance["pluginInstance"].(string); ok {
							r.ClientCertificate.CertificateProviderInstance.PluginInstance = dcl.String(s)
						} else {
							return nil, fmt.Errorf("r.ClientCertificate.CertificateProviderInstance.PluginInstance: expected string")
						}
					}
				} else {
					return nil, fmt.Errorf("r.ClientCertificate.CertificateProviderInstance: expected map[string]interface{}")
				}
			}
			if _, ok := rClientCertificate["grpcEndpoint"]; ok {
				if rClientCertificateGrpcEndpoint, ok := rClientCertificate["grpcEndpoint"].(map[string]interface{}); ok {
					r.ClientCertificate.GrpcEndpoint = &dclService.ClientTlsPolicyClientCertificateGrpcEndpoint{}
					if _, ok := rClientCertificateGrpcEndpoint["targetUri"]; ok {
						if s, ok := rClientCertificateGrpcEndpoint["targetUri"].(string); ok {
							r.ClientCertificate.GrpcEndpoint.TargetUri = dcl.String(s)
						} else {
							return nil, fmt.Errorf("r.ClientCertificate.GrpcEndpoint.TargetUri: expected string")
						}
					}
				} else {
					return nil, fmt.Errorf("r.ClientCertificate.GrpcEndpoint: expected map[string]interface{}")
				}
			}
			if _, ok := rClientCertificate["localFilepath"]; ok {
				if rClientCertificateLocalFilepath, ok := rClientCertificate["localFilepath"].(map[string]interface{}); ok {
					r.ClientCertificate.LocalFilepath = &dclService.ClientTlsPolicyClientCertificateLocalFilepath{}
					if _, ok := rClientCertificateLocalFilepath["certificatePath"]; ok {
						if s, ok := rClientCertificateLocalFilepath["certificatePath"].(string); ok {
							r.ClientCertificate.LocalFilepath.CertificatePath = dcl.String(s)
						} else {
							return nil, fmt.Errorf("r.ClientCertificate.LocalFilepath.CertificatePath: expected string")
						}
					}
					if _, ok := rClientCertificateLocalFilepath["privateKeyPath"]; ok {
						if s, ok := rClientCertificateLocalFilepath["privateKeyPath"].(string); ok {
							r.ClientCertificate.LocalFilepath.PrivateKeyPath = dcl.String(s)
						} else {
							return nil, fmt.Errorf("r.ClientCertificate.LocalFilepath.PrivateKeyPath: expected string")
						}
					}
				} else {
					return nil, fmt.Errorf("r.ClientCertificate.LocalFilepath: expected map[string]interface{}")
				}
			}
		} else {
			return nil, fmt.Errorf("r.ClientCertificate: expected map[string]interface{}")
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
	if _, ok := u.Object["serverValidationCa"]; ok {
		if s, ok := u.Object["serverValidationCa"].([]interface{}); ok {
			for _, o := range s {
				if objval, ok := o.(map[string]interface{}); ok {
					var rServerValidationCa dclService.ClientTlsPolicyServerValidationCa
					if _, ok := objval["caCertPath"]; ok {
						if s, ok := objval["caCertPath"].(string); ok {
							rServerValidationCa.CaCertPath = dcl.String(s)
						} else {
							return nil, fmt.Errorf("rServerValidationCa.CaCertPath: expected string")
						}
					}
					if _, ok := objval["certificateProviderInstance"]; ok {
						if rServerValidationCaCertificateProviderInstance, ok := objval["certificateProviderInstance"].(map[string]interface{}); ok {
							rServerValidationCa.CertificateProviderInstance = &dclService.ClientTlsPolicyServerValidationCaCertificateProviderInstance{}
							if _, ok := rServerValidationCaCertificateProviderInstance["pluginInstance"]; ok {
								if s, ok := rServerValidationCaCertificateProviderInstance["pluginInstance"].(string); ok {
									rServerValidationCa.CertificateProviderInstance.PluginInstance = dcl.String(s)
								} else {
									return nil, fmt.Errorf("rServerValidationCa.CertificateProviderInstance.PluginInstance: expected string")
								}
							}
						} else {
							return nil, fmt.Errorf("rServerValidationCa.CertificateProviderInstance: expected map[string]interface{}")
						}
					}
					if _, ok := objval["grpcEndpoint"]; ok {
						if rServerValidationCaGrpcEndpoint, ok := objval["grpcEndpoint"].(map[string]interface{}); ok {
							rServerValidationCa.GrpcEndpoint = &dclService.ClientTlsPolicyServerValidationCaGrpcEndpoint{}
							if _, ok := rServerValidationCaGrpcEndpoint["targetUri"]; ok {
								if s, ok := rServerValidationCaGrpcEndpoint["targetUri"].(string); ok {
									rServerValidationCa.GrpcEndpoint.TargetUri = dcl.String(s)
								} else {
									return nil, fmt.Errorf("rServerValidationCa.GrpcEndpoint.TargetUri: expected string")
								}
							}
						} else {
							return nil, fmt.Errorf("rServerValidationCa.GrpcEndpoint: expected map[string]interface{}")
						}
					}
					r.ServerValidationCa = append(r.ServerValidationCa, rServerValidationCa)
				}
			}
		} else {
			return nil, fmt.Errorf("r.ServerValidationCa: expected []interface{}")
		}
	}
	if _, ok := u.Object["sni"]; ok {
		if s, ok := u.Object["sni"].(string); ok {
			r.Sni = dcl.String(s)
		} else {
			return nil, fmt.Errorf("r.Sni: expected string")
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

func GetClientTlsPolicy(ctx context.Context, config *dcl.Config, u *unstructured.Resource) (*unstructured.Resource, error) {
	c := dclService.NewClient(config)
	r, err := UnstructuredToClientTlsPolicy(u)
	if err != nil {
		return nil, err
	}
	r, err = c.GetClientTlsPolicy(ctx, r)
	if err != nil {
		return nil, err
	}
	return ClientTlsPolicyToUnstructured(r), nil
}

func ListClientTlsPolicy(ctx context.Context, config *dcl.Config, project string, location string) ([]*unstructured.Resource, error) {
	c := dclService.NewClient(config)
	l, err := c.ListClientTlsPolicy(ctx, project, location)
	if err != nil {
		return nil, err
	}
	var resources []*unstructured.Resource
	for {
		for _, r := range l.Items {
			resources = append(resources, ClientTlsPolicyToUnstructured(r))
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

func ApplyClientTlsPolicy(ctx context.Context, config *dcl.Config, u *unstructured.Resource, opts ...dcl.ApplyOption) (*unstructured.Resource, error) {
	c := dclService.NewClient(config)
	r, err := UnstructuredToClientTlsPolicy(u)
	if err != nil {
		return nil, err
	}
	if ush := unstructured.FetchStateHint(opts); ush != nil {
		sh, err := UnstructuredToClientTlsPolicy(ush)
		if err != nil {
			return nil, err
		}
		opts = append(opts, dcl.WithStateHint(sh))
	}
	r, err = c.ApplyClientTlsPolicy(ctx, r, opts...)
	if err != nil {
		return nil, err
	}
	return ClientTlsPolicyToUnstructured(r), nil
}

func ClientTlsPolicyHasDiff(ctx context.Context, config *dcl.Config, u *unstructured.Resource, opts ...dcl.ApplyOption) (bool, error) {
	c := dclService.NewClient(config)
	r, err := UnstructuredToClientTlsPolicy(u)
	if err != nil {
		return false, err
	}
	if ush := unstructured.FetchStateHint(opts); ush != nil {
		sh, err := UnstructuredToClientTlsPolicy(ush)
		if err != nil {
			return false, err
		}
		opts = append(opts, dcl.WithStateHint(sh))
	}
	opts = append(opts, dcl.WithLifecycleParam(dcl.BlockDestruction), dcl.WithLifecycleParam(dcl.BlockCreation), dcl.WithLifecycleParam(dcl.BlockModification))
	_, err = c.ApplyClientTlsPolicy(ctx, r, opts...)
	if err != nil {
		if _, ok := err.(dcl.ApplyInfeasibleError); ok {
			return true, nil
		}
		return false, err
	}
	return false, nil
}

func DeleteClientTlsPolicy(ctx context.Context, config *dcl.Config, u *unstructured.Resource) error {
	c := dclService.NewClient(config)
	r, err := UnstructuredToClientTlsPolicy(u)
	if err != nil {
		return err
	}
	return c.DeleteClientTlsPolicy(ctx, r)
}

func ClientTlsPolicyID(u *unstructured.Resource) (string, error) {
	r, err := UnstructuredToClientTlsPolicy(u)
	if err != nil {
		return "", err
	}
	return r.ID()
}

func (r *ClientTlsPolicy) STV() unstructured.ServiceTypeVersion {
	return unstructured.ServiceTypeVersion{
		"networksecurity",
		"ClientTlsPolicy",
		"alpha",
	}
}

func SetPolicyClientTlsPolicy(ctx context.Context, config *dcl.Config, u *unstructured.Resource, p *unstructured.Resource) (*unstructured.Resource, error) {
	r, err := UnstructuredToClientTlsPolicy(u)
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

func SetPolicyWithEtagClientTlsPolicy(ctx context.Context, config *dcl.Config, u *unstructured.Resource, p *unstructured.Resource) (*unstructured.Resource, error) {
	r, err := UnstructuredToClientTlsPolicy(u)
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

func GetPolicyClientTlsPolicy(ctx context.Context, config *dcl.Config, u *unstructured.Resource) (*unstructured.Resource, error) {
	r, err := UnstructuredToClientTlsPolicy(u)
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

func SetPolicyMemberClientTlsPolicy(ctx context.Context, config *dcl.Config, u *unstructured.Resource, m *unstructured.Resource) (*unstructured.Resource, error) {
	r, err := UnstructuredToClientTlsPolicy(u)
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

func GetPolicyMemberClientTlsPolicy(ctx context.Context, config *dcl.Config, u *unstructured.Resource, role, member string) (*unstructured.Resource, error) {
	r, err := UnstructuredToClientTlsPolicy(u)
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

func DeletePolicyMemberClientTlsPolicy(ctx context.Context, config *dcl.Config, u *unstructured.Resource, m *unstructured.Resource) error {
	r, err := UnstructuredToClientTlsPolicy(u)
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

func (r *ClientTlsPolicy) SetPolicyMember(ctx context.Context, config *dcl.Config, resource *unstructured.Resource, member *unstructured.Resource) (*unstructured.Resource, error) {
	return SetPolicyMemberClientTlsPolicy(ctx, config, resource, member)
}

func (r *ClientTlsPolicy) GetPolicyMember(ctx context.Context, config *dcl.Config, resource *unstructured.Resource, role, member string) (*unstructured.Resource, error) {
	return GetPolicyMemberClientTlsPolicy(ctx, config, resource, role, member)
}

func (r *ClientTlsPolicy) DeletePolicyMember(ctx context.Context, config *dcl.Config, resource *unstructured.Resource, member *unstructured.Resource) error {
	return DeletePolicyMemberClientTlsPolicy(ctx, config, resource, member)
}

func (r *ClientTlsPolicy) SetPolicy(ctx context.Context, config *dcl.Config, resource *unstructured.Resource, policy *unstructured.Resource) (*unstructured.Resource, error) {
	return SetPolicyClientTlsPolicy(ctx, config, resource, policy)
}

func (r *ClientTlsPolicy) SetPolicyWithEtag(ctx context.Context, config *dcl.Config, resource *unstructured.Resource, policy *unstructured.Resource) (*unstructured.Resource, error) {
	return SetPolicyWithEtagClientTlsPolicy(ctx, config, resource, policy)
}

func (r *ClientTlsPolicy) GetPolicy(ctx context.Context, config *dcl.Config, resource *unstructured.Resource) (*unstructured.Resource, error) {
	return GetPolicyClientTlsPolicy(ctx, config, resource)
}

func (r *ClientTlsPolicy) Get(ctx context.Context, config *dcl.Config, resource *unstructured.Resource) (*unstructured.Resource, error) {
	return GetClientTlsPolicy(ctx, config, resource)
}

func (r *ClientTlsPolicy) Apply(ctx context.Context, config *dcl.Config, resource *unstructured.Resource, opts ...dcl.ApplyOption) (*unstructured.Resource, error) {
	return ApplyClientTlsPolicy(ctx, config, resource, opts...)
}

func (r *ClientTlsPolicy) HasDiff(ctx context.Context, config *dcl.Config, resource *unstructured.Resource, opts ...dcl.ApplyOption) (bool, error) {
	return ClientTlsPolicyHasDiff(ctx, config, resource, opts...)
}

func (r *ClientTlsPolicy) Delete(ctx context.Context, config *dcl.Config, resource *unstructured.Resource) error {
	return DeleteClientTlsPolicy(ctx, config, resource)
}

func (r *ClientTlsPolicy) ID(resource *unstructured.Resource) (string, error) {
	return ClientTlsPolicyID(resource)
}

func init() {
	unstructured.Register(&ClientTlsPolicy{})
}
