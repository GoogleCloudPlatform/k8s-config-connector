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
resource "google_iam_workload_identity_pool" "pool" {
  workload_identity_pool_id = "example-pool"
}

resource "google_iam_workload_identity_pool_provider" "example" {
  workload_identity_pool_id          = google_iam_workload_identity_pool.pool.workload_identity_pool_id
  workload_identity_pool_provider_id = "example-prvdr"
  display_name                       = "Name of provider"
  description                        = "OIDC identity pool provider for automated test"
  disabled                           = true
  attribute_condition                = "\"e968c2ef-047c-498d-8d79-16ca1b61e77e\" in assertion.groups"
  attribute_mapping                  = {
    "google.subject"                  = "\"azure::\" + assertion.tid + \"::\" + assertion.sub"
    "attribute.tid"                   = "assertion.tid"
    "attribute.managed_identity_name" = <<EOT
      {
        "8bb39bdb-1cc5-4447-b7db-a19e920eb111":"workload1",
        "55d36609-9bcf-48e0-a366-a3cf19027d2a":"workload2"
      }[assertion.oid]
EOT
  }
  oidc {
    allowed_audiences = ["https://example.com/gcp-oidc-federation", "example.com/gcp-oidc-federation"]
    issuer_uri        = "https://sts.windows.net/azure-tenant-id"
    jwks_json         = "{\"keys\":[{\"kty\":\"RSA\",\"alg\":\"RS256\",\"kid\":\"sif0AR-F6MuvksAyAOv-Pds08Bcf2eUMlxE30NofddA\",\"use\":\"sig\",\"e\":\"AQAB\",\"n\":\"ylH1Chl1tpfti3lh51E1g5dPogzXDaQseqjsefGLknaNl5W6Wd4frBhHyE2t41Q5zgz_Ll0-NvWm0FlaG6brhrN9QZu6sJP1bM8WPfJVPgXOanxi7d7TXCkeNubGeiLTf5R3UXtS9Lm_guemU7MxDjDTelxnlgGCihOVTcL526suNJUdfXtpwUsvdU6_ZnAp9IpsuYjCtwPm9hPumlcZGMbxstdh07O4y4O90cVQClJOKSGQjAUCKJWXIQ0cqffGS_HuS_725CPzQ85SzYZzaNpgfhAER7kx_9P16ARM3BJz0PI5fe2hECE61J4GYU_BY43sxDfs7HyJpEXKLU9eWw\"}]}"
  }
}
```
