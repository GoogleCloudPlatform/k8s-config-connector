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
from google3.cloud.graphite.mmv2.services.google.sql import instance_pb2
from google3.cloud.graphite.mmv2.services.google.sql import instance_pb2_grpc

from typing import List


class Instance(object):
    def __init__(
        self,
        backend_type: str = None,
        connection_name: str = None,
        database_version: str = None,
        etag: str = None,
        gce_zone: str = None,
        instance_type: str = None,
        master_instance_name: str = None,
        max_disk_size: dict = None,
        name: str = None,
        project: str = None,
        region: str = None,
        root_password: str = None,
        current_disk_size: dict = None,
        disk_encryption_configuration: dict = None,
        failover_replica: dict = None,
        ip_addresses: list = None,
        master_instance: dict = None,
        replica_configuration: dict = None,
        scheduled_maintenance: dict = None,
        settings: dict = None,
        service_account_file: str = "",
    ):

        channel.initialize()
        self.backend_type = backend_type
        self.connection_name = connection_name
        self.database_version = database_version
        self.etag = etag
        self.gce_zone = gce_zone
        self.instance_type = instance_type
        self.master_instance_name = master_instance_name
        self.max_disk_size = max_disk_size
        self.name = name
        self.project = project
        self.region = region
        self.root_password = root_password
        self.current_disk_size = current_disk_size
        self.disk_encryption_configuration = disk_encryption_configuration
        self.failover_replica = failover_replica
        self.ip_addresses = ip_addresses
        self.master_instance = master_instance
        self.replica_configuration = replica_configuration
        self.scheduled_maintenance = scheduled_maintenance
        self.settings = settings
        self.service_account_file = service_account_file

    def apply(self):
        stub = instance_pb2_grpc.SqlInstanceServiceStub(channel.Channel())
        request = instance_pb2.ApplySqlInstanceRequest()
        if InstanceBackendTypeEnum.to_proto(self.backend_type):
            request.resource.backend_type = InstanceBackendTypeEnum.to_proto(
                self.backend_type
            )

        if Primitive.to_proto(self.connection_name):
            request.resource.connection_name = Primitive.to_proto(self.connection_name)

        if InstanceDatabaseVersionEnum.to_proto(self.database_version):
            request.resource.database_version = InstanceDatabaseVersionEnum.to_proto(
                self.database_version
            )

        if Primitive.to_proto(self.etag):
            request.resource.etag = Primitive.to_proto(self.etag)

        if Primitive.to_proto(self.gce_zone):
            request.resource.gce_zone = Primitive.to_proto(self.gce_zone)

        if InstanceInstanceTypeEnum.to_proto(self.instance_type):
            request.resource.instance_type = InstanceInstanceTypeEnum.to_proto(
                self.instance_type
            )

        if Primitive.to_proto(self.master_instance_name):
            request.resource.master_instance_name = Primitive.to_proto(
                self.master_instance_name
            )

        if InstanceMaxDiskSize.to_proto(self.max_disk_size):
            request.resource.max_disk_size.CopyFrom(
                InstanceMaxDiskSize.to_proto(self.max_disk_size)
            )
        else:
            request.resource.ClearField("max_disk_size")
        if Primitive.to_proto(self.name):
            request.resource.name = Primitive.to_proto(self.name)

        if Primitive.to_proto(self.project):
            request.resource.project = Primitive.to_proto(self.project)

        if Primitive.to_proto(self.region):
            request.resource.region = Primitive.to_proto(self.region)

        if Primitive.to_proto(self.root_password):
            request.resource.root_password = Primitive.to_proto(self.root_password)

        if InstanceCurrentDiskSize.to_proto(self.current_disk_size):
            request.resource.current_disk_size.CopyFrom(
                InstanceCurrentDiskSize.to_proto(self.current_disk_size)
            )
        else:
            request.resource.ClearField("current_disk_size")
        if InstanceDiskEncryptionConfiguration.to_proto(
            self.disk_encryption_configuration
        ):
            request.resource.disk_encryption_configuration.CopyFrom(
                InstanceDiskEncryptionConfiguration.to_proto(
                    self.disk_encryption_configuration
                )
            )
        else:
            request.resource.ClearField("disk_encryption_configuration")
        if InstanceFailoverReplica.to_proto(self.failover_replica):
            request.resource.failover_replica.CopyFrom(
                InstanceFailoverReplica.to_proto(self.failover_replica)
            )
        else:
            request.resource.ClearField("failover_replica")
        if InstanceIPAddressesArray.to_proto(self.ip_addresses):
            request.resource.ip_addresses.extend(
                InstanceIPAddressesArray.to_proto(self.ip_addresses)
            )
        if InstanceMasterInstance.to_proto(self.master_instance):
            request.resource.master_instance.CopyFrom(
                InstanceMasterInstance.to_proto(self.master_instance)
            )
        else:
            request.resource.ClearField("master_instance")
        if InstanceReplicaConfiguration.to_proto(self.replica_configuration):
            request.resource.replica_configuration.CopyFrom(
                InstanceReplicaConfiguration.to_proto(self.replica_configuration)
            )
        else:
            request.resource.ClearField("replica_configuration")
        if InstanceScheduledMaintenance.to_proto(self.scheduled_maintenance):
            request.resource.scheduled_maintenance.CopyFrom(
                InstanceScheduledMaintenance.to_proto(self.scheduled_maintenance)
            )
        else:
            request.resource.ClearField("scheduled_maintenance")
        if InstanceSettings.to_proto(self.settings):
            request.resource.settings.CopyFrom(InstanceSettings.to_proto(self.settings))
        else:
            request.resource.ClearField("settings")
        request.service_account_file = self.service_account_file

        response = stub.ApplySqlInstance(request)
        self.backend_type = InstanceBackendTypeEnum.from_proto(response.backend_type)
        self.connection_name = Primitive.from_proto(response.connection_name)
        self.database_version = InstanceDatabaseVersionEnum.from_proto(
            response.database_version
        )
        self.etag = Primitive.from_proto(response.etag)
        self.gce_zone = Primitive.from_proto(response.gce_zone)
        self.instance_type = InstanceInstanceTypeEnum.from_proto(response.instance_type)
        self.master_instance_name = Primitive.from_proto(response.master_instance_name)
        self.max_disk_size = InstanceMaxDiskSize.from_proto(response.max_disk_size)
        self.name = Primitive.from_proto(response.name)
        self.project = Primitive.from_proto(response.project)
        self.region = Primitive.from_proto(response.region)
        self.root_password = Primitive.from_proto(response.root_password)
        self.current_disk_size = InstanceCurrentDiskSize.from_proto(
            response.current_disk_size
        )
        self.disk_encryption_configuration = InstanceDiskEncryptionConfiguration.from_proto(
            response.disk_encryption_configuration
        )
        self.failover_replica = InstanceFailoverReplica.from_proto(
            response.failover_replica
        )
        self.ip_addresses = InstanceIPAddressesArray.from_proto(response.ip_addresses)
        self.master_instance = InstanceMasterInstance.from_proto(
            response.master_instance
        )
        self.replica_configuration = InstanceReplicaConfiguration.from_proto(
            response.replica_configuration
        )
        self.scheduled_maintenance = InstanceScheduledMaintenance.from_proto(
            response.scheduled_maintenance
        )
        self.settings = InstanceSettings.from_proto(response.settings)

    @classmethod
    def delete(self, project, name, service_account_file=""):
        stub = instance_pb2_grpc.SqlInstanceServiceStub(channel.Channel())
        request = instance_pb2.DeleteSqlInstanceRequest()
        request.service_account_file = service_account_file
        request.Project = project

        request.Name = name

        response = stub.DeleteSqlInstance(request)

    @classmethod
    def list(self, project, service_account_file=""):
        stub = instance_pb2_grpc.SqlInstanceServiceStub(channel.Channel())
        request = instance_pb2.ListSqlInstanceRequest()
        request.service_account_file = service_account_file
        request.Project = project

        return stub.ListSqlInstance(request).items

    @classmethod
    def from_any(self, any_proto):
        # Marshal any proto to regular proto.
        res_proto = instance_pb2.SqlInstance()
        any_proto.Unpack(res_proto)

        res = Instance()
        res.backend_type = InstanceBackendTypeEnum.from_proto(res_proto.backend_type)
        res.connection_name = Primitive.from_proto(res_proto.connection_name)
        res.database_version = InstanceDatabaseVersionEnum.from_proto(
            res_proto.database_version
        )
        res.etag = Primitive.from_proto(res_proto.etag)
        res.gce_zone = Primitive.from_proto(res_proto.gce_zone)
        res.instance_type = InstanceInstanceTypeEnum.from_proto(res_proto.instance_type)
        res.master_instance_name = Primitive.from_proto(res_proto.master_instance_name)
        res.max_disk_size = InstanceMaxDiskSize.from_proto(res_proto.max_disk_size)
        res.name = Primitive.from_proto(res_proto.name)
        res.project = Primitive.from_proto(res_proto.project)
        res.region = Primitive.from_proto(res_proto.region)
        res.root_password = Primitive.from_proto(res_proto.root_password)
        res.current_disk_size = InstanceCurrentDiskSize.from_proto(
            res_proto.current_disk_size
        )
        res.disk_encryption_configuration = InstanceDiskEncryptionConfiguration.from_proto(
            res_proto.disk_encryption_configuration
        )
        res.failover_replica = InstanceFailoverReplica.from_proto(
            res_proto.failover_replica
        )
        res.ip_addresses = InstanceIPAddressesArray.from_proto(res_proto.ip_addresses)
        res.master_instance = InstanceMasterInstance.from_proto(
            res_proto.master_instance
        )
        res.replica_configuration = InstanceReplicaConfiguration.from_proto(
            res_proto.replica_configuration
        )
        res.scheduled_maintenance = InstanceScheduledMaintenance.from_proto(
            res_proto.scheduled_maintenance
        )
        res.settings = InstanceSettings.from_proto(res_proto.settings)
        return res


