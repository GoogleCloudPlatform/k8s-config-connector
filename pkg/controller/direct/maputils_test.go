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

package direct

import (
	"testing"

	"google.golang.org/protobuf/types/known/durationpb"
)

func TestStringDuration_FromProto(t *testing.T) {
	mapctx := &MapContext{}
	d := &durationpb.Duration{Seconds: 34312, Nanos: 20}
	krm := StringDuration_FromProto(mapctx, d)
	if *krm != "9h31m52.00000002s" {
		t.Fatalf("google.protobuf.Duration -> string, expect \"9h31m52.00000002s\", got %s", *krm)
	}
	if mapctx.Err() != nil {
		t.Fatalf("google.protobuf.Duration -> string error: %s", mapctx.Err())
	}
}

func TestStringDuration_ToProto(t *testing.T) {
	mapctx := &MapContext{}
	s := "1h1m"
	d := StringDuration_ToProto(mapctx, &s)
	if d.Seconds != 3660 || d.Nanos != 0 {
		t.Fatalf("string -> google.protobuf.Duration, expect \"seconds:3660 nanos:00\", got %s", d)
	}
	if mapctx.Err() != nil {
		t.Fatalf("google.protobuf.Duration -> String error: %s", mapctx.Err())
	}
}
