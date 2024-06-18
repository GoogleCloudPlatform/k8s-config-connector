/*
Copyright 2024.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

	http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/
package externalresource

import (
	"testing"

	"cloud.google.com/go/cloudbuild/apiv1/v2/cloudbuildpb"
	"cloud.google.com/go/compute/apiv1/computepb"
)

var (
	networkSelfLink = "https://www.googleapis.com/mockservice/v1/projects/yuwenma-gke-dev/global/networks/computenetwork-52ldcmpbgxfmizhrk74q"
)

func TestNew(t *testing.T) {
	tests := []struct {
		name                string
		gcpResource         GCPResource
		expectedHasSelfLink bool
		expectedExternalRef string
	}{
		{
			name:                "cloudbuild workerpool, no selfLink",
			gcpResource:         &cloudbuildpb.WorkerPool{Name: "projects/mock-project/locations/us-central1/workerPools/mock-pool"},
			expectedHasSelfLink: false,
			expectedExternalRef: "https://www.googleapis.com/mockservice/v1/projects/mock-project/locations/us-central1/workerPools/mock-pool",
		},
		{
			name:                "compute network, has seflLink as string pointer",
			gcpResource:         &computepb.Network{SelfLink: &networkSelfLink},
			expectedHasSelfLink: true,
			expectedExternalRef: networkSelfLink,
		},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			gcpObj := test.gcpResource
			actualRef := New("https://www.googleapis.com/mockservice/v1/", gcpObj)
			if actualRef.hasSelfLink != test.expectedHasSelfLink {
				t.Errorf("expected %v, got %v", test.expectedHasSelfLink, actualRef.hasSelfLink)
			}
			if actualRef.externalRef != test.expectedExternalRef {
				t.Errorf("expected %v, got %v", test.expectedExternalRef, actualRef.externalRef)
			}
		})
	}
}
