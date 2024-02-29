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

package asset_test

import (
	"errors"
	"io"
	"testing"

	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/cli/asset"
	testservicemappingloader "github.com/GoogleCloudPlatform/k8s-config-connector/pkg/test/servicemappingloader"
)

func TestGetServiceMappingAndResourceConfig(t *testing.T) {
	stream := newStreamFromFile(t, "testdata/cai-export.json")
	defer closeStream(t, stream)
	smLoader := testservicemappingloader.New(t)
	for a, err := stream.Next(); !errors.Is(err, io.EOF); a, err = stream.Next() {
		if err != nil {
			t.Fatalf("unexpected error reading asset: %v", err)
		}
		sm, rc, err := asset.GetServiceMappingAndResourceConfig(smLoader, a)
		if err != nil {
			t.Fatalf("error getting service mapping and resource config: %v", err)
		}
		if sm == nil || rc == nil {
			t.Fatalf("unexpected nil value")
		}
	}
}
