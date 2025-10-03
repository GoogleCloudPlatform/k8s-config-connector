# Copyright 2022 Google LLC. All Rights Reserved.
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
from google3.cloud.graphite.mmv2.services.google.os_config import patch_deployment_pb2
from google3.cloud.graphite.mmv2.services.google.os_config import (
    patch_deployment_pb2_grpc,
)

from typing import List


class PatchDeployment(object):
    def __init__(
        self,
        name: str = None,
        description: str = None,
        instance_filter: dict = None,
        patch_config: dict = None,
        duration: str = None,
        one_time_schedule: dict = None,
        recurring_schedule: dict = None,
        create_time: str = None,
        update_time: str = None,
        last_execute_time: str = None,
        rollout: dict = None,
        project: str = None,
        service_account_file: str = "",
    ):

        channel.initialize()
        self.name = name
        self.description = description
        self.instance_filter = instance_filter
        self.patch_config = patch_config
        self.duration = duration
        self.one_time_schedule = one_time_schedule
        self.recurring_schedule = recurring_schedule
        self.rollout = rollout
        self.project = project
        self.service_account_file = service_account_file

    def apply(self):
        stub = patch_deployment_pb2_grpc.OsconfigBetaPatchDeploymentServiceStub(
            channel.Channel()
        )
        request = patch_deployment_pb2.ApplyOsconfigBetaPatchDeploymentRequest()
        if Primitive.to_proto(self.name):
            request.resource.name = Primitive.to_proto(self.name)

        if Primitive.to_proto(self.description):
            request.resource.description = Primitive.to_proto(self.description)

        if PatchDeploymentInstanceFilter.to_proto(self.instance_filter):
            request.resource.instance_filter.CopyFrom(
                PatchDeploymentInstanceFilter.to_proto(self.instance_filter)
            )
        else:
            request.resource.ClearField("instance_filter")
        if PatchDeploymentPatchConfig.to_proto(self.patch_config):
            request.resource.patch_config.CopyFrom(
                PatchDeploymentPatchConfig.to_proto(self.patch_config)
            )
        else:
            request.resource.ClearField("patch_config")
        if Primitive.to_proto(self.duration):
            request.resource.duration = Primitive.to_proto(self.duration)

        if PatchDeploymentOneTimeSchedule.to_proto(self.one_time_schedule):
            request.resource.one_time_schedule.CopyFrom(
                PatchDeploymentOneTimeSchedule.to_proto(self.one_time_schedule)
            )
        else:
            request.resource.ClearField("one_time_schedule")
        if PatchDeploymentRecurringSchedule.to_proto(self.recurring_schedule):
            request.resource.recurring_schedule.CopyFrom(
                PatchDeploymentRecurringSchedule.to_proto(self.recurring_schedule)
            )
        else:
            request.resource.ClearField("recurring_schedule")
        if PatchDeploymentRollout.to_proto(self.rollout):
            request.resource.rollout.CopyFrom(
                PatchDeploymentRollout.to_proto(self.rollout)
            )
        else:
            request.resource.ClearField("rollout")
        if Primitive.to_proto(self.project):
            request.resource.project = Primitive.to_proto(self.project)

        request.service_account_file = self.service_account_file

        response = stub.ApplyOsconfigBetaPatchDeployment(request)
        self.name = Primitive.from_proto(response.name)
        self.description = Primitive.from_proto(response.description)
        self.instance_filter = PatchDeploymentInstanceFilter.from_proto(
            response.instance_filter
        )
        self.patch_config = PatchDeploymentPatchConfig.from_proto(response.patch_config)
        self.duration = Primitive.from_proto(response.duration)
        self.one_time_schedule = PatchDeploymentOneTimeSchedule.from_proto(
            response.one_time_schedule
        )
        self.recurring_schedule = PatchDeploymentRecurringSchedule.from_proto(
            response.recurring_schedule
        )
        self.create_time = Primitive.from_proto(response.create_time)
        self.update_time = Primitive.from_proto(response.update_time)
        self.last_execute_time = Primitive.from_proto(response.last_execute_time)
        self.rollout = PatchDeploymentRollout.from_proto(response.rollout)
        self.project = Primitive.from_proto(response.project)

    def delete(self):
        stub = patch_deployment_pb2_grpc.OsconfigBetaPatchDeploymentServiceStub(
            channel.Channel()
        )
        request = patch_deployment_pb2.DeleteOsconfigBetaPatchDeploymentRequest()
        request.service_account_file = self.service_account_file
        if Primitive.to_proto(self.name):
            request.resource.name = Primitive.to_proto(self.name)

        if Primitive.to_proto(self.description):
            request.resource.description = Primitive.to_proto(self.description)

        if PatchDeploymentInstanceFilter.to_proto(self.instance_filter):
            request.resource.instance_filter.CopyFrom(
                PatchDeploymentInstanceFilter.to_proto(self.instance_filter)
            )
        else:
            request.resource.ClearField("instance_filter")
        if PatchDeploymentPatchConfig.to_proto(self.patch_config):
            request.resource.patch_config.CopyFrom(
                PatchDeploymentPatchConfig.to_proto(self.patch_config)
            )
        else:
            request.resource.ClearField("patch_config")
        if Primitive.to_proto(self.duration):
            request.resource.duration = Primitive.to_proto(self.duration)

        if PatchDeploymentOneTimeSchedule.to_proto(self.one_time_schedule):
            request.resource.one_time_schedule.CopyFrom(
                PatchDeploymentOneTimeSchedule.to_proto(self.one_time_schedule)
            )
        else:
            request.resource.ClearField("one_time_schedule")
        if PatchDeploymentRecurringSchedule.to_proto(self.recurring_schedule):
            request.resource.recurring_schedule.CopyFrom(
                PatchDeploymentRecurringSchedule.to_proto(self.recurring_schedule)
            )
        else:
            request.resource.ClearField("recurring_schedule")
        if PatchDeploymentRollout.to_proto(self.rollout):
            request.resource.rollout.CopyFrom(
                PatchDeploymentRollout.to_proto(self.rollout)
            )
        else:
            request.resource.ClearField("rollout")
        if Primitive.to_proto(self.project):
            request.resource.project = Primitive.to_proto(self.project)

        response = stub.DeleteOsconfigBetaPatchDeployment(request)

    @classmethod
    def list(self, project, service_account_file=""):
        stub = patch_deployment_pb2_grpc.OsconfigBetaPatchDeploymentServiceStub(
            channel.Channel()
        )
        request = patch_deployment_pb2.ListOsconfigBetaPatchDeploymentRequest()
        request.service_account_file = service_account_file
        request.Project = project

        return stub.ListOsconfigBetaPatchDeployment(request).items

    def to_proto(self):
        resource = patch_deployment_pb2.OsconfigBetaPatchDeployment()
        if Primitive.to_proto(self.name):
            resource.name = Primitive.to_proto(self.name)
        if Primitive.to_proto(self.description):
            resource.description = Primitive.to_proto(self.description)
        if PatchDeploymentInstanceFilter.to_proto(self.instance_filter):
            resource.instance_filter.CopyFrom(
                PatchDeploymentInstanceFilter.to_proto(self.instance_filter)
            )
        else:
            resource.ClearField("instance_filter")
        if PatchDeploymentPatchConfig.to_proto(self.patch_config):
            resource.patch_config.CopyFrom(
                PatchDeploymentPatchConfig.to_proto(self.patch_config)
            )
        else:
            resource.ClearField("patch_config")
        if Primitive.to_proto(self.duration):
            resource.duration = Primitive.to_proto(self.duration)
        if PatchDeploymentOneTimeSchedule.to_proto(self.one_time_schedule):
            resource.one_time_schedule.CopyFrom(
                PatchDeploymentOneTimeSchedule.to_proto(self.one_time_schedule)
            )
        else:
            resource.ClearField("one_time_schedule")
        if PatchDeploymentRecurringSchedule.to_proto(self.recurring_schedule):
            resource.recurring_schedule.CopyFrom(
                PatchDeploymentRecurringSchedule.to_proto(self.recurring_schedule)
            )
        else:
            resource.ClearField("recurring_schedule")
        if PatchDeploymentRollout.to_proto(self.rollout):
            resource.rollout.CopyFrom(PatchDeploymentRollout.to_proto(self.rollout))
        else:
            resource.ClearField("rollout")
        if Primitive.to_proto(self.project):
            resource.project = Primitive.to_proto(self.project)
        return resource


