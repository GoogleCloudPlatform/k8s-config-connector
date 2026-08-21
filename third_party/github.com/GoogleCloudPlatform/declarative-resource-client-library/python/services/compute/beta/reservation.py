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
from google3.cloud.graphite.mmv2.services.google.compute import reservation_pb2
from google3.cloud.graphite.mmv2.services.google.compute import reservation_pb2_grpc

from typing import List


class Reservation(object):
    def __init__(
        self,
        id: int = None,
        self_link: str = None,
        zone: str = None,
        description: str = None,
        name: str = None,
        specific_reservation: dict = None,
        commitment: str = None,
        specific_reservation_required: bool = None,
        status: str = None,
        project: str = None,
        service_account_file: str = "",
    ):

        channel.initialize()
        self.id = id
        self.zone = zone
        self.description = description
        self.name = name
        self.specific_reservation = specific_reservation
        self.commitment = commitment
        self.specific_reservation_required = specific_reservation_required
        self.project = project
        self.service_account_file = service_account_file

    def apply(self):
        stub = reservation_pb2_grpc.ComputeBetaReservationServiceStub(channel.Channel())
        request = reservation_pb2.ApplyComputeBetaReservationRequest()
        if Primitive.to_proto(self.id):
            request.resource.id = Primitive.to_proto(self.id)

        if Primitive.to_proto(self.zone):
            request.resource.zone = Primitive.to_proto(self.zone)

        if Primitive.to_proto(self.description):
            request.resource.description = Primitive.to_proto(self.description)

        if Primitive.to_proto(self.name):
            request.resource.name = Primitive.to_proto(self.name)

        if ReservationSpecificReservation.to_proto(self.specific_reservation):
            request.resource.specific_reservation.CopyFrom(
                ReservationSpecificReservation.to_proto(self.specific_reservation)
            )
        else:
            request.resource.ClearField("specific_reservation")
        if Primitive.to_proto(self.commitment):
            request.resource.commitment = Primitive.to_proto(self.commitment)

        if Primitive.to_proto(self.specific_reservation_required):
            request.resource.specific_reservation_required = Primitive.to_proto(
                self.specific_reservation_required
            )

        if Primitive.to_proto(self.project):
            request.resource.project = Primitive.to_proto(self.project)

        request.service_account_file = self.service_account_file

        response = stub.ApplyComputeBetaReservation(request)
        self.id = Primitive.from_proto(response.id)
        self.self_link = Primitive.from_proto(response.self_link)
        self.zone = Primitive.from_proto(response.zone)
        self.description = Primitive.from_proto(response.description)
        self.name = Primitive.from_proto(response.name)
        self.specific_reservation = ReservationSpecificReservation.from_proto(
            response.specific_reservation
        )
        self.commitment = Primitive.from_proto(response.commitment)
        self.specific_reservation_required = Primitive.from_proto(
            response.specific_reservation_required
        )
        self.status = ReservationStatusEnum.from_proto(response.status)
        self.project = Primitive.from_proto(response.project)

    def delete(self):
        stub = reservation_pb2_grpc.ComputeBetaReservationServiceStub(channel.Channel())
        request = reservation_pb2.DeleteComputeBetaReservationRequest()
        request.service_account_file = self.service_account_file
        if Primitive.to_proto(self.id):
            request.resource.id = Primitive.to_proto(self.id)

        if Primitive.to_proto(self.zone):
            request.resource.zone = Primitive.to_proto(self.zone)

        if Primitive.to_proto(self.description):
            request.resource.description = Primitive.to_proto(self.description)

        if Primitive.to_proto(self.name):
            request.resource.name = Primitive.to_proto(self.name)

        if ReservationSpecificReservation.to_proto(self.specific_reservation):
            request.resource.specific_reservation.CopyFrom(
                ReservationSpecificReservation.to_proto(self.specific_reservation)
            )
        else:
            request.resource.ClearField("specific_reservation")
        if Primitive.to_proto(self.commitment):
            request.resource.commitment = Primitive.to_proto(self.commitment)

        if Primitive.to_proto(self.specific_reservation_required):
            request.resource.specific_reservation_required = Primitive.to_proto(
                self.specific_reservation_required
            )

        if Primitive.to_proto(self.project):
            request.resource.project = Primitive.to_proto(self.project)

        response = stub.DeleteComputeBetaReservation(request)

    @classmethod
    def list(self, project, zone, service_account_file=""):
        stub = reservation_pb2_grpc.ComputeBetaReservationServiceStub(channel.Channel())
        request = reservation_pb2.ListComputeBetaReservationRequest()
        request.service_account_file = service_account_file
        request.Project = project

        request.Zone = zone

        return stub.ListComputeBetaReservation(request).items

    def to_proto(self):
        resource = reservation_pb2.ComputeBetaReservation()
        if Primitive.to_proto(self.id):
            resource.id = Primitive.to_proto(self.id)
        if Primitive.to_proto(self.zone):
            resource.zone = Primitive.to_proto(self.zone)
        if Primitive.to_proto(self.description):
            resource.description = Primitive.to_proto(self.description)
        if Primitive.to_proto(self.name):
            resource.name = Primitive.to_proto(self.name)
        if ReservationSpecificReservation.to_proto(self.specific_reservation):
            resource.specific_reservation.CopyFrom(
                ReservationSpecificReservation.to_proto(self.specific_reservation)
            )
        else:
            resource.ClearField("specific_reservation")
        if Primitive.to_proto(self.commitment):
            resource.commitment = Primitive.to_proto(self.commitment)
        if Primitive.to_proto(self.specific_reservation_required):
            resource.specific_reservation_required = Primitive.to_proto(
                self.specific_reservation_required
            )
        if Primitive.to_proto(self.project):
            resource.project = Primitive.to_proto(self.project)
        return resource


