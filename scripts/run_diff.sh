#!/bin/bash
set -e
set -x
if [ -z "$1" ]; then
  echo "Must provide 1 argument - name of resource to diff, e.g. 'google_compute_forwarding_rule'"
  exit 1
fi

function cleanup() {
  go mod edit -dropreplace=github.com/hashicorp/terraform-provider-clean-google-beta
  go mod edit -droprequire=github.com/hashicorp/terraform-provider-clean-google-beta
}

trap cleanup EXIT
if [[ -d ~/go/src/github.com/hashicorp/terraform-provider-clean-google-beta ]]; then
  pushd ~/go/src/github.com/hashicorp/terraform-provider-clean-google-beta
  git clean -fdx
  git reset --hard
  git checkout main
  git pull
  popd
else
  mkdir -p ~/go/src/github.com/hashicorp
  git clone https://github.com/hashicorp/terraform-provider-google-beta ~/go/src/github.com/hashicorp/terraform-provider-clean-google-beta
fi


go mod edit -require=github.com/hashicorp/terraform-provider-clean-google-beta@v0.0.0
go mod edit -replace github.com/hashicorp/terraform-provider-clean-google-beta=$(realpath ~/go/src/github.com/hashicorp/terraform-provider-clean-google-beta)
go run scripts/diff.go --resource $1 --verbose
