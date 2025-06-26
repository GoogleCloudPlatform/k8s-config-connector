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

package lifecyclehandler

import (
	"context"
)

const MultiClusterFieldManager = "multicluster-controller-manager"

type LifecycleHandler interface {
	// OnNewLeader is triggered when a LeaderElector detects a change in leadership.
	// This includes the first leader detected when the LeaderElector starts.
	// leaderID identifies the new leader.
	// isLeader indicates if the new leader is the current LeaderElector.
	OnNewLeader(ctx context.Context, leaderID string, isLeader bool) error

	// OnStopping is triggered when a LeaderElector is in the process of stopping.
	OnStopping() error
}
