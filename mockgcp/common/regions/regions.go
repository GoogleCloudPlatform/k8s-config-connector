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

package regions

import (
	"context"
	"encoding/json"
	"fmt"
	"strings"

	_ "embed"
)

type Region struct {
	Name string

	// DefaultCIDR is the CIDR created for the subnet, when using auto subnet creation.
	DefaultCIDR string
}

var regions = []Region{
	{Name: "us-central1", DefaultCIDR: "10.128.0.0/20"},
	{Name: "europe-west1", DefaultCIDR: "10.132.0.0/20"},
	{Name: "us-west1", DefaultCIDR: "10.138.0.0/20"},
	{Name: "asia-east1", DefaultCIDR: "10.140.0.0/20"},
	{Name: "us-east1", DefaultCIDR: "10.142.0.0/20"},
	{Name: "asia-northeast1", DefaultCIDR: "10.146.0.0/20"},
	{Name: "asia-southeast1", DefaultCIDR: "10.148.0.0/20"},
	{Name: "us-east4", DefaultCIDR: "10.150.0.0/20"},
	{Name: "australia-southeast1", DefaultCIDR: "10.152.0.0/20"},
	{Name: "europe-west2", DefaultCIDR: "10.154.0.0/20"},
	{Name: "europe-west3", DefaultCIDR: "10.156.0.0/20"},
	{Name: "southamerica-east1", DefaultCIDR: "10.158.0.0/20"},
	{Name: "asia-south1", DefaultCIDR: "10.160.0.0/20"},
	{Name: "northamerica-northeast1", DefaultCIDR: "10.162.0.0/20"},
	{Name: "europe-west4", DefaultCIDR: "10.164.0.0/20"},
	{Name: "europe-north1", DefaultCIDR: "10.166.0.0/20"},
	{Name: "us-west2", DefaultCIDR: "10.168.0.0/20"},
	{Name: "asia-east2", DefaultCIDR: "10.170.0.0/20"},
	{Name: "europe-west6", DefaultCIDR: "10.172.0.0/20"},
	{Name: "asia-northeast2", DefaultCIDR: "10.174.0.0/20"},
	{Name: "asia-northeast3", DefaultCIDR: "10.178.0.0/20"},
	{Name: "us-west3", DefaultCIDR: "10.180.0.0/20"},
	{Name: "us-west4", DefaultCIDR: "10.182.0.0/20"},
	{Name: "asia-southeast2", DefaultCIDR: "10.184.0.0/20"},
	{Name: "europe-central2", DefaultCIDR: "10.186.0.0/20"},
	{Name: "northamerica-northeast2", DefaultCIDR: "10.188.0.0/20"},
	{Name: "asia-south2", DefaultCIDR: "10.190.0.0/20"},
	{Name: "australia-southeast2", DefaultCIDR: "10.192.0.0/20"},
	{Name: "southamerica-west1", DefaultCIDR: "10.194.0.0/20"},
	{Name: "europe-west8", DefaultCIDR: "10.198.0.0/20"},
	{Name: "europe-west9", DefaultCIDR: "10.200.0.0/20"},
	{Name: "us-east5", DefaultCIDR: "10.202.0.0/20"},
	{Name: "europe-southwest1", DefaultCIDR: "10.204.0.0/20"},
	{Name: "us-south1", DefaultCIDR: "10.206.0.0/20"},
	{Name: "me-west1", DefaultCIDR: "10.208.0.0/20"},
	{Name: "europe-west12", DefaultCIDR: "10.210.0.0/20"},
	{Name: "me-central1", DefaultCIDR: "10.212.0.0/20"},
	{Name: "europe-west10", DefaultCIDR: "10.214.0.0/20"},
	{Name: "africa-south1", DefaultCIDR: "10.218.0.0/20"},
	{Name: "northamerica-south1", DefaultCIDR: "10.224.0.0/20"},
	{Name: "europe-north2", DefaultCIDR: "10.226.0.0/20"},

	// Some regions appear in results, but don't e.g. have default subnets.
	{Name: "me-central2"},
}

// GetAllRegions returns all accessible regions.
func GetAllRegions(ctx context.Context) []Region {
	return regions
}

// Zones returns all zones in the region.
func (r *Region) Zones(ctx context.Context) []string {
	results := []string{}
	for _, zone := range zones {
		if lastComponent(zone.Region) == r.Name {
			results = append(results, zone.Name)
		}
	}
	return results
}

//go:embed zones.json
var zonesJSON []byte

type Zone struct {
	AvailableCpuPlatforms []string `json:"availableCpuPlatforms"`
	CreationTimestamp     string   `json:"creationTimestamp"`
	Description           string   `json:"description"`
	Name                  string   `json:"name"`
	Id                    string   `json:"id"`
	Kind                  string   `json:"kind"`
	Region                string   `json:"region"`
	SelfLink              string   `json:"selfLink"`
	Status                string   `json:"status"`
	SupportsPzs           bool     `json:"supportsPzs"`
}

var zones []Zone

func init() {
	if err := json.Unmarshal(zonesJSON, &zones); err != nil {
		panic(fmt.Sprintf("failed to unmarshal zones.json: %v", err))
	}
}

func lastComponent(s string) string {
	i := strings.LastIndex(s, "/")
	return s[i+1:]
}
