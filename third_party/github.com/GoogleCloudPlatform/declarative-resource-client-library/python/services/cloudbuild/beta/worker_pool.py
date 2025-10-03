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
from google3.cloud.graphite.mmv2.services.google.cloud_build import worker_pool_pb2
from google3.cloud.graphite.mmv2.services.google.cloud_build import worker_pool_pb2_grpc

from typing import List


class WorkerPool(object):
    def __init__(
        self,
        name: str = None,
        display_name: str = None,
        uid: str = None,
        annotations: dict = None,
        create_time: str = None,
        update_time: str = None,
        delete_time: str = None,
        state: str = None,
        private_pool_v1_config: dict = None,
        etag: str = None,
        worker_config: dict = None,
        network_config: dict = None,
        project: str = None,
        location: str = None,
        service_account_file: str = "",
    ):
        channel.initialize()
        self.name = name
        self.display_name = display_name
        self.annotations = annotations
        self.private_pool_v1_config = private_pool_v1_config
        self.worker_config = worker_config
        self.network_config = network_config
        self.project = project
        self.location = location
        self.service_account_file = service_account_file

    def apply(self):
        stub = worker_pool_pb2_grpc.CloudbuildBetaWorkerPoolServiceStub(
            channel.Channel()
        )
        request = worker_pool_pb2.ApplyCloudbuildBetaWorkerPoolRequest()
        if Primitive.to_proto(self.name):
            request.resource.name = Primitive.to_proto(self.name)

        if Primitive.to_proto(self.display_name):
            request.resource.display_name = Primitive.to_proto(self.display_name)

        if Primitive.to_proto(self.annotations):
            request.resource.annotations = Primitive.to_proto(self.annotations)

        if WorkerPoolPrivatePoolV1Config.to_proto(self.private_pool_v1_config):
            request.resource.private_pool_v1_config.CopyFrom(
                WorkerPoolPrivatePoolV1Config.to_proto(self.private_pool_v1_config)
            )
        else:
            request.resource.ClearField("private_pool_v1_config")
        if WorkerPoolWorkerConfig.to_proto(self.worker_config):
            request.resource.worker_config.CopyFrom(
                WorkerPoolWorkerConfig.to_proto(self.worker_config)
            )
        else:
            request.resource.ClearField("worker_config")
        if WorkerPoolNetworkConfig.to_proto(self.network_config):
            request.resource.network_config.CopyFrom(
                WorkerPoolNetworkConfig.to_proto(self.network_config)
            )
        else:
            request.resource.ClearField("network_config")
        if Primitive.to_proto(self.project):
            request.resource.project = Primitive.to_proto(self.project)

        if Primitive.to_proto(self.location):
            request.resource.location = Primitive.to_proto(self.location)

        request.service_account_file = self.service_account_file

        response = stub.ApplyCloudbuildBetaWorkerPool(request)
        self.name = Primitive.from_proto(response.name)
        self.display_name = Primitive.from_proto(response.display_name)
        self.uid = Primitive.from_proto(response.uid)
        self.annotations = Primitive.from_proto(response.annotations)
        self.create_time = Primitive.from_proto(response.create_time)
        self.update_time = Primitive.from_proto(response.update_time)
        self.delete_time = Primitive.from_proto(response.delete_time)
        self.state = WorkerPoolStateEnum.from_proto(response.state)
        self.private_pool_v1_config = WorkerPoolPrivatePoolV1Config.from_proto(
            response.private_pool_v1_config
        )
        self.etag = Primitive.from_proto(response.etag)
        self.worker_config = WorkerPoolWorkerConfig.from_proto(response.worker_config)
        self.network_config = WorkerPoolNetworkConfig.from_proto(
            response.network_config
        )
        self.project = Primitive.from_proto(response.project)
        self.location = Primitive.from_proto(response.location)

    def delete(self):
        stub = worker_pool_pb2_grpc.CloudbuildBetaWorkerPoolServiceStub(
            channel.Channel()
        )
        request = worker_pool_pb2.DeleteCloudbuildBetaWorkerPoolRequest()
        request.service_account_file = self.service_account_file
        if Primitive.to_proto(self.name):
            request.resource.name = Primitive.to_proto(self.name)

        if Primitive.to_proto(self.display_name):
            request.resource.display_name = Primitive.to_proto(self.display_name)

        if Primitive.to_proto(self.annotations):
            request.resource.annotations = Primitive.to_proto(self.annotations)

        if WorkerPoolPrivatePoolV1Config.to_proto(self.private_pool_v1_config):
            request.resource.private_pool_v1_config.CopyFrom(
                WorkerPoolPrivatePoolV1Config.to_proto(self.private_pool_v1_config)
            )
        else:
            request.resource.ClearField("private_pool_v1_config")
        if WorkerPoolWorkerConfig.to_proto(self.worker_config):
            request.resource.worker_config.CopyFrom(
                WorkerPoolWorkerConfig.to_proto(self.worker_config)
            )
        else:
            request.resource.ClearField("worker_config")
        if WorkerPoolNetworkConfig.to_proto(self.network_config):
            request.resource.network_config.CopyFrom(
                WorkerPoolNetworkConfig.to_proto(self.network_config)
            )
        else:
            request.resource.ClearField("network_config")
        if Primitive.to_proto(self.project):
            request.resource.project = Primitive.to_proto(self.project)

        if Primitive.to_proto(self.location):
            request.resource.location = Primitive.to_proto(self.location)

        response = stub.DeleteCloudbuildBetaWorkerPool(request)

    @classmethod
    def list(self, project, location, service_account_file=""):
        stub = worker_pool_pb2_grpc.CloudbuildBetaWorkerPoolServiceStub(
            channel.Channel()
        )
        request = worker_pool_pb2.ListCloudbuildBetaWorkerPoolRequest()
        request.service_account_file = service_account_file
        request.Project = project

        request.Location = location

        return stub.ListCloudbuildBetaWorkerPool(request).items

    def to_proto(self):
        resource = worker_pool_pb2.CloudbuildBetaWorkerPool()
        if Primitive.to_proto(self.name):
            resource.name = Primitive.to_proto(self.name)
        if Primitive.to_proto(self.display_name):
            resource.display_name = Primitive.to_proto(self.display_name)
        if Primitive.to_proto(self.annotations):
            resource.annotations = Primitive.to_proto(self.annotations)
        if WorkerPoolPrivatePoolV1Config.to_proto(self.private_pool_v1_config):
            resource.private_pool_v1_config.CopyFrom(
                WorkerPoolPrivatePoolV1Config.to_proto(self.private_pool_v1_config)
            )
        else:
            resource.ClearField("private_pool_v1_config")
        if WorkerPoolWorkerConfig.to_proto(self.worker_config):
            resource.worker_config.CopyFrom(
                WorkerPoolWorkerConfig.to_proto(self.worker_config)
            )
        else:
            resource.ClearField("worker_config")
        if WorkerPoolNetworkConfig.to_proto(self.network_config):
            resource.network_config.CopyFrom(
                WorkerPoolNetworkConfig.to_proto(self.network_config)
            )
        else:
            resource.ClearField("network_config")
        if Primitive.to_proto(self.project):
            resource.project = Primitive.to_proto(self.project)
        if Primitive.to_proto(self.location):
            resource.location = Primitive.to_proto(self.location)
        return resource


