// Copyright 2024 Google LLC
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

package plan

import (
	"cmp"
	"strings"
)

func compareStringSlices(l []string, r []string) int {
	lLen := len(l)
	rLen := len(r)
	minLen := min(lLen, rLen)
	for i := 0; i < minLen; i++ {
		lv := l[i]
		rv := r[i]
		if x := strings.Compare(lv, rv); x != 0 {
			return x
		}
	}
	if lLen < rLen {
		return -1
	}
	if rLen > lLen {
		return 1
	}
	return 0
}

func min[T cmp.Ordered](l, r T) T {
	if l < r {
		return l
	}
	return r
}
