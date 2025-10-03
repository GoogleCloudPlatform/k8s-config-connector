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
from google3.cloud.graphite.mmv2.services.google.assured_workloads import workload_pb2
from google3.cloud.graphite.mmv2.services.google.assured_workloads import (
    workload_pb2_grpc,
)

from typing import List


class Workload(object):
    def __init__(
        self,
        name: str = None,
        display_name: str = None,
        resources: list = None,
        compliance_regime: str = None,
        create_time: str = None,
        billing_account: str = None,
        labels: dict = None,
        provisioned_resources_parent: str = None,
        kms_settings: dict = None,
        resource_settings: list = None,
        kaj_enrollment_state: str = None,
        enable_sovereign_controls: bool = None,
        saa_enrollment_response: dict = None,
        compliance_status: dict = None,
        compliant_but_disallowed_services: list = None,
        partner: str = None,
        partner_permissions: dict = None,
        ekm_provisioning_response: dict = None,
        violation_notifications_enabled: bool = None,
        organization: str = None,
        location: str = None,
        service_account_file: str = "",
    ):
        channel.initialize()
        self.name = name
        self.display_name = display_name
        self.compliance_regime = compliance_regime
        self.billing_account = billing_account
        self.labels = labels
        self.provisioned_resources_parent = provisioned_resources_parent
        self.kms_settings = kms_settings
        self.resource_settings = resource_settings
        self.enable_sovereign_controls = enable_sovereign_controls
        self.partner = partner
        self.partner_permissions = partner_permissions
        self.violation_notifications_enabled = violation_notifications_enabled
        self.organization = organization
        self.location = location
        self.service_account_file = service_account_file

    def apply(self):
        stub = workload_pb2_grpc.AssuredworkloadsBetaWorkloadServiceStub(
            channel.Channel()
        )
        request = workload_pb2.ApplyAssuredworkloadsBetaWorkloadRequest()
        if Primitive.to_proto(self.name):
            request.resource.name = Primitive.to_proto(self.name)

        if Primitive.to_proto(self.display_name):
            request.resource.display_name = Primitive.to_proto(self.display_name)

        if WorkloadComplianceRegimeEnum.to_proto(self.compliance_regime):
            request.resource.compliance_regime = WorkloadComplianceRegimeEnum.to_proto(
                self.compliance_regime
            )

        if Primitive.to_proto(self.billing_account):
            request.resource.billing_account = Primitive.to_proto(self.billing_account)

        if Primitive.to_proto(self.labels):
            request.resource.labels = Primitive.to_proto(self.labels)

        if Primitive.to_proto(self.provisioned_resources_parent):
            request.resource.provisioned_resources_parent = Primitive.to_proto(
                self.provisioned_resources_parent
            )

        if WorkloadKmsSettings.to_proto(self.kms_settings):
            request.resource.kms_settings.CopyFrom(
                WorkloadKmsSettings.to_proto(self.kms_settings)
            )
        else:
            request.resource.ClearField("kms_settings")
        if WorkloadResourceSettingsArray.to_proto(self.resource_settings):
            request.resource.resource_settings.extend(
                WorkloadResourceSettingsArray.to_proto(self.resource_settings)
            )
        if Primitive.to_proto(self.enable_sovereign_controls):
            request.resource.enable_sovereign_controls = Primitive.to_proto(
                self.enable_sovereign_controls
            )

        if WorkloadPartnerEnum.to_proto(self.partner):
            request.resource.partner = WorkloadPartnerEnum.to_proto(self.partner)

        if WorkloadPartnerPermissions.to_proto(self.partner_permissions):
            request.resource.partner_permissions.CopyFrom(
                WorkloadPartnerPermissions.to_proto(self.partner_permissions)
            )
        else:
            request.resource.ClearField("partner_permissions")
        if Primitive.to_proto(self.violation_notifications_enabled):
            request.resource.violation_notifications_enabled = Primitive.to_proto(
                self.violation_notifications_enabled
            )

        if Primitive.to_proto(self.organization):
            request.resource.organization = Primitive.to_proto(self.organization)

        if Primitive.to_proto(self.location):
            request.resource.location = Primitive.to_proto(self.location)

        request.service_account_file = self.service_account_file

        response = stub.ApplyAssuredworkloadsBetaWorkload(request)
        self.name = Primitive.from_proto(response.name)
        self.display_name = Primitive.from_proto(response.display_name)
        self.resources = WorkloadResourcesArray.from_proto(response.resources)
        self.compliance_regime = WorkloadComplianceRegimeEnum.from_proto(
            response.compliance_regime
        )
        self.create_time = Primitive.from_proto(response.create_time)
        self.billing_account = Primitive.from_proto(response.billing_account)
        self.labels = Primitive.from_proto(response.labels)
        self.provisioned_resources_parent = Primitive.from_proto(
            response.provisioned_resources_parent
        )
        self.kms_settings = WorkloadKmsSettings.from_proto(response.kms_settings)
        self.resource_settings = WorkloadResourceSettingsArray.from_proto(
            response.resource_settings
        )
        self.kaj_enrollment_state = WorkloadKajEnrollmentStateEnum.from_proto(
            response.kaj_enrollment_state
        )
        self.enable_sovereign_controls = Primitive.from_proto(
            response.enable_sovereign_controls
        )
        self.saa_enrollment_response = WorkloadSaaEnrollmentResponse.from_proto(
            response.saa_enrollment_response
        )
        self.compliance_status = WorkloadComplianceStatus.from_proto(
            response.compliance_status
        )
        self.compliant_but_disallowed_services = Primitive.from_proto(
            response.compliant_but_disallowed_services
        )
        self.partner = WorkloadPartnerEnum.from_proto(response.partner)
        self.partner_permissions = WorkloadPartnerPermissions.from_proto(
            response.partner_permissions
        )
        self.ekm_provisioning_response = WorkloadEkmProvisioningResponse.from_proto(
            response.ekm_provisioning_response
        )
        self.violation_notifications_enabled = Primitive.from_proto(
            response.violation_notifications_enabled
        )
        self.organization = Primitive.from_proto(response.organization)
        self.location = Primitive.from_proto(response.location)

    def delete(self):
        stub = workload_pb2_grpc.AssuredworkloadsBetaWorkloadServiceStub(
            channel.Channel()
        )
        request = workload_pb2.DeleteAssuredworkloadsBetaWorkloadRequest()
        request.service_account_file = self.service_account_file
        if Primitive.to_proto(self.name):
            request.resource.name = Primitive.to_proto(self.name)

        if Primitive.to_proto(self.display_name):
            request.resource.display_name = Primitive.to_proto(self.display_name)

        if WorkloadComplianceRegimeEnum.to_proto(self.compliance_regime):
            request.resource.compliance_regime = WorkloadComplianceRegimeEnum.to_proto(
                self.compliance_regime
            )

        if Primitive.to_proto(self.billing_account):
            request.resource.billing_account = Primitive.to_proto(self.billing_account)

        if Primitive.to_proto(self.labels):
            request.resource.labels = Primitive.to_proto(self.labels)

        if Primitive.to_proto(self.provisioned_resources_parent):
            request.resource.provisioned_resources_parent = Primitive.to_proto(
                self.provisioned_resources_parent
            )

        if WorkloadKmsSettings.to_proto(self.kms_settings):
            request.resource.kms_settings.CopyFrom(
                WorkloadKmsSettings.to_proto(self.kms_settings)
            )
        else:
            request.resource.ClearField("kms_settings")
        if WorkloadResourceSettingsArray.to_proto(self.resource_settings):
            request.resource.resource_settings.extend(
                WorkloadResourceSettingsArray.to_proto(self.resource_settings)
            )
        if Primitive.to_proto(self.enable_sovereign_controls):
            request.resource.enable_sovereign_controls = Primitive.to_proto(
                self.enable_sovereign_controls
            )

        if WorkloadPartnerEnum.to_proto(self.partner):
            request.resource.partner = WorkloadPartnerEnum.to_proto(self.partner)

        if WorkloadPartnerPermissions.to_proto(self.partner_permissions):
            request.resource.partner_permissions.CopyFrom(
                WorkloadPartnerPermissions.to_proto(self.partner_permissions)
            )
        else:
            request.resource.ClearField("partner_permissions")
        if Primitive.to_proto(self.violation_notifications_enabled):
            request.resource.violation_notifications_enabled = Primitive.to_proto(
                self.violation_notifications_enabled
            )

        if Primitive.to_proto(self.organization):
            request.resource.organization = Primitive.to_proto(self.organization)

        if Primitive.to_proto(self.location):
            request.resource.location = Primitive.to_proto(self.location)

        response = stub.DeleteAssuredworkloadsBetaWorkload(request)

    @classmethod
    def list(self, organization, location, service_account_file=""):
        stub = workload_pb2_grpc.AssuredworkloadsBetaWorkloadServiceStub(
            channel.Channel()
        )
        request = workload_pb2.ListAssuredworkloadsBetaWorkloadRequest()
        request.service_account_file = service_account_file
        request.Organization = organization

        request.Location = location

        return stub.ListAssuredworkloadsBetaWorkload(request).items

    def to_proto(self):
        resource = workload_pb2.AssuredworkloadsBetaWorkload()
        if Primitive.to_proto(self.name):
            resource.name = Primitive.to_proto(self.name)
        if Primitive.to_proto(self.display_name):
            resource.display_name = Primitive.to_proto(self.display_name)
        if WorkloadComplianceRegimeEnum.to_proto(self.compliance_regime):
            resource.compliance_regime = WorkloadComplianceRegimeEnum.to_proto(
                self.compliance_regime
            )
        if Primitive.to_proto(self.billing_account):
            resource.billing_account = Primitive.to_proto(self.billing_account)
        if Primitive.to_proto(self.labels):
            resource.labels = Primitive.to_proto(self.labels)
        if Primitive.to_proto(self.provisioned_resources_parent):
            resource.provisioned_resources_parent = Primitive.to_proto(
                self.provisioned_resources_parent
            )
        if WorkloadKmsSettings.to_proto(self.kms_settings):
            resource.kms_settings.CopyFrom(
                WorkloadKmsSettings.to_proto(self.kms_settings)
            )
        else:
            resource.ClearField("kms_settings")
        if WorkloadResourceSettingsArray.to_proto(self.resource_settings):
            resource.resource_settings.extend(
                WorkloadResourceSettingsArray.to_proto(self.resource_settings)
            )
        if Primitive.to_proto(self.enable_sovereign_controls):
            resource.enable_sovereign_controls = Primitive.to_proto(
                self.enable_sovereign_controls
            )
        if WorkloadPartnerEnum.to_proto(self.partner):
            resource.partner = WorkloadPartnerEnum.to_proto(self.partner)
        if WorkloadPartnerPermissions.to_proto(self.partner_permissions):
            resource.partner_permissions.CopyFrom(
                WorkloadPartnerPermissions.to_proto(self.partner_permissions)
            )
        else:
            resource.ClearField("partner_permissions")
        if Primitive.to_proto(self.violation_notifications_enabled):
            resource.violation_notifications_enabled = Primitive.to_proto(
                self.violation_notifications_enabled
            )
        if Primitive.to_proto(self.organization):
            resource.organization = Primitive.to_proto(self.organization)
        if Primitive.to_proto(self.location):
            resource.location = Primitive.to_proto(self.location)
        return resource


