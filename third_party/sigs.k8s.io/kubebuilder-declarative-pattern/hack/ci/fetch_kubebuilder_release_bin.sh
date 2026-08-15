#!/bin/bash

# Copyright 2020 The Kubernetes Authors.
#
# Licensed under the Apache License, Version 2.0 (the "License");
# you may not use this file except in compliance with the License.
# You may obtain a copy of the License at
#
#     http://www.apache.org/licenses/LICENSE-2.0
#
# Unless required by applicable law or agreed to in writing, software
# distributed under the License is distributed on an "AS IS" BASIS,
# WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
# See the License for the specific language governing permissions and
# limitations under the License.


# This shell fetches kubebuilder release binary and puts it under
# /tmp/kubebuilder_bin/kubebuilder/bin/
# Then sets KUBEBUILDER_ASSETS to that path to make go test using
# envtest package in kubernetes-sigs/controller-runtime work properly.
#
# This script is largely based on CI script file "chek-everything.sh"
# of kubernetes-sigs/controller-runtime

# Version of kubebuilder release binary which this script attempts to fetch
kubebuilder_release_version="2.3.1"

goos=$(go env GOHOSTOS)
goarch=$(go env GOHOSTARCH)

# Supported arch under above release version
# See source code of standard library "build" ("go/build") for os&arch constant
os_darwin='darwin'
darwin_supported_arch=('amd64')

# Supported arch under above release version
# See source code of standard library "build" ("go/build") for os&arch constant
os_linux='linux'
linux_supported_arch=('amd64' 'arm64' 'ppc64le')

kubebuilder_release_base_url='https://github.com/kubernetes-sigs/kubebuilder/releases/download'
fetch_url_base="${kubebuilder_release_base_url}/v${kubebuilder_release_version}"


# Fetch kubebuilder release binaries and place them under "<argument>/kubebuilder/bin/"
function fetch_kb_bin () {
  local dest_dir="${1}/kubebuilder"
  local archive_file
  local checksum_file="checksums.txt"
  local found
  
  # Check whether kubebuilder release binary is provided for current environment's OS & Arch
  if [[ ${goos} == ${os_darwin} ]]; then
    for (( i=0; i < ${#darwin_supported_arch[@]}; i++ )); do
      if [[ ${goarch} == ${darwin_supported_arch[${i}]} ]]; then
        found="true"
        break
      fi
    done
  
    if [[ -z ${found} ]]; then
      echo "No kubebuilder supported arch under ${goos}: ${goarch}"
      exit 1
    fi
  
  elif [[ ${goos} == ${os_linux} ]]; then
    for (( i=0; i < ${#linux_supported_arch[@]}; i++ )); do
      if [[ ${goarch} == ${linux_supported_arch[${i}]} ]]; then
        found="true"
        break
      fi
    done
  
    if [[ -z ${found} ]]; then
      echo "No kubebuilder supported arch under ${goos}: ${goarch}"
      exit 1
    fi
  
  else
    echo "Not kubebuilder supported os: ${goos}"
    exit 1
  fi
  
  archive_file="kubebuilder_${kubebuilder_release_version}_${goos}_${goarch}.tar.gz"
  
  mkdir -p ${dest_dir}

  # Fetch kubebuilder release binary
  curl -fsL "${fetch_url_base}/${archive_file}" -o "/tmp/${archive_file}"
  if [[ ${?} != 0 ]]; then
    echo "Failed to curl kubebuilder release binary"
    exit 1
  fi

  # Fetch checksum file
  curl -fsL "${fetch_url_base}/${checksum_file}" -o "/tmp/${checksum_file}"
  if [[ ${?} != 0 ]]; then
    echo "Failed to curl checksum file"
    exit 1
  fi
 
  # Hash & check fetched binary
  (
  cd /tmp
  sha256sum -c --quiet --ignore-missing "/tmp/${checksum_file}"
  if [[ ${?} != 0 ]]; then
    echo "Maybe, kubebuilder release binary got broken & doesn't match hash value in ${checksum_file}"
    exit 1
  fi
  )

  # Unpack
  tar -C ${dest_dir} -xzf "/tmp/${archive_file}" --strip-components=1
  
  # Check expected binaries exist
  if [[ !(-x ${dest_dir}/bin/etcd && -x ${dest_dir}/bin/kubectl && -x ${dest_dir}/bin/kube-apiserver) ]]; then
    echo "There are no expected binaries at ${dset_dir}/bin : etcd, kubectl & kube-apiserver"
    exit 1
  fi

  # Set KUBEBUILDER_ASSETS environment variable
  export KUBEBUILDER_ASSETS="${dest_dir}/bin"
}

fetch_kb_bin "/tmp"
