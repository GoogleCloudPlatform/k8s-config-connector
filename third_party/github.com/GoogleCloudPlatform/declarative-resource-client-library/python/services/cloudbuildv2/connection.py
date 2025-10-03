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
from google3.cloud.graphite.mmv2.services.google.cloudbuildv2 import connection_pb2
from google3.cloud.graphite.mmv2.services.google.cloudbuildv2 import connection_pb2_grpc

from typing import List


class Connection(object):
    def __init__(
        self,
        name: str = None,
        create_time: str = None,
        update_time: str = None,
        github_config: dict = None,
        github_enterprise_config: dict = None,
        gitlab_config: dict = None,
        installation_state: dict = None,
        disabled: bool = None,
        reconciling: bool = None,
        annotations: dict = None,
        etag: str = None,
        project: str = None,
        location: str = None,
        service_account_file: str = "",
    ):
        channel.initialize()
        self.name = name
        self.github_config = github_config
        self.github_enterprise_config = github_enterprise_config
        self.gitlab_config = gitlab_config
        self.disabled = disabled
        self.annotations = annotations
        self.project = project
        self.location = location
        self.service_account_file = service_account_file

    def apply(self):
        stub = connection_pb2_grpc.Cloudbuildv2ConnectionServiceStub(channel.Channel())
        request = connection_pb2.ApplyCloudbuildv2ConnectionRequest()
        if Primitive.to_proto(self.name):
            request.resource.name = Primitive.to_proto(self.name)

        if ConnectionGithubConfig.to_proto(self.github_config):
            request.resource.github_config.CopyFrom(
                ConnectionGithubConfig.to_proto(self.github_config)
            )
        else:
            request.resource.ClearField("github_config")
        if ConnectionGithubEnterpriseConfig.to_proto(self.github_enterprise_config):
            request.resource.github_enterprise_config.CopyFrom(
                ConnectionGithubEnterpriseConfig.to_proto(self.github_enterprise_config)
            )
        else:
            request.resource.ClearField("github_enterprise_config")
        if ConnectionGitlabConfig.to_proto(self.gitlab_config):
            request.resource.gitlab_config.CopyFrom(
                ConnectionGitlabConfig.to_proto(self.gitlab_config)
            )
        else:
            request.resource.ClearField("gitlab_config")
        if Primitive.to_proto(self.disabled):
            request.resource.disabled = Primitive.to_proto(self.disabled)

        if Primitive.to_proto(self.annotations):
            request.resource.annotations = Primitive.to_proto(self.annotations)

        if Primitive.to_proto(self.project):
            request.resource.project = Primitive.to_proto(self.project)

        if Primitive.to_proto(self.location):
            request.resource.location = Primitive.to_proto(self.location)

        request.service_account_file = self.service_account_file

        response = stub.ApplyCloudbuildv2Connection(request)
        self.name = Primitive.from_proto(response.name)
        self.create_time = Primitive.from_proto(response.create_time)
        self.update_time = Primitive.from_proto(response.update_time)
        self.github_config = ConnectionGithubConfig.from_proto(response.github_config)
        self.github_enterprise_config = ConnectionGithubEnterpriseConfig.from_proto(
            response.github_enterprise_config
        )
        self.gitlab_config = ConnectionGitlabConfig.from_proto(response.gitlab_config)
        self.installation_state = ConnectionInstallationState.from_proto(
            response.installation_state
        )
        self.disabled = Primitive.from_proto(response.disabled)
        self.reconciling = Primitive.from_proto(response.reconciling)
        self.annotations = Primitive.from_proto(response.annotations)
        self.etag = Primitive.from_proto(response.etag)
        self.project = Primitive.from_proto(response.project)
        self.location = Primitive.from_proto(response.location)

    def delete(self):
        stub = connection_pb2_grpc.Cloudbuildv2ConnectionServiceStub(channel.Channel())
        request = connection_pb2.DeleteCloudbuildv2ConnectionRequest()
        request.service_account_file = self.service_account_file
        if Primitive.to_proto(self.name):
            request.resource.name = Primitive.to_proto(self.name)

        if ConnectionGithubConfig.to_proto(self.github_config):
            request.resource.github_config.CopyFrom(
                ConnectionGithubConfig.to_proto(self.github_config)
            )
        else:
            request.resource.ClearField("github_config")
        if ConnectionGithubEnterpriseConfig.to_proto(self.github_enterprise_config):
            request.resource.github_enterprise_config.CopyFrom(
                ConnectionGithubEnterpriseConfig.to_proto(self.github_enterprise_config)
            )
        else:
            request.resource.ClearField("github_enterprise_config")
        if ConnectionGitlabConfig.to_proto(self.gitlab_config):
            request.resource.gitlab_config.CopyFrom(
                ConnectionGitlabConfig.to_proto(self.gitlab_config)
            )
        else:
            request.resource.ClearField("gitlab_config")
        if Primitive.to_proto(self.disabled):
            request.resource.disabled = Primitive.to_proto(self.disabled)

        if Primitive.to_proto(self.annotations):
            request.resource.annotations = Primitive.to_proto(self.annotations)

        if Primitive.to_proto(self.project):
            request.resource.project = Primitive.to_proto(self.project)

        if Primitive.to_proto(self.location):
            request.resource.location = Primitive.to_proto(self.location)

        response = stub.DeleteCloudbuildv2Connection(request)

    @classmethod
    def list(self, project, location, service_account_file=""):
        stub = connection_pb2_grpc.Cloudbuildv2ConnectionServiceStub(channel.Channel())
        request = connection_pb2.ListCloudbuildv2ConnectionRequest()
        request.service_account_file = service_account_file
        request.Project = project

        request.Location = location

        return stub.ListCloudbuildv2Connection(request).items

    def to_proto(self):
        resource = connection_pb2.Cloudbuildv2Connection()
        if Primitive.to_proto(self.name):
            resource.name = Primitive.to_proto(self.name)
        if ConnectionGithubConfig.to_proto(self.github_config):
            resource.github_config.CopyFrom(
                ConnectionGithubConfig.to_proto(self.github_config)
            )
        else:
            resource.ClearField("github_config")
        if ConnectionGithubEnterpriseConfig.to_proto(self.github_enterprise_config):
            resource.github_enterprise_config.CopyFrom(
                ConnectionGithubEnterpriseConfig.to_proto(self.github_enterprise_config)
            )
        else:
            resource.ClearField("github_enterprise_config")
        if ConnectionGitlabConfig.to_proto(self.gitlab_config):
            resource.gitlab_config.CopyFrom(
                ConnectionGitlabConfig.to_proto(self.gitlab_config)
            )
        else:
            resource.ClearField("gitlab_config")
        if Primitive.to_proto(self.disabled):
            resource.disabled = Primitive.to_proto(self.disabled)
        if Primitive.to_proto(self.annotations):
            resource.annotations = Primitive.to_proto(self.annotations)
        if Primitive.to_proto(self.project):
            resource.project = Primitive.to_proto(self.project)
        if Primitive.to_proto(self.location):
            resource.location = Primitive.to_proto(self.location)
        return resource


