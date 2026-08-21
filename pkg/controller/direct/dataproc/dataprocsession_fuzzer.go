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
// proto.message: google.cloud.dataproc.v1.Session
// api.group: dataproc.cnrm.cloud.google.com

package dataproc

import (
	pb "cloud.google.com/go/dataproc/v2/apiv1/dataprocpb"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/fuzztesting"
)

func init() {
	fuzztesting.RegisterKRMSpecFuzzer(dataprocSessionFuzzer())
}

func dataprocSessionFuzzer() fuzztesting.KRMFuzzer {
	f := fuzztesting.NewKRMTypedSpecFuzzer(&pb.Session{},
		DataprocSessionSpec_v1alpha1_FromProto, DataprocSessionSpec_v1alpha1_ToProto,
	)

	f.SpecField(".jupyter_session")
	f.SpecField(".labels")
	f.SpecField(".runtime_config")
	f.SpecField(".environment_config")
	f.SpecField(".user")
	f.SpecField(".session_template")

	f.Unimplemented_Identity(".name")
	f.Unimplemented_Identity(".uuid")

	// Mapped under observedState / output-only
	f.Unimplemented_Identity(".create_time")
	f.Unimplemented_Identity(".state")
	f.Unimplemented_Identity(".state_message")
	f.Unimplemented_Identity(".state_time")
	f.Unimplemented_Identity(".creator")
	f.Unimplemented_Identity(".state_history")
	f.Unimplemented_Identity(".runtime_info")

	// Missing / Not yet supported
	f.Unimplemented_NotYetTriaged(".spark_connect_session")
	f.Unimplemented_NotYetTriaged(".environment_config.execution_config.authentication_config")
	f.Unimplemented_NotYetTriaged(".environment_config.execution_config.resource_manager_tags")

	return f
}
