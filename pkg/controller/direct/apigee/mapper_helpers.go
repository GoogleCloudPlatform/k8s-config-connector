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

package apigee

import (
	"time"

	apigeev1beta1 "github.com/GoogleCloudPlatform/k8s-config-connector/apis/apigee/v1beta1"
)

// ConvertEpochMillisToTimestamp converts an int64 representing milliseconds since epoch
// into a human-readable timestamp string in RFC3339 format.
func ConvertEpochMillisToTimestamp(epochMillis int64) string {
	t := time.UnixMilli(epochMillis).UTC()
	return t.Format(time.RFC3339)
}

// ConvertTimestampToEpochMillis converts an RFC3339 timestamp string
// into an int64 representing milliseconds since epoch.
func ConvertTimestampToEpochMillis(timestamp string) (int64, error) {
	t, err := time.Parse(time.RFC3339, timestamp)
	if err != nil {
		return 0, err
	}
	return t.UnixMilli(), nil
}

// GetNameOfEnvironment gets the name of the environment by parsing its external reference string
func GetNameOfEnvironment(external string) (string, error) {
	environmentID := &apigeev1beta1.ApigeeEnvironmentIdentity{}
	if err := environmentID.FromExternal(external); err != nil {
		return "", err
	}
	return environmentID.ResourceID, nil
}
