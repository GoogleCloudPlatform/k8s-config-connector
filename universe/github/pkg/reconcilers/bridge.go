// Copyright 2023 Google LLC
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

package reconcilers

import (
	"context"

	ctrl "sigs.k8s.io/controller-runtime"
	"sigs.k8s.io/controller-runtime/pkg/client"
	"sigs.k8s.io/controller-runtime/pkg/event"
	"sigs.k8s.io/controller-runtime/pkg/handler"
	"sigs.k8s.io/controller-runtime/pkg/manager"
	"sigs.k8s.io/controller-runtime/pkg/source"

	"github.com/GoogleCloudPlatform/k8s-config-connector/universe/github/pkg/githubapi"
)

type ReconcilerBridge struct {
	client client.Client
	mgr    manager.Manager

	impl           Reconciler
	externalClient *githubapi.Client
	notify         chan event.GenericEvent
}

func NewReconcilerBridge(impl Reconciler, externalClient *githubapi.Client) *ReconcilerBridge {
	return &ReconcilerBridge{
		impl:           impl,
		externalClient: externalClient,
		notify:         make(chan event.GenericEvent),
	}
}

func (r *ReconcilerBridge) SetupWithManager(mgr manager.Manager) error {
	r.client = mgr.GetClient()
	r.mgr = mgr

	watchInfo := r.impl.GetWatchInfo()
	return ctrl.NewControllerManagedBy(mgr).
		For(watchInfo.Primary).
		WatchesRawSource(&source.Channel{Source: r.notify}, &handler.EnqueueRequestForObject{}).
		Complete(r)
}

func (r *ReconcilerBridge) Reconcile(ctx context.Context, req ctrl.Request) (ctrl.Result, error) {
	reconcileReq := &ReconcileRequest{
		ID:             req,
		Client:         r.client,
		ExternalClient: r.externalClient,
	}
	op := r.impl.NewOp(reconcileReq)
	_, err := op.Reconcile(ctx)
	if err != nil {
		return ctrl.Result{}, err
	}
	return ctrl.Result{}, nil
}
