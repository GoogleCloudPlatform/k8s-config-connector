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

package vision

import (
	"context"
	"sort"
	"testing"

	pb "cloud.google.com/go/vision/v2/apiv1/visionpb"
	krm "github.com/GoogleCloudPlatform/k8s-config-connector/apis/vision/v1alpha1"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct"
	"github.com/google/go-cmp/cmp"
	"google.golang.org/protobuf/testing/protocmp"
	"k8s.io/utils/ptr"
)

func TestVisionProductSpecMapping(t *testing.T) {
	originalKRM := &krm.VisionProductSpec{
		DisplayName:     ptr.To("My awesome product"),
		Description:     ptr.To("Product description"),
		ProductCategory: ptr.To("homegoods-v2"),
		ProductLabels: []krm.ProductKeyValue{
			{
				Key:   ptr.To("color"),
				Value: ptr.To("red"),
			},
			{
				Key:   ptr.To("material"),
				Value: ptr.To("wood"),
			},
		},
	}

	mapCtx := &direct.MapContext{}
	gotProto := VisionProductSpec_ToProto(mapCtx, originalKRM)
	if mapCtx.Err() != nil {
		t.Fatalf("error mapping Spec to Proto: %v", mapCtx.Err())
	}

	wantProto := &pb.Product{
		DisplayName:     "My awesome product",
		Description:     "Product description",
		ProductCategory: "homegoods-v2",
		ProductLabels: []*pb.Product_KeyValue{
			{
				Key:   "color",
				Value: "red",
			},
			{
				Key:   "material",
				Value: "wood",
			},
		},
	}

	if diff := cmp.Diff(gotProto, wantProto, protocmp.Transform()); diff != "" {
		t.Errorf("unexpected diff in ToProto mapping: %s", diff)
	}

	gotKRM := VisionProductSpec_FromProto(mapCtx, gotProto)
	if mapCtx.Err() != nil {
		t.Fatalf("error mapping Proto to Spec: %v", mapCtx.Err())
	}

	// ProductCategory is immutable, and the GET API returns it.
	// DisplayName, Description, ProductLabels should roundtrip perfectly.
	if diff := cmp.Diff(gotKRM, originalKRM); diff != "" {
		t.Errorf("unexpected diff in roundtrip FromProto mapping: %s", diff)
	}
}

func TestCompareProduct(t *testing.T) {
	actual := &pb.Product{
		Name:            "projects/p1/locations/l1/products/pID",
		DisplayName:     "Old Name",
		Description:     "Old Description",
		ProductCategory: "homegoods-v2",
		ProductLabels: []*pb.Product_KeyValue{
			{
				Key:   "color",
				Value: "blue",
			},
		},
	}

	desired := &pb.Product{
		Name:            "projects/p1/locations/l1/products/pID",
		DisplayName:     "New Name",
		Description:     "New Description",
		ProductCategory: "homegoods-v2",
		ProductLabels: []*pb.Product_KeyValue{
			{
				Key:   "color",
				Value: "red",
			},
		},
	}

	ctx := context.Background()
	diffs, updateMask, err := compareProduct(ctx, actual, desired)
	if err != nil {
		t.Fatalf("error comparing products: %v", err)
	}

	if !diffs.HasDiff() {
		t.Errorf("expected to find a diff, but found none")
	}

	wantPaths := []string{"display_name", "description", "product_labels"}
	gotPaths := updateMask.GetPaths()
	sort.Strings(gotPaths)
	sort.Strings(wantPaths)
	if diff := cmp.Diff(gotPaths, wantPaths); diff != "" {
		t.Errorf("unexpected update mask paths: %s", diff)
	}
}
