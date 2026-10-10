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
// proto.message: google.cloud.oslogin.common.SshPublicKey
// api.group: oslogin.cnrm.cloud.google.com

package oslogin

import (
	commonpb "cloud.google.com/go/oslogin/common/commonpb"
	krm "github.com/GoogleCloudPlatform/k8s-config-connector/apis/oslogin/v1alpha1"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/fuzztesting"
)

func init() {
	fuzztesting.RegisterKRMFuzzer(osloginSSHPublicKeyFuzzer())
}

func osloginSSHPublicKeyFuzzer() fuzztesting.KRMFuzzer {
	f := fuzztesting.NewKRMTypedFuzzer[*commonpb.SshPublicKey, krm.OSLoginSSHPublicKeySpec, krm.OSLoginSSHPublicKeyStatus](&commonpb.SshPublicKey{},
		OSLoginSSHPublicKeySpec_v1alpha1_FromProto, OSLoginSSHPublicKeySpec_v1alpha1_ToProto,
		OSLoginSSHPublicKeyStatus_v1alpha1_FromProto, OSLoginSSHPublicKeyStatus_v1alpha1_ToProto,
	)

	// Field comparison: OSLoginSSHPublicKeySpec vs commonpb.SshPublicKey Proto
	// - Spec.ExpirationTimeUsec maps to proto field .expiration_time_usec
	// - Spec.Key                maps to proto field .key
	// - Spec.Project            not represented in commonpb.SshPublicKey (KCC project reference)
	// - Spec.ResourceID         maps to proto field .fingerprint (service-generated identifier for acquisition)
	// - Spec.User               not represented in commonpb.SshPublicKey (parent user email identifier)

	// Field comparison: OSLoginSSHPublicKeyStatus vs commonpb.SshPublicKey Proto
	// - Status.Conditions       not represented in commonpb.SshPublicKey
	// - Status.Fingerprint      maps to proto field .fingerprint
	// - Status.ObservedGeneration not represented in commonpb.SshPublicKey

	// Proto fields comparison:
	// - .key                    maps to Spec.Key
	// - .expiration_time_usec   maps to Spec.ExpirationTimeUsec
	// - .fingerprint            maps to Status.Fingerprint
	// - .name                   identity field (users/{user}/sshPublicKeys/{fingerprint})

	f.SpecField(".key")
	f.SpecField(".expiration_time_usec")

	f.StatusField(".fingerprint")

	f.Unimplemented_Identity(".name")

	return f
}
