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

package mockgcp

import (
	"bytes"
	"crypto/md5"
	"encoding/json"
	"fmt"
	"io"
	"net/http"

	"cloud.google.com/go/iam/apiv1/iampb"
	"google.golang.org/protobuf/proto"
)

type mockIAMPolicies struct {
	policies map[string]*iampb.Policy
}

func newMockIAMPolicies() *mockIAMPolicies {
	return &mockIAMPolicies{
		policies: make(map[string]*iampb.Policy),
	}
}

func (m *mockIAMPolicies) serveGetIAMPolicy(resourcePath string) (*http.Response, error) {
	policy := m.policies[resourcePath]
	if policy == nil {
		policy = &iampb.Policy{}
		policy.Version = 3
		policy.Etag = computeEtag(policy)
	}
	b, err := json.Marshal(policy)
	if err != nil {
		return nil, err
	}
	body := io.NopCloser(bytes.NewReader(b))
	return &http.Response{StatusCode: http.StatusOK, Body: body}, nil
}

func (m *mockIAMPolicies) serveSetIAMPolicy(resourcePath string, httpRequest *http.Request) (*http.Response, error) {
	request := &iampb.SetIamPolicyRequest{}

	requestBytes, err := io.ReadAll(httpRequest.Body)
	if err != nil {
		return nil, err
	}
	if err := json.Unmarshal(requestBytes, request); err != nil {
		return nil, err
	}

	oldPolicy := m.policies[resourcePath]
	if oldPolicy == nil {
		oldPolicy = &iampb.Policy{}
		oldPolicy.Version = 3
		oldPolicy.Etag = computeEtag(oldPolicy)
	}

	if request.Policy.Etag != nil && !bytes.Equal(request.Policy.Etag, oldPolicy.Etag) {
		responseBytes := []byte("{}")
		responseBody := io.NopCloser(bytes.NewReader(responseBytes))
		return &http.Response{StatusCode: http.StatusConflict, Body: responseBody}, nil
	}

	request.Policy.Version = 3
	request.Policy.Etag = computeEtag(request.Policy)
	m.policies[resourcePath] = request.Policy

	responseBytes, err := json.Marshal(request.Policy)
	if err != nil {
		return nil, err
	}

	responseBody := io.NopCloser(bytes.NewReader(responseBytes))
	return &http.Response{StatusCode: http.StatusOK, Body: responseBody}, nil
}

func computeEtag(policy *iampb.Policy) []byte {
	b, err := proto.Marshal(policy)
	if err != nil {
		panic(fmt.Sprintf("converting to proto: %v", err))
	}
	hash := md5.Sum(b)
	return hash[:]
}
