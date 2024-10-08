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

#/bin/bash
#export PROM_2_SD_VERSION="$(gcrane ls gke.gcr.io/prometheus-to-sd | egrep "prometheus-to-sd:v" | sort -r --version-sort | head -1)"
export PROM_2_SD_VERSION="$(gcrane ls gcr.io/gke-release/prometheus-to-sd | egrep "prometheus-to-sd:v" | sort -r --version-sort | head -1)"
echo "Switching to use the $PROM_2_SD_VERSION image"
for file in `find . -name '*.yaml'`
do 
	#if grep -q "gke.gcr.io/prometheus-to-sd" $file 
	if grep -q "gcr.io/gke-release/prometheus-to-sd" $file 
	then
		sed -i -E "s|gcr.io/gke-release/prometheus-to-sd:.*|$PROM_2_SD_VERSION|g" $file
	fi
done
