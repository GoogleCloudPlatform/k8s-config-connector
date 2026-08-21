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
from google3.cloud.graphite.mmv2.services.google.os_config import (
    os_policy_assignment_pb2,
)
from google3.cloud.graphite.mmv2.services.google.os_config import (
    os_policy_assignment_pb2_grpc,
)

from typing import List


class OSPolicyAssignment(object):
    def __init__(
        self,
        name: str = None,
        description: str = None,
        os_policies: list = None,
        instance_filter: dict = None,
        rollout: dict = None,
        revision_id: str = None,
        revision_create_time: str = None,
        etag: str = None,
        rollout_state: str = None,
        baseline: bool = None,
        deleted: bool = None,
        reconciling: bool = None,
        uid: str = None,
        project: str = None,
        location: str = None,
        skip_await_rollout: bool = None,
        service_account_file: str = "",
    ):
        channel.initialize()
        self.name = name
        self.description = description
        self.os_policies = os_policies
        self.instance_filter = instance_filter
        self.rollout = rollout
        self.project = project
        self.location = location
        self.skip_await_rollout = skip_await_rollout
        self.service_account_file = service_account_file

    def apply(self):
        stub = os_policy_assignment_pb2_grpc.OsconfigBetaOSPolicyAssignmentServiceStub(
            channel.Channel()
        )
        request = os_policy_assignment_pb2.ApplyOsconfigBetaOSPolicyAssignmentRequest()
        if Primitive.to_proto(self.name):
            request.resource.name = Primitive.to_proto(self.name)

        if Primitive.to_proto(self.description):
            request.resource.description = Primitive.to_proto(self.description)

        if OSPolicyAssignmentOSPoliciesArray.to_proto(self.os_policies):
            request.resource.os_policies.extend(
                OSPolicyAssignmentOSPoliciesArray.to_proto(self.os_policies)
            )
        if OSPolicyAssignmentInstanceFilter.to_proto(self.instance_filter):
            request.resource.instance_filter.CopyFrom(
                OSPolicyAssignmentInstanceFilter.to_proto(self.instance_filter)
            )
        else:
            request.resource.ClearField("instance_filter")
        if OSPolicyAssignmentRollout.to_proto(self.rollout):
            request.resource.rollout.CopyFrom(
                OSPolicyAssignmentRollout.to_proto(self.rollout)
            )
        else:
            request.resource.ClearField("rollout")
        if Primitive.to_proto(self.project):
            request.resource.project = Primitive.to_proto(self.project)

        if Primitive.to_proto(self.location):
            request.resource.location = Primitive.to_proto(self.location)

        if Primitive.to_proto(self.skip_await_rollout):
            request.resource.skip_await_rollout = Primitive.to_proto(
                self.skip_await_rollout
            )

        request.service_account_file = self.service_account_file

        response = stub.ApplyOsconfigBetaOSPolicyAssignment(request)
        self.name = Primitive.from_proto(response.name)
        self.description = Primitive.from_proto(response.description)
        self.os_policies = OSPolicyAssignmentOSPoliciesArray.from_proto(
            response.os_policies
        )
        self.instance_filter = OSPolicyAssignmentInstanceFilter.from_proto(
            response.instance_filter
        )
        self.rollout = OSPolicyAssignmentRollout.from_proto(response.rollout)
        self.revision_id = Primitive.from_proto(response.revision_id)
        self.revision_create_time = Primitive.from_proto(response.revision_create_time)
        self.etag = Primitive.from_proto(response.etag)
        self.rollout_state = OSPolicyAssignmentRolloutStateEnum.from_proto(
            response.rollout_state
        )
        self.baseline = Primitive.from_proto(response.baseline)
        self.deleted = Primitive.from_proto(response.deleted)
        self.reconciling = Primitive.from_proto(response.reconciling)
        self.uid = Primitive.from_proto(response.uid)
        self.project = Primitive.from_proto(response.project)
        self.location = Primitive.from_proto(response.location)
        self.skip_await_rollout = Primitive.from_proto(response.skip_await_rollout)

    def delete(self):
        stub = os_policy_assignment_pb2_grpc.OsconfigBetaOSPolicyAssignmentServiceStub(
            channel.Channel()
        )
        request = os_policy_assignment_pb2.DeleteOsconfigBetaOSPolicyAssignmentRequest()
        request.service_account_file = self.service_account_file
        if Primitive.to_proto(self.name):
            request.resource.name = Primitive.to_proto(self.name)

        if Primitive.to_proto(self.description):
            request.resource.description = Primitive.to_proto(self.description)

        if OSPolicyAssignmentOSPoliciesArray.to_proto(self.os_policies):
            request.resource.os_policies.extend(
                OSPolicyAssignmentOSPoliciesArray.to_proto(self.os_policies)
            )
        if OSPolicyAssignmentInstanceFilter.to_proto(self.instance_filter):
            request.resource.instance_filter.CopyFrom(
                OSPolicyAssignmentInstanceFilter.to_proto(self.instance_filter)
            )
        else:
            request.resource.ClearField("instance_filter")
        if OSPolicyAssignmentRollout.to_proto(self.rollout):
            request.resource.rollout.CopyFrom(
                OSPolicyAssignmentRollout.to_proto(self.rollout)
            )
        else:
            request.resource.ClearField("rollout")
        if Primitive.to_proto(self.project):
            request.resource.project = Primitive.to_proto(self.project)

        if Primitive.to_proto(self.location):
            request.resource.location = Primitive.to_proto(self.location)

        if Primitive.to_proto(self.skip_await_rollout):
            request.resource.skip_await_rollout = Primitive.to_proto(
                self.skip_await_rollout
            )

        response = stub.DeleteOsconfigBetaOSPolicyAssignment(request)

    @classmethod
    def list(self, project, location, service_account_file=""):
        stub = os_policy_assignment_pb2_grpc.OsconfigBetaOSPolicyAssignmentServiceStub(
            channel.Channel()
        )
        request = os_policy_assignment_pb2.ListOsconfigBetaOSPolicyAssignmentRequest()
        request.service_account_file = service_account_file
        request.Project = project

        request.Location = location

        return stub.ListOsconfigBetaOSPolicyAssignment(request).items

    def to_proto(self):
        resource = os_policy_assignment_pb2.OsconfigBetaOSPolicyAssignment()
        if Primitive.to_proto(self.name):
            resource.name = Primitive.to_proto(self.name)
        if Primitive.to_proto(self.description):
            resource.description = Primitive.to_proto(self.description)
        if OSPolicyAssignmentOSPoliciesArray.to_proto(self.os_policies):
            resource.os_policies.extend(
                OSPolicyAssignmentOSPoliciesArray.to_proto(self.os_policies)
            )
        if OSPolicyAssignmentInstanceFilter.to_proto(self.instance_filter):
            resource.instance_filter.CopyFrom(
                OSPolicyAssignmentInstanceFilter.to_proto(self.instance_filter)
            )
        else:
            resource.ClearField("instance_filter")
        if OSPolicyAssignmentRollout.to_proto(self.rollout):
            resource.rollout.CopyFrom(OSPolicyAssignmentRollout.to_proto(self.rollout))
        else:
            resource.ClearField("rollout")
        if Primitive.to_proto(self.project):
            resource.project = Primitive.to_proto(self.project)
        if Primitive.to_proto(self.location):
            resource.location = Primitive.to_proto(self.location)
        if Primitive.to_proto(self.skip_await_rollout):
            resource.skip_await_rollout = Primitive.to_proto(self.skip_await_rollout)
        return resource


class OSPolicyAssignmentOSPolicies(object):
    def __init__(
        self,
        id: str = None,
        description: str = None,
        mode: str = None,
        resource_groups: list = None,
        allow_no_resource_group_match: bool = None,
    ):
        self.id = id
        self.description = description
        self.mode = mode
        self.resource_groups = resource_groups
        self.allow_no_resource_group_match = allow_no_resource_group_match

    @classmethod
    def to_proto(self, resource):
        if not resource:
            return None

        res = os_policy_assignment_pb2.OsconfigBetaOSPolicyAssignmentOSPolicies()
        if Primitive.to_proto(resource.id):
            res.id = Primitive.to_proto(resource.id)
        if Primitive.to_proto(resource.description):
            res.description = Primitive.to_proto(resource.description)
        if OSPolicyAssignmentOSPoliciesModeEnum.to_proto(resource.mode):
            res.mode = OSPolicyAssignmentOSPoliciesModeEnum.to_proto(resource.mode)
        if OSPolicyAssignmentOSPoliciesResourceGroupsArray.to_proto(
            resource.resource_groups
        ):
            res.resource_groups.extend(
                OSPolicyAssignmentOSPoliciesResourceGroupsArray.to_proto(
                    resource.resource_groups
                )
            )
        if Primitive.to_proto(resource.allow_no_resource_group_match):
            res.allow_no_resource_group_match = Primitive.to_proto(
                resource.allow_no_resource_group_match
            )
        return res

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return None

        return OSPolicyAssignmentOSPolicies(
            id=Primitive.from_proto(resource.id),
            description=Primitive.from_proto(resource.description),
            mode=OSPolicyAssignmentOSPoliciesModeEnum.from_proto(resource.mode),
            resource_groups=OSPolicyAssignmentOSPoliciesResourceGroupsArray.from_proto(
                resource.resource_groups
            ),
            allow_no_resource_group_match=Primitive.from_proto(
                resource.allow_no_resource_group_match
            ),
        )


class OSPolicyAssignmentOSPoliciesArray(object):
    @classmethod
    def to_proto(self, resources):
        if not resources:
            return resources
        return [OSPolicyAssignmentOSPolicies.to_proto(i) for i in resources]

    @classmethod
    def from_proto(self, resources):
        return [OSPolicyAssignmentOSPolicies.from_proto(i) for i in resources]


class OSPolicyAssignmentOSPoliciesResourceGroups(object):
    def __init__(self, inventory_filters: list = None, resources: list = None):
        self.inventory_filters = inventory_filters
        self.resources = resources

    @classmethod
    def to_proto(self, resource):
        if not resource:
            return None

        res = (
            os_policy_assignment_pb2.OsconfigBetaOSPolicyAssignmentOSPoliciesResourceGroups()
        )
        if OSPolicyAssignmentOSPoliciesResourceGroupsInventoryFiltersArray.to_proto(
            resource.inventory_filters
        ):
            res.inventory_filters.extend(
                OSPolicyAssignmentOSPoliciesResourceGroupsInventoryFiltersArray.to_proto(
                    resource.inventory_filters
                )
            )
        if OSPolicyAssignmentOSPoliciesResourceGroupsResourcesArray.to_proto(
            resource.resources
        ):
            res.resources.extend(
                OSPolicyAssignmentOSPoliciesResourceGroupsResourcesArray.to_proto(
                    resource.resources
                )
            )
        return res

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return None

        return OSPolicyAssignmentOSPoliciesResourceGroups(
            inventory_filters=OSPolicyAssignmentOSPoliciesResourceGroupsInventoryFiltersArray.from_proto(
                resource.inventory_filters
            ),
            resources=OSPolicyAssignmentOSPoliciesResourceGroupsResourcesArray.from_proto(
                resource.resources
            ),
        )


class OSPolicyAssignmentOSPoliciesResourceGroupsArray(object):
    @classmethod
    def to_proto(self, resources):
        if not resources:
            return resources
        return [
            OSPolicyAssignmentOSPoliciesResourceGroups.to_proto(i) for i in resources
        ]

    @classmethod
    def from_proto(self, resources):
        return [
            OSPolicyAssignmentOSPoliciesResourceGroups.from_proto(i) for i in resources
        ]


class OSPolicyAssignmentOSPoliciesResourceGroupsInventoryFilters(object):
    def __init__(self, os_short_name: str = None, os_version: str = None):
        self.os_short_name = os_short_name
        self.os_version = os_version

    @classmethod
    def to_proto(self, resource):
        if not resource:
            return None

        res = (
            os_policy_assignment_pb2.OsconfigBetaOSPolicyAssignmentOSPoliciesResourceGroupsInventoryFilters()
        )
        if Primitive.to_proto(resource.os_short_name):
            res.os_short_name = Primitive.to_proto(resource.os_short_name)
        if Primitive.to_proto(resource.os_version):
            res.os_version = Primitive.to_proto(resource.os_version)
        return res

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return None

        return OSPolicyAssignmentOSPoliciesResourceGroupsInventoryFilters(
            os_short_name=Primitive.from_proto(resource.os_short_name),
            os_version=Primitive.from_proto(resource.os_version),
        )


class OSPolicyAssignmentOSPoliciesResourceGroupsInventoryFiltersArray(object):
    @classmethod
    def to_proto(self, resources):
        if not resources:
            return resources
        return [
            OSPolicyAssignmentOSPoliciesResourceGroupsInventoryFilters.to_proto(i)
            for i in resources
        ]

    @classmethod
    def from_proto(self, resources):
        return [
            OSPolicyAssignmentOSPoliciesResourceGroupsInventoryFilters.from_proto(i)
            for i in resources
        ]


