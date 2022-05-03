package google

import (
	"fmt"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
)

func TestAccArtifactRegistryRepository_update(t *testing.T) {
	t.Parallel()

	repositoryID := fmt.Sprintf("tf-test-%d", randInt(t))

	vcrTest(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		Providers:    testAccProvidersOiCS,
		CheckDestroy: testAccCheckArtifactRegistryRepositoryDestroyProducer(t),
		Steps: []resource.TestStep{
			{
				Config: testAccArtifactRegistryRepository_update(repositoryID),
			},
			{
				ResourceName:      "google_artifact_registry_repository.test",
				ImportState:       true,
				ImportStateVerify: true,
			},
			{
				Config: testAccArtifactRegistryRepository_update2(repositoryID),
			},
			{
				ResourceName:      "google_artifact_registry_repository.test",
				ImportState:       true,
				ImportStateVerify: true,
			},
		},
	})
}

func TestAccArtifactRegistryRepository_create_mvn_snapshot(t *testing.T) {
	t.Parallel()

	repositoryID := fmt.Sprintf("tf-test-%d", randInt(t))

	vcrTest(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		Providers:    testAccProvidersOiCS,
		CheckDestroy: testAccCheckArtifactRegistryRepositoryDestroyProducer(t),
		Steps: []resource.TestStep{
			{
				Config: testAccArtifactRegistryRepository_create(repositoryID, "SNAPSHOT"),
			},
			{
				ResourceName:      "google_artifact_registry_repository.test",
				ImportState:       true,
				ImportStateVerify: true,
			},
		},
	})
}

func TestAccArtifactRegistryRepository_create_mvn_release(t *testing.T) {
	t.Parallel()

	repositoryID := fmt.Sprintf("tf-test-%d", randInt(t))

	vcrTest(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		Providers:    testAccProvidersOiCS,
		CheckDestroy: testAccCheckArtifactRegistryRepositoryDestroyProducer(t),
		Steps: []resource.TestStep{
			{
				Config: testAccArtifactRegistryRepository_create(repositoryID, "RELEASE"),
			},
			{
				ResourceName:      "google_artifact_registry_repository.test",
				ImportState:       true,
				ImportStateVerify: true,
			},
		},
	})
}

func testAccArtifactRegistryRepository_update(repositoryID string) string {
	return fmt.Sprintf(`
resource "google_artifact_registry_repository" "test" {
  provider = google-beta

  repository_id = "%s"
  location = "us-central1"
  description = "pre-update"
  format = "DOCKER"

  labels = {
    my_key    = "my_val"
    other_key = "other_val"
  }
}
`, repositoryID)
}

func testAccArtifactRegistryRepository_update2(repositoryID string) string {
	return fmt.Sprintf(`
resource "google_artifact_registry_repository" "test" {
  provider = google-beta

  repository_id = "%s"
  location = "us-central1"
  description = "post-update"
  format = "DOCKER"

  labels = {
    my_key    = "my_val"
    other_key = "new_val"
  }
}
`, repositoryID)
}

func testAccArtifactRegistryRepository_create(repositoryID string, versionPolicy string) string {
	return fmt.Sprintf(`
resource "google_artifact_registry_repository" "test" {
  provider = google-beta

  repository_id = "%s"
  location = "us-central1"
  description = "post-update"
  format = "MAVEN"
  maven_config {
    version_policy = "%s"
  }
}
`, repositoryID, versionPolicy)
}
