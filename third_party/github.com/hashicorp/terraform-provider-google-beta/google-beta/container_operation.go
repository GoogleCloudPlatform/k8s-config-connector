// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0
package google

import (
	"time"

	tpgcontainer "github.com/hashicorp/terraform-provider-google-beta/google-beta/services/container"
	transport_tpg "github.com/hashicorp/terraform-provider-google-beta/google-beta/transport"

	container "google.golang.org/api/container/v1beta1"
)

// Deprecated: For backward compatibility ContainerOperationWait is still working,
// but all new code should use ContainerOperationWait in the tpgcontainer package instead.
func ContainerOperationWait(config *transport_tpg.Config, op *container.Operation, project, location, activity, userAgent string, timeout time.Duration) error {
	return tpgcontainer.ContainerOperationWait(config, op, project, location, activity, userAgent, timeout)
}