class OSPolicyAssignmentOSPoliciesResourceGroupsResources(object):
    def __init__(
        self,
        id: str = None,
        pkg: dict = None,
        repository: dict = None,
        exec: dict = None,
        file: dict = None,
    ):
        self.id = id
        self.pkg = pkg
        self.repository = repository
        self.exec = exec
        self.file = file

    @classmethod
    def to_proto(self, resource):
        if not resource:
            return None

        res = (
            os_policy_assignment_pb2.OsconfigBetaOSPolicyAssignmentOSPoliciesResourceGroupsResources()
        )
        if Primitive.to_proto(resource.id):
            res.id = Primitive.to_proto(resource.id)
        if OSPolicyAssignmentOSPoliciesResourceGroupsResourcesPkg.to_proto(
            resource.pkg
        ):
            res.pkg.CopyFrom(
                OSPolicyAssignmentOSPoliciesResourceGroupsResourcesPkg.to_proto(
                    resource.pkg
                )
            )
        else:
            res.ClearField("pkg")
        if OSPolicyAssignmentOSPoliciesResourceGroupsResourcesRepository.to_proto(
            resource.repository
        ):
            res.repository.CopyFrom(
                OSPolicyAssignmentOSPoliciesResourceGroupsResourcesRepository.to_proto(
                    resource.repository
                )
            )
        else:
            res.ClearField("repository")
        if OSPolicyAssignmentOSPoliciesResourceGroupsResourcesExec.to_proto(
            resource.exec
        ):
            res.exec.CopyFrom(
                OSPolicyAssignmentOSPoliciesResourceGroupsResourcesExec.to_proto(
                    resource.exec
                )
            )
        else:
            res.ClearField("exec")
        if OSPolicyAssignmentOSPoliciesResourceGroupsResourcesFile.to_proto(
            resource.file
        ):
            res.file.CopyFrom(
                OSPolicyAssignmentOSPoliciesResourceGroupsResourcesFile.to_proto(
                    resource.file
                )
            )
        else:
            res.ClearField("file")
        return res

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return None

        return OSPolicyAssignmentOSPoliciesResourceGroupsResources(
            id=Primitive.from_proto(resource.id),
            pkg=OSPolicyAssignmentOSPoliciesResourceGroupsResourcesPkg.from_proto(
                resource.pkg
            ),
            repository=OSPolicyAssignmentOSPoliciesResourceGroupsResourcesRepository.from_proto(
                resource.repository
            ),
            exec=OSPolicyAssignmentOSPoliciesResourceGroupsResourcesExec.from_proto(
                resource.exec
            ),
            file=OSPolicyAssignmentOSPoliciesResourceGroupsResourcesFile.from_proto(
                resource.file
            ),
        )


class OSPolicyAssignmentOSPoliciesResourceGroupsResourcesArray(object):
    @classmethod
    def to_proto(self, resources):
        if not resources:
            return resources
        return [
            OSPolicyAssignmentOSPoliciesResourceGroupsResources.to_proto(i)
            for i in resources
        ]

    @classmethod
    def from_proto(self, resources):
        return [
            OSPolicyAssignmentOSPoliciesResourceGroupsResources.from_proto(i)
            for i in resources
        ]


class OSPolicyAssignmentOSPoliciesResourceGroupsResourcesPkg(object):
    def __init__(
        self,
        desired_state: str = None,
        apt: dict = None,
        deb: dict = None,
        yum: dict = None,
        zypper: dict = None,
        rpm: dict = None,
        googet: dict = None,
        msi: dict = None,
    ):
        self.desired_state = desired_state
        self.apt = apt
        self.deb = deb
        self.yum = yum
        self.zypper = zypper
        self.rpm = rpm
        self.googet = googet
        self.msi = msi

    @classmethod
    def to_proto(self, resource):
        if not resource:
            return None

        res = (
            os_policy_assignment_pb2.OsconfigBetaOSPolicyAssignmentOSPoliciesResourceGroupsResourcesPkg()
        )
        if OSPolicyAssignmentOSPoliciesResourceGroupsResourcesPkgDesiredStateEnum.to_proto(
            resource.desired_state
        ):
            res.desired_state = OSPolicyAssignmentOSPoliciesResourceGroupsResourcesPkgDesiredStateEnum.to_proto(
                resource.desired_state
            )
        if OSPolicyAssignmentOSPoliciesResourceGroupsResourcesPkgApt.to_proto(
            resource.apt
        ):
            res.apt.CopyFrom(
                OSPolicyAssignmentOSPoliciesResourceGroupsResourcesPkgApt.to_proto(
                    resource.apt
                )
            )
        else:
            res.ClearField("apt")
        if OSPolicyAssignmentOSPoliciesResourceGroupsResourcesPkgDeb.to_proto(
            resource.deb
        ):
            res.deb.CopyFrom(
                OSPolicyAssignmentOSPoliciesResourceGroupsResourcesPkgDeb.to_proto(
                    resource.deb
                )
            )
        else:
            res.ClearField("deb")
        if OSPolicyAssignmentOSPoliciesResourceGroupsResourcesPkgYum.to_proto(
            resource.yum
        ):
            res.yum.CopyFrom(
                OSPolicyAssignmentOSPoliciesResourceGroupsResourcesPkgYum.to_proto(
                    resource.yum
                )
            )
        else:
            res.ClearField("yum")
        if OSPolicyAssignmentOSPoliciesResourceGroupsResourcesPkgZypper.to_proto(
            resource.zypper
        ):
            res.zypper.CopyFrom(
                OSPolicyAssignmentOSPoliciesResourceGroupsResourcesPkgZypper.to_proto(
                    resource.zypper
                )
            )
        else:
            res.ClearField("zypper")
        if OSPolicyAssignmentOSPoliciesResourceGroupsResourcesPkgRpm.to_proto(
            resource.rpm
        ):
            res.rpm.CopyFrom(
                OSPolicyAssignmentOSPoliciesResourceGroupsResourcesPkgRpm.to_proto(
                    resource.rpm
                )
            )
        else:
            res.ClearField("rpm")
        if OSPolicyAssignmentOSPoliciesResourceGroupsResourcesPkgGooget.to_proto(
            resource.googet
        ):
            res.googet.CopyFrom(
                OSPolicyAssignmentOSPoliciesResourceGroupsResourcesPkgGooget.to_proto(
                    resource.googet
                )
            )
        else:
            res.ClearField("googet")
        if OSPolicyAssignmentOSPoliciesResourceGroupsResourcesPkgMsi.to_proto(
            resource.msi
        ):
            res.msi.CopyFrom(
                OSPolicyAssignmentOSPoliciesResourceGroupsResourcesPkgMsi.to_proto(
                    resource.msi
                )
            )
        else:
            res.ClearField("msi")
        return res

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return None

        return OSPolicyAssignmentOSPoliciesResourceGroupsResourcesPkg(
            desired_state=OSPolicyAssignmentOSPoliciesResourceGroupsResourcesPkgDesiredStateEnum.from_proto(
                resource.desired_state
            ),
            apt=OSPolicyAssignmentOSPoliciesResourceGroupsResourcesPkgApt.from_proto(
                resource.apt
            ),
            deb=OSPolicyAssignmentOSPoliciesResourceGroupsResourcesPkgDeb.from_proto(
                resource.deb
            ),
            yum=OSPolicyAssignmentOSPoliciesResourceGroupsResourcesPkgYum.from_proto(
                resource.yum
            ),
            zypper=OSPolicyAssignmentOSPoliciesResourceGroupsResourcesPkgZypper.from_proto(
                resource.zypper
            ),
            rpm=OSPolicyAssignmentOSPoliciesResourceGroupsResourcesPkgRpm.from_proto(
                resource.rpm
            ),
            googet=OSPolicyAssignmentOSPoliciesResourceGroupsResourcesPkgGooget.from_proto(
                resource.googet
            ),
            msi=OSPolicyAssignmentOSPoliciesResourceGroupsResourcesPkgMsi.from_proto(
                resource.msi
            ),
        )


class OSPolicyAssignmentOSPoliciesResourceGroupsResourcesPkgArray(object):
    @classmethod
    def to_proto(self, resources):
        if not resources:
            return resources
        return [
            OSPolicyAssignmentOSPoliciesResourceGroupsResourcesPkg.to_proto(i)
            for i in resources
        ]

    @classmethod
    def from_proto(self, resources):
        return [
            OSPolicyAssignmentOSPoliciesResourceGroupsResourcesPkg.from_proto(i)
            for i in resources
        ]


class OSPolicyAssignmentOSPoliciesResourceGroupsResourcesPkgApt(object):
    def __init__(self, name: str = None):
        self.name = name

    @classmethod
    def to_proto(self, resource):
        if not resource:
            return None

        res = (
            os_policy_assignment_pb2.OsconfigBetaOSPolicyAssignmentOSPoliciesResourceGroupsResourcesPkgApt()
        )
        if Primitive.to_proto(resource.name):
            res.name = Primitive.to_proto(resource.name)
        return res

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return None

        return OSPolicyAssignmentOSPoliciesResourceGroupsResourcesPkgApt(
            name=Primitive.from_proto(resource.name),
        )


class OSPolicyAssignmentOSPoliciesResourceGroupsResourcesPkgAptArray(object):
    @classmethod
    def to_proto(self, resources):
        if not resources:
            return resources
        return [
            OSPolicyAssignmentOSPoliciesResourceGroupsResourcesPkgApt.to_proto(i)
            for i in resources
        ]

    @classmethod
    def from_proto(self, resources):
        return [
            OSPolicyAssignmentOSPoliciesResourceGroupsResourcesPkgApt.from_proto(i)
            for i in resources
        ]


class OSPolicyAssignmentOSPoliciesResourceGroupsResourcesPkgDeb(object):
    def __init__(self, source: dict = None, pull_deps: bool = None):
        self.source = source
        self.pull_deps = pull_deps

    @classmethod
    def to_proto(self, resource):
        if not resource:
            return None

        res = (
            os_policy_assignment_pb2.OsconfigBetaOSPolicyAssignmentOSPoliciesResourceGroupsResourcesPkgDeb()
        )
        if OSPolicyAssignmentOSPoliciesResourceGroupsResourcesPkgDebSource.to_proto(
            resource.source
        ):
            res.source.CopyFrom(
                OSPolicyAssignmentOSPoliciesResourceGroupsResourcesPkgDebSource.to_proto(
                    resource.source
                )
            )
        else:
            res.ClearField("source")
        if Primitive.to_proto(resource.pull_deps):
            res.pull_deps = Primitive.to_proto(resource.pull_deps)
        return res

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return None

        return OSPolicyAssignmentOSPoliciesResourceGroupsResourcesPkgDeb(
            source=OSPolicyAssignmentOSPoliciesResourceGroupsResourcesPkgDebSource.from_proto(
                resource.source
            ),
            pull_deps=Primitive.from_proto(resource.pull_deps),
        )


class OSPolicyAssignmentOSPoliciesResourceGroupsResourcesPkgDebArray(object):
    @classmethod
    def to_proto(self, resources):
        if not resources:
            return resources
        return [
            OSPolicyAssignmentOSPoliciesResourceGroupsResourcesPkgDeb.to_proto(i)
            for i in resources
        ]

    @classmethod
    def from_proto(self, resources):
        return [
            OSPolicyAssignmentOSPoliciesResourceGroupsResourcesPkgDeb.from_proto(i)
            for i in resources
        ]


class OSPolicyAssignmentOSPoliciesResourceGroupsResourcesPkgDebSource(object):
    def __init__(
        self,
        remote: dict = None,
        gcs: dict = None,
        local_path: str = None,
        allow_insecure: bool = None,
    ):
        self.remote = remote
        self.gcs = gcs
        self.local_path = local_path
        self.allow_insecure = allow_insecure

    @classmethod
    def to_proto(self, resource):
        if not resource:
            return None

        res = (
            os_policy_assignment_pb2.OsconfigBetaOSPolicyAssignmentOSPoliciesResourceGroupsResourcesPkgDebSource()
        )
        if OSPolicyAssignmentOSPoliciesResourceGroupsResourcesPkgDebSourceRemote.to_proto(
            resource.remote
        ):
            res.remote.CopyFrom(
                OSPolicyAssignmentOSPoliciesResourceGroupsResourcesPkgDebSourceRemote.to_proto(
                    resource.remote
                )
            )
        else:
            res.ClearField("remote")
        if OSPolicyAssignmentOSPoliciesResourceGroupsResourcesPkgDebSourceGcs.to_proto(
            resource.gcs
        ):
            res.gcs.CopyFrom(
                OSPolicyAssignmentOSPoliciesResourceGroupsResourcesPkgDebSourceGcs.to_proto(
                    resource.gcs
                )
            )
        else:
            res.ClearField("gcs")
        if Primitive.to_proto(resource.local_path):
            res.local_path = Primitive.to_proto(resource.local_path)
        if Primitive.to_proto(resource.allow_insecure):
            res.allow_insecure = Primitive.to_proto(resource.allow_insecure)
        return res

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return None

        return OSPolicyAssignmentOSPoliciesResourceGroupsResourcesPkgDebSource(
            remote=OSPolicyAssignmentOSPoliciesResourceGroupsResourcesPkgDebSourceRemote.from_proto(
                resource.remote
            ),
            gcs=OSPolicyAssignmentOSPoliciesResourceGroupsResourcesPkgDebSourceGcs.from_proto(
                resource.gcs
            ),
            local_path=Primitive.from_proto(resource.local_path),
            allow_insecure=Primitive.from_proto(resource.allow_insecure),
        )


