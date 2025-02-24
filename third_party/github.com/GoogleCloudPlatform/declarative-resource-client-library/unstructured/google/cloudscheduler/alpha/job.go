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
package cloudscheduler

import (
	"context"
	"fmt"
	"github.com/GoogleCloudPlatform/declarative-resource-client-library/dcl"
	dclService "github.com/GoogleCloudPlatform/declarative-resource-client-library/services/google/cloudscheduler/alpha"
	"github.com/GoogleCloudPlatform/declarative-resource-client-library/unstructured"
)

type Job struct{}

func JobToUnstructured(r *dclService.Job) *unstructured.Resource {
	u := &unstructured.Resource{
		STV: unstructured.ServiceTypeVersion{
			Service: "cloudscheduler",
			Version: "alpha",
			Type:    "Job",
		},
		Object: make(map[string]interface{}),
	}
	if r.AppEngineHttpTarget != nil && r.AppEngineHttpTarget != dclService.EmptyJobAppEngineHttpTarget {
		rAppEngineHttpTarget := make(map[string]interface{})
		if r.AppEngineHttpTarget.AppEngineRouting != nil && r.AppEngineHttpTarget.AppEngineRouting != dclService.EmptyJobAppEngineHttpTargetAppEngineRouting {
			rAppEngineHttpTargetAppEngineRouting := make(map[string]interface{})
			if r.AppEngineHttpTarget.AppEngineRouting.Host != nil {
				rAppEngineHttpTargetAppEngineRouting["host"] = *r.AppEngineHttpTarget.AppEngineRouting.Host
			}
			if r.AppEngineHttpTarget.AppEngineRouting.Instance != nil {
				rAppEngineHttpTargetAppEngineRouting["instance"] = *r.AppEngineHttpTarget.AppEngineRouting.Instance
			}
			if r.AppEngineHttpTarget.AppEngineRouting.Service != nil {
				rAppEngineHttpTargetAppEngineRouting["service"] = *r.AppEngineHttpTarget.AppEngineRouting.Service
			}
			if r.AppEngineHttpTarget.AppEngineRouting.Version != nil {
				rAppEngineHttpTargetAppEngineRouting["version"] = *r.AppEngineHttpTarget.AppEngineRouting.Version
			}
			rAppEngineHttpTarget["appEngineRouting"] = rAppEngineHttpTargetAppEngineRouting
		}
		if r.AppEngineHttpTarget.Body != nil {
			rAppEngineHttpTarget["body"] = *r.AppEngineHttpTarget.Body
		}
		if r.AppEngineHttpTarget.Headers != nil {
			rAppEngineHttpTargetHeaders := make(map[string]interface{})
			for k, v := range r.AppEngineHttpTarget.Headers {
				rAppEngineHttpTargetHeaders[k] = v
			}
			rAppEngineHttpTarget["headers"] = rAppEngineHttpTargetHeaders
		}
		if r.AppEngineHttpTarget.HttpMethod != nil {
			rAppEngineHttpTarget["httpMethod"] = string(*r.AppEngineHttpTarget.HttpMethod)
		}
		if r.AppEngineHttpTarget.RelativeUri != nil {
			rAppEngineHttpTarget["relativeUri"] = *r.AppEngineHttpTarget.RelativeUri
		}
		u.Object["appEngineHttpTarget"] = rAppEngineHttpTarget
	}
	if r.AttemptDeadline != nil {
		u.Object["attemptDeadline"] = *r.AttemptDeadline
	}
	if r.Description != nil {
		u.Object["description"] = *r.Description
	}
	if r.HttpTarget != nil && r.HttpTarget != dclService.EmptyJobHttpTarget {
		rHttpTarget := make(map[string]interface{})
		if r.HttpTarget.Body != nil {
			rHttpTarget["body"] = *r.HttpTarget.Body
		}
		if r.HttpTarget.Headers != nil {
			rHttpTargetHeaders := make(map[string]interface{})
			for k, v := range r.HttpTarget.Headers {
				rHttpTargetHeaders[k] = v
			}
			rHttpTarget["headers"] = rHttpTargetHeaders
		}
		if r.HttpTarget.HttpMethod != nil {
			rHttpTarget["httpMethod"] = string(*r.HttpTarget.HttpMethod)
		}
		if r.HttpTarget.OAuthToken != nil && r.HttpTarget.OAuthToken != dclService.EmptyJobHttpTargetOAuthToken {
			rHttpTargetOAuthToken := make(map[string]interface{})
			if r.HttpTarget.OAuthToken.Scope != nil {
				rHttpTargetOAuthToken["scope"] = *r.HttpTarget.OAuthToken.Scope
			}
			if r.HttpTarget.OAuthToken.ServiceAccountEmail != nil {
				rHttpTargetOAuthToken["serviceAccountEmail"] = *r.HttpTarget.OAuthToken.ServiceAccountEmail
			}
			rHttpTarget["oauthToken"] = rHttpTargetOAuthToken
		}
		if r.HttpTarget.OidcToken != nil && r.HttpTarget.OidcToken != dclService.EmptyJobHttpTargetOidcToken {
			rHttpTargetOidcToken := make(map[string]interface{})
			if r.HttpTarget.OidcToken.Audience != nil {
				rHttpTargetOidcToken["audience"] = *r.HttpTarget.OidcToken.Audience
			}
			if r.HttpTarget.OidcToken.ServiceAccountEmail != nil {
				rHttpTargetOidcToken["serviceAccountEmail"] = *r.HttpTarget.OidcToken.ServiceAccountEmail
			}
			rHttpTarget["oidcToken"] = rHttpTargetOidcToken
		}
		if r.HttpTarget.Uri != nil {
			rHttpTarget["uri"] = *r.HttpTarget.Uri
		}
		u.Object["httpTarget"] = rHttpTarget
	}
	if r.LastAttemptTime != nil {
		u.Object["lastAttemptTime"] = *r.LastAttemptTime
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
	if r.PubsubTarget != nil && r.PubsubTarget != dclService.EmptyJobPubsubTarget {
		rPubsubTarget := make(map[string]interface{})
		if r.PubsubTarget.Attributes != nil {
			rPubsubTargetAttributes := make(map[string]interface{})
			for k, v := range r.PubsubTarget.Attributes {
				rPubsubTargetAttributes[k] = v
			}
			rPubsubTarget["attributes"] = rPubsubTargetAttributes
		}
		if r.PubsubTarget.Data != nil {
			rPubsubTarget["data"] = *r.PubsubTarget.Data
		}
		if r.PubsubTarget.TopicName != nil {
			rPubsubTarget["topicName"] = *r.PubsubTarget.TopicName
		}
		u.Object["pubsubTarget"] = rPubsubTarget
	}
	if r.RetryConfig != nil && r.RetryConfig != dclService.EmptyJobRetryConfig {
		rRetryConfig := make(map[string]interface{})
		if r.RetryConfig.MaxBackoffDuration != nil {
			rRetryConfig["maxBackoffDuration"] = *r.RetryConfig.MaxBackoffDuration
		}
		if r.RetryConfig.MaxDoublings != nil {
			rRetryConfig["maxDoublings"] = *r.RetryConfig.MaxDoublings
		}
		if r.RetryConfig.MaxRetryDuration != nil {
			rRetryConfig["maxRetryDuration"] = *r.RetryConfig.MaxRetryDuration
		}
		if r.RetryConfig.MinBackoffDuration != nil {
			rRetryConfig["minBackoffDuration"] = *r.RetryConfig.MinBackoffDuration
		}
		if r.RetryConfig.RetryCount != nil {
			rRetryConfig["retryCount"] = *r.RetryConfig.RetryCount
		}
		u.Object["retryConfig"] = rRetryConfig
	}
	if r.Schedule != nil {
		u.Object["schedule"] = *r.Schedule
	}
	if r.ScheduleTime != nil {
		u.Object["scheduleTime"] = *r.ScheduleTime
	}
	if r.State != nil {
		u.Object["state"] = string(*r.State)
	}
	if r.Status != nil && r.Status != dclService.EmptyJobStatus {
		rStatus := make(map[string]interface{})
		if r.Status.Code != nil {
			rStatus["code"] = *r.Status.Code
		}
		var rStatusDetails []interface{}
		for _, rStatusDetailsVal := range r.Status.Details {
			rStatusDetailsObject := make(map[string]interface{})
			if rStatusDetailsVal.TypeUrl != nil {
				rStatusDetailsObject["typeUrl"] = *rStatusDetailsVal.TypeUrl
			}
			if rStatusDetailsVal.Value != nil {
				rStatusDetailsObject["value"] = *rStatusDetailsVal.Value
			}
			rStatusDetails = append(rStatusDetails, rStatusDetailsObject)
		}
		rStatus["details"] = rStatusDetails
		if r.Status.Message != nil {
			rStatus["message"] = *r.Status.Message
		}
		u.Object["status"] = rStatus
	}
	if r.TimeZone != nil {
		u.Object["timeZone"] = *r.TimeZone
	}
	if r.UserUpdateTime != nil {
		u.Object["userUpdateTime"] = *r.UserUpdateTime
	}
	return u
}

func UnstructuredToJob(u *unstructured.Resource) (*dclService.Job, error) {
	r := &dclService.Job{}
	if _, ok := u.Object["appEngineHttpTarget"]; ok {
		if rAppEngineHttpTarget, ok := u.Object["appEngineHttpTarget"].(map[string]interface{}); ok {
			r.AppEngineHttpTarget = &dclService.JobAppEngineHttpTarget{}
			if _, ok := rAppEngineHttpTarget["appEngineRouting"]; ok {
				if rAppEngineHttpTargetAppEngineRouting, ok := rAppEngineHttpTarget["appEngineRouting"].(map[string]interface{}); ok {
					r.AppEngineHttpTarget.AppEngineRouting = &dclService.JobAppEngineHttpTargetAppEngineRouting{}
					if _, ok := rAppEngineHttpTargetAppEngineRouting["host"]; ok {
						if s, ok := rAppEngineHttpTargetAppEngineRouting["host"].(string); ok {
							r.AppEngineHttpTarget.AppEngineRouting.Host = dcl.String(s)
						} else {
							return nil, fmt.Errorf("r.AppEngineHttpTarget.AppEngineRouting.Host: expected string")
						}
					}
					if _, ok := rAppEngineHttpTargetAppEngineRouting["instance"]; ok {
						if s, ok := rAppEngineHttpTargetAppEngineRouting["instance"].(string); ok {
							r.AppEngineHttpTarget.AppEngineRouting.Instance = dcl.String(s)
						} else {
							return nil, fmt.Errorf("r.AppEngineHttpTarget.AppEngineRouting.Instance: expected string")
						}
					}
					if _, ok := rAppEngineHttpTargetAppEngineRouting["service"]; ok {
						if s, ok := rAppEngineHttpTargetAppEngineRouting["service"].(string); ok {
							r.AppEngineHttpTarget.AppEngineRouting.Service = dcl.String(s)
						} else {
							return nil, fmt.Errorf("r.AppEngineHttpTarget.AppEngineRouting.Service: expected string")
						}
					}
					if _, ok := rAppEngineHttpTargetAppEngineRouting["version"]; ok {
						if s, ok := rAppEngineHttpTargetAppEngineRouting["version"].(string); ok {
							r.AppEngineHttpTarget.AppEngineRouting.Version = dcl.String(s)
						} else {
							return nil, fmt.Errorf("r.AppEngineHttpTarget.AppEngineRouting.Version: expected string")
						}
					}
				} else {
					return nil, fmt.Errorf("r.AppEngineHttpTarget.AppEngineRouting: expected map[string]interface{}")
				}
			}
			if _, ok := rAppEngineHttpTarget["body"]; ok {
				if s, ok := rAppEngineHttpTarget["body"].(string); ok {
					r.AppEngineHttpTarget.Body = dcl.String(s)
				} else {
					return nil, fmt.Errorf("r.AppEngineHttpTarget.Body: expected string")
				}
			}
			if _, ok := rAppEngineHttpTarget["headers"]; ok {
				if rAppEngineHttpTargetHeaders, ok := rAppEngineHttpTarget["headers"].(map[string]interface{}); ok {
					m := make(map[string]string)
					for k, v := range rAppEngineHttpTargetHeaders {
						if s, ok := v.(string); ok {
							m[k] = s
						}
					}
					r.AppEngineHttpTarget.Headers = m
				} else {
					return nil, fmt.Errorf("r.AppEngineHttpTarget.Headers: expected map[string]interface{}")
				}
			}
			if _, ok := rAppEngineHttpTarget["httpMethod"]; ok {
				if s, ok := rAppEngineHttpTarget["httpMethod"].(string); ok {
					r.AppEngineHttpTarget.HttpMethod = dclService.JobAppEngineHttpTargetHttpMethodEnumRef(s)
				} else {
					return nil, fmt.Errorf("r.AppEngineHttpTarget.HttpMethod: expected string")
				}
			}
			if _, ok := rAppEngineHttpTarget["relativeUri"]; ok {
				if s, ok := rAppEngineHttpTarget["relativeUri"].(string); ok {
					r.AppEngineHttpTarget.RelativeUri = dcl.String(s)
				} else {
					return nil, fmt.Errorf("r.AppEngineHttpTarget.RelativeUri: expected string")
				}
			}
		} else {
			return nil, fmt.Errorf("r.AppEngineHttpTarget: expected map[string]interface{}")
		}
	}
	if _, ok := u.Object["attemptDeadline"]; ok {
		if s, ok := u.Object["attemptDeadline"].(string); ok {
			r.AttemptDeadline = dcl.String(s)
		} else {
			return nil, fmt.Errorf("r.AttemptDeadline: expected string")
		}
	}
	if _, ok := u.Object["description"]; ok {
		if s, ok := u.Object["description"].(string); ok {
			r.Description = dcl.String(s)
		} else {
			return nil, fmt.Errorf("r.Description: expected string")
		}
	}
	if _, ok := u.Object["httpTarget"]; ok {
		if rHttpTarget, ok := u.Object["httpTarget"].(map[string]interface{}); ok {
			r.HttpTarget = &dclService.JobHttpTarget{}
			if _, ok := rHttpTarget["body"]; ok {
				if s, ok := rHttpTarget["body"].(string); ok {
					r.HttpTarget.Body = dcl.String(s)
				} else {
					return nil, fmt.Errorf("r.HttpTarget.Body: expected string")
				}
			}
			if _, ok := rHttpTarget["headers"]; ok {
				if rHttpTargetHeaders, ok := rHttpTarget["headers"].(map[string]interface{}); ok {
					m := make(map[string]string)
					for k, v := range rHttpTargetHeaders {
						if s, ok := v.(string); ok {
							m[k] = s
						}
					}
					r.HttpTarget.Headers = m
				} else {
					return nil, fmt.Errorf("r.HttpTarget.Headers: expected map[string]interface{}")
				}
			}
			if _, ok := rHttpTarget["httpMethod"]; ok {
				if s, ok := rHttpTarget["httpMethod"].(string); ok {
					r.HttpTarget.HttpMethod = dclService.JobHttpTargetHttpMethodEnumRef(s)
				} else {
					return nil, fmt.Errorf("r.HttpTarget.HttpMethod: expected string")
				}
			}
			if _, ok := rHttpTarget["oauthToken"]; ok {
				if rHttpTargetOAuthToken, ok := rHttpTarget["oauthToken"].(map[string]interface{}); ok {
					r.HttpTarget.OAuthToken = &dclService.JobHttpTargetOAuthToken{}
					if _, ok := rHttpTargetOAuthToken["scope"]; ok {
						if s, ok := rHttpTargetOAuthToken["scope"].(string); ok {
							r.HttpTarget.OAuthToken.Scope = dcl.String(s)
						} else {
							return nil, fmt.Errorf("r.HttpTarget.OAuthToken.Scope: expected string")
						}
					}
					if _, ok := rHttpTargetOAuthToken["serviceAccountEmail"]; ok {
						if s, ok := rHttpTargetOAuthToken["serviceAccountEmail"].(string); ok {
							r.HttpTarget.OAuthToken.ServiceAccountEmail = dcl.String(s)
						} else {
							return nil, fmt.Errorf("r.HttpTarget.OAuthToken.ServiceAccountEmail: expected string")
						}
					}
				} else {
					return nil, fmt.Errorf("r.HttpTarget.OAuthToken: expected map[string]interface{}")
				}
			}
			if _, ok := rHttpTarget["oidcToken"]; ok {
				if rHttpTargetOidcToken, ok := rHttpTarget["oidcToken"].(map[string]interface{}); ok {
					r.HttpTarget.OidcToken = &dclService.JobHttpTargetOidcToken{}
					if _, ok := rHttpTargetOidcToken["audience"]; ok {
						if s, ok := rHttpTargetOidcToken["audience"].(string); ok {
							r.HttpTarget.OidcToken.Audience = dcl.String(s)
						} else {
							return nil, fmt.Errorf("r.HttpTarget.OidcToken.Audience: expected string")
						}
					}
					if _, ok := rHttpTargetOidcToken["serviceAccountEmail"]; ok {
						if s, ok := rHttpTargetOidcToken["serviceAccountEmail"].(string); ok {
							r.HttpTarget.OidcToken.ServiceAccountEmail = dcl.String(s)
						} else {
							return nil, fmt.Errorf("r.HttpTarget.OidcToken.ServiceAccountEmail: expected string")
						}
					}
				} else {
					return nil, fmt.Errorf("r.HttpTarget.OidcToken: expected map[string]interface{}")
				}
			}
			if _, ok := rHttpTarget["uri"]; ok {
				if s, ok := rHttpTarget["uri"].(string); ok {
					r.HttpTarget.Uri = dcl.String(s)
				} else {
					return nil, fmt.Errorf("r.HttpTarget.Uri: expected string")
				}
			}
		} else {
			return nil, fmt.Errorf("r.HttpTarget: expected map[string]interface{}")
		}
	}
	if _, ok := u.Object["lastAttemptTime"]; ok {
		if s, ok := u.Object["lastAttemptTime"].(string); ok {
			r.LastAttemptTime = dcl.String(s)
		} else {
			return nil, fmt.Errorf("r.LastAttemptTime: expected string")
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
	if _, ok := u.Object["pubsubTarget"]; ok {
		if rPubsubTarget, ok := u.Object["pubsubTarget"].(map[string]interface{}); ok {
			r.PubsubTarget = &dclService.JobPubsubTarget{}
			if _, ok := rPubsubTarget["attributes"]; ok {
				if rPubsubTargetAttributes, ok := rPubsubTarget["attributes"].(map[string]interface{}); ok {
					m := make(map[string]string)
					for k, v := range rPubsubTargetAttributes {
						if s, ok := v.(string); ok {
							m[k] = s
						}
					}
					r.PubsubTarget.Attributes = m
				} else {
					return nil, fmt.Errorf("r.PubsubTarget.Attributes: expected map[string]interface{}")
				}
			}
			if _, ok := rPubsubTarget["data"]; ok {
				if s, ok := rPubsubTarget["data"].(string); ok {
					r.PubsubTarget.Data = dcl.String(s)
				} else {
					return nil, fmt.Errorf("r.PubsubTarget.Data: expected string")
				}
			}
			if _, ok := rPubsubTarget["topicName"]; ok {
				if s, ok := rPubsubTarget["topicName"].(string); ok {
					r.PubsubTarget.TopicName = dcl.String(s)
				} else {
					return nil, fmt.Errorf("r.PubsubTarget.TopicName: expected string")
				}
			}
		} else {
			return nil, fmt.Errorf("r.PubsubTarget: expected map[string]interface{}")
		}
	}
	if _, ok := u.Object["retryConfig"]; ok {
		if rRetryConfig, ok := u.Object["retryConfig"].(map[string]interface{}); ok {
			r.RetryConfig = &dclService.JobRetryConfig{}
			if _, ok := rRetryConfig["maxBackoffDuration"]; ok {
				if s, ok := rRetryConfig["maxBackoffDuration"].(string); ok {
					r.RetryConfig.MaxBackoffDuration = dcl.String(s)
				} else {
					return nil, fmt.Errorf("r.RetryConfig.MaxBackoffDuration: expected string")
				}
			}
			if _, ok := rRetryConfig["maxDoublings"]; ok {
				if i, ok := rRetryConfig["maxDoublings"].(int64); ok {
					r.RetryConfig.MaxDoublings = dcl.Int64(i)
				} else {
					return nil, fmt.Errorf("r.RetryConfig.MaxDoublings: expected int64")
				}
			}
			if _, ok := rRetryConfig["maxRetryDuration"]; ok {
				if s, ok := rRetryConfig["maxRetryDuration"].(string); ok {
					r.RetryConfig.MaxRetryDuration = dcl.String(s)
				} else {
					return nil, fmt.Errorf("r.RetryConfig.MaxRetryDuration: expected string")
				}
			}
			if _, ok := rRetryConfig["minBackoffDuration"]; ok {
				if s, ok := rRetryConfig["minBackoffDuration"].(string); ok {
					r.RetryConfig.MinBackoffDuration = dcl.String(s)
				} else {
					return nil, fmt.Errorf("r.RetryConfig.MinBackoffDuration: expected string")
				}
			}
			if _, ok := rRetryConfig["retryCount"]; ok {
				if i, ok := rRetryConfig["retryCount"].(int64); ok {
					r.RetryConfig.RetryCount = dcl.Int64(i)
				} else {
					return nil, fmt.Errorf("r.RetryConfig.RetryCount: expected int64")
				}
			}
		} else {
			return nil, fmt.Errorf("r.RetryConfig: expected map[string]interface{}")
		}
	}
	if _, ok := u.Object["schedule"]; ok {
		if s, ok := u.Object["schedule"].(string); ok {
			r.Schedule = dcl.String(s)
		} else {
			return nil, fmt.Errorf("r.Schedule: expected string")
		}
	}
	if _, ok := u.Object["scheduleTime"]; ok {
		if s, ok := u.Object["scheduleTime"].(string); ok {
			r.ScheduleTime = dcl.String(s)
		} else {
			return nil, fmt.Errorf("r.ScheduleTime: expected string")
		}
	}
	if _, ok := u.Object["state"]; ok {
		if s, ok := u.Object["state"].(string); ok {
			r.State = dclService.JobStateEnumRef(s)
		} else {
			return nil, fmt.Errorf("r.State: expected string")
		}
	}
	if _, ok := u.Object["status"]; ok {
		if rStatus, ok := u.Object["status"].(map[string]interface{}); ok {
			r.Status = &dclService.JobStatus{}
			if _, ok := rStatus["code"]; ok {
				if i, ok := rStatus["code"].(int64); ok {
					r.Status.Code = dcl.Int64(i)
				} else {
					return nil, fmt.Errorf("r.Status.Code: expected int64")
				}
			}
			if _, ok := rStatus["details"]; ok {
				if s, ok := rStatus["details"].([]interface{}); ok {
					for _, o := range s {
						if objval, ok := o.(map[string]interface{}); ok {
							var rStatusDetails dclService.JobStatusDetails
							if _, ok := objval["typeUrl"]; ok {
								if s, ok := objval["typeUrl"].(string); ok {
									rStatusDetails.TypeUrl = dcl.String(s)
								} else {
									return nil, fmt.Errorf("rStatusDetails.TypeUrl: expected string")
								}
							}
							if _, ok := objval["value"]; ok {
								if s, ok := objval["value"].(string); ok {
									rStatusDetails.Value = dcl.String(s)
								} else {
									return nil, fmt.Errorf("rStatusDetails.Value: expected string")
								}
							}
							r.Status.Details = append(r.Status.Details, rStatusDetails)
						}
					}
				} else {
					return nil, fmt.Errorf("r.Status.Details: expected []interface{}")
				}
			}
			if _, ok := rStatus["message"]; ok {
				if s, ok := rStatus["message"].(string); ok {
					r.Status.Message = dcl.String(s)
				} else {
					return nil, fmt.Errorf("r.Status.Message: expected string")
				}
			}
		} else {
			return nil, fmt.Errorf("r.Status: expected map[string]interface{}")
		}
	}
	if _, ok := u.Object["timeZone"]; ok {
		if s, ok := u.Object["timeZone"].(string); ok {
			r.TimeZone = dcl.String(s)
		} else {
			return nil, fmt.Errorf("r.TimeZone: expected string")
		}
	}
	if _, ok := u.Object["userUpdateTime"]; ok {
		if s, ok := u.Object["userUpdateTime"].(string); ok {
			r.UserUpdateTime = dcl.String(s)
		} else {
			return nil, fmt.Errorf("r.UserUpdateTime: expected string")
		}
	}
	return r, nil
}

func GetJob(ctx context.Context, config *dcl.Config, u *unstructured.Resource) (*unstructured.Resource, error) {
	c := dclService.NewClient(config)
	r, err := UnstructuredToJob(u)
	if err != nil {
		return nil, err
	}
	r, err = c.GetJob(ctx, r)
	if err != nil {
		return nil, err
	}
	return JobToUnstructured(r), nil
}

func ListJob(ctx context.Context, config *dcl.Config, project string, location string) ([]*unstructured.Resource, error) {
	c := dclService.NewClient(config)
	l, err := c.ListJob(ctx, project, location)
	if err != nil {
		return nil, err
	}
	var resources []*unstructured.Resource
	for {
		for _, r := range l.Items {
			resources = append(resources, JobToUnstructured(r))
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

func ApplyJob(ctx context.Context, config *dcl.Config, u *unstructured.Resource, opts ...dcl.ApplyOption) (*unstructured.Resource, error) {
	c := dclService.NewClient(config)
	r, err := UnstructuredToJob(u)
	if err != nil {
		return nil, err
	}
	if ush := unstructured.FetchStateHint(opts); ush != nil {
		sh, err := UnstructuredToJob(ush)
		if err != nil {
			return nil, err
		}
		opts = append(opts, dcl.WithStateHint(sh))
	}
	r, err = c.ApplyJob(ctx, r, opts...)
	if err != nil {
		return nil, err
	}
	return JobToUnstructured(r), nil
}

func JobHasDiff(ctx context.Context, config *dcl.Config, u *unstructured.Resource, opts ...dcl.ApplyOption) (bool, error) {
	c := dclService.NewClient(config)
	r, err := UnstructuredToJob(u)
	if err != nil {
		return false, err
	}
	if ush := unstructured.FetchStateHint(opts); ush != nil {
		sh, err := UnstructuredToJob(ush)
		if err != nil {
			return false, err
		}
		opts = append(opts, dcl.WithStateHint(sh))
	}
	opts = append(opts, dcl.WithLifecycleParam(dcl.BlockDestruction), dcl.WithLifecycleParam(dcl.BlockCreation), dcl.WithLifecycleParam(dcl.BlockModification))
	_, err = c.ApplyJob(ctx, r, opts...)
	if err != nil {
		if _, ok := err.(dcl.ApplyInfeasibleError); ok {
			return true, nil
		}
		return false, err
	}
	return false, nil
}

func DeleteJob(ctx context.Context, config *dcl.Config, u *unstructured.Resource) error {
	c := dclService.NewClient(config)
	r, err := UnstructuredToJob(u)
	if err != nil {
		return err
	}
	return c.DeleteJob(ctx, r)
}

func JobID(u *unstructured.Resource) (string, error) {
	r, err := UnstructuredToJob(u)
	if err != nil {
		return "", err
	}
	return r.ID()
}

func (r *Job) STV() unstructured.ServiceTypeVersion {
	return unstructured.ServiceTypeVersion{
		"cloudscheduler",
		"Job",
		"alpha",
	}
}

func (r *Job) SetPolicyMember(ctx context.Context, config *dcl.Config, resource *unstructured.Resource, member *unstructured.Resource) (*unstructured.Resource, error) {
	return nil, unstructured.ErrNoSuchMethod
}

func (r *Job) GetPolicyMember(ctx context.Context, config *dcl.Config, resource *unstructured.Resource, role, member string) (*unstructured.Resource, error) {
	return nil, unstructured.ErrNoSuchMethod
}

func (r *Job) DeletePolicyMember(ctx context.Context, config *dcl.Config, resource *unstructured.Resource, member *unstructured.Resource) error {
	return unstructured.ErrNoSuchMethod
}

func (r *Job) SetPolicy(ctx context.Context, config *dcl.Config, resource *unstructured.Resource, policy *unstructured.Resource) (*unstructured.Resource, error) {
	return nil, unstructured.ErrNoSuchMethod
}

func (r *Job) SetPolicyWithEtag(ctx context.Context, config *dcl.Config, resource *unstructured.Resource, policy *unstructured.Resource) (*unstructured.Resource, error) {
	return nil, unstructured.ErrNoSuchMethod
}

func (r *Job) GetPolicy(ctx context.Context, config *dcl.Config, resource *unstructured.Resource) (*unstructured.Resource, error) {
	return nil, unstructured.ErrNoSuchMethod
}

func (r *Job) Get(ctx context.Context, config *dcl.Config, resource *unstructured.Resource) (*unstructured.Resource, error) {
	return GetJob(ctx, config, resource)
}

func (r *Job) Apply(ctx context.Context, config *dcl.Config, resource *unstructured.Resource, opts ...dcl.ApplyOption) (*unstructured.Resource, error) {
	return ApplyJob(ctx, config, resource, opts...)
}

func (r *Job) HasDiff(ctx context.Context, config *dcl.Config, resource *unstructured.Resource, opts ...dcl.ApplyOption) (bool, error) {
	return JobHasDiff(ctx, config, resource, opts...)
}

func (r *Job) Delete(ctx context.Context, config *dcl.Config, resource *unstructured.Resource) error {
	return DeleteJob(ctx, config, resource)
}

func (r *Job) ID(resource *unstructured.Resource) (string, error) {
	return JobID(resource)
}

func init() {
	unstructured.Register(&Job{})
}