class InstanceMaxDiskSize(object):
    def __init__(self, value: int = None):
        self.value = value

    @classmethod
    def to_proto(self, resource):
        if not resource:
            return None

        res = instance_pb2.SqlInstanceMaxDiskSize()
        if Primitive.to_proto(resource.value):
            res.value = Primitive.to_proto(resource.value)
        return res

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return None

        return InstanceMaxDiskSize(value=resource.value,)


class InstanceMaxDiskSizeArray(object):
    @classmethod
    def to_proto(self, resources):
        if not resources:
            return resources
        return [InstanceMaxDiskSize.to_proto(i) for i in resources]

    @classmethod
    def from_proto(self, resources):
        return [InstanceMaxDiskSize.from_proto(i) for i in resources]


class InstanceCurrentDiskSize(object):
    def __init__(self, value: int = None):
        self.value = value

    @classmethod
    def to_proto(self, resource):
        if not resource:
            return None

        res = instance_pb2.SqlInstanceCurrentDiskSize()
        if Primitive.to_proto(resource.value):
            res.value = Primitive.to_proto(resource.value)
        return res

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return None

        return InstanceCurrentDiskSize(value=resource.value,)


class InstanceCurrentDiskSizeArray(object):
    @classmethod
    def to_proto(self, resources):
        if not resources:
            return resources
        return [InstanceCurrentDiskSize.to_proto(i) for i in resources]

    @classmethod
    def from_proto(self, resources):
        return [InstanceCurrentDiskSize.from_proto(i) for i in resources]


