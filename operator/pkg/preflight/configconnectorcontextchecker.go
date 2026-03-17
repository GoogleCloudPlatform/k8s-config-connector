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
	"regexp"

	corev1beta1 "github.com/GoogleCloudPlatform/k8s-config-connector/operator/pkg/apis/core/v1beta1"
	"github.com/GoogleCloudPlatform/k8s-config-connector/operator/pkg/k8s"

	"k8s.io/apimachinery/pkg/types"
	ctrl "sigs.k8s.io/controller-runtime"
	"sigs.k8s.io/controller-runtime/pkg/client"
	"sigs.k8s.io/kubebuilder-declarative-pattern/pkg/patterns/declarative"
)

var (
	clog = ctrl.Log.WithName("ConfigConnectorContextChecker")
)

type ConfigConnectorContextChecker struct {
	client client.Client
}

func NewConfigConnectorContextChecker(client client.Client) *ConfigConnectorContextChecker {
	return &ConfigConnectorContextChecker{
		client: client,
	}
}

func (c *ConfigConnectorContextChecker) Preflight(ctx context.Context, o declarative.DeclarativeObject) error {
	clog.Info("preflight check before reconciling the object", "kind", o.GetObjectKind().GroupVersionKind().Kind, "name", o.GetName(), "namespace", o.GetNamespace())

	ccc, ok := o.(*corev1beta1.ConfigConnectorContext)
	if !ok {
		return fmt.Errorf("expected the resource to be a ConfigConnectorContext, but it was not. Object: %v", o)
	}

	if ccc.GetRequestProjectPolicy() != k8s.BillingProjectPolicy && ccc.Spec.BillingProject != "" {
		return fmt.Errorf("spec.billingProject cannot be set if spec.requestProjectPolicy is not set to %v", k8s.BillingProjectPolicy)
	}

	if ccc.GetRequestProjectPolicy() == k8s.BillingProjectPolicy && ccc.Spec.BillingProject == "" {
		return fmt.Errorf("spec.billingProject must be set if spec.requestProjectPolicy is set to %v", k8s.BillingProjectPolicy)
	}

	if err := validateGSAFormat(ccc.Spec.GoogleServiceAccount); err != nil {
		return err
	}

	// Validate mode consistency with ConfigConnector
	cc := &corev1beta1.ConfigConnector{}
	if c.client != nil {
		if err := c.client.Get(ctx, types.NamespacedName{Name: "configconnector.core.cnrm.cloud.google.com"}, cc); err == nil {
			if err := validateResourceSettingsMode(cc, ccc); err != nil {
				return err
			}
		}
	}

	return nil
}

func validateResourceSettingsMode(cc *corev1beta1.ConfigConnector, ccc *corev1beta1.ConfigConnectorContext) error {
	if cc.Spec.Experiments == nil || cc.Spec.Experiments.ResourceSettings == nil {
		return nil
	}
	if ccc.Spec.Experiments == nil || ccc.Spec.Experiments.ResourceSettings == nil {
		return nil
	}
	ccSettings := cc.Spec.Experiments.ResourceSettings
	cccSettings := ccc.Spec.Experiments.ResourceSettings

	ccInclusive := ccSettings.Enabled != nil && *ccSettings.Enabled
	cccInclusive := cccSettings.Enabled != nil && *cccSettings.Enabled

	if ccInclusive != cccInclusive {
		return fmt.Errorf("conflict: ConfigConnector and ConfigConnectorContext cannot mix inclusive (enabled: true) and exclusive (enabled: false) modes")
	}
	return nil
}

func validateGSAFormat(gsa string) error {
	if gsa == "" { // GoogleServiceAccount is a required field. We do not need to fail here.
		return nil
	}
	validGSAPattern := `^[A-Za-z0-9._%+\-]+@[a-z0-9.\-]+\.gserviceaccount.com$`
	emailRegex := regexp.MustCompile(validGSAPattern)
	if !emailRegex.MatchString(gsa) {
		return fmt.Errorf("invalid GoogleServiceAccount format for %q", gsa)
	}
	return nil
}
