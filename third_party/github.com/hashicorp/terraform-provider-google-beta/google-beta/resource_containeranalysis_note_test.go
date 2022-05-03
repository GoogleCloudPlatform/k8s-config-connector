package google

import (
	"fmt"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
)

func TestAccContainerAnalysisNote_basic(t *testing.T) {
	t.Parallel()

	name := randString(t, 10)
	readableName := randString(t, 10)
	vcrTest(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		Providers:    testAccProviders,
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

	name := randString(t, 10)
	readableName := randString(t, 10)
	readableName2 := randString(t, 10)
	vcrTest(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		Providers:    testAccProviders,
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
