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

// +tool:fuzz-gen
// proto.message: google.cloud.bigquery.v2.RowAccessPolicy
// api.group: bigquery.cnrm.cloud.google.com

package bigquery

import (
	api "google.golang.org/api/bigquery/v2"

	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/fuzztesting"
)

func init() {
	fuzztesting.RegisterKRMFuzzer_NoProto(bigqueryRowAccessPolicyFuzzer())
}

func bigqueryRowAccessPolicyFuzzer() fuzztesting.KRMFuzzer_NoProto {
	f := fuzztesting.NewKRMTypedFuzzer_NoProto(&api.RowAccessPolicy{},
		BigQueryRowAccessPolicySpec_FromAPI, BigQueryRowAccessPolicySpec_ToAPI,
		BigQueryRowAccessPolicyStatus_FromAPI, BigQueryRowAccessPolicyStatus_ToAPI,
	)

	// Spec fields
	f.SpecField(".FilterPredicate")
	f.SpecField(".Grantees")

	// Status fields
	f.StatusField(".CreationTime")
	f.StatusField(".Etag")
	f.StatusField(".LastModifiedTime")

	// Fields handled by hierarchical identity (ParentRef / ResourceID)
	f.IdentityField(".RowAccessPolicyReference.DatasetId")
	f.IdentityField(".RowAccessPolicyReference.PolicyId")
	f.IdentityField(".RowAccessPolicyReference.ProjectId")
	f.IdentityField(".RowAccessPolicyReference.TableId")

	// JSON bookkeeping fields
	f.Ignore_JSONBookkeeping(".ForceSendFields")
	f.Ignore_JSONBookkeeping(".NullFields")
	f.Ignore_JSONBookkeeping(".RowAccessPolicyReference.ForceSendFields")
	f.Ignore_JSONBookkeeping(".RowAccessPolicyReference.NullFields")

	f.FilterSpec = func(in *api.RowAccessPolicy) {
		in.RowAccessPolicyReference = nil
	}
	f.FilterStatus = func(in *api.RowAccessPolicy) {
		in.RowAccessPolicyReference = nil
	}

	return f
}
