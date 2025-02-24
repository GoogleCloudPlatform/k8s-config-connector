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
from google3.cloud.graphite.mmv2.services.google.compute import backend_service_pb2
from google3.cloud.graphite.mmv2.services.google.compute import backend_service_pb2_grpc

from typing import List


class BackendService(object):
    def __init__(
        self,
        name: str = None,
        description: str = None,
        self_link: str = None,
        self_link_with_id: str = None,
        backends: list = None,
        health_checks: list = None,
        timeout_sec: int = None,
        port: int = None,
        protocol: str = None,
        fingerprint: str = None,
        port_name: str = None,
        enable_cdn: bool = None,
        session_affinity: str = None,
        affinity_cookie_ttl_sec: int = None,
        location: str = None,
        failover_policy: dict = None,
        load_balancing_scheme: str = None,
        connection_draining: dict = None,
        iap: dict = None,
        cdn_policy: dict = None,
        custom_request_headers: list = None,
        custom_response_headers: list = None,
        security_policy: str = None,
        log_config: dict = None,
        security_settings: dict = None,
        locality_lb_policy: str = None,
        consistent_hash: dict = None,
        circuit_breakers: dict = None,
        outlier_detection: dict = None,
        network: str = None,
        subsetting: dict = None,
        connection_tracking_policy: dict = None,
        max_stream_duration: dict = None,
        project: str = None,
        service_account_file: str = "",
    ):

        channel.initialize()
        self.name = name
        self.description = description
        self.backends = backends
        self.health_checks = health_checks
        self.timeout_sec = timeout_sec
        self.port = port
        self.protocol = protocol
        self.port_name = port_name
        self.enable_cdn = enable_cdn
        self.session_affinity = session_affinity
        self.affinity_cookie_ttl_sec = affinity_cookie_ttl_sec
        self.location = location
        self.failover_policy = failover_policy
        self.load_balancing_scheme = load_balancing_scheme
        self.connection_draining = connection_draining
        self.iap = iap
        self.cdn_policy = cdn_policy
        self.custom_request_headers = custom_request_headers
        self.custom_response_headers = custom_response_headers
        self.security_policy = security_policy
        self.log_config = log_config
        self.security_settings = security_settings
        self.locality_lb_policy = locality_lb_policy
        self.consistent_hash = consistent_hash
        self.circuit_breakers = circuit_breakers
        self.outlier_detection = outlier_detection
        self.network = network
        self.subsetting = subsetting
        self.connection_tracking_policy = connection_tracking_policy
        self.max_stream_duration = max_stream_duration
        self.project = project
        self.service_account_file = service_account_file

    def apply(self):
        stub = backend_service_pb2_grpc.ComputeBetaBackendServiceServiceStub(
            channel.Channel()
        )
        request = backend_service_pb2.ApplyComputeBetaBackendServiceRequest()
        if Primitive.to_proto(self.name):
            request.resource.name = Primitive.to_proto(self.name)

        if Primitive.to_proto(self.description):
            request.resource.description = Primitive.to_proto(self.description)

        if BackendServiceBackendsArray.to_proto(self.backends):
            request.resource.backends.extend(
                BackendServiceBackendsArray.to_proto(self.backends)
            )
        if Primitive.to_proto(self.health_checks):
            request.resource.health_checks.extend(
                Primitive.to_proto(self.health_checks)
            )
        if Primitive.to_proto(self.timeout_sec):
            request.resource.timeout_sec = Primitive.to_proto(self.timeout_sec)

        if Primitive.to_proto(self.port):
            request.resource.port = Primitive.to_proto(self.port)

        if BackendServiceProtocolEnum.to_proto(self.protocol):
            request.resource.protocol = BackendServiceProtocolEnum.to_proto(
                self.protocol
            )

        if Primitive.to_proto(self.port_name):
            request.resource.port_name = Primitive.to_proto(self.port_name)

        if Primitive.to_proto(self.enable_cdn):
            request.resource.enable_cdn = Primitive.to_proto(self.enable_cdn)

        if BackendServiceSessionAffinityEnum.to_proto(self.session_affinity):
            request.resource.session_affinity = BackendServiceSessionAffinityEnum.to_proto(
                self.session_affinity
            )

        if Primitive.to_proto(self.affinity_cookie_ttl_sec):
            request.resource.affinity_cookie_ttl_sec = Primitive.to_proto(
                self.affinity_cookie_ttl_sec
            )

        if Primitive.to_proto(self.location):
            request.resource.location = Primitive.to_proto(self.location)

        if BackendServiceFailoverPolicy.to_proto(self.failover_policy):
            request.resource.failover_policy.CopyFrom(
                BackendServiceFailoverPolicy.to_proto(self.failover_policy)
            )
        else:
            request.resource.ClearField("failover_policy")
        if BackendServiceLoadBalancingSchemeEnum.to_proto(self.load_balancing_scheme):
            request.resource.load_balancing_scheme = BackendServiceLoadBalancingSchemeEnum.to_proto(
                self.load_balancing_scheme
            )

        if BackendServiceConnectionDraining.to_proto(self.connection_draining):
            request.resource.connection_draining.CopyFrom(
                BackendServiceConnectionDraining.to_proto(self.connection_draining)
            )
        else:
            request.resource.ClearField("connection_draining")
        if BackendServiceIap.to_proto(self.iap):
            request.resource.iap.CopyFrom(BackendServiceIap.to_proto(self.iap))
        else:
            request.resource.ClearField("iap")
        if BackendServiceCdnPolicy.to_proto(self.cdn_policy):
            request.resource.cdn_policy.CopyFrom(
                BackendServiceCdnPolicy.to_proto(self.cdn_policy)
            )
        else:
            request.resource.ClearField("cdn_policy")
        if Primitive.to_proto(self.custom_request_headers):
            request.resource.custom_request_headers.extend(
                Primitive.to_proto(self.custom_request_headers)
            )
        if Primitive.to_proto(self.custom_response_headers):
            request.resource.custom_response_headers.extend(
                Primitive.to_proto(self.custom_response_headers)
            )
        if Primitive.to_proto(self.security_policy):
            request.resource.security_policy = Primitive.to_proto(self.security_policy)

        if BackendServiceLogConfig.to_proto(self.log_config):
            request.resource.log_config.CopyFrom(
                BackendServiceLogConfig.to_proto(self.log_config)
            )
        else:
            request.resource.ClearField("log_config")
        if BackendServiceSecuritySettings.to_proto(self.security_settings):
            request.resource.security_settings.CopyFrom(
                BackendServiceSecuritySettings.to_proto(self.security_settings)
            )
        else:
            request.resource.ClearField("security_settings")
        if BackendServiceLocalityLbPolicyEnum.to_proto(self.locality_lb_policy):
            request.resource.locality_lb_policy = BackendServiceLocalityLbPolicyEnum.to_proto(
                self.locality_lb_policy
            )

        if BackendServiceConsistentHash.to_proto(self.consistent_hash):
            request.resource.consistent_hash.CopyFrom(
                BackendServiceConsistentHash.to_proto(self.consistent_hash)
            )
        else:
            request.resource.ClearField("consistent_hash")
        if BackendServiceCircuitBreakers.to_proto(self.circuit_breakers):
            request.resource.circuit_breakers.CopyFrom(
                BackendServiceCircuitBreakers.to_proto(self.circuit_breakers)
            )
        else:
            request.resource.ClearField("circuit_breakers")
        if BackendServiceOutlierDetection.to_proto(self.outlier_detection):
            request.resource.outlier_detection.CopyFrom(
                BackendServiceOutlierDetection.to_proto(self.outlier_detection)
            )
        else:
            request.resource.ClearField("outlier_detection")
        if Primitive.to_proto(self.network):
            request.resource.network = Primitive.to_proto(self.network)

        if BackendServiceSubsetting.to_proto(self.subsetting):
            request.resource.subsetting.CopyFrom(
                BackendServiceSubsetting.to_proto(self.subsetting)
            )
        else:
            request.resource.ClearField("subsetting")
        if BackendServiceConnectionTrackingPolicy.to_proto(
            self.connection_tracking_policy
        ):
            request.resource.connection_tracking_policy.CopyFrom(
                BackendServiceConnectionTrackingPolicy.to_proto(
                    self.connection_tracking_policy
                )
            )
        else:
            request.resource.ClearField("connection_tracking_policy")
        if BackendServiceMaxStreamDuration.to_proto(self.max_stream_duration):
            request.resource.max_stream_duration.CopyFrom(
                BackendServiceMaxStreamDuration.to_proto(self.max_stream_duration)
            )
        else:
            request.resource.ClearField("max_stream_duration")
        if Primitive.to_proto(self.project):
            request.resource.project = Primitive.to_proto(self.project)

        request.service_account_file = self.service_account_file

        response = stub.ApplyComputeBetaBackendService(request)
        self.name = Primitive.from_proto(response.name)
        self.description = Primitive.from_proto(response.description)
        self.self_link = Primitive.from_proto(response.self_link)
        self.self_link_with_id = Primitive.from_proto(response.self_link_with_id)
        self.backends = BackendServiceBackendsArray.from_proto(response.backends)
        self.health_checks = Primitive.from_proto(response.health_checks)
        self.timeout_sec = Primitive.from_proto(response.timeout_sec)
        self.port = Primitive.from_proto(response.port)
        self.protocol = BackendServiceProtocolEnum.from_proto(response.protocol)
        self.fingerprint = Primitive.from_proto(response.fingerprint)
        self.port_name = Primitive.from_proto(response.port_name)
        self.enable_cdn = Primitive.from_proto(response.enable_cdn)
        self.session_affinity = BackendServiceSessionAffinityEnum.from_proto(
            response.session_affinity
        )
        self.affinity_cookie_ttl_sec = Primitive.from_proto(
            response.affinity_cookie_ttl_sec
        )
        self.location = Primitive.from_proto(response.location)
        self.failover_policy = BackendServiceFailoverPolicy.from_proto(
            response.failover_policy
        )
        self.load_balancing_scheme = BackendServiceLoadBalancingSchemeEnum.from_proto(
            response.load_balancing_scheme
        )
        self.connection_draining = BackendServiceConnectionDraining.from_proto(
            response.connection_draining
        )
        self.iap = BackendServiceIap.from_proto(response.iap)
        self.cdn_policy = BackendServiceCdnPolicy.from_proto(response.cdn_policy)
        self.custom_request_headers = Primitive.from_proto(
            response.custom_request_headers
        )
        self.custom_response_headers = Primitive.from_proto(
            response.custom_response_headers
        )
        self.security_policy = Primitive.from_proto(response.security_policy)
        self.log_config = BackendServiceLogConfig.from_proto(response.log_config)
        self.security_settings = BackendServiceSecuritySettings.from_proto(
            response.security_settings
        )
        self.locality_lb_policy = BackendServiceLocalityLbPolicyEnum.from_proto(
            response.locality_lb_policy
        )
        self.consistent_hash = BackendServiceConsistentHash.from_proto(
            response.consistent_hash
        )
        self.circuit_breakers = BackendServiceCircuitBreakers.from_proto(
            response.circuit_breakers
        )
        self.outlier_detection = BackendServiceOutlierDetection.from_proto(
            response.outlier_detection
        )
        self.network = Primitive.from_proto(response.network)
        self.subsetting = BackendServiceSubsetting.from_proto(response.subsetting)
        self.connection_tracking_policy = BackendServiceConnectionTrackingPolicy.from_proto(
            response.connection_tracking_policy
        )
        self.max_stream_duration = BackendServiceMaxStreamDuration.from_proto(
            response.max_stream_duration
        )
        self.project = Primitive.from_proto(response.project)

    def delete(self):
        stub = backend_service_pb2_grpc.ComputeBetaBackendServiceServiceStub(
            channel.Channel()
        )
        request = backend_service_pb2.DeleteComputeBetaBackendServiceRequest()
        request.service_account_file = self.service_account_file
        if Primitive.to_proto(self.name):
            request.resource.name = Primitive.to_proto(self.name)

        if Primitive.to_proto(self.description):
            request.resource.description = Primitive.to_proto(self.description)

        if BackendServiceBackendsArray.to_proto(self.backends):
            request.resource.backends.extend(
                BackendServiceBackendsArray.to_proto(self.backends)
            )
        if Primitive.to_proto(self.health_checks):
            request.resource.health_checks.extend(
                Primitive.to_proto(self.health_checks)
            )
        if Primitive.to_proto(self.timeout_sec):
            request.resource.timeout_sec = Primitive.to_proto(self.timeout_sec)

        if Primitive.to_proto(self.port):
            request.resource.port = Primitive.to_proto(self.port)

        if BackendServiceProtocolEnum.to_proto(self.protocol):
            request.resource.protocol = BackendServiceProtocolEnum.to_proto(
                self.protocol
            )

        if Primitive.to_proto(self.port_name):
            request.resource.port_name = Primitive.to_proto(self.port_name)

        if Primitive.to_proto(self.enable_cdn):
            request.resource.enable_cdn = Primitive.to_proto(self.enable_cdn)

        if BackendServiceSessionAffinityEnum.to_proto(self.session_affinity):
            request.resource.session_affinity = BackendServiceSessionAffinityEnum.to_proto(
                self.session_affinity
            )

        if Primitive.to_proto(self.affinity_cookie_ttl_sec):
            request.resource.affinity_cookie_ttl_sec = Primitive.to_proto(
                self.affinity_cookie_ttl_sec
            )

        if Primitive.to_proto(self.location):
            request.resource.location = Primitive.to_proto(self.location)

        if BackendServiceFailoverPolicy.to_proto(self.failover_policy):
            request.resource.failover_policy.CopyFrom(
                BackendServiceFailoverPolicy.to_proto(self.failover_policy)
            )
        else:
            request.resource.ClearField("failover_policy")
        if BackendServiceLoadBalancingSchemeEnum.to_proto(self.load_balancing_scheme):
            request.resource.load_balancing_scheme = BackendServiceLoadBalancingSchemeEnum.to_proto(
                self.load_balancing_scheme
            )

        if BackendServiceConnectionDraining.to_proto(self.connection_draining):
            request.resource.connection_draining.CopyFrom(
                BackendServiceConnectionDraining.to_proto(self.connection_draining)
            )
        else:
            request.resource.ClearField("connection_draining")
        if BackendServiceIap.to_proto(self.iap):
            request.resource.iap.CopyFrom(BackendServiceIap.to_proto(self.iap))
        else:
            request.resource.ClearField("iap")
        if BackendServiceCdnPolicy.to_proto(self.cdn_policy):
            request.resource.cdn_policy.CopyFrom(
                BackendServiceCdnPolicy.to_proto(self.cdn_policy)
            )
        else:
            request.resource.ClearField("cdn_policy")
        if Primitive.to_proto(self.custom_request_headers):
            request.resource.custom_request_headers.extend(
                Primitive.to_proto(self.custom_request_headers)
            )
        if Primitive.to_proto(self.custom_response_headers):
            request.resource.custom_response_headers.extend(
                Primitive.to_proto(self.custom_response_headers)
            )
        if Primitive.to_proto(self.security_policy):
            request.resource.security_policy = Primitive.to_proto(self.security_policy)

        if BackendServiceLogConfig.to_proto(self.log_config):
            request.resource.log_config.CopyFrom(
                BackendServiceLogConfig.to_proto(self.log_config)
            )
        else:
            request.resource.ClearField("log_config")
        if BackendServiceSecuritySettings.to_proto(self.security_settings):
            request.resource.security_settings.CopyFrom(
                BackendServiceSecuritySettings.to_proto(self.security_settings)
            )
        else:
            request.resource.ClearField("security_settings")
        if BackendServiceLocalityLbPolicyEnum.to_proto(self.locality_lb_policy):
            request.resource.locality_lb_policy = BackendServiceLocalityLbPolicyEnum.to_proto(
                self.locality_lb_policy
            )

        if BackendServiceConsistentHash.to_proto(self.consistent_hash):
            request.resource.consistent_hash.CopyFrom(
                BackendServiceConsistentHash.to_proto(self.consistent_hash)
            )
        else:
            request.resource.ClearField("consistent_hash")
        if BackendServiceCircuitBreakers.to_proto(self.circuit_breakers):
            request.resource.circuit_breakers.CopyFrom(
                BackendServiceCircuitBreakers.to_proto(self.circuit_breakers)
            )
        else:
            request.resource.ClearField("circuit_breakers")
        if BackendServiceOutlierDetection.to_proto(self.outlier_detection):
            request.resource.outlier_detection.CopyFrom(
                BackendServiceOutlierDetection.to_proto(self.outlier_detection)
            )
        else:
            request.resource.ClearField("outlier_detection")
        if Primitive.to_proto(self.network):
            request.resource.network = Primitive.to_proto(self.network)

        if BackendServiceSubsetting.to_proto(self.subsetting):
            request.resource.subsetting.CopyFrom(
                BackendServiceSubsetting.to_proto(self.subsetting)
            )
        else:
            request.resource.ClearField("subsetting")
        if BackendServiceConnectionTrackingPolicy.to_proto(
            self.connection_tracking_policy
        ):
            request.resource.connection_tracking_policy.CopyFrom(
                BackendServiceConnectionTrackingPolicy.to_proto(
                    self.connection_tracking_policy
                )
            )
        else:
            request.resource.ClearField("connection_tracking_policy")
        if BackendServiceMaxStreamDuration.to_proto(self.max_stream_duration):
            request.resource.max_stream_duration.CopyFrom(
                BackendServiceMaxStreamDuration.to_proto(self.max_stream_duration)
            )
        else:
            request.resource.ClearField("max_stream_duration")
        if Primitive.to_proto(self.project):
            request.resource.project = Primitive.to_proto(self.project)

        response = stub.DeleteComputeBetaBackendService(request)

    @classmethod
    def list(self, project, location, service_account_file=""):
        stub = backend_service_pb2_grpc.ComputeBetaBackendServiceServiceStub(
            channel.Channel()
        )
        request = backend_service_pb2.ListComputeBetaBackendServiceRequest()
        request.service_account_file = service_account_file
        request.Project = project

        request.Location = location

        return stub.ListComputeBetaBackendService(request).items

    def to_proto(self):
        resource = backend_service_pb2.ComputeBetaBackendService()
        if Primitive.to_proto(self.name):
            resource.name = Primitive.to_proto(self.name)
        if Primitive.to_proto(self.description):
            resource.description = Primitive.to_proto(self.description)
        if BackendServiceBackendsArray.to_proto(self.backends):
            resource.backends.extend(
                BackendServiceBackendsArray.to_proto(self.backends)
            )
        if Primitive.to_proto(self.health_checks):
            resource.health_checks.extend(Primitive.to_proto(self.health_checks))
        if Primitive.to_proto(self.timeout_sec):
            resource.timeout_sec = Primitive.to_proto(self.timeout_sec)
        if Primitive.to_proto(self.port):
            resource.port = Primitive.to_proto(self.port)
        if BackendServiceProtocolEnum.to_proto(self.protocol):
            resource.protocol = BackendServiceProtocolEnum.to_proto(self.protocol)
        if Primitive.to_proto(self.port_name):
            resource.port_name = Primitive.to_proto(self.port_name)
        if Primitive.to_proto(self.enable_cdn):
            resource.enable_cdn = Primitive.to_proto(self.enable_cdn)
        if BackendServiceSessionAffinityEnum.to_proto(self.session_affinity):
            resource.session_affinity = BackendServiceSessionAffinityEnum.to_proto(
                self.session_affinity
            )
        if Primitive.to_proto(self.affinity_cookie_ttl_sec):
            resource.affinity_cookie_ttl_sec = Primitive.to_proto(
                self.affinity_cookie_ttl_sec
            )
        if Primitive.to_proto(self.location):
            resource.location = Primitive.to_proto(self.location)
        if BackendServiceFailoverPolicy.to_proto(self.failover_policy):
            resource.failover_policy.CopyFrom(
                BackendServiceFailoverPolicy.to_proto(self.failover_policy)
            )
        else:
            resource.ClearField("failover_policy")
        if BackendServiceLoadBalancingSchemeEnum.to_proto(self.load_balancing_scheme):
            resource.load_balancing_scheme = BackendServiceLoadBalancingSchemeEnum.to_proto(
                self.load_balancing_scheme
            )
        if BackendServiceConnectionDraining.to_proto(self.connection_draining):
            resource.connection_draining.CopyFrom(
                BackendServiceConnectionDraining.to_proto(self.connection_draining)
            )
        else:
            resource.ClearField("connection_draining")
        if BackendServiceIap.to_proto(self.iap):
            resource.iap.CopyFrom(BackendServiceIap.to_proto(self.iap))
        else:
            resource.ClearField("iap")
        if BackendServiceCdnPolicy.to_proto(self.cdn_policy):
            resource.cdn_policy.CopyFrom(
                BackendServiceCdnPolicy.to_proto(self.cdn_policy)
            )
        else:
            resource.ClearField("cdn_policy")
        if Primitive.to_proto(self.custom_request_headers):
            resource.custom_request_headers.extend(
                Primitive.to_proto(self.custom_request_headers)
            )
        if Primitive.to_proto(self.custom_response_headers):
            resource.custom_response_headers.extend(
                Primitive.to_proto(self.custom_response_headers)
            )
        if Primitive.to_proto(self.security_policy):
            resource.security_policy = Primitive.to_proto(self.security_policy)
        if BackendServiceLogConfig.to_proto(self.log_config):
            resource.log_config.CopyFrom(
                BackendServiceLogConfig.to_proto(self.log_config)
            )
        else:
            resource.ClearField("log_config")
        if BackendServiceSecuritySettings.to_proto(self.security_settings):
            resource.security_settings.CopyFrom(
                BackendServiceSecuritySettings.to_proto(self.security_settings)
            )
        else:
            resource.ClearField("security_settings")
        if BackendServiceLocalityLbPolicyEnum.to_proto(self.locality_lb_policy):
            resource.locality_lb_policy = BackendServiceLocalityLbPolicyEnum.to_proto(
                self.locality_lb_policy
            )
        if BackendServiceConsistentHash.to_proto(self.consistent_hash):
            resource.consistent_hash.CopyFrom(
                BackendServiceConsistentHash.to_proto(self.consistent_hash)
            )
        else:
            resource.ClearField("consistent_hash")
        if BackendServiceCircuitBreakers.to_proto(self.circuit_breakers):
            resource.circuit_breakers.CopyFrom(
                BackendServiceCircuitBreakers.to_proto(self.circuit_breakers)
            )
        else:
            resource.ClearField("circuit_breakers")
        if BackendServiceOutlierDetection.to_proto(self.outlier_detection):
            resource.outlier_detection.CopyFrom(
                BackendServiceOutlierDetection.to_proto(self.outlier_detection)
            )
        else:
            resource.ClearField("outlier_detection")
        if Primitive.to_proto(self.network):
            resource.network = Primitive.to_proto(self.network)
        if BackendServiceSubsetting.to_proto(self.subsetting):
            resource.subsetting.CopyFrom(
                BackendServiceSubsetting.to_proto(self.subsetting)
            )
        else:
            resource.ClearField("subsetting")
        if BackendServiceConnectionTrackingPolicy.to_proto(
            self.connection_tracking_policy
        ):
            resource.connection_tracking_policy.CopyFrom(
                BackendServiceConnectionTrackingPolicy.to_proto(
                    self.connection_tracking_policy
                )
            )
        else:
            resource.ClearField("connection_tracking_policy")
        if BackendServiceMaxStreamDuration.to_proto(self.max_stream_duration):
            resource.max_stream_duration.CopyFrom(
                BackendServiceMaxStreamDuration.to_proto(self.max_stream_duration)
            )
        else:
            resource.ClearField("max_stream_duration")
        if Primitive.to_proto(self.project):
            resource.project = Primitive.to_proto(self.project)
        return resource


