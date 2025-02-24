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
from google3.cloud.graphite.mmv2.services.google.compute import health_check_pb2
from google3.cloud.graphite.mmv2.services.google.compute import health_check_pb2_grpc

from typing import List


class HealthCheck(object):
    def __init__(
        self,
        check_interval_sec: int = None,
        description: str = None,
        healthy_threshold: int = None,
        http2_health_check: dict = None,
        http_health_check: dict = None,
        https_health_check: dict = None,
        name: str = None,
        ssl_health_check: dict = None,
        tcp_health_check: dict = None,
        type: str = None,
        unhealthy_threshold: int = None,
        timeout_sec: int = None,
        region: str = None,
        project: str = None,
        self_link: str = None,
        location: str = None,
        service_account_file: str = "",
    ):

        channel.initialize()
        self.check_interval_sec = check_interval_sec
        self.description = description
        self.healthy_threshold = healthy_threshold
        self.http2_health_check = http2_health_check
        self.http_health_check = http_health_check
        self.https_health_check = https_health_check
        self.name = name
        self.ssl_health_check = ssl_health_check
        self.tcp_health_check = tcp_health_check
        self.type = type
        self.unhealthy_threshold = unhealthy_threshold
        self.timeout_sec = timeout_sec
        self.region = region
        self.project = project
        self.location = location
        self.service_account_file = service_account_file

    def apply(self):
        stub = health_check_pb2_grpc.ComputeHealthCheckServiceStub(channel.Channel())
        request = health_check_pb2.ApplyComputeHealthCheckRequest()
        if Primitive.to_proto(self.check_interval_sec):
            request.resource.check_interval_sec = Primitive.to_proto(
                self.check_interval_sec
            )

        if Primitive.to_proto(self.description):
            request.resource.description = Primitive.to_proto(self.description)

        if Primitive.to_proto(self.healthy_threshold):
            request.resource.healthy_threshold = Primitive.to_proto(
                self.healthy_threshold
            )

        if HealthCheckHttp2HealthCheck.to_proto(self.http2_health_check):
            request.resource.http2_health_check.CopyFrom(
                HealthCheckHttp2HealthCheck.to_proto(self.http2_health_check)
            )
        else:
            request.resource.ClearField("http2_health_check")
        if HealthCheckHttpHealthCheck.to_proto(self.http_health_check):
            request.resource.http_health_check.CopyFrom(
                HealthCheckHttpHealthCheck.to_proto(self.http_health_check)
            )
        else:
            request.resource.ClearField("http_health_check")
        if HealthCheckHttpsHealthCheck.to_proto(self.https_health_check):
            request.resource.https_health_check.CopyFrom(
                HealthCheckHttpsHealthCheck.to_proto(self.https_health_check)
            )
        else:
            request.resource.ClearField("https_health_check")
        if Primitive.to_proto(self.name):
            request.resource.name = Primitive.to_proto(self.name)

        if HealthCheckSslHealthCheck.to_proto(self.ssl_health_check):
            request.resource.ssl_health_check.CopyFrom(
                HealthCheckSslHealthCheck.to_proto(self.ssl_health_check)
            )
        else:
            request.resource.ClearField("ssl_health_check")
        if HealthCheckTcpHealthCheck.to_proto(self.tcp_health_check):
            request.resource.tcp_health_check.CopyFrom(
                HealthCheckTcpHealthCheck.to_proto(self.tcp_health_check)
            )
        else:
            request.resource.ClearField("tcp_health_check")
        if HealthCheckTypeEnum.to_proto(self.type):
            request.resource.type = HealthCheckTypeEnum.to_proto(self.type)

        if Primitive.to_proto(self.unhealthy_threshold):
            request.resource.unhealthy_threshold = Primitive.to_proto(
                self.unhealthy_threshold
            )

        if Primitive.to_proto(self.timeout_sec):
            request.resource.timeout_sec = Primitive.to_proto(self.timeout_sec)

        if Primitive.to_proto(self.region):
            request.resource.region = Primitive.to_proto(self.region)

        if Primitive.to_proto(self.project):
            request.resource.project = Primitive.to_proto(self.project)

        if Primitive.to_proto(self.location):
            request.resource.location = Primitive.to_proto(self.location)

        request.service_account_file = self.service_account_file

        response = stub.ApplyComputeHealthCheck(request)
        self.check_interval_sec = Primitive.from_proto(response.check_interval_sec)
        self.description = Primitive.from_proto(response.description)
        self.healthy_threshold = Primitive.from_proto(response.healthy_threshold)
        self.http2_health_check = HealthCheckHttp2HealthCheck.from_proto(
            response.http2_health_check
        )
        self.http_health_check = HealthCheckHttpHealthCheck.from_proto(
            response.http_health_check
        )
        self.https_health_check = HealthCheckHttpsHealthCheck.from_proto(
            response.https_health_check
        )
        self.name = Primitive.from_proto(response.name)
        self.ssl_health_check = HealthCheckSslHealthCheck.from_proto(
            response.ssl_health_check
        )
        self.tcp_health_check = HealthCheckTcpHealthCheck.from_proto(
            response.tcp_health_check
        )
        self.type = HealthCheckTypeEnum.from_proto(response.type)
        self.unhealthy_threshold = Primitive.from_proto(response.unhealthy_threshold)
        self.timeout_sec = Primitive.from_proto(response.timeout_sec)
        self.region = Primitive.from_proto(response.region)
        self.project = Primitive.from_proto(response.project)
        self.self_link = Primitive.from_proto(response.self_link)
        self.location = Primitive.from_proto(response.location)

    def delete(self):
        stub = health_check_pb2_grpc.ComputeHealthCheckServiceStub(channel.Channel())
        request = health_check_pb2.DeleteComputeHealthCheckRequest()
        request.service_account_file = self.service_account_file
        if Primitive.to_proto(self.check_interval_sec):
            request.resource.check_interval_sec = Primitive.to_proto(
                self.check_interval_sec
            )

        if Primitive.to_proto(self.description):
            request.resource.description = Primitive.to_proto(self.description)

        if Primitive.to_proto(self.healthy_threshold):
            request.resource.healthy_threshold = Primitive.to_proto(
                self.healthy_threshold
            )

        if HealthCheckHttp2HealthCheck.to_proto(self.http2_health_check):
            request.resource.http2_health_check.CopyFrom(
                HealthCheckHttp2HealthCheck.to_proto(self.http2_health_check)
            )
        else:
            request.resource.ClearField("http2_health_check")
        if HealthCheckHttpHealthCheck.to_proto(self.http_health_check):
            request.resource.http_health_check.CopyFrom(
                HealthCheckHttpHealthCheck.to_proto(self.http_health_check)
            )
        else:
            request.resource.ClearField("http_health_check")
        if HealthCheckHttpsHealthCheck.to_proto(self.https_health_check):
            request.resource.https_health_check.CopyFrom(
                HealthCheckHttpsHealthCheck.to_proto(self.https_health_check)
            )
        else:
            request.resource.ClearField("https_health_check")
        if Primitive.to_proto(self.name):
            request.resource.name = Primitive.to_proto(self.name)

        if HealthCheckSslHealthCheck.to_proto(self.ssl_health_check):
            request.resource.ssl_health_check.CopyFrom(
                HealthCheckSslHealthCheck.to_proto(self.ssl_health_check)
            )
        else:
            request.resource.ClearField("ssl_health_check")
        if HealthCheckTcpHealthCheck.to_proto(self.tcp_health_check):
            request.resource.tcp_health_check.CopyFrom(
                HealthCheckTcpHealthCheck.to_proto(self.tcp_health_check)
            )
        else:
            request.resource.ClearField("tcp_health_check")
        if HealthCheckTypeEnum.to_proto(self.type):
            request.resource.type = HealthCheckTypeEnum.to_proto(self.type)

        if Primitive.to_proto(self.unhealthy_threshold):
            request.resource.unhealthy_threshold = Primitive.to_proto(
                self.unhealthy_threshold
            )

        if Primitive.to_proto(self.timeout_sec):
            request.resource.timeout_sec = Primitive.to_proto(self.timeout_sec)

        if Primitive.to_proto(self.region):
            request.resource.region = Primitive.to_proto(self.region)

        if Primitive.to_proto(self.project):
            request.resource.project = Primitive.to_proto(self.project)

        if Primitive.to_proto(self.location):
            request.resource.location = Primitive.to_proto(self.location)

        response = stub.DeleteComputeHealthCheck(request)

    @classmethod
    def list(self, project, location, service_account_file=""):
        stub = health_check_pb2_grpc.ComputeHealthCheckServiceStub(channel.Channel())
        request = health_check_pb2.ListComputeHealthCheckRequest()
        request.service_account_file = service_account_file
        request.Project = project

        request.Location = location

        return stub.ListComputeHealthCheck(request).items

    def to_proto(self):
        resource = health_check_pb2.ComputeHealthCheck()
        if Primitive.to_proto(self.check_interval_sec):
            resource.check_interval_sec = Primitive.to_proto(self.check_interval_sec)
        if Primitive.to_proto(self.description):
            resource.description = Primitive.to_proto(self.description)
        if Primitive.to_proto(self.healthy_threshold):
            resource.healthy_threshold = Primitive.to_proto(self.healthy_threshold)
        if HealthCheckHttp2HealthCheck.to_proto(self.http2_health_check):
            resource.http2_health_check.CopyFrom(
                HealthCheckHttp2HealthCheck.to_proto(self.http2_health_check)
            )
        else:
            resource.ClearField("http2_health_check")
        if HealthCheckHttpHealthCheck.to_proto(self.http_health_check):
            resource.http_health_check.CopyFrom(
                HealthCheckHttpHealthCheck.to_proto(self.http_health_check)
            )
        else:
            resource.ClearField("http_health_check")
        if HealthCheckHttpsHealthCheck.to_proto(self.https_health_check):
            resource.https_health_check.CopyFrom(
                HealthCheckHttpsHealthCheck.to_proto(self.https_health_check)
            )
        else:
            resource.ClearField("https_health_check")
        if Primitive.to_proto(self.name):
            resource.name = Primitive.to_proto(self.name)
        if HealthCheckSslHealthCheck.to_proto(self.ssl_health_check):
            resource.ssl_health_check.CopyFrom(
                HealthCheckSslHealthCheck.to_proto(self.ssl_health_check)
            )
        else:
            resource.ClearField("ssl_health_check")
        if HealthCheckTcpHealthCheck.to_proto(self.tcp_health_check):
            resource.tcp_health_check.CopyFrom(
                HealthCheckTcpHealthCheck.to_proto(self.tcp_health_check)
            )
        else:
            resource.ClearField("tcp_health_check")
        if HealthCheckTypeEnum.to_proto(self.type):
            resource.type = HealthCheckTypeEnum.to_proto(self.type)
        if Primitive.to_proto(self.unhealthy_threshold):
            resource.unhealthy_threshold = Primitive.to_proto(self.unhealthy_threshold)
        if Primitive.to_proto(self.timeout_sec):
            resource.timeout_sec = Primitive.to_proto(self.timeout_sec)
        if Primitive.to_proto(self.region):
            resource.region = Primitive.to_proto(self.region)
        if Primitive.to_proto(self.project):
            resource.project = Primitive.to_proto(self.project)
        if Primitive.to_proto(self.location):
            resource.location = Primitive.to_proto(self.location)
        return resource


