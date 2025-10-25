#!/usr/bin/env bash
# Copyright 2018 The Kubernetes Authors.
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

set -o errexit
set -o nounset
set -o pipefail

# needs KUBEBUILDER_VERSION
REPO_ROOT="$(git rev-parse --show-toplevel)"

source ${REPO_ROOT}/scripts/shared-vars-public.sh

# Enable tracing in this script off by setting the TRACE variable in your
# environment to any value:
#
# $ TRACE=1 test.sh
TRACE=${TRACE:-""}
if [ -n "$TRACE" ]; then
  set -x
fi

goarch=amd64
goos="unknown"

if [[ "$OSTYPE" == "linux-gnu" ]]; then
  goos="linux"
elif [[ "$OSTYPE" == "darwin"* ]]; then
  goos="darwin"
fi

if [[ "$goos" == "unknown" ]]; then
  echo "OS '$OSTYPE' not supported. Aborting." >&2
  exit 1
fi

# Turn colors in this script off by setting the NO_COLOR variable in your
# environment to any value:
#
# $ NO_COLOR=1 test.sh
NO_COLOR=${NO_COLOR:-""}
if [ -z "$NO_COLOR" ]; then
  header=$'\e[1;33m'
  reset=$'\e[0m'
else
  header=''
  reset=''
fi

function header_text {
  echo "$header$*$reset"
}

rc=0
tmp_root=/tmp

kb_root_dir=$tmp_root/kubebuilder
kube_bin_dir=$tmp_root/kube/bin
kustomize_bin_dir=$tmp_root/kustomize/bin
kb_orig=$(pwd)

# Skip fetching and untaring the tools by setting the SKIP_FETCH_TOOLS variable
# in your environment to any value:
#
# $ SKIP_FETCH_TOOLS=1 ./fetch_ext_bins.sh
#
# If you skip fetching tools, this script will use the tools already on your
# machine, but rebuild the kubebuilder and kubebuilder-bin binaries.
SKIP_FETCH_TOOLS=${SKIP_FETCH_TOOLS:-""}

function prepare_staging_dir {
  header_text "preparing staging dir"

  if [ -z "$SKIP_FETCH_TOOLS" ]; then
    rm -rf "$kb_root_dir"
  else
    rm -f "$kb_root_dir/kubebuilder/bin/kubebuilder"
    rm -f "$kb_root_dir/kubebuilder/bin/kubebuilder-gen"
    rm -f "$kb_root_dir/kubebuilder/bin/vendor.tar.gz"
  fi
}

# fetch tools for build:
# * k8s API gen tools and make it available under kb_root_dir/bin.
function fetch_tools {
  if [ -n "$SKIP_FETCH_TOOLS" ]; then
    return 0
  fi

  header_text "fetching kubebuilder"

  kb_tools_archive_name="kubebuilder-${KUBEBUILDER_VERSION}.tar.gz"
  kb_tools_download_url="https://github.com/kubernetes-sigs/kubebuilder/releases/download/v${KUBEBUILDER_VERSION}/kubebuilder_${KUBEBUILDER_VERSION}_${goos}_${goarch}.tar.gz"

  kb_tools_archive_path="$tmp_root/$kb_tools_archive_name"
  if [ ! -f $kb_tools_archive_path ]; then
    curl -sL --retry 5 ${kb_tools_download_url} -o "$kb_tools_archive_path"
  fi
  mkdir -p "$kb_root_dir"
  tar -zvxf "$kb_tools_archive_path" -C "$kb_root_dir" --strip-components=1

  # Download kustomize
  curl -O -L https://github.com/kubernetes-sigs/kustomize/releases/download/kustomize%2Fv${KUSTOMIZE_VERSION}/kustomize_v${KUSTOMIZE_VERSION}_${goos}_${goarch}.tar.gz
  tar -zxvf kustomize_v${KUSTOMIZE_VERSION}_${goos}_${goarch}.tar.gz
  mkdir -p $kustomize_bin_dir
  mv kustomize $kustomize_bin_dir/kustomize
  chmod u+x $kustomize_bin_dir/kustomize
}

function setup_envs {
  header_text "setting up env vars"

  # Setup env vars
  export PATH=$PATH:/tmp/kubebuilder/bin
  export PATH=$PATH:$kustomize_bin_dir
  export TEST_ASSET_KUBECTL=/tmp/kubebuilder/bin/kubectl
  export TEST_ASSET_KUBE_APISERVER=$kube_bin_dir/kube-apiserver
  export TEST_ASSET_ETCD=/tmp/kubebuilder/bin/etcd
}
