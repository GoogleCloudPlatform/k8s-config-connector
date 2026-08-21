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
from proto.redis import instance_pb2
from proto.redis import instance_pb2_grpc

from typing import List


class Instance(object):
    def __init__(
        self,
        name: str = None,
        display_name: str = None,
        labels: dict = None,
        location_id: str = None,
        alternative_location_id: str = None,
        redis_version: str = None,
        reserved_ip_range: str = None,
        host: str = None,
        port: int = None,
        current_location_id: str = None,
        create_time: str = None,
        state: str = None,
        status_message: str = None,
        redis_configs: dict = None,
        tier: str = None,
        memory_size_gb: int = None,
        authorized_network: str = None,
        persistence_iam_identity: str = None,
        connect_mode: str = None,
        auth_enabled: bool = None,
        server_ca_certs: list = None,
        transit_encryption_mode: str = None,
        maintenance_policy: dict = None,
        maintenance_schedule: dict = None,
        project: str = None,
        location: str = None,
        service_account_file: str = "",
    ):

        channel.initialize()
        self.name = name
        self.display_name = display_name
        self.labels = labels
        self.location_id = location_id
        self.alternative_location_id = alternative_location_id
        self.redis_version = redis_version
        self.reserved_ip_range = reserved_ip_range
        self.redis_configs = redis_configs
        self.tier = tier
        self.memory_size_gb = memory_size_gb
        self.authorized_network = authorized_network
        self.connect_mode = connect_mode
        self.auth_enabled = auth_enabled
        self.transit_encryption_mode = transit_encryption_mode
        self.maintenance_policy = maintenance_policy
        self.project = project
        self.location = location
        self.service_account_file = service_account_file

    def apply(self):
        stub = instance_pb2_grpc.RedisBetaInstanceServiceStub(channel.Channel())
        request = instance_pb2.ApplyRedisBetaInstanceRequest()
        if Primitive.to_proto(self.name):
            request.resource.name = Primitive.to_proto(self.name)

        if Primitive.to_proto(self.display_name):
            request.resource.display_name = Primitive.to_proto(self.display_name)

        if Primitive.to_proto(self.labels):
            request.resource.labels = Primitive.to_proto(self.labels)

        if Primitive.to_proto(self.location_id):
            request.resource.location_id = Primitive.to_proto(self.location_id)

        if Primitive.to_proto(self.alternative_location_id):
            request.resource.alternative_location_id = Primitive.to_proto(
                self.alternative_location_id
            )

        if Primitive.to_proto(self.redis_version):
            request.resource.redis_version = Primitive.to_proto(self.redis_version)

        if Primitive.to_proto(self.reserved_ip_range):
            request.resource.reserved_ip_range = Primitive.to_proto(
                self.reserved_ip_range
            )

        if Primitive.to_proto(self.redis_configs):
            request.resource.redis_configs = Primitive.to_proto(self.redis_configs)

        if InstanceTierEnum.to_proto(self.tier):
            request.resource.tier = InstanceTierEnum.to_proto(self.tier)

        if Primitive.to_proto(self.memory_size_gb):
            request.resource.memory_size_gb = Primitive.to_proto(self.memory_size_gb)

        if Primitive.to_proto(self.authorized_network):
            request.resource.authorized_network = Primitive.to_proto(
                self.authorized_network
            )

        if InstanceConnectModeEnum.to_proto(self.connect_mode):
            request.resource.connect_mode = InstanceConnectModeEnum.to_proto(
                self.connect_mode
            )

        if Primitive.to_proto(self.auth_enabled):
            request.resource.auth_enabled = Primitive.to_proto(self.auth_enabled)

        if InstanceTransitEncryptionModeEnum.to_proto(self.transit_encryption_mode):
            request.resource.transit_encryption_mode = InstanceTransitEncryptionModeEnum.to_proto(
                self.transit_encryption_mode
            )

        if InstanceMaintenancePolicy.to_proto(self.maintenance_policy):
            request.resource.maintenance_policy.CopyFrom(
                InstanceMaintenancePolicy.to_proto(self.maintenance_policy)
            )
        else:
            request.resource.ClearField("maintenance_policy")
        if Primitive.to_proto(self.project):
            request.resource.project = Primitive.to_proto(self.project)

        if Primitive.to_proto(self.location):
            request.resource.location = Primitive.to_proto(self.location)

        request.service_account_file = self.service_account_file

        response = stub.ApplyRedisBetaInstance(request)
        self.name = Primitive.from_proto(response.name)
        self.display_name = Primitive.from_proto(response.display_name)
        self.labels = Primitive.from_proto(response.labels)
        self.location_id = Primitive.from_proto(response.location_id)
        self.alternative_location_id = Primitive.from_proto(
            response.alternative_location_id
        )
        self.redis_version = Primitive.from_proto(response.redis_version)
        self.reserved_ip_range = Primitive.from_proto(response.reserved_ip_range)
        self.host = Primitive.from_proto(response.host)
        self.port = Primitive.from_proto(response.port)
        self.current_location_id = Primitive.from_proto(response.current_location_id)
        self.create_time = Primitive.from_proto(response.create_time)
        self.state = InstanceStateEnum.from_proto(response.state)
        self.status_message = Primitive.from_proto(response.status_message)
        self.redis_configs = Primitive.from_proto(response.redis_configs)
        self.tier = InstanceTierEnum.from_proto(response.tier)
        self.memory_size_gb = Primitive.from_proto(response.memory_size_gb)
        self.authorized_network = Primitive.from_proto(response.authorized_network)
        self.persistence_iam_identity = Primitive.from_proto(
            response.persistence_iam_identity
        )
        self.connect_mode = InstanceConnectModeEnum.from_proto(response.connect_mode)
        self.auth_enabled = Primitive.from_proto(response.auth_enabled)
        self.server_ca_certs = InstanceServerCaCertsArray.from_proto(
            response.server_ca_certs
        )
        self.transit_encryption_mode = InstanceTransitEncryptionModeEnum.from_proto(
            response.transit_encryption_mode
        )
        self.maintenance_policy = InstanceMaintenancePolicy.from_proto(
            response.maintenance_policy
        )
        self.maintenance_schedule = InstanceMaintenanceSchedule.from_proto(
            response.maintenance_schedule
        )
        self.project = Primitive.from_proto(response.project)
        self.location = Primitive.from_proto(response.location)

    def delete(self):
        stub = instance_pb2_grpc.RedisBetaInstanceServiceStub(channel.Channel())
        request = instance_pb2.DeleteRedisBetaInstanceRequest()
        request.service_account_file = self.service_account_file
        if Primitive.to_proto(self.name):
            request.resource.name = Primitive.to_proto(self.name)

        if Primitive.to_proto(self.display_name):
            request.resource.display_name = Primitive.to_proto(self.display_name)

        if Primitive.to_proto(self.labels):
            request.resource.labels = Primitive.to_proto(self.labels)

        if Primitive.to_proto(self.location_id):
            request.resource.location_id = Primitive.to_proto(self.location_id)

        if Primitive.to_proto(self.alternative_location_id):
            request.resource.alternative_location_id = Primitive.to_proto(
                self.alternative_location_id
            )

        if Primitive.to_proto(self.redis_version):
            request.resource.redis_version = Primitive.to_proto(self.redis_version)

        if Primitive.to_proto(self.reserved_ip_range):
            request.resource.reserved_ip_range = Primitive.to_proto(
                self.reserved_ip_range
            )

        if Primitive.to_proto(self.redis_configs):
            request.resource.redis_configs = Primitive.to_proto(self.redis_configs)

        if InstanceTierEnum.to_proto(self.tier):
            request.resource.tier = InstanceTierEnum.to_proto(self.tier)

        if Primitive.to_proto(self.memory_size_gb):
            request.resource.memory_size_gb = Primitive.to_proto(self.memory_size_gb)

        if Primitive.to_proto(self.authorized_network):
            request.resource.authorized_network = Primitive.to_proto(
                self.authorized_network
            )

        if InstanceConnectModeEnum.to_proto(self.connect_mode):
            request.resource.connect_mode = InstanceConnectModeEnum.to_proto(
                self.connect_mode
            )

        if Primitive.to_proto(self.auth_enabled):
            request.resource.auth_enabled = Primitive.to_proto(self.auth_enabled)

        if InstanceTransitEncryptionModeEnum.to_proto(self.transit_encryption_mode):
            request.resource.transit_encryption_mode = InstanceTransitEncryptionModeEnum.to_proto(
                self.transit_encryption_mode
            )

        if InstanceMaintenancePolicy.to_proto(self.maintenance_policy):
            request.resource.maintenance_policy.CopyFrom(
                InstanceMaintenancePolicy.to_proto(self.maintenance_policy)
            )
        else:
            request.resource.ClearField("maintenance_policy")
        if Primitive.to_proto(self.project):
            request.resource.project = Primitive.to_proto(self.project)

        if Primitive.to_proto(self.location):
            request.resource.location = Primitive.to_proto(self.location)

        response = stub.DeleteRedisBetaInstance(request)

    @classmethod
    def list(self, project, location, service_account_file=""):
        stub = instance_pb2_grpc.RedisBetaInstanceServiceStub(channel.Channel())
        request = instance_pb2.ListRedisBetaInstanceRequest()
        request.service_account_file = service_account_file
        request.Project = project

        request.Location = location

        return stub.ListRedisBetaInstance(request).items

    def to_proto(self):
        resource = instance_pb2.RedisBetaInstance()
        if Primitive.to_proto(self.name):
            resource.name = Primitive.to_proto(self.name)
        if Primitive.to_proto(self.display_name):
            resource.display_name = Primitive.to_proto(self.display_name)
        if Primitive.to_proto(self.labels):
            resource.labels = Primitive.to_proto(self.labels)
        if Primitive.to_proto(self.location_id):
            resource.location_id = Primitive.to_proto(self.location_id)
        if Primitive.to_proto(self.alternative_location_id):
            resource.alternative_location_id = Primitive.to_proto(
                self.alternative_location_id
            )
        if Primitive.to_proto(self.redis_version):
            resource.redis_version = Primitive.to_proto(self.redis_version)
        if Primitive.to_proto(self.reserved_ip_range):
            resource.reserved_ip_range = Primitive.to_proto(self.reserved_ip_range)
        if Primitive.to_proto(self.redis_configs):
            resource.redis_configs = Primitive.to_proto(self.redis_configs)
        if InstanceTierEnum.to_proto(self.tier):
            resource.tier = InstanceTierEnum.to_proto(self.tier)
        if Primitive.to_proto(self.memory_size_gb):
            resource.memory_size_gb = Primitive.to_proto(self.memory_size_gb)
        if Primitive.to_proto(self.authorized_network):
            resource.authorized_network = Primitive.to_proto(self.authorized_network)
        if InstanceConnectModeEnum.to_proto(self.connect_mode):
            resource.connect_mode = InstanceConnectModeEnum.to_proto(self.connect_mode)
        if Primitive.to_proto(self.auth_enabled):
            resource.auth_enabled = Primitive.to_proto(self.auth_enabled)
        if InstanceTransitEncryptionModeEnum.to_proto(self.transit_encryption_mode):
            resource.transit_encryption_mode = InstanceTransitEncryptionModeEnum.to_proto(
                self.transit_encryption_mode
            )
        if InstanceMaintenancePolicy.to_proto(self.maintenance_policy):
            resource.maintenance_policy.CopyFrom(
                InstanceMaintenancePolicy.to_proto(self.maintenance_policy)
            )
        else:
            resource.ClearField("maintenance_policy")
        if Primitive.to_proto(self.project):
            resource.project = Primitive.to_proto(self.project)
        if Primitive.to_proto(self.location):
            resource.location = Primitive.to_proto(self.location)
        return resource