class WorkloadResources(object):
    def __init__(self, resource_id: int = None, resource_type: str = None):
        self.resource_id = resource_id
        self.resource_type = resource_type

    @classmethod
    def to_proto(self, resource):
        if not resource:
            return None

        res = workload_pb2.AssuredworkloadsBetaWorkloadResources()
        if Primitive.to_proto(resource.resource_id):
            res.resource_id = Primitive.to_proto(resource.resource_id)
        if WorkloadResourcesResourceTypeEnum.to_proto(resource.resource_type):
            res.resource_type = WorkloadResourcesResourceTypeEnum.to_proto(
                resource.resource_type
            )
        return res

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return None

        return WorkloadResources(
            resource_id=Primitive.from_proto(resource.resource_id),
            resource_type=WorkloadResourcesResourceTypeEnum.from_proto(
                resource.resource_type
            ),
        )


class WorkloadResourcesArray(object):
    @classmethod
    def to_proto(self, resources):
        if not resources:
            return resources
        return [WorkloadResources.to_proto(i) for i in resources]

    @classmethod
    def from_proto(self, resources):
        return [WorkloadResources.from_proto(i) for i in resources]


class WorkloadKmsSettings(object):
    def __init__(self, next_rotation_time: str = None, rotation_period: str = None):
        self.next_rotation_time = next_rotation_time
        self.rotation_period = rotation_period

    @classmethod
    def to_proto(self, resource):
        if not resource:
            return None

        res = workload_pb2.AssuredworkloadsBetaWorkloadKmsSettings()
        if Primitive.to_proto(resource.next_rotation_time):
            res.next_rotation_time = Primitive.to_proto(resource.next_rotation_time)
        if Primitive.to_proto(resource.rotation_period):
            res.rotation_period = Primitive.to_proto(resource.rotation_period)
        return res

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return None

        return WorkloadKmsSettings(
            next_rotation_time=Primitive.from_proto(resource.next_rotation_time),
            rotation_period=Primitive.from_proto(resource.rotation_period),
        )


