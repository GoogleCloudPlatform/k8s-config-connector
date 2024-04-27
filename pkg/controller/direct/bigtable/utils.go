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

package bigtable

import (
	"errors"
	"fmt"
	"strings"

	"github.com/googleapis/gax-go/v2/apierror"
	"k8s.io/apimachinery/pkg/apis/meta/v1/unstructured"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/klog/v2"
)

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

func PtrTo[T any](t T) *T {
	return &t
}

func areSame[T comparable](l, r *T) bool {
	if l == nil {
		return r == nil
	}
	if r == nil {
		return l == nil
	}
	return *l == *r
}

// HasHTTPCode returns true if the given error is an HTTP response with the given code.
func HasHTTPCode(err error, code int) bool {
	if err == nil {
		return false
	}
	apiError := &apierror.APIError{}
	if errors.As(err, &apiError) {
		if apiError.HTTPCode() == code {
			return true
		}
	} else {
		klog.Warningf("unexpected error type %T", err)
	}
	return false
}

// IsNotFound returns true if the given error is an HTTP 404.
func IsNotFound(err error) bool {
	return HasHTTPCode(err, 404)
}

func lastComponent(s string) string {
	i := strings.LastIndex(s, "/")
	return s[i+1:]
}
