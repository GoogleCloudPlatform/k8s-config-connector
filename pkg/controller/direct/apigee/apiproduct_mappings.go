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
	krm "github.com/GoogleCloudPlatform/k8s-config-connector/apis/apigee/v1alpha1"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct"
	api "google.golang.org/api/apigee/v1"
)

func GoogleCloudApigeeV1Attribute_FromApi(mapCtx *direct.MapContext, in *api.GoogleCloudApigeeV1Attribute) *krm.GoogleCloudApigeeV1Attribute {
	if in == nil {
		return nil
	}
	out := &krm.GoogleCloudApigeeV1Attribute{}
	out.Name = direct.LazyPtr(in.Name)
	out.Value = direct.LazyPtr(in.Value)
	return out
}

func GoogleCloudApigeeV1Attribute_ToApi(mapCtx *direct.MapContext, in *krm.GoogleCloudApigeeV1Attribute) *api.GoogleCloudApigeeV1Attribute {
	if in == nil {
		return nil
	}
	out := &api.GoogleCloudApigeeV1Attribute{}
	out.Name = direct.ValueOf(in.Name)
	out.Value = direct.ValueOf(in.Value)
	return out
}

func GoogleCloudApigeeV1Attributes_FromApi(mapCtx *direct.MapContext, in []*api.GoogleCloudApigeeV1Attribute) []krm.GoogleCloudApigeeV1Attribute {
	if in == nil {
		return nil
	}
	out := make([]krm.GoogleCloudApigeeV1Attribute, len(in))
	for i, item := range in {
		if item != nil {
			out[i] = *GoogleCloudApigeeV1Attribute_FromApi(mapCtx, item)
		}
	}
	return out
}

func GoogleCloudApigeeV1Attributes_ToApi(mapCtx *direct.MapContext, in []krm.GoogleCloudApigeeV1Attribute) []*api.GoogleCloudApigeeV1Attribute {
	if in == nil {
		return nil
	}
	out := make([]*api.GoogleCloudApigeeV1Attribute, len(in))
	for i := range in {
		out[i] = GoogleCloudApigeeV1Attribute_ToApi(mapCtx, &in[i])
	}
	return out
}

func GoogleCloudApigeeV1GraphQlOperation_FromApi(mapCtx *direct.MapContext, in *api.GoogleCloudApigeeV1GraphQLOperation) *krm.GoogleCloudApigeeV1GraphQlOperation {
	if in == nil {
		return nil
	}
	out := &krm.GoogleCloudApigeeV1GraphQlOperation{}
	out.Operation = direct.LazyPtr(in.Operation)
	out.OperationTypes = in.OperationTypes
	return out
}

func GoogleCloudApigeeV1GraphQlOperation_ToApi(mapCtx *direct.MapContext, in *krm.GoogleCloudApigeeV1GraphQlOperation) *api.GoogleCloudApigeeV1GraphQLOperation {
	if in == nil {
		return nil
	}
	out := &api.GoogleCloudApigeeV1GraphQLOperation{}
	out.Operation = direct.ValueOf(in.Operation)
	out.OperationTypes = in.OperationTypes
	return out
}

func GoogleCloudApigeeV1GraphQlOperations_FromApi(mapCtx *direct.MapContext, in []*api.GoogleCloudApigeeV1GraphQLOperation) []krm.GoogleCloudApigeeV1GraphQlOperation {
	if in == nil {
		return nil
	}
	out := make([]krm.GoogleCloudApigeeV1GraphQlOperation, len(in))
	for i, item := range in {
		if item != nil {
			out[i] = *GoogleCloudApigeeV1GraphQlOperation_FromApi(mapCtx, item)
		}
	}
	return out
}

func GoogleCloudApigeeV1GraphQlOperations_ToApi(mapCtx *direct.MapContext, in []krm.GoogleCloudApigeeV1GraphQlOperation) []*api.GoogleCloudApigeeV1GraphQLOperation {
	if in == nil {
		return nil
	}
	out := make([]*api.GoogleCloudApigeeV1GraphQLOperation, len(in))
	for i := range in {
		out[i] = GoogleCloudApigeeV1GraphQlOperation_ToApi(mapCtx, &in[i])
	}
	return out
}