class WorkloadKmsSettingsArray(object):
    @classmethod
    def to_proto(self, resources):
        if not resources:
            return resources
        return [WorkloadKmsSettings.to_proto(i) for i in resources]

    @classmethod
    def from_proto(self, resources):
        return [WorkloadKmsSettings.from_proto(i) for i in resources]


class WorkloadResourceSettings(object):
    def __init__(
        self,
        resource_id: str = None,
        resource_type: str = None,
        display_name: str = None,
    ):
        self.resource_id = resource_id
        self.resource_type = resource_type
        self.display_name = display_name

    @classmethod
    def to_proto(self, resource):
        if not resource:
            return None

        res = workload_pb2.AssuredworkloadsBetaWorkloadResourceSettings()
        if Primitive.to_proto(resource.resource_id):
            res.resource_id = Primitive.to_proto(resource.resource_id)
        if WorkloadResourceSettingsResourceTypeEnum.to_proto(resource.resource_type):
            res.resource_type = WorkloadResourceSettingsResourceTypeEnum.to_proto(
                resource.resource_type
            )
        if Primitive.to_proto(resource.display_name):
            res.display_name = Primitive.to_proto(resource.display_name)
        return res

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return None

        return WorkloadResourceSettings(
            resource_id=Primitive.from_proto(resource.resource_id),
            resource_type=WorkloadResourceSettingsResourceTypeEnum.from_proto(
                resource.resource_type
            ),
            display_name=Primitive.from_proto(resource.display_name),
        )


