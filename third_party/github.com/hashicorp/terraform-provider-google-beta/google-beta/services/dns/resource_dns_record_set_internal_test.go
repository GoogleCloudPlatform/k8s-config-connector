// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0
package dns

import (
	"testing"

	"github.com/hashicorp/terraform-provider-google-beta/google-beta/verify"
)

func TestValidateRecordNameTrailingDot(t *testing.T) {
	cases := []verify.StringValidationTestCase{
		// No errors
		{TestName: "trailing dot", Value: "test-record.hashicorptest.com."},

		// With errors
		{TestName: "empty string", Value: "", ExpectError: true},
		{TestName: "no trailing dot", Value: "test-record.hashicorptest.com", ExpectError: true},
	}

	es := verify.TestStringValidationCases(cases, validateRecordNameTrailingDot)
	if len(es) > 0 {
		t.Errorf("Failed to validate DNS Record name with value: %v", es)
	}
}

func TestExpandDnsRecordSetRoutingPolicyIgnoresNilItems(t *testing.T) {
	cases := []struct {
		name   string
		policy map[string]interface{}
	}{
		{
			name: "wrr",
			policy: map[string]interface{}{
				"health_check":   "",
				"wrr":            []interface{}{nil, nil},
				"geo":            []interface{}{},
				"primary_backup": []interface{}{},
			},
		},
		{
			name: "geo",
			policy: map[string]interface{}{
				"health_check":       "",
				"enable_geo_fencing": false,
				"wrr":                []interface{}{},
				"geo":                []interface{}{nil},
				"primary_backup":     []interface{}{},
			},
		},
		{
			name: "primary backup",
			policy: map[string]interface{}{
				"health_check":   "",
				"wrr":            []interface{}{},
				"geo":            []interface{}{},
				"primary_backup": []interface{}{nil},
			},
		},
	}
	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			got, err := expandDnsRecordSetRoutingPolicy([]interface{}{tc.policy}, nil, nil)
			if err != nil {
				t.Fatalf("unexpected error: %v", err)
			}
			if got != nil {
				t.Fatalf("expected no routing policy, got %#v", got)
			}
		})
	}
}

func TestExpandDnsRecordSetRoutingPolicyWrrSkipsNilItems(t *testing.T) {
	policy := map[string]interface{}{
		"health_check": "",
		"wrr": []interface{}{
			nil,
			map[string]interface{}{
				"rrdatas":                []interface{}{"198.51.100.5"},
				"weight":                 0.5,
				"health_checked_targets": []interface{}{},
			},
		},
		"geo":            []interface{}{},
		"primary_backup": []interface{}{},
	}
	got, err := expandDnsRecordSetRoutingPolicy([]interface{}{policy}, nil, nil)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if got == nil || got.Wrr == nil || len(got.Wrr.Items) != 1 {
		t.Fatalf("expected one WRR item, got %#v", got)
	}
	if item := got.Wrr.Items[0]; item.Weight != 0.5 || len(item.Rrdatas) != 1 || item.Rrdatas[0] != "198.51.100.5" {
		t.Fatalf("unexpected WRR item %#v", item)
	}
}
