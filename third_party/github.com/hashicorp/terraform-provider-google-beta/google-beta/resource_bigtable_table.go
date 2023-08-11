// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0
package google

import (
	"github.com/hashicorp/terraform-provider-google-beta/google-beta/services/bigtable"
)

func flattenColumnFamily(families []string) []map[string]interface{} {
	return bigtable.FlattenColumnFamily(families)
}