class ConnectionGithubConfig(object):
    def __init__(
        self, authorizer_credential: dict = None, app_installation_id: int = None
    ):
        self.authorizer_credential = authorizer_credential
        self.app_installation_id = app_installation_id

    @classmethod
    def to_proto(self, resource):
        if not resource:
            return None

        res = connection_pb2.Cloudbuildv2ConnectionGithubConfig()
        if ConnectionGithubConfigAuthorizerCredential.to_proto(
            resource.authorizer_credential
        ):
            res.authorizer_credential.CopyFrom(
                ConnectionGithubConfigAuthorizerCredential.to_proto(
                    resource.authorizer_credential
                )
            )
        else:
            res.ClearField("authorizer_credential")
        if Primitive.to_proto(resource.app_installation_id):
            res.app_installation_id = Primitive.to_proto(resource.app_installation_id)
        return res

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return None

        return ConnectionGithubConfig(
            authorizer_credential=ConnectionGithubConfigAuthorizerCredential.from_proto(
                resource.authorizer_credential
            ),
            app_installation_id=Primitive.from_proto(resource.app_installation_id),
        )


class ConnectionGithubConfigArray(object):
    @classmethod
    def to_proto(self, resources):
        if not resources:
            return resources
        return [ConnectionGithubConfig.to_proto(i) for i in resources]

    @classmethod
    def from_proto(self, resources):
        return [ConnectionGithubConfig.from_proto(i) for i in resources]


