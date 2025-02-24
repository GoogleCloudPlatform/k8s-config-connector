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
package bigquery

import (
	"context"
	"fmt"
	"github.com/GoogleCloudPlatform/declarative-resource-client-library/dcl"
	dclService "github.com/GoogleCloudPlatform/declarative-resource-client-library/services/google/bigquery/beta"
	"github.com/GoogleCloudPlatform/declarative-resource-client-library/unstructured"
)

type Dataset struct{}

func DatasetToUnstructured(r *dclService.Dataset) *unstructured.Resource {
	u := &unstructured.Resource{
		STV: unstructured.ServiceTypeVersion{
			Service: "bigquery",
			Version: "beta",
			Type:    "Dataset",
		},
		Object: make(map[string]interface{}),
	}
	var rAccess []interface{}
	for _, rAccessVal := range r.Access {
		rAccessObject := make(map[string]interface{})
		if rAccessVal.Domain != nil {
			rAccessObject["domain"] = *rAccessVal.Domain
		}
		if rAccessVal.GroupByEmail != nil {
			rAccessObject["groupByEmail"] = *rAccessVal.GroupByEmail
		}
		if rAccessVal.IamMember != nil {
			rAccessObject["iamMember"] = *rAccessVal.IamMember
		}
		if rAccessVal.Role != nil {
			rAccessObject["role"] = *rAccessVal.Role
		}
		if rAccessVal.Routine != nil && rAccessVal.Routine != dclService.EmptyDatasetAccessRoutine {
			rAccessValRoutine := make(map[string]interface{})
			if rAccessVal.Routine.DatasetId != nil {
				rAccessValRoutine["datasetId"] = *rAccessVal.Routine.DatasetId
			}
			if rAccessVal.Routine.ProjectId != nil {
				rAccessValRoutine["projectId"] = *rAccessVal.Routine.ProjectId
			}
			if rAccessVal.Routine.RoutineId != nil {
				rAccessValRoutine["routineId"] = *rAccessVal.Routine.RoutineId
			}
			rAccessObject["routine"] = rAccessValRoutine
		}
		if rAccessVal.SpecialGroup != nil {
			rAccessObject["specialGroup"] = *rAccessVal.SpecialGroup
		}
		if rAccessVal.UserByEmail != nil {
			rAccessObject["userByEmail"] = *rAccessVal.UserByEmail
		}
		if rAccessVal.View != nil && rAccessVal.View != dclService.EmptyDatasetAccessView {
			rAccessValView := make(map[string]interface{})
			if rAccessVal.View.DatasetId != nil {
				rAccessValView["datasetId"] = *rAccessVal.View.DatasetId
			}
			if rAccessVal.View.ProjectId != nil {
				rAccessValView["projectId"] = *rAccessVal.View.ProjectId
			}
			if rAccessVal.View.TableId != nil {
				rAccessValView["tableId"] = *rAccessVal.View.TableId
			}
			rAccessObject["view"] = rAccessValView
		}
		rAccess = append(rAccess, rAccessObject)
	}
	u.Object["access"] = rAccess
	if r.CreationTime != nil {
		u.Object["creationTime"] = *r.CreationTime
	}
	if r.DefaultEncryptionConfiguration != nil && r.DefaultEncryptionConfiguration != dclService.EmptyDatasetDefaultEncryptionConfiguration {
		rDefaultEncryptionConfiguration := make(map[string]interface{})
		if r.DefaultEncryptionConfiguration.KmsKeyName != nil {
			rDefaultEncryptionConfiguration["kmsKeyName"] = *r.DefaultEncryptionConfiguration.KmsKeyName
		}
		u.Object["defaultEncryptionConfiguration"] = rDefaultEncryptionConfiguration
	}
	if r.DefaultPartitionExpirationMs != nil {
		u.Object["defaultPartitionExpirationMs"] = *r.DefaultPartitionExpirationMs
	}
	if r.DefaultTableExpirationMs != nil {
		u.Object["defaultTableExpirationMs"] = *r.DefaultTableExpirationMs
	}
	if r.Description != nil {
		u.Object["description"] = *r.Description
	}
	if r.Etag != nil {
		u.Object["etag"] = *r.Etag
	}
	if r.FriendlyName != nil {
		u.Object["friendlyName"] = *r.FriendlyName
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
	if r.LastModifiedTime != nil {
		u.Object["lastModifiedTime"] = *r.LastModifiedTime
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
	if r.Published != nil {
		u.Object["published"] = *r.Published
	}
	if r.SelfLink != nil {
		u.Object["selfLink"] = *r.SelfLink
	}
	return u
}

func UnstructuredToDataset(u *unstructured.Resource) (*dclService.Dataset, error) {
	r := &dclService.Dataset{}
	if _, ok := u.Object["access"]; ok {
		if s, ok := u.Object["access"].([]interface{}); ok {
			for _, o := range s {
				if objval, ok := o.(map[string]interface{}); ok {
					var rAccess dclService.DatasetAccess
					if _, ok := objval["domain"]; ok {
						if s, ok := objval["domain"].(string); ok {
							rAccess.Domain = dcl.String(s)
						} else {
							return nil, fmt.Errorf("rAccess.Domain: expected string")
						}
					}
					if _, ok := objval["groupByEmail"]; ok {
						if s, ok := objval["groupByEmail"].(string); ok {
							rAccess.GroupByEmail = dcl.String(s)
						} else {
							return nil, fmt.Errorf("rAccess.GroupByEmail: expected string")
						}
					}
					if _, ok := objval["iamMember"]; ok {
						if s, ok := objval["iamMember"].(string); ok {
							rAccess.IamMember = dcl.String(s)
						} else {
							return nil, fmt.Errorf("rAccess.IamMember: expected string")
						}
					}
					if _, ok := objval["role"]; ok {
						if s, ok := objval["role"].(string); ok {
							rAccess.Role = dcl.String(s)
						} else {
							return nil, fmt.Errorf("rAccess.Role: expected string")
						}
					}
					if _, ok := objval["routine"]; ok {
						if rAccessRoutine, ok := objval["routine"].(map[string]interface{}); ok {
							rAccess.Routine = &dclService.DatasetAccessRoutine{}
							if _, ok := rAccessRoutine["datasetId"]; ok {
								if s, ok := rAccessRoutine["datasetId"].(string); ok {
									rAccess.Routine.DatasetId = dcl.String(s)
								} else {
									return nil, fmt.Errorf("rAccess.Routine.DatasetId: expected string")
								}
							}
							if _, ok := rAccessRoutine["projectId"]; ok {
								if s, ok := rAccessRoutine["projectId"].(string); ok {
									rAccess.Routine.ProjectId = dcl.String(s)
								} else {
									return nil, fmt.Errorf("rAccess.Routine.ProjectId: expected string")
								}
							}
							if _, ok := rAccessRoutine["routineId"]; ok {
								if s, ok := rAccessRoutine["routineId"].(string); ok {
									rAccess.Routine.RoutineId = dcl.String(s)
								} else {
									return nil, fmt.Errorf("rAccess.Routine.RoutineId: expected string")
								}
							}
						} else {
							return nil, fmt.Errorf("rAccess.Routine: expected map[string]interface{}")
						}
					}
					if _, ok := objval["specialGroup"]; ok {
						if s, ok := objval["specialGroup"].(string); ok {
							rAccess.SpecialGroup = dcl.String(s)
						} else {
							return nil, fmt.Errorf("rAccess.SpecialGroup: expected string")
						}
					}
					if _, ok := objval["userByEmail"]; ok {
						if s, ok := objval["userByEmail"].(string); ok {
							rAccess.UserByEmail = dcl.String(s)
						} else {
							return nil, fmt.Errorf("rAccess.UserByEmail: expected string")
						}
					}
					if _, ok := objval["view"]; ok {
						if rAccessView, ok := objval["view"].(map[string]interface{}); ok {
							rAccess.View = &dclService.DatasetAccessView{}
							if _, ok := rAccessView["datasetId"]; ok {
								if s, ok := rAccessView["datasetId"].(string); ok {
									rAccess.View.DatasetId = dcl.String(s)
								} else {
									return nil, fmt.Errorf("rAccess.View.DatasetId: expected string")
								}
							}
							if _, ok := rAccessView["projectId"]; ok {
								if s, ok := rAccessView["projectId"].(string); ok {
									rAccess.View.ProjectId = dcl.String(s)
								} else {
									return nil, fmt.Errorf("rAccess.View.ProjectId: expected string")
								}
							}
							if _, ok := rAccessView["tableId"]; ok {
								if s, ok := rAccessView["tableId"].(string); ok {
									rAccess.View.TableId = dcl.String(s)
								} else {
									return nil, fmt.Errorf("rAccess.View.TableId: expected string")
								}
							}
						} else {
							return nil, fmt.Errorf("rAccess.View: expected map[string]interface{}")
						}
					}
					r.Access = append(r.Access, rAccess)
				}
			}
		} else {
			return nil, fmt.Errorf("r.Access: expected []interface{}")
		}
	}
	if _, ok := u.Object["creationTime"]; ok {
		if i, ok := u.Object["creationTime"].(int64); ok {
			r.CreationTime = dcl.Int64(i)
		} else {
			return nil, fmt.Errorf("r.CreationTime: expected int64")
		}
	}
	if _, ok := u.Object["defaultEncryptionConfiguration"]; ok {
		if rDefaultEncryptionConfiguration, ok := u.Object["defaultEncryptionConfiguration"].(map[string]interface{}); ok {
			r.DefaultEncryptionConfiguration = &dclService.DatasetDefaultEncryptionConfiguration{}
			if _, ok := rDefaultEncryptionConfiguration["kmsKeyName"]; ok {
				if s, ok := rDefaultEncryptionConfiguration["kmsKeyName"].(string); ok {
					r.DefaultEncryptionConfiguration.KmsKeyName = dcl.String(s)
				} else {
					return nil, fmt.Errorf("r.DefaultEncryptionConfiguration.KmsKeyName: expected string")
				}
			}
		} else {
			return nil, fmt.Errorf("r.DefaultEncryptionConfiguration: expected map[string]interface{}")
		}
	}
	if _, ok := u.Object["defaultPartitionExpirationMs"]; ok {
		if s, ok := u.Object["defaultPartitionExpirationMs"].(string); ok {
			r.DefaultPartitionExpirationMs = dcl.String(s)
		} else {
			return nil, fmt.Errorf("r.DefaultPartitionExpirationMs: expected string")
		}
	}
	if _, ok := u.Object["defaultTableExpirationMs"]; ok {
		if s, ok := u.Object["defaultTableExpirationMs"].(string); ok {
			r.DefaultTableExpirationMs = dcl.String(s)
		} else {
			return nil, fmt.Errorf("r.DefaultTableExpirationMs: expected string")
		}
	}
	if _, ok := u.Object["description"]; ok {
		if s, ok := u.Object["description"].(string); ok {
			r.Description = dcl.String(s)
		} else {
			return nil, fmt.Errorf("r.Description: expected string")
		}
	}
	if _, ok := u.Object["etag"]; ok {
		if s, ok := u.Object["etag"].(string); ok {
			r.Etag = dcl.String(s)
		} else {
			return nil, fmt.Errorf("r.Etag: expected string")
		}
	}
	if _, ok := u.Object["friendlyName"]; ok {
		if s, ok := u.Object["friendlyName"].(string); ok {
			r.FriendlyName = dcl.String(s)
		} else {
			return nil, fmt.Errorf("r.FriendlyName: expected string")
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
	if _, ok := u.Object["lastModifiedTime"]; ok {
		if i, ok := u.Object["lastModifiedTime"].(int64); ok {
			r.LastModifiedTime = dcl.Int64(i)
		} else {
			return nil, fmt.Errorf("r.LastModifiedTime: expected int64")
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
	if _, ok := u.Object["published"]; ok {
		if b, ok := u.Object["published"].(bool); ok {
			r.Published = dcl.Bool(b)
		} else {
			return nil, fmt.Errorf("r.Published: expected bool")
		}
	}
	if _, ok := u.Object["selfLink"]; ok {
		if s, ok := u.Object["selfLink"].(string); ok {
			r.SelfLink = dcl.String(s)
		} else {
			return nil, fmt.Errorf("r.SelfLink: expected string")
		}
	}
	return r, nil
}

func GetDataset(ctx context.Context, config *dcl.Config, u *unstructured.Resource) (*unstructured.Resource, error) {
	c := dclService.NewClient(config)
	r, err := UnstructuredToDataset(u)
	if err != nil {
		return nil, err
	}
	r, err = c.GetDataset(ctx, r)
	if err != nil {
		return nil, err
	}
	return DatasetToUnstructured(r), nil
}

func ListDataset(ctx context.Context, config *dcl.Config, project string) ([]*unstructured.Resource, error) {
	c := dclService.NewClient(config)
	l, err := c.ListDataset(ctx, project)
	if err != nil {
		return nil, err
	}
	var resources []*unstructured.Resource
	for {
		for _, r := range l.Items {
			resources = append(resources, DatasetToUnstructured(r))
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

func ApplyDataset(ctx context.Context, config *dcl.Config, u *unstructured.Resource, opts ...dcl.ApplyOption) (*unstructured.Resource, error) {
	c := dclService.NewClient(config)
	r, err := UnstructuredToDataset(u)
	if err != nil {
		return nil, err
	}
	if ush := unstructured.FetchStateHint(opts); ush != nil {
		sh, err := UnstructuredToDataset(ush)
		if err != nil {
			return nil, err
		}
		opts = append(opts, dcl.WithStateHint(sh))
	}
	r, err = c.ApplyDataset(ctx, r, opts...)
	if err != nil {
		return nil, err
	}
	return DatasetToUnstructured(r), nil
}

func DatasetHasDiff(ctx context.Context, config *dcl.Config, u *unstructured.Resource, opts ...dcl.ApplyOption) (bool, error) {
	c := dclService.NewClient(config)
	r, err := UnstructuredToDataset(u)
	if err != nil {
		return false, err
	}
	if ush := unstructured.FetchStateHint(opts); ush != nil {
		sh, err := UnstructuredToDataset(ush)
		if err != nil {
			return false, err
		}
		opts = append(opts, dcl.WithStateHint(sh))
	}
	opts = append(opts, dcl.WithLifecycleParam(dcl.BlockDestruction), dcl.WithLifecycleParam(dcl.BlockCreation), dcl.WithLifecycleParam(dcl.BlockModification))
	_, err = c.ApplyDataset(ctx, r, opts...)
	if err != nil {
		if _, ok := err.(dcl.ApplyInfeasibleError); ok {
			return true, nil
		}
		return false, err
	}
	return false, nil
}

func DeleteDataset(ctx context.Context, config *dcl.Config, u *unstructured.Resource) error {
	c := dclService.NewClient(config)
	r, err := UnstructuredToDataset(u)
	if err != nil {
		return err
	}
	return c.DeleteDataset(ctx, r)
}

func DatasetID(u *unstructured.Resource) (string, error) {
	r, err := UnstructuredToDataset(u)
	if err != nil {
		return "", err
	}
	return r.ID()
}

func (r *Dataset) STV() unstructured.ServiceTypeVersion {
	return unstructured.ServiceTypeVersion{
		"bigquery",
		"Dataset",
		"beta",
	}
}

func (r *Dataset) SetPolicyMember(ctx context.Context, config *dcl.Config, resource *unstructured.Resource, member *unstructured.Resource) (*unstructured.Resource, error) {
	return nil, unstructured.ErrNoSuchMethod
}

func (r *Dataset) GetPolicyMember(ctx context.Context, config *dcl.Config, resource *unstructured.Resource, role, member string) (*unstructured.Resource, error) {
	return nil, unstructured.ErrNoSuchMethod
}

func (r *Dataset) DeletePolicyMember(ctx context.Context, config *dcl.Config, resource *unstructured.Resource, member *unstructured.Resource) error {
	return unstructured.ErrNoSuchMethod
}

func (r *Dataset) SetPolicy(ctx context.Context, config *dcl.Config, resource *unstructured.Resource, policy *unstructured.Resource) (*unstructured.Resource, error) {
	return nil, unstructured.ErrNoSuchMethod
}

func (r *Dataset) SetPolicyWithEtag(ctx context.Context, config *dcl.Config, resource *unstructured.Resource, policy *unstructured.Resource) (*unstructured.Resource, error) {
	return nil, unstructured.ErrNoSuchMethod
}

func (r *Dataset) GetPolicy(ctx context.Context, config *dcl.Config, resource *unstructured.Resource) (*unstructured.Resource, error) {
	return nil, unstructured.ErrNoSuchMethod
}

func (r *Dataset) Get(ctx context.Context, config *dcl.Config, resource *unstructured.Resource) (*unstructured.Resource, error) {
	return GetDataset(ctx, config, resource)
}

func (r *Dataset) Apply(ctx context.Context, config *dcl.Config, resource *unstructured.Resource, opts ...dcl.ApplyOption) (*unstructured.Resource, error) {
	return ApplyDataset(ctx, config, resource, opts...)
}

func (r *Dataset) HasDiff(ctx context.Context, config *dcl.Config, resource *unstructured.Resource, opts ...dcl.ApplyOption) (bool, error) {
	return DatasetHasDiff(ctx, config, resource, opts...)
}

func (r *Dataset) Delete(ctx context.Context, config *dcl.Config, resource *unstructured.Resource) error {
	return DeleteDataset(ctx, config, resource)
}

func (r *Dataset) ID(resource *unstructured.Resource) (string, error) {
	return DatasetID(resource)
}

func init() {
	unstructured.Register(&Dataset{})
}
