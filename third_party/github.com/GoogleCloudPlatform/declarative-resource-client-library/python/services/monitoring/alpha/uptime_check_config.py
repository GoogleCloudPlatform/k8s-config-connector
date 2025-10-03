# Copyright 2024 Google LLC. All Rights Reserved.
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
from google3.cloud.graphite.mmv2.services.google.monitoring import (
    uptime_check_config_pb2,
)
from google3.cloud.graphite.mmv2.services.google.monitoring import (
    uptime_check_config_pb2_grpc,
)

from typing import List


class UptimeCheckConfig(object):
    def __init__(
        self,
        name: str = None,
        display_name: str = None,
        monitored_resource: dict = None,
        resource_group: dict = None,
        http_check: dict = None,
        tcp_check: dict = None,
        period: str = None,
        timeout: str = None,
        content_matchers: list = None,
        selected_regions: list = None,
        project: str = None,
        service_account_file: str = "",
    ):
        channel.initialize()
        self.name = name
        self.display_name = display_name
        self.monitored_resource = monitored_resource
        self.resource_group = resource_group
        self.http_check = http_check
        self.tcp_check = tcp_check
        self.period = period
        self.timeout = timeout
        self.content_matchers = content_matchers
        self.selected_regions = selected_regions
        self.project = project
        self.service_account_file = service_account_file

    def apply(self):
        stub = uptime_check_config_pb2_grpc.MonitoringAlphaUptimeCheckConfigServiceStub(
            channel.Channel()
        )
        request = uptime_check_config_pb2.ApplyMonitoringAlphaUptimeCheckConfigRequest()
        if Primitive.to_proto(self.name):
            request.resource.name = Primitive.to_proto(self.name)

        if Primitive.to_proto(self.display_name):
            request.resource.display_name = Primitive.to_proto(self.display_name)

        if UptimeCheckConfigMonitoredResource.to_proto(self.monitored_resource):
            request.resource.monitored_resource.CopyFrom(
                UptimeCheckConfigMonitoredResource.to_proto(self.monitored_resource)
            )
        else:
            request.resource.ClearField("monitored_resource")
        if UptimeCheckConfigResourceGroup.to_proto(self.resource_group):
            request.resource.resource_group.CopyFrom(
                UptimeCheckConfigResourceGroup.to_proto(self.resource_group)
            )
        else:
            request.resource.ClearField("resource_group")
        if UptimeCheckConfigHttpCheck.to_proto(self.http_check):
            request.resource.http_check.CopyFrom(
                UptimeCheckConfigHttpCheck.to_proto(self.http_check)
            )
        else:
            request.resource.ClearField("http_check")
        if UptimeCheckConfigTcpCheck.to_proto(self.tcp_check):
            request.resource.tcp_check.CopyFrom(
                UptimeCheckConfigTcpCheck.to_proto(self.tcp_check)
            )
        else:
            request.resource.ClearField("tcp_check")
        if Primitive.to_proto(self.period):
            request.resource.period = Primitive.to_proto(self.period)

        if Primitive.to_proto(self.timeout):
            request.resource.timeout = Primitive.to_proto(self.timeout)

        if UptimeCheckConfigContentMatchersArray.to_proto(self.content_matchers):
            request.resource.content_matchers.extend(
                UptimeCheckConfigContentMatchersArray.to_proto(self.content_matchers)
            )
        if Primitive.to_proto(self.selected_regions):
            request.resource.selected_regions.extend(
                Primitive.to_proto(self.selected_regions)
            )
        if Primitive.to_proto(self.project):
            request.resource.project = Primitive.to_proto(self.project)

        request.service_account_file = self.service_account_file

        response = stub.ApplyMonitoringAlphaUptimeCheckConfig(request)
        self.name = Primitive.from_proto(response.name)
        self.display_name = Primitive.from_proto(response.display_name)
        self.monitored_resource = UptimeCheckConfigMonitoredResource.from_proto(
            response.monitored_resource
        )
        self.resource_group = UptimeCheckConfigResourceGroup.from_proto(
            response.resource_group
        )
        self.http_check = UptimeCheckConfigHttpCheck.from_proto(response.http_check)
        self.tcp_check = UptimeCheckConfigTcpCheck.from_proto(response.tcp_check)
        self.period = Primitive.from_proto(response.period)
        self.timeout = Primitive.from_proto(response.timeout)
        self.content_matchers = UptimeCheckConfigContentMatchersArray.from_proto(
            response.content_matchers
        )
        self.selected_regions = Primitive.from_proto(response.selected_regions)
        self.project = Primitive.from_proto(response.project)

    def delete(self):
        stub = uptime_check_config_pb2_grpc.MonitoringAlphaUptimeCheckConfigServiceStub(
            channel.Channel()
        )
        request = (
            uptime_check_config_pb2.DeleteMonitoringAlphaUptimeCheckConfigRequest()
        )
        request.service_account_file = self.service_account_file
        if Primitive.to_proto(self.name):
            request.resource.name = Primitive.to_proto(self.name)

        if Primitive.to_proto(self.display_name):
            request.resource.display_name = Primitive.to_proto(self.display_name)

        if UptimeCheckConfigMonitoredResource.to_proto(self.monitored_resource):
            request.resource.monitored_resource.CopyFrom(
                UptimeCheckConfigMonitoredResource.to_proto(self.monitored_resource)
            )
        else:
            request.resource.ClearField("monitored_resource")
        if UptimeCheckConfigResourceGroup.to_proto(self.resource_group):
            request.resource.resource_group.CopyFrom(
                UptimeCheckConfigResourceGroup.to_proto(self.resource_group)
            )
        else:
            request.resource.ClearField("resource_group")
        if UptimeCheckConfigHttpCheck.to_proto(self.http_check):
            request.resource.http_check.CopyFrom(
                UptimeCheckConfigHttpCheck.to_proto(self.http_check)
            )
        else:
            request.resource.ClearField("http_check")
        if UptimeCheckConfigTcpCheck.to_proto(self.tcp_check):
            request.resource.tcp_check.CopyFrom(
                UptimeCheckConfigTcpCheck.to_proto(self.tcp_check)
            )
        else:
            request.resource.ClearField("tcp_check")
        if Primitive.to_proto(self.period):
            request.resource.period = Primitive.to_proto(self.period)

        if Primitive.to_proto(self.timeout):
            request.resource.timeout = Primitive.to_proto(self.timeout)

        if UptimeCheckConfigContentMatchersArray.to_proto(self.content_matchers):
            request.resource.content_matchers.extend(
                UptimeCheckConfigContentMatchersArray.to_proto(self.content_matchers)
            )
        if Primitive.to_proto(self.selected_regions):
            request.resource.selected_regions.extend(
                Primitive.to_proto(self.selected_regions)
            )
        if Primitive.to_proto(self.project):
            request.resource.project = Primitive.to_proto(self.project)

        response = stub.DeleteMonitoringAlphaUptimeCheckConfig(request)

    @classmethod
    def list(self, project, service_account_file=""):
        stub = uptime_check_config_pb2_grpc.MonitoringAlphaUptimeCheckConfigServiceStub(
            channel.Channel()
        )
        request = uptime_check_config_pb2.ListMonitoringAlphaUptimeCheckConfigRequest()
        request.service_account_file = service_account_file
        request.Project = project

        return stub.ListMonitoringAlphaUptimeCheckConfig(request).items

    def to_proto(self):
        resource = uptime_check_config_pb2.MonitoringAlphaUptimeCheckConfig()
        if Primitive.to_proto(self.name):
            resource.name = Primitive.to_proto(self.name)
        if Primitive.to_proto(self.display_name):
            resource.display_name = Primitive.to_proto(self.display_name)
        if UptimeCheckConfigMonitoredResource.to_proto(self.monitored_resource):
            resource.monitored_resource.CopyFrom(
                UptimeCheckConfigMonitoredResource.to_proto(self.monitored_resource)
            )
        else:
            resource.ClearField("monitored_resource")
        if UptimeCheckConfigResourceGroup.to_proto(self.resource_group):
            resource.resource_group.CopyFrom(
                UptimeCheckConfigResourceGroup.to_proto(self.resource_group)
            )
        else:
            resource.ClearField("resource_group")
        if UptimeCheckConfigHttpCheck.to_proto(self.http_check):
            resource.http_check.CopyFrom(
                UptimeCheckConfigHttpCheck.to_proto(self.http_check)
            )
        else:
            resource.ClearField("http_check")
        if UptimeCheckConfigTcpCheck.to_proto(self.tcp_check):
            resource.tcp_check.CopyFrom(
                UptimeCheckConfigTcpCheck.to_proto(self.tcp_check)
            )
        else:
            resource.ClearField("tcp_check")
        if Primitive.to_proto(self.period):
            resource.period = Primitive.to_proto(self.period)
        if Primitive.to_proto(self.timeout):
            resource.timeout = Primitive.to_proto(self.timeout)
        if UptimeCheckConfigContentMatchersArray.to_proto(self.content_matchers):
            resource.content_matchers.extend(
                UptimeCheckConfigContentMatchersArray.to_proto(self.content_matchers)
            )
        if Primitive.to_proto(self.selected_regions):
            resource.selected_regions.extend(Primitive.to_proto(self.selected_regions))
        if Primitive.to_proto(self.project):
            resource.project = Primitive.to_proto(self.project)
        return resource