class ReservationSpecificReservation(object):
    def __init__(
        self,
        instance_properties: dict = None,
        count: int = None,
        in_use_count: int = None,
    ):
        self.instance_properties = instance_properties
        self.count = count
        self.in_use_count = in_use_count

    @classmethod
    def to_proto(self, resource):
        if not resource:
            return None

        res = reservation_pb2.ComputeBetaReservationSpecificReservation()
        if ReservationSpecificReservationInstanceProperties.to_proto(
            resource.instance_properties
        ):
            res.instance_properties.CopyFrom(
                ReservationSpecificReservationInstanceProperties.to_proto(
                    resource.instance_properties
                )
            )
        else:
            res.ClearField("instance_properties")
        if Primitive.to_proto(resource.count):
            res.count = Primitive.to_proto(resource.count)
        if Primitive.to_proto(resource.in_use_count):
            res.in_use_count = Primitive.to_proto(resource.in_use_count)
        return res

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return None

        return ReservationSpecificReservation(
            instance_properties=ReservationSpecificReservationInstanceProperties.from_proto(
                resource.instance_properties
            ),
            count=Primitive.from_proto(resource.count),
            in_use_count=Primitive.from_proto(resource.in_use_count),
        )


class ReservationSpecificReservationArray(object):
    @classmethod
    def to_proto(self, resources):
        if not resources:
            return resources
        return [ReservationSpecificReservation.to_proto(i) for i in resources]

    @classmethod
    def from_proto(self, resources):
        return [ReservationSpecificReservation.from_proto(i) for i in resources]


class ReservationSpecificReservationInstanceProperties(object):
    def __init__(
        self,
        machine_type: str = None,
        guest_accelerators: list = None,
        min_cpu_platform: str = None,
        local_ssds: list = None,
    ):
        self.machine_type = machine_type
        self.guest_accelerators = guest_accelerators
        self.min_cpu_platform = min_cpu_platform
        self.local_ssds = local_ssds

    @classmethod
    def to_proto(self, resource):
        if not resource:
            return None

        res = (
            reservation_pb2.ComputeBetaReservationSpecificReservationInstanceProperties()
        )
        if Primitive.to_proto(resource.machine_type):
            res.machine_type = Primitive.to_proto(resource.machine_type)
        if ReservationSpecificReservationInstancePropertiesGuestAcceleratorsArray.to_proto(
            resource.guest_accelerators
        ):
            res.guest_accelerators.extend(
                ReservationSpecificReservationInstancePropertiesGuestAcceleratorsArray.to_proto(
                    resource.guest_accelerators
                )
            )
        if Primitive.to_proto(resource.min_cpu_platform):
            res.min_cpu_platform = Primitive.to_proto(resource.min_cpu_platform)
        if ReservationSpecificReservationInstancePropertiesLocalSsdsArray.to_proto(
            resource.local_ssds
        ):
            res.local_ssds.extend(
                ReservationSpecificReservationInstancePropertiesLocalSsdsArray.to_proto(
                    resource.local_ssds
                )
            )
        return res

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return None

        return ReservationSpecificReservationInstanceProperties(
            machine_type=Primitive.from_proto(resource.machine_type),
            guest_accelerators=ReservationSpecificReservationInstancePropertiesGuestAcceleratorsArray.from_proto(
                resource.guest_accelerators
            ),
            min_cpu_platform=Primitive.from_proto(resource.min_cpu_platform),
            local_ssds=ReservationSpecificReservationInstancePropertiesLocalSsdsArray.from_proto(
                resource.local_ssds
            ),
        )