class WorkerPoolPrivatePoolV1Config(object):
    def __init__(self, worker_config: dict = None, network_config: dict = None):
        self.worker_config = worker_config
        self.network_config = network_config

    @classmethod
    def to_proto(self, resource):
        if not resource:
            return None

        res = worker_pool_pb2.CloudbuildBetaWorkerPoolPrivatePoolV1Config()
        if WorkerPoolPrivatePoolV1ConfigWorkerConfig.to_proto(resource.worker_config):
            res.worker_config.CopyFrom(
                WorkerPoolPrivatePoolV1ConfigWorkerConfig.to_proto(
                    resource.worker_config
                )
            )
        else:
            res.ClearField("worker_config")
        if WorkerPoolPrivatePoolV1ConfigNetworkConfig.to_proto(resource.network_config):
            res.network_config.CopyFrom(
                WorkerPoolPrivatePoolV1ConfigNetworkConfig.to_proto(
                    resource.network_config
                )
            )
        else:
            res.ClearField("network_config")
        return res

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return None

        return WorkerPoolPrivatePoolV1Config(
            worker_config=WorkerPoolPrivatePoolV1ConfigWorkerConfig.from_proto(
                resource.worker_config
            ),
            network_config=WorkerPoolPrivatePoolV1ConfigNetworkConfig.from_proto(
                resource.network_config
            ),
        )


