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

package certificatemanager

import (
	"context"
	"sort"
	"testing"

	pb "cloud.google.com/go/certificatemanager/apiv1/certificatemanagerpb"
	krm "github.com/GoogleCloudPlatform/k8s-config-connector/apis/certificatemanager/v1beta1"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct"
	"github.com/google/go-cmp/cmp"
)

func TestCompareCertificate(t *testing.T) {
	ctx := context.Background()

	tests := []struct {
		name      string
		actual    *pb.Certificate
		desired   *pb.Certificate
		wantPaths []string
	}{
		{
			name: "managed cert identical spec with server-computed fields in actual",
			actual: &pb.Certificate{
				Name:        "projects/test-project/locations/global/certificates/cert-1",
				Description: "sample managed cert for kcc",
				Labels: map[string]string{
					"cnrm-test": "true",
				},
				Type: &pb.Certificate_Managed{
					Managed: &pb.Certificate_ManagedCertificate{
						Domains: []string{
							"domain1.example.com",
							"domain2.example.com",
						},
						DnsAuthorizations: []string{
							"projects/test-project/locations/global/dnsAuthorizations/dns-1",
							"projects/test-project/locations/global/dnsAuthorizations/dns-2",
						},
						State: pb.Certificate_ManagedCertificate_ACTIVE,
						AuthorizationAttemptInfo: []*pb.Certificate_ManagedCertificate_AuthorizationAttemptInfo{
							{
								Domain: "domain1.example.com",
								State:  pb.Certificate_ManagedCertificate_AuthorizationAttemptInfo_AUTHORIZED,
							},
							{
								Domain: "domain2.example.com",
								State:  pb.Certificate_ManagedCertificate_AuthorizationAttemptInfo_AUTHORIZED,
							},
						},
					},
				},
				Scope: pb.Certificate_EDGE_CACHE,
			},
			desired: func() *pb.Certificate {
				mapCtx := &direct.MapContext{}
				desc := "sample managed cert for kcc"
				scope := "EDGE_CACHE"
				spec := &krm.CertificateManagerCertificateSpec{
					Description: &desc,
					Location:    "global",
					Scope:       &scope,
					Managed: &krm.CertificateManaged{
						Domains: []string{
							"domain1.example.com",
							"domain2.example.com",
						},
						DnsAuthorizationsRefs: []krm.CertificateManagerDNSAuthorizationRef{
							{External: "projects/test-project/locations/global/dnsAuthorizations/dns-1"},
							{External: "projects/test-project/locations/global/dnsAuthorizations/dns-2"},
						},
					},
				}
				d := CertificateManagerCertificateSpec_v1beta1_ToProto(mapCtx, spec)
				d.Labels = map[string]string{
					"cnrm-test": "true",
				}
				return d
			}(),
			wantPaths: []string{},
		},
		{
			name: "self_managed cert identical spec with secret pemPrivateKey omitted from actual",
			actual: &pb.Certificate{
				Name:        "projects/test-project/locations/europe-west1/certificates/cert-2",
				Description: "regional self-managed certificate",
				Labels: map[string]string{
					"cnrm-test": "true",
				},
				Type: &pb.Certificate_SelfManaged{
					SelfManaged: &pb.Certificate_SelfManagedCertificate{
						PemCertificate: "-----BEGIN CERTIFICATE-----\nMIID...cert\n-----END CERTIFICATE-----",
					},
				},
			},
			desired: func() *pb.Certificate {
				mapCtx := &direct.MapContext{}
				desc := "regional self-managed certificate"
				pemCert := "-----BEGIN CERTIFICATE-----\nMIID...cert\n-----END CERTIFICATE-----"
				pemKey := "-----BEGIN PRIVATE KEY-----\nMIIE...key\n-----END PRIVATE KEY-----"
				spec := &krm.CertificateManagerCertificateSpec{
					Description: &desc,
					Location:    "europe-west1",
					SelfManaged: &krm.CertificateSelfManaged{
						PemCertificate: &pemCert,
						PemPrivateKey: &krm.CertificatePemPrivateKey{
							Value: &pemKey,
						},
					},
				}
				d := CertificateManagerCertificateSpec_v1beta1_ToProto(mapCtx, spec)
				d.Labels = map[string]string{
					"cnrm-test": "true",
				}
				return d
			}(),
			wantPaths: []string{},
		},
		{
			name: "mutable description and labels updated",
			actual: &pb.Certificate{
				Name:        "projects/test-project/locations/global/certificates/cert-1",
				Description: "old description",
				Labels: map[string]string{
					"old-key": "old-val",
				},
				Type: &pb.Certificate_Managed{
					Managed: &pb.Certificate_ManagedCertificate{
						Domains: []string{"example.com"},
						State:   pb.Certificate_ManagedCertificate_ACTIVE,
					},
				},
			},
			desired: func() *pb.Certificate {
				mapCtx := &direct.MapContext{}
				desc := "new description"
				spec := &krm.CertificateManagerCertificateSpec{
					Description: &desc,
					Location:    "global",
					Managed: &krm.CertificateManaged{
						Domains: []string{"example.com"},
					},
				}
				d := CertificateManagerCertificateSpec_v1beta1_ToProto(mapCtx, spec)
				d.Labels = map[string]string{
					"new-key": "new-val",
				}
				return d
			}(),
			wantPaths: []string{"description", "labels"},
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			_, updateMask, err := compareCertificate(ctx, tc.actual, tc.desired)
			if err != nil {
				t.Fatalf("unexpected error: %v", err)
			}
			var gotPaths []string
			if updateMask != nil {
				gotPaths = append([]string{}, updateMask.GetPaths()...)
			}
			sort.Strings(gotPaths)
			wantPaths := append([]string{}, tc.wantPaths...)
			sort.Strings(wantPaths)
			if diff := cmp.Diff(wantPaths, gotPaths); diff != "" {
				t.Errorf("updateMask paths mismatch (-want +got):\n%s", diff)
			}
		})
	}
}
