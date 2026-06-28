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
	pb "github.com/GoogleCloudPlatform/k8s-config-connector/mockgcp/generated/mockgcp/cloud/compute/v1"
)

func (s *MockService) populateHealthCheckDefaults(obj *pb.HealthCheck) {
	if obj.HttpsHealthCheck != nil && obj.HttpsHealthCheck.ProxyHeader == nil {
		obj.HttpsHealthCheck.ProxyHeader = PtrTo("NONE")
	}
	if obj.HttpHealthCheck != nil && obj.HttpHealthCheck.ProxyHeader == nil {
		obj.HttpHealthCheck.ProxyHeader = PtrTo("NONE")
	}
}
