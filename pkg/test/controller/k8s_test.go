// Copyright 2026 Google LLC
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

package testcontroller

import (
	"testing"
	"time"
)

func TestReplaceFutureTimestampDays(t *testing.T) {
	now := time.Date(2026, 11, 20, 17, 54, 16, 0, time.FixedZone("PST", -8*60*60))

	tests := []struct {
		name string
		in   string
		want string
	}{
		{
			name: "single placeholder",
			in:   `expireTime: "${futureTimestampDays:30}"`,
			want: `expireTime: "2026-12-21T01:00:00Z"`,
		},
		{
			name: "multiple placeholders",
			in:   "a: ${futureTimestampDays:0}\nb: ${futureTimestampDays:15}",
			want: "a: 2026-11-21T01:00:00Z\nb: 2026-12-06T01:00:00Z",
		},
		{
			name: "fixed futureTimestamp untouched",
			in:   "${futureTimestamp}",
			want: "${futureTimestamp}",
		},
		{
			name: "malformed placeholder untouched",
			in:   "${futureTimestampDays:abc}",
			want: "${futureTimestampDays:abc}",
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			if got := replaceFutureTimestampDays(tc.in, now); got != tc.want {
				t.Errorf("replaceFutureTimestampDays(%q) = %q, want %q", tc.in, got, tc.want)
			}
		})
	}
}
