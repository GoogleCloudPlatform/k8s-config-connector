// Copyright 2022 Google LLC
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

package dcl

import "fmt"

// CanonicalizeIntegerValue converts the numeric value for integer type to int64 because that's the integer type
// used by DCL. During json marshalling, all numeric values are converted to be of type float64;
// we canonicalize them to int64 before sending to DCL.
func CanonicalizeIntegerValue(val interface{}) (int64, error) {
	switch val.(type) {
	case int64:
		return val.(int64), nil
	case int:
		return int64(val.(int)), nil
	case float64:
		return int64(val.(float64)), nil
	default:
		return 0, fmt.Errorf("expect to have one of the types (int, int64, float64) for the integer value, but got %T", val)
	}
}

// CanonicalizeNumberValue converts the numeric value for number type to float64 because that's the double type
// used by DCL.
func CanonicalizeNumberValue(val interface{}) (float64, error) {
	switch val.(type) {
	case float64:
		return val.(float64), nil
	case float32:
		return float64(val.(float32)), nil
	case int64:
		return float64(val.(int64)), nil
	case int:
		return float64(val.(int)), nil
	default:
		return 0, fmt.Errorf("expect to have one of the types (float64, float32, int64, int) for number value, but got %T", val)
	}
}
