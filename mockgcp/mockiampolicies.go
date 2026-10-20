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
	"slices"
	"sort"

	"cloud.google.com/go/iam/apiv1/iampb"
	"google.golang.org/protobuf/encoding/protojson"
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

func (m *mockIAMPolicies) buildResponse(obj proto.Message) (*http.Response, error) {
	// Use json.Marshal via a map to get alphabetical field order for stability with existing golden files
	jsonBytes, err := protojson.Marshal(obj)
	if err != nil {
		return nil, err
	}
	var data map[string]any
	if err := json.Unmarshal(jsonBytes, &data); err != nil {
		return nil, err
	}
	b, err := json.Marshal(data)
	if err != nil {
		return nil, err
	}

	body := io.NopCloser(bytes.NewReader(b))
	w := &http.Response{StatusCode: http.StatusOK, Body: body}

	w.Status = fmt.Sprintf("%d %s", w.StatusCode, http.StatusText(w.StatusCode))

	if w.Header == nil {
		w.Header = make(http.Header)
	}
	w.Header.Set("Content-Type", "application/json; charset=UTF-8")
	w.Header.Set("Cache-Control", "private")
	w.Header.Set("Server", "ESF")
	w.Header["Vary"] = []string{"Origin", "X-Origin", "Referer"}
	w.Header.Set("X-Content-Type-Options", "nosniff")
	w.Header.Set("X-Frame-Options", "SAMEORIGIN")
	w.Header.Set("X-Xss-Protection", "0")

	return w, nil
}

func (m *mockIAMPolicies) getIAMPolicy(resourcePath string) (*iampb.Policy, error) {
	policy := m.policies[resourcePath]
	if policy == nil {
		policy = &iampb.Policy{}
	}

	// Return a copy to avoid mutation
	policy = proto.Clone(policy).(*iampb.Policy)

	// Sort for determinism
	sortPolicy(policy)

	if policy.Etag == nil {
		policy.Etag = computeEtag(policy)
	}

	return policy, nil
}

func (m *mockIAMPolicies) serveGetIAMPolicy(resourcePath string) (*http.Response, error) {
	policy, err := m.getIAMPolicy(resourcePath)
	if err != nil {
		return nil, err
	}
	return m.buildResponse(policy)
}

func (m *mockIAMPolicies) serveSetIAMPolicy(resourcePath string, httpRequest *http.Request) (*http.Response, error) {
	request := &iampb.SetIamPolicyRequest{}

	requestBytes, err := io.ReadAll(httpRequest.Body)
	if err != nil {
		return nil, err
	}
	if err := protojson.Unmarshal(requestBytes, request); err != nil {
		return nil, err
	}

	oldPolicy, err := m.getIAMPolicy(resourcePath)
	if err != nil {
		return nil, err
	}

	if request.Policy.Etag != nil && !bytes.Equal(request.Policy.Etag, oldPolicy.Etag) {
		responseBytes := []byte("{}")
		responseBody := io.NopCloser(bytes.NewReader(responseBytes))
		return &http.Response{StatusCode: http.StatusConflict, Body: responseBody}, nil
	}

	// conditional role bindings must specify version 3
	hasConditions := false
	for _, binding := range request.Policy.Bindings {
		if binding.Condition != nil {
			hasConditions = true
			break
		}
	}
	// GCP returns the version as 1 if there are no conditions
	if !hasConditions {
		request.Policy.Version = 1
	}

	sortPolicy(request.Policy)

	request.Policy.Etag = computeEtag(request.Policy)
	m.policies[resourcePath] = request.Policy

	return m.buildResponse(request.Policy)
}

func sortPolicy(policy *iampb.Policy) {
	sort.Slice(policy.Bindings, func(i, j int) bool {
		if policy.Bindings[i].Role != policy.Bindings[j].Role {
			return policy.Bindings[i].Role < policy.Bindings[j].Role
		}
		if policy.Bindings[i].Condition.GetTitle() != policy.Bindings[j].Condition.GetTitle() {
			return policy.Bindings[i].Condition.GetTitle() < policy.Bindings[j].Condition.GetTitle()
		}
		return policy.Bindings[i].Condition.GetExpression() < policy.Bindings[j].Condition.GetExpression()
	})
	for _, binding := range policy.Bindings {
		slices.Sort(binding.Members)
	}
	sort.Slice(policy.AuditConfigs, func(i, j int) bool {
		return policy.AuditConfigs[i].Service < policy.AuditConfigs[j].Service
	})
	for _, auditConfig := range policy.AuditConfigs {
		sort.Slice(auditConfig.AuditLogConfigs, func(i, j int) bool {
			return auditConfig.AuditLogConfigs[i].LogType < auditConfig.AuditLogConfigs[j].LogType
		})
	}
}

func computeEtag(policy *iampb.Policy) []byte {
	// Create a copy and clear fields that shouldn't affect etag
	temp := proto.Clone(policy).(*iampb.Policy)
	temp.Etag = nil

	b, err := proto.Marshal(temp)
	if err != nil {
		panic(fmt.Sprintf("converting to proto: %v", err))
	}
	hash := md5.Sum(b)
	return hash[:]
}
