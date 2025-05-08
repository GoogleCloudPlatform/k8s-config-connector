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

package resourcefixture

import (
	"fmt"
	"slices"
	"testing"

	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/apis/core/v1alpha1"
	iamapi "github.com/GoogleCloudPlatform/k8s-config-connector/pkg/apis/iam/v1beta1"
	dclmetadata "github.com/GoogleCloudPlatform/k8s-config-connector/pkg/dcl/metadata"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/k8s"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/servicemapping/servicemappingloader"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/test"
	testservicemapping "github.com/GoogleCloudPlatform/k8s-config-connector/pkg/test/servicemapping"
	testservicemappingloader "github.com/GoogleCloudPlatform/k8s-config-connector/pkg/test/servicemappingloader"
	testyaml "github.com/GoogleCloudPlatform/k8s-config-connector/pkg/test/yaml"

	"k8s.io/apimachinery/pkg/apis/meta/v1/unstructured"
	"k8s.io/apimachinery/pkg/runtime/schema"
)

// GetFilteredSetCover is an implementation of https://en.wikipedia.org/wiki/Set_cover_problem#Greedy_algorithm:
// returns a minimal set cover from the resource fixtures that match the given filters.
//
// The set cover is determined by resource config, that is to say, for every resource config there is at least one
// resource returned in the set cover. The result of this function is useful when you wish to run a test for every
// resource type that is supported, but for performance and quota reasons, not across every fixture. For example,
// when you want to test every resource but don't want to test a given resource more than once. For example, the
// basic/pubsubsubscription fixture covers both the PubSubTopic and PubSubSubscription resources. There is no reason
// to also run a test on the basic/pubsubtopic fixture if covering each unique resource is your goal.
func GetFilteredSetCover(t *testing.T, lightFilterFunc LightFilter, heavyFilterFunc HeavyFilter) []ResourceFixture {
	fixtures := LoadWithFilter(t, lightFilterFunc, heavyFilterFunc)
	smLoader := testservicemappingloader.New(t)
	serviceMetadataLoader := dclmetadata.New()
	fixtureRCIds := buildResourceFixtureRCIdGraph(t, smLoader, serviceMetadataLoader, fixtures)
	minFixtureSet := findSetCover(fixtureRCIds)
	return fixtureRCIdsToFixtures(minFixtureSet)
}

func GetBasicTypeSetCover(t *testing.T) []ResourceFixture {
	lightFilter := func(name string, testType TestType) bool {
		return testType == Basic
	}
	heavyFilter := func(fixture ResourceFixture) bool {
		// Skip v1alpha1 CRDs when testing set cover as they may not yet be
		// correctly supported.
		return fixture.GVK.Version == k8s.KCCAPIVersionV1Beta1
	}
	return GetFilteredSetCover(t, lightFilter, heavyFilter)
}

// returns all the resource config ids in use by the resources defined for a given fixture
func getResourceConfigIds(t *testing.T, smLoader *servicemappingloader.ServiceMappingLoader, serviceMetadataLoader dclmetadata.ServiceMetadataLoader, fixture ResourceFixture) map[string]bool {
	resourceConfigIds := make(map[string]bool)
	addResourceConfig(t, smLoader, serviceMetadataLoader, fixture.Create, resourceConfigIds)
	if fixture.Dependencies != nil {
		dependencyYamls := testyaml.SplitYAML(t, fixture.Dependencies)
		for _, d := range dependencyYamls {
			addResourceConfig(t, smLoader, serviceMetadataLoader, d, resourceConfigIds)
		}
	}
	return resourceConfigIds
}

func addResourceConfig(t *testing.T, smLoader *servicemappingloader.ServiceMappingLoader, serviceMetadataLoader dclmetadata.ServiceMetadataLoader, yamlBytes []byte, resourceConfigIds map[string]bool) {
	u := test.ToUnstruct(t, yamlBytes)
	if !ShouldHaveResourceConfig(u, serviceMetadataLoader) {
		return
	}
	rc := testservicemapping.GetResourceConfig(t, smLoader, u)
	resourceConfigIds[GetUniqueResourceConfigID(*rc)] = true
}

