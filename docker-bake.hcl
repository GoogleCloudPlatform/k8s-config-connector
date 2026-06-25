// Copyright 2026 Google LLC
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//      http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

group "default" {
  targets = ["manager", "recorder", "webhook", "deletiondefender", "unmanageddetector", "config-connector", "operator"]
}

variable "IMAGE_PREFIX" {
  default = ""
}

variable "IMAGE_TAG" {
  default = "latest"
}

variable "GKE_DISTROLESS_IMG" {
  default = "gcr.io/gke-release/gke-distroless/static:gke_distroless_20260207.00_p0"
}

variable "PLATFORM" {
  default = "linux/amd64"
}

target "builder" {
  dockerfile = "build/builder/Dockerfile"
  context    = "."
}

target "manager" {
  dockerfile = "build/manager/Dockerfile"
  context    = "."
  platforms  = split(",", PLATFORM)
  tags       = ["${IMAGE_PREFIX}controller:${IMAGE_TAG}"]
  contexts = {
    "builder:${IMAGE_TAG}" = "target:builder"
  }
  args = {
    BUILDER_IMG = "builder:${IMAGE_TAG}"
    GKE_DISTROLESS_IMG = "${GKE_DISTROLESS_IMG}"
  }
}

target "recorder" {
  dockerfile = "build/recorder/Dockerfile"
  context    = "."
  platforms  = split(",", PLATFORM)
  tags       = ["${IMAGE_PREFIX}recorder:${IMAGE_TAG}"]
  contexts = {
    "builder:${IMAGE_TAG}" = "target:builder"
  }
  args = {
    BUILDER_IMG = "builder:${IMAGE_TAG}"
    GKE_DISTROLESS_IMG = "${GKE_DISTROLESS_IMG}"
  }
}

target "webhook" {
  dockerfile = "build/webhook/Dockerfile"
  context    = "."
  platforms  = split(",", PLATFORM)
  tags       = ["${IMAGE_PREFIX}webhook:${IMAGE_TAG}"]
  contexts = {
    "builder:${IMAGE_TAG}" = "target:builder"
  }
  args = {
    BUILDER_IMG = "builder:${IMAGE_TAG}"
    GKE_DISTROLESS_IMG = "${GKE_DISTROLESS_IMG}"
  }
}

target "deletiondefender" {
  dockerfile = "build/deletiondefender/Dockerfile"
  context    = "."
  platforms  = split(",", PLATFORM)
  tags       = ["${IMAGE_PREFIX}deletiondefender:${IMAGE_TAG}"]
  contexts = {
    "builder:${IMAGE_TAG}" = "target:builder"
  }
  args = {
    BUILDER_IMG = "builder:${IMAGE_TAG}"
    GKE_DISTROLESS_IMG = "${GKE_DISTROLESS_IMG}"
  }
}

target "unmanageddetector" {
  dockerfile = "build/unmanageddetector/Dockerfile"
  context    = "."
  platforms  = split(",", PLATFORM)
  tags       = ["${IMAGE_PREFIX}unmanageddetector:${IMAGE_TAG}"]
  contexts = {
    "builder:${IMAGE_TAG}" = "target:builder"
  }
  args = {
    BUILDER_IMG = "builder:${IMAGE_TAG}"
    GKE_DISTROLESS_IMG = "${GKE_DISTROLESS_IMG}"
  }
}

target "config-connector" {
  dockerfile = "build/config-connector/Dockerfile"
  context    = "."
  platforms  = split(",", PLATFORM)
  tags       = ["${IMAGE_PREFIX}config-connector-cli:${IMAGE_TAG}"]
  contexts = {
    "builder:${IMAGE_TAG}" = "target:builder"
  }
  args = {
    BUILDER_IMG = "builder:${IMAGE_TAG}"
    GKE_DISTROLESS_IMG = "${GKE_DISTROLESS_IMG}"
  }
}

target "operator" {
  dockerfile = "operator/Dockerfile"
  context    = "."
  platforms  = split(",", PLATFORM)
  tags       = ["${IMAGE_PREFIX}operator:${IMAGE_TAG}"]
  args = {
    GKE_DISTROLESS_IMG = "${GKE_DISTROLESS_IMG}"
  }
}
