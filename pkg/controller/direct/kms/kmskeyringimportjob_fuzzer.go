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
// proto.message: google.cloud.kms.v1.ImportJob
// api.group: kms.cnrm.cloud.google.com

package kms

import (
	pb "cloud.google.com/go/kms/apiv1/kmspb"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/fuzztesting"
)

func init() {
	fuzztesting.RegisterKRMFuzzer(KMSKeyRingImportJobFuzzer())
}

func KMSKeyRingImportJobFuzzer() fuzztesting.KRMFuzzer {
	f := fuzztesting.NewKRMTypedFuzzer(&pb.ImportJob{},
		KMSKeyRingImportJobSpec_FromProto, KMSKeyRingImportJobSpec_ToProto,
		KMSKeyRingImportJobStatus_FromProto, KMSKeyRingImportJobStatus_ToProto,
	)

	// Spec fields mapping:
	// - importJobId                       -> (part of identity / .name resource ID)
	// - keyRing                           -> (part of parent path of .name GCP resource identity)
	// - resourceID                        -> (part of identity / .name resource ID)
	// - importMethod                      -> .import_method
	// - protectionLevel                   -> .protection_level
	f.SpecField(".import_method")
	f.SpecField(".protection_level")

	// Status fields mapping:
	// - attestation                       -> .attestation
	//   - content                         -> .attestation.content (custom Base64-encoded in KRM)
	//   - format                          -> .attestation.format
	// - expireTime                        -> .expire_time
	// - name                              -> .name (Identity)
	// - publicKey                         -> .public_key
	//   - pem                             -> .public_key.pem
	// - state                             -> .state
	f.StatusField(".expire_time")
	f.StatusField(".state")
	f.StatusField(".public_key")
	f.StatusField(".attestation")

	// Identity / Special fields
	f.Unimplemented_Identity(".name")

	// Unimplemented / Not Yet Triaged fields
	f.Unimplemented_Internal(".create_time")
	f.Unimplemented_Internal(".generate_time")
	f.Unimplemented_Internal(".expire_event_time")
	f.Unimplemented_NotYetTriaged(".crypto_key_backend")
	f.Unimplemented_NotYetTriaged(".attestation.cert_chains")
	f.Unimplemented_NotYetTriaged(".attestation.cert_chains.cavium_certs")
	f.Unimplemented_NotYetTriaged(".attestation.cert_chains.google_card_certs")
	f.Unimplemented_NotYetTriaged(".attestation.cert_chains.google_partition_certs")

	return f
}
