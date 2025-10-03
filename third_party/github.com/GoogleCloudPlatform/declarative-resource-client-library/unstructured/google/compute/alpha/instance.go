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
package compute

import (
	"context"
	"fmt"
	"github.com/GoogleCloudPlatform/declarative-resource-client-library/dcl"
	dclService "github.com/GoogleCloudPlatform/declarative-resource-client-library/services/google/compute/alpha"
	"github.com/GoogleCloudPlatform/declarative-resource-client-library/services/google/iam"
	"github.com/GoogleCloudPlatform/declarative-resource-client-library/unstructured"
	iamUnstruct "github.com/GoogleCloudPlatform/declarative-resource-client-library/unstructured/google/iam"
)

type Instance struct{}

func InstanceToUnstructured(r *dclService.Instance) *unstructured.Resource {
	u := &unstructured.Resource{
		STV: unstructured.ServiceTypeVersion{
			Service: "compute",
			Version: "alpha",
			Type:    "Instance",
		},
		Object: make(map[string]interface{}),
	}
	if r.CanIPForward != nil {
		u.Object["canIPForward"] = *r.CanIPForward
	}
	if r.CpuPlatform != nil {
		u.Object["cpuPlatform"] = *r.CpuPlatform
	}
	if r.CreationTimestamp != nil {
		u.Object["creationTimestamp"] = *r.CreationTimestamp
	}
	if r.DeletionProtection != nil {
		u.Object["deletionProtection"] = *r.DeletionProtection
	}
	if r.Description != nil {
		u.Object["description"] = *r.Description
	}
	var rDisks []interface{}
	for _, rDisksVal := range r.Disks {
		rDisksObject := make(map[string]interface{})
		if rDisksVal.AutoDelete != nil {
			rDisksObject["autoDelete"] = *rDisksVal.AutoDelete
		}
		if rDisksVal.Boot != nil {
			rDisksObject["boot"] = *rDisksVal.Boot
		}
		if rDisksVal.DeviceName != nil {
			rDisksObject["deviceName"] = *rDisksVal.DeviceName
		}
		if rDisksVal.DiskEncryptionKey != nil && rDisksVal.DiskEncryptionKey != dclService.EmptyInstanceDisksDiskEncryptionKey {
			rDisksValDiskEncryptionKey := make(map[string]interface{})
			if rDisksVal.DiskEncryptionKey.RawKey != nil {
				rDisksValDiskEncryptionKey["rawKey"] = *rDisksVal.DiskEncryptionKey.RawKey
			}
			if rDisksVal.DiskEncryptionKey.RsaEncryptedKey != nil {
				rDisksValDiskEncryptionKey["rsaEncryptedKey"] = *rDisksVal.DiskEncryptionKey.RsaEncryptedKey
			}
			if rDisksVal.DiskEncryptionKey.Sha256 != nil {
				rDisksValDiskEncryptionKey["sha256"] = *rDisksVal.DiskEncryptionKey.Sha256
			}
			rDisksObject["diskEncryptionKey"] = rDisksValDiskEncryptionKey
		}
		if rDisksVal.Index != nil {
			rDisksObject["index"] = *rDisksVal.Index
		}
		if rDisksVal.InitializeParams != nil && rDisksVal.InitializeParams != dclService.EmptyInstanceDisksInitializeParams {
			rDisksValInitializeParams := make(map[string]interface{})
			if rDisksVal.InitializeParams.DiskName != nil {
				rDisksValInitializeParams["diskName"] = *rDisksVal.InitializeParams.DiskName
			}
			if rDisksVal.InitializeParams.DiskSizeGb != nil {
				rDisksValInitializeParams["diskSizeGb"] = *rDisksVal.InitializeParams.DiskSizeGb
			}
			if rDisksVal.InitializeParams.DiskType != nil {
				rDisksValInitializeParams["diskType"] = *rDisksVal.InitializeParams.DiskType
			}
			if rDisksVal.InitializeParams.SourceImage != nil {
				rDisksValInitializeParams["sourceImage"] = *rDisksVal.InitializeParams.SourceImage
			}
			if rDisksVal.InitializeParams.SourceImageEncryptionKey != nil && rDisksVal.InitializeParams.SourceImageEncryptionKey != dclService.EmptyInstanceDisksInitializeParamsSourceImageEncryptionKey {
				rDisksValInitializeParamsSourceImageEncryptionKey := make(map[string]interface{})
				if rDisksVal.InitializeParams.SourceImageEncryptionKey.RawKey != nil {
					rDisksValInitializeParamsSourceImageEncryptionKey["rawKey"] = *rDisksVal.InitializeParams.SourceImageEncryptionKey.RawKey
				}
				if rDisksVal.InitializeParams.SourceImageEncryptionKey.Sha256 != nil {
					rDisksValInitializeParamsSourceImageEncryptionKey["sha256"] = *rDisksVal.InitializeParams.SourceImageEncryptionKey.Sha256
				}
				rDisksValInitializeParams["sourceImageEncryptionKey"] = rDisksValInitializeParamsSourceImageEncryptionKey
			}
			rDisksObject["initializeParams"] = rDisksValInitializeParams
		}
		if rDisksVal.Interface != nil {
			rDisksObject["interface"] = string(*rDisksVal.Interface)
		}
		if rDisksVal.Mode != nil {
			rDisksObject["mode"] = string(*rDisksVal.Mode)
		}
		if rDisksVal.Source != nil {
			rDisksObject["source"] = *rDisksVal.Source
		}
		if rDisksVal.Type != nil {
			rDisksObject["type"] = string(*rDisksVal.Type)
		}
		rDisks = append(rDisks, rDisksObject)
	}
	u.Object["disks"] = rDisks
	var rGuestAccelerators []interface{}
	for _, rGuestAcceleratorsVal := range r.GuestAccelerators {
		rGuestAcceleratorsObject := make(map[string]interface{})
		if rGuestAcceleratorsVal.AcceleratorCount != nil {
			rGuestAcceleratorsObject["acceleratorCount"] = *rGuestAcceleratorsVal.AcceleratorCount
		}
		if rGuestAcceleratorsVal.AcceleratorType != nil {
			rGuestAcceleratorsObject["acceleratorType"] = *rGuestAcceleratorsVal.AcceleratorType
		}
		rGuestAccelerators = append(rGuestAccelerators, rGuestAcceleratorsObject)
	}
	u.Object["guestAccelerators"] = rGuestAccelerators
	if r.Hostname != nil {
		u.Object["hostname"] = *r.Hostname
	}
	if r.Id != nil {
		u.Object["id"] = *r.Id
	}
	if r.Labels != nil {
		rLabels := make(map[string]interface{})
		for k, v := range r.Labels {
			rLabels[k] = v
		}
		u.Object["labels"] = rLabels
	}
	if r.MachineType != nil {
		u.Object["machineType"] = *r.MachineType
	}
	if r.Metadata != nil {
		rMetadata := make(map[string]interface{})
		for k, v := range r.Metadata {
			rMetadata[k] = v
		}
		u.Object["metadata"] = rMetadata
	}
	if r.MinCpuPlatform != nil {
		u.Object["minCpuPlatform"] = *r.MinCpuPlatform
	}
	if r.Name != nil {
		u.Object["name"] = *r.Name
	}
	var rNetworkInterfaces []interface{}
	for _, rNetworkInterfacesVal := range r.NetworkInterfaces {
		rNetworkInterfacesObject := make(map[string]interface{})
		var rNetworkInterfacesValAccessConfigs []interface{}
		for _, rNetworkInterfacesValAccessConfigsVal := range rNetworkInterfacesVal.AccessConfigs {
			rNetworkInterfacesValAccessConfigsObject := make(map[string]interface{})
			if rNetworkInterfacesValAccessConfigsVal.ExternalIPv6 != nil {
				rNetworkInterfacesValAccessConfigsObject["externalIPv6"] = *rNetworkInterfacesValAccessConfigsVal.ExternalIPv6
			}
			if rNetworkInterfacesValAccessConfigsVal.ExternalIPv6PrefixLength != nil {
				rNetworkInterfacesValAccessConfigsObject["externalIPv6PrefixLength"] = *rNetworkInterfacesValAccessConfigsVal.ExternalIPv6PrefixLength
			}
			if rNetworkInterfacesValAccessConfigsVal.Name != nil {
				rNetworkInterfacesValAccessConfigsObject["name"] = *rNetworkInterfacesValAccessConfigsVal.Name
			}
			if rNetworkInterfacesValAccessConfigsVal.NatIP != nil {
				rNetworkInterfacesValAccessConfigsObject["natIP"] = *rNetworkInterfacesValAccessConfigsVal.NatIP
			}
			if rNetworkInterfacesValAccessConfigsVal.NetworkTier != nil {
				rNetworkInterfacesValAccessConfigsObject["networkTier"] = string(*rNetworkInterfacesValAccessConfigsVal.NetworkTier)
			}
			if rNetworkInterfacesValAccessConfigsVal.PublicPtrDomainName != nil {
				rNetworkInterfacesValAccessConfigsObject["publicPtrDomainName"] = *rNetworkInterfacesValAccessConfigsVal.PublicPtrDomainName
			}
			if rNetworkInterfacesValAccessConfigsVal.SetPublicPtr != nil {
				rNetworkInterfacesValAccessConfigsObject["setPublicPtr"] = *rNetworkInterfacesValAccessConfigsVal.SetPublicPtr
			}
			if rNetworkInterfacesValAccessConfigsVal.Type != nil {
				rNetworkInterfacesValAccessConfigsObject["type"] = string(*rNetworkInterfacesValAccessConfigsVal.Type)
			}
			rNetworkInterfacesValAccessConfigs = append(rNetworkInterfacesValAccessConfigs, rNetworkInterfacesValAccessConfigsObject)
		}
		rNetworkInterfacesObject["accessConfigs"] = rNetworkInterfacesValAccessConfigs
		var rNetworkInterfacesValAliasIPRanges []interface{}
		for _, rNetworkInterfacesValAliasIPRangesVal := range rNetworkInterfacesVal.AliasIPRanges {
			rNetworkInterfacesValAliasIPRangesObject := make(map[string]interface{})
			if rNetworkInterfacesValAliasIPRangesVal.IPCidrRange != nil {
				rNetworkInterfacesValAliasIPRangesObject["ipCidrRange"] = *rNetworkInterfacesValAliasIPRangesVal.IPCidrRange
			}
			if rNetworkInterfacesValAliasIPRangesVal.SubnetworkRangeName != nil {
				rNetworkInterfacesValAliasIPRangesObject["subnetworkRangeName"] = *rNetworkInterfacesValAliasIPRangesVal.SubnetworkRangeName
			}
			rNetworkInterfacesValAliasIPRanges = append(rNetworkInterfacesValAliasIPRanges, rNetworkInterfacesValAliasIPRangesObject)
		}
		rNetworkInterfacesObject["aliasIPRanges"] = rNetworkInterfacesValAliasIPRanges
		var rNetworkInterfacesValIPv6AccessConfigs []interface{}
		for _, rNetworkInterfacesValIPv6AccessConfigsVal := range rNetworkInterfacesVal.IPv6AccessConfigs {
			rNetworkInterfacesValIPv6AccessConfigsObject := make(map[string]interface{})
			if rNetworkInterfacesValIPv6AccessConfigsVal.ExternalIPv6 != nil {
				rNetworkInterfacesValIPv6AccessConfigsObject["externalIPv6"] = *rNetworkInterfacesValIPv6AccessConfigsVal.ExternalIPv6
			}
			if rNetworkInterfacesValIPv6AccessConfigsVal.ExternalIPv6PrefixLength != nil {
				rNetworkInterfacesValIPv6AccessConfigsObject["externalIPv6PrefixLength"] = *rNetworkInterfacesValIPv6AccessConfigsVal.ExternalIPv6PrefixLength
			}
			if rNetworkInterfacesValIPv6AccessConfigsVal.Name != nil {
				rNetworkInterfacesValIPv6AccessConfigsObject["name"] = *rNetworkInterfacesValIPv6AccessConfigsVal.Name
			}
			if rNetworkInterfacesValIPv6AccessConfigsVal.NatIP != nil {
				rNetworkInterfacesValIPv6AccessConfigsObject["natIP"] = *rNetworkInterfacesValIPv6AccessConfigsVal.NatIP
			}
			if rNetworkInterfacesValIPv6AccessConfigsVal.NetworkTier != nil {
				rNetworkInterfacesValIPv6AccessConfigsObject["networkTier"] = string(*rNetworkInterfacesValIPv6AccessConfigsVal.NetworkTier)
			}
			if rNetworkInterfacesValIPv6AccessConfigsVal.PublicPtrDomainName != nil {
				rNetworkInterfacesValIPv6AccessConfigsObject["publicPtrDomainName"] = *rNetworkInterfacesValIPv6AccessConfigsVal.PublicPtrDomainName
			}
			if rNetworkInterfacesValIPv6AccessConfigsVal.SetPublicPtr != nil {
				rNetworkInterfacesValIPv6AccessConfigsObject["setPublicPtr"] = *rNetworkInterfacesValIPv6AccessConfigsVal.SetPublicPtr
			}
			if rNetworkInterfacesValIPv6AccessConfigsVal.Type != nil {
				rNetworkInterfacesValIPv6AccessConfigsObject["type"] = string(*rNetworkInterfacesValIPv6AccessConfigsVal.Type)
			}
			rNetworkInterfacesValIPv6AccessConfigs = append(rNetworkInterfacesValIPv6AccessConfigs, rNetworkInterfacesValIPv6AccessConfigsObject)
		}
		rNetworkInterfacesObject["ipv6AccessConfigs"] = rNetworkInterfacesValIPv6AccessConfigs
		if rNetworkInterfacesVal.Name != nil {
			rNetworkInterfacesObject["name"] = *rNetworkInterfacesVal.Name
		}
		if rNetworkInterfacesVal.Network != nil {
			rNetworkInterfacesObject["network"] = *rNetworkInterfacesVal.Network
		}
		if rNetworkInterfacesVal.NetworkIP != nil {
			rNetworkInterfacesObject["networkIP"] = *rNetworkInterfacesVal.NetworkIP
		}
		if rNetworkInterfacesVal.Subnetwork != nil {
			rNetworkInterfacesObject["subnetwork"] = *rNetworkInterfacesVal.Subnetwork
		}
		rNetworkInterfaces = append(rNetworkInterfaces, rNetworkInterfacesObject)
	}
	u.Object["networkInterfaces"] = rNetworkInterfaces
	if r.Project != nil {
		u.Object["project"] = *r.Project
	}
	if r.Scheduling != nil && r.Scheduling != dclService.EmptyInstanceScheduling {
		rScheduling := make(map[string]interface{})
		if r.Scheduling.AutomaticRestart != nil {
			rScheduling["automaticRestart"] = *r.Scheduling.AutomaticRestart
		}
		if r.Scheduling.OnHostMaintenance != nil {
			rScheduling["onHostMaintenance"] = *r.Scheduling.OnHostMaintenance
		}
		if r.Scheduling.Preemptible != nil {
			rScheduling["preemptible"] = *r.Scheduling.Preemptible
		}
		u.Object["scheduling"] = rScheduling
	}
	if r.SelfLink != nil {
		u.Object["selfLink"] = *r.SelfLink
	}
	var rServiceAccounts []interface{}
	for _, rServiceAccountsVal := range r.ServiceAccounts {
		rServiceAccountsObject := make(map[string]interface{})
		if rServiceAccountsVal.Email != nil {
			rServiceAccountsObject["email"] = *rServiceAccountsVal.Email
		}
		var rServiceAccountsValScopes []interface{}
		for _, rServiceAccountsValScopesVal := range rServiceAccountsVal.Scopes {
			rServiceAccountsValScopes = append(rServiceAccountsValScopes, rServiceAccountsValScopesVal)
		}
		rServiceAccountsObject["scopes"] = rServiceAccountsValScopes
		rServiceAccounts = append(rServiceAccounts, rServiceAccountsObject)
	}
	u.Object["serviceAccounts"] = rServiceAccounts
	if r.ShieldedInstanceConfig != nil && r.ShieldedInstanceConfig != dclService.EmptyInstanceShieldedInstanceConfig {
		rShieldedInstanceConfig := make(map[string]interface{})
		if r.ShieldedInstanceConfig.EnableIntegrityMonitoring != nil {
			rShieldedInstanceConfig["enableIntegrityMonitoring"] = *r.ShieldedInstanceConfig.EnableIntegrityMonitoring
		}
		if r.ShieldedInstanceConfig.EnableSecureBoot != nil {
			rShieldedInstanceConfig["enableSecureBoot"] = *r.ShieldedInstanceConfig.EnableSecureBoot
		}
		if r.ShieldedInstanceConfig.EnableVtpm != nil {
			rShieldedInstanceConfig["enableVtpm"] = *r.ShieldedInstanceConfig.EnableVtpm
		}
		u.Object["shieldedInstanceConfig"] = rShieldedInstanceConfig
	}
	if r.Status != nil {
		u.Object["status"] = string(*r.Status)
	}
	if r.StatusMessage != nil {
		u.Object["statusMessage"] = *r.StatusMessage
	}
	var rTags []interface{}
	for _, rTagsVal := range r.Tags {
		rTags = append(rTags, rTagsVal)
	}
	u.Object["tags"] = rTags
	if r.Zone != nil {
		u.Object["zone"] = *r.Zone
	}
	return u
}

