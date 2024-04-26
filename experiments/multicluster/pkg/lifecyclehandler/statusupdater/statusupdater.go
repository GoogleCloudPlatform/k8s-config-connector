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

package statusupdater

import (
	"context"
	"fmt"
	"time"

	"k8s.io/apimachinery/pkg/types"
	"k8s.io/client-go/tools/leaderelection"
	"sigs.k8s.io/controller-runtime/pkg/client"

	"github.com/GoogleCloudPlatform/k8s-config-connector/experiments/multicluster/pkg/lifecyclehandler"
	"github.com/GoogleCloudPlatform/k8s-config-connector/experiments/multicluster/pkg/util"
)

func New(c Config) *StatusUpdater {
	return &StatusUpdater{
		client:   c.Client,
		nn:       c.NN,
		identity: c.Identity,
	}
}

type Config struct {
	Client   client.Client
	NN       types.NamespacedName
	Identity string
}

type StatusUpdater struct {
	client   client.Client
	nn       types.NamespacedName
	identity string
}

var _ lifecyclehandler.LifecycleHandler = &StatusUpdater{}

func (s *StatusUpdater) Callbacks(ctx context.Context) leaderelection.LeaderCallbacks {
	return leaderelection.LeaderCallbacks{
		OnStartedLeading: func(ctx context.Context) {
			s.OnStartedLeadingFunc()(ctx)
		},
		OnStoppedLeading: func() {
			s.OnStoppedLeadingFunc()(context.Background()) // need to use a new context so that the function calls wonn't be cancelled upon original ctx cancellation.
		},
		OnNewLeader: func(leaderID string) {
			s.OnNewLeaderFunc()(ctx, leaderID)
		},
	}
}

func (s *StatusUpdater) OnStartedLeadingFunc() func(ctx context.Context) {
	return func(ctx context.Context) {
		s.setIsLeader(ctx, true)
	}
}

func (s *StatusUpdater) OnStoppedLeadingFunc() func(ctx context.Context) {
	return func(ctx context.Context) {
		s.setIsLeader(ctx, false)
	}
}

func (s *StatusUpdater) OnNewLeaderFunc() func(ctx context.Context, leaderID string) {
	return func(ctx context.Context, leaderID string) {
		s.setIsLeader(ctx, leaderID == s.identity)
	}
}

func (s *StatusUpdater) setIsLeader(ctx context.Context, isLeader bool) error {
	mcl, err := util.GetMultiClusterLease(ctx, s.client, s.nn)
	if err != nil {
		return fmt.Errorf("error getting MultiClusterLease: %w", err)
	}
	mcl.Status.IsLeader = isLeader
	mcl.Status.LastObservedTime = time.Now().Format(time.RFC3339)

	err = s.client.Status().Update(ctx, mcl)
	if err != nil {
		return fmt.Errorf("error updating MultiClusterLease status: %w", err)
	}
	return nil
}
