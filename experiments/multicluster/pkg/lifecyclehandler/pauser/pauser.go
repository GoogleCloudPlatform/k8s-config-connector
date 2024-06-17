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

package pauser

import (
	"context"

	"github.com/go-logr/logr"
	apierrors "k8s.io/apimachinery/pkg/api/errors"
	"k8s.io/apimachinery/pkg/types"
	"sigs.k8s.io/controller-runtime/pkg/client"

	"github.com/GoogleCloudPlatform/k8s-config-connector/experiments/multicluster/pkg/lifecyclehandler"
	corev1beta1 "github.com/GoogleCloudPlatform/k8s-config-connector/operator/pkg/apis/core/v1beta1"
	k8s "github.com/GoogleCloudPlatform/k8s-config-connector/operator/pkg/k8s"
)

func New(c PauserConfig) *Pauser {
	return &Pauser{
		client:    c.Client,
		namespace: c.Namespace,
		identity:  c.Identity,
		log:       c.Log,
	}
}

type PauserConfig struct {
	Client    client.Client
	Identity  string
	Namespace string
	Log       logr.Logger
}

var _ lifecyclehandler.LifecycleHandler = &Pauser{}

type Pauser struct {
	client    client.Client
	identity  string
	namespace string
	log       logr.Logger
}

func (p *Pauser) OnStartedLeading(ctx context.Context) error {
	p.log.Info("started leading")
	// TODO: retry in case there is a conflict
	return p.unPause(ctx)
}

func (p *Pauser) OnStoppedLeading(ctx context.Context) error {
	p.log.Info("stopped leading")
	// TODO: retry in case there is a conflict
	return p.pause(ctx)
}

func (p *Pauser) OnNewLeader(ctx context.Context, leaderID string) error {
	if leaderID == p.identity {
		// I just got the lock
		return nil
	}
	p.log.Info("new leader observed", "leaderID", leaderID)
	// TODO: retry in case there is a conflict
	return p.pause(ctx)
}

func (p *Pauser) OnStopping() error {
	p.log.Info("stopping")
	return p.pause(context.Background())
}

func (p *Pauser) pause(ctx context.Context) error {
	p.log.Info("pausing KCC", "namespace", p.namespace)
	// fetch CCC
	ccc, err := getConfigConnectorContext(ctx, p.client, types.NamespacedName{
		Name:      k8s.ConfigConnectorContextAllowedName,
		Namespace: p.namespace,
	})
	if err != nil {
		if apierrors.IsNotFound(err) {
			p.log.Info("ConfigConnectorContext object not found, skipping 'OnStartedLeading' operation", "namespace", p.namespace)
			return err
		}
		p.log.Error(err, "failed to fetch ccc object")
		return err
	}
	p.log.V(4).Info("fetched ccc object", "ConfigConnectorContext object", ccc)

	// ensure CCC is paused
	ccc.Spec.Actuation = corev1beta1.Paused
	if err := updateConfigConnectorContext(ctx, p.client, ccc); err != nil {
		p.log.Error(err, "failed to update ccc object")
		return err
	}
	p.log.V(4).Info("updated ccc object", "ConfigConnectorContext object", ccc)

	p.log.Info("KCC paused", "namespace", p.namespace)
	return nil
}

func (p *Pauser) unPause(ctx context.Context) error {
	p.log.Info("un-pausing KCC", "namespace", p.namespace)
	// fetch CCC
	ccc, err := getConfigConnectorContext(ctx, p.client, types.NamespacedName{
		Name:      k8s.ConfigConnectorContextAllowedName,
		Namespace: p.namespace,
	})
	if err != nil {
		if apierrors.IsNotFound(err) {
			p.log.Info("ConfigConnectorContext object not found, skipping 'OnStartedLeading' operation", "namespace", p.namespace)
			return err
		}
		p.log.Error(err, "failed to fetch ccc object")
		return err
	}
	p.log.V(4).Info("fetched ccc object", "ConfigConnectorContext object", ccc)

	// ensure CCC is unpaused
	ccc.Spec.Actuation = corev1beta1.Reconciling
	if err := updateConfigConnectorContext(ctx, p.client, ccc); err != nil {
		p.log.Error(err, "failed to update ccc object")
		return err
	}
	p.log.V(4).Info("updated ccc object", "ConfigConnectorContext object", ccc)

	p.log.Info("KCC un-paused", "namespace", p.namespace)
	return nil
}

func getConfigConnectorContext(ctx context.Context, client client.Client, nn types.NamespacedName) (*corev1beta1.ConfigConnectorContext, error) {
	ccc := &corev1beta1.ConfigConnectorContext{}
	if err := client.Get(ctx, nn, ccc); err != nil {
		return nil, err
	}
	return ccc, nil
}

func updateConfigConnectorContext(ctx context.Context, client client.Client, ccc *corev1beta1.ConfigConnectorContext) error {
	if err := client.Update(ctx, ccc); err != nil {
		return err
	}
	return nil
}