class WorkerPoolPrivatePoolV1ConfigArray(object):
    @classmethod
    def to_proto(self, resources):
        if not resources:
            return resources
        return [WorkerPoolPrivatePoolV1Config.to_proto(i) for i in resources]

    @classmethod
    def from_proto(self, resources):
        return [WorkerPoolPrivatePoolV1Config.from_proto(i) for i in resources]


class WorkerPoolPrivatePoolV1ConfigWorkerConfig(object):
    def __init__(self, machine_type: str = None, disk_size_gb: int = None):
        self.machine_type = machine_type
        self.disk_size_gb = disk_size_gb

    @classmethod
    def to_proto(self, resource):
        if not resource:
            return None

        res = worker_pool_pb2.CloudbuildBetaWorkerPoolPrivatePoolV1ConfigWorkerConfig()
        if Primitive.to_proto(resource.machine_type):
            res.machine_type = Primitive.to_proto(resource.machine_type)
        if Primitive.to_proto(resource.disk_size_gb):
            res.disk_size_gb = Primitive.to_proto(resource.disk_size_gb)
        return res

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return None

        return WorkerPoolPrivatePoolV1ConfigWorkerConfig(
            machine_type=Primitive.from_proto(resource.machine_type),
            disk_size_gb=Primitive.from_proto(resource.disk_size_gb),
        )


class WorkerPoolPrivatePoolV1ConfigWorkerConfigArray(object):
    @classmethod
    def to_proto(self, resources):
        if not resources:
            return resources
        return [
            WorkerPoolPrivatePoolV1ConfigWorkerConfig.to_proto(i) for i in resources
        ]

    @classmethod
    def from_proto(self, resources):
        return [
            WorkerPoolPrivatePoolV1ConfigWorkerConfig.from_proto(i) for i in resources
        ]


class WorkerPoolPrivatePoolV1ConfigNetworkConfig(object):
    def __init__(
        self,
        peered_network: str = None,
        peered_network_ip_range: str = None,
        egress_option: str = None,
    ):
        self.peered_network = peered_network
        self.peered_network_ip_range = peered_network_ip_range
        self.egress_option = egress_option

    @classmethod
    def to_proto(self, resource):
        if not resource:
            return None

        res = worker_pool_pb2.CloudbuildBetaWorkerPoolPrivatePoolV1ConfigNetworkConfig()
        if Primitive.to_proto(resource.peered_network):
            res.peered_network = Primitive.to_proto(resource.peered_network)
        if Primitive.to_proto(resource.peered_network_ip_range):
            res.peered_network_ip_range = Primitive.to_proto(
                resource.peered_network_ip_range
            )
        if WorkerPoolPrivatePoolV1ConfigNetworkConfigEgressOptionEnum.to_proto(
            resource.egress_option
        ):
            res.egress_option = (
                WorkerPoolPrivatePoolV1ConfigNetworkConfigEgressOptionEnum.to_proto(
                    resource.egress_option
                )
            )
        return res

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return None

        return WorkerPoolPrivatePoolV1ConfigNetworkConfig(
            peered_network=Primitive.from_proto(resource.peered_network),
            peered_network_ip_range=Primitive.from_proto(
                resource.peered_network_ip_range
            ),
            egress_option=WorkerPoolPrivatePoolV1ConfigNetworkConfigEgressOptionEnum.from_proto(
                resource.egress_option
            ),
        )


class WorkerPoolPrivatePoolV1ConfigNetworkConfigArray(object):
    @classmethod
    def to_proto(self, resources):
        if not resources:
            return resources
        return [
            WorkerPoolPrivatePoolV1ConfigNetworkConfig.to_proto(i) for i in resources
        ]

    @classmethod
    def from_proto(self, resources):
        return [
            WorkerPoolPrivatePoolV1ConfigNetworkConfig.from_proto(i) for i in resources
        ]


