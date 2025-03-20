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
// proto.message: google.cloud.netapp.v1.ActiveDirectory
// api.group: netapp.cnrm.cloud.google.com

package netapp

import (
	pb "cloud.google.com/go/netapp/apiv1/netapppb"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/fuzztesting"
)

func init() {
	fuzztesting.RegisterKRMFuzzer(netAppActiveDirectoryFuzzer())
}

func netAppActiveDirectoryFuzzer() fuzztesting.KRMFuzzer {
	f := fuzztesting.NewKRMTypedFuzzer(&pb.ActiveDirectory{},
		NetAppActiveDirectorySpec_FromProto, NetAppActiveDirectorySpec_ToProto,
		NetAppActiveDirectoryObservedState_FromProto, NetAppActiveDirectoryObservedState_ToProto,
	)

	f.SpecFields.Insert(".domain")
	f.SpecFields.Insert(".site")
	f.SpecFields.Insert(".dns")
	f.SpecFields.Insert(".net_bios_prefix")
	f.SpecFields.Insert(".organizational_unit")
	f.SpecFields.Insert(".aes_encryption")
	f.SpecFields.Insert(".username")
	f.SpecFields.Insert(".password")
	f.SpecFields.Insert(".backup_operators")
	f.SpecFields.Insert(".administrators")
	f.SpecFields.Insert(".security_operators")
	f.SpecFields.Insert(".kdc_hostname")
	f.SpecFields.Insert(".kdc_ip")
	f.SpecFields.Insert(".nfs_users_with_ldap")
	f.SpecFields.Insert(".description")
	f.SpecFields.Insert(".ldap_signing")
	f.SpecFields.Insert(".encrypt_dc_connections")

	f.UnimplementedFields.Insert(".name")          // special field
	f.UnimplementedFields.Insert(".create_time")   // output only
	f.UnimplementedFields.Insert(".state")         // output only
	f.UnimplementedFields.Insert(".state_details") // output only
	f.UnimplementedFields.Insert(".labels")        // NOTYET
	return f
}
