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

package mockassuredworkloads

import (
	"strings"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type workloadName struct {
	OrganizationID string
	Location       string
	WorkloadID     string
}

func (n *workloadName) String() string {
	return "organizations/" + n.OrganizationID + "/locations/" + n.Location + "/workloads/" + n.WorkloadID
}

func (s *MockService) parseWorkloadName(name string) (*workloadName, error) {
	tokens := strings.Split(name, "/")
	if len(tokens) == 6 && tokens[0] == "organizations" && tokens[2] == "locations" && tokens[4] == "workloads" {
		return &workloadName{
			OrganizationID: tokens[1],
			Location:       tokens[3],
			WorkloadID:     tokens[5],
		}, nil
	}
	return nil, status.Errorf(codes.InvalidArgument, "name %q is not valid", name)
}
