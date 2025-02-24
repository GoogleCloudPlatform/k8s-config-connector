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
from google3.cloud.graphite.mmv2.services.google.cloud_build import build_trigger_pb2
from google3.cloud.graphite.mmv2.services.google.cloud_build import (
    build_trigger_pb2_grpc,
)

from typing import List


class BuildTrigger(object):
    def __init__(
        self,
        name: str = None,
        description: str = None,
        tags: list = None,
        disabled: bool = None,
        substitutions: dict = None,
        filename: str = None,
        ignored_files: list = None,
        included_files: list = None,
        trigger_template: dict = None,
        github: dict = None,
        project: str = None,
        build: dict = None,
        id: str = None,
        create_time: str = None,
        service_account_file: str = "",
    ):

        channel.initialize()
        self.name = name
        self.description = description
        self.tags = tags
        self.disabled = disabled
        self.substitutions = substitutions
        self.filename = filename
        self.ignored_files = ignored_files
        self.included_files = included_files
        self.trigger_template = trigger_template
        self.github = github
        self.project = project
        self.build = build
        self.service_account_file = service_account_file

    def apply(self):
        stub = build_trigger_pb2_grpc.CloudbuildBetaBuildTriggerServiceStub(
            channel.Channel()
        )
        request = build_trigger_pb2.ApplyCloudbuildBetaBuildTriggerRequest()
        if Primitive.to_proto(self.name):
            request.resource.name = Primitive.to_proto(self.name)

        if Primitive.to_proto(self.description):
            request.resource.description = Primitive.to_proto(self.description)

        if Primitive.to_proto(self.tags):
            request.resource.tags.extend(Primitive.to_proto(self.tags))
        if Primitive.to_proto(self.disabled):
            request.resource.disabled = Primitive.to_proto(self.disabled)

        if Primitive.to_proto(self.substitutions):
            request.resource.substitutions = Primitive.to_proto(self.substitutions)

        if Primitive.to_proto(self.filename):
            request.resource.filename = Primitive.to_proto(self.filename)

        if Primitive.to_proto(self.ignored_files):
            request.resource.ignored_files.extend(
                Primitive.to_proto(self.ignored_files)
            )
        if Primitive.to_proto(self.included_files):
            request.resource.included_files.extend(
                Primitive.to_proto(self.included_files)
            )
        if BuildTriggerTriggerTemplate.to_proto(self.trigger_template):
            request.resource.trigger_template.CopyFrom(
                BuildTriggerTriggerTemplate.to_proto(self.trigger_template)
            )
        else:
            request.resource.ClearField("trigger_template")
        if BuildTriggerGithub.to_proto(self.github):
            request.resource.github.CopyFrom(BuildTriggerGithub.to_proto(self.github))
        else:
            request.resource.ClearField("github")
        if Primitive.to_proto(self.project):
            request.resource.project = Primitive.to_proto(self.project)

        if BuildTriggerBuild.to_proto(self.build):
            request.resource.build.CopyFrom(BuildTriggerBuild.to_proto(self.build))
        else:
            request.resource.ClearField("build")
        request.service_account_file = self.service_account_file

        response = stub.ApplyCloudbuildBetaBuildTrigger(request)
        self.name = Primitive.from_proto(response.name)
        self.description = Primitive.from_proto(response.description)
        self.tags = Primitive.from_proto(response.tags)
        self.disabled = Primitive.from_proto(response.disabled)
        self.substitutions = Primitive.from_proto(response.substitutions)
        self.filename = Primitive.from_proto(response.filename)
        self.ignored_files = Primitive.from_proto(response.ignored_files)
        self.included_files = Primitive.from_proto(response.included_files)
        self.trigger_template = BuildTriggerTriggerTemplate.from_proto(
            response.trigger_template
        )
        self.github = BuildTriggerGithub.from_proto(response.github)
        self.project = Primitive.from_proto(response.project)
        self.build = BuildTriggerBuild.from_proto(response.build)
        self.id = Primitive.from_proto(response.id)
        self.create_time = Primitive.from_proto(response.create_time)

    def delete(self):
        stub = build_trigger_pb2_grpc.CloudbuildBetaBuildTriggerServiceStub(
            channel.Channel()
        )
        request = build_trigger_pb2.DeleteCloudbuildBetaBuildTriggerRequest()
        request.service_account_file = self.service_account_file
        if Primitive.to_proto(self.name):
            request.resource.name = Primitive.to_proto(self.name)

        if Primitive.to_proto(self.description):
            request.resource.description = Primitive.to_proto(self.description)

        if Primitive.to_proto(self.tags):
            request.resource.tags.extend(Primitive.to_proto(self.tags))
        if Primitive.to_proto(self.disabled):
            request.resource.disabled = Primitive.to_proto(self.disabled)

        if Primitive.to_proto(self.substitutions):
            request.resource.substitutions = Primitive.to_proto(self.substitutions)

        if Primitive.to_proto(self.filename):
            request.resource.filename = Primitive.to_proto(self.filename)

        if Primitive.to_proto(self.ignored_files):
            request.resource.ignored_files.extend(
                Primitive.to_proto(self.ignored_files)
            )
        if Primitive.to_proto(self.included_files):
            request.resource.included_files.extend(
                Primitive.to_proto(self.included_files)
            )
        if BuildTriggerTriggerTemplate.to_proto(self.trigger_template):
            request.resource.trigger_template.CopyFrom(
                BuildTriggerTriggerTemplate.to_proto(self.trigger_template)
            )
        else:
            request.resource.ClearField("trigger_template")
        if BuildTriggerGithub.to_proto(self.github):
            request.resource.github.CopyFrom(BuildTriggerGithub.to_proto(self.github))
        else:
            request.resource.ClearField("github")
        if Primitive.to_proto(self.project):
            request.resource.project = Primitive.to_proto(self.project)

        if BuildTriggerBuild.to_proto(self.build):
            request.resource.build.CopyFrom(BuildTriggerBuild.to_proto(self.build))
        else:
            request.resource.ClearField("build")
        response = stub.DeleteCloudbuildBetaBuildTrigger(request)

    @classmethod
    def list(self, project, service_account_file=""):
        stub = build_trigger_pb2_grpc.CloudbuildBetaBuildTriggerServiceStub(
            channel.Channel()
        )
        request = build_trigger_pb2.ListCloudbuildBetaBuildTriggerRequest()
        request.service_account_file = service_account_file
        request.Project = project

        return stub.ListCloudbuildBetaBuildTrigger(request).items

    def to_proto(self):
        resource = build_trigger_pb2.CloudbuildBetaBuildTrigger()
        if Primitive.to_proto(self.name):
            resource.name = Primitive.to_proto(self.name)
        if Primitive.to_proto(self.description):
            resource.description = Primitive.to_proto(self.description)
        if Primitive.to_proto(self.tags):
            resource.tags.extend(Primitive.to_proto(self.tags))
        if Primitive.to_proto(self.disabled):
            resource.disabled = Primitive.to_proto(self.disabled)
        if Primitive.to_proto(self.substitutions):
            resource.substitutions = Primitive.to_proto(self.substitutions)
        if Primitive.to_proto(self.filename):
            resource.filename = Primitive.to_proto(self.filename)
        if Primitive.to_proto(self.ignored_files):
            resource.ignored_files.extend(Primitive.to_proto(self.ignored_files))
        if Primitive.to_proto(self.included_files):
            resource.included_files.extend(Primitive.to_proto(self.included_files))
        if BuildTriggerTriggerTemplate.to_proto(self.trigger_template):
            resource.trigger_template.CopyFrom(
                BuildTriggerTriggerTemplate.to_proto(self.trigger_template)
            )
        else:
            resource.ClearField("trigger_template")
        if BuildTriggerGithub.to_proto(self.github):
            resource.github.CopyFrom(BuildTriggerGithub.to_proto(self.github))
        else:
            resource.ClearField("github")
        if Primitive.to_proto(self.project):
            resource.project = Primitive.to_proto(self.project)
        if BuildTriggerBuild.to_proto(self.build):
            resource.build.CopyFrom(BuildTriggerBuild.to_proto(self.build))
        else:
            resource.ClearField("build")
        return resource


