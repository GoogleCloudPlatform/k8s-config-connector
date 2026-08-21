/**
 * Copyright 2022 Google LLC
 *
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *
 *      http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 */

```hcl
data "google_billing_account" "account" {
  billing_account = "000000-0000000-0000000-000000"
}

resource "google_billing_budget" "budget" {
  billing_account = data.google_billing_account.account.id
  display_name = "Example Billing Budget"

  amount {
    specified_amount {
      currency_code = "USD"
      units = "100000"
    }
  }

  all_updates_rule {
    disable_default_iam_recipients = true
    pubsub_topic = google_pubsub_topic.budget.id
  }
}

resource "google_pubsub_topic" "budget" {
  name = "example-topic"
}
```
