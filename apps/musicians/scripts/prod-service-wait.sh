#!/bin/bash

#Copyright 2019 Google LLC
#
#Licensed under the Apache License, Version 2.0 (the "License");
#you may not use this file except in compliance with the License.
#You may obtain a copy of the License at
#
#    https://www.apache.org/licenses/LICENSE-2.0
#
#Unless required by applicable law or agreed to in writing, software
#distributed under the License is distributed on an "AS IS" BASIS,
#WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
#See the License for the specific language governing permissions and
#limitations under the License.

while [[ $(kubectl get sqlinstance musicians-demo-prod -o json | jq --raw-output '.status.conditions[0].status') != "True" ]]; do
  echo "SQLInstance is not ready, waiting 15 seconds..."
  sleep 15
done


while [[
   $(
      curl -s http://$(kubectl get svc musicians-prod -o json | \
        jq --raw-output ".status.loadBalancer.ingress[0].ip")
    ) != *page*
  ]]; do
  echo "musicians-prod service not ready, retrying in 5 seconds"
  sleep 5
done
