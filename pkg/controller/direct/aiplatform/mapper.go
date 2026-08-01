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

package aiplatform

import (
	"encoding/json"
	"strconv"

	aiplatformpb "cloud.google.com/go/aiplatform/apiv1beta1/aiplatformpb"
	agentsearchv1alpha1 "github.com/GoogleCloudPlatform/k8s-config-connector/apis/agentsearch/v1alpha1"
	krm "github.com/GoogleCloudPlatform/k8s-config-connector/apis/aiplatform/v1alpha1"
	refsv1beta1 "github.com/GoogleCloudPlatform/k8s-config-connector/apis/refs/v1beta1"
	servicedirectoryv1beta1 "github.com/GoogleCloudPlatform/k8s-config-connector/apis/servicedirectory/v1beta1"
	storagev1beta1 "github.com/GoogleCloudPlatform/k8s-config-connector/apis/storage/v1beta1"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct"
	"google.golang.org/protobuf/types/known/structpb"
	"google.golang.org/protobuf/types/known/wrapperspb"
	apiextensionsv1 "k8s.io/apiextensions-apiserver/pkg/apis/apiextensions/v1"
)

func Int32Value_FromProto(mapCtx *direct.MapContext, in *wrapperspb.Int32Value) *krm.Int32Value {
	if in == nil {
		return nil
	}
	out := in.Value
	return &krm.Int32Value{
		Value: &out,
	}
}

func Int32Value_ToProto(mapCtx *direct.MapContext, in *krm.Int32Value) *wrapperspb.Int32Value {
	if in == nil || in.Value == nil {
		return nil
	}
	return wrapperspb.Int32(*in.Value)
}

func Schema_FromProto(mapCtx *direct.MapContext, in *aiplatformpb.Schema) *krm.Schema {
	if in == nil {
		return nil
	}
	out := &krm.Schema{}
	out.Type = direct.Enum_FromProto(mapCtx, in.GetType())
	out.Format = direct.LazyPtr(in.GetFormat())
	out.Title = direct.LazyPtr(in.GetTitle())
	out.Description = direct.LazyPtr(in.GetDescription())
	out.Nullable = direct.LazyPtr(in.GetNullable())
	out.Default = Value_FromProto(mapCtx, in.GetDefault())
	if in.GetItems() != nil {
		nested := Schema_FromProto(mapCtx, in.GetItems())
		if nested != nil {
			b, err := json.Marshal(nested)
			if err != nil {
				mapCtx.Errorf("error marshalling nested schema to JSON: %v", err)
			} else {
				out.Items = &apiextensionsv1.JSON{Raw: b}
			}
		}
	}
	out.MinItems = direct.LazyPtr(in.GetMinItems())
	out.MaxItems = direct.LazyPtr(in.GetMaxItems())
	out.Enum = in.Enum
	out.PropertyOrdering = in.PropertyOrdering
	out.Required = in.Required
	out.MinProperties = direct.LazyPtr(in.GetMinProperties())
	out.MaxProperties = direct.LazyPtr(in.GetMaxProperties())
	out.Minimum = direct.LazyPtr(in.GetMinimum())
	out.Maximum = direct.LazyPtr(in.GetMaximum())
	out.MinLength = direct.LazyPtr(in.GetMinLength())
	out.MaxLength = direct.LazyPtr(in.GetMaxLength())
	out.Pattern = direct.LazyPtr(in.GetPattern())
	out.Example = Value_FromProto(mapCtx, in.GetExample())
	if in.AnyOf != nil {
		out.AnyOf = make([]apiextensionsv1.JSON, 0, len(in.AnyOf))
		for _, x := range in.AnyOf {
			nested := Schema_FromProto(mapCtx, x)
			if nested != nil {
				b, err := json.Marshal(nested)
				if err != nil {
					mapCtx.Errorf("error marshalling nested anyOf schema to JSON: %v", err)
				} else {
					out.AnyOf = append(out.AnyOf, apiextensionsv1.JSON{Raw: b})
				}
			}
		}
	}
	out.AdditionalProperties = Value_FromProto(mapCtx, in.GetAdditionalProperties())
	out.Ref = direct.LazyPtr(in.GetRef())
	return out
}

