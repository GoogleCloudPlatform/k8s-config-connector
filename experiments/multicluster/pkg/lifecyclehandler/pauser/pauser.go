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
	meta "k8s.io/apimachinery/pkg/api/meta"
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
		log:       c.Log,
	}
}

type PauserConfig struct {
	Client    client.Client
	Namespace string
	Log       logr.Logger
}

var _ lifecyclehandler.LifecycleHandler = &Pauser{}

type Pauser struct {
	client    client.Client
	namespace string
	log       logr.Logger
}

func (p *Pauser) OnNewLeader(ctx context.Context, leaderID string, isLeader bool) error {
	p.log.Info("new leader observed", "leaderID", leaderID)
	// TODO: retry in case there is a conflict
	if isLeader {
		return p.setActuationMode(ctx, corev1beta1.Reconciling)
	}
	return p.setActuationMode(ctx, corev1beta1.Paused)
}

func (p *Pauser) OnStopping() error {
	p.log.Info("stopping")
	return p.setActuationMode(context.Background(), corev1beta1.Paused)
}

func (p *Pauser) setActuationMode(ctx context.Context, mode corev1beta1.ActuationMode) error {
	p.log.Info("setting actuation mode", "mode", mode, "namespace", p.namespace)

	// Fetch ConfigConnectorContext (CCC) object
	ccc, err := getConfigConnectorContext(ctx, p.client, types.NamespacedName{
		Name:      k8s.ConfigConnectorContextAllowedName,
		Namespace: p.namespace,
	})
	if err != nil {
		if apierrors.IsNotFound(err) {
			p.log.Info("ConfigConnectorContext object not found; skipping 'OnStartedLeading' operation", "namespace", p.namespace)
			return nil
		}
		p.log.Error(err, "failed to fetch ConfigConnectorContext object")
		return err
	}
	p.log.V(4).Info("fetched ConfigConnectorContext object", "object", ccc)

	// Update the actuation mode in the CCC object
	ccc.Spec.Actuation = mode
	if err := updateConfigConnectorContext(ctx, p.client, ccc); err != nil {
		p.log.Error(err, "failed to update ConfigConnectorContext object")
		return err
	}
	p.log.V(4).Info("updated ConfigConnectorContext object", "object", ccc)

	p.log.Info("actuation mode successfully set", "mode", mode, "namespace", p.namespace)
	return nil
}

func getConfigConnectorContext(ctx context.Context, client client.Client, nn types.NamespacedName) (*corev1beta1.ConfigConnectorContext, error) {
	ccc := &corev1beta1.ConfigConnectorContext{}
	if err := client.Get(ctx, nn, ccc); err != nil {
		return nil, err
	}
	return ccc, nil
}

func updateConfigConnectorContext(ctx context.Context, c client.Client, ccc *corev1beta1.ConfigConnectorContext) error {
	// ensure that managedFields is nil
	if accessor, err := meta.Accessor(ccc); err == nil {
		accessor.SetManagedFields(nil)
	} else {
		return err
	}

	opts := []client.PatchOption{client.ForceOwnership}
	opts = append(opts, client.FieldOwner(lifecyclehandler.MultiClusterFieldManager))
	if err := c.Patch(ctx, ccc, client.Apply, opts...); err != nil {
		return err
	}
	return nil
}
