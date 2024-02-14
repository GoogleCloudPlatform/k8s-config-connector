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

package cert_test

import (
	"testing"

	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/webhook/cert/generator"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/webhook/cert/writer"

	"k8s.io/client-go/util/cert"
)

const CommonName = "foo.example.com"

// TestDoesCertificateWorkWithKubernetes verifies that certificates
// that are created without a SAN are considered invalid, and
// thereby triggering overwriting the certificate with one that does
// have one.
//
// This verifies a fix for https://github.com/GoogleCloudPlatform/k8s-config-connector/issues/406.
func TestDoesCertificateWorkWithKubernetes(t *testing.T) {
	certConfig := cert.Config{CommonName: CommonName}
	privateKey, err := generator.NewPrivateKey()
	if err != nil {
		t.Fatalf("error when creating a private key: %s", err)
	}
	cert, err := generator.NewSelfSignedCACert(certConfig, privateKey)
	if err != nil {
		t.Fatalf("error when creating a new-style certificate: %s", err)
	}
	if !writer.DoesCertificateWorkWithK8sAPIClient(cert) {
		t.Fatalf("writer detected a new-style certificate is invalid. This implies" +
			"that newly generated certificates may not be compatible with all versions" +
			"of the Kubernetes HTTP client")
	}
	cert.DNSNames = nil
	if writer.DoesCertificateWorkWithK8sAPIClient(cert) {
		t.Fatalf("DoesCertificateWorkWithK8sAPIClient recognized an invalid cert as valid.")
	}
}
