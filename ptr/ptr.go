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

// Package ptr provides generic pointer and value conversion helpers with zero
// external dependencies.
package ptr

// ValueOf returns the value pointed to by p, or the zero value of T if p is nil.
func ValueOf[T any](p *T) T {
	var zeroVal T
	if p == nil {
		return zeroVal
	}
	return *p
}

// PtrTo returns a pointer to a copy of t.
func PtrTo[T any](t T) *T {
	return &t
}

// To returns a pointer to a copy of t (alias for PtrTo).
func To[T any](t T) *T {
	return &t
}

// LazyPtr returns a pointer to a copy of v, or nil if v is the zero value of V.
func LazyPtr[V comparable](v V) *V {
	var defaultV V
	if v == defaultV {
		return nil
	}
	return &v
}