class HealthCheckHttp2HealthCheck(object):
    def __init__(
        self,
        port: int = None,
        port_name: str = None,
        port_specification: str = None,
        host: str = None,
        request_path: str = None,
        proxy_header: str = None,
        response: str = None,
    ):
        self.port = port
        self.port_name = port_name
        self.port_specification = port_specification
        self.host = host
        self.request_path = request_path
        self.proxy_header = proxy_header
        self.response = response

    @classmethod
    def to_proto(self, resource):
        if not resource:
            return None

        res = health_check_pb2.ComputeHealthCheckHttp2HealthCheck()
        if Primitive.to_proto(resource.port):
            res.port = Primitive.to_proto(resource.port)
        if Primitive.to_proto(resource.port_name):
            res.port_name = Primitive.to_proto(resource.port_name)
        if HealthCheckHttp2HealthCheckPortSpecificationEnum.to_proto(
            resource.port_specification
        ):
            res.port_specification = HealthCheckHttp2HealthCheckPortSpecificationEnum.to_proto(
                resource.port_specification
            )
        if Primitive.to_proto(resource.host):
            res.host = Primitive.to_proto(resource.host)
        if Primitive.to_proto(resource.request_path):
            res.request_path = Primitive.to_proto(resource.request_path)
        if HealthCheckHttp2HealthCheckProxyHeaderEnum.to_proto(resource.proxy_header):
            res.proxy_header = HealthCheckHttp2HealthCheckProxyHeaderEnum.to_proto(
                resource.proxy_header
            )
        if Primitive.to_proto(resource.response):
            res.response = Primitive.to_proto(resource.response)
        return res

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return None

        return HealthCheckHttp2HealthCheck(
            port=Primitive.from_proto(resource.port),
            port_name=Primitive.from_proto(resource.port_name),
            port_specification=HealthCheckHttp2HealthCheckPortSpecificationEnum.from_proto(
                resource.port_specification
            ),
            host=Primitive.from_proto(resource.host),
            request_path=Primitive.from_proto(resource.request_path),
            proxy_header=HealthCheckHttp2HealthCheckProxyHeaderEnum.from_proto(
                resource.proxy_header
            ),
            response=Primitive.from_proto(resource.response),
        )


