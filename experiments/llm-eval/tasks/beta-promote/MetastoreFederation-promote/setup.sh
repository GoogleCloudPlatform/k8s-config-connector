#!/bin/bash
# Copyright 2025 Google LLC
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

set -e
SUBDIR="MetastoreFederation-promote-metastore"
git clone git@github.com:GoogleCloudPlatform/k8s-config-connector.git $SUBDIR
rm -rf $SUBDIR/.gemini # Avoid using git cloned .gemini.

export MCPWorkDir=${SUBDIR} # placeholder. evaluator runs setup.sh as a temp shell. This env var won't take effect when the shell finishes.
echo ${MCPWorkDir}