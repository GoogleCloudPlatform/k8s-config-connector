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

package test

import (
	"testing"
)

func TestExtractResourceKey(t *testing.T) {
	tests := []struct {
		url  string
		want string
	}{
		{
			url:  "https://container.googleapis.com/v1beta1/projects/mock-project/locations/us-central1-a/clusters/cluster-sample-123",
			want: "container/clusters/cluster-sample-123",
		},
		{
			url:  "https://container.googleapis.com/v1beta1/projects/mock-project/locations/us-central1-a/clusters/cluster-sample-123/nodePools/nodepool-sample-123?alt=json",
			want: "container/nodePools/nodepool-sample-123",
		},
		{
			url:  "https://compute.googleapis.com/compute/v1/projects/mock-project/global/networks/computenetwork-123",
			want: "compute/networks/computenetwork-123",
		},
		{
			url:  "https://cloudkms.googleapis.com/v1/projects/mock-project/locations/us-central1/keyRings/keyring-123/cryptoKeys/cryptokey-123:getIamPolicy",
			want: "cloudkms/cryptoKeys/cryptokey-123",
		},
		{
			url:  "https://container.googleapis.com/v1beta1/projects/mock-project/locations/us-central1-a/operations/operation-123",
			want: "container/us-central1-a/operations/operation-123",
		},
		{
			url:  "https://compute.googleapis.com/compute/v1/projects/mock-project/regions/us-central1/operations/operation-456",
			want: "compute/us-central1/operations/operation-456",
		},
		{
			url:  "https://compute.googleapis.com/compute/v1/projects/mock-project/global/operations/operation-789",
			want: "compute/global/operations/operation-789",
		},
	}

	for _, tc := range tests {
		got := extractResourceKey(tc.url)
		if got != tc.want {
			t.Errorf("extractResourceKey(%q) = %q; want %q", tc.url, got, tc.want)
		}
	}
}

func TestCompareHTTPLogs(t *testing.T) {
	testCases := []struct {
		name    string
		wantLog string
		gotLog  string
		wantErr bool
	}{
		{
			name: "exact match",
			wantLog: `GET https://compute.googleapis.com/compute/v1/projects/mock-project/global/networks/net-123
200 OK

{
  "name": "net-123"
}
---
GET https://compute.googleapis.com/compute/v1/projects/mock-project/global/networks/net-456
200 OK

{
  "name": "net-456"
}`,
			gotLog: `GET https://compute.googleapis.com/compute/v1/projects/mock-project/global/networks/net-123
200 OK

{
  "name": "net-123"
}
---
GET https://compute.googleapis.com/compute/v1/projects/mock-project/global/networks/net-456
200 OK

{
  "name": "net-456"
}`,
			wantErr: false,
		},
		{
			name: "out of order resource calls match successfully",
			wantLog: `GET https://compute.googleapis.com/compute/v1/projects/mock-project/global/networks/net-123
200 OK

{
  "name": "net-123"
}
---
GET https://compute.googleapis.com/compute/v1/projects/mock-project/global/networks/net-456
200 OK

{
  "name": "net-456"
}`,
			gotLog: `GET https://compute.googleapis.com/compute/v1/projects/mock-project/global/networks/net-456
200 OK

{
  "name": "net-456"
}
---
GET https://compute.googleapis.com/compute/v1/projects/mock-project/global/networks/net-123
200 OK

{
  "name": "net-123"
}`,
			wantErr: false,
		},
		{
			name: "internal resource sequence order mismatch fails",
			wantLog: `GET https://compute.googleapis.com/compute/v1/projects/mock-project/global/networks/net-123
404 Not Found
---
POST https://compute.googleapis.com/compute/v1/projects/mock-project/global/networks/net-123
200 OK`,
			gotLog: `POST https://compute.googleapis.com/compute/v1/projects/mock-project/global/networks/net-123
200 OK
---
GET https://compute.googleapis.com/compute/v1/projects/mock-project/global/networks/net-123
404 Not Found`,
			wantErr: true,
		},
		{
			name: "payload JSON diff fails",
			wantLog: `POST https://compute.googleapis.com/compute/v1/projects/mock-project/global/networks/net-123
200 OK

{
  "name": "net-123",
  "routingMode": "GLOBAL"
}`,
			gotLog: `POST https://compute.googleapis.com/compute/v1/projects/mock-project/global/networks/net-123
200 OK

{
  "name": "net-123",
  "routingMode": "REGIONAL"
}`,
			wantErr: true,
		},
		{
			name: "volatile URL query parameters ignored",
			wantLog: `GET https://compute.googleapis.com/compute/v1/projects/mock-project/global/networks/net-123?alt=json&prettyPrint=false
200 OK`,
			gotLog: `GET https://compute.googleapis.com/compute/v1/projects/mock-project/global/networks/net-123?alt=json
200 OK`,
			wantErr: false,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			err := compareHTTPLogs(tc.wantLog, tc.gotLog)
			if (err != nil) != tc.wantErr {
				t.Errorf("compareHTTPLogs() error = %v, wantErr %v", err, tc.wantErr)
			}
		})
	}
}
