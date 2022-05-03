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

"""This file contains tests for the dcl sample converter."""

import os

from absl import logging
from absl.testing import absltest

import resource
import strings


class Test(absltest.TestCase):

  def test_conversion(self):
    print(__file__)
    base_dir = f'{os.path.dirname(os.path.abspath(__file__))}/testdata'
    for test_case in os.listdir(base_dir):
      kcc_testdata_dir = f'{base_dir}/{test_case}/kcc/testdata'
      kcc_samples_dir = f'{base_dir}/{test_case}/kcc/samples'

      r = resource.Resource(
          service='testservice',
          name='test_resource',
          kcc_kind=strings.snake_to_title('testservice_test_resource'),
          dcl_dir=f'{base_dir}/{test_case}/dcl',
          kcc_testdata_dir=kcc_testdata_dir,
          kcc_samples_dir=kcc_samples_dir,
          samples=[],
      )
      r.load_samples()
      r.convert_samples()
      r.write_samples()

      self.compare_dirs(kcc_samples_dir, f'{kcc_samples_dir}/canonical')
      self.compare_dirs(kcc_testdata_dir, f'{kcc_testdata_dir}/canonical')

  def compare_dirs(self, created_path: str, canonical_path: str):
    for file_name in os.listdir(created_path):
      if file_name == 'canonical':
        continue
      created_file_path = f'{created_path}/{file_name}'
      canonical_file_path = f'{canonical_path}/{file_name}'
      if os.path.isdir(created_file_path):
        self.compare_dirs(created_file_path, canonical_file_path)
        continue
      with open(canonical_file_path) as canonical_file, open(
          created_file_path) as created_file:
        logging.info(
            f'Comparing files:\n{created_file_path}\n{canonical_file_path}')
        self.assertEqual(canonical_file.read(), created_file.read())
      os.remove(created_file_path)


if __name__ == '__main__':
  absltest.main()