class BuildTriggerTriggerTemplate(object):
    def __init__(
        self,
        project_id: str = None,
        repo_name: str = None,
        branch_name: str = None,
        tag_name: str = None,
        commit_sha: str = None,
        dir: str = None,
        invert_regex: bool = None,
    ):
        self.project_id = project_id
        self.repo_name = repo_name
        self.branch_name = branch_name
        self.tag_name = tag_name
        self.commit_sha = commit_sha
        self.dir = dir
        self.invert_regex = invert_regex

    @classmethod
    def to_proto(self, resource):
        if not resource:
            return None

        res = build_trigger_pb2.CloudbuildBetaBuildTriggerTriggerTemplate()
        if Primitive.to_proto(resource.project_id):
            res.project_id = Primitive.to_proto(resource.project_id)
        if Primitive.to_proto(resource.repo_name):
            res.repo_name = Primitive.to_proto(resource.repo_name)
        if Primitive.to_proto(resource.branch_name):
            res.branch_name = Primitive.to_proto(resource.branch_name)
        if Primitive.to_proto(resource.tag_name):
            res.tag_name = Primitive.to_proto(resource.tag_name)
        if Primitive.to_proto(resource.commit_sha):
            res.commit_sha = Primitive.to_proto(resource.commit_sha)
        if Primitive.to_proto(resource.dir):
            res.dir = Primitive.to_proto(resource.dir)
        if Primitive.to_proto(resource.invert_regex):
            res.invert_regex = Primitive.to_proto(resource.invert_regex)
        return res

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return None

        return BuildTriggerTriggerTemplate(
            project_id=Primitive.from_proto(resource.project_id),
            repo_name=Primitive.from_proto(resource.repo_name),
            branch_name=Primitive.from_proto(resource.branch_name),
            tag_name=Primitive.from_proto(resource.tag_name),
            commit_sha=Primitive.from_proto(resource.commit_sha),
            dir=Primitive.from_proto(resource.dir),
            invert_regex=Primitive.from_proto(resource.invert_regex),
        )


