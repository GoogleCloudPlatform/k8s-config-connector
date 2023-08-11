// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0
package google

import (
	"fmt"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-provider-google-beta/google-beta/acctest"
)

func TestAccContainerAnalysisNote_basic(t *testing.T) {
	t.Parallel()

	name := acctest.RandString(t, 10)
	readableName := acctest.RandString(t, 10)
	acctest.VcrTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.AccTestPreCheck(t) },
		ProtoV5ProviderFactories: acctest.ProtoV5ProviderFactories(t),
		CheckDestroy:             testAccCheckContainerAnalysisNoteDestroyProducer(t),
		Steps: []resource.TestStep{
			{
				Config: testAccContainerAnalysisNoteBasic(name, readableName),
			},
			{
				ResourceName:      "google_container_analysis_note.note",
				ImportState:       true,
				ImportStateVerify: true,
			},
		},
	})
}

func TestAccContainerAnalysisNote_update(t *testing.T) {
	t.Parallel()

	name := acctest.RandString(t, 10)
	readableName := acctest.RandString(t, 10)
	readableName2 := acctest.RandString(t, 10)
	acctest.VcrTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.AccTestPreCheck(t) },
		ProtoV5ProviderFactories: acctest.ProtoV5ProviderFactories(t),
		CheckDestroy:             testAccCheckContainerAnalysisNoteDestroyProducer(t),
		Steps: []resource.TestStep{
			{
				Config: testAccContainerAnalysisNoteBasic(name, readableName),
			},
			{
				ResourceName:      "google_container_analysis_note.note",
				ImportState:       true,
				ImportStateVerify: true,
			},
			{
				Config: testAccContainerAnalysisNoteBasic(name, readableName2),
			},
			{
				ResourceName:      "google_container_analysis_note.note",
				ImportState:       true,
				ImportStateVerify: true,
			},
		},
	})
}

func testAccContainerAnalysisNoteBasic(name, readableName string) string {
	return fmt.Sprintf(`
resource "google_container_analysis_note" "note" {
  name = "tf-test-%s"
  attestation_authority {
    hint {
      human_readable_name = "My Attestor %s"
    }
  }
}
`, name, readableName)
}
