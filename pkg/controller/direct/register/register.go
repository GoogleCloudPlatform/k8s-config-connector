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

package register

import (
	_ "github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct/alloydb"
	_ "github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct/apikeys"
	_ "github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct/bigqueryanalyticshub"
	_ "github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct/bigqueryconnection"
	_ "github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct/bigquerydataset"
	_ "github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct/bigquerydatatransfer"
	_ "github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct/certificatemanager"
	_ "github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct/cloudbuild"
	_ "github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct/compute/firewallpolicyrule"
	_ "github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct/compute/forwardingrule"
	_ "github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct/dataflow"
	_ "github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct/dataform"
	_ "github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct/discoveryengine"
	_ "github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct/firestore"
	_ "github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct/gkehub"
	_ "github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct/kms/autokeyconfig"
	_ "github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct/kms/keyhandle"
	_ "github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct/logging"
	_ "github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct/monitoring"
	_ "github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct/networkconnectivity"
	_ "github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct/privateca"
	_ "github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct/privilegedaccessmanager"
	_ "github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct/redis/cluster"
	_ "github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct/resourcemanager"
	_ "github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct/secretmanager"
	_ "github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct/securesourcemanager"
	_ "github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct/sql"
	_ "github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct/workstations"
)
