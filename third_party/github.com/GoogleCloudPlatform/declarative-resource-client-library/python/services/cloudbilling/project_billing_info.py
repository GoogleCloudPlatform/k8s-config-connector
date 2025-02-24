# Copyright 2021 Google LLC. All Rights Reserved.
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
from connector import channel
from google3.cloud.graphite.mmv2.services.google.cloud_billing import (
    project_billing_info_pb2,
)
from google3.cloud.graphite.mmv2.services.google.cloud_billing import (
    project_billing_info_pb2_grpc,
)

from typing import List


class ProjectBillingInfo(object):
    def __init__(
        self,
        name: str = None,
        billing_account_name: str = None,
        billing_enabled: str = None,
        service_account_file: str = "",
    ):

        channel.initialize()
        self.name = name
        self.billing_account_name = billing_account_name
        self.service_account_file = service_account_file

    def apply(self):
        stub = project_billing_info_pb2_grpc.CloudbillingProjectBillingInfoServiceStub(
            channel.Channel()
        )
        request = project_billing_info_pb2.ApplyCloudbillingProjectBillingInfoRequest()
        if Primitive.to_proto(self.name):
            request.resource.name = Primitive.to_proto(self.name)

        if Primitive.to_proto(self.billing_account_name):
            request.resource.billing_account_name = Primitive.to_proto(
                self.billing_account_name
            )

        request.service_account_file = self.service_account_file

        response = stub.ApplyCloudbillingProjectBillingInfo(request)
        self.name = Primitive.from_proto(response.name)
        self.billing_account_name = Primitive.from_proto(response.billing_account_name)
        self.billing_enabled = Primitive.from_proto(response.billing_enabled)

    def delete(self):
        stub = project_billing_info_pb2_grpc.CloudbillingProjectBillingInfoServiceStub(
            channel.Channel()
        )
        request = project_billing_info_pb2.DeleteCloudbillingProjectBillingInfoRequest()
        request.service_account_file = self.service_account_file
        if Primitive.to_proto(self.name):
            request.resource.name = Primitive.to_proto(self.name)

        if Primitive.to_proto(self.billing_account_name):
            request.resource.billing_account_name = Primitive.to_proto(
                self.billing_account_name
            )

        response = stub.DeleteCloudbillingProjectBillingInfo(request)

    @classmethod
    def list(self, service_account_file=""):
        stub = project_billing_info_pb2_grpc.CloudbillingProjectBillingInfoServiceStub(
            channel.Channel()
        )
        request = project_billing_info_pb2.ListCloudbillingProjectBillingInfoRequest()
        request.service_account_file = service_account_file

        return stub.ListCloudbillingProjectBillingInfo(request).items

    def to_proto(self):
        resource = project_billing_info_pb2.CloudbillingProjectBillingInfo()
        if Primitive.to_proto(self.name):
            resource.name = Primitive.to_proto(self.name)
        if Primitive.to_proto(self.billing_account_name):
            resource.billing_account_name = Primitive.to_proto(
                self.billing_account_name
            )
        return resource


class Primitive(object):
    @classmethod
    def to_proto(self, s):
        if not s:
            return ""
        return s

    @classmethod
    def from_proto(self, s):
        return s