class ConnectionGithubConfigAuthorizerCredential(object):
    def __init__(self, oauth_token_secret_version: str = None, username: str = None):
        self.oauth_token_secret_version = oauth_token_secret_version
        self.username = username

    @classmethod
    def to_proto(self, resource):
        if not resource:
            return None

        res = connection_pb2.Cloudbuildv2ConnectionGithubConfigAuthorizerCredential()
        if Primitive.to_proto(resource.oauth_token_secret_version):
            res.oauth_token_secret_version = Primitive.to_proto(
                resource.oauth_token_secret_version
            )
        if Primitive.to_proto(resource.username):
            res.username = Primitive.to_proto(resource.username)
        return res

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return None

        return ConnectionGithubConfigAuthorizerCredential(
            oauth_token_secret_version=Primitive.from_proto(
                resource.oauth_token_secret_version
            ),
            username=Primitive.from_proto(resource.username),
        )


class ConnectionGithubConfigAuthorizerCredentialArray(object):
    @classmethod
    def to_proto(self, resources):
        if not resources:
            return resources
        return [
            ConnectionGithubConfigAuthorizerCredential.to_proto(i) for i in resources
        ]

    @classmethod
    def from_proto(self, resources):
        return [
            ConnectionGithubConfigAuthorizerCredential.from_proto(i) for i in resources
        ]


class ConnectionGithubEnterpriseConfig(object):
    def __init__(
        self,
        host_uri: str = None,
        app_id: int = None,
        app_slug: str = None,
        private_key_secret_version: str = None,
        webhook_secret_secret_version: str = None,
        app_installation_id: int = None,
        service_directory_config: dict = None,
        ssl_ca: str = None,
    ):
        self.host_uri = host_uri
        self.app_id = app_id
        self.app_slug = app_slug
        self.private_key_secret_version = private_key_secret_version
        self.webhook_secret_secret_version = webhook_secret_secret_version
        self.app_installation_id = app_installation_id
        self.service_directory_config = service_directory_config
        self.ssl_ca = ssl_ca

    @classmethod
    def to_proto(self, resource):
        if not resource:
            return None

        res = connection_pb2.Cloudbuildv2ConnectionGithubEnterpriseConfig()
        if Primitive.to_proto(resource.host_uri):
            res.host_uri = Primitive.to_proto(resource.host_uri)
        if Primitive.to_proto(resource.app_id):
            res.app_id = Primitive.to_proto(resource.app_id)
        if Primitive.to_proto(resource.app_slug):
            res.app_slug = Primitive.to_proto(resource.app_slug)
        if Primitive.to_proto(resource.private_key_secret_version):
            res.private_key_secret_version = Primitive.to_proto(
                resource.private_key_secret_version
            )
        if Primitive.to_proto(resource.webhook_secret_secret_version):
            res.webhook_secret_secret_version = Primitive.to_proto(
                resource.webhook_secret_secret_version
            )
        if Primitive.to_proto(resource.app_installation_id):
            res.app_installation_id = Primitive.to_proto(resource.app_installation_id)
        if ConnectionGithubEnterpriseConfigServiceDirectoryConfig.to_proto(
            resource.service_directory_config
        ):
            res.service_directory_config.CopyFrom(
                ConnectionGithubEnterpriseConfigServiceDirectoryConfig.to_proto(
                    resource.service_directory_config
                )
            )
        else:
            res.ClearField("service_directory_config")
        if Primitive.to_proto(resource.ssl_ca):
            res.ssl_ca = Primitive.to_proto(resource.ssl_ca)
        return res

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return None

        return ConnectionGithubEnterpriseConfig(
            host_uri=Primitive.from_proto(resource.host_uri),
            app_id=Primitive.from_proto(resource.app_id),
            app_slug=Primitive.from_proto(resource.app_slug),
            private_key_secret_version=Primitive.from_proto(
                resource.private_key_secret_version
            ),
            webhook_secret_secret_version=Primitive.from_proto(
                resource.webhook_secret_secret_version
            ),
            app_installation_id=Primitive.from_proto(resource.app_installation_id),
            service_directory_config=ConnectionGithubEnterpriseConfigServiceDirectoryConfig.from_proto(
                resource.service_directory_config
            ),
            ssl_ca=Primitive.from_proto(resource.ssl_ca),
        )


