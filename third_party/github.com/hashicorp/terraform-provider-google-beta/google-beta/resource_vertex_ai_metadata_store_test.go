// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0
package google

import (
	"fmt"
	"github.com/hashicorp/terraform-provider-google-beta/google-beta/acctest"
	"strings"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/terraform"
	"github.com/hashicorp/terraform-provider-google-beta/google-beta/tpgresource"
	transport_tpg "github.com/hashicorp/terraform-provider-google-beta/google-beta/transport"
)

func TestAccVertexAIMetadataStore_vertexAiMetadataStoreExample(t *testing.T) {
	t.Parallel()

	kms := BootstrapKMSKeyInLocation(t, "us-central1")
	name := fmt.Sprintf("tf-test-%s", RandString(t, 10))

	VcrTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.AccTestPreCheck(t) },
		ProtoV5ProviderFactories: ProtoV5ProviderFactories(t),
		CheckDestroy:             testAccCheckVertexAIMetadataStoreDestroyProducer(t),
		Steps: []resource.TestStep{
			{
				Config: testAccVertexAIMetadataStore_vertexAiMetadataStoreExample(name, kms.CryptoKey.Name),
			},
			{
				ResourceName:            "google_vertex_ai_metadata_store.store",
				ImportState:             true,
				ImportStateVerify:       true,
				ImportStateVerifyIgnore: []string{"region"},
			},
		},
	})
}

func testAccVertexAIMetadataStore_vertexAiMetadataStoreExample(name, kmsKey string) string {
	return fmt.Sprintf(`
resource "google_vertex_ai_metadata_store" "store" {
  name          = "%s"
  description   = "Magic"
  region        = "us-central1"
  encryption_spec {
	kms_key_name = "%s"
  }
}
`, name, kmsKey)
}

func testAccCheckVertexAIMetadataStoreDestroyProducer(t *testing.T) func(s *terraform.State) error {
	return func(s *terraform.State) error {
		for name, rs := range s.RootModule().Resources {
			if rs.Type != "google_vertex_ai_metadata_store" {
				continue
			}
			if strings.HasPrefix(name, "data.") {
				continue
			}

			config := GoogleProviderConfig(t)

			url, err := tpgresource.ReplaceVarsForTest(config, rs, "{{VertexAIBasePath}}{{name}}")
			if err != nil {
				return err
			}

			billingProject := ""

			if config.BillingProject != "" {
				billingProject = config.BillingProject
			}

			_, err = transport_tpg.SendRequest(transport_tpg.SendRequestOptions{
				Config:    config,
				Method:    "GET",
				Project:   billingProject,
				RawURL:    url,
				UserAgent: config.UserAgent,
			})
			if err == nil {
				return fmt.Errorf("VertexAIMetadataStore still exists at %s", url)
			}
		}

		return nil
	}
}
