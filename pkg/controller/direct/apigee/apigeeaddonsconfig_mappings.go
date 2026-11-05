// Copyright 2026 Google LLC
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//      http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package apigee

import (
	"strconv"

	krm "github.com/GoogleCloudPlatform/k8s-config-connector/apis/apigee/v1alpha1"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct"
	api "google.golang.org/api/apigee/v1"
)

func ApigeeAddonsConfigSpec_FromAPI(mapCtx *direct.MapContext, in *api.GoogleCloudApigeeV1AddonsConfig) *krm.ApigeeAddonsConfigSpec {
	if in == nil {
		return nil
	}
	out := &krm.ApigeeAddonsConfigSpec{}
	out.AddonsConfig = GoogleCloudApigeeV1AddonsConfig_FromAPI(mapCtx, in)
	return out
}

func ApigeeAddonsConfigSpec_ToAPI(mapCtx *direct.MapContext, in *krm.ApigeeAddonsConfigSpec) *api.GoogleCloudApigeeV1AddonsConfig {
	if in == nil {
		return nil
	}
	return GoogleCloudApigeeV1AddonsConfig_ToAPI(mapCtx, in.AddonsConfig)
}

func GoogleCloudApigeeV1AddonsConfig_FromAPI(mapCtx *direct.MapContext, in *api.GoogleCloudApigeeV1AddonsConfig) *krm.GoogleCloudApigeeV1AddonsConfig {
	if in == nil {
		return nil
	}
	out := &krm.GoogleCloudApigeeV1AddonsConfig{}
	out.AdvancedApiOpsConfig = GoogleCloudApigeeV1AdvancedApiOpsConfig_FromAPI(mapCtx, in.AdvancedApiOpsConfig)
	out.ApiSecurityConfig = GoogleCloudApigeeV1ApiSecurityConfig_FromAPI(mapCtx, in.ApiSecurityConfig)
	out.ConnectorsPlatformConfig = GoogleCloudApigeeV1ConnectorsPlatformConfig_FromAPI(mapCtx, in.ConnectorsPlatformConfig)
	out.IntegrationConfig = GoogleCloudApigeeV1IntegrationConfig_FromAPI(mapCtx, in.IntegrationConfig)
	out.MonetizationConfig = GoogleCloudApigeeV1MonetizationConfig_FromAPI(mapCtx, in.MonetizationConfig)
	return out
}

func GoogleCloudApigeeV1AddonsConfig_ToAPI(mapCtx *direct.MapContext, in *krm.GoogleCloudApigeeV1AddonsConfig) *api.GoogleCloudApigeeV1AddonsConfig {
	if in == nil {
		return nil
	}
	out := &api.GoogleCloudApigeeV1AddonsConfig{}
	out.AdvancedApiOpsConfig = GoogleCloudApigeeV1AdvancedApiOpsConfig_ToAPI(mapCtx, in.AdvancedApiOpsConfig)
	out.ApiSecurityConfig = GoogleCloudApigeeV1ApiSecurityConfig_ToAPI(mapCtx, in.ApiSecurityConfig)
	out.ConnectorsPlatformConfig = GoogleCloudApigeeV1ConnectorsPlatformConfig_ToAPI(mapCtx, in.ConnectorsPlatformConfig)
	out.IntegrationConfig = GoogleCloudApigeeV1IntegrationConfig_ToAPI(mapCtx, in.IntegrationConfig)
	out.MonetizationConfig = GoogleCloudApigeeV1MonetizationConfig_ToAPI(mapCtx, in.MonetizationConfig)
	return out
}

func GoogleCloudApigeeV1AdvancedApiOpsConfig_FromAPI(mapCtx *direct.MapContext, in *api.GoogleCloudApigeeV1AdvancedApiOpsConfig) *krm.GoogleCloudApigeeV1AdvancedApiOpsConfig {
	if in == nil {
		return nil
	}
	out := &krm.GoogleCloudApigeeV1AdvancedApiOpsConfig{}
	out.Enabled = direct.LazyPtr(in.Enabled)
	return out
}

func GoogleCloudApigeeV1AdvancedApiOpsConfig_ToAPI(mapCtx *direct.MapContext, in *krm.GoogleCloudApigeeV1AdvancedApiOpsConfig) *api.GoogleCloudApigeeV1AdvancedApiOpsConfig {
	if in == nil {
		return nil
	}
	out := &api.GoogleCloudApigeeV1AdvancedApiOpsConfig{}
	out.Enabled = direct.ValueOf(in.Enabled)
	return out
}

