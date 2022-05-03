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

"""dclsampleconverter converts samples from DCL format into KCC format."""

from absl import app
from absl import flags
from absl import logging
from collections.abc import Sequence

import resource
import strings

FLAGS = flags.FLAGS

flags.DEFINE_string('service', '', 'The name of the service, snake cased')
flags.DEFINE_string('resource', '', 'The name of the resource, snake cased')
flags.DEFINE_string('dcl_path', '', 'The local path of the DCL GoB repo.')
flags.DEFINE_string('kcc_path', '', 'The local path of the KCC GoB repo.')


def main(argv: Sequence[str]) -> None:
  logging.info(
      f'Converting samples for service {FLAGS.service} resource {FLAGS.resource}.'
  )
  kcc_kind = strings.snake_to_title(
      f'{FLAGS.service}_{FLAGS.resource}'
  )  # e.g. compute_service_attachment -> ComputeServiceAttachment
  r = resource.Resource(
      service=FLAGS.service,
      name=FLAGS.resource,
      kcc_kind=kcc_kind,
      dcl_dir=f'{FLAGS.dcl_path}/services/google/{FLAGS.service}',
      kcc_testdata_dir=f'{FLAGS.kcc_path}/pkg/test/resourcefixture/testdata/basic/{FLAGS.service}/v1beta1/{kcc_kind.lower()}',
      kcc_samples_dir=f'{FLAGS.kcc_path}/config/samples/resources/{kcc_kind.lower()}',
      samples=[],
  )
  r.load_samples()
  r.convert_samples()
  r.write_samples()


if __name__ == '__main__':
  app.run(main)
