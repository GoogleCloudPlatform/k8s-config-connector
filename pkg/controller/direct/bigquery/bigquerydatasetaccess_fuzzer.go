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
// proto.message: google.cloud.bigquery.v2.Access
// api.group: bigquery.cnrm.cloud.google.com

package bigquery

import (
	api "google.golang.org/api/bigquery/v2"

	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/fuzztesting"
)

func init() {
	fuzztesting.RegisterKRMFuzzer_NoProto(bigqueryDatasetAccessFuzzer())
}

func bigqueryDatasetAccessFuzzer() fuzztesting.KRMFuzzer_NoProto {
	f := fuzztesting.NewKRMTypedFuzzer_NoProto(&api.DatasetAccess{},
		BigQueryDatasetAccessSpec_FromAPI, BigQueryDatasetAccessSpec_ToAPI,
		BigQueryDatasetAccessStatus_FromAPI, BigQueryDatasetAccessStatus_ToAPI,
	)

	// Field Comparison & Mapping Status:
	// KRM Spec Fields:
	// - dataset: Maps to .Dataset
	// - datasetId: Parent dataset ID (KCC hierarchical identity field, not part of DatasetAccess body)
	// - domain: Maps to .Domain
	// - groupByEmail: Maps to .GroupByEmail
	// - iamMember: Maps to .IamMember
	// - projectRef: Parent project ref (KCC hierarchical identity field, not part of DatasetAccess body)
	// - resourceID: Resource identity (KCC identity field, not part of DatasetAccess body)
	// - role: Maps to .Role
	// - specialGroup: Maps to .SpecialGroup
	// - userByEmail: Maps to .UserByEmail
	// - view: Maps to .View
	//
	// KRM Status Fields:
	// - apiUpdatedMember: KCC client-side status flag, not an API field
	// - conditions: Standard KCC Status condition list, not an API field
	// - observedGeneration: Standard KCC Status generation tracking, not an API field

	// Spec fields
	f.SpecField(".Dataset")
	f.SpecField(".Domain")
	f.SpecField(".GroupByEmail")
	f.SpecField(".IamMember")
	f.SpecField(".Role")
	f.SpecField(".SpecialGroup")
	f.SpecField(".UserByEmail")
	f.SpecField(".View")

	// Untriaged / unimplemented API fields
	f.Unimplemented_NotYetTriaged(".Condition")
	f.Unimplemented_NotYetTriaged(".Routine")

	// JSON bookkeeping fields
	f.Ignore_JSONBookkeeping(".ForceSendFields")
	f.Ignore_JSONBookkeeping(".NullFields")
	f.Ignore_JSONBookkeeping(".Dataset.ForceSendFields")
	f.Ignore_JSONBookkeeping(".Dataset.NullFields")
	f.Ignore_JSONBookkeeping(".Dataset.Dataset.ForceSendFields")
	f.Ignore_JSONBookkeeping(".Dataset.Dataset.NullFields")
	f.Ignore_JSONBookkeeping(".View.ForceSendFields")
	f.Ignore_JSONBookkeeping(".View.NullFields")

	f.FilterSpec = func(in *api.DatasetAccess) {
		if in.Dataset != nil && in.Dataset.Dataset == nil {
			in.Dataset.Dataset = &api.DatasetReference{}
		}
	}

	return f
}
