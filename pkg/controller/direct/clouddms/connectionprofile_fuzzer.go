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
// proto.message: google.cloud.clouddms.v1.ConnectionProfile
// api.group: clouddms.cnrm.cloud.google.com

package clouddms

import (
	pb "cloud.google.com/go/clouddms/apiv1/clouddmspb"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/fuzztesting"
)

func init() {
	fuzztesting.RegisterKRMFuzzer(cloudDMSConnectionProfileFuzzer())
}

func cloudDMSConnectionProfileFuzzer() fuzztesting.KRMFuzzer {
	f := fuzztesting.NewKRMTypedFuzzer(&pb.ConnectionProfile{},
		CloudDMSConnectionProfileSpec_FromProto, CloudDMSConnectionProfileSpec_ToProto,
		CloudDMSConnectionProfileObservedState_FromProto, CloudDMSConnectionProfileObservedState_ToProto,
	)

	f.IdentityField(".name")

	f.SpecField(".labels")
	f.SpecField(".display_name")
	f.SpecField(".mysql")
	f.SpecField(".postgresql")
	f.SpecField(".oracle")
	f.SpecField(".cloudsql")
	f.SpecField(".alloydb")
	f.SpecField(".provider")

	f.StatusField(".create_time")
	f.StatusField(".update_time")
	f.StatusField(".state")
	f.StatusField(".error")

	f.Unimplemented_NotYetTriaged(".alloydb.settings.initial_user.password_set")
	f.Unimplemented_NotYetTriaged(".alloydb.settings.primary_instance_settings.private_ip")
	f.Unimplemented_NotYetTriaged(".cloudsql.additional_public_ip")
	f.Unimplemented_NotYetTriaged(".cloudsql.cloud_sql_id")
	f.Unimplemented_NotYetTriaged(".cloudsql.private_ip")
	f.Unimplemented_NotYetTriaged(".cloudsql.public_ip")
	f.Unimplemented_NotYetTriaged(".cloudsql.settings.root_password_set")
	f.Unimplemented_NotYetTriaged(".error.details")
	f.Unimplemented_NotYetTriaged(".mysql.password_set")
	f.Unimplemented_NotYetTriaged(".mysql.ssl.type")
	f.Unimplemented_NotYetTriaged(".oracle.password_set")
	f.Unimplemented_NotYetTriaged(".oracle.ssl.type")
	f.Unimplemented_NotYetTriaged(".postgresql.network_architecture")
	f.Unimplemented_NotYetTriaged(".postgresql.password_set")
	f.Unimplemented_NotYetTriaged(".postgresql.ssl.type")

	return f
}
