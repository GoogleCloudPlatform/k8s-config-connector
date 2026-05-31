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

package session

import (
	pb "cloud.google.com/go/dataproc/v2/apiv1/dataprocpb"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/fuzztesting"
)

func init() {
	fuzztesting.RegisterKRMFuzzer(dataprocSessionFuzzer())
}

func dataprocSessionFuzzer() fuzztesting.KRMFuzzer {
	f := fuzztesting.NewKRMTypedFuzzer(&pb.Session{},
		DataprocSessionSpec_FromProto, DataprocSessionSpec_ToProto,
		DataprocSessionObservedState_FromProto, DataprocSessionObservedState_ToProto,
	)

	f.Unimplemented_Identity(".name")

	f.Unimplemented_LabelsAnnotations(".labels")

	f.SpecField(".jupyter_session")
	f.SpecField(".runtime_config")
	f.SpecField(".environment_config")
	f.SpecField(".user")
	f.SpecField(".session_template")

	f.Unimplemented_NotYetTriaged(".spark_connect_session")
	f.Unimplemented_NotYetTriaged(".environment_config.execution_config.authentication_config")

	f.StatusField(".uuid")
	f.StatusField(".create_time")
	f.StatusField(".runtime_info")
	f.StatusField(".state")
	f.StatusField(".state_message")
	f.StatusField(".state_time")
	f.StatusField(".creator")
	f.StatusField(".state_history")

	return f
}
