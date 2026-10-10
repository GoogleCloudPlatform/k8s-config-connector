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

package codegen

// Reachable returns the nodes that seeds reach through next, seeds included.
// The walk doesn't enter a node for which stop returns true. A nil stop lets
// it enter every node.
func Reachable[T comparable](seeds []T, next func(T) []T, stop func(T) bool) map[T]bool {
	reached := map[T]bool{}
	queue := append([]T(nil), seeds...)
	for len(queue) > 0 {
		n := queue[0]
		queue = queue[1:]
		if reached[n] || (stop != nil && stop(n)) {
			continue
		}
		reached[n] = true
		queue = append(queue, next(n)...)
	}
	return reached
}