class InstanceServerCaCerts(object):
    def __init__(
        self,
        serial_number: str = None,
        cert: str = None,
        create_time: str = None,
        expire_time: str = None,
        sha1_fingerprint: str = None,
    ):
        self.serial_number = serial_number
        self.cert = cert
        self.create_time = create_time
        self.expire_time = expire_time
        self.sha1_fingerprint = sha1_fingerprint

    @classmethod
    def to_proto(self, resource):
        if not resource:
            return None

        res = instance_pb2.RedisBetaInstanceServerCaCerts()
        if Primitive.to_proto(resource.serial_number):
            res.serial_number = Primitive.to_proto(resource.serial_number)
        if Primitive.to_proto(resource.cert):
            res.cert = Primitive.to_proto(resource.cert)
        if Primitive.to_proto(resource.create_time):
            res.create_time = Primitive.to_proto(resource.create_time)
        if Primitive.to_proto(resource.expire_time):
            res.expire_time = Primitive.to_proto(resource.expire_time)
        if Primitive.to_proto(resource.sha1_fingerprint):
            res.sha1_fingerprint = Primitive.to_proto(resource.sha1_fingerprint)
        return res

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return None

        return InstanceServerCaCerts(
            serial_number=Primitive.from_proto(resource.serial_number),
            cert=Primitive.from_proto(resource.cert),
            create_time=Primitive.from_proto(resource.create_time),
            expire_time=Primitive.from_proto(resource.expire_time),
            sha1_fingerprint=Primitive.from_proto(resource.sha1_fingerprint),
        )


