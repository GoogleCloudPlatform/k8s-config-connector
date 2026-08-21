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
from google3.cloud.graphite.mmv2.services.google.bigquery import table_pb2
from google3.cloud.graphite.mmv2.services.google.bigquery import table_pb2_grpc

from typing import List


class Table(object):
    def __init__(
        self,
        etag: str = None,
        id: str = None,
        self_link: str = None,
        name: str = None,
        dataset: str = None,
        project: str = None,
        friendly_name: str = None,
        description: str = None,
        labels: dict = None,
        model: dict = None,
        schema: dict = None,
        time_partitioning: dict = None,
        range_partitioning: dict = None,
        clustering: dict = None,
        require_partition_filter: bool = None,
        num_bytes: str = None,
        num_physical_bytes: str = None,
        num_long_term_bytes: str = None,
        num_rows: int = None,
        creation_time: int = None,
        expiration_time: int = None,
        last_modified_time: int = None,
        type: str = None,
        view: dict = None,
        materialized_view: dict = None,
        external_data_configuration: dict = None,
        location: str = None,
        streaming_buffer: dict = None,
        encryption_configuration: dict = None,
        snapshot_definition: dict = None,
        default_collation: str = None,
        service_account_file: str = "",
    ):

        channel.initialize()
        self.name = name
        self.dataset = dataset
        self.project = project
        self.friendly_name = friendly_name
        self.description = description
        self.labels = labels
        self.schema = schema
        self.time_partitioning = time_partitioning
        self.range_partitioning = range_partitioning
        self.clustering = clustering
        self.require_partition_filter = require_partition_filter
        self.expiration_time = expiration_time
        self.view = view
        self.materialized_view = materialized_view
        self.external_data_configuration = external_data_configuration
        self.encryption_configuration = encryption_configuration
        self.default_collation = default_collation
        self.service_account_file = service_account_file

    def apply(self):
        stub = table_pb2_grpc.BigqueryAlphaTableServiceStub(channel.Channel())
        request = table_pb2.ApplyBigqueryAlphaTableRequest()
        if Primitive.to_proto(self.name):
            request.resource.name = Primitive.to_proto(self.name)

        if Primitive.to_proto(self.dataset):
            request.resource.dataset = Primitive.to_proto(self.dataset)

        if Primitive.to_proto(self.project):
            request.resource.project = Primitive.to_proto(self.project)

        if Primitive.to_proto(self.friendly_name):
            request.resource.friendly_name = Primitive.to_proto(self.friendly_name)

        if Primitive.to_proto(self.description):
            request.resource.description = Primitive.to_proto(self.description)

        if Primitive.to_proto(self.labels):
            request.resource.labels = Primitive.to_proto(self.labels)

        if TableSchema.to_proto(self.schema):
            request.resource.schema.CopyFrom(TableSchema.to_proto(self.schema))
        else:
            request.resource.ClearField("schema")
        if TableTimePartitioning.to_proto(self.time_partitioning):
            request.resource.time_partitioning.CopyFrom(
                TableTimePartitioning.to_proto(self.time_partitioning)
            )
        else:
            request.resource.ClearField("time_partitioning")
        if TableRangePartitioning.to_proto(self.range_partitioning):
            request.resource.range_partitioning.CopyFrom(
                TableRangePartitioning.to_proto(self.range_partitioning)
            )
        else:
            request.resource.ClearField("range_partitioning")
        if TableClustering.to_proto(self.clustering):
            request.resource.clustering.CopyFrom(
                TableClustering.to_proto(self.clustering)
            )
        else:
            request.resource.ClearField("clustering")
        if Primitive.to_proto(self.require_partition_filter):
            request.resource.require_partition_filter = Primitive.to_proto(
                self.require_partition_filter
            )

        if Primitive.to_proto(self.expiration_time):
            request.resource.expiration_time = Primitive.to_proto(self.expiration_time)

        if TableView.to_proto(self.view):
            request.resource.view.CopyFrom(TableView.to_proto(self.view))
        else:
            request.resource.ClearField("view")
        if TableMaterializedView.to_proto(self.materialized_view):
            request.resource.materialized_view.CopyFrom(
                TableMaterializedView.to_proto(self.materialized_view)
            )
        else:
            request.resource.ClearField("materialized_view")
        if TableExternalDataConfiguration.to_proto(self.external_data_configuration):
            request.resource.external_data_configuration.CopyFrom(
                TableExternalDataConfiguration.to_proto(
                    self.external_data_configuration
                )
            )
        else:
            request.resource.ClearField("external_data_configuration")
        if TableEncryptionConfiguration.to_proto(self.encryption_configuration):
            request.resource.encryption_configuration.CopyFrom(
                TableEncryptionConfiguration.to_proto(self.encryption_configuration)
            )
        else:
            request.resource.ClearField("encryption_configuration")
        if Primitive.to_proto(self.default_collation):
            request.resource.default_collation = Primitive.to_proto(
                self.default_collation
            )

        request.service_account_file = self.service_account_file

        response = stub.ApplyBigqueryAlphaTable(request)
        self.etag = Primitive.from_proto(response.etag)
        self.id = Primitive.from_proto(response.id)
        self.self_link = Primitive.from_proto(response.self_link)
        self.name = Primitive.from_proto(response.name)
        self.dataset = Primitive.from_proto(response.dataset)
        self.project = Primitive.from_proto(response.project)
        self.friendly_name = Primitive.from_proto(response.friendly_name)
        self.description = Primitive.from_proto(response.description)
        self.labels = Primitive.from_proto(response.labels)
        self.model = TableModel.from_proto(response.model)
        self.schema = TableSchema.from_proto(response.schema)
        self.time_partitioning = TableTimePartitioning.from_proto(
            response.time_partitioning
        )
        self.range_partitioning = TableRangePartitioning.from_proto(
            response.range_partitioning
        )
        self.clustering = TableClustering.from_proto(response.clustering)
        self.require_partition_filter = Primitive.from_proto(
            response.require_partition_filter
        )
        self.num_bytes = Primitive.from_proto(response.num_bytes)
        self.num_physical_bytes = Primitive.from_proto(response.num_physical_bytes)
        self.num_long_term_bytes = Primitive.from_proto(response.num_long_term_bytes)
        self.num_rows = Primitive.from_proto(response.num_rows)
        self.creation_time = Primitive.from_proto(response.creation_time)
        self.expiration_time = Primitive.from_proto(response.expiration_time)
        self.last_modified_time = Primitive.from_proto(response.last_modified_time)
        self.type = Primitive.from_proto(response.type)
        self.view = TableView.from_proto(response.view)
        self.materialized_view = TableMaterializedView.from_proto(
            response.materialized_view
        )
        self.external_data_configuration = TableExternalDataConfiguration.from_proto(
            response.external_data_configuration
        )
        self.location = Primitive.from_proto(response.location)
        self.streaming_buffer = TableStreamingBuffer.from_proto(
            response.streaming_buffer
        )
        self.encryption_configuration = TableEncryptionConfiguration.from_proto(
            response.encryption_configuration
        )
        self.snapshot_definition = TableSnapshotDefinition.from_proto(
            response.snapshot_definition
        )
        self.default_collation = Primitive.from_proto(response.default_collation)

    def delete(self):
        stub = table_pb2_grpc.BigqueryAlphaTableServiceStub(channel.Channel())
        request = table_pb2.DeleteBigqueryAlphaTableRequest()
        request.service_account_file = self.service_account_file
        if Primitive.to_proto(self.name):
            request.resource.name = Primitive.to_proto(self.name)

        if Primitive.to_proto(self.dataset):
            request.resource.dataset = Primitive.to_proto(self.dataset)

        if Primitive.to_proto(self.project):
            request.resource.project = Primitive.to_proto(self.project)

        if Primitive.to_proto(self.friendly_name):
            request.resource.friendly_name = Primitive.to_proto(self.friendly_name)

        if Primitive.to_proto(self.description):
            request.resource.description = Primitive.to_proto(self.description)

        if Primitive.to_proto(self.labels):
            request.resource.labels = Primitive.to_proto(self.labels)

        if TableSchema.to_proto(self.schema):
            request.resource.schema.CopyFrom(TableSchema.to_proto(self.schema))
        else:
            request.resource.ClearField("schema")
        if TableTimePartitioning.to_proto(self.time_partitioning):
            request.resource.time_partitioning.CopyFrom(
                TableTimePartitioning.to_proto(self.time_partitioning)
            )
        else:
            request.resource.ClearField("time_partitioning")
        if TableRangePartitioning.to_proto(self.range_partitioning):
            request.resource.range_partitioning.CopyFrom(
                TableRangePartitioning.to_proto(self.range_partitioning)
            )
        else:
            request.resource.ClearField("range_partitioning")
        if TableClustering.to_proto(self.clustering):
            request.resource.clustering.CopyFrom(
                TableClustering.to_proto(self.clustering)
            )
        else:
            request.resource.ClearField("clustering")
        if Primitive.to_proto(self.require_partition_filter):
            request.resource.require_partition_filter = Primitive.to_proto(
                self.require_partition_filter
            )

        if Primitive.to_proto(self.expiration_time):
            request.resource.expiration_time = Primitive.to_proto(self.expiration_time)

        if TableView.to_proto(self.view):
            request.resource.view.CopyFrom(TableView.to_proto(self.view))
        else:
            request.resource.ClearField("view")
        if TableMaterializedView.to_proto(self.materialized_view):
            request.resource.materialized_view.CopyFrom(
                TableMaterializedView.to_proto(self.materialized_view)
            )
        else:
            request.resource.ClearField("materialized_view")
        if TableExternalDataConfiguration.to_proto(self.external_data_configuration):
            request.resource.external_data_configuration.CopyFrom(
                TableExternalDataConfiguration.to_proto(
                    self.external_data_configuration
                )
            )
        else:
            request.resource.ClearField("external_data_configuration")
        if TableEncryptionConfiguration.to_proto(self.encryption_configuration):
            request.resource.encryption_configuration.CopyFrom(
                TableEncryptionConfiguration.to_proto(self.encryption_configuration)
            )
        else:
            request.resource.ClearField("encryption_configuration")
        if Primitive.to_proto(self.default_collation):
            request.resource.default_collation = Primitive.to_proto(
                self.default_collation
            )

        response = stub.DeleteBigqueryAlphaTable(request)

    @classmethod
    def list(self, project, dataset, service_account_file=""):
        stub = table_pb2_grpc.BigqueryAlphaTableServiceStub(channel.Channel())
        request = table_pb2.ListBigqueryAlphaTableRequest()
        request.service_account_file = service_account_file
        request.Project = project

        request.Dataset = dataset

        return stub.ListBigqueryAlphaTable(request).items

    def to_proto(self):
        resource = table_pb2.BigqueryAlphaTable()
        if Primitive.to_proto(self.name):
            resource.name = Primitive.to_proto(self.name)
        if Primitive.to_proto(self.dataset):
            resource.dataset = Primitive.to_proto(self.dataset)
        if Primitive.to_proto(self.project):
            resource.project = Primitive.to_proto(self.project)
        if Primitive.to_proto(self.friendly_name):
            resource.friendly_name = Primitive.to_proto(self.friendly_name)
        if Primitive.to_proto(self.description):
            resource.description = Primitive.to_proto(self.description)
        if Primitive.to_proto(self.labels):
            resource.labels = Primitive.to_proto(self.labels)
        if TableSchema.to_proto(self.schema):
            resource.schema.CopyFrom(TableSchema.to_proto(self.schema))
        else:
            resource.ClearField("schema")
        if TableTimePartitioning.to_proto(self.time_partitioning):
            resource.time_partitioning.CopyFrom(
                TableTimePartitioning.to_proto(self.time_partitioning)
            )
        else:
            resource.ClearField("time_partitioning")
        if TableRangePartitioning.to_proto(self.range_partitioning):
            resource.range_partitioning.CopyFrom(
                TableRangePartitioning.to_proto(self.range_partitioning)
            )
        else:
            resource.ClearField("range_partitioning")
        if TableClustering.to_proto(self.clustering):
            resource.clustering.CopyFrom(TableClustering.to_proto(self.clustering))
        else:
            resource.ClearField("clustering")
        if Primitive.to_proto(self.require_partition_filter):
            resource.require_partition_filter = Primitive.to_proto(
                self.require_partition_filter
            )
        if Primitive.to_proto(self.expiration_time):
            resource.expiration_time = Primitive.to_proto(self.expiration_time)
        if TableView.to_proto(self.view):
            resource.view.CopyFrom(TableView.to_proto(self.view))
        else:
            resource.ClearField("view")
        if TableMaterializedView.to_proto(self.materialized_view):
            resource.materialized_view.CopyFrom(
                TableMaterializedView.to_proto(self.materialized_view)
            )
        else:
            resource.ClearField("materialized_view")
        if TableExternalDataConfiguration.to_proto(self.external_data_configuration):
            resource.external_data_configuration.CopyFrom(
                TableExternalDataConfiguration.to_proto(
                    self.external_data_configuration
                )
            )
        else:
            resource.ClearField("external_data_configuration")
        if TableEncryptionConfiguration.to_proto(self.encryption_configuration):
            resource.encryption_configuration.CopyFrom(
                TableEncryptionConfiguration.to_proto(self.encryption_configuration)
            )
        else:
            resource.ClearField("encryption_configuration")
        if Primitive.to_proto(self.default_collation):
            resource.default_collation = Primitive.to_proto(self.default_collation)
        return resource


