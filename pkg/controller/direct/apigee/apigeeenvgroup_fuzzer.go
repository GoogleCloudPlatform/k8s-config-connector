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
	krm "github.com/GoogleCloudPlatform/k8s-config-connector/apis/apigee/v1beta1"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/fuzztesting"
	api "google.golang.org/api/apigee/v1"
)

func init() {
	fuzztesting.RegisterKRMFuzzer_NoProto(envgroupFuzzer())
}

func envgroupFuzzer() fuzztesting.KRMFuzzer_NoProto {
	f := fuzztesting.NewKRMTypedFuzzer_NoProto(&api.GoogleCloudApigeeV1EnvironmentGroup{},
		func(ctx *direct.MapContext, in *api.GoogleCloudApigeeV1EnvironmentGroup) *krm.ApigeeEnvgroupSpec {
			return ApigeeEnvgroupSpec_FromApi(ctx, in)
		},
		func(ctx *direct.MapContext, in *krm.ApigeeEnvgroupSpec) *api.GoogleCloudApigeeV1EnvironmentGroup {
			resourceID := ""
			if in.ResourceID != nil {
				resourceID = *in.ResourceID
			}
			return ApigeeEnvgroupSpec_ToApi(ctx, in, resourceID)
		},
		ApigeeEnvgroupObservedState_FromApi, ApigeeEnvgroupObservedState_ToApi,
	)

	f.SpecField(".Hostnames")

	f.StatusField(".CreatedAt")
	f.StatusField(".LastModifiedAt")
	f.StatusField(".State")
	f.StatusField(".Name")

	f.Ignore_JSONBookkeeping(".ForceSendFields")
	f.Ignore_JSONBookkeeping(".NullFields")
	f.Ignore_JSONBookkeeping(".ServerResponse")

	return f
}