class InstanceServerCaCertsArray(object):
    @classmethod
    def to_proto(self, resources):
        if not resources:
            return resources
        return [InstanceServerCaCerts.to_proto(i) for i in resources]

    @classmethod
    def from_proto(self, resources):
        return [InstanceServerCaCerts.from_proto(i) for i in resources]


class InstanceMaintenancePolicy(object):
    def __init__(
        self,
        create_time: str = None,
        update_time: str = None,
        description: str = None,
        weekly_maintenance_window: list = None,
    ):
        self.create_time = create_time
        self.update_time = update_time
        self.description = description
        self.weekly_maintenance_window = weekly_maintenance_window

    @classmethod
    def to_proto(self, resource):
        if not resource:
            return None

        res = instance_pb2.RedisBetaInstanceMaintenancePolicy()
        if Primitive.to_proto(resource.create_time):
            res.create_time = Primitive.to_proto(resource.create_time)
        if Primitive.to_proto(resource.update_time):
            res.update_time = Primitive.to_proto(resource.update_time)
        if Primitive.to_proto(resource.description):
            res.description = Primitive.to_proto(resource.description)
        if InstanceMaintenancePolicyWeeklyMaintenanceWindowArray.to_proto(
            resource.weekly_maintenance_window
        ):
            res.weekly_maintenance_window.extend(
                InstanceMaintenancePolicyWeeklyMaintenanceWindowArray.to_proto(
                    resource.weekly_maintenance_window
                )
            )
        return res

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return None

        return InstanceMaintenancePolicy(
            create_time=Primitive.from_proto(resource.create_time),
            update_time=Primitive.from_proto(resource.update_time),
            description=Primitive.from_proto(resource.description),
            weekly_maintenance_window=InstanceMaintenancePolicyWeeklyMaintenanceWindowArray.from_proto(
                resource.weekly_maintenance_window
            ),
        )


