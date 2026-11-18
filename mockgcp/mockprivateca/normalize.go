// Copyright 2022 Google LLC
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

package mockprivateca

import (
	"encoding/base64"
	"fmt"
	"strings"

	"github.com/GoogleCloudPlatform/k8s-config-connector/mockgcp/mockgcpregistry"
)

var _ mockgcpregistry.SupportsNormalization = &MockService{}

func (s *MockService) ConfigureVisitor(url string, visitor mockgcpregistry.NormalizingVisitor) {
	if !strings.Contains(url, "privateca.googleapis.com") {
		return
	}

	visitor.ReplacePath(".createTime", mockgcpregistry.PlaceholderTimestamp)
	visitor.ReplacePath(".updateTime", mockgcpregistry.PlaceholderTimestamp)
	visitor.ReplacePath(".deleteTime", mockgcpregistry.PlaceholderTimestamp)
	visitor.ReplacePath(".expireTime", mockgcpregistry.PlaceholderTimestamp)
	visitor.ReplacePath(".endTime", mockgcpregistry.PlaceholderTimestamp)

	visitor.ReplacePath(".response.createTime", mockgcpregistry.PlaceholderTimestamp)
	visitor.ReplacePath(".response.updateTime", mockgcpregistry.PlaceholderTimestamp)
	visitor.ReplacePath(".response.deleteTime", mockgcpregistry.PlaceholderTimestamp)
	visitor.ReplacePath(".response.expireTime", mockgcpregistry.PlaceholderTimestamp)

	visitor.ReplacePath(".metadata.createTime", mockgcpregistry.PlaceholderTimestamp)
	visitor.ReplacePath(".metadata.endTime", mockgcpregistry.PlaceholderTimestamp)
	visitor.RemovePath(".metadata.requestedCancellation")
	visitor.ReplacePath(".done", true)

	// SHA-1 hashes (keyId) should be 40 hex characters
	const keyId = "0123456789abcdef0123456789abcdef01234567"
	// SHA-256 hashes (fingerprint, digest) should be 64 hex characters
	const sha256 = "0123456789abcdef0123456789abcdef0123456789abcdef0123456789abcdef"

	visitor.ReplacePath(".caCertificateDescriptions[].subjectDescription.notBeforeTime", mockgcpregistry.PlaceholderTimestamp)
	visitor.ReplacePath(".caCertificateDescriptions[].subjectDescription.notAfterTime", mockgcpregistry.PlaceholderTimestamp)
	visitor.ReplacePath(".caCertificateDescriptions[].subjectDescription.hexSerialNumber", "0123456789abcdef")
	visitor.ReplacePath(".caCertificateDescriptions[].authorityKeyId.keyId", keyId)
	visitor.ReplacePath(".caCertificateDescriptions[].subjectKeyId.keyId", keyId)

	visitor.ReplacePath(".caCertificateDescriptions[].certFingerprint.sha256Hash", sha256)
	visitor.ReplacePath(".caCertificateDescriptions[].tbsCertificateDigest", sha256)
	visitor.RemovePath(".caCertificateDescriptions[].publicKey.format")
	visitor.ReplacePath(".accessUrls.caCertificateAccessUrl", "http://privateca-content-00000000-0000-0000-0000-000000000000.storage.googleapis.com/ca.crt")
	visitor.ReplacePath(".accessUrls.crlAccessUrls[]", "http://privateca-content-00000000-0000-0000-0000-000000000000.storage.googleapis.com/crl")

	visitor.ReplacePath(".status.caCertificateDescriptions[].subjectDescription.notBeforeTime", mockgcpregistry.PlaceholderTimestamp)
	visitor.ReplacePath(".status.caCertificateDescriptions[].subjectDescription.notAfterTime", mockgcpregistry.PlaceholderTimestamp)
	visitor.ReplacePath(".status.caCertificateDescriptions[].subjectDescription.hexSerialNumber", "0123456789abcdef")
	visitor.ReplacePath(".status.caCertificateDescriptions[].authorityKeyId.keyId", keyId)
	visitor.ReplacePath(".status.caCertificateDescriptions[].subjectKeyId.keyId", keyId)
	visitor.ReplacePath(".status.caCertificateDescriptions[].certFingerprint.sha256Hash", sha256)
	visitor.ReplacePath(".status.caCertificateDescriptions[].tbsCertificateDigest", sha256)
	visitor.RemovePath(".status.caCertificateDescriptions[].publicKey.format")
	visitor.ReplacePath(".status.accessUrls.caCertificateAccessUrl", "http://privateca-content-00000000-0000-0000-0000-000000000000.storage.googleapis.com/ca.crt")
	visitor.ReplacePath(".status.accessUrls.crlAccessUrls[]", "http://privateca-content-00000000-0000-0000-0000-000000000000.storage.googleapis.com/crl")

	const base64PublicKey = "LS0tLS1CRUdJTiBQVUJMSUMgS0VZLS0tLS0KTUlJQ0lqQU5CZ2txaGtpRzl3MEJBUUVGQUFPQ0FnOEFNSUlDQ2dLQ0FnRUFyeGt3dVBoREZTNlc1eEgvMW1MVApTYXhzMTNONnpGYlRXSUY4aG4vTlk1cXlJYmFpd0FYdEVOeU53NkhSSmd4R2c0WElkRXlMVGpCK1VNVVVLYU9OCnd0WFJTUW9CeGR2VitKeWRlL05jTUUzM3QyT3d1UDBRVzY2UnRLak52b2R5dzRLTHphVmp6T1hPY0YwV0NNYy8KRjV4cU5uUHpOalVncUlBanozSHkrejROWmNnT0lnM3dVdEJRRlNTUm16TmtzdjMwejhLQjdXUXJkaElWb3JuOQpuWVFPeW11eTZPT0dLM3FIUXRZYk1MYU9oREY1VnJ3amozblUxeStQMzdRR1kwdG5KS3VYbGNwR3hJN0tkWTI5CjRJM1F6K0JHemI0Wll0WE1uTzNOZFZVaTRteG14VTBMbVE3VlJkdHpkTGN6cTJDUEoyV2JjTTZkUldmVERoNisKYzdDQys3K1ZBdzlxeW1OSnFXN0wxb2JNSkNuTHpwcHBPVWg0RjNXU1V0SXVLcUZ2alo1eFMzUXBFSGJsRFoybgpUaTY1Tm4yR2JzeVYrc2FxTkpOdUdmVEpvUzRJOGFoakljY2hDaXUzRDMxNm5MczlENmMwckRLNmxsbUpHdFRLCmZsSVEyQkZmY2FtV2VlSXlNU1dIK3Uza2lKTTY2YWF1NVJrNlFXWWlKMTRhaWswNmsvK1pMRnNMazUvbVV3eVEKTWFUMXNPUGQ1Z3FSeUsrdGgyVXNkb1p1dE5PZFZYMWdRNjk4cXdVZk1oTmpkSzNLZUkzNWQvY2xuR2F1UGVPSgo2ZlgySy9WN1hTblQrcGRjRTExZjNFU0FZVEIybnJJSXgzK3NjYXdZalREd0Qrd2JQZ0p4ZG4wc1ppOTNtV2FKCkFPdTV4QjBOSStsYXhZT2tPZHhrYklVQ0F3RUFBUT09Ci0tLS0tRU5EIFBVQkxJQyBLRVktLS0tLQo="
	decodedKey, err := base64.StdEncoding.DecodeString(base64PublicKey)
	if err != nil {
		panic(fmt.Sprintf("decoding placeholder public key: %v", err))
	}
	visitor.ReplacePath(".caCertificateDescriptions[].publicKey.key", base64PublicKey)
	visitor.ReplacePath(".response.caCertificateDescriptions[].publicKey.key", base64PublicKey)
	// Direct controller stores decoded PEM in status while legacy TF controller stores base64.
	visitor.TransformString(".status.caCertificateDescriptions[].publicKey.key", func(s string) string {
		if strings.HasPrefix(s, "-----BEGIN PUBLIC KEY-----") {
			return string(decodedKey)
		}
		return base64PublicKey
	})

	const placeholderPEMCert = "-----BEGIN CERTIFICATE-----\n-----END CERTIFICATE-----\n"
	visitor.ReplacePath(".pemCaCertificates[]", placeholderPEMCert)
	visitor.ReplacePath(".response.pemCaCertificates[]", placeholderPEMCert)
	visitor.ReplacePath(".status.pemCaCertificates[]", placeholderPEMCert)

	visitor.ReplacePath(".response.caCertificateDescriptions[].subjectDescription.notBeforeTime", mockgcpregistry.PlaceholderTimestamp)
	visitor.ReplacePath(".response.caCertificateDescriptions[].subjectDescription.notAfterTime", mockgcpregistry.PlaceholderTimestamp)
	visitor.ReplacePath(".response.caCertificateDescriptions[].subjectDescription.hexSerialNumber", "0123456789abcdef")
	visitor.ReplacePath(".response.caCertificateDescriptions[].authorityKeyId.keyId", keyId)
	visitor.ReplacePath(".response.caCertificateDescriptions[].subjectKeyId.keyId", keyId)
	visitor.ReplacePath(".response.caCertificateDescriptions[].certFingerprint.sha256Hash", sha256)
	visitor.ReplacePath(".response.caCertificateDescriptions[].tbsCertificateDigest", sha256)
	visitor.RemovePath(".response.caCertificateDescriptions[].publicKey.format")
	visitor.ReplacePath(".response.accessUrls.caCertificateAccessUrl", "http://privateca-content-00000000-0000-0000-0000-000000000000.storage.googleapis.com/ca.crt")
	visitor.ReplacePath(".response.accessUrls.crlAccessUrls[]", "http://privateca-content-00000000-0000-0000-0000-000000000000.storage.googleapis.com/crl")
}

func (s *MockService) Previsit(event mockgcpregistry.Event, replacements mockgcpregistry.NormalizingVisitor) {
}