class InstanceDiskEncryptionConfiguration(object):
    def __init__(self, kms_key_name: str = None, kind: str = None):
        self.kms_key_name = kms_key_name
        self.kind = kind

    @classmethod
    def to_proto(self, resource):
        if not resource:
            return None

        res = instance_pb2.SqlInstanceDiskEncryptionConfiguration()
        if Primitive.to_proto(resource.kms_key_name):
            res.kms_key_name = Primitive.to_proto(resource.kms_key_name)
        if Primitive.to_proto(resource.kind):
            res.kind = Primitive.to_proto(resource.kind)
        return res

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return None

        return InstanceDiskEncryptionConfiguration(
            kms_key_name=resource.kms_key_name, kind=resource.kind,
        )


class InstanceDiskEncryptionConfigurationArray(object):
    @classmethod
    def to_proto(self, resources):
        if not resources:
            return resources
        return [InstanceDiskEncryptionConfiguration.to_proto(i) for i in resources]

    @classmethod
    def from_proto(self, resources):
        return [InstanceDiskEncryptionConfiguration.from_proto(i) for i in resources]


class InstanceFailoverReplica(object):
    def __init__(self, name: str = None, available: bool = None):
        self.name = name
        self.available = available

    @classmethod
    def to_proto(self, resource):
        if not resource:
            return None

        res = instance_pb2.SqlInstanceFailoverReplica()
        if Primitive.to_proto(resource.name):
            res.name = Primitive.to_proto(resource.name)
        if Primitive.to_proto(resource.available):
            res.available = Primitive.to_proto(resource.available)
        return res

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return None

        return InstanceFailoverReplica(
            name=resource.name, available=resource.available,
        )


class InstanceFailoverReplicaArray(object):
    @classmethod
    def to_proto(self, resources):
        if not resources:
            return resources
        return [InstanceFailoverReplica.to_proto(i) for i in resources]

    @classmethod
    def from_proto(self, resources):
        return [InstanceFailoverReplica.from_proto(i) for i in resources]


class InstanceIPAddresses(object):
    def __init__(
        self, type: str = None, ip_address: str = None, time_to_retire: dict = None
    ):
        self.type = type
        self.ip_address = ip_address
        self.time_to_retire = time_to_retire

    @classmethod
    def to_proto(self, resource):
        if not resource:
            return None

        res = instance_pb2.SqlInstanceIPAddresses()
        if InstanceIPAddressesTypeEnum.to_proto(resource.type):
            res.type = InstanceIPAddressesTypeEnum.to_proto(resource.type)
        if Primitive.to_proto(resource.ip_address):
            res.ip_address = Primitive.to_proto(resource.ip_address)
        if InstanceIPAddressesTimeToRetire.to_proto(resource.time_to_retire):
            res.time_to_retire.CopyFrom(
                InstanceIPAddressesTimeToRetire.to_proto(resource.time_to_retire)
            )
        else:
            res.ClearField("time_to_retire")
        return res

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return None

        return InstanceIPAddresses(
            type=resource.type,
            ip_address=resource.ip_address,
            time_to_retire=resource.time_to_retire,
        )


class InstanceIPAddressesArray(object):
    @classmethod
    def to_proto(self, resources):
        if not resources:
            return resources
        return [InstanceIPAddresses.to_proto(i) for i in resources]

    @classmethod
    def from_proto(self, resources):
        return [InstanceIPAddresses.from_proto(i) for i in resources]