class HealthCheckHttp2HealthCheckArray(object):
    @classmethod
    def to_proto(self, resources):
        if not resources:
            return resources
        return [HealthCheckHttp2HealthCheck.to_proto(i) for i in resources]

    @classmethod
    def from_proto(self, resources):
        return [HealthCheckHttp2HealthCheck.from_proto(i) for i in resources]


class HealthCheckHttpHealthCheck(object):
    def __init__(
        self,
        port: int = None,
        port_name: str = None,
        port_specification: str = None,
        host: str = None,
        request_path: str = None,
        proxy_header: str = None,
        response: str = None,
    ):
        self.port = port
        self.port_name = port_name
        self.port_specification = port_specification
        self.host = host
        self.request_path = request_path
        self.proxy_header = proxy_header
        self.response = response

    @classmethod
    def to_proto(self, resource):
        if not resource:
            return None

        res = health_check_pb2.ComputeHealthCheckHttpHealthCheck()
        if Primitive.to_proto(resource.port):
            res.port = Primitive.to_proto(resource.port)
        if Primitive.to_proto(resource.port_name):
            res.port_name = Primitive.to_proto(resource.port_name)
        if HealthCheckHttpHealthCheckPortSpecificationEnum.to_proto(
            resource.port_specification
        ):
            res.port_specification = HealthCheckHttpHealthCheckPortSpecificationEnum.to_proto(
                resource.port_specification
            )
        if Primitive.to_proto(resource.host):
            res.host = Primitive.to_proto(resource.host)
        if Primitive.to_proto(resource.request_path):
            res.request_path = Primitive.to_proto(resource.request_path)
        if HealthCheckHttpHealthCheckProxyHeaderEnum.to_proto(resource.proxy_header):
            res.proxy_header = HealthCheckHttpHealthCheckProxyHeaderEnum.to_proto(
                resource.proxy_header
            )
        if Primitive.to_proto(resource.response):
            res.response = Primitive.to_proto(resource.response)
        return res

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return None

        return HealthCheckHttpHealthCheck(
            port=Primitive.from_proto(resource.port),
            port_name=Primitive.from_proto(resource.port_name),
            port_specification=HealthCheckHttpHealthCheckPortSpecificationEnum.from_proto(
                resource.port_specification
            ),
            host=Primitive.from_proto(resource.host),
            request_path=Primitive.from_proto(resource.request_path),
            proxy_header=HealthCheckHttpHealthCheckProxyHeaderEnum.from_proto(
                resource.proxy_header
            ),
            response=Primitive.from_proto(resource.response),
        )


