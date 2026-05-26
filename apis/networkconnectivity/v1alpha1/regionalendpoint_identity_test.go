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

package v1alpha1_test

import (
	"testing"

	"github.com/GoogleCloudPlatform/k8s-config-connector/apis/networkconnectivity/v1alpha1"
)

func TestNetworkConnectivityRegionalEndpointIdentity(t *testing.T) {
	id := &v1alpha1.NetworkConnectivityRegionalEndpointIdentity{
		Project:          "my-project",
		Location:         "us-central1",
		RegionalEndpoint: "my-endpoint",
	}

	expected := "projects/my-project/locations/us-central1/regionalEndpoints/my-endpoint"
	if actual := id.String(); actual != expected {
		t.Errorf("expected %q, got %q", expected, actual)
	}

	parsed := &v1alpha1.NetworkConnectivityRegionalEndpointIdentity{}
	if err := parsed.FromExternal(expected); err != nil {
		t.Fatalf("unexpected error parsing external: %v", err)
	}

	if *parsed != *id {
		t.Errorf("parsed %+v != original %+v", parsed, id)
	}
}
