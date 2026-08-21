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
from google3.cloud.graphite.mmv2.services.google.os_config import guest_policy_pb2
from google3.cloud.graphite.mmv2.services.google.os_config import guest_policy_pb2_grpc

from typing import List


class GuestPolicy(object):
    def __init__(
        self,
        name: str = None,
        description: str = None,
        create_time: str = None,
        update_time: str = None,
        assignment: dict = None,
        packages: list = None,
        package_repositories: list = None,
        recipes: list = None,
        etag: str = None,
        project: str = None,
        service_account_file: str = "",
    ):
        channel.initialize()
        self.name = name
        self.description = description
        self.assignment = assignment
        self.packages = packages
        self.package_repositories = package_repositories
        self.recipes = recipes
        self.project = project
        self.service_account_file = service_account_file

    def apply(self):
        stub = guest_policy_pb2_grpc.OsconfigBetaGuestPolicyServiceStub(
            channel.Channel()
        )
        request = guest_policy_pb2.ApplyOsconfigBetaGuestPolicyRequest()
        if Primitive.to_proto(self.name):
            request.resource.name = Primitive.to_proto(self.name)

        if Primitive.to_proto(self.description):
            request.resource.description = Primitive.to_proto(self.description)

        if GuestPolicyAssignment.to_proto(self.assignment):
            request.resource.assignment.CopyFrom(
                GuestPolicyAssignment.to_proto(self.assignment)
            )
        else:
            request.resource.ClearField("assignment")
        if GuestPolicyPackagesArray.to_proto(self.packages):
            request.resource.packages.extend(
                GuestPolicyPackagesArray.to_proto(self.packages)
            )
        if GuestPolicyPackageRepositoriesArray.to_proto(self.package_repositories):
            request.resource.package_repositories.extend(
                GuestPolicyPackageRepositoriesArray.to_proto(self.package_repositories)
            )
        if GuestPolicyRecipesArray.to_proto(self.recipes):
            request.resource.recipes.extend(
                GuestPolicyRecipesArray.to_proto(self.recipes)
            )
        if Primitive.to_proto(self.project):
            request.resource.project = Primitive.to_proto(self.project)

        request.service_account_file = self.service_account_file

        response = stub.ApplyOsconfigBetaGuestPolicy(request)
        self.name = Primitive.from_proto(response.name)
        self.description = Primitive.from_proto(response.description)
        self.create_time = Primitive.from_proto(response.create_time)
        self.update_time = Primitive.from_proto(response.update_time)
        self.assignment = GuestPolicyAssignment.from_proto(response.assignment)
        self.packages = GuestPolicyPackagesArray.from_proto(response.packages)
        self.package_repositories = GuestPolicyPackageRepositoriesArray.from_proto(
            response.package_repositories
        )
        self.recipes = GuestPolicyRecipesArray.from_proto(response.recipes)
        self.etag = Primitive.from_proto(response.etag)
        self.project = Primitive.from_proto(response.project)

    def delete(self):
        stub = guest_policy_pb2_grpc.OsconfigBetaGuestPolicyServiceStub(
            channel.Channel()
        )
        request = guest_policy_pb2.DeleteOsconfigBetaGuestPolicyRequest()
        request.service_account_file = self.service_account_file
        if Primitive.to_proto(self.name):
            request.resource.name = Primitive.to_proto(self.name)

        if Primitive.to_proto(self.description):
            request.resource.description = Primitive.to_proto(self.description)

        if GuestPolicyAssignment.to_proto(self.assignment):
            request.resource.assignment.CopyFrom(
                GuestPolicyAssignment.to_proto(self.assignment)
            )
        else:
            request.resource.ClearField("assignment")
        if GuestPolicyPackagesArray.to_proto(self.packages):
            request.resource.packages.extend(
                GuestPolicyPackagesArray.to_proto(self.packages)
            )
        if GuestPolicyPackageRepositoriesArray.to_proto(self.package_repositories):
            request.resource.package_repositories.extend(
                GuestPolicyPackageRepositoriesArray.to_proto(self.package_repositories)
            )
        if GuestPolicyRecipesArray.to_proto(self.recipes):
            request.resource.recipes.extend(
                GuestPolicyRecipesArray.to_proto(self.recipes)
            )
        if Primitive.to_proto(self.project):
            request.resource.project = Primitive.to_proto(self.project)

        response = stub.DeleteOsconfigBetaGuestPolicy(request)

    @classmethod
    def list(self, project, service_account_file=""):
        stub = guest_policy_pb2_grpc.OsconfigBetaGuestPolicyServiceStub(
            channel.Channel()
        )
        request = guest_policy_pb2.ListOsconfigBetaGuestPolicyRequest()
        request.service_account_file = service_account_file
        request.Project = project

        return stub.ListOsconfigBetaGuestPolicy(request).items

    def to_proto(self):
        resource = guest_policy_pb2.OsconfigBetaGuestPolicy()
        if Primitive.to_proto(self.name):
            resource.name = Primitive.to_proto(self.name)
        if Primitive.to_proto(self.description):
            resource.description = Primitive.to_proto(self.description)
        if GuestPolicyAssignment.to_proto(self.assignment):
            resource.assignment.CopyFrom(
                GuestPolicyAssignment.to_proto(self.assignment)
            )
        else:
            resource.ClearField("assignment")
        if GuestPolicyPackagesArray.to_proto(self.packages):
            resource.packages.extend(GuestPolicyPackagesArray.to_proto(self.packages))
        if GuestPolicyPackageRepositoriesArray.to_proto(self.package_repositories):
            resource.package_repositories.extend(
                GuestPolicyPackageRepositoriesArray.to_proto(self.package_repositories)
            )
        if GuestPolicyRecipesArray.to_proto(self.recipes):
            resource.recipes.extend(GuestPolicyRecipesArray.to_proto(self.recipes))
        if Primitive.to_proto(self.project):
            resource.project = Primitive.to_proto(self.project)
        return resource


class GuestPolicyAssignment(object):
    def __init__(
        self,
        group_labels: list = None,
        zones: list = None,
        instances: list = None,
        instance_name_prefixes: list = None,
        os_types: list = None,
    ):
        self.group_labels = group_labels
        self.zones = zones
        self.instances = instances
        self.instance_name_prefixes = instance_name_prefixes
        self.os_types = os_types

    @classmethod
    def to_proto(self, resource):
        if not resource:
            return None

        res = guest_policy_pb2.OsconfigBetaGuestPolicyAssignment()
        if GuestPolicyAssignmentGroupLabelsArray.to_proto(resource.group_labels):
            res.group_labels.extend(
                GuestPolicyAssignmentGroupLabelsArray.to_proto(resource.group_labels)
            )
        if Primitive.to_proto(resource.zones):
            res.zones.extend(Primitive.to_proto(resource.zones))
        if Primitive.to_proto(resource.instances):
            res.instances.extend(Primitive.to_proto(resource.instances))
        if Primitive.to_proto(resource.instance_name_prefixes):
            res.instance_name_prefixes.extend(
                Primitive.to_proto(resource.instance_name_prefixes)
            )
        if GuestPolicyAssignmentOSTypesArray.to_proto(resource.os_types):
            res.os_types.extend(
                GuestPolicyAssignmentOSTypesArray.to_proto(resource.os_types)
            )
        return res

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return None

        return GuestPolicyAssignment(
            group_labels=GuestPolicyAssignmentGroupLabelsArray.from_proto(
                resource.group_labels
            ),
            zones=Primitive.from_proto(resource.zones),
            instances=Primitive.from_proto(resource.instances),
            instance_name_prefixes=Primitive.from_proto(
                resource.instance_name_prefixes
            ),
            os_types=GuestPolicyAssignmentOSTypesArray.from_proto(resource.os_types),
        )