class BackendServiceBackends(object):
    def __init__(
        self,
        description: str = None,
        group: str = None,
        balancing_mode: str = None,
        max_utilization: float = None,
        max_rate: int = None,
        max_rate_per_instance: float = None,
        max_rate_per_endpoint: float = None,
        max_connections: int = None,
        max_connections_per_instance: int = None,
        max_connections_per_endpoint: int = None,
        capacity_scaler: float = None,
        failover: bool = None,
    ):
        self.description = description
        self.group = group
        self.balancing_mode = balancing_mode
        self.max_utilization = max_utilization
        self.max_rate = max_rate
        self.max_rate_per_instance = max_rate_per_instance
        self.max_rate_per_endpoint = max_rate_per_endpoint
        self.max_connections = max_connections
        self.max_connections_per_instance = max_connections_per_instance
        self.max_connections_per_endpoint = max_connections_per_endpoint
        self.capacity_scaler = capacity_scaler
        self.failover = failover

    @classmethod
    def to_proto(self, resource):
        if not resource:
            return None

        res = backend_service_pb2.ComputeBetaBackendServiceBackends()
        if Primitive.to_proto(resource.description):
            res.description = Primitive.to_proto(resource.description)
        if Primitive.to_proto(resource.group):
            res.group = Primitive.to_proto(resource.group)
        if BackendServiceBackendsBalancingModeEnum.to_proto(resource.balancing_mode):
            res.balancing_mode = BackendServiceBackendsBalancingModeEnum.to_proto(
                resource.balancing_mode
            )
        if Primitive.to_proto(resource.max_utilization):
            res.max_utilization = Primitive.to_proto(resource.max_utilization)
        if Primitive.to_proto(resource.max_rate):
            res.max_rate = Primitive.to_proto(resource.max_rate)
        if Primitive.to_proto(resource.max_rate_per_instance):
            res.max_rate_per_instance = Primitive.to_proto(
                resource.max_rate_per_instance
            )
        if Primitive.to_proto(resource.max_rate_per_endpoint):
            res.max_rate_per_endpoint = Primitive.to_proto(
                resource.max_rate_per_endpoint
            )
        if Primitive.to_proto(resource.max_connections):
            res.max_connections = Primitive.to_proto(resource.max_connections)
        if Primitive.to_proto(resource.max_connections_per_instance):
            res.max_connections_per_instance = Primitive.to_proto(
                resource.max_connections_per_instance
            )
        if Primitive.to_proto(resource.max_connections_per_endpoint):
            res.max_connections_per_endpoint = Primitive.to_proto(
                resource.max_connections_per_endpoint
            )
        if Primitive.to_proto(resource.capacity_scaler):
            res.capacity_scaler = Primitive.to_proto(resource.capacity_scaler)
        if Primitive.to_proto(resource.failover):
            res.failover = Primitive.to_proto(resource.failover)
        return res

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return None

        return BackendServiceBackends(
            description=Primitive.from_proto(resource.description),
            group=Primitive.from_proto(resource.group),
            balancing_mode=BackendServiceBackendsBalancingModeEnum.from_proto(
                resource.balancing_mode
            ),
            max_utilization=Primitive.from_proto(resource.max_utilization),
            max_rate=Primitive.from_proto(resource.max_rate),
            max_rate_per_instance=Primitive.from_proto(resource.max_rate_per_instance),
            max_rate_per_endpoint=Primitive.from_proto(resource.max_rate_per_endpoint),
            max_connections=Primitive.from_proto(resource.max_connections),
            max_connections_per_instance=Primitive.from_proto(
                resource.max_connections_per_instance
            ),
            max_connections_per_endpoint=Primitive.from_proto(
                resource.max_connections_per_endpoint
            ),
            capacity_scaler=Primitive.from_proto(resource.capacity_scaler),
            failover=Primitive.from_proto(resource.failover),
        )