func GoogleCloudApigeeV1GraphQlOperationConfig_FromApi(mapCtx *direct.MapContext, in *api.GoogleCloudApigeeV1GraphQLOperationConfig) *krm.GoogleCloudApigeeV1GraphQlOperationConfig {
	if in == nil {
		return nil
	}
	out := &krm.GoogleCloudApigeeV1GraphQlOperationConfig{}
	out.APISource = direct.LazyPtr(in.ApiSource)
	out.Attributes = GoogleCloudApigeeV1Attributes_FromApi(mapCtx, in.Attributes)
	out.Operations = GoogleCloudApigeeV1GraphQlOperations_FromApi(mapCtx, in.Operations)
	out.Quota = GoogleCloudApigeeV1Quota_FromApi(mapCtx, in.Quota)
	return out
}

func GoogleCloudApigeeV1GraphQlOperationConfig_ToApi(mapCtx *direct.MapContext, in *krm.GoogleCloudApigeeV1GraphQlOperationConfig) *api.GoogleCloudApigeeV1GraphQLOperationConfig {
	if in == nil {
		return nil
	}
	out := &api.GoogleCloudApigeeV1GraphQLOperationConfig{}
	out.ApiSource = direct.ValueOf(in.APISource)
	out.Attributes = GoogleCloudApigeeV1Attributes_ToApi(mapCtx, in.Attributes)
	out.Operations = GoogleCloudApigeeV1GraphQlOperations_ToApi(mapCtx, in.Operations)
	out.Quota = GoogleCloudApigeeV1Quota_ToApi(mapCtx, in.Quota)
	return out
}

func GoogleCloudApigeeV1GraphQlOperationConfigs_FromApi(mapCtx *direct.MapContext, in []*api.GoogleCloudApigeeV1GraphQLOperationConfig) []krm.GoogleCloudApigeeV1GraphQlOperationConfig {
	if in == nil {
		return nil
	}
	out := make([]krm.GoogleCloudApigeeV1GraphQlOperationConfig, len(in))
	for i, item := range in {
		if item != nil {
			out[i] = *GoogleCloudApigeeV1GraphQlOperationConfig_FromApi(mapCtx, item)
		}
	}
	return out
}

func GoogleCloudApigeeV1GraphQlOperationConfigs_ToApi(mapCtx *direct.MapContext, in []krm.GoogleCloudApigeeV1GraphQlOperationConfig) []*api.GoogleCloudApigeeV1GraphQLOperationConfig {
	if in == nil {
		return nil
	}
	out := make([]*api.GoogleCloudApigeeV1GraphQLOperationConfig, len(in))
	for i := range in {
		out[i] = GoogleCloudApigeeV1GraphQlOperationConfig_ToApi(mapCtx, &in[i])
	}
	return out
}

func GoogleCloudApigeeV1GraphQlOperationGroup_FromApi(mapCtx *direct.MapContext, in *api.GoogleCloudApigeeV1GraphQLOperationGroup) *krm.GoogleCloudApigeeV1GraphQlOperationGroup {
	if in == nil {
		return nil
	}
	out := &krm.GoogleCloudApigeeV1GraphQlOperationGroup{}
	out.OperationConfigType = direct.LazyPtr(in.OperationConfigType)
	out.OperationConfigs = GoogleCloudApigeeV1GraphQlOperationConfigs_FromApi(mapCtx, in.OperationConfigs)
	return out
}

func GoogleCloudApigeeV1GraphQlOperationGroup_ToApi(mapCtx *direct.MapContext, in *krm.GoogleCloudApigeeV1GraphQlOperationGroup) *api.GoogleCloudApigeeV1GraphQLOperationGroup {
	if in == nil {
		return nil
	}
	out := &api.GoogleCloudApigeeV1GraphQLOperationGroup{}
	out.OperationConfigType = direct.ValueOf(in.OperationConfigType)
	out.OperationConfigs = GoogleCloudApigeeV1GraphQlOperationConfigs_ToApi(mapCtx, in.OperationConfigs)
	return out
}

func GoogleCloudApigeeV1GrpcOperationConfig_FromApi(mapCtx *direct.MapContext, in *api.GoogleCloudApigeeV1GrpcOperationConfig) *krm.GoogleCloudApigeeV1GrpcOperationConfig {
	if in == nil {
		return nil
	}
	out := &krm.GoogleCloudApigeeV1GrpcOperationConfig{}
	out.APISource = direct.LazyPtr(in.ApiSource)
	out.Attributes = GoogleCloudApigeeV1Attributes_FromApi(mapCtx, in.Attributes)
	out.Methods = in.Methods
	out.Quota = GoogleCloudApigeeV1Quota_FromApi(mapCtx, in.Quota)
	out.Service = direct.LazyPtr(in.Service)
	return out
}

