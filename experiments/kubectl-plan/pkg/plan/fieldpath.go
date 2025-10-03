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

import "strings"

type FieldPath struct {
	parent *FieldPath
	part   string
}

func (f *FieldPath) With(part string) *FieldPath {
	return &FieldPath{parent: f, part: part}
}

func (f *FieldPath) asSlice() []string {
	n := 0
	pos := f
	for pos != nil {
		pos = pos.parent
		n++
	}
	ret := make([]string, n)
	i := n - 1
	pos = f
	for i >= 0 {
		ret[i] = pos.part
		pos = pos.parent
		i--
	}
	return ret
}

func (f *FieldPath) String() string {
	s := f.asSlice()
	return strings.Join(s, ".")
}
