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
        state: str = None,
        replica_instances: list = None,
        server_ca_cert: dict = None,
        ipv6_address: str = None,
        service_account_email_address: str = None,
        on_premises_configuration: dict = None,
        suspension_reason: list = None,
        disk_encryption_status: dict = None,
        instance_uid: str = None,
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
        self.state = state
        self.replica_instances = replica_instances
        self.server_ca_cert = server_ca_cert
        self.ipv6_address = ipv6_address
        self.service_account_email_address = service_account_email_address
        self.on_premises_configuration = on_premises_configuration
        self.suspension_reason = suspension_reason
        self.disk_encryption_status = disk_encryption_status
        self.instance_uid = instance_uid
        self.service_account_file = service_account_file

    def apply(self):
        stub = instance_pb2_grpc.SqlBetaInstanceServiceStub(channel.Channel())
        request = instance_pb2.ApplySqlBetaInstanceRequest()
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
        if Primitive.to_proto(self.state):
            request.resource.state = Primitive.to_proto(self.state)

        if InstanceReplicaInstancesArray.to_proto(self.replica_instances):
            request.resource.replica_instances.extend(
                InstanceReplicaInstancesArray.to_proto(self.replica_instances)
            )
        if InstanceServerCaCert.to_proto(self.server_ca_cert):
            request.resource.server_ca_cert.CopyFrom(
                InstanceServerCaCert.to_proto(self.server_ca_cert)
            )
        else:
            request.resource.ClearField("server_ca_cert")
        if Primitive.to_proto(self.ipv6_address):
            request.resource.ipv6_address = Primitive.to_proto(self.ipv6_address)

        if Primitive.to_proto(self.service_account_email_address):
            request.resource.service_account_email_address = Primitive.to_proto(
                self.service_account_email_address
            )

        if InstanceOnPremisesConfiguration.to_proto(self.on_premises_configuration):
            request.resource.on_premises_configuration.CopyFrom(
                InstanceOnPremisesConfiguration.to_proto(self.on_premises_configuration)
            )
        else:
            request.resource.ClearField("on_premises_configuration")
        if Primitive.to_proto(self.suspension_reason):
            request.resource.suspension_reason.extend(
                Primitive.to_proto(self.suspension_reason)
            )
        if InstanceDiskEncryptionStatus.to_proto(self.disk_encryption_status):
            request.resource.disk_encryption_status.CopyFrom(
                InstanceDiskEncryptionStatus.to_proto(self.disk_encryption_status)
            )
        else:
            request.resource.ClearField("disk_encryption_status")
        if Primitive.to_proto(self.instance_uid):
            request.resource.instance_uid = Primitive.to_proto(self.instance_uid)

        request.service_account_file = self.service_account_file

        response = stub.ApplySqlBetaInstance(request)
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
        self.state = Primitive.from_proto(response.state)
        self.replica_instances = InstanceReplicaInstancesArray.from_proto(
            response.replica_instances
        )
        self.server_ca_cert = InstanceServerCaCert.from_proto(response.server_ca_cert)
        self.ipv6_address = Primitive.from_proto(response.ipv6_address)
        self.service_account_email_address = Primitive.from_proto(
            response.service_account_email_address
        )
        self.on_premises_configuration = InstanceOnPremisesConfiguration.from_proto(
            response.on_premises_configuration
        )
        self.suspension_reason = Primitive.from_proto(response.suspension_reason)
        self.disk_encryption_status = InstanceDiskEncryptionStatus.from_proto(
            response.disk_encryption_status
        )
        self.instance_uid = Primitive.from_proto(response.instance_uid)

    def delete(self):
        stub = instance_pb2_grpc.SqlBetaInstanceServiceStub(channel.Channel())
        request = instance_pb2.DeleteSqlBetaInstanceRequest()
        request.service_account_file = self.service_account_file
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
        if Primitive.to_proto(self.state):
            request.resource.state = Primitive.to_proto(self.state)

        if InstanceReplicaInstancesArray.to_proto(self.replica_instances):
            request.resource.replica_instances.extend(
                InstanceReplicaInstancesArray.to_proto(self.replica_instances)
            )
        if InstanceServerCaCert.to_proto(self.server_ca_cert):
            request.resource.server_ca_cert.CopyFrom(
                InstanceServerCaCert.to_proto(self.server_ca_cert)
            )
        else:
            request.resource.ClearField("server_ca_cert")
        if Primitive.to_proto(self.ipv6_address):
            request.resource.ipv6_address = Primitive.to_proto(self.ipv6_address)

        if Primitive.to_proto(self.service_account_email_address):
            request.resource.service_account_email_address = Primitive.to_proto(
                self.service_account_email_address
            )

        if InstanceOnPremisesConfiguration.to_proto(self.on_premises_configuration):
            request.resource.on_premises_configuration.CopyFrom(
                InstanceOnPremisesConfiguration.to_proto(self.on_premises_configuration)
            )
        else:
            request.resource.ClearField("on_premises_configuration")
        if Primitive.to_proto(self.suspension_reason):
            request.resource.suspension_reason.extend(
                Primitive.to_proto(self.suspension_reason)
            )
        if InstanceDiskEncryptionStatus.to_proto(self.disk_encryption_status):
            request.resource.disk_encryption_status.CopyFrom(
                InstanceDiskEncryptionStatus.to_proto(self.disk_encryption_status)
            )
        else:
            request.resource.ClearField("disk_encryption_status")
        if Primitive.to_proto(self.instance_uid):
            request.resource.instance_uid = Primitive.to_proto(self.instance_uid)

        response = stub.DeleteSqlBetaInstance(request)

    @classmethod
    def list(self, project, service_account_file=""):
        stub = instance_pb2_grpc.SqlBetaInstanceServiceStub(channel.Channel())
        request = instance_pb2.ListSqlBetaInstanceRequest()
        request.service_account_file = service_account_file
        request.Project = project

        return stub.ListSqlBetaInstance(request).items

    def to_proto(self):
        resource = instance_pb2.SqlBetaInstance()
        if InstanceBackendTypeEnum.to_proto(self.backend_type):
            resource.backend_type = InstanceBackendTypeEnum.to_proto(self.backend_type)
        if Primitive.to_proto(self.connection_name):
            resource.connection_name = Primitive.to_proto(self.connection_name)
        if InstanceDatabaseVersionEnum.to_proto(self.database_version):
            resource.database_version = InstanceDatabaseVersionEnum.to_proto(
                self.database_version
            )
        if Primitive.to_proto(self.etag):
            resource.etag = Primitive.to_proto(self.etag)
        if Primitive.to_proto(self.gce_zone):
            resource.gce_zone = Primitive.to_proto(self.gce_zone)
        if InstanceInstanceTypeEnum.to_proto(self.instance_type):
            resource.instance_type = InstanceInstanceTypeEnum.to_proto(
                self.instance_type
            )
        if Primitive.to_proto(self.master_instance_name):
            resource.master_instance_name = Primitive.to_proto(
                self.master_instance_name
            )
        if InstanceMaxDiskSize.to_proto(self.max_disk_size):
            resource.max_disk_size.CopyFrom(
                InstanceMaxDiskSize.to_proto(self.max_disk_size)
            )
        else:
            resource.ClearField("max_disk_size")
        if Primitive.to_proto(self.name):
            resource.name = Primitive.to_proto(self.name)
        if Primitive.to_proto(self.project):
            resource.project = Primitive.to_proto(self.project)
        if Primitive.to_proto(self.region):
            resource.region = Primitive.to_proto(self.region)
        if Primitive.to_proto(self.root_password):
            resource.root_password = Primitive.to_proto(self.root_password)
        if InstanceCurrentDiskSize.to_proto(self.current_disk_size):
            resource.current_disk_size.CopyFrom(
                InstanceCurrentDiskSize.to_proto(self.current_disk_size)
            )
        else:
            resource.ClearField("current_disk_size")
        if InstanceDiskEncryptionConfiguration.to_proto(
            self.disk_encryption_configuration
        ):
            resource.disk_encryption_configuration.CopyFrom(
                InstanceDiskEncryptionConfiguration.to_proto(
                    self.disk_encryption_configuration
                )
            )
        else:
            resource.ClearField("disk_encryption_configuration")
        if InstanceFailoverReplica.to_proto(self.failover_replica):
            resource.failover_replica.CopyFrom(
                InstanceFailoverReplica.to_proto(self.failover_replica)
            )
        else:
            resource.ClearField("failover_replica")
        if InstanceIPAddressesArray.to_proto(self.ip_addresses):
            resource.ip_addresses.extend(
                InstanceIPAddressesArray.to_proto(self.ip_addresses)
            )
        if InstanceMasterInstance.to_proto(self.master_instance):
            resource.master_instance.CopyFrom(
                InstanceMasterInstance.to_proto(self.master_instance)
            )
        else:
            resource.ClearField("master_instance")
        if InstanceReplicaConfiguration.to_proto(self.replica_configuration):
            resource.replica_configuration.CopyFrom(
                InstanceReplicaConfiguration.to_proto(self.replica_configuration)
            )
        else:
            resource.ClearField("replica_configuration")
        if InstanceScheduledMaintenance.to_proto(self.scheduled_maintenance):
            resource.scheduled_maintenance.CopyFrom(
                InstanceScheduledMaintenance.to_proto(self.scheduled_maintenance)
            )
        else:
            resource.ClearField("scheduled_maintenance")
        if InstanceSettings.to_proto(self.settings):
            resource.settings.CopyFrom(InstanceSettings.to_proto(self.settings))
        else:
            resource.ClearField("settings")
        if Primitive.to_proto(self.state):
            resource.state = Primitive.to_proto(self.state)
        if InstanceReplicaInstancesArray.to_proto(self.replica_instances):
            resource.replica_instances.extend(
                InstanceReplicaInstancesArray.to_proto(self.replica_instances)
            )
        if InstanceServerCaCert.to_proto(self.server_ca_cert):
            resource.server_ca_cert.CopyFrom(
                InstanceServerCaCert.to_proto(self.server_ca_cert)
            )
        else:
            resource.ClearField("server_ca_cert")
        if Primitive.to_proto(self.ipv6_address):
            resource.ipv6_address = Primitive.to_proto(self.ipv6_address)
        if Primitive.to_proto(self.service_account_email_address):
            resource.service_account_email_address = Primitive.to_proto(
                self.service_account_email_address
            )
        if InstanceOnPremisesConfiguration.to_proto(self.on_premises_configuration):
            resource.on_premises_configuration.CopyFrom(
                InstanceOnPremisesConfiguration.to_proto(self.on_premises_configuration)
            )
        else:
            resource.ClearField("on_premises_configuration")
        if Primitive.to_proto(self.suspension_reason):
            resource.suspension_reason.extend(
                Primitive.to_proto(self.suspension_reason)
            )
        if InstanceDiskEncryptionStatus.to_proto(self.disk_encryption_status):
            resource.disk_encryption_status.CopyFrom(
                InstanceDiskEncryptionStatus.to_proto(self.disk_encryption_status)
            )
        else:
            resource.ClearField("disk_encryption_status")
        if Primitive.to_proto(self.instance_uid):
            resource.instance_uid = Primitive.to_proto(self.instance_uid)
        return resource