func GoogleCloudApigeeV1GrpcOperationConfig_ToApi(mapCtx *direct.MapContext, in *krm.GoogleCloudApigeeV1GrpcOperationConfig) *api.GoogleCloudApigeeV1GrpcOperationConfig {
	if in == nil {
		return nil
	}
	out := &api.GoogleCloudApigeeV1GrpcOperationConfig{}
	out.ApiSource = direct.ValueOf(in.APISource)
	out.Attributes = GoogleCloudApigeeV1Attributes_ToApi(mapCtx, in.Attributes)
	out.Methods = in.Methods
	out.Quota = GoogleCloudApigeeV1Quota_ToApi(mapCtx, in.Quota)
	out.Service = direct.ValueOf(in.Service)
	return out
}

func GoogleCloudApigeeV1GrpcOperationConfigs_FromApi(mapCtx *direct.MapContext, in []*api.GoogleCloudApigeeV1GrpcOperationConfig) []krm.GoogleCloudApigeeV1GrpcOperationConfig {
	if in == nil {
		return nil
	}
	out := make([]krm.GoogleCloudApigeeV1GrpcOperationConfig, len(in))
	for i, item := range in {
		if item != nil {
			out[i] = *GoogleCloudApigeeV1GrpcOperationConfig_FromApi(mapCtx, item)
		}
	}
	return out
}

func GoogleCloudApigeeV1GrpcOperationConfigs_ToApi(mapCtx *direct.MapContext, in []krm.GoogleCloudApigeeV1GrpcOperationConfig) []*api.GoogleCloudApigeeV1GrpcOperationConfig {
	if in == nil {
		return nil
	}
	out := make([]*api.GoogleCloudApigeeV1GrpcOperationConfig, len(in))
	for i := range in {
		out[i] = GoogleCloudApigeeV1GrpcOperationConfig_ToApi(mapCtx, &in[i])
	}
	return out
}

func GoogleCloudApigeeV1GrpcOperationGroup_FromApi(mapCtx *direct.MapContext, in *api.GoogleCloudApigeeV1GrpcOperationGroup) *krm.GoogleCloudApigeeV1GrpcOperationGroup {
	if in == nil {
		return nil
	}
	out := &krm.GoogleCloudApigeeV1GrpcOperationGroup{}
	out.OperationConfigs = GoogleCloudApigeeV1GrpcOperationConfigs_FromApi(mapCtx, in.OperationConfigs)
	return out
}

func GoogleCloudApigeeV1GrpcOperationGroup_ToApi(mapCtx *direct.MapContext, in *krm.GoogleCloudApigeeV1GrpcOperationGroup) *api.GoogleCloudApigeeV1GrpcOperationGroup {
	if in == nil {
		return nil
	}
	out := &api.GoogleCloudApigeeV1GrpcOperationGroup{}
	out.OperationConfigs = GoogleCloudApigeeV1GrpcOperationConfigs_ToApi(mapCtx, in.OperationConfigs)
	return out
}

func GoogleCloudApigeeV1LlmOperation_FromApi(mapCtx *direct.MapContext, in *api.GoogleCloudApigeeV1LlmOperation) *krm.GoogleCloudApigeeV1LlmOperation {
	if in == nil {
		return nil
	}
	out := &krm.GoogleCloudApigeeV1LlmOperation{}
	out.Methods = in.Methods
	out.Model = direct.LazyPtr(in.Model)
	out.Resource = direct.LazyPtr(in.Resource)
	return out
}

func GoogleCloudApigeeV1LlmOperation_ToApi(mapCtx *direct.MapContext, in *krm.GoogleCloudApigeeV1LlmOperation) *api.GoogleCloudApigeeV1LlmOperation {
	if in == nil {
		return nil
	}
	out := &api.GoogleCloudApigeeV1LlmOperation{}
	out.Methods = in.Methods
	out.Model = direct.ValueOf(in.Model)
	out.Resource = direct.ValueOf(in.Resource)
	return out
}