class PatchDeploymentInstanceFilter(object):
    def __init__(
        self,
        all: bool = None,
        group_labels: list = None,
        zones: list = None,
        instances: list = None,
        instance_name_prefixes: list = None,
    ):
        self.all = all
        self.group_labels = group_labels
        self.zones = zones
        self.instances = instances
        self.instance_name_prefixes = instance_name_prefixes

    @classmethod
    def to_proto(self, resource):
        if not resource:
            return None

        res = patch_deployment_pb2.OsconfigBetaPatchDeploymentInstanceFilter()
        if Primitive.to_proto(resource.all):
            res.all = Primitive.to_proto(resource.all)
        if PatchDeploymentInstanceFilterGroupLabelsArray.to_proto(
            resource.group_labels
        ):
            res.group_labels.extend(
                PatchDeploymentInstanceFilterGroupLabelsArray.to_proto(
                    resource.group_labels
                )
            )
        if Primitive.to_proto(resource.zones):
            res.zones.extend(Primitive.to_proto(resource.zones))
        if Primitive.to_proto(resource.instances):
            res.instances.extend(Primitive.to_proto(resource.instances))
        if Primitive.to_proto(resource.instance_name_prefixes):
            res.instance_name_prefixes.extend(
                Primitive.to_proto(resource.instance_name_prefixes)
            )
        return res

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return None

        return PatchDeploymentInstanceFilter(
            all=Primitive.from_proto(resource.all),
            group_labels=PatchDeploymentInstanceFilterGroupLabelsArray.from_proto(
                resource.group_labels
            ),
            zones=Primitive.from_proto(resource.zones),
            instances=Primitive.from_proto(resource.instances),
            instance_name_prefixes=Primitive.from_proto(
                resource.instance_name_prefixes
            ),
        )


class PatchDeploymentInstanceFilterArray(object):
    @classmethod
    def to_proto(self, resources):
        if not resources:
            return resources
        return [PatchDeploymentInstanceFilter.to_proto(i) for i in resources]

    @classmethod
    def from_proto(self, resources):
        return [PatchDeploymentInstanceFilter.from_proto(i) for i in resources]


class PatchDeploymentInstanceFilterGroupLabels(object):
    def __init__(self, labels: dict = None):
        self.labels = labels

    @classmethod
    def to_proto(self, resource):
        if not resource:
            return None

        res = (
            patch_deployment_pb2.OsconfigBetaPatchDeploymentInstanceFilterGroupLabels()
        )
        if Primitive.to_proto(resource.labels):
            res.labels = Primitive.to_proto(resource.labels)
        return res

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return None

        return PatchDeploymentInstanceFilterGroupLabels(
            labels=Primitive.from_proto(resource.labels),
        )


class PatchDeploymentInstanceFilterGroupLabelsArray(object):
    @classmethod
    def to_proto(self, resources):
        if not resources:
            return resources
        return [PatchDeploymentInstanceFilterGroupLabels.to_proto(i) for i in resources]

    @classmethod
    def from_proto(self, resources):
        return [
            PatchDeploymentInstanceFilterGroupLabels.from_proto(i) for i in resources
        ]


class PatchDeploymentPatchConfig(object):
    def __init__(
        self,
        reboot_config: str = None,
        apt: dict = None,
        yum: dict = None,
        goo: dict = None,
        zypper: dict = None,
        windows_update: dict = None,
        pre_step: dict = None,
        post_step: dict = None,
    ):
        self.reboot_config = reboot_config
        self.apt = apt
        self.yum = yum
        self.goo = goo
        self.zypper = zypper
        self.windows_update = windows_update
        self.pre_step = pre_step
        self.post_step = post_step

    @classmethod
    def to_proto(self, resource):
        if not resource:
            return None

        res = patch_deployment_pb2.OsconfigBetaPatchDeploymentPatchConfig()
        if PatchDeploymentPatchConfigRebootConfigEnum.to_proto(resource.reboot_config):
            res.reboot_config = PatchDeploymentPatchConfigRebootConfigEnum.to_proto(
                resource.reboot_config
            )
        if PatchDeploymentPatchConfigApt.to_proto(resource.apt):
            res.apt.CopyFrom(PatchDeploymentPatchConfigApt.to_proto(resource.apt))
        else:
            res.ClearField("apt")
        if PatchDeploymentPatchConfigYum.to_proto(resource.yum):
            res.yum.CopyFrom(PatchDeploymentPatchConfigYum.to_proto(resource.yum))
        else:
            res.ClearField("yum")
        if PatchDeploymentPatchConfigGoo.to_proto(resource.goo):
            res.goo.CopyFrom(PatchDeploymentPatchConfigGoo.to_proto(resource.goo))
        else:
            res.ClearField("goo")
        if PatchDeploymentPatchConfigZypper.to_proto(resource.zypper):
            res.zypper.CopyFrom(
                PatchDeploymentPatchConfigZypper.to_proto(resource.zypper)
            )
        else:
            res.ClearField("zypper")
        if PatchDeploymentPatchConfigWindowsUpdate.to_proto(resource.windows_update):
            res.windows_update.CopyFrom(
                PatchDeploymentPatchConfigWindowsUpdate.to_proto(
                    resource.windows_update
                )
            )
        else:
            res.ClearField("windows_update")
        if PatchDeploymentPatchConfigPreStep.to_proto(resource.pre_step):
            res.pre_step.CopyFrom(
                PatchDeploymentPatchConfigPreStep.to_proto(resource.pre_step)
            )
        else:
            res.ClearField("pre_step")
        if PatchDeploymentPatchConfigPostStep.to_proto(resource.post_step):
            res.post_step.CopyFrom(
                PatchDeploymentPatchConfigPostStep.to_proto(resource.post_step)
            )
        else:
            res.ClearField("post_step")
        return res

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return None

        return PatchDeploymentPatchConfig(
            reboot_config=PatchDeploymentPatchConfigRebootConfigEnum.from_proto(
                resource.reboot_config
            ),
            apt=PatchDeploymentPatchConfigApt.from_proto(resource.apt),
            yum=PatchDeploymentPatchConfigYum.from_proto(resource.yum),
            goo=PatchDeploymentPatchConfigGoo.from_proto(resource.goo),
            zypper=PatchDeploymentPatchConfigZypper.from_proto(resource.zypper),
            windows_update=PatchDeploymentPatchConfigWindowsUpdate.from_proto(
                resource.windows_update
            ),
            pre_step=PatchDeploymentPatchConfigPreStep.from_proto(resource.pre_step),
            post_step=PatchDeploymentPatchConfigPostStep.from_proto(resource.post_step),
        )


class PatchDeploymentPatchConfigArray(object):
    @classmethod
    def to_proto(self, resources):
        if not resources:
            return resources
        return [PatchDeploymentPatchConfig.to_proto(i) for i in resources]

    @classmethod
    def from_proto(self, resources):
        return [PatchDeploymentPatchConfig.from_proto(i) for i in resources]


class PatchDeploymentPatchConfigApt(object):
    def __init__(
        self, type: str = None, excludes: list = None, exclusive_packages: list = None
    ):
        self.type = type
        self.excludes = excludes
        self.exclusive_packages = exclusive_packages

    @classmethod
    def to_proto(self, resource):
        if not resource:
            return None

        res = patch_deployment_pb2.OsconfigBetaPatchDeploymentPatchConfigApt()
        if PatchDeploymentPatchConfigAptTypeEnum.to_proto(resource.type):
            res.type = PatchDeploymentPatchConfigAptTypeEnum.to_proto(resource.type)
        if Primitive.to_proto(resource.excludes):
            res.excludes.extend(Primitive.to_proto(resource.excludes))
        if Primitive.to_proto(resource.exclusive_packages):
            res.exclusive_packages.extend(
                Primitive.to_proto(resource.exclusive_packages)
            )
        return res

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return None

        return PatchDeploymentPatchConfigApt(
            type=PatchDeploymentPatchConfigAptTypeEnum.from_proto(resource.type),
            excludes=Primitive.from_proto(resource.excludes),
            exclusive_packages=Primitive.from_proto(resource.exclusive_packages),
        )


