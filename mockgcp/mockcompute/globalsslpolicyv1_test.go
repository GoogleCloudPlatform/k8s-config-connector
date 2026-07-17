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

	pb "github.com/GoogleCloudPlatform/k8s-config-connector/mockgcp/generated/mockgcp/cloud/compute/v1"
)

func TestPopulateSslPolicy(t *testing.T) {
	tests := []struct {
		name        string
		input       *pb.SslPolicy
		wantMinTls  string
		wantProfile string
		wantCiphers []string
		wantErr     bool
	}{
		{
			name:        "default values (COMPATIBLE)",
			input:       &pb.SslPolicy{},
			wantMinTls:  "TLS_1_0",
			wantProfile: "COMPATIBLE",
			wantCiphers: []string{
				"TLS_ECDHE_ECDSA_WITH_AES_128_CBC_SHA",
				"TLS_ECDHE_ECDSA_WITH_AES_128_GCM_SHA256",
				"TLS_ECDHE_ECDSA_WITH_AES_256_CBC_SHA",
				"TLS_ECDHE_ECDSA_WITH_AES_256_GCM_SHA384",
				"TLS_ECDHE_ECDSA_WITH_CHACHA20_POLY1305_SHA256",
				"TLS_ECDHE_RSA_WITH_AES_128_CBC_SHA",
				"TLS_ECDHE_RSA_WITH_AES_128_GCM_SHA256",
				"TLS_ECDHE_RSA_WITH_AES_256_CBC_SHA",
				"TLS_ECDHE_RSA_WITH_AES_256_GCM_SHA384",
				"TLS_ECDHE_RSA_WITH_CHACHA20_POLY1305_SHA256",
				"TLS_RSA_WITH_3DES_EDE_CBC_SHA",
				"TLS_RSA_WITH_AES_128_CBC_SHA",
				"TLS_RSA_WITH_AES_128_GCM_SHA256",
				"TLS_RSA_WITH_AES_256_CBC_SHA",
				"TLS_RSA_WITH_AES_256_GCM_SHA384",
			},
		},
		{
			name: "RESTRICTED profile",
			input: &pb.SslPolicy{
				Profile: PtrTo("RESTRICTED"),
			},
			wantMinTls:  "TLS_1_0",
			wantProfile: "RESTRICTED",
			wantCiphers: []string{
				"TLS_ECDHE_ECDSA_WITH_AES_128_GCM_SHA256",
				"TLS_ECDHE_ECDSA_WITH_AES_256_GCM_SHA384",
				"TLS_ECDHE_ECDSA_WITH_CHACHA20_POLY1305_SHA256",
				"TLS_ECDHE_RSA_WITH_AES_128_GCM_SHA256",
				"TLS_ECDHE_RSA_WITH_AES_256_GCM_SHA384",
				"TLS_ECDHE_RSA_WITH_CHACHA20_POLY1305_SHA256",
			},
		},
		{
			name: "MODERN profile",
			input: &pb.SslPolicy{
				Profile: PtrTo("MODERN"),
			},
			wantMinTls:  "TLS_1_0",
			wantProfile: "MODERN",
			wantCiphers: []string{
				"TLS_ECDHE_ECDSA_WITH_AES_128_CBC_SHA",
				"TLS_ECDHE_ECDSA_WITH_AES_128_GCM_SHA256",
				"TLS_ECDHE_ECDSA_WITH_AES_256_CBC_SHA",
				"TLS_ECDHE_ECDSA_WITH_AES_256_GCM_SHA384",
				"TLS_ECDHE_ECDSA_WITH_CHACHA20_POLY1305_SHA256",
				"TLS_ECDHE_RSA_WITH_AES_128_CBC_SHA",
				"TLS_ECDHE_RSA_WITH_AES_128_GCM_SHA256",
				"TLS_ECDHE_RSA_WITH_AES_256_CBC_SHA",
				"TLS_ECDHE_RSA_WITH_AES_256_GCM_SHA384",
				"TLS_ECDHE_RSA_WITH_CHACHA20_POLY1305_SHA256",
			},
		},
		{
			name: "FIPS_202205 profile with min TLS 1.2",
			input: &pb.SslPolicy{
				Profile:       PtrTo("FIPS_202205"),
				MinTlsVersion: PtrTo("TLS_1_2"),
			},
			wantMinTls:  "TLS_1_2",
			wantProfile: "FIPS_202205",
			wantCiphers: []string{
				"TLS_ECDHE_ECDSA_WITH_AES_128_GCM_SHA256",
				"TLS_ECDHE_ECDSA_WITH_AES_256_GCM_SHA384",
				"TLS_ECDHE_RSA_WITH_AES_128_GCM_SHA256",
				"TLS_ECDHE_RSA_WITH_AES_256_GCM_SHA384",
			},
		},
		{
			name: "CUSTOM profile with ciphers",
			input: &pb.SslPolicy{
				Profile: PtrTo("CUSTOM"),
				CustomFeatures: []string{
					"TLS_ECDHE_ECDSA_WITH_AES_256_GCM_SHA384",
					"TLS_ECDHE_RSA_WITH_AES_256_GCM_SHA384",
				},
			},
			wantMinTls:  "TLS_1_0",
			wantProfile: "CUSTOM",
			wantCiphers: []string{
				"TLS_ECDHE_ECDSA_WITH_AES_256_GCM_SHA384",
				"TLS_ECDHE_RSA_WITH_AES_256_GCM_SHA384",
			},
		},
		{
			name: "Invalid FIPS minTlsVersion",
			input: &pb.SslPolicy{
				Profile:       PtrTo("FIPS_202205"),
				MinTlsVersion: PtrTo("TLS_1_0"),
			},
			wantErr: true,
		},
		{
			name: "Invalid TLS 1.3 profile",
			input: &pb.SslPolicy{
				Profile:       PtrTo("COMPATIBLE"),
				MinTlsVersion: PtrTo("TLS_1_3"),
			},
			wantErr: true,
		},
		{
			name: "CUSTOM profile empty ciphers",
			input: &pb.SslPolicy{
				Profile: PtrTo("CUSTOM"),
			},
			wantErr: true,
		},
		{
			name: "Non-CUSTOM profile with custom ciphers",
			input: &pb.SslPolicy{
				Profile: PtrTo("MODERN"),
				CustomFeatures: []string{
					"TLS_ECDHE_ECDSA_WITH_AES_256_GCM_SHA384",
				},
			},
			wantErr: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := populateSslPolicy(tt.input)
			if (err != nil) != tt.wantErr {
				t.Fatalf("populateSslPolicy() error = %v, wantErr = %v", err, tt.wantErr)
			}
			if tt.wantErr {
				return
			}
			if *tt.input.MinTlsVersion != tt.wantMinTls {
				t.Errorf("MinTlsVersion = %q, want %q", *tt.input.MinTlsVersion, tt.wantMinTls)
			}
			if *tt.input.Profile != tt.wantProfile {
				t.Errorf("Profile = %q, want %q", *tt.input.Profile, tt.wantProfile)
			}
			if !reflect.DeepEqual(tt.input.EnabledFeatures, tt.wantCiphers) {
				t.Errorf("EnabledFeatures =\n%v\nwant\n%v", tt.input.EnabledFeatures, tt.wantCiphers)
			}
		})
	}
}
