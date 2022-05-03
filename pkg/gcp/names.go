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

func FormatBigtableInstanceName(projectId string, name string) string {
	return fmt.Sprintf("projects/%v/instances/%v", projectId, name)
}

func FormatBigtableClusterName(projectId, instanceName, clusterName string) string {
	return fmt.Sprintf("projects/%v/instances/%v/clusters/%v", projectId, instanceName, clusterName)
}

func FormatFullProjectID(projectId string) string {
	return fmt.Sprintf("projects/%v", projectId)
}

func FormatComputeNetworkName(projectId, name string) string {
	return fmt.Sprintf("projects/%v/global/networks/%v", projectId, name)
}

func FormatComputeRouterName(projectId, region, name string) string {
	return fmt.Sprintf("projects/%v/regions/%v/routers/%v", projectId, region, name)
}

func FormatComputeSubNetworkName(projectId, region, name string) string {
	return fmt.Sprintf("projects/%v/regions/%v/subnetworks/%v", projectId, region, name)
}

func FormatGlobalComputeSslCertificateName(projectId, name string) string {
	return fmt.Sprintf("projects/%v/global/sslCertificates/%v", projectId, name)
}

func FormatPubSubTopicName(projectId, name string) string {
	return fmt.Sprintf("projects/%v/topics/%v", projectId, name)
}

func FormatInstanceLocationName(projectId, name string) string {
	return fmt.Sprintf("projects/%v/locations/%v", projectId, name)
}

func FullResourceNameToShortName(fullName string) string {
	idx := strings.LastIndexAny(fullName, "/")
	if idx == -1 {
		return fullName
	}
	return fullName[idx+1:]
}
