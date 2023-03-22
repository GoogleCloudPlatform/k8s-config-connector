package google

import (
	"fmt"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
)

func TestAccContainerAnalysisNote_basic(t *testing.T) {
	t.Parallel()

	name := RandString(t, 10)
	readableName := RandString(t, 10)
	VcrTest(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		Providers:    TestAccProviders,
		CheckDestroy: testAccCheckContainerAnalysisNoteDestroyProducer(t),
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

	name := RandString(t, 10)
	readableName := RandString(t, 10)
	readableName2 := RandString(t, 10)
	VcrTest(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		Providers:    TestAccProviders,
		CheckDestroy: testAccCheckContainerAnalysisNoteDestroyProducer(t),
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
