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

package containeranalysis

import (
	"fmt"
	"time"

	"google.golang.org/protobuf/types/known/timestamppb"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"k8s.io/apimachinery/pkg/apis/meta/v1/unstructured"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/klog/v2"
)

func ToOpenAPIDateTime(ts *timestamppb.Timestamp) *string {
	formatted := ts.AsTime().Format(time.RFC3339)
	return &formatted
}

func setStatus(u *unstructured.Unstructured, typedStatus any) error {
	// TODO: Just fetch this object?
	status, err := runtime.DefaultUnstructuredConverter.ToUnstructured(typedStatus)
	if err != nil {
		return fmt.Errorf("error converting status to unstructured: %w", err)
	}
	// TODO: Merge to avoid overwriting conditions?
	u.Object["status"] = status

	return nil
}

func ValueOf[T any](p *T) T {
	var v T
	if p != nil {
		v = *p
	}
	return v
}

// IsNotFound returns true if the given error is canonical error code 5
// Canonical error codes: http://google3/util/task/codes.proto
func IsNotFound(err error) bool {
	if err == nil {
		return false
	}

	if status.Code(err) == codes.NotFound {
		return true
	} else {
		klog.Warningf("unexpected error type %T", err)
	}
	return false
}
