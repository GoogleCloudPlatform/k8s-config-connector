// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0
package google

import (
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-provider-google-beta/google-beta/services/dataproc"
	"github.com/hashicorp/terraform-provider-google-beta/google-beta/tpgiamresource"
	"github.com/hashicorp/terraform-provider-google-beta/google-beta/tpgresource"
	transport_tpg "github.com/hashicorp/terraform-provider-google-beta/google-beta/transport"
)

var IamDataprocClusterSchema = dataproc.IamDataprocClusterSchema

func NewDataprocClusterUpdater(d tpgresource.TerraformResourceData, config *transport_tpg.Config) (tpgiamresource.ResourceIamUpdater, error) {
	return dataproc.NewDataprocClusterUpdater(d, config)
}

func DataprocClusterIdParseFunc(d *schema.ResourceData, config *transport_tpg.Config) error {
	return dataproc.DataprocClusterIdParseFunc(d, config)
}
