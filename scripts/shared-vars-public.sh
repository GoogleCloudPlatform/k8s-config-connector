#!/usr/bin/env bash
# Copyright 2022 Google LLC
#
# Licensed under the Apache License, Version 2.0 (the "License");
# you may not use this file except in compliance with the License.
# You may obtain a copy of the License at
#
#      http://www.apache.org/licenses/LICENSE-2.0
#
# Unless required by applicable law or agreed to in writing, software
# distributed under the License is distributed on an "AS IS" BASIS,
# WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
# See the License for the specific language governing permissions and
# limitations under the License.


# Paths
REPO_ROOT="$(git rev-parse --show-toplevel)"

# general purpose folders
BIN_DIR=bin

# binary names
CONFIG_CONNECTOR_BINARY_NAME=config-connector

# dependency versions
KUBEBUILDER_VERSION=2.3.1
KUBEAPISERVER_VERSION=1.21.0
KUSTOMIZE_VERSION=3.5.4

# Supported GCP services API endpoints in Config Connector
SUPPORTED_SERVICES=(
  accesscontextmanager.googleapis.com
  aiplatform.googleapis.com
  alloydb.googleapis.com
  anthos.googleapis.com
  anthosaudit.googleapis.com
  anthosgke.googleapis.com
  apigee.googleapis.com
  apikeys.googleapis.com
  appengine.googleapis.com
  artifactregistry.googleapis.com
  backupdr.googleapis.com
  bigquery.googleapis.com
  bigqueryconnection.googleapis.com
  bigquerydatatransfer.googleapis.com
  bigtableadmin.googleapis.com
  billingbudgets.googleapis.com
  binaryauthorization.googleapis.com
  certificatemanager.googleapis.com
  cloudasset.googleapis.com
  cloudbilling.googleapis.com
  cloudbuild.googleapis.com
  cloudidentity.googleapis.com
  cloudfunctions.googleapis.com
  cloudkms.googleapis.com
  cloudresourcemanager.googleapis.com
  cloudscheduler.googleapis.com
  compute.googleapis.com
  connectgateway.googleapis.com
  container.googleapis.com
  containeranalysis.googleapis.com
  datacatalog.googleapis.com
  dataflow.googleapis.com
  datafusion.googleapis.com
  dataproc.googleapis.com
  datastream.googleapis.com
  discoveryengine.googleapis.com
  dlp.googleapis.com
  dns.googleapis.com
  edgenetwork.googleapis.com
  edgecontainer.googleapis.com
  eventarc.googleapis.com
  file.googleapis.com
  firestore.googleapis.com
  gkebackup.googleapis.com
  gkeconnect.googleapis.com
  gkehub.googleapis.com
  gkemulticloud.googleapis.com
  iap.googleapis.com
  iam.googleapis.com
  identitytoolkit.googleapis.com
  ids.googleapis.com
  krmapihosting.googleapis.com
  logging.googleapis.com
  managedkafka.googleapis.com
  monitoring.googleapis.com
  memcache.googleapis.com
  networkconnectivity.googleapis.com
  networksecurity.googleapis.com
  networkservices.googleapis.com
  osconfig.googleapis.com
  opsconfigmonitoring.googleapis.com
  privateca.googleapis.com
  privilegedaccessmanager.googleapis.com
  pubsub.googleapis.com
  pubsublite.googleapis.com
  recaptchaenterprise.googleapis.com
  redis.googleapis.com
  run.googleapis.com
  secretmanager.googleapis.com
  securesourcemanager.googleapis.com
  servicedirectory.googleapis.com
  servicenetworking.googleapis.com
  serviceusage.googleapis.com
  sourcerepo.googleapis.com
  stackdriver.googleapis.com
  spanner.googleapis.com
  sqladmin.googleapis.com
  storagetransfer.googleapis.com
  vpcaccess.googleapis.com
  vmwareengine.googleapis.com
  workstations.googleapis.com
)

# Regex used to match long running tests cases (10m+ runtime). Any new
# long-running tests should be added to this regex in alphabetical order.
LONG_RUNNING_CRUD_TESTS_REGEX="basicalloydbbackup|\
basicalloydbinstance|\
basicalloydbsecondarycluster|\
basicalloydbsecondaryinstance|\
cidrconnector|\
cloudidsendpoint|\
configcontrollerinstance|\
containercluster|\
containernodepool|\
databasealloydbuser|\
datafusioninstance|\
filestorebackup|\
filestoreinstance|\
fullalloydbbackup|\
fullalloydbcluster|\
gkehubfeaturemembership|\
gkehubmembership|\
iamalloydbuser|\
memcacheinstance|\
postgresinstance|\
readalloydbinstance|\
redisinstance|\
removedefaultnodepool|\
restorebackupalloydbcluster|\
securesourcemanagerinstancebasic|\
securesourcemanagerinstancecmek|\
securesourcemanagerrepositorybasic|\
securesourcemanagerrepositoryfull|\
sqlinstanceencryptionkey|\
subnetconnector|\
vertexaidatasetencryptionkey|\
vertexaiendpointencryptionkey|\
zonalalloydbinstance"
