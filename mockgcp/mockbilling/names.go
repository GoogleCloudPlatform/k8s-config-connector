// Copyright 2023 Google LLC
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

package mockbilling

import (
	"strings"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type billingAccountName struct {
	BillingAccountName string
}

func (n *billingAccountName) String() string {
	return "billingAccounts/" + n.BillingAccountName
}

// parseBillingAccountName parses a string into a billingAccountName.
// The expected form is `billingAccounts/*`.
func (s *MockService) parseBillingAccountName(name string) (*billingAccountName, error) {
	tokens := strings.Split(name, "/")

	if len(tokens) == 2 && tokens[0] == "billingAccounts" {
		name := &billingAccountName{
			BillingAccountName: tokens[1],
		}

		return name, nil
	} else {
		return nil, status.Errorf(codes.InvalidArgument, "name %q is not valid", name)
	}
}
