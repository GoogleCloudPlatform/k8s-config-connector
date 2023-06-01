// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0
package google

import (
	"time"

	"github.com/hashicorp/terraform-provider-google-beta/google-beta/services/compute"
	transport_tpg "github.com/hashicorp/terraform-provider-google-beta/google-beta/transport"
)

// Deprecated: For backward compatibility ComputeOperationWaitTime is still working,
// but all new code should use ComputeOperationWaitTime in the compute package instead.
func ComputeOperationWaitTime(config *transport_tpg.Config, res interface{}, project, activity, userAgent string, timeout time.Duration) error {
	return compute.ComputeOperationWaitTime(config, res, project, activity, userAgent, timeout)
}

// Deprecated: For backward compatibility ComputeOrgOperationWaitTimeWithResponse is still working,
// but all new code should use ComputeOrgOperationWaitTimeWithResponse in the compute package instead.
func ComputeOrgOperationWaitTimeWithResponse(config *transport_tpg.Config, res interface{}, response *map[string]interface{}, parent, activity, userAgent string, timeout time.Duration) error {
	return compute.ComputeOrgOperationWaitTimeWithResponse(config, res, response, parent, activity, userAgent, timeout)
}