class ReservationSpecificReservationInstancePropertiesArray(object):
    @classmethod
    def to_proto(self, resources):
        if not resources:
            return resources
        return [
            ReservationSpecificReservationInstanceProperties.to_proto(i)
            for i in resources
        ]

    @classmethod
    def from_proto(self, resources):
        return [
            ReservationSpecificReservationInstanceProperties.from_proto(i)
            for i in resources
        ]


class ReservationSpecificReservationInstancePropertiesGuestAccelerators(object):
    def __init__(self, accelerator_type: str = None, accelerator_count: int = None):
        self.accelerator_type = accelerator_type
        self.accelerator_count = accelerator_count

    @classmethod
    def to_proto(self, resource):
        if not resource:
            return None

        res = (
            reservation_pb2.ComputeBetaReservationSpecificReservationInstancePropertiesGuestAccelerators()
        )
        if Primitive.to_proto(resource.accelerator_type):
            res.accelerator_type = Primitive.to_proto(resource.accelerator_type)
        if Primitive.to_proto(resource.accelerator_count):
            res.accelerator_count = Primitive.to_proto(resource.accelerator_count)
        return res

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return None

        return ReservationSpecificReservationInstancePropertiesGuestAccelerators(
            accelerator_type=Primitive.from_proto(resource.accelerator_type),
            accelerator_count=Primitive.from_proto(resource.accelerator_count),
        )


class ReservationSpecificReservationInstancePropertiesGuestAcceleratorsArray(object):
    @classmethod
    def to_proto(self, resources):
        if not resources:
            return resources
        return [
            ReservationSpecificReservationInstancePropertiesGuestAccelerators.to_proto(
                i
            )
            for i in resources
        ]

    @classmethod
    def from_proto(self, resources):
        return [
            ReservationSpecificReservationInstancePropertiesGuestAccelerators.from_proto(
                i
            )
            for i in resources
        ]


class ReservationSpecificReservationInstancePropertiesLocalSsds(object):
    def __init__(self, disk_size_gb: int = None, interface: str = None):
        self.disk_size_gb = disk_size_gb
        self.interface = interface

    @classmethod
    def to_proto(self, resource):
        if not resource:
            return None

        res = (
            reservation_pb2.ComputeBetaReservationSpecificReservationInstancePropertiesLocalSsds()
        )
        if Primitive.to_proto(resource.disk_size_gb):
            res.disk_size_gb = Primitive.to_proto(resource.disk_size_gb)
        if ReservationSpecificReservationInstancePropertiesLocalSsdsInterfaceEnum.to_proto(
            resource.interface
        ):
            res.interface = ReservationSpecificReservationInstancePropertiesLocalSsdsInterfaceEnum.to_proto(
                resource.interface
            )
        return res

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return None

        return ReservationSpecificReservationInstancePropertiesLocalSsds(
            disk_size_gb=Primitive.from_proto(resource.disk_size_gb),
            interface=ReservationSpecificReservationInstancePropertiesLocalSsdsInterfaceEnum.from_proto(
                resource.interface
            ),
        )


class ReservationSpecificReservationInstancePropertiesLocalSsdsArray(object):
    @classmethod
    def to_proto(self, resources):
        if not resources:
            return resources
        return [
            ReservationSpecificReservationInstancePropertiesLocalSsds.to_proto(i)
            for i in resources
        ]

    @classmethod
    def from_proto(self, resources):
        return [
            ReservationSpecificReservationInstancePropertiesLocalSsds.from_proto(i)
            for i in resources
        ]


class ReservationSpecificReservationInstancePropertiesLocalSsdsInterfaceEnum(object):
    @classmethod
    def to_proto(self, resource):
        if not resource:
            return resource
        return reservation_pb2.ComputeBetaReservationSpecificReservationInstancePropertiesLocalSsdsInterfaceEnum.Value(
            "ComputeBetaReservationSpecificReservationInstancePropertiesLocalSsdsInterfaceEnum%s"
            % resource
        )

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return resource
        return reservation_pb2.ComputeBetaReservationSpecificReservationInstancePropertiesLocalSsdsInterfaceEnum.Name(
            resource
        )[
            len(
                "ComputeBetaReservationSpecificReservationInstancePropertiesLocalSsdsInterfaceEnum"
            ) :
        ]


class ReservationStatusEnum(object):
    @classmethod
    def to_proto(self, resource):
        if not resource:
            return resource
        return reservation_pb2.ComputeBetaReservationStatusEnum.Value(
            "ComputeBetaReservationStatusEnum%s" % resource
        )

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return resource
        return reservation_pb2.ComputeBetaReservationStatusEnum.Name(resource)[
            len("ComputeBetaReservationStatusEnum") :
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
