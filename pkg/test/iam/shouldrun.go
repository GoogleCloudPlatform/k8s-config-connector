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

package testiam

import "github.com/GoogleCloudPlatform/k8s-config-connector/pkg/test/resourcefixture"

func ShouldRunWithExternalRef(fixture resourcefixture.ResourceFixture) bool {
	// We only need to test the case of "IAMPolicy (or IAMPolicyMember) having
	// an external reference" for a few resources. We could test both cases
	// (IAMPolicy with regular reference and IAMPolicy with external reference)
	// for all resources that can be referenced by IAMPolicy, but this is very
	// expensive and not really necessary.
	//
	// Note: only resources with user-specified IDs are supported since
	// NewExternalRef() cannot generate external references to resources with
	// server-generated IDs (e.g. Folder).
	switch fixture.GVK.Kind {
	case "Project", // Commonly referenced resource for IAMPolicy/IAMPolicyMember
		"PubSubTopic",     // Basic resource with no dependencies
		"SpannerDatabase": // Resource whose IAMPolicy/IAMPolicyMember spec must contain info about a dependency (name of the SpannerInstance)
		return true
	default:
		return false
	}
}

func ShouldRunWithIAMConditions(fixture resourcefixture.ResourceFixture) bool {
	// We only need to test the case of "IAMPolicy (or IAMPolicyMember) with
	// IAM condition" for a few resources. We could test IAM conditions for all
	// resources that support conditions, but this is very expensive and not
	// really necessary.
	switch fixture.GVK.Kind {
	case "Project", // Commonly referenced resource for IAMPolicy/IAMPolicyMember
		"KMSKeyRing": // Basic resource that supports IAM conditions
		return true
	default:
		return false
	}
}

func ShouldRunWithAuditConfigs(fixture resourcefixture.ResourceFixture) bool {
	// Only the following resources support IAM audit configs in KCC currently
	switch fixture.GVK.Kind {
	case "Project", "Folder":
		return true
	default:
		return false
	}
}

func ShouldRunDeleteParentFirst(fixture resourcefixture.ResourceFixture) bool {
	// Only the following resource(s) will be used for deletion ordering tests
	return fixture.GVK.Kind == "PubSubTopic"
}

func ShouldRunAcquire(fixture resourcefixture.ResourceFixture) bool {
	// Only the following resource(s) will be used for IAM resource acquisition tests
	return fixture.GVK.Kind == "PubSubTopic"
}

func ShouldRunWithTFResourcesOnly(fixture resourcefixture.ResourceFixture) bool {
	switch fixture.GVK.Kind {
	case "BigtableInstance", "KMSKeyRing", "Project", "Folder",
		"PubSubTopic", "PubSubSubscription", "SpannerInstance", "StorageBucket", "IAMServiceAccount":
		return true
	default:
		return false
	}
}
