#!/usr/bin/env bash
# Copyright 2022 Google LLC
#
# Licensed under the Apache License, Version 2.0 (the "License");
# you may not use this file except in compliance with the License.
# You may obtain a copy of the License at
#
#      http://www.apache.org/licenses/LICENSE-2.0
#
# Unless required by applicable law or agreed to in writing, software
# distributed under the License is distributed on an "AS IS" BASIS,
# WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
# See the License for the specific language governing permissions and
# limitations under the License.


# Paths
REPO_ROOT="$(git rev-parse --show-toplevel)"

# general purpose folders
BIN_DIR=bin

# binary names
CONFIG_CONNECTOR_BINARY_NAME=config-connector

# dependency versions
KUBEBUILDER_VERSION=2.3.1
KUBEAPISERVER_VERSION=1.21.0
KUSTOMIZE_VERSION=3.5.4