class TableModel(object):
    def __init__(self, model_options: dict = None, training_runs: list = None):
        self.model_options = model_options
        self.training_runs = training_runs

    @classmethod
    def to_proto(self, resource):
        if not resource:
            return None

        res = table_pb2.BigqueryAlphaTableModel()
        if TableModelModelOptions.to_proto(resource.model_options):
            res.model_options.CopyFrom(
                TableModelModelOptions.to_proto(resource.model_options)
            )
        else:
            res.ClearField("model_options")
        if TableModelTrainingRunsArray.to_proto(resource.training_runs):
            res.training_runs.extend(
                TableModelTrainingRunsArray.to_proto(resource.training_runs)
            )
        return res

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return None

        return TableModel(
            model_options=TableModelModelOptions.from_proto(resource.model_options),
            training_runs=TableModelTrainingRunsArray.from_proto(
                resource.training_runs
            ),
        )


class TableModelArray(object):
    @classmethod
    def to_proto(self, resources):
        if not resources:
            return resources
        return [TableModel.to_proto(i) for i in resources]

    @classmethod
    def from_proto(self, resources):
        return [TableModel.from_proto(i) for i in resources]


class TableModelModelOptions(object):
    def __init__(
        self, model_type: str = None, labels: list = None, loss_type: str = None
    ):
        self.model_type = model_type
        self.labels = labels
        self.loss_type = loss_type

    @classmethod
    def to_proto(self, resource):
        if not resource:
            return None

        res = table_pb2.BigqueryAlphaTableModelModelOptions()
        if Primitive.to_proto(resource.model_type):
            res.model_type = Primitive.to_proto(resource.model_type)
        if Primitive.to_proto(resource.labels):
            res.labels.extend(Primitive.to_proto(resource.labels))
        if Primitive.to_proto(resource.loss_type):
            res.loss_type = Primitive.to_proto(resource.loss_type)
        return res

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return None

        return TableModelModelOptions(
            model_type=Primitive.from_proto(resource.model_type),
            labels=Primitive.from_proto(resource.labels),
            loss_type=Primitive.from_proto(resource.loss_type),
        )


class TableModelModelOptionsArray(object):
    @classmethod
    def to_proto(self, resources):
        if not resources:
            return resources
        return [TableModelModelOptions.to_proto(i) for i in resources]

    @classmethod
    def from_proto(self, resources):
        return [TableModelModelOptions.from_proto(i) for i in resources]


class TableModelTrainingRuns(object):
    def __init__(
        self,
        state: str = None,
        start_time: str = None,
        training_options: dict = None,
        iteration_results: list = None,
    ):
        self.state = state
        self.start_time = start_time
        self.training_options = training_options
        self.iteration_results = iteration_results

    @classmethod
    def to_proto(self, resource):
        if not resource:
            return None

        res = table_pb2.BigqueryAlphaTableModelTrainingRuns()
        if Primitive.to_proto(resource.state):
            res.state = Primitive.to_proto(resource.state)
        if Primitive.to_proto(resource.start_time):
            res.start_time = Primitive.to_proto(resource.start_time)
        if TableModelTrainingRunsTrainingOptions.to_proto(resource.training_options):
            res.training_options.CopyFrom(
                TableModelTrainingRunsTrainingOptions.to_proto(
                    resource.training_options
                )
            )
        else:
            res.ClearField("training_options")
        if TableModelTrainingRunsIterationResultsArray.to_proto(
            resource.iteration_results
        ):
            res.iteration_results.extend(
                TableModelTrainingRunsIterationResultsArray.to_proto(
                    resource.iteration_results
                )
            )
        return res

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return None

        return TableModelTrainingRuns(
            state=Primitive.from_proto(resource.state),
            start_time=Primitive.from_proto(resource.start_time),
            training_options=TableModelTrainingRunsTrainingOptions.from_proto(
                resource.training_options
            ),
            iteration_results=TableModelTrainingRunsIterationResultsArray.from_proto(
                resource.iteration_results
            ),
        )


class TableModelTrainingRunsArray(object):
    @classmethod
    def to_proto(self, resources):
        if not resources:
            return resources
        return [TableModelTrainingRuns.to_proto(i) for i in resources]

    @classmethod
    def from_proto(self, resources):
        return [TableModelTrainingRuns.from_proto(i) for i in resources]