class OSPolicyAssignmentOSPoliciesResourceGroupsResourcesPkgDebSourceArray(object):
    @classmethod
    def to_proto(self, resources):
        if not resources:
            return resources
        return [
            OSPolicyAssignmentOSPoliciesResourceGroupsResourcesPkgDebSource.to_proto(i)
            for i in resources
        ]

    @classmethod
    def from_proto(self, resources):
        return [
            OSPolicyAssignmentOSPoliciesResourceGroupsResourcesPkgDebSource.from_proto(
                i
            )
            for i in resources
        ]


class OSPolicyAssignmentOSPoliciesResourceGroupsResourcesPkgDebSourceRemote(object):
    def __init__(self, uri: str = None, sha256_checksum: str = None):
        self.uri = uri
        self.sha256_checksum = sha256_checksum

    @classmethod
    def to_proto(self, resource):
        if not resource:
            return None

        res = (
            os_policy_assignment_pb2.OsconfigBetaOSPolicyAssignmentOSPoliciesResourceGroupsResourcesPkgDebSourceRemote()
        )
        if Primitive.to_proto(resource.uri):
            res.uri = Primitive.to_proto(resource.uri)
        if Primitive.to_proto(resource.sha256_checksum):
            res.sha256_checksum = Primitive.to_proto(resource.sha256_checksum)
        return res

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return None

        return OSPolicyAssignmentOSPoliciesResourceGroupsResourcesPkgDebSourceRemote(
            uri=Primitive.from_proto(resource.uri),
            sha256_checksum=Primitive.from_proto(resource.sha256_checksum),
        )


class OSPolicyAssignmentOSPoliciesResourceGroupsResourcesPkgDebSourceRemoteArray(
    object
):
    @classmethod
    def to_proto(self, resources):
        if not resources:
            return resources
        return [
            OSPolicyAssignmentOSPoliciesResourceGroupsResourcesPkgDebSourceRemote.to_proto(
                i
            )
            for i in resources
        ]

    @classmethod
    def from_proto(self, resources):
        return [
            OSPolicyAssignmentOSPoliciesResourceGroupsResourcesPkgDebSourceRemote.from_proto(
                i
            )
            for i in resources
        ]


class OSPolicyAssignmentOSPoliciesResourceGroupsResourcesPkgDebSourceGcs(object):
    def __init__(self, bucket: str = None, object: str = None, generation: int = None):
        self.bucket = bucket
        self.object = object
        self.generation = generation

    @classmethod
    def to_proto(self, resource):
        if not resource:
            return None

        res = (
            os_policy_assignment_pb2.OsconfigBetaOSPolicyAssignmentOSPoliciesResourceGroupsResourcesPkgDebSourceGcs()
        )
        if Primitive.to_proto(resource.bucket):
            res.bucket = Primitive.to_proto(resource.bucket)
        if Primitive.to_proto(resource.object):
            res.object = Primitive.to_proto(resource.object)
        if Primitive.to_proto(resource.generation):
            res.generation = Primitive.to_proto(resource.generation)
        return res

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return None

        return OSPolicyAssignmentOSPoliciesResourceGroupsResourcesPkgDebSourceGcs(
            bucket=Primitive.from_proto(resource.bucket),
            object=Primitive.from_proto(resource.object),
            generation=Primitive.from_proto(resource.generation),
        )


class OSPolicyAssignmentOSPoliciesResourceGroupsResourcesPkgDebSourceGcsArray(object):
    @classmethod
    def to_proto(self, resources):
        if not resources:
            return resources
        return [
            OSPolicyAssignmentOSPoliciesResourceGroupsResourcesPkgDebSourceGcs.to_proto(
                i
            )
            for i in resources
        ]

    @classmethod
    def from_proto(self, resources):
        return [
            OSPolicyAssignmentOSPoliciesResourceGroupsResourcesPkgDebSourceGcs.from_proto(
                i
            )
            for i in resources
        ]


class OSPolicyAssignmentOSPoliciesResourceGroupsResourcesPkgYum(object):
    def __init__(self, name: str = None):
        self.name = name

    @classmethod
    def to_proto(self, resource):
        if not resource:
            return None

        res = (
            os_policy_assignment_pb2.OsconfigBetaOSPolicyAssignmentOSPoliciesResourceGroupsResourcesPkgYum()
        )
        if Primitive.to_proto(resource.name):
            res.name = Primitive.to_proto(resource.name)
        return res

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return None

        return OSPolicyAssignmentOSPoliciesResourceGroupsResourcesPkgYum(
            name=Primitive.from_proto(resource.name),
        )


class OSPolicyAssignmentOSPoliciesResourceGroupsResourcesPkgYumArray(object):
    @classmethod
    def to_proto(self, resources):
        if not resources:
            return resources
        return [
            OSPolicyAssignmentOSPoliciesResourceGroupsResourcesPkgYum.to_proto(i)
            for i in resources
        ]

    @classmethod
    def from_proto(self, resources):
        return [
            OSPolicyAssignmentOSPoliciesResourceGroupsResourcesPkgYum.from_proto(i)
            for i in resources
        ]


class OSPolicyAssignmentOSPoliciesResourceGroupsResourcesPkgZypper(object):
    def __init__(self, name: str = None):
        self.name = name

    @classmethod
    def to_proto(self, resource):
        if not resource:
            return None

        res = (
            os_policy_assignment_pb2.OsconfigBetaOSPolicyAssignmentOSPoliciesResourceGroupsResourcesPkgZypper()
        )
        if Primitive.to_proto(resource.name):
            res.name = Primitive.to_proto(resource.name)
        return res

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return None

        return OSPolicyAssignmentOSPoliciesResourceGroupsResourcesPkgZypper(
            name=Primitive.from_proto(resource.name),
        )


class OSPolicyAssignmentOSPoliciesResourceGroupsResourcesPkgZypperArray(object):
    @classmethod
    def to_proto(self, resources):
        if not resources:
            return resources
        return [
            OSPolicyAssignmentOSPoliciesResourceGroupsResourcesPkgZypper.to_proto(i)
            for i in resources
        ]

    @classmethod
    def from_proto(self, resources):
        return [
            OSPolicyAssignmentOSPoliciesResourceGroupsResourcesPkgZypper.from_proto(i)
            for i in resources
        ]


class OSPolicyAssignmentOSPoliciesResourceGroupsResourcesPkgRpm(object):
    def __init__(self, source: dict = None, pull_deps: bool = None):
        self.source = source
        self.pull_deps = pull_deps

    @classmethod
    def to_proto(self, resource):
        if not resource:
            return None

        res = (
            os_policy_assignment_pb2.OsconfigBetaOSPolicyAssignmentOSPoliciesResourceGroupsResourcesPkgRpm()
        )
        if OSPolicyAssignmentOSPoliciesResourceGroupsResourcesPkgRpmSource.to_proto(
            resource.source
        ):
            res.source.CopyFrom(
                OSPolicyAssignmentOSPoliciesResourceGroupsResourcesPkgRpmSource.to_proto(
                    resource.source
                )
            )
        else:
            res.ClearField("source")
        if Primitive.to_proto(resource.pull_deps):
            res.pull_deps = Primitive.to_proto(resource.pull_deps)
        return res

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return None

        return OSPolicyAssignmentOSPoliciesResourceGroupsResourcesPkgRpm(
            source=OSPolicyAssignmentOSPoliciesResourceGroupsResourcesPkgRpmSource.from_proto(
                resource.source
            ),
            pull_deps=Primitive.from_proto(resource.pull_deps),
        )


class OSPolicyAssignmentOSPoliciesResourceGroupsResourcesPkgRpmArray(object):
    @classmethod
    def to_proto(self, resources):
        if not resources:
            return resources
        return [
            OSPolicyAssignmentOSPoliciesResourceGroupsResourcesPkgRpm.to_proto(i)
            for i in resources
        ]

    @classmethod
    def from_proto(self, resources):
        return [
            OSPolicyAssignmentOSPoliciesResourceGroupsResourcesPkgRpm.from_proto(i)
            for i in resources
        ]


class OSPolicyAssignmentOSPoliciesResourceGroupsResourcesPkgRpmSource(object):
    def __init__(
        self,
        remote: dict = None,
        gcs: dict = None,
        local_path: str = None,
        allow_insecure: bool = None,
    ):
        self.remote = remote
        self.gcs = gcs
        self.local_path = local_path
        self.allow_insecure = allow_insecure

    @classmethod
    def to_proto(self, resource):
        if not resource:
            return None

        res = (
            os_policy_assignment_pb2.OsconfigBetaOSPolicyAssignmentOSPoliciesResourceGroupsResourcesPkgRpmSource()
        )
        if OSPolicyAssignmentOSPoliciesResourceGroupsResourcesPkgRpmSourceRemote.to_proto(
            resource.remote
        ):
            res.remote.CopyFrom(
                OSPolicyAssignmentOSPoliciesResourceGroupsResourcesPkgRpmSourceRemote.to_proto(
                    resource.remote
                )
            )
        else:
            res.ClearField("remote")
        if OSPolicyAssignmentOSPoliciesResourceGroupsResourcesPkgRpmSourceGcs.to_proto(
            resource.gcs
        ):
            res.gcs.CopyFrom(
                OSPolicyAssignmentOSPoliciesResourceGroupsResourcesPkgRpmSourceGcs.to_proto(
                    resource.gcs
                )
            )
        else:
            res.ClearField("gcs")
        if Primitive.to_proto(resource.local_path):
            res.local_path = Primitive.to_proto(resource.local_path)
        if Primitive.to_proto(resource.allow_insecure):
            res.allow_insecure = Primitive.to_proto(resource.allow_insecure)
        return res

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return None

        return OSPolicyAssignmentOSPoliciesResourceGroupsResourcesPkgRpmSource(
            remote=OSPolicyAssignmentOSPoliciesResourceGroupsResourcesPkgRpmSourceRemote.from_proto(
                resource.remote
            ),
            gcs=OSPolicyAssignmentOSPoliciesResourceGroupsResourcesPkgRpmSourceGcs.from_proto(
                resource.gcs
            ),
            local_path=Primitive.from_proto(resource.local_path),
            allow_insecure=Primitive.from_proto(resource.allow_insecure),
        )


class OSPolicyAssignmentOSPoliciesResourceGroupsResourcesPkgRpmSourceArray(object):
    @classmethod
    def to_proto(self, resources):
        if not resources:
            return resources
        return [
            OSPolicyAssignmentOSPoliciesResourceGroupsResourcesPkgRpmSource.to_proto(i)
            for i in resources
        ]

    @classmethod
    def from_proto(self, resources):
        return [
            OSPolicyAssignmentOSPoliciesResourceGroupsResourcesPkgRpmSource.from_proto(
                i
            )
            for i in resources
        ]


class OSPolicyAssignmentOSPoliciesResourceGroupsResourcesPkgRpmSourceRemote(object):
    def __init__(self, uri: str = None, sha256_checksum: str = None):
        self.uri = uri
        self.sha256_checksum = sha256_checksum

    @classmethod
    def to_proto(self, resource):
        if not resource:
            return None

        res = (
            os_policy_assignment_pb2.OsconfigBetaOSPolicyAssignmentOSPoliciesResourceGroupsResourcesPkgRpmSourceRemote()
        )
        if Primitive.to_proto(resource.uri):
            res.uri = Primitive.to_proto(resource.uri)
        if Primitive.to_proto(resource.sha256_checksum):
            res.sha256_checksum = Primitive.to_proto(resource.sha256_checksum)
        return res

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return None

        return OSPolicyAssignmentOSPoliciesResourceGroupsResourcesPkgRpmSourceRemote(
            uri=Primitive.from_proto(resource.uri),
            sha256_checksum=Primitive.from_proto(resource.sha256_checksum),
        )