class InstanceIPAddressesTimeToRetire(object):
    def __init__(self, seconds: int = None, nanos: int = None):
        self.seconds = seconds
        self.nanos = nanos

    @classmethod
    def to_proto(self, resource):
        if not resource:
            return None

        res = instance_pb2.SqlInstanceIPAddressesTimeToRetire()
        if Primitive.to_proto(resource.seconds):
            res.seconds = Primitive.to_proto(resource.seconds)
        if Primitive.to_proto(resource.nanos):
            res.nanos = Primitive.to_proto(resource.nanos)
        return res

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return None

        return InstanceIPAddressesTimeToRetire(
            seconds=resource.seconds, nanos=resource.nanos,
        )


class InstanceIPAddressesTimeToRetireArray(object):
    @classmethod
    def to_proto(self, resources):
        if not resources:
            return resources
        return [InstanceIPAddressesTimeToRetire.to_proto(i) for i in resources]

    @classmethod
    def from_proto(self, resources):
        return [InstanceIPAddressesTimeToRetire.from_proto(i) for i in resources]


class InstanceMasterInstance(object):
    def __init__(self, name: str = None, region: str = None):
        self.name = name
        self.region = region

    @classmethod
    def to_proto(self, resource):
        if not resource:
            return None

        res = instance_pb2.SqlInstanceMasterInstance()
        if Primitive.to_proto(resource.name):
            res.name = Primitive.to_proto(resource.name)
        if Primitive.to_proto(resource.region):
            res.region = Primitive.to_proto(resource.region)
        return res

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return None

        return InstanceMasterInstance(name=resource.name, region=resource.region,)


class InstanceMasterInstanceArray(object):
    @classmethod
    def to_proto(self, resources):
        if not resources:
            return resources
        return [InstanceMasterInstance.to_proto(i) for i in resources]

    @classmethod
    def from_proto(self, resources):
        return [InstanceMasterInstance.from_proto(i) for i in resources]


class InstanceReplicaConfiguration(object):
    def __init__(
        self,
        kind: str = None,
        mysql_replica_configuration: dict = None,
        failover_target: bool = None,
        replica_pool_configuration: dict = None,
    ):
        self.kind = kind
        self.mysql_replica_configuration = mysql_replica_configuration
        self.failover_target = failover_target
        self.replica_pool_configuration = replica_pool_configuration

    @classmethod
    def to_proto(self, resource):
        if not resource:
            return None

        res = instance_pb2.SqlInstanceReplicaConfiguration()
        if Primitive.to_proto(resource.kind):
            res.kind = Primitive.to_proto(resource.kind)
        if InstanceReplicaConfigurationMysqlReplicaConfiguration.to_proto(
            resource.mysql_replica_configuration
        ):
            res.mysql_replica_configuration.CopyFrom(
                InstanceReplicaConfigurationMysqlReplicaConfiguration.to_proto(
                    resource.mysql_replica_configuration
                )
            )
        else:
            res.ClearField("mysql_replica_configuration")
        if Primitive.to_proto(resource.failover_target):
            res.failover_target = Primitive.to_proto(resource.failover_target)
        if InstanceReplicaConfigurationReplicaPoolConfiguration.to_proto(
            resource.replica_pool_configuration
        ):
            res.replica_pool_configuration.CopyFrom(
                InstanceReplicaConfigurationReplicaPoolConfiguration.to_proto(
                    resource.replica_pool_configuration
                )
            )
        else:
            res.ClearField("replica_pool_configuration")
        return res

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return None

        return InstanceReplicaConfiguration(
            kind=resource.kind,
            mysql_replica_configuration=resource.mysql_replica_configuration,
            failover_target=resource.failover_target,
            replica_pool_configuration=resource.replica_pool_configuration,
        )


class InstanceReplicaConfigurationArray(object):
    @classmethod
    def to_proto(self, resources):
        if not resources:
            return resources
        return [InstanceReplicaConfiguration.to_proto(i) for i in resources]

    @classmethod
    def from_proto(self, resources):
        return [InstanceReplicaConfiguration.from_proto(i) for i in resources]


