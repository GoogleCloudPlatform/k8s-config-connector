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
from google3.cloud.graphite.mmv2.services.google.access_context_manager import (
    service_perimeter_pb2,
)
from google3.cloud.graphite.mmv2.services.google.access_context_manager import (
    service_perimeter_pb2_grpc,
)

from typing import List


class ServicePerimeter(object):
    def __init__(
        self,
        title: str = None,
        description: str = None,
        create_time: str = None,
        update_time: str = None,
        perimeter_type: str = None,
        status: dict = None,
        policy: str = None,
        name: str = None,
        use_explicit_dry_run_spec: bool = None,
        spec: dict = None,
        service_account_file: str = "",
    ):

        channel.initialize()
        self.title = title
        self.description = description
        self.perimeter_type = perimeter_type
        self.status = status
        self.policy = policy
        self.name = name
        self.use_explicit_dry_run_spec = use_explicit_dry_run_spec
        self.spec = spec
        self.service_account_file = service_account_file

    def apply(self):
        stub = service_perimeter_pb2_grpc.AccesscontextmanagerServicePerimeterServiceStub(
            channel.Channel()
        )
        request = (
            service_perimeter_pb2.ApplyAccesscontextmanagerServicePerimeterRequest()
        )
        if Primitive.to_proto(self.title):
            request.resource.title = Primitive.to_proto(self.title)

        if Primitive.to_proto(self.description):
            request.resource.description = Primitive.to_proto(self.description)

        if ServicePerimeterPerimeterTypeEnum.to_proto(self.perimeter_type):
            request.resource.perimeter_type = ServicePerimeterPerimeterTypeEnum.to_proto(
                self.perimeter_type
            )

        if ServicePerimeterStatus.to_proto(self.status):
            request.resource.status.CopyFrom(
                ServicePerimeterStatus.to_proto(self.status)
            )
        else:
            request.resource.ClearField("status")
        if Primitive.to_proto(self.policy):
            request.resource.policy = Primitive.to_proto(self.policy)

        if Primitive.to_proto(self.name):
            request.resource.name = Primitive.to_proto(self.name)

        if Primitive.to_proto(self.use_explicit_dry_run_spec):
            request.resource.use_explicit_dry_run_spec = Primitive.to_proto(
                self.use_explicit_dry_run_spec
            )

        if ServicePerimeterSpec.to_proto(self.spec):
            request.resource.spec.CopyFrom(ServicePerimeterSpec.to_proto(self.spec))
        else:
            request.resource.ClearField("spec")
        request.service_account_file = self.service_account_file

        response = stub.ApplyAccesscontextmanagerServicePerimeter(request)
        self.title = Primitive.from_proto(response.title)
        self.description = Primitive.from_proto(response.description)
        self.create_time = Primitive.from_proto(response.create_time)
        self.update_time = Primitive.from_proto(response.update_time)
        self.perimeter_type = ServicePerimeterPerimeterTypeEnum.from_proto(
            response.perimeter_type
        )
        self.status = ServicePerimeterStatus.from_proto(response.status)
        self.policy = Primitive.from_proto(response.policy)
        self.name = Primitive.from_proto(response.name)
        self.use_explicit_dry_run_spec = Primitive.from_proto(
            response.use_explicit_dry_run_spec
        )
        self.spec = ServicePerimeterSpec.from_proto(response.spec)

    def delete(self):
        stub = service_perimeter_pb2_grpc.AccesscontextmanagerServicePerimeterServiceStub(
            channel.Channel()
        )
        request = (
            service_perimeter_pb2.DeleteAccesscontextmanagerServicePerimeterRequest()
        )
        request.service_account_file = self.service_account_file
        if Primitive.to_proto(self.title):
            request.resource.title = Primitive.to_proto(self.title)

        if Primitive.to_proto(self.description):
            request.resource.description = Primitive.to_proto(self.description)

        if ServicePerimeterPerimeterTypeEnum.to_proto(self.perimeter_type):
            request.resource.perimeter_type = ServicePerimeterPerimeterTypeEnum.to_proto(
                self.perimeter_type
            )

        if ServicePerimeterStatus.to_proto(self.status):
            request.resource.status.CopyFrom(
                ServicePerimeterStatus.to_proto(self.status)
            )
        else:
            request.resource.ClearField("status")
        if Primitive.to_proto(self.policy):
            request.resource.policy = Primitive.to_proto(self.policy)

        if Primitive.to_proto(self.name):
            request.resource.name = Primitive.to_proto(self.name)

        if Primitive.to_proto(self.use_explicit_dry_run_spec):
            request.resource.use_explicit_dry_run_spec = Primitive.to_proto(
                self.use_explicit_dry_run_spec
            )

        if ServicePerimeterSpec.to_proto(self.spec):
            request.resource.spec.CopyFrom(ServicePerimeterSpec.to_proto(self.spec))
        else:
            request.resource.ClearField("spec")
        response = stub.DeleteAccesscontextmanagerServicePerimeter(request)

    @classmethod
    def list(self, policy, service_account_file=""):
        stub = service_perimeter_pb2_grpc.AccesscontextmanagerServicePerimeterServiceStub(
            channel.Channel()
        )
        request = (
            service_perimeter_pb2.ListAccesscontextmanagerServicePerimeterRequest()
        )
        request.service_account_file = service_account_file
        request.Policy = policy

        return stub.ListAccesscontextmanagerServicePerimeter(request).items

    def to_proto(self):
        resource = service_perimeter_pb2.AccesscontextmanagerServicePerimeter()
        if Primitive.to_proto(self.title):
            resource.title = Primitive.to_proto(self.title)
        if Primitive.to_proto(self.description):
            resource.description = Primitive.to_proto(self.description)
        if ServicePerimeterPerimeterTypeEnum.to_proto(self.perimeter_type):
            resource.perimeter_type = ServicePerimeterPerimeterTypeEnum.to_proto(
                self.perimeter_type
            )
        if ServicePerimeterStatus.to_proto(self.status):
            resource.status.CopyFrom(ServicePerimeterStatus.to_proto(self.status))
        else:
            resource.ClearField("status")
        if Primitive.to_proto(self.policy):
            resource.policy = Primitive.to_proto(self.policy)
        if Primitive.to_proto(self.name):
            resource.name = Primitive.to_proto(self.name)
        if Primitive.to_proto(self.use_explicit_dry_run_spec):
            resource.use_explicit_dry_run_spec = Primitive.to_proto(
                self.use_explicit_dry_run_spec
            )
        if ServicePerimeterSpec.to_proto(self.spec):
            resource.spec.CopyFrom(ServicePerimeterSpec.to_proto(self.spec))
        else:
            resource.ClearField("spec")
        return resource