class BackendServiceBackendsArray(object):
    @classmethod
    def to_proto(self, resources):
        if not resources:
            return resources
        return [BackendServiceBackends.to_proto(i) for i in resources]

    @classmethod
    def from_proto(self, resources):
        return [BackendServiceBackends.from_proto(i) for i in resources]


class BackendServiceFailoverPolicy(object):
    def __init__(
        self,
        disable_connection_drain_on_failover: bool = None,
        drop_traffic_if_unhealthy: bool = None,
        failover_ratio: float = None,
    ):
        self.disable_connection_drain_on_failover = disable_connection_drain_on_failover
        self.drop_traffic_if_unhealthy = drop_traffic_if_unhealthy
        self.failover_ratio = failover_ratio

    @classmethod
    def to_proto(self, resource):
        if not resource:
            return None

        res = backend_service_pb2.ComputeBetaBackendServiceFailoverPolicy()
        if Primitive.to_proto(resource.disable_connection_drain_on_failover):
            res.disable_connection_drain_on_failover = Primitive.to_proto(
                resource.disable_connection_drain_on_failover
            )
        if Primitive.to_proto(resource.drop_traffic_if_unhealthy):
            res.drop_traffic_if_unhealthy = Primitive.to_proto(
                resource.drop_traffic_if_unhealthy
            )
        if Primitive.to_proto(resource.failover_ratio):
            res.failover_ratio = Primitive.to_proto(resource.failover_ratio)
        return res

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return None

        return BackendServiceFailoverPolicy(
            disable_connection_drain_on_failover=Primitive.from_proto(
                resource.disable_connection_drain_on_failover
            ),
            drop_traffic_if_unhealthy=Primitive.from_proto(
                resource.drop_traffic_if_unhealthy
            ),
            failover_ratio=Primitive.from_proto(resource.failover_ratio),
        )


