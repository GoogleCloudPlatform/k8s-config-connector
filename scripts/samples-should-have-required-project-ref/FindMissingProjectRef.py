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

# Sample usage: python3 FindMissingProjectRef.py /usr/local/google/home/shuxiancai/go/src/github.com/GoogleCloudPlatform/k8s-config-connector
import glob
import sys
import yaml

rootPath = sys.argv[1]
kindsWithRequiredProjectRef = set()
for filepath in glob.iglob(rootPath + '/config/crds/resources/*.yaml'):
  with open(filepath) as f:
    dataMap = yaml.safe_load(f)
    schemaprops = dataMap["spec"]["versions"][0]["schema"]["openAPIV3Schema"]["properties"]
    if "spec" in schemaprops.keys():
      spec = dataMap["spec"]["versions"][0]["schema"]["openAPIV3Schema"]["properties"]["spec"]
      properties = spec["properties"]
      if "projectRef" in properties.keys() and "required" in spec.keys():
        required = spec["required"]
        if "projectRef" in required:
          kindsWithRequiredProjectRef.add(dataMap["spec"]["names"]["kind"])

for filepath in glob.glob(rootPath + '/config/samples/resources/**/*.yaml', recursive=True):
  with open(filepath) as f:
    dataMaps = yaml.safe_load_all(f)
    for dataMap in dataMaps:
      kind = dataMap["kind"]
      if kind in kindsWithRequiredProjectRef and "spec" in dataMap.keys() and "projectRef" not in dataMap["spec"]:
        print(filepath)
