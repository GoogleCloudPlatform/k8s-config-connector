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
package gkehub

import (
	"context"
	"fmt"
	"github.com/GoogleCloudPlatform/declarative-resource-client-library/dcl"
	dclService "github.com/GoogleCloudPlatform/declarative-resource-client-library/services/google/gkehub/beta"
	"github.com/GoogleCloudPlatform/declarative-resource-client-library/unstructured"
)

type Membership struct{}

func MembershipToUnstructured(r *dclService.Membership) *unstructured.Resource {
	u := &unstructured.Resource{
		STV: unstructured.ServiceTypeVersion{
			Service: "gkehub",
			Version: "beta",
			Type:    "Membership",
		},
		Object: make(map[string]interface{}),
	}
	if r.Authority != nil && r.Authority != dclService.EmptyMembershipAuthority {
		rAuthority := make(map[string]interface{})
		if r.Authority.IdentityProvider != nil {
			rAuthority["identityProvider"] = *r.Authority.IdentityProvider
		}
		if r.Authority.Issuer != nil {
			rAuthority["issuer"] = *r.Authority.Issuer
		}
		if r.Authority.WorkloadIdentityPool != nil {
			rAuthority["workloadIdentityPool"] = *r.Authority.WorkloadIdentityPool
		}
		u.Object["authority"] = rAuthority
	}
	if r.CreateTime != nil {
		u.Object["createTime"] = *r.CreateTime
	}
	if r.DeleteTime != nil {
		u.Object["deleteTime"] = *r.DeleteTime
	}
	if r.Description != nil {
		u.Object["description"] = *r.Description
	}
	if r.Endpoint != nil && r.Endpoint != dclService.EmptyMembershipEndpoint {
		rEndpoint := make(map[string]interface{})
		if r.Endpoint.GkeCluster != nil && r.Endpoint.GkeCluster != dclService.EmptyMembershipEndpointGkeCluster {
			rEndpointGkeCluster := make(map[string]interface{})
			if r.Endpoint.GkeCluster.ResourceLink != nil {
				rEndpointGkeCluster["resourceLink"] = *r.Endpoint.GkeCluster.ResourceLink
			}
			rEndpoint["gkeCluster"] = rEndpointGkeCluster
		}
		if r.Endpoint.KubernetesMetadata != nil && r.Endpoint.KubernetesMetadata != dclService.EmptyMembershipEndpointKubernetesMetadata {
			rEndpointKubernetesMetadata := make(map[string]interface{})
			if r.Endpoint.KubernetesMetadata.KubernetesApiServerVersion != nil {
				rEndpointKubernetesMetadata["kubernetesApiServerVersion"] = *r.Endpoint.KubernetesMetadata.KubernetesApiServerVersion
			}
			if r.Endpoint.KubernetesMetadata.MemoryMb != nil {
				rEndpointKubernetesMetadata["memoryMb"] = *r.Endpoint.KubernetesMetadata.MemoryMb
			}
			if r.Endpoint.KubernetesMetadata.NodeCount != nil {
				rEndpointKubernetesMetadata["nodeCount"] = *r.Endpoint.KubernetesMetadata.NodeCount
			}
			if r.Endpoint.KubernetesMetadata.NodeProviderId != nil {
				rEndpointKubernetesMetadata["nodeProviderId"] = *r.Endpoint.KubernetesMetadata.NodeProviderId
			}
			if r.Endpoint.KubernetesMetadata.UpdateTime != nil {
				rEndpointKubernetesMetadata["updateTime"] = *r.Endpoint.KubernetesMetadata.UpdateTime
			}
			if r.Endpoint.KubernetesMetadata.VcpuCount != nil {
				rEndpointKubernetesMetadata["vcpuCount"] = *r.Endpoint.KubernetesMetadata.VcpuCount
			}
			rEndpoint["kubernetesMetadata"] = rEndpointKubernetesMetadata
		}
		if r.Endpoint.KubernetesResource != nil && r.Endpoint.KubernetesResource != dclService.EmptyMembershipEndpointKubernetesResource {
			rEndpointKubernetesResource := make(map[string]interface{})
			var rEndpointKubernetesResourceConnectResources []interface{}
			for _, rEndpointKubernetesResourceConnectResourcesVal := range r.Endpoint.KubernetesResource.ConnectResources {
				rEndpointKubernetesResourceConnectResourcesObject := make(map[string]interface{})
				if rEndpointKubernetesResourceConnectResourcesVal.ClusterScoped != nil {
					rEndpointKubernetesResourceConnectResourcesObject["clusterScoped"] = *rEndpointKubernetesResourceConnectResourcesVal.ClusterScoped
				}
				if rEndpointKubernetesResourceConnectResourcesVal.Manifest != nil {
					rEndpointKubernetesResourceConnectResourcesObject["manifest"] = *rEndpointKubernetesResourceConnectResourcesVal.Manifest
				}
				rEndpointKubernetesResourceConnectResources = append(rEndpointKubernetesResourceConnectResources, rEndpointKubernetesResourceConnectResourcesObject)
			}
			rEndpointKubernetesResource["connectResources"] = rEndpointKubernetesResourceConnectResources
			if r.Endpoint.KubernetesResource.MembershipCrManifest != nil {
				rEndpointKubernetesResource["membershipCrManifest"] = *r.Endpoint.KubernetesResource.MembershipCrManifest
			}
			var rEndpointKubernetesResourceMembershipResources []interface{}
			for _, rEndpointKubernetesResourceMembershipResourcesVal := range r.Endpoint.KubernetesResource.MembershipResources {
				rEndpointKubernetesResourceMembershipResourcesObject := make(map[string]interface{})
				if rEndpointKubernetesResourceMembershipResourcesVal.ClusterScoped != nil {
					rEndpointKubernetesResourceMembershipResourcesObject["clusterScoped"] = *rEndpointKubernetesResourceMembershipResourcesVal.ClusterScoped
				}
				if rEndpointKubernetesResourceMembershipResourcesVal.Manifest != nil {
					rEndpointKubernetesResourceMembershipResourcesObject["manifest"] = *rEndpointKubernetesResourceMembershipResourcesVal.Manifest
				}
				rEndpointKubernetesResourceMembershipResources = append(rEndpointKubernetesResourceMembershipResources, rEndpointKubernetesResourceMembershipResourcesObject)
			}
			rEndpointKubernetesResource["membershipResources"] = rEndpointKubernetesResourceMembershipResources
			if r.Endpoint.KubernetesResource.ResourceOptions != nil && r.Endpoint.KubernetesResource.ResourceOptions != dclService.EmptyMembershipEndpointKubernetesResourceResourceOptions {
				rEndpointKubernetesResourceResourceOptions := make(map[string]interface{})
				if r.Endpoint.KubernetesResource.ResourceOptions.ConnectVersion != nil {
					rEndpointKubernetesResourceResourceOptions["connectVersion"] = *r.Endpoint.KubernetesResource.ResourceOptions.ConnectVersion
				}
				if r.Endpoint.KubernetesResource.ResourceOptions.V1Beta1Crd != nil {
					rEndpointKubernetesResourceResourceOptions["v1beta1Crd"] = *r.Endpoint.KubernetesResource.ResourceOptions.V1Beta1Crd
				}
				rEndpointKubernetesResource["resourceOptions"] = rEndpointKubernetesResourceResourceOptions
			}
			rEndpoint["kubernetesResource"] = rEndpointKubernetesResource
		}
		u.Object["endpoint"] = rEndpoint
	}
	if r.ExternalId != nil {
		u.Object["externalId"] = *r.ExternalId
	}
	if r.InfrastructureType != nil {
		u.Object["infrastructureType"] = string(*r.InfrastructureType)
	}
	if r.Labels != nil {
		rLabels := make(map[string]interface{})
		for k, v := range r.Labels {
			rLabels[k] = v
		}
		u.Object["labels"] = rLabels
	}
	if r.LastConnectionTime != nil {
		u.Object["lastConnectionTime"] = *r.LastConnectionTime
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
	if r.State != nil && r.State != dclService.EmptyMembershipState {
		rState := make(map[string]interface{})
		if r.State.Code != nil {
			rState["code"] = string(*r.State.Code)
		}
		u.Object["state"] = rState
	}
	if r.UniqueId != nil {
		u.Object["uniqueId"] = *r.UniqueId
	}
	if r.UpdateTime != nil {
		u.Object["updateTime"] = *r.UpdateTime
	}
	return u
}

func UnstructuredToMembership(u *unstructured.Resource) (*dclService.Membership, error) {
	r := &dclService.Membership{}
	if _, ok := u.Object["authority"]; ok {
		if rAuthority, ok := u.Object["authority"].(map[string]interface{}); ok {
			r.Authority = &dclService.MembershipAuthority{}
			if _, ok := rAuthority["identityProvider"]; ok {
				if s, ok := rAuthority["identityProvider"].(string); ok {
					r.Authority.IdentityProvider = dcl.String(s)
				} else {
					return nil, fmt.Errorf("r.Authority.IdentityProvider: expected string")
				}
			}
			if _, ok := rAuthority["issuer"]; ok {
				if s, ok := rAuthority["issuer"].(string); ok {
					r.Authority.Issuer = dcl.String(s)
				} else {
					return nil, fmt.Errorf("r.Authority.Issuer: expected string")
				}
			}
			if _, ok := rAuthority["workloadIdentityPool"]; ok {
				if s, ok := rAuthority["workloadIdentityPool"].(string); ok {
					r.Authority.WorkloadIdentityPool = dcl.String(s)
				} else {
					return nil, fmt.Errorf("r.Authority.WorkloadIdentityPool: expected string")
				}
			}
		} else {
			return nil, fmt.Errorf("r.Authority: expected map[string]interface{}")
		}
	}
	if _, ok := u.Object["createTime"]; ok {
		if s, ok := u.Object["createTime"].(string); ok {
			r.CreateTime = dcl.String(s)
		} else {
			return nil, fmt.Errorf("r.CreateTime: expected string")
		}
	}
	if _, ok := u.Object["deleteTime"]; ok {
		if s, ok := u.Object["deleteTime"].(string); ok {
			r.DeleteTime = dcl.String(s)
		} else {
			return nil, fmt.Errorf("r.DeleteTime: expected string")
		}
	}
	if _, ok := u.Object["description"]; ok {
		if s, ok := u.Object["description"].(string); ok {
			r.Description = dcl.String(s)
		} else {
			return nil, fmt.Errorf("r.Description: expected string")
		}
	}
	if _, ok := u.Object["endpoint"]; ok {
		if rEndpoint, ok := u.Object["endpoint"].(map[string]interface{}); ok {
			r.Endpoint = &dclService.MembershipEndpoint{}
			if _, ok := rEndpoint["gkeCluster"]; ok {
				if rEndpointGkeCluster, ok := rEndpoint["gkeCluster"].(map[string]interface{}); ok {
					r.Endpoint.GkeCluster = &dclService.MembershipEndpointGkeCluster{}
					if _, ok := rEndpointGkeCluster["resourceLink"]; ok {
						if s, ok := rEndpointGkeCluster["resourceLink"].(string); ok {
							r.Endpoint.GkeCluster.ResourceLink = dcl.String(s)
						} else {
							return nil, fmt.Errorf("r.Endpoint.GkeCluster.ResourceLink: expected string")
						}
					}
				} else {
					return nil, fmt.Errorf("r.Endpoint.GkeCluster: expected map[string]interface{}")
				}
			}
			if _, ok := rEndpoint["kubernetesMetadata"]; ok {
				if rEndpointKubernetesMetadata, ok := rEndpoint["kubernetesMetadata"].(map[string]interface{}); ok {
					r.Endpoint.KubernetesMetadata = &dclService.MembershipEndpointKubernetesMetadata{}
					if _, ok := rEndpointKubernetesMetadata["kubernetesApiServerVersion"]; ok {
						if s, ok := rEndpointKubernetesMetadata["kubernetesApiServerVersion"].(string); ok {
							r.Endpoint.KubernetesMetadata.KubernetesApiServerVersion = dcl.String(s)
						} else {
							return nil, fmt.Errorf("r.Endpoint.KubernetesMetadata.KubernetesApiServerVersion: expected string")
						}
					}
					if _, ok := rEndpointKubernetesMetadata["memoryMb"]; ok {
						if i, ok := rEndpointKubernetesMetadata["memoryMb"].(int64); ok {
							r.Endpoint.KubernetesMetadata.MemoryMb = dcl.Int64(i)
						} else {
							return nil, fmt.Errorf("r.Endpoint.KubernetesMetadata.MemoryMb: expected int64")
						}
					}
					if _, ok := rEndpointKubernetesMetadata["nodeCount"]; ok {
						if i, ok := rEndpointKubernetesMetadata["nodeCount"].(int64); ok {
							r.Endpoint.KubernetesMetadata.NodeCount = dcl.Int64(i)
						} else {
							return nil, fmt.Errorf("r.Endpoint.KubernetesMetadata.NodeCount: expected int64")
						}
					}
					if _, ok := rEndpointKubernetesMetadata["nodeProviderId"]; ok {
						if s, ok := rEndpointKubernetesMetadata["nodeProviderId"].(string); ok {
							r.Endpoint.KubernetesMetadata.NodeProviderId = dcl.String(s)
						} else {
							return nil, fmt.Errorf("r.Endpoint.KubernetesMetadata.NodeProviderId: expected string")
						}
					}
					if _, ok := rEndpointKubernetesMetadata["updateTime"]; ok {
						if s, ok := rEndpointKubernetesMetadata["updateTime"].(string); ok {
							r.Endpoint.KubernetesMetadata.UpdateTime = dcl.String(s)
						} else {
							return nil, fmt.Errorf("r.Endpoint.KubernetesMetadata.UpdateTime: expected string")
						}
					}
					if _, ok := rEndpointKubernetesMetadata["vcpuCount"]; ok {
						if i, ok := rEndpointKubernetesMetadata["vcpuCount"].(int64); ok {
							r.Endpoint.KubernetesMetadata.VcpuCount = dcl.Int64(i)
						} else {
							return nil, fmt.Errorf("r.Endpoint.KubernetesMetadata.VcpuCount: expected int64")
						}
					}
				} else {
					return nil, fmt.Errorf("r.Endpoint.KubernetesMetadata: expected map[string]interface{}")
				}
			}
			if _, ok := rEndpoint["kubernetesResource"]; ok {
				if rEndpointKubernetesResource, ok := rEndpoint["kubernetesResource"].(map[string]interface{}); ok {
					r.Endpoint.KubernetesResource = &dclService.MembershipEndpointKubernetesResource{}
					if _, ok := rEndpointKubernetesResource["connectResources"]; ok {
						if s, ok := rEndpointKubernetesResource["connectResources"].([]interface{}); ok {
							for _, o := range s {
								if objval, ok := o.(map[string]interface{}); ok {
									var rEndpointKubernetesResourceConnectResources dclService.MembershipEndpointKubernetesResourceConnectResources
									if _, ok := objval["clusterScoped"]; ok {
										if b, ok := objval["clusterScoped"].(bool); ok {
											rEndpointKubernetesResourceConnectResources.ClusterScoped = dcl.Bool(b)
										} else {
											return nil, fmt.Errorf("rEndpointKubernetesResourceConnectResources.ClusterScoped: expected bool")
										}
									}
									if _, ok := objval["manifest"]; ok {
										if s, ok := objval["manifest"].(string); ok {
											rEndpointKubernetesResourceConnectResources.Manifest = dcl.String(s)
										} else {
											return nil, fmt.Errorf("rEndpointKubernetesResourceConnectResources.Manifest: expected string")
										}
									}
									r.Endpoint.KubernetesResource.ConnectResources = append(r.Endpoint.KubernetesResource.ConnectResources, rEndpointKubernetesResourceConnectResources)
								}
							}
						} else {
							return nil, fmt.Errorf("r.Endpoint.KubernetesResource.ConnectResources: expected []interface{}")
						}
					}
					if _, ok := rEndpointKubernetesResource["membershipCrManifest"]; ok {
						if s, ok := rEndpointKubernetesResource["membershipCrManifest"].(string); ok {
							r.Endpoint.KubernetesResource.MembershipCrManifest = dcl.String(s)
						} else {
							return nil, fmt.Errorf("r.Endpoint.KubernetesResource.MembershipCrManifest: expected string")
						}
					}
					if _, ok := rEndpointKubernetesResource["membershipResources"]; ok {
						if s, ok := rEndpointKubernetesResource["membershipResources"].([]interface{}); ok {
							for _, o := range s {
								if objval, ok := o.(map[string]interface{}); ok {
									var rEndpointKubernetesResourceMembershipResources dclService.MembershipEndpointKubernetesResourceMembershipResources
									if _, ok := objval["clusterScoped"]; ok {
										if b, ok := objval["clusterScoped"].(bool); ok {
											rEndpointKubernetesResourceMembershipResources.ClusterScoped = dcl.Bool(b)
										} else {
											return nil, fmt.Errorf("rEndpointKubernetesResourceMembershipResources.ClusterScoped: expected bool")
										}
									}
									if _, ok := objval["manifest"]; ok {
										if s, ok := objval["manifest"].(string); ok {
											rEndpointKubernetesResourceMembershipResources.Manifest = dcl.String(s)
										} else {
											return nil, fmt.Errorf("rEndpointKubernetesResourceMembershipResources.Manifest: expected string")
										}
									}
									r.Endpoint.KubernetesResource.MembershipResources = append(r.Endpoint.KubernetesResource.MembershipResources, rEndpointKubernetesResourceMembershipResources)
								}
							}
						} else {
							return nil, fmt.Errorf("r.Endpoint.KubernetesResource.MembershipResources: expected []interface{}")
						}
					}
					if _, ok := rEndpointKubernetesResource["resourceOptions"]; ok {
						if rEndpointKubernetesResourceResourceOptions, ok := rEndpointKubernetesResource["resourceOptions"].(map[string]interface{}); ok {
							r.Endpoint.KubernetesResource.ResourceOptions = &dclService.MembershipEndpointKubernetesResourceResourceOptions{}
							if _, ok := rEndpointKubernetesResourceResourceOptions["connectVersion"]; ok {
								if s, ok := rEndpointKubernetesResourceResourceOptions["connectVersion"].(string); ok {
									r.Endpoint.KubernetesResource.ResourceOptions.ConnectVersion = dcl.String(s)
								} else {
									return nil, fmt.Errorf("r.Endpoint.KubernetesResource.ResourceOptions.ConnectVersion: expected string")
								}
							}
							if _, ok := rEndpointKubernetesResourceResourceOptions["v1beta1Crd"]; ok {
								if b, ok := rEndpointKubernetesResourceResourceOptions["v1beta1Crd"].(bool); ok {
									r.Endpoint.KubernetesResource.ResourceOptions.V1Beta1Crd = dcl.Bool(b)
								} else {
									return nil, fmt.Errorf("r.Endpoint.KubernetesResource.ResourceOptions.V1Beta1Crd: expected bool")
								}
							}
						} else {
							return nil, fmt.Errorf("r.Endpoint.KubernetesResource.ResourceOptions: expected map[string]interface{}")
						}
					}
				} else {
					return nil, fmt.Errorf("r.Endpoint.KubernetesResource: expected map[string]interface{}")
				}
			}
		} else {
			return nil, fmt.Errorf("r.Endpoint: expected map[string]interface{}")
		}
	}
	if _, ok := u.Object["externalId"]; ok {
		if s, ok := u.Object["externalId"].(string); ok {
			r.ExternalId = dcl.String(s)
		} else {
			return nil, fmt.Errorf("r.ExternalId: expected string")
		}
	}
	if _, ok := u.Object["infrastructureType"]; ok {
		if s, ok := u.Object["infrastructureType"].(string); ok {
			r.InfrastructureType = dclService.MembershipInfrastructureTypeEnumRef(s)
		} else {
			return nil, fmt.Errorf("r.InfrastructureType: expected string")
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
	if _, ok := u.Object["lastConnectionTime"]; ok {
		if s, ok := u.Object["lastConnectionTime"].(string); ok {
			r.LastConnectionTime = dcl.String(s)
		} else {
			return nil, fmt.Errorf("r.LastConnectionTime: expected string")
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
	if _, ok := u.Object["state"]; ok {
		if rState, ok := u.Object["state"].(map[string]interface{}); ok {
			r.State = &dclService.MembershipState{}
			if _, ok := rState["code"]; ok {
				if s, ok := rState["code"].(string); ok {
					r.State.Code = dclService.MembershipStateCodeEnumRef(s)
				} else {
					return nil, fmt.Errorf("r.State.Code: expected string")
				}
			}
		} else {
			return nil, fmt.Errorf("r.State: expected map[string]interface{}")
		}
	}
	if _, ok := u.Object["uniqueId"]; ok {
		if s, ok := u.Object["uniqueId"].(string); ok {
			r.UniqueId = dcl.String(s)
		} else {
			return nil, fmt.Errorf("r.UniqueId: expected string")
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

func GetMembership(ctx context.Context, config *dcl.Config, u *unstructured.Resource) (*unstructured.Resource, error) {
	c := dclService.NewClient(config)
	r, err := UnstructuredToMembership(u)
	if err != nil {
		return nil, err
	}
	r, err = c.GetMembership(ctx, r)
	if err != nil {
		return nil, err
	}
	return MembershipToUnstructured(r), nil
}

func ListMembership(ctx context.Context, config *dcl.Config, project string, location string) ([]*unstructured.Resource, error) {
	c := dclService.NewClient(config)
	l, err := c.ListMembership(ctx, project, location)
	if err != nil {
		return nil, err
	}
	var resources []*unstructured.Resource
	for {
		for _, r := range l.Items {
			resources = append(resources, MembershipToUnstructured(r))
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

func ApplyMembership(ctx context.Context, config *dcl.Config, u *unstructured.Resource, opts ...dcl.ApplyOption) (*unstructured.Resource, error) {
	c := dclService.NewClient(config)
	r, err := UnstructuredToMembership(u)
	if err != nil {
		return nil, err
	}
	if ush := unstructured.FetchStateHint(opts); ush != nil {
		sh, err := UnstructuredToMembership(ush)
		if err != nil {
			return nil, err
		}
		opts = append(opts, dcl.WithStateHint(sh))
	}
	r, err = c.ApplyMembership(ctx, r, opts...)
	if err != nil {
		return nil, err
	}
	return MembershipToUnstructured(r), nil
}

func MembershipHasDiff(ctx context.Context, config *dcl.Config, u *unstructured.Resource, opts ...dcl.ApplyOption) (bool, error) {
	c := dclService.NewClient(config)
	r, err := UnstructuredToMembership(u)
	if err != nil {
		return false, err
	}
	if ush := unstructured.FetchStateHint(opts); ush != nil {
		sh, err := UnstructuredToMembership(ush)
		if err != nil {
			return false, err
		}
		opts = append(opts, dcl.WithStateHint(sh))
	}
	opts = append(opts, dcl.WithLifecycleParam(dcl.BlockDestruction), dcl.WithLifecycleParam(dcl.BlockCreation), dcl.WithLifecycleParam(dcl.BlockModification))
	_, err = c.ApplyMembership(ctx, r, opts...)
	if err != nil {
		if _, ok := err.(dcl.ApplyInfeasibleError); ok {
			return true, nil
		}
		return false, err
	}
	return false, nil
}

func DeleteMembership(ctx context.Context, config *dcl.Config, u *unstructured.Resource) error {
	c := dclService.NewClient(config)
	r, err := UnstructuredToMembership(u)
	if err != nil {
		return err
	}
	return c.DeleteMembership(ctx, r)
}

func MembershipID(u *unstructured.Resource) (string, error) {
	r, err := UnstructuredToMembership(u)
	if err != nil {
		return "", err
	}
	return r.ID()
}

func (r *Membership) STV() unstructured.ServiceTypeVersion {
	return unstructured.ServiceTypeVersion{
		"gkehub",
		"Membership",
		"beta",
	}
}

func (r *Membership) SetPolicyMember(ctx context.Context, config *dcl.Config, resource *unstructured.Resource, member *unstructured.Resource) (*unstructured.Resource, error) {
	return nil, unstructured.ErrNoSuchMethod
}

func (r *Membership) GetPolicyMember(ctx context.Context, config *dcl.Config, resource *unstructured.Resource, role, member string) (*unstructured.Resource, error) {
	return nil, unstructured.ErrNoSuchMethod
}

func (r *Membership) DeletePolicyMember(ctx context.Context, config *dcl.Config, resource *unstructured.Resource, member *unstructured.Resource) error {
	return unstructured.ErrNoSuchMethod
}

func (r *Membership) SetPolicy(ctx context.Context, config *dcl.Config, resource *unstructured.Resource, policy *unstructured.Resource) (*unstructured.Resource, error) {
	return nil, unstructured.ErrNoSuchMethod
}

func (r *Membership) SetPolicyWithEtag(ctx context.Context, config *dcl.Config, resource *unstructured.Resource, policy *unstructured.Resource) (*unstructured.Resource, error) {
	return nil, unstructured.ErrNoSuchMethod
}

func (r *Membership) GetPolicy(ctx context.Context, config *dcl.Config, resource *unstructured.Resource) (*unstructured.Resource, error) {
	return nil, unstructured.ErrNoSuchMethod
}

func (r *Membership) Get(ctx context.Context, config *dcl.Config, resource *unstructured.Resource) (*unstructured.Resource, error) {
	return GetMembership(ctx, config, resource)
}

func (r *Membership) Apply(ctx context.Context, config *dcl.Config, resource *unstructured.Resource, opts ...dcl.ApplyOption) (*unstructured.Resource, error) {
	return ApplyMembership(ctx, config, resource, opts...)
}

func (r *Membership) HasDiff(ctx context.Context, config *dcl.Config, resource *unstructured.Resource, opts ...dcl.ApplyOption) (bool, error) {
	return MembershipHasDiff(ctx, config, resource, opts...)
}

func (r *Membership) Delete(ctx context.Context, config *dcl.Config, resource *unstructured.Resource) error {
	return DeleteMembership(ctx, config, resource)
}

func (r *Membership) ID(resource *unstructured.Resource) (string, error) {
	return MembershipID(resource)
}

func init() {
	unstructured.Register(&Membership{})
}