class GuestPolicyAssignmentArray(object):
    @classmethod
    def to_proto(self, resources):
        if not resources:
            return resources
        return [GuestPolicyAssignment.to_proto(i) for i in resources]

    @classmethod
    def from_proto(self, resources):
        return [GuestPolicyAssignment.from_proto(i) for i in resources]


class GuestPolicyAssignmentGroupLabels(object):
    def __init__(self, labels: dict = None):
        self.labels = labels

    @classmethod
    def to_proto(self, resource):
        if not resource:
            return None

        res = guest_policy_pb2.OsconfigBetaGuestPolicyAssignmentGroupLabels()
        if Primitive.to_proto(resource.labels):
            res.labels = Primitive.to_proto(resource.labels)
        return res

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return None

        return GuestPolicyAssignmentGroupLabels(
            labels=Primitive.from_proto(resource.labels),
        )


class GuestPolicyAssignmentGroupLabelsArray(object):
    @classmethod
    def to_proto(self, resources):
        if not resources:
            return resources
        return [GuestPolicyAssignmentGroupLabels.to_proto(i) for i in resources]

    @classmethod
    def from_proto(self, resources):
        return [GuestPolicyAssignmentGroupLabels.from_proto(i) for i in resources]


class GuestPolicyAssignmentOSTypes(object):
    def __init__(
        self,
        os_short_name: str = None,
        os_version: str = None,
        os_architecture: str = None,
    ):
        self.os_short_name = os_short_name
        self.os_version = os_version
        self.os_architecture = os_architecture

    @classmethod
    def to_proto(self, resource):
        if not resource:
            return None

        res = guest_policy_pb2.OsconfigBetaGuestPolicyAssignmentOSTypes()
        if Primitive.to_proto(resource.os_short_name):
            res.os_short_name = Primitive.to_proto(resource.os_short_name)
        if Primitive.to_proto(resource.os_version):
            res.os_version = Primitive.to_proto(resource.os_version)
        if Primitive.to_proto(resource.os_architecture):
            res.os_architecture = Primitive.to_proto(resource.os_architecture)
        return res

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return None

        return GuestPolicyAssignmentOSTypes(
            os_short_name=Primitive.from_proto(resource.os_short_name),
            os_version=Primitive.from_proto(resource.os_version),
            os_architecture=Primitive.from_proto(resource.os_architecture),
        )


class GuestPolicyAssignmentOSTypesArray(object):
    @classmethod
    def to_proto(self, resources):
        if not resources:
            return resources
        return [GuestPolicyAssignmentOSTypes.to_proto(i) for i in resources]

    @classmethod
    def from_proto(self, resources):
        return [GuestPolicyAssignmentOSTypes.from_proto(i) for i in resources]


class GuestPolicyPackages(object):
    def __init__(
        self, name: str = None, desired_state: str = None, manager: str = None
    ):
        self.name = name
        self.desired_state = desired_state
        self.manager = manager

    @classmethod
    def to_proto(self, resource):
        if not resource:
            return None

        res = guest_policy_pb2.OsconfigBetaGuestPolicyPackages()
        if Primitive.to_proto(resource.name):
            res.name = Primitive.to_proto(resource.name)
        if GuestPolicyPackagesDesiredStateEnum.to_proto(resource.desired_state):
            res.desired_state = GuestPolicyPackagesDesiredStateEnum.to_proto(
                resource.desired_state
            )
        if GuestPolicyPackagesManagerEnum.to_proto(resource.manager):
            res.manager = GuestPolicyPackagesManagerEnum.to_proto(resource.manager)
        return res

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return None

        return GuestPolicyPackages(
            name=Primitive.from_proto(resource.name),
            desired_state=GuestPolicyPackagesDesiredStateEnum.from_proto(
                resource.desired_state
            ),
            manager=GuestPolicyPackagesManagerEnum.from_proto(resource.manager),
        )


class GuestPolicyPackagesArray(object):
    @classmethod
    def to_proto(self, resources):
        if not resources:
            return resources
        return [GuestPolicyPackages.to_proto(i) for i in resources]

    @classmethod
    def from_proto(self, resources):
        return [GuestPolicyPackages.from_proto(i) for i in resources]


class GuestPolicyPackageRepositories(object):
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

        res = guest_policy_pb2.OsconfigBetaGuestPolicyPackageRepositories()
        if GuestPolicyPackageRepositoriesApt.to_proto(resource.apt):
            res.apt.CopyFrom(GuestPolicyPackageRepositoriesApt.to_proto(resource.apt))
        else:
            res.ClearField("apt")
        if GuestPolicyPackageRepositoriesYum.to_proto(resource.yum):
            res.yum.CopyFrom(GuestPolicyPackageRepositoriesYum.to_proto(resource.yum))
        else:
            res.ClearField("yum")
        if GuestPolicyPackageRepositoriesZypper.to_proto(resource.zypper):
            res.zypper.CopyFrom(
                GuestPolicyPackageRepositoriesZypper.to_proto(resource.zypper)
            )
        else:
            res.ClearField("zypper")
        if GuestPolicyPackageRepositoriesGoo.to_proto(resource.goo):
            res.goo.CopyFrom(GuestPolicyPackageRepositoriesGoo.to_proto(resource.goo))
        else:
            res.ClearField("goo")
        return res

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return None

        return GuestPolicyPackageRepositories(
            apt=GuestPolicyPackageRepositoriesApt.from_proto(resource.apt),
            yum=GuestPolicyPackageRepositoriesYum.from_proto(resource.yum),
            zypper=GuestPolicyPackageRepositoriesZypper.from_proto(resource.zypper),
            goo=GuestPolicyPackageRepositoriesGoo.from_proto(resource.goo),
        )


class GuestPolicyPackageRepositoriesArray(object):
    @classmethod
    def to_proto(self, resources):
        if not resources:
            return resources
        return [GuestPolicyPackageRepositories.to_proto(i) for i in resources]

    @classmethod
    def from_proto(self, resources):
        return [GuestPolicyPackageRepositories.from_proto(i) for i in resources]


