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
package dlp

import (
	"context"
	"fmt"
	"github.com/GoogleCloudPlatform/declarative-resource-client-library/dcl"
	dclService "github.com/GoogleCloudPlatform/declarative-resource-client-library/services/google/dlp"
	"github.com/GoogleCloudPlatform/declarative-resource-client-library/unstructured"
)

type StoredInfoType struct{}

func StoredInfoTypeToUnstructured(r *dclService.StoredInfoType) *unstructured.Resource {
	u := &unstructured.Resource{
		STV: unstructured.ServiceTypeVersion{
			Service: "dlp",
			Version: "ga",
			Type:    "StoredInfoType",
		},
		Object: make(map[string]interface{}),
	}
	if r.Description != nil {
		u.Object["description"] = *r.Description
	}
	if r.Dictionary != nil && r.Dictionary != dclService.EmptyStoredInfoTypeDictionary {
		rDictionary := make(map[string]interface{})
		if r.Dictionary.CloudStoragePath != nil && r.Dictionary.CloudStoragePath != dclService.EmptyStoredInfoTypeDictionaryCloudStoragePath {
			rDictionaryCloudStoragePath := make(map[string]interface{})
			if r.Dictionary.CloudStoragePath.Path != nil {
				rDictionaryCloudStoragePath["path"] = *r.Dictionary.CloudStoragePath.Path
			}
			rDictionary["cloudStoragePath"] = rDictionaryCloudStoragePath
		}
		if r.Dictionary.WordList != nil && r.Dictionary.WordList != dclService.EmptyStoredInfoTypeDictionaryWordList {
			rDictionaryWordList := make(map[string]interface{})
			var rDictionaryWordListWords []interface{}
			for _, rDictionaryWordListWordsVal := range r.Dictionary.WordList.Words {
				rDictionaryWordListWords = append(rDictionaryWordListWords, rDictionaryWordListWordsVal)
			}
			rDictionaryWordList["words"] = rDictionaryWordListWords
			rDictionary["wordList"] = rDictionaryWordList
		}
		u.Object["dictionary"] = rDictionary
	}
	if r.DisplayName != nil {
		u.Object["displayName"] = *r.DisplayName
	}
	if r.LargeCustomDictionary != nil && r.LargeCustomDictionary != dclService.EmptyStoredInfoTypeLargeCustomDictionary {
		rLargeCustomDictionary := make(map[string]interface{})
		if r.LargeCustomDictionary.BigQueryField != nil && r.LargeCustomDictionary.BigQueryField != dclService.EmptyStoredInfoTypeLargeCustomDictionaryBigQueryField {
			rLargeCustomDictionaryBigQueryField := make(map[string]interface{})
			if r.LargeCustomDictionary.BigQueryField.Field != nil && r.LargeCustomDictionary.BigQueryField.Field != dclService.EmptyStoredInfoTypeLargeCustomDictionaryBigQueryFieldField {
				rLargeCustomDictionaryBigQueryFieldField := make(map[string]interface{})
				if r.LargeCustomDictionary.BigQueryField.Field.Name != nil {
					rLargeCustomDictionaryBigQueryFieldField["name"] = *r.LargeCustomDictionary.BigQueryField.Field.Name
				}
				rLargeCustomDictionaryBigQueryField["field"] = rLargeCustomDictionaryBigQueryFieldField
			}
			if r.LargeCustomDictionary.BigQueryField.Table != nil && r.LargeCustomDictionary.BigQueryField.Table != dclService.EmptyStoredInfoTypeLargeCustomDictionaryBigQueryFieldTable {
				rLargeCustomDictionaryBigQueryFieldTable := make(map[string]interface{})
				if r.LargeCustomDictionary.BigQueryField.Table.DatasetId != nil {
					rLargeCustomDictionaryBigQueryFieldTable["datasetId"] = *r.LargeCustomDictionary.BigQueryField.Table.DatasetId
				}
				if r.LargeCustomDictionary.BigQueryField.Table.ProjectId != nil {
					rLargeCustomDictionaryBigQueryFieldTable["projectId"] = *r.LargeCustomDictionary.BigQueryField.Table.ProjectId
				}
				if r.LargeCustomDictionary.BigQueryField.Table.TableId != nil {
					rLargeCustomDictionaryBigQueryFieldTable["tableId"] = *r.LargeCustomDictionary.BigQueryField.Table.TableId
				}
				rLargeCustomDictionaryBigQueryField["table"] = rLargeCustomDictionaryBigQueryFieldTable
			}
			rLargeCustomDictionary["bigQueryField"] = rLargeCustomDictionaryBigQueryField
		}
		if r.LargeCustomDictionary.CloudStorageFileSet != nil && r.LargeCustomDictionary.CloudStorageFileSet != dclService.EmptyStoredInfoTypeLargeCustomDictionaryCloudStorageFileSet {
			rLargeCustomDictionaryCloudStorageFileSet := make(map[string]interface{})
			if r.LargeCustomDictionary.CloudStorageFileSet.Url != nil {
				rLargeCustomDictionaryCloudStorageFileSet["url"] = *r.LargeCustomDictionary.CloudStorageFileSet.Url
			}
			rLargeCustomDictionary["cloudStorageFileSet"] = rLargeCustomDictionaryCloudStorageFileSet
		}
		if r.LargeCustomDictionary.OutputPath != nil && r.LargeCustomDictionary.OutputPath != dclService.EmptyStoredInfoTypeLargeCustomDictionaryOutputPath {
			rLargeCustomDictionaryOutputPath := make(map[string]interface{})
			if r.LargeCustomDictionary.OutputPath.Path != nil {
				rLargeCustomDictionaryOutputPath["path"] = *r.LargeCustomDictionary.OutputPath.Path
			}
			rLargeCustomDictionary["outputPath"] = rLargeCustomDictionaryOutputPath
		}
		u.Object["largeCustomDictionary"] = rLargeCustomDictionary
	}
	if r.Location != nil {
		u.Object["location"] = *r.Location
	}
	if r.Name != nil {
		u.Object["name"] = *r.Name
	}
	if r.Parent != nil {
		u.Object["parent"] = *r.Parent
	}
	if r.Regex != nil && r.Regex != dclService.EmptyStoredInfoTypeRegex {
		rRegex := make(map[string]interface{})
		var rRegexGroupIndexes []interface{}
		for _, rRegexGroupIndexesVal := range r.Regex.GroupIndexes {
			rRegexGroupIndexes = append(rRegexGroupIndexes, rRegexGroupIndexesVal)
		}
		rRegex["groupIndexes"] = rRegexGroupIndexes
		if r.Regex.Pattern != nil {
			rRegex["pattern"] = *r.Regex.Pattern
		}
		u.Object["regex"] = rRegex
	}
	return u
}