class BuildTriggerTriggerTemplateArray(object):
    @classmethod
    def to_proto(self, resources):
        if not resources:
            return resources
        return [BuildTriggerTriggerTemplate.to_proto(i) for i in resources]

    @classmethod
    def from_proto(self, resources):
        return [BuildTriggerTriggerTemplate.from_proto(i) for i in resources]


class BuildTriggerGithub(object):
    def __init__(
        self,
        owner: str = None,
        name: str = None,
        pull_request: dict = None,
        push: dict = None,
    ):
        self.owner = owner
        self.name = name
        self.pull_request = pull_request
        self.push = push

    @classmethod
    def to_proto(self, resource):
        if not resource:
            return None

        res = build_trigger_pb2.CloudbuildBetaBuildTriggerGithub()
        if Primitive.to_proto(resource.owner):
            res.owner = Primitive.to_proto(resource.owner)
        if Primitive.to_proto(resource.name):
            res.name = Primitive.to_proto(resource.name)
        if BuildTriggerGithubPullRequest.to_proto(resource.pull_request):
            res.pull_request.CopyFrom(
                BuildTriggerGithubPullRequest.to_proto(resource.pull_request)
            )
        else:
            res.ClearField("pull_request")
        if BuildTriggerGithubPush.to_proto(resource.push):
            res.push.CopyFrom(BuildTriggerGithubPush.to_proto(resource.push))
        else:
            res.ClearField("push")
        return res

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return None

        return BuildTriggerGithub(
            owner=Primitive.from_proto(resource.owner),
            name=Primitive.from_proto(resource.name),
            pull_request=BuildTriggerGithubPullRequest.from_proto(
                resource.pull_request
            ),
            push=BuildTriggerGithubPush.from_proto(resource.push),
        )


class BuildTriggerGithubArray(object):
    @classmethod
    def to_proto(self, resources):
        if not resources:
            return resources
        return [BuildTriggerGithub.to_proto(i) for i in resources]

    @classmethod
    def from_proto(self, resources):
        return [BuildTriggerGithub.from_proto(i) for i in resources]


