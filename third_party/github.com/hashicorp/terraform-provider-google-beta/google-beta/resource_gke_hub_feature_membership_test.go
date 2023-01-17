package google

import (
	"context"
	"fmt"
	"testing"

	dcl "github.com/GoogleCloudPlatform/declarative-resource-client-library/dcl"
	gkehub "github.com/GoogleCloudPlatform/declarative-resource-client-library/services/google/gkehub/beta"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/terraform"
)

func TestAccGkeHubFeatureMembership_gkehubFeatureAcmUpdate(t *testing.T) {
	// Multiple fine-grained resources cause VCR to fail
	skipIfVcr(t)
	t.Parallel()

	context := map[string]interface{}{
		"random_suffix":   randString(t, 10),
		"org_id":          getTestOrgFromEnv(t),
		"billing_account": getTestBillingAccountFromEnv(t),
	}

	vcrTest(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		Providers:    testAccProvidersOiCS,
		CheckDestroy: testAccCheckGKEHubFeatureDestroyProducer(t),
		Steps: []resource.TestStep{
			{
				Config: testAccGkeHubFeatureMembership_gkehubFeatureAcmUpdateStart(context),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckGkeHubFeatureMembershipPresent(t, fmt.Sprintf("tf-test-gkehub%s", context["random_suffix"]), "global", "configmanagement", fmt.Sprintf("tf-test1%s", context["random_suffix"])),
					testAccCheckGkeHubFeatureMembershipPresent(t, fmt.Sprintf("tf-test-gkehub%s", context["random_suffix"]), "global", "configmanagement", fmt.Sprintf("tf-test2%s", context["random_suffix"])),
				),
			},
			{
				ResourceName:      "google_gke_hub_feature_membership.feature_member_1",
				ImportState:       true,
				ImportStateVerify: true,
			},
			{
				Config: testAccGkeHubFeatureMembership_gkehubFeatureAcmMembershipUpdate(context),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckGkeHubFeatureMembershipPresent(t, fmt.Sprintf("tf-test-gkehub%s", context["random_suffix"]), "global", "configmanagement", fmt.Sprintf("tf-test1%s", context["random_suffix"])),
					testAccCheckGkeHubFeatureMembershipPresent(t, fmt.Sprintf("tf-test-gkehub%s", context["random_suffix"]), "global", "configmanagement", fmt.Sprintf("tf-test2%s", context["random_suffix"])),
				),
			},
			{
				ResourceName:      "google_gke_hub_feature_membership.feature_member_2",
				ImportState:       true,
				ImportStateVerify: true,
			},
			{
				Config: testAccGkeHubFeatureMembership_gkehubFeatureAcmAddHierarchyController(context),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckGkeHubFeatureMembershipNotPresent(t, fmt.Sprintf("tf-test-gkehub%s", context["random_suffix"]), "global", "configmanagement", fmt.Sprintf("tf-test1%s", context["random_suffix"])),
					testAccCheckGkeHubFeatureMembershipPresent(t, fmt.Sprintf("tf-test-gkehub%s", context["random_suffix"]), "global", "configmanagement", fmt.Sprintf("tf-test2%s", context["random_suffix"])),
					testAccCheckGkeHubFeatureMembershipPresent(t, fmt.Sprintf("tf-test-gkehub%s", context["random_suffix"]), "global", "configmanagement", fmt.Sprintf("tf-test3%s", context["random_suffix"])),
				),
			},
			{
				ResourceName:      "google_gke_hub_feature_membership.feature_member_3",
				ImportState:       true,
				ImportStateVerify: true,
			},
			{
				Config: testAccGkeHubFeatureMembership_gkehubFeatureAcmRemoveFields(context),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckGkeHubFeatureMembershipNotPresent(t, fmt.Sprintf("tf-test-gkehub%s", context["random_suffix"]), "global", "configmanagement", fmt.Sprintf("tf-test2%s", context["random_suffix"])),
					testAccCheckGkeHubFeatureMembershipNotPresent(t, fmt.Sprintf("tf-test-gkehub%s", context["random_suffix"]), "global", "configmanagement", fmt.Sprintf("basic1%s", context["random_suffix"])),
					testAccCheckGkeHubFeatureMembershipPresent(t, fmt.Sprintf("tf-test-gkehub%s", context["random_suffix"]), "global", "configmanagement", fmt.Sprintf("tf-test3%s", context["random_suffix"])),
				),
			},
			{
				ResourceName:      "google_gke_hub_feature_membership.feature_member_3",
				ImportState:       true,
				ImportStateVerify: true,
			},
		},
	})
}