func UnstructuredToInstance(u *unstructured.Resource) (*dclService.Instance, error) {
	r := &dclService.Instance{}
	if _, ok := u.Object["canIPForward"]; ok {
		if b, ok := u.Object["canIPForward"].(bool); ok {
			r.CanIPForward = dcl.Bool(b)
		} else {
			return nil, fmt.Errorf("r.CanIPForward: expected bool")
		}
	}
	if _, ok := u.Object["cpuPlatform"]; ok {
		if s, ok := u.Object["cpuPlatform"].(string); ok {
			r.CpuPlatform = dcl.String(s)
		} else {
			return nil, fmt.Errorf("r.CpuPlatform: expected string")
		}
	}
	if _, ok := u.Object["creationTimestamp"]; ok {
		if s, ok := u.Object["creationTimestamp"].(string); ok {
			r.CreationTimestamp = dcl.String(s)
		} else {
			return nil, fmt.Errorf("r.CreationTimestamp: expected string")
		}
	}
	if _, ok := u.Object["deletionProtection"]; ok {
		if b, ok := u.Object["deletionProtection"].(bool); ok {
			r.DeletionProtection = dcl.Bool(b)
		} else {
			return nil, fmt.Errorf("r.DeletionProtection: expected bool")
		}
	}
	if _, ok := u.Object["description"]; ok {
		if s, ok := u.Object["description"].(string); ok {
			r.Description = dcl.String(s)
		} else {
			return nil, fmt.Errorf("r.Description: expected string")
		}
	}
	if _, ok := u.Object["disks"]; ok {
		if s, ok := u.Object["disks"].([]interface{}); ok {
			for _, o := range s {
				if objval, ok := o.(map[string]interface{}); ok {
					var rDisks dclService.InstanceDisks
					if _, ok := objval["autoDelete"]; ok {
						if b, ok := objval["autoDelete"].(bool); ok {
							rDisks.AutoDelete = dcl.Bool(b)
						} else {
							return nil, fmt.Errorf("rDisks.AutoDelete: expected bool")
						}
					}
					if _, ok := objval["boot"]; ok {
						if b, ok := objval["boot"].(bool); ok {
							rDisks.Boot = dcl.Bool(b)
						} else {
							return nil, fmt.Errorf("rDisks.Boot: expected bool")
						}
					}
					if _, ok := objval["deviceName"]; ok {
						if s, ok := objval["deviceName"].(string); ok {
							rDisks.DeviceName = dcl.String(s)
						} else {
							return nil, fmt.Errorf("rDisks.DeviceName: expected string")
						}
					}
					if _, ok := objval["diskEncryptionKey"]; ok {
						if rDisksDiskEncryptionKey, ok := objval["diskEncryptionKey"].(map[string]interface{}); ok {
							rDisks.DiskEncryptionKey = &dclService.InstanceDisksDiskEncryptionKey{}
							if _, ok := rDisksDiskEncryptionKey["rawKey"]; ok {
								if s, ok := rDisksDiskEncryptionKey["rawKey"].(string); ok {
									rDisks.DiskEncryptionKey.RawKey = dcl.String(s)
								} else {
									return nil, fmt.Errorf("rDisks.DiskEncryptionKey.RawKey: expected string")
								}
							}
							if _, ok := rDisksDiskEncryptionKey["rsaEncryptedKey"]; ok {
								if s, ok := rDisksDiskEncryptionKey["rsaEncryptedKey"].(string); ok {
									rDisks.DiskEncryptionKey.RsaEncryptedKey = dcl.String(s)
								} else {
									return nil, fmt.Errorf("rDisks.DiskEncryptionKey.RsaEncryptedKey: expected string")
								}
							}
							if _, ok := rDisksDiskEncryptionKey["sha256"]; ok {
								if s, ok := rDisksDiskEncryptionKey["sha256"].(string); ok {
									rDisks.DiskEncryptionKey.Sha256 = dcl.String(s)
								} else {
									return nil, fmt.Errorf("rDisks.DiskEncryptionKey.Sha256: expected string")
								}
							}
						} else {
							return nil, fmt.Errorf("rDisks.DiskEncryptionKey: expected map[string]interface{}")
						}
					}
					if _, ok := objval["index"]; ok {
						if i, ok := objval["index"].(int64); ok {
							rDisks.Index = dcl.Int64(i)
						} else {
							return nil, fmt.Errorf("rDisks.Index: expected int64")
						}
					}
					if _, ok := objval["initializeParams"]; ok {
						if rDisksInitializeParams, ok := objval["initializeParams"].(map[string]interface{}); ok {
							rDisks.InitializeParams = &dclService.InstanceDisksInitializeParams{}
							if _, ok := rDisksInitializeParams["diskName"]; ok {
								if s, ok := rDisksInitializeParams["diskName"].(string); ok {
									rDisks.InitializeParams.DiskName = dcl.String(s)
								} else {
									return nil, fmt.Errorf("rDisks.InitializeParams.DiskName: expected string")
								}
							}
							if _, ok := rDisksInitializeParams["diskSizeGb"]; ok {
								if i, ok := rDisksInitializeParams["diskSizeGb"].(int64); ok {
									rDisks.InitializeParams.DiskSizeGb = dcl.Int64(i)
								} else {
									return nil, fmt.Errorf("rDisks.InitializeParams.DiskSizeGb: expected int64")
								}
							}
							if _, ok := rDisksInitializeParams["diskType"]; ok {
								if s, ok := rDisksInitializeParams["diskType"].(string); ok {
									rDisks.InitializeParams.DiskType = dcl.String(s)
								} else {
									return nil, fmt.Errorf("rDisks.InitializeParams.DiskType: expected string")
								}
							}
							if _, ok := rDisksInitializeParams["sourceImage"]; ok {
								if s, ok := rDisksInitializeParams["sourceImage"].(string); ok {
									rDisks.InitializeParams.SourceImage = dcl.String(s)
								} else {
									return nil, fmt.Errorf("rDisks.InitializeParams.SourceImage: expected string")
								}
							}
							if _, ok := rDisksInitializeParams["sourceImageEncryptionKey"]; ok {
								if rDisksInitializeParamsSourceImageEncryptionKey, ok := rDisksInitializeParams["sourceImageEncryptionKey"].(map[string]interface{}); ok {
									rDisks.InitializeParams.SourceImageEncryptionKey = &dclService.InstanceDisksInitializeParamsSourceImageEncryptionKey{}
									if _, ok := rDisksInitializeParamsSourceImageEncryptionKey["rawKey"]; ok {
										if s, ok := rDisksInitializeParamsSourceImageEncryptionKey["rawKey"].(string); ok {
											rDisks.InitializeParams.SourceImageEncryptionKey.RawKey = dcl.String(s)
										} else {
											return nil, fmt.Errorf("rDisks.InitializeParams.SourceImageEncryptionKey.RawKey: expected string")
										}
									}
									if _, ok := rDisksInitializeParamsSourceImageEncryptionKey["sha256"]; ok {
										if s, ok := rDisksInitializeParamsSourceImageEncryptionKey["sha256"].(string); ok {
											rDisks.InitializeParams.SourceImageEncryptionKey.Sha256 = dcl.String(s)
										} else {
											return nil, fmt.Errorf("rDisks.InitializeParams.SourceImageEncryptionKey.Sha256: expected string")
										}
									}
								} else {
									return nil, fmt.Errorf("rDisks.InitializeParams.SourceImageEncryptionKey: expected map[string]interface{}")
								}
							}
						} else {
							return nil, fmt.Errorf("rDisks.InitializeParams: expected map[string]interface{}")
						}
					}
					if _, ok := objval["interface"]; ok {
						if s, ok := objval["interface"].(string); ok {
							rDisks.Interface = dclService.InstanceDisksInterfaceEnumRef(s)
						} else {
							return nil, fmt.Errorf("rDisks.Interface: expected string")
						}
					}
					if _, ok := objval["mode"]; ok {
						if s, ok := objval["mode"].(string); ok {
							rDisks.Mode = dclService.InstanceDisksModeEnumRef(s)
						} else {
							return nil, fmt.Errorf("rDisks.Mode: expected string")
						}
					}
					if _, ok := objval["source"]; ok {
						if s, ok := objval["source"].(string); ok {
							rDisks.Source = dcl.String(s)
						} else {
							return nil, fmt.Errorf("rDisks.Source: expected string")
						}
					}
					if _, ok := objval["type"]; ok {
						if s, ok := objval["type"].(string); ok {
							rDisks.Type = dclService.InstanceDisksTypeEnumRef(s)
						} else {
							return nil, fmt.Errorf("rDisks.Type: expected string")
						}
					}
					r.Disks = append(r.Disks, rDisks)
				}
			}
		} else {
			return nil, fmt.Errorf("r.Disks: expected []interface{}")
		}
	}
	if _, ok := u.Object["guestAccelerators"]; ok {
		if s, ok := u.Object["guestAccelerators"].([]interface{}); ok {
			for _, o := range s {
				if objval, ok := o.(map[string]interface{}); ok {
					var rGuestAccelerators dclService.InstanceGuestAccelerators
					if _, ok := objval["acceleratorCount"]; ok {
						if i, ok := objval["acceleratorCount"].(int64); ok {
							rGuestAccelerators.AcceleratorCount = dcl.Int64(i)
						} else {
							return nil, fmt.Errorf("rGuestAccelerators.AcceleratorCount: expected int64")
						}
					}
					if _, ok := objval["acceleratorType"]; ok {
						if s, ok := objval["acceleratorType"].(string); ok {
							rGuestAccelerators.AcceleratorType = dcl.String(s)
						} else {
							return nil, fmt.Errorf("rGuestAccelerators.AcceleratorType: expected string")
						}
					}
					r.GuestAccelerators = append(r.GuestAccelerators, rGuestAccelerators)
				}
			}
		} else {
			return nil, fmt.Errorf("r.GuestAccelerators: expected []interface{}")
		}
	}
	if _, ok := u.Object["hostname"]; ok {
		if s, ok := u.Object["hostname"].(string); ok {
			r.Hostname = dcl.String(s)
		} else {
			return nil, fmt.Errorf("r.Hostname: expected string")
		}
	}
	if _, ok := u.Object["id"]; ok {
		if s, ok := u.Object["id"].(string); ok {
			r.Id = dcl.String(s)
		} else {
			return nil, fmt.Errorf("r.Id: expected string")
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
	if _, ok := u.Object["machineType"]; ok {
		if s, ok := u.Object["machineType"].(string); ok {
			r.MachineType = dcl.String(s)
		} else {
			return nil, fmt.Errorf("r.MachineType: expected string")
		}
	}
	if _, ok := u.Object["metadata"]; ok {
		if rMetadata, ok := u.Object["metadata"].(map[string]interface{}); ok {
			m := make(map[string]string)
			for k, v := range rMetadata {
				if s, ok := v.(string); ok {
					m[k] = s
				}
			}
			r.Metadata = m
		} else {
			return nil, fmt.Errorf("r.Metadata: expected map[string]interface{}")
		}
	}
	if _, ok := u.Object["minCpuPlatform"]; ok {
		if s, ok := u.Object["minCpuPlatform"].(string); ok {
			r.MinCpuPlatform = dcl.String(s)
		} else {
			return nil, fmt.Errorf("r.MinCpuPlatform: expected string")
		}
	}
	if _, ok := u.Object["name"]; ok {
		if s, ok := u.Object["name"].(string); ok {
			r.Name = dcl.String(s)
		} else {
			return nil, fmt.Errorf("r.Name: expected string")
		}
	}
	if _, ok := u.Object["networkInterfaces"]; ok {
		if s, ok := u.Object["networkInterfaces"].([]interface{}); ok {
			for _, o := range s {
				if objval, ok := o.(map[string]interface{}); ok {
					var rNetworkInterfaces dclService.InstanceNetworkInterfaces
					if _, ok := objval["accessConfigs"]; ok {
						if s, ok := objval["accessConfigs"].([]interface{}); ok {
							for _, o := range s {
								if objval, ok := o.(map[string]interface{}); ok {
									var rNetworkInterfacesAccessConfigs dclService.InstanceNetworkInterfacesAccessConfigs
									if _, ok := objval["externalIPv6"]; ok {
										if s, ok := objval["externalIPv6"].(string); ok {
											rNetworkInterfacesAccessConfigs.ExternalIPv6 = dcl.String(s)
										} else {
											return nil, fmt.Errorf("rNetworkInterfacesAccessConfigs.ExternalIPv6: expected string")
										}
									}
									if _, ok := objval["externalIPv6PrefixLength"]; ok {
										if s, ok := objval["externalIPv6PrefixLength"].(string); ok {
											rNetworkInterfacesAccessConfigs.ExternalIPv6PrefixLength = dcl.String(s)
										} else {
											return nil, fmt.Errorf("rNetworkInterfacesAccessConfigs.ExternalIPv6PrefixLength: expected string")
										}
									}
									if _, ok := objval["name"]; ok {
										if s, ok := objval["name"].(string); ok {
											rNetworkInterfacesAccessConfigs.Name = dcl.String(s)
										} else {
											return nil, fmt.Errorf("rNetworkInterfacesAccessConfigs.Name: expected string")
										}
									}
									if _, ok := objval["natIP"]; ok {
										if s, ok := objval["natIP"].(string); ok {
											rNetworkInterfacesAccessConfigs.NatIP = dcl.String(s)
										} else {
											return nil, fmt.Errorf("rNetworkInterfacesAccessConfigs.NatIP: expected string")
										}
									}
									if _, ok := objval["networkTier"]; ok {
										if s, ok := objval["networkTier"].(string); ok {
											rNetworkInterfacesAccessConfigs.NetworkTier = dclService.InstanceNetworkInterfacesAccessConfigsNetworkTierEnumRef(s)
										} else {
											return nil, fmt.Errorf("rNetworkInterfacesAccessConfigs.NetworkTier: expected string")
										}
									}
									if _, ok := objval["publicPtrDomainName"]; ok {
										if s, ok := objval["publicPtrDomainName"].(string); ok {
											rNetworkInterfacesAccessConfigs.PublicPtrDomainName = dcl.String(s)
										} else {
											return nil, fmt.Errorf("rNetworkInterfacesAccessConfigs.PublicPtrDomainName: expected string")
										}
									}
									if _, ok := objval["setPublicPtr"]; ok {
										if b, ok := objval["setPublicPtr"].(bool); ok {
											rNetworkInterfacesAccessConfigs.SetPublicPtr = dcl.Bool(b)
										} else {
											return nil, fmt.Errorf("rNetworkInterfacesAccessConfigs.SetPublicPtr: expected bool")
										}
									}
									if _, ok := objval["type"]; ok {
										if s, ok := objval["type"].(string); ok {
											rNetworkInterfacesAccessConfigs.Type = dclService.InstanceNetworkInterfacesAccessConfigsTypeEnumRef(s)
										} else {
											return nil, fmt.Errorf("rNetworkInterfacesAccessConfigs.Type: expected string")
										}
									}
									rNetworkInterfaces.AccessConfigs = append(rNetworkInterfaces.AccessConfigs, rNetworkInterfacesAccessConfigs)
								}
							}
						} else {
							return nil, fmt.Errorf("rNetworkInterfaces.AccessConfigs: expected []interface{}")
						}
					}
					if _, ok := objval["aliasIPRanges"]; ok {
						if s, ok := objval["aliasIPRanges"].([]interface{}); ok {
							for _, o := range s {
								if objval, ok := o.(map[string]interface{}); ok {
									var rNetworkInterfacesAliasIPRanges dclService.InstanceNetworkInterfacesAliasIPRanges
									if _, ok := objval["ipCidrRange"]; ok {
										if s, ok := objval["ipCidrRange"].(string); ok {
											rNetworkInterfacesAliasIPRanges.IPCidrRange = dcl.String(s)
										} else {
											return nil, fmt.Errorf("rNetworkInterfacesAliasIPRanges.IPCidrRange: expected string")
										}
									}
									if _, ok := objval["subnetworkRangeName"]; ok {
										if s, ok := objval["subnetworkRangeName"].(string); ok {
											rNetworkInterfacesAliasIPRanges.SubnetworkRangeName = dcl.String(s)
										} else {
											return nil, fmt.Errorf("rNetworkInterfacesAliasIPRanges.SubnetworkRangeName: expected string")
										}
									}
									rNetworkInterfaces.AliasIPRanges = append(rNetworkInterfaces.AliasIPRanges, rNetworkInterfacesAliasIPRanges)
								}
							}
						} else {
							return nil, fmt.Errorf("rNetworkInterfaces.AliasIPRanges: expected []interface{}")
						}
					}
					if _, ok := objval["ipv6AccessConfigs"]; ok {
						if s, ok := objval["ipv6AccessConfigs"].([]interface{}); ok {
							for _, o := range s {
								if objval, ok := o.(map[string]interface{}); ok {
									var rNetworkInterfacesIPv6AccessConfigs dclService.InstanceNetworkInterfacesIPv6AccessConfigs
									if _, ok := objval["externalIPv6"]; ok {
										if s, ok := objval["externalIPv6"].(string); ok {
											rNetworkInterfacesIPv6AccessConfigs.ExternalIPv6 = dcl.String(s)
										} else {
											return nil, fmt.Errorf("rNetworkInterfacesIPv6AccessConfigs.ExternalIPv6: expected string")
										}
									}
									if _, ok := objval["externalIPv6PrefixLength"]; ok {
										if s, ok := objval["externalIPv6PrefixLength"].(string); ok {
											rNetworkInterfacesIPv6AccessConfigs.ExternalIPv6PrefixLength = dcl.String(s)
										} else {
											return nil, fmt.Errorf("rNetworkInterfacesIPv6AccessConfigs.ExternalIPv6PrefixLength: expected string")
										}
									}
									if _, ok := objval["name"]; ok {
										if s, ok := objval["name"].(string); ok {
											rNetworkInterfacesIPv6AccessConfigs.Name = dcl.String(s)
										} else {
											return nil, fmt.Errorf("rNetworkInterfacesIPv6AccessConfigs.Name: expected string")
										}
									}
									if _, ok := objval["natIP"]; ok {
										if s, ok := objval["natIP"].(string); ok {
											rNetworkInterfacesIPv6AccessConfigs.NatIP = dcl.String(s)
										} else {
											return nil, fmt.Errorf("rNetworkInterfacesIPv6AccessConfigs.NatIP: expected string")
										}
									}
									if _, ok := objval["networkTier"]; ok {
										if s, ok := objval["networkTier"].(string); ok {
											rNetworkInterfacesIPv6AccessConfigs.NetworkTier = dclService.InstanceNetworkInterfacesIPv6AccessConfigsNetworkTierEnumRef(s)
										} else {
											return nil, fmt.Errorf("rNetworkInterfacesIPv6AccessConfigs.NetworkTier: expected string")
										}
									}
									if _, ok := objval["publicPtrDomainName"]; ok {
										if s, ok := objval["publicPtrDomainName"].(string); ok {
											rNetworkInterfacesIPv6AccessConfigs.PublicPtrDomainName = dcl.String(s)
										} else {
											return nil, fmt.Errorf("rNetworkInterfacesIPv6AccessConfigs.PublicPtrDomainName: expected string")
										}
									}
									if _, ok := objval["setPublicPtr"]; ok {
										if b, ok := objval["setPublicPtr"].(bool); ok {
											rNetworkInterfacesIPv6AccessConfigs.SetPublicPtr = dcl.Bool(b)
										} else {
											return nil, fmt.Errorf("rNetworkInterfacesIPv6AccessConfigs.SetPublicPtr: expected bool")
										}
									}
									if _, ok := objval["type"]; ok {
										if s, ok := objval["type"].(string); ok {
											rNetworkInterfacesIPv6AccessConfigs.Type = dclService.InstanceNetworkInterfacesIPv6AccessConfigsTypeEnumRef(s)
										} else {
											return nil, fmt.Errorf("rNetworkInterfacesIPv6AccessConfigs.Type: expected string")
										}
									}
									rNetworkInterfaces.IPv6AccessConfigs = append(rNetworkInterfaces.IPv6AccessConfigs, rNetworkInterfacesIPv6AccessConfigs)
								}
							}
						} else {
							return nil, fmt.Errorf("rNetworkInterfaces.IPv6AccessConfigs: expected []interface{}")
						}
					}
					if _, ok := objval["name"]; ok {
						if s, ok := objval["name"].(string); ok {
							rNetworkInterfaces.Name = dcl.String(s)
						} else {
							return nil, fmt.Errorf("rNetworkInterfaces.Name: expected string")
						}
					}
					if _, ok := objval["network"]; ok {
						if s, ok := objval["network"].(string); ok {
							rNetworkInterfaces.Network = dcl.String(s)
						} else {
							return nil, fmt.Errorf("rNetworkInterfaces.Network: expected string")
						}
					}
					if _, ok := objval["networkIP"]; ok {
						if s, ok := objval["networkIP"].(string); ok {
							rNetworkInterfaces.NetworkIP = dcl.String(s)
						} else {
							return nil, fmt.Errorf("rNetworkInterfaces.NetworkIP: expected string")
						}
					}
					if _, ok := objval["subnetwork"]; ok {
						if s, ok := objval["subnetwork"].(string); ok {
							rNetworkInterfaces.Subnetwork = dcl.String(s)
						} else {
							return nil, fmt.Errorf("rNetworkInterfaces.Subnetwork: expected string")
						}
					}
					r.NetworkInterfaces = append(r.NetworkInterfaces, rNetworkInterfaces)
				}
			}
		} else {
			return nil, fmt.Errorf("r.NetworkInterfaces: expected []interface{}")
		}
	}
	if _, ok := u.Object["project"]; ok {
		if s, ok := u.Object["project"].(string); ok {
			r.Project = dcl.String(s)
		} else {
			return nil, fmt.Errorf("r.Project: expected string")
		}
	}
	if _, ok := u.Object["scheduling"]; ok {
		if rScheduling, ok := u.Object["scheduling"].(map[string]interface{}); ok {
			r.Scheduling = &dclService.InstanceScheduling{}
			if _, ok := rScheduling["automaticRestart"]; ok {
				if b, ok := rScheduling["automaticRestart"].(bool); ok {
					r.Scheduling.AutomaticRestart = dcl.Bool(b)
				} else {
					return nil, fmt.Errorf("r.Scheduling.AutomaticRestart: expected bool")
				}
			}
			if _, ok := rScheduling["onHostMaintenance"]; ok {
				if s, ok := rScheduling["onHostMaintenance"].(string); ok {
					r.Scheduling.OnHostMaintenance = dcl.String(s)
				} else {
					return nil, fmt.Errorf("r.Scheduling.OnHostMaintenance: expected string")
				}
			}
			if _, ok := rScheduling["preemptible"]; ok {
				if b, ok := rScheduling["preemptible"].(bool); ok {
					r.Scheduling.Preemptible = dcl.Bool(b)
				} else {
					return nil, fmt.Errorf("r.Scheduling.Preemptible: expected bool")
				}
			}
		} else {
			return nil, fmt.Errorf("r.Scheduling: expected map[string]interface{}")
		}
	}
	if _, ok := u.Object["selfLink"]; ok {
		if s, ok := u.Object["selfLink"].(string); ok {
			r.SelfLink = dcl.String(s)
		} else {
			return nil, fmt.Errorf("r.SelfLink: expected string")
		}
	}
	if _, ok := u.Object["serviceAccounts"]; ok {
		if s, ok := u.Object["serviceAccounts"].([]interface{}); ok {
			for _, o := range s {
				if objval, ok := o.(map[string]interface{}); ok {
					var rServiceAccounts dclService.InstanceServiceAccounts
					if _, ok := objval["email"]; ok {
						if s, ok := objval["email"].(string); ok {
							rServiceAccounts.Email = dcl.String(s)
						} else {
							return nil, fmt.Errorf("rServiceAccounts.Email: expected string")
						}
					}
					if _, ok := objval["scopes"]; ok {
						if s, ok := objval["scopes"].([]interface{}); ok {
							for _, ss := range s {
								if strval, ok := ss.(string); ok {
									rServiceAccounts.Scopes = append(rServiceAccounts.Scopes, strval)
								}
							}
						} else {
							return nil, fmt.Errorf("rServiceAccounts.Scopes: expected []interface{}")
						}
					}
					r.ServiceAccounts = append(r.ServiceAccounts, rServiceAccounts)
				}
			}
		} else {
			return nil, fmt.Errorf("r.ServiceAccounts: expected []interface{}")
		}
	}
	if _, ok := u.Object["shieldedInstanceConfig"]; ok {
		if rShieldedInstanceConfig, ok := u.Object["shieldedInstanceConfig"].(map[string]interface{}); ok {
			r.ShieldedInstanceConfig = &dclService.InstanceShieldedInstanceConfig{}
			if _, ok := rShieldedInstanceConfig["enableIntegrityMonitoring"]; ok {
				if b, ok := rShieldedInstanceConfig["enableIntegrityMonitoring"].(bool); ok {
					r.ShieldedInstanceConfig.EnableIntegrityMonitoring = dcl.Bool(b)
				} else {
					return nil, fmt.Errorf("r.ShieldedInstanceConfig.EnableIntegrityMonitoring: expected bool")
				}
			}
			if _, ok := rShieldedInstanceConfig["enableSecureBoot"]; ok {
				if b, ok := rShieldedInstanceConfig["enableSecureBoot"].(bool); ok {
					r.ShieldedInstanceConfig.EnableSecureBoot = dcl.Bool(b)
				} else {
					return nil, fmt.Errorf("r.ShieldedInstanceConfig.EnableSecureBoot: expected bool")
				}
			}
			if _, ok := rShieldedInstanceConfig["enableVtpm"]; ok {
				if b, ok := rShieldedInstanceConfig["enableVtpm"].(bool); ok {
					r.ShieldedInstanceConfig.EnableVtpm = dcl.Bool(b)
				} else {
					return nil, fmt.Errorf("r.ShieldedInstanceConfig.EnableVtpm: expected bool")
				}
			}
		} else {
			return nil, fmt.Errorf("r.ShieldedInstanceConfig: expected map[string]interface{}")
		}
	}
	if _, ok := u.Object["status"]; ok {
		if s, ok := u.Object["status"].(string); ok {
			r.Status = dclService.InstanceStatusEnumRef(s)
		} else {
			return nil, fmt.Errorf("r.Status: expected string")
		}
	}
	if _, ok := u.Object["statusMessage"]; ok {
		if s, ok := u.Object["statusMessage"].(string); ok {
			r.StatusMessage = dcl.String(s)
		} else {
			return nil, fmt.Errorf("r.StatusMessage: expected string")
		}
	}
	if _, ok := u.Object["tags"]; ok {
		if s, ok := u.Object["tags"].([]interface{}); ok {
			for _, ss := range s {
				if strval, ok := ss.(string); ok {
					r.Tags = append(r.Tags, strval)
				}
			}
		} else {
			return nil, fmt.Errorf("r.Tags: expected []interface{}")
		}
	}
	if _, ok := u.Object["zone"]; ok {
		if s, ok := u.Object["zone"].(string); ok {
			r.Zone = dcl.String(s)
		} else {
			return nil, fmt.Errorf("r.Zone: expected string")
		}
	}
	return r, nil
}

func GetInstance(ctx context.Context, config *dcl.Config, u *unstructured.Resource) (*unstructured.Resource, error) {
	c := dclService.NewClient(config)
	r, err := UnstructuredToInstance(u)
	if err != nil {
		return nil, err
	}
	r, err = c.GetInstance(ctx, r)
	if err != nil {
		return nil, err
	}
	return InstanceToUnstructured(r), nil
}

func ListInstance(ctx context.Context, config *dcl.Config, project string, zone string) ([]*unstructured.Resource, error) {
	c := dclService.NewClient(config)
	l, err := c.ListInstance(ctx, project, zone)
	if err != nil {
		return nil, err
	}
	var resources []*unstructured.Resource
	for {
		for _, r := range l.Items {
			resources = append(resources, InstanceToUnstructured(r))
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

func ApplyInstance(ctx context.Context, config *dcl.Config, u *unstructured.Resource, opts ...dcl.ApplyOption) (*unstructured.Resource, error) {
	c := dclService.NewClient(config)
	r, err := UnstructuredToInstance(u)
	if err != nil {
		return nil, err
	}
	if ush := unstructured.FetchStateHint(opts); ush != nil {
		sh, err := UnstructuredToInstance(ush)
		if err != nil {
			return nil, err
		}
		opts = append(opts, dcl.WithStateHint(sh))
	}
	r, err = c.ApplyInstance(ctx, r, opts...)
	if err != nil {
		return nil, err
	}
	return InstanceToUnstructured(r), nil
}

func InstanceHasDiff(ctx context.Context, config *dcl.Config, u *unstructured.Resource, opts ...dcl.ApplyOption) (bool, error) {
	c := dclService.NewClient(config)
	r, err := UnstructuredToInstance(u)
	if err != nil {
		return false, err
	}
	if ush := unstructured.FetchStateHint(opts); ush != nil {
		sh, err := UnstructuredToInstance(ush)
		if err != nil {
			return false, err
		}
		opts = append(opts, dcl.WithStateHint(sh))
	}
	opts = append(opts, dcl.WithLifecycleParam(dcl.BlockDestruction), dcl.WithLifecycleParam(dcl.BlockCreation), dcl.WithLifecycleParam(dcl.BlockModification))
	_, err = c.ApplyInstance(ctx, r, opts...)
	if err != nil {
		if _, ok := err.(dcl.ApplyInfeasibleError); ok {
			return true, nil
		}
		return false, err
	}
	return false, nil
}

func DeleteInstance(ctx context.Context, config *dcl.Config, u *unstructured.Resource) error {
	c := dclService.NewClient(config)
	r, err := UnstructuredToInstance(u)
	if err != nil {
		return err
	}
	return c.DeleteInstance(ctx, r)
}

func InstanceID(u *unstructured.Resource) (string, error) {
	r, err := UnstructuredToInstance(u)
	if err != nil {
		return "", err
	}
	return r.ID()
}

func (r *Instance) STV() unstructured.ServiceTypeVersion {
	return unstructured.ServiceTypeVersion{
		"compute",
		"Instance",
		"alpha",
	}
}

func SetPolicyInstance(ctx context.Context, config *dcl.Config, u *unstructured.Resource, p *unstructured.Resource) (*unstructured.Resource, error) {
	r, err := UnstructuredToInstance(u)
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

func SetPolicyWithEtagInstance(ctx context.Context, config *dcl.Config, u *unstructured.Resource, p *unstructured.Resource) (*unstructured.Resource, error) {
	r, err := UnstructuredToInstance(u)
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

func GetPolicyInstance(ctx context.Context, config *dcl.Config, u *unstructured.Resource) (*unstructured.Resource, error) {
	r, err := UnstructuredToInstance(u)
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

func SetPolicyMemberInstance(ctx context.Context, config *dcl.Config, u *unstructured.Resource, m *unstructured.Resource) (*unstructured.Resource, error) {
	r, err := UnstructuredToInstance(u)
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

func GetPolicyMemberInstance(ctx context.Context, config *dcl.Config, u *unstructured.Resource, role, member string) (*unstructured.Resource, error) {
	r, err := UnstructuredToInstance(u)
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

func DeletePolicyMemberInstance(ctx context.Context, config *dcl.Config, u *unstructured.Resource, m *unstructured.Resource) error {
	r, err := UnstructuredToInstance(u)
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

func (r *Instance) SetPolicyMember(ctx context.Context, config *dcl.Config, resource *unstructured.Resource, member *unstructured.Resource) (*unstructured.Resource, error) {
	return SetPolicyMemberInstance(ctx, config, resource, member)
}

func (r *Instance) GetPolicyMember(ctx context.Context, config *dcl.Config, resource *unstructured.Resource, role, member string) (*unstructured.Resource, error) {
	return GetPolicyMemberInstance(ctx, config, resource, role, member)
}

func (r *Instance) DeletePolicyMember(ctx context.Context, config *dcl.Config, resource *unstructured.Resource, member *unstructured.Resource) error {
	return DeletePolicyMemberInstance(ctx, config, resource, member)
}

func (r *Instance) SetPolicy(ctx context.Context, config *dcl.Config, resource *unstructured.Resource, policy *unstructured.Resource) (*unstructured.Resource, error) {
	return SetPolicyInstance(ctx, config, resource, policy)
}

func (r *Instance) SetPolicyWithEtag(ctx context.Context, config *dcl.Config, resource *unstructured.Resource, policy *unstructured.Resource) (*unstructured.Resource, error) {
	return SetPolicyWithEtagInstance(ctx, config, resource, policy)
}

func (r *Instance) GetPolicy(ctx context.Context, config *dcl.Config, resource *unstructured.Resource) (*unstructured.Resource, error) {
	return GetPolicyInstance(ctx, config, resource)
}

func (r *Instance) Get(ctx context.Context, config *dcl.Config, resource *unstructured.Resource) (*unstructured.Resource, error) {
	return GetInstance(ctx, config, resource)
}

func (r *Instance) Apply(ctx context.Context, config *dcl.Config, resource *unstructured.Resource, opts ...dcl.ApplyOption) (*unstructured.Resource, error) {
	return ApplyInstance(ctx, config, resource, opts...)
}

func (r *Instance) HasDiff(ctx context.Context, config *dcl.Config, resource *unstructured.Resource, opts ...dcl.ApplyOption) (bool, error) {
	return InstanceHasDiff(ctx, config, resource, opts...)
}

func (r *Instance) Delete(ctx context.Context, config *dcl.Config, resource *unstructured.Resource) error {
	return DeleteInstance(ctx, config, resource)
}

func (r *Instance) ID(resource *unstructured.Resource) (string, error) {
	return InstanceID(resource)
}

func init() {
	unstructured.Register(&Instance{})
}