class WorkloadResourceSettingsArray(object):
    @classmethod
    def to_proto(self, resources):
        if not resources:
            return resources
        return [WorkloadResourceSettings.to_proto(i) for i in resources]

    @classmethod
    def from_proto(self, resources):
        return [WorkloadResourceSettings.from_proto(i) for i in resources]


class WorkloadSaaEnrollmentResponse(object):
    def __init__(self, setup_errors: list = None, setup_status: str = None):
        self.setup_errors = setup_errors
        self.setup_status = setup_status

    @classmethod
    def to_proto(self, resource):
        if not resource:
            return None

        res = workload_pb2.AssuredworkloadsBetaWorkloadSaaEnrollmentResponse()
        if WorkloadSaaEnrollmentResponseSetupErrorsEnumArray.to_proto(
            resource.setup_errors
        ):
            res.setup_errors.extend(
                WorkloadSaaEnrollmentResponseSetupErrorsEnumArray.to_proto(
                    resource.setup_errors
                )
            )
        if WorkloadSaaEnrollmentResponseSetupStatusEnum.to_proto(resource.setup_status):
            res.setup_status = WorkloadSaaEnrollmentResponseSetupStatusEnum.to_proto(
                resource.setup_status
            )
        return res

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return None

        return WorkloadSaaEnrollmentResponse(
            setup_errors=WorkloadSaaEnrollmentResponseSetupErrorsEnumArray.from_proto(
                resource.setup_errors
            ),
            setup_status=WorkloadSaaEnrollmentResponseSetupStatusEnum.from_proto(
                resource.setup_status
            ),
        )


class WorkloadSaaEnrollmentResponseArray(object):
    @classmethod
    def to_proto(self, resources):
        if not resources:
            return resources
        return [WorkloadSaaEnrollmentResponse.to_proto(i) for i in resources]

    @classmethod
    def from_proto(self, resources):
        return [WorkloadSaaEnrollmentResponse.from_proto(i) for i in resources]