class InstanceReplicaConfigurationMysqlReplicaConfiguration(object):
    def __init__(
        self,
        dump_file_path: str = None,
        username: str = None,
        password: str = None,
        connect_retry_interval: int = None,
        master_heartbeat_period: dict = None,
        ca_certificate: str = None,
        client_certificate: str = None,
        client_key: str = None,
        ssl_cipher: str = None,
        verify_server_certificate: bool = None,
        kind: str = None,
    ):
        self.dump_file_path = dump_file_path
        self.username = username
        self.password = password
        self.connect_retry_interval = connect_retry_interval
        self.master_heartbeat_period = master_heartbeat_period
        self.ca_certificate = ca_certificate
        self.client_certificate = client_certificate
        self.client_key = client_key
        self.ssl_cipher = ssl_cipher
        self.verify_server_certificate = verify_server_certificate
        self.kind = kind

    @classmethod
    def to_proto(self, resource):
        if not resource:
            return None

        res = instance_pb2.SqlInstanceReplicaConfigurationMysqlReplicaConfiguration()
        if Primitive.to_proto(resource.dump_file_path):
            res.dump_file_path = Primitive.to_proto(resource.dump_file_path)
        if Primitive.to_proto(resource.username):
            res.username = Primitive.to_proto(resource.username)
        if Primitive.to_proto(resource.password):
            res.password = Primitive.to_proto(resource.password)
        if Primitive.to_proto(resource.connect_retry_interval):
            res.connect_retry_interval = Primitive.to_proto(
                resource.connect_retry_interval
            )
        if InstanceReplicaConfigurationMysqlReplicaConfigurationMasterHeartbeatPeriod.to_proto(
            resource.master_heartbeat_period
        ):
            res.master_heartbeat_period.CopyFrom(
                InstanceReplicaConfigurationMysqlReplicaConfigurationMasterHeartbeatPeriod.to_proto(
                    resource.master_heartbeat_period
                )
            )
        else:
            res.ClearField("master_heartbeat_period")
        if Primitive.to_proto(resource.ca_certificate):
            res.ca_certificate = Primitive.to_proto(resource.ca_certificate)
        if Primitive.to_proto(resource.client_certificate):
            res.client_certificate = Primitive.to_proto(resource.client_certificate)
        if Primitive.to_proto(resource.client_key):
            res.client_key = Primitive.to_proto(resource.client_key)
        if Primitive.to_proto(resource.ssl_cipher):
            res.ssl_cipher = Primitive.to_proto(resource.ssl_cipher)
        if Primitive.to_proto(resource.verify_server_certificate):
            res.verify_server_certificate = Primitive.to_proto(
                resource.verify_server_certificate
            )
        if Primitive.to_proto(resource.kind):
            res.kind = Primitive.to_proto(resource.kind)
        return res

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return None

        return InstanceReplicaConfigurationMysqlReplicaConfiguration(
            dump_file_path=resource.dump_file_path,
            username=resource.username,
            password=resource.password,
            connect_retry_interval=resource.connect_retry_interval,
            master_heartbeat_period=resource.master_heartbeat_period,
            ca_certificate=resource.ca_certificate,
            client_certificate=resource.client_certificate,
            client_key=resource.client_key,
            ssl_cipher=resource.ssl_cipher,
            verify_server_certificate=resource.verify_server_certificate,
            kind=resource.kind,
        )


class InstanceReplicaConfigurationMysqlReplicaConfigurationArray(object):
    @classmethod
    def to_proto(self, resources):
        if not resources:
            return resources
        return [
            InstanceReplicaConfigurationMysqlReplicaConfiguration.to_proto(i)
            for i in resources
        ]

    @classmethod
    def from_proto(self, resources):
        return [
            InstanceReplicaConfigurationMysqlReplicaConfiguration.from_proto(i)
            for i in resources
        ]


class InstanceReplicaConfigurationMysqlReplicaConfigurationMasterHeartbeatPeriod(
    object
):
    def __init__(self, value: int = None):
        self.value = value

    @classmethod
    def to_proto(self, resource):
        if not resource:
            return None

        res = (
            instance_pb2.SqlInstanceReplicaConfigurationMysqlReplicaConfigurationMasterHeartbeatPeriod()
        )
        if Primitive.to_proto(resource.value):
            res.value = Primitive.to_proto(resource.value)
        return res

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return None

        return InstanceReplicaConfigurationMysqlReplicaConfigurationMasterHeartbeatPeriod(
            value=resource.value,
        )


class InstanceReplicaConfigurationMysqlReplicaConfigurationMasterHeartbeatPeriodArray(
    object
):
    @classmethod
    def to_proto(self, resources):
        if not resources:
            return resources
        return [
            InstanceReplicaConfigurationMysqlReplicaConfigurationMasterHeartbeatPeriod.to_proto(
                i
            )
            for i in resources
        ]

    @classmethod
    def from_proto(self, resources):
        return [
            InstanceReplicaConfigurationMysqlReplicaConfigurationMasterHeartbeatPeriod.from_proto(
                i
            )
            for i in resources
        ]


class InstanceReplicaConfigurationReplicaPoolConfiguration(object):
    def __init__(
        self,
        kind: str = None,
        static_pool_configuration: dict = None,
        autoscaling_pool_configuration: dict = None,
    ):
        self.kind = kind
        self.static_pool_configuration = static_pool_configuration
        self.autoscaling_pool_configuration = autoscaling_pool_configuration

    @classmethod
    def to_proto(self, resource):
        if not resource:
            return None

        res = instance_pb2.SqlInstanceReplicaConfigurationReplicaPoolConfiguration()
        if Primitive.to_proto(resource.kind):
            res.kind = Primitive.to_proto(resource.kind)
        if InstanceReplicaConfigurationReplicaPoolConfigurationStaticPoolConfiguration.to_proto(
            resource.static_pool_configuration
        ):
            res.static_pool_configuration.CopyFrom(
                InstanceReplicaConfigurationReplicaPoolConfigurationStaticPoolConfiguration.to_proto(
                    resource.static_pool_configuration
                )
            )
        else:
            res.ClearField("static_pool_configuration")
        if InstanceReplicaConfigurationReplicaPoolConfigurationAutoscalingPoolConfiguration.to_proto(
            resource.autoscaling_pool_configuration
        ):
            res.autoscaling_pool_configuration.CopyFrom(
                InstanceReplicaConfigurationReplicaPoolConfigurationAutoscalingPoolConfiguration.to_proto(
                    resource.autoscaling_pool_configuration
                )
            )
        else:
            res.ClearField("autoscaling_pool_configuration")
        return res

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return None

        return InstanceReplicaConfigurationReplicaPoolConfiguration(
            kind=resource.kind,
            static_pool_configuration=resource.static_pool_configuration,
            autoscaling_pool_configuration=resource.autoscaling_pool_configuration,
        )


