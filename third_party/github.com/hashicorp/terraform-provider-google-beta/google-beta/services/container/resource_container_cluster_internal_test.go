// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0
package container

import (
	"testing"

	container "google.golang.org/api/container/v1beta1"
)

func TestValidateNodePoolAutoConfig(t *testing.T) {
	withTags := &container.NodePoolAutoConfig{
		NetworkTags: &container.NetworkTags{
			Tags: []string{"not-empty"},
		},
	}
	noTags := &container.NodePoolAutoConfig{}

	cases := map[string]struct {
		Input       *container.Cluster
		ExpectError bool
	}{
		"with tags, nap nil, autopilot nil": {
			Input:       &container.Cluster{NodePoolAutoConfig: withTags},
			ExpectError: true,
		},
		"with tags, autopilot disabled": {
			Input: &container.Cluster{
				Autopilot:          &container.Autopilot{Enabled: false},
				NodePoolAutoConfig: withTags,
			},
			ExpectError: true,
		},
		"with tags, nap disabled": {
			Input: &container.Cluster{
				Autoscaling:        &container.ClusterAutoscaling{EnableNodeAutoprovisioning: false},
				NodePoolAutoConfig: withTags,
			},
			ExpectError: true,
		},
		"with tags, autopilot enabled": {
			Input: &container.Cluster{
				Autopilot:          &container.Autopilot{Enabled: true},
				NodePoolAutoConfig: withTags,
			},
			ExpectError: false,
		},
		"with tags, nap enabled": {
			Input: &container.Cluster{
				Autoscaling:        &container.ClusterAutoscaling{EnableNodeAutoprovisioning: true},
				NodePoolAutoConfig: withTags,
			},
			ExpectError: false,
		},
		"no tags, autopilot enabled": {
			Input: &container.Cluster{
				Autopilot:          &container.Autopilot{Enabled: true},
				NodePoolAutoConfig: noTags,
			},
			ExpectError: false,
		},
		"no tags, nap enabled": {
			Input: &container.Cluster{
				Autoscaling:        &container.ClusterAutoscaling{EnableNodeAutoprovisioning: true},
				NodePoolAutoConfig: noTags,
			},
			ExpectError: false,
		},
		"no tags, autopilot disabled": {
			Input: &container.Cluster{
				Autopilot:          &container.Autopilot{Enabled: false},
				NodePoolAutoConfig: noTags,
			},
			ExpectError: false,
		},
		"no tags, nap disabled": {
			Input: &container.Cluster{
				Autoscaling:        &container.ClusterAutoscaling{EnableNodeAutoprovisioning: false},
				NodePoolAutoConfig: noTags,
			},
			ExpectError: false,
		},
	}

	for tn, tc := range cases {
		if err := validateNodePoolAutoConfig(tc.Input); (err != nil) != tc.ExpectError {
			t.Fatalf("bad: '%s', expected error: %t, received error: %t", tn, tc.ExpectError, (err != nil))
		}
	}
}
