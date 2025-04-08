// Copyright 2024 Google LLC
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//	http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

// +tool:fuzz-gen
// proto.message: google.cloud.dataplex.v1.Task
// api.group: dataplex.cnrm.cloud.google.com

package dataplex

import (
	pb "cloud.google.com/go/dataplex/apiv1/dataplexpb"
	krmv1alpha1 "github.com/GoogleCloudPlatform/k8s-config-connector/apis/dataplex/v1alpha1"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/fuzztesting"
)

func init() {
	fuzztesting.RegisterKRMFuzzer(dataplexTaskFuzzer())
}

func dataplexTaskFuzzer() fuzztesting.KRMFuzzer {
	f := fuzztesting.NewKRMTypedFuzzer(&pb.Task{},
		krmv1alpha1.DataplexTaskSpec_FromProto, krmv1alpha1.DataplexTaskSpec_ToProto,
		krmv1alpha1.DataplexTaskObservedState_FromProto, krmv1alpha1.DataplexTaskObservedState_ToProto,
	)

	f.SpecFields.Insert(".description")
	f.SpecFields.Insert(".display_name")
	f.SpecFields.Insert(".labels")
	f.SpecFields.Insert(".trigger_spec")
	f.SpecFields.Insert(".execution_spec")
	f.SpecFields.Insert(".spark")
	f.SpecFields.Insert(".notebook")

	f.StatusFields.Insert(".uid")
	f.StatusFields.Insert(".create_time")
	f.StatusFields.Insert(".update_time")
	f.StatusFields.Insert(".state")
	f.StatusFields.Insert(".execution_status")

	f.UnimplementedFields.Insert(".name")
	f.UnimplementedFields.Insert(".execution_spec.kms_key")               // type *refsv1beta1.KMSCryptoKeyRef in krm
	f.UnimplementedFields.Insert(".notebook.infrastructure_spec.network") // type oneof in proto
	f.UnimplementedFields.Insert(".spark.infrastructure_spec.network")    // type oneof in proto

	return f
}