class InstanceReplicaConfigurationReplicaPoolConfigurationArray(object):
    @classmethod
    def to_proto(self, resources):
        if not resources:
            return resources
        return [
            InstanceReplicaConfigurationReplicaPoolConfiguration.to_proto(i)
            for i in resources
        ]

    @classmethod
    def from_proto(self, resources):
        return [
            InstanceReplicaConfigurationReplicaPoolConfiguration.from_proto(i)
            for i in resources
        ]


class InstanceReplicaConfigurationReplicaPoolConfigurationStaticPoolConfiguration(
    object
):
    def __init__(
        self,
        kind: str = None,
        replica_count: int = None,
        expose_replica_ip: bool = None,
    ):
        self.kind = kind
        self.replica_count = replica_count
        self.expose_replica_ip = expose_replica_ip

    @classmethod
    def to_proto(self, resource):
        if not resource:
            return None

        res = (
            instance_pb2.SqlInstanceReplicaConfigurationReplicaPoolConfigurationStaticPoolConfiguration()
        )
        if Primitive.to_proto(resource.kind):
            res.kind = Primitive.to_proto(resource.kind)
        if Primitive.to_proto(resource.replica_count):
            res.replica_count = Primitive.to_proto(resource.replica_count)
        if Primitive.to_proto(resource.expose_replica_ip):
            res.expose_replica_ip = Primitive.to_proto(resource.expose_replica_ip)
        return res

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return None

        return InstanceReplicaConfigurationReplicaPoolConfigurationStaticPoolConfiguration(
            kind=resource.kind,
            replica_count=resource.replica_count,
            expose_replica_ip=resource.expose_replica_ip,
        )


class InstanceReplicaConfigurationReplicaPoolConfigurationStaticPoolConfigurationArray(
    object
):
    @classmethod
    def to_proto(self, resources):
        if not resources:
            return resources
        return [
            InstanceReplicaConfigurationReplicaPoolConfigurationStaticPoolConfiguration.to_proto(
                i
            )
            for i in resources
        ]

    @classmethod
    def from_proto(self, resources):
        return [
            InstanceReplicaConfigurationReplicaPoolConfigurationStaticPoolConfiguration.from_proto(
                i
            )
            for i in resources
        ]


class InstanceReplicaConfigurationReplicaPoolConfigurationAutoscalingPoolConfiguration(
    object
):
    def __init__(
        self,
        kind: str = None,
        min_replica_count: int = None,
        max_replica_count: int = None,
        target_cpu_util: float = None,
    ):
        self.kind = kind
        self.min_replica_count = min_replica_count
        self.max_replica_count = max_replica_count
        self.target_cpu_util = target_cpu_util

    @classmethod
    def to_proto(self, resource):
        if not resource:
            return None

        res = (
            instance_pb2.SqlInstanceReplicaConfigurationReplicaPoolConfigurationAutoscalingPoolConfiguration()
        )
        if Primitive.to_proto(resource.kind):
            res.kind = Primitive.to_proto(resource.kind)
        if Primitive.to_proto(resource.min_replica_count):
            res.min_replica_count = Primitive.to_proto(resource.min_replica_count)
        if Primitive.to_proto(resource.max_replica_count):
            res.max_replica_count = Primitive.to_proto(resource.max_replica_count)
        if Primitive.to_proto(resource.target_cpu_util):
            res.target_cpu_util = Primitive.to_proto(resource.target_cpu_util)
        return res

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return None

        return InstanceReplicaConfigurationReplicaPoolConfigurationAutoscalingPoolConfiguration(
            kind=resource.kind,
            min_replica_count=resource.min_replica_count,
            max_replica_count=resource.max_replica_count,
            target_cpu_util=resource.target_cpu_util,
        )


class InstanceReplicaConfigurationReplicaPoolConfigurationAutoscalingPoolConfigurationArray(
    object
):
    @classmethod
    def to_proto(self, resources):
        if not resources:
            return resources
        return [
            InstanceReplicaConfigurationReplicaPoolConfigurationAutoscalingPoolConfiguration.to_proto(
                i
            )
            for i in resources
        ]

    @classmethod
    def from_proto(self, resources):
        return [
            InstanceReplicaConfigurationReplicaPoolConfigurationAutoscalingPoolConfiguration.from_proto(
                i
            )
            for i in resources
        ]