class TableModelTrainingRunsTrainingOptions(object):
    def __init__(
        self,
        max_iteration: int = None,
        learn_rate: float = None,
        l1_reg: float = None,
        l2_reg: float = None,
        min_rel_progress: float = None,
        warm_start: bool = None,
        early_stop: bool = None,
        learn_rate_strategy: str = None,
        line_search_init_learn_rate: float = None,
    ):
        self.max_iteration = max_iteration
        self.learn_rate = learn_rate
        self.l1_reg = l1_reg
        self.l2_reg = l2_reg
        self.min_rel_progress = min_rel_progress
        self.warm_start = warm_start
        self.early_stop = early_stop
        self.learn_rate_strategy = learn_rate_strategy
        self.line_search_init_learn_rate = line_search_init_learn_rate

    @classmethod
    def to_proto(self, resource):
        if not resource:
            return None

        res = table_pb2.BigqueryAlphaTableModelTrainingRunsTrainingOptions()
        if Primitive.to_proto(resource.max_iteration):
            res.max_iteration = Primitive.to_proto(resource.max_iteration)
        if Primitive.to_proto(resource.learn_rate):
            res.learn_rate = Primitive.to_proto(resource.learn_rate)
        if Primitive.to_proto(resource.l1_reg):
            res.l1_reg = Primitive.to_proto(resource.l1_reg)
        if Primitive.to_proto(resource.l2_reg):
            res.l2_reg = Primitive.to_proto(resource.l2_reg)
        if Primitive.to_proto(resource.min_rel_progress):
            res.min_rel_progress = Primitive.to_proto(resource.min_rel_progress)
        if Primitive.to_proto(resource.warm_start):
            res.warm_start = Primitive.to_proto(resource.warm_start)
        if Primitive.to_proto(resource.early_stop):
            res.early_stop = Primitive.to_proto(resource.early_stop)
        if Primitive.to_proto(resource.learn_rate_strategy):
            res.learn_rate_strategy = Primitive.to_proto(resource.learn_rate_strategy)
        if Primitive.to_proto(resource.line_search_init_learn_rate):
            res.line_search_init_learn_rate = Primitive.to_proto(
                resource.line_search_init_learn_rate
            )
        return res

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return None

        return TableModelTrainingRunsTrainingOptions(
            max_iteration=Primitive.from_proto(resource.max_iteration),
            learn_rate=Primitive.from_proto(resource.learn_rate),
            l1_reg=Primitive.from_proto(resource.l1_reg),
            l2_reg=Primitive.from_proto(resource.l2_reg),
            min_rel_progress=Primitive.from_proto(resource.min_rel_progress),
            warm_start=Primitive.from_proto(resource.warm_start),
            early_stop=Primitive.from_proto(resource.early_stop),
            learn_rate_strategy=Primitive.from_proto(resource.learn_rate_strategy),
            line_search_init_learn_rate=Primitive.from_proto(
                resource.line_search_init_learn_rate
            ),
        )


class TableModelTrainingRunsTrainingOptionsArray(object):
    @classmethod
    def to_proto(self, resources):
        if not resources:
            return resources
        return [TableModelTrainingRunsTrainingOptions.to_proto(i) for i in resources]

    @classmethod
    def from_proto(self, resources):
        return [TableModelTrainingRunsTrainingOptions.from_proto(i) for i in resources]


class TableModelTrainingRunsIterationResults(object):
    def __init__(
        self,
        index: int = None,
        learn_rate: float = None,
        training_loss: float = None,
        eval_loss: float = None,
        duration_ms: str = None,
    ):
        self.index = index
        self.learn_rate = learn_rate
        self.training_loss = training_loss
        self.eval_loss = eval_loss
        self.duration_ms = duration_ms

    @classmethod
    def to_proto(self, resource):
        if not resource:
            return None

        res = table_pb2.BigqueryAlphaTableModelTrainingRunsIterationResults()
        if Primitive.to_proto(resource.index):
            res.index = Primitive.to_proto(resource.index)
        if Primitive.to_proto(resource.learn_rate):
            res.learn_rate = Primitive.to_proto(resource.learn_rate)
        if Primitive.to_proto(resource.training_loss):
            res.training_loss = Primitive.to_proto(resource.training_loss)
        if Primitive.to_proto(resource.eval_loss):
            res.eval_loss = Primitive.to_proto(resource.eval_loss)
        if Primitive.to_proto(resource.duration_ms):
            res.duration_ms = Primitive.to_proto(resource.duration_ms)
        return res

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return None

        return TableModelTrainingRunsIterationResults(
            index=Primitive.from_proto(resource.index),
            learn_rate=Primitive.from_proto(resource.learn_rate),
            training_loss=Primitive.from_proto(resource.training_loss),
            eval_loss=Primitive.from_proto(resource.eval_loss),
            duration_ms=Primitive.from_proto(resource.duration_ms),
        )


class TableModelTrainingRunsIterationResultsArray(object):
    @classmethod
    def to_proto(self, resources):
        if not resources:
            return resources
        return [TableModelTrainingRunsIterationResults.to_proto(i) for i in resources]

    @classmethod
    def from_proto(self, resources):
        return [TableModelTrainingRunsIterationResults.from_proto(i) for i in resources]


class TableSchema(object):
    def __init__(self, fields: list = None):
        self.fields = fields

    @classmethod
    def to_proto(self, resource):
        if not resource:
            return None

        res = table_pb2.BigqueryAlphaTableSchema()
        if TableGooglecloudbigqueryv2TablefieldschemaArray.to_proto(resource.fields):
            res.fields.extend(
                TableGooglecloudbigqueryv2TablefieldschemaArray.to_proto(
                    resource.fields
                )
            )
        return res

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return None

        return TableSchema(
            fields=TableGooglecloudbigqueryv2TablefieldschemaArray.from_proto(
                resource.fields
            ),
        )


class TableSchemaArray(object):
    @classmethod
    def to_proto(self, resources):
        if not resources:
            return resources
        return [TableSchema.to_proto(i) for i in resources]

    @classmethod
    def from_proto(self, resources):
        return [TableSchema.from_proto(i) for i in resources]


class TableGooglecloudbigqueryv2Tablefieldschema(object):
    def __init__(
        self,
        name: str = None,
        type: str = None,
        mode: str = None,
        fields: list = None,
        description: str = None,
        categories: dict = None,
        policy_tags: dict = None,
        name_alternative: list = None,
        max_length: int = None,
        precision: int = None,
        scale: int = None,
        collation: str = None,
        default_value_expression: str = None,
    ):
        self.name = name
        self.type = type
        self.mode = mode
        self.fields = fields
        self.description = description
        self.categories = categories
        self.policy_tags = policy_tags
        self.name_alternative = name_alternative
        self.max_length = max_length
        self.precision = precision
        self.scale = scale
        self.collation = collation
        self.default_value_expression = default_value_expression

    @classmethod
    def to_proto(self, resource):
        if not resource:
            return None

        res = table_pb2.BigqueryAlphaTableGooglecloudbigqueryv2Tablefieldschema()
        if Primitive.to_proto(resource.name):
            res.name = Primitive.to_proto(resource.name)
        if Primitive.to_proto(resource.type):
            res.type = Primitive.to_proto(resource.type)
        if Primitive.to_proto(resource.mode):
            res.mode = Primitive.to_proto(resource.mode)
        if TableGooglecloudbigqueryv2TablefieldschemaArray.to_proto(resource.fields):
            res.fields.extend(
                TableGooglecloudbigqueryv2TablefieldschemaArray.to_proto(
                    resource.fields
                )
            )
        if Primitive.to_proto(resource.description):
            res.description = Primitive.to_proto(resource.description)
        if TableGooglecloudbigqueryv2TablefieldschemaCategories.to_proto(
            resource.categories
        ):
            res.categories.CopyFrom(
                TableGooglecloudbigqueryv2TablefieldschemaCategories.to_proto(
                    resource.categories
                )
            )
        else:
            res.ClearField("categories")
        if TableGooglecloudbigqueryv2TablefieldschemaPolicyTags.to_proto(
            resource.policy_tags
        ):
            res.policy_tags.CopyFrom(
                TableGooglecloudbigqueryv2TablefieldschemaPolicyTags.to_proto(
                    resource.policy_tags
                )
            )
        else:
            res.ClearField("policy_tags")
        if Primitive.to_proto(resource.name_alternative):
            res.name_alternative.extend(Primitive.to_proto(resource.name_alternative))
        if Primitive.to_proto(resource.max_length):
            res.max_length = Primitive.to_proto(resource.max_length)
        if Primitive.to_proto(resource.precision):
            res.precision = Primitive.to_proto(resource.precision)
        if Primitive.to_proto(resource.scale):
            res.scale = Primitive.to_proto(resource.scale)
        if Primitive.to_proto(resource.collation):
            res.collation = Primitive.to_proto(resource.collation)
        if Primitive.to_proto(resource.default_value_expression):
            res.default_value_expression = Primitive.to_proto(
                resource.default_value_expression
            )
        return res

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return None

        return TableGooglecloudbigqueryv2Tablefieldschema(
            name=Primitive.from_proto(resource.name),
            type=Primitive.from_proto(resource.type),
            mode=Primitive.from_proto(resource.mode),
            fields=TableGooglecloudbigqueryv2TablefieldschemaArray.from_proto(
                resource.fields
            ),
            description=Primitive.from_proto(resource.description),
            categories=TableGooglecloudbigqueryv2TablefieldschemaCategories.from_proto(
                resource.categories
            ),
            policy_tags=TableGooglecloudbigqueryv2TablefieldschemaPolicyTags.from_proto(
                resource.policy_tags
            ),
            name_alternative=Primitive.from_proto(resource.name_alternative),
            max_length=Primitive.from_proto(resource.max_length),
            precision=Primitive.from_proto(resource.precision),
            scale=Primitive.from_proto(resource.scale),
            collation=Primitive.from_proto(resource.collation),
            default_value_expression=Primitive.from_proto(
                resource.default_value_expression
            ),
        )


