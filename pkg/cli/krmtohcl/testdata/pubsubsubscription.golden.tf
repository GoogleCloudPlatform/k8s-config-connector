resource "google_pubsub_subscription" "pubsubsubscription_sample" {
  ack_deadline_seconds = 15

  expiration_policy {
    ttl = "2678400s"
  }

  labels = {
    cnrm-lease-expiration = "1603984859"
    cnrm-lease-holder-id  = "btpp498colih6qs1pe5g"
    label-one             = "value-one"
  }

  message_retention_duration = "86400s"
  name                       = "pubsubsubscription-sample"
  project                    = "my-project"
  topic                      = "projects/my-project/topics/pubsubsubscription-dep"
}
# terraform import google_pubsub_subscription.pubsubsubscription_sample projects/my-project/subscriptions/pubsubsubscription-sample