func Schema_ToProto(mapCtx *direct.MapContext, in *krm.Schema) *aiplatformpb.Schema {
	if in == nil {
		return nil
	}
	out := &aiplatformpb.Schema{}
	out.Type = direct.Enum_ToProto[aiplatformpb.Type](mapCtx, in.Type)
	out.Format = direct.ValueOf(in.Format)
	out.Title = direct.ValueOf(in.Title)
	out.Description = direct.ValueOf(in.Description)
	out.Nullable = direct.ValueOf(in.Nullable)
	out.Default = Value_ToProto(mapCtx, in.Default)
	if in.Items != nil {
		var nested krm.Schema
		if err := json.Unmarshal(in.Items.Raw, &nested); err != nil {
			mapCtx.Errorf("error unmarshalling nested schema from JSON: %v", err)
		} else {
			out.Items = Schema_ToProto(mapCtx, &nested)
		}
	}
	out.MinItems = direct.ValueOf(in.MinItems)
	out.MaxItems = direct.ValueOf(in.MaxItems)
	out.Enum = in.Enum
	out.PropertyOrdering = in.PropertyOrdering
	out.Required = in.Required
	out.MinProperties = direct.ValueOf(in.MinProperties)
	out.MaxProperties = direct.ValueOf(in.MaxProperties)
	out.Minimum = direct.ValueOf(in.Minimum)
	out.Maximum = direct.ValueOf(in.Maximum)
	out.MinLength = direct.ValueOf(in.MinLength)
	out.MaxLength = direct.ValueOf(in.MaxLength)
	out.Pattern = direct.ValueOf(in.Pattern)
	out.Example = Value_ToProto(mapCtx, in.Example)
	if len(in.AnyOf) > 0 {
		out.AnyOf = make([]*aiplatformpb.Schema, 0, len(in.AnyOf))
		for _, x := range in.AnyOf {
			var nested krm.Schema
			if err := json.Unmarshal(x.Raw, &nested); err != nil {
				mapCtx.Errorf("error unmarshalling anyOf schema from JSON: %v", err)
			} else {
				out.AnyOf = append(out.AnyOf, Schema_ToProto(mapCtx, &nested))
			}
		}
	}
	out.AdditionalProperties = Value_ToProto(mapCtx, in.AdditionalProperties)
	out.Ref = direct.ValueOf(in.Ref)
	return out
}

func AuthConfig_OIDCConfig_FromProto(mapCtx *direct.MapContext, in *aiplatformpb.AuthConfig_OidcConfig) *krm.AuthConfig_OIDCConfig {
	if in == nil {
		return nil
	}
	out := &krm.AuthConfig_OIDCConfig{}
	switch oneof := in.OidcConfig.(type) {
	case *aiplatformpb.AuthConfig_OidcConfig_IdToken:
		if oneof != nil {
			out.IDToken = &oneof.IdToken
		}
	case *aiplatformpb.AuthConfig_OidcConfig_ServiceAccount:
		if oneof != nil && oneof.ServiceAccount != "" {
			out.ServiceAccountRef = &refsv1beta1.IAMServiceAccountRef{External: oneof.ServiceAccount}
		}
	}
	return out
}

func AuthConfig_OIDCConfig_ToProto(mapCtx *direct.MapContext, in *krm.AuthConfig_OIDCConfig) *aiplatformpb.AuthConfig_OidcConfig {
	if in == nil {
		return nil
	}
	out := &aiplatformpb.AuthConfig_OidcConfig{}
	if in.IDToken != nil {
		out.OidcConfig = &aiplatformpb.AuthConfig_OidcConfig_IdToken{IdToken: *in.IDToken}
	} else if in.ServiceAccountRef != nil {
		out.OidcConfig = &aiplatformpb.AuthConfig_OidcConfig_ServiceAccount{ServiceAccount: in.ServiceAccountRef.External}
	}
	return out
}

func AuthConfig_OauthConfig_FromProto(mapCtx *direct.MapContext, in *aiplatformpb.AuthConfig_OauthConfig) *krm.AuthConfig_OauthConfig {
	if in == nil {
		return nil
	}
	out := &krm.AuthConfig_OauthConfig{}
	switch oneof := in.OauthConfig.(type) {
	case *aiplatformpb.AuthConfig_OauthConfig_AccessToken:
		if oneof != nil {
			out.AccessToken = &oneof.AccessToken
		}
	case *aiplatformpb.AuthConfig_OauthConfig_ServiceAccount:
		if oneof != nil && oneof.ServiceAccount != "" {
			out.ServiceAccountRef = &refsv1beta1.IAMServiceAccountRef{External: oneof.ServiceAccount}
		}
	}
	return out
}

