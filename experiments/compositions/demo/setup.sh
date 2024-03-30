#!/bin/bash
# Copyright 2024 Google LLC
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


scriptpath=$(realpath $0)
base=$(dirname $scriptpath)

# we are demoing on CC in config-control namespace
# create Composition CRDs
kubectl apply -f $base/../composition/config/crd/bases/composition.google.com_compositions.yaml
kubectl apply -f $base/../composition/config/crd/bases/composition.google.com_contexts.yaml
kubectl apply -f $base/../composition/config/crd/bases/composition.google.com_plans.yaml
kubectl apply -f $base/../alice/config/crd/bases/alice.alice_cloudsqls.yaml
kubectl apply -f $base/../alice/config/crd/bases/alice.alice_appteams.yaml

kubectl apply -f composition-hasql.yaml  # create composition CR
kubectl apply -f composition-appteam.yaml  # create composition CR
kubectl apply -f context.yaml      # create context CR