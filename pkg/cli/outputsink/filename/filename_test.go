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

package filename_test

import (
	"context"
	"fmt"
	"testing"

	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/cli/outputsink/filename"
	testservicemappingloader "github.com/GoogleCloudPlatform/k8s-config-connector/pkg/test/servicemappingloader"
	tfprovider "github.com/GoogleCloudPlatform/k8s-config-connector/pkg/tf/provider"

	testyaml "github.com/GoogleCloudPlatform/k8s-config-connector/pkg/test/yaml"
	"k8s.io/apimachinery/pkg/apis/meta/v1/unstructured"
)

func TestPubSubTopic(t *testing.T) {
	testResource(t, "pubsubtopic.yaml", "projects/my-project-id/PubSubTopic/pubsubtopic-bl5f09tftk5sk8gc0q4g", "")
}

func TestPubSuSubscription(t *testing.T) {
	testResource(t, "pubsubsubscription.yaml", "projects/kcc-test/PubSubSubscription/pubsubtopic1-nt3re4g91rdcc0clj3a1.subscription-4158350662426637332", "")
}

func TestSpannerDatabase(t *testing.T) {
	testResource(t, "spannerdatabase.yaml", "projects/kcc-test/SpannerInstance/spannerdatabase-dep/SpannerDatabase/asdf", "")
}

func TestComputeSubnetwork(t *testing.T) {
	testResource(t, "computesubnetwork.yaml", "projects/kcc-test/ComputeSubnetwork/us-west1/default", "")
}

func TestProject(t *testing.T) {
	testResource(t, "project.yaml", "organizations/1234567890/Project/test-project", "")
}

func TestIAMPolicyPubSubTopic(t *testing.T) {
	testResource(t, "iampolicy-pubsubtopic.yaml", "projects/kcc-test/PubSubTopic/my-new-topic-iampolicy", "")
}

func TestIAMPolicySpannerDatabase(t *testing.T) {
	testResource(t, "iampolicy-spannerdatabase.yaml", "projects/test-project/SpannerInstance/spannerdatabase-dep/SpannerDatabase/asdf-iampolicy", "")
}

func testResource(t *testing.T, testResourceFile, expectedPath, expectedErrMsg string) {
	ctx := context.TODO()

	smLoader := testservicemappingloader.New(t)
	tfProvider := tfprovider.NewOrLogFatal(tfprovider.UnitTestConfig())
	var u unstructured.Unstructured
	testyaml.UnmarshalFile(t, fmt.Sprintf("testdata/%v", testResourceFile), &u)
	path, err := filename.Get(ctx, &u, smLoader, tfProvider)
	errMsg := ""
	if err != nil {
		errMsg = err.Error()
	}
	if errMsg != expectedErrMsg {
		t.Fatalf("unexpected error message: got '%v', want '%v'", errMsg, expectedErrMsg)
	}
	if path != expectedPath {
		t.Fatalf("unexpected filepath result: got '%v', want '%v'", path, expectedPath)
	}
}
