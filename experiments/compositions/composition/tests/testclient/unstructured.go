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

package testclient

import (
	"testing"

	"k8s.io/apimachinery/pkg/apis/meta/v1/unstructured"
)

// Note: The functions below can be used in Polling because they are called on
// paths holding fields that should absolutely be of the specified type. If not,
// there is a fundamental error in which case we can exit the test immediately
// instead of telling Poll to err.
// MustGetBool - returns bool on u at path if set
func MustGetBool(t *testing.T, u *unstructured.Unstructured, path ...string) (value, set bool) {
	t.Helper()
	value, set, err := unstructured.NestedBool(u.Object, path...)
	if err != nil {
		t.Errorf("value at path %v not bool", path)
		t.FailNow()
	}
	return
}

// MustGetString - return string on u at path if set
func MustGetString(t *testing.T, u *unstructured.Unstructured, path ...string) (string, bool) {
	t.Helper()
	value, set, err := unstructured.NestedString(u.Object, path...)
	if err != nil {
		t.Errorf("value at path %v not string", path)
		t.FailNow()
	}
	return value, set
}

// GetInt - return int64 on u at path if set
func MustGetInt64(t *testing.T, u *unstructured.Unstructured, path ...string) (int64, bool) {
	t.Helper()
	value, set, err := unstructured.NestedInt64(u.Object, path...)
	if err != nil {
		t.Errorf("value at path %v not int64", path)
		t.FailNow()
	}
	return value, set
}
