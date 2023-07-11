// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0
package google

import (
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"

	"github.com/hashicorp/terraform-provider-google-beta/google-beta/services/logging"
	transport_tpg "github.com/hashicorp/terraform-provider-google-beta/google-beta/transport"
)

var OrganizationLoggingExclusionSchema = logging.OrganizationLoggingExclusionSchema

func NewOrganizationLoggingExclusionUpdater(d *schema.ResourceData, config *transport_tpg.Config) (logging.ResourceLoggingExclusionUpdater, error) {
	return logging.NewOrganizationLoggingExclusionUpdater(d, config)
}

func OrganizationLoggingExclusionIdParseFunc(d *schema.ResourceData, _ *transport_tpg.Config) error {
	return logging.OrganizationLoggingExclusionIdParseFunc(d, nil)
}