class OSPolicyAssignmentOSPoliciesResourceGroupsResourcesPkgRpmSourceRemoteArray(
    object
):
    @classmethod
    def to_proto(self, resources):
        if not resources:
            return resources
        return [
            OSPolicyAssignmentOSPoliciesResourceGroupsResourcesPkgRpmSourceRemote.to_proto(
                i
            )
            for i in resources
        ]

    @classmethod
    def from_proto(self, resources):
        return [
            OSPolicyAssignmentOSPoliciesResourceGroupsResourcesPkgRpmSourceRemote.from_proto(
                i
            )
            for i in resources
        ]


class OSPolicyAssignmentOSPoliciesResourceGroupsResourcesPkgRpmSourceGcs(object):
    def __init__(self, bucket: str = None, object: str = None, generation: int = None):
        self.bucket = bucket
        self.object = object
        self.generation = generation

    @classmethod
    def to_proto(self, resource):
        if not resource:
            return None

        res = (
            os_policy_assignment_pb2.OsconfigBetaOSPolicyAssignmentOSPoliciesResourceGroupsResourcesPkgRpmSourceGcs()
        )
        if Primitive.to_proto(resource.bucket):
            res.bucket = Primitive.to_proto(resource.bucket)
        if Primitive.to_proto(resource.object):
            res.object = Primitive.to_proto(resource.object)
        if Primitive.to_proto(resource.generation):
            res.generation = Primitive.to_proto(resource.generation)
        return res

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return None

        return OSPolicyAssignmentOSPoliciesResourceGroupsResourcesPkgRpmSourceGcs(
            bucket=Primitive.from_proto(resource.bucket),
            object=Primitive.from_proto(resource.object),
            generation=Primitive.from_proto(resource.generation),
        )


class OSPolicyAssignmentOSPoliciesResourceGroupsResourcesPkgRpmSourceGcsArray(object):
    @classmethod
    def to_proto(self, resources):
        if not resources:
            return resources
        return [
            OSPolicyAssignmentOSPoliciesResourceGroupsResourcesPkgRpmSourceGcs.to_proto(
                i
            )
            for i in resources
        ]

    @classmethod
    def from_proto(self, resources):
        return [
            OSPolicyAssignmentOSPoliciesResourceGroupsResourcesPkgRpmSourceGcs.from_proto(
                i
            )
            for i in resources
        ]


class OSPolicyAssignmentOSPoliciesResourceGroupsResourcesPkgGooget(object):
    def __init__(self, name: str = None):
        self.name = name

    @classmethod
    def to_proto(self, resource):
        if not resource:
            return None

        res = (
            os_policy_assignment_pb2.OsconfigBetaOSPolicyAssignmentOSPoliciesResourceGroupsResourcesPkgGooget()
        )
        if Primitive.to_proto(resource.name):
            res.name = Primitive.to_proto(resource.name)
        return res

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return None

        return OSPolicyAssignmentOSPoliciesResourceGroupsResourcesPkgGooget(
            name=Primitive.from_proto(resource.name),
        )


class OSPolicyAssignmentOSPoliciesResourceGroupsResourcesPkgGoogetArray(object):
    @classmethod
    def to_proto(self, resources):
        if not resources:
            return resources
        return [
            OSPolicyAssignmentOSPoliciesResourceGroupsResourcesPkgGooget.to_proto(i)
            for i in resources
        ]

    @classmethod
    def from_proto(self, resources):
        return [
            OSPolicyAssignmentOSPoliciesResourceGroupsResourcesPkgGooget.from_proto(i)
            for i in resources
        ]


class OSPolicyAssignmentOSPoliciesResourceGroupsResourcesPkgMsi(object):
    def __init__(self, source: dict = None, properties: list = None):
        self.source = source
        self.properties = properties

    @classmethod
    def to_proto(self, resource):
        if not resource:
            return None

        res = (
            os_policy_assignment_pb2.OsconfigBetaOSPolicyAssignmentOSPoliciesResourceGroupsResourcesPkgMsi()
        )
        if OSPolicyAssignmentOSPoliciesResourceGroupsResourcesPkgMsiSource.to_proto(
            resource.source
        ):
            res.source.CopyFrom(
                OSPolicyAssignmentOSPoliciesResourceGroupsResourcesPkgMsiSource.to_proto(
                    resource.source
                )
            )
        else:
            res.ClearField("source")
        if Primitive.to_proto(resource.properties):
            res.properties.extend(Primitive.to_proto(resource.properties))
        return res

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return None

        return OSPolicyAssignmentOSPoliciesResourceGroupsResourcesPkgMsi(
            source=OSPolicyAssignmentOSPoliciesResourceGroupsResourcesPkgMsiSource.from_proto(
                resource.source
            ),
            properties=Primitive.from_proto(resource.properties),
        )


class OSPolicyAssignmentOSPoliciesResourceGroupsResourcesPkgMsiArray(object):
    @classmethod
    def to_proto(self, resources):
        if not resources:
            return resources
        return [
            OSPolicyAssignmentOSPoliciesResourceGroupsResourcesPkgMsi.to_proto(i)
            for i in resources
        ]

    @classmethod
    def from_proto(self, resources):
        return [
            OSPolicyAssignmentOSPoliciesResourceGroupsResourcesPkgMsi.from_proto(i)
            for i in resources
        ]


class OSPolicyAssignmentOSPoliciesResourceGroupsResourcesPkgMsiSource(object):
    def __init__(
        self,
        remote: dict = None,
        gcs: dict = None,
        local_path: str = None,
        allow_insecure: bool = None,
    ):
        self.remote = remote
        self.gcs = gcs
        self.local_path = local_path
        self.allow_insecure = allow_insecure

    @classmethod
    def to_proto(self, resource):
        if not resource:
            return None

        res = (
            os_policy_assignment_pb2.OsconfigBetaOSPolicyAssignmentOSPoliciesResourceGroupsResourcesPkgMsiSource()
        )
        if OSPolicyAssignmentOSPoliciesResourceGroupsResourcesPkgMsiSourceRemote.to_proto(
            resource.remote
        ):
            res.remote.CopyFrom(
                OSPolicyAssignmentOSPoliciesResourceGroupsResourcesPkgMsiSourceRemote.to_proto(
                    resource.remote
                )
            )
        else:
            res.ClearField("remote")
        if OSPolicyAssignmentOSPoliciesResourceGroupsResourcesPkgMsiSourceGcs.to_proto(
            resource.gcs
        ):
            res.gcs.CopyFrom(
                OSPolicyAssignmentOSPoliciesResourceGroupsResourcesPkgMsiSourceGcs.to_proto(
                    resource.gcs
                )
            )
        else:
            res.ClearField("gcs")
        if Primitive.to_proto(resource.local_path):
            res.local_path = Primitive.to_proto(resource.local_path)
        if Primitive.to_proto(resource.allow_insecure):
            res.allow_insecure = Primitive.to_proto(resource.allow_insecure)
        return res

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return None

        return OSPolicyAssignmentOSPoliciesResourceGroupsResourcesPkgMsiSource(
            remote=OSPolicyAssignmentOSPoliciesResourceGroupsResourcesPkgMsiSourceRemote.from_proto(
                resource.remote
            ),
            gcs=OSPolicyAssignmentOSPoliciesResourceGroupsResourcesPkgMsiSourceGcs.from_proto(
                resource.gcs
            ),
            local_path=Primitive.from_proto(resource.local_path),
            allow_insecure=Primitive.from_proto(resource.allow_insecure),
        )


class OSPolicyAssignmentOSPoliciesResourceGroupsResourcesPkgMsiSourceArray(object):
    @classmethod
    def to_proto(self, resources):
        if not resources:
            return resources
        return [
            OSPolicyAssignmentOSPoliciesResourceGroupsResourcesPkgMsiSource.to_proto(i)
            for i in resources
        ]

    @classmethod
    def from_proto(self, resources):
        return [
            OSPolicyAssignmentOSPoliciesResourceGroupsResourcesPkgMsiSource.from_proto(
                i
            )
            for i in resources
        ]


class OSPolicyAssignmentOSPoliciesResourceGroupsResourcesPkgMsiSourceRemote(object):
    def __init__(self, uri: str = None, sha256_checksum: str = None):
        self.uri = uri
        self.sha256_checksum = sha256_checksum

    @classmethod
    def to_proto(self, resource):
        if not resource:
            return None

        res = (
            os_policy_assignment_pb2.OsconfigBetaOSPolicyAssignmentOSPoliciesResourceGroupsResourcesPkgMsiSourceRemote()
        )
        if Primitive.to_proto(resource.uri):
            res.uri = Primitive.to_proto(resource.uri)
        if Primitive.to_proto(resource.sha256_checksum):
            res.sha256_checksum = Primitive.to_proto(resource.sha256_checksum)
        return res

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return None

        return OSPolicyAssignmentOSPoliciesResourceGroupsResourcesPkgMsiSourceRemote(
            uri=Primitive.from_proto(resource.uri),
            sha256_checksum=Primitive.from_proto(resource.sha256_checksum),
        )


class OSPolicyAssignmentOSPoliciesResourceGroupsResourcesPkgMsiSourceRemoteArray(
    object
):
    @classmethod
    def to_proto(self, resources):
        if not resources:
            return resources
        return [
            OSPolicyAssignmentOSPoliciesResourceGroupsResourcesPkgMsiSourceRemote.to_proto(
                i
            )
            for i in resources
        ]

    @classmethod
    def from_proto(self, resources):
        return [
            OSPolicyAssignmentOSPoliciesResourceGroupsResourcesPkgMsiSourceRemote.from_proto(
                i
            )
            for i in resources
        ]


class OSPolicyAssignmentOSPoliciesResourceGroupsResourcesPkgMsiSourceGcs(object):
    def __init__(self, bucket: str = None, object: str = None, generation: int = None):
        self.bucket = bucket
        self.object = object
        self.generation = generation

    @classmethod
    def to_proto(self, resource):
        if not resource:
            return None

        res = (
            os_policy_assignment_pb2.OsconfigBetaOSPolicyAssignmentOSPoliciesResourceGroupsResourcesPkgMsiSourceGcs()
        )
        if Primitive.to_proto(resource.bucket):
            res.bucket = Primitive.to_proto(resource.bucket)
        if Primitive.to_proto(resource.object):
            res.object = Primitive.to_proto(resource.object)
        if Primitive.to_proto(resource.generation):
            res.generation = Primitive.to_proto(resource.generation)
        return res

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return None

        return OSPolicyAssignmentOSPoliciesResourceGroupsResourcesPkgMsiSourceGcs(
            bucket=Primitive.from_proto(resource.bucket),
            object=Primitive.from_proto(resource.object),
            generation=Primitive.from_proto(resource.generation),
        )


class OSPolicyAssignmentOSPoliciesResourceGroupsResourcesPkgMsiSourceGcsArray(object):
    @classmethod
    def to_proto(self, resources):
        if not resources:
            return resources
        return [
            OSPolicyAssignmentOSPoliciesResourceGroupsResourcesPkgMsiSourceGcs.to_proto(
                i
            )
            for i in resources
        ]

    @classmethod
    def from_proto(self, resources):
        return [
            OSPolicyAssignmentOSPoliciesResourceGroupsResourcesPkgMsiSourceGcs.from_proto(
                i
            )
            for i in resources
        ]


class OSPolicyAssignmentOSPoliciesResourceGroupsResourcesRepository(object):
    def __init__(
        self, apt: dict = None, yum: dict = None, zypper: dict = None, goo: dict = None
    ):
        self.apt = apt
        self.yum = yum
        self.zypper = zypper
        self.goo = goo

    @classmethod
    def to_proto(self, resource):
        if not resource:
            return None

        res = (
            os_policy_assignment_pb2.OsconfigBetaOSPolicyAssignmentOSPoliciesResourceGroupsResourcesRepository()
        )
        if OSPolicyAssignmentOSPoliciesResourceGroupsResourcesRepositoryApt.to_proto(
            resource.apt
        ):
            res.apt.CopyFrom(
                OSPolicyAssignmentOSPoliciesResourceGroupsResourcesRepositoryApt.to_proto(
                    resource.apt
                )
            )
        else:
            res.ClearField("apt")
        if OSPolicyAssignmentOSPoliciesResourceGroupsResourcesRepositoryYum.to_proto(
            resource.yum
        ):
            res.yum.CopyFrom(
                OSPolicyAssignmentOSPoliciesResourceGroupsResourcesRepositoryYum.to_proto(
                    resource.yum
                )
            )
        else:
            res.ClearField("yum")
        if OSPolicyAssignmentOSPoliciesResourceGroupsResourcesRepositoryZypper.to_proto(
            resource.zypper
        ):
            res.zypper.CopyFrom(
                OSPolicyAssignmentOSPoliciesResourceGroupsResourcesRepositoryZypper.to_proto(
                    resource.zypper
                )
            )
        else:
            res.ClearField("zypper")
        if OSPolicyAssignmentOSPoliciesResourceGroupsResourcesRepositoryGoo.to_proto(
            resource.goo
        ):
            res.goo.CopyFrom(
                OSPolicyAssignmentOSPoliciesResourceGroupsResourcesRepositoryGoo.to_proto(
                    resource.goo
                )
            )
        else:
            res.ClearField("goo")
        return res

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return None

        return OSPolicyAssignmentOSPoliciesResourceGroupsResourcesRepository(
            apt=OSPolicyAssignmentOSPoliciesResourceGroupsResourcesRepositoryApt.from_proto(
                resource.apt
            ),
            yum=OSPolicyAssignmentOSPoliciesResourceGroupsResourcesRepositoryYum.from_proto(
                resource.yum
            ),
            zypper=OSPolicyAssignmentOSPoliciesResourceGroupsResourcesRepositoryZypper.from_proto(
                resource.zypper
            ),
            goo=OSPolicyAssignmentOSPoliciesResourceGroupsResourcesRepositoryGoo.from_proto(
                resource.goo
            ),
        )