class PatchDeploymentPatchConfigAptArray(object):
    @classmethod
    def to_proto(self, resources):
        if not resources:
            return resources
        return [PatchDeploymentPatchConfigApt.to_proto(i) for i in resources]

    @classmethod
    def from_proto(self, resources):
        return [PatchDeploymentPatchConfigApt.from_proto(i) for i in resources]


class PatchDeploymentPatchConfigYum(object):
    def __init__(
        self,
        security: bool = None,
        minimal: bool = None,
        excludes: list = None,
        exclusive_packages: list = None,
    ):
        self.security = security
        self.minimal = minimal
        self.excludes = excludes
        self.exclusive_packages = exclusive_packages

    @classmethod
    def to_proto(self, resource):
        if not resource:
            return None

        res = patch_deployment_pb2.OsconfigBetaPatchDeploymentPatchConfigYum()
        if Primitive.to_proto(resource.security):
            res.security = Primitive.to_proto(resource.security)
        if Primitive.to_proto(resource.minimal):
            res.minimal = Primitive.to_proto(resource.minimal)
        if Primitive.to_proto(resource.excludes):
            res.excludes.extend(Primitive.to_proto(resource.excludes))
        if Primitive.to_proto(resource.exclusive_packages):
            res.exclusive_packages.extend(
                Primitive.to_proto(resource.exclusive_packages)
            )
        return res

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return None

        return PatchDeploymentPatchConfigYum(
            security=Primitive.from_proto(resource.security),
            minimal=Primitive.from_proto(resource.minimal),
            excludes=Primitive.from_proto(resource.excludes),
            exclusive_packages=Primitive.from_proto(resource.exclusive_packages),
        )


class PatchDeploymentPatchConfigYumArray(object):
    @classmethod
    def to_proto(self, resources):
        if not resources:
            return resources
        return [PatchDeploymentPatchConfigYum.to_proto(i) for i in resources]

    @classmethod
    def from_proto(self, resources):
        return [PatchDeploymentPatchConfigYum.from_proto(i) for i in resources]


class PatchDeploymentPatchConfigGoo(object):
    @classmethod
    def to_proto(self, resource):
        if not resource:
            return None

        res = patch_deployment_pb2.OsconfigBetaPatchDeploymentPatchConfigGoo()
        return res

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return None

        return PatchDeploymentPatchConfigGoo()


class PatchDeploymentPatchConfigGooArray(object):
    @classmethod
    def to_proto(self, resources):
        if not resources:
            return resources
        return [PatchDeploymentPatchConfigGoo.to_proto(i) for i in resources]

    @classmethod
    def from_proto(self, resources):
        return [PatchDeploymentPatchConfigGoo.from_proto(i) for i in resources]


class PatchDeploymentPatchConfigZypper(object):
    def __init__(
        self,
        with_optional: bool = None,
        with_update: bool = None,
        categories: list = None,
        severities: list = None,
        excludes: list = None,
        exclusive_patches: list = None,
    ):
        self.with_optional = with_optional
        self.with_update = with_update
        self.categories = categories
        self.severities = severities
        self.excludes = excludes
        self.exclusive_patches = exclusive_patches

    @classmethod
    def to_proto(self, resource):
        if not resource:
            return None

        res = patch_deployment_pb2.OsconfigBetaPatchDeploymentPatchConfigZypper()
        if Primitive.to_proto(resource.with_optional):
            res.with_optional = Primitive.to_proto(resource.with_optional)
        if Primitive.to_proto(resource.with_update):
            res.with_update = Primitive.to_proto(resource.with_update)
        if Primitive.to_proto(resource.categories):
            res.categories.extend(Primitive.to_proto(resource.categories))
        if Primitive.to_proto(resource.severities):
            res.severities.extend(Primitive.to_proto(resource.severities))
        if Primitive.to_proto(resource.excludes):
            res.excludes.extend(Primitive.to_proto(resource.excludes))
        if Primitive.to_proto(resource.exclusive_patches):
            res.exclusive_patches.extend(Primitive.to_proto(resource.exclusive_patches))
        return res

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return None

        return PatchDeploymentPatchConfigZypper(
            with_optional=Primitive.from_proto(resource.with_optional),
            with_update=Primitive.from_proto(resource.with_update),
            categories=Primitive.from_proto(resource.categories),
            severities=Primitive.from_proto(resource.severities),
            excludes=Primitive.from_proto(resource.excludes),
            exclusive_patches=Primitive.from_proto(resource.exclusive_patches),
        )


class PatchDeploymentPatchConfigZypperArray(object):
    @classmethod
    def to_proto(self, resources):
        if not resources:
            return resources
        return [PatchDeploymentPatchConfigZypper.to_proto(i) for i in resources]

    @classmethod
    def from_proto(self, resources):
        return [PatchDeploymentPatchConfigZypper.from_proto(i) for i in resources]


class PatchDeploymentPatchConfigWindowsUpdate(object):
    def __init__(
        self,
        classifications: list = None,
        excludes: list = None,
        exclusive_patches: list = None,
    ):
        self.classifications = classifications
        self.excludes = excludes
        self.exclusive_patches = exclusive_patches

    @classmethod
    def to_proto(self, resource):
        if not resource:
            return None

        res = patch_deployment_pb2.OsconfigBetaPatchDeploymentPatchConfigWindowsUpdate()
        if PatchDeploymentPatchConfigWindowsUpdateClassificationsEnumArray.to_proto(
            resource.classifications
        ):
            res.classifications.extend(
                PatchDeploymentPatchConfigWindowsUpdateClassificationsEnumArray.to_proto(
                    resource.classifications
                )
            )
        if Primitive.to_proto(resource.excludes):
            res.excludes.extend(Primitive.to_proto(resource.excludes))
        if Primitive.to_proto(resource.exclusive_patches):
            res.exclusive_patches.extend(Primitive.to_proto(resource.exclusive_patches))
        return res

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return None

        return PatchDeploymentPatchConfigWindowsUpdate(
            classifications=PatchDeploymentPatchConfigWindowsUpdateClassificationsEnumArray.from_proto(
                resource.classifications
            ),
            excludes=Primitive.from_proto(resource.excludes),
            exclusive_patches=Primitive.from_proto(resource.exclusive_patches),
        )


class PatchDeploymentPatchConfigWindowsUpdateArray(object):
    @classmethod
    def to_proto(self, resources):
        if not resources:
            return resources
        return [PatchDeploymentPatchConfigWindowsUpdate.to_proto(i) for i in resources]

    @classmethod
    def from_proto(self, resources):
        return [
            PatchDeploymentPatchConfigWindowsUpdate.from_proto(i) for i in resources
        ]


class PatchDeploymentPatchConfigPreStep(object):
    def __init__(
        self, linux_exec_step_config: dict = None, windows_exec_step_config: dict = None
    ):
        self.linux_exec_step_config = linux_exec_step_config
        self.windows_exec_step_config = windows_exec_step_config

    @classmethod
    def to_proto(self, resource):
        if not resource:
            return None

        res = patch_deployment_pb2.OsconfigBetaPatchDeploymentPatchConfigPreStep()
        if PatchDeploymentPatchConfigPreStepLinuxExecStepConfig.to_proto(
            resource.linux_exec_step_config
        ):
            res.linux_exec_step_config.CopyFrom(
                PatchDeploymentPatchConfigPreStepLinuxExecStepConfig.to_proto(
                    resource.linux_exec_step_config
                )
            )
        else:
            res.ClearField("linux_exec_step_config")
        if PatchDeploymentPatchConfigPreStepWindowsExecStepConfig.to_proto(
            resource.windows_exec_step_config
        ):
            res.windows_exec_step_config.CopyFrom(
                PatchDeploymentPatchConfigPreStepWindowsExecStepConfig.to_proto(
                    resource.windows_exec_step_config
                )
            )
        else:
            res.ClearField("windows_exec_step_config")
        return res

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return None

        return PatchDeploymentPatchConfigPreStep(
            linux_exec_step_config=PatchDeploymentPatchConfigPreStepLinuxExecStepConfig.from_proto(
                resource.linux_exec_step_config
            ),
            windows_exec_step_config=PatchDeploymentPatchConfigPreStepWindowsExecStepConfig.from_proto(
                resource.windows_exec_step_config
            ),
        )