class BackendServiceFailoverPolicyArray(object):
    @classmethod
    def to_proto(self, resources):
        if not resources:
            return resources
        return [BackendServiceFailoverPolicy.to_proto(i) for i in resources]

    @classmethod
    def from_proto(self, resources):
        return [BackendServiceFailoverPolicy.from_proto(i) for i in resources]


class BackendServiceConnectionDraining(object):
    def __init__(self, draining_timeout_sec: int = None):
        self.draining_timeout_sec = draining_timeout_sec

    @classmethod
    def to_proto(self, resource):
        if not resource:
            return None

        res = backend_service_pb2.ComputeBetaBackendServiceConnectionDraining()
        if Primitive.to_proto(resource.draining_timeout_sec):
            res.draining_timeout_sec = Primitive.to_proto(resource.draining_timeout_sec)
        return res

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return None

        return BackendServiceConnectionDraining(
            draining_timeout_sec=Primitive.from_proto(resource.draining_timeout_sec),
        )


class BackendServiceConnectionDrainingArray(object):
    @classmethod
    def to_proto(self, resources):
        if not resources:
            return resources
        return [BackendServiceConnectionDraining.to_proto(i) for i in resources]

    @classmethod
    def from_proto(self, resources):
        return [BackendServiceConnectionDraining.from_proto(i) for i in resources]


class BackendServiceIap(object):
    def __init__(
        self,
        enabled: bool = None,
        oauth2_client_id: str = None,
        oauth2_client_secret: str = None,
        oauth2_client_secret_sha256: str = None,
    ):
        self.enabled = enabled
        self.oauth2_client_id = oauth2_client_id
        self.oauth2_client_secret = oauth2_client_secret
        self.oauth2_client_secret_sha256 = oauth2_client_secret_sha256

    @classmethod
    def to_proto(self, resource):
        if not resource:
            return None

        res = backend_service_pb2.ComputeBetaBackendServiceIap()
        if Primitive.to_proto(resource.enabled):
            res.enabled = Primitive.to_proto(resource.enabled)
        if Primitive.to_proto(resource.oauth2_client_id):
            res.oauth2_client_id = Primitive.to_proto(resource.oauth2_client_id)
        if Primitive.to_proto(resource.oauth2_client_secret):
            res.oauth2_client_secret = Primitive.to_proto(resource.oauth2_client_secret)
        if Primitive.to_proto(resource.oauth2_client_secret_sha256):
            res.oauth2_client_secret_sha256 = Primitive.to_proto(
                resource.oauth2_client_secret_sha256
            )
        return res

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return None

        return BackendServiceIap(
            enabled=Primitive.from_proto(resource.enabled),
            oauth2_client_id=Primitive.from_proto(resource.oauth2_client_id),
            oauth2_client_secret=Primitive.from_proto(resource.oauth2_client_secret),
            oauth2_client_secret_sha256=Primitive.from_proto(
                resource.oauth2_client_secret_sha256
            ),
        )


class BackendServiceIapArray(object):
    @classmethod
    def to_proto(self, resources):
        if not resources:
            return resources
        return [BackendServiceIap.to_proto(i) for i in resources]

    @classmethod
    def from_proto(self, resources):
        return [BackendServiceIap.from_proto(i) for i in resources]


class BackendServiceCdnPolicy(object):
    def __init__(
        self,
        cache_key_policy: dict = None,
        signed_url_key_names: list = None,
        signed_url_cache_max_age_sec: int = None,
        request_coalescing: bool = None,
        cache_mode: str = None,
        default_ttl: int = None,
        max_ttl: int = None,
        client_ttl: int = None,
        negative_caching: bool = None,
        negative_caching_policy: list = None,
        bypass_cache_on_request_headers: list = None,
        serve_while_stale: int = None,
    ):
        self.cache_key_policy = cache_key_policy
        self.signed_url_key_names = signed_url_key_names
        self.signed_url_cache_max_age_sec = signed_url_cache_max_age_sec
        self.request_coalescing = request_coalescing
        self.cache_mode = cache_mode
        self.default_ttl = default_ttl
        self.max_ttl = max_ttl
        self.client_ttl = client_ttl
        self.negative_caching = negative_caching
        self.negative_caching_policy = negative_caching_policy
        self.bypass_cache_on_request_headers = bypass_cache_on_request_headers
        self.serve_while_stale = serve_while_stale

    @classmethod
    def to_proto(self, resource):
        if not resource:
            return None

        res = backend_service_pb2.ComputeBetaBackendServiceCdnPolicy()
        if BackendServiceCdnPolicyCacheKeyPolicy.to_proto(resource.cache_key_policy):
            res.cache_key_policy.CopyFrom(
                BackendServiceCdnPolicyCacheKeyPolicy.to_proto(
                    resource.cache_key_policy
                )
            )
        else:
            res.ClearField("cache_key_policy")
        if Primitive.to_proto(resource.signed_url_key_names):
            res.signed_url_key_names.extend(
                Primitive.to_proto(resource.signed_url_key_names)
            )
        if Primitive.to_proto(resource.signed_url_cache_max_age_sec):
            res.signed_url_cache_max_age_sec = Primitive.to_proto(
                resource.signed_url_cache_max_age_sec
            )
        if Primitive.to_proto(resource.request_coalescing):
            res.request_coalescing = Primitive.to_proto(resource.request_coalescing)
        if BackendServiceCdnPolicyCacheModeEnum.to_proto(resource.cache_mode):
            res.cache_mode = BackendServiceCdnPolicyCacheModeEnum.to_proto(
                resource.cache_mode
            )
        if Primitive.to_proto(resource.default_ttl):
            res.default_ttl = Primitive.to_proto(resource.default_ttl)
        if Primitive.to_proto(resource.max_ttl):
            res.max_ttl = Primitive.to_proto(resource.max_ttl)
        if Primitive.to_proto(resource.client_ttl):
            res.client_ttl = Primitive.to_proto(resource.client_ttl)
        if Primitive.to_proto(resource.negative_caching):
            res.negative_caching = Primitive.to_proto(resource.negative_caching)
        if BackendServiceCdnPolicyNegativeCachingPolicyArray.to_proto(
            resource.negative_caching_policy
        ):
            res.negative_caching_policy.extend(
                BackendServiceCdnPolicyNegativeCachingPolicyArray.to_proto(
                    resource.negative_caching_policy
                )
            )
        if BackendServiceCdnPolicyBypassCacheOnRequestHeadersArray.to_proto(
            resource.bypass_cache_on_request_headers
        ):
            res.bypass_cache_on_request_headers.extend(
                BackendServiceCdnPolicyBypassCacheOnRequestHeadersArray.to_proto(
                    resource.bypass_cache_on_request_headers
                )
            )
        if Primitive.to_proto(resource.serve_while_stale):
            res.serve_while_stale = Primitive.to_proto(resource.serve_while_stale)
        return res

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return None

        return BackendServiceCdnPolicy(
            cache_key_policy=BackendServiceCdnPolicyCacheKeyPolicy.from_proto(
                resource.cache_key_policy
            ),
            signed_url_key_names=Primitive.from_proto(resource.signed_url_key_names),
            signed_url_cache_max_age_sec=Primitive.from_proto(
                resource.signed_url_cache_max_age_sec
            ),
            request_coalescing=Primitive.from_proto(resource.request_coalescing),
            cache_mode=BackendServiceCdnPolicyCacheModeEnum.from_proto(
                resource.cache_mode
            ),
            default_ttl=Primitive.from_proto(resource.default_ttl),
            max_ttl=Primitive.from_proto(resource.max_ttl),
            client_ttl=Primitive.from_proto(resource.client_ttl),
            negative_caching=Primitive.from_proto(resource.negative_caching),
            negative_caching_policy=BackendServiceCdnPolicyNegativeCachingPolicyArray.from_proto(
                resource.negative_caching_policy
            ),
            bypass_cache_on_request_headers=BackendServiceCdnPolicyBypassCacheOnRequestHeadersArray.from_proto(
                resource.bypass_cache_on_request_headers
            ),
            serve_while_stale=Primitive.from_proto(resource.serve_while_stale),
        )