class HealthCheckHttpHealthCheckArray(object):
    @classmethod
    def to_proto(self, resources):
        if not resources:
            return resources
        return [HealthCheckHttpHealthCheck.to_proto(i) for i in resources]

    @classmethod
    def from_proto(self, resources):
        return [HealthCheckHttpHealthCheck.from_proto(i) for i in resources]


class HealthCheckHttpsHealthCheck(object):
    def __init__(
        self,
        port: int = None,
        port_name: str = None,
        port_specification: str = None,
        host: str = None,
        request_path: str = None,
        proxy_header: str = None,
        response: str = None,
    ):
        self.port = port
        self.port_name = port_name
        self.port_specification = port_specification
        self.host = host
        self.request_path = request_path
        self.proxy_header = proxy_header
        self.response = response

    @classmethod
    def to_proto(self, resource):
        if not resource:
            return None

        res = health_check_pb2.ComputeHealthCheckHttpsHealthCheck()
        if Primitive.to_proto(resource.port):
            res.port = Primitive.to_proto(resource.port)
        if Primitive.to_proto(resource.port_name):
            res.port_name = Primitive.to_proto(resource.port_name)
        if HealthCheckHttpsHealthCheckPortSpecificationEnum.to_proto(
            resource.port_specification
        ):
            res.port_specification = HealthCheckHttpsHealthCheckPortSpecificationEnum.to_proto(
                resource.port_specification
            )
        if Primitive.to_proto(resource.host):
            res.host = Primitive.to_proto(resource.host)
        if Primitive.to_proto(resource.request_path):
            res.request_path = Primitive.to_proto(resource.request_path)
        if HealthCheckHttpsHealthCheckProxyHeaderEnum.to_proto(resource.proxy_header):
            res.proxy_header = HealthCheckHttpsHealthCheckProxyHeaderEnum.to_proto(
                resource.proxy_header
            )
        if Primitive.to_proto(resource.response):
            res.response = Primitive.to_proto(resource.response)
        return res

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return None

        return HealthCheckHttpsHealthCheck(
            port=Primitive.from_proto(resource.port),
            port_name=Primitive.from_proto(resource.port_name),
            port_specification=HealthCheckHttpsHealthCheckPortSpecificationEnum.from_proto(
                resource.port_specification
            ),
            host=Primitive.from_proto(resource.host),
            request_path=Primitive.from_proto(resource.request_path),
            proxy_header=HealthCheckHttpsHealthCheckProxyHeaderEnum.from_proto(
                resource.proxy_header
            ),
            response=Primitive.from_proto(resource.response),
        )


