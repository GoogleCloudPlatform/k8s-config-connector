// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0
package compute

import (
	"reflect"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func TestResourceComputeSecurityPolicy_enforceOnKeySchema(t *testing.T) {
	res := ResourceComputeSecurityPolicy()
	ruleSchema := res.Schema["rule"].Elem.(*schema.Resource).Schema
	rateLimitOptionsSchema := ruleSchema["rate_limit_options"].Elem.(*schema.Resource).Schema
	enforceOnKeySchema := rateLimitOptionsSchema["enforce_on_key"]

	if !enforceOnKeySchema.Optional {
		t.Errorf("expected enforce_on_key to be Optional")
	}
	if !enforceOnKeySchema.Computed {
		t.Errorf("expected enforce_on_key to be Computed")
	}
	if enforceOnKeySchema.Default != nil {
		t.Errorf("expected enforce_on_key.Default to be nil, got %v", enforceOnKeySchema.Default)
	}
}

func TestExpandSecurityPolicyRuleRateLimitOptions(t *testing.T) {
	threshold := []interface{}{
		map[string]interface{}{
			"count":        100,
			"interval_sec": 60,
		},
	}

	tests := []struct {
		name                 string
		input                []interface{}
		expectedEnforceOnKey string
		expectedConfigsCount int
		expectedForceSend    []string
	}{
		{
			name: "enforce_on_key_configs only",
			input: []interface{}{
				map[string]interface{}{
					"ban_threshold":           []interface{}{},
					"rate_limit_threshold":    threshold,
					"exceed_action":          "deny(429)",
					"conform_action":         "allow",
					"ban_duration_sec":        0,
					"exceed_redirect_options": []interface{}{},
					"enforce_on_key":          "",
					"enforce_on_key_name":     "",
					"enforce_on_key_configs": []interface{}{
						map[string]interface{}{
							"enforce_on_key_type": "IP",
							"enforce_on_key_name": "",
						},
					},
				},
			},
			expectedEnforceOnKey: "",
			expectedConfigsCount: 1,
			expectedForceSend:    []string{"EnforceOnKeyConfigs"},
		},
		{
			name: "enforce_on_key only",
			input: []interface{}{
				map[string]interface{}{
					"ban_threshold":           []interface{}{},
					"rate_limit_threshold":    threshold,
					"exceed_action":          "deny(429)",
					"conform_action":         "allow",
					"ban_duration_sec":        0,
					"exceed_redirect_options": []interface{}{},
					"enforce_on_key":          "IP",
					"enforce_on_key_name":     "",
					"enforce_on_key_configs":  []interface{}{},
				},
			},
			expectedEnforceOnKey: "IP",
			expectedConfigsCount: 0,
			expectedForceSend:    []string{"EnforceOnKey", "EnforceOnKeyName"},
		},
		{
			name: "neither configured - defaults to ALL",
			input: []interface{}{
				map[string]interface{}{
					"ban_threshold":           []interface{}{},
					"rate_limit_threshold":    threshold,
					"exceed_action":          "deny(429)",
					"conform_action":         "allow",
					"ban_duration_sec":        0,
					"exceed_redirect_options": []interface{}{},
					"enforce_on_key":          "",
					"enforce_on_key_name":     "",
					"enforce_on_key_configs":  []interface{}{},
				},
			},
			expectedEnforceOnKey: "ALL",
			expectedConfigsCount: 0,
			expectedForceSend:    []string{"EnforceOnKey", "EnforceOnKeyName"},
		},
		{
			name: "both configured",
			input: []interface{}{
				map[string]interface{}{
					"ban_threshold":           []interface{}{},
					"rate_limit_threshold":    threshold,
					"exceed_action":          "deny(429)",
					"conform_action":         "allow",
					"ban_duration_sec":        0,
					"exceed_redirect_options": []interface{}{},
					"enforce_on_key":          "IP",
					"enforce_on_key_name":     "",
					"enforce_on_key_configs": []interface{}{
						map[string]interface{}{
							"enforce_on_key_type": "HTTP_PATH",
							"enforce_on_key_name": "",
						},
					},
				},
			},
			expectedEnforceOnKey: "IP",
			expectedConfigsCount: 1,
			expectedForceSend:    []string{"EnforceOnKeyConfigs", "EnforceOnKey"},
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			got := expandSecurityPolicyRuleRateLimitOptions(tc.input)
			if got == nil {
				t.Fatalf("expected non-nil result")
			}
			if got.EnforceOnKey != tc.expectedEnforceOnKey {
				t.Errorf("expected EnforceOnKey %q, got %q", tc.expectedEnforceOnKey, got.EnforceOnKey)
			}
			if len(got.EnforceOnKeyConfigs) != tc.expectedConfigsCount {
				t.Errorf("expected %d EnforceOnKeyConfigs, got %d", tc.expectedConfigsCount, len(got.EnforceOnKeyConfigs))
			}
			if !reflect.DeepEqual(got.ForceSendFields, tc.expectedForceSend) {
				t.Errorf("expected ForceSendFields %v, got %v", tc.expectedForceSend, got.ForceSendFields)
			}
		})
	}
}
