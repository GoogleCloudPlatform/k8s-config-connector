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
resource "google_secret_manager_secret" "secret-basic" {
  secret_id = "secret-name"

  replication {
    automatic = true
  }
}

resource "google_secret_manager_secret_version" "secret-version-basic" {
  secret = google_secret_manager_secret.secret-basic.id

  secret_data = "secret-data"
}

resource "google_network_services_edge_cache_keyset" "keyset" {
  name        = "keyset-name"
  description = "The default keyset"
  public_key {
    id      = "my-public-key"
    managed = true
  }
  validation_shared_keys {
    secret_version = google_secret_manager_secret_version.secret-version-basic.id
  }
}

resource "google_network_services_edge_cache_origin" "instance" {
  name                 = "my-origin"
  origin_address       = "gs://media-edge-default"
  description          = "The default bucket for media edge test"
}

resource "google_network_services_edge_cache_service" "instance" {
  name                 = "my-service"
  description          = "some description"
  routing {
    host_rule {
      description = "host rule description"
      hosts = ["sslcert.tf-test.club"]
      path_matcher = "routes"
    }
    path_matcher {
      name = "routes"
      route_rule {
        description = "a route rule to match against master playlist"
        priority = 1
        match_rule {
          path_template_match = "/master.m3u8"
	}	
        origin = google_network_services_edge_cache_origin.instance.name
        route_action {
          cdn_policy {
	    signed_request_mode = "REQUIRE_TOKENS"
	    signed_request_keyset = google_network_services_edge_cache_keyset.keyset.id
	    signed_token_options {
	      token_query_parameter = "edge-cache-token"
	    }
	    signed_request_maximum_expiration_ttl = "600s"
	    add_signatures {
	      actions = ["GENERATE_COOKIE"]
	      keyset = google_network_services_edge_cache_keyset.keyset.id
	      copied_parameters = ["PathGlobs", "SessionID"]
	    }
          }
        }
      }
      route_rule {
        description = "a route rule to match against all playlists"
        priority = 2
        match_rule {
          path_template_match = "/*.m3u8"
        }
        origin = google_network_services_edge_cache_origin.instance.name
        route_action {
          cdn_policy {
	    signed_request_mode = "REQUIRE_TOKENS"
	    signed_request_keyset = google_network_services_edge_cache_keyset.keyset.id
	    signed_token_options {
	      token_query_parameter = "hdnts"
	      allowed_signature_algorithms = ["ED25519", "HMAC_SHA_256", "HMAC_SHA1"]
	    }
	    add_signatures {
	      actions = ["GENERATE_TOKEN_HLS_COOKIELESS"]
	      keyset = google_network_services_edge_cache_keyset.keyset.id
	      token_ttl = "1200s"
	      token_query_parameter = "hdntl"
	      copied_parameters = ["URLPrefix"]
	    }
          }
        }
      }
      route_rule {
        description = "a route rule to match against"
        priority = 3
        match_rule {
          path_template_match = "/**.m3u8"
        }
        origin = google_network_services_edge_cache_origin.instance.name
        route_action {
          cdn_policy {
	    signed_request_mode = "REQUIRE_TOKENS"
	    signed_request_keyset = google_network_services_edge_cache_keyset.keyset.id
	    signed_token_options {
	      token_query_parameter = "hdntl"
	    }
	    add_signatures {
	      actions = ["PROPAGATE_TOKEN_HLS_COOKIELESS"]
	      token_query_parameter = "hdntl"
	    }
          }
        }
      }
    }
  }
}
```
