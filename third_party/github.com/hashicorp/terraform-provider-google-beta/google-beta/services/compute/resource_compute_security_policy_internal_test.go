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
	ruleSchema := res.Schema["rule"]
	if ruleSchema.Set == nil {
		t.Errorf("expected rule Set hash function to be non-nil")
	}

	ruleElemSchema := ruleSchema.Elem.(*schema.Resource).Schema
	rateLimitOptionsSchema := ruleElemSchema["rate_limit_options"].Elem.(*schema.Resource).Schema
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
	expectedForceSend := []string{"EnforceOnKey", "EnforceOnKeyName", "EnforceOnKeyConfigs"}

	tests := []struct {
		name                 string
		input                []interface{}
		expectedEnforceOnKey string
		expectedConfigsCount int
	}{
		{
			name: "enforce_on_key_configs only",
			input: []interface{}{
				map[string]interface{}{
					"ban_threshold":           []interface{}{},
					"rate_limit_threshold":    threshold,
					"exceed_action":           "deny(429)",
					"conform_action":          "allow",
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
		},
		{
			name: "enforce_on_key only",
			input: []interface{}{
				map[string]interface{}{
					"ban_threshold":           []interface{}{},
					"rate_limit_threshold":    threshold,
					"exceed_action":           "deny(429)",
					"conform_action":          "allow",
					"ban_duration_sec":        0,
					"exceed_redirect_options": []interface{}{},
					"enforce_on_key":          "IP",
					"enforce_on_key_name":     "",
					"enforce_on_key_configs":  []interface{}{},
				},
			},
			expectedEnforceOnKey: "IP",
			expectedConfigsCount: 0,
		},
		{
			name: "neither configured - defaults to ALL",
			input: []interface{}{
				map[string]interface{}{
					"ban_threshold":           []interface{}{},
					"rate_limit_threshold":    threshold,
					"exceed_action":           "deny(429)",
					"conform_action":          "allow",
					"ban_duration_sec":        0,
					"exceed_redirect_options": []interface{}{},
					"enforce_on_key":          "",
					"enforce_on_key_name":     "",
					"enforce_on_key_configs":  []interface{}{},
				},
			},
			expectedEnforceOnKey: "ALL",
			expectedConfigsCount: 0,
		},
		{
			name: "both configured",
			input: []interface{}{
				map[string]interface{}{
					"ban_threshold":           []interface{}{},
					"rate_limit_threshold":    threshold,
					"exceed_action":           "deny(429)",
					"conform_action":          "allow",
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
			if !reflect.DeepEqual(got.ForceSendFields, expectedForceSend) {
				t.Errorf("expected ForceSendFields %v, got %v", expectedForceSend, got.ForceSendFields)
			}
		})
	}
}

func TestResourceComputeSecurityPolicyRuleHash(t *testing.T) {
	makeRule := func(enforceOnKey string, configs []interface{}) map[string]interface{} {
		return map[string]interface{}{
			"action":      "throttle",
			"priority":    1000,
			"description": "rate limit rule",
			"preview":     false,
			"match": []interface{}{
				map[string]interface{}{
					"versioned_expr": "SRC_IPS_V1",
					"config": []interface{}{
						map[string]interface{}{
							"src_ip_ranges": schema.NewSet(schema.HashString, []interface{}{"*"}),
						},
					},
				},
			},
			"rate_limit_options": []interface{}{
				map[string]interface{}{
					"conform_action":       "allow",
					"exceed_action":        "deny(429)",
					"enforce_on_key":       enforceOnKey,
					"enforce_on_key_name":  "",
					"enforce_on_key_configs": configs,
					"rate_limit_threshold": []interface{}{
						map[string]interface{}{
							"count":        100,
							"interval_sec": 60,
						},
					},
				},
			},
		}
	}

	ruleOmitted := makeRule("", []interface{}{})
	ruleDefaultAll := makeRule("ALL", []interface{}{})
	ruleIP := makeRule("IP", []interface{}{})
	ruleWithConfigs := makeRule("", []interface{}{
		map[string]interface{}{
			"enforce_on_key_type": "IP",
			"enforce_on_key_name": "",
		},
	})

	hashOmitted := resourceComputeSecurityPolicyRuleHash(ruleOmitted)
	hashDefaultAll := resourceComputeSecurityPolicyRuleHash(ruleDefaultAll)
	hashIP := resourceComputeSecurityPolicyRuleHash(ruleIP)
	hashWithConfigs := resourceComputeSecurityPolicyRuleHash(ruleWithConfigs)

	if hashOmitted != hashDefaultAll {
		t.Errorf("expected omitted enforce_on_key hash (%d) to equal default ALL hash (%d)", hashOmitted, hashDefaultAll)
	}
	if hashOmitted == hashIP {
		t.Errorf("expected omitted enforce_on_key hash (%d) to differ from IP hash (%d)", hashOmitted, hashIP)
	}
	if hashOmitted == hashWithConfigs {
		t.Errorf("expected omitted enforce_on_key hash (%d) to differ from enforce_on_key_configs hash (%d)", hashOmitted, hashWithConfigs)
	}
}

func TestEffectiveEnforceOnKey(t *testing.T) {
	oMapDefaultAll := map[string]interface{}{
		"enforce_on_key":         "ALL",
		"enforce_on_key_configs": []interface{}{},
	}
	nMapOmitted := map[string]interface{}{
		"enforce_on_key":         "",
		"enforce_on_key_configs": []interface{}{},
	}
	nMapWithConfigs := map[string]interface{}{
		"enforce_on_key": "",
		"enforce_on_key_configs": []interface{}{
			map[string]interface{}{"enforce_on_key_type": "IP"},
		},
	}

	if effectiveEnforceOnKey(oMapDefaultAll) != effectiveEnforceOnKey(nMapOmitted) {
		t.Errorf("expected no diff between ALL and omitted when configs are empty")
	}
	if effectiveEnforceOnKey(oMapDefaultAll) == effectiveEnforceOnKey(nMapWithConfigs) {
		t.Errorf("expected diff when transitioning from ALL to enforce_on_key_configs")
	}
}
