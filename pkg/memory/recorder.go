// Copyright 2025 Google LLC
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//	http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package memory

import "runtime"

// MemoryReport holds memory usage statistics, and makes it easy to report basic memory usage for a test.
type MemoryRecorder struct {
	before runtime.MemStats
	after  runtime.MemStats
}

// StartMemoryRecorder captures the memory usage at the start of a test.
func StartMemoryRecorder() *MemoryRecorder {
	r := &MemoryRecorder{}
	runtime.GC()
	runtime.ReadMemStats(&r.before)
	return r
}

// Stop captures the memory usage at the end of a test, and returns a report of the differences.
func (r *MemoryRecorder) Stop() *MemoryRecorderReport {
	runtime.GC()
	runtime.ReadMemStats(&r.after)

	report := &MemoryRecorderReport{
		HeapAlloc:  int64(r.after.HeapAlloc) - int64(r.before.HeapAlloc),
		HeapInuse:  int64(r.after.HeapInuse) - int64(r.before.HeapInuse),
		TotalAlloc: int64(r.after.TotalAlloc) - int64(r.before.TotalAlloc),
	}
	return report
}

// MemoryRecorderReport holds memory usage statistics from a MemoryRecorder.
type MemoryRecorderReport struct {
	HeapAlloc  int64
	HeapInuse  int64
	TotalAlloc int64
}