class InstanceMaintenancePolicyArray(object):
    @classmethod
    def to_proto(self, resources):
        if not resources:
            return resources
        return [InstanceMaintenancePolicy.to_proto(i) for i in resources]

    @classmethod
    def from_proto(self, resources):
        return [InstanceMaintenancePolicy.from_proto(i) for i in resources]


class InstanceMaintenancePolicyWeeklyMaintenanceWindow(object):
    def __init__(self, day: str = None, start_time: dict = None, duration: str = None):
        self.day = day
        self.start_time = start_time
        self.duration = duration

    @classmethod
    def to_proto(self, resource):
        if not resource:
            return None

        res = instance_pb2.RedisBetaInstanceMaintenancePolicyWeeklyMaintenanceWindow()
        if InstanceMaintenancePolicyWeeklyMaintenanceWindowDayEnum.to_proto(
            resource.day
        ):
            res.day = InstanceMaintenancePolicyWeeklyMaintenanceWindowDayEnum.to_proto(
                resource.day
            )
        if InstanceMaintenancePolicyWeeklyMaintenanceWindowStartTime.to_proto(
            resource.start_time
        ):
            res.start_time.CopyFrom(
                InstanceMaintenancePolicyWeeklyMaintenanceWindowStartTime.to_proto(
                    resource.start_time
                )
            )
        else:
            res.ClearField("start_time")
        if Primitive.to_proto(resource.duration):
            res.duration = Primitive.to_proto(resource.duration)
        return res

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return None

        return InstanceMaintenancePolicyWeeklyMaintenanceWindow(
            day=InstanceMaintenancePolicyWeeklyMaintenanceWindowDayEnum.from_proto(
                resource.day
            ),
            start_time=InstanceMaintenancePolicyWeeklyMaintenanceWindowStartTime.from_proto(
                resource.start_time
            ),
            duration=Primitive.from_proto(resource.duration),
        )