// TODO(yuwenma): This is a temp fix. We should use a more generic approach.
func IsPureDirectResource(gk schema.GroupKind) bool {
	pureDirectResources := []string{
		"ApigeeEndpointAttachment",
		"ApigeeEnvgroup",
		"ApigeeEnvgroupAttachment",
		"ApigeeInstance",
		"ApigeeInstanceAttachment",
		"BigQueryConnectionConnection",
		"BigQueryDataTransferConfig",
		"CloudBuildWorkerPool",
		"DataformRepository",
		"FirestoreDatabase",
		"NetworkConnectivityServiceConnectionPolicy",
		"PrivilegedAccessManagerEntitlement",
		"RedisCluster",
		"BigQueryAnalyticsHubDataExchange",
		"BigQueryAnalyticsHubListing",
		"WorkstationCluster",
		"WorkstationConfig",
		"Workstation",
		"KMSAutokeyConfig",
		"KMSKeyHandle",
		"SecureSourceManagerInstance",
		"SecureSourceManagerRepository",
		"ManagedKafkaCluster",
		"ManagedKafkaTopic",
		"WorkflowsWorkflow",
		"IAPSettings",
	}
	return slices.Contains(pureDirectResources, gk.Kind)
}

func ShouldHaveResourceConfig(u *unstructured.Unstructured, serviceMetadataLoader dclmetadata.ServiceMetadataLoader) bool {
	return k8s.IsManagedByKCC(u.GroupVersionKind()) &&
		!iamapi.IsHandwrittenIAM(u.GroupVersionKind()) &&
		!dclmetadata.IsDCLBasedResourceKind(u.GroupVersionKind(), serviceMetadataLoader) &&
		!IsPureDirectResource(u.GroupVersionKind().GroupKind())
}

// returns an id that is unique for each resource config
func GetUniqueResourceConfigID(rc v1alpha1.ResourceConfig) string {
	if rc.Locationality != "" {
		return fmt.Sprintf("%v:%v", rc.Kind, rc.Locationality)
	}
	if rc.Name == "google_compute_instance" || rc.Name == "google_compute_instance_from_template" {
		return fmt.Sprintf("%v:%v", rc.Kind, rc.Name)
	}

	return rc.Kind
}

// this struct is used to construct a graph where the nodes are ResourceFixtures and the edges are resource config IDs
type fixtureRCId struct {
	Fixture ResourceFixture
	RCIds   map[string]bool
}

func buildResourceFixtureRCIdGraph(t *testing.T, smLoader *servicemappingloader.ServiceMappingLoader, serviceMetadataLoader dclmetadata.ServiceMetadataLoader, fixtures []ResourceFixture) []fixtureRCId {
	fixtureRCIds := make([]fixtureRCId, 0)
	for _, f := range fixtures {
		fRCId := fixtureRCId{
			Fixture: f,
			RCIds:   make(map[string]bool),
		}
		fixtureRCIds = append(fixtureRCIds, fRCId)
		resourceConfigIds := getResourceConfigIds(t, smLoader, serviceMetadataLoader, f)
		for k := range resourceConfigIds {
			fRCId.RCIds[k] = true
		}
	}
	return fixtureRCIds
}

func findSetCover(fixtureRCIds []fixtureRCId) []fixtureRCId {
	minFixtureSet := make([]fixtureRCId, 0)
	rcIDToCovered := make(map[string]bool)
	for _, f := range fixtureRCIds {
		for rcID := range f.RCIds {
			rcIDToCovered[rcID] = false
		}
	}
	coverCount := 0
	for coverCount < len(rcIDToCovered) {
		// find set with maximum number uncovered
		var maxUncoverFixture fixtureRCId
		maxUncoverFixtureNewCoverCount := 0
		for _, fk := range fixtureRCIds {
			uncoverCount := getUncoveredCount(fk, rcIDToCovered)
			if uncoverCount > maxUncoverFixtureNewCoverCount {
				maxUncoverFixtureNewCoverCount = uncoverCount
				maxUncoverFixture = fk
			}
		}
		for rcID := range maxUncoverFixture.RCIds {
			rcIDToCovered[rcID] = true
		}
		coverCount += maxUncoverFixtureNewCoverCount
		minFixtureSet = append(minFixtureSet, maxUncoverFixture)
	}
	return minFixtureSet
}

func getUncoveredCount(f fixtureRCId, rcIDToCovered map[string]bool) int {
	count := 0
	for r := range f.RCIds {
		covered, ok := rcIDToCovered[r]
		if !ok {
			panic(fmt.Sprintf("expected resource config id '%v' to be in the map", r))
		}
		if !covered {
			count++
		}
	}
	return count
}

func fixtureRCIdsToFixtures(fixtureRCIds []fixtureRCId) []ResourceFixture {
	resourceFixtures := make([]ResourceFixture, 0, len(fixtureRCIds))
	for _, f := range fixtureRCIds {
		resourceFixtures = append(resourceFixtures, f.Fixture)
	}
	return resourceFixtures
}