class InstanceMaxDiskSize(object):
    def __init__(self, value: int = None):
        self.value = value

    @classmethod
    def to_proto(self, resource):
        if not resource:
            return None

        res = instance_pb2.SqlBetaInstanceMaxDiskSize()
        if Primitive.to_proto(resource.value):
            res.value = Primitive.to_proto(resource.value)
        return res

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return None

        return InstanceMaxDiskSize(value=Primitive.from_proto(resource.value),)


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

        res = instance_pb2.SqlBetaInstanceCurrentDiskSize()
        if Primitive.to_proto(resource.value):
            res.value = Primitive.to_proto(resource.value)
        return res

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return None

        return InstanceCurrentDiskSize(value=Primitive.from_proto(resource.value),)


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

        res = instance_pb2.SqlBetaInstanceDiskEncryptionConfiguration()
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
            kms_key_name=Primitive.from_proto(resource.kms_key_name),
            kind=Primitive.from_proto(resource.kind),
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
    def __init__(
        self, name: str = None, available: bool = None, failover_instance: dict = None
    ):
        self.name = name
        self.available = available
        self.failover_instance = failover_instance

    @classmethod
    def to_proto(self, resource):
        if not resource:
            return None

        res = instance_pb2.SqlBetaInstanceFailoverReplica()
        if Primitive.to_proto(resource.name):
            res.name = Primitive.to_proto(resource.name)
        if Primitive.to_proto(resource.available):
            res.available = Primitive.to_proto(resource.available)
        if InstanceFailoverReplicaFailoverInstance.to_proto(resource.failover_instance):
            res.failover_instance.CopyFrom(
                InstanceFailoverReplicaFailoverInstance.to_proto(
                    resource.failover_instance
                )
            )
        else:
            res.ClearField("failover_instance")
        return res

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return None

        return InstanceFailoverReplica(
            name=Primitive.from_proto(resource.name),
            available=Primitive.from_proto(resource.available),
            failover_instance=InstanceFailoverReplicaFailoverInstance.from_proto(
                resource.failover_instance
            ),
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


class InstanceFailoverReplicaFailoverInstance(object):
    def __init__(self, name: str = None, region: str = None):
        self.name = name
        self.region = region

    @classmethod
    def to_proto(self, resource):
        if not resource:
            return None

        res = instance_pb2.SqlBetaInstanceFailoverReplicaFailoverInstance()
        if Primitive.to_proto(resource.name):
            res.name = Primitive.to_proto(resource.name)
        if Primitive.to_proto(resource.region):
            res.region = Primitive.to_proto(resource.region)
        return res

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return None

        return InstanceFailoverReplicaFailoverInstance(
            name=Primitive.from_proto(resource.name),
            region=Primitive.from_proto(resource.region),
        )


class InstanceFailoverReplicaFailoverInstanceArray(object):
    @classmethod
    def to_proto(self, resources):
        if not resources:
            return resources
        return [InstanceFailoverReplicaFailoverInstance.to_proto(i) for i in resources]

    @classmethod
    def from_proto(self, resources):
        return [
            InstanceFailoverReplicaFailoverInstance.from_proto(i) for i in resources
        ]


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

        res = instance_pb2.SqlBetaInstanceIPAddresses()
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
            type=InstanceIPAddressesTypeEnum.from_proto(resource.type),
            ip_address=Primitive.from_proto(resource.ip_address),
            time_to_retire=InstanceIPAddressesTimeToRetire.from_proto(
                resource.time_to_retire
            ),
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

        res = instance_pb2.SqlBetaInstanceIPAddressesTimeToRetire()
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
            seconds=Primitive.from_proto(resource.seconds),
            nanos=Primitive.from_proto(resource.nanos),
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

        res = instance_pb2.SqlBetaInstanceMasterInstance()
        if Primitive.to_proto(resource.name):
            res.name = Primitive.to_proto(resource.name)
        if Primitive.to_proto(resource.region):
            res.region = Primitive.to_proto(resource.region)
        return res

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return None

        return InstanceMasterInstance(
            name=Primitive.from_proto(resource.name),
            region=Primitive.from_proto(resource.region),
        )


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

        res = instance_pb2.SqlBetaInstanceReplicaConfiguration()
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
            kind=Primitive.from_proto(resource.kind),
            mysql_replica_configuration=InstanceReplicaConfigurationMysqlReplicaConfiguration.from_proto(
                resource.mysql_replica_configuration
            ),
            failover_target=Primitive.from_proto(resource.failover_target),
            replica_pool_configuration=InstanceReplicaConfigurationReplicaPoolConfiguration.from_proto(
                resource.replica_pool_configuration
            ),
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

        res = (
            instance_pb2.SqlBetaInstanceReplicaConfigurationMysqlReplicaConfiguration()
        )
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
            dump_file_path=Primitive.from_proto(resource.dump_file_path),
            username=Primitive.from_proto(resource.username),
            password=Primitive.from_proto(resource.password),
            connect_retry_interval=Primitive.from_proto(
                resource.connect_retry_interval
            ),
            master_heartbeat_period=InstanceReplicaConfigurationMysqlReplicaConfigurationMasterHeartbeatPeriod.from_proto(
                resource.master_heartbeat_period
            ),
            ca_certificate=Primitive.from_proto(resource.ca_certificate),
            client_certificate=Primitive.from_proto(resource.client_certificate),
            client_key=Primitive.from_proto(resource.client_key),
            ssl_cipher=Primitive.from_proto(resource.ssl_cipher),
            verify_server_certificate=Primitive.from_proto(
                resource.verify_server_certificate
            ),
            kind=Primitive.from_proto(resource.kind),
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
            instance_pb2.SqlBetaInstanceReplicaConfigurationMysqlReplicaConfigurationMasterHeartbeatPeriod()
        )
        if Primitive.to_proto(resource.value):
            res.value = Primitive.to_proto(resource.value)
        return res

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return None

        return InstanceReplicaConfigurationMysqlReplicaConfigurationMasterHeartbeatPeriod(
            value=Primitive.from_proto(resource.value),
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
        replica_count: int = None,
        expose_replica_ip: bool = None,
    ):
        self.kind = kind
        self.static_pool_configuration = static_pool_configuration
        self.autoscaling_pool_configuration = autoscaling_pool_configuration
        self.replica_count = replica_count
        self.expose_replica_ip = expose_replica_ip

    @classmethod
    def to_proto(self, resource):
        if not resource:
            return None

        res = instance_pb2.SqlBetaInstanceReplicaConfigurationReplicaPoolConfiguration()
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
        if Primitive.to_proto(resource.replica_count):
            res.replica_count = Primitive.to_proto(resource.replica_count)
        if Primitive.to_proto(resource.expose_replica_ip):
            res.expose_replica_ip = Primitive.to_proto(resource.expose_replica_ip)
        return res

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return None

        return InstanceReplicaConfigurationReplicaPoolConfiguration(
            kind=Primitive.from_proto(resource.kind),
            static_pool_configuration=InstanceReplicaConfigurationReplicaPoolConfigurationStaticPoolConfiguration.from_proto(
                resource.static_pool_configuration
            ),
            autoscaling_pool_configuration=InstanceReplicaConfigurationReplicaPoolConfigurationAutoscalingPoolConfiguration.from_proto(
                resource.autoscaling_pool_configuration
            ),
            replica_count=Primitive.from_proto(resource.replica_count),
            expose_replica_ip=Primitive.from_proto(resource.expose_replica_ip),
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
            instance_pb2.SqlBetaInstanceReplicaConfigurationReplicaPoolConfigurationStaticPoolConfiguration()
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
            kind=Primitive.from_proto(resource.kind),
            replica_count=Primitive.from_proto(resource.replica_count),
            expose_replica_ip=Primitive.from_proto(resource.expose_replica_ip),
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
            instance_pb2.SqlBetaInstanceReplicaConfigurationReplicaPoolConfigurationAutoscalingPoolConfiguration()
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
            kind=Primitive.from_proto(resource.kind),
            min_replica_count=Primitive.from_proto(resource.min_replica_count),
            max_replica_count=Primitive.from_proto(resource.max_replica_count),
            target_cpu_util=Primitive.from_proto(resource.target_cpu_util),
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

        res = instance_pb2.SqlBetaInstanceScheduledMaintenance()
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
            start_time=InstanceScheduledMaintenanceStartTime.from_proto(
                resource.start_time
            ),
            can_defer=Primitive.from_proto(resource.can_defer),
            can_reschedule=Primitive.from_proto(resource.can_reschedule),
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

        res = instance_pb2.SqlBetaInstanceScheduledMaintenanceStartTime()
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
            seconds=Primitive.from_proto(resource.seconds),
            nanos=Primitive.from_proto(resource.nanos),
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
        settings_version: dict = None,
        user_labels: dict = None,
        storage_auto_resize_limit: dict = None,
        ip_configuration: dict = None,
        location_preference: dict = None,
        database_flags: list = None,
        maintenance_window: dict = None,
        backup_configuration: dict = None,
        data_disk_size_gb: dict = None,
        active_directory_config: dict = None,
        collation: str = None,
        deny_maintenance_periods: list = None,
        insights_config: dict = None,
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
        self.settings_version = settings_version
        self.user_labels = user_labels
        self.storage_auto_resize_limit = storage_auto_resize_limit
        self.ip_configuration = ip_configuration
        self.location_preference = location_preference
        self.database_flags = database_flags
        self.maintenance_window = maintenance_window
        self.backup_configuration = backup_configuration
        self.data_disk_size_gb = data_disk_size_gb
        self.active_directory_config = active_directory_config
        self.collation = collation
        self.deny_maintenance_periods = deny_maintenance_periods
        self.insights_config = insights_config

    @classmethod
    def to_proto(self, resource):
        if not resource:
            return None

        res = instance_pb2.SqlBetaInstanceSettings()
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
        if InstanceSettingsSettingsVersion.to_proto(resource.settings_version):
            res.settings_version.CopyFrom(
                InstanceSettingsSettingsVersion.to_proto(resource.settings_version)
            )
        else:
            res.ClearField("settings_version")
        if Primitive.to_proto(resource.user_labels):
            res.user_labels = Primitive.to_proto(resource.user_labels)
        if InstanceSettingsStorageAutoResizeLimit.to_proto(
            resource.storage_auto_resize_limit
        ):
            res.storage_auto_resize_limit.CopyFrom(
                InstanceSettingsStorageAutoResizeLimit.to_proto(
                    resource.storage_auto_resize_limit
                )
            )
        else:
            res.ClearField("storage_auto_resize_limit")
        if InstanceSettingsIPConfiguration.to_proto(resource.ip_configuration):
            res.ip_configuration.CopyFrom(
                InstanceSettingsIPConfiguration.to_proto(resource.ip_configuration)
            )
        else:
            res.ClearField("ip_configuration")
        if InstanceSettingsLocationPreference.to_proto(resource.location_preference):
            res.location_preference.CopyFrom(
                InstanceSettingsLocationPreference.to_proto(
                    resource.location_preference
                )
            )
        else:
            res.ClearField("location_preference")
        if InstanceSettingsDatabaseFlagsArray.to_proto(resource.database_flags):
            res.database_flags.extend(
                InstanceSettingsDatabaseFlagsArray.to_proto(resource.database_flags)
            )
        if InstanceSettingsMaintenanceWindow.to_proto(resource.maintenance_window):
            res.maintenance_window.CopyFrom(
                InstanceSettingsMaintenanceWindow.to_proto(resource.maintenance_window)
            )
        else:
            res.ClearField("maintenance_window")
        if InstanceSettingsBackupConfiguration.to_proto(resource.backup_configuration):
            res.backup_configuration.CopyFrom(
                InstanceSettingsBackupConfiguration.to_proto(
                    resource.backup_configuration
                )
            )
        else:
            res.ClearField("backup_configuration")
        if InstanceSettingsDataDiskSizeGb.to_proto(resource.data_disk_size_gb):
            res.data_disk_size_gb.CopyFrom(
                InstanceSettingsDataDiskSizeGb.to_proto(resource.data_disk_size_gb)
            )
        else:
            res.ClearField("data_disk_size_gb")
        if InstanceSettingsActiveDirectoryConfig.to_proto(
            resource.active_directory_config
        ):
            res.active_directory_config.CopyFrom(
                InstanceSettingsActiveDirectoryConfig.to_proto(
                    resource.active_directory_config
                )
            )
        else:
            res.ClearField("active_directory_config")
        if Primitive.to_proto(resource.collation):
            res.collation = Primitive.to_proto(resource.collation)
        if InstanceSettingsDenyMaintenancePeriodsArray.to_proto(
            resource.deny_maintenance_periods
        ):
            res.deny_maintenance_periods.extend(
                InstanceSettingsDenyMaintenancePeriodsArray.to_proto(
                    resource.deny_maintenance_periods
                )
            )
        if InstanceSettingsInsightsConfig.to_proto(resource.insights_config):
            res.insights_config.CopyFrom(
                InstanceSettingsInsightsConfig.to_proto(resource.insights_config)
            )
        else:
            res.ClearField("insights_config")
        return res

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return None

        return InstanceSettings(
            authorized_gae_applications=Primitive.from_proto(
                resource.authorized_gae_applications
            ),
            tier=Primitive.from_proto(resource.tier),
            kind=Primitive.from_proto(resource.kind),
            availability_type=InstanceSettingsAvailabilityTypeEnum.from_proto(
                resource.availability_type
            ),
            pricing_plan=InstanceSettingsPricingPlanEnum.from_proto(
                resource.pricing_plan
            ),
            replication_type=InstanceSettingsReplicationTypeEnum.from_proto(
                resource.replication_type
            ),
            activation_policy=InstanceSettingsActivationPolicyEnum.from_proto(
                resource.activation_policy
            ),
            storage_auto_resize=Primitive.from_proto(resource.storage_auto_resize),
            data_disk_type=InstanceSettingsDataDiskTypeEnum.from_proto(
                resource.data_disk_type
            ),
            database_replication_enabled=Primitive.from_proto(
                resource.database_replication_enabled
            ),
            crash_safe_replication_enabled=Primitive.from_proto(
                resource.crash_safe_replication_enabled
            ),
            settings_version=InstanceSettingsSettingsVersion.from_proto(
                resource.settings_version
            ),
            user_labels=Primitive.from_proto(resource.user_labels),
            storage_auto_resize_limit=InstanceSettingsStorageAutoResizeLimit.from_proto(
                resource.storage_auto_resize_limit
            ),
            ip_configuration=InstanceSettingsIPConfiguration.from_proto(
                resource.ip_configuration
            ),
            location_preference=InstanceSettingsLocationPreference.from_proto(
                resource.location_preference
            ),
            database_flags=InstanceSettingsDatabaseFlagsArray.from_proto(
                resource.database_flags
            ),
            maintenance_window=InstanceSettingsMaintenanceWindow.from_proto(
                resource.maintenance_window
            ),
            backup_configuration=InstanceSettingsBackupConfiguration.from_proto(
                resource.backup_configuration
            ),
            data_disk_size_gb=InstanceSettingsDataDiskSizeGb.from_proto(
                resource.data_disk_size_gb
            ),
            active_directory_config=InstanceSettingsActiveDirectoryConfig.from_proto(
                resource.active_directory_config
            ),
            collation=Primitive.from_proto(resource.collation),
            deny_maintenance_periods=InstanceSettingsDenyMaintenancePeriodsArray.from_proto(
                resource.deny_maintenance_periods
            ),
            insights_config=InstanceSettingsInsightsConfig.from_proto(
                resource.insights_config
            ),
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


class InstanceSettingsSettingsVersion(object):
    def __init__(self, value: int = None):
        self.value = value

    @classmethod
    def to_proto(self, resource):
        if not resource:
            return None

        res = instance_pb2.SqlBetaInstanceSettingsSettingsVersion()
        if Primitive.to_proto(resource.value):
            res.value = Primitive.to_proto(resource.value)
        return res

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return None

        return InstanceSettingsSettingsVersion(
            value=Primitive.from_proto(resource.value),
        )


class InstanceSettingsSettingsVersionArray(object):
    @classmethod
    def to_proto(self, resources):
        if not resources:
            return resources
        return [InstanceSettingsSettingsVersion.to_proto(i) for i in resources]

    @classmethod
    def from_proto(self, resources):
        return [InstanceSettingsSettingsVersion.from_proto(i) for i in resources]


class InstanceSettingsStorageAutoResizeLimit(object):
    def __init__(self, value: int = None):
        self.value = value

    @classmethod
    def to_proto(self, resource):
        if not resource:
            return None

        res = instance_pb2.SqlBetaInstanceSettingsStorageAutoResizeLimit()
        if Primitive.to_proto(resource.value):
            res.value = Primitive.to_proto(resource.value)
        return res

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return None

        return InstanceSettingsStorageAutoResizeLimit(
            value=Primitive.from_proto(resource.value),
        )


class InstanceSettingsStorageAutoResizeLimitArray(object):
    @classmethod
    def to_proto(self, resources):
        if not resources:
            return resources
        return [InstanceSettingsStorageAutoResizeLimit.to_proto(i) for i in resources]

    @classmethod
    def from_proto(self, resources):
        return [InstanceSettingsStorageAutoResizeLimit.from_proto(i) for i in resources]


class InstanceSettingsIPConfiguration(object):
    def __init__(
        self,
        ipv4_enabled: bool = None,
        private_network: str = None,
        require_ssl: bool = None,
        authorized_networks: list = None,
    ):
        self.ipv4_enabled = ipv4_enabled
        self.private_network = private_network
        self.require_ssl = require_ssl
        self.authorized_networks = authorized_networks

    @classmethod
    def to_proto(self, resource):
        if not resource:
            return None

        res = instance_pb2.SqlBetaInstanceSettingsIPConfiguration()
        if Primitive.to_proto(resource.ipv4_enabled):
            res.ipv4_enabled = Primitive.to_proto(resource.ipv4_enabled)
        if Primitive.to_proto(resource.private_network):
            res.private_network = Primitive.to_proto(resource.private_network)
        if Primitive.to_proto(resource.require_ssl):
            res.require_ssl = Primitive.to_proto(resource.require_ssl)
        if InstanceSettingsIPConfigurationAuthorizedNetworksArray.to_proto(
            resource.authorized_networks
        ):
            res.authorized_networks.extend(
                InstanceSettingsIPConfigurationAuthorizedNetworksArray.to_proto(
                    resource.authorized_networks
                )
            )
        return res

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return None

        return InstanceSettingsIPConfiguration(
            ipv4_enabled=Primitive.from_proto(resource.ipv4_enabled),
            private_network=Primitive.from_proto(resource.private_network),
            require_ssl=Primitive.from_proto(resource.require_ssl),
            authorized_networks=InstanceSettingsIPConfigurationAuthorizedNetworksArray.from_proto(
                resource.authorized_networks
            ),
        )


class InstanceSettingsIPConfigurationArray(object):
    @classmethod
    def to_proto(self, resources):
        if not resources:
            return resources
        return [InstanceSettingsIPConfiguration.to_proto(i) for i in resources]

    @classmethod
    def from_proto(self, resources):
        return [InstanceSettingsIPConfiguration.from_proto(i) for i in resources]


class InstanceSettingsIPConfigurationAuthorizedNetworks(object):
    def __init__(
        self,
        value: str = None,
        expiration_time: str = None,
        name: str = None,
        kind: str = None,
    ):
        self.value = value
        self.expiration_time = expiration_time
        self.name = name
        self.kind = kind

    @classmethod
    def to_proto(self, resource):
        if not resource:
            return None

        res = instance_pb2.SqlBetaInstanceSettingsIPConfigurationAuthorizedNetworks()
        if Primitive.to_proto(resource.value):
            res.value = Primitive.to_proto(resource.value)
        if Primitive.to_proto(resource.expiration_time):
            res.expiration_time = Primitive.to_proto(resource.expiration_time)
        if Primitive.to_proto(resource.name):
            res.name = Primitive.to_proto(resource.name)
        if Primitive.to_proto(resource.kind):
            res.kind = Primitive.to_proto(resource.kind)
        return res

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return None

        return InstanceSettingsIPConfigurationAuthorizedNetworks(
            value=Primitive.from_proto(resource.value),
            expiration_time=Primitive.from_proto(resource.expiration_time),
            name=Primitive.from_proto(resource.name),
            kind=Primitive.from_proto(resource.kind),
        )


class InstanceSettingsIPConfigurationAuthorizedNetworksArray(object):
    @classmethod
    def to_proto(self, resources):
        if not resources:
            return resources
        return [
            InstanceSettingsIPConfigurationAuthorizedNetworks.to_proto(i)
            for i in resources
        ]

    @classmethod
    def from_proto(self, resources):
        return [
            InstanceSettingsIPConfigurationAuthorizedNetworks.from_proto(i)
            for i in resources
        ]


class InstanceSettingsLocationPreference(object):
    def __init__(self, zone: str = None, kind: str = None):
        self.zone = zone
        self.kind = kind

    @classmethod
    def to_proto(self, resource):
        if not resource:
            return None

        res = instance_pb2.SqlBetaInstanceSettingsLocationPreference()
        if Primitive.to_proto(resource.zone):
            res.zone = Primitive.to_proto(resource.zone)
        if Primitive.to_proto(resource.kind):
            res.kind = Primitive.to_proto(resource.kind)
        return res

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return None

        return InstanceSettingsLocationPreference(
            zone=Primitive.from_proto(resource.zone),
            kind=Primitive.from_proto(resource.kind),
        )


class InstanceSettingsLocationPreferenceArray(object):
    @classmethod
    def to_proto(self, resources):
        if not resources:
            return resources
        return [InstanceSettingsLocationPreference.to_proto(i) for i in resources]

    @classmethod
    def from_proto(self, resources):
        return [InstanceSettingsLocationPreference.from_proto(i) for i in resources]


class InstanceSettingsDatabaseFlags(object):
    def __init__(self, name: str = None, value: str = None):
        self.name = name
        self.value = value

    @classmethod
    def to_proto(self, resource):
        if not resource:
            return None

        res = instance_pb2.SqlBetaInstanceSettingsDatabaseFlags()
        if Primitive.to_proto(resource.name):
            res.name = Primitive.to_proto(resource.name)
        if Primitive.to_proto(resource.value):
            res.value = Primitive.to_proto(resource.value)
        return res

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return None

        return InstanceSettingsDatabaseFlags(
            name=Primitive.from_proto(resource.name),
            value=Primitive.from_proto(resource.value),
        )


class InstanceSettingsDatabaseFlagsArray(object):
    @classmethod
    def to_proto(self, resources):
        if not resources:
            return resources
        return [InstanceSettingsDatabaseFlags.to_proto(i) for i in resources]

    @classmethod
    def from_proto(self, resources):
        return [InstanceSettingsDatabaseFlags.from_proto(i) for i in resources]


class InstanceSettingsMaintenanceWindow(object):
    def __init__(
        self,
        hour: int = None,
        day: int = None,
        update_track: str = None,
        kind: str = None,
    ):
        self.hour = hour
        self.day = day
        self.update_track = update_track
        self.kind = kind

    @classmethod
    def to_proto(self, resource):
        if not resource:
            return None

        res = instance_pb2.SqlBetaInstanceSettingsMaintenanceWindow()
        if Primitive.to_proto(resource.hour):
            res.hour = Primitive.to_proto(resource.hour)
        if Primitive.to_proto(resource.day):
            res.day = Primitive.to_proto(resource.day)
        if InstanceSettingsMaintenanceWindowUpdateTrackEnum.to_proto(
            resource.update_track
        ):
            res.update_track = InstanceSettingsMaintenanceWindowUpdateTrackEnum.to_proto(
                resource.update_track
            )
        if Primitive.to_proto(resource.kind):
            res.kind = Primitive.to_proto(resource.kind)
        return res

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return None

        return InstanceSettingsMaintenanceWindow(
            hour=Primitive.from_proto(resource.hour),
            day=Primitive.from_proto(resource.day),
            update_track=InstanceSettingsMaintenanceWindowUpdateTrackEnum.from_proto(
                resource.update_track
            ),
            kind=Primitive.from_proto(resource.kind),
        )


class InstanceSettingsMaintenanceWindowArray(object):
    @classmethod
    def to_proto(self, resources):
        if not resources:
            return resources
        return [InstanceSettingsMaintenanceWindow.to_proto(i) for i in resources]

    @classmethod
    def from_proto(self, resources):
        return [InstanceSettingsMaintenanceWindow.from_proto(i) for i in resources]


class InstanceSettingsBackupConfiguration(object):
    def __init__(
        self,
        start_time: str = None,
        enabled: bool = None,
        kind: str = None,
        binary_log_enabled: bool = None,
        location: str = None,
        backup_retention_settings: dict = None,
        transaction_log_retention_days: int = None,
    ):
        self.start_time = start_time
        self.enabled = enabled
        self.kind = kind
        self.binary_log_enabled = binary_log_enabled
        self.location = location
        self.backup_retention_settings = backup_retention_settings
        self.transaction_log_retention_days = transaction_log_retention_days

    @classmethod
    def to_proto(self, resource):
        if not resource:
            return None

        res = instance_pb2.SqlBetaInstanceSettingsBackupConfiguration()
        if Primitive.to_proto(resource.start_time):
            res.start_time = Primitive.to_proto(resource.start_time)
        if Primitive.to_proto(resource.enabled):
            res.enabled = Primitive.to_proto(resource.enabled)
        if Primitive.to_proto(resource.kind):
            res.kind = Primitive.to_proto(resource.kind)
        if Primitive.to_proto(resource.binary_log_enabled):
            res.binary_log_enabled = Primitive.to_proto(resource.binary_log_enabled)
        if Primitive.to_proto(resource.location):
            res.location = Primitive.to_proto(resource.location)
        if InstanceSettingsBackupConfigurationBackupRetentionSettings.to_proto(
            resource.backup_retention_settings
        ):
            res.backup_retention_settings.CopyFrom(
                InstanceSettingsBackupConfigurationBackupRetentionSettings.to_proto(
                    resource.backup_retention_settings
                )
            )
        else:
            res.ClearField("backup_retention_settings")
        if Primitive.to_proto(resource.transaction_log_retention_days):
            res.transaction_log_retention_days = Primitive.to_proto(
                resource.transaction_log_retention_days
            )
        return res

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return None

        return InstanceSettingsBackupConfiguration(
            start_time=Primitive.from_proto(resource.start_time),
            enabled=Primitive.from_proto(resource.enabled),
            kind=Primitive.from_proto(resource.kind),
            binary_log_enabled=Primitive.from_proto(resource.binary_log_enabled),
            location=Primitive.from_proto(resource.location),
            backup_retention_settings=InstanceSettingsBackupConfigurationBackupRetentionSettings.from_proto(
                resource.backup_retention_settings
            ),
            transaction_log_retention_days=Primitive.from_proto(
                resource.transaction_log_retention_days
            ),
        )


class InstanceSettingsBackupConfigurationArray(object):
    @classmethod
    def to_proto(self, resources):
        if not resources:
            return resources
        return [InstanceSettingsBackupConfiguration.to_proto(i) for i in resources]

    @classmethod
    def from_proto(self, resources):
        return [InstanceSettingsBackupConfiguration.from_proto(i) for i in resources]


class InstanceSettingsBackupConfigurationBackupRetentionSettings(object):
    def __init__(self, retention_unit: str = None, retained_backups: int = None):
        self.retention_unit = retention_unit
        self.retained_backups = retained_backups

    @classmethod
    def to_proto(self, resource):
        if not resource:
            return None

        res = (
            instance_pb2.SqlBetaInstanceSettingsBackupConfigurationBackupRetentionSettings()
        )
        if InstanceSettingsBackupConfigurationBackupRetentionSettingsRetentionUnitEnum.to_proto(
            resource.retention_unit
        ):
            res.retention_unit = InstanceSettingsBackupConfigurationBackupRetentionSettingsRetentionUnitEnum.to_proto(
                resource.retention_unit
            )
        if Primitive.to_proto(resource.retained_backups):
            res.retained_backups = Primitive.to_proto(resource.retained_backups)
        return res

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return None

        return InstanceSettingsBackupConfigurationBackupRetentionSettings(
            retention_unit=InstanceSettingsBackupConfigurationBackupRetentionSettingsRetentionUnitEnum.from_proto(
                resource.retention_unit
            ),
            retained_backups=Primitive.from_proto(resource.retained_backups),
        )


class InstanceSettingsBackupConfigurationBackupRetentionSettingsArray(object):
    @classmethod
    def to_proto(self, resources):
        if not resources:
            return resources
        return [
            InstanceSettingsBackupConfigurationBackupRetentionSettings.to_proto(i)
            for i in resources
        ]

    @classmethod
    def from_proto(self, resources):
        return [
            InstanceSettingsBackupConfigurationBackupRetentionSettings.from_proto(i)
            for i in resources
        ]


class InstanceSettingsDataDiskSizeGb(object):
    def __init__(self, value: int = None):
        self.value = value

    @classmethod
    def to_proto(self, resource):
        if not resource:
            return None

        res = instance_pb2.SqlBetaInstanceSettingsDataDiskSizeGb()
        if Primitive.to_proto(resource.value):
            res.value = Primitive.to_proto(resource.value)
        return res

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return None

        return InstanceSettingsDataDiskSizeGb(
            value=Primitive.from_proto(resource.value),
        )


class InstanceSettingsDataDiskSizeGbArray(object):
    @classmethod
    def to_proto(self, resources):
        if not resources:
            return resources
        return [InstanceSettingsDataDiskSizeGb.to_proto(i) for i in resources]

    @classmethod
    def from_proto(self, resources):
        return [InstanceSettingsDataDiskSizeGb.from_proto(i) for i in resources]


class InstanceSettingsActiveDirectoryConfig(object):
    def __init__(self, kind: str = None, domain: str = None):
        self.kind = kind
        self.domain = domain

    @classmethod
    def to_proto(self, resource):
        if not resource:
            return None

        res = instance_pb2.SqlBetaInstanceSettingsActiveDirectoryConfig()
        if Primitive.to_proto(resource.kind):
            res.kind = Primitive.to_proto(resource.kind)
        if Primitive.to_proto(resource.domain):
            res.domain = Primitive.to_proto(resource.domain)
        return res

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return None

        return InstanceSettingsActiveDirectoryConfig(
            kind=Primitive.from_proto(resource.kind),
            domain=Primitive.from_proto(resource.domain),
        )


class InstanceSettingsActiveDirectoryConfigArray(object):
    @classmethod
    def to_proto(self, resources):
        if not resources:
            return resources
        return [InstanceSettingsActiveDirectoryConfig.to_proto(i) for i in resources]

    @classmethod
    def from_proto(self, resources):
        return [InstanceSettingsActiveDirectoryConfig.from_proto(i) for i in resources]


class InstanceSettingsDenyMaintenancePeriods(object):
    def __init__(self, start_date: str = None, end_date: str = None, time: str = None):
        self.start_date = start_date
        self.end_date = end_date
        self.time = time

    @classmethod
    def to_proto(self, resource):
        if not resource:
            return None

        res = instance_pb2.SqlBetaInstanceSettingsDenyMaintenancePeriods()
        if Primitive.to_proto(resource.start_date):
            res.start_date = Primitive.to_proto(resource.start_date)
        if Primitive.to_proto(resource.end_date):
            res.end_date = Primitive.to_proto(resource.end_date)
        if Primitive.to_proto(resource.time):
            res.time = Primitive.to_proto(resource.time)
        return res

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return None

        return InstanceSettingsDenyMaintenancePeriods(
            start_date=Primitive.from_proto(resource.start_date),
            end_date=Primitive.from_proto(resource.end_date),
            time=Primitive.from_proto(resource.time),
        )


class InstanceSettingsDenyMaintenancePeriodsArray(object):
    @classmethod
    def to_proto(self, resources):
        if not resources:
            return resources
        return [InstanceSettingsDenyMaintenancePeriods.to_proto(i) for i in resources]

    @classmethod
    def from_proto(self, resources):
        return [InstanceSettingsDenyMaintenancePeriods.from_proto(i) for i in resources]


class InstanceSettingsInsightsConfig(object):
    def __init__(
        self,
        query_insights_enabled: bool = None,
        record_client_address: bool = None,
        record_application_tags: bool = None,
        query_string_length: int = None,
    ):
        self.query_insights_enabled = query_insights_enabled
        self.record_client_address = record_client_address
        self.record_application_tags = record_application_tags
        self.query_string_length = query_string_length

    @classmethod
    def to_proto(self, resource):
        if not resource:
            return None

        res = instance_pb2.SqlBetaInstanceSettingsInsightsConfig()
        if Primitive.to_proto(resource.query_insights_enabled):
            res.query_insights_enabled = Primitive.to_proto(
                resource.query_insights_enabled
            )
        if Primitive.to_proto(resource.record_client_address):
            res.record_client_address = Primitive.to_proto(
                resource.record_client_address
            )
        if Primitive.to_proto(resource.record_application_tags):
            res.record_application_tags = Primitive.to_proto(
                resource.record_application_tags
            )
        if Primitive.to_proto(resource.query_string_length):
            res.query_string_length = Primitive.to_proto(resource.query_string_length)
        return res

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return None

        return InstanceSettingsInsightsConfig(
            query_insights_enabled=Primitive.from_proto(
                resource.query_insights_enabled
            ),
            record_client_address=Primitive.from_proto(resource.record_client_address),
            record_application_tags=Primitive.from_proto(
                resource.record_application_tags
            ),
            query_string_length=Primitive.from_proto(resource.query_string_length),
        )


class InstanceSettingsInsightsConfigArray(object):
    @classmethod
    def to_proto(self, resources):
        if not resources:
            return resources
        return [InstanceSettingsInsightsConfig.to_proto(i) for i in resources]

    @classmethod
    def from_proto(self, resources):
        return [InstanceSettingsInsightsConfig.from_proto(i) for i in resources]


class InstanceReplicaInstances(object):
    def __init__(self, name: str = None, region: str = None):
        self.name = name
        self.region = region

    @classmethod
    def to_proto(self, resource):
        if not resource:
            return None

        res = instance_pb2.SqlBetaInstanceReplicaInstances()
        if Primitive.to_proto(resource.name):
            res.name = Primitive.to_proto(resource.name)
        if Primitive.to_proto(resource.region):
            res.region = Primitive.to_proto(resource.region)
        return res

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return None

        return InstanceReplicaInstances(
            name=Primitive.from_proto(resource.name),
            region=Primitive.from_proto(resource.region),
        )


class InstanceReplicaInstancesArray(object):
    @classmethod
    def to_proto(self, resources):
        if not resources:
            return resources
        return [InstanceReplicaInstances.to_proto(i) for i in resources]

    @classmethod
    def from_proto(self, resources):
        return [InstanceReplicaInstances.from_proto(i) for i in resources]


class InstanceServerCaCert(object):
    def __init__(
        self,
        kind: str = None,
        cert_serial_number: str = None,
        cert: str = None,
        create_time: str = None,
        common_name: str = None,
        expiration_time: str = None,
        sha1_fingerprint: str = None,
        instance: str = None,
    ):
        self.kind = kind
        self.cert_serial_number = cert_serial_number
        self.cert = cert
        self.create_time = create_time
        self.common_name = common_name
        self.expiration_time = expiration_time
        self.sha1_fingerprint = sha1_fingerprint
        self.instance = instance

    @classmethod
    def to_proto(self, resource):
        if not resource:
            return None

        res = instance_pb2.SqlBetaInstanceServerCaCert()
        if Primitive.to_proto(resource.kind):
            res.kind = Primitive.to_proto(resource.kind)
        if Primitive.to_proto(resource.cert_serial_number):
            res.cert_serial_number = Primitive.to_proto(resource.cert_serial_number)
        if Primitive.to_proto(resource.cert):
            res.cert = Primitive.to_proto(resource.cert)
        if Primitive.to_proto(resource.create_time):
            res.create_time = Primitive.to_proto(resource.create_time)
        if Primitive.to_proto(resource.common_name):
            res.common_name = Primitive.to_proto(resource.common_name)
        if Primitive.to_proto(resource.expiration_time):
            res.expiration_time = Primitive.to_proto(resource.expiration_time)
        if Primitive.to_proto(resource.sha1_fingerprint):
            res.sha1_fingerprint = Primitive.to_proto(resource.sha1_fingerprint)
        if Primitive.to_proto(resource.instance):
            res.instance = Primitive.to_proto(resource.instance)
        return res

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return None

        return InstanceServerCaCert(
            kind=Primitive.from_proto(resource.kind),
            cert_serial_number=Primitive.from_proto(resource.cert_serial_number),
            cert=Primitive.from_proto(resource.cert),
            create_time=Primitive.from_proto(resource.create_time),
            common_name=Primitive.from_proto(resource.common_name),
            expiration_time=Primitive.from_proto(resource.expiration_time),
            sha1_fingerprint=Primitive.from_proto(resource.sha1_fingerprint),
            instance=Primitive.from_proto(resource.instance),
        )


class InstanceServerCaCertArray(object):
    @classmethod
    def to_proto(self, resources):
        if not resources:
            return resources
        return [InstanceServerCaCert.to_proto(i) for i in resources]

    @classmethod
    def from_proto(self, resources):
        return [InstanceServerCaCert.from_proto(i) for i in resources]


class InstanceOnPremisesConfiguration(object):
    def __init__(
        self,
        host_port: str = None,
        kind: str = None,
        username: str = None,
        password: str = None,
        ca_certificate: str = None,
        client_certificate: str = None,
        client_key: str = None,
        dump_file_path: str = None,
        database: str = None,
        replicated_databases: list = None,
    ):
        self.host_port = host_port
        self.kind = kind
        self.username = username
        self.password = password
        self.ca_certificate = ca_certificate
        self.client_certificate = client_certificate
        self.client_key = client_key
        self.dump_file_path = dump_file_path
        self.database = database
        self.replicated_databases = replicated_databases

    @classmethod
    def to_proto(self, resource):
        if not resource:
            return None

        res = instance_pb2.SqlBetaInstanceOnPremisesConfiguration()
        if Primitive.to_proto(resource.host_port):
            res.host_port = Primitive.to_proto(resource.host_port)
        if Primitive.to_proto(resource.kind):
            res.kind = Primitive.to_proto(resource.kind)
        if Primitive.to_proto(resource.username):
            res.username = Primitive.to_proto(resource.username)
        if Primitive.to_proto(resource.password):
            res.password = Primitive.to_proto(resource.password)
        if Primitive.to_proto(resource.ca_certificate):
            res.ca_certificate = Primitive.to_proto(resource.ca_certificate)
        if Primitive.to_proto(resource.client_certificate):
            res.client_certificate = Primitive.to_proto(resource.client_certificate)
        if Primitive.to_proto(resource.client_key):
            res.client_key = Primitive.to_proto(resource.client_key)
        if Primitive.to_proto(resource.dump_file_path):
            res.dump_file_path = Primitive.to_proto(resource.dump_file_path)
        if Primitive.to_proto(resource.database):
            res.database = Primitive.to_proto(resource.database)
        if Primitive.to_proto(resource.replicated_databases):
            res.replicated_databases.extend(
                Primitive.to_proto(resource.replicated_databases)
            )
        return res

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return None

        return InstanceOnPremisesConfiguration(
            host_port=Primitive.from_proto(resource.host_port),
            kind=Primitive.from_proto(resource.kind),
            username=Primitive.from_proto(resource.username),
            password=Primitive.from_proto(resource.password),
            ca_certificate=Primitive.from_proto(resource.ca_certificate),
            client_certificate=Primitive.from_proto(resource.client_certificate),
            client_key=Primitive.from_proto(resource.client_key),
            dump_file_path=Primitive.from_proto(resource.dump_file_path),
            database=Primitive.from_proto(resource.database),
            replicated_databases=Primitive.from_proto(resource.replicated_databases),
        )


class InstanceOnPremisesConfigurationArray(object):
    @classmethod
    def to_proto(self, resources):
        if not resources:
            return resources
        return [InstanceOnPremisesConfiguration.to_proto(i) for i in resources]

    @classmethod
    def from_proto(self, resources):
        return [InstanceOnPremisesConfiguration.from_proto(i) for i in resources]


class InstanceDiskEncryptionStatus(object):
    def __init__(self, kms_key_version_name: str = None, kind: str = None):
        self.kms_key_version_name = kms_key_version_name
        self.kind = kind

    @classmethod
    def to_proto(self, resource):
        if not resource:
            return None

        res = instance_pb2.SqlBetaInstanceDiskEncryptionStatus()
        if Primitive.to_proto(resource.kms_key_version_name):
            res.kms_key_version_name = Primitive.to_proto(resource.kms_key_version_name)
        if Primitive.to_proto(resource.kind):
            res.kind = Primitive.to_proto(resource.kind)
        return res

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return None

        return InstanceDiskEncryptionStatus(
            kms_key_version_name=Primitive.from_proto(resource.kms_key_version_name),
            kind=Primitive.from_proto(resource.kind),
        )


class InstanceDiskEncryptionStatusArray(object):
    @classmethod
    def to_proto(self, resources):
        if not resources:
            return resources
        return [InstanceDiskEncryptionStatus.to_proto(i) for i in resources]

    @classmethod
    def from_proto(self, resources):
        return [InstanceDiskEncryptionStatus.from_proto(i) for i in resources]


class InstanceBackendTypeEnum(object):
    @classmethod
    def to_proto(self, resource):
        if not resource:
            return resource
        return instance_pb2.SqlBetaInstanceBackendTypeEnum.Value(
            "SqlBetaInstanceBackendTypeEnum%s" % resource
        )

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return resource
        return instance_pb2.SqlBetaInstanceBackendTypeEnum.Name(resource)[
            len("SqlBetaInstanceBackendTypeEnum") :
        ]


class InstanceDatabaseVersionEnum(object):
    @classmethod
    def to_proto(self, resource):
        if not resource:
            return resource
        return instance_pb2.SqlBetaInstanceDatabaseVersionEnum.Value(
            "SqlBetaInstanceDatabaseVersionEnum%s" % resource
        )

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return resource
        return instance_pb2.SqlBetaInstanceDatabaseVersionEnum.Name(resource)[
            len("SqlBetaInstanceDatabaseVersionEnum") :
        ]


class InstanceInstanceTypeEnum(object):
    @classmethod
    def to_proto(self, resource):
        if not resource:
            return resource
        return instance_pb2.SqlBetaInstanceInstanceTypeEnum.Value(
            "SqlBetaInstanceInstanceTypeEnum%s" % resource
        )

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return resource
        return instance_pb2.SqlBetaInstanceInstanceTypeEnum.Name(resource)[
            len("SqlBetaInstanceInstanceTypeEnum") :
        ]


class InstanceIPAddressesTypeEnum(object):
    @classmethod
    def to_proto(self, resource):
        if not resource:
            return resource
        return instance_pb2.SqlBetaInstanceIPAddressesTypeEnum.Value(
            "SqlBetaInstanceIPAddressesTypeEnum%s" % resource
        )

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return resource
        return instance_pb2.SqlBetaInstanceIPAddressesTypeEnum.Name(resource)[
            len("SqlBetaInstanceIPAddressesTypeEnum") :
        ]


class InstanceSettingsAvailabilityTypeEnum(object):
    @classmethod
    def to_proto(self, resource):
        if not resource:
            return resource
        return instance_pb2.SqlBetaInstanceSettingsAvailabilityTypeEnum.Value(
            "SqlBetaInstanceSettingsAvailabilityTypeEnum%s" % resource
        )

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return resource
        return instance_pb2.SqlBetaInstanceSettingsAvailabilityTypeEnum.Name(resource)[
            len("SqlBetaInstanceSettingsAvailabilityTypeEnum") :
        ]


class InstanceSettingsPricingPlanEnum(object):
    @classmethod
    def to_proto(self, resource):
        if not resource:
            return resource
        return instance_pb2.SqlBetaInstanceSettingsPricingPlanEnum.Value(
            "SqlBetaInstanceSettingsPricingPlanEnum%s" % resource
        )

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return resource
        return instance_pb2.SqlBetaInstanceSettingsPricingPlanEnum.Name(resource)[
            len("SqlBetaInstanceSettingsPricingPlanEnum") :
        ]


class InstanceSettingsReplicationTypeEnum(object):
    @classmethod
    def to_proto(self, resource):
        if not resource:
            return resource
        return instance_pb2.SqlBetaInstanceSettingsReplicationTypeEnum.Value(
            "SqlBetaInstanceSettingsReplicationTypeEnum%s" % resource
        )

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return resource
        return instance_pb2.SqlBetaInstanceSettingsReplicationTypeEnum.Name(resource)[
            len("SqlBetaInstanceSettingsReplicationTypeEnum") :
        ]


class InstanceSettingsActivationPolicyEnum(object):
    @classmethod
    def to_proto(self, resource):
        if not resource:
            return resource
        return instance_pb2.SqlBetaInstanceSettingsActivationPolicyEnum.Value(
            "SqlBetaInstanceSettingsActivationPolicyEnum%s" % resource
        )

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return resource
        return instance_pb2.SqlBetaInstanceSettingsActivationPolicyEnum.Name(resource)[
            len("SqlBetaInstanceSettingsActivationPolicyEnum") :
        ]


class InstanceSettingsDataDiskTypeEnum(object):
    @classmethod
    def to_proto(self, resource):
        if not resource:
            return resource
        return instance_pb2.SqlBetaInstanceSettingsDataDiskTypeEnum.Value(
            "SqlBetaInstanceSettingsDataDiskTypeEnum%s" % resource
        )

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return resource
        return instance_pb2.SqlBetaInstanceSettingsDataDiskTypeEnum.Name(resource)[
            len("SqlBetaInstanceSettingsDataDiskTypeEnum") :
        ]


class InstanceSettingsMaintenanceWindowUpdateTrackEnum(object):
    @classmethod
    def to_proto(self, resource):
        if not resource:
            return resource
        return instance_pb2.SqlBetaInstanceSettingsMaintenanceWindowUpdateTrackEnum.Value(
            "SqlBetaInstanceSettingsMaintenanceWindowUpdateTrackEnum%s" % resource
        )

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return resource
        return instance_pb2.SqlBetaInstanceSettingsMaintenanceWindowUpdateTrackEnum.Name(
            resource
        )[
            len("SqlBetaInstanceSettingsMaintenanceWindowUpdateTrackEnum") :
        ]


class InstanceSettingsBackupConfigurationBackupRetentionSettingsRetentionUnitEnum(
    object
):
    @classmethod
    def to_proto(self, resource):
        if not resource:
            return resource
        return instance_pb2.SqlBetaInstanceSettingsBackupConfigurationBackupRetentionSettingsRetentionUnitEnum.Value(
            "SqlBetaInstanceSettingsBackupConfigurationBackupRetentionSettingsRetentionUnitEnum%s"
            % resource
        )

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return resource
        return instance_pb2.SqlBetaInstanceSettingsBackupConfigurationBackupRetentionSettingsRetentionUnitEnum.Name(
            resource
        )[
            len(
                "SqlBetaInstanceSettingsBackupConfigurationBackupRetentionSettingsRetentionUnitEnum"
            ) :
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
