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
package containerazure

import (
	"context"
	"fmt"
	"github.com/GoogleCloudPlatform/declarative-resource-client-library/dcl"
	dclService "github.com/GoogleCloudPlatform/declarative-resource-client-library/services/google/containerazure/beta"
	"github.com/GoogleCloudPlatform/declarative-resource-client-library/unstructured"
)

type Cluster struct{}

func ClusterToUnstructured(r *dclService.Cluster) *unstructured.Resource {
	u := &unstructured.Resource{
		STV: unstructured.ServiceTypeVersion{
			Service: "containerazure",
			Version: "beta",
			Type:    "Cluster",
		},
		Object: make(map[string]interface{}),
	}
	if r.Annotations != nil {
		rAnnotations := make(map[string]interface{})
		for k, v := range r.Annotations {
			rAnnotations[k] = v
		}
		u.Object["annotations"] = rAnnotations
	}
	if r.Authorization != nil && r.Authorization != dclService.EmptyClusterAuthorization {
		rAuthorization := make(map[string]interface{})
		var rAuthorizationAdminGroups []interface{}
		for _, rAuthorizationAdminGroupsVal := range r.Authorization.AdminGroups {
			rAuthorizationAdminGroupsObject := make(map[string]interface{})
			if rAuthorizationAdminGroupsVal.Group != nil {
				rAuthorizationAdminGroupsObject["group"] = *rAuthorizationAdminGroupsVal.Group
			}
			rAuthorizationAdminGroups = append(rAuthorizationAdminGroups, rAuthorizationAdminGroupsObject)
		}
		rAuthorization["adminGroups"] = rAuthorizationAdminGroups
		var rAuthorizationAdminUsers []interface{}
		for _, rAuthorizationAdminUsersVal := range r.Authorization.AdminUsers {
			rAuthorizationAdminUsersObject := make(map[string]interface{})
			if rAuthorizationAdminUsersVal.Username != nil {
				rAuthorizationAdminUsersObject["username"] = *rAuthorizationAdminUsersVal.Username
			}
			rAuthorizationAdminUsers = append(rAuthorizationAdminUsers, rAuthorizationAdminUsersObject)
		}
		rAuthorization["adminUsers"] = rAuthorizationAdminUsers
		u.Object["authorization"] = rAuthorization
	}
	if r.AzureRegion != nil {
		u.Object["azureRegion"] = *r.AzureRegion
	}
	if r.AzureServicesAuthentication != nil && r.AzureServicesAuthentication != dclService.EmptyClusterAzureServicesAuthentication {
		rAzureServicesAuthentication := make(map[string]interface{})
		if r.AzureServicesAuthentication.ApplicationId != nil {
			rAzureServicesAuthentication["applicationId"] = *r.AzureServicesAuthentication.ApplicationId
		}
		if r.AzureServicesAuthentication.TenantId != nil {
			rAzureServicesAuthentication["tenantId"] = *r.AzureServicesAuthentication.TenantId
		}
		u.Object["azureServicesAuthentication"] = rAzureServicesAuthentication
	}
	if r.Client != nil {
		u.Object["client"] = *r.Client
	}
	if r.ControlPlane != nil && r.ControlPlane != dclService.EmptyClusterControlPlane {
		rControlPlane := make(map[string]interface{})
		if r.ControlPlane.DatabaseEncryption != nil && r.ControlPlane.DatabaseEncryption != dclService.EmptyClusterControlPlaneDatabaseEncryption {
			rControlPlaneDatabaseEncryption := make(map[string]interface{})
			if r.ControlPlane.DatabaseEncryption.KeyId != nil {
				rControlPlaneDatabaseEncryption["keyId"] = *r.ControlPlane.DatabaseEncryption.KeyId
			}
			rControlPlane["databaseEncryption"] = rControlPlaneDatabaseEncryption
		}
		if r.ControlPlane.MainVolume != nil && r.ControlPlane.MainVolume != dclService.EmptyClusterControlPlaneMainVolume {
			rControlPlaneMainVolume := make(map[string]interface{})
			if r.ControlPlane.MainVolume.SizeGib != nil {
				rControlPlaneMainVolume["sizeGib"] = *r.ControlPlane.MainVolume.SizeGib
			}
			rControlPlane["mainVolume"] = rControlPlaneMainVolume
		}
		if r.ControlPlane.ProxyConfig != nil && r.ControlPlane.ProxyConfig != dclService.EmptyClusterControlPlaneProxyConfig {
			rControlPlaneProxyConfig := make(map[string]interface{})
			if r.ControlPlane.ProxyConfig.ResourceGroupId != nil {
				rControlPlaneProxyConfig["resourceGroupId"] = *r.ControlPlane.ProxyConfig.ResourceGroupId
			}
			if r.ControlPlane.ProxyConfig.SecretId != nil {
				rControlPlaneProxyConfig["secretId"] = *r.ControlPlane.ProxyConfig.SecretId
			}
			rControlPlane["proxyConfig"] = rControlPlaneProxyConfig
		}
		var rControlPlaneReplicaPlacements []interface{}
		for _, rControlPlaneReplicaPlacementsVal := range r.ControlPlane.ReplicaPlacements {
			rControlPlaneReplicaPlacementsObject := make(map[string]interface{})
			if rControlPlaneReplicaPlacementsVal.AzureAvailabilityZone != nil {
				rControlPlaneReplicaPlacementsObject["azureAvailabilityZone"] = *rControlPlaneReplicaPlacementsVal.AzureAvailabilityZone
			}
			if rControlPlaneReplicaPlacementsVal.SubnetId != nil {
				rControlPlaneReplicaPlacementsObject["subnetId"] = *rControlPlaneReplicaPlacementsVal.SubnetId
			}
			rControlPlaneReplicaPlacements = append(rControlPlaneReplicaPlacements, rControlPlaneReplicaPlacementsObject)
		}
		rControlPlane["replicaPlacements"] = rControlPlaneReplicaPlacements
		if r.ControlPlane.RootVolume != nil && r.ControlPlane.RootVolume != dclService.EmptyClusterControlPlaneRootVolume {
			rControlPlaneRootVolume := make(map[string]interface{})
			if r.ControlPlane.RootVolume.SizeGib != nil {
				rControlPlaneRootVolume["sizeGib"] = *r.ControlPlane.RootVolume.SizeGib
			}
			rControlPlane["rootVolume"] = rControlPlaneRootVolume
		}
		if r.ControlPlane.SshConfig != nil && r.ControlPlane.SshConfig != dclService.EmptyClusterControlPlaneSshConfig {
			rControlPlaneSshConfig := make(map[string]interface{})
			if r.ControlPlane.SshConfig.AuthorizedKey != nil {
				rControlPlaneSshConfig["authorizedKey"] = *r.ControlPlane.SshConfig.AuthorizedKey
			}
			rControlPlane["sshConfig"] = rControlPlaneSshConfig
		}
		if r.ControlPlane.SubnetId != nil {
			rControlPlane["subnetId"] = *r.ControlPlane.SubnetId
		}
		if r.ControlPlane.Tags != nil {
			rControlPlaneTags := make(map[string]interface{})
			for k, v := range r.ControlPlane.Tags {
				rControlPlaneTags[k] = v
			}
			rControlPlane["tags"] = rControlPlaneTags
		}
		if r.ControlPlane.Version != nil {
			rControlPlane["version"] = *r.ControlPlane.Version
		}
		if r.ControlPlane.VmSize != nil {
			rControlPlane["vmSize"] = *r.ControlPlane.VmSize
		}
		u.Object["controlPlane"] = rControlPlane
	}
	if r.CreateTime != nil {
		u.Object["createTime"] = *r.CreateTime
	}
	if r.Description != nil {
		u.Object["description"] = *r.Description
	}
	if r.Endpoint != nil {
		u.Object["endpoint"] = *r.Endpoint
	}
	if r.Etag != nil {
		u.Object["etag"] = *r.Etag
	}
	if r.Fleet != nil && r.Fleet != dclService.EmptyClusterFleet {
		rFleet := make(map[string]interface{})
		if r.Fleet.Membership != nil {
			rFleet["membership"] = *r.Fleet.Membership
		}
		if r.Fleet.Project != nil {
			rFleet["project"] = *r.Fleet.Project
		}
		u.Object["fleet"] = rFleet
	}
	if r.Location != nil {
		u.Object["location"] = *r.Location
	}
	if r.LoggingConfig != nil && r.LoggingConfig != dclService.EmptyClusterLoggingConfig {
		rLoggingConfig := make(map[string]interface{})
		if r.LoggingConfig.ComponentConfig != nil && r.LoggingConfig.ComponentConfig != dclService.EmptyClusterLoggingConfigComponentConfig {
			rLoggingConfigComponentConfig := make(map[string]interface{})
			var rLoggingConfigComponentConfigEnableComponents []interface{}
			for _, rLoggingConfigComponentConfigEnableComponentsVal := range r.LoggingConfig.ComponentConfig.EnableComponents {
				rLoggingConfigComponentConfigEnableComponents = append(rLoggingConfigComponentConfigEnableComponents, string(rLoggingConfigComponentConfigEnableComponentsVal))
			}
			rLoggingConfigComponentConfig["enableComponents"] = rLoggingConfigComponentConfigEnableComponents
			rLoggingConfig["componentConfig"] = rLoggingConfigComponentConfig
		}
		u.Object["loggingConfig"] = rLoggingConfig
	}
	if r.MonitoringConfig != nil && r.MonitoringConfig != dclService.EmptyClusterMonitoringConfig {
		rMonitoringConfig := make(map[string]interface{})
		if r.MonitoringConfig.ManagedPrometheusConfig != nil && r.MonitoringConfig.ManagedPrometheusConfig != dclService.EmptyClusterMonitoringConfigManagedPrometheusConfig {
			rMonitoringConfigManagedPrometheusConfig := make(map[string]interface{})
			if r.MonitoringConfig.ManagedPrometheusConfig.Enabled != nil {
				rMonitoringConfigManagedPrometheusConfig["enabled"] = *r.MonitoringConfig.ManagedPrometheusConfig.Enabled
			}
			rMonitoringConfig["managedPrometheusConfig"] = rMonitoringConfigManagedPrometheusConfig
		}
		u.Object["monitoringConfig"] = rMonitoringConfig
	}
	if r.Name != nil {
		u.Object["name"] = *r.Name
	}
	if r.Networking != nil && r.Networking != dclService.EmptyClusterNetworking {
		rNetworking := make(map[string]interface{})
		var rNetworkingPodAddressCidrBlocks []interface{}
		for _, rNetworkingPodAddressCidrBlocksVal := range r.Networking.PodAddressCidrBlocks {
			rNetworkingPodAddressCidrBlocks = append(rNetworkingPodAddressCidrBlocks, rNetworkingPodAddressCidrBlocksVal)
		}
		rNetworking["podAddressCidrBlocks"] = rNetworkingPodAddressCidrBlocks
		var rNetworkingServiceAddressCidrBlocks []interface{}
		for _, rNetworkingServiceAddressCidrBlocksVal := range r.Networking.ServiceAddressCidrBlocks {
			rNetworkingServiceAddressCidrBlocks = append(rNetworkingServiceAddressCidrBlocks, rNetworkingServiceAddressCidrBlocksVal)
		}
		rNetworking["serviceAddressCidrBlocks"] = rNetworkingServiceAddressCidrBlocks
		if r.Networking.VirtualNetworkId != nil {
			rNetworking["virtualNetworkId"] = *r.Networking.VirtualNetworkId
		}
		u.Object["networking"] = rNetworking
	}
	if r.Project != nil {
		u.Object["project"] = *r.Project
	}
	if r.Reconciling != nil {
		u.Object["reconciling"] = *r.Reconciling
	}
	if r.ResourceGroupId != nil {
		u.Object["resourceGroupId"] = *r.ResourceGroupId
	}
	if r.State != nil {
		u.Object["state"] = string(*r.State)
	}
	if r.Uid != nil {
		u.Object["uid"] = *r.Uid
	}
	if r.UpdateTime != nil {
		u.Object["updateTime"] = *r.UpdateTime
	}
	if r.WorkloadIdentityConfig != nil && r.WorkloadIdentityConfig != dclService.EmptyClusterWorkloadIdentityConfig {
		rWorkloadIdentityConfig := make(map[string]interface{})
		if r.WorkloadIdentityConfig.IdentityProvider != nil {
			rWorkloadIdentityConfig["identityProvider"] = *r.WorkloadIdentityConfig.IdentityProvider
		}
		if r.WorkloadIdentityConfig.IssuerUri != nil {
			rWorkloadIdentityConfig["issuerUri"] = *r.WorkloadIdentityConfig.IssuerUri
		}
		if r.WorkloadIdentityConfig.WorkloadPool != nil {
			rWorkloadIdentityConfig["workloadPool"] = *r.WorkloadIdentityConfig.WorkloadPool
		}
		u.Object["workloadIdentityConfig"] = rWorkloadIdentityConfig
	}
	return u
}