class WorkerPoolWorkerConfig(object):
    def __init__(
        self,
        machine_type: str = None,
        disk_size_gb: int = None,
        no_external_ip: bool = None,
    ):
        self.machine_type = machine_type
        self.disk_size_gb = disk_size_gb
        self.no_external_ip = no_external_ip

    @classmethod
    def to_proto(self, resource):
        if not resource:
            return None

        res = worker_pool_pb2.CloudbuildBetaWorkerPoolWorkerConfig()
        if Primitive.to_proto(resource.machine_type):
            res.machine_type = Primitive.to_proto(resource.machine_type)
        if Primitive.to_proto(resource.disk_size_gb):
            res.disk_size_gb = Primitive.to_proto(resource.disk_size_gb)
        if Primitive.to_proto(resource.no_external_ip):
            res.no_external_ip = Primitive.to_proto(resource.no_external_ip)
        return res

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return None

        return WorkerPoolWorkerConfig(
            machine_type=Primitive.from_proto(resource.machine_type),
            disk_size_gb=Primitive.from_proto(resource.disk_size_gb),
            no_external_ip=Primitive.from_proto(resource.no_external_ip),
        )


class WorkerPoolWorkerConfigArray(object):
    @classmethod
    def to_proto(self, resources):
        if not resources:
            return resources
        return [WorkerPoolWorkerConfig.to_proto(i) for i in resources]

    @classmethod
    def from_proto(self, resources):
        return [WorkerPoolWorkerConfig.from_proto(i) for i in resources]


class WorkerPoolNetworkConfig(object):
    def __init__(self, peered_network: str = None, peered_network_ip_range: str = None):
        self.peered_network = peered_network
        self.peered_network_ip_range = peered_network_ip_range

    @classmethod
    def to_proto(self, resource):
        if not resource:
            return None

        res = worker_pool_pb2.CloudbuildBetaWorkerPoolNetworkConfig()
        if Primitive.to_proto(resource.peered_network):
            res.peered_network = Primitive.to_proto(resource.peered_network)
        if Primitive.to_proto(resource.peered_network_ip_range):
            res.peered_network_ip_range = Primitive.to_proto(
                resource.peered_network_ip_range
            )
        return res

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return None

        return WorkerPoolNetworkConfig(
            peered_network=Primitive.from_proto(resource.peered_network),
            peered_network_ip_range=Primitive.from_proto(
                resource.peered_network_ip_range
            ),
        )


class WorkerPoolNetworkConfigArray(object):
    @classmethod
    def to_proto(self, resources):
        if not resources:
            return resources
        return [WorkerPoolNetworkConfig.to_proto(i) for i in resources]

    @classmethod
    def from_proto(self, resources):
        return [WorkerPoolNetworkConfig.from_proto(i) for i in resources]


class WorkerPoolStateEnum(object):
    @classmethod
    def to_proto(self, resource):
        if not resource:
            return resource
        return worker_pool_pb2.CloudbuildBetaWorkerPoolStateEnum.Value(
            "CloudbuildBetaWorkerPoolStateEnum%s" % resource
        )

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return resource
        return worker_pool_pb2.CloudbuildBetaWorkerPoolStateEnum.Name(resource)[
            len("CloudbuildBetaWorkerPoolStateEnum") :
        ]


class WorkerPoolPrivatePoolV1ConfigNetworkConfigEgressOptionEnum(object):
    @classmethod
    def to_proto(self, resource):
        if not resource:
            return resource
        return worker_pool_pb2.CloudbuildBetaWorkerPoolPrivatePoolV1ConfigNetworkConfigEgressOptionEnum.Value(
            "CloudbuildBetaWorkerPoolPrivatePoolV1ConfigNetworkConfigEgressOptionEnum%s"
            % resource
        )

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return resource
        return worker_pool_pb2.CloudbuildBetaWorkerPoolPrivatePoolV1ConfigNetworkConfigEgressOptionEnum.Name(
            resource
        )[
            len(
                "CloudbuildBetaWorkerPoolPrivatePoolV1ConfigNetworkConfigEgressOptionEnum"
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