class ConnectionGithubEnterpriseConfigArray(object):
    @classmethod
    def to_proto(self, resources):
        if not resources:
            return resources
        return [ConnectionGithubEnterpriseConfig.to_proto(i) for i in resources]

    @classmethod
    def from_proto(self, resources):
        return [ConnectionGithubEnterpriseConfig.from_proto(i) for i in resources]


class ConnectionGithubEnterpriseConfigServiceDirectoryConfig(object):
    def __init__(self, service: str = None):
        self.service = service

    @classmethod
    def to_proto(self, resource):
        if not resource:
            return None

        res = (
            connection_pb2.Cloudbuildv2ConnectionGithubEnterpriseConfigServiceDirectoryConfig()
        )
        if Primitive.to_proto(resource.service):
            res.service = Primitive.to_proto(resource.service)
        return res

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return None

        return ConnectionGithubEnterpriseConfigServiceDirectoryConfig(
            service=Primitive.from_proto(resource.service),
        )


class ConnectionGithubEnterpriseConfigServiceDirectoryConfigArray(object):
    @classmethod
    def to_proto(self, resources):
        if not resources:
            return resources
        return [
            ConnectionGithubEnterpriseConfigServiceDirectoryConfig.to_proto(i)
            for i in resources
        ]

    @classmethod
    def from_proto(self, resources):
        return [
            ConnectionGithubEnterpriseConfigServiceDirectoryConfig.from_proto(i)
            for i in resources
        ]


