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
from google3.cloud.graphite.mmv2.services.google.tier2 import instance_pb2
from google3.cloud.graphite.mmv2.services.google.tier2 import instance_pb2_grpc

from typing import List


class Instance(object):
    def __init__(
        self,
        name: str = None,
        display_name: str = None,
        labels: dict = None,
        zone: str = None,
        alternative_zone: str = None,
        sku: dict = None,
        authorized_network_id: str = None,
        reserved_ip_range: str = None,
        host: str = None,
        port: int = None,
        current_zone: str = None,
        create_time: str = None,
        state: str = None,
        status_message: str = None,
        update_time: str = None,
        mutate_user_id: int = None,
        read_user_id: int = None,
        references: list = None,
        encryption_keys: list = None,
        preprocess_create_recipe: dict = None,
        create_recipe: dict = None,
        delete_recipe: dict = None,
        update_recipe: dict = None,
        preprocess_reset_recipe: dict = None,
        reset_recipe: dict = None,
        preprocess_repair_recipe: dict = None,
        repair_recipe: dict = None,
        preprocess_delete_recipe: dict = None,
        preprocess_update_recipe: dict = None,
        preprocess_freeze_recipe: dict = None,
        freeze_recipe: dict = None,
        preprocess_unfreeze_recipe: dict = None,
        unfreeze_recipe: dict = None,
        preprocess_report_instance_health_recipe: dict = None,
        report_instance_health_recipe: dict = None,
        preprocess_get_recipe: dict = None,
        notify_key_available_recipe: dict = None,
        notify_key_unavailable_recipe: dict = None,
        readonly_recipe: dict = None,
        reconcile_recipe: dict = None,
        preprocess_passthrough_recipe: dict = None,
        preprocess_reconcile_recipe: dict = None,
        enable_call_history: bool = None,
        history: list = None,
        public_resource_view_override: str = None,
        extra_info: str = None,
        uid: str = None,
        etag: str = None,
        project: str = None,
        location: str = None,
        service_account_file: str = "",
    ):

        channel.initialize()
        self.name = name
        self.display_name = display_name
        self.labels = labels
        self.zone = zone
        self.alternative_zone = alternative_zone
        self.sku = sku
        self.authorized_network_id = authorized_network_id
        self.reserved_ip_range = reserved_ip_range
        self.references = references
        self.preprocess_create_recipe = preprocess_create_recipe
        self.create_recipe = create_recipe
        self.delete_recipe = delete_recipe
        self.update_recipe = update_recipe
        self.preprocess_reset_recipe = preprocess_reset_recipe
        self.reset_recipe = reset_recipe
        self.preprocess_repair_recipe = preprocess_repair_recipe
        self.repair_recipe = repair_recipe
        self.preprocess_delete_recipe = preprocess_delete_recipe
        self.preprocess_update_recipe = preprocess_update_recipe
        self.preprocess_freeze_recipe = preprocess_freeze_recipe
        self.freeze_recipe = freeze_recipe
        self.preprocess_unfreeze_recipe = preprocess_unfreeze_recipe
        self.unfreeze_recipe = unfreeze_recipe
        self.preprocess_report_instance_health_recipe = (
            preprocess_report_instance_health_recipe
        )
        self.report_instance_health_recipe = report_instance_health_recipe
        self.preprocess_get_recipe = preprocess_get_recipe
        self.notify_key_available_recipe = notify_key_available_recipe
        self.notify_key_unavailable_recipe = notify_key_unavailable_recipe
        self.readonly_recipe = readonly_recipe
        self.reconcile_recipe = reconcile_recipe
        self.preprocess_passthrough_recipe = preprocess_passthrough_recipe
        self.preprocess_reconcile_recipe = preprocess_reconcile_recipe
        self.enable_call_history = enable_call_history
        self.history = history
        self.public_resource_view_override = public_resource_view_override
        self.uid = uid
        self.project = project
        self.location = location
        self.service_account_file = service_account_file

    def apply(self):
        stub = instance_pb2_grpc.Tier2AlphaInstanceServiceStub(channel.Channel())
        request = instance_pb2.ApplyTier2AlphaInstanceRequest()
        if Primitive.to_proto(self.name):
            request.resource.name = Primitive.to_proto(self.name)

        if Primitive.to_proto(self.display_name):
            request.resource.display_name = Primitive.to_proto(self.display_name)

        if Primitive.to_proto(self.labels):
            request.resource.labels = Primitive.to_proto(self.labels)

        if Primitive.to_proto(self.zone):
            request.resource.zone = Primitive.to_proto(self.zone)

        if Primitive.to_proto(self.alternative_zone):
            request.resource.alternative_zone = Primitive.to_proto(
                self.alternative_zone
            )

        if InstanceSku.to_proto(self.sku):
            request.resource.sku.CopyFrom(InstanceSku.to_proto(self.sku))
        else:
            request.resource.ClearField("sku")
        if Primitive.to_proto(self.authorized_network_id):
            request.resource.authorized_network_id = Primitive.to_proto(
                self.authorized_network_id
            )

        if Primitive.to_proto(self.reserved_ip_range):
            request.resource.reserved_ip_range = Primitive.to_proto(
                self.reserved_ip_range
            )

        if InstanceReferencesArray.to_proto(self.references):
            request.resource.references.extend(
                InstanceReferencesArray.to_proto(self.references)
            )
        if InstancePreprocessCreateRecipe.to_proto(self.preprocess_create_recipe):
            request.resource.preprocess_create_recipe.CopyFrom(
                InstancePreprocessCreateRecipe.to_proto(self.preprocess_create_recipe)
            )
        else:
            request.resource.ClearField("preprocess_create_recipe")
        if InstanceCreateRecipe.to_proto(self.create_recipe):
            request.resource.create_recipe.CopyFrom(
                InstanceCreateRecipe.to_proto(self.create_recipe)
            )
        else:
            request.resource.ClearField("create_recipe")
        if InstanceDeleteRecipe.to_proto(self.delete_recipe):
            request.resource.delete_recipe.CopyFrom(
                InstanceDeleteRecipe.to_proto(self.delete_recipe)
            )
        else:
            request.resource.ClearField("delete_recipe")
        if InstanceUpdateRecipe.to_proto(self.update_recipe):
            request.resource.update_recipe.CopyFrom(
                InstanceUpdateRecipe.to_proto(self.update_recipe)
            )
        else:
            request.resource.ClearField("update_recipe")
        if InstancePreprocessResetRecipe.to_proto(self.preprocess_reset_recipe):
            request.resource.preprocess_reset_recipe.CopyFrom(
                InstancePreprocessResetRecipe.to_proto(self.preprocess_reset_recipe)
            )
        else:
            request.resource.ClearField("preprocess_reset_recipe")
        if InstanceResetRecipe.to_proto(self.reset_recipe):
            request.resource.reset_recipe.CopyFrom(
                InstanceResetRecipe.to_proto(self.reset_recipe)
            )
        else:
            request.resource.ClearField("reset_recipe")
        if InstancePreprocessRepairRecipe.to_proto(self.preprocess_repair_recipe):
            request.resource.preprocess_repair_recipe.CopyFrom(
                InstancePreprocessRepairRecipe.to_proto(self.preprocess_repair_recipe)
            )
        else:
            request.resource.ClearField("preprocess_repair_recipe")
        if InstanceRepairRecipe.to_proto(self.repair_recipe):
            request.resource.repair_recipe.CopyFrom(
                InstanceRepairRecipe.to_proto(self.repair_recipe)
            )
        else:
            request.resource.ClearField("repair_recipe")
        if InstancePreprocessDeleteRecipe.to_proto(self.preprocess_delete_recipe):
            request.resource.preprocess_delete_recipe.CopyFrom(
                InstancePreprocessDeleteRecipe.to_proto(self.preprocess_delete_recipe)
            )
        else:
            request.resource.ClearField("preprocess_delete_recipe")
        if InstancePreprocessUpdateRecipe.to_proto(self.preprocess_update_recipe):
            request.resource.preprocess_update_recipe.CopyFrom(
                InstancePreprocessUpdateRecipe.to_proto(self.preprocess_update_recipe)
            )
        else:
            request.resource.ClearField("preprocess_update_recipe")
        if InstancePreprocessFreezeRecipe.to_proto(self.preprocess_freeze_recipe):
            request.resource.preprocess_freeze_recipe.CopyFrom(
                InstancePreprocessFreezeRecipe.to_proto(self.preprocess_freeze_recipe)
            )
        else:
            request.resource.ClearField("preprocess_freeze_recipe")
        if InstanceFreezeRecipe.to_proto(self.freeze_recipe):
            request.resource.freeze_recipe.CopyFrom(
                InstanceFreezeRecipe.to_proto(self.freeze_recipe)
            )
        else:
            request.resource.ClearField("freeze_recipe")
        if InstancePreprocessUnfreezeRecipe.to_proto(self.preprocess_unfreeze_recipe):
            request.resource.preprocess_unfreeze_recipe.CopyFrom(
                InstancePreprocessUnfreezeRecipe.to_proto(
                    self.preprocess_unfreeze_recipe
                )
            )
        else:
            request.resource.ClearField("preprocess_unfreeze_recipe")
        if InstanceUnfreezeRecipe.to_proto(self.unfreeze_recipe):
            request.resource.unfreeze_recipe.CopyFrom(
                InstanceUnfreezeRecipe.to_proto(self.unfreeze_recipe)
            )
        else:
            request.resource.ClearField("unfreeze_recipe")
        if InstancePreprocessReportInstanceHealthRecipe.to_proto(
            self.preprocess_report_instance_health_recipe
        ):
            request.resource.preprocess_report_instance_health_recipe.CopyFrom(
                InstancePreprocessReportInstanceHealthRecipe.to_proto(
                    self.preprocess_report_instance_health_recipe
                )
            )
        else:
            request.resource.ClearField("preprocess_report_instance_health_recipe")
        if InstanceReportInstanceHealthRecipe.to_proto(
            self.report_instance_health_recipe
        ):
            request.resource.report_instance_health_recipe.CopyFrom(
                InstanceReportInstanceHealthRecipe.to_proto(
                    self.report_instance_health_recipe
                )
            )
        else:
            request.resource.ClearField("report_instance_health_recipe")
        if InstancePreprocessGetRecipe.to_proto(self.preprocess_get_recipe):
            request.resource.preprocess_get_recipe.CopyFrom(
                InstancePreprocessGetRecipe.to_proto(self.preprocess_get_recipe)
            )
        else:
            request.resource.ClearField("preprocess_get_recipe")
        if InstanceNotifyKeyAvailableRecipe.to_proto(self.notify_key_available_recipe):
            request.resource.notify_key_available_recipe.CopyFrom(
                InstanceNotifyKeyAvailableRecipe.to_proto(
                    self.notify_key_available_recipe
                )
            )
        else:
            request.resource.ClearField("notify_key_available_recipe")
        if InstanceNotifyKeyUnavailableRecipe.to_proto(
            self.notify_key_unavailable_recipe
        ):
            request.resource.notify_key_unavailable_recipe.CopyFrom(
                InstanceNotifyKeyUnavailableRecipe.to_proto(
                    self.notify_key_unavailable_recipe
                )
            )
        else:
            request.resource.ClearField("notify_key_unavailable_recipe")
        if InstanceReadonlyRecipe.to_proto(self.readonly_recipe):
            request.resource.readonly_recipe.CopyFrom(
                InstanceReadonlyRecipe.to_proto(self.readonly_recipe)
            )
        else:
            request.resource.ClearField("readonly_recipe")
        if InstanceReconcileRecipe.to_proto(self.reconcile_recipe):
            request.resource.reconcile_recipe.CopyFrom(
                InstanceReconcileRecipe.to_proto(self.reconcile_recipe)
            )
        else:
            request.resource.ClearField("reconcile_recipe")
        if InstancePreprocessPassthroughRecipe.to_proto(
            self.preprocess_passthrough_recipe
        ):
            request.resource.preprocess_passthrough_recipe.CopyFrom(
                InstancePreprocessPassthroughRecipe.to_proto(
                    self.preprocess_passthrough_recipe
                )
            )
        else:
            request.resource.ClearField("preprocess_passthrough_recipe")
        if InstancePreprocessReconcileRecipe.to_proto(self.preprocess_reconcile_recipe):
            request.resource.preprocess_reconcile_recipe.CopyFrom(
                InstancePreprocessReconcileRecipe.to_proto(
                    self.preprocess_reconcile_recipe
                )
            )
        else:
            request.resource.ClearField("preprocess_reconcile_recipe")
        if Primitive.to_proto(self.enable_call_history):
            request.resource.enable_call_history = Primitive.to_proto(
                self.enable_call_history
            )

        if InstanceHistoryArray.to_proto(self.history):
            request.resource.history.extend(InstanceHistoryArray.to_proto(self.history))
        if Primitive.to_proto(self.public_resource_view_override):
            request.resource.public_resource_view_override = Primitive.to_proto(
                self.public_resource_view_override
            )

        if Primitive.to_proto(self.uid):
            request.resource.uid = Primitive.to_proto(self.uid)

        if Primitive.to_proto(self.project):
            request.resource.project = Primitive.to_proto(self.project)

        if Primitive.to_proto(self.location):
            request.resource.location = Primitive.to_proto(self.location)

        request.service_account_file = self.service_account_file

        response = stub.ApplyTier2AlphaInstance(request)
        self.name = Primitive.from_proto(response.name)
        self.display_name = Primitive.from_proto(response.display_name)
        self.labels = Primitive.from_proto(response.labels)
        self.zone = Primitive.from_proto(response.zone)
        self.alternative_zone = Primitive.from_proto(response.alternative_zone)
        self.sku = InstanceSku.from_proto(response.sku)
        self.authorized_network_id = Primitive.from_proto(
            response.authorized_network_id
        )
        self.reserved_ip_range = Primitive.from_proto(response.reserved_ip_range)
        self.host = Primitive.from_proto(response.host)
        self.port = Primitive.from_proto(response.port)
        self.current_zone = Primitive.from_proto(response.current_zone)
        self.create_time = Primitive.from_proto(response.create_time)
        self.state = InstanceStateEnum.from_proto(response.state)
        self.status_message = Primitive.from_proto(response.status_message)
        self.update_time = Primitive.from_proto(response.update_time)
        self.mutate_user_id = Primitive.from_proto(response.mutate_user_id)
        self.read_user_id = Primitive.from_proto(response.read_user_id)
        self.references = InstanceReferencesArray.from_proto(response.references)
        self.encryption_keys = InstanceEncryptionKeysArray.from_proto(
            response.encryption_keys
        )
        self.preprocess_create_recipe = InstancePreprocessCreateRecipe.from_proto(
            response.preprocess_create_recipe
        )
        self.create_recipe = InstanceCreateRecipe.from_proto(response.create_recipe)
        self.delete_recipe = InstanceDeleteRecipe.from_proto(response.delete_recipe)
        self.update_recipe = InstanceUpdateRecipe.from_proto(response.update_recipe)
        self.preprocess_reset_recipe = InstancePreprocessResetRecipe.from_proto(
            response.preprocess_reset_recipe
        )
        self.reset_recipe = InstanceResetRecipe.from_proto(response.reset_recipe)
        self.preprocess_repair_recipe = InstancePreprocessRepairRecipe.from_proto(
            response.preprocess_repair_recipe
        )
        self.repair_recipe = InstanceRepairRecipe.from_proto(response.repair_recipe)
        self.preprocess_delete_recipe = InstancePreprocessDeleteRecipe.from_proto(
            response.preprocess_delete_recipe
        )
        self.preprocess_update_recipe = InstancePreprocessUpdateRecipe.from_proto(
            response.preprocess_update_recipe
        )
        self.preprocess_freeze_recipe = InstancePreprocessFreezeRecipe.from_proto(
            response.preprocess_freeze_recipe
        )
        self.freeze_recipe = InstanceFreezeRecipe.from_proto(response.freeze_recipe)
        self.preprocess_unfreeze_recipe = InstancePreprocessUnfreezeRecipe.from_proto(
            response.preprocess_unfreeze_recipe
        )
        self.unfreeze_recipe = InstanceUnfreezeRecipe.from_proto(
            response.unfreeze_recipe
        )
        self.preprocess_report_instance_health_recipe = InstancePreprocessReportInstanceHealthRecipe.from_proto(
            response.preprocess_report_instance_health_recipe
        )
        self.report_instance_health_recipe = InstanceReportInstanceHealthRecipe.from_proto(
            response.report_instance_health_recipe
        )
        self.preprocess_get_recipe = InstancePreprocessGetRecipe.from_proto(
            response.preprocess_get_recipe
        )
        self.notify_key_available_recipe = InstanceNotifyKeyAvailableRecipe.from_proto(
            response.notify_key_available_recipe
        )
        self.notify_key_unavailable_recipe = InstanceNotifyKeyUnavailableRecipe.from_proto(
            response.notify_key_unavailable_recipe
        )
        self.readonly_recipe = InstanceReadonlyRecipe.from_proto(
            response.readonly_recipe
        )
        self.reconcile_recipe = InstanceReconcileRecipe.from_proto(
            response.reconcile_recipe
        )
        self.preprocess_passthrough_recipe = InstancePreprocessPassthroughRecipe.from_proto(
            response.preprocess_passthrough_recipe
        )
        self.preprocess_reconcile_recipe = InstancePreprocessReconcileRecipe.from_proto(
            response.preprocess_reconcile_recipe
        )
        self.enable_call_history = Primitive.from_proto(response.enable_call_history)
        self.history = InstanceHistoryArray.from_proto(response.history)
        self.public_resource_view_override = Primitive.from_proto(
            response.public_resource_view_override
        )
        self.extra_info = Primitive.from_proto(response.extra_info)
        self.uid = Primitive.from_proto(response.uid)
        self.etag = Primitive.from_proto(response.etag)
        self.project = Primitive.from_proto(response.project)
        self.location = Primitive.from_proto(response.location)

    def delete(self):
        stub = instance_pb2_grpc.Tier2AlphaInstanceServiceStub(channel.Channel())
        request = instance_pb2.DeleteTier2AlphaInstanceRequest()
        request.service_account_file = self.service_account_file
        if Primitive.to_proto(self.name):
            request.resource.name = Primitive.to_proto(self.name)

        if Primitive.to_proto(self.display_name):
            request.resource.display_name = Primitive.to_proto(self.display_name)

        if Primitive.to_proto(self.labels):
            request.resource.labels = Primitive.to_proto(self.labels)

        if Primitive.to_proto(self.zone):
            request.resource.zone = Primitive.to_proto(self.zone)

        if Primitive.to_proto(self.alternative_zone):
            request.resource.alternative_zone = Primitive.to_proto(
                self.alternative_zone
            )

        if InstanceSku.to_proto(self.sku):
            request.resource.sku.CopyFrom(InstanceSku.to_proto(self.sku))
        else:
            request.resource.ClearField("sku")
        if Primitive.to_proto(self.authorized_network_id):
            request.resource.authorized_network_id = Primitive.to_proto(
                self.authorized_network_id
            )

        if Primitive.to_proto(self.reserved_ip_range):
            request.resource.reserved_ip_range = Primitive.to_proto(
                self.reserved_ip_range
            )

        if InstanceReferencesArray.to_proto(self.references):
            request.resource.references.extend(
                InstanceReferencesArray.to_proto(self.references)
            )
        if InstancePreprocessCreateRecipe.to_proto(self.preprocess_create_recipe):
            request.resource.preprocess_create_recipe.CopyFrom(
                InstancePreprocessCreateRecipe.to_proto(self.preprocess_create_recipe)
            )
        else:
            request.resource.ClearField("preprocess_create_recipe")
        if InstanceCreateRecipe.to_proto(self.create_recipe):
            request.resource.create_recipe.CopyFrom(
                InstanceCreateRecipe.to_proto(self.create_recipe)
            )
        else:
            request.resource.ClearField("create_recipe")
        if InstanceDeleteRecipe.to_proto(self.delete_recipe):
            request.resource.delete_recipe.CopyFrom(
                InstanceDeleteRecipe.to_proto(self.delete_recipe)
            )
        else:
            request.resource.ClearField("delete_recipe")
        if InstanceUpdateRecipe.to_proto(self.update_recipe):
            request.resource.update_recipe.CopyFrom(
                InstanceUpdateRecipe.to_proto(self.update_recipe)
            )
        else:
            request.resource.ClearField("update_recipe")
        if InstancePreprocessResetRecipe.to_proto(self.preprocess_reset_recipe):
            request.resource.preprocess_reset_recipe.CopyFrom(
                InstancePreprocessResetRecipe.to_proto(self.preprocess_reset_recipe)
            )
        else:
            request.resource.ClearField("preprocess_reset_recipe")
        if InstanceResetRecipe.to_proto(self.reset_recipe):
            request.resource.reset_recipe.CopyFrom(
                InstanceResetRecipe.to_proto(self.reset_recipe)
            )
        else:
            request.resource.ClearField("reset_recipe")
        if InstancePreprocessRepairRecipe.to_proto(self.preprocess_repair_recipe):
            request.resource.preprocess_repair_recipe.CopyFrom(
                InstancePreprocessRepairRecipe.to_proto(self.preprocess_repair_recipe)
            )
        else:
            request.resource.ClearField("preprocess_repair_recipe")
        if InstanceRepairRecipe.to_proto(self.repair_recipe):
            request.resource.repair_recipe.CopyFrom(
                InstanceRepairRecipe.to_proto(self.repair_recipe)
            )
        else:
            request.resource.ClearField("repair_recipe")
        if InstancePreprocessDeleteRecipe.to_proto(self.preprocess_delete_recipe):
            request.resource.preprocess_delete_recipe.CopyFrom(
                InstancePreprocessDeleteRecipe.to_proto(self.preprocess_delete_recipe)
            )
        else:
            request.resource.ClearField("preprocess_delete_recipe")
        if InstancePreprocessUpdateRecipe.to_proto(self.preprocess_update_recipe):
            request.resource.preprocess_update_recipe.CopyFrom(
                InstancePreprocessUpdateRecipe.to_proto(self.preprocess_update_recipe)
            )
        else:
            request.resource.ClearField("preprocess_update_recipe")
        if InstancePreprocessFreezeRecipe.to_proto(self.preprocess_freeze_recipe):
            request.resource.preprocess_freeze_recipe.CopyFrom(
                InstancePreprocessFreezeRecipe.to_proto(self.preprocess_freeze_recipe)
            )
        else:
            request.resource.ClearField("preprocess_freeze_recipe")
        if InstanceFreezeRecipe.to_proto(self.freeze_recipe):
            request.resource.freeze_recipe.CopyFrom(
                InstanceFreezeRecipe.to_proto(self.freeze_recipe)
            )
        else:
            request.resource.ClearField("freeze_recipe")
        if InstancePreprocessUnfreezeRecipe.to_proto(self.preprocess_unfreeze_recipe):
            request.resource.preprocess_unfreeze_recipe.CopyFrom(
                InstancePreprocessUnfreezeRecipe.to_proto(
                    self.preprocess_unfreeze_recipe
                )
            )
        else:
            request.resource.ClearField("preprocess_unfreeze_recipe")
        if InstanceUnfreezeRecipe.to_proto(self.unfreeze_recipe):
            request.resource.unfreeze_recipe.CopyFrom(
                InstanceUnfreezeRecipe.to_proto(self.unfreeze_recipe)
            )
        else:
            request.resource.ClearField("unfreeze_recipe")
        if InstancePreprocessReportInstanceHealthRecipe.to_proto(
            self.preprocess_report_instance_health_recipe
        ):
            request.resource.preprocess_report_instance_health_recipe.CopyFrom(
                InstancePreprocessReportInstanceHealthRecipe.to_proto(
                    self.preprocess_report_instance_health_recipe
                )
            )
        else:
            request.resource.ClearField("preprocess_report_instance_health_recipe")
        if InstanceReportInstanceHealthRecipe.to_proto(
            self.report_instance_health_recipe
        ):
            request.resource.report_instance_health_recipe.CopyFrom(
                InstanceReportInstanceHealthRecipe.to_proto(
                    self.report_instance_health_recipe
                )
            )
        else:
            request.resource.ClearField("report_instance_health_recipe")
        if InstancePreprocessGetRecipe.to_proto(self.preprocess_get_recipe):
            request.resource.preprocess_get_recipe.CopyFrom(
                InstancePreprocessGetRecipe.to_proto(self.preprocess_get_recipe)
            )
        else:
            request.resource.ClearField("preprocess_get_recipe")
        if InstanceNotifyKeyAvailableRecipe.to_proto(self.notify_key_available_recipe):
            request.resource.notify_key_available_recipe.CopyFrom(
                InstanceNotifyKeyAvailableRecipe.to_proto(
                    self.notify_key_available_recipe
                )
            )
        else:
            request.resource.ClearField("notify_key_available_recipe")
        if InstanceNotifyKeyUnavailableRecipe.to_proto(
            self.notify_key_unavailable_recipe
        ):
            request.resource.notify_key_unavailable_recipe.CopyFrom(
                InstanceNotifyKeyUnavailableRecipe.to_proto(
                    self.notify_key_unavailable_recipe
                )
            )
        else:
            request.resource.ClearField("notify_key_unavailable_recipe")
        if InstanceReadonlyRecipe.to_proto(self.readonly_recipe):
            request.resource.readonly_recipe.CopyFrom(
                InstanceReadonlyRecipe.to_proto(self.readonly_recipe)
            )
        else:
            request.resource.ClearField("readonly_recipe")
        if InstanceReconcileRecipe.to_proto(self.reconcile_recipe):
            request.resource.reconcile_recipe.CopyFrom(
                InstanceReconcileRecipe.to_proto(self.reconcile_recipe)
            )
        else:
            request.resource.ClearField("reconcile_recipe")
        if InstancePreprocessPassthroughRecipe.to_proto(
            self.preprocess_passthrough_recipe
        ):
            request.resource.preprocess_passthrough_recipe.CopyFrom(
                InstancePreprocessPassthroughRecipe.to_proto(
                    self.preprocess_passthrough_recipe
                )
            )
        else:
            request.resource.ClearField("preprocess_passthrough_recipe")
        if InstancePreprocessReconcileRecipe.to_proto(self.preprocess_reconcile_recipe):
            request.resource.preprocess_reconcile_recipe.CopyFrom(
                InstancePreprocessReconcileRecipe.to_proto(
                    self.preprocess_reconcile_recipe
                )
            )
        else:
            request.resource.ClearField("preprocess_reconcile_recipe")
        if Primitive.to_proto(self.enable_call_history):
            request.resource.enable_call_history = Primitive.to_proto(
                self.enable_call_history
            )

        if InstanceHistoryArray.to_proto(self.history):
            request.resource.history.extend(InstanceHistoryArray.to_proto(self.history))
        if Primitive.to_proto(self.public_resource_view_override):
            request.resource.public_resource_view_override = Primitive.to_proto(
                self.public_resource_view_override
            )

        if Primitive.to_proto(self.uid):
            request.resource.uid = Primitive.to_proto(self.uid)

        if Primitive.to_proto(self.project):
            request.resource.project = Primitive.to_proto(self.project)

        if Primitive.to_proto(self.location):
            request.resource.location = Primitive.to_proto(self.location)

        response = stub.DeleteTier2AlphaInstance(request)

    def list(self):
        stub = instance_pb2_grpc.Tier2AlphaInstanceServiceStub(channel.Channel())
        request = instance_pb2.ListTier2AlphaInstanceRequest()
        request.service_account_file = self.service_account_file
        if Primitive.to_proto(self.name):
            request.resource.name = Primitive.to_proto(self.name)

        if Primitive.to_proto(self.display_name):
            request.resource.display_name = Primitive.to_proto(self.display_name)

        if Primitive.to_proto(self.labels):
            request.resource.labels = Primitive.to_proto(self.labels)

        if Primitive.to_proto(self.zone):
            request.resource.zone = Primitive.to_proto(self.zone)

        if Primitive.to_proto(self.alternative_zone):
            request.resource.alternative_zone = Primitive.to_proto(
                self.alternative_zone
            )

        if InstanceSku.to_proto(self.sku):
            request.resource.sku.CopyFrom(InstanceSku.to_proto(self.sku))
        else:
            request.resource.ClearField("sku")
        if Primitive.to_proto(self.authorized_network_id):
            request.resource.authorized_network_id = Primitive.to_proto(
                self.authorized_network_id
            )

        if Primitive.to_proto(self.reserved_ip_range):
            request.resource.reserved_ip_range = Primitive.to_proto(
                self.reserved_ip_range
            )

        if InstanceReferencesArray.to_proto(self.references):
            request.resource.references.extend(
                InstanceReferencesArray.to_proto(self.references)
            )
        if InstancePreprocessCreateRecipe.to_proto(self.preprocess_create_recipe):
            request.resource.preprocess_create_recipe.CopyFrom(
                InstancePreprocessCreateRecipe.to_proto(self.preprocess_create_recipe)
            )
        else:
            request.resource.ClearField("preprocess_create_recipe")
        if InstanceCreateRecipe.to_proto(self.create_recipe):
            request.resource.create_recipe.CopyFrom(
                InstanceCreateRecipe.to_proto(self.create_recipe)
            )
        else:
            request.resource.ClearField("create_recipe")
        if InstanceDeleteRecipe.to_proto(self.delete_recipe):
            request.resource.delete_recipe.CopyFrom(
                InstanceDeleteRecipe.to_proto(self.delete_recipe)
            )
        else:
            request.resource.ClearField("delete_recipe")
        if InstanceUpdateRecipe.to_proto(self.update_recipe):
            request.resource.update_recipe.CopyFrom(
                InstanceUpdateRecipe.to_proto(self.update_recipe)
            )
        else:
            request.resource.ClearField("update_recipe")
        if InstancePreprocessResetRecipe.to_proto(self.preprocess_reset_recipe):
            request.resource.preprocess_reset_recipe.CopyFrom(
                InstancePreprocessResetRecipe.to_proto(self.preprocess_reset_recipe)
            )
        else:
            request.resource.ClearField("preprocess_reset_recipe")
        if InstanceResetRecipe.to_proto(self.reset_recipe):
            request.resource.reset_recipe.CopyFrom(
                InstanceResetRecipe.to_proto(self.reset_recipe)
            )
        else:
            request.resource.ClearField("reset_recipe")
        if InstancePreprocessRepairRecipe.to_proto(self.preprocess_repair_recipe):
            request.resource.preprocess_repair_recipe.CopyFrom(
                InstancePreprocessRepairRecipe.to_proto(self.preprocess_repair_recipe)
            )
        else:
            request.resource.ClearField("preprocess_repair_recipe")
        if InstanceRepairRecipe.to_proto(self.repair_recipe):
            request.resource.repair_recipe.CopyFrom(
                InstanceRepairRecipe.to_proto(self.repair_recipe)
            )
        else:
            request.resource.ClearField("repair_recipe")
        if InstancePreprocessDeleteRecipe.to_proto(self.preprocess_delete_recipe):
            request.resource.preprocess_delete_recipe.CopyFrom(
                InstancePreprocessDeleteRecipe.to_proto(self.preprocess_delete_recipe)
            )
        else:
            request.resource.ClearField("preprocess_delete_recipe")
        if InstancePreprocessUpdateRecipe.to_proto(self.preprocess_update_recipe):
            request.resource.preprocess_update_recipe.CopyFrom(
                InstancePreprocessUpdateRecipe.to_proto(self.preprocess_update_recipe)
            )
        else:
            request.resource.ClearField("preprocess_update_recipe")
        if InstancePreprocessFreezeRecipe.to_proto(self.preprocess_freeze_recipe):
            request.resource.preprocess_freeze_recipe.CopyFrom(
                InstancePreprocessFreezeRecipe.to_proto(self.preprocess_freeze_recipe)
            )
        else:
            request.resource.ClearField("preprocess_freeze_recipe")
        if InstanceFreezeRecipe.to_proto(self.freeze_recipe):
            request.resource.freeze_recipe.CopyFrom(
                InstanceFreezeRecipe.to_proto(self.freeze_recipe)
            )
        else:
            request.resource.ClearField("freeze_recipe")
        if InstancePreprocessUnfreezeRecipe.to_proto(self.preprocess_unfreeze_recipe):
            request.resource.preprocess_unfreeze_recipe.CopyFrom(
                InstancePreprocessUnfreezeRecipe.to_proto(
                    self.preprocess_unfreeze_recipe
                )
            )
        else:
            request.resource.ClearField("preprocess_unfreeze_recipe")
        if InstanceUnfreezeRecipe.to_proto(self.unfreeze_recipe):
            request.resource.unfreeze_recipe.CopyFrom(
                InstanceUnfreezeRecipe.to_proto(self.unfreeze_recipe)
            )
        else:
            request.resource.ClearField("unfreeze_recipe")
        if InstancePreprocessReportInstanceHealthRecipe.to_proto(
            self.preprocess_report_instance_health_recipe
        ):
            request.resource.preprocess_report_instance_health_recipe.CopyFrom(
                InstancePreprocessReportInstanceHealthRecipe.to_proto(
                    self.preprocess_report_instance_health_recipe
                )
            )
        else:
            request.resource.ClearField("preprocess_report_instance_health_recipe")
        if InstanceReportInstanceHealthRecipe.to_proto(
            self.report_instance_health_recipe
        ):
            request.resource.report_instance_health_recipe.CopyFrom(
                InstanceReportInstanceHealthRecipe.to_proto(
                    self.report_instance_health_recipe
                )
            )
        else:
            request.resource.ClearField("report_instance_health_recipe")
        if InstancePreprocessGetRecipe.to_proto(self.preprocess_get_recipe):
            request.resource.preprocess_get_recipe.CopyFrom(
                InstancePreprocessGetRecipe.to_proto(self.preprocess_get_recipe)
            )
        else:
            request.resource.ClearField("preprocess_get_recipe")
        if InstanceNotifyKeyAvailableRecipe.to_proto(self.notify_key_available_recipe):
            request.resource.notify_key_available_recipe.CopyFrom(
                InstanceNotifyKeyAvailableRecipe.to_proto(
                    self.notify_key_available_recipe
                )
            )
        else:
            request.resource.ClearField("notify_key_available_recipe")
        if InstanceNotifyKeyUnavailableRecipe.to_proto(
            self.notify_key_unavailable_recipe
        ):
            request.resource.notify_key_unavailable_recipe.CopyFrom(
                InstanceNotifyKeyUnavailableRecipe.to_proto(
                    self.notify_key_unavailable_recipe
                )
            )
        else:
            request.resource.ClearField("notify_key_unavailable_recipe")
        if InstanceReadonlyRecipe.to_proto(self.readonly_recipe):
            request.resource.readonly_recipe.CopyFrom(
                InstanceReadonlyRecipe.to_proto(self.readonly_recipe)
            )
        else:
            request.resource.ClearField("readonly_recipe")
        if InstanceReconcileRecipe.to_proto(self.reconcile_recipe):
            request.resource.reconcile_recipe.CopyFrom(
                InstanceReconcileRecipe.to_proto(self.reconcile_recipe)
            )
        else:
            request.resource.ClearField("reconcile_recipe")
        if InstancePreprocessPassthroughRecipe.to_proto(
            self.preprocess_passthrough_recipe
        ):
            request.resource.preprocess_passthrough_recipe.CopyFrom(
                InstancePreprocessPassthroughRecipe.to_proto(
                    self.preprocess_passthrough_recipe
                )
            )
        else:
            request.resource.ClearField("preprocess_passthrough_recipe")
        if InstancePreprocessReconcileRecipe.to_proto(self.preprocess_reconcile_recipe):
            request.resource.preprocess_reconcile_recipe.CopyFrom(
                InstancePreprocessReconcileRecipe.to_proto(
                    self.preprocess_reconcile_recipe
                )
            )
        else:
            request.resource.ClearField("preprocess_reconcile_recipe")
        if Primitive.to_proto(self.enable_call_history):
            request.resource.enable_call_history = Primitive.to_proto(
                self.enable_call_history
            )

        if InstanceHistoryArray.to_proto(self.history):
            request.resource.history.extend(InstanceHistoryArray.to_proto(self.history))
        if Primitive.to_proto(self.public_resource_view_override):
            request.resource.public_resource_view_override = Primitive.to_proto(
                self.public_resource_view_override
            )

        if Primitive.to_proto(self.uid):
            request.resource.uid = Primitive.to_proto(self.uid)

        if Primitive.to_proto(self.project):
            request.resource.project = Primitive.to_proto(self.project)

        if Primitive.to_proto(self.location):
            request.resource.location = Primitive.to_proto(self.location)

        return stub.ListTier2AlphaInstance(request).items

    def to_proto(self):
        resource = instance_pb2.Tier2AlphaInstance()
        if Primitive.to_proto(self.name):
            resource.name = Primitive.to_proto(self.name)
        if Primitive.to_proto(self.display_name):
            resource.display_name = Primitive.to_proto(self.display_name)
        if Primitive.to_proto(self.labels):
            resource.labels = Primitive.to_proto(self.labels)
        if Primitive.to_proto(self.zone):
            resource.zone = Primitive.to_proto(self.zone)
        if Primitive.to_proto(self.alternative_zone):
            resource.alternative_zone = Primitive.to_proto(self.alternative_zone)
        if InstanceSku.to_proto(self.sku):
            resource.sku.CopyFrom(InstanceSku.to_proto(self.sku))
        else:
            resource.ClearField("sku")
        if Primitive.to_proto(self.authorized_network_id):
            resource.authorized_network_id = Primitive.to_proto(
                self.authorized_network_id
            )
        if Primitive.to_proto(self.reserved_ip_range):
            resource.reserved_ip_range = Primitive.to_proto(self.reserved_ip_range)
        if InstanceReferencesArray.to_proto(self.references):
            resource.references.extend(
                InstanceReferencesArray.to_proto(self.references)
            )
        if InstancePreprocessCreateRecipe.to_proto(self.preprocess_create_recipe):
            resource.preprocess_create_recipe.CopyFrom(
                InstancePreprocessCreateRecipe.to_proto(self.preprocess_create_recipe)
            )
        else:
            resource.ClearField("preprocess_create_recipe")
        if InstanceCreateRecipe.to_proto(self.create_recipe):
            resource.create_recipe.CopyFrom(
                InstanceCreateRecipe.to_proto(self.create_recipe)
            )
        else:
            resource.ClearField("create_recipe")
        if InstanceDeleteRecipe.to_proto(self.delete_recipe):
            resource.delete_recipe.CopyFrom(
                InstanceDeleteRecipe.to_proto(self.delete_recipe)
            )
        else:
            resource.ClearField("delete_recipe")
        if InstanceUpdateRecipe.to_proto(self.update_recipe):
            resource.update_recipe.CopyFrom(
                InstanceUpdateRecipe.to_proto(self.update_recipe)
            )
        else:
            resource.ClearField("update_recipe")
        if InstancePreprocessResetRecipe.to_proto(self.preprocess_reset_recipe):
            resource.preprocess_reset_recipe.CopyFrom(
                InstancePreprocessResetRecipe.to_proto(self.preprocess_reset_recipe)
            )
        else:
            resource.ClearField("preprocess_reset_recipe")
        if InstanceResetRecipe.to_proto(self.reset_recipe):
            resource.reset_recipe.CopyFrom(
                InstanceResetRecipe.to_proto(self.reset_recipe)
            )
        else:
            resource.ClearField("reset_recipe")
        if InstancePreprocessRepairRecipe.to_proto(self.preprocess_repair_recipe):
            resource.preprocess_repair_recipe.CopyFrom(
                InstancePreprocessRepairRecipe.to_proto(self.preprocess_repair_recipe)
            )
        else:
            resource.ClearField("preprocess_repair_recipe")
        if InstanceRepairRecipe.to_proto(self.repair_recipe):
            resource.repair_recipe.CopyFrom(
                InstanceRepairRecipe.to_proto(self.repair_recipe)
            )
        else:
            resource.ClearField("repair_recipe")
        if InstancePreprocessDeleteRecipe.to_proto(self.preprocess_delete_recipe):
            resource.preprocess_delete_recipe.CopyFrom(
                InstancePreprocessDeleteRecipe.to_proto(self.preprocess_delete_recipe)
            )
        else:
            resource.ClearField("preprocess_delete_recipe")
        if InstancePreprocessUpdateRecipe.to_proto(self.preprocess_update_recipe):
            resource.preprocess_update_recipe.CopyFrom(
                InstancePreprocessUpdateRecipe.to_proto(self.preprocess_update_recipe)
            )
        else:
            resource.ClearField("preprocess_update_recipe")
        if InstancePreprocessFreezeRecipe.to_proto(self.preprocess_freeze_recipe):
            resource.preprocess_freeze_recipe.CopyFrom(
                InstancePreprocessFreezeRecipe.to_proto(self.preprocess_freeze_recipe)
            )
        else:
            resource.ClearField("preprocess_freeze_recipe")
        if InstanceFreezeRecipe.to_proto(self.freeze_recipe):
            resource.freeze_recipe.CopyFrom(
                InstanceFreezeRecipe.to_proto(self.freeze_recipe)
            )
        else:
            resource.ClearField("freeze_recipe")
        if InstancePreprocessUnfreezeRecipe.to_proto(self.preprocess_unfreeze_recipe):
            resource.preprocess_unfreeze_recipe.CopyFrom(
                InstancePreprocessUnfreezeRecipe.to_proto(
                    self.preprocess_unfreeze_recipe
                )
            )
        else:
            resource.ClearField("preprocess_unfreeze_recipe")
        if InstanceUnfreezeRecipe.to_proto(self.unfreeze_recipe):
            resource.unfreeze_recipe.CopyFrom(
                InstanceUnfreezeRecipe.to_proto(self.unfreeze_recipe)
            )
        else:
            resource.ClearField("unfreeze_recipe")
        if InstancePreprocessReportInstanceHealthRecipe.to_proto(
            self.preprocess_report_instance_health_recipe
        ):
            resource.preprocess_report_instance_health_recipe.CopyFrom(
                InstancePreprocessReportInstanceHealthRecipe.to_proto(
                    self.preprocess_report_instance_health_recipe
                )
            )
        else:
            resource.ClearField("preprocess_report_instance_health_recipe")
        if InstanceReportInstanceHealthRecipe.to_proto(
            self.report_instance_health_recipe
        ):
            resource.report_instance_health_recipe.CopyFrom(
                InstanceReportInstanceHealthRecipe.to_proto(
                    self.report_instance_health_recipe
                )
            )
        else:
            resource.ClearField("report_instance_health_recipe")
        if InstancePreprocessGetRecipe.to_proto(self.preprocess_get_recipe):
            resource.preprocess_get_recipe.CopyFrom(
                InstancePreprocessGetRecipe.to_proto(self.preprocess_get_recipe)
            )
        else:
            resource.ClearField("preprocess_get_recipe")
        if InstanceNotifyKeyAvailableRecipe.to_proto(self.notify_key_available_recipe):
            resource.notify_key_available_recipe.CopyFrom(
                InstanceNotifyKeyAvailableRecipe.to_proto(
                    self.notify_key_available_recipe
                )
            )
        else:
            resource.ClearField("notify_key_available_recipe")
        if InstanceNotifyKeyUnavailableRecipe.to_proto(
            self.notify_key_unavailable_recipe
        ):
            resource.notify_key_unavailable_recipe.CopyFrom(
                InstanceNotifyKeyUnavailableRecipe.to_proto(
                    self.notify_key_unavailable_recipe
                )
            )
        else:
            resource.ClearField("notify_key_unavailable_recipe")
        if InstanceReadonlyRecipe.to_proto(self.readonly_recipe):
            resource.readonly_recipe.CopyFrom(
                InstanceReadonlyRecipe.to_proto(self.readonly_recipe)
            )
        else:
            resource.ClearField("readonly_recipe")
        if InstanceReconcileRecipe.to_proto(self.reconcile_recipe):
            resource.reconcile_recipe.CopyFrom(
                InstanceReconcileRecipe.to_proto(self.reconcile_recipe)
            )
        else:
            resource.ClearField("reconcile_recipe")
        if InstancePreprocessPassthroughRecipe.to_proto(
            self.preprocess_passthrough_recipe
        ):
            resource.preprocess_passthrough_recipe.CopyFrom(
                InstancePreprocessPassthroughRecipe.to_proto(
                    self.preprocess_passthrough_recipe
                )
            )
        else:
            resource.ClearField("preprocess_passthrough_recipe")
        if InstancePreprocessReconcileRecipe.to_proto(self.preprocess_reconcile_recipe):
            resource.preprocess_reconcile_recipe.CopyFrom(
                InstancePreprocessReconcileRecipe.to_proto(
                    self.preprocess_reconcile_recipe
                )
            )
        else:
            resource.ClearField("preprocess_reconcile_recipe")
        if Primitive.to_proto(self.enable_call_history):
            resource.enable_call_history = Primitive.to_proto(self.enable_call_history)
        if InstanceHistoryArray.to_proto(self.history):
            resource.history.extend(InstanceHistoryArray.to_proto(self.history))
        if Primitive.to_proto(self.public_resource_view_override):
            resource.public_resource_view_override = Primitive.to_proto(
                self.public_resource_view_override
            )
        if Primitive.to_proto(self.uid):
            resource.uid = Primitive.to_proto(self.uid)
        if Primitive.to_proto(self.project):
            resource.project = Primitive.to_proto(self.project)
        if Primitive.to_proto(self.location):
            resource.location = Primitive.to_proto(self.location)
        return resource


class InstanceSku(object):
    def __init__(self, tier: str = None, size: str = None):
        self.tier = tier
        self.size = size

    @classmethod
    def to_proto(self, resource):
        if not resource:
            return None

        res = instance_pb2.Tier2AlphaInstanceSku()
        if InstanceSkuTierEnum.to_proto(resource.tier):
            res.tier = InstanceSkuTierEnum.to_proto(resource.tier)
        if InstanceSkuSizeEnum.to_proto(resource.size):
            res.size = InstanceSkuSizeEnum.to_proto(resource.size)
        return res

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return None

        return InstanceSku(
            tier=InstanceSkuTierEnum.from_proto(resource.tier),
            size=InstanceSkuSizeEnum.from_proto(resource.size),
        )


class InstanceSkuArray(object):
    @classmethod
    def to_proto(self, resources):
        if not resources:
            return resources
        return [InstanceSku.to_proto(i) for i in resources]

    @classmethod
    def from_proto(self, resources):
        return [InstanceSku.from_proto(i) for i in resources]


class InstanceReferences(object):
    def __init__(
        self,
        name: str = None,
        type: str = None,
        source_resource: str = None,
        details: list = None,
        create_time: str = None,
    ):
        self.name = name
        self.type = type
        self.source_resource = source_resource
        self.details = details
        self.create_time = create_time

    @classmethod
    def to_proto(self, resource):
        if not resource:
            return None

        res = instance_pb2.Tier2AlphaInstanceReferences()
        if Primitive.to_proto(resource.name):
            res.name = Primitive.to_proto(resource.name)
        if Primitive.to_proto(resource.type):
            res.type = Primitive.to_proto(resource.type)
        if Primitive.to_proto(resource.source_resource):
            res.source_resource = Primitive.to_proto(resource.source_resource)
        if InstanceReferencesDetailsArray.to_proto(resource.details):
            res.details.extend(
                InstanceReferencesDetailsArray.to_proto(resource.details)
            )
        if Primitive.to_proto(resource.create_time):
            res.create_time = Primitive.to_proto(resource.create_time)
        return res

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return None

        return InstanceReferences(
            name=Primitive.from_proto(resource.name),
            type=Primitive.from_proto(resource.type),
            source_resource=Primitive.from_proto(resource.source_resource),
            details=InstanceReferencesDetailsArray.from_proto(resource.details),
            create_time=Primitive.from_proto(resource.create_time),
        )


class InstanceReferencesArray(object):
    @classmethod
    def to_proto(self, resources):
        if not resources:
            return resources
        return [InstanceReferences.to_proto(i) for i in resources]

    @classmethod
    def from_proto(self, resources):
        return [InstanceReferences.from_proto(i) for i in resources]


class InstanceReferencesDetails(object):
    def __init__(self, type_url: str = None, value: str = None):
        self.type_url = type_url
        self.value = value

    @classmethod
    def to_proto(self, resource):
        if not resource:
            return None

        res = instance_pb2.Tier2AlphaInstanceReferencesDetails()
        if Primitive.to_proto(resource.type_url):
            res.type_url = Primitive.to_proto(resource.type_url)
        if Primitive.to_proto(resource.value):
            res.value = Primitive.to_proto(resource.value)
        return res

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return None

        return InstanceReferencesDetails(
            type_url=Primitive.from_proto(resource.type_url),
            value=Primitive.from_proto(resource.value),
        )


class InstanceReferencesDetailsArray(object):
    @classmethod
    def to_proto(self, resources):
        if not resources:
            return resources
        return [InstanceReferencesDetails.to_proto(i) for i in resources]

    @classmethod
    def from_proto(self, resources):
        return [InstanceReferencesDetails.from_proto(i) for i in resources]


class InstanceEncryptionKeys(object):
    def __init__(
        self,
        key_or_version: str = None,
        grant: str = None,
        delegate: str = None,
        key_state: dict = None,
    ):
        self.key_or_version = key_or_version
        self.grant = grant
        self.delegate = delegate
        self.key_state = key_state

    @classmethod
    def to_proto(self, resource):
        if not resource:
            return None

        res = instance_pb2.Tier2AlphaInstanceEncryptionKeys()
        if Primitive.to_proto(resource.key_or_version):
            res.key_or_version = Primitive.to_proto(resource.key_or_version)
        if Primitive.to_proto(resource.grant):
            res.grant = Primitive.to_proto(resource.grant)
        if Primitive.to_proto(resource.delegate):
            res.delegate = Primitive.to_proto(resource.delegate)
        if InstanceEncryptionKeysKeyState.to_proto(resource.key_state):
            res.key_state.CopyFrom(
                InstanceEncryptionKeysKeyState.to_proto(resource.key_state)
            )
        else:
            res.ClearField("key_state")
        return res

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return None

        return InstanceEncryptionKeys(
            key_or_version=Primitive.from_proto(resource.key_or_version),
            grant=Primitive.from_proto(resource.grant),
            delegate=Primitive.from_proto(resource.delegate),
            key_state=InstanceEncryptionKeysKeyState.from_proto(resource.key_state),
        )


class InstanceEncryptionKeysArray(object):
    @classmethod
    def to_proto(self, resources):
        if not resources:
            return resources
        return [InstanceEncryptionKeys.to_proto(i) for i in resources]

    @classmethod
    def from_proto(self, resources):
        return [InstanceEncryptionKeys.from_proto(i) for i in resources]


class InstanceEncryptionKeysKeyState(object):
    def __init__(self, key_state_version: int = None, availability: dict = None):
        self.key_state_version = key_state_version
        self.availability = availability

    @classmethod
    def to_proto(self, resource):
        if not resource:
            return None

        res = instance_pb2.Tier2AlphaInstanceEncryptionKeysKeyState()
        if Primitive.to_proto(resource.key_state_version):
            res.key_state_version = Primitive.to_proto(resource.key_state_version)
        if InstanceEncryptionKeysKeyStateAvailability.to_proto(resource.availability):
            res.availability.CopyFrom(
                InstanceEncryptionKeysKeyStateAvailability.to_proto(
                    resource.availability
                )
            )
        else:
            res.ClearField("availability")
        return res

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return None

        return InstanceEncryptionKeysKeyState(
            key_state_version=Primitive.from_proto(resource.key_state_version),
            availability=InstanceEncryptionKeysKeyStateAvailability.from_proto(
                resource.availability
            ),
        )


class InstanceEncryptionKeysKeyStateArray(object):
    @classmethod
    def to_proto(self, resources):
        if not resources:
            return resources
        return [InstanceEncryptionKeysKeyState.to_proto(i) for i in resources]

    @classmethod
    def from_proto(self, resources):
        return [InstanceEncryptionKeysKeyState.from_proto(i) for i in resources]


class InstanceEncryptionKeysKeyStateAvailability(object):
    def __init__(
        self,
        permission_denied: bool = None,
        unknown_failure: bool = None,
        key_version_state: str = None,
    ):
        self.permission_denied = permission_denied
        self.unknown_failure = unknown_failure
        self.key_version_state = key_version_state

    @classmethod
    def to_proto(self, resource):
        if not resource:
            return None

        res = instance_pb2.Tier2AlphaInstanceEncryptionKeysKeyStateAvailability()
        if Primitive.to_proto(resource.permission_denied):
            res.permission_denied = Primitive.to_proto(resource.permission_denied)
        if Primitive.to_proto(resource.unknown_failure):
            res.unknown_failure = Primitive.to_proto(resource.unknown_failure)
        if InstanceEncryptionKeysKeyStateAvailabilityKeyVersionStateEnum.to_proto(
            resource.key_version_state
        ):
            res.key_version_state = InstanceEncryptionKeysKeyStateAvailabilityKeyVersionStateEnum.to_proto(
                resource.key_version_state
            )
        return res

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return None

        return InstanceEncryptionKeysKeyStateAvailability(
            permission_denied=Primitive.from_proto(resource.permission_denied),
            unknown_failure=Primitive.from_proto(resource.unknown_failure),
            key_version_state=InstanceEncryptionKeysKeyStateAvailabilityKeyVersionStateEnum.from_proto(
                resource.key_version_state
            ),
        )


class InstanceEncryptionKeysKeyStateAvailabilityArray(object):
    @classmethod
    def to_proto(self, resources):
        if not resources:
            return resources
        return [
            InstanceEncryptionKeysKeyStateAvailability.to_proto(i) for i in resources
        ]

    @classmethod
    def from_proto(self, resources):
        return [
            InstanceEncryptionKeysKeyStateAvailability.from_proto(i) for i in resources
        ]


class InstancePreprocessCreateRecipe(object):
    def __init__(
        self,
        steps: list = None,
        honor_cancel_request: bool = None,
        ignore_recipe_after: int = None,
        verify_deadline_seconds_below: float = None,
        populate_operation_result: bool = None,
        readonly_recipe_start_time: str = None,
        resource_names_stored_in_clh_with_delay: list = None,
        delay_to_store_resources_in_clh_db_nanos: int = None,
    ):
        self.steps = steps
        self.honor_cancel_request = honor_cancel_request
        self.ignore_recipe_after = ignore_recipe_after
        self.verify_deadline_seconds_below = verify_deadline_seconds_below
        self.populate_operation_result = populate_operation_result
        self.readonly_recipe_start_time = readonly_recipe_start_time
        self.resource_names_stored_in_clh_with_delay = (
            resource_names_stored_in_clh_with_delay
        )
        self.delay_to_store_resources_in_clh_db_nanos = (
            delay_to_store_resources_in_clh_db_nanos
        )

    @classmethod
    def to_proto(self, resource):
        if not resource:
            return None

        res = instance_pb2.Tier2AlphaInstancePreprocessCreateRecipe()
        if InstancePreprocessCreateRecipeStepsArray.to_proto(resource.steps):
            res.steps.extend(
                InstancePreprocessCreateRecipeStepsArray.to_proto(resource.steps)
            )
        if Primitive.to_proto(resource.honor_cancel_request):
            res.honor_cancel_request = Primitive.to_proto(resource.honor_cancel_request)
        if Primitive.to_proto(resource.ignore_recipe_after):
            res.ignore_recipe_after = Primitive.to_proto(resource.ignore_recipe_after)
        if Primitive.to_proto(resource.verify_deadline_seconds_below):
            res.verify_deadline_seconds_below = Primitive.to_proto(
                resource.verify_deadline_seconds_below
            )
        if Primitive.to_proto(resource.populate_operation_result):
            res.populate_operation_result = Primitive.to_proto(
                resource.populate_operation_result
            )
        if Primitive.to_proto(resource.readonly_recipe_start_time):
            res.readonly_recipe_start_time = Primitive.to_proto(
                resource.readonly_recipe_start_time
            )
        if Primitive.to_proto(resource.resource_names_stored_in_clh_with_delay):
            res.resource_names_stored_in_clh_with_delay.extend(
                Primitive.to_proto(resource.resource_names_stored_in_clh_with_delay)
            )
        if Primitive.to_proto(resource.delay_to_store_resources_in_clh_db_nanos):
            res.delay_to_store_resources_in_clh_db_nanos = Primitive.to_proto(
                resource.delay_to_store_resources_in_clh_db_nanos
            )
        return res

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return None

        return InstancePreprocessCreateRecipe(
            steps=InstancePreprocessCreateRecipeStepsArray.from_proto(resource.steps),
            honor_cancel_request=Primitive.from_proto(resource.honor_cancel_request),
            ignore_recipe_after=Primitive.from_proto(resource.ignore_recipe_after),
            verify_deadline_seconds_below=Primitive.from_proto(
                resource.verify_deadline_seconds_below
            ),
            populate_operation_result=Primitive.from_proto(
                resource.populate_operation_result
            ),
            readonly_recipe_start_time=Primitive.from_proto(
                resource.readonly_recipe_start_time
            ),
            resource_names_stored_in_clh_with_delay=Primitive.from_proto(
                resource.resource_names_stored_in_clh_with_delay
            ),
            delay_to_store_resources_in_clh_db_nanos=Primitive.from_proto(
                resource.delay_to_store_resources_in_clh_db_nanos
            ),
        )


class InstancePreprocessCreateRecipeArray(object):
    @classmethod
    def to_proto(self, resources):
        if not resources:
            return resources
        return [InstancePreprocessCreateRecipe.to_proto(i) for i in resources]

    @classmethod
    def from_proto(self, resources):
        return [InstancePreprocessCreateRecipe.from_proto(i) for i in resources]


class InstancePreprocessCreateRecipeSteps(object):
    def __init__(
        self,
        relative_time: int = None,
        sleep_duration: int = None,
        action: str = None,
        status: dict = None,
        error_space: str = None,
        p4_service_account: str = None,
        resource_metadata_size: int = None,
        description: str = None,
        updated_repeat_operation_delay_sec: float = None,
        quota_request_deltas: list = None,
        preprocess_update: dict = None,
        public_operation_metadata: str = None,
        requested_tenant_project: dict = None,
        permissions_info: list = None,
        key_notifications_update: dict = None,
        clh_data_update_time: str = None,
        public_error_message: str = None,
    ):
        self.relative_time = relative_time
        self.sleep_duration = sleep_duration
        self.action = action
        self.status = status
        self.error_space = error_space
        self.p4_service_account = p4_service_account
        self.resource_metadata_size = resource_metadata_size
        self.description = description
        self.updated_repeat_operation_delay_sec = updated_repeat_operation_delay_sec
        self.quota_request_deltas = quota_request_deltas
        self.preprocess_update = preprocess_update
        self.public_operation_metadata = public_operation_metadata
        self.requested_tenant_project = requested_tenant_project
        self.permissions_info = permissions_info
        self.key_notifications_update = key_notifications_update
        self.clh_data_update_time = clh_data_update_time
        self.public_error_message = public_error_message

    @classmethod
    def to_proto(self, resource):
        if not resource:
            return None

        res = instance_pb2.Tier2AlphaInstancePreprocessCreateRecipeSteps()
        if Primitive.to_proto(resource.relative_time):
            res.relative_time = Primitive.to_proto(resource.relative_time)
        if Primitive.to_proto(resource.sleep_duration):
            res.sleep_duration = Primitive.to_proto(resource.sleep_duration)
        if InstancePreprocessCreateRecipeStepsActionEnum.to_proto(resource.action):
            res.action = InstancePreprocessCreateRecipeStepsActionEnum.to_proto(
                resource.action
            )
        if InstancePreprocessCreateRecipeStepsStatus.to_proto(resource.status):
            res.status.CopyFrom(
                InstancePreprocessCreateRecipeStepsStatus.to_proto(resource.status)
            )
        else:
            res.ClearField("status")
        if Primitive.to_proto(resource.error_space):
            res.error_space = Primitive.to_proto(resource.error_space)
        if Primitive.to_proto(resource.p4_service_account):
            res.p4_service_account = Primitive.to_proto(resource.p4_service_account)
        if Primitive.to_proto(resource.resource_metadata_size):
            res.resource_metadata_size = Primitive.to_proto(
                resource.resource_metadata_size
            )
        if Primitive.to_proto(resource.description):
            res.description = Primitive.to_proto(resource.description)
        if Primitive.to_proto(resource.updated_repeat_operation_delay_sec):
            res.updated_repeat_operation_delay_sec = Primitive.to_proto(
                resource.updated_repeat_operation_delay_sec
            )
        if InstancePreprocessCreateRecipeStepsQuotaRequestDeltasArray.to_proto(
            resource.quota_request_deltas
        ):
            res.quota_request_deltas.extend(
                InstancePreprocessCreateRecipeStepsQuotaRequestDeltasArray.to_proto(
                    resource.quota_request_deltas
                )
            )
        if InstancePreprocessCreateRecipeStepsPreprocessUpdate.to_proto(
            resource.preprocess_update
        ):
            res.preprocess_update.CopyFrom(
                InstancePreprocessCreateRecipeStepsPreprocessUpdate.to_proto(
                    resource.preprocess_update
                )
            )
        else:
            res.ClearField("preprocess_update")
        if Primitive.to_proto(resource.public_operation_metadata):
            res.public_operation_metadata = Primitive.to_proto(
                resource.public_operation_metadata
            )
        if InstancePreprocessCreateRecipeStepsRequestedTenantProject.to_proto(
            resource.requested_tenant_project
        ):
            res.requested_tenant_project.CopyFrom(
                InstancePreprocessCreateRecipeStepsRequestedTenantProject.to_proto(
                    resource.requested_tenant_project
                )
            )
        else:
            res.ClearField("requested_tenant_project")
        if InstancePreprocessCreateRecipeStepsPermissionsInfoArray.to_proto(
            resource.permissions_info
        ):
            res.permissions_info.extend(
                InstancePreprocessCreateRecipeStepsPermissionsInfoArray.to_proto(
                    resource.permissions_info
                )
            )
        if InstancePreprocessCreateRecipeStepsKeyNotificationsUpdate.to_proto(
            resource.key_notifications_update
        ):
            res.key_notifications_update.CopyFrom(
                InstancePreprocessCreateRecipeStepsKeyNotificationsUpdate.to_proto(
                    resource.key_notifications_update
                )
            )
        else:
            res.ClearField("key_notifications_update")
        if Primitive.to_proto(resource.clh_data_update_time):
            res.clh_data_update_time = Primitive.to_proto(resource.clh_data_update_time)
        if Primitive.to_proto(resource.public_error_message):
            res.public_error_message = Primitive.to_proto(resource.public_error_message)
        return res

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return None

        return InstancePreprocessCreateRecipeSteps(
            relative_time=Primitive.from_proto(resource.relative_time),
            sleep_duration=Primitive.from_proto(resource.sleep_duration),
            action=InstancePreprocessCreateRecipeStepsActionEnum.from_proto(
                resource.action
            ),
            status=InstancePreprocessCreateRecipeStepsStatus.from_proto(
                resource.status
            ),
            error_space=Primitive.from_proto(resource.error_space),
            p4_service_account=Primitive.from_proto(resource.p4_service_account),
            resource_metadata_size=Primitive.from_proto(
                resource.resource_metadata_size
            ),
            description=Primitive.from_proto(resource.description),
            updated_repeat_operation_delay_sec=Primitive.from_proto(
                resource.updated_repeat_operation_delay_sec
            ),
            quota_request_deltas=InstancePreprocessCreateRecipeStepsQuotaRequestDeltasArray.from_proto(
                resource.quota_request_deltas
            ),
            preprocess_update=InstancePreprocessCreateRecipeStepsPreprocessUpdate.from_proto(
                resource.preprocess_update
            ),
            public_operation_metadata=Primitive.from_proto(
                resource.public_operation_metadata
            ),
            requested_tenant_project=InstancePreprocessCreateRecipeStepsRequestedTenantProject.from_proto(
                resource.requested_tenant_project
            ),
            permissions_info=InstancePreprocessCreateRecipeStepsPermissionsInfoArray.from_proto(
                resource.permissions_info
            ),
            key_notifications_update=InstancePreprocessCreateRecipeStepsKeyNotificationsUpdate.from_proto(
                resource.key_notifications_update
            ),
            clh_data_update_time=Primitive.from_proto(resource.clh_data_update_time),
            public_error_message=Primitive.from_proto(resource.public_error_message),
        )


class InstancePreprocessCreateRecipeStepsArray(object):
    @classmethod
    def to_proto(self, resources):
        if not resources:
            return resources
        return [InstancePreprocessCreateRecipeSteps.to_proto(i) for i in resources]

    @classmethod
    def from_proto(self, resources):
        return [InstancePreprocessCreateRecipeSteps.from_proto(i) for i in resources]


class InstancePreprocessCreateRecipeStepsStatus(object):
    def __init__(self, code: int = None, message: str = None, details: list = None):
        self.code = code
        self.message = message
        self.details = details

    @classmethod
    def to_proto(self, resource):
        if not resource:
            return None

        res = instance_pb2.Tier2AlphaInstancePreprocessCreateRecipeStepsStatus()
        if Primitive.to_proto(resource.code):
            res.code = Primitive.to_proto(resource.code)
        if Primitive.to_proto(resource.message):
            res.message = Primitive.to_proto(resource.message)
        if InstancePreprocessCreateRecipeStepsStatusDetailsArray.to_proto(
            resource.details
        ):
            res.details.extend(
                InstancePreprocessCreateRecipeStepsStatusDetailsArray.to_proto(
                    resource.details
                )
            )
        return res

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return None

        return InstancePreprocessCreateRecipeStepsStatus(
            code=Primitive.from_proto(resource.code),
            message=Primitive.from_proto(resource.message),
            details=InstancePreprocessCreateRecipeStepsStatusDetailsArray.from_proto(
                resource.details
            ),
        )


class InstancePreprocessCreateRecipeStepsStatusArray(object):
    @classmethod
    def to_proto(self, resources):
        if not resources:
            return resources
        return [
            InstancePreprocessCreateRecipeStepsStatus.to_proto(i) for i in resources
        ]

    @classmethod
    def from_proto(self, resources):
        return [
            InstancePreprocessCreateRecipeStepsStatus.from_proto(i) for i in resources
        ]


class InstancePreprocessCreateRecipeStepsStatusDetails(object):
    def __init__(self, type_url: str = None, value: str = None):
        self.type_url = type_url
        self.value = value

    @classmethod
    def to_proto(self, resource):
        if not resource:
            return None

        res = instance_pb2.Tier2AlphaInstancePreprocessCreateRecipeStepsStatusDetails()
        if Primitive.to_proto(resource.type_url):
            res.type_url = Primitive.to_proto(resource.type_url)
        if Primitive.to_proto(resource.value):
            res.value = Primitive.to_proto(resource.value)
        return res

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return None

        return InstancePreprocessCreateRecipeStepsStatusDetails(
            type_url=Primitive.from_proto(resource.type_url),
            value=Primitive.from_proto(resource.value),
        )


class InstancePreprocessCreateRecipeStepsStatusDetailsArray(object):
    @classmethod
    def to_proto(self, resources):
        if not resources:
            return resources
        return [
            InstancePreprocessCreateRecipeStepsStatusDetails.to_proto(i)
            for i in resources
        ]

    @classmethod
    def from_proto(self, resources):
        return [
            InstancePreprocessCreateRecipeStepsStatusDetails.from_proto(i)
            for i in resources
        ]


class InstancePreprocessCreateRecipeStepsQuotaRequestDeltas(object):
    def __init__(
        self,
        metric_name: str = None,
        amount: int = None,
        quota_location_name: str = None,
    ):
        self.metric_name = metric_name
        self.amount = amount
        self.quota_location_name = quota_location_name

    @classmethod
    def to_proto(self, resource):
        if not resource:
            return None

        res = (
            instance_pb2.Tier2AlphaInstancePreprocessCreateRecipeStepsQuotaRequestDeltas()
        )
        if Primitive.to_proto(resource.metric_name):
            res.metric_name = Primitive.to_proto(resource.metric_name)
        if Primitive.to_proto(resource.amount):
            res.amount = Primitive.to_proto(resource.amount)
        if Primitive.to_proto(resource.quota_location_name):
            res.quota_location_name = Primitive.to_proto(resource.quota_location_name)
        return res

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return None

        return InstancePreprocessCreateRecipeStepsQuotaRequestDeltas(
            metric_name=Primitive.from_proto(resource.metric_name),
            amount=Primitive.from_proto(resource.amount),
            quota_location_name=Primitive.from_proto(resource.quota_location_name),
        )


class InstancePreprocessCreateRecipeStepsQuotaRequestDeltasArray(object):
    @classmethod
    def to_proto(self, resources):
        if not resources:
            return resources
        return [
            InstancePreprocessCreateRecipeStepsQuotaRequestDeltas.to_proto(i)
            for i in resources
        ]

    @classmethod
    def from_proto(self, resources):
        return [
            InstancePreprocessCreateRecipeStepsQuotaRequestDeltas.from_proto(i)
            for i in resources
        ]


class InstancePreprocessCreateRecipeStepsPreprocessUpdate(object):
    def __init__(
        self, latency_slo_bucket_name: str = None, public_operation_metadata: str = None
    ):
        self.latency_slo_bucket_name = latency_slo_bucket_name
        self.public_operation_metadata = public_operation_metadata

    @classmethod
    def to_proto(self, resource):
        if not resource:
            return None

        res = (
            instance_pb2.Tier2AlphaInstancePreprocessCreateRecipeStepsPreprocessUpdate()
        )
        if Primitive.to_proto(resource.latency_slo_bucket_name):
            res.latency_slo_bucket_name = Primitive.to_proto(
                resource.latency_slo_bucket_name
            )
        if Primitive.to_proto(resource.public_operation_metadata):
            res.public_operation_metadata = Primitive.to_proto(
                resource.public_operation_metadata
            )
        return res

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return None

        return InstancePreprocessCreateRecipeStepsPreprocessUpdate(
            latency_slo_bucket_name=Primitive.from_proto(
                resource.latency_slo_bucket_name
            ),
            public_operation_metadata=Primitive.from_proto(
                resource.public_operation_metadata
            ),
        )


class InstancePreprocessCreateRecipeStepsPreprocessUpdateArray(object):
    @classmethod
    def to_proto(self, resources):
        if not resources:
            return resources
        return [
            InstancePreprocessCreateRecipeStepsPreprocessUpdate.to_proto(i)
            for i in resources
        ]

    @classmethod
    def from_proto(self, resources):
        return [
            InstancePreprocessCreateRecipeStepsPreprocessUpdate.from_proto(i)
            for i in resources
        ]


class InstancePreprocessCreateRecipeStepsRequestedTenantProject(object):
    def __init__(self, tag: str = None, folder: str = None, scope: str = None):
        self.tag = tag
        self.folder = folder
        self.scope = scope

    @classmethod
    def to_proto(self, resource):
        if not resource:
            return None

        res = (
            instance_pb2.Tier2AlphaInstancePreprocessCreateRecipeStepsRequestedTenantProject()
        )
        if Primitive.to_proto(resource.tag):
            res.tag = Primitive.to_proto(resource.tag)
        if Primitive.to_proto(resource.folder):
            res.folder = Primitive.to_proto(resource.folder)
        if InstancePreprocessCreateRecipeStepsRequestedTenantProjectScopeEnum.to_proto(
            resource.scope
        ):
            res.scope = InstancePreprocessCreateRecipeStepsRequestedTenantProjectScopeEnum.to_proto(
                resource.scope
            )
        return res

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return None

        return InstancePreprocessCreateRecipeStepsRequestedTenantProject(
            tag=Primitive.from_proto(resource.tag),
            folder=Primitive.from_proto(resource.folder),
            scope=InstancePreprocessCreateRecipeStepsRequestedTenantProjectScopeEnum.from_proto(
                resource.scope
            ),
        )


class InstancePreprocessCreateRecipeStepsRequestedTenantProjectArray(object):
    @classmethod
    def to_proto(self, resources):
        if not resources:
            return resources
        return [
            InstancePreprocessCreateRecipeStepsRequestedTenantProject.to_proto(i)
            for i in resources
        ]

    @classmethod
    def from_proto(self, resources):
        return [
            InstancePreprocessCreateRecipeStepsRequestedTenantProject.from_proto(i)
            for i in resources
        ]


class InstancePreprocessCreateRecipeStepsPermissionsInfo(object):
    def __init__(
        self,
        policy_name: dict = None,
        iam_permissions: list = None,
        resource_path: str = None,
        api_attrs: dict = None,
        policy_name_mode: str = None,
        resource: dict = None,
    ):
        self.policy_name = policy_name
        self.iam_permissions = iam_permissions
        self.resource_path = resource_path
        self.api_attrs = api_attrs
        self.policy_name_mode = policy_name_mode
        self.resource = resource

    @classmethod
    def to_proto(self, resource):
        if not resource:
            return None

        res = (
            instance_pb2.Tier2AlphaInstancePreprocessCreateRecipeStepsPermissionsInfo()
        )
        if InstancePreprocessCreateRecipeStepsPermissionsInfoPolicyName.to_proto(
            resource.policy_name
        ):
            res.policy_name.CopyFrom(
                InstancePreprocessCreateRecipeStepsPermissionsInfoPolicyName.to_proto(
                    resource.policy_name
                )
            )
        else:
            res.ClearField("policy_name")
        if InstancePreprocessCreateRecipeStepsPermissionsInfoIamPermissionsArray.to_proto(
            resource.iam_permissions
        ):
            res.iam_permissions.extend(
                InstancePreprocessCreateRecipeStepsPermissionsInfoIamPermissionsArray.to_proto(
                    resource.iam_permissions
                )
            )
        if Primitive.to_proto(resource.resource_path):
            res.resource_path = Primitive.to_proto(resource.resource_path)
        if InstanceGoogleprotobufstruct.to_proto(resource.api_attrs):
            res.api_attrs.CopyFrom(
                InstanceGoogleprotobufstruct.to_proto(resource.api_attrs)
            )
        else:
            res.ClearField("api_attrs")
        if InstancePreprocessCreateRecipeStepsPermissionsInfoPolicyNameModeEnum.to_proto(
            resource.policy_name_mode
        ):
            res.policy_name_mode = InstancePreprocessCreateRecipeStepsPermissionsInfoPolicyNameModeEnum.to_proto(
                resource.policy_name_mode
            )
        if InstancePreprocessCreateRecipeStepsPermissionsInfoResource.to_proto(
            resource.resource
        ):
            res.resource.CopyFrom(
                InstancePreprocessCreateRecipeStepsPermissionsInfoResource.to_proto(
                    resource.resource
                )
            )
        else:
            res.ClearField("resource")
        return res

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return None

        return InstancePreprocessCreateRecipeStepsPermissionsInfo(
            policy_name=InstancePreprocessCreateRecipeStepsPermissionsInfoPolicyName.from_proto(
                resource.policy_name
            ),
            iam_permissions=InstancePreprocessCreateRecipeStepsPermissionsInfoIamPermissionsArray.from_proto(
                resource.iam_permissions
            ),
            resource_path=Primitive.from_proto(resource.resource_path),
            api_attrs=InstanceGoogleprotobufstruct.from_proto(resource.api_attrs),
            policy_name_mode=InstancePreprocessCreateRecipeStepsPermissionsInfoPolicyNameModeEnum.from_proto(
                resource.policy_name_mode
            ),
            resource=InstancePreprocessCreateRecipeStepsPermissionsInfoResource.from_proto(
                resource.resource
            ),
        )


class InstancePreprocessCreateRecipeStepsPermissionsInfoArray(object):
    @classmethod
    def to_proto(self, resources):
        if not resources:
            return resources
        return [
            InstancePreprocessCreateRecipeStepsPermissionsInfo.to_proto(i)
            for i in resources
        ]

    @classmethod
    def from_proto(self, resources):
        return [
            InstancePreprocessCreateRecipeStepsPermissionsInfo.from_proto(i)
            for i in resources
        ]


class InstancePreprocessCreateRecipeStepsPermissionsInfoPolicyName(object):
    def __init__(self, type: str = None, id: str = None, region: str = None):
        self.type = type
        self.id = id
        self.region = region

    @classmethod
    def to_proto(self, resource):
        if not resource:
            return None

        res = (
            instance_pb2.Tier2AlphaInstancePreprocessCreateRecipeStepsPermissionsInfoPolicyName()
        )
        if Primitive.to_proto(resource.type):
            res.type = Primitive.to_proto(resource.type)
        if Primitive.to_proto(resource.id):
            res.id = Primitive.to_proto(resource.id)
        if Primitive.to_proto(resource.region):
            res.region = Primitive.to_proto(resource.region)
        return res

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return None

        return InstancePreprocessCreateRecipeStepsPermissionsInfoPolicyName(
            type=Primitive.from_proto(resource.type),
            id=Primitive.from_proto(resource.id),
            region=Primitive.from_proto(resource.region),
        )


class InstancePreprocessCreateRecipeStepsPermissionsInfoPolicyNameArray(object):
    @classmethod
    def to_proto(self, resources):
        if not resources:
            return resources
        return [
            InstancePreprocessCreateRecipeStepsPermissionsInfoPolicyName.to_proto(i)
            for i in resources
        ]

    @classmethod
    def from_proto(self, resources):
        return [
            InstancePreprocessCreateRecipeStepsPermissionsInfoPolicyName.from_proto(i)
            for i in resources
        ]


class InstancePreprocessCreateRecipeStepsPermissionsInfoIamPermissions(object):
    def __init__(self, permission: str = None):
        self.permission = permission

    @classmethod
    def to_proto(self, resource):
        if not resource:
            return None

        res = (
            instance_pb2.Tier2AlphaInstancePreprocessCreateRecipeStepsPermissionsInfoIamPermissions()
        )
        if Primitive.to_proto(resource.permission):
            res.permission = Primitive.to_proto(resource.permission)
        return res

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return None

        return InstancePreprocessCreateRecipeStepsPermissionsInfoIamPermissions(
            permission=Primitive.from_proto(resource.permission),
        )


class InstancePreprocessCreateRecipeStepsPermissionsInfoIamPermissionsArray(object):
    @classmethod
    def to_proto(self, resources):
        if not resources:
            return resources
        return [
            InstancePreprocessCreateRecipeStepsPermissionsInfoIamPermissions.to_proto(i)
            for i in resources
        ]

    @classmethod
    def from_proto(self, resources):
        return [
            InstancePreprocessCreateRecipeStepsPermissionsInfoIamPermissions.from_proto(
                i
            )
            for i in resources
        ]


class InstanceGoogleprotobufstruct(object):
    @classmethod
    def to_proto(self, resource):
        if not resource:
            return None

        res = instance_pb2.Tier2AlphaInstanceGoogleprotobufstruct()
        return res

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return None

        return InstanceGoogleprotobufstruct()


class InstanceGoogleprotobufstructArray(object):
    @classmethod
    def to_proto(self, resources):
        if not resources:
            return resources
        return [InstanceGoogleprotobufstruct.to_proto(i) for i in resources]

    @classmethod
    def from_proto(self, resources):
        return [InstanceGoogleprotobufstruct.from_proto(i) for i in resources]


class InstancePreprocessCreateRecipeStepsPermissionsInfoResource(object):
    def __init__(
        self,
        name: str = None,
        type: str = None,
        service: str = None,
        labels: dict = None,
    ):
        self.name = name
        self.type = type
        self.service = service
        self.labels = labels

    @classmethod
    def to_proto(self, resource):
        if not resource:
            return None

        res = (
            instance_pb2.Tier2AlphaInstancePreprocessCreateRecipeStepsPermissionsInfoResource()
        )
        if Primitive.to_proto(resource.name):
            res.name = Primitive.to_proto(resource.name)
        if Primitive.to_proto(resource.type):
            res.type = Primitive.to_proto(resource.type)
        if Primitive.to_proto(resource.service):
            res.service = Primitive.to_proto(resource.service)
        if Primitive.to_proto(resource.labels):
            res.labels = Primitive.to_proto(resource.labels)
        return res

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return None

        return InstancePreprocessCreateRecipeStepsPermissionsInfoResource(
            name=Primitive.from_proto(resource.name),
            type=Primitive.from_proto(resource.type),
            service=Primitive.from_proto(resource.service),
            labels=Primitive.from_proto(resource.labels),
        )


class InstancePreprocessCreateRecipeStepsPermissionsInfoResourceArray(object):
    @classmethod
    def to_proto(self, resources):
        if not resources:
            return resources
        return [
            InstancePreprocessCreateRecipeStepsPermissionsInfoResource.to_proto(i)
            for i in resources
        ]

    @classmethod
    def from_proto(self, resources):
        return [
            InstancePreprocessCreateRecipeStepsPermissionsInfoResource.from_proto(i)
            for i in resources
        ]


class InstancePreprocessCreateRecipeStepsKeyNotificationsUpdate(object):
    def __init__(self, key_notifications_info: dict = None):
        self.key_notifications_info = key_notifications_info

    @classmethod
    def to_proto(self, resource):
        if not resource:
            return None

        res = (
            instance_pb2.Tier2AlphaInstancePreprocessCreateRecipeStepsKeyNotificationsUpdate()
        )
        if InstancePreprocessCreateRecipeStepsKeyNotificationsUpdateKeyNotificationsInfo.to_proto(
            resource.key_notifications_info
        ):
            res.key_notifications_info.CopyFrom(
                InstancePreprocessCreateRecipeStepsKeyNotificationsUpdateKeyNotificationsInfo.to_proto(
                    resource.key_notifications_info
                )
            )
        else:
            res.ClearField("key_notifications_info")
        return res

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return None

        return InstancePreprocessCreateRecipeStepsKeyNotificationsUpdate(
            key_notifications_info=InstancePreprocessCreateRecipeStepsKeyNotificationsUpdateKeyNotificationsInfo.from_proto(
                resource.key_notifications_info
            ),
        )


class InstancePreprocessCreateRecipeStepsKeyNotificationsUpdateArray(object):
    @classmethod
    def to_proto(self, resources):
        if not resources:
            return resources
        return [
            InstancePreprocessCreateRecipeStepsKeyNotificationsUpdate.to_proto(i)
            for i in resources
        ]

    @classmethod
    def from_proto(self, resources):
        return [
            InstancePreprocessCreateRecipeStepsKeyNotificationsUpdate.from_proto(i)
            for i in resources
        ]


class InstancePreprocessCreateRecipeStepsKeyNotificationsUpdateKeyNotificationsInfo(
    object
):
    def __init__(
        self,
        data_version: int = None,
        delegate: str = None,
        key_notification_configs: list = None,
    ):
        self.data_version = data_version
        self.delegate = delegate
        self.key_notification_configs = key_notification_configs

    @classmethod
    def to_proto(self, resource):
        if not resource:
            return None

        res = (
            instance_pb2.Tier2AlphaInstancePreprocessCreateRecipeStepsKeyNotificationsUpdateKeyNotificationsInfo()
        )
        if Primitive.to_proto(resource.data_version):
            res.data_version = Primitive.to_proto(resource.data_version)
        if Primitive.to_proto(resource.delegate):
            res.delegate = Primitive.to_proto(resource.delegate)
        if InstancePreprocessCreateRecipeStepsKeyNotificationsUpdateKeyNotificationsInfoKeyNotificationConfigsArray.to_proto(
            resource.key_notification_configs
        ):
            res.key_notification_configs.extend(
                InstancePreprocessCreateRecipeStepsKeyNotificationsUpdateKeyNotificationsInfoKeyNotificationConfigsArray.to_proto(
                    resource.key_notification_configs
                )
            )
        return res

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return None

        return InstancePreprocessCreateRecipeStepsKeyNotificationsUpdateKeyNotificationsInfo(
            data_version=Primitive.from_proto(resource.data_version),
            delegate=Primitive.from_proto(resource.delegate),
            key_notification_configs=InstancePreprocessCreateRecipeStepsKeyNotificationsUpdateKeyNotificationsInfoKeyNotificationConfigsArray.from_proto(
                resource.key_notification_configs
            ),
        )


class InstancePreprocessCreateRecipeStepsKeyNotificationsUpdateKeyNotificationsInfoArray(
    object
):
    @classmethod
    def to_proto(self, resources):
        if not resources:
            return resources
        return [
            InstancePreprocessCreateRecipeStepsKeyNotificationsUpdateKeyNotificationsInfo.to_proto(
                i
            )
            for i in resources
        ]

    @classmethod
    def from_proto(self, resources):
        return [
            InstancePreprocessCreateRecipeStepsKeyNotificationsUpdateKeyNotificationsInfo.from_proto(
                i
            )
            for i in resources
        ]


class InstancePreprocessCreateRecipeStepsKeyNotificationsUpdateKeyNotificationsInfoKeyNotificationConfigs(
    object
):
    def __init__(
        self,
        key_or_version_name: str = None,
        grant: str = None,
        delegator_gaia_id: int = None,
    ):
        self.key_or_version_name = key_or_version_name
        self.grant = grant
        self.delegator_gaia_id = delegator_gaia_id

    @classmethod
    def to_proto(self, resource):
        if not resource:
            return None

        res = (
            instance_pb2.Tier2AlphaInstancePreprocessCreateRecipeStepsKeyNotificationsUpdateKeyNotificationsInfoKeyNotificationConfigs()
        )
        if Primitive.to_proto(resource.key_or_version_name):
            res.key_or_version_name = Primitive.to_proto(resource.key_or_version_name)
        if Primitive.to_proto(resource.grant):
            res.grant = Primitive.to_proto(resource.grant)
        if Primitive.to_proto(resource.delegator_gaia_id):
            res.delegator_gaia_id = Primitive.to_proto(resource.delegator_gaia_id)
        return res

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return None

        return InstancePreprocessCreateRecipeStepsKeyNotificationsUpdateKeyNotificationsInfoKeyNotificationConfigs(
            key_or_version_name=Primitive.from_proto(resource.key_or_version_name),
            grant=Primitive.from_proto(resource.grant),
            delegator_gaia_id=Primitive.from_proto(resource.delegator_gaia_id),
        )


class InstancePreprocessCreateRecipeStepsKeyNotificationsUpdateKeyNotificationsInfoKeyNotificationConfigsArray(
    object
):
    @classmethod
    def to_proto(self, resources):
        if not resources:
            return resources
        return [
            InstancePreprocessCreateRecipeStepsKeyNotificationsUpdateKeyNotificationsInfoKeyNotificationConfigs.to_proto(
                i
            )
            for i in resources
        ]

    @classmethod
    def from_proto(self, resources):
        return [
            InstancePreprocessCreateRecipeStepsKeyNotificationsUpdateKeyNotificationsInfoKeyNotificationConfigs.from_proto(
                i
            )
            for i in resources
        ]


class InstanceCreateRecipe(object):
    def __init__(
        self,
        steps: list = None,
        honor_cancel_request: bool = None,
        ignore_recipe_after: int = None,
        verify_deadline_seconds_below: float = None,
        populate_operation_result: bool = None,
        readonly_recipe_start_time: str = None,
        resource_names_stored_in_clh_with_delay: list = None,
        delay_to_store_resources_in_clh_db_nanos: int = None,
    ):
        self.steps = steps
        self.honor_cancel_request = honor_cancel_request
        self.ignore_recipe_after = ignore_recipe_after
        self.verify_deadline_seconds_below = verify_deadline_seconds_below
        self.populate_operation_result = populate_operation_result
        self.readonly_recipe_start_time = readonly_recipe_start_time
        self.resource_names_stored_in_clh_with_delay = (
            resource_names_stored_in_clh_with_delay
        )
        self.delay_to_store_resources_in_clh_db_nanos = (
            delay_to_store_resources_in_clh_db_nanos
        )

    @classmethod
    def to_proto(self, resource):
        if not resource:
            return None

        res = instance_pb2.Tier2AlphaInstanceCreateRecipe()
        if InstanceCreateRecipeStepsArray.to_proto(resource.steps):
            res.steps.extend(InstanceCreateRecipeStepsArray.to_proto(resource.steps))
        if Primitive.to_proto(resource.honor_cancel_request):
            res.honor_cancel_request = Primitive.to_proto(resource.honor_cancel_request)
        if Primitive.to_proto(resource.ignore_recipe_after):
            res.ignore_recipe_after = Primitive.to_proto(resource.ignore_recipe_after)
        if Primitive.to_proto(resource.verify_deadline_seconds_below):
            res.verify_deadline_seconds_below = Primitive.to_proto(
                resource.verify_deadline_seconds_below
            )
        if Primitive.to_proto(resource.populate_operation_result):
            res.populate_operation_result = Primitive.to_proto(
                resource.populate_operation_result
            )
        if Primitive.to_proto(resource.readonly_recipe_start_time):
            res.readonly_recipe_start_time = Primitive.to_proto(
                resource.readonly_recipe_start_time
            )
        if Primitive.to_proto(resource.resource_names_stored_in_clh_with_delay):
            res.resource_names_stored_in_clh_with_delay.extend(
                Primitive.to_proto(resource.resource_names_stored_in_clh_with_delay)
            )
        if Primitive.to_proto(resource.delay_to_store_resources_in_clh_db_nanos):
            res.delay_to_store_resources_in_clh_db_nanos = Primitive.to_proto(
                resource.delay_to_store_resources_in_clh_db_nanos
            )
        return res

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return None

        return InstanceCreateRecipe(
            steps=InstanceCreateRecipeStepsArray.from_proto(resource.steps),
            honor_cancel_request=Primitive.from_proto(resource.honor_cancel_request),
            ignore_recipe_after=Primitive.from_proto(resource.ignore_recipe_after),
            verify_deadline_seconds_below=Primitive.from_proto(
                resource.verify_deadline_seconds_below
            ),
            populate_operation_result=Primitive.from_proto(
                resource.populate_operation_result
            ),
            readonly_recipe_start_time=Primitive.from_proto(
                resource.readonly_recipe_start_time
            ),
            resource_names_stored_in_clh_with_delay=Primitive.from_proto(
                resource.resource_names_stored_in_clh_with_delay
            ),
            delay_to_store_resources_in_clh_db_nanos=Primitive.from_proto(
                resource.delay_to_store_resources_in_clh_db_nanos
            ),
        )


class InstanceCreateRecipeArray(object):
    @classmethod
    def to_proto(self, resources):
        if not resources:
            return resources
        return [InstanceCreateRecipe.to_proto(i) for i in resources]

    @classmethod
    def from_proto(self, resources):
        return [InstanceCreateRecipe.from_proto(i) for i in resources]


class InstanceCreateRecipeSteps(object):
    def __init__(
        self,
        relative_time: int = None,
        sleep_duration: int = None,
        action: str = None,
        status: dict = None,
        error_space: str = None,
        p4_service_account: str = None,
        resource_metadata_size: int = None,
        description: str = None,
        updated_repeat_operation_delay_sec: float = None,
        quota_request_deltas: list = None,
        preprocess_update: dict = None,
        public_operation_metadata: str = None,
        requested_tenant_project: dict = None,
        permissions_info: list = None,
        key_notifications_update: dict = None,
        clh_data_update_time: str = None,
        public_error_message: str = None,
    ):
        self.relative_time = relative_time
        self.sleep_duration = sleep_duration
        self.action = action
        self.status = status
        self.error_space = error_space
        self.p4_service_account = p4_service_account
        self.resource_metadata_size = resource_metadata_size
        self.description = description
        self.updated_repeat_operation_delay_sec = updated_repeat_operation_delay_sec
        self.quota_request_deltas = quota_request_deltas
        self.preprocess_update = preprocess_update
        self.public_operation_metadata = public_operation_metadata
        self.requested_tenant_project = requested_tenant_project
        self.permissions_info = permissions_info
        self.key_notifications_update = key_notifications_update
        self.clh_data_update_time = clh_data_update_time
        self.public_error_message = public_error_message

    @classmethod
    def to_proto(self, resource):
        if not resource:
            return None

        res = instance_pb2.Tier2AlphaInstanceCreateRecipeSteps()
        if Primitive.to_proto(resource.relative_time):
            res.relative_time = Primitive.to_proto(resource.relative_time)
        if Primitive.to_proto(resource.sleep_duration):
            res.sleep_duration = Primitive.to_proto(resource.sleep_duration)
        if InstanceCreateRecipeStepsActionEnum.to_proto(resource.action):
            res.action = InstanceCreateRecipeStepsActionEnum.to_proto(resource.action)
        if InstanceCreateRecipeStepsStatus.to_proto(resource.status):
            res.status.CopyFrom(
                InstanceCreateRecipeStepsStatus.to_proto(resource.status)
            )
        else:
            res.ClearField("status")
        if Primitive.to_proto(resource.error_space):
            res.error_space = Primitive.to_proto(resource.error_space)
        if Primitive.to_proto(resource.p4_service_account):
            res.p4_service_account = Primitive.to_proto(resource.p4_service_account)
        if Primitive.to_proto(resource.resource_metadata_size):
            res.resource_metadata_size = Primitive.to_proto(
                resource.resource_metadata_size
            )
        if Primitive.to_proto(resource.description):
            res.description = Primitive.to_proto(resource.description)
        if Primitive.to_proto(resource.updated_repeat_operation_delay_sec):
            res.updated_repeat_operation_delay_sec = Primitive.to_proto(
                resource.updated_repeat_operation_delay_sec
            )
        if InstanceCreateRecipeStepsQuotaRequestDeltasArray.to_proto(
            resource.quota_request_deltas
        ):
            res.quota_request_deltas.extend(
                InstanceCreateRecipeStepsQuotaRequestDeltasArray.to_proto(
                    resource.quota_request_deltas
                )
            )
        if InstanceCreateRecipeStepsPreprocessUpdate.to_proto(
            resource.preprocess_update
        ):
            res.preprocess_update.CopyFrom(
                InstanceCreateRecipeStepsPreprocessUpdate.to_proto(
                    resource.preprocess_update
                )
            )
        else:
            res.ClearField("preprocess_update")
        if Primitive.to_proto(resource.public_operation_metadata):
            res.public_operation_metadata = Primitive.to_proto(
                resource.public_operation_metadata
            )
        if InstanceCreateRecipeStepsRequestedTenantProject.to_proto(
            resource.requested_tenant_project
        ):
            res.requested_tenant_project.CopyFrom(
                InstanceCreateRecipeStepsRequestedTenantProject.to_proto(
                    resource.requested_tenant_project
                )
            )
        else:
            res.ClearField("requested_tenant_project")
        if InstanceCreateRecipeStepsPermissionsInfoArray.to_proto(
            resource.permissions_info
        ):
            res.permissions_info.extend(
                InstanceCreateRecipeStepsPermissionsInfoArray.to_proto(
                    resource.permissions_info
                )
            )
        if InstanceCreateRecipeStepsKeyNotificationsUpdate.to_proto(
            resource.key_notifications_update
        ):
            res.key_notifications_update.CopyFrom(
                InstanceCreateRecipeStepsKeyNotificationsUpdate.to_proto(
                    resource.key_notifications_update
                )
            )
        else:
            res.ClearField("key_notifications_update")
        if Primitive.to_proto(resource.clh_data_update_time):
            res.clh_data_update_time = Primitive.to_proto(resource.clh_data_update_time)
        if Primitive.to_proto(resource.public_error_message):
            res.public_error_message = Primitive.to_proto(resource.public_error_message)
        return res

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return None

        return InstanceCreateRecipeSteps(
            relative_time=Primitive.from_proto(resource.relative_time),
            sleep_duration=Primitive.from_proto(resource.sleep_duration),
            action=InstanceCreateRecipeStepsActionEnum.from_proto(resource.action),
            status=InstanceCreateRecipeStepsStatus.from_proto(resource.status),
            error_space=Primitive.from_proto(resource.error_space),
            p4_service_account=Primitive.from_proto(resource.p4_service_account),
            resource_metadata_size=Primitive.from_proto(
                resource.resource_metadata_size
            ),
            description=Primitive.from_proto(resource.description),
            updated_repeat_operation_delay_sec=Primitive.from_proto(
                resource.updated_repeat_operation_delay_sec
            ),
            quota_request_deltas=InstanceCreateRecipeStepsQuotaRequestDeltasArray.from_proto(
                resource.quota_request_deltas
            ),
            preprocess_update=InstanceCreateRecipeStepsPreprocessUpdate.from_proto(
                resource.preprocess_update
            ),
            public_operation_metadata=Primitive.from_proto(
                resource.public_operation_metadata
            ),
            requested_tenant_project=InstanceCreateRecipeStepsRequestedTenantProject.from_proto(
                resource.requested_tenant_project
            ),
            permissions_info=InstanceCreateRecipeStepsPermissionsInfoArray.from_proto(
                resource.permissions_info
            ),
            key_notifications_update=InstanceCreateRecipeStepsKeyNotificationsUpdate.from_proto(
                resource.key_notifications_update
            ),
            clh_data_update_time=Primitive.from_proto(resource.clh_data_update_time),
            public_error_message=Primitive.from_proto(resource.public_error_message),
        )


class InstanceCreateRecipeStepsArray(object):
    @classmethod
    def to_proto(self, resources):
        if not resources:
            return resources
        return [InstanceCreateRecipeSteps.to_proto(i) for i in resources]

    @classmethod
    def from_proto(self, resources):
        return [InstanceCreateRecipeSteps.from_proto(i) for i in resources]


class InstanceCreateRecipeStepsStatus(object):
    def __init__(self, code: int = None, message: str = None, details: list = None):
        self.code = code
        self.message = message
        self.details = details

    @classmethod
    def to_proto(self, resource):
        if not resource:
            return None

        res = instance_pb2.Tier2AlphaInstanceCreateRecipeStepsStatus()
        if Primitive.to_proto(resource.code):
            res.code = Primitive.to_proto(resource.code)
        if Primitive.to_proto(resource.message):
            res.message = Primitive.to_proto(resource.message)
        if InstanceCreateRecipeStepsStatusDetailsArray.to_proto(resource.details):
            res.details.extend(
                InstanceCreateRecipeStepsStatusDetailsArray.to_proto(resource.details)
            )
        return res

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return None

        return InstanceCreateRecipeStepsStatus(
            code=Primitive.from_proto(resource.code),
            message=Primitive.from_proto(resource.message),
            details=InstanceCreateRecipeStepsStatusDetailsArray.from_proto(
                resource.details
            ),
        )


class InstanceCreateRecipeStepsStatusArray(object):
    @classmethod
    def to_proto(self, resources):
        if not resources:
            return resources
        return [InstanceCreateRecipeStepsStatus.to_proto(i) for i in resources]

    @classmethod
    def from_proto(self, resources):
        return [InstanceCreateRecipeStepsStatus.from_proto(i) for i in resources]


class InstanceCreateRecipeStepsStatusDetails(object):
    def __init__(self, type_url: str = None, value: str = None):
        self.type_url = type_url
        self.value = value

    @classmethod
    def to_proto(self, resource):
        if not resource:
            return None

        res = instance_pb2.Tier2AlphaInstanceCreateRecipeStepsStatusDetails()
        if Primitive.to_proto(resource.type_url):
            res.type_url = Primitive.to_proto(resource.type_url)
        if Primitive.to_proto(resource.value):
            res.value = Primitive.to_proto(resource.value)
        return res

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return None

        return InstanceCreateRecipeStepsStatusDetails(
            type_url=Primitive.from_proto(resource.type_url),
            value=Primitive.from_proto(resource.value),
        )


class InstanceCreateRecipeStepsStatusDetailsArray(object):
    @classmethod
    def to_proto(self, resources):
        if not resources:
            return resources
        return [InstanceCreateRecipeStepsStatusDetails.to_proto(i) for i in resources]

    @classmethod
    def from_proto(self, resources):
        return [InstanceCreateRecipeStepsStatusDetails.from_proto(i) for i in resources]


class InstanceCreateRecipeStepsQuotaRequestDeltas(object):
    def __init__(
        self,
        metric_name: str = None,
        amount: int = None,
        quota_location_name: str = None,
    ):
        self.metric_name = metric_name
        self.amount = amount
        self.quota_location_name = quota_location_name

    @classmethod
    def to_proto(self, resource):
        if not resource:
            return None

        res = instance_pb2.Tier2AlphaInstanceCreateRecipeStepsQuotaRequestDeltas()
        if Primitive.to_proto(resource.metric_name):
            res.metric_name = Primitive.to_proto(resource.metric_name)
        if Primitive.to_proto(resource.amount):
            res.amount = Primitive.to_proto(resource.amount)
        if Primitive.to_proto(resource.quota_location_name):
            res.quota_location_name = Primitive.to_proto(resource.quota_location_name)
        return res

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return None

        return InstanceCreateRecipeStepsQuotaRequestDeltas(
            metric_name=Primitive.from_proto(resource.metric_name),
            amount=Primitive.from_proto(resource.amount),
            quota_location_name=Primitive.from_proto(resource.quota_location_name),
        )


class InstanceCreateRecipeStepsQuotaRequestDeltasArray(object):
    @classmethod
    def to_proto(self, resources):
        if not resources:
            return resources
        return [
            InstanceCreateRecipeStepsQuotaRequestDeltas.to_proto(i) for i in resources
        ]

    @classmethod
    def from_proto(self, resources):
        return [
            InstanceCreateRecipeStepsQuotaRequestDeltas.from_proto(i) for i in resources
        ]


class InstanceCreateRecipeStepsPreprocessUpdate(object):
    def __init__(
        self, latency_slo_bucket_name: str = None, public_operation_metadata: str = None
    ):
        self.latency_slo_bucket_name = latency_slo_bucket_name
        self.public_operation_metadata = public_operation_metadata

    @classmethod
    def to_proto(self, resource):
        if not resource:
            return None

        res = instance_pb2.Tier2AlphaInstanceCreateRecipeStepsPreprocessUpdate()
        if Primitive.to_proto(resource.latency_slo_bucket_name):
            res.latency_slo_bucket_name = Primitive.to_proto(
                resource.latency_slo_bucket_name
            )
        if Primitive.to_proto(resource.public_operation_metadata):
            res.public_operation_metadata = Primitive.to_proto(
                resource.public_operation_metadata
            )
        return res

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return None

        return InstanceCreateRecipeStepsPreprocessUpdate(
            latency_slo_bucket_name=Primitive.from_proto(
                resource.latency_slo_bucket_name
            ),
            public_operation_metadata=Primitive.from_proto(
                resource.public_operation_metadata
            ),
        )


class InstanceCreateRecipeStepsPreprocessUpdateArray(object):
    @classmethod
    def to_proto(self, resources):
        if not resources:
            return resources
        return [
            InstanceCreateRecipeStepsPreprocessUpdate.to_proto(i) for i in resources
        ]

    @classmethod
    def from_proto(self, resources):
        return [
            InstanceCreateRecipeStepsPreprocessUpdate.from_proto(i) for i in resources
        ]


class InstanceCreateRecipeStepsRequestedTenantProject(object):
    def __init__(self, tag: str = None, folder: str = None, scope: str = None):
        self.tag = tag
        self.folder = folder
        self.scope = scope

    @classmethod
    def to_proto(self, resource):
        if not resource:
            return None

        res = instance_pb2.Tier2AlphaInstanceCreateRecipeStepsRequestedTenantProject()
        if Primitive.to_proto(resource.tag):
            res.tag = Primitive.to_proto(resource.tag)
        if Primitive.to_proto(resource.folder):
            res.folder = Primitive.to_proto(resource.folder)
        if InstanceCreateRecipeStepsRequestedTenantProjectScopeEnum.to_proto(
            resource.scope
        ):
            res.scope = InstanceCreateRecipeStepsRequestedTenantProjectScopeEnum.to_proto(
                resource.scope
            )
        return res

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return None

        return InstanceCreateRecipeStepsRequestedTenantProject(
            tag=Primitive.from_proto(resource.tag),
            folder=Primitive.from_proto(resource.folder),
            scope=InstanceCreateRecipeStepsRequestedTenantProjectScopeEnum.from_proto(
                resource.scope
            ),
        )


class InstanceCreateRecipeStepsRequestedTenantProjectArray(object):
    @classmethod
    def to_proto(self, resources):
        if not resources:
            return resources
        return [
            InstanceCreateRecipeStepsRequestedTenantProject.to_proto(i)
            for i in resources
        ]

    @classmethod
    def from_proto(self, resources):
        return [
            InstanceCreateRecipeStepsRequestedTenantProject.from_proto(i)
            for i in resources
        ]


class InstanceCreateRecipeStepsPermissionsInfo(object):
    def __init__(
        self,
        policy_name: dict = None,
        iam_permissions: list = None,
        resource_path: str = None,
        api_attrs: dict = None,
        policy_name_mode: str = None,
        resource: dict = None,
    ):
        self.policy_name = policy_name
        self.iam_permissions = iam_permissions
        self.resource_path = resource_path
        self.api_attrs = api_attrs
        self.policy_name_mode = policy_name_mode
        self.resource = resource

    @classmethod
    def to_proto(self, resource):
        if not resource:
            return None

        res = instance_pb2.Tier2AlphaInstanceCreateRecipeStepsPermissionsInfo()
        if InstanceCreateRecipeStepsPermissionsInfoPolicyName.to_proto(
            resource.policy_name
        ):
            res.policy_name.CopyFrom(
                InstanceCreateRecipeStepsPermissionsInfoPolicyName.to_proto(
                    resource.policy_name
                )
            )
        else:
            res.ClearField("policy_name")
        if InstanceCreateRecipeStepsPermissionsInfoIamPermissionsArray.to_proto(
            resource.iam_permissions
        ):
            res.iam_permissions.extend(
                InstanceCreateRecipeStepsPermissionsInfoIamPermissionsArray.to_proto(
                    resource.iam_permissions
                )
            )
        if Primitive.to_proto(resource.resource_path):
            res.resource_path = Primitive.to_proto(resource.resource_path)
        if InstanceGoogleprotobufstruct.to_proto(resource.api_attrs):
            res.api_attrs.CopyFrom(
                InstanceGoogleprotobufstruct.to_proto(resource.api_attrs)
            )
        else:
            res.ClearField("api_attrs")
        if InstanceCreateRecipeStepsPermissionsInfoPolicyNameModeEnum.to_proto(
            resource.policy_name_mode
        ):
            res.policy_name_mode = InstanceCreateRecipeStepsPermissionsInfoPolicyNameModeEnum.to_proto(
                resource.policy_name_mode
            )
        if InstanceCreateRecipeStepsPermissionsInfoResource.to_proto(resource.resource):
            res.resource.CopyFrom(
                InstanceCreateRecipeStepsPermissionsInfoResource.to_proto(
                    resource.resource
                )
            )
        else:
            res.ClearField("resource")
        return res

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return None

        return InstanceCreateRecipeStepsPermissionsInfo(
            policy_name=InstanceCreateRecipeStepsPermissionsInfoPolicyName.from_proto(
                resource.policy_name
            ),
            iam_permissions=InstanceCreateRecipeStepsPermissionsInfoIamPermissionsArray.from_proto(
                resource.iam_permissions
            ),
            resource_path=Primitive.from_proto(resource.resource_path),
            api_attrs=InstanceGoogleprotobufstruct.from_proto(resource.api_attrs),
            policy_name_mode=InstanceCreateRecipeStepsPermissionsInfoPolicyNameModeEnum.from_proto(
                resource.policy_name_mode
            ),
            resource=InstanceCreateRecipeStepsPermissionsInfoResource.from_proto(
                resource.resource
            ),
        )


class InstanceCreateRecipeStepsPermissionsInfoArray(object):
    @classmethod
    def to_proto(self, resources):
        if not resources:
            return resources
        return [InstanceCreateRecipeStepsPermissionsInfo.to_proto(i) for i in resources]

    @classmethod
    def from_proto(self, resources):
        return [
            InstanceCreateRecipeStepsPermissionsInfo.from_proto(i) for i in resources
        ]


class InstanceCreateRecipeStepsPermissionsInfoPolicyName(object):
    def __init__(self, type: str = None, id: str = None, region: str = None):
        self.type = type
        self.id = id
        self.region = region

    @classmethod
    def to_proto(self, resource):
        if not resource:
            return None

        res = (
            instance_pb2.Tier2AlphaInstanceCreateRecipeStepsPermissionsInfoPolicyName()
        )
        if Primitive.to_proto(resource.type):
            res.type = Primitive.to_proto(resource.type)
        if Primitive.to_proto(resource.id):
            res.id = Primitive.to_proto(resource.id)
        if Primitive.to_proto(resource.region):
            res.region = Primitive.to_proto(resource.region)
        return res

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return None

        return InstanceCreateRecipeStepsPermissionsInfoPolicyName(
            type=Primitive.from_proto(resource.type),
            id=Primitive.from_proto(resource.id),
            region=Primitive.from_proto(resource.region),
        )


class InstanceCreateRecipeStepsPermissionsInfoPolicyNameArray(object):
    @classmethod
    def to_proto(self, resources):
        if not resources:
            return resources
        return [
            InstanceCreateRecipeStepsPermissionsInfoPolicyName.to_proto(i)
            for i in resources
        ]

    @classmethod
    def from_proto(self, resources):
        return [
            InstanceCreateRecipeStepsPermissionsInfoPolicyName.from_proto(i)
            for i in resources
        ]


class InstanceCreateRecipeStepsPermissionsInfoIamPermissions(object):
    def __init__(self, permission: str = None):
        self.permission = permission

    @classmethod
    def to_proto(self, resource):
        if not resource:
            return None

        res = (
            instance_pb2.Tier2AlphaInstanceCreateRecipeStepsPermissionsInfoIamPermissions()
        )
        if Primitive.to_proto(resource.permission):
            res.permission = Primitive.to_proto(resource.permission)
        return res

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return None

        return InstanceCreateRecipeStepsPermissionsInfoIamPermissions(
            permission=Primitive.from_proto(resource.permission),
        )


class InstanceCreateRecipeStepsPermissionsInfoIamPermissionsArray(object):
    @classmethod
    def to_proto(self, resources):
        if not resources:
            return resources
        return [
            InstanceCreateRecipeStepsPermissionsInfoIamPermissions.to_proto(i)
            for i in resources
        ]

    @classmethod
    def from_proto(self, resources):
        return [
            InstanceCreateRecipeStepsPermissionsInfoIamPermissions.from_proto(i)
            for i in resources
        ]


class InstanceCreateRecipeStepsPermissionsInfoResource(object):
    def __init__(
        self,
        name: str = None,
        type: str = None,
        service: str = None,
        labels: dict = None,
    ):
        self.name = name
        self.type = type
        self.service = service
        self.labels = labels

    @classmethod
    def to_proto(self, resource):
        if not resource:
            return None

        res = instance_pb2.Tier2AlphaInstanceCreateRecipeStepsPermissionsInfoResource()
        if Primitive.to_proto(resource.name):
            res.name = Primitive.to_proto(resource.name)
        if Primitive.to_proto(resource.type):
            res.type = Primitive.to_proto(resource.type)
        if Primitive.to_proto(resource.service):
            res.service = Primitive.to_proto(resource.service)
        if Primitive.to_proto(resource.labels):
            res.labels = Primitive.to_proto(resource.labels)
        return res

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return None

        return InstanceCreateRecipeStepsPermissionsInfoResource(
            name=Primitive.from_proto(resource.name),
            type=Primitive.from_proto(resource.type),
            service=Primitive.from_proto(resource.service),
            labels=Primitive.from_proto(resource.labels),
        )


class InstanceCreateRecipeStepsPermissionsInfoResourceArray(object):
    @classmethod
    def to_proto(self, resources):
        if not resources:
            return resources
        return [
            InstanceCreateRecipeStepsPermissionsInfoResource.to_proto(i)
            for i in resources
        ]

    @classmethod
    def from_proto(self, resources):
        return [
            InstanceCreateRecipeStepsPermissionsInfoResource.from_proto(i)
            for i in resources
        ]


class InstanceCreateRecipeStepsKeyNotificationsUpdate(object):
    def __init__(self, key_notifications_info: dict = None):
        self.key_notifications_info = key_notifications_info

    @classmethod
    def to_proto(self, resource):
        if not resource:
            return None

        res = instance_pb2.Tier2AlphaInstanceCreateRecipeStepsKeyNotificationsUpdate()
        if InstanceCreateRecipeStepsKeyNotificationsUpdateKeyNotificationsInfo.to_proto(
            resource.key_notifications_info
        ):
            res.key_notifications_info.CopyFrom(
                InstanceCreateRecipeStepsKeyNotificationsUpdateKeyNotificationsInfo.to_proto(
                    resource.key_notifications_info
                )
            )
        else:
            res.ClearField("key_notifications_info")
        return res

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return None

        return InstanceCreateRecipeStepsKeyNotificationsUpdate(
            key_notifications_info=InstanceCreateRecipeStepsKeyNotificationsUpdateKeyNotificationsInfo.from_proto(
                resource.key_notifications_info
            ),
        )


class InstanceCreateRecipeStepsKeyNotificationsUpdateArray(object):
    @classmethod
    def to_proto(self, resources):
        if not resources:
            return resources
        return [
            InstanceCreateRecipeStepsKeyNotificationsUpdate.to_proto(i)
            for i in resources
        ]

    @classmethod
    def from_proto(self, resources):
        return [
            InstanceCreateRecipeStepsKeyNotificationsUpdate.from_proto(i)
            for i in resources
        ]


class InstanceCreateRecipeStepsKeyNotificationsUpdateKeyNotificationsInfo(object):
    def __init__(
        self,
        data_version: int = None,
        delegate: str = None,
        key_notification_configs: list = None,
    ):
        self.data_version = data_version
        self.delegate = delegate
        self.key_notification_configs = key_notification_configs

    @classmethod
    def to_proto(self, resource):
        if not resource:
            return None

        res = (
            instance_pb2.Tier2AlphaInstanceCreateRecipeStepsKeyNotificationsUpdateKeyNotificationsInfo()
        )
        if Primitive.to_proto(resource.data_version):
            res.data_version = Primitive.to_proto(resource.data_version)
        if Primitive.to_proto(resource.delegate):
            res.delegate = Primitive.to_proto(resource.delegate)
        if InstanceCreateRecipeStepsKeyNotificationsUpdateKeyNotificationsInfoKeyNotificationConfigsArray.to_proto(
            resource.key_notification_configs
        ):
            res.key_notification_configs.extend(
                InstanceCreateRecipeStepsKeyNotificationsUpdateKeyNotificationsInfoKeyNotificationConfigsArray.to_proto(
                    resource.key_notification_configs
                )
            )
        return res

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return None

        return InstanceCreateRecipeStepsKeyNotificationsUpdateKeyNotificationsInfo(
            data_version=Primitive.from_proto(resource.data_version),
            delegate=Primitive.from_proto(resource.delegate),
            key_notification_configs=InstanceCreateRecipeStepsKeyNotificationsUpdateKeyNotificationsInfoKeyNotificationConfigsArray.from_proto(
                resource.key_notification_configs
            ),
        )


class InstanceCreateRecipeStepsKeyNotificationsUpdateKeyNotificationsInfoArray(object):
    @classmethod
    def to_proto(self, resources):
        if not resources:
            return resources
        return [
            InstanceCreateRecipeStepsKeyNotificationsUpdateKeyNotificationsInfo.to_proto(
                i
            )
            for i in resources
        ]

    @classmethod
    def from_proto(self, resources):
        return [
            InstanceCreateRecipeStepsKeyNotificationsUpdateKeyNotificationsInfo.from_proto(
                i
            )
            for i in resources
        ]


class InstanceCreateRecipeStepsKeyNotificationsUpdateKeyNotificationsInfoKeyNotificationConfigs(
    object
):
    def __init__(
        self,
        key_or_version_name: str = None,
        grant: str = None,
        delegator_gaia_id: int = None,
    ):
        self.key_or_version_name = key_or_version_name
        self.grant = grant
        self.delegator_gaia_id = delegator_gaia_id

    @classmethod
    def to_proto(self, resource):
        if not resource:
            return None

        res = (
            instance_pb2.Tier2AlphaInstanceCreateRecipeStepsKeyNotificationsUpdateKeyNotificationsInfoKeyNotificationConfigs()
        )
        if Primitive.to_proto(resource.key_or_version_name):
            res.key_or_version_name = Primitive.to_proto(resource.key_or_version_name)
        if Primitive.to_proto(resource.grant):
            res.grant = Primitive.to_proto(resource.grant)
        if Primitive.to_proto(resource.delegator_gaia_id):
            res.delegator_gaia_id = Primitive.to_proto(resource.delegator_gaia_id)
        return res

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return None

        return InstanceCreateRecipeStepsKeyNotificationsUpdateKeyNotificationsInfoKeyNotificationConfigs(
            key_or_version_name=Primitive.from_proto(resource.key_or_version_name),
            grant=Primitive.from_proto(resource.grant),
            delegator_gaia_id=Primitive.from_proto(resource.delegator_gaia_id),
        )


class InstanceCreateRecipeStepsKeyNotificationsUpdateKeyNotificationsInfoKeyNotificationConfigsArray(
    object
):
    @classmethod
    def to_proto(self, resources):
        if not resources:
            return resources
        return [
            InstanceCreateRecipeStepsKeyNotificationsUpdateKeyNotificationsInfoKeyNotificationConfigs.to_proto(
                i
            )
            for i in resources
        ]

    @classmethod
    def from_proto(self, resources):
        return [
            InstanceCreateRecipeStepsKeyNotificationsUpdateKeyNotificationsInfoKeyNotificationConfigs.from_proto(
                i
            )
            for i in resources
        ]


class InstanceDeleteRecipe(object):
    def __init__(
        self,
        steps: list = None,
        honor_cancel_request: bool = None,
        ignore_recipe_after: int = None,
        verify_deadline_seconds_below: float = None,
        populate_operation_result: bool = None,
        readonly_recipe_start_time: str = None,
        resource_names_stored_in_clh_with_delay: list = None,
        delay_to_store_resources_in_clh_db_nanos: int = None,
    ):
        self.steps = steps
        self.honor_cancel_request = honor_cancel_request
        self.ignore_recipe_after = ignore_recipe_after
        self.verify_deadline_seconds_below = verify_deadline_seconds_below
        self.populate_operation_result = populate_operation_result
        self.readonly_recipe_start_time = readonly_recipe_start_time
        self.resource_names_stored_in_clh_with_delay = (
            resource_names_stored_in_clh_with_delay
        )
        self.delay_to_store_resources_in_clh_db_nanos = (
            delay_to_store_resources_in_clh_db_nanos
        )

    @classmethod
    def to_proto(self, resource):
        if not resource:
            return None

        res = instance_pb2.Tier2AlphaInstanceDeleteRecipe()
        if InstanceDeleteRecipeStepsArray.to_proto(resource.steps):
            res.steps.extend(InstanceDeleteRecipeStepsArray.to_proto(resource.steps))
        if Primitive.to_proto(resource.honor_cancel_request):
            res.honor_cancel_request = Primitive.to_proto(resource.honor_cancel_request)
        if Primitive.to_proto(resource.ignore_recipe_after):
            res.ignore_recipe_after = Primitive.to_proto(resource.ignore_recipe_after)
        if Primitive.to_proto(resource.verify_deadline_seconds_below):
            res.verify_deadline_seconds_below = Primitive.to_proto(
                resource.verify_deadline_seconds_below
            )
        if Primitive.to_proto(resource.populate_operation_result):
            res.populate_operation_result = Primitive.to_proto(
                resource.populate_operation_result
            )
        if Primitive.to_proto(resource.readonly_recipe_start_time):
            res.readonly_recipe_start_time = Primitive.to_proto(
                resource.readonly_recipe_start_time
            )
        if Primitive.to_proto(resource.resource_names_stored_in_clh_with_delay):
            res.resource_names_stored_in_clh_with_delay.extend(
                Primitive.to_proto(resource.resource_names_stored_in_clh_with_delay)
            )
        if Primitive.to_proto(resource.delay_to_store_resources_in_clh_db_nanos):
            res.delay_to_store_resources_in_clh_db_nanos = Primitive.to_proto(
                resource.delay_to_store_resources_in_clh_db_nanos
            )
        return res

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return None

        return InstanceDeleteRecipe(
            steps=InstanceDeleteRecipeStepsArray.from_proto(resource.steps),
            honor_cancel_request=Primitive.from_proto(resource.honor_cancel_request),
            ignore_recipe_after=Primitive.from_proto(resource.ignore_recipe_after),
            verify_deadline_seconds_below=Primitive.from_proto(
                resource.verify_deadline_seconds_below
            ),
            populate_operation_result=Primitive.from_proto(
                resource.populate_operation_result
            ),
            readonly_recipe_start_time=Primitive.from_proto(
                resource.readonly_recipe_start_time
            ),
            resource_names_stored_in_clh_with_delay=Primitive.from_proto(
                resource.resource_names_stored_in_clh_with_delay
            ),
            delay_to_store_resources_in_clh_db_nanos=Primitive.from_proto(
                resource.delay_to_store_resources_in_clh_db_nanos
            ),
        )


class InstanceDeleteRecipeArray(object):
    @classmethod
    def to_proto(self, resources):
        if not resources:
            return resources
        return [InstanceDeleteRecipe.to_proto(i) for i in resources]

    @classmethod
    def from_proto(self, resources):
        return [InstanceDeleteRecipe.from_proto(i) for i in resources]


class InstanceDeleteRecipeSteps(object):
    def __init__(
        self,
        relative_time: int = None,
        sleep_duration: int = None,
        action: str = None,
        status: dict = None,
        error_space: str = None,
        p4_service_account: str = None,
        resource_metadata_size: int = None,
        description: str = None,
        updated_repeat_operation_delay_sec: float = None,
        quota_request_deltas: list = None,
        preprocess_update: dict = None,
        public_operation_metadata: str = None,
        requested_tenant_project: dict = None,
        permissions_info: list = None,
        key_notifications_update: dict = None,
        clh_data_update_time: str = None,
        public_error_message: str = None,
    ):
        self.relative_time = relative_time
        self.sleep_duration = sleep_duration
        self.action = action
        self.status = status
        self.error_space = error_space
        self.p4_service_account = p4_service_account
        self.resource_metadata_size = resource_metadata_size
        self.description = description
        self.updated_repeat_operation_delay_sec = updated_repeat_operation_delay_sec
        self.quota_request_deltas = quota_request_deltas
        self.preprocess_update = preprocess_update
        self.public_operation_metadata = public_operation_metadata
        self.requested_tenant_project = requested_tenant_project
        self.permissions_info = permissions_info
        self.key_notifications_update = key_notifications_update
        self.clh_data_update_time = clh_data_update_time
        self.public_error_message = public_error_message

    @classmethod
    def to_proto(self, resource):
        if not resource:
            return None

        res = instance_pb2.Tier2AlphaInstanceDeleteRecipeSteps()
        if Primitive.to_proto(resource.relative_time):
            res.relative_time = Primitive.to_proto(resource.relative_time)
        if Primitive.to_proto(resource.sleep_duration):
            res.sleep_duration = Primitive.to_proto(resource.sleep_duration)
        if InstanceDeleteRecipeStepsActionEnum.to_proto(resource.action):
            res.action = InstanceDeleteRecipeStepsActionEnum.to_proto(resource.action)
        if InstanceDeleteRecipeStepsStatus.to_proto(resource.status):
            res.status.CopyFrom(
                InstanceDeleteRecipeStepsStatus.to_proto(resource.status)
            )
        else:
            res.ClearField("status")
        if Primitive.to_proto(resource.error_space):
            res.error_space = Primitive.to_proto(resource.error_space)
        if Primitive.to_proto(resource.p4_service_account):
            res.p4_service_account = Primitive.to_proto(resource.p4_service_account)
        if Primitive.to_proto(resource.resource_metadata_size):
            res.resource_metadata_size = Primitive.to_proto(
                resource.resource_metadata_size
            )
        if Primitive.to_proto(resource.description):
            res.description = Primitive.to_proto(resource.description)
        if Primitive.to_proto(resource.updated_repeat_operation_delay_sec):
            res.updated_repeat_operation_delay_sec = Primitive.to_proto(
                resource.updated_repeat_operation_delay_sec
            )
        if InstanceDeleteRecipeStepsQuotaRequestDeltasArray.to_proto(
            resource.quota_request_deltas
        ):
            res.quota_request_deltas.extend(
                InstanceDeleteRecipeStepsQuotaRequestDeltasArray.to_proto(
                    resource.quota_request_deltas
                )
            )
        if InstanceDeleteRecipeStepsPreprocessUpdate.to_proto(
            resource.preprocess_update
        ):
            res.preprocess_update.CopyFrom(
                InstanceDeleteRecipeStepsPreprocessUpdate.to_proto(
                    resource.preprocess_update
                )
            )
        else:
            res.ClearField("preprocess_update")
        if Primitive.to_proto(resource.public_operation_metadata):
            res.public_operation_metadata = Primitive.to_proto(
                resource.public_operation_metadata
            )
        if InstanceDeleteRecipeStepsRequestedTenantProject.to_proto(
            resource.requested_tenant_project
        ):
            res.requested_tenant_project.CopyFrom(
                InstanceDeleteRecipeStepsRequestedTenantProject.to_proto(
                    resource.requested_tenant_project
                )
            )
        else:
            res.ClearField("requested_tenant_project")
        if InstanceDeleteRecipeStepsPermissionsInfoArray.to_proto(
            resource.permissions_info
        ):
            res.permissions_info.extend(
                InstanceDeleteRecipeStepsPermissionsInfoArray.to_proto(
                    resource.permissions_info
                )
            )
        if InstanceDeleteRecipeStepsKeyNotificationsUpdate.to_proto(
            resource.key_notifications_update
        ):
            res.key_notifications_update.CopyFrom(
                InstanceDeleteRecipeStepsKeyNotificationsUpdate.to_proto(
                    resource.key_notifications_update
                )
            )
        else:
            res.ClearField("key_notifications_update")
        if Primitive.to_proto(resource.clh_data_update_time):
            res.clh_data_update_time = Primitive.to_proto(resource.clh_data_update_time)
        if Primitive.to_proto(resource.public_error_message):
            res.public_error_message = Primitive.to_proto(resource.public_error_message)
        return res

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return None

        return InstanceDeleteRecipeSteps(
            relative_time=Primitive.from_proto(resource.relative_time),
            sleep_duration=Primitive.from_proto(resource.sleep_duration),
            action=InstanceDeleteRecipeStepsActionEnum.from_proto(resource.action),
            status=InstanceDeleteRecipeStepsStatus.from_proto(resource.status),
            error_space=Primitive.from_proto(resource.error_space),
            p4_service_account=Primitive.from_proto(resource.p4_service_account),
            resource_metadata_size=Primitive.from_proto(
                resource.resource_metadata_size
            ),
            description=Primitive.from_proto(resource.description),
            updated_repeat_operation_delay_sec=Primitive.from_proto(
                resource.updated_repeat_operation_delay_sec
            ),
            quota_request_deltas=InstanceDeleteRecipeStepsQuotaRequestDeltasArray.from_proto(
                resource.quota_request_deltas
            ),
            preprocess_update=InstanceDeleteRecipeStepsPreprocessUpdate.from_proto(
                resource.preprocess_update
            ),
            public_operation_metadata=Primitive.from_proto(
                resource.public_operation_metadata
            ),
            requested_tenant_project=InstanceDeleteRecipeStepsRequestedTenantProject.from_proto(
                resource.requested_tenant_project
            ),
            permissions_info=InstanceDeleteRecipeStepsPermissionsInfoArray.from_proto(
                resource.permissions_info
            ),
            key_notifications_update=InstanceDeleteRecipeStepsKeyNotificationsUpdate.from_proto(
                resource.key_notifications_update
            ),
            clh_data_update_time=Primitive.from_proto(resource.clh_data_update_time),
            public_error_message=Primitive.from_proto(resource.public_error_message),
        )


class InstanceDeleteRecipeStepsArray(object):
    @classmethod
    def to_proto(self, resources):
        if not resources:
            return resources
        return [InstanceDeleteRecipeSteps.to_proto(i) for i in resources]

    @classmethod
    def from_proto(self, resources):
        return [InstanceDeleteRecipeSteps.from_proto(i) for i in resources]


class InstanceDeleteRecipeStepsStatus(object):
    def __init__(self, code: int = None, message: str = None, details: list = None):
        self.code = code
        self.message = message
        self.details = details

    @classmethod
    def to_proto(self, resource):
        if not resource:
            return None

        res = instance_pb2.Tier2AlphaInstanceDeleteRecipeStepsStatus()
        if Primitive.to_proto(resource.code):
            res.code = Primitive.to_proto(resource.code)
        if Primitive.to_proto(resource.message):
            res.message = Primitive.to_proto(resource.message)
        if InstanceDeleteRecipeStepsStatusDetailsArray.to_proto(resource.details):
            res.details.extend(
                InstanceDeleteRecipeStepsStatusDetailsArray.to_proto(resource.details)
            )
        return res

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return None

        return InstanceDeleteRecipeStepsStatus(
            code=Primitive.from_proto(resource.code),
            message=Primitive.from_proto(resource.message),
            details=InstanceDeleteRecipeStepsStatusDetailsArray.from_proto(
                resource.details
            ),
        )


class InstanceDeleteRecipeStepsStatusArray(object):
    @classmethod
    def to_proto(self, resources):
        if not resources:
            return resources
        return [InstanceDeleteRecipeStepsStatus.to_proto(i) for i in resources]

    @classmethod
    def from_proto(self, resources):
        return [InstanceDeleteRecipeStepsStatus.from_proto(i) for i in resources]


class InstanceDeleteRecipeStepsStatusDetails(object):
    def __init__(self, type_url: str = None, value: str = None):
        self.type_url = type_url
        self.value = value

    @classmethod
    def to_proto(self, resource):
        if not resource:
            return None

        res = instance_pb2.Tier2AlphaInstanceDeleteRecipeStepsStatusDetails()
        if Primitive.to_proto(resource.type_url):
            res.type_url = Primitive.to_proto(resource.type_url)
        if Primitive.to_proto(resource.value):
            res.value = Primitive.to_proto(resource.value)
        return res

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return None

        return InstanceDeleteRecipeStepsStatusDetails(
            type_url=Primitive.from_proto(resource.type_url),
            value=Primitive.from_proto(resource.value),
        )


class InstanceDeleteRecipeStepsStatusDetailsArray(object):
    @classmethod
    def to_proto(self, resources):
        if not resources:
            return resources
        return [InstanceDeleteRecipeStepsStatusDetails.to_proto(i) for i in resources]

    @classmethod
    def from_proto(self, resources):
        return [InstanceDeleteRecipeStepsStatusDetails.from_proto(i) for i in resources]


class InstanceDeleteRecipeStepsQuotaRequestDeltas(object):
    def __init__(
        self,
        metric_name: str = None,
        amount: int = None,
        quota_location_name: str = None,
    ):
        self.metric_name = metric_name
        self.amount = amount
        self.quota_location_name = quota_location_name

    @classmethod
    def to_proto(self, resource):
        if not resource:
            return None

        res = instance_pb2.Tier2AlphaInstanceDeleteRecipeStepsQuotaRequestDeltas()
        if Primitive.to_proto(resource.metric_name):
            res.metric_name = Primitive.to_proto(resource.metric_name)
        if Primitive.to_proto(resource.amount):
            res.amount = Primitive.to_proto(resource.amount)
        if Primitive.to_proto(resource.quota_location_name):
            res.quota_location_name = Primitive.to_proto(resource.quota_location_name)
        return res

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return None

        return InstanceDeleteRecipeStepsQuotaRequestDeltas(
            metric_name=Primitive.from_proto(resource.metric_name),
            amount=Primitive.from_proto(resource.amount),
            quota_location_name=Primitive.from_proto(resource.quota_location_name),
        )


class InstanceDeleteRecipeStepsQuotaRequestDeltasArray(object):
    @classmethod
    def to_proto(self, resources):
        if not resources:
            return resources
        return [
            InstanceDeleteRecipeStepsQuotaRequestDeltas.to_proto(i) for i in resources
        ]

    @classmethod
    def from_proto(self, resources):
        return [
            InstanceDeleteRecipeStepsQuotaRequestDeltas.from_proto(i) for i in resources
        ]


class InstanceDeleteRecipeStepsPreprocessUpdate(object):
    def __init__(
        self, latency_slo_bucket_name: str = None, public_operation_metadata: str = None
    ):
        self.latency_slo_bucket_name = latency_slo_bucket_name
        self.public_operation_metadata = public_operation_metadata

    @classmethod
    def to_proto(self, resource):
        if not resource:
            return None

        res = instance_pb2.Tier2AlphaInstanceDeleteRecipeStepsPreprocessUpdate()
        if Primitive.to_proto(resource.latency_slo_bucket_name):
            res.latency_slo_bucket_name = Primitive.to_proto(
                resource.latency_slo_bucket_name
            )
        if Primitive.to_proto(resource.public_operation_metadata):
            res.public_operation_metadata = Primitive.to_proto(
                resource.public_operation_metadata
            )
        return res

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return None

        return InstanceDeleteRecipeStepsPreprocessUpdate(
            latency_slo_bucket_name=Primitive.from_proto(
                resource.latency_slo_bucket_name
            ),
            public_operation_metadata=Primitive.from_proto(
                resource.public_operation_metadata
            ),
        )


class InstanceDeleteRecipeStepsPreprocessUpdateArray(object):
    @classmethod
    def to_proto(self, resources):
        if not resources:
            return resources
        return [
            InstanceDeleteRecipeStepsPreprocessUpdate.to_proto(i) for i in resources
        ]

    @classmethod
    def from_proto(self, resources):
        return [
            InstanceDeleteRecipeStepsPreprocessUpdate.from_proto(i) for i in resources
        ]


class InstanceDeleteRecipeStepsRequestedTenantProject(object):
    def __init__(self, tag: str = None, folder: str = None, scope: str = None):
        self.tag = tag
        self.folder = folder
        self.scope = scope

    @classmethod
    def to_proto(self, resource):
        if not resource:
            return None

        res = instance_pb2.Tier2AlphaInstanceDeleteRecipeStepsRequestedTenantProject()
        if Primitive.to_proto(resource.tag):
            res.tag = Primitive.to_proto(resource.tag)
        if Primitive.to_proto(resource.folder):
            res.folder = Primitive.to_proto(resource.folder)
        if InstanceDeleteRecipeStepsRequestedTenantProjectScopeEnum.to_proto(
            resource.scope
        ):
            res.scope = InstanceDeleteRecipeStepsRequestedTenantProjectScopeEnum.to_proto(
                resource.scope
            )
        return res

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return None

        return InstanceDeleteRecipeStepsRequestedTenantProject(
            tag=Primitive.from_proto(resource.tag),
            folder=Primitive.from_proto(resource.folder),
            scope=InstanceDeleteRecipeStepsRequestedTenantProjectScopeEnum.from_proto(
                resource.scope
            ),
        )


class InstanceDeleteRecipeStepsRequestedTenantProjectArray(object):
    @classmethod
    def to_proto(self, resources):
        if not resources:
            return resources
        return [
            InstanceDeleteRecipeStepsRequestedTenantProject.to_proto(i)
            for i in resources
        ]

    @classmethod
    def from_proto(self, resources):
        return [
            InstanceDeleteRecipeStepsRequestedTenantProject.from_proto(i)
            for i in resources
        ]


class InstanceDeleteRecipeStepsPermissionsInfo(object):
    def __init__(
        self,
        policy_name: dict = None,
        iam_permissions: list = None,
        resource_path: str = None,
        api_attrs: dict = None,
        policy_name_mode: str = None,
        resource: dict = None,
    ):
        self.policy_name = policy_name
        self.iam_permissions = iam_permissions
        self.resource_path = resource_path
        self.api_attrs = api_attrs
        self.policy_name_mode = policy_name_mode
        self.resource = resource

    @classmethod
    def to_proto(self, resource):
        if not resource:
            return None

        res = instance_pb2.Tier2AlphaInstanceDeleteRecipeStepsPermissionsInfo()
        if InstanceDeleteRecipeStepsPermissionsInfoPolicyName.to_proto(
            resource.policy_name
        ):
            res.policy_name.CopyFrom(
                InstanceDeleteRecipeStepsPermissionsInfoPolicyName.to_proto(
                    resource.policy_name
                )
            )
        else:
            res.ClearField("policy_name")
        if InstanceDeleteRecipeStepsPermissionsInfoIamPermissionsArray.to_proto(
            resource.iam_permissions
        ):
            res.iam_permissions.extend(
                InstanceDeleteRecipeStepsPermissionsInfoIamPermissionsArray.to_proto(
                    resource.iam_permissions
                )
            )
        if Primitive.to_proto(resource.resource_path):
            res.resource_path = Primitive.to_proto(resource.resource_path)
        if InstanceGoogleprotobufstruct.to_proto(resource.api_attrs):
            res.api_attrs.CopyFrom(
                InstanceGoogleprotobufstruct.to_proto(resource.api_attrs)
            )
        else:
            res.ClearField("api_attrs")
        if InstanceDeleteRecipeStepsPermissionsInfoPolicyNameModeEnum.to_proto(
            resource.policy_name_mode
        ):
            res.policy_name_mode = InstanceDeleteRecipeStepsPermissionsInfoPolicyNameModeEnum.to_proto(
                resource.policy_name_mode
            )
        if InstanceDeleteRecipeStepsPermissionsInfoResource.to_proto(resource.resource):
            res.resource.CopyFrom(
                InstanceDeleteRecipeStepsPermissionsInfoResource.to_proto(
                    resource.resource
                )
            )
        else:
            res.ClearField("resource")
        return res

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return None

        return InstanceDeleteRecipeStepsPermissionsInfo(
            policy_name=InstanceDeleteRecipeStepsPermissionsInfoPolicyName.from_proto(
                resource.policy_name
            ),
            iam_permissions=InstanceDeleteRecipeStepsPermissionsInfoIamPermissionsArray.from_proto(
                resource.iam_permissions
            ),
            resource_path=Primitive.from_proto(resource.resource_path),
            api_attrs=InstanceGoogleprotobufstruct.from_proto(resource.api_attrs),
            policy_name_mode=InstanceDeleteRecipeStepsPermissionsInfoPolicyNameModeEnum.from_proto(
                resource.policy_name_mode
            ),
            resource=InstanceDeleteRecipeStepsPermissionsInfoResource.from_proto(
                resource.resource
            ),
        )


class InstanceDeleteRecipeStepsPermissionsInfoArray(object):
    @classmethod
    def to_proto(self, resources):
        if not resources:
            return resources
        return [InstanceDeleteRecipeStepsPermissionsInfo.to_proto(i) for i in resources]

    @classmethod
    def from_proto(self, resources):
        return [
            InstanceDeleteRecipeStepsPermissionsInfo.from_proto(i) for i in resources
        ]


class InstanceDeleteRecipeStepsPermissionsInfoPolicyName(object):
    def __init__(self, type: str = None, id: str = None, region: str = None):
        self.type = type
        self.id = id
        self.region = region

    @classmethod
    def to_proto(self, resource):
        if not resource:
            return None

        res = (
            instance_pb2.Tier2AlphaInstanceDeleteRecipeStepsPermissionsInfoPolicyName()
        )
        if Primitive.to_proto(resource.type):
            res.type = Primitive.to_proto(resource.type)
        if Primitive.to_proto(resource.id):
            res.id = Primitive.to_proto(resource.id)
        if Primitive.to_proto(resource.region):
            res.region = Primitive.to_proto(resource.region)
        return res

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return None

        return InstanceDeleteRecipeStepsPermissionsInfoPolicyName(
            type=Primitive.from_proto(resource.type),
            id=Primitive.from_proto(resource.id),
            region=Primitive.from_proto(resource.region),
        )


class InstanceDeleteRecipeStepsPermissionsInfoPolicyNameArray(object):
    @classmethod
    def to_proto(self, resources):
        if not resources:
            return resources
        return [
            InstanceDeleteRecipeStepsPermissionsInfoPolicyName.to_proto(i)
            for i in resources
        ]

    @classmethod
    def from_proto(self, resources):
        return [
            InstanceDeleteRecipeStepsPermissionsInfoPolicyName.from_proto(i)
            for i in resources
        ]


class InstanceDeleteRecipeStepsPermissionsInfoIamPermissions(object):
    def __init__(self, permission: str = None):
        self.permission = permission

    @classmethod
    def to_proto(self, resource):
        if not resource:
            return None

        res = (
            instance_pb2.Tier2AlphaInstanceDeleteRecipeStepsPermissionsInfoIamPermissions()
        )
        if Primitive.to_proto(resource.permission):
            res.permission = Primitive.to_proto(resource.permission)
        return res

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return None

        return InstanceDeleteRecipeStepsPermissionsInfoIamPermissions(
            permission=Primitive.from_proto(resource.permission),
        )


class InstanceDeleteRecipeStepsPermissionsInfoIamPermissionsArray(object):
    @classmethod
    def to_proto(self, resources):
        if not resources:
            return resources
        return [
            InstanceDeleteRecipeStepsPermissionsInfoIamPermissions.to_proto(i)
            for i in resources
        ]

    @classmethod
    def from_proto(self, resources):
        return [
            InstanceDeleteRecipeStepsPermissionsInfoIamPermissions.from_proto(i)
            for i in resources
        ]


class InstanceDeleteRecipeStepsPermissionsInfoResource(object):
    def __init__(
        self,
        name: str = None,
        type: str = None,
        service: str = None,
        labels: dict = None,
    ):
        self.name = name
        self.type = type
        self.service = service
        self.labels = labels

    @classmethod
    def to_proto(self, resource):
        if not resource:
            return None

        res = instance_pb2.Tier2AlphaInstanceDeleteRecipeStepsPermissionsInfoResource()
        if Primitive.to_proto(resource.name):
            res.name = Primitive.to_proto(resource.name)
        if Primitive.to_proto(resource.type):
            res.type = Primitive.to_proto(resource.type)
        if Primitive.to_proto(resource.service):
            res.service = Primitive.to_proto(resource.service)
        if Primitive.to_proto(resource.labels):
            res.labels = Primitive.to_proto(resource.labels)
        return res

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return None

        return InstanceDeleteRecipeStepsPermissionsInfoResource(
            name=Primitive.from_proto(resource.name),
            type=Primitive.from_proto(resource.type),
            service=Primitive.from_proto(resource.service),
            labels=Primitive.from_proto(resource.labels),
        )


class InstanceDeleteRecipeStepsPermissionsInfoResourceArray(object):
    @classmethod
    def to_proto(self, resources):
        if not resources:
            return resources
        return [
            InstanceDeleteRecipeStepsPermissionsInfoResource.to_proto(i)
            for i in resources
        ]

    @classmethod
    def from_proto(self, resources):
        return [
            InstanceDeleteRecipeStepsPermissionsInfoResource.from_proto(i)
            for i in resources
        ]


class InstanceDeleteRecipeStepsKeyNotificationsUpdate(object):
    def __init__(self, key_notifications_info: dict = None):
        self.key_notifications_info = key_notifications_info

    @classmethod
    def to_proto(self, resource):
        if not resource:
            return None

        res = instance_pb2.Tier2AlphaInstanceDeleteRecipeStepsKeyNotificationsUpdate()
        if InstanceDeleteRecipeStepsKeyNotificationsUpdateKeyNotificationsInfo.to_proto(
            resource.key_notifications_info
        ):
            res.key_notifications_info.CopyFrom(
                InstanceDeleteRecipeStepsKeyNotificationsUpdateKeyNotificationsInfo.to_proto(
                    resource.key_notifications_info
                )
            )
        else:
            res.ClearField("key_notifications_info")
        return res

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return None

        return InstanceDeleteRecipeStepsKeyNotificationsUpdate(
            key_notifications_info=InstanceDeleteRecipeStepsKeyNotificationsUpdateKeyNotificationsInfo.from_proto(
                resource.key_notifications_info
            ),
        )


class InstanceDeleteRecipeStepsKeyNotificationsUpdateArray(object):
    @classmethod
    def to_proto(self, resources):
        if not resources:
            return resources
        return [
            InstanceDeleteRecipeStepsKeyNotificationsUpdate.to_proto(i)
            for i in resources
        ]

    @classmethod
    def from_proto(self, resources):
        return [
            InstanceDeleteRecipeStepsKeyNotificationsUpdate.from_proto(i)
            for i in resources
        ]


class InstanceDeleteRecipeStepsKeyNotificationsUpdateKeyNotificationsInfo(object):
    def __init__(
        self,
        data_version: int = None,
        delegate: str = None,
        key_notification_configs: list = None,
    ):
        self.data_version = data_version
        self.delegate = delegate
        self.key_notification_configs = key_notification_configs

    @classmethod
    def to_proto(self, resource):
        if not resource:
            return None

        res = (
            instance_pb2.Tier2AlphaInstanceDeleteRecipeStepsKeyNotificationsUpdateKeyNotificationsInfo()
        )
        if Primitive.to_proto(resource.data_version):
            res.data_version = Primitive.to_proto(resource.data_version)
        if Primitive.to_proto(resource.delegate):
            res.delegate = Primitive.to_proto(resource.delegate)
        if InstanceDeleteRecipeStepsKeyNotificationsUpdateKeyNotificationsInfoKeyNotificationConfigsArray.to_proto(
            resource.key_notification_configs
        ):
            res.key_notification_configs.extend(
                InstanceDeleteRecipeStepsKeyNotificationsUpdateKeyNotificationsInfoKeyNotificationConfigsArray.to_proto(
                    resource.key_notification_configs
                )
            )
        return res

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return None

        return InstanceDeleteRecipeStepsKeyNotificationsUpdateKeyNotificationsInfo(
            data_version=Primitive.from_proto(resource.data_version),
            delegate=Primitive.from_proto(resource.delegate),
            key_notification_configs=InstanceDeleteRecipeStepsKeyNotificationsUpdateKeyNotificationsInfoKeyNotificationConfigsArray.from_proto(
                resource.key_notification_configs
            ),
        )


class InstanceDeleteRecipeStepsKeyNotificationsUpdateKeyNotificationsInfoArray(object):
    @classmethod
    def to_proto(self, resources):
        if not resources:
            return resources
        return [
            InstanceDeleteRecipeStepsKeyNotificationsUpdateKeyNotificationsInfo.to_proto(
                i
            )
            for i in resources
        ]

    @classmethod
    def from_proto(self, resources):
        return [
            InstanceDeleteRecipeStepsKeyNotificationsUpdateKeyNotificationsInfo.from_proto(
                i
            )
            for i in resources
        ]


class InstanceDeleteRecipeStepsKeyNotificationsUpdateKeyNotificationsInfoKeyNotificationConfigs(
    object
):
    def __init__(
        self,
        key_or_version_name: str = None,
        grant: str = None,
        delegator_gaia_id: int = None,
    ):
        self.key_or_version_name = key_or_version_name
        self.grant = grant
        self.delegator_gaia_id = delegator_gaia_id

    @classmethod
    def to_proto(self, resource):
        if not resource:
            return None

        res = (
            instance_pb2.Tier2AlphaInstanceDeleteRecipeStepsKeyNotificationsUpdateKeyNotificationsInfoKeyNotificationConfigs()
        )
        if Primitive.to_proto(resource.key_or_version_name):
            res.key_or_version_name = Primitive.to_proto(resource.key_or_version_name)
        if Primitive.to_proto(resource.grant):
            res.grant = Primitive.to_proto(resource.grant)
        if Primitive.to_proto(resource.delegator_gaia_id):
            res.delegator_gaia_id = Primitive.to_proto(resource.delegator_gaia_id)
        return res

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return None

        return InstanceDeleteRecipeStepsKeyNotificationsUpdateKeyNotificationsInfoKeyNotificationConfigs(
            key_or_version_name=Primitive.from_proto(resource.key_or_version_name),
            grant=Primitive.from_proto(resource.grant),
            delegator_gaia_id=Primitive.from_proto(resource.delegator_gaia_id),
        )


class InstanceDeleteRecipeStepsKeyNotificationsUpdateKeyNotificationsInfoKeyNotificationConfigsArray(
    object
):
    @classmethod
    def to_proto(self, resources):
        if not resources:
            return resources
        return [
            InstanceDeleteRecipeStepsKeyNotificationsUpdateKeyNotificationsInfoKeyNotificationConfigs.to_proto(
                i
            )
            for i in resources
        ]

    @classmethod
    def from_proto(self, resources):
        return [
            InstanceDeleteRecipeStepsKeyNotificationsUpdateKeyNotificationsInfoKeyNotificationConfigs.from_proto(
                i
            )
            for i in resources
        ]


class InstanceUpdateRecipe(object):
    def __init__(
        self,
        steps: list = None,
        honor_cancel_request: bool = None,
        ignore_recipe_after: int = None,
        verify_deadline_seconds_below: float = None,
        populate_operation_result: bool = None,
        readonly_recipe_start_time: str = None,
        resource_names_stored_in_clh_with_delay: list = None,
        delay_to_store_resources_in_clh_db_nanos: int = None,
    ):
        self.steps = steps
        self.honor_cancel_request = honor_cancel_request
        self.ignore_recipe_after = ignore_recipe_after
        self.verify_deadline_seconds_below = verify_deadline_seconds_below
        self.populate_operation_result = populate_operation_result
        self.readonly_recipe_start_time = readonly_recipe_start_time
        self.resource_names_stored_in_clh_with_delay = (
            resource_names_stored_in_clh_with_delay
        )
        self.delay_to_store_resources_in_clh_db_nanos = (
            delay_to_store_resources_in_clh_db_nanos
        )

    @classmethod
    def to_proto(self, resource):
        if not resource:
            return None

        res = instance_pb2.Tier2AlphaInstanceUpdateRecipe()
        if InstanceUpdateRecipeStepsArray.to_proto(resource.steps):
            res.steps.extend(InstanceUpdateRecipeStepsArray.to_proto(resource.steps))
        if Primitive.to_proto(resource.honor_cancel_request):
            res.honor_cancel_request = Primitive.to_proto(resource.honor_cancel_request)
        if Primitive.to_proto(resource.ignore_recipe_after):
            res.ignore_recipe_after = Primitive.to_proto(resource.ignore_recipe_after)
        if Primitive.to_proto(resource.verify_deadline_seconds_below):
            res.verify_deadline_seconds_below = Primitive.to_proto(
                resource.verify_deadline_seconds_below
            )
        if Primitive.to_proto(resource.populate_operation_result):
            res.populate_operation_result = Primitive.to_proto(
                resource.populate_operation_result
            )
        if Primitive.to_proto(resource.readonly_recipe_start_time):
            res.readonly_recipe_start_time = Primitive.to_proto(
                resource.readonly_recipe_start_time
            )
        if Primitive.to_proto(resource.resource_names_stored_in_clh_with_delay):
            res.resource_names_stored_in_clh_with_delay.extend(
                Primitive.to_proto(resource.resource_names_stored_in_clh_with_delay)
            )
        if Primitive.to_proto(resource.delay_to_store_resources_in_clh_db_nanos):
            res.delay_to_store_resources_in_clh_db_nanos = Primitive.to_proto(
                resource.delay_to_store_resources_in_clh_db_nanos
            )
        return res

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return None

        return InstanceUpdateRecipe(
            steps=InstanceUpdateRecipeStepsArray.from_proto(resource.steps),
            honor_cancel_request=Primitive.from_proto(resource.honor_cancel_request),
            ignore_recipe_after=Primitive.from_proto(resource.ignore_recipe_after),
            verify_deadline_seconds_below=Primitive.from_proto(
                resource.verify_deadline_seconds_below
            ),
            populate_operation_result=Primitive.from_proto(
                resource.populate_operation_result
            ),
            readonly_recipe_start_time=Primitive.from_proto(
                resource.readonly_recipe_start_time
            ),
            resource_names_stored_in_clh_with_delay=Primitive.from_proto(
                resource.resource_names_stored_in_clh_with_delay
            ),
            delay_to_store_resources_in_clh_db_nanos=Primitive.from_proto(
                resource.delay_to_store_resources_in_clh_db_nanos
            ),
        )


class InstanceUpdateRecipeArray(object):
    @classmethod
    def to_proto(self, resources):
        if not resources:
            return resources
        return [InstanceUpdateRecipe.to_proto(i) for i in resources]

    @classmethod
    def from_proto(self, resources):
        return [InstanceUpdateRecipe.from_proto(i) for i in resources]


class InstanceUpdateRecipeSteps(object):
    def __init__(
        self,
        relative_time: int = None,
        sleep_duration: int = None,
        action: str = None,
        status: dict = None,
        error_space: str = None,
        p4_service_account: str = None,
        resource_metadata_size: int = None,
        description: str = None,
        updated_repeat_operation_delay_sec: float = None,
        quota_request_deltas: list = None,
        preprocess_update: dict = None,
        public_operation_metadata: str = None,
        requested_tenant_project: dict = None,
        permissions_info: list = None,
        key_notifications_update: dict = None,
        clh_data_update_time: str = None,
        public_error_message: str = None,
    ):
        self.relative_time = relative_time
        self.sleep_duration = sleep_duration
        self.action = action
        self.status = status
        self.error_space = error_space
        self.p4_service_account = p4_service_account
        self.resource_metadata_size = resource_metadata_size
        self.description = description
        self.updated_repeat_operation_delay_sec = updated_repeat_operation_delay_sec
        self.quota_request_deltas = quota_request_deltas
        self.preprocess_update = preprocess_update
        self.public_operation_metadata = public_operation_metadata
        self.requested_tenant_project = requested_tenant_project
        self.permissions_info = permissions_info
        self.key_notifications_update = key_notifications_update
        self.clh_data_update_time = clh_data_update_time
        self.public_error_message = public_error_message

    @classmethod
    def to_proto(self, resource):
        if not resource:
            return None

        res = instance_pb2.Tier2AlphaInstanceUpdateRecipeSteps()
        if Primitive.to_proto(resource.relative_time):
            res.relative_time = Primitive.to_proto(resource.relative_time)
        if Primitive.to_proto(resource.sleep_duration):
            res.sleep_duration = Primitive.to_proto(resource.sleep_duration)
        if InstanceUpdateRecipeStepsActionEnum.to_proto(resource.action):
            res.action = InstanceUpdateRecipeStepsActionEnum.to_proto(resource.action)
        if InstanceUpdateRecipeStepsStatus.to_proto(resource.status):
            res.status.CopyFrom(
                InstanceUpdateRecipeStepsStatus.to_proto(resource.status)
            )
        else:
            res.ClearField("status")
        if Primitive.to_proto(resource.error_space):
            res.error_space = Primitive.to_proto(resource.error_space)
        if Primitive.to_proto(resource.p4_service_account):
            res.p4_service_account = Primitive.to_proto(resource.p4_service_account)
        if Primitive.to_proto(resource.resource_metadata_size):
            res.resource_metadata_size = Primitive.to_proto(
                resource.resource_metadata_size
            )
        if Primitive.to_proto(resource.description):
            res.description = Primitive.to_proto(resource.description)
        if Primitive.to_proto(resource.updated_repeat_operation_delay_sec):
            res.updated_repeat_operation_delay_sec = Primitive.to_proto(
                resource.updated_repeat_operation_delay_sec
            )
        if InstanceUpdateRecipeStepsQuotaRequestDeltasArray.to_proto(
            resource.quota_request_deltas
        ):
            res.quota_request_deltas.extend(
                InstanceUpdateRecipeStepsQuotaRequestDeltasArray.to_proto(
                    resource.quota_request_deltas
                )
            )
        if InstanceUpdateRecipeStepsPreprocessUpdate.to_proto(
            resource.preprocess_update
        ):
            res.preprocess_update.CopyFrom(
                InstanceUpdateRecipeStepsPreprocessUpdate.to_proto(
                    resource.preprocess_update
                )
            )
        else:
            res.ClearField("preprocess_update")
        if Primitive.to_proto(resource.public_operation_metadata):
            res.public_operation_metadata = Primitive.to_proto(
                resource.public_operation_metadata
            )
        if InstanceUpdateRecipeStepsRequestedTenantProject.to_proto(
            resource.requested_tenant_project
        ):
            res.requested_tenant_project.CopyFrom(
                InstanceUpdateRecipeStepsRequestedTenantProject.to_proto(
                    resource.requested_tenant_project
                )
            )
        else:
            res.ClearField("requested_tenant_project")
        if InstanceUpdateRecipeStepsPermissionsInfoArray.to_proto(
            resource.permissions_info
        ):
            res.permissions_info.extend(
                InstanceUpdateRecipeStepsPermissionsInfoArray.to_proto(
                    resource.permissions_info
                )
            )
        if InstanceUpdateRecipeStepsKeyNotificationsUpdate.to_proto(
            resource.key_notifications_update
        ):
            res.key_notifications_update.CopyFrom(
                InstanceUpdateRecipeStepsKeyNotificationsUpdate.to_proto(
                    resource.key_notifications_update
                )
            )
        else:
            res.ClearField("key_notifications_update")
        if Primitive.to_proto(resource.clh_data_update_time):
            res.clh_data_update_time = Primitive.to_proto(resource.clh_data_update_time)
        if Primitive.to_proto(resource.public_error_message):
            res.public_error_message = Primitive.to_proto(resource.public_error_message)
        return res

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return None

        return InstanceUpdateRecipeSteps(
            relative_time=Primitive.from_proto(resource.relative_time),
            sleep_duration=Primitive.from_proto(resource.sleep_duration),
            action=InstanceUpdateRecipeStepsActionEnum.from_proto(resource.action),
            status=InstanceUpdateRecipeStepsStatus.from_proto(resource.status),
            error_space=Primitive.from_proto(resource.error_space),
            p4_service_account=Primitive.from_proto(resource.p4_service_account),
            resource_metadata_size=Primitive.from_proto(
                resource.resource_metadata_size
            ),
            description=Primitive.from_proto(resource.description),
            updated_repeat_operation_delay_sec=Primitive.from_proto(
                resource.updated_repeat_operation_delay_sec
            ),
            quota_request_deltas=InstanceUpdateRecipeStepsQuotaRequestDeltasArray.from_proto(
                resource.quota_request_deltas
            ),
            preprocess_update=InstanceUpdateRecipeStepsPreprocessUpdate.from_proto(
                resource.preprocess_update
            ),
            public_operation_metadata=Primitive.from_proto(
                resource.public_operation_metadata
            ),
            requested_tenant_project=InstanceUpdateRecipeStepsRequestedTenantProject.from_proto(
                resource.requested_tenant_project
            ),
            permissions_info=InstanceUpdateRecipeStepsPermissionsInfoArray.from_proto(
                resource.permissions_info
            ),
            key_notifications_update=InstanceUpdateRecipeStepsKeyNotificationsUpdate.from_proto(
                resource.key_notifications_update
            ),
            clh_data_update_time=Primitive.from_proto(resource.clh_data_update_time),
            public_error_message=Primitive.from_proto(resource.public_error_message),
        )


class InstanceUpdateRecipeStepsArray(object):
    @classmethod
    def to_proto(self, resources):
        if not resources:
            return resources
        return [InstanceUpdateRecipeSteps.to_proto(i) for i in resources]

    @classmethod
    def from_proto(self, resources):
        return [InstanceUpdateRecipeSteps.from_proto(i) for i in resources]


class InstanceUpdateRecipeStepsStatus(object):
    def __init__(self, code: int = None, message: str = None, details: list = None):
        self.code = code
        self.message = message
        self.details = details

    @classmethod
    def to_proto(self, resource):
        if not resource:
            return None

        res = instance_pb2.Tier2AlphaInstanceUpdateRecipeStepsStatus()
        if Primitive.to_proto(resource.code):
            res.code = Primitive.to_proto(resource.code)
        if Primitive.to_proto(resource.message):
            res.message = Primitive.to_proto(resource.message)
        if InstanceUpdateRecipeStepsStatusDetailsArray.to_proto(resource.details):
            res.details.extend(
                InstanceUpdateRecipeStepsStatusDetailsArray.to_proto(resource.details)
            )
        return res

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return None

        return InstanceUpdateRecipeStepsStatus(
            code=Primitive.from_proto(resource.code),
            message=Primitive.from_proto(resource.message),
            details=InstanceUpdateRecipeStepsStatusDetailsArray.from_proto(
                resource.details
            ),
        )


class InstanceUpdateRecipeStepsStatusArray(object):
    @classmethod
    def to_proto(self, resources):
        if not resources:
            return resources
        return [InstanceUpdateRecipeStepsStatus.to_proto(i) for i in resources]

    @classmethod
    def from_proto(self, resources):
        return [InstanceUpdateRecipeStepsStatus.from_proto(i) for i in resources]


class InstanceUpdateRecipeStepsStatusDetails(object):
    def __init__(self, type_url: str = None, value: str = None):
        self.type_url = type_url
        self.value = value

    @classmethod
    def to_proto(self, resource):
        if not resource:
            return None

        res = instance_pb2.Tier2AlphaInstanceUpdateRecipeStepsStatusDetails()
        if Primitive.to_proto(resource.type_url):
            res.type_url = Primitive.to_proto(resource.type_url)
        if Primitive.to_proto(resource.value):
            res.value = Primitive.to_proto(resource.value)
        return res

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return None

        return InstanceUpdateRecipeStepsStatusDetails(
            type_url=Primitive.from_proto(resource.type_url),
            value=Primitive.from_proto(resource.value),
        )


class InstanceUpdateRecipeStepsStatusDetailsArray(object):
    @classmethod
    def to_proto(self, resources):
        if not resources:
            return resources
        return [InstanceUpdateRecipeStepsStatusDetails.to_proto(i) for i in resources]

    @classmethod
    def from_proto(self, resources):
        return [InstanceUpdateRecipeStepsStatusDetails.from_proto(i) for i in resources]


class InstanceUpdateRecipeStepsQuotaRequestDeltas(object):
    def __init__(
        self,
        metric_name: str = None,
        amount: int = None,
        quota_location_name: str = None,
    ):
        self.metric_name = metric_name
        self.amount = amount
        self.quota_location_name = quota_location_name

    @classmethod
    def to_proto(self, resource):
        if not resource:
            return None

        res = instance_pb2.Tier2AlphaInstanceUpdateRecipeStepsQuotaRequestDeltas()
        if Primitive.to_proto(resource.metric_name):
            res.metric_name = Primitive.to_proto(resource.metric_name)
        if Primitive.to_proto(resource.amount):
            res.amount = Primitive.to_proto(resource.amount)
        if Primitive.to_proto(resource.quota_location_name):
            res.quota_location_name = Primitive.to_proto(resource.quota_location_name)
        return res

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return None

        return InstanceUpdateRecipeStepsQuotaRequestDeltas(
            metric_name=Primitive.from_proto(resource.metric_name),
            amount=Primitive.from_proto(resource.amount),
            quota_location_name=Primitive.from_proto(resource.quota_location_name),
        )


class InstanceUpdateRecipeStepsQuotaRequestDeltasArray(object):
    @classmethod
    def to_proto(self, resources):
        if not resources:
            return resources
        return [
            InstanceUpdateRecipeStepsQuotaRequestDeltas.to_proto(i) for i in resources
        ]

    @classmethod
    def from_proto(self, resources):
        return [
            InstanceUpdateRecipeStepsQuotaRequestDeltas.from_proto(i) for i in resources
        ]


class InstanceUpdateRecipeStepsPreprocessUpdate(object):
    def __init__(
        self, latency_slo_bucket_name: str = None, public_operation_metadata: str = None
    ):
        self.latency_slo_bucket_name = latency_slo_bucket_name
        self.public_operation_metadata = public_operation_metadata

    @classmethod
    def to_proto(self, resource):
        if not resource:
            return None

        res = instance_pb2.Tier2AlphaInstanceUpdateRecipeStepsPreprocessUpdate()
        if Primitive.to_proto(resource.latency_slo_bucket_name):
            res.latency_slo_bucket_name = Primitive.to_proto(
                resource.latency_slo_bucket_name
            )
        if Primitive.to_proto(resource.public_operation_metadata):
            res.public_operation_metadata = Primitive.to_proto(
                resource.public_operation_metadata
            )
        return res

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return None

        return InstanceUpdateRecipeStepsPreprocessUpdate(
            latency_slo_bucket_name=Primitive.from_proto(
                resource.latency_slo_bucket_name
            ),
            public_operation_metadata=Primitive.from_proto(
                resource.public_operation_metadata
            ),
        )


class InstanceUpdateRecipeStepsPreprocessUpdateArray(object):
    @classmethod
    def to_proto(self, resources):
        if not resources:
            return resources
        return [
            InstanceUpdateRecipeStepsPreprocessUpdate.to_proto(i) for i in resources
        ]

    @classmethod
    def from_proto(self, resources):
        return [
            InstanceUpdateRecipeStepsPreprocessUpdate.from_proto(i) for i in resources
        ]


class InstanceUpdateRecipeStepsRequestedTenantProject(object):
    def __init__(self, tag: str = None, folder: str = None, scope: str = None):
        self.tag = tag
        self.folder = folder
        self.scope = scope

    @classmethod
    def to_proto(self, resource):
        if not resource:
            return None

        res = instance_pb2.Tier2AlphaInstanceUpdateRecipeStepsRequestedTenantProject()
        if Primitive.to_proto(resource.tag):
            res.tag = Primitive.to_proto(resource.tag)
        if Primitive.to_proto(resource.folder):
            res.folder = Primitive.to_proto(resource.folder)
        if InstanceUpdateRecipeStepsRequestedTenantProjectScopeEnum.to_proto(
            resource.scope
        ):
            res.scope = InstanceUpdateRecipeStepsRequestedTenantProjectScopeEnum.to_proto(
                resource.scope
            )
        return res

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return None

        return InstanceUpdateRecipeStepsRequestedTenantProject(
            tag=Primitive.from_proto(resource.tag),
            folder=Primitive.from_proto(resource.folder),
            scope=InstanceUpdateRecipeStepsRequestedTenantProjectScopeEnum.from_proto(
                resource.scope
            ),
        )


class InstanceUpdateRecipeStepsRequestedTenantProjectArray(object):
    @classmethod
    def to_proto(self, resources):
        if not resources:
            return resources
        return [
            InstanceUpdateRecipeStepsRequestedTenantProject.to_proto(i)
            for i in resources
        ]

    @classmethod
    def from_proto(self, resources):
        return [
            InstanceUpdateRecipeStepsRequestedTenantProject.from_proto(i)
            for i in resources
        ]


class InstanceUpdateRecipeStepsPermissionsInfo(object):
    def __init__(
        self,
        policy_name: dict = None,
        iam_permissions: list = None,
        resource_path: str = None,
        api_attrs: dict = None,
        policy_name_mode: str = None,
        resource: dict = None,
    ):
        self.policy_name = policy_name
        self.iam_permissions = iam_permissions
        self.resource_path = resource_path
        self.api_attrs = api_attrs
        self.policy_name_mode = policy_name_mode
        self.resource = resource

    @classmethod
    def to_proto(self, resource):
        if not resource:
            return None

        res = instance_pb2.Tier2AlphaInstanceUpdateRecipeStepsPermissionsInfo()
        if InstanceUpdateRecipeStepsPermissionsInfoPolicyName.to_proto(
            resource.policy_name
        ):
            res.policy_name.CopyFrom(
                InstanceUpdateRecipeStepsPermissionsInfoPolicyName.to_proto(
                    resource.policy_name
                )
            )
        else:
            res.ClearField("policy_name")
        if InstanceUpdateRecipeStepsPermissionsInfoIamPermissionsArray.to_proto(
            resource.iam_permissions
        ):
            res.iam_permissions.extend(
                InstanceUpdateRecipeStepsPermissionsInfoIamPermissionsArray.to_proto(
                    resource.iam_permissions
                )
            )
        if Primitive.to_proto(resource.resource_path):
            res.resource_path = Primitive.to_proto(resource.resource_path)
        if InstanceGoogleprotobufstruct.to_proto(resource.api_attrs):
            res.api_attrs.CopyFrom(
                InstanceGoogleprotobufstruct.to_proto(resource.api_attrs)
            )
        else:
            res.ClearField("api_attrs")
        if InstanceUpdateRecipeStepsPermissionsInfoPolicyNameModeEnum.to_proto(
            resource.policy_name_mode
        ):
            res.policy_name_mode = InstanceUpdateRecipeStepsPermissionsInfoPolicyNameModeEnum.to_proto(
                resource.policy_name_mode
            )
        if InstanceUpdateRecipeStepsPermissionsInfoResource.to_proto(resource.resource):
            res.resource.CopyFrom(
                InstanceUpdateRecipeStepsPermissionsInfoResource.to_proto(
                    resource.resource
                )
            )
        else:
            res.ClearField("resource")
        return res

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return None

        return InstanceUpdateRecipeStepsPermissionsInfo(
            policy_name=InstanceUpdateRecipeStepsPermissionsInfoPolicyName.from_proto(
                resource.policy_name
            ),
            iam_permissions=InstanceUpdateRecipeStepsPermissionsInfoIamPermissionsArray.from_proto(
                resource.iam_permissions
            ),
            resource_path=Primitive.from_proto(resource.resource_path),
            api_attrs=InstanceGoogleprotobufstruct.from_proto(resource.api_attrs),
            policy_name_mode=InstanceUpdateRecipeStepsPermissionsInfoPolicyNameModeEnum.from_proto(
                resource.policy_name_mode
            ),
            resource=InstanceUpdateRecipeStepsPermissionsInfoResource.from_proto(
                resource.resource
            ),
        )


class InstanceUpdateRecipeStepsPermissionsInfoArray(object):
    @classmethod
    def to_proto(self, resources):
        if not resources:
            return resources
        return [InstanceUpdateRecipeStepsPermissionsInfo.to_proto(i) for i in resources]

    @classmethod
    def from_proto(self, resources):
        return [
            InstanceUpdateRecipeStepsPermissionsInfo.from_proto(i) for i in resources
        ]


class InstanceUpdateRecipeStepsPermissionsInfoPolicyName(object):
    def __init__(self, type: str = None, id: str = None, region: str = None):
        self.type = type
        self.id = id
        self.region = region

    @classmethod
    def to_proto(self, resource):
        if not resource:
            return None

        res = (
            instance_pb2.Tier2AlphaInstanceUpdateRecipeStepsPermissionsInfoPolicyName()
        )
        if Primitive.to_proto(resource.type):
            res.type = Primitive.to_proto(resource.type)
        if Primitive.to_proto(resource.id):
            res.id = Primitive.to_proto(resource.id)
        if Primitive.to_proto(resource.region):
            res.region = Primitive.to_proto(resource.region)
        return res

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return None

        return InstanceUpdateRecipeStepsPermissionsInfoPolicyName(
            type=Primitive.from_proto(resource.type),
            id=Primitive.from_proto(resource.id),
            region=Primitive.from_proto(resource.region),
        )


class InstanceUpdateRecipeStepsPermissionsInfoPolicyNameArray(object):
    @classmethod
    def to_proto(self, resources):
        if not resources:
            return resources
        return [
            InstanceUpdateRecipeStepsPermissionsInfoPolicyName.to_proto(i)
            for i in resources
        ]

    @classmethod
    def from_proto(self, resources):
        return [
            InstanceUpdateRecipeStepsPermissionsInfoPolicyName.from_proto(i)
            for i in resources
        ]


class InstanceUpdateRecipeStepsPermissionsInfoIamPermissions(object):
    def __init__(self, permission: str = None):
        self.permission = permission

    @classmethod
    def to_proto(self, resource):
        if not resource:
            return None

        res = (
            instance_pb2.Tier2AlphaInstanceUpdateRecipeStepsPermissionsInfoIamPermissions()
        )
        if Primitive.to_proto(resource.permission):
            res.permission = Primitive.to_proto(resource.permission)
        return res

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return None

        return InstanceUpdateRecipeStepsPermissionsInfoIamPermissions(
            permission=Primitive.from_proto(resource.permission),
        )


class InstanceUpdateRecipeStepsPermissionsInfoIamPermissionsArray(object):
    @classmethod
    def to_proto(self, resources):
        if not resources:
            return resources
        return [
            InstanceUpdateRecipeStepsPermissionsInfoIamPermissions.to_proto(i)
            for i in resources
        ]

    @classmethod
    def from_proto(self, resources):
        return [
            InstanceUpdateRecipeStepsPermissionsInfoIamPermissions.from_proto(i)
            for i in resources
        ]


class InstanceUpdateRecipeStepsPermissionsInfoResource(object):
    def __init__(
        self,
        name: str = None,
        type: str = None,
        service: str = None,
        labels: dict = None,
    ):
        self.name = name
        self.type = type
        self.service = service
        self.labels = labels

    @classmethod
    def to_proto(self, resource):
        if not resource:
            return None

        res = instance_pb2.Tier2AlphaInstanceUpdateRecipeStepsPermissionsInfoResource()
        if Primitive.to_proto(resource.name):
            res.name = Primitive.to_proto(resource.name)
        if Primitive.to_proto(resource.type):
            res.type = Primitive.to_proto(resource.type)
        if Primitive.to_proto(resource.service):
            res.service = Primitive.to_proto(resource.service)
        if Primitive.to_proto(resource.labels):
            res.labels = Primitive.to_proto(resource.labels)
        return res

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return None

        return InstanceUpdateRecipeStepsPermissionsInfoResource(
            name=Primitive.from_proto(resource.name),
            type=Primitive.from_proto(resource.type),
            service=Primitive.from_proto(resource.service),
            labels=Primitive.from_proto(resource.labels),
        )


class InstanceUpdateRecipeStepsPermissionsInfoResourceArray(object):
    @classmethod
    def to_proto(self, resources):
        if not resources:
            return resources
        return [
            InstanceUpdateRecipeStepsPermissionsInfoResource.to_proto(i)
            for i in resources
        ]

    @classmethod
    def from_proto(self, resources):
        return [
            InstanceUpdateRecipeStepsPermissionsInfoResource.from_proto(i)
            for i in resources
        ]


class InstanceUpdateRecipeStepsKeyNotificationsUpdate(object):
    def __init__(self, key_notifications_info: dict = None):
        self.key_notifications_info = key_notifications_info

    @classmethod
    def to_proto(self, resource):
        if not resource:
            return None

        res = instance_pb2.Tier2AlphaInstanceUpdateRecipeStepsKeyNotificationsUpdate()
        if InstanceUpdateRecipeStepsKeyNotificationsUpdateKeyNotificationsInfo.to_proto(
            resource.key_notifications_info
        ):
            res.key_notifications_info.CopyFrom(
                InstanceUpdateRecipeStepsKeyNotificationsUpdateKeyNotificationsInfo.to_proto(
                    resource.key_notifications_info
                )
            )
        else:
            res.ClearField("key_notifications_info")
        return res

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return None

        return InstanceUpdateRecipeStepsKeyNotificationsUpdate(
            key_notifications_info=InstanceUpdateRecipeStepsKeyNotificationsUpdateKeyNotificationsInfo.from_proto(
                resource.key_notifications_info
            ),
        )


class InstanceUpdateRecipeStepsKeyNotificationsUpdateArray(object):
    @classmethod
    def to_proto(self, resources):
        if not resources:
            return resources
        return [
            InstanceUpdateRecipeStepsKeyNotificationsUpdate.to_proto(i)
            for i in resources
        ]

    @classmethod
    def from_proto(self, resources):
        return [
            InstanceUpdateRecipeStepsKeyNotificationsUpdate.from_proto(i)
            for i in resources
        ]


class InstanceUpdateRecipeStepsKeyNotificationsUpdateKeyNotificationsInfo(object):
    def __init__(
        self,
        data_version: int = None,
        delegate: str = None,
        key_notification_configs: list = None,
    ):
        self.data_version = data_version
        self.delegate = delegate
        self.key_notification_configs = key_notification_configs

    @classmethod
    def to_proto(self, resource):
        if not resource:
            return None

        res = (
            instance_pb2.Tier2AlphaInstanceUpdateRecipeStepsKeyNotificationsUpdateKeyNotificationsInfo()
        )
        if Primitive.to_proto(resource.data_version):
            res.data_version = Primitive.to_proto(resource.data_version)
        if Primitive.to_proto(resource.delegate):
            res.delegate = Primitive.to_proto(resource.delegate)
        if InstanceUpdateRecipeStepsKeyNotificationsUpdateKeyNotificationsInfoKeyNotificationConfigsArray.to_proto(
            resource.key_notification_configs
        ):
            res.key_notification_configs.extend(
                InstanceUpdateRecipeStepsKeyNotificationsUpdateKeyNotificationsInfoKeyNotificationConfigsArray.to_proto(
                    resource.key_notification_configs
                )
            )
        return res

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return None

        return InstanceUpdateRecipeStepsKeyNotificationsUpdateKeyNotificationsInfo(
            data_version=Primitive.from_proto(resource.data_version),
            delegate=Primitive.from_proto(resource.delegate),
            key_notification_configs=InstanceUpdateRecipeStepsKeyNotificationsUpdateKeyNotificationsInfoKeyNotificationConfigsArray.from_proto(
                resource.key_notification_configs
            ),
        )


class InstanceUpdateRecipeStepsKeyNotificationsUpdateKeyNotificationsInfoArray(object):
    @classmethod
    def to_proto(self, resources):
        if not resources:
            return resources
        return [
            InstanceUpdateRecipeStepsKeyNotificationsUpdateKeyNotificationsInfo.to_proto(
                i
            )
            for i in resources
        ]

    @classmethod
    def from_proto(self, resources):
        return [
            InstanceUpdateRecipeStepsKeyNotificationsUpdateKeyNotificationsInfo.from_proto(
                i
            )
            for i in resources
        ]


class InstanceUpdateRecipeStepsKeyNotificationsUpdateKeyNotificationsInfoKeyNotificationConfigs(
    object
):
    def __init__(
        self,
        key_or_version_name: str = None,
        grant: str = None,
        delegator_gaia_id: int = None,
    ):
        self.key_or_version_name = key_or_version_name
        self.grant = grant
        self.delegator_gaia_id = delegator_gaia_id

    @classmethod
    def to_proto(self, resource):
        if not resource:
            return None

        res = (
            instance_pb2.Tier2AlphaInstanceUpdateRecipeStepsKeyNotificationsUpdateKeyNotificationsInfoKeyNotificationConfigs()
        )
        if Primitive.to_proto(resource.key_or_version_name):
            res.key_or_version_name = Primitive.to_proto(resource.key_or_version_name)
        if Primitive.to_proto(resource.grant):
            res.grant = Primitive.to_proto(resource.grant)
        if Primitive.to_proto(resource.delegator_gaia_id):
            res.delegator_gaia_id = Primitive.to_proto(resource.delegator_gaia_id)
        return res

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return None

        return InstanceUpdateRecipeStepsKeyNotificationsUpdateKeyNotificationsInfoKeyNotificationConfigs(
            key_or_version_name=Primitive.from_proto(resource.key_or_version_name),
            grant=Primitive.from_proto(resource.grant),
            delegator_gaia_id=Primitive.from_proto(resource.delegator_gaia_id),
        )


class InstanceUpdateRecipeStepsKeyNotificationsUpdateKeyNotificationsInfoKeyNotificationConfigsArray(
    object
):
    @classmethod
    def to_proto(self, resources):
        if not resources:
            return resources
        return [
            InstanceUpdateRecipeStepsKeyNotificationsUpdateKeyNotificationsInfoKeyNotificationConfigs.to_proto(
                i
            )
            for i in resources
        ]

    @classmethod
    def from_proto(self, resources):
        return [
            InstanceUpdateRecipeStepsKeyNotificationsUpdateKeyNotificationsInfoKeyNotificationConfigs.from_proto(
                i
            )
            for i in resources
        ]


class InstancePreprocessResetRecipe(object):
    def __init__(
        self,
        steps: list = None,
        honor_cancel_request: bool = None,
        ignore_recipe_after: int = None,
        verify_deadline_seconds_below: float = None,
        populate_operation_result: bool = None,
        readonly_recipe_start_time: str = None,
        resource_names_stored_in_clh_with_delay: list = None,
        delay_to_store_resources_in_clh_db_nanos: int = None,
    ):
        self.steps = steps
        self.honor_cancel_request = honor_cancel_request
        self.ignore_recipe_after = ignore_recipe_after
        self.verify_deadline_seconds_below = verify_deadline_seconds_below
        self.populate_operation_result = populate_operation_result
        self.readonly_recipe_start_time = readonly_recipe_start_time
        self.resource_names_stored_in_clh_with_delay = (
            resource_names_stored_in_clh_with_delay
        )
        self.delay_to_store_resources_in_clh_db_nanos = (
            delay_to_store_resources_in_clh_db_nanos
        )

    @classmethod
    def to_proto(self, resource):
        if not resource:
            return None

        res = instance_pb2.Tier2AlphaInstancePreprocessResetRecipe()
        if InstancePreprocessResetRecipeStepsArray.to_proto(resource.steps):
            res.steps.extend(
                InstancePreprocessResetRecipeStepsArray.to_proto(resource.steps)
            )
        if Primitive.to_proto(resource.honor_cancel_request):
            res.honor_cancel_request = Primitive.to_proto(resource.honor_cancel_request)
        if Primitive.to_proto(resource.ignore_recipe_after):
            res.ignore_recipe_after = Primitive.to_proto(resource.ignore_recipe_after)
        if Primitive.to_proto(resource.verify_deadline_seconds_below):
            res.verify_deadline_seconds_below = Primitive.to_proto(
                resource.verify_deadline_seconds_below
            )
        if Primitive.to_proto(resource.populate_operation_result):
            res.populate_operation_result = Primitive.to_proto(
                resource.populate_operation_result
            )
        if Primitive.to_proto(resource.readonly_recipe_start_time):
            res.readonly_recipe_start_time = Primitive.to_proto(
                resource.readonly_recipe_start_time
            )
        if Primitive.to_proto(resource.resource_names_stored_in_clh_with_delay):
            res.resource_names_stored_in_clh_with_delay.extend(
                Primitive.to_proto(resource.resource_names_stored_in_clh_with_delay)
            )
        if Primitive.to_proto(resource.delay_to_store_resources_in_clh_db_nanos):
            res.delay_to_store_resources_in_clh_db_nanos = Primitive.to_proto(
                resource.delay_to_store_resources_in_clh_db_nanos
            )
        return res

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return None

        return InstancePreprocessResetRecipe(
            steps=InstancePreprocessResetRecipeStepsArray.from_proto(resource.steps),
            honor_cancel_request=Primitive.from_proto(resource.honor_cancel_request),
            ignore_recipe_after=Primitive.from_proto(resource.ignore_recipe_after),
            verify_deadline_seconds_below=Primitive.from_proto(
                resource.verify_deadline_seconds_below
            ),
            populate_operation_result=Primitive.from_proto(
                resource.populate_operation_result
            ),
            readonly_recipe_start_time=Primitive.from_proto(
                resource.readonly_recipe_start_time
            ),
            resource_names_stored_in_clh_with_delay=Primitive.from_proto(
                resource.resource_names_stored_in_clh_with_delay
            ),
            delay_to_store_resources_in_clh_db_nanos=Primitive.from_proto(
                resource.delay_to_store_resources_in_clh_db_nanos
            ),
        )


class InstancePreprocessResetRecipeArray(object):
    @classmethod
    def to_proto(self, resources):
        if not resources:
            return resources
        return [InstancePreprocessResetRecipe.to_proto(i) for i in resources]

    @classmethod
    def from_proto(self, resources):
        return [InstancePreprocessResetRecipe.from_proto(i) for i in resources]


class InstancePreprocessResetRecipeSteps(object):
    def __init__(
        self,
        relative_time: int = None,
        sleep_duration: int = None,
        action: str = None,
        status: dict = None,
        error_space: str = None,
        p4_service_account: str = None,
        resource_metadata_size: int = None,
        description: str = None,
        updated_repeat_operation_delay_sec: float = None,
        quota_request_deltas: list = None,
        preprocess_update: dict = None,
        public_operation_metadata: str = None,
        requested_tenant_project: dict = None,
        permissions_info: list = None,
        key_notifications_update: dict = None,
        clh_data_update_time: str = None,
        public_error_message: str = None,
    ):
        self.relative_time = relative_time
        self.sleep_duration = sleep_duration
        self.action = action
        self.status = status
        self.error_space = error_space
        self.p4_service_account = p4_service_account
        self.resource_metadata_size = resource_metadata_size
        self.description = description
        self.updated_repeat_operation_delay_sec = updated_repeat_operation_delay_sec
        self.quota_request_deltas = quota_request_deltas
        self.preprocess_update = preprocess_update
        self.public_operation_metadata = public_operation_metadata
        self.requested_tenant_project = requested_tenant_project
        self.permissions_info = permissions_info
        self.key_notifications_update = key_notifications_update
        self.clh_data_update_time = clh_data_update_time
        self.public_error_message = public_error_message

    @classmethod
    def to_proto(self, resource):
        if not resource:
            return None

        res = instance_pb2.Tier2AlphaInstancePreprocessResetRecipeSteps()
        if Primitive.to_proto(resource.relative_time):
            res.relative_time = Primitive.to_proto(resource.relative_time)
        if Primitive.to_proto(resource.sleep_duration):
            res.sleep_duration = Primitive.to_proto(resource.sleep_duration)
        if InstancePreprocessResetRecipeStepsActionEnum.to_proto(resource.action):
            res.action = InstancePreprocessResetRecipeStepsActionEnum.to_proto(
                resource.action
            )
        if InstancePreprocessResetRecipeStepsStatus.to_proto(resource.status):
            res.status.CopyFrom(
                InstancePreprocessResetRecipeStepsStatus.to_proto(resource.status)
            )
        else:
            res.ClearField("status")
        if Primitive.to_proto(resource.error_space):
            res.error_space = Primitive.to_proto(resource.error_space)
        if Primitive.to_proto(resource.p4_service_account):
            res.p4_service_account = Primitive.to_proto(resource.p4_service_account)
        if Primitive.to_proto(resource.resource_metadata_size):
            res.resource_metadata_size = Primitive.to_proto(
                resource.resource_metadata_size
            )
        if Primitive.to_proto(resource.description):
            res.description = Primitive.to_proto(resource.description)
        if Primitive.to_proto(resource.updated_repeat_operation_delay_sec):
            res.updated_repeat_operation_delay_sec = Primitive.to_proto(
                resource.updated_repeat_operation_delay_sec
            )
        if InstancePreprocessResetRecipeStepsQuotaRequestDeltasArray.to_proto(
            resource.quota_request_deltas
        ):
            res.quota_request_deltas.extend(
                InstancePreprocessResetRecipeStepsQuotaRequestDeltasArray.to_proto(
                    resource.quota_request_deltas
                )
            )
        if InstancePreprocessResetRecipeStepsPreprocessUpdate.to_proto(
            resource.preprocess_update
        ):
            res.preprocess_update.CopyFrom(
                InstancePreprocessResetRecipeStepsPreprocessUpdate.to_proto(
                    resource.preprocess_update
                )
            )
        else:
            res.ClearField("preprocess_update")
        if Primitive.to_proto(resource.public_operation_metadata):
            res.public_operation_metadata = Primitive.to_proto(
                resource.public_operation_metadata
            )
        if InstancePreprocessResetRecipeStepsRequestedTenantProject.to_proto(
            resource.requested_tenant_project
        ):
            res.requested_tenant_project.CopyFrom(
                InstancePreprocessResetRecipeStepsRequestedTenantProject.to_proto(
                    resource.requested_tenant_project
                )
            )
        else:
            res.ClearField("requested_tenant_project")
        if InstancePreprocessResetRecipeStepsPermissionsInfoArray.to_proto(
            resource.permissions_info
        ):
            res.permissions_info.extend(
                InstancePreprocessResetRecipeStepsPermissionsInfoArray.to_proto(
                    resource.permissions_info
                )
            )
        if InstancePreprocessResetRecipeStepsKeyNotificationsUpdate.to_proto(
            resource.key_notifications_update
        ):
            res.key_notifications_update.CopyFrom(
                InstancePreprocessResetRecipeStepsKeyNotificationsUpdate.to_proto(
                    resource.key_notifications_update
                )
            )
        else:
            res.ClearField("key_notifications_update")
        if Primitive.to_proto(resource.clh_data_update_time):
            res.clh_data_update_time = Primitive.to_proto(resource.clh_data_update_time)
        if Primitive.to_proto(resource.public_error_message):
            res.public_error_message = Primitive.to_proto(resource.public_error_message)
        return res

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return None

        return InstancePreprocessResetRecipeSteps(
            relative_time=Primitive.from_proto(resource.relative_time),
            sleep_duration=Primitive.from_proto(resource.sleep_duration),
            action=InstancePreprocessResetRecipeStepsActionEnum.from_proto(
                resource.action
            ),
            status=InstancePreprocessResetRecipeStepsStatus.from_proto(resource.status),
            error_space=Primitive.from_proto(resource.error_space),
            p4_service_account=Primitive.from_proto(resource.p4_service_account),
            resource_metadata_size=Primitive.from_proto(
                resource.resource_metadata_size
            ),
            description=Primitive.from_proto(resource.description),
            updated_repeat_operation_delay_sec=Primitive.from_proto(
                resource.updated_repeat_operation_delay_sec
            ),
            quota_request_deltas=InstancePreprocessResetRecipeStepsQuotaRequestDeltasArray.from_proto(
                resource.quota_request_deltas
            ),
            preprocess_update=InstancePreprocessResetRecipeStepsPreprocessUpdate.from_proto(
                resource.preprocess_update
            ),
            public_operation_metadata=Primitive.from_proto(
                resource.public_operation_metadata
            ),
            requested_tenant_project=InstancePreprocessResetRecipeStepsRequestedTenantProject.from_proto(
                resource.requested_tenant_project
            ),
            permissions_info=InstancePreprocessResetRecipeStepsPermissionsInfoArray.from_proto(
                resource.permissions_info
            ),
            key_notifications_update=InstancePreprocessResetRecipeStepsKeyNotificationsUpdate.from_proto(
                resource.key_notifications_update
            ),
            clh_data_update_time=Primitive.from_proto(resource.clh_data_update_time),
            public_error_message=Primitive.from_proto(resource.public_error_message),
        )


class InstancePreprocessResetRecipeStepsArray(object):
    @classmethod
    def to_proto(self, resources):
        if not resources:
            return resources
        return [InstancePreprocessResetRecipeSteps.to_proto(i) for i in resources]

    @classmethod
    def from_proto(self, resources):
        return [InstancePreprocessResetRecipeSteps.from_proto(i) for i in resources]


class InstancePreprocessResetRecipeStepsStatus(object):
    def __init__(self, code: int = None, message: str = None, details: list = None):
        self.code = code
        self.message = message
        self.details = details

    @classmethod
    def to_proto(self, resource):
        if not resource:
            return None

        res = instance_pb2.Tier2AlphaInstancePreprocessResetRecipeStepsStatus()
        if Primitive.to_proto(resource.code):
            res.code = Primitive.to_proto(resource.code)
        if Primitive.to_proto(resource.message):
            res.message = Primitive.to_proto(resource.message)
        if InstancePreprocessResetRecipeStepsStatusDetailsArray.to_proto(
            resource.details
        ):
            res.details.extend(
                InstancePreprocessResetRecipeStepsStatusDetailsArray.to_proto(
                    resource.details
                )
            )
        return res

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return None

        return InstancePreprocessResetRecipeStepsStatus(
            code=Primitive.from_proto(resource.code),
            message=Primitive.from_proto(resource.message),
            details=InstancePreprocessResetRecipeStepsStatusDetailsArray.from_proto(
                resource.details
            ),
        )


class InstancePreprocessResetRecipeStepsStatusArray(object):
    @classmethod
    def to_proto(self, resources):
        if not resources:
            return resources
        return [InstancePreprocessResetRecipeStepsStatus.to_proto(i) for i in resources]

    @classmethod
    def from_proto(self, resources):
        return [
            InstancePreprocessResetRecipeStepsStatus.from_proto(i) for i in resources
        ]


class InstancePreprocessResetRecipeStepsStatusDetails(object):
    def __init__(self, type_url: str = None, value: str = None):
        self.type_url = type_url
        self.value = value

    @classmethod
    def to_proto(self, resource):
        if not resource:
            return None

        res = instance_pb2.Tier2AlphaInstancePreprocessResetRecipeStepsStatusDetails()
        if Primitive.to_proto(resource.type_url):
            res.type_url = Primitive.to_proto(resource.type_url)
        if Primitive.to_proto(resource.value):
            res.value = Primitive.to_proto(resource.value)
        return res

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return None

        return InstancePreprocessResetRecipeStepsStatusDetails(
            type_url=Primitive.from_proto(resource.type_url),
            value=Primitive.from_proto(resource.value),
        )


class InstancePreprocessResetRecipeStepsStatusDetailsArray(object):
    @classmethod
    def to_proto(self, resources):
        if not resources:
            return resources
        return [
            InstancePreprocessResetRecipeStepsStatusDetails.to_proto(i)
            for i in resources
        ]

    @classmethod
    def from_proto(self, resources):
        return [
            InstancePreprocessResetRecipeStepsStatusDetails.from_proto(i)
            for i in resources
        ]


class InstancePreprocessResetRecipeStepsQuotaRequestDeltas(object):
    def __init__(
        self,
        metric_name: str = None,
        amount: int = None,
        quota_location_name: str = None,
    ):
        self.metric_name = metric_name
        self.amount = amount
        self.quota_location_name = quota_location_name

    @classmethod
    def to_proto(self, resource):
        if not resource:
            return None

        res = (
            instance_pb2.Tier2AlphaInstancePreprocessResetRecipeStepsQuotaRequestDeltas()
        )
        if Primitive.to_proto(resource.metric_name):
            res.metric_name = Primitive.to_proto(resource.metric_name)
        if Primitive.to_proto(resource.amount):
            res.amount = Primitive.to_proto(resource.amount)
        if Primitive.to_proto(resource.quota_location_name):
            res.quota_location_name = Primitive.to_proto(resource.quota_location_name)
        return res

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return None

        return InstancePreprocessResetRecipeStepsQuotaRequestDeltas(
            metric_name=Primitive.from_proto(resource.metric_name),
            amount=Primitive.from_proto(resource.amount),
            quota_location_name=Primitive.from_proto(resource.quota_location_name),
        )


class InstancePreprocessResetRecipeStepsQuotaRequestDeltasArray(object):
    @classmethod
    def to_proto(self, resources):
        if not resources:
            return resources
        return [
            InstancePreprocessResetRecipeStepsQuotaRequestDeltas.to_proto(i)
            for i in resources
        ]

    @classmethod
    def from_proto(self, resources):
        return [
            InstancePreprocessResetRecipeStepsQuotaRequestDeltas.from_proto(i)
            for i in resources
        ]


class InstancePreprocessResetRecipeStepsPreprocessUpdate(object):
    def __init__(
        self, latency_slo_bucket_name: str = None, public_operation_metadata: str = None
    ):
        self.latency_slo_bucket_name = latency_slo_bucket_name
        self.public_operation_metadata = public_operation_metadata

    @classmethod
    def to_proto(self, resource):
        if not resource:
            return None

        res = (
            instance_pb2.Tier2AlphaInstancePreprocessResetRecipeStepsPreprocessUpdate()
        )
        if Primitive.to_proto(resource.latency_slo_bucket_name):
            res.latency_slo_bucket_name = Primitive.to_proto(
                resource.latency_slo_bucket_name
            )
        if Primitive.to_proto(resource.public_operation_metadata):
            res.public_operation_metadata = Primitive.to_proto(
                resource.public_operation_metadata
            )
        return res

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return None

        return InstancePreprocessResetRecipeStepsPreprocessUpdate(
            latency_slo_bucket_name=Primitive.from_proto(
                resource.latency_slo_bucket_name
            ),
            public_operation_metadata=Primitive.from_proto(
                resource.public_operation_metadata
            ),
        )


class InstancePreprocessResetRecipeStepsPreprocessUpdateArray(object):
    @classmethod
    def to_proto(self, resources):
        if not resources:
            return resources
        return [
            InstancePreprocessResetRecipeStepsPreprocessUpdate.to_proto(i)
            for i in resources
        ]

    @classmethod
    def from_proto(self, resources):
        return [
            InstancePreprocessResetRecipeStepsPreprocessUpdate.from_proto(i)
            for i in resources
        ]


class InstancePreprocessResetRecipeStepsRequestedTenantProject(object):
    def __init__(self, tag: str = None, folder: str = None, scope: str = None):
        self.tag = tag
        self.folder = folder
        self.scope = scope

    @classmethod
    def to_proto(self, resource):
        if not resource:
            return None

        res = (
            instance_pb2.Tier2AlphaInstancePreprocessResetRecipeStepsRequestedTenantProject()
        )
        if Primitive.to_proto(resource.tag):
            res.tag = Primitive.to_proto(resource.tag)
        if Primitive.to_proto(resource.folder):
            res.folder = Primitive.to_proto(resource.folder)
        if InstancePreprocessResetRecipeStepsRequestedTenantProjectScopeEnum.to_proto(
            resource.scope
        ):
            res.scope = InstancePreprocessResetRecipeStepsRequestedTenantProjectScopeEnum.to_proto(
                resource.scope
            )
        return res

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return None

        return InstancePreprocessResetRecipeStepsRequestedTenantProject(
            tag=Primitive.from_proto(resource.tag),
            folder=Primitive.from_proto(resource.folder),
            scope=InstancePreprocessResetRecipeStepsRequestedTenantProjectScopeEnum.from_proto(
                resource.scope
            ),
        )


class InstancePreprocessResetRecipeStepsRequestedTenantProjectArray(object):
    @classmethod
    def to_proto(self, resources):
        if not resources:
            return resources
        return [
            InstancePreprocessResetRecipeStepsRequestedTenantProject.to_proto(i)
            for i in resources
        ]

    @classmethod
    def from_proto(self, resources):
        return [
            InstancePreprocessResetRecipeStepsRequestedTenantProject.from_proto(i)
            for i in resources
        ]


class InstancePreprocessResetRecipeStepsPermissionsInfo(object):
    def __init__(
        self,
        policy_name: dict = None,
        iam_permissions: list = None,
        resource_path: str = None,
        api_attrs: dict = None,
        policy_name_mode: str = None,
        resource: dict = None,
    ):
        self.policy_name = policy_name
        self.iam_permissions = iam_permissions
        self.resource_path = resource_path
        self.api_attrs = api_attrs
        self.policy_name_mode = policy_name_mode
        self.resource = resource

    @classmethod
    def to_proto(self, resource):
        if not resource:
            return None

        res = instance_pb2.Tier2AlphaInstancePreprocessResetRecipeStepsPermissionsInfo()
        if InstancePreprocessResetRecipeStepsPermissionsInfoPolicyName.to_proto(
            resource.policy_name
        ):
            res.policy_name.CopyFrom(
                InstancePreprocessResetRecipeStepsPermissionsInfoPolicyName.to_proto(
                    resource.policy_name
                )
            )
        else:
            res.ClearField("policy_name")
        if InstancePreprocessResetRecipeStepsPermissionsInfoIamPermissionsArray.to_proto(
            resource.iam_permissions
        ):
            res.iam_permissions.extend(
                InstancePreprocessResetRecipeStepsPermissionsInfoIamPermissionsArray.to_proto(
                    resource.iam_permissions
                )
            )
        if Primitive.to_proto(resource.resource_path):
            res.resource_path = Primitive.to_proto(resource.resource_path)
        if InstanceGoogleprotobufstruct.to_proto(resource.api_attrs):
            res.api_attrs.CopyFrom(
                InstanceGoogleprotobufstruct.to_proto(resource.api_attrs)
            )
        else:
            res.ClearField("api_attrs")
        if InstancePreprocessResetRecipeStepsPermissionsInfoPolicyNameModeEnum.to_proto(
            resource.policy_name_mode
        ):
            res.policy_name_mode = InstancePreprocessResetRecipeStepsPermissionsInfoPolicyNameModeEnum.to_proto(
                resource.policy_name_mode
            )
        if InstancePreprocessResetRecipeStepsPermissionsInfoResource.to_proto(
            resource.resource
        ):
            res.resource.CopyFrom(
                InstancePreprocessResetRecipeStepsPermissionsInfoResource.to_proto(
                    resource.resource
                )
            )
        else:
            res.ClearField("resource")
        return res

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return None

        return InstancePreprocessResetRecipeStepsPermissionsInfo(
            policy_name=InstancePreprocessResetRecipeStepsPermissionsInfoPolicyName.from_proto(
                resource.policy_name
            ),
            iam_permissions=InstancePreprocessResetRecipeStepsPermissionsInfoIamPermissionsArray.from_proto(
                resource.iam_permissions
            ),
            resource_path=Primitive.from_proto(resource.resource_path),
            api_attrs=InstanceGoogleprotobufstruct.from_proto(resource.api_attrs),
            policy_name_mode=InstancePreprocessResetRecipeStepsPermissionsInfoPolicyNameModeEnum.from_proto(
                resource.policy_name_mode
            ),
            resource=InstancePreprocessResetRecipeStepsPermissionsInfoResource.from_proto(
                resource.resource
            ),
        )


class InstancePreprocessResetRecipeStepsPermissionsInfoArray(object):
    @classmethod
    def to_proto(self, resources):
        if not resources:
            return resources
        return [
            InstancePreprocessResetRecipeStepsPermissionsInfo.to_proto(i)
            for i in resources
        ]

    @classmethod
    def from_proto(self, resources):
        return [
            InstancePreprocessResetRecipeStepsPermissionsInfo.from_proto(i)
            for i in resources
        ]


class InstancePreprocessResetRecipeStepsPermissionsInfoPolicyName(object):
    def __init__(self, type: str = None, id: str = None, region: str = None):
        self.type = type
        self.id = id
        self.region = region

    @classmethod
    def to_proto(self, resource):
        if not resource:
            return None

        res = (
            instance_pb2.Tier2AlphaInstancePreprocessResetRecipeStepsPermissionsInfoPolicyName()
        )
        if Primitive.to_proto(resource.type):
            res.type = Primitive.to_proto(resource.type)
        if Primitive.to_proto(resource.id):
            res.id = Primitive.to_proto(resource.id)
        if Primitive.to_proto(resource.region):
            res.region = Primitive.to_proto(resource.region)
        return res

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return None

        return InstancePreprocessResetRecipeStepsPermissionsInfoPolicyName(
            type=Primitive.from_proto(resource.type),
            id=Primitive.from_proto(resource.id),
            region=Primitive.from_proto(resource.region),
        )


class InstancePreprocessResetRecipeStepsPermissionsInfoPolicyNameArray(object):
    @classmethod
    def to_proto(self, resources):
        if not resources:
            return resources
        return [
            InstancePreprocessResetRecipeStepsPermissionsInfoPolicyName.to_proto(i)
            for i in resources
        ]

    @classmethod
    def from_proto(self, resources):
        return [
            InstancePreprocessResetRecipeStepsPermissionsInfoPolicyName.from_proto(i)
            for i in resources
        ]


class InstancePreprocessResetRecipeStepsPermissionsInfoIamPermissions(object):
    def __init__(self, permission: str = None):
        self.permission = permission

    @classmethod
    def to_proto(self, resource):
        if not resource:
            return None

        res = (
            instance_pb2.Tier2AlphaInstancePreprocessResetRecipeStepsPermissionsInfoIamPermissions()
        )
        if Primitive.to_proto(resource.permission):
            res.permission = Primitive.to_proto(resource.permission)
        return res

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return None

        return InstancePreprocessResetRecipeStepsPermissionsInfoIamPermissions(
            permission=Primitive.from_proto(resource.permission),
        )


class InstancePreprocessResetRecipeStepsPermissionsInfoIamPermissionsArray(object):
    @classmethod
    def to_proto(self, resources):
        if not resources:
            return resources
        return [
            InstancePreprocessResetRecipeStepsPermissionsInfoIamPermissions.to_proto(i)
            for i in resources
        ]

    @classmethod
    def from_proto(self, resources):
        return [
            InstancePreprocessResetRecipeStepsPermissionsInfoIamPermissions.from_proto(
                i
            )
            for i in resources
        ]


class InstancePreprocessResetRecipeStepsPermissionsInfoResource(object):
    def __init__(
        self,
        name: str = None,
        type: str = None,
        service: str = None,
        labels: dict = None,
    ):
        self.name = name
        self.type = type
        self.service = service
        self.labels = labels

    @classmethod
    def to_proto(self, resource):
        if not resource:
            return None

        res = (
            instance_pb2.Tier2AlphaInstancePreprocessResetRecipeStepsPermissionsInfoResource()
        )
        if Primitive.to_proto(resource.name):
            res.name = Primitive.to_proto(resource.name)
        if Primitive.to_proto(resource.type):
            res.type = Primitive.to_proto(resource.type)
        if Primitive.to_proto(resource.service):
            res.service = Primitive.to_proto(resource.service)
        if Primitive.to_proto(resource.labels):
            res.labels = Primitive.to_proto(resource.labels)
        return res

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return None

        return InstancePreprocessResetRecipeStepsPermissionsInfoResource(
            name=Primitive.from_proto(resource.name),
            type=Primitive.from_proto(resource.type),
            service=Primitive.from_proto(resource.service),
            labels=Primitive.from_proto(resource.labels),
        )


class InstancePreprocessResetRecipeStepsPermissionsInfoResourceArray(object):
    @classmethod
    def to_proto(self, resources):
        if not resources:
            return resources
        return [
            InstancePreprocessResetRecipeStepsPermissionsInfoResource.to_proto(i)
            for i in resources
        ]

    @classmethod
    def from_proto(self, resources):
        return [
            InstancePreprocessResetRecipeStepsPermissionsInfoResource.from_proto(i)
            for i in resources
        ]


class InstancePreprocessResetRecipeStepsKeyNotificationsUpdate(object):
    def __init__(self, key_notifications_info: dict = None):
        self.key_notifications_info = key_notifications_info

    @classmethod
    def to_proto(self, resource):
        if not resource:
            return None

        res = (
            instance_pb2.Tier2AlphaInstancePreprocessResetRecipeStepsKeyNotificationsUpdate()
        )
        if InstancePreprocessResetRecipeStepsKeyNotificationsUpdateKeyNotificationsInfo.to_proto(
            resource.key_notifications_info
        ):
            res.key_notifications_info.CopyFrom(
                InstancePreprocessResetRecipeStepsKeyNotificationsUpdateKeyNotificationsInfo.to_proto(
                    resource.key_notifications_info
                )
            )
        else:
            res.ClearField("key_notifications_info")
        return res

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return None

        return InstancePreprocessResetRecipeStepsKeyNotificationsUpdate(
            key_notifications_info=InstancePreprocessResetRecipeStepsKeyNotificationsUpdateKeyNotificationsInfo.from_proto(
                resource.key_notifications_info
            ),
        )


class InstancePreprocessResetRecipeStepsKeyNotificationsUpdateArray(object):
    @classmethod
    def to_proto(self, resources):
        if not resources:
            return resources
        return [
            InstancePreprocessResetRecipeStepsKeyNotificationsUpdate.to_proto(i)
            for i in resources
        ]

    @classmethod
    def from_proto(self, resources):
        return [
            InstancePreprocessResetRecipeStepsKeyNotificationsUpdate.from_proto(i)
            for i in resources
        ]


class InstancePreprocessResetRecipeStepsKeyNotificationsUpdateKeyNotificationsInfo(
    object
):
    def __init__(
        self,
        data_version: int = None,
        delegate: str = None,
        key_notification_configs: list = None,
    ):
        self.data_version = data_version
        self.delegate = delegate
        self.key_notification_configs = key_notification_configs

    @classmethod
    def to_proto(self, resource):
        if not resource:
            return None

        res = (
            instance_pb2.Tier2AlphaInstancePreprocessResetRecipeStepsKeyNotificationsUpdateKeyNotificationsInfo()
        )
        if Primitive.to_proto(resource.data_version):
            res.data_version = Primitive.to_proto(resource.data_version)
        if Primitive.to_proto(resource.delegate):
            res.delegate = Primitive.to_proto(resource.delegate)
        if InstancePreprocessResetRecipeStepsKeyNotificationsUpdateKeyNotificationsInfoKeyNotificationConfigsArray.to_proto(
            resource.key_notification_configs
        ):
            res.key_notification_configs.extend(
                InstancePreprocessResetRecipeStepsKeyNotificationsUpdateKeyNotificationsInfoKeyNotificationConfigsArray.to_proto(
                    resource.key_notification_configs
                )
            )
        return res

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return None

        return InstancePreprocessResetRecipeStepsKeyNotificationsUpdateKeyNotificationsInfo(
            data_version=Primitive.from_proto(resource.data_version),
            delegate=Primitive.from_proto(resource.delegate),
            key_notification_configs=InstancePreprocessResetRecipeStepsKeyNotificationsUpdateKeyNotificationsInfoKeyNotificationConfigsArray.from_proto(
                resource.key_notification_configs
            ),
        )


class InstancePreprocessResetRecipeStepsKeyNotificationsUpdateKeyNotificationsInfoArray(
    object
):
    @classmethod
    def to_proto(self, resources):
        if not resources:
            return resources
        return [
            InstancePreprocessResetRecipeStepsKeyNotificationsUpdateKeyNotificationsInfo.to_proto(
                i
            )
            for i in resources
        ]

    @classmethod
    def from_proto(self, resources):
        return [
            InstancePreprocessResetRecipeStepsKeyNotificationsUpdateKeyNotificationsInfo.from_proto(
                i
            )
            for i in resources
        ]


class InstancePreprocessResetRecipeStepsKeyNotificationsUpdateKeyNotificationsInfoKeyNotificationConfigs(
    object
):
    def __init__(
        self,
        key_or_version_name: str = None,
        grant: str = None,
        delegator_gaia_id: int = None,
    ):
        self.key_or_version_name = key_or_version_name
        self.grant = grant
        self.delegator_gaia_id = delegator_gaia_id

    @classmethod
    def to_proto(self, resource):
        if not resource:
            return None

        res = (
            instance_pb2.Tier2AlphaInstancePreprocessResetRecipeStepsKeyNotificationsUpdateKeyNotificationsInfoKeyNotificationConfigs()
        )
        if Primitive.to_proto(resource.key_or_version_name):
            res.key_or_version_name = Primitive.to_proto(resource.key_or_version_name)
        if Primitive.to_proto(resource.grant):
            res.grant = Primitive.to_proto(resource.grant)
        if Primitive.to_proto(resource.delegator_gaia_id):
            res.delegator_gaia_id = Primitive.to_proto(resource.delegator_gaia_id)
        return res

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return None

        return InstancePreprocessResetRecipeStepsKeyNotificationsUpdateKeyNotificationsInfoKeyNotificationConfigs(
            key_or_version_name=Primitive.from_proto(resource.key_or_version_name),
            grant=Primitive.from_proto(resource.grant),
            delegator_gaia_id=Primitive.from_proto(resource.delegator_gaia_id),
        )


class InstancePreprocessResetRecipeStepsKeyNotificationsUpdateKeyNotificationsInfoKeyNotificationConfigsArray(
    object
):
    @classmethod
    def to_proto(self, resources):
        if not resources:
            return resources
        return [
            InstancePreprocessResetRecipeStepsKeyNotificationsUpdateKeyNotificationsInfoKeyNotificationConfigs.to_proto(
                i
            )
            for i in resources
        ]

    @classmethod
    def from_proto(self, resources):
        return [
            InstancePreprocessResetRecipeStepsKeyNotificationsUpdateKeyNotificationsInfoKeyNotificationConfigs.from_proto(
                i
            )
            for i in resources
        ]


class InstanceResetRecipe(object):
    def __init__(
        self,
        steps: list = None,
        honor_cancel_request: bool = None,
        ignore_recipe_after: int = None,
        verify_deadline_seconds_below: float = None,
        populate_operation_result: bool = None,
        readonly_recipe_start_time: str = None,
        resource_names_stored_in_clh_with_delay: list = None,
        delay_to_store_resources_in_clh_db_nanos: int = None,
    ):
        self.steps = steps
        self.honor_cancel_request = honor_cancel_request
        self.ignore_recipe_after = ignore_recipe_after
        self.verify_deadline_seconds_below = verify_deadline_seconds_below
        self.populate_operation_result = populate_operation_result
        self.readonly_recipe_start_time = readonly_recipe_start_time
        self.resource_names_stored_in_clh_with_delay = (
            resource_names_stored_in_clh_with_delay
        )
        self.delay_to_store_resources_in_clh_db_nanos = (
            delay_to_store_resources_in_clh_db_nanos
        )

    @classmethod
    def to_proto(self, resource):
        if not resource:
            return None

        res = instance_pb2.Tier2AlphaInstanceResetRecipe()
        if InstanceResetRecipeStepsArray.to_proto(resource.steps):
            res.steps.extend(InstanceResetRecipeStepsArray.to_proto(resource.steps))
        if Primitive.to_proto(resource.honor_cancel_request):
            res.honor_cancel_request = Primitive.to_proto(resource.honor_cancel_request)
        if Primitive.to_proto(resource.ignore_recipe_after):
            res.ignore_recipe_after = Primitive.to_proto(resource.ignore_recipe_after)
        if Primitive.to_proto(resource.verify_deadline_seconds_below):
            res.verify_deadline_seconds_below = Primitive.to_proto(
                resource.verify_deadline_seconds_below
            )
        if Primitive.to_proto(resource.populate_operation_result):
            res.populate_operation_result = Primitive.to_proto(
                resource.populate_operation_result
            )
        if Primitive.to_proto(resource.readonly_recipe_start_time):
            res.readonly_recipe_start_time = Primitive.to_proto(
                resource.readonly_recipe_start_time
            )
        if Primitive.to_proto(resource.resource_names_stored_in_clh_with_delay):
            res.resource_names_stored_in_clh_with_delay.extend(
                Primitive.to_proto(resource.resource_names_stored_in_clh_with_delay)
            )
        if Primitive.to_proto(resource.delay_to_store_resources_in_clh_db_nanos):
            res.delay_to_store_resources_in_clh_db_nanos = Primitive.to_proto(
                resource.delay_to_store_resources_in_clh_db_nanos
            )
        return res

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return None

        return InstanceResetRecipe(
            steps=InstanceResetRecipeStepsArray.from_proto(resource.steps),
            honor_cancel_request=Primitive.from_proto(resource.honor_cancel_request),
            ignore_recipe_after=Primitive.from_proto(resource.ignore_recipe_after),
            verify_deadline_seconds_below=Primitive.from_proto(
                resource.verify_deadline_seconds_below
            ),
            populate_operation_result=Primitive.from_proto(
                resource.populate_operation_result
            ),
            readonly_recipe_start_time=Primitive.from_proto(
                resource.readonly_recipe_start_time
            ),
            resource_names_stored_in_clh_with_delay=Primitive.from_proto(
                resource.resource_names_stored_in_clh_with_delay
            ),
            delay_to_store_resources_in_clh_db_nanos=Primitive.from_proto(
                resource.delay_to_store_resources_in_clh_db_nanos
            ),
        )


class InstanceResetRecipeArray(object):
    @classmethod
    def to_proto(self, resources):
        if not resources:
            return resources
        return [InstanceResetRecipe.to_proto(i) for i in resources]

    @classmethod
    def from_proto(self, resources):
        return [InstanceResetRecipe.from_proto(i) for i in resources]


class InstanceResetRecipeSteps(object):
    def __init__(
        self,
        relative_time: int = None,
        sleep_duration: int = None,
        action: str = None,
        status: dict = None,
        error_space: str = None,
        p4_service_account: str = None,
        resource_metadata_size: int = None,
        description: str = None,
        updated_repeat_operation_delay_sec: float = None,
        quota_request_deltas: list = None,
        preprocess_update: dict = None,
        public_operation_metadata: str = None,
        requested_tenant_project: dict = None,
        permissions_info: list = None,
        key_notifications_update: dict = None,
        clh_data_update_time: str = None,
        public_error_message: str = None,
    ):
        self.relative_time = relative_time
        self.sleep_duration = sleep_duration
        self.action = action
        self.status = status
        self.error_space = error_space
        self.p4_service_account = p4_service_account
        self.resource_metadata_size = resource_metadata_size
        self.description = description
        self.updated_repeat_operation_delay_sec = updated_repeat_operation_delay_sec
        self.quota_request_deltas = quota_request_deltas
        self.preprocess_update = preprocess_update
        self.public_operation_metadata = public_operation_metadata
        self.requested_tenant_project = requested_tenant_project
        self.permissions_info = permissions_info
        self.key_notifications_update = key_notifications_update
        self.clh_data_update_time = clh_data_update_time
        self.public_error_message = public_error_message

    @classmethod
    def to_proto(self, resource):
        if not resource:
            return None

        res = instance_pb2.Tier2AlphaInstanceResetRecipeSteps()
        if Primitive.to_proto(resource.relative_time):
            res.relative_time = Primitive.to_proto(resource.relative_time)
        if Primitive.to_proto(resource.sleep_duration):
            res.sleep_duration = Primitive.to_proto(resource.sleep_duration)
        if InstanceResetRecipeStepsActionEnum.to_proto(resource.action):
            res.action = InstanceResetRecipeStepsActionEnum.to_proto(resource.action)
        if InstanceResetRecipeStepsStatus.to_proto(resource.status):
            res.status.CopyFrom(
                InstanceResetRecipeStepsStatus.to_proto(resource.status)
            )
        else:
            res.ClearField("status")
        if Primitive.to_proto(resource.error_space):
            res.error_space = Primitive.to_proto(resource.error_space)
        if Primitive.to_proto(resource.p4_service_account):
            res.p4_service_account = Primitive.to_proto(resource.p4_service_account)
        if Primitive.to_proto(resource.resource_metadata_size):
            res.resource_metadata_size = Primitive.to_proto(
                resource.resource_metadata_size
            )
        if Primitive.to_proto(resource.description):
            res.description = Primitive.to_proto(resource.description)
        if Primitive.to_proto(resource.updated_repeat_operation_delay_sec):
            res.updated_repeat_operation_delay_sec = Primitive.to_proto(
                resource.updated_repeat_operation_delay_sec
            )
        if InstanceResetRecipeStepsQuotaRequestDeltasArray.to_proto(
            resource.quota_request_deltas
        ):
            res.quota_request_deltas.extend(
                InstanceResetRecipeStepsQuotaRequestDeltasArray.to_proto(
                    resource.quota_request_deltas
                )
            )
        if InstanceResetRecipeStepsPreprocessUpdate.to_proto(
            resource.preprocess_update
        ):
            res.preprocess_update.CopyFrom(
                InstanceResetRecipeStepsPreprocessUpdate.to_proto(
                    resource.preprocess_update
                )
            )
        else:
            res.ClearField("preprocess_update")
        if Primitive.to_proto(resource.public_operation_metadata):
            res.public_operation_metadata = Primitive.to_proto(
                resource.public_operation_metadata
            )
        if InstanceResetRecipeStepsRequestedTenantProject.to_proto(
            resource.requested_tenant_project
        ):
            res.requested_tenant_project.CopyFrom(
                InstanceResetRecipeStepsRequestedTenantProject.to_proto(
                    resource.requested_tenant_project
                )
            )
        else:
            res.ClearField("requested_tenant_project")
        if InstanceResetRecipeStepsPermissionsInfoArray.to_proto(
            resource.permissions_info
        ):
            res.permissions_info.extend(
                InstanceResetRecipeStepsPermissionsInfoArray.to_proto(
                    resource.permissions_info
                )
            )
        if InstanceResetRecipeStepsKeyNotificationsUpdate.to_proto(
            resource.key_notifications_update
        ):
            res.key_notifications_update.CopyFrom(
                InstanceResetRecipeStepsKeyNotificationsUpdate.to_proto(
                    resource.key_notifications_update
                )
            )
        else:
            res.ClearField("key_notifications_update")
        if Primitive.to_proto(resource.clh_data_update_time):
            res.clh_data_update_time = Primitive.to_proto(resource.clh_data_update_time)
        if Primitive.to_proto(resource.public_error_message):
            res.public_error_message = Primitive.to_proto(resource.public_error_message)
        return res

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return None

        return InstanceResetRecipeSteps(
            relative_time=Primitive.from_proto(resource.relative_time),
            sleep_duration=Primitive.from_proto(resource.sleep_duration),
            action=InstanceResetRecipeStepsActionEnum.from_proto(resource.action),
            status=InstanceResetRecipeStepsStatus.from_proto(resource.status),
            error_space=Primitive.from_proto(resource.error_space),
            p4_service_account=Primitive.from_proto(resource.p4_service_account),
            resource_metadata_size=Primitive.from_proto(
                resource.resource_metadata_size
            ),
            description=Primitive.from_proto(resource.description),
            updated_repeat_operation_delay_sec=Primitive.from_proto(
                resource.updated_repeat_operation_delay_sec
            ),
            quota_request_deltas=InstanceResetRecipeStepsQuotaRequestDeltasArray.from_proto(
                resource.quota_request_deltas
            ),
            preprocess_update=InstanceResetRecipeStepsPreprocessUpdate.from_proto(
                resource.preprocess_update
            ),
            public_operation_metadata=Primitive.from_proto(
                resource.public_operation_metadata
            ),
            requested_tenant_project=InstanceResetRecipeStepsRequestedTenantProject.from_proto(
                resource.requested_tenant_project
            ),
            permissions_info=InstanceResetRecipeStepsPermissionsInfoArray.from_proto(
                resource.permissions_info
            ),
            key_notifications_update=InstanceResetRecipeStepsKeyNotificationsUpdate.from_proto(
                resource.key_notifications_update
            ),
            clh_data_update_time=Primitive.from_proto(resource.clh_data_update_time),
            public_error_message=Primitive.from_proto(resource.public_error_message),
        )


class InstanceResetRecipeStepsArray(object):
    @classmethod
    def to_proto(self, resources):
        if not resources:
            return resources
        return [InstanceResetRecipeSteps.to_proto(i) for i in resources]

    @classmethod
    def from_proto(self, resources):
        return [InstanceResetRecipeSteps.from_proto(i) for i in resources]


class InstanceResetRecipeStepsStatus(object):
    def __init__(self, code: int = None, message: str = None, details: list = None):
        self.code = code
        self.message = message
        self.details = details

    @classmethod
    def to_proto(self, resource):
        if not resource:
            return None

        res = instance_pb2.Tier2AlphaInstanceResetRecipeStepsStatus()
        if Primitive.to_proto(resource.code):
            res.code = Primitive.to_proto(resource.code)
        if Primitive.to_proto(resource.message):
            res.message = Primitive.to_proto(resource.message)
        if InstanceResetRecipeStepsStatusDetailsArray.to_proto(resource.details):
            res.details.extend(
                InstanceResetRecipeStepsStatusDetailsArray.to_proto(resource.details)
            )
        return res

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return None

        return InstanceResetRecipeStepsStatus(
            code=Primitive.from_proto(resource.code),
            message=Primitive.from_proto(resource.message),
            details=InstanceResetRecipeStepsStatusDetailsArray.from_proto(
                resource.details
            ),
        )


class InstanceResetRecipeStepsStatusArray(object):
    @classmethod
    def to_proto(self, resources):
        if not resources:
            return resources
        return [InstanceResetRecipeStepsStatus.to_proto(i) for i in resources]

    @classmethod
    def from_proto(self, resources):
        return [InstanceResetRecipeStepsStatus.from_proto(i) for i in resources]


class InstanceResetRecipeStepsStatusDetails(object):
    def __init__(self, type_url: str = None, value: str = None):
        self.type_url = type_url
        self.value = value

    @classmethod
    def to_proto(self, resource):
        if not resource:
            return None

        res = instance_pb2.Tier2AlphaInstanceResetRecipeStepsStatusDetails()
        if Primitive.to_proto(resource.type_url):
            res.type_url = Primitive.to_proto(resource.type_url)
        if Primitive.to_proto(resource.value):
            res.value = Primitive.to_proto(resource.value)
        return res

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return None

        return InstanceResetRecipeStepsStatusDetails(
            type_url=Primitive.from_proto(resource.type_url),
            value=Primitive.from_proto(resource.value),
        )


class InstanceResetRecipeStepsStatusDetailsArray(object):
    @classmethod
    def to_proto(self, resources):
        if not resources:
            return resources
        return [InstanceResetRecipeStepsStatusDetails.to_proto(i) for i in resources]

    @classmethod
    def from_proto(self, resources):
        return [InstanceResetRecipeStepsStatusDetails.from_proto(i) for i in resources]


class InstanceResetRecipeStepsQuotaRequestDeltas(object):
    def __init__(
        self,
        metric_name: str = None,
        amount: int = None,
        quota_location_name: str = None,
    ):
        self.metric_name = metric_name
        self.amount = amount
        self.quota_location_name = quota_location_name

    @classmethod
    def to_proto(self, resource):
        if not resource:
            return None

        res = instance_pb2.Tier2AlphaInstanceResetRecipeStepsQuotaRequestDeltas()
        if Primitive.to_proto(resource.metric_name):
            res.metric_name = Primitive.to_proto(resource.metric_name)
        if Primitive.to_proto(resource.amount):
            res.amount = Primitive.to_proto(resource.amount)
        if Primitive.to_proto(resource.quota_location_name):
            res.quota_location_name = Primitive.to_proto(resource.quota_location_name)
        return res

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return None

        return InstanceResetRecipeStepsQuotaRequestDeltas(
            metric_name=Primitive.from_proto(resource.metric_name),
            amount=Primitive.from_proto(resource.amount),
            quota_location_name=Primitive.from_proto(resource.quota_location_name),
        )


class InstanceResetRecipeStepsQuotaRequestDeltasArray(object):
    @classmethod
    def to_proto(self, resources):
        if not resources:
            return resources
        return [
            InstanceResetRecipeStepsQuotaRequestDeltas.to_proto(i) for i in resources
        ]

    @classmethod
    def from_proto(self, resources):
        return [
            InstanceResetRecipeStepsQuotaRequestDeltas.from_proto(i) for i in resources
        ]


class InstanceResetRecipeStepsPreprocessUpdate(object):
    def __init__(
        self, latency_slo_bucket_name: str = None, public_operation_metadata: str = None
    ):
        self.latency_slo_bucket_name = latency_slo_bucket_name
        self.public_operation_metadata = public_operation_metadata

    @classmethod
    def to_proto(self, resource):
        if not resource:
            return None

        res = instance_pb2.Tier2AlphaInstanceResetRecipeStepsPreprocessUpdate()
        if Primitive.to_proto(resource.latency_slo_bucket_name):
            res.latency_slo_bucket_name = Primitive.to_proto(
                resource.latency_slo_bucket_name
            )
        if Primitive.to_proto(resource.public_operation_metadata):
            res.public_operation_metadata = Primitive.to_proto(
                resource.public_operation_metadata
            )
        return res

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return None

        return InstanceResetRecipeStepsPreprocessUpdate(
            latency_slo_bucket_name=Primitive.from_proto(
                resource.latency_slo_bucket_name
            ),
            public_operation_metadata=Primitive.from_proto(
                resource.public_operation_metadata
            ),
        )


class InstanceResetRecipeStepsPreprocessUpdateArray(object):
    @classmethod
    def to_proto(self, resources):
        if not resources:
            return resources
        return [InstanceResetRecipeStepsPreprocessUpdate.to_proto(i) for i in resources]

    @classmethod
    def from_proto(self, resources):
        return [
            InstanceResetRecipeStepsPreprocessUpdate.from_proto(i) for i in resources
        ]


class InstanceResetRecipeStepsRequestedTenantProject(object):
    def __init__(self, tag: str = None, folder: str = None, scope: str = None):
        self.tag = tag
        self.folder = folder
        self.scope = scope

    @classmethod
    def to_proto(self, resource):
        if not resource:
            return None

        res = instance_pb2.Tier2AlphaInstanceResetRecipeStepsRequestedTenantProject()
        if Primitive.to_proto(resource.tag):
            res.tag = Primitive.to_proto(resource.tag)
        if Primitive.to_proto(resource.folder):
            res.folder = Primitive.to_proto(resource.folder)
        if InstanceResetRecipeStepsRequestedTenantProjectScopeEnum.to_proto(
            resource.scope
        ):
            res.scope = InstanceResetRecipeStepsRequestedTenantProjectScopeEnum.to_proto(
                resource.scope
            )
        return res

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return None

        return InstanceResetRecipeStepsRequestedTenantProject(
            tag=Primitive.from_proto(resource.tag),
            folder=Primitive.from_proto(resource.folder),
            scope=InstanceResetRecipeStepsRequestedTenantProjectScopeEnum.from_proto(
                resource.scope
            ),
        )


class InstanceResetRecipeStepsRequestedTenantProjectArray(object):
    @classmethod
    def to_proto(self, resources):
        if not resources:
            return resources
        return [
            InstanceResetRecipeStepsRequestedTenantProject.to_proto(i)
            for i in resources
        ]

    @classmethod
    def from_proto(self, resources):
        return [
            InstanceResetRecipeStepsRequestedTenantProject.from_proto(i)
            for i in resources
        ]


class InstanceResetRecipeStepsPermissionsInfo(object):
    def __init__(
        self,
        policy_name: dict = None,
        iam_permissions: list = None,
        resource_path: str = None,
        api_attrs: dict = None,
        policy_name_mode: str = None,
        resource: dict = None,
    ):
        self.policy_name = policy_name
        self.iam_permissions = iam_permissions
        self.resource_path = resource_path
        self.api_attrs = api_attrs
        self.policy_name_mode = policy_name_mode
        self.resource = resource

    @classmethod
    def to_proto(self, resource):
        if not resource:
            return None

        res = instance_pb2.Tier2AlphaInstanceResetRecipeStepsPermissionsInfo()
        if InstanceResetRecipeStepsPermissionsInfoPolicyName.to_proto(
            resource.policy_name
        ):
            res.policy_name.CopyFrom(
                InstanceResetRecipeStepsPermissionsInfoPolicyName.to_proto(
                    resource.policy_name
                )
            )
        else:
            res.ClearField("policy_name")
        if InstanceResetRecipeStepsPermissionsInfoIamPermissionsArray.to_proto(
            resource.iam_permissions
        ):
            res.iam_permissions.extend(
                InstanceResetRecipeStepsPermissionsInfoIamPermissionsArray.to_proto(
                    resource.iam_permissions
                )
            )
        if Primitive.to_proto(resource.resource_path):
            res.resource_path = Primitive.to_proto(resource.resource_path)
        if InstanceGoogleprotobufstruct.to_proto(resource.api_attrs):
            res.api_attrs.CopyFrom(
                InstanceGoogleprotobufstruct.to_proto(resource.api_attrs)
            )
        else:
            res.ClearField("api_attrs")
        if InstanceResetRecipeStepsPermissionsInfoPolicyNameModeEnum.to_proto(
            resource.policy_name_mode
        ):
            res.policy_name_mode = InstanceResetRecipeStepsPermissionsInfoPolicyNameModeEnum.to_proto(
                resource.policy_name_mode
            )
        if InstanceResetRecipeStepsPermissionsInfoResource.to_proto(resource.resource):
            res.resource.CopyFrom(
                InstanceResetRecipeStepsPermissionsInfoResource.to_proto(
                    resource.resource
                )
            )
        else:
            res.ClearField("resource")
        return res

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return None

        return InstanceResetRecipeStepsPermissionsInfo(
            policy_name=InstanceResetRecipeStepsPermissionsInfoPolicyName.from_proto(
                resource.policy_name
            ),
            iam_permissions=InstanceResetRecipeStepsPermissionsInfoIamPermissionsArray.from_proto(
                resource.iam_permissions
            ),
            resource_path=Primitive.from_proto(resource.resource_path),
            api_attrs=InstanceGoogleprotobufstruct.from_proto(resource.api_attrs),
            policy_name_mode=InstanceResetRecipeStepsPermissionsInfoPolicyNameModeEnum.from_proto(
                resource.policy_name_mode
            ),
            resource=InstanceResetRecipeStepsPermissionsInfoResource.from_proto(
                resource.resource
            ),
        )


class InstanceResetRecipeStepsPermissionsInfoArray(object):
    @classmethod
    def to_proto(self, resources):
        if not resources:
            return resources
        return [InstanceResetRecipeStepsPermissionsInfo.to_proto(i) for i in resources]

    @classmethod
    def from_proto(self, resources):
        return [
            InstanceResetRecipeStepsPermissionsInfo.from_proto(i) for i in resources
        ]


class InstanceResetRecipeStepsPermissionsInfoPolicyName(object):
    def __init__(self, type: str = None, id: str = None, region: str = None):
        self.type = type
        self.id = id
        self.region = region

    @classmethod
    def to_proto(self, resource):
        if not resource:
            return None

        res = instance_pb2.Tier2AlphaInstanceResetRecipeStepsPermissionsInfoPolicyName()
        if Primitive.to_proto(resource.type):
            res.type = Primitive.to_proto(resource.type)
        if Primitive.to_proto(resource.id):
            res.id = Primitive.to_proto(resource.id)
        if Primitive.to_proto(resource.region):
            res.region = Primitive.to_proto(resource.region)
        return res

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return None

        return InstanceResetRecipeStepsPermissionsInfoPolicyName(
            type=Primitive.from_proto(resource.type),
            id=Primitive.from_proto(resource.id),
            region=Primitive.from_proto(resource.region),
        )


class InstanceResetRecipeStepsPermissionsInfoPolicyNameArray(object):
    @classmethod
    def to_proto(self, resources):
        if not resources:
            return resources
        return [
            InstanceResetRecipeStepsPermissionsInfoPolicyName.to_proto(i)
            for i in resources
        ]

    @classmethod
    def from_proto(self, resources):
        return [
            InstanceResetRecipeStepsPermissionsInfoPolicyName.from_proto(i)
            for i in resources
        ]


class InstanceResetRecipeStepsPermissionsInfoIamPermissions(object):
    def __init__(self, permission: str = None):
        self.permission = permission

    @classmethod
    def to_proto(self, resource):
        if not resource:
            return None

        res = (
            instance_pb2.Tier2AlphaInstanceResetRecipeStepsPermissionsInfoIamPermissions()
        )
        if Primitive.to_proto(resource.permission):
            res.permission = Primitive.to_proto(resource.permission)
        return res

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return None

        return InstanceResetRecipeStepsPermissionsInfoIamPermissions(
            permission=Primitive.from_proto(resource.permission),
        )


class InstanceResetRecipeStepsPermissionsInfoIamPermissionsArray(object):
    @classmethod
    def to_proto(self, resources):
        if not resources:
            return resources
        return [
            InstanceResetRecipeStepsPermissionsInfoIamPermissions.to_proto(i)
            for i in resources
        ]

    @classmethod
    def from_proto(self, resources):
        return [
            InstanceResetRecipeStepsPermissionsInfoIamPermissions.from_proto(i)
            for i in resources
        ]


class InstanceResetRecipeStepsPermissionsInfoResource(object):
    def __init__(
        self,
        name: str = None,
        type: str = None,
        service: str = None,
        labels: dict = None,
    ):
        self.name = name
        self.type = type
        self.service = service
        self.labels = labels

    @classmethod
    def to_proto(self, resource):
        if not resource:
            return None

        res = instance_pb2.Tier2AlphaInstanceResetRecipeStepsPermissionsInfoResource()
        if Primitive.to_proto(resource.name):
            res.name = Primitive.to_proto(resource.name)
        if Primitive.to_proto(resource.type):
            res.type = Primitive.to_proto(resource.type)
        if Primitive.to_proto(resource.service):
            res.service = Primitive.to_proto(resource.service)
        if Primitive.to_proto(resource.labels):
            res.labels = Primitive.to_proto(resource.labels)
        return res

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return None

        return InstanceResetRecipeStepsPermissionsInfoResource(
            name=Primitive.from_proto(resource.name),
            type=Primitive.from_proto(resource.type),
            service=Primitive.from_proto(resource.service),
            labels=Primitive.from_proto(resource.labels),
        )


class InstanceResetRecipeStepsPermissionsInfoResourceArray(object):
    @classmethod
    def to_proto(self, resources):
        if not resources:
            return resources
        return [
            InstanceResetRecipeStepsPermissionsInfoResource.to_proto(i)
            for i in resources
        ]

    @classmethod
    def from_proto(self, resources):
        return [
            InstanceResetRecipeStepsPermissionsInfoResource.from_proto(i)
            for i in resources
        ]


class InstanceResetRecipeStepsKeyNotificationsUpdate(object):
    def __init__(self, key_notifications_info: dict = None):
        self.key_notifications_info = key_notifications_info

    @classmethod
    def to_proto(self, resource):
        if not resource:
            return None

        res = instance_pb2.Tier2AlphaInstanceResetRecipeStepsKeyNotificationsUpdate()
        if InstanceResetRecipeStepsKeyNotificationsUpdateKeyNotificationsInfo.to_proto(
            resource.key_notifications_info
        ):
            res.key_notifications_info.CopyFrom(
                InstanceResetRecipeStepsKeyNotificationsUpdateKeyNotificationsInfo.to_proto(
                    resource.key_notifications_info
                )
            )
        else:
            res.ClearField("key_notifications_info")
        return res

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return None

        return InstanceResetRecipeStepsKeyNotificationsUpdate(
            key_notifications_info=InstanceResetRecipeStepsKeyNotificationsUpdateKeyNotificationsInfo.from_proto(
                resource.key_notifications_info
            ),
        )


class InstanceResetRecipeStepsKeyNotificationsUpdateArray(object):
    @classmethod
    def to_proto(self, resources):
        if not resources:
            return resources
        return [
            InstanceResetRecipeStepsKeyNotificationsUpdate.to_proto(i)
            for i in resources
        ]

    @classmethod
    def from_proto(self, resources):
        return [
            InstanceResetRecipeStepsKeyNotificationsUpdate.from_proto(i)
            for i in resources
        ]


class InstanceResetRecipeStepsKeyNotificationsUpdateKeyNotificationsInfo(object):
    def __init__(
        self,
        data_version: int = None,
        delegate: str = None,
        key_notification_configs: list = None,
    ):
        self.data_version = data_version
        self.delegate = delegate
        self.key_notification_configs = key_notification_configs

    @classmethod
    def to_proto(self, resource):
        if not resource:
            return None

        res = (
            instance_pb2.Tier2AlphaInstanceResetRecipeStepsKeyNotificationsUpdateKeyNotificationsInfo()
        )
        if Primitive.to_proto(resource.data_version):
            res.data_version = Primitive.to_proto(resource.data_version)
        if Primitive.to_proto(resource.delegate):
            res.delegate = Primitive.to_proto(resource.delegate)
        if InstanceResetRecipeStepsKeyNotificationsUpdateKeyNotificationsInfoKeyNotificationConfigsArray.to_proto(
            resource.key_notification_configs
        ):
            res.key_notification_configs.extend(
                InstanceResetRecipeStepsKeyNotificationsUpdateKeyNotificationsInfoKeyNotificationConfigsArray.to_proto(
                    resource.key_notification_configs
                )
            )
        return res

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return None

        return InstanceResetRecipeStepsKeyNotificationsUpdateKeyNotificationsInfo(
            data_version=Primitive.from_proto(resource.data_version),
            delegate=Primitive.from_proto(resource.delegate),
            key_notification_configs=InstanceResetRecipeStepsKeyNotificationsUpdateKeyNotificationsInfoKeyNotificationConfigsArray.from_proto(
                resource.key_notification_configs
            ),
        )


class InstanceResetRecipeStepsKeyNotificationsUpdateKeyNotificationsInfoArray(object):
    @classmethod
    def to_proto(self, resources):
        if not resources:
            return resources
        return [
            InstanceResetRecipeStepsKeyNotificationsUpdateKeyNotificationsInfo.to_proto(
                i
            )
            for i in resources
        ]

    @classmethod
    def from_proto(self, resources):
        return [
            InstanceResetRecipeStepsKeyNotificationsUpdateKeyNotificationsInfo.from_proto(
                i
            )
            for i in resources
        ]


class InstanceResetRecipeStepsKeyNotificationsUpdateKeyNotificationsInfoKeyNotificationConfigs(
    object
):
    def __init__(
        self,
        key_or_version_name: str = None,
        grant: str = None,
        delegator_gaia_id: int = None,
    ):
        self.key_or_version_name = key_or_version_name
        self.grant = grant
        self.delegator_gaia_id = delegator_gaia_id

    @classmethod
    def to_proto(self, resource):
        if not resource:
            return None

        res = (
            instance_pb2.Tier2AlphaInstanceResetRecipeStepsKeyNotificationsUpdateKeyNotificationsInfoKeyNotificationConfigs()
        )
        if Primitive.to_proto(resource.key_or_version_name):
            res.key_or_version_name = Primitive.to_proto(resource.key_or_version_name)
        if Primitive.to_proto(resource.grant):
            res.grant = Primitive.to_proto(resource.grant)
        if Primitive.to_proto(resource.delegator_gaia_id):
            res.delegator_gaia_id = Primitive.to_proto(resource.delegator_gaia_id)
        return res

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return None

        return InstanceResetRecipeStepsKeyNotificationsUpdateKeyNotificationsInfoKeyNotificationConfigs(
            key_or_version_name=Primitive.from_proto(resource.key_or_version_name),
            grant=Primitive.from_proto(resource.grant),
            delegator_gaia_id=Primitive.from_proto(resource.delegator_gaia_id),
        )


class InstanceResetRecipeStepsKeyNotificationsUpdateKeyNotificationsInfoKeyNotificationConfigsArray(
    object
):
    @classmethod
    def to_proto(self, resources):
        if not resources:
            return resources
        return [
            InstanceResetRecipeStepsKeyNotificationsUpdateKeyNotificationsInfoKeyNotificationConfigs.to_proto(
                i
            )
            for i in resources
        ]

    @classmethod
    def from_proto(self, resources):
        return [
            InstanceResetRecipeStepsKeyNotificationsUpdateKeyNotificationsInfoKeyNotificationConfigs.from_proto(
                i
            )
            for i in resources
        ]


class InstancePreprocessRepairRecipe(object):
    def __init__(
        self,
        steps: list = None,
        honor_cancel_request: bool = None,
        ignore_recipe_after: int = None,
        verify_deadline_seconds_below: float = None,
        populate_operation_result: bool = None,
        readonly_recipe_start_time: str = None,
        resource_names_stored_in_clh_with_delay: list = None,
        delay_to_store_resources_in_clh_db_nanos: int = None,
    ):
        self.steps = steps
        self.honor_cancel_request = honor_cancel_request
        self.ignore_recipe_after = ignore_recipe_after
        self.verify_deadline_seconds_below = verify_deadline_seconds_below
        self.populate_operation_result = populate_operation_result
        self.readonly_recipe_start_time = readonly_recipe_start_time
        self.resource_names_stored_in_clh_with_delay = (
            resource_names_stored_in_clh_with_delay
        )
        self.delay_to_store_resources_in_clh_db_nanos = (
            delay_to_store_resources_in_clh_db_nanos
        )

    @classmethod
    def to_proto(self, resource):
        if not resource:
            return None

        res = instance_pb2.Tier2AlphaInstancePreprocessRepairRecipe()
        if InstancePreprocessRepairRecipeStepsArray.to_proto(resource.steps):
            res.steps.extend(
                InstancePreprocessRepairRecipeStepsArray.to_proto(resource.steps)
            )
        if Primitive.to_proto(resource.honor_cancel_request):
            res.honor_cancel_request = Primitive.to_proto(resource.honor_cancel_request)
        if Primitive.to_proto(resource.ignore_recipe_after):
            res.ignore_recipe_after = Primitive.to_proto(resource.ignore_recipe_after)
        if Primitive.to_proto(resource.verify_deadline_seconds_below):
            res.verify_deadline_seconds_below = Primitive.to_proto(
                resource.verify_deadline_seconds_below
            )
        if Primitive.to_proto(resource.populate_operation_result):
            res.populate_operation_result = Primitive.to_proto(
                resource.populate_operation_result
            )
        if Primitive.to_proto(resource.readonly_recipe_start_time):
            res.readonly_recipe_start_time = Primitive.to_proto(
                resource.readonly_recipe_start_time
            )
        if Primitive.to_proto(resource.resource_names_stored_in_clh_with_delay):
            res.resource_names_stored_in_clh_with_delay.extend(
                Primitive.to_proto(resource.resource_names_stored_in_clh_with_delay)
            )
        if Primitive.to_proto(resource.delay_to_store_resources_in_clh_db_nanos):
            res.delay_to_store_resources_in_clh_db_nanos = Primitive.to_proto(
                resource.delay_to_store_resources_in_clh_db_nanos
            )
        return res

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return None

        return InstancePreprocessRepairRecipe(
            steps=InstancePreprocessRepairRecipeStepsArray.from_proto(resource.steps),
            honor_cancel_request=Primitive.from_proto(resource.honor_cancel_request),
            ignore_recipe_after=Primitive.from_proto(resource.ignore_recipe_after),
            verify_deadline_seconds_below=Primitive.from_proto(
                resource.verify_deadline_seconds_below
            ),
            populate_operation_result=Primitive.from_proto(
                resource.populate_operation_result
            ),
            readonly_recipe_start_time=Primitive.from_proto(
                resource.readonly_recipe_start_time
            ),
            resource_names_stored_in_clh_with_delay=Primitive.from_proto(
                resource.resource_names_stored_in_clh_with_delay
            ),
            delay_to_store_resources_in_clh_db_nanos=Primitive.from_proto(
                resource.delay_to_store_resources_in_clh_db_nanos
            ),
        )


class InstancePreprocessRepairRecipeArray(object):
    @classmethod
    def to_proto(self, resources):
        if not resources:
            return resources
        return [InstancePreprocessRepairRecipe.to_proto(i) for i in resources]

    @classmethod
    def from_proto(self, resources):
        return [InstancePreprocessRepairRecipe.from_proto(i) for i in resources]


class InstancePreprocessRepairRecipeSteps(object):
    def __init__(
        self,
        relative_time: int = None,
        sleep_duration: int = None,
        action: str = None,
        status: dict = None,
        error_space: str = None,
        p4_service_account: str = None,
        resource_metadata_size: int = None,
        description: str = None,
        updated_repeat_operation_delay_sec: float = None,
        quota_request_deltas: list = None,
        preprocess_update: dict = None,
        public_operation_metadata: str = None,
        requested_tenant_project: dict = None,
        permissions_info: list = None,
        key_notifications_update: dict = None,
        clh_data_update_time: str = None,
        public_error_message: str = None,
    ):
        self.relative_time = relative_time
        self.sleep_duration = sleep_duration
        self.action = action
        self.status = status
        self.error_space = error_space
        self.p4_service_account = p4_service_account
        self.resource_metadata_size = resource_metadata_size
        self.description = description
        self.updated_repeat_operation_delay_sec = updated_repeat_operation_delay_sec
        self.quota_request_deltas = quota_request_deltas
        self.preprocess_update = preprocess_update
        self.public_operation_metadata = public_operation_metadata
        self.requested_tenant_project = requested_tenant_project
        self.permissions_info = permissions_info
        self.key_notifications_update = key_notifications_update
        self.clh_data_update_time = clh_data_update_time
        self.public_error_message = public_error_message

    @classmethod
    def to_proto(self, resource):
        if not resource:
            return None

        res = instance_pb2.Tier2AlphaInstancePreprocessRepairRecipeSteps()
        if Primitive.to_proto(resource.relative_time):
            res.relative_time = Primitive.to_proto(resource.relative_time)
        if Primitive.to_proto(resource.sleep_duration):
            res.sleep_duration = Primitive.to_proto(resource.sleep_duration)
        if InstancePreprocessRepairRecipeStepsActionEnum.to_proto(resource.action):
            res.action = InstancePreprocessRepairRecipeStepsActionEnum.to_proto(
                resource.action
            )
        if InstancePreprocessRepairRecipeStepsStatus.to_proto(resource.status):
            res.status.CopyFrom(
                InstancePreprocessRepairRecipeStepsStatus.to_proto(resource.status)
            )
        else:
            res.ClearField("status")
        if Primitive.to_proto(resource.error_space):
            res.error_space = Primitive.to_proto(resource.error_space)
        if Primitive.to_proto(resource.p4_service_account):
            res.p4_service_account = Primitive.to_proto(resource.p4_service_account)
        if Primitive.to_proto(resource.resource_metadata_size):
            res.resource_metadata_size = Primitive.to_proto(
                resource.resource_metadata_size
            )
        if Primitive.to_proto(resource.description):
            res.description = Primitive.to_proto(resource.description)
        if Primitive.to_proto(resource.updated_repeat_operation_delay_sec):
            res.updated_repeat_operation_delay_sec = Primitive.to_proto(
                resource.updated_repeat_operation_delay_sec
            )
        if InstancePreprocessRepairRecipeStepsQuotaRequestDeltasArray.to_proto(
            resource.quota_request_deltas
        ):
            res.quota_request_deltas.extend(
                InstancePreprocessRepairRecipeStepsQuotaRequestDeltasArray.to_proto(
                    resource.quota_request_deltas
                )
            )
        if InstancePreprocessRepairRecipeStepsPreprocessUpdate.to_proto(
            resource.preprocess_update
        ):
            res.preprocess_update.CopyFrom(
                InstancePreprocessRepairRecipeStepsPreprocessUpdate.to_proto(
                    resource.preprocess_update
                )
            )
        else:
            res.ClearField("preprocess_update")
        if Primitive.to_proto(resource.public_operation_metadata):
            res.public_operation_metadata = Primitive.to_proto(
                resource.public_operation_metadata
            )
        if InstancePreprocessRepairRecipeStepsRequestedTenantProject.to_proto(
            resource.requested_tenant_project
        ):
            res.requested_tenant_project.CopyFrom(
                InstancePreprocessRepairRecipeStepsRequestedTenantProject.to_proto(
                    resource.requested_tenant_project
                )
            )
        else:
            res.ClearField("requested_tenant_project")
        if InstancePreprocessRepairRecipeStepsPermissionsInfoArray.to_proto(
            resource.permissions_info
        ):
            res.permissions_info.extend(
                InstancePreprocessRepairRecipeStepsPermissionsInfoArray.to_proto(
                    resource.permissions_info
                )
            )
        if InstancePreprocessRepairRecipeStepsKeyNotificationsUpdate.to_proto(
            resource.key_notifications_update
        ):
            res.key_notifications_update.CopyFrom(
                InstancePreprocessRepairRecipeStepsKeyNotificationsUpdate.to_proto(
                    resource.key_notifications_update
                )
            )
        else:
            res.ClearField("key_notifications_update")
        if Primitive.to_proto(resource.clh_data_update_time):
            res.clh_data_update_time = Primitive.to_proto(resource.clh_data_update_time)
        if Primitive.to_proto(resource.public_error_message):
            res.public_error_message = Primitive.to_proto(resource.public_error_message)
        return res

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return None

        return InstancePreprocessRepairRecipeSteps(
            relative_time=Primitive.from_proto(resource.relative_time),
            sleep_duration=Primitive.from_proto(resource.sleep_duration),
            action=InstancePreprocessRepairRecipeStepsActionEnum.from_proto(
                resource.action
            ),
            status=InstancePreprocessRepairRecipeStepsStatus.from_proto(
                resource.status
            ),
            error_space=Primitive.from_proto(resource.error_space),
            p4_service_account=Primitive.from_proto(resource.p4_service_account),
            resource_metadata_size=Primitive.from_proto(
                resource.resource_metadata_size
            ),
            description=Primitive.from_proto(resource.description),
            updated_repeat_operation_delay_sec=Primitive.from_proto(
                resource.updated_repeat_operation_delay_sec
            ),
            quota_request_deltas=InstancePreprocessRepairRecipeStepsQuotaRequestDeltasArray.from_proto(
                resource.quota_request_deltas
            ),
            preprocess_update=InstancePreprocessRepairRecipeStepsPreprocessUpdate.from_proto(
                resource.preprocess_update
            ),
            public_operation_metadata=Primitive.from_proto(
                resource.public_operation_metadata
            ),
            requested_tenant_project=InstancePreprocessRepairRecipeStepsRequestedTenantProject.from_proto(
                resource.requested_tenant_project
            ),
            permissions_info=InstancePreprocessRepairRecipeStepsPermissionsInfoArray.from_proto(
                resource.permissions_info
            ),
            key_notifications_update=InstancePreprocessRepairRecipeStepsKeyNotificationsUpdate.from_proto(
                resource.key_notifications_update
            ),
            clh_data_update_time=Primitive.from_proto(resource.clh_data_update_time),
            public_error_message=Primitive.from_proto(resource.public_error_message),
        )


class InstancePreprocessRepairRecipeStepsArray(object):
    @classmethod
    def to_proto(self, resources):
        if not resources:
            return resources
        return [InstancePreprocessRepairRecipeSteps.to_proto(i) for i in resources]

    @classmethod
    def from_proto(self, resources):
        return [InstancePreprocessRepairRecipeSteps.from_proto(i) for i in resources]


class InstancePreprocessRepairRecipeStepsStatus(object):
    def __init__(self, code: int = None, message: str = None, details: list = None):
        self.code = code
        self.message = message
        self.details = details

    @classmethod
    def to_proto(self, resource):
        if not resource:
            return None

        res = instance_pb2.Tier2AlphaInstancePreprocessRepairRecipeStepsStatus()
        if Primitive.to_proto(resource.code):
            res.code = Primitive.to_proto(resource.code)
        if Primitive.to_proto(resource.message):
            res.message = Primitive.to_proto(resource.message)
        if InstancePreprocessRepairRecipeStepsStatusDetailsArray.to_proto(
            resource.details
        ):
            res.details.extend(
                InstancePreprocessRepairRecipeStepsStatusDetailsArray.to_proto(
                    resource.details
                )
            )
        return res

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return None

        return InstancePreprocessRepairRecipeStepsStatus(
            code=Primitive.from_proto(resource.code),
            message=Primitive.from_proto(resource.message),
            details=InstancePreprocessRepairRecipeStepsStatusDetailsArray.from_proto(
                resource.details
            ),
        )


class InstancePreprocessRepairRecipeStepsStatusArray(object):
    @classmethod
    def to_proto(self, resources):
        if not resources:
            return resources
        return [
            InstancePreprocessRepairRecipeStepsStatus.to_proto(i) for i in resources
        ]

    @classmethod
    def from_proto(self, resources):
        return [
            InstancePreprocessRepairRecipeStepsStatus.from_proto(i) for i in resources
        ]


class InstancePreprocessRepairRecipeStepsStatusDetails(object):
    def __init__(self, type_url: str = None, value: str = None):
        self.type_url = type_url
        self.value = value

    @classmethod
    def to_proto(self, resource):
        if not resource:
            return None

        res = instance_pb2.Tier2AlphaInstancePreprocessRepairRecipeStepsStatusDetails()
        if Primitive.to_proto(resource.type_url):
            res.type_url = Primitive.to_proto(resource.type_url)
        if Primitive.to_proto(resource.value):
            res.value = Primitive.to_proto(resource.value)
        return res

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return None

        return InstancePreprocessRepairRecipeStepsStatusDetails(
            type_url=Primitive.from_proto(resource.type_url),
            value=Primitive.from_proto(resource.value),
        )


class InstancePreprocessRepairRecipeStepsStatusDetailsArray(object):
    @classmethod
    def to_proto(self, resources):
        if not resources:
            return resources
        return [
            InstancePreprocessRepairRecipeStepsStatusDetails.to_proto(i)
            for i in resources
        ]

    @classmethod
    def from_proto(self, resources):
        return [
            InstancePreprocessRepairRecipeStepsStatusDetails.from_proto(i)
            for i in resources
        ]


class InstancePreprocessRepairRecipeStepsQuotaRequestDeltas(object):
    def __init__(
        self,
        metric_name: str = None,
        amount: int = None,
        quota_location_name: str = None,
    ):
        self.metric_name = metric_name
        self.amount = amount
        self.quota_location_name = quota_location_name

    @classmethod
    def to_proto(self, resource):
        if not resource:
            return None

        res = (
            instance_pb2.Tier2AlphaInstancePreprocessRepairRecipeStepsQuotaRequestDeltas()
        )
        if Primitive.to_proto(resource.metric_name):
            res.metric_name = Primitive.to_proto(resource.metric_name)
        if Primitive.to_proto(resource.amount):
            res.amount = Primitive.to_proto(resource.amount)
        if Primitive.to_proto(resource.quota_location_name):
            res.quota_location_name = Primitive.to_proto(resource.quota_location_name)
        return res

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return None

        return InstancePreprocessRepairRecipeStepsQuotaRequestDeltas(
            metric_name=Primitive.from_proto(resource.metric_name),
            amount=Primitive.from_proto(resource.amount),
            quota_location_name=Primitive.from_proto(resource.quota_location_name),
        )


class InstancePreprocessRepairRecipeStepsQuotaRequestDeltasArray(object):
    @classmethod
    def to_proto(self, resources):
        if not resources:
            return resources
        return [
            InstancePreprocessRepairRecipeStepsQuotaRequestDeltas.to_proto(i)
            for i in resources
        ]

    @classmethod
    def from_proto(self, resources):
        return [
            InstancePreprocessRepairRecipeStepsQuotaRequestDeltas.from_proto(i)
            for i in resources
        ]


class InstancePreprocessRepairRecipeStepsPreprocessUpdate(object):
    def __init__(
        self, latency_slo_bucket_name: str = None, public_operation_metadata: str = None
    ):
        self.latency_slo_bucket_name = latency_slo_bucket_name
        self.public_operation_metadata = public_operation_metadata

    @classmethod
    def to_proto(self, resource):
        if not resource:
            return None

        res = (
            instance_pb2.Tier2AlphaInstancePreprocessRepairRecipeStepsPreprocessUpdate()
        )
        if Primitive.to_proto(resource.latency_slo_bucket_name):
            res.latency_slo_bucket_name = Primitive.to_proto(
                resource.latency_slo_bucket_name
            )
        if Primitive.to_proto(resource.public_operation_metadata):
            res.public_operation_metadata = Primitive.to_proto(
                resource.public_operation_metadata
            )
        return res

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return None

        return InstancePreprocessRepairRecipeStepsPreprocessUpdate(
            latency_slo_bucket_name=Primitive.from_proto(
                resource.latency_slo_bucket_name
            ),
            public_operation_metadata=Primitive.from_proto(
                resource.public_operation_metadata
            ),
        )


class InstancePreprocessRepairRecipeStepsPreprocessUpdateArray(object):
    @classmethod
    def to_proto(self, resources):
        if not resources:
            return resources
        return [
            InstancePreprocessRepairRecipeStepsPreprocessUpdate.to_proto(i)
            for i in resources
        ]

    @classmethod
    def from_proto(self, resources):
        return [
            InstancePreprocessRepairRecipeStepsPreprocessUpdate.from_proto(i)
            for i in resources
        ]


class InstancePreprocessRepairRecipeStepsRequestedTenantProject(object):
    def __init__(self, tag: str = None, folder: str = None, scope: str = None):
        self.tag = tag
        self.folder = folder
        self.scope = scope

    @classmethod
    def to_proto(self, resource):
        if not resource:
            return None

        res = (
            instance_pb2.Tier2AlphaInstancePreprocessRepairRecipeStepsRequestedTenantProject()
        )
        if Primitive.to_proto(resource.tag):
            res.tag = Primitive.to_proto(resource.tag)
        if Primitive.to_proto(resource.folder):
            res.folder = Primitive.to_proto(resource.folder)
        if InstancePreprocessRepairRecipeStepsRequestedTenantProjectScopeEnum.to_proto(
            resource.scope
        ):
            res.scope = InstancePreprocessRepairRecipeStepsRequestedTenantProjectScopeEnum.to_proto(
                resource.scope
            )
        return res

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return None

        return InstancePreprocessRepairRecipeStepsRequestedTenantProject(
            tag=Primitive.from_proto(resource.tag),
            folder=Primitive.from_proto(resource.folder),
            scope=InstancePreprocessRepairRecipeStepsRequestedTenantProjectScopeEnum.from_proto(
                resource.scope
            ),
        )


class InstancePreprocessRepairRecipeStepsRequestedTenantProjectArray(object):
    @classmethod
    def to_proto(self, resources):
        if not resources:
            return resources
        return [
            InstancePreprocessRepairRecipeStepsRequestedTenantProject.to_proto(i)
            for i in resources
        ]

    @classmethod
    def from_proto(self, resources):
        return [
            InstancePreprocessRepairRecipeStepsRequestedTenantProject.from_proto(i)
            for i in resources
        ]


class InstancePreprocessRepairRecipeStepsPermissionsInfo(object):
    def __init__(
        self,
        policy_name: dict = None,
        iam_permissions: list = None,
        resource_path: str = None,
        api_attrs: dict = None,
        policy_name_mode: str = None,
        resource: dict = None,
    ):
        self.policy_name = policy_name
        self.iam_permissions = iam_permissions
        self.resource_path = resource_path
        self.api_attrs = api_attrs
        self.policy_name_mode = policy_name_mode
        self.resource = resource

    @classmethod
    def to_proto(self, resource):
        if not resource:
            return None

        res = (
            instance_pb2.Tier2AlphaInstancePreprocessRepairRecipeStepsPermissionsInfo()
        )
        if InstancePreprocessRepairRecipeStepsPermissionsInfoPolicyName.to_proto(
            resource.policy_name
        ):
            res.policy_name.CopyFrom(
                InstancePreprocessRepairRecipeStepsPermissionsInfoPolicyName.to_proto(
                    resource.policy_name
                )
            )
        else:
            res.ClearField("policy_name")
        if InstancePreprocessRepairRecipeStepsPermissionsInfoIamPermissionsArray.to_proto(
            resource.iam_permissions
        ):
            res.iam_permissions.extend(
                InstancePreprocessRepairRecipeStepsPermissionsInfoIamPermissionsArray.to_proto(
                    resource.iam_permissions
                )
            )
        if Primitive.to_proto(resource.resource_path):
            res.resource_path = Primitive.to_proto(resource.resource_path)
        if InstanceGoogleprotobufstruct.to_proto(resource.api_attrs):
            res.api_attrs.CopyFrom(
                InstanceGoogleprotobufstruct.to_proto(resource.api_attrs)
            )
        else:
            res.ClearField("api_attrs")
        if InstancePreprocessRepairRecipeStepsPermissionsInfoPolicyNameModeEnum.to_proto(
            resource.policy_name_mode
        ):
            res.policy_name_mode = InstancePreprocessRepairRecipeStepsPermissionsInfoPolicyNameModeEnum.to_proto(
                resource.policy_name_mode
            )
        if InstancePreprocessRepairRecipeStepsPermissionsInfoResource.to_proto(
            resource.resource
        ):
            res.resource.CopyFrom(
                InstancePreprocessRepairRecipeStepsPermissionsInfoResource.to_proto(
                    resource.resource
                )
            )
        else:
            res.ClearField("resource")
        return res

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return None

        return InstancePreprocessRepairRecipeStepsPermissionsInfo(
            policy_name=InstancePreprocessRepairRecipeStepsPermissionsInfoPolicyName.from_proto(
                resource.policy_name
            ),
            iam_permissions=InstancePreprocessRepairRecipeStepsPermissionsInfoIamPermissionsArray.from_proto(
                resource.iam_permissions
            ),
            resource_path=Primitive.from_proto(resource.resource_path),
            api_attrs=InstanceGoogleprotobufstruct.from_proto(resource.api_attrs),
            policy_name_mode=InstancePreprocessRepairRecipeStepsPermissionsInfoPolicyNameModeEnum.from_proto(
                resource.policy_name_mode
            ),
            resource=InstancePreprocessRepairRecipeStepsPermissionsInfoResource.from_proto(
                resource.resource
            ),
        )


class InstancePreprocessRepairRecipeStepsPermissionsInfoArray(object):
    @classmethod
    def to_proto(self, resources):
        if not resources:
            return resources
        return [
            InstancePreprocessRepairRecipeStepsPermissionsInfo.to_proto(i)
            for i in resources
        ]

    @classmethod
    def from_proto(self, resources):
        return [
            InstancePreprocessRepairRecipeStepsPermissionsInfo.from_proto(i)
            for i in resources
        ]


class InstancePreprocessRepairRecipeStepsPermissionsInfoPolicyName(object):
    def __init__(self, type: str = None, id: str = None, region: str = None):
        self.type = type
        self.id = id
        self.region = region

    @classmethod
    def to_proto(self, resource):
        if not resource:
            return None

        res = (
            instance_pb2.Tier2AlphaInstancePreprocessRepairRecipeStepsPermissionsInfoPolicyName()
        )
        if Primitive.to_proto(resource.type):
            res.type = Primitive.to_proto(resource.type)
        if Primitive.to_proto(resource.id):
            res.id = Primitive.to_proto(resource.id)
        if Primitive.to_proto(resource.region):
            res.region = Primitive.to_proto(resource.region)
        return res

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return None

        return InstancePreprocessRepairRecipeStepsPermissionsInfoPolicyName(
            type=Primitive.from_proto(resource.type),
            id=Primitive.from_proto(resource.id),
            region=Primitive.from_proto(resource.region),
        )


class InstancePreprocessRepairRecipeStepsPermissionsInfoPolicyNameArray(object):
    @classmethod
    def to_proto(self, resources):
        if not resources:
            return resources
        return [
            InstancePreprocessRepairRecipeStepsPermissionsInfoPolicyName.to_proto(i)
            for i in resources
        ]

    @classmethod
    def from_proto(self, resources):
        return [
            InstancePreprocessRepairRecipeStepsPermissionsInfoPolicyName.from_proto(i)
            for i in resources
        ]


class InstancePreprocessRepairRecipeStepsPermissionsInfoIamPermissions(object):
    def __init__(self, permission: str = None):
        self.permission = permission

    @classmethod
    def to_proto(self, resource):
        if not resource:
            return None

        res = (
            instance_pb2.Tier2AlphaInstancePreprocessRepairRecipeStepsPermissionsInfoIamPermissions()
        )
        if Primitive.to_proto(resource.permission):
            res.permission = Primitive.to_proto(resource.permission)
        return res

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return None

        return InstancePreprocessRepairRecipeStepsPermissionsInfoIamPermissions(
            permission=Primitive.from_proto(resource.permission),
        )


class InstancePreprocessRepairRecipeStepsPermissionsInfoIamPermissionsArray(object):
    @classmethod
    def to_proto(self, resources):
        if not resources:
            return resources
        return [
            InstancePreprocessRepairRecipeStepsPermissionsInfoIamPermissions.to_proto(i)
            for i in resources
        ]

    @classmethod
    def from_proto(self, resources):
        return [
            InstancePreprocessRepairRecipeStepsPermissionsInfoIamPermissions.from_proto(
                i
            )
            for i in resources
        ]


class InstancePreprocessRepairRecipeStepsPermissionsInfoResource(object):
    def __init__(
        self,
        name: str = None,
        type: str = None,
        service: str = None,
        labels: dict = None,
    ):
        self.name = name
        self.type = type
        self.service = service
        self.labels = labels

    @classmethod
    def to_proto(self, resource):
        if not resource:
            return None

        res = (
            instance_pb2.Tier2AlphaInstancePreprocessRepairRecipeStepsPermissionsInfoResource()
        )
        if Primitive.to_proto(resource.name):
            res.name = Primitive.to_proto(resource.name)
        if Primitive.to_proto(resource.type):
            res.type = Primitive.to_proto(resource.type)
        if Primitive.to_proto(resource.service):
            res.service = Primitive.to_proto(resource.service)
        if Primitive.to_proto(resource.labels):
            res.labels = Primitive.to_proto(resource.labels)
        return res

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return None

        return InstancePreprocessRepairRecipeStepsPermissionsInfoResource(
            name=Primitive.from_proto(resource.name),
            type=Primitive.from_proto(resource.type),
            service=Primitive.from_proto(resource.service),
            labels=Primitive.from_proto(resource.labels),
        )


class InstancePreprocessRepairRecipeStepsPermissionsInfoResourceArray(object):
    @classmethod
    def to_proto(self, resources):
        if not resources:
            return resources
        return [
            InstancePreprocessRepairRecipeStepsPermissionsInfoResource.to_proto(i)
            for i in resources
        ]

    @classmethod
    def from_proto(self, resources):
        return [
            InstancePreprocessRepairRecipeStepsPermissionsInfoResource.from_proto(i)
            for i in resources
        ]


class InstancePreprocessRepairRecipeStepsKeyNotificationsUpdate(object):
    def __init__(self, key_notifications_info: dict = None):
        self.key_notifications_info = key_notifications_info

    @classmethod
    def to_proto(self, resource):
        if not resource:
            return None

        res = (
            instance_pb2.Tier2AlphaInstancePreprocessRepairRecipeStepsKeyNotificationsUpdate()
        )
        if InstancePreprocessRepairRecipeStepsKeyNotificationsUpdateKeyNotificationsInfo.to_proto(
            resource.key_notifications_info
        ):
            res.key_notifications_info.CopyFrom(
                InstancePreprocessRepairRecipeStepsKeyNotificationsUpdateKeyNotificationsInfo.to_proto(
                    resource.key_notifications_info
                )
            )
        else:
            res.ClearField("key_notifications_info")
        return res

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return None

        return InstancePreprocessRepairRecipeStepsKeyNotificationsUpdate(
            key_notifications_info=InstancePreprocessRepairRecipeStepsKeyNotificationsUpdateKeyNotificationsInfo.from_proto(
                resource.key_notifications_info
            ),
        )


class InstancePreprocessRepairRecipeStepsKeyNotificationsUpdateArray(object):
    @classmethod
    def to_proto(self, resources):
        if not resources:
            return resources
        return [
            InstancePreprocessRepairRecipeStepsKeyNotificationsUpdate.to_proto(i)
            for i in resources
        ]

    @classmethod
    def from_proto(self, resources):
        return [
            InstancePreprocessRepairRecipeStepsKeyNotificationsUpdate.from_proto(i)
            for i in resources
        ]


class InstancePreprocessRepairRecipeStepsKeyNotificationsUpdateKeyNotificationsInfo(
    object
):
    def __init__(
        self,
        data_version: int = None,
        delegate: str = None,
        key_notification_configs: list = None,
    ):
        self.data_version = data_version
        self.delegate = delegate
        self.key_notification_configs = key_notification_configs

    @classmethod
    def to_proto(self, resource):
        if not resource:
            return None

        res = (
            instance_pb2.Tier2AlphaInstancePreprocessRepairRecipeStepsKeyNotificationsUpdateKeyNotificationsInfo()
        )
        if Primitive.to_proto(resource.data_version):
            res.data_version = Primitive.to_proto(resource.data_version)
        if Primitive.to_proto(resource.delegate):
            res.delegate = Primitive.to_proto(resource.delegate)
        if InstancePreprocessRepairRecipeStepsKeyNotificationsUpdateKeyNotificationsInfoKeyNotificationConfigsArray.to_proto(
            resource.key_notification_configs
        ):
            res.key_notification_configs.extend(
                InstancePreprocessRepairRecipeStepsKeyNotificationsUpdateKeyNotificationsInfoKeyNotificationConfigsArray.to_proto(
                    resource.key_notification_configs
                )
            )
        return res

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return None

        return InstancePreprocessRepairRecipeStepsKeyNotificationsUpdateKeyNotificationsInfo(
            data_version=Primitive.from_proto(resource.data_version),
            delegate=Primitive.from_proto(resource.delegate),
            key_notification_configs=InstancePreprocessRepairRecipeStepsKeyNotificationsUpdateKeyNotificationsInfoKeyNotificationConfigsArray.from_proto(
                resource.key_notification_configs
            ),
        )


class InstancePreprocessRepairRecipeStepsKeyNotificationsUpdateKeyNotificationsInfoArray(
    object
):
    @classmethod
    def to_proto(self, resources):
        if not resources:
            return resources
        return [
            InstancePreprocessRepairRecipeStepsKeyNotificationsUpdateKeyNotificationsInfo.to_proto(
                i
            )
            for i in resources
        ]

    @classmethod
    def from_proto(self, resources):
        return [
            InstancePreprocessRepairRecipeStepsKeyNotificationsUpdateKeyNotificationsInfo.from_proto(
                i
            )
            for i in resources
        ]


class InstancePreprocessRepairRecipeStepsKeyNotificationsUpdateKeyNotificationsInfoKeyNotificationConfigs(
    object
):
    def __init__(
        self,
        key_or_version_name: str = None,
        grant: str = None,
        delegator_gaia_id: int = None,
    ):
        self.key_or_version_name = key_or_version_name
        self.grant = grant
        self.delegator_gaia_id = delegator_gaia_id

    @classmethod
    def to_proto(self, resource):
        if not resource:
            return None

        res = (
            instance_pb2.Tier2AlphaInstancePreprocessRepairRecipeStepsKeyNotificationsUpdateKeyNotificationsInfoKeyNotificationConfigs()
        )
        if Primitive.to_proto(resource.key_or_version_name):
            res.key_or_version_name = Primitive.to_proto(resource.key_or_version_name)
        if Primitive.to_proto(resource.grant):
            res.grant = Primitive.to_proto(resource.grant)
        if Primitive.to_proto(resource.delegator_gaia_id):
            res.delegator_gaia_id = Primitive.to_proto(resource.delegator_gaia_id)
        return res

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return None

        return InstancePreprocessRepairRecipeStepsKeyNotificationsUpdateKeyNotificationsInfoKeyNotificationConfigs(
            key_or_version_name=Primitive.from_proto(resource.key_or_version_name),
            grant=Primitive.from_proto(resource.grant),
            delegator_gaia_id=Primitive.from_proto(resource.delegator_gaia_id),
        )


class InstancePreprocessRepairRecipeStepsKeyNotificationsUpdateKeyNotificationsInfoKeyNotificationConfigsArray(
    object
):
    @classmethod
    def to_proto(self, resources):
        if not resources:
            return resources
        return [
            InstancePreprocessRepairRecipeStepsKeyNotificationsUpdateKeyNotificationsInfoKeyNotificationConfigs.to_proto(
                i
            )
            for i in resources
        ]

    @classmethod
    def from_proto(self, resources):
        return [
            InstancePreprocessRepairRecipeStepsKeyNotificationsUpdateKeyNotificationsInfoKeyNotificationConfigs.from_proto(
                i
            )
            for i in resources
        ]


class InstanceRepairRecipe(object):
    def __init__(
        self,
        steps: list = None,
        honor_cancel_request: bool = None,
        ignore_recipe_after: int = None,
        verify_deadline_seconds_below: float = None,
        populate_operation_result: bool = None,
        readonly_recipe_start_time: str = None,
        resource_names_stored_in_clh_with_delay: list = None,
        delay_to_store_resources_in_clh_db_nanos: int = None,
    ):
        self.steps = steps
        self.honor_cancel_request = honor_cancel_request
        self.ignore_recipe_after = ignore_recipe_after
        self.verify_deadline_seconds_below = verify_deadline_seconds_below
        self.populate_operation_result = populate_operation_result
        self.readonly_recipe_start_time = readonly_recipe_start_time
        self.resource_names_stored_in_clh_with_delay = (
            resource_names_stored_in_clh_with_delay
        )
        self.delay_to_store_resources_in_clh_db_nanos = (
            delay_to_store_resources_in_clh_db_nanos
        )

    @classmethod
    def to_proto(self, resource):
        if not resource:
            return None

        res = instance_pb2.Tier2AlphaInstanceRepairRecipe()
        if InstanceRepairRecipeStepsArray.to_proto(resource.steps):
            res.steps.extend(InstanceRepairRecipeStepsArray.to_proto(resource.steps))
        if Primitive.to_proto(resource.honor_cancel_request):
            res.honor_cancel_request = Primitive.to_proto(resource.honor_cancel_request)
        if Primitive.to_proto(resource.ignore_recipe_after):
            res.ignore_recipe_after = Primitive.to_proto(resource.ignore_recipe_after)
        if Primitive.to_proto(resource.verify_deadline_seconds_below):
            res.verify_deadline_seconds_below = Primitive.to_proto(
                resource.verify_deadline_seconds_below
            )
        if Primitive.to_proto(resource.populate_operation_result):
            res.populate_operation_result = Primitive.to_proto(
                resource.populate_operation_result
            )
        if Primitive.to_proto(resource.readonly_recipe_start_time):
            res.readonly_recipe_start_time = Primitive.to_proto(
                resource.readonly_recipe_start_time
            )
        if Primitive.to_proto(resource.resource_names_stored_in_clh_with_delay):
            res.resource_names_stored_in_clh_with_delay.extend(
                Primitive.to_proto(resource.resource_names_stored_in_clh_with_delay)
            )
        if Primitive.to_proto(resource.delay_to_store_resources_in_clh_db_nanos):
            res.delay_to_store_resources_in_clh_db_nanos = Primitive.to_proto(
                resource.delay_to_store_resources_in_clh_db_nanos
            )
        return res

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return None

        return InstanceRepairRecipe(
            steps=InstanceRepairRecipeStepsArray.from_proto(resource.steps),
            honor_cancel_request=Primitive.from_proto(resource.honor_cancel_request),
            ignore_recipe_after=Primitive.from_proto(resource.ignore_recipe_after),
            verify_deadline_seconds_below=Primitive.from_proto(
                resource.verify_deadline_seconds_below
            ),
            populate_operation_result=Primitive.from_proto(
                resource.populate_operation_result
            ),
            readonly_recipe_start_time=Primitive.from_proto(
                resource.readonly_recipe_start_time
            ),
            resource_names_stored_in_clh_with_delay=Primitive.from_proto(
                resource.resource_names_stored_in_clh_with_delay
            ),
            delay_to_store_resources_in_clh_db_nanos=Primitive.from_proto(
                resource.delay_to_store_resources_in_clh_db_nanos
            ),
        )


class InstanceRepairRecipeArray(object):
    @classmethod
    def to_proto(self, resources):
        if not resources:
            return resources
        return [InstanceRepairRecipe.to_proto(i) for i in resources]

    @classmethod
    def from_proto(self, resources):
        return [InstanceRepairRecipe.from_proto(i) for i in resources]


class InstanceRepairRecipeSteps(object):
    def __init__(
        self,
        relative_time: int = None,
        sleep_duration: int = None,
        action: str = None,
        status: dict = None,
        error_space: str = None,
        p4_service_account: str = None,
        resource_metadata_size: int = None,
        description: str = None,
        updated_repeat_operation_delay_sec: float = None,
        quota_request_deltas: list = None,
        preprocess_update: dict = None,
        public_operation_metadata: str = None,
        requested_tenant_project: dict = None,
        permissions_info: list = None,
        key_notifications_update: dict = None,
        clh_data_update_time: str = None,
        public_error_message: str = None,
    ):
        self.relative_time = relative_time
        self.sleep_duration = sleep_duration
        self.action = action
        self.status = status
        self.error_space = error_space
        self.p4_service_account = p4_service_account
        self.resource_metadata_size = resource_metadata_size
        self.description = description
        self.updated_repeat_operation_delay_sec = updated_repeat_operation_delay_sec
        self.quota_request_deltas = quota_request_deltas
        self.preprocess_update = preprocess_update
        self.public_operation_metadata = public_operation_metadata
        self.requested_tenant_project = requested_tenant_project
        self.permissions_info = permissions_info
        self.key_notifications_update = key_notifications_update
        self.clh_data_update_time = clh_data_update_time
        self.public_error_message = public_error_message

    @classmethod
    def to_proto(self, resource):
        if not resource:
            return None

        res = instance_pb2.Tier2AlphaInstanceRepairRecipeSteps()
        if Primitive.to_proto(resource.relative_time):
            res.relative_time = Primitive.to_proto(resource.relative_time)
        if Primitive.to_proto(resource.sleep_duration):
            res.sleep_duration = Primitive.to_proto(resource.sleep_duration)
        if InstanceRepairRecipeStepsActionEnum.to_proto(resource.action):
            res.action = InstanceRepairRecipeStepsActionEnum.to_proto(resource.action)
        if InstanceRepairRecipeStepsStatus.to_proto(resource.status):
            res.status.CopyFrom(
                InstanceRepairRecipeStepsStatus.to_proto(resource.status)
            )
        else:
            res.ClearField("status")
        if Primitive.to_proto(resource.error_space):
            res.error_space = Primitive.to_proto(resource.error_space)
        if Primitive.to_proto(resource.p4_service_account):
            res.p4_service_account = Primitive.to_proto(resource.p4_service_account)
        if Primitive.to_proto(resource.resource_metadata_size):
            res.resource_metadata_size = Primitive.to_proto(
                resource.resource_metadata_size
            )
        if Primitive.to_proto(resource.description):
            res.description = Primitive.to_proto(resource.description)
        if Primitive.to_proto(resource.updated_repeat_operation_delay_sec):
            res.updated_repeat_operation_delay_sec = Primitive.to_proto(
                resource.updated_repeat_operation_delay_sec
            )
        if InstanceRepairRecipeStepsQuotaRequestDeltasArray.to_proto(
            resource.quota_request_deltas
        ):
            res.quota_request_deltas.extend(
                InstanceRepairRecipeStepsQuotaRequestDeltasArray.to_proto(
                    resource.quota_request_deltas
                )
            )
        if InstanceRepairRecipeStepsPreprocessUpdate.to_proto(
            resource.preprocess_update
        ):
            res.preprocess_update.CopyFrom(
                InstanceRepairRecipeStepsPreprocessUpdate.to_proto(
                    resource.preprocess_update
                )
            )
        else:
            res.ClearField("preprocess_update")
        if Primitive.to_proto(resource.public_operation_metadata):
            res.public_operation_metadata = Primitive.to_proto(
                resource.public_operation_metadata
            )
        if InstanceRepairRecipeStepsRequestedTenantProject.to_proto(
            resource.requested_tenant_project
        ):
            res.requested_tenant_project.CopyFrom(
                InstanceRepairRecipeStepsRequestedTenantProject.to_proto(
                    resource.requested_tenant_project
                )
            )
        else:
            res.ClearField("requested_tenant_project")
        if InstanceRepairRecipeStepsPermissionsInfoArray.to_proto(
            resource.permissions_info
        ):
            res.permissions_info.extend(
                InstanceRepairRecipeStepsPermissionsInfoArray.to_proto(
                    resource.permissions_info
                )
            )
        if InstanceRepairRecipeStepsKeyNotificationsUpdate.to_proto(
            resource.key_notifications_update
        ):
            res.key_notifications_update.CopyFrom(
                InstanceRepairRecipeStepsKeyNotificationsUpdate.to_proto(
                    resource.key_notifications_update
                )
            )
        else:
            res.ClearField("key_notifications_update")
        if Primitive.to_proto(resource.clh_data_update_time):
            res.clh_data_update_time = Primitive.to_proto(resource.clh_data_update_time)
        if Primitive.to_proto(resource.public_error_message):
            res.public_error_message = Primitive.to_proto(resource.public_error_message)
        return res

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return None

        return InstanceRepairRecipeSteps(
            relative_time=Primitive.from_proto(resource.relative_time),
            sleep_duration=Primitive.from_proto(resource.sleep_duration),
            action=InstanceRepairRecipeStepsActionEnum.from_proto(resource.action),
            status=InstanceRepairRecipeStepsStatus.from_proto(resource.status),
            error_space=Primitive.from_proto(resource.error_space),
            p4_service_account=Primitive.from_proto(resource.p4_service_account),
            resource_metadata_size=Primitive.from_proto(
                resource.resource_metadata_size
            ),
            description=Primitive.from_proto(resource.description),
            updated_repeat_operation_delay_sec=Primitive.from_proto(
                resource.updated_repeat_operation_delay_sec
            ),
            quota_request_deltas=InstanceRepairRecipeStepsQuotaRequestDeltasArray.from_proto(
                resource.quota_request_deltas
            ),
            preprocess_update=InstanceRepairRecipeStepsPreprocessUpdate.from_proto(
                resource.preprocess_update
            ),
            public_operation_metadata=Primitive.from_proto(
                resource.public_operation_metadata
            ),
            requested_tenant_project=InstanceRepairRecipeStepsRequestedTenantProject.from_proto(
                resource.requested_tenant_project
            ),
            permissions_info=InstanceRepairRecipeStepsPermissionsInfoArray.from_proto(
                resource.permissions_info
            ),
            key_notifications_update=InstanceRepairRecipeStepsKeyNotificationsUpdate.from_proto(
                resource.key_notifications_update
            ),
            clh_data_update_time=Primitive.from_proto(resource.clh_data_update_time),
            public_error_message=Primitive.from_proto(resource.public_error_message),
        )


class InstanceRepairRecipeStepsArray(object):
    @classmethod
    def to_proto(self, resources):
        if not resources:
            return resources
        return [InstanceRepairRecipeSteps.to_proto(i) for i in resources]

    @classmethod
    def from_proto(self, resources):
        return [InstanceRepairRecipeSteps.from_proto(i) for i in resources]


class InstanceRepairRecipeStepsStatus(object):
    def __init__(self, code: int = None, message: str = None, details: list = None):
        self.code = code
        self.message = message
        self.details = details

    @classmethod
    def to_proto(self, resource):
        if not resource:
            return None

        res = instance_pb2.Tier2AlphaInstanceRepairRecipeStepsStatus()
        if Primitive.to_proto(resource.code):
            res.code = Primitive.to_proto(resource.code)
        if Primitive.to_proto(resource.message):
            res.message = Primitive.to_proto(resource.message)
        if InstanceRepairRecipeStepsStatusDetailsArray.to_proto(resource.details):
            res.details.extend(
                InstanceRepairRecipeStepsStatusDetailsArray.to_proto(resource.details)
            )
        return res

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return None

        return InstanceRepairRecipeStepsStatus(
            code=Primitive.from_proto(resource.code),
            message=Primitive.from_proto(resource.message),
            details=InstanceRepairRecipeStepsStatusDetailsArray.from_proto(
                resource.details
            ),
        )


class InstanceRepairRecipeStepsStatusArray(object):
    @classmethod
    def to_proto(self, resources):
        if not resources:
            return resources
        return [InstanceRepairRecipeStepsStatus.to_proto(i) for i in resources]

    @classmethod
    def from_proto(self, resources):
        return [InstanceRepairRecipeStepsStatus.from_proto(i) for i in resources]


class InstanceRepairRecipeStepsStatusDetails(object):
    def __init__(self, type_url: str = None, value: str = None):
        self.type_url = type_url
        self.value = value

    @classmethod
    def to_proto(self, resource):
        if not resource:
            return None

        res = instance_pb2.Tier2AlphaInstanceRepairRecipeStepsStatusDetails()
        if Primitive.to_proto(resource.type_url):
            res.type_url = Primitive.to_proto(resource.type_url)
        if Primitive.to_proto(resource.value):
            res.value = Primitive.to_proto(resource.value)
        return res

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return None

        return InstanceRepairRecipeStepsStatusDetails(
            type_url=Primitive.from_proto(resource.type_url),
            value=Primitive.from_proto(resource.value),
        )


class InstanceRepairRecipeStepsStatusDetailsArray(object):
    @classmethod
    def to_proto(self, resources):
        if not resources:
            return resources
        return [InstanceRepairRecipeStepsStatusDetails.to_proto(i) for i in resources]

    @classmethod
    def from_proto(self, resources):
        return [InstanceRepairRecipeStepsStatusDetails.from_proto(i) for i in resources]


class InstanceRepairRecipeStepsQuotaRequestDeltas(object):
    def __init__(
        self,
        metric_name: str = None,
        amount: int = None,
        quota_location_name: str = None,
    ):
        self.metric_name = metric_name
        self.amount = amount
        self.quota_location_name = quota_location_name

    @classmethod
    def to_proto(self, resource):
        if not resource:
            return None

        res = instance_pb2.Tier2AlphaInstanceRepairRecipeStepsQuotaRequestDeltas()
        if Primitive.to_proto(resource.metric_name):
            res.metric_name = Primitive.to_proto(resource.metric_name)
        if Primitive.to_proto(resource.amount):
            res.amount = Primitive.to_proto(resource.amount)
        if Primitive.to_proto(resource.quota_location_name):
            res.quota_location_name = Primitive.to_proto(resource.quota_location_name)
        return res

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return None

        return InstanceRepairRecipeStepsQuotaRequestDeltas(
            metric_name=Primitive.from_proto(resource.metric_name),
            amount=Primitive.from_proto(resource.amount),
            quota_location_name=Primitive.from_proto(resource.quota_location_name),
        )


class InstanceRepairRecipeStepsQuotaRequestDeltasArray(object):
    @classmethod
    def to_proto(self, resources):
        if not resources:
            return resources
        return [
            InstanceRepairRecipeStepsQuotaRequestDeltas.to_proto(i) for i in resources
        ]

    @classmethod
    def from_proto(self, resources):
        return [
            InstanceRepairRecipeStepsQuotaRequestDeltas.from_proto(i) for i in resources
        ]


class InstanceRepairRecipeStepsPreprocessUpdate(object):
    def __init__(
        self, latency_slo_bucket_name: str = None, public_operation_metadata: str = None
    ):
        self.latency_slo_bucket_name = latency_slo_bucket_name
        self.public_operation_metadata = public_operation_metadata

    @classmethod
    def to_proto(self, resource):
        if not resource:
            return None

        res = instance_pb2.Tier2AlphaInstanceRepairRecipeStepsPreprocessUpdate()
        if Primitive.to_proto(resource.latency_slo_bucket_name):
            res.latency_slo_bucket_name = Primitive.to_proto(
                resource.latency_slo_bucket_name
            )
        if Primitive.to_proto(resource.public_operation_metadata):
            res.public_operation_metadata = Primitive.to_proto(
                resource.public_operation_metadata
            )
        return res

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return None

        return InstanceRepairRecipeStepsPreprocessUpdate(
            latency_slo_bucket_name=Primitive.from_proto(
                resource.latency_slo_bucket_name
            ),
            public_operation_metadata=Primitive.from_proto(
                resource.public_operation_metadata
            ),
        )


class InstanceRepairRecipeStepsPreprocessUpdateArray(object):
    @classmethod
    def to_proto(self, resources):
        if not resources:
            return resources
        return [
            InstanceRepairRecipeStepsPreprocessUpdate.to_proto(i) for i in resources
        ]

    @classmethod
    def from_proto(self, resources):
        return [
            InstanceRepairRecipeStepsPreprocessUpdate.from_proto(i) for i in resources
        ]


class InstanceRepairRecipeStepsRequestedTenantProject(object):
    def __init__(self, tag: str = None, folder: str = None, scope: str = None):
        self.tag = tag
        self.folder = folder
        self.scope = scope

    @classmethod
    def to_proto(self, resource):
        if not resource:
            return None

        res = instance_pb2.Tier2AlphaInstanceRepairRecipeStepsRequestedTenantProject()
        if Primitive.to_proto(resource.tag):
            res.tag = Primitive.to_proto(resource.tag)
        if Primitive.to_proto(resource.folder):
            res.folder = Primitive.to_proto(resource.folder)
        if InstanceRepairRecipeStepsRequestedTenantProjectScopeEnum.to_proto(
            resource.scope
        ):
            res.scope = InstanceRepairRecipeStepsRequestedTenantProjectScopeEnum.to_proto(
                resource.scope
            )
        return res

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return None

        return InstanceRepairRecipeStepsRequestedTenantProject(
            tag=Primitive.from_proto(resource.tag),
            folder=Primitive.from_proto(resource.folder),
            scope=InstanceRepairRecipeStepsRequestedTenantProjectScopeEnum.from_proto(
                resource.scope
            ),
        )


class InstanceRepairRecipeStepsRequestedTenantProjectArray(object):
    @classmethod
    def to_proto(self, resources):
        if not resources:
            return resources
        return [
            InstanceRepairRecipeStepsRequestedTenantProject.to_proto(i)
            for i in resources
        ]

    @classmethod
    def from_proto(self, resources):
        return [
            InstanceRepairRecipeStepsRequestedTenantProject.from_proto(i)
            for i in resources
        ]


class InstanceRepairRecipeStepsPermissionsInfo(object):
    def __init__(
        self,
        policy_name: dict = None,
        iam_permissions: list = None,
        resource_path: str = None,
        api_attrs: dict = None,
        policy_name_mode: str = None,
        resource: dict = None,
    ):
        self.policy_name = policy_name
        self.iam_permissions = iam_permissions
        self.resource_path = resource_path
        self.api_attrs = api_attrs
        self.policy_name_mode = policy_name_mode
        self.resource = resource

    @classmethod
    def to_proto(self, resource):
        if not resource:
            return None

        res = instance_pb2.Tier2AlphaInstanceRepairRecipeStepsPermissionsInfo()
        if InstanceRepairRecipeStepsPermissionsInfoPolicyName.to_proto(
            resource.policy_name
        ):
            res.policy_name.CopyFrom(
                InstanceRepairRecipeStepsPermissionsInfoPolicyName.to_proto(
                    resource.policy_name
                )
            )
        else:
            res.ClearField("policy_name")
        if InstanceRepairRecipeStepsPermissionsInfoIamPermissionsArray.to_proto(
            resource.iam_permissions
        ):
            res.iam_permissions.extend(
                InstanceRepairRecipeStepsPermissionsInfoIamPermissionsArray.to_proto(
                    resource.iam_permissions
                )
            )
        if Primitive.to_proto(resource.resource_path):
            res.resource_path = Primitive.to_proto(resource.resource_path)
        if InstanceGoogleprotobufstruct.to_proto(resource.api_attrs):
            res.api_attrs.CopyFrom(
                InstanceGoogleprotobufstruct.to_proto(resource.api_attrs)
            )
        else:
            res.ClearField("api_attrs")
        if InstanceRepairRecipeStepsPermissionsInfoPolicyNameModeEnum.to_proto(
            resource.policy_name_mode
        ):
            res.policy_name_mode = InstanceRepairRecipeStepsPermissionsInfoPolicyNameModeEnum.to_proto(
                resource.policy_name_mode
            )
        if InstanceRepairRecipeStepsPermissionsInfoResource.to_proto(resource.resource):
            res.resource.CopyFrom(
                InstanceRepairRecipeStepsPermissionsInfoResource.to_proto(
                    resource.resource
                )
            )
        else:
            res.ClearField("resource")
        return res

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return None

        return InstanceRepairRecipeStepsPermissionsInfo(
            policy_name=InstanceRepairRecipeStepsPermissionsInfoPolicyName.from_proto(
                resource.policy_name
            ),
            iam_permissions=InstanceRepairRecipeStepsPermissionsInfoIamPermissionsArray.from_proto(
                resource.iam_permissions
            ),
            resource_path=Primitive.from_proto(resource.resource_path),
            api_attrs=InstanceGoogleprotobufstruct.from_proto(resource.api_attrs),
            policy_name_mode=InstanceRepairRecipeStepsPermissionsInfoPolicyNameModeEnum.from_proto(
                resource.policy_name_mode
            ),
            resource=InstanceRepairRecipeStepsPermissionsInfoResource.from_proto(
                resource.resource
            ),
        )


class InstanceRepairRecipeStepsPermissionsInfoArray(object):
    @classmethod
    def to_proto(self, resources):
        if not resources:
            return resources
        return [InstanceRepairRecipeStepsPermissionsInfo.to_proto(i) for i in resources]

    @classmethod
    def from_proto(self, resources):
        return [
            InstanceRepairRecipeStepsPermissionsInfo.from_proto(i) for i in resources
        ]


class InstanceRepairRecipeStepsPermissionsInfoPolicyName(object):
    def __init__(self, type: str = None, id: str = None, region: str = None):
        self.type = type
        self.id = id
        self.region = region

    @classmethod
    def to_proto(self, resource):
        if not resource:
            return None

        res = (
            instance_pb2.Tier2AlphaInstanceRepairRecipeStepsPermissionsInfoPolicyName()
        )
        if Primitive.to_proto(resource.type):
            res.type = Primitive.to_proto(resource.type)
        if Primitive.to_proto(resource.id):
            res.id = Primitive.to_proto(resource.id)
        if Primitive.to_proto(resource.region):
            res.region = Primitive.to_proto(resource.region)
        return res

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return None

        return InstanceRepairRecipeStepsPermissionsInfoPolicyName(
            type=Primitive.from_proto(resource.type),
            id=Primitive.from_proto(resource.id),
            region=Primitive.from_proto(resource.region),
        )


class InstanceRepairRecipeStepsPermissionsInfoPolicyNameArray(object):
    @classmethod
    def to_proto(self, resources):
        if not resources:
            return resources
        return [
            InstanceRepairRecipeStepsPermissionsInfoPolicyName.to_proto(i)
            for i in resources
        ]

    @classmethod
    def from_proto(self, resources):
        return [
            InstanceRepairRecipeStepsPermissionsInfoPolicyName.from_proto(i)
            for i in resources
        ]


class InstanceRepairRecipeStepsPermissionsInfoIamPermissions(object):
    def __init__(self, permission: str = None):
        self.permission = permission

    @classmethod
    def to_proto(self, resource):
        if not resource:
            return None

        res = (
            instance_pb2.Tier2AlphaInstanceRepairRecipeStepsPermissionsInfoIamPermissions()
        )
        if Primitive.to_proto(resource.permission):
            res.permission = Primitive.to_proto(resource.permission)
        return res

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return None

        return InstanceRepairRecipeStepsPermissionsInfoIamPermissions(
            permission=Primitive.from_proto(resource.permission),
        )


class InstanceRepairRecipeStepsPermissionsInfoIamPermissionsArray(object):
    @classmethod
    def to_proto(self, resources):
        if not resources:
            return resources
        return [
            InstanceRepairRecipeStepsPermissionsInfoIamPermissions.to_proto(i)
            for i in resources
        ]

    @classmethod
    def from_proto(self, resources):
        return [
            InstanceRepairRecipeStepsPermissionsInfoIamPermissions.from_proto(i)
            for i in resources
        ]


class InstanceRepairRecipeStepsPermissionsInfoResource(object):
    def __init__(
        self,
        name: str = None,
        type: str = None,
        service: str = None,
        labels: dict = None,
    ):
        self.name = name
        self.type = type
        self.service = service
        self.labels = labels

    @classmethod
    def to_proto(self, resource):
        if not resource:
            return None

        res = instance_pb2.Tier2AlphaInstanceRepairRecipeStepsPermissionsInfoResource()
        if Primitive.to_proto(resource.name):
            res.name = Primitive.to_proto(resource.name)
        if Primitive.to_proto(resource.type):
            res.type = Primitive.to_proto(resource.type)
        if Primitive.to_proto(resource.service):
            res.service = Primitive.to_proto(resource.service)
        if Primitive.to_proto(resource.labels):
            res.labels = Primitive.to_proto(resource.labels)
        return res

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return None

        return InstanceRepairRecipeStepsPermissionsInfoResource(
            name=Primitive.from_proto(resource.name),
            type=Primitive.from_proto(resource.type),
            service=Primitive.from_proto(resource.service),
            labels=Primitive.from_proto(resource.labels),
        )


class InstanceRepairRecipeStepsPermissionsInfoResourceArray(object):
    @classmethod
    def to_proto(self, resources):
        if not resources:
            return resources
        return [
            InstanceRepairRecipeStepsPermissionsInfoResource.to_proto(i)
            for i in resources
        ]

    @classmethod
    def from_proto(self, resources):
        return [
            InstanceRepairRecipeStepsPermissionsInfoResource.from_proto(i)
            for i in resources
        ]


class InstanceRepairRecipeStepsKeyNotificationsUpdate(object):
    def __init__(self, key_notifications_info: dict = None):
        self.key_notifications_info = key_notifications_info

    @classmethod
    def to_proto(self, resource):
        if not resource:
            return None

        res = instance_pb2.Tier2AlphaInstanceRepairRecipeStepsKeyNotificationsUpdate()
        if InstanceRepairRecipeStepsKeyNotificationsUpdateKeyNotificationsInfo.to_proto(
            resource.key_notifications_info
        ):
            res.key_notifications_info.CopyFrom(
                InstanceRepairRecipeStepsKeyNotificationsUpdateKeyNotificationsInfo.to_proto(
                    resource.key_notifications_info
                )
            )
        else:
            res.ClearField("key_notifications_info")
        return res

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return None

        return InstanceRepairRecipeStepsKeyNotificationsUpdate(
            key_notifications_info=InstanceRepairRecipeStepsKeyNotificationsUpdateKeyNotificationsInfo.from_proto(
                resource.key_notifications_info
            ),
        )


class InstanceRepairRecipeStepsKeyNotificationsUpdateArray(object):
    @classmethod
    def to_proto(self, resources):
        if not resources:
            return resources
        return [
            InstanceRepairRecipeStepsKeyNotificationsUpdate.to_proto(i)
            for i in resources
        ]

    @classmethod
    def from_proto(self, resources):
        return [
            InstanceRepairRecipeStepsKeyNotificationsUpdate.from_proto(i)
            for i in resources
        ]


class InstanceRepairRecipeStepsKeyNotificationsUpdateKeyNotificationsInfo(object):
    def __init__(
        self,
        data_version: int = None,
        delegate: str = None,
        key_notification_configs: list = None,
    ):
        self.data_version = data_version
        self.delegate = delegate
        self.key_notification_configs = key_notification_configs

    @classmethod
    def to_proto(self, resource):
        if not resource:
            return None

        res = (
            instance_pb2.Tier2AlphaInstanceRepairRecipeStepsKeyNotificationsUpdateKeyNotificationsInfo()
        )
        if Primitive.to_proto(resource.data_version):
            res.data_version = Primitive.to_proto(resource.data_version)
        if Primitive.to_proto(resource.delegate):
            res.delegate = Primitive.to_proto(resource.delegate)
        if InstanceRepairRecipeStepsKeyNotificationsUpdateKeyNotificationsInfoKeyNotificationConfigsArray.to_proto(
            resource.key_notification_configs
        ):
            res.key_notification_configs.extend(
                InstanceRepairRecipeStepsKeyNotificationsUpdateKeyNotificationsInfoKeyNotificationConfigsArray.to_proto(
                    resource.key_notification_configs
                )
            )
        return res

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return None

        return InstanceRepairRecipeStepsKeyNotificationsUpdateKeyNotificationsInfo(
            data_version=Primitive.from_proto(resource.data_version),
            delegate=Primitive.from_proto(resource.delegate),
            key_notification_configs=InstanceRepairRecipeStepsKeyNotificationsUpdateKeyNotificationsInfoKeyNotificationConfigsArray.from_proto(
                resource.key_notification_configs
            ),
        )


class InstanceRepairRecipeStepsKeyNotificationsUpdateKeyNotificationsInfoArray(object):
    @classmethod
    def to_proto(self, resources):
        if not resources:
            return resources
        return [
            InstanceRepairRecipeStepsKeyNotificationsUpdateKeyNotificationsInfo.to_proto(
                i
            )
            for i in resources
        ]

    @classmethod
    def from_proto(self, resources):
        return [
            InstanceRepairRecipeStepsKeyNotificationsUpdateKeyNotificationsInfo.from_proto(
                i
            )
            for i in resources
        ]


class InstanceRepairRecipeStepsKeyNotificationsUpdateKeyNotificationsInfoKeyNotificationConfigs(
    object
):
    def __init__(
        self,
        key_or_version_name: str = None,
        grant: str = None,
        delegator_gaia_id: int = None,
    ):
        self.key_or_version_name = key_or_version_name
        self.grant = grant
        self.delegator_gaia_id = delegator_gaia_id

    @classmethod
    def to_proto(self, resource):
        if not resource:
            return None

        res = (
            instance_pb2.Tier2AlphaInstanceRepairRecipeStepsKeyNotificationsUpdateKeyNotificationsInfoKeyNotificationConfigs()
        )
        if Primitive.to_proto(resource.key_or_version_name):
            res.key_or_version_name = Primitive.to_proto(resource.key_or_version_name)
        if Primitive.to_proto(resource.grant):
            res.grant = Primitive.to_proto(resource.grant)
        if Primitive.to_proto(resource.delegator_gaia_id):
            res.delegator_gaia_id = Primitive.to_proto(resource.delegator_gaia_id)
        return res

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return None

        return InstanceRepairRecipeStepsKeyNotificationsUpdateKeyNotificationsInfoKeyNotificationConfigs(
            key_or_version_name=Primitive.from_proto(resource.key_or_version_name),
            grant=Primitive.from_proto(resource.grant),
            delegator_gaia_id=Primitive.from_proto(resource.delegator_gaia_id),
        )


class InstanceRepairRecipeStepsKeyNotificationsUpdateKeyNotificationsInfoKeyNotificationConfigsArray(
    object
):
    @classmethod
    def to_proto(self, resources):
        if not resources:
            return resources
        return [
            InstanceRepairRecipeStepsKeyNotificationsUpdateKeyNotificationsInfoKeyNotificationConfigs.to_proto(
                i
            )
            for i in resources
        ]

    @classmethod
    def from_proto(self, resources):
        return [
            InstanceRepairRecipeStepsKeyNotificationsUpdateKeyNotificationsInfoKeyNotificationConfigs.from_proto(
                i
            )
            for i in resources
        ]


class InstancePreprocessDeleteRecipe(object):
    def __init__(
        self,
        steps: list = None,
        honor_cancel_request: bool = None,
        ignore_recipe_after: int = None,
        verify_deadline_seconds_below: float = None,
        populate_operation_result: bool = None,
        readonly_recipe_start_time: str = None,
        resource_names_stored_in_clh_with_delay: list = None,
        delay_to_store_resources_in_clh_db_nanos: int = None,
    ):
        self.steps = steps
        self.honor_cancel_request = honor_cancel_request
        self.ignore_recipe_after = ignore_recipe_after
        self.verify_deadline_seconds_below = verify_deadline_seconds_below
        self.populate_operation_result = populate_operation_result
        self.readonly_recipe_start_time = readonly_recipe_start_time
        self.resource_names_stored_in_clh_with_delay = (
            resource_names_stored_in_clh_with_delay
        )
        self.delay_to_store_resources_in_clh_db_nanos = (
            delay_to_store_resources_in_clh_db_nanos
        )

    @classmethod
    def to_proto(self, resource):
        if not resource:
            return None

        res = instance_pb2.Tier2AlphaInstancePreprocessDeleteRecipe()
        if InstancePreprocessDeleteRecipeStepsArray.to_proto(resource.steps):
            res.steps.extend(
                InstancePreprocessDeleteRecipeStepsArray.to_proto(resource.steps)
            )
        if Primitive.to_proto(resource.honor_cancel_request):
            res.honor_cancel_request = Primitive.to_proto(resource.honor_cancel_request)
        if Primitive.to_proto(resource.ignore_recipe_after):
            res.ignore_recipe_after = Primitive.to_proto(resource.ignore_recipe_after)
        if Primitive.to_proto(resource.verify_deadline_seconds_below):
            res.verify_deadline_seconds_below = Primitive.to_proto(
                resource.verify_deadline_seconds_below
            )
        if Primitive.to_proto(resource.populate_operation_result):
            res.populate_operation_result = Primitive.to_proto(
                resource.populate_operation_result
            )
        if Primitive.to_proto(resource.readonly_recipe_start_time):
            res.readonly_recipe_start_time = Primitive.to_proto(
                resource.readonly_recipe_start_time
            )
        if Primitive.to_proto(resource.resource_names_stored_in_clh_with_delay):
            res.resource_names_stored_in_clh_with_delay.extend(
                Primitive.to_proto(resource.resource_names_stored_in_clh_with_delay)
            )
        if Primitive.to_proto(resource.delay_to_store_resources_in_clh_db_nanos):
            res.delay_to_store_resources_in_clh_db_nanos = Primitive.to_proto(
                resource.delay_to_store_resources_in_clh_db_nanos
            )
        return res

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return None

        return InstancePreprocessDeleteRecipe(
            steps=InstancePreprocessDeleteRecipeStepsArray.from_proto(resource.steps),
            honor_cancel_request=Primitive.from_proto(resource.honor_cancel_request),
            ignore_recipe_after=Primitive.from_proto(resource.ignore_recipe_after),
            verify_deadline_seconds_below=Primitive.from_proto(
                resource.verify_deadline_seconds_below
            ),
            populate_operation_result=Primitive.from_proto(
                resource.populate_operation_result
            ),
            readonly_recipe_start_time=Primitive.from_proto(
                resource.readonly_recipe_start_time
            ),
            resource_names_stored_in_clh_with_delay=Primitive.from_proto(
                resource.resource_names_stored_in_clh_with_delay
            ),
            delay_to_store_resources_in_clh_db_nanos=Primitive.from_proto(
                resource.delay_to_store_resources_in_clh_db_nanos
            ),
        )


class InstancePreprocessDeleteRecipeArray(object):
    @classmethod
    def to_proto(self, resources):
        if not resources:
            return resources
        return [InstancePreprocessDeleteRecipe.to_proto(i) for i in resources]

    @classmethod
    def from_proto(self, resources):
        return [InstancePreprocessDeleteRecipe.from_proto(i) for i in resources]


class InstancePreprocessDeleteRecipeSteps(object):
    def __init__(
        self,
        relative_time: int = None,
        sleep_duration: int = None,
        action: str = None,
        status: dict = None,
        error_space: str = None,
        p4_service_account: str = None,
        resource_metadata_size: int = None,
        description: str = None,
        updated_repeat_operation_delay_sec: float = None,
        quota_request_deltas: list = None,
        preprocess_update: dict = None,
        public_operation_metadata: str = None,
        requested_tenant_project: dict = None,
        permissions_info: list = None,
        key_notifications_update: dict = None,
        clh_data_update_time: str = None,
        public_error_message: str = None,
    ):
        self.relative_time = relative_time
        self.sleep_duration = sleep_duration
        self.action = action
        self.status = status
        self.error_space = error_space
        self.p4_service_account = p4_service_account
        self.resource_metadata_size = resource_metadata_size
        self.description = description
        self.updated_repeat_operation_delay_sec = updated_repeat_operation_delay_sec
        self.quota_request_deltas = quota_request_deltas
        self.preprocess_update = preprocess_update
        self.public_operation_metadata = public_operation_metadata
        self.requested_tenant_project = requested_tenant_project
        self.permissions_info = permissions_info
        self.key_notifications_update = key_notifications_update
        self.clh_data_update_time = clh_data_update_time
        self.public_error_message = public_error_message

    @classmethod
    def to_proto(self, resource):
        if not resource:
            return None

        res = instance_pb2.Tier2AlphaInstancePreprocessDeleteRecipeSteps()
        if Primitive.to_proto(resource.relative_time):
            res.relative_time = Primitive.to_proto(resource.relative_time)
        if Primitive.to_proto(resource.sleep_duration):
            res.sleep_duration = Primitive.to_proto(resource.sleep_duration)
        if InstancePreprocessDeleteRecipeStepsActionEnum.to_proto(resource.action):
            res.action = InstancePreprocessDeleteRecipeStepsActionEnum.to_proto(
                resource.action
            )
        if InstancePreprocessDeleteRecipeStepsStatus.to_proto(resource.status):
            res.status.CopyFrom(
                InstancePreprocessDeleteRecipeStepsStatus.to_proto(resource.status)
            )
        else:
            res.ClearField("status")
        if Primitive.to_proto(resource.error_space):
            res.error_space = Primitive.to_proto(resource.error_space)
        if Primitive.to_proto(resource.p4_service_account):
            res.p4_service_account = Primitive.to_proto(resource.p4_service_account)
        if Primitive.to_proto(resource.resource_metadata_size):
            res.resource_metadata_size = Primitive.to_proto(
                resource.resource_metadata_size
            )
        if Primitive.to_proto(resource.description):
            res.description = Primitive.to_proto(resource.description)
        if Primitive.to_proto(resource.updated_repeat_operation_delay_sec):
            res.updated_repeat_operation_delay_sec = Primitive.to_proto(
                resource.updated_repeat_operation_delay_sec
            )
        if InstancePreprocessDeleteRecipeStepsQuotaRequestDeltasArray.to_proto(
            resource.quota_request_deltas
        ):
            res.quota_request_deltas.extend(
                InstancePreprocessDeleteRecipeStepsQuotaRequestDeltasArray.to_proto(
                    resource.quota_request_deltas
                )
            )
        if InstancePreprocessDeleteRecipeStepsPreprocessUpdate.to_proto(
            resource.preprocess_update
        ):
            res.preprocess_update.CopyFrom(
                InstancePreprocessDeleteRecipeStepsPreprocessUpdate.to_proto(
                    resource.preprocess_update
                )
            )
        else:
            res.ClearField("preprocess_update")
        if Primitive.to_proto(resource.public_operation_metadata):
            res.public_operation_metadata = Primitive.to_proto(
                resource.public_operation_metadata
            )
        if InstancePreprocessDeleteRecipeStepsRequestedTenantProject.to_proto(
            resource.requested_tenant_project
        ):
            res.requested_tenant_project.CopyFrom(
                InstancePreprocessDeleteRecipeStepsRequestedTenantProject.to_proto(
                    resource.requested_tenant_project
                )
            )
        else:
            res.ClearField("requested_tenant_project")
        if InstancePreprocessDeleteRecipeStepsPermissionsInfoArray.to_proto(
            resource.permissions_info
        ):
            res.permissions_info.extend(
                InstancePreprocessDeleteRecipeStepsPermissionsInfoArray.to_proto(
                    resource.permissions_info
                )
            )
        if InstancePreprocessDeleteRecipeStepsKeyNotificationsUpdate.to_proto(
            resource.key_notifications_update
        ):
            res.key_notifications_update.CopyFrom(
                InstancePreprocessDeleteRecipeStepsKeyNotificationsUpdate.to_proto(
                    resource.key_notifications_update
                )
            )
        else:
            res.ClearField("key_notifications_update")
        if Primitive.to_proto(resource.clh_data_update_time):
            res.clh_data_update_time = Primitive.to_proto(resource.clh_data_update_time)
        if Primitive.to_proto(resource.public_error_message):
            res.public_error_message = Primitive.to_proto(resource.public_error_message)
        return res

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return None

        return InstancePreprocessDeleteRecipeSteps(
            relative_time=Primitive.from_proto(resource.relative_time),
            sleep_duration=Primitive.from_proto(resource.sleep_duration),
            action=InstancePreprocessDeleteRecipeStepsActionEnum.from_proto(
                resource.action
            ),
            status=InstancePreprocessDeleteRecipeStepsStatus.from_proto(
                resource.status
            ),
            error_space=Primitive.from_proto(resource.error_space),
            p4_service_account=Primitive.from_proto(resource.p4_service_account),
            resource_metadata_size=Primitive.from_proto(
                resource.resource_metadata_size
            ),
            description=Primitive.from_proto(resource.description),
            updated_repeat_operation_delay_sec=Primitive.from_proto(
                resource.updated_repeat_operation_delay_sec
            ),
            quota_request_deltas=InstancePreprocessDeleteRecipeStepsQuotaRequestDeltasArray.from_proto(
                resource.quota_request_deltas
            ),
            preprocess_update=InstancePreprocessDeleteRecipeStepsPreprocessUpdate.from_proto(
                resource.preprocess_update
            ),
            public_operation_metadata=Primitive.from_proto(
                resource.public_operation_metadata
            ),
            requested_tenant_project=InstancePreprocessDeleteRecipeStepsRequestedTenantProject.from_proto(
                resource.requested_tenant_project
            ),
            permissions_info=InstancePreprocessDeleteRecipeStepsPermissionsInfoArray.from_proto(
                resource.permissions_info
            ),
            key_notifications_update=InstancePreprocessDeleteRecipeStepsKeyNotificationsUpdate.from_proto(
                resource.key_notifications_update
            ),
            clh_data_update_time=Primitive.from_proto(resource.clh_data_update_time),
            public_error_message=Primitive.from_proto(resource.public_error_message),
        )


class InstancePreprocessDeleteRecipeStepsArray(object):
    @classmethod
    def to_proto(self, resources):
        if not resources:
            return resources
        return [InstancePreprocessDeleteRecipeSteps.to_proto(i) for i in resources]

    @classmethod
    def from_proto(self, resources):
        return [InstancePreprocessDeleteRecipeSteps.from_proto(i) for i in resources]


class InstancePreprocessDeleteRecipeStepsStatus(object):
    def __init__(self, code: int = None, message: str = None, details: list = None):
        self.code = code
        self.message = message
        self.details = details

    @classmethod
    def to_proto(self, resource):
        if not resource:
            return None

        res = instance_pb2.Tier2AlphaInstancePreprocessDeleteRecipeStepsStatus()
        if Primitive.to_proto(resource.code):
            res.code = Primitive.to_proto(resource.code)
        if Primitive.to_proto(resource.message):
            res.message = Primitive.to_proto(resource.message)
        if InstancePreprocessDeleteRecipeStepsStatusDetailsArray.to_proto(
            resource.details
        ):
            res.details.extend(
                InstancePreprocessDeleteRecipeStepsStatusDetailsArray.to_proto(
                    resource.details
                )
            )
        return res

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return None

        return InstancePreprocessDeleteRecipeStepsStatus(
            code=Primitive.from_proto(resource.code),
            message=Primitive.from_proto(resource.message),
            details=InstancePreprocessDeleteRecipeStepsStatusDetailsArray.from_proto(
                resource.details
            ),
        )


class InstancePreprocessDeleteRecipeStepsStatusArray(object):
    @classmethod
    def to_proto(self, resources):
        if not resources:
            return resources
        return [
            InstancePreprocessDeleteRecipeStepsStatus.to_proto(i) for i in resources
        ]

    @classmethod
    def from_proto(self, resources):
        return [
            InstancePreprocessDeleteRecipeStepsStatus.from_proto(i) for i in resources
        ]


class InstancePreprocessDeleteRecipeStepsStatusDetails(object):
    def __init__(self, type_url: str = None, value: str = None):
        self.type_url = type_url
        self.value = value

    @classmethod
    def to_proto(self, resource):
        if not resource:
            return None

        res = instance_pb2.Tier2AlphaInstancePreprocessDeleteRecipeStepsStatusDetails()
        if Primitive.to_proto(resource.type_url):
            res.type_url = Primitive.to_proto(resource.type_url)
        if Primitive.to_proto(resource.value):
            res.value = Primitive.to_proto(resource.value)
        return res

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return None

        return InstancePreprocessDeleteRecipeStepsStatusDetails(
            type_url=Primitive.from_proto(resource.type_url),
            value=Primitive.from_proto(resource.value),
        )


class InstancePreprocessDeleteRecipeStepsStatusDetailsArray(object):
    @classmethod
    def to_proto(self, resources):
        if not resources:
            return resources
        return [
            InstancePreprocessDeleteRecipeStepsStatusDetails.to_proto(i)
            for i in resources
        ]

    @classmethod
    def from_proto(self, resources):
        return [
            InstancePreprocessDeleteRecipeStepsStatusDetails.from_proto(i)
            for i in resources
        ]


class InstancePreprocessDeleteRecipeStepsQuotaRequestDeltas(object):
    def __init__(
        self,
        metric_name: str = None,
        amount: int = None,
        quota_location_name: str = None,
    ):
        self.metric_name = metric_name
        self.amount = amount
        self.quota_location_name = quota_location_name

    @classmethod
    def to_proto(self, resource):
        if not resource:
            return None

        res = (
            instance_pb2.Tier2AlphaInstancePreprocessDeleteRecipeStepsQuotaRequestDeltas()
        )
        if Primitive.to_proto(resource.metric_name):
            res.metric_name = Primitive.to_proto(resource.metric_name)
        if Primitive.to_proto(resource.amount):
            res.amount = Primitive.to_proto(resource.amount)
        if Primitive.to_proto(resource.quota_location_name):
            res.quota_location_name = Primitive.to_proto(resource.quota_location_name)
        return res

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return None

        return InstancePreprocessDeleteRecipeStepsQuotaRequestDeltas(
            metric_name=Primitive.from_proto(resource.metric_name),
            amount=Primitive.from_proto(resource.amount),
            quota_location_name=Primitive.from_proto(resource.quota_location_name),
        )


class InstancePreprocessDeleteRecipeStepsQuotaRequestDeltasArray(object):
    @classmethod
    def to_proto(self, resources):
        if not resources:
            return resources
        return [
            InstancePreprocessDeleteRecipeStepsQuotaRequestDeltas.to_proto(i)
            for i in resources
        ]

    @classmethod
    def from_proto(self, resources):
        return [
            InstancePreprocessDeleteRecipeStepsQuotaRequestDeltas.from_proto(i)
            for i in resources
        ]


class InstancePreprocessDeleteRecipeStepsPreprocessUpdate(object):
    def __init__(
        self, latency_slo_bucket_name: str = None, public_operation_metadata: str = None
    ):
        self.latency_slo_bucket_name = latency_slo_bucket_name
        self.public_operation_metadata = public_operation_metadata

    @classmethod
    def to_proto(self, resource):
        if not resource:
            return None

        res = (
            instance_pb2.Tier2AlphaInstancePreprocessDeleteRecipeStepsPreprocessUpdate()
        )
        if Primitive.to_proto(resource.latency_slo_bucket_name):
            res.latency_slo_bucket_name = Primitive.to_proto(
                resource.latency_slo_bucket_name
            )
        if Primitive.to_proto(resource.public_operation_metadata):
            res.public_operation_metadata = Primitive.to_proto(
                resource.public_operation_metadata
            )
        return res

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return None

        return InstancePreprocessDeleteRecipeStepsPreprocessUpdate(
            latency_slo_bucket_name=Primitive.from_proto(
                resource.latency_slo_bucket_name
            ),
            public_operation_metadata=Primitive.from_proto(
                resource.public_operation_metadata
            ),
        )


class InstancePreprocessDeleteRecipeStepsPreprocessUpdateArray(object):
    @classmethod
    def to_proto(self, resources):
        if not resources:
            return resources
        return [
            InstancePreprocessDeleteRecipeStepsPreprocessUpdate.to_proto(i)
            for i in resources
        ]

    @classmethod
    def from_proto(self, resources):
        return [
            InstancePreprocessDeleteRecipeStepsPreprocessUpdate.from_proto(i)
            for i in resources
        ]


class InstancePreprocessDeleteRecipeStepsRequestedTenantProject(object):
    def __init__(self, tag: str = None, folder: str = None, scope: str = None):
        self.tag = tag
        self.folder = folder
        self.scope = scope

    @classmethod
    def to_proto(self, resource):
        if not resource:
            return None

        res = (
            instance_pb2.Tier2AlphaInstancePreprocessDeleteRecipeStepsRequestedTenantProject()
        )
        if Primitive.to_proto(resource.tag):
            res.tag = Primitive.to_proto(resource.tag)
        if Primitive.to_proto(resource.folder):
            res.folder = Primitive.to_proto(resource.folder)
        if InstancePreprocessDeleteRecipeStepsRequestedTenantProjectScopeEnum.to_proto(
            resource.scope
        ):
            res.scope = InstancePreprocessDeleteRecipeStepsRequestedTenantProjectScopeEnum.to_proto(
                resource.scope
            )
        return res

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return None

        return InstancePreprocessDeleteRecipeStepsRequestedTenantProject(
            tag=Primitive.from_proto(resource.tag),
            folder=Primitive.from_proto(resource.folder),
            scope=InstancePreprocessDeleteRecipeStepsRequestedTenantProjectScopeEnum.from_proto(
                resource.scope
            ),
        )


class InstancePreprocessDeleteRecipeStepsRequestedTenantProjectArray(object):
    @classmethod
    def to_proto(self, resources):
        if not resources:
            return resources
        return [
            InstancePreprocessDeleteRecipeStepsRequestedTenantProject.to_proto(i)
            for i in resources
        ]

    @classmethod
    def from_proto(self, resources):
        return [
            InstancePreprocessDeleteRecipeStepsRequestedTenantProject.from_proto(i)
            for i in resources
        ]


class InstancePreprocessDeleteRecipeStepsPermissionsInfo(object):
    def __init__(
        self,
        policy_name: dict = None,
        iam_permissions: list = None,
        resource_path: str = None,
        api_attrs: dict = None,
        policy_name_mode: str = None,
        resource: dict = None,
    ):
        self.policy_name = policy_name
        self.iam_permissions = iam_permissions
        self.resource_path = resource_path
        self.api_attrs = api_attrs
        self.policy_name_mode = policy_name_mode
        self.resource = resource

    @classmethod
    def to_proto(self, resource):
        if not resource:
            return None

        res = (
            instance_pb2.Tier2AlphaInstancePreprocessDeleteRecipeStepsPermissionsInfo()
        )
        if InstancePreprocessDeleteRecipeStepsPermissionsInfoPolicyName.to_proto(
            resource.policy_name
        ):
            res.policy_name.CopyFrom(
                InstancePreprocessDeleteRecipeStepsPermissionsInfoPolicyName.to_proto(
                    resource.policy_name
                )
            )
        else:
            res.ClearField("policy_name")
        if InstancePreprocessDeleteRecipeStepsPermissionsInfoIamPermissionsArray.to_proto(
            resource.iam_permissions
        ):
            res.iam_permissions.extend(
                InstancePreprocessDeleteRecipeStepsPermissionsInfoIamPermissionsArray.to_proto(
                    resource.iam_permissions
                )
            )
        if Primitive.to_proto(resource.resource_path):
            res.resource_path = Primitive.to_proto(resource.resource_path)
        if InstanceGoogleprotobufstruct.to_proto(resource.api_attrs):
            res.api_attrs.CopyFrom(
                InstanceGoogleprotobufstruct.to_proto(resource.api_attrs)
            )
        else:
            res.ClearField("api_attrs")
        if InstancePreprocessDeleteRecipeStepsPermissionsInfoPolicyNameModeEnum.to_proto(
            resource.policy_name_mode
        ):
            res.policy_name_mode = InstancePreprocessDeleteRecipeStepsPermissionsInfoPolicyNameModeEnum.to_proto(
                resource.policy_name_mode
            )
        if InstancePreprocessDeleteRecipeStepsPermissionsInfoResource.to_proto(
            resource.resource
        ):
            res.resource.CopyFrom(
                InstancePreprocessDeleteRecipeStepsPermissionsInfoResource.to_proto(
                    resource.resource
                )
            )
        else:
            res.ClearField("resource")
        return res

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return None

        return InstancePreprocessDeleteRecipeStepsPermissionsInfo(
            policy_name=InstancePreprocessDeleteRecipeStepsPermissionsInfoPolicyName.from_proto(
                resource.policy_name
            ),
            iam_permissions=InstancePreprocessDeleteRecipeStepsPermissionsInfoIamPermissionsArray.from_proto(
                resource.iam_permissions
            ),
            resource_path=Primitive.from_proto(resource.resource_path),
            api_attrs=InstanceGoogleprotobufstruct.from_proto(resource.api_attrs),
            policy_name_mode=InstancePreprocessDeleteRecipeStepsPermissionsInfoPolicyNameModeEnum.from_proto(
                resource.policy_name_mode
            ),
            resource=InstancePreprocessDeleteRecipeStepsPermissionsInfoResource.from_proto(
                resource.resource
            ),
        )


class InstancePreprocessDeleteRecipeStepsPermissionsInfoArray(object):
    @classmethod
    def to_proto(self, resources):
        if not resources:
            return resources
        return [
            InstancePreprocessDeleteRecipeStepsPermissionsInfo.to_proto(i)
            for i in resources
        ]

    @classmethod
    def from_proto(self, resources):
        return [
            InstancePreprocessDeleteRecipeStepsPermissionsInfo.from_proto(i)
            for i in resources
        ]


class InstancePreprocessDeleteRecipeStepsPermissionsInfoPolicyName(object):
    def __init__(self, type: str = None, id: str = None, region: str = None):
        self.type = type
        self.id = id
        self.region = region

    @classmethod
    def to_proto(self, resource):
        if not resource:
            return None

        res = (
            instance_pb2.Tier2AlphaInstancePreprocessDeleteRecipeStepsPermissionsInfoPolicyName()
        )
        if Primitive.to_proto(resource.type):
            res.type = Primitive.to_proto(resource.type)
        if Primitive.to_proto(resource.id):
            res.id = Primitive.to_proto(resource.id)
        if Primitive.to_proto(resource.region):
            res.region = Primitive.to_proto(resource.region)
        return res

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return None

        return InstancePreprocessDeleteRecipeStepsPermissionsInfoPolicyName(
            type=Primitive.from_proto(resource.type),
            id=Primitive.from_proto(resource.id),
            region=Primitive.from_proto(resource.region),
        )


class InstancePreprocessDeleteRecipeStepsPermissionsInfoPolicyNameArray(object):
    @classmethod
    def to_proto(self, resources):
        if not resources:
            return resources
        return [
            InstancePreprocessDeleteRecipeStepsPermissionsInfoPolicyName.to_proto(i)
            for i in resources
        ]

    @classmethod
    def from_proto(self, resources):
        return [
            InstancePreprocessDeleteRecipeStepsPermissionsInfoPolicyName.from_proto(i)
            for i in resources
        ]


class InstancePreprocessDeleteRecipeStepsPermissionsInfoIamPermissions(object):
    def __init__(self, permission: str = None):
        self.permission = permission

    @classmethod
    def to_proto(self, resource):
        if not resource:
            return None

        res = (
            instance_pb2.Tier2AlphaInstancePreprocessDeleteRecipeStepsPermissionsInfoIamPermissions()
        )
        if Primitive.to_proto(resource.permission):
            res.permission = Primitive.to_proto(resource.permission)
        return res

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return None

        return InstancePreprocessDeleteRecipeStepsPermissionsInfoIamPermissions(
            permission=Primitive.from_proto(resource.permission),
        )


class InstancePreprocessDeleteRecipeStepsPermissionsInfoIamPermissionsArray(object):
    @classmethod
    def to_proto(self, resources):
        if not resources:
            return resources
        return [
            InstancePreprocessDeleteRecipeStepsPermissionsInfoIamPermissions.to_proto(i)
            for i in resources
        ]

    @classmethod
    def from_proto(self, resources):
        return [
            InstancePreprocessDeleteRecipeStepsPermissionsInfoIamPermissions.from_proto(
                i
            )
            for i in resources
        ]


class InstancePreprocessDeleteRecipeStepsPermissionsInfoResource(object):
    def __init__(
        self,
        name: str = None,
        type: str = None,
        service: str = None,
        labels: dict = None,
    ):
        self.name = name
        self.type = type
        self.service = service
        self.labels = labels

    @classmethod
    def to_proto(self, resource):
        if not resource:
            return None

        res = (
            instance_pb2.Tier2AlphaInstancePreprocessDeleteRecipeStepsPermissionsInfoResource()
        )
        if Primitive.to_proto(resource.name):
            res.name = Primitive.to_proto(resource.name)
        if Primitive.to_proto(resource.type):
            res.type = Primitive.to_proto(resource.type)
        if Primitive.to_proto(resource.service):
            res.service = Primitive.to_proto(resource.service)
        if Primitive.to_proto(resource.labels):
            res.labels = Primitive.to_proto(resource.labels)
        return res

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return None

        return InstancePreprocessDeleteRecipeStepsPermissionsInfoResource(
            name=Primitive.from_proto(resource.name),
            type=Primitive.from_proto(resource.type),
            service=Primitive.from_proto(resource.service),
            labels=Primitive.from_proto(resource.labels),
        )


class InstancePreprocessDeleteRecipeStepsPermissionsInfoResourceArray(object):
    @classmethod
    def to_proto(self, resources):
        if not resources:
            return resources
        return [
            InstancePreprocessDeleteRecipeStepsPermissionsInfoResource.to_proto(i)
            for i in resources
        ]

    @classmethod
    def from_proto(self, resources):
        return [
            InstancePreprocessDeleteRecipeStepsPermissionsInfoResource.from_proto(i)
            for i in resources
        ]


class InstancePreprocessDeleteRecipeStepsKeyNotificationsUpdate(object):
    def __init__(self, key_notifications_info: dict = None):
        self.key_notifications_info = key_notifications_info

    @classmethod
    def to_proto(self, resource):
        if not resource:
            return None

        res = (
            instance_pb2.Tier2AlphaInstancePreprocessDeleteRecipeStepsKeyNotificationsUpdate()
        )
        if InstancePreprocessDeleteRecipeStepsKeyNotificationsUpdateKeyNotificationsInfo.to_proto(
            resource.key_notifications_info
        ):
            res.key_notifications_info.CopyFrom(
                InstancePreprocessDeleteRecipeStepsKeyNotificationsUpdateKeyNotificationsInfo.to_proto(
                    resource.key_notifications_info
                )
            )
        else:
            res.ClearField("key_notifications_info")
        return res

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return None

        return InstancePreprocessDeleteRecipeStepsKeyNotificationsUpdate(
            key_notifications_info=InstancePreprocessDeleteRecipeStepsKeyNotificationsUpdateKeyNotificationsInfo.from_proto(
                resource.key_notifications_info
            ),
        )


class InstancePreprocessDeleteRecipeStepsKeyNotificationsUpdateArray(object):
    @classmethod
    def to_proto(self, resources):
        if not resources:
            return resources
        return [
            InstancePreprocessDeleteRecipeStepsKeyNotificationsUpdate.to_proto(i)
            for i in resources
        ]

    @classmethod
    def from_proto(self, resources):
        return [
            InstancePreprocessDeleteRecipeStepsKeyNotificationsUpdate.from_proto(i)
            for i in resources
        ]


class InstancePreprocessDeleteRecipeStepsKeyNotificationsUpdateKeyNotificationsInfo(
    object
):
    def __init__(
        self,
        data_version: int = None,
        delegate: str = None,
        key_notification_configs: list = None,
    ):
        self.data_version = data_version
        self.delegate = delegate
        self.key_notification_configs = key_notification_configs

    @classmethod
    def to_proto(self, resource):
        if not resource:
            return None

        res = (
            instance_pb2.Tier2AlphaInstancePreprocessDeleteRecipeStepsKeyNotificationsUpdateKeyNotificationsInfo()
        )
        if Primitive.to_proto(resource.data_version):
            res.data_version = Primitive.to_proto(resource.data_version)
        if Primitive.to_proto(resource.delegate):
            res.delegate = Primitive.to_proto(resource.delegate)
        if InstancePreprocessDeleteRecipeStepsKeyNotificationsUpdateKeyNotificationsInfoKeyNotificationConfigsArray.to_proto(
            resource.key_notification_configs
        ):
            res.key_notification_configs.extend(
                InstancePreprocessDeleteRecipeStepsKeyNotificationsUpdateKeyNotificationsInfoKeyNotificationConfigsArray.to_proto(
                    resource.key_notification_configs
                )
            )
        return res

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return None

        return InstancePreprocessDeleteRecipeStepsKeyNotificationsUpdateKeyNotificationsInfo(
            data_version=Primitive.from_proto(resource.data_version),
            delegate=Primitive.from_proto(resource.delegate),
            key_notification_configs=InstancePreprocessDeleteRecipeStepsKeyNotificationsUpdateKeyNotificationsInfoKeyNotificationConfigsArray.from_proto(
                resource.key_notification_configs
            ),
        )


class InstancePreprocessDeleteRecipeStepsKeyNotificationsUpdateKeyNotificationsInfoArray(
    object
):
    @classmethod
    def to_proto(self, resources):
        if not resources:
            return resources
        return [
            InstancePreprocessDeleteRecipeStepsKeyNotificationsUpdateKeyNotificationsInfo.to_proto(
                i
            )
            for i in resources
        ]

    @classmethod
    def from_proto(self, resources):
        return [
            InstancePreprocessDeleteRecipeStepsKeyNotificationsUpdateKeyNotificationsInfo.from_proto(
                i
            )
            for i in resources
        ]


class InstancePreprocessDeleteRecipeStepsKeyNotificationsUpdateKeyNotificationsInfoKeyNotificationConfigs(
    object
):
    def __init__(
        self,
        key_or_version_name: str = None,
        grant: str = None,
        delegator_gaia_id: int = None,
    ):
        self.key_or_version_name = key_or_version_name
        self.grant = grant
        self.delegator_gaia_id = delegator_gaia_id

    @classmethod
    def to_proto(self, resource):
        if not resource:
            return None

        res = (
            instance_pb2.Tier2AlphaInstancePreprocessDeleteRecipeStepsKeyNotificationsUpdateKeyNotificationsInfoKeyNotificationConfigs()
        )
        if Primitive.to_proto(resource.key_or_version_name):
            res.key_or_version_name = Primitive.to_proto(resource.key_or_version_name)
        if Primitive.to_proto(resource.grant):
            res.grant = Primitive.to_proto(resource.grant)
        if Primitive.to_proto(resource.delegator_gaia_id):
            res.delegator_gaia_id = Primitive.to_proto(resource.delegator_gaia_id)
        return res

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return None

        return InstancePreprocessDeleteRecipeStepsKeyNotificationsUpdateKeyNotificationsInfoKeyNotificationConfigs(
            key_or_version_name=Primitive.from_proto(resource.key_or_version_name),
            grant=Primitive.from_proto(resource.grant),
            delegator_gaia_id=Primitive.from_proto(resource.delegator_gaia_id),
        )


class InstancePreprocessDeleteRecipeStepsKeyNotificationsUpdateKeyNotificationsInfoKeyNotificationConfigsArray(
    object
):
    @classmethod
    def to_proto(self, resources):
        if not resources:
            return resources
        return [
            InstancePreprocessDeleteRecipeStepsKeyNotificationsUpdateKeyNotificationsInfoKeyNotificationConfigs.to_proto(
                i
            )
            for i in resources
        ]

    @classmethod
    def from_proto(self, resources):
        return [
            InstancePreprocessDeleteRecipeStepsKeyNotificationsUpdateKeyNotificationsInfoKeyNotificationConfigs.from_proto(
                i
            )
            for i in resources
        ]


class InstancePreprocessUpdateRecipe(object):
    def __init__(
        self,
        steps: list = None,
        honor_cancel_request: bool = None,
        ignore_recipe_after: int = None,
        verify_deadline_seconds_below: float = None,
        populate_operation_result: bool = None,
        readonly_recipe_start_time: str = None,
        resource_names_stored_in_clh_with_delay: list = None,
        delay_to_store_resources_in_clh_db_nanos: int = None,
    ):
        self.steps = steps
        self.honor_cancel_request = honor_cancel_request
        self.ignore_recipe_after = ignore_recipe_after
        self.verify_deadline_seconds_below = verify_deadline_seconds_below
        self.populate_operation_result = populate_operation_result
        self.readonly_recipe_start_time = readonly_recipe_start_time
        self.resource_names_stored_in_clh_with_delay = (
            resource_names_stored_in_clh_with_delay
        )
        self.delay_to_store_resources_in_clh_db_nanos = (
            delay_to_store_resources_in_clh_db_nanos
        )

    @classmethod
    def to_proto(self, resource):
        if not resource:
            return None

        res = instance_pb2.Tier2AlphaInstancePreprocessUpdateRecipe()
        if InstancePreprocessUpdateRecipeStepsArray.to_proto(resource.steps):
            res.steps.extend(
                InstancePreprocessUpdateRecipeStepsArray.to_proto(resource.steps)
            )
        if Primitive.to_proto(resource.honor_cancel_request):
            res.honor_cancel_request = Primitive.to_proto(resource.honor_cancel_request)
        if Primitive.to_proto(resource.ignore_recipe_after):
            res.ignore_recipe_after = Primitive.to_proto(resource.ignore_recipe_after)
        if Primitive.to_proto(resource.verify_deadline_seconds_below):
            res.verify_deadline_seconds_below = Primitive.to_proto(
                resource.verify_deadline_seconds_below
            )
        if Primitive.to_proto(resource.populate_operation_result):
            res.populate_operation_result = Primitive.to_proto(
                resource.populate_operation_result
            )
        if Primitive.to_proto(resource.readonly_recipe_start_time):
            res.readonly_recipe_start_time = Primitive.to_proto(
                resource.readonly_recipe_start_time
            )
        if Primitive.to_proto(resource.resource_names_stored_in_clh_with_delay):
            res.resource_names_stored_in_clh_with_delay.extend(
                Primitive.to_proto(resource.resource_names_stored_in_clh_with_delay)
            )
        if Primitive.to_proto(resource.delay_to_store_resources_in_clh_db_nanos):
            res.delay_to_store_resources_in_clh_db_nanos = Primitive.to_proto(
                resource.delay_to_store_resources_in_clh_db_nanos
            )
        return res

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return None

        return InstancePreprocessUpdateRecipe(
            steps=InstancePreprocessUpdateRecipeStepsArray.from_proto(resource.steps),
            honor_cancel_request=Primitive.from_proto(resource.honor_cancel_request),
            ignore_recipe_after=Primitive.from_proto(resource.ignore_recipe_after),
            verify_deadline_seconds_below=Primitive.from_proto(
                resource.verify_deadline_seconds_below
            ),
            populate_operation_result=Primitive.from_proto(
                resource.populate_operation_result
            ),
            readonly_recipe_start_time=Primitive.from_proto(
                resource.readonly_recipe_start_time
            ),
            resource_names_stored_in_clh_with_delay=Primitive.from_proto(
                resource.resource_names_stored_in_clh_with_delay
            ),
            delay_to_store_resources_in_clh_db_nanos=Primitive.from_proto(
                resource.delay_to_store_resources_in_clh_db_nanos
            ),
        )


class InstancePreprocessUpdateRecipeArray(object):
    @classmethod
    def to_proto(self, resources):
        if not resources:
            return resources
        return [InstancePreprocessUpdateRecipe.to_proto(i) for i in resources]

    @classmethod
    def from_proto(self, resources):
        return [InstancePreprocessUpdateRecipe.from_proto(i) for i in resources]


class InstancePreprocessUpdateRecipeSteps(object):
    def __init__(
        self,
        relative_time: int = None,
        sleep_duration: int = None,
        action: str = None,
        status: dict = None,
        error_space: str = None,
        p4_service_account: str = None,
        resource_metadata_size: int = None,
        description: str = None,
        updated_repeat_operation_delay_sec: float = None,
        quota_request_deltas: list = None,
        preprocess_update: dict = None,
        public_operation_metadata: str = None,
        requested_tenant_project: dict = None,
        permissions_info: list = None,
        key_notifications_update: dict = None,
        clh_data_update_time: str = None,
        public_error_message: str = None,
    ):
        self.relative_time = relative_time
        self.sleep_duration = sleep_duration
        self.action = action
        self.status = status
        self.error_space = error_space
        self.p4_service_account = p4_service_account
        self.resource_metadata_size = resource_metadata_size
        self.description = description
        self.updated_repeat_operation_delay_sec = updated_repeat_operation_delay_sec
        self.quota_request_deltas = quota_request_deltas
        self.preprocess_update = preprocess_update
        self.public_operation_metadata = public_operation_metadata
        self.requested_tenant_project = requested_tenant_project
        self.permissions_info = permissions_info
        self.key_notifications_update = key_notifications_update
        self.clh_data_update_time = clh_data_update_time
        self.public_error_message = public_error_message

    @classmethod
    def to_proto(self, resource):
        if not resource:
            return None

        res = instance_pb2.Tier2AlphaInstancePreprocessUpdateRecipeSteps()
        if Primitive.to_proto(resource.relative_time):
            res.relative_time = Primitive.to_proto(resource.relative_time)
        if Primitive.to_proto(resource.sleep_duration):
            res.sleep_duration = Primitive.to_proto(resource.sleep_duration)
        if InstancePreprocessUpdateRecipeStepsActionEnum.to_proto(resource.action):
            res.action = InstancePreprocessUpdateRecipeStepsActionEnum.to_proto(
                resource.action
            )
        if InstancePreprocessUpdateRecipeStepsStatus.to_proto(resource.status):
            res.status.CopyFrom(
                InstancePreprocessUpdateRecipeStepsStatus.to_proto(resource.status)
            )
        else:
            res.ClearField("status")
        if Primitive.to_proto(resource.error_space):
            res.error_space = Primitive.to_proto(resource.error_space)
        if Primitive.to_proto(resource.p4_service_account):
            res.p4_service_account = Primitive.to_proto(resource.p4_service_account)
        if Primitive.to_proto(resource.resource_metadata_size):
            res.resource_metadata_size = Primitive.to_proto(
                resource.resource_metadata_size
            )
        if Primitive.to_proto(resource.description):
            res.description = Primitive.to_proto(resource.description)
        if Primitive.to_proto(resource.updated_repeat_operation_delay_sec):
            res.updated_repeat_operation_delay_sec = Primitive.to_proto(
                resource.updated_repeat_operation_delay_sec
            )
        if InstancePreprocessUpdateRecipeStepsQuotaRequestDeltasArray.to_proto(
            resource.quota_request_deltas
        ):
            res.quota_request_deltas.extend(
                InstancePreprocessUpdateRecipeStepsQuotaRequestDeltasArray.to_proto(
                    resource.quota_request_deltas
                )
            )
        if InstancePreprocessUpdateRecipeStepsPreprocessUpdate.to_proto(
            resource.preprocess_update
        ):
            res.preprocess_update.CopyFrom(
                InstancePreprocessUpdateRecipeStepsPreprocessUpdate.to_proto(
                    resource.preprocess_update
                )
            )
        else:
            res.ClearField("preprocess_update")
        if Primitive.to_proto(resource.public_operation_metadata):
            res.public_operation_metadata = Primitive.to_proto(
                resource.public_operation_metadata
            )
        if InstancePreprocessUpdateRecipeStepsRequestedTenantProject.to_proto(
            resource.requested_tenant_project
        ):
            res.requested_tenant_project.CopyFrom(
                InstancePreprocessUpdateRecipeStepsRequestedTenantProject.to_proto(
                    resource.requested_tenant_project
                )
            )
        else:
            res.ClearField("requested_tenant_project")
        if InstancePreprocessUpdateRecipeStepsPermissionsInfoArray.to_proto(
            resource.permissions_info
        ):
            res.permissions_info.extend(
                InstancePreprocessUpdateRecipeStepsPermissionsInfoArray.to_proto(
                    resource.permissions_info
                )
            )
        if InstancePreprocessUpdateRecipeStepsKeyNotificationsUpdate.to_proto(
            resource.key_notifications_update
        ):
            res.key_notifications_update.CopyFrom(
                InstancePreprocessUpdateRecipeStepsKeyNotificationsUpdate.to_proto(
                    resource.key_notifications_update
                )
            )
        else:
            res.ClearField("key_notifications_update")
        if Primitive.to_proto(resource.clh_data_update_time):
            res.clh_data_update_time = Primitive.to_proto(resource.clh_data_update_time)
        if Primitive.to_proto(resource.public_error_message):
            res.public_error_message = Primitive.to_proto(resource.public_error_message)
        return res

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return None

        return InstancePreprocessUpdateRecipeSteps(
            relative_time=Primitive.from_proto(resource.relative_time),
            sleep_duration=Primitive.from_proto(resource.sleep_duration),
            action=InstancePreprocessUpdateRecipeStepsActionEnum.from_proto(
                resource.action
            ),
            status=InstancePreprocessUpdateRecipeStepsStatus.from_proto(
                resource.status
            ),
            error_space=Primitive.from_proto(resource.error_space),
            p4_service_account=Primitive.from_proto(resource.p4_service_account),
            resource_metadata_size=Primitive.from_proto(
                resource.resource_metadata_size
            ),
            description=Primitive.from_proto(resource.description),
            updated_repeat_operation_delay_sec=Primitive.from_proto(
                resource.updated_repeat_operation_delay_sec
            ),
            quota_request_deltas=InstancePreprocessUpdateRecipeStepsQuotaRequestDeltasArray.from_proto(
                resource.quota_request_deltas
            ),
            preprocess_update=InstancePreprocessUpdateRecipeStepsPreprocessUpdate.from_proto(
                resource.preprocess_update
            ),
            public_operation_metadata=Primitive.from_proto(
                resource.public_operation_metadata
            ),
            requested_tenant_project=InstancePreprocessUpdateRecipeStepsRequestedTenantProject.from_proto(
                resource.requested_tenant_project
            ),
            permissions_info=InstancePreprocessUpdateRecipeStepsPermissionsInfoArray.from_proto(
                resource.permissions_info
            ),
            key_notifications_update=InstancePreprocessUpdateRecipeStepsKeyNotificationsUpdate.from_proto(
                resource.key_notifications_update
            ),
            clh_data_update_time=Primitive.from_proto(resource.clh_data_update_time),
            public_error_message=Primitive.from_proto(resource.public_error_message),
        )


class InstancePreprocessUpdateRecipeStepsArray(object):
    @classmethod
    def to_proto(self, resources):
        if not resources:
            return resources
        return [InstancePreprocessUpdateRecipeSteps.to_proto(i) for i in resources]

    @classmethod
    def from_proto(self, resources):
        return [InstancePreprocessUpdateRecipeSteps.from_proto(i) for i in resources]


class InstancePreprocessUpdateRecipeStepsStatus(object):
    def __init__(self, code: int = None, message: str = None, details: list = None):
        self.code = code
        self.message = message
        self.details = details

    @classmethod
    def to_proto(self, resource):
        if not resource:
            return None

        res = instance_pb2.Tier2AlphaInstancePreprocessUpdateRecipeStepsStatus()
        if Primitive.to_proto(resource.code):
            res.code = Primitive.to_proto(resource.code)
        if Primitive.to_proto(resource.message):
            res.message = Primitive.to_proto(resource.message)
        if InstancePreprocessUpdateRecipeStepsStatusDetailsArray.to_proto(
            resource.details
        ):
            res.details.extend(
                InstancePreprocessUpdateRecipeStepsStatusDetailsArray.to_proto(
                    resource.details
                )
            )
        return res

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return None

        return InstancePreprocessUpdateRecipeStepsStatus(
            code=Primitive.from_proto(resource.code),
            message=Primitive.from_proto(resource.message),
            details=InstancePreprocessUpdateRecipeStepsStatusDetailsArray.from_proto(
                resource.details
            ),
        )


class InstancePreprocessUpdateRecipeStepsStatusArray(object):
    @classmethod
    def to_proto(self, resources):
        if not resources:
            return resources
        return [
            InstancePreprocessUpdateRecipeStepsStatus.to_proto(i) for i in resources
        ]

    @classmethod
    def from_proto(self, resources):
        return [
            InstancePreprocessUpdateRecipeStepsStatus.from_proto(i) for i in resources
        ]


class InstancePreprocessUpdateRecipeStepsStatusDetails(object):
    def __init__(self, type_url: str = None, value: str = None):
        self.type_url = type_url
        self.value = value

    @classmethod
    def to_proto(self, resource):
        if not resource:
            return None

        res = instance_pb2.Tier2AlphaInstancePreprocessUpdateRecipeStepsStatusDetails()
        if Primitive.to_proto(resource.type_url):
            res.type_url = Primitive.to_proto(resource.type_url)
        if Primitive.to_proto(resource.value):
            res.value = Primitive.to_proto(resource.value)
        return res

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return None

        return InstancePreprocessUpdateRecipeStepsStatusDetails(
            type_url=Primitive.from_proto(resource.type_url),
            value=Primitive.from_proto(resource.value),
        )


class InstancePreprocessUpdateRecipeStepsStatusDetailsArray(object):
    @classmethod
    def to_proto(self, resources):
        if not resources:
            return resources
        return [
            InstancePreprocessUpdateRecipeStepsStatusDetails.to_proto(i)
            for i in resources
        ]

    @classmethod
    def from_proto(self, resources):
        return [
            InstancePreprocessUpdateRecipeStepsStatusDetails.from_proto(i)
            for i in resources
        ]


class InstancePreprocessUpdateRecipeStepsQuotaRequestDeltas(object):
    def __init__(
        self,
        metric_name: str = None,
        amount: int = None,
        quota_location_name: str = None,
    ):
        self.metric_name = metric_name
        self.amount = amount
        self.quota_location_name = quota_location_name

    @classmethod
    def to_proto(self, resource):
        if not resource:
            return None

        res = (
            instance_pb2.Tier2AlphaInstancePreprocessUpdateRecipeStepsQuotaRequestDeltas()
        )
        if Primitive.to_proto(resource.metric_name):
            res.metric_name = Primitive.to_proto(resource.metric_name)
        if Primitive.to_proto(resource.amount):
            res.amount = Primitive.to_proto(resource.amount)
        if Primitive.to_proto(resource.quota_location_name):
            res.quota_location_name = Primitive.to_proto(resource.quota_location_name)
        return res

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return None

        return InstancePreprocessUpdateRecipeStepsQuotaRequestDeltas(
            metric_name=Primitive.from_proto(resource.metric_name),
            amount=Primitive.from_proto(resource.amount),
            quota_location_name=Primitive.from_proto(resource.quota_location_name),
        )


class InstancePreprocessUpdateRecipeStepsQuotaRequestDeltasArray(object):
    @classmethod
    def to_proto(self, resources):
        if not resources:
            return resources
        return [
            InstancePreprocessUpdateRecipeStepsQuotaRequestDeltas.to_proto(i)
            for i in resources
        ]

    @classmethod
    def from_proto(self, resources):
        return [
            InstancePreprocessUpdateRecipeStepsQuotaRequestDeltas.from_proto(i)
            for i in resources
        ]


class InstancePreprocessUpdateRecipeStepsPreprocessUpdate(object):
    def __init__(
        self, latency_slo_bucket_name: str = None, public_operation_metadata: str = None
    ):
        self.latency_slo_bucket_name = latency_slo_bucket_name
        self.public_operation_metadata = public_operation_metadata

    @classmethod
    def to_proto(self, resource):
        if not resource:
            return None

        res = (
            instance_pb2.Tier2AlphaInstancePreprocessUpdateRecipeStepsPreprocessUpdate()
        )
        if Primitive.to_proto(resource.latency_slo_bucket_name):
            res.latency_slo_bucket_name = Primitive.to_proto(
                resource.latency_slo_bucket_name
            )
        if Primitive.to_proto(resource.public_operation_metadata):
            res.public_operation_metadata = Primitive.to_proto(
                resource.public_operation_metadata
            )
        return res

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return None

        return InstancePreprocessUpdateRecipeStepsPreprocessUpdate(
            latency_slo_bucket_name=Primitive.from_proto(
                resource.latency_slo_bucket_name
            ),
            public_operation_metadata=Primitive.from_proto(
                resource.public_operation_metadata
            ),
        )


class InstancePreprocessUpdateRecipeStepsPreprocessUpdateArray(object):
    @classmethod
    def to_proto(self, resources):
        if not resources:
            return resources
        return [
            InstancePreprocessUpdateRecipeStepsPreprocessUpdate.to_proto(i)
            for i in resources
        ]

    @classmethod
    def from_proto(self, resources):
        return [
            InstancePreprocessUpdateRecipeStepsPreprocessUpdate.from_proto(i)
            for i in resources
        ]


class InstancePreprocessUpdateRecipeStepsRequestedTenantProject(object):
    def __init__(self, tag: str = None, folder: str = None, scope: str = None):
        self.tag = tag
        self.folder = folder
        self.scope = scope

    @classmethod
    def to_proto(self, resource):
        if not resource:
            return None

        res = (
            instance_pb2.Tier2AlphaInstancePreprocessUpdateRecipeStepsRequestedTenantProject()
        )
        if Primitive.to_proto(resource.tag):
            res.tag = Primitive.to_proto(resource.tag)
        if Primitive.to_proto(resource.folder):
            res.folder = Primitive.to_proto(resource.folder)
        if InstancePreprocessUpdateRecipeStepsRequestedTenantProjectScopeEnum.to_proto(
            resource.scope
        ):
            res.scope = InstancePreprocessUpdateRecipeStepsRequestedTenantProjectScopeEnum.to_proto(
                resource.scope
            )
        return res

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return None

        return InstancePreprocessUpdateRecipeStepsRequestedTenantProject(
            tag=Primitive.from_proto(resource.tag),
            folder=Primitive.from_proto(resource.folder),
            scope=InstancePreprocessUpdateRecipeStepsRequestedTenantProjectScopeEnum.from_proto(
                resource.scope
            ),
        )


class InstancePreprocessUpdateRecipeStepsRequestedTenantProjectArray(object):
    @classmethod
    def to_proto(self, resources):
        if not resources:
            return resources
        return [
            InstancePreprocessUpdateRecipeStepsRequestedTenantProject.to_proto(i)
            for i in resources
        ]

    @classmethod
    def from_proto(self, resources):
        return [
            InstancePreprocessUpdateRecipeStepsRequestedTenantProject.from_proto(i)
            for i in resources
        ]


class InstancePreprocessUpdateRecipeStepsPermissionsInfo(object):
    def __init__(
        self,
        policy_name: dict = None,
        iam_permissions: list = None,
        resource_path: str = None,
        api_attrs: dict = None,
        policy_name_mode: str = None,
        resource: dict = None,
    ):
        self.policy_name = policy_name
        self.iam_permissions = iam_permissions
        self.resource_path = resource_path
        self.api_attrs = api_attrs
        self.policy_name_mode = policy_name_mode
        self.resource = resource

    @classmethod
    def to_proto(self, resource):
        if not resource:
            return None

        res = (
            instance_pb2.Tier2AlphaInstancePreprocessUpdateRecipeStepsPermissionsInfo()
        )
        if InstancePreprocessUpdateRecipeStepsPermissionsInfoPolicyName.to_proto(
            resource.policy_name
        ):
            res.policy_name.CopyFrom(
                InstancePreprocessUpdateRecipeStepsPermissionsInfoPolicyName.to_proto(
                    resource.policy_name
                )
            )
        else:
            res.ClearField("policy_name")
        if InstancePreprocessUpdateRecipeStepsPermissionsInfoIamPermissionsArray.to_proto(
            resource.iam_permissions
        ):
            res.iam_permissions.extend(
                InstancePreprocessUpdateRecipeStepsPermissionsInfoIamPermissionsArray.to_proto(
                    resource.iam_permissions
                )
            )
        if Primitive.to_proto(resource.resource_path):
            res.resource_path = Primitive.to_proto(resource.resource_path)
        if InstanceGoogleprotobufstruct.to_proto(resource.api_attrs):
            res.api_attrs.CopyFrom(
                InstanceGoogleprotobufstruct.to_proto(resource.api_attrs)
            )
        else:
            res.ClearField("api_attrs")
        if InstancePreprocessUpdateRecipeStepsPermissionsInfoPolicyNameModeEnum.to_proto(
            resource.policy_name_mode
        ):
            res.policy_name_mode = InstancePreprocessUpdateRecipeStepsPermissionsInfoPolicyNameModeEnum.to_proto(
                resource.policy_name_mode
            )
        if InstancePreprocessUpdateRecipeStepsPermissionsInfoResource.to_proto(
            resource.resource
        ):
            res.resource.CopyFrom(
                InstancePreprocessUpdateRecipeStepsPermissionsInfoResource.to_proto(
                    resource.resource
                )
            )
        else:
            res.ClearField("resource")
        return res

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return None

        return InstancePreprocessUpdateRecipeStepsPermissionsInfo(
            policy_name=InstancePreprocessUpdateRecipeStepsPermissionsInfoPolicyName.from_proto(
                resource.policy_name
            ),
            iam_permissions=InstancePreprocessUpdateRecipeStepsPermissionsInfoIamPermissionsArray.from_proto(
                resource.iam_permissions
            ),
            resource_path=Primitive.from_proto(resource.resource_path),
            api_attrs=InstanceGoogleprotobufstruct.from_proto(resource.api_attrs),
            policy_name_mode=InstancePreprocessUpdateRecipeStepsPermissionsInfoPolicyNameModeEnum.from_proto(
                resource.policy_name_mode
            ),
            resource=InstancePreprocessUpdateRecipeStepsPermissionsInfoResource.from_proto(
                resource.resource
            ),
        )


class InstancePreprocessUpdateRecipeStepsPermissionsInfoArray(object):
    @classmethod
    def to_proto(self, resources):
        if not resources:
            return resources
        return [
            InstancePreprocessUpdateRecipeStepsPermissionsInfo.to_proto(i)
            for i in resources
        ]

    @classmethod
    def from_proto(self, resources):
        return [
            InstancePreprocessUpdateRecipeStepsPermissionsInfo.from_proto(i)
            for i in resources
        ]


class InstancePreprocessUpdateRecipeStepsPermissionsInfoPolicyName(object):
    def __init__(self, type: str = None, id: str = None, region: str = None):
        self.type = type
        self.id = id
        self.region = region

    @classmethod
    def to_proto(self, resource):
        if not resource:
            return None

        res = (
            instance_pb2.Tier2AlphaInstancePreprocessUpdateRecipeStepsPermissionsInfoPolicyName()
        )
        if Primitive.to_proto(resource.type):
            res.type = Primitive.to_proto(resource.type)
        if Primitive.to_proto(resource.id):
            res.id = Primitive.to_proto(resource.id)
        if Primitive.to_proto(resource.region):
            res.region = Primitive.to_proto(resource.region)
        return res

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return None

        return InstancePreprocessUpdateRecipeStepsPermissionsInfoPolicyName(
            type=Primitive.from_proto(resource.type),
            id=Primitive.from_proto(resource.id),
            region=Primitive.from_proto(resource.region),
        )


class InstancePreprocessUpdateRecipeStepsPermissionsInfoPolicyNameArray(object):
    @classmethod
    def to_proto(self, resources):
        if not resources:
            return resources
        return [
            InstancePreprocessUpdateRecipeStepsPermissionsInfoPolicyName.to_proto(i)
            for i in resources
        ]

    @classmethod
    def from_proto(self, resources):
        return [
            InstancePreprocessUpdateRecipeStepsPermissionsInfoPolicyName.from_proto(i)
            for i in resources
        ]


class InstancePreprocessUpdateRecipeStepsPermissionsInfoIamPermissions(object):
    def __init__(self, permission: str = None):
        self.permission = permission

    @classmethod
    def to_proto(self, resource):
        if not resource:
            return None

        res = (
            instance_pb2.Tier2AlphaInstancePreprocessUpdateRecipeStepsPermissionsInfoIamPermissions()
        )
        if Primitive.to_proto(resource.permission):
            res.permission = Primitive.to_proto(resource.permission)
        return res

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return None

        return InstancePreprocessUpdateRecipeStepsPermissionsInfoIamPermissions(
            permission=Primitive.from_proto(resource.permission),
        )


class InstancePreprocessUpdateRecipeStepsPermissionsInfoIamPermissionsArray(object):
    @classmethod
    def to_proto(self, resources):
        if not resources:
            return resources
        return [
            InstancePreprocessUpdateRecipeStepsPermissionsInfoIamPermissions.to_proto(i)
            for i in resources
        ]

    @classmethod
    def from_proto(self, resources):
        return [
            InstancePreprocessUpdateRecipeStepsPermissionsInfoIamPermissions.from_proto(
                i
            )
            for i in resources
        ]


class InstancePreprocessUpdateRecipeStepsPermissionsInfoResource(object):
    def __init__(
        self,
        name: str = None,
        type: str = None,
        service: str = None,
        labels: dict = None,
    ):
        self.name = name
        self.type = type
        self.service = service
        self.labels = labels

    @classmethod
    def to_proto(self, resource):
        if not resource:
            return None

        res = (
            instance_pb2.Tier2AlphaInstancePreprocessUpdateRecipeStepsPermissionsInfoResource()
        )
        if Primitive.to_proto(resource.name):
            res.name = Primitive.to_proto(resource.name)
        if Primitive.to_proto(resource.type):
            res.type = Primitive.to_proto(resource.type)
        if Primitive.to_proto(resource.service):
            res.service = Primitive.to_proto(resource.service)
        if Primitive.to_proto(resource.labels):
            res.labels = Primitive.to_proto(resource.labels)
        return res

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return None

        return InstancePreprocessUpdateRecipeStepsPermissionsInfoResource(
            name=Primitive.from_proto(resource.name),
            type=Primitive.from_proto(resource.type),
            service=Primitive.from_proto(resource.service),
            labels=Primitive.from_proto(resource.labels),
        )


class InstancePreprocessUpdateRecipeStepsPermissionsInfoResourceArray(object):
    @classmethod
    def to_proto(self, resources):
        if not resources:
            return resources
        return [
            InstancePreprocessUpdateRecipeStepsPermissionsInfoResource.to_proto(i)
            for i in resources
        ]

    @classmethod
    def from_proto(self, resources):
        return [
            InstancePreprocessUpdateRecipeStepsPermissionsInfoResource.from_proto(i)
            for i in resources
        ]


class InstancePreprocessUpdateRecipeStepsKeyNotificationsUpdate(object):
    def __init__(self, key_notifications_info: dict = None):
        self.key_notifications_info = key_notifications_info

    @classmethod
    def to_proto(self, resource):
        if not resource:
            return None

        res = (
            instance_pb2.Tier2AlphaInstancePreprocessUpdateRecipeStepsKeyNotificationsUpdate()
        )
        if InstancePreprocessUpdateRecipeStepsKeyNotificationsUpdateKeyNotificationsInfo.to_proto(
            resource.key_notifications_info
        ):
            res.key_notifications_info.CopyFrom(
                InstancePreprocessUpdateRecipeStepsKeyNotificationsUpdateKeyNotificationsInfo.to_proto(
                    resource.key_notifications_info
                )
            )
        else:
            res.ClearField("key_notifications_info")
        return res

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return None

        return InstancePreprocessUpdateRecipeStepsKeyNotificationsUpdate(
            key_notifications_info=InstancePreprocessUpdateRecipeStepsKeyNotificationsUpdateKeyNotificationsInfo.from_proto(
                resource.key_notifications_info
            ),
        )


class InstancePreprocessUpdateRecipeStepsKeyNotificationsUpdateArray(object):
    @classmethod
    def to_proto(self, resources):
        if not resources:
            return resources
        return [
            InstancePreprocessUpdateRecipeStepsKeyNotificationsUpdate.to_proto(i)
            for i in resources
        ]

    @classmethod
    def from_proto(self, resources):
        return [
            InstancePreprocessUpdateRecipeStepsKeyNotificationsUpdate.from_proto(i)
            for i in resources
        ]


class InstancePreprocessUpdateRecipeStepsKeyNotificationsUpdateKeyNotificationsInfo(
    object
):
    def __init__(
        self,
        data_version: int = None,
        delegate: str = None,
        key_notification_configs: list = None,
    ):
        self.data_version = data_version
        self.delegate = delegate
        self.key_notification_configs = key_notification_configs

    @classmethod
    def to_proto(self, resource):
        if not resource:
            return None

        res = (
            instance_pb2.Tier2AlphaInstancePreprocessUpdateRecipeStepsKeyNotificationsUpdateKeyNotificationsInfo()
        )
        if Primitive.to_proto(resource.data_version):
            res.data_version = Primitive.to_proto(resource.data_version)
        if Primitive.to_proto(resource.delegate):
            res.delegate = Primitive.to_proto(resource.delegate)
        if InstancePreprocessUpdateRecipeStepsKeyNotificationsUpdateKeyNotificationsInfoKeyNotificationConfigsArray.to_proto(
            resource.key_notification_configs
        ):
            res.key_notification_configs.extend(
                InstancePreprocessUpdateRecipeStepsKeyNotificationsUpdateKeyNotificationsInfoKeyNotificationConfigsArray.to_proto(
                    resource.key_notification_configs
                )
            )
        return res

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return None

        return InstancePreprocessUpdateRecipeStepsKeyNotificationsUpdateKeyNotificationsInfo(
            data_version=Primitive.from_proto(resource.data_version),
            delegate=Primitive.from_proto(resource.delegate),
            key_notification_configs=InstancePreprocessUpdateRecipeStepsKeyNotificationsUpdateKeyNotificationsInfoKeyNotificationConfigsArray.from_proto(
                resource.key_notification_configs
            ),
        )


class InstancePreprocessUpdateRecipeStepsKeyNotificationsUpdateKeyNotificationsInfoArray(
    object
):
    @classmethod
    def to_proto(self, resources):
        if not resources:
            return resources
        return [
            InstancePreprocessUpdateRecipeStepsKeyNotificationsUpdateKeyNotificationsInfo.to_proto(
                i
            )
            for i in resources
        ]

    @classmethod
    def from_proto(self, resources):
        return [
            InstancePreprocessUpdateRecipeStepsKeyNotificationsUpdateKeyNotificationsInfo.from_proto(
                i
            )
            for i in resources
        ]


class InstancePreprocessUpdateRecipeStepsKeyNotificationsUpdateKeyNotificationsInfoKeyNotificationConfigs(
    object
):
    def __init__(
        self,
        key_or_version_name: str = None,
        grant: str = None,
        delegator_gaia_id: int = None,
    ):
        self.key_or_version_name = key_or_version_name
        self.grant = grant
        self.delegator_gaia_id = delegator_gaia_id

    @classmethod
    def to_proto(self, resource):
        if not resource:
            return None

        res = (
            instance_pb2.Tier2AlphaInstancePreprocessUpdateRecipeStepsKeyNotificationsUpdateKeyNotificationsInfoKeyNotificationConfigs()
        )
        if Primitive.to_proto(resource.key_or_version_name):
            res.key_or_version_name = Primitive.to_proto(resource.key_or_version_name)
        if Primitive.to_proto(resource.grant):
            res.grant = Primitive.to_proto(resource.grant)
        if Primitive.to_proto(resource.delegator_gaia_id):
            res.delegator_gaia_id = Primitive.to_proto(resource.delegator_gaia_id)
        return res

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return None

        return InstancePreprocessUpdateRecipeStepsKeyNotificationsUpdateKeyNotificationsInfoKeyNotificationConfigs(
            key_or_version_name=Primitive.from_proto(resource.key_or_version_name),
            grant=Primitive.from_proto(resource.grant),
            delegator_gaia_id=Primitive.from_proto(resource.delegator_gaia_id),
        )


class InstancePreprocessUpdateRecipeStepsKeyNotificationsUpdateKeyNotificationsInfoKeyNotificationConfigsArray(
    object
):
    @classmethod
    def to_proto(self, resources):
        if not resources:
            return resources
        return [
            InstancePreprocessUpdateRecipeStepsKeyNotificationsUpdateKeyNotificationsInfoKeyNotificationConfigs.to_proto(
                i
            )
            for i in resources
        ]

    @classmethod
    def from_proto(self, resources):
        return [
            InstancePreprocessUpdateRecipeStepsKeyNotificationsUpdateKeyNotificationsInfoKeyNotificationConfigs.from_proto(
                i
            )
            for i in resources
        ]


class InstancePreprocessFreezeRecipe(object):
    def __init__(
        self,
        steps: list = None,
        honor_cancel_request: bool = None,
        ignore_recipe_after: int = None,
        verify_deadline_seconds_below: float = None,
        populate_operation_result: bool = None,
        readonly_recipe_start_time: str = None,
        resource_names_stored_in_clh_with_delay: list = None,
        delay_to_store_resources_in_clh_db_nanos: int = None,
    ):
        self.steps = steps
        self.honor_cancel_request = honor_cancel_request
        self.ignore_recipe_after = ignore_recipe_after
        self.verify_deadline_seconds_below = verify_deadline_seconds_below
        self.populate_operation_result = populate_operation_result
        self.readonly_recipe_start_time = readonly_recipe_start_time
        self.resource_names_stored_in_clh_with_delay = (
            resource_names_stored_in_clh_with_delay
        )
        self.delay_to_store_resources_in_clh_db_nanos = (
            delay_to_store_resources_in_clh_db_nanos
        )

    @classmethod
    def to_proto(self, resource):
        if not resource:
            return None

        res = instance_pb2.Tier2AlphaInstancePreprocessFreezeRecipe()
        if InstancePreprocessFreezeRecipeStepsArray.to_proto(resource.steps):
            res.steps.extend(
                InstancePreprocessFreezeRecipeStepsArray.to_proto(resource.steps)
            )
        if Primitive.to_proto(resource.honor_cancel_request):
            res.honor_cancel_request = Primitive.to_proto(resource.honor_cancel_request)
        if Primitive.to_proto(resource.ignore_recipe_after):
            res.ignore_recipe_after = Primitive.to_proto(resource.ignore_recipe_after)
        if Primitive.to_proto(resource.verify_deadline_seconds_below):
            res.verify_deadline_seconds_below = Primitive.to_proto(
                resource.verify_deadline_seconds_below
            )
        if Primitive.to_proto(resource.populate_operation_result):
            res.populate_operation_result = Primitive.to_proto(
                resource.populate_operation_result
            )
        if Primitive.to_proto(resource.readonly_recipe_start_time):
            res.readonly_recipe_start_time = Primitive.to_proto(
                resource.readonly_recipe_start_time
            )
        if Primitive.to_proto(resource.resource_names_stored_in_clh_with_delay):
            res.resource_names_stored_in_clh_with_delay.extend(
                Primitive.to_proto(resource.resource_names_stored_in_clh_with_delay)
            )
        if Primitive.to_proto(resource.delay_to_store_resources_in_clh_db_nanos):
            res.delay_to_store_resources_in_clh_db_nanos = Primitive.to_proto(
                resource.delay_to_store_resources_in_clh_db_nanos
            )
        return res

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return None

        return InstancePreprocessFreezeRecipe(
            steps=InstancePreprocessFreezeRecipeStepsArray.from_proto(resource.steps),
            honor_cancel_request=Primitive.from_proto(resource.honor_cancel_request),
            ignore_recipe_after=Primitive.from_proto(resource.ignore_recipe_after),
            verify_deadline_seconds_below=Primitive.from_proto(
                resource.verify_deadline_seconds_below
            ),
            populate_operation_result=Primitive.from_proto(
                resource.populate_operation_result
            ),
            readonly_recipe_start_time=Primitive.from_proto(
                resource.readonly_recipe_start_time
            ),
            resource_names_stored_in_clh_with_delay=Primitive.from_proto(
                resource.resource_names_stored_in_clh_with_delay
            ),
            delay_to_store_resources_in_clh_db_nanos=Primitive.from_proto(
                resource.delay_to_store_resources_in_clh_db_nanos
            ),
        )


class InstancePreprocessFreezeRecipeArray(object):
    @classmethod
    def to_proto(self, resources):
        if not resources:
            return resources
        return [InstancePreprocessFreezeRecipe.to_proto(i) for i in resources]

    @classmethod
    def from_proto(self, resources):
        return [InstancePreprocessFreezeRecipe.from_proto(i) for i in resources]


class InstancePreprocessFreezeRecipeSteps(object):
    def __init__(
        self,
        relative_time: int = None,
        sleep_duration: int = None,
        action: str = None,
        status: dict = None,
        error_space: str = None,
        p4_service_account: str = None,
        resource_metadata_size: int = None,
        description: str = None,
        updated_repeat_operation_delay_sec: float = None,
        quota_request_deltas: list = None,
        preprocess_update: dict = None,
        public_operation_metadata: str = None,
        requested_tenant_project: dict = None,
        permissions_info: list = None,
        key_notifications_update: dict = None,
        clh_data_update_time: str = None,
        public_error_message: str = None,
    ):
        self.relative_time = relative_time
        self.sleep_duration = sleep_duration
        self.action = action
        self.status = status
        self.error_space = error_space
        self.p4_service_account = p4_service_account
        self.resource_metadata_size = resource_metadata_size
        self.description = description
        self.updated_repeat_operation_delay_sec = updated_repeat_operation_delay_sec
        self.quota_request_deltas = quota_request_deltas
        self.preprocess_update = preprocess_update
        self.public_operation_metadata = public_operation_metadata
        self.requested_tenant_project = requested_tenant_project
        self.permissions_info = permissions_info
        self.key_notifications_update = key_notifications_update
        self.clh_data_update_time = clh_data_update_time
        self.public_error_message = public_error_message

    @classmethod
    def to_proto(self, resource):
        if not resource:
            return None

        res = instance_pb2.Tier2AlphaInstancePreprocessFreezeRecipeSteps()
        if Primitive.to_proto(resource.relative_time):
            res.relative_time = Primitive.to_proto(resource.relative_time)
        if Primitive.to_proto(resource.sleep_duration):
            res.sleep_duration = Primitive.to_proto(resource.sleep_duration)
        if InstancePreprocessFreezeRecipeStepsActionEnum.to_proto(resource.action):
            res.action = InstancePreprocessFreezeRecipeStepsActionEnum.to_proto(
                resource.action
            )
        if InstancePreprocessFreezeRecipeStepsStatus.to_proto(resource.status):
            res.status.CopyFrom(
                InstancePreprocessFreezeRecipeStepsStatus.to_proto(resource.status)
            )
        else:
            res.ClearField("status")
        if Primitive.to_proto(resource.error_space):
            res.error_space = Primitive.to_proto(resource.error_space)
        if Primitive.to_proto(resource.p4_service_account):
            res.p4_service_account = Primitive.to_proto(resource.p4_service_account)
        if Primitive.to_proto(resource.resource_metadata_size):
            res.resource_metadata_size = Primitive.to_proto(
                resource.resource_metadata_size
            )
        if Primitive.to_proto(resource.description):
            res.description = Primitive.to_proto(resource.description)
        if Primitive.to_proto(resource.updated_repeat_operation_delay_sec):
            res.updated_repeat_operation_delay_sec = Primitive.to_proto(
                resource.updated_repeat_operation_delay_sec
            )
        if InstancePreprocessFreezeRecipeStepsQuotaRequestDeltasArray.to_proto(
            resource.quota_request_deltas
        ):
            res.quota_request_deltas.extend(
                InstancePreprocessFreezeRecipeStepsQuotaRequestDeltasArray.to_proto(
                    resource.quota_request_deltas
                )
            )
        if InstancePreprocessFreezeRecipeStepsPreprocessUpdate.to_proto(
            resource.preprocess_update
        ):
            res.preprocess_update.CopyFrom(
                InstancePreprocessFreezeRecipeStepsPreprocessUpdate.to_proto(
                    resource.preprocess_update
                )
            )
        else:
            res.ClearField("preprocess_update")
        if Primitive.to_proto(resource.public_operation_metadata):
            res.public_operation_metadata = Primitive.to_proto(
                resource.public_operation_metadata
            )
        if InstancePreprocessFreezeRecipeStepsRequestedTenantProject.to_proto(
            resource.requested_tenant_project
        ):
            res.requested_tenant_project.CopyFrom(
                InstancePreprocessFreezeRecipeStepsRequestedTenantProject.to_proto(
                    resource.requested_tenant_project
                )
            )
        else:
            res.ClearField("requested_tenant_project")
        if InstancePreprocessFreezeRecipeStepsPermissionsInfoArray.to_proto(
            resource.permissions_info
        ):
            res.permissions_info.extend(
                InstancePreprocessFreezeRecipeStepsPermissionsInfoArray.to_proto(
                    resource.permissions_info
                )
            )
        if InstancePreprocessFreezeRecipeStepsKeyNotificationsUpdate.to_proto(
            resource.key_notifications_update
        ):
            res.key_notifications_update.CopyFrom(
                InstancePreprocessFreezeRecipeStepsKeyNotificationsUpdate.to_proto(
                    resource.key_notifications_update
                )
            )
        else:
            res.ClearField("key_notifications_update")
        if Primitive.to_proto(resource.clh_data_update_time):
            res.clh_data_update_time = Primitive.to_proto(resource.clh_data_update_time)
        if Primitive.to_proto(resource.public_error_message):
            res.public_error_message = Primitive.to_proto(resource.public_error_message)
        return res

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return None

        return InstancePreprocessFreezeRecipeSteps(
            relative_time=Primitive.from_proto(resource.relative_time),
            sleep_duration=Primitive.from_proto(resource.sleep_duration),
            action=InstancePreprocessFreezeRecipeStepsActionEnum.from_proto(
                resource.action
            ),
            status=InstancePreprocessFreezeRecipeStepsStatus.from_proto(
                resource.status
            ),
            error_space=Primitive.from_proto(resource.error_space),
            p4_service_account=Primitive.from_proto(resource.p4_service_account),
            resource_metadata_size=Primitive.from_proto(
                resource.resource_metadata_size
            ),
            description=Primitive.from_proto(resource.description),
            updated_repeat_operation_delay_sec=Primitive.from_proto(
                resource.updated_repeat_operation_delay_sec
            ),
            quota_request_deltas=InstancePreprocessFreezeRecipeStepsQuotaRequestDeltasArray.from_proto(
                resource.quota_request_deltas
            ),
            preprocess_update=InstancePreprocessFreezeRecipeStepsPreprocessUpdate.from_proto(
                resource.preprocess_update
            ),
            public_operation_metadata=Primitive.from_proto(
                resource.public_operation_metadata
            ),
            requested_tenant_project=InstancePreprocessFreezeRecipeStepsRequestedTenantProject.from_proto(
                resource.requested_tenant_project
            ),
            permissions_info=InstancePreprocessFreezeRecipeStepsPermissionsInfoArray.from_proto(
                resource.permissions_info
            ),
            key_notifications_update=InstancePreprocessFreezeRecipeStepsKeyNotificationsUpdate.from_proto(
                resource.key_notifications_update
            ),
            clh_data_update_time=Primitive.from_proto(resource.clh_data_update_time),
            public_error_message=Primitive.from_proto(resource.public_error_message),
        )


class InstancePreprocessFreezeRecipeStepsArray(object):
    @classmethod
    def to_proto(self, resources):
        if not resources:
            return resources
        return [InstancePreprocessFreezeRecipeSteps.to_proto(i) for i in resources]

    @classmethod
    def from_proto(self, resources):
        return [InstancePreprocessFreezeRecipeSteps.from_proto(i) for i in resources]


class InstancePreprocessFreezeRecipeStepsStatus(object):
    def __init__(self, code: int = None, message: str = None, details: list = None):
        self.code = code
        self.message = message
        self.details = details

    @classmethod
    def to_proto(self, resource):
        if not resource:
            return None

        res = instance_pb2.Tier2AlphaInstancePreprocessFreezeRecipeStepsStatus()
        if Primitive.to_proto(resource.code):
            res.code = Primitive.to_proto(resource.code)
        if Primitive.to_proto(resource.message):
            res.message = Primitive.to_proto(resource.message)
        if InstancePreprocessFreezeRecipeStepsStatusDetailsArray.to_proto(
            resource.details
        ):
            res.details.extend(
                InstancePreprocessFreezeRecipeStepsStatusDetailsArray.to_proto(
                    resource.details
                )
            )
        return res

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return None

        return InstancePreprocessFreezeRecipeStepsStatus(
            code=Primitive.from_proto(resource.code),
            message=Primitive.from_proto(resource.message),
            details=InstancePreprocessFreezeRecipeStepsStatusDetailsArray.from_proto(
                resource.details
            ),
        )


class InstancePreprocessFreezeRecipeStepsStatusArray(object):
    @classmethod
    def to_proto(self, resources):
        if not resources:
            return resources
        return [
            InstancePreprocessFreezeRecipeStepsStatus.to_proto(i) for i in resources
        ]

    @classmethod
    def from_proto(self, resources):
        return [
            InstancePreprocessFreezeRecipeStepsStatus.from_proto(i) for i in resources
        ]


class InstancePreprocessFreezeRecipeStepsStatusDetails(object):
    def __init__(self, type_url: str = None, value: str = None):
        self.type_url = type_url
        self.value = value

    @classmethod
    def to_proto(self, resource):
        if not resource:
            return None

        res = instance_pb2.Tier2AlphaInstancePreprocessFreezeRecipeStepsStatusDetails()
        if Primitive.to_proto(resource.type_url):
            res.type_url = Primitive.to_proto(resource.type_url)
        if Primitive.to_proto(resource.value):
            res.value = Primitive.to_proto(resource.value)
        return res

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return None

        return InstancePreprocessFreezeRecipeStepsStatusDetails(
            type_url=Primitive.from_proto(resource.type_url),
            value=Primitive.from_proto(resource.value),
        )


class InstancePreprocessFreezeRecipeStepsStatusDetailsArray(object):
    @classmethod
    def to_proto(self, resources):
        if not resources:
            return resources
        return [
            InstancePreprocessFreezeRecipeStepsStatusDetails.to_proto(i)
            for i in resources
        ]

    @classmethod
    def from_proto(self, resources):
        return [
            InstancePreprocessFreezeRecipeStepsStatusDetails.from_proto(i)
            for i in resources
        ]


class InstancePreprocessFreezeRecipeStepsQuotaRequestDeltas(object):
    def __init__(
        self,
        metric_name: str = None,
        amount: int = None,
        quota_location_name: str = None,
    ):
        self.metric_name = metric_name
        self.amount = amount
        self.quota_location_name = quota_location_name

    @classmethod
    def to_proto(self, resource):
        if not resource:
            return None

        res = (
            instance_pb2.Tier2AlphaInstancePreprocessFreezeRecipeStepsQuotaRequestDeltas()
        )
        if Primitive.to_proto(resource.metric_name):
            res.metric_name = Primitive.to_proto(resource.metric_name)
        if Primitive.to_proto(resource.amount):
            res.amount = Primitive.to_proto(resource.amount)
        if Primitive.to_proto(resource.quota_location_name):
            res.quota_location_name = Primitive.to_proto(resource.quota_location_name)
        return res

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return None

        return InstancePreprocessFreezeRecipeStepsQuotaRequestDeltas(
            metric_name=Primitive.from_proto(resource.metric_name),
            amount=Primitive.from_proto(resource.amount),
            quota_location_name=Primitive.from_proto(resource.quota_location_name),
        )


class InstancePreprocessFreezeRecipeStepsQuotaRequestDeltasArray(object):
    @classmethod
    def to_proto(self, resources):
        if not resources:
            return resources
        return [
            InstancePreprocessFreezeRecipeStepsQuotaRequestDeltas.to_proto(i)
            for i in resources
        ]

    @classmethod
    def from_proto(self, resources):
        return [
            InstancePreprocessFreezeRecipeStepsQuotaRequestDeltas.from_proto(i)
            for i in resources
        ]


class InstancePreprocessFreezeRecipeStepsPreprocessUpdate(object):
    def __init__(
        self, latency_slo_bucket_name: str = None, public_operation_metadata: str = None
    ):
        self.latency_slo_bucket_name = latency_slo_bucket_name
        self.public_operation_metadata = public_operation_metadata

    @classmethod
    def to_proto(self, resource):
        if not resource:
            return None

        res = (
            instance_pb2.Tier2AlphaInstancePreprocessFreezeRecipeStepsPreprocessUpdate()
        )
        if Primitive.to_proto(resource.latency_slo_bucket_name):
            res.latency_slo_bucket_name = Primitive.to_proto(
                resource.latency_slo_bucket_name
            )
        if Primitive.to_proto(resource.public_operation_metadata):
            res.public_operation_metadata = Primitive.to_proto(
                resource.public_operation_metadata
            )
        return res

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return None

        return InstancePreprocessFreezeRecipeStepsPreprocessUpdate(
            latency_slo_bucket_name=Primitive.from_proto(
                resource.latency_slo_bucket_name
            ),
            public_operation_metadata=Primitive.from_proto(
                resource.public_operation_metadata
            ),
        )


class InstancePreprocessFreezeRecipeStepsPreprocessUpdateArray(object):
    @classmethod
    def to_proto(self, resources):
        if not resources:
            return resources
        return [
            InstancePreprocessFreezeRecipeStepsPreprocessUpdate.to_proto(i)
            for i in resources
        ]

    @classmethod
    def from_proto(self, resources):
        return [
            InstancePreprocessFreezeRecipeStepsPreprocessUpdate.from_proto(i)
            for i in resources
        ]


class InstancePreprocessFreezeRecipeStepsRequestedTenantProject(object):
    def __init__(self, tag: str = None, folder: str = None, scope: str = None):
        self.tag = tag
        self.folder = folder
        self.scope = scope

    @classmethod
    def to_proto(self, resource):
        if not resource:
            return None

        res = (
            instance_pb2.Tier2AlphaInstancePreprocessFreezeRecipeStepsRequestedTenantProject()
        )
        if Primitive.to_proto(resource.tag):
            res.tag = Primitive.to_proto(resource.tag)
        if Primitive.to_proto(resource.folder):
            res.folder = Primitive.to_proto(resource.folder)
        if InstancePreprocessFreezeRecipeStepsRequestedTenantProjectScopeEnum.to_proto(
            resource.scope
        ):
            res.scope = InstancePreprocessFreezeRecipeStepsRequestedTenantProjectScopeEnum.to_proto(
                resource.scope
            )
        return res

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return None

        return InstancePreprocessFreezeRecipeStepsRequestedTenantProject(
            tag=Primitive.from_proto(resource.tag),
            folder=Primitive.from_proto(resource.folder),
            scope=InstancePreprocessFreezeRecipeStepsRequestedTenantProjectScopeEnum.from_proto(
                resource.scope
            ),
        )


class InstancePreprocessFreezeRecipeStepsRequestedTenantProjectArray(object):
    @classmethod
    def to_proto(self, resources):
        if not resources:
            return resources
        return [
            InstancePreprocessFreezeRecipeStepsRequestedTenantProject.to_proto(i)
            for i in resources
        ]

    @classmethod
    def from_proto(self, resources):
        return [
            InstancePreprocessFreezeRecipeStepsRequestedTenantProject.from_proto(i)
            for i in resources
        ]


class InstancePreprocessFreezeRecipeStepsPermissionsInfo(object):
    def __init__(
        self,
        policy_name: dict = None,
        iam_permissions: list = None,
        resource_path: str = None,
        api_attrs: dict = None,
        policy_name_mode: str = None,
        resource: dict = None,
    ):
        self.policy_name = policy_name
        self.iam_permissions = iam_permissions
        self.resource_path = resource_path
        self.api_attrs = api_attrs
        self.policy_name_mode = policy_name_mode
        self.resource = resource

    @classmethod
    def to_proto(self, resource):
        if not resource:
            return None

        res = (
            instance_pb2.Tier2AlphaInstancePreprocessFreezeRecipeStepsPermissionsInfo()
        )
        if InstancePreprocessFreezeRecipeStepsPermissionsInfoPolicyName.to_proto(
            resource.policy_name
        ):
            res.policy_name.CopyFrom(
                InstancePreprocessFreezeRecipeStepsPermissionsInfoPolicyName.to_proto(
                    resource.policy_name
                )
            )
        else:
            res.ClearField("policy_name")
        if InstancePreprocessFreezeRecipeStepsPermissionsInfoIamPermissionsArray.to_proto(
            resource.iam_permissions
        ):
            res.iam_permissions.extend(
                InstancePreprocessFreezeRecipeStepsPermissionsInfoIamPermissionsArray.to_proto(
                    resource.iam_permissions
                )
            )
        if Primitive.to_proto(resource.resource_path):
            res.resource_path = Primitive.to_proto(resource.resource_path)
        if InstanceGoogleprotobufstruct.to_proto(resource.api_attrs):
            res.api_attrs.CopyFrom(
                InstanceGoogleprotobufstruct.to_proto(resource.api_attrs)
            )
        else:
            res.ClearField("api_attrs")
        if InstancePreprocessFreezeRecipeStepsPermissionsInfoPolicyNameModeEnum.to_proto(
            resource.policy_name_mode
        ):
            res.policy_name_mode = InstancePreprocessFreezeRecipeStepsPermissionsInfoPolicyNameModeEnum.to_proto(
                resource.policy_name_mode
            )
        if InstancePreprocessFreezeRecipeStepsPermissionsInfoResource.to_proto(
            resource.resource
        ):
            res.resource.CopyFrom(
                InstancePreprocessFreezeRecipeStepsPermissionsInfoResource.to_proto(
                    resource.resource
                )
            )
        else:
            res.ClearField("resource")
        return res

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return None

        return InstancePreprocessFreezeRecipeStepsPermissionsInfo(
            policy_name=InstancePreprocessFreezeRecipeStepsPermissionsInfoPolicyName.from_proto(
                resource.policy_name
            ),
            iam_permissions=InstancePreprocessFreezeRecipeStepsPermissionsInfoIamPermissionsArray.from_proto(
                resource.iam_permissions
            ),
            resource_path=Primitive.from_proto(resource.resource_path),
            api_attrs=InstanceGoogleprotobufstruct.from_proto(resource.api_attrs),
            policy_name_mode=InstancePreprocessFreezeRecipeStepsPermissionsInfoPolicyNameModeEnum.from_proto(
                resource.policy_name_mode
            ),
            resource=InstancePreprocessFreezeRecipeStepsPermissionsInfoResource.from_proto(
                resource.resource
            ),
        )


class InstancePreprocessFreezeRecipeStepsPermissionsInfoArray(object):
    @classmethod
    def to_proto(self, resources):
        if not resources:
            return resources
        return [
            InstancePreprocessFreezeRecipeStepsPermissionsInfo.to_proto(i)
            for i in resources
        ]

    @classmethod
    def from_proto(self, resources):
        return [
            InstancePreprocessFreezeRecipeStepsPermissionsInfo.from_proto(i)
            for i in resources
        ]


class InstancePreprocessFreezeRecipeStepsPermissionsInfoPolicyName(object):
    def __init__(self, type: str = None, id: str = None, region: str = None):
        self.type = type
        self.id = id
        self.region = region

    @classmethod
    def to_proto(self, resource):
        if not resource:
            return None

        res = (
            instance_pb2.Tier2AlphaInstancePreprocessFreezeRecipeStepsPermissionsInfoPolicyName()
        )
        if Primitive.to_proto(resource.type):
            res.type = Primitive.to_proto(resource.type)
        if Primitive.to_proto(resource.id):
            res.id = Primitive.to_proto(resource.id)
        if Primitive.to_proto(resource.region):
            res.region = Primitive.to_proto(resource.region)
        return res

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return None

        return InstancePreprocessFreezeRecipeStepsPermissionsInfoPolicyName(
            type=Primitive.from_proto(resource.type),
            id=Primitive.from_proto(resource.id),
            region=Primitive.from_proto(resource.region),
        )


class InstancePreprocessFreezeRecipeStepsPermissionsInfoPolicyNameArray(object):
    @classmethod
    def to_proto(self, resources):
        if not resources:
            return resources
        return [
            InstancePreprocessFreezeRecipeStepsPermissionsInfoPolicyName.to_proto(i)
            for i in resources
        ]

    @classmethod
    def from_proto(self, resources):
        return [
            InstancePreprocessFreezeRecipeStepsPermissionsInfoPolicyName.from_proto(i)
            for i in resources
        ]


class InstancePreprocessFreezeRecipeStepsPermissionsInfoIamPermissions(object):
    def __init__(self, permission: str = None):
        self.permission = permission

    @classmethod
    def to_proto(self, resource):
        if not resource:
            return None

        res = (
            instance_pb2.Tier2AlphaInstancePreprocessFreezeRecipeStepsPermissionsInfoIamPermissions()
        )
        if Primitive.to_proto(resource.permission):
            res.permission = Primitive.to_proto(resource.permission)
        return res

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return None

        return InstancePreprocessFreezeRecipeStepsPermissionsInfoIamPermissions(
            permission=Primitive.from_proto(resource.permission),
        )


class InstancePreprocessFreezeRecipeStepsPermissionsInfoIamPermissionsArray(object):
    @classmethod
    def to_proto(self, resources):
        if not resources:
            return resources
        return [
            InstancePreprocessFreezeRecipeStepsPermissionsInfoIamPermissions.to_proto(i)
            for i in resources
        ]

    @classmethod
    def from_proto(self, resources):
        return [
            InstancePreprocessFreezeRecipeStepsPermissionsInfoIamPermissions.from_proto(
                i
            )
            for i in resources
        ]


class InstancePreprocessFreezeRecipeStepsPermissionsInfoResource(object):
    def __init__(
        self,
        name: str = None,
        type: str = None,
        service: str = None,
        labels: dict = None,
    ):
        self.name = name
        self.type = type
        self.service = service
        self.labels = labels

    @classmethod
    def to_proto(self, resource):
        if not resource:
            return None

        res = (
            instance_pb2.Tier2AlphaInstancePreprocessFreezeRecipeStepsPermissionsInfoResource()
        )
        if Primitive.to_proto(resource.name):
            res.name = Primitive.to_proto(resource.name)
        if Primitive.to_proto(resource.type):
            res.type = Primitive.to_proto(resource.type)
        if Primitive.to_proto(resource.service):
            res.service = Primitive.to_proto(resource.service)
        if Primitive.to_proto(resource.labels):
            res.labels = Primitive.to_proto(resource.labels)
        return res

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return None

        return InstancePreprocessFreezeRecipeStepsPermissionsInfoResource(
            name=Primitive.from_proto(resource.name),
            type=Primitive.from_proto(resource.type),
            service=Primitive.from_proto(resource.service),
            labels=Primitive.from_proto(resource.labels),
        )


class InstancePreprocessFreezeRecipeStepsPermissionsInfoResourceArray(object):
    @classmethod
    def to_proto(self, resources):
        if not resources:
            return resources
        return [
            InstancePreprocessFreezeRecipeStepsPermissionsInfoResource.to_proto(i)
            for i in resources
        ]

    @classmethod
    def from_proto(self, resources):
        return [
            InstancePreprocessFreezeRecipeStepsPermissionsInfoResource.from_proto(i)
            for i in resources
        ]


class InstancePreprocessFreezeRecipeStepsKeyNotificationsUpdate(object):
    def __init__(self, key_notifications_info: dict = None):
        self.key_notifications_info = key_notifications_info

    @classmethod
    def to_proto(self, resource):
        if not resource:
            return None

        res = (
            instance_pb2.Tier2AlphaInstancePreprocessFreezeRecipeStepsKeyNotificationsUpdate()
        )
        if InstancePreprocessFreezeRecipeStepsKeyNotificationsUpdateKeyNotificationsInfo.to_proto(
            resource.key_notifications_info
        ):
            res.key_notifications_info.CopyFrom(
                InstancePreprocessFreezeRecipeStepsKeyNotificationsUpdateKeyNotificationsInfo.to_proto(
                    resource.key_notifications_info
                )
            )
        else:
            res.ClearField("key_notifications_info")
        return res

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return None

        return InstancePreprocessFreezeRecipeStepsKeyNotificationsUpdate(
            key_notifications_info=InstancePreprocessFreezeRecipeStepsKeyNotificationsUpdateKeyNotificationsInfo.from_proto(
                resource.key_notifications_info
            ),
        )


class InstancePreprocessFreezeRecipeStepsKeyNotificationsUpdateArray(object):
    @classmethod
    def to_proto(self, resources):
        if not resources:
            return resources
        return [
            InstancePreprocessFreezeRecipeStepsKeyNotificationsUpdate.to_proto(i)
            for i in resources
        ]

    @classmethod
    def from_proto(self, resources):
        return [
            InstancePreprocessFreezeRecipeStepsKeyNotificationsUpdate.from_proto(i)
            for i in resources
        ]


class InstancePreprocessFreezeRecipeStepsKeyNotificationsUpdateKeyNotificationsInfo(
    object
):
    def __init__(
        self,
        data_version: int = None,
        delegate: str = None,
        key_notification_configs: list = None,
    ):
        self.data_version = data_version
        self.delegate = delegate
        self.key_notification_configs = key_notification_configs

    @classmethod
    def to_proto(self, resource):
        if not resource:
            return None

        res = (
            instance_pb2.Tier2AlphaInstancePreprocessFreezeRecipeStepsKeyNotificationsUpdateKeyNotificationsInfo()
        )
        if Primitive.to_proto(resource.data_version):
            res.data_version = Primitive.to_proto(resource.data_version)
        if Primitive.to_proto(resource.delegate):
            res.delegate = Primitive.to_proto(resource.delegate)
        if InstancePreprocessFreezeRecipeStepsKeyNotificationsUpdateKeyNotificationsInfoKeyNotificationConfigsArray.to_proto(
            resource.key_notification_configs
        ):
            res.key_notification_configs.extend(
                InstancePreprocessFreezeRecipeStepsKeyNotificationsUpdateKeyNotificationsInfoKeyNotificationConfigsArray.to_proto(
                    resource.key_notification_configs
                )
            )
        return res

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return None

        return InstancePreprocessFreezeRecipeStepsKeyNotificationsUpdateKeyNotificationsInfo(
            data_version=Primitive.from_proto(resource.data_version),
            delegate=Primitive.from_proto(resource.delegate),
            key_notification_configs=InstancePreprocessFreezeRecipeStepsKeyNotificationsUpdateKeyNotificationsInfoKeyNotificationConfigsArray.from_proto(
                resource.key_notification_configs
            ),
        )


class InstancePreprocessFreezeRecipeStepsKeyNotificationsUpdateKeyNotificationsInfoArray(
    object
):
    @classmethod
    def to_proto(self, resources):
        if not resources:
            return resources
        return [
            InstancePreprocessFreezeRecipeStepsKeyNotificationsUpdateKeyNotificationsInfo.to_proto(
                i
            )
            for i in resources
        ]

    @classmethod
    def from_proto(self, resources):
        return [
            InstancePreprocessFreezeRecipeStepsKeyNotificationsUpdateKeyNotificationsInfo.from_proto(
                i
            )
            for i in resources
        ]


class InstancePreprocessFreezeRecipeStepsKeyNotificationsUpdateKeyNotificationsInfoKeyNotificationConfigs(
    object
):
    def __init__(
        self,
        key_or_version_name: str = None,
        grant: str = None,
        delegator_gaia_id: int = None,
    ):
        self.key_or_version_name = key_or_version_name
        self.grant = grant
        self.delegator_gaia_id = delegator_gaia_id

    @classmethod
    def to_proto(self, resource):
        if not resource:
            return None

        res = (
            instance_pb2.Tier2AlphaInstancePreprocessFreezeRecipeStepsKeyNotificationsUpdateKeyNotificationsInfoKeyNotificationConfigs()
        )
        if Primitive.to_proto(resource.key_or_version_name):
            res.key_or_version_name = Primitive.to_proto(resource.key_or_version_name)
        if Primitive.to_proto(resource.grant):
            res.grant = Primitive.to_proto(resource.grant)
        if Primitive.to_proto(resource.delegator_gaia_id):
            res.delegator_gaia_id = Primitive.to_proto(resource.delegator_gaia_id)
        return res

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return None

        return InstancePreprocessFreezeRecipeStepsKeyNotificationsUpdateKeyNotificationsInfoKeyNotificationConfigs(
            key_or_version_name=Primitive.from_proto(resource.key_or_version_name),
            grant=Primitive.from_proto(resource.grant),
            delegator_gaia_id=Primitive.from_proto(resource.delegator_gaia_id),
        )


class InstancePreprocessFreezeRecipeStepsKeyNotificationsUpdateKeyNotificationsInfoKeyNotificationConfigsArray(
    object
):
    @classmethod
    def to_proto(self, resources):
        if not resources:
            return resources
        return [
            InstancePreprocessFreezeRecipeStepsKeyNotificationsUpdateKeyNotificationsInfoKeyNotificationConfigs.to_proto(
                i
            )
            for i in resources
        ]

    @classmethod
    def from_proto(self, resources):
        return [
            InstancePreprocessFreezeRecipeStepsKeyNotificationsUpdateKeyNotificationsInfoKeyNotificationConfigs.from_proto(
                i
            )
            for i in resources
        ]


class InstanceFreezeRecipe(object):
    def __init__(
        self,
        steps: list = None,
        honor_cancel_request: bool = None,
        ignore_recipe_after: int = None,
        verify_deadline_seconds_below: float = None,
        populate_operation_result: bool = None,
        readonly_recipe_start_time: str = None,
        resource_names_stored_in_clh_with_delay: list = None,
        delay_to_store_resources_in_clh_db_nanos: int = None,
    ):
        self.steps = steps
        self.honor_cancel_request = honor_cancel_request
        self.ignore_recipe_after = ignore_recipe_after
        self.verify_deadline_seconds_below = verify_deadline_seconds_below
        self.populate_operation_result = populate_operation_result
        self.readonly_recipe_start_time = readonly_recipe_start_time
        self.resource_names_stored_in_clh_with_delay = (
            resource_names_stored_in_clh_with_delay
        )
        self.delay_to_store_resources_in_clh_db_nanos = (
            delay_to_store_resources_in_clh_db_nanos
        )

    @classmethod
    def to_proto(self, resource):
        if not resource:
            return None

        res = instance_pb2.Tier2AlphaInstanceFreezeRecipe()
        if InstanceFreezeRecipeStepsArray.to_proto(resource.steps):
            res.steps.extend(InstanceFreezeRecipeStepsArray.to_proto(resource.steps))
        if Primitive.to_proto(resource.honor_cancel_request):
            res.honor_cancel_request = Primitive.to_proto(resource.honor_cancel_request)
        if Primitive.to_proto(resource.ignore_recipe_after):
            res.ignore_recipe_after = Primitive.to_proto(resource.ignore_recipe_after)
        if Primitive.to_proto(resource.verify_deadline_seconds_below):
            res.verify_deadline_seconds_below = Primitive.to_proto(
                resource.verify_deadline_seconds_below
            )
        if Primitive.to_proto(resource.populate_operation_result):
            res.populate_operation_result = Primitive.to_proto(
                resource.populate_operation_result
            )
        if Primitive.to_proto(resource.readonly_recipe_start_time):
            res.readonly_recipe_start_time = Primitive.to_proto(
                resource.readonly_recipe_start_time
            )
        if Primitive.to_proto(resource.resource_names_stored_in_clh_with_delay):
            res.resource_names_stored_in_clh_with_delay.extend(
                Primitive.to_proto(resource.resource_names_stored_in_clh_with_delay)
            )
        if Primitive.to_proto(resource.delay_to_store_resources_in_clh_db_nanos):
            res.delay_to_store_resources_in_clh_db_nanos = Primitive.to_proto(
                resource.delay_to_store_resources_in_clh_db_nanos
            )
        return res

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return None

        return InstanceFreezeRecipe(
            steps=InstanceFreezeRecipeStepsArray.from_proto(resource.steps),
            honor_cancel_request=Primitive.from_proto(resource.honor_cancel_request),
            ignore_recipe_after=Primitive.from_proto(resource.ignore_recipe_after),
            verify_deadline_seconds_below=Primitive.from_proto(
                resource.verify_deadline_seconds_below
            ),
            populate_operation_result=Primitive.from_proto(
                resource.populate_operation_result
            ),
            readonly_recipe_start_time=Primitive.from_proto(
                resource.readonly_recipe_start_time
            ),
            resource_names_stored_in_clh_with_delay=Primitive.from_proto(
                resource.resource_names_stored_in_clh_with_delay
            ),
            delay_to_store_resources_in_clh_db_nanos=Primitive.from_proto(
                resource.delay_to_store_resources_in_clh_db_nanos
            ),
        )


class InstanceFreezeRecipeArray(object):
    @classmethod
    def to_proto(self, resources):
        if not resources:
            return resources
        return [InstanceFreezeRecipe.to_proto(i) for i in resources]

    @classmethod
    def from_proto(self, resources):
        return [InstanceFreezeRecipe.from_proto(i) for i in resources]


class InstanceFreezeRecipeSteps(object):
    def __init__(
        self,
        relative_time: int = None,
        sleep_duration: int = None,
        action: str = None,
        status: dict = None,
        error_space: str = None,
        p4_service_account: str = None,
        resource_metadata_size: int = None,
        description: str = None,
        updated_repeat_operation_delay_sec: float = None,
        quota_request_deltas: list = None,
        preprocess_update: dict = None,
        public_operation_metadata: str = None,
        requested_tenant_project: dict = None,
        permissions_info: list = None,
        key_notifications_update: dict = None,
        clh_data_update_time: str = None,
        public_error_message: str = None,
    ):
        self.relative_time = relative_time
        self.sleep_duration = sleep_duration
        self.action = action
        self.status = status
        self.error_space = error_space
        self.p4_service_account = p4_service_account
        self.resource_metadata_size = resource_metadata_size
        self.description = description
        self.updated_repeat_operation_delay_sec = updated_repeat_operation_delay_sec
        self.quota_request_deltas = quota_request_deltas
        self.preprocess_update = preprocess_update
        self.public_operation_metadata = public_operation_metadata
        self.requested_tenant_project = requested_tenant_project
        self.permissions_info = permissions_info
        self.key_notifications_update = key_notifications_update
        self.clh_data_update_time = clh_data_update_time
        self.public_error_message = public_error_message

    @classmethod
    def to_proto(self, resource):
        if not resource:
            return None

        res = instance_pb2.Tier2AlphaInstanceFreezeRecipeSteps()
        if Primitive.to_proto(resource.relative_time):
            res.relative_time = Primitive.to_proto(resource.relative_time)
        if Primitive.to_proto(resource.sleep_duration):
            res.sleep_duration = Primitive.to_proto(resource.sleep_duration)
        if InstanceFreezeRecipeStepsActionEnum.to_proto(resource.action):
            res.action = InstanceFreezeRecipeStepsActionEnum.to_proto(resource.action)
        if InstanceFreezeRecipeStepsStatus.to_proto(resource.status):
            res.status.CopyFrom(
                InstanceFreezeRecipeStepsStatus.to_proto(resource.status)
            )
        else:
            res.ClearField("status")
        if Primitive.to_proto(resource.error_space):
            res.error_space = Primitive.to_proto(resource.error_space)
        if Primitive.to_proto(resource.p4_service_account):
            res.p4_service_account = Primitive.to_proto(resource.p4_service_account)
        if Primitive.to_proto(resource.resource_metadata_size):
            res.resource_metadata_size = Primitive.to_proto(
                resource.resource_metadata_size
            )
        if Primitive.to_proto(resource.description):
            res.description = Primitive.to_proto(resource.description)
        if Primitive.to_proto(resource.updated_repeat_operation_delay_sec):
            res.updated_repeat_operation_delay_sec = Primitive.to_proto(
                resource.updated_repeat_operation_delay_sec
            )
        if InstanceFreezeRecipeStepsQuotaRequestDeltasArray.to_proto(
            resource.quota_request_deltas
        ):
            res.quota_request_deltas.extend(
                InstanceFreezeRecipeStepsQuotaRequestDeltasArray.to_proto(
                    resource.quota_request_deltas
                )
            )
        if InstanceFreezeRecipeStepsPreprocessUpdate.to_proto(
            resource.preprocess_update
        ):
            res.preprocess_update.CopyFrom(
                InstanceFreezeRecipeStepsPreprocessUpdate.to_proto(
                    resource.preprocess_update
                )
            )
        else:
            res.ClearField("preprocess_update")
        if Primitive.to_proto(resource.public_operation_metadata):
            res.public_operation_metadata = Primitive.to_proto(
                resource.public_operation_metadata
            )
        if InstanceFreezeRecipeStepsRequestedTenantProject.to_proto(
            resource.requested_tenant_project
        ):
            res.requested_tenant_project.CopyFrom(
                InstanceFreezeRecipeStepsRequestedTenantProject.to_proto(
                    resource.requested_tenant_project
                )
            )
        else:
            res.ClearField("requested_tenant_project")
        if InstanceFreezeRecipeStepsPermissionsInfoArray.to_proto(
            resource.permissions_info
        ):
            res.permissions_info.extend(
                InstanceFreezeRecipeStepsPermissionsInfoArray.to_proto(
                    resource.permissions_info
                )
            )
        if InstanceFreezeRecipeStepsKeyNotificationsUpdate.to_proto(
            resource.key_notifications_update
        ):
            res.key_notifications_update.CopyFrom(
                InstanceFreezeRecipeStepsKeyNotificationsUpdate.to_proto(
                    resource.key_notifications_update
                )
            )
        else:
            res.ClearField("key_notifications_update")
        if Primitive.to_proto(resource.clh_data_update_time):
            res.clh_data_update_time = Primitive.to_proto(resource.clh_data_update_time)
        if Primitive.to_proto(resource.public_error_message):
            res.public_error_message = Primitive.to_proto(resource.public_error_message)
        return res

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return None

        return InstanceFreezeRecipeSteps(
            relative_time=Primitive.from_proto(resource.relative_time),
            sleep_duration=Primitive.from_proto(resource.sleep_duration),
            action=InstanceFreezeRecipeStepsActionEnum.from_proto(resource.action),
            status=InstanceFreezeRecipeStepsStatus.from_proto(resource.status),
            error_space=Primitive.from_proto(resource.error_space),
            p4_service_account=Primitive.from_proto(resource.p4_service_account),
            resource_metadata_size=Primitive.from_proto(
                resource.resource_metadata_size
            ),
            description=Primitive.from_proto(resource.description),
            updated_repeat_operation_delay_sec=Primitive.from_proto(
                resource.updated_repeat_operation_delay_sec
            ),
            quota_request_deltas=InstanceFreezeRecipeStepsQuotaRequestDeltasArray.from_proto(
                resource.quota_request_deltas
            ),
            preprocess_update=InstanceFreezeRecipeStepsPreprocessUpdate.from_proto(
                resource.preprocess_update
            ),
            public_operation_metadata=Primitive.from_proto(
                resource.public_operation_metadata
            ),
            requested_tenant_project=InstanceFreezeRecipeStepsRequestedTenantProject.from_proto(
                resource.requested_tenant_project
            ),
            permissions_info=InstanceFreezeRecipeStepsPermissionsInfoArray.from_proto(
                resource.permissions_info
            ),
            key_notifications_update=InstanceFreezeRecipeStepsKeyNotificationsUpdate.from_proto(
                resource.key_notifications_update
            ),
            clh_data_update_time=Primitive.from_proto(resource.clh_data_update_time),
            public_error_message=Primitive.from_proto(resource.public_error_message),
        )


class InstanceFreezeRecipeStepsArray(object):
    @classmethod
    def to_proto(self, resources):
        if not resources:
            return resources
        return [InstanceFreezeRecipeSteps.to_proto(i) for i in resources]

    @classmethod
    def from_proto(self, resources):
        return [InstanceFreezeRecipeSteps.from_proto(i) for i in resources]


class InstanceFreezeRecipeStepsStatus(object):
    def __init__(self, code: int = None, message: str = None, details: list = None):
        self.code = code
        self.message = message
        self.details = details

    @classmethod
    def to_proto(self, resource):
        if not resource:
            return None

        res = instance_pb2.Tier2AlphaInstanceFreezeRecipeStepsStatus()
        if Primitive.to_proto(resource.code):
            res.code = Primitive.to_proto(resource.code)
        if Primitive.to_proto(resource.message):
            res.message = Primitive.to_proto(resource.message)
        if InstanceFreezeRecipeStepsStatusDetailsArray.to_proto(resource.details):
            res.details.extend(
                InstanceFreezeRecipeStepsStatusDetailsArray.to_proto(resource.details)
            )
        return res

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return None

        return InstanceFreezeRecipeStepsStatus(
            code=Primitive.from_proto(resource.code),
            message=Primitive.from_proto(resource.message),
            details=InstanceFreezeRecipeStepsStatusDetailsArray.from_proto(
                resource.details
            ),
        )


class InstanceFreezeRecipeStepsStatusArray(object):
    @classmethod
    def to_proto(self, resources):
        if not resources:
            return resources
        return [InstanceFreezeRecipeStepsStatus.to_proto(i) for i in resources]

    @classmethod
    def from_proto(self, resources):
        return [InstanceFreezeRecipeStepsStatus.from_proto(i) for i in resources]


class InstanceFreezeRecipeStepsStatusDetails(object):
    def __init__(self, type_url: str = None, value: str = None):
        self.type_url = type_url
        self.value = value

    @classmethod
    def to_proto(self, resource):
        if not resource:
            return None

        res = instance_pb2.Tier2AlphaInstanceFreezeRecipeStepsStatusDetails()
        if Primitive.to_proto(resource.type_url):
            res.type_url = Primitive.to_proto(resource.type_url)
        if Primitive.to_proto(resource.value):
            res.value = Primitive.to_proto(resource.value)
        return res

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return None

        return InstanceFreezeRecipeStepsStatusDetails(
            type_url=Primitive.from_proto(resource.type_url),
            value=Primitive.from_proto(resource.value),
        )


class InstanceFreezeRecipeStepsStatusDetailsArray(object):
    @classmethod
    def to_proto(self, resources):
        if not resources:
            return resources
        return [InstanceFreezeRecipeStepsStatusDetails.to_proto(i) for i in resources]

    @classmethod
    def from_proto(self, resources):
        return [InstanceFreezeRecipeStepsStatusDetails.from_proto(i) for i in resources]


class InstanceFreezeRecipeStepsQuotaRequestDeltas(object):
    def __init__(
        self,
        metric_name: str = None,
        amount: int = None,
        quota_location_name: str = None,
    ):
        self.metric_name = metric_name
        self.amount = amount
        self.quota_location_name = quota_location_name

    @classmethod
    def to_proto(self, resource):
        if not resource:
            return None

        res = instance_pb2.Tier2AlphaInstanceFreezeRecipeStepsQuotaRequestDeltas()
        if Primitive.to_proto(resource.metric_name):
            res.metric_name = Primitive.to_proto(resource.metric_name)
        if Primitive.to_proto(resource.amount):
            res.amount = Primitive.to_proto(resource.amount)
        if Primitive.to_proto(resource.quota_location_name):
            res.quota_location_name = Primitive.to_proto(resource.quota_location_name)
        return res

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return None

        return InstanceFreezeRecipeStepsQuotaRequestDeltas(
            metric_name=Primitive.from_proto(resource.metric_name),
            amount=Primitive.from_proto(resource.amount),
            quota_location_name=Primitive.from_proto(resource.quota_location_name),
        )


class InstanceFreezeRecipeStepsQuotaRequestDeltasArray(object):
    @classmethod
    def to_proto(self, resources):
        if not resources:
            return resources
        return [
            InstanceFreezeRecipeStepsQuotaRequestDeltas.to_proto(i) for i in resources
        ]

    @classmethod
    def from_proto(self, resources):
        return [
            InstanceFreezeRecipeStepsQuotaRequestDeltas.from_proto(i) for i in resources
        ]


class InstanceFreezeRecipeStepsPreprocessUpdate(object):
    def __init__(
        self, latency_slo_bucket_name: str = None, public_operation_metadata: str = None
    ):
        self.latency_slo_bucket_name = latency_slo_bucket_name
        self.public_operation_metadata = public_operation_metadata

    @classmethod
    def to_proto(self, resource):
        if not resource:
            return None

        res = instance_pb2.Tier2AlphaInstanceFreezeRecipeStepsPreprocessUpdate()
        if Primitive.to_proto(resource.latency_slo_bucket_name):
            res.latency_slo_bucket_name = Primitive.to_proto(
                resource.latency_slo_bucket_name
            )
        if Primitive.to_proto(resource.public_operation_metadata):
            res.public_operation_metadata = Primitive.to_proto(
                resource.public_operation_metadata
            )
        return res

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return None

        return InstanceFreezeRecipeStepsPreprocessUpdate(
            latency_slo_bucket_name=Primitive.from_proto(
                resource.latency_slo_bucket_name
            ),
            public_operation_metadata=Primitive.from_proto(
                resource.public_operation_metadata
            ),
        )


class InstanceFreezeRecipeStepsPreprocessUpdateArray(object):
    @classmethod
    def to_proto(self, resources):
        if not resources:
            return resources
        return [
            InstanceFreezeRecipeStepsPreprocessUpdate.to_proto(i) for i in resources
        ]

    @classmethod
    def from_proto(self, resources):
        return [
            InstanceFreezeRecipeStepsPreprocessUpdate.from_proto(i) for i in resources
        ]


class InstanceFreezeRecipeStepsRequestedTenantProject(object):
    def __init__(self, tag: str = None, folder: str = None, scope: str = None):
        self.tag = tag
        self.folder = folder
        self.scope = scope

    @classmethod
    def to_proto(self, resource):
        if not resource:
            return None

        res = instance_pb2.Tier2AlphaInstanceFreezeRecipeStepsRequestedTenantProject()
        if Primitive.to_proto(resource.tag):
            res.tag = Primitive.to_proto(resource.tag)
        if Primitive.to_proto(resource.folder):
            res.folder = Primitive.to_proto(resource.folder)
        if InstanceFreezeRecipeStepsRequestedTenantProjectScopeEnum.to_proto(
            resource.scope
        ):
            res.scope = InstanceFreezeRecipeStepsRequestedTenantProjectScopeEnum.to_proto(
                resource.scope
            )
        return res

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return None

        return InstanceFreezeRecipeStepsRequestedTenantProject(
            tag=Primitive.from_proto(resource.tag),
            folder=Primitive.from_proto(resource.folder),
            scope=InstanceFreezeRecipeStepsRequestedTenantProjectScopeEnum.from_proto(
                resource.scope
            ),
        )


class InstanceFreezeRecipeStepsRequestedTenantProjectArray(object):
    @classmethod
    def to_proto(self, resources):
        if not resources:
            return resources
        return [
            InstanceFreezeRecipeStepsRequestedTenantProject.to_proto(i)
            for i in resources
        ]

    @classmethod
    def from_proto(self, resources):
        return [
            InstanceFreezeRecipeStepsRequestedTenantProject.from_proto(i)
            for i in resources
        ]


class InstanceFreezeRecipeStepsPermissionsInfo(object):
    def __init__(
        self,
        policy_name: dict = None,
        iam_permissions: list = None,
        resource_path: str = None,
        api_attrs: dict = None,
        policy_name_mode: str = None,
        resource: dict = None,
    ):
        self.policy_name = policy_name
        self.iam_permissions = iam_permissions
        self.resource_path = resource_path
        self.api_attrs = api_attrs
        self.policy_name_mode = policy_name_mode
        self.resource = resource

    @classmethod
    def to_proto(self, resource):
        if not resource:
            return None

        res = instance_pb2.Tier2AlphaInstanceFreezeRecipeStepsPermissionsInfo()
        if InstanceFreezeRecipeStepsPermissionsInfoPolicyName.to_proto(
            resource.policy_name
        ):
            res.policy_name.CopyFrom(
                InstanceFreezeRecipeStepsPermissionsInfoPolicyName.to_proto(
                    resource.policy_name
                )
            )
        else:
            res.ClearField("policy_name")
        if InstanceFreezeRecipeStepsPermissionsInfoIamPermissionsArray.to_proto(
            resource.iam_permissions
        ):
            res.iam_permissions.extend(
                InstanceFreezeRecipeStepsPermissionsInfoIamPermissionsArray.to_proto(
                    resource.iam_permissions
                )
            )
        if Primitive.to_proto(resource.resource_path):
            res.resource_path = Primitive.to_proto(resource.resource_path)
        if InstanceGoogleprotobufstruct.to_proto(resource.api_attrs):
            res.api_attrs.CopyFrom(
                InstanceGoogleprotobufstruct.to_proto(resource.api_attrs)
            )
        else:
            res.ClearField("api_attrs")
        if InstanceFreezeRecipeStepsPermissionsInfoPolicyNameModeEnum.to_proto(
            resource.policy_name_mode
        ):
            res.policy_name_mode = InstanceFreezeRecipeStepsPermissionsInfoPolicyNameModeEnum.to_proto(
                resource.policy_name_mode
            )
        if InstanceFreezeRecipeStepsPermissionsInfoResource.to_proto(resource.resource):
            res.resource.CopyFrom(
                InstanceFreezeRecipeStepsPermissionsInfoResource.to_proto(
                    resource.resource
                )
            )
        else:
            res.ClearField("resource")
        return res

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return None

        return InstanceFreezeRecipeStepsPermissionsInfo(
            policy_name=InstanceFreezeRecipeStepsPermissionsInfoPolicyName.from_proto(
                resource.policy_name
            ),
            iam_permissions=InstanceFreezeRecipeStepsPermissionsInfoIamPermissionsArray.from_proto(
                resource.iam_permissions
            ),
            resource_path=Primitive.from_proto(resource.resource_path),
            api_attrs=InstanceGoogleprotobufstruct.from_proto(resource.api_attrs),
            policy_name_mode=InstanceFreezeRecipeStepsPermissionsInfoPolicyNameModeEnum.from_proto(
                resource.policy_name_mode
            ),
            resource=InstanceFreezeRecipeStepsPermissionsInfoResource.from_proto(
                resource.resource
            ),
        )


class InstanceFreezeRecipeStepsPermissionsInfoArray(object):
    @classmethod
    def to_proto(self, resources):
        if not resources:
            return resources
        return [InstanceFreezeRecipeStepsPermissionsInfo.to_proto(i) for i in resources]

    @classmethod
    def from_proto(self, resources):
        return [
            InstanceFreezeRecipeStepsPermissionsInfo.from_proto(i) for i in resources
        ]


class InstanceFreezeRecipeStepsPermissionsInfoPolicyName(object):
    def __init__(self, type: str = None, id: str = None, region: str = None):
        self.type = type
        self.id = id
        self.region = region

    @classmethod
    def to_proto(self, resource):
        if not resource:
            return None

        res = (
            instance_pb2.Tier2AlphaInstanceFreezeRecipeStepsPermissionsInfoPolicyName()
        )
        if Primitive.to_proto(resource.type):
            res.type = Primitive.to_proto(resource.type)
        if Primitive.to_proto(resource.id):
            res.id = Primitive.to_proto(resource.id)
        if Primitive.to_proto(resource.region):
            res.region = Primitive.to_proto(resource.region)
        return res

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return None

        return InstanceFreezeRecipeStepsPermissionsInfoPolicyName(
            type=Primitive.from_proto(resource.type),
            id=Primitive.from_proto(resource.id),
            region=Primitive.from_proto(resource.region),
        )


class InstanceFreezeRecipeStepsPermissionsInfoPolicyNameArray(object):
    @classmethod
    def to_proto(self, resources):
        if not resources:
            return resources
        return [
            InstanceFreezeRecipeStepsPermissionsInfoPolicyName.to_proto(i)
            for i in resources
        ]

    @classmethod
    def from_proto(self, resources):
        return [
            InstanceFreezeRecipeStepsPermissionsInfoPolicyName.from_proto(i)
            for i in resources
        ]


class InstanceFreezeRecipeStepsPermissionsInfoIamPermissions(object):
    def __init__(self, permission: str = None):
        self.permission = permission

    @classmethod
    def to_proto(self, resource):
        if not resource:
            return None

        res = (
            instance_pb2.Tier2AlphaInstanceFreezeRecipeStepsPermissionsInfoIamPermissions()
        )
        if Primitive.to_proto(resource.permission):
            res.permission = Primitive.to_proto(resource.permission)
        return res

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return None

        return InstanceFreezeRecipeStepsPermissionsInfoIamPermissions(
            permission=Primitive.from_proto(resource.permission),
        )


class InstanceFreezeRecipeStepsPermissionsInfoIamPermissionsArray(object):
    @classmethod
    def to_proto(self, resources):
        if not resources:
            return resources
        return [
            InstanceFreezeRecipeStepsPermissionsInfoIamPermissions.to_proto(i)
            for i in resources
        ]

    @classmethod
    def from_proto(self, resources):
        return [
            InstanceFreezeRecipeStepsPermissionsInfoIamPermissions.from_proto(i)
            for i in resources
        ]


class InstanceFreezeRecipeStepsPermissionsInfoResource(object):
    def __init__(
        self,
        name: str = None,
        type: str = None,
        service: str = None,
        labels: dict = None,
    ):
        self.name = name
        self.type = type
        self.service = service
        self.labels = labels

    @classmethod
    def to_proto(self, resource):
        if not resource:
            return None

        res = instance_pb2.Tier2AlphaInstanceFreezeRecipeStepsPermissionsInfoResource()
        if Primitive.to_proto(resource.name):
            res.name = Primitive.to_proto(resource.name)
        if Primitive.to_proto(resource.type):
            res.type = Primitive.to_proto(resource.type)
        if Primitive.to_proto(resource.service):
            res.service = Primitive.to_proto(resource.service)
        if Primitive.to_proto(resource.labels):
            res.labels = Primitive.to_proto(resource.labels)
        return res

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return None

        return InstanceFreezeRecipeStepsPermissionsInfoResource(
            name=Primitive.from_proto(resource.name),
            type=Primitive.from_proto(resource.type),
            service=Primitive.from_proto(resource.service),
            labels=Primitive.from_proto(resource.labels),
        )


class InstanceFreezeRecipeStepsPermissionsInfoResourceArray(object):
    @classmethod
    def to_proto(self, resources):
        if not resources:
            return resources
        return [
            InstanceFreezeRecipeStepsPermissionsInfoResource.to_proto(i)
            for i in resources
        ]

    @classmethod
    def from_proto(self, resources):
        return [
            InstanceFreezeRecipeStepsPermissionsInfoResource.from_proto(i)
            for i in resources
        ]


class InstanceFreezeRecipeStepsKeyNotificationsUpdate(object):
    def __init__(self, key_notifications_info: dict = None):
        self.key_notifications_info = key_notifications_info

    @classmethod
    def to_proto(self, resource):
        if not resource:
            return None

        res = instance_pb2.Tier2AlphaInstanceFreezeRecipeStepsKeyNotificationsUpdate()
        if InstanceFreezeRecipeStepsKeyNotificationsUpdateKeyNotificationsInfo.to_proto(
            resource.key_notifications_info
        ):
            res.key_notifications_info.CopyFrom(
                InstanceFreezeRecipeStepsKeyNotificationsUpdateKeyNotificationsInfo.to_proto(
                    resource.key_notifications_info
                )
            )
        else:
            res.ClearField("key_notifications_info")
        return res

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return None

        return InstanceFreezeRecipeStepsKeyNotificationsUpdate(
            key_notifications_info=InstanceFreezeRecipeStepsKeyNotificationsUpdateKeyNotificationsInfo.from_proto(
                resource.key_notifications_info
            ),
        )


class InstanceFreezeRecipeStepsKeyNotificationsUpdateArray(object):
    @classmethod
    def to_proto(self, resources):
        if not resources:
            return resources
        return [
            InstanceFreezeRecipeStepsKeyNotificationsUpdate.to_proto(i)
            for i in resources
        ]

    @classmethod
    def from_proto(self, resources):
        return [
            InstanceFreezeRecipeStepsKeyNotificationsUpdate.from_proto(i)
            for i in resources
        ]


class InstanceFreezeRecipeStepsKeyNotificationsUpdateKeyNotificationsInfo(object):
    def __init__(
        self,
        data_version: int = None,
        delegate: str = None,
        key_notification_configs: list = None,
    ):
        self.data_version = data_version
        self.delegate = delegate
        self.key_notification_configs = key_notification_configs

    @classmethod
    def to_proto(self, resource):
        if not resource:
            return None

        res = (
            instance_pb2.Tier2AlphaInstanceFreezeRecipeStepsKeyNotificationsUpdateKeyNotificationsInfo()
        )
        if Primitive.to_proto(resource.data_version):
            res.data_version = Primitive.to_proto(resource.data_version)
        if Primitive.to_proto(resource.delegate):
            res.delegate = Primitive.to_proto(resource.delegate)
        if InstanceFreezeRecipeStepsKeyNotificationsUpdateKeyNotificationsInfoKeyNotificationConfigsArray.to_proto(
            resource.key_notification_configs
        ):
            res.key_notification_configs.extend(
                InstanceFreezeRecipeStepsKeyNotificationsUpdateKeyNotificationsInfoKeyNotificationConfigsArray.to_proto(
                    resource.key_notification_configs
                )
            )
        return res

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return None

        return InstanceFreezeRecipeStepsKeyNotificationsUpdateKeyNotificationsInfo(
            data_version=Primitive.from_proto(resource.data_version),
            delegate=Primitive.from_proto(resource.delegate),
            key_notification_configs=InstanceFreezeRecipeStepsKeyNotificationsUpdateKeyNotificationsInfoKeyNotificationConfigsArray.from_proto(
                resource.key_notification_configs
            ),
        )


class InstanceFreezeRecipeStepsKeyNotificationsUpdateKeyNotificationsInfoArray(object):
    @classmethod
    def to_proto(self, resources):
        if not resources:
            return resources
        return [
            InstanceFreezeRecipeStepsKeyNotificationsUpdateKeyNotificationsInfo.to_proto(
                i
            )
            for i in resources
        ]

    @classmethod
    def from_proto(self, resources):
        return [
            InstanceFreezeRecipeStepsKeyNotificationsUpdateKeyNotificationsInfo.from_proto(
                i
            )
            for i in resources
        ]


class InstanceFreezeRecipeStepsKeyNotificationsUpdateKeyNotificationsInfoKeyNotificationConfigs(
    object
):
    def __init__(
        self,
        key_or_version_name: str = None,
        grant: str = None,
        delegator_gaia_id: int = None,
    ):
        self.key_or_version_name = key_or_version_name
        self.grant = grant
        self.delegator_gaia_id = delegator_gaia_id

    @classmethod
    def to_proto(self, resource):
        if not resource:
            return None

        res = (
            instance_pb2.Tier2AlphaInstanceFreezeRecipeStepsKeyNotificationsUpdateKeyNotificationsInfoKeyNotificationConfigs()
        )
        if Primitive.to_proto(resource.key_or_version_name):
            res.key_or_version_name = Primitive.to_proto(resource.key_or_version_name)
        if Primitive.to_proto(resource.grant):
            res.grant = Primitive.to_proto(resource.grant)
        if Primitive.to_proto(resource.delegator_gaia_id):
            res.delegator_gaia_id = Primitive.to_proto(resource.delegator_gaia_id)
        return res

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return None

        return InstanceFreezeRecipeStepsKeyNotificationsUpdateKeyNotificationsInfoKeyNotificationConfigs(
            key_or_version_name=Primitive.from_proto(resource.key_or_version_name),
            grant=Primitive.from_proto(resource.grant),
            delegator_gaia_id=Primitive.from_proto(resource.delegator_gaia_id),
        )


class InstanceFreezeRecipeStepsKeyNotificationsUpdateKeyNotificationsInfoKeyNotificationConfigsArray(
    object
):
    @classmethod
    def to_proto(self, resources):
        if not resources:
            return resources
        return [
            InstanceFreezeRecipeStepsKeyNotificationsUpdateKeyNotificationsInfoKeyNotificationConfigs.to_proto(
                i
            )
            for i in resources
        ]

    @classmethod
    def from_proto(self, resources):
        return [
            InstanceFreezeRecipeStepsKeyNotificationsUpdateKeyNotificationsInfoKeyNotificationConfigs.from_proto(
                i
            )
            for i in resources
        ]


class InstancePreprocessUnfreezeRecipe(object):
    def __init__(
        self,
        steps: list = None,
        honor_cancel_request: bool = None,
        ignore_recipe_after: int = None,
        verify_deadline_seconds_below: float = None,
        populate_operation_result: bool = None,
        readonly_recipe_start_time: str = None,
        resource_names_stored_in_clh_with_delay: list = None,
        delay_to_store_resources_in_clh_db_nanos: int = None,
    ):
        self.steps = steps
        self.honor_cancel_request = honor_cancel_request
        self.ignore_recipe_after = ignore_recipe_after
        self.verify_deadline_seconds_below = verify_deadline_seconds_below
        self.populate_operation_result = populate_operation_result
        self.readonly_recipe_start_time = readonly_recipe_start_time
        self.resource_names_stored_in_clh_with_delay = (
            resource_names_stored_in_clh_with_delay
        )
        self.delay_to_store_resources_in_clh_db_nanos = (
            delay_to_store_resources_in_clh_db_nanos
        )

    @classmethod
    def to_proto(self, resource):
        if not resource:
            return None

        res = instance_pb2.Tier2AlphaInstancePreprocessUnfreezeRecipe()
        if InstancePreprocessUnfreezeRecipeStepsArray.to_proto(resource.steps):
            res.steps.extend(
                InstancePreprocessUnfreezeRecipeStepsArray.to_proto(resource.steps)
            )
        if Primitive.to_proto(resource.honor_cancel_request):
            res.honor_cancel_request = Primitive.to_proto(resource.honor_cancel_request)
        if Primitive.to_proto(resource.ignore_recipe_after):
            res.ignore_recipe_after = Primitive.to_proto(resource.ignore_recipe_after)
        if Primitive.to_proto(resource.verify_deadline_seconds_below):
            res.verify_deadline_seconds_below = Primitive.to_proto(
                resource.verify_deadline_seconds_below
            )
        if Primitive.to_proto(resource.populate_operation_result):
            res.populate_operation_result = Primitive.to_proto(
                resource.populate_operation_result
            )
        if Primitive.to_proto(resource.readonly_recipe_start_time):
            res.readonly_recipe_start_time = Primitive.to_proto(
                resource.readonly_recipe_start_time
            )
        if Primitive.to_proto(resource.resource_names_stored_in_clh_with_delay):
            res.resource_names_stored_in_clh_with_delay.extend(
                Primitive.to_proto(resource.resource_names_stored_in_clh_with_delay)
            )
        if Primitive.to_proto(resource.delay_to_store_resources_in_clh_db_nanos):
            res.delay_to_store_resources_in_clh_db_nanos = Primitive.to_proto(
                resource.delay_to_store_resources_in_clh_db_nanos
            )
        return res

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return None

        return InstancePreprocessUnfreezeRecipe(
            steps=InstancePreprocessUnfreezeRecipeStepsArray.from_proto(resource.steps),
            honor_cancel_request=Primitive.from_proto(resource.honor_cancel_request),
            ignore_recipe_after=Primitive.from_proto(resource.ignore_recipe_after),
            verify_deadline_seconds_below=Primitive.from_proto(
                resource.verify_deadline_seconds_below
            ),
            populate_operation_result=Primitive.from_proto(
                resource.populate_operation_result
            ),
            readonly_recipe_start_time=Primitive.from_proto(
                resource.readonly_recipe_start_time
            ),
            resource_names_stored_in_clh_with_delay=Primitive.from_proto(
                resource.resource_names_stored_in_clh_with_delay
            ),
            delay_to_store_resources_in_clh_db_nanos=Primitive.from_proto(
                resource.delay_to_store_resources_in_clh_db_nanos
            ),
        )


class InstancePreprocessUnfreezeRecipeArray(object):
    @classmethod
    def to_proto(self, resources):
        if not resources:
            return resources
        return [InstancePreprocessUnfreezeRecipe.to_proto(i) for i in resources]

    @classmethod
    def from_proto(self, resources):
        return [InstancePreprocessUnfreezeRecipe.from_proto(i) for i in resources]


class InstancePreprocessUnfreezeRecipeSteps(object):
    def __init__(
        self,
        relative_time: int = None,
        sleep_duration: int = None,
        action: str = None,
        status: dict = None,
        error_space: str = None,
        p4_service_account: str = None,
        resource_metadata_size: int = None,
        description: str = None,
        updated_repeat_operation_delay_sec: float = None,
        quota_request_deltas: list = None,
        preprocess_update: dict = None,
        public_operation_metadata: str = None,
        requested_tenant_project: dict = None,
        permissions_info: list = None,
        key_notifications_update: dict = None,
        clh_data_update_time: str = None,
        public_error_message: str = None,
    ):
        self.relative_time = relative_time
        self.sleep_duration = sleep_duration
        self.action = action
        self.status = status
        self.error_space = error_space
        self.p4_service_account = p4_service_account
        self.resource_metadata_size = resource_metadata_size
        self.description = description
        self.updated_repeat_operation_delay_sec = updated_repeat_operation_delay_sec
        self.quota_request_deltas = quota_request_deltas
        self.preprocess_update = preprocess_update
        self.public_operation_metadata = public_operation_metadata
        self.requested_tenant_project = requested_tenant_project
        self.permissions_info = permissions_info
        self.key_notifications_update = key_notifications_update
        self.clh_data_update_time = clh_data_update_time
        self.public_error_message = public_error_message

    @classmethod
    def to_proto(self, resource):
        if not resource:
            return None

        res = instance_pb2.Tier2AlphaInstancePreprocessUnfreezeRecipeSteps()
        if Primitive.to_proto(resource.relative_time):
            res.relative_time = Primitive.to_proto(resource.relative_time)
        if Primitive.to_proto(resource.sleep_duration):
            res.sleep_duration = Primitive.to_proto(resource.sleep_duration)
        if InstancePreprocessUnfreezeRecipeStepsActionEnum.to_proto(resource.action):
            res.action = InstancePreprocessUnfreezeRecipeStepsActionEnum.to_proto(
                resource.action
            )
        if InstancePreprocessUnfreezeRecipeStepsStatus.to_proto(resource.status):
            res.status.CopyFrom(
                InstancePreprocessUnfreezeRecipeStepsStatus.to_proto(resource.status)
            )
        else:
            res.ClearField("status")
        if Primitive.to_proto(resource.error_space):
            res.error_space = Primitive.to_proto(resource.error_space)
        if Primitive.to_proto(resource.p4_service_account):
            res.p4_service_account = Primitive.to_proto(resource.p4_service_account)
        if Primitive.to_proto(resource.resource_metadata_size):
            res.resource_metadata_size = Primitive.to_proto(
                resource.resource_metadata_size
            )
        if Primitive.to_proto(resource.description):
            res.description = Primitive.to_proto(resource.description)
        if Primitive.to_proto(resource.updated_repeat_operation_delay_sec):
            res.updated_repeat_operation_delay_sec = Primitive.to_proto(
                resource.updated_repeat_operation_delay_sec
            )
        if InstancePreprocessUnfreezeRecipeStepsQuotaRequestDeltasArray.to_proto(
            resource.quota_request_deltas
        ):
            res.quota_request_deltas.extend(
                InstancePreprocessUnfreezeRecipeStepsQuotaRequestDeltasArray.to_proto(
                    resource.quota_request_deltas
                )
            )
        if InstancePreprocessUnfreezeRecipeStepsPreprocessUpdate.to_proto(
            resource.preprocess_update
        ):
            res.preprocess_update.CopyFrom(
                InstancePreprocessUnfreezeRecipeStepsPreprocessUpdate.to_proto(
                    resource.preprocess_update
                )
            )
        else:
            res.ClearField("preprocess_update")
        if Primitive.to_proto(resource.public_operation_metadata):
            res.public_operation_metadata = Primitive.to_proto(
                resource.public_operation_metadata
            )
        if InstancePreprocessUnfreezeRecipeStepsRequestedTenantProject.to_proto(
            resource.requested_tenant_project
        ):
            res.requested_tenant_project.CopyFrom(
                InstancePreprocessUnfreezeRecipeStepsRequestedTenantProject.to_proto(
                    resource.requested_tenant_project
                )
            )
        else:
            res.ClearField("requested_tenant_project")
        if InstancePreprocessUnfreezeRecipeStepsPermissionsInfoArray.to_proto(
            resource.permissions_info
        ):
            res.permissions_info.extend(
                InstancePreprocessUnfreezeRecipeStepsPermissionsInfoArray.to_proto(
                    resource.permissions_info
                )
            )
        if InstancePreprocessUnfreezeRecipeStepsKeyNotificationsUpdate.to_proto(
            resource.key_notifications_update
        ):
            res.key_notifications_update.CopyFrom(
                InstancePreprocessUnfreezeRecipeStepsKeyNotificationsUpdate.to_proto(
                    resource.key_notifications_update
                )
            )
        else:
            res.ClearField("key_notifications_update")
        if Primitive.to_proto(resource.clh_data_update_time):
            res.clh_data_update_time = Primitive.to_proto(resource.clh_data_update_time)
        if Primitive.to_proto(resource.public_error_message):
            res.public_error_message = Primitive.to_proto(resource.public_error_message)
        return res

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return None

        return InstancePreprocessUnfreezeRecipeSteps(
            relative_time=Primitive.from_proto(resource.relative_time),
            sleep_duration=Primitive.from_proto(resource.sleep_duration),
            action=InstancePreprocessUnfreezeRecipeStepsActionEnum.from_proto(
                resource.action
            ),
            status=InstancePreprocessUnfreezeRecipeStepsStatus.from_proto(
                resource.status
            ),
            error_space=Primitive.from_proto(resource.error_space),
            p4_service_account=Primitive.from_proto(resource.p4_service_account),
            resource_metadata_size=Primitive.from_proto(
                resource.resource_metadata_size
            ),
            description=Primitive.from_proto(resource.description),
            updated_repeat_operation_delay_sec=Primitive.from_proto(
                resource.updated_repeat_operation_delay_sec
            ),
            quota_request_deltas=InstancePreprocessUnfreezeRecipeStepsQuotaRequestDeltasArray.from_proto(
                resource.quota_request_deltas
            ),
            preprocess_update=InstancePreprocessUnfreezeRecipeStepsPreprocessUpdate.from_proto(
                resource.preprocess_update
            ),
            public_operation_metadata=Primitive.from_proto(
                resource.public_operation_metadata
            ),
            requested_tenant_project=InstancePreprocessUnfreezeRecipeStepsRequestedTenantProject.from_proto(
                resource.requested_tenant_project
            ),
            permissions_info=InstancePreprocessUnfreezeRecipeStepsPermissionsInfoArray.from_proto(
                resource.permissions_info
            ),
            key_notifications_update=InstancePreprocessUnfreezeRecipeStepsKeyNotificationsUpdate.from_proto(
                resource.key_notifications_update
            ),
            clh_data_update_time=Primitive.from_proto(resource.clh_data_update_time),
            public_error_message=Primitive.from_proto(resource.public_error_message),
        )


class InstancePreprocessUnfreezeRecipeStepsArray(object):
    @classmethod
    def to_proto(self, resources):
        if not resources:
            return resources
        return [InstancePreprocessUnfreezeRecipeSteps.to_proto(i) for i in resources]

    @classmethod
    def from_proto(self, resources):
        return [InstancePreprocessUnfreezeRecipeSteps.from_proto(i) for i in resources]


class InstancePreprocessUnfreezeRecipeStepsStatus(object):
    def __init__(self, code: int = None, message: str = None, details: list = None):
        self.code = code
        self.message = message
        self.details = details

    @classmethod
    def to_proto(self, resource):
        if not resource:
            return None

        res = instance_pb2.Tier2AlphaInstancePreprocessUnfreezeRecipeStepsStatus()
        if Primitive.to_proto(resource.code):
            res.code = Primitive.to_proto(resource.code)
        if Primitive.to_proto(resource.message):
            res.message = Primitive.to_proto(resource.message)
        if InstancePreprocessUnfreezeRecipeStepsStatusDetailsArray.to_proto(
            resource.details
        ):
            res.details.extend(
                InstancePreprocessUnfreezeRecipeStepsStatusDetailsArray.to_proto(
                    resource.details
                )
            )
        return res

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return None

        return InstancePreprocessUnfreezeRecipeStepsStatus(
            code=Primitive.from_proto(resource.code),
            message=Primitive.from_proto(resource.message),
            details=InstancePreprocessUnfreezeRecipeStepsStatusDetailsArray.from_proto(
                resource.details
            ),
        )


class InstancePreprocessUnfreezeRecipeStepsStatusArray(object):
    @classmethod
    def to_proto(self, resources):
        if not resources:
            return resources
        return [
            InstancePreprocessUnfreezeRecipeStepsStatus.to_proto(i) for i in resources
        ]

    @classmethod
    def from_proto(self, resources):
        return [
            InstancePreprocessUnfreezeRecipeStepsStatus.from_proto(i) for i in resources
        ]


class InstancePreprocessUnfreezeRecipeStepsStatusDetails(object):
    def __init__(self, type_url: str = None, value: str = None):
        self.type_url = type_url
        self.value = value

    @classmethod
    def to_proto(self, resource):
        if not resource:
            return None

        res = (
            instance_pb2.Tier2AlphaInstancePreprocessUnfreezeRecipeStepsStatusDetails()
        )
        if Primitive.to_proto(resource.type_url):
            res.type_url = Primitive.to_proto(resource.type_url)
        if Primitive.to_proto(resource.value):
            res.value = Primitive.to_proto(resource.value)
        return res

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return None

        return InstancePreprocessUnfreezeRecipeStepsStatusDetails(
            type_url=Primitive.from_proto(resource.type_url),
            value=Primitive.from_proto(resource.value),
        )


class InstancePreprocessUnfreezeRecipeStepsStatusDetailsArray(object):
    @classmethod
    def to_proto(self, resources):
        if not resources:
            return resources
        return [
            InstancePreprocessUnfreezeRecipeStepsStatusDetails.to_proto(i)
            for i in resources
        ]

    @classmethod
    def from_proto(self, resources):
        return [
            InstancePreprocessUnfreezeRecipeStepsStatusDetails.from_proto(i)
            for i in resources
        ]


class InstancePreprocessUnfreezeRecipeStepsQuotaRequestDeltas(object):
    def __init__(
        self,
        metric_name: str = None,
        amount: int = None,
        quota_location_name: str = None,
    ):
        self.metric_name = metric_name
        self.amount = amount
        self.quota_location_name = quota_location_name

    @classmethod
    def to_proto(self, resource):
        if not resource:
            return None

        res = (
            instance_pb2.Tier2AlphaInstancePreprocessUnfreezeRecipeStepsQuotaRequestDeltas()
        )
        if Primitive.to_proto(resource.metric_name):
            res.metric_name = Primitive.to_proto(resource.metric_name)
        if Primitive.to_proto(resource.amount):
            res.amount = Primitive.to_proto(resource.amount)
        if Primitive.to_proto(resource.quota_location_name):
            res.quota_location_name = Primitive.to_proto(resource.quota_location_name)
        return res

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return None

        return InstancePreprocessUnfreezeRecipeStepsQuotaRequestDeltas(
            metric_name=Primitive.from_proto(resource.metric_name),
            amount=Primitive.from_proto(resource.amount),
            quota_location_name=Primitive.from_proto(resource.quota_location_name),
        )


class InstancePreprocessUnfreezeRecipeStepsQuotaRequestDeltasArray(object):
    @classmethod
    def to_proto(self, resources):
        if not resources:
            return resources
        return [
            InstancePreprocessUnfreezeRecipeStepsQuotaRequestDeltas.to_proto(i)
            for i in resources
        ]

    @classmethod
    def from_proto(self, resources):
        return [
            InstancePreprocessUnfreezeRecipeStepsQuotaRequestDeltas.from_proto(i)
            for i in resources
        ]


class InstancePreprocessUnfreezeRecipeStepsPreprocessUpdate(object):
    def __init__(
        self, latency_slo_bucket_name: str = None, public_operation_metadata: str = None
    ):
        self.latency_slo_bucket_name = latency_slo_bucket_name
        self.public_operation_metadata = public_operation_metadata

    @classmethod
    def to_proto(self, resource):
        if not resource:
            return None

        res = (
            instance_pb2.Tier2AlphaInstancePreprocessUnfreezeRecipeStepsPreprocessUpdate()
        )
        if Primitive.to_proto(resource.latency_slo_bucket_name):
            res.latency_slo_bucket_name = Primitive.to_proto(
                resource.latency_slo_bucket_name
            )
        if Primitive.to_proto(resource.public_operation_metadata):
            res.public_operation_metadata = Primitive.to_proto(
                resource.public_operation_metadata
            )
        return res

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return None

        return InstancePreprocessUnfreezeRecipeStepsPreprocessUpdate(
            latency_slo_bucket_name=Primitive.from_proto(
                resource.latency_slo_bucket_name
            ),
            public_operation_metadata=Primitive.from_proto(
                resource.public_operation_metadata
            ),
        )


class InstancePreprocessUnfreezeRecipeStepsPreprocessUpdateArray(object):
    @classmethod
    def to_proto(self, resources):
        if not resources:
            return resources
        return [
            InstancePreprocessUnfreezeRecipeStepsPreprocessUpdate.to_proto(i)
            for i in resources
        ]

    @classmethod
    def from_proto(self, resources):
        return [
            InstancePreprocessUnfreezeRecipeStepsPreprocessUpdate.from_proto(i)
            for i in resources
        ]


class InstancePreprocessUnfreezeRecipeStepsRequestedTenantProject(object):
    def __init__(self, tag: str = None, folder: str = None, scope: str = None):
        self.tag = tag
        self.folder = folder
        self.scope = scope

    @classmethod
    def to_proto(self, resource):
        if not resource:
            return None

        res = (
            instance_pb2.Tier2AlphaInstancePreprocessUnfreezeRecipeStepsRequestedTenantProject()
        )
        if Primitive.to_proto(resource.tag):
            res.tag = Primitive.to_proto(resource.tag)
        if Primitive.to_proto(resource.folder):
            res.folder = Primitive.to_proto(resource.folder)
        if InstancePreprocessUnfreezeRecipeStepsRequestedTenantProjectScopeEnum.to_proto(
            resource.scope
        ):
            res.scope = InstancePreprocessUnfreezeRecipeStepsRequestedTenantProjectScopeEnum.to_proto(
                resource.scope
            )
        return res

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return None

        return InstancePreprocessUnfreezeRecipeStepsRequestedTenantProject(
            tag=Primitive.from_proto(resource.tag),
            folder=Primitive.from_proto(resource.folder),
            scope=InstancePreprocessUnfreezeRecipeStepsRequestedTenantProjectScopeEnum.from_proto(
                resource.scope
            ),
        )


class InstancePreprocessUnfreezeRecipeStepsRequestedTenantProjectArray(object):
    @classmethod
    def to_proto(self, resources):
        if not resources:
            return resources
        return [
            InstancePreprocessUnfreezeRecipeStepsRequestedTenantProject.to_proto(i)
            for i in resources
        ]

    @classmethod
    def from_proto(self, resources):
        return [
            InstancePreprocessUnfreezeRecipeStepsRequestedTenantProject.from_proto(i)
            for i in resources
        ]


class InstancePreprocessUnfreezeRecipeStepsPermissionsInfo(object):
    def __init__(
        self,
        policy_name: dict = None,
        iam_permissions: list = None,
        resource_path: str = None,
        api_attrs: dict = None,
        policy_name_mode: str = None,
        resource: dict = None,
    ):
        self.policy_name = policy_name
        self.iam_permissions = iam_permissions
        self.resource_path = resource_path
        self.api_attrs = api_attrs
        self.policy_name_mode = policy_name_mode
        self.resource = resource

    @classmethod
    def to_proto(self, resource):
        if not resource:
            return None

        res = (
            instance_pb2.Tier2AlphaInstancePreprocessUnfreezeRecipeStepsPermissionsInfo()
        )
        if InstancePreprocessUnfreezeRecipeStepsPermissionsInfoPolicyName.to_proto(
            resource.policy_name
        ):
            res.policy_name.CopyFrom(
                InstancePreprocessUnfreezeRecipeStepsPermissionsInfoPolicyName.to_proto(
                    resource.policy_name
                )
            )
        else:
            res.ClearField("policy_name")
        if InstancePreprocessUnfreezeRecipeStepsPermissionsInfoIamPermissionsArray.to_proto(
            resource.iam_permissions
        ):
            res.iam_permissions.extend(
                InstancePreprocessUnfreezeRecipeStepsPermissionsInfoIamPermissionsArray.to_proto(
                    resource.iam_permissions
                )
            )
        if Primitive.to_proto(resource.resource_path):
            res.resource_path = Primitive.to_proto(resource.resource_path)
        if InstanceGoogleprotobufstruct.to_proto(resource.api_attrs):
            res.api_attrs.CopyFrom(
                InstanceGoogleprotobufstruct.to_proto(resource.api_attrs)
            )
        else:
            res.ClearField("api_attrs")
        if InstancePreprocessUnfreezeRecipeStepsPermissionsInfoPolicyNameModeEnum.to_proto(
            resource.policy_name_mode
        ):
            res.policy_name_mode = InstancePreprocessUnfreezeRecipeStepsPermissionsInfoPolicyNameModeEnum.to_proto(
                resource.policy_name_mode
            )
        if InstancePreprocessUnfreezeRecipeStepsPermissionsInfoResource.to_proto(
            resource.resource
        ):
            res.resource.CopyFrom(
                InstancePreprocessUnfreezeRecipeStepsPermissionsInfoResource.to_proto(
                    resource.resource
                )
            )
        else:
            res.ClearField("resource")
        return res

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return None

        return InstancePreprocessUnfreezeRecipeStepsPermissionsInfo(
            policy_name=InstancePreprocessUnfreezeRecipeStepsPermissionsInfoPolicyName.from_proto(
                resource.policy_name
            ),
            iam_permissions=InstancePreprocessUnfreezeRecipeStepsPermissionsInfoIamPermissionsArray.from_proto(
                resource.iam_permissions
            ),
            resource_path=Primitive.from_proto(resource.resource_path),
            api_attrs=InstanceGoogleprotobufstruct.from_proto(resource.api_attrs),
            policy_name_mode=InstancePreprocessUnfreezeRecipeStepsPermissionsInfoPolicyNameModeEnum.from_proto(
                resource.policy_name_mode
            ),
            resource=InstancePreprocessUnfreezeRecipeStepsPermissionsInfoResource.from_proto(
                resource.resource
            ),
        )


class InstancePreprocessUnfreezeRecipeStepsPermissionsInfoArray(object):
    @classmethod
    def to_proto(self, resources):
        if not resources:
            return resources
        return [
            InstancePreprocessUnfreezeRecipeStepsPermissionsInfo.to_proto(i)
            for i in resources
        ]

    @classmethod
    def from_proto(self, resources):
        return [
            InstancePreprocessUnfreezeRecipeStepsPermissionsInfo.from_proto(i)
            for i in resources
        ]


class InstancePreprocessUnfreezeRecipeStepsPermissionsInfoPolicyName(object):
    def __init__(self, type: str = None, id: str = None, region: str = None):
        self.type = type
        self.id = id
        self.region = region

    @classmethod
    def to_proto(self, resource):
        if not resource:
            return None

        res = (
            instance_pb2.Tier2AlphaInstancePreprocessUnfreezeRecipeStepsPermissionsInfoPolicyName()
        )
        if Primitive.to_proto(resource.type):
            res.type = Primitive.to_proto(resource.type)
        if Primitive.to_proto(resource.id):
            res.id = Primitive.to_proto(resource.id)
        if Primitive.to_proto(resource.region):
            res.region = Primitive.to_proto(resource.region)
        return res

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return None

        return InstancePreprocessUnfreezeRecipeStepsPermissionsInfoPolicyName(
            type=Primitive.from_proto(resource.type),
            id=Primitive.from_proto(resource.id),
            region=Primitive.from_proto(resource.region),
        )


class InstancePreprocessUnfreezeRecipeStepsPermissionsInfoPolicyNameArray(object):
    @classmethod
    def to_proto(self, resources):
        if not resources:
            return resources
        return [
            InstancePreprocessUnfreezeRecipeStepsPermissionsInfoPolicyName.to_proto(i)
            for i in resources
        ]

    @classmethod
    def from_proto(self, resources):
        return [
            InstancePreprocessUnfreezeRecipeStepsPermissionsInfoPolicyName.from_proto(i)
            for i in resources
        ]


class InstancePreprocessUnfreezeRecipeStepsPermissionsInfoIamPermissions(object):
    def __init__(self, permission: str = None):
        self.permission = permission

    @classmethod
    def to_proto(self, resource):
        if not resource:
            return None

        res = (
            instance_pb2.Tier2AlphaInstancePreprocessUnfreezeRecipeStepsPermissionsInfoIamPermissions()
        )
        if Primitive.to_proto(resource.permission):
            res.permission = Primitive.to_proto(resource.permission)
        return res

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return None

        return InstancePreprocessUnfreezeRecipeStepsPermissionsInfoIamPermissions(
            permission=Primitive.from_proto(resource.permission),
        )


class InstancePreprocessUnfreezeRecipeStepsPermissionsInfoIamPermissionsArray(object):
    @classmethod
    def to_proto(self, resources):
        if not resources:
            return resources
        return [
            InstancePreprocessUnfreezeRecipeStepsPermissionsInfoIamPermissions.to_proto(
                i
            )
            for i in resources
        ]

    @classmethod
    def from_proto(self, resources):
        return [
            InstancePreprocessUnfreezeRecipeStepsPermissionsInfoIamPermissions.from_proto(
                i
            )
            for i in resources
        ]


class InstancePreprocessUnfreezeRecipeStepsPermissionsInfoResource(object):
    def __init__(
        self,
        name: str = None,
        type: str = None,
        service: str = None,
        labels: dict = None,
    ):
        self.name = name
        self.type = type
        self.service = service
        self.labels = labels

    @classmethod
    def to_proto(self, resource):
        if not resource:
            return None

        res = (
            instance_pb2.Tier2AlphaInstancePreprocessUnfreezeRecipeStepsPermissionsInfoResource()
        )
        if Primitive.to_proto(resource.name):
            res.name = Primitive.to_proto(resource.name)
        if Primitive.to_proto(resource.type):
            res.type = Primitive.to_proto(resource.type)
        if Primitive.to_proto(resource.service):
            res.service = Primitive.to_proto(resource.service)
        if Primitive.to_proto(resource.labels):
            res.labels = Primitive.to_proto(resource.labels)
        return res

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return None

        return InstancePreprocessUnfreezeRecipeStepsPermissionsInfoResource(
            name=Primitive.from_proto(resource.name),
            type=Primitive.from_proto(resource.type),
            service=Primitive.from_proto(resource.service),
            labels=Primitive.from_proto(resource.labels),
        )


class InstancePreprocessUnfreezeRecipeStepsPermissionsInfoResourceArray(object):
    @classmethod
    def to_proto(self, resources):
        if not resources:
            return resources
        return [
            InstancePreprocessUnfreezeRecipeStepsPermissionsInfoResource.to_proto(i)
            for i in resources
        ]

    @classmethod
    def from_proto(self, resources):
        return [
            InstancePreprocessUnfreezeRecipeStepsPermissionsInfoResource.from_proto(i)
            for i in resources
        ]


class InstancePreprocessUnfreezeRecipeStepsKeyNotificationsUpdate(object):
    def __init__(self, key_notifications_info: dict = None):
        self.key_notifications_info = key_notifications_info

    @classmethod
    def to_proto(self, resource):
        if not resource:
            return None

        res = (
            instance_pb2.Tier2AlphaInstancePreprocessUnfreezeRecipeStepsKeyNotificationsUpdate()
        )
        if InstancePreprocessUnfreezeRecipeStepsKeyNotificationsUpdateKeyNotificationsInfo.to_proto(
            resource.key_notifications_info
        ):
            res.key_notifications_info.CopyFrom(
                InstancePreprocessUnfreezeRecipeStepsKeyNotificationsUpdateKeyNotificationsInfo.to_proto(
                    resource.key_notifications_info
                )
            )
        else:
            res.ClearField("key_notifications_info")
        return res

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return None

        return InstancePreprocessUnfreezeRecipeStepsKeyNotificationsUpdate(
            key_notifications_info=InstancePreprocessUnfreezeRecipeStepsKeyNotificationsUpdateKeyNotificationsInfo.from_proto(
                resource.key_notifications_info
            ),
        )


class InstancePreprocessUnfreezeRecipeStepsKeyNotificationsUpdateArray(object):
    @classmethod
    def to_proto(self, resources):
        if not resources:
            return resources
        return [
            InstancePreprocessUnfreezeRecipeStepsKeyNotificationsUpdate.to_proto(i)
            for i in resources
        ]

    @classmethod
    def from_proto(self, resources):
        return [
            InstancePreprocessUnfreezeRecipeStepsKeyNotificationsUpdate.from_proto(i)
            for i in resources
        ]


class InstancePreprocessUnfreezeRecipeStepsKeyNotificationsUpdateKeyNotificationsInfo(
    object
):
    def __init__(
        self,
        data_version: int = None,
        delegate: str = None,
        key_notification_configs: list = None,
    ):
        self.data_version = data_version
        self.delegate = delegate
        self.key_notification_configs = key_notification_configs

    @classmethod
    def to_proto(self, resource):
        if not resource:
            return None

        res = (
            instance_pb2.Tier2AlphaInstancePreprocessUnfreezeRecipeStepsKeyNotificationsUpdateKeyNotificationsInfo()
        )
        if Primitive.to_proto(resource.data_version):
            res.data_version = Primitive.to_proto(resource.data_version)
        if Primitive.to_proto(resource.delegate):
            res.delegate = Primitive.to_proto(resource.delegate)
        if InstancePreprocessUnfreezeRecipeStepsKeyNotificationsUpdateKeyNotificationsInfoKeyNotificationConfigsArray.to_proto(
            resource.key_notification_configs
        ):
            res.key_notification_configs.extend(
                InstancePreprocessUnfreezeRecipeStepsKeyNotificationsUpdateKeyNotificationsInfoKeyNotificationConfigsArray.to_proto(
                    resource.key_notification_configs
                )
            )
        return res

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return None

        return InstancePreprocessUnfreezeRecipeStepsKeyNotificationsUpdateKeyNotificationsInfo(
            data_version=Primitive.from_proto(resource.data_version),
            delegate=Primitive.from_proto(resource.delegate),
            key_notification_configs=InstancePreprocessUnfreezeRecipeStepsKeyNotificationsUpdateKeyNotificationsInfoKeyNotificationConfigsArray.from_proto(
                resource.key_notification_configs
            ),
        )


class InstancePreprocessUnfreezeRecipeStepsKeyNotificationsUpdateKeyNotificationsInfoArray(
    object
):
    @classmethod
    def to_proto(self, resources):
        if not resources:
            return resources
        return [
            InstancePreprocessUnfreezeRecipeStepsKeyNotificationsUpdateKeyNotificationsInfo.to_proto(
                i
            )
            for i in resources
        ]

    @classmethod
    def from_proto(self, resources):
        return [
            InstancePreprocessUnfreezeRecipeStepsKeyNotificationsUpdateKeyNotificationsInfo.from_proto(
                i
            )
            for i in resources
        ]


class InstancePreprocessUnfreezeRecipeStepsKeyNotificationsUpdateKeyNotificationsInfoKeyNotificationConfigs(
    object
):
    def __init__(
        self,
        key_or_version_name: str = None,
        grant: str = None,
        delegator_gaia_id: int = None,
    ):
        self.key_or_version_name = key_or_version_name
        self.grant = grant
        self.delegator_gaia_id = delegator_gaia_id

    @classmethod
    def to_proto(self, resource):
        if not resource:
            return None

        res = (
            instance_pb2.Tier2AlphaInstancePreprocessUnfreezeRecipeStepsKeyNotificationsUpdateKeyNotificationsInfoKeyNotificationConfigs()
        )
        if Primitive.to_proto(resource.key_or_version_name):
            res.key_or_version_name = Primitive.to_proto(resource.key_or_version_name)
        if Primitive.to_proto(resource.grant):
            res.grant = Primitive.to_proto(resource.grant)
        if Primitive.to_proto(resource.delegator_gaia_id):
            res.delegator_gaia_id = Primitive.to_proto(resource.delegator_gaia_id)
        return res

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return None

        return InstancePreprocessUnfreezeRecipeStepsKeyNotificationsUpdateKeyNotificationsInfoKeyNotificationConfigs(
            key_or_version_name=Primitive.from_proto(resource.key_or_version_name),
            grant=Primitive.from_proto(resource.grant),
            delegator_gaia_id=Primitive.from_proto(resource.delegator_gaia_id),
        )


class InstancePreprocessUnfreezeRecipeStepsKeyNotificationsUpdateKeyNotificationsInfoKeyNotificationConfigsArray(
    object
):
    @classmethod
    def to_proto(self, resources):
        if not resources:
            return resources
        return [
            InstancePreprocessUnfreezeRecipeStepsKeyNotificationsUpdateKeyNotificationsInfoKeyNotificationConfigs.to_proto(
                i
            )
            for i in resources
        ]

    @classmethod
    def from_proto(self, resources):
        return [
            InstancePreprocessUnfreezeRecipeStepsKeyNotificationsUpdateKeyNotificationsInfoKeyNotificationConfigs.from_proto(
                i
            )
            for i in resources
        ]


class InstanceUnfreezeRecipe(object):
    def __init__(
        self,
        steps: list = None,
        honor_cancel_request: bool = None,
        ignore_recipe_after: int = None,
        verify_deadline_seconds_below: float = None,
        populate_operation_result: bool = None,
        readonly_recipe_start_time: str = None,
        resource_names_stored_in_clh_with_delay: list = None,
        delay_to_store_resources_in_clh_db_nanos: int = None,
    ):
        self.steps = steps
        self.honor_cancel_request = honor_cancel_request
        self.ignore_recipe_after = ignore_recipe_after
        self.verify_deadline_seconds_below = verify_deadline_seconds_below
        self.populate_operation_result = populate_operation_result
        self.readonly_recipe_start_time = readonly_recipe_start_time
        self.resource_names_stored_in_clh_with_delay = (
            resource_names_stored_in_clh_with_delay
        )
        self.delay_to_store_resources_in_clh_db_nanos = (
            delay_to_store_resources_in_clh_db_nanos
        )

    @classmethod
    def to_proto(self, resource):
        if not resource:
            return None

        res = instance_pb2.Tier2AlphaInstanceUnfreezeRecipe()
        if InstanceUnfreezeRecipeStepsArray.to_proto(resource.steps):
            res.steps.extend(InstanceUnfreezeRecipeStepsArray.to_proto(resource.steps))
        if Primitive.to_proto(resource.honor_cancel_request):
            res.honor_cancel_request = Primitive.to_proto(resource.honor_cancel_request)
        if Primitive.to_proto(resource.ignore_recipe_after):
            res.ignore_recipe_after = Primitive.to_proto(resource.ignore_recipe_after)
        if Primitive.to_proto(resource.verify_deadline_seconds_below):
            res.verify_deadline_seconds_below = Primitive.to_proto(
                resource.verify_deadline_seconds_below
            )
        if Primitive.to_proto(resource.populate_operation_result):
            res.populate_operation_result = Primitive.to_proto(
                resource.populate_operation_result
            )
        if Primitive.to_proto(resource.readonly_recipe_start_time):
            res.readonly_recipe_start_time = Primitive.to_proto(
                resource.readonly_recipe_start_time
            )
        if Primitive.to_proto(resource.resource_names_stored_in_clh_with_delay):
            res.resource_names_stored_in_clh_with_delay.extend(
                Primitive.to_proto(resource.resource_names_stored_in_clh_with_delay)
            )
        if Primitive.to_proto(resource.delay_to_store_resources_in_clh_db_nanos):
            res.delay_to_store_resources_in_clh_db_nanos = Primitive.to_proto(
                resource.delay_to_store_resources_in_clh_db_nanos
            )
        return res

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return None

        return InstanceUnfreezeRecipe(
            steps=InstanceUnfreezeRecipeStepsArray.from_proto(resource.steps),
            honor_cancel_request=Primitive.from_proto(resource.honor_cancel_request),
            ignore_recipe_after=Primitive.from_proto(resource.ignore_recipe_after),
            verify_deadline_seconds_below=Primitive.from_proto(
                resource.verify_deadline_seconds_below
            ),
            populate_operation_result=Primitive.from_proto(
                resource.populate_operation_result
            ),
            readonly_recipe_start_time=Primitive.from_proto(
                resource.readonly_recipe_start_time
            ),
            resource_names_stored_in_clh_with_delay=Primitive.from_proto(
                resource.resource_names_stored_in_clh_with_delay
            ),
            delay_to_store_resources_in_clh_db_nanos=Primitive.from_proto(
                resource.delay_to_store_resources_in_clh_db_nanos
            ),
        )


class InstanceUnfreezeRecipeArray(object):
    @classmethod
    def to_proto(self, resources):
        if not resources:
            return resources
        return [InstanceUnfreezeRecipe.to_proto(i) for i in resources]

    @classmethod
    def from_proto(self, resources):
        return [InstanceUnfreezeRecipe.from_proto(i) for i in resources]


class InstanceUnfreezeRecipeSteps(object):
    def __init__(
        self,
        relative_time: int = None,
        sleep_duration: int = None,
        action: str = None,
        status: dict = None,
        error_space: str = None,
        p4_service_account: str = None,
        resource_metadata_size: int = None,
        description: str = None,
        updated_repeat_operation_delay_sec: float = None,
        quota_request_deltas: list = None,
        preprocess_update: dict = None,
        public_operation_metadata: str = None,
        requested_tenant_project: dict = None,
        permissions_info: list = None,
        key_notifications_update: dict = None,
        clh_data_update_time: str = None,
        public_error_message: str = None,
    ):
        self.relative_time = relative_time
        self.sleep_duration = sleep_duration
        self.action = action
        self.status = status
        self.error_space = error_space
        self.p4_service_account = p4_service_account
        self.resource_metadata_size = resource_metadata_size
        self.description = description
        self.updated_repeat_operation_delay_sec = updated_repeat_operation_delay_sec
        self.quota_request_deltas = quota_request_deltas
        self.preprocess_update = preprocess_update
        self.public_operation_metadata = public_operation_metadata
        self.requested_tenant_project = requested_tenant_project
        self.permissions_info = permissions_info
        self.key_notifications_update = key_notifications_update
        self.clh_data_update_time = clh_data_update_time
        self.public_error_message = public_error_message

    @classmethod
    def to_proto(self, resource):
        if not resource:
            return None

        res = instance_pb2.Tier2AlphaInstanceUnfreezeRecipeSteps()
        if Primitive.to_proto(resource.relative_time):
            res.relative_time = Primitive.to_proto(resource.relative_time)
        if Primitive.to_proto(resource.sleep_duration):
            res.sleep_duration = Primitive.to_proto(resource.sleep_duration)
        if InstanceUnfreezeRecipeStepsActionEnum.to_proto(resource.action):
            res.action = InstanceUnfreezeRecipeStepsActionEnum.to_proto(resource.action)
        if InstanceUnfreezeRecipeStepsStatus.to_proto(resource.status):
            res.status.CopyFrom(
                InstanceUnfreezeRecipeStepsStatus.to_proto(resource.status)
            )
        else:
            res.ClearField("status")
        if Primitive.to_proto(resource.error_space):
            res.error_space = Primitive.to_proto(resource.error_space)
        if Primitive.to_proto(resource.p4_service_account):
            res.p4_service_account = Primitive.to_proto(resource.p4_service_account)
        if Primitive.to_proto(resource.resource_metadata_size):
            res.resource_metadata_size = Primitive.to_proto(
                resource.resource_metadata_size
            )
        if Primitive.to_proto(resource.description):
            res.description = Primitive.to_proto(resource.description)
        if Primitive.to_proto(resource.updated_repeat_operation_delay_sec):
            res.updated_repeat_operation_delay_sec = Primitive.to_proto(
                resource.updated_repeat_operation_delay_sec
            )
        if InstanceUnfreezeRecipeStepsQuotaRequestDeltasArray.to_proto(
            resource.quota_request_deltas
        ):
            res.quota_request_deltas.extend(
                InstanceUnfreezeRecipeStepsQuotaRequestDeltasArray.to_proto(
                    resource.quota_request_deltas
                )
            )
        if InstanceUnfreezeRecipeStepsPreprocessUpdate.to_proto(
            resource.preprocess_update
        ):
            res.preprocess_update.CopyFrom(
                InstanceUnfreezeRecipeStepsPreprocessUpdate.to_proto(
                    resource.preprocess_update
                )
            )
        else:
            res.ClearField("preprocess_update")
        if Primitive.to_proto(resource.public_operation_metadata):
            res.public_operation_metadata = Primitive.to_proto(
                resource.public_operation_metadata
            )
        if InstanceUnfreezeRecipeStepsRequestedTenantProject.to_proto(
            resource.requested_tenant_project
        ):
            res.requested_tenant_project.CopyFrom(
                InstanceUnfreezeRecipeStepsRequestedTenantProject.to_proto(
                    resource.requested_tenant_project
                )
            )
        else:
            res.ClearField("requested_tenant_project")
        if InstanceUnfreezeRecipeStepsPermissionsInfoArray.to_proto(
            resource.permissions_info
        ):
            res.permissions_info.extend(
                InstanceUnfreezeRecipeStepsPermissionsInfoArray.to_proto(
                    resource.permissions_info
                )
            )
        if InstanceUnfreezeRecipeStepsKeyNotificationsUpdate.to_proto(
            resource.key_notifications_update
        ):
            res.key_notifications_update.CopyFrom(
                InstanceUnfreezeRecipeStepsKeyNotificationsUpdate.to_proto(
                    resource.key_notifications_update
                )
            )
        else:
            res.ClearField("key_notifications_update")
        if Primitive.to_proto(resource.clh_data_update_time):
            res.clh_data_update_time = Primitive.to_proto(resource.clh_data_update_time)
        if Primitive.to_proto(resource.public_error_message):
            res.public_error_message = Primitive.to_proto(resource.public_error_message)
        return res

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return None

        return InstanceUnfreezeRecipeSteps(
            relative_time=Primitive.from_proto(resource.relative_time),
            sleep_duration=Primitive.from_proto(resource.sleep_duration),
            action=InstanceUnfreezeRecipeStepsActionEnum.from_proto(resource.action),
            status=InstanceUnfreezeRecipeStepsStatus.from_proto(resource.status),
            error_space=Primitive.from_proto(resource.error_space),
            p4_service_account=Primitive.from_proto(resource.p4_service_account),
            resource_metadata_size=Primitive.from_proto(
                resource.resource_metadata_size
            ),
            description=Primitive.from_proto(resource.description),
            updated_repeat_operation_delay_sec=Primitive.from_proto(
                resource.updated_repeat_operation_delay_sec
            ),
            quota_request_deltas=InstanceUnfreezeRecipeStepsQuotaRequestDeltasArray.from_proto(
                resource.quota_request_deltas
            ),
            preprocess_update=InstanceUnfreezeRecipeStepsPreprocessUpdate.from_proto(
                resource.preprocess_update
            ),
            public_operation_metadata=Primitive.from_proto(
                resource.public_operation_metadata
            ),
            requested_tenant_project=InstanceUnfreezeRecipeStepsRequestedTenantProject.from_proto(
                resource.requested_tenant_project
            ),
            permissions_info=InstanceUnfreezeRecipeStepsPermissionsInfoArray.from_proto(
                resource.permissions_info
            ),
            key_notifications_update=InstanceUnfreezeRecipeStepsKeyNotificationsUpdate.from_proto(
                resource.key_notifications_update
            ),
            clh_data_update_time=Primitive.from_proto(resource.clh_data_update_time),
            public_error_message=Primitive.from_proto(resource.public_error_message),
        )


class InstanceUnfreezeRecipeStepsArray(object):
    @classmethod
    def to_proto(self, resources):
        if not resources:
            return resources
        return [InstanceUnfreezeRecipeSteps.to_proto(i) for i in resources]

    @classmethod
    def from_proto(self, resources):
        return [InstanceUnfreezeRecipeSteps.from_proto(i) for i in resources]


class InstanceUnfreezeRecipeStepsStatus(object):
    def __init__(self, code: int = None, message: str = None, details: list = None):
        self.code = code
        self.message = message
        self.details = details

    @classmethod
    def to_proto(self, resource):
        if not resource:
            return None

        res = instance_pb2.Tier2AlphaInstanceUnfreezeRecipeStepsStatus()
        if Primitive.to_proto(resource.code):
            res.code = Primitive.to_proto(resource.code)
        if Primitive.to_proto(resource.message):
            res.message = Primitive.to_proto(resource.message)
        if InstanceUnfreezeRecipeStepsStatusDetailsArray.to_proto(resource.details):
            res.details.extend(
                InstanceUnfreezeRecipeStepsStatusDetailsArray.to_proto(resource.details)
            )
        return res

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return None

        return InstanceUnfreezeRecipeStepsStatus(
            code=Primitive.from_proto(resource.code),
            message=Primitive.from_proto(resource.message),
            details=InstanceUnfreezeRecipeStepsStatusDetailsArray.from_proto(
                resource.details
            ),
        )


class InstanceUnfreezeRecipeStepsStatusArray(object):
    @classmethod
    def to_proto(self, resources):
        if not resources:
            return resources
        return [InstanceUnfreezeRecipeStepsStatus.to_proto(i) for i in resources]

    @classmethod
    def from_proto(self, resources):
        return [InstanceUnfreezeRecipeStepsStatus.from_proto(i) for i in resources]


class InstanceUnfreezeRecipeStepsStatusDetails(object):
    def __init__(self, type_url: str = None, value: str = None):
        self.type_url = type_url
        self.value = value

    @classmethod
    def to_proto(self, resource):
        if not resource:
            return None

        res = instance_pb2.Tier2AlphaInstanceUnfreezeRecipeStepsStatusDetails()
        if Primitive.to_proto(resource.type_url):
            res.type_url = Primitive.to_proto(resource.type_url)
        if Primitive.to_proto(resource.value):
            res.value = Primitive.to_proto(resource.value)
        return res

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return None

        return InstanceUnfreezeRecipeStepsStatusDetails(
            type_url=Primitive.from_proto(resource.type_url),
            value=Primitive.from_proto(resource.value),
        )


class InstanceUnfreezeRecipeStepsStatusDetailsArray(object):
    @classmethod
    def to_proto(self, resources):
        if not resources:
            return resources
        return [InstanceUnfreezeRecipeStepsStatusDetails.to_proto(i) for i in resources]

    @classmethod
    def from_proto(self, resources):
        return [
            InstanceUnfreezeRecipeStepsStatusDetails.from_proto(i) for i in resources
        ]


class InstanceUnfreezeRecipeStepsQuotaRequestDeltas(object):
    def __init__(
        self,
        metric_name: str = None,
        amount: int = None,
        quota_location_name: str = None,
    ):
        self.metric_name = metric_name
        self.amount = amount
        self.quota_location_name = quota_location_name

    @classmethod
    def to_proto(self, resource):
        if not resource:
            return None

        res = instance_pb2.Tier2AlphaInstanceUnfreezeRecipeStepsQuotaRequestDeltas()
        if Primitive.to_proto(resource.metric_name):
            res.metric_name = Primitive.to_proto(resource.metric_name)
        if Primitive.to_proto(resource.amount):
            res.amount = Primitive.to_proto(resource.amount)
        if Primitive.to_proto(resource.quota_location_name):
            res.quota_location_name = Primitive.to_proto(resource.quota_location_name)
        return res

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return None

        return InstanceUnfreezeRecipeStepsQuotaRequestDeltas(
            metric_name=Primitive.from_proto(resource.metric_name),
            amount=Primitive.from_proto(resource.amount),
            quota_location_name=Primitive.from_proto(resource.quota_location_name),
        )


class InstanceUnfreezeRecipeStepsQuotaRequestDeltasArray(object):
    @classmethod
    def to_proto(self, resources):
        if not resources:
            return resources
        return [
            InstanceUnfreezeRecipeStepsQuotaRequestDeltas.to_proto(i) for i in resources
        ]

    @classmethod
    def from_proto(self, resources):
        return [
            InstanceUnfreezeRecipeStepsQuotaRequestDeltas.from_proto(i)
            for i in resources
        ]


class InstanceUnfreezeRecipeStepsPreprocessUpdate(object):
    def __init__(
        self, latency_slo_bucket_name: str = None, public_operation_metadata: str = None
    ):
        self.latency_slo_bucket_name = latency_slo_bucket_name
        self.public_operation_metadata = public_operation_metadata

    @classmethod
    def to_proto(self, resource):
        if not resource:
            return None

        res = instance_pb2.Tier2AlphaInstanceUnfreezeRecipeStepsPreprocessUpdate()
        if Primitive.to_proto(resource.latency_slo_bucket_name):
            res.latency_slo_bucket_name = Primitive.to_proto(
                resource.latency_slo_bucket_name
            )
        if Primitive.to_proto(resource.public_operation_metadata):
            res.public_operation_metadata = Primitive.to_proto(
                resource.public_operation_metadata
            )
        return res

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return None

        return InstanceUnfreezeRecipeStepsPreprocessUpdate(
            latency_slo_bucket_name=Primitive.from_proto(
                resource.latency_slo_bucket_name
            ),
            public_operation_metadata=Primitive.from_proto(
                resource.public_operation_metadata
            ),
        )


class InstanceUnfreezeRecipeStepsPreprocessUpdateArray(object):
    @classmethod
    def to_proto(self, resources):
        if not resources:
            return resources
        return [
            InstanceUnfreezeRecipeStepsPreprocessUpdate.to_proto(i) for i in resources
        ]

    @classmethod
    def from_proto(self, resources):
        return [
            InstanceUnfreezeRecipeStepsPreprocessUpdate.from_proto(i) for i in resources
        ]


class InstanceUnfreezeRecipeStepsRequestedTenantProject(object):
    def __init__(self, tag: str = None, folder: str = None, scope: str = None):
        self.tag = tag
        self.folder = folder
        self.scope = scope

    @classmethod
    def to_proto(self, resource):
        if not resource:
            return None

        res = instance_pb2.Tier2AlphaInstanceUnfreezeRecipeStepsRequestedTenantProject()
        if Primitive.to_proto(resource.tag):
            res.tag = Primitive.to_proto(resource.tag)
        if Primitive.to_proto(resource.folder):
            res.folder = Primitive.to_proto(resource.folder)
        if InstanceUnfreezeRecipeStepsRequestedTenantProjectScopeEnum.to_proto(
            resource.scope
        ):
            res.scope = InstanceUnfreezeRecipeStepsRequestedTenantProjectScopeEnum.to_proto(
                resource.scope
            )
        return res

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return None

        return InstanceUnfreezeRecipeStepsRequestedTenantProject(
            tag=Primitive.from_proto(resource.tag),
            folder=Primitive.from_proto(resource.folder),
            scope=InstanceUnfreezeRecipeStepsRequestedTenantProjectScopeEnum.from_proto(
                resource.scope
            ),
        )


class InstanceUnfreezeRecipeStepsRequestedTenantProjectArray(object):
    @classmethod
    def to_proto(self, resources):
        if not resources:
            return resources
        return [
            InstanceUnfreezeRecipeStepsRequestedTenantProject.to_proto(i)
            for i in resources
        ]

    @classmethod
    def from_proto(self, resources):
        return [
            InstanceUnfreezeRecipeStepsRequestedTenantProject.from_proto(i)
            for i in resources
        ]


class InstanceUnfreezeRecipeStepsPermissionsInfo(object):
    def __init__(
        self,
        policy_name: dict = None,
        iam_permissions: list = None,
        resource_path: str = None,
        api_attrs: dict = None,
        policy_name_mode: str = None,
        resource: dict = None,
    ):
        self.policy_name = policy_name
        self.iam_permissions = iam_permissions
        self.resource_path = resource_path
        self.api_attrs = api_attrs
        self.policy_name_mode = policy_name_mode
        self.resource = resource

    @classmethod
    def to_proto(self, resource):
        if not resource:
            return None

        res = instance_pb2.Tier2AlphaInstanceUnfreezeRecipeStepsPermissionsInfo()
        if InstanceUnfreezeRecipeStepsPermissionsInfoPolicyName.to_proto(
            resource.policy_name
        ):
            res.policy_name.CopyFrom(
                InstanceUnfreezeRecipeStepsPermissionsInfoPolicyName.to_proto(
                    resource.policy_name
                )
            )
        else:
            res.ClearField("policy_name")
        if InstanceUnfreezeRecipeStepsPermissionsInfoIamPermissionsArray.to_proto(
            resource.iam_permissions
        ):
            res.iam_permissions.extend(
                InstanceUnfreezeRecipeStepsPermissionsInfoIamPermissionsArray.to_proto(
                    resource.iam_permissions
                )
            )
        if Primitive.to_proto(resource.resource_path):
            res.resource_path = Primitive.to_proto(resource.resource_path)
        if InstanceGoogleprotobufstruct.to_proto(resource.api_attrs):
            res.api_attrs.CopyFrom(
                InstanceGoogleprotobufstruct.to_proto(resource.api_attrs)
            )
        else:
            res.ClearField("api_attrs")
        if InstanceUnfreezeRecipeStepsPermissionsInfoPolicyNameModeEnum.to_proto(
            resource.policy_name_mode
        ):
            res.policy_name_mode = InstanceUnfreezeRecipeStepsPermissionsInfoPolicyNameModeEnum.to_proto(
                resource.policy_name_mode
            )
        if InstanceUnfreezeRecipeStepsPermissionsInfoResource.to_proto(
            resource.resource
        ):
            res.resource.CopyFrom(
                InstanceUnfreezeRecipeStepsPermissionsInfoResource.to_proto(
                    resource.resource
                )
            )
        else:
            res.ClearField("resource")
        return res

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return None

        return InstanceUnfreezeRecipeStepsPermissionsInfo(
            policy_name=InstanceUnfreezeRecipeStepsPermissionsInfoPolicyName.from_proto(
                resource.policy_name
            ),
            iam_permissions=InstanceUnfreezeRecipeStepsPermissionsInfoIamPermissionsArray.from_proto(
                resource.iam_permissions
            ),
            resource_path=Primitive.from_proto(resource.resource_path),
            api_attrs=InstanceGoogleprotobufstruct.from_proto(resource.api_attrs),
            policy_name_mode=InstanceUnfreezeRecipeStepsPermissionsInfoPolicyNameModeEnum.from_proto(
                resource.policy_name_mode
            ),
            resource=InstanceUnfreezeRecipeStepsPermissionsInfoResource.from_proto(
                resource.resource
            ),
        )


class InstanceUnfreezeRecipeStepsPermissionsInfoArray(object):
    @classmethod
    def to_proto(self, resources):
        if not resources:
            return resources
        return [
            InstanceUnfreezeRecipeStepsPermissionsInfo.to_proto(i) for i in resources
        ]

    @classmethod
    def from_proto(self, resources):
        return [
            InstanceUnfreezeRecipeStepsPermissionsInfo.from_proto(i) for i in resources
        ]


class InstanceUnfreezeRecipeStepsPermissionsInfoPolicyName(object):
    def __init__(self, type: str = None, id: str = None, region: str = None):
        self.type = type
        self.id = id
        self.region = region

    @classmethod
    def to_proto(self, resource):
        if not resource:
            return None

        res = (
            instance_pb2.Tier2AlphaInstanceUnfreezeRecipeStepsPermissionsInfoPolicyName()
        )
        if Primitive.to_proto(resource.type):
            res.type = Primitive.to_proto(resource.type)
        if Primitive.to_proto(resource.id):
            res.id = Primitive.to_proto(resource.id)
        if Primitive.to_proto(resource.region):
            res.region = Primitive.to_proto(resource.region)
        return res

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return None

        return InstanceUnfreezeRecipeStepsPermissionsInfoPolicyName(
            type=Primitive.from_proto(resource.type),
            id=Primitive.from_proto(resource.id),
            region=Primitive.from_proto(resource.region),
        )


class InstanceUnfreezeRecipeStepsPermissionsInfoPolicyNameArray(object):
    @classmethod
    def to_proto(self, resources):
        if not resources:
            return resources
        return [
            InstanceUnfreezeRecipeStepsPermissionsInfoPolicyName.to_proto(i)
            for i in resources
        ]

    @classmethod
    def from_proto(self, resources):
        return [
            InstanceUnfreezeRecipeStepsPermissionsInfoPolicyName.from_proto(i)
            for i in resources
        ]


class InstanceUnfreezeRecipeStepsPermissionsInfoIamPermissions(object):
    def __init__(self, permission: str = None):
        self.permission = permission

    @classmethod
    def to_proto(self, resource):
        if not resource:
            return None

        res = (
            instance_pb2.Tier2AlphaInstanceUnfreezeRecipeStepsPermissionsInfoIamPermissions()
        )
        if Primitive.to_proto(resource.permission):
            res.permission = Primitive.to_proto(resource.permission)
        return res

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return None

        return InstanceUnfreezeRecipeStepsPermissionsInfoIamPermissions(
            permission=Primitive.from_proto(resource.permission),
        )


class InstanceUnfreezeRecipeStepsPermissionsInfoIamPermissionsArray(object):
    @classmethod
    def to_proto(self, resources):
        if not resources:
            return resources
        return [
            InstanceUnfreezeRecipeStepsPermissionsInfoIamPermissions.to_proto(i)
            for i in resources
        ]

    @classmethod
    def from_proto(self, resources):
        return [
            InstanceUnfreezeRecipeStepsPermissionsInfoIamPermissions.from_proto(i)
            for i in resources
        ]


class InstanceUnfreezeRecipeStepsPermissionsInfoResource(object):
    def __init__(
        self,
        name: str = None,
        type: str = None,
        service: str = None,
        labels: dict = None,
    ):
        self.name = name
        self.type = type
        self.service = service
        self.labels = labels

    @classmethod
    def to_proto(self, resource):
        if not resource:
            return None

        res = (
            instance_pb2.Tier2AlphaInstanceUnfreezeRecipeStepsPermissionsInfoResource()
        )
        if Primitive.to_proto(resource.name):
            res.name = Primitive.to_proto(resource.name)
        if Primitive.to_proto(resource.type):
            res.type = Primitive.to_proto(resource.type)
        if Primitive.to_proto(resource.service):
            res.service = Primitive.to_proto(resource.service)
        if Primitive.to_proto(resource.labels):
            res.labels = Primitive.to_proto(resource.labels)
        return res

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return None

        return InstanceUnfreezeRecipeStepsPermissionsInfoResource(
            name=Primitive.from_proto(resource.name),
            type=Primitive.from_proto(resource.type),
            service=Primitive.from_proto(resource.service),
            labels=Primitive.from_proto(resource.labels),
        )


class InstanceUnfreezeRecipeStepsPermissionsInfoResourceArray(object):
    @classmethod
    def to_proto(self, resources):
        if not resources:
            return resources
        return [
            InstanceUnfreezeRecipeStepsPermissionsInfoResource.to_proto(i)
            for i in resources
        ]

    @classmethod
    def from_proto(self, resources):
        return [
            InstanceUnfreezeRecipeStepsPermissionsInfoResource.from_proto(i)
            for i in resources
        ]


class InstanceUnfreezeRecipeStepsKeyNotificationsUpdate(object):
    def __init__(self, key_notifications_info: dict = None):
        self.key_notifications_info = key_notifications_info

    @classmethod
    def to_proto(self, resource):
        if not resource:
            return None

        res = instance_pb2.Tier2AlphaInstanceUnfreezeRecipeStepsKeyNotificationsUpdate()
        if InstanceUnfreezeRecipeStepsKeyNotificationsUpdateKeyNotificationsInfo.to_proto(
            resource.key_notifications_info
        ):
            res.key_notifications_info.CopyFrom(
                InstanceUnfreezeRecipeStepsKeyNotificationsUpdateKeyNotificationsInfo.to_proto(
                    resource.key_notifications_info
                )
            )
        else:
            res.ClearField("key_notifications_info")
        return res

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return None

        return InstanceUnfreezeRecipeStepsKeyNotificationsUpdate(
            key_notifications_info=InstanceUnfreezeRecipeStepsKeyNotificationsUpdateKeyNotificationsInfo.from_proto(
                resource.key_notifications_info
            ),
        )


class InstanceUnfreezeRecipeStepsKeyNotificationsUpdateArray(object):
    @classmethod
    def to_proto(self, resources):
        if not resources:
            return resources
        return [
            InstanceUnfreezeRecipeStepsKeyNotificationsUpdate.to_proto(i)
            for i in resources
        ]

    @classmethod
    def from_proto(self, resources):
        return [
            InstanceUnfreezeRecipeStepsKeyNotificationsUpdate.from_proto(i)
            for i in resources
        ]


class InstanceUnfreezeRecipeStepsKeyNotificationsUpdateKeyNotificationsInfo(object):
    def __init__(
        self,
        data_version: int = None,
        delegate: str = None,
        key_notification_configs: list = None,
    ):
        self.data_version = data_version
        self.delegate = delegate
        self.key_notification_configs = key_notification_configs

    @classmethod
    def to_proto(self, resource):
        if not resource:
            return None

        res = (
            instance_pb2.Tier2AlphaInstanceUnfreezeRecipeStepsKeyNotificationsUpdateKeyNotificationsInfo()
        )
        if Primitive.to_proto(resource.data_version):
            res.data_version = Primitive.to_proto(resource.data_version)
        if Primitive.to_proto(resource.delegate):
            res.delegate = Primitive.to_proto(resource.delegate)
        if InstanceUnfreezeRecipeStepsKeyNotificationsUpdateKeyNotificationsInfoKeyNotificationConfigsArray.to_proto(
            resource.key_notification_configs
        ):
            res.key_notification_configs.extend(
                InstanceUnfreezeRecipeStepsKeyNotificationsUpdateKeyNotificationsInfoKeyNotificationConfigsArray.to_proto(
                    resource.key_notification_configs
                )
            )
        return res

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return None

        return InstanceUnfreezeRecipeStepsKeyNotificationsUpdateKeyNotificationsInfo(
            data_version=Primitive.from_proto(resource.data_version),
            delegate=Primitive.from_proto(resource.delegate),
            key_notification_configs=InstanceUnfreezeRecipeStepsKeyNotificationsUpdateKeyNotificationsInfoKeyNotificationConfigsArray.from_proto(
                resource.key_notification_configs
            ),
        )


class InstanceUnfreezeRecipeStepsKeyNotificationsUpdateKeyNotificationsInfoArray(
    object
):
    @classmethod
    def to_proto(self, resources):
        if not resources:
            return resources
        return [
            InstanceUnfreezeRecipeStepsKeyNotificationsUpdateKeyNotificationsInfo.to_proto(
                i
            )
            for i in resources
        ]

    @classmethod
    def from_proto(self, resources):
        return [
            InstanceUnfreezeRecipeStepsKeyNotificationsUpdateKeyNotificationsInfo.from_proto(
                i
            )
            for i in resources
        ]


class InstanceUnfreezeRecipeStepsKeyNotificationsUpdateKeyNotificationsInfoKeyNotificationConfigs(
    object
):
    def __init__(
        self,
        key_or_version_name: str = None,
        grant: str = None,
        delegator_gaia_id: int = None,
    ):
        self.key_or_version_name = key_or_version_name
        self.grant = grant
        self.delegator_gaia_id = delegator_gaia_id

    @classmethod
    def to_proto(self, resource):
        if not resource:
            return None

        res = (
            instance_pb2.Tier2AlphaInstanceUnfreezeRecipeStepsKeyNotificationsUpdateKeyNotificationsInfoKeyNotificationConfigs()
        )
        if Primitive.to_proto(resource.key_or_version_name):
            res.key_or_version_name = Primitive.to_proto(resource.key_or_version_name)
        if Primitive.to_proto(resource.grant):
            res.grant = Primitive.to_proto(resource.grant)
        if Primitive.to_proto(resource.delegator_gaia_id):
            res.delegator_gaia_id = Primitive.to_proto(resource.delegator_gaia_id)
        return res

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return None

        return InstanceUnfreezeRecipeStepsKeyNotificationsUpdateKeyNotificationsInfoKeyNotificationConfigs(
            key_or_version_name=Primitive.from_proto(resource.key_or_version_name),
            grant=Primitive.from_proto(resource.grant),
            delegator_gaia_id=Primitive.from_proto(resource.delegator_gaia_id),
        )


class InstanceUnfreezeRecipeStepsKeyNotificationsUpdateKeyNotificationsInfoKeyNotificationConfigsArray(
    object
):
    @classmethod
    def to_proto(self, resources):
        if not resources:
            return resources
        return [
            InstanceUnfreezeRecipeStepsKeyNotificationsUpdateKeyNotificationsInfoKeyNotificationConfigs.to_proto(
                i
            )
            for i in resources
        ]

    @classmethod
    def from_proto(self, resources):
        return [
            InstanceUnfreezeRecipeStepsKeyNotificationsUpdateKeyNotificationsInfoKeyNotificationConfigs.from_proto(
                i
            )
            for i in resources
        ]


class InstancePreprocessReportInstanceHealthRecipe(object):
    def __init__(
        self,
        steps: list = None,
        honor_cancel_request: bool = None,
        ignore_recipe_after: int = None,
        verify_deadline_seconds_below: float = None,
        populate_operation_result: bool = None,
        readonly_recipe_start_time: str = None,
        resource_names_stored_in_clh_with_delay: list = None,
        delay_to_store_resources_in_clh_db_nanos: int = None,
    ):
        self.steps = steps
        self.honor_cancel_request = honor_cancel_request
        self.ignore_recipe_after = ignore_recipe_after
        self.verify_deadline_seconds_below = verify_deadline_seconds_below
        self.populate_operation_result = populate_operation_result
        self.readonly_recipe_start_time = readonly_recipe_start_time
        self.resource_names_stored_in_clh_with_delay = (
            resource_names_stored_in_clh_with_delay
        )
        self.delay_to_store_resources_in_clh_db_nanos = (
            delay_to_store_resources_in_clh_db_nanos
        )

    @classmethod
    def to_proto(self, resource):
        if not resource:
            return None

        res = instance_pb2.Tier2AlphaInstancePreprocessReportInstanceHealthRecipe()
        if InstancePreprocessReportInstanceHealthRecipeStepsArray.to_proto(
            resource.steps
        ):
            res.steps.extend(
                InstancePreprocessReportInstanceHealthRecipeStepsArray.to_proto(
                    resource.steps
                )
            )
        if Primitive.to_proto(resource.honor_cancel_request):
            res.honor_cancel_request = Primitive.to_proto(resource.honor_cancel_request)
        if Primitive.to_proto(resource.ignore_recipe_after):
            res.ignore_recipe_after = Primitive.to_proto(resource.ignore_recipe_after)
        if Primitive.to_proto(resource.verify_deadline_seconds_below):
            res.verify_deadline_seconds_below = Primitive.to_proto(
                resource.verify_deadline_seconds_below
            )
        if Primitive.to_proto(resource.populate_operation_result):
            res.populate_operation_result = Primitive.to_proto(
                resource.populate_operation_result
            )
        if Primitive.to_proto(resource.readonly_recipe_start_time):
            res.readonly_recipe_start_time = Primitive.to_proto(
                resource.readonly_recipe_start_time
            )
        if Primitive.to_proto(resource.resource_names_stored_in_clh_with_delay):
            res.resource_names_stored_in_clh_with_delay.extend(
                Primitive.to_proto(resource.resource_names_stored_in_clh_with_delay)
            )
        if Primitive.to_proto(resource.delay_to_store_resources_in_clh_db_nanos):
            res.delay_to_store_resources_in_clh_db_nanos = Primitive.to_proto(
                resource.delay_to_store_resources_in_clh_db_nanos
            )
        return res

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return None

        return InstancePreprocessReportInstanceHealthRecipe(
            steps=InstancePreprocessReportInstanceHealthRecipeStepsArray.from_proto(
                resource.steps
            ),
            honor_cancel_request=Primitive.from_proto(resource.honor_cancel_request),
            ignore_recipe_after=Primitive.from_proto(resource.ignore_recipe_after),
            verify_deadline_seconds_below=Primitive.from_proto(
                resource.verify_deadline_seconds_below
            ),
            populate_operation_result=Primitive.from_proto(
                resource.populate_operation_result
            ),
            readonly_recipe_start_time=Primitive.from_proto(
                resource.readonly_recipe_start_time
            ),
            resource_names_stored_in_clh_with_delay=Primitive.from_proto(
                resource.resource_names_stored_in_clh_with_delay
            ),
            delay_to_store_resources_in_clh_db_nanos=Primitive.from_proto(
                resource.delay_to_store_resources_in_clh_db_nanos
            ),
        )


class InstancePreprocessReportInstanceHealthRecipeArray(object):
    @classmethod
    def to_proto(self, resources):
        if not resources:
            return resources
        return [
            InstancePreprocessReportInstanceHealthRecipe.to_proto(i) for i in resources
        ]

    @classmethod
    def from_proto(self, resources):
        return [
            InstancePreprocessReportInstanceHealthRecipe.from_proto(i)
            for i in resources
        ]


class InstancePreprocessReportInstanceHealthRecipeSteps(object):
    def __init__(
        self,
        relative_time: int = None,
        sleep_duration: int = None,
        action: str = None,
        status: dict = None,
        error_space: str = None,
        p4_service_account: str = None,
        resource_metadata_size: int = None,
        description: str = None,
        updated_repeat_operation_delay_sec: float = None,
        quota_request_deltas: list = None,
        preprocess_update: dict = None,
        public_operation_metadata: str = None,
        requested_tenant_project: dict = None,
        permissions_info: list = None,
        key_notifications_update: dict = None,
        clh_data_update_time: str = None,
        public_error_message: str = None,
    ):
        self.relative_time = relative_time
        self.sleep_duration = sleep_duration
        self.action = action
        self.status = status
        self.error_space = error_space
        self.p4_service_account = p4_service_account
        self.resource_metadata_size = resource_metadata_size
        self.description = description
        self.updated_repeat_operation_delay_sec = updated_repeat_operation_delay_sec
        self.quota_request_deltas = quota_request_deltas
        self.preprocess_update = preprocess_update
        self.public_operation_metadata = public_operation_metadata
        self.requested_tenant_project = requested_tenant_project
        self.permissions_info = permissions_info
        self.key_notifications_update = key_notifications_update
        self.clh_data_update_time = clh_data_update_time
        self.public_error_message = public_error_message

    @classmethod
    def to_proto(self, resource):
        if not resource:
            return None

        res = instance_pb2.Tier2AlphaInstancePreprocessReportInstanceHealthRecipeSteps()
        if Primitive.to_proto(resource.relative_time):
            res.relative_time = Primitive.to_proto(resource.relative_time)
        if Primitive.to_proto(resource.sleep_duration):
            res.sleep_duration = Primitive.to_proto(resource.sleep_duration)
        if InstancePreprocessReportInstanceHealthRecipeStepsActionEnum.to_proto(
            resource.action
        ):
            res.action = InstancePreprocessReportInstanceHealthRecipeStepsActionEnum.to_proto(
                resource.action
            )
        if InstancePreprocessReportInstanceHealthRecipeStepsStatus.to_proto(
            resource.status
        ):
            res.status.CopyFrom(
                InstancePreprocessReportInstanceHealthRecipeStepsStatus.to_proto(
                    resource.status
                )
            )
        else:
            res.ClearField("status")
        if Primitive.to_proto(resource.error_space):
            res.error_space = Primitive.to_proto(resource.error_space)
        if Primitive.to_proto(resource.p4_service_account):
            res.p4_service_account = Primitive.to_proto(resource.p4_service_account)
        if Primitive.to_proto(resource.resource_metadata_size):
            res.resource_metadata_size = Primitive.to_proto(
                resource.resource_metadata_size
            )
        if Primitive.to_proto(resource.description):
            res.description = Primitive.to_proto(resource.description)
        if Primitive.to_proto(resource.updated_repeat_operation_delay_sec):
            res.updated_repeat_operation_delay_sec = Primitive.to_proto(
                resource.updated_repeat_operation_delay_sec
            )
        if InstancePreprocessReportInstanceHealthRecipeStepsQuotaRequestDeltasArray.to_proto(
            resource.quota_request_deltas
        ):
            res.quota_request_deltas.extend(
                InstancePreprocessReportInstanceHealthRecipeStepsQuotaRequestDeltasArray.to_proto(
                    resource.quota_request_deltas
                )
            )
        if InstancePreprocessReportInstanceHealthRecipeStepsPreprocessUpdate.to_proto(
            resource.preprocess_update
        ):
            res.preprocess_update.CopyFrom(
                InstancePreprocessReportInstanceHealthRecipeStepsPreprocessUpdate.to_proto(
                    resource.preprocess_update
                )
            )
        else:
            res.ClearField("preprocess_update")
        if Primitive.to_proto(resource.public_operation_metadata):
            res.public_operation_metadata = Primitive.to_proto(
                resource.public_operation_metadata
            )
        if InstancePreprocessReportInstanceHealthRecipeStepsRequestedTenantProject.to_proto(
            resource.requested_tenant_project
        ):
            res.requested_tenant_project.CopyFrom(
                InstancePreprocessReportInstanceHealthRecipeStepsRequestedTenantProject.to_proto(
                    resource.requested_tenant_project
                )
            )
        else:
            res.ClearField("requested_tenant_project")
        if InstancePreprocessReportInstanceHealthRecipeStepsPermissionsInfoArray.to_proto(
            resource.permissions_info
        ):
            res.permissions_info.extend(
                InstancePreprocessReportInstanceHealthRecipeStepsPermissionsInfoArray.to_proto(
                    resource.permissions_info
                )
            )
        if InstancePreprocessReportInstanceHealthRecipeStepsKeyNotificationsUpdate.to_proto(
            resource.key_notifications_update
        ):
            res.key_notifications_update.CopyFrom(
                InstancePreprocessReportInstanceHealthRecipeStepsKeyNotificationsUpdate.to_proto(
                    resource.key_notifications_update
                )
            )
        else:
            res.ClearField("key_notifications_update")
        if Primitive.to_proto(resource.clh_data_update_time):
            res.clh_data_update_time = Primitive.to_proto(resource.clh_data_update_time)
        if Primitive.to_proto(resource.public_error_message):
            res.public_error_message = Primitive.to_proto(resource.public_error_message)
        return res

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return None

        return InstancePreprocessReportInstanceHealthRecipeSteps(
            relative_time=Primitive.from_proto(resource.relative_time),
            sleep_duration=Primitive.from_proto(resource.sleep_duration),
            action=InstancePreprocessReportInstanceHealthRecipeStepsActionEnum.from_proto(
                resource.action
            ),
            status=InstancePreprocessReportInstanceHealthRecipeStepsStatus.from_proto(
                resource.status
            ),
            error_space=Primitive.from_proto(resource.error_space),
            p4_service_account=Primitive.from_proto(resource.p4_service_account),
            resource_metadata_size=Primitive.from_proto(
                resource.resource_metadata_size
            ),
            description=Primitive.from_proto(resource.description),
            updated_repeat_operation_delay_sec=Primitive.from_proto(
                resource.updated_repeat_operation_delay_sec
            ),
            quota_request_deltas=InstancePreprocessReportInstanceHealthRecipeStepsQuotaRequestDeltasArray.from_proto(
                resource.quota_request_deltas
            ),
            preprocess_update=InstancePreprocessReportInstanceHealthRecipeStepsPreprocessUpdate.from_proto(
                resource.preprocess_update
            ),
            public_operation_metadata=Primitive.from_proto(
                resource.public_operation_metadata
            ),
            requested_tenant_project=InstancePreprocessReportInstanceHealthRecipeStepsRequestedTenantProject.from_proto(
                resource.requested_tenant_project
            ),
            permissions_info=InstancePreprocessReportInstanceHealthRecipeStepsPermissionsInfoArray.from_proto(
                resource.permissions_info
            ),
            key_notifications_update=InstancePreprocessReportInstanceHealthRecipeStepsKeyNotificationsUpdate.from_proto(
                resource.key_notifications_update
            ),
            clh_data_update_time=Primitive.from_proto(resource.clh_data_update_time),
            public_error_message=Primitive.from_proto(resource.public_error_message),
        )


class InstancePreprocessReportInstanceHealthRecipeStepsArray(object):
    @classmethod
    def to_proto(self, resources):
        if not resources:
            return resources
        return [
            InstancePreprocessReportInstanceHealthRecipeSteps.to_proto(i)
            for i in resources
        ]

    @classmethod
    def from_proto(self, resources):
        return [
            InstancePreprocessReportInstanceHealthRecipeSteps.from_proto(i)
            for i in resources
        ]


class InstancePreprocessReportInstanceHealthRecipeStepsStatus(object):
    def __init__(self, code: int = None, message: str = None, details: list = None):
        self.code = code
        self.message = message
        self.details = details

    @classmethod
    def to_proto(self, resource):
        if not resource:
            return None

        res = (
            instance_pb2.Tier2AlphaInstancePreprocessReportInstanceHealthRecipeStepsStatus()
        )
        if Primitive.to_proto(resource.code):
            res.code = Primitive.to_proto(resource.code)
        if Primitive.to_proto(resource.message):
            res.message = Primitive.to_proto(resource.message)
        if InstancePreprocessReportInstanceHealthRecipeStepsStatusDetailsArray.to_proto(
            resource.details
        ):
            res.details.extend(
                InstancePreprocessReportInstanceHealthRecipeStepsStatusDetailsArray.to_proto(
                    resource.details
                )
            )
        return res

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return None

        return InstancePreprocessReportInstanceHealthRecipeStepsStatus(
            code=Primitive.from_proto(resource.code),
            message=Primitive.from_proto(resource.message),
            details=InstancePreprocessReportInstanceHealthRecipeStepsStatusDetailsArray.from_proto(
                resource.details
            ),
        )


class InstancePreprocessReportInstanceHealthRecipeStepsStatusArray(object):
    @classmethod
    def to_proto(self, resources):
        if not resources:
            return resources
        return [
            InstancePreprocessReportInstanceHealthRecipeStepsStatus.to_proto(i)
            for i in resources
        ]

    @classmethod
    def from_proto(self, resources):
        return [
            InstancePreprocessReportInstanceHealthRecipeStepsStatus.from_proto(i)
            for i in resources
        ]


class InstancePreprocessReportInstanceHealthRecipeStepsStatusDetails(object):
    def __init__(self, type_url: str = None, value: str = None):
        self.type_url = type_url
        self.value = value

    @classmethod
    def to_proto(self, resource):
        if not resource:
            return None

        res = (
            instance_pb2.Tier2AlphaInstancePreprocessReportInstanceHealthRecipeStepsStatusDetails()
        )
        if Primitive.to_proto(resource.type_url):
            res.type_url = Primitive.to_proto(resource.type_url)
        if Primitive.to_proto(resource.value):
            res.value = Primitive.to_proto(resource.value)
        return res

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return None

        return InstancePreprocessReportInstanceHealthRecipeStepsStatusDetails(
            type_url=Primitive.from_proto(resource.type_url),
            value=Primitive.from_proto(resource.value),
        )


class InstancePreprocessReportInstanceHealthRecipeStepsStatusDetailsArray(object):
    @classmethod
    def to_proto(self, resources):
        if not resources:
            return resources
        return [
            InstancePreprocessReportInstanceHealthRecipeStepsStatusDetails.to_proto(i)
            for i in resources
        ]

    @classmethod
    def from_proto(self, resources):
        return [
            InstancePreprocessReportInstanceHealthRecipeStepsStatusDetails.from_proto(i)
            for i in resources
        ]


class InstancePreprocessReportInstanceHealthRecipeStepsQuotaRequestDeltas(object):
    def __init__(
        self,
        metric_name: str = None,
        amount: int = None,
        quota_location_name: str = None,
    ):
        self.metric_name = metric_name
        self.amount = amount
        self.quota_location_name = quota_location_name

    @classmethod
    def to_proto(self, resource):
        if not resource:
            return None

        res = (
            instance_pb2.Tier2AlphaInstancePreprocessReportInstanceHealthRecipeStepsQuotaRequestDeltas()
        )
        if Primitive.to_proto(resource.metric_name):
            res.metric_name = Primitive.to_proto(resource.metric_name)
        if Primitive.to_proto(resource.amount):
            res.amount = Primitive.to_proto(resource.amount)
        if Primitive.to_proto(resource.quota_location_name):
            res.quota_location_name = Primitive.to_proto(resource.quota_location_name)
        return res

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return None

        return InstancePreprocessReportInstanceHealthRecipeStepsQuotaRequestDeltas(
            metric_name=Primitive.from_proto(resource.metric_name),
            amount=Primitive.from_proto(resource.amount),
            quota_location_name=Primitive.from_proto(resource.quota_location_name),
        )


class InstancePreprocessReportInstanceHealthRecipeStepsQuotaRequestDeltasArray(object):
    @classmethod
    def to_proto(self, resources):
        if not resources:
            return resources
        return [
            InstancePreprocessReportInstanceHealthRecipeStepsQuotaRequestDeltas.to_proto(
                i
            )
            for i in resources
        ]

    @classmethod
    def from_proto(self, resources):
        return [
            InstancePreprocessReportInstanceHealthRecipeStepsQuotaRequestDeltas.from_proto(
                i
            )
            for i in resources
        ]


class InstancePreprocessReportInstanceHealthRecipeStepsPreprocessUpdate(object):
    def __init__(
        self, latency_slo_bucket_name: str = None, public_operation_metadata: str = None
    ):
        self.latency_slo_bucket_name = latency_slo_bucket_name
        self.public_operation_metadata = public_operation_metadata

    @classmethod
    def to_proto(self, resource):
        if not resource:
            return None

        res = (
            instance_pb2.Tier2AlphaInstancePreprocessReportInstanceHealthRecipeStepsPreprocessUpdate()
        )
        if Primitive.to_proto(resource.latency_slo_bucket_name):
            res.latency_slo_bucket_name = Primitive.to_proto(
                resource.latency_slo_bucket_name
            )
        if Primitive.to_proto(resource.public_operation_metadata):
            res.public_operation_metadata = Primitive.to_proto(
                resource.public_operation_metadata
            )
        return res

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return None

        return InstancePreprocessReportInstanceHealthRecipeStepsPreprocessUpdate(
            latency_slo_bucket_name=Primitive.from_proto(
                resource.latency_slo_bucket_name
            ),
            public_operation_metadata=Primitive.from_proto(
                resource.public_operation_metadata
            ),
        )


class InstancePreprocessReportInstanceHealthRecipeStepsPreprocessUpdateArray(object):
    @classmethod
    def to_proto(self, resources):
        if not resources:
            return resources
        return [
            InstancePreprocessReportInstanceHealthRecipeStepsPreprocessUpdate.to_proto(
                i
            )
            for i in resources
        ]

    @classmethod
    def from_proto(self, resources):
        return [
            InstancePreprocessReportInstanceHealthRecipeStepsPreprocessUpdate.from_proto(
                i
            )
            for i in resources
        ]


class InstancePreprocessReportInstanceHealthRecipeStepsRequestedTenantProject(object):
    def __init__(self, tag: str = None, folder: str = None, scope: str = None):
        self.tag = tag
        self.folder = folder
        self.scope = scope

    @classmethod
    def to_proto(self, resource):
        if not resource:
            return None

        res = (
            instance_pb2.Tier2AlphaInstancePreprocessReportInstanceHealthRecipeStepsRequestedTenantProject()
        )
        if Primitive.to_proto(resource.tag):
            res.tag = Primitive.to_proto(resource.tag)
        if Primitive.to_proto(resource.folder):
            res.folder = Primitive.to_proto(resource.folder)
        if InstancePreprocessReportInstanceHealthRecipeStepsRequestedTenantProjectScopeEnum.to_proto(
            resource.scope
        ):
            res.scope = InstancePreprocessReportInstanceHealthRecipeStepsRequestedTenantProjectScopeEnum.to_proto(
                resource.scope
            )
        return res

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return None

        return InstancePreprocessReportInstanceHealthRecipeStepsRequestedTenantProject(
            tag=Primitive.from_proto(resource.tag),
            folder=Primitive.from_proto(resource.folder),
            scope=InstancePreprocessReportInstanceHealthRecipeStepsRequestedTenantProjectScopeEnum.from_proto(
                resource.scope
            ),
        )


class InstancePreprocessReportInstanceHealthRecipeStepsRequestedTenantProjectArray(
    object
):
    @classmethod
    def to_proto(self, resources):
        if not resources:
            return resources
        return [
            InstancePreprocessReportInstanceHealthRecipeStepsRequestedTenantProject.to_proto(
                i
            )
            for i in resources
        ]

    @classmethod
    def from_proto(self, resources):
        return [
            InstancePreprocessReportInstanceHealthRecipeStepsRequestedTenantProject.from_proto(
                i
            )
            for i in resources
        ]


class InstancePreprocessReportInstanceHealthRecipeStepsPermissionsInfo(object):
    def __init__(
        self,
        policy_name: dict = None,
        iam_permissions: list = None,
        resource_path: str = None,
        api_attrs: dict = None,
        policy_name_mode: str = None,
        resource: dict = None,
    ):
        self.policy_name = policy_name
        self.iam_permissions = iam_permissions
        self.resource_path = resource_path
        self.api_attrs = api_attrs
        self.policy_name_mode = policy_name_mode
        self.resource = resource

    @classmethod
    def to_proto(self, resource):
        if not resource:
            return None

        res = (
            instance_pb2.Tier2AlphaInstancePreprocessReportInstanceHealthRecipeStepsPermissionsInfo()
        )
        if InstancePreprocessReportInstanceHealthRecipeStepsPermissionsInfoPolicyName.to_proto(
            resource.policy_name
        ):
            res.policy_name.CopyFrom(
                InstancePreprocessReportInstanceHealthRecipeStepsPermissionsInfoPolicyName.to_proto(
                    resource.policy_name
                )
            )
        else:
            res.ClearField("policy_name")
        if InstancePreprocessReportInstanceHealthRecipeStepsPermissionsInfoIamPermissionsArray.to_proto(
            resource.iam_permissions
        ):
            res.iam_permissions.extend(
                InstancePreprocessReportInstanceHealthRecipeStepsPermissionsInfoIamPermissionsArray.to_proto(
                    resource.iam_permissions
                )
            )
        if Primitive.to_proto(resource.resource_path):
            res.resource_path = Primitive.to_proto(resource.resource_path)
        if InstanceGoogleprotobufstruct.to_proto(resource.api_attrs):
            res.api_attrs.CopyFrom(
                InstanceGoogleprotobufstruct.to_proto(resource.api_attrs)
            )
        else:
            res.ClearField("api_attrs")
        if InstancePreprocessReportInstanceHealthRecipeStepsPermissionsInfoPolicyNameModeEnum.to_proto(
            resource.policy_name_mode
        ):
            res.policy_name_mode = InstancePreprocessReportInstanceHealthRecipeStepsPermissionsInfoPolicyNameModeEnum.to_proto(
                resource.policy_name_mode
            )
        if InstancePreprocessReportInstanceHealthRecipeStepsPermissionsInfoResource.to_proto(
            resource.resource
        ):
            res.resource.CopyFrom(
                InstancePreprocessReportInstanceHealthRecipeStepsPermissionsInfoResource.to_proto(
                    resource.resource
                )
            )
        else:
            res.ClearField("resource")
        return res

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return None

        return InstancePreprocessReportInstanceHealthRecipeStepsPermissionsInfo(
            policy_name=InstancePreprocessReportInstanceHealthRecipeStepsPermissionsInfoPolicyName.from_proto(
                resource.policy_name
            ),
            iam_permissions=InstancePreprocessReportInstanceHealthRecipeStepsPermissionsInfoIamPermissionsArray.from_proto(
                resource.iam_permissions
            ),
            resource_path=Primitive.from_proto(resource.resource_path),
            api_attrs=InstanceGoogleprotobufstruct.from_proto(resource.api_attrs),
            policy_name_mode=InstancePreprocessReportInstanceHealthRecipeStepsPermissionsInfoPolicyNameModeEnum.from_proto(
                resource.policy_name_mode
            ),
            resource=InstancePreprocessReportInstanceHealthRecipeStepsPermissionsInfoResource.from_proto(
                resource.resource
            ),
        )


class InstancePreprocessReportInstanceHealthRecipeStepsPermissionsInfoArray(object):
    @classmethod
    def to_proto(self, resources):
        if not resources:
            return resources
        return [
            InstancePreprocessReportInstanceHealthRecipeStepsPermissionsInfo.to_proto(i)
            for i in resources
        ]

    @classmethod
    def from_proto(self, resources):
        return [
            InstancePreprocessReportInstanceHealthRecipeStepsPermissionsInfo.from_proto(
                i
            )
            for i in resources
        ]


class InstancePreprocessReportInstanceHealthRecipeStepsPermissionsInfoPolicyName(
    object
):
    def __init__(self, type: str = None, id: str = None, region: str = None):
        self.type = type
        self.id = id
        self.region = region

    @classmethod
    def to_proto(self, resource):
        if not resource:
            return None

        res = (
            instance_pb2.Tier2AlphaInstancePreprocessReportInstanceHealthRecipeStepsPermissionsInfoPolicyName()
        )
        if Primitive.to_proto(resource.type):
            res.type = Primitive.to_proto(resource.type)
        if Primitive.to_proto(resource.id):
            res.id = Primitive.to_proto(resource.id)
        if Primitive.to_proto(resource.region):
            res.region = Primitive.to_proto(resource.region)
        return res

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return None

        return InstancePreprocessReportInstanceHealthRecipeStepsPermissionsInfoPolicyName(
            type=Primitive.from_proto(resource.type),
            id=Primitive.from_proto(resource.id),
            region=Primitive.from_proto(resource.region),
        )


class InstancePreprocessReportInstanceHealthRecipeStepsPermissionsInfoPolicyNameArray(
    object
):
    @classmethod
    def to_proto(self, resources):
        if not resources:
            return resources
        return [
            InstancePreprocessReportInstanceHealthRecipeStepsPermissionsInfoPolicyName.to_proto(
                i
            )
            for i in resources
        ]

    @classmethod
    def from_proto(self, resources):
        return [
            InstancePreprocessReportInstanceHealthRecipeStepsPermissionsInfoPolicyName.from_proto(
                i
            )
            for i in resources
        ]


class InstancePreprocessReportInstanceHealthRecipeStepsPermissionsInfoIamPermissions(
    object
):
    def __init__(self, permission: str = None):
        self.permission = permission

    @classmethod
    def to_proto(self, resource):
        if not resource:
            return None

        res = (
            instance_pb2.Tier2AlphaInstancePreprocessReportInstanceHealthRecipeStepsPermissionsInfoIamPermissions()
        )
        if Primitive.to_proto(resource.permission):
            res.permission = Primitive.to_proto(resource.permission)
        return res

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return None

        return InstancePreprocessReportInstanceHealthRecipeStepsPermissionsInfoIamPermissions(
            permission=Primitive.from_proto(resource.permission),
        )


class InstancePreprocessReportInstanceHealthRecipeStepsPermissionsInfoIamPermissionsArray(
    object
):
    @classmethod
    def to_proto(self, resources):
        if not resources:
            return resources
        return [
            InstancePreprocessReportInstanceHealthRecipeStepsPermissionsInfoIamPermissions.to_proto(
                i
            )
            for i in resources
        ]

    @classmethod
    def from_proto(self, resources):
        return [
            InstancePreprocessReportInstanceHealthRecipeStepsPermissionsInfoIamPermissions.from_proto(
                i
            )
            for i in resources
        ]


class InstancePreprocessReportInstanceHealthRecipeStepsPermissionsInfoResource(object):
    def __init__(
        self,
        name: str = None,
        type: str = None,
        service: str = None,
        labels: dict = None,
    ):
        self.name = name
        self.type = type
        self.service = service
        self.labels = labels

    @classmethod
    def to_proto(self, resource):
        if not resource:
            return None

        res = (
            instance_pb2.Tier2AlphaInstancePreprocessReportInstanceHealthRecipeStepsPermissionsInfoResource()
        )
        if Primitive.to_proto(resource.name):
            res.name = Primitive.to_proto(resource.name)
        if Primitive.to_proto(resource.type):
            res.type = Primitive.to_proto(resource.type)
        if Primitive.to_proto(resource.service):
            res.service = Primitive.to_proto(resource.service)
        if Primitive.to_proto(resource.labels):
            res.labels = Primitive.to_proto(resource.labels)
        return res

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return None

        return InstancePreprocessReportInstanceHealthRecipeStepsPermissionsInfoResource(
            name=Primitive.from_proto(resource.name),
            type=Primitive.from_proto(resource.type),
            service=Primitive.from_proto(resource.service),
            labels=Primitive.from_proto(resource.labels),
        )


class InstancePreprocessReportInstanceHealthRecipeStepsPermissionsInfoResourceArray(
    object
):
    @classmethod
    def to_proto(self, resources):
        if not resources:
            return resources
        return [
            InstancePreprocessReportInstanceHealthRecipeStepsPermissionsInfoResource.to_proto(
                i
            )
            for i in resources
        ]

    @classmethod
    def from_proto(self, resources):
        return [
            InstancePreprocessReportInstanceHealthRecipeStepsPermissionsInfoResource.from_proto(
                i
            )
            for i in resources
        ]


class InstancePreprocessReportInstanceHealthRecipeStepsKeyNotificationsUpdate(object):
    def __init__(self, key_notifications_info: dict = None):
        self.key_notifications_info = key_notifications_info

    @classmethod
    def to_proto(self, resource):
        if not resource:
            return None

        res = (
            instance_pb2.Tier2AlphaInstancePreprocessReportInstanceHealthRecipeStepsKeyNotificationsUpdate()
        )
        if InstancePreprocessReportInstanceHealthRecipeStepsKeyNotificationsUpdateKeyNotificationsInfo.to_proto(
            resource.key_notifications_info
        ):
            res.key_notifications_info.CopyFrom(
                InstancePreprocessReportInstanceHealthRecipeStepsKeyNotificationsUpdateKeyNotificationsInfo.to_proto(
                    resource.key_notifications_info
                )
            )
        else:
            res.ClearField("key_notifications_info")
        return res

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return None

        return InstancePreprocessReportInstanceHealthRecipeStepsKeyNotificationsUpdate(
            key_notifications_info=InstancePreprocessReportInstanceHealthRecipeStepsKeyNotificationsUpdateKeyNotificationsInfo.from_proto(
                resource.key_notifications_info
            ),
        )


class InstancePreprocessReportInstanceHealthRecipeStepsKeyNotificationsUpdateArray(
    object
):
    @classmethod
    def to_proto(self, resources):
        if not resources:
            return resources
        return [
            InstancePreprocessReportInstanceHealthRecipeStepsKeyNotificationsUpdate.to_proto(
                i
            )
            for i in resources
        ]

    @classmethod
    def from_proto(self, resources):
        return [
            InstancePreprocessReportInstanceHealthRecipeStepsKeyNotificationsUpdate.from_proto(
                i
            )
            for i in resources
        ]


class InstancePreprocessReportInstanceHealthRecipeStepsKeyNotificationsUpdateKeyNotificationsInfo(
    object
):
    def __init__(
        self,
        data_version: int = None,
        delegate: str = None,
        key_notification_configs: list = None,
    ):
        self.data_version = data_version
        self.delegate = delegate
        self.key_notification_configs = key_notification_configs

    @classmethod
    def to_proto(self, resource):
        if not resource:
            return None

        res = (
            instance_pb2.Tier2AlphaInstancePreprocessReportInstanceHealthRecipeStepsKeyNotificationsUpdateKeyNotificationsInfo()
        )
        if Primitive.to_proto(resource.data_version):
            res.data_version = Primitive.to_proto(resource.data_version)
        if Primitive.to_proto(resource.delegate):
            res.delegate = Primitive.to_proto(resource.delegate)
        if InstancePreprocessReportInstanceHealthRecipeStepsKeyNotificationsUpdateKeyNotificationsInfoKeyNotificationConfigsArray.to_proto(
            resource.key_notification_configs
        ):
            res.key_notification_configs.extend(
                InstancePreprocessReportInstanceHealthRecipeStepsKeyNotificationsUpdateKeyNotificationsInfoKeyNotificationConfigsArray.to_proto(
                    resource.key_notification_configs
                )
            )
        return res

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return None

        return InstancePreprocessReportInstanceHealthRecipeStepsKeyNotificationsUpdateKeyNotificationsInfo(
            data_version=Primitive.from_proto(resource.data_version),
            delegate=Primitive.from_proto(resource.delegate),
            key_notification_configs=InstancePreprocessReportInstanceHealthRecipeStepsKeyNotificationsUpdateKeyNotificationsInfoKeyNotificationConfigsArray.from_proto(
                resource.key_notification_configs
            ),
        )


class InstancePreprocessReportInstanceHealthRecipeStepsKeyNotificationsUpdateKeyNotificationsInfoArray(
    object
):
    @classmethod
    def to_proto(self, resources):
        if not resources:
            return resources
        return [
            InstancePreprocessReportInstanceHealthRecipeStepsKeyNotificationsUpdateKeyNotificationsInfo.to_proto(
                i
            )
            for i in resources
        ]

    @classmethod
    def from_proto(self, resources):
        return [
            InstancePreprocessReportInstanceHealthRecipeStepsKeyNotificationsUpdateKeyNotificationsInfo.from_proto(
                i
            )
            for i in resources
        ]


class InstancePreprocessReportInstanceHealthRecipeStepsKeyNotificationsUpdateKeyNotificationsInfoKeyNotificationConfigs(
    object
):
    def __init__(
        self,
        key_or_version_name: str = None,
        grant: str = None,
        delegator_gaia_id: int = None,
    ):
        self.key_or_version_name = key_or_version_name
        self.grant = grant
        self.delegator_gaia_id = delegator_gaia_id

    @classmethod
    def to_proto(self, resource):
        if not resource:
            return None

        res = (
            instance_pb2.Tier2AlphaInstancePreprocessReportInstanceHealthRecipeStepsKeyNotificationsUpdateKeyNotificationsInfoKeyNotificationConfigs()
        )
        if Primitive.to_proto(resource.key_or_version_name):
            res.key_or_version_name = Primitive.to_proto(resource.key_or_version_name)
        if Primitive.to_proto(resource.grant):
            res.grant = Primitive.to_proto(resource.grant)
        if Primitive.to_proto(resource.delegator_gaia_id):
            res.delegator_gaia_id = Primitive.to_proto(resource.delegator_gaia_id)
        return res

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return None

        return InstancePreprocessReportInstanceHealthRecipeStepsKeyNotificationsUpdateKeyNotificationsInfoKeyNotificationConfigs(
            key_or_version_name=Primitive.from_proto(resource.key_or_version_name),
            grant=Primitive.from_proto(resource.grant),
            delegator_gaia_id=Primitive.from_proto(resource.delegator_gaia_id),
        )


class InstancePreprocessReportInstanceHealthRecipeStepsKeyNotificationsUpdateKeyNotificationsInfoKeyNotificationConfigsArray(
    object
):
    @classmethod
    def to_proto(self, resources):
        if not resources:
            return resources
        return [
            InstancePreprocessReportInstanceHealthRecipeStepsKeyNotificationsUpdateKeyNotificationsInfoKeyNotificationConfigs.to_proto(
                i
            )
            for i in resources
        ]

    @classmethod
    def from_proto(self, resources):
        return [
            InstancePreprocessReportInstanceHealthRecipeStepsKeyNotificationsUpdateKeyNotificationsInfoKeyNotificationConfigs.from_proto(
                i
            )
            for i in resources
        ]


class InstanceReportInstanceHealthRecipe(object):
    def __init__(
        self,
        steps: list = None,
        honor_cancel_request: bool = None,
        ignore_recipe_after: int = None,
        verify_deadline_seconds_below: float = None,
        populate_operation_result: bool = None,
        readonly_recipe_start_time: str = None,
        resource_names_stored_in_clh_with_delay: list = None,
        delay_to_store_resources_in_clh_db_nanos: int = None,
    ):
        self.steps = steps
        self.honor_cancel_request = honor_cancel_request
        self.ignore_recipe_after = ignore_recipe_after
        self.verify_deadline_seconds_below = verify_deadline_seconds_below
        self.populate_operation_result = populate_operation_result
        self.readonly_recipe_start_time = readonly_recipe_start_time
        self.resource_names_stored_in_clh_with_delay = (
            resource_names_stored_in_clh_with_delay
        )
        self.delay_to_store_resources_in_clh_db_nanos = (
            delay_to_store_resources_in_clh_db_nanos
        )

    @classmethod
    def to_proto(self, resource):
        if not resource:
            return None

        res = instance_pb2.Tier2AlphaInstanceReportInstanceHealthRecipe()
        if InstanceReportInstanceHealthRecipeStepsArray.to_proto(resource.steps):
            res.steps.extend(
                InstanceReportInstanceHealthRecipeStepsArray.to_proto(resource.steps)
            )
        if Primitive.to_proto(resource.honor_cancel_request):
            res.honor_cancel_request = Primitive.to_proto(resource.honor_cancel_request)
        if Primitive.to_proto(resource.ignore_recipe_after):
            res.ignore_recipe_after = Primitive.to_proto(resource.ignore_recipe_after)
        if Primitive.to_proto(resource.verify_deadline_seconds_below):
            res.verify_deadline_seconds_below = Primitive.to_proto(
                resource.verify_deadline_seconds_below
            )
        if Primitive.to_proto(resource.populate_operation_result):
            res.populate_operation_result = Primitive.to_proto(
                resource.populate_operation_result
            )
        if Primitive.to_proto(resource.readonly_recipe_start_time):
            res.readonly_recipe_start_time = Primitive.to_proto(
                resource.readonly_recipe_start_time
            )
        if Primitive.to_proto(resource.resource_names_stored_in_clh_with_delay):
            res.resource_names_stored_in_clh_with_delay.extend(
                Primitive.to_proto(resource.resource_names_stored_in_clh_with_delay)
            )
        if Primitive.to_proto(resource.delay_to_store_resources_in_clh_db_nanos):
            res.delay_to_store_resources_in_clh_db_nanos = Primitive.to_proto(
                resource.delay_to_store_resources_in_clh_db_nanos
            )
        return res

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return None

        return InstanceReportInstanceHealthRecipe(
            steps=InstanceReportInstanceHealthRecipeStepsArray.from_proto(
                resource.steps
            ),
            honor_cancel_request=Primitive.from_proto(resource.honor_cancel_request),
            ignore_recipe_after=Primitive.from_proto(resource.ignore_recipe_after),
            verify_deadline_seconds_below=Primitive.from_proto(
                resource.verify_deadline_seconds_below
            ),
            populate_operation_result=Primitive.from_proto(
                resource.populate_operation_result
            ),
            readonly_recipe_start_time=Primitive.from_proto(
                resource.readonly_recipe_start_time
            ),
            resource_names_stored_in_clh_with_delay=Primitive.from_proto(
                resource.resource_names_stored_in_clh_with_delay
            ),
            delay_to_store_resources_in_clh_db_nanos=Primitive.from_proto(
                resource.delay_to_store_resources_in_clh_db_nanos
            ),
        )


class InstanceReportInstanceHealthRecipeArray(object):
    @classmethod
    def to_proto(self, resources):
        if not resources:
            return resources
        return [InstanceReportInstanceHealthRecipe.to_proto(i) for i in resources]

    @classmethod
    def from_proto(self, resources):
        return [InstanceReportInstanceHealthRecipe.from_proto(i) for i in resources]


class InstanceReportInstanceHealthRecipeSteps(object):
    def __init__(
        self,
        relative_time: int = None,
        sleep_duration: int = None,
        action: str = None,
        status: dict = None,
        error_space: str = None,
        p4_service_account: str = None,
        resource_metadata_size: int = None,
        description: str = None,
        updated_repeat_operation_delay_sec: float = None,
        quota_request_deltas: list = None,
        preprocess_update: dict = None,
        public_operation_metadata: str = None,
        requested_tenant_project: dict = None,
        permissions_info: list = None,
        key_notifications_update: dict = None,
        clh_data_update_time: str = None,
        public_error_message: str = None,
    ):
        self.relative_time = relative_time
        self.sleep_duration = sleep_duration
        self.action = action
        self.status = status
        self.error_space = error_space
        self.p4_service_account = p4_service_account
        self.resource_metadata_size = resource_metadata_size
        self.description = description
        self.updated_repeat_operation_delay_sec = updated_repeat_operation_delay_sec
        self.quota_request_deltas = quota_request_deltas
        self.preprocess_update = preprocess_update
        self.public_operation_metadata = public_operation_metadata
        self.requested_tenant_project = requested_tenant_project
        self.permissions_info = permissions_info
        self.key_notifications_update = key_notifications_update
        self.clh_data_update_time = clh_data_update_time
        self.public_error_message = public_error_message

    @classmethod
    def to_proto(self, resource):
        if not resource:
            return None

        res = instance_pb2.Tier2AlphaInstanceReportInstanceHealthRecipeSteps()
        if Primitive.to_proto(resource.relative_time):
            res.relative_time = Primitive.to_proto(resource.relative_time)
        if Primitive.to_proto(resource.sleep_duration):
            res.sleep_duration = Primitive.to_proto(resource.sleep_duration)
        if InstanceReportInstanceHealthRecipeStepsActionEnum.to_proto(resource.action):
            res.action = InstanceReportInstanceHealthRecipeStepsActionEnum.to_proto(
                resource.action
            )
        if InstanceReportInstanceHealthRecipeStepsStatus.to_proto(resource.status):
            res.status.CopyFrom(
                InstanceReportInstanceHealthRecipeStepsStatus.to_proto(resource.status)
            )
        else:
            res.ClearField("status")
        if Primitive.to_proto(resource.error_space):
            res.error_space = Primitive.to_proto(resource.error_space)
        if Primitive.to_proto(resource.p4_service_account):
            res.p4_service_account = Primitive.to_proto(resource.p4_service_account)
        if Primitive.to_proto(resource.resource_metadata_size):
            res.resource_metadata_size = Primitive.to_proto(
                resource.resource_metadata_size
            )
        if Primitive.to_proto(resource.description):
            res.description = Primitive.to_proto(resource.description)
        if Primitive.to_proto(resource.updated_repeat_operation_delay_sec):
            res.updated_repeat_operation_delay_sec = Primitive.to_proto(
                resource.updated_repeat_operation_delay_sec
            )
        if InstanceReportInstanceHealthRecipeStepsQuotaRequestDeltasArray.to_proto(
            resource.quota_request_deltas
        ):
            res.quota_request_deltas.extend(
                InstanceReportInstanceHealthRecipeStepsQuotaRequestDeltasArray.to_proto(
                    resource.quota_request_deltas
                )
            )
        if InstanceReportInstanceHealthRecipeStepsPreprocessUpdate.to_proto(
            resource.preprocess_update
        ):
            res.preprocess_update.CopyFrom(
                InstanceReportInstanceHealthRecipeStepsPreprocessUpdate.to_proto(
                    resource.preprocess_update
                )
            )
        else:
            res.ClearField("preprocess_update")
        if Primitive.to_proto(resource.public_operation_metadata):
            res.public_operation_metadata = Primitive.to_proto(
                resource.public_operation_metadata
            )
        if InstanceReportInstanceHealthRecipeStepsRequestedTenantProject.to_proto(
            resource.requested_tenant_project
        ):
            res.requested_tenant_project.CopyFrom(
                InstanceReportInstanceHealthRecipeStepsRequestedTenantProject.to_proto(
                    resource.requested_tenant_project
                )
            )
        else:
            res.ClearField("requested_tenant_project")
        if InstanceReportInstanceHealthRecipeStepsPermissionsInfoArray.to_proto(
            resource.permissions_info
        ):
            res.permissions_info.extend(
                InstanceReportInstanceHealthRecipeStepsPermissionsInfoArray.to_proto(
                    resource.permissions_info
                )
            )
        if InstanceReportInstanceHealthRecipeStepsKeyNotificationsUpdate.to_proto(
            resource.key_notifications_update
        ):
            res.key_notifications_update.CopyFrom(
                InstanceReportInstanceHealthRecipeStepsKeyNotificationsUpdate.to_proto(
                    resource.key_notifications_update
                )
            )
        else:
            res.ClearField("key_notifications_update")
        if Primitive.to_proto(resource.clh_data_update_time):
            res.clh_data_update_time = Primitive.to_proto(resource.clh_data_update_time)
        if Primitive.to_proto(resource.public_error_message):
            res.public_error_message = Primitive.to_proto(resource.public_error_message)
        return res

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return None

        return InstanceReportInstanceHealthRecipeSteps(
            relative_time=Primitive.from_proto(resource.relative_time),
            sleep_duration=Primitive.from_proto(resource.sleep_duration),
            action=InstanceReportInstanceHealthRecipeStepsActionEnum.from_proto(
                resource.action
            ),
            status=InstanceReportInstanceHealthRecipeStepsStatus.from_proto(
                resource.status
            ),
            error_space=Primitive.from_proto(resource.error_space),
            p4_service_account=Primitive.from_proto(resource.p4_service_account),
            resource_metadata_size=Primitive.from_proto(
                resource.resource_metadata_size
            ),
            description=Primitive.from_proto(resource.description),
            updated_repeat_operation_delay_sec=Primitive.from_proto(
                resource.updated_repeat_operation_delay_sec
            ),
            quota_request_deltas=InstanceReportInstanceHealthRecipeStepsQuotaRequestDeltasArray.from_proto(
                resource.quota_request_deltas
            ),
            preprocess_update=InstanceReportInstanceHealthRecipeStepsPreprocessUpdate.from_proto(
                resource.preprocess_update
            ),
            public_operation_metadata=Primitive.from_proto(
                resource.public_operation_metadata
            ),
            requested_tenant_project=InstanceReportInstanceHealthRecipeStepsRequestedTenantProject.from_proto(
                resource.requested_tenant_project
            ),
            permissions_info=InstanceReportInstanceHealthRecipeStepsPermissionsInfoArray.from_proto(
                resource.permissions_info
            ),
            key_notifications_update=InstanceReportInstanceHealthRecipeStepsKeyNotificationsUpdate.from_proto(
                resource.key_notifications_update
            ),
            clh_data_update_time=Primitive.from_proto(resource.clh_data_update_time),
            public_error_message=Primitive.from_proto(resource.public_error_message),
        )


class InstanceReportInstanceHealthRecipeStepsArray(object):
    @classmethod
    def to_proto(self, resources):
        if not resources:
            return resources
        return [InstanceReportInstanceHealthRecipeSteps.to_proto(i) for i in resources]

    @classmethod
    def from_proto(self, resources):
        return [
            InstanceReportInstanceHealthRecipeSteps.from_proto(i) for i in resources
        ]


class InstanceReportInstanceHealthRecipeStepsStatus(object):
    def __init__(self, code: int = None, message: str = None, details: list = None):
        self.code = code
        self.message = message
        self.details = details

    @classmethod
    def to_proto(self, resource):
        if not resource:
            return None

        res = instance_pb2.Tier2AlphaInstanceReportInstanceHealthRecipeStepsStatus()
        if Primitive.to_proto(resource.code):
            res.code = Primitive.to_proto(resource.code)
        if Primitive.to_proto(resource.message):
            res.message = Primitive.to_proto(resource.message)
        if InstanceReportInstanceHealthRecipeStepsStatusDetailsArray.to_proto(
            resource.details
        ):
            res.details.extend(
                InstanceReportInstanceHealthRecipeStepsStatusDetailsArray.to_proto(
                    resource.details
                )
            )
        return res

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return None

        return InstanceReportInstanceHealthRecipeStepsStatus(
            code=Primitive.from_proto(resource.code),
            message=Primitive.from_proto(resource.message),
            details=InstanceReportInstanceHealthRecipeStepsStatusDetailsArray.from_proto(
                resource.details
            ),
        )


class InstanceReportInstanceHealthRecipeStepsStatusArray(object):
    @classmethod
    def to_proto(self, resources):
        if not resources:
            return resources
        return [
            InstanceReportInstanceHealthRecipeStepsStatus.to_proto(i) for i in resources
        ]

    @classmethod
    def from_proto(self, resources):
        return [
            InstanceReportInstanceHealthRecipeStepsStatus.from_proto(i)
            for i in resources
        ]


class InstanceReportInstanceHealthRecipeStepsStatusDetails(object):
    def __init__(self, type_url: str = None, value: str = None):
        self.type_url = type_url
        self.value = value

    @classmethod
    def to_proto(self, resource):
        if not resource:
            return None

        res = (
            instance_pb2.Tier2AlphaInstanceReportInstanceHealthRecipeStepsStatusDetails()
        )
        if Primitive.to_proto(resource.type_url):
            res.type_url = Primitive.to_proto(resource.type_url)
        if Primitive.to_proto(resource.value):
            res.value = Primitive.to_proto(resource.value)
        return res

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return None

        return InstanceReportInstanceHealthRecipeStepsStatusDetails(
            type_url=Primitive.from_proto(resource.type_url),
            value=Primitive.from_proto(resource.value),
        )


class InstanceReportInstanceHealthRecipeStepsStatusDetailsArray(object):
    @classmethod
    def to_proto(self, resources):
        if not resources:
            return resources
        return [
            InstanceReportInstanceHealthRecipeStepsStatusDetails.to_proto(i)
            for i in resources
        ]

    @classmethod
    def from_proto(self, resources):
        return [
            InstanceReportInstanceHealthRecipeStepsStatusDetails.from_proto(i)
            for i in resources
        ]


class InstanceReportInstanceHealthRecipeStepsQuotaRequestDeltas(object):
    def __init__(
        self,
        metric_name: str = None,
        amount: int = None,
        quota_location_name: str = None,
    ):
        self.metric_name = metric_name
        self.amount = amount
        self.quota_location_name = quota_location_name

    @classmethod
    def to_proto(self, resource):
        if not resource:
            return None

        res = (
            instance_pb2.Tier2AlphaInstanceReportInstanceHealthRecipeStepsQuotaRequestDeltas()
        )
        if Primitive.to_proto(resource.metric_name):
            res.metric_name = Primitive.to_proto(resource.metric_name)
        if Primitive.to_proto(resource.amount):
            res.amount = Primitive.to_proto(resource.amount)
        if Primitive.to_proto(resource.quota_location_name):
            res.quota_location_name = Primitive.to_proto(resource.quota_location_name)
        return res

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return None

        return InstanceReportInstanceHealthRecipeStepsQuotaRequestDeltas(
            metric_name=Primitive.from_proto(resource.metric_name),
            amount=Primitive.from_proto(resource.amount),
            quota_location_name=Primitive.from_proto(resource.quota_location_name),
        )


class InstanceReportInstanceHealthRecipeStepsQuotaRequestDeltasArray(object):
    @classmethod
    def to_proto(self, resources):
        if not resources:
            return resources
        return [
            InstanceReportInstanceHealthRecipeStepsQuotaRequestDeltas.to_proto(i)
            for i in resources
        ]

    @classmethod
    def from_proto(self, resources):
        return [
            InstanceReportInstanceHealthRecipeStepsQuotaRequestDeltas.from_proto(i)
            for i in resources
        ]


class InstanceReportInstanceHealthRecipeStepsPreprocessUpdate(object):
    def __init__(
        self, latency_slo_bucket_name: str = None, public_operation_metadata: str = None
    ):
        self.latency_slo_bucket_name = latency_slo_bucket_name
        self.public_operation_metadata = public_operation_metadata

    @classmethod
    def to_proto(self, resource):
        if not resource:
            return None

        res = (
            instance_pb2.Tier2AlphaInstanceReportInstanceHealthRecipeStepsPreprocessUpdate()
        )
        if Primitive.to_proto(resource.latency_slo_bucket_name):
            res.latency_slo_bucket_name = Primitive.to_proto(
                resource.latency_slo_bucket_name
            )
        if Primitive.to_proto(resource.public_operation_metadata):
            res.public_operation_metadata = Primitive.to_proto(
                resource.public_operation_metadata
            )
        return res

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return None

        return InstanceReportInstanceHealthRecipeStepsPreprocessUpdate(
            latency_slo_bucket_name=Primitive.from_proto(
                resource.latency_slo_bucket_name
            ),
            public_operation_metadata=Primitive.from_proto(
                resource.public_operation_metadata
            ),
        )


class InstanceReportInstanceHealthRecipeStepsPreprocessUpdateArray(object):
    @classmethod
    def to_proto(self, resources):
        if not resources:
            return resources
        return [
            InstanceReportInstanceHealthRecipeStepsPreprocessUpdate.to_proto(i)
            for i in resources
        ]

    @classmethod
    def from_proto(self, resources):
        return [
            InstanceReportInstanceHealthRecipeStepsPreprocessUpdate.from_proto(i)
            for i in resources
        ]


class InstanceReportInstanceHealthRecipeStepsRequestedTenantProject(object):
    def __init__(self, tag: str = None, folder: str = None, scope: str = None):
        self.tag = tag
        self.folder = folder
        self.scope = scope

    @classmethod
    def to_proto(self, resource):
        if not resource:
            return None

        res = (
            instance_pb2.Tier2AlphaInstanceReportInstanceHealthRecipeStepsRequestedTenantProject()
        )
        if Primitive.to_proto(resource.tag):
            res.tag = Primitive.to_proto(resource.tag)
        if Primitive.to_proto(resource.folder):
            res.folder = Primitive.to_proto(resource.folder)
        if InstanceReportInstanceHealthRecipeStepsRequestedTenantProjectScopeEnum.to_proto(
            resource.scope
        ):
            res.scope = InstanceReportInstanceHealthRecipeStepsRequestedTenantProjectScopeEnum.to_proto(
                resource.scope
            )
        return res

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return None

        return InstanceReportInstanceHealthRecipeStepsRequestedTenantProject(
            tag=Primitive.from_proto(resource.tag),
            folder=Primitive.from_proto(resource.folder),
            scope=InstanceReportInstanceHealthRecipeStepsRequestedTenantProjectScopeEnum.from_proto(
                resource.scope
            ),
        )


class InstanceReportInstanceHealthRecipeStepsRequestedTenantProjectArray(object):
    @classmethod
    def to_proto(self, resources):
        if not resources:
            return resources
        return [
            InstanceReportInstanceHealthRecipeStepsRequestedTenantProject.to_proto(i)
            for i in resources
        ]

    @classmethod
    def from_proto(self, resources):
        return [
            InstanceReportInstanceHealthRecipeStepsRequestedTenantProject.from_proto(i)
            for i in resources
        ]


class InstanceReportInstanceHealthRecipeStepsPermissionsInfo(object):
    def __init__(
        self,
        policy_name: dict = None,
        iam_permissions: list = None,
        resource_path: str = None,
        api_attrs: dict = None,
        policy_name_mode: str = None,
        resource: dict = None,
    ):
        self.policy_name = policy_name
        self.iam_permissions = iam_permissions
        self.resource_path = resource_path
        self.api_attrs = api_attrs
        self.policy_name_mode = policy_name_mode
        self.resource = resource

    @classmethod
    def to_proto(self, resource):
        if not resource:
            return None

        res = (
            instance_pb2.Tier2AlphaInstanceReportInstanceHealthRecipeStepsPermissionsInfo()
        )
        if InstanceReportInstanceHealthRecipeStepsPermissionsInfoPolicyName.to_proto(
            resource.policy_name
        ):
            res.policy_name.CopyFrom(
                InstanceReportInstanceHealthRecipeStepsPermissionsInfoPolicyName.to_proto(
                    resource.policy_name
                )
            )
        else:
            res.ClearField("policy_name")
        if InstanceReportInstanceHealthRecipeStepsPermissionsInfoIamPermissionsArray.to_proto(
            resource.iam_permissions
        ):
            res.iam_permissions.extend(
                InstanceReportInstanceHealthRecipeStepsPermissionsInfoIamPermissionsArray.to_proto(
                    resource.iam_permissions
                )
            )
        if Primitive.to_proto(resource.resource_path):
            res.resource_path = Primitive.to_proto(resource.resource_path)
        if InstanceGoogleprotobufstruct.to_proto(resource.api_attrs):
            res.api_attrs.CopyFrom(
                InstanceGoogleprotobufstruct.to_proto(resource.api_attrs)
            )
        else:
            res.ClearField("api_attrs")
        if InstanceReportInstanceHealthRecipeStepsPermissionsInfoPolicyNameModeEnum.to_proto(
            resource.policy_name_mode
        ):
            res.policy_name_mode = InstanceReportInstanceHealthRecipeStepsPermissionsInfoPolicyNameModeEnum.to_proto(
                resource.policy_name_mode
            )
        if InstanceReportInstanceHealthRecipeStepsPermissionsInfoResource.to_proto(
            resource.resource
        ):
            res.resource.CopyFrom(
                InstanceReportInstanceHealthRecipeStepsPermissionsInfoResource.to_proto(
                    resource.resource
                )
            )
        else:
            res.ClearField("resource")
        return res

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return None

        return InstanceReportInstanceHealthRecipeStepsPermissionsInfo(
            policy_name=InstanceReportInstanceHealthRecipeStepsPermissionsInfoPolicyName.from_proto(
                resource.policy_name
            ),
            iam_permissions=InstanceReportInstanceHealthRecipeStepsPermissionsInfoIamPermissionsArray.from_proto(
                resource.iam_permissions
            ),
            resource_path=Primitive.from_proto(resource.resource_path),
            api_attrs=InstanceGoogleprotobufstruct.from_proto(resource.api_attrs),
            policy_name_mode=InstanceReportInstanceHealthRecipeStepsPermissionsInfoPolicyNameModeEnum.from_proto(
                resource.policy_name_mode
            ),
            resource=InstanceReportInstanceHealthRecipeStepsPermissionsInfoResource.from_proto(
                resource.resource
            ),
        )


class InstanceReportInstanceHealthRecipeStepsPermissionsInfoArray(object):
    @classmethod
    def to_proto(self, resources):
        if not resources:
            return resources
        return [
            InstanceReportInstanceHealthRecipeStepsPermissionsInfo.to_proto(i)
            for i in resources
        ]

    @classmethod
    def from_proto(self, resources):
        return [
            InstanceReportInstanceHealthRecipeStepsPermissionsInfo.from_proto(i)
            for i in resources
        ]


class InstanceReportInstanceHealthRecipeStepsPermissionsInfoPolicyName(object):
    def __init__(self, type: str = None, id: str = None, region: str = None):
        self.type = type
        self.id = id
        self.region = region

    @classmethod
    def to_proto(self, resource):
        if not resource:
            return None

        res = (
            instance_pb2.Tier2AlphaInstanceReportInstanceHealthRecipeStepsPermissionsInfoPolicyName()
        )
        if Primitive.to_proto(resource.type):
            res.type = Primitive.to_proto(resource.type)
        if Primitive.to_proto(resource.id):
            res.id = Primitive.to_proto(resource.id)
        if Primitive.to_proto(resource.region):
            res.region = Primitive.to_proto(resource.region)
        return res

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return None

        return InstanceReportInstanceHealthRecipeStepsPermissionsInfoPolicyName(
            type=Primitive.from_proto(resource.type),
            id=Primitive.from_proto(resource.id),
            region=Primitive.from_proto(resource.region),
        )


class InstanceReportInstanceHealthRecipeStepsPermissionsInfoPolicyNameArray(object):
    @classmethod
    def to_proto(self, resources):
        if not resources:
            return resources
        return [
            InstanceReportInstanceHealthRecipeStepsPermissionsInfoPolicyName.to_proto(i)
            for i in resources
        ]

    @classmethod
    def from_proto(self, resources):
        return [
            InstanceReportInstanceHealthRecipeStepsPermissionsInfoPolicyName.from_proto(
                i
            )
            for i in resources
        ]


class InstanceReportInstanceHealthRecipeStepsPermissionsInfoIamPermissions(object):
    def __init__(self, permission: str = None):
        self.permission = permission

    @classmethod
    def to_proto(self, resource):
        if not resource:
            return None

        res = (
            instance_pb2.Tier2AlphaInstanceReportInstanceHealthRecipeStepsPermissionsInfoIamPermissions()
        )
        if Primitive.to_proto(resource.permission):
            res.permission = Primitive.to_proto(resource.permission)
        return res

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return None

        return InstanceReportInstanceHealthRecipeStepsPermissionsInfoIamPermissions(
            permission=Primitive.from_proto(resource.permission),
        )


class InstanceReportInstanceHealthRecipeStepsPermissionsInfoIamPermissionsArray(object):
    @classmethod
    def to_proto(self, resources):
        if not resources:
            return resources
        return [
            InstanceReportInstanceHealthRecipeStepsPermissionsInfoIamPermissions.to_proto(
                i
            )
            for i in resources
        ]

    @classmethod
    def from_proto(self, resources):
        return [
            InstanceReportInstanceHealthRecipeStepsPermissionsInfoIamPermissions.from_proto(
                i
            )
            for i in resources
        ]


class InstanceReportInstanceHealthRecipeStepsPermissionsInfoResource(object):
    def __init__(
        self,
        name: str = None,
        type: str = None,
        service: str = None,
        labels: dict = None,
    ):
        self.name = name
        self.type = type
        self.service = service
        self.labels = labels

    @classmethod
    def to_proto(self, resource):
        if not resource:
            return None

        res = (
            instance_pb2.Tier2AlphaInstanceReportInstanceHealthRecipeStepsPermissionsInfoResource()
        )
        if Primitive.to_proto(resource.name):
            res.name = Primitive.to_proto(resource.name)
        if Primitive.to_proto(resource.type):
            res.type = Primitive.to_proto(resource.type)
        if Primitive.to_proto(resource.service):
            res.service = Primitive.to_proto(resource.service)
        if Primitive.to_proto(resource.labels):
            res.labels = Primitive.to_proto(resource.labels)
        return res

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return None

        return InstanceReportInstanceHealthRecipeStepsPermissionsInfoResource(
            name=Primitive.from_proto(resource.name),
            type=Primitive.from_proto(resource.type),
            service=Primitive.from_proto(resource.service),
            labels=Primitive.from_proto(resource.labels),
        )


class InstanceReportInstanceHealthRecipeStepsPermissionsInfoResourceArray(object):
    @classmethod
    def to_proto(self, resources):
        if not resources:
            return resources
        return [
            InstanceReportInstanceHealthRecipeStepsPermissionsInfoResource.to_proto(i)
            for i in resources
        ]

    @classmethod
    def from_proto(self, resources):
        return [
            InstanceReportInstanceHealthRecipeStepsPermissionsInfoResource.from_proto(i)
            for i in resources
        ]


class InstanceReportInstanceHealthRecipeStepsKeyNotificationsUpdate(object):
    def __init__(self, key_notifications_info: dict = None):
        self.key_notifications_info = key_notifications_info

    @classmethod
    def to_proto(self, resource):
        if not resource:
            return None

        res = (
            instance_pb2.Tier2AlphaInstanceReportInstanceHealthRecipeStepsKeyNotificationsUpdate()
        )
        if InstanceReportInstanceHealthRecipeStepsKeyNotificationsUpdateKeyNotificationsInfo.to_proto(
            resource.key_notifications_info
        ):
            res.key_notifications_info.CopyFrom(
                InstanceReportInstanceHealthRecipeStepsKeyNotificationsUpdateKeyNotificationsInfo.to_proto(
                    resource.key_notifications_info
                )
            )
        else:
            res.ClearField("key_notifications_info")
        return res

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return None

        return InstanceReportInstanceHealthRecipeStepsKeyNotificationsUpdate(
            key_notifications_info=InstanceReportInstanceHealthRecipeStepsKeyNotificationsUpdateKeyNotificationsInfo.from_proto(
                resource.key_notifications_info
            ),
        )


class InstanceReportInstanceHealthRecipeStepsKeyNotificationsUpdateArray(object):
    @classmethod
    def to_proto(self, resources):
        if not resources:
            return resources
        return [
            InstanceReportInstanceHealthRecipeStepsKeyNotificationsUpdate.to_proto(i)
            for i in resources
        ]

    @classmethod
    def from_proto(self, resources):
        return [
            InstanceReportInstanceHealthRecipeStepsKeyNotificationsUpdate.from_proto(i)
            for i in resources
        ]


class InstanceReportInstanceHealthRecipeStepsKeyNotificationsUpdateKeyNotificationsInfo(
    object
):
    def __init__(
        self,
        data_version: int = None,
        delegate: str = None,
        key_notification_configs: list = None,
    ):
        self.data_version = data_version
        self.delegate = delegate
        self.key_notification_configs = key_notification_configs

    @classmethod
    def to_proto(self, resource):
        if not resource:
            return None

        res = (
            instance_pb2.Tier2AlphaInstanceReportInstanceHealthRecipeStepsKeyNotificationsUpdateKeyNotificationsInfo()
        )
        if Primitive.to_proto(resource.data_version):
            res.data_version = Primitive.to_proto(resource.data_version)
        if Primitive.to_proto(resource.delegate):
            res.delegate = Primitive.to_proto(resource.delegate)
        if InstanceReportInstanceHealthRecipeStepsKeyNotificationsUpdateKeyNotificationsInfoKeyNotificationConfigsArray.to_proto(
            resource.key_notification_configs
        ):
            res.key_notification_configs.extend(
                InstanceReportInstanceHealthRecipeStepsKeyNotificationsUpdateKeyNotificationsInfoKeyNotificationConfigsArray.to_proto(
                    resource.key_notification_configs
                )
            )
        return res

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return None

        return InstanceReportInstanceHealthRecipeStepsKeyNotificationsUpdateKeyNotificationsInfo(
            data_version=Primitive.from_proto(resource.data_version),
            delegate=Primitive.from_proto(resource.delegate),
            key_notification_configs=InstanceReportInstanceHealthRecipeStepsKeyNotificationsUpdateKeyNotificationsInfoKeyNotificationConfigsArray.from_proto(
                resource.key_notification_configs
            ),
        )


class InstanceReportInstanceHealthRecipeStepsKeyNotificationsUpdateKeyNotificationsInfoArray(
    object
):
    @classmethod
    def to_proto(self, resources):
        if not resources:
            return resources
        return [
            InstanceReportInstanceHealthRecipeStepsKeyNotificationsUpdateKeyNotificationsInfo.to_proto(
                i
            )
            for i in resources
        ]

    @classmethod
    def from_proto(self, resources):
        return [
            InstanceReportInstanceHealthRecipeStepsKeyNotificationsUpdateKeyNotificationsInfo.from_proto(
                i
            )
            for i in resources
        ]


class InstanceReportInstanceHealthRecipeStepsKeyNotificationsUpdateKeyNotificationsInfoKeyNotificationConfigs(
    object
):
    def __init__(
        self,
        key_or_version_name: str = None,
        grant: str = None,
        delegator_gaia_id: int = None,
    ):
        self.key_or_version_name = key_or_version_name
        self.grant = grant
        self.delegator_gaia_id = delegator_gaia_id

    @classmethod
    def to_proto(self, resource):
        if not resource:
            return None

        res = (
            instance_pb2.Tier2AlphaInstanceReportInstanceHealthRecipeStepsKeyNotificationsUpdateKeyNotificationsInfoKeyNotificationConfigs()
        )
        if Primitive.to_proto(resource.key_or_version_name):
            res.key_or_version_name = Primitive.to_proto(resource.key_or_version_name)
        if Primitive.to_proto(resource.grant):
            res.grant = Primitive.to_proto(resource.grant)
        if Primitive.to_proto(resource.delegator_gaia_id):
            res.delegator_gaia_id = Primitive.to_proto(resource.delegator_gaia_id)
        return res

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return None

        return InstanceReportInstanceHealthRecipeStepsKeyNotificationsUpdateKeyNotificationsInfoKeyNotificationConfigs(
            key_or_version_name=Primitive.from_proto(resource.key_or_version_name),
            grant=Primitive.from_proto(resource.grant),
            delegator_gaia_id=Primitive.from_proto(resource.delegator_gaia_id),
        )


class InstanceReportInstanceHealthRecipeStepsKeyNotificationsUpdateKeyNotificationsInfoKeyNotificationConfigsArray(
    object
):
    @classmethod
    def to_proto(self, resources):
        if not resources:
            return resources
        return [
            InstanceReportInstanceHealthRecipeStepsKeyNotificationsUpdateKeyNotificationsInfoKeyNotificationConfigs.to_proto(
                i
            )
            for i in resources
        ]

    @classmethod
    def from_proto(self, resources):
        return [
            InstanceReportInstanceHealthRecipeStepsKeyNotificationsUpdateKeyNotificationsInfoKeyNotificationConfigs.from_proto(
                i
            )
            for i in resources
        ]


class InstancePreprocessGetRecipe(object):
    def __init__(
        self,
        steps: list = None,
        honor_cancel_request: bool = None,
        ignore_recipe_after: int = None,
        verify_deadline_seconds_below: float = None,
        populate_operation_result: bool = None,
        readonly_recipe_start_time: str = None,
        resource_names_stored_in_clh_with_delay: list = None,
        delay_to_store_resources_in_clh_db_nanos: int = None,
    ):
        self.steps = steps
        self.honor_cancel_request = honor_cancel_request
        self.ignore_recipe_after = ignore_recipe_after
        self.verify_deadline_seconds_below = verify_deadline_seconds_below
        self.populate_operation_result = populate_operation_result
        self.readonly_recipe_start_time = readonly_recipe_start_time
        self.resource_names_stored_in_clh_with_delay = (
            resource_names_stored_in_clh_with_delay
        )
        self.delay_to_store_resources_in_clh_db_nanos = (
            delay_to_store_resources_in_clh_db_nanos
        )

    @classmethod
    def to_proto(self, resource):
        if not resource:
            return None

        res = instance_pb2.Tier2AlphaInstancePreprocessGetRecipe()
        if InstancePreprocessGetRecipeStepsArray.to_proto(resource.steps):
            res.steps.extend(
                InstancePreprocessGetRecipeStepsArray.to_proto(resource.steps)
            )
        if Primitive.to_proto(resource.honor_cancel_request):
            res.honor_cancel_request = Primitive.to_proto(resource.honor_cancel_request)
        if Primitive.to_proto(resource.ignore_recipe_after):
            res.ignore_recipe_after = Primitive.to_proto(resource.ignore_recipe_after)
        if Primitive.to_proto(resource.verify_deadline_seconds_below):
            res.verify_deadline_seconds_below = Primitive.to_proto(
                resource.verify_deadline_seconds_below
            )
        if Primitive.to_proto(resource.populate_operation_result):
            res.populate_operation_result = Primitive.to_proto(
                resource.populate_operation_result
            )
        if Primitive.to_proto(resource.readonly_recipe_start_time):
            res.readonly_recipe_start_time = Primitive.to_proto(
                resource.readonly_recipe_start_time
            )
        if Primitive.to_proto(resource.resource_names_stored_in_clh_with_delay):
            res.resource_names_stored_in_clh_with_delay.extend(
                Primitive.to_proto(resource.resource_names_stored_in_clh_with_delay)
            )
        if Primitive.to_proto(resource.delay_to_store_resources_in_clh_db_nanos):
            res.delay_to_store_resources_in_clh_db_nanos = Primitive.to_proto(
                resource.delay_to_store_resources_in_clh_db_nanos
            )
        return res

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return None

        return InstancePreprocessGetRecipe(
            steps=InstancePreprocessGetRecipeStepsArray.from_proto(resource.steps),
            honor_cancel_request=Primitive.from_proto(resource.honor_cancel_request),
            ignore_recipe_after=Primitive.from_proto(resource.ignore_recipe_after),
            verify_deadline_seconds_below=Primitive.from_proto(
                resource.verify_deadline_seconds_below
            ),
            populate_operation_result=Primitive.from_proto(
                resource.populate_operation_result
            ),
            readonly_recipe_start_time=Primitive.from_proto(
                resource.readonly_recipe_start_time
            ),
            resource_names_stored_in_clh_with_delay=Primitive.from_proto(
                resource.resource_names_stored_in_clh_with_delay
            ),
            delay_to_store_resources_in_clh_db_nanos=Primitive.from_proto(
                resource.delay_to_store_resources_in_clh_db_nanos
            ),
        )


class InstancePreprocessGetRecipeArray(object):
    @classmethod
    def to_proto(self, resources):
        if not resources:
            return resources
        return [InstancePreprocessGetRecipe.to_proto(i) for i in resources]

    @classmethod
    def from_proto(self, resources):
        return [InstancePreprocessGetRecipe.from_proto(i) for i in resources]


class InstancePreprocessGetRecipeSteps(object):
    def __init__(
        self,
        relative_time: int = None,
        sleep_duration: int = None,
        action: str = None,
        status: dict = None,
        error_space: str = None,
        p4_service_account: str = None,
        resource_metadata_size: int = None,
        description: str = None,
        updated_repeat_operation_delay_sec: float = None,
        quota_request_deltas: list = None,
        preprocess_update: dict = None,
        public_operation_metadata: str = None,
        requested_tenant_project: dict = None,
        permissions_info: list = None,
        key_notifications_update: dict = None,
        clh_data_update_time: str = None,
        public_error_message: str = None,
    ):
        self.relative_time = relative_time
        self.sleep_duration = sleep_duration
        self.action = action
        self.status = status
        self.error_space = error_space
        self.p4_service_account = p4_service_account
        self.resource_metadata_size = resource_metadata_size
        self.description = description
        self.updated_repeat_operation_delay_sec = updated_repeat_operation_delay_sec
        self.quota_request_deltas = quota_request_deltas
        self.preprocess_update = preprocess_update
        self.public_operation_metadata = public_operation_metadata
        self.requested_tenant_project = requested_tenant_project
        self.permissions_info = permissions_info
        self.key_notifications_update = key_notifications_update
        self.clh_data_update_time = clh_data_update_time
        self.public_error_message = public_error_message

    @classmethod
    def to_proto(self, resource):
        if not resource:
            return None

        res = instance_pb2.Tier2AlphaInstancePreprocessGetRecipeSteps()
        if Primitive.to_proto(resource.relative_time):
            res.relative_time = Primitive.to_proto(resource.relative_time)
        if Primitive.to_proto(resource.sleep_duration):
            res.sleep_duration = Primitive.to_proto(resource.sleep_duration)
        if InstancePreprocessGetRecipeStepsActionEnum.to_proto(resource.action):
            res.action = InstancePreprocessGetRecipeStepsActionEnum.to_proto(
                resource.action
            )
        if InstancePreprocessGetRecipeStepsStatus.to_proto(resource.status):
            res.status.CopyFrom(
                InstancePreprocessGetRecipeStepsStatus.to_proto(resource.status)
            )
        else:
            res.ClearField("status")
        if Primitive.to_proto(resource.error_space):
            res.error_space = Primitive.to_proto(resource.error_space)
        if Primitive.to_proto(resource.p4_service_account):
            res.p4_service_account = Primitive.to_proto(resource.p4_service_account)
        if Primitive.to_proto(resource.resource_metadata_size):
            res.resource_metadata_size = Primitive.to_proto(
                resource.resource_metadata_size
            )
        if Primitive.to_proto(resource.description):
            res.description = Primitive.to_proto(resource.description)
        if Primitive.to_proto(resource.updated_repeat_operation_delay_sec):
            res.updated_repeat_operation_delay_sec = Primitive.to_proto(
                resource.updated_repeat_operation_delay_sec
            )
        if InstancePreprocessGetRecipeStepsQuotaRequestDeltasArray.to_proto(
            resource.quota_request_deltas
        ):
            res.quota_request_deltas.extend(
                InstancePreprocessGetRecipeStepsQuotaRequestDeltasArray.to_proto(
                    resource.quota_request_deltas
                )
            )
        if InstancePreprocessGetRecipeStepsPreprocessUpdate.to_proto(
            resource.preprocess_update
        ):
            res.preprocess_update.CopyFrom(
                InstancePreprocessGetRecipeStepsPreprocessUpdate.to_proto(
                    resource.preprocess_update
                )
            )
        else:
            res.ClearField("preprocess_update")
        if Primitive.to_proto(resource.public_operation_metadata):
            res.public_operation_metadata = Primitive.to_proto(
                resource.public_operation_metadata
            )
        if InstancePreprocessGetRecipeStepsRequestedTenantProject.to_proto(
            resource.requested_tenant_project
        ):
            res.requested_tenant_project.CopyFrom(
                InstancePreprocessGetRecipeStepsRequestedTenantProject.to_proto(
                    resource.requested_tenant_project
                )
            )
        else:
            res.ClearField("requested_tenant_project")
        if InstancePreprocessGetRecipeStepsPermissionsInfoArray.to_proto(
            resource.permissions_info
        ):
            res.permissions_info.extend(
                InstancePreprocessGetRecipeStepsPermissionsInfoArray.to_proto(
                    resource.permissions_info
                )
            )
        if InstancePreprocessGetRecipeStepsKeyNotificationsUpdate.to_proto(
            resource.key_notifications_update
        ):
            res.key_notifications_update.CopyFrom(
                InstancePreprocessGetRecipeStepsKeyNotificationsUpdate.to_proto(
                    resource.key_notifications_update
                )
            )
        else:
            res.ClearField("key_notifications_update")
        if Primitive.to_proto(resource.clh_data_update_time):
            res.clh_data_update_time = Primitive.to_proto(resource.clh_data_update_time)
        if Primitive.to_proto(resource.public_error_message):
            res.public_error_message = Primitive.to_proto(resource.public_error_message)
        return res

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return None

        return InstancePreprocessGetRecipeSteps(
            relative_time=Primitive.from_proto(resource.relative_time),
            sleep_duration=Primitive.from_proto(resource.sleep_duration),
            action=InstancePreprocessGetRecipeStepsActionEnum.from_proto(
                resource.action
            ),
            status=InstancePreprocessGetRecipeStepsStatus.from_proto(resource.status),
            error_space=Primitive.from_proto(resource.error_space),
            p4_service_account=Primitive.from_proto(resource.p4_service_account),
            resource_metadata_size=Primitive.from_proto(
                resource.resource_metadata_size
            ),
            description=Primitive.from_proto(resource.description),
            updated_repeat_operation_delay_sec=Primitive.from_proto(
                resource.updated_repeat_operation_delay_sec
            ),
            quota_request_deltas=InstancePreprocessGetRecipeStepsQuotaRequestDeltasArray.from_proto(
                resource.quota_request_deltas
            ),
            preprocess_update=InstancePreprocessGetRecipeStepsPreprocessUpdate.from_proto(
                resource.preprocess_update
            ),
            public_operation_metadata=Primitive.from_proto(
                resource.public_operation_metadata
            ),
            requested_tenant_project=InstancePreprocessGetRecipeStepsRequestedTenantProject.from_proto(
                resource.requested_tenant_project
            ),
            permissions_info=InstancePreprocessGetRecipeStepsPermissionsInfoArray.from_proto(
                resource.permissions_info
            ),
            key_notifications_update=InstancePreprocessGetRecipeStepsKeyNotificationsUpdate.from_proto(
                resource.key_notifications_update
            ),
            clh_data_update_time=Primitive.from_proto(resource.clh_data_update_time),
            public_error_message=Primitive.from_proto(resource.public_error_message),
        )


class InstancePreprocessGetRecipeStepsArray(object):
    @classmethod
    def to_proto(self, resources):
        if not resources:
            return resources
        return [InstancePreprocessGetRecipeSteps.to_proto(i) for i in resources]

    @classmethod
    def from_proto(self, resources):
        return [InstancePreprocessGetRecipeSteps.from_proto(i) for i in resources]


class InstancePreprocessGetRecipeStepsStatus(object):
    def __init__(self, code: int = None, message: str = None, details: list = None):
        self.code = code
        self.message = message
        self.details = details

    @classmethod
    def to_proto(self, resource):
        if not resource:
            return None

        res = instance_pb2.Tier2AlphaInstancePreprocessGetRecipeStepsStatus()
        if Primitive.to_proto(resource.code):
            res.code = Primitive.to_proto(resource.code)
        if Primitive.to_proto(resource.message):
            res.message = Primitive.to_proto(resource.message)
        if InstancePreprocessGetRecipeStepsStatusDetailsArray.to_proto(
            resource.details
        ):
            res.details.extend(
                InstancePreprocessGetRecipeStepsStatusDetailsArray.to_proto(
                    resource.details
                )
            )
        return res

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return None

        return InstancePreprocessGetRecipeStepsStatus(
            code=Primitive.from_proto(resource.code),
            message=Primitive.from_proto(resource.message),
            details=InstancePreprocessGetRecipeStepsStatusDetailsArray.from_proto(
                resource.details
            ),
        )


class InstancePreprocessGetRecipeStepsStatusArray(object):
    @classmethod
    def to_proto(self, resources):
        if not resources:
            return resources
        return [InstancePreprocessGetRecipeStepsStatus.to_proto(i) for i in resources]

    @classmethod
    def from_proto(self, resources):
        return [InstancePreprocessGetRecipeStepsStatus.from_proto(i) for i in resources]


class InstancePreprocessGetRecipeStepsStatusDetails(object):
    def __init__(self, type_url: str = None, value: str = None):
        self.type_url = type_url
        self.value = value

    @classmethod
    def to_proto(self, resource):
        if not resource:
            return None

        res = instance_pb2.Tier2AlphaInstancePreprocessGetRecipeStepsStatusDetails()
        if Primitive.to_proto(resource.type_url):
            res.type_url = Primitive.to_proto(resource.type_url)
        if Primitive.to_proto(resource.value):
            res.value = Primitive.to_proto(resource.value)
        return res

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return None

        return InstancePreprocessGetRecipeStepsStatusDetails(
            type_url=Primitive.from_proto(resource.type_url),
            value=Primitive.from_proto(resource.value),
        )


class InstancePreprocessGetRecipeStepsStatusDetailsArray(object):
    @classmethod
    def to_proto(self, resources):
        if not resources:
            return resources
        return [
            InstancePreprocessGetRecipeStepsStatusDetails.to_proto(i) for i in resources
        ]

    @classmethod
    def from_proto(self, resources):
        return [
            InstancePreprocessGetRecipeStepsStatusDetails.from_proto(i)
            for i in resources
        ]


class InstancePreprocessGetRecipeStepsQuotaRequestDeltas(object):
    def __init__(
        self,
        metric_name: str = None,
        amount: int = None,
        quota_location_name: str = None,
    ):
        self.metric_name = metric_name
        self.amount = amount
        self.quota_location_name = quota_location_name

    @classmethod
    def to_proto(self, resource):
        if not resource:
            return None

        res = (
            instance_pb2.Tier2AlphaInstancePreprocessGetRecipeStepsQuotaRequestDeltas()
        )
        if Primitive.to_proto(resource.metric_name):
            res.metric_name = Primitive.to_proto(resource.metric_name)
        if Primitive.to_proto(resource.amount):
            res.amount = Primitive.to_proto(resource.amount)
        if Primitive.to_proto(resource.quota_location_name):
            res.quota_location_name = Primitive.to_proto(resource.quota_location_name)
        return res

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return None

        return InstancePreprocessGetRecipeStepsQuotaRequestDeltas(
            metric_name=Primitive.from_proto(resource.metric_name),
            amount=Primitive.from_proto(resource.amount),
            quota_location_name=Primitive.from_proto(resource.quota_location_name),
        )


class InstancePreprocessGetRecipeStepsQuotaRequestDeltasArray(object):
    @classmethod
    def to_proto(self, resources):
        if not resources:
            return resources
        return [
            InstancePreprocessGetRecipeStepsQuotaRequestDeltas.to_proto(i)
            for i in resources
        ]

    @classmethod
    def from_proto(self, resources):
        return [
            InstancePreprocessGetRecipeStepsQuotaRequestDeltas.from_proto(i)
            for i in resources
        ]


class InstancePreprocessGetRecipeStepsPreprocessUpdate(object):
    def __init__(
        self, latency_slo_bucket_name: str = None, public_operation_metadata: str = None
    ):
        self.latency_slo_bucket_name = latency_slo_bucket_name
        self.public_operation_metadata = public_operation_metadata

    @classmethod
    def to_proto(self, resource):
        if not resource:
            return None

        res = instance_pb2.Tier2AlphaInstancePreprocessGetRecipeStepsPreprocessUpdate()
        if Primitive.to_proto(resource.latency_slo_bucket_name):
            res.latency_slo_bucket_name = Primitive.to_proto(
                resource.latency_slo_bucket_name
            )
        if Primitive.to_proto(resource.public_operation_metadata):
            res.public_operation_metadata = Primitive.to_proto(
                resource.public_operation_metadata
            )
        return res

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return None

        return InstancePreprocessGetRecipeStepsPreprocessUpdate(
            latency_slo_bucket_name=Primitive.from_proto(
                resource.latency_slo_bucket_name
            ),
            public_operation_metadata=Primitive.from_proto(
                resource.public_operation_metadata
            ),
        )


class InstancePreprocessGetRecipeStepsPreprocessUpdateArray(object):
    @classmethod
    def to_proto(self, resources):
        if not resources:
            return resources
        return [
            InstancePreprocessGetRecipeStepsPreprocessUpdate.to_proto(i)
            for i in resources
        ]

    @classmethod
    def from_proto(self, resources):
        return [
            InstancePreprocessGetRecipeStepsPreprocessUpdate.from_proto(i)
            for i in resources
        ]


class InstancePreprocessGetRecipeStepsRequestedTenantProject(object):
    def __init__(self, tag: str = None, folder: str = None, scope: str = None):
        self.tag = tag
        self.folder = folder
        self.scope = scope

    @classmethod
    def to_proto(self, resource):
        if not resource:
            return None

        res = (
            instance_pb2.Tier2AlphaInstancePreprocessGetRecipeStepsRequestedTenantProject()
        )
        if Primitive.to_proto(resource.tag):
            res.tag = Primitive.to_proto(resource.tag)
        if Primitive.to_proto(resource.folder):
            res.folder = Primitive.to_proto(resource.folder)
        if InstancePreprocessGetRecipeStepsRequestedTenantProjectScopeEnum.to_proto(
            resource.scope
        ):
            res.scope = InstancePreprocessGetRecipeStepsRequestedTenantProjectScopeEnum.to_proto(
                resource.scope
            )
        return res

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return None

        return InstancePreprocessGetRecipeStepsRequestedTenantProject(
            tag=Primitive.from_proto(resource.tag),
            folder=Primitive.from_proto(resource.folder),
            scope=InstancePreprocessGetRecipeStepsRequestedTenantProjectScopeEnum.from_proto(
                resource.scope
            ),
        )


class InstancePreprocessGetRecipeStepsRequestedTenantProjectArray(object):
    @classmethod
    def to_proto(self, resources):
        if not resources:
            return resources
        return [
            InstancePreprocessGetRecipeStepsRequestedTenantProject.to_proto(i)
            for i in resources
        ]

    @classmethod
    def from_proto(self, resources):
        return [
            InstancePreprocessGetRecipeStepsRequestedTenantProject.from_proto(i)
            for i in resources
        ]


class InstancePreprocessGetRecipeStepsPermissionsInfo(object):
    def __init__(
        self,
        policy_name: dict = None,
        iam_permissions: list = None,
        resource_path: str = None,
        api_attrs: dict = None,
        policy_name_mode: str = None,
        resource: dict = None,
    ):
        self.policy_name = policy_name
        self.iam_permissions = iam_permissions
        self.resource_path = resource_path
        self.api_attrs = api_attrs
        self.policy_name_mode = policy_name_mode
        self.resource = resource

    @classmethod
    def to_proto(self, resource):
        if not resource:
            return None

        res = instance_pb2.Tier2AlphaInstancePreprocessGetRecipeStepsPermissionsInfo()
        if InstancePreprocessGetRecipeStepsPermissionsInfoPolicyName.to_proto(
            resource.policy_name
        ):
            res.policy_name.CopyFrom(
                InstancePreprocessGetRecipeStepsPermissionsInfoPolicyName.to_proto(
                    resource.policy_name
                )
            )
        else:
            res.ClearField("policy_name")
        if InstancePreprocessGetRecipeStepsPermissionsInfoIamPermissionsArray.to_proto(
            resource.iam_permissions
        ):
            res.iam_permissions.extend(
                InstancePreprocessGetRecipeStepsPermissionsInfoIamPermissionsArray.to_proto(
                    resource.iam_permissions
                )
            )
        if Primitive.to_proto(resource.resource_path):
            res.resource_path = Primitive.to_proto(resource.resource_path)
        if InstanceGoogleprotobufstruct.to_proto(resource.api_attrs):
            res.api_attrs.CopyFrom(
                InstanceGoogleprotobufstruct.to_proto(resource.api_attrs)
            )
        else:
            res.ClearField("api_attrs")
        if InstancePreprocessGetRecipeStepsPermissionsInfoPolicyNameModeEnum.to_proto(
            resource.policy_name_mode
        ):
            res.policy_name_mode = InstancePreprocessGetRecipeStepsPermissionsInfoPolicyNameModeEnum.to_proto(
                resource.policy_name_mode
            )
        if InstancePreprocessGetRecipeStepsPermissionsInfoResource.to_proto(
            resource.resource
        ):
            res.resource.CopyFrom(
                InstancePreprocessGetRecipeStepsPermissionsInfoResource.to_proto(
                    resource.resource
                )
            )
        else:
            res.ClearField("resource")
        return res

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return None

        return InstancePreprocessGetRecipeStepsPermissionsInfo(
            policy_name=InstancePreprocessGetRecipeStepsPermissionsInfoPolicyName.from_proto(
                resource.policy_name
            ),
            iam_permissions=InstancePreprocessGetRecipeStepsPermissionsInfoIamPermissionsArray.from_proto(
                resource.iam_permissions
            ),
            resource_path=Primitive.from_proto(resource.resource_path),
            api_attrs=InstanceGoogleprotobufstruct.from_proto(resource.api_attrs),
            policy_name_mode=InstancePreprocessGetRecipeStepsPermissionsInfoPolicyNameModeEnum.from_proto(
                resource.policy_name_mode
            ),
            resource=InstancePreprocessGetRecipeStepsPermissionsInfoResource.from_proto(
                resource.resource
            ),
        )


class InstancePreprocessGetRecipeStepsPermissionsInfoArray(object):
    @classmethod
    def to_proto(self, resources):
        if not resources:
            return resources
        return [
            InstancePreprocessGetRecipeStepsPermissionsInfo.to_proto(i)
            for i in resources
        ]

    @classmethod
    def from_proto(self, resources):
        return [
            InstancePreprocessGetRecipeStepsPermissionsInfo.from_proto(i)
            for i in resources
        ]


class InstancePreprocessGetRecipeStepsPermissionsInfoPolicyName(object):
    def __init__(self, type: str = None, id: str = None, region: str = None):
        self.type = type
        self.id = id
        self.region = region

    @classmethod
    def to_proto(self, resource):
        if not resource:
            return None

        res = (
            instance_pb2.Tier2AlphaInstancePreprocessGetRecipeStepsPermissionsInfoPolicyName()
        )
        if Primitive.to_proto(resource.type):
            res.type = Primitive.to_proto(resource.type)
        if Primitive.to_proto(resource.id):
            res.id = Primitive.to_proto(resource.id)
        if Primitive.to_proto(resource.region):
            res.region = Primitive.to_proto(resource.region)
        return res

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return None

        return InstancePreprocessGetRecipeStepsPermissionsInfoPolicyName(
            type=Primitive.from_proto(resource.type),
            id=Primitive.from_proto(resource.id),
            region=Primitive.from_proto(resource.region),
        )


class InstancePreprocessGetRecipeStepsPermissionsInfoPolicyNameArray(object):
    @classmethod
    def to_proto(self, resources):
        if not resources:
            return resources
        return [
            InstancePreprocessGetRecipeStepsPermissionsInfoPolicyName.to_proto(i)
            for i in resources
        ]

    @classmethod
    def from_proto(self, resources):
        return [
            InstancePreprocessGetRecipeStepsPermissionsInfoPolicyName.from_proto(i)
            for i in resources
        ]


class InstancePreprocessGetRecipeStepsPermissionsInfoIamPermissions(object):
    def __init__(self, permission: str = None):
        self.permission = permission

    @classmethod
    def to_proto(self, resource):
        if not resource:
            return None

        res = (
            instance_pb2.Tier2AlphaInstancePreprocessGetRecipeStepsPermissionsInfoIamPermissions()
        )
        if Primitive.to_proto(resource.permission):
            res.permission = Primitive.to_proto(resource.permission)
        return res

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return None

        return InstancePreprocessGetRecipeStepsPermissionsInfoIamPermissions(
            permission=Primitive.from_proto(resource.permission),
        )


class InstancePreprocessGetRecipeStepsPermissionsInfoIamPermissionsArray(object):
    @classmethod
    def to_proto(self, resources):
        if not resources:
            return resources
        return [
            InstancePreprocessGetRecipeStepsPermissionsInfoIamPermissions.to_proto(i)
            for i in resources
        ]

    @classmethod
    def from_proto(self, resources):
        return [
            InstancePreprocessGetRecipeStepsPermissionsInfoIamPermissions.from_proto(i)
            for i in resources
        ]


class InstancePreprocessGetRecipeStepsPermissionsInfoResource(object):
    def __init__(
        self,
        name: str = None,
        type: str = None,
        service: str = None,
        labels: dict = None,
    ):
        self.name = name
        self.type = type
        self.service = service
        self.labels = labels

    @classmethod
    def to_proto(self, resource):
        if not resource:
            return None

        res = (
            instance_pb2.Tier2AlphaInstancePreprocessGetRecipeStepsPermissionsInfoResource()
        )
        if Primitive.to_proto(resource.name):
            res.name = Primitive.to_proto(resource.name)
        if Primitive.to_proto(resource.type):
            res.type = Primitive.to_proto(resource.type)
        if Primitive.to_proto(resource.service):
            res.service = Primitive.to_proto(resource.service)
        if Primitive.to_proto(resource.labels):
            res.labels = Primitive.to_proto(resource.labels)
        return res

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return None

        return InstancePreprocessGetRecipeStepsPermissionsInfoResource(
            name=Primitive.from_proto(resource.name),
            type=Primitive.from_proto(resource.type),
            service=Primitive.from_proto(resource.service),
            labels=Primitive.from_proto(resource.labels),
        )


class InstancePreprocessGetRecipeStepsPermissionsInfoResourceArray(object):
    @classmethod
    def to_proto(self, resources):
        if not resources:
            return resources
        return [
            InstancePreprocessGetRecipeStepsPermissionsInfoResource.to_proto(i)
            for i in resources
        ]

    @classmethod
    def from_proto(self, resources):
        return [
            InstancePreprocessGetRecipeStepsPermissionsInfoResource.from_proto(i)
            for i in resources
        ]


class InstancePreprocessGetRecipeStepsKeyNotificationsUpdate(object):
    def __init__(self, key_notifications_info: dict = None):
        self.key_notifications_info = key_notifications_info

    @classmethod
    def to_proto(self, resource):
        if not resource:
            return None

        res = (
            instance_pb2.Tier2AlphaInstancePreprocessGetRecipeStepsKeyNotificationsUpdate()
        )
        if InstancePreprocessGetRecipeStepsKeyNotificationsUpdateKeyNotificationsInfo.to_proto(
            resource.key_notifications_info
        ):
            res.key_notifications_info.CopyFrom(
                InstancePreprocessGetRecipeStepsKeyNotificationsUpdateKeyNotificationsInfo.to_proto(
                    resource.key_notifications_info
                )
            )
        else:
            res.ClearField("key_notifications_info")
        return res

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return None

        return InstancePreprocessGetRecipeStepsKeyNotificationsUpdate(
            key_notifications_info=InstancePreprocessGetRecipeStepsKeyNotificationsUpdateKeyNotificationsInfo.from_proto(
                resource.key_notifications_info
            ),
        )


class InstancePreprocessGetRecipeStepsKeyNotificationsUpdateArray(object):
    @classmethod
    def to_proto(self, resources):
        if not resources:
            return resources
        return [
            InstancePreprocessGetRecipeStepsKeyNotificationsUpdate.to_proto(i)
            for i in resources
        ]

    @classmethod
    def from_proto(self, resources):
        return [
            InstancePreprocessGetRecipeStepsKeyNotificationsUpdate.from_proto(i)
            for i in resources
        ]


class InstancePreprocessGetRecipeStepsKeyNotificationsUpdateKeyNotificationsInfo(
    object
):
    def __init__(
        self,
        data_version: int = None,
        delegate: str = None,
        key_notification_configs: list = None,
    ):
        self.data_version = data_version
        self.delegate = delegate
        self.key_notification_configs = key_notification_configs

    @classmethod
    def to_proto(self, resource):
        if not resource:
            return None

        res = (
            instance_pb2.Tier2AlphaInstancePreprocessGetRecipeStepsKeyNotificationsUpdateKeyNotificationsInfo()
        )
        if Primitive.to_proto(resource.data_version):
            res.data_version = Primitive.to_proto(resource.data_version)
        if Primitive.to_proto(resource.delegate):
            res.delegate = Primitive.to_proto(resource.delegate)
        if InstancePreprocessGetRecipeStepsKeyNotificationsUpdateKeyNotificationsInfoKeyNotificationConfigsArray.to_proto(
            resource.key_notification_configs
        ):
            res.key_notification_configs.extend(
                InstancePreprocessGetRecipeStepsKeyNotificationsUpdateKeyNotificationsInfoKeyNotificationConfigsArray.to_proto(
                    resource.key_notification_configs
                )
            )
        return res

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return None

        return InstancePreprocessGetRecipeStepsKeyNotificationsUpdateKeyNotificationsInfo(
            data_version=Primitive.from_proto(resource.data_version),
            delegate=Primitive.from_proto(resource.delegate),
            key_notification_configs=InstancePreprocessGetRecipeStepsKeyNotificationsUpdateKeyNotificationsInfoKeyNotificationConfigsArray.from_proto(
                resource.key_notification_configs
            ),
        )


class InstancePreprocessGetRecipeStepsKeyNotificationsUpdateKeyNotificationsInfoArray(
    object
):
    @classmethod
    def to_proto(self, resources):
        if not resources:
            return resources
        return [
            InstancePreprocessGetRecipeStepsKeyNotificationsUpdateKeyNotificationsInfo.to_proto(
                i
            )
            for i in resources
        ]

    @classmethod
    def from_proto(self, resources):
        return [
            InstancePreprocessGetRecipeStepsKeyNotificationsUpdateKeyNotificationsInfo.from_proto(
                i
            )
            for i in resources
        ]


class InstancePreprocessGetRecipeStepsKeyNotificationsUpdateKeyNotificationsInfoKeyNotificationConfigs(
    object
):
    def __init__(
        self,
        key_or_version_name: str = None,
        grant: str = None,
        delegator_gaia_id: int = None,
    ):
        self.key_or_version_name = key_or_version_name
        self.grant = grant
        self.delegator_gaia_id = delegator_gaia_id

    @classmethod
    def to_proto(self, resource):
        if not resource:
            return None

        res = (
            instance_pb2.Tier2AlphaInstancePreprocessGetRecipeStepsKeyNotificationsUpdateKeyNotificationsInfoKeyNotificationConfigs()
        )
        if Primitive.to_proto(resource.key_or_version_name):
            res.key_or_version_name = Primitive.to_proto(resource.key_or_version_name)
        if Primitive.to_proto(resource.grant):
            res.grant = Primitive.to_proto(resource.grant)
        if Primitive.to_proto(resource.delegator_gaia_id):
            res.delegator_gaia_id = Primitive.to_proto(resource.delegator_gaia_id)
        return res

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return None

        return InstancePreprocessGetRecipeStepsKeyNotificationsUpdateKeyNotificationsInfoKeyNotificationConfigs(
            key_or_version_name=Primitive.from_proto(resource.key_or_version_name),
            grant=Primitive.from_proto(resource.grant),
            delegator_gaia_id=Primitive.from_proto(resource.delegator_gaia_id),
        )


class InstancePreprocessGetRecipeStepsKeyNotificationsUpdateKeyNotificationsInfoKeyNotificationConfigsArray(
    object
):
    @classmethod
    def to_proto(self, resources):
        if not resources:
            return resources
        return [
            InstancePreprocessGetRecipeStepsKeyNotificationsUpdateKeyNotificationsInfoKeyNotificationConfigs.to_proto(
                i
            )
            for i in resources
        ]

    @classmethod
    def from_proto(self, resources):
        return [
            InstancePreprocessGetRecipeStepsKeyNotificationsUpdateKeyNotificationsInfoKeyNotificationConfigs.from_proto(
                i
            )
            for i in resources
        ]


class InstanceNotifyKeyAvailableRecipe(object):
    def __init__(
        self,
        steps: list = None,
        honor_cancel_request: bool = None,
        ignore_recipe_after: int = None,
        verify_deadline_seconds_below: float = None,
        populate_operation_result: bool = None,
        readonly_recipe_start_time: str = None,
        resource_names_stored_in_clh_with_delay: list = None,
        delay_to_store_resources_in_clh_db_nanos: int = None,
    ):
        self.steps = steps
        self.honor_cancel_request = honor_cancel_request
        self.ignore_recipe_after = ignore_recipe_after
        self.verify_deadline_seconds_below = verify_deadline_seconds_below
        self.populate_operation_result = populate_operation_result
        self.readonly_recipe_start_time = readonly_recipe_start_time
        self.resource_names_stored_in_clh_with_delay = (
            resource_names_stored_in_clh_with_delay
        )
        self.delay_to_store_resources_in_clh_db_nanos = (
            delay_to_store_resources_in_clh_db_nanos
        )

    @classmethod
    def to_proto(self, resource):
        if not resource:
            return None

        res = instance_pb2.Tier2AlphaInstanceNotifyKeyAvailableRecipe()
        if InstanceNotifyKeyAvailableRecipeStepsArray.to_proto(resource.steps):
            res.steps.extend(
                InstanceNotifyKeyAvailableRecipeStepsArray.to_proto(resource.steps)
            )
        if Primitive.to_proto(resource.honor_cancel_request):
            res.honor_cancel_request = Primitive.to_proto(resource.honor_cancel_request)
        if Primitive.to_proto(resource.ignore_recipe_after):
            res.ignore_recipe_after = Primitive.to_proto(resource.ignore_recipe_after)
        if Primitive.to_proto(resource.verify_deadline_seconds_below):
            res.verify_deadline_seconds_below = Primitive.to_proto(
                resource.verify_deadline_seconds_below
            )
        if Primitive.to_proto(resource.populate_operation_result):
            res.populate_operation_result = Primitive.to_proto(
                resource.populate_operation_result
            )
        if Primitive.to_proto(resource.readonly_recipe_start_time):
            res.readonly_recipe_start_time = Primitive.to_proto(
                resource.readonly_recipe_start_time
            )
        if Primitive.to_proto(resource.resource_names_stored_in_clh_with_delay):
            res.resource_names_stored_in_clh_with_delay.extend(
                Primitive.to_proto(resource.resource_names_stored_in_clh_with_delay)
            )
        if Primitive.to_proto(resource.delay_to_store_resources_in_clh_db_nanos):
            res.delay_to_store_resources_in_clh_db_nanos = Primitive.to_proto(
                resource.delay_to_store_resources_in_clh_db_nanos
            )
        return res

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return None

        return InstanceNotifyKeyAvailableRecipe(
            steps=InstanceNotifyKeyAvailableRecipeStepsArray.from_proto(resource.steps),
            honor_cancel_request=Primitive.from_proto(resource.honor_cancel_request),
            ignore_recipe_after=Primitive.from_proto(resource.ignore_recipe_after),
            verify_deadline_seconds_below=Primitive.from_proto(
                resource.verify_deadline_seconds_below
            ),
            populate_operation_result=Primitive.from_proto(
                resource.populate_operation_result
            ),
            readonly_recipe_start_time=Primitive.from_proto(
                resource.readonly_recipe_start_time
            ),
            resource_names_stored_in_clh_with_delay=Primitive.from_proto(
                resource.resource_names_stored_in_clh_with_delay
            ),
            delay_to_store_resources_in_clh_db_nanos=Primitive.from_proto(
                resource.delay_to_store_resources_in_clh_db_nanos
            ),
        )


class InstanceNotifyKeyAvailableRecipeArray(object):
    @classmethod
    def to_proto(self, resources):
        if not resources:
            return resources
        return [InstanceNotifyKeyAvailableRecipe.to_proto(i) for i in resources]

    @classmethod
    def from_proto(self, resources):
        return [InstanceNotifyKeyAvailableRecipe.from_proto(i) for i in resources]


class InstanceNotifyKeyAvailableRecipeSteps(object):
    def __init__(
        self,
        relative_time: int = None,
        sleep_duration: int = None,
        action: str = None,
        status: dict = None,
        error_space: str = None,
        p4_service_account: str = None,
        resource_metadata_size: int = None,
        description: str = None,
        updated_repeat_operation_delay_sec: float = None,
        quota_request_deltas: list = None,
        preprocess_update: dict = None,
        public_operation_metadata: str = None,
        requested_tenant_project: dict = None,
        permissions_info: list = None,
        key_notifications_update: dict = None,
        clh_data_update_time: str = None,
        public_error_message: str = None,
    ):
        self.relative_time = relative_time
        self.sleep_duration = sleep_duration
        self.action = action
        self.status = status
        self.error_space = error_space
        self.p4_service_account = p4_service_account
        self.resource_metadata_size = resource_metadata_size
        self.description = description
        self.updated_repeat_operation_delay_sec = updated_repeat_operation_delay_sec
        self.quota_request_deltas = quota_request_deltas
        self.preprocess_update = preprocess_update
        self.public_operation_metadata = public_operation_metadata
        self.requested_tenant_project = requested_tenant_project
        self.permissions_info = permissions_info
        self.key_notifications_update = key_notifications_update
        self.clh_data_update_time = clh_data_update_time
        self.public_error_message = public_error_message

    @classmethod
    def to_proto(self, resource):
        if not resource:
            return None

        res = instance_pb2.Tier2AlphaInstanceNotifyKeyAvailableRecipeSteps()
        if Primitive.to_proto(resource.relative_time):
            res.relative_time = Primitive.to_proto(resource.relative_time)
        if Primitive.to_proto(resource.sleep_duration):
            res.sleep_duration = Primitive.to_proto(resource.sleep_duration)
        if InstanceNotifyKeyAvailableRecipeStepsActionEnum.to_proto(resource.action):
            res.action = InstanceNotifyKeyAvailableRecipeStepsActionEnum.to_proto(
                resource.action
            )
        if InstanceNotifyKeyAvailableRecipeStepsStatus.to_proto(resource.status):
            res.status.CopyFrom(
                InstanceNotifyKeyAvailableRecipeStepsStatus.to_proto(resource.status)
            )
        else:
            res.ClearField("status")
        if Primitive.to_proto(resource.error_space):
            res.error_space = Primitive.to_proto(resource.error_space)
        if Primitive.to_proto(resource.p4_service_account):
            res.p4_service_account = Primitive.to_proto(resource.p4_service_account)
        if Primitive.to_proto(resource.resource_metadata_size):
            res.resource_metadata_size = Primitive.to_proto(
                resource.resource_metadata_size
            )
        if Primitive.to_proto(resource.description):
            res.description = Primitive.to_proto(resource.description)
        if Primitive.to_proto(resource.updated_repeat_operation_delay_sec):
            res.updated_repeat_operation_delay_sec = Primitive.to_proto(
                resource.updated_repeat_operation_delay_sec
            )
        if InstanceNotifyKeyAvailableRecipeStepsQuotaRequestDeltasArray.to_proto(
            resource.quota_request_deltas
        ):
            res.quota_request_deltas.extend(
                InstanceNotifyKeyAvailableRecipeStepsQuotaRequestDeltasArray.to_proto(
                    resource.quota_request_deltas
                )
            )
        if InstanceNotifyKeyAvailableRecipeStepsPreprocessUpdate.to_proto(
            resource.preprocess_update
        ):
            res.preprocess_update.CopyFrom(
                InstanceNotifyKeyAvailableRecipeStepsPreprocessUpdate.to_proto(
                    resource.preprocess_update
                )
            )
        else:
            res.ClearField("preprocess_update")
        if Primitive.to_proto(resource.public_operation_metadata):
            res.public_operation_metadata = Primitive.to_proto(
                resource.public_operation_metadata
            )
        if InstanceNotifyKeyAvailableRecipeStepsRequestedTenantProject.to_proto(
            resource.requested_tenant_project
        ):
            res.requested_tenant_project.CopyFrom(
                InstanceNotifyKeyAvailableRecipeStepsRequestedTenantProject.to_proto(
                    resource.requested_tenant_project
                )
            )
        else:
            res.ClearField("requested_tenant_project")
        if InstanceNotifyKeyAvailableRecipeStepsPermissionsInfoArray.to_proto(
            resource.permissions_info
        ):
            res.permissions_info.extend(
                InstanceNotifyKeyAvailableRecipeStepsPermissionsInfoArray.to_proto(
                    resource.permissions_info
                )
            )
        if InstanceNotifyKeyAvailableRecipeStepsKeyNotificationsUpdate.to_proto(
            resource.key_notifications_update
        ):
            res.key_notifications_update.CopyFrom(
                InstanceNotifyKeyAvailableRecipeStepsKeyNotificationsUpdate.to_proto(
                    resource.key_notifications_update
                )
            )
        else:
            res.ClearField("key_notifications_update")
        if Primitive.to_proto(resource.clh_data_update_time):
            res.clh_data_update_time = Primitive.to_proto(resource.clh_data_update_time)
        if Primitive.to_proto(resource.public_error_message):
            res.public_error_message = Primitive.to_proto(resource.public_error_message)
        return res

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return None

        return InstanceNotifyKeyAvailableRecipeSteps(
            relative_time=Primitive.from_proto(resource.relative_time),
            sleep_duration=Primitive.from_proto(resource.sleep_duration),
            action=InstanceNotifyKeyAvailableRecipeStepsActionEnum.from_proto(
                resource.action
            ),
            status=InstanceNotifyKeyAvailableRecipeStepsStatus.from_proto(
                resource.status
            ),
            error_space=Primitive.from_proto(resource.error_space),
            p4_service_account=Primitive.from_proto(resource.p4_service_account),
            resource_metadata_size=Primitive.from_proto(
                resource.resource_metadata_size
            ),
            description=Primitive.from_proto(resource.description),
            updated_repeat_operation_delay_sec=Primitive.from_proto(
                resource.updated_repeat_operation_delay_sec
            ),
            quota_request_deltas=InstanceNotifyKeyAvailableRecipeStepsQuotaRequestDeltasArray.from_proto(
                resource.quota_request_deltas
            ),
            preprocess_update=InstanceNotifyKeyAvailableRecipeStepsPreprocessUpdate.from_proto(
                resource.preprocess_update
            ),
            public_operation_metadata=Primitive.from_proto(
                resource.public_operation_metadata
            ),
            requested_tenant_project=InstanceNotifyKeyAvailableRecipeStepsRequestedTenantProject.from_proto(
                resource.requested_tenant_project
            ),
            permissions_info=InstanceNotifyKeyAvailableRecipeStepsPermissionsInfoArray.from_proto(
                resource.permissions_info
            ),
            key_notifications_update=InstanceNotifyKeyAvailableRecipeStepsKeyNotificationsUpdate.from_proto(
                resource.key_notifications_update
            ),
            clh_data_update_time=Primitive.from_proto(resource.clh_data_update_time),
            public_error_message=Primitive.from_proto(resource.public_error_message),
        )


class InstanceNotifyKeyAvailableRecipeStepsArray(object):
    @classmethod
    def to_proto(self, resources):
        if not resources:
            return resources
        return [InstanceNotifyKeyAvailableRecipeSteps.to_proto(i) for i in resources]

    @classmethod
    def from_proto(self, resources):
        return [InstanceNotifyKeyAvailableRecipeSteps.from_proto(i) for i in resources]


class InstanceNotifyKeyAvailableRecipeStepsStatus(object):
    def __init__(self, code: int = None, message: str = None, details: list = None):
        self.code = code
        self.message = message
        self.details = details

    @classmethod
    def to_proto(self, resource):
        if not resource:
            return None

        res = instance_pb2.Tier2AlphaInstanceNotifyKeyAvailableRecipeStepsStatus()
        if Primitive.to_proto(resource.code):
            res.code = Primitive.to_proto(resource.code)
        if Primitive.to_proto(resource.message):
            res.message = Primitive.to_proto(resource.message)
        if InstanceNotifyKeyAvailableRecipeStepsStatusDetailsArray.to_proto(
            resource.details
        ):
            res.details.extend(
                InstanceNotifyKeyAvailableRecipeStepsStatusDetailsArray.to_proto(
                    resource.details
                )
            )
        return res

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return None

        return InstanceNotifyKeyAvailableRecipeStepsStatus(
            code=Primitive.from_proto(resource.code),
            message=Primitive.from_proto(resource.message),
            details=InstanceNotifyKeyAvailableRecipeStepsStatusDetailsArray.from_proto(
                resource.details
            ),
        )


class InstanceNotifyKeyAvailableRecipeStepsStatusArray(object):
    @classmethod
    def to_proto(self, resources):
        if not resources:
            return resources
        return [
            InstanceNotifyKeyAvailableRecipeStepsStatus.to_proto(i) for i in resources
        ]

    @classmethod
    def from_proto(self, resources):
        return [
            InstanceNotifyKeyAvailableRecipeStepsStatus.from_proto(i) for i in resources
        ]


class InstanceNotifyKeyAvailableRecipeStepsStatusDetails(object):
    def __init__(self, type_url: str = None, value: str = None):
        self.type_url = type_url
        self.value = value

    @classmethod
    def to_proto(self, resource):
        if not resource:
            return None

        res = (
            instance_pb2.Tier2AlphaInstanceNotifyKeyAvailableRecipeStepsStatusDetails()
        )
        if Primitive.to_proto(resource.type_url):
            res.type_url = Primitive.to_proto(resource.type_url)
        if Primitive.to_proto(resource.value):
            res.value = Primitive.to_proto(resource.value)
        return res

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return None

        return InstanceNotifyKeyAvailableRecipeStepsStatusDetails(
            type_url=Primitive.from_proto(resource.type_url),
            value=Primitive.from_proto(resource.value),
        )


class InstanceNotifyKeyAvailableRecipeStepsStatusDetailsArray(object):
    @classmethod
    def to_proto(self, resources):
        if not resources:
            return resources
        return [
            InstanceNotifyKeyAvailableRecipeStepsStatusDetails.to_proto(i)
            for i in resources
        ]

    @classmethod
    def from_proto(self, resources):
        return [
            InstanceNotifyKeyAvailableRecipeStepsStatusDetails.from_proto(i)
            for i in resources
        ]


class InstanceNotifyKeyAvailableRecipeStepsQuotaRequestDeltas(object):
    def __init__(
        self,
        metric_name: str = None,
        amount: int = None,
        quota_location_name: str = None,
    ):
        self.metric_name = metric_name
        self.amount = amount
        self.quota_location_name = quota_location_name

    @classmethod
    def to_proto(self, resource):
        if not resource:
            return None

        res = (
            instance_pb2.Tier2AlphaInstanceNotifyKeyAvailableRecipeStepsQuotaRequestDeltas()
        )
        if Primitive.to_proto(resource.metric_name):
            res.metric_name = Primitive.to_proto(resource.metric_name)
        if Primitive.to_proto(resource.amount):
            res.amount = Primitive.to_proto(resource.amount)
        if Primitive.to_proto(resource.quota_location_name):
            res.quota_location_name = Primitive.to_proto(resource.quota_location_name)
        return res

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return None

        return InstanceNotifyKeyAvailableRecipeStepsQuotaRequestDeltas(
            metric_name=Primitive.from_proto(resource.metric_name),
            amount=Primitive.from_proto(resource.amount),
            quota_location_name=Primitive.from_proto(resource.quota_location_name),
        )


class InstanceNotifyKeyAvailableRecipeStepsQuotaRequestDeltasArray(object):
    @classmethod
    def to_proto(self, resources):
        if not resources:
            return resources
        return [
            InstanceNotifyKeyAvailableRecipeStepsQuotaRequestDeltas.to_proto(i)
            for i in resources
        ]

    @classmethod
    def from_proto(self, resources):
        return [
            InstanceNotifyKeyAvailableRecipeStepsQuotaRequestDeltas.from_proto(i)
            for i in resources
        ]


class InstanceNotifyKeyAvailableRecipeStepsPreprocessUpdate(object):
    def __init__(
        self, latency_slo_bucket_name: str = None, public_operation_metadata: str = None
    ):
        self.latency_slo_bucket_name = latency_slo_bucket_name
        self.public_operation_metadata = public_operation_metadata

    @classmethod
    def to_proto(self, resource):
        if not resource:
            return None

        res = (
            instance_pb2.Tier2AlphaInstanceNotifyKeyAvailableRecipeStepsPreprocessUpdate()
        )
        if Primitive.to_proto(resource.latency_slo_bucket_name):
            res.latency_slo_bucket_name = Primitive.to_proto(
                resource.latency_slo_bucket_name
            )
        if Primitive.to_proto(resource.public_operation_metadata):
            res.public_operation_metadata = Primitive.to_proto(
                resource.public_operation_metadata
            )
        return res

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return None

        return InstanceNotifyKeyAvailableRecipeStepsPreprocessUpdate(
            latency_slo_bucket_name=Primitive.from_proto(
                resource.latency_slo_bucket_name
            ),
            public_operation_metadata=Primitive.from_proto(
                resource.public_operation_metadata
            ),
        )


class InstanceNotifyKeyAvailableRecipeStepsPreprocessUpdateArray(object):
    @classmethod
    def to_proto(self, resources):
        if not resources:
            return resources
        return [
            InstanceNotifyKeyAvailableRecipeStepsPreprocessUpdate.to_proto(i)
            for i in resources
        ]

    @classmethod
    def from_proto(self, resources):
        return [
            InstanceNotifyKeyAvailableRecipeStepsPreprocessUpdate.from_proto(i)
            for i in resources
        ]


class InstanceNotifyKeyAvailableRecipeStepsRequestedTenantProject(object):
    def __init__(self, tag: str = None, folder: str = None, scope: str = None):
        self.tag = tag
        self.folder = folder
        self.scope = scope

    @classmethod
    def to_proto(self, resource):
        if not resource:
            return None

        res = (
            instance_pb2.Tier2AlphaInstanceNotifyKeyAvailableRecipeStepsRequestedTenantProject()
        )
        if Primitive.to_proto(resource.tag):
            res.tag = Primitive.to_proto(resource.tag)
        if Primitive.to_proto(resource.folder):
            res.folder = Primitive.to_proto(resource.folder)
        if InstanceNotifyKeyAvailableRecipeStepsRequestedTenantProjectScopeEnum.to_proto(
            resource.scope
        ):
            res.scope = InstanceNotifyKeyAvailableRecipeStepsRequestedTenantProjectScopeEnum.to_proto(
                resource.scope
            )
        return res

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return None

        return InstanceNotifyKeyAvailableRecipeStepsRequestedTenantProject(
            tag=Primitive.from_proto(resource.tag),
            folder=Primitive.from_proto(resource.folder),
            scope=InstanceNotifyKeyAvailableRecipeStepsRequestedTenantProjectScopeEnum.from_proto(
                resource.scope
            ),
        )


class InstanceNotifyKeyAvailableRecipeStepsRequestedTenantProjectArray(object):
    @classmethod
    def to_proto(self, resources):
        if not resources:
            return resources
        return [
            InstanceNotifyKeyAvailableRecipeStepsRequestedTenantProject.to_proto(i)
            for i in resources
        ]

    @classmethod
    def from_proto(self, resources):
        return [
            InstanceNotifyKeyAvailableRecipeStepsRequestedTenantProject.from_proto(i)
            for i in resources
        ]


class InstanceNotifyKeyAvailableRecipeStepsPermissionsInfo(object):
    def __init__(
        self,
        policy_name: dict = None,
        iam_permissions: list = None,
        resource_path: str = None,
        api_attrs: dict = None,
        policy_name_mode: str = None,
        resource: dict = None,
    ):
        self.policy_name = policy_name
        self.iam_permissions = iam_permissions
        self.resource_path = resource_path
        self.api_attrs = api_attrs
        self.policy_name_mode = policy_name_mode
        self.resource = resource

    @classmethod
    def to_proto(self, resource):
        if not resource:
            return None

        res = (
            instance_pb2.Tier2AlphaInstanceNotifyKeyAvailableRecipeStepsPermissionsInfo()
        )
        if InstanceNotifyKeyAvailableRecipeStepsPermissionsInfoPolicyName.to_proto(
            resource.policy_name
        ):
            res.policy_name.CopyFrom(
                InstanceNotifyKeyAvailableRecipeStepsPermissionsInfoPolicyName.to_proto(
                    resource.policy_name
                )
            )
        else:
            res.ClearField("policy_name")
        if InstanceNotifyKeyAvailableRecipeStepsPermissionsInfoIamPermissionsArray.to_proto(
            resource.iam_permissions
        ):
            res.iam_permissions.extend(
                InstanceNotifyKeyAvailableRecipeStepsPermissionsInfoIamPermissionsArray.to_proto(
                    resource.iam_permissions
                )
            )
        if Primitive.to_proto(resource.resource_path):
            res.resource_path = Primitive.to_proto(resource.resource_path)
        if InstanceGoogleprotobufstruct.to_proto(resource.api_attrs):
            res.api_attrs.CopyFrom(
                InstanceGoogleprotobufstruct.to_proto(resource.api_attrs)
            )
        else:
            res.ClearField("api_attrs")
        if InstanceNotifyKeyAvailableRecipeStepsPermissionsInfoPolicyNameModeEnum.to_proto(
            resource.policy_name_mode
        ):
            res.policy_name_mode = InstanceNotifyKeyAvailableRecipeStepsPermissionsInfoPolicyNameModeEnum.to_proto(
                resource.policy_name_mode
            )
        if InstanceNotifyKeyAvailableRecipeStepsPermissionsInfoResource.to_proto(
            resource.resource
        ):
            res.resource.CopyFrom(
                InstanceNotifyKeyAvailableRecipeStepsPermissionsInfoResource.to_proto(
                    resource.resource
                )
            )
        else:
            res.ClearField("resource")
        return res

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return None

        return InstanceNotifyKeyAvailableRecipeStepsPermissionsInfo(
            policy_name=InstanceNotifyKeyAvailableRecipeStepsPermissionsInfoPolicyName.from_proto(
                resource.policy_name
            ),
            iam_permissions=InstanceNotifyKeyAvailableRecipeStepsPermissionsInfoIamPermissionsArray.from_proto(
                resource.iam_permissions
            ),
            resource_path=Primitive.from_proto(resource.resource_path),
            api_attrs=InstanceGoogleprotobufstruct.from_proto(resource.api_attrs),
            policy_name_mode=InstanceNotifyKeyAvailableRecipeStepsPermissionsInfoPolicyNameModeEnum.from_proto(
                resource.policy_name_mode
            ),
            resource=InstanceNotifyKeyAvailableRecipeStepsPermissionsInfoResource.from_proto(
                resource.resource
            ),
        )


class InstanceNotifyKeyAvailableRecipeStepsPermissionsInfoArray(object):
    @classmethod
    def to_proto(self, resources):
        if not resources:
            return resources
        return [
            InstanceNotifyKeyAvailableRecipeStepsPermissionsInfo.to_proto(i)
            for i in resources
        ]

    @classmethod
    def from_proto(self, resources):
        return [
            InstanceNotifyKeyAvailableRecipeStepsPermissionsInfo.from_proto(i)
            for i in resources
        ]


class InstanceNotifyKeyAvailableRecipeStepsPermissionsInfoPolicyName(object):
    def __init__(self, type: str = None, id: str = None, region: str = None):
        self.type = type
        self.id = id
        self.region = region

    @classmethod
    def to_proto(self, resource):
        if not resource:
            return None

        res = (
            instance_pb2.Tier2AlphaInstanceNotifyKeyAvailableRecipeStepsPermissionsInfoPolicyName()
        )
        if Primitive.to_proto(resource.type):
            res.type = Primitive.to_proto(resource.type)
        if Primitive.to_proto(resource.id):
            res.id = Primitive.to_proto(resource.id)
        if Primitive.to_proto(resource.region):
            res.region = Primitive.to_proto(resource.region)
        return res

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return None

        return InstanceNotifyKeyAvailableRecipeStepsPermissionsInfoPolicyName(
            type=Primitive.from_proto(resource.type),
            id=Primitive.from_proto(resource.id),
            region=Primitive.from_proto(resource.region),
        )


class InstanceNotifyKeyAvailableRecipeStepsPermissionsInfoPolicyNameArray(object):
    @classmethod
    def to_proto(self, resources):
        if not resources:
            return resources
        return [
            InstanceNotifyKeyAvailableRecipeStepsPermissionsInfoPolicyName.to_proto(i)
            for i in resources
        ]

    @classmethod
    def from_proto(self, resources):
        return [
            InstanceNotifyKeyAvailableRecipeStepsPermissionsInfoPolicyName.from_proto(i)
            for i in resources
        ]


class InstanceNotifyKeyAvailableRecipeStepsPermissionsInfoIamPermissions(object):
    def __init__(self, permission: str = None):
        self.permission = permission

    @classmethod
    def to_proto(self, resource):
        if not resource:
            return None

        res = (
            instance_pb2.Tier2AlphaInstanceNotifyKeyAvailableRecipeStepsPermissionsInfoIamPermissions()
        )
        if Primitive.to_proto(resource.permission):
            res.permission = Primitive.to_proto(resource.permission)
        return res

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return None

        return InstanceNotifyKeyAvailableRecipeStepsPermissionsInfoIamPermissions(
            permission=Primitive.from_proto(resource.permission),
        )


class InstanceNotifyKeyAvailableRecipeStepsPermissionsInfoIamPermissionsArray(object):
    @classmethod
    def to_proto(self, resources):
        if not resources:
            return resources
        return [
            InstanceNotifyKeyAvailableRecipeStepsPermissionsInfoIamPermissions.to_proto(
                i
            )
            for i in resources
        ]

    @classmethod
    def from_proto(self, resources):
        return [
            InstanceNotifyKeyAvailableRecipeStepsPermissionsInfoIamPermissions.from_proto(
                i
            )
            for i in resources
        ]


class InstanceNotifyKeyAvailableRecipeStepsPermissionsInfoResource(object):
    def __init__(
        self,
        name: str = None,
        type: str = None,
        service: str = None,
        labels: dict = None,
    ):
        self.name = name
        self.type = type
        self.service = service
        self.labels = labels

    @classmethod
    def to_proto(self, resource):
        if not resource:
            return None

        res = (
            instance_pb2.Tier2AlphaInstanceNotifyKeyAvailableRecipeStepsPermissionsInfoResource()
        )
        if Primitive.to_proto(resource.name):
            res.name = Primitive.to_proto(resource.name)
        if Primitive.to_proto(resource.type):
            res.type = Primitive.to_proto(resource.type)
        if Primitive.to_proto(resource.service):
            res.service = Primitive.to_proto(resource.service)
        if Primitive.to_proto(resource.labels):
            res.labels = Primitive.to_proto(resource.labels)
        return res

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return None

        return InstanceNotifyKeyAvailableRecipeStepsPermissionsInfoResource(
            name=Primitive.from_proto(resource.name),
            type=Primitive.from_proto(resource.type),
            service=Primitive.from_proto(resource.service),
            labels=Primitive.from_proto(resource.labels),
        )


class InstanceNotifyKeyAvailableRecipeStepsPermissionsInfoResourceArray(object):
    @classmethod
    def to_proto(self, resources):
        if not resources:
            return resources
        return [
            InstanceNotifyKeyAvailableRecipeStepsPermissionsInfoResource.to_proto(i)
            for i in resources
        ]

    @classmethod
    def from_proto(self, resources):
        return [
            InstanceNotifyKeyAvailableRecipeStepsPermissionsInfoResource.from_proto(i)
            for i in resources
        ]


class InstanceNotifyKeyAvailableRecipeStepsKeyNotificationsUpdate(object):
    def __init__(self, key_notifications_info: dict = None):
        self.key_notifications_info = key_notifications_info

    @classmethod
    def to_proto(self, resource):
        if not resource:
            return None

        res = (
            instance_pb2.Tier2AlphaInstanceNotifyKeyAvailableRecipeStepsKeyNotificationsUpdate()
        )
        if InstanceNotifyKeyAvailableRecipeStepsKeyNotificationsUpdateKeyNotificationsInfo.to_proto(
            resource.key_notifications_info
        ):
            res.key_notifications_info.CopyFrom(
                InstanceNotifyKeyAvailableRecipeStepsKeyNotificationsUpdateKeyNotificationsInfo.to_proto(
                    resource.key_notifications_info
                )
            )
        else:
            res.ClearField("key_notifications_info")
        return res

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return None

        return InstanceNotifyKeyAvailableRecipeStepsKeyNotificationsUpdate(
            key_notifications_info=InstanceNotifyKeyAvailableRecipeStepsKeyNotificationsUpdateKeyNotificationsInfo.from_proto(
                resource.key_notifications_info
            ),
        )


class InstanceNotifyKeyAvailableRecipeStepsKeyNotificationsUpdateArray(object):
    @classmethod
    def to_proto(self, resources):
        if not resources:
            return resources
        return [
            InstanceNotifyKeyAvailableRecipeStepsKeyNotificationsUpdate.to_proto(i)
            for i in resources
        ]

    @classmethod
    def from_proto(self, resources):
        return [
            InstanceNotifyKeyAvailableRecipeStepsKeyNotificationsUpdate.from_proto(i)
            for i in resources
        ]


class InstanceNotifyKeyAvailableRecipeStepsKeyNotificationsUpdateKeyNotificationsInfo(
    object
):
    def __init__(
        self,
        data_version: int = None,
        delegate: str = None,
        key_notification_configs: list = None,
    ):
        self.data_version = data_version
        self.delegate = delegate
        self.key_notification_configs = key_notification_configs

    @classmethod
    def to_proto(self, resource):
        if not resource:
            return None

        res = (
            instance_pb2.Tier2AlphaInstanceNotifyKeyAvailableRecipeStepsKeyNotificationsUpdateKeyNotificationsInfo()
        )
        if Primitive.to_proto(resource.data_version):
            res.data_version = Primitive.to_proto(resource.data_version)
        if Primitive.to_proto(resource.delegate):
            res.delegate = Primitive.to_proto(resource.delegate)
        if InstanceNotifyKeyAvailableRecipeStepsKeyNotificationsUpdateKeyNotificationsInfoKeyNotificationConfigsArray.to_proto(
            resource.key_notification_configs
        ):
            res.key_notification_configs.extend(
                InstanceNotifyKeyAvailableRecipeStepsKeyNotificationsUpdateKeyNotificationsInfoKeyNotificationConfigsArray.to_proto(
                    resource.key_notification_configs
                )
            )
        return res

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return None

        return InstanceNotifyKeyAvailableRecipeStepsKeyNotificationsUpdateKeyNotificationsInfo(
            data_version=Primitive.from_proto(resource.data_version),
            delegate=Primitive.from_proto(resource.delegate),
            key_notification_configs=InstanceNotifyKeyAvailableRecipeStepsKeyNotificationsUpdateKeyNotificationsInfoKeyNotificationConfigsArray.from_proto(
                resource.key_notification_configs
            ),
        )


class InstanceNotifyKeyAvailableRecipeStepsKeyNotificationsUpdateKeyNotificationsInfoArray(
    object
):
    @classmethod
    def to_proto(self, resources):
        if not resources:
            return resources
        return [
            InstanceNotifyKeyAvailableRecipeStepsKeyNotificationsUpdateKeyNotificationsInfo.to_proto(
                i
            )
            for i in resources
        ]

    @classmethod
    def from_proto(self, resources):
        return [
            InstanceNotifyKeyAvailableRecipeStepsKeyNotificationsUpdateKeyNotificationsInfo.from_proto(
                i
            )
            for i in resources
        ]


class InstanceNotifyKeyAvailableRecipeStepsKeyNotificationsUpdateKeyNotificationsInfoKeyNotificationConfigs(
    object
):
    def __init__(
        self,
        key_or_version_name: str = None,
        grant: str = None,
        delegator_gaia_id: int = None,
    ):
        self.key_or_version_name = key_or_version_name
        self.grant = grant
        self.delegator_gaia_id = delegator_gaia_id

    @classmethod
    def to_proto(self, resource):
        if not resource:
            return None

        res = (
            instance_pb2.Tier2AlphaInstanceNotifyKeyAvailableRecipeStepsKeyNotificationsUpdateKeyNotificationsInfoKeyNotificationConfigs()
        )
        if Primitive.to_proto(resource.key_or_version_name):
            res.key_or_version_name = Primitive.to_proto(resource.key_or_version_name)
        if Primitive.to_proto(resource.grant):
            res.grant = Primitive.to_proto(resource.grant)
        if Primitive.to_proto(resource.delegator_gaia_id):
            res.delegator_gaia_id = Primitive.to_proto(resource.delegator_gaia_id)
        return res

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return None

        return InstanceNotifyKeyAvailableRecipeStepsKeyNotificationsUpdateKeyNotificationsInfoKeyNotificationConfigs(
            key_or_version_name=Primitive.from_proto(resource.key_or_version_name),
            grant=Primitive.from_proto(resource.grant),
            delegator_gaia_id=Primitive.from_proto(resource.delegator_gaia_id),
        )


class InstanceNotifyKeyAvailableRecipeStepsKeyNotificationsUpdateKeyNotificationsInfoKeyNotificationConfigsArray(
    object
):
    @classmethod
    def to_proto(self, resources):
        if not resources:
            return resources
        return [
            InstanceNotifyKeyAvailableRecipeStepsKeyNotificationsUpdateKeyNotificationsInfoKeyNotificationConfigs.to_proto(
                i
            )
            for i in resources
        ]

    @classmethod
    def from_proto(self, resources):
        return [
            InstanceNotifyKeyAvailableRecipeStepsKeyNotificationsUpdateKeyNotificationsInfoKeyNotificationConfigs.from_proto(
                i
            )
            for i in resources
        ]


class InstanceNotifyKeyUnavailableRecipe(object):
    def __init__(
        self,
        steps: list = None,
        honor_cancel_request: bool = None,
        ignore_recipe_after: int = None,
        verify_deadline_seconds_below: float = None,
        populate_operation_result: bool = None,
        readonly_recipe_start_time: str = None,
        resource_names_stored_in_clh_with_delay: list = None,
        delay_to_store_resources_in_clh_db_nanos: int = None,
    ):
        self.steps = steps
        self.honor_cancel_request = honor_cancel_request
        self.ignore_recipe_after = ignore_recipe_after
        self.verify_deadline_seconds_below = verify_deadline_seconds_below
        self.populate_operation_result = populate_operation_result
        self.readonly_recipe_start_time = readonly_recipe_start_time
        self.resource_names_stored_in_clh_with_delay = (
            resource_names_stored_in_clh_with_delay
        )
        self.delay_to_store_resources_in_clh_db_nanos = (
            delay_to_store_resources_in_clh_db_nanos
        )

    @classmethod
    def to_proto(self, resource):
        if not resource:
            return None

        res = instance_pb2.Tier2AlphaInstanceNotifyKeyUnavailableRecipe()
        if InstanceNotifyKeyUnavailableRecipeStepsArray.to_proto(resource.steps):
            res.steps.extend(
                InstanceNotifyKeyUnavailableRecipeStepsArray.to_proto(resource.steps)
            )
        if Primitive.to_proto(resource.honor_cancel_request):
            res.honor_cancel_request = Primitive.to_proto(resource.honor_cancel_request)
        if Primitive.to_proto(resource.ignore_recipe_after):
            res.ignore_recipe_after = Primitive.to_proto(resource.ignore_recipe_after)
        if Primitive.to_proto(resource.verify_deadline_seconds_below):
            res.verify_deadline_seconds_below = Primitive.to_proto(
                resource.verify_deadline_seconds_below
            )
        if Primitive.to_proto(resource.populate_operation_result):
            res.populate_operation_result = Primitive.to_proto(
                resource.populate_operation_result
            )
        if Primitive.to_proto(resource.readonly_recipe_start_time):
            res.readonly_recipe_start_time = Primitive.to_proto(
                resource.readonly_recipe_start_time
            )
        if Primitive.to_proto(resource.resource_names_stored_in_clh_with_delay):
            res.resource_names_stored_in_clh_with_delay.extend(
                Primitive.to_proto(resource.resource_names_stored_in_clh_with_delay)
            )
        if Primitive.to_proto(resource.delay_to_store_resources_in_clh_db_nanos):
            res.delay_to_store_resources_in_clh_db_nanos = Primitive.to_proto(
                resource.delay_to_store_resources_in_clh_db_nanos
            )
        return res

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return None

        return InstanceNotifyKeyUnavailableRecipe(
            steps=InstanceNotifyKeyUnavailableRecipeStepsArray.from_proto(
                resource.steps
            ),
            honor_cancel_request=Primitive.from_proto(resource.honor_cancel_request),
            ignore_recipe_after=Primitive.from_proto(resource.ignore_recipe_after),
            verify_deadline_seconds_below=Primitive.from_proto(
                resource.verify_deadline_seconds_below
            ),
            populate_operation_result=Primitive.from_proto(
                resource.populate_operation_result
            ),
            readonly_recipe_start_time=Primitive.from_proto(
                resource.readonly_recipe_start_time
            ),
            resource_names_stored_in_clh_with_delay=Primitive.from_proto(
                resource.resource_names_stored_in_clh_with_delay
            ),
            delay_to_store_resources_in_clh_db_nanos=Primitive.from_proto(
                resource.delay_to_store_resources_in_clh_db_nanos
            ),
        )


class InstanceNotifyKeyUnavailableRecipeArray(object):
    @classmethod
    def to_proto(self, resources):
        if not resources:
            return resources
        return [InstanceNotifyKeyUnavailableRecipe.to_proto(i) for i in resources]

    @classmethod
    def from_proto(self, resources):
        return [InstanceNotifyKeyUnavailableRecipe.from_proto(i) for i in resources]


class InstanceNotifyKeyUnavailableRecipeSteps(object):
    def __init__(
        self,
        relative_time: int = None,
        sleep_duration: int = None,
        action: str = None,
        status: dict = None,
        error_space: str = None,
        p4_service_account: str = None,
        resource_metadata_size: int = None,
        description: str = None,
        updated_repeat_operation_delay_sec: float = None,
        quota_request_deltas: list = None,
        preprocess_update: dict = None,
        public_operation_metadata: str = None,
        requested_tenant_project: dict = None,
        permissions_info: list = None,
        key_notifications_update: dict = None,
        clh_data_update_time: str = None,
        public_error_message: str = None,
    ):
        self.relative_time = relative_time
        self.sleep_duration = sleep_duration
        self.action = action
        self.status = status
        self.error_space = error_space
        self.p4_service_account = p4_service_account
        self.resource_metadata_size = resource_metadata_size
        self.description = description
        self.updated_repeat_operation_delay_sec = updated_repeat_operation_delay_sec
        self.quota_request_deltas = quota_request_deltas
        self.preprocess_update = preprocess_update
        self.public_operation_metadata = public_operation_metadata
        self.requested_tenant_project = requested_tenant_project
        self.permissions_info = permissions_info
        self.key_notifications_update = key_notifications_update
        self.clh_data_update_time = clh_data_update_time
        self.public_error_message = public_error_message

    @classmethod
    def to_proto(self, resource):
        if not resource:
            return None

        res = instance_pb2.Tier2AlphaInstanceNotifyKeyUnavailableRecipeSteps()
        if Primitive.to_proto(resource.relative_time):
            res.relative_time = Primitive.to_proto(resource.relative_time)
        if Primitive.to_proto(resource.sleep_duration):
            res.sleep_duration = Primitive.to_proto(resource.sleep_duration)
        if InstanceNotifyKeyUnavailableRecipeStepsActionEnum.to_proto(resource.action):
            res.action = InstanceNotifyKeyUnavailableRecipeStepsActionEnum.to_proto(
                resource.action
            )
        if InstanceNotifyKeyUnavailableRecipeStepsStatus.to_proto(resource.status):
            res.status.CopyFrom(
                InstanceNotifyKeyUnavailableRecipeStepsStatus.to_proto(resource.status)
            )
        else:
            res.ClearField("status")
        if Primitive.to_proto(resource.error_space):
            res.error_space = Primitive.to_proto(resource.error_space)
        if Primitive.to_proto(resource.p4_service_account):
            res.p4_service_account = Primitive.to_proto(resource.p4_service_account)
        if Primitive.to_proto(resource.resource_metadata_size):
            res.resource_metadata_size = Primitive.to_proto(
                resource.resource_metadata_size
            )
        if Primitive.to_proto(resource.description):
            res.description = Primitive.to_proto(resource.description)
        if Primitive.to_proto(resource.updated_repeat_operation_delay_sec):
            res.updated_repeat_operation_delay_sec = Primitive.to_proto(
                resource.updated_repeat_operation_delay_sec
            )
        if InstanceNotifyKeyUnavailableRecipeStepsQuotaRequestDeltasArray.to_proto(
            resource.quota_request_deltas
        ):
            res.quota_request_deltas.extend(
                InstanceNotifyKeyUnavailableRecipeStepsQuotaRequestDeltasArray.to_proto(
                    resource.quota_request_deltas
                )
            )
        if InstanceNotifyKeyUnavailableRecipeStepsPreprocessUpdate.to_proto(
            resource.preprocess_update
        ):
            res.preprocess_update.CopyFrom(
                InstanceNotifyKeyUnavailableRecipeStepsPreprocessUpdate.to_proto(
                    resource.preprocess_update
                )
            )
        else:
            res.ClearField("preprocess_update")
        if Primitive.to_proto(resource.public_operation_metadata):
            res.public_operation_metadata = Primitive.to_proto(
                resource.public_operation_metadata
            )
        if InstanceNotifyKeyUnavailableRecipeStepsRequestedTenantProject.to_proto(
            resource.requested_tenant_project
        ):
            res.requested_tenant_project.CopyFrom(
                InstanceNotifyKeyUnavailableRecipeStepsRequestedTenantProject.to_proto(
                    resource.requested_tenant_project
                )
            )
        else:
            res.ClearField("requested_tenant_project")
        if InstanceNotifyKeyUnavailableRecipeStepsPermissionsInfoArray.to_proto(
            resource.permissions_info
        ):
            res.permissions_info.extend(
                InstanceNotifyKeyUnavailableRecipeStepsPermissionsInfoArray.to_proto(
                    resource.permissions_info
                )
            )
        if InstanceNotifyKeyUnavailableRecipeStepsKeyNotificationsUpdate.to_proto(
            resource.key_notifications_update
        ):
            res.key_notifications_update.CopyFrom(
                InstanceNotifyKeyUnavailableRecipeStepsKeyNotificationsUpdate.to_proto(
                    resource.key_notifications_update
                )
            )
        else:
            res.ClearField("key_notifications_update")
        if Primitive.to_proto(resource.clh_data_update_time):
            res.clh_data_update_time = Primitive.to_proto(resource.clh_data_update_time)
        if Primitive.to_proto(resource.public_error_message):
            res.public_error_message = Primitive.to_proto(resource.public_error_message)
        return res

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return None

        return InstanceNotifyKeyUnavailableRecipeSteps(
            relative_time=Primitive.from_proto(resource.relative_time),
            sleep_duration=Primitive.from_proto(resource.sleep_duration),
            action=InstanceNotifyKeyUnavailableRecipeStepsActionEnum.from_proto(
                resource.action
            ),
            status=InstanceNotifyKeyUnavailableRecipeStepsStatus.from_proto(
                resource.status
            ),
            error_space=Primitive.from_proto(resource.error_space),
            p4_service_account=Primitive.from_proto(resource.p4_service_account),
            resource_metadata_size=Primitive.from_proto(
                resource.resource_metadata_size
            ),
            description=Primitive.from_proto(resource.description),
            updated_repeat_operation_delay_sec=Primitive.from_proto(
                resource.updated_repeat_operation_delay_sec
            ),
            quota_request_deltas=InstanceNotifyKeyUnavailableRecipeStepsQuotaRequestDeltasArray.from_proto(
                resource.quota_request_deltas
            ),
            preprocess_update=InstanceNotifyKeyUnavailableRecipeStepsPreprocessUpdate.from_proto(
                resource.preprocess_update
            ),
            public_operation_metadata=Primitive.from_proto(
                resource.public_operation_metadata
            ),
            requested_tenant_project=InstanceNotifyKeyUnavailableRecipeStepsRequestedTenantProject.from_proto(
                resource.requested_tenant_project
            ),
            permissions_info=InstanceNotifyKeyUnavailableRecipeStepsPermissionsInfoArray.from_proto(
                resource.permissions_info
            ),
            key_notifications_update=InstanceNotifyKeyUnavailableRecipeStepsKeyNotificationsUpdate.from_proto(
                resource.key_notifications_update
            ),
            clh_data_update_time=Primitive.from_proto(resource.clh_data_update_time),
            public_error_message=Primitive.from_proto(resource.public_error_message),
        )


class InstanceNotifyKeyUnavailableRecipeStepsArray(object):
    @classmethod
    def to_proto(self, resources):
        if not resources:
            return resources
        return [InstanceNotifyKeyUnavailableRecipeSteps.to_proto(i) for i in resources]

    @classmethod
    def from_proto(self, resources):
        return [
            InstanceNotifyKeyUnavailableRecipeSteps.from_proto(i) for i in resources
        ]


class InstanceNotifyKeyUnavailableRecipeStepsStatus(object):
    def __init__(self, code: int = None, message: str = None, details: list = None):
        self.code = code
        self.message = message
        self.details = details

    @classmethod
    def to_proto(self, resource):
        if not resource:
            return None

        res = instance_pb2.Tier2AlphaInstanceNotifyKeyUnavailableRecipeStepsStatus()
        if Primitive.to_proto(resource.code):
            res.code = Primitive.to_proto(resource.code)
        if Primitive.to_proto(resource.message):
            res.message = Primitive.to_proto(resource.message)
        if InstanceNotifyKeyUnavailableRecipeStepsStatusDetailsArray.to_proto(
            resource.details
        ):
            res.details.extend(
                InstanceNotifyKeyUnavailableRecipeStepsStatusDetailsArray.to_proto(
                    resource.details
                )
            )
        return res

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return None

        return InstanceNotifyKeyUnavailableRecipeStepsStatus(
            code=Primitive.from_proto(resource.code),
            message=Primitive.from_proto(resource.message),
            details=InstanceNotifyKeyUnavailableRecipeStepsStatusDetailsArray.from_proto(
                resource.details
            ),
        )


class InstanceNotifyKeyUnavailableRecipeStepsStatusArray(object):
    @classmethod
    def to_proto(self, resources):
        if not resources:
            return resources
        return [
            InstanceNotifyKeyUnavailableRecipeStepsStatus.to_proto(i) for i in resources
        ]

    @classmethod
    def from_proto(self, resources):
        return [
            InstanceNotifyKeyUnavailableRecipeStepsStatus.from_proto(i)
            for i in resources
        ]


class InstanceNotifyKeyUnavailableRecipeStepsStatusDetails(object):
    def __init__(self, type_url: str = None, value: str = None):
        self.type_url = type_url
        self.value = value

    @classmethod
    def to_proto(self, resource):
        if not resource:
            return None

        res = (
            instance_pb2.Tier2AlphaInstanceNotifyKeyUnavailableRecipeStepsStatusDetails()
        )
        if Primitive.to_proto(resource.type_url):
            res.type_url = Primitive.to_proto(resource.type_url)
        if Primitive.to_proto(resource.value):
            res.value = Primitive.to_proto(resource.value)
        return res

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return None

        return InstanceNotifyKeyUnavailableRecipeStepsStatusDetails(
            type_url=Primitive.from_proto(resource.type_url),
            value=Primitive.from_proto(resource.value),
        )


class InstanceNotifyKeyUnavailableRecipeStepsStatusDetailsArray(object):
    @classmethod
    def to_proto(self, resources):
        if not resources:
            return resources
        return [
            InstanceNotifyKeyUnavailableRecipeStepsStatusDetails.to_proto(i)
            for i in resources
        ]

    @classmethod
    def from_proto(self, resources):
        return [
            InstanceNotifyKeyUnavailableRecipeStepsStatusDetails.from_proto(i)
            for i in resources
        ]


class InstanceNotifyKeyUnavailableRecipeStepsQuotaRequestDeltas(object):
    def __init__(
        self,
        metric_name: str = None,
        amount: int = None,
        quota_location_name: str = None,
    ):
        self.metric_name = metric_name
        self.amount = amount
        self.quota_location_name = quota_location_name

    @classmethod
    def to_proto(self, resource):
        if not resource:
            return None

        res = (
            instance_pb2.Tier2AlphaInstanceNotifyKeyUnavailableRecipeStepsQuotaRequestDeltas()
        )
        if Primitive.to_proto(resource.metric_name):
            res.metric_name = Primitive.to_proto(resource.metric_name)
        if Primitive.to_proto(resource.amount):
            res.amount = Primitive.to_proto(resource.amount)
        if Primitive.to_proto(resource.quota_location_name):
            res.quota_location_name = Primitive.to_proto(resource.quota_location_name)
        return res

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return None

        return InstanceNotifyKeyUnavailableRecipeStepsQuotaRequestDeltas(
            metric_name=Primitive.from_proto(resource.metric_name),
            amount=Primitive.from_proto(resource.amount),
            quota_location_name=Primitive.from_proto(resource.quota_location_name),
        )


class InstanceNotifyKeyUnavailableRecipeStepsQuotaRequestDeltasArray(object):
    @classmethod
    def to_proto(self, resources):
        if not resources:
            return resources
        return [
            InstanceNotifyKeyUnavailableRecipeStepsQuotaRequestDeltas.to_proto(i)
            for i in resources
        ]

    @classmethod
    def from_proto(self, resources):
        return [
            InstanceNotifyKeyUnavailableRecipeStepsQuotaRequestDeltas.from_proto(i)
            for i in resources
        ]


class InstanceNotifyKeyUnavailableRecipeStepsPreprocessUpdate(object):
    def __init__(
        self, latency_slo_bucket_name: str = None, public_operation_metadata: str = None
    ):
        self.latency_slo_bucket_name = latency_slo_bucket_name
        self.public_operation_metadata = public_operation_metadata

    @classmethod
    def to_proto(self, resource):
        if not resource:
            return None

        res = (
            instance_pb2.Tier2AlphaInstanceNotifyKeyUnavailableRecipeStepsPreprocessUpdate()
        )
        if Primitive.to_proto(resource.latency_slo_bucket_name):
            res.latency_slo_bucket_name = Primitive.to_proto(
                resource.latency_slo_bucket_name
            )
        if Primitive.to_proto(resource.public_operation_metadata):
            res.public_operation_metadata = Primitive.to_proto(
                resource.public_operation_metadata
            )
        return res

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return None

        return InstanceNotifyKeyUnavailableRecipeStepsPreprocessUpdate(
            latency_slo_bucket_name=Primitive.from_proto(
                resource.latency_slo_bucket_name
            ),
            public_operation_metadata=Primitive.from_proto(
                resource.public_operation_metadata
            ),
        )


class InstanceNotifyKeyUnavailableRecipeStepsPreprocessUpdateArray(object):
    @classmethod
    def to_proto(self, resources):
        if not resources:
            return resources
        return [
            InstanceNotifyKeyUnavailableRecipeStepsPreprocessUpdate.to_proto(i)
            for i in resources
        ]

    @classmethod
    def from_proto(self, resources):
        return [
            InstanceNotifyKeyUnavailableRecipeStepsPreprocessUpdate.from_proto(i)
            for i in resources
        ]


class InstanceNotifyKeyUnavailableRecipeStepsRequestedTenantProject(object):
    def __init__(self, tag: str = None, folder: str = None, scope: str = None):
        self.tag = tag
        self.folder = folder
        self.scope = scope

    @classmethod
    def to_proto(self, resource):
        if not resource:
            return None

        res = (
            instance_pb2.Tier2AlphaInstanceNotifyKeyUnavailableRecipeStepsRequestedTenantProject()
        )
        if Primitive.to_proto(resource.tag):
            res.tag = Primitive.to_proto(resource.tag)
        if Primitive.to_proto(resource.folder):
            res.folder = Primitive.to_proto(resource.folder)
        if InstanceNotifyKeyUnavailableRecipeStepsRequestedTenantProjectScopeEnum.to_proto(
            resource.scope
        ):
            res.scope = InstanceNotifyKeyUnavailableRecipeStepsRequestedTenantProjectScopeEnum.to_proto(
                resource.scope
            )
        return res

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return None

        return InstanceNotifyKeyUnavailableRecipeStepsRequestedTenantProject(
            tag=Primitive.from_proto(resource.tag),
            folder=Primitive.from_proto(resource.folder),
            scope=InstanceNotifyKeyUnavailableRecipeStepsRequestedTenantProjectScopeEnum.from_proto(
                resource.scope
            ),
        )


class InstanceNotifyKeyUnavailableRecipeStepsRequestedTenantProjectArray(object):
    @classmethod
    def to_proto(self, resources):
        if not resources:
            return resources
        return [
            InstanceNotifyKeyUnavailableRecipeStepsRequestedTenantProject.to_proto(i)
            for i in resources
        ]

    @classmethod
    def from_proto(self, resources):
        return [
            InstanceNotifyKeyUnavailableRecipeStepsRequestedTenantProject.from_proto(i)
            for i in resources
        ]


class InstanceNotifyKeyUnavailableRecipeStepsPermissionsInfo(object):
    def __init__(
        self,
        policy_name: dict = None,
        iam_permissions: list = None,
        resource_path: str = None,
        api_attrs: dict = None,
        policy_name_mode: str = None,
        resource: dict = None,
    ):
        self.policy_name = policy_name
        self.iam_permissions = iam_permissions
        self.resource_path = resource_path
        self.api_attrs = api_attrs
        self.policy_name_mode = policy_name_mode
        self.resource = resource

    @classmethod
    def to_proto(self, resource):
        if not resource:
            return None

        res = (
            instance_pb2.Tier2AlphaInstanceNotifyKeyUnavailableRecipeStepsPermissionsInfo()
        )
        if InstanceNotifyKeyUnavailableRecipeStepsPermissionsInfoPolicyName.to_proto(
            resource.policy_name
        ):
            res.policy_name.CopyFrom(
                InstanceNotifyKeyUnavailableRecipeStepsPermissionsInfoPolicyName.to_proto(
                    resource.policy_name
                )
            )
        else:
            res.ClearField("policy_name")
        if InstanceNotifyKeyUnavailableRecipeStepsPermissionsInfoIamPermissionsArray.to_proto(
            resource.iam_permissions
        ):
            res.iam_permissions.extend(
                InstanceNotifyKeyUnavailableRecipeStepsPermissionsInfoIamPermissionsArray.to_proto(
                    resource.iam_permissions
                )
            )
        if Primitive.to_proto(resource.resource_path):
            res.resource_path = Primitive.to_proto(resource.resource_path)
        if InstanceGoogleprotobufstruct.to_proto(resource.api_attrs):
            res.api_attrs.CopyFrom(
                InstanceGoogleprotobufstruct.to_proto(resource.api_attrs)
            )
        else:
            res.ClearField("api_attrs")
        if InstanceNotifyKeyUnavailableRecipeStepsPermissionsInfoPolicyNameModeEnum.to_proto(
            resource.policy_name_mode
        ):
            res.policy_name_mode = InstanceNotifyKeyUnavailableRecipeStepsPermissionsInfoPolicyNameModeEnum.to_proto(
                resource.policy_name_mode
            )
        if InstanceNotifyKeyUnavailableRecipeStepsPermissionsInfoResource.to_proto(
            resource.resource
        ):
            res.resource.CopyFrom(
                InstanceNotifyKeyUnavailableRecipeStepsPermissionsInfoResource.to_proto(
                    resource.resource
                )
            )
        else:
            res.ClearField("resource")
        return res

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return None

        return InstanceNotifyKeyUnavailableRecipeStepsPermissionsInfo(
            policy_name=InstanceNotifyKeyUnavailableRecipeStepsPermissionsInfoPolicyName.from_proto(
                resource.policy_name
            ),
            iam_permissions=InstanceNotifyKeyUnavailableRecipeStepsPermissionsInfoIamPermissionsArray.from_proto(
                resource.iam_permissions
            ),
            resource_path=Primitive.from_proto(resource.resource_path),
            api_attrs=InstanceGoogleprotobufstruct.from_proto(resource.api_attrs),
            policy_name_mode=InstanceNotifyKeyUnavailableRecipeStepsPermissionsInfoPolicyNameModeEnum.from_proto(
                resource.policy_name_mode
            ),
            resource=InstanceNotifyKeyUnavailableRecipeStepsPermissionsInfoResource.from_proto(
                resource.resource
            ),
        )


class InstanceNotifyKeyUnavailableRecipeStepsPermissionsInfoArray(object):
    @classmethod
    def to_proto(self, resources):
        if not resources:
            return resources
        return [
            InstanceNotifyKeyUnavailableRecipeStepsPermissionsInfo.to_proto(i)
            for i in resources
        ]

    @classmethod
    def from_proto(self, resources):
        return [
            InstanceNotifyKeyUnavailableRecipeStepsPermissionsInfo.from_proto(i)
            for i in resources
        ]


class InstanceNotifyKeyUnavailableRecipeStepsPermissionsInfoPolicyName(object):
    def __init__(self, type: str = None, id: str = None, region: str = None):
        self.type = type
        self.id = id
        self.region = region

    @classmethod
    def to_proto(self, resource):
        if not resource:
            return None

        res = (
            instance_pb2.Tier2AlphaInstanceNotifyKeyUnavailableRecipeStepsPermissionsInfoPolicyName()
        )
        if Primitive.to_proto(resource.type):
            res.type = Primitive.to_proto(resource.type)
        if Primitive.to_proto(resource.id):
            res.id = Primitive.to_proto(resource.id)
        if Primitive.to_proto(resource.region):
            res.region = Primitive.to_proto(resource.region)
        return res

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return None

        return InstanceNotifyKeyUnavailableRecipeStepsPermissionsInfoPolicyName(
            type=Primitive.from_proto(resource.type),
            id=Primitive.from_proto(resource.id),
            region=Primitive.from_proto(resource.region),
        )


class InstanceNotifyKeyUnavailableRecipeStepsPermissionsInfoPolicyNameArray(object):
    @classmethod
    def to_proto(self, resources):
        if not resources:
            return resources
        return [
            InstanceNotifyKeyUnavailableRecipeStepsPermissionsInfoPolicyName.to_proto(i)
            for i in resources
        ]

    @classmethod
    def from_proto(self, resources):
        return [
            InstanceNotifyKeyUnavailableRecipeStepsPermissionsInfoPolicyName.from_proto(
                i
            )
            for i in resources
        ]


class InstanceNotifyKeyUnavailableRecipeStepsPermissionsInfoIamPermissions(object):
    def __init__(self, permission: str = None):
        self.permission = permission

    @classmethod
    def to_proto(self, resource):
        if not resource:
            return None

        res = (
            instance_pb2.Tier2AlphaInstanceNotifyKeyUnavailableRecipeStepsPermissionsInfoIamPermissions()
        )
        if Primitive.to_proto(resource.permission):
            res.permission = Primitive.to_proto(resource.permission)
        return res

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return None

        return InstanceNotifyKeyUnavailableRecipeStepsPermissionsInfoIamPermissions(
            permission=Primitive.from_proto(resource.permission),
        )


class InstanceNotifyKeyUnavailableRecipeStepsPermissionsInfoIamPermissionsArray(object):
    @classmethod
    def to_proto(self, resources):
        if not resources:
            return resources
        return [
            InstanceNotifyKeyUnavailableRecipeStepsPermissionsInfoIamPermissions.to_proto(
                i
            )
            for i in resources
        ]

    @classmethod
    def from_proto(self, resources):
        return [
            InstanceNotifyKeyUnavailableRecipeStepsPermissionsInfoIamPermissions.from_proto(
                i
            )
            for i in resources
        ]


class InstanceNotifyKeyUnavailableRecipeStepsPermissionsInfoResource(object):
    def __init__(
        self,
        name: str = None,
        type: str = None,
        service: str = None,
        labels: dict = None,
    ):
        self.name = name
        self.type = type
        self.service = service
        self.labels = labels

    @classmethod
    def to_proto(self, resource):
        if not resource:
            return None

        res = (
            instance_pb2.Tier2AlphaInstanceNotifyKeyUnavailableRecipeStepsPermissionsInfoResource()
        )
        if Primitive.to_proto(resource.name):
            res.name = Primitive.to_proto(resource.name)
        if Primitive.to_proto(resource.type):
            res.type = Primitive.to_proto(resource.type)
        if Primitive.to_proto(resource.service):
            res.service = Primitive.to_proto(resource.service)
        if Primitive.to_proto(resource.labels):
            res.labels = Primitive.to_proto(resource.labels)
        return res

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return None

        return InstanceNotifyKeyUnavailableRecipeStepsPermissionsInfoResource(
            name=Primitive.from_proto(resource.name),
            type=Primitive.from_proto(resource.type),
            service=Primitive.from_proto(resource.service),
            labels=Primitive.from_proto(resource.labels),
        )


class InstanceNotifyKeyUnavailableRecipeStepsPermissionsInfoResourceArray(object):
    @classmethod
    def to_proto(self, resources):
        if not resources:
            return resources
        return [
            InstanceNotifyKeyUnavailableRecipeStepsPermissionsInfoResource.to_proto(i)
            for i in resources
        ]

    @classmethod
    def from_proto(self, resources):
        return [
            InstanceNotifyKeyUnavailableRecipeStepsPermissionsInfoResource.from_proto(i)
            for i in resources
        ]


class InstanceNotifyKeyUnavailableRecipeStepsKeyNotificationsUpdate(object):
    def __init__(self, key_notifications_info: dict = None):
        self.key_notifications_info = key_notifications_info

    @classmethod
    def to_proto(self, resource):
        if not resource:
            return None

        res = (
            instance_pb2.Tier2AlphaInstanceNotifyKeyUnavailableRecipeStepsKeyNotificationsUpdate()
        )
        if InstanceNotifyKeyUnavailableRecipeStepsKeyNotificationsUpdateKeyNotificationsInfo.to_proto(
            resource.key_notifications_info
        ):
            res.key_notifications_info.CopyFrom(
                InstanceNotifyKeyUnavailableRecipeStepsKeyNotificationsUpdateKeyNotificationsInfo.to_proto(
                    resource.key_notifications_info
                )
            )
        else:
            res.ClearField("key_notifications_info")
        return res

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return None

        return InstanceNotifyKeyUnavailableRecipeStepsKeyNotificationsUpdate(
            key_notifications_info=InstanceNotifyKeyUnavailableRecipeStepsKeyNotificationsUpdateKeyNotificationsInfo.from_proto(
                resource.key_notifications_info
            ),
        )


class InstanceNotifyKeyUnavailableRecipeStepsKeyNotificationsUpdateArray(object):
    @classmethod
    def to_proto(self, resources):
        if not resources:
            return resources
        return [
            InstanceNotifyKeyUnavailableRecipeStepsKeyNotificationsUpdate.to_proto(i)
            for i in resources
        ]

    @classmethod
    def from_proto(self, resources):
        return [
            InstanceNotifyKeyUnavailableRecipeStepsKeyNotificationsUpdate.from_proto(i)
            for i in resources
        ]


class InstanceNotifyKeyUnavailableRecipeStepsKeyNotificationsUpdateKeyNotificationsInfo(
    object
):
    def __init__(
        self,
        data_version: int = None,
        delegate: str = None,
        key_notification_configs: list = None,
    ):
        self.data_version = data_version
        self.delegate = delegate
        self.key_notification_configs = key_notification_configs

    @classmethod
    def to_proto(self, resource):
        if not resource:
            return None

        res = (
            instance_pb2.Tier2AlphaInstanceNotifyKeyUnavailableRecipeStepsKeyNotificationsUpdateKeyNotificationsInfo()
        )
        if Primitive.to_proto(resource.data_version):
            res.data_version = Primitive.to_proto(resource.data_version)
        if Primitive.to_proto(resource.delegate):
            res.delegate = Primitive.to_proto(resource.delegate)
        if InstanceNotifyKeyUnavailableRecipeStepsKeyNotificationsUpdateKeyNotificationsInfoKeyNotificationConfigsArray.to_proto(
            resource.key_notification_configs
        ):
            res.key_notification_configs.extend(
                InstanceNotifyKeyUnavailableRecipeStepsKeyNotificationsUpdateKeyNotificationsInfoKeyNotificationConfigsArray.to_proto(
                    resource.key_notification_configs
                )
            )
        return res

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return None

        return InstanceNotifyKeyUnavailableRecipeStepsKeyNotificationsUpdateKeyNotificationsInfo(
            data_version=Primitive.from_proto(resource.data_version),
            delegate=Primitive.from_proto(resource.delegate),
            key_notification_configs=InstanceNotifyKeyUnavailableRecipeStepsKeyNotificationsUpdateKeyNotificationsInfoKeyNotificationConfigsArray.from_proto(
                resource.key_notification_configs
            ),
        )


class InstanceNotifyKeyUnavailableRecipeStepsKeyNotificationsUpdateKeyNotificationsInfoArray(
    object
):
    @classmethod
    def to_proto(self, resources):
        if not resources:
            return resources
        return [
            InstanceNotifyKeyUnavailableRecipeStepsKeyNotificationsUpdateKeyNotificationsInfo.to_proto(
                i
            )
            for i in resources
        ]

    @classmethod
    def from_proto(self, resources):
        return [
            InstanceNotifyKeyUnavailableRecipeStepsKeyNotificationsUpdateKeyNotificationsInfo.from_proto(
                i
            )
            for i in resources
        ]


class InstanceNotifyKeyUnavailableRecipeStepsKeyNotificationsUpdateKeyNotificationsInfoKeyNotificationConfigs(
    object
):
    def __init__(
        self,
        key_or_version_name: str = None,
        grant: str = None,
        delegator_gaia_id: int = None,
    ):
        self.key_or_version_name = key_or_version_name
        self.grant = grant
        self.delegator_gaia_id = delegator_gaia_id

    @classmethod
    def to_proto(self, resource):
        if not resource:
            return None

        res = (
            instance_pb2.Tier2AlphaInstanceNotifyKeyUnavailableRecipeStepsKeyNotificationsUpdateKeyNotificationsInfoKeyNotificationConfigs()
        )
        if Primitive.to_proto(resource.key_or_version_name):
            res.key_or_version_name = Primitive.to_proto(resource.key_or_version_name)
        if Primitive.to_proto(resource.grant):
            res.grant = Primitive.to_proto(resource.grant)
        if Primitive.to_proto(resource.delegator_gaia_id):
            res.delegator_gaia_id = Primitive.to_proto(resource.delegator_gaia_id)
        return res

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return None

        return InstanceNotifyKeyUnavailableRecipeStepsKeyNotificationsUpdateKeyNotificationsInfoKeyNotificationConfigs(
            key_or_version_name=Primitive.from_proto(resource.key_or_version_name),
            grant=Primitive.from_proto(resource.grant),
            delegator_gaia_id=Primitive.from_proto(resource.delegator_gaia_id),
        )


class InstanceNotifyKeyUnavailableRecipeStepsKeyNotificationsUpdateKeyNotificationsInfoKeyNotificationConfigsArray(
    object
):
    @classmethod
    def to_proto(self, resources):
        if not resources:
            return resources
        return [
            InstanceNotifyKeyUnavailableRecipeStepsKeyNotificationsUpdateKeyNotificationsInfoKeyNotificationConfigs.to_proto(
                i
            )
            for i in resources
        ]

    @classmethod
    def from_proto(self, resources):
        return [
            InstanceNotifyKeyUnavailableRecipeStepsKeyNotificationsUpdateKeyNotificationsInfoKeyNotificationConfigs.from_proto(
                i
            )
            for i in resources
        ]


class InstanceReadonlyRecipe(object):
    def __init__(
        self,
        steps: list = None,
        honor_cancel_request: bool = None,
        ignore_recipe_after: int = None,
        verify_deadline_seconds_below: float = None,
        populate_operation_result: bool = None,
        readonly_recipe_start_time: str = None,
        resource_names_stored_in_clh_with_delay: list = None,
        delay_to_store_resources_in_clh_db_nanos: int = None,
    ):
        self.steps = steps
        self.honor_cancel_request = honor_cancel_request
        self.ignore_recipe_after = ignore_recipe_after
        self.verify_deadline_seconds_below = verify_deadline_seconds_below
        self.populate_operation_result = populate_operation_result
        self.readonly_recipe_start_time = readonly_recipe_start_time
        self.resource_names_stored_in_clh_with_delay = (
            resource_names_stored_in_clh_with_delay
        )
        self.delay_to_store_resources_in_clh_db_nanos = (
            delay_to_store_resources_in_clh_db_nanos
        )

    @classmethod
    def to_proto(self, resource):
        if not resource:
            return None

        res = instance_pb2.Tier2AlphaInstanceReadonlyRecipe()
        if InstanceReadonlyRecipeStepsArray.to_proto(resource.steps):
            res.steps.extend(InstanceReadonlyRecipeStepsArray.to_proto(resource.steps))
        if Primitive.to_proto(resource.honor_cancel_request):
            res.honor_cancel_request = Primitive.to_proto(resource.honor_cancel_request)
        if Primitive.to_proto(resource.ignore_recipe_after):
            res.ignore_recipe_after = Primitive.to_proto(resource.ignore_recipe_after)
        if Primitive.to_proto(resource.verify_deadline_seconds_below):
            res.verify_deadline_seconds_below = Primitive.to_proto(
                resource.verify_deadline_seconds_below
            )
        if Primitive.to_proto(resource.populate_operation_result):
            res.populate_operation_result = Primitive.to_proto(
                resource.populate_operation_result
            )
        if Primitive.to_proto(resource.readonly_recipe_start_time):
            res.readonly_recipe_start_time = Primitive.to_proto(
                resource.readonly_recipe_start_time
            )
        if Primitive.to_proto(resource.resource_names_stored_in_clh_with_delay):
            res.resource_names_stored_in_clh_with_delay.extend(
                Primitive.to_proto(resource.resource_names_stored_in_clh_with_delay)
            )
        if Primitive.to_proto(resource.delay_to_store_resources_in_clh_db_nanos):
            res.delay_to_store_resources_in_clh_db_nanos = Primitive.to_proto(
                resource.delay_to_store_resources_in_clh_db_nanos
            )
        return res

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return None

        return InstanceReadonlyRecipe(
            steps=InstanceReadonlyRecipeStepsArray.from_proto(resource.steps),
            honor_cancel_request=Primitive.from_proto(resource.honor_cancel_request),
            ignore_recipe_after=Primitive.from_proto(resource.ignore_recipe_after),
            verify_deadline_seconds_below=Primitive.from_proto(
                resource.verify_deadline_seconds_below
            ),
            populate_operation_result=Primitive.from_proto(
                resource.populate_operation_result
            ),
            readonly_recipe_start_time=Primitive.from_proto(
                resource.readonly_recipe_start_time
            ),
            resource_names_stored_in_clh_with_delay=Primitive.from_proto(
                resource.resource_names_stored_in_clh_with_delay
            ),
            delay_to_store_resources_in_clh_db_nanos=Primitive.from_proto(
                resource.delay_to_store_resources_in_clh_db_nanos
            ),
        )


class InstanceReadonlyRecipeArray(object):
    @classmethod
    def to_proto(self, resources):
        if not resources:
            return resources
        return [InstanceReadonlyRecipe.to_proto(i) for i in resources]

    @classmethod
    def from_proto(self, resources):
        return [InstanceReadonlyRecipe.from_proto(i) for i in resources]


class InstanceReadonlyRecipeSteps(object):
    def __init__(
        self,
        relative_time: int = None,
        sleep_duration: int = None,
        action: str = None,
        status: dict = None,
        error_space: str = None,
        p4_service_account: str = None,
        resource_metadata_size: int = None,
        description: str = None,
        updated_repeat_operation_delay_sec: float = None,
        quota_request_deltas: list = None,
        preprocess_update: dict = None,
        public_operation_metadata: str = None,
        requested_tenant_project: dict = None,
        permissions_info: list = None,
        key_notifications_update: dict = None,
        clh_data_update_time: str = None,
        public_error_message: str = None,
    ):
        self.relative_time = relative_time
        self.sleep_duration = sleep_duration
        self.action = action
        self.status = status
        self.error_space = error_space
        self.p4_service_account = p4_service_account
        self.resource_metadata_size = resource_metadata_size
        self.description = description
        self.updated_repeat_operation_delay_sec = updated_repeat_operation_delay_sec
        self.quota_request_deltas = quota_request_deltas
        self.preprocess_update = preprocess_update
        self.public_operation_metadata = public_operation_metadata
        self.requested_tenant_project = requested_tenant_project
        self.permissions_info = permissions_info
        self.key_notifications_update = key_notifications_update
        self.clh_data_update_time = clh_data_update_time
        self.public_error_message = public_error_message

    @classmethod
    def to_proto(self, resource):
        if not resource:
            return None

        res = instance_pb2.Tier2AlphaInstanceReadonlyRecipeSteps()
        if Primitive.to_proto(resource.relative_time):
            res.relative_time = Primitive.to_proto(resource.relative_time)
        if Primitive.to_proto(resource.sleep_duration):
            res.sleep_duration = Primitive.to_proto(resource.sleep_duration)
        if InstanceReadonlyRecipeStepsActionEnum.to_proto(resource.action):
            res.action = InstanceReadonlyRecipeStepsActionEnum.to_proto(resource.action)
        if InstanceReadonlyRecipeStepsStatus.to_proto(resource.status):
            res.status.CopyFrom(
                InstanceReadonlyRecipeStepsStatus.to_proto(resource.status)
            )
        else:
            res.ClearField("status")
        if Primitive.to_proto(resource.error_space):
            res.error_space = Primitive.to_proto(resource.error_space)
        if Primitive.to_proto(resource.p4_service_account):
            res.p4_service_account = Primitive.to_proto(resource.p4_service_account)
        if Primitive.to_proto(resource.resource_metadata_size):
            res.resource_metadata_size = Primitive.to_proto(
                resource.resource_metadata_size
            )
        if Primitive.to_proto(resource.description):
            res.description = Primitive.to_proto(resource.description)
        if Primitive.to_proto(resource.updated_repeat_operation_delay_sec):
            res.updated_repeat_operation_delay_sec = Primitive.to_proto(
                resource.updated_repeat_operation_delay_sec
            )
        if InstanceReadonlyRecipeStepsQuotaRequestDeltasArray.to_proto(
            resource.quota_request_deltas
        ):
            res.quota_request_deltas.extend(
                InstanceReadonlyRecipeStepsQuotaRequestDeltasArray.to_proto(
                    resource.quota_request_deltas
                )
            )
        if InstanceReadonlyRecipeStepsPreprocessUpdate.to_proto(
            resource.preprocess_update
        ):
            res.preprocess_update.CopyFrom(
                InstanceReadonlyRecipeStepsPreprocessUpdate.to_proto(
                    resource.preprocess_update
                )
            )
        else:
            res.ClearField("preprocess_update")
        if Primitive.to_proto(resource.public_operation_metadata):
            res.public_operation_metadata = Primitive.to_proto(
                resource.public_operation_metadata
            )
        if InstanceReadonlyRecipeStepsRequestedTenantProject.to_proto(
            resource.requested_tenant_project
        ):
            res.requested_tenant_project.CopyFrom(
                InstanceReadonlyRecipeStepsRequestedTenantProject.to_proto(
                    resource.requested_tenant_project
                )
            )
        else:
            res.ClearField("requested_tenant_project")
        if InstanceReadonlyRecipeStepsPermissionsInfoArray.to_proto(
            resource.permissions_info
        ):
            res.permissions_info.extend(
                InstanceReadonlyRecipeStepsPermissionsInfoArray.to_proto(
                    resource.permissions_info
                )
            )
        if InstanceReadonlyRecipeStepsKeyNotificationsUpdate.to_proto(
            resource.key_notifications_update
        ):
            res.key_notifications_update.CopyFrom(
                InstanceReadonlyRecipeStepsKeyNotificationsUpdate.to_proto(
                    resource.key_notifications_update
                )
            )
        else:
            res.ClearField("key_notifications_update")
        if Primitive.to_proto(resource.clh_data_update_time):
            res.clh_data_update_time = Primitive.to_proto(resource.clh_data_update_time)
        if Primitive.to_proto(resource.public_error_message):
            res.public_error_message = Primitive.to_proto(resource.public_error_message)
        return res

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return None

        return InstanceReadonlyRecipeSteps(
            relative_time=Primitive.from_proto(resource.relative_time),
            sleep_duration=Primitive.from_proto(resource.sleep_duration),
            action=InstanceReadonlyRecipeStepsActionEnum.from_proto(resource.action),
            status=InstanceReadonlyRecipeStepsStatus.from_proto(resource.status),
            error_space=Primitive.from_proto(resource.error_space),
            p4_service_account=Primitive.from_proto(resource.p4_service_account),
            resource_metadata_size=Primitive.from_proto(
                resource.resource_metadata_size
            ),
            description=Primitive.from_proto(resource.description),
            updated_repeat_operation_delay_sec=Primitive.from_proto(
                resource.updated_repeat_operation_delay_sec
            ),
            quota_request_deltas=InstanceReadonlyRecipeStepsQuotaRequestDeltasArray.from_proto(
                resource.quota_request_deltas
            ),
            preprocess_update=InstanceReadonlyRecipeStepsPreprocessUpdate.from_proto(
                resource.preprocess_update
            ),
            public_operation_metadata=Primitive.from_proto(
                resource.public_operation_metadata
            ),
            requested_tenant_project=InstanceReadonlyRecipeStepsRequestedTenantProject.from_proto(
                resource.requested_tenant_project
            ),
            permissions_info=InstanceReadonlyRecipeStepsPermissionsInfoArray.from_proto(
                resource.permissions_info
            ),
            key_notifications_update=InstanceReadonlyRecipeStepsKeyNotificationsUpdate.from_proto(
                resource.key_notifications_update
            ),
            clh_data_update_time=Primitive.from_proto(resource.clh_data_update_time),
            public_error_message=Primitive.from_proto(resource.public_error_message),
        )


class InstanceReadonlyRecipeStepsArray(object):
    @classmethod
    def to_proto(self, resources):
        if not resources:
            return resources
        return [InstanceReadonlyRecipeSteps.to_proto(i) for i in resources]

    @classmethod
    def from_proto(self, resources):
        return [InstanceReadonlyRecipeSteps.from_proto(i) for i in resources]


class InstanceReadonlyRecipeStepsStatus(object):
    def __init__(self, code: int = None, message: str = None, details: list = None):
        self.code = code
        self.message = message
        self.details = details

    @classmethod
    def to_proto(self, resource):
        if not resource:
            return None

        res = instance_pb2.Tier2AlphaInstanceReadonlyRecipeStepsStatus()
        if Primitive.to_proto(resource.code):
            res.code = Primitive.to_proto(resource.code)
        if Primitive.to_proto(resource.message):
            res.message = Primitive.to_proto(resource.message)
        if InstanceReadonlyRecipeStepsStatusDetailsArray.to_proto(resource.details):
            res.details.extend(
                InstanceReadonlyRecipeStepsStatusDetailsArray.to_proto(resource.details)
            )
        return res

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return None

        return InstanceReadonlyRecipeStepsStatus(
            code=Primitive.from_proto(resource.code),
            message=Primitive.from_proto(resource.message),
            details=InstanceReadonlyRecipeStepsStatusDetailsArray.from_proto(
                resource.details
            ),
        )


class InstanceReadonlyRecipeStepsStatusArray(object):
    @classmethod
    def to_proto(self, resources):
        if not resources:
            return resources
        return [InstanceReadonlyRecipeStepsStatus.to_proto(i) for i in resources]

    @classmethod
    def from_proto(self, resources):
        return [InstanceReadonlyRecipeStepsStatus.from_proto(i) for i in resources]


class InstanceReadonlyRecipeStepsStatusDetails(object):
    def __init__(self, type_url: str = None, value: str = None):
        self.type_url = type_url
        self.value = value

    @classmethod
    def to_proto(self, resource):
        if not resource:
            return None

        res = instance_pb2.Tier2AlphaInstanceReadonlyRecipeStepsStatusDetails()
        if Primitive.to_proto(resource.type_url):
            res.type_url = Primitive.to_proto(resource.type_url)
        if Primitive.to_proto(resource.value):
            res.value = Primitive.to_proto(resource.value)
        return res

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return None

        return InstanceReadonlyRecipeStepsStatusDetails(
            type_url=Primitive.from_proto(resource.type_url),
            value=Primitive.from_proto(resource.value),
        )


class InstanceReadonlyRecipeStepsStatusDetailsArray(object):
    @classmethod
    def to_proto(self, resources):
        if not resources:
            return resources
        return [InstanceReadonlyRecipeStepsStatusDetails.to_proto(i) for i in resources]

    @classmethod
    def from_proto(self, resources):
        return [
            InstanceReadonlyRecipeStepsStatusDetails.from_proto(i) for i in resources
        ]


class InstanceReadonlyRecipeStepsQuotaRequestDeltas(object):
    def __init__(
        self,
        metric_name: str = None,
        amount: int = None,
        quota_location_name: str = None,
    ):
        self.metric_name = metric_name
        self.amount = amount
        self.quota_location_name = quota_location_name

    @classmethod
    def to_proto(self, resource):
        if not resource:
            return None

        res = instance_pb2.Tier2AlphaInstanceReadonlyRecipeStepsQuotaRequestDeltas()
        if Primitive.to_proto(resource.metric_name):
            res.metric_name = Primitive.to_proto(resource.metric_name)
        if Primitive.to_proto(resource.amount):
            res.amount = Primitive.to_proto(resource.amount)
        if Primitive.to_proto(resource.quota_location_name):
            res.quota_location_name = Primitive.to_proto(resource.quota_location_name)
        return res

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return None

        return InstanceReadonlyRecipeStepsQuotaRequestDeltas(
            metric_name=Primitive.from_proto(resource.metric_name),
            amount=Primitive.from_proto(resource.amount),
            quota_location_name=Primitive.from_proto(resource.quota_location_name),
        )


class InstanceReadonlyRecipeStepsQuotaRequestDeltasArray(object):
    @classmethod
    def to_proto(self, resources):
        if not resources:
            return resources
        return [
            InstanceReadonlyRecipeStepsQuotaRequestDeltas.to_proto(i) for i in resources
        ]

    @classmethod
    def from_proto(self, resources):
        return [
            InstanceReadonlyRecipeStepsQuotaRequestDeltas.from_proto(i)
            for i in resources
        ]


class InstanceReadonlyRecipeStepsPreprocessUpdate(object):
    def __init__(
        self, latency_slo_bucket_name: str = None, public_operation_metadata: str = None
    ):
        self.latency_slo_bucket_name = latency_slo_bucket_name
        self.public_operation_metadata = public_operation_metadata

    @classmethod
    def to_proto(self, resource):
        if not resource:
            return None

        res = instance_pb2.Tier2AlphaInstanceReadonlyRecipeStepsPreprocessUpdate()
        if Primitive.to_proto(resource.latency_slo_bucket_name):
            res.latency_slo_bucket_name = Primitive.to_proto(
                resource.latency_slo_bucket_name
            )
        if Primitive.to_proto(resource.public_operation_metadata):
            res.public_operation_metadata = Primitive.to_proto(
                resource.public_operation_metadata
            )
        return res

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return None

        return InstanceReadonlyRecipeStepsPreprocessUpdate(
            latency_slo_bucket_name=Primitive.from_proto(
                resource.latency_slo_bucket_name
            ),
            public_operation_metadata=Primitive.from_proto(
                resource.public_operation_metadata
            ),
        )


class InstanceReadonlyRecipeStepsPreprocessUpdateArray(object):
    @classmethod
    def to_proto(self, resources):
        if not resources:
            return resources
        return [
            InstanceReadonlyRecipeStepsPreprocessUpdate.to_proto(i) for i in resources
        ]

    @classmethod
    def from_proto(self, resources):
        return [
            InstanceReadonlyRecipeStepsPreprocessUpdate.from_proto(i) for i in resources
        ]


class InstanceReadonlyRecipeStepsRequestedTenantProject(object):
    def __init__(self, tag: str = None, folder: str = None, scope: str = None):
        self.tag = tag
        self.folder = folder
        self.scope = scope

    @classmethod
    def to_proto(self, resource):
        if not resource:
            return None

        res = instance_pb2.Tier2AlphaInstanceReadonlyRecipeStepsRequestedTenantProject()
        if Primitive.to_proto(resource.tag):
            res.tag = Primitive.to_proto(resource.tag)
        if Primitive.to_proto(resource.folder):
            res.folder = Primitive.to_proto(resource.folder)
        if InstanceReadonlyRecipeStepsRequestedTenantProjectScopeEnum.to_proto(
            resource.scope
        ):
            res.scope = InstanceReadonlyRecipeStepsRequestedTenantProjectScopeEnum.to_proto(
                resource.scope
            )
        return res

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return None

        return InstanceReadonlyRecipeStepsRequestedTenantProject(
            tag=Primitive.from_proto(resource.tag),
            folder=Primitive.from_proto(resource.folder),
            scope=InstanceReadonlyRecipeStepsRequestedTenantProjectScopeEnum.from_proto(
                resource.scope
            ),
        )


class InstanceReadonlyRecipeStepsRequestedTenantProjectArray(object):
    @classmethod
    def to_proto(self, resources):
        if not resources:
            return resources
        return [
            InstanceReadonlyRecipeStepsRequestedTenantProject.to_proto(i)
            for i in resources
        ]

    @classmethod
    def from_proto(self, resources):
        return [
            InstanceReadonlyRecipeStepsRequestedTenantProject.from_proto(i)
            for i in resources
        ]


class InstanceReadonlyRecipeStepsPermissionsInfo(object):
    def __init__(
        self,
        policy_name: dict = None,
        iam_permissions: list = None,
        resource_path: str = None,
        api_attrs: dict = None,
        policy_name_mode: str = None,
        resource: dict = None,
    ):
        self.policy_name = policy_name
        self.iam_permissions = iam_permissions
        self.resource_path = resource_path
        self.api_attrs = api_attrs
        self.policy_name_mode = policy_name_mode
        self.resource = resource

    @classmethod
    def to_proto(self, resource):
        if not resource:
            return None

        res = instance_pb2.Tier2AlphaInstanceReadonlyRecipeStepsPermissionsInfo()
        if InstanceReadonlyRecipeStepsPermissionsInfoPolicyName.to_proto(
            resource.policy_name
        ):
            res.policy_name.CopyFrom(
                InstanceReadonlyRecipeStepsPermissionsInfoPolicyName.to_proto(
                    resource.policy_name
                )
            )
        else:
            res.ClearField("policy_name")
        if InstanceReadonlyRecipeStepsPermissionsInfoIamPermissionsArray.to_proto(
            resource.iam_permissions
        ):
            res.iam_permissions.extend(
                InstanceReadonlyRecipeStepsPermissionsInfoIamPermissionsArray.to_proto(
                    resource.iam_permissions
                )
            )
        if Primitive.to_proto(resource.resource_path):
            res.resource_path = Primitive.to_proto(resource.resource_path)
        if InstanceGoogleprotobufstruct.to_proto(resource.api_attrs):
            res.api_attrs.CopyFrom(
                InstanceGoogleprotobufstruct.to_proto(resource.api_attrs)
            )
        else:
            res.ClearField("api_attrs")
        if InstanceReadonlyRecipeStepsPermissionsInfoPolicyNameModeEnum.to_proto(
            resource.policy_name_mode
        ):
            res.policy_name_mode = InstanceReadonlyRecipeStepsPermissionsInfoPolicyNameModeEnum.to_proto(
                resource.policy_name_mode
            )
        if InstanceReadonlyRecipeStepsPermissionsInfoResource.to_proto(
            resource.resource
        ):
            res.resource.CopyFrom(
                InstanceReadonlyRecipeStepsPermissionsInfoResource.to_proto(
                    resource.resource
                )
            )
        else:
            res.ClearField("resource")
        return res

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return None

        return InstanceReadonlyRecipeStepsPermissionsInfo(
            policy_name=InstanceReadonlyRecipeStepsPermissionsInfoPolicyName.from_proto(
                resource.policy_name
            ),
            iam_permissions=InstanceReadonlyRecipeStepsPermissionsInfoIamPermissionsArray.from_proto(
                resource.iam_permissions
            ),
            resource_path=Primitive.from_proto(resource.resource_path),
            api_attrs=InstanceGoogleprotobufstruct.from_proto(resource.api_attrs),
            policy_name_mode=InstanceReadonlyRecipeStepsPermissionsInfoPolicyNameModeEnum.from_proto(
                resource.policy_name_mode
            ),
            resource=InstanceReadonlyRecipeStepsPermissionsInfoResource.from_proto(
                resource.resource
            ),
        )


class InstanceReadonlyRecipeStepsPermissionsInfoArray(object):
    @classmethod
    def to_proto(self, resources):
        if not resources:
            return resources
        return [
            InstanceReadonlyRecipeStepsPermissionsInfo.to_proto(i) for i in resources
        ]

    @classmethod
    def from_proto(self, resources):
        return [
            InstanceReadonlyRecipeStepsPermissionsInfo.from_proto(i) for i in resources
        ]


class InstanceReadonlyRecipeStepsPermissionsInfoPolicyName(object):
    def __init__(self, type: str = None, id: str = None, region: str = None):
        self.type = type
        self.id = id
        self.region = region

    @classmethod
    def to_proto(self, resource):
        if not resource:
            return None

        res = (
            instance_pb2.Tier2AlphaInstanceReadonlyRecipeStepsPermissionsInfoPolicyName()
        )
        if Primitive.to_proto(resource.type):
            res.type = Primitive.to_proto(resource.type)
        if Primitive.to_proto(resource.id):
            res.id = Primitive.to_proto(resource.id)
        if Primitive.to_proto(resource.region):
            res.region = Primitive.to_proto(resource.region)
        return res

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return None

        return InstanceReadonlyRecipeStepsPermissionsInfoPolicyName(
            type=Primitive.from_proto(resource.type),
            id=Primitive.from_proto(resource.id),
            region=Primitive.from_proto(resource.region),
        )


class InstanceReadonlyRecipeStepsPermissionsInfoPolicyNameArray(object):
    @classmethod
    def to_proto(self, resources):
        if not resources:
            return resources
        return [
            InstanceReadonlyRecipeStepsPermissionsInfoPolicyName.to_proto(i)
            for i in resources
        ]

    @classmethod
    def from_proto(self, resources):
        return [
            InstanceReadonlyRecipeStepsPermissionsInfoPolicyName.from_proto(i)
            for i in resources
        ]


class InstanceReadonlyRecipeStepsPermissionsInfoIamPermissions(object):
    def __init__(self, permission: str = None):
        self.permission = permission

    @classmethod
    def to_proto(self, resource):
        if not resource:
            return None

        res = (
            instance_pb2.Tier2AlphaInstanceReadonlyRecipeStepsPermissionsInfoIamPermissions()
        )
        if Primitive.to_proto(resource.permission):
            res.permission = Primitive.to_proto(resource.permission)
        return res

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return None

        return InstanceReadonlyRecipeStepsPermissionsInfoIamPermissions(
            permission=Primitive.from_proto(resource.permission),
        )


class InstanceReadonlyRecipeStepsPermissionsInfoIamPermissionsArray(object):
    @classmethod
    def to_proto(self, resources):
        if not resources:
            return resources
        return [
            InstanceReadonlyRecipeStepsPermissionsInfoIamPermissions.to_proto(i)
            for i in resources
        ]

    @classmethod
    def from_proto(self, resources):
        return [
            InstanceReadonlyRecipeStepsPermissionsInfoIamPermissions.from_proto(i)
            for i in resources
        ]


class InstanceReadonlyRecipeStepsPermissionsInfoResource(object):
    def __init__(
        self,
        name: str = None,
        type: str = None,
        service: str = None,
        labels: dict = None,
    ):
        self.name = name
        self.type = type
        self.service = service
        self.labels = labels

    @classmethod
    def to_proto(self, resource):
        if not resource:
            return None

        res = (
            instance_pb2.Tier2AlphaInstanceReadonlyRecipeStepsPermissionsInfoResource()
        )
        if Primitive.to_proto(resource.name):
            res.name = Primitive.to_proto(resource.name)
        if Primitive.to_proto(resource.type):
            res.type = Primitive.to_proto(resource.type)
        if Primitive.to_proto(resource.service):
            res.service = Primitive.to_proto(resource.service)
        if Primitive.to_proto(resource.labels):
            res.labels = Primitive.to_proto(resource.labels)
        return res

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return None

        return InstanceReadonlyRecipeStepsPermissionsInfoResource(
            name=Primitive.from_proto(resource.name),
            type=Primitive.from_proto(resource.type),
            service=Primitive.from_proto(resource.service),
            labels=Primitive.from_proto(resource.labels),
        )


class InstanceReadonlyRecipeStepsPermissionsInfoResourceArray(object):
    @classmethod
    def to_proto(self, resources):
        if not resources:
            return resources
        return [
            InstanceReadonlyRecipeStepsPermissionsInfoResource.to_proto(i)
            for i in resources
        ]

    @classmethod
    def from_proto(self, resources):
        return [
            InstanceReadonlyRecipeStepsPermissionsInfoResource.from_proto(i)
            for i in resources
        ]


class InstanceReadonlyRecipeStepsKeyNotificationsUpdate(object):
    def __init__(self, key_notifications_info: dict = None):
        self.key_notifications_info = key_notifications_info

    @classmethod
    def to_proto(self, resource):
        if not resource:
            return None

        res = instance_pb2.Tier2AlphaInstanceReadonlyRecipeStepsKeyNotificationsUpdate()
        if InstanceReadonlyRecipeStepsKeyNotificationsUpdateKeyNotificationsInfo.to_proto(
            resource.key_notifications_info
        ):
            res.key_notifications_info.CopyFrom(
                InstanceReadonlyRecipeStepsKeyNotificationsUpdateKeyNotificationsInfo.to_proto(
                    resource.key_notifications_info
                )
            )
        else:
            res.ClearField("key_notifications_info")
        return res

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return None

        return InstanceReadonlyRecipeStepsKeyNotificationsUpdate(
            key_notifications_info=InstanceReadonlyRecipeStepsKeyNotificationsUpdateKeyNotificationsInfo.from_proto(
                resource.key_notifications_info
            ),
        )


class InstanceReadonlyRecipeStepsKeyNotificationsUpdateArray(object):
    @classmethod
    def to_proto(self, resources):
        if not resources:
            return resources
        return [
            InstanceReadonlyRecipeStepsKeyNotificationsUpdate.to_proto(i)
            for i in resources
        ]

    @classmethod
    def from_proto(self, resources):
        return [
            InstanceReadonlyRecipeStepsKeyNotificationsUpdate.from_proto(i)
            for i in resources
        ]


class InstanceReadonlyRecipeStepsKeyNotificationsUpdateKeyNotificationsInfo(object):
    def __init__(
        self,
        data_version: int = None,
        delegate: str = None,
        key_notification_configs: list = None,
    ):
        self.data_version = data_version
        self.delegate = delegate
        self.key_notification_configs = key_notification_configs

    @classmethod
    def to_proto(self, resource):
        if not resource:
            return None

        res = (
            instance_pb2.Tier2AlphaInstanceReadonlyRecipeStepsKeyNotificationsUpdateKeyNotificationsInfo()
        )
        if Primitive.to_proto(resource.data_version):
            res.data_version = Primitive.to_proto(resource.data_version)
        if Primitive.to_proto(resource.delegate):
            res.delegate = Primitive.to_proto(resource.delegate)
        if InstanceReadonlyRecipeStepsKeyNotificationsUpdateKeyNotificationsInfoKeyNotificationConfigsArray.to_proto(
            resource.key_notification_configs
        ):
            res.key_notification_configs.extend(
                InstanceReadonlyRecipeStepsKeyNotificationsUpdateKeyNotificationsInfoKeyNotificationConfigsArray.to_proto(
                    resource.key_notification_configs
                )
            )
        return res

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return None

        return InstanceReadonlyRecipeStepsKeyNotificationsUpdateKeyNotificationsInfo(
            data_version=Primitive.from_proto(resource.data_version),
            delegate=Primitive.from_proto(resource.delegate),
            key_notification_configs=InstanceReadonlyRecipeStepsKeyNotificationsUpdateKeyNotificationsInfoKeyNotificationConfigsArray.from_proto(
                resource.key_notification_configs
            ),
        )


class InstanceReadonlyRecipeStepsKeyNotificationsUpdateKeyNotificationsInfoArray(
    object
):
    @classmethod
    def to_proto(self, resources):
        if not resources:
            return resources
        return [
            InstanceReadonlyRecipeStepsKeyNotificationsUpdateKeyNotificationsInfo.to_proto(
                i
            )
            for i in resources
        ]

    @classmethod
    def from_proto(self, resources):
        return [
            InstanceReadonlyRecipeStepsKeyNotificationsUpdateKeyNotificationsInfo.from_proto(
                i
            )
            for i in resources
        ]


class InstanceReadonlyRecipeStepsKeyNotificationsUpdateKeyNotificationsInfoKeyNotificationConfigs(
    object
):
    def __init__(
        self,
        key_or_version_name: str = None,
        grant: str = None,
        delegator_gaia_id: int = None,
    ):
        self.key_or_version_name = key_or_version_name
        self.grant = grant
        self.delegator_gaia_id = delegator_gaia_id

    @classmethod
    def to_proto(self, resource):
        if not resource:
            return None

        res = (
            instance_pb2.Tier2AlphaInstanceReadonlyRecipeStepsKeyNotificationsUpdateKeyNotificationsInfoKeyNotificationConfigs()
        )
        if Primitive.to_proto(resource.key_or_version_name):
            res.key_or_version_name = Primitive.to_proto(resource.key_or_version_name)
        if Primitive.to_proto(resource.grant):
            res.grant = Primitive.to_proto(resource.grant)
        if Primitive.to_proto(resource.delegator_gaia_id):
            res.delegator_gaia_id = Primitive.to_proto(resource.delegator_gaia_id)
        return res

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return None

        return InstanceReadonlyRecipeStepsKeyNotificationsUpdateKeyNotificationsInfoKeyNotificationConfigs(
            key_or_version_name=Primitive.from_proto(resource.key_or_version_name),
            grant=Primitive.from_proto(resource.grant),
            delegator_gaia_id=Primitive.from_proto(resource.delegator_gaia_id),
        )


class InstanceReadonlyRecipeStepsKeyNotificationsUpdateKeyNotificationsInfoKeyNotificationConfigsArray(
    object
):
    @classmethod
    def to_proto(self, resources):
        if not resources:
            return resources
        return [
            InstanceReadonlyRecipeStepsKeyNotificationsUpdateKeyNotificationsInfoKeyNotificationConfigs.to_proto(
                i
            )
            for i in resources
        ]

    @classmethod
    def from_proto(self, resources):
        return [
            InstanceReadonlyRecipeStepsKeyNotificationsUpdateKeyNotificationsInfoKeyNotificationConfigs.from_proto(
                i
            )
            for i in resources
        ]


class InstanceReconcileRecipe(object):
    def __init__(
        self,
        steps: list = None,
        honor_cancel_request: bool = None,
        ignore_recipe_after: int = None,
        verify_deadline_seconds_below: float = None,
        populate_operation_result: bool = None,
        readonly_recipe_start_time: str = None,
        resource_names_stored_in_clh_with_delay: list = None,
        delay_to_store_resources_in_clh_db_nanos: int = None,
    ):
        self.steps = steps
        self.honor_cancel_request = honor_cancel_request
        self.ignore_recipe_after = ignore_recipe_after
        self.verify_deadline_seconds_below = verify_deadline_seconds_below
        self.populate_operation_result = populate_operation_result
        self.readonly_recipe_start_time = readonly_recipe_start_time
        self.resource_names_stored_in_clh_with_delay = (
            resource_names_stored_in_clh_with_delay
        )
        self.delay_to_store_resources_in_clh_db_nanos = (
            delay_to_store_resources_in_clh_db_nanos
        )

    @classmethod
    def to_proto(self, resource):
        if not resource:
            return None

        res = instance_pb2.Tier2AlphaInstanceReconcileRecipe()
        if InstanceReconcileRecipeStepsArray.to_proto(resource.steps):
            res.steps.extend(InstanceReconcileRecipeStepsArray.to_proto(resource.steps))
        if Primitive.to_proto(resource.honor_cancel_request):
            res.honor_cancel_request = Primitive.to_proto(resource.honor_cancel_request)
        if Primitive.to_proto(resource.ignore_recipe_after):
            res.ignore_recipe_after = Primitive.to_proto(resource.ignore_recipe_after)
        if Primitive.to_proto(resource.verify_deadline_seconds_below):
            res.verify_deadline_seconds_below = Primitive.to_proto(
                resource.verify_deadline_seconds_below
            )
        if Primitive.to_proto(resource.populate_operation_result):
            res.populate_operation_result = Primitive.to_proto(
                resource.populate_operation_result
            )
        if Primitive.to_proto(resource.readonly_recipe_start_time):
            res.readonly_recipe_start_time = Primitive.to_proto(
                resource.readonly_recipe_start_time
            )
        if Primitive.to_proto(resource.resource_names_stored_in_clh_with_delay):
            res.resource_names_stored_in_clh_with_delay.extend(
                Primitive.to_proto(resource.resource_names_stored_in_clh_with_delay)
            )
        if Primitive.to_proto(resource.delay_to_store_resources_in_clh_db_nanos):
            res.delay_to_store_resources_in_clh_db_nanos = Primitive.to_proto(
                resource.delay_to_store_resources_in_clh_db_nanos
            )
        return res

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return None

        return InstanceReconcileRecipe(
            steps=InstanceReconcileRecipeStepsArray.from_proto(resource.steps),
            honor_cancel_request=Primitive.from_proto(resource.honor_cancel_request),
            ignore_recipe_after=Primitive.from_proto(resource.ignore_recipe_after),
            verify_deadline_seconds_below=Primitive.from_proto(
                resource.verify_deadline_seconds_below
            ),
            populate_operation_result=Primitive.from_proto(
                resource.populate_operation_result
            ),
            readonly_recipe_start_time=Primitive.from_proto(
                resource.readonly_recipe_start_time
            ),
            resource_names_stored_in_clh_with_delay=Primitive.from_proto(
                resource.resource_names_stored_in_clh_with_delay
            ),
            delay_to_store_resources_in_clh_db_nanos=Primitive.from_proto(
                resource.delay_to_store_resources_in_clh_db_nanos
            ),
        )


class InstanceReconcileRecipeArray(object):
    @classmethod
    def to_proto(self, resources):
        if not resources:
            return resources
        return [InstanceReconcileRecipe.to_proto(i) for i in resources]

    @classmethod
    def from_proto(self, resources):
        return [InstanceReconcileRecipe.from_proto(i) for i in resources]


class InstanceReconcileRecipeSteps(object):
    def __init__(
        self,
        relative_time: int = None,
        sleep_duration: int = None,
        action: str = None,
        status: dict = None,
        error_space: str = None,
        p4_service_account: str = None,
        resource_metadata_size: int = None,
        description: str = None,
        updated_repeat_operation_delay_sec: float = None,
        quota_request_deltas: list = None,
        preprocess_update: dict = None,
        public_operation_metadata: str = None,
        requested_tenant_project: dict = None,
        permissions_info: list = None,
        key_notifications_update: dict = None,
        clh_data_update_time: str = None,
        public_error_message: str = None,
    ):
        self.relative_time = relative_time
        self.sleep_duration = sleep_duration
        self.action = action
        self.status = status
        self.error_space = error_space
        self.p4_service_account = p4_service_account
        self.resource_metadata_size = resource_metadata_size
        self.description = description
        self.updated_repeat_operation_delay_sec = updated_repeat_operation_delay_sec
        self.quota_request_deltas = quota_request_deltas
        self.preprocess_update = preprocess_update
        self.public_operation_metadata = public_operation_metadata
        self.requested_tenant_project = requested_tenant_project
        self.permissions_info = permissions_info
        self.key_notifications_update = key_notifications_update
        self.clh_data_update_time = clh_data_update_time
        self.public_error_message = public_error_message

    @classmethod
    def to_proto(self, resource):
        if not resource:
            return None

        res = instance_pb2.Tier2AlphaInstanceReconcileRecipeSteps()
        if Primitive.to_proto(resource.relative_time):
            res.relative_time = Primitive.to_proto(resource.relative_time)
        if Primitive.to_proto(resource.sleep_duration):
            res.sleep_duration = Primitive.to_proto(resource.sleep_duration)
        if InstanceReconcileRecipeStepsActionEnum.to_proto(resource.action):
            res.action = InstanceReconcileRecipeStepsActionEnum.to_proto(
                resource.action
            )
        if InstanceReconcileRecipeStepsStatus.to_proto(resource.status):
            res.status.CopyFrom(
                InstanceReconcileRecipeStepsStatus.to_proto(resource.status)
            )
        else:
            res.ClearField("status")
        if Primitive.to_proto(resource.error_space):
            res.error_space = Primitive.to_proto(resource.error_space)
        if Primitive.to_proto(resource.p4_service_account):
            res.p4_service_account = Primitive.to_proto(resource.p4_service_account)
        if Primitive.to_proto(resource.resource_metadata_size):
            res.resource_metadata_size = Primitive.to_proto(
                resource.resource_metadata_size
            )
        if Primitive.to_proto(resource.description):
            res.description = Primitive.to_proto(resource.description)
        if Primitive.to_proto(resource.updated_repeat_operation_delay_sec):
            res.updated_repeat_operation_delay_sec = Primitive.to_proto(
                resource.updated_repeat_operation_delay_sec
            )
        if InstanceReconcileRecipeStepsQuotaRequestDeltasArray.to_proto(
            resource.quota_request_deltas
        ):
            res.quota_request_deltas.extend(
                InstanceReconcileRecipeStepsQuotaRequestDeltasArray.to_proto(
                    resource.quota_request_deltas
                )
            )
        if InstanceReconcileRecipeStepsPreprocessUpdate.to_proto(
            resource.preprocess_update
        ):
            res.preprocess_update.CopyFrom(
                InstanceReconcileRecipeStepsPreprocessUpdate.to_proto(
                    resource.preprocess_update
                )
            )
        else:
            res.ClearField("preprocess_update")
        if Primitive.to_proto(resource.public_operation_metadata):
            res.public_operation_metadata = Primitive.to_proto(
                resource.public_operation_metadata
            )
        if InstanceReconcileRecipeStepsRequestedTenantProject.to_proto(
            resource.requested_tenant_project
        ):
            res.requested_tenant_project.CopyFrom(
                InstanceReconcileRecipeStepsRequestedTenantProject.to_proto(
                    resource.requested_tenant_project
                )
            )
        else:
            res.ClearField("requested_tenant_project")
        if InstanceReconcileRecipeStepsPermissionsInfoArray.to_proto(
            resource.permissions_info
        ):
            res.permissions_info.extend(
                InstanceReconcileRecipeStepsPermissionsInfoArray.to_proto(
                    resource.permissions_info
                )
            )
        if InstanceReconcileRecipeStepsKeyNotificationsUpdate.to_proto(
            resource.key_notifications_update
        ):
            res.key_notifications_update.CopyFrom(
                InstanceReconcileRecipeStepsKeyNotificationsUpdate.to_proto(
                    resource.key_notifications_update
                )
            )
        else:
            res.ClearField("key_notifications_update")
        if Primitive.to_proto(resource.clh_data_update_time):
            res.clh_data_update_time = Primitive.to_proto(resource.clh_data_update_time)
        if Primitive.to_proto(resource.public_error_message):
            res.public_error_message = Primitive.to_proto(resource.public_error_message)
        return res

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return None

        return InstanceReconcileRecipeSteps(
            relative_time=Primitive.from_proto(resource.relative_time),
            sleep_duration=Primitive.from_proto(resource.sleep_duration),
            action=InstanceReconcileRecipeStepsActionEnum.from_proto(resource.action),
            status=InstanceReconcileRecipeStepsStatus.from_proto(resource.status),
            error_space=Primitive.from_proto(resource.error_space),
            p4_service_account=Primitive.from_proto(resource.p4_service_account),
            resource_metadata_size=Primitive.from_proto(
                resource.resource_metadata_size
            ),
            description=Primitive.from_proto(resource.description),
            updated_repeat_operation_delay_sec=Primitive.from_proto(
                resource.updated_repeat_operation_delay_sec
            ),
            quota_request_deltas=InstanceReconcileRecipeStepsQuotaRequestDeltasArray.from_proto(
                resource.quota_request_deltas
            ),
            preprocess_update=InstanceReconcileRecipeStepsPreprocessUpdate.from_proto(
                resource.preprocess_update
            ),
            public_operation_metadata=Primitive.from_proto(
                resource.public_operation_metadata
            ),
            requested_tenant_project=InstanceReconcileRecipeStepsRequestedTenantProject.from_proto(
                resource.requested_tenant_project
            ),
            permissions_info=InstanceReconcileRecipeStepsPermissionsInfoArray.from_proto(
                resource.permissions_info
            ),
            key_notifications_update=InstanceReconcileRecipeStepsKeyNotificationsUpdate.from_proto(
                resource.key_notifications_update
            ),
            clh_data_update_time=Primitive.from_proto(resource.clh_data_update_time),
            public_error_message=Primitive.from_proto(resource.public_error_message),
        )


class InstanceReconcileRecipeStepsArray(object):
    @classmethod
    def to_proto(self, resources):
        if not resources:
            return resources
        return [InstanceReconcileRecipeSteps.to_proto(i) for i in resources]

    @classmethod
    def from_proto(self, resources):
        return [InstanceReconcileRecipeSteps.from_proto(i) for i in resources]


class InstanceReconcileRecipeStepsStatus(object):
    def __init__(self, code: int = None, message: str = None, details: list = None):
        self.code = code
        self.message = message
        self.details = details

    @classmethod
    def to_proto(self, resource):
        if not resource:
            return None

        res = instance_pb2.Tier2AlphaInstanceReconcileRecipeStepsStatus()
        if Primitive.to_proto(resource.code):
            res.code = Primitive.to_proto(resource.code)
        if Primitive.to_proto(resource.message):
            res.message = Primitive.to_proto(resource.message)
        if InstanceReconcileRecipeStepsStatusDetailsArray.to_proto(resource.details):
            res.details.extend(
                InstanceReconcileRecipeStepsStatusDetailsArray.to_proto(
                    resource.details
                )
            )
        return res

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return None

        return InstanceReconcileRecipeStepsStatus(
            code=Primitive.from_proto(resource.code),
            message=Primitive.from_proto(resource.message),
            details=InstanceReconcileRecipeStepsStatusDetailsArray.from_proto(
                resource.details
            ),
        )


class InstanceReconcileRecipeStepsStatusArray(object):
    @classmethod
    def to_proto(self, resources):
        if not resources:
            return resources
        return [InstanceReconcileRecipeStepsStatus.to_proto(i) for i in resources]

    @classmethod
    def from_proto(self, resources):
        return [InstanceReconcileRecipeStepsStatus.from_proto(i) for i in resources]


class InstanceReconcileRecipeStepsStatusDetails(object):
    def __init__(self, type_url: str = None, value: str = None):
        self.type_url = type_url
        self.value = value

    @classmethod
    def to_proto(self, resource):
        if not resource:
            return None

        res = instance_pb2.Tier2AlphaInstanceReconcileRecipeStepsStatusDetails()
        if Primitive.to_proto(resource.type_url):
            res.type_url = Primitive.to_proto(resource.type_url)
        if Primitive.to_proto(resource.value):
            res.value = Primitive.to_proto(resource.value)
        return res

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return None

        return InstanceReconcileRecipeStepsStatusDetails(
            type_url=Primitive.from_proto(resource.type_url),
            value=Primitive.from_proto(resource.value),
        )


class InstanceReconcileRecipeStepsStatusDetailsArray(object):
    @classmethod
    def to_proto(self, resources):
        if not resources:
            return resources
        return [
            InstanceReconcileRecipeStepsStatusDetails.to_proto(i) for i in resources
        ]

    @classmethod
    def from_proto(self, resources):
        return [
            InstanceReconcileRecipeStepsStatusDetails.from_proto(i) for i in resources
        ]


class InstanceReconcileRecipeStepsQuotaRequestDeltas(object):
    def __init__(
        self,
        metric_name: str = None,
        amount: int = None,
        quota_location_name: str = None,
    ):
        self.metric_name = metric_name
        self.amount = amount
        self.quota_location_name = quota_location_name

    @classmethod
    def to_proto(self, resource):
        if not resource:
            return None

        res = instance_pb2.Tier2AlphaInstanceReconcileRecipeStepsQuotaRequestDeltas()
        if Primitive.to_proto(resource.metric_name):
            res.metric_name = Primitive.to_proto(resource.metric_name)
        if Primitive.to_proto(resource.amount):
            res.amount = Primitive.to_proto(resource.amount)
        if Primitive.to_proto(resource.quota_location_name):
            res.quota_location_name = Primitive.to_proto(resource.quota_location_name)
        return res

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return None

        return InstanceReconcileRecipeStepsQuotaRequestDeltas(
            metric_name=Primitive.from_proto(resource.metric_name),
            amount=Primitive.from_proto(resource.amount),
            quota_location_name=Primitive.from_proto(resource.quota_location_name),
        )


class InstanceReconcileRecipeStepsQuotaRequestDeltasArray(object):
    @classmethod
    def to_proto(self, resources):
        if not resources:
            return resources
        return [
            InstanceReconcileRecipeStepsQuotaRequestDeltas.to_proto(i)
            for i in resources
        ]

    @classmethod
    def from_proto(self, resources):
        return [
            InstanceReconcileRecipeStepsQuotaRequestDeltas.from_proto(i)
            for i in resources
        ]


class InstanceReconcileRecipeStepsPreprocessUpdate(object):
    def __init__(
        self, latency_slo_bucket_name: str = None, public_operation_metadata: str = None
    ):
        self.latency_slo_bucket_name = latency_slo_bucket_name
        self.public_operation_metadata = public_operation_metadata

    @classmethod
    def to_proto(self, resource):
        if not resource:
            return None

        res = instance_pb2.Tier2AlphaInstanceReconcileRecipeStepsPreprocessUpdate()
        if Primitive.to_proto(resource.latency_slo_bucket_name):
            res.latency_slo_bucket_name = Primitive.to_proto(
                resource.latency_slo_bucket_name
            )
        if Primitive.to_proto(resource.public_operation_metadata):
            res.public_operation_metadata = Primitive.to_proto(
                resource.public_operation_metadata
            )
        return res

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return None

        return InstanceReconcileRecipeStepsPreprocessUpdate(
            latency_slo_bucket_name=Primitive.from_proto(
                resource.latency_slo_bucket_name
            ),
            public_operation_metadata=Primitive.from_proto(
                resource.public_operation_metadata
            ),
        )


class InstanceReconcileRecipeStepsPreprocessUpdateArray(object):
    @classmethod
    def to_proto(self, resources):
        if not resources:
            return resources
        return [
            InstanceReconcileRecipeStepsPreprocessUpdate.to_proto(i) for i in resources
        ]

    @classmethod
    def from_proto(self, resources):
        return [
            InstanceReconcileRecipeStepsPreprocessUpdate.from_proto(i)
            for i in resources
        ]


class InstanceReconcileRecipeStepsRequestedTenantProject(object):
    def __init__(self, tag: str = None, folder: str = None, scope: str = None):
        self.tag = tag
        self.folder = folder
        self.scope = scope

    @classmethod
    def to_proto(self, resource):
        if not resource:
            return None

        res = (
            instance_pb2.Tier2AlphaInstanceReconcileRecipeStepsRequestedTenantProject()
        )
        if Primitive.to_proto(resource.tag):
            res.tag = Primitive.to_proto(resource.tag)
        if Primitive.to_proto(resource.folder):
            res.folder = Primitive.to_proto(resource.folder)
        if InstanceReconcileRecipeStepsRequestedTenantProjectScopeEnum.to_proto(
            resource.scope
        ):
            res.scope = InstanceReconcileRecipeStepsRequestedTenantProjectScopeEnum.to_proto(
                resource.scope
            )
        return res

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return None

        return InstanceReconcileRecipeStepsRequestedTenantProject(
            tag=Primitive.from_proto(resource.tag),
            folder=Primitive.from_proto(resource.folder),
            scope=InstanceReconcileRecipeStepsRequestedTenantProjectScopeEnum.from_proto(
                resource.scope
            ),
        )


class InstanceReconcileRecipeStepsRequestedTenantProjectArray(object):
    @classmethod
    def to_proto(self, resources):
        if not resources:
            return resources
        return [
            InstanceReconcileRecipeStepsRequestedTenantProject.to_proto(i)
            for i in resources
        ]

    @classmethod
    def from_proto(self, resources):
        return [
            InstanceReconcileRecipeStepsRequestedTenantProject.from_proto(i)
            for i in resources
        ]


class InstanceReconcileRecipeStepsPermissionsInfo(object):
    def __init__(
        self,
        policy_name: dict = None,
        iam_permissions: list = None,
        resource_path: str = None,
        api_attrs: dict = None,
        policy_name_mode: str = None,
        resource: dict = None,
    ):
        self.policy_name = policy_name
        self.iam_permissions = iam_permissions
        self.resource_path = resource_path
        self.api_attrs = api_attrs
        self.policy_name_mode = policy_name_mode
        self.resource = resource

    @classmethod
    def to_proto(self, resource):
        if not resource:
            return None

        res = instance_pb2.Tier2AlphaInstanceReconcileRecipeStepsPermissionsInfo()
        if InstanceReconcileRecipeStepsPermissionsInfoPolicyName.to_proto(
            resource.policy_name
        ):
            res.policy_name.CopyFrom(
                InstanceReconcileRecipeStepsPermissionsInfoPolicyName.to_proto(
                    resource.policy_name
                )
            )
        else:
            res.ClearField("policy_name")
        if InstanceReconcileRecipeStepsPermissionsInfoIamPermissionsArray.to_proto(
            resource.iam_permissions
        ):
            res.iam_permissions.extend(
                InstanceReconcileRecipeStepsPermissionsInfoIamPermissionsArray.to_proto(
                    resource.iam_permissions
                )
            )
        if Primitive.to_proto(resource.resource_path):
            res.resource_path = Primitive.to_proto(resource.resource_path)
        if InstanceGoogleprotobufstruct.to_proto(resource.api_attrs):
            res.api_attrs.CopyFrom(
                InstanceGoogleprotobufstruct.to_proto(resource.api_attrs)
            )
        else:
            res.ClearField("api_attrs")
        if InstanceReconcileRecipeStepsPermissionsInfoPolicyNameModeEnum.to_proto(
            resource.policy_name_mode
        ):
            res.policy_name_mode = InstanceReconcileRecipeStepsPermissionsInfoPolicyNameModeEnum.to_proto(
                resource.policy_name_mode
            )
        if InstanceReconcileRecipeStepsPermissionsInfoResource.to_proto(
            resource.resource
        ):
            res.resource.CopyFrom(
                InstanceReconcileRecipeStepsPermissionsInfoResource.to_proto(
                    resource.resource
                )
            )
        else:
            res.ClearField("resource")
        return res

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return None

        return InstanceReconcileRecipeStepsPermissionsInfo(
            policy_name=InstanceReconcileRecipeStepsPermissionsInfoPolicyName.from_proto(
                resource.policy_name
            ),
            iam_permissions=InstanceReconcileRecipeStepsPermissionsInfoIamPermissionsArray.from_proto(
                resource.iam_permissions
            ),
            resource_path=Primitive.from_proto(resource.resource_path),
            api_attrs=InstanceGoogleprotobufstruct.from_proto(resource.api_attrs),
            policy_name_mode=InstanceReconcileRecipeStepsPermissionsInfoPolicyNameModeEnum.from_proto(
                resource.policy_name_mode
            ),
            resource=InstanceReconcileRecipeStepsPermissionsInfoResource.from_proto(
                resource.resource
            ),
        )


class InstanceReconcileRecipeStepsPermissionsInfoArray(object):
    @classmethod
    def to_proto(self, resources):
        if not resources:
            return resources
        return [
            InstanceReconcileRecipeStepsPermissionsInfo.to_proto(i) for i in resources
        ]

    @classmethod
    def from_proto(self, resources):
        return [
            InstanceReconcileRecipeStepsPermissionsInfo.from_proto(i) for i in resources
        ]


class InstanceReconcileRecipeStepsPermissionsInfoPolicyName(object):
    def __init__(self, type: str = None, id: str = None, region: str = None):
        self.type = type
        self.id = id
        self.region = region

    @classmethod
    def to_proto(self, resource):
        if not resource:
            return None

        res = (
            instance_pb2.Tier2AlphaInstanceReconcileRecipeStepsPermissionsInfoPolicyName()
        )
        if Primitive.to_proto(resource.type):
            res.type = Primitive.to_proto(resource.type)
        if Primitive.to_proto(resource.id):
            res.id = Primitive.to_proto(resource.id)
        if Primitive.to_proto(resource.region):
            res.region = Primitive.to_proto(resource.region)
        return res

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return None

        return InstanceReconcileRecipeStepsPermissionsInfoPolicyName(
            type=Primitive.from_proto(resource.type),
            id=Primitive.from_proto(resource.id),
            region=Primitive.from_proto(resource.region),
        )


class InstanceReconcileRecipeStepsPermissionsInfoPolicyNameArray(object):
    @classmethod
    def to_proto(self, resources):
        if not resources:
            return resources
        return [
            InstanceReconcileRecipeStepsPermissionsInfoPolicyName.to_proto(i)
            for i in resources
        ]

    @classmethod
    def from_proto(self, resources):
        return [
            InstanceReconcileRecipeStepsPermissionsInfoPolicyName.from_proto(i)
            for i in resources
        ]


class InstanceReconcileRecipeStepsPermissionsInfoIamPermissions(object):
    def __init__(self, permission: str = None):
        self.permission = permission

    @classmethod
    def to_proto(self, resource):
        if not resource:
            return None

        res = (
            instance_pb2.Tier2AlphaInstanceReconcileRecipeStepsPermissionsInfoIamPermissions()
        )
        if Primitive.to_proto(resource.permission):
            res.permission = Primitive.to_proto(resource.permission)
        return res

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return None

        return InstanceReconcileRecipeStepsPermissionsInfoIamPermissions(
            permission=Primitive.from_proto(resource.permission),
        )


class InstanceReconcileRecipeStepsPermissionsInfoIamPermissionsArray(object):
    @classmethod
    def to_proto(self, resources):
        if not resources:
            return resources
        return [
            InstanceReconcileRecipeStepsPermissionsInfoIamPermissions.to_proto(i)
            for i in resources
        ]

    @classmethod
    def from_proto(self, resources):
        return [
            InstanceReconcileRecipeStepsPermissionsInfoIamPermissions.from_proto(i)
            for i in resources
        ]


class InstanceReconcileRecipeStepsPermissionsInfoResource(object):
    def __init__(
        self,
        name: str = None,
        type: str = None,
        service: str = None,
        labels: dict = None,
    ):
        self.name = name
        self.type = type
        self.service = service
        self.labels = labels

    @classmethod
    def to_proto(self, resource):
        if not resource:
            return None

        res = (
            instance_pb2.Tier2AlphaInstanceReconcileRecipeStepsPermissionsInfoResource()
        )
        if Primitive.to_proto(resource.name):
            res.name = Primitive.to_proto(resource.name)
        if Primitive.to_proto(resource.type):
            res.type = Primitive.to_proto(resource.type)
        if Primitive.to_proto(resource.service):
            res.service = Primitive.to_proto(resource.service)
        if Primitive.to_proto(resource.labels):
            res.labels = Primitive.to_proto(resource.labels)
        return res

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return None

        return InstanceReconcileRecipeStepsPermissionsInfoResource(
            name=Primitive.from_proto(resource.name),
            type=Primitive.from_proto(resource.type),
            service=Primitive.from_proto(resource.service),
            labels=Primitive.from_proto(resource.labels),
        )


class InstanceReconcileRecipeStepsPermissionsInfoResourceArray(object):
    @classmethod
    def to_proto(self, resources):
        if not resources:
            return resources
        return [
            InstanceReconcileRecipeStepsPermissionsInfoResource.to_proto(i)
            for i in resources
        ]

    @classmethod
    def from_proto(self, resources):
        return [
            InstanceReconcileRecipeStepsPermissionsInfoResource.from_proto(i)
            for i in resources
        ]


class InstanceReconcileRecipeStepsKeyNotificationsUpdate(object):
    def __init__(self, key_notifications_info: dict = None):
        self.key_notifications_info = key_notifications_info

    @classmethod
    def to_proto(self, resource):
        if not resource:
            return None

        res = (
            instance_pb2.Tier2AlphaInstanceReconcileRecipeStepsKeyNotificationsUpdate()
        )
        if InstanceReconcileRecipeStepsKeyNotificationsUpdateKeyNotificationsInfo.to_proto(
            resource.key_notifications_info
        ):
            res.key_notifications_info.CopyFrom(
                InstanceReconcileRecipeStepsKeyNotificationsUpdateKeyNotificationsInfo.to_proto(
                    resource.key_notifications_info
                )
            )
        else:
            res.ClearField("key_notifications_info")
        return res

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return None

        return InstanceReconcileRecipeStepsKeyNotificationsUpdate(
            key_notifications_info=InstanceReconcileRecipeStepsKeyNotificationsUpdateKeyNotificationsInfo.from_proto(
                resource.key_notifications_info
            ),
        )


class InstanceReconcileRecipeStepsKeyNotificationsUpdateArray(object):
    @classmethod
    def to_proto(self, resources):
        if not resources:
            return resources
        return [
            InstanceReconcileRecipeStepsKeyNotificationsUpdate.to_proto(i)
            for i in resources
        ]

    @classmethod
    def from_proto(self, resources):
        return [
            InstanceReconcileRecipeStepsKeyNotificationsUpdate.from_proto(i)
            for i in resources
        ]


class InstanceReconcileRecipeStepsKeyNotificationsUpdateKeyNotificationsInfo(object):
    def __init__(
        self,
        data_version: int = None,
        delegate: str = None,
        key_notification_configs: list = None,
    ):
        self.data_version = data_version
        self.delegate = delegate
        self.key_notification_configs = key_notification_configs

    @classmethod
    def to_proto(self, resource):
        if not resource:
            return None

        res = (
            instance_pb2.Tier2AlphaInstanceReconcileRecipeStepsKeyNotificationsUpdateKeyNotificationsInfo()
        )
        if Primitive.to_proto(resource.data_version):
            res.data_version = Primitive.to_proto(resource.data_version)
        if Primitive.to_proto(resource.delegate):
            res.delegate = Primitive.to_proto(resource.delegate)
        if InstanceReconcileRecipeStepsKeyNotificationsUpdateKeyNotificationsInfoKeyNotificationConfigsArray.to_proto(
            resource.key_notification_configs
        ):
            res.key_notification_configs.extend(
                InstanceReconcileRecipeStepsKeyNotificationsUpdateKeyNotificationsInfoKeyNotificationConfigsArray.to_proto(
                    resource.key_notification_configs
                )
            )
        return res

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return None

        return InstanceReconcileRecipeStepsKeyNotificationsUpdateKeyNotificationsInfo(
            data_version=Primitive.from_proto(resource.data_version),
            delegate=Primitive.from_proto(resource.delegate),
            key_notification_configs=InstanceReconcileRecipeStepsKeyNotificationsUpdateKeyNotificationsInfoKeyNotificationConfigsArray.from_proto(
                resource.key_notification_configs
            ),
        )


class InstanceReconcileRecipeStepsKeyNotificationsUpdateKeyNotificationsInfoArray(
    object
):
    @classmethod
    def to_proto(self, resources):
        if not resources:
            return resources
        return [
            InstanceReconcileRecipeStepsKeyNotificationsUpdateKeyNotificationsInfo.to_proto(
                i
            )
            for i in resources
        ]

    @classmethod
    def from_proto(self, resources):
        return [
            InstanceReconcileRecipeStepsKeyNotificationsUpdateKeyNotificationsInfo.from_proto(
                i
            )
            for i in resources
        ]


class InstanceReconcileRecipeStepsKeyNotificationsUpdateKeyNotificationsInfoKeyNotificationConfigs(
    object
):
    def __init__(
        self,
        key_or_version_name: str = None,
        grant: str = None,
        delegator_gaia_id: int = None,
    ):
        self.key_or_version_name = key_or_version_name
        self.grant = grant
        self.delegator_gaia_id = delegator_gaia_id

    @classmethod
    def to_proto(self, resource):
        if not resource:
            return None

        res = (
            instance_pb2.Tier2AlphaInstanceReconcileRecipeStepsKeyNotificationsUpdateKeyNotificationsInfoKeyNotificationConfigs()
        )
        if Primitive.to_proto(resource.key_or_version_name):
            res.key_or_version_name = Primitive.to_proto(resource.key_or_version_name)
        if Primitive.to_proto(resource.grant):
            res.grant = Primitive.to_proto(resource.grant)
        if Primitive.to_proto(resource.delegator_gaia_id):
            res.delegator_gaia_id = Primitive.to_proto(resource.delegator_gaia_id)
        return res

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return None

        return InstanceReconcileRecipeStepsKeyNotificationsUpdateKeyNotificationsInfoKeyNotificationConfigs(
            key_or_version_name=Primitive.from_proto(resource.key_or_version_name),
            grant=Primitive.from_proto(resource.grant),
            delegator_gaia_id=Primitive.from_proto(resource.delegator_gaia_id),
        )


class InstanceReconcileRecipeStepsKeyNotificationsUpdateKeyNotificationsInfoKeyNotificationConfigsArray(
    object
):
    @classmethod
    def to_proto(self, resources):
        if not resources:
            return resources
        return [
            InstanceReconcileRecipeStepsKeyNotificationsUpdateKeyNotificationsInfoKeyNotificationConfigs.to_proto(
                i
            )
            for i in resources
        ]

    @classmethod
    def from_proto(self, resources):
        return [
            InstanceReconcileRecipeStepsKeyNotificationsUpdateKeyNotificationsInfoKeyNotificationConfigs.from_proto(
                i
            )
            for i in resources
        ]


class InstancePreprocessPassthroughRecipe(object):
    def __init__(
        self,
        steps: list = None,
        honor_cancel_request: bool = None,
        ignore_recipe_after: int = None,
        verify_deadline_seconds_below: float = None,
        populate_operation_result: bool = None,
        readonly_recipe_start_time: str = None,
        resource_names_stored_in_clh_with_delay: list = None,
        delay_to_store_resources_in_clh_db_nanos: int = None,
    ):
        self.steps = steps
        self.honor_cancel_request = honor_cancel_request
        self.ignore_recipe_after = ignore_recipe_after
        self.verify_deadline_seconds_below = verify_deadline_seconds_below
        self.populate_operation_result = populate_operation_result
        self.readonly_recipe_start_time = readonly_recipe_start_time
        self.resource_names_stored_in_clh_with_delay = (
            resource_names_stored_in_clh_with_delay
        )
        self.delay_to_store_resources_in_clh_db_nanos = (
            delay_to_store_resources_in_clh_db_nanos
        )

    @classmethod
    def to_proto(self, resource):
        if not resource:
            return None

        res = instance_pb2.Tier2AlphaInstancePreprocessPassthroughRecipe()
        if InstancePreprocessPassthroughRecipeStepsArray.to_proto(resource.steps):
            res.steps.extend(
                InstancePreprocessPassthroughRecipeStepsArray.to_proto(resource.steps)
            )
        if Primitive.to_proto(resource.honor_cancel_request):
            res.honor_cancel_request = Primitive.to_proto(resource.honor_cancel_request)
        if Primitive.to_proto(resource.ignore_recipe_after):
            res.ignore_recipe_after = Primitive.to_proto(resource.ignore_recipe_after)
        if Primitive.to_proto(resource.verify_deadline_seconds_below):
            res.verify_deadline_seconds_below = Primitive.to_proto(
                resource.verify_deadline_seconds_below
            )
        if Primitive.to_proto(resource.populate_operation_result):
            res.populate_operation_result = Primitive.to_proto(
                resource.populate_operation_result
            )
        if Primitive.to_proto(resource.readonly_recipe_start_time):
            res.readonly_recipe_start_time = Primitive.to_proto(
                resource.readonly_recipe_start_time
            )
        if Primitive.to_proto(resource.resource_names_stored_in_clh_with_delay):
            res.resource_names_stored_in_clh_with_delay.extend(
                Primitive.to_proto(resource.resource_names_stored_in_clh_with_delay)
            )
        if Primitive.to_proto(resource.delay_to_store_resources_in_clh_db_nanos):
            res.delay_to_store_resources_in_clh_db_nanos = Primitive.to_proto(
                resource.delay_to_store_resources_in_clh_db_nanos
            )
        return res

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return None

        return InstancePreprocessPassthroughRecipe(
            steps=InstancePreprocessPassthroughRecipeStepsArray.from_proto(
                resource.steps
            ),
            honor_cancel_request=Primitive.from_proto(resource.honor_cancel_request),
            ignore_recipe_after=Primitive.from_proto(resource.ignore_recipe_after),
            verify_deadline_seconds_below=Primitive.from_proto(
                resource.verify_deadline_seconds_below
            ),
            populate_operation_result=Primitive.from_proto(
                resource.populate_operation_result
            ),
            readonly_recipe_start_time=Primitive.from_proto(
                resource.readonly_recipe_start_time
            ),
            resource_names_stored_in_clh_with_delay=Primitive.from_proto(
                resource.resource_names_stored_in_clh_with_delay
            ),
            delay_to_store_resources_in_clh_db_nanos=Primitive.from_proto(
                resource.delay_to_store_resources_in_clh_db_nanos
            ),
        )


class InstancePreprocessPassthroughRecipeArray(object):
    @classmethod
    def to_proto(self, resources):
        if not resources:
            return resources
        return [InstancePreprocessPassthroughRecipe.to_proto(i) for i in resources]

    @classmethod
    def from_proto(self, resources):
        return [InstancePreprocessPassthroughRecipe.from_proto(i) for i in resources]


class InstancePreprocessPassthroughRecipeSteps(object):
    def __init__(
        self,
        relative_time: int = None,
        sleep_duration: int = None,
        action: str = None,
        status: dict = None,
        error_space: str = None,
        p4_service_account: str = None,
        resource_metadata_size: int = None,
        description: str = None,
        updated_repeat_operation_delay_sec: float = None,
        quota_request_deltas: list = None,
        preprocess_update: dict = None,
        public_operation_metadata: str = None,
        requested_tenant_project: dict = None,
        permissions_info: list = None,
        key_notifications_update: dict = None,
        clh_data_update_time: str = None,
        public_error_message: str = None,
    ):
        self.relative_time = relative_time
        self.sleep_duration = sleep_duration
        self.action = action
        self.status = status
        self.error_space = error_space
        self.p4_service_account = p4_service_account
        self.resource_metadata_size = resource_metadata_size
        self.description = description
        self.updated_repeat_operation_delay_sec = updated_repeat_operation_delay_sec
        self.quota_request_deltas = quota_request_deltas
        self.preprocess_update = preprocess_update
        self.public_operation_metadata = public_operation_metadata
        self.requested_tenant_project = requested_tenant_project
        self.permissions_info = permissions_info
        self.key_notifications_update = key_notifications_update
        self.clh_data_update_time = clh_data_update_time
        self.public_error_message = public_error_message

    @classmethod
    def to_proto(self, resource):
        if not resource:
            return None

        res = instance_pb2.Tier2AlphaInstancePreprocessPassthroughRecipeSteps()
        if Primitive.to_proto(resource.relative_time):
            res.relative_time = Primitive.to_proto(resource.relative_time)
        if Primitive.to_proto(resource.sleep_duration):
            res.sleep_duration = Primitive.to_proto(resource.sleep_duration)
        if InstancePreprocessPassthroughRecipeStepsActionEnum.to_proto(resource.action):
            res.action = InstancePreprocessPassthroughRecipeStepsActionEnum.to_proto(
                resource.action
            )
        if InstancePreprocessPassthroughRecipeStepsStatus.to_proto(resource.status):
            res.status.CopyFrom(
                InstancePreprocessPassthroughRecipeStepsStatus.to_proto(resource.status)
            )
        else:
            res.ClearField("status")
        if Primitive.to_proto(resource.error_space):
            res.error_space = Primitive.to_proto(resource.error_space)
        if Primitive.to_proto(resource.p4_service_account):
            res.p4_service_account = Primitive.to_proto(resource.p4_service_account)
        if Primitive.to_proto(resource.resource_metadata_size):
            res.resource_metadata_size = Primitive.to_proto(
                resource.resource_metadata_size
            )
        if Primitive.to_proto(resource.description):
            res.description = Primitive.to_proto(resource.description)
        if Primitive.to_proto(resource.updated_repeat_operation_delay_sec):
            res.updated_repeat_operation_delay_sec = Primitive.to_proto(
                resource.updated_repeat_operation_delay_sec
            )
        if InstancePreprocessPassthroughRecipeStepsQuotaRequestDeltasArray.to_proto(
            resource.quota_request_deltas
        ):
            res.quota_request_deltas.extend(
                InstancePreprocessPassthroughRecipeStepsQuotaRequestDeltasArray.to_proto(
                    resource.quota_request_deltas
                )
            )
        if InstancePreprocessPassthroughRecipeStepsPreprocessUpdate.to_proto(
            resource.preprocess_update
        ):
            res.preprocess_update.CopyFrom(
                InstancePreprocessPassthroughRecipeStepsPreprocessUpdate.to_proto(
                    resource.preprocess_update
                )
            )
        else:
            res.ClearField("preprocess_update")
        if Primitive.to_proto(resource.public_operation_metadata):
            res.public_operation_metadata = Primitive.to_proto(
                resource.public_operation_metadata
            )
        if InstancePreprocessPassthroughRecipeStepsRequestedTenantProject.to_proto(
            resource.requested_tenant_project
        ):
            res.requested_tenant_project.CopyFrom(
                InstancePreprocessPassthroughRecipeStepsRequestedTenantProject.to_proto(
                    resource.requested_tenant_project
                )
            )
        else:
            res.ClearField("requested_tenant_project")
        if InstancePreprocessPassthroughRecipeStepsPermissionsInfoArray.to_proto(
            resource.permissions_info
        ):
            res.permissions_info.extend(
                InstancePreprocessPassthroughRecipeStepsPermissionsInfoArray.to_proto(
                    resource.permissions_info
                )
            )
        if InstancePreprocessPassthroughRecipeStepsKeyNotificationsUpdate.to_proto(
            resource.key_notifications_update
        ):
            res.key_notifications_update.CopyFrom(
                InstancePreprocessPassthroughRecipeStepsKeyNotificationsUpdate.to_proto(
                    resource.key_notifications_update
                )
            )
        else:
            res.ClearField("key_notifications_update")
        if Primitive.to_proto(resource.clh_data_update_time):
            res.clh_data_update_time = Primitive.to_proto(resource.clh_data_update_time)
        if Primitive.to_proto(resource.public_error_message):
            res.public_error_message = Primitive.to_proto(resource.public_error_message)
        return res

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return None

        return InstancePreprocessPassthroughRecipeSteps(
            relative_time=Primitive.from_proto(resource.relative_time),
            sleep_duration=Primitive.from_proto(resource.sleep_duration),
            action=InstancePreprocessPassthroughRecipeStepsActionEnum.from_proto(
                resource.action
            ),
            status=InstancePreprocessPassthroughRecipeStepsStatus.from_proto(
                resource.status
            ),
            error_space=Primitive.from_proto(resource.error_space),
            p4_service_account=Primitive.from_proto(resource.p4_service_account),
            resource_metadata_size=Primitive.from_proto(
                resource.resource_metadata_size
            ),
            description=Primitive.from_proto(resource.description),
            updated_repeat_operation_delay_sec=Primitive.from_proto(
                resource.updated_repeat_operation_delay_sec
            ),
            quota_request_deltas=InstancePreprocessPassthroughRecipeStepsQuotaRequestDeltasArray.from_proto(
                resource.quota_request_deltas
            ),
            preprocess_update=InstancePreprocessPassthroughRecipeStepsPreprocessUpdate.from_proto(
                resource.preprocess_update
            ),
            public_operation_metadata=Primitive.from_proto(
                resource.public_operation_metadata
            ),
            requested_tenant_project=InstancePreprocessPassthroughRecipeStepsRequestedTenantProject.from_proto(
                resource.requested_tenant_project
            ),
            permissions_info=InstancePreprocessPassthroughRecipeStepsPermissionsInfoArray.from_proto(
                resource.permissions_info
            ),
            key_notifications_update=InstancePreprocessPassthroughRecipeStepsKeyNotificationsUpdate.from_proto(
                resource.key_notifications_update
            ),
            clh_data_update_time=Primitive.from_proto(resource.clh_data_update_time),
            public_error_message=Primitive.from_proto(resource.public_error_message),
        )


class InstancePreprocessPassthroughRecipeStepsArray(object):
    @classmethod
    def to_proto(self, resources):
        if not resources:
            return resources
        return [InstancePreprocessPassthroughRecipeSteps.to_proto(i) for i in resources]

    @classmethod
    def from_proto(self, resources):
        return [
            InstancePreprocessPassthroughRecipeSteps.from_proto(i) for i in resources
        ]


class InstancePreprocessPassthroughRecipeStepsStatus(object):
    def __init__(self, code: int = None, message: str = None, details: list = None):
        self.code = code
        self.message = message
        self.details = details

    @classmethod
    def to_proto(self, resource):
        if not resource:
            return None

        res = instance_pb2.Tier2AlphaInstancePreprocessPassthroughRecipeStepsStatus()
        if Primitive.to_proto(resource.code):
            res.code = Primitive.to_proto(resource.code)
        if Primitive.to_proto(resource.message):
            res.message = Primitive.to_proto(resource.message)
        if InstancePreprocessPassthroughRecipeStepsStatusDetailsArray.to_proto(
            resource.details
        ):
            res.details.extend(
                InstancePreprocessPassthroughRecipeStepsStatusDetailsArray.to_proto(
                    resource.details
                )
            )
        return res

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return None

        return InstancePreprocessPassthroughRecipeStepsStatus(
            code=Primitive.from_proto(resource.code),
            message=Primitive.from_proto(resource.message),
            details=InstancePreprocessPassthroughRecipeStepsStatusDetailsArray.from_proto(
                resource.details
            ),
        )


class InstancePreprocessPassthroughRecipeStepsStatusArray(object):
    @classmethod
    def to_proto(self, resources):
        if not resources:
            return resources
        return [
            InstancePreprocessPassthroughRecipeStepsStatus.to_proto(i)
            for i in resources
        ]

    @classmethod
    def from_proto(self, resources):
        return [
            InstancePreprocessPassthroughRecipeStepsStatus.from_proto(i)
            for i in resources
        ]


class InstancePreprocessPassthroughRecipeStepsStatusDetails(object):
    def __init__(self, type_url: str = None, value: str = None):
        self.type_url = type_url
        self.value = value

    @classmethod
    def to_proto(self, resource):
        if not resource:
            return None

        res = (
            instance_pb2.Tier2AlphaInstancePreprocessPassthroughRecipeStepsStatusDetails()
        )
        if Primitive.to_proto(resource.type_url):
            res.type_url = Primitive.to_proto(resource.type_url)
        if Primitive.to_proto(resource.value):
            res.value = Primitive.to_proto(resource.value)
        return res

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return None

        return InstancePreprocessPassthroughRecipeStepsStatusDetails(
            type_url=Primitive.from_proto(resource.type_url),
            value=Primitive.from_proto(resource.value),
        )


class InstancePreprocessPassthroughRecipeStepsStatusDetailsArray(object):
    @classmethod
    def to_proto(self, resources):
        if not resources:
            return resources
        return [
            InstancePreprocessPassthroughRecipeStepsStatusDetails.to_proto(i)
            for i in resources
        ]

    @classmethod
    def from_proto(self, resources):
        return [
            InstancePreprocessPassthroughRecipeStepsStatusDetails.from_proto(i)
            for i in resources
        ]


class InstancePreprocessPassthroughRecipeStepsQuotaRequestDeltas(object):
    def __init__(
        self,
        metric_name: str = None,
        amount: int = None,
        quota_location_name: str = None,
    ):
        self.metric_name = metric_name
        self.amount = amount
        self.quota_location_name = quota_location_name

    @classmethod
    def to_proto(self, resource):
        if not resource:
            return None

        res = (
            instance_pb2.Tier2AlphaInstancePreprocessPassthroughRecipeStepsQuotaRequestDeltas()
        )
        if Primitive.to_proto(resource.metric_name):
            res.metric_name = Primitive.to_proto(resource.metric_name)
        if Primitive.to_proto(resource.amount):
            res.amount = Primitive.to_proto(resource.amount)
        if Primitive.to_proto(resource.quota_location_name):
            res.quota_location_name = Primitive.to_proto(resource.quota_location_name)
        return res

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return None

        return InstancePreprocessPassthroughRecipeStepsQuotaRequestDeltas(
            metric_name=Primitive.from_proto(resource.metric_name),
            amount=Primitive.from_proto(resource.amount),
            quota_location_name=Primitive.from_proto(resource.quota_location_name),
        )


class InstancePreprocessPassthroughRecipeStepsQuotaRequestDeltasArray(object):
    @classmethod
    def to_proto(self, resources):
        if not resources:
            return resources
        return [
            InstancePreprocessPassthroughRecipeStepsQuotaRequestDeltas.to_proto(i)
            for i in resources
        ]

    @classmethod
    def from_proto(self, resources):
        return [
            InstancePreprocessPassthroughRecipeStepsQuotaRequestDeltas.from_proto(i)
            for i in resources
        ]


class InstancePreprocessPassthroughRecipeStepsPreprocessUpdate(object):
    def __init__(
        self, latency_slo_bucket_name: str = None, public_operation_metadata: str = None
    ):
        self.latency_slo_bucket_name = latency_slo_bucket_name
        self.public_operation_metadata = public_operation_metadata

    @classmethod
    def to_proto(self, resource):
        if not resource:
            return None

        res = (
            instance_pb2.Tier2AlphaInstancePreprocessPassthroughRecipeStepsPreprocessUpdate()
        )
        if Primitive.to_proto(resource.latency_slo_bucket_name):
            res.latency_slo_bucket_name = Primitive.to_proto(
                resource.latency_slo_bucket_name
            )
        if Primitive.to_proto(resource.public_operation_metadata):
            res.public_operation_metadata = Primitive.to_proto(
                resource.public_operation_metadata
            )
        return res

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return None

        return InstancePreprocessPassthroughRecipeStepsPreprocessUpdate(
            latency_slo_bucket_name=Primitive.from_proto(
                resource.latency_slo_bucket_name
            ),
            public_operation_metadata=Primitive.from_proto(
                resource.public_operation_metadata
            ),
        )


class InstancePreprocessPassthroughRecipeStepsPreprocessUpdateArray(object):
    @classmethod
    def to_proto(self, resources):
        if not resources:
            return resources
        return [
            InstancePreprocessPassthroughRecipeStepsPreprocessUpdate.to_proto(i)
            for i in resources
        ]

    @classmethod
    def from_proto(self, resources):
        return [
            InstancePreprocessPassthroughRecipeStepsPreprocessUpdate.from_proto(i)
            for i in resources
        ]


class InstancePreprocessPassthroughRecipeStepsRequestedTenantProject(object):
    def __init__(self, tag: str = None, folder: str = None, scope: str = None):
        self.tag = tag
        self.folder = folder
        self.scope = scope

    @classmethod
    def to_proto(self, resource):
        if not resource:
            return None

        res = (
            instance_pb2.Tier2AlphaInstancePreprocessPassthroughRecipeStepsRequestedTenantProject()
        )
        if Primitive.to_proto(resource.tag):
            res.tag = Primitive.to_proto(resource.tag)
        if Primitive.to_proto(resource.folder):
            res.folder = Primitive.to_proto(resource.folder)
        if InstancePreprocessPassthroughRecipeStepsRequestedTenantProjectScopeEnum.to_proto(
            resource.scope
        ):
            res.scope = InstancePreprocessPassthroughRecipeStepsRequestedTenantProjectScopeEnum.to_proto(
                resource.scope
            )
        return res

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return None

        return InstancePreprocessPassthroughRecipeStepsRequestedTenantProject(
            tag=Primitive.from_proto(resource.tag),
            folder=Primitive.from_proto(resource.folder),
            scope=InstancePreprocessPassthroughRecipeStepsRequestedTenantProjectScopeEnum.from_proto(
                resource.scope
            ),
        )


class InstancePreprocessPassthroughRecipeStepsRequestedTenantProjectArray(object):
    @classmethod
    def to_proto(self, resources):
        if not resources:
            return resources
        return [
            InstancePreprocessPassthroughRecipeStepsRequestedTenantProject.to_proto(i)
            for i in resources
        ]

    @classmethod
    def from_proto(self, resources):
        return [
            InstancePreprocessPassthroughRecipeStepsRequestedTenantProject.from_proto(i)
            for i in resources
        ]


class InstancePreprocessPassthroughRecipeStepsPermissionsInfo(object):
    def __init__(
        self,
        policy_name: dict = None,
        iam_permissions: list = None,
        resource_path: str = None,
        api_attrs: dict = None,
        policy_name_mode: str = None,
        resource: dict = None,
    ):
        self.policy_name = policy_name
        self.iam_permissions = iam_permissions
        self.resource_path = resource_path
        self.api_attrs = api_attrs
        self.policy_name_mode = policy_name_mode
        self.resource = resource

    @classmethod
    def to_proto(self, resource):
        if not resource:
            return None

        res = (
            instance_pb2.Tier2AlphaInstancePreprocessPassthroughRecipeStepsPermissionsInfo()
        )
        if InstancePreprocessPassthroughRecipeStepsPermissionsInfoPolicyName.to_proto(
            resource.policy_name
        ):
            res.policy_name.CopyFrom(
                InstancePreprocessPassthroughRecipeStepsPermissionsInfoPolicyName.to_proto(
                    resource.policy_name
                )
            )
        else:
            res.ClearField("policy_name")
        if InstancePreprocessPassthroughRecipeStepsPermissionsInfoIamPermissionsArray.to_proto(
            resource.iam_permissions
        ):
            res.iam_permissions.extend(
                InstancePreprocessPassthroughRecipeStepsPermissionsInfoIamPermissionsArray.to_proto(
                    resource.iam_permissions
                )
            )
        if Primitive.to_proto(resource.resource_path):
            res.resource_path = Primitive.to_proto(resource.resource_path)
        if InstanceGoogleprotobufstruct.to_proto(resource.api_attrs):
            res.api_attrs.CopyFrom(
                InstanceGoogleprotobufstruct.to_proto(resource.api_attrs)
            )
        else:
            res.ClearField("api_attrs")
        if InstancePreprocessPassthroughRecipeStepsPermissionsInfoPolicyNameModeEnum.to_proto(
            resource.policy_name_mode
        ):
            res.policy_name_mode = InstancePreprocessPassthroughRecipeStepsPermissionsInfoPolicyNameModeEnum.to_proto(
                resource.policy_name_mode
            )
        if InstancePreprocessPassthroughRecipeStepsPermissionsInfoResource.to_proto(
            resource.resource
        ):
            res.resource.CopyFrom(
                InstancePreprocessPassthroughRecipeStepsPermissionsInfoResource.to_proto(
                    resource.resource
                )
            )
        else:
            res.ClearField("resource")
        return res

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return None

        return InstancePreprocessPassthroughRecipeStepsPermissionsInfo(
            policy_name=InstancePreprocessPassthroughRecipeStepsPermissionsInfoPolicyName.from_proto(
                resource.policy_name
            ),
            iam_permissions=InstancePreprocessPassthroughRecipeStepsPermissionsInfoIamPermissionsArray.from_proto(
                resource.iam_permissions
            ),
            resource_path=Primitive.from_proto(resource.resource_path),
            api_attrs=InstanceGoogleprotobufstruct.from_proto(resource.api_attrs),
            policy_name_mode=InstancePreprocessPassthroughRecipeStepsPermissionsInfoPolicyNameModeEnum.from_proto(
                resource.policy_name_mode
            ),
            resource=InstancePreprocessPassthroughRecipeStepsPermissionsInfoResource.from_proto(
                resource.resource
            ),
        )


class InstancePreprocessPassthroughRecipeStepsPermissionsInfoArray(object):
    @classmethod
    def to_proto(self, resources):
        if not resources:
            return resources
        return [
            InstancePreprocessPassthroughRecipeStepsPermissionsInfo.to_proto(i)
            for i in resources
        ]

    @classmethod
    def from_proto(self, resources):
        return [
            InstancePreprocessPassthroughRecipeStepsPermissionsInfo.from_proto(i)
            for i in resources
        ]


class InstancePreprocessPassthroughRecipeStepsPermissionsInfoPolicyName(object):
    def __init__(self, type: str = None, id: str = None, region: str = None):
        self.type = type
        self.id = id
        self.region = region

    @classmethod
    def to_proto(self, resource):
        if not resource:
            return None

        res = (
            instance_pb2.Tier2AlphaInstancePreprocessPassthroughRecipeStepsPermissionsInfoPolicyName()
        )
        if Primitive.to_proto(resource.type):
            res.type = Primitive.to_proto(resource.type)
        if Primitive.to_proto(resource.id):
            res.id = Primitive.to_proto(resource.id)
        if Primitive.to_proto(resource.region):
            res.region = Primitive.to_proto(resource.region)
        return res

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return None

        return InstancePreprocessPassthroughRecipeStepsPermissionsInfoPolicyName(
            type=Primitive.from_proto(resource.type),
            id=Primitive.from_proto(resource.id),
            region=Primitive.from_proto(resource.region),
        )


class InstancePreprocessPassthroughRecipeStepsPermissionsInfoPolicyNameArray(object):
    @classmethod
    def to_proto(self, resources):
        if not resources:
            return resources
        return [
            InstancePreprocessPassthroughRecipeStepsPermissionsInfoPolicyName.to_proto(
                i
            )
            for i in resources
        ]

    @classmethod
    def from_proto(self, resources):
        return [
            InstancePreprocessPassthroughRecipeStepsPermissionsInfoPolicyName.from_proto(
                i
            )
            for i in resources
        ]


class InstancePreprocessPassthroughRecipeStepsPermissionsInfoIamPermissions(object):
    def __init__(self, permission: str = None):
        self.permission = permission

    @classmethod
    def to_proto(self, resource):
        if not resource:
            return None

        res = (
            instance_pb2.Tier2AlphaInstancePreprocessPassthroughRecipeStepsPermissionsInfoIamPermissions()
        )
        if Primitive.to_proto(resource.permission):
            res.permission = Primitive.to_proto(resource.permission)
        return res

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return None

        return InstancePreprocessPassthroughRecipeStepsPermissionsInfoIamPermissions(
            permission=Primitive.from_proto(resource.permission),
        )


class InstancePreprocessPassthroughRecipeStepsPermissionsInfoIamPermissionsArray(
    object
):
    @classmethod
    def to_proto(self, resources):
        if not resources:
            return resources
        return [
            InstancePreprocessPassthroughRecipeStepsPermissionsInfoIamPermissions.to_proto(
                i
            )
            for i in resources
        ]

    @classmethod
    def from_proto(self, resources):
        return [
            InstancePreprocessPassthroughRecipeStepsPermissionsInfoIamPermissions.from_proto(
                i
            )
            for i in resources
        ]


class InstancePreprocessPassthroughRecipeStepsPermissionsInfoResource(object):
    def __init__(
        self,
        name: str = None,
        type: str = None,
        service: str = None,
        labels: dict = None,
    ):
        self.name = name
        self.type = type
        self.service = service
        self.labels = labels

    @classmethod
    def to_proto(self, resource):
        if not resource:
            return None

        res = (
            instance_pb2.Tier2AlphaInstancePreprocessPassthroughRecipeStepsPermissionsInfoResource()
        )
        if Primitive.to_proto(resource.name):
            res.name = Primitive.to_proto(resource.name)
        if Primitive.to_proto(resource.type):
            res.type = Primitive.to_proto(resource.type)
        if Primitive.to_proto(resource.service):
            res.service = Primitive.to_proto(resource.service)
        if Primitive.to_proto(resource.labels):
            res.labels = Primitive.to_proto(resource.labels)
        return res

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return None

        return InstancePreprocessPassthroughRecipeStepsPermissionsInfoResource(
            name=Primitive.from_proto(resource.name),
            type=Primitive.from_proto(resource.type),
            service=Primitive.from_proto(resource.service),
            labels=Primitive.from_proto(resource.labels),
        )


class InstancePreprocessPassthroughRecipeStepsPermissionsInfoResourceArray(object):
    @classmethod
    def to_proto(self, resources):
        if not resources:
            return resources
        return [
            InstancePreprocessPassthroughRecipeStepsPermissionsInfoResource.to_proto(i)
            for i in resources
        ]

    @classmethod
    def from_proto(self, resources):
        return [
            InstancePreprocessPassthroughRecipeStepsPermissionsInfoResource.from_proto(
                i
            )
            for i in resources
        ]


class InstancePreprocessPassthroughRecipeStepsKeyNotificationsUpdate(object):
    def __init__(self, key_notifications_info: dict = None):
        self.key_notifications_info = key_notifications_info

    @classmethod
    def to_proto(self, resource):
        if not resource:
            return None

        res = (
            instance_pb2.Tier2AlphaInstancePreprocessPassthroughRecipeStepsKeyNotificationsUpdate()
        )
        if InstancePreprocessPassthroughRecipeStepsKeyNotificationsUpdateKeyNotificationsInfo.to_proto(
            resource.key_notifications_info
        ):
            res.key_notifications_info.CopyFrom(
                InstancePreprocessPassthroughRecipeStepsKeyNotificationsUpdateKeyNotificationsInfo.to_proto(
                    resource.key_notifications_info
                )
            )
        else:
            res.ClearField("key_notifications_info")
        return res

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return None

        return InstancePreprocessPassthroughRecipeStepsKeyNotificationsUpdate(
            key_notifications_info=InstancePreprocessPassthroughRecipeStepsKeyNotificationsUpdateKeyNotificationsInfo.from_proto(
                resource.key_notifications_info
            ),
        )


class InstancePreprocessPassthroughRecipeStepsKeyNotificationsUpdateArray(object):
    @classmethod
    def to_proto(self, resources):
        if not resources:
            return resources
        return [
            InstancePreprocessPassthroughRecipeStepsKeyNotificationsUpdate.to_proto(i)
            for i in resources
        ]

    @classmethod
    def from_proto(self, resources):
        return [
            InstancePreprocessPassthroughRecipeStepsKeyNotificationsUpdate.from_proto(i)
            for i in resources
        ]


class InstancePreprocessPassthroughRecipeStepsKeyNotificationsUpdateKeyNotificationsInfo(
    object
):
    def __init__(
        self,
        data_version: int = None,
        delegate: str = None,
        key_notification_configs: list = None,
    ):
        self.data_version = data_version
        self.delegate = delegate
        self.key_notification_configs = key_notification_configs

    @classmethod
    def to_proto(self, resource):
        if not resource:
            return None

        res = (
            instance_pb2.Tier2AlphaInstancePreprocessPassthroughRecipeStepsKeyNotificationsUpdateKeyNotificationsInfo()
        )
        if Primitive.to_proto(resource.data_version):
            res.data_version = Primitive.to_proto(resource.data_version)
        if Primitive.to_proto(resource.delegate):
            res.delegate = Primitive.to_proto(resource.delegate)
        if InstancePreprocessPassthroughRecipeStepsKeyNotificationsUpdateKeyNotificationsInfoKeyNotificationConfigsArray.to_proto(
            resource.key_notification_configs
        ):
            res.key_notification_configs.extend(
                InstancePreprocessPassthroughRecipeStepsKeyNotificationsUpdateKeyNotificationsInfoKeyNotificationConfigsArray.to_proto(
                    resource.key_notification_configs
                )
            )
        return res

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return None

        return InstancePreprocessPassthroughRecipeStepsKeyNotificationsUpdateKeyNotificationsInfo(
            data_version=Primitive.from_proto(resource.data_version),
            delegate=Primitive.from_proto(resource.delegate),
            key_notification_configs=InstancePreprocessPassthroughRecipeStepsKeyNotificationsUpdateKeyNotificationsInfoKeyNotificationConfigsArray.from_proto(
                resource.key_notification_configs
            ),
        )


class InstancePreprocessPassthroughRecipeStepsKeyNotificationsUpdateKeyNotificationsInfoArray(
    object
):
    @classmethod
    def to_proto(self, resources):
        if not resources:
            return resources
        return [
            InstancePreprocessPassthroughRecipeStepsKeyNotificationsUpdateKeyNotificationsInfo.to_proto(
                i
            )
            for i in resources
        ]

    @classmethod
    def from_proto(self, resources):
        return [
            InstancePreprocessPassthroughRecipeStepsKeyNotificationsUpdateKeyNotificationsInfo.from_proto(
                i
            )
            for i in resources
        ]


class InstancePreprocessPassthroughRecipeStepsKeyNotificationsUpdateKeyNotificationsInfoKeyNotificationConfigs(
    object
):
    def __init__(
        self,
        key_or_version_name: str = None,
        grant: str = None,
        delegator_gaia_id: int = None,
    ):
        self.key_or_version_name = key_or_version_name
        self.grant = grant
        self.delegator_gaia_id = delegator_gaia_id

    @classmethod
    def to_proto(self, resource):
        if not resource:
            return None

        res = (
            instance_pb2.Tier2AlphaInstancePreprocessPassthroughRecipeStepsKeyNotificationsUpdateKeyNotificationsInfoKeyNotificationConfigs()
        )
        if Primitive.to_proto(resource.key_or_version_name):
            res.key_or_version_name = Primitive.to_proto(resource.key_or_version_name)
        if Primitive.to_proto(resource.grant):
            res.grant = Primitive.to_proto(resource.grant)
        if Primitive.to_proto(resource.delegator_gaia_id):
            res.delegator_gaia_id = Primitive.to_proto(resource.delegator_gaia_id)
        return res

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return None

        return InstancePreprocessPassthroughRecipeStepsKeyNotificationsUpdateKeyNotificationsInfoKeyNotificationConfigs(
            key_or_version_name=Primitive.from_proto(resource.key_or_version_name),
            grant=Primitive.from_proto(resource.grant),
            delegator_gaia_id=Primitive.from_proto(resource.delegator_gaia_id),
        )


class InstancePreprocessPassthroughRecipeStepsKeyNotificationsUpdateKeyNotificationsInfoKeyNotificationConfigsArray(
    object
):
    @classmethod
    def to_proto(self, resources):
        if not resources:
            return resources
        return [
            InstancePreprocessPassthroughRecipeStepsKeyNotificationsUpdateKeyNotificationsInfoKeyNotificationConfigs.to_proto(
                i
            )
            for i in resources
        ]

    @classmethod
    def from_proto(self, resources):
        return [
            InstancePreprocessPassthroughRecipeStepsKeyNotificationsUpdateKeyNotificationsInfoKeyNotificationConfigs.from_proto(
                i
            )
            for i in resources
        ]


class InstancePreprocessReconcileRecipe(object):
    def __init__(
        self,
        steps: list = None,
        honor_cancel_request: bool = None,
        ignore_recipe_after: int = None,
        verify_deadline_seconds_below: float = None,
        populate_operation_result: bool = None,
        readonly_recipe_start_time: str = None,
        resource_names_stored_in_clh_with_delay: list = None,
        delay_to_store_resources_in_clh_db_nanos: int = None,
    ):
        self.steps = steps
        self.honor_cancel_request = honor_cancel_request
        self.ignore_recipe_after = ignore_recipe_after
        self.verify_deadline_seconds_below = verify_deadline_seconds_below
        self.populate_operation_result = populate_operation_result
        self.readonly_recipe_start_time = readonly_recipe_start_time
        self.resource_names_stored_in_clh_with_delay = (
            resource_names_stored_in_clh_with_delay
        )
        self.delay_to_store_resources_in_clh_db_nanos = (
            delay_to_store_resources_in_clh_db_nanos
        )

    @classmethod
    def to_proto(self, resource):
        if not resource:
            return None

        res = instance_pb2.Tier2AlphaInstancePreprocessReconcileRecipe()
        if InstancePreprocessReconcileRecipeStepsArray.to_proto(resource.steps):
            res.steps.extend(
                InstancePreprocessReconcileRecipeStepsArray.to_proto(resource.steps)
            )
        if Primitive.to_proto(resource.honor_cancel_request):
            res.honor_cancel_request = Primitive.to_proto(resource.honor_cancel_request)
        if Primitive.to_proto(resource.ignore_recipe_after):
            res.ignore_recipe_after = Primitive.to_proto(resource.ignore_recipe_after)
        if Primitive.to_proto(resource.verify_deadline_seconds_below):
            res.verify_deadline_seconds_below = Primitive.to_proto(
                resource.verify_deadline_seconds_below
            )
        if Primitive.to_proto(resource.populate_operation_result):
            res.populate_operation_result = Primitive.to_proto(
                resource.populate_operation_result
            )
        if Primitive.to_proto(resource.readonly_recipe_start_time):
            res.readonly_recipe_start_time = Primitive.to_proto(
                resource.readonly_recipe_start_time
            )
        if Primitive.to_proto(resource.resource_names_stored_in_clh_with_delay):
            res.resource_names_stored_in_clh_with_delay.extend(
                Primitive.to_proto(resource.resource_names_stored_in_clh_with_delay)
            )
        if Primitive.to_proto(resource.delay_to_store_resources_in_clh_db_nanos):
            res.delay_to_store_resources_in_clh_db_nanos = Primitive.to_proto(
                resource.delay_to_store_resources_in_clh_db_nanos
            )
        return res

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return None

        return InstancePreprocessReconcileRecipe(
            steps=InstancePreprocessReconcileRecipeStepsArray.from_proto(
                resource.steps
            ),
            honor_cancel_request=Primitive.from_proto(resource.honor_cancel_request),
            ignore_recipe_after=Primitive.from_proto(resource.ignore_recipe_after),
            verify_deadline_seconds_below=Primitive.from_proto(
                resource.verify_deadline_seconds_below
            ),
            populate_operation_result=Primitive.from_proto(
                resource.populate_operation_result
            ),
            readonly_recipe_start_time=Primitive.from_proto(
                resource.readonly_recipe_start_time
            ),
            resource_names_stored_in_clh_with_delay=Primitive.from_proto(
                resource.resource_names_stored_in_clh_with_delay
            ),
            delay_to_store_resources_in_clh_db_nanos=Primitive.from_proto(
                resource.delay_to_store_resources_in_clh_db_nanos
            ),
        )


class InstancePreprocessReconcileRecipeArray(object):
    @classmethod
    def to_proto(self, resources):
        if not resources:
            return resources
        return [InstancePreprocessReconcileRecipe.to_proto(i) for i in resources]

    @classmethod
    def from_proto(self, resources):
        return [InstancePreprocessReconcileRecipe.from_proto(i) for i in resources]


class InstancePreprocessReconcileRecipeSteps(object):
    def __init__(
        self,
        relative_time: int = None,
        sleep_duration: int = None,
        action: str = None,
        status: dict = None,
        error_space: str = None,
        p4_service_account: str = None,
        resource_metadata_size: int = None,
        description: str = None,
        updated_repeat_operation_delay_sec: float = None,
        quota_request_deltas: list = None,
        preprocess_update: dict = None,
        public_operation_metadata: str = None,
        requested_tenant_project: dict = None,
        permissions_info: list = None,
        key_notifications_update: dict = None,
        clh_data_update_time: str = None,
        public_error_message: str = None,
    ):
        self.relative_time = relative_time
        self.sleep_duration = sleep_duration
        self.action = action
        self.status = status
        self.error_space = error_space
        self.p4_service_account = p4_service_account
        self.resource_metadata_size = resource_metadata_size
        self.description = description
        self.updated_repeat_operation_delay_sec = updated_repeat_operation_delay_sec
        self.quota_request_deltas = quota_request_deltas
        self.preprocess_update = preprocess_update
        self.public_operation_metadata = public_operation_metadata
        self.requested_tenant_project = requested_tenant_project
        self.permissions_info = permissions_info
        self.key_notifications_update = key_notifications_update
        self.clh_data_update_time = clh_data_update_time
        self.public_error_message = public_error_message

    @classmethod
    def to_proto(self, resource):
        if not resource:
            return None

        res = instance_pb2.Tier2AlphaInstancePreprocessReconcileRecipeSteps()
        if Primitive.to_proto(resource.relative_time):
            res.relative_time = Primitive.to_proto(resource.relative_time)
        if Primitive.to_proto(resource.sleep_duration):
            res.sleep_duration = Primitive.to_proto(resource.sleep_duration)
        if InstancePreprocessReconcileRecipeStepsActionEnum.to_proto(resource.action):
            res.action = InstancePreprocessReconcileRecipeStepsActionEnum.to_proto(
                resource.action
            )
        if InstancePreprocessReconcileRecipeStepsStatus.to_proto(resource.status):
            res.status.CopyFrom(
                InstancePreprocessReconcileRecipeStepsStatus.to_proto(resource.status)
            )
        else:
            res.ClearField("status")
        if Primitive.to_proto(resource.error_space):
            res.error_space = Primitive.to_proto(resource.error_space)
        if Primitive.to_proto(resource.p4_service_account):
            res.p4_service_account = Primitive.to_proto(resource.p4_service_account)
        if Primitive.to_proto(resource.resource_metadata_size):
            res.resource_metadata_size = Primitive.to_proto(
                resource.resource_metadata_size
            )
        if Primitive.to_proto(resource.description):
            res.description = Primitive.to_proto(resource.description)
        if Primitive.to_proto(resource.updated_repeat_operation_delay_sec):
            res.updated_repeat_operation_delay_sec = Primitive.to_proto(
                resource.updated_repeat_operation_delay_sec
            )
        if InstancePreprocessReconcileRecipeStepsQuotaRequestDeltasArray.to_proto(
            resource.quota_request_deltas
        ):
            res.quota_request_deltas.extend(
                InstancePreprocessReconcileRecipeStepsQuotaRequestDeltasArray.to_proto(
                    resource.quota_request_deltas
                )
            )
        if InstancePreprocessReconcileRecipeStepsPreprocessUpdate.to_proto(
            resource.preprocess_update
        ):
            res.preprocess_update.CopyFrom(
                InstancePreprocessReconcileRecipeStepsPreprocessUpdate.to_proto(
                    resource.preprocess_update
                )
            )
        else:
            res.ClearField("preprocess_update")
        if Primitive.to_proto(resource.public_operation_metadata):
            res.public_operation_metadata = Primitive.to_proto(
                resource.public_operation_metadata
            )
        if InstancePreprocessReconcileRecipeStepsRequestedTenantProject.to_proto(
            resource.requested_tenant_project
        ):
            res.requested_tenant_project.CopyFrom(
                InstancePreprocessReconcileRecipeStepsRequestedTenantProject.to_proto(
                    resource.requested_tenant_project
                )
            )
        else:
            res.ClearField("requested_tenant_project")
        if InstancePreprocessReconcileRecipeStepsPermissionsInfoArray.to_proto(
            resource.permissions_info
        ):
            res.permissions_info.extend(
                InstancePreprocessReconcileRecipeStepsPermissionsInfoArray.to_proto(
                    resource.permissions_info
                )
            )
        if InstancePreprocessReconcileRecipeStepsKeyNotificationsUpdate.to_proto(
            resource.key_notifications_update
        ):
            res.key_notifications_update.CopyFrom(
                InstancePreprocessReconcileRecipeStepsKeyNotificationsUpdate.to_proto(
                    resource.key_notifications_update
                )
            )
        else:
            res.ClearField("key_notifications_update")
        if Primitive.to_proto(resource.clh_data_update_time):
            res.clh_data_update_time = Primitive.to_proto(resource.clh_data_update_time)
        if Primitive.to_proto(resource.public_error_message):
            res.public_error_message = Primitive.to_proto(resource.public_error_message)
        return res

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return None

        return InstancePreprocessReconcileRecipeSteps(
            relative_time=Primitive.from_proto(resource.relative_time),
            sleep_duration=Primitive.from_proto(resource.sleep_duration),
            action=InstancePreprocessReconcileRecipeStepsActionEnum.from_proto(
                resource.action
            ),
            status=InstancePreprocessReconcileRecipeStepsStatus.from_proto(
                resource.status
            ),
            error_space=Primitive.from_proto(resource.error_space),
            p4_service_account=Primitive.from_proto(resource.p4_service_account),
            resource_metadata_size=Primitive.from_proto(
                resource.resource_metadata_size
            ),
            description=Primitive.from_proto(resource.description),
            updated_repeat_operation_delay_sec=Primitive.from_proto(
                resource.updated_repeat_operation_delay_sec
            ),
            quota_request_deltas=InstancePreprocessReconcileRecipeStepsQuotaRequestDeltasArray.from_proto(
                resource.quota_request_deltas
            ),
            preprocess_update=InstancePreprocessReconcileRecipeStepsPreprocessUpdate.from_proto(
                resource.preprocess_update
            ),
            public_operation_metadata=Primitive.from_proto(
                resource.public_operation_metadata
            ),
            requested_tenant_project=InstancePreprocessReconcileRecipeStepsRequestedTenantProject.from_proto(
                resource.requested_tenant_project
            ),
            permissions_info=InstancePreprocessReconcileRecipeStepsPermissionsInfoArray.from_proto(
                resource.permissions_info
            ),
            key_notifications_update=InstancePreprocessReconcileRecipeStepsKeyNotificationsUpdate.from_proto(
                resource.key_notifications_update
            ),
            clh_data_update_time=Primitive.from_proto(resource.clh_data_update_time),
            public_error_message=Primitive.from_proto(resource.public_error_message),
        )


class InstancePreprocessReconcileRecipeStepsArray(object):
    @classmethod
    def to_proto(self, resources):
        if not resources:
            return resources
        return [InstancePreprocessReconcileRecipeSteps.to_proto(i) for i in resources]

    @classmethod
    def from_proto(self, resources):
        return [InstancePreprocessReconcileRecipeSteps.from_proto(i) for i in resources]


class InstancePreprocessReconcileRecipeStepsStatus(object):
    def __init__(self, code: int = None, message: str = None, details: list = None):
        self.code = code
        self.message = message
        self.details = details

    @classmethod
    def to_proto(self, resource):
        if not resource:
            return None

        res = instance_pb2.Tier2AlphaInstancePreprocessReconcileRecipeStepsStatus()
        if Primitive.to_proto(resource.code):
            res.code = Primitive.to_proto(resource.code)
        if Primitive.to_proto(resource.message):
            res.message = Primitive.to_proto(resource.message)
        if InstancePreprocessReconcileRecipeStepsStatusDetailsArray.to_proto(
            resource.details
        ):
            res.details.extend(
                InstancePreprocessReconcileRecipeStepsStatusDetailsArray.to_proto(
                    resource.details
                )
            )
        return res

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return None

        return InstancePreprocessReconcileRecipeStepsStatus(
            code=Primitive.from_proto(resource.code),
            message=Primitive.from_proto(resource.message),
            details=InstancePreprocessReconcileRecipeStepsStatusDetailsArray.from_proto(
                resource.details
            ),
        )


class InstancePreprocessReconcileRecipeStepsStatusArray(object):
    @classmethod
    def to_proto(self, resources):
        if not resources:
            return resources
        return [
            InstancePreprocessReconcileRecipeStepsStatus.to_proto(i) for i in resources
        ]

    @classmethod
    def from_proto(self, resources):
        return [
            InstancePreprocessReconcileRecipeStepsStatus.from_proto(i)
            for i in resources
        ]


class InstancePreprocessReconcileRecipeStepsStatusDetails(object):
    def __init__(self, type_url: str = None, value: str = None):
        self.type_url = type_url
        self.value = value

    @classmethod
    def to_proto(self, resource):
        if not resource:
            return None

        res = (
            instance_pb2.Tier2AlphaInstancePreprocessReconcileRecipeStepsStatusDetails()
        )
        if Primitive.to_proto(resource.type_url):
            res.type_url = Primitive.to_proto(resource.type_url)
        if Primitive.to_proto(resource.value):
            res.value = Primitive.to_proto(resource.value)
        return res

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return None

        return InstancePreprocessReconcileRecipeStepsStatusDetails(
            type_url=Primitive.from_proto(resource.type_url),
            value=Primitive.from_proto(resource.value),
        )


class InstancePreprocessReconcileRecipeStepsStatusDetailsArray(object):
    @classmethod
    def to_proto(self, resources):
        if not resources:
            return resources
        return [
            InstancePreprocessReconcileRecipeStepsStatusDetails.to_proto(i)
            for i in resources
        ]

    @classmethod
    def from_proto(self, resources):
        return [
            InstancePreprocessReconcileRecipeStepsStatusDetails.from_proto(i)
            for i in resources
        ]


class InstancePreprocessReconcileRecipeStepsQuotaRequestDeltas(object):
    def __init__(
        self,
        metric_name: str = None,
        amount: int = None,
        quota_location_name: str = None,
    ):
        self.metric_name = metric_name
        self.amount = amount
        self.quota_location_name = quota_location_name

    @classmethod
    def to_proto(self, resource):
        if not resource:
            return None

        res = (
            instance_pb2.Tier2AlphaInstancePreprocessReconcileRecipeStepsQuotaRequestDeltas()
        )
        if Primitive.to_proto(resource.metric_name):
            res.metric_name = Primitive.to_proto(resource.metric_name)
        if Primitive.to_proto(resource.amount):
            res.amount = Primitive.to_proto(resource.amount)
        if Primitive.to_proto(resource.quota_location_name):
            res.quota_location_name = Primitive.to_proto(resource.quota_location_name)
        return res

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return None

        return InstancePreprocessReconcileRecipeStepsQuotaRequestDeltas(
            metric_name=Primitive.from_proto(resource.metric_name),
            amount=Primitive.from_proto(resource.amount),
            quota_location_name=Primitive.from_proto(resource.quota_location_name),
        )


class InstancePreprocessReconcileRecipeStepsQuotaRequestDeltasArray(object):
    @classmethod
    def to_proto(self, resources):
        if not resources:
            return resources
        return [
            InstancePreprocessReconcileRecipeStepsQuotaRequestDeltas.to_proto(i)
            for i in resources
        ]

    @classmethod
    def from_proto(self, resources):
        return [
            InstancePreprocessReconcileRecipeStepsQuotaRequestDeltas.from_proto(i)
            for i in resources
        ]


class InstancePreprocessReconcileRecipeStepsPreprocessUpdate(object):
    def __init__(
        self, latency_slo_bucket_name: str = None, public_operation_metadata: str = None
    ):
        self.latency_slo_bucket_name = latency_slo_bucket_name
        self.public_operation_metadata = public_operation_metadata

    @classmethod
    def to_proto(self, resource):
        if not resource:
            return None

        res = (
            instance_pb2.Tier2AlphaInstancePreprocessReconcileRecipeStepsPreprocessUpdate()
        )
        if Primitive.to_proto(resource.latency_slo_bucket_name):
            res.latency_slo_bucket_name = Primitive.to_proto(
                resource.latency_slo_bucket_name
            )
        if Primitive.to_proto(resource.public_operation_metadata):
            res.public_operation_metadata = Primitive.to_proto(
                resource.public_operation_metadata
            )
        return res

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return None

        return InstancePreprocessReconcileRecipeStepsPreprocessUpdate(
            latency_slo_bucket_name=Primitive.from_proto(
                resource.latency_slo_bucket_name
            ),
            public_operation_metadata=Primitive.from_proto(
                resource.public_operation_metadata
            ),
        )


class InstancePreprocessReconcileRecipeStepsPreprocessUpdateArray(object):
    @classmethod
    def to_proto(self, resources):
        if not resources:
            return resources
        return [
            InstancePreprocessReconcileRecipeStepsPreprocessUpdate.to_proto(i)
            for i in resources
        ]

    @classmethod
    def from_proto(self, resources):
        return [
            InstancePreprocessReconcileRecipeStepsPreprocessUpdate.from_proto(i)
            for i in resources
        ]


class InstancePreprocessReconcileRecipeStepsRequestedTenantProject(object):
    def __init__(self, tag: str = None, folder: str = None, scope: str = None):
        self.tag = tag
        self.folder = folder
        self.scope = scope

    @classmethod
    def to_proto(self, resource):
        if not resource:
            return None

        res = (
            instance_pb2.Tier2AlphaInstancePreprocessReconcileRecipeStepsRequestedTenantProject()
        )
        if Primitive.to_proto(resource.tag):
            res.tag = Primitive.to_proto(resource.tag)
        if Primitive.to_proto(resource.folder):
            res.folder = Primitive.to_proto(resource.folder)
        if InstancePreprocessReconcileRecipeStepsRequestedTenantProjectScopeEnum.to_proto(
            resource.scope
        ):
            res.scope = InstancePreprocessReconcileRecipeStepsRequestedTenantProjectScopeEnum.to_proto(
                resource.scope
            )
        return res

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return None

        return InstancePreprocessReconcileRecipeStepsRequestedTenantProject(
            tag=Primitive.from_proto(resource.tag),
            folder=Primitive.from_proto(resource.folder),
            scope=InstancePreprocessReconcileRecipeStepsRequestedTenantProjectScopeEnum.from_proto(
                resource.scope
            ),
        )


class InstancePreprocessReconcileRecipeStepsRequestedTenantProjectArray(object):
    @classmethod
    def to_proto(self, resources):
        if not resources:
            return resources
        return [
            InstancePreprocessReconcileRecipeStepsRequestedTenantProject.to_proto(i)
            for i in resources
        ]

    @classmethod
    def from_proto(self, resources):
        return [
            InstancePreprocessReconcileRecipeStepsRequestedTenantProject.from_proto(i)
            for i in resources
        ]


class InstancePreprocessReconcileRecipeStepsPermissionsInfo(object):
    def __init__(
        self,
        policy_name: dict = None,
        iam_permissions: list = None,
        resource_path: str = None,
        api_attrs: dict = None,
        policy_name_mode: str = None,
        resource: dict = None,
    ):
        self.policy_name = policy_name
        self.iam_permissions = iam_permissions
        self.resource_path = resource_path
        self.api_attrs = api_attrs
        self.policy_name_mode = policy_name_mode
        self.resource = resource

    @classmethod
    def to_proto(self, resource):
        if not resource:
            return None

        res = (
            instance_pb2.Tier2AlphaInstancePreprocessReconcileRecipeStepsPermissionsInfo()
        )
        if InstancePreprocessReconcileRecipeStepsPermissionsInfoPolicyName.to_proto(
            resource.policy_name
        ):
            res.policy_name.CopyFrom(
                InstancePreprocessReconcileRecipeStepsPermissionsInfoPolicyName.to_proto(
                    resource.policy_name
                )
            )
        else:
            res.ClearField("policy_name")
        if InstancePreprocessReconcileRecipeStepsPermissionsInfoIamPermissionsArray.to_proto(
            resource.iam_permissions
        ):
            res.iam_permissions.extend(
                InstancePreprocessReconcileRecipeStepsPermissionsInfoIamPermissionsArray.to_proto(
                    resource.iam_permissions
                )
            )
        if Primitive.to_proto(resource.resource_path):
            res.resource_path = Primitive.to_proto(resource.resource_path)
        if InstanceGoogleprotobufstruct.to_proto(resource.api_attrs):
            res.api_attrs.CopyFrom(
                InstanceGoogleprotobufstruct.to_proto(resource.api_attrs)
            )
        else:
            res.ClearField("api_attrs")
        if InstancePreprocessReconcileRecipeStepsPermissionsInfoPolicyNameModeEnum.to_proto(
            resource.policy_name_mode
        ):
            res.policy_name_mode = InstancePreprocessReconcileRecipeStepsPermissionsInfoPolicyNameModeEnum.to_proto(
                resource.policy_name_mode
            )
        if InstancePreprocessReconcileRecipeStepsPermissionsInfoResource.to_proto(
            resource.resource
        ):
            res.resource.CopyFrom(
                InstancePreprocessReconcileRecipeStepsPermissionsInfoResource.to_proto(
                    resource.resource
                )
            )
        else:
            res.ClearField("resource")
        return res

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return None

        return InstancePreprocessReconcileRecipeStepsPermissionsInfo(
            policy_name=InstancePreprocessReconcileRecipeStepsPermissionsInfoPolicyName.from_proto(
                resource.policy_name
            ),
            iam_permissions=InstancePreprocessReconcileRecipeStepsPermissionsInfoIamPermissionsArray.from_proto(
                resource.iam_permissions
            ),
            resource_path=Primitive.from_proto(resource.resource_path),
            api_attrs=InstanceGoogleprotobufstruct.from_proto(resource.api_attrs),
            policy_name_mode=InstancePreprocessReconcileRecipeStepsPermissionsInfoPolicyNameModeEnum.from_proto(
                resource.policy_name_mode
            ),
            resource=InstancePreprocessReconcileRecipeStepsPermissionsInfoResource.from_proto(
                resource.resource
            ),
        )


class InstancePreprocessReconcileRecipeStepsPermissionsInfoArray(object):
    @classmethod
    def to_proto(self, resources):
        if not resources:
            return resources
        return [
            InstancePreprocessReconcileRecipeStepsPermissionsInfo.to_proto(i)
            for i in resources
        ]

    @classmethod
    def from_proto(self, resources):
        return [
            InstancePreprocessReconcileRecipeStepsPermissionsInfo.from_proto(i)
            for i in resources
        ]


class InstancePreprocessReconcileRecipeStepsPermissionsInfoPolicyName(object):
    def __init__(self, type: str = None, id: str = None, region: str = None):
        self.type = type
        self.id = id
        self.region = region

    @classmethod
    def to_proto(self, resource):
        if not resource:
            return None

        res = (
            instance_pb2.Tier2AlphaInstancePreprocessReconcileRecipeStepsPermissionsInfoPolicyName()
        )
        if Primitive.to_proto(resource.type):
            res.type = Primitive.to_proto(resource.type)
        if Primitive.to_proto(resource.id):
            res.id = Primitive.to_proto(resource.id)
        if Primitive.to_proto(resource.region):
            res.region = Primitive.to_proto(resource.region)
        return res

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return None

        return InstancePreprocessReconcileRecipeStepsPermissionsInfoPolicyName(
            type=Primitive.from_proto(resource.type),
            id=Primitive.from_proto(resource.id),
            region=Primitive.from_proto(resource.region),
        )


class InstancePreprocessReconcileRecipeStepsPermissionsInfoPolicyNameArray(object):
    @classmethod
    def to_proto(self, resources):
        if not resources:
            return resources
        return [
            InstancePreprocessReconcileRecipeStepsPermissionsInfoPolicyName.to_proto(i)
            for i in resources
        ]

    @classmethod
    def from_proto(self, resources):
        return [
            InstancePreprocessReconcileRecipeStepsPermissionsInfoPolicyName.from_proto(
                i
            )
            for i in resources
        ]


class InstancePreprocessReconcileRecipeStepsPermissionsInfoIamPermissions(object):
    def __init__(self, permission: str = None):
        self.permission = permission

    @classmethod
    def to_proto(self, resource):
        if not resource:
            return None

        res = (
            instance_pb2.Tier2AlphaInstancePreprocessReconcileRecipeStepsPermissionsInfoIamPermissions()
        )
        if Primitive.to_proto(resource.permission):
            res.permission = Primitive.to_proto(resource.permission)
        return res

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return None

        return InstancePreprocessReconcileRecipeStepsPermissionsInfoIamPermissions(
            permission=Primitive.from_proto(resource.permission),
        )


class InstancePreprocessReconcileRecipeStepsPermissionsInfoIamPermissionsArray(object):
    @classmethod
    def to_proto(self, resources):
        if not resources:
            return resources
        return [
            InstancePreprocessReconcileRecipeStepsPermissionsInfoIamPermissions.to_proto(
                i
            )
            for i in resources
        ]

    @classmethod
    def from_proto(self, resources):
        return [
            InstancePreprocessReconcileRecipeStepsPermissionsInfoIamPermissions.from_proto(
                i
            )
            for i in resources
        ]


class InstancePreprocessReconcileRecipeStepsPermissionsInfoResource(object):
    def __init__(
        self,
        name: str = None,
        type: str = None,
        service: str = None,
        labels: dict = None,
    ):
        self.name = name
        self.type = type
        self.service = service
        self.labels = labels

    @classmethod
    def to_proto(self, resource):
        if not resource:
            return None

        res = (
            instance_pb2.Tier2AlphaInstancePreprocessReconcileRecipeStepsPermissionsInfoResource()
        )
        if Primitive.to_proto(resource.name):
            res.name = Primitive.to_proto(resource.name)
        if Primitive.to_proto(resource.type):
            res.type = Primitive.to_proto(resource.type)
        if Primitive.to_proto(resource.service):
            res.service = Primitive.to_proto(resource.service)
        if Primitive.to_proto(resource.labels):
            res.labels = Primitive.to_proto(resource.labels)
        return res

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return None

        return InstancePreprocessReconcileRecipeStepsPermissionsInfoResource(
            name=Primitive.from_proto(resource.name),
            type=Primitive.from_proto(resource.type),
            service=Primitive.from_proto(resource.service),
            labels=Primitive.from_proto(resource.labels),
        )


class InstancePreprocessReconcileRecipeStepsPermissionsInfoResourceArray(object):
    @classmethod
    def to_proto(self, resources):
        if not resources:
            return resources
        return [
            InstancePreprocessReconcileRecipeStepsPermissionsInfoResource.to_proto(i)
            for i in resources
        ]

    @classmethod
    def from_proto(self, resources):
        return [
            InstancePreprocessReconcileRecipeStepsPermissionsInfoResource.from_proto(i)
            for i in resources
        ]


class InstancePreprocessReconcileRecipeStepsKeyNotificationsUpdate(object):
    def __init__(self, key_notifications_info: dict = None):
        self.key_notifications_info = key_notifications_info

    @classmethod
    def to_proto(self, resource):
        if not resource:
            return None

        res = (
            instance_pb2.Tier2AlphaInstancePreprocessReconcileRecipeStepsKeyNotificationsUpdate()
        )
        if InstancePreprocessReconcileRecipeStepsKeyNotificationsUpdateKeyNotificationsInfo.to_proto(
            resource.key_notifications_info
        ):
            res.key_notifications_info.CopyFrom(
                InstancePreprocessReconcileRecipeStepsKeyNotificationsUpdateKeyNotificationsInfo.to_proto(
                    resource.key_notifications_info
                )
            )
        else:
            res.ClearField("key_notifications_info")
        return res

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return None

        return InstancePreprocessReconcileRecipeStepsKeyNotificationsUpdate(
            key_notifications_info=InstancePreprocessReconcileRecipeStepsKeyNotificationsUpdateKeyNotificationsInfo.from_proto(
                resource.key_notifications_info
            ),
        )


class InstancePreprocessReconcileRecipeStepsKeyNotificationsUpdateArray(object):
    @classmethod
    def to_proto(self, resources):
        if not resources:
            return resources
        return [
            InstancePreprocessReconcileRecipeStepsKeyNotificationsUpdate.to_proto(i)
            for i in resources
        ]

    @classmethod
    def from_proto(self, resources):
        return [
            InstancePreprocessReconcileRecipeStepsKeyNotificationsUpdate.from_proto(i)
            for i in resources
        ]


class InstancePreprocessReconcileRecipeStepsKeyNotificationsUpdateKeyNotificationsInfo(
    object
):
    def __init__(
        self,
        data_version: int = None,
        delegate: str = None,
        key_notification_configs: list = None,
    ):
        self.data_version = data_version
        self.delegate = delegate
        self.key_notification_configs = key_notification_configs

    @classmethod
    def to_proto(self, resource):
        if not resource:
            return None

        res = (
            instance_pb2.Tier2AlphaInstancePreprocessReconcileRecipeStepsKeyNotificationsUpdateKeyNotificationsInfo()
        )
        if Primitive.to_proto(resource.data_version):
            res.data_version = Primitive.to_proto(resource.data_version)
        if Primitive.to_proto(resource.delegate):
            res.delegate = Primitive.to_proto(resource.delegate)
        if InstancePreprocessReconcileRecipeStepsKeyNotificationsUpdateKeyNotificationsInfoKeyNotificationConfigsArray.to_proto(
            resource.key_notification_configs
        ):
            res.key_notification_configs.extend(
                InstancePreprocessReconcileRecipeStepsKeyNotificationsUpdateKeyNotificationsInfoKeyNotificationConfigsArray.to_proto(
                    resource.key_notification_configs
                )
            )
        return res

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return None

        return InstancePreprocessReconcileRecipeStepsKeyNotificationsUpdateKeyNotificationsInfo(
            data_version=Primitive.from_proto(resource.data_version),
            delegate=Primitive.from_proto(resource.delegate),
            key_notification_configs=InstancePreprocessReconcileRecipeStepsKeyNotificationsUpdateKeyNotificationsInfoKeyNotificationConfigsArray.from_proto(
                resource.key_notification_configs
            ),
        )


class InstancePreprocessReconcileRecipeStepsKeyNotificationsUpdateKeyNotificationsInfoArray(
    object
):
    @classmethod
    def to_proto(self, resources):
        if not resources:
            return resources
        return [
            InstancePreprocessReconcileRecipeStepsKeyNotificationsUpdateKeyNotificationsInfo.to_proto(
                i
            )
            for i in resources
        ]

    @classmethod
    def from_proto(self, resources):
        return [
            InstancePreprocessReconcileRecipeStepsKeyNotificationsUpdateKeyNotificationsInfo.from_proto(
                i
            )
            for i in resources
        ]


class InstancePreprocessReconcileRecipeStepsKeyNotificationsUpdateKeyNotificationsInfoKeyNotificationConfigs(
    object
):
    def __init__(
        self,
        key_or_version_name: str = None,
        grant: str = None,
        delegator_gaia_id: int = None,
    ):
        self.key_or_version_name = key_or_version_name
        self.grant = grant
        self.delegator_gaia_id = delegator_gaia_id

    @classmethod
    def to_proto(self, resource):
        if not resource:
            return None

        res = (
            instance_pb2.Tier2AlphaInstancePreprocessReconcileRecipeStepsKeyNotificationsUpdateKeyNotificationsInfoKeyNotificationConfigs()
        )
        if Primitive.to_proto(resource.key_or_version_name):
            res.key_or_version_name = Primitive.to_proto(resource.key_or_version_name)
        if Primitive.to_proto(resource.grant):
            res.grant = Primitive.to_proto(resource.grant)
        if Primitive.to_proto(resource.delegator_gaia_id):
            res.delegator_gaia_id = Primitive.to_proto(resource.delegator_gaia_id)
        return res

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return None

        return InstancePreprocessReconcileRecipeStepsKeyNotificationsUpdateKeyNotificationsInfoKeyNotificationConfigs(
            key_or_version_name=Primitive.from_proto(resource.key_or_version_name),
            grant=Primitive.from_proto(resource.grant),
            delegator_gaia_id=Primitive.from_proto(resource.delegator_gaia_id),
        )


class InstancePreprocessReconcileRecipeStepsKeyNotificationsUpdateKeyNotificationsInfoKeyNotificationConfigsArray(
    object
):
    @classmethod
    def to_proto(self, resources):
        if not resources:
            return resources
        return [
            InstancePreprocessReconcileRecipeStepsKeyNotificationsUpdateKeyNotificationsInfoKeyNotificationConfigs.to_proto(
                i
            )
            for i in resources
        ]

    @classmethod
    def from_proto(self, resources):
        return [
            InstancePreprocessReconcileRecipeStepsKeyNotificationsUpdateKeyNotificationsInfoKeyNotificationConfigs.from_proto(
                i
            )
            for i in resources
        ]


class InstanceHistory(object):
    def __init__(
        self,
        timestamp: str = None,
        operation_handle: str = None,
        description: str = None,
        step_index: int = None,
        tenant_project_number: int = None,
        tenant_project_id: str = None,
        p4_service_account: str = None,
    ):
        self.timestamp = timestamp
        self.operation_handle = operation_handle
        self.description = description
        self.step_index = step_index
        self.tenant_project_number = tenant_project_number
        self.tenant_project_id = tenant_project_id
        self.p4_service_account = p4_service_account

    @classmethod
    def to_proto(self, resource):
        if not resource:
            return None

        res = instance_pb2.Tier2AlphaInstanceHistory()
        if Primitive.to_proto(resource.timestamp):
            res.timestamp = Primitive.to_proto(resource.timestamp)
        if Primitive.to_proto(resource.operation_handle):
            res.operation_handle = Primitive.to_proto(resource.operation_handle)
        if Primitive.to_proto(resource.description):
            res.description = Primitive.to_proto(resource.description)
        if Primitive.to_proto(resource.step_index):
            res.step_index = Primitive.to_proto(resource.step_index)
        if Primitive.to_proto(resource.tenant_project_number):
            res.tenant_project_number = Primitive.to_proto(
                resource.tenant_project_number
            )
        if Primitive.to_proto(resource.tenant_project_id):
            res.tenant_project_id = Primitive.to_proto(resource.tenant_project_id)
        if Primitive.to_proto(resource.p4_service_account):
            res.p4_service_account = Primitive.to_proto(resource.p4_service_account)
        return res

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return None

        return InstanceHistory(
            timestamp=Primitive.from_proto(resource.timestamp),
            operation_handle=Primitive.from_proto(resource.operation_handle),
            description=Primitive.from_proto(resource.description),
            step_index=Primitive.from_proto(resource.step_index),
            tenant_project_number=Primitive.from_proto(resource.tenant_project_number),
            tenant_project_id=Primitive.from_proto(resource.tenant_project_id),
            p4_service_account=Primitive.from_proto(resource.p4_service_account),
        )


class InstanceHistoryArray(object):
    @classmethod
    def to_proto(self, resources):
        if not resources:
            return resources
        return [InstanceHistory.to_proto(i) for i in resources]

    @classmethod
    def from_proto(self, resources):
        return [InstanceHistory.from_proto(i) for i in resources]


class InstanceSkuTierEnum(object):
    @classmethod
    def to_proto(self, resource):
        if not resource:
            return resource
        return instance_pb2.Tier2AlphaInstanceSkuTierEnum.Value(
            "Tier2AlphaInstanceSkuTierEnum%s" % resource
        )

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return resource
        return instance_pb2.Tier2AlphaInstanceSkuTierEnum.Name(resource)[
            len("Tier2AlphaInstanceSkuTierEnum") :
        ]


class InstanceSkuSizeEnum(object):
    @classmethod
    def to_proto(self, resource):
        if not resource:
            return resource
        return instance_pb2.Tier2AlphaInstanceSkuSizeEnum.Value(
            "Tier2AlphaInstanceSkuSizeEnum%s" % resource
        )

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return resource
        return instance_pb2.Tier2AlphaInstanceSkuSizeEnum.Name(resource)[
            len("Tier2AlphaInstanceSkuSizeEnum") :
        ]


class InstanceStateEnum(object):
    @classmethod
    def to_proto(self, resource):
        if not resource:
            return resource
        return instance_pb2.Tier2AlphaInstanceStateEnum.Value(
            "Tier2AlphaInstanceStateEnum%s" % resource
        )

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return resource
        return instance_pb2.Tier2AlphaInstanceStateEnum.Name(resource)[
            len("Tier2AlphaInstanceStateEnum") :
        ]


class InstanceEncryptionKeysKeyStateAvailabilityKeyVersionStateEnum(object):
    @classmethod
    def to_proto(self, resource):
        if not resource:
            return resource
        return instance_pb2.Tier2AlphaInstanceEncryptionKeysKeyStateAvailabilityKeyVersionStateEnum.Value(
            "Tier2AlphaInstanceEncryptionKeysKeyStateAvailabilityKeyVersionStateEnum%s"
            % resource
        )

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return resource
        return instance_pb2.Tier2AlphaInstanceEncryptionKeysKeyStateAvailabilityKeyVersionStateEnum.Name(
            resource
        )[
            len(
                "Tier2AlphaInstanceEncryptionKeysKeyStateAvailabilityKeyVersionStateEnum"
            ) :
        ]


class InstancePreprocessCreateRecipeStepsActionEnum(object):
    @classmethod
    def to_proto(self, resource):
        if not resource:
            return resource
        return instance_pb2.Tier2AlphaInstancePreprocessCreateRecipeStepsActionEnum.Value(
            "Tier2AlphaInstancePreprocessCreateRecipeStepsActionEnum%s" % resource
        )

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return resource
        return instance_pb2.Tier2AlphaInstancePreprocessCreateRecipeStepsActionEnum.Name(
            resource
        )[
            len("Tier2AlphaInstancePreprocessCreateRecipeStepsActionEnum") :
        ]


class InstancePreprocessCreateRecipeStepsRequestedTenantProjectScopeEnum(object):
    @classmethod
    def to_proto(self, resource):
        if not resource:
            return resource
        return instance_pb2.Tier2AlphaInstancePreprocessCreateRecipeStepsRequestedTenantProjectScopeEnum.Value(
            "Tier2AlphaInstancePreprocessCreateRecipeStepsRequestedTenantProjectScopeEnum%s"
            % resource
        )

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return resource
        return instance_pb2.Tier2AlphaInstancePreprocessCreateRecipeStepsRequestedTenantProjectScopeEnum.Name(
            resource
        )[
            len(
                "Tier2AlphaInstancePreprocessCreateRecipeStepsRequestedTenantProjectScopeEnum"
            ) :
        ]


class InstancePreprocessCreateRecipeStepsPermissionsInfoPolicyNameModeEnum(object):
    @classmethod
    def to_proto(self, resource):
        if not resource:
            return resource
        return instance_pb2.Tier2AlphaInstancePreprocessCreateRecipeStepsPermissionsInfoPolicyNameModeEnum.Value(
            "Tier2AlphaInstancePreprocessCreateRecipeStepsPermissionsInfoPolicyNameModeEnum%s"
            % resource
        )

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return resource
        return instance_pb2.Tier2AlphaInstancePreprocessCreateRecipeStepsPermissionsInfoPolicyNameModeEnum.Name(
            resource
        )[
            len(
                "Tier2AlphaInstancePreprocessCreateRecipeStepsPermissionsInfoPolicyNameModeEnum"
            ) :
        ]


class InstanceCreateRecipeStepsActionEnum(object):
    @classmethod
    def to_proto(self, resource):
        if not resource:
            return resource
        return instance_pb2.Tier2AlphaInstanceCreateRecipeStepsActionEnum.Value(
            "Tier2AlphaInstanceCreateRecipeStepsActionEnum%s" % resource
        )

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return resource
        return instance_pb2.Tier2AlphaInstanceCreateRecipeStepsActionEnum.Name(
            resource
        )[len("Tier2AlphaInstanceCreateRecipeStepsActionEnum") :]


class InstanceCreateRecipeStepsRequestedTenantProjectScopeEnum(object):
    @classmethod
    def to_proto(self, resource):
        if not resource:
            return resource
        return instance_pb2.Tier2AlphaInstanceCreateRecipeStepsRequestedTenantProjectScopeEnum.Value(
            "Tier2AlphaInstanceCreateRecipeStepsRequestedTenantProjectScopeEnum%s"
            % resource
        )

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return resource
        return instance_pb2.Tier2AlphaInstanceCreateRecipeStepsRequestedTenantProjectScopeEnum.Name(
            resource
        )[
            len("Tier2AlphaInstanceCreateRecipeStepsRequestedTenantProjectScopeEnum") :
        ]


class InstanceCreateRecipeStepsPermissionsInfoPolicyNameModeEnum(object):
    @classmethod
    def to_proto(self, resource):
        if not resource:
            return resource
        return instance_pb2.Tier2AlphaInstanceCreateRecipeStepsPermissionsInfoPolicyNameModeEnum.Value(
            "Tier2AlphaInstanceCreateRecipeStepsPermissionsInfoPolicyNameModeEnum%s"
            % resource
        )

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return resource
        return instance_pb2.Tier2AlphaInstanceCreateRecipeStepsPermissionsInfoPolicyNameModeEnum.Name(
            resource
        )[
            len(
                "Tier2AlphaInstanceCreateRecipeStepsPermissionsInfoPolicyNameModeEnum"
            ) :
        ]


class InstanceDeleteRecipeStepsActionEnum(object):
    @classmethod
    def to_proto(self, resource):
        if not resource:
            return resource
        return instance_pb2.Tier2AlphaInstanceDeleteRecipeStepsActionEnum.Value(
            "Tier2AlphaInstanceDeleteRecipeStepsActionEnum%s" % resource
        )

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return resource
        return instance_pb2.Tier2AlphaInstanceDeleteRecipeStepsActionEnum.Name(
            resource
        )[len("Tier2AlphaInstanceDeleteRecipeStepsActionEnum") :]


class InstanceDeleteRecipeStepsRequestedTenantProjectScopeEnum(object):
    @classmethod
    def to_proto(self, resource):
        if not resource:
            return resource
        return instance_pb2.Tier2AlphaInstanceDeleteRecipeStepsRequestedTenantProjectScopeEnum.Value(
            "Tier2AlphaInstanceDeleteRecipeStepsRequestedTenantProjectScopeEnum%s"
            % resource
        )

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return resource
        return instance_pb2.Tier2AlphaInstanceDeleteRecipeStepsRequestedTenantProjectScopeEnum.Name(
            resource
        )[
            len("Tier2AlphaInstanceDeleteRecipeStepsRequestedTenantProjectScopeEnum") :
        ]


class InstanceDeleteRecipeStepsPermissionsInfoPolicyNameModeEnum(object):
    @classmethod
    def to_proto(self, resource):
        if not resource:
            return resource
        return instance_pb2.Tier2AlphaInstanceDeleteRecipeStepsPermissionsInfoPolicyNameModeEnum.Value(
            "Tier2AlphaInstanceDeleteRecipeStepsPermissionsInfoPolicyNameModeEnum%s"
            % resource
        )

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return resource
        return instance_pb2.Tier2AlphaInstanceDeleteRecipeStepsPermissionsInfoPolicyNameModeEnum.Name(
            resource
        )[
            len(
                "Tier2AlphaInstanceDeleteRecipeStepsPermissionsInfoPolicyNameModeEnum"
            ) :
        ]


class InstanceUpdateRecipeStepsActionEnum(object):
    @classmethod
    def to_proto(self, resource):
        if not resource:
            return resource
        return instance_pb2.Tier2AlphaInstanceUpdateRecipeStepsActionEnum.Value(
            "Tier2AlphaInstanceUpdateRecipeStepsActionEnum%s" % resource
        )

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return resource
        return instance_pb2.Tier2AlphaInstanceUpdateRecipeStepsActionEnum.Name(
            resource
        )[len("Tier2AlphaInstanceUpdateRecipeStepsActionEnum") :]


class InstanceUpdateRecipeStepsRequestedTenantProjectScopeEnum(object):
    @classmethod
    def to_proto(self, resource):
        if not resource:
            return resource
        return instance_pb2.Tier2AlphaInstanceUpdateRecipeStepsRequestedTenantProjectScopeEnum.Value(
            "Tier2AlphaInstanceUpdateRecipeStepsRequestedTenantProjectScopeEnum%s"
            % resource
        )

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return resource
        return instance_pb2.Tier2AlphaInstanceUpdateRecipeStepsRequestedTenantProjectScopeEnum.Name(
            resource
        )[
            len("Tier2AlphaInstanceUpdateRecipeStepsRequestedTenantProjectScopeEnum") :
        ]


class InstanceUpdateRecipeStepsPermissionsInfoPolicyNameModeEnum(object):
    @classmethod
    def to_proto(self, resource):
        if not resource:
            return resource
        return instance_pb2.Tier2AlphaInstanceUpdateRecipeStepsPermissionsInfoPolicyNameModeEnum.Value(
            "Tier2AlphaInstanceUpdateRecipeStepsPermissionsInfoPolicyNameModeEnum%s"
            % resource
        )

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return resource
        return instance_pb2.Tier2AlphaInstanceUpdateRecipeStepsPermissionsInfoPolicyNameModeEnum.Name(
            resource
        )[
            len(
                "Tier2AlphaInstanceUpdateRecipeStepsPermissionsInfoPolicyNameModeEnum"
            ) :
        ]


class InstancePreprocessResetRecipeStepsActionEnum(object):
    @classmethod
    def to_proto(self, resource):
        if not resource:
            return resource
        return instance_pb2.Tier2AlphaInstancePreprocessResetRecipeStepsActionEnum.Value(
            "Tier2AlphaInstancePreprocessResetRecipeStepsActionEnum%s" % resource
        )

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return resource
        return instance_pb2.Tier2AlphaInstancePreprocessResetRecipeStepsActionEnum.Name(
            resource
        )[len("Tier2AlphaInstancePreprocessResetRecipeStepsActionEnum") :]


class InstancePreprocessResetRecipeStepsRequestedTenantProjectScopeEnum(object):
    @classmethod
    def to_proto(self, resource):
        if not resource:
            return resource
        return instance_pb2.Tier2AlphaInstancePreprocessResetRecipeStepsRequestedTenantProjectScopeEnum.Value(
            "Tier2AlphaInstancePreprocessResetRecipeStepsRequestedTenantProjectScopeEnum%s"
            % resource
        )

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return resource
        return instance_pb2.Tier2AlphaInstancePreprocessResetRecipeStepsRequestedTenantProjectScopeEnum.Name(
            resource
        )[
            len(
                "Tier2AlphaInstancePreprocessResetRecipeStepsRequestedTenantProjectScopeEnum"
            ) :
        ]


class InstancePreprocessResetRecipeStepsPermissionsInfoPolicyNameModeEnum(object):
    @classmethod
    def to_proto(self, resource):
        if not resource:
            return resource
        return instance_pb2.Tier2AlphaInstancePreprocessResetRecipeStepsPermissionsInfoPolicyNameModeEnum.Value(
            "Tier2AlphaInstancePreprocessResetRecipeStepsPermissionsInfoPolicyNameModeEnum%s"
            % resource
        )

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return resource
        return instance_pb2.Tier2AlphaInstancePreprocessResetRecipeStepsPermissionsInfoPolicyNameModeEnum.Name(
            resource
        )[
            len(
                "Tier2AlphaInstancePreprocessResetRecipeStepsPermissionsInfoPolicyNameModeEnum"
            ) :
        ]


class InstanceResetRecipeStepsActionEnum(object):
    @classmethod
    def to_proto(self, resource):
        if not resource:
            return resource
        return instance_pb2.Tier2AlphaInstanceResetRecipeStepsActionEnum.Value(
            "Tier2AlphaInstanceResetRecipeStepsActionEnum%s" % resource
        )

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return resource
        return instance_pb2.Tier2AlphaInstanceResetRecipeStepsActionEnum.Name(resource)[
            len("Tier2AlphaInstanceResetRecipeStepsActionEnum") :
        ]


class InstanceResetRecipeStepsRequestedTenantProjectScopeEnum(object):
    @classmethod
    def to_proto(self, resource):
        if not resource:
            return resource
        return instance_pb2.Tier2AlphaInstanceResetRecipeStepsRequestedTenantProjectScopeEnum.Value(
            "Tier2AlphaInstanceResetRecipeStepsRequestedTenantProjectScopeEnum%s"
            % resource
        )

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return resource
        return instance_pb2.Tier2AlphaInstanceResetRecipeStepsRequestedTenantProjectScopeEnum.Name(
            resource
        )[
            len("Tier2AlphaInstanceResetRecipeStepsRequestedTenantProjectScopeEnum") :
        ]


class InstanceResetRecipeStepsPermissionsInfoPolicyNameModeEnum(object):
    @classmethod
    def to_proto(self, resource):
        if not resource:
            return resource
        return instance_pb2.Tier2AlphaInstanceResetRecipeStepsPermissionsInfoPolicyNameModeEnum.Value(
            "Tier2AlphaInstanceResetRecipeStepsPermissionsInfoPolicyNameModeEnum%s"
            % resource
        )

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return resource
        return instance_pb2.Tier2AlphaInstanceResetRecipeStepsPermissionsInfoPolicyNameModeEnum.Name(
            resource
        )[
            len("Tier2AlphaInstanceResetRecipeStepsPermissionsInfoPolicyNameModeEnum") :
        ]


class InstancePreprocessRepairRecipeStepsActionEnum(object):
    @classmethod
    def to_proto(self, resource):
        if not resource:
            return resource
        return instance_pb2.Tier2AlphaInstancePreprocessRepairRecipeStepsActionEnum.Value(
            "Tier2AlphaInstancePreprocessRepairRecipeStepsActionEnum%s" % resource
        )

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return resource
        return instance_pb2.Tier2AlphaInstancePreprocessRepairRecipeStepsActionEnum.Name(
            resource
        )[
            len("Tier2AlphaInstancePreprocessRepairRecipeStepsActionEnum") :
        ]


class InstancePreprocessRepairRecipeStepsRequestedTenantProjectScopeEnum(object):
    @classmethod
    def to_proto(self, resource):
        if not resource:
            return resource
        return instance_pb2.Tier2AlphaInstancePreprocessRepairRecipeStepsRequestedTenantProjectScopeEnum.Value(
            "Tier2AlphaInstancePreprocessRepairRecipeStepsRequestedTenantProjectScopeEnum%s"
            % resource
        )

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return resource
        return instance_pb2.Tier2AlphaInstancePreprocessRepairRecipeStepsRequestedTenantProjectScopeEnum.Name(
            resource
        )[
            len(
                "Tier2AlphaInstancePreprocessRepairRecipeStepsRequestedTenantProjectScopeEnum"
            ) :
        ]


class InstancePreprocessRepairRecipeStepsPermissionsInfoPolicyNameModeEnum(object):
    @classmethod
    def to_proto(self, resource):
        if not resource:
            return resource
        return instance_pb2.Tier2AlphaInstancePreprocessRepairRecipeStepsPermissionsInfoPolicyNameModeEnum.Value(
            "Tier2AlphaInstancePreprocessRepairRecipeStepsPermissionsInfoPolicyNameModeEnum%s"
            % resource
        )

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return resource
        return instance_pb2.Tier2AlphaInstancePreprocessRepairRecipeStepsPermissionsInfoPolicyNameModeEnum.Name(
            resource
        )[
            len(
                "Tier2AlphaInstancePreprocessRepairRecipeStepsPermissionsInfoPolicyNameModeEnum"
            ) :
        ]


class InstanceRepairRecipeStepsActionEnum(object):
    @classmethod
    def to_proto(self, resource):
        if not resource:
            return resource
        return instance_pb2.Tier2AlphaInstanceRepairRecipeStepsActionEnum.Value(
            "Tier2AlphaInstanceRepairRecipeStepsActionEnum%s" % resource
        )

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return resource
        return instance_pb2.Tier2AlphaInstanceRepairRecipeStepsActionEnum.Name(
            resource
        )[len("Tier2AlphaInstanceRepairRecipeStepsActionEnum") :]


class InstanceRepairRecipeStepsRequestedTenantProjectScopeEnum(object):
    @classmethod
    def to_proto(self, resource):
        if not resource:
            return resource
        return instance_pb2.Tier2AlphaInstanceRepairRecipeStepsRequestedTenantProjectScopeEnum.Value(
            "Tier2AlphaInstanceRepairRecipeStepsRequestedTenantProjectScopeEnum%s"
            % resource
        )

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return resource
        return instance_pb2.Tier2AlphaInstanceRepairRecipeStepsRequestedTenantProjectScopeEnum.Name(
            resource
        )[
            len("Tier2AlphaInstanceRepairRecipeStepsRequestedTenantProjectScopeEnum") :
        ]


class InstanceRepairRecipeStepsPermissionsInfoPolicyNameModeEnum(object):
    @classmethod
    def to_proto(self, resource):
        if not resource:
            return resource
        return instance_pb2.Tier2AlphaInstanceRepairRecipeStepsPermissionsInfoPolicyNameModeEnum.Value(
            "Tier2AlphaInstanceRepairRecipeStepsPermissionsInfoPolicyNameModeEnum%s"
            % resource
        )

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return resource
        return instance_pb2.Tier2AlphaInstanceRepairRecipeStepsPermissionsInfoPolicyNameModeEnum.Name(
            resource
        )[
            len(
                "Tier2AlphaInstanceRepairRecipeStepsPermissionsInfoPolicyNameModeEnum"
            ) :
        ]


class InstancePreprocessDeleteRecipeStepsActionEnum(object):
    @classmethod
    def to_proto(self, resource):
        if not resource:
            return resource
        return instance_pb2.Tier2AlphaInstancePreprocessDeleteRecipeStepsActionEnum.Value(
            "Tier2AlphaInstancePreprocessDeleteRecipeStepsActionEnum%s" % resource
        )

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return resource
        return instance_pb2.Tier2AlphaInstancePreprocessDeleteRecipeStepsActionEnum.Name(
            resource
        )[
            len("Tier2AlphaInstancePreprocessDeleteRecipeStepsActionEnum") :
        ]


class InstancePreprocessDeleteRecipeStepsRequestedTenantProjectScopeEnum(object):
    @classmethod
    def to_proto(self, resource):
        if not resource:
            return resource
        return instance_pb2.Tier2AlphaInstancePreprocessDeleteRecipeStepsRequestedTenantProjectScopeEnum.Value(
            "Tier2AlphaInstancePreprocessDeleteRecipeStepsRequestedTenantProjectScopeEnum%s"
            % resource
        )

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return resource
        return instance_pb2.Tier2AlphaInstancePreprocessDeleteRecipeStepsRequestedTenantProjectScopeEnum.Name(
            resource
        )[
            len(
                "Tier2AlphaInstancePreprocessDeleteRecipeStepsRequestedTenantProjectScopeEnum"
            ) :
        ]


class InstancePreprocessDeleteRecipeStepsPermissionsInfoPolicyNameModeEnum(object):
    @classmethod
    def to_proto(self, resource):
        if not resource:
            return resource
        return instance_pb2.Tier2AlphaInstancePreprocessDeleteRecipeStepsPermissionsInfoPolicyNameModeEnum.Value(
            "Tier2AlphaInstancePreprocessDeleteRecipeStepsPermissionsInfoPolicyNameModeEnum%s"
            % resource
        )

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return resource
        return instance_pb2.Tier2AlphaInstancePreprocessDeleteRecipeStepsPermissionsInfoPolicyNameModeEnum.Name(
            resource
        )[
            len(
                "Tier2AlphaInstancePreprocessDeleteRecipeStepsPermissionsInfoPolicyNameModeEnum"
            ) :
        ]


class InstancePreprocessUpdateRecipeStepsActionEnum(object):
    @classmethod
    def to_proto(self, resource):
        if not resource:
            return resource
        return instance_pb2.Tier2AlphaInstancePreprocessUpdateRecipeStepsActionEnum.Value(
            "Tier2AlphaInstancePreprocessUpdateRecipeStepsActionEnum%s" % resource
        )

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return resource
        return instance_pb2.Tier2AlphaInstancePreprocessUpdateRecipeStepsActionEnum.Name(
            resource
        )[
            len("Tier2AlphaInstancePreprocessUpdateRecipeStepsActionEnum") :
        ]


class InstancePreprocessUpdateRecipeStepsRequestedTenantProjectScopeEnum(object):
    @classmethod
    def to_proto(self, resource):
        if not resource:
            return resource
        return instance_pb2.Tier2AlphaInstancePreprocessUpdateRecipeStepsRequestedTenantProjectScopeEnum.Value(
            "Tier2AlphaInstancePreprocessUpdateRecipeStepsRequestedTenantProjectScopeEnum%s"
            % resource
        )

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return resource
        return instance_pb2.Tier2AlphaInstancePreprocessUpdateRecipeStepsRequestedTenantProjectScopeEnum.Name(
            resource
        )[
            len(
                "Tier2AlphaInstancePreprocessUpdateRecipeStepsRequestedTenantProjectScopeEnum"
            ) :
        ]


class InstancePreprocessUpdateRecipeStepsPermissionsInfoPolicyNameModeEnum(object):
    @classmethod
    def to_proto(self, resource):
        if not resource:
            return resource
        return instance_pb2.Tier2AlphaInstancePreprocessUpdateRecipeStepsPermissionsInfoPolicyNameModeEnum.Value(
            "Tier2AlphaInstancePreprocessUpdateRecipeStepsPermissionsInfoPolicyNameModeEnum%s"
            % resource
        )

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return resource
        return instance_pb2.Tier2AlphaInstancePreprocessUpdateRecipeStepsPermissionsInfoPolicyNameModeEnum.Name(
            resource
        )[
            len(
                "Tier2AlphaInstancePreprocessUpdateRecipeStepsPermissionsInfoPolicyNameModeEnum"
            ) :
        ]


class InstancePreprocessFreezeRecipeStepsActionEnum(object):
    @classmethod
    def to_proto(self, resource):
        if not resource:
            return resource
        return instance_pb2.Tier2AlphaInstancePreprocessFreezeRecipeStepsActionEnum.Value(
            "Tier2AlphaInstancePreprocessFreezeRecipeStepsActionEnum%s" % resource
        )

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return resource
        return instance_pb2.Tier2AlphaInstancePreprocessFreezeRecipeStepsActionEnum.Name(
            resource
        )[
            len("Tier2AlphaInstancePreprocessFreezeRecipeStepsActionEnum") :
        ]


class InstancePreprocessFreezeRecipeStepsRequestedTenantProjectScopeEnum(object):
    @classmethod
    def to_proto(self, resource):
        if not resource:
            return resource
        return instance_pb2.Tier2AlphaInstancePreprocessFreezeRecipeStepsRequestedTenantProjectScopeEnum.Value(
            "Tier2AlphaInstancePreprocessFreezeRecipeStepsRequestedTenantProjectScopeEnum%s"
            % resource
        )

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return resource
        return instance_pb2.Tier2AlphaInstancePreprocessFreezeRecipeStepsRequestedTenantProjectScopeEnum.Name(
            resource
        )[
            len(
                "Tier2AlphaInstancePreprocessFreezeRecipeStepsRequestedTenantProjectScopeEnum"
            ) :
        ]


class InstancePreprocessFreezeRecipeStepsPermissionsInfoPolicyNameModeEnum(object):
    @classmethod
    def to_proto(self, resource):
        if not resource:
            return resource
        return instance_pb2.Tier2AlphaInstancePreprocessFreezeRecipeStepsPermissionsInfoPolicyNameModeEnum.Value(
            "Tier2AlphaInstancePreprocessFreezeRecipeStepsPermissionsInfoPolicyNameModeEnum%s"
            % resource
        )

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return resource
        return instance_pb2.Tier2AlphaInstancePreprocessFreezeRecipeStepsPermissionsInfoPolicyNameModeEnum.Name(
            resource
        )[
            len(
                "Tier2AlphaInstancePreprocessFreezeRecipeStepsPermissionsInfoPolicyNameModeEnum"
            ) :
        ]


class InstanceFreezeRecipeStepsActionEnum(object):
    @classmethod
    def to_proto(self, resource):
        if not resource:
            return resource
        return instance_pb2.Tier2AlphaInstanceFreezeRecipeStepsActionEnum.Value(
            "Tier2AlphaInstanceFreezeRecipeStepsActionEnum%s" % resource
        )

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return resource
        return instance_pb2.Tier2AlphaInstanceFreezeRecipeStepsActionEnum.Name(
            resource
        )[len("Tier2AlphaInstanceFreezeRecipeStepsActionEnum") :]


class InstanceFreezeRecipeStepsRequestedTenantProjectScopeEnum(object):
    @classmethod
    def to_proto(self, resource):
        if not resource:
            return resource
        return instance_pb2.Tier2AlphaInstanceFreezeRecipeStepsRequestedTenantProjectScopeEnum.Value(
            "Tier2AlphaInstanceFreezeRecipeStepsRequestedTenantProjectScopeEnum%s"
            % resource
        )

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return resource
        return instance_pb2.Tier2AlphaInstanceFreezeRecipeStepsRequestedTenantProjectScopeEnum.Name(
            resource
        )[
            len("Tier2AlphaInstanceFreezeRecipeStepsRequestedTenantProjectScopeEnum") :
        ]


class InstanceFreezeRecipeStepsPermissionsInfoPolicyNameModeEnum(object):
    @classmethod
    def to_proto(self, resource):
        if not resource:
            return resource
        return instance_pb2.Tier2AlphaInstanceFreezeRecipeStepsPermissionsInfoPolicyNameModeEnum.Value(
            "Tier2AlphaInstanceFreezeRecipeStepsPermissionsInfoPolicyNameModeEnum%s"
            % resource
        )

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return resource
        return instance_pb2.Tier2AlphaInstanceFreezeRecipeStepsPermissionsInfoPolicyNameModeEnum.Name(
            resource
        )[
            len(
                "Tier2AlphaInstanceFreezeRecipeStepsPermissionsInfoPolicyNameModeEnum"
            ) :
        ]


class InstancePreprocessUnfreezeRecipeStepsActionEnum(object):
    @classmethod
    def to_proto(self, resource):
        if not resource:
            return resource
        return instance_pb2.Tier2AlphaInstancePreprocessUnfreezeRecipeStepsActionEnum.Value(
            "Tier2AlphaInstancePreprocessUnfreezeRecipeStepsActionEnum%s" % resource
        )

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return resource
        return instance_pb2.Tier2AlphaInstancePreprocessUnfreezeRecipeStepsActionEnum.Name(
            resource
        )[
            len("Tier2AlphaInstancePreprocessUnfreezeRecipeStepsActionEnum") :
        ]


class InstancePreprocessUnfreezeRecipeStepsRequestedTenantProjectScopeEnum(object):
    @classmethod
    def to_proto(self, resource):
        if not resource:
            return resource
        return instance_pb2.Tier2AlphaInstancePreprocessUnfreezeRecipeStepsRequestedTenantProjectScopeEnum.Value(
            "Tier2AlphaInstancePreprocessUnfreezeRecipeStepsRequestedTenantProjectScopeEnum%s"
            % resource
        )

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return resource
        return instance_pb2.Tier2AlphaInstancePreprocessUnfreezeRecipeStepsRequestedTenantProjectScopeEnum.Name(
            resource
        )[
            len(
                "Tier2AlphaInstancePreprocessUnfreezeRecipeStepsRequestedTenantProjectScopeEnum"
            ) :
        ]


class InstancePreprocessUnfreezeRecipeStepsPermissionsInfoPolicyNameModeEnum(object):
    @classmethod
    def to_proto(self, resource):
        if not resource:
            return resource
        return instance_pb2.Tier2AlphaInstancePreprocessUnfreezeRecipeStepsPermissionsInfoPolicyNameModeEnum.Value(
            "Tier2AlphaInstancePreprocessUnfreezeRecipeStepsPermissionsInfoPolicyNameModeEnum%s"
            % resource
        )

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return resource
        return instance_pb2.Tier2AlphaInstancePreprocessUnfreezeRecipeStepsPermissionsInfoPolicyNameModeEnum.Name(
            resource
        )[
            len(
                "Tier2AlphaInstancePreprocessUnfreezeRecipeStepsPermissionsInfoPolicyNameModeEnum"
            ) :
        ]


class InstanceUnfreezeRecipeStepsActionEnum(object):
    @classmethod
    def to_proto(self, resource):
        if not resource:
            return resource
        return instance_pb2.Tier2AlphaInstanceUnfreezeRecipeStepsActionEnum.Value(
            "Tier2AlphaInstanceUnfreezeRecipeStepsActionEnum%s" % resource
        )

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return resource
        return instance_pb2.Tier2AlphaInstanceUnfreezeRecipeStepsActionEnum.Name(
            resource
        )[len("Tier2AlphaInstanceUnfreezeRecipeStepsActionEnum") :]


class InstanceUnfreezeRecipeStepsRequestedTenantProjectScopeEnum(object):
    @classmethod
    def to_proto(self, resource):
        if not resource:
            return resource
        return instance_pb2.Tier2AlphaInstanceUnfreezeRecipeStepsRequestedTenantProjectScopeEnum.Value(
            "Tier2AlphaInstanceUnfreezeRecipeStepsRequestedTenantProjectScopeEnum%s"
            % resource
        )

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return resource
        return instance_pb2.Tier2AlphaInstanceUnfreezeRecipeStepsRequestedTenantProjectScopeEnum.Name(
            resource
        )[
            len(
                "Tier2AlphaInstanceUnfreezeRecipeStepsRequestedTenantProjectScopeEnum"
            ) :
        ]


class InstanceUnfreezeRecipeStepsPermissionsInfoPolicyNameModeEnum(object):
    @classmethod
    def to_proto(self, resource):
        if not resource:
            return resource
        return instance_pb2.Tier2AlphaInstanceUnfreezeRecipeStepsPermissionsInfoPolicyNameModeEnum.Value(
            "Tier2AlphaInstanceUnfreezeRecipeStepsPermissionsInfoPolicyNameModeEnum%s"
            % resource
        )

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return resource
        return instance_pb2.Tier2AlphaInstanceUnfreezeRecipeStepsPermissionsInfoPolicyNameModeEnum.Name(
            resource
        )[
            len(
                "Tier2AlphaInstanceUnfreezeRecipeStepsPermissionsInfoPolicyNameModeEnum"
            ) :
        ]


class InstancePreprocessReportInstanceHealthRecipeStepsActionEnum(object):
    @classmethod
    def to_proto(self, resource):
        if not resource:
            return resource
        return instance_pb2.Tier2AlphaInstancePreprocessReportInstanceHealthRecipeStepsActionEnum.Value(
            "Tier2AlphaInstancePreprocessReportInstanceHealthRecipeStepsActionEnum%s"
            % resource
        )

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return resource
        return instance_pb2.Tier2AlphaInstancePreprocessReportInstanceHealthRecipeStepsActionEnum.Name(
            resource
        )[
            len(
                "Tier2AlphaInstancePreprocessReportInstanceHealthRecipeStepsActionEnum"
            ) :
        ]


class InstancePreprocessReportInstanceHealthRecipeStepsRequestedTenantProjectScopeEnum(
    object
):
    @classmethod
    def to_proto(self, resource):
        if not resource:
            return resource
        return instance_pb2.Tier2AlphaInstancePreprocessReportInstanceHealthRecipeStepsRequestedTenantProjectScopeEnum.Value(
            "Tier2AlphaInstancePreprocessReportInstanceHealthRecipeStepsRequestedTenantProjectScopeEnum%s"
            % resource
        )

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return resource
        return instance_pb2.Tier2AlphaInstancePreprocessReportInstanceHealthRecipeStepsRequestedTenantProjectScopeEnum.Name(
            resource
        )[
            len(
                "Tier2AlphaInstancePreprocessReportInstanceHealthRecipeStepsRequestedTenantProjectScopeEnum"
            ) :
        ]


class InstancePreprocessReportInstanceHealthRecipeStepsPermissionsInfoPolicyNameModeEnum(
    object
):
    @classmethod
    def to_proto(self, resource):
        if not resource:
            return resource
        return instance_pb2.Tier2AlphaInstancePreprocessReportInstanceHealthRecipeStepsPermissionsInfoPolicyNameModeEnum.Value(
            "Tier2AlphaInstancePreprocessReportInstanceHealthRecipeStepsPermissionsInfoPolicyNameModeEnum%s"
            % resource
        )

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return resource
        return instance_pb2.Tier2AlphaInstancePreprocessReportInstanceHealthRecipeStepsPermissionsInfoPolicyNameModeEnum.Name(
            resource
        )[
            len(
                "Tier2AlphaInstancePreprocessReportInstanceHealthRecipeStepsPermissionsInfoPolicyNameModeEnum"
            ) :
        ]


class InstanceReportInstanceHealthRecipeStepsActionEnum(object):
    @classmethod
    def to_proto(self, resource):
        if not resource:
            return resource
        return instance_pb2.Tier2AlphaInstanceReportInstanceHealthRecipeStepsActionEnum.Value(
            "Tier2AlphaInstanceReportInstanceHealthRecipeStepsActionEnum%s" % resource
        )

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return resource
        return instance_pb2.Tier2AlphaInstanceReportInstanceHealthRecipeStepsActionEnum.Name(
            resource
        )[
            len("Tier2AlphaInstanceReportInstanceHealthRecipeStepsActionEnum") :
        ]


class InstanceReportInstanceHealthRecipeStepsRequestedTenantProjectScopeEnum(object):
    @classmethod
    def to_proto(self, resource):
        if not resource:
            return resource
        return instance_pb2.Tier2AlphaInstanceReportInstanceHealthRecipeStepsRequestedTenantProjectScopeEnum.Value(
            "Tier2AlphaInstanceReportInstanceHealthRecipeStepsRequestedTenantProjectScopeEnum%s"
            % resource
        )

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return resource
        return instance_pb2.Tier2AlphaInstanceReportInstanceHealthRecipeStepsRequestedTenantProjectScopeEnum.Name(
            resource
        )[
            len(
                "Tier2AlphaInstanceReportInstanceHealthRecipeStepsRequestedTenantProjectScopeEnum"
            ) :
        ]


class InstanceReportInstanceHealthRecipeStepsPermissionsInfoPolicyNameModeEnum(object):
    @classmethod
    def to_proto(self, resource):
        if not resource:
            return resource
        return instance_pb2.Tier2AlphaInstanceReportInstanceHealthRecipeStepsPermissionsInfoPolicyNameModeEnum.Value(
            "Tier2AlphaInstanceReportInstanceHealthRecipeStepsPermissionsInfoPolicyNameModeEnum%s"
            % resource
        )

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return resource
        return instance_pb2.Tier2AlphaInstanceReportInstanceHealthRecipeStepsPermissionsInfoPolicyNameModeEnum.Name(
            resource
        )[
            len(
                "Tier2AlphaInstanceReportInstanceHealthRecipeStepsPermissionsInfoPolicyNameModeEnum"
            ) :
        ]


class InstancePreprocessGetRecipeStepsActionEnum(object):
    @classmethod
    def to_proto(self, resource):
        if not resource:
            return resource
        return instance_pb2.Tier2AlphaInstancePreprocessGetRecipeStepsActionEnum.Value(
            "Tier2AlphaInstancePreprocessGetRecipeStepsActionEnum%s" % resource
        )

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return resource
        return instance_pb2.Tier2AlphaInstancePreprocessGetRecipeStepsActionEnum.Name(
            resource
        )[len("Tier2AlphaInstancePreprocessGetRecipeStepsActionEnum") :]


class InstancePreprocessGetRecipeStepsRequestedTenantProjectScopeEnum(object):
    @classmethod
    def to_proto(self, resource):
        if not resource:
            return resource
        return instance_pb2.Tier2AlphaInstancePreprocessGetRecipeStepsRequestedTenantProjectScopeEnum.Value(
            "Tier2AlphaInstancePreprocessGetRecipeStepsRequestedTenantProjectScopeEnum%s"
            % resource
        )

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return resource
        return instance_pb2.Tier2AlphaInstancePreprocessGetRecipeStepsRequestedTenantProjectScopeEnum.Name(
            resource
        )[
            len(
                "Tier2AlphaInstancePreprocessGetRecipeStepsRequestedTenantProjectScopeEnum"
            ) :
        ]


class InstancePreprocessGetRecipeStepsPermissionsInfoPolicyNameModeEnum(object):
    @classmethod
    def to_proto(self, resource):
        if not resource:
            return resource
        return instance_pb2.Tier2AlphaInstancePreprocessGetRecipeStepsPermissionsInfoPolicyNameModeEnum.Value(
            "Tier2AlphaInstancePreprocessGetRecipeStepsPermissionsInfoPolicyNameModeEnum%s"
            % resource
        )

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return resource
        return instance_pb2.Tier2AlphaInstancePreprocessGetRecipeStepsPermissionsInfoPolicyNameModeEnum.Name(
            resource
        )[
            len(
                "Tier2AlphaInstancePreprocessGetRecipeStepsPermissionsInfoPolicyNameModeEnum"
            ) :
        ]


class InstanceNotifyKeyAvailableRecipeStepsActionEnum(object):
    @classmethod
    def to_proto(self, resource):
        if not resource:
            return resource
        return instance_pb2.Tier2AlphaInstanceNotifyKeyAvailableRecipeStepsActionEnum.Value(
            "Tier2AlphaInstanceNotifyKeyAvailableRecipeStepsActionEnum%s" % resource
        )

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return resource
        return instance_pb2.Tier2AlphaInstanceNotifyKeyAvailableRecipeStepsActionEnum.Name(
            resource
        )[
            len("Tier2AlphaInstanceNotifyKeyAvailableRecipeStepsActionEnum") :
        ]


class InstanceNotifyKeyAvailableRecipeStepsRequestedTenantProjectScopeEnum(object):
    @classmethod
    def to_proto(self, resource):
        if not resource:
            return resource
        return instance_pb2.Tier2AlphaInstanceNotifyKeyAvailableRecipeStepsRequestedTenantProjectScopeEnum.Value(
            "Tier2AlphaInstanceNotifyKeyAvailableRecipeStepsRequestedTenantProjectScopeEnum%s"
            % resource
        )

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return resource
        return instance_pb2.Tier2AlphaInstanceNotifyKeyAvailableRecipeStepsRequestedTenantProjectScopeEnum.Name(
            resource
        )[
            len(
                "Tier2AlphaInstanceNotifyKeyAvailableRecipeStepsRequestedTenantProjectScopeEnum"
            ) :
        ]


class InstanceNotifyKeyAvailableRecipeStepsPermissionsInfoPolicyNameModeEnum(object):
    @classmethod
    def to_proto(self, resource):
        if not resource:
            return resource
        return instance_pb2.Tier2AlphaInstanceNotifyKeyAvailableRecipeStepsPermissionsInfoPolicyNameModeEnum.Value(
            "Tier2AlphaInstanceNotifyKeyAvailableRecipeStepsPermissionsInfoPolicyNameModeEnum%s"
            % resource
        )

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return resource
        return instance_pb2.Tier2AlphaInstanceNotifyKeyAvailableRecipeStepsPermissionsInfoPolicyNameModeEnum.Name(
            resource
        )[
            len(
                "Tier2AlphaInstanceNotifyKeyAvailableRecipeStepsPermissionsInfoPolicyNameModeEnum"
            ) :
        ]


class InstanceNotifyKeyUnavailableRecipeStepsActionEnum(object):
    @classmethod
    def to_proto(self, resource):
        if not resource:
            return resource
        return instance_pb2.Tier2AlphaInstanceNotifyKeyUnavailableRecipeStepsActionEnum.Value(
            "Tier2AlphaInstanceNotifyKeyUnavailableRecipeStepsActionEnum%s" % resource
        )

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return resource
        return instance_pb2.Tier2AlphaInstanceNotifyKeyUnavailableRecipeStepsActionEnum.Name(
            resource
        )[
            len("Tier2AlphaInstanceNotifyKeyUnavailableRecipeStepsActionEnum") :
        ]


class InstanceNotifyKeyUnavailableRecipeStepsRequestedTenantProjectScopeEnum(object):
    @classmethod
    def to_proto(self, resource):
        if not resource:
            return resource
        return instance_pb2.Tier2AlphaInstanceNotifyKeyUnavailableRecipeStepsRequestedTenantProjectScopeEnum.Value(
            "Tier2AlphaInstanceNotifyKeyUnavailableRecipeStepsRequestedTenantProjectScopeEnum%s"
            % resource
        )

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return resource
        return instance_pb2.Tier2AlphaInstanceNotifyKeyUnavailableRecipeStepsRequestedTenantProjectScopeEnum.Name(
            resource
        )[
            len(
                "Tier2AlphaInstanceNotifyKeyUnavailableRecipeStepsRequestedTenantProjectScopeEnum"
            ) :
        ]


class InstanceNotifyKeyUnavailableRecipeStepsPermissionsInfoPolicyNameModeEnum(object):
    @classmethod
    def to_proto(self, resource):
        if not resource:
            return resource
        return instance_pb2.Tier2AlphaInstanceNotifyKeyUnavailableRecipeStepsPermissionsInfoPolicyNameModeEnum.Value(
            "Tier2AlphaInstanceNotifyKeyUnavailableRecipeStepsPermissionsInfoPolicyNameModeEnum%s"
            % resource
        )

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return resource
        return instance_pb2.Tier2AlphaInstanceNotifyKeyUnavailableRecipeStepsPermissionsInfoPolicyNameModeEnum.Name(
            resource
        )[
            len(
                "Tier2AlphaInstanceNotifyKeyUnavailableRecipeStepsPermissionsInfoPolicyNameModeEnum"
            ) :
        ]


class InstanceReadonlyRecipeStepsActionEnum(object):
    @classmethod
    def to_proto(self, resource):
        if not resource:
            return resource
        return instance_pb2.Tier2AlphaInstanceReadonlyRecipeStepsActionEnum.Value(
            "Tier2AlphaInstanceReadonlyRecipeStepsActionEnum%s" % resource
        )

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return resource
        return instance_pb2.Tier2AlphaInstanceReadonlyRecipeStepsActionEnum.Name(
            resource
        )[len("Tier2AlphaInstanceReadonlyRecipeStepsActionEnum") :]


class InstanceReadonlyRecipeStepsRequestedTenantProjectScopeEnum(object):
    @classmethod
    def to_proto(self, resource):
        if not resource:
            return resource
        return instance_pb2.Tier2AlphaInstanceReadonlyRecipeStepsRequestedTenantProjectScopeEnum.Value(
            "Tier2AlphaInstanceReadonlyRecipeStepsRequestedTenantProjectScopeEnum%s"
            % resource
        )

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return resource
        return instance_pb2.Tier2AlphaInstanceReadonlyRecipeStepsRequestedTenantProjectScopeEnum.Name(
            resource
        )[
            len(
                "Tier2AlphaInstanceReadonlyRecipeStepsRequestedTenantProjectScopeEnum"
            ) :
        ]


class InstanceReadonlyRecipeStepsPermissionsInfoPolicyNameModeEnum(object):
    @classmethod
    def to_proto(self, resource):
        if not resource:
            return resource
        return instance_pb2.Tier2AlphaInstanceReadonlyRecipeStepsPermissionsInfoPolicyNameModeEnum.Value(
            "Tier2AlphaInstanceReadonlyRecipeStepsPermissionsInfoPolicyNameModeEnum%s"
            % resource
        )

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return resource
        return instance_pb2.Tier2AlphaInstanceReadonlyRecipeStepsPermissionsInfoPolicyNameModeEnum.Name(
            resource
        )[
            len(
                "Tier2AlphaInstanceReadonlyRecipeStepsPermissionsInfoPolicyNameModeEnum"
            ) :
        ]


class InstanceReconcileRecipeStepsActionEnum(object):
    @classmethod
    def to_proto(self, resource):
        if not resource:
            return resource
        return instance_pb2.Tier2AlphaInstanceReconcileRecipeStepsActionEnum.Value(
            "Tier2AlphaInstanceReconcileRecipeStepsActionEnum%s" % resource
        )

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return resource
        return instance_pb2.Tier2AlphaInstanceReconcileRecipeStepsActionEnum.Name(
            resource
        )[len("Tier2AlphaInstanceReconcileRecipeStepsActionEnum") :]


class InstanceReconcileRecipeStepsRequestedTenantProjectScopeEnum(object):
    @classmethod
    def to_proto(self, resource):
        if not resource:
            return resource
        return instance_pb2.Tier2AlphaInstanceReconcileRecipeStepsRequestedTenantProjectScopeEnum.Value(
            "Tier2AlphaInstanceReconcileRecipeStepsRequestedTenantProjectScopeEnum%s"
            % resource
        )

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return resource
        return instance_pb2.Tier2AlphaInstanceReconcileRecipeStepsRequestedTenantProjectScopeEnum.Name(
            resource
        )[
            len(
                "Tier2AlphaInstanceReconcileRecipeStepsRequestedTenantProjectScopeEnum"
            ) :
        ]


class InstanceReconcileRecipeStepsPermissionsInfoPolicyNameModeEnum(object):
    @classmethod
    def to_proto(self, resource):
        if not resource:
            return resource
        return instance_pb2.Tier2AlphaInstanceReconcileRecipeStepsPermissionsInfoPolicyNameModeEnum.Value(
            "Tier2AlphaInstanceReconcileRecipeStepsPermissionsInfoPolicyNameModeEnum%s"
            % resource
        )

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return resource
        return instance_pb2.Tier2AlphaInstanceReconcileRecipeStepsPermissionsInfoPolicyNameModeEnum.Name(
            resource
        )[
            len(
                "Tier2AlphaInstanceReconcileRecipeStepsPermissionsInfoPolicyNameModeEnum"
            ) :
        ]


class InstancePreprocessPassthroughRecipeStepsActionEnum(object):
    @classmethod
    def to_proto(self, resource):
        if not resource:
            return resource
        return instance_pb2.Tier2AlphaInstancePreprocessPassthroughRecipeStepsActionEnum.Value(
            "Tier2AlphaInstancePreprocessPassthroughRecipeStepsActionEnum%s" % resource
        )

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return resource
        return instance_pb2.Tier2AlphaInstancePreprocessPassthroughRecipeStepsActionEnum.Name(
            resource
        )[
            len("Tier2AlphaInstancePreprocessPassthroughRecipeStepsActionEnum") :
        ]


class InstancePreprocessPassthroughRecipeStepsRequestedTenantProjectScopeEnum(object):
    @classmethod
    def to_proto(self, resource):
        if not resource:
            return resource
        return instance_pb2.Tier2AlphaInstancePreprocessPassthroughRecipeStepsRequestedTenantProjectScopeEnum.Value(
            "Tier2AlphaInstancePreprocessPassthroughRecipeStepsRequestedTenantProjectScopeEnum%s"
            % resource
        )

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return resource
        return instance_pb2.Tier2AlphaInstancePreprocessPassthroughRecipeStepsRequestedTenantProjectScopeEnum.Name(
            resource
        )[
            len(
                "Tier2AlphaInstancePreprocessPassthroughRecipeStepsRequestedTenantProjectScopeEnum"
            ) :
        ]


class InstancePreprocessPassthroughRecipeStepsPermissionsInfoPolicyNameModeEnum(object):
    @classmethod
    def to_proto(self, resource):
        if not resource:
            return resource
        return instance_pb2.Tier2AlphaInstancePreprocessPassthroughRecipeStepsPermissionsInfoPolicyNameModeEnum.Value(
            "Tier2AlphaInstancePreprocessPassthroughRecipeStepsPermissionsInfoPolicyNameModeEnum%s"
            % resource
        )

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return resource
        return instance_pb2.Tier2AlphaInstancePreprocessPassthroughRecipeStepsPermissionsInfoPolicyNameModeEnum.Name(
            resource
        )[
            len(
                "Tier2AlphaInstancePreprocessPassthroughRecipeStepsPermissionsInfoPolicyNameModeEnum"
            ) :
        ]


class InstancePreprocessReconcileRecipeStepsActionEnum(object):
    @classmethod
    def to_proto(self, resource):
        if not resource:
            return resource
        return instance_pb2.Tier2AlphaInstancePreprocessReconcileRecipeStepsActionEnum.Value(
            "Tier2AlphaInstancePreprocessReconcileRecipeStepsActionEnum%s" % resource
        )

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return resource
        return instance_pb2.Tier2AlphaInstancePreprocessReconcileRecipeStepsActionEnum.Name(
            resource
        )[
            len("Tier2AlphaInstancePreprocessReconcileRecipeStepsActionEnum") :
        ]


class InstancePreprocessReconcileRecipeStepsRequestedTenantProjectScopeEnum(object):
    @classmethod
    def to_proto(self, resource):
        if not resource:
            return resource
        return instance_pb2.Tier2AlphaInstancePreprocessReconcileRecipeStepsRequestedTenantProjectScopeEnum.Value(
            "Tier2AlphaInstancePreprocessReconcileRecipeStepsRequestedTenantProjectScopeEnum%s"
            % resource
        )

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return resource
        return instance_pb2.Tier2AlphaInstancePreprocessReconcileRecipeStepsRequestedTenantProjectScopeEnum.Name(
            resource
        )[
            len(
                "Tier2AlphaInstancePreprocessReconcileRecipeStepsRequestedTenantProjectScopeEnum"
            ) :
        ]


class InstancePreprocessReconcileRecipeStepsPermissionsInfoPolicyNameModeEnum(object):
    @classmethod
    def to_proto(self, resource):
        if not resource:
            return resource
        return instance_pb2.Tier2AlphaInstancePreprocessReconcileRecipeStepsPermissionsInfoPolicyNameModeEnum.Value(
            "Tier2AlphaInstancePreprocessReconcileRecipeStepsPermissionsInfoPolicyNameModeEnum%s"
            % resource
        )

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return resource
        return instance_pb2.Tier2AlphaInstancePreprocessReconcileRecipeStepsPermissionsInfoPolicyNameModeEnum.Name(
            resource
        )[
            len(
                "Tier2AlphaInstancePreprocessReconcileRecipeStepsPermissionsInfoPolicyNameModeEnum"
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