func GoogleCloudApigeeV1LlmOperations_FromApi(mapCtx *direct.MapContext, in []*api.GoogleCloudApigeeV1LlmOperation) []krm.GoogleCloudApigeeV1LlmOperation {
	if in == nil {
		return nil
	}
	out := make([]krm.GoogleCloudApigeeV1LlmOperation, len(in))
	for i, item := range in {
		if item != nil {
			out[i] = *GoogleCloudApigeeV1LlmOperation_FromApi(mapCtx, item)
		}
	}
	return out
}

func GoogleCloudApigeeV1LlmOperations_ToApi(mapCtx *direct.MapContext, in []krm.GoogleCloudApigeeV1LlmOperation) []*api.GoogleCloudApigeeV1LlmOperation {
	if in == nil {
		return nil
	}
	out := make([]*api.GoogleCloudApigeeV1LlmOperation, len(in))
	for i := range in {
		out[i] = GoogleCloudApigeeV1LlmOperation_ToApi(mapCtx, &in[i])
	}
	return out
}

func GoogleCloudApigeeV1LlmOperationConfig_FromApi(mapCtx *direct.MapContext, in *api.GoogleCloudApigeeV1LlmOperationConfig) *krm.GoogleCloudApigeeV1LlmOperationConfig {
	if in == nil {
		return nil
	}
	out := &krm.GoogleCloudApigeeV1LlmOperationConfig{}
	out.APISource = direct.LazyPtr(in.ApiSource)
	out.Attributes = GoogleCloudApigeeV1Attributes_FromApi(mapCtx, in.Attributes)
	out.LlmOperations = GoogleCloudApigeeV1LlmOperations_FromApi(mapCtx, in.LlmOperations)
	out.LlmTokenQuota = GoogleCloudApigeeV1LlmTokenQuota_FromApi(mapCtx, in.LlmTokenQuota)
	return out
}

func GoogleCloudApigeeV1LlmOperationConfig_ToApi(mapCtx *direct.MapContext, in *krm.GoogleCloudApigeeV1LlmOperationConfig) *api.GoogleCloudApigeeV1LlmOperationConfig {
	if in == nil {
		return nil
	}
	out := &api.GoogleCloudApigeeV1LlmOperationConfig{}
	out.ApiSource = direct.ValueOf(in.APISource)
	out.Attributes = GoogleCloudApigeeV1Attributes_ToApi(mapCtx, in.Attributes)
	out.LlmOperations = GoogleCloudApigeeV1LlmOperations_ToApi(mapCtx, in.LlmOperations)
	out.LlmTokenQuota = GoogleCloudApigeeV1LlmTokenQuota_ToApi(mapCtx, in.LlmTokenQuota)
	return out
}

func GoogleCloudApigeeV1LlmOperationConfigs_FromApi(mapCtx *direct.MapContext, in []*api.GoogleCloudApigeeV1LlmOperationConfig) []krm.GoogleCloudApigeeV1LlmOperationConfig {
	if in == nil {
		return nil
	}
	out := make([]krm.GoogleCloudApigeeV1LlmOperationConfig, len(in))
	for i, item := range in {
		if item != nil {
			out[i] = *GoogleCloudApigeeV1LlmOperationConfig_FromApi(mapCtx, item)
		}
	}
	return out
}

func GoogleCloudApigeeV1LlmOperationConfigs_ToApi(mapCtx *direct.MapContext, in []krm.GoogleCloudApigeeV1LlmOperationConfig) []*api.GoogleCloudApigeeV1LlmOperationConfig {
	if in == nil {
		return nil
	}
	out := make([]*api.GoogleCloudApigeeV1LlmOperationConfig, len(in))
	for i := range in {
		out[i] = GoogleCloudApigeeV1LlmOperationConfig_ToApi(mapCtx, &in[i])
	}
	return out
}

func GoogleCloudApigeeV1LlmOperationGroup_FromApi(mapCtx *direct.MapContext, in *api.GoogleCloudApigeeV1LlmOperationGroup) *krm.GoogleCloudApigeeV1LlmOperationGroup {
	if in == nil {
		return nil
	}
	out := &krm.GoogleCloudApigeeV1LlmOperationGroup{}
	out.OperationConfigs = GoogleCloudApigeeV1LlmOperationConfigs_FromApi(mapCtx, in.OperationConfigs)
	return out
}

