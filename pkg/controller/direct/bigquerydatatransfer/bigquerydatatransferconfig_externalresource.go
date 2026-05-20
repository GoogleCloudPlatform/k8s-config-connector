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

package bigquerydatatransfer

import (
	"fmt"
	"strings"

	krm "github.com/GoogleCloudPlatform/k8s-config-connector/apis/bigquerydatatransfer/v1beta1"
)

type BigQueryDataTransferConfigIdentity = krm.BigQueryDataTransferConfigIdentity

// parseServiceGeneratedIDFromName extracts the service generated UUID from the name field of the resource. e.g. "projects/{project_id}/locations/{region}/transferConfigs/{config_id}"
func parseServiceGeneratedIDFromName(s string) (string, error) {
	tokens := strings.Split(s, "/")
	if len(tokens) != 6 || tokens[0] != "projects" || tokens[2] != "locations" || tokens[4] != "transferConfigs" {
		return "", fmt.Errorf("service generated name should have format projects/<project>/locations/<location>/transferConfigs/<transferConfigID>, got %s", s)
	}
	return tokens[5], nil
}