func testAccGkeHubFeatureMembership_gkehubFeatureAcmUpdateStart(context map[string]interface{}) string {
	return gkeHubFeatureProjectSetup(context) + gkeHubClusterMembershipSetup(context) + Nprintf(`
resource "google_gke_hub_feature" "feature" {
  project = google_project.project.project_id
  name = "configmanagement"
  location = "global"

  labels = {
    foo = "bar"
  }
  provider = google-beta
  depends_on = [google_project_service.mci, google_project_service.container, google_project_service.container, google_project_service.gkehub]
}

resource "google_gke_hub_feature_membership" "feature_member_1" {
  project = google_project.project.project_id
  location = "global"
  feature = google_gke_hub_feature.feature.name
  membership = google_gke_hub_membership.membership.membership_id
  configmanagement {
    version = "1.12.0"
    config_sync {
      source_format = "hierarchy"
      git {
        sync_repo   = "https://github.com/GoogleCloudPlatform/magic-modules"
        secret_type = "none"
      }
    }
  }
  provider = google-beta
}

resource "google_gke_hub_feature_membership" "feature_member_2" {
  project = google_project.project.project_id
  location = "global"
  feature = google_gke_hub_feature.feature.name
  membership = google_gke_hub_membership.membership_second.membership_id
  configmanagement {
    version = "1.12.0"
    config_sync {
      source_format = "hierarchy"
      git {
        sync_repo   = "https://github.com/terraform-providers/terraform-provider-google"
        secret_type = "none"
      }
    }
  }
  provider = google-beta
}
`, context)
}

func testAccGkeHubFeatureMembership_gkehubFeatureAcmMembershipUpdate(context map[string]interface{}) string {
	return gkeHubFeatureProjectSetup(context) + gkeHubClusterMembershipSetup(context) + Nprintf(`
resource "google_gke_hub_feature" "feature" {
  project = google_project.project.project_id
  name = "configmanagement"
  location = "global"

  labels = {
    foo = "changed"
  }
  provider = google-beta
  depends_on = [google_project_service.mci, google_project_service.container, google_project_service.container, google_project_service.gkehub]
}

resource "google_gke_hub_feature_membership" "feature_member_1" {
  project = google_project.project.project_id
  location = "global"
  feature = google_gke_hub_feature.feature.name
  membership = google_gke_hub_membership.membership.membership_id
  configmanagement {
    version = "1.12.0"
    config_sync {
      source_format = "hierarchy"
      git {
        sync_repo   = "https://github.com/GoogleCloudPlatform/magic-modules"
        secret_type = "none"
      }
    }
  }
  provider = google-beta
}

resource "google_gke_hub_feature_membership" "feature_member_2" {
  project = google_project.project.project_id
  location = "global"
  feature = google_gke_hub_feature.feature.name
  membership = google_gke_hub_membership.membership_second.membership_id
  configmanagement {
    version = "1.12.0"
    config_sync {
      source_format = "hierarchy"
      git {
        sync_repo   = "https://github.com/terraform-providers/terraform-provider-google-beta"
        secret_type = "none"
      }
    }
    policy_controller {
      enabled = true
      audit_interval_seconds = "10"
      exemptable_namespaces = ["asdf", "1234"]
      template_library_installed = true
    }
  }
  provider = google-beta
}
`, context)
}

func testAccGkeHubFeatureMembership_gkehubFeatureAcmAddHierarchyController(context map[string]interface{}) string {
	return gkeHubFeatureProjectSetup(context) + gkeHubClusterMembershipSetup(context) + Nprintf(`
resource "google_gke_hub_feature" "feature" {
  project = google_project.project.project_id
  name = "configmanagement"
  location = "global"

  labels = {
    foo = "changed"
  }
  provider = google-beta
  depends_on = [google_project_service.mci, google_project_service.container, google_project_service.container, google_project_service.gkehub]
}

resource "google_gke_hub_feature_membership" "feature_member_2" {
  project = google_project.project.project_id
  location = "global"
  feature = google_gke_hub_feature.feature.name
  membership = google_gke_hub_membership.membership_second.membership_id
  configmanagement {
    version = "1.12.0"
    config_sync {
      source_format = "unstructured"
      git {
        sync_repo   = "https://github.com/terraform-providers/terraform-provider-google-beta"
        secret_type = "none"
      }
    }
    policy_controller {
      enabled = true
      audit_interval_seconds = "9"
      exemptable_namespaces = ["different", "1234"]
      template_library_installed = false
    }
    hierarchy_controller {
      enable_hierarchical_resource_quota = true
      enable_pod_tree_labels = false
      enabled = true
    }
  }
  provider = google-beta
}

resource "google_gke_hub_feature_membership" "feature_member_3" {
  project = google_project.project.project_id
  location = "global"
  feature = google_gke_hub_feature.feature.name
  membership = google_gke_hub_membership.membership_third.membership_id
  configmanagement {
    version = "1.12.0"
    config_sync {
      source_format = "hierarchy"
      git {
        sync_repo   = "https://github.com/hashicorp/terraform"
        secret_type = "none"
      }
    }
    policy_controller {
      enabled = false
      audit_interval_seconds = "100"
      exemptable_namespaces = ["onetwothree", "fourfive"]
      template_library_installed = true
    }
    hierarchy_controller {
      enable_hierarchical_resource_quota = false
      enable_pod_tree_labels = true
      enabled = false
    }
  }
  provider = google-beta
}

resource "google_gke_hub_feature_membership" "feature_member_4" {
  project = google_project.project.project_id
  location = "global"
  feature = google_gke_hub_feature.feature.name
  membership = google_gke_hub_membership.membership_fourth.membership_id
  configmanagement {
    version = "1.12.0"
    policy_controller {
      enabled = true
      audit_interval_seconds = "100"
      template_library_installed = true
      mutation_enabled = true
      monitoring {
        backends = ["CLOUD_MONITORING", "PROMETHEUS"]
      }
    }
  }
  provider = google-beta
}



`, context)
}