class PatchDeploymentPatchConfigPreStepArray(object):
    @classmethod
    def to_proto(self, resources):
        if not resources:
            return resources
        return [PatchDeploymentPatchConfigPreStep.to_proto(i) for i in resources]

    @classmethod
    def from_proto(self, resources):
        return [PatchDeploymentPatchConfigPreStep.from_proto(i) for i in resources]


class PatchDeploymentPatchConfigPreStepLinuxExecStepConfig(object):
    def __init__(
        self,
        local_path: str = None,
        gcs_object: dict = None,
        allowed_success_codes: list = None,
        interpreter: str = None,
    ):
        self.local_path = local_path
        self.gcs_object = gcs_object
        self.allowed_success_codes = allowed_success_codes
        self.interpreter = interpreter

    @classmethod
    def to_proto(self, resource):
        if not resource:
            return None

        res = (
            patch_deployment_pb2.OsconfigBetaPatchDeploymentPatchConfigPreStepLinuxExecStepConfig()
        )
        if Primitive.to_proto(resource.local_path):
            res.local_path = Primitive.to_proto(resource.local_path)
        if PatchDeploymentPatchConfigPreStepLinuxExecStepConfigGcsObject.to_proto(
            resource.gcs_object
        ):
            res.gcs_object.CopyFrom(
                PatchDeploymentPatchConfigPreStepLinuxExecStepConfigGcsObject.to_proto(
                    resource.gcs_object
                )
            )
        else:
            res.ClearField("gcs_object")
        if int64Array.to_proto(resource.allowed_success_codes):
            res.allowed_success_codes.extend(
                int64Array.to_proto(resource.allowed_success_codes)
            )
        if PatchDeploymentPatchConfigPreStepLinuxExecStepConfigInterpreterEnum.to_proto(
            resource.interpreter
        ):
            res.interpreter = PatchDeploymentPatchConfigPreStepLinuxExecStepConfigInterpreterEnum.to_proto(
                resource.interpreter
            )
        return res

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return None

        return PatchDeploymentPatchConfigPreStepLinuxExecStepConfig(
            local_path=Primitive.from_proto(resource.local_path),
            gcs_object=PatchDeploymentPatchConfigPreStepLinuxExecStepConfigGcsObject.from_proto(
                resource.gcs_object
            ),
            allowed_success_codes=int64Array.from_proto(resource.allowed_success_codes),
            interpreter=PatchDeploymentPatchConfigPreStepLinuxExecStepConfigInterpreterEnum.from_proto(
                resource.interpreter
            ),
        )


class PatchDeploymentPatchConfigPreStepLinuxExecStepConfigArray(object):
    @classmethod
    def to_proto(self, resources):
        if not resources:
            return resources
        return [
            PatchDeploymentPatchConfigPreStepLinuxExecStepConfig.to_proto(i)
            for i in resources
        ]

    @classmethod
    def from_proto(self, resources):
        return [
            PatchDeploymentPatchConfigPreStepLinuxExecStepConfig.from_proto(i)
            for i in resources
        ]


class PatchDeploymentPatchConfigPreStepLinuxExecStepConfigGcsObject(object):
    def __init__(
        self, bucket: str = None, object: str = None, generation_number: int = None
    ):
        self.bucket = bucket
        self.object = object
        self.generation_number = generation_number

    @classmethod
    def to_proto(self, resource):
        if not resource:
            return None

        res = (
            patch_deployment_pb2.OsconfigBetaPatchDeploymentPatchConfigPreStepLinuxExecStepConfigGcsObject()
        )
        if Primitive.to_proto(resource.bucket):
            res.bucket = Primitive.to_proto(resource.bucket)
        if Primitive.to_proto(resource.object):
            res.object = Primitive.to_proto(resource.object)
        if Primitive.to_proto(resource.generation_number):
            res.generation_number = Primitive.to_proto(resource.generation_number)
        return res

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return None

        return PatchDeploymentPatchConfigPreStepLinuxExecStepConfigGcsObject(
            bucket=Primitive.from_proto(resource.bucket),
            object=Primitive.from_proto(resource.object),
            generation_number=Primitive.from_proto(resource.generation_number),
        )


class PatchDeploymentPatchConfigPreStepLinuxExecStepConfigGcsObjectArray(object):
    @classmethod
    def to_proto(self, resources):
        if not resources:
            return resources
        return [
            PatchDeploymentPatchConfigPreStepLinuxExecStepConfigGcsObject.to_proto(i)
            for i in resources
        ]

    @classmethod
    def from_proto(self, resources):
        return [
            PatchDeploymentPatchConfigPreStepLinuxExecStepConfigGcsObject.from_proto(i)
            for i in resources
        ]


class PatchDeploymentPatchConfigPreStepWindowsExecStepConfig(object):
    def __init__(
        self,
        local_path: str = None,
        gcs_object: dict = None,
        allowed_success_codes: list = None,
        interpreter: str = None,
    ):
        self.local_path = local_path
        self.gcs_object = gcs_object
        self.allowed_success_codes = allowed_success_codes
        self.interpreter = interpreter

    @classmethod
    def to_proto(self, resource):
        if not resource:
            return None

        res = (
            patch_deployment_pb2.OsconfigBetaPatchDeploymentPatchConfigPreStepWindowsExecStepConfig()
        )
        if Primitive.to_proto(resource.local_path):
            res.local_path = Primitive.to_proto(resource.local_path)
        if PatchDeploymentPatchConfigPreStepWindowsExecStepConfigGcsObject.to_proto(
            resource.gcs_object
        ):
            res.gcs_object.CopyFrom(
                PatchDeploymentPatchConfigPreStepWindowsExecStepConfigGcsObject.to_proto(
                    resource.gcs_object
                )
            )
        else:
            res.ClearField("gcs_object")
        if int64Array.to_proto(resource.allowed_success_codes):
            res.allowed_success_codes.extend(
                int64Array.to_proto(resource.allowed_success_codes)
            )
        if PatchDeploymentPatchConfigPreStepWindowsExecStepConfigInterpreterEnum.to_proto(
            resource.interpreter
        ):
            res.interpreter = PatchDeploymentPatchConfigPreStepWindowsExecStepConfigInterpreterEnum.to_proto(
                resource.interpreter
            )
        return res

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return None

        return PatchDeploymentPatchConfigPreStepWindowsExecStepConfig(
            local_path=Primitive.from_proto(resource.local_path),
            gcs_object=PatchDeploymentPatchConfigPreStepWindowsExecStepConfigGcsObject.from_proto(
                resource.gcs_object
            ),
            allowed_success_codes=int64Array.from_proto(resource.allowed_success_codes),
            interpreter=PatchDeploymentPatchConfigPreStepWindowsExecStepConfigInterpreterEnum.from_proto(
                resource.interpreter
            ),
        )


class PatchDeploymentPatchConfigPreStepWindowsExecStepConfigArray(object):
    @classmethod
    def to_proto(self, resources):
        if not resources:
            return resources
        return [
            PatchDeploymentPatchConfigPreStepWindowsExecStepConfig.to_proto(i)
            for i in resources
        ]

    @classmethod
    def from_proto(self, resources):
        return [
            PatchDeploymentPatchConfigPreStepWindowsExecStepConfig.from_proto(i)
            for i in resources
        ]


class PatchDeploymentPatchConfigPreStepWindowsExecStepConfigGcsObject(object):
    def __init__(
        self, bucket: str = None, object: str = None, generation_number: int = None
    ):
        self.bucket = bucket
        self.object = object
        self.generation_number = generation_number

    @classmethod
    def to_proto(self, resource):
        if not resource:
            return None

        res = (
            patch_deployment_pb2.OsconfigBetaPatchDeploymentPatchConfigPreStepWindowsExecStepConfigGcsObject()
        )
        if Primitive.to_proto(resource.bucket):
            res.bucket = Primitive.to_proto(resource.bucket)
        if Primitive.to_proto(resource.object):
            res.object = Primitive.to_proto(resource.object)
        if Primitive.to_proto(resource.generation_number):
            res.generation_number = Primitive.to_proto(resource.generation_number)
        return res

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return None

        return PatchDeploymentPatchConfigPreStepWindowsExecStepConfigGcsObject(
            bucket=Primitive.from_proto(resource.bucket),
            object=Primitive.from_proto(resource.object),
            generation_number=Primitive.from_proto(resource.generation_number),
        )


