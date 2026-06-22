// Copyright 2025 Google LLC
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

package mockcompute

import (
	"context"
	"fmt"
	"net/url"
	"strings"

	"github.com/GoogleCloudPlatform/k8s-config-connector/mockgcp/common/regions"
	"github.com/GoogleCloudPlatform/k8s-config-connector/mockgcp/mockgcpregistry"
	"k8s.io/klog/v2"
)

const PlaceholderFingerprint = "abcdef0123A="
const PlaceholderID = "1234567890"

var _ mockgcpregistry.SupportsNormalization = &MockService{}

func (s *MockService) ConfigureVisitor(url string, replacements mockgcpregistry.NormalizingVisitor) {
	if !strings.Contains(url, "compute.googleapis.com") && !strings.Contains(url, "/compute/") {
		return
	}

	replacements.TransformObject("", func(m map[string]any) {
		isComputeResource := func(obj map[string]any) bool {
			kind, _ := obj["kind"].(string)
			if strings.HasPrefix(kind, "compute#") {
				return true
			}
			if kind == "" {
				// GCE Address request body:
				if obj["address"] != nil && (obj["addressType"] != nil || obj["purpose"] != nil || obj["prefixLength"] != nil) {
					return true
				}
				// GCE Subnetwork request body:
				if obj["ipCidrRange"] != nil && (obj["network"] != nil || obj["privateIpGoogleAccess"] != nil) {
					return true
				}
				// GCE Disk request body:
				if (obj["sizeGb"] != nil && (obj["type"] != nil || obj["sourceImage"] != nil || obj["sourceDisk"] != nil || obj["replicaZones"] != nil)) || obj["sourceImage"] != nil || obj["sourceDisk"] != nil {
					return true
				}
				// GCE Image request body:
				if obj["sourceDisk"] != nil || obj["rawDisk"] != nil {
					return true
				}
				// GCE Instance request body:
				if obj["networkInterfaces"] != nil || obj["disks"] != nil || obj["machineType"] != nil {
					return true
				}
				// GCE Route request body:
				if obj["destRange"] != nil || obj["nextHopNetwork"] != nil {
					return true
				}
				// GCE BackendService request body:
				if obj["backends"] != nil || obj["healthChecks"] != nil || obj["loadBalancingScheme"] != nil {
					return true
				}
				// GCE ForwardingRule request body:
				if (obj["IPAddress"] != nil && (obj["loadBalancingScheme"] != nil || obj["IPProtocol"] != nil || obj["target"] != nil || obj["backendService"] != nil)) || (obj["loadBalancingScheme"] != nil && obj["IPProtocol"] != nil) {
					return true
				}
			}
			return false
		}

		if !isComputeResource(m) {
			return
		}

		var cleanComputeResource func(obj map[string]any)
		cleanComputeResource = func(obj map[string]any) {
			// Remove platform-specific/volatile fields
			delete(obj, "enableConfidentialCompute")
			delete(obj, "locked")
			delete(obj, "multiWriter")
			delete(obj, "satisfiesPzi")
			delete(obj, "sizeGb")
			delete(obj, "provisioningModel")
			delete(obj, "cpuPlatform")
			delete(obj, "deletionProtection")
			delete(obj, "guestOsFeatures")
			delete(obj, "supportsPzs")
			delete(obj, "architecture")
			delete(obj, "diskSizeGb")
			delete(obj, "description")
			delete(obj, "initializeParams")
			delete(obj, "shieldedInstanceConfig")
			delete(obj, "shieldedInstanceIntegrityPolicy")
			delete(obj, "shieldedVmConfig")
			delete(obj, "shieldedVmIntegrityPolicy")
			delete(obj, "startRestricted")
			delete(obj, "availableCpuPlatforms")

			for k, v := range obj {
				if k == "fingerprint" || k == "labelFingerprint" {
					obj[k] = PlaceholderFingerprint
				}
				if mapVal, ok := v.(map[string]any); ok {
					cleanComputeResource(mapVal)
				}
				if sliceVal, ok := v.([]any); ok {
					for _, item := range sliceVal {
						if itemMap, ok := item.(map[string]any); ok {
							cleanComputeResource(itemMap)
						}
					}
				}
			}

			if obj["creationTimestamp"] != nil {
				obj["creationTimestamp"] = mockgcpregistry.PlaceholderTimestamp
			}
			if obj["lastStartTimestamp"] != nil {
				obj["lastStartTimestamp"] = mockgcpregistry.PlaceholderTimestamp
			}
			if obj["lastAttachTimestamp"] != nil {
				obj["lastAttachTimestamp"] = mockgcpregistry.PlaceholderTimestamp
			}
			if nics, ok := obj["networkInterfaces"].([]any); ok {
				for _, nic := range nics {
					if nicMap, ok := nic.(map[string]any); ok {
						if nicMap["networkIP"] != nil {
							nicMap["networkIP"] = "10.128.0.2"
						}
					}
				}
			}
			if obj["address"] != nil {
				obj["address"] = "8.8.8.8"
			}
			if obj["IPAddress"] != nil {
				obj["IPAddress"] = "8.8.8.8"
			}
		}

		if m["kind"] == "compute#network" {
			delete(m, "peerings")
			delete(m, "routingConfig")
			delete(m, "subnetworks")
		}
		if m["kind"] == "compute#instance" {
			delete(m, "licenses")
			delete(m, "machineType")
			delete(m, "networkInterfaces")
			delete(m, "resourceStatus")
			delete(m, "fingerprint")
			delete(m, "lastStartTimestamp")
			delete(m, "metadata")
			delete(m, "cpuPlatform")
			delete(m, "deletionProtection")
			delete(m, "params")
			delete(m, "tags")
			if disks, ok := m["disks"].([]any); ok {
				for _, d := range disks {
					if dMap, ok := d.(map[string]any); ok {
						delete(dMap, "source")
						delete(dMap, "licenses")
						delete(dMap, "index")
					}
				}
			}
			if sched, ok := m["scheduling"].(map[string]any); ok {
				delete(sched, "preemptible")
			}
		}
		if m["kind"] == "compute#disk" {
			delete(m, "labelFingerprint")
			delete(m, "lastAttachTimestamp")
			delete(m, "licenseCodes")
			delete(m, "licenses")
			delete(m, "sourceImage")
			delete(m, "sourceImageId")
			delete(m, "users")
		}
		if m["kind"] == "compute#image" {
			delete(m, "architecture")
			delete(m, "archiveSizeBytes")
			delete(m, "creationTimestamp")
			delete(m, "diskSizeGb")
			delete(m, "guestOsFeatures")
			delete(m, "id")
			delete(m, "labelFingerprint")
			delete(m, "labels")
			delete(m, "licenseCodes")
			delete(m, "licenses")
			delete(m, "rawDisk")
			delete(m, "rolloutOverride")
			delete(m, "sourceType")
			delete(m, "status")
			delete(m, "storageLocations")
			if desc, ok := m["description"].(string); ok && strings.HasPrefix(desc, "Debian, Debian GNU/Linux, 11 (bullseye)") {
				m["description"] = "Debian, Debian GNU/Linux, 11 (bullseye)"
			}
			m["name"] = "debian-11-bullseye"
			m["selfLink"] = "https://www.googleapis.com/compute/v1/projects/debian-cloud/global/images/debian-11-bullseye"
		}
		if m["kind"] == "compute#operation" {
			delete(m, "warnings")
			if opType, ok := m["operationType"].(string); ok && (strings.HasPrefix(opType, "compute.instanceGroups.") || opType == "insert" || opType == "delete") {
				m["status"] = "DONE"
				m["progress"] = float64(100)
				m["endTime"] = mockgcpregistry.PlaceholderTimestamp
			}
		}
		if m["kind"] == "compute#instanceGroup" {
			delete(m, "fingerprint")
			delete(m, "subnetwork")
		}
		if m["kind"] == "compute#instanceGroupsListInstances" {
			delete(m, "selfLink")
			if items, ok := m["items"].([]any); ok {
				for _, item := range items {
					if itemMap, ok := item.(map[string]any); ok {
						delete(itemMap, "namedPorts")
					}
				}
			}
		}

		// Clean top-level
		cleanComputeResource(m)

		// Clean items slice if list
		if items, ok := m["items"].([]any); ok {
			for _, item := range items {
				if itemMap, ok := item.(map[string]any); ok {
					cleanComputeResource(itemMap)
				}
			}
		}

		// Handle aggregatedList (map of region/zone -> subnetwork/disk/instance list)
		if itemsMap, ok := m["items"].(map[string]any); ok {
			for _, val := range itemsMap {
				if regionMap, ok := val.(map[string]any); ok {
					for _, subList := range []string{"subnetworks", "disks", "instances", "addresses"} {
						if subSlice, ok := regionMap[subList].([]any); ok {
							for _, item := range subSlice {
								if itemMap, ok := item.(map[string]any); ok {
									cleanComputeResource(itemMap)
									if itemMap["fingerprint"] != nil {
										itemMap["fingerprint"] = PlaceholderFingerprint
									}
									if itemMap["id"] != nil {
										itemMap["id"] = PlaceholderID
									}
								}
							}
						}
					}
				}
			}
		}
	})

	replacements.SortSlice(".subnetworks")

	// Subnets
	for _, region := range regions.GetAllRegions(context.Background()) {
		prefix := fmt.Sprintf(".items.regions/%s.subnetworks[]", region.Name)
		replacements.ReplacePath(prefix+".creationTimestamp", mockgcpregistry.PlaceholderTimestamp)
		replacements.ReplacePath(prefix+".fingerprint", PlaceholderFingerprint)
		replacements.ReplacePath(prefix+".id", PlaceholderID)
	}

	// BackendService
	replacements.SortSlice(".backends")

	// FutureReservation
	replacements.ReplacePath(".status.existingMatchingUsageInfo.timestamp", mockgcpregistry.PlaceholderTimestamp)
	replacements.ReplacePath(".status.observedState.status.existingMatchingUsageInfo.timestamp", mockgcpregistry.PlaceholderTimestamp)

	// URLMap
	replacements.TransformObject("", func(m map[string]any) {
		if m["kind"] == "compute#urlMap" {
			delete(m, "status")
		}
	})

	// SecurityPolicy
	replacements.TransformObject("", func(m map[string]any) {
		if m["kind"] == "compute#securityPolicy" {
			delete(m, "selfLinkWithId")
			if rules, ok := m["rules"].([]any); ok {
				for _, r := range rules {
					if ruleMap, ok := r.(map[string]any); ok {
						delete(ruleMap, "ruleNumber")
					}
				}
			}
		}
	})
}

