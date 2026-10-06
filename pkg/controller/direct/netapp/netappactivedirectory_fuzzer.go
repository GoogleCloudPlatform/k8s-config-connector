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

	f.IdentityField(".name")

	f.SpecField(".domain")
	f.SpecField(".site")
	f.SpecField(".dns")
	f.SpecField(".net_bios_prefix")
	f.SpecField(".organizational_unit")
	f.SpecField(".aes_encryption")
	f.SpecField(".username")
	f.SpecField(".password")
	f.SpecField(".backup_operators")
	f.SpecField(".administrators")
	f.SpecField(".security_operators")
	f.SpecField(".kdc_hostname")
	f.SpecField(".kdc_ip")
	f.SpecField(".nfs_users_with_ldap")
	f.SpecField(".description")
	f.SpecField(".ldap_signing")
	f.SpecField(".encrypt_dc_connections")
	f.SpecField(".labels")

	f.StatusField(".create_time")
	f.StatusField(".state")
	f.StatusField(".state_details")

	return f
}
