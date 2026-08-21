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

package preflight

import (
	"context"
	"fmt"

	corev1beta1 "github.com/GoogleCloudPlatform/k8s-config-connector/operator/pkg/apis/core/v1beta1"
	ctrl "sigs.k8s.io/controller-runtime"
	"sigs.k8s.io/controller-runtime/pkg/client"
	"sigs.k8s.io/kubebuilder-declarative-pattern/pkg/patterns/declarative"
)

var (
	nlog = ctrl.Log.WithName("NameChecker")
)

type NameChecker struct {
	allowedName string
	client      client.Client
}

func NewNameChecker(client client.Client, allowedName string) *NameChecker {
	return &NameChecker{allowedName: allowedName, client: client}
}

func (n *NameChecker) Preflight(_ context.Context, o declarative.DeclarativeObject) error {
	nlog.Info("preflight check before reconciling the object", "kind", o.GetObjectKind().GroupVersionKind().Kind, "name", o.GetName(), "namespace", o.GetNamespace())
	if o.GetName() != n.allowedName {
		switch v := o.(type) {
		case *corev1beta1.ConfigConnector:
			return fmt.Errorf("the only allowed name for ConfigConnector object is '%v'. The name restriction is required to ensure that there is only one ConfigConnector instance in your cluster", n.allowedName)
		case *corev1beta1.ConfigConnectorContext:
			return fmt.Errorf("the only allowed name for ConfigConnectorContext object is '%v'. The name restriction is required to ensure that there is only one ConfigConnectorContext instance in your namespace", n.allowedName)
		default:
			return fmt.Errorf("unrecongized type %v", v)
		}
	}
	return nil
}
