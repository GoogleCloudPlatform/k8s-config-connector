package google

import (
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
)

func TestAccCloudFunctions2Function_update(t *testing.T) {
	t.Parallel()

	context := map[string]interface{}{
		"zip_path":      "./test-fixtures/cloudfunctions2/function-source.zip",
		"random_suffix": randString(t, 10),
	}

	vcrTest(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		Providers:    testAccProvidersOiCS,
		CheckDestroy: testAccCheckCloudfunctions2functionDestroyProducer(t),
		Steps: []resource.TestStep{
			{
				Config: testAccCloudfunctions2function_basic(context),
			},
			{
				ResourceName:            "google_cloudfunctions2_function.terraform-test2",
				ImportState:             true,
				ImportStateVerify:       true,
				ImportStateVerifyIgnore: []string{"location", "build_config.0.source.0.storage_source.0.object", "build_config.0.source.0.storage_source.0.bucket"},
			},
			{
				Config: testAccCloudFunctions2Function_test_update(context),
			},
			{
				ResourceName:            "google_cloudfunctions2_function.terraform-test2",
				ImportState:             true,
				ImportStateVerify:       true,
				ImportStateVerifyIgnore: []string{"location", "build_config.0.source.0.storage_source.0.object", "build_config.0.source.0.storage_source.0.bucket"},
			},
			{
				Config: testAccCloudFunctions2Function_test_redeploy(context),
			},
			{
				ResourceName:            "google_cloudfunctions2_function.terraform-test2",
				ImportState:             true,
				ImportStateVerify:       true,
				ImportStateVerifyIgnore: []string{"location", "build_config.0.source.0.storage_source.0.object", "build_config.0.source.0.storage_source.0.bucket"},
			},
		},
	})
}

func testAccCloudfunctions2function_basic(context map[string]interface{}) string {
	return Nprintf(`
resource "google_storage_bucket" "bucket" {
  provider = google-beta
  name     = "tf-test-cloudfunctions2-function-bucket%{random_suffix}"
  location = "US"
  uniform_bucket_level_access = true
}
 
resource "google_storage_bucket_object" "object" {
  provider = google-beta
  name   = "function-source.zip"
  bucket = google_storage_bucket.bucket.name
  source = "%{zip_path}"
}
 
resource "google_cloudfunctions2_function" "terraform-test2" {
  provider = google-beta
  name = "tf-test-test-function%{random_suffix}"
  location = "us-central1"
  description = "a new function"
 
  build_config {
    runtime = "nodejs12"
    entry_point = "helloHttp"
    source {
      storage_source {
        bucket = google_storage_bucket.bucket.name
        object = google_storage_bucket_object.object.name
      }
    }
  }
 
  service_config {
    max_instance_count  = 1
    available_memory    = "1536Mi"
    timeout_seconds     = 30
  }
}
`, context)
}

func testAccCloudFunctions2Function_test_update(context map[string]interface{}) string {
	return Nprintf(`
resource "google_storage_bucket" "bucket" {
  provider = google-beta
  name     = "tf-test-cloudfunctions2-function-bucket%{random_suffix}"
  location = "US"
  uniform_bucket_level_access = true
}
 
resource "google_storage_bucket_object" "object" {
  provider = google-beta
  name   = "function-source.zip"
  bucket = google_storage_bucket.bucket.name
  source = "%{zip_path}"
}
 
resource "google_cloudfunctions2_function" "terraform-test2" {
  provider = google-beta
  name = "tf-test-test-function%{random_suffix}"
  location = "us-central1"
  description = "an updated function"
 
  build_config {
    runtime = "nodejs12"
    entry_point = "helloHttp"
    source {
      storage_source {
        bucket = google_storage_bucket.bucket.name
        object = google_storage_bucket_object.object.name
      }
    }
  }
 
  service_config {
    max_instance_count  = 1
    available_memory    = "1536Mi"
    timeout_seconds     = 30
  }
}
`, context)
}

func testAccCloudFunctions2Function_test_redeploy(context map[string]interface{}) string {
	return Nprintf(`
resource "google_storage_bucket" "bucket" {
  provider = google-beta
  name     = "tf-test-cloudfunctions2-function-bucket%{random_suffix}"
  location = "US"
  uniform_bucket_level_access = true
}
 
resource "google_storage_bucket_object" "object" {
  provider = google-beta
  name   = "function-source.zip"
  bucket = google_storage_bucket.bucket.name
  source = "%{zip_path}"
}
 
resource "google_cloudfunctions2_function" "terraform-test2" {
  provider = google-beta
  name = "tf-test-test-function%{random_suffix}"
  location = "us-west1"
  description = "function test"
 
  build_config {
    runtime = "nodejs16"
    entry_point = "helloHttp"
    environment_variables = {
        BUILD_CONFIG_TEST = "build_test"
    }
    source {
      storage_source {
        bucket = google_storage_bucket.bucket.name
        object = google_storage_bucket_object.object.name
      }
    }
  }
 
  service_config {
    max_instance_count  = 5
    min_instance_count = 1
    available_memory    = "256M"
    timeout_seconds     = 60
    environment_variables = {
        SERVICE_CONFIG_TEST = "build_test"
    }
  }
}
`, context)
}