class BackendServiceCdnPolicyArray(object):
    @classmethod
    def to_proto(self, resources):
        if not resources:
            return resources
        return [BackendServiceCdnPolicy.to_proto(i) for i in resources]

    @classmethod
    def from_proto(self, resources):
        return [BackendServiceCdnPolicy.from_proto(i) for i in resources]


class BackendServiceCdnPolicyCacheKeyPolicy(object):
    def __init__(
        self,
        include_protocol: bool = None,
        include_host: bool = None,
        include_query_string: bool = None,
        query_string_whitelist: list = None,
        query_string_blacklist: list = None,
        include_http_headers: list = None,
        include_named_cookies: list = None,
    ):
        self.include_protocol = include_protocol
        self.include_host = include_host
        self.include_query_string = include_query_string
        self.query_string_whitelist = query_string_whitelist
        self.query_string_blacklist = query_string_blacklist
        self.include_http_headers = include_http_headers
        self.include_named_cookies = include_named_cookies

    @classmethod
    def to_proto(self, resource):
        if not resource:
            return None

        res = backend_service_pb2.ComputeBetaBackendServiceCdnPolicyCacheKeyPolicy()
        if Primitive.to_proto(resource.include_protocol):
            res.include_protocol = Primitive.to_proto(resource.include_protocol)
        if Primitive.to_proto(resource.include_host):
            res.include_host = Primitive.to_proto(resource.include_host)
        if Primitive.to_proto(resource.include_query_string):
            res.include_query_string = Primitive.to_proto(resource.include_query_string)
        if Primitive.to_proto(resource.query_string_whitelist):
            res.query_string_whitelist.extend(
                Primitive.to_proto(resource.query_string_whitelist)
            )
        if Primitive.to_proto(resource.query_string_blacklist):
            res.query_string_blacklist.extend(
                Primitive.to_proto(resource.query_string_blacklist)
            )
        if Primitive.to_proto(resource.include_http_headers):
            res.include_http_headers.extend(
                Primitive.to_proto(resource.include_http_headers)
            )
        if Primitive.to_proto(resource.include_named_cookies):
            res.include_named_cookies.extend(
                Primitive.to_proto(resource.include_named_cookies)
            )
        return res

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return None

        return BackendServiceCdnPolicyCacheKeyPolicy(
            include_protocol=Primitive.from_proto(resource.include_protocol),
            include_host=Primitive.from_proto(resource.include_host),
            include_query_string=Primitive.from_proto(resource.include_query_string),
            query_string_whitelist=Primitive.from_proto(
                resource.query_string_whitelist
            ),
            query_string_blacklist=Primitive.from_proto(
                resource.query_string_blacklist
            ),
            include_http_headers=Primitive.from_proto(resource.include_http_headers),
            include_named_cookies=Primitive.from_proto(resource.include_named_cookies),
        )


class BackendServiceCdnPolicyCacheKeyPolicyArray(object):
    @classmethod
    def to_proto(self, resources):
        if not resources:
            return resources
        return [BackendServiceCdnPolicyCacheKeyPolicy.to_proto(i) for i in resources]

    @classmethod
    def from_proto(self, resources):
        return [BackendServiceCdnPolicyCacheKeyPolicy.from_proto(i) for i in resources]


class BackendServiceCdnPolicyNegativeCachingPolicy(object):
    def __init__(self, code: int = None, ttl: int = None):
        self.code = code
        self.ttl = ttl

    @classmethod
    def to_proto(self, resource):
        if not resource:
            return None

        res = (
            backend_service_pb2.ComputeBetaBackendServiceCdnPolicyNegativeCachingPolicy()
        )
        if Primitive.to_proto(resource.code):
            res.code = Primitive.to_proto(resource.code)
        if Primitive.to_proto(resource.ttl):
            res.ttl = Primitive.to_proto(resource.ttl)
        return res

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return None

        return BackendServiceCdnPolicyNegativeCachingPolicy(
            code=Primitive.from_proto(resource.code),
            ttl=Primitive.from_proto(resource.ttl),
        )


class BackendServiceCdnPolicyNegativeCachingPolicyArray(object):
    @classmethod
    def to_proto(self, resources):
        if not resources:
            return resources
        return [
            BackendServiceCdnPolicyNegativeCachingPolicy.to_proto(i) for i in resources
        ]

    @classmethod
    def from_proto(self, resources):
        return [
            BackendServiceCdnPolicyNegativeCachingPolicy.from_proto(i)
            for i in resources
        ]


class BackendServiceCdnPolicyBypassCacheOnRequestHeaders(object):
    def __init__(self, header_name: str = None):
        self.header_name = header_name

    @classmethod
    def to_proto(self, resource):
        if not resource:
            return None

        res = (
            backend_service_pb2.ComputeBetaBackendServiceCdnPolicyBypassCacheOnRequestHeaders()
        )
        if Primitive.to_proto(resource.header_name):
            res.header_name = Primitive.to_proto(resource.header_name)
        return res

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return None

        return BackendServiceCdnPolicyBypassCacheOnRequestHeaders(
            header_name=Primitive.from_proto(resource.header_name),
        )


class BackendServiceCdnPolicyBypassCacheOnRequestHeadersArray(object):
    @classmethod
    def to_proto(self, resources):
        if not resources:
            return resources
        return [
            BackendServiceCdnPolicyBypassCacheOnRequestHeaders.to_proto(i)
            for i in resources
        ]

    @classmethod
    def from_proto(self, resources):
        return [
            BackendServiceCdnPolicyBypassCacheOnRequestHeaders.from_proto(i)
            for i in resources
        ]


class BackendServiceLogConfig(object):
    def __init__(self, enable: bool = None, sample_rate: float = None):
        self.enable = enable
        self.sample_rate = sample_rate

    @classmethod
    def to_proto(self, resource):
        if not resource:
            return None

        res = backend_service_pb2.ComputeBetaBackendServiceLogConfig()
        if Primitive.to_proto(resource.enable):
            res.enable = Primitive.to_proto(resource.enable)
        if Primitive.to_proto(resource.sample_rate):
            res.sample_rate = Primitive.to_proto(resource.sample_rate)
        return res

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return None

        return BackendServiceLogConfig(
            enable=Primitive.from_proto(resource.enable),
            sample_rate=Primitive.from_proto(resource.sample_rate),
        )


class BackendServiceLogConfigArray(object):
    @classmethod
    def to_proto(self, resources):
        if not resources:
            return resources
        return [BackendServiceLogConfig.to_proto(i) for i in resources]

    @classmethod
    def from_proto(self, resources):
        return [BackendServiceLogConfig.from_proto(i) for i in resources]