func testAccGkeHubFeatureMembership_gkehubFeatureAcmRemoveFields(context map[string]interface{}) string {
	return gkeHubFeatureProjectSetup(context) + gkeHubClusterMembershipSetup(context) + Nprintf(`
resource "google_gke_hub_feature" "feature" {
  project = google_project.project.project_id
  name = "configmanagement"
  location = "global"

  labels = {
    foo = "changed"
  }
  provider = google-beta
  depends_on = [google_project_service.mci, google_project_service.container, google_project_service.container, google_project_service.gkehub]
}

resource "google_gke_hub_feature_membership" "feature_member_3" {
  project = google_project.project.project_id
  location = "global"
  feature = google_gke_hub_feature.feature.name
  membership = google_gke_hub_membership.membership_third.membership_id
  configmanagement {
    version = "1.12.0"
    policy_controller {
      enabled = true
      audit_interval_seconds = "100"
      exemptable_namespaces = ["onetwothree", "fourfive"]
      template_library_installed = true
    }
  }
  provider = google-beta
}
`, context)
}

func TestAccGkeHubFeatureMembership_gkehubFeatureAcmAllFields(t *testing.T) {
	// VCR fails to handle batched project services
	skipIfVcr(t)
	t.Parallel()

	context := map[string]interface{}{
		"random_suffix":   randString(t, 10),
		"org_id":          getTestOrgFromEnv(t),
		"billing_account": getTestBillingAccountFromEnv(t),
	}

	vcrTest(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		Providers:    testAccProvidersOiCS,
		CheckDestroy: testAccCheckGKEHubFeatureDestroyProducer(t),
		Steps: []resource.TestStep{
			{
				Config: testAccGkeHubFeatureMembership_gkehubFeatureAcmFewFields(context),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckGkeHubFeatureMembershipPresent(t, fmt.Sprintf("tf-test-gkehub%s", context["random_suffix"]), "global", "configmanagement", fmt.Sprintf("tf-test1%s", context["random_suffix"])),
				),
			},
			{
				ResourceName:      "google_gke_hub_feature_membership.feature_member",
				ImportState:       true,
				ImportStateVerify: true,
			},
			{
				Config: testAccGkeHubFeatureMembership_gkehubFeatureAcmAllFields(context),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckGkeHubFeatureMembershipPresent(t, fmt.Sprintf("tf-test-gkehub%s", context["random_suffix"]), "global", "configmanagement", fmt.Sprintf("tf-test1%s", context["random_suffix"])),
				),
			},
			{
				ResourceName:      "google_gke_hub_feature_membership.feature_member",
				ImportState:       true,
				ImportStateVerify: true,
			},
			{
				Config: testAccGkeHubFeatureMembership_gkehubFeatureAcmFewFields(context),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckGkeHubFeatureMembershipPresent(t, fmt.Sprintf("tf-test-gkehub%s", context["random_suffix"]), "global", "configmanagement", fmt.Sprintf("tf-test1%s", context["random_suffix"])),
				),
			},
			{
				ResourceName:      "google_gke_hub_feature_membership.feature_member",
				ImportState:       true,
				ImportStateVerify: true,
			},
			{
				Config: testAccGkeHubFeatureMembership_gkehubFeatureWithPreventDriftField(context),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckGkeHubFeatureMembershipPresent(t, fmt.Sprintf("tf-test-gkehub%s", context["random_suffix"]), "global", "configmanagement", fmt.Sprintf("tf-test1%s", context["random_suffix"])),
				),
			},
			{
				ResourceName:      "google_gke_hub_feature_membership.feature_member",
				ImportState:       true,
				ImportStateVerify: true,
			},
		},
	})
}

