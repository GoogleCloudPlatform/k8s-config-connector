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

package common

import "google.golang.org/protobuf/types/known/timestamppb"

func DeepEqual_StringAndTimestampPb(a string, b *timestamppb.Timestamp) bool {
	norm := NormalizeStringToTimestamp(a)
	if norm == nil {
		return b == nil
	}
	if b == nil {
		return false
	}
	x := norm.String()
	y := b.String()
	return x == y
}