class HealthCheckHttpsHealthCheckArray(object):
    @classmethod
    def to_proto(self, resources):
        if not resources:
            return resources
        return [HealthCheckHttpsHealthCheck.to_proto(i) for i in resources]

    @classmethod
    def from_proto(self, resources):
        return [HealthCheckHttpsHealthCheck.from_proto(i) for i in resources]


class HealthCheckSslHealthCheck(object):
    def __init__(
        self,
        port: int = None,
        port_name: str = None,
        port_specification: str = None,
        request: str = None,
        response: str = None,
        proxy_header: str = None,
    ):
        self.port = port
        self.port_name = port_name
        self.port_specification = port_specification
        self.request = request
        self.response = response
        self.proxy_header = proxy_header

    @classmethod
    def to_proto(self, resource):
        if not resource:
            return None

        res = health_check_pb2.ComputeHealthCheckSslHealthCheck()
        if Primitive.to_proto(resource.port):
            res.port = Primitive.to_proto(resource.port)
        if Primitive.to_proto(resource.port_name):
            res.port_name = Primitive.to_proto(resource.port_name)
        if HealthCheckSslHealthCheckPortSpecificationEnum.to_proto(
            resource.port_specification
        ):
            res.port_specification = HealthCheckSslHealthCheckPortSpecificationEnum.to_proto(
                resource.port_specification
            )
        if Primitive.to_proto(resource.request):
            res.request = Primitive.to_proto(resource.request)
        if Primitive.to_proto(resource.response):
            res.response = Primitive.to_proto(resource.response)
        if HealthCheckSslHealthCheckProxyHeaderEnum.to_proto(resource.proxy_header):
            res.proxy_header = HealthCheckSslHealthCheckProxyHeaderEnum.to_proto(
                resource.proxy_header
            )
        return res

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return None

        return HealthCheckSslHealthCheck(
            port=Primitive.from_proto(resource.port),
            port_name=Primitive.from_proto(resource.port_name),
            port_specification=HealthCheckSslHealthCheckPortSpecificationEnum.from_proto(
                resource.port_specification
            ),
            request=Primitive.from_proto(resource.request),
            response=Primitive.from_proto(resource.response),
            proxy_header=HealthCheckSslHealthCheckProxyHeaderEnum.from_proto(
                resource.proxy_header
            ),
        )


class HealthCheckSslHealthCheckArray(object):
    @classmethod
    def to_proto(self, resources):
        if not resources:
            return resources
        return [HealthCheckSslHealthCheck.to_proto(i) for i in resources]

    @classmethod
    def from_proto(self, resources):
        return [HealthCheckSslHealthCheck.from_proto(i) for i in resources]


