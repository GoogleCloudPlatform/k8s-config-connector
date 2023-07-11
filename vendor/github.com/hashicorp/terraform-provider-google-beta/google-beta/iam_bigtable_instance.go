// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0
package google

import (
	"github.com/hashicorp/terraform-provider-google-beta/google-beta/services/bigtable"
	"github.com/hashicorp/terraform-provider-google-beta/google-beta/tpgiamresource"
	"github.com/hashicorp/terraform-provider-google-beta/google-beta/tpgresource"
	transport_tpg "github.com/hashicorp/terraform-provider-google-beta/google-beta/transport"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

var IamBigtableInstanceSchema = bigtable.IamBigtableInstanceSchema

func NewBigtableInstanceUpdater(d tpgresource.TerraformResourceData, config *transport_tpg.Config) (tpgiamresource.ResourceIamUpdater, error) {
	return bigtable.NewBigtableInstanceUpdater(d, config)
}

func BigtableInstanceIdParseFunc(d *schema.ResourceData, config *transport_tpg.Config) error {
	return bigtable.BigtableInstanceIdParseFunc(d, config)
}