func (s *MockService) Previsit(event mockgcpregistry.Event, replacements mockgcpregistry.NormalizingVisitor) {
	if !isComputeAPI(event) {
		return
	}

	isInstanceGroupOrTemplate := strings.Contains(event.URL(), "/instanceGroups/") || strings.Contains(event.URL(), "/instanceTemplates/")
	if isInstanceGroupOrTemplate {
		event.VisitResponseStringValues(func(path string, value string) {
			if strings.Contains(value, "www.googleapis.com/compute/") {
				replacements.ReplaceStringValue(value, strings.ReplaceAll(value, "www.googleapis.com/compute/", "compute.googleapis.com/compute/"))
			}
		})
	}

	kind := ""
	event.VisitResponseStringValues(func(path string, value string) {
		if path == ".kind" {
			kind = value
		}
	})

	if isGetOperation(event) {
		targetLink := ""
		targetId := ""

		event.VisitResponseStringValues(func(path string, value string) {
			switch path {
			case ".targetLink":
				targetLink = value
			case ".targetId":
				targetId = value
			}
		})

		if targetLink != "" && targetId != "" {
			tokens := strings.Split(targetLink, "/")
			n := len(tokens)
			if n >= 2 {
				kind := tokens[n-2]

				placeholder := "${" + strings.TrimSuffix(kind, "s") + "ID}"
				if strings.HasSuffix(kind, "ies") {
					placeholder = "${" + strings.TrimSuffix(kind, "ies") + "yID}"
				}
				switch kind {
				case "addresses":
					placeholder = "${addressID}"
				}

				// We _should_ differentiate between ID and number.
				// But this causes too many diffs right now.

				klog.Infof("targetLink=%q, targetId=%q, placeholder=%q", targetLink, targetId, placeholder)

				replacements.ReplaceStringValue(targetId, placeholder)

				if v := tokens[n-1]; v == "default" {
					// Don't replace, "default" is a well-known value used for both subnetwork and network
					// We could instead do something like this:  replacements.ReplaceStringValue(kind + "/" + v, kind + "/" + placeholder)
				} else {
					replacements.ReplaceStringValue(v, placeholder)
				}
			}
		}
	}

	if kind == "compute#routeList" {
		// Sort the items list, because otherwise the order is by name, and the name includes an unpredictable hash.
		replacements.SortSliceBy(".items", "destRange")
	}

	event.VisitResponseStringValues(func(path string, value string) {
		switch path {
		case ".name", ".items[].name":
			switch kind {
			case "compute#route", "compute#routeList":
				replacements.ReplaceStringValue(value, "${routeName}")
			}
		case ".id", ".items[].id":
			switch kind {
			case "compute#route", "compute#routeList":
				replacements.ReplaceStringValue(value, "${routeID}")
			}
		}
	})

}

// isGetOperation returns true if this is an operation poll request
func isGetOperation(event mockgcpregistry.Event) bool {
	u := event.URL()
	// A normal GET poll
	if event.Method() == "GET" && strings.Contains(u, "/operations/") {
		return true
	}
	// A call to the /wait endpoint
	if event.Method() == "POST" && strings.Contains(u, "/operations/") && strings.Contains(u, "/wait") {
		return true
	}
	// A GRPC call
	if u == "/google.longrunning.Operations/GetOperation" {
		return true
	}
	return false
}

// isComputeAPI returns true if this is a compute URL
func isComputeAPI(event mockgcpregistry.Event) bool {
	u, err := url.Parse(event.URL())
	if err != nil {
		klog.Fatalf("cannot parse URL %q", event.URL())
	}
	switch u.Host {
	case "compute.googleapis.com":
		return true
	case "www.googleapis.com":
		return strings.Contains(u.Path, "/compute/")
	}
	return false
}
