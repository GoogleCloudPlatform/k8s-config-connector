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
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/fuzztesting"
	api "google.golang.org/api/apigee/v1"
)

func init() {
	fuzztesting.RegisterKRMFuzzer_NoProto(apiproductFuzzer())
}

func apiproductFuzzer() fuzztesting.KRMFuzzer_NoProto {
	f := fuzztesting.NewKRMTypedFuzzer_NoProto(&api.GoogleCloudApigeeV1ApiProduct{},
		func(ctx *direct.MapContext, in *api.GoogleCloudApigeeV1ApiProduct) *krm.ApigeeAPIProductSpec {
			return ApigeeAPIProductSpec_FromApi(ctx, in)
		},
		func(ctx *direct.MapContext, in *krm.ApigeeAPIProductSpec) *api.GoogleCloudApigeeV1ApiProduct {
			resourceID := ""
			if in.ResourceID != nil {
				resourceID = *in.ResourceID
			}
			return ApigeeAPIProductSpec_ToApi(ctx, in, resourceID)
		},
		ApigeeAPIProductObservedState_FromApi, ApigeeAPIProductObservedState_ToApi,
	)

	f.SpecField(".ApiResources")
	f.SpecField(".ApprovalType")
	f.SpecField(".Attributes")
	f.SpecField(".Description")
	f.SpecField(".DisplayName")
	f.SpecField(".Environments")
	f.SpecField(".GraphqlOperationGroup")
	f.SpecField(".GrpcOperationGroup")
	f.SpecField(".LlmOperationGroup")
	f.SpecField(".LlmQuota")
	f.SpecField(".LlmQuotaInterval")
	f.SpecField(".LlmQuotaTimeUnit")
	f.SpecField(".OperationGroup")
	f.SpecField(".Proxies")
	f.SpecField(".Quota")
	f.SpecField(".QuotaCounterScope")
	f.SpecField(".QuotaInterval")
	f.SpecField(".QuotaTimeUnit")
	f.SpecField(".Scopes")

	f.StatusField(".CreatedAt")
	f.StatusField(".LastModifiedAt")

	f.IdentityField(".Name")
	f.Unimplemented_NotYetTriaged(".Space")
	f.Unimplemented_NotYetTriaged(".PayloadOperationGroup")

	f.Ignore_JSONBookkeeping(".ForceSendFields")
	f.Ignore_JSONBookkeeping(".NullFields")
	f.Ignore_JSONBookkeeping(".ServerResponse")

	return f
}
