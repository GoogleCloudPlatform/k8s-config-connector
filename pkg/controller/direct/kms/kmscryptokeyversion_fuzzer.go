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
// proto.message: google.cloud.kms.v1.CryptoKeyVersion
// api.group: kms.cnrm.cloud.google.com

package kms

import (
	pb "cloud.google.com/go/kms/apiv1/kmspb"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/fuzztesting"
)

func init() {
	fuzztesting.RegisterKRMFuzzer(KMSCryptoKeyVersionFuzzer())
}

func KMSCryptoKeyVersionFuzzer() fuzztesting.KRMFuzzer {
	f := fuzztesting.NewKRMTypedFuzzer(&pb.CryptoKeyVersion{},
		KMSCryptoKeyVersionSpec_FromProto, KMSCryptoKeyVersionSpec_ToProto,
		KMSCryptoKeyVersionStatus_FromProto, KMSCryptoKeyVersionStatus_ToProto,
	)

	// Field comparison:
	// KMSCryptoKeyVersionSpec:
	// - .cryptoKey: KRM-only reference to parent, resolved to parent name.
	// - .resourceID: KRM-only field for acquisition.
	// - .state: maps to .state in the proto (SpecField).
	// KMSCryptoKeyVersionStatus:
	// - .algorithm: maps to .algorithm (StatusField).
	// - .attestation: maps to .attestation (StatusField) and its nested fields.
	// - .generateTime: maps to .generate_time (StatusField).
	// - .name: maps to .name (StatusField / Identity).
	// - .protectionLevel: maps to .protection_level (StatusField).
	f.SpecField(".state")

	f.StatusField(".name")
	f.StatusField(".protection_level")
	f.StatusField(".algorithm")
	f.StatusField(".generate_time")

	f.StatusField(".attestation")
	f.StatusField(".attestation.format")
	f.StatusField(".attestation.content")
	f.StatusField(".attestation.cert_chains")
	f.StatusField(".attestation.cert_chains.cavium_certs")
	f.StatusField(".attestation.cert_chains.google_card_certs")
	f.StatusField(".attestation.cert_chains.google_partition_certs")

	f.StatusField(".external_protection_level_options")
	f.StatusField(".external_protection_level_options.ekm_connection_key_path")
	f.StatusField(".external_protection_level_options.external_key_uri")

	f.Unimplemented_Identity(".name")

	f.Unimplemented_NotYetTriaged(".create_time")
	f.Unimplemented_NotYetTriaged(".destroy_time")
	f.Unimplemented_NotYetTriaged(".destroy_event_time")
	f.Unimplemented_NotYetTriaged(".import_job")
	f.Unimplemented_NotYetTriaged(".import_time")
	f.Unimplemented_NotYetTriaged(".import_failure_reason")
	f.Unimplemented_NotYetTriaged(".generation_failure_reason")
	f.Unimplemented_NotYetTriaged(".external_destruction_failure_reason")
	f.Unimplemented_NotYetTriaged(".reimport_eligible")

	f.FilterStatus = func(in *pb.CryptoKeyVersion) {
		if in.ExternalProtectionLevelOptions != nil && in.Attestation == nil {
			in.Attestation = &pb.KeyOperationAttestation{}
		}

		if in.Attestation != nil && in.Attestation.CertChains != nil {
			cc := in.Attestation.CertChains
			if len(cc.CaviumCerts) > 1 {
				cc.CaviumCerts = cc.CaviumCerts[:1]
			}
			if len(cc.GoogleCardCerts) > 1 {
				cc.GoogleCardCerts = cc.GoogleCardCerts[:1]
			}
			if len(cc.GooglePartitionCerts) > 1 {
				cc.GooglePartitionCerts = cc.GooglePartitionCerts[:1]
			}
		}
	}

	return f
}