func GoogleCloudApigeeV1LlmOperationGroup_ToApi(mapCtx *direct.MapContext, in *krm.GoogleCloudApigeeV1LlmOperationGroup) *api.GoogleCloudApigeeV1LlmOperationGroup {
	if in == nil {
		return nil
	}
	out := &api.GoogleCloudApigeeV1LlmOperationGroup{}
	out.OperationConfigs = GoogleCloudApigeeV1LlmOperationConfigs_ToApi(mapCtx, in.OperationConfigs)
	return out
}

func GoogleCloudApigeeV1LlmTokenQuota_FromApi(mapCtx *direct.MapContext, in *api.GoogleCloudApigeeV1LlmTokenQuota) *krm.GoogleCloudApigeeV1LlmTokenQuota {
	if in == nil {
		return nil
	}
	out := &krm.GoogleCloudApigeeV1LlmTokenQuota{}
	out.Interval = direct.LazyPtr(in.Interval)
	out.Limit = direct.LazyPtr(in.Limit)
	out.TimeUnit = direct.LazyPtr(in.TimeUnit)
	return out
}

func GoogleCloudApigeeV1LlmTokenQuota_ToApi(mapCtx *direct.MapContext, in *krm.GoogleCloudApigeeV1LlmTokenQuota) *api.GoogleCloudApigeeV1LlmTokenQuota {
	if in == nil {
		return nil
	}
	out := &api.GoogleCloudApigeeV1LlmTokenQuota{}
	out.Interval = direct.ValueOf(in.Interval)
	out.Limit = direct.ValueOf(in.Limit)
	out.TimeUnit = direct.ValueOf(in.TimeUnit)
	return out
}

func GoogleCloudApigeeV1Operation_FromApi(mapCtx *direct.MapContext, in *api.GoogleCloudApigeeV1Operation) *krm.GoogleCloudApigeeV1Operation {
	if in == nil {
		return nil
	}
	out := &krm.GoogleCloudApigeeV1Operation{}
	out.Methods = in.Methods
	out.Resource = direct.LazyPtr(in.Resource)
	return out
}

func GoogleCloudApigeeV1Operation_ToApi(mapCtx *direct.MapContext, in *krm.GoogleCloudApigeeV1Operation) *api.GoogleCloudApigeeV1Operation {
	if in == nil {
		return nil
	}
	out := &api.GoogleCloudApigeeV1Operation{}
	out.Methods = in.Methods
	out.Resource = direct.ValueOf(in.Resource)
	return out
}

func GoogleCloudApigeeV1Operations_FromApi(mapCtx *direct.MapContext, in []*api.GoogleCloudApigeeV1Operation) []krm.GoogleCloudApigeeV1Operation {
	if in == nil {
		return nil
	}
	out := make([]krm.GoogleCloudApigeeV1Operation, len(in))
	for i, item := range in {
		if item != nil {
			out[i] = *GoogleCloudApigeeV1Operation_FromApi(mapCtx, item)
		}
	}
	return out
}

func GoogleCloudApigeeV1Operations_ToApi(mapCtx *direct.MapContext, in []krm.GoogleCloudApigeeV1Operation) []*api.GoogleCloudApigeeV1Operation {
	if in == nil {
		return nil
	}
	out := make([]*api.GoogleCloudApigeeV1Operation, len(in))
	for i := range in {
		out[i] = GoogleCloudApigeeV1Operation_ToApi(mapCtx, &in[i])
	}
	return out
}

func GoogleCloudApigeeV1OperationConfig_FromApi(mapCtx *direct.MapContext, in *api.GoogleCloudApigeeV1OperationConfig) *krm.GoogleCloudApigeeV1OperationConfig {
	if in == nil {
		return nil
	}
	out := &krm.GoogleCloudApigeeV1OperationConfig{}
	out.APISource = direct.LazyPtr(in.ApiSource)
	out.Attributes = GoogleCloudApigeeV1Attributes_FromApi(mapCtx, in.Attributes)
	out.Operations = GoogleCloudApigeeV1Operations_FromApi(mapCtx, in.Operations)
	out.Quota = GoogleCloudApigeeV1Quota_FromApi(mapCtx, in.Quota)
	return out
}