class UptimeCheckConfigMonitoredResource(object):
    def __init__(self, type: str = None, filter_labels: dict = None):
        self.type = type
        self.filter_labels = filter_labels

    @classmethod
    def to_proto(self, resource):
        if not resource:
            return None

        res = (
            uptime_check_config_pb2.MonitoringAlphaUptimeCheckConfigMonitoredResource()
        )
        if Primitive.to_proto(resource.type):
            res.type = Primitive.to_proto(resource.type)
        if Primitive.to_proto(resource.filter_labels):
            res.filter_labels = Primitive.to_proto(resource.filter_labels)
        return res

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return None

        return UptimeCheckConfigMonitoredResource(
            type=Primitive.from_proto(resource.type),
            filter_labels=Primitive.from_proto(resource.filter_labels),
        )


class UptimeCheckConfigMonitoredResourceArray(object):
    @classmethod
    def to_proto(self, resources):
        if not resources:
            return resources
        return [UptimeCheckConfigMonitoredResource.to_proto(i) for i in resources]

    @classmethod
    def from_proto(self, resources):
        return [UptimeCheckConfigMonitoredResource.from_proto(i) for i in resources]


class UptimeCheckConfigResourceGroup(object):
    def __init__(self, group_id: str = None, resource_type: str = None):
        self.group_id = group_id
        self.resource_type = resource_type

    @classmethod
    def to_proto(self, resource):
        if not resource:
            return None

        res = uptime_check_config_pb2.MonitoringAlphaUptimeCheckConfigResourceGroup()
        if Primitive.to_proto(resource.group_id):
            res.group_id = Primitive.to_proto(resource.group_id)
        if UptimeCheckConfigResourceGroupResourceTypeEnum.to_proto(
            resource.resource_type
        ):
            res.resource_type = UptimeCheckConfigResourceGroupResourceTypeEnum.to_proto(
                resource.resource_type
            )
        return res

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return None

        return UptimeCheckConfigResourceGroup(
            group_id=Primitive.from_proto(resource.group_id),
            resource_type=UptimeCheckConfigResourceGroupResourceTypeEnum.from_proto(
                resource.resource_type
            ),
        )


