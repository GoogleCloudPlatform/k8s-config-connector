// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0
package google

import (
	"time"

	"github.com/hashicorp/terraform-provider-google-beta/google-beta/services/gkeonprem"
	transport_tpg "github.com/hashicorp/terraform-provider-google-beta/google-beta/transport"
)

// nolint: deadcode,unused
//
// Deprecated: For backward compatibility GkeonpremOperationWaitTimeWithResponse is still working,
// but all new code should use GkeonpremOperationWaitTimeWithResponse in the gkeonprem package instead.
func GkeonpremOperationWaitTimeWithResponse(config *transport_tpg.Config, op map[string]interface{}, response *map[string]interface{}, project, activity, userAgent string, timeout time.Duration) error {
	return gkeonprem.GkeonpremOperationWaitTimeWithResponse(config, op, response, project, activity, userAgent, timeout)
}

// Deprecated: For backward compatibility GkeonpremOperationWaitTime is still working,
// but all new code should use GkeonpremOperationWaitTime in the gkeonprem package instead.
func GkeonpremOperationWaitTime(config *transport_tpg.Config, op map[string]interface{}, project, activity, userAgent string, timeout time.Duration) error {
	return gkeonprem.GkeonpremOperationWaitTime(config, op, project, activity, userAgent, timeout)
}