class OSPolicyAssignmentOSPoliciesResourceGroupsResourcesRepositoryArray(object):
    @classmethod
    def to_proto(self, resources):
        if not resources:
            return resources
        return [
            OSPolicyAssignmentOSPoliciesResourceGroupsResourcesRepository.to_proto(i)
            for i in resources
        ]

    @classmethod
    def from_proto(self, resources):
        return [
            OSPolicyAssignmentOSPoliciesResourceGroupsResourcesRepository.from_proto(i)
            for i in resources
        ]


class OSPolicyAssignmentOSPoliciesResourceGroupsResourcesRepositoryApt(object):
    def __init__(
        self,
        archive_type: str = None,
        uri: str = None,
        distribution: str = None,
        components: list = None,
        gpg_key: str = None,
    ):
        self.archive_type = archive_type
        self.uri = uri
        self.distribution = distribution
        self.components = components
        self.gpg_key = gpg_key

    @classmethod
    def to_proto(self, resource):
        if not resource:
            return None

        res = (
            os_policy_assignment_pb2.OsconfigBetaOSPolicyAssignmentOSPoliciesResourceGroupsResourcesRepositoryApt()
        )
        if OSPolicyAssignmentOSPoliciesResourceGroupsResourcesRepositoryAptArchiveTypeEnum.to_proto(
            resource.archive_type
        ):
            res.archive_type = OSPolicyAssignmentOSPoliciesResourceGroupsResourcesRepositoryAptArchiveTypeEnum.to_proto(
                resource.archive_type
            )
        if Primitive.to_proto(resource.uri):
            res.uri = Primitive.to_proto(resource.uri)
        if Primitive.to_proto(resource.distribution):
            res.distribution = Primitive.to_proto(resource.distribution)
        if Primitive.to_proto(resource.components):
            res.components.extend(Primitive.to_proto(resource.components))
        if Primitive.to_proto(resource.gpg_key):
            res.gpg_key = Primitive.to_proto(resource.gpg_key)
        return res

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return None

        return OSPolicyAssignmentOSPoliciesResourceGroupsResourcesRepositoryApt(
            archive_type=OSPolicyAssignmentOSPoliciesResourceGroupsResourcesRepositoryAptArchiveTypeEnum.from_proto(
                resource.archive_type
            ),
            uri=Primitive.from_proto(resource.uri),
            distribution=Primitive.from_proto(resource.distribution),
            components=Primitive.from_proto(resource.components),
            gpg_key=Primitive.from_proto(resource.gpg_key),
        )


class OSPolicyAssignmentOSPoliciesResourceGroupsResourcesRepositoryAptArray(object):
    @classmethod
    def to_proto(self, resources):
        if not resources:
            return resources
        return [
            OSPolicyAssignmentOSPoliciesResourceGroupsResourcesRepositoryApt.to_proto(i)
            for i in resources
        ]

    @classmethod
    def from_proto(self, resources):
        return [
            OSPolicyAssignmentOSPoliciesResourceGroupsResourcesRepositoryApt.from_proto(
                i
            )
            for i in resources
        ]


class OSPolicyAssignmentOSPoliciesResourceGroupsResourcesRepositoryYum(object):
    def __init__(
        self,
        id: str = None,
        display_name: str = None,
        base_url: str = None,
        gpg_keys: list = None,
    ):
        self.id = id
        self.display_name = display_name
        self.base_url = base_url
        self.gpg_keys = gpg_keys

    @classmethod
    def to_proto(self, resource):
        if not resource:
            return None

        res = (
            os_policy_assignment_pb2.OsconfigBetaOSPolicyAssignmentOSPoliciesResourceGroupsResourcesRepositoryYum()
        )
        if Primitive.to_proto(resource.id):
            res.id = Primitive.to_proto(resource.id)
        if Primitive.to_proto(resource.display_name):
            res.display_name = Primitive.to_proto(resource.display_name)
        if Primitive.to_proto(resource.base_url):
            res.base_url = Primitive.to_proto(resource.base_url)
        if Primitive.to_proto(resource.gpg_keys):
            res.gpg_keys.extend(Primitive.to_proto(resource.gpg_keys))
        return res

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return None

        return OSPolicyAssignmentOSPoliciesResourceGroupsResourcesRepositoryYum(
            id=Primitive.from_proto(resource.id),
            display_name=Primitive.from_proto(resource.display_name),
            base_url=Primitive.from_proto(resource.base_url),
            gpg_keys=Primitive.from_proto(resource.gpg_keys),
        )


class OSPolicyAssignmentOSPoliciesResourceGroupsResourcesRepositoryYumArray(object):
    @classmethod
    def to_proto(self, resources):
        if not resources:
            return resources
        return [
            OSPolicyAssignmentOSPoliciesResourceGroupsResourcesRepositoryYum.to_proto(i)
            for i in resources
        ]

    @classmethod
    def from_proto(self, resources):
        return [
            OSPolicyAssignmentOSPoliciesResourceGroupsResourcesRepositoryYum.from_proto(
                i
            )
            for i in resources
        ]


class OSPolicyAssignmentOSPoliciesResourceGroupsResourcesRepositoryZypper(object):
    def __init__(
        self,
        id: str = None,
        display_name: str = None,
        base_url: str = None,
        gpg_keys: list = None,
    ):
        self.id = id
        self.display_name = display_name
        self.base_url = base_url
        self.gpg_keys = gpg_keys

    @classmethod
    def to_proto(self, resource):
        if not resource:
            return None

        res = (
            os_policy_assignment_pb2.OsconfigBetaOSPolicyAssignmentOSPoliciesResourceGroupsResourcesRepositoryZypper()
        )
        if Primitive.to_proto(resource.id):
            res.id = Primitive.to_proto(resource.id)
        if Primitive.to_proto(resource.display_name):
            res.display_name = Primitive.to_proto(resource.display_name)
        if Primitive.to_proto(resource.base_url):
            res.base_url = Primitive.to_proto(resource.base_url)
        if Primitive.to_proto(resource.gpg_keys):
            res.gpg_keys.extend(Primitive.to_proto(resource.gpg_keys))
        return res

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return None

        return OSPolicyAssignmentOSPoliciesResourceGroupsResourcesRepositoryZypper(
            id=Primitive.from_proto(resource.id),
            display_name=Primitive.from_proto(resource.display_name),
            base_url=Primitive.from_proto(resource.base_url),
            gpg_keys=Primitive.from_proto(resource.gpg_keys),
        )


class OSPolicyAssignmentOSPoliciesResourceGroupsResourcesRepositoryZypperArray(object):
    @classmethod
    def to_proto(self, resources):
        if not resources:
            return resources
        return [
            OSPolicyAssignmentOSPoliciesResourceGroupsResourcesRepositoryZypper.to_proto(
                i
            )
            for i in resources
        ]

    @classmethod
    def from_proto(self, resources):
        return [
            OSPolicyAssignmentOSPoliciesResourceGroupsResourcesRepositoryZypper.from_proto(
                i
            )
            for i in resources
        ]


class OSPolicyAssignmentOSPoliciesResourceGroupsResourcesRepositoryGoo(object):
    def __init__(self, name: str = None, url: str = None):
        self.name = name
        self.url = url

    @classmethod
    def to_proto(self, resource):
        if not resource:
            return None

        res = (
            os_policy_assignment_pb2.OsconfigBetaOSPolicyAssignmentOSPoliciesResourceGroupsResourcesRepositoryGoo()
        )
        if Primitive.to_proto(resource.name):
            res.name = Primitive.to_proto(resource.name)
        if Primitive.to_proto(resource.url):
            res.url = Primitive.to_proto(resource.url)
        return res

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return None

        return OSPolicyAssignmentOSPoliciesResourceGroupsResourcesRepositoryGoo(
            name=Primitive.from_proto(resource.name),
            url=Primitive.from_proto(resource.url),
        )


class OSPolicyAssignmentOSPoliciesResourceGroupsResourcesRepositoryGooArray(object):
    @classmethod
    def to_proto(self, resources):
        if not resources:
            return resources
        return [
            OSPolicyAssignmentOSPoliciesResourceGroupsResourcesRepositoryGoo.to_proto(i)
            for i in resources
        ]

    @classmethod
    def from_proto(self, resources):
        return [
            OSPolicyAssignmentOSPoliciesResourceGroupsResourcesRepositoryGoo.from_proto(
                i
            )
            for i in resources
        ]


class OSPolicyAssignmentOSPoliciesResourceGroupsResourcesExec(object):
    def __init__(self, validate: dict = None, enforce: dict = None):
        self.validate = validate
        self.enforce = enforce

    @classmethod
    def to_proto(self, resource):
        if not resource:
            return None

        res = (
            os_policy_assignment_pb2.OsconfigBetaOSPolicyAssignmentOSPoliciesResourceGroupsResourcesExec()
        )
        if OSPolicyAssignmentOSPoliciesResourceGroupsResourcesExecValidate.to_proto(
            resource.validate
        ):
            res.validate.CopyFrom(
                OSPolicyAssignmentOSPoliciesResourceGroupsResourcesExecValidate.to_proto(
                    resource.validate
                )
            )
        else:
            res.ClearField("validate")
        if OSPolicyAssignmentOSPoliciesResourceGroupsResourcesExecEnforce.to_proto(
            resource.enforce
        ):
            res.enforce.CopyFrom(
                OSPolicyAssignmentOSPoliciesResourceGroupsResourcesExecEnforce.to_proto(
                    resource.enforce
                )
            )
        else:
            res.ClearField("enforce")
        return res

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return None

        return OSPolicyAssignmentOSPoliciesResourceGroupsResourcesExec(
            validate=OSPolicyAssignmentOSPoliciesResourceGroupsResourcesExecValidate.from_proto(
                resource.validate
            ),
            enforce=OSPolicyAssignmentOSPoliciesResourceGroupsResourcesExecEnforce.from_proto(
                resource.enforce
            ),
        )


class OSPolicyAssignmentOSPoliciesResourceGroupsResourcesExecArray(object):
    @classmethod
    def to_proto(self, resources):
        if not resources:
            return resources
        return [
            OSPolicyAssignmentOSPoliciesResourceGroupsResourcesExec.to_proto(i)
            for i in resources
        ]

    @classmethod
    def from_proto(self, resources):
        return [
            OSPolicyAssignmentOSPoliciesResourceGroupsResourcesExec.from_proto(i)
            for i in resources
        ]


class OSPolicyAssignmentOSPoliciesResourceGroupsResourcesExecValidate(object):
    def __init__(
        self,
        file: dict = None,
        script: str = None,
        args: list = None,
        interpreter: str = None,
        output_file_path: str = None,
    ):
        self.file = file
        self.script = script
        self.args = args
        self.interpreter = interpreter
        self.output_file_path = output_file_path

    @classmethod
    def to_proto(self, resource):
        if not resource:
            return None

        res = (
            os_policy_assignment_pb2.OsconfigBetaOSPolicyAssignmentOSPoliciesResourceGroupsResourcesExecValidate()
        )
        if OSPolicyAssignmentOSPoliciesResourceGroupsResourcesExecValidateFile.to_proto(
            resource.file
        ):
            res.file.CopyFrom(
                OSPolicyAssignmentOSPoliciesResourceGroupsResourcesExecValidateFile.to_proto(
                    resource.file
                )
            )
        else:
            res.ClearField("file")
        if Primitive.to_proto(resource.script):
            res.script = Primitive.to_proto(resource.script)
        if Primitive.to_proto(resource.args):
            res.args.extend(Primitive.to_proto(resource.args))
        if OSPolicyAssignmentOSPoliciesResourceGroupsResourcesExecValidateInterpreterEnum.to_proto(
            resource.interpreter
        ):
            res.interpreter = OSPolicyAssignmentOSPoliciesResourceGroupsResourcesExecValidateInterpreterEnum.to_proto(
                resource.interpreter
            )
        if Primitive.to_proto(resource.output_file_path):
            res.output_file_path = Primitive.to_proto(resource.output_file_path)
        return res

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return None

        return OSPolicyAssignmentOSPoliciesResourceGroupsResourcesExecValidate(
            file=OSPolicyAssignmentOSPoliciesResourceGroupsResourcesExecValidateFile.from_proto(
                resource.file
            ),
            script=Primitive.from_proto(resource.script),
            args=Primitive.from_proto(resource.args),
            interpreter=OSPolicyAssignmentOSPoliciesResourceGroupsResourcesExecValidateInterpreterEnum.from_proto(
                resource.interpreter
            ),
            output_file_path=Primitive.from_proto(resource.output_file_path),
        )


