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
// proto.message: google.cloud.sql.v1beta4.SslCert
// api.group: sql.cnrm.cloud.google.com

package sql

import (
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/fuzztesting"
	pb "github.com/GoogleCloudPlatform/k8s-config-connector/pkg/gcpclients/generated/google/cloud/sql/v1beta4"
)

func init() {
	fuzztesting.RegisterKRMFuzzer(sqlSSLCertFuzzer())
}

func sqlSSLCertFuzzer() fuzztesting.KRMFuzzer {
	f := fuzztesting.NewKRMTypedFuzzer(&pb.SslCert{},
		SQLSSLCertSpec_FromProto, SQLSSLCertSpec_ToProto,
		SQLSSLCertStatus_FromProto, SQLSSLCertStatus_ToProto,
	)

	// KRM Spec field comparison:
	// - CommonName maps to .common_name
	// - InstanceRef is KRM reference to Cloud SQL instance, not part of pb.SslCert body directly (proto .instance is resource path/name)
	// - ResourceID is KCC-specific resource ID and not mapped directly in spec
	f.SpecField(".common_name")

	// KRM Status field comparison:
	// - Cert maps to .cert
	// - CertSerialNumber maps to .cert_serial_number
	// - CreateTime maps to .create_time
	// - ExpirationTime maps to .expiration_time
	// - Sha1Fingerprint maps to .sha1_fingerprint
	// - PrivateKey is returned only on creation in SslCertDetail.CertPrivateKey, not in SslCert proto
	// - ServerCaCert is CA cert of server returned on creation, not in SslCert proto
	f.StatusField(".cert")
	f.StatusField(".cert_serial_number")
	f.StatusField(".create_time")
	f.StatusField(".expiration_time")
	f.StatusField(".sha1_fingerprint")

	// Identity / Path fields
	f.Unimplemented_Identity(".instance")
	f.Unimplemented_Identity(".self_link")

	// Unimplemented / Not Yet Triaged fields
	f.Unimplemented_NotYetTriaged(".kind")

	return f
}
