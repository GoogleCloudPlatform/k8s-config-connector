# Copyright 2020 Google LLC. All Rights Reserved.
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
from google3.cloud.graphite.mmv2.services.google.sql import ssl_cert_pb2
from google3.cloud.graphite.mmv2.services.google.sql import ssl_cert_pb2_grpc

from typing import List


class SslCert(object):
    def __init__(
        self,
        cert_serial_number: str = None,
        cert: str = None,
        create_time: str = None,
        common_name: str = None,
        expiration_time: str = None,
        name: str = None,
        instance: str = None,
        project: str = None,
        service_account_file: str = "",
    ):

        channel.initialize()
        self.common_name = common_name
        self.instance = instance
        self.project = project
        self.service_account_file = service_account_file

    def apply(self):
        stub = ssl_cert_pb2_grpc.SqlSslCertServiceStub(channel.Channel())
        request = ssl_cert_pb2.ApplySqlSslCertRequest()
        if Primitive.to_proto(self.common_name):
            request.resource.common_name = Primitive.to_proto(self.common_name)

        if Primitive.to_proto(self.instance):
            request.resource.instance = Primitive.to_proto(self.instance)

        if Primitive.to_proto(self.project):
            request.resource.project = Primitive.to_proto(self.project)

        request.service_account_file = self.service_account_file

        response = stub.ApplySqlSslCert(request)
        self.cert_serial_number = Primitive.from_proto(response.cert_serial_number)
        self.cert = Primitive.from_proto(response.cert)
        self.create_time = Primitive.from_proto(response.create_time)
        self.common_name = Primitive.from_proto(response.common_name)
        self.expiration_time = Primitive.from_proto(response.expiration_time)
        self.name = Primitive.from_proto(response.name)
        self.instance = Primitive.from_proto(response.instance)
        self.project = Primitive.from_proto(response.project)

    @classmethod
    def delete(self, project, instance, name, service_account_file=""):
        stub = ssl_cert_pb2_grpc.SqlSslCertServiceStub(channel.Channel())
        request = ssl_cert_pb2.DeleteSqlSslCertRequest()
        request.service_account_file = service_account_file
        request.Project = project

        request.Instance = instance

        request.Name = name

        response = stub.DeleteSqlSslCert(request)

    @classmethod
    def list(self, project, instance, service_account_file=""):
        stub = ssl_cert_pb2_grpc.SqlSslCertServiceStub(channel.Channel())
        request = ssl_cert_pb2.ListSqlSslCertRequest()
        request.service_account_file = service_account_file
        request.Project = project

        request.Instance = instance

        return stub.ListSqlSslCert(request).items

    @classmethod
    def from_any(self, any_proto):
        # Marshal any proto to regular proto.
        res_proto = ssl_cert_pb2.SqlSslCert()
        any_proto.Unpack(res_proto)

        res = SslCert()
        res.cert_serial_number = Primitive.from_proto(res_proto.cert_serial_number)
        res.cert = Primitive.from_proto(res_proto.cert)
        res.create_time = Primitive.from_proto(res_proto.create_time)
        res.common_name = Primitive.from_proto(res_proto.common_name)
        res.expiration_time = Primitive.from_proto(res_proto.expiration_time)
        res.name = Primitive.from_proto(res_proto.name)
        res.instance = Primitive.from_proto(res_proto.instance)
        res.project = Primitive.from_proto(res_proto.project)
        return res


class Primitive(object):
    @classmethod
    def to_proto(self, s):
        if not s:
            return ""
        return s

    @classmethod
    def from_proto(self, s):
        return s