class OSPolicyAssignmentOSPoliciesResourceGroupsResourcesExecValidateArray(object):
    @classmethod
    def to_proto(self, resources):
        if not resources:
            return resources
        return [
            OSPolicyAssignmentOSPoliciesResourceGroupsResourcesExecValidate.to_proto(i)
            for i in resources
        ]

    @classmethod
    def from_proto(self, resources):
        return [
            OSPolicyAssignmentOSPoliciesResourceGroupsResourcesExecValidate.from_proto(
                i
            )
            for i in resources
        ]


class OSPolicyAssignmentOSPoliciesResourceGroupsResourcesExecValidateFile(object):
    def __init__(
        self,
        remote: dict = None,
        gcs: dict = None,
        local_path: str = None,
        allow_insecure: bool = None,
    ):
        self.remote = remote
        self.gcs = gcs
        self.local_path = local_path
        self.allow_insecure = allow_insecure

    @classmethod
    def to_proto(self, resource):
        if not resource:
            return None

        res = (
            os_policy_assignment_pb2.OsconfigBetaOSPolicyAssignmentOSPoliciesResourceGroupsResourcesExecValidateFile()
        )
        if OSPolicyAssignmentOSPoliciesResourceGroupsResourcesExecValidateFileRemote.to_proto(
            resource.remote
        ):
            res.remote.CopyFrom(
                OSPolicyAssignmentOSPoliciesResourceGroupsResourcesExecValidateFileRemote.to_proto(
                    resource.remote
                )
            )
        else:
            res.ClearField("remote")
        if OSPolicyAssignmentOSPoliciesResourceGroupsResourcesExecValidateFileGcs.to_proto(
            resource.gcs
        ):
            res.gcs.CopyFrom(
                OSPolicyAssignmentOSPoliciesResourceGroupsResourcesExecValidateFileGcs.to_proto(
                    resource.gcs
                )
            )
        else:
            res.ClearField("gcs")
        if Primitive.to_proto(resource.local_path):
            res.local_path = Primitive.to_proto(resource.local_path)
        if Primitive.to_proto(resource.allow_insecure):
            res.allow_insecure = Primitive.to_proto(resource.allow_insecure)
        return res

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return None

        return OSPolicyAssignmentOSPoliciesResourceGroupsResourcesExecValidateFile(
            remote=OSPolicyAssignmentOSPoliciesResourceGroupsResourcesExecValidateFileRemote.from_proto(
                resource.remote
            ),
            gcs=OSPolicyAssignmentOSPoliciesResourceGroupsResourcesExecValidateFileGcs.from_proto(
                resource.gcs
            ),
            local_path=Primitive.from_proto(resource.local_path),
            allow_insecure=Primitive.from_proto(resource.allow_insecure),
        )


class OSPolicyAssignmentOSPoliciesResourceGroupsResourcesExecValidateFileArray(object):
    @classmethod
    def to_proto(self, resources):
        if not resources:
            return resources
        return [
            OSPolicyAssignmentOSPoliciesResourceGroupsResourcesExecValidateFile.to_proto(
                i
            )
            for i in resources
        ]

    @classmethod
    def from_proto(self, resources):
        return [
            OSPolicyAssignmentOSPoliciesResourceGroupsResourcesExecValidateFile.from_proto(
                i
            )
            for i in resources
        ]


class OSPolicyAssignmentOSPoliciesResourceGroupsResourcesExecValidateFileRemote(object):
    def __init__(self, uri: str = None, sha256_checksum: str = None):
        self.uri = uri
        self.sha256_checksum = sha256_checksum

    @classmethod
    def to_proto(self, resource):
        if not resource:
            return None

        res = (
            os_policy_assignment_pb2.OsconfigBetaOSPolicyAssignmentOSPoliciesResourceGroupsResourcesExecValidateFileRemote()
        )
        if Primitive.to_proto(resource.uri):
            res.uri = Primitive.to_proto(resource.uri)
        if Primitive.to_proto(resource.sha256_checksum):
            res.sha256_checksum = Primitive.to_proto(resource.sha256_checksum)
        return res

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return None

        return (
            OSPolicyAssignmentOSPoliciesResourceGroupsResourcesExecValidateFileRemote(
                uri=Primitive.from_proto(resource.uri),
                sha256_checksum=Primitive.from_proto(resource.sha256_checksum),
            )
        )


class OSPolicyAssignmentOSPoliciesResourceGroupsResourcesExecValidateFileRemoteArray(
    object
):
    @classmethod
    def to_proto(self, resources):
        if not resources:
            return resources
        return [
            OSPolicyAssignmentOSPoliciesResourceGroupsResourcesExecValidateFileRemote.to_proto(
                i
            )
            for i in resources
        ]

    @classmethod
    def from_proto(self, resources):
        return [
            OSPolicyAssignmentOSPoliciesResourceGroupsResourcesExecValidateFileRemote.from_proto(
                i
            )
            for i in resources
        ]


class OSPolicyAssignmentOSPoliciesResourceGroupsResourcesExecValidateFileGcs(object):
    def __init__(self, bucket: str = None, object: str = None, generation: int = None):
        self.bucket = bucket
        self.object = object
        self.generation = generation

    @classmethod
    def to_proto(self, resource):
        if not resource:
            return None

        res = (
            os_policy_assignment_pb2.OsconfigBetaOSPolicyAssignmentOSPoliciesResourceGroupsResourcesExecValidateFileGcs()
        )
        if Primitive.to_proto(resource.bucket):
            res.bucket = Primitive.to_proto(resource.bucket)
        if Primitive.to_proto(resource.object):
            res.object = Primitive.to_proto(resource.object)
        if Primitive.to_proto(resource.generation):
            res.generation = Primitive.to_proto(resource.generation)
        return res

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return None

        return OSPolicyAssignmentOSPoliciesResourceGroupsResourcesExecValidateFileGcs(
            bucket=Primitive.from_proto(resource.bucket),
            object=Primitive.from_proto(resource.object),
            generation=Primitive.from_proto(resource.generation),
        )


class OSPolicyAssignmentOSPoliciesResourceGroupsResourcesExecValidateFileGcsArray(
    object
):
    @classmethod
    def to_proto(self, resources):
        if not resources:
            return resources
        return [
            OSPolicyAssignmentOSPoliciesResourceGroupsResourcesExecValidateFileGcs.to_proto(
                i
            )
            for i in resources
        ]

    @classmethod
    def from_proto(self, resources):
        return [
            OSPolicyAssignmentOSPoliciesResourceGroupsResourcesExecValidateFileGcs.from_proto(
                i
            )
            for i in resources
        ]


class OSPolicyAssignmentOSPoliciesResourceGroupsResourcesExecEnforce(object):
    def __init__(
        self,
        file: dict = None,
        script: str = None,
        args: list = None,
        interpreter: str = None,
        output_file_path: str = None,
    ):
        self.file = file
        self.script = script
        self.args = args
        self.interpreter = interpreter
        self.output_file_path = output_file_path

    @classmethod
    def to_proto(self, resource):
        if not resource:
            return None

        res = (
            os_policy_assignment_pb2.OsconfigBetaOSPolicyAssignmentOSPoliciesResourceGroupsResourcesExecEnforce()
        )
        if OSPolicyAssignmentOSPoliciesResourceGroupsResourcesExecEnforceFile.to_proto(
            resource.file
        ):
            res.file.CopyFrom(
                OSPolicyAssignmentOSPoliciesResourceGroupsResourcesExecEnforceFile.to_proto(
                    resource.file
                )
            )
        else:
            res.ClearField("file")
        if Primitive.to_proto(resource.script):
            res.script = Primitive.to_proto(resource.script)
        if Primitive.to_proto(resource.args):
            res.args.extend(Primitive.to_proto(resource.args))
        if OSPolicyAssignmentOSPoliciesResourceGroupsResourcesExecEnforceInterpreterEnum.to_proto(
            resource.interpreter
        ):
            res.interpreter = OSPolicyAssignmentOSPoliciesResourceGroupsResourcesExecEnforceInterpreterEnum.to_proto(
                resource.interpreter
            )
        if Primitive.to_proto(resource.output_file_path):
            res.output_file_path = Primitive.to_proto(resource.output_file_path)
        return res

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return None

        return OSPolicyAssignmentOSPoliciesResourceGroupsResourcesExecEnforce(
            file=OSPolicyAssignmentOSPoliciesResourceGroupsResourcesExecEnforceFile.from_proto(
                resource.file
            ),
            script=Primitive.from_proto(resource.script),
            args=Primitive.from_proto(resource.args),
            interpreter=OSPolicyAssignmentOSPoliciesResourceGroupsResourcesExecEnforceInterpreterEnum.from_proto(
                resource.interpreter
            ),
            output_file_path=Primitive.from_proto(resource.output_file_path),
        )


class OSPolicyAssignmentOSPoliciesResourceGroupsResourcesExecEnforceArray(object):
    @classmethod
    def to_proto(self, resources):
        if not resources:
            return resources
        return [
            OSPolicyAssignmentOSPoliciesResourceGroupsResourcesExecEnforce.to_proto(i)
            for i in resources
        ]

    @classmethod
    def from_proto(self, resources):
        return [
            OSPolicyAssignmentOSPoliciesResourceGroupsResourcesExecEnforce.from_proto(i)
            for i in resources
        ]


class OSPolicyAssignmentOSPoliciesResourceGroupsResourcesExecEnforceFile(object):
    def __init__(
        self,
        remote: dict = None,
        gcs: dict = None,
        local_path: str = None,
        allow_insecure: bool = None,
    ):
        self.remote = remote
        self.gcs = gcs
        self.local_path = local_path
        self.allow_insecure = allow_insecure

    @classmethod
    def to_proto(self, resource):
        if not resource:
            return None

        res = (
            os_policy_assignment_pb2.OsconfigBetaOSPolicyAssignmentOSPoliciesResourceGroupsResourcesExecEnforceFile()
        )
        if OSPolicyAssignmentOSPoliciesResourceGroupsResourcesExecEnforceFileRemote.to_proto(
            resource.remote
        ):
            res.remote.CopyFrom(
                OSPolicyAssignmentOSPoliciesResourceGroupsResourcesExecEnforceFileRemote.to_proto(
                    resource.remote
                )
            )
        else:
            res.ClearField("remote")
        if OSPolicyAssignmentOSPoliciesResourceGroupsResourcesExecEnforceFileGcs.to_proto(
            resource.gcs
        ):
            res.gcs.CopyFrom(
                OSPolicyAssignmentOSPoliciesResourceGroupsResourcesExecEnforceFileGcs.to_proto(
                    resource.gcs
                )
            )
        else:
            res.ClearField("gcs")
        if Primitive.to_proto(resource.local_path):
            res.local_path = Primitive.to_proto(resource.local_path)
        if Primitive.to_proto(resource.allow_insecure):
            res.allow_insecure = Primitive.to_proto(resource.allow_insecure)
        return res

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return None

        return OSPolicyAssignmentOSPoliciesResourceGroupsResourcesExecEnforceFile(
            remote=OSPolicyAssignmentOSPoliciesResourceGroupsResourcesExecEnforceFileRemote.from_proto(
                resource.remote
            ),
            gcs=OSPolicyAssignmentOSPoliciesResourceGroupsResourcesExecEnforceFileGcs.from_proto(
                resource.gcs
            ),
            local_path=Primitive.from_proto(resource.local_path),
            allow_insecure=Primitive.from_proto(resource.allow_insecure),
        )


class OSPolicyAssignmentOSPoliciesResourceGroupsResourcesExecEnforceFileArray(object):
    @classmethod
    def to_proto(self, resources):
        if not resources:
            return resources
        return [
            OSPolicyAssignmentOSPoliciesResourceGroupsResourcesExecEnforceFile.to_proto(
                i
            )
            for i in resources
        ]

    @classmethod
    def from_proto(self, resources):
        return [
            OSPolicyAssignmentOSPoliciesResourceGroupsResourcesExecEnforceFile.from_proto(
                i
            )
            for i in resources
        ]