class ServicePerimeterStatus(object):
    def __init__(
        self,
        resources: list = None,
        access_levels: list = None,
        restricted_services: list = None,
        vpc_accessible_services: dict = None,
    ):
        self.resources = resources
        self.access_levels = access_levels
        self.restricted_services = restricted_services
        self.vpc_accessible_services = vpc_accessible_services

    @classmethod
    def to_proto(self, resource):
        if not resource:
            return None

        res = service_perimeter_pb2.AccesscontextmanagerServicePerimeterStatus()
        if Primitive.to_proto(resource.resources):
            res.resources.extend(Primitive.to_proto(resource.resources))
        if Primitive.to_proto(resource.access_levels):
            res.access_levels.extend(Primitive.to_proto(resource.access_levels))
        if Primitive.to_proto(resource.restricted_services):
            res.restricted_services.extend(
                Primitive.to_proto(resource.restricted_services)
            )
        if ServicePerimeterStatusVPCAccessibleServices.to_proto(
            resource.vpc_accessible_services
        ):
            res.vpc_accessible_services.CopyFrom(
                ServicePerimeterStatusVPCAccessibleServices.to_proto(
                    resource.vpc_accessible_services
                )
            )
        else:
            res.ClearField("vpc_accessible_services")
        return res

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return None

        return ServicePerimeterStatus(
            resources=Primitive.from_proto(resource.resources),
            access_levels=Primitive.from_proto(resource.access_levels),
            restricted_services=Primitive.from_proto(resource.restricted_services),
            vpc_accessible_services=ServicePerimeterStatusVPCAccessibleServices.from_proto(
                resource.vpc_accessible_services
            ),
        )


class ServicePerimeterStatusArray(object):
    @classmethod
    def to_proto(self, resources):
        if not resources:
            return resources
        return [ServicePerimeterStatus.to_proto(i) for i in resources]

    @classmethod
    def from_proto(self, resources):
        return [ServicePerimeterStatus.from_proto(i) for i in resources]


class ServicePerimeterStatusVPCAccessibleServices(object):
    def __init__(self, enable_restriction: bool = None, allowed_services: list = None):
        self.enable_restriction = enable_restriction
        self.allowed_services = allowed_services

    @classmethod
    def to_proto(self, resource):
        if not resource:
            return None

        res = (
            service_perimeter_pb2.AccesscontextmanagerServicePerimeterStatusVPCAccessibleServices()
        )
        if Primitive.to_proto(resource.enable_restriction):
            res.enable_restriction = Primitive.to_proto(resource.enable_restriction)
        if Primitive.to_proto(resource.allowed_services):
            res.allowed_services.extend(Primitive.to_proto(resource.allowed_services))
        return res

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return None

        return ServicePerimeterStatusVPCAccessibleServices(
            enable_restriction=Primitive.from_proto(resource.enable_restriction),
            allowed_services=Primitive.from_proto(resource.allowed_services),
        )


