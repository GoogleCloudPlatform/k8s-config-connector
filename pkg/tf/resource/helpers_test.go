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

package resource

import (
	"strings"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/terraform"
)

func TestFieldValuesAreShown(t *testing.T) {
	// Define the attribute difference
	attrDiff := &terraform.ResourceAttrDiff{
		Old:         "foo",
		New:         "bar",
		RequiresNew: true,
	}

	// Create a map of attributes containing the attrDiff
	attributes := map[string]*terraform.ResourceAttrDiff{
		"testfield": attrDiff,
	}

	// Create a new InstanceDiff using the attributes map
	instanceDiff := &terraform.InstanceDiff{
		Attributes: attributes,
	}

	fields := ImmutableFieldsFromDiff(instanceDiff)

	if !keywordIncluded(fields, "foo") || !keywordIncluded(fields, "bar") {
		t.Fatalf("Field value missed from diff: %v: ", fields)
	}
}

func TestSensitiveFieldValuesAreHidden(t *testing.T) {
	// Define the attribute difference
	attrDiff := &terraform.ResourceAttrDiff{
		Old:         "foo",
		New:         "bar",
		RequiresNew: true,
		Sensitive:   true,
	}

	// Create a map of attributes containing the attrDiff
	attributes := map[string]*terraform.ResourceAttrDiff{
		"testfield": attrDiff,
	}

	// Create a new InstanceDiff using the attributes map
	instanceDiff := &terraform.InstanceDiff{
		Attributes: attributes,
	}

	fields := ImmutableFieldsFromDiff(instanceDiff)

	if keywordIncluded(fields, "foo") || keywordIncluded(fields, "bar") {
		t.Fatalf("Sensitive field value exposed in diff: %v: ", fields)
	}
}

func keywordIncluded(slice []string, keyword string) bool {
	for _, s := range slice {
		if strings.Contains(s, keyword) {
			return true
		}
	}
	return false
}
