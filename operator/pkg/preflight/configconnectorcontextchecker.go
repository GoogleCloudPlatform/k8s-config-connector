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

	ctrl "sigs.k8s.io/controller-runtime"
	"sigs.k8s.io/kubebuilder-declarative-pattern/pkg/patterns/declarative"
)

var (
	clog = ctrl.Log.WithName("ConfigConnectorContextChecker")
)

type ConfigConnectorContextChecker struct {
}

func NewConfigConnectorContextChecker() *ConfigConnectorContextChecker {
	return &ConfigConnectorContextChecker{}
}

func (c *ConfigConnectorContextChecker) Preflight(_ context.Context, o declarative.DeclarativeObject) error {
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