class ServicePerimeterStatusVPCAccessibleServicesArray(object):
    @classmethod
    def to_proto(self, resources):
        if not resources:
            return resources
        return [
            ServicePerimeterStatusVPCAccessibleServices.to_proto(i) for i in resources
        ]

    @classmethod
    def from_proto(self, resources):
        return [
            ServicePerimeterStatusVPCAccessibleServices.from_proto(i) for i in resources
        ]


class ServicePerimeterSpec(object):
    def __init__(
        self,
        resources: list = None,
        access_levels: list = None,
        restricted_services: list = None,
        vpc_accessible_services: dict = None,
    ):
        self.resources = resources
        self.access_levels = access_levels
        self.restricted_services = restricted_services
        self.vpc_accessible_services = vpc_accessible_services

    @classmethod
    def to_proto(self, resource):
        if not resource:
            return None

        res = service_perimeter_pb2.AccesscontextmanagerServicePerimeterSpec()
        if Primitive.to_proto(resource.resources):
            res.resources.extend(Primitive.to_proto(resource.resources))
        if Primitive.to_proto(resource.access_levels):
            res.access_levels.extend(Primitive.to_proto(resource.access_levels))
        if Primitive.to_proto(resource.restricted_services):
            res.restricted_services.extend(
                Primitive.to_proto(resource.restricted_services)
            )
        if ServicePerimeterSpecVPCAccessibleServices.to_proto(
            resource.vpc_accessible_services
        ):
            res.vpc_accessible_services.CopyFrom(
                ServicePerimeterSpecVPCAccessibleServices.to_proto(
                    resource.vpc_accessible_services
                )
            )
        else:
            res.ClearField("vpc_accessible_services")
        return res

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return None

        return ServicePerimeterSpec(
            resources=Primitive.from_proto(resource.resources),
            access_levels=Primitive.from_proto(resource.access_levels),
            restricted_services=Primitive.from_proto(resource.restricted_services),
            vpc_accessible_services=ServicePerimeterSpecVPCAccessibleServices.from_proto(
                resource.vpc_accessible_services
            ),
        )


class ServicePerimeterSpecArray(object):
    @classmethod
    def to_proto(self, resources):
        if not resources:
            return resources
        return [ServicePerimeterSpec.to_proto(i) for i in resources]

    @classmethod
    def from_proto(self, resources):
        return [ServicePerimeterSpec.from_proto(i) for i in resources]


class ServicePerimeterSpecVPCAccessibleServices(object):
    def __init__(self, enable_restriction: bool = None, allowed_services: list = None):
        self.enable_restriction = enable_restriction
        self.allowed_services = allowed_services

    @classmethod
    def to_proto(self, resource):
        if not resource:
            return None

        res = (
            service_perimeter_pb2.AccesscontextmanagerServicePerimeterSpecVPCAccessibleServices()
        )
        if Primitive.to_proto(resource.enable_restriction):
            res.enable_restriction = Primitive.to_proto(resource.enable_restriction)
        if Primitive.to_proto(resource.allowed_services):
            res.allowed_services.extend(Primitive.to_proto(resource.allowed_services))
        return res

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return None

        return ServicePerimeterSpecVPCAccessibleServices(
            enable_restriction=Primitive.from_proto(resource.enable_restriction),
            allowed_services=Primitive.from_proto(resource.allowed_services),
        )


class ServicePerimeterSpecVPCAccessibleServicesArray(object):
    @classmethod
    def to_proto(self, resources):
        if not resources:
            return resources
        return [
            ServicePerimeterSpecVPCAccessibleServices.to_proto(i) for i in resources
        ]

    @classmethod
    def from_proto(self, resources):
        return [
            ServicePerimeterSpecVPCAccessibleServices.from_proto(i) for i in resources
        ]


class ServicePerimeterPerimeterTypeEnum(object):
    @classmethod
    def to_proto(self, resource):
        if not resource:
            return resource
        return service_perimeter_pb2.AccesscontextmanagerServicePerimeterPerimeterTypeEnum.Value(
            "AccesscontextmanagerServicePerimeterPerimeterTypeEnum%s" % resource
        )

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return resource
        return service_perimeter_pb2.AccesscontextmanagerServicePerimeterPerimeterTypeEnum.Name(
            resource
        )[
            len("AccesscontextmanagerServicePerimeterPerimeterTypeEnum") :
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
