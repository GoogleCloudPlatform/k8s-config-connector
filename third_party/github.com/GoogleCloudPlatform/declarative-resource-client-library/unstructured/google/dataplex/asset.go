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
package dataplex

import (
	"context"
	"fmt"
	"github.com/GoogleCloudPlatform/declarative-resource-client-library/dcl"
	dclService "github.com/GoogleCloudPlatform/declarative-resource-client-library/services/google/dataplex"
	"github.com/GoogleCloudPlatform/declarative-resource-client-library/unstructured"
)

type Asset struct{}

func AssetToUnstructured(r *dclService.Asset) *unstructured.Resource {
	u := &unstructured.Resource{
		STV: unstructured.ServiceTypeVersion{
			Service: "dataplex",
			Version: "ga",
			Type:    "Asset",
		},
		Object: make(map[string]interface{}),
	}
	if r.CreateTime != nil {
		u.Object["createTime"] = *r.CreateTime
	}
	if r.DataplexZone != nil {
		u.Object["dataplexZone"] = *r.DataplexZone
	}
	if r.Description != nil {
		u.Object["description"] = *r.Description
	}
	if r.DiscoverySpec != nil && r.DiscoverySpec != dclService.EmptyAssetDiscoverySpec {
		rDiscoverySpec := make(map[string]interface{})
		if r.DiscoverySpec.CsvOptions != nil && r.DiscoverySpec.CsvOptions != dclService.EmptyAssetDiscoverySpecCsvOptions {
			rDiscoverySpecCsvOptions := make(map[string]interface{})
			if r.DiscoverySpec.CsvOptions.Delimiter != nil {
				rDiscoverySpecCsvOptions["delimiter"] = *r.DiscoverySpec.CsvOptions.Delimiter
			}
			if r.DiscoverySpec.CsvOptions.DisableTypeInference != nil {
				rDiscoverySpecCsvOptions["disableTypeInference"] = *r.DiscoverySpec.CsvOptions.DisableTypeInference
			}
			if r.DiscoverySpec.CsvOptions.Encoding != nil {
				rDiscoverySpecCsvOptions["encoding"] = *r.DiscoverySpec.CsvOptions.Encoding
			}
			if r.DiscoverySpec.CsvOptions.HeaderRows != nil {
				rDiscoverySpecCsvOptions["headerRows"] = *r.DiscoverySpec.CsvOptions.HeaderRows
			}
			rDiscoverySpec["csvOptions"] = rDiscoverySpecCsvOptions
		}
		if r.DiscoverySpec.Enabled != nil {
			rDiscoverySpec["enabled"] = *r.DiscoverySpec.Enabled
		}
		var rDiscoverySpecExcludePatterns []interface{}
		for _, rDiscoverySpecExcludePatternsVal := range r.DiscoverySpec.ExcludePatterns {
			rDiscoverySpecExcludePatterns = append(rDiscoverySpecExcludePatterns, rDiscoverySpecExcludePatternsVal)
		}
		rDiscoverySpec["excludePatterns"] = rDiscoverySpecExcludePatterns
		var rDiscoverySpecIncludePatterns []interface{}
		for _, rDiscoverySpecIncludePatternsVal := range r.DiscoverySpec.IncludePatterns {
			rDiscoverySpecIncludePatterns = append(rDiscoverySpecIncludePatterns, rDiscoverySpecIncludePatternsVal)
		}
		rDiscoverySpec["includePatterns"] = rDiscoverySpecIncludePatterns
		if r.DiscoverySpec.JsonOptions != nil && r.DiscoverySpec.JsonOptions != dclService.EmptyAssetDiscoverySpecJsonOptions {
			rDiscoverySpecJsonOptions := make(map[string]interface{})
			if r.DiscoverySpec.JsonOptions.DisableTypeInference != nil {
				rDiscoverySpecJsonOptions["disableTypeInference"] = *r.DiscoverySpec.JsonOptions.DisableTypeInference
			}
			if r.DiscoverySpec.JsonOptions.Encoding != nil {
				rDiscoverySpecJsonOptions["encoding"] = *r.DiscoverySpec.JsonOptions.Encoding
			}
			rDiscoverySpec["jsonOptions"] = rDiscoverySpecJsonOptions
		}
		if r.DiscoverySpec.Schedule != nil {
			rDiscoverySpec["schedule"] = *r.DiscoverySpec.Schedule
		}
		u.Object["discoverySpec"] = rDiscoverySpec
	}
	if r.DiscoveryStatus != nil && r.DiscoveryStatus != dclService.EmptyAssetDiscoveryStatus {
		rDiscoveryStatus := make(map[string]interface{})
		if r.DiscoveryStatus.LastRunDuration != nil {
			rDiscoveryStatus["lastRunDuration"] = *r.DiscoveryStatus.LastRunDuration
		}
		if r.DiscoveryStatus.LastRunTime != nil {
			rDiscoveryStatus["lastRunTime"] = *r.DiscoveryStatus.LastRunTime
		}
		if r.DiscoveryStatus.Message != nil {
			rDiscoveryStatus["message"] = *r.DiscoveryStatus.Message
		}
		if r.DiscoveryStatus.State != nil {
			rDiscoveryStatus["state"] = string(*r.DiscoveryStatus.State)
		}
		if r.DiscoveryStatus.Stats != nil && r.DiscoveryStatus.Stats != dclService.EmptyAssetDiscoveryStatusStats {
			rDiscoveryStatusStats := make(map[string]interface{})
			if r.DiscoveryStatus.Stats.DataItems != nil {
				rDiscoveryStatusStats["dataItems"] = *r.DiscoveryStatus.Stats.DataItems
			}
			if r.DiscoveryStatus.Stats.DataSize != nil {
				rDiscoveryStatusStats["dataSize"] = *r.DiscoveryStatus.Stats.DataSize
			}
			if r.DiscoveryStatus.Stats.Filesets != nil {
				rDiscoveryStatusStats["filesets"] = *r.DiscoveryStatus.Stats.Filesets
			}
			if r.DiscoveryStatus.Stats.Tables != nil {
				rDiscoveryStatusStats["tables"] = *r.DiscoveryStatus.Stats.Tables
			}
			rDiscoveryStatus["stats"] = rDiscoveryStatusStats
		}
		if r.DiscoveryStatus.UpdateTime != nil {
			rDiscoveryStatus["updateTime"] = *r.DiscoveryStatus.UpdateTime
		}
		u.Object["discoveryStatus"] = rDiscoveryStatus
	}
	if r.DisplayName != nil {
		u.Object["displayName"] = *r.DisplayName
	}
	if r.Labels != nil {
		rLabels := make(map[string]interface{})
		for k, v := range r.Labels {
			rLabels[k] = v
		}
		u.Object["labels"] = rLabels
	}
	if r.Lake != nil {
		u.Object["lake"] = *r.Lake
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
	if r.ResourceSpec != nil && r.ResourceSpec != dclService.EmptyAssetResourceSpec {
		rResourceSpec := make(map[string]interface{})
		if r.ResourceSpec.Name != nil {
			rResourceSpec["name"] = *r.ResourceSpec.Name
		}
		if r.ResourceSpec.ReadAccessMode != nil {
			rResourceSpec["readAccessMode"] = string(*r.ResourceSpec.ReadAccessMode)
		}
		if r.ResourceSpec.Type != nil {
			rResourceSpec["type"] = string(*r.ResourceSpec.Type)
		}
		u.Object["resourceSpec"] = rResourceSpec
	}
	if r.ResourceStatus != nil && r.ResourceStatus != dclService.EmptyAssetResourceStatus {
		rResourceStatus := make(map[string]interface{})
		if r.ResourceStatus.Message != nil {
			rResourceStatus["message"] = *r.ResourceStatus.Message
		}
		if r.ResourceStatus.State != nil {
			rResourceStatus["state"] = string(*r.ResourceStatus.State)
		}
		if r.ResourceStatus.UpdateTime != nil {
			rResourceStatus["updateTime"] = *r.ResourceStatus.UpdateTime
		}
		u.Object["resourceStatus"] = rResourceStatus
	}
	if r.SecurityStatus != nil && r.SecurityStatus != dclService.EmptyAssetSecurityStatus {
		rSecurityStatus := make(map[string]interface{})
		if r.SecurityStatus.Message != nil {
			rSecurityStatus["message"] = *r.SecurityStatus.Message
		}
		if r.SecurityStatus.State != nil {
			rSecurityStatus["state"] = string(*r.SecurityStatus.State)
		}
		if r.SecurityStatus.UpdateTime != nil {
			rSecurityStatus["updateTime"] = *r.SecurityStatus.UpdateTime
		}
		u.Object["securityStatus"] = rSecurityStatus
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
	return u
}

func UnstructuredToAsset(u *unstructured.Resource) (*dclService.Asset, error) {
	r := &dclService.Asset{}
	if _, ok := u.Object["createTime"]; ok {
		if s, ok := u.Object["createTime"].(string); ok {
			r.CreateTime = dcl.String(s)
		} else {
			return nil, fmt.Errorf("r.CreateTime: expected string")
		}
	}
	if _, ok := u.Object["dataplexZone"]; ok {
		if s, ok := u.Object["dataplexZone"].(string); ok {
			r.DataplexZone = dcl.String(s)
		} else {
			return nil, fmt.Errorf("r.DataplexZone: expected string")
		}
	}
	if _, ok := u.Object["description"]; ok {
		if s, ok := u.Object["description"].(string); ok {
			r.Description = dcl.String(s)
		} else {
			return nil, fmt.Errorf("r.Description: expected string")
		}
	}
	if _, ok := u.Object["discoverySpec"]; ok {
		if rDiscoverySpec, ok := u.Object["discoverySpec"].(map[string]interface{}); ok {
			r.DiscoverySpec = &dclService.AssetDiscoverySpec{}
			if _, ok := rDiscoverySpec["csvOptions"]; ok {
				if rDiscoverySpecCsvOptions, ok := rDiscoverySpec["csvOptions"].(map[string]interface{}); ok {
					r.DiscoverySpec.CsvOptions = &dclService.AssetDiscoverySpecCsvOptions{}
					if _, ok := rDiscoverySpecCsvOptions["delimiter"]; ok {
						if s, ok := rDiscoverySpecCsvOptions["delimiter"].(string); ok {
							r.DiscoverySpec.CsvOptions.Delimiter = dcl.String(s)
						} else {
							return nil, fmt.Errorf("r.DiscoverySpec.CsvOptions.Delimiter: expected string")
						}
					}
					if _, ok := rDiscoverySpecCsvOptions["disableTypeInference"]; ok {
						if b, ok := rDiscoverySpecCsvOptions["disableTypeInference"].(bool); ok {
							r.DiscoverySpec.CsvOptions.DisableTypeInference = dcl.Bool(b)
						} else {
							return nil, fmt.Errorf("r.DiscoverySpec.CsvOptions.DisableTypeInference: expected bool")
						}
					}
					if _, ok := rDiscoverySpecCsvOptions["encoding"]; ok {
						if s, ok := rDiscoverySpecCsvOptions["encoding"].(string); ok {
							r.DiscoverySpec.CsvOptions.Encoding = dcl.String(s)
						} else {
							return nil, fmt.Errorf("r.DiscoverySpec.CsvOptions.Encoding: expected string")
						}
					}
					if _, ok := rDiscoverySpecCsvOptions["headerRows"]; ok {
						if i, ok := rDiscoverySpecCsvOptions["headerRows"].(int64); ok {
							r.DiscoverySpec.CsvOptions.HeaderRows = dcl.Int64(i)
						} else {
							return nil, fmt.Errorf("r.DiscoverySpec.CsvOptions.HeaderRows: expected int64")
						}
					}
				} else {
					return nil, fmt.Errorf("r.DiscoverySpec.CsvOptions: expected map[string]interface{}")
				}
			}
			if _, ok := rDiscoverySpec["enabled"]; ok {
				if b, ok := rDiscoverySpec["enabled"].(bool); ok {
					r.DiscoverySpec.Enabled = dcl.Bool(b)
				} else {
					return nil, fmt.Errorf("r.DiscoverySpec.Enabled: expected bool")
				}
			}
			if _, ok := rDiscoverySpec["excludePatterns"]; ok {
				if s, ok := rDiscoverySpec["excludePatterns"].([]interface{}); ok {
					for _, ss := range s {
						if strval, ok := ss.(string); ok {
							r.DiscoverySpec.ExcludePatterns = append(r.DiscoverySpec.ExcludePatterns, strval)
						}
					}
				} else {
					return nil, fmt.Errorf("r.DiscoverySpec.ExcludePatterns: expected []interface{}")
				}
			}
			if _, ok := rDiscoverySpec["includePatterns"]; ok {
				if s, ok := rDiscoverySpec["includePatterns"].([]interface{}); ok {
					for _, ss := range s {
						if strval, ok := ss.(string); ok {
							r.DiscoverySpec.IncludePatterns = append(r.DiscoverySpec.IncludePatterns, strval)
						}
					}
				} else {
					return nil, fmt.Errorf("r.DiscoverySpec.IncludePatterns: expected []interface{}")
				}
			}
			if _, ok := rDiscoverySpec["jsonOptions"]; ok {
				if rDiscoverySpecJsonOptions, ok := rDiscoverySpec["jsonOptions"].(map[string]interface{}); ok {
					r.DiscoverySpec.JsonOptions = &dclService.AssetDiscoverySpecJsonOptions{}
					if _, ok := rDiscoverySpecJsonOptions["disableTypeInference"]; ok {
						if b, ok := rDiscoverySpecJsonOptions["disableTypeInference"].(bool); ok {
							r.DiscoverySpec.JsonOptions.DisableTypeInference = dcl.Bool(b)
						} else {
							return nil, fmt.Errorf("r.DiscoverySpec.JsonOptions.DisableTypeInference: expected bool")
						}
					}
					if _, ok := rDiscoverySpecJsonOptions["encoding"]; ok {
						if s, ok := rDiscoverySpecJsonOptions["encoding"].(string); ok {
							r.DiscoverySpec.JsonOptions.Encoding = dcl.String(s)
						} else {
							return nil, fmt.Errorf("r.DiscoverySpec.JsonOptions.Encoding: expected string")
						}
					}
				} else {
					return nil, fmt.Errorf("r.DiscoverySpec.JsonOptions: expected map[string]interface{}")
				}
			}
			if _, ok := rDiscoverySpec["schedule"]; ok {
				if s, ok := rDiscoverySpec["schedule"].(string); ok {
					r.DiscoverySpec.Schedule = dcl.String(s)
				} else {
					return nil, fmt.Errorf("r.DiscoverySpec.Schedule: expected string")
				}
			}
		} else {
			return nil, fmt.Errorf("r.DiscoverySpec: expected map[string]interface{}")
		}
	}
	if _, ok := u.Object["discoveryStatus"]; ok {
		if rDiscoveryStatus, ok := u.Object["discoveryStatus"].(map[string]interface{}); ok {
			r.DiscoveryStatus = &dclService.AssetDiscoveryStatus{}
			if _, ok := rDiscoveryStatus["lastRunDuration"]; ok {
				if s, ok := rDiscoveryStatus["lastRunDuration"].(string); ok {
					r.DiscoveryStatus.LastRunDuration = dcl.String(s)
				} else {
					return nil, fmt.Errorf("r.DiscoveryStatus.LastRunDuration: expected string")
				}
			}
			if _, ok := rDiscoveryStatus["lastRunTime"]; ok {
				if s, ok := rDiscoveryStatus["lastRunTime"].(string); ok {
					r.DiscoveryStatus.LastRunTime = dcl.String(s)
				} else {
					return nil, fmt.Errorf("r.DiscoveryStatus.LastRunTime: expected string")
				}
			}
			if _, ok := rDiscoveryStatus["message"]; ok {
				if s, ok := rDiscoveryStatus["message"].(string); ok {
					r.DiscoveryStatus.Message = dcl.String(s)
				} else {
					return nil, fmt.Errorf("r.DiscoveryStatus.Message: expected string")
				}
			}
			if _, ok := rDiscoveryStatus["state"]; ok {
				if s, ok := rDiscoveryStatus["state"].(string); ok {
					r.DiscoveryStatus.State = dclService.AssetDiscoveryStatusStateEnumRef(s)
				} else {
					return nil, fmt.Errorf("r.DiscoveryStatus.State: expected string")
				}
			}
			if _, ok := rDiscoveryStatus["stats"]; ok {
				if rDiscoveryStatusStats, ok := rDiscoveryStatus["stats"].(map[string]interface{}); ok {
					r.DiscoveryStatus.Stats = &dclService.AssetDiscoveryStatusStats{}
					if _, ok := rDiscoveryStatusStats["dataItems"]; ok {
						if i, ok := rDiscoveryStatusStats["dataItems"].(int64); ok {
							r.DiscoveryStatus.Stats.DataItems = dcl.Int64(i)
						} else {
							return nil, fmt.Errorf("r.DiscoveryStatus.Stats.DataItems: expected int64")
						}
					}
					if _, ok := rDiscoveryStatusStats["dataSize"]; ok {
						if i, ok := rDiscoveryStatusStats["dataSize"].(int64); ok {
							r.DiscoveryStatus.Stats.DataSize = dcl.Int64(i)
						} else {
							return nil, fmt.Errorf("r.DiscoveryStatus.Stats.DataSize: expected int64")
						}
					}
					if _, ok := rDiscoveryStatusStats["filesets"]; ok {
						if i, ok := rDiscoveryStatusStats["filesets"].(int64); ok {
							r.DiscoveryStatus.Stats.Filesets = dcl.Int64(i)
						} else {
							return nil, fmt.Errorf("r.DiscoveryStatus.Stats.Filesets: expected int64")
						}
					}
					if _, ok := rDiscoveryStatusStats["tables"]; ok {
						if i, ok := rDiscoveryStatusStats["tables"].(int64); ok {
							r.DiscoveryStatus.Stats.Tables = dcl.Int64(i)
						} else {
							return nil, fmt.Errorf("r.DiscoveryStatus.Stats.Tables: expected int64")
						}
					}
				} else {
					return nil, fmt.Errorf("r.DiscoveryStatus.Stats: expected map[string]interface{}")
				}
			}
			if _, ok := rDiscoveryStatus["updateTime"]; ok {
				if s, ok := rDiscoveryStatus["updateTime"].(string); ok {
					r.DiscoveryStatus.UpdateTime = dcl.String(s)
				} else {
					return nil, fmt.Errorf("r.DiscoveryStatus.UpdateTime: expected string")
				}
			}
		} else {
			return nil, fmt.Errorf("r.DiscoveryStatus: expected map[string]interface{}")
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
	if _, ok := u.Object["lake"]; ok {
		if s, ok := u.Object["lake"].(string); ok {
			r.Lake = dcl.String(s)
		} else {
			return nil, fmt.Errorf("r.Lake: expected string")
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
	if _, ok := u.Object["resourceSpec"]; ok {
		if rResourceSpec, ok := u.Object["resourceSpec"].(map[string]interface{}); ok {
			r.ResourceSpec = &dclService.AssetResourceSpec{}
			if _, ok := rResourceSpec["name"]; ok {
				if s, ok := rResourceSpec["name"].(string); ok {
					r.ResourceSpec.Name = dcl.String(s)
				} else {
					return nil, fmt.Errorf("r.ResourceSpec.Name: expected string")
				}
			}
			if _, ok := rResourceSpec["readAccessMode"]; ok {
				if s, ok := rResourceSpec["readAccessMode"].(string); ok {
					r.ResourceSpec.ReadAccessMode = dclService.AssetResourceSpecReadAccessModeEnumRef(s)
				} else {
					return nil, fmt.Errorf("r.ResourceSpec.ReadAccessMode: expected string")
				}
			}
			if _, ok := rResourceSpec["type"]; ok {
				if s, ok := rResourceSpec["type"].(string); ok {
					r.ResourceSpec.Type = dclService.AssetResourceSpecTypeEnumRef(s)
				} else {
					return nil, fmt.Errorf("r.ResourceSpec.Type: expected string")
				}
			}
		} else {
			return nil, fmt.Errorf("r.ResourceSpec: expected map[string]interface{}")
		}
	}
	if _, ok := u.Object["resourceStatus"]; ok {
		if rResourceStatus, ok := u.Object["resourceStatus"].(map[string]interface{}); ok {
			r.ResourceStatus = &dclService.AssetResourceStatus{}
			if _, ok := rResourceStatus["message"]; ok {
				if s, ok := rResourceStatus["message"].(string); ok {
					r.ResourceStatus.Message = dcl.String(s)
				} else {
					return nil, fmt.Errorf("r.ResourceStatus.Message: expected string")
				}
			}
			if _, ok := rResourceStatus["state"]; ok {
				if s, ok := rResourceStatus["state"].(string); ok {
					r.ResourceStatus.State = dclService.AssetResourceStatusStateEnumRef(s)
				} else {
					return nil, fmt.Errorf("r.ResourceStatus.State: expected string")
				}
			}
			if _, ok := rResourceStatus["updateTime"]; ok {
				if s, ok := rResourceStatus["updateTime"].(string); ok {
					r.ResourceStatus.UpdateTime = dcl.String(s)
				} else {
					return nil, fmt.Errorf("r.ResourceStatus.UpdateTime: expected string")
				}
			}
		} else {
			return nil, fmt.Errorf("r.ResourceStatus: expected map[string]interface{}")
		}
	}
	if _, ok := u.Object["securityStatus"]; ok {
		if rSecurityStatus, ok := u.Object["securityStatus"].(map[string]interface{}); ok {
			r.SecurityStatus = &dclService.AssetSecurityStatus{}
			if _, ok := rSecurityStatus["message"]; ok {
				if s, ok := rSecurityStatus["message"].(string); ok {
					r.SecurityStatus.Message = dcl.String(s)
				} else {
					return nil, fmt.Errorf("r.SecurityStatus.Message: expected string")
				}
			}
			if _, ok := rSecurityStatus["state"]; ok {
				if s, ok := rSecurityStatus["state"].(string); ok {
					r.SecurityStatus.State = dclService.AssetSecurityStatusStateEnumRef(s)
				} else {
					return nil, fmt.Errorf("r.SecurityStatus.State: expected string")
				}
			}
			if _, ok := rSecurityStatus["updateTime"]; ok {
				if s, ok := rSecurityStatus["updateTime"].(string); ok {
					r.SecurityStatus.UpdateTime = dcl.String(s)
				} else {
					return nil, fmt.Errorf("r.SecurityStatus.UpdateTime: expected string")
				}
			}
		} else {
			return nil, fmt.Errorf("r.SecurityStatus: expected map[string]interface{}")
		}
	}
	if _, ok := u.Object["state"]; ok {
		if s, ok := u.Object["state"].(string); ok {
			r.State = dclService.AssetStateEnumRef(s)
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
	return r, nil
}

func GetAsset(ctx context.Context, config *dcl.Config, u *unstructured.Resource) (*unstructured.Resource, error) {
	c := dclService.NewClient(config)
	r, err := UnstructuredToAsset(u)
	if err != nil {
		return nil, err
	}
	r, err = c.GetAsset(ctx, r)
	if err != nil {
		return nil, err
	}
	return AssetToUnstructured(r), nil
}

func ListAsset(ctx context.Context, config *dcl.Config, project string, location string, dataplexZone string, lake string) ([]*unstructured.Resource, error) {
	c := dclService.NewClient(config)
	l, err := c.ListAsset(ctx, project, location, dataplexZone, lake)
	if err != nil {
		return nil, err
	}
	var resources []*unstructured.Resource
	for {
		for _, r := range l.Items {
			resources = append(resources, AssetToUnstructured(r))
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

func ApplyAsset(ctx context.Context, config *dcl.Config, u *unstructured.Resource, opts ...dcl.ApplyOption) (*unstructured.Resource, error) {
	c := dclService.NewClient(config)
	r, err := UnstructuredToAsset(u)
	if err != nil {
		return nil, err
	}
	if ush := unstructured.FetchStateHint(opts); ush != nil {
		sh, err := UnstructuredToAsset(ush)
		if err != nil {
			return nil, err
		}
		opts = append(opts, dcl.WithStateHint(sh))
	}
	r, err = c.ApplyAsset(ctx, r, opts...)
	if err != nil {
		return nil, err
	}
	return AssetToUnstructured(r), nil
}

func AssetHasDiff(ctx context.Context, config *dcl.Config, u *unstructured.Resource, opts ...dcl.ApplyOption) (bool, error) {
	c := dclService.NewClient(config)
	r, err := UnstructuredToAsset(u)
	if err != nil {
		return false, err
	}
	if ush := unstructured.FetchStateHint(opts); ush != nil {
		sh, err := UnstructuredToAsset(ush)
		if err != nil {
			return false, err
		}
		opts = append(opts, dcl.WithStateHint(sh))
	}
	opts = append(opts, dcl.WithLifecycleParam(dcl.BlockDestruction), dcl.WithLifecycleParam(dcl.BlockCreation), dcl.WithLifecycleParam(dcl.BlockModification))
	_, err = c.ApplyAsset(ctx, r, opts...)
	if err != nil {
		if _, ok := err.(dcl.ApplyInfeasibleError); ok {
			return true, nil
		}
		return false, err
	}
	return false, nil
}

func DeleteAsset(ctx context.Context, config *dcl.Config, u *unstructured.Resource) error {
	c := dclService.NewClient(config)
	r, err := UnstructuredToAsset(u)
	if err != nil {
		return err
	}
	return c.DeleteAsset(ctx, r)
}

func AssetID(u *unstructured.Resource) (string, error) {
	r, err := UnstructuredToAsset(u)
	if err != nil {
		return "", err
	}
	return r.ID()
}

func (r *Asset) STV() unstructured.ServiceTypeVersion {
	return unstructured.ServiceTypeVersion{
		"dataplex",
		"Asset",
		"ga",
	}
}

func (r *Asset) SetPolicyMember(ctx context.Context, config *dcl.Config, resource *unstructured.Resource, member *unstructured.Resource) (*unstructured.Resource, error) {
	return nil, unstructured.ErrNoSuchMethod
}

func (r *Asset) GetPolicyMember(ctx context.Context, config *dcl.Config, resource *unstructured.Resource, role, member string) (*unstructured.Resource, error) {
	return nil, unstructured.ErrNoSuchMethod
}

func (r *Asset) DeletePolicyMember(ctx context.Context, config *dcl.Config, resource *unstructured.Resource, member *unstructured.Resource) error {
	return unstructured.ErrNoSuchMethod
}

func (r *Asset) SetPolicy(ctx context.Context, config *dcl.Config, resource *unstructured.Resource, policy *unstructured.Resource) (*unstructured.Resource, error) {
	return nil, unstructured.ErrNoSuchMethod
}

func (r *Asset) SetPolicyWithEtag(ctx context.Context, config *dcl.Config, resource *unstructured.Resource, policy *unstructured.Resource) (*unstructured.Resource, error) {
	return nil, unstructured.ErrNoSuchMethod
}

func (r *Asset) GetPolicy(ctx context.Context, config *dcl.Config, resource *unstructured.Resource) (*unstructured.Resource, error) {
	return nil, unstructured.ErrNoSuchMethod
}

func (r *Asset) Get(ctx context.Context, config *dcl.Config, resource *unstructured.Resource) (*unstructured.Resource, error) {
	return GetAsset(ctx, config, resource)
}

func (r *Asset) Apply(ctx context.Context, config *dcl.Config, resource *unstructured.Resource, opts ...dcl.ApplyOption) (*unstructured.Resource, error) {
	return ApplyAsset(ctx, config, resource, opts...)
}

func (r *Asset) HasDiff(ctx context.Context, config *dcl.Config, resource *unstructured.Resource, opts ...dcl.ApplyOption) (bool, error) {
	return AssetHasDiff(ctx, config, resource, opts...)
}

func (r *Asset) Delete(ctx context.Context, config *dcl.Config, resource *unstructured.Resource) error {
	return DeleteAsset(ctx, config, resource)
}

func (r *Asset) ID(resource *unstructured.Resource) (string, error) {
	return AssetID(resource)
}

func init() {
	unstructured.Register(&Asset{})
}
