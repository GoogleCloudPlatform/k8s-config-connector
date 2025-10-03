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

import "testing"

func TestMakeSafeFilename(t *testing.T) {
	grid := []struct {
		In   string
		Want string
	}{
		{In: "foo", Want: "foo"},
		{In: "123", Want: "123"},
		{In: "Foo123", Want: "Foo123"},
		{In: "Foo-123", Want: "Foo-123"},
		{In: "Foo.123", Want: "Foo.123"},
		{In: "Foo_123", Want: "Foo_123"},
		{In: ".Foo_123", Want: "_.Foo_123"},
		{In: "$%?!@Foo", Want: "_24_25_3f_21_40Foo"},
	}

	for _, g := range grid {
		got := MakeSafeFilename(g.In)
		if got != g.Want {
			t.Errorf("got %v, want %v", got, g.Want)
		}
	}
}