class InstanceScheduledMaintenance(object):
    def __init__(
        self,
        start_time: dict = None,
        can_defer: bool = None,
        can_reschedule: bool = None,
    ):
        self.start_time = start_time
        self.can_defer = can_defer
        self.can_reschedule = can_reschedule

    @classmethod
    def to_proto(self, resource):
        if not resource:
            return None

        res = instance_pb2.SqlInstanceScheduledMaintenance()
        if InstanceScheduledMaintenanceStartTime.to_proto(resource.start_time):
            res.start_time.CopyFrom(
                InstanceScheduledMaintenanceStartTime.to_proto(resource.start_time)
            )
        else:
            res.ClearField("start_time")
        if Primitive.to_proto(resource.can_defer):
            res.can_defer = Primitive.to_proto(resource.can_defer)
        if Primitive.to_proto(resource.can_reschedule):
            res.can_reschedule = Primitive.to_proto(resource.can_reschedule)
        return res

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return None

        return InstanceScheduledMaintenance(
            start_time=resource.start_time,
            can_defer=resource.can_defer,
            can_reschedule=resource.can_reschedule,
        )


class InstanceScheduledMaintenanceArray(object):
    @classmethod
    def to_proto(self, resources):
        if not resources:
            return resources
        return [InstanceScheduledMaintenance.to_proto(i) for i in resources]

    @classmethod
    def from_proto(self, resources):
        return [InstanceScheduledMaintenance.from_proto(i) for i in resources]


class InstanceScheduledMaintenanceStartTime(object):
    def __init__(self, seconds: int = None, nanos: int = None):
        self.seconds = seconds
        self.nanos = nanos

    @classmethod
    def to_proto(self, resource):
        if not resource:
            return None

        res = instance_pb2.SqlInstanceScheduledMaintenanceStartTime()
        if Primitive.to_proto(resource.seconds):
            res.seconds = Primitive.to_proto(resource.seconds)
        if Primitive.to_proto(resource.nanos):
            res.nanos = Primitive.to_proto(resource.nanos)
        return res

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return None

        return InstanceScheduledMaintenanceStartTime(
            seconds=resource.seconds, nanos=resource.nanos,
        )


class InstanceScheduledMaintenanceStartTimeArray(object):
    @classmethod
    def to_proto(self, resources):
        if not resources:
            return resources
        return [InstanceScheduledMaintenanceStartTime.to_proto(i) for i in resources]

    @classmethod
    def from_proto(self, resources):
        return [InstanceScheduledMaintenanceStartTime.from_proto(i) for i in resources]


class InstanceSettings(object):
    def __init__(
        self,
        authorized_gae_applications: list = None,
        tier: str = None,
        kind: str = None,
        availability_type: str = None,
        pricing_plan: str = None,
        replication_type: str = None,
        activation_policy: str = None,
        storage_auto_resize: bool = None,
        data_disk_type: str = None,
        database_replication_enabled: bool = None,
        crash_safe_replication_enabled: bool = None,
    ):
        self.authorized_gae_applications = authorized_gae_applications
        self.tier = tier
        self.kind = kind
        self.availability_type = availability_type
        self.pricing_plan = pricing_plan
        self.replication_type = replication_type
        self.activation_policy = activation_policy
        self.storage_auto_resize = storage_auto_resize
        self.data_disk_type = data_disk_type
        self.database_replication_enabled = database_replication_enabled
        self.crash_safe_replication_enabled = crash_safe_replication_enabled

    @classmethod
    def to_proto(self, resource):
        if not resource:
            return None

        res = instance_pb2.SqlInstanceSettings()
        if Primitive.to_proto(resource.authorized_gae_applications):
            res.authorized_gae_applications.extend(
                Primitive.to_proto(resource.authorized_gae_applications)
            )
        if Primitive.to_proto(resource.tier):
            res.tier = Primitive.to_proto(resource.tier)
        if Primitive.to_proto(resource.kind):
            res.kind = Primitive.to_proto(resource.kind)
        if InstanceSettingsAvailabilityTypeEnum.to_proto(resource.availability_type):
            res.availability_type = InstanceSettingsAvailabilityTypeEnum.to_proto(
                resource.availability_type
            )
        if InstanceSettingsPricingPlanEnum.to_proto(resource.pricing_plan):
            res.pricing_plan = InstanceSettingsPricingPlanEnum.to_proto(
                resource.pricing_plan
            )
        if InstanceSettingsReplicationTypeEnum.to_proto(resource.replication_type):
            res.replication_type = InstanceSettingsReplicationTypeEnum.to_proto(
                resource.replication_type
            )
        if InstanceSettingsActivationPolicyEnum.to_proto(resource.activation_policy):
            res.activation_policy = InstanceSettingsActivationPolicyEnum.to_proto(
                resource.activation_policy
            )
        if Primitive.to_proto(resource.storage_auto_resize):
            res.storage_auto_resize = Primitive.to_proto(resource.storage_auto_resize)
        if InstanceSettingsDataDiskTypeEnum.to_proto(resource.data_disk_type):
            res.data_disk_type = InstanceSettingsDataDiskTypeEnum.to_proto(
                resource.data_disk_type
            )
        if Primitive.to_proto(resource.database_replication_enabled):
            res.database_replication_enabled = Primitive.to_proto(
                resource.database_replication_enabled
            )
        if Primitive.to_proto(resource.crash_safe_replication_enabled):
            res.crash_safe_replication_enabled = Primitive.to_proto(
                resource.crash_safe_replication_enabled
            )
        return res

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return None

        return InstanceSettings(
            authorized_gae_applications=resource.authorized_gae_applications,
            tier=resource.tier,
            kind=resource.kind,
            availability_type=resource.availability_type,
            pricing_plan=resource.pricing_plan,
            replication_type=resource.replication_type,
            activation_policy=resource.activation_policy,
            storage_auto_resize=resource.storage_auto_resize,
            data_disk_type=resource.data_disk_type,
            database_replication_enabled=resource.database_replication_enabled,
            crash_safe_replication_enabled=resource.crash_safe_replication_enabled,
        )