func GoogleCloudApigeeV1ApiSecurityConfig_FromAPI(mapCtx *direct.MapContext, in *api.GoogleCloudApigeeV1ApiSecurityConfig) *krm.GoogleCloudApigeeV1ApiSecurityConfig {
	if in == nil {
		return nil
	}
	out := &krm.GoogleCloudApigeeV1ApiSecurityConfig{}
	out.Enabled = direct.LazyPtr(in.Enabled)
	if in.ExpiresAt != 0 {
		out.ExpiresAt = direct.LazyPtr(strconv.FormatInt(in.ExpiresAt, 10))
	}
	return out
}

func GoogleCloudApigeeV1ApiSecurityConfig_ToAPI(mapCtx *direct.MapContext, in *krm.GoogleCloudApigeeV1ApiSecurityConfig) *api.GoogleCloudApigeeV1ApiSecurityConfig {
	if in == nil {
		return nil
	}
	out := &api.GoogleCloudApigeeV1ApiSecurityConfig{}
	out.Enabled = direct.ValueOf(in.Enabled)
	if in.ExpiresAt != nil {
		val, err := strconv.ParseInt(*in.ExpiresAt, 10, 64)
		if err != nil {
			mapCtx.Errorf("cannot parse expiresAt %q as int64: %v", *in.ExpiresAt, err)
		} else {
			out.ExpiresAt = val
		}
	}
	return out
}

func GoogleCloudApigeeV1ConnectorsPlatformConfig_FromAPI(mapCtx *direct.MapContext, in *api.GoogleCloudApigeeV1ConnectorsPlatformConfig) *krm.GoogleCloudApigeeV1ConnectorsPlatformConfig {
	if in == nil {
		return nil
	}
	out := &krm.GoogleCloudApigeeV1ConnectorsPlatformConfig{}
	out.Enabled = direct.LazyPtr(in.Enabled)
	if in.ExpiresAt != 0 {
		out.ExpiresAt = direct.LazyPtr(strconv.FormatInt(in.ExpiresAt, 10))
	}
	return out
}

func GoogleCloudApigeeV1ConnectorsPlatformConfig_ToAPI(mapCtx *direct.MapContext, in *krm.GoogleCloudApigeeV1ConnectorsPlatformConfig) *api.GoogleCloudApigeeV1ConnectorsPlatformConfig {
	if in == nil {
		return nil
	}
	out := &api.GoogleCloudApigeeV1ConnectorsPlatformConfig{}
	out.Enabled = direct.ValueOf(in.Enabled)
	if in.ExpiresAt != nil {
		val, err := strconv.ParseInt(*in.ExpiresAt, 10, 64)
		if err != nil {
			mapCtx.Errorf("cannot parse expiresAt %q as int64: %v", *in.ExpiresAt, err)
		} else {
			out.ExpiresAt = val
		}
	}
	return out
}

func GoogleCloudApigeeV1IntegrationConfig_FromAPI(mapCtx *direct.MapContext, in *api.GoogleCloudApigeeV1IntegrationConfig) *krm.GoogleCloudApigeeV1IntegrationConfig {
	if in == nil {
		return nil
	}
	out := &krm.GoogleCloudApigeeV1IntegrationConfig{}
	out.Enabled = direct.LazyPtr(in.Enabled)
	return out
}

func GoogleCloudApigeeV1IntegrationConfig_ToAPI(mapCtx *direct.MapContext, in *krm.GoogleCloudApigeeV1IntegrationConfig) *api.GoogleCloudApigeeV1IntegrationConfig {
	if in == nil {
		return nil
	}
	out := &api.GoogleCloudApigeeV1IntegrationConfig{}
	out.Enabled = direct.ValueOf(in.Enabled)
	return out
}

func GoogleCloudApigeeV1MonetizationConfig_FromAPI(mapCtx *direct.MapContext, in *api.GoogleCloudApigeeV1MonetizationConfig) *krm.GoogleCloudApigeeV1MonetizationConfig {
	if in == nil {
		return nil
	}
	out := &krm.GoogleCloudApigeeV1MonetizationConfig{}
	out.Enabled = direct.LazyPtr(in.Enabled)
	return out
}

func GoogleCloudApigeeV1MonetizationConfig_ToAPI(mapCtx *direct.MapContext, in *krm.GoogleCloudApigeeV1MonetizationConfig) *api.GoogleCloudApigeeV1MonetizationConfig {
	if in == nil {
		return nil
	}
	out := &api.GoogleCloudApigeeV1MonetizationConfig{}
	out.Enabled = direct.ValueOf(in.Enabled)
	return out
}