class GuestPolicyPackageRepositoriesApt(object):
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

        res = guest_policy_pb2.OsconfigBetaGuestPolicyPackageRepositoriesApt()
        if GuestPolicyPackageRepositoriesAptArchiveTypeEnum.to_proto(
            resource.archive_type
        ):
            res.archive_type = (
                GuestPolicyPackageRepositoriesAptArchiveTypeEnum.to_proto(
                    resource.archive_type
                )
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

        return GuestPolicyPackageRepositoriesApt(
            archive_type=GuestPolicyPackageRepositoriesAptArchiveTypeEnum.from_proto(
                resource.archive_type
            ),
            uri=Primitive.from_proto(resource.uri),
            distribution=Primitive.from_proto(resource.distribution),
            components=Primitive.from_proto(resource.components),
            gpg_key=Primitive.from_proto(resource.gpg_key),
        )


class GuestPolicyPackageRepositoriesAptArray(object):
    @classmethod
    def to_proto(self, resources):
        if not resources:
            return resources
        return [GuestPolicyPackageRepositoriesApt.to_proto(i) for i in resources]

    @classmethod
    def from_proto(self, resources):
        return [GuestPolicyPackageRepositoriesApt.from_proto(i) for i in resources]


class GuestPolicyPackageRepositoriesYum(object):
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

        res = guest_policy_pb2.OsconfigBetaGuestPolicyPackageRepositoriesYum()
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

        return GuestPolicyPackageRepositoriesYum(
            id=Primitive.from_proto(resource.id),
            display_name=Primitive.from_proto(resource.display_name),
            base_url=Primitive.from_proto(resource.base_url),
            gpg_keys=Primitive.from_proto(resource.gpg_keys),
        )


class GuestPolicyPackageRepositoriesYumArray(object):
    @classmethod
    def to_proto(self, resources):
        if not resources:
            return resources
        return [GuestPolicyPackageRepositoriesYum.to_proto(i) for i in resources]

    @classmethod
    def from_proto(self, resources):
        return [GuestPolicyPackageRepositoriesYum.from_proto(i) for i in resources]


class GuestPolicyPackageRepositoriesZypper(object):
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

        res = guest_policy_pb2.OsconfigBetaGuestPolicyPackageRepositoriesZypper()
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

        return GuestPolicyPackageRepositoriesZypper(
            id=Primitive.from_proto(resource.id),
            display_name=Primitive.from_proto(resource.display_name),
            base_url=Primitive.from_proto(resource.base_url),
            gpg_keys=Primitive.from_proto(resource.gpg_keys),
        )


class GuestPolicyPackageRepositoriesZypperArray(object):
    @classmethod
    def to_proto(self, resources):
        if not resources:
            return resources
        return [GuestPolicyPackageRepositoriesZypper.to_proto(i) for i in resources]

    @classmethod
    def from_proto(self, resources):
        return [GuestPolicyPackageRepositoriesZypper.from_proto(i) for i in resources]


class GuestPolicyPackageRepositoriesGoo(object):
    def __init__(self, name: str = None, url: str = None):
        self.name = name
        self.url = url

    @classmethod
    def to_proto(self, resource):
        if not resource:
            return None

        res = guest_policy_pb2.OsconfigBetaGuestPolicyPackageRepositoriesGoo()
        if Primitive.to_proto(resource.name):
            res.name = Primitive.to_proto(resource.name)
        if Primitive.to_proto(resource.url):
            res.url = Primitive.to_proto(resource.url)
        return res

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return None

        return GuestPolicyPackageRepositoriesGoo(
            name=Primitive.from_proto(resource.name),
            url=Primitive.from_proto(resource.url),
        )


class GuestPolicyPackageRepositoriesGooArray(object):
    @classmethod
    def to_proto(self, resources):
        if not resources:
            return resources
        return [GuestPolicyPackageRepositoriesGoo.to_proto(i) for i in resources]

    @classmethod
    def from_proto(self, resources):
        return [GuestPolicyPackageRepositoriesGoo.from_proto(i) for i in resources]


class GuestPolicyRecipes(object):
    def __init__(
        self,
        name: str = None,
        version: str = None,
        artifacts: list = None,
        install_steps: list = None,
        update_steps: list = None,
        desired_state: str = None,
    ):
        self.name = name
        self.version = version
        self.artifacts = artifacts
        self.install_steps = install_steps
        self.update_steps = update_steps
        self.desired_state = desired_state

    @classmethod
    def to_proto(self, resource):
        if not resource:
            return None

        res = guest_policy_pb2.OsconfigBetaGuestPolicyRecipes()
        if Primitive.to_proto(resource.name):
            res.name = Primitive.to_proto(resource.name)
        if Primitive.to_proto(resource.version):
            res.version = Primitive.to_proto(resource.version)
        if GuestPolicyRecipesArtifactsArray.to_proto(resource.artifacts):
            res.artifacts.extend(
                GuestPolicyRecipesArtifactsArray.to_proto(resource.artifacts)
            )
        if GuestPolicyRecipesInstallStepsArray.to_proto(resource.install_steps):
            res.install_steps.extend(
                GuestPolicyRecipesInstallStepsArray.to_proto(resource.install_steps)
            )
        if GuestPolicyRecipesUpdateStepsArray.to_proto(resource.update_steps):
            res.update_steps.extend(
                GuestPolicyRecipesUpdateStepsArray.to_proto(resource.update_steps)
            )
        if GuestPolicyRecipesDesiredStateEnum.to_proto(resource.desired_state):
            res.desired_state = GuestPolicyRecipesDesiredStateEnum.to_proto(
                resource.desired_state
            )
        return res

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return None

        return GuestPolicyRecipes(
            name=Primitive.from_proto(resource.name),
            version=Primitive.from_proto(resource.version),
            artifacts=GuestPolicyRecipesArtifactsArray.from_proto(resource.artifacts),
            install_steps=GuestPolicyRecipesInstallStepsArray.from_proto(
                resource.install_steps
            ),
            update_steps=GuestPolicyRecipesUpdateStepsArray.from_proto(
                resource.update_steps
            ),
            desired_state=GuestPolicyRecipesDesiredStateEnum.from_proto(
                resource.desired_state
            ),
        )


class GuestPolicyRecipesArray(object):
    @classmethod
    def to_proto(self, resources):
        if not resources:
            return resources
        return [GuestPolicyRecipes.to_proto(i) for i in resources]

    @classmethod
    def from_proto(self, resources):
        return [GuestPolicyRecipes.from_proto(i) for i in resources]


class GuestPolicyRecipesArtifacts(object):
    def __init__(
        self,
        id: str = None,
        remote: dict = None,
        gcs: dict = None,
        allow_insecure: bool = None,
    ):
        self.id = id
        self.remote = remote
        self.gcs = gcs
        self.allow_insecure = allow_insecure

    @classmethod
    def to_proto(self, resource):
        if not resource:
            return None

        res = guest_policy_pb2.OsconfigBetaGuestPolicyRecipesArtifacts()
        if Primitive.to_proto(resource.id):
            res.id = Primitive.to_proto(resource.id)
        if GuestPolicyRecipesArtifactsRemote.to_proto(resource.remote):
            res.remote.CopyFrom(
                GuestPolicyRecipesArtifactsRemote.to_proto(resource.remote)
            )
        else:
            res.ClearField("remote")
        if GuestPolicyRecipesArtifactsGcs.to_proto(resource.gcs):
            res.gcs.CopyFrom(GuestPolicyRecipesArtifactsGcs.to_proto(resource.gcs))
        else:
            res.ClearField("gcs")
        if Primitive.to_proto(resource.allow_insecure):
            res.allow_insecure = Primitive.to_proto(resource.allow_insecure)
        return res

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return None

        return GuestPolicyRecipesArtifacts(
            id=Primitive.from_proto(resource.id),
            remote=GuestPolicyRecipesArtifactsRemote.from_proto(resource.remote),
            gcs=GuestPolicyRecipesArtifactsGcs.from_proto(resource.gcs),
            allow_insecure=Primitive.from_proto(resource.allow_insecure),
        )


class GuestPolicyRecipesArtifactsArray(object):
    @classmethod
    def to_proto(self, resources):
        if not resources:
            return resources
        return [GuestPolicyRecipesArtifacts.to_proto(i) for i in resources]

    @classmethod
    def from_proto(self, resources):
        return [GuestPolicyRecipesArtifacts.from_proto(i) for i in resources]


class GuestPolicyRecipesArtifactsRemote(object):
    def __init__(self, uri: str = None, checksum: str = None):
        self.uri = uri
        self.checksum = checksum

    @classmethod
    def to_proto(self, resource):
        if not resource:
            return None

        res = guest_policy_pb2.OsconfigBetaGuestPolicyRecipesArtifactsRemote()
        if Primitive.to_proto(resource.uri):
            res.uri = Primitive.to_proto(resource.uri)
        if Primitive.to_proto(resource.checksum):
            res.checksum = Primitive.to_proto(resource.checksum)
        return res

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return None

        return GuestPolicyRecipesArtifactsRemote(
            uri=Primitive.from_proto(resource.uri),
            checksum=Primitive.from_proto(resource.checksum),
        )


class GuestPolicyRecipesArtifactsRemoteArray(object):
    @classmethod
    def to_proto(self, resources):
        if not resources:
            return resources
        return [GuestPolicyRecipesArtifactsRemote.to_proto(i) for i in resources]

    @classmethod
    def from_proto(self, resources):
        return [GuestPolicyRecipesArtifactsRemote.from_proto(i) for i in resources]


class GuestPolicyRecipesArtifactsGcs(object):
    def __init__(self, bucket: str = None, object: str = None, generation: int = None):
        self.bucket = bucket
        self.object = object
        self.generation = generation

    @classmethod
    def to_proto(self, resource):
        if not resource:
            return None

        res = guest_policy_pb2.OsconfigBetaGuestPolicyRecipesArtifactsGcs()
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

        return GuestPolicyRecipesArtifactsGcs(
            bucket=Primitive.from_proto(resource.bucket),
            object=Primitive.from_proto(resource.object),
            generation=Primitive.from_proto(resource.generation),
        )


class GuestPolicyRecipesArtifactsGcsArray(object):
    @classmethod
    def to_proto(self, resources):
        if not resources:
            return resources
        return [GuestPolicyRecipesArtifactsGcs.to_proto(i) for i in resources]

    @classmethod
    def from_proto(self, resources):
        return [GuestPolicyRecipesArtifactsGcs.from_proto(i) for i in resources]


class GuestPolicyRecipesInstallSteps(object):
    def __init__(
        self,
        file_copy: dict = None,
        archive_extraction: dict = None,
        msi_installation: dict = None,
        dpkg_installation: dict = None,
        rpm_installation: dict = None,
        file_exec: dict = None,
        script_run: dict = None,
    ):
        self.file_copy = file_copy
        self.archive_extraction = archive_extraction
        self.msi_installation = msi_installation
        self.dpkg_installation = dpkg_installation
        self.rpm_installation = rpm_installation
        self.file_exec = file_exec
        self.script_run = script_run

    @classmethod
    def to_proto(self, resource):
        if not resource:
            return None

        res = guest_policy_pb2.OsconfigBetaGuestPolicyRecipesInstallSteps()
        if GuestPolicyRecipesInstallStepsFileCopy.to_proto(resource.file_copy):
            res.file_copy.CopyFrom(
                GuestPolicyRecipesInstallStepsFileCopy.to_proto(resource.file_copy)
            )
        else:
            res.ClearField("file_copy")
        if GuestPolicyRecipesInstallStepsArchiveExtraction.to_proto(
            resource.archive_extraction
        ):
            res.archive_extraction.CopyFrom(
                GuestPolicyRecipesInstallStepsArchiveExtraction.to_proto(
                    resource.archive_extraction
                )
            )
        else:
            res.ClearField("archive_extraction")
        if GuestPolicyRecipesInstallStepsMsiInstallation.to_proto(
            resource.msi_installation
        ):
            res.msi_installation.CopyFrom(
                GuestPolicyRecipesInstallStepsMsiInstallation.to_proto(
                    resource.msi_installation
                )
            )
        else:
            res.ClearField("msi_installation")
        if GuestPolicyRecipesInstallStepsDpkgInstallation.to_proto(
            resource.dpkg_installation
        ):
            res.dpkg_installation.CopyFrom(
                GuestPolicyRecipesInstallStepsDpkgInstallation.to_proto(
                    resource.dpkg_installation
                )
            )
        else:
            res.ClearField("dpkg_installation")
        if GuestPolicyRecipesInstallStepsRpmInstallation.to_proto(
            resource.rpm_installation
        ):
            res.rpm_installation.CopyFrom(
                GuestPolicyRecipesInstallStepsRpmInstallation.to_proto(
                    resource.rpm_installation
                )
            )
        else:
            res.ClearField("rpm_installation")
        if GuestPolicyRecipesInstallStepsFileExec.to_proto(resource.file_exec):
            res.file_exec.CopyFrom(
                GuestPolicyRecipesInstallStepsFileExec.to_proto(resource.file_exec)
            )
        else:
            res.ClearField("file_exec")
        if GuestPolicyRecipesInstallStepsScriptRun.to_proto(resource.script_run):
            res.script_run.CopyFrom(
                GuestPolicyRecipesInstallStepsScriptRun.to_proto(resource.script_run)
            )
        else:
            res.ClearField("script_run")
        return res

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return None

        return GuestPolicyRecipesInstallSteps(
            file_copy=GuestPolicyRecipesInstallStepsFileCopy.from_proto(
                resource.file_copy
            ),
            archive_extraction=GuestPolicyRecipesInstallStepsArchiveExtraction.from_proto(
                resource.archive_extraction
            ),
            msi_installation=GuestPolicyRecipesInstallStepsMsiInstallation.from_proto(
                resource.msi_installation
            ),
            dpkg_installation=GuestPolicyRecipesInstallStepsDpkgInstallation.from_proto(
                resource.dpkg_installation
            ),
            rpm_installation=GuestPolicyRecipesInstallStepsRpmInstallation.from_proto(
                resource.rpm_installation
            ),
            file_exec=GuestPolicyRecipesInstallStepsFileExec.from_proto(
                resource.file_exec
            ),
            script_run=GuestPolicyRecipesInstallStepsScriptRun.from_proto(
                resource.script_run
            ),
        )


class GuestPolicyRecipesInstallStepsArray(object):
    @classmethod
    def to_proto(self, resources):
        if not resources:
            return resources
        return [GuestPolicyRecipesInstallSteps.to_proto(i) for i in resources]

    @classmethod
    def from_proto(self, resources):
        return [GuestPolicyRecipesInstallSteps.from_proto(i) for i in resources]


class GuestPolicyRecipesInstallStepsFileCopy(object):
    def __init__(
        self,
        artifact_id: str = None,
        destination: str = None,
        overwrite: bool = None,
        permissions: str = None,
    ):
        self.artifact_id = artifact_id
        self.destination = destination
        self.overwrite = overwrite
        self.permissions = permissions

    @classmethod
    def to_proto(self, resource):
        if not resource:
            return None

        res = guest_policy_pb2.OsconfigBetaGuestPolicyRecipesInstallStepsFileCopy()
        if Primitive.to_proto(resource.artifact_id):
            res.artifact_id = Primitive.to_proto(resource.artifact_id)
        if Primitive.to_proto(resource.destination):
            res.destination = Primitive.to_proto(resource.destination)
        if Primitive.to_proto(resource.overwrite):
            res.overwrite = Primitive.to_proto(resource.overwrite)
        if Primitive.to_proto(resource.permissions):
            res.permissions = Primitive.to_proto(resource.permissions)
        return res

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return None

        return GuestPolicyRecipesInstallStepsFileCopy(
            artifact_id=Primitive.from_proto(resource.artifact_id),
            destination=Primitive.from_proto(resource.destination),
            overwrite=Primitive.from_proto(resource.overwrite),
            permissions=Primitive.from_proto(resource.permissions),
        )


class GuestPolicyRecipesInstallStepsFileCopyArray(object):
    @classmethod
    def to_proto(self, resources):
        if not resources:
            return resources
        return [GuestPolicyRecipesInstallStepsFileCopy.to_proto(i) for i in resources]

    @classmethod
    def from_proto(self, resources):
        return [GuestPolicyRecipesInstallStepsFileCopy.from_proto(i) for i in resources]


class GuestPolicyRecipesInstallStepsArchiveExtraction(object):
    def __init__(
        self, artifact_id: str = None, destination: str = None, type: str = None
    ):
        self.artifact_id = artifact_id
        self.destination = destination
        self.type = type

    @classmethod
    def to_proto(self, resource):
        if not resource:
            return None

        res = (
            guest_policy_pb2.OsconfigBetaGuestPolicyRecipesInstallStepsArchiveExtraction()
        )
        if Primitive.to_proto(resource.artifact_id):
            res.artifact_id = Primitive.to_proto(resource.artifact_id)
        if Primitive.to_proto(resource.destination):
            res.destination = Primitive.to_proto(resource.destination)
        if GuestPolicyRecipesInstallStepsArchiveExtractionTypeEnum.to_proto(
            resource.type
        ):
            res.type = GuestPolicyRecipesInstallStepsArchiveExtractionTypeEnum.to_proto(
                resource.type
            )
        return res

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return None

        return GuestPolicyRecipesInstallStepsArchiveExtraction(
            artifact_id=Primitive.from_proto(resource.artifact_id),
            destination=Primitive.from_proto(resource.destination),
            type=GuestPolicyRecipesInstallStepsArchiveExtractionTypeEnum.from_proto(
                resource.type
            ),
        )


class GuestPolicyRecipesInstallStepsArchiveExtractionArray(object):
    @classmethod
    def to_proto(self, resources):
        if not resources:
            return resources
        return [
            GuestPolicyRecipesInstallStepsArchiveExtraction.to_proto(i)
            for i in resources
        ]

    @classmethod
    def from_proto(self, resources):
        return [
            GuestPolicyRecipesInstallStepsArchiveExtraction.from_proto(i)
            for i in resources
        ]


class GuestPolicyRecipesInstallStepsMsiInstallation(object):
    def __init__(
        self,
        artifact_id: str = None,
        flags: list = None,
        allowed_exit_codes: list = None,
    ):
        self.artifact_id = artifact_id
        self.flags = flags
        self.allowed_exit_codes = allowed_exit_codes

    @classmethod
    def to_proto(self, resource):
        if not resource:
            return None

        res = (
            guest_policy_pb2.OsconfigBetaGuestPolicyRecipesInstallStepsMsiInstallation()
        )
        if Primitive.to_proto(resource.artifact_id):
            res.artifact_id = Primitive.to_proto(resource.artifact_id)
        if Primitive.to_proto(resource.flags):
            res.flags.extend(Primitive.to_proto(resource.flags))
        if int64Array.to_proto(resource.allowed_exit_codes):
            res.allowed_exit_codes.extend(
                int64Array.to_proto(resource.allowed_exit_codes)
            )
        return res

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return None

        return GuestPolicyRecipesInstallStepsMsiInstallation(
            artifact_id=Primitive.from_proto(resource.artifact_id),
            flags=Primitive.from_proto(resource.flags),
            allowed_exit_codes=int64Array.from_proto(resource.allowed_exit_codes),
        )


class GuestPolicyRecipesInstallStepsMsiInstallationArray(object):
    @classmethod
    def to_proto(self, resources):
        if not resources:
            return resources
        return [
            GuestPolicyRecipesInstallStepsMsiInstallation.to_proto(i) for i in resources
        ]

    @classmethod
    def from_proto(self, resources):
        return [
            GuestPolicyRecipesInstallStepsMsiInstallation.from_proto(i)
            for i in resources
        ]


class GuestPolicyRecipesInstallStepsDpkgInstallation(object):
    def __init__(self, artifact_id: str = None):
        self.artifact_id = artifact_id

    @classmethod
    def to_proto(self, resource):
        if not resource:
            return None

        res = (
            guest_policy_pb2.OsconfigBetaGuestPolicyRecipesInstallStepsDpkgInstallation()
        )
        if Primitive.to_proto(resource.artifact_id):
            res.artifact_id = Primitive.to_proto(resource.artifact_id)
        return res

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return None

        return GuestPolicyRecipesInstallStepsDpkgInstallation(
            artifact_id=Primitive.from_proto(resource.artifact_id),
        )


class GuestPolicyRecipesInstallStepsDpkgInstallationArray(object):
    @classmethod
    def to_proto(self, resources):
        if not resources:
            return resources
        return [
            GuestPolicyRecipesInstallStepsDpkgInstallation.to_proto(i)
            for i in resources
        ]

    @classmethod
    def from_proto(self, resources):
        return [
            GuestPolicyRecipesInstallStepsDpkgInstallation.from_proto(i)
            for i in resources
        ]


class GuestPolicyRecipesInstallStepsRpmInstallation(object):
    def __init__(self, artifact_id: str = None):
        self.artifact_id = artifact_id

    @classmethod
    def to_proto(self, resource):
        if not resource:
            return None

        res = (
            guest_policy_pb2.OsconfigBetaGuestPolicyRecipesInstallStepsRpmInstallation()
        )
        if Primitive.to_proto(resource.artifact_id):
            res.artifact_id = Primitive.to_proto(resource.artifact_id)
        return res

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return None

        return GuestPolicyRecipesInstallStepsRpmInstallation(
            artifact_id=Primitive.from_proto(resource.artifact_id),
        )


class GuestPolicyRecipesInstallStepsRpmInstallationArray(object):
    @classmethod
    def to_proto(self, resources):
        if not resources:
            return resources
        return [
            GuestPolicyRecipesInstallStepsRpmInstallation.to_proto(i) for i in resources
        ]

    @classmethod
    def from_proto(self, resources):
        return [
            GuestPolicyRecipesInstallStepsRpmInstallation.from_proto(i)
            for i in resources
        ]


class GuestPolicyRecipesInstallStepsFileExec(object):
    def __init__(
        self,
        artifact_id: str = None,
        local_path: str = None,
        args: list = None,
        allowed_exit_codes: list = None,
    ):
        self.artifact_id = artifact_id
        self.local_path = local_path
        self.args = args
        self.allowed_exit_codes = allowed_exit_codes

    @classmethod
    def to_proto(self, resource):
        if not resource:
            return None

        res = guest_policy_pb2.OsconfigBetaGuestPolicyRecipesInstallStepsFileExec()
        if Primitive.to_proto(resource.artifact_id):
            res.artifact_id = Primitive.to_proto(resource.artifact_id)
        if Primitive.to_proto(resource.local_path):
            res.local_path = Primitive.to_proto(resource.local_path)
        if Primitive.to_proto(resource.args):
            res.args.extend(Primitive.to_proto(resource.args))
        if int64Array.to_proto(resource.allowed_exit_codes):
            res.allowed_exit_codes.extend(
                int64Array.to_proto(resource.allowed_exit_codes)
            )
        return res

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return None

        return GuestPolicyRecipesInstallStepsFileExec(
            artifact_id=Primitive.from_proto(resource.artifact_id),
            local_path=Primitive.from_proto(resource.local_path),
            args=Primitive.from_proto(resource.args),
            allowed_exit_codes=int64Array.from_proto(resource.allowed_exit_codes),
        )


class GuestPolicyRecipesInstallStepsFileExecArray(object):
    @classmethod
    def to_proto(self, resources):
        if not resources:
            return resources
        return [GuestPolicyRecipesInstallStepsFileExec.to_proto(i) for i in resources]

    @classmethod
    def from_proto(self, resources):
        return [GuestPolicyRecipesInstallStepsFileExec.from_proto(i) for i in resources]


class GuestPolicyRecipesInstallStepsScriptRun(object):
    def __init__(
        self,
        script: str = None,
        allowed_exit_codes: list = None,
        interpreter: str = None,
    ):
        self.script = script
        self.allowed_exit_codes = allowed_exit_codes
        self.interpreter = interpreter

    @classmethod
    def to_proto(self, resource):
        if not resource:
            return None

        res = guest_policy_pb2.OsconfigBetaGuestPolicyRecipesInstallStepsScriptRun()
        if Primitive.to_proto(resource.script):
            res.script = Primitive.to_proto(resource.script)
        if int64Array.to_proto(resource.allowed_exit_codes):
            res.allowed_exit_codes.extend(
                int64Array.to_proto(resource.allowed_exit_codes)
            )
        if GuestPolicyRecipesInstallStepsScriptRunInterpreterEnum.to_proto(
            resource.interpreter
        ):
            res.interpreter = (
                GuestPolicyRecipesInstallStepsScriptRunInterpreterEnum.to_proto(
                    resource.interpreter
                )
            )
        return res

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return None

        return GuestPolicyRecipesInstallStepsScriptRun(
            script=Primitive.from_proto(resource.script),
            allowed_exit_codes=int64Array.from_proto(resource.allowed_exit_codes),
            interpreter=GuestPolicyRecipesInstallStepsScriptRunInterpreterEnum.from_proto(
                resource.interpreter
            ),
        )


class GuestPolicyRecipesInstallStepsScriptRunArray(object):
    @classmethod
    def to_proto(self, resources):
        if not resources:
            return resources
        return [GuestPolicyRecipesInstallStepsScriptRun.to_proto(i) for i in resources]

    @classmethod
    def from_proto(self, resources):
        return [
            GuestPolicyRecipesInstallStepsScriptRun.from_proto(i) for i in resources
        ]


class GuestPolicyRecipesUpdateSteps(object):
    def __init__(
        self,
        file_copy: dict = None,
        archive_extraction: dict = None,
        msi_installation: dict = None,
        dpkg_installation: dict = None,
        rpm_installation: dict = None,
        file_exec: dict = None,
        script_run: dict = None,
    ):
        self.file_copy = file_copy
        self.archive_extraction = archive_extraction
        self.msi_installation = msi_installation
        self.dpkg_installation = dpkg_installation
        self.rpm_installation = rpm_installation
        self.file_exec = file_exec
        self.script_run = script_run

    @classmethod
    def to_proto(self, resource):
        if not resource:
            return None

        res = guest_policy_pb2.OsconfigBetaGuestPolicyRecipesUpdateSteps()
        if GuestPolicyRecipesUpdateStepsFileCopy.to_proto(resource.file_copy):
            res.file_copy.CopyFrom(
                GuestPolicyRecipesUpdateStepsFileCopy.to_proto(resource.file_copy)
            )
        else:
            res.ClearField("file_copy")
        if GuestPolicyRecipesUpdateStepsArchiveExtraction.to_proto(
            resource.archive_extraction
        ):
            res.archive_extraction.CopyFrom(
                GuestPolicyRecipesUpdateStepsArchiveExtraction.to_proto(
                    resource.archive_extraction
                )
            )
        else:
            res.ClearField("archive_extraction")
        if GuestPolicyRecipesUpdateStepsMsiInstallation.to_proto(
            resource.msi_installation
        ):
            res.msi_installation.CopyFrom(
                GuestPolicyRecipesUpdateStepsMsiInstallation.to_proto(
                    resource.msi_installation
                )
            )
        else:
            res.ClearField("msi_installation")
        if GuestPolicyRecipesUpdateStepsDpkgInstallation.to_proto(
            resource.dpkg_installation
        ):
            res.dpkg_installation.CopyFrom(
                GuestPolicyRecipesUpdateStepsDpkgInstallation.to_proto(
                    resource.dpkg_installation
                )
            )
        else:
            res.ClearField("dpkg_installation")
        if GuestPolicyRecipesUpdateStepsRpmInstallation.to_proto(
            resource.rpm_installation
        ):
            res.rpm_installation.CopyFrom(
                GuestPolicyRecipesUpdateStepsRpmInstallation.to_proto(
                    resource.rpm_installation
                )
            )
        else:
            res.ClearField("rpm_installation")
        if GuestPolicyRecipesUpdateStepsFileExec.to_proto(resource.file_exec):
            res.file_exec.CopyFrom(
                GuestPolicyRecipesUpdateStepsFileExec.to_proto(resource.file_exec)
            )
        else:
            res.ClearField("file_exec")
        if GuestPolicyRecipesUpdateStepsScriptRun.to_proto(resource.script_run):
            res.script_run.CopyFrom(
                GuestPolicyRecipesUpdateStepsScriptRun.to_proto(resource.script_run)
            )
        else:
            res.ClearField("script_run")
        return res

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return None

        return GuestPolicyRecipesUpdateSteps(
            file_copy=GuestPolicyRecipesUpdateStepsFileCopy.from_proto(
                resource.file_copy
            ),
            archive_extraction=GuestPolicyRecipesUpdateStepsArchiveExtraction.from_proto(
                resource.archive_extraction
            ),
            msi_installation=GuestPolicyRecipesUpdateStepsMsiInstallation.from_proto(
                resource.msi_installation
            ),
            dpkg_installation=GuestPolicyRecipesUpdateStepsDpkgInstallation.from_proto(
                resource.dpkg_installation
            ),
            rpm_installation=GuestPolicyRecipesUpdateStepsRpmInstallation.from_proto(
                resource.rpm_installation
            ),
            file_exec=GuestPolicyRecipesUpdateStepsFileExec.from_proto(
                resource.file_exec
            ),
            script_run=GuestPolicyRecipesUpdateStepsScriptRun.from_proto(
                resource.script_run
            ),
        )


class GuestPolicyRecipesUpdateStepsArray(object):
    @classmethod
    def to_proto(self, resources):
        if not resources:
            return resources
        return [GuestPolicyRecipesUpdateSteps.to_proto(i) for i in resources]

    @classmethod
    def from_proto(self, resources):
        return [GuestPolicyRecipesUpdateSteps.from_proto(i) for i in resources]


class GuestPolicyRecipesUpdateStepsFileCopy(object):
    def __init__(
        self,
        artifact_id: str = None,
        destination: str = None,
        overwrite: bool = None,
        permissions: str = None,
    ):
        self.artifact_id = artifact_id
        self.destination = destination
        self.overwrite = overwrite
        self.permissions = permissions

    @classmethod
    def to_proto(self, resource):
        if not resource:
            return None

        res = guest_policy_pb2.OsconfigBetaGuestPolicyRecipesUpdateStepsFileCopy()
        if Primitive.to_proto(resource.artifact_id):
            res.artifact_id = Primitive.to_proto(resource.artifact_id)
        if Primitive.to_proto(resource.destination):
            res.destination = Primitive.to_proto(resource.destination)
        if Primitive.to_proto(resource.overwrite):
            res.overwrite = Primitive.to_proto(resource.overwrite)
        if Primitive.to_proto(resource.permissions):
            res.permissions = Primitive.to_proto(resource.permissions)
        return res

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return None

        return GuestPolicyRecipesUpdateStepsFileCopy(
            artifact_id=Primitive.from_proto(resource.artifact_id),
            destination=Primitive.from_proto(resource.destination),
            overwrite=Primitive.from_proto(resource.overwrite),
            permissions=Primitive.from_proto(resource.permissions),
        )


class GuestPolicyRecipesUpdateStepsFileCopyArray(object):
    @classmethod
    def to_proto(self, resources):
        if not resources:
            return resources
        return [GuestPolicyRecipesUpdateStepsFileCopy.to_proto(i) for i in resources]

    @classmethod
    def from_proto(self, resources):
        return [GuestPolicyRecipesUpdateStepsFileCopy.from_proto(i) for i in resources]


class GuestPolicyRecipesUpdateStepsArchiveExtraction(object):
    def __init__(
        self, artifact_id: str = None, destination: str = None, type: str = None
    ):
        self.artifact_id = artifact_id
        self.destination = destination
        self.type = type

    @classmethod
    def to_proto(self, resource):
        if not resource:
            return None

        res = (
            guest_policy_pb2.OsconfigBetaGuestPolicyRecipesUpdateStepsArchiveExtraction()
        )
        if Primitive.to_proto(resource.artifact_id):
            res.artifact_id = Primitive.to_proto(resource.artifact_id)
        if Primitive.to_proto(resource.destination):
            res.destination = Primitive.to_proto(resource.destination)
        if GuestPolicyRecipesUpdateStepsArchiveExtractionTypeEnum.to_proto(
            resource.type
        ):
            res.type = GuestPolicyRecipesUpdateStepsArchiveExtractionTypeEnum.to_proto(
                resource.type
            )
        return res

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return None

        return GuestPolicyRecipesUpdateStepsArchiveExtraction(
            artifact_id=Primitive.from_proto(resource.artifact_id),
            destination=Primitive.from_proto(resource.destination),
            type=GuestPolicyRecipesUpdateStepsArchiveExtractionTypeEnum.from_proto(
                resource.type
            ),
        )


class GuestPolicyRecipesUpdateStepsArchiveExtractionArray(object):
    @classmethod
    def to_proto(self, resources):
        if not resources:
            return resources
        return [
            GuestPolicyRecipesUpdateStepsArchiveExtraction.to_proto(i)
            for i in resources
        ]

    @classmethod
    def from_proto(self, resources):
        return [
            GuestPolicyRecipesUpdateStepsArchiveExtraction.from_proto(i)
            for i in resources
        ]


class GuestPolicyRecipesUpdateStepsMsiInstallation(object):
    def __init__(
        self,
        artifact_id: str = None,
        flags: list = None,
        allowed_exit_codes: list = None,
    ):
        self.artifact_id = artifact_id
        self.flags = flags
        self.allowed_exit_codes = allowed_exit_codes

    @classmethod
    def to_proto(self, resource):
        if not resource:
            return None

        res = (
            guest_policy_pb2.OsconfigBetaGuestPolicyRecipesUpdateStepsMsiInstallation()
        )
        if Primitive.to_proto(resource.artifact_id):
            res.artifact_id = Primitive.to_proto(resource.artifact_id)
        if Primitive.to_proto(resource.flags):
            res.flags.extend(Primitive.to_proto(resource.flags))
        if int64Array.to_proto(resource.allowed_exit_codes):
            res.allowed_exit_codes.extend(
                int64Array.to_proto(resource.allowed_exit_codes)
            )
        return res

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return None

        return GuestPolicyRecipesUpdateStepsMsiInstallation(
            artifact_id=Primitive.from_proto(resource.artifact_id),
            flags=Primitive.from_proto(resource.flags),
            allowed_exit_codes=int64Array.from_proto(resource.allowed_exit_codes),
        )


class GuestPolicyRecipesUpdateStepsMsiInstallationArray(object):
    @classmethod
    def to_proto(self, resources):
        if not resources:
            return resources
        return [
            GuestPolicyRecipesUpdateStepsMsiInstallation.to_proto(i) for i in resources
        ]

    @classmethod
    def from_proto(self, resources):
        return [
            GuestPolicyRecipesUpdateStepsMsiInstallation.from_proto(i)
            for i in resources
        ]


class GuestPolicyRecipesUpdateStepsDpkgInstallation(object):
    def __init__(self, artifact_id: str = None):
        self.artifact_id = artifact_id

    @classmethod
    def to_proto(self, resource):
        if not resource:
            return None

        res = (
            guest_policy_pb2.OsconfigBetaGuestPolicyRecipesUpdateStepsDpkgInstallation()
        )
        if Primitive.to_proto(resource.artifact_id):
            res.artifact_id = Primitive.to_proto(resource.artifact_id)
        return res

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return None

        return GuestPolicyRecipesUpdateStepsDpkgInstallation(
            artifact_id=Primitive.from_proto(resource.artifact_id),
        )


class GuestPolicyRecipesUpdateStepsDpkgInstallationArray(object):
    @classmethod
    def to_proto(self, resources):
        if not resources:
            return resources
        return [
            GuestPolicyRecipesUpdateStepsDpkgInstallation.to_proto(i) for i in resources
        ]

    @classmethod
    def from_proto(self, resources):
        return [
            GuestPolicyRecipesUpdateStepsDpkgInstallation.from_proto(i)
            for i in resources
        ]


class GuestPolicyRecipesUpdateStepsRpmInstallation(object):
    def __init__(self, artifact_id: str = None):
        self.artifact_id = artifact_id

    @classmethod
    def to_proto(self, resource):
        if not resource:
            return None

        res = (
            guest_policy_pb2.OsconfigBetaGuestPolicyRecipesUpdateStepsRpmInstallation()
        )
        if Primitive.to_proto(resource.artifact_id):
            res.artifact_id = Primitive.to_proto(resource.artifact_id)
        return res

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return None

        return GuestPolicyRecipesUpdateStepsRpmInstallation(
            artifact_id=Primitive.from_proto(resource.artifact_id),
        )


class GuestPolicyRecipesUpdateStepsRpmInstallationArray(object):
    @classmethod
    def to_proto(self, resources):
        if not resources:
            return resources
        return [
            GuestPolicyRecipesUpdateStepsRpmInstallation.to_proto(i) for i in resources
        ]

    @classmethod
    def from_proto(self, resources):
        return [
            GuestPolicyRecipesUpdateStepsRpmInstallation.from_proto(i)
            for i in resources
        ]


class GuestPolicyRecipesUpdateStepsFileExec(object):
    def __init__(
        self,
        artifact_id: str = None,
        local_path: str = None,
        args: list = None,
        allowed_exit_codes: list = None,
    ):
        self.artifact_id = artifact_id
        self.local_path = local_path
        self.args = args
        self.allowed_exit_codes = allowed_exit_codes

    @classmethod
    def to_proto(self, resource):
        if not resource:
            return None

        res = guest_policy_pb2.OsconfigBetaGuestPolicyRecipesUpdateStepsFileExec()
        if Primitive.to_proto(resource.artifact_id):
            res.artifact_id = Primitive.to_proto(resource.artifact_id)
        if Primitive.to_proto(resource.local_path):
            res.local_path = Primitive.to_proto(resource.local_path)
        if Primitive.to_proto(resource.args):
            res.args.extend(Primitive.to_proto(resource.args))
        if int64Array.to_proto(resource.allowed_exit_codes):
            res.allowed_exit_codes.extend(
                int64Array.to_proto(resource.allowed_exit_codes)
            )
        return res

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return None

        return GuestPolicyRecipesUpdateStepsFileExec(
            artifact_id=Primitive.from_proto(resource.artifact_id),
            local_path=Primitive.from_proto(resource.local_path),
            args=Primitive.from_proto(resource.args),
            allowed_exit_codes=int64Array.from_proto(resource.allowed_exit_codes),
        )


class GuestPolicyRecipesUpdateStepsFileExecArray(object):
    @classmethod
    def to_proto(self, resources):
        if not resources:
            return resources
        return [GuestPolicyRecipesUpdateStepsFileExec.to_proto(i) for i in resources]

    @classmethod
    def from_proto(self, resources):
        return [GuestPolicyRecipesUpdateStepsFileExec.from_proto(i) for i in resources]


class GuestPolicyRecipesUpdateStepsScriptRun(object):
    def __init__(
        self,
        script: str = None,
        allowed_exit_codes: list = None,
        interpreter: str = None,
    ):
        self.script = script
        self.allowed_exit_codes = allowed_exit_codes
        self.interpreter = interpreter

    @classmethod
    def to_proto(self, resource):
        if not resource:
            return None

        res = guest_policy_pb2.OsconfigBetaGuestPolicyRecipesUpdateStepsScriptRun()
        if Primitive.to_proto(resource.script):
            res.script = Primitive.to_proto(resource.script)
        if int64Array.to_proto(resource.allowed_exit_codes):
            res.allowed_exit_codes.extend(
                int64Array.to_proto(resource.allowed_exit_codes)
            )
        if GuestPolicyRecipesUpdateStepsScriptRunInterpreterEnum.to_proto(
            resource.interpreter
        ):
            res.interpreter = (
                GuestPolicyRecipesUpdateStepsScriptRunInterpreterEnum.to_proto(
                    resource.interpreter
                )
            )
        return res

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return None

        return GuestPolicyRecipesUpdateStepsScriptRun(
            script=Primitive.from_proto(resource.script),
            allowed_exit_codes=int64Array.from_proto(resource.allowed_exit_codes),
            interpreter=GuestPolicyRecipesUpdateStepsScriptRunInterpreterEnum.from_proto(
                resource.interpreter
            ),
        )


class GuestPolicyRecipesUpdateStepsScriptRunArray(object):
    @classmethod
    def to_proto(self, resources):
        if not resources:
            return resources
        return [GuestPolicyRecipesUpdateStepsScriptRun.to_proto(i) for i in resources]

    @classmethod
    def from_proto(self, resources):
        return [GuestPolicyRecipesUpdateStepsScriptRun.from_proto(i) for i in resources]


class GuestPolicyPackagesDesiredStateEnum(object):
    @classmethod
    def to_proto(self, resource):
        if not resource:
            return resource
        return guest_policy_pb2.OsconfigBetaGuestPolicyPackagesDesiredStateEnum.Value(
            "OsconfigBetaGuestPolicyPackagesDesiredStateEnum%s" % resource
        )

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return resource
        return guest_policy_pb2.OsconfigBetaGuestPolicyPackagesDesiredStateEnum.Name(
            resource
        )[len("OsconfigBetaGuestPolicyPackagesDesiredStateEnum") :]


class GuestPolicyPackagesManagerEnum(object):
    @classmethod
    def to_proto(self, resource):
        if not resource:
            return resource
        return guest_policy_pb2.OsconfigBetaGuestPolicyPackagesManagerEnum.Value(
            "OsconfigBetaGuestPolicyPackagesManagerEnum%s" % resource
        )

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return resource
        return guest_policy_pb2.OsconfigBetaGuestPolicyPackagesManagerEnum.Name(
            resource
        )[len("OsconfigBetaGuestPolicyPackagesManagerEnum") :]


class GuestPolicyPackageRepositoriesAptArchiveTypeEnum(object):
    @classmethod
    def to_proto(self, resource):
        if not resource:
            return resource
        return guest_policy_pb2.OsconfigBetaGuestPolicyPackageRepositoriesAptArchiveTypeEnum.Value(
            "OsconfigBetaGuestPolicyPackageRepositoriesAptArchiveTypeEnum%s" % resource
        )

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return resource
        return guest_policy_pb2.OsconfigBetaGuestPolicyPackageRepositoriesAptArchiveTypeEnum.Name(
            resource
        )[
            len("OsconfigBetaGuestPolicyPackageRepositoriesAptArchiveTypeEnum") :
        ]


class GuestPolicyRecipesInstallStepsArchiveExtractionTypeEnum(object):
    @classmethod
    def to_proto(self, resource):
        if not resource:
            return resource
        return guest_policy_pb2.OsconfigBetaGuestPolicyRecipesInstallStepsArchiveExtractionTypeEnum.Value(
            "OsconfigBetaGuestPolicyRecipesInstallStepsArchiveExtractionTypeEnum%s"
            % resource
        )

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return resource
        return guest_policy_pb2.OsconfigBetaGuestPolicyRecipesInstallStepsArchiveExtractionTypeEnum.Name(
            resource
        )[
            len("OsconfigBetaGuestPolicyRecipesInstallStepsArchiveExtractionTypeEnum") :
        ]


class GuestPolicyRecipesInstallStepsScriptRunInterpreterEnum(object):
    @classmethod
    def to_proto(self, resource):
        if not resource:
            return resource
        return guest_policy_pb2.OsconfigBetaGuestPolicyRecipesInstallStepsScriptRunInterpreterEnum.Value(
            "OsconfigBetaGuestPolicyRecipesInstallStepsScriptRunInterpreterEnum%s"
            % resource
        )

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return resource
        return guest_policy_pb2.OsconfigBetaGuestPolicyRecipesInstallStepsScriptRunInterpreterEnum.Name(
            resource
        )[
            len("OsconfigBetaGuestPolicyRecipesInstallStepsScriptRunInterpreterEnum") :
        ]


class GuestPolicyRecipesUpdateStepsArchiveExtractionTypeEnum(object):
    @classmethod
    def to_proto(self, resource):
        if not resource:
            return resource
        return guest_policy_pb2.OsconfigBetaGuestPolicyRecipesUpdateStepsArchiveExtractionTypeEnum.Value(
            "OsconfigBetaGuestPolicyRecipesUpdateStepsArchiveExtractionTypeEnum%s"
            % resource
        )

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return resource
        return guest_policy_pb2.OsconfigBetaGuestPolicyRecipesUpdateStepsArchiveExtractionTypeEnum.Name(
            resource
        )[
            len("OsconfigBetaGuestPolicyRecipesUpdateStepsArchiveExtractionTypeEnum") :
        ]


class GuestPolicyRecipesUpdateStepsScriptRunInterpreterEnum(object):
    @classmethod
    def to_proto(self, resource):
        if not resource:
            return resource
        return guest_policy_pb2.OsconfigBetaGuestPolicyRecipesUpdateStepsScriptRunInterpreterEnum.Value(
            "OsconfigBetaGuestPolicyRecipesUpdateStepsScriptRunInterpreterEnum%s"
            % resource
        )

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return resource
        return guest_policy_pb2.OsconfigBetaGuestPolicyRecipesUpdateStepsScriptRunInterpreterEnum.Name(
            resource
        )[
            len("OsconfigBetaGuestPolicyRecipesUpdateStepsScriptRunInterpreterEnum") :
        ]


class GuestPolicyRecipesDesiredStateEnum(object):
    @classmethod
    def to_proto(self, resource):
        if not resource:
            return resource
        return guest_policy_pb2.OsconfigBetaGuestPolicyRecipesDesiredStateEnum.Value(
            "OsconfigBetaGuestPolicyRecipesDesiredStateEnum%s" % resource
        )

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return resource
        return guest_policy_pb2.OsconfigBetaGuestPolicyRecipesDesiredStateEnum.Name(
            resource
        )[len("OsconfigBetaGuestPolicyRecipesDesiredStateEnum") :]


class Primitive(object):
    @classmethod
    def to_proto(self, s):
        if not s:
            return ""
        return s

    @classmethod
    def from_proto(self, s):
        return s
