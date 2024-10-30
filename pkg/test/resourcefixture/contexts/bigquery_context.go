// Copyright 2022 Google LLC
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

package contexts

func init() {
	resourceContextMap["bigqueryjob"] = ResourceContext{
		ResourceKind: "BigQueryJob",
		SkipUpdate:   true,
	}

	// BigQueryConnectionConnection is a service-generate-ID resource.
	// Drift-deletion removes the object from the GCP server and test
	// if Config Connector can recreate the object. This does not fit the
	// service-generated-ID resource.
	resourceContextMap["bigqueryconnectionconnectionbasic"] = ResourceContext{
		ResourceKind: "BigQueryConnectionConnection",

		SkipDriftDetection: true,
	}
	resourceContextMap["bigqueryconnectionconnectionfull"] = ResourceContext{
		ResourceKind:       "BigQueryConnectionConnection",
		SkipDriftDetection: true,
	}
	resourceContextMap["awsconnectionbasic"] = ResourceContext{
		ResourceKind:       "BigQueryConnectionConnection",
		SkipDriftDetection: true,
	}
	resourceContextMap["azureconnectionbasic"] = ResourceContext{
		ResourceKind:       "BigQueryConnectionConnection",
		SkipDriftDetection: true,
	}
	resourceContextMap["cloudspannerconnectionbasic"] = ResourceContext{
		ResourceKind:       "BigQueryConnectionConnection",
		SkipDriftDetection: true,
	}
	resourceContextMap["cloudsqlconnectionbasic"] = ResourceContext{
		ResourceKind:       "BigQueryConnectionConnection",
		SkipDriftDetection: true,
	}
	resourceContextMap["sparkconnectionbasic"] = ResourceContext{
		ResourceKind:       "BigQueryConnectionConnection",
		SkipDriftDetection: true,
	}
}