class WorkloadComplianceStatus(object):
    def __init__(
        self,
        active_violation_count: list = None,
        acknowledged_violation_count: list = None,
    ):
        self.active_violation_count = active_violation_count
        self.acknowledged_violation_count = acknowledged_violation_count

    @classmethod
    def to_proto(self, resource):
        if not resource:
            return None

        res = workload_pb2.AssuredworkloadsBetaWorkloadComplianceStatus()
        if int64Array.to_proto(resource.active_violation_count):
            res.active_violation_count.extend(
                int64Array.to_proto(resource.active_violation_count)
            )
        if int64Array.to_proto(resource.acknowledged_violation_count):
            res.acknowledged_violation_count.extend(
                int64Array.to_proto(resource.acknowledged_violation_count)
            )
        return res

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return None

        return WorkloadComplianceStatus(
            active_violation_count=int64Array.from_proto(
                resource.active_violation_count
            ),
            acknowledged_violation_count=int64Array.from_proto(
                resource.acknowledged_violation_count
            ),
        )


class WorkloadComplianceStatusArray(object):
    @classmethod
    def to_proto(self, resources):
        if not resources:
            return resources
        return [WorkloadComplianceStatus.to_proto(i) for i in resources]

    @classmethod
    def from_proto(self, resources):
        return [WorkloadComplianceStatus.from_proto(i) for i in resources]


class WorkloadPartnerPermissions(object):
    def __init__(
        self,
        data_logs_viewer: bool = None,
        service_access_approver: bool = None,
        assured_workloads_monitoring: bool = None,
    ):
        self.data_logs_viewer = data_logs_viewer
        self.service_access_approver = service_access_approver
        self.assured_workloads_monitoring = assured_workloads_monitoring

    @classmethod
    def to_proto(self, resource):
        if not resource:
            return None

        res = workload_pb2.AssuredworkloadsBetaWorkloadPartnerPermissions()
        if Primitive.to_proto(resource.data_logs_viewer):
            res.data_logs_viewer = Primitive.to_proto(resource.data_logs_viewer)
        if Primitive.to_proto(resource.service_access_approver):
            res.service_access_approver = Primitive.to_proto(
                resource.service_access_approver
            )
        if Primitive.to_proto(resource.assured_workloads_monitoring):
            res.assured_workloads_monitoring = Primitive.to_proto(
                resource.assured_workloads_monitoring
            )
        return res

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return None

        return WorkloadPartnerPermissions(
            data_logs_viewer=Primitive.from_proto(resource.data_logs_viewer),
            service_access_approver=Primitive.from_proto(
                resource.service_access_approver
            ),
            assured_workloads_monitoring=Primitive.from_proto(
                resource.assured_workloads_monitoring
            ),
        )


class WorkloadPartnerPermissionsArray(object):
    @classmethod
    def to_proto(self, resources):
        if not resources:
            return resources
        return [WorkloadPartnerPermissions.to_proto(i) for i in resources]

    @classmethod
    def from_proto(self, resources):
        return [WorkloadPartnerPermissions.from_proto(i) for i in resources]