class UptimeCheckConfigResourceGroupArray(object):
    @classmethod
    def to_proto(self, resources):
        if not resources:
            return resources
        return [UptimeCheckConfigResourceGroup.to_proto(i) for i in resources]

    @classmethod
    def from_proto(self, resources):
        return [UptimeCheckConfigResourceGroup.from_proto(i) for i in resources]


class UptimeCheckConfigHttpCheck(object):
    def __init__(
        self,
        request_method: str = None,
        use_ssl: bool = None,
        path: str = None,
        port: int = None,
        auth_info: dict = None,
        mask_headers: bool = None,
        headers: dict = None,
        content_type: str = None,
        validate_ssl: bool = None,
        body: str = None,
    ):
        self.request_method = request_method
        self.use_ssl = use_ssl
        self.path = path
        self.port = port
        self.auth_info = auth_info
        self.mask_headers = mask_headers
        self.headers = headers
        self.content_type = content_type
        self.validate_ssl = validate_ssl
        self.body = body

    @classmethod
    def to_proto(self, resource):
        if not resource:
            return None

        res = uptime_check_config_pb2.MonitoringAlphaUptimeCheckConfigHttpCheck()
        if UptimeCheckConfigHttpCheckRequestMethodEnum.to_proto(
            resource.request_method
        ):
            res.request_method = UptimeCheckConfigHttpCheckRequestMethodEnum.to_proto(
                resource.request_method
            )
        if Primitive.to_proto(resource.use_ssl):
            res.use_ssl = Primitive.to_proto(resource.use_ssl)
        if Primitive.to_proto(resource.path):
            res.path = Primitive.to_proto(resource.path)
        if Primitive.to_proto(resource.port):
            res.port = Primitive.to_proto(resource.port)
        if UptimeCheckConfigHttpCheckAuthInfo.to_proto(resource.auth_info):
            res.auth_info.CopyFrom(
                UptimeCheckConfigHttpCheckAuthInfo.to_proto(resource.auth_info)
            )
        else:
            res.ClearField("auth_info")
        if Primitive.to_proto(resource.mask_headers):
            res.mask_headers = Primitive.to_proto(resource.mask_headers)
        if Primitive.to_proto(resource.headers):
            res.headers = Primitive.to_proto(resource.headers)
        if UptimeCheckConfigHttpCheckContentTypeEnum.to_proto(resource.content_type):
            res.content_type = UptimeCheckConfigHttpCheckContentTypeEnum.to_proto(
                resource.content_type
            )
        if Primitive.to_proto(resource.validate_ssl):
            res.validate_ssl = Primitive.to_proto(resource.validate_ssl)
        if Primitive.to_proto(resource.body):
            res.body = Primitive.to_proto(resource.body)
        return res

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return None

        return UptimeCheckConfigHttpCheck(
            request_method=UptimeCheckConfigHttpCheckRequestMethodEnum.from_proto(
                resource.request_method
            ),
            use_ssl=Primitive.from_proto(resource.use_ssl),
            path=Primitive.from_proto(resource.path),
            port=Primitive.from_proto(resource.port),
            auth_info=UptimeCheckConfigHttpCheckAuthInfo.from_proto(resource.auth_info),
            mask_headers=Primitive.from_proto(resource.mask_headers),
            headers=Primitive.from_proto(resource.headers),
            content_type=UptimeCheckConfigHttpCheckContentTypeEnum.from_proto(
                resource.content_type
            ),
            validate_ssl=Primitive.from_proto(resource.validate_ssl),
            body=Primitive.from_proto(resource.body),
        )