class BuildTriggerGithubPullRequest(object):
    def __init__(
        self, branch: str = None, comment_control: str = None, invert_regex: bool = None
    ):
        self.branch = branch
        self.comment_control = comment_control
        self.invert_regex = invert_regex

    @classmethod
    def to_proto(self, resource):
        if not resource:
            return None

        res = build_trigger_pb2.CloudbuildBetaBuildTriggerGithubPullRequest()
        if Primitive.to_proto(resource.branch):
            res.branch = Primitive.to_proto(resource.branch)
        if BuildTriggerGithubPullRequestCommentControlEnum.to_proto(
            resource.comment_control
        ):
            res.comment_control = BuildTriggerGithubPullRequestCommentControlEnum.to_proto(
                resource.comment_control
            )
        if Primitive.to_proto(resource.invert_regex):
            res.invert_regex = Primitive.to_proto(resource.invert_regex)
        return res

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return None

        return BuildTriggerGithubPullRequest(
            branch=Primitive.from_proto(resource.branch),
            comment_control=BuildTriggerGithubPullRequestCommentControlEnum.from_proto(
                resource.comment_control
            ),
            invert_regex=Primitive.from_proto(resource.invert_regex),
        )


class BuildTriggerGithubPullRequestArray(object):
    @classmethod
    def to_proto(self, resources):
        if not resources:
            return resources
        return [BuildTriggerGithubPullRequest.to_proto(i) for i in resources]

    @classmethod
    def from_proto(self, resources):
        return [BuildTriggerGithubPullRequest.from_proto(i) for i in resources]


class BuildTriggerGithubPush(object):
    def __init__(self, branch: str = None, tag: str = None, invert_regex: bool = None):
        self.branch = branch
        self.tag = tag
        self.invert_regex = invert_regex

    @classmethod
    def to_proto(self, resource):
        if not resource:
            return None

        res = build_trigger_pb2.CloudbuildBetaBuildTriggerGithubPush()
        if Primitive.to_proto(resource.branch):
            res.branch = Primitive.to_proto(resource.branch)
        if Primitive.to_proto(resource.tag):
            res.tag = Primitive.to_proto(resource.tag)
        if Primitive.to_proto(resource.invert_regex):
            res.invert_regex = Primitive.to_proto(resource.invert_regex)
        return res

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return None

        return BuildTriggerGithubPush(
            branch=Primitive.from_proto(resource.branch),
            tag=Primitive.from_proto(resource.tag),
            invert_regex=Primitive.from_proto(resource.invert_regex),
        )


class BuildTriggerGithubPushArray(object):
    @classmethod
    def to_proto(self, resources):
        if not resources:
            return resources
        return [BuildTriggerGithubPush.to_proto(i) for i in resources]

    @classmethod
    def from_proto(self, resources):
        return [BuildTriggerGithubPush.from_proto(i) for i in resources]


class BuildTriggerBuild(object):
    def __init__(
        self,
        tags: list = None,
        images: list = None,
        substitutions: dict = None,
        queue_ttl: str = None,
        logs_bucket: str = None,
        timeout: str = None,
        secrets: list = None,
        steps: list = None,
        source: dict = None,
    ):
        self.tags = tags
        self.images = images
        self.substitutions = substitutions
        self.queue_ttl = queue_ttl
        self.logs_bucket = logs_bucket
        self.timeout = timeout
        self.secrets = secrets
        self.steps = steps
        self.source = source

    @classmethod
    def to_proto(self, resource):
        if not resource:
            return None

        res = build_trigger_pb2.CloudbuildBetaBuildTriggerBuild()
        if Primitive.to_proto(resource.tags):
            res.tags.extend(Primitive.to_proto(resource.tags))
        if Primitive.to_proto(resource.images):
            res.images.extend(Primitive.to_proto(resource.images))
        if Primitive.to_proto(resource.substitutions):
            res.substitutions = Primitive.to_proto(resource.substitutions)
        if Primitive.to_proto(resource.queue_ttl):
            res.queue_ttl = Primitive.to_proto(resource.queue_ttl)
        if Primitive.to_proto(resource.logs_bucket):
            res.logs_bucket = Primitive.to_proto(resource.logs_bucket)
        if Primitive.to_proto(resource.timeout):
            res.timeout = Primitive.to_proto(resource.timeout)
        if BuildTriggerBuildSecretsArray.to_proto(resource.secrets):
            res.secrets.extend(BuildTriggerBuildSecretsArray.to_proto(resource.secrets))
        if BuildTriggerBuildStepsArray.to_proto(resource.steps):
            res.steps.extend(BuildTriggerBuildStepsArray.to_proto(resource.steps))
        if BuildTriggerBuildSource.to_proto(resource.source):
            res.source.CopyFrom(BuildTriggerBuildSource.to_proto(resource.source))
        else:
            res.ClearField("source")
        return res

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return None

        return BuildTriggerBuild(
            tags=Primitive.from_proto(resource.tags),
            images=Primitive.from_proto(resource.images),
            substitutions=Primitive.from_proto(resource.substitutions),
            queue_ttl=Primitive.from_proto(resource.queue_ttl),
            logs_bucket=Primitive.from_proto(resource.logs_bucket),
            timeout=Primitive.from_proto(resource.timeout),
            secrets=BuildTriggerBuildSecretsArray.from_proto(resource.secrets),
            steps=BuildTriggerBuildStepsArray.from_proto(resource.steps),
            source=BuildTriggerBuildSource.from_proto(resource.source),
        )