func UnstructuredToCluster(u *unstructured.Resource) (*dclService.Cluster, error) {
	r := &dclService.Cluster{}
	if _, ok := u.Object["annotations"]; ok {
		if rAnnotations, ok := u.Object["annotations"].(map[string]interface{}); ok {
			m := make(map[string]string)
			for k, v := range rAnnotations {
				if s, ok := v.(string); ok {
					m[k] = s
				}
			}
			r.Annotations = m
		} else {
			return nil, fmt.Errorf("r.Annotations: expected map[string]interface{}")
		}
	}
	if _, ok := u.Object["authorization"]; ok {
		if rAuthorization, ok := u.Object["authorization"].(map[string]interface{}); ok {
			r.Authorization = &dclService.ClusterAuthorization{}
			if _, ok := rAuthorization["adminGroups"]; ok {
				if s, ok := rAuthorization["adminGroups"].([]interface{}); ok {
					for _, o := range s {
						if objval, ok := o.(map[string]interface{}); ok {
							var rAuthorizationAdminGroups dclService.ClusterAuthorizationAdminGroups
							if _, ok := objval["group"]; ok {
								if s, ok := objval["group"].(string); ok {
									rAuthorizationAdminGroups.Group = dcl.String(s)
								} else {
									return nil, fmt.Errorf("rAuthorizationAdminGroups.Group: expected string")
								}
							}
							r.Authorization.AdminGroups = append(r.Authorization.AdminGroups, rAuthorizationAdminGroups)
						}
					}
				} else {
					return nil, fmt.Errorf("r.Authorization.AdminGroups: expected []interface{}")
				}
			}
			if _, ok := rAuthorization["adminUsers"]; ok {
				if s, ok := rAuthorization["adminUsers"].([]interface{}); ok {
					for _, o := range s {
						if objval, ok := o.(map[string]interface{}); ok {
							var rAuthorizationAdminUsers dclService.ClusterAuthorizationAdminUsers
							if _, ok := objval["username"]; ok {
								if s, ok := objval["username"].(string); ok {
									rAuthorizationAdminUsers.Username = dcl.String(s)
								} else {
									return nil, fmt.Errorf("rAuthorizationAdminUsers.Username: expected string")
								}
							}
							r.Authorization.AdminUsers = append(r.Authorization.AdminUsers, rAuthorizationAdminUsers)
						}
					}
				} else {
					return nil, fmt.Errorf("r.Authorization.AdminUsers: expected []interface{}")
				}
			}
		} else {
			return nil, fmt.Errorf("r.Authorization: expected map[string]interface{}")
		}
	}
	if _, ok := u.Object["azureRegion"]; ok {
		if s, ok := u.Object["azureRegion"].(string); ok {
			r.AzureRegion = dcl.String(s)
		} else {
			return nil, fmt.Errorf("r.AzureRegion: expected string")
		}
	}
	if _, ok := u.Object["azureServicesAuthentication"]; ok {
		if rAzureServicesAuthentication, ok := u.Object["azureServicesAuthentication"].(map[string]interface{}); ok {
			r.AzureServicesAuthentication = &dclService.ClusterAzureServicesAuthentication{}
			if _, ok := rAzureServicesAuthentication["applicationId"]; ok {
				if s, ok := rAzureServicesAuthentication["applicationId"].(string); ok {
					r.AzureServicesAuthentication.ApplicationId = dcl.String(s)
				} else {
					return nil, fmt.Errorf("r.AzureServicesAuthentication.ApplicationId: expected string")
				}
			}
			if _, ok := rAzureServicesAuthentication["tenantId"]; ok {
				if s, ok := rAzureServicesAuthentication["tenantId"].(string); ok {
					r.AzureServicesAuthentication.TenantId = dcl.String(s)
				} else {
					return nil, fmt.Errorf("r.AzureServicesAuthentication.TenantId: expected string")
				}
			}
		} else {
			return nil, fmt.Errorf("r.AzureServicesAuthentication: expected map[string]interface{}")
		}
	}
	if _, ok := u.Object["client"]; ok {
		if s, ok := u.Object["client"].(string); ok {
			r.Client = dcl.String(s)
		} else {
			return nil, fmt.Errorf("r.Client: expected string")
		}
	}
	if _, ok := u.Object["controlPlane"]; ok {
		if rControlPlane, ok := u.Object["controlPlane"].(map[string]interface{}); ok {
			r.ControlPlane = &dclService.ClusterControlPlane{}
			if _, ok := rControlPlane["databaseEncryption"]; ok {
				if rControlPlaneDatabaseEncryption, ok := rControlPlane["databaseEncryption"].(map[string]interface{}); ok {
					r.ControlPlane.DatabaseEncryption = &dclService.ClusterControlPlaneDatabaseEncryption{}
					if _, ok := rControlPlaneDatabaseEncryption["keyId"]; ok {
						if s, ok := rControlPlaneDatabaseEncryption["keyId"].(string); ok {
							r.ControlPlane.DatabaseEncryption.KeyId = dcl.String(s)
						} else {
							return nil, fmt.Errorf("r.ControlPlane.DatabaseEncryption.KeyId: expected string")
						}
					}
				} else {
					return nil, fmt.Errorf("r.ControlPlane.DatabaseEncryption: expected map[string]interface{}")
				}
			}
			if _, ok := rControlPlane["mainVolume"]; ok {
				if rControlPlaneMainVolume, ok := rControlPlane["mainVolume"].(map[string]interface{}); ok {
					r.ControlPlane.MainVolume = &dclService.ClusterControlPlaneMainVolume{}
					if _, ok := rControlPlaneMainVolume["sizeGib"]; ok {
						if i, ok := rControlPlaneMainVolume["sizeGib"].(int64); ok {
							r.ControlPlane.MainVolume.SizeGib = dcl.Int64(i)
						} else {
							return nil, fmt.Errorf("r.ControlPlane.MainVolume.SizeGib: expected int64")
						}
					}
				} else {
					return nil, fmt.Errorf("r.ControlPlane.MainVolume: expected map[string]interface{}")
				}
			}
			if _, ok := rControlPlane["proxyConfig"]; ok {
				if rControlPlaneProxyConfig, ok := rControlPlane["proxyConfig"].(map[string]interface{}); ok {
					r.ControlPlane.ProxyConfig = &dclService.ClusterControlPlaneProxyConfig{}
					if _, ok := rControlPlaneProxyConfig["resourceGroupId"]; ok {
						if s, ok := rControlPlaneProxyConfig["resourceGroupId"].(string); ok {
							r.ControlPlane.ProxyConfig.ResourceGroupId = dcl.String(s)
						} else {
							return nil, fmt.Errorf("r.ControlPlane.ProxyConfig.ResourceGroupId: expected string")
						}
					}
					if _, ok := rControlPlaneProxyConfig["secretId"]; ok {
						if s, ok := rControlPlaneProxyConfig["secretId"].(string); ok {
							r.ControlPlane.ProxyConfig.SecretId = dcl.String(s)
						} else {
							return nil, fmt.Errorf("r.ControlPlane.ProxyConfig.SecretId: expected string")
						}
					}
				} else {
					return nil, fmt.Errorf("r.ControlPlane.ProxyConfig: expected map[string]interface{}")
				}
			}
			if _, ok := rControlPlane["replicaPlacements"]; ok {
				if s, ok := rControlPlane["replicaPlacements"].([]interface{}); ok {
					for _, o := range s {
						if objval, ok := o.(map[string]interface{}); ok {
							var rControlPlaneReplicaPlacements dclService.ClusterControlPlaneReplicaPlacements
							if _, ok := objval["azureAvailabilityZone"]; ok {
								if s, ok := objval["azureAvailabilityZone"].(string); ok {
									rControlPlaneReplicaPlacements.AzureAvailabilityZone = dcl.String(s)
								} else {
									return nil, fmt.Errorf("rControlPlaneReplicaPlacements.AzureAvailabilityZone: expected string")
								}
							}
							if _, ok := objval["subnetId"]; ok {
								if s, ok := objval["subnetId"].(string); ok {
									rControlPlaneReplicaPlacements.SubnetId = dcl.String(s)
								} else {
									return nil, fmt.Errorf("rControlPlaneReplicaPlacements.SubnetId: expected string")
								}
							}
							r.ControlPlane.ReplicaPlacements = append(r.ControlPlane.ReplicaPlacements, rControlPlaneReplicaPlacements)
						}
					}
				} else {
					return nil, fmt.Errorf("r.ControlPlane.ReplicaPlacements: expected []interface{}")
				}
			}
			if _, ok := rControlPlane["rootVolume"]; ok {
				if rControlPlaneRootVolume, ok := rControlPlane["rootVolume"].(map[string]interface{}); ok {
					r.ControlPlane.RootVolume = &dclService.ClusterControlPlaneRootVolume{}
					if _, ok := rControlPlaneRootVolume["sizeGib"]; ok {
						if i, ok := rControlPlaneRootVolume["sizeGib"].(int64); ok {
							r.ControlPlane.RootVolume.SizeGib = dcl.Int64(i)
						} else {
							return nil, fmt.Errorf("r.ControlPlane.RootVolume.SizeGib: expected int64")
						}
					}
				} else {
					return nil, fmt.Errorf("r.ControlPlane.RootVolume: expected map[string]interface{}")
				}
			}
			if _, ok := rControlPlane["sshConfig"]; ok {
				if rControlPlaneSshConfig, ok := rControlPlane["sshConfig"].(map[string]interface{}); ok {
					r.ControlPlane.SshConfig = &dclService.ClusterControlPlaneSshConfig{}
					if _, ok := rControlPlaneSshConfig["authorizedKey"]; ok {
						if s, ok := rControlPlaneSshConfig["authorizedKey"].(string); ok {
							r.ControlPlane.SshConfig.AuthorizedKey = dcl.String(s)
						} else {
							return nil, fmt.Errorf("r.ControlPlane.SshConfig.AuthorizedKey: expected string")
						}
					}
				} else {
					return nil, fmt.Errorf("r.ControlPlane.SshConfig: expected map[string]interface{}")
				}
			}
			if _, ok := rControlPlane["subnetId"]; ok {
				if s, ok := rControlPlane["subnetId"].(string); ok {
					r.ControlPlane.SubnetId = dcl.String(s)
				} else {
					return nil, fmt.Errorf("r.ControlPlane.SubnetId: expected string")
				}
			}
			if _, ok := rControlPlane["tags"]; ok {
				if rControlPlaneTags, ok := rControlPlane["tags"].(map[string]interface{}); ok {
					m := make(map[string]string)
					for k, v := range rControlPlaneTags {
						if s, ok := v.(string); ok {
							m[k] = s
						}
					}
					r.ControlPlane.Tags = m
				} else {
					return nil, fmt.Errorf("r.ControlPlane.Tags: expected map[string]interface{}")
				}
			}
			if _, ok := rControlPlane["version"]; ok {
				if s, ok := rControlPlane["version"].(string); ok {
					r.ControlPlane.Version = dcl.String(s)
				} else {
					return nil, fmt.Errorf("r.ControlPlane.Version: expected string")
				}
			}
			if _, ok := rControlPlane["vmSize"]; ok {
				if s, ok := rControlPlane["vmSize"].(string); ok {
					r.ControlPlane.VmSize = dcl.String(s)
				} else {
					return nil, fmt.Errorf("r.ControlPlane.VmSize: expected string")
				}
			}
		} else {
			return nil, fmt.Errorf("r.ControlPlane: expected map[string]interface{}")
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
	if _, ok := u.Object["endpoint"]; ok {
		if s, ok := u.Object["endpoint"].(string); ok {
			r.Endpoint = dcl.String(s)
		} else {
			return nil, fmt.Errorf("r.Endpoint: expected string")
		}
	}
	if _, ok := u.Object["etag"]; ok {
		if s, ok := u.Object["etag"].(string); ok {
			r.Etag = dcl.String(s)
		} else {
			return nil, fmt.Errorf("r.Etag: expected string")
		}
	}
	if _, ok := u.Object["fleet"]; ok {
		if rFleet, ok := u.Object["fleet"].(map[string]interface{}); ok {
			r.Fleet = &dclService.ClusterFleet{}
			if _, ok := rFleet["membership"]; ok {
				if s, ok := rFleet["membership"].(string); ok {
					r.Fleet.Membership = dcl.String(s)
				} else {
					return nil, fmt.Errorf("r.Fleet.Membership: expected string")
				}
			}
			if _, ok := rFleet["project"]; ok {
				if s, ok := rFleet["project"].(string); ok {
					r.Fleet.Project = dcl.String(s)
				} else {
					return nil, fmt.Errorf("r.Fleet.Project: expected string")
				}
			}
		} else {
			return nil, fmt.Errorf("r.Fleet: expected map[string]interface{}")
		}
	}
	if _, ok := u.Object["location"]; ok {
		if s, ok := u.Object["location"].(string); ok {
			r.Location = dcl.String(s)
		} else {
			return nil, fmt.Errorf("r.Location: expected string")
		}
	}
	if _, ok := u.Object["loggingConfig"]; ok {
		if rLoggingConfig, ok := u.Object["loggingConfig"].(map[string]interface{}); ok {
			r.LoggingConfig = &dclService.ClusterLoggingConfig{}
			if _, ok := rLoggingConfig["componentConfig"]; ok {
				if rLoggingConfigComponentConfig, ok := rLoggingConfig["componentConfig"].(map[string]interface{}); ok {
					r.LoggingConfig.ComponentConfig = &dclService.ClusterLoggingConfigComponentConfig{}
					if _, ok := rLoggingConfigComponentConfig["enableComponents"]; ok {
						if s, ok := rLoggingConfigComponentConfig["enableComponents"].([]interface{}); ok {
							for _, ss := range s {
								if strval, ok := ss.(string); ok {
									r.LoggingConfig.ComponentConfig.EnableComponents = append(r.LoggingConfig.ComponentConfig.EnableComponents, dclService.ClusterLoggingConfigComponentConfigEnableComponentsEnum(strval))
								}
							}
						} else {
							return nil, fmt.Errorf("r.LoggingConfig.ComponentConfig.EnableComponents: expected []interface{}")
						}
					}
				} else {
					return nil, fmt.Errorf("r.LoggingConfig.ComponentConfig: expected map[string]interface{}")
				}
			}
		} else {
			return nil, fmt.Errorf("r.LoggingConfig: expected map[string]interface{}")
		}
	}
	if _, ok := u.Object["monitoringConfig"]; ok {
		if rMonitoringConfig, ok := u.Object["monitoringConfig"].(map[string]interface{}); ok {
			r.MonitoringConfig = &dclService.ClusterMonitoringConfig{}
			if _, ok := rMonitoringConfig["managedPrometheusConfig"]; ok {
				if rMonitoringConfigManagedPrometheusConfig, ok := rMonitoringConfig["managedPrometheusConfig"].(map[string]interface{}); ok {
					r.MonitoringConfig.ManagedPrometheusConfig = &dclService.ClusterMonitoringConfigManagedPrometheusConfig{}
					if _, ok := rMonitoringConfigManagedPrometheusConfig["enabled"]; ok {
						if b, ok := rMonitoringConfigManagedPrometheusConfig["enabled"].(bool); ok {
							r.MonitoringConfig.ManagedPrometheusConfig.Enabled = dcl.Bool(b)
						} else {
							return nil, fmt.Errorf("r.MonitoringConfig.ManagedPrometheusConfig.Enabled: expected bool")
						}
					}
				} else {
					return nil, fmt.Errorf("r.MonitoringConfig.ManagedPrometheusConfig: expected map[string]interface{}")
				}
			}
		} else {
			return nil, fmt.Errorf("r.MonitoringConfig: expected map[string]interface{}")
		}
	}
	if _, ok := u.Object["name"]; ok {
		if s, ok := u.Object["name"].(string); ok {
			r.Name = dcl.String(s)
		} else {
			return nil, fmt.Errorf("r.Name: expected string")
		}
	}
	if _, ok := u.Object["networking"]; ok {
		if rNetworking, ok := u.Object["networking"].(map[string]interface{}); ok {
			r.Networking = &dclService.ClusterNetworking{}
			if _, ok := rNetworking["podAddressCidrBlocks"]; ok {
				if s, ok := rNetworking["podAddressCidrBlocks"].([]interface{}); ok {
					for _, ss := range s {
						if strval, ok := ss.(string); ok {
							r.Networking.PodAddressCidrBlocks = append(r.Networking.PodAddressCidrBlocks, strval)
						}
					}
				} else {
					return nil, fmt.Errorf("r.Networking.PodAddressCidrBlocks: expected []interface{}")
				}
			}
			if _, ok := rNetworking["serviceAddressCidrBlocks"]; ok {
				if s, ok := rNetworking["serviceAddressCidrBlocks"].([]interface{}); ok {
					for _, ss := range s {
						if strval, ok := ss.(string); ok {
							r.Networking.ServiceAddressCidrBlocks = append(r.Networking.ServiceAddressCidrBlocks, strval)
						}
					}
				} else {
					return nil, fmt.Errorf("r.Networking.ServiceAddressCidrBlocks: expected []interface{}")
				}
			}
			if _, ok := rNetworking["virtualNetworkId"]; ok {
				if s, ok := rNetworking["virtualNetworkId"].(string); ok {
					r.Networking.VirtualNetworkId = dcl.String(s)
				} else {
					return nil, fmt.Errorf("r.Networking.VirtualNetworkId: expected string")
				}
			}
		} else {
			return nil, fmt.Errorf("r.Networking: expected map[string]interface{}")
		}
	}
	if _, ok := u.Object["project"]; ok {
		if s, ok := u.Object["project"].(string); ok {
			r.Project = dcl.String(s)
		} else {
			return nil, fmt.Errorf("r.Project: expected string")
		}
	}
	if _, ok := u.Object["reconciling"]; ok {
		if b, ok := u.Object["reconciling"].(bool); ok {
			r.Reconciling = dcl.Bool(b)
		} else {
			return nil, fmt.Errorf("r.Reconciling: expected bool")
		}
	}
	if _, ok := u.Object["resourceGroupId"]; ok {
		if s, ok := u.Object["resourceGroupId"].(string); ok {
			r.ResourceGroupId = dcl.String(s)
		} else {
			return nil, fmt.Errorf("r.ResourceGroupId: expected string")
		}
	}
	if _, ok := u.Object["state"]; ok {
		if s, ok := u.Object["state"].(string); ok {
			r.State = dclService.ClusterStateEnumRef(s)
		} else {
			return nil, fmt.Errorf("r.State: expected string")
		}
	}
	if _, ok := u.Object["uid"]; ok {
		if s, ok := u.Object["uid"].(string); ok {
			r.Uid = dcl.String(s)
		} else {
			return nil, fmt.Errorf("r.Uid: expected string")
		}
	}
	if _, ok := u.Object["updateTime"]; ok {
		if s, ok := u.Object["updateTime"].(string); ok {
			r.UpdateTime = dcl.String(s)
		} else {
			return nil, fmt.Errorf("r.UpdateTime: expected string")
		}
	}
	if _, ok := u.Object["workloadIdentityConfig"]; ok {
		if rWorkloadIdentityConfig, ok := u.Object["workloadIdentityConfig"].(map[string]interface{}); ok {
			r.WorkloadIdentityConfig = &dclService.ClusterWorkloadIdentityConfig{}
			if _, ok := rWorkloadIdentityConfig["identityProvider"]; ok {
				if s, ok := rWorkloadIdentityConfig["identityProvider"].(string); ok {
					r.WorkloadIdentityConfig.IdentityProvider = dcl.String(s)
				} else {
					return nil, fmt.Errorf("r.WorkloadIdentityConfig.IdentityProvider: expected string")
				}
			}
			if _, ok := rWorkloadIdentityConfig["issuerUri"]; ok {
				if s, ok := rWorkloadIdentityConfig["issuerUri"].(string); ok {
					r.WorkloadIdentityConfig.IssuerUri = dcl.String(s)
				} else {
					return nil, fmt.Errorf("r.WorkloadIdentityConfig.IssuerUri: expected string")
				}
			}
			if _, ok := rWorkloadIdentityConfig["workloadPool"]; ok {
				if s, ok := rWorkloadIdentityConfig["workloadPool"].(string); ok {
					r.WorkloadIdentityConfig.WorkloadPool = dcl.String(s)
				} else {
					return nil, fmt.Errorf("r.WorkloadIdentityConfig.WorkloadPool: expected string")
				}
			}
		} else {
			return nil, fmt.Errorf("r.WorkloadIdentityConfig: expected map[string]interface{}")
		}
	}
	return r, nil
}

func GetCluster(ctx context.Context, config *dcl.Config, u *unstructured.Resource) (*unstructured.Resource, error) {
	c := dclService.NewClient(config)
	r, err := UnstructuredToCluster(u)
	if err != nil {
		return nil, err
	}
	r, err = c.GetCluster(ctx, r)
	if err != nil {
		return nil, err
	}
	return ClusterToUnstructured(r), nil
}

func ListCluster(ctx context.Context, config *dcl.Config, project string, location string) ([]*unstructured.Resource, error) {
	c := dclService.NewClient(config)
	l, err := c.ListCluster(ctx, project, location)
	if err != nil {
		return nil, err
	}
	var resources []*unstructured.Resource
	for {
		for _, r := range l.Items {
			resources = append(resources, ClusterToUnstructured(r))
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

func ApplyCluster(ctx context.Context, config *dcl.Config, u *unstructured.Resource, opts ...dcl.ApplyOption) (*unstructured.Resource, error) {
	c := dclService.NewClient(config)
	r, err := UnstructuredToCluster(u)
	if err != nil {
		return nil, err
	}
	if ush := unstructured.FetchStateHint(opts); ush != nil {
		sh, err := UnstructuredToCluster(ush)
		if err != nil {
			return nil, err
		}
		opts = append(opts, dcl.WithStateHint(sh))
	}
	r, err = c.ApplyCluster(ctx, r, opts...)
	if err != nil {
		return nil, err
	}
	return ClusterToUnstructured(r), nil
}

func ClusterHasDiff(ctx context.Context, config *dcl.Config, u *unstructured.Resource, opts ...dcl.ApplyOption) (bool, error) {
	c := dclService.NewClient(config)
	r, err := UnstructuredToCluster(u)
	if err != nil {
		return false, err
	}
	if ush := unstructured.FetchStateHint(opts); ush != nil {
		sh, err := UnstructuredToCluster(ush)
		if err != nil {
			return false, err
		}
		opts = append(opts, dcl.WithStateHint(sh))
	}
	opts = append(opts, dcl.WithLifecycleParam(dcl.BlockDestruction), dcl.WithLifecycleParam(dcl.BlockCreation), dcl.WithLifecycleParam(dcl.BlockModification))
	_, err = c.ApplyCluster(ctx, r, opts...)
	if err != nil {
		if _, ok := err.(dcl.ApplyInfeasibleError); ok {
			return true, nil
		}
		return false, err
	}
	return false, nil
}

func DeleteCluster(ctx context.Context, config *dcl.Config, u *unstructured.Resource) error {
	c := dclService.NewClient(config)
	r, err := UnstructuredToCluster(u)
	if err != nil {
		return err
	}
	return c.DeleteCluster(ctx, r)
}

func ClusterID(u *unstructured.Resource) (string, error) {
	r, err := UnstructuredToCluster(u)
	if err != nil {
		return "", err
	}
	return r.ID()
}

func (r *Cluster) STV() unstructured.ServiceTypeVersion {
	return unstructured.ServiceTypeVersion{
		"containerazure",
		"Cluster",
		"beta",
	}
}

func (r *Cluster) SetPolicyMember(ctx context.Context, config *dcl.Config, resource *unstructured.Resource, member *unstructured.Resource) (*unstructured.Resource, error) {
	return nil, unstructured.ErrNoSuchMethod
}

func (r *Cluster) GetPolicyMember(ctx context.Context, config *dcl.Config, resource *unstructured.Resource, role, member string) (*unstructured.Resource, error) {
	return nil, unstructured.ErrNoSuchMethod
}

func (r *Cluster) DeletePolicyMember(ctx context.Context, config *dcl.Config, resource *unstructured.Resource, member *unstructured.Resource) error {
	return unstructured.ErrNoSuchMethod
}

func (r *Cluster) SetPolicy(ctx context.Context, config *dcl.Config, resource *unstructured.Resource, policy *unstructured.Resource) (*unstructured.Resource, error) {
	return nil, unstructured.ErrNoSuchMethod
}

func (r *Cluster) SetPolicyWithEtag(ctx context.Context, config *dcl.Config, resource *unstructured.Resource, policy *unstructured.Resource) (*unstructured.Resource, error) {
	return nil, unstructured.ErrNoSuchMethod
}

func (r *Cluster) GetPolicy(ctx context.Context, config *dcl.Config, resource *unstructured.Resource) (*unstructured.Resource, error) {
	return nil, unstructured.ErrNoSuchMethod
}

func (r *Cluster) Get(ctx context.Context, config *dcl.Config, resource *unstructured.Resource) (*unstructured.Resource, error) {
	return GetCluster(ctx, config, resource)
}

func (r *Cluster) Apply(ctx context.Context, config *dcl.Config, resource *unstructured.Resource, opts ...dcl.ApplyOption) (*unstructured.Resource, error) {
	return ApplyCluster(ctx, config, resource, opts...)
}

func (r *Cluster) HasDiff(ctx context.Context, config *dcl.Config, resource *unstructured.Resource, opts ...dcl.ApplyOption) (bool, error) {
	return ClusterHasDiff(ctx, config, resource, opts...)
}

func (r *Cluster) Delete(ctx context.Context, config *dcl.Config, resource *unstructured.Resource) error {
	return DeleteCluster(ctx, config, resource)
}

func (r *Cluster) ID(resource *unstructured.Resource) (string, error) {
	return ClusterID(resource)
}

func init() {
	unstructured.Register(&Cluster{})
}