class UptimeCheckConfigHttpCheckArray(object):
    @classmethod
    def to_proto(self, resources):
        if not resources:
            return resources
        return [UptimeCheckConfigHttpCheck.to_proto(i) for i in resources]

    @classmethod
    def from_proto(self, resources):
        return [UptimeCheckConfigHttpCheck.from_proto(i) for i in resources]


class UptimeCheckConfigHttpCheckAuthInfo(object):
    def __init__(self, username: str = None, password: str = None):
        self.username = username
        self.password = password

    @classmethod
    def to_proto(self, resource):
        if not resource:
            return None

        res = (
            uptime_check_config_pb2.MonitoringAlphaUptimeCheckConfigHttpCheckAuthInfo()
        )
        if Primitive.to_proto(resource.username):
            res.username = Primitive.to_proto(resource.username)
        if Primitive.to_proto(resource.password):
            res.password = Primitive.to_proto(resource.password)
        return res

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return None

        return UptimeCheckConfigHttpCheckAuthInfo(
            username=Primitive.from_proto(resource.username),
            password=Primitive.from_proto(resource.password),
        )


class UptimeCheckConfigHttpCheckAuthInfoArray(object):
    @classmethod
    def to_proto(self, resources):
        if not resources:
            return resources
        return [UptimeCheckConfigHttpCheckAuthInfo.to_proto(i) for i in resources]

    @classmethod
    def from_proto(self, resources):
        return [UptimeCheckConfigHttpCheckAuthInfo.from_proto(i) for i in resources]


class UptimeCheckConfigTcpCheck(object):
    def __init__(self, port: int = None):
        self.port = port

    @classmethod
    def to_proto(self, resource):
        if not resource:
            return None

        res = uptime_check_config_pb2.MonitoringAlphaUptimeCheckConfigTcpCheck()
        if Primitive.to_proto(resource.port):
            res.port = Primitive.to_proto(resource.port)
        return res

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return None

        return UptimeCheckConfigTcpCheck(
            port=Primitive.from_proto(resource.port),
        )