class BuildTriggerBuildArray(object):
    @classmethod
    def to_proto(self, resources):
        if not resources:
            return resources
        return [BuildTriggerBuild.to_proto(i) for i in resources]

    @classmethod
    def from_proto(self, resources):
        return [BuildTriggerBuild.from_proto(i) for i in resources]


class BuildTriggerBuildSecrets(object):
    def __init__(self, kms_key_name: str = None, secret_env: dict = None):
        self.kms_key_name = kms_key_name
        self.secret_env = secret_env

    @classmethod
    def to_proto(self, resource):
        if not resource:
            return None

        res = build_trigger_pb2.CloudbuildBetaBuildTriggerBuildSecrets()
        if Primitive.to_proto(resource.kms_key_name):
            res.kms_key_name = Primitive.to_proto(resource.kms_key_name)
        if Primitive.to_proto(resource.secret_env):
            res.secret_env = Primitive.to_proto(resource.secret_env)
        return res

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return None

        return BuildTriggerBuildSecrets(
            kms_key_name=Primitive.from_proto(resource.kms_key_name),
            secret_env=Primitive.from_proto(resource.secret_env),
        )


class BuildTriggerBuildSecretsArray(object):
    @classmethod
    def to_proto(self, resources):
        if not resources:
            return resources
        return [BuildTriggerBuildSecrets.to_proto(i) for i in resources]

    @classmethod
    def from_proto(self, resources):
        return [BuildTriggerBuildSecrets.from_proto(i) for i in resources]


