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

package direct

import (
	"fmt"
	"strings"

	refs "github.com/GoogleCloudPlatform/k8s-config-connector/apis/refs/v1beta1"
	"k8s.io/klog/v2"
)

// TODO If need support for identifying TagValueRef by Name, need to do apis/refs/v1beta1/tagrefs.go ResolveTagValueRef with a client.Reader

func Tags_ToProto(mapCtx *MapContext, in []*refs.TagValueRef) map[string]string {
	const WantedExternalTokenCount int = 3

	if in == nil {
		return nil
	}

	out := make(map[string]string)
	for k, v := range in {
		tokens := strings.Split(v.External, "/")
		if len(tokens) == WantedExternalTokenCount {
			key := strings.Join(tokens[:2], "/")
			if tokens[0] == "tagKeys" {
				// For `tagKeys/[tag_key_id]` mapped to `tagValues/[tag_value_id]`, stored as `tagKeys/[tag_key_id]/[tag_value_id]`
				value := tokens[2]
				out[key] = "tagValues/" + value
			} else {
				// For `[org id, project id, or project number]/[tag_key_shortname]` mapped to `[value_shortname]`
				value := tokens[2]
				out[key] = value
			}
		} else {
			// Shouldn't reach here if tags were well formed
			if v.Name != "" {
				klog.Warningf("Skipping TagValueRef %v because we cannot handle identification by Name alone.", k)
			} else {
				klog.Warningf("Skipping TagValueRef %v because it is not well formed. External token count is %v but wanted %v", k, len(tokens), WantedExternalTokenCount)
			}
		}
	}

	return out
}

func Tags_FromProto(mapCtx *MapContext, in map[string]string) []*refs.TagValueRef {
	const WantedKeyTokenCount int = 2
	const WantedValueTokenCount int = 2

	if in == nil {
		return nil
	}

	out := make([]*refs.TagValueRef, 0)
	for k, v := range in {
		tagValueRef := &refs.TagValueRef{}
		keyTokens := strings.Split(k, "/")
		valueTokens := strings.Split(v, "/")
		if len(keyTokens) == WantedKeyTokenCount {
			// For `tagKeys/[tag_key_id]` mapped to `tagValues/[tag_value_id]`
			// and `[org id, project id, or project number]/[tag_key_shortname]` mapped to `[value_shortname]`
			tagValueRef.External = fmt.Sprintf("%s/%s", keyTokens[0], keyTokens[1], valueTokens[len(valueTokens) - 1])
		} else {
			// Shouldn't reach here if tags were well formed
			klog.Warningf("Skipping Tag because it is not well formed. Key token count is %v but wanted %v, value token count is %v but wanted %v. Consult Resource Manager Tags Overview for more info", len(keyTokens), WantedKeyTokenCount, len(valueTokens), WantedValueTokenCount)
			continue
		}
		out = append(out, tagValueRef)
	}

	return out
}
