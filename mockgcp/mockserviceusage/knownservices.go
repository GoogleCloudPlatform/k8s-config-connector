// Copyright 2022 Google LLC
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//	http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package mockserviceusage

var allServices = []string{
	"apigee.googleapis.com",
	"bigquery.googleapis.com",
	"compute.googleapis.com",
	"pubsub.googleapis.com",
	"runtimeconfig.googleapis.com",
	"storage.googleapis.com",
	"gkehub.googleapis.com",
	"anthos.googleapis.com",
	"anthosconfigmanagement.googleapis.com",
	"anthospolicycontroller.googleapis.com",
	"managedkafka.googleapis.com",
	"multiclusteringress.googleapis.com",
	"multiclusterservicediscovery.googleapis.com",
	"mesh.googleapis.com",
	"servicenetworking.googleapis.com",
	"spanner.googleapis.com",
	"vpcaccess.googleapis.com",
	"container.googleapis.com",
	"workflows.googleapis.com",
}

func isKnownService(serviceName string) bool {
	for _, service := range allServices {
		if service == serviceName {
			return true
		}
	}
	return false
}