class BuildTriggerBuildSteps(object):
    def __init__(
        self,
        name: str = None,
        env: list = None,
        args: list = None,
        dir: str = None,
        id: str = None,
        wait_for: list = None,
        entrypoint: str = None,
        secret_env: list = None,
        volumes: list = None,
        timing: dict = None,
        pull_timing: dict = None,
        timeout: str = None,
        status: str = None,
    ):
        self.name = name
        self.env = env
        self.args = args
        self.dir = dir
        self.id = id
        self.wait_for = wait_for
        self.entrypoint = entrypoint
        self.secret_env = secret_env
        self.volumes = volumes
        self.timing = timing
        self.pull_timing = pull_timing
        self.timeout = timeout
        self.status = status

    @classmethod
    def to_proto(self, resource):
        if not resource:
            return None

        res = build_trigger_pb2.CloudbuildBetaBuildTriggerBuildSteps()
        if Primitive.to_proto(resource.name):
            res.name = Primitive.to_proto(resource.name)
        if Primitive.to_proto(resource.env):
            res.env.extend(Primitive.to_proto(resource.env))
        if Primitive.to_proto(resource.args):
            res.args.extend(Primitive.to_proto(resource.args))
        if Primitive.to_proto(resource.dir):
            res.dir = Primitive.to_proto(resource.dir)
        if Primitive.to_proto(resource.id):
            res.id = Primitive.to_proto(resource.id)
        if Primitive.to_proto(resource.wait_for):
            res.wait_for.extend(Primitive.to_proto(resource.wait_for))
        if Primitive.to_proto(resource.entrypoint):
            res.entrypoint = Primitive.to_proto(resource.entrypoint)
        if Primitive.to_proto(resource.secret_env):
            res.secret_env.extend(Primitive.to_proto(resource.secret_env))
        if BuildTriggerBuildStepsVolumesArray.to_proto(resource.volumes):
            res.volumes.extend(
                BuildTriggerBuildStepsVolumesArray.to_proto(resource.volumes)
            )
        if BuildTriggerBuildStepsTiming.to_proto(resource.timing):
            res.timing.CopyFrom(BuildTriggerBuildStepsTiming.to_proto(resource.timing))
        else:
            res.ClearField("timing")
        if BuildTriggerBuildStepsPullTiming.to_proto(resource.pull_timing):
            res.pull_timing.CopyFrom(
                BuildTriggerBuildStepsPullTiming.to_proto(resource.pull_timing)
            )
        else:
            res.ClearField("pull_timing")
        if Primitive.to_proto(resource.timeout):
            res.timeout = Primitive.to_proto(resource.timeout)
        if BuildTriggerBuildStepsStatusEnum.to_proto(resource.status):
            res.status = BuildTriggerBuildStepsStatusEnum.to_proto(resource.status)
        return res

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return None

        return BuildTriggerBuildSteps(
            name=Primitive.from_proto(resource.name),
            env=Primitive.from_proto(resource.env),
            args=Primitive.from_proto(resource.args),
            dir=Primitive.from_proto(resource.dir),
            id=Primitive.from_proto(resource.id),
            wait_for=Primitive.from_proto(resource.wait_for),
            entrypoint=Primitive.from_proto(resource.entrypoint),
            secret_env=Primitive.from_proto(resource.secret_env),
            volumes=BuildTriggerBuildStepsVolumesArray.from_proto(resource.volumes),
            timing=BuildTriggerBuildStepsTiming.from_proto(resource.timing),
            pull_timing=BuildTriggerBuildStepsPullTiming.from_proto(
                resource.pull_timing
            ),
            timeout=Primitive.from_proto(resource.timeout),
            status=BuildTriggerBuildStepsStatusEnum.from_proto(resource.status),
        )


class BuildTriggerBuildStepsArray(object):
    @classmethod
    def to_proto(self, resources):
        if not resources:
            return resources
        return [BuildTriggerBuildSteps.to_proto(i) for i in resources]

    @classmethod
    def from_proto(self, resources):
        return [BuildTriggerBuildSteps.from_proto(i) for i in resources]


class BuildTriggerBuildStepsVolumes(object):
    def __init__(self, name: str = None, path: str = None):
        self.name = name
        self.path = path

    @classmethod
    def to_proto(self, resource):
        if not resource:
            return None

        res = build_trigger_pb2.CloudbuildBetaBuildTriggerBuildStepsVolumes()
        if Primitive.to_proto(resource.name):
            res.name = Primitive.to_proto(resource.name)
        if Primitive.to_proto(resource.path):
            res.path = Primitive.to_proto(resource.path)
        return res

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return None

        return BuildTriggerBuildStepsVolumes(
            name=Primitive.from_proto(resource.name),
            path=Primitive.from_proto(resource.path),
        )


class BuildTriggerBuildStepsVolumesArray(object):
    @classmethod
    def to_proto(self, resources):
        if not resources:
            return resources
        return [BuildTriggerBuildStepsVolumes.to_proto(i) for i in resources]

    @classmethod
    def from_proto(self, resources):
        return [BuildTriggerBuildStepsVolumes.from_proto(i) for i in resources]


class BuildTriggerBuildStepsTiming(object):
    def __init__(self, start_time: str = None, end_time: str = None):
        self.start_time = start_time
        self.end_time = end_time

    @classmethod
    def to_proto(self, resource):
        if not resource:
            return None

        res = build_trigger_pb2.CloudbuildBetaBuildTriggerBuildStepsTiming()
        if Primitive.to_proto(resource.start_time):
            res.start_time = Primitive.to_proto(resource.start_time)
        if Primitive.to_proto(resource.end_time):
            res.end_time = Primitive.to_proto(resource.end_time)
        return res

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return None

        return BuildTriggerBuildStepsTiming(
            start_time=Primitive.from_proto(resource.start_time),
            end_time=Primitive.from_proto(resource.end_time),
        )