class TableGooglecloudbigqueryv2TablefieldschemaArray(object):
    @classmethod
    def to_proto(self, resources):
        if not resources:
            return resources
        return [
            TableGooglecloudbigqueryv2Tablefieldschema.to_proto(i) for i in resources
        ]

    @classmethod
    def from_proto(self, resources):
        return [
            TableGooglecloudbigqueryv2Tablefieldschema.from_proto(i) for i in resources
        ]


class TableGooglecloudbigqueryv2TablefieldschemaCategories(object):
    def __init__(self, names: list = None):
        self.names = names

    @classmethod
    def to_proto(self, resource):
        if not resource:
            return None

        res = (
            table_pb2.BigqueryAlphaTableGooglecloudbigqueryv2TablefieldschemaCategories()
        )
        if Primitive.to_proto(resource.names):
            res.names.extend(Primitive.to_proto(resource.names))
        return res

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return None

        return TableGooglecloudbigqueryv2TablefieldschemaCategories(
            names=Primitive.from_proto(resource.names),
        )


class TableGooglecloudbigqueryv2TablefieldschemaCategoriesArray(object):
    @classmethod
    def to_proto(self, resources):
        if not resources:
            return resources
        return [
            TableGooglecloudbigqueryv2TablefieldschemaCategories.to_proto(i)
            for i in resources
        ]

    @classmethod
    def from_proto(self, resources):
        return [
            TableGooglecloudbigqueryv2TablefieldschemaCategories.from_proto(i)
            for i in resources
        ]


class TableGooglecloudbigqueryv2TablefieldschemaPolicyTags(object):
    def __init__(self, names: list = None):
        self.names = names

    @classmethod
    def to_proto(self, resource):
        if not resource:
            return None

        res = (
            table_pb2.BigqueryAlphaTableGooglecloudbigqueryv2TablefieldschemaPolicyTags()
        )
        if Primitive.to_proto(resource.names):
            res.names.extend(Primitive.to_proto(resource.names))
        return res

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return None

        return TableGooglecloudbigqueryv2TablefieldschemaPolicyTags(
            names=Primitive.from_proto(resource.names),
        )


class TableGooglecloudbigqueryv2TablefieldschemaPolicyTagsArray(object):
    @classmethod
    def to_proto(self, resources):
        if not resources:
            return resources
        return [
            TableGooglecloudbigqueryv2TablefieldschemaPolicyTags.to_proto(i)
            for i in resources
        ]

    @classmethod
    def from_proto(self, resources):
        return [
            TableGooglecloudbigqueryv2TablefieldschemaPolicyTags.from_proto(i)
            for i in resources
        ]


class TableTimePartitioning(object):
    def __init__(self, type: str = None, expiration_ms: str = None, field: str = None):
        self.type = type
        self.expiration_ms = expiration_ms
        self.field = field

    @classmethod
    def to_proto(self, resource):
        if not resource:
            return None

        res = table_pb2.BigqueryAlphaTableTimePartitioning()
        if Primitive.to_proto(resource.type):
            res.type = Primitive.to_proto(resource.type)
        if Primitive.to_proto(resource.expiration_ms):
            res.expiration_ms = Primitive.to_proto(resource.expiration_ms)
        if Primitive.to_proto(resource.field):
            res.field = Primitive.to_proto(resource.field)
        return res

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return None

        return TableTimePartitioning(
            type=Primitive.from_proto(resource.type),
            expiration_ms=Primitive.from_proto(resource.expiration_ms),
            field=Primitive.from_proto(resource.field),
        )


class TableTimePartitioningArray(object):
    @classmethod
    def to_proto(self, resources):
        if not resources:
            return resources
        return [TableTimePartitioning.to_proto(i) for i in resources]

    @classmethod
    def from_proto(self, resources):
        return [TableTimePartitioning.from_proto(i) for i in resources]


class TableRangePartitioning(object):
    def __init__(self, field: str = None, range: dict = None):
        self.field = field
        self.range = range

    @classmethod
    def to_proto(self, resource):
        if not resource:
            return None

        res = table_pb2.BigqueryAlphaTableRangePartitioning()
        if Primitive.to_proto(resource.field):
            res.field = Primitive.to_proto(resource.field)
        if TableRangePartitioningRange.to_proto(resource.range):
            res.range.CopyFrom(TableRangePartitioningRange.to_proto(resource.range))
        else:
            res.ClearField("range")
        return res

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return None

        return TableRangePartitioning(
            field=Primitive.from_proto(resource.field),
            range=TableRangePartitioningRange.from_proto(resource.range),
        )


class TableRangePartitioningArray(object):
    @classmethod
    def to_proto(self, resources):
        if not resources:
            return resources
        return [TableRangePartitioning.to_proto(i) for i in resources]

    @classmethod
    def from_proto(self, resources):
        return [TableRangePartitioning.from_proto(i) for i in resources]


class TableRangePartitioningRange(object):
    def __init__(self, start: str = None, end: str = None, interval: str = None):
        self.start = start
        self.end = end
        self.interval = interval

    @classmethod
    def to_proto(self, resource):
        if not resource:
            return None

        res = table_pb2.BigqueryAlphaTableRangePartitioningRange()
        if Primitive.to_proto(resource.start):
            res.start = Primitive.to_proto(resource.start)
        if Primitive.to_proto(resource.end):
            res.end = Primitive.to_proto(resource.end)
        if Primitive.to_proto(resource.interval):
            res.interval = Primitive.to_proto(resource.interval)
        return res

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return None

        return TableRangePartitioningRange(
            start=Primitive.from_proto(resource.start),
            end=Primitive.from_proto(resource.end),
            interval=Primitive.from_proto(resource.interval),
        )


class TableRangePartitioningRangeArray(object):
    @classmethod
    def to_proto(self, resources):
        if not resources:
            return resources
        return [TableRangePartitioningRange.to_proto(i) for i in resources]

    @classmethod
    def from_proto(self, resources):
        return [TableRangePartitioningRange.from_proto(i) for i in resources]


class TableClustering(object):
    def __init__(self, fields: list = None):
        self.fields = fields

    @classmethod
    def to_proto(self, resource):
        if not resource:
            return None

        res = table_pb2.BigqueryAlphaTableClustering()
        if Primitive.to_proto(resource.fields):
            res.fields.extend(Primitive.to_proto(resource.fields))
        return res

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return None

        return TableClustering(fields=Primitive.from_proto(resource.fields),)


class TableClusteringArray(object):
    @classmethod
    def to_proto(self, resources):
        if not resources:
            return resources
        return [TableClustering.to_proto(i) for i in resources]

    @classmethod
    def from_proto(self, resources):
        return [TableClustering.from_proto(i) for i in resources]


class TableView(object):
    def __init__(
        self,
        query: str = None,
        user_defined_function_resources: list = None,
        use_legacy_sql: bool = None,
        use_explicit_column_names: bool = None,
    ):
        self.query = query
        self.user_defined_function_resources = user_defined_function_resources
        self.use_legacy_sql = use_legacy_sql
        self.use_explicit_column_names = use_explicit_column_names

    @classmethod
    def to_proto(self, resource):
        if not resource:
            return None

        res = table_pb2.BigqueryAlphaTableView()
        if Primitive.to_proto(resource.query):
            res.query = Primitive.to_proto(resource.query)
        if TableViewUserDefinedFunctionResourcesArray.to_proto(
            resource.user_defined_function_resources
        ):
            res.user_defined_function_resources.extend(
                TableViewUserDefinedFunctionResourcesArray.to_proto(
                    resource.user_defined_function_resources
                )
            )
        if Primitive.to_proto(resource.use_legacy_sql):
            res.use_legacy_sql = Primitive.to_proto(resource.use_legacy_sql)
        if Primitive.to_proto(resource.use_explicit_column_names):
            res.use_explicit_column_names = Primitive.to_proto(
                resource.use_explicit_column_names
            )
        return res

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return None

        return TableView(
            query=Primitive.from_proto(resource.query),
            user_defined_function_resources=TableViewUserDefinedFunctionResourcesArray.from_proto(
                resource.user_defined_function_resources
            ),
            use_legacy_sql=Primitive.from_proto(resource.use_legacy_sql),
            use_explicit_column_names=Primitive.from_proto(
                resource.use_explicit_column_names
            ),
        )