func AuthConfig_OauthConfig_ToProto(mapCtx *direct.MapContext, in *krm.AuthConfig_OauthConfig) *aiplatformpb.AuthConfig_OauthConfig {
	if in == nil {
		return nil
	}
	out := &aiplatformpb.AuthConfig_OauthConfig{}
	if in.AccessToken != nil {
		out.OauthConfig = &aiplatformpb.AuthConfig_OauthConfig_AccessToken{AccessToken: *in.AccessToken}
	} else if in.ServiceAccountRef != nil {
		out.OauthConfig = &aiplatformpb.AuthConfig_OauthConfig_ServiceAccount{ServiceAccount: in.ServiceAccountRef.External}
	}
	return out
}

func AuthConfig_APIKeyConfig_FromProto(mapCtx *direct.MapContext, in *aiplatformpb.AuthConfig_ApiKeyConfig) *krm.AuthConfig_APIKeyConfig {
	if in == nil {
		return nil
	}
	out := &krm.AuthConfig_APIKeyConfig{}
	out.Name = direct.LazyPtr(in.GetName())
	if in.GetApiKeySecret() != "" {
		out.APIKeySecretRef = &refsv1beta1.SecretManagerSecretVersionRef{External: in.GetApiKeySecret()}
	}
	out.HTTPElementLocation = direct.Enum_FromProto(mapCtx, in.GetHttpElementLocation())
	return out
}

func AuthConfig_APIKeyConfig_ToProto(mapCtx *direct.MapContext, in *krm.AuthConfig_APIKeyConfig) *aiplatformpb.AuthConfig_ApiKeyConfig {
	if in == nil {
		return nil
	}
	out := &aiplatformpb.AuthConfig_ApiKeyConfig{}
	out.Name = direct.ValueOf(in.Name)
	if in.APIKeySecretRef != nil {
		out.ApiKeySecret = in.APIKeySecretRef.External
	}
	out.HttpElementLocation = direct.Enum_ToProto[aiplatformpb.HttpElementLocation](mapCtx, in.HTTPElementLocation)
	return out
}

func AuthConfig_GoogleServiceAccountConfig_FromProto(mapCtx *direct.MapContext, in *aiplatformpb.AuthConfig_GoogleServiceAccountConfig) *krm.AuthConfig_GoogleServiceAccountConfig {
	if in == nil {
		return nil
	}
	out := &krm.AuthConfig_GoogleServiceAccountConfig{}
	if in.GetServiceAccount() != "" {
		out.ServiceAccountRef = &refsv1beta1.IAMServiceAccountRef{External: in.GetServiceAccount()}
	}
	return out
}

func AuthConfig_GoogleServiceAccountConfig_ToProto(mapCtx *direct.MapContext, in *krm.AuthConfig_GoogleServiceAccountConfig) *aiplatformpb.AuthConfig_GoogleServiceAccountConfig {
	if in == nil {
		return nil
	}
	out := &aiplatformpb.AuthConfig_GoogleServiceAccountConfig{}
	if in.ServiceAccountRef != nil {
		out.ServiceAccount = in.ServiceAccountRef.External
	}
	return out
}

func AuthConfig_HTTPBasicAuthConfig_FromProto(mapCtx *direct.MapContext, in *aiplatformpb.AuthConfig_HttpBasicAuthConfig) *krm.AuthConfig_HTTPBasicAuthConfig {
	if in == nil {
		return nil
	}
	out := &krm.AuthConfig_HTTPBasicAuthConfig{}
	if in.GetCredentialSecret() != "" {
		out.CredentialSecretRef = &refsv1beta1.SecretManagerSecretVersionRef{External: in.GetCredentialSecret()}
	}
	return out
}