class InstanceMaintenancePolicyWeeklyMaintenanceWindowArray(object):
    @classmethod
    def to_proto(self, resources):
        if not resources:
            return resources
        return [
            InstanceMaintenancePolicyWeeklyMaintenanceWindow.to_proto(i)
            for i in resources
        ]

    @classmethod
    def from_proto(self, resources):
        return [
            InstanceMaintenancePolicyWeeklyMaintenanceWindow.from_proto(i)
            for i in resources
        ]


class InstanceMaintenancePolicyWeeklyMaintenanceWindowStartTime(object):
    def __init__(
        self,
        hours: int = None,
        minutes: int = None,
        seconds: int = None,
        nanos: int = None,
    ):
        self.hours = hours
        self.minutes = minutes
        self.seconds = seconds
        self.nanos = nanos

    @classmethod
    def to_proto(self, resource):
        if not resource:
            return None

        res = (
            instance_pb2.RedisBetaInstanceMaintenancePolicyWeeklyMaintenanceWindowStartTime()
        )
        if Primitive.to_proto(resource.hours):
            res.hours = Primitive.to_proto(resource.hours)
        if Primitive.to_proto(resource.minutes):
            res.minutes = Primitive.to_proto(resource.minutes)
        if Primitive.to_proto(resource.seconds):
            res.seconds = Primitive.to_proto(resource.seconds)
        if Primitive.to_proto(resource.nanos):
            res.nanos = Primitive.to_proto(resource.nanos)
        return res

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return None

        return InstanceMaintenancePolicyWeeklyMaintenanceWindowStartTime(
            hours=Primitive.from_proto(resource.hours),
            minutes=Primitive.from_proto(resource.minutes),
            seconds=Primitive.from_proto(resource.seconds),
            nanos=Primitive.from_proto(resource.nanos),
        )


class InstanceMaintenancePolicyWeeklyMaintenanceWindowStartTimeArray(object):
    @classmethod
    def to_proto(self, resources):
        if not resources:
            return resources
        return [
            InstanceMaintenancePolicyWeeklyMaintenanceWindowStartTime.to_proto(i)
            for i in resources
        ]

    @classmethod
    def from_proto(self, resources):
        return [
            InstanceMaintenancePolicyWeeklyMaintenanceWindowStartTime.from_proto(i)
            for i in resources
        ]