class TableViewArray(object):
    @classmethod
    def to_proto(self, resources):
        if not resources:
            return resources
        return [TableView.to_proto(i) for i in resources]

    @classmethod
    def from_proto(self, resources):
        return [TableView.from_proto(i) for i in resources]


class TableViewUserDefinedFunctionResources(object):
    def __init__(
        self,
        resource_uri: str = None,
        inline_code: str = None,
        inline_code_alternative: list = None,
    ):
        self.resource_uri = resource_uri
        self.inline_code = inline_code
        self.inline_code_alternative = inline_code_alternative

    @classmethod
    def to_proto(self, resource):
        if not resource:
            return None

        res = table_pb2.BigqueryAlphaTableViewUserDefinedFunctionResources()
        if Primitive.to_proto(resource.resource_uri):
            res.resource_uri = Primitive.to_proto(resource.resource_uri)
        if Primitive.to_proto(resource.inline_code):
            res.inline_code = Primitive.to_proto(resource.inline_code)
        if Primitive.to_proto(resource.inline_code_alternative):
            res.inline_code_alternative.extend(
                Primitive.to_proto(resource.inline_code_alternative)
            )
        return res

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return None

        return TableViewUserDefinedFunctionResources(
            resource_uri=Primitive.from_proto(resource.resource_uri),
            inline_code=Primitive.from_proto(resource.inline_code),
            inline_code_alternative=Primitive.from_proto(
                resource.inline_code_alternative
            ),
        )


class TableViewUserDefinedFunctionResourcesArray(object):
    @classmethod
    def to_proto(self, resources):
        if not resources:
            return resources
        return [TableViewUserDefinedFunctionResources.to_proto(i) for i in resources]

    @classmethod
    def from_proto(self, resources):
        return [TableViewUserDefinedFunctionResources.from_proto(i) for i in resources]


class TableMaterializedView(object):
    def __init__(
        self,
        query: str = None,
        last_refresh_time: int = None,
        enable_refresh: bool = None,
        refresh_interval_ms: int = None,
    ):
        self.query = query
        self.last_refresh_time = last_refresh_time
        self.enable_refresh = enable_refresh
        self.refresh_interval_ms = refresh_interval_ms

    @classmethod
    def to_proto(self, resource):
        if not resource:
            return None

        res = table_pb2.BigqueryAlphaTableMaterializedView()
        if Primitive.to_proto(resource.query):
            res.query = Primitive.to_proto(resource.query)
        if Primitive.to_proto(resource.last_refresh_time):
            res.last_refresh_time = Primitive.to_proto(resource.last_refresh_time)
        if Primitive.to_proto(resource.enable_refresh):
            res.enable_refresh = Primitive.to_proto(resource.enable_refresh)
        if Primitive.to_proto(resource.refresh_interval_ms):
            res.refresh_interval_ms = Primitive.to_proto(resource.refresh_interval_ms)
        return res

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return None

        return TableMaterializedView(
            query=Primitive.from_proto(resource.query),
            last_refresh_time=Primitive.from_proto(resource.last_refresh_time),
            enable_refresh=Primitive.from_proto(resource.enable_refresh),
            refresh_interval_ms=Primitive.from_proto(resource.refresh_interval_ms),
        )


class TableMaterializedViewArray(object):
    @classmethod
    def to_proto(self, resources):
        if not resources:
            return resources
        return [TableMaterializedView.to_proto(i) for i in resources]

    @classmethod
    def from_proto(self, resources):
        return [TableMaterializedView.from_proto(i) for i in resources]


class TableExternalDataConfiguration(object):
    def __init__(
        self,
        source_uris: list = None,
        schema: dict = None,
        source_format: str = None,
        max_bad_records: int = None,
        autodetect: bool = None,
        ignore_unknown_values: bool = None,
        compression: str = None,
        csv_options: dict = None,
        bigtable_options: dict = None,
        google_sheets_options: dict = None,
        max_bad_records_alternative: list = None,
        hive_partitioning_options: dict = None,
        connection_id: str = None,
        value_conversion_modes: dict = None,
        decimal_target_types: list = None,
        avro_options: dict = None,
        json_extension: str = None,
        parquet_options: dict = None,
    ):
        self.source_uris = source_uris
        self.schema = schema
        self.source_format = source_format
        self.max_bad_records = max_bad_records
        self.autodetect = autodetect
        self.ignore_unknown_values = ignore_unknown_values
        self.compression = compression
        self.csv_options = csv_options
        self.bigtable_options = bigtable_options
        self.google_sheets_options = google_sheets_options
        self.max_bad_records_alternative = max_bad_records_alternative
        self.hive_partitioning_options = hive_partitioning_options
        self.connection_id = connection_id
        self.value_conversion_modes = value_conversion_modes
        self.decimal_target_types = decimal_target_types
        self.avro_options = avro_options
        self.json_extension = json_extension
        self.parquet_options = parquet_options

    @classmethod
    def to_proto(self, resource):
        if not resource:
            return None

        res = table_pb2.BigqueryAlphaTableExternalDataConfiguration()
        if Primitive.to_proto(resource.source_uris):
            res.source_uris.extend(Primitive.to_proto(resource.source_uris))
        if TableExternalDataConfigurationSchema.to_proto(resource.schema):
            res.schema.CopyFrom(
                TableExternalDataConfigurationSchema.to_proto(resource.schema)
            )
        else:
            res.ClearField("schema")
        if Primitive.to_proto(resource.source_format):
            res.source_format = Primitive.to_proto(resource.source_format)
        if Primitive.to_proto(resource.max_bad_records):
            res.max_bad_records = Primitive.to_proto(resource.max_bad_records)
        if Primitive.to_proto(resource.autodetect):
            res.autodetect = Primitive.to_proto(resource.autodetect)
        if Primitive.to_proto(resource.ignore_unknown_values):
            res.ignore_unknown_values = Primitive.to_proto(
                resource.ignore_unknown_values
            )
        if Primitive.to_proto(resource.compression):
            res.compression = Primitive.to_proto(resource.compression)
        if TableExternalDataConfigurationCsvOptions.to_proto(resource.csv_options):
            res.csv_options.CopyFrom(
                TableExternalDataConfigurationCsvOptions.to_proto(resource.csv_options)
            )
        else:
            res.ClearField("csv_options")
        if TableExternalDataConfigurationBigtableOptions.to_proto(
            resource.bigtable_options
        ):
            res.bigtable_options.CopyFrom(
                TableExternalDataConfigurationBigtableOptions.to_proto(
                    resource.bigtable_options
                )
            )
        else:
            res.ClearField("bigtable_options")
        if TableExternalDataConfigurationGoogleSheetsOptions.to_proto(
            resource.google_sheets_options
        ):
            res.google_sheets_options.CopyFrom(
                TableExternalDataConfigurationGoogleSheetsOptions.to_proto(
                    resource.google_sheets_options
                )
            )
        else:
            res.ClearField("google_sheets_options")
        if int64Array.to_proto(resource.max_bad_records_alternative):
            res.max_bad_records_alternative.extend(
                int64Array.to_proto(resource.max_bad_records_alternative)
            )
        if TableExternalDataConfigurationHivePartitioningOptions.to_proto(
            resource.hive_partitioning_options
        ):
            res.hive_partitioning_options.CopyFrom(
                TableExternalDataConfigurationHivePartitioningOptions.to_proto(
                    resource.hive_partitioning_options
                )
            )
        else:
            res.ClearField("hive_partitioning_options")
        if Primitive.to_proto(resource.connection_id):
            res.connection_id = Primitive.to_proto(resource.connection_id)
        if TableExternalDataConfigurationValueConversionModes.to_proto(
            resource.value_conversion_modes
        ):
            res.value_conversion_modes.CopyFrom(
                TableExternalDataConfigurationValueConversionModes.to_proto(
                    resource.value_conversion_modes
                )
            )
        else:
            res.ClearField("value_conversion_modes")
        if TableExternalDataConfigurationDecimalTargetTypesEnumArray.to_proto(
            resource.decimal_target_types
        ):
            res.decimal_target_types.extend(
                TableExternalDataConfigurationDecimalTargetTypesEnumArray.to_proto(
                    resource.decimal_target_types
                )
            )
        if TableExternalDataConfigurationAvroOptions.to_proto(resource.avro_options):
            res.avro_options.CopyFrom(
                TableExternalDataConfigurationAvroOptions.to_proto(
                    resource.avro_options
                )
            )
        else:
            res.ClearField("avro_options")
        if TableExternalDataConfigurationJsonExtensionEnum.to_proto(
            resource.json_extension
        ):
            res.json_extension = TableExternalDataConfigurationJsonExtensionEnum.to_proto(
                resource.json_extension
            )
        if TableExternalDataConfigurationParquetOptions.to_proto(
            resource.parquet_options
        ):
            res.parquet_options.CopyFrom(
                TableExternalDataConfigurationParquetOptions.to_proto(
                    resource.parquet_options
                )
            )
        else:
            res.ClearField("parquet_options")
        return res

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return None

        return TableExternalDataConfiguration(
            source_uris=Primitive.from_proto(resource.source_uris),
            schema=TableExternalDataConfigurationSchema.from_proto(resource.schema),
            source_format=Primitive.from_proto(resource.source_format),
            max_bad_records=Primitive.from_proto(resource.max_bad_records),
            autodetect=Primitive.from_proto(resource.autodetect),
            ignore_unknown_values=Primitive.from_proto(resource.ignore_unknown_values),
            compression=Primitive.from_proto(resource.compression),
            csv_options=TableExternalDataConfigurationCsvOptions.from_proto(
                resource.csv_options
            ),
            bigtable_options=TableExternalDataConfigurationBigtableOptions.from_proto(
                resource.bigtable_options
            ),
            google_sheets_options=TableExternalDataConfigurationGoogleSheetsOptions.from_proto(
                resource.google_sheets_options
            ),
            max_bad_records_alternative=int64Array.from_proto(
                resource.max_bad_records_alternative
            ),
            hive_partitioning_options=TableExternalDataConfigurationHivePartitioningOptions.from_proto(
                resource.hive_partitioning_options
            ),
            connection_id=Primitive.from_proto(resource.connection_id),
            value_conversion_modes=TableExternalDataConfigurationValueConversionModes.from_proto(
                resource.value_conversion_modes
            ),
            decimal_target_types=TableExternalDataConfigurationDecimalTargetTypesEnumArray.from_proto(
                resource.decimal_target_types
            ),
            avro_options=TableExternalDataConfigurationAvroOptions.from_proto(
                resource.avro_options
            ),
            json_extension=TableExternalDataConfigurationJsonExtensionEnum.from_proto(
                resource.json_extension
            ),
            parquet_options=TableExternalDataConfigurationParquetOptions.from_proto(
                resource.parquet_options
            ),
        )


