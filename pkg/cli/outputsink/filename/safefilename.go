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

package filename

import (
	"fmt"
	"strings"
)

// MakeSafeFilename escapes a filepath component to avoid problematic characters,
// and to avoid hidden files.
func MakeSafeFilename(in string) string {
	var sb strings.Builder
	for _, r := range in {
		if '0' <= r && r <= '9' {
			sb.WriteRune(r)
		} else if 'a' <= r && r <= 'z' {
			sb.WriteRune(r)
		} else if 'A' <= r && r <= 'Z' {
			sb.WriteRune(r)
		} else {
			switch r {
			case '_', '-', '.':
				sb.WriteRune(r)
			default:
				sb.WriteString(fmt.Sprintf("_%x", r))
			}
		}
	}
	s := sb.String()
	if strings.HasPrefix(s, ".") {
		// Avoid hidden files and ..
		s = "_" + s
	}
	return s
}
