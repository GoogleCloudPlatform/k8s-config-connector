package google

import (
	"github.com/hashicorp/terraform-plugin-sdk/v2/terraform"
	"github.com/hashicorp/terraform-provider-google-beta/google-beta/tpgresource"
	transport_tpg "github.com/hashicorp/terraform-provider-google-beta/google-beta/transport"
)

// This function isn't a test of transport.go; instead, it is used as an alternative
// to ReplaceVars inside tests.
//
// Deprecated: For backward compatibility replaceVarsForTest is still working,
// but all new code should use ReplaceVarsForTest in the tpgresource package instead.
func replaceVarsForTest(config *transport_tpg.Config, rs *terraform.ResourceState, linkTmpl string) (string, error) {
	return tpgresource.ReplaceVarsForTest(config, rs, linkTmpl)
}
