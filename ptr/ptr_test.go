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

package ptr

import "testing"

func TestValueOf(t *testing.T) {
	var nilStr *string
	if got := ValueOf(nilStr); got != "" {
		t.Errorf("ValueOf(nil) = %q, want empty string", got)
	}

	s := "hello"
	if got := ValueOf(&s); got != "hello" {
		t.Errorf("ValueOf(&s) = %q, want %q", got, "hello")
	}
}

func TestPtrTo(t *testing.T) {
	p := PtrTo(42)
	if p == nil || *p != 42 {
		t.Errorf("PtrTo(42) = %v, want pointer to 42", p)
	}

	p2 := To("abc")
	if p2 == nil || *p2 != "abc" {
		t.Errorf("To(\"abc\") = %v, want pointer to \"abc\"", p2)
	}
}

func TestLazyPtr(t *testing.T) {
	if got := LazyPtr(""); got != nil {
		t.Errorf("LazyPtr(\"\") = %v, want nil", got)
	}
	if got := LazyPtr(0); got != nil {
		t.Errorf("LazyPtr(0) = %v, want nil", got)
	}
	if got := LazyPtr("val"); got == nil || *got != "val" {
		t.Errorf("LazyPtr(\"val\") = %v, want pointer to \"val\"", got)
	}
}
