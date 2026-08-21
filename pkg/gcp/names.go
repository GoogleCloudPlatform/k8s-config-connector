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

package gcp

import (
	"fmt"
	"strings"
)

func FormatBigtableInstanceName(projectID string, name string) string {
	return fmt.Sprintf("projects/%v/instances/%v", projectID, name)
}

func FormatBigtableClusterName(projectID, instanceName, clusterName string) string {
	return fmt.Sprintf("projects/%v/instances/%v/clusters/%v", projectID, instanceName, clusterName)
}

func FormatFullProjectID(projectID string) string {
	return fmt.Sprintf("projects/%v", projectID)
}

func FormatComputeNetworkName(projectID, name string) string {
	return fmt.Sprintf("projects/%v/global/networks/%v", projectID, name)
}

func FormatComputeRouterName(projectID, region, name string) string {
	return fmt.Sprintf("projects/%v/regions/%v/routers/%v", projectID, region, name)
}

func FormatComputeSubNetworkName(projectID, region, name string) string {
	return fmt.Sprintf("projects/%v/regions/%v/subnetworks/%v", projectID, region, name)
}

func FormatGlobalComputeSslCertificateName(projectID, name string) string {
	return fmt.Sprintf("projects/%v/global/sslCertificates/%v", projectID, name)
}

func FormatPubSubTopicName(projectID, name string) string {
	return fmt.Sprintf("projects/%v/topics/%v", projectID, name)
}

func FormatInstanceLocationName(projectID, name string) string {
	return fmt.Sprintf("projects/%v/locations/%v", projectID, name)
}

func FullResourceNameToShortName(fullName string) string {
	idx := strings.LastIndexAny(fullName, "/")
	if idx == -1 {
		return fullName
	}
	return fullName[idx+1:]
}
