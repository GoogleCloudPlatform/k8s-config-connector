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

package allowlist

import (
	"testing"
)

func TestAllowlistsOrderedAlphabetically(t *testing.T) {
	var prevItem string
	for _, currItem := range alphaAllowlist {
		if prevItem > currItem {
			t.Errorf("alphaAllowlist items are not listed alphabetically: %s listed before %s", prevItem, currItem)
		}
		prevItem = currItem
	}
	prevItem = ""
	for _, currItem := range betaAllowlist {
		if prevItem > currItem {
			t.Errorf("betaAllowlist items are not listed alphabetically: %s listed before %s", prevItem, currItem)
		}
		prevItem = currItem
	}
}
