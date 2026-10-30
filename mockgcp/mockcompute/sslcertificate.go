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

package mockcompute

import (
	"strings"
	"time"

	pb "github.com/GoogleCloudPlatform/k8s-config-connector/mockgcp/generated/mockgcp/cloud/compute/v1"
)

const dummyManagedSSLCertificate = "-----BEGIN CERTIFICATE-----\n" +
	"MIIDJTCCAg0CFHdD3ZGYMCmF3O4PvMwsP5i8d/V0MA0GCSqGSIb3DQEBCwUAME8x\n" +
	"CzAJBgNVBAYTAlVTMQswCQYDVQQIDAJXQTEhMB8GA1UECgwYSW50ZXJuZXQgV2lk\n" +
	"Z2l0cyBQdHkgTHRkMRAwDgYDVQQDDAdFeGFtcGxlMB4XDTE5MDkyOTIyMjgyOVoX\n" +
	"DTIwMDkyODIyMjgyOVowTzELMAkGA1UEBhMCVVMxCzAJBgNVBAgMAldBMSEwHwYD\n" +
	"VQQKDBhJbnRlcm5ldCBXaWRnaXRzIFB0eSBMdGQxEDAOBgNVBAMMB0V4YW1wbGUw\n" +
	"ggEiMA0GCSqGSIb3DQEBAQUAA4IBDwAwggEKAoIBAQDWLvOZIail12i6NXIqOspV\n" +
	"corkuS1Nl0ayrl0VuKHCvheun/s7lLLgEfifzRueYlSUtdGg4atWIwEKsbIE+AF9\n" +
	"uUTzkq/t6zHxFAAWgVZ6/hW696jqcZX3yU+LCuHPLSN0ruqD6ZygnYDVciDmYwxe\n" +
	"601xNfOOYRlm6dGRx6uTxGDZtfu8zsaNI0UxTugTp2x5cKB66SbgdlIJvc2Hb54a\n" +
	"7qOsb9CIf+rrK2xUdJUj4ueUEIMxjnY2u/Dc71SgfBVn+yFfN9MHNdcTWPXEUClE\n" +
	"Fxd/MB3dGn7hVavXyvy3NT4tWhBgYBphfEUudDFej5MmVq56JOEQ2UtaQ+Imscud\n" +
	"AgMBAAEwDQYJKoZIhvcNAQELBQADggEBAMYTQyjVlo6TCYoyK6akjPX7vRiwCCAh\n" +
	"jqsEu3bZqwUreOhZgRAyEXrq68dtXwTbwdisQmnhpBeBQuX4WWeas9TiycZ13TA1\n" +
	"Z+h518D9OVXjrNs7oE3QNFeTom807IW16YydlrZMLKO8mQg6/BXfSHbLwuQHSIYS\n" +
	"JD+uOfnkr08ORBbLGgBKKpy7ngflIkdSrQPmCYmYlvoy+goMAEVi0K3Y1wVzAF4k\n" +
	"O4v8f7GXkNarsFT1QM82JboVV5uwX+uDmi858WKDHYGv2Ypv6yy93vdV0Xt/IBj3\n" +
	"95/RDisBzcL7Ynpl34AAr5MLm7yCSsPrAmgevX4BOtcVc4rSXj5rcoE=\n" +
	"-----END CERTIFICATE-----\n"

func (s *MockService) populateSSLCertificateDefaults(now time.Time, obj *pb.SslCertificate) {
	if obj.CreationTimestamp == nil {
		obj.CreationTimestamp = PtrTo(s.formatDate(now))
	} else if t, err := time.Parse(time.RFC3339Nano, obj.GetCreationTimestamp()); err == nil {
		now = t
	}

	if obj.Type == nil {
		if obj.Managed != nil {
			obj.Type = PtrTo("MANAGED")
		} else {
			obj.Type = PtrTo("SELF_MANAGED")
		}
	}

	if obj.GetType() == "SELF_MANAGED" {
		if obj.SelfManaged == nil {
			obj.SelfManaged = &pb.SslCertificateSelfManagedSslCertificate{}
		}
		if obj.SelfManaged.Certificate == nil {
			obj.SelfManaged.Certificate = obj.Certificate
		}
		if obj.ExpireTime == nil {
			obj.ExpireTime = obj.CreationTimestamp
		}
		return
	}

	if obj.GetType() == "MANAGED" {
		if obj.Managed == nil {
			obj.Managed = &pb.SslCertificateManagedSslCertificate{}
		}
		if obj.Managed.Status == nil || *obj.Managed.Status == "" {
			obj.Managed.Status = PtrTo("PROVISIONING")
		}

		var normalizedDomains []string
		for _, d := range obj.Managed.Domains {
			normalizedDomains = append(normalizedDomains, strings.TrimSuffix(d, "."))
		}
		obj.Managed.Domains = normalizedDomains

		normalizedDomainStatus := make(map[string]string)
		for k, v := range obj.Managed.DomainStatus {
			normalizedDomainStatus[strings.TrimSuffix(k, ".")] = v
		}

		switch obj.Managed.GetStatus() {
		case "ACTIVE":
			// Issued: Yes
			if obj.Certificate == nil {
				obj.Certificate = PtrTo(dummyManagedSSLCertificate)
			}
			if obj.ExpireTime == nil {
				obj.ExpireTime = PtrTo(s.formatDate(now.Add(90 * 24 * time.Hour)))
			}
			if len(obj.SubjectAlternativeNames) == 0 {
				obj.SubjectAlternativeNames = normalizedDomains
			}
			for _, d := range normalizedDomains {
				if _, ok := normalizedDomainStatus[d]; !ok {
					normalizedDomainStatus[d] = "ACTIVE"
				}
			}

		case "PROVISIONING":
			// Issued: No
			obj.Certificate = nil
			obj.ExpireTime = nil
			obj.SubjectAlternativeNames = nil
			for _, d := range normalizedDomains {
				if _, ok := normalizedDomainStatus[d]; !ok {
					normalizedDomainStatus[d] = "PROVISIONING"
				}
			}

		case "PROVISIONING_FAILED", "PROVISIONING_FAILED_PERMANENTLY":
			// Issued: No
			obj.Certificate = nil
			obj.ExpireTime = nil
			obj.SubjectAlternativeNames = nil
			for _, d := range normalizedDomains {
				if _, ok := normalizedDomainStatus[d]; !ok {
					normalizedDomainStatus[d] = "FAILED_NOT_VISIBLE"
				}
			}

		case "RENEWAL_FAILED":
			// Issued: Yes (old cert)
			if obj.Certificate == nil {
				obj.Certificate = PtrTo(dummyManagedSSLCertificate)
			}
			if obj.ExpireTime == nil {
				obj.ExpireTime = PtrTo(s.formatDate(now.Add(30 * 24 * time.Hour)))
			}
			if len(obj.SubjectAlternativeNames) == 0 {
				obj.SubjectAlternativeNames = normalizedDomains
			}
			for _, d := range normalizedDomains {
				if _, ok := normalizedDomainStatus[d]; !ok {
					normalizedDomainStatus[d] = "FAILED_NOT_VISIBLE"
				}
			}

		default:
			for _, d := range normalizedDomains {
				if _, ok := normalizedDomainStatus[d]; !ok {
					normalizedDomainStatus[d] = obj.Managed.GetStatus()
				}
			}
		}

		obj.Managed.DomainStatus = normalizedDomainStatus
	}
}