class BackendServiceSecuritySettings(object):
    def __init__(
        self,
        client_tls_policy: str = None,
        authentication: str = None,
        subject_alt_names: list = None,
    ):
        self.client_tls_policy = client_tls_policy
        self.authentication = authentication
        self.subject_alt_names = subject_alt_names

    @classmethod
    def to_proto(self, resource):
        if not resource:
            return None

        res = backend_service_pb2.ComputeBetaBackendServiceSecuritySettings()
        if Primitive.to_proto(resource.client_tls_policy):
            res.client_tls_policy = Primitive.to_proto(resource.client_tls_policy)
        if Primitive.to_proto(resource.authentication):
            res.authentication = Primitive.to_proto(resource.authentication)
        if Primitive.to_proto(resource.subject_alt_names):
            res.subject_alt_names.extend(Primitive.to_proto(resource.subject_alt_names))
        return res

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return None

        return BackendServiceSecuritySettings(
            client_tls_policy=Primitive.from_proto(resource.client_tls_policy),
            authentication=Primitive.from_proto(resource.authentication),
            subject_alt_names=Primitive.from_proto(resource.subject_alt_names),
        )


class BackendServiceSecuritySettingsArray(object):
    @classmethod
    def to_proto(self, resources):
        if not resources:
            return resources
        return [BackendServiceSecuritySettings.to_proto(i) for i in resources]

    @classmethod
    def from_proto(self, resources):
        return [BackendServiceSecuritySettings.from_proto(i) for i in resources]


class BackendServiceConsistentHash(object):
    def __init__(
        self,
        http_cookie: dict = None,
        http_header_name: str = None,
        minimum_ring_size: int = None,
    ):
        self.http_cookie = http_cookie
        self.http_header_name = http_header_name
        self.minimum_ring_size = minimum_ring_size

    @classmethod
    def to_proto(self, resource):
        if not resource:
            return None

        res = backend_service_pb2.ComputeBetaBackendServiceConsistentHash()
        if BackendServiceConsistentHashHttpCookie.to_proto(resource.http_cookie):
            res.http_cookie.CopyFrom(
                BackendServiceConsistentHashHttpCookie.to_proto(resource.http_cookie)
            )
        else:
            res.ClearField("http_cookie")
        if Primitive.to_proto(resource.http_header_name):
            res.http_header_name = Primitive.to_proto(resource.http_header_name)
        if Primitive.to_proto(resource.minimum_ring_size):
            res.minimum_ring_size = Primitive.to_proto(resource.minimum_ring_size)
        return res

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return None

        return BackendServiceConsistentHash(
            http_cookie=BackendServiceConsistentHashHttpCookie.from_proto(
                resource.http_cookie
            ),
            http_header_name=Primitive.from_proto(resource.http_header_name),
            minimum_ring_size=Primitive.from_proto(resource.minimum_ring_size),
        )


class BackendServiceConsistentHashArray(object):
    @classmethod
    def to_proto(self, resources):
        if not resources:
            return resources
        return [BackendServiceConsistentHash.to_proto(i) for i in resources]

    @classmethod
    def from_proto(self, resources):
        return [BackendServiceConsistentHash.from_proto(i) for i in resources]


class BackendServiceConsistentHashHttpCookie(object):
    def __init__(self, name: str = None, path: str = None, ttl: dict = None):
        self.name = name
        self.path = path
        self.ttl = ttl

    @classmethod
    def to_proto(self, resource):
        if not resource:
            return None

        res = backend_service_pb2.ComputeBetaBackendServiceConsistentHashHttpCookie()
        if Primitive.to_proto(resource.name):
            res.name = Primitive.to_proto(resource.name)
        if Primitive.to_proto(resource.path):
            res.path = Primitive.to_proto(resource.path)
        if BackendServiceConsistentHashHttpCookieTtl.to_proto(resource.ttl):
            res.ttl.CopyFrom(
                BackendServiceConsistentHashHttpCookieTtl.to_proto(resource.ttl)
            )
        else:
            res.ClearField("ttl")
        return res

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return None

        return BackendServiceConsistentHashHttpCookie(
            name=Primitive.from_proto(resource.name),
            path=Primitive.from_proto(resource.path),
            ttl=BackendServiceConsistentHashHttpCookieTtl.from_proto(resource.ttl),
        )


class BackendServiceConsistentHashHttpCookieArray(object):
    @classmethod
    def to_proto(self, resources):
        if not resources:
            return resources
        return [BackendServiceConsistentHashHttpCookie.to_proto(i) for i in resources]

    @classmethod
    def from_proto(self, resources):
        return [BackendServiceConsistentHashHttpCookie.from_proto(i) for i in resources]


class BackendServiceConsistentHashHttpCookieTtl(object):
    def __init__(self, seconds: int = None, nanos: int = None):
        self.seconds = seconds
        self.nanos = nanos

    @classmethod
    def to_proto(self, resource):
        if not resource:
            return None

        res = backend_service_pb2.ComputeBetaBackendServiceConsistentHashHttpCookieTtl()
        if Primitive.to_proto(resource.seconds):
            res.seconds = Primitive.to_proto(resource.seconds)
        if Primitive.to_proto(resource.nanos):
            res.nanos = Primitive.to_proto(resource.nanos)
        return res

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return None

        return BackendServiceConsistentHashHttpCookieTtl(
            seconds=Primitive.from_proto(resource.seconds),
            nanos=Primitive.from_proto(resource.nanos),
        )


class BackendServiceConsistentHashHttpCookieTtlArray(object):
    @classmethod
    def to_proto(self, resources):
        if not resources:
            return resources
        return [
            BackendServiceConsistentHashHttpCookieTtl.to_proto(i) for i in resources
        ]

    @classmethod
    def from_proto(self, resources):
        return [
            BackendServiceConsistentHashHttpCookieTtl.from_proto(i) for i in resources
        ]


class BackendServiceCircuitBreakers(object):
    def __init__(
        self,
        connect_timeout: dict = None,
        max_requests_per_connection: int = None,
        max_connections: int = None,
        max_pending_requests: int = None,
        max_requests: int = None,
        max_retries: int = None,
    ):
        self.connect_timeout = connect_timeout
        self.max_requests_per_connection = max_requests_per_connection
        self.max_connections = max_connections
        self.max_pending_requests = max_pending_requests
        self.max_requests = max_requests
        self.max_retries = max_retries

    @classmethod
    def to_proto(self, resource):
        if not resource:
            return None

        res = backend_service_pb2.ComputeBetaBackendServiceCircuitBreakers()
        if BackendServiceCircuitBreakersConnectTimeout.to_proto(
            resource.connect_timeout
        ):
            res.connect_timeout.CopyFrom(
                BackendServiceCircuitBreakersConnectTimeout.to_proto(
                    resource.connect_timeout
                )
            )
        else:
            res.ClearField("connect_timeout")
        if Primitive.to_proto(resource.max_requests_per_connection):
            res.max_requests_per_connection = Primitive.to_proto(
                resource.max_requests_per_connection
            )
        if Primitive.to_proto(resource.max_connections):
            res.max_connections = Primitive.to_proto(resource.max_connections)
        if Primitive.to_proto(resource.max_pending_requests):
            res.max_pending_requests = Primitive.to_proto(resource.max_pending_requests)
        if Primitive.to_proto(resource.max_requests):
            res.max_requests = Primitive.to_proto(resource.max_requests)
        if Primitive.to_proto(resource.max_retries):
            res.max_retries = Primitive.to_proto(resource.max_retries)
        return res

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return None

        return BackendServiceCircuitBreakers(
            connect_timeout=BackendServiceCircuitBreakersConnectTimeout.from_proto(
                resource.connect_timeout
            ),
            max_requests_per_connection=Primitive.from_proto(
                resource.max_requests_per_connection
            ),
            max_connections=Primitive.from_proto(resource.max_connections),
            max_pending_requests=Primitive.from_proto(resource.max_pending_requests),
            max_requests=Primitive.from_proto(resource.max_requests),
            max_retries=Primitive.from_proto(resource.max_retries),
        )


class BackendServiceCircuitBreakersArray(object):
    @classmethod
    def to_proto(self, resources):
        if not resources:
            return resources
        return [BackendServiceCircuitBreakers.to_proto(i) for i in resources]

    @classmethod
    def from_proto(self, resources):
        return [BackendServiceCircuitBreakers.from_proto(i) for i in resources]


class BackendServiceCircuitBreakersConnectTimeout(object):
    def __init__(self, seconds: int = None, nanos: int = None):
        self.seconds = seconds
        self.nanos = nanos

    @classmethod
    def to_proto(self, resource):
        if not resource:
            return None

        res = (
            backend_service_pb2.ComputeBetaBackendServiceCircuitBreakersConnectTimeout()
        )
        if Primitive.to_proto(resource.seconds):
            res.seconds = Primitive.to_proto(resource.seconds)
        if Primitive.to_proto(resource.nanos):
            res.nanos = Primitive.to_proto(resource.nanos)
        return res

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return None

        return BackendServiceCircuitBreakersConnectTimeout(
            seconds=Primitive.from_proto(resource.seconds),
            nanos=Primitive.from_proto(resource.nanos),
        )