func GoogleCloudApigeeV1OperationConfig_ToApi(mapCtx *direct.MapContext, in *krm.GoogleCloudApigeeV1OperationConfig) *api.GoogleCloudApigeeV1OperationConfig {
	if in == nil {
		return nil
	}
	out := &api.GoogleCloudApigeeV1OperationConfig{}
	out.ApiSource = direct.ValueOf(in.APISource)
	out.Attributes = GoogleCloudApigeeV1Attributes_ToApi(mapCtx, in.Attributes)
	out.Operations = GoogleCloudApigeeV1Operations_ToApi(mapCtx, in.Operations)
	out.Quota = GoogleCloudApigeeV1Quota_ToApi(mapCtx, in.Quota)
	return out
}

func GoogleCloudApigeeV1OperationConfigs_FromApi(mapCtx *direct.MapContext, in []*api.GoogleCloudApigeeV1OperationConfig) []krm.GoogleCloudApigeeV1OperationConfig {
	if in == nil {
		return nil
	}
	out := make([]krm.GoogleCloudApigeeV1OperationConfig, len(in))
	for i, item := range in {
		if item != nil {
			out[i] = *GoogleCloudApigeeV1OperationConfig_FromApi(mapCtx, item)
		}
	}
	return out
}

func GoogleCloudApigeeV1OperationConfigs_ToApi(mapCtx *direct.MapContext, in []krm.GoogleCloudApigeeV1OperationConfig) []*api.GoogleCloudApigeeV1OperationConfig {
	if in == nil {
		return nil
	}
	out := make([]*api.GoogleCloudApigeeV1OperationConfig, len(in))
	for i := range in {
		out[i] = GoogleCloudApigeeV1OperationConfig_ToApi(mapCtx, &in[i])
	}
	return out
}

func GoogleCloudApigeeV1OperationGroup_FromApi(mapCtx *direct.MapContext, in *api.GoogleCloudApigeeV1OperationGroup) *krm.GoogleCloudApigeeV1OperationGroup {
	if in == nil {
		return nil
	}
	out := &krm.GoogleCloudApigeeV1OperationGroup{}
	out.OperationConfigType = direct.LazyPtr(in.OperationConfigType)
	out.OperationConfigs = GoogleCloudApigeeV1OperationConfigs_FromApi(mapCtx, in.OperationConfigs)
	return out
}

func GoogleCloudApigeeV1OperationGroup_ToApi(mapCtx *direct.MapContext, in *krm.GoogleCloudApigeeV1OperationGroup) *api.GoogleCloudApigeeV1OperationGroup {
	if in == nil {
		return nil
	}
	out := &api.GoogleCloudApigeeV1OperationGroup{}
	out.OperationConfigType = direct.ValueOf(in.OperationConfigType)
	out.OperationConfigs = GoogleCloudApigeeV1OperationConfigs_ToApi(mapCtx, in.OperationConfigs)
	return out
}

func GoogleCloudApigeeV1Quota_FromApi(mapCtx *direct.MapContext, in *api.GoogleCloudApigeeV1Quota) *krm.GoogleCloudApigeeV1Quota {
	if in == nil {
		return nil
	}
	out := &krm.GoogleCloudApigeeV1Quota{}
	out.Interval = direct.LazyPtr(in.Interval)
	out.Limit = direct.LazyPtr(in.Limit)
	out.TimeUnit = direct.LazyPtr(in.TimeUnit)
	return out
}

func GoogleCloudApigeeV1Quota_ToApi(mapCtx *direct.MapContext, in *krm.GoogleCloudApigeeV1Quota) *api.GoogleCloudApigeeV1Quota {
	if in == nil {
		return nil
	}
	out := &api.GoogleCloudApigeeV1Quota{}
	out.Interval = direct.ValueOf(in.Interval)
	out.Limit = direct.ValueOf(in.Limit)
	out.TimeUnit = direct.ValueOf(in.TimeUnit)
	return out
}