func UnstructuredToStoredInfoType(u *unstructured.Resource) (*dclService.StoredInfoType, error) {
	r := &dclService.StoredInfoType{}
	if _, ok := u.Object["description"]; ok {
		if s, ok := u.Object["description"].(string); ok {
			r.Description = dcl.String(s)
		} else {
			return nil, fmt.Errorf("r.Description: expected string")
		}
	}
	if _, ok := u.Object["dictionary"]; ok {
		if rDictionary, ok := u.Object["dictionary"].(map[string]interface{}); ok {
			r.Dictionary = &dclService.StoredInfoTypeDictionary{}
			if _, ok := rDictionary["cloudStoragePath"]; ok {
				if rDictionaryCloudStoragePath, ok := rDictionary["cloudStoragePath"].(map[string]interface{}); ok {
					r.Dictionary.CloudStoragePath = &dclService.StoredInfoTypeDictionaryCloudStoragePath{}
					if _, ok := rDictionaryCloudStoragePath["path"]; ok {
						if s, ok := rDictionaryCloudStoragePath["path"].(string); ok {
							r.Dictionary.CloudStoragePath.Path = dcl.String(s)
						} else {
							return nil, fmt.Errorf("r.Dictionary.CloudStoragePath.Path: expected string")
						}
					}
				} else {
					return nil, fmt.Errorf("r.Dictionary.CloudStoragePath: expected map[string]interface{}")
				}
			}
			if _, ok := rDictionary["wordList"]; ok {
				if rDictionaryWordList, ok := rDictionary["wordList"].(map[string]interface{}); ok {
					r.Dictionary.WordList = &dclService.StoredInfoTypeDictionaryWordList{}
					if _, ok := rDictionaryWordList["words"]; ok {
						if s, ok := rDictionaryWordList["words"].([]interface{}); ok {
							for _, ss := range s {
								if strval, ok := ss.(string); ok {
									r.Dictionary.WordList.Words = append(r.Dictionary.WordList.Words, strval)
								}
							}
						} else {
							return nil, fmt.Errorf("r.Dictionary.WordList.Words: expected []interface{}")
						}
					}
				} else {
					return nil, fmt.Errorf("r.Dictionary.WordList: expected map[string]interface{}")
				}
			}
		} else {
			return nil, fmt.Errorf("r.Dictionary: expected map[string]interface{}")
		}
	}
	if _, ok := u.Object["displayName"]; ok {
		if s, ok := u.Object["displayName"].(string); ok {
			r.DisplayName = dcl.String(s)
		} else {
			return nil, fmt.Errorf("r.DisplayName: expected string")
		}
	}
	if _, ok := u.Object["largeCustomDictionary"]; ok {
		if rLargeCustomDictionary, ok := u.Object["largeCustomDictionary"].(map[string]interface{}); ok {
			r.LargeCustomDictionary = &dclService.StoredInfoTypeLargeCustomDictionary{}
			if _, ok := rLargeCustomDictionary["bigQueryField"]; ok {
				if rLargeCustomDictionaryBigQueryField, ok := rLargeCustomDictionary["bigQueryField"].(map[string]interface{}); ok {
					r.LargeCustomDictionary.BigQueryField = &dclService.StoredInfoTypeLargeCustomDictionaryBigQueryField{}
					if _, ok := rLargeCustomDictionaryBigQueryField["field"]; ok {
						if rLargeCustomDictionaryBigQueryFieldField, ok := rLargeCustomDictionaryBigQueryField["field"].(map[string]interface{}); ok {
							r.LargeCustomDictionary.BigQueryField.Field = &dclService.StoredInfoTypeLargeCustomDictionaryBigQueryFieldField{}
							if _, ok := rLargeCustomDictionaryBigQueryFieldField["name"]; ok {
								if s, ok := rLargeCustomDictionaryBigQueryFieldField["name"].(string); ok {
									r.LargeCustomDictionary.BigQueryField.Field.Name = dcl.String(s)
								} else {
									return nil, fmt.Errorf("r.LargeCustomDictionary.BigQueryField.Field.Name: expected string")
								}
							}
						} else {
							return nil, fmt.Errorf("r.LargeCustomDictionary.BigQueryField.Field: expected map[string]interface{}")
						}
					}
					if _, ok := rLargeCustomDictionaryBigQueryField["table"]; ok {
						if rLargeCustomDictionaryBigQueryFieldTable, ok := rLargeCustomDictionaryBigQueryField["table"].(map[string]interface{}); ok {
							r.LargeCustomDictionary.BigQueryField.Table = &dclService.StoredInfoTypeLargeCustomDictionaryBigQueryFieldTable{}
							if _, ok := rLargeCustomDictionaryBigQueryFieldTable["datasetId"]; ok {
								if s, ok := rLargeCustomDictionaryBigQueryFieldTable["datasetId"].(string); ok {
									r.LargeCustomDictionary.BigQueryField.Table.DatasetId = dcl.String(s)
								} else {
									return nil, fmt.Errorf("r.LargeCustomDictionary.BigQueryField.Table.DatasetId: expected string")
								}
							}
							if _, ok := rLargeCustomDictionaryBigQueryFieldTable["projectId"]; ok {
								if s, ok := rLargeCustomDictionaryBigQueryFieldTable["projectId"].(string); ok {
									r.LargeCustomDictionary.BigQueryField.Table.ProjectId = dcl.String(s)
								} else {
									return nil, fmt.Errorf("r.LargeCustomDictionary.BigQueryField.Table.ProjectId: expected string")
								}
							}
							if _, ok := rLargeCustomDictionaryBigQueryFieldTable["tableId"]; ok {
								if s, ok := rLargeCustomDictionaryBigQueryFieldTable["tableId"].(string); ok {
									r.LargeCustomDictionary.BigQueryField.Table.TableId = dcl.String(s)
								} else {
									return nil, fmt.Errorf("r.LargeCustomDictionary.BigQueryField.Table.TableId: expected string")
								}
							}
						} else {
							return nil, fmt.Errorf("r.LargeCustomDictionary.BigQueryField.Table: expected map[string]interface{}")
						}
					}
				} else {
					return nil, fmt.Errorf("r.LargeCustomDictionary.BigQueryField: expected map[string]interface{}")
				}
			}
			if _, ok := rLargeCustomDictionary["cloudStorageFileSet"]; ok {
				if rLargeCustomDictionaryCloudStorageFileSet, ok := rLargeCustomDictionary["cloudStorageFileSet"].(map[string]interface{}); ok {
					r.LargeCustomDictionary.CloudStorageFileSet = &dclService.StoredInfoTypeLargeCustomDictionaryCloudStorageFileSet{}
					if _, ok := rLargeCustomDictionaryCloudStorageFileSet["url"]; ok {
						if s, ok := rLargeCustomDictionaryCloudStorageFileSet["url"].(string); ok {
							r.LargeCustomDictionary.CloudStorageFileSet.Url = dcl.String(s)
						} else {
							return nil, fmt.Errorf("r.LargeCustomDictionary.CloudStorageFileSet.Url: expected string")
						}
					}
				} else {
					return nil, fmt.Errorf("r.LargeCustomDictionary.CloudStorageFileSet: expected map[string]interface{}")
				}
			}
			if _, ok := rLargeCustomDictionary["outputPath"]; ok {
				if rLargeCustomDictionaryOutputPath, ok := rLargeCustomDictionary["outputPath"].(map[string]interface{}); ok {
					r.LargeCustomDictionary.OutputPath = &dclService.StoredInfoTypeLargeCustomDictionaryOutputPath{}
					if _, ok := rLargeCustomDictionaryOutputPath["path"]; ok {
						if s, ok := rLargeCustomDictionaryOutputPath["path"].(string); ok {
							r.LargeCustomDictionary.OutputPath.Path = dcl.String(s)
						} else {
							return nil, fmt.Errorf("r.LargeCustomDictionary.OutputPath.Path: expected string")
						}
					}
				} else {
					return nil, fmt.Errorf("r.LargeCustomDictionary.OutputPath: expected map[string]interface{}")
				}
			}
		} else {
			return nil, fmt.Errorf("r.LargeCustomDictionary: expected map[string]interface{}")
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
	if _, ok := u.Object["parent"]; ok {
		if s, ok := u.Object["parent"].(string); ok {
			r.Parent = dcl.String(s)
		} else {
			return nil, fmt.Errorf("r.Parent: expected string")
		}
	}
	if _, ok := u.Object["regex"]; ok {
		if rRegex, ok := u.Object["regex"].(map[string]interface{}); ok {
			r.Regex = &dclService.StoredInfoTypeRegex{}
			if _, ok := rRegex["groupIndexes"]; ok {
				if s, ok := rRegex["groupIndexes"].([]interface{}); ok {
					for _, ss := range s {
						if intval, ok := ss.(int64); ok {
							r.Regex.GroupIndexes = append(r.Regex.GroupIndexes, intval)
						}
					}
				} else {
					return nil, fmt.Errorf("r.Regex.GroupIndexes: expected []interface{}")
				}
			}
			if _, ok := rRegex["pattern"]; ok {
				if s, ok := rRegex["pattern"].(string); ok {
					r.Regex.Pattern = dcl.String(s)
				} else {
					return nil, fmt.Errorf("r.Regex.Pattern: expected string")
				}
			}
		} else {
			return nil, fmt.Errorf("r.Regex: expected map[string]interface{}")
		}
	}
	return r, nil
}

func GetStoredInfoType(ctx context.Context, config *dcl.Config, u *unstructured.Resource) (*unstructured.Resource, error) {
	c := dclService.NewClient(config)
	r, err := UnstructuredToStoredInfoType(u)
	if err != nil {
		return nil, err
	}
	r, err = c.GetStoredInfoType(ctx, r)
	if err != nil {
		return nil, err
	}
	return StoredInfoTypeToUnstructured(r), nil
}

func ListStoredInfoType(ctx context.Context, config *dcl.Config, location string, parent string) ([]*unstructured.Resource, error) {
	c := dclService.NewClient(config)
	l, err := c.ListStoredInfoType(ctx, location, parent)
	if err != nil {
		return nil, err
	}
	var resources []*unstructured.Resource
	for {
		for _, r := range l.Items {
			resources = append(resources, StoredInfoTypeToUnstructured(r))
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

func ApplyStoredInfoType(ctx context.Context, config *dcl.Config, u *unstructured.Resource, opts ...dcl.ApplyOption) (*unstructured.Resource, error) {
	c := dclService.NewClient(config)
	r, err := UnstructuredToStoredInfoType(u)
	if err != nil {
		return nil, err
	}
	if ush := unstructured.FetchStateHint(opts); ush != nil {
		sh, err := UnstructuredToStoredInfoType(ush)
		if err != nil {
			return nil, err
		}
		opts = append(opts, dcl.WithStateHint(sh))
	}
	r, err = c.ApplyStoredInfoType(ctx, r, opts...)
	if err != nil {
		return nil, err
	}
	return StoredInfoTypeToUnstructured(r), nil
}

func StoredInfoTypeHasDiff(ctx context.Context, config *dcl.Config, u *unstructured.Resource, opts ...dcl.ApplyOption) (bool, error) {
	c := dclService.NewClient(config)
	r, err := UnstructuredToStoredInfoType(u)
	if err != nil {
		return false, err
	}
	if ush := unstructured.FetchStateHint(opts); ush != nil {
		sh, err := UnstructuredToStoredInfoType(ush)
		if err != nil {
			return false, err
		}
		opts = append(opts, dcl.WithStateHint(sh))
	}
	opts = append(opts, dcl.WithLifecycleParam(dcl.BlockDestruction), dcl.WithLifecycleParam(dcl.BlockCreation), dcl.WithLifecycleParam(dcl.BlockModification))
	_, err = c.ApplyStoredInfoType(ctx, r, opts...)
	if err != nil {
		if _, ok := err.(dcl.ApplyInfeasibleError); ok {
			return true, nil
		}
		return false, err
	}
	return false, nil
}

func DeleteStoredInfoType(ctx context.Context, config *dcl.Config, u *unstructured.Resource) error {
	c := dclService.NewClient(config)
	r, err := UnstructuredToStoredInfoType(u)
	if err != nil {
		return err
	}
	return c.DeleteStoredInfoType(ctx, r)
}

func StoredInfoTypeID(u *unstructured.Resource) (string, error) {
	r, err := UnstructuredToStoredInfoType(u)
	if err != nil {
		return "", err
	}
	return r.ID()
}

func (r *StoredInfoType) STV() unstructured.ServiceTypeVersion {
	return unstructured.ServiceTypeVersion{
		"dlp",
		"StoredInfoType",
		"ga",
	}
}

func (r *StoredInfoType) SetPolicyMember(ctx context.Context, config *dcl.Config, resource *unstructured.Resource, member *unstructured.Resource) (*unstructured.Resource, error) {
	return nil, unstructured.ErrNoSuchMethod
}

func (r *StoredInfoType) GetPolicyMember(ctx context.Context, config *dcl.Config, resource *unstructured.Resource, role, member string) (*unstructured.Resource, error) {
	return nil, unstructured.ErrNoSuchMethod
}

func (r *StoredInfoType) DeletePolicyMember(ctx context.Context, config *dcl.Config, resource *unstructured.Resource, member *unstructured.Resource) error {
	return unstructured.ErrNoSuchMethod
}

func (r *StoredInfoType) SetPolicy(ctx context.Context, config *dcl.Config, resource *unstructured.Resource, policy *unstructured.Resource) (*unstructured.Resource, error) {
	return nil, unstructured.ErrNoSuchMethod
}

func (r *StoredInfoType) SetPolicyWithEtag(ctx context.Context, config *dcl.Config, resource *unstructured.Resource, policy *unstructured.Resource) (*unstructured.Resource, error) {
	return nil, unstructured.ErrNoSuchMethod
}

func (r *StoredInfoType) GetPolicy(ctx context.Context, config *dcl.Config, resource *unstructured.Resource) (*unstructured.Resource, error) {
	return nil, unstructured.ErrNoSuchMethod
}

func (r *StoredInfoType) Get(ctx context.Context, config *dcl.Config, resource *unstructured.Resource) (*unstructured.Resource, error) {
	return GetStoredInfoType(ctx, config, resource)
}

func (r *StoredInfoType) Apply(ctx context.Context, config *dcl.Config, resource *unstructured.Resource, opts ...dcl.ApplyOption) (*unstructured.Resource, error) {
	return ApplyStoredInfoType(ctx, config, resource, opts...)
}

func (r *StoredInfoType) HasDiff(ctx context.Context, config *dcl.Config, resource *unstructured.Resource, opts ...dcl.ApplyOption) (bool, error) {
	return StoredInfoTypeHasDiff(ctx, config, resource, opts...)
}

func (r *StoredInfoType) Delete(ctx context.Context, config *dcl.Config, resource *unstructured.Resource) error {
	return DeleteStoredInfoType(ctx, config, resource)
}

func (r *StoredInfoType) ID(resource *unstructured.Resource) (string, error) {
	return StoredInfoTypeID(resource)
}

func init() {
	unstructured.Register(&StoredInfoType{})
}
