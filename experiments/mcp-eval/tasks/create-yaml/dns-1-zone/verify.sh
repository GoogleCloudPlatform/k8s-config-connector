#!/bin/bash
set -o errexit
set -o nounset
set -o pipefail

TOP_LEVEL=$(git rev-parse --show-toplevel)
FILE_PATH="${TOP_LEVEL}/.build/tasks/dns-1-zone/zone.yaml"

# Check if the file exists
if [[ ! -f "${FILE_PATH}" ]]; then
  echo "File not found: ${FILE_PATH}"
  exit 1
fi

# Check if the file is a YAML file with kind "DNSManagedZone"
if ! grep -q "kind: DNSManagedZone" "${FILE_PATH}"; then
  echo "File does not have kind: DNSManagedZone"
  exit 1
fi

# Check for the correct settings
if ! grep -q "dnsName: my-zone.example.com." "${FILE_PATH}"; then
  echo "dnsName is not set to my-zone.example.com."
  exit 1
fi

if ! grep -q "description: My DNS zone" "${FILE_PATH}"; then
  echo "description is not set to My DNS zone"
  exit 1
fi

echo "Validation successful!"
exit 0
