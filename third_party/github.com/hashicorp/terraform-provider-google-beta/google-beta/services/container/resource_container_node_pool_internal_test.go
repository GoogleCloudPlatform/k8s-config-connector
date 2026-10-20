// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0
package container

import (
	"reflect"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	transport_tpg "github.com/hashicorp/terraform-provider-google-beta/google-beta/transport"
	container "google.golang.org/api/container/v1beta1"
)

func TestFlattenNodePool_Autoscaling(t *testing.T) {
	config := &transport_tpg.Config{}
	cases := map[string]struct {
		configValues      map[string]interface{}
		apiNodePool       *container.NodePool
		expectedAutoscale []map[string]interface{}
	}{
		"legacy disable path: min=max=0 in config and api enabled=false returns populated block": {
			configValues: map[string]interface{}{
				"autoscaling": []interface{}{
					map[string]interface{}{
						"min_node_count": 0,
						"max_node_count": 0,
					},
				},
			},
			apiNodePool: &container.NodePool{
				Management: &container.NodeManagement{},
				Autoscaling: &container.NodePoolAutoscaling{
					Enabled:           false,
					MinNodeCount:      0,
					MaxNodeCount:      0,
					TotalMinNodeCount: 0,
					TotalMaxNodeCount: 0,
				},
			},
			expectedAutoscale: []map[string]interface{}{
				{
					"enabled":              false,
					"min_node_count":       int64(0),
					"max_node_count":       int64(0),
					"total_min_node_count": int64(0),
					"total_max_node_count": int64(0),
					"location_policy":      "",
				},
			},
		},
		"enabled autoscaling returns populated block with enabled=true": {
			configValues: map[string]interface{}{
				"autoscaling": []interface{}{
					map[string]interface{}{
						"min_node_count": 1,
						"max_node_count": 3,
					},
				},
			},
			apiNodePool: &container.NodePool{
				Management: &container.NodeManagement{},
				Autoscaling: &container.NodePoolAutoscaling{
					Enabled:           true,
					MinNodeCount:      1,
					MaxNodeCount:      3,
					TotalMinNodeCount: 0,
					TotalMaxNodeCount: 0,
				},
			},
			expectedAutoscale: []map[string]interface{}{
				{
					"enabled":              true,
					"min_node_count":       int64(1),
					"max_node_count":       int64(3),
					"total_min_node_count": int64(0),
					"total_max_node_count": int64(0),
					"location_policy":      "",
				},
			},
		},
		"disabled in api when not configured returns empty slice": {
			configValues: map[string]interface{}{},
			apiNodePool: &container.NodePool{
				Management: &container.NodeManagement{},
				Autoscaling: &container.NodePoolAutoscaling{
					Enabled:           false,
					MinNodeCount:      0,
					MaxNodeCount:      0,
					TotalMinNodeCount: 0,
					TotalMaxNodeCount: 0,
				},
			},
			expectedAutoscale: []map[string]interface{}{},
		},
		"explicit enabled=false in config with non-zero min/max and nil api autoscaling preserves block": {
			configValues: map[string]interface{}{
				"autoscaling": []interface{}{
					map[string]interface{}{
						"enabled":        false,
						"min_node_count": 1,
						"max_node_count": 3,
					},
				},
			},
			apiNodePool: &container.NodePool{
				Management:  &container.NodeManagement{},
				Autoscaling: nil,
			},
			expectedAutoscale: []map[string]interface{}{
				{
					"enabled":              false,
					"min_node_count":       1,
					"max_node_count":       3,
					"total_min_node_count": 0,
					"total_max_node_count": 0,
					"location_policy":      "",
				},
			},
		},
	}

	for tn, tc := range cases {
		t.Run(tn, func(t *testing.T) {
			d := schema.TestResourceDataRaw(t, ResourceContainerNodePool().Schema, tc.configValues)
			flattened, err := flattenNodePool(d, config, tc.apiNodePool, "")
			if err != nil {
				t.Fatalf("unexpected error flattening node pool: %v", err)
			}
			gotAutoscale, ok := flattened["autoscaling"].([]map[string]interface{})
			if !ok {
				t.Fatalf("expected autoscaling to be []map[string]interface{}, got %T", flattened["autoscaling"])
			}
			if !reflect.DeepEqual(gotAutoscale, tc.expectedAutoscale) {
				t.Errorf("flattenNodePool autoscaling = %+v, want %+v", gotAutoscale, tc.expectedAutoscale)
			}
		})
	}
}