class WorkloadEkmProvisioningResponse(object):
    def __init__(
        self,
        ekm_provisioning_state: str = None,
        ekm_Provisioning_error_domain: str = None,
        ekm_provisioning_error_mapping: str = None,
    ):
        self.ekm_provisioning_state = ekm_provisioning_state
        self.ekm_Provisioning_error_domain = ekm_Provisioning_error_domain
        self.ekm_provisioning_error_mapping = ekm_provisioning_error_mapping

    @classmethod
    def to_proto(self, resource):
        if not resource:
            return None

        res = workload_pb2.AssuredworkloadsBetaWorkloadEkmProvisioningResponse()
        if WorkloadEkmProvisioningResponseEkmProvisioningStateEnum.to_proto(
            resource.ekm_provisioning_state
        ):
            res.ekm_provisioning_state = (
                WorkloadEkmProvisioningResponseEkmProvisioningStateEnum.to_proto(
                    resource.ekm_provisioning_state
                )
            )
        if WorkloadEkmProvisioningResponseEkmProvisioningErrorDomainEnum.to_proto(
            resource.ekm_Provisioning_error_domain
        ):
            res.ekm_Provisioning_error_domain = (
                WorkloadEkmProvisioningResponseEkmProvisioningErrorDomainEnum.to_proto(
                    resource.ekm_Provisioning_error_domain
                )
            )
        if WorkloadEkmProvisioningResponseEkmProvisioningErrorMappingEnum.to_proto(
            resource.ekm_provisioning_error_mapping
        ):
            res.ekm_provisioning_error_mapping = (
                WorkloadEkmProvisioningResponseEkmProvisioningErrorMappingEnum.to_proto(
                    resource.ekm_provisioning_error_mapping
                )
            )
        return res

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return None

        return WorkloadEkmProvisioningResponse(
            ekm_provisioning_state=WorkloadEkmProvisioningResponseEkmProvisioningStateEnum.from_proto(
                resource.ekm_provisioning_state
            ),
            ekm_Provisioning_error_domain=WorkloadEkmProvisioningResponseEkmProvisioningErrorDomainEnum.from_proto(
                resource.ekm_Provisioning_error_domain
            ),
            ekm_provisioning_error_mapping=WorkloadEkmProvisioningResponseEkmProvisioningErrorMappingEnum.from_proto(
                resource.ekm_provisioning_error_mapping
            ),
        )


class WorkloadEkmProvisioningResponseArray(object):
    @classmethod
    def to_proto(self, resources):
        if not resources:
            return resources
        return [WorkloadEkmProvisioningResponse.to_proto(i) for i in resources]

    @classmethod
    def from_proto(self, resources):
        return [WorkloadEkmProvisioningResponse.from_proto(i) for i in resources]


class WorkloadResourcesResourceTypeEnum(object):
    @classmethod
    def to_proto(self, resource):
        if not resource:
            return resource
        return workload_pb2.AssuredworkloadsBetaWorkloadResourcesResourceTypeEnum.Value(
            "AssuredworkloadsBetaWorkloadResourcesResourceTypeEnum%s" % resource
        )

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return resource
        return workload_pb2.AssuredworkloadsBetaWorkloadResourcesResourceTypeEnum.Name(
            resource
        )[len("AssuredworkloadsBetaWorkloadResourcesResourceTypeEnum") :]


class WorkloadComplianceRegimeEnum(object):
    @classmethod
    def to_proto(self, resource):
        if not resource:
            return resource
        return workload_pb2.AssuredworkloadsBetaWorkloadComplianceRegimeEnum.Value(
            "AssuredworkloadsBetaWorkloadComplianceRegimeEnum%s" % resource
        )

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return resource
        return workload_pb2.AssuredworkloadsBetaWorkloadComplianceRegimeEnum.Name(
            resource
        )[len("AssuredworkloadsBetaWorkloadComplianceRegimeEnum") :]


class WorkloadResourceSettingsResourceTypeEnum(object):
    @classmethod
    def to_proto(self, resource):
        if not resource:
            return resource
        return workload_pb2.AssuredworkloadsBetaWorkloadResourceSettingsResourceTypeEnum.Value(
            "AssuredworkloadsBetaWorkloadResourceSettingsResourceTypeEnum%s" % resource
        )

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return resource
        return workload_pb2.AssuredworkloadsBetaWorkloadResourceSettingsResourceTypeEnum.Name(
            resource
        )[
            len("AssuredworkloadsBetaWorkloadResourceSettingsResourceTypeEnum") :
        ]


class WorkloadKajEnrollmentStateEnum(object):
    @classmethod
    def to_proto(self, resource):
        if not resource:
            return resource
        return workload_pb2.AssuredworkloadsBetaWorkloadKajEnrollmentStateEnum.Value(
            "AssuredworkloadsBetaWorkloadKajEnrollmentStateEnum%s" % resource
        )

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return resource
        return workload_pb2.AssuredworkloadsBetaWorkloadKajEnrollmentStateEnum.Name(
            resource
        )[len("AssuredworkloadsBetaWorkloadKajEnrollmentStateEnum") :]