class TableExternalDataConfigurationArray(object):
    @classmethod
    def to_proto(self, resources):
        if not resources:
            return resources
        return [TableExternalDataConfiguration.to_proto(i) for i in resources]

    @classmethod
    def from_proto(self, resources):
        return [TableExternalDataConfiguration.from_proto(i) for i in resources]


class TableExternalDataConfigurationSchema(object):
    def __init__(self, fields: list = None):
        self.fields = fields

    @classmethod
    def to_proto(self, resource):
        if not resource:
            return None

        res = table_pb2.BigqueryAlphaTableExternalDataConfigurationSchema()
        if TableGooglecloudbigqueryv2TablefieldschemaArray.to_proto(resource.fields):
            res.fields.extend(
                TableGooglecloudbigqueryv2TablefieldschemaArray.to_proto(
                    resource.fields
                )
            )
        return res

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return None

        return TableExternalDataConfigurationSchema(
            fields=TableGooglecloudbigqueryv2TablefieldschemaArray.from_proto(
                resource.fields
            ),
        )


class TableExternalDataConfigurationSchemaArray(object):
    @classmethod
    def to_proto(self, resources):
        if not resources:
            return resources
        return [TableExternalDataConfigurationSchema.to_proto(i) for i in resources]

    @classmethod
    def from_proto(self, resources):
        return [TableExternalDataConfigurationSchema.from_proto(i) for i in resources]


class TableExternalDataConfigurationCsvOptions(object):
    def __init__(
        self,
        field_delimiter: str = None,
        skip_leading_rows: str = None,
        quote: str = None,
        allow_quoted_newlines: bool = None,
        allow_jagged_rows: bool = None,
        encoding: str = None,
    ):
        self.field_delimiter = field_delimiter
        self.skip_leading_rows = skip_leading_rows
        self.quote = quote
        self.allow_quoted_newlines = allow_quoted_newlines
        self.allow_jagged_rows = allow_jagged_rows
        self.encoding = encoding

    @classmethod
    def to_proto(self, resource):
        if not resource:
            return None

        res = table_pb2.BigqueryAlphaTableExternalDataConfigurationCsvOptions()
        if Primitive.to_proto(resource.field_delimiter):
            res.field_delimiter = Primitive.to_proto(resource.field_delimiter)
        if Primitive.to_proto(resource.skip_leading_rows):
            res.skip_leading_rows = Primitive.to_proto(resource.skip_leading_rows)
        if Primitive.to_proto(resource.quote):
            res.quote = Primitive.to_proto(resource.quote)
        if Primitive.to_proto(resource.allow_quoted_newlines):
            res.allow_quoted_newlines = Primitive.to_proto(
                resource.allow_quoted_newlines
            )
        if Primitive.to_proto(resource.allow_jagged_rows):
            res.allow_jagged_rows = Primitive.to_proto(resource.allow_jagged_rows)
        if Primitive.to_proto(resource.encoding):
            res.encoding = Primitive.to_proto(resource.encoding)
        return res

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return None

        return TableExternalDataConfigurationCsvOptions(
            field_delimiter=Primitive.from_proto(resource.field_delimiter),
            skip_leading_rows=Primitive.from_proto(resource.skip_leading_rows),
            quote=Primitive.from_proto(resource.quote),
            allow_quoted_newlines=Primitive.from_proto(resource.allow_quoted_newlines),
            allow_jagged_rows=Primitive.from_proto(resource.allow_jagged_rows),
            encoding=Primitive.from_proto(resource.encoding),
        )


class TableExternalDataConfigurationCsvOptionsArray(object):
    @classmethod
    def to_proto(self, resources):
        if not resources:
            return resources
        return [TableExternalDataConfigurationCsvOptions.to_proto(i) for i in resources]

    @classmethod
    def from_proto(self, resources):
        return [
            TableExternalDataConfigurationCsvOptions.from_proto(i) for i in resources
        ]


class TableExternalDataConfigurationBigtableOptions(object):
    def __init__(
        self,
        column_families: list = None,
        ignore_unspecified_column_families: bool = None,
        read_rowkey_as_string: bool = None,
    ):
        self.column_families = column_families
        self.ignore_unspecified_column_families = ignore_unspecified_column_families
        self.read_rowkey_as_string = read_rowkey_as_string

    @classmethod
    def to_proto(self, resource):
        if not resource:
            return None

        res = table_pb2.BigqueryAlphaTableExternalDataConfigurationBigtableOptions()
        if TableExternalDataConfigurationBigtableOptionsColumnFamiliesArray.to_proto(
            resource.column_families
        ):
            res.column_families.extend(
                TableExternalDataConfigurationBigtableOptionsColumnFamiliesArray.to_proto(
                    resource.column_families
                )
            )
        if Primitive.to_proto(resource.ignore_unspecified_column_families):
            res.ignore_unspecified_column_families = Primitive.to_proto(
                resource.ignore_unspecified_column_families
            )
        if Primitive.to_proto(resource.read_rowkey_as_string):
            res.read_rowkey_as_string = Primitive.to_proto(
                resource.read_rowkey_as_string
            )
        return res

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return None

        return TableExternalDataConfigurationBigtableOptions(
            column_families=TableExternalDataConfigurationBigtableOptionsColumnFamiliesArray.from_proto(
                resource.column_families
            ),
            ignore_unspecified_column_families=Primitive.from_proto(
                resource.ignore_unspecified_column_families
            ),
            read_rowkey_as_string=Primitive.from_proto(resource.read_rowkey_as_string),
        )


class TableExternalDataConfigurationBigtableOptionsArray(object):
    @classmethod
    def to_proto(self, resources):
        if not resources:
            return resources
        return [
            TableExternalDataConfigurationBigtableOptions.to_proto(i) for i in resources
        ]

    @classmethod
    def from_proto(self, resources):
        return [
            TableExternalDataConfigurationBigtableOptions.from_proto(i)
            for i in resources
        ]


class TableExternalDataConfigurationBigtableOptionsColumnFamilies(object):
    def __init__(
        self,
        family_id: str = None,
        type: str = None,
        encoding: str = None,
        columns: list = None,
        only_read_latest: bool = None,
    ):
        self.family_id = family_id
        self.type = type
        self.encoding = encoding
        self.columns = columns
        self.only_read_latest = only_read_latest

    @classmethod
    def to_proto(self, resource):
        if not resource:
            return None

        res = (
            table_pb2.BigqueryAlphaTableExternalDataConfigurationBigtableOptionsColumnFamilies()
        )
        if Primitive.to_proto(resource.family_id):
            res.family_id = Primitive.to_proto(resource.family_id)
        if Primitive.to_proto(resource.type):
            res.type = Primitive.to_proto(resource.type)
        if Primitive.to_proto(resource.encoding):
            res.encoding = Primitive.to_proto(resource.encoding)
        if TableExternalDataConfigurationBigtableOptionsColumnFamiliesColumnsArray.to_proto(
            resource.columns
        ):
            res.columns.extend(
                TableExternalDataConfigurationBigtableOptionsColumnFamiliesColumnsArray.to_proto(
                    resource.columns
                )
            )
        if Primitive.to_proto(resource.only_read_latest):
            res.only_read_latest = Primitive.to_proto(resource.only_read_latest)
        return res

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return None

        return TableExternalDataConfigurationBigtableOptionsColumnFamilies(
            family_id=Primitive.from_proto(resource.family_id),
            type=Primitive.from_proto(resource.type),
            encoding=Primitive.from_proto(resource.encoding),
            columns=TableExternalDataConfigurationBigtableOptionsColumnFamiliesColumnsArray.from_proto(
                resource.columns
            ),
            only_read_latest=Primitive.from_proto(resource.only_read_latest),
        )