func AuthConfig_HTTPBasicAuthConfig_ToProto(mapCtx *direct.MapContext, in *krm.AuthConfig_HTTPBasicAuthConfig) *aiplatformpb.AuthConfig_HttpBasicAuthConfig {
	if in == nil {
		return nil
	}
	out := &aiplatformpb.AuthConfig_HttpBasicAuthConfig{}
	if in.CredentialSecretRef != nil {
		out.CredentialSecret = in.CredentialSecretRef.External
	}
	return out
}

func ExtensionPrivateServiceConnectConfig_FromProto(mapCtx *direct.MapContext, in *aiplatformpb.ExtensionPrivateServiceConnectConfig) *krm.ExtensionPrivateServiceConnectConfig {
	if in == nil {
		return nil
	}
	out := &krm.ExtensionPrivateServiceConnectConfig{}
	if in.GetServiceDirectory() != "" {
		out.ServiceDirectoryRef = &servicedirectoryv1beta1.ServiceDirectoryServiceRef{External: in.GetServiceDirectory()}
	}
	return out
}

func ExtensionPrivateServiceConnectConfig_ToProto(mapCtx *direct.MapContext, in *krm.ExtensionPrivateServiceConnectConfig) *aiplatformpb.ExtensionPrivateServiceConnectConfig {
	if in == nil {
		return nil
	}
	out := &aiplatformpb.ExtensionPrivateServiceConnectConfig{}
	if in.ServiceDirectoryRef != nil {
		out.ServiceDirectory = in.ServiceDirectoryRef.External
	}
	return out
}

func RuntimeConfig_CodeInterpreterRuntimeConfig_FromProto(mapCtx *direct.MapContext, in *aiplatformpb.RuntimeConfig_CodeInterpreterRuntimeConfig) *krm.RuntimeConfig_CodeInterpreterRuntimeConfig {
	if in == nil {
		return nil
	}
	out := &krm.RuntimeConfig_CodeInterpreterRuntimeConfig{}
	if in.GetFileInputGcsBucket() != "" {
		out.FileInputGCSBucketRef = &storagev1beta1.StorageBucketRef{External: in.GetFileInputGcsBucket()}
	}
	if in.GetFileOutputGcsBucket() != "" {
		out.FileOutputGCSBucketRef = &storagev1beta1.StorageBucketRef{External: in.GetFileOutputGcsBucket()}
	}
	return out
}

func RuntimeConfig_CodeInterpreterRuntimeConfig_ToProto(mapCtx *direct.MapContext, in *krm.RuntimeConfig_CodeInterpreterRuntimeConfig) *aiplatformpb.RuntimeConfig_CodeInterpreterRuntimeConfig {
	if in == nil {
		return nil
	}
	out := &aiplatformpb.RuntimeConfig_CodeInterpreterRuntimeConfig{}
	if in.FileInputGCSBucketRef != nil {
		out.FileInputGcsBucket = in.FileInputGCSBucketRef.External
	}
	if in.FileOutputGCSBucketRef != nil {
		out.FileOutputGcsBucket = in.FileOutputGCSBucketRef.External
	}
	return out
}

func RuntimeConfig_VertexAISearchRuntimeConfig_FromProto(mapCtx *direct.MapContext, in *aiplatformpb.RuntimeConfig_VertexAISearchRuntimeConfig) *krm.RuntimeConfig_VertexAISearchRuntimeConfig {
	if in == nil {
		return nil
	}
	out := &krm.RuntimeConfig_VertexAISearchRuntimeConfig{}
	if in.GetServingConfigName() != "" {
		out.ServingConfigRef = &agentsearchv1alpha1.AgentSearchServingConfigRef{External: in.GetServingConfigName()}
	}
	if in.GetEngineId() != "" {
		out.EngineRef = &agentsearchv1alpha1.AgentSearchEngineRef{External: in.GetEngineId()}
	}
	return out
}

func RuntimeConfig_VertexAISearchRuntimeConfig_ToProto(mapCtx *direct.MapContext, in *krm.RuntimeConfig_VertexAISearchRuntimeConfig) *aiplatformpb.RuntimeConfig_VertexAISearchRuntimeConfig {
	if in == nil {
		return nil
	}
	out := &aiplatformpb.RuntimeConfig_VertexAISearchRuntimeConfig{}
	if in.ServingConfigRef != nil {
		out.ServingConfigName = in.ServingConfigRef.External
	}
	if in.EngineRef != nil {
		out.EngineId = in.EngineRef.External
	}
	return out
}

