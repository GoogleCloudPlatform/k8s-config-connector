// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0
package google

import (
	"github.com/hashicorp/terraform-provider-google-beta/google-beta/acctest"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
)

func testAccFirebaseDatabaseInstance_firebaseDatabaseInstanceInState(context map[string]interface{}, state string) string {
	context["desired_state"] = state
	return Nprintf(`
resource "google_firebase_database_instance" "updated" {
  provider = google-beta
  project  = "%{project_id}"
  region   = "%{region}"
  instance_id = "tf-test-state-change-db%{random_suffix}"
  desired_state  = "%{desired_state}"
}
`, context)
}

func TestAccFirebaseDatabaseInstance_firebaseDatabaseInstanceStateChange(t *testing.T) {
	t.Parallel()

	context := map[string]interface{}{
		"project_id":    acctest.GetTestProjectFromEnv(),
		"region":        acctest.GetTestRegionFromEnv(),
		"random_suffix": RandString(t, 10),
	}

	VcrTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.AccTestPreCheck(t) },
		ProtoV5ProviderFactories: ProtoV5ProviderBetaFactories(t),
		CheckDestroy:             testAccCheckFirebaseDatabaseInstanceDestroyProducer(t),
		Steps: []resource.TestStep{
			{
				Config: testAccFirebaseDatabaseInstance_firebaseDatabaseInstanceInState(context, "ACTIVE"),
			},
			{
				ResourceName:            "google_firebase_database_instance.updated",
				ImportState:             true,
				ImportStateVerify:       true,
				ImportStateVerifyIgnore: []string{"region", "instance_id", "desired_state"},
			},
			{
				Config: testAccFirebaseDatabaseInstance_firebaseDatabaseInstanceInState(context, "DISABLED"),
			},
			{
				ResourceName:            "google_firebase_database_instance.updated",
				ImportState:             true,
				ImportStateVerify:       true,
				ImportStateVerifyIgnore: []string{"region", "instance_id", "desired_state"},
			},
			{
				Config: testAccFirebaseDatabaseInstance_firebaseDatabaseInstanceInState(context, "ACTIVE"),
			},
			{
				ResourceName:            "google_firebase_database_instance.updated",
				ImportState:             true,
				ImportStateVerify:       true,
				ImportStateVerifyIgnore: []string{"region", "instance_id", "desired_state"},
			},
		},
	})
}