class TableExternalDataConfigurationBigtableOptionsColumnFamiliesArray(object):
    @classmethod
    def to_proto(self, resources):
        if not resources:
            return resources
        return [
            TableExternalDataConfigurationBigtableOptionsColumnFamilies.to_proto(i)
            for i in resources
        ]

    @classmethod
    def from_proto(self, resources):
        return [
            TableExternalDataConfigurationBigtableOptionsColumnFamilies.from_proto(i)
            for i in resources
        ]


class TableExternalDataConfigurationBigtableOptionsColumnFamiliesColumns(object):
    def __init__(
        self,
        qualifier_encoded: str = None,
        qualifier_string: str = None,
        field_name: str = None,
        type: str = None,
        encoding: str = None,
        only_read_latest: bool = None,
    ):
        self.qualifier_encoded = qualifier_encoded
        self.qualifier_string = qualifier_string
        self.field_name = field_name
        self.type = type
        self.encoding = encoding
        self.only_read_latest = only_read_latest

    @classmethod
    def to_proto(self, resource):
        if not resource:
            return None

        res = (
            table_pb2.BigqueryAlphaTableExternalDataConfigurationBigtableOptionsColumnFamiliesColumns()
        )
        if Primitive.to_proto(resource.qualifier_encoded):
            res.qualifier_encoded = Primitive.to_proto(resource.qualifier_encoded)
        if Primitive.to_proto(resource.qualifier_string):
            res.qualifier_string = Primitive.to_proto(resource.qualifier_string)
        if Primitive.to_proto(resource.field_name):
            res.field_name = Primitive.to_proto(resource.field_name)
        if Primitive.to_proto(resource.type):
            res.type = Primitive.to_proto(resource.type)
        if Primitive.to_proto(resource.encoding):
            res.encoding = Primitive.to_proto(resource.encoding)
        if Primitive.to_proto(resource.only_read_latest):
            res.only_read_latest = Primitive.to_proto(resource.only_read_latest)
        return res

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return None

        return TableExternalDataConfigurationBigtableOptionsColumnFamiliesColumns(
            qualifier_encoded=Primitive.from_proto(resource.qualifier_encoded),
            qualifier_string=Primitive.from_proto(resource.qualifier_string),
            field_name=Primitive.from_proto(resource.field_name),
            type=Primitive.from_proto(resource.type),
            encoding=Primitive.from_proto(resource.encoding),
            only_read_latest=Primitive.from_proto(resource.only_read_latest),
        )


class TableExternalDataConfigurationBigtableOptionsColumnFamiliesColumnsArray(object):
    @classmethod
    def to_proto(self, resources):
        if not resources:
            return resources
        return [
            TableExternalDataConfigurationBigtableOptionsColumnFamiliesColumns.to_proto(
                i
            )
            for i in resources
        ]

    @classmethod
    def from_proto(self, resources):
        return [
            TableExternalDataConfigurationBigtableOptionsColumnFamiliesColumns.from_proto(
                i
            )
            for i in resources
        ]


class TableExternalDataConfigurationGoogleSheetsOptions(object):
    def __init__(self, skip_leading_rows: str = None, range: str = None):
        self.skip_leading_rows = skip_leading_rows
        self.range = range

    @classmethod
    def to_proto(self, resource):
        if not resource:
            return None

        res = table_pb2.BigqueryAlphaTableExternalDataConfigurationGoogleSheetsOptions()
        if Primitive.to_proto(resource.skip_leading_rows):
            res.skip_leading_rows = Primitive.to_proto(resource.skip_leading_rows)
        if Primitive.to_proto(resource.range):
            res.range = Primitive.to_proto(resource.range)
        return res

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return None

        return TableExternalDataConfigurationGoogleSheetsOptions(
            skip_leading_rows=Primitive.from_proto(resource.skip_leading_rows),
            range=Primitive.from_proto(resource.range),
        )


class TableExternalDataConfigurationGoogleSheetsOptionsArray(object):
    @classmethod
    def to_proto(self, resources):
        if not resources:
            return resources
        return [
            TableExternalDataConfigurationGoogleSheetsOptions.to_proto(i)
            for i in resources
        ]

    @classmethod
    def from_proto(self, resources):
        return [
            TableExternalDataConfigurationGoogleSheetsOptions.from_proto(i)
            for i in resources
        ]


class TableExternalDataConfigurationHivePartitioningOptions(object):
    def __init__(
        self,
        mode: str = None,
        source_uri_prefix: str = None,
        require_partition_filter: bool = None,
        fields: list = None,
    ):
        self.mode = mode
        self.source_uri_prefix = source_uri_prefix
        self.require_partition_filter = require_partition_filter
        self.fields = fields

    @classmethod
    def to_proto(self, resource):
        if not resource:
            return None

        res = (
            table_pb2.BigqueryAlphaTableExternalDataConfigurationHivePartitioningOptions()
        )
        if Primitive.to_proto(resource.mode):
            res.mode = Primitive.to_proto(resource.mode)
        if Primitive.to_proto(resource.source_uri_prefix):
            res.source_uri_prefix = Primitive.to_proto(resource.source_uri_prefix)
        if Primitive.to_proto(resource.require_partition_filter):
            res.require_partition_filter = Primitive.to_proto(
                resource.require_partition_filter
            )
        if Primitive.to_proto(resource.fields):
            res.fields.extend(Primitive.to_proto(resource.fields))
        return res

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return None

        return TableExternalDataConfigurationHivePartitioningOptions(
            mode=Primitive.from_proto(resource.mode),
            source_uri_prefix=Primitive.from_proto(resource.source_uri_prefix),
            require_partition_filter=Primitive.from_proto(
                resource.require_partition_filter
            ),
            fields=Primitive.from_proto(resource.fields),
        )


class TableExternalDataConfigurationHivePartitioningOptionsArray(object):
    @classmethod
    def to_proto(self, resources):
        if not resources:
            return resources
        return [
            TableExternalDataConfigurationHivePartitioningOptions.to_proto(i)
            for i in resources
        ]

    @classmethod
    def from_proto(self, resources):
        return [
            TableExternalDataConfigurationHivePartitioningOptions.from_proto(i)
            for i in resources
        ]


class TableExternalDataConfigurationValueConversionModes(object):
    def __init__(
        self,
        temporal_types_out_of_range_conversion_mode: str = None,
        numeric_type_out_of_range_conversion_mode: str = None,
    ):
        self.temporal_types_out_of_range_conversion_mode = (
            temporal_types_out_of_range_conversion_mode
        )
        self.numeric_type_out_of_range_conversion_mode = (
            numeric_type_out_of_range_conversion_mode
        )

    @classmethod
    def to_proto(self, resource):
        if not resource:
            return None

        res = (
            table_pb2.BigqueryAlphaTableExternalDataConfigurationValueConversionModes()
        )
        if TableExternalDataConfigurationValueConversionModesTemporalTypesOutOfRangeConversionModeEnum.to_proto(
            resource.temporal_types_out_of_range_conversion_mode
        ):
            res.temporal_types_out_of_range_conversion_mode = TableExternalDataConfigurationValueConversionModesTemporalTypesOutOfRangeConversionModeEnum.to_proto(
                resource.temporal_types_out_of_range_conversion_mode
            )
        if TableExternalDataConfigurationValueConversionModesNumericTypeOutOfRangeConversionModeEnum.to_proto(
            resource.numeric_type_out_of_range_conversion_mode
        ):
            res.numeric_type_out_of_range_conversion_mode = TableExternalDataConfigurationValueConversionModesNumericTypeOutOfRangeConversionModeEnum.to_proto(
                resource.numeric_type_out_of_range_conversion_mode
            )
        return res

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return None

        return TableExternalDataConfigurationValueConversionModes(
            temporal_types_out_of_range_conversion_mode=TableExternalDataConfigurationValueConversionModesTemporalTypesOutOfRangeConversionModeEnum.from_proto(
                resource.temporal_types_out_of_range_conversion_mode
            ),
            numeric_type_out_of_range_conversion_mode=TableExternalDataConfigurationValueConversionModesNumericTypeOutOfRangeConversionModeEnum.from_proto(
                resource.numeric_type_out_of_range_conversion_mode
            ),
        )


class TableExternalDataConfigurationValueConversionModesArray(object):
    @classmethod
    def to_proto(self, resources):
        if not resources:
            return resources
        return [
            TableExternalDataConfigurationValueConversionModes.to_proto(i)
            for i in resources
        ]

    @classmethod
    def from_proto(self, resources):
        return [
            TableExternalDataConfigurationValueConversionModes.from_proto(i)
            for i in resources
        ]


