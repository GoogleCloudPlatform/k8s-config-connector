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

"""This file contains a definition for a resource class with methods for converting samples of a single DCL resource to KCC format."""

import os

from absl import logging
from dataclasses import dataclass
from typing import List

import sample


@dataclass
class Resource:
  # Represents a resource whose samples are being converted from DCL to
  # KCC.
  service: str
  name: str
  kcc_kind: str
  dcl_dir: str
  kcc_testdata_dir: str
  kcc_samples_dir: str
  samples: List[sample.Sample]

  def load_samples(self) -> None:
    logging.info(f'Looking for DCL samples in:\n{self.dcl_dir}/samples')
    self.samples = [
        sample.Sample(
            service=self.service,
            dcl_dir=self.dcl_dir,
            kcc_testdata_dir=self.kcc_testdata_dir,
            kcc_samples_dir=self.kcc_samples_dir,
            file_name=file_name,
        ) for file_name in os.listdir(f'{self.dcl_dir}/samples') if
        file_name.endswith(f'{self.name}.yaml') and 'minimal' not in file_name
    ]

  def convert_samples(self) -> None:
    for s in self.samples:
      s.convert()

  def write_samples(self) -> None:
    is_only = len(self.samples) == 1
    for s in self.samples:
      s.write(is_only)
