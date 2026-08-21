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

	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/cli/asset"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/cli/cmd/bulkexport/inputstream"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/cli/cmd/bulkexport/outputstream"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/cli/cmd/bulkexport/parameters"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/cli/stream"
	testos "github.com/GoogleCloudPlatform/k8s-config-connector/pkg/cli/test/os"
	testservicemappingloader "github.com/GoogleCloudPlatform/k8s-config-connector/pkg/test/servicemappingloader"
	tfprovider "github.com/GoogleCloudPlatform/k8s-config-connector/pkg/tf/provider"
)

func TestNewUnstructuredSTreamWithUnknownIAMOption(t *testing.T) {
	cleanup, stdin := testos.GetStdin(t, "garbage string")
	defer cleanup()
	params := parameters.Parameters{}
	params.IAMFormat = "garbage"
	params.OAuth2Token = "dummyToken"
	assetStream, err := inputstream.NewAssetStream(&params, stdin)
	if err != nil {
		t.Fatalf("error creating asset stream: %v", err)
	}
	defer closeStream(t, assetStream)
	_, err = outputstream.NewUnstructuredStream(&params, assetStream, tfprovider.NewOrLogFatal(tfprovider.UnitTestConfig()),
		testservicemappingloader.New(t))
	if err == nil {
		t.Fatal("invalid error value: got 'nil', want an error")
	}
	expectedErrMsg := "invalid policy format option 'garbage'"
	actualErrMsg := err.Error()
	if actualErrMsg != expectedErrMsg {
		t.Fatalf("unexpected error message: got '%v', want '%v'", actualErrMsg, expectedErrMsg)
	}
}

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
	cleanup, stdin := testos.GetStdin(t, "garbage string")
	defer cleanup()
	params := parameters.Parameters{}
	assetStream, err := inputstream.NewAssetStream(&params, stdin)
	if err != nil {
		t.Fatalf("error creating asset stream: %v", err)
	}
	defer closeStream(t, assetStream)
	params.IAMFormat = iamFormatOption
	params.OAuth2Token = "dummyToken"
	unstructuredStream, err := outputstream.NewUnstructuredStream(&params, assetStream, tfprovider.NewOrLogFatal(tfprovider.UnitTestConfig()),
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

func closeStream(t *testing.T, assetStream *asset.Stream) {
	err := assetStream.Close()
	if err != nil {
		t.Fatalf("error closing stream: %v", err)
	}
}
