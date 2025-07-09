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

package webhook

import (
	"fmt"

	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/webhook"
)

func GetTestCommonWebhookConfigs() ([]webhook.Config, error) {
	whCfgs, err := webhook.GetCommonWebhookConfigs()
	if err != nil {
		return nil, fmt.Errorf("error getting common wehbook configs: %w", err)
	}
	res := make([]webhook.Config, 0)
	for _, config := range whCfgs {
		// deny-immutable-field-updates webhook currently cannot work with envtest
		// because updates from controller cannot be distinguished by service account
		// TODO: figure out a way to enable it in envtest

		// res = append(res, config)

		if config.Name != "deny-immutable-field-updates.cnrm.cloud.google.com" {
			res = append(res, config)
		}
	}
	return res, nil
}
