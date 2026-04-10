// Copyright 2024 Google LLC
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

package dataform

import (
	dataformpb "cloud.google.com/go/dataform/apiv1beta1/dataformpb"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/fuzztesting"
)

func init() {
	fuzztesting.RegisterKRMSpecFuzzer(repositoryFuzzer())
}

func repositoryFuzzer() fuzztesting.KRMFuzzer {
	fuzzer := fuzztesting.NewKRMTypedSpecFuzzer(
		&dataformpb.Repository{},
		DataformRepositorySpec_FromProto,
		DataformRepositorySpec_ToProto,
	)

	fuzzer.UnimplementedFields.Insert(".name")
	fuzzer.UnimplementedFields.Insert(".labels")
	fuzzer.UnimplementedFields.Insert(".git_remote_settings.token_status") // not supported by KCC anymore
	fuzzer.UnimplementedFields.Insert(".create_time")
	fuzzer.UnimplementedFields.Insert(".data_encryption_state")
	fuzzer.UnimplementedFields.Insert(".kms_key_name")
	fuzzer.UnimplementedFields.Insert(".internal_metadata")

	return fuzzer
}