class PatchDeploymentPatchConfigPreStepWindowsExecStepConfigGcsObjectArray(object):
    @classmethod
    def to_proto(self, resources):
        if not resources:
            return resources
        return [
            PatchDeploymentPatchConfigPreStepWindowsExecStepConfigGcsObject.to_proto(i)
            for i in resources
        ]

    @classmethod
    def from_proto(self, resources):
        return [
            PatchDeploymentPatchConfigPreStepWindowsExecStepConfigGcsObject.from_proto(
                i
            )
            for i in resources
        ]


class PatchDeploymentPatchConfigPostStep(object):
    def __init__(
        self, linux_exec_step_config: dict = None, windows_exec_step_config: dict = None
    ):
        self.linux_exec_step_config = linux_exec_step_config
        self.windows_exec_step_config = windows_exec_step_config

    @classmethod
    def to_proto(self, resource):
        if not resource:
            return None

        res = patch_deployment_pb2.OsconfigBetaPatchDeploymentPatchConfigPostStep()
        if PatchDeploymentPatchConfigPostStepLinuxExecStepConfig.to_proto(
            resource.linux_exec_step_config
        ):
            res.linux_exec_step_config.CopyFrom(
                PatchDeploymentPatchConfigPostStepLinuxExecStepConfig.to_proto(
                    resource.linux_exec_step_config
                )
            )
        else:
            res.ClearField("linux_exec_step_config")
        if PatchDeploymentPatchConfigPostStepWindowsExecStepConfig.to_proto(
            resource.windows_exec_step_config
        ):
            res.windows_exec_step_config.CopyFrom(
                PatchDeploymentPatchConfigPostStepWindowsExecStepConfig.to_proto(
                    resource.windows_exec_step_config
                )
            )
        else:
            res.ClearField("windows_exec_step_config")
        return res

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return None

        return PatchDeploymentPatchConfigPostStep(
            linux_exec_step_config=PatchDeploymentPatchConfigPostStepLinuxExecStepConfig.from_proto(
                resource.linux_exec_step_config
            ),
            windows_exec_step_config=PatchDeploymentPatchConfigPostStepWindowsExecStepConfig.from_proto(
                resource.windows_exec_step_config
            ),
        )


class PatchDeploymentPatchConfigPostStepArray(object):
    @classmethod
    def to_proto(self, resources):
        if not resources:
            return resources
        return [PatchDeploymentPatchConfigPostStep.to_proto(i) for i in resources]

    @classmethod
    def from_proto(self, resources):
        return [PatchDeploymentPatchConfigPostStep.from_proto(i) for i in resources]


class PatchDeploymentPatchConfigPostStepLinuxExecStepConfig(object):
    def __init__(
        self,
        local_path: str = None,
        gcs_object: dict = None,
        allowed_success_codes: list = None,
        interpreter: str = None,
    ):
        self.local_path = local_path
        self.gcs_object = gcs_object
        self.allowed_success_codes = allowed_success_codes
        self.interpreter = interpreter

    @classmethod
    def to_proto(self, resource):
        if not resource:
            return None

        res = (
            patch_deployment_pb2.OsconfigBetaPatchDeploymentPatchConfigPostStepLinuxExecStepConfig()
        )
        if Primitive.to_proto(resource.local_path):
            res.local_path = Primitive.to_proto(resource.local_path)
        if PatchDeploymentPatchConfigPostStepLinuxExecStepConfigGcsObject.to_proto(
            resource.gcs_object
        ):
            res.gcs_object.CopyFrom(
                PatchDeploymentPatchConfigPostStepLinuxExecStepConfigGcsObject.to_proto(
                    resource.gcs_object
                )
            )
        else:
            res.ClearField("gcs_object")
        if int64Array.to_proto(resource.allowed_success_codes):
            res.allowed_success_codes.extend(
                int64Array.to_proto(resource.allowed_success_codes)
            )
        if PatchDeploymentPatchConfigPostStepLinuxExecStepConfigInterpreterEnum.to_proto(
            resource.interpreter
        ):
            res.interpreter = PatchDeploymentPatchConfigPostStepLinuxExecStepConfigInterpreterEnum.to_proto(
                resource.interpreter
            )
        return res

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return None

        return PatchDeploymentPatchConfigPostStepLinuxExecStepConfig(
            local_path=Primitive.from_proto(resource.local_path),
            gcs_object=PatchDeploymentPatchConfigPostStepLinuxExecStepConfigGcsObject.from_proto(
                resource.gcs_object
            ),
            allowed_success_codes=int64Array.from_proto(resource.allowed_success_codes),
            interpreter=PatchDeploymentPatchConfigPostStepLinuxExecStepConfigInterpreterEnum.from_proto(
                resource.interpreter
            ),
        )


class PatchDeploymentPatchConfigPostStepLinuxExecStepConfigArray(object):
    @classmethod
    def to_proto(self, resources):
        if not resources:
            return resources
        return [
            PatchDeploymentPatchConfigPostStepLinuxExecStepConfig.to_proto(i)
            for i in resources
        ]

    @classmethod
    def from_proto(self, resources):
        return [
            PatchDeploymentPatchConfigPostStepLinuxExecStepConfig.from_proto(i)
            for i in resources
        ]


class PatchDeploymentPatchConfigPostStepLinuxExecStepConfigGcsObject(object):
    def __init__(
        self, bucket: str = None, object: str = None, generation_number: int = None
    ):
        self.bucket = bucket
        self.object = object
        self.generation_number = generation_number

    @classmethod
    def to_proto(self, resource):
        if not resource:
            return None

        res = (
            patch_deployment_pb2.OsconfigBetaPatchDeploymentPatchConfigPostStepLinuxExecStepConfigGcsObject()
        )
        if Primitive.to_proto(resource.bucket):
            res.bucket = Primitive.to_proto(resource.bucket)
        if Primitive.to_proto(resource.object):
            res.object = Primitive.to_proto(resource.object)
        if Primitive.to_proto(resource.generation_number):
            res.generation_number = Primitive.to_proto(resource.generation_number)
        return res

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return None

        return PatchDeploymentPatchConfigPostStepLinuxExecStepConfigGcsObject(
            bucket=Primitive.from_proto(resource.bucket),
            object=Primitive.from_proto(resource.object),
            generation_number=Primitive.from_proto(resource.generation_number),
        )


class PatchDeploymentPatchConfigPostStepLinuxExecStepConfigGcsObjectArray(object):
    @classmethod
    def to_proto(self, resources):
        if not resources:
            return resources
        return [
            PatchDeploymentPatchConfigPostStepLinuxExecStepConfigGcsObject.to_proto(i)
            for i in resources
        ]

    @classmethod
    def from_proto(self, resources):
        return [
            PatchDeploymentPatchConfigPostStepLinuxExecStepConfigGcsObject.from_proto(i)
            for i in resources
        ]


class PatchDeploymentPatchConfigPostStepWindowsExecStepConfig(object):
    def __init__(
        self,
        local_path: str = None,
        gcs_object: dict = None,
        allowed_success_codes: list = None,
        interpreter: str = None,
    ):
        self.local_path = local_path
        self.gcs_object = gcs_object
        self.allowed_success_codes = allowed_success_codes
        self.interpreter = interpreter

    @classmethod
    def to_proto(self, resource):
        if not resource:
            return None

        res = (
            patch_deployment_pb2.OsconfigBetaPatchDeploymentPatchConfigPostStepWindowsExecStepConfig()
        )
        if Primitive.to_proto(resource.local_path):
            res.local_path = Primitive.to_proto(resource.local_path)
        if PatchDeploymentPatchConfigPostStepWindowsExecStepConfigGcsObject.to_proto(
            resource.gcs_object
        ):
            res.gcs_object.CopyFrom(
                PatchDeploymentPatchConfigPostStepWindowsExecStepConfigGcsObject.to_proto(
                    resource.gcs_object
                )
            )
        else:
            res.ClearField("gcs_object")
        if int64Array.to_proto(resource.allowed_success_codes):
            res.allowed_success_codes.extend(
                int64Array.to_proto(resource.allowed_success_codes)
            )
        if PatchDeploymentPatchConfigPostStepWindowsExecStepConfigInterpreterEnum.to_proto(
            resource.interpreter
        ):
            res.interpreter = PatchDeploymentPatchConfigPostStepWindowsExecStepConfigInterpreterEnum.to_proto(
                resource.interpreter
            )
        return res

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return None

        return PatchDeploymentPatchConfigPostStepWindowsExecStepConfig(
            local_path=Primitive.from_proto(resource.local_path),
            gcs_object=PatchDeploymentPatchConfigPostStepWindowsExecStepConfigGcsObject.from_proto(
                resource.gcs_object
            ),
            allowed_success_codes=int64Array.from_proto(resource.allowed_success_codes),
            interpreter=PatchDeploymentPatchConfigPostStepWindowsExecStepConfigInterpreterEnum.from_proto(
                resource.interpreter
            ),
        )


