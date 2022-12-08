resource "google_project_iam_custom_role" "iamcustomrolesampleproject" {
  description = "This role only contains two permissions - publish and update"
  permissions = ["pubsub.topics.publish", "pubsub.topics.update"]
  project     = "my-project"
  role_id     = "iamcustomrolesampleproject"
  stage       = "GA"
  title       = "Example Project-Level Custom Role"
}
# terraform import google_project_iam_custom_role.iamcustomrolesampleproject my-project##iamcustomrolesampleproject
