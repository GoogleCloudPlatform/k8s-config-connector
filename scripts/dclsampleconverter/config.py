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
"""This file provides the Config class which manages a single resource config

file in a sample.
"""
from __future__ import annotations

import os
import re
import ruamel.yaml

from absl import logging
from collections.abc import Mapping
from collections.abc import Set
from dataclasses import dataclass
from typing import Any, Callable

import strings

# This regex matches a dcl resource reference. The submatch is the resource
# name.
REFERENCE_REGEXP = re.compile(
    r'{{\s*ref:([a-z0-9_]*\.[a-z_]*\.[a-z_]*(\.[a-z_]*)?):[a-zA-Z0-9_\.\[\]]*\s*}}'
)

# Maps variable names to the comment associated with them in doc samples.
VARIABLE_COMMENTS = {
    '${ORG_ID?}':
        '# Replace "${ORG_ID?}" with the numeric ID for your organization',
    '${PROJECT_ID?}':
        '# Replace "${PROJECT_ID?}" with your project ID',
}

# These suffixes should be removed from reference field names.
# TODO(kcc-eng) Source this constant directly from
# pkg/dcl/extension/extension.go
TRIMMABLE_REFERENCE_SUFFIXES = [
    'Name', 'Id', 'IdOrNum', 'Email', 'Link', 'Reference'
]


def add_dependency(dependencies: Set[str], value: str) -> None:
  """Add the file name of the sample the given field value depends on, if any."""
  # Take only the last DCL reference in the string as this is most likely the
  # referenced resource.
  reference_match = None
  for reference_match in re.finditer(REFERENCE_REGEXP, value):
    pass
  if reference_match:
    dependencies.add(f'samples/{reference_match.group(1)}')


def find_dependencies(base: Mapping[str, Any]) -> Set[str]:
  """Returns a set of config file names that the given config depends on."""
  dependencies = set()
  for key, value in base.items():
    if isinstance(value, dict):
      dependencies |= find_dependencies(value)
    elif isinstance(value, list):
      for element in value:
        if isinstance(element, dict):
          dependencies |= find_dependencies(element)
        elif isinstance(element, str):
          add_dependency(dependencies, element)
    elif isinstance(value, str):
      add_dependency(dependencies, value)
  return dependencies


@dataclass
class Config:
  """Corresponds to a single resource json file in DCL and a yaml file in KCC."""

  def __init__(self, service: str, dcl_file: str):
    self.dcl_file = dcl_file
    # Infer service and resource name from filename.
    file_name = os.path.split(self.dcl_file)[1]
    file_name_parts = file_name.split('.')
    if len(file_name_parts) == 3:
      # This config belongs to the main service of the sample.
      self.service = service
    elif len(file_name_parts) == 4:
      self.service = file_name_parts[-3]
    else:
      logging.warning(
          f'{strings.colors.WARNING}Unable to infer service and resource name from {file_name}.{strings.colors.ENDC}'
      )
      return
    self.resource = file_name_parts[-2]
    self.kcc_kind = strings.snake_to_title(f'{self.service}_{self.resource}')
    logging.info(f'Loading a {self.kcc_kind} from:\n{self.dcl_file}')
    yaml = ruamel.yaml.YAML(typ='safe', pure=True)
    with open(self.dcl_file) as file_object:
      self.dcl_config = yaml.load(file_object)

  def get(self, field) -> Any:
    """Returns the value of the given field in this resource's dcl config."""
    return self.dcl_config.get(field)

  def depends_on(self) -> Set[str]:
    return find_dependencies(self.dcl_config)

  def convert_project(
      self, variables_replaced: Mapping[str, Any]) -> Mapping[str, Any]:
    """Convert the config of a project resource file from DCL to KCC format."""
    metadata = {'name': variables_replaced['name']}
    parent = variables_replaced['parent']
    parent_type, parent_id = parent.split('s/')
    spec = {
        f'{parent_type}Ref': {
            'external': parent_id,
        },
    }
    display_name = variables_replaced.get('displayName')
    if display_name:
      spec['name'] = display_name
    else:
      # No display name, use project ID.
      spec['name'] = variables_replaced['name']
    return {
        'apiVersion': 'resourcemanager.cnrm.cloud.google.com/v1beta1',
        'kind': 'Project',
        'metadata': metadata,
        'spec': spec,
    }

  def convert(self, dependencies: Mapping[str, Config],
              replacements: Mapping[str, str]) -> Mapping[str, Any]:
    """Convert the config of a single resource file from DCL to KCC format."""
    # Replace references before variables because the reference value will
    # become a variable then be converted to a KCC value.
    references_replaced = replace_strings(
        self.dcl_config, lambda s: resolve_reference(s, dependencies))
    variables_replaced = replace_strings(
        references_replaced, lambda s: strings.replace_map(replacements, s))
    if self.kcc_kind == 'Project':
      return self.convert_project(variables_replaced)
    if 'name' not in variables_replaced:
      variables_replaced['name'] = f'{self.kcc_kind.lower()}-${{uniqueId}}'
    # TODO(kcc-eng): Handle cases where 'labels' is not the metadata
    # labels field.
    metadata = {
        key: variables_replaced[key]
        for key in ('name', 'labels')
        if key in variables_replaced
    }
    spec = {
        key: value
        for key, value in variables_replaced.items()
        if key not in ('name', 'labels', 'project')
    }
    project = variables_replaced.get('project')
    if isinstance(project, dict):
      # The resource is in a project other than the base test project.
      spec['projectRef'] = project
    elif isinstance(project, str):
      spec['projectRef'] = {'external': f'projects/{project}'}
    converted = {
        'apiVersion': f'{self.service}.cnrm.cloud.google.com/v1beta1',
        'kind': self.kcc_kind,
    }
    if metadata:
      converted['metadata'] = metadata
    if spec:
      converted['spec'] = spec
    return converted


