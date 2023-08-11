// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0
package google

import (
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"

	"github.com/hashicorp/terraform-provider-google-beta/google-beta/acctest"
)

func TestAccFirebaseExtensionsInstance_firebaseExtentionsInstanceResizeImageUpdate(t *testing.T) {
	t.Parallel()

	context := map[string]interface{}{
		"project_id":    acctest.GetTestProjectFromEnv(),
		"location":      "us-central1",
		"random_suffix": acctest.RandString(t, 10),
	}

	acctest.VcrTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.AccTestPreCheck(t) },
		ProtoV5ProviderFactories: acctest.ProtoV5ProviderBetaFactories(t),
		CheckDestroy:             testAccCheckFirebaseExtensionsInstanceDestroyProducer(t),
		Steps: []resource.TestStep{
			{
				Config: testAccFirebaseExtensionsInstance_firebaseExtentionsInstanceResizeImageBefore(context),
			},
			{
				ResourceName:            "google_firebase_extensions_instance.resize_image",
				ImportState:             true,
				ImportStateVerify:       true,
				ImportStateVerifyIgnore: []string{"instance_id"},
			},
			{
				Config: testAccFirebaseExtensionsInstance_firebaseExtentionsInstanceResizeImageAfter(context),
			},
			{
				ResourceName:            "google_firebase_extensions_instance.resize_image",
				ImportState:             true,
				ImportStateVerify:       true,
				ImportStateVerifyIgnore: []string{"instance_id"},
			},
		},
	})
}

func testAccFirebaseExtensionsInstance_firebaseExtentionsInstanceResizeImageBefore(context map[string]interface{}) string {
	return Nprintf(`
resource "google_storage_bucket" "images" {
  provider                    = google-beta
  project                     = "%{project_id}"
  name                        = "tf-test-bucket-id%{random_suffix}"
  location                    = "US"
  uniform_bucket_level_access = true

  # Delete all objects when the bucket is deleted
  force_destroy = true
}

resource "google_firebase_extensions_instance" "resize_image" {
  provider = google-beta
  project  = "%{project_id}"
  instance_id = "tf-test-storage-resize-images%{random_suffix}"
  config {
    extension_ref = "firebase/storage-resize-images"
    extension_version = "0.1.37"

    # The following params apply to the firebase/storage-resize-images extension. 
    # Different extensions may have different params
    params = {
      DELETE_ORIGINAL_FILE = false
      MAKE_PUBLIC          = false
      IMAGE_TYPE           = "jpeg"
      IS_ANIMATED          = true
      FUNCTION_MEMORY      = 1024
      DO_BACKFILL          = false
      IMG_SIZES            = "200x200"
      IMG_BUCKET           = google_storage_bucket.images.name
      LOCATION             = "%{location}"
    }

    system_params = {
      "firebaseextensions.v1beta.function/maxInstances"               = 3000
      "firebaseextensions.v1beta.function/memory"                     = 256
      "firebaseextensions.v1beta.function/minInstances"               = 0
      "firebaseextensions.v1beta.function/vpcConnectorEgressSettings" = "VPC_CONNECTOR_EGRESS_SETTINGS_UNSPECIFIED"
    }

    allowed_event_types = [
      "firebase.extensions.storage-resize-images.v1.complete"
    ]

    eventarc_channel = "projects/%{project_id}/locations/%{location}/channels/firebase"
  }
}
`, context)
}

func testAccFirebaseExtensionsInstance_firebaseExtentionsInstanceResizeImageAfter(context map[string]interface{}) string {
	return Nprintf(`
resource "google_storage_bucket" "images" {
  provider                    = google-beta
  project                     = "%{project_id}"
  name                        = "tf-test-bucket-id%{random_suffix}"
  location                    = "US"
  uniform_bucket_level_access = true

  # Delete all objects when the bucket is deleted
  force_destroy = true
}

resource "google_firebase_extensions_instance" "resize_image" {
  provider = google-beta
  project  = "%{project_id}"
  instance_id = "tf-test-storage-resize-images%{random_suffix}"
  config {
    extension_ref = "firebase/storage-resize-images"
    extension_version = "0.1.37"

    # The following params apply to the firebase/storage-resize-images extension. 
    # Different extensions may have different params
    params = {
      # Changed params
      DELETE_ORIGINAL_FILE = true
      MAKE_PUBLIC          = true
      IMAGE_TYPE           = "jpeg"
      IS_ANIMATED          = true
      FUNCTION_MEMORY      = 512
      DO_BACKFILL          = true
      IMG_SIZES            = "400x400"
      IMG_BUCKET           = google_storage_bucket.images.name
      LOCATION             = "%{location}"
    }

    system_params = {
      # Changed params
      "firebaseextensions.v1beta.function/maxInstances"               = 100
      "firebaseextensions.v1beta.function/memory"                     = 128
      "firebaseextensions.v1beta.function/minInstances"               = 0
      "firebaseextensions.v1beta.function/vpcConnectorEgressSettings" = "VPC_CONNECTOR_EGRESS_SETTINGS_UNSPECIFIED"
    }

    # Disable events
  }
}
`, context)
}