class BuildTriggerBuildStepsTimingArray(object):
    @classmethod
    def to_proto(self, resources):
        if not resources:
            return resources
        return [BuildTriggerBuildStepsTiming.to_proto(i) for i in resources]

    @classmethod
    def from_proto(self, resources):
        return [BuildTriggerBuildStepsTiming.from_proto(i) for i in resources]


class BuildTriggerBuildStepsPullTiming(object):
    def __init__(self, start_time: str = None, end_time: str = None):
        self.start_time = start_time
        self.end_time = end_time

    @classmethod
    def to_proto(self, resource):
        if not resource:
            return None

        res = build_trigger_pb2.CloudbuildBetaBuildTriggerBuildStepsPullTiming()
        if Primitive.to_proto(resource.start_time):
            res.start_time = Primitive.to_proto(resource.start_time)
        if Primitive.to_proto(resource.end_time):
            res.end_time = Primitive.to_proto(resource.end_time)
        return res

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return None

        return BuildTriggerBuildStepsPullTiming(
            start_time=Primitive.from_proto(resource.start_time),
            end_time=Primitive.from_proto(resource.end_time),
        )


class BuildTriggerBuildStepsPullTimingArray(object):
    @classmethod
    def to_proto(self, resources):
        if not resources:
            return resources
        return [BuildTriggerBuildStepsPullTiming.to_proto(i) for i in resources]

    @classmethod
    def from_proto(self, resources):
        return [BuildTriggerBuildStepsPullTiming.from_proto(i) for i in resources]


class BuildTriggerBuildSource(object):
    def __init__(self, storage_source: dict = None, repo_source: dict = None):
        self.storage_source = storage_source
        self.repo_source = repo_source

    @classmethod
    def to_proto(self, resource):
        if not resource:
            return None

        res = build_trigger_pb2.CloudbuildBetaBuildTriggerBuildSource()
        if BuildTriggerBuildSourceStorageSource.to_proto(resource.storage_source):
            res.storage_source.CopyFrom(
                BuildTriggerBuildSourceStorageSource.to_proto(resource.storage_source)
            )
        else:
            res.ClearField("storage_source")
        if BuildTriggerBuildSourceRepoSource.to_proto(resource.repo_source):
            res.repo_source.CopyFrom(
                BuildTriggerBuildSourceRepoSource.to_proto(resource.repo_source)
            )
        else:
            res.ClearField("repo_source")
        return res

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return None

        return BuildTriggerBuildSource(
            storage_source=BuildTriggerBuildSourceStorageSource.from_proto(
                resource.storage_source
            ),
            repo_source=BuildTriggerBuildSourceRepoSource.from_proto(
                resource.repo_source
            ),
        )


class BuildTriggerBuildSourceArray(object):
    @classmethod
    def to_proto(self, resources):
        if not resources:
            return resources
        return [BuildTriggerBuildSource.to_proto(i) for i in resources]

    @classmethod
    def from_proto(self, resources):
        return [BuildTriggerBuildSource.from_proto(i) for i in resources]


class BuildTriggerBuildSourceStorageSource(object):
    def __init__(self, bucket: str = None, object: str = None, generation: str = None):
        self.bucket = bucket
        self.object = object
        self.generation = generation

    @classmethod
    def to_proto(self, resource):
        if not resource:
            return None

        res = build_trigger_pb2.CloudbuildBetaBuildTriggerBuildSourceStorageSource()
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

        return BuildTriggerBuildSourceStorageSource(
            bucket=Primitive.from_proto(resource.bucket),
            object=Primitive.from_proto(resource.object),
            generation=Primitive.from_proto(resource.generation),
        )


class BuildTriggerBuildSourceStorageSourceArray(object):
    @classmethod
    def to_proto(self, resources):
        if not resources:
            return resources
        return [BuildTriggerBuildSourceStorageSource.to_proto(i) for i in resources]

    @classmethod
    def from_proto(self, resources):
        return [BuildTriggerBuildSourceStorageSource.from_proto(i) for i in resources]