class ConnectionGitlabConfig(object):
    def __init__(
        self,
        host_uri: str = None,
        webhook_secret_secret_version: str = None,
        read_authorizer_credential: dict = None,
        authorizer_credential: dict = None,
        service_directory_config: dict = None,
        ssl_ca: str = None,
        server_version: str = None,
    ):
        self.host_uri = host_uri
        self.webhook_secret_secret_version = webhook_secret_secret_version
        self.read_authorizer_credential = read_authorizer_credential
        self.authorizer_credential = authorizer_credential
        self.service_directory_config = service_directory_config
        self.ssl_ca = ssl_ca
        self.server_version = server_version

    @classmethod
    def to_proto(self, resource):
        if not resource:
            return None

        res = connection_pb2.Cloudbuildv2ConnectionGitlabConfig()
        if Primitive.to_proto(resource.host_uri):
            res.host_uri = Primitive.to_proto(resource.host_uri)
        if Primitive.to_proto(resource.webhook_secret_secret_version):
            res.webhook_secret_secret_version = Primitive.to_proto(
                resource.webhook_secret_secret_version
            )
        if ConnectionGitlabConfigReadAuthorizerCredential.to_proto(
            resource.read_authorizer_credential
        ):
            res.read_authorizer_credential.CopyFrom(
                ConnectionGitlabConfigReadAuthorizerCredential.to_proto(
                    resource.read_authorizer_credential
                )
            )
        else:
            res.ClearField("read_authorizer_credential")
        if ConnectionGitlabConfigAuthorizerCredential.to_proto(
            resource.authorizer_credential
        ):
            res.authorizer_credential.CopyFrom(
                ConnectionGitlabConfigAuthorizerCredential.to_proto(
                    resource.authorizer_credential
                )
            )
        else:
            res.ClearField("authorizer_credential")
        if ConnectionGitlabConfigServiceDirectoryConfig.to_proto(
            resource.service_directory_config
        ):
            res.service_directory_config.CopyFrom(
                ConnectionGitlabConfigServiceDirectoryConfig.to_proto(
                    resource.service_directory_config
                )
            )
        else:
            res.ClearField("service_directory_config")
        if Primitive.to_proto(resource.ssl_ca):
            res.ssl_ca = Primitive.to_proto(resource.ssl_ca)
        if Primitive.to_proto(resource.server_version):
            res.server_version = Primitive.to_proto(resource.server_version)
        return res

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return None

        return ConnectionGitlabConfig(
            host_uri=Primitive.from_proto(resource.host_uri),
            webhook_secret_secret_version=Primitive.from_proto(
                resource.webhook_secret_secret_version
            ),
            read_authorizer_credential=ConnectionGitlabConfigReadAuthorizerCredential.from_proto(
                resource.read_authorizer_credential
            ),
            authorizer_credential=ConnectionGitlabConfigAuthorizerCredential.from_proto(
                resource.authorizer_credential
            ),
            service_directory_config=ConnectionGitlabConfigServiceDirectoryConfig.from_proto(
                resource.service_directory_config
            ),
            ssl_ca=Primitive.from_proto(resource.ssl_ca),
            server_version=Primitive.from_proto(resource.server_version),
        )


class ConnectionGitlabConfigArray(object):
    @classmethod
    def to_proto(self, resources):
        if not resources:
            return resources
        return [ConnectionGitlabConfig.to_proto(i) for i in resources]

    @classmethod
    def from_proto(self, resources):
        return [ConnectionGitlabConfig.from_proto(i) for i in resources]


class ConnectionGitlabConfigReadAuthorizerCredential(object):
    def __init__(self, user_token_secret_version: str = None, username: str = None):
        self.user_token_secret_version = user_token_secret_version
        self.username = username

    @classmethod
    def to_proto(self, resource):
        if not resource:
            return None

        res = (
            connection_pb2.Cloudbuildv2ConnectionGitlabConfigReadAuthorizerCredential()
        )
        if Primitive.to_proto(resource.user_token_secret_version):
            res.user_token_secret_version = Primitive.to_proto(
                resource.user_token_secret_version
            )
        if Primitive.to_proto(resource.username):
            res.username = Primitive.to_proto(resource.username)
        return res

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return None

        return ConnectionGitlabConfigReadAuthorizerCredential(
            user_token_secret_version=Primitive.from_proto(
                resource.user_token_secret_version
            ),
            username=Primitive.from_proto(resource.username),
        )


class ConnectionGitlabConfigReadAuthorizerCredentialArray(object):
    @classmethod
    def to_proto(self, resources):
        if not resources:
            return resources
        return [
            ConnectionGitlabConfigReadAuthorizerCredential.to_proto(i)
            for i in resources
        ]

    @classmethod
    def from_proto(self, resources):
        return [
            ConnectionGitlabConfigReadAuthorizerCredential.from_proto(i)
            for i in resources
        ]