class UptimeCheckConfigTcpCheckArray(object):
    @classmethod
    def to_proto(self, resources):
        if not resources:
            return resources
        return [UptimeCheckConfigTcpCheck.to_proto(i) for i in resources]

    @classmethod
    def from_proto(self, resources):
        return [UptimeCheckConfigTcpCheck.from_proto(i) for i in resources]


class UptimeCheckConfigContentMatchers(object):
    def __init__(self, content: str = None, matcher: str = None):
        self.content = content
        self.matcher = matcher

    @classmethod
    def to_proto(self, resource):
        if not resource:
            return None

        res = uptime_check_config_pb2.MonitoringAlphaUptimeCheckConfigContentMatchers()
        if Primitive.to_proto(resource.content):
            res.content = Primitive.to_proto(resource.content)
        if UptimeCheckConfigContentMatchersMatcherEnum.to_proto(resource.matcher):
            res.matcher = UptimeCheckConfigContentMatchersMatcherEnum.to_proto(
                resource.matcher
            )
        return res

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return None

        return UptimeCheckConfigContentMatchers(
            content=Primitive.from_proto(resource.content),
            matcher=UptimeCheckConfigContentMatchersMatcherEnum.from_proto(
                resource.matcher
            ),
        )


class UptimeCheckConfigContentMatchersArray(object):
    @classmethod
    def to_proto(self, resources):
        if not resources:
            return resources
        return [UptimeCheckConfigContentMatchers.to_proto(i) for i in resources]

    @classmethod
    def from_proto(self, resources):
        return [UptimeCheckConfigContentMatchers.from_proto(i) for i in resources]


class UptimeCheckConfigResourceGroupResourceTypeEnum(object):
    @classmethod
    def to_proto(self, resource):
        if not resource:
            return resource
        return uptime_check_config_pb2.MonitoringAlphaUptimeCheckConfigResourceGroupResourceTypeEnum.Value(
            "MonitoringAlphaUptimeCheckConfigResourceGroupResourceTypeEnum%s" % resource
        )

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return resource
        return uptime_check_config_pb2.MonitoringAlphaUptimeCheckConfigResourceGroupResourceTypeEnum.Name(
            resource
        )[
            len("MonitoringAlphaUptimeCheckConfigResourceGroupResourceTypeEnum") :
        ]


class UptimeCheckConfigHttpCheckRequestMethodEnum(object):
    @classmethod
    def to_proto(self, resource):
        if not resource:
            return resource
        return uptime_check_config_pb2.MonitoringAlphaUptimeCheckConfigHttpCheckRequestMethodEnum.Value(
            "MonitoringAlphaUptimeCheckConfigHttpCheckRequestMethodEnum%s" % resource
        )

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return resource
        return uptime_check_config_pb2.MonitoringAlphaUptimeCheckConfigHttpCheckRequestMethodEnum.Name(
            resource
        )[
            len("MonitoringAlphaUptimeCheckConfigHttpCheckRequestMethodEnum") :
        ]


class UptimeCheckConfigHttpCheckContentTypeEnum(object):
    @classmethod
    def to_proto(self, resource):
        if not resource:
            return resource
        return uptime_check_config_pb2.MonitoringAlphaUptimeCheckConfigHttpCheckContentTypeEnum.Value(
            "MonitoringAlphaUptimeCheckConfigHttpCheckContentTypeEnum%s" % resource
        )

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return resource
        return uptime_check_config_pb2.MonitoringAlphaUptimeCheckConfigHttpCheckContentTypeEnum.Name(
            resource
        )[
            len("MonitoringAlphaUptimeCheckConfigHttpCheckContentTypeEnum") :
        ]


class UptimeCheckConfigContentMatchersMatcherEnum(object):
    @classmethod
    def to_proto(self, resource):
        if not resource:
            return resource
        return uptime_check_config_pb2.MonitoringAlphaUptimeCheckConfigContentMatchersMatcherEnum.Value(
            "MonitoringAlphaUptimeCheckConfigContentMatchersMatcherEnum%s" % resource
        )

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return resource
        return uptime_check_config_pb2.MonitoringAlphaUptimeCheckConfigContentMatchersMatcherEnum.Name(
            resource
        )[
            len("MonitoringAlphaUptimeCheckConfigContentMatchersMatcherEnum") :
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
