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

package parameters

import (
	"net/http"

	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/cli/cmd/commonparams"
)

type Parameters struct {
	IAMFormat               string
	FilterDeletedIAMMembers bool

	// GCPAccessToken is the (optional) static authentication token to use for GCP authentication.
	GCPAccessToken string

	Output         string
	ResourceFormat string
	URI            string
	Verbose        bool

	// HTTPClient allows for overriding the default HTTP Client
	HTTPClient *http.Client
}

func Validate(p *Parameters) error {
	if err := commonparams.ValidateIAMFormat(p.IAMFormat); err != nil {
		return err
	}

	return commonparams.ValidateResourceFormat(p.ResourceFormat, p.IAMFormat)
}
