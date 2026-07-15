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
// proto.message: google.cloud.securesourcemanager.v1.Instance
// api.group: securesourcemanager.cnrm.cloud.google.com

package securesourcemanager

import (
	pb "cloud.google.com/go/securesourcemanager/apiv1/securesourcemanagerpb"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/fuzztesting"
)

func init() {
	fuzztesting.RegisterKRMFuzzer(secureSourceManagerInstanceFuzzer())
}

func secureSourceManagerInstanceFuzzer() fuzztesting.KRMFuzzer {
	f := fuzztesting.NewKRMTypedFuzzer(&pb.Instance{},
		SecureSourceManagerInstanceSpec_FromProto, SecureSourceManagerInstanceSpec_ToProto,
		SecureSourceManagerInstanceObservedState_FromProto, SecureSourceManagerInstanceObservedState_ToProto,
	)

	// Identity Field
	f.IdentityField(".name")

	// Field Comparison Map (KRM Spec -> Proto Field):
	// - kmsKeyRef                      -> .kms_key
	// - labels                         -> .labels
	// - location                       -> [KRM resource path / identifier]
	// - projectRef                     -> [KRM resource path / identifier]
	// - resourceID                     -> [KRM resource path / identifier]
	// - privateConfig.caPoolRef        -> .private_config.ca_pool
	// - privateConfig.isPrivate        -> .private_config.is_private
	f.SpecField(".labels")
	f.SpecField(".kms_key")
	f.SpecField(".private_config.ca_pool")
	f.SpecField(".private_config.is_private")

	// Field Comparison Map (KRM ObservedState -> Proto Field):
	// - createTime                              -> .create_time
	// - updateTime                              -> .update_time
	// - state                                   -> .state
	// - stateNote                               -> .state_note
	// - hostConfig                              -> .host_config
	// - privateConfig.httpServiceAttachment     -> .private_config.http_service_attachment
	// - privateConfig.sshServiceAttachment      -> .private_config.ssh_service_attachment
	f.StatusField(".create_time")
	f.StatusField(".update_time")
	f.StatusField(".state")
	f.StatusField(".state_note")
	f.StatusField(".host_config")
	f.StatusField(".private_config.http_service_attachment")
	f.StatusField(".private_config.ssh_service_attachment")

	// Unimplemented and Not Yet Triaged fields in GCP proto
	f.Unimplemented_NotYetTriaged(".private_config.psc_allowed_projects")
	f.Unimplemented_NotYetTriaged(".workforce_identity_federation_config")
	f.Unimplemented_NotYetTriaged(".private_config.custom_host_config")

	return f
}