class BackendServiceCircuitBreakersConnectTimeoutArray(object):
    @classmethod
    def to_proto(self, resources):
        if not resources:
            return resources
        return [
            BackendServiceCircuitBreakersConnectTimeout.to_proto(i) for i in resources
        ]

    @classmethod
    def from_proto(self, resources):
        return [
            BackendServiceCircuitBreakersConnectTimeout.from_proto(i) for i in resources
        ]


class BackendServiceOutlierDetection(object):
    def __init__(
        self,
        consecutive_errors: int = None,
        interval: dict = None,
        base_ejection_time: dict = None,
        max_ejection_percent: int = None,
        enforcing_consecutive_errors: int = None,
        enforcing_success_rate: int = None,
        success_rate_minimum_hosts: int = None,
        success_rate_request_volume: int = None,
        success_rate_stdev_factor: int = None,
        consecutive_gateway_failure: int = None,
        enforcing_consecutive_gateway_failure: int = None,
    ):
        self.consecutive_errors = consecutive_errors
        self.interval = interval
        self.base_ejection_time = base_ejection_time
        self.max_ejection_percent = max_ejection_percent
        self.enforcing_consecutive_errors = enforcing_consecutive_errors
        self.enforcing_success_rate = enforcing_success_rate
        self.success_rate_minimum_hosts = success_rate_minimum_hosts
        self.success_rate_request_volume = success_rate_request_volume
        self.success_rate_stdev_factor = success_rate_stdev_factor
        self.consecutive_gateway_failure = consecutive_gateway_failure
        self.enforcing_consecutive_gateway_failure = (
            enforcing_consecutive_gateway_failure
        )

    @classmethod
    def to_proto(self, resource):
        if not resource:
            return None

        res = backend_service_pb2.ComputeBetaBackendServiceOutlierDetection()
        if Primitive.to_proto(resource.consecutive_errors):
            res.consecutive_errors = Primitive.to_proto(resource.consecutive_errors)
        if BackendServiceOutlierDetectionInterval.to_proto(resource.interval):
            res.interval.CopyFrom(
                BackendServiceOutlierDetectionInterval.to_proto(resource.interval)
            )
        else:
            res.ClearField("interval")
        if BackendServiceOutlierDetectionBaseEjectionTime.to_proto(
            resource.base_ejection_time
        ):
            res.base_ejection_time.CopyFrom(
                BackendServiceOutlierDetectionBaseEjectionTime.to_proto(
                    resource.base_ejection_time
                )
            )
        else:
            res.ClearField("base_ejection_time")
        if Primitive.to_proto(resource.max_ejection_percent):
            res.max_ejection_percent = Primitive.to_proto(resource.max_ejection_percent)
        if Primitive.to_proto(resource.enforcing_consecutive_errors):
            res.enforcing_consecutive_errors = Primitive.to_proto(
                resource.enforcing_consecutive_errors
            )
        if Primitive.to_proto(resource.enforcing_success_rate):
            res.enforcing_success_rate = Primitive.to_proto(
                resource.enforcing_success_rate
            )
        if Primitive.to_proto(resource.success_rate_minimum_hosts):
            res.success_rate_minimum_hosts = Primitive.to_proto(
                resource.success_rate_minimum_hosts
            )
        if Primitive.to_proto(resource.success_rate_request_volume):
            res.success_rate_request_volume = Primitive.to_proto(
                resource.success_rate_request_volume
            )
        if Primitive.to_proto(resource.success_rate_stdev_factor):
            res.success_rate_stdev_factor = Primitive.to_proto(
                resource.success_rate_stdev_factor
            )
        if Primitive.to_proto(resource.consecutive_gateway_failure):
            res.consecutive_gateway_failure = Primitive.to_proto(
                resource.consecutive_gateway_failure
            )
        if Primitive.to_proto(resource.enforcing_consecutive_gateway_failure):
            res.enforcing_consecutive_gateway_failure = Primitive.to_proto(
                resource.enforcing_consecutive_gateway_failure
            )
        return res

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return None

        return BackendServiceOutlierDetection(
            consecutive_errors=Primitive.from_proto(resource.consecutive_errors),
            interval=BackendServiceOutlierDetectionInterval.from_proto(
                resource.interval
            ),
            base_ejection_time=BackendServiceOutlierDetectionBaseEjectionTime.from_proto(
                resource.base_ejection_time
            ),
            max_ejection_percent=Primitive.from_proto(resource.max_ejection_percent),
            enforcing_consecutive_errors=Primitive.from_proto(
                resource.enforcing_consecutive_errors
            ),
            enforcing_success_rate=Primitive.from_proto(
                resource.enforcing_success_rate
            ),
            success_rate_minimum_hosts=Primitive.from_proto(
                resource.success_rate_minimum_hosts
            ),
            success_rate_request_volume=Primitive.from_proto(
                resource.success_rate_request_volume
            ),
            success_rate_stdev_factor=Primitive.from_proto(
                resource.success_rate_stdev_factor
            ),
            consecutive_gateway_failure=Primitive.from_proto(
                resource.consecutive_gateway_failure
            ),
            enforcing_consecutive_gateway_failure=Primitive.from_proto(
                resource.enforcing_consecutive_gateway_failure
            ),
        )


class BackendServiceOutlierDetectionArray(object):
    @classmethod
    def to_proto(self, resources):
        if not resources:
            return resources
        return [BackendServiceOutlierDetection.to_proto(i) for i in resources]

    @classmethod
    def from_proto(self, resources):
        return [BackendServiceOutlierDetection.from_proto(i) for i in resources]


class BackendServiceOutlierDetectionInterval(object):
    def __init__(self, seconds: int = None, nanos: int = None):
        self.seconds = seconds
        self.nanos = nanos

    @classmethod
    def to_proto(self, resource):
        if not resource:
            return None

        res = backend_service_pb2.ComputeBetaBackendServiceOutlierDetectionInterval()
        if Primitive.to_proto(resource.seconds):
            res.seconds = Primitive.to_proto(resource.seconds)
        if Primitive.to_proto(resource.nanos):
            res.nanos = Primitive.to_proto(resource.nanos)
        return res

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return None

        return BackendServiceOutlierDetectionInterval(
            seconds=Primitive.from_proto(resource.seconds),
            nanos=Primitive.from_proto(resource.nanos),
        )


class BackendServiceOutlierDetectionIntervalArray(object):
    @classmethod
    def to_proto(self, resources):
        if not resources:
            return resources
        return [BackendServiceOutlierDetectionInterval.to_proto(i) for i in resources]

    @classmethod
    def from_proto(self, resources):
        return [BackendServiceOutlierDetectionInterval.from_proto(i) for i in resources]


class BackendServiceOutlierDetectionBaseEjectionTime(object):
    def __init__(self, seconds: int = None, nanos: int = None):
        self.seconds = seconds
        self.nanos = nanos

    @classmethod
    def to_proto(self, resource):
        if not resource:
            return None

        res = (
            backend_service_pb2.ComputeBetaBackendServiceOutlierDetectionBaseEjectionTime()
        )
        if Primitive.to_proto(resource.seconds):
            res.seconds = Primitive.to_proto(resource.seconds)
        if Primitive.to_proto(resource.nanos):
            res.nanos = Primitive.to_proto(resource.nanos)
        return res

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return None

        return BackendServiceOutlierDetectionBaseEjectionTime(
            seconds=Primitive.from_proto(resource.seconds),
            nanos=Primitive.from_proto(resource.nanos),
        )


class BackendServiceOutlierDetectionBaseEjectionTimeArray(object):
    @classmethod
    def to_proto(self, resources):
        if not resources:
            return resources
        return [
            BackendServiceOutlierDetectionBaseEjectionTime.to_proto(i)
            for i in resources
        ]

    @classmethod
    def from_proto(self, resources):
        return [
            BackendServiceOutlierDetectionBaseEjectionTime.from_proto(i)
            for i in resources
        ]


class BackendServiceSubsetting(object):
    def __init__(self, policy: str = None):
        self.policy = policy

    @classmethod
    def to_proto(self, resource):
        if not resource:
            return None

        res = backend_service_pb2.ComputeBetaBackendServiceSubsetting()
        if BackendServiceSubsettingPolicyEnum.to_proto(resource.policy):
            res.policy = BackendServiceSubsettingPolicyEnum.to_proto(resource.policy)
        return res

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return None

        return BackendServiceSubsetting(
            policy=BackendServiceSubsettingPolicyEnum.from_proto(resource.policy),
        )


class BackendServiceSubsettingArray(object):
    @classmethod
    def to_proto(self, resources):
        if not resources:
            return resources
        return [BackendServiceSubsetting.to_proto(i) for i in resources]

    @classmethod
    def from_proto(self, resources):
        return [BackendServiceSubsetting.from_proto(i) for i in resources]


