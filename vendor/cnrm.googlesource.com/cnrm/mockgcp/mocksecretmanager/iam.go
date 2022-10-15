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

package mocksecretmanager

import (
	"context"

	iamv1 "google.golang.org/genproto/googleapis/iam/v1"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

// Sets the access control policy on the specified secret. Replaces any
// existing policy.
//
// Permissions on [SecretVersions][google.cloud.secretmanager.v1.SecretVersion] are enforced according
// to the policy set on the associated [Secret][google.cloud.secretmanager.v1.Secret].
func (s *MockService) SetIamPolicy(context.Context, *iamv1.SetIamPolicyRequest) (*iamv1.Policy, error) {
	return nil, status.Errorf(codes.Unimplemented, "method not implemented")
}

// Gets the access control policy for a secret.
// Returns empty policy if the secret exists and does not have a policy set.
func (s *MockService) GetIamPolicy(context.Context, *iamv1.GetIamPolicyRequest) (*iamv1.Policy, error) {
	return nil, status.Errorf(codes.Unimplemented, "method not implemented")
}

// Returns permissions that a caller has for the specified secret.
// If the secret does not exist, this call returns an empty set of
// permissions, not a NOT_FOUND error.
//
// Note: This operation is designed to be used for building permission-aware
// UIs and command-line tools, not for authorization checking. This operation
// may "fail open" without warning.
func (s *MockService) TestIamPermissions(context.Context, *iamv1.TestIamPermissionsRequest) (*iamv1.TestIamPermissionsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method not implemented")
}
