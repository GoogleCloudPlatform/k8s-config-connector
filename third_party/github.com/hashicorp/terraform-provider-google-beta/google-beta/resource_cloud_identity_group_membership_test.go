package google

import (
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
)

func TestAccCloudIdentityGroupMembership_update(t *testing.T) {
	t.Parallel()

	context := map[string]interface{}{
		"org_domain":    getTestOrgDomainFromEnv(t),
		"cust_id":       getTestCustIdFromEnv(t),
		"identity_user": getTestIdentityUserFromEnv(t),
		"random_suffix": randString(t, 10),
	}

	vcrTest(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		Providers:    testAccProviders,
		CheckDestroy: testAccCheckCloudIdentityGroupMembershipDestroyProducer(t),
		Steps: []resource.TestStep{
			{
				Config: testAccCloudIdentityGroupMembership_update1(context),
			},
			{
				ResourceName:      "google_cloud_identity_group_membership.basic",
				ImportState:       true,
				ImportStateVerify: true,
			},
			{
				Config: testAccCloudIdentityGroupMembership_update2(context),
			},
			{
				ResourceName:      "google_cloud_identity_group_membership.basic",
				ImportState:       true,
				ImportStateVerify: true,
			},
			{
				Config: testAccCloudIdentityGroupMembership_update1(context),
			},
			{
				ResourceName:      "google_cloud_identity_group_membership.basic",
				ImportState:       true,
				ImportStateVerify: true,
			},
		},
	})
}

func testAccCloudIdentityGroupMembership_update1(context map[string]interface{}) string {
	return Nprintf(`
resource "google_cloud_identity_group" "group" {
  display_name = "tf-test-my-identity-group%{random_suffix}"

  parent = "customers/%{cust_id}"

  group_key {
    id = "tf-test-my-identity-group%{random_suffix}@%{org_domain}"
  }

  labels = {
    "cloudidentity.googleapis.com/groups.discussion_forum" = ""
  }
}

resource "google_cloud_identity_group_membership" "basic" {
  group    = google_cloud_identity_group.group.id

  preferred_member_key {
    id = "%{identity_user}@%{org_domain}"
  }

  roles {
    name = "MEMBER"
  }

}
`, context)
}

func testAccCloudIdentityGroupMembership_update2(context map[string]interface{}) string {
	return Nprintf(`
resource "google_cloud_identity_group" "group" {
  display_name = "tf-test-my-identity-group%{random_suffix}"

  parent = "customers/%{cust_id}"

  group_key {
    id = "tf-test-my-identity-group%{random_suffix}@%{org_domain}"
  }

  labels = {
    "cloudidentity.googleapis.com/groups.discussion_forum" = ""
  }
}

resource "google_cloud_identity_group_membership" "basic" {
  group    = google_cloud_identity_group.group.id

  preferred_member_key {
    id = "%{identity_user}@%{org_domain}"
  }

  roles {
    name = "MEMBER"
  }

  roles {
    name = "MANAGER"
  }
}
`, context)
}

func TestAccCloudIdentityGroupMembership_import(t *testing.T) {
	t.Parallel()

	context := map[string]interface{}{
		"org_domain":    getTestOrgDomainFromEnv(t),
		"cust_id":       getTestCustIdFromEnv(t),
		"identity_user": getTestIdentityUserFromEnv(t),
		"random_suffix": randString(t, 10),
	}

	vcrTest(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		Providers:    testAccProviders,
		CheckDestroy: testAccCheckCloudIdentityGroupMembershipDestroyProducer(t),
		Steps: []resource.TestStep{
			{
				Config: testAccCloudIdentityGroupMembership_import(context),
			},
			{
				ResourceName:      "google_cloud_identity_group_membership.basic",
				ImportState:       true,
				ImportStateVerify: true,
			},
		},
	})
}

func testAccCloudIdentityGroupMembership_import(context map[string]interface{}) string {
	return Nprintf(`
resource "google_cloud_identity_group" "group" {
  display_name = "tf-test-my-identity-group%{random_suffix}"

  parent = "customers/%{cust_id}"

  group_key {
    id = "tf-test-my-identity-group%{random_suffix}@%{org_domain}"
  }

  labels = {
    "cloudidentity.googleapis.com/groups.discussion_forum" = ""
  }
}

resource "google_cloud_identity_group_membership" "basic" {
  group    = google_cloud_identity_group.group.id

  preferred_member_key {
    id = "%{identity_user}@%{org_domain}"
  }

  roles {
    name = "MEMBER"
  }

  roles {
    name = "MANAGER"
  }
}
`, context)
}

