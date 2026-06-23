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
// proto.message: google.bigtable.admin.v2.GcRule
// api.group: bigtable.cnrm.cloud.google.com

package bigtable

import (
	pb "cloud.google.com/go/bigtable/admin/apiv2/adminpb"

	krm "github.com/GoogleCloudPlatform/k8s-config-connector/apis/bigtable/v1beta1"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/fuzztesting"
)

func init() {
	fuzztesting.RegisterKRMSpecFuzzer(bigtableGCPolicyFuzzer())
}

func bigtableGCPolicyFuzzer() fuzztesting.KRMFuzzer {
	f := fuzztesting.NewKRMTypedFuzzer[*pb.GcRule, krm.BigtableGCPolicySpec, krm.BigtableGCPolicyStatus](&pb.GcRule{},
		BigtableGCPolicySpec_v1beta1_FromProto, BigtableGCPolicySpec_v1beta1_ToProto,
		nil, nil,
	)

	// KRM Spec type fields correspond to fields in the fuzzer/proto:
	// - MaxVersion maps to GcRule.Rule.GcRule_MaxNumVersions (max_num_versions)
	// - MaxAge maps to GcRule.Rule.GcRule_MaxAge (max_age)
	f.SpecField(".max_num_versions")
	f.SpecField(".max_age")

	// Unimplemented fields
	f.Unimplemented_NotYetTriaged(".intersection")
	f.Unimplemented_NotYetTriaged(".union")

	return f
}
