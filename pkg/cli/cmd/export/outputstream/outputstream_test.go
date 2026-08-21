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

package outputstream_test

import (
	"reflect"
	"testing"

	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/cli/cmd/export/outputstream"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/cli/cmd/export/parameters"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/cli/stream"
	testservicemappingloader "github.com/GoogleCloudPlatform/k8s-config-connector/pkg/test/servicemappingloader"
	tfprovider "github.com/GoogleCloudPlatform/k8s-config-connector/pkg/tf/provider"
)

func TestNewUnstructuredStreamWithNoneIAMOption(t *testing.T) {
	testNewUnstructuredStreamWithIAMOption(t, "none", stream.UnstructuredResourceFixupStream{})
}

func TestNewUnstructuredStreamWithPolicyIAMOption(t *testing.T) {
	testNewUnstructuredStreamWithIAMOption(t, "policy", stream.UnstructuredResourceAndIAMPolicyStream{})
}

func TestNewUnstructuredStreamWithPolicyMemberIAMOption(t *testing.T) {
	testNewUnstructuredStreamWithIAMOption(t, "policymember", stream.UnstructuredResourceAndIAMPolicyStream{})
}

func testNewUnstructuredStreamWithIAMOption(t *testing.T, iamFormatOption string, instanceOfExpectedType interface{}) {
	params := parameters.Parameters{
		IAMFormat:      iamFormatOption,
		GCPAccessToken: "dummyToken",
	}
	unstructuredStream, err := outputstream.NewUnstructuredStream(&params, tfprovider.NewOrLogFatal(tfprovider.UnitTestConfig()),
		testservicemappingloader.New(t))
	if err != nil {
		t.Fatalf("error creating stream: %v", err)
	}
	expectedType := reflect.TypeOf(instanceOfExpectedType)
	actualType := reflect.TypeOf(unstructuredStream).Elem()
	if expectedType != actualType {
		t.Fatalf("unexpected type for unstructured stream: got '%v', want '%v'", actualType.Name(), expectedType.Name())
	}
}