class HealthCheckTcpHealthCheck(object):
    def __init__(
        self,
        port: int = None,
        port_name: str = None,
        port_specification: str = None,
        request: str = None,
        response: str = None,
        proxy_header: str = None,
    ):
        self.port = port
        self.port_name = port_name
        self.port_specification = port_specification
        self.request = request
        self.response = response
        self.proxy_header = proxy_header

    @classmethod
    def to_proto(self, resource):
        if not resource:
            return None

        res = health_check_pb2.ComputeHealthCheckTcpHealthCheck()
        if Primitive.to_proto(resource.port):
            res.port = Primitive.to_proto(resource.port)
        if Primitive.to_proto(resource.port_name):
            res.port_name = Primitive.to_proto(resource.port_name)
        if HealthCheckTcpHealthCheckPortSpecificationEnum.to_proto(
            resource.port_specification
        ):
            res.port_specification = HealthCheckTcpHealthCheckPortSpecificationEnum.to_proto(
                resource.port_specification
            )
        if Primitive.to_proto(resource.request):
            res.request = Primitive.to_proto(resource.request)
        if Primitive.to_proto(resource.response):
            res.response = Primitive.to_proto(resource.response)
        if HealthCheckTcpHealthCheckProxyHeaderEnum.to_proto(resource.proxy_header):
            res.proxy_header = HealthCheckTcpHealthCheckProxyHeaderEnum.to_proto(
                resource.proxy_header
            )
        return res

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return None

        return HealthCheckTcpHealthCheck(
            port=Primitive.from_proto(resource.port),
            port_name=Primitive.from_proto(resource.port_name),
            port_specification=HealthCheckTcpHealthCheckPortSpecificationEnum.from_proto(
                resource.port_specification
            ),
            request=Primitive.from_proto(resource.request),
            response=Primitive.from_proto(resource.response),
            proxy_header=HealthCheckTcpHealthCheckProxyHeaderEnum.from_proto(
                resource.proxy_header
            ),
        )


class HealthCheckTcpHealthCheckArray(object):
    @classmethod
    def to_proto(self, resources):
        if not resources:
            return resources
        return [HealthCheckTcpHealthCheck.to_proto(i) for i in resources]

    @classmethod
    def from_proto(self, resources):
        return [HealthCheckTcpHealthCheck.from_proto(i) for i in resources]


class HealthCheckHttp2HealthCheckPortSpecificationEnum(object):
    @classmethod
    def to_proto(self, resource):
        if not resource:
            return resource
        return health_check_pb2.ComputeHealthCheckHttp2HealthCheckPortSpecificationEnum.Value(
            "ComputeHealthCheckHttp2HealthCheckPortSpecificationEnum%s" % resource
        )

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return resource
        return health_check_pb2.ComputeHealthCheckHttp2HealthCheckPortSpecificationEnum.Name(
            resource
        )[
            len("ComputeHealthCheckHttp2HealthCheckPortSpecificationEnum") :
        ]


class HealthCheckHttp2HealthCheckProxyHeaderEnum(object):
    @classmethod
    def to_proto(self, resource):
        if not resource:
            return resource
        return health_check_pb2.ComputeHealthCheckHttp2HealthCheckProxyHeaderEnum.Value(
            "ComputeHealthCheckHttp2HealthCheckProxyHeaderEnum%s" % resource
        )

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return resource
        return health_check_pb2.ComputeHealthCheckHttp2HealthCheckProxyHeaderEnum.Name(
            resource
        )[len("ComputeHealthCheckHttp2HealthCheckProxyHeaderEnum") :]


class HealthCheckHttpHealthCheckPortSpecificationEnum(object):
    @classmethod
    def to_proto(self, resource):
        if not resource:
            return resource
        return health_check_pb2.ComputeHealthCheckHttpHealthCheckPortSpecificationEnum.Value(
            "ComputeHealthCheckHttpHealthCheckPortSpecificationEnum%s" % resource
        )

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return resource
        return health_check_pb2.ComputeHealthCheckHttpHealthCheckPortSpecificationEnum.Name(
            resource
        )[
            len("ComputeHealthCheckHttpHealthCheckPortSpecificationEnum") :
        ]


class HealthCheckHttpHealthCheckProxyHeaderEnum(object):
    @classmethod
    def to_proto(self, resource):
        if not resource:
            return resource
        return health_check_pb2.ComputeHealthCheckHttpHealthCheckProxyHeaderEnum.Value(
            "ComputeHealthCheckHttpHealthCheckProxyHeaderEnum%s" % resource
        )

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return resource
        return health_check_pb2.ComputeHealthCheckHttpHealthCheckProxyHeaderEnum.Name(
            resource
        )[len("ComputeHealthCheckHttpHealthCheckProxyHeaderEnum") :]