class PatchDeploymentPatchConfigPostStepWindowsExecStepConfigArray(object):
    @classmethod
    def to_proto(self, resources):
        if not resources:
            return resources
        return [
            PatchDeploymentPatchConfigPostStepWindowsExecStepConfig.to_proto(i)
            for i in resources
        ]

    @classmethod
    def from_proto(self, resources):
        return [
            PatchDeploymentPatchConfigPostStepWindowsExecStepConfig.from_proto(i)
            for i in resources
        ]


class PatchDeploymentPatchConfigPostStepWindowsExecStepConfigGcsObject(object):
    def __init__(
        self, bucket: str = None, object: str = None, generation_number: int = None
    ):
        self.bucket = bucket
        self.object = object
        self.generation_number = generation_number

    @classmethod
    def to_proto(self, resource):
        if not resource:
            return None

        res = (
            patch_deployment_pb2.OsconfigBetaPatchDeploymentPatchConfigPostStepWindowsExecStepConfigGcsObject()
        )
        if Primitive.to_proto(resource.bucket):
            res.bucket = Primitive.to_proto(resource.bucket)
        if Primitive.to_proto(resource.object):
            res.object = Primitive.to_proto(resource.object)
        if Primitive.to_proto(resource.generation_number):
            res.generation_number = Primitive.to_proto(resource.generation_number)
        return res

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return None

        return PatchDeploymentPatchConfigPostStepWindowsExecStepConfigGcsObject(
            bucket=Primitive.from_proto(resource.bucket),
            object=Primitive.from_proto(resource.object),
            generation_number=Primitive.from_proto(resource.generation_number),
        )


class PatchDeploymentPatchConfigPostStepWindowsExecStepConfigGcsObjectArray(object):
    @classmethod
    def to_proto(self, resources):
        if not resources:
            return resources
        return [
            PatchDeploymentPatchConfigPostStepWindowsExecStepConfigGcsObject.to_proto(i)
            for i in resources
        ]

    @classmethod
    def from_proto(self, resources):
        return [
            PatchDeploymentPatchConfigPostStepWindowsExecStepConfigGcsObject.from_proto(
                i
            )
            for i in resources
        ]


class PatchDeploymentOneTimeSchedule(object):
    def __init__(self, execute_time: str = None):
        self.execute_time = execute_time

    @classmethod
    def to_proto(self, resource):
        if not resource:
            return None

        res = patch_deployment_pb2.OsconfigBetaPatchDeploymentOneTimeSchedule()
        if Primitive.to_proto(resource.execute_time):
            res.execute_time = Primitive.to_proto(resource.execute_time)
        return res

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return None

        return PatchDeploymentOneTimeSchedule(
            execute_time=Primitive.from_proto(resource.execute_time),
        )


class PatchDeploymentOneTimeScheduleArray(object):
    @classmethod
    def to_proto(self, resources):
        if not resources:
            return resources
        return [PatchDeploymentOneTimeSchedule.to_proto(i) for i in resources]

    @classmethod
    def from_proto(self, resources):
        return [PatchDeploymentOneTimeSchedule.from_proto(i) for i in resources]


class PatchDeploymentRecurringSchedule(object):
    def __init__(
        self,
        time_zone: dict = None,
        start_time: str = None,
        end_time: str = None,
        time_of_day: dict = None,
        frequency: str = None,
        weekly: dict = None,
        monthly: dict = None,
        last_execute_time: str = None,
        next_execute_time: str = None,
    ):
        self.time_zone = time_zone
        self.start_time = start_time
        self.end_time = end_time
        self.time_of_day = time_of_day
        self.frequency = frequency
        self.weekly = weekly
        self.monthly = monthly
        self.last_execute_time = last_execute_time
        self.next_execute_time = next_execute_time

    @classmethod
    def to_proto(self, resource):
        if not resource:
            return None

        res = patch_deployment_pb2.OsconfigBetaPatchDeploymentRecurringSchedule()
        if PatchDeploymentRecurringScheduleTimeZone.to_proto(resource.time_zone):
            res.time_zone.CopyFrom(
                PatchDeploymentRecurringScheduleTimeZone.to_proto(resource.time_zone)
            )
        else:
            res.ClearField("time_zone")
        if Primitive.to_proto(resource.start_time):
            res.start_time = Primitive.to_proto(resource.start_time)
        if Primitive.to_proto(resource.end_time):
            res.end_time = Primitive.to_proto(resource.end_time)
        if PatchDeploymentRecurringScheduleTimeOfDay.to_proto(resource.time_of_day):
            res.time_of_day.CopyFrom(
                PatchDeploymentRecurringScheduleTimeOfDay.to_proto(resource.time_of_day)
            )
        else:
            res.ClearField("time_of_day")
        if PatchDeploymentRecurringScheduleFrequencyEnum.to_proto(resource.frequency):
            res.frequency = PatchDeploymentRecurringScheduleFrequencyEnum.to_proto(
                resource.frequency
            )
        if PatchDeploymentRecurringScheduleWeekly.to_proto(resource.weekly):
            res.weekly.CopyFrom(
                PatchDeploymentRecurringScheduleWeekly.to_proto(resource.weekly)
            )
        else:
            res.ClearField("weekly")
        if PatchDeploymentRecurringScheduleMonthly.to_proto(resource.monthly):
            res.monthly.CopyFrom(
                PatchDeploymentRecurringScheduleMonthly.to_proto(resource.monthly)
            )
        else:
            res.ClearField("monthly")
        if Primitive.to_proto(resource.last_execute_time):
            res.last_execute_time = Primitive.to_proto(resource.last_execute_time)
        if Primitive.to_proto(resource.next_execute_time):
            res.next_execute_time = Primitive.to_proto(resource.next_execute_time)
        return res

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return None

        return PatchDeploymentRecurringSchedule(
            time_zone=PatchDeploymentRecurringScheduleTimeZone.from_proto(
                resource.time_zone
            ),
            start_time=Primitive.from_proto(resource.start_time),
            end_time=Primitive.from_proto(resource.end_time),
            time_of_day=PatchDeploymentRecurringScheduleTimeOfDay.from_proto(
                resource.time_of_day
            ),
            frequency=PatchDeploymentRecurringScheduleFrequencyEnum.from_proto(
                resource.frequency
            ),
            weekly=PatchDeploymentRecurringScheduleWeekly.from_proto(resource.weekly),
            monthly=PatchDeploymentRecurringScheduleMonthly.from_proto(
                resource.monthly
            ),
            last_execute_time=Primitive.from_proto(resource.last_execute_time),
            next_execute_time=Primitive.from_proto(resource.next_execute_time),
        )


class PatchDeploymentRecurringScheduleArray(object):
    @classmethod
    def to_proto(self, resources):
        if not resources:
            return resources
        return [PatchDeploymentRecurringSchedule.to_proto(i) for i in resources]

    @classmethod
    def from_proto(self, resources):
        return [PatchDeploymentRecurringSchedule.from_proto(i) for i in resources]


class PatchDeploymentRecurringScheduleTimeZone(object):
    def __init__(self, id: str = None, version: str = None):
        self.id = id
        self.version = version

    @classmethod
    def to_proto(self, resource):
        if not resource:
            return None

        res = (
            patch_deployment_pb2.OsconfigBetaPatchDeploymentRecurringScheduleTimeZone()
        )
        if Primitive.to_proto(resource.id):
            res.id = Primitive.to_proto(resource.id)
        if Primitive.to_proto(resource.version):
            res.version = Primitive.to_proto(resource.version)
        return res

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return None

        return PatchDeploymentRecurringScheduleTimeZone(
            id=Primitive.from_proto(resource.id),
            version=Primitive.from_proto(resource.version),
        )


