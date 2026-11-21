// Copyright 2026 Google LLC
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     https://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package iap

import (
	"testing"

	pb "cloud.google.com/go/iap/apiv1/iappb"
	"github.com/google/go-cmp/cmp"
	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/types/known/wrapperspb"
)

func TestCalculateUpdateMaskOAuthSettingsFromJSON(t *testing.T) {
	for _, tt := range []struct {
		name    string
		clients []string
		want    []string
	}{
		{
			name:    "unchanged decoded settings need no update",
			clients: []string{"client-a"},
		},
		{
			name:    "client update does not select unchanged decoded login hint",
			clients: []string{"client-b"},
			want:    []string{"iap_settings.access_settings.oauth_settings.programmatic_clients"},
		},
	} {
		t.Run(tt.name, func(t *testing.T) {
			actual := &pb.IapSettings{}
			if err := protojson.Unmarshal([]byte(`{
				"accessSettings": {
					"oauthSettings": {
						"loginHint": "example.com",
						"programmaticClients": ["client-a"]
					}
				}
			}`), actual); err != nil {
				t.Fatal(err)
			}
			desired := &pb.IapSettings{
				AccessSettings: &pb.AccessSettings{
					OauthSettings: &pb.OAuthSettings{
						LoginHint:           wrapperspb.String("example.com"),
						ProgrammaticClients: tt.clients,
					},
				},
			}
			if diff := cmp.Diff(tt.want, calculateUpdateMask(desired, actual)); diff != "" {
				t.Errorf("unexpected update mask (-want +got):\n%s", diff)
			}
		})
	}
}

func TestCalculateUpdateMaskOAuthSettings(t *testing.T) {
	const clientsPath = "iap_settings.access_settings.oauth_settings.programmatic_clients"
	const hintPath = "iap_settings.access_settings.oauth_settings.login_hint"
	settings := func(oauth *pb.OAuthSettings) *pb.IapSettings {
		return &pb.IapSettings{AccessSettings: &pb.AccessSettings{OauthSettings: oauth}}
	}
	tests := []struct {
		name    string
		desired *pb.IapSettings
		actual  *pb.IapSettings
		want    []string
	}{
		{
			name:    "initialize clients without access settings",
			desired: settings(&pb.OAuthSettings{ProgrammaticClients: []string{"client-a"}}),
			actual:  &pb.IapSettings{},
			want:    []string{clientsPath},
		},
		{
			name:    "initialize clients without oauth settings",
			desired: settings(&pb.OAuthSettings{ProgrammaticClients: []string{"client-a"}}),
			actual:  settings(nil),
			want:    []string{clientsPath},
		},
		{
			name:    "add client",
			desired: settings(&pb.OAuthSettings{ProgrammaticClients: []string{"client-a", "client-b"}}),
			actual:  settings(&pb.OAuthSettings{ProgrammaticClients: []string{"client-a"}}),
			want:    []string{clientsPath},
		},
		{
			name:    "client update excludes existing oauth credentials",
			desired: settings(&pb.OAuthSettings{ProgrammaticClients: []string{"client-b"}}),
			actual: settings(&pb.OAuthSettings{
				ProgrammaticClients: []string{"client-a"},
				ClientId:            wrapperspb.String("browser-client"),
				ClientSecretSha256:  wrapperspb.String("example-hash"),
			}),
			want: []string{clientsPath},
		},
		{
			name:    "unchanged clients ignore existing oauth credentials",
			desired: settings(&pb.OAuthSettings{ProgrammaticClients: []string{"client-a"}}),
			actual: settings(&pb.OAuthSettings{
				ProgrammaticClients: []string{"client-a"},
				ClientId:            wrapperspb.String("browser-client"),
				ClientSecretSha256:  wrapperspb.String("example-hash"),
			}),
		},
		{
			name:    "remove client",
			desired: settings(&pb.OAuthSettings{ProgrammaticClients: []string{"client-b"}}),
			actual:  settings(&pb.OAuthSettings{ProgrammaticClients: []string{"client-a", "client-b"}}),
			want:    []string{clientsPath},
		},
		{
			name:    "clear clients",
			desired: settings(&pb.OAuthSettings{}),
			actual:  settings(&pb.OAuthSettings{ProgrammaticClients: []string{"client-a"}}),
			want:    []string{clientsPath},
		},
		{
			name:    "initialize login hint",
			desired: settings(&pb.OAuthSettings{LoginHint: wrapperspb.String("example.com")}),
			actual:  &pb.IapSettings{},
			want:    []string{hintPath},
		},
		{
			name:    "change login hint without changing clients",
			desired: settings(&pb.OAuthSettings{LoginHint: wrapperspb.String("new.example.com"), ProgrammaticClients: []string{"client-a"}}),
			actual:  settings(&pb.OAuthSettings{LoginHint: wrapperspb.String("old.example.com"), ProgrammaticClients: []string{"client-a"}}),
			want:    []string{hintPath},
		},
		{
			name:    "clear login hint",
			desired: settings(&pb.OAuthSettings{}),
			actual:  settings(&pb.OAuthSettings{LoginHint: wrapperspb.String("example.com")}),
			want:    []string{hintPath},
		},
		{
			name:    "explicit empty login hint",
			desired: settings(&pb.OAuthSettings{LoginHint: wrapperspb.String("")}),
			actual:  &pb.IapSettings{},
			want:    []string{hintPath},
		},
		{
			name:    "change both fields",
			desired: settings(&pb.OAuthSettings{LoginHint: wrapperspb.String("example.com"), ProgrammaticClients: []string{"client-b"}}),
			actual:  settings(&pb.OAuthSettings{ProgrammaticClients: []string{"client-a"}}),
			want:    []string{hintPath, clientsPath},
		},
		{
			name:    "unchanged settings",
			desired: settings(&pb.OAuthSettings{LoginHint: wrapperspb.String("example.com"), ProgrammaticClients: []string{"client-a"}}),
			actual:  settings(&pb.OAuthSettings{LoginHint: wrapperspb.String("example.com"), ProgrammaticClients: []string{"client-a"}}),
		},
		{
			name:    "nil and empty clients are equivalent",
			desired: settings(&pb.OAuthSettings{ProgrammaticClients: []string{}}),
			actual:  settings(&pb.OAuthSettings{}),
		},
		{
			name:    "empty oauth settings need no initialization",
			desired: settings(&pb.OAuthSettings{}),
			actual:  &pb.IapSettings{},
		},
		{
			name:    "omitted oauth settings are unmanaged",
			desired: settings(nil),
			actual:  settings(&pb.OAuthSettings{LoginHint: wrapperspb.String("example.com"), ProgrammaticClients: []string{"client-a"}}),
		},
		{
			name:    "omitted access settings are unmanaged",
			desired: &pb.IapSettings{},
			actual:  settings(&pb.OAuthSettings{ProgrammaticClients: []string{"client-a"}}),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if diff := cmp.Diff(tt.want, calculateUpdateMask(tt.desired, tt.actual)); diff != "" {
				t.Errorf("unexpected update mask (-want +got):\n%s", diff)
			}
		})
	}
}
