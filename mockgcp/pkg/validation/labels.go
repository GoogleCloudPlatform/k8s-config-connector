// Copyright 2024 Google LLC
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

package validation

import (
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

// ValidateLabels checks if a map of labels conforms to GCP's label requirements.
func ValidateLabels(labels map[string]string) error {
	for k, v := range labels {
		if len(k) > 63 {
			return status.Errorf(codes.InvalidArgument, "resource labels are invalid: key %q is longer than 63 characters", k)
		}
		if len(k) == 0 {
			return status.Errorf(codes.InvalidArgument, "resource labels are invalid: key cannot be empty")
		}
		if (k[0] < 'a' || k[0] > 'z') && (k[0] < 'A' || k[0] > 'Z') {
			return status.Errorf(codes.InvalidArgument, "resource labels are invalid: key %q must start with a letter", k)
		}
		for i, r := range k {
			if (r >= 'a' && r <= 'z') || (r >= 'A' && r <= 'Z') || (r >= '0' && r <= '9') || r == '_' || r == '-' {
				continue
			}
			return status.Errorf(codes.InvalidArgument, "resource labels are invalid: key %q contains invalid character '%c' at index %d", k, r, i)
		}

		if len(v) > 63 {
			return status.Errorf(codes.InvalidArgument, "resource labels are invalid: value for key %q is longer than 63 characters", k)
		}
		for i, r := range v {
			if (r >= 'a' && r <= 'z') || (r >= 'A' && r <= 'Z') || (r >= '0' && r <= '9') || r == '_' || r == '-' {
				continue
			}
			return status.Errorf(codes.InvalidArgument, "resource labels are invalid: value for key %q contains invalid character '%c' at index %d", k, r, i)
		}
	}
	return nil
}