class BuildTriggerBuildSourceRepoSource(object):
    def __init__(
        self,
        project_id: str = None,
        repo_name: str = None,
        branch_name: str = None,
        tag_name: str = None,
        commit_sha: str = None,
        dir: str = None,
        invert_regex: bool = None,
        substitutions: dict = None,
    ):
        self.project_id = project_id
        self.repo_name = repo_name
        self.branch_name = branch_name
        self.tag_name = tag_name
        self.commit_sha = commit_sha
        self.dir = dir
        self.invert_regex = invert_regex
        self.substitutions = substitutions

    @classmethod
    def to_proto(self, resource):
        if not resource:
            return None

        res = build_trigger_pb2.CloudbuildBetaBuildTriggerBuildSourceRepoSource()
        if Primitive.to_proto(resource.project_id):
            res.project_id = Primitive.to_proto(resource.project_id)
        if Primitive.to_proto(resource.repo_name):
            res.repo_name = Primitive.to_proto(resource.repo_name)
        if Primitive.to_proto(resource.branch_name):
            res.branch_name = Primitive.to_proto(resource.branch_name)
        if Primitive.to_proto(resource.tag_name):
            res.tag_name = Primitive.to_proto(resource.tag_name)
        if Primitive.to_proto(resource.commit_sha):
            res.commit_sha = Primitive.to_proto(resource.commit_sha)
        if Primitive.to_proto(resource.dir):
            res.dir = Primitive.to_proto(resource.dir)
        if Primitive.to_proto(resource.invert_regex):
            res.invert_regex = Primitive.to_proto(resource.invert_regex)
        if Primitive.to_proto(resource.substitutions):
            res.substitutions = Primitive.to_proto(resource.substitutions)
        return res

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return None

        return BuildTriggerBuildSourceRepoSource(
            project_id=Primitive.from_proto(resource.project_id),
            repo_name=Primitive.from_proto(resource.repo_name),
            branch_name=Primitive.from_proto(resource.branch_name),
            tag_name=Primitive.from_proto(resource.tag_name),
            commit_sha=Primitive.from_proto(resource.commit_sha),
            dir=Primitive.from_proto(resource.dir),
            invert_regex=Primitive.from_proto(resource.invert_regex),
            substitutions=Primitive.from_proto(resource.substitutions),
        )


class BuildTriggerBuildSourceRepoSourceArray(object):
    @classmethod
    def to_proto(self, resources):
        if not resources:
            return resources
        return [BuildTriggerBuildSourceRepoSource.to_proto(i) for i in resources]

    @classmethod
    def from_proto(self, resources):
        return [BuildTriggerBuildSourceRepoSource.from_proto(i) for i in resources]


class BuildTriggerGithubPullRequestCommentControlEnum(object):
    @classmethod
    def to_proto(self, resource):
        if not resource:
            return resource
        return build_trigger_pb2.CloudbuildBetaBuildTriggerGithubPullRequestCommentControlEnum.Value(
            "CloudbuildBetaBuildTriggerGithubPullRequestCommentControlEnum%s" % resource
        )

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return resource
        return build_trigger_pb2.CloudbuildBetaBuildTriggerGithubPullRequestCommentControlEnum.Name(
            resource
        )[
            len("CloudbuildBetaBuildTriggerGithubPullRequestCommentControlEnum") :
        ]


class BuildTriggerBuildStepsStatusEnum(object):
    @classmethod
    def to_proto(self, resource):
        if not resource:
            return resource
        return build_trigger_pb2.CloudbuildBetaBuildTriggerBuildStepsStatusEnum.Value(
            "CloudbuildBetaBuildTriggerBuildStepsStatusEnum%s" % resource
        )

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return resource
        return build_trigger_pb2.CloudbuildBetaBuildTriggerBuildStepsStatusEnum.Name(
            resource
        )[len("CloudbuildBetaBuildTriggerBuildStepsStatusEnum") :]


class Primitive(object):
    @classmethod
    def to_proto(self, s):
        if not s:
            return ""
        return s

    @classmethod
    def from_proto(self, s):
        return s