class HealthCheckHttpsHealthCheckPortSpecificationEnum(object):
    @classmethod
    def to_proto(self, resource):
        if not resource:
            return resource
        return health_check_pb2.ComputeHealthCheckHttpsHealthCheckPortSpecificationEnum.Value(
            "ComputeHealthCheckHttpsHealthCheckPortSpecificationEnum%s" % resource
        )

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return resource
        return health_check_pb2.ComputeHealthCheckHttpsHealthCheckPortSpecificationEnum.Name(
            resource
        )[
            len("ComputeHealthCheckHttpsHealthCheckPortSpecificationEnum") :
        ]


class HealthCheckHttpsHealthCheckProxyHeaderEnum(object):
    @classmethod
    def to_proto(self, resource):
        if not resource:
            return resource
        return health_check_pb2.ComputeHealthCheckHttpsHealthCheckProxyHeaderEnum.Value(
            "ComputeHealthCheckHttpsHealthCheckProxyHeaderEnum%s" % resource
        )

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return resource
        return health_check_pb2.ComputeHealthCheckHttpsHealthCheckProxyHeaderEnum.Name(
            resource
        )[len("ComputeHealthCheckHttpsHealthCheckProxyHeaderEnum") :]


class HealthCheckSslHealthCheckPortSpecificationEnum(object):
    @classmethod
    def to_proto(self, resource):
        if not resource:
            return resource
        return health_check_pb2.ComputeHealthCheckSslHealthCheckPortSpecificationEnum.Value(
            "ComputeHealthCheckSslHealthCheckPortSpecificationEnum%s" % resource
        )

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return resource
        return health_check_pb2.ComputeHealthCheckSslHealthCheckPortSpecificationEnum.Name(
            resource
        )[
            len("ComputeHealthCheckSslHealthCheckPortSpecificationEnum") :
        ]


class HealthCheckSslHealthCheckProxyHeaderEnum(object):
    @classmethod
    def to_proto(self, resource):
        if not resource:
            return resource
        return health_check_pb2.ComputeHealthCheckSslHealthCheckProxyHeaderEnum.Value(
            "ComputeHealthCheckSslHealthCheckProxyHeaderEnum%s" % resource
        )

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return resource
        return health_check_pb2.ComputeHealthCheckSslHealthCheckProxyHeaderEnum.Name(
            resource
        )[len("ComputeHealthCheckSslHealthCheckProxyHeaderEnum") :]


class HealthCheckTcpHealthCheckPortSpecificationEnum(object):
    @classmethod
    def to_proto(self, resource):
        if not resource:
            return resource
        return health_check_pb2.ComputeHealthCheckTcpHealthCheckPortSpecificationEnum.Value(
            "ComputeHealthCheckTcpHealthCheckPortSpecificationEnum%s" % resource
        )

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return resource
        return health_check_pb2.ComputeHealthCheckTcpHealthCheckPortSpecificationEnum.Name(
            resource
        )[
            len("ComputeHealthCheckTcpHealthCheckPortSpecificationEnum") :
        ]


class HealthCheckTcpHealthCheckProxyHeaderEnum(object):
    @classmethod
    def to_proto(self, resource):
        if not resource:
            return resource
        return health_check_pb2.ComputeHealthCheckTcpHealthCheckProxyHeaderEnum.Value(
            "ComputeHealthCheckTcpHealthCheckProxyHeaderEnum%s" % resource
        )

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return resource
        return health_check_pb2.ComputeHealthCheckTcpHealthCheckProxyHeaderEnum.Name(
            resource
        )[len("ComputeHealthCheckTcpHealthCheckProxyHeaderEnum") :]


class HealthCheckTypeEnum(object):
    @classmethod
    def to_proto(self, resource):
        if not resource:
            return resource
        return health_check_pb2.ComputeHealthCheckTypeEnum.Value(
            "ComputeHealthCheckTypeEnum%s" % resource
        )

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return resource
        return health_check_pb2.ComputeHealthCheckTypeEnum.Name(resource)[
            len("ComputeHealthCheckTypeEnum") :
        ]


class Primitive(object):
    @classmethod
    def to_proto(self, s):
        if not s:
            return ""
        return s

    @classmethod
    def from_proto(self, s):
        return s