func testAccGkeHubFeatureMembership_gkehubFeatureAcmAllFields(context map[string]interface{}) string {
	return gkeHubFeatureProjectSetup(context) + Nprintf(`
resource "google_container_cluster" "primary" {
  project = google_project.project.project_id
  name               = "tf-test-cl%{random_suffix}"
  location           = "us-central1-a"
  initial_node_count = 1
  provider = google-beta
  depends_on = [google_project_service.mci, google_project_service.container, google_project_service.container, google_project_service.gkehub]
}

resource "google_gke_hub_membership" "membership" {
  project = google_project.project.project_id
  membership_id = "tf-test1%{random_suffix}"
  endpoint {
    gke_cluster {
      resource_link = "//container.googleapis.com/${google_container_cluster.primary.id}"
    }
  }
  description = "test resource."
  provider = google-beta
}

resource "google_gke_hub_feature" "feature" {
  project = google_project.project.project_id
  name = "configmanagement"
  location = "global"

  labels = {
    foo = "bar"
  }
  provider = google-beta
  depends_on = [google_project_service.mci, google_project_service.container, google_project_service.container, google_project_service.gkehub]
}

resource "google_gke_hub_feature_membership" "feature_member" {
  project = google_project.project.project_id
  location = "global"
  feature = google_gke_hub_feature.feature.name
  membership = google_gke_hub_membership.membership.membership_id
  configmanagement {
    version = "1.12.0"
    config_sync {
      git {
        sync_repo      = "https://github.com/hashicorp/terraform"
        https_proxy    = "https://example.com"
        policy_dir     = "google/"
        secret_type    = "none"
        sync_branch    = "some-branch"
        sync_rev       = "v3.60.0"
        sync_wait_secs = "30"
      }
    }
    policy_controller {
      enabled = true
      audit_interval_seconds = "100"
      exemptable_namespaces = ["onetwothree", "fourfive"]
      template_library_installed = true
      referential_rules_enabled = true
      log_denies_enabled = true
    }
  }
  provider = google-beta
}
`, context)
}

func testAccGkeHubFeatureMembership_gkehubFeatureWithPreventDriftField(context map[string]interface{}) string {
	return gkeHubFeatureProjectSetup(context) + Nprintf(`
resource "google_container_cluster" "primary" {
  project = google_project.project.project_id
  name               = "tf-test-cl%{random_suffix}"
  location           = "us-central1-a"
  initial_node_count = 1
  provider = google-beta
  depends_on = [google_project_service.mci, google_project_service.container, google_project_service.container, google_project_service.gkehub]
}

resource "google_gke_hub_membership" "membership" {
  project = google_project.project.project_id
  membership_id = "tf-test1%{random_suffix}"
  endpoint {
    gke_cluster {
      resource_link = "//container.googleapis.com/${google_container_cluster.primary.id}"
    }
  }
  description = "test resource."
  provider = google-beta
}

resource "google_gke_hub_feature" "feature" {
  project = google_project.project.project_id
  name = "configmanagement"
  location = "global"

  labels = {
    foo = "bar"
  }
  provider = google-beta
  depends_on = [google_project_service.mci, google_project_service.container, google_project_service.container, google_project_service.gkehub]
}

resource "google_gke_hub_feature_membership" "feature_member" {
  project = google_project.project.project_id
  location = "global"
  feature = google_gke_hub_feature.feature.name
  membership = google_gke_hub_membership.membership.membership_id
  configmanagement {
    version = "1.12.0"
    config_sync {
      git {
        sync_repo      = "https://github.com/hashicorp/terraform"
        https_proxy    = "https://example.com"
        policy_dir     = "google/"
        secret_type    = "none"
        sync_branch    = "some-branch"
        sync_rev       = "v3.60.0"
        sync_wait_secs = "30"
      }
      prevent_drift = true
    }
    policy_controller {
      enabled = true
      audit_interval_seconds = "100"
      exemptable_namespaces = ["onetwothree", "fourfive"]
      template_library_installed = true
      referential_rules_enabled = true
      log_denies_enabled = true
    }
  }
  provider = google-beta
}
`, context)
}

func testAccGkeHubFeatureMembership_gkehubFeatureAcmFewFields(context map[string]interface{}) string {
	return gkeHubFeatureProjectSetup(context) + Nprintf(`
resource "google_container_cluster" "primary" {
  project = google_project.project.project_id
  name               = "tf-test-cl%{random_suffix}"
  location           = "us-central1-a"
  initial_node_count = 1
  provider = google-beta
  depends_on = [google_project_service.mci, google_project_service.container, google_project_service.container, google_project_service.gkehub, google_project_service.acm]
}

resource "google_gke_hub_membership" "membership" {
  project = google_project.project.project_id
  membership_id = "tf-test1%{random_suffix}"
  endpoint {
    gke_cluster {
      resource_link = "//container.googleapis.com/${google_container_cluster.primary.id}"
    }
  }
  description = "test resource."
  provider = google-beta
}

resource "google_gke_hub_feature" "feature" {
  project = google_project.project.project_id
  name = "configmanagement"
  location = "global"

  labels = {
    foo = "bar"
  }
  provider = google-beta
  depends_on = [google_project_service.mci, google_project_service.container, google_project_service.container, google_project_service.gkehub]
}

resource "google_service_account" "feature_sa" {
  project = google_project.project.project_id
  account_id = "feature-sa"
  provider = google-beta
}

resource "google_gke_hub_feature_membership" "feature_member" {
  project = google_project.project.project_id
  location = "global"
  feature = google_gke_hub_feature.feature.name
  membership = google_gke_hub_membership.membership.membership_id
  configmanagement {
    version = "1.12.0"
    config_sync {
      git {
        sync_repo   = "https://github.com/hashicorp/terraform"
        secret_type = "none"
      }
    }
  }
  provider = google-beta
}
`, context)
}

