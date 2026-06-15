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
	visitor.ReplacePath(".metadata.requestedCancellation", false)
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