class ConnectionGitlabConfigAuthorizerCredential(object):
    def __init__(self, user_token_secret_version: str = None, username: str = None):
        self.user_token_secret_version = user_token_secret_version
        self.username = username

    @classmethod
    def to_proto(self, resource):
        if not resource:
            return None

        res = connection_pb2.Cloudbuildv2ConnectionGitlabConfigAuthorizerCredential()
        if Primitive.to_proto(resource.user_token_secret_version):
            res.user_token_secret_version = Primitive.to_proto(
                resource.user_token_secret_version
            )
        if Primitive.to_proto(resource.username):
            res.username = Primitive.to_proto(resource.username)
        return res

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return None

        return ConnectionGitlabConfigAuthorizerCredential(
            user_token_secret_version=Primitive.from_proto(
                resource.user_token_secret_version
            ),
            username=Primitive.from_proto(resource.username),
        )


class ConnectionGitlabConfigAuthorizerCredentialArray(object):
    @classmethod
    def to_proto(self, resources):
        if not resources:
            return resources
        return [
            ConnectionGitlabConfigAuthorizerCredential.to_proto(i) for i in resources
        ]

    @classmethod
    def from_proto(self, resources):
        return [
            ConnectionGitlabConfigAuthorizerCredential.from_proto(i) for i in resources
        ]


class ConnectionGitlabConfigServiceDirectoryConfig(object):
    def __init__(self, service: str = None):
        self.service = service

    @classmethod
    def to_proto(self, resource):
        if not resource:
            return None

        res = connection_pb2.Cloudbuildv2ConnectionGitlabConfigServiceDirectoryConfig()
        if Primitive.to_proto(resource.service):
            res.service = Primitive.to_proto(resource.service)
        return res

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return None

        return ConnectionGitlabConfigServiceDirectoryConfig(
            service=Primitive.from_proto(resource.service),
        )


class ConnectionGitlabConfigServiceDirectoryConfigArray(object):
    @classmethod
    def to_proto(self, resources):
        if not resources:
            return resources
        return [
            ConnectionGitlabConfigServiceDirectoryConfig.to_proto(i) for i in resources
        ]

    @classmethod
    def from_proto(self, resources):
        return [
            ConnectionGitlabConfigServiceDirectoryConfig.from_proto(i)
            for i in resources
        ]


class ConnectionInstallationState(object):
    def __init__(self, stage: str = None, message: str = None, action_uri: str = None):
        self.stage = stage
        self.message = message
        self.action_uri = action_uri

    @classmethod
    def to_proto(self, resource):
        if not resource:
            return None

        res = connection_pb2.Cloudbuildv2ConnectionInstallationState()
        if ConnectionInstallationStateStageEnum.to_proto(resource.stage):
            res.stage = ConnectionInstallationStateStageEnum.to_proto(resource.stage)
        if Primitive.to_proto(resource.message):
            res.message = Primitive.to_proto(resource.message)
        if Primitive.to_proto(resource.action_uri):
            res.action_uri = Primitive.to_proto(resource.action_uri)
        return res

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return None

        return ConnectionInstallationState(
            stage=ConnectionInstallationStateStageEnum.from_proto(resource.stage),
            message=Primitive.from_proto(resource.message),
            action_uri=Primitive.from_proto(resource.action_uri),
        )


class ConnectionInstallationStateArray(object):
    @classmethod
    def to_proto(self, resources):
        if not resources:
            return resources
        return [ConnectionInstallationState.to_proto(i) for i in resources]

    @classmethod
    def from_proto(self, resources):
        return [ConnectionInstallationState.from_proto(i) for i in resources]


class ConnectionInstallationStateStageEnum(object):
    @classmethod
    def to_proto(self, resource):
        if not resource:
            return resource
        return connection_pb2.Cloudbuildv2ConnectionInstallationStateStageEnum.Value(
            "Cloudbuildv2ConnectionInstallationStateStageEnum%s" % resource
        )

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return resource
        return connection_pb2.Cloudbuildv2ConnectionInstallationStateStageEnum.Name(
            resource
        )[len("Cloudbuildv2ConnectionInstallationStateStageEnum") :]


class Primitive(object):
    @classmethod
    def to_proto(self, s):
        if not s:
            return ""
        return s

    @classmethod
    def from_proto(self, s):
        return s