func TestAccGkeHubFeatureMembership_gkehubFeatureAcmOci(t *testing.T) {
	// Multiple fine-grained resources cause VCR to fail
	skipIfVcr(t)
	t.Parallel()

	context := map[string]interface{}{
		"random_suffix":   randString(t, 10),
		"org_id":          getTestOrgFromEnv(t),
		"billing_account": getTestBillingAccountFromEnv(t),
	}

	vcrTest(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		Providers:    testAccProvidersOiCS,
		CheckDestroy: testAccCheckGKEHubFeatureDestroyProducer(t),
		Steps: []resource.TestStep{
			{
				Config: testAccGkeHubFeatureMembership_gkehubFeatureAcmOciStart(context),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckGkeHubFeatureMembershipPresent(t, fmt.Sprintf("tf-test-gkehub%s", context["random_suffix"]), "global", "configmanagement", fmt.Sprintf("tf-test1%s", context["random_suffix"])),
				),
			},
			{
				ResourceName:      "google_gke_hub_feature_membership.feature_member",
				ImportState:       true,
				ImportStateVerify: true,
			},
			{
				Config: testAccGkeHubFeatureMembership_gkehubFeatureAcmOciUpdate(context),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckGkeHubFeatureMembershipPresent(t, fmt.Sprintf("tf-test-gkehub%s", context["random_suffix"]), "global", "configmanagement", fmt.Sprintf("tf-test1%s", context["random_suffix"])),
				),
			},
			{
				ResourceName:      "google_gke_hub_feature_membership.feature_member",
				ImportState:       true,
				ImportStateVerify: true,
			},
			{
				Config: testAccGkeHubFeatureMembership_gkehubFeatureAcmOciRemoveFields(context),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckGkeHubFeatureMembershipPresent(t, fmt.Sprintf("tf-test-gkehub%s", context["random_suffix"]), "global", "configmanagement", fmt.Sprintf("tf-test1%s", context["random_suffix"])),
				),
			},
			{
				ResourceName:      "google_gke_hub_feature_membership.feature_member",
				ImportState:       true,
				ImportStateVerify: true,
			},
		},
	})
}

func testAccGkeHubFeatureMembership_gkehubFeatureAcmOciStart(context map[string]interface{}) string {
	return gkeHubFeatureProjectSetup(context) + gkeHubClusterMembershipSetup_ACMOCI(context) + Nprintf(`
resource "google_gke_hub_feature" "feature" {
  project = google_project.project.project_id
  name = "configmanagement"
  location = "global"

  labels = {
    foo = "bar"
  }
  provider = google-beta
  depends_on = [google_project_service.container, google_project_service.gkehub]
}

resource "google_service_account" "feature_sa" {
  project = google_project.project.project_id
  account_id = "feature-sa"
  provider = google-beta
}

resource "google_gke_hub_feature_membership" "feature_member" {
  project = google_project.project.project_id
  location = "global"
  feature = google_gke_hub_feature.feature.name
  membership = google_gke_hub_membership.membership_acmoci.membership_id
  configmanagement {
    version = "1.12.0"
    config_sync {
      source_format = "unstructured"
      oci {
        sync_repo = "us-central1-docker.pkg.dev/sample-project/config-repo/config-sync-gke:latest"
        policy_dir = "config-connector"
        sync_wait_secs = "20"
        secret_type = "gcpserviceaccount"
        gcp_service_account_email = google_service_account.feature_sa.email
      }
      prevent_drift = true
    }
    policy_controller {
      enabled = true
      audit_interval_seconds = "100"
      exemptable_namespaces = ["onetwothree", "fourfive"]
      template_library_installed = true
      referential_rules_enabled = true
      log_denies_enabled = true
    }
  }
  provider = google-beta
}
`, context)
}

