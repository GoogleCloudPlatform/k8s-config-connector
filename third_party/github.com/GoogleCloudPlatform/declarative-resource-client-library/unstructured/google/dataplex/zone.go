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

type Zone struct{}

func ZoneToUnstructured(r *dclService.Zone) *unstructured.Resource {
	u := &unstructured.Resource{
		STV: unstructured.ServiceTypeVersion{
			Service: "dataplex",
			Version: "ga",
			Type:    "Zone",
		},
		Object: make(map[string]interface{}),
	}
	if r.AssetStatus != nil && r.AssetStatus != dclService.EmptyZoneAssetStatus {
		rAssetStatus := make(map[string]interface{})
		if r.AssetStatus.ActiveAssets != nil {
			rAssetStatus["activeAssets"] = *r.AssetStatus.ActiveAssets
		}
		if r.AssetStatus.SecurityPolicyApplyingAssets != nil {
			rAssetStatus["securityPolicyApplyingAssets"] = *r.AssetStatus.SecurityPolicyApplyingAssets
		}
		if r.AssetStatus.UpdateTime != nil {
			rAssetStatus["updateTime"] = *r.AssetStatus.UpdateTime
		}
		u.Object["assetStatus"] = rAssetStatus
	}
	if r.CreateTime != nil {
		u.Object["createTime"] = *r.CreateTime
	}
	if r.Description != nil {
		u.Object["description"] = *r.Description
	}
	if r.DiscoverySpec != nil && r.DiscoverySpec != dclService.EmptyZoneDiscoverySpec {
		rDiscoverySpec := make(map[string]interface{})
		if r.DiscoverySpec.CsvOptions != nil && r.DiscoverySpec.CsvOptions != dclService.EmptyZoneDiscoverySpecCsvOptions {
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
		if r.DiscoverySpec.JsonOptions != nil && r.DiscoverySpec.JsonOptions != dclService.EmptyZoneDiscoverySpecJsonOptions {
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
	if r.ResourceSpec != nil && r.ResourceSpec != dclService.EmptyZoneResourceSpec {
		rResourceSpec := make(map[string]interface{})
		if r.ResourceSpec.LocationType != nil {
			rResourceSpec["locationType"] = string(*r.ResourceSpec.LocationType)
		}
		u.Object["resourceSpec"] = rResourceSpec
	}
	if r.State != nil {
		u.Object["state"] = string(*r.State)
	}
	if r.Type != nil {
		u.Object["type"] = string(*r.Type)
	}
	if r.Uid != nil {
		u.Object["uid"] = *r.Uid
	}
	if r.UpdateTime != nil {
		u.Object["updateTime"] = *r.UpdateTime
	}
	return u
}

func UnstructuredToZone(u *unstructured.Resource) (*dclService.Zone, error) {
	r := &dclService.Zone{}
	if _, ok := u.Object["assetStatus"]; ok {
		if rAssetStatus, ok := u.Object["assetStatus"].(map[string]interface{}); ok {
			r.AssetStatus = &dclService.ZoneAssetStatus{}
			if _, ok := rAssetStatus["activeAssets"]; ok {
				if i, ok := rAssetStatus["activeAssets"].(int64); ok {
					r.AssetStatus.ActiveAssets = dcl.Int64(i)
				} else {
					return nil, fmt.Errorf("r.AssetStatus.ActiveAssets: expected int64")
				}
			}
			if _, ok := rAssetStatus["securityPolicyApplyingAssets"]; ok {
				if i, ok := rAssetStatus["securityPolicyApplyingAssets"].(int64); ok {
					r.AssetStatus.SecurityPolicyApplyingAssets = dcl.Int64(i)
				} else {
					return nil, fmt.Errorf("r.AssetStatus.SecurityPolicyApplyingAssets: expected int64")
				}
			}
			if _, ok := rAssetStatus["updateTime"]; ok {
				if s, ok := rAssetStatus["updateTime"].(string); ok {
					r.AssetStatus.UpdateTime = dcl.String(s)
				} else {
					return nil, fmt.Errorf("r.AssetStatus.UpdateTime: expected string")
				}
			}
		} else {
			return nil, fmt.Errorf("r.AssetStatus: expected map[string]interface{}")
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
	if _, ok := u.Object["discoverySpec"]; ok {
		if rDiscoverySpec, ok := u.Object["discoverySpec"].(map[string]interface{}); ok {
			r.DiscoverySpec = &dclService.ZoneDiscoverySpec{}
			if _, ok := rDiscoverySpec["csvOptions"]; ok {
				if rDiscoverySpecCsvOptions, ok := rDiscoverySpec["csvOptions"].(map[string]interface{}); ok {
					r.DiscoverySpec.CsvOptions = &dclService.ZoneDiscoverySpecCsvOptions{}
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
					r.DiscoverySpec.JsonOptions = &dclService.ZoneDiscoverySpecJsonOptions{}
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
			r.ResourceSpec = &dclService.ZoneResourceSpec{}
			if _, ok := rResourceSpec["locationType"]; ok {
				if s, ok := rResourceSpec["locationType"].(string); ok {
					r.ResourceSpec.LocationType = dclService.ZoneResourceSpecLocationTypeEnumRef(s)
				} else {
					return nil, fmt.Errorf("r.ResourceSpec.LocationType: expected string")
				}
			}
		} else {
			return nil, fmt.Errorf("r.ResourceSpec: expected map[string]interface{}")
		}
	}
	if _, ok := u.Object["state"]; ok {
		if s, ok := u.Object["state"].(string); ok {
			r.State = dclService.ZoneStateEnumRef(s)
		} else {
			return nil, fmt.Errorf("r.State: expected string")
		}
	}
	if _, ok := u.Object["type"]; ok {
		if s, ok := u.Object["type"].(string); ok {
			r.Type = dclService.ZoneTypeEnumRef(s)
		} else {
			return nil, fmt.Errorf("r.Type: expected string")
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

func GetZone(ctx context.Context, config *dcl.Config, u *unstructured.Resource) (*unstructured.Resource, error) {
	c := dclService.NewClient(config)
	r, err := UnstructuredToZone(u)
	if err != nil {
		return nil, err
	}
	r, err = c.GetZone(ctx, r)
	if err != nil {
		return nil, err
	}
	return ZoneToUnstructured(r), nil
}

func ListZone(ctx context.Context, config *dcl.Config, project string, location string, lake string) ([]*unstructured.Resource, error) {
	c := dclService.NewClient(config)
	l, err := c.ListZone(ctx, project, location, lake)
	if err != nil {
		return nil, err
	}
	var resources []*unstructured.Resource
	for {
		for _, r := range l.Items {
			resources = append(resources, ZoneToUnstructured(r))
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

func ApplyZone(ctx context.Context, config *dcl.Config, u *unstructured.Resource, opts ...dcl.ApplyOption) (*unstructured.Resource, error) {
	c := dclService.NewClient(config)
	r, err := UnstructuredToZone(u)
	if err != nil {
		return nil, err
	}
	if ush := unstructured.FetchStateHint(opts); ush != nil {
		sh, err := UnstructuredToZone(ush)
		if err != nil {
			return nil, err
		}
		opts = append(opts, dcl.WithStateHint(sh))
	}
	r, err = c.ApplyZone(ctx, r, opts...)
	if err != nil {
		return nil, err
	}
	return ZoneToUnstructured(r), nil
}

func ZoneHasDiff(ctx context.Context, config *dcl.Config, u *unstructured.Resource, opts ...dcl.ApplyOption) (bool, error) {
	c := dclService.NewClient(config)
	r, err := UnstructuredToZone(u)
	if err != nil {
		return false, err
	}
	if ush := unstructured.FetchStateHint(opts); ush != nil {
		sh, err := UnstructuredToZone(ush)
		if err != nil {
			return false, err
		}
		opts = append(opts, dcl.WithStateHint(sh))
	}
	opts = append(opts, dcl.WithLifecycleParam(dcl.BlockDestruction), dcl.WithLifecycleParam(dcl.BlockCreation), dcl.WithLifecycleParam(dcl.BlockModification))
	_, err = c.ApplyZone(ctx, r, opts...)
	if err != nil {
		if _, ok := err.(dcl.ApplyInfeasibleError); ok {
			return true, nil
		}
		return false, err
	}
	return false, nil
}

func DeleteZone(ctx context.Context, config *dcl.Config, u *unstructured.Resource) error {
	c := dclService.NewClient(config)
	r, err := UnstructuredToZone(u)
	if err != nil {
		return err
	}
	return c.DeleteZone(ctx, r)
}

func ZoneID(u *unstructured.Resource) (string, error) {
	r, err := UnstructuredToZone(u)
	if err != nil {
		return "", err
	}
	return r.ID()
}

func (r *Zone) STV() unstructured.ServiceTypeVersion {
	return unstructured.ServiceTypeVersion{
		"dataplex",
		"Zone",
		"ga",
	}
}

func (r *Zone) SetPolicyMember(ctx context.Context, config *dcl.Config, resource *unstructured.Resource, member *unstructured.Resource) (*unstructured.Resource, error) {
	return nil, unstructured.ErrNoSuchMethod
}

func (r *Zone) GetPolicyMember(ctx context.Context, config *dcl.Config, resource *unstructured.Resource, role, member string) (*unstructured.Resource, error) {
	return nil, unstructured.ErrNoSuchMethod
}

func (r *Zone) DeletePolicyMember(ctx context.Context, config *dcl.Config, resource *unstructured.Resource, member *unstructured.Resource) error {
	return unstructured.ErrNoSuchMethod
}

func (r *Zone) SetPolicy(ctx context.Context, config *dcl.Config, resource *unstructured.Resource, policy *unstructured.Resource) (*unstructured.Resource, error) {
	return nil, unstructured.ErrNoSuchMethod
}

func (r *Zone) SetPolicyWithEtag(ctx context.Context, config *dcl.Config, resource *unstructured.Resource, policy *unstructured.Resource) (*unstructured.Resource, error) {
	return nil, unstructured.ErrNoSuchMethod
}

func (r *Zone) GetPolicy(ctx context.Context, config *dcl.Config, resource *unstructured.Resource) (*unstructured.Resource, error) {
	return nil, unstructured.ErrNoSuchMethod
}

func (r *Zone) Get(ctx context.Context, config *dcl.Config, resource *unstructured.Resource) (*unstructured.Resource, error) {
	return GetZone(ctx, config, resource)
}

func (r *Zone) Apply(ctx context.Context, config *dcl.Config, resource *unstructured.Resource, opts ...dcl.ApplyOption) (*unstructured.Resource, error) {
	return ApplyZone(ctx, config, resource, opts...)
}

func (r *Zone) HasDiff(ctx context.Context, config *dcl.Config, resource *unstructured.Resource, opts ...dcl.ApplyOption) (bool, error) {
	return ZoneHasDiff(ctx, config, resource, opts...)
}

func (r *Zone) Delete(ctx context.Context, config *dcl.Config, resource *unstructured.Resource) error {
	return DeleteZone(ctx, config, resource)
}

func (r *Zone) ID(resource *unstructured.Resource) (string, error) {
	return ZoneID(resource)
}

func init() {
	unstructured.Register(&Zone{})
}
