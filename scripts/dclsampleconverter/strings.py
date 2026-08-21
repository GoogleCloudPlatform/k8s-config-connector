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

"""This file provides utility functions for working with strings."""

import re

from functools import cache

SERVICE_REPLACEMENTS = {
    'bigquery': 'BigQuery',
    'billingbudgets': 'BillingBudgets',
    'cloudbuild': 'CloudBuild',
    'cloudfunctions': 'CloudFunctions',
    'cloudresourcemanager':
        '',  #TODO(kcc-eng): Handle ResourceManagerLien, etc.
    'dependencyservice': 'DependencyService',
    'dlp': 'DLP',
    'dns': 'DNS',
    'gcp': 'GCP',
    'http': 'HTTP',
    'https': 'HTTPS',
    'iam': 'IAM',
    'identitytoolkit': 'IdentityPlatform',
    'kms': 'KMS',
    'osconfig': 'OSConfig',
    'ospolicyassignment': 'OSPolicyAssignment',
    'pubsub': 'PubSub',
    'sql': 'SQL',
    'tcp': 'TCP',
    'testservice': 'TestService',
    'url': 'URL',
    'vpcaccess': 'VPCAccess',
    'vpn': 'VPN',
}  # Replacement map for KCC Kinds. Keys are case-insensitive.


class colors:
  HEADER = '\033[95m'
  OKBLUE = '\033[94m'
  OKCYAN = '\033[96m'
  OKGREEN = '\033[92m'
  WARNING = '\033[93m'
  FAIL = '\033[91m'
  ENDC = '\033[0m'
  BOLD = '\033[1m'
  UNDERLINE = '\033[4m'


class patterns:
  SNAKE_FIRST_LETTERS = re.compile('(^|_)([a-z])')


def snake_to_title(string):
  return replace_map(
      SERVICE_REPLACEMENTS,
      re.sub(patterns.SNAKE_FIRST_LETTERS, lambda m: m.group(2).upper(),
             string), True)


@cache
def replace_map_regex(keys, ignore_case):
  """Return a regex which matches any of the keys in the given set."""
  if ignore_case:
    return re.compile(r'(' + keys + r')', re.IGNORECASE)
  return re.compile(r'(' + keys + r')')


def replace_map(replacements, string, ignore_case=False):
  """Return the given string with all instances of keys in the given map replaced with their values."""
  if not replacements:
    return string
  return re.sub(
      replace_map_regex(
          r'|'.join([re.escape(key) for key in replacements.keys()]),
          ignore_case),
      lambda match: replacements[match.group().lower()
                                 if ignore_case else match.group()], string)
