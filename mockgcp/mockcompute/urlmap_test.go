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
	"testing"

	pb "github.com/GoogleCloudPlatform/k8s-config-connector/mockgcp/generated/mockgcp/cloud/compute/v1"
)

func TestValidateURLMap(t *testing.T) {
	service := "https://www.googleapis.com/compute/v1/projects/p/global/backendServices/bs"
	redirect := &pb.HttpRedirectAction{
		HostRedirect: PtrTo("example.com"),
	}

	tests := []struct {
		name    string
		urlMap  *pb.UrlMap
		wantErr bool
	}{
		{
			name: "only defaultService",
			urlMap: &pb.UrlMap{
				DefaultService: &service,
			},
			wantErr: false,
		},
		{
			name: "only defaultUrlRedirect",
			urlMap: &pb.UrlMap{
				DefaultUrlRedirect: redirect,
			},
			wantErr: false,
		},
		{
			name: "both defaultService and defaultUrlRedirect",
			urlMap: &pb.UrlMap{
				DefaultService:     &service,
				DefaultUrlRedirect: redirect,
			},
			wantErr: true,
		},
		{
			name: "both defaultService and defaultRouteAction",
			urlMap: &pb.UrlMap{
				DefaultService: &service,
				DefaultRouteAction: &pb.HttpRouteAction{
					WeightedBackendServices: []*pb.WeightedBackendService{
						{BackendService: &service},
					},
				},
			},
			wantErr: true,
		},
		{
			name: "pathMatcher has both defaultService and defaultUrlRedirect",
			urlMap: &pb.UrlMap{
				DefaultService: &service,
				PathMatchers: []*pb.PathMatcher{
					{
						DefaultService:     &service,
						DefaultUrlRedirect: redirect,
					},
				},
			},
			wantErr: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := validateURLMap(tt.urlMap)
			if (err != nil) != tt.wantErr {
				t.Errorf("validateURLMap() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
