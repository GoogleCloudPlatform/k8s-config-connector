// Copyright 2024 Google LLC. All Rights Reserved.
// 
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
// 
//     http://www.apache.org/licenses/LICENSE-2.0
// 
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.
package beta

import (
	"strings"

	"github.com/GoogleCloudPlatform/declarative-resource-client-library/dcl"
)

// ExpandFunctionEventTriggerResource handles the special case for StorageBucket
// and makes sure resource field has format projects/{{project}}/buckets/{{name}}.
// http://b/210923505
func ExpandFunctionEventTriggerResource(_ *Client, _ *string, res *Function) (*string, error) {
	if res.EventTrigger == nil || res.EventTrigger.Resource == nil {
		return nil, nil
	}

	if dcl.ValueOrEmptyString(res.EventTrigger.EventType) == "providers/cloud.storage/eventTypes/object.change" {
		return dcl.DeriveField("projects/%s/buckets/%s", res.EventTrigger.Resource, dcl.SelfLinkToName(res.Project), dcl.SelfLinkToName(res.EventTrigger.Resource))
	}

	return res.EventTrigger.Resource, nil
}

// ExpandFunctionEventTriggerFailurePolicy inverts the flattenFunctionEventTriggerFailurePolicy transformation.
func ExpandFunctionEventTriggerFailurePolicy(_ *Client, _ *bool, res *Function) (interface{}, error) {
	if res.EventTrigger == nil || res.EventTrigger.FailurePolicy == nil || !*(res.EventTrigger.FailurePolicy) {
		return nil, nil
	}
	return map[string]interface{}{
		"retry": map[string]interface{}{},
	}, nil
}

// flattenFunctionEventTriggerFailurePolicy converts the API reprensentation of an event
// trigger retry policy, which is true or false based on the presence or absence
// of an empty object, to an actual bool for convenience purposes.
func flattenFunctionEventTriggerFailurePolicy(_ *Client, i interface{}, _ *Function) *bool {
	if _, ok := i.(map[string]interface{}); ok {
		return dcl.Bool(true)
	}
	return nil
}

// CanonicalizeFunctionSourceRepoURL compares source repo url because /paths/ can be omitted but will be added by the API.
func CanonicalizeFunctionSourceRepoURL(m interface{}, n interface{}) bool {
	mVal, ok := m.(*string)
	if !ok {
		return false
	}
	nVal, ok := n.(*string)
	if !ok {
		return false
	}
	if mVal == nil && nVal == nil {
		return true
	}
	if mVal == nil || nVal == nil {
		return false
	}
	// Both values are pointers to strings. Compare values with suffix trimmed.
	return strings.TrimSuffix(*mVal, "/paths/") == strings.TrimSuffix(*nVal, "/paths/")
}
