// Copyright 2026 Google LLC
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

package sorter

import (
	"bytes"
	"fmt"
	"io"
	"os"
	"path/filepath"
	"sort"
	"strings"

	"gopkg.in/yaml.v3"
)

// ResourceTier assigns topological priority to KCC resource kinds.
// Lower tier numbers must be applied before higher tier numbers.
var ResourceTier = map[string]int{
	// Tier 0: Root Foundations & Identities
	"ResourceManagerProject": 0,
	"ComputeNetwork":         0,
	"KMSKeyRing":             0,
	"IAMServiceAccount":      0,
	"StorageBucket":          0,

	// Tier 1: Sub-foundations, Keys, Secrets & Subnets
	"KMSCryptoKey":           1,
	"ComputeSubnetwork":      1,
	"SecretManagerSecret":    1,
	"PubSubTopic":            1,
	"ServiceDirectoryNamespace": 1,

	// Tier 2: Core Stateful Engines, Networking Routes & DNS Zones
	"SQLInstance":            2,
	"SpannerInstance":        2,
	"RedisInstance":          2,
	"ComputeFirewall":        2,
	"ComputeRouter":          2,
	"ComputeAddress":         2,
	"DNSManagedZone":         2,

	// Tier 3: Stateful Children, Users, Databases & Secret Versions
	"SQLDatabase":                3,
	"SQLUser":                    3,
	"SQLSSLConfig":               3,
	"SpannerDatabase":            3,
	"SecretManagerSecretVersion": 3,
	"PubSubSubscription":         3,
	"DNSRecordSet":               3,

	// Tier 4: Routing, Forwarding & Services
	"ComputeBackendService":      4,
	"ComputeURLMap":              4,
	"ComputeTargetHttpProxy":     4,
	"ComputeForwardingRule":      4,
	"ServiceDirectoryService":    4,

	// Tier 5: IAM Policy Members & Access Control Bindings
	"IAMPolicyMember":        5,
	"IAMPolicy":              5,
	"IAMPartialPolicy":       5,
	"IAMAuditConfig":         5,
	"StorageBucketAccessControl": 5,
}

const DefaultTier = 10

// GetKindTier returns the topological tier for a given KCC Kind.
func GetKindTier(kind string) int {
	if tier, found := ResourceTier[kind]; found {
		return tier
	}
	return DefaultTier
}

// ResourceItem holds a parsed KCC resource and its topological metadata.
type ResourceItem struct {
	Kind      string
	Name      string
	Namespace string
	Tier      int
	Raw       map[string]interface{}
}

// ParseResourcesFromData decodes YAML data into a slice of ResourceItem.
func ParseResourcesFromData(data []byte) ([]ResourceItem, error) {
	decoder := yaml.NewDecoder(bytes.NewReader(data))
	var items []ResourceItem

	for {
		var doc map[string]interface{}
		err := decoder.Decode(&doc)
		if err == io.EOF {
			break
		}
		if err != nil {
			return nil, fmt.Errorf("decoding YAML: %w", err)
		}
		if len(doc) == 0 {
			continue
		}

		// Handle List kind
		if kind, ok := doc["kind"].(string); ok && kind == "List" {
			if rawItems, ok := doc["items"].([]interface{}); ok {
				for _, rawItem := range rawItems {
					if itemMap, ok := rawItem.(map[string]interface{}); ok {
						item := createResourceItem(itemMap)
						items = append(items, item)
					}
				}
			}
		} else {
			item := createResourceItem(doc)
			items = append(items, item)
		}
	}

	return items, nil
}

func createResourceItem(doc map[string]interface{}) ResourceItem {
	kind, _ := doc["kind"].(string)
	name := ""
	namespace := ""
	if meta, ok := doc["metadata"].(map[string]interface{}); ok {
		if n, ok := meta["name"].(string); ok {
			name = n
		}
		if ns, ok := meta["namespace"].(string); ok {
			namespace = ns
		}
	}

	return ResourceItem{
		Kind:      kind,
		Name:      name,
		Namespace: namespace,
		Tier:      GetKindTier(kind),
		Raw:       doc,
	}
}

// SortResources sorts resources topologically by tier, then kind, then name.
func SortResources(items []ResourceItem) []ResourceItem {
	sorted := make([]ResourceItem, len(items))
	copy(sorted, items)

	sort.SliceStable(sorted, func(i, j int) bool {
		if sorted[i].Tier != sorted[j].Tier {
			return sorted[i].Tier < sorted[j].Tier
		}
		if sorted[i].Kind != sorted[j].Kind {
			return sorted[i].Kind < sorted[j].Kind
		}
		return sorted[i].Name < sorted[j].Name
	})

	return sorted
}

// SortDirectory reads all YAML files in srcDir, extracts and sorts resources, and writes
// ordered tiered manifests into dstDir (e.g. tier0_ordered.yaml, tier1_ordered.yaml, etc.).
func SortDirectory(srcDir, dstDir string) (int, error) {
	if _, err := os.Stat(srcDir); os.IsNotExist(err) {
		return 0, nil
	}
	files, err := os.ReadDir(srcDir)
	if err != nil {
		return 0, fmt.Errorf("reading srcDir %s: %w", srcDir, err)
	}

	var allItems []ResourceItem
	for _, f := range files {
		if f.IsDir() || (!strings.HasSuffix(f.Name(), ".yaml") && !strings.HasSuffix(f.Name(), ".yml")) {
			continue
		}
		p := filepath.Join(srcDir, f.Name())
		data, err := os.ReadFile(p)
		if err != nil {
			return 0, fmt.Errorf("reading %s: %w", p, err)
		}
		items, err := ParseResourcesFromData(data)
		if err != nil {
			return 0, fmt.Errorf("parsing %s: %w", p, err)
		}
		allItems = append(allItems, items...)
	}

	if len(allItems) == 0 {
		return 0, nil
	}

	sorted := SortResources(allItems)

	if err := os.MkdirAll(dstDir, 0755); err != nil {
		return 0, fmt.Errorf("creating dstDir %s: %w", dstDir, err)
	}

	// Group by Tier
	tierMap := make(map[int][]ResourceItem)
	for _, it := range sorted {
		tierMap[it.Tier] = append(tierMap[it.Tier], it)
	}

	var tiers []int
	for t := range tierMap {
		tiers = append(tiers, t)
	}
	sort.Ints(tiers)

	for _, t := range tiers {
		tierItems := tierMap[t]
		outFileName := fmt.Sprintf("tier%d_ordered.yaml", t)
		outPath := filepath.Join(dstDir, outFileName)

		outFile, err := os.Create(outPath)
		if err != nil {
			return 0, fmt.Errorf("creating %s: %w", outPath, err)
		}

		encoder := yaml.NewEncoder(outFile)
		encoder.SetIndent(2)

		for _, item := range tierItems {
			if err := encoder.Encode(item.Raw); err != nil {
				outFile.Close()
				return 0, fmt.Errorf("encoding to %s: %w", outPath, err)
			}
		}
		encoder.Close()
		outFile.Close()
	}

	return len(sorted), nil
}
