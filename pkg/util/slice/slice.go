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

package slice

import "sort"

func StringSliceContains(l []string, s string) bool {
	for _, elem := range l {
		if elem == s {
			return true
		}
	}
	return false
}

// IncludeString inserts the given string to the sorted slice if not already
// included. It is assumed that the slice is already in sorted order.
func IncludeString(l []string, s string) []string {
	i := sort.Search(
		len(l),
		func(i int) bool {
			return l[i] >= s
		},
	)
	if i < len(l) && l[i] == s {
		// string is already in slice
		return l
	}
	l = append(l, "")
	copy(l[i+1:], l[i:])
	l[i] = s
	return l
}

func RemoveStringFromStringSlice(l []string, s string) []string {
	newSlice := make([]string, 0)
	for _, elem := range l {
		if elem != s {
			newSlice = append(newSlice, elem)
		}
	}
	return newSlice
}

func ConcatStringSlices(slices ...[]string) []string {
	newSlice := make([]string, 0)
	for _, s := range slices {
		newSlice = append(newSlice, s...)
	}
	return newSlice
}

// IsListOfStringInterfaceMaps returns true if the given list of interface{}
// values is actually a list of map[string]interface{} values. Returns false if
// the given list is empty.
func IsListOfStringInterfaceMaps(list []interface{}) bool {
	if len(list) == 0 {
		return false
	}
	for _, e := range list {
		_, ok := e.(map[string]interface{})
		if !ok {
			return false
		}
	}
	return true
}

func Reverse(slice []string) {
	length := len(slice)
	midIndex := length / 2

	for i := 0; i < midIndex; i++ {
		j := length - i - 1

		slice[i], slice[j] = slice[j], slice[i]
	}
}
