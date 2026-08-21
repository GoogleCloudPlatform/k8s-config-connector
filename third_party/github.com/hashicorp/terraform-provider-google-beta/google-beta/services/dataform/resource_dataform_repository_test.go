// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0
package dataform_test

import (
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"

	"github.com/hashicorp/terraform-provider-google-beta/google-beta/acctest"
)

func TestAccDataformRepository_updated(t *testing.T) {
	t.Parallel()

	context := map[string]interface{}{
		"random_suffix": acctest.RandString(t, 10),
	}

	acctest.VcrTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.AccTestPreCheck(t) },
		ProtoV5ProviderFactories: acctest.ProtoV5ProviderBetaFactories(t),
		CheckDestroy:             testAccCheckDataformRepositoryDestroyProducer(t),
		Steps: []resource.TestStep{
			{
				Config: testAccDataformRepository_basic(context),
			},
			{
				ResourceName:            "google_dataform_repository.dataform_respository",
				ImportState:             true,
				ImportStateVerify:       true,
				ImportStateVerifyIgnore: []string{"region"},
			},
			{
				Config: testAccDataformRepository_updated(context),
			},
			{
				ResourceName:            "google_dataform_repository.dataform_respository",
				ImportState:             true,
				ImportStateVerify:       true,
				ImportStateVerifyIgnore: []string{"region"},
			},
		},
	})
}

func testAccDataformRepository_basic(context map[string]interface{}) string {
	return acctest.Nprintf(`
resource "google_sourcerepo_repository" "git_repository" {
  provider = google-beta
  name = "my/repository%{random_suffix}"
}

resource "google_secret_manager_secret" "secret" {
  provider = google-beta
  secret_id = "secret"

  replication {
    auto {}
  }
}

resource "google_secret_manager_secret_version" "secret_version" {
  provider = google-beta
  secret = google_secret_manager_secret.secret.id

  secret_data = "tf-test-secret-data%{random_suffix}"
}

resource "google_dataform_repository" "dataform_respository" {
  provider = google-beta
  name = "tf_test_dataform_repository%{random_suffix}"

  git_remote_settings {
      url = google_sourcerepo_repository.git_repository.url
      default_branch = "main"
      authentication_token_secret_version = google_secret_manager_secret_version.secret_version.id
  }

  workspace_compilation_overrides {
    default_database = "database"
    schema_suffix = "_suffix"
    table_prefix = "prefix_"
  }
}
`, context)
}

func testAccDataformRepository_updated(context map[string]interface{}) string {
	return acctest.Nprintf(`
resource "google_sourcerepo_repository" "git_repository" {
  provider = google-beta
  name = "my/repository%{random_suffix}"
}

resource "google_secret_manager_secret" "secret" {
  provider = google-beta
  secret_id = "secret"

  replication {
    auto {}
  }
}

resource "google_secret_manager_secret_version" "secret_version" {
  provider = google-beta
  secret = google_secret_manager_secret.secret.id

  secret_data = "tf-test-secret-data%{random_suffix}"
}

resource "google_dataform_repository" "dataform_respository" {
  provider = google-beta
  name = "tf_test_dataform_repository%{random_suffix}"

  git_remote_settings {
      url = google_sourcerepo_repository.git_repository.url
      default_branch = "main"
      authentication_token_secret_version = google_secret_manager_secret_version.secret_version.id
  }

  workspace_compilation_overrides {
    schema_suffix = "_suffix_v2"
    table_prefix = "prefix_v2_"
  }
}
`, context)
}
