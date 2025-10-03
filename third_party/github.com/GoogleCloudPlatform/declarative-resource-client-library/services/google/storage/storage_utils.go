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
// Package storage contains handwritten support code for the storage service.
package storage

import (
	"fmt"
)

// expandStorageBucketLifecycleWithState expands the with_state enum into the
// ternary boolean is_live. It can be true, false, or unset each corresponding
// to a different value.
func expandStorageBucketLifecycleWithState(_ *Client, v *BucketLifecycleRuleConditionWithStateEnum, _ *Bucket) (interface{}, error) {
	switch *v {
	case "LIVE":
		b := true
		return &b, nil
	case "ARCHIVED":
		b := false
		return &b, nil
	case "ANY":
		return nil, nil
	}

	return nil, fmt.Errorf("unrecognized BucketLifecycleRuleConditionWithStateEnum value: %v", v)
}

// flattenStorageBucketLifecycleWithState flattens the ternary boolean is_live
// into the with_state enum. It can be true, false, or unset each corresponding
// to a different value.
func flattenStorageBucketLifecycleWithState(_ *Client, v interface{}, _ *Bucket) *BucketLifecycleRuleConditionWithStateEnum {
	b, ok := v.(bool)
	if !ok { // b is unset
		return BucketLifecycleRuleConditionWithStateEnumRef("ANY")
	}

	if b {
		return BucketLifecycleRuleConditionWithStateEnumRef("LIVE")
	}

	// b is false
	return BucketLifecycleRuleConditionWithStateEnumRef("ARCHIVED")
}