class InstanceMaintenanceSchedule(object):
    def __init__(
        self,
        start_time: str = None,
        end_time: str = None,
        can_reschedule: bool = None,
        schedule_deadline_time: str = None,
    ):
        self.start_time = start_time
        self.end_time = end_time
        self.can_reschedule = can_reschedule
        self.schedule_deadline_time = schedule_deadline_time

    @classmethod
    def to_proto(self, resource):
        if not resource:
            return None

        res = instance_pb2.RedisBetaInstanceMaintenanceSchedule()
        if Primitive.to_proto(resource.start_time):
            res.start_time = Primitive.to_proto(resource.start_time)
        if Primitive.to_proto(resource.end_time):
            res.end_time = Primitive.to_proto(resource.end_time)
        if Primitive.to_proto(resource.can_reschedule):
            res.can_reschedule = Primitive.to_proto(resource.can_reschedule)
        if Primitive.to_proto(resource.schedule_deadline_time):
            res.schedule_deadline_time = Primitive.to_proto(
                resource.schedule_deadline_time
            )
        return res

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return None

        return InstanceMaintenanceSchedule(
            start_time=Primitive.from_proto(resource.start_time),
            end_time=Primitive.from_proto(resource.end_time),
            can_reschedule=Primitive.from_proto(resource.can_reschedule),
            schedule_deadline_time=Primitive.from_proto(
                resource.schedule_deadline_time
            ),
        )


class InstanceMaintenanceScheduleArray(object):
    @classmethod
    def to_proto(self, resources):
        if not resources:
            return resources
        return [InstanceMaintenanceSchedule.to_proto(i) for i in resources]

    @classmethod
    def from_proto(self, resources):
        return [InstanceMaintenanceSchedule.from_proto(i) for i in resources]


class InstanceStateEnum(object):
    @classmethod
    def to_proto(self, resource):
        if not resource:
            return resource
        return instance_pb2.RedisBetaInstanceStateEnum.Value(
            "RedisBetaInstanceStateEnum%s" % resource
        )

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return resource
        return instance_pb2.RedisBetaInstanceStateEnum.Name(resource)[
            len("RedisBetaInstanceStateEnum") :
        ]


class InstanceTierEnum(object):
    @classmethod
    def to_proto(self, resource):
        if not resource:
            return resource
        return instance_pb2.RedisBetaInstanceTierEnum.Value(
            "RedisBetaInstanceTierEnum%s" % resource
        )

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return resource
        return instance_pb2.RedisBetaInstanceTierEnum.Name(resource)[
            len("RedisBetaInstanceTierEnum") :
        ]


class InstanceConnectModeEnum(object):
    @classmethod
    def to_proto(self, resource):
        if not resource:
            return resource
        return instance_pb2.RedisBetaInstanceConnectModeEnum.Value(
            "RedisBetaInstanceConnectModeEnum%s" % resource
        )

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return resource
        return instance_pb2.RedisBetaInstanceConnectModeEnum.Name(resource)[
            len("RedisBetaInstanceConnectModeEnum") :
        ]


class InstanceTransitEncryptionModeEnum(object):
    @classmethod
    def to_proto(self, resource):
        if not resource:
            return resource
        return instance_pb2.RedisBetaInstanceTransitEncryptionModeEnum.Value(
            "RedisBetaInstanceTransitEncryptionModeEnum%s" % resource
        )

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return resource
        return instance_pb2.RedisBetaInstanceTransitEncryptionModeEnum.Name(resource)[
            len("RedisBetaInstanceTransitEncryptionModeEnum") :
        ]


class InstanceMaintenancePolicyWeeklyMaintenanceWindowDayEnum(object):
    @classmethod
    def to_proto(self, resource):
        if not resource:
            return resource
        return instance_pb2.RedisBetaInstanceMaintenancePolicyWeeklyMaintenanceWindowDayEnum.Value(
            "RedisBetaInstanceMaintenancePolicyWeeklyMaintenanceWindowDayEnum%s"
            % resource
        )

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return resource
        return instance_pb2.RedisBetaInstanceMaintenancePolicyWeeklyMaintenanceWindowDayEnum.Name(
            resource
        )[
            len("RedisBetaInstanceMaintenancePolicyWeeklyMaintenanceWindowDayEnum") :
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
