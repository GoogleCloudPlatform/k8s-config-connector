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
resource "google_iam_workforce_pool" "pool" {
  provider = google-beta

  workforce_pool_id = "example-pool"
  parent            = "organizations/123456789"
  location          = "global"
}

resource "google_iam_workforce_pool_provider" "example" {
  provider = google-beta

  workforce_pool_id  = google_iam_workforce_pool.pool.workforce_pool_id
  location           = google_iam_workforce_pool.pool.location
  provider_id        = "example-prvdr"
  attribute_mapping  = {
    "google.subject" = "assertion.sub"
  }
  oidc {
    issuer_uri       = "https://accounts.thirdparty.com"
    client_id        = "client-id"
    client_secret {
      value {
        plain_text = "client-secret"
      }
    }
    web_sso_config {
      response_type             = "ID_TOKEN"
      assertion_claims_behavior = "ONLY_ID_TOKEN_CLAIMS"
    }
    jwks_json        = "{\"keys\":[{\"kty\":\"RSA\",\"e\":\"AQAB\",\"use\":\"sig\",\"kid\":\"1i-PmZZrF1j2rOUAxkcQaaz3MnOXcwwziuch_XWjvqI\",\"alg\":\"RS256\",\"n\":\"kFpYE2Zm32y--cnUiFLm4cYmFO8tR4-5KU5-aqhRwiHPP0FkgdQZSoSyp_1DO6PruYfluRMviwOpbmM6LH7KemxVdxLKqLDkHSG0XC3dZkACRFNvBBOdFrvJ0ABXv3vVx592lFE0m-Je5-FerRSQCml6E7icNiTSxizEmvDsTIe8mvArjsODDrgWP25bEFwDPBd5cCl3_2gtW6YdaCRewLXdzuB5Wmp_vOu6trTUzEKbnQlWFtDDCPfOpywYXF8dY1Lbwas5iwwIZozwD2_CuTiyXa3T2_4oa119_rQrIC2BAv7q_S1Xoa2lk3q2GZUSVQ5i3gIbJuDHmp-6yh3k4w\"}]}"
  }
}
```
