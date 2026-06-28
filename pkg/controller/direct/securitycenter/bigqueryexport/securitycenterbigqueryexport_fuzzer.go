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

package bigqueryexport

import (
	pb "cloud.google.com/go/securitycenter/apiv1/securitycenterpb"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct/securitycenter"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/fuzztesting"
)

func init() {
	fuzztesting.RegisterKRMFuzzer(securitycenterBigQueryExportFuzzer())
}

func securitycenterBigQueryExportFuzzer() fuzztesting.KRMFuzzer {
	f := fuzztesting.NewKRMTypedFuzzer(&pb.BigQueryExport{},
		securitycenter.SecurityCenterBigQueryExportSpec_FromProto, securitycenter.SecurityCenterBigQueryExportSpec_ToProto,
		securitycenter.SecurityCenterBigQueryExportObservedState_FromProto, securitycenter.SecurityCenterBigQueryExportObservedState_ToProto,
	)

	f.UnimplementedFields.Insert(".name")

	f.SpecFields.Insert(".dataset")
	f.SpecFields.Insert(".description")
	f.SpecFields.Insert(".filter")

	f.StatusFields.Insert(".create_time")
	f.StatusFields.Insert(".update_time")
	f.StatusFields.Insert(".most_recent_editor")
	f.StatusFields.Insert(".principal")

	return f
}
