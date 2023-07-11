// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0
package google

import (
	tpgcompute "github.com/hashicorp/terraform-provider-google-beta/google-beta/services/compute"
	transport_tpg "github.com/hashicorp/terraform-provider-google-beta/google-beta/transport"
)

// Deprecated: For backward compatibility PollCheckInstanceConfigDeleted is still working,
// but all new code should use PollCheckInstanceConfigDeleted in the tpgcompute package instead.
func PollCheckInstanceConfigDeleted(resp map[string]interface{}, respErr error) transport_tpg.PollResult {
	return tpgcompute.PollCheckInstanceConfigDeleted(resp, respErr)
}