def replace_strings(base: Mapping[str, Any],
                    replacer: Callable[str, Any]) -> Mapping[str, Any]:
  """Recursively replace string values in the given mapping according to the given function, then return the result.

  Also renames keys when the function returns a dict instead of a str.
  """
  replaced_base = {}
  for key, value in base.items():
    if isinstance(value, dict):
      replaced_base[key] = replace_strings(value, replacer)
    elif isinstance(value, list):
      replaced_elements = []
      for element in value:
        if isinstance(element, dict):
          replaced_elements.append(replace_strings(element, replacer))
        elif isinstance(element, str):
          replaced_elements.append(replacer(element))
        else:
          replaced_elements.append(element)
      replaced_base[key] = replaced_elements
    elif isinstance(value, str):
      resolved = replacer(value)
      if isinstance(resolved, dict):
        # Only a resolved reference will return a dict from the replacer
        # function.
        for reference_suffix in TRIMMABLE_REFERENCE_SUFFIXES:
          key = key.removesuffix(reference_suffix)
        # TODO(kcc-eng): Handle references with multiple referenced types.
        replaced_base[f'{key}Ref'] = resolved
      else:
        replaced_base[key] = resolved
    else:
      replaced_base[key] = value
  return replaced_base


def resolve_reference(value: str, dependencies: Mapping[str, Config]) -> Any:
  """Return a KCC reference to the sample indicated by the given reference value."""
  reference_match = re.search(REFERENCE_REGEXP, value)
  if not reference_match:
    return value
  return {
      'name':
          dependencies.get(f'samples/{reference_match.group(1)}').get('name')
  }


def prepare_sample(config: Mapping[str, Any],
                   replacer: Callable[str, Any]) -> Mapping[str, Any]:
  """Return the given KCC config in the format for a user sample."""
  return add_comments(quote_spec(replace_strings(config, replacer)))


def quote_spec(config: Mapping[str, Any]) -> Mapping[str, Any]:
  """Return the given KCC config with all strings in spec quoted."""
  quoted = {key: value for key, value in config.items() if key != 'spec'}
  spec = config.get('spec')
  if not spec:
    return quoted
  quoted['spec'] = replace_strings(
      spec, lambda s: ruamel.yaml.scalarstring.DoubleQuotedScalarString(s))
  return quoted


def add_comments(base: Mapping[str, Any], depth: int = 0) -> Mapping[str, Any]:
  """Return the given KCC config with comments needed for user samples."""
  commented = ruamel.yaml.comments.CommentedMap()
  for key, value in base.items():
    if isinstance(value, dict):
      commented[key] = add_comments(value, depth + 2)
    else:
      commented[key] = value
      if isinstance(value, str):
        for variable, comment in VARIABLE_COMMENTS.items():
          if variable in value:
            commented.yaml_set_start_comment(comment, depth)
  return commented