func testAccGkeHubFeatureMembership_gkehubFeatureAcmOciUpdate(context map[string]interface{}) string {
	return gkeHubFeatureProjectSetup(context) + gkeHubClusterMembershipSetup_ACMOCI(context) + Nprintf(`
resource "google_gke_hub_feature" "feature" {
  project = google_project.project.project_id
  name = "configmanagement"
  location = "global"

  labels = {
    foo = "bar"
  }
  provider = google-beta
  depends_on = [google_project_service.container, google_project_service.gkehub]
}

resource "google_service_account" "feature_sa" {
  project = google_project.project.project_id
  account_id = "feature-sa"
  provider = google-beta
}

resource "google_gke_hub_feature_membership" "feature_member" {
  project = google_project.project.project_id
  location = "global"
  feature = google_gke_hub_feature.feature.name
  membership = google_gke_hub_membership.membership_acmoci.membership_id
  configmanagement {
    version = "1.12.0"
    config_sync {
      source_format = "hierarchy"
      oci {
        sync_repo = "us-central1-docker.pkg.dev/sample-project/config-repo/config-sync-gke:latest"
        policy_dir = "config-sync"
        sync_wait_secs = "15"
        secret_type = "gcenode"
        gcp_service_account_email = google_service_account.feature_sa.email
      }
      prevent_drift = true
    }
    policy_controller {
      enabled = true
      audit_interval_seconds = "100"
      exemptable_namespaces = ["onetwothree", "fourfive"]
      template_library_installed = true
      referential_rules_enabled = true
      log_denies_enabled = true
    }
  }
  provider = google-beta
}
`, context)
}

func testAccGkeHubFeatureMembership_gkehubFeatureAcmOciRemoveFields(context map[string]interface{}) string {
	return gkeHubFeatureProjectSetup(context) + gkeHubClusterMembershipSetup_ACMOCI(context) + Nprintf(`
resource "google_gke_hub_feature" "feature" {
  project = google_project.project.project_id
  name = "configmanagement"
  location = "global"

  labels = {
    foo = "bar"
  }
  provider = google-beta
  depends_on = [google_project_service.container, google_project_service.gkehub]
}

resource "google_service_account" "feature_sa" {
  project = google_project.project.project_id
  account_id = "feature-sa"
  provider = google-beta
}

resource "google_gke_hub_feature_membership" "feature_member" {
  project = google_project.project.project_id
  location = "global"
  feature = google_gke_hub_feature.feature.name
  membership = google_gke_hub_membership.membership_acmoci.membership_id
  configmanagement {
    version = "1.12.0"
    policy_controller {
      enabled = true
      audit_interval_seconds = "100"
      exemptable_namespaces = ["onetwothree", "fourfive"]
      template_library_installed = true
      referential_rules_enabled = true
      log_denies_enabled = true
    }
  }
  provider = google-beta
}
`, context)
}

func TestAccGkeHubFeatureMembership_gkehubFeatureMesh(t *testing.T) {
	// VCR fails to handle batched project services
	skipIfVcr(t)
	t.Parallel()

	context := map[string]interface{}{
		"random_suffix":   randString(t, 10),
		"org_id":          getTestOrgFromEnv(t),
		"billing_account": getTestBillingAccountFromEnv(t),
	}

	vcrTest(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		Providers:    testAccProvidersOiCS,
		CheckDestroy: testAccCheckGKEHubFeatureDestroyProducer(t),
		Steps: []resource.TestStep{
			{
				Config: testAccGkeHubFeatureMembership_meshStart(context),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckGkeHubFeatureMembershipPresent(t, fmt.Sprintf("tf-test-gkehub%s", context["random_suffix"]), "global", "servicemesh", fmt.Sprintf("tf-test1%s", context["random_suffix"])),
				),
			},
			{
				ResourceName:      "google_gke_hub_feature_membership.feature_member",
				ImportState:       true,
				ImportStateVerify: true,
			},
			{
				Config: testAccGkeHubFeatureMembership_meshUpdateManagement(context),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckGkeHubFeatureMembershipPresent(t, fmt.Sprintf("tf-test-gkehub%s", context["random_suffix"]), "global", "servicemesh", fmt.Sprintf("tf-test1%s", context["random_suffix"])),
				),
			},
			{
				ResourceName:      "google_gke_hub_feature_membership.feature_member",
				ImportState:       true,
				ImportStateVerify: true,
			},
			{
				Config: testAccGkeHubFeatureMembership_meshUpdateControlPlane(context),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckGkeHubFeatureMembershipPresent(t, fmt.Sprintf("tf-test-gkehub%s", context["random_suffix"]), "global", "servicemesh", fmt.Sprintf("tf-test1%s", context["random_suffix"])),
				),
			},
			{
				ResourceName:      "google_gke_hub_feature_membership.feature_member",
				ImportState:       true,
				ImportStateVerify: true,
			},
		},
	})
}

