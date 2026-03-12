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

package firestore

import (
	"encoding/json"
	"testing"

	pb "cloud.google.com/go/firestore/apiv1/firestorepb"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct"
	"github.com/google/go-cmp/cmp"
	"google.golang.org/protobuf/testing/protocmp"
)

// There are some difficult FirestoreDocument field values which we want to test specially
func TestSpecialRoundTripDocumentValues(t *testing.T) {
	grid := []struct {
		Name         string
		ExpectedJSON string
		Value        *pb.Value
	}{
		{
			Name:         "Integer and Float are preserved",
			ExpectedJSON: `{"databaseRef":{},"fields":{"specialValues":[100.0,100]}}`,
			Value: &pb.Value{
				ValueType: &pb.Value_ArrayValue{
					ArrayValue: &pb.ArrayValue{
						Values: []*pb.Value{
							{
								// 100.0 would normally be converted to an integer in JSON
								ValueType: &pb.Value_DoubleValue{DoubleValue: 100.0},
							},
							{
								// The same value as an integer
								ValueType: &pb.Value_IntegerValue{IntegerValue: 100},
							},
						},
					},
				},
			},
		},
		{
			Name:         "Large Integer and Float",
			ExpectedJSON: `{"databaseRef":{},"fields":{"specialValues":[5903169270288737330,5.903169270288737e+18]}}`,
			Value: &pb.Value{
				ValueType: &pb.Value_ArrayValue{
					ArrayValue: &pb.ArrayValue{
						Values: []*pb.Value{
							{
								// This value is large and would normally be converted to a float in JSON
								ValueType: &pb.Value_IntegerValue{IntegerValue: 5903169270288737330},
							},
							{
								// The same value as a float
								ValueType: &pb.Value_DoubleValue{DoubleValue: 5.903169270288737e+18},
							},
						},
					},
				},
			},
		},
		{
			Name:         "Empty array",
			ExpectedJSON: `{"databaseRef":{},"fields":{"specialValues":[]}}`,
			Value: &pb.Value{
				ValueType: &pb.Value_ArrayValue{
					ArrayValue: &pb.ArrayValue{
						Values: []*pb.Value{},
					},
				},
			},
		},
	}

	for _, g := range grid {
		t.Run(g.Name, func(t *testing.T) {
			inputProto := &pb.Document{
				Fields: map[string]*pb.Value{
					"specialValues": g.Value,
				},
			}

			// Proto -> KRM
			mapCtx := &direct.MapContext{}
			inputKRM := FirestoreDocumentSpec_v1alpha1_FromProto(mapCtx, inputProto)
			if mapCtx.Err() != nil {
				t.Fatalf("unexpected error mapping from proto: %v", mapCtx.Err())
			}

			// Verify expected JSON representation
			jsonData, err := json.Marshal(inputKRM)
			if err != nil {
				t.Fatalf("unexpected error marshalling to JSON: %v", err)
			}

			if diff := cmp.Diff(g.ExpectedJSON, string(jsonData)); diff != "" {
				t.Errorf("JSON representation mismatch (-want +got):\n%s", diff)
			}

			// // JSON -> KRM
			// var outputKRM krm.FirestoreDocumentSpec
			// if err := json.Unmarshal(jsonData, &outputKRM); err != nil {
			// 	t.Fatalf("unexpected error unmarshalling from JSON: %v", err)
			// }

			// if diff := cmp.Diff(inputKRM, outputKRM); diff != "" {
			// 	t.Errorf("KRM round-trip mismatch (-want +got):\n%s", diff)
			// }

			// KRM -> Proto
			mapCtx = &direct.MapContext{}
			outputProto := FirestoreDocumentSpec_v1alpha1_ToProto(mapCtx, inputKRM)
			if mapCtx.Err() != nil {
				t.Fatalf("unexpected error mapping to proto: %v", mapCtx.Err())
			}

			// Verify round-trip
			if diff := cmp.Diff(inputProto, outputProto, protocmp.Transform()); diff != "" {
				t.Errorf("round-trip mismatch (-want +got):\n%s", diff)
			}
		})
	}
}
