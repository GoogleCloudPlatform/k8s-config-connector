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

package mockcompute

import "time"

func (s *MockService) nowString() string {
	now := time.Now()
	return s.formatDate(now)
}

func (s *MockService) formatDate(t time.Time) string {
	format := "2006-01-02T15:04:05.999Z07:00"

	return t.Format(format)
}

func (s *MockService) generateID() uint64 {
	return uint64(time.Now().UnixNano())
}

func PtrTo[T any](t T) *T {
	return &t
}