class OSPolicyAssignmentOSPoliciesResourceGroupsResourcesExecEnforceFileRemote(object):
    def __init__(self, uri: str = None, sha256_checksum: str = None):
        self.uri = uri
        self.sha256_checksum = sha256_checksum

    @classmethod
    def to_proto(self, resource):
        if not resource:
            return None

        res = (
            os_policy_assignment_pb2.OsconfigBetaOSPolicyAssignmentOSPoliciesResourceGroupsResourcesExecEnforceFileRemote()
        )
        if Primitive.to_proto(resource.uri):
            res.uri = Primitive.to_proto(resource.uri)
        if Primitive.to_proto(resource.sha256_checksum):
            res.sha256_checksum = Primitive.to_proto(resource.sha256_checksum)
        return res

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return None

        return OSPolicyAssignmentOSPoliciesResourceGroupsResourcesExecEnforceFileRemote(
            uri=Primitive.from_proto(resource.uri),
            sha256_checksum=Primitive.from_proto(resource.sha256_checksum),
        )


class OSPolicyAssignmentOSPoliciesResourceGroupsResourcesExecEnforceFileRemoteArray(
    object
):
    @classmethod
    def to_proto(self, resources):
        if not resources:
            return resources
        return [
            OSPolicyAssignmentOSPoliciesResourceGroupsResourcesExecEnforceFileRemote.to_proto(
                i
            )
            for i in resources
        ]

    @classmethod
    def from_proto(self, resources):
        return [
            OSPolicyAssignmentOSPoliciesResourceGroupsResourcesExecEnforceFileRemote.from_proto(
                i
            )
            for i in resources
        ]


class OSPolicyAssignmentOSPoliciesResourceGroupsResourcesExecEnforceFileGcs(object):
    def __init__(self, bucket: str = None, object: str = None, generation: int = None):
        self.bucket = bucket
        self.object = object
        self.generation = generation

    @classmethod
    def to_proto(self, resource):
        if not resource:
            return None

        res = (
            os_policy_assignment_pb2.OsconfigBetaOSPolicyAssignmentOSPoliciesResourceGroupsResourcesExecEnforceFileGcs()
        )
        if Primitive.to_proto(resource.bucket):
            res.bucket = Primitive.to_proto(resource.bucket)
        if Primitive.to_proto(resource.object):
            res.object = Primitive.to_proto(resource.object)
        if Primitive.to_proto(resource.generation):
            res.generation = Primitive.to_proto(resource.generation)
        return res

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return None

        return OSPolicyAssignmentOSPoliciesResourceGroupsResourcesExecEnforceFileGcs(
            bucket=Primitive.from_proto(resource.bucket),
            object=Primitive.from_proto(resource.object),
            generation=Primitive.from_proto(resource.generation),
        )


class OSPolicyAssignmentOSPoliciesResourceGroupsResourcesExecEnforceFileGcsArray(
    object
):
    @classmethod
    def to_proto(self, resources):
        if not resources:
            return resources
        return [
            OSPolicyAssignmentOSPoliciesResourceGroupsResourcesExecEnforceFileGcs.to_proto(
                i
            )
            for i in resources
        ]

    @classmethod
    def from_proto(self, resources):
        return [
            OSPolicyAssignmentOSPoliciesResourceGroupsResourcesExecEnforceFileGcs.from_proto(
                i
            )
            for i in resources
        ]


class OSPolicyAssignmentOSPoliciesResourceGroupsResourcesFile(object):
    def __init__(
        self,
        file: dict = None,
        content: str = None,
        path: str = None,
        state: str = None,
        permissions: str = None,
    ):
        self.file = file
        self.content = content
        self.path = path
        self.state = state
        self.permissions = permissions

    @classmethod
    def to_proto(self, resource):
        if not resource:
            return None

        res = (
            os_policy_assignment_pb2.OsconfigBetaOSPolicyAssignmentOSPoliciesResourceGroupsResourcesFile()
        )
        if OSPolicyAssignmentOSPoliciesResourceGroupsResourcesFileFile.to_proto(
            resource.file
        ):
            res.file.CopyFrom(
                OSPolicyAssignmentOSPoliciesResourceGroupsResourcesFileFile.to_proto(
                    resource.file
                )
            )
        else:
            res.ClearField("file")
        if Primitive.to_proto(resource.content):
            res.content = Primitive.to_proto(resource.content)
        if Primitive.to_proto(resource.path):
            res.path = Primitive.to_proto(resource.path)
        if OSPolicyAssignmentOSPoliciesResourceGroupsResourcesFileStateEnum.to_proto(
            resource.state
        ):
            res.state = OSPolicyAssignmentOSPoliciesResourceGroupsResourcesFileStateEnum.to_proto(
                resource.state
            )
        if Primitive.to_proto(resource.permissions):
            res.permissions = Primitive.to_proto(resource.permissions)
        return res

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return None

        return OSPolicyAssignmentOSPoliciesResourceGroupsResourcesFile(
            file=OSPolicyAssignmentOSPoliciesResourceGroupsResourcesFileFile.from_proto(
                resource.file
            ),
            content=Primitive.from_proto(resource.content),
            path=Primitive.from_proto(resource.path),
            state=OSPolicyAssignmentOSPoliciesResourceGroupsResourcesFileStateEnum.from_proto(
                resource.state
            ),
            permissions=Primitive.from_proto(resource.permissions),
        )


class OSPolicyAssignmentOSPoliciesResourceGroupsResourcesFileArray(object):
    @classmethod
    def to_proto(self, resources):
        if not resources:
            return resources
        return [
            OSPolicyAssignmentOSPoliciesResourceGroupsResourcesFile.to_proto(i)
            for i in resources
        ]

    @classmethod
    def from_proto(self, resources):
        return [
            OSPolicyAssignmentOSPoliciesResourceGroupsResourcesFile.from_proto(i)
            for i in resources
        ]


class OSPolicyAssignmentOSPoliciesResourceGroupsResourcesFileFile(object):
    def __init__(
        self,
        remote: dict = None,
        gcs: dict = None,
        local_path: str = None,
        allow_insecure: bool = None,
    ):
        self.remote = remote
        self.gcs = gcs
        self.local_path = local_path
        self.allow_insecure = allow_insecure

    @classmethod
    def to_proto(self, resource):
        if not resource:
            return None

        res = (
            os_policy_assignment_pb2.OsconfigBetaOSPolicyAssignmentOSPoliciesResourceGroupsResourcesFileFile()
        )
        if OSPolicyAssignmentOSPoliciesResourceGroupsResourcesFileFileRemote.to_proto(
            resource.remote
        ):
            res.remote.CopyFrom(
                OSPolicyAssignmentOSPoliciesResourceGroupsResourcesFileFileRemote.to_proto(
                    resource.remote
                )
            )
        else:
            res.ClearField("remote")
        if OSPolicyAssignmentOSPoliciesResourceGroupsResourcesFileFileGcs.to_proto(
            resource.gcs
        ):
            res.gcs.CopyFrom(
                OSPolicyAssignmentOSPoliciesResourceGroupsResourcesFileFileGcs.to_proto(
                    resource.gcs
                )
            )
        else:
            res.ClearField("gcs")
        if Primitive.to_proto(resource.local_path):
            res.local_path = Primitive.to_proto(resource.local_path)
        if Primitive.to_proto(resource.allow_insecure):
            res.allow_insecure = Primitive.to_proto(resource.allow_insecure)
        return res

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return None

        return OSPolicyAssignmentOSPoliciesResourceGroupsResourcesFileFile(
            remote=OSPolicyAssignmentOSPoliciesResourceGroupsResourcesFileFileRemote.from_proto(
                resource.remote
            ),
            gcs=OSPolicyAssignmentOSPoliciesResourceGroupsResourcesFileFileGcs.from_proto(
                resource.gcs
            ),
            local_path=Primitive.from_proto(resource.local_path),
            allow_insecure=Primitive.from_proto(resource.allow_insecure),
        )


class OSPolicyAssignmentOSPoliciesResourceGroupsResourcesFileFileArray(object):
    @classmethod
    def to_proto(self, resources):
        if not resources:
            return resources
        return [
            OSPolicyAssignmentOSPoliciesResourceGroupsResourcesFileFile.to_proto(i)
            for i in resources
        ]

    @classmethod
    def from_proto(self, resources):
        return [
            OSPolicyAssignmentOSPoliciesResourceGroupsResourcesFileFile.from_proto(i)
            for i in resources
        ]


class OSPolicyAssignmentOSPoliciesResourceGroupsResourcesFileFileRemote(object):
    def __init__(self, uri: str = None, sha256_checksum: str = None):
        self.uri = uri
        self.sha256_checksum = sha256_checksum

    @classmethod
    def to_proto(self, resource):
        if not resource:
            return None

        res = (
            os_policy_assignment_pb2.OsconfigBetaOSPolicyAssignmentOSPoliciesResourceGroupsResourcesFileFileRemote()
        )
        if Primitive.to_proto(resource.uri):
            res.uri = Primitive.to_proto(resource.uri)
        if Primitive.to_proto(resource.sha256_checksum):
            res.sha256_checksum = Primitive.to_proto(resource.sha256_checksum)
        return res

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return None

        return OSPolicyAssignmentOSPoliciesResourceGroupsResourcesFileFileRemote(
            uri=Primitive.from_proto(resource.uri),
            sha256_checksum=Primitive.from_proto(resource.sha256_checksum),
        )


class OSPolicyAssignmentOSPoliciesResourceGroupsResourcesFileFileRemoteArray(object):
    @classmethod
    def to_proto(self, resources):
        if not resources:
            return resources
        return [
            OSPolicyAssignmentOSPoliciesResourceGroupsResourcesFileFileRemote.to_proto(
                i
            )
            for i in resources
        ]

    @classmethod
    def from_proto(self, resources):
        return [
            OSPolicyAssignmentOSPoliciesResourceGroupsResourcesFileFileRemote.from_proto(
                i
            )
            for i in resources
        ]


class OSPolicyAssignmentOSPoliciesResourceGroupsResourcesFileFileGcs(object):
    def __init__(self, bucket: str = None, object: str = None, generation: int = None):
        self.bucket = bucket
        self.object = object
        self.generation = generation

    @classmethod
    def to_proto(self, resource):
        if not resource:
            return None

        res = (
            os_policy_assignment_pb2.OsconfigBetaOSPolicyAssignmentOSPoliciesResourceGroupsResourcesFileFileGcs()
        )
        if Primitive.to_proto(resource.bucket):
            res.bucket = Primitive.to_proto(resource.bucket)
        if Primitive.to_proto(resource.object):
            res.object = Primitive.to_proto(resource.object)
        if Primitive.to_proto(resource.generation):
            res.generation = Primitive.to_proto(resource.generation)
        return res

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return None

        return OSPolicyAssignmentOSPoliciesResourceGroupsResourcesFileFileGcs(
            bucket=Primitive.from_proto(resource.bucket),
            object=Primitive.from_proto(resource.object),
            generation=Primitive.from_proto(resource.generation),
        )


class OSPolicyAssignmentOSPoliciesResourceGroupsResourcesFileFileGcsArray(object):
    @classmethod
    def to_proto(self, resources):
        if not resources:
            return resources
        return [
            OSPolicyAssignmentOSPoliciesResourceGroupsResourcesFileFileGcs.to_proto(i)
            for i in resources
        ]

    @classmethod
    def from_proto(self, resources):
        return [
            OSPolicyAssignmentOSPoliciesResourceGroupsResourcesFileFileGcs.from_proto(i)
            for i in resources
        ]


class OSPolicyAssignmentInstanceFilter(object):
    def __init__(
        self,
        all: bool = None,
        inclusion_labels: list = None,
        exclusion_labels: list = None,
        inventories: list = None,
    ):
        self.all = all
        self.inclusion_labels = inclusion_labels
        self.exclusion_labels = exclusion_labels
        self.inventories = inventories

    @classmethod
    def to_proto(self, resource):
        if not resource:
            return None

        res = os_policy_assignment_pb2.OsconfigBetaOSPolicyAssignmentInstanceFilter()
        if Primitive.to_proto(resource.all):
            res.all = Primitive.to_proto(resource.all)
        if OSPolicyAssignmentInstanceFilterInclusionLabelsArray.to_proto(
            resource.inclusion_labels
        ):
            res.inclusion_labels.extend(
                OSPolicyAssignmentInstanceFilterInclusionLabelsArray.to_proto(
                    resource.inclusion_labels
                )
            )
        if OSPolicyAssignmentInstanceFilterExclusionLabelsArray.to_proto(
            resource.exclusion_labels
        ):
            res.exclusion_labels.extend(
                OSPolicyAssignmentInstanceFilterExclusionLabelsArray.to_proto(
                    resource.exclusion_labels
                )
            )
        if OSPolicyAssignmentInstanceFilterInventoriesArray.to_proto(
            resource.inventories
        ):
            res.inventories.extend(
                OSPolicyAssignmentInstanceFilterInventoriesArray.to_proto(
                    resource.inventories
                )
            )
        return res

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return None

        return OSPolicyAssignmentInstanceFilter(
            all=Primitive.from_proto(resource.all),
            inclusion_labels=OSPolicyAssignmentInstanceFilterInclusionLabelsArray.from_proto(
                resource.inclusion_labels
            ),
            exclusion_labels=OSPolicyAssignmentInstanceFilterExclusionLabelsArray.from_proto(
                resource.exclusion_labels
            ),
            inventories=OSPolicyAssignmentInstanceFilterInventoriesArray.from_proto(
                resource.inventories
            ),
        )


