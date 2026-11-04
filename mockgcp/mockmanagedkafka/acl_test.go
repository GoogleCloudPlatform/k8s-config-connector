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

package mockmanagedkafka

import (
	"testing"
)

func TestDeriveAclFields(t *testing.T) {
	tests := []struct {
		aclID            string
		wantResourceType string
		wantResourceName string
		wantPatternType  string
		wantErr          bool
	}{
		{
			aclID:            "cluster",
			wantResourceType: "CLUSTER",
			wantResourceName: "kafka-cluster",
			wantPatternType:  "LITERAL",
		},
		{
			aclID:            "allTopics",
			wantResourceType: "TOPIC",
			wantResourceName: "*",
			wantPatternType:  "LITERAL",
		},
		{
			aclID:            "allConsumerGroups",
			wantResourceType: "GROUP",
			wantResourceName: "*",
			wantPatternType:  "LITERAL",
		},
		{
			aclID:            "allTransactionalIds",
			wantResourceType: "TRANSACTIONAL_ID",
			wantResourceName: "*",
			wantPatternType:  "LITERAL",
		},
		{
			aclID:            "topic/my-topic",
			wantResourceType: "TOPIC",
			wantResourceName: "my-topic",
			wantPatternType:  "LITERAL",
		},
		{
			aclID:            "consumerGroup/my-group",
			wantResourceType: "GROUP",
			wantResourceName: "my-group",
			wantPatternType:  "LITERAL",
		},
		{
			aclID:            "transactionalId/my-tx",
			wantResourceType: "TRANSACTIONAL_ID",
			wantResourceName: "my-tx",
			wantPatternType:  "LITERAL",
		},
		{
			aclID:            "topicPrefixed/my-prefix",
			wantResourceType: "TOPIC",
			wantResourceName: "my-prefix",
			wantPatternType:  "PREFIXED",
		},
		{
			aclID:            "consumerGroupPrefixed/my-group-prefix",
			wantResourceType: "GROUP",
			wantResourceName: "my-group-prefix",
			wantPatternType:  "PREFIXED",
		},
		{
			aclID:            "transactionalIdPrefixed/my-tx-prefix",
			wantResourceType: "TRANSACTIONAL_ID",
			wantResourceName: "my-tx-prefix",
			wantPatternType:  "PREFIXED",
		},
		{
			aclID:   "invalidAclId",
			wantErr: true,
		},
	}

	for _, tc := range tests {
		t.Run(tc.aclID, func(t *testing.T) {
			rType, rName, pType, err := deriveAclFields(tc.aclID)
			if tc.wantErr {
				if err == nil {
					t.Fatalf("deriveAclFields(%q) expected error, got nil", tc.aclID)
				}
				return
			}
			if err != nil {
				t.Fatalf("deriveAclFields(%q) unexpected error: %v", tc.aclID, err)
			}
			if rType != tc.wantResourceType {
				t.Errorf("deriveAclFields(%q) resourceType = %q, want %q", tc.aclID, rType, tc.wantResourceType)
			}
			if rName != tc.wantResourceName {
				t.Errorf("deriveAclFields(%q) resourceName = %q, want %q", tc.aclID, rName, tc.wantResourceName)
			}
			if pType != tc.wantPatternType {
				t.Errorf("deriveAclFields(%q) patternType = %q, want %q", tc.aclID, pType, tc.wantPatternType)
			}
		})
	}
}
