// Copyright 2025 Google LLC
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

package v1alpha1


// +kcc:proto=google.cloud.recommender.v1beta1.InsightType
type InsightType struct {
	// The insight_type’s name in format insightTypes/{insight_type}
	//  eg: insightTypes/google.iam.policy.Insight
	// +kcc:proto:field=google.cloud.recommender.v1beta1.InsightType.name
	Name *string `json:"name,omitempty"`
}