class PatchDeploymentRecurringScheduleTimeZoneArray(object):
    @classmethod
    def to_proto(self, resources):
        if not resources:
            return resources
        return [PatchDeploymentRecurringScheduleTimeZone.to_proto(i) for i in resources]

    @classmethod
    def from_proto(self, resources):
        return [
            PatchDeploymentRecurringScheduleTimeZone.from_proto(i) for i in resources
        ]


class PatchDeploymentRecurringScheduleTimeOfDay(object):
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
            patch_deployment_pb2.OsconfigBetaPatchDeploymentRecurringScheduleTimeOfDay()
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

        return PatchDeploymentRecurringScheduleTimeOfDay(
            hours=Primitive.from_proto(resource.hours),
            minutes=Primitive.from_proto(resource.minutes),
            seconds=Primitive.from_proto(resource.seconds),
            nanos=Primitive.from_proto(resource.nanos),
        )


class PatchDeploymentRecurringScheduleTimeOfDayArray(object):
    @classmethod
    def to_proto(self, resources):
        if not resources:
            return resources
        return [
            PatchDeploymentRecurringScheduleTimeOfDay.to_proto(i) for i in resources
        ]

    @classmethod
    def from_proto(self, resources):
        return [
            PatchDeploymentRecurringScheduleTimeOfDay.from_proto(i) for i in resources
        ]


class PatchDeploymentRecurringScheduleWeekly(object):
    def __init__(self, day_of_week: str = None):
        self.day_of_week = day_of_week

    @classmethod
    def to_proto(self, resource):
        if not resource:
            return None

        res = patch_deployment_pb2.OsconfigBetaPatchDeploymentRecurringScheduleWeekly()
        if PatchDeploymentRecurringScheduleWeeklyDayOfWeekEnum.to_proto(
            resource.day_of_week
        ):
            res.day_of_week = (
                PatchDeploymentRecurringScheduleWeeklyDayOfWeekEnum.to_proto(
                    resource.day_of_week
                )
            )
        return res

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return None

        return PatchDeploymentRecurringScheduleWeekly(
            day_of_week=PatchDeploymentRecurringScheduleWeeklyDayOfWeekEnum.from_proto(
                resource.day_of_week
            ),
        )


class PatchDeploymentRecurringScheduleWeeklyArray(object):
    @classmethod
    def to_proto(self, resources):
        if not resources:
            return resources
        return [PatchDeploymentRecurringScheduleWeekly.to_proto(i) for i in resources]

    @classmethod
    def from_proto(self, resources):
        return [PatchDeploymentRecurringScheduleWeekly.from_proto(i) for i in resources]


class PatchDeploymentRecurringScheduleMonthly(object):
    def __init__(self, week_day_of_month: dict = None, month_day: int = None):
        self.week_day_of_month = week_day_of_month
        self.month_day = month_day

    @classmethod
    def to_proto(self, resource):
        if not resource:
            return None

        res = patch_deployment_pb2.OsconfigBetaPatchDeploymentRecurringScheduleMonthly()
        if PatchDeploymentRecurringScheduleMonthlyWeekDayOfMonth.to_proto(
            resource.week_day_of_month
        ):
            res.week_day_of_month.CopyFrom(
                PatchDeploymentRecurringScheduleMonthlyWeekDayOfMonth.to_proto(
                    resource.week_day_of_month
                )
            )
        else:
            res.ClearField("week_day_of_month")
        if Primitive.to_proto(resource.month_day):
            res.month_day = Primitive.to_proto(resource.month_day)
        return res

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return None

        return PatchDeploymentRecurringScheduleMonthly(
            week_day_of_month=PatchDeploymentRecurringScheduleMonthlyWeekDayOfMonth.from_proto(
                resource.week_day_of_month
            ),
            month_day=Primitive.from_proto(resource.month_day),
        )


class PatchDeploymentRecurringScheduleMonthlyArray(object):
    @classmethod
    def to_proto(self, resources):
        if not resources:
            return resources
        return [PatchDeploymentRecurringScheduleMonthly.to_proto(i) for i in resources]

    @classmethod
    def from_proto(self, resources):
        return [
            PatchDeploymentRecurringScheduleMonthly.from_proto(i) for i in resources
        ]


class PatchDeploymentRecurringScheduleMonthlyWeekDayOfMonth(object):
    def __init__(self, week_ordinal: int = None, day_of_week: str = None):
        self.week_ordinal = week_ordinal
        self.day_of_week = day_of_week

    @classmethod
    def to_proto(self, resource):
        if not resource:
            return None

        res = (
            patch_deployment_pb2.OsconfigBetaPatchDeploymentRecurringScheduleMonthlyWeekDayOfMonth()
        )
        if Primitive.to_proto(resource.week_ordinal):
            res.week_ordinal = Primitive.to_proto(resource.week_ordinal)
        if PatchDeploymentRecurringScheduleMonthlyWeekDayOfMonthDayOfWeekEnum.to_proto(
            resource.day_of_week
        ):
            res.day_of_week = PatchDeploymentRecurringScheduleMonthlyWeekDayOfMonthDayOfWeekEnum.to_proto(
                resource.day_of_week
            )
        return res

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return None

        return PatchDeploymentRecurringScheduleMonthlyWeekDayOfMonth(
            week_ordinal=Primitive.from_proto(resource.week_ordinal),
            day_of_week=PatchDeploymentRecurringScheduleMonthlyWeekDayOfMonthDayOfWeekEnum.from_proto(
                resource.day_of_week
            ),
        )


class PatchDeploymentRecurringScheduleMonthlyWeekDayOfMonthArray(object):
    @classmethod
    def to_proto(self, resources):
        if not resources:
            return resources
        return [
            PatchDeploymentRecurringScheduleMonthlyWeekDayOfMonth.to_proto(i)
            for i in resources
        ]

    @classmethod
    def from_proto(self, resources):
        return [
            PatchDeploymentRecurringScheduleMonthlyWeekDayOfMonth.from_proto(i)
            for i in resources
        ]


class PatchDeploymentRollout(object):
    def __init__(self, mode: str = None, disruption_budget: dict = None):
        self.mode = mode
        self.disruption_budget = disruption_budget

    @classmethod
    def to_proto(self, resource):
        if not resource:
            return None

        res = patch_deployment_pb2.OsconfigBetaPatchDeploymentRollout()
        if PatchDeploymentRolloutModeEnum.to_proto(resource.mode):
            res.mode = PatchDeploymentRolloutModeEnum.to_proto(resource.mode)
        if PatchDeploymentRolloutDisruptionBudget.to_proto(resource.disruption_budget):
            res.disruption_budget.CopyFrom(
                PatchDeploymentRolloutDisruptionBudget.to_proto(
                    resource.disruption_budget
                )
            )
        else:
            res.ClearField("disruption_budget")
        return res

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return None

        return PatchDeploymentRollout(
            mode=PatchDeploymentRolloutModeEnum.from_proto(resource.mode),
            disruption_budget=PatchDeploymentRolloutDisruptionBudget.from_proto(
                resource.disruption_budget
            ),
        )


class PatchDeploymentRolloutArray(object):
    @classmethod
    def to_proto(self, resources):
        if not resources:
            return resources
        return [PatchDeploymentRollout.to_proto(i) for i in resources]

    @classmethod
    def from_proto(self, resources):
        return [PatchDeploymentRollout.from_proto(i) for i in resources]


class PatchDeploymentRolloutDisruptionBudget(object):
    def __init__(self, fixed: int = None, percent: int = None):
        self.fixed = fixed
        self.percent = percent

    @classmethod
    def to_proto(self, resource):
        if not resource:
            return None

        res = patch_deployment_pb2.OsconfigBetaPatchDeploymentRolloutDisruptionBudget()
        if Primitive.to_proto(resource.fixed):
            res.fixed = Primitive.to_proto(resource.fixed)
        if Primitive.to_proto(resource.percent):
            res.percent = Primitive.to_proto(resource.percent)
        return res

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return None

        return PatchDeploymentRolloutDisruptionBudget(
            fixed=Primitive.from_proto(resource.fixed),
            percent=Primitive.from_proto(resource.percent),
        )


class PatchDeploymentRolloutDisruptionBudgetArray(object):
    @classmethod
    def to_proto(self, resources):
        if not resources:
            return resources
        return [PatchDeploymentRolloutDisruptionBudget.to_proto(i) for i in resources]

    @classmethod
    def from_proto(self, resources):
        return [PatchDeploymentRolloutDisruptionBudget.from_proto(i) for i in resources]


