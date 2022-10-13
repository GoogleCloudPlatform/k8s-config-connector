resource "google_storage_bucket" "cc_cli" {
  force_destroy               = false
  location                    = "US"
  name                        = "cc-cli"
  project                     = "my-project"
  storage_class               = "STANDARD"
  uniform_bucket_level_access = true
}
# terraform import google_storage_bucket.cc_cli cc-cli
