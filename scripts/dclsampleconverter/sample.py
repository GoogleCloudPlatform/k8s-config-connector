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

"""This file provides the Sample class which contains methods for converting

DCL samples to KCC format
"""
from __future__ import annotations

import os
import re
import ruamel.yaml
import string

from absl import logging
from collections import Counter
from collections.abc import Mapping
from collections.abc import Set
from datetime import date
from typing import Any
from typing import List

import config
import strings

APACHE_LICENSE = f"""# Copyright {date.today().year} Google LLC
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

"""


def kcc_service_version_and_kind(kcc_config: Mapping[str, Any]):
  """Return the service and kind from a kcc resource config."""
  version_parts = kcc_config['apiVersion'].split('/')
  version_parts[0] = version_parts[0].split('.')[0]
  return (version_parts[0], version_parts[1], kcc_config['kind'].lower())


class Sample:
  """Corresponds to a sample yaml file in DCL and a directory in KCC for a resource."""

  def __init__(self, service: str, dcl_dir: str, kcc_testdata_dir: str,
               kcc_samples_dir: str, file_name: str):
    """Load and convert the configs for a single sample."""
    self.service = service
    self.dcl_dir = dcl_dir
    self.kcc_testdata_dir = kcc_testdata_dir
    self.kcc_samples_dir = kcc_samples_dir
    # Load sample file.
    dcl_sample_file = f'{self.dcl_dir}/samples/{file_name}'
    logging.info(f'Loading sample file:\n{dcl_sample_file}')
    yaml = ruamel.yaml.YAML(typ='safe', pure=True)
    with open(dcl_sample_file) as file_object:
      self.dcl_config = yaml.load(file_object)
    self.name = self.dcl_config.get('name')
    dependencies = self.dcl_config.get('dependencies')
    if dependencies:
      self.dcl_dependencies = {
          dep: config.Config(self.service, f'{self.dcl_dir}/{dep}')
          for dep in dependencies
      }
    else:
      self.dcl_dependencies = {}
    self.dcl_create = config.Config(
        self.service, f'{self.dcl_dir}/{self.dcl_config.get("resource")}')
    updates = self.dcl_config.get('updates')
    if updates:
      if len(updates) > 1:
        logging.warning(
            f'{strings.colors.WARNING}Sample {self.name} has more than one update.{strings.colors.ENDC}'
        )
      self.dcl_update = config.Config(
          self.service, f'{self.dcl_dir}/{updates[0].get("resource")}')
      #TODO: Handle update dependencies which are not already in the create dependencies array.
    else:
      self.dcl_update = None
    self.variables = self.dcl_config.get('variables', [])
    # Replace variables in all fields.
    self.replacements = {
        '{{org_id}}':
            '${TEST_ORG_ID}',  #TODO(kcc-eng): Handle resources that expect long form reference for org id.
        '{{project}}': '${projectId}',
        '{{region}}': 'us-west2',
        '{{zone}}': 'us-west2-a',
    }
    self.add_replacements_for_variables()
    self.kcc_dependencies = []
    self.kcc_create = {}
    self.kcc_update = {}

  def add_replacements_for_variables(self) -> None:
    """Add the proper replacement to this sample's conversion map to convert each DCL variable to KCC format."""
    for variable in self.variables:
      variable_type = variable.get('type')
      if variable_type == 'resource_name':
        variable_name = variable.get('name')
        dcl_var = '{{' + variable_name + '}}'
        kcc_var = variable_name.replace('_', '-')
        self.replacements[dcl_var] = f'{kcc_var}-${{uniqueId}}'

  def convert(self) -> None:
    """Convert all dependency configs to KCC format."""
    converted_dependencies = set()
    while len(converted_dependencies) < len(self.dcl_dependencies):
      # Convert the first resource config that doesn't depend on any resources
      # that haven't been converted.
      for file_name, dependency in self.dcl_dependencies.items():
        if file_name in converted_dependencies:
          continue
        if not (dependency.depends_on() - converted_dependencies):
          converted_dependencies.add(file_name)
          self.kcc_dependencies.append(
              dependency.convert(self.dcl_dependencies, self.replacements))
          break
      else:
        logging.error(
            f'{strings.colors.FAIL}Could not find next dependency to apply.{strings.colors.ENDC}'
        )
    # Convert main resource to KCC format.
    self.kcc_create = self.dcl_create.convert(self.dcl_dependencies,
                                              self.replacements)
    if self.dcl_update:
      self.kcc_update = self.dcl_update.convert(self.dcl_dependencies,
                                                self.replacements)

  def write(self, is_only: bool) -> None:
    """After converting samples write them to files."""
    if is_only:
      # No subdirectories should be used if there is only one test case.
      kcc_testdata_path = self.kcc_testdata_dir
      kcc_samples_path = self.kcc_samples_dir
    else:
      kcc_name = self.dcl_config.get('name').replace('_', '-')
      kcc_short_name = kcc_name.replace('-', '')
      kcc_testdata_path = f'{self.kcc_testdata_dir}/{kcc_short_name}'
      kcc_samples_path = f'{self.kcc_samples_dir}/{kcc_name}'
    # Ensure that testdata and samples paths exist.
    for path in (kcc_testdata_path, kcc_samples_path):
      cur_path = '/' if path[0] == '/' else ''
      for step in path.split(os.sep):
        if not step:
          continue
        cur_path = os.path.join(cur_path, step)
        if not os.path.isdir(cur_path):
          os.mkdir(cur_path)
    self.write_testdata(kcc_testdata_path)
    self.write_samples(kcc_samples_path,
                       '' if is_only else f'-{kcc_short_name}')

  def total_kind_counts(self) -> Counter:
    """Return a counter of how many samples have each kcc kind."""
    counts = Counter()
    for kcc_dependency in self.kcc_dependencies:
      counts[kcc_dependency['kind'].lower()] += 1
    return counts

  def rename_replacements(
      self, renamer: Callable[[str, int], str]) -> Mapping[str, str]:
    """Return a map of current config names to new names using the given renaming

    function which takes the kind of the current sample and the index of that
    sample within its kind if there is more than one.
    """
    create_name = self.kcc_create.get('metadata', {}).get('name')
    if create_name:
      replacements = {create_name: renamer(self.kcc_create['kind'].lower(), 0)}
    else:
      replacements = {}
    total_kind_counts = self.total_kind_counts()
    kind_indices = Counter()
    for kcc_dependency in self.kcc_dependencies:
      service, _, kind = kcc_service_version_and_kind(kcc_dependency)
      kcc_dependency_name = kcc_dependency.get('metadata', {}).get('name')
      if not kcc_dependency_name:
        continue
      if total_kind_counts[kind] > 1:
        kind_indices[kind] += 1
        replacements[kcc_dependency_name] = renamer(kind, kind_indices[kind])
      else:
        replacements[kcc_dependency_name] = renamer(kind, 0)
    return replacements

  def write_testdata(self, kcc_testdata_path: str):
    """Write testdata samples."""
    # Count total number of dependencies of each kind to determine which ones
    # need a number suffix.
    replacements = self.rename_replacements(
        lambda k, i: f'{k}-{i}-${{uniqueId}}' if i else f'{k}-${{uniqueId}}')
    replacer = lambda s: strings.replace_map(replacements, s)
    yaml = ruamel.yaml.YAML()
    yaml.indent(mapping=2)
    kcc_dependencies_file = f'{kcc_testdata_path}/dependencies.yaml'
    if self.kcc_dependencies:
      logging.info(f'Writing dependencies to:\n{kcc_dependencies_file}')
      with open(kcc_dependencies_file, 'w') as file_object:
        yaml.dump_all([
            config.quote_spec(config.replace_strings(dep, replacer))
            for dep in self.kcc_dependencies
        ], file_object)
    kcc_create_file = f'{kcc_testdata_path}/create.yaml'
    logging.info(f'Writing create sample to:\n{kcc_create_file}')
    with open(kcc_create_file, 'w') as file_object:
      yaml.dump(
          config.quote_spec(config.replace_strings(self.kcc_create, replacer)),
          file_object)
    if self.kcc_update:
      kcc_update_file = f'{kcc_testdata_path}/update.yaml'
      logging.info(f'Writing update sample to:\n{kcc_update_file}')
      with open(kcc_update_file, 'w') as file_object:
        yaml.dump(
            config.quote_spec(
                config.replace_strings(self.kcc_update, replacer)), file_object)

  def write_samples(self, kcc_samples_path: str, name_suffix: str):
    """Write config samples."""
    _, kcc_create_version, kcc_create_kind = kcc_service_version_and_kind(
        self.kcc_create)
    replacements = self.rename_replacements(
        lambda k, i: f'{kcc_create_kind}-dep{i}{name_suffix}'
        if i else f'{kcc_create_kind}-dep{name_suffix}')
    create_name = self.kcc_create.get('metadata', {}).get('name')
    if create_name:
      replacements[create_name] = f'{kcc_create_kind}-sample{name_suffix}'
    replacements['${TEST_ORG_ID}'] = '${ORG_ID?}'
    replacements['${projectId}'] = '${PROJECT_ID?}'
    replacer = lambda s: strings.replace_map(replacements, s)
    # Map file names to configs so that all configs of a given kind go in the
    # same file.
    file_names_to_kcc_configs = {}
    for kcc_dependency in self.kcc_dependencies:
      service, version, kind = kcc_service_version_and_kind(kcc_dependency)
      file_name = f'{kcc_samples_path}/{service}_{version}_{kind}.yaml'
      existing_configs = file_names_to_kcc_configs.get(file_name)
      if existing_configs:
        # There is already at least one dependency of this kind.
        if isinstance(existing_configs, dict):
          # There is already exactly one dependency of this kind.
          existing_configs = [existing_configs]
          file_names_to_kcc_configs[file_name] = existing_configs
        existing_configs.append(config.prepare_sample(kcc_dependency, replacer))
      else:
        file_names_to_kcc_configs[file_name] = config.prepare_sample(
            kcc_dependency, replacer)
    file_names_to_kcc_configs[
        f'{kcc_samples_path}/{self.service}_{kcc_create_version}_{kcc_create_kind}.yaml'] = config.prepare_sample(
            self.kcc_create, replacer)
    for file_name, kcc_config in file_names_to_kcc_configs.items():
      logging.info(f'Writing sample to:\n{file_name}')
      with open(file_name, 'w') as file_object:
        file_object.write(APACHE_LICENSE)
        yaml = ruamel.yaml.YAML()
        yaml.indent(mapping=2)
        if isinstance(kcc_config, list):
          yaml.dump_all(kcc_config, file_object)
        else:
          yaml.dump(kcc_config, file_object)
