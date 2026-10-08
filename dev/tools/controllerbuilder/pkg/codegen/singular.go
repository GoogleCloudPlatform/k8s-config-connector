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

import (
	"strings"

	"k8s.io/klog/v2"
)

// Singular turns the plural collection segment of a resource pattern into the
// singular noun a reference is named after. For "keyRings/{key_ring}" it
// gives "keyRing", so the scaffolder looks for a type such as KMSKeyRingRef
// and names the field keyRingRef rather than keyRingsRef.
//
// It handles only regular plurals (-s, -es, -ies) and leaves words such as
// "status" or "analysis" alone, to avoid false positives. Its answer is only
// ever a guess: whether or not referenceTo finds a matching type, it queues an
// entry in judgement_queue.yaml for a reviewer to confirm.
func Singular(s string) string {
	switch {
	case strings.HasSuffix(s, "ies") && len(s) > 4:
		// policies -> policy
		return s[:len(s)-3] + "y"
	case strings.HasSuffix(s, "sses"), strings.HasSuffix(s, "shes"),
		strings.HasSuffix(s, "ches"), strings.HasSuffix(s, "xes"),
		strings.HasSuffix(s, "zes"):
		// addresses -> address
		return s[:len(s)-2]
	case strings.HasSuffix(s, "ss"), strings.HasSuffix(s, "us"),
		strings.HasSuffix(s, "is"):
		// access, status, analysis: not plurals at all.
		return s
	case strings.HasSuffix(s, "s") && len(s) > 2:
		// subnetworks -> subnetwork
		return s[:len(s)-1]
	}
	// No rule matched: s is already singular, or an irregular plural these
	// rules do not cover.
	klog.Infof("Singular: %q matched no plural rule; returning it unchanged", s)
	return s
}
