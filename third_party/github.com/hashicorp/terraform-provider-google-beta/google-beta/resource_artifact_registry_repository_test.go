// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0
package google

import (
	"fmt"
	"github.com/hashicorp/terraform-provider-google-beta/google-beta/acctest"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
)

func TestAccArtifactRegistryRepository_update(t *testing.T) {
	t.Parallel()

	repositoryID := fmt.Sprintf("tf-test-%d", RandInt(t))

	VcrTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.AccTestPreCheck(t) },
		ProtoV5ProviderFactories: ProtoV5ProviderFactories(t),
		CheckDestroy:             testAccCheckArtifactRegistryRepositoryDestroyProducer(t),
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

func TestAccArtifactRegistryRepository_createMvnSnapshot(t *testing.T) {
	t.Parallel()

	repositoryID := fmt.Sprintf("tf-test-%d", RandInt(t))

	VcrTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.AccTestPreCheck(t) },
		ProtoV5ProviderFactories: ProtoV5ProviderFactories(t),
		CheckDestroy:             testAccCheckArtifactRegistryRepositoryDestroyProducer(t),
		Steps: []resource.TestStep{
			{
				Config: testAccArtifactRegistryRepository_createMvnWithVersionPolicy(repositoryID, "SNAPSHOT"),
			},
			{
				ResourceName:      "google_artifact_registry_repository.test",
				ImportState:       true,
				ImportStateVerify: true,
			},
		},
	})
}

func TestAccArtifactRegistryRepository_createMvnRelease(t *testing.T) {
	t.Parallel()

	repositoryID := fmt.Sprintf("tf-test-%d", RandInt(t))

	VcrTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.AccTestPreCheck(t) },
		ProtoV5ProviderFactories: ProtoV5ProviderFactories(t),
		CheckDestroy:             testAccCheckArtifactRegistryRepositoryDestroyProducer(t),
		Steps: []resource.TestStep{
			{
				Config: testAccArtifactRegistryRepository_createMvnWithVersionPolicy(repositoryID, "RELEASE"),
			},
			{
				ResourceName:      "google_artifact_registry_repository.test",
				ImportState:       true,
				ImportStateVerify: true,
			},
		},
	})
}

func TestAccArtifactRegistryRepository_kfp(t *testing.T) {
	t.Parallel()

	repositoryID := fmt.Sprintf("tf-test-%d", RandInt(t))

	VcrTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.AccTestPreCheck(t) },
		ProtoV5ProviderFactories: ProtoV5ProviderFactories(t),
		CheckDestroy:             testAccCheckArtifactRegistryRepositoryDestroyProducer(t),
		Steps: []resource.TestStep{
			{
				Config: testAccArtifactRegistryRepository_kfp(repositoryID),
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

func testAccArtifactRegistryRepository_createMvnWithVersionPolicy(repositoryID string, versionPolicy string) string {
	return fmt.Sprintf(`
resource "google_artifact_registry_repository" "test" {
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

func testAccArtifactRegistryRepository_kfp(repositoryID string) string {
	return fmt.Sprintf(`
resource "google_artifact_registry_repository" "test" {
  repository_id = "%s"
  location = "us-central1"
  description = "my-kfp-repository"
  format = "KFP"
}
`, repositoryID)
}

func TestAccArtifactRegistryRepository_virtual(t *testing.T) {
	t.Parallel()

	upstreamRepositoryID := fmt.Sprintf("tf-test-%d", RandInt(t))
	repositoryID := fmt.Sprintf("%s-virtual", upstreamRepositoryID)

	VcrTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.AccTestPreCheck(t) },
		ProtoV5ProviderFactories: ProtoV5ProviderFactories(t),
		CheckDestroy:             testAccCheckArtifactRegistryRepositoryDestroyProducer(t),
		Steps: []resource.TestStep{
			{
				Config: testAccArtifactRegistryRepository_virtual(repositoryID, upstreamRepositoryID, false),
			},
			{
				ResourceName:      "google_artifact_registry_repository.vr-test",
				ImportState:       true,
				ImportStateVerify: true,
			},
			{
				Config: testAccArtifactRegistryRepository_virtual(repositoryID, upstreamRepositoryID, true),
			},
			{
				ResourceName:      "google_artifact_registry_repository.vr-test",
				ImportState:       true,
				ImportStateVerify: true,
			},
		},
	})
}

func TestAccArtifactRegistryRepository_remote(t *testing.T) {
	t.Parallel()

	repositoryID := fmt.Sprintf("tf-test-%d", RandInt(t))

	VcrTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.AccTestPreCheck(t) },
		ProtoV5ProviderFactories: ProtoV5ProviderFactories(t),
		CheckDestroy:             testAccCheckArtifactRegistryRepositoryDestroyProducer(t),
		Steps: []resource.TestStep{
			{
				Config: testAccArtifactRegistryRepository_remote(repositoryID, "upstream"),
			},
			{
				ResourceName:      "google_artifact_registry_repository.rr-test",
				ImportState:       true,
				ImportStateVerify: true,
			},
			{
				Config: testAccArtifactRegistryRepository_remote(repositoryID, "docker hub"),
			},
			{
				ResourceName:      "google_artifact_registry_repository.rr-test",
				ImportState:       true,
				ImportStateVerify: true,
			},
		},
	})
}

func testAccArtifactRegistryRepository_virtual(repositoryID string, upstreamRepositoryID string, two_policies bool) string {
	policy_a := `
    upstream_policies {
      id = "upstream-a"
      repository = google_artifact_registry_repository.test-a.id
      priority = 1
    }
`
	policy_b := `
  upstream_policies {
    id = "upstream-b"
    repository = google_artifact_registry_repository.test-b.id
    priority = 2
  }
`
	if !two_policies {
		policy_b = ""
	}

	return fmt.Sprintf(`
resource "google_artifact_registry_repository" "test-a" {
  repository_id = "%s-a"
  location = "us-central1"
  description = "upstream repo"
  format = "DOCKER"
}
resource "google_artifact_registry_repository" "test-b" {
  repository_id = "%s-b"
  location = "us-central1"
  description = "alt upstream repo"
  format = "DOCKER"
}
resource "google_artifact_registry_repository" "vr-test" {
  repository_id = "%s"
  location = "us-central1"
  description = "virtual repo"
  format = "DOCKER"
  mode = "VIRTUAL_REPOSITORY"

  virtual_repository_config {
%s
%s
  }
}
`, upstreamRepositoryID, upstreamRepositoryID, repositoryID, policy_a, policy_b)
}

func testAccArtifactRegistryRepository_remote(repositoryID string, remoteDescription string) string {
	return fmt.Sprintf(`
resource "google_artifact_registry_repository" "rr-test" {
  repository_id = "%s"
  location = "us-central1"
  description = "remote repo"
  format = "DOCKER"
  mode = "REMOTE_REPOSITORY"

  remote_repository_config {
    description = "%s"
    docker_repository {
      public_repository = "DOCKER_HUB"
    }
  }
}
`, repositoryID, remoteDescription)
}
