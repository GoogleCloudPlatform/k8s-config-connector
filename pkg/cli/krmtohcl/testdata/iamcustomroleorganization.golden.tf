resource "google_iam_custom_role" "iamcustomrolesampleorganization" {
  description = "This role only contains two permissions - publish and update"
  org_id      = "1234567"
  permissions = ["pubsub.topics.publish", "pubsub.topics.update"]
  role_id     = "iamcustomrolesampleorganization"
  stage       = "GA"
  title       = "Example Organization-Level Custom Role Created by Config Connector"
}
# terraform import google_iam_custom_role.iamcustomrolesampleorganization #1234567#iamcustomrolesampleorganization
