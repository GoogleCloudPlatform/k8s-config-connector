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
package cloudkms

import (
	"context"
	"fmt"
	"github.com/GoogleCloudPlatform/declarative-resource-client-library/dcl"
	dclService "github.com/GoogleCloudPlatform/declarative-resource-client-library/services/google/cloudkms"
	"github.com/GoogleCloudPlatform/declarative-resource-client-library/services/google/iam"
	"github.com/GoogleCloudPlatform/declarative-resource-client-library/unstructured"
	iamUnstruct "github.com/GoogleCloudPlatform/declarative-resource-client-library/unstructured/google/iam"
)

type EkmConnection struct{}

func EkmConnectionToUnstructured(r *dclService.EkmConnection) *unstructured.Resource {
	u := &unstructured.Resource{
		STV: unstructured.ServiceTypeVersion{
			Service: "cloudkms",
			Version: "ga",
			Type:    "EkmConnection",
		},
		Object: make(map[string]interface{}),
	}
	if r.CreateTime != nil {
		u.Object["createTime"] = *r.CreateTime
	}
	if r.Etag != nil {
		u.Object["etag"] = *r.Etag
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
	var rServiceResolvers []interface{}
	for _, rServiceResolversVal := range r.ServiceResolvers {
		rServiceResolversObject := make(map[string]interface{})
		if rServiceResolversVal.EndpointFilter != nil {
			rServiceResolversObject["endpointFilter"] = *rServiceResolversVal.EndpointFilter
		}
		if rServiceResolversVal.Hostname != nil {
			rServiceResolversObject["hostname"] = *rServiceResolversVal.Hostname
		}
		var rServiceResolversValServerCertificates []interface{}
		for _, rServiceResolversValServerCertificatesVal := range rServiceResolversVal.ServerCertificates {
			rServiceResolversValServerCertificatesObject := make(map[string]interface{})
			if rServiceResolversValServerCertificatesVal.Issuer != nil {
				rServiceResolversValServerCertificatesObject["issuer"] = *rServiceResolversValServerCertificatesVal.Issuer
			}
			if rServiceResolversValServerCertificatesVal.NotAfterTime != nil {
				rServiceResolversValServerCertificatesObject["notAfterTime"] = *rServiceResolversValServerCertificatesVal.NotAfterTime
			}
			if rServiceResolversValServerCertificatesVal.NotBeforeTime != nil {
				rServiceResolversValServerCertificatesObject["notBeforeTime"] = *rServiceResolversValServerCertificatesVal.NotBeforeTime
			}
			if rServiceResolversValServerCertificatesVal.Parsed != nil {
				rServiceResolversValServerCertificatesObject["parsed"] = *rServiceResolversValServerCertificatesVal.Parsed
			}
			if rServiceResolversValServerCertificatesVal.RawDer != nil {
				rServiceResolversValServerCertificatesObject["rawDer"] = *rServiceResolversValServerCertificatesVal.RawDer
			}
			if rServiceResolversValServerCertificatesVal.SerialNumber != nil {
				rServiceResolversValServerCertificatesObject["serialNumber"] = *rServiceResolversValServerCertificatesVal.SerialNumber
			}
			if rServiceResolversValServerCertificatesVal.Sha256Fingerprint != nil {
				rServiceResolversValServerCertificatesObject["sha256Fingerprint"] = *rServiceResolversValServerCertificatesVal.Sha256Fingerprint
			}
			if rServiceResolversValServerCertificatesVal.Subject != nil {
				rServiceResolversValServerCertificatesObject["subject"] = *rServiceResolversValServerCertificatesVal.Subject
			}
			var rServiceResolversValServerCertificatesValSubjectAlternativeDnsNames []interface{}
			for _, rServiceResolversValServerCertificatesValSubjectAlternativeDnsNamesVal := range rServiceResolversValServerCertificatesVal.SubjectAlternativeDnsNames {
				rServiceResolversValServerCertificatesValSubjectAlternativeDnsNames = append(rServiceResolversValServerCertificatesValSubjectAlternativeDnsNames, rServiceResolversValServerCertificatesValSubjectAlternativeDnsNamesVal)
			}
			rServiceResolversValServerCertificatesObject["subjectAlternativeDnsNames"] = rServiceResolversValServerCertificatesValSubjectAlternativeDnsNames
			rServiceResolversValServerCertificates = append(rServiceResolversValServerCertificates, rServiceResolversValServerCertificatesObject)
		}
		rServiceResolversObject["serverCertificates"] = rServiceResolversValServerCertificates
		if rServiceResolversVal.ServiceDirectoryService != nil {
			rServiceResolversObject["serviceDirectoryService"] = *rServiceResolversVal.ServiceDirectoryService
		}
		rServiceResolvers = append(rServiceResolvers, rServiceResolversObject)
	}
	u.Object["serviceResolvers"] = rServiceResolvers
	return u
}

func UnstructuredToEkmConnection(u *unstructured.Resource) (*dclService.EkmConnection, error) {
	r := &dclService.EkmConnection{}
	if _, ok := u.Object["createTime"]; ok {
		if s, ok := u.Object["createTime"].(string); ok {
			r.CreateTime = dcl.String(s)
		} else {
			return nil, fmt.Errorf("r.CreateTime: expected string")
		}
	}
	if _, ok := u.Object["etag"]; ok {
		if s, ok := u.Object["etag"].(string); ok {
			r.Etag = dcl.String(s)
		} else {
			return nil, fmt.Errorf("r.Etag: expected string")
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
	if _, ok := u.Object["serviceResolvers"]; ok {
		if s, ok := u.Object["serviceResolvers"].([]interface{}); ok {
			for _, o := range s {
				if objval, ok := o.(map[string]interface{}); ok {
					var rServiceResolvers dclService.EkmConnectionServiceResolvers
					if _, ok := objval["endpointFilter"]; ok {
						if s, ok := objval["endpointFilter"].(string); ok {
							rServiceResolvers.EndpointFilter = dcl.String(s)
						} else {
							return nil, fmt.Errorf("rServiceResolvers.EndpointFilter: expected string")
						}
					}
					if _, ok := objval["hostname"]; ok {
						if s, ok := objval["hostname"].(string); ok {
							rServiceResolvers.Hostname = dcl.String(s)
						} else {
							return nil, fmt.Errorf("rServiceResolvers.Hostname: expected string")
						}
					}
					if _, ok := objval["serverCertificates"]; ok {
						if s, ok := objval["serverCertificates"].([]interface{}); ok {
							for _, o := range s {
								if objval, ok := o.(map[string]interface{}); ok {
									var rServiceResolversServerCertificates dclService.EkmConnectionServiceResolversServerCertificates
									if _, ok := objval["issuer"]; ok {
										if s, ok := objval["issuer"].(string); ok {
											rServiceResolversServerCertificates.Issuer = dcl.String(s)
										} else {
											return nil, fmt.Errorf("rServiceResolversServerCertificates.Issuer: expected string")
										}
									}
									if _, ok := objval["notAfterTime"]; ok {
										if s, ok := objval["notAfterTime"].(string); ok {
											rServiceResolversServerCertificates.NotAfterTime = dcl.String(s)
										} else {
											return nil, fmt.Errorf("rServiceResolversServerCertificates.NotAfterTime: expected string")
										}
									}
									if _, ok := objval["notBeforeTime"]; ok {
										if s, ok := objval["notBeforeTime"].(string); ok {
											rServiceResolversServerCertificates.NotBeforeTime = dcl.String(s)
										} else {
											return nil, fmt.Errorf("rServiceResolversServerCertificates.NotBeforeTime: expected string")
										}
									}
									if _, ok := objval["parsed"]; ok {
										if b, ok := objval["parsed"].(bool); ok {
											rServiceResolversServerCertificates.Parsed = dcl.Bool(b)
										} else {
											return nil, fmt.Errorf("rServiceResolversServerCertificates.Parsed: expected bool")
										}
									}
									if _, ok := objval["rawDer"]; ok {
										if s, ok := objval["rawDer"].(string); ok {
											rServiceResolversServerCertificates.RawDer = dcl.String(s)
										} else {
											return nil, fmt.Errorf("rServiceResolversServerCertificates.RawDer: expected string")
										}
									}
									if _, ok := objval["serialNumber"]; ok {
										if s, ok := objval["serialNumber"].(string); ok {
											rServiceResolversServerCertificates.SerialNumber = dcl.String(s)
										} else {
											return nil, fmt.Errorf("rServiceResolversServerCertificates.SerialNumber: expected string")
										}
									}
									if _, ok := objval["sha256Fingerprint"]; ok {
										if s, ok := objval["sha256Fingerprint"].(string); ok {
											rServiceResolversServerCertificates.Sha256Fingerprint = dcl.String(s)
										} else {
											return nil, fmt.Errorf("rServiceResolversServerCertificates.Sha256Fingerprint: expected string")
										}
									}
									if _, ok := objval["subject"]; ok {
										if s, ok := objval["subject"].(string); ok {
											rServiceResolversServerCertificates.Subject = dcl.String(s)
										} else {
											return nil, fmt.Errorf("rServiceResolversServerCertificates.Subject: expected string")
										}
									}
									if _, ok := objval["subjectAlternativeDnsNames"]; ok {
										if s, ok := objval["subjectAlternativeDnsNames"].([]interface{}); ok {
											for _, ss := range s {
												if strval, ok := ss.(string); ok {
													rServiceResolversServerCertificates.SubjectAlternativeDnsNames = append(rServiceResolversServerCertificates.SubjectAlternativeDnsNames, strval)
												}
											}
										} else {
											return nil, fmt.Errorf("rServiceResolversServerCertificates.SubjectAlternativeDnsNames: expected []interface{}")
										}
									}
									rServiceResolvers.ServerCertificates = append(rServiceResolvers.ServerCertificates, rServiceResolversServerCertificates)
								}
							}
						} else {
							return nil, fmt.Errorf("rServiceResolvers.ServerCertificates: expected []interface{}")
						}
					}
					if _, ok := objval["serviceDirectoryService"]; ok {
						if s, ok := objval["serviceDirectoryService"].(string); ok {
							rServiceResolvers.ServiceDirectoryService = dcl.String(s)
						} else {
							return nil, fmt.Errorf("rServiceResolvers.ServiceDirectoryService: expected string")
						}
					}
					r.ServiceResolvers = append(r.ServiceResolvers, rServiceResolvers)
				}
			}
		} else {
			return nil, fmt.Errorf("r.ServiceResolvers: expected []interface{}")
		}
	}
	return r, nil
}

func GetEkmConnection(ctx context.Context, config *dcl.Config, u *unstructured.Resource) (*unstructured.Resource, error) {
	c := dclService.NewClient(config)
	r, err := UnstructuredToEkmConnection(u)
	if err != nil {
		return nil, err
	}
	r, err = c.GetEkmConnection(ctx, r)
	if err != nil {
		return nil, err
	}
	return EkmConnectionToUnstructured(r), nil
}

func ListEkmConnection(ctx context.Context, config *dcl.Config, project string, location string) ([]*unstructured.Resource, error) {
	c := dclService.NewClient(config)
	l, err := c.ListEkmConnection(ctx, project, location)
	if err != nil {
		return nil, err
	}
	var resources []*unstructured.Resource
	for {
		for _, r := range l.Items {
			resources = append(resources, EkmConnectionToUnstructured(r))
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

func ApplyEkmConnection(ctx context.Context, config *dcl.Config, u *unstructured.Resource, opts ...dcl.ApplyOption) (*unstructured.Resource, error) {
	c := dclService.NewClient(config)
	r, err := UnstructuredToEkmConnection(u)
	if err != nil {
		return nil, err
	}
	if ush := unstructured.FetchStateHint(opts); ush != nil {
		sh, err := UnstructuredToEkmConnection(ush)
		if err != nil {
			return nil, err
		}
		opts = append(opts, dcl.WithStateHint(sh))
	}
	r, err = c.ApplyEkmConnection(ctx, r, opts...)
	if err != nil {
		return nil, err
	}
	return EkmConnectionToUnstructured(r), nil
}

func EkmConnectionHasDiff(ctx context.Context, config *dcl.Config, u *unstructured.Resource, opts ...dcl.ApplyOption) (bool, error) {
	c := dclService.NewClient(config)
	r, err := UnstructuredToEkmConnection(u)
	if err != nil {
		return false, err
	}
	if ush := unstructured.FetchStateHint(opts); ush != nil {
		sh, err := UnstructuredToEkmConnection(ush)
		if err != nil {
			return false, err
		}
		opts = append(opts, dcl.WithStateHint(sh))
	}
	opts = append(opts, dcl.WithLifecycleParam(dcl.BlockDestruction), dcl.WithLifecycleParam(dcl.BlockCreation), dcl.WithLifecycleParam(dcl.BlockModification))
	_, err = c.ApplyEkmConnection(ctx, r, opts...)
	if err != nil {
		if _, ok := err.(dcl.ApplyInfeasibleError); ok {
			return true, nil
		}
		return false, err
	}
	return false, nil
}

func DeleteEkmConnection(ctx context.Context, config *dcl.Config, u *unstructured.Resource) error {
	return unstructured.ErrNoSuchMethod
}

func EkmConnectionID(u *unstructured.Resource) (string, error) {
	r, err := UnstructuredToEkmConnection(u)
	if err != nil {
		return "", err
	}
	return r.ID()
}

func (r *EkmConnection) STV() unstructured.ServiceTypeVersion {
	return unstructured.ServiceTypeVersion{
		"cloudkms",
		"EkmConnection",
		"ga",
	}
}

func SetPolicyEkmConnection(ctx context.Context, config *dcl.Config, u *unstructured.Resource, p *unstructured.Resource) (*unstructured.Resource, error) {
	r, err := UnstructuredToEkmConnection(u)
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

func SetPolicyWithEtagEkmConnection(ctx context.Context, config *dcl.Config, u *unstructured.Resource, p *unstructured.Resource) (*unstructured.Resource, error) {
	r, err := UnstructuredToEkmConnection(u)
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

func GetPolicyEkmConnection(ctx context.Context, config *dcl.Config, u *unstructured.Resource) (*unstructured.Resource, error) {
	r, err := UnstructuredToEkmConnection(u)
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

func SetPolicyMemberEkmConnection(ctx context.Context, config *dcl.Config, u *unstructured.Resource, m *unstructured.Resource) (*unstructured.Resource, error) {
	r, err := UnstructuredToEkmConnection(u)
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

func GetPolicyMemberEkmConnection(ctx context.Context, config *dcl.Config, u *unstructured.Resource, role, member string) (*unstructured.Resource, error) {
	r, err := UnstructuredToEkmConnection(u)
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

func DeletePolicyMemberEkmConnection(ctx context.Context, config *dcl.Config, u *unstructured.Resource, m *unstructured.Resource) error {
	r, err := UnstructuredToEkmConnection(u)
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

func (r *EkmConnection) SetPolicyMember(ctx context.Context, config *dcl.Config, resource *unstructured.Resource, member *unstructured.Resource) (*unstructured.Resource, error) {
	return SetPolicyMemberEkmConnection(ctx, config, resource, member)
}

func (r *EkmConnection) GetPolicyMember(ctx context.Context, config *dcl.Config, resource *unstructured.Resource, role, member string) (*unstructured.Resource, error) {
	return GetPolicyMemberEkmConnection(ctx, config, resource, role, member)
}

func (r *EkmConnection) DeletePolicyMember(ctx context.Context, config *dcl.Config, resource *unstructured.Resource, member *unstructured.Resource) error {
	return DeletePolicyMemberEkmConnection(ctx, config, resource, member)
}

func (r *EkmConnection) SetPolicy(ctx context.Context, config *dcl.Config, resource *unstructured.Resource, policy *unstructured.Resource) (*unstructured.Resource, error) {
	return SetPolicyEkmConnection(ctx, config, resource, policy)
}

func (r *EkmConnection) SetPolicyWithEtag(ctx context.Context, config *dcl.Config, resource *unstructured.Resource, policy *unstructured.Resource) (*unstructured.Resource, error) {
	return SetPolicyWithEtagEkmConnection(ctx, config, resource, policy)
}

func (r *EkmConnection) GetPolicy(ctx context.Context, config *dcl.Config, resource *unstructured.Resource) (*unstructured.Resource, error) {
	return GetPolicyEkmConnection(ctx, config, resource)
}

func (r *EkmConnection) Get(ctx context.Context, config *dcl.Config, resource *unstructured.Resource) (*unstructured.Resource, error) {
	return GetEkmConnection(ctx, config, resource)
}

func (r *EkmConnection) Apply(ctx context.Context, config *dcl.Config, resource *unstructured.Resource, opts ...dcl.ApplyOption) (*unstructured.Resource, error) {
	return ApplyEkmConnection(ctx, config, resource, opts...)
}

func (r *EkmConnection) HasDiff(ctx context.Context, config *dcl.Config, resource *unstructured.Resource, opts ...dcl.ApplyOption) (bool, error) {
	return EkmConnectionHasDiff(ctx, config, resource, opts...)
}

func (r *EkmConnection) Delete(ctx context.Context, config *dcl.Config, resource *unstructured.Resource) error {
	return DeleteEkmConnection(ctx, config, resource)
}

func (r *EkmConnection) ID(resource *unstructured.Resource) (string, error) {
	return EkmConnectionID(resource)
}

func init() {
	unstructured.Register(&EkmConnection{})
}
