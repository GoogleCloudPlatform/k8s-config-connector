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

package sampleconversion

import (
	"fmt"
	"regexp"
	"strings"

	"k8s.io/apimachinery/pkg/runtime/schema"
)

// GetTFReferenceValue extracts "TF Reference Value" in the format of
// `[referenced_tf_type].[referenced_resource_name].[referenced_field_name]`,
// "Value Template" if the input contains more than the "TF Reference Value",
// and a bool result of whether the input contains a TF ReferenceValue.
func GetTFReferenceValue(value interface{}) (tfRefValue, valueTemplate string, containsTFReference bool, err error) {
	str, ok := value.(string)
	if !ok {
		return "", "", false, nil
	}
	tfRefValueRegex := regexp.MustCompile(`\${(google_[a-z_]+[a-z]\.[a-z](?:[a-z_-]*[a-z])*\.[a-z](?:[a-z_]*[a-z])*)}`)
	matchResult := tfRefValueRegex.FindStringSubmatch(str)
	if len(matchResult) == 0 {
		return "", "", false, nil
	}

	// If the value itself is a TF reference.
	if matchResult[0] == str {
		return matchResult[1], "", true, nil
	}

	if len(matchResult) > 2 {
		return "", "", false, fmt.Errorf("cannot handle more than one TF references: %v", str)
	}

	valueTemplate = strings.ReplaceAll(str, matchResult[0], "{{value}}")
	return matchResult[1], valueTemplate, true, nil
}

func ConstructKRMNameReferenceObject(value string, tfToGVK map[string]schema.GroupVersionKind) (map[string]interface{}, error) {
	tfType := strings.Split(value, ".")[0]
	gvk, ok := tfToGVK[tfType]
	if !ok {
		return nil, fmt.Errorf("unsupported reference TF type: %v", tfType)
	}
	nameVal := fmt.Sprintf("%s-${uniqueID}", strings.ToLower(gvk.Kind))
	refVal := make(map[string]interface{})
	refVal["name"] = nameVal
	return refVal, nil
}

func ConstructKRMExternalReferenceObject(value string) map[string]interface{} {
	refVal := make(map[string]interface{})
	refVal["external"] = value
	return refVal
}

// ConstructKRMExternalRefValFromTFRefVal constructs the external referenced
// value of the reference field with "TF Reference Value" in the format of
// `[referenced_tf_type].[referenced_resource_name].[referenced_field_name]` and
// the "Value Template".
func ConstructKRMExternalRefValFromTFRefVal(tfRefVal, valueTemplate string, tfToGVK map[string]schema.GroupVersionKind) (string, error) {
	tfType, field := extractReferencedTFTypeAndField(tfRefVal)

	if field != "name" {
		return "", fmt.Errorf("unsupported referenced field: %v", field)
	}
	gvk, ok := tfToGVK[tfType]
	if !ok {
		return "", fmt.Errorf("unsupported reference type: %v", tfType)
	}
	nameVal := fmt.Sprintf("%s-${uniqueID}", strings.ToLower(gvk.Kind))
	return strings.ReplaceAll(valueTemplate, "{{value}}", nameVal), nil
}

func extractReferencedTFTypeAndField(tfRefVal string) (tfType, field string) {
	parts := strings.Split(tfRefVal, ".")
	tfType = parts[0]
	field = parts[2]
	return tfType, field
}