class WorkloadSaaEnrollmentResponseSetupErrorsEnum(object):
    @classmethod
    def to_proto(self, resource):
        if not resource:
            return resource
        return workload_pb2.AssuredworkloadsBetaWorkloadSaaEnrollmentResponseSetupErrorsEnum.Value(
            "AssuredworkloadsBetaWorkloadSaaEnrollmentResponseSetupErrorsEnum%s"
            % resource
        )

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return resource
        return workload_pb2.AssuredworkloadsBetaWorkloadSaaEnrollmentResponseSetupErrorsEnum.Name(
            resource
        )[
            len("AssuredworkloadsBetaWorkloadSaaEnrollmentResponseSetupErrorsEnum") :
        ]


class WorkloadSaaEnrollmentResponseSetupStatusEnum(object):
    @classmethod
    def to_proto(self, resource):
        if not resource:
            return resource
        return workload_pb2.AssuredworkloadsBetaWorkloadSaaEnrollmentResponseSetupStatusEnum.Value(
            "AssuredworkloadsBetaWorkloadSaaEnrollmentResponseSetupStatusEnum%s"
            % resource
        )

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return resource
        return workload_pb2.AssuredworkloadsBetaWorkloadSaaEnrollmentResponseSetupStatusEnum.Name(
            resource
        )[
            len("AssuredworkloadsBetaWorkloadSaaEnrollmentResponseSetupStatusEnum") :
        ]


class WorkloadPartnerEnum(object):
    @classmethod
    def to_proto(self, resource):
        if not resource:
            return resource
        return workload_pb2.AssuredworkloadsBetaWorkloadPartnerEnum.Value(
            "AssuredworkloadsBetaWorkloadPartnerEnum%s" % resource
        )

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return resource
        return workload_pb2.AssuredworkloadsBetaWorkloadPartnerEnum.Name(resource)[
            len("AssuredworkloadsBetaWorkloadPartnerEnum") :
        ]


class WorkloadEkmProvisioningResponseEkmProvisioningStateEnum(object):
    @classmethod
    def to_proto(self, resource):
        if not resource:
            return resource
        return workload_pb2.AssuredworkloadsBetaWorkloadEkmProvisioningResponseEkmProvisioningStateEnum.Value(
            "AssuredworkloadsBetaWorkloadEkmProvisioningResponseEkmProvisioningStateEnum%s"
            % resource
        )

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return resource
        return workload_pb2.AssuredworkloadsBetaWorkloadEkmProvisioningResponseEkmProvisioningStateEnum.Name(
            resource
        )[
            len(
                "AssuredworkloadsBetaWorkloadEkmProvisioningResponseEkmProvisioningStateEnum"
            ) :
        ]


class WorkloadEkmProvisioningResponseEkmProvisioningErrorDomainEnum(object):
    @classmethod
    def to_proto(self, resource):
        if not resource:
            return resource
        return workload_pb2.AssuredworkloadsBetaWorkloadEkmProvisioningResponseEkmProvisioningErrorDomainEnum.Value(
            "AssuredworkloadsBetaWorkloadEkmProvisioningResponseEkmProvisioningErrorDomainEnum%s"
            % resource
        )

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return resource
        return workload_pb2.AssuredworkloadsBetaWorkloadEkmProvisioningResponseEkmProvisioningErrorDomainEnum.Name(
            resource
        )[
            len(
                "AssuredworkloadsBetaWorkloadEkmProvisioningResponseEkmProvisioningErrorDomainEnum"
            ) :
        ]


class WorkloadEkmProvisioningResponseEkmProvisioningErrorMappingEnum(object):
    @classmethod
    def to_proto(self, resource):
        if not resource:
            return resource
        return workload_pb2.AssuredworkloadsBetaWorkloadEkmProvisioningResponseEkmProvisioningErrorMappingEnum.Value(
            "AssuredworkloadsBetaWorkloadEkmProvisioningResponseEkmProvisioningErrorMappingEnum%s"
            % resource
        )

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return resource
        return workload_pb2.AssuredworkloadsBetaWorkloadEkmProvisioningResponseEkmProvisioningErrorMappingEnum.Name(
            resource
        )[
            len(
                "AssuredworkloadsBetaWorkloadEkmProvisioningResponseEkmProvisioningErrorMappingEnum"
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
