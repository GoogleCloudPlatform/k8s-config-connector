package google

import (
	transport_tpg "github.com/hashicorp/terraform-provider-google-beta/google-beta/transport"
)

func MultiEnvSearch(ks []string) string {
	return transport_tpg.MultiEnvSearch(ks)
}
