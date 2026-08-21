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

package mocks

// PtrTo returns a pointer to the given value.
func PtrTo[T any](v T) *T {
	return &v
}

// ValueOf returns the value pointed to by the given pointer.
// If the pointer is nil, it returns the zero value of the type.
func ValueOf[T any](v *T) T {
	if v == nil {
		var zero T
		return zero
	}
	return *v
}
