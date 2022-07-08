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
    version = "1.9.0"
    config_sync {
      source_format = "hierarchy"
      git {
        sync_repo = "https://github.com/GoogleCloudPlatform/magic-modules"
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
    version = "1.9.0"
    config_sync {
      source_format = "hierarchy"
      git {
        sync_repo = "https://github.com/terraform-providers/terraform-provider-google"
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
    version = "1.9.0"
    config_sync {
      source_format = "hierarchy"
      git {
        sync_repo = "https://github.com/GoogleCloudPlatform/magic-modules"
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
    version = "1.9.0"
    config_sync {
      source_format = "hierarchy"
      git {
        sync_repo = "https://github.com/terraform-providers/terraform-provider-google-beta"
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
    version = "1.9.0"
    config_sync {
      source_format = "unstructured"
      git {
        sync_repo = "https://github.com/terraform-providers/terraform-provider-google-beta"
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
    version = "1.9.0"
    config_sync {
      source_format = "hierarchy"
      git {
        sync_repo = "https://github.com/hashicorp/terraform"
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
    version = "1.9.0"
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
    version = "1.9.0"
    config_sync {
      git {
        sync_repo = "https://github.com/hashicorp/terraform"
        https_proxy = "https://example.com"
        policy_dir = "google/"
        sync_branch = "some-branch"
        sync_rev = "v3.60.0"
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
    version = "1.10.1"
    config_sync {
      git {
        sync_repo = "https://github.com/hashicorp/terraform"
        https_proxy = "https://example.com"
        policy_dir = "google/"
        sync_branch = "some-branch"
        sync_rev = "v3.60.0"
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
    version = "1.9.0"
    config_sync {
      git {
        sync_repo = "https://github.com/hashicorp/terraform"
      }
    }
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