func RuntimeConfig_FromProto(mapCtx *direct.MapContext, in *aiplatformpb.RuntimeConfig) *krm.RuntimeConfig {
	if in == nil {
		return nil
	}
	out := &krm.RuntimeConfig{}
	out.CodeInterpreterRuntimeConfig = RuntimeConfig_CodeInterpreterRuntimeConfig_FromProto(mapCtx, in.GetCodeInterpreterRuntimeConfig())
	out.VertexAISearchRuntimeConfig = RuntimeConfig_VertexAISearchRuntimeConfig_FromProto(mapCtx, in.GetVertexAiSearchRuntimeConfig())
	if v := direct.Struct_FromProto(mapCtx, in.GetDefaultParams()); v != nil {
		out.DefaultParams = *v
	}
	return out
}

func RuntimeConfig_ToProto(mapCtx *direct.MapContext, in *krm.RuntimeConfig) *aiplatformpb.RuntimeConfig {
	if in == nil {
		return nil
	}
	out := &aiplatformpb.RuntimeConfig{}
	if oneof := RuntimeConfig_CodeInterpreterRuntimeConfig_ToProto(mapCtx, in.CodeInterpreterRuntimeConfig); oneof != nil {
		out.GoogleFirstPartyExtensionConfig = &aiplatformpb.RuntimeConfig_CodeInterpreterRuntimeConfig_{CodeInterpreterRuntimeConfig: oneof}
	}
	if oneof := RuntimeConfig_VertexAISearchRuntimeConfig_ToProto(mapCtx, in.VertexAISearchRuntimeConfig); oneof != nil {
		out.GoogleFirstPartyExtensionConfig = &aiplatformpb.RuntimeConfig_VertexAiSearchRuntimeConfig{VertexAiSearchRuntimeConfig: oneof}
	}
	out.DefaultParams = direct.Struct_ToProto(mapCtx, &in.DefaultParams)
	return out
}

func Value_ToProto(mapCtx *direct.MapContext, in *krm.Value) *structpb.Value {
	if in == nil {
		return nil
	}
	out := &structpb.Value{}
	if in.BoolValue != nil {
		out.Kind = &structpb.Value_BoolValue{
			BoolValue: direct.ValueOf(in.BoolValue),
		}
	}
	if in.NullValue != nil {
		strVal := direct.ValueOf(in.NullValue)
		var value int
		if val, ok := structpb.NullValue_value[strVal]; ok {
			value = int(val)
		} else {
			var err error
			value, err = strconv.Atoi(strVal)
			if err != nil {
				mapCtx.Errorf("error converting value %s from string to int", strVal)
			}
		}
		out.Kind = &structpb.Value_NullValue{
			NullValue: structpb.NullValue(value),
		}
	}
	if in.NumberValue != nil {
		out.Kind = &structpb.Value_NumberValue{
			NumberValue: direct.ValueOf(in.NumberValue),
		}
	}
	if in.StringValue != nil {
		out.Kind = &structpb.Value_StringValue{
			StringValue: direct.ValueOf(in.StringValue),
		}
	}
	if len(in.StructValue.Raw) > 0 {
		out.Kind = &structpb.Value_StructValue{
			StructValue: direct.Struct_ToProto(mapCtx, &in.StructValue),
		}
	}
	return out
}

func Value_FromProto(mapCtx *direct.MapContext, in *structpb.Value) *krm.Value {
	if in == nil {
		return nil
	}
	out := &krm.Value{}
	switch in.GetKind().(type) {
	case *structpb.Value_StringValue:
		value := in.GetStringValue()
		out.StringValue = &value
	case *structpb.Value_NumberValue:
		value := in.GetNumberValue()
		out.NumberValue = &value
	case *structpb.Value_NullValue:
		value := in.GetNullValue().String()
		out.NullValue = &value
	case *structpb.Value_BoolValue:
		value := in.GetBoolValue()
		out.BoolValue = &value
	case *structpb.Value_StructValue:
		if val := direct.Struct_FromProto(mapCtx, in.GetStructValue()); val != nil {
			out.StructValue = *val
		}
	}
	return out
}