func ApigeeAPIProductSpec_FromApi(mapCtx *direct.MapContext, in *api.GoogleCloudApigeeV1ApiProduct) *krm.ApigeeAPIProductSpec {
	if in == nil {
		return nil
	}
	out := &krm.ApigeeAPIProductSpec{}
	out.APIResources = in.ApiResources
	out.ApprovalType = direct.LazyPtr(in.ApprovalType)
	out.Attributes = GoogleCloudApigeeV1Attributes_FromApi(mapCtx, in.Attributes)
	out.Description = direct.LazyPtr(in.Description)
	out.DisplayName = direct.LazyPtr(in.DisplayName)
	out.Environments = in.Environments
	out.GraphQLOperationGroup = GoogleCloudApigeeV1GraphQlOperationGroup_FromApi(mapCtx, in.GraphqlOperationGroup)
	out.GrpcOperationGroup = GoogleCloudApigeeV1GrpcOperationGroup_FromApi(mapCtx, in.GrpcOperationGroup)
	out.LlmOperationGroup = GoogleCloudApigeeV1LlmOperationGroup_FromApi(mapCtx, in.LlmOperationGroup)
	out.LlmQuota = direct.LazyPtr(in.LlmQuota)
	out.LlmQuotaInterval = direct.LazyPtr(in.LlmQuotaInterval)
	out.LlmQuotaTimeUnit = direct.LazyPtr(in.LlmQuotaTimeUnit)
	out.OperationGroup = GoogleCloudApigeeV1OperationGroup_FromApi(mapCtx, in.OperationGroup)
	out.Proxies = in.Proxies
	out.Quota = direct.LazyPtr(in.Quota)
	out.QuotaCounterScope = direct.LazyPtr(in.QuotaCounterScope)
	out.QuotaInterval = direct.LazyPtr(in.QuotaInterval)
	out.QuotaTimeUnit = direct.LazyPtr(in.QuotaTimeUnit)
	out.Scopes = in.Scopes

	return out
}

func ApigeeAPIProductSpec_ToApi(mapCtx *direct.MapContext, in *krm.ApigeeAPIProductSpec, name string) *api.GoogleCloudApigeeV1ApiProduct {
	if in == nil {
		return nil
	}
	out := &api.GoogleCloudApigeeV1ApiProduct{}
	out.Name = name
	out.ApiResources = in.APIResources
	out.ApprovalType = direct.ValueOf(in.ApprovalType)
	out.Attributes = GoogleCloudApigeeV1Attributes_ToApi(mapCtx, in.Attributes)
	out.Description = direct.ValueOf(in.Description)
	out.DisplayName = direct.ValueOf(in.DisplayName)
	out.Environments = in.Environments
	out.GraphqlOperationGroup = GoogleCloudApigeeV1GraphQlOperationGroup_ToApi(mapCtx, in.GraphQLOperationGroup)
	out.GrpcOperationGroup = GoogleCloudApigeeV1GrpcOperationGroup_ToApi(mapCtx, in.GrpcOperationGroup)
	out.LlmOperationGroup = GoogleCloudApigeeV1LlmOperationGroup_ToApi(mapCtx, in.LlmOperationGroup)
	out.LlmQuota = direct.ValueOf(in.LlmQuota)
	out.LlmQuotaInterval = direct.ValueOf(in.LlmQuotaInterval)
	out.LlmQuotaTimeUnit = direct.ValueOf(in.LlmQuotaTimeUnit)
	out.OperationGroup = GoogleCloudApigeeV1OperationGroup_ToApi(mapCtx, in.OperationGroup)
	out.Proxies = in.Proxies
	out.Quota = direct.ValueOf(in.Quota)
	out.QuotaCounterScope = direct.ValueOf(in.QuotaCounterScope)
	out.QuotaInterval = direct.ValueOf(in.QuotaInterval)
	out.QuotaTimeUnit = direct.ValueOf(in.QuotaTimeUnit)
	out.Scopes = in.Scopes

	return out
}

func ApigeeAPIProductObservedState_FromApi(mapCtx *direct.MapContext, in *api.GoogleCloudApigeeV1ApiProduct) *krm.ApigeeAPIProductObservedState {
	if in == nil {
		return nil
	}
	out := &krm.ApigeeAPIProductObservedState{}
	out.CreatedAt = direct.LazyPtr(in.CreatedAt)
	out.LastModifiedAt = direct.LazyPtr(in.LastModifiedAt)
	return out
}

func ApigeeAPIProductObservedState_ToApi(mapCtx *direct.MapContext, in *krm.ApigeeAPIProductObservedState) *api.GoogleCloudApigeeV1ApiProduct {
	if in == nil {
		return nil
	}
	out := &api.GoogleCloudApigeeV1ApiProduct{}
	out.CreatedAt = direct.ValueOf(in.CreatedAt)
	out.LastModifiedAt = direct.ValueOf(in.LastModifiedAt)
	return out
}
