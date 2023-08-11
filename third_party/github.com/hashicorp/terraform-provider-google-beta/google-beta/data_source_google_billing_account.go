// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0
package google

import (
	"github.com/hashicorp/terraform-provider-google-beta/google-beta/services/billing"
)

func canonicalBillingAccountName(ba string) string {
	return billing.CanonicalBillingAccountName(ba)
}
