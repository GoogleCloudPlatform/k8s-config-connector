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
	"reflect"
	"testing"
	"time"

	pb "github.com/GoogleCloudPlatform/k8s-config-connector/mockgcp/generated/mockgcp/cloud/compute/v1"
)

func TestPopulateSSLCertificateDefaults(t *testing.T) {
	s := &MockService{}
	fixedNow := time.Date(2026, 10, 30, 12, 0, 0, 0, time.UTC)
	expectedCreationTimestamp := s.formatDate(fixedNow)
	expectedActiveExpireTime := s.formatDate(fixedNow.Add(90 * 24 * time.Hour))
	expectedRenewalFailedExpireTime := s.formatDate(fixedNow.Add(30 * 24 * time.Hour))

	tests := []struct {
		name                   string
		input                  *pb.SslCertificate
		wantType               string
		wantStatus             string
		wantDomains            []string
		wantDomainStatus       map[string]string
		wantCertPresent        bool
		wantCertValue          *string
		wantExpireTime         *string
		wantSubjectAltNames    []string
		wantSelfManagedPresent bool
	}{
		{
			name: "self-managed certificate defaulted",
			input: &pb.SslCertificate{
				Certificate: PtrTo("my-cert-content"),
			},
			wantType:               "SELF_MANAGED",
			wantCertPresent:        true,
			wantCertValue:          PtrTo("my-cert-content"),
			wantExpireTime:         PtrTo(expectedCreationTimestamp),
			wantSelfManagedPresent: true,
		},
		{
			name: "managed certificate implicit type and default PROVISIONING status",
			input: &pb.SslCertificate{
				Managed: &pb.SslCertificateManagedSslCertificate{
					Domains: []string{"example.com.", "sub.example.com"},
				},
			},
			wantType:    "MANAGED",
			wantStatus:  "PROVISIONING",
			wantDomains: []string{"example.com", "sub.example.com"},
			wantDomainStatus: map[string]string{
				"example.com":     "PROVISIONING",
				"sub.example.com": "PROVISIONING",
			},
			wantCertPresent:     false,
			wantExpireTime:      nil,
			wantSubjectAltNames: nil,
		},
		{
			name: "managed certificate explicit PROVISIONING with custom domainStatus",
			input: &pb.SslCertificate{
				Type: PtrTo("MANAGED"),
				Managed: &pb.SslCertificateManagedSslCertificate{
					Status:  PtrTo("PROVISIONING"),
					Domains: []string{"example.com", "fail.com."},
					DomainStatus: map[string]string{
						"fail.com": "FAILED_NOT_VISIBLE",
					},
				},
			},
			wantType:    "MANAGED",
			wantStatus:  "PROVISIONING",
			wantDomains: []string{"example.com", "fail.com"},
			wantDomainStatus: map[string]string{
				"example.com": "PROVISIONING",
				"fail.com":    "FAILED_NOT_VISIBLE",
			},
			wantCertPresent:     false,
			wantExpireTime:      nil,
			wantSubjectAltNames: nil,
		},
		{
			name: "managed certificate ACTIVE status",
			input: &pb.SslCertificate{
				Type: PtrTo("MANAGED"),
				Managed: &pb.SslCertificateManagedSslCertificate{
					Status:  PtrTo("ACTIVE"),
					Domains: []string{"domain1.com.", "domain2.com"},
				},
			},
			wantType:    "MANAGED",
			wantStatus:  "ACTIVE",
			wantDomains: []string{"domain1.com", "domain2.com"},
			wantDomainStatus: map[string]string{
				"domain1.com": "ACTIVE",
				"domain2.com": "ACTIVE",
			},
			wantCertPresent:     true,
			wantCertValue:       PtrTo(dummyManagedSSLCertificate),
			wantExpireTime:      PtrTo(expectedActiveExpireTime),
			wantSubjectAltNames: []string{"domain1.com", "domain2.com"},
		},
		{
			name: "managed certificate PROVISIONING_FAILED status",
			input: &pb.SslCertificate{
				Type: PtrTo("MANAGED"),
				Managed: &pb.SslCertificateManagedSslCertificate{
					Status:  PtrTo("PROVISIONING_FAILED"),
					Domains: []string{"failed.com"},
				},
			},
			wantType:    "MANAGED",
			wantStatus:  "PROVISIONING_FAILED",
			wantDomains: []string{"failed.com"},
			wantDomainStatus: map[string]string{
				"failed.com": "FAILED_NOT_VISIBLE",
			},
			wantCertPresent:     false,
			wantExpireTime:      nil,
			wantSubjectAltNames: nil,
		},
		{
			name: "managed certificate PROVISIONING_FAILED_PERMANENTLY status",
			input: &pb.SslCertificate{
				Type: PtrTo("MANAGED"),
				Managed: &pb.SslCertificateManagedSslCertificate{
					Status:  PtrTo("PROVISIONING_FAILED_PERMANENTLY"),
					Domains: []string{"perm-failed.com"},
				},
			},
			wantType:    "MANAGED",
			wantStatus:  "PROVISIONING_FAILED_PERMANENTLY",
			wantDomains: []string{"perm-failed.com"},
			wantDomainStatus: map[string]string{
				"perm-failed.com": "FAILED_NOT_VISIBLE",
			},
			wantCertPresent:     false,
			wantExpireTime:      nil,
			wantSubjectAltNames: nil,
		},
		{
			name: "managed certificate RENEWAL_FAILED status",
			input: &pb.SslCertificate{
				Type: PtrTo("MANAGED"),
				Managed: &pb.SslCertificateManagedSslCertificate{
					Status:  PtrTo("RENEWAL_FAILED"),
					Domains: []string{"active-domain.com", "failed-domain.com"},
					DomainStatus: map[string]string{
						"active-domain.com": "ACTIVE",
					},
				},
			},
			wantType:    "MANAGED",
			wantStatus:  "RENEWAL_FAILED",
			wantDomains: []string{"active-domain.com", "failed-domain.com"},
			wantDomainStatus: map[string]string{
				"active-domain.com": "ACTIVE",
				"failed-domain.com": "FAILED_NOT_VISIBLE",
			},
			wantCertPresent:     true,
			wantCertValue:       PtrTo(dummyManagedSSLCertificate),
			wantExpireTime:      PtrTo(expectedRenewalFailedExpireTime),
			wantSubjectAltNames: []string{"active-domain.com", "failed-domain.com"},
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			obj := tc.input
			s.populateSSLCertificateDefaults(fixedNow, obj)

			if obj.GetType() != tc.wantType {
				t.Errorf("Type = %q, want %q", obj.GetType(), tc.wantType)
			}

			if tc.wantType == "SELF_MANAGED" {
				if tc.wantSelfManagedPresent && obj.SelfManaged == nil {
					t.Errorf("SelfManaged is nil, want populated")
				}
				if tc.wantCertPresent && obj.Certificate == nil {
					t.Errorf("Certificate is nil, want populated")
				}
				if tc.wantExpireTime != nil && (obj.ExpireTime == nil || *obj.ExpireTime != *tc.wantExpireTime) {
					t.Errorf("ExpireTime = %v, want %v", obj.ExpireTime, tc.wantExpireTime)
				}
				return
			}

			if obj.Managed == nil {
				t.Fatalf("Managed is nil")
			}

			if obj.Managed.GetStatus() != tc.wantStatus {
				t.Errorf("Managed.Status = %q, want %q", obj.Managed.GetStatus(), tc.wantStatus)
			}

			if !reflect.DeepEqual(obj.Managed.Domains, tc.wantDomains) {
				t.Errorf("Managed.Domains = %v, want %v", obj.Managed.Domains, tc.wantDomains)
			}

			if !reflect.DeepEqual(obj.Managed.DomainStatus, tc.wantDomainStatus) {
				t.Errorf("Managed.DomainStatus = %v, want %v", obj.Managed.DomainStatus, tc.wantDomainStatus)
			}

			if tc.wantCertPresent {
				if obj.Certificate == nil {
					t.Errorf("Certificate is nil, want present")
				} else if tc.wantCertValue != nil && *obj.Certificate != *tc.wantCertValue {
					t.Errorf("Certificate = %q, want %q", *obj.Certificate, *tc.wantCertValue)
				}
			} else {
				if obj.Certificate != nil {
					t.Errorf("Certificate = %q, want nil", *obj.Certificate)
				}
			}

			if tc.wantExpireTime != nil {
				if obj.ExpireTime == nil {
					t.Errorf("ExpireTime is nil, want %v", *tc.wantExpireTime)
				} else if *obj.ExpireTime != *tc.wantExpireTime {
					t.Errorf("ExpireTime = %q, want %q", *obj.ExpireTime, *tc.wantExpireTime)
				}
			} else {
				if obj.ExpireTime != nil {
					t.Errorf("ExpireTime = %q, want nil", *obj.ExpireTime)
				}
			}

			if !reflect.DeepEqual(obj.SubjectAlternativeNames, tc.wantSubjectAltNames) {
				t.Errorf("SubjectAlternativeNames = %v, want %v", obj.SubjectAlternativeNames, tc.wantSubjectAltNames)
			}
		})
	}
}