class OSPolicyAssignmentInstanceFilterArray(object):
    @classmethod
    def to_proto(self, resources):
        if not resources:
            return resources
        return [OSPolicyAssignmentInstanceFilter.to_proto(i) for i in resources]

    @classmethod
    def from_proto(self, resources):
        return [OSPolicyAssignmentInstanceFilter.from_proto(i) for i in resources]


class OSPolicyAssignmentInstanceFilterInclusionLabels(object):
    def __init__(self, labels: dict = None):
        self.labels = labels

    @classmethod
    def to_proto(self, resource):
        if not resource:
            return None

        res = (
            os_policy_assignment_pb2.OsconfigBetaOSPolicyAssignmentInstanceFilterInclusionLabels()
        )
        if Primitive.to_proto(resource.labels):
            res.labels = Primitive.to_proto(resource.labels)
        return res

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return None

        return OSPolicyAssignmentInstanceFilterInclusionLabels(
            labels=Primitive.from_proto(resource.labels),
        )


class OSPolicyAssignmentInstanceFilterInclusionLabelsArray(object):
    @classmethod
    def to_proto(self, resources):
        if not resources:
            return resources
        return [
            OSPolicyAssignmentInstanceFilterInclusionLabels.to_proto(i)
            for i in resources
        ]

    @classmethod
    def from_proto(self, resources):
        return [
            OSPolicyAssignmentInstanceFilterInclusionLabels.from_proto(i)
            for i in resources
        ]


class OSPolicyAssignmentInstanceFilterExclusionLabels(object):
    def __init__(self, labels: dict = None):
        self.labels = labels

    @classmethod
    def to_proto(self, resource):
        if not resource:
            return None

        res = (
            os_policy_assignment_pb2.OsconfigBetaOSPolicyAssignmentInstanceFilterExclusionLabels()
        )
        if Primitive.to_proto(resource.labels):
            res.labels = Primitive.to_proto(resource.labels)
        return res

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return None

        return OSPolicyAssignmentInstanceFilterExclusionLabels(
            labels=Primitive.from_proto(resource.labels),
        )


class OSPolicyAssignmentInstanceFilterExclusionLabelsArray(object):
    @classmethod
    def to_proto(self, resources):
        if not resources:
            return resources
        return [
            OSPolicyAssignmentInstanceFilterExclusionLabels.to_proto(i)
            for i in resources
        ]

    @classmethod
    def from_proto(self, resources):
        return [
            OSPolicyAssignmentInstanceFilterExclusionLabels.from_proto(i)
            for i in resources
        ]


class OSPolicyAssignmentInstanceFilterInventories(object):
    def __init__(self, os_short_name: str = None, os_version: str = None):
        self.os_short_name = os_short_name
        self.os_version = os_version

    @classmethod
    def to_proto(self, resource):
        if not resource:
            return None

        res = (
            os_policy_assignment_pb2.OsconfigBetaOSPolicyAssignmentInstanceFilterInventories()
        )
        if Primitive.to_proto(resource.os_short_name):
            res.os_short_name = Primitive.to_proto(resource.os_short_name)
        if Primitive.to_proto(resource.os_version):
            res.os_version = Primitive.to_proto(resource.os_version)
        return res

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return None

        return OSPolicyAssignmentInstanceFilterInventories(
            os_short_name=Primitive.from_proto(resource.os_short_name),
            os_version=Primitive.from_proto(resource.os_version),
        )


class OSPolicyAssignmentInstanceFilterInventoriesArray(object):
    @classmethod
    def to_proto(self, resources):
        if not resources:
            return resources
        return [
            OSPolicyAssignmentInstanceFilterInventories.to_proto(i) for i in resources
        ]

    @classmethod
    def from_proto(self, resources):
        return [
            OSPolicyAssignmentInstanceFilterInventories.from_proto(i) for i in resources
        ]


class OSPolicyAssignmentRollout(object):
    def __init__(self, disruption_budget: dict = None, min_wait_duration: str = None):
        self.disruption_budget = disruption_budget
        self.min_wait_duration = min_wait_duration

    @classmethod
    def to_proto(self, resource):
        if not resource:
            return None

        res = os_policy_assignment_pb2.OsconfigBetaOSPolicyAssignmentRollout()
        if OSPolicyAssignmentRolloutDisruptionBudget.to_proto(
            resource.disruption_budget
        ):
            res.disruption_budget.CopyFrom(
                OSPolicyAssignmentRolloutDisruptionBudget.to_proto(
                    resource.disruption_budget
                )
            )
        else:
            res.ClearField("disruption_budget")
        if Primitive.to_proto(resource.min_wait_duration):
            res.min_wait_duration = Primitive.to_proto(resource.min_wait_duration)
        return res

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return None

        return OSPolicyAssignmentRollout(
            disruption_budget=OSPolicyAssignmentRolloutDisruptionBudget.from_proto(
                resource.disruption_budget
            ),
            min_wait_duration=Primitive.from_proto(resource.min_wait_duration),
        )


class OSPolicyAssignmentRolloutArray(object):
    @classmethod
    def to_proto(self, resources):
        if not resources:
            return resources
        return [OSPolicyAssignmentRollout.to_proto(i) for i in resources]

    @classmethod
    def from_proto(self, resources):
        return [OSPolicyAssignmentRollout.from_proto(i) for i in resources]


class OSPolicyAssignmentRolloutDisruptionBudget(object):
    def __init__(self, fixed: int = None, percent: int = None):
        self.fixed = fixed
        self.percent = percent

    @classmethod
    def to_proto(self, resource):
        if not resource:
            return None

        res = (
            os_policy_assignment_pb2.OsconfigBetaOSPolicyAssignmentRolloutDisruptionBudget()
        )
        if Primitive.to_proto(resource.fixed):
            res.fixed = Primitive.to_proto(resource.fixed)
        if Primitive.to_proto(resource.percent):
            res.percent = Primitive.to_proto(resource.percent)
        return res

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return None

        return OSPolicyAssignmentRolloutDisruptionBudget(
            fixed=Primitive.from_proto(resource.fixed),
            percent=Primitive.from_proto(resource.percent),
        )


class OSPolicyAssignmentRolloutDisruptionBudgetArray(object):
    @classmethod
    def to_proto(self, resources):
        if not resources:
            return resources
        return [
            OSPolicyAssignmentRolloutDisruptionBudget.to_proto(i) for i in resources
        ]

    @classmethod
    def from_proto(self, resources):
        return [
            OSPolicyAssignmentRolloutDisruptionBudget.from_proto(i) for i in resources
        ]


class OSPolicyAssignmentOSPoliciesModeEnum(object):
    @classmethod
    def to_proto(self, resource):
        if not resource:
            return resource
        return os_policy_assignment_pb2.OsconfigBetaOSPolicyAssignmentOSPoliciesModeEnum.Value(
            "OsconfigBetaOSPolicyAssignmentOSPoliciesModeEnum%s" % resource
        )

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return resource
        return os_policy_assignment_pb2.OsconfigBetaOSPolicyAssignmentOSPoliciesModeEnum.Name(
            resource
        )[
            len("OsconfigBetaOSPolicyAssignmentOSPoliciesModeEnum") :
        ]


class OSPolicyAssignmentOSPoliciesResourceGroupsResourcesPkgDesiredStateEnum(object):
    @classmethod
    def to_proto(self, resource):
        if not resource:
            return resource
        return os_policy_assignment_pb2.OsconfigBetaOSPolicyAssignmentOSPoliciesResourceGroupsResourcesPkgDesiredStateEnum.Value(
            "OsconfigBetaOSPolicyAssignmentOSPoliciesResourceGroupsResourcesPkgDesiredStateEnum%s"
            % resource
        )

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return resource
        return os_policy_assignment_pb2.OsconfigBetaOSPolicyAssignmentOSPoliciesResourceGroupsResourcesPkgDesiredStateEnum.Name(
            resource
        )[
            len(
                "OsconfigBetaOSPolicyAssignmentOSPoliciesResourceGroupsResourcesPkgDesiredStateEnum"
            ) :
        ]


class OSPolicyAssignmentOSPoliciesResourceGroupsResourcesRepositoryAptArchiveTypeEnum(
    object
):
    @classmethod
    def to_proto(self, resource):
        if not resource:
            return resource
        return os_policy_assignment_pb2.OsconfigBetaOSPolicyAssignmentOSPoliciesResourceGroupsResourcesRepositoryAptArchiveTypeEnum.Value(
            "OsconfigBetaOSPolicyAssignmentOSPoliciesResourceGroupsResourcesRepositoryAptArchiveTypeEnum%s"
            % resource
        )

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return resource
        return os_policy_assignment_pb2.OsconfigBetaOSPolicyAssignmentOSPoliciesResourceGroupsResourcesRepositoryAptArchiveTypeEnum.Name(
            resource
        )[
            len(
                "OsconfigBetaOSPolicyAssignmentOSPoliciesResourceGroupsResourcesRepositoryAptArchiveTypeEnum"
            ) :
        ]


class OSPolicyAssignmentOSPoliciesResourceGroupsResourcesExecValidateInterpreterEnum(
    object
):
    @classmethod
    def to_proto(self, resource):
        if not resource:
            return resource
        return os_policy_assignment_pb2.OsconfigBetaOSPolicyAssignmentOSPoliciesResourceGroupsResourcesExecValidateInterpreterEnum.Value(
            "OsconfigBetaOSPolicyAssignmentOSPoliciesResourceGroupsResourcesExecValidateInterpreterEnum%s"
            % resource
        )

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return resource
        return os_policy_assignment_pb2.OsconfigBetaOSPolicyAssignmentOSPoliciesResourceGroupsResourcesExecValidateInterpreterEnum.Name(
            resource
        )[
            len(
                "OsconfigBetaOSPolicyAssignmentOSPoliciesResourceGroupsResourcesExecValidateInterpreterEnum"
            ) :
        ]


class OSPolicyAssignmentOSPoliciesResourceGroupsResourcesExecEnforceInterpreterEnum(
    object
):
    @classmethod
    def to_proto(self, resource):
        if not resource:
            return resource
        return os_policy_assignment_pb2.OsconfigBetaOSPolicyAssignmentOSPoliciesResourceGroupsResourcesExecEnforceInterpreterEnum.Value(
            "OsconfigBetaOSPolicyAssignmentOSPoliciesResourceGroupsResourcesExecEnforceInterpreterEnum%s"
            % resource
        )

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return resource
        return os_policy_assignment_pb2.OsconfigBetaOSPolicyAssignmentOSPoliciesResourceGroupsResourcesExecEnforceInterpreterEnum.Name(
            resource
        )[
            len(
                "OsconfigBetaOSPolicyAssignmentOSPoliciesResourceGroupsResourcesExecEnforceInterpreterEnum"
            ) :
        ]


class OSPolicyAssignmentOSPoliciesResourceGroupsResourcesFileStateEnum(object):
    @classmethod
    def to_proto(self, resource):
        if not resource:
            return resource
        return os_policy_assignment_pb2.OsconfigBetaOSPolicyAssignmentOSPoliciesResourceGroupsResourcesFileStateEnum.Value(
            "OsconfigBetaOSPolicyAssignmentOSPoliciesResourceGroupsResourcesFileStateEnum%s"
            % resource
        )

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return resource
        return os_policy_assignment_pb2.OsconfigBetaOSPolicyAssignmentOSPoliciesResourceGroupsResourcesFileStateEnum.Name(
            resource
        )[
            len(
                "OsconfigBetaOSPolicyAssignmentOSPoliciesResourceGroupsResourcesFileStateEnum"
            ) :
        ]


class OSPolicyAssignmentRolloutStateEnum(object):
    @classmethod
    def to_proto(self, resource):
        if not resource:
            return resource
        return os_policy_assignment_pb2.OsconfigBetaOSPolicyAssignmentRolloutStateEnum.Value(
            "OsconfigBetaOSPolicyAssignmentRolloutStateEnum%s" % resource
        )

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return resource
        return os_policy_assignment_pb2.OsconfigBetaOSPolicyAssignmentRolloutStateEnum.Name(
            resource
        )[
            len("OsconfigBetaOSPolicyAssignmentRolloutStateEnum") :
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