func TestAccCloudIdentityGroupMembership_cloudIdentityGroupMembershipWithMemberKey(t *testing.T) {
	t.Parallel()

	context := map[string]interface{}{
		"org_domain":    getTestOrgDomainFromEnv(t),
		"cust_id":       getTestCustIdFromEnv(t),
		"random_suffix": randString(t, 10),
	}

	vcrTest(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		Providers:    testAccProviders,
		CheckDestroy: testAccCheckCloudIdentityGroupMembershipDestroyProducer(t),
		Steps: []resource.TestStep{
			{
				Config: testAccCloudIdentityGroupMembership_cloudIdentityGroupMembershipWithMemberKey(context),
			},
			{
				ResourceName:      "google_cloud_identity_group_membership.cloud_identity_group_membership_basic",
				ImportState:       true,
				ImportStateVerify: true,
			},
		},
	})
}

func testAccCloudIdentityGroupMembership_cloudIdentityGroupMembershipWithMemberKey(context map[string]interface{}) string {
	return Nprintf(`
resource "google_cloud_identity_group" "group" {
  display_name = "tf-test-my-identity-group%{random_suffix}"

  parent = "customers/%{cust_id}"

  group_key {
  	id = "tf-test-my-identity-group%{random_suffix}@%{org_domain}"
  }

  labels = {
    "cloudidentity.googleapis.com/groups.discussion_forum" = ""
  }
}

resource "google_cloud_identity_group" "child-group" {
  display_name = "tf-test-my-identity-group%{random_suffix}-child"

  parent = "customers/%{cust_id}"

  group_key {
  	id = "tf-test-my-identity-group%{random_suffix}-child@%{org_domain}"
  }

  labels = {
    "cloudidentity.googleapis.com/groups.discussion_forum" = ""
  }
}

resource "google_cloud_identity_group_membership" "cloud_identity_group_membership_basic" {
  group    = google_cloud_identity_group.group.id

  member_key {
    id = google_cloud_identity_group.child-group.group_key[0].id
  }

  roles {
  	name = "MEMBER"
  }
}
`, context)
}

func TestAccCloudIdentityGroupMembership_cloudIdentityGroupMembershipUserWithMemberKey(t *testing.T) {
	t.Parallel()

	context := map[string]interface{}{
		"org_domain":    getTestOrgDomainFromEnv(t),
		"cust_id":       getTestCustIdFromEnv(t),
		"identity_user": getTestIdentityUserFromEnv(t),
		"random_suffix": randString(t, 10),
	}

	vcrTest(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		Providers:    testAccProviders,
		CheckDestroy: testAccCheckCloudIdentityGroupMembershipDestroyProducer(t),
		Steps: []resource.TestStep{
			{
				Config: testAccCloudIdentityGroupMembership_cloudIdentityGroupMembershipUserWithMemberKey(context),
			},
			{
				ResourceName:      "google_cloud_identity_group_membership.cloud_identity_group_membership_basic",
				ImportState:       true,
				ImportStateVerify: true,
			},
		},
	})
}

func testAccCloudIdentityGroupMembership_cloudIdentityGroupMembershipUserWithMemberKey(context map[string]interface{}) string {
	return Nprintf(`
resource "google_cloud_identity_group" "group" {
  display_name = "tf-test-my-identity-group%{random_suffix}"

  parent = "customers/%{cust_id}"

  group_key {
    id = "tf-test-my-identity-group%{random_suffix}@%{org_domain}"
  }

  labels = {
    "cloudidentity.googleapis.com/groups.discussion_forum" = ""
  }
}

resource "google_cloud_identity_group_membership" "cloud_identity_group_membership_basic" {
  group    = google_cloud_identity_group.group.id

  member_key {
    id = "%{identity_user}@%{org_domain}"
  }

  roles {
    name = "MEMBER"
  }

  roles {
    name = "MANAGER"
  }
}
`, context)
}