func testAccGkeHubFeatureMembership_meshStart(context map[string]interface{}) string {
	return gkeHubFeatureProjectSetup(context) + Nprintf(`
resource "google_container_cluster" "primary" {
  project = google_project.project.project_id
  name               = "tf-test-cl%{random_suffix}"
  location           = "us-central1-a"
  initial_node_count = 1
  provider = google-beta
  depends_on = [google_project_service.container, google_project_service.gkehub]
}

resource "google_gke_hub_membership" "membership" {
  project = google_project.project.project_id
  membership_id = "tf-test1%{random_suffix}"
  endpoint {
    gke_cluster {
      resource_link = "//container.googleapis.com/${google_container_cluster.primary.id}"
    }
  }
  description = "test resource."
  provider = google-beta
}

resource "google_gke_hub_feature" "feature" {
  project = google_project.project.project_id
  name = "servicemesh"
  location = "global"

  labels = {
    foo = "bar"
  }
  provider = google-beta
  depends_on = [google_project_service.container, google_project_service.gkehub, google_project_service.mesh]
}

resource "google_service_account" "feature_sa" {
  project = google_project.project.project_id
  account_id = "feature-sa"
  provider = google-beta
}

resource "google_gke_hub_feature_membership" "feature_member" {
  project = google_project.project.project_id
  location = "global"
  feature = google_gke_hub_feature.feature.name
  membership = google_gke_hub_membership.membership.membership_id
  mesh {
    management = "MANAGEMENT_AUTOMATIC"
    control_plane = "AUTOMATIC"
  }
  provider = google-beta
}
`, context)
}

func testAccGkeHubFeatureMembership_meshUpdateManagement(context map[string]interface{}) string {
	return gkeHubFeatureProjectSetup(context) + Nprintf(`
resource "google_container_cluster" "primary" {
  project = google_project.project.project_id
  name               = "tf-test-cl%{random_suffix}"
  location           = "us-central1-a"
  initial_node_count = 1
  provider = google-beta
  depends_on = [google_project_service.container, google_project_service.gkehub]
}

resource "google_gke_hub_membership" "membership" {
  project = google_project.project.project_id
  membership_id = "tf-test1%{random_suffix}"
  endpoint {
    gke_cluster {
      resource_link = "//container.googleapis.com/${google_container_cluster.primary.id}"
    }
  }
  description = "test resource."
  provider = google-beta
}

resource "google_gke_hub_feature" "feature" {
  project = google_project.project.project_id
  name = "servicemesh"
  location = "global"

  labels = {
    foo = "bar"
  }
  provider = google-beta
  depends_on = [google_project_service.container, google_project_service.gkehub, google_project_service.mesh]
}

resource "google_service_account" "feature_sa" {
  project = google_project.project.project_id
  account_id = "feature-sa"
  provider = google-beta
}

resource "google_gke_hub_feature_membership" "feature_member" {
  project = google_project.project.project_id
  location = "global"
  feature = google_gke_hub_feature.feature.name
  membership = google_gke_hub_membership.membership.membership_id
  mesh {
    management = "MANAGEMENT_MANUAL"
  }
  provider = google-beta
}
`, context)
}

func testAccGkeHubFeatureMembership_meshUpdateControlPlane(context map[string]interface{}) string {
	return gkeHubFeatureProjectSetup(context) + Nprintf(`
resource "google_container_cluster" "primary" {
  project = google_project.project.project_id
  name               = "tf-test-cl%{random_suffix}"
  location           = "us-central1-a"
  initial_node_count = 1
  provider = google-beta
  depends_on = [google_project_service.container, google_project_service.gkehub]
}

resource "google_gke_hub_membership" "membership" {
  project = google_project.project.project_id
  membership_id = "tf-test1%{random_suffix}"
  endpoint {
    gke_cluster {
      resource_link = "//container.googleapis.com/${google_container_cluster.primary.id}"
    }
  }
  description = "test resource."
  provider = google-beta
}

resource "google_gke_hub_feature" "feature" {
  project = google_project.project.project_id
  name = "servicemesh"
  location = "global"

  labels = {
    foo = "bar"
  }
  provider = google-beta
  depends_on = [google_project_service.container, google_project_service.gkehub, google_project_service.mesh]
}

resource "google_service_account" "feature_sa" {
  project = google_project.project.project_id
  account_id = "feature-sa"
  provider = google-beta
}

resource "google_gke_hub_feature_membership" "feature_member" {
  project = google_project.project.project_id
  location = "global"
  feature = google_gke_hub_feature.feature.name
  membership = google_gke_hub_membership.membership.membership_id
  mesh {
    control_plane = "MANUAL"
  }
  provider = google-beta
}
`, context)
}

