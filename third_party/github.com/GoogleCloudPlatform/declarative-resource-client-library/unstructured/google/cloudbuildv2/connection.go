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
package cloudbuildv2

import (
	"context"
	"fmt"
	"github.com/GoogleCloudPlatform/declarative-resource-client-library/dcl"
	dclService "github.com/GoogleCloudPlatform/declarative-resource-client-library/services/google/cloudbuildv2"
	"github.com/GoogleCloudPlatform/declarative-resource-client-library/unstructured"
)

type Connection struct{}

func ConnectionToUnstructured(r *dclService.Connection) *unstructured.Resource {
	u := &unstructured.Resource{
		STV: unstructured.ServiceTypeVersion{
			Service: "cloudbuildv2",
			Version: "ga",
			Type:    "Connection",
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
	if r.CreateTime != nil {
		u.Object["createTime"] = *r.CreateTime
	}
	if r.Disabled != nil {
		u.Object["disabled"] = *r.Disabled
	}
	if r.Etag != nil {
		u.Object["etag"] = *r.Etag
	}
	if r.GithubConfig != nil && r.GithubConfig != dclService.EmptyConnectionGithubConfig {
		rGithubConfig := make(map[string]interface{})
		if r.GithubConfig.AppInstallationId != nil {
			rGithubConfig["appInstallationId"] = *r.GithubConfig.AppInstallationId
		}
		if r.GithubConfig.AuthorizerCredential != nil && r.GithubConfig.AuthorizerCredential != dclService.EmptyConnectionGithubConfigAuthorizerCredential {
			rGithubConfigAuthorizerCredential := make(map[string]interface{})
			if r.GithubConfig.AuthorizerCredential.OAuthTokenSecretVersion != nil {
				rGithubConfigAuthorizerCredential["oauthTokenSecretVersion"] = *r.GithubConfig.AuthorizerCredential.OAuthTokenSecretVersion
			}
			if r.GithubConfig.AuthorizerCredential.Username != nil {
				rGithubConfigAuthorizerCredential["username"] = *r.GithubConfig.AuthorizerCredential.Username
			}
			rGithubConfig["authorizerCredential"] = rGithubConfigAuthorizerCredential
		}
		u.Object["githubConfig"] = rGithubConfig
	}
	if r.GithubEnterpriseConfig != nil && r.GithubEnterpriseConfig != dclService.EmptyConnectionGithubEnterpriseConfig {
		rGithubEnterpriseConfig := make(map[string]interface{})
		if r.GithubEnterpriseConfig.AppId != nil {
			rGithubEnterpriseConfig["appId"] = *r.GithubEnterpriseConfig.AppId
		}
		if r.GithubEnterpriseConfig.AppInstallationId != nil {
			rGithubEnterpriseConfig["appInstallationId"] = *r.GithubEnterpriseConfig.AppInstallationId
		}
		if r.GithubEnterpriseConfig.AppSlug != nil {
			rGithubEnterpriseConfig["appSlug"] = *r.GithubEnterpriseConfig.AppSlug
		}
		if r.GithubEnterpriseConfig.HostUri != nil {
			rGithubEnterpriseConfig["hostUri"] = *r.GithubEnterpriseConfig.HostUri
		}
		if r.GithubEnterpriseConfig.PrivateKeySecretVersion != nil {
			rGithubEnterpriseConfig["privateKeySecretVersion"] = *r.GithubEnterpriseConfig.PrivateKeySecretVersion
		}
		if r.GithubEnterpriseConfig.ServiceDirectoryConfig != nil && r.GithubEnterpriseConfig.ServiceDirectoryConfig != dclService.EmptyConnectionGithubEnterpriseConfigServiceDirectoryConfig {
			rGithubEnterpriseConfigServiceDirectoryConfig := make(map[string]interface{})
			if r.GithubEnterpriseConfig.ServiceDirectoryConfig.Service != nil {
				rGithubEnterpriseConfigServiceDirectoryConfig["service"] = *r.GithubEnterpriseConfig.ServiceDirectoryConfig.Service
			}
			rGithubEnterpriseConfig["serviceDirectoryConfig"] = rGithubEnterpriseConfigServiceDirectoryConfig
		}
		if r.GithubEnterpriseConfig.SslCa != nil {
			rGithubEnterpriseConfig["sslCa"] = *r.GithubEnterpriseConfig.SslCa
		}
		if r.GithubEnterpriseConfig.WebhookSecretSecretVersion != nil {
			rGithubEnterpriseConfig["webhookSecretSecretVersion"] = *r.GithubEnterpriseConfig.WebhookSecretSecretVersion
		}
		u.Object["githubEnterpriseConfig"] = rGithubEnterpriseConfig
	}
	if r.GitlabConfig != nil && r.GitlabConfig != dclService.EmptyConnectionGitlabConfig {
		rGitlabConfig := make(map[string]interface{})
		if r.GitlabConfig.AuthorizerCredential != nil && r.GitlabConfig.AuthorizerCredential != dclService.EmptyConnectionGitlabConfigAuthorizerCredential {
			rGitlabConfigAuthorizerCredential := make(map[string]interface{})
			if r.GitlabConfig.AuthorizerCredential.UserTokenSecretVersion != nil {
				rGitlabConfigAuthorizerCredential["userTokenSecretVersion"] = *r.GitlabConfig.AuthorizerCredential.UserTokenSecretVersion
			}
			if r.GitlabConfig.AuthorizerCredential.Username != nil {
				rGitlabConfigAuthorizerCredential["username"] = *r.GitlabConfig.AuthorizerCredential.Username
			}
			rGitlabConfig["authorizerCredential"] = rGitlabConfigAuthorizerCredential
		}
		if r.GitlabConfig.HostUri != nil {
			rGitlabConfig["hostUri"] = *r.GitlabConfig.HostUri
		}
		if r.GitlabConfig.ReadAuthorizerCredential != nil && r.GitlabConfig.ReadAuthorizerCredential != dclService.EmptyConnectionGitlabConfigReadAuthorizerCredential {
			rGitlabConfigReadAuthorizerCredential := make(map[string]interface{})
			if r.GitlabConfig.ReadAuthorizerCredential.UserTokenSecretVersion != nil {
				rGitlabConfigReadAuthorizerCredential["userTokenSecretVersion"] = *r.GitlabConfig.ReadAuthorizerCredential.UserTokenSecretVersion
			}
			if r.GitlabConfig.ReadAuthorizerCredential.Username != nil {
				rGitlabConfigReadAuthorizerCredential["username"] = *r.GitlabConfig.ReadAuthorizerCredential.Username
			}
			rGitlabConfig["readAuthorizerCredential"] = rGitlabConfigReadAuthorizerCredential
		}
		if r.GitlabConfig.ServerVersion != nil {
			rGitlabConfig["serverVersion"] = *r.GitlabConfig.ServerVersion
		}
		if r.GitlabConfig.ServiceDirectoryConfig != nil && r.GitlabConfig.ServiceDirectoryConfig != dclService.EmptyConnectionGitlabConfigServiceDirectoryConfig {
			rGitlabConfigServiceDirectoryConfig := make(map[string]interface{})
			if r.GitlabConfig.ServiceDirectoryConfig.Service != nil {
				rGitlabConfigServiceDirectoryConfig["service"] = *r.GitlabConfig.ServiceDirectoryConfig.Service
			}
			rGitlabConfig["serviceDirectoryConfig"] = rGitlabConfigServiceDirectoryConfig
		}
		if r.GitlabConfig.SslCa != nil {
			rGitlabConfig["sslCa"] = *r.GitlabConfig.SslCa
		}
		if r.GitlabConfig.WebhookSecretSecretVersion != nil {
			rGitlabConfig["webhookSecretSecretVersion"] = *r.GitlabConfig.WebhookSecretSecretVersion
		}
		u.Object["gitlabConfig"] = rGitlabConfig
	}
	if r.InstallationState != nil && r.InstallationState != dclService.EmptyConnectionInstallationState {
		rInstallationState := make(map[string]interface{})
		if r.InstallationState.ActionUri != nil {
			rInstallationState["actionUri"] = *r.InstallationState.ActionUri
		}
		if r.InstallationState.Message != nil {
			rInstallationState["message"] = *r.InstallationState.Message
		}
		if r.InstallationState.Stage != nil {
			rInstallationState["stage"] = string(*r.InstallationState.Stage)
		}
		u.Object["installationState"] = rInstallationState
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
	if r.Reconciling != nil {
		u.Object["reconciling"] = *r.Reconciling
	}
	if r.UpdateTime != nil {
		u.Object["updateTime"] = *r.UpdateTime
	}
	return u
}

func UnstructuredToConnection(u *unstructured.Resource) (*dclService.Connection, error) {
	r := &dclService.Connection{}
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
	if _, ok := u.Object["createTime"]; ok {
		if s, ok := u.Object["createTime"].(string); ok {
			r.CreateTime = dcl.String(s)
		} else {
			return nil, fmt.Errorf("r.CreateTime: expected string")
		}
	}
	if _, ok := u.Object["disabled"]; ok {
		if b, ok := u.Object["disabled"].(bool); ok {
			r.Disabled = dcl.Bool(b)
		} else {
			return nil, fmt.Errorf("r.Disabled: expected bool")
		}
	}
	if _, ok := u.Object["etag"]; ok {
		if s, ok := u.Object["etag"].(string); ok {
			r.Etag = dcl.String(s)
		} else {
			return nil, fmt.Errorf("r.Etag: expected string")
		}
	}
	if _, ok := u.Object["githubConfig"]; ok {
		if rGithubConfig, ok := u.Object["githubConfig"].(map[string]interface{}); ok {
			r.GithubConfig = &dclService.ConnectionGithubConfig{}
			if _, ok := rGithubConfig["appInstallationId"]; ok {
				if i, ok := rGithubConfig["appInstallationId"].(int64); ok {
					r.GithubConfig.AppInstallationId = dcl.Int64(i)
				} else {
					return nil, fmt.Errorf("r.GithubConfig.AppInstallationId: expected int64")
				}
			}
			if _, ok := rGithubConfig["authorizerCredential"]; ok {
				if rGithubConfigAuthorizerCredential, ok := rGithubConfig["authorizerCredential"].(map[string]interface{}); ok {
					r.GithubConfig.AuthorizerCredential = &dclService.ConnectionGithubConfigAuthorizerCredential{}
					if _, ok := rGithubConfigAuthorizerCredential["oauthTokenSecretVersion"]; ok {
						if s, ok := rGithubConfigAuthorizerCredential["oauthTokenSecretVersion"].(string); ok {
							r.GithubConfig.AuthorizerCredential.OAuthTokenSecretVersion = dcl.String(s)
						} else {
							return nil, fmt.Errorf("r.GithubConfig.AuthorizerCredential.OAuthTokenSecretVersion: expected string")
						}
					}
					if _, ok := rGithubConfigAuthorizerCredential["username"]; ok {
						if s, ok := rGithubConfigAuthorizerCredential["username"].(string); ok {
							r.GithubConfig.AuthorizerCredential.Username = dcl.String(s)
						} else {
							return nil, fmt.Errorf("r.GithubConfig.AuthorizerCredential.Username: expected string")
						}
					}
				} else {
					return nil, fmt.Errorf("r.GithubConfig.AuthorizerCredential: expected map[string]interface{}")
				}
			}
		} else {
			return nil, fmt.Errorf("r.GithubConfig: expected map[string]interface{}")
		}
	}
	if _, ok := u.Object["githubEnterpriseConfig"]; ok {
		if rGithubEnterpriseConfig, ok := u.Object["githubEnterpriseConfig"].(map[string]interface{}); ok {
			r.GithubEnterpriseConfig = &dclService.ConnectionGithubEnterpriseConfig{}
			if _, ok := rGithubEnterpriseConfig["appId"]; ok {
				if i, ok := rGithubEnterpriseConfig["appId"].(int64); ok {
					r.GithubEnterpriseConfig.AppId = dcl.Int64(i)
				} else {
					return nil, fmt.Errorf("r.GithubEnterpriseConfig.AppId: expected int64")
				}
			}
			if _, ok := rGithubEnterpriseConfig["appInstallationId"]; ok {
				if i, ok := rGithubEnterpriseConfig["appInstallationId"].(int64); ok {
					r.GithubEnterpriseConfig.AppInstallationId = dcl.Int64(i)
				} else {
					return nil, fmt.Errorf("r.GithubEnterpriseConfig.AppInstallationId: expected int64")
				}
			}
			if _, ok := rGithubEnterpriseConfig["appSlug"]; ok {
				if s, ok := rGithubEnterpriseConfig["appSlug"].(string); ok {
					r.GithubEnterpriseConfig.AppSlug = dcl.String(s)
				} else {
					return nil, fmt.Errorf("r.GithubEnterpriseConfig.AppSlug: expected string")
				}
			}
			if _, ok := rGithubEnterpriseConfig["hostUri"]; ok {
				if s, ok := rGithubEnterpriseConfig["hostUri"].(string); ok {
					r.GithubEnterpriseConfig.HostUri = dcl.String(s)
				} else {
					return nil, fmt.Errorf("r.GithubEnterpriseConfig.HostUri: expected string")
				}
			}
			if _, ok := rGithubEnterpriseConfig["privateKeySecretVersion"]; ok {
				if s, ok := rGithubEnterpriseConfig["privateKeySecretVersion"].(string); ok {
					r.GithubEnterpriseConfig.PrivateKeySecretVersion = dcl.String(s)
				} else {
					return nil, fmt.Errorf("r.GithubEnterpriseConfig.PrivateKeySecretVersion: expected string")
				}
			}
			if _, ok := rGithubEnterpriseConfig["serviceDirectoryConfig"]; ok {
				if rGithubEnterpriseConfigServiceDirectoryConfig, ok := rGithubEnterpriseConfig["serviceDirectoryConfig"].(map[string]interface{}); ok {
					r.GithubEnterpriseConfig.ServiceDirectoryConfig = &dclService.ConnectionGithubEnterpriseConfigServiceDirectoryConfig{}
					if _, ok := rGithubEnterpriseConfigServiceDirectoryConfig["service"]; ok {
						if s, ok := rGithubEnterpriseConfigServiceDirectoryConfig["service"].(string); ok {
							r.GithubEnterpriseConfig.ServiceDirectoryConfig.Service = dcl.String(s)
						} else {
							return nil, fmt.Errorf("r.GithubEnterpriseConfig.ServiceDirectoryConfig.Service: expected string")
						}
					}
				} else {
					return nil, fmt.Errorf("r.GithubEnterpriseConfig.ServiceDirectoryConfig: expected map[string]interface{}")
				}
			}
			if _, ok := rGithubEnterpriseConfig["sslCa"]; ok {
				if s, ok := rGithubEnterpriseConfig["sslCa"].(string); ok {
					r.GithubEnterpriseConfig.SslCa = dcl.String(s)
				} else {
					return nil, fmt.Errorf("r.GithubEnterpriseConfig.SslCa: expected string")
				}
			}
			if _, ok := rGithubEnterpriseConfig["webhookSecretSecretVersion"]; ok {
				if s, ok := rGithubEnterpriseConfig["webhookSecretSecretVersion"].(string); ok {
					r.GithubEnterpriseConfig.WebhookSecretSecretVersion = dcl.String(s)
				} else {
					return nil, fmt.Errorf("r.GithubEnterpriseConfig.WebhookSecretSecretVersion: expected string")
				}
			}
		} else {
			return nil, fmt.Errorf("r.GithubEnterpriseConfig: expected map[string]interface{}")
		}
	}
	if _, ok := u.Object["gitlabConfig"]; ok {
		if rGitlabConfig, ok := u.Object["gitlabConfig"].(map[string]interface{}); ok {
			r.GitlabConfig = &dclService.ConnectionGitlabConfig{}
			if _, ok := rGitlabConfig["authorizerCredential"]; ok {
				if rGitlabConfigAuthorizerCredential, ok := rGitlabConfig["authorizerCredential"].(map[string]interface{}); ok {
					r.GitlabConfig.AuthorizerCredential = &dclService.ConnectionGitlabConfigAuthorizerCredential{}
					if _, ok := rGitlabConfigAuthorizerCredential["userTokenSecretVersion"]; ok {
						if s, ok := rGitlabConfigAuthorizerCredential["userTokenSecretVersion"].(string); ok {
							r.GitlabConfig.AuthorizerCredential.UserTokenSecretVersion = dcl.String(s)
						} else {
							return nil, fmt.Errorf("r.GitlabConfig.AuthorizerCredential.UserTokenSecretVersion: expected string")
						}
					}
					if _, ok := rGitlabConfigAuthorizerCredential["username"]; ok {
						if s, ok := rGitlabConfigAuthorizerCredential["username"].(string); ok {
							r.GitlabConfig.AuthorizerCredential.Username = dcl.String(s)
						} else {
							return nil, fmt.Errorf("r.GitlabConfig.AuthorizerCredential.Username: expected string")
						}
					}
				} else {
					return nil, fmt.Errorf("r.GitlabConfig.AuthorizerCredential: expected map[string]interface{}")
				}
			}
			if _, ok := rGitlabConfig["hostUri"]; ok {
				if s, ok := rGitlabConfig["hostUri"].(string); ok {
					r.GitlabConfig.HostUri = dcl.String(s)
				} else {
					return nil, fmt.Errorf("r.GitlabConfig.HostUri: expected string")
				}
			}
			if _, ok := rGitlabConfig["readAuthorizerCredential"]; ok {
				if rGitlabConfigReadAuthorizerCredential, ok := rGitlabConfig["readAuthorizerCredential"].(map[string]interface{}); ok {
					r.GitlabConfig.ReadAuthorizerCredential = &dclService.ConnectionGitlabConfigReadAuthorizerCredential{}
					if _, ok := rGitlabConfigReadAuthorizerCredential["userTokenSecretVersion"]; ok {
						if s, ok := rGitlabConfigReadAuthorizerCredential["userTokenSecretVersion"].(string); ok {
							r.GitlabConfig.ReadAuthorizerCredential.UserTokenSecretVersion = dcl.String(s)
						} else {
							return nil, fmt.Errorf("r.GitlabConfig.ReadAuthorizerCredential.UserTokenSecretVersion: expected string")
						}
					}
					if _, ok := rGitlabConfigReadAuthorizerCredential["username"]; ok {
						if s, ok := rGitlabConfigReadAuthorizerCredential["username"].(string); ok {
							r.GitlabConfig.ReadAuthorizerCredential.Username = dcl.String(s)
						} else {
							return nil, fmt.Errorf("r.GitlabConfig.ReadAuthorizerCredential.Username: expected string")
						}
					}
				} else {
					return nil, fmt.Errorf("r.GitlabConfig.ReadAuthorizerCredential: expected map[string]interface{}")
				}
			}
			if _, ok := rGitlabConfig["serverVersion"]; ok {
				if s, ok := rGitlabConfig["serverVersion"].(string); ok {
					r.GitlabConfig.ServerVersion = dcl.String(s)
				} else {
					return nil, fmt.Errorf("r.GitlabConfig.ServerVersion: expected string")
				}
			}
			if _, ok := rGitlabConfig["serviceDirectoryConfig"]; ok {
				if rGitlabConfigServiceDirectoryConfig, ok := rGitlabConfig["serviceDirectoryConfig"].(map[string]interface{}); ok {
					r.GitlabConfig.ServiceDirectoryConfig = &dclService.ConnectionGitlabConfigServiceDirectoryConfig{}
					if _, ok := rGitlabConfigServiceDirectoryConfig["service"]; ok {
						if s, ok := rGitlabConfigServiceDirectoryConfig["service"].(string); ok {
							r.GitlabConfig.ServiceDirectoryConfig.Service = dcl.String(s)
						} else {
							return nil, fmt.Errorf("r.GitlabConfig.ServiceDirectoryConfig.Service: expected string")
						}
					}
				} else {
					return nil, fmt.Errorf("r.GitlabConfig.ServiceDirectoryConfig: expected map[string]interface{}")
				}
			}
			if _, ok := rGitlabConfig["sslCa"]; ok {
				if s, ok := rGitlabConfig["sslCa"].(string); ok {
					r.GitlabConfig.SslCa = dcl.String(s)
				} else {
					return nil, fmt.Errorf("r.GitlabConfig.SslCa: expected string")
				}
			}
			if _, ok := rGitlabConfig["webhookSecretSecretVersion"]; ok {
				if s, ok := rGitlabConfig["webhookSecretSecretVersion"].(string); ok {
					r.GitlabConfig.WebhookSecretSecretVersion = dcl.String(s)
				} else {
					return nil, fmt.Errorf("r.GitlabConfig.WebhookSecretSecretVersion: expected string")
				}
			}
		} else {
			return nil, fmt.Errorf("r.GitlabConfig: expected map[string]interface{}")
		}
	}
	if _, ok := u.Object["installationState"]; ok {
		if rInstallationState, ok := u.Object["installationState"].(map[string]interface{}); ok {
			r.InstallationState = &dclService.ConnectionInstallationState{}
			if _, ok := rInstallationState["actionUri"]; ok {
				if s, ok := rInstallationState["actionUri"].(string); ok {
					r.InstallationState.ActionUri = dcl.String(s)
				} else {
					return nil, fmt.Errorf("r.InstallationState.ActionUri: expected string")
				}
			}
			if _, ok := rInstallationState["message"]; ok {
				if s, ok := rInstallationState["message"].(string); ok {
					r.InstallationState.Message = dcl.String(s)
				} else {
					return nil, fmt.Errorf("r.InstallationState.Message: expected string")
				}
			}
			if _, ok := rInstallationState["stage"]; ok {
				if s, ok := rInstallationState["stage"].(string); ok {
					r.InstallationState.Stage = dclService.ConnectionInstallationStateStageEnumRef(s)
				} else {
					return nil, fmt.Errorf("r.InstallationState.Stage: expected string")
				}
			}
		} else {
			return nil, fmt.Errorf("r.InstallationState: expected map[string]interface{}")
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
	if _, ok := u.Object["reconciling"]; ok {
		if b, ok := u.Object["reconciling"].(bool); ok {
			r.Reconciling = dcl.Bool(b)
		} else {
			return nil, fmt.Errorf("r.Reconciling: expected bool")
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

func GetConnection(ctx context.Context, config *dcl.Config, u *unstructured.Resource) (*unstructured.Resource, error) {
	c := dclService.NewClient(config)
	r, err := UnstructuredToConnection(u)
	if err != nil {
		return nil, err
	}
	r, err = c.GetConnection(ctx, r)
	if err != nil {
		return nil, err
	}
	return ConnectionToUnstructured(r), nil
}

func ListConnection(ctx context.Context, config *dcl.Config, project string, location string) ([]*unstructured.Resource, error) {
	c := dclService.NewClient(config)
	l, err := c.ListConnection(ctx, project, location)
	if err != nil {
		return nil, err
	}
	var resources []*unstructured.Resource
	for {
		for _, r := range l.Items {
			resources = append(resources, ConnectionToUnstructured(r))
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

func ApplyConnection(ctx context.Context, config *dcl.Config, u *unstructured.Resource, opts ...dcl.ApplyOption) (*unstructured.Resource, error) {
	c := dclService.NewClient(config)
	r, err := UnstructuredToConnection(u)
	if err != nil {
		return nil, err
	}
	if ush := unstructured.FetchStateHint(opts); ush != nil {
		sh, err := UnstructuredToConnection(ush)
		if err != nil {
			return nil, err
		}
		opts = append(opts, dcl.WithStateHint(sh))
	}
	r, err = c.ApplyConnection(ctx, r, opts...)
	if err != nil {
		return nil, err
	}
	return ConnectionToUnstructured(r), nil
}

func ConnectionHasDiff(ctx context.Context, config *dcl.Config, u *unstructured.Resource, opts ...dcl.ApplyOption) (bool, error) {
	c := dclService.NewClient(config)
	r, err := UnstructuredToConnection(u)
	if err != nil {
		return false, err
	}
	if ush := unstructured.FetchStateHint(opts); ush != nil {
		sh, err := UnstructuredToConnection(ush)
		if err != nil {
			return false, err
		}
		opts = append(opts, dcl.WithStateHint(sh))
	}
	opts = append(opts, dcl.WithLifecycleParam(dcl.BlockDestruction), dcl.WithLifecycleParam(dcl.BlockCreation), dcl.WithLifecycleParam(dcl.BlockModification))
	_, err = c.ApplyConnection(ctx, r, opts...)
	if err != nil {
		if _, ok := err.(dcl.ApplyInfeasibleError); ok {
			return true, nil
		}
		return false, err
	}
	return false, nil
}

func DeleteConnection(ctx context.Context, config *dcl.Config, u *unstructured.Resource) error {
	c := dclService.NewClient(config)
	r, err := UnstructuredToConnection(u)
	if err != nil {
		return err
	}
	return c.DeleteConnection(ctx, r)
}

func ConnectionID(u *unstructured.Resource) (string, error) {
	r, err := UnstructuredToConnection(u)
	if err != nil {
		return "", err
	}
	return r.ID()
}

func (r *Connection) STV() unstructured.ServiceTypeVersion {
	return unstructured.ServiceTypeVersion{
		"cloudbuildv2",
		"Connection",
		"ga",
	}
}

func (r *Connection) SetPolicyMember(ctx context.Context, config *dcl.Config, resource *unstructured.Resource, member *unstructured.Resource) (*unstructured.Resource, error) {
	return nil, unstructured.ErrNoSuchMethod
}

func (r *Connection) GetPolicyMember(ctx context.Context, config *dcl.Config, resource *unstructured.Resource, role, member string) (*unstructured.Resource, error) {
	return nil, unstructured.ErrNoSuchMethod
}

func (r *Connection) DeletePolicyMember(ctx context.Context, config *dcl.Config, resource *unstructured.Resource, member *unstructured.Resource) error {
	return unstructured.ErrNoSuchMethod
}

func (r *Connection) SetPolicy(ctx context.Context, config *dcl.Config, resource *unstructured.Resource, policy *unstructured.Resource) (*unstructured.Resource, error) {
	return nil, unstructured.ErrNoSuchMethod
}

func (r *Connection) SetPolicyWithEtag(ctx context.Context, config *dcl.Config, resource *unstructured.Resource, policy *unstructured.Resource) (*unstructured.Resource, error) {
	return nil, unstructured.ErrNoSuchMethod
}

func (r *Connection) GetPolicy(ctx context.Context, config *dcl.Config, resource *unstructured.Resource) (*unstructured.Resource, error) {
	return nil, unstructured.ErrNoSuchMethod
}

func (r *Connection) Get(ctx context.Context, config *dcl.Config, resource *unstructured.Resource) (*unstructured.Resource, error) {
	return GetConnection(ctx, config, resource)
}

func (r *Connection) Apply(ctx context.Context, config *dcl.Config, resource *unstructured.Resource, opts ...dcl.ApplyOption) (*unstructured.Resource, error) {
	return ApplyConnection(ctx, config, resource, opts...)
}

func (r *Connection) HasDiff(ctx context.Context, config *dcl.Config, resource *unstructured.Resource, opts ...dcl.ApplyOption) (bool, error) {
	return ConnectionHasDiff(ctx, config, resource, opts...)
}

func (r *Connection) Delete(ctx context.Context, config *dcl.Config, resource *unstructured.Resource) error {
	return DeleteConnection(ctx, config, resource)
}

func (r *Connection) ID(resource *unstructured.Resource) (string, error) {
	return ConnectionID(resource)
}

func init() {
	unstructured.Register(&Connection{})
}