class PatchDeploymentPatchConfigRebootConfigEnum(object):
    @classmethod
    def to_proto(self, resource):
        if not resource:
            return resource
        return patch_deployment_pb2.OsconfigBetaPatchDeploymentPatchConfigRebootConfigEnum.Value(
            "OsconfigBetaPatchDeploymentPatchConfigRebootConfigEnum%s" % resource
        )

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return resource
        return patch_deployment_pb2.OsconfigBetaPatchDeploymentPatchConfigRebootConfigEnum.Name(
            resource
        )[
            len("OsconfigBetaPatchDeploymentPatchConfigRebootConfigEnum") :
        ]


class PatchDeploymentPatchConfigAptTypeEnum(object):
    @classmethod
    def to_proto(self, resource):
        if not resource:
            return resource
        return patch_deployment_pb2.OsconfigBetaPatchDeploymentPatchConfigAptTypeEnum.Value(
            "OsconfigBetaPatchDeploymentPatchConfigAptTypeEnum%s" % resource
        )

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return resource
        return (
            patch_deployment_pb2.OsconfigBetaPatchDeploymentPatchConfigAptTypeEnum.Name(
                resource
            )[len("OsconfigBetaPatchDeploymentPatchConfigAptTypeEnum") :]
        )


class PatchDeploymentPatchConfigWindowsUpdateClassificationsEnum(object):
    @classmethod
    def to_proto(self, resource):
        if not resource:
            return resource
        return patch_deployment_pb2.OsconfigBetaPatchDeploymentPatchConfigWindowsUpdateClassificationsEnum.Value(
            "OsconfigBetaPatchDeploymentPatchConfigWindowsUpdateClassificationsEnum%s"
            % resource
        )

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return resource
        return patch_deployment_pb2.OsconfigBetaPatchDeploymentPatchConfigWindowsUpdateClassificationsEnum.Name(
            resource
        )[
            len(
                "OsconfigBetaPatchDeploymentPatchConfigWindowsUpdateClassificationsEnum"
            ) :
        ]


class PatchDeploymentPatchConfigPreStepLinuxExecStepConfigInterpreterEnum(object):
    @classmethod
    def to_proto(self, resource):
        if not resource:
            return resource
        return patch_deployment_pb2.OsconfigBetaPatchDeploymentPatchConfigPreStepLinuxExecStepConfigInterpreterEnum.Value(
            "OsconfigBetaPatchDeploymentPatchConfigPreStepLinuxExecStepConfigInterpreterEnum%s"
            % resource
        )

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return resource
        return patch_deployment_pb2.OsconfigBetaPatchDeploymentPatchConfigPreStepLinuxExecStepConfigInterpreterEnum.Name(
            resource
        )[
            len(
                "OsconfigBetaPatchDeploymentPatchConfigPreStepLinuxExecStepConfigInterpreterEnum"
            ) :
        ]


class PatchDeploymentPatchConfigPreStepWindowsExecStepConfigInterpreterEnum(object):
    @classmethod
    def to_proto(self, resource):
        if not resource:
            return resource
        return patch_deployment_pb2.OsconfigBetaPatchDeploymentPatchConfigPreStepWindowsExecStepConfigInterpreterEnum.Value(
            "OsconfigBetaPatchDeploymentPatchConfigPreStepWindowsExecStepConfigInterpreterEnum%s"
            % resource
        )

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return resource
        return patch_deployment_pb2.OsconfigBetaPatchDeploymentPatchConfigPreStepWindowsExecStepConfigInterpreterEnum.Name(
            resource
        )[
            len(
                "OsconfigBetaPatchDeploymentPatchConfigPreStepWindowsExecStepConfigInterpreterEnum"
            ) :
        ]


class PatchDeploymentPatchConfigPostStepLinuxExecStepConfigInterpreterEnum(object):
    @classmethod
    def to_proto(self, resource):
        if not resource:
            return resource
        return patch_deployment_pb2.OsconfigBetaPatchDeploymentPatchConfigPostStepLinuxExecStepConfigInterpreterEnum.Value(
            "OsconfigBetaPatchDeploymentPatchConfigPostStepLinuxExecStepConfigInterpreterEnum%s"
            % resource
        )

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return resource
        return patch_deployment_pb2.OsconfigBetaPatchDeploymentPatchConfigPostStepLinuxExecStepConfigInterpreterEnum.Name(
            resource
        )[
            len(
                "OsconfigBetaPatchDeploymentPatchConfigPostStepLinuxExecStepConfigInterpreterEnum"
            ) :
        ]


class PatchDeploymentPatchConfigPostStepWindowsExecStepConfigInterpreterEnum(object):
    @classmethod
    def to_proto(self, resource):
        if not resource:
            return resource
        return patch_deployment_pb2.OsconfigBetaPatchDeploymentPatchConfigPostStepWindowsExecStepConfigInterpreterEnum.Value(
            "OsconfigBetaPatchDeploymentPatchConfigPostStepWindowsExecStepConfigInterpreterEnum%s"
            % resource
        )

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return resource
        return patch_deployment_pb2.OsconfigBetaPatchDeploymentPatchConfigPostStepWindowsExecStepConfigInterpreterEnum.Name(
            resource
        )[
            len(
                "OsconfigBetaPatchDeploymentPatchConfigPostStepWindowsExecStepConfigInterpreterEnum"
            ) :
        ]


class PatchDeploymentRecurringScheduleFrequencyEnum(object):
    @classmethod
    def to_proto(self, resource):
        if not resource:
            return resource
        return patch_deployment_pb2.OsconfigBetaPatchDeploymentRecurringScheduleFrequencyEnum.Value(
            "OsconfigBetaPatchDeploymentRecurringScheduleFrequencyEnum%s" % resource
        )

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return resource
        return patch_deployment_pb2.OsconfigBetaPatchDeploymentRecurringScheduleFrequencyEnum.Name(
            resource
        )[
            len("OsconfigBetaPatchDeploymentRecurringScheduleFrequencyEnum") :
        ]


class PatchDeploymentRecurringScheduleWeeklyDayOfWeekEnum(object):
    @classmethod
    def to_proto(self, resource):
        if not resource:
            return resource
        return patch_deployment_pb2.OsconfigBetaPatchDeploymentRecurringScheduleWeeklyDayOfWeekEnum.Value(
            "OsconfigBetaPatchDeploymentRecurringScheduleWeeklyDayOfWeekEnum%s"
            % resource
        )

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return resource
        return patch_deployment_pb2.OsconfigBetaPatchDeploymentRecurringScheduleWeeklyDayOfWeekEnum.Name(
            resource
        )[
            len("OsconfigBetaPatchDeploymentRecurringScheduleWeeklyDayOfWeekEnum") :
        ]


class PatchDeploymentRecurringScheduleMonthlyWeekDayOfMonthDayOfWeekEnum(object):
    @classmethod
    def to_proto(self, resource):
        if not resource:
            return resource
        return patch_deployment_pb2.OsconfigBetaPatchDeploymentRecurringScheduleMonthlyWeekDayOfMonthDayOfWeekEnum.Value(
            "OsconfigBetaPatchDeploymentRecurringScheduleMonthlyWeekDayOfMonthDayOfWeekEnum%s"
            % resource
        )

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return resource
        return patch_deployment_pb2.OsconfigBetaPatchDeploymentRecurringScheduleMonthlyWeekDayOfMonthDayOfWeekEnum.Name(
            resource
        )[
            len(
                "OsconfigBetaPatchDeploymentRecurringScheduleMonthlyWeekDayOfMonthDayOfWeekEnum"
            ) :
        ]


class PatchDeploymentRolloutModeEnum(object):
    @classmethod
    def to_proto(self, resource):
        if not resource:
            return resource
        return patch_deployment_pb2.OsconfigBetaPatchDeploymentRolloutModeEnum.Value(
            "OsconfigBetaPatchDeploymentRolloutModeEnum%s" % resource
        )

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return resource
        return patch_deployment_pb2.OsconfigBetaPatchDeploymentRolloutModeEnum.Name(
            resource
        )[len("OsconfigBetaPatchDeploymentRolloutModeEnum") :]


class Primitive(object):
    @classmethod
    def to_proto(self, s):
        if not s:
            return ""
        return s

    @classmethod
    def from_proto(self, s):
        return s
