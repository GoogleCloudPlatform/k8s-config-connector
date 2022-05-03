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

package jitter

import (
	"time"

	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/k8s"

	"k8s.io/apimachinery/pkg/util/wait"
)

// GenerateJitteredReenqueuePeriod() returns a wait duration to reenqueue the request between
// 1/2 * MeanReconcileReenqueuePeriod and 3/2 * MeanReconcileReenqueuePeriod (not inclusive of
// upper bound). The mean duration to reenqueue is MeanReconcileReenqueuePeriod.
func GenerateJitteredReenqueuePeriod() time.Duration {
	return wait.Jitter(k8s.MeanReconcileReenqueuePeriod/2, k8s.JitterFactor)
}