class InstanceSettingsArray(object):
    @classmethod
    def to_proto(self, resources):
        if not resources:
            return resources
        return [InstanceSettings.to_proto(i) for i in resources]

    @classmethod
    def from_proto(self, resources):
        return [InstanceSettings.from_proto(i) for i in resources]


class InstanceBackendTypeEnum(object):
    @classmethod
    def to_proto(self, resource):
        if not resource:
            return resource
        return instance_pb2.SqlInstanceBackendTypeEnum.Value(
            "InstanceBackendTypeEnum%s" % resource
        )

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return resource
        return instance_pb2.SqlInstanceBackendTypeEnum.Name(resource)[
            len("InstanceBackendTypeEnum") :
        ]


class InstanceDatabaseVersionEnum(object):
    @classmethod
    def to_proto(self, resource):
        if not resource:
            return resource
        return instance_pb2.SqlInstanceDatabaseVersionEnum.Value(
            "InstanceDatabaseVersionEnum%s" % resource
        )

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return resource
        return instance_pb2.SqlInstanceDatabaseVersionEnum.Name(resource)[
            len("InstanceDatabaseVersionEnum") :
        ]


class InstanceInstanceTypeEnum(object):
    @classmethod
    def to_proto(self, resource):
        if not resource:
            return resource
        return instance_pb2.SqlInstanceInstanceTypeEnum.Value(
            "InstanceInstanceTypeEnum%s" % resource
        )

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return resource
        return instance_pb2.SqlInstanceInstanceTypeEnum.Name(resource)[
            len("InstanceInstanceTypeEnum") :
        ]


class InstanceIPAddressesTypeEnum(object):
    @classmethod
    def to_proto(self, resource):
        if not resource:
            return resource
        return instance_pb2.SqlInstanceIPAddressesTypeEnum.Value(
            "InstanceIPAddressesTypeEnum%s" % resource
        )

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return resource
        return instance_pb2.SqlInstanceIPAddressesTypeEnum.Name(resource)[
            len("InstanceIPAddressesTypeEnum") :
        ]


class InstanceSettingsAvailabilityTypeEnum(object):
    @classmethod
    def to_proto(self, resource):
        if not resource:
            return resource
        return instance_pb2.SqlInstanceSettingsAvailabilityTypeEnum.Value(
            "InstanceSettingsAvailabilityTypeEnum%s" % resource
        )

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return resource
        return instance_pb2.SqlInstanceSettingsAvailabilityTypeEnum.Name(resource)[
            len("InstanceSettingsAvailabilityTypeEnum") :
        ]


class InstanceSettingsPricingPlanEnum(object):
    @classmethod
    def to_proto(self, resource):
        if not resource:
            return resource
        return instance_pb2.SqlInstanceSettingsPricingPlanEnum.Value(
            "InstanceSettingsPricingPlanEnum%s" % resource
        )

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return resource
        return instance_pb2.SqlInstanceSettingsPricingPlanEnum.Name(resource)[
            len("InstanceSettingsPricingPlanEnum") :
        ]


class InstanceSettingsReplicationTypeEnum(object):
    @classmethod
    def to_proto(self, resource):
        if not resource:
            return resource
        return instance_pb2.SqlInstanceSettingsReplicationTypeEnum.Value(
            "InstanceSettingsReplicationTypeEnum%s" % resource
        )

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return resource
        return instance_pb2.SqlInstanceSettingsReplicationTypeEnum.Name(resource)[
            len("InstanceSettingsReplicationTypeEnum") :
        ]


class InstanceSettingsActivationPolicyEnum(object):
    @classmethod
    def to_proto(self, resource):
        if not resource:
            return resource
        return instance_pb2.SqlInstanceSettingsActivationPolicyEnum.Value(
            "InstanceSettingsActivationPolicyEnum%s" % resource
        )

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return resource
        return instance_pb2.SqlInstanceSettingsActivationPolicyEnum.Name(resource)[
            len("InstanceSettingsActivationPolicyEnum") :
        ]


class InstanceSettingsDataDiskTypeEnum(object):
    @classmethod
    def to_proto(self, resource):
        if not resource:
            return resource
        return instance_pb2.SqlInstanceSettingsDataDiskTypeEnum.Value(
            "InstanceSettingsDataDiskTypeEnum%s" % resource
        )

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return resource
        return instance_pb2.SqlInstanceSettingsDataDiskTypeEnum.Name(resource)[
            len("InstanceSettingsDataDiskTypeEnum") :
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