func gkeHubClusterMembershipSetup(context map[string]interface{}) string {
	return Nprintf(`
resource "google_container_cluster" "primary" {
  name               = "tf-test-cl%{random_suffix}"
  location           = "us-central1-a"
  initial_node_count = 1
  project = google_project.project.project_id
  provider = google-beta
  depends_on = [google_project_service.container, google_project_service.container, google_project_service.gkehub]
}

resource "google_container_cluster" "secondary" {
  name               = "tf-test-cl2%{random_suffix}"
  location           = "us-central1-a"
  initial_node_count = 1
  project = google_project.project.project_id
  provider = google-beta
  depends_on = [google_project_service.container, google_project_service.container, google_project_service.gkehub]
}

resource "google_container_cluster" "tertiary" {
  name               = "tf-test-cl3%{random_suffix}"
  location           = "us-central1-a"
  initial_node_count = 1
  project = google_project.project.project_id
  provider = google-beta
  depends_on = [google_project_service.container, google_project_service.container, google_project_service.gkehub]
}


resource "google_container_cluster" "quarternary" {
  name               = "tf-test-cl4%{random_suffix}"
  location           = "us-central1-a"
  initial_node_count = 1
  project = google_project.project.project_id
  provider = google-beta
  depends_on = [google_project_service.container, google_project_service.container, google_project_service.gkehub]
}

resource "google_gke_hub_membership" "membership" {
  project = google_project.project.project_id
  membership_id = "tf-test1%{random_suffix}"
  endpoint {
    gke_cluster {
      resource_link = "//container.googleapis.com/${google_container_cluster.primary.id}"
    }
  }
  description = "test resource."
  provider = google-beta
}

resource "google_gke_hub_membership" "membership_second" {
  project = google_project.project.project_id
  membership_id = "tf-test2%{random_suffix}"
  endpoint {
    gke_cluster {
      resource_link = "//container.googleapis.com/${google_container_cluster.secondary.id}"
    }
  }
  description = "test resource."
  provider = google-beta
}

resource "google_gke_hub_membership" "membership_third" {
  project = google_project.project.project_id
  membership_id = "tf-test3%{random_suffix}"
  endpoint {
    gke_cluster {
      resource_link = "//container.googleapis.com/${google_container_cluster.tertiary.id}"
    }
  }
  description = "test resource."
  provider = google-beta
}

resource "google_gke_hub_membership" "membership_fourth" {
  project = google_project.project.project_id
  membership_id = "tf-test4%{random_suffix}"
  endpoint {
    gke_cluster {
      resource_link = "//container.googleapis.com/${google_container_cluster.quarternary.id}"
    }
  }
  description = "test resource."
  provider = google-beta
}
`, context)
}

func gkeHubClusterMembershipSetup_ACMOCI(context map[string]interface{}) string {
	return Nprintf(`

resource "google_compute_network" "testnetwork" {
    project                 = google_project.project.project_id
    name                    = "testnetwork"
    auto_create_subnetworks = true
    provider = google-beta
    depends_on = [google_project_service.compute]
}

resource "google_container_cluster" "container_acmoci" {
  name               = "tf-test-cl%{random_suffix}"
  location           = "us-central1-a"
  initial_node_count = 1
  network = google_compute_network.testnetwork.self_link
  project = google_project.project.project_id
  provider = google-beta
  depends_on = [google_project_service.container, google_project_service.container, google_project_service.gkehub]
}

resource "google_gke_hub_membership" "membership_acmoci" {
  project = google_project.project.project_id
  membership_id = "tf-test1%{random_suffix}"
  endpoint {
    gke_cluster {
      resource_link = "//container.googleapis.com/${google_container_cluster.container_acmoci.id}"
    }
  }
  description = "test resource."
  provider = google-beta
}
`, context)
}

func testAccCheckGkeHubFeatureMembershipPresent(t *testing.T, project, location, feature, membership string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		config := googleProviderConfig(t)
		obj := &gkehub.FeatureMembership{
			Feature:    dcl.StringOrNil(feature),
			Location:   dcl.StringOrNil(location),
			Membership: dcl.StringOrNil(membership),
			Project:    dcl.String(project),
		}

		_, err := NewDCLGkeHubClient(config, "", "", 0).GetFeatureMembership(context.Background(), obj)
		if err != nil {
			return err
		}
		return nil
	}
}

func testAccCheckGkeHubFeatureMembershipNotPresent(t *testing.T, project, location, feature, membership string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		config := googleProviderConfig(t)
		obj := &gkehub.FeatureMembership{
			Feature:    dcl.StringOrNil(feature),
			Location:   dcl.StringOrNil(location),
			Membership: dcl.StringOrNil(membership),
			Project:    dcl.String(project),
		}

		_, err := NewDCLGkeHubClient(config, "", "", 0).GetFeatureMembership(context.Background(), obj)
		if err == nil {
			return fmt.Errorf("Did not expect to find GKE Feature Membership for projects/%s/locations/%s/features/%s/membershipId/%s", project, location, feature, membership)
		}
		if dcl.IsNotFound(err) {
			return nil
		}
		return err
	}
}