class TableExternalDataConfigurationAvroOptions(object):
    def __init__(self, use_avro_logical_types: bool = None):
        self.use_avro_logical_types = use_avro_logical_types

    @classmethod
    def to_proto(self, resource):
        if not resource:
            return None

        res = table_pb2.BigqueryAlphaTableExternalDataConfigurationAvroOptions()
        if Primitive.to_proto(resource.use_avro_logical_types):
            res.use_avro_logical_types = Primitive.to_proto(
                resource.use_avro_logical_types
            )
        return res

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return None

        return TableExternalDataConfigurationAvroOptions(
            use_avro_logical_types=Primitive.from_proto(
                resource.use_avro_logical_types
            ),
        )


class TableExternalDataConfigurationAvroOptionsArray(object):
    @classmethod
    def to_proto(self, resources):
        if not resources:
            return resources
        return [
            TableExternalDataConfigurationAvroOptions.to_proto(i) for i in resources
        ]

    @classmethod
    def from_proto(self, resources):
        return [
            TableExternalDataConfigurationAvroOptions.from_proto(i) for i in resources
        ]


class TableExternalDataConfigurationParquetOptions(object):
    def __init__(self, enum_as_string: bool = None, enable_list_inference: bool = None):
        self.enum_as_string = enum_as_string
        self.enable_list_inference = enable_list_inference

    @classmethod
    def to_proto(self, resource):
        if not resource:
            return None

        res = table_pb2.BigqueryAlphaTableExternalDataConfigurationParquetOptions()
        if Primitive.to_proto(resource.enum_as_string):
            res.enum_as_string = Primitive.to_proto(resource.enum_as_string)
        if Primitive.to_proto(resource.enable_list_inference):
            res.enable_list_inference = Primitive.to_proto(
                resource.enable_list_inference
            )
        return res

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return None

        return TableExternalDataConfigurationParquetOptions(
            enum_as_string=Primitive.from_proto(resource.enum_as_string),
            enable_list_inference=Primitive.from_proto(resource.enable_list_inference),
        )


class TableExternalDataConfigurationParquetOptionsArray(object):
    @classmethod
    def to_proto(self, resources):
        if not resources:
            return resources
        return [
            TableExternalDataConfigurationParquetOptions.to_proto(i) for i in resources
        ]

    @classmethod
    def from_proto(self, resources):
        return [
            TableExternalDataConfigurationParquetOptions.from_proto(i)
            for i in resources
        ]


class TableStreamingBuffer(object):
    def __init__(
        self,
        estimated_bytes: int = None,
        estimated_rows: int = None,
        oldest_entry_time: int = None,
    ):
        self.estimated_bytes = estimated_bytes
        self.estimated_rows = estimated_rows
        self.oldest_entry_time = oldest_entry_time

    @classmethod
    def to_proto(self, resource):
        if not resource:
            return None

        res = table_pb2.BigqueryAlphaTableStreamingBuffer()
        if Primitive.to_proto(resource.estimated_bytes):
            res.estimated_bytes = Primitive.to_proto(resource.estimated_bytes)
        if Primitive.to_proto(resource.estimated_rows):
            res.estimated_rows = Primitive.to_proto(resource.estimated_rows)
        if Primitive.to_proto(resource.oldest_entry_time):
            res.oldest_entry_time = Primitive.to_proto(resource.oldest_entry_time)
        return res

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return None

        return TableStreamingBuffer(
            estimated_bytes=Primitive.from_proto(resource.estimated_bytes),
            estimated_rows=Primitive.from_proto(resource.estimated_rows),
            oldest_entry_time=Primitive.from_proto(resource.oldest_entry_time),
        )


class TableStreamingBufferArray(object):
    @classmethod
    def to_proto(self, resources):
        if not resources:
            return resources
        return [TableStreamingBuffer.to_proto(i) for i in resources]

    @classmethod
    def from_proto(self, resources):
        return [TableStreamingBuffer.from_proto(i) for i in resources]


class TableEncryptionConfiguration(object):
    def __init__(self, kms_key_name: str = None):
        self.kms_key_name = kms_key_name

    @classmethod
    def to_proto(self, resource):
        if not resource:
            return None

        res = table_pb2.BigqueryAlphaTableEncryptionConfiguration()
        if Primitive.to_proto(resource.kms_key_name):
            res.kms_key_name = Primitive.to_proto(resource.kms_key_name)
        return res

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return None

        return TableEncryptionConfiguration(
            kms_key_name=Primitive.from_proto(resource.kms_key_name),
        )


class TableEncryptionConfigurationArray(object):
    @classmethod
    def to_proto(self, resources):
        if not resources:
            return resources
        return [TableEncryptionConfiguration.to_proto(i) for i in resources]

    @classmethod
    def from_proto(self, resources):
        return [TableEncryptionConfiguration.from_proto(i) for i in resources]


class TableSnapshotDefinition(object):
    def __init__(
        self,
        table: str = None,
        dataset: str = None,
        project: str = None,
        snapshot_time: str = None,
    ):
        self.table = table
        self.dataset = dataset
        self.project = project
        self.snapshot_time = snapshot_time

    @classmethod
    def to_proto(self, resource):
        if not resource:
            return None

        res = table_pb2.BigqueryAlphaTableSnapshotDefinition()
        if Primitive.to_proto(resource.table):
            res.table = Primitive.to_proto(resource.table)
        if Primitive.to_proto(resource.dataset):
            res.dataset = Primitive.to_proto(resource.dataset)
        if Primitive.to_proto(resource.project):
            res.project = Primitive.to_proto(resource.project)
        if Primitive.to_proto(resource.snapshot_time):
            res.snapshot_time = Primitive.to_proto(resource.snapshot_time)
        return res

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return None

        return TableSnapshotDefinition(
            table=Primitive.from_proto(resource.table),
            dataset=Primitive.from_proto(resource.dataset),
            project=Primitive.from_proto(resource.project),
            snapshot_time=Primitive.from_proto(resource.snapshot_time),
        )


class TableSnapshotDefinitionArray(object):
    @classmethod
    def to_proto(self, resources):
        if not resources:
            return resources
        return [TableSnapshotDefinition.to_proto(i) for i in resources]

    @classmethod
    def from_proto(self, resources):
        return [TableSnapshotDefinition.from_proto(i) for i in resources]


class TableExternalDataConfigurationValueConversionModesTemporalTypesOutOfRangeConversionModeEnum(
    object
):
    @classmethod
    def to_proto(self, resource):
        if not resource:
            return resource
        return table_pb2.BigqueryAlphaTableExternalDataConfigurationValueConversionModesTemporalTypesOutOfRangeConversionModeEnum.Value(
            "BigqueryAlphaTableExternalDataConfigurationValueConversionModesTemporalTypesOutOfRangeConversionModeEnum%s"
            % resource
        )

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return resource
        return table_pb2.BigqueryAlphaTableExternalDataConfigurationValueConversionModesTemporalTypesOutOfRangeConversionModeEnum.Name(
            resource
        )[
            len(
                "BigqueryAlphaTableExternalDataConfigurationValueConversionModesTemporalTypesOutOfRangeConversionModeEnum"
            ) :
        ]


class TableExternalDataConfigurationValueConversionModesNumericTypeOutOfRangeConversionModeEnum(
    object
):
    @classmethod
    def to_proto(self, resource):
        if not resource:
            return resource
        return table_pb2.BigqueryAlphaTableExternalDataConfigurationValueConversionModesNumericTypeOutOfRangeConversionModeEnum.Value(
            "BigqueryAlphaTableExternalDataConfigurationValueConversionModesNumericTypeOutOfRangeConversionModeEnum%s"
            % resource
        )

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return resource
        return table_pb2.BigqueryAlphaTableExternalDataConfigurationValueConversionModesNumericTypeOutOfRangeConversionModeEnum.Name(
            resource
        )[
            len(
                "BigqueryAlphaTableExternalDataConfigurationValueConversionModesNumericTypeOutOfRangeConversionModeEnum"
            ) :
        ]


class TableExternalDataConfigurationDecimalTargetTypesEnum(object):
    @classmethod
    def to_proto(self, resource):
        if not resource:
            return resource
        return table_pb2.BigqueryAlphaTableExternalDataConfigurationDecimalTargetTypesEnum.Value(
            "BigqueryAlphaTableExternalDataConfigurationDecimalTargetTypesEnum%s"
            % resource
        )

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return resource
        return table_pb2.BigqueryAlphaTableExternalDataConfigurationDecimalTargetTypesEnum.Name(
            resource
        )[
            len("BigqueryAlphaTableExternalDataConfigurationDecimalTargetTypesEnum") :
        ]


class TableExternalDataConfigurationJsonExtensionEnum(object):
    @classmethod
    def to_proto(self, resource):
        if not resource:
            return resource
        return table_pb2.BigqueryAlphaTableExternalDataConfigurationJsonExtensionEnum.Value(
            "BigqueryAlphaTableExternalDataConfigurationJsonExtensionEnum%s" % resource
        )

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return resource
        return table_pb2.BigqueryAlphaTableExternalDataConfigurationJsonExtensionEnum.Name(
            resource
        )[
            len("BigqueryAlphaTableExternalDataConfigurationJsonExtensionEnum") :
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