class BackendServiceConnectionTrackingPolicy(object):
    def __init__(
        self,
        tracking_mode: str = None,
        connection_persistence_on_unhealthy_backends: str = None,
        idle_timeout_sec: int = None,
    ):
        self.tracking_mode = tracking_mode
        self.connection_persistence_on_unhealthy_backends = (
            connection_persistence_on_unhealthy_backends
        )
        self.idle_timeout_sec = idle_timeout_sec

    @classmethod
    def to_proto(self, resource):
        if not resource:
            return None

        res = backend_service_pb2.ComputeBetaBackendServiceConnectionTrackingPolicy()
        if BackendServiceConnectionTrackingPolicyTrackingModeEnum.to_proto(
            resource.tracking_mode
        ):
            res.tracking_mode = BackendServiceConnectionTrackingPolicyTrackingModeEnum.to_proto(
                resource.tracking_mode
            )
        if BackendServiceConnectionTrackingPolicyConnectionPersistenceOnUnhealthyBackendsEnum.to_proto(
            resource.connection_persistence_on_unhealthy_backends
        ):
            res.connection_persistence_on_unhealthy_backends = BackendServiceConnectionTrackingPolicyConnectionPersistenceOnUnhealthyBackendsEnum.to_proto(
                resource.connection_persistence_on_unhealthy_backends
            )
        if Primitive.to_proto(resource.idle_timeout_sec):
            res.idle_timeout_sec = Primitive.to_proto(resource.idle_timeout_sec)
        return res

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return None

        return BackendServiceConnectionTrackingPolicy(
            tracking_mode=BackendServiceConnectionTrackingPolicyTrackingModeEnum.from_proto(
                resource.tracking_mode
            ),
            connection_persistence_on_unhealthy_backends=BackendServiceConnectionTrackingPolicyConnectionPersistenceOnUnhealthyBackendsEnum.from_proto(
                resource.connection_persistence_on_unhealthy_backends
            ),
            idle_timeout_sec=Primitive.from_proto(resource.idle_timeout_sec),
        )


class BackendServiceConnectionTrackingPolicyArray(object):
    @classmethod
    def to_proto(self, resources):
        if not resources:
            return resources
        return [BackendServiceConnectionTrackingPolicy.to_proto(i) for i in resources]

    @classmethod
    def from_proto(self, resources):
        return [BackendServiceConnectionTrackingPolicy.from_proto(i) for i in resources]


class BackendServiceMaxStreamDuration(object):
    def __init__(self, seconds: int = None, nanos: int = None):
        self.seconds = seconds
        self.nanos = nanos

    @classmethod
    def to_proto(self, resource):
        if not resource:
            return None

        res = backend_service_pb2.ComputeBetaBackendServiceMaxStreamDuration()
        if Primitive.to_proto(resource.seconds):
            res.seconds = Primitive.to_proto(resource.seconds)
        if Primitive.to_proto(resource.nanos):
            res.nanos = Primitive.to_proto(resource.nanos)
        return res

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return None

        return BackendServiceMaxStreamDuration(
            seconds=Primitive.from_proto(resource.seconds),
            nanos=Primitive.from_proto(resource.nanos),
        )


class BackendServiceMaxStreamDurationArray(object):
    @classmethod
    def to_proto(self, resources):
        if not resources:
            return resources
        return [BackendServiceMaxStreamDuration.to_proto(i) for i in resources]

    @classmethod
    def from_proto(self, resources):
        return [BackendServiceMaxStreamDuration.from_proto(i) for i in resources]


class BackendServiceBackendsBalancingModeEnum(object):
    @classmethod
    def to_proto(self, resource):
        if not resource:
            return resource
        return backend_service_pb2.ComputeBetaBackendServiceBackendsBalancingModeEnum.Value(
            "ComputeBetaBackendServiceBackendsBalancingModeEnum%s" % resource
        )

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return resource
        return backend_service_pb2.ComputeBetaBackendServiceBackendsBalancingModeEnum.Name(
            resource
        )[
            len("ComputeBetaBackendServiceBackendsBalancingModeEnum") :
        ]


class BackendServiceProtocolEnum(object):
    @classmethod
    def to_proto(self, resource):
        if not resource:
            return resource
        return backend_service_pb2.ComputeBetaBackendServiceProtocolEnum.Value(
            "ComputeBetaBackendServiceProtocolEnum%s" % resource
        )

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return resource
        return backend_service_pb2.ComputeBetaBackendServiceProtocolEnum.Name(resource)[
            len("ComputeBetaBackendServiceProtocolEnum") :
        ]


class BackendServiceSessionAffinityEnum(object):
    @classmethod
    def to_proto(self, resource):
        if not resource:
            return resource
        return backend_service_pb2.ComputeBetaBackendServiceSessionAffinityEnum.Value(
            "ComputeBetaBackendServiceSessionAffinityEnum%s" % resource
        )

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return resource
        return backend_service_pb2.ComputeBetaBackendServiceSessionAffinityEnum.Name(
            resource
        )[len("ComputeBetaBackendServiceSessionAffinityEnum") :]


class BackendServiceLoadBalancingSchemeEnum(object):
    @classmethod
    def to_proto(self, resource):
        if not resource:
            return resource
        return backend_service_pb2.ComputeBetaBackendServiceLoadBalancingSchemeEnum.Value(
            "ComputeBetaBackendServiceLoadBalancingSchemeEnum%s" % resource
        )

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return resource
        return backend_service_pb2.ComputeBetaBackendServiceLoadBalancingSchemeEnum.Name(
            resource
        )[
            len("ComputeBetaBackendServiceLoadBalancingSchemeEnum") :
        ]


class BackendServiceCdnPolicyCacheModeEnum(object):
    @classmethod
    def to_proto(self, resource):
        if not resource:
            return resource
        return backend_service_pb2.ComputeBetaBackendServiceCdnPolicyCacheModeEnum.Value(
            "ComputeBetaBackendServiceCdnPolicyCacheModeEnum%s" % resource
        )

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return resource
        return backend_service_pb2.ComputeBetaBackendServiceCdnPolicyCacheModeEnum.Name(
            resource
        )[len("ComputeBetaBackendServiceCdnPolicyCacheModeEnum") :]


class BackendServiceLocalityLbPolicyEnum(object):
    @classmethod
    def to_proto(self, resource):
        if not resource:
            return resource
        return backend_service_pb2.ComputeBetaBackendServiceLocalityLbPolicyEnum.Value(
            "ComputeBetaBackendServiceLocalityLbPolicyEnum%s" % resource
        )

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return resource
        return backend_service_pb2.ComputeBetaBackendServiceLocalityLbPolicyEnum.Name(
            resource
        )[len("ComputeBetaBackendServiceLocalityLbPolicyEnum") :]


class BackendServiceSubsettingPolicyEnum(object):
    @classmethod
    def to_proto(self, resource):
        if not resource:
            return resource
        return backend_service_pb2.ComputeBetaBackendServiceSubsettingPolicyEnum.Value(
            "ComputeBetaBackendServiceSubsettingPolicyEnum%s" % resource
        )

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return resource
        return backend_service_pb2.ComputeBetaBackendServiceSubsettingPolicyEnum.Name(
            resource
        )[len("ComputeBetaBackendServiceSubsettingPolicyEnum") :]


class BackendServiceConnectionTrackingPolicyTrackingModeEnum(object):
    @classmethod
    def to_proto(self, resource):
        if not resource:
            return resource
        return backend_service_pb2.ComputeBetaBackendServiceConnectionTrackingPolicyTrackingModeEnum.Value(
            "ComputeBetaBackendServiceConnectionTrackingPolicyTrackingModeEnum%s"
            % resource
        )

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return resource
        return backend_service_pb2.ComputeBetaBackendServiceConnectionTrackingPolicyTrackingModeEnum.Name(
            resource
        )[
            len("ComputeBetaBackendServiceConnectionTrackingPolicyTrackingModeEnum") :
        ]


class BackendServiceConnectionTrackingPolicyConnectionPersistenceOnUnhealthyBackendsEnum(
    object
):
    @classmethod
    def to_proto(self, resource):
        if not resource:
            return resource
        return backend_service_pb2.ComputeBetaBackendServiceConnectionTrackingPolicyConnectionPersistenceOnUnhealthyBackendsEnum.Value(
            "ComputeBetaBackendServiceConnectionTrackingPolicyConnectionPersistenceOnUnhealthyBackendsEnum%s"
            % resource
        )

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return resource
        return backend_service_pb2.ComputeBetaBackendServiceConnectionTrackingPolicyConnectionPersistenceOnUnhealthyBackendsEnum.Name(
            resource
        )[
            len(
                "ComputeBetaBackendServiceConnectionTrackingPolicyConnectionPersistenceOnUnhealthyBackendsEnum"
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
