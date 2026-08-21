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
from google3.cloud.graphite.mmv2.services.google.monitoring import dashboard_pb2
from google3.cloud.graphite.mmv2.services.google.monitoring import dashboard_pb2_grpc

from typing import List


class Dashboard(object):
    def __init__(
        self,
        name: str = None,
        display_name: str = None,
        grid_layout: dict = None,
        mosaic_layout: dict = None,
        row_layout: dict = None,
        column_layout: dict = None,
        project: str = None,
        etag: str = None,
        service_account_file: str = "",
    ):
        channel.initialize()
        self.name = name
        self.display_name = display_name
        self.grid_layout = grid_layout
        self.mosaic_layout = mosaic_layout
        self.row_layout = row_layout
        self.column_layout = column_layout
        self.project = project
        self.service_account_file = service_account_file

    def apply(self):
        stub = dashboard_pb2_grpc.MonitoringBetaDashboardServiceStub(channel.Channel())
        request = dashboard_pb2.ApplyMonitoringBetaDashboardRequest()
        if Primitive.to_proto(self.name):
            request.resource.name = Primitive.to_proto(self.name)

        if Primitive.to_proto(self.display_name):
            request.resource.display_name = Primitive.to_proto(self.display_name)

        if DashboardGridLayout.to_proto(self.grid_layout):
            request.resource.grid_layout.CopyFrom(
                DashboardGridLayout.to_proto(self.grid_layout)
            )
        else:
            request.resource.ClearField("grid_layout")
        if DashboardMosaicLayout.to_proto(self.mosaic_layout):
            request.resource.mosaic_layout.CopyFrom(
                DashboardMosaicLayout.to_proto(self.mosaic_layout)
            )
        else:
            request.resource.ClearField("mosaic_layout")
        if DashboardRowLayout.to_proto(self.row_layout):
            request.resource.row_layout.CopyFrom(
                DashboardRowLayout.to_proto(self.row_layout)
            )
        else:
            request.resource.ClearField("row_layout")
        if DashboardColumnLayout.to_proto(self.column_layout):
            request.resource.column_layout.CopyFrom(
                DashboardColumnLayout.to_proto(self.column_layout)
            )
        else:
            request.resource.ClearField("column_layout")
        if Primitive.to_proto(self.project):
            request.resource.project = Primitive.to_proto(self.project)

        request.service_account_file = self.service_account_file

        response = stub.ApplyMonitoringBetaDashboard(request)
        self.name = Primitive.from_proto(response.name)
        self.display_name = Primitive.from_proto(response.display_name)
        self.grid_layout = DashboardGridLayout.from_proto(response.grid_layout)
        self.mosaic_layout = DashboardMosaicLayout.from_proto(response.mosaic_layout)
        self.row_layout = DashboardRowLayout.from_proto(response.row_layout)
        self.column_layout = DashboardColumnLayout.from_proto(response.column_layout)
        self.project = Primitive.from_proto(response.project)
        self.etag = Primitive.from_proto(response.etag)

    def delete(self):
        stub = dashboard_pb2_grpc.MonitoringBetaDashboardServiceStub(channel.Channel())
        request = dashboard_pb2.DeleteMonitoringBetaDashboardRequest()
        request.service_account_file = self.service_account_file
        if Primitive.to_proto(self.name):
            request.resource.name = Primitive.to_proto(self.name)

        if Primitive.to_proto(self.display_name):
            request.resource.display_name = Primitive.to_proto(self.display_name)

        if DashboardGridLayout.to_proto(self.grid_layout):
            request.resource.grid_layout.CopyFrom(
                DashboardGridLayout.to_proto(self.grid_layout)
            )
        else:
            request.resource.ClearField("grid_layout")
        if DashboardMosaicLayout.to_proto(self.mosaic_layout):
            request.resource.mosaic_layout.CopyFrom(
                DashboardMosaicLayout.to_proto(self.mosaic_layout)
            )
        else:
            request.resource.ClearField("mosaic_layout")
        if DashboardRowLayout.to_proto(self.row_layout):
            request.resource.row_layout.CopyFrom(
                DashboardRowLayout.to_proto(self.row_layout)
            )
        else:
            request.resource.ClearField("row_layout")
        if DashboardColumnLayout.to_proto(self.column_layout):
            request.resource.column_layout.CopyFrom(
                DashboardColumnLayout.to_proto(self.column_layout)
            )
        else:
            request.resource.ClearField("column_layout")
        if Primitive.to_proto(self.project):
            request.resource.project = Primitive.to_proto(self.project)

        response = stub.DeleteMonitoringBetaDashboard(request)

    @classmethod
    def list(self, project, service_account_file=""):
        stub = dashboard_pb2_grpc.MonitoringBetaDashboardServiceStub(channel.Channel())
        request = dashboard_pb2.ListMonitoringBetaDashboardRequest()
        request.service_account_file = service_account_file
        request.Project = project

        return stub.ListMonitoringBetaDashboard(request).items

    def to_proto(self):
        resource = dashboard_pb2.MonitoringBetaDashboard()
        if Primitive.to_proto(self.name):
            resource.name = Primitive.to_proto(self.name)
        if Primitive.to_proto(self.display_name):
            resource.display_name = Primitive.to_proto(self.display_name)
        if DashboardGridLayout.to_proto(self.grid_layout):
            resource.grid_layout.CopyFrom(
                DashboardGridLayout.to_proto(self.grid_layout)
            )
        else:
            resource.ClearField("grid_layout")
        if DashboardMosaicLayout.to_proto(self.mosaic_layout):
            resource.mosaic_layout.CopyFrom(
                DashboardMosaicLayout.to_proto(self.mosaic_layout)
            )
        else:
            resource.ClearField("mosaic_layout")
        if DashboardRowLayout.to_proto(self.row_layout):
            resource.row_layout.CopyFrom(DashboardRowLayout.to_proto(self.row_layout))
        else:
            resource.ClearField("row_layout")
        if DashboardColumnLayout.to_proto(self.column_layout):
            resource.column_layout.CopyFrom(
                DashboardColumnLayout.to_proto(self.column_layout)
            )
        else:
            resource.ClearField("column_layout")
        if Primitive.to_proto(self.project):
            resource.project = Primitive.to_proto(self.project)
        return resource


class DashboardGridLayout(object):
    def __init__(self, columns: int = None, widgets: list = None):
        self.columns = columns
        self.widgets = widgets

    @classmethod
    def to_proto(self, resource):
        if not resource:
            return None

        res = dashboard_pb2.MonitoringBetaDashboardGridLayout()
        if Primitive.to_proto(resource.columns):
            res.columns = Primitive.to_proto(resource.columns)
        if DashboardGridLayoutWidgetsArray.to_proto(resource.widgets):
            res.widgets.extend(
                DashboardGridLayoutWidgetsArray.to_proto(resource.widgets)
            )
        return res

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return None

        return DashboardGridLayout(
            columns=Primitive.from_proto(resource.columns),
            widgets=DashboardGridLayoutWidgetsArray.from_proto(resource.widgets),
        )


class DashboardGridLayoutArray(object):
    @classmethod
    def to_proto(self, resources):
        if not resources:
            return resources
        return [DashboardGridLayout.to_proto(i) for i in resources]

    @classmethod
    def from_proto(self, resources):
        return [DashboardGridLayout.from_proto(i) for i in resources]


class DashboardGridLayoutWidgets(object):
    def __init__(
        self,
        title: str = None,
        xy_chart: dict = None,
        scorecard: dict = None,
        text: dict = None,
        blank: dict = None,
        logs_panel: dict = None,
    ):
        self.title = title
        self.xy_chart = xy_chart
        self.scorecard = scorecard
        self.text = text
        self.blank = blank
        self.logs_panel = logs_panel

    @classmethod
    def to_proto(self, resource):
        if not resource:
            return None

        res = dashboard_pb2.MonitoringBetaDashboardGridLayoutWidgets()
        if Primitive.to_proto(resource.title):
            res.title = Primitive.to_proto(resource.title)
        if DashboardGridLayoutWidgetsXyChart.to_proto(resource.xy_chart):
            res.xy_chart.CopyFrom(
                DashboardGridLayoutWidgetsXyChart.to_proto(resource.xy_chart)
            )
        else:
            res.ClearField("xy_chart")
        if DashboardGridLayoutWidgetsScorecard.to_proto(resource.scorecard):
            res.scorecard.CopyFrom(
                DashboardGridLayoutWidgetsScorecard.to_proto(resource.scorecard)
            )
        else:
            res.ClearField("scorecard")
        if DashboardGridLayoutWidgetsText.to_proto(resource.text):
            res.text.CopyFrom(DashboardGridLayoutWidgetsText.to_proto(resource.text))
        else:
            res.ClearField("text")
        if DashboardGridLayoutWidgetsBlank.to_proto(resource.blank):
            res.blank.CopyFrom(DashboardGridLayoutWidgetsBlank.to_proto(resource.blank))
        else:
            res.ClearField("blank")
        if DashboardGridLayoutWidgetsLogsPanel.to_proto(resource.logs_panel):
            res.logs_panel.CopyFrom(
                DashboardGridLayoutWidgetsLogsPanel.to_proto(resource.logs_panel)
            )
        else:
            res.ClearField("logs_panel")
        return res

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return None

        return DashboardGridLayoutWidgets(
            title=Primitive.from_proto(resource.title),
            xy_chart=DashboardGridLayoutWidgetsXyChart.from_proto(resource.xy_chart),
            scorecard=DashboardGridLayoutWidgetsScorecard.from_proto(
                resource.scorecard
            ),
            text=DashboardGridLayoutWidgetsText.from_proto(resource.text),
            blank=DashboardGridLayoutWidgetsBlank.from_proto(resource.blank),
            logs_panel=DashboardGridLayoutWidgetsLogsPanel.from_proto(
                resource.logs_panel
            ),
        )


class DashboardGridLayoutWidgetsArray(object):
    @classmethod
    def to_proto(self, resources):
        if not resources:
            return resources
        return [DashboardGridLayoutWidgets.to_proto(i) for i in resources]

    @classmethod
    def from_proto(self, resources):
        return [DashboardGridLayoutWidgets.from_proto(i) for i in resources]


class DashboardGridLayoutWidgetsXyChart(object):
    def __init__(
        self,
        data_sets: list = None,
        timeshift_duration: str = None,
        thresholds: list = None,
        x_axis: dict = None,
        y_axis: dict = None,
        chart_options: dict = None,
    ):
        self.data_sets = data_sets
        self.timeshift_duration = timeshift_duration
        self.thresholds = thresholds
        self.x_axis = x_axis
        self.y_axis = y_axis
        self.chart_options = chart_options

    @classmethod
    def to_proto(self, resource):
        if not resource:
            return None

        res = dashboard_pb2.MonitoringBetaDashboardGridLayoutWidgetsXyChart()
        if DashboardGridLayoutWidgetsXyChartDataSetsArray.to_proto(resource.data_sets):
            res.data_sets.extend(
                DashboardGridLayoutWidgetsXyChartDataSetsArray.to_proto(
                    resource.data_sets
                )
            )
        if Primitive.to_proto(resource.timeshift_duration):
            res.timeshift_duration = Primitive.to_proto(resource.timeshift_duration)
        if DashboardGridLayoutWidgetsXyChartThresholdsArray.to_proto(
            resource.thresholds
        ):
            res.thresholds.extend(
                DashboardGridLayoutWidgetsXyChartThresholdsArray.to_proto(
                    resource.thresholds
                )
            )
        if DashboardGridLayoutWidgetsXyChartXAxis.to_proto(resource.x_axis):
            res.x_axis.CopyFrom(
                DashboardGridLayoutWidgetsXyChartXAxis.to_proto(resource.x_axis)
            )
        else:
            res.ClearField("x_axis")
        if DashboardGridLayoutWidgetsXyChartYAxis.to_proto(resource.y_axis):
            res.y_axis.CopyFrom(
                DashboardGridLayoutWidgetsXyChartYAxis.to_proto(resource.y_axis)
            )
        else:
            res.ClearField("y_axis")
        if DashboardGridLayoutWidgetsXyChartChartOptions.to_proto(
            resource.chart_options
        ):
            res.chart_options.CopyFrom(
                DashboardGridLayoutWidgetsXyChartChartOptions.to_proto(
                    resource.chart_options
                )
            )
        else:
            res.ClearField("chart_options")
        return res

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return None

        return DashboardGridLayoutWidgetsXyChart(
            data_sets=DashboardGridLayoutWidgetsXyChartDataSetsArray.from_proto(
                resource.data_sets
            ),
            timeshift_duration=Primitive.from_proto(resource.timeshift_duration),
            thresholds=DashboardGridLayoutWidgetsXyChartThresholdsArray.from_proto(
                resource.thresholds
            ),
            x_axis=DashboardGridLayoutWidgetsXyChartXAxis.from_proto(resource.x_axis),
            y_axis=DashboardGridLayoutWidgetsXyChartYAxis.from_proto(resource.y_axis),
            chart_options=DashboardGridLayoutWidgetsXyChartChartOptions.from_proto(
                resource.chart_options
            ),
        )


class DashboardGridLayoutWidgetsXyChartArray(object):
    @classmethod
    def to_proto(self, resources):
        if not resources:
            return resources
        return [DashboardGridLayoutWidgetsXyChart.to_proto(i) for i in resources]

    @classmethod
    def from_proto(self, resources):
        return [DashboardGridLayoutWidgetsXyChart.from_proto(i) for i in resources]


class DashboardGridLayoutWidgetsXyChartDataSets(object):
    def __init__(
        self,
        time_series_query: dict = None,
        plot_type: str = None,
        legend_template: str = None,
        min_alignment_period: str = None,
    ):
        self.time_series_query = time_series_query
        self.plot_type = plot_type
        self.legend_template = legend_template
        self.min_alignment_period = min_alignment_period

    @classmethod
    def to_proto(self, resource):
        if not resource:
            return None

        res = dashboard_pb2.MonitoringBetaDashboardGridLayoutWidgetsXyChartDataSets()
        if DashboardGridLayoutWidgetsXyChartDataSetsTimeSeriesQuery.to_proto(
            resource.time_series_query
        ):
            res.time_series_query.CopyFrom(
                DashboardGridLayoutWidgetsXyChartDataSetsTimeSeriesQuery.to_proto(
                    resource.time_series_query
                )
            )
        else:
            res.ClearField("time_series_query")
        if DashboardGridLayoutWidgetsXyChartDataSetsPlotTypeEnum.to_proto(
            resource.plot_type
        ):
            res.plot_type = (
                DashboardGridLayoutWidgetsXyChartDataSetsPlotTypeEnum.to_proto(
                    resource.plot_type
                )
            )
        if Primitive.to_proto(resource.legend_template):
            res.legend_template = Primitive.to_proto(resource.legend_template)
        if Primitive.to_proto(resource.min_alignment_period):
            res.min_alignment_period = Primitive.to_proto(resource.min_alignment_period)
        return res

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return None

        return DashboardGridLayoutWidgetsXyChartDataSets(
            time_series_query=DashboardGridLayoutWidgetsXyChartDataSetsTimeSeriesQuery.from_proto(
                resource.time_series_query
            ),
            plot_type=DashboardGridLayoutWidgetsXyChartDataSetsPlotTypeEnum.from_proto(
                resource.plot_type
            ),
            legend_template=Primitive.from_proto(resource.legend_template),
            min_alignment_period=Primitive.from_proto(resource.min_alignment_period),
        )


class DashboardGridLayoutWidgetsXyChartDataSetsArray(object):
    @classmethod
    def to_proto(self, resources):
        if not resources:
            return resources
        return [
            DashboardGridLayoutWidgetsXyChartDataSets.to_proto(i) for i in resources
        ]

    @classmethod
    def from_proto(self, resources):
        return [
            DashboardGridLayoutWidgetsXyChartDataSets.from_proto(i) for i in resources
        ]


class DashboardGridLayoutWidgetsXyChartDataSetsTimeSeriesQuery(object):
    def __init__(
        self,
        time_series_filter: dict = None,
        time_series_filter_ratio: dict = None,
        time_series_query_language: str = None,
        unit_override: str = None,
    ):
        self.time_series_filter = time_series_filter
        self.time_series_filter_ratio = time_series_filter_ratio
        self.time_series_query_language = time_series_query_language
        self.unit_override = unit_override

    @classmethod
    def to_proto(self, resource):
        if not resource:
            return None

        res = (
            dashboard_pb2.MonitoringBetaDashboardGridLayoutWidgetsXyChartDataSetsTimeSeriesQuery()
        )
        if DashboardGridLayoutWidgetsXyChartDataSetsTimeSeriesQueryTimeSeriesFilter.to_proto(
            resource.time_series_filter
        ):
            res.time_series_filter.CopyFrom(
                DashboardGridLayoutWidgetsXyChartDataSetsTimeSeriesQueryTimeSeriesFilter.to_proto(
                    resource.time_series_filter
                )
            )
        else:
            res.ClearField("time_series_filter")
        if DashboardGridLayoutWidgetsXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatio.to_proto(
            resource.time_series_filter_ratio
        ):
            res.time_series_filter_ratio.CopyFrom(
                DashboardGridLayoutWidgetsXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatio.to_proto(
                    resource.time_series_filter_ratio
                )
            )
        else:
            res.ClearField("time_series_filter_ratio")
        if Primitive.to_proto(resource.time_series_query_language):
            res.time_series_query_language = Primitive.to_proto(
                resource.time_series_query_language
            )
        if Primitive.to_proto(resource.unit_override):
            res.unit_override = Primitive.to_proto(resource.unit_override)
        return res

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return None

        return DashboardGridLayoutWidgetsXyChartDataSetsTimeSeriesQuery(
            time_series_filter=DashboardGridLayoutWidgetsXyChartDataSetsTimeSeriesQueryTimeSeriesFilter.from_proto(
                resource.time_series_filter
            ),
            time_series_filter_ratio=DashboardGridLayoutWidgetsXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatio.from_proto(
                resource.time_series_filter_ratio
            ),
            time_series_query_language=Primitive.from_proto(
                resource.time_series_query_language
            ),
            unit_override=Primitive.from_proto(resource.unit_override),
        )


class DashboardGridLayoutWidgetsXyChartDataSetsTimeSeriesQueryArray(object):
    @classmethod
    def to_proto(self, resources):
        if not resources:
            return resources
        return [
            DashboardGridLayoutWidgetsXyChartDataSetsTimeSeriesQuery.to_proto(i)
            for i in resources
        ]

    @classmethod
    def from_proto(self, resources):
        return [
            DashboardGridLayoutWidgetsXyChartDataSetsTimeSeriesQuery.from_proto(i)
            for i in resources
        ]


class DashboardGridLayoutWidgetsXyChartDataSetsTimeSeriesQueryTimeSeriesFilter(object):
    def __init__(
        self,
        filter: str = None,
        aggregation: dict = None,
        secondary_aggregation: dict = None,
        pick_time_series_filter: dict = None,
    ):
        self.filter = filter
        self.aggregation = aggregation
        self.secondary_aggregation = secondary_aggregation
        self.pick_time_series_filter = pick_time_series_filter

    @classmethod
    def to_proto(self, resource):
        if not resource:
            return None

        res = (
            dashboard_pb2.MonitoringBetaDashboardGridLayoutWidgetsXyChartDataSetsTimeSeriesQueryTimeSeriesFilter()
        )
        if Primitive.to_proto(resource.filter):
            res.filter = Primitive.to_proto(resource.filter)
        if DashboardGridLayoutWidgetsXyChartDataSetsTimeSeriesQueryTimeSeriesFilterAggregation.to_proto(
            resource.aggregation
        ):
            res.aggregation.CopyFrom(
                DashboardGridLayoutWidgetsXyChartDataSetsTimeSeriesQueryTimeSeriesFilterAggregation.to_proto(
                    resource.aggregation
                )
            )
        else:
            res.ClearField("aggregation")
        if DashboardGridLayoutWidgetsXyChartDataSetsTimeSeriesQueryTimeSeriesFilterSecondaryAggregation.to_proto(
            resource.secondary_aggregation
        ):
            res.secondary_aggregation.CopyFrom(
                DashboardGridLayoutWidgetsXyChartDataSetsTimeSeriesQueryTimeSeriesFilterSecondaryAggregation.to_proto(
                    resource.secondary_aggregation
                )
            )
        else:
            res.ClearField("secondary_aggregation")
        if DashboardGridLayoutWidgetsXyChartDataSetsTimeSeriesQueryTimeSeriesFilterPickTimeSeriesFilter.to_proto(
            resource.pick_time_series_filter
        ):
            res.pick_time_series_filter.CopyFrom(
                DashboardGridLayoutWidgetsXyChartDataSetsTimeSeriesQueryTimeSeriesFilterPickTimeSeriesFilter.to_proto(
                    resource.pick_time_series_filter
                )
            )
        else:
            res.ClearField("pick_time_series_filter")
        return res

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return None

        return DashboardGridLayoutWidgetsXyChartDataSetsTimeSeriesQueryTimeSeriesFilter(
            filter=Primitive.from_proto(resource.filter),
            aggregation=DashboardGridLayoutWidgetsXyChartDataSetsTimeSeriesQueryTimeSeriesFilterAggregation.from_proto(
                resource.aggregation
            ),
            secondary_aggregation=DashboardGridLayoutWidgetsXyChartDataSetsTimeSeriesQueryTimeSeriesFilterSecondaryAggregation.from_proto(
                resource.secondary_aggregation
            ),
            pick_time_series_filter=DashboardGridLayoutWidgetsXyChartDataSetsTimeSeriesQueryTimeSeriesFilterPickTimeSeriesFilter.from_proto(
                resource.pick_time_series_filter
            ),
        )


class DashboardGridLayoutWidgetsXyChartDataSetsTimeSeriesQueryTimeSeriesFilterArray(
    object
):
    @classmethod
    def to_proto(self, resources):
        if not resources:
            return resources
        return [
            DashboardGridLayoutWidgetsXyChartDataSetsTimeSeriesQueryTimeSeriesFilter.to_proto(
                i
            )
            for i in resources
        ]

    @classmethod
    def from_proto(self, resources):
        return [
            DashboardGridLayoutWidgetsXyChartDataSetsTimeSeriesQueryTimeSeriesFilter.from_proto(
                i
            )
            for i in resources
        ]


class DashboardGridLayoutWidgetsXyChartDataSetsTimeSeriesQueryTimeSeriesFilterAggregation(
    object
):
    def __init__(
        self,
        alignment_period: str = None,
        per_series_aligner: str = None,
        cross_series_reducer: str = None,
        group_by_fields: list = None,
    ):
        self.alignment_period = alignment_period
        self.per_series_aligner = per_series_aligner
        self.cross_series_reducer = cross_series_reducer
        self.group_by_fields = group_by_fields

    @classmethod
    def to_proto(self, resource):
        if not resource:
            return None

        res = (
            dashboard_pb2.MonitoringBetaDashboardGridLayoutWidgetsXyChartDataSetsTimeSeriesQueryTimeSeriesFilterAggregation()
        )
        if Primitive.to_proto(resource.alignment_period):
            res.alignment_period = Primitive.to_proto(resource.alignment_period)
        if DashboardGridLayoutWidgetsXyChartDataSetsTimeSeriesQueryTimeSeriesFilterAggregationPerSeriesAlignerEnum.to_proto(
            resource.per_series_aligner
        ):
            res.per_series_aligner = DashboardGridLayoutWidgetsXyChartDataSetsTimeSeriesQueryTimeSeriesFilterAggregationPerSeriesAlignerEnum.to_proto(
                resource.per_series_aligner
            )
        if DashboardGridLayoutWidgetsXyChartDataSetsTimeSeriesQueryTimeSeriesFilterAggregationCrossSeriesReducerEnum.to_proto(
            resource.cross_series_reducer
        ):
            res.cross_series_reducer = DashboardGridLayoutWidgetsXyChartDataSetsTimeSeriesQueryTimeSeriesFilterAggregationCrossSeriesReducerEnum.to_proto(
                resource.cross_series_reducer
            )
        if Primitive.to_proto(resource.group_by_fields):
            res.group_by_fields.extend(Primitive.to_proto(resource.group_by_fields))
        return res

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return None

        return DashboardGridLayoutWidgetsXyChartDataSetsTimeSeriesQueryTimeSeriesFilterAggregation(
            alignment_period=Primitive.from_proto(resource.alignment_period),
            per_series_aligner=DashboardGridLayoutWidgetsXyChartDataSetsTimeSeriesQueryTimeSeriesFilterAggregationPerSeriesAlignerEnum.from_proto(
                resource.per_series_aligner
            ),
            cross_series_reducer=DashboardGridLayoutWidgetsXyChartDataSetsTimeSeriesQueryTimeSeriesFilterAggregationCrossSeriesReducerEnum.from_proto(
                resource.cross_series_reducer
            ),
            group_by_fields=Primitive.from_proto(resource.group_by_fields),
        )


class DashboardGridLayoutWidgetsXyChartDataSetsTimeSeriesQueryTimeSeriesFilterAggregationArray(
    object
):
    @classmethod
    def to_proto(self, resources):
        if not resources:
            return resources
        return [
            DashboardGridLayoutWidgetsXyChartDataSetsTimeSeriesQueryTimeSeriesFilterAggregation.to_proto(
                i
            )
            for i in resources
        ]

    @classmethod
    def from_proto(self, resources):
        return [
            DashboardGridLayoutWidgetsXyChartDataSetsTimeSeriesQueryTimeSeriesFilterAggregation.from_proto(
                i
            )
            for i in resources
        ]


class DashboardGridLayoutWidgetsXyChartDataSetsTimeSeriesQueryTimeSeriesFilterSecondaryAggregation(
    object
):
    def __init__(
        self,
        alignment_period: str = None,
        per_series_aligner: str = None,
        cross_series_reducer: str = None,
        group_by_fields: list = None,
    ):
        self.alignment_period = alignment_period
        self.per_series_aligner = per_series_aligner
        self.cross_series_reducer = cross_series_reducer
        self.group_by_fields = group_by_fields

    @classmethod
    def to_proto(self, resource):
        if not resource:
            return None

        res = (
            dashboard_pb2.MonitoringBetaDashboardGridLayoutWidgetsXyChartDataSetsTimeSeriesQueryTimeSeriesFilterSecondaryAggregation()
        )
        if Primitive.to_proto(resource.alignment_period):
            res.alignment_period = Primitive.to_proto(resource.alignment_period)
        if DashboardGridLayoutWidgetsXyChartDataSetsTimeSeriesQueryTimeSeriesFilterSecondaryAggregationPerSeriesAlignerEnum.to_proto(
            resource.per_series_aligner
        ):
            res.per_series_aligner = DashboardGridLayoutWidgetsXyChartDataSetsTimeSeriesQueryTimeSeriesFilterSecondaryAggregationPerSeriesAlignerEnum.to_proto(
                resource.per_series_aligner
            )
        if DashboardGridLayoutWidgetsXyChartDataSetsTimeSeriesQueryTimeSeriesFilterSecondaryAggregationCrossSeriesReducerEnum.to_proto(
            resource.cross_series_reducer
        ):
            res.cross_series_reducer = DashboardGridLayoutWidgetsXyChartDataSetsTimeSeriesQueryTimeSeriesFilterSecondaryAggregationCrossSeriesReducerEnum.to_proto(
                resource.cross_series_reducer
            )
        if Primitive.to_proto(resource.group_by_fields):
            res.group_by_fields.extend(Primitive.to_proto(resource.group_by_fields))
        return res

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return None

        return DashboardGridLayoutWidgetsXyChartDataSetsTimeSeriesQueryTimeSeriesFilterSecondaryAggregation(
            alignment_period=Primitive.from_proto(resource.alignment_period),
            per_series_aligner=DashboardGridLayoutWidgetsXyChartDataSetsTimeSeriesQueryTimeSeriesFilterSecondaryAggregationPerSeriesAlignerEnum.from_proto(
                resource.per_series_aligner
            ),
            cross_series_reducer=DashboardGridLayoutWidgetsXyChartDataSetsTimeSeriesQueryTimeSeriesFilterSecondaryAggregationCrossSeriesReducerEnum.from_proto(
                resource.cross_series_reducer
            ),
            group_by_fields=Primitive.from_proto(resource.group_by_fields),
        )


class DashboardGridLayoutWidgetsXyChartDataSetsTimeSeriesQueryTimeSeriesFilterSecondaryAggregationArray(
    object
):
    @classmethod
    def to_proto(self, resources):
        if not resources:
            return resources
        return [
            DashboardGridLayoutWidgetsXyChartDataSetsTimeSeriesQueryTimeSeriesFilterSecondaryAggregation.to_proto(
                i
            )
            for i in resources
        ]

    @classmethod
    def from_proto(self, resources):
        return [
            DashboardGridLayoutWidgetsXyChartDataSetsTimeSeriesQueryTimeSeriesFilterSecondaryAggregation.from_proto(
                i
            )
            for i in resources
        ]


class DashboardGridLayoutWidgetsXyChartDataSetsTimeSeriesQueryTimeSeriesFilterPickTimeSeriesFilter(
    object
):
    def __init__(
        self,
        ranking_method: str = None,
        num_time_series: int = None,
        direction: str = None,
    ):
        self.ranking_method = ranking_method
        self.num_time_series = num_time_series
        self.direction = direction

    @classmethod
    def to_proto(self, resource):
        if not resource:
            return None

        res = (
            dashboard_pb2.MonitoringBetaDashboardGridLayoutWidgetsXyChartDataSetsTimeSeriesQueryTimeSeriesFilterPickTimeSeriesFilter()
        )
        if DashboardGridLayoutWidgetsXyChartDataSetsTimeSeriesQueryTimeSeriesFilterPickTimeSeriesFilterRankingMethodEnum.to_proto(
            resource.ranking_method
        ):
            res.ranking_method = DashboardGridLayoutWidgetsXyChartDataSetsTimeSeriesQueryTimeSeriesFilterPickTimeSeriesFilterRankingMethodEnum.to_proto(
                resource.ranking_method
            )
        if Primitive.to_proto(resource.num_time_series):
            res.num_time_series = Primitive.to_proto(resource.num_time_series)
        if DashboardGridLayoutWidgetsXyChartDataSetsTimeSeriesQueryTimeSeriesFilterPickTimeSeriesFilterDirectionEnum.to_proto(
            resource.direction
        ):
            res.direction = DashboardGridLayoutWidgetsXyChartDataSetsTimeSeriesQueryTimeSeriesFilterPickTimeSeriesFilterDirectionEnum.to_proto(
                resource.direction
            )
        return res

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return None

        return DashboardGridLayoutWidgetsXyChartDataSetsTimeSeriesQueryTimeSeriesFilterPickTimeSeriesFilter(
            ranking_method=DashboardGridLayoutWidgetsXyChartDataSetsTimeSeriesQueryTimeSeriesFilterPickTimeSeriesFilterRankingMethodEnum.from_proto(
                resource.ranking_method
            ),
            num_time_series=Primitive.from_proto(resource.num_time_series),
            direction=DashboardGridLayoutWidgetsXyChartDataSetsTimeSeriesQueryTimeSeriesFilterPickTimeSeriesFilterDirectionEnum.from_proto(
                resource.direction
            ),
        )


class DashboardGridLayoutWidgetsXyChartDataSetsTimeSeriesQueryTimeSeriesFilterPickTimeSeriesFilterArray(
    object
):
    @classmethod
    def to_proto(self, resources):
        if not resources:
            return resources
        return [
            DashboardGridLayoutWidgetsXyChartDataSetsTimeSeriesQueryTimeSeriesFilterPickTimeSeriesFilter.to_proto(
                i
            )
            for i in resources
        ]

    @classmethod
    def from_proto(self, resources):
        return [
            DashboardGridLayoutWidgetsXyChartDataSetsTimeSeriesQueryTimeSeriesFilterPickTimeSeriesFilter.from_proto(
                i
            )
            for i in resources
        ]


class DashboardGridLayoutWidgetsXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatio(
    object
):
    def __init__(
        self,
        numerator: dict = None,
        denominator: dict = None,
        secondary_aggregation: dict = None,
        pick_time_series_filter: dict = None,
    ):
        self.numerator = numerator
        self.denominator = denominator
        self.secondary_aggregation = secondary_aggregation
        self.pick_time_series_filter = pick_time_series_filter

    @classmethod
    def to_proto(self, resource):
        if not resource:
            return None

        res = (
            dashboard_pb2.MonitoringBetaDashboardGridLayoutWidgetsXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatio()
        )
        if DashboardGridLayoutWidgetsXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioNumerator.to_proto(
            resource.numerator
        ):
            res.numerator.CopyFrom(
                DashboardGridLayoutWidgetsXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioNumerator.to_proto(
                    resource.numerator
                )
            )
        else:
            res.ClearField("numerator")
        if DashboardGridLayoutWidgetsXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioDenominator.to_proto(
            resource.denominator
        ):
            res.denominator.CopyFrom(
                DashboardGridLayoutWidgetsXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioDenominator.to_proto(
                    resource.denominator
                )
            )
        else:
            res.ClearField("denominator")
        if DashboardGridLayoutWidgetsXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioSecondaryAggregation.to_proto(
            resource.secondary_aggregation
        ):
            res.secondary_aggregation.CopyFrom(
                DashboardGridLayoutWidgetsXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioSecondaryAggregation.to_proto(
                    resource.secondary_aggregation
                )
            )
        else:
            res.ClearField("secondary_aggregation")
        if DashboardGridLayoutWidgetsXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioPickTimeSeriesFilter.to_proto(
            resource.pick_time_series_filter
        ):
            res.pick_time_series_filter.CopyFrom(
                DashboardGridLayoutWidgetsXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioPickTimeSeriesFilter.to_proto(
                    resource.pick_time_series_filter
                )
            )
        else:
            res.ClearField("pick_time_series_filter")
        return res

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return None

        return DashboardGridLayoutWidgetsXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatio(
            numerator=DashboardGridLayoutWidgetsXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioNumerator.from_proto(
                resource.numerator
            ),
            denominator=DashboardGridLayoutWidgetsXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioDenominator.from_proto(
                resource.denominator
            ),
            secondary_aggregation=DashboardGridLayoutWidgetsXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioSecondaryAggregation.from_proto(
                resource.secondary_aggregation
            ),
            pick_time_series_filter=DashboardGridLayoutWidgetsXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioPickTimeSeriesFilter.from_proto(
                resource.pick_time_series_filter
            ),
        )


class DashboardGridLayoutWidgetsXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioArray(
    object
):
    @classmethod
    def to_proto(self, resources):
        if not resources:
            return resources
        return [
            DashboardGridLayoutWidgetsXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatio.to_proto(
                i
            )
            for i in resources
        ]

    @classmethod
    def from_proto(self, resources):
        return [
            DashboardGridLayoutWidgetsXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatio.from_proto(
                i
            )
            for i in resources
        ]


class DashboardGridLayoutWidgetsXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioNumerator(
    object
):
    def __init__(self, filter: str = None, aggregation: dict = None):
        self.filter = filter
        self.aggregation = aggregation

    @classmethod
    def to_proto(self, resource):
        if not resource:
            return None

        res = (
            dashboard_pb2.MonitoringBetaDashboardGridLayoutWidgetsXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioNumerator()
        )
        if Primitive.to_proto(resource.filter):
            res.filter = Primitive.to_proto(resource.filter)
        if DashboardGridLayoutWidgetsXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioNumeratorAggregation.to_proto(
            resource.aggregation
        ):
            res.aggregation.CopyFrom(
                DashboardGridLayoutWidgetsXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioNumeratorAggregation.to_proto(
                    resource.aggregation
                )
            )
        else:
            res.ClearField("aggregation")
        return res

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return None

        return DashboardGridLayoutWidgetsXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioNumerator(
            filter=Primitive.from_proto(resource.filter),
            aggregation=DashboardGridLayoutWidgetsXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioNumeratorAggregation.from_proto(
                resource.aggregation
            ),
        )


class DashboardGridLayoutWidgetsXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioNumeratorArray(
    object
):
    @classmethod
    def to_proto(self, resources):
        if not resources:
            return resources
        return [
            DashboardGridLayoutWidgetsXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioNumerator.to_proto(
                i
            )
            for i in resources
        ]

    @classmethod
    def from_proto(self, resources):
        return [
            DashboardGridLayoutWidgetsXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioNumerator.from_proto(
                i
            )
            for i in resources
        ]


class DashboardGridLayoutWidgetsXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioNumeratorAggregation(
    object
):
    def __init__(
        self,
        alignment_period: str = None,
        per_series_aligner: str = None,
        cross_series_reducer: str = None,
        group_by_fields: list = None,
    ):
        self.alignment_period = alignment_period
        self.per_series_aligner = per_series_aligner
        self.cross_series_reducer = cross_series_reducer
        self.group_by_fields = group_by_fields

    @classmethod
    def to_proto(self, resource):
        if not resource:
            return None

        res = (
            dashboard_pb2.MonitoringBetaDashboardGridLayoutWidgetsXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioNumeratorAggregation()
        )
        if Primitive.to_proto(resource.alignment_period):
            res.alignment_period = Primitive.to_proto(resource.alignment_period)
        if DashboardGridLayoutWidgetsXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioNumeratorAggregationPerSeriesAlignerEnum.to_proto(
            resource.per_series_aligner
        ):
            res.per_series_aligner = DashboardGridLayoutWidgetsXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioNumeratorAggregationPerSeriesAlignerEnum.to_proto(
                resource.per_series_aligner
            )
        if DashboardGridLayoutWidgetsXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioNumeratorAggregationCrossSeriesReducerEnum.to_proto(
            resource.cross_series_reducer
        ):
            res.cross_series_reducer = DashboardGridLayoutWidgetsXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioNumeratorAggregationCrossSeriesReducerEnum.to_proto(
                resource.cross_series_reducer
            )
        if Primitive.to_proto(resource.group_by_fields):
            res.group_by_fields.extend(Primitive.to_proto(resource.group_by_fields))
        return res

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return None

        return DashboardGridLayoutWidgetsXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioNumeratorAggregation(
            alignment_period=Primitive.from_proto(resource.alignment_period),
            per_series_aligner=DashboardGridLayoutWidgetsXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioNumeratorAggregationPerSeriesAlignerEnum.from_proto(
                resource.per_series_aligner
            ),
            cross_series_reducer=DashboardGridLayoutWidgetsXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioNumeratorAggregationCrossSeriesReducerEnum.from_proto(
                resource.cross_series_reducer
            ),
            group_by_fields=Primitive.from_proto(resource.group_by_fields),
        )


class DashboardGridLayoutWidgetsXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioNumeratorAggregationArray(
    object
):
    @classmethod
    def to_proto(self, resources):
        if not resources:
            return resources
        return [
            DashboardGridLayoutWidgetsXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioNumeratorAggregation.to_proto(
                i
            )
            for i in resources
        ]

    @classmethod
    def from_proto(self, resources):
        return [
            DashboardGridLayoutWidgetsXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioNumeratorAggregation.from_proto(
                i
            )
            for i in resources
        ]


class DashboardGridLayoutWidgetsXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioDenominator(
    object
):
    def __init__(self, filter: str = None, aggregation: dict = None):
        self.filter = filter
        self.aggregation = aggregation

    @classmethod
    def to_proto(self, resource):
        if not resource:
            return None

        res = (
            dashboard_pb2.MonitoringBetaDashboardGridLayoutWidgetsXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioDenominator()
        )
        if Primitive.to_proto(resource.filter):
            res.filter = Primitive.to_proto(resource.filter)
        if DashboardGridLayoutWidgetsXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioDenominatorAggregation.to_proto(
            resource.aggregation
        ):
            res.aggregation.CopyFrom(
                DashboardGridLayoutWidgetsXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioDenominatorAggregation.to_proto(
                    resource.aggregation
                )
            )
        else:
            res.ClearField("aggregation")
        return res

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return None

        return DashboardGridLayoutWidgetsXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioDenominator(
            filter=Primitive.from_proto(resource.filter),
            aggregation=DashboardGridLayoutWidgetsXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioDenominatorAggregation.from_proto(
                resource.aggregation
            ),
        )


class DashboardGridLayoutWidgetsXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioDenominatorArray(
    object
):
    @classmethod
    def to_proto(self, resources):
        if not resources:
            return resources
        return [
            DashboardGridLayoutWidgetsXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioDenominator.to_proto(
                i
            )
            for i in resources
        ]

    @classmethod
    def from_proto(self, resources):
        return [
            DashboardGridLayoutWidgetsXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioDenominator.from_proto(
                i
            )
            for i in resources
        ]


class DashboardGridLayoutWidgetsXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioDenominatorAggregation(
    object
):
    def __init__(
        self,
        alignment_period: str = None,
        per_series_aligner: str = None,
        cross_series_reducer: str = None,
        group_by_fields: list = None,
    ):
        self.alignment_period = alignment_period
        self.per_series_aligner = per_series_aligner
        self.cross_series_reducer = cross_series_reducer
        self.group_by_fields = group_by_fields

    @classmethod
    def to_proto(self, resource):
        if not resource:
            return None

        res = (
            dashboard_pb2.MonitoringBetaDashboardGridLayoutWidgetsXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioDenominatorAggregation()
        )
        if Primitive.to_proto(resource.alignment_period):
            res.alignment_period = Primitive.to_proto(resource.alignment_period)
        if DashboardGridLayoutWidgetsXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioDenominatorAggregationPerSeriesAlignerEnum.to_proto(
            resource.per_series_aligner
        ):
            res.per_series_aligner = DashboardGridLayoutWidgetsXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioDenominatorAggregationPerSeriesAlignerEnum.to_proto(
                resource.per_series_aligner
            )
        if DashboardGridLayoutWidgetsXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioDenominatorAggregationCrossSeriesReducerEnum.to_proto(
            resource.cross_series_reducer
        ):
            res.cross_series_reducer = DashboardGridLayoutWidgetsXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioDenominatorAggregationCrossSeriesReducerEnum.to_proto(
                resource.cross_series_reducer
            )
        if Primitive.to_proto(resource.group_by_fields):
            res.group_by_fields.extend(Primitive.to_proto(resource.group_by_fields))
        return res

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return None

        return DashboardGridLayoutWidgetsXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioDenominatorAggregation(
            alignment_period=Primitive.from_proto(resource.alignment_period),
            per_series_aligner=DashboardGridLayoutWidgetsXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioDenominatorAggregationPerSeriesAlignerEnum.from_proto(
                resource.per_series_aligner
            ),
            cross_series_reducer=DashboardGridLayoutWidgetsXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioDenominatorAggregationCrossSeriesReducerEnum.from_proto(
                resource.cross_series_reducer
            ),
            group_by_fields=Primitive.from_proto(resource.group_by_fields),
        )


class DashboardGridLayoutWidgetsXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioDenominatorAggregationArray(
    object
):
    @classmethod
    def to_proto(self, resources):
        if not resources:
            return resources
        return [
            DashboardGridLayoutWidgetsXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioDenominatorAggregation.to_proto(
                i
            )
            for i in resources
        ]

    @classmethod
    def from_proto(self, resources):
        return [
            DashboardGridLayoutWidgetsXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioDenominatorAggregation.from_proto(
                i
            )
            for i in resources
        ]


class DashboardGridLayoutWidgetsXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioSecondaryAggregation(
    object
):
    def __init__(
        self,
        alignment_period: str = None,
        per_series_aligner: str = None,
        cross_series_reducer: str = None,
        group_by_fields: list = None,
    ):
        self.alignment_period = alignment_period
        self.per_series_aligner = per_series_aligner
        self.cross_series_reducer = cross_series_reducer
        self.group_by_fields = group_by_fields

    @classmethod
    def to_proto(self, resource):
        if not resource:
            return None

        res = (
            dashboard_pb2.MonitoringBetaDashboardGridLayoutWidgetsXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioSecondaryAggregation()
        )
        if Primitive.to_proto(resource.alignment_period):
            res.alignment_period = Primitive.to_proto(resource.alignment_period)
        if DashboardGridLayoutWidgetsXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioSecondaryAggregationPerSeriesAlignerEnum.to_proto(
            resource.per_series_aligner
        ):
            res.per_series_aligner = DashboardGridLayoutWidgetsXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioSecondaryAggregationPerSeriesAlignerEnum.to_proto(
                resource.per_series_aligner
            )
        if DashboardGridLayoutWidgetsXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioSecondaryAggregationCrossSeriesReducerEnum.to_proto(
            resource.cross_series_reducer
        ):
            res.cross_series_reducer = DashboardGridLayoutWidgetsXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioSecondaryAggregationCrossSeriesReducerEnum.to_proto(
                resource.cross_series_reducer
            )
        if Primitive.to_proto(resource.group_by_fields):
            res.group_by_fields.extend(Primitive.to_proto(resource.group_by_fields))
        return res

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return None

        return DashboardGridLayoutWidgetsXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioSecondaryAggregation(
            alignment_period=Primitive.from_proto(resource.alignment_period),
            per_series_aligner=DashboardGridLayoutWidgetsXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioSecondaryAggregationPerSeriesAlignerEnum.from_proto(
                resource.per_series_aligner
            ),
            cross_series_reducer=DashboardGridLayoutWidgetsXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioSecondaryAggregationCrossSeriesReducerEnum.from_proto(
                resource.cross_series_reducer
            ),
            group_by_fields=Primitive.from_proto(resource.group_by_fields),
        )


class DashboardGridLayoutWidgetsXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioSecondaryAggregationArray(
    object
):
    @classmethod
    def to_proto(self, resources):
        if not resources:
            return resources
        return [
            DashboardGridLayoutWidgetsXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioSecondaryAggregation.to_proto(
                i
            )
            for i in resources
        ]

    @classmethod
    def from_proto(self, resources):
        return [
            DashboardGridLayoutWidgetsXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioSecondaryAggregation.from_proto(
                i
            )
            for i in resources
        ]


class DashboardGridLayoutWidgetsXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioPickTimeSeriesFilter(
    object
):
    def __init__(
        self,
        ranking_method: str = None,
        num_time_series: int = None,
        direction: str = None,
    ):
        self.ranking_method = ranking_method
        self.num_time_series = num_time_series
        self.direction = direction

    @classmethod
    def to_proto(self, resource):
        if not resource:
            return None

        res = (
            dashboard_pb2.MonitoringBetaDashboardGridLayoutWidgetsXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioPickTimeSeriesFilter()
        )
        if DashboardGridLayoutWidgetsXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioPickTimeSeriesFilterRankingMethodEnum.to_proto(
            resource.ranking_method
        ):
            res.ranking_method = DashboardGridLayoutWidgetsXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioPickTimeSeriesFilterRankingMethodEnum.to_proto(
                resource.ranking_method
            )
        if Primitive.to_proto(resource.num_time_series):
            res.num_time_series = Primitive.to_proto(resource.num_time_series)
        if DashboardGridLayoutWidgetsXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioPickTimeSeriesFilterDirectionEnum.to_proto(
            resource.direction
        ):
            res.direction = DashboardGridLayoutWidgetsXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioPickTimeSeriesFilterDirectionEnum.to_proto(
                resource.direction
            )
        return res

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return None

        return DashboardGridLayoutWidgetsXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioPickTimeSeriesFilter(
            ranking_method=DashboardGridLayoutWidgetsXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioPickTimeSeriesFilterRankingMethodEnum.from_proto(
                resource.ranking_method
            ),
            num_time_series=Primitive.from_proto(resource.num_time_series),
            direction=DashboardGridLayoutWidgetsXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioPickTimeSeriesFilterDirectionEnum.from_proto(
                resource.direction
            ),
        )


class DashboardGridLayoutWidgetsXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioPickTimeSeriesFilterArray(
    object
):
    @classmethod
    def to_proto(self, resources):
        if not resources:
            return resources
        return [
            DashboardGridLayoutWidgetsXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioPickTimeSeriesFilter.to_proto(
                i
            )
            for i in resources
        ]

    @classmethod
    def from_proto(self, resources):
        return [
            DashboardGridLayoutWidgetsXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioPickTimeSeriesFilter.from_proto(
                i
            )
            for i in resources
        ]


class DashboardGridLayoutWidgetsXyChartThresholds(object):
    def __init__(
        self,
        label: str = None,
        value: float = None,
        color: str = None,
        direction: str = None,
    ):
        self.label = label
        self.value = value
        self.color = color
        self.direction = direction

    @classmethod
    def to_proto(self, resource):
        if not resource:
            return None

        res = dashboard_pb2.MonitoringBetaDashboardGridLayoutWidgetsXyChartThresholds()
        if Primitive.to_proto(resource.label):
            res.label = Primitive.to_proto(resource.label)
        if Primitive.to_proto(resource.value):
            res.value = Primitive.to_proto(resource.value)
        if DashboardGridLayoutWidgetsXyChartThresholdsColorEnum.to_proto(
            resource.color
        ):
            res.color = DashboardGridLayoutWidgetsXyChartThresholdsColorEnum.to_proto(
                resource.color
            )
        if DashboardGridLayoutWidgetsXyChartThresholdsDirectionEnum.to_proto(
            resource.direction
        ):
            res.direction = (
                DashboardGridLayoutWidgetsXyChartThresholdsDirectionEnum.to_proto(
                    resource.direction
                )
            )
        return res

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return None

        return DashboardGridLayoutWidgetsXyChartThresholds(
            label=Primitive.from_proto(resource.label),
            value=Primitive.from_proto(resource.value),
            color=DashboardGridLayoutWidgetsXyChartThresholdsColorEnum.from_proto(
                resource.color
            ),
            direction=DashboardGridLayoutWidgetsXyChartThresholdsDirectionEnum.from_proto(
                resource.direction
            ),
        )


class DashboardGridLayoutWidgetsXyChartThresholdsArray(object):
    @classmethod
    def to_proto(self, resources):
        if not resources:
            return resources
        return [
            DashboardGridLayoutWidgetsXyChartThresholds.to_proto(i) for i in resources
        ]

    @classmethod
    def from_proto(self, resources):
        return [
            DashboardGridLayoutWidgetsXyChartThresholds.from_proto(i) for i in resources
        ]


class DashboardGridLayoutWidgetsXyChartXAxis(object):
    def __init__(self, label: str = None, scale: str = None):
        self.label = label
        self.scale = scale

    @classmethod
    def to_proto(self, resource):
        if not resource:
            return None

        res = dashboard_pb2.MonitoringBetaDashboardGridLayoutWidgetsXyChartXAxis()
        if Primitive.to_proto(resource.label):
            res.label = Primitive.to_proto(resource.label)
        if DashboardGridLayoutWidgetsXyChartXAxisScaleEnum.to_proto(resource.scale):
            res.scale = DashboardGridLayoutWidgetsXyChartXAxisScaleEnum.to_proto(
                resource.scale
            )
        return res

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return None

        return DashboardGridLayoutWidgetsXyChartXAxis(
            label=Primitive.from_proto(resource.label),
            scale=DashboardGridLayoutWidgetsXyChartXAxisScaleEnum.from_proto(
                resource.scale
            ),
        )


class DashboardGridLayoutWidgetsXyChartXAxisArray(object):
    @classmethod
    def to_proto(self, resources):
        if not resources:
            return resources
        return [DashboardGridLayoutWidgetsXyChartXAxis.to_proto(i) for i in resources]

    @classmethod
    def from_proto(self, resources):
        return [DashboardGridLayoutWidgetsXyChartXAxis.from_proto(i) for i in resources]


class DashboardGridLayoutWidgetsXyChartYAxis(object):
    def __init__(self, label: str = None, scale: str = None):
        self.label = label
        self.scale = scale

    @classmethod
    def to_proto(self, resource):
        if not resource:
            return None

        res = dashboard_pb2.MonitoringBetaDashboardGridLayoutWidgetsXyChartYAxis()
        if Primitive.to_proto(resource.label):
            res.label = Primitive.to_proto(resource.label)
        if DashboardGridLayoutWidgetsXyChartYAxisScaleEnum.to_proto(resource.scale):
            res.scale = DashboardGridLayoutWidgetsXyChartYAxisScaleEnum.to_proto(
                resource.scale
            )
        return res

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return None

        return DashboardGridLayoutWidgetsXyChartYAxis(
            label=Primitive.from_proto(resource.label),
            scale=DashboardGridLayoutWidgetsXyChartYAxisScaleEnum.from_proto(
                resource.scale
            ),
        )


class DashboardGridLayoutWidgetsXyChartYAxisArray(object):
    @classmethod
    def to_proto(self, resources):
        if not resources:
            return resources
        return [DashboardGridLayoutWidgetsXyChartYAxis.to_proto(i) for i in resources]

    @classmethod
    def from_proto(self, resources):
        return [DashboardGridLayoutWidgetsXyChartYAxis.from_proto(i) for i in resources]


class DashboardGridLayoutWidgetsXyChartChartOptions(object):
    def __init__(self, mode: str = None):
        self.mode = mode

    @classmethod
    def to_proto(self, resource):
        if not resource:
            return None

        res = (
            dashboard_pb2.MonitoringBetaDashboardGridLayoutWidgetsXyChartChartOptions()
        )
        if DashboardGridLayoutWidgetsXyChartChartOptionsModeEnum.to_proto(
            resource.mode
        ):
            res.mode = DashboardGridLayoutWidgetsXyChartChartOptionsModeEnum.to_proto(
                resource.mode
            )
        return res

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return None

        return DashboardGridLayoutWidgetsXyChartChartOptions(
            mode=DashboardGridLayoutWidgetsXyChartChartOptionsModeEnum.from_proto(
                resource.mode
            ),
        )


class DashboardGridLayoutWidgetsXyChartChartOptionsArray(object):
    @classmethod
    def to_proto(self, resources):
        if not resources:
            return resources
        return [
            DashboardGridLayoutWidgetsXyChartChartOptions.to_proto(i) for i in resources
        ]

    @classmethod
    def from_proto(self, resources):
        return [
            DashboardGridLayoutWidgetsXyChartChartOptions.from_proto(i)
            for i in resources
        ]


class DashboardGridLayoutWidgetsScorecard(object):
    def __init__(
        self,
        time_series_query: dict = None,
        gauge_view: dict = None,
        spark_chart_view: dict = None,
        thresholds: list = None,
    ):
        self.time_series_query = time_series_query
        self.gauge_view = gauge_view
        self.spark_chart_view = spark_chart_view
        self.thresholds = thresholds

    @classmethod
    def to_proto(self, resource):
        if not resource:
            return None

        res = dashboard_pb2.MonitoringBetaDashboardGridLayoutWidgetsScorecard()
        if DashboardGridLayoutWidgetsScorecardTimeSeriesQuery.to_proto(
            resource.time_series_query
        ):
            res.time_series_query.CopyFrom(
                DashboardGridLayoutWidgetsScorecardTimeSeriesQuery.to_proto(
                    resource.time_series_query
                )
            )
        else:
            res.ClearField("time_series_query")
        if DashboardGridLayoutWidgetsScorecardGaugeView.to_proto(resource.gauge_view):
            res.gauge_view.CopyFrom(
                DashboardGridLayoutWidgetsScorecardGaugeView.to_proto(
                    resource.gauge_view
                )
            )
        else:
            res.ClearField("gauge_view")
        if DashboardGridLayoutWidgetsScorecardSparkChartView.to_proto(
            resource.spark_chart_view
        ):
            res.spark_chart_view.CopyFrom(
                DashboardGridLayoutWidgetsScorecardSparkChartView.to_proto(
                    resource.spark_chart_view
                )
            )
        else:
            res.ClearField("spark_chart_view")
        if DashboardGridLayoutWidgetsScorecardThresholdsArray.to_proto(
            resource.thresholds
        ):
            res.thresholds.extend(
                DashboardGridLayoutWidgetsScorecardThresholdsArray.to_proto(
                    resource.thresholds
                )
            )
        return res

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return None

        return DashboardGridLayoutWidgetsScorecard(
            time_series_query=DashboardGridLayoutWidgetsScorecardTimeSeriesQuery.from_proto(
                resource.time_series_query
            ),
            gauge_view=DashboardGridLayoutWidgetsScorecardGaugeView.from_proto(
                resource.gauge_view
            ),
            spark_chart_view=DashboardGridLayoutWidgetsScorecardSparkChartView.from_proto(
                resource.spark_chart_view
            ),
            thresholds=DashboardGridLayoutWidgetsScorecardThresholdsArray.from_proto(
                resource.thresholds
            ),
        )


class DashboardGridLayoutWidgetsScorecardArray(object):
    @classmethod
    def to_proto(self, resources):
        if not resources:
            return resources
        return [DashboardGridLayoutWidgetsScorecard.to_proto(i) for i in resources]

    @classmethod
    def from_proto(self, resources):
        return [DashboardGridLayoutWidgetsScorecard.from_proto(i) for i in resources]


class DashboardGridLayoutWidgetsScorecardTimeSeriesQuery(object):
    def __init__(
        self,
        time_series_filter: dict = None,
        time_series_filter_ratio: dict = None,
        time_series_query_language: str = None,
        unit_override: str = None,
    ):
        self.time_series_filter = time_series_filter
        self.time_series_filter_ratio = time_series_filter_ratio
        self.time_series_query_language = time_series_query_language
        self.unit_override = unit_override

    @classmethod
    def to_proto(self, resource):
        if not resource:
            return None

        res = (
            dashboard_pb2.MonitoringBetaDashboardGridLayoutWidgetsScorecardTimeSeriesQuery()
        )
        if DashboardGridLayoutWidgetsScorecardTimeSeriesQueryTimeSeriesFilter.to_proto(
            resource.time_series_filter
        ):
            res.time_series_filter.CopyFrom(
                DashboardGridLayoutWidgetsScorecardTimeSeriesQueryTimeSeriesFilter.to_proto(
                    resource.time_series_filter
                )
            )
        else:
            res.ClearField("time_series_filter")
        if DashboardGridLayoutWidgetsScorecardTimeSeriesQueryTimeSeriesFilterRatio.to_proto(
            resource.time_series_filter_ratio
        ):
            res.time_series_filter_ratio.CopyFrom(
                DashboardGridLayoutWidgetsScorecardTimeSeriesQueryTimeSeriesFilterRatio.to_proto(
                    resource.time_series_filter_ratio
                )
            )
        else:
            res.ClearField("time_series_filter_ratio")
        if Primitive.to_proto(resource.time_series_query_language):
            res.time_series_query_language = Primitive.to_proto(
                resource.time_series_query_language
            )
        if Primitive.to_proto(resource.unit_override):
            res.unit_override = Primitive.to_proto(resource.unit_override)
        return res

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return None

        return DashboardGridLayoutWidgetsScorecardTimeSeriesQuery(
            time_series_filter=DashboardGridLayoutWidgetsScorecardTimeSeriesQueryTimeSeriesFilter.from_proto(
                resource.time_series_filter
            ),
            time_series_filter_ratio=DashboardGridLayoutWidgetsScorecardTimeSeriesQueryTimeSeriesFilterRatio.from_proto(
                resource.time_series_filter_ratio
            ),
            time_series_query_language=Primitive.from_proto(
                resource.time_series_query_language
            ),
            unit_override=Primitive.from_proto(resource.unit_override),
        )


class DashboardGridLayoutWidgetsScorecardTimeSeriesQueryArray(object):
    @classmethod
    def to_proto(self, resources):
        if not resources:
            return resources
        return [
            DashboardGridLayoutWidgetsScorecardTimeSeriesQuery.to_proto(i)
            for i in resources
        ]

    @classmethod
    def from_proto(self, resources):
        return [
            DashboardGridLayoutWidgetsScorecardTimeSeriesQuery.from_proto(i)
            for i in resources
        ]


class DashboardGridLayoutWidgetsScorecardTimeSeriesQueryTimeSeriesFilter(object):
    def __init__(
        self,
        filter: str = None,
        aggregation: dict = None,
        secondary_aggregation: dict = None,
        pick_time_series_filter: dict = None,
    ):
        self.filter = filter
        self.aggregation = aggregation
        self.secondary_aggregation = secondary_aggregation
        self.pick_time_series_filter = pick_time_series_filter

    @classmethod
    def to_proto(self, resource):
        if not resource:
            return None

        res = (
            dashboard_pb2.MonitoringBetaDashboardGridLayoutWidgetsScorecardTimeSeriesQueryTimeSeriesFilter()
        )
        if Primitive.to_proto(resource.filter):
            res.filter = Primitive.to_proto(resource.filter)
        if DashboardGridLayoutWidgetsScorecardTimeSeriesQueryTimeSeriesFilterAggregation.to_proto(
            resource.aggregation
        ):
            res.aggregation.CopyFrom(
                DashboardGridLayoutWidgetsScorecardTimeSeriesQueryTimeSeriesFilterAggregation.to_proto(
                    resource.aggregation
                )
            )
        else:
            res.ClearField("aggregation")
        if DashboardGridLayoutWidgetsScorecardTimeSeriesQueryTimeSeriesFilterSecondaryAggregation.to_proto(
            resource.secondary_aggregation
        ):
            res.secondary_aggregation.CopyFrom(
                DashboardGridLayoutWidgetsScorecardTimeSeriesQueryTimeSeriesFilterSecondaryAggregation.to_proto(
                    resource.secondary_aggregation
                )
            )
        else:
            res.ClearField("secondary_aggregation")
        if DashboardGridLayoutWidgetsScorecardTimeSeriesQueryTimeSeriesFilterPickTimeSeriesFilter.to_proto(
            resource.pick_time_series_filter
        ):
            res.pick_time_series_filter.CopyFrom(
                DashboardGridLayoutWidgetsScorecardTimeSeriesQueryTimeSeriesFilterPickTimeSeriesFilter.to_proto(
                    resource.pick_time_series_filter
                )
            )
        else:
            res.ClearField("pick_time_series_filter")
        return res

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return None

        return DashboardGridLayoutWidgetsScorecardTimeSeriesQueryTimeSeriesFilter(
            filter=Primitive.from_proto(resource.filter),
            aggregation=DashboardGridLayoutWidgetsScorecardTimeSeriesQueryTimeSeriesFilterAggregation.from_proto(
                resource.aggregation
            ),
            secondary_aggregation=DashboardGridLayoutWidgetsScorecardTimeSeriesQueryTimeSeriesFilterSecondaryAggregation.from_proto(
                resource.secondary_aggregation
            ),
            pick_time_series_filter=DashboardGridLayoutWidgetsScorecardTimeSeriesQueryTimeSeriesFilterPickTimeSeriesFilter.from_proto(
                resource.pick_time_series_filter
            ),
        )


class DashboardGridLayoutWidgetsScorecardTimeSeriesQueryTimeSeriesFilterArray(object):
    @classmethod
    def to_proto(self, resources):
        if not resources:
            return resources
        return [
            DashboardGridLayoutWidgetsScorecardTimeSeriesQueryTimeSeriesFilter.to_proto(
                i
            )
            for i in resources
        ]

    @classmethod
    def from_proto(self, resources):
        return [
            DashboardGridLayoutWidgetsScorecardTimeSeriesQueryTimeSeriesFilter.from_proto(
                i
            )
            for i in resources
        ]


class DashboardGridLayoutWidgetsScorecardTimeSeriesQueryTimeSeriesFilterAggregation(
    object
):
    def __init__(
        self,
        alignment_period: str = None,
        per_series_aligner: str = None,
        cross_series_reducer: str = None,
        group_by_fields: list = None,
    ):
        self.alignment_period = alignment_period
        self.per_series_aligner = per_series_aligner
        self.cross_series_reducer = cross_series_reducer
        self.group_by_fields = group_by_fields

    @classmethod
    def to_proto(self, resource):
        if not resource:
            return None

        res = (
            dashboard_pb2.MonitoringBetaDashboardGridLayoutWidgetsScorecardTimeSeriesQueryTimeSeriesFilterAggregation()
        )
        if Primitive.to_proto(resource.alignment_period):
            res.alignment_period = Primitive.to_proto(resource.alignment_period)
        if DashboardGridLayoutWidgetsScorecardTimeSeriesQueryTimeSeriesFilterAggregationPerSeriesAlignerEnum.to_proto(
            resource.per_series_aligner
        ):
            res.per_series_aligner = DashboardGridLayoutWidgetsScorecardTimeSeriesQueryTimeSeriesFilterAggregationPerSeriesAlignerEnum.to_proto(
                resource.per_series_aligner
            )
        if DashboardGridLayoutWidgetsScorecardTimeSeriesQueryTimeSeriesFilterAggregationCrossSeriesReducerEnum.to_proto(
            resource.cross_series_reducer
        ):
            res.cross_series_reducer = DashboardGridLayoutWidgetsScorecardTimeSeriesQueryTimeSeriesFilterAggregationCrossSeriesReducerEnum.to_proto(
                resource.cross_series_reducer
            )
        if Primitive.to_proto(resource.group_by_fields):
            res.group_by_fields.extend(Primitive.to_proto(resource.group_by_fields))
        return res

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return None

        return DashboardGridLayoutWidgetsScorecardTimeSeriesQueryTimeSeriesFilterAggregation(
            alignment_period=Primitive.from_proto(resource.alignment_period),
            per_series_aligner=DashboardGridLayoutWidgetsScorecardTimeSeriesQueryTimeSeriesFilterAggregationPerSeriesAlignerEnum.from_proto(
                resource.per_series_aligner
            ),
            cross_series_reducer=DashboardGridLayoutWidgetsScorecardTimeSeriesQueryTimeSeriesFilterAggregationCrossSeriesReducerEnum.from_proto(
                resource.cross_series_reducer
            ),
            group_by_fields=Primitive.from_proto(resource.group_by_fields),
        )


class DashboardGridLayoutWidgetsScorecardTimeSeriesQueryTimeSeriesFilterAggregationArray(
    object
):
    @classmethod
    def to_proto(self, resources):
        if not resources:
            return resources
        return [
            DashboardGridLayoutWidgetsScorecardTimeSeriesQueryTimeSeriesFilterAggregation.to_proto(
                i
            )
            for i in resources
        ]

    @classmethod
    def from_proto(self, resources):
        return [
            DashboardGridLayoutWidgetsScorecardTimeSeriesQueryTimeSeriesFilterAggregation.from_proto(
                i
            )
            for i in resources
        ]


class DashboardGridLayoutWidgetsScorecardTimeSeriesQueryTimeSeriesFilterSecondaryAggregation(
    object
):
    def __init__(
        self,
        alignment_period: str = None,
        per_series_aligner: str = None,
        cross_series_reducer: str = None,
        group_by_fields: list = None,
    ):
        self.alignment_period = alignment_period
        self.per_series_aligner = per_series_aligner
        self.cross_series_reducer = cross_series_reducer
        self.group_by_fields = group_by_fields

    @classmethod
    def to_proto(self, resource):
        if not resource:
            return None

        res = (
            dashboard_pb2.MonitoringBetaDashboardGridLayoutWidgetsScorecardTimeSeriesQueryTimeSeriesFilterSecondaryAggregation()
        )
        if Primitive.to_proto(resource.alignment_period):
            res.alignment_period = Primitive.to_proto(resource.alignment_period)
        if DashboardGridLayoutWidgetsScorecardTimeSeriesQueryTimeSeriesFilterSecondaryAggregationPerSeriesAlignerEnum.to_proto(
            resource.per_series_aligner
        ):
            res.per_series_aligner = DashboardGridLayoutWidgetsScorecardTimeSeriesQueryTimeSeriesFilterSecondaryAggregationPerSeriesAlignerEnum.to_proto(
                resource.per_series_aligner
            )
        if DashboardGridLayoutWidgetsScorecardTimeSeriesQueryTimeSeriesFilterSecondaryAggregationCrossSeriesReducerEnum.to_proto(
            resource.cross_series_reducer
        ):
            res.cross_series_reducer = DashboardGridLayoutWidgetsScorecardTimeSeriesQueryTimeSeriesFilterSecondaryAggregationCrossSeriesReducerEnum.to_proto(
                resource.cross_series_reducer
            )
        if Primitive.to_proto(resource.group_by_fields):
            res.group_by_fields.extend(Primitive.to_proto(resource.group_by_fields))
        return res

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return None

        return DashboardGridLayoutWidgetsScorecardTimeSeriesQueryTimeSeriesFilterSecondaryAggregation(
            alignment_period=Primitive.from_proto(resource.alignment_period),
            per_series_aligner=DashboardGridLayoutWidgetsScorecardTimeSeriesQueryTimeSeriesFilterSecondaryAggregationPerSeriesAlignerEnum.from_proto(
                resource.per_series_aligner
            ),
            cross_series_reducer=DashboardGridLayoutWidgetsScorecardTimeSeriesQueryTimeSeriesFilterSecondaryAggregationCrossSeriesReducerEnum.from_proto(
                resource.cross_series_reducer
            ),
            group_by_fields=Primitive.from_proto(resource.group_by_fields),
        )


class DashboardGridLayoutWidgetsScorecardTimeSeriesQueryTimeSeriesFilterSecondaryAggregationArray(
    object
):
    @classmethod
    def to_proto(self, resources):
        if not resources:
            return resources
        return [
            DashboardGridLayoutWidgetsScorecardTimeSeriesQueryTimeSeriesFilterSecondaryAggregation.to_proto(
                i
            )
            for i in resources
        ]

    @classmethod
    def from_proto(self, resources):
        return [
            DashboardGridLayoutWidgetsScorecardTimeSeriesQueryTimeSeriesFilterSecondaryAggregation.from_proto(
                i
            )
            for i in resources
        ]


class DashboardGridLayoutWidgetsScorecardTimeSeriesQueryTimeSeriesFilterPickTimeSeriesFilter(
    object
):
    def __init__(
        self,
        ranking_method: str = None,
        num_time_series: int = None,
        direction: str = None,
    ):
        self.ranking_method = ranking_method
        self.num_time_series = num_time_series
        self.direction = direction

    @classmethod
    def to_proto(self, resource):
        if not resource:
            return None

        res = (
            dashboard_pb2.MonitoringBetaDashboardGridLayoutWidgetsScorecardTimeSeriesQueryTimeSeriesFilterPickTimeSeriesFilter()
        )
        if DashboardGridLayoutWidgetsScorecardTimeSeriesQueryTimeSeriesFilterPickTimeSeriesFilterRankingMethodEnum.to_proto(
            resource.ranking_method
        ):
            res.ranking_method = DashboardGridLayoutWidgetsScorecardTimeSeriesQueryTimeSeriesFilterPickTimeSeriesFilterRankingMethodEnum.to_proto(
                resource.ranking_method
            )
        if Primitive.to_proto(resource.num_time_series):
            res.num_time_series = Primitive.to_proto(resource.num_time_series)
        if DashboardGridLayoutWidgetsScorecardTimeSeriesQueryTimeSeriesFilterPickTimeSeriesFilterDirectionEnum.to_proto(
            resource.direction
        ):
            res.direction = DashboardGridLayoutWidgetsScorecardTimeSeriesQueryTimeSeriesFilterPickTimeSeriesFilterDirectionEnum.to_proto(
                resource.direction
            )
        return res

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return None

        return DashboardGridLayoutWidgetsScorecardTimeSeriesQueryTimeSeriesFilterPickTimeSeriesFilter(
            ranking_method=DashboardGridLayoutWidgetsScorecardTimeSeriesQueryTimeSeriesFilterPickTimeSeriesFilterRankingMethodEnum.from_proto(
                resource.ranking_method
            ),
            num_time_series=Primitive.from_proto(resource.num_time_series),
            direction=DashboardGridLayoutWidgetsScorecardTimeSeriesQueryTimeSeriesFilterPickTimeSeriesFilterDirectionEnum.from_proto(
                resource.direction
            ),
        )


class DashboardGridLayoutWidgetsScorecardTimeSeriesQueryTimeSeriesFilterPickTimeSeriesFilterArray(
    object
):
    @classmethod
    def to_proto(self, resources):
        if not resources:
            return resources
        return [
            DashboardGridLayoutWidgetsScorecardTimeSeriesQueryTimeSeriesFilterPickTimeSeriesFilter.to_proto(
                i
            )
            for i in resources
        ]

    @classmethod
    def from_proto(self, resources):
        return [
            DashboardGridLayoutWidgetsScorecardTimeSeriesQueryTimeSeriesFilterPickTimeSeriesFilter.from_proto(
                i
            )
            for i in resources
        ]


class DashboardGridLayoutWidgetsScorecardTimeSeriesQueryTimeSeriesFilterRatio(object):
    def __init__(
        self,
        numerator: dict = None,
        denominator: dict = None,
        secondary_aggregation: dict = None,
        pick_time_series_filter: dict = None,
    ):
        self.numerator = numerator
        self.denominator = denominator
        self.secondary_aggregation = secondary_aggregation
        self.pick_time_series_filter = pick_time_series_filter

    @classmethod
    def to_proto(self, resource):
        if not resource:
            return None

        res = (
            dashboard_pb2.MonitoringBetaDashboardGridLayoutWidgetsScorecardTimeSeriesQueryTimeSeriesFilterRatio()
        )
        if DashboardGridLayoutWidgetsScorecardTimeSeriesQueryTimeSeriesFilterRatioNumerator.to_proto(
            resource.numerator
        ):
            res.numerator.CopyFrom(
                DashboardGridLayoutWidgetsScorecardTimeSeriesQueryTimeSeriesFilterRatioNumerator.to_proto(
                    resource.numerator
                )
            )
        else:
            res.ClearField("numerator")
        if DashboardGridLayoutWidgetsScorecardTimeSeriesQueryTimeSeriesFilterRatioDenominator.to_proto(
            resource.denominator
        ):
            res.denominator.CopyFrom(
                DashboardGridLayoutWidgetsScorecardTimeSeriesQueryTimeSeriesFilterRatioDenominator.to_proto(
                    resource.denominator
                )
            )
        else:
            res.ClearField("denominator")
        if DashboardGridLayoutWidgetsScorecardTimeSeriesQueryTimeSeriesFilterRatioSecondaryAggregation.to_proto(
            resource.secondary_aggregation
        ):
            res.secondary_aggregation.CopyFrom(
                DashboardGridLayoutWidgetsScorecardTimeSeriesQueryTimeSeriesFilterRatioSecondaryAggregation.to_proto(
                    resource.secondary_aggregation
                )
            )
        else:
            res.ClearField("secondary_aggregation")
        if DashboardGridLayoutWidgetsScorecardTimeSeriesQueryTimeSeriesFilterRatioPickTimeSeriesFilter.to_proto(
            resource.pick_time_series_filter
        ):
            res.pick_time_series_filter.CopyFrom(
                DashboardGridLayoutWidgetsScorecardTimeSeriesQueryTimeSeriesFilterRatioPickTimeSeriesFilter.to_proto(
                    resource.pick_time_series_filter
                )
            )
        else:
            res.ClearField("pick_time_series_filter")
        return res

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return None

        return DashboardGridLayoutWidgetsScorecardTimeSeriesQueryTimeSeriesFilterRatio(
            numerator=DashboardGridLayoutWidgetsScorecardTimeSeriesQueryTimeSeriesFilterRatioNumerator.from_proto(
                resource.numerator
            ),
            denominator=DashboardGridLayoutWidgetsScorecardTimeSeriesQueryTimeSeriesFilterRatioDenominator.from_proto(
                resource.denominator
            ),
            secondary_aggregation=DashboardGridLayoutWidgetsScorecardTimeSeriesQueryTimeSeriesFilterRatioSecondaryAggregation.from_proto(
                resource.secondary_aggregation
            ),
            pick_time_series_filter=DashboardGridLayoutWidgetsScorecardTimeSeriesQueryTimeSeriesFilterRatioPickTimeSeriesFilter.from_proto(
                resource.pick_time_series_filter
            ),
        )


class DashboardGridLayoutWidgetsScorecardTimeSeriesQueryTimeSeriesFilterRatioArray(
    object
):
    @classmethod
    def to_proto(self, resources):
        if not resources:
            return resources
        return [
            DashboardGridLayoutWidgetsScorecardTimeSeriesQueryTimeSeriesFilterRatio.to_proto(
                i
            )
            for i in resources
        ]

    @classmethod
    def from_proto(self, resources):
        return [
            DashboardGridLayoutWidgetsScorecardTimeSeriesQueryTimeSeriesFilterRatio.from_proto(
                i
            )
            for i in resources
        ]


class DashboardGridLayoutWidgetsScorecardTimeSeriesQueryTimeSeriesFilterRatioNumerator(
    object
):
    def __init__(self, filter: str = None, aggregation: dict = None):
        self.filter = filter
        self.aggregation = aggregation

    @classmethod
    def to_proto(self, resource):
        if not resource:
            return None

        res = (
            dashboard_pb2.MonitoringBetaDashboardGridLayoutWidgetsScorecardTimeSeriesQueryTimeSeriesFilterRatioNumerator()
        )
        if Primitive.to_proto(resource.filter):
            res.filter = Primitive.to_proto(resource.filter)
        if DashboardGridLayoutWidgetsScorecardTimeSeriesQueryTimeSeriesFilterRatioNumeratorAggregation.to_proto(
            resource.aggregation
        ):
            res.aggregation.CopyFrom(
                DashboardGridLayoutWidgetsScorecardTimeSeriesQueryTimeSeriesFilterRatioNumeratorAggregation.to_proto(
                    resource.aggregation
                )
            )
        else:
            res.ClearField("aggregation")
        return res

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return None

        return DashboardGridLayoutWidgetsScorecardTimeSeriesQueryTimeSeriesFilterRatioNumerator(
            filter=Primitive.from_proto(resource.filter),
            aggregation=DashboardGridLayoutWidgetsScorecardTimeSeriesQueryTimeSeriesFilterRatioNumeratorAggregation.from_proto(
                resource.aggregation
            ),
        )


class DashboardGridLayoutWidgetsScorecardTimeSeriesQueryTimeSeriesFilterRatioNumeratorArray(
    object
):
    @classmethod
    def to_proto(self, resources):
        if not resources:
            return resources
        return [
            DashboardGridLayoutWidgetsScorecardTimeSeriesQueryTimeSeriesFilterRatioNumerator.to_proto(
                i
            )
            for i in resources
        ]

    @classmethod
    def from_proto(self, resources):
        return [
            DashboardGridLayoutWidgetsScorecardTimeSeriesQueryTimeSeriesFilterRatioNumerator.from_proto(
                i
            )
            for i in resources
        ]


class DashboardGridLayoutWidgetsScorecardTimeSeriesQueryTimeSeriesFilterRatioNumeratorAggregation(
    object
):
    def __init__(
        self,
        alignment_period: str = None,
        per_series_aligner: str = None,
        cross_series_reducer: str = None,
        group_by_fields: list = None,
    ):
        self.alignment_period = alignment_period
        self.per_series_aligner = per_series_aligner
        self.cross_series_reducer = cross_series_reducer
        self.group_by_fields = group_by_fields

    @classmethod
    def to_proto(self, resource):
        if not resource:
            return None

        res = (
            dashboard_pb2.MonitoringBetaDashboardGridLayoutWidgetsScorecardTimeSeriesQueryTimeSeriesFilterRatioNumeratorAggregation()
        )
        if Primitive.to_proto(resource.alignment_period):
            res.alignment_period = Primitive.to_proto(resource.alignment_period)
        if DashboardGridLayoutWidgetsScorecardTimeSeriesQueryTimeSeriesFilterRatioNumeratorAggregationPerSeriesAlignerEnum.to_proto(
            resource.per_series_aligner
        ):
            res.per_series_aligner = DashboardGridLayoutWidgetsScorecardTimeSeriesQueryTimeSeriesFilterRatioNumeratorAggregationPerSeriesAlignerEnum.to_proto(
                resource.per_series_aligner
            )
        if DashboardGridLayoutWidgetsScorecardTimeSeriesQueryTimeSeriesFilterRatioNumeratorAggregationCrossSeriesReducerEnum.to_proto(
            resource.cross_series_reducer
        ):
            res.cross_series_reducer = DashboardGridLayoutWidgetsScorecardTimeSeriesQueryTimeSeriesFilterRatioNumeratorAggregationCrossSeriesReducerEnum.to_proto(
                resource.cross_series_reducer
            )
        if Primitive.to_proto(resource.group_by_fields):
            res.group_by_fields.extend(Primitive.to_proto(resource.group_by_fields))
        return res

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return None

        return DashboardGridLayoutWidgetsScorecardTimeSeriesQueryTimeSeriesFilterRatioNumeratorAggregation(
            alignment_period=Primitive.from_proto(resource.alignment_period),
            per_series_aligner=DashboardGridLayoutWidgetsScorecardTimeSeriesQueryTimeSeriesFilterRatioNumeratorAggregationPerSeriesAlignerEnum.from_proto(
                resource.per_series_aligner
            ),
            cross_series_reducer=DashboardGridLayoutWidgetsScorecardTimeSeriesQueryTimeSeriesFilterRatioNumeratorAggregationCrossSeriesReducerEnum.from_proto(
                resource.cross_series_reducer
            ),
            group_by_fields=Primitive.from_proto(resource.group_by_fields),
        )


class DashboardGridLayoutWidgetsScorecardTimeSeriesQueryTimeSeriesFilterRatioNumeratorAggregationArray(
    object
):
    @classmethod
    def to_proto(self, resources):
        if not resources:
            return resources
        return [
            DashboardGridLayoutWidgetsScorecardTimeSeriesQueryTimeSeriesFilterRatioNumeratorAggregation.to_proto(
                i
            )
            for i in resources
        ]

    @classmethod
    def from_proto(self, resources):
        return [
            DashboardGridLayoutWidgetsScorecardTimeSeriesQueryTimeSeriesFilterRatioNumeratorAggregation.from_proto(
                i
            )
            for i in resources
        ]


class DashboardGridLayoutWidgetsScorecardTimeSeriesQueryTimeSeriesFilterRatioDenominator(
    object
):
    def __init__(self, filter: str = None, aggregation: dict = None):
        self.filter = filter
        self.aggregation = aggregation

    @classmethod
    def to_proto(self, resource):
        if not resource:
            return None

        res = (
            dashboard_pb2.MonitoringBetaDashboardGridLayoutWidgetsScorecardTimeSeriesQueryTimeSeriesFilterRatioDenominator()
        )
        if Primitive.to_proto(resource.filter):
            res.filter = Primitive.to_proto(resource.filter)
        if DashboardGridLayoutWidgetsScorecardTimeSeriesQueryTimeSeriesFilterRatioDenominatorAggregation.to_proto(
            resource.aggregation
        ):
            res.aggregation.CopyFrom(
                DashboardGridLayoutWidgetsScorecardTimeSeriesQueryTimeSeriesFilterRatioDenominatorAggregation.to_proto(
                    resource.aggregation
                )
            )
        else:
            res.ClearField("aggregation")
        return res

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return None

        return DashboardGridLayoutWidgetsScorecardTimeSeriesQueryTimeSeriesFilterRatioDenominator(
            filter=Primitive.from_proto(resource.filter),
            aggregation=DashboardGridLayoutWidgetsScorecardTimeSeriesQueryTimeSeriesFilterRatioDenominatorAggregation.from_proto(
                resource.aggregation
            ),
        )


class DashboardGridLayoutWidgetsScorecardTimeSeriesQueryTimeSeriesFilterRatioDenominatorArray(
    object
):
    @classmethod
    def to_proto(self, resources):
        if not resources:
            return resources
        return [
            DashboardGridLayoutWidgetsScorecardTimeSeriesQueryTimeSeriesFilterRatioDenominator.to_proto(
                i
            )
            for i in resources
        ]

    @classmethod
    def from_proto(self, resources):
        return [
            DashboardGridLayoutWidgetsScorecardTimeSeriesQueryTimeSeriesFilterRatioDenominator.from_proto(
                i
            )
            for i in resources
        ]


class DashboardGridLayoutWidgetsScorecardTimeSeriesQueryTimeSeriesFilterRatioDenominatorAggregation(
    object
):
    def __init__(
        self,
        alignment_period: str = None,
        per_series_aligner: str = None,
        cross_series_reducer: str = None,
        group_by_fields: list = None,
    ):
        self.alignment_period = alignment_period
        self.per_series_aligner = per_series_aligner
        self.cross_series_reducer = cross_series_reducer
        self.group_by_fields = group_by_fields

    @classmethod
    def to_proto(self, resource):
        if not resource:
            return None

        res = (
            dashboard_pb2.MonitoringBetaDashboardGridLayoutWidgetsScorecardTimeSeriesQueryTimeSeriesFilterRatioDenominatorAggregation()
        )
        if Primitive.to_proto(resource.alignment_period):
            res.alignment_period = Primitive.to_proto(resource.alignment_period)
        if DashboardGridLayoutWidgetsScorecardTimeSeriesQueryTimeSeriesFilterRatioDenominatorAggregationPerSeriesAlignerEnum.to_proto(
            resource.per_series_aligner
        ):
            res.per_series_aligner = DashboardGridLayoutWidgetsScorecardTimeSeriesQueryTimeSeriesFilterRatioDenominatorAggregationPerSeriesAlignerEnum.to_proto(
                resource.per_series_aligner
            )
        if DashboardGridLayoutWidgetsScorecardTimeSeriesQueryTimeSeriesFilterRatioDenominatorAggregationCrossSeriesReducerEnum.to_proto(
            resource.cross_series_reducer
        ):
            res.cross_series_reducer = DashboardGridLayoutWidgetsScorecardTimeSeriesQueryTimeSeriesFilterRatioDenominatorAggregationCrossSeriesReducerEnum.to_proto(
                resource.cross_series_reducer
            )
        if Primitive.to_proto(resource.group_by_fields):
            res.group_by_fields.extend(Primitive.to_proto(resource.group_by_fields))
        return res

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return None

        return DashboardGridLayoutWidgetsScorecardTimeSeriesQueryTimeSeriesFilterRatioDenominatorAggregation(
            alignment_period=Primitive.from_proto(resource.alignment_period),
            per_series_aligner=DashboardGridLayoutWidgetsScorecardTimeSeriesQueryTimeSeriesFilterRatioDenominatorAggregationPerSeriesAlignerEnum.from_proto(
                resource.per_series_aligner
            ),
            cross_series_reducer=DashboardGridLayoutWidgetsScorecardTimeSeriesQueryTimeSeriesFilterRatioDenominatorAggregationCrossSeriesReducerEnum.from_proto(
                resource.cross_series_reducer
            ),
            group_by_fields=Primitive.from_proto(resource.group_by_fields),
        )


class DashboardGridLayoutWidgetsScorecardTimeSeriesQueryTimeSeriesFilterRatioDenominatorAggregationArray(
    object
):
    @classmethod
    def to_proto(self, resources):
        if not resources:
            return resources
        return [
            DashboardGridLayoutWidgetsScorecardTimeSeriesQueryTimeSeriesFilterRatioDenominatorAggregation.to_proto(
                i
            )
            for i in resources
        ]

    @classmethod
    def from_proto(self, resources):
        return [
            DashboardGridLayoutWidgetsScorecardTimeSeriesQueryTimeSeriesFilterRatioDenominatorAggregation.from_proto(
                i
            )
            for i in resources
        ]


class DashboardGridLayoutWidgetsScorecardTimeSeriesQueryTimeSeriesFilterRatioSecondaryAggregation(
    object
):
    def __init__(
        self,
        alignment_period: str = None,
        per_series_aligner: str = None,
        cross_series_reducer: str = None,
        group_by_fields: list = None,
    ):
        self.alignment_period = alignment_period
        self.per_series_aligner = per_series_aligner
        self.cross_series_reducer = cross_series_reducer
        self.group_by_fields = group_by_fields

    @classmethod
    def to_proto(self, resource):
        if not resource:
            return None

        res = (
            dashboard_pb2.MonitoringBetaDashboardGridLayoutWidgetsScorecardTimeSeriesQueryTimeSeriesFilterRatioSecondaryAggregation()
        )
        if Primitive.to_proto(resource.alignment_period):
            res.alignment_period = Primitive.to_proto(resource.alignment_period)
        if DashboardGridLayoutWidgetsScorecardTimeSeriesQueryTimeSeriesFilterRatioSecondaryAggregationPerSeriesAlignerEnum.to_proto(
            resource.per_series_aligner
        ):
            res.per_series_aligner = DashboardGridLayoutWidgetsScorecardTimeSeriesQueryTimeSeriesFilterRatioSecondaryAggregationPerSeriesAlignerEnum.to_proto(
                resource.per_series_aligner
            )
        if DashboardGridLayoutWidgetsScorecardTimeSeriesQueryTimeSeriesFilterRatioSecondaryAggregationCrossSeriesReducerEnum.to_proto(
            resource.cross_series_reducer
        ):
            res.cross_series_reducer = DashboardGridLayoutWidgetsScorecardTimeSeriesQueryTimeSeriesFilterRatioSecondaryAggregationCrossSeriesReducerEnum.to_proto(
                resource.cross_series_reducer
            )
        if Primitive.to_proto(resource.group_by_fields):
            res.group_by_fields.extend(Primitive.to_proto(resource.group_by_fields))
        return res

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return None

        return DashboardGridLayoutWidgetsScorecardTimeSeriesQueryTimeSeriesFilterRatioSecondaryAggregation(
            alignment_period=Primitive.from_proto(resource.alignment_period),
            per_series_aligner=DashboardGridLayoutWidgetsScorecardTimeSeriesQueryTimeSeriesFilterRatioSecondaryAggregationPerSeriesAlignerEnum.from_proto(
                resource.per_series_aligner
            ),
            cross_series_reducer=DashboardGridLayoutWidgetsScorecardTimeSeriesQueryTimeSeriesFilterRatioSecondaryAggregationCrossSeriesReducerEnum.from_proto(
                resource.cross_series_reducer
            ),
            group_by_fields=Primitive.from_proto(resource.group_by_fields),
        )


class DashboardGridLayoutWidgetsScorecardTimeSeriesQueryTimeSeriesFilterRatioSecondaryAggregationArray(
    object
):
    @classmethod
    def to_proto(self, resources):
        if not resources:
            return resources
        return [
            DashboardGridLayoutWidgetsScorecardTimeSeriesQueryTimeSeriesFilterRatioSecondaryAggregation.to_proto(
                i
            )
            for i in resources
        ]

    @classmethod
    def from_proto(self, resources):
        return [
            DashboardGridLayoutWidgetsScorecardTimeSeriesQueryTimeSeriesFilterRatioSecondaryAggregation.from_proto(
                i
            )
            for i in resources
        ]


class DashboardGridLayoutWidgetsScorecardTimeSeriesQueryTimeSeriesFilterRatioPickTimeSeriesFilter(
    object
):
    def __init__(
        self,
        ranking_method: str = None,
        num_time_series: int = None,
        direction: str = None,
    ):
        self.ranking_method = ranking_method
        self.num_time_series = num_time_series
        self.direction = direction

    @classmethod
    def to_proto(self, resource):
        if not resource:
            return None

        res = (
            dashboard_pb2.MonitoringBetaDashboardGridLayoutWidgetsScorecardTimeSeriesQueryTimeSeriesFilterRatioPickTimeSeriesFilter()
        )
        if DashboardGridLayoutWidgetsScorecardTimeSeriesQueryTimeSeriesFilterRatioPickTimeSeriesFilterRankingMethodEnum.to_proto(
            resource.ranking_method
        ):
            res.ranking_method = DashboardGridLayoutWidgetsScorecardTimeSeriesQueryTimeSeriesFilterRatioPickTimeSeriesFilterRankingMethodEnum.to_proto(
                resource.ranking_method
            )
        if Primitive.to_proto(resource.num_time_series):
            res.num_time_series = Primitive.to_proto(resource.num_time_series)
        if DashboardGridLayoutWidgetsScorecardTimeSeriesQueryTimeSeriesFilterRatioPickTimeSeriesFilterDirectionEnum.to_proto(
            resource.direction
        ):
            res.direction = DashboardGridLayoutWidgetsScorecardTimeSeriesQueryTimeSeriesFilterRatioPickTimeSeriesFilterDirectionEnum.to_proto(
                resource.direction
            )
        return res

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return None

        return DashboardGridLayoutWidgetsScorecardTimeSeriesQueryTimeSeriesFilterRatioPickTimeSeriesFilter(
            ranking_method=DashboardGridLayoutWidgetsScorecardTimeSeriesQueryTimeSeriesFilterRatioPickTimeSeriesFilterRankingMethodEnum.from_proto(
                resource.ranking_method
            ),
            num_time_series=Primitive.from_proto(resource.num_time_series),
            direction=DashboardGridLayoutWidgetsScorecardTimeSeriesQueryTimeSeriesFilterRatioPickTimeSeriesFilterDirectionEnum.from_proto(
                resource.direction
            ),
        )


class DashboardGridLayoutWidgetsScorecardTimeSeriesQueryTimeSeriesFilterRatioPickTimeSeriesFilterArray(
    object
):
    @classmethod
    def to_proto(self, resources):
        if not resources:
            return resources
        return [
            DashboardGridLayoutWidgetsScorecardTimeSeriesQueryTimeSeriesFilterRatioPickTimeSeriesFilter.to_proto(
                i
            )
            for i in resources
        ]

    @classmethod
    def from_proto(self, resources):
        return [
            DashboardGridLayoutWidgetsScorecardTimeSeriesQueryTimeSeriesFilterRatioPickTimeSeriesFilter.from_proto(
                i
            )
            for i in resources
        ]


class DashboardGridLayoutWidgetsScorecardGaugeView(object):
    def __init__(self, lower_bound: float = None, upper_bound: float = None):
        self.lower_bound = lower_bound
        self.upper_bound = upper_bound

    @classmethod
    def to_proto(self, resource):
        if not resource:
            return None

        res = dashboard_pb2.MonitoringBetaDashboardGridLayoutWidgetsScorecardGaugeView()
        if Primitive.to_proto(resource.lower_bound):
            res.lower_bound = Primitive.to_proto(resource.lower_bound)
        if Primitive.to_proto(resource.upper_bound):
            res.upper_bound = Primitive.to_proto(resource.upper_bound)
        return res

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return None

        return DashboardGridLayoutWidgetsScorecardGaugeView(
            lower_bound=Primitive.from_proto(resource.lower_bound),
            upper_bound=Primitive.from_proto(resource.upper_bound),
        )


class DashboardGridLayoutWidgetsScorecardGaugeViewArray(object):
    @classmethod
    def to_proto(self, resources):
        if not resources:
            return resources
        return [
            DashboardGridLayoutWidgetsScorecardGaugeView.to_proto(i) for i in resources
        ]

    @classmethod
    def from_proto(self, resources):
        return [
            DashboardGridLayoutWidgetsScorecardGaugeView.from_proto(i)
            for i in resources
        ]


class DashboardGridLayoutWidgetsScorecardSparkChartView(object):
    def __init__(self, spark_chart_type: str = None, min_alignment_period: str = None):
        self.spark_chart_type = spark_chart_type
        self.min_alignment_period = min_alignment_period

    @classmethod
    def to_proto(self, resource):
        if not resource:
            return None

        res = (
            dashboard_pb2.MonitoringBetaDashboardGridLayoutWidgetsScorecardSparkChartView()
        )
        if DashboardGridLayoutWidgetsScorecardSparkChartViewSparkChartTypeEnum.to_proto(
            resource.spark_chart_type
        ):
            res.spark_chart_type = DashboardGridLayoutWidgetsScorecardSparkChartViewSparkChartTypeEnum.to_proto(
                resource.spark_chart_type
            )
        if Primitive.to_proto(resource.min_alignment_period):
            res.min_alignment_period = Primitive.to_proto(resource.min_alignment_period)
        return res

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return None

        return DashboardGridLayoutWidgetsScorecardSparkChartView(
            spark_chart_type=DashboardGridLayoutWidgetsScorecardSparkChartViewSparkChartTypeEnum.from_proto(
                resource.spark_chart_type
            ),
            min_alignment_period=Primitive.from_proto(resource.min_alignment_period),
        )


class DashboardGridLayoutWidgetsScorecardSparkChartViewArray(object):
    @classmethod
    def to_proto(self, resources):
        if not resources:
            return resources
        return [
            DashboardGridLayoutWidgetsScorecardSparkChartView.to_proto(i)
            for i in resources
        ]

    @classmethod
    def from_proto(self, resources):
        return [
            DashboardGridLayoutWidgetsScorecardSparkChartView.from_proto(i)
            for i in resources
        ]


class DashboardGridLayoutWidgetsScorecardThresholds(object):
    def __init__(
        self,
        label: str = None,
        value: float = None,
        color: str = None,
        direction: str = None,
    ):
        self.label = label
        self.value = value
        self.color = color
        self.direction = direction

    @classmethod
    def to_proto(self, resource):
        if not resource:
            return None

        res = (
            dashboard_pb2.MonitoringBetaDashboardGridLayoutWidgetsScorecardThresholds()
        )
        if Primitive.to_proto(resource.label):
            res.label = Primitive.to_proto(resource.label)
        if Primitive.to_proto(resource.value):
            res.value = Primitive.to_proto(resource.value)
        if DashboardGridLayoutWidgetsScorecardThresholdsColorEnum.to_proto(
            resource.color
        ):
            res.color = DashboardGridLayoutWidgetsScorecardThresholdsColorEnum.to_proto(
                resource.color
            )
        if DashboardGridLayoutWidgetsScorecardThresholdsDirectionEnum.to_proto(
            resource.direction
        ):
            res.direction = (
                DashboardGridLayoutWidgetsScorecardThresholdsDirectionEnum.to_proto(
                    resource.direction
                )
            )
        return res

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return None

        return DashboardGridLayoutWidgetsScorecardThresholds(
            label=Primitive.from_proto(resource.label),
            value=Primitive.from_proto(resource.value),
            color=DashboardGridLayoutWidgetsScorecardThresholdsColorEnum.from_proto(
                resource.color
            ),
            direction=DashboardGridLayoutWidgetsScorecardThresholdsDirectionEnum.from_proto(
                resource.direction
            ),
        )


class DashboardGridLayoutWidgetsScorecardThresholdsArray(object):
    @classmethod
    def to_proto(self, resources):
        if not resources:
            return resources
        return [
            DashboardGridLayoutWidgetsScorecardThresholds.to_proto(i) for i in resources
        ]

    @classmethod
    def from_proto(self, resources):
        return [
            DashboardGridLayoutWidgetsScorecardThresholds.from_proto(i)
            for i in resources
        ]


class DashboardGridLayoutWidgetsText(object):
    def __init__(self, content: str = None, format: str = None):
        self.content = content
        self.format = format

    @classmethod
    def to_proto(self, resource):
        if not resource:
            return None

        res = dashboard_pb2.MonitoringBetaDashboardGridLayoutWidgetsText()
        if Primitive.to_proto(resource.content):
            res.content = Primitive.to_proto(resource.content)
        if DashboardGridLayoutWidgetsTextFormatEnum.to_proto(resource.format):
            res.format = DashboardGridLayoutWidgetsTextFormatEnum.to_proto(
                resource.format
            )
        return res

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return None

        return DashboardGridLayoutWidgetsText(
            content=Primitive.from_proto(resource.content),
            format=DashboardGridLayoutWidgetsTextFormatEnum.from_proto(resource.format),
        )


class DashboardGridLayoutWidgetsTextArray(object):
    @classmethod
    def to_proto(self, resources):
        if not resources:
            return resources
        return [DashboardGridLayoutWidgetsText.to_proto(i) for i in resources]

    @classmethod
    def from_proto(self, resources):
        return [DashboardGridLayoutWidgetsText.from_proto(i) for i in resources]


class DashboardGridLayoutWidgetsBlank(object):
    @classmethod
    def to_proto(self, resource):
        if not resource:
            return None

        res = dashboard_pb2.MonitoringBetaDashboardGridLayoutWidgetsBlank()
        return res

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return None

        return DashboardGridLayoutWidgetsBlank()


class DashboardGridLayoutWidgetsBlankArray(object):
    @classmethod
    def to_proto(self, resources):
        if not resources:
            return resources
        return [DashboardGridLayoutWidgetsBlank.to_proto(i) for i in resources]

    @classmethod
    def from_proto(self, resources):
        return [DashboardGridLayoutWidgetsBlank.from_proto(i) for i in resources]


class DashboardGridLayoutWidgetsLogsPanel(object):
    def __init__(self, filter: str = None, resource_names: list = None):
        self.filter = filter
        self.resource_names = resource_names

    @classmethod
    def to_proto(self, resource):
        if not resource:
            return None

        res = dashboard_pb2.MonitoringBetaDashboardGridLayoutWidgetsLogsPanel()
        if Primitive.to_proto(resource.filter):
            res.filter = Primitive.to_proto(resource.filter)
        if Primitive.to_proto(resource.resource_names):
            res.resource_names.extend(Primitive.to_proto(resource.resource_names))
        return res

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return None

        return DashboardGridLayoutWidgetsLogsPanel(
            filter=Primitive.from_proto(resource.filter),
            resource_names=Primitive.from_proto(resource.resource_names),
        )


class DashboardGridLayoutWidgetsLogsPanelArray(object):
    @classmethod
    def to_proto(self, resources):
        if not resources:
            return resources
        return [DashboardGridLayoutWidgetsLogsPanel.to_proto(i) for i in resources]

    @classmethod
    def from_proto(self, resources):
        return [DashboardGridLayoutWidgetsLogsPanel.from_proto(i) for i in resources]


class DashboardMosaicLayout(object):
    def __init__(self, columns: int = None, tiles: list = None):
        self.columns = columns
        self.tiles = tiles

    @classmethod
    def to_proto(self, resource):
        if not resource:
            return None

        res = dashboard_pb2.MonitoringBetaDashboardMosaicLayout()
        if Primitive.to_proto(resource.columns):
            res.columns = Primitive.to_proto(resource.columns)
        if DashboardMosaicLayoutTilesArray.to_proto(resource.tiles):
            res.tiles.extend(DashboardMosaicLayoutTilesArray.to_proto(resource.tiles))
        return res

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return None

        return DashboardMosaicLayout(
            columns=Primitive.from_proto(resource.columns),
            tiles=DashboardMosaicLayoutTilesArray.from_proto(resource.tiles),
        )


class DashboardMosaicLayoutArray(object):
    @classmethod
    def to_proto(self, resources):
        if not resources:
            return resources
        return [DashboardMosaicLayout.to_proto(i) for i in resources]

    @classmethod
    def from_proto(self, resources):
        return [DashboardMosaicLayout.from_proto(i) for i in resources]


class DashboardMosaicLayoutTiles(object):
    def __init__(
        self,
        x_pos: int = None,
        y_pos: int = None,
        width: int = None,
        height: int = None,
        widget: dict = None,
    ):
        self.x_pos = x_pos
        self.y_pos = y_pos
        self.width = width
        self.height = height
        self.widget = widget

    @classmethod
    def to_proto(self, resource):
        if not resource:
            return None

        res = dashboard_pb2.MonitoringBetaDashboardMosaicLayoutTiles()
        if Primitive.to_proto(resource.x_pos):
            res.x_pos = Primitive.to_proto(resource.x_pos)
        if Primitive.to_proto(resource.y_pos):
            res.y_pos = Primitive.to_proto(resource.y_pos)
        if Primitive.to_proto(resource.width):
            res.width = Primitive.to_proto(resource.width)
        if Primitive.to_proto(resource.height):
            res.height = Primitive.to_proto(resource.height)
        if DashboardMosaicLayoutTilesWidget.to_proto(resource.widget):
            res.widget.CopyFrom(
                DashboardMosaicLayoutTilesWidget.to_proto(resource.widget)
            )
        else:
            res.ClearField("widget")
        return res

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return None

        return DashboardMosaicLayoutTiles(
            x_pos=Primitive.from_proto(resource.x_pos),
            y_pos=Primitive.from_proto(resource.y_pos),
            width=Primitive.from_proto(resource.width),
            height=Primitive.from_proto(resource.height),
            widget=DashboardMosaicLayoutTilesWidget.from_proto(resource.widget),
        )


class DashboardMosaicLayoutTilesArray(object):
    @classmethod
    def to_proto(self, resources):
        if not resources:
            return resources
        return [DashboardMosaicLayoutTiles.to_proto(i) for i in resources]

    @classmethod
    def from_proto(self, resources):
        return [DashboardMosaicLayoutTiles.from_proto(i) for i in resources]


class DashboardMosaicLayoutTilesWidget(object):
    def __init__(
        self,
        title: str = None,
        xy_chart: dict = None,
        scorecard: dict = None,
        text: dict = None,
        blank: dict = None,
        logs_panel: dict = None,
    ):
        self.title = title
        self.xy_chart = xy_chart
        self.scorecard = scorecard
        self.text = text
        self.blank = blank
        self.logs_panel = logs_panel

    @classmethod
    def to_proto(self, resource):
        if not resource:
            return None

        res = dashboard_pb2.MonitoringBetaDashboardMosaicLayoutTilesWidget()
        if Primitive.to_proto(resource.title):
            res.title = Primitive.to_proto(resource.title)
        if DashboardMosaicLayoutTilesWidgetXyChart.to_proto(resource.xy_chart):
            res.xy_chart.CopyFrom(
                DashboardMosaicLayoutTilesWidgetXyChart.to_proto(resource.xy_chart)
            )
        else:
            res.ClearField("xy_chart")
        if DashboardMosaicLayoutTilesWidgetScorecard.to_proto(resource.scorecard):
            res.scorecard.CopyFrom(
                DashboardMosaicLayoutTilesWidgetScorecard.to_proto(resource.scorecard)
            )
        else:
            res.ClearField("scorecard")
        if DashboardMosaicLayoutTilesWidgetText.to_proto(resource.text):
            res.text.CopyFrom(
                DashboardMosaicLayoutTilesWidgetText.to_proto(resource.text)
            )
        else:
            res.ClearField("text")
        if DashboardMosaicLayoutTilesWidgetBlank.to_proto(resource.blank):
            res.blank.CopyFrom(
                DashboardMosaicLayoutTilesWidgetBlank.to_proto(resource.blank)
            )
        else:
            res.ClearField("blank")
        if DashboardMosaicLayoutTilesWidgetLogsPanel.to_proto(resource.logs_panel):
            res.logs_panel.CopyFrom(
                DashboardMosaicLayoutTilesWidgetLogsPanel.to_proto(resource.logs_panel)
            )
        else:
            res.ClearField("logs_panel")
        return res

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return None

        return DashboardMosaicLayoutTilesWidget(
            title=Primitive.from_proto(resource.title),
            xy_chart=DashboardMosaicLayoutTilesWidgetXyChart.from_proto(
                resource.xy_chart
            ),
            scorecard=DashboardMosaicLayoutTilesWidgetScorecard.from_proto(
                resource.scorecard
            ),
            text=DashboardMosaicLayoutTilesWidgetText.from_proto(resource.text),
            blank=DashboardMosaicLayoutTilesWidgetBlank.from_proto(resource.blank),
            logs_panel=DashboardMosaicLayoutTilesWidgetLogsPanel.from_proto(
                resource.logs_panel
            ),
        )


class DashboardMosaicLayoutTilesWidgetArray(object):
    @classmethod
    def to_proto(self, resources):
        if not resources:
            return resources
        return [DashboardMosaicLayoutTilesWidget.to_proto(i) for i in resources]

    @classmethod
    def from_proto(self, resources):
        return [DashboardMosaicLayoutTilesWidget.from_proto(i) for i in resources]


class DashboardMosaicLayoutTilesWidgetXyChart(object):
    def __init__(
        self,
        data_sets: list = None,
        timeshift_duration: str = None,
        thresholds: list = None,
        x_axis: dict = None,
        y_axis: dict = None,
        chart_options: dict = None,
    ):
        self.data_sets = data_sets
        self.timeshift_duration = timeshift_duration
        self.thresholds = thresholds
        self.x_axis = x_axis
        self.y_axis = y_axis
        self.chart_options = chart_options

    @classmethod
    def to_proto(self, resource):
        if not resource:
            return None

        res = dashboard_pb2.MonitoringBetaDashboardMosaicLayoutTilesWidgetXyChart()
        if DashboardMosaicLayoutTilesWidgetXyChartDataSetsArray.to_proto(
            resource.data_sets
        ):
            res.data_sets.extend(
                DashboardMosaicLayoutTilesWidgetXyChartDataSetsArray.to_proto(
                    resource.data_sets
                )
            )
        if Primitive.to_proto(resource.timeshift_duration):
            res.timeshift_duration = Primitive.to_proto(resource.timeshift_duration)
        if DashboardMosaicLayoutTilesWidgetXyChartThresholdsArray.to_proto(
            resource.thresholds
        ):
            res.thresholds.extend(
                DashboardMosaicLayoutTilesWidgetXyChartThresholdsArray.to_proto(
                    resource.thresholds
                )
            )
        if DashboardMosaicLayoutTilesWidgetXyChartXAxis.to_proto(resource.x_axis):
            res.x_axis.CopyFrom(
                DashboardMosaicLayoutTilesWidgetXyChartXAxis.to_proto(resource.x_axis)
            )
        else:
            res.ClearField("x_axis")
        if DashboardMosaicLayoutTilesWidgetXyChartYAxis.to_proto(resource.y_axis):
            res.y_axis.CopyFrom(
                DashboardMosaicLayoutTilesWidgetXyChartYAxis.to_proto(resource.y_axis)
            )
        else:
            res.ClearField("y_axis")
        if DashboardMosaicLayoutTilesWidgetXyChartChartOptions.to_proto(
            resource.chart_options
        ):
            res.chart_options.CopyFrom(
                DashboardMosaicLayoutTilesWidgetXyChartChartOptions.to_proto(
                    resource.chart_options
                )
            )
        else:
            res.ClearField("chart_options")
        return res

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return None

        return DashboardMosaicLayoutTilesWidgetXyChart(
            data_sets=DashboardMosaicLayoutTilesWidgetXyChartDataSetsArray.from_proto(
                resource.data_sets
            ),
            timeshift_duration=Primitive.from_proto(resource.timeshift_duration),
            thresholds=DashboardMosaicLayoutTilesWidgetXyChartThresholdsArray.from_proto(
                resource.thresholds
            ),
            x_axis=DashboardMosaicLayoutTilesWidgetXyChartXAxis.from_proto(
                resource.x_axis
            ),
            y_axis=DashboardMosaicLayoutTilesWidgetXyChartYAxis.from_proto(
                resource.y_axis
            ),
            chart_options=DashboardMosaicLayoutTilesWidgetXyChartChartOptions.from_proto(
                resource.chart_options
            ),
        )


class DashboardMosaicLayoutTilesWidgetXyChartArray(object):
    @classmethod
    def to_proto(self, resources):
        if not resources:
            return resources
        return [DashboardMosaicLayoutTilesWidgetXyChart.to_proto(i) for i in resources]

    @classmethod
    def from_proto(self, resources):
        return [
            DashboardMosaicLayoutTilesWidgetXyChart.from_proto(i) for i in resources
        ]


class DashboardMosaicLayoutTilesWidgetXyChartDataSets(object):
    def __init__(
        self,
        time_series_query: dict = None,
        plot_type: str = None,
        legend_template: str = None,
        min_alignment_period: str = None,
    ):
        self.time_series_query = time_series_query
        self.plot_type = plot_type
        self.legend_template = legend_template
        self.min_alignment_period = min_alignment_period

    @classmethod
    def to_proto(self, resource):
        if not resource:
            return None

        res = (
            dashboard_pb2.MonitoringBetaDashboardMosaicLayoutTilesWidgetXyChartDataSets()
        )
        if DashboardMosaicLayoutTilesWidgetXyChartDataSetsTimeSeriesQuery.to_proto(
            resource.time_series_query
        ):
            res.time_series_query.CopyFrom(
                DashboardMosaicLayoutTilesWidgetXyChartDataSetsTimeSeriesQuery.to_proto(
                    resource.time_series_query
                )
            )
        else:
            res.ClearField("time_series_query")
        if DashboardMosaicLayoutTilesWidgetXyChartDataSetsPlotTypeEnum.to_proto(
            resource.plot_type
        ):
            res.plot_type = (
                DashboardMosaicLayoutTilesWidgetXyChartDataSetsPlotTypeEnum.to_proto(
                    resource.plot_type
                )
            )
        if Primitive.to_proto(resource.legend_template):
            res.legend_template = Primitive.to_proto(resource.legend_template)
        if Primitive.to_proto(resource.min_alignment_period):
            res.min_alignment_period = Primitive.to_proto(resource.min_alignment_period)
        return res

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return None

        return DashboardMosaicLayoutTilesWidgetXyChartDataSets(
            time_series_query=DashboardMosaicLayoutTilesWidgetXyChartDataSetsTimeSeriesQuery.from_proto(
                resource.time_series_query
            ),
            plot_type=DashboardMosaicLayoutTilesWidgetXyChartDataSetsPlotTypeEnum.from_proto(
                resource.plot_type
            ),
            legend_template=Primitive.from_proto(resource.legend_template),
            min_alignment_period=Primitive.from_proto(resource.min_alignment_period),
        )


class DashboardMosaicLayoutTilesWidgetXyChartDataSetsArray(object):
    @classmethod
    def to_proto(self, resources):
        if not resources:
            return resources
        return [
            DashboardMosaicLayoutTilesWidgetXyChartDataSets.to_proto(i)
            for i in resources
        ]

    @classmethod
    def from_proto(self, resources):
        return [
            DashboardMosaicLayoutTilesWidgetXyChartDataSets.from_proto(i)
            for i in resources
        ]


class DashboardMosaicLayoutTilesWidgetXyChartDataSetsTimeSeriesQuery(object):
    def __init__(
        self,
        time_series_filter: dict = None,
        time_series_filter_ratio: dict = None,
        time_series_query_language: str = None,
        unit_override: str = None,
    ):
        self.time_series_filter = time_series_filter
        self.time_series_filter_ratio = time_series_filter_ratio
        self.time_series_query_language = time_series_query_language
        self.unit_override = unit_override

    @classmethod
    def to_proto(self, resource):
        if not resource:
            return None

        res = (
            dashboard_pb2.MonitoringBetaDashboardMosaicLayoutTilesWidgetXyChartDataSetsTimeSeriesQuery()
        )
        if DashboardMosaicLayoutTilesWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilter.to_proto(
            resource.time_series_filter
        ):
            res.time_series_filter.CopyFrom(
                DashboardMosaicLayoutTilesWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilter.to_proto(
                    resource.time_series_filter
                )
            )
        else:
            res.ClearField("time_series_filter")
        if DashboardMosaicLayoutTilesWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatio.to_proto(
            resource.time_series_filter_ratio
        ):
            res.time_series_filter_ratio.CopyFrom(
                DashboardMosaicLayoutTilesWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatio.to_proto(
                    resource.time_series_filter_ratio
                )
            )
        else:
            res.ClearField("time_series_filter_ratio")
        if Primitive.to_proto(resource.time_series_query_language):
            res.time_series_query_language = Primitive.to_proto(
                resource.time_series_query_language
            )
        if Primitive.to_proto(resource.unit_override):
            res.unit_override = Primitive.to_proto(resource.unit_override)
        return res

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return None

        return DashboardMosaicLayoutTilesWidgetXyChartDataSetsTimeSeriesQuery(
            time_series_filter=DashboardMosaicLayoutTilesWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilter.from_proto(
                resource.time_series_filter
            ),
            time_series_filter_ratio=DashboardMosaicLayoutTilesWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatio.from_proto(
                resource.time_series_filter_ratio
            ),
            time_series_query_language=Primitive.from_proto(
                resource.time_series_query_language
            ),
            unit_override=Primitive.from_proto(resource.unit_override),
        )


class DashboardMosaicLayoutTilesWidgetXyChartDataSetsTimeSeriesQueryArray(object):
    @classmethod
    def to_proto(self, resources):
        if not resources:
            return resources
        return [
            DashboardMosaicLayoutTilesWidgetXyChartDataSetsTimeSeriesQuery.to_proto(i)
            for i in resources
        ]

    @classmethod
    def from_proto(self, resources):
        return [
            DashboardMosaicLayoutTilesWidgetXyChartDataSetsTimeSeriesQuery.from_proto(i)
            for i in resources
        ]


class DashboardMosaicLayoutTilesWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilter(
    object
):
    def __init__(
        self,
        filter: str = None,
        aggregation: dict = None,
        secondary_aggregation: dict = None,
        pick_time_series_filter: dict = None,
    ):
        self.filter = filter
        self.aggregation = aggregation
        self.secondary_aggregation = secondary_aggregation
        self.pick_time_series_filter = pick_time_series_filter

    @classmethod
    def to_proto(self, resource):
        if not resource:
            return None

        res = (
            dashboard_pb2.MonitoringBetaDashboardMosaicLayoutTilesWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilter()
        )
        if Primitive.to_proto(resource.filter):
            res.filter = Primitive.to_proto(resource.filter)
        if DashboardMosaicLayoutTilesWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterAggregation.to_proto(
            resource.aggregation
        ):
            res.aggregation.CopyFrom(
                DashboardMosaicLayoutTilesWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterAggregation.to_proto(
                    resource.aggregation
                )
            )
        else:
            res.ClearField("aggregation")
        if DashboardMosaicLayoutTilesWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterSecondaryAggregation.to_proto(
            resource.secondary_aggregation
        ):
            res.secondary_aggregation.CopyFrom(
                DashboardMosaicLayoutTilesWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterSecondaryAggregation.to_proto(
                    resource.secondary_aggregation
                )
            )
        else:
            res.ClearField("secondary_aggregation")
        if DashboardMosaicLayoutTilesWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterPickTimeSeriesFilter.to_proto(
            resource.pick_time_series_filter
        ):
            res.pick_time_series_filter.CopyFrom(
                DashboardMosaicLayoutTilesWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterPickTimeSeriesFilter.to_proto(
                    resource.pick_time_series_filter
                )
            )
        else:
            res.ClearField("pick_time_series_filter")
        return res

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return None

        return DashboardMosaicLayoutTilesWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilter(
            filter=Primitive.from_proto(resource.filter),
            aggregation=DashboardMosaicLayoutTilesWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterAggregation.from_proto(
                resource.aggregation
            ),
            secondary_aggregation=DashboardMosaicLayoutTilesWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterSecondaryAggregation.from_proto(
                resource.secondary_aggregation
            ),
            pick_time_series_filter=DashboardMosaicLayoutTilesWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterPickTimeSeriesFilter.from_proto(
                resource.pick_time_series_filter
            ),
        )


class DashboardMosaicLayoutTilesWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterArray(
    object
):
    @classmethod
    def to_proto(self, resources):
        if not resources:
            return resources
        return [
            DashboardMosaicLayoutTilesWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilter.to_proto(
                i
            )
            for i in resources
        ]

    @classmethod
    def from_proto(self, resources):
        return [
            DashboardMosaicLayoutTilesWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilter.from_proto(
                i
            )
            for i in resources
        ]


class DashboardMosaicLayoutTilesWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterAggregation(
    object
):
    def __init__(
        self,
        alignment_period: str = None,
        per_series_aligner: str = None,
        cross_series_reducer: str = None,
        group_by_fields: list = None,
    ):
        self.alignment_period = alignment_period
        self.per_series_aligner = per_series_aligner
        self.cross_series_reducer = cross_series_reducer
        self.group_by_fields = group_by_fields

    @classmethod
    def to_proto(self, resource):
        if not resource:
            return None

        res = (
            dashboard_pb2.MonitoringBetaDashboardMosaicLayoutTilesWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterAggregation()
        )
        if Primitive.to_proto(resource.alignment_period):
            res.alignment_period = Primitive.to_proto(resource.alignment_period)
        if DashboardMosaicLayoutTilesWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterAggregationPerSeriesAlignerEnum.to_proto(
            resource.per_series_aligner
        ):
            res.per_series_aligner = DashboardMosaicLayoutTilesWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterAggregationPerSeriesAlignerEnum.to_proto(
                resource.per_series_aligner
            )
        if DashboardMosaicLayoutTilesWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterAggregationCrossSeriesReducerEnum.to_proto(
            resource.cross_series_reducer
        ):
            res.cross_series_reducer = DashboardMosaicLayoutTilesWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterAggregationCrossSeriesReducerEnum.to_proto(
                resource.cross_series_reducer
            )
        if Primitive.to_proto(resource.group_by_fields):
            res.group_by_fields.extend(Primitive.to_proto(resource.group_by_fields))
        return res

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return None

        return DashboardMosaicLayoutTilesWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterAggregation(
            alignment_period=Primitive.from_proto(resource.alignment_period),
            per_series_aligner=DashboardMosaicLayoutTilesWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterAggregationPerSeriesAlignerEnum.from_proto(
                resource.per_series_aligner
            ),
            cross_series_reducer=DashboardMosaicLayoutTilesWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterAggregationCrossSeriesReducerEnum.from_proto(
                resource.cross_series_reducer
            ),
            group_by_fields=Primitive.from_proto(resource.group_by_fields),
        )


class DashboardMosaicLayoutTilesWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterAggregationArray(
    object
):
    @classmethod
    def to_proto(self, resources):
        if not resources:
            return resources
        return [
            DashboardMosaicLayoutTilesWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterAggregation.to_proto(
                i
            )
            for i in resources
        ]

    @classmethod
    def from_proto(self, resources):
        return [
            DashboardMosaicLayoutTilesWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterAggregation.from_proto(
                i
            )
            for i in resources
        ]


class DashboardMosaicLayoutTilesWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterSecondaryAggregation(
    object
):
    def __init__(
        self,
        alignment_period: str = None,
        per_series_aligner: str = None,
        cross_series_reducer: str = None,
        group_by_fields: list = None,
    ):
        self.alignment_period = alignment_period
        self.per_series_aligner = per_series_aligner
        self.cross_series_reducer = cross_series_reducer
        self.group_by_fields = group_by_fields

    @classmethod
    def to_proto(self, resource):
        if not resource:
            return None

        res = (
            dashboard_pb2.MonitoringBetaDashboardMosaicLayoutTilesWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterSecondaryAggregation()
        )
        if Primitive.to_proto(resource.alignment_period):
            res.alignment_period = Primitive.to_proto(resource.alignment_period)
        if DashboardMosaicLayoutTilesWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterSecondaryAggregationPerSeriesAlignerEnum.to_proto(
            resource.per_series_aligner
        ):
            res.per_series_aligner = DashboardMosaicLayoutTilesWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterSecondaryAggregationPerSeriesAlignerEnum.to_proto(
                resource.per_series_aligner
            )
        if DashboardMosaicLayoutTilesWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterSecondaryAggregationCrossSeriesReducerEnum.to_proto(
            resource.cross_series_reducer
        ):
            res.cross_series_reducer = DashboardMosaicLayoutTilesWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterSecondaryAggregationCrossSeriesReducerEnum.to_proto(
                resource.cross_series_reducer
            )
        if Primitive.to_proto(resource.group_by_fields):
            res.group_by_fields.extend(Primitive.to_proto(resource.group_by_fields))
        return res

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return None

        return DashboardMosaicLayoutTilesWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterSecondaryAggregation(
            alignment_period=Primitive.from_proto(resource.alignment_period),
            per_series_aligner=DashboardMosaicLayoutTilesWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterSecondaryAggregationPerSeriesAlignerEnum.from_proto(
                resource.per_series_aligner
            ),
            cross_series_reducer=DashboardMosaicLayoutTilesWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterSecondaryAggregationCrossSeriesReducerEnum.from_proto(
                resource.cross_series_reducer
            ),
            group_by_fields=Primitive.from_proto(resource.group_by_fields),
        )


class DashboardMosaicLayoutTilesWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterSecondaryAggregationArray(
    object
):
    @classmethod
    def to_proto(self, resources):
        if not resources:
            return resources
        return [
            DashboardMosaicLayoutTilesWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterSecondaryAggregation.to_proto(
                i
            )
            for i in resources
        ]

    @classmethod
    def from_proto(self, resources):
        return [
            DashboardMosaicLayoutTilesWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterSecondaryAggregation.from_proto(
                i
            )
            for i in resources
        ]


class DashboardMosaicLayoutTilesWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterPickTimeSeriesFilter(
    object
):
    def __init__(
        self,
        ranking_method: str = None,
        num_time_series: int = None,
        direction: str = None,
    ):
        self.ranking_method = ranking_method
        self.num_time_series = num_time_series
        self.direction = direction

    @classmethod
    def to_proto(self, resource):
        if not resource:
            return None

        res = (
            dashboard_pb2.MonitoringBetaDashboardMosaicLayoutTilesWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterPickTimeSeriesFilter()
        )
        if DashboardMosaicLayoutTilesWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterPickTimeSeriesFilterRankingMethodEnum.to_proto(
            resource.ranking_method
        ):
            res.ranking_method = DashboardMosaicLayoutTilesWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterPickTimeSeriesFilterRankingMethodEnum.to_proto(
                resource.ranking_method
            )
        if Primitive.to_proto(resource.num_time_series):
            res.num_time_series = Primitive.to_proto(resource.num_time_series)
        if DashboardMosaicLayoutTilesWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterPickTimeSeriesFilterDirectionEnum.to_proto(
            resource.direction
        ):
            res.direction = DashboardMosaicLayoutTilesWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterPickTimeSeriesFilterDirectionEnum.to_proto(
                resource.direction
            )
        return res

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return None

        return DashboardMosaicLayoutTilesWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterPickTimeSeriesFilter(
            ranking_method=DashboardMosaicLayoutTilesWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterPickTimeSeriesFilterRankingMethodEnum.from_proto(
                resource.ranking_method
            ),
            num_time_series=Primitive.from_proto(resource.num_time_series),
            direction=DashboardMosaicLayoutTilesWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterPickTimeSeriesFilterDirectionEnum.from_proto(
                resource.direction
            ),
        )


class DashboardMosaicLayoutTilesWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterPickTimeSeriesFilterArray(
    object
):
    @classmethod
    def to_proto(self, resources):
        if not resources:
            return resources
        return [
            DashboardMosaicLayoutTilesWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterPickTimeSeriesFilter.to_proto(
                i
            )
            for i in resources
        ]

    @classmethod
    def from_proto(self, resources):
        return [
            DashboardMosaicLayoutTilesWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterPickTimeSeriesFilter.from_proto(
                i
            )
            for i in resources
        ]


class DashboardMosaicLayoutTilesWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatio(
    object
):
    def __init__(
        self,
        numerator: dict = None,
        denominator: dict = None,
        secondary_aggregation: dict = None,
        pick_time_series_filter: dict = None,
    ):
        self.numerator = numerator
        self.denominator = denominator
        self.secondary_aggregation = secondary_aggregation
        self.pick_time_series_filter = pick_time_series_filter

    @classmethod
    def to_proto(self, resource):
        if not resource:
            return None

        res = (
            dashboard_pb2.MonitoringBetaDashboardMosaicLayoutTilesWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatio()
        )
        if DashboardMosaicLayoutTilesWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioNumerator.to_proto(
            resource.numerator
        ):
            res.numerator.CopyFrom(
                DashboardMosaicLayoutTilesWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioNumerator.to_proto(
                    resource.numerator
                )
            )
        else:
            res.ClearField("numerator")
        if DashboardMosaicLayoutTilesWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioDenominator.to_proto(
            resource.denominator
        ):
            res.denominator.CopyFrom(
                DashboardMosaicLayoutTilesWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioDenominator.to_proto(
                    resource.denominator
                )
            )
        else:
            res.ClearField("denominator")
        if DashboardMosaicLayoutTilesWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioSecondaryAggregation.to_proto(
            resource.secondary_aggregation
        ):
            res.secondary_aggregation.CopyFrom(
                DashboardMosaicLayoutTilesWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioSecondaryAggregation.to_proto(
                    resource.secondary_aggregation
                )
            )
        else:
            res.ClearField("secondary_aggregation")
        if DashboardMosaicLayoutTilesWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioPickTimeSeriesFilter.to_proto(
            resource.pick_time_series_filter
        ):
            res.pick_time_series_filter.CopyFrom(
                DashboardMosaicLayoutTilesWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioPickTimeSeriesFilter.to_proto(
                    resource.pick_time_series_filter
                )
            )
        else:
            res.ClearField("pick_time_series_filter")
        return res

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return None

        return DashboardMosaicLayoutTilesWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatio(
            numerator=DashboardMosaicLayoutTilesWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioNumerator.from_proto(
                resource.numerator
            ),
            denominator=DashboardMosaicLayoutTilesWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioDenominator.from_proto(
                resource.denominator
            ),
            secondary_aggregation=DashboardMosaicLayoutTilesWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioSecondaryAggregation.from_proto(
                resource.secondary_aggregation
            ),
            pick_time_series_filter=DashboardMosaicLayoutTilesWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioPickTimeSeriesFilter.from_proto(
                resource.pick_time_series_filter
            ),
        )


class DashboardMosaicLayoutTilesWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioArray(
    object
):
    @classmethod
    def to_proto(self, resources):
        if not resources:
            return resources
        return [
            DashboardMosaicLayoutTilesWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatio.to_proto(
                i
            )
            for i in resources
        ]

    @classmethod
    def from_proto(self, resources):
        return [
            DashboardMosaicLayoutTilesWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatio.from_proto(
                i
            )
            for i in resources
        ]


class DashboardMosaicLayoutTilesWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioNumerator(
    object
):
    def __init__(self, filter: str = None, aggregation: dict = None):
        self.filter = filter
        self.aggregation = aggregation

    @classmethod
    def to_proto(self, resource):
        if not resource:
            return None

        res = (
            dashboard_pb2.MonitoringBetaDashboardMosaicLayoutTilesWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioNumerator()
        )
        if Primitive.to_proto(resource.filter):
            res.filter = Primitive.to_proto(resource.filter)
        if DashboardMosaicLayoutTilesWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioNumeratorAggregation.to_proto(
            resource.aggregation
        ):
            res.aggregation.CopyFrom(
                DashboardMosaicLayoutTilesWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioNumeratorAggregation.to_proto(
                    resource.aggregation
                )
            )
        else:
            res.ClearField("aggregation")
        return res

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return None

        return DashboardMosaicLayoutTilesWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioNumerator(
            filter=Primitive.from_proto(resource.filter),
            aggregation=DashboardMosaicLayoutTilesWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioNumeratorAggregation.from_proto(
                resource.aggregation
            ),
        )


class DashboardMosaicLayoutTilesWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioNumeratorArray(
    object
):
    @classmethod
    def to_proto(self, resources):
        if not resources:
            return resources
        return [
            DashboardMosaicLayoutTilesWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioNumerator.to_proto(
                i
            )
            for i in resources
        ]

    @classmethod
    def from_proto(self, resources):
        return [
            DashboardMosaicLayoutTilesWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioNumerator.from_proto(
                i
            )
            for i in resources
        ]


class DashboardMosaicLayoutTilesWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioNumeratorAggregation(
    object
):
    def __init__(
        self,
        alignment_period: str = None,
        per_series_aligner: str = None,
        cross_series_reducer: str = None,
        group_by_fields: list = None,
    ):
        self.alignment_period = alignment_period
        self.per_series_aligner = per_series_aligner
        self.cross_series_reducer = cross_series_reducer
        self.group_by_fields = group_by_fields

    @classmethod
    def to_proto(self, resource):
        if not resource:
            return None

        res = (
            dashboard_pb2.MonitoringBetaDashboardMosaicLayoutTilesWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioNumeratorAggregation()
        )
        if Primitive.to_proto(resource.alignment_period):
            res.alignment_period = Primitive.to_proto(resource.alignment_period)
        if DashboardMosaicLayoutTilesWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioNumeratorAggregationPerSeriesAlignerEnum.to_proto(
            resource.per_series_aligner
        ):
            res.per_series_aligner = DashboardMosaicLayoutTilesWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioNumeratorAggregationPerSeriesAlignerEnum.to_proto(
                resource.per_series_aligner
            )
        if DashboardMosaicLayoutTilesWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioNumeratorAggregationCrossSeriesReducerEnum.to_proto(
            resource.cross_series_reducer
        ):
            res.cross_series_reducer = DashboardMosaicLayoutTilesWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioNumeratorAggregationCrossSeriesReducerEnum.to_proto(
                resource.cross_series_reducer
            )
        if Primitive.to_proto(resource.group_by_fields):
            res.group_by_fields.extend(Primitive.to_proto(resource.group_by_fields))
        return res

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return None

        return DashboardMosaicLayoutTilesWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioNumeratorAggregation(
            alignment_period=Primitive.from_proto(resource.alignment_period),
            per_series_aligner=DashboardMosaicLayoutTilesWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioNumeratorAggregationPerSeriesAlignerEnum.from_proto(
                resource.per_series_aligner
            ),
            cross_series_reducer=DashboardMosaicLayoutTilesWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioNumeratorAggregationCrossSeriesReducerEnum.from_proto(
                resource.cross_series_reducer
            ),
            group_by_fields=Primitive.from_proto(resource.group_by_fields),
        )


class DashboardMosaicLayoutTilesWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioNumeratorAggregationArray(
    object
):
    @classmethod
    def to_proto(self, resources):
        if not resources:
            return resources
        return [
            DashboardMosaicLayoutTilesWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioNumeratorAggregation.to_proto(
                i
            )
            for i in resources
        ]

    @classmethod
    def from_proto(self, resources):
        return [
            DashboardMosaicLayoutTilesWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioNumeratorAggregation.from_proto(
                i
            )
            for i in resources
        ]


class DashboardMosaicLayoutTilesWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioDenominator(
    object
):
    def __init__(self, filter: str = None, aggregation: dict = None):
        self.filter = filter
        self.aggregation = aggregation

    @classmethod
    def to_proto(self, resource):
        if not resource:
            return None

        res = (
            dashboard_pb2.MonitoringBetaDashboardMosaicLayoutTilesWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioDenominator()
        )
        if Primitive.to_proto(resource.filter):
            res.filter = Primitive.to_proto(resource.filter)
        if DashboardMosaicLayoutTilesWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioDenominatorAggregation.to_proto(
            resource.aggregation
        ):
            res.aggregation.CopyFrom(
                DashboardMosaicLayoutTilesWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioDenominatorAggregation.to_proto(
                    resource.aggregation
                )
            )
        else:
            res.ClearField("aggregation")
        return res

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return None

        return DashboardMosaicLayoutTilesWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioDenominator(
            filter=Primitive.from_proto(resource.filter),
            aggregation=DashboardMosaicLayoutTilesWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioDenominatorAggregation.from_proto(
                resource.aggregation
            ),
        )


class DashboardMosaicLayoutTilesWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioDenominatorArray(
    object
):
    @classmethod
    def to_proto(self, resources):
        if not resources:
            return resources
        return [
            DashboardMosaicLayoutTilesWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioDenominator.to_proto(
                i
            )
            for i in resources
        ]

    @classmethod
    def from_proto(self, resources):
        return [
            DashboardMosaicLayoutTilesWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioDenominator.from_proto(
                i
            )
            for i in resources
        ]


class DashboardMosaicLayoutTilesWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioDenominatorAggregation(
    object
):
    def __init__(
        self,
        alignment_period: str = None,
        per_series_aligner: str = None,
        cross_series_reducer: str = None,
        group_by_fields: list = None,
    ):
        self.alignment_period = alignment_period
        self.per_series_aligner = per_series_aligner
        self.cross_series_reducer = cross_series_reducer
        self.group_by_fields = group_by_fields

    @classmethod
    def to_proto(self, resource):
        if not resource:
            return None

        res = (
            dashboard_pb2.MonitoringBetaDashboardMosaicLayoutTilesWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioDenominatorAggregation()
        )
        if Primitive.to_proto(resource.alignment_period):
            res.alignment_period = Primitive.to_proto(resource.alignment_period)
        if DashboardMosaicLayoutTilesWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioDenominatorAggregationPerSeriesAlignerEnum.to_proto(
            resource.per_series_aligner
        ):
            res.per_series_aligner = DashboardMosaicLayoutTilesWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioDenominatorAggregationPerSeriesAlignerEnum.to_proto(
                resource.per_series_aligner
            )
        if DashboardMosaicLayoutTilesWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioDenominatorAggregationCrossSeriesReducerEnum.to_proto(
            resource.cross_series_reducer
        ):
            res.cross_series_reducer = DashboardMosaicLayoutTilesWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioDenominatorAggregationCrossSeriesReducerEnum.to_proto(
                resource.cross_series_reducer
            )
        if Primitive.to_proto(resource.group_by_fields):
            res.group_by_fields.extend(Primitive.to_proto(resource.group_by_fields))
        return res

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return None

        return DashboardMosaicLayoutTilesWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioDenominatorAggregation(
            alignment_period=Primitive.from_proto(resource.alignment_period),
            per_series_aligner=DashboardMosaicLayoutTilesWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioDenominatorAggregationPerSeriesAlignerEnum.from_proto(
                resource.per_series_aligner
            ),
            cross_series_reducer=DashboardMosaicLayoutTilesWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioDenominatorAggregationCrossSeriesReducerEnum.from_proto(
                resource.cross_series_reducer
            ),
            group_by_fields=Primitive.from_proto(resource.group_by_fields),
        )


class DashboardMosaicLayoutTilesWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioDenominatorAggregationArray(
    object
):
    @classmethod
    def to_proto(self, resources):
        if not resources:
            return resources
        return [
            DashboardMosaicLayoutTilesWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioDenominatorAggregation.to_proto(
                i
            )
            for i in resources
        ]

    @classmethod
    def from_proto(self, resources):
        return [
            DashboardMosaicLayoutTilesWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioDenominatorAggregation.from_proto(
                i
            )
            for i in resources
        ]


class DashboardMosaicLayoutTilesWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioSecondaryAggregation(
    object
):
    def __init__(
        self,
        alignment_period: str = None,
        per_series_aligner: str = None,
        cross_series_reducer: str = None,
        group_by_fields: list = None,
    ):
        self.alignment_period = alignment_period
        self.per_series_aligner = per_series_aligner
        self.cross_series_reducer = cross_series_reducer
        self.group_by_fields = group_by_fields

    @classmethod
    def to_proto(self, resource):
        if not resource:
            return None

        res = (
            dashboard_pb2.MonitoringBetaDashboardMosaicLayoutTilesWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioSecondaryAggregation()
        )
        if Primitive.to_proto(resource.alignment_period):
            res.alignment_period = Primitive.to_proto(resource.alignment_period)
        if DashboardMosaicLayoutTilesWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioSecondaryAggregationPerSeriesAlignerEnum.to_proto(
            resource.per_series_aligner
        ):
            res.per_series_aligner = DashboardMosaicLayoutTilesWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioSecondaryAggregationPerSeriesAlignerEnum.to_proto(
                resource.per_series_aligner
            )
        if DashboardMosaicLayoutTilesWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioSecondaryAggregationCrossSeriesReducerEnum.to_proto(
            resource.cross_series_reducer
        ):
            res.cross_series_reducer = DashboardMosaicLayoutTilesWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioSecondaryAggregationCrossSeriesReducerEnum.to_proto(
                resource.cross_series_reducer
            )
        if Primitive.to_proto(resource.group_by_fields):
            res.group_by_fields.extend(Primitive.to_proto(resource.group_by_fields))
        return res

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return None

        return DashboardMosaicLayoutTilesWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioSecondaryAggregation(
            alignment_period=Primitive.from_proto(resource.alignment_period),
            per_series_aligner=DashboardMosaicLayoutTilesWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioSecondaryAggregationPerSeriesAlignerEnum.from_proto(
                resource.per_series_aligner
            ),
            cross_series_reducer=DashboardMosaicLayoutTilesWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioSecondaryAggregationCrossSeriesReducerEnum.from_proto(
                resource.cross_series_reducer
            ),
            group_by_fields=Primitive.from_proto(resource.group_by_fields),
        )


class DashboardMosaicLayoutTilesWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioSecondaryAggregationArray(
    object
):
    @classmethod
    def to_proto(self, resources):
        if not resources:
            return resources
        return [
            DashboardMosaicLayoutTilesWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioSecondaryAggregation.to_proto(
                i
            )
            for i in resources
        ]

    @classmethod
    def from_proto(self, resources):
        return [
            DashboardMosaicLayoutTilesWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioSecondaryAggregation.from_proto(
                i
            )
            for i in resources
        ]


class DashboardMosaicLayoutTilesWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioPickTimeSeriesFilter(
    object
):
    def __init__(
        self,
        ranking_method: str = None,
        num_time_series: int = None,
        direction: str = None,
    ):
        self.ranking_method = ranking_method
        self.num_time_series = num_time_series
        self.direction = direction

    @classmethod
    def to_proto(self, resource):
        if not resource:
            return None

        res = (
            dashboard_pb2.MonitoringBetaDashboardMosaicLayoutTilesWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioPickTimeSeriesFilter()
        )
        if DashboardMosaicLayoutTilesWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioPickTimeSeriesFilterRankingMethodEnum.to_proto(
            resource.ranking_method
        ):
            res.ranking_method = DashboardMosaicLayoutTilesWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioPickTimeSeriesFilterRankingMethodEnum.to_proto(
                resource.ranking_method
            )
        if Primitive.to_proto(resource.num_time_series):
            res.num_time_series = Primitive.to_proto(resource.num_time_series)
        if DashboardMosaicLayoutTilesWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioPickTimeSeriesFilterDirectionEnum.to_proto(
            resource.direction
        ):
            res.direction = DashboardMosaicLayoutTilesWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioPickTimeSeriesFilterDirectionEnum.to_proto(
                resource.direction
            )
        return res

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return None

        return DashboardMosaicLayoutTilesWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioPickTimeSeriesFilter(
            ranking_method=DashboardMosaicLayoutTilesWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioPickTimeSeriesFilterRankingMethodEnum.from_proto(
                resource.ranking_method
            ),
            num_time_series=Primitive.from_proto(resource.num_time_series),
            direction=DashboardMosaicLayoutTilesWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioPickTimeSeriesFilterDirectionEnum.from_proto(
                resource.direction
            ),
        )


class DashboardMosaicLayoutTilesWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioPickTimeSeriesFilterArray(
    object
):
    @classmethod
    def to_proto(self, resources):
        if not resources:
            return resources
        return [
            DashboardMosaicLayoutTilesWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioPickTimeSeriesFilter.to_proto(
                i
            )
            for i in resources
        ]

    @classmethod
    def from_proto(self, resources):
        return [
            DashboardMosaicLayoutTilesWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioPickTimeSeriesFilter.from_proto(
                i
            )
            for i in resources
        ]


class DashboardMosaicLayoutTilesWidgetXyChartThresholds(object):
    def __init__(
        self,
        label: str = None,
        value: float = None,
        color: str = None,
        direction: str = None,
    ):
        self.label = label
        self.value = value
        self.color = color
        self.direction = direction

    @classmethod
    def to_proto(self, resource):
        if not resource:
            return None

        res = (
            dashboard_pb2.MonitoringBetaDashboardMosaicLayoutTilesWidgetXyChartThresholds()
        )
        if Primitive.to_proto(resource.label):
            res.label = Primitive.to_proto(resource.label)
        if Primitive.to_proto(resource.value):
            res.value = Primitive.to_proto(resource.value)
        if DashboardMosaicLayoutTilesWidgetXyChartThresholdsColorEnum.to_proto(
            resource.color
        ):
            res.color = (
                DashboardMosaicLayoutTilesWidgetXyChartThresholdsColorEnum.to_proto(
                    resource.color
                )
            )
        if DashboardMosaicLayoutTilesWidgetXyChartThresholdsDirectionEnum.to_proto(
            resource.direction
        ):
            res.direction = (
                DashboardMosaicLayoutTilesWidgetXyChartThresholdsDirectionEnum.to_proto(
                    resource.direction
                )
            )
        return res

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return None

        return DashboardMosaicLayoutTilesWidgetXyChartThresholds(
            label=Primitive.from_proto(resource.label),
            value=Primitive.from_proto(resource.value),
            color=DashboardMosaicLayoutTilesWidgetXyChartThresholdsColorEnum.from_proto(
                resource.color
            ),
            direction=DashboardMosaicLayoutTilesWidgetXyChartThresholdsDirectionEnum.from_proto(
                resource.direction
            ),
        )


class DashboardMosaicLayoutTilesWidgetXyChartThresholdsArray(object):
    @classmethod
    def to_proto(self, resources):
        if not resources:
            return resources
        return [
            DashboardMosaicLayoutTilesWidgetXyChartThresholds.to_proto(i)
            for i in resources
        ]

    @classmethod
    def from_proto(self, resources):
        return [
            DashboardMosaicLayoutTilesWidgetXyChartThresholds.from_proto(i)
            for i in resources
        ]


class DashboardMosaicLayoutTilesWidgetXyChartXAxis(object):
    def __init__(self, label: str = None, scale: str = None):
        self.label = label
        self.scale = scale

    @classmethod
    def to_proto(self, resource):
        if not resource:
            return None

        res = dashboard_pb2.MonitoringBetaDashboardMosaicLayoutTilesWidgetXyChartXAxis()
        if Primitive.to_proto(resource.label):
            res.label = Primitive.to_proto(resource.label)
        if DashboardMosaicLayoutTilesWidgetXyChartXAxisScaleEnum.to_proto(
            resource.scale
        ):
            res.scale = DashboardMosaicLayoutTilesWidgetXyChartXAxisScaleEnum.to_proto(
                resource.scale
            )
        return res

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return None

        return DashboardMosaicLayoutTilesWidgetXyChartXAxis(
            label=Primitive.from_proto(resource.label),
            scale=DashboardMosaicLayoutTilesWidgetXyChartXAxisScaleEnum.from_proto(
                resource.scale
            ),
        )


class DashboardMosaicLayoutTilesWidgetXyChartXAxisArray(object):
    @classmethod
    def to_proto(self, resources):
        if not resources:
            return resources
        return [
            DashboardMosaicLayoutTilesWidgetXyChartXAxis.to_proto(i) for i in resources
        ]

    @classmethod
    def from_proto(self, resources):
        return [
            DashboardMosaicLayoutTilesWidgetXyChartXAxis.from_proto(i)
            for i in resources
        ]


class DashboardMosaicLayoutTilesWidgetXyChartYAxis(object):
    def __init__(self, label: str = None, scale: str = None):
        self.label = label
        self.scale = scale

    @classmethod
    def to_proto(self, resource):
        if not resource:
            return None

        res = dashboard_pb2.MonitoringBetaDashboardMosaicLayoutTilesWidgetXyChartYAxis()
        if Primitive.to_proto(resource.label):
            res.label = Primitive.to_proto(resource.label)
        if DashboardMosaicLayoutTilesWidgetXyChartYAxisScaleEnum.to_proto(
            resource.scale
        ):
            res.scale = DashboardMosaicLayoutTilesWidgetXyChartYAxisScaleEnum.to_proto(
                resource.scale
            )
        return res

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return None

        return DashboardMosaicLayoutTilesWidgetXyChartYAxis(
            label=Primitive.from_proto(resource.label),
            scale=DashboardMosaicLayoutTilesWidgetXyChartYAxisScaleEnum.from_proto(
                resource.scale
            ),
        )


class DashboardMosaicLayoutTilesWidgetXyChartYAxisArray(object):
    @classmethod
    def to_proto(self, resources):
        if not resources:
            return resources
        return [
            DashboardMosaicLayoutTilesWidgetXyChartYAxis.to_proto(i) for i in resources
        ]

    @classmethod
    def from_proto(self, resources):
        return [
            DashboardMosaicLayoutTilesWidgetXyChartYAxis.from_proto(i)
            for i in resources
        ]


class DashboardMosaicLayoutTilesWidgetXyChartChartOptions(object):
    def __init__(self, mode: str = None):
        self.mode = mode

    @classmethod
    def to_proto(self, resource):
        if not resource:
            return None

        res = (
            dashboard_pb2.MonitoringBetaDashboardMosaicLayoutTilesWidgetXyChartChartOptions()
        )
        if DashboardMosaicLayoutTilesWidgetXyChartChartOptionsModeEnum.to_proto(
            resource.mode
        ):
            res.mode = (
                DashboardMosaicLayoutTilesWidgetXyChartChartOptionsModeEnum.to_proto(
                    resource.mode
                )
            )
        return res

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return None

        return DashboardMosaicLayoutTilesWidgetXyChartChartOptions(
            mode=DashboardMosaicLayoutTilesWidgetXyChartChartOptionsModeEnum.from_proto(
                resource.mode
            ),
        )


class DashboardMosaicLayoutTilesWidgetXyChartChartOptionsArray(object):
    @classmethod
    def to_proto(self, resources):
        if not resources:
            return resources
        return [
            DashboardMosaicLayoutTilesWidgetXyChartChartOptions.to_proto(i)
            for i in resources
        ]

    @classmethod
    def from_proto(self, resources):
        return [
            DashboardMosaicLayoutTilesWidgetXyChartChartOptions.from_proto(i)
            for i in resources
        ]


class DashboardMosaicLayoutTilesWidgetScorecard(object):
    def __init__(
        self,
        time_series_query: dict = None,
        gauge_view: dict = None,
        spark_chart_view: dict = None,
        thresholds: list = None,
    ):
        self.time_series_query = time_series_query
        self.gauge_view = gauge_view
        self.spark_chart_view = spark_chart_view
        self.thresholds = thresholds

    @classmethod
    def to_proto(self, resource):
        if not resource:
            return None

        res = dashboard_pb2.MonitoringBetaDashboardMosaicLayoutTilesWidgetScorecard()
        if DashboardMosaicLayoutTilesWidgetScorecardTimeSeriesQuery.to_proto(
            resource.time_series_query
        ):
            res.time_series_query.CopyFrom(
                DashboardMosaicLayoutTilesWidgetScorecardTimeSeriesQuery.to_proto(
                    resource.time_series_query
                )
            )
        else:
            res.ClearField("time_series_query")
        if DashboardMosaicLayoutTilesWidgetScorecardGaugeView.to_proto(
            resource.gauge_view
        ):
            res.gauge_view.CopyFrom(
                DashboardMosaicLayoutTilesWidgetScorecardGaugeView.to_proto(
                    resource.gauge_view
                )
            )
        else:
            res.ClearField("gauge_view")
        if DashboardMosaicLayoutTilesWidgetScorecardSparkChartView.to_proto(
            resource.spark_chart_view
        ):
            res.spark_chart_view.CopyFrom(
                DashboardMosaicLayoutTilesWidgetScorecardSparkChartView.to_proto(
                    resource.spark_chart_view
                )
            )
        else:
            res.ClearField("spark_chart_view")
        if DashboardMosaicLayoutTilesWidgetScorecardThresholdsArray.to_proto(
            resource.thresholds
        ):
            res.thresholds.extend(
                DashboardMosaicLayoutTilesWidgetScorecardThresholdsArray.to_proto(
                    resource.thresholds
                )
            )
        return res

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return None

        return DashboardMosaicLayoutTilesWidgetScorecard(
            time_series_query=DashboardMosaicLayoutTilesWidgetScorecardTimeSeriesQuery.from_proto(
                resource.time_series_query
            ),
            gauge_view=DashboardMosaicLayoutTilesWidgetScorecardGaugeView.from_proto(
                resource.gauge_view
            ),
            spark_chart_view=DashboardMosaicLayoutTilesWidgetScorecardSparkChartView.from_proto(
                resource.spark_chart_view
            ),
            thresholds=DashboardMosaicLayoutTilesWidgetScorecardThresholdsArray.from_proto(
                resource.thresholds
            ),
        )


class DashboardMosaicLayoutTilesWidgetScorecardArray(object):
    @classmethod
    def to_proto(self, resources):
        if not resources:
            return resources
        return [
            DashboardMosaicLayoutTilesWidgetScorecard.to_proto(i) for i in resources
        ]

    @classmethod
    def from_proto(self, resources):
        return [
            DashboardMosaicLayoutTilesWidgetScorecard.from_proto(i) for i in resources
        ]


class DashboardMosaicLayoutTilesWidgetScorecardTimeSeriesQuery(object):
    def __init__(
        self,
        time_series_filter: dict = None,
        time_series_filter_ratio: dict = None,
        time_series_query_language: str = None,
        unit_override: str = None,
    ):
        self.time_series_filter = time_series_filter
        self.time_series_filter_ratio = time_series_filter_ratio
        self.time_series_query_language = time_series_query_language
        self.unit_override = unit_override

    @classmethod
    def to_proto(self, resource):
        if not resource:
            return None

        res = (
            dashboard_pb2.MonitoringBetaDashboardMosaicLayoutTilesWidgetScorecardTimeSeriesQuery()
        )
        if DashboardMosaicLayoutTilesWidgetScorecardTimeSeriesQueryTimeSeriesFilter.to_proto(
            resource.time_series_filter
        ):
            res.time_series_filter.CopyFrom(
                DashboardMosaicLayoutTilesWidgetScorecardTimeSeriesQueryTimeSeriesFilter.to_proto(
                    resource.time_series_filter
                )
            )
        else:
            res.ClearField("time_series_filter")
        if DashboardMosaicLayoutTilesWidgetScorecardTimeSeriesQueryTimeSeriesFilterRatio.to_proto(
            resource.time_series_filter_ratio
        ):
            res.time_series_filter_ratio.CopyFrom(
                DashboardMosaicLayoutTilesWidgetScorecardTimeSeriesQueryTimeSeriesFilterRatio.to_proto(
                    resource.time_series_filter_ratio
                )
            )
        else:
            res.ClearField("time_series_filter_ratio")
        if Primitive.to_proto(resource.time_series_query_language):
            res.time_series_query_language = Primitive.to_proto(
                resource.time_series_query_language
            )
        if Primitive.to_proto(resource.unit_override):
            res.unit_override = Primitive.to_proto(resource.unit_override)
        return res

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return None

        return DashboardMosaicLayoutTilesWidgetScorecardTimeSeriesQuery(
            time_series_filter=DashboardMosaicLayoutTilesWidgetScorecardTimeSeriesQueryTimeSeriesFilter.from_proto(
                resource.time_series_filter
            ),
            time_series_filter_ratio=DashboardMosaicLayoutTilesWidgetScorecardTimeSeriesQueryTimeSeriesFilterRatio.from_proto(
                resource.time_series_filter_ratio
            ),
            time_series_query_language=Primitive.from_proto(
                resource.time_series_query_language
            ),
            unit_override=Primitive.from_proto(resource.unit_override),
        )


class DashboardMosaicLayoutTilesWidgetScorecardTimeSeriesQueryArray(object):
    @classmethod
    def to_proto(self, resources):
        if not resources:
            return resources
        return [
            DashboardMosaicLayoutTilesWidgetScorecardTimeSeriesQuery.to_proto(i)
            for i in resources
        ]

    @classmethod
    def from_proto(self, resources):
        return [
            DashboardMosaicLayoutTilesWidgetScorecardTimeSeriesQuery.from_proto(i)
            for i in resources
        ]


class DashboardMosaicLayoutTilesWidgetScorecardTimeSeriesQueryTimeSeriesFilter(object):
    def __init__(
        self,
        filter: str = None,
        aggregation: dict = None,
        secondary_aggregation: dict = None,
        pick_time_series_filter: dict = None,
    ):
        self.filter = filter
        self.aggregation = aggregation
        self.secondary_aggregation = secondary_aggregation
        self.pick_time_series_filter = pick_time_series_filter

    @classmethod
    def to_proto(self, resource):
        if not resource:
            return None

        res = (
            dashboard_pb2.MonitoringBetaDashboardMosaicLayoutTilesWidgetScorecardTimeSeriesQueryTimeSeriesFilter()
        )
        if Primitive.to_proto(resource.filter):
            res.filter = Primitive.to_proto(resource.filter)
        if DashboardMosaicLayoutTilesWidgetScorecardTimeSeriesQueryTimeSeriesFilterAggregation.to_proto(
            resource.aggregation
        ):
            res.aggregation.CopyFrom(
                DashboardMosaicLayoutTilesWidgetScorecardTimeSeriesQueryTimeSeriesFilterAggregation.to_proto(
                    resource.aggregation
                )
            )
        else:
            res.ClearField("aggregation")
        if DashboardMosaicLayoutTilesWidgetScorecardTimeSeriesQueryTimeSeriesFilterSecondaryAggregation.to_proto(
            resource.secondary_aggregation
        ):
            res.secondary_aggregation.CopyFrom(
                DashboardMosaicLayoutTilesWidgetScorecardTimeSeriesQueryTimeSeriesFilterSecondaryAggregation.to_proto(
                    resource.secondary_aggregation
                )
            )
        else:
            res.ClearField("secondary_aggregation")
        if DashboardMosaicLayoutTilesWidgetScorecardTimeSeriesQueryTimeSeriesFilterPickTimeSeriesFilter.to_proto(
            resource.pick_time_series_filter
        ):
            res.pick_time_series_filter.CopyFrom(
                DashboardMosaicLayoutTilesWidgetScorecardTimeSeriesQueryTimeSeriesFilterPickTimeSeriesFilter.to_proto(
                    resource.pick_time_series_filter
                )
            )
        else:
            res.ClearField("pick_time_series_filter")
        return res

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return None

        return DashboardMosaicLayoutTilesWidgetScorecardTimeSeriesQueryTimeSeriesFilter(
            filter=Primitive.from_proto(resource.filter),
            aggregation=DashboardMosaicLayoutTilesWidgetScorecardTimeSeriesQueryTimeSeriesFilterAggregation.from_proto(
                resource.aggregation
            ),
            secondary_aggregation=DashboardMosaicLayoutTilesWidgetScorecardTimeSeriesQueryTimeSeriesFilterSecondaryAggregation.from_proto(
                resource.secondary_aggregation
            ),
            pick_time_series_filter=DashboardMosaicLayoutTilesWidgetScorecardTimeSeriesQueryTimeSeriesFilterPickTimeSeriesFilter.from_proto(
                resource.pick_time_series_filter
            ),
        )


class DashboardMosaicLayoutTilesWidgetScorecardTimeSeriesQueryTimeSeriesFilterArray(
    object
):
    @classmethod
    def to_proto(self, resources):
        if not resources:
            return resources
        return [
            DashboardMosaicLayoutTilesWidgetScorecardTimeSeriesQueryTimeSeriesFilter.to_proto(
                i
            )
            for i in resources
        ]

    @classmethod
    def from_proto(self, resources):
        return [
            DashboardMosaicLayoutTilesWidgetScorecardTimeSeriesQueryTimeSeriesFilter.from_proto(
                i
            )
            for i in resources
        ]


class DashboardMosaicLayoutTilesWidgetScorecardTimeSeriesQueryTimeSeriesFilterAggregation(
    object
):
    def __init__(
        self,
        alignment_period: str = None,
        per_series_aligner: str = None,
        cross_series_reducer: str = None,
        group_by_fields: list = None,
    ):
        self.alignment_period = alignment_period
        self.per_series_aligner = per_series_aligner
        self.cross_series_reducer = cross_series_reducer
        self.group_by_fields = group_by_fields

    @classmethod
    def to_proto(self, resource):
        if not resource:
            return None

        res = (
            dashboard_pb2.MonitoringBetaDashboardMosaicLayoutTilesWidgetScorecardTimeSeriesQueryTimeSeriesFilterAggregation()
        )
        if Primitive.to_proto(resource.alignment_period):
            res.alignment_period = Primitive.to_proto(resource.alignment_period)
        if DashboardMosaicLayoutTilesWidgetScorecardTimeSeriesQueryTimeSeriesFilterAggregationPerSeriesAlignerEnum.to_proto(
            resource.per_series_aligner
        ):
            res.per_series_aligner = DashboardMosaicLayoutTilesWidgetScorecardTimeSeriesQueryTimeSeriesFilterAggregationPerSeriesAlignerEnum.to_proto(
                resource.per_series_aligner
            )
        if DashboardMosaicLayoutTilesWidgetScorecardTimeSeriesQueryTimeSeriesFilterAggregationCrossSeriesReducerEnum.to_proto(
            resource.cross_series_reducer
        ):
            res.cross_series_reducer = DashboardMosaicLayoutTilesWidgetScorecardTimeSeriesQueryTimeSeriesFilterAggregationCrossSeriesReducerEnum.to_proto(
                resource.cross_series_reducer
            )
        if Primitive.to_proto(resource.group_by_fields):
            res.group_by_fields.extend(Primitive.to_proto(resource.group_by_fields))
        return res

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return None

        return DashboardMosaicLayoutTilesWidgetScorecardTimeSeriesQueryTimeSeriesFilterAggregation(
            alignment_period=Primitive.from_proto(resource.alignment_period),
            per_series_aligner=DashboardMosaicLayoutTilesWidgetScorecardTimeSeriesQueryTimeSeriesFilterAggregationPerSeriesAlignerEnum.from_proto(
                resource.per_series_aligner
            ),
            cross_series_reducer=DashboardMosaicLayoutTilesWidgetScorecardTimeSeriesQueryTimeSeriesFilterAggregationCrossSeriesReducerEnum.from_proto(
                resource.cross_series_reducer
            ),
            group_by_fields=Primitive.from_proto(resource.group_by_fields),
        )


class DashboardMosaicLayoutTilesWidgetScorecardTimeSeriesQueryTimeSeriesFilterAggregationArray(
    object
):
    @classmethod
    def to_proto(self, resources):
        if not resources:
            return resources
        return [
            DashboardMosaicLayoutTilesWidgetScorecardTimeSeriesQueryTimeSeriesFilterAggregation.to_proto(
                i
            )
            for i in resources
        ]

    @classmethod
    def from_proto(self, resources):
        return [
            DashboardMosaicLayoutTilesWidgetScorecardTimeSeriesQueryTimeSeriesFilterAggregation.from_proto(
                i
            )
            for i in resources
        ]


class DashboardMosaicLayoutTilesWidgetScorecardTimeSeriesQueryTimeSeriesFilterSecondaryAggregation(
    object
):
    def __init__(
        self,
        alignment_period: str = None,
        per_series_aligner: str = None,
        cross_series_reducer: str = None,
        group_by_fields: list = None,
    ):
        self.alignment_period = alignment_period
        self.per_series_aligner = per_series_aligner
        self.cross_series_reducer = cross_series_reducer
        self.group_by_fields = group_by_fields

    @classmethod
    def to_proto(self, resource):
        if not resource:
            return None

        res = (
            dashboard_pb2.MonitoringBetaDashboardMosaicLayoutTilesWidgetScorecardTimeSeriesQueryTimeSeriesFilterSecondaryAggregation()
        )
        if Primitive.to_proto(resource.alignment_period):
            res.alignment_period = Primitive.to_proto(resource.alignment_period)
        if DashboardMosaicLayoutTilesWidgetScorecardTimeSeriesQueryTimeSeriesFilterSecondaryAggregationPerSeriesAlignerEnum.to_proto(
            resource.per_series_aligner
        ):
            res.per_series_aligner = DashboardMosaicLayoutTilesWidgetScorecardTimeSeriesQueryTimeSeriesFilterSecondaryAggregationPerSeriesAlignerEnum.to_proto(
                resource.per_series_aligner
            )
        if DashboardMosaicLayoutTilesWidgetScorecardTimeSeriesQueryTimeSeriesFilterSecondaryAggregationCrossSeriesReducerEnum.to_proto(
            resource.cross_series_reducer
        ):
            res.cross_series_reducer = DashboardMosaicLayoutTilesWidgetScorecardTimeSeriesQueryTimeSeriesFilterSecondaryAggregationCrossSeriesReducerEnum.to_proto(
                resource.cross_series_reducer
            )
        if Primitive.to_proto(resource.group_by_fields):
            res.group_by_fields.extend(Primitive.to_proto(resource.group_by_fields))
        return res

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return None

        return DashboardMosaicLayoutTilesWidgetScorecardTimeSeriesQueryTimeSeriesFilterSecondaryAggregation(
            alignment_period=Primitive.from_proto(resource.alignment_period),
            per_series_aligner=DashboardMosaicLayoutTilesWidgetScorecardTimeSeriesQueryTimeSeriesFilterSecondaryAggregationPerSeriesAlignerEnum.from_proto(
                resource.per_series_aligner
            ),
            cross_series_reducer=DashboardMosaicLayoutTilesWidgetScorecardTimeSeriesQueryTimeSeriesFilterSecondaryAggregationCrossSeriesReducerEnum.from_proto(
                resource.cross_series_reducer
            ),
            group_by_fields=Primitive.from_proto(resource.group_by_fields),
        )


class DashboardMosaicLayoutTilesWidgetScorecardTimeSeriesQueryTimeSeriesFilterSecondaryAggregationArray(
    object
):
    @classmethod
    def to_proto(self, resources):
        if not resources:
            return resources
        return [
            DashboardMosaicLayoutTilesWidgetScorecardTimeSeriesQueryTimeSeriesFilterSecondaryAggregation.to_proto(
                i
            )
            for i in resources
        ]

    @classmethod
    def from_proto(self, resources):
        return [
            DashboardMosaicLayoutTilesWidgetScorecardTimeSeriesQueryTimeSeriesFilterSecondaryAggregation.from_proto(
                i
            )
            for i in resources
        ]


class DashboardMosaicLayoutTilesWidgetScorecardTimeSeriesQueryTimeSeriesFilterPickTimeSeriesFilter(
    object
):
    def __init__(
        self,
        ranking_method: str = None,
        num_time_series: int = None,
        direction: str = None,
    ):
        self.ranking_method = ranking_method
        self.num_time_series = num_time_series
        self.direction = direction

    @classmethod
    def to_proto(self, resource):
        if not resource:
            return None

        res = (
            dashboard_pb2.MonitoringBetaDashboardMosaicLayoutTilesWidgetScorecardTimeSeriesQueryTimeSeriesFilterPickTimeSeriesFilter()
        )
        if DashboardMosaicLayoutTilesWidgetScorecardTimeSeriesQueryTimeSeriesFilterPickTimeSeriesFilterRankingMethodEnum.to_proto(
            resource.ranking_method
        ):
            res.ranking_method = DashboardMosaicLayoutTilesWidgetScorecardTimeSeriesQueryTimeSeriesFilterPickTimeSeriesFilterRankingMethodEnum.to_proto(
                resource.ranking_method
            )
        if Primitive.to_proto(resource.num_time_series):
            res.num_time_series = Primitive.to_proto(resource.num_time_series)
        if DashboardMosaicLayoutTilesWidgetScorecardTimeSeriesQueryTimeSeriesFilterPickTimeSeriesFilterDirectionEnum.to_proto(
            resource.direction
        ):
            res.direction = DashboardMosaicLayoutTilesWidgetScorecardTimeSeriesQueryTimeSeriesFilterPickTimeSeriesFilterDirectionEnum.to_proto(
                resource.direction
            )
        return res

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return None

        return DashboardMosaicLayoutTilesWidgetScorecardTimeSeriesQueryTimeSeriesFilterPickTimeSeriesFilter(
            ranking_method=DashboardMosaicLayoutTilesWidgetScorecardTimeSeriesQueryTimeSeriesFilterPickTimeSeriesFilterRankingMethodEnum.from_proto(
                resource.ranking_method
            ),
            num_time_series=Primitive.from_proto(resource.num_time_series),
            direction=DashboardMosaicLayoutTilesWidgetScorecardTimeSeriesQueryTimeSeriesFilterPickTimeSeriesFilterDirectionEnum.from_proto(
                resource.direction
            ),
        )


class DashboardMosaicLayoutTilesWidgetScorecardTimeSeriesQueryTimeSeriesFilterPickTimeSeriesFilterArray(
    object
):
    @classmethod
    def to_proto(self, resources):
        if not resources:
            return resources
        return [
            DashboardMosaicLayoutTilesWidgetScorecardTimeSeriesQueryTimeSeriesFilterPickTimeSeriesFilter.to_proto(
                i
            )
            for i in resources
        ]

    @classmethod
    def from_proto(self, resources):
        return [
            DashboardMosaicLayoutTilesWidgetScorecardTimeSeriesQueryTimeSeriesFilterPickTimeSeriesFilter.from_proto(
                i
            )
            for i in resources
        ]


class DashboardMosaicLayoutTilesWidgetScorecardTimeSeriesQueryTimeSeriesFilterRatio(
    object
):
    def __init__(
        self,
        numerator: dict = None,
        denominator: dict = None,
        secondary_aggregation: dict = None,
        pick_time_series_filter: dict = None,
    ):
        self.numerator = numerator
        self.denominator = denominator
        self.secondary_aggregation = secondary_aggregation
        self.pick_time_series_filter = pick_time_series_filter

    @classmethod
    def to_proto(self, resource):
        if not resource:
            return None

        res = (
            dashboard_pb2.MonitoringBetaDashboardMosaicLayoutTilesWidgetScorecardTimeSeriesQueryTimeSeriesFilterRatio()
        )
        if DashboardMosaicLayoutTilesWidgetScorecardTimeSeriesQueryTimeSeriesFilterRatioNumerator.to_proto(
            resource.numerator
        ):
            res.numerator.CopyFrom(
                DashboardMosaicLayoutTilesWidgetScorecardTimeSeriesQueryTimeSeriesFilterRatioNumerator.to_proto(
                    resource.numerator
                )
            )
        else:
            res.ClearField("numerator")
        if DashboardMosaicLayoutTilesWidgetScorecardTimeSeriesQueryTimeSeriesFilterRatioDenominator.to_proto(
            resource.denominator
        ):
            res.denominator.CopyFrom(
                DashboardMosaicLayoutTilesWidgetScorecardTimeSeriesQueryTimeSeriesFilterRatioDenominator.to_proto(
                    resource.denominator
                )
            )
        else:
            res.ClearField("denominator")
        if DashboardMosaicLayoutTilesWidgetScorecardTimeSeriesQueryTimeSeriesFilterRatioSecondaryAggregation.to_proto(
            resource.secondary_aggregation
        ):
            res.secondary_aggregation.CopyFrom(
                DashboardMosaicLayoutTilesWidgetScorecardTimeSeriesQueryTimeSeriesFilterRatioSecondaryAggregation.to_proto(
                    resource.secondary_aggregation
                )
            )
        else:
            res.ClearField("secondary_aggregation")
        if DashboardMosaicLayoutTilesWidgetScorecardTimeSeriesQueryTimeSeriesFilterRatioPickTimeSeriesFilter.to_proto(
            resource.pick_time_series_filter
        ):
            res.pick_time_series_filter.CopyFrom(
                DashboardMosaicLayoutTilesWidgetScorecardTimeSeriesQueryTimeSeriesFilterRatioPickTimeSeriesFilter.to_proto(
                    resource.pick_time_series_filter
                )
            )
        else:
            res.ClearField("pick_time_series_filter")
        return res

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return None

        return DashboardMosaicLayoutTilesWidgetScorecardTimeSeriesQueryTimeSeriesFilterRatio(
            numerator=DashboardMosaicLayoutTilesWidgetScorecardTimeSeriesQueryTimeSeriesFilterRatioNumerator.from_proto(
                resource.numerator
            ),
            denominator=DashboardMosaicLayoutTilesWidgetScorecardTimeSeriesQueryTimeSeriesFilterRatioDenominator.from_proto(
                resource.denominator
            ),
            secondary_aggregation=DashboardMosaicLayoutTilesWidgetScorecardTimeSeriesQueryTimeSeriesFilterRatioSecondaryAggregation.from_proto(
                resource.secondary_aggregation
            ),
            pick_time_series_filter=DashboardMosaicLayoutTilesWidgetScorecardTimeSeriesQueryTimeSeriesFilterRatioPickTimeSeriesFilter.from_proto(
                resource.pick_time_series_filter
            ),
        )


class DashboardMosaicLayoutTilesWidgetScorecardTimeSeriesQueryTimeSeriesFilterRatioArray(
    object
):
    @classmethod
    def to_proto(self, resources):
        if not resources:
            return resources
        return [
            DashboardMosaicLayoutTilesWidgetScorecardTimeSeriesQueryTimeSeriesFilterRatio.to_proto(
                i
            )
            for i in resources
        ]

    @classmethod
    def from_proto(self, resources):
        return [
            DashboardMosaicLayoutTilesWidgetScorecardTimeSeriesQueryTimeSeriesFilterRatio.from_proto(
                i
            )
            for i in resources
        ]


class DashboardMosaicLayoutTilesWidgetScorecardTimeSeriesQueryTimeSeriesFilterRatioNumerator(
    object
):
    def __init__(self, filter: str = None, aggregation: dict = None):
        self.filter = filter
        self.aggregation = aggregation

    @classmethod
    def to_proto(self, resource):
        if not resource:
            return None

        res = (
            dashboard_pb2.MonitoringBetaDashboardMosaicLayoutTilesWidgetScorecardTimeSeriesQueryTimeSeriesFilterRatioNumerator()
        )
        if Primitive.to_proto(resource.filter):
            res.filter = Primitive.to_proto(resource.filter)
        if DashboardMosaicLayoutTilesWidgetScorecardTimeSeriesQueryTimeSeriesFilterRatioNumeratorAggregation.to_proto(
            resource.aggregation
        ):
            res.aggregation.CopyFrom(
                DashboardMosaicLayoutTilesWidgetScorecardTimeSeriesQueryTimeSeriesFilterRatioNumeratorAggregation.to_proto(
                    resource.aggregation
                )
            )
        else:
            res.ClearField("aggregation")
        return res

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return None

        return DashboardMosaicLayoutTilesWidgetScorecardTimeSeriesQueryTimeSeriesFilterRatioNumerator(
            filter=Primitive.from_proto(resource.filter),
            aggregation=DashboardMosaicLayoutTilesWidgetScorecardTimeSeriesQueryTimeSeriesFilterRatioNumeratorAggregation.from_proto(
                resource.aggregation
            ),
        )


class DashboardMosaicLayoutTilesWidgetScorecardTimeSeriesQueryTimeSeriesFilterRatioNumeratorArray(
    object
):
    @classmethod
    def to_proto(self, resources):
        if not resources:
            return resources
        return [
            DashboardMosaicLayoutTilesWidgetScorecardTimeSeriesQueryTimeSeriesFilterRatioNumerator.to_proto(
                i
            )
            for i in resources
        ]

    @classmethod
    def from_proto(self, resources):
        return [
            DashboardMosaicLayoutTilesWidgetScorecardTimeSeriesQueryTimeSeriesFilterRatioNumerator.from_proto(
                i
            )
            for i in resources
        ]


class DashboardMosaicLayoutTilesWidgetScorecardTimeSeriesQueryTimeSeriesFilterRatioNumeratorAggregation(
    object
):
    def __init__(
        self,
        alignment_period: str = None,
        per_series_aligner: str = None,
        cross_series_reducer: str = None,
        group_by_fields: list = None,
    ):
        self.alignment_period = alignment_period
        self.per_series_aligner = per_series_aligner
        self.cross_series_reducer = cross_series_reducer
        self.group_by_fields = group_by_fields

    @classmethod
    def to_proto(self, resource):
        if not resource:
            return None

        res = (
            dashboard_pb2.MonitoringBetaDashboardMosaicLayoutTilesWidgetScorecardTimeSeriesQueryTimeSeriesFilterRatioNumeratorAggregation()
        )
        if Primitive.to_proto(resource.alignment_period):
            res.alignment_period = Primitive.to_proto(resource.alignment_period)
        if DashboardMosaicLayoutTilesWidgetScorecardTimeSeriesQueryTimeSeriesFilterRatioNumeratorAggregationPerSeriesAlignerEnum.to_proto(
            resource.per_series_aligner
        ):
            res.per_series_aligner = DashboardMosaicLayoutTilesWidgetScorecardTimeSeriesQueryTimeSeriesFilterRatioNumeratorAggregationPerSeriesAlignerEnum.to_proto(
                resource.per_series_aligner
            )
        if DashboardMosaicLayoutTilesWidgetScorecardTimeSeriesQueryTimeSeriesFilterRatioNumeratorAggregationCrossSeriesReducerEnum.to_proto(
            resource.cross_series_reducer
        ):
            res.cross_series_reducer = DashboardMosaicLayoutTilesWidgetScorecardTimeSeriesQueryTimeSeriesFilterRatioNumeratorAggregationCrossSeriesReducerEnum.to_proto(
                resource.cross_series_reducer
            )
        if Primitive.to_proto(resource.group_by_fields):
            res.group_by_fields.extend(Primitive.to_proto(resource.group_by_fields))
        return res

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return None

        return DashboardMosaicLayoutTilesWidgetScorecardTimeSeriesQueryTimeSeriesFilterRatioNumeratorAggregation(
            alignment_period=Primitive.from_proto(resource.alignment_period),
            per_series_aligner=DashboardMosaicLayoutTilesWidgetScorecardTimeSeriesQueryTimeSeriesFilterRatioNumeratorAggregationPerSeriesAlignerEnum.from_proto(
                resource.per_series_aligner
            ),
            cross_series_reducer=DashboardMosaicLayoutTilesWidgetScorecardTimeSeriesQueryTimeSeriesFilterRatioNumeratorAggregationCrossSeriesReducerEnum.from_proto(
                resource.cross_series_reducer
            ),
            group_by_fields=Primitive.from_proto(resource.group_by_fields),
        )


class DashboardMosaicLayoutTilesWidgetScorecardTimeSeriesQueryTimeSeriesFilterRatioNumeratorAggregationArray(
    object
):
    @classmethod
    def to_proto(self, resources):
        if not resources:
            return resources
        return [
            DashboardMosaicLayoutTilesWidgetScorecardTimeSeriesQueryTimeSeriesFilterRatioNumeratorAggregation.to_proto(
                i
            )
            for i in resources
        ]

    @classmethod
    def from_proto(self, resources):
        return [
            DashboardMosaicLayoutTilesWidgetScorecardTimeSeriesQueryTimeSeriesFilterRatioNumeratorAggregation.from_proto(
                i
            )
            for i in resources
        ]


class DashboardMosaicLayoutTilesWidgetScorecardTimeSeriesQueryTimeSeriesFilterRatioDenominator(
    object
):
    def __init__(self, filter: str = None, aggregation: dict = None):
        self.filter = filter
        self.aggregation = aggregation

    @classmethod
    def to_proto(self, resource):
        if not resource:
            return None

        res = (
            dashboard_pb2.MonitoringBetaDashboardMosaicLayoutTilesWidgetScorecardTimeSeriesQueryTimeSeriesFilterRatioDenominator()
        )
        if Primitive.to_proto(resource.filter):
            res.filter = Primitive.to_proto(resource.filter)
        if DashboardMosaicLayoutTilesWidgetScorecardTimeSeriesQueryTimeSeriesFilterRatioDenominatorAggregation.to_proto(
            resource.aggregation
        ):
            res.aggregation.CopyFrom(
                DashboardMosaicLayoutTilesWidgetScorecardTimeSeriesQueryTimeSeriesFilterRatioDenominatorAggregation.to_proto(
                    resource.aggregation
                )
            )
        else:
            res.ClearField("aggregation")
        return res

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return None

        return DashboardMosaicLayoutTilesWidgetScorecardTimeSeriesQueryTimeSeriesFilterRatioDenominator(
            filter=Primitive.from_proto(resource.filter),
            aggregation=DashboardMosaicLayoutTilesWidgetScorecardTimeSeriesQueryTimeSeriesFilterRatioDenominatorAggregation.from_proto(
                resource.aggregation
            ),
        )


class DashboardMosaicLayoutTilesWidgetScorecardTimeSeriesQueryTimeSeriesFilterRatioDenominatorArray(
    object
):
    @classmethod
    def to_proto(self, resources):
        if not resources:
            return resources
        return [
            DashboardMosaicLayoutTilesWidgetScorecardTimeSeriesQueryTimeSeriesFilterRatioDenominator.to_proto(
                i
            )
            for i in resources
        ]

    @classmethod
    def from_proto(self, resources):
        return [
            DashboardMosaicLayoutTilesWidgetScorecardTimeSeriesQueryTimeSeriesFilterRatioDenominator.from_proto(
                i
            )
            for i in resources
        ]


class DashboardMosaicLayoutTilesWidgetScorecardTimeSeriesQueryTimeSeriesFilterRatioDenominatorAggregation(
    object
):
    def __init__(
        self,
        alignment_period: str = None,
        per_series_aligner: str = None,
        cross_series_reducer: str = None,
        group_by_fields: list = None,
    ):
        self.alignment_period = alignment_period
        self.per_series_aligner = per_series_aligner
        self.cross_series_reducer = cross_series_reducer
        self.group_by_fields = group_by_fields

    @classmethod
    def to_proto(self, resource):
        if not resource:
            return None

        res = (
            dashboard_pb2.MonitoringBetaDashboardMosaicLayoutTilesWidgetScorecardTimeSeriesQueryTimeSeriesFilterRatioDenominatorAggregation()
        )
        if Primitive.to_proto(resource.alignment_period):
            res.alignment_period = Primitive.to_proto(resource.alignment_period)
        if DashboardMosaicLayoutTilesWidgetScorecardTimeSeriesQueryTimeSeriesFilterRatioDenominatorAggregationPerSeriesAlignerEnum.to_proto(
            resource.per_series_aligner
        ):
            res.per_series_aligner = DashboardMosaicLayoutTilesWidgetScorecardTimeSeriesQueryTimeSeriesFilterRatioDenominatorAggregationPerSeriesAlignerEnum.to_proto(
                resource.per_series_aligner
            )
        if DashboardMosaicLayoutTilesWidgetScorecardTimeSeriesQueryTimeSeriesFilterRatioDenominatorAggregationCrossSeriesReducerEnum.to_proto(
            resource.cross_series_reducer
        ):
            res.cross_series_reducer = DashboardMosaicLayoutTilesWidgetScorecardTimeSeriesQueryTimeSeriesFilterRatioDenominatorAggregationCrossSeriesReducerEnum.to_proto(
                resource.cross_series_reducer
            )
        if Primitive.to_proto(resource.group_by_fields):
            res.group_by_fields.extend(Primitive.to_proto(resource.group_by_fields))
        return res

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return None

        return DashboardMosaicLayoutTilesWidgetScorecardTimeSeriesQueryTimeSeriesFilterRatioDenominatorAggregation(
            alignment_period=Primitive.from_proto(resource.alignment_period),
            per_series_aligner=DashboardMosaicLayoutTilesWidgetScorecardTimeSeriesQueryTimeSeriesFilterRatioDenominatorAggregationPerSeriesAlignerEnum.from_proto(
                resource.per_series_aligner
            ),
            cross_series_reducer=DashboardMosaicLayoutTilesWidgetScorecardTimeSeriesQueryTimeSeriesFilterRatioDenominatorAggregationCrossSeriesReducerEnum.from_proto(
                resource.cross_series_reducer
            ),
            group_by_fields=Primitive.from_proto(resource.group_by_fields),
        )


class DashboardMosaicLayoutTilesWidgetScorecardTimeSeriesQueryTimeSeriesFilterRatioDenominatorAggregationArray(
    object
):
    @classmethod
    def to_proto(self, resources):
        if not resources:
            return resources
        return [
            DashboardMosaicLayoutTilesWidgetScorecardTimeSeriesQueryTimeSeriesFilterRatioDenominatorAggregation.to_proto(
                i
            )
            for i in resources
        ]

    @classmethod
    def from_proto(self, resources):
        return [
            DashboardMosaicLayoutTilesWidgetScorecardTimeSeriesQueryTimeSeriesFilterRatioDenominatorAggregation.from_proto(
                i
            )
            for i in resources
        ]


class DashboardMosaicLayoutTilesWidgetScorecardTimeSeriesQueryTimeSeriesFilterRatioSecondaryAggregation(
    object
):
    def __init__(
        self,
        alignment_period: str = None,
        per_series_aligner: str = None,
        cross_series_reducer: str = None,
        group_by_fields: list = None,
    ):
        self.alignment_period = alignment_period
        self.per_series_aligner = per_series_aligner
        self.cross_series_reducer = cross_series_reducer
        self.group_by_fields = group_by_fields

    @classmethod
    def to_proto(self, resource):
        if not resource:
            return None

        res = (
            dashboard_pb2.MonitoringBetaDashboardMosaicLayoutTilesWidgetScorecardTimeSeriesQueryTimeSeriesFilterRatioSecondaryAggregation()
        )
        if Primitive.to_proto(resource.alignment_period):
            res.alignment_period = Primitive.to_proto(resource.alignment_period)
        if DashboardMosaicLayoutTilesWidgetScorecardTimeSeriesQueryTimeSeriesFilterRatioSecondaryAggregationPerSeriesAlignerEnum.to_proto(
            resource.per_series_aligner
        ):
            res.per_series_aligner = DashboardMosaicLayoutTilesWidgetScorecardTimeSeriesQueryTimeSeriesFilterRatioSecondaryAggregationPerSeriesAlignerEnum.to_proto(
                resource.per_series_aligner
            )
        if DashboardMosaicLayoutTilesWidgetScorecardTimeSeriesQueryTimeSeriesFilterRatioSecondaryAggregationCrossSeriesReducerEnum.to_proto(
            resource.cross_series_reducer
        ):
            res.cross_series_reducer = DashboardMosaicLayoutTilesWidgetScorecardTimeSeriesQueryTimeSeriesFilterRatioSecondaryAggregationCrossSeriesReducerEnum.to_proto(
                resource.cross_series_reducer
            )
        if Primitive.to_proto(resource.group_by_fields):
            res.group_by_fields.extend(Primitive.to_proto(resource.group_by_fields))
        return res

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return None

        return DashboardMosaicLayoutTilesWidgetScorecardTimeSeriesQueryTimeSeriesFilterRatioSecondaryAggregation(
            alignment_period=Primitive.from_proto(resource.alignment_period),
            per_series_aligner=DashboardMosaicLayoutTilesWidgetScorecardTimeSeriesQueryTimeSeriesFilterRatioSecondaryAggregationPerSeriesAlignerEnum.from_proto(
                resource.per_series_aligner
            ),
            cross_series_reducer=DashboardMosaicLayoutTilesWidgetScorecardTimeSeriesQueryTimeSeriesFilterRatioSecondaryAggregationCrossSeriesReducerEnum.from_proto(
                resource.cross_series_reducer
            ),
            group_by_fields=Primitive.from_proto(resource.group_by_fields),
        )


class DashboardMosaicLayoutTilesWidgetScorecardTimeSeriesQueryTimeSeriesFilterRatioSecondaryAggregationArray(
    object
):
    @classmethod
    def to_proto(self, resources):
        if not resources:
            return resources
        return [
            DashboardMosaicLayoutTilesWidgetScorecardTimeSeriesQueryTimeSeriesFilterRatioSecondaryAggregation.to_proto(
                i
            )
            for i in resources
        ]

    @classmethod
    def from_proto(self, resources):
        return [
            DashboardMosaicLayoutTilesWidgetScorecardTimeSeriesQueryTimeSeriesFilterRatioSecondaryAggregation.from_proto(
                i
            )
            for i in resources
        ]


class DashboardMosaicLayoutTilesWidgetScorecardTimeSeriesQueryTimeSeriesFilterRatioPickTimeSeriesFilter(
    object
):
    def __init__(
        self,
        ranking_method: str = None,
        num_time_series: int = None,
        direction: str = None,
    ):
        self.ranking_method = ranking_method
        self.num_time_series = num_time_series
        self.direction = direction

    @classmethod
    def to_proto(self, resource):
        if not resource:
            return None

        res = (
            dashboard_pb2.MonitoringBetaDashboardMosaicLayoutTilesWidgetScorecardTimeSeriesQueryTimeSeriesFilterRatioPickTimeSeriesFilter()
        )
        if DashboardMosaicLayoutTilesWidgetScorecardTimeSeriesQueryTimeSeriesFilterRatioPickTimeSeriesFilterRankingMethodEnum.to_proto(
            resource.ranking_method
        ):
            res.ranking_method = DashboardMosaicLayoutTilesWidgetScorecardTimeSeriesQueryTimeSeriesFilterRatioPickTimeSeriesFilterRankingMethodEnum.to_proto(
                resource.ranking_method
            )
        if Primitive.to_proto(resource.num_time_series):
            res.num_time_series = Primitive.to_proto(resource.num_time_series)
        if DashboardMosaicLayoutTilesWidgetScorecardTimeSeriesQueryTimeSeriesFilterRatioPickTimeSeriesFilterDirectionEnum.to_proto(
            resource.direction
        ):
            res.direction = DashboardMosaicLayoutTilesWidgetScorecardTimeSeriesQueryTimeSeriesFilterRatioPickTimeSeriesFilterDirectionEnum.to_proto(
                resource.direction
            )
        return res

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return None

        return DashboardMosaicLayoutTilesWidgetScorecardTimeSeriesQueryTimeSeriesFilterRatioPickTimeSeriesFilter(
            ranking_method=DashboardMosaicLayoutTilesWidgetScorecardTimeSeriesQueryTimeSeriesFilterRatioPickTimeSeriesFilterRankingMethodEnum.from_proto(
                resource.ranking_method
            ),
            num_time_series=Primitive.from_proto(resource.num_time_series),
            direction=DashboardMosaicLayoutTilesWidgetScorecardTimeSeriesQueryTimeSeriesFilterRatioPickTimeSeriesFilterDirectionEnum.from_proto(
                resource.direction
            ),
        )


class DashboardMosaicLayoutTilesWidgetScorecardTimeSeriesQueryTimeSeriesFilterRatioPickTimeSeriesFilterArray(
    object
):
    @classmethod
    def to_proto(self, resources):
        if not resources:
            return resources
        return [
            DashboardMosaicLayoutTilesWidgetScorecardTimeSeriesQueryTimeSeriesFilterRatioPickTimeSeriesFilter.to_proto(
                i
            )
            for i in resources
        ]

    @classmethod
    def from_proto(self, resources):
        return [
            DashboardMosaicLayoutTilesWidgetScorecardTimeSeriesQueryTimeSeriesFilterRatioPickTimeSeriesFilter.from_proto(
                i
            )
            for i in resources
        ]


class DashboardMosaicLayoutTilesWidgetScorecardGaugeView(object):
    def __init__(self, lower_bound: float = None, upper_bound: float = None):
        self.lower_bound = lower_bound
        self.upper_bound = upper_bound

    @classmethod
    def to_proto(self, resource):
        if not resource:
            return None

        res = (
            dashboard_pb2.MonitoringBetaDashboardMosaicLayoutTilesWidgetScorecardGaugeView()
        )
        if Primitive.to_proto(resource.lower_bound):
            res.lower_bound = Primitive.to_proto(resource.lower_bound)
        if Primitive.to_proto(resource.upper_bound):
            res.upper_bound = Primitive.to_proto(resource.upper_bound)
        return res

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return None

        return DashboardMosaicLayoutTilesWidgetScorecardGaugeView(
            lower_bound=Primitive.from_proto(resource.lower_bound),
            upper_bound=Primitive.from_proto(resource.upper_bound),
        )


class DashboardMosaicLayoutTilesWidgetScorecardGaugeViewArray(object):
    @classmethod
    def to_proto(self, resources):
        if not resources:
            return resources
        return [
            DashboardMosaicLayoutTilesWidgetScorecardGaugeView.to_proto(i)
            for i in resources
        ]

    @classmethod
    def from_proto(self, resources):
        return [
            DashboardMosaicLayoutTilesWidgetScorecardGaugeView.from_proto(i)
            for i in resources
        ]


class DashboardMosaicLayoutTilesWidgetScorecardSparkChartView(object):
    def __init__(self, spark_chart_type: str = None, min_alignment_period: str = None):
        self.spark_chart_type = spark_chart_type
        self.min_alignment_period = min_alignment_period

    @classmethod
    def to_proto(self, resource):
        if not resource:
            return None

        res = (
            dashboard_pb2.MonitoringBetaDashboardMosaicLayoutTilesWidgetScorecardSparkChartView()
        )
        if DashboardMosaicLayoutTilesWidgetScorecardSparkChartViewSparkChartTypeEnum.to_proto(
            resource.spark_chart_type
        ):
            res.spark_chart_type = DashboardMosaicLayoutTilesWidgetScorecardSparkChartViewSparkChartTypeEnum.to_proto(
                resource.spark_chart_type
            )
        if Primitive.to_proto(resource.min_alignment_period):
            res.min_alignment_period = Primitive.to_proto(resource.min_alignment_period)
        return res

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return None

        return DashboardMosaicLayoutTilesWidgetScorecardSparkChartView(
            spark_chart_type=DashboardMosaicLayoutTilesWidgetScorecardSparkChartViewSparkChartTypeEnum.from_proto(
                resource.spark_chart_type
            ),
            min_alignment_period=Primitive.from_proto(resource.min_alignment_period),
        )


class DashboardMosaicLayoutTilesWidgetScorecardSparkChartViewArray(object):
    @classmethod
    def to_proto(self, resources):
        if not resources:
            return resources
        return [
            DashboardMosaicLayoutTilesWidgetScorecardSparkChartView.to_proto(i)
            for i in resources
        ]

    @classmethod
    def from_proto(self, resources):
        return [
            DashboardMosaicLayoutTilesWidgetScorecardSparkChartView.from_proto(i)
            for i in resources
        ]


class DashboardMosaicLayoutTilesWidgetScorecardThresholds(object):
    def __init__(
        self,
        label: str = None,
        value: float = None,
        color: str = None,
        direction: str = None,
    ):
        self.label = label
        self.value = value
        self.color = color
        self.direction = direction

    @classmethod
    def to_proto(self, resource):
        if not resource:
            return None

        res = (
            dashboard_pb2.MonitoringBetaDashboardMosaicLayoutTilesWidgetScorecardThresholds()
        )
        if Primitive.to_proto(resource.label):
            res.label = Primitive.to_proto(resource.label)
        if Primitive.to_proto(resource.value):
            res.value = Primitive.to_proto(resource.value)
        if DashboardMosaicLayoutTilesWidgetScorecardThresholdsColorEnum.to_proto(
            resource.color
        ):
            res.color = (
                DashboardMosaicLayoutTilesWidgetScorecardThresholdsColorEnum.to_proto(
                    resource.color
                )
            )
        if DashboardMosaicLayoutTilesWidgetScorecardThresholdsDirectionEnum.to_proto(
            resource.direction
        ):
            res.direction = DashboardMosaicLayoutTilesWidgetScorecardThresholdsDirectionEnum.to_proto(
                resource.direction
            )
        return res

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return None

        return DashboardMosaicLayoutTilesWidgetScorecardThresholds(
            label=Primitive.from_proto(resource.label),
            value=Primitive.from_proto(resource.value),
            color=DashboardMosaicLayoutTilesWidgetScorecardThresholdsColorEnum.from_proto(
                resource.color
            ),
            direction=DashboardMosaicLayoutTilesWidgetScorecardThresholdsDirectionEnum.from_proto(
                resource.direction
            ),
        )


class DashboardMosaicLayoutTilesWidgetScorecardThresholdsArray(object):
    @classmethod
    def to_proto(self, resources):
        if not resources:
            return resources
        return [
            DashboardMosaicLayoutTilesWidgetScorecardThresholds.to_proto(i)
            for i in resources
        ]

    @classmethod
    def from_proto(self, resources):
        return [
            DashboardMosaicLayoutTilesWidgetScorecardThresholds.from_proto(i)
            for i in resources
        ]


class DashboardMosaicLayoutTilesWidgetText(object):
    def __init__(self, content: str = None, format: str = None):
        self.content = content
        self.format = format

    @classmethod
    def to_proto(self, resource):
        if not resource:
            return None

        res = dashboard_pb2.MonitoringBetaDashboardMosaicLayoutTilesWidgetText()
        if Primitive.to_proto(resource.content):
            res.content = Primitive.to_proto(resource.content)
        if DashboardMosaicLayoutTilesWidgetTextFormatEnum.to_proto(resource.format):
            res.format = DashboardMosaicLayoutTilesWidgetTextFormatEnum.to_proto(
                resource.format
            )
        return res

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return None

        return DashboardMosaicLayoutTilesWidgetText(
            content=Primitive.from_proto(resource.content),
            format=DashboardMosaicLayoutTilesWidgetTextFormatEnum.from_proto(
                resource.format
            ),
        )


class DashboardMosaicLayoutTilesWidgetTextArray(object):
    @classmethod
    def to_proto(self, resources):
        if not resources:
            return resources
        return [DashboardMosaicLayoutTilesWidgetText.to_proto(i) for i in resources]

    @classmethod
    def from_proto(self, resources):
        return [DashboardMosaicLayoutTilesWidgetText.from_proto(i) for i in resources]


class DashboardMosaicLayoutTilesWidgetBlank(object):
    @classmethod
    def to_proto(self, resource):
        if not resource:
            return None

        res = dashboard_pb2.MonitoringBetaDashboardMosaicLayoutTilesWidgetBlank()
        return res

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return None

        return DashboardMosaicLayoutTilesWidgetBlank()


class DashboardMosaicLayoutTilesWidgetBlankArray(object):
    @classmethod
    def to_proto(self, resources):
        if not resources:
            return resources
        return [DashboardMosaicLayoutTilesWidgetBlank.to_proto(i) for i in resources]

    @classmethod
    def from_proto(self, resources):
        return [DashboardMosaicLayoutTilesWidgetBlank.from_proto(i) for i in resources]


class DashboardMosaicLayoutTilesWidgetLogsPanel(object):
    def __init__(self, filter: str = None, resource_names: list = None):
        self.filter = filter
        self.resource_names = resource_names

    @classmethod
    def to_proto(self, resource):
        if not resource:
            return None

        res = dashboard_pb2.MonitoringBetaDashboardMosaicLayoutTilesWidgetLogsPanel()
        if Primitive.to_proto(resource.filter):
            res.filter = Primitive.to_proto(resource.filter)
        if Primitive.to_proto(resource.resource_names):
            res.resource_names.extend(Primitive.to_proto(resource.resource_names))
        return res

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return None

        return DashboardMosaicLayoutTilesWidgetLogsPanel(
            filter=Primitive.from_proto(resource.filter),
            resource_names=Primitive.from_proto(resource.resource_names),
        )


class DashboardMosaicLayoutTilesWidgetLogsPanelArray(object):
    @classmethod
    def to_proto(self, resources):
        if not resources:
            return resources
        return [
            DashboardMosaicLayoutTilesWidgetLogsPanel.to_proto(i) for i in resources
        ]

    @classmethod
    def from_proto(self, resources):
        return [
            DashboardMosaicLayoutTilesWidgetLogsPanel.from_proto(i) for i in resources
        ]


class DashboardRowLayout(object):
    def __init__(self, rows: list = None):
        self.rows = rows

    @classmethod
    def to_proto(self, resource):
        if not resource:
            return None

        res = dashboard_pb2.MonitoringBetaDashboardRowLayout()
        if DashboardRowLayoutRowsArray.to_proto(resource.rows):
            res.rows.extend(DashboardRowLayoutRowsArray.to_proto(resource.rows))
        return res

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return None

        return DashboardRowLayout(
            rows=DashboardRowLayoutRowsArray.from_proto(resource.rows),
        )


class DashboardRowLayoutArray(object):
    @classmethod
    def to_proto(self, resources):
        if not resources:
            return resources
        return [DashboardRowLayout.to_proto(i) for i in resources]

    @classmethod
    def from_proto(self, resources):
        return [DashboardRowLayout.from_proto(i) for i in resources]


class DashboardRowLayoutRows(object):
    def __init__(self, weight: int = None, widgets: list = None):
        self.weight = weight
        self.widgets = widgets

    @classmethod
    def to_proto(self, resource):
        if not resource:
            return None

        res = dashboard_pb2.MonitoringBetaDashboardRowLayoutRows()
        if Primitive.to_proto(resource.weight):
            res.weight = Primitive.to_proto(resource.weight)
        if DashboardRowLayoutRowsWidgetsArray.to_proto(resource.widgets):
            res.widgets.extend(
                DashboardRowLayoutRowsWidgetsArray.to_proto(resource.widgets)
            )
        return res

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return None

        return DashboardRowLayoutRows(
            weight=Primitive.from_proto(resource.weight),
            widgets=DashboardRowLayoutRowsWidgetsArray.from_proto(resource.widgets),
        )


class DashboardRowLayoutRowsArray(object):
    @classmethod
    def to_proto(self, resources):
        if not resources:
            return resources
        return [DashboardRowLayoutRows.to_proto(i) for i in resources]

    @classmethod
    def from_proto(self, resources):
        return [DashboardRowLayoutRows.from_proto(i) for i in resources]


class DashboardRowLayoutRowsWidgets(object):
    def __init__(
        self,
        title: str = None,
        xy_chart: dict = None,
        scorecard: dict = None,
        text: dict = None,
        blank: dict = None,
        logs_panel: dict = None,
    ):
        self.title = title
        self.xy_chart = xy_chart
        self.scorecard = scorecard
        self.text = text
        self.blank = blank
        self.logs_panel = logs_panel

    @classmethod
    def to_proto(self, resource):
        if not resource:
            return None

        res = dashboard_pb2.MonitoringBetaDashboardRowLayoutRowsWidgets()
        if Primitive.to_proto(resource.title):
            res.title = Primitive.to_proto(resource.title)
        if DashboardRowLayoutRowsWidgetsXyChart.to_proto(resource.xy_chart):
            res.xy_chart.CopyFrom(
                DashboardRowLayoutRowsWidgetsXyChart.to_proto(resource.xy_chart)
            )
        else:
            res.ClearField("xy_chart")
        if DashboardRowLayoutRowsWidgetsScorecard.to_proto(resource.scorecard):
            res.scorecard.CopyFrom(
                DashboardRowLayoutRowsWidgetsScorecard.to_proto(resource.scorecard)
            )
        else:
            res.ClearField("scorecard")
        if DashboardRowLayoutRowsWidgetsText.to_proto(resource.text):
            res.text.CopyFrom(DashboardRowLayoutRowsWidgetsText.to_proto(resource.text))
        else:
            res.ClearField("text")
        if DashboardRowLayoutRowsWidgetsBlank.to_proto(resource.blank):
            res.blank.CopyFrom(
                DashboardRowLayoutRowsWidgetsBlank.to_proto(resource.blank)
            )
        else:
            res.ClearField("blank")
        if DashboardRowLayoutRowsWidgetsLogsPanel.to_proto(resource.logs_panel):
            res.logs_panel.CopyFrom(
                DashboardRowLayoutRowsWidgetsLogsPanel.to_proto(resource.logs_panel)
            )
        else:
            res.ClearField("logs_panel")
        return res

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return None

        return DashboardRowLayoutRowsWidgets(
            title=Primitive.from_proto(resource.title),
            xy_chart=DashboardRowLayoutRowsWidgetsXyChart.from_proto(resource.xy_chart),
            scorecard=DashboardRowLayoutRowsWidgetsScorecard.from_proto(
                resource.scorecard
            ),
            text=DashboardRowLayoutRowsWidgetsText.from_proto(resource.text),
            blank=DashboardRowLayoutRowsWidgetsBlank.from_proto(resource.blank),
            logs_panel=DashboardRowLayoutRowsWidgetsLogsPanel.from_proto(
                resource.logs_panel
            ),
        )


class DashboardRowLayoutRowsWidgetsArray(object):
    @classmethod
    def to_proto(self, resources):
        if not resources:
            return resources
        return [DashboardRowLayoutRowsWidgets.to_proto(i) for i in resources]

    @classmethod
    def from_proto(self, resources):
        return [DashboardRowLayoutRowsWidgets.from_proto(i) for i in resources]


class DashboardRowLayoutRowsWidgetsXyChart(object):
    def __init__(
        self,
        data_sets: list = None,
        timeshift_duration: str = None,
        thresholds: list = None,
        x_axis: dict = None,
        y_axis: dict = None,
        chart_options: dict = None,
    ):
        self.data_sets = data_sets
        self.timeshift_duration = timeshift_duration
        self.thresholds = thresholds
        self.x_axis = x_axis
        self.y_axis = y_axis
        self.chart_options = chart_options

    @classmethod
    def to_proto(self, resource):
        if not resource:
            return None

        res = dashboard_pb2.MonitoringBetaDashboardRowLayoutRowsWidgetsXyChart()
        if DashboardRowLayoutRowsWidgetsXyChartDataSetsArray.to_proto(
            resource.data_sets
        ):
            res.data_sets.extend(
                DashboardRowLayoutRowsWidgetsXyChartDataSetsArray.to_proto(
                    resource.data_sets
                )
            )
        if Primitive.to_proto(resource.timeshift_duration):
            res.timeshift_duration = Primitive.to_proto(resource.timeshift_duration)
        if DashboardRowLayoutRowsWidgetsXyChartThresholdsArray.to_proto(
            resource.thresholds
        ):
            res.thresholds.extend(
                DashboardRowLayoutRowsWidgetsXyChartThresholdsArray.to_proto(
                    resource.thresholds
                )
            )
        if DashboardRowLayoutRowsWidgetsXyChartXAxis.to_proto(resource.x_axis):
            res.x_axis.CopyFrom(
                DashboardRowLayoutRowsWidgetsXyChartXAxis.to_proto(resource.x_axis)
            )
        else:
            res.ClearField("x_axis")
        if DashboardRowLayoutRowsWidgetsXyChartYAxis.to_proto(resource.y_axis):
            res.y_axis.CopyFrom(
                DashboardRowLayoutRowsWidgetsXyChartYAxis.to_proto(resource.y_axis)
            )
        else:
            res.ClearField("y_axis")
        if DashboardRowLayoutRowsWidgetsXyChartChartOptions.to_proto(
            resource.chart_options
        ):
            res.chart_options.CopyFrom(
                DashboardRowLayoutRowsWidgetsXyChartChartOptions.to_proto(
                    resource.chart_options
                )
            )
        else:
            res.ClearField("chart_options")
        return res

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return None

        return DashboardRowLayoutRowsWidgetsXyChart(
            data_sets=DashboardRowLayoutRowsWidgetsXyChartDataSetsArray.from_proto(
                resource.data_sets
            ),
            timeshift_duration=Primitive.from_proto(resource.timeshift_duration),
            thresholds=DashboardRowLayoutRowsWidgetsXyChartThresholdsArray.from_proto(
                resource.thresholds
            ),
            x_axis=DashboardRowLayoutRowsWidgetsXyChartXAxis.from_proto(
                resource.x_axis
            ),
            y_axis=DashboardRowLayoutRowsWidgetsXyChartYAxis.from_proto(
                resource.y_axis
            ),
            chart_options=DashboardRowLayoutRowsWidgetsXyChartChartOptions.from_proto(
                resource.chart_options
            ),
        )


class DashboardRowLayoutRowsWidgetsXyChartArray(object):
    @classmethod
    def to_proto(self, resources):
        if not resources:
            return resources
        return [DashboardRowLayoutRowsWidgetsXyChart.to_proto(i) for i in resources]

    @classmethod
    def from_proto(self, resources):
        return [DashboardRowLayoutRowsWidgetsXyChart.from_proto(i) for i in resources]


class DashboardRowLayoutRowsWidgetsXyChartDataSets(object):
    def __init__(
        self,
        time_series_query: dict = None,
        plot_type: str = None,
        legend_template: str = None,
        min_alignment_period: str = None,
    ):
        self.time_series_query = time_series_query
        self.plot_type = plot_type
        self.legend_template = legend_template
        self.min_alignment_period = min_alignment_period

    @classmethod
    def to_proto(self, resource):
        if not resource:
            return None

        res = dashboard_pb2.MonitoringBetaDashboardRowLayoutRowsWidgetsXyChartDataSets()
        if DashboardRowLayoutRowsWidgetsXyChartDataSetsTimeSeriesQuery.to_proto(
            resource.time_series_query
        ):
            res.time_series_query.CopyFrom(
                DashboardRowLayoutRowsWidgetsXyChartDataSetsTimeSeriesQuery.to_proto(
                    resource.time_series_query
                )
            )
        else:
            res.ClearField("time_series_query")
        if DashboardRowLayoutRowsWidgetsXyChartDataSetsPlotTypeEnum.to_proto(
            resource.plot_type
        ):
            res.plot_type = (
                DashboardRowLayoutRowsWidgetsXyChartDataSetsPlotTypeEnum.to_proto(
                    resource.plot_type
                )
            )
        if Primitive.to_proto(resource.legend_template):
            res.legend_template = Primitive.to_proto(resource.legend_template)
        if Primitive.to_proto(resource.min_alignment_period):
            res.min_alignment_period = Primitive.to_proto(resource.min_alignment_period)
        return res

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return None

        return DashboardRowLayoutRowsWidgetsXyChartDataSets(
            time_series_query=DashboardRowLayoutRowsWidgetsXyChartDataSetsTimeSeriesQuery.from_proto(
                resource.time_series_query
            ),
            plot_type=DashboardRowLayoutRowsWidgetsXyChartDataSetsPlotTypeEnum.from_proto(
                resource.plot_type
            ),
            legend_template=Primitive.from_proto(resource.legend_template),
            min_alignment_period=Primitive.from_proto(resource.min_alignment_period),
        )


class DashboardRowLayoutRowsWidgetsXyChartDataSetsArray(object):
    @classmethod
    def to_proto(self, resources):
        if not resources:
            return resources
        return [
            DashboardRowLayoutRowsWidgetsXyChartDataSets.to_proto(i) for i in resources
        ]

    @classmethod
    def from_proto(self, resources):
        return [
            DashboardRowLayoutRowsWidgetsXyChartDataSets.from_proto(i)
            for i in resources
        ]


class DashboardRowLayoutRowsWidgetsXyChartDataSetsTimeSeriesQuery(object):
    def __init__(
        self,
        time_series_filter: dict = None,
        time_series_filter_ratio: dict = None,
        time_series_query_language: str = None,
        unit_override: str = None,
    ):
        self.time_series_filter = time_series_filter
        self.time_series_filter_ratio = time_series_filter_ratio
        self.time_series_query_language = time_series_query_language
        self.unit_override = unit_override

    @classmethod
    def to_proto(self, resource):
        if not resource:
            return None

        res = (
            dashboard_pb2.MonitoringBetaDashboardRowLayoutRowsWidgetsXyChartDataSetsTimeSeriesQuery()
        )
        if DashboardRowLayoutRowsWidgetsXyChartDataSetsTimeSeriesQueryTimeSeriesFilter.to_proto(
            resource.time_series_filter
        ):
            res.time_series_filter.CopyFrom(
                DashboardRowLayoutRowsWidgetsXyChartDataSetsTimeSeriesQueryTimeSeriesFilter.to_proto(
                    resource.time_series_filter
                )
            )
        else:
            res.ClearField("time_series_filter")
        if DashboardRowLayoutRowsWidgetsXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatio.to_proto(
            resource.time_series_filter_ratio
        ):
            res.time_series_filter_ratio.CopyFrom(
                DashboardRowLayoutRowsWidgetsXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatio.to_proto(
                    resource.time_series_filter_ratio
                )
            )
        else:
            res.ClearField("time_series_filter_ratio")
        if Primitive.to_proto(resource.time_series_query_language):
            res.time_series_query_language = Primitive.to_proto(
                resource.time_series_query_language
            )
        if Primitive.to_proto(resource.unit_override):
            res.unit_override = Primitive.to_proto(resource.unit_override)
        return res

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return None

        return DashboardRowLayoutRowsWidgetsXyChartDataSetsTimeSeriesQuery(
            time_series_filter=DashboardRowLayoutRowsWidgetsXyChartDataSetsTimeSeriesQueryTimeSeriesFilter.from_proto(
                resource.time_series_filter
            ),
            time_series_filter_ratio=DashboardRowLayoutRowsWidgetsXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatio.from_proto(
                resource.time_series_filter_ratio
            ),
            time_series_query_language=Primitive.from_proto(
                resource.time_series_query_language
            ),
            unit_override=Primitive.from_proto(resource.unit_override),
        )


class DashboardRowLayoutRowsWidgetsXyChartDataSetsTimeSeriesQueryArray(object):
    @classmethod
    def to_proto(self, resources):
        if not resources:
            return resources
        return [
            DashboardRowLayoutRowsWidgetsXyChartDataSetsTimeSeriesQuery.to_proto(i)
            for i in resources
        ]

    @classmethod
    def from_proto(self, resources):
        return [
            DashboardRowLayoutRowsWidgetsXyChartDataSetsTimeSeriesQuery.from_proto(i)
            for i in resources
        ]


class DashboardRowLayoutRowsWidgetsXyChartDataSetsTimeSeriesQueryTimeSeriesFilter(
    object
):
    def __init__(
        self,
        filter: str = None,
        aggregation: dict = None,
        secondary_aggregation: dict = None,
        pick_time_series_filter: dict = None,
    ):
        self.filter = filter
        self.aggregation = aggregation
        self.secondary_aggregation = secondary_aggregation
        self.pick_time_series_filter = pick_time_series_filter

    @classmethod
    def to_proto(self, resource):
        if not resource:
            return None

        res = (
            dashboard_pb2.MonitoringBetaDashboardRowLayoutRowsWidgetsXyChartDataSetsTimeSeriesQueryTimeSeriesFilter()
        )
        if Primitive.to_proto(resource.filter):
            res.filter = Primitive.to_proto(resource.filter)
        if DashboardRowLayoutRowsWidgetsXyChartDataSetsTimeSeriesQueryTimeSeriesFilterAggregation.to_proto(
            resource.aggregation
        ):
            res.aggregation.CopyFrom(
                DashboardRowLayoutRowsWidgetsXyChartDataSetsTimeSeriesQueryTimeSeriesFilterAggregation.to_proto(
                    resource.aggregation
                )
            )
        else:
            res.ClearField("aggregation")
        if DashboardRowLayoutRowsWidgetsXyChartDataSetsTimeSeriesQueryTimeSeriesFilterSecondaryAggregation.to_proto(
            resource.secondary_aggregation
        ):
            res.secondary_aggregation.CopyFrom(
                DashboardRowLayoutRowsWidgetsXyChartDataSetsTimeSeriesQueryTimeSeriesFilterSecondaryAggregation.to_proto(
                    resource.secondary_aggregation
                )
            )
        else:
            res.ClearField("secondary_aggregation")
        if DashboardRowLayoutRowsWidgetsXyChartDataSetsTimeSeriesQueryTimeSeriesFilterPickTimeSeriesFilter.to_proto(
            resource.pick_time_series_filter
        ):
            res.pick_time_series_filter.CopyFrom(
                DashboardRowLayoutRowsWidgetsXyChartDataSetsTimeSeriesQueryTimeSeriesFilterPickTimeSeriesFilter.to_proto(
                    resource.pick_time_series_filter
                )
            )
        else:
            res.ClearField("pick_time_series_filter")
        return res

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return None

        return DashboardRowLayoutRowsWidgetsXyChartDataSetsTimeSeriesQueryTimeSeriesFilter(
            filter=Primitive.from_proto(resource.filter),
            aggregation=DashboardRowLayoutRowsWidgetsXyChartDataSetsTimeSeriesQueryTimeSeriesFilterAggregation.from_proto(
                resource.aggregation
            ),
            secondary_aggregation=DashboardRowLayoutRowsWidgetsXyChartDataSetsTimeSeriesQueryTimeSeriesFilterSecondaryAggregation.from_proto(
                resource.secondary_aggregation
            ),
            pick_time_series_filter=DashboardRowLayoutRowsWidgetsXyChartDataSetsTimeSeriesQueryTimeSeriesFilterPickTimeSeriesFilter.from_proto(
                resource.pick_time_series_filter
            ),
        )


class DashboardRowLayoutRowsWidgetsXyChartDataSetsTimeSeriesQueryTimeSeriesFilterArray(
    object
):
    @classmethod
    def to_proto(self, resources):
        if not resources:
            return resources
        return [
            DashboardRowLayoutRowsWidgetsXyChartDataSetsTimeSeriesQueryTimeSeriesFilter.to_proto(
                i
            )
            for i in resources
        ]

    @classmethod
    def from_proto(self, resources):
        return [
            DashboardRowLayoutRowsWidgetsXyChartDataSetsTimeSeriesQueryTimeSeriesFilter.from_proto(
                i
            )
            for i in resources
        ]


class DashboardRowLayoutRowsWidgetsXyChartDataSetsTimeSeriesQueryTimeSeriesFilterAggregation(
    object
):
    def __init__(
        self,
        alignment_period: str = None,
        per_series_aligner: str = None,
        cross_series_reducer: str = None,
        group_by_fields: list = None,
    ):
        self.alignment_period = alignment_period
        self.per_series_aligner = per_series_aligner
        self.cross_series_reducer = cross_series_reducer
        self.group_by_fields = group_by_fields

    @classmethod
    def to_proto(self, resource):
        if not resource:
            return None

        res = (
            dashboard_pb2.MonitoringBetaDashboardRowLayoutRowsWidgetsXyChartDataSetsTimeSeriesQueryTimeSeriesFilterAggregation()
        )
        if Primitive.to_proto(resource.alignment_period):
            res.alignment_period = Primitive.to_proto(resource.alignment_period)
        if DashboardRowLayoutRowsWidgetsXyChartDataSetsTimeSeriesQueryTimeSeriesFilterAggregationPerSeriesAlignerEnum.to_proto(
            resource.per_series_aligner
        ):
            res.per_series_aligner = DashboardRowLayoutRowsWidgetsXyChartDataSetsTimeSeriesQueryTimeSeriesFilterAggregationPerSeriesAlignerEnum.to_proto(
                resource.per_series_aligner
            )
        if DashboardRowLayoutRowsWidgetsXyChartDataSetsTimeSeriesQueryTimeSeriesFilterAggregationCrossSeriesReducerEnum.to_proto(
            resource.cross_series_reducer
        ):
            res.cross_series_reducer = DashboardRowLayoutRowsWidgetsXyChartDataSetsTimeSeriesQueryTimeSeriesFilterAggregationCrossSeriesReducerEnum.to_proto(
                resource.cross_series_reducer
            )
        if Primitive.to_proto(resource.group_by_fields):
            res.group_by_fields.extend(Primitive.to_proto(resource.group_by_fields))
        return res

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return None

        return DashboardRowLayoutRowsWidgetsXyChartDataSetsTimeSeriesQueryTimeSeriesFilterAggregation(
            alignment_period=Primitive.from_proto(resource.alignment_period),
            per_series_aligner=DashboardRowLayoutRowsWidgetsXyChartDataSetsTimeSeriesQueryTimeSeriesFilterAggregationPerSeriesAlignerEnum.from_proto(
                resource.per_series_aligner
            ),
            cross_series_reducer=DashboardRowLayoutRowsWidgetsXyChartDataSetsTimeSeriesQueryTimeSeriesFilterAggregationCrossSeriesReducerEnum.from_proto(
                resource.cross_series_reducer
            ),
            group_by_fields=Primitive.from_proto(resource.group_by_fields),
        )


class DashboardRowLayoutRowsWidgetsXyChartDataSetsTimeSeriesQueryTimeSeriesFilterAggregationArray(
    object
):
    @classmethod
    def to_proto(self, resources):
        if not resources:
            return resources
        return [
            DashboardRowLayoutRowsWidgetsXyChartDataSetsTimeSeriesQueryTimeSeriesFilterAggregation.to_proto(
                i
            )
            for i in resources
        ]

    @classmethod
    def from_proto(self, resources):
        return [
            DashboardRowLayoutRowsWidgetsXyChartDataSetsTimeSeriesQueryTimeSeriesFilterAggregation.from_proto(
                i
            )
            for i in resources
        ]


class DashboardRowLayoutRowsWidgetsXyChartDataSetsTimeSeriesQueryTimeSeriesFilterSecondaryAggregation(
    object
):
    def __init__(
        self,
        alignment_period: str = None,
        per_series_aligner: str = None,
        cross_series_reducer: str = None,
        group_by_fields: list = None,
    ):
        self.alignment_period = alignment_period
        self.per_series_aligner = per_series_aligner
        self.cross_series_reducer = cross_series_reducer
        self.group_by_fields = group_by_fields

    @classmethod
    def to_proto(self, resource):
        if not resource:
            return None

        res = (
            dashboard_pb2.MonitoringBetaDashboardRowLayoutRowsWidgetsXyChartDataSetsTimeSeriesQueryTimeSeriesFilterSecondaryAggregation()
        )
        if Primitive.to_proto(resource.alignment_period):
            res.alignment_period = Primitive.to_proto(resource.alignment_period)
        if DashboardRowLayoutRowsWidgetsXyChartDataSetsTimeSeriesQueryTimeSeriesFilterSecondaryAggregationPerSeriesAlignerEnum.to_proto(
            resource.per_series_aligner
        ):
            res.per_series_aligner = DashboardRowLayoutRowsWidgetsXyChartDataSetsTimeSeriesQueryTimeSeriesFilterSecondaryAggregationPerSeriesAlignerEnum.to_proto(
                resource.per_series_aligner
            )
        if DashboardRowLayoutRowsWidgetsXyChartDataSetsTimeSeriesQueryTimeSeriesFilterSecondaryAggregationCrossSeriesReducerEnum.to_proto(
            resource.cross_series_reducer
        ):
            res.cross_series_reducer = DashboardRowLayoutRowsWidgetsXyChartDataSetsTimeSeriesQueryTimeSeriesFilterSecondaryAggregationCrossSeriesReducerEnum.to_proto(
                resource.cross_series_reducer
            )
        if Primitive.to_proto(resource.group_by_fields):
            res.group_by_fields.extend(Primitive.to_proto(resource.group_by_fields))
        return res

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return None

        return DashboardRowLayoutRowsWidgetsXyChartDataSetsTimeSeriesQueryTimeSeriesFilterSecondaryAggregation(
            alignment_period=Primitive.from_proto(resource.alignment_period),
            per_series_aligner=DashboardRowLayoutRowsWidgetsXyChartDataSetsTimeSeriesQueryTimeSeriesFilterSecondaryAggregationPerSeriesAlignerEnum.from_proto(
                resource.per_series_aligner
            ),
            cross_series_reducer=DashboardRowLayoutRowsWidgetsXyChartDataSetsTimeSeriesQueryTimeSeriesFilterSecondaryAggregationCrossSeriesReducerEnum.from_proto(
                resource.cross_series_reducer
            ),
            group_by_fields=Primitive.from_proto(resource.group_by_fields),
        )


class DashboardRowLayoutRowsWidgetsXyChartDataSetsTimeSeriesQueryTimeSeriesFilterSecondaryAggregationArray(
    object
):
    @classmethod
    def to_proto(self, resources):
        if not resources:
            return resources
        return [
            DashboardRowLayoutRowsWidgetsXyChartDataSetsTimeSeriesQueryTimeSeriesFilterSecondaryAggregation.to_proto(
                i
            )
            for i in resources
        ]

    @classmethod
    def from_proto(self, resources):
        return [
            DashboardRowLayoutRowsWidgetsXyChartDataSetsTimeSeriesQueryTimeSeriesFilterSecondaryAggregation.from_proto(
                i
            )
            for i in resources
        ]


class DashboardRowLayoutRowsWidgetsXyChartDataSetsTimeSeriesQueryTimeSeriesFilterPickTimeSeriesFilter(
    object
):
    def __init__(
        self,
        ranking_method: str = None,
        num_time_series: int = None,
        direction: str = None,
    ):
        self.ranking_method = ranking_method
        self.num_time_series = num_time_series
        self.direction = direction

    @classmethod
    def to_proto(self, resource):
        if not resource:
            return None

        res = (
            dashboard_pb2.MonitoringBetaDashboardRowLayoutRowsWidgetsXyChartDataSetsTimeSeriesQueryTimeSeriesFilterPickTimeSeriesFilter()
        )
        if DashboardRowLayoutRowsWidgetsXyChartDataSetsTimeSeriesQueryTimeSeriesFilterPickTimeSeriesFilterRankingMethodEnum.to_proto(
            resource.ranking_method
        ):
            res.ranking_method = DashboardRowLayoutRowsWidgetsXyChartDataSetsTimeSeriesQueryTimeSeriesFilterPickTimeSeriesFilterRankingMethodEnum.to_proto(
                resource.ranking_method
            )
        if Primitive.to_proto(resource.num_time_series):
            res.num_time_series = Primitive.to_proto(resource.num_time_series)
        if DashboardRowLayoutRowsWidgetsXyChartDataSetsTimeSeriesQueryTimeSeriesFilterPickTimeSeriesFilterDirectionEnum.to_proto(
            resource.direction
        ):
            res.direction = DashboardRowLayoutRowsWidgetsXyChartDataSetsTimeSeriesQueryTimeSeriesFilterPickTimeSeriesFilterDirectionEnum.to_proto(
                resource.direction
            )
        return res

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return None

        return DashboardRowLayoutRowsWidgetsXyChartDataSetsTimeSeriesQueryTimeSeriesFilterPickTimeSeriesFilter(
            ranking_method=DashboardRowLayoutRowsWidgetsXyChartDataSetsTimeSeriesQueryTimeSeriesFilterPickTimeSeriesFilterRankingMethodEnum.from_proto(
                resource.ranking_method
            ),
            num_time_series=Primitive.from_proto(resource.num_time_series),
            direction=DashboardRowLayoutRowsWidgetsXyChartDataSetsTimeSeriesQueryTimeSeriesFilterPickTimeSeriesFilterDirectionEnum.from_proto(
                resource.direction
            ),
        )


class DashboardRowLayoutRowsWidgetsXyChartDataSetsTimeSeriesQueryTimeSeriesFilterPickTimeSeriesFilterArray(
    object
):
    @classmethod
    def to_proto(self, resources):
        if not resources:
            return resources
        return [
            DashboardRowLayoutRowsWidgetsXyChartDataSetsTimeSeriesQueryTimeSeriesFilterPickTimeSeriesFilter.to_proto(
                i
            )
            for i in resources
        ]

    @classmethod
    def from_proto(self, resources):
        return [
            DashboardRowLayoutRowsWidgetsXyChartDataSetsTimeSeriesQueryTimeSeriesFilterPickTimeSeriesFilter.from_proto(
                i
            )
            for i in resources
        ]


class DashboardRowLayoutRowsWidgetsXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatio(
    object
):
    def __init__(
        self,
        numerator: dict = None,
        denominator: dict = None,
        secondary_aggregation: dict = None,
        pick_time_series_filter: dict = None,
    ):
        self.numerator = numerator
        self.denominator = denominator
        self.secondary_aggregation = secondary_aggregation
        self.pick_time_series_filter = pick_time_series_filter

    @classmethod
    def to_proto(self, resource):
        if not resource:
            return None

        res = (
            dashboard_pb2.MonitoringBetaDashboardRowLayoutRowsWidgetsXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatio()
        )
        if DashboardRowLayoutRowsWidgetsXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioNumerator.to_proto(
            resource.numerator
        ):
            res.numerator.CopyFrom(
                DashboardRowLayoutRowsWidgetsXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioNumerator.to_proto(
                    resource.numerator
                )
            )
        else:
            res.ClearField("numerator")
        if DashboardRowLayoutRowsWidgetsXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioDenominator.to_proto(
            resource.denominator
        ):
            res.denominator.CopyFrom(
                DashboardRowLayoutRowsWidgetsXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioDenominator.to_proto(
                    resource.denominator
                )
            )
        else:
            res.ClearField("denominator")
        if DashboardRowLayoutRowsWidgetsXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioSecondaryAggregation.to_proto(
            resource.secondary_aggregation
        ):
            res.secondary_aggregation.CopyFrom(
                DashboardRowLayoutRowsWidgetsXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioSecondaryAggregation.to_proto(
                    resource.secondary_aggregation
                )
            )
        else:
            res.ClearField("secondary_aggregation")
        if DashboardRowLayoutRowsWidgetsXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioPickTimeSeriesFilter.to_proto(
            resource.pick_time_series_filter
        ):
            res.pick_time_series_filter.CopyFrom(
                DashboardRowLayoutRowsWidgetsXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioPickTimeSeriesFilter.to_proto(
                    resource.pick_time_series_filter
                )
            )
        else:
            res.ClearField("pick_time_series_filter")
        return res

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return None

        return DashboardRowLayoutRowsWidgetsXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatio(
            numerator=DashboardRowLayoutRowsWidgetsXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioNumerator.from_proto(
                resource.numerator
            ),
            denominator=DashboardRowLayoutRowsWidgetsXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioDenominator.from_proto(
                resource.denominator
            ),
            secondary_aggregation=DashboardRowLayoutRowsWidgetsXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioSecondaryAggregation.from_proto(
                resource.secondary_aggregation
            ),
            pick_time_series_filter=DashboardRowLayoutRowsWidgetsXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioPickTimeSeriesFilter.from_proto(
                resource.pick_time_series_filter
            ),
        )


class DashboardRowLayoutRowsWidgetsXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioArray(
    object
):
    @classmethod
    def to_proto(self, resources):
        if not resources:
            return resources
        return [
            DashboardRowLayoutRowsWidgetsXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatio.to_proto(
                i
            )
            for i in resources
        ]

    @classmethod
    def from_proto(self, resources):
        return [
            DashboardRowLayoutRowsWidgetsXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatio.from_proto(
                i
            )
            for i in resources
        ]


class DashboardRowLayoutRowsWidgetsXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioNumerator(
    object
):
    def __init__(self, filter: str = None, aggregation: dict = None):
        self.filter = filter
        self.aggregation = aggregation

    @classmethod
    def to_proto(self, resource):
        if not resource:
            return None

        res = (
            dashboard_pb2.MonitoringBetaDashboardRowLayoutRowsWidgetsXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioNumerator()
        )
        if Primitive.to_proto(resource.filter):
            res.filter = Primitive.to_proto(resource.filter)
        if DashboardRowLayoutRowsWidgetsXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioNumeratorAggregation.to_proto(
            resource.aggregation
        ):
            res.aggregation.CopyFrom(
                DashboardRowLayoutRowsWidgetsXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioNumeratorAggregation.to_proto(
                    resource.aggregation
                )
            )
        else:
            res.ClearField("aggregation")
        return res

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return None

        return DashboardRowLayoutRowsWidgetsXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioNumerator(
            filter=Primitive.from_proto(resource.filter),
            aggregation=DashboardRowLayoutRowsWidgetsXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioNumeratorAggregation.from_proto(
                resource.aggregation
            ),
        )


class DashboardRowLayoutRowsWidgetsXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioNumeratorArray(
    object
):
    @classmethod
    def to_proto(self, resources):
        if not resources:
            return resources
        return [
            DashboardRowLayoutRowsWidgetsXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioNumerator.to_proto(
                i
            )
            for i in resources
        ]

    @classmethod
    def from_proto(self, resources):
        return [
            DashboardRowLayoutRowsWidgetsXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioNumerator.from_proto(
                i
            )
            for i in resources
        ]


class DashboardRowLayoutRowsWidgetsXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioNumeratorAggregation(
    object
):
    def __init__(
        self,
        alignment_period: str = None,
        per_series_aligner: str = None,
        cross_series_reducer: str = None,
        group_by_fields: list = None,
    ):
        self.alignment_period = alignment_period
        self.per_series_aligner = per_series_aligner
        self.cross_series_reducer = cross_series_reducer
        self.group_by_fields = group_by_fields

    @classmethod
    def to_proto(self, resource):
        if not resource:
            return None

        res = (
            dashboard_pb2.MonitoringBetaDashboardRowLayoutRowsWidgetsXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioNumeratorAggregation()
        )
        if Primitive.to_proto(resource.alignment_period):
            res.alignment_period = Primitive.to_proto(resource.alignment_period)
        if DashboardRowLayoutRowsWidgetsXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioNumeratorAggregationPerSeriesAlignerEnum.to_proto(
            resource.per_series_aligner
        ):
            res.per_series_aligner = DashboardRowLayoutRowsWidgetsXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioNumeratorAggregationPerSeriesAlignerEnum.to_proto(
                resource.per_series_aligner
            )
        if DashboardRowLayoutRowsWidgetsXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioNumeratorAggregationCrossSeriesReducerEnum.to_proto(
            resource.cross_series_reducer
        ):
            res.cross_series_reducer = DashboardRowLayoutRowsWidgetsXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioNumeratorAggregationCrossSeriesReducerEnum.to_proto(
                resource.cross_series_reducer
            )
        if Primitive.to_proto(resource.group_by_fields):
            res.group_by_fields.extend(Primitive.to_proto(resource.group_by_fields))
        return res

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return None

        return DashboardRowLayoutRowsWidgetsXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioNumeratorAggregation(
            alignment_period=Primitive.from_proto(resource.alignment_period),
            per_series_aligner=DashboardRowLayoutRowsWidgetsXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioNumeratorAggregationPerSeriesAlignerEnum.from_proto(
                resource.per_series_aligner
            ),
            cross_series_reducer=DashboardRowLayoutRowsWidgetsXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioNumeratorAggregationCrossSeriesReducerEnum.from_proto(
                resource.cross_series_reducer
            ),
            group_by_fields=Primitive.from_proto(resource.group_by_fields),
        )


class DashboardRowLayoutRowsWidgetsXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioNumeratorAggregationArray(
    object
):
    @classmethod
    def to_proto(self, resources):
        if not resources:
            return resources
        return [
            DashboardRowLayoutRowsWidgetsXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioNumeratorAggregation.to_proto(
                i
            )
            for i in resources
        ]

    @classmethod
    def from_proto(self, resources):
        return [
            DashboardRowLayoutRowsWidgetsXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioNumeratorAggregation.from_proto(
                i
            )
            for i in resources
        ]


class DashboardRowLayoutRowsWidgetsXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioDenominator(
    object
):
    def __init__(self, filter: str = None, aggregation: dict = None):
        self.filter = filter
        self.aggregation = aggregation

    @classmethod
    def to_proto(self, resource):
        if not resource:
            return None

        res = (
            dashboard_pb2.MonitoringBetaDashboardRowLayoutRowsWidgetsXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioDenominator()
        )
        if Primitive.to_proto(resource.filter):
            res.filter = Primitive.to_proto(resource.filter)
        if DashboardRowLayoutRowsWidgetsXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioDenominatorAggregation.to_proto(
            resource.aggregation
        ):
            res.aggregation.CopyFrom(
                DashboardRowLayoutRowsWidgetsXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioDenominatorAggregation.to_proto(
                    resource.aggregation
                )
            )
        else:
            res.ClearField("aggregation")
        return res

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return None

        return DashboardRowLayoutRowsWidgetsXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioDenominator(
            filter=Primitive.from_proto(resource.filter),
            aggregation=DashboardRowLayoutRowsWidgetsXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioDenominatorAggregation.from_proto(
                resource.aggregation
            ),
        )


class DashboardRowLayoutRowsWidgetsXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioDenominatorArray(
    object
):
    @classmethod
    def to_proto(self, resources):
        if not resources:
            return resources
        return [
            DashboardRowLayoutRowsWidgetsXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioDenominator.to_proto(
                i
            )
            for i in resources
        ]

    @classmethod
    def from_proto(self, resources):
        return [
            DashboardRowLayoutRowsWidgetsXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioDenominator.from_proto(
                i
            )
            for i in resources
        ]


class DashboardRowLayoutRowsWidgetsXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioDenominatorAggregation(
    object
):
    def __init__(
        self,
        alignment_period: str = None,
        per_series_aligner: str = None,
        cross_series_reducer: str = None,
        group_by_fields: list = None,
    ):
        self.alignment_period = alignment_period
        self.per_series_aligner = per_series_aligner
        self.cross_series_reducer = cross_series_reducer
        self.group_by_fields = group_by_fields

    @classmethod
    def to_proto(self, resource):
        if not resource:
            return None

        res = (
            dashboard_pb2.MonitoringBetaDashboardRowLayoutRowsWidgetsXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioDenominatorAggregation()
        )
        if Primitive.to_proto(resource.alignment_period):
            res.alignment_period = Primitive.to_proto(resource.alignment_period)
        if DashboardRowLayoutRowsWidgetsXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioDenominatorAggregationPerSeriesAlignerEnum.to_proto(
            resource.per_series_aligner
        ):
            res.per_series_aligner = DashboardRowLayoutRowsWidgetsXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioDenominatorAggregationPerSeriesAlignerEnum.to_proto(
                resource.per_series_aligner
            )
        if DashboardRowLayoutRowsWidgetsXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioDenominatorAggregationCrossSeriesReducerEnum.to_proto(
            resource.cross_series_reducer
        ):
            res.cross_series_reducer = DashboardRowLayoutRowsWidgetsXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioDenominatorAggregationCrossSeriesReducerEnum.to_proto(
                resource.cross_series_reducer
            )
        if Primitive.to_proto(resource.group_by_fields):
            res.group_by_fields.extend(Primitive.to_proto(resource.group_by_fields))
        return res

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return None

        return DashboardRowLayoutRowsWidgetsXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioDenominatorAggregation(
            alignment_period=Primitive.from_proto(resource.alignment_period),
            per_series_aligner=DashboardRowLayoutRowsWidgetsXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioDenominatorAggregationPerSeriesAlignerEnum.from_proto(
                resource.per_series_aligner
            ),
            cross_series_reducer=DashboardRowLayoutRowsWidgetsXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioDenominatorAggregationCrossSeriesReducerEnum.from_proto(
                resource.cross_series_reducer
            ),
            group_by_fields=Primitive.from_proto(resource.group_by_fields),
        )


class DashboardRowLayoutRowsWidgetsXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioDenominatorAggregationArray(
    object
):
    @classmethod
    def to_proto(self, resources):
        if not resources:
            return resources
        return [
            DashboardRowLayoutRowsWidgetsXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioDenominatorAggregation.to_proto(
                i
            )
            for i in resources
        ]

    @classmethod
    def from_proto(self, resources):
        return [
            DashboardRowLayoutRowsWidgetsXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioDenominatorAggregation.from_proto(
                i
            )
            for i in resources
        ]


class DashboardRowLayoutRowsWidgetsXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioSecondaryAggregation(
    object
):
    def __init__(
        self,
        alignment_period: str = None,
        per_series_aligner: str = None,
        cross_series_reducer: str = None,
        group_by_fields: list = None,
    ):
        self.alignment_period = alignment_period
        self.per_series_aligner = per_series_aligner
        self.cross_series_reducer = cross_series_reducer
        self.group_by_fields = group_by_fields

    @classmethod
    def to_proto(self, resource):
        if not resource:
            return None

        res = (
            dashboard_pb2.MonitoringBetaDashboardRowLayoutRowsWidgetsXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioSecondaryAggregation()
        )
        if Primitive.to_proto(resource.alignment_period):
            res.alignment_period = Primitive.to_proto(resource.alignment_period)
        if DashboardRowLayoutRowsWidgetsXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioSecondaryAggregationPerSeriesAlignerEnum.to_proto(
            resource.per_series_aligner
        ):
            res.per_series_aligner = DashboardRowLayoutRowsWidgetsXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioSecondaryAggregationPerSeriesAlignerEnum.to_proto(
                resource.per_series_aligner
            )
        if DashboardRowLayoutRowsWidgetsXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioSecondaryAggregationCrossSeriesReducerEnum.to_proto(
            resource.cross_series_reducer
        ):
            res.cross_series_reducer = DashboardRowLayoutRowsWidgetsXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioSecondaryAggregationCrossSeriesReducerEnum.to_proto(
                resource.cross_series_reducer
            )
        if Primitive.to_proto(resource.group_by_fields):
            res.group_by_fields.extend(Primitive.to_proto(resource.group_by_fields))
        return res

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return None

        return DashboardRowLayoutRowsWidgetsXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioSecondaryAggregation(
            alignment_period=Primitive.from_proto(resource.alignment_period),
            per_series_aligner=DashboardRowLayoutRowsWidgetsXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioSecondaryAggregationPerSeriesAlignerEnum.from_proto(
                resource.per_series_aligner
            ),
            cross_series_reducer=DashboardRowLayoutRowsWidgetsXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioSecondaryAggregationCrossSeriesReducerEnum.from_proto(
                resource.cross_series_reducer
            ),
            group_by_fields=Primitive.from_proto(resource.group_by_fields),
        )


class DashboardRowLayoutRowsWidgetsXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioSecondaryAggregationArray(
    object
):
    @classmethod
    def to_proto(self, resources):
        if not resources:
            return resources
        return [
            DashboardRowLayoutRowsWidgetsXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioSecondaryAggregation.to_proto(
                i
            )
            for i in resources
        ]

    @classmethod
    def from_proto(self, resources):
        return [
            DashboardRowLayoutRowsWidgetsXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioSecondaryAggregation.from_proto(
                i
            )
            for i in resources
        ]


class DashboardRowLayoutRowsWidgetsXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioPickTimeSeriesFilter(
    object
):
    def __init__(
        self,
        ranking_method: str = None,
        num_time_series: int = None,
        direction: str = None,
    ):
        self.ranking_method = ranking_method
        self.num_time_series = num_time_series
        self.direction = direction

    @classmethod
    def to_proto(self, resource):
        if not resource:
            return None

        res = (
            dashboard_pb2.MonitoringBetaDashboardRowLayoutRowsWidgetsXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioPickTimeSeriesFilter()
        )
        if DashboardRowLayoutRowsWidgetsXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioPickTimeSeriesFilterRankingMethodEnum.to_proto(
            resource.ranking_method
        ):
            res.ranking_method = DashboardRowLayoutRowsWidgetsXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioPickTimeSeriesFilterRankingMethodEnum.to_proto(
                resource.ranking_method
            )
        if Primitive.to_proto(resource.num_time_series):
            res.num_time_series = Primitive.to_proto(resource.num_time_series)
        if DashboardRowLayoutRowsWidgetsXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioPickTimeSeriesFilterDirectionEnum.to_proto(
            resource.direction
        ):
            res.direction = DashboardRowLayoutRowsWidgetsXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioPickTimeSeriesFilterDirectionEnum.to_proto(
                resource.direction
            )
        return res

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return None

        return DashboardRowLayoutRowsWidgetsXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioPickTimeSeriesFilter(
            ranking_method=DashboardRowLayoutRowsWidgetsXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioPickTimeSeriesFilterRankingMethodEnum.from_proto(
                resource.ranking_method
            ),
            num_time_series=Primitive.from_proto(resource.num_time_series),
            direction=DashboardRowLayoutRowsWidgetsXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioPickTimeSeriesFilterDirectionEnum.from_proto(
                resource.direction
            ),
        )


class DashboardRowLayoutRowsWidgetsXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioPickTimeSeriesFilterArray(
    object
):
    @classmethod
    def to_proto(self, resources):
        if not resources:
            return resources
        return [
            DashboardRowLayoutRowsWidgetsXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioPickTimeSeriesFilter.to_proto(
                i
            )
            for i in resources
        ]

    @classmethod
    def from_proto(self, resources):
        return [
            DashboardRowLayoutRowsWidgetsXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioPickTimeSeriesFilter.from_proto(
                i
            )
            for i in resources
        ]


class DashboardRowLayoutRowsWidgetsXyChartThresholds(object):
    def __init__(
        self,
        label: str = None,
        value: float = None,
        color: str = None,
        direction: str = None,
    ):
        self.label = label
        self.value = value
        self.color = color
        self.direction = direction

    @classmethod
    def to_proto(self, resource):
        if not resource:
            return None

        res = (
            dashboard_pb2.MonitoringBetaDashboardRowLayoutRowsWidgetsXyChartThresholds()
        )
        if Primitive.to_proto(resource.label):
            res.label = Primitive.to_proto(resource.label)
        if Primitive.to_proto(resource.value):
            res.value = Primitive.to_proto(resource.value)
        if DashboardRowLayoutRowsWidgetsXyChartThresholdsColorEnum.to_proto(
            resource.color
        ):
            res.color = (
                DashboardRowLayoutRowsWidgetsXyChartThresholdsColorEnum.to_proto(
                    resource.color
                )
            )
        if DashboardRowLayoutRowsWidgetsXyChartThresholdsDirectionEnum.to_proto(
            resource.direction
        ):
            res.direction = (
                DashboardRowLayoutRowsWidgetsXyChartThresholdsDirectionEnum.to_proto(
                    resource.direction
                )
            )
        return res

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return None

        return DashboardRowLayoutRowsWidgetsXyChartThresholds(
            label=Primitive.from_proto(resource.label),
            value=Primitive.from_proto(resource.value),
            color=DashboardRowLayoutRowsWidgetsXyChartThresholdsColorEnum.from_proto(
                resource.color
            ),
            direction=DashboardRowLayoutRowsWidgetsXyChartThresholdsDirectionEnum.from_proto(
                resource.direction
            ),
        )


class DashboardRowLayoutRowsWidgetsXyChartThresholdsArray(object):
    @classmethod
    def to_proto(self, resources):
        if not resources:
            return resources
        return [
            DashboardRowLayoutRowsWidgetsXyChartThresholds.to_proto(i)
            for i in resources
        ]

    @classmethod
    def from_proto(self, resources):
        return [
            DashboardRowLayoutRowsWidgetsXyChartThresholds.from_proto(i)
            for i in resources
        ]


class DashboardRowLayoutRowsWidgetsXyChartXAxis(object):
    def __init__(self, label: str = None, scale: str = None):
        self.label = label
        self.scale = scale

    @classmethod
    def to_proto(self, resource):
        if not resource:
            return None

        res = dashboard_pb2.MonitoringBetaDashboardRowLayoutRowsWidgetsXyChartXAxis()
        if Primitive.to_proto(resource.label):
            res.label = Primitive.to_proto(resource.label)
        if DashboardRowLayoutRowsWidgetsXyChartXAxisScaleEnum.to_proto(resource.scale):
            res.scale = DashboardRowLayoutRowsWidgetsXyChartXAxisScaleEnum.to_proto(
                resource.scale
            )
        return res

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return None

        return DashboardRowLayoutRowsWidgetsXyChartXAxis(
            label=Primitive.from_proto(resource.label),
            scale=DashboardRowLayoutRowsWidgetsXyChartXAxisScaleEnum.from_proto(
                resource.scale
            ),
        )


class DashboardRowLayoutRowsWidgetsXyChartXAxisArray(object):
    @classmethod
    def to_proto(self, resources):
        if not resources:
            return resources
        return [
            DashboardRowLayoutRowsWidgetsXyChartXAxis.to_proto(i) for i in resources
        ]

    @classmethod
    def from_proto(self, resources):
        return [
            DashboardRowLayoutRowsWidgetsXyChartXAxis.from_proto(i) for i in resources
        ]


class DashboardRowLayoutRowsWidgetsXyChartYAxis(object):
    def __init__(self, label: str = None, scale: str = None):
        self.label = label
        self.scale = scale

    @classmethod
    def to_proto(self, resource):
        if not resource:
            return None

        res = dashboard_pb2.MonitoringBetaDashboardRowLayoutRowsWidgetsXyChartYAxis()
        if Primitive.to_proto(resource.label):
            res.label = Primitive.to_proto(resource.label)
        if DashboardRowLayoutRowsWidgetsXyChartYAxisScaleEnum.to_proto(resource.scale):
            res.scale = DashboardRowLayoutRowsWidgetsXyChartYAxisScaleEnum.to_proto(
                resource.scale
            )
        return res

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return None

        return DashboardRowLayoutRowsWidgetsXyChartYAxis(
            label=Primitive.from_proto(resource.label),
            scale=DashboardRowLayoutRowsWidgetsXyChartYAxisScaleEnum.from_proto(
                resource.scale
            ),
        )


class DashboardRowLayoutRowsWidgetsXyChartYAxisArray(object):
    @classmethod
    def to_proto(self, resources):
        if not resources:
            return resources
        return [
            DashboardRowLayoutRowsWidgetsXyChartYAxis.to_proto(i) for i in resources
        ]

    @classmethod
    def from_proto(self, resources):
        return [
            DashboardRowLayoutRowsWidgetsXyChartYAxis.from_proto(i) for i in resources
        ]


class DashboardRowLayoutRowsWidgetsXyChartChartOptions(object):
    def __init__(self, mode: str = None):
        self.mode = mode

    @classmethod
    def to_proto(self, resource):
        if not resource:
            return None

        res = (
            dashboard_pb2.MonitoringBetaDashboardRowLayoutRowsWidgetsXyChartChartOptions()
        )
        if DashboardRowLayoutRowsWidgetsXyChartChartOptionsModeEnum.to_proto(
            resource.mode
        ):
            res.mode = (
                DashboardRowLayoutRowsWidgetsXyChartChartOptionsModeEnum.to_proto(
                    resource.mode
                )
            )
        return res

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return None

        return DashboardRowLayoutRowsWidgetsXyChartChartOptions(
            mode=DashboardRowLayoutRowsWidgetsXyChartChartOptionsModeEnum.from_proto(
                resource.mode
            ),
        )


class DashboardRowLayoutRowsWidgetsXyChartChartOptionsArray(object):
    @classmethod
    def to_proto(self, resources):
        if not resources:
            return resources
        return [
            DashboardRowLayoutRowsWidgetsXyChartChartOptions.to_proto(i)
            for i in resources
        ]

    @classmethod
    def from_proto(self, resources):
        return [
            DashboardRowLayoutRowsWidgetsXyChartChartOptions.from_proto(i)
            for i in resources
        ]


class DashboardRowLayoutRowsWidgetsScorecard(object):
    def __init__(
        self,
        time_series_query: dict = None,
        gauge_view: dict = None,
        spark_chart_view: dict = None,
        thresholds: list = None,
    ):
        self.time_series_query = time_series_query
        self.gauge_view = gauge_view
        self.spark_chart_view = spark_chart_view
        self.thresholds = thresholds

    @classmethod
    def to_proto(self, resource):
        if not resource:
            return None

        res = dashboard_pb2.MonitoringBetaDashboardRowLayoutRowsWidgetsScorecard()
        if DashboardRowLayoutRowsWidgetsScorecardTimeSeriesQuery.to_proto(
            resource.time_series_query
        ):
            res.time_series_query.CopyFrom(
                DashboardRowLayoutRowsWidgetsScorecardTimeSeriesQuery.to_proto(
                    resource.time_series_query
                )
            )
        else:
            res.ClearField("time_series_query")
        if DashboardRowLayoutRowsWidgetsScorecardGaugeView.to_proto(
            resource.gauge_view
        ):
            res.gauge_view.CopyFrom(
                DashboardRowLayoutRowsWidgetsScorecardGaugeView.to_proto(
                    resource.gauge_view
                )
            )
        else:
            res.ClearField("gauge_view")
        if DashboardRowLayoutRowsWidgetsScorecardSparkChartView.to_proto(
            resource.spark_chart_view
        ):
            res.spark_chart_view.CopyFrom(
                DashboardRowLayoutRowsWidgetsScorecardSparkChartView.to_proto(
                    resource.spark_chart_view
                )
            )
        else:
            res.ClearField("spark_chart_view")
        if DashboardRowLayoutRowsWidgetsScorecardThresholdsArray.to_proto(
            resource.thresholds
        ):
            res.thresholds.extend(
                DashboardRowLayoutRowsWidgetsScorecardThresholdsArray.to_proto(
                    resource.thresholds
                )
            )
        return res

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return None

        return DashboardRowLayoutRowsWidgetsScorecard(
            time_series_query=DashboardRowLayoutRowsWidgetsScorecardTimeSeriesQuery.from_proto(
                resource.time_series_query
            ),
            gauge_view=DashboardRowLayoutRowsWidgetsScorecardGaugeView.from_proto(
                resource.gauge_view
            ),
            spark_chart_view=DashboardRowLayoutRowsWidgetsScorecardSparkChartView.from_proto(
                resource.spark_chart_view
            ),
            thresholds=DashboardRowLayoutRowsWidgetsScorecardThresholdsArray.from_proto(
                resource.thresholds
            ),
        )


class DashboardRowLayoutRowsWidgetsScorecardArray(object):
    @classmethod
    def to_proto(self, resources):
        if not resources:
            return resources
        return [DashboardRowLayoutRowsWidgetsScorecard.to_proto(i) for i in resources]

    @classmethod
    def from_proto(self, resources):
        return [DashboardRowLayoutRowsWidgetsScorecard.from_proto(i) for i in resources]


class DashboardRowLayoutRowsWidgetsScorecardTimeSeriesQuery(object):
    def __init__(
        self,
        time_series_filter: dict = None,
        time_series_filter_ratio: dict = None,
        time_series_query_language: str = None,
        unit_override: str = None,
    ):
        self.time_series_filter = time_series_filter
        self.time_series_filter_ratio = time_series_filter_ratio
        self.time_series_query_language = time_series_query_language
        self.unit_override = unit_override

    @classmethod
    def to_proto(self, resource):
        if not resource:
            return None

        res = (
            dashboard_pb2.MonitoringBetaDashboardRowLayoutRowsWidgetsScorecardTimeSeriesQuery()
        )
        if DashboardRowLayoutRowsWidgetsScorecardTimeSeriesQueryTimeSeriesFilter.to_proto(
            resource.time_series_filter
        ):
            res.time_series_filter.CopyFrom(
                DashboardRowLayoutRowsWidgetsScorecardTimeSeriesQueryTimeSeriesFilter.to_proto(
                    resource.time_series_filter
                )
            )
        else:
            res.ClearField("time_series_filter")
        if DashboardRowLayoutRowsWidgetsScorecardTimeSeriesQueryTimeSeriesFilterRatio.to_proto(
            resource.time_series_filter_ratio
        ):
            res.time_series_filter_ratio.CopyFrom(
                DashboardRowLayoutRowsWidgetsScorecardTimeSeriesQueryTimeSeriesFilterRatio.to_proto(
                    resource.time_series_filter_ratio
                )
            )
        else:
            res.ClearField("time_series_filter_ratio")
        if Primitive.to_proto(resource.time_series_query_language):
            res.time_series_query_language = Primitive.to_proto(
                resource.time_series_query_language
            )
        if Primitive.to_proto(resource.unit_override):
            res.unit_override = Primitive.to_proto(resource.unit_override)
        return res

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return None

        return DashboardRowLayoutRowsWidgetsScorecardTimeSeriesQuery(
            time_series_filter=DashboardRowLayoutRowsWidgetsScorecardTimeSeriesQueryTimeSeriesFilter.from_proto(
                resource.time_series_filter
            ),
            time_series_filter_ratio=DashboardRowLayoutRowsWidgetsScorecardTimeSeriesQueryTimeSeriesFilterRatio.from_proto(
                resource.time_series_filter_ratio
            ),
            time_series_query_language=Primitive.from_proto(
                resource.time_series_query_language
            ),
            unit_override=Primitive.from_proto(resource.unit_override),
        )


class DashboardRowLayoutRowsWidgetsScorecardTimeSeriesQueryArray(object):
    @classmethod
    def to_proto(self, resources):
        if not resources:
            return resources
        return [
            DashboardRowLayoutRowsWidgetsScorecardTimeSeriesQuery.to_proto(i)
            for i in resources
        ]

    @classmethod
    def from_proto(self, resources):
        return [
            DashboardRowLayoutRowsWidgetsScorecardTimeSeriesQuery.from_proto(i)
            for i in resources
        ]


class DashboardRowLayoutRowsWidgetsScorecardTimeSeriesQueryTimeSeriesFilter(object):
    def __init__(
        self,
        filter: str = None,
        aggregation: dict = None,
        secondary_aggregation: dict = None,
        pick_time_series_filter: dict = None,
    ):
        self.filter = filter
        self.aggregation = aggregation
        self.secondary_aggregation = secondary_aggregation
        self.pick_time_series_filter = pick_time_series_filter

    @classmethod
    def to_proto(self, resource):
        if not resource:
            return None

        res = (
            dashboard_pb2.MonitoringBetaDashboardRowLayoutRowsWidgetsScorecardTimeSeriesQueryTimeSeriesFilter()
        )
        if Primitive.to_proto(resource.filter):
            res.filter = Primitive.to_proto(resource.filter)
        if DashboardRowLayoutRowsWidgetsScorecardTimeSeriesQueryTimeSeriesFilterAggregation.to_proto(
            resource.aggregation
        ):
            res.aggregation.CopyFrom(
                DashboardRowLayoutRowsWidgetsScorecardTimeSeriesQueryTimeSeriesFilterAggregation.to_proto(
                    resource.aggregation
                )
            )
        else:
            res.ClearField("aggregation")
        if DashboardRowLayoutRowsWidgetsScorecardTimeSeriesQueryTimeSeriesFilterSecondaryAggregation.to_proto(
            resource.secondary_aggregation
        ):
            res.secondary_aggregation.CopyFrom(
                DashboardRowLayoutRowsWidgetsScorecardTimeSeriesQueryTimeSeriesFilterSecondaryAggregation.to_proto(
                    resource.secondary_aggregation
                )
            )
        else:
            res.ClearField("secondary_aggregation")
        if DashboardRowLayoutRowsWidgetsScorecardTimeSeriesQueryTimeSeriesFilterPickTimeSeriesFilter.to_proto(
            resource.pick_time_series_filter
        ):
            res.pick_time_series_filter.CopyFrom(
                DashboardRowLayoutRowsWidgetsScorecardTimeSeriesQueryTimeSeriesFilterPickTimeSeriesFilter.to_proto(
                    resource.pick_time_series_filter
                )
            )
        else:
            res.ClearField("pick_time_series_filter")
        return res

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return None

        return DashboardRowLayoutRowsWidgetsScorecardTimeSeriesQueryTimeSeriesFilter(
            filter=Primitive.from_proto(resource.filter),
            aggregation=DashboardRowLayoutRowsWidgetsScorecardTimeSeriesQueryTimeSeriesFilterAggregation.from_proto(
                resource.aggregation
            ),
            secondary_aggregation=DashboardRowLayoutRowsWidgetsScorecardTimeSeriesQueryTimeSeriesFilterSecondaryAggregation.from_proto(
                resource.secondary_aggregation
            ),
            pick_time_series_filter=DashboardRowLayoutRowsWidgetsScorecardTimeSeriesQueryTimeSeriesFilterPickTimeSeriesFilter.from_proto(
                resource.pick_time_series_filter
            ),
        )


class DashboardRowLayoutRowsWidgetsScorecardTimeSeriesQueryTimeSeriesFilterArray(
    object
):
    @classmethod
    def to_proto(self, resources):
        if not resources:
            return resources
        return [
            DashboardRowLayoutRowsWidgetsScorecardTimeSeriesQueryTimeSeriesFilter.to_proto(
                i
            )
            for i in resources
        ]

    @classmethod
    def from_proto(self, resources):
        return [
            DashboardRowLayoutRowsWidgetsScorecardTimeSeriesQueryTimeSeriesFilter.from_proto(
                i
            )
            for i in resources
        ]


class DashboardRowLayoutRowsWidgetsScorecardTimeSeriesQueryTimeSeriesFilterAggregation(
    object
):
    def __init__(
        self,
        alignment_period: str = None,
        per_series_aligner: str = None,
        cross_series_reducer: str = None,
        group_by_fields: list = None,
    ):
        self.alignment_period = alignment_period
        self.per_series_aligner = per_series_aligner
        self.cross_series_reducer = cross_series_reducer
        self.group_by_fields = group_by_fields

    @classmethod
    def to_proto(self, resource):
        if not resource:
            return None

        res = (
            dashboard_pb2.MonitoringBetaDashboardRowLayoutRowsWidgetsScorecardTimeSeriesQueryTimeSeriesFilterAggregation()
        )
        if Primitive.to_proto(resource.alignment_period):
            res.alignment_period = Primitive.to_proto(resource.alignment_period)
        if DashboardRowLayoutRowsWidgetsScorecardTimeSeriesQueryTimeSeriesFilterAggregationPerSeriesAlignerEnum.to_proto(
            resource.per_series_aligner
        ):
            res.per_series_aligner = DashboardRowLayoutRowsWidgetsScorecardTimeSeriesQueryTimeSeriesFilterAggregationPerSeriesAlignerEnum.to_proto(
                resource.per_series_aligner
            )
        if DashboardRowLayoutRowsWidgetsScorecardTimeSeriesQueryTimeSeriesFilterAggregationCrossSeriesReducerEnum.to_proto(
            resource.cross_series_reducer
        ):
            res.cross_series_reducer = DashboardRowLayoutRowsWidgetsScorecardTimeSeriesQueryTimeSeriesFilterAggregationCrossSeriesReducerEnum.to_proto(
                resource.cross_series_reducer
            )
        if Primitive.to_proto(resource.group_by_fields):
            res.group_by_fields.extend(Primitive.to_proto(resource.group_by_fields))
        return res

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return None

        return DashboardRowLayoutRowsWidgetsScorecardTimeSeriesQueryTimeSeriesFilterAggregation(
            alignment_period=Primitive.from_proto(resource.alignment_period),
            per_series_aligner=DashboardRowLayoutRowsWidgetsScorecardTimeSeriesQueryTimeSeriesFilterAggregationPerSeriesAlignerEnum.from_proto(
                resource.per_series_aligner
            ),
            cross_series_reducer=DashboardRowLayoutRowsWidgetsScorecardTimeSeriesQueryTimeSeriesFilterAggregationCrossSeriesReducerEnum.from_proto(
                resource.cross_series_reducer
            ),
            group_by_fields=Primitive.from_proto(resource.group_by_fields),
        )


class DashboardRowLayoutRowsWidgetsScorecardTimeSeriesQueryTimeSeriesFilterAggregationArray(
    object
):
    @classmethod
    def to_proto(self, resources):
        if not resources:
            return resources
        return [
            DashboardRowLayoutRowsWidgetsScorecardTimeSeriesQueryTimeSeriesFilterAggregation.to_proto(
                i
            )
            for i in resources
        ]

    @classmethod
    def from_proto(self, resources):
        return [
            DashboardRowLayoutRowsWidgetsScorecardTimeSeriesQueryTimeSeriesFilterAggregation.from_proto(
                i
            )
            for i in resources
        ]


class DashboardRowLayoutRowsWidgetsScorecardTimeSeriesQueryTimeSeriesFilterSecondaryAggregation(
    object
):
    def __init__(
        self,
        alignment_period: str = None,
        per_series_aligner: str = None,
        cross_series_reducer: str = None,
        group_by_fields: list = None,
    ):
        self.alignment_period = alignment_period
        self.per_series_aligner = per_series_aligner
        self.cross_series_reducer = cross_series_reducer
        self.group_by_fields = group_by_fields

    @classmethod
    def to_proto(self, resource):
        if not resource:
            return None

        res = (
            dashboard_pb2.MonitoringBetaDashboardRowLayoutRowsWidgetsScorecardTimeSeriesQueryTimeSeriesFilterSecondaryAggregation()
        )
        if Primitive.to_proto(resource.alignment_period):
            res.alignment_period = Primitive.to_proto(resource.alignment_period)
        if DashboardRowLayoutRowsWidgetsScorecardTimeSeriesQueryTimeSeriesFilterSecondaryAggregationPerSeriesAlignerEnum.to_proto(
            resource.per_series_aligner
        ):
            res.per_series_aligner = DashboardRowLayoutRowsWidgetsScorecardTimeSeriesQueryTimeSeriesFilterSecondaryAggregationPerSeriesAlignerEnum.to_proto(
                resource.per_series_aligner
            )
        if DashboardRowLayoutRowsWidgetsScorecardTimeSeriesQueryTimeSeriesFilterSecondaryAggregationCrossSeriesReducerEnum.to_proto(
            resource.cross_series_reducer
        ):
            res.cross_series_reducer = DashboardRowLayoutRowsWidgetsScorecardTimeSeriesQueryTimeSeriesFilterSecondaryAggregationCrossSeriesReducerEnum.to_proto(
                resource.cross_series_reducer
            )
        if Primitive.to_proto(resource.group_by_fields):
            res.group_by_fields.extend(Primitive.to_proto(resource.group_by_fields))
        return res

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return None

        return DashboardRowLayoutRowsWidgetsScorecardTimeSeriesQueryTimeSeriesFilterSecondaryAggregation(
            alignment_period=Primitive.from_proto(resource.alignment_period),
            per_series_aligner=DashboardRowLayoutRowsWidgetsScorecardTimeSeriesQueryTimeSeriesFilterSecondaryAggregationPerSeriesAlignerEnum.from_proto(
                resource.per_series_aligner
            ),
            cross_series_reducer=DashboardRowLayoutRowsWidgetsScorecardTimeSeriesQueryTimeSeriesFilterSecondaryAggregationCrossSeriesReducerEnum.from_proto(
                resource.cross_series_reducer
            ),
            group_by_fields=Primitive.from_proto(resource.group_by_fields),
        )


class DashboardRowLayoutRowsWidgetsScorecardTimeSeriesQueryTimeSeriesFilterSecondaryAggregationArray(
    object
):
    @classmethod
    def to_proto(self, resources):
        if not resources:
            return resources
        return [
            DashboardRowLayoutRowsWidgetsScorecardTimeSeriesQueryTimeSeriesFilterSecondaryAggregation.to_proto(
                i
            )
            for i in resources
        ]

    @classmethod
    def from_proto(self, resources):
        return [
            DashboardRowLayoutRowsWidgetsScorecardTimeSeriesQueryTimeSeriesFilterSecondaryAggregation.from_proto(
                i
            )
            for i in resources
        ]


class DashboardRowLayoutRowsWidgetsScorecardTimeSeriesQueryTimeSeriesFilterPickTimeSeriesFilter(
    object
):
    def __init__(
        self,
        ranking_method: str = None,
        num_time_series: int = None,
        direction: str = None,
    ):
        self.ranking_method = ranking_method
        self.num_time_series = num_time_series
        self.direction = direction

    @classmethod
    def to_proto(self, resource):
        if not resource:
            return None

        res = (
            dashboard_pb2.MonitoringBetaDashboardRowLayoutRowsWidgetsScorecardTimeSeriesQueryTimeSeriesFilterPickTimeSeriesFilter()
        )
        if DashboardRowLayoutRowsWidgetsScorecardTimeSeriesQueryTimeSeriesFilterPickTimeSeriesFilterRankingMethodEnum.to_proto(
            resource.ranking_method
        ):
            res.ranking_method = DashboardRowLayoutRowsWidgetsScorecardTimeSeriesQueryTimeSeriesFilterPickTimeSeriesFilterRankingMethodEnum.to_proto(
                resource.ranking_method
            )
        if Primitive.to_proto(resource.num_time_series):
            res.num_time_series = Primitive.to_proto(resource.num_time_series)
        if DashboardRowLayoutRowsWidgetsScorecardTimeSeriesQueryTimeSeriesFilterPickTimeSeriesFilterDirectionEnum.to_proto(
            resource.direction
        ):
            res.direction = DashboardRowLayoutRowsWidgetsScorecardTimeSeriesQueryTimeSeriesFilterPickTimeSeriesFilterDirectionEnum.to_proto(
                resource.direction
            )
        return res

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return None

        return DashboardRowLayoutRowsWidgetsScorecardTimeSeriesQueryTimeSeriesFilterPickTimeSeriesFilter(
            ranking_method=DashboardRowLayoutRowsWidgetsScorecardTimeSeriesQueryTimeSeriesFilterPickTimeSeriesFilterRankingMethodEnum.from_proto(
                resource.ranking_method
            ),
            num_time_series=Primitive.from_proto(resource.num_time_series),
            direction=DashboardRowLayoutRowsWidgetsScorecardTimeSeriesQueryTimeSeriesFilterPickTimeSeriesFilterDirectionEnum.from_proto(
                resource.direction
            ),
        )


class DashboardRowLayoutRowsWidgetsScorecardTimeSeriesQueryTimeSeriesFilterPickTimeSeriesFilterArray(
    object
):
    @classmethod
    def to_proto(self, resources):
        if not resources:
            return resources
        return [
            DashboardRowLayoutRowsWidgetsScorecardTimeSeriesQueryTimeSeriesFilterPickTimeSeriesFilter.to_proto(
                i
            )
            for i in resources
        ]

    @classmethod
    def from_proto(self, resources):
        return [
            DashboardRowLayoutRowsWidgetsScorecardTimeSeriesQueryTimeSeriesFilterPickTimeSeriesFilter.from_proto(
                i
            )
            for i in resources
        ]


class DashboardRowLayoutRowsWidgetsScorecardTimeSeriesQueryTimeSeriesFilterRatio(
    object
):
    def __init__(
        self,
        numerator: dict = None,
        denominator: dict = None,
        secondary_aggregation: dict = None,
        pick_time_series_filter: dict = None,
    ):
        self.numerator = numerator
        self.denominator = denominator
        self.secondary_aggregation = secondary_aggregation
        self.pick_time_series_filter = pick_time_series_filter

    @classmethod
    def to_proto(self, resource):
        if not resource:
            return None

        res = (
            dashboard_pb2.MonitoringBetaDashboardRowLayoutRowsWidgetsScorecardTimeSeriesQueryTimeSeriesFilterRatio()
        )
        if DashboardRowLayoutRowsWidgetsScorecardTimeSeriesQueryTimeSeriesFilterRatioNumerator.to_proto(
            resource.numerator
        ):
            res.numerator.CopyFrom(
                DashboardRowLayoutRowsWidgetsScorecardTimeSeriesQueryTimeSeriesFilterRatioNumerator.to_proto(
                    resource.numerator
                )
            )
        else:
            res.ClearField("numerator")
        if DashboardRowLayoutRowsWidgetsScorecardTimeSeriesQueryTimeSeriesFilterRatioDenominator.to_proto(
            resource.denominator
        ):
            res.denominator.CopyFrom(
                DashboardRowLayoutRowsWidgetsScorecardTimeSeriesQueryTimeSeriesFilterRatioDenominator.to_proto(
                    resource.denominator
                )
            )
        else:
            res.ClearField("denominator")
        if DashboardRowLayoutRowsWidgetsScorecardTimeSeriesQueryTimeSeriesFilterRatioSecondaryAggregation.to_proto(
            resource.secondary_aggregation
        ):
            res.secondary_aggregation.CopyFrom(
                DashboardRowLayoutRowsWidgetsScorecardTimeSeriesQueryTimeSeriesFilterRatioSecondaryAggregation.to_proto(
                    resource.secondary_aggregation
                )
            )
        else:
            res.ClearField("secondary_aggregation")
        if DashboardRowLayoutRowsWidgetsScorecardTimeSeriesQueryTimeSeriesFilterRatioPickTimeSeriesFilter.to_proto(
            resource.pick_time_series_filter
        ):
            res.pick_time_series_filter.CopyFrom(
                DashboardRowLayoutRowsWidgetsScorecardTimeSeriesQueryTimeSeriesFilterRatioPickTimeSeriesFilter.to_proto(
                    resource.pick_time_series_filter
                )
            )
        else:
            res.ClearField("pick_time_series_filter")
        return res

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return None

        return DashboardRowLayoutRowsWidgetsScorecardTimeSeriesQueryTimeSeriesFilterRatio(
            numerator=DashboardRowLayoutRowsWidgetsScorecardTimeSeriesQueryTimeSeriesFilterRatioNumerator.from_proto(
                resource.numerator
            ),
            denominator=DashboardRowLayoutRowsWidgetsScorecardTimeSeriesQueryTimeSeriesFilterRatioDenominator.from_proto(
                resource.denominator
            ),
            secondary_aggregation=DashboardRowLayoutRowsWidgetsScorecardTimeSeriesQueryTimeSeriesFilterRatioSecondaryAggregation.from_proto(
                resource.secondary_aggregation
            ),
            pick_time_series_filter=DashboardRowLayoutRowsWidgetsScorecardTimeSeriesQueryTimeSeriesFilterRatioPickTimeSeriesFilter.from_proto(
                resource.pick_time_series_filter
            ),
        )


class DashboardRowLayoutRowsWidgetsScorecardTimeSeriesQueryTimeSeriesFilterRatioArray(
    object
):
    @classmethod
    def to_proto(self, resources):
        if not resources:
            return resources
        return [
            DashboardRowLayoutRowsWidgetsScorecardTimeSeriesQueryTimeSeriesFilterRatio.to_proto(
                i
            )
            for i in resources
        ]

    @classmethod
    def from_proto(self, resources):
        return [
            DashboardRowLayoutRowsWidgetsScorecardTimeSeriesQueryTimeSeriesFilterRatio.from_proto(
                i
            )
            for i in resources
        ]


class DashboardRowLayoutRowsWidgetsScorecardTimeSeriesQueryTimeSeriesFilterRatioNumerator(
    object
):
    def __init__(self, filter: str = None, aggregation: dict = None):
        self.filter = filter
        self.aggregation = aggregation

    @classmethod
    def to_proto(self, resource):
        if not resource:
            return None

        res = (
            dashboard_pb2.MonitoringBetaDashboardRowLayoutRowsWidgetsScorecardTimeSeriesQueryTimeSeriesFilterRatioNumerator()
        )
        if Primitive.to_proto(resource.filter):
            res.filter = Primitive.to_proto(resource.filter)
        if DashboardRowLayoutRowsWidgetsScorecardTimeSeriesQueryTimeSeriesFilterRatioNumeratorAggregation.to_proto(
            resource.aggregation
        ):
            res.aggregation.CopyFrom(
                DashboardRowLayoutRowsWidgetsScorecardTimeSeriesQueryTimeSeriesFilterRatioNumeratorAggregation.to_proto(
                    resource.aggregation
                )
            )
        else:
            res.ClearField("aggregation")
        return res

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return None

        return DashboardRowLayoutRowsWidgetsScorecardTimeSeriesQueryTimeSeriesFilterRatioNumerator(
            filter=Primitive.from_proto(resource.filter),
            aggregation=DashboardRowLayoutRowsWidgetsScorecardTimeSeriesQueryTimeSeriesFilterRatioNumeratorAggregation.from_proto(
                resource.aggregation
            ),
        )


class DashboardRowLayoutRowsWidgetsScorecardTimeSeriesQueryTimeSeriesFilterRatioNumeratorArray(
    object
):
    @classmethod
    def to_proto(self, resources):
        if not resources:
            return resources
        return [
            DashboardRowLayoutRowsWidgetsScorecardTimeSeriesQueryTimeSeriesFilterRatioNumerator.to_proto(
                i
            )
            for i in resources
        ]

    @classmethod
    def from_proto(self, resources):
        return [
            DashboardRowLayoutRowsWidgetsScorecardTimeSeriesQueryTimeSeriesFilterRatioNumerator.from_proto(
                i
            )
            for i in resources
        ]


class DashboardRowLayoutRowsWidgetsScorecardTimeSeriesQueryTimeSeriesFilterRatioNumeratorAggregation(
    object
):
    def __init__(
        self,
        alignment_period: str = None,
        per_series_aligner: str = None,
        cross_series_reducer: str = None,
        group_by_fields: list = None,
    ):
        self.alignment_period = alignment_period
        self.per_series_aligner = per_series_aligner
        self.cross_series_reducer = cross_series_reducer
        self.group_by_fields = group_by_fields

    @classmethod
    def to_proto(self, resource):
        if not resource:
            return None

        res = (
            dashboard_pb2.MonitoringBetaDashboardRowLayoutRowsWidgetsScorecardTimeSeriesQueryTimeSeriesFilterRatioNumeratorAggregation()
        )
        if Primitive.to_proto(resource.alignment_period):
            res.alignment_period = Primitive.to_proto(resource.alignment_period)
        if DashboardRowLayoutRowsWidgetsScorecardTimeSeriesQueryTimeSeriesFilterRatioNumeratorAggregationPerSeriesAlignerEnum.to_proto(
            resource.per_series_aligner
        ):
            res.per_series_aligner = DashboardRowLayoutRowsWidgetsScorecardTimeSeriesQueryTimeSeriesFilterRatioNumeratorAggregationPerSeriesAlignerEnum.to_proto(
                resource.per_series_aligner
            )
        if DashboardRowLayoutRowsWidgetsScorecardTimeSeriesQueryTimeSeriesFilterRatioNumeratorAggregationCrossSeriesReducerEnum.to_proto(
            resource.cross_series_reducer
        ):
            res.cross_series_reducer = DashboardRowLayoutRowsWidgetsScorecardTimeSeriesQueryTimeSeriesFilterRatioNumeratorAggregationCrossSeriesReducerEnum.to_proto(
                resource.cross_series_reducer
            )
        if Primitive.to_proto(resource.group_by_fields):
            res.group_by_fields.extend(Primitive.to_proto(resource.group_by_fields))
        return res

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return None

        return DashboardRowLayoutRowsWidgetsScorecardTimeSeriesQueryTimeSeriesFilterRatioNumeratorAggregation(
            alignment_period=Primitive.from_proto(resource.alignment_period),
            per_series_aligner=DashboardRowLayoutRowsWidgetsScorecardTimeSeriesQueryTimeSeriesFilterRatioNumeratorAggregationPerSeriesAlignerEnum.from_proto(
                resource.per_series_aligner
            ),
            cross_series_reducer=DashboardRowLayoutRowsWidgetsScorecardTimeSeriesQueryTimeSeriesFilterRatioNumeratorAggregationCrossSeriesReducerEnum.from_proto(
                resource.cross_series_reducer
            ),
            group_by_fields=Primitive.from_proto(resource.group_by_fields),
        )


class DashboardRowLayoutRowsWidgetsScorecardTimeSeriesQueryTimeSeriesFilterRatioNumeratorAggregationArray(
    object
):
    @classmethod
    def to_proto(self, resources):
        if not resources:
            return resources
        return [
            DashboardRowLayoutRowsWidgetsScorecardTimeSeriesQueryTimeSeriesFilterRatioNumeratorAggregation.to_proto(
                i
            )
            for i in resources
        ]

    @classmethod
    def from_proto(self, resources):
        return [
            DashboardRowLayoutRowsWidgetsScorecardTimeSeriesQueryTimeSeriesFilterRatioNumeratorAggregation.from_proto(
                i
            )
            for i in resources
        ]


class DashboardRowLayoutRowsWidgetsScorecardTimeSeriesQueryTimeSeriesFilterRatioDenominator(
    object
):
    def __init__(self, filter: str = None, aggregation: dict = None):
        self.filter = filter
        self.aggregation = aggregation

    @classmethod
    def to_proto(self, resource):
        if not resource:
            return None

        res = (
            dashboard_pb2.MonitoringBetaDashboardRowLayoutRowsWidgetsScorecardTimeSeriesQueryTimeSeriesFilterRatioDenominator()
        )
        if Primitive.to_proto(resource.filter):
            res.filter = Primitive.to_proto(resource.filter)
        if DashboardRowLayoutRowsWidgetsScorecardTimeSeriesQueryTimeSeriesFilterRatioDenominatorAggregation.to_proto(
            resource.aggregation
        ):
            res.aggregation.CopyFrom(
                DashboardRowLayoutRowsWidgetsScorecardTimeSeriesQueryTimeSeriesFilterRatioDenominatorAggregation.to_proto(
                    resource.aggregation
                )
            )
        else:
            res.ClearField("aggregation")
        return res

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return None

        return DashboardRowLayoutRowsWidgetsScorecardTimeSeriesQueryTimeSeriesFilterRatioDenominator(
            filter=Primitive.from_proto(resource.filter),
            aggregation=DashboardRowLayoutRowsWidgetsScorecardTimeSeriesQueryTimeSeriesFilterRatioDenominatorAggregation.from_proto(
                resource.aggregation
            ),
        )


class DashboardRowLayoutRowsWidgetsScorecardTimeSeriesQueryTimeSeriesFilterRatioDenominatorArray(
    object
):
    @classmethod
    def to_proto(self, resources):
        if not resources:
            return resources
        return [
            DashboardRowLayoutRowsWidgetsScorecardTimeSeriesQueryTimeSeriesFilterRatioDenominator.to_proto(
                i
            )
            for i in resources
        ]

    @classmethod
    def from_proto(self, resources):
        return [
            DashboardRowLayoutRowsWidgetsScorecardTimeSeriesQueryTimeSeriesFilterRatioDenominator.from_proto(
                i
            )
            for i in resources
        ]


class DashboardRowLayoutRowsWidgetsScorecardTimeSeriesQueryTimeSeriesFilterRatioDenominatorAggregation(
    object
):
    def __init__(
        self,
        alignment_period: str = None,
        per_series_aligner: str = None,
        cross_series_reducer: str = None,
        group_by_fields: list = None,
    ):
        self.alignment_period = alignment_period
        self.per_series_aligner = per_series_aligner
        self.cross_series_reducer = cross_series_reducer
        self.group_by_fields = group_by_fields

    @classmethod
    def to_proto(self, resource):
        if not resource:
            return None

        res = (
            dashboard_pb2.MonitoringBetaDashboardRowLayoutRowsWidgetsScorecardTimeSeriesQueryTimeSeriesFilterRatioDenominatorAggregation()
        )
        if Primitive.to_proto(resource.alignment_period):
            res.alignment_period = Primitive.to_proto(resource.alignment_period)
        if DashboardRowLayoutRowsWidgetsScorecardTimeSeriesQueryTimeSeriesFilterRatioDenominatorAggregationPerSeriesAlignerEnum.to_proto(
            resource.per_series_aligner
        ):
            res.per_series_aligner = DashboardRowLayoutRowsWidgetsScorecardTimeSeriesQueryTimeSeriesFilterRatioDenominatorAggregationPerSeriesAlignerEnum.to_proto(
                resource.per_series_aligner
            )
        if DashboardRowLayoutRowsWidgetsScorecardTimeSeriesQueryTimeSeriesFilterRatioDenominatorAggregationCrossSeriesReducerEnum.to_proto(
            resource.cross_series_reducer
        ):
            res.cross_series_reducer = DashboardRowLayoutRowsWidgetsScorecardTimeSeriesQueryTimeSeriesFilterRatioDenominatorAggregationCrossSeriesReducerEnum.to_proto(
                resource.cross_series_reducer
            )
        if Primitive.to_proto(resource.group_by_fields):
            res.group_by_fields.extend(Primitive.to_proto(resource.group_by_fields))
        return res

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return None

        return DashboardRowLayoutRowsWidgetsScorecardTimeSeriesQueryTimeSeriesFilterRatioDenominatorAggregation(
            alignment_period=Primitive.from_proto(resource.alignment_period),
            per_series_aligner=DashboardRowLayoutRowsWidgetsScorecardTimeSeriesQueryTimeSeriesFilterRatioDenominatorAggregationPerSeriesAlignerEnum.from_proto(
                resource.per_series_aligner
            ),
            cross_series_reducer=DashboardRowLayoutRowsWidgetsScorecardTimeSeriesQueryTimeSeriesFilterRatioDenominatorAggregationCrossSeriesReducerEnum.from_proto(
                resource.cross_series_reducer
            ),
            group_by_fields=Primitive.from_proto(resource.group_by_fields),
        )


class DashboardRowLayoutRowsWidgetsScorecardTimeSeriesQueryTimeSeriesFilterRatioDenominatorAggregationArray(
    object
):
    @classmethod
    def to_proto(self, resources):
        if not resources:
            return resources
        return [
            DashboardRowLayoutRowsWidgetsScorecardTimeSeriesQueryTimeSeriesFilterRatioDenominatorAggregation.to_proto(
                i
            )
            for i in resources
        ]

    @classmethod
    def from_proto(self, resources):
        return [
            DashboardRowLayoutRowsWidgetsScorecardTimeSeriesQueryTimeSeriesFilterRatioDenominatorAggregation.from_proto(
                i
            )
            for i in resources
        ]


class DashboardRowLayoutRowsWidgetsScorecardTimeSeriesQueryTimeSeriesFilterRatioSecondaryAggregation(
    object
):
    def __init__(
        self,
        alignment_period: str = None,
        per_series_aligner: str = None,
        cross_series_reducer: str = None,
        group_by_fields: list = None,
    ):
        self.alignment_period = alignment_period
        self.per_series_aligner = per_series_aligner
        self.cross_series_reducer = cross_series_reducer
        self.group_by_fields = group_by_fields

    @classmethod
    def to_proto(self, resource):
        if not resource:
            return None

        res = (
            dashboard_pb2.MonitoringBetaDashboardRowLayoutRowsWidgetsScorecardTimeSeriesQueryTimeSeriesFilterRatioSecondaryAggregation()
        )
        if Primitive.to_proto(resource.alignment_period):
            res.alignment_period = Primitive.to_proto(resource.alignment_period)
        if DashboardRowLayoutRowsWidgetsScorecardTimeSeriesQueryTimeSeriesFilterRatioSecondaryAggregationPerSeriesAlignerEnum.to_proto(
            resource.per_series_aligner
        ):
            res.per_series_aligner = DashboardRowLayoutRowsWidgetsScorecardTimeSeriesQueryTimeSeriesFilterRatioSecondaryAggregationPerSeriesAlignerEnum.to_proto(
                resource.per_series_aligner
            )
        if DashboardRowLayoutRowsWidgetsScorecardTimeSeriesQueryTimeSeriesFilterRatioSecondaryAggregationCrossSeriesReducerEnum.to_proto(
            resource.cross_series_reducer
        ):
            res.cross_series_reducer = DashboardRowLayoutRowsWidgetsScorecardTimeSeriesQueryTimeSeriesFilterRatioSecondaryAggregationCrossSeriesReducerEnum.to_proto(
                resource.cross_series_reducer
            )
        if Primitive.to_proto(resource.group_by_fields):
            res.group_by_fields.extend(Primitive.to_proto(resource.group_by_fields))
        return res

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return None

        return DashboardRowLayoutRowsWidgetsScorecardTimeSeriesQueryTimeSeriesFilterRatioSecondaryAggregation(
            alignment_period=Primitive.from_proto(resource.alignment_period),
            per_series_aligner=DashboardRowLayoutRowsWidgetsScorecardTimeSeriesQueryTimeSeriesFilterRatioSecondaryAggregationPerSeriesAlignerEnum.from_proto(
                resource.per_series_aligner
            ),
            cross_series_reducer=DashboardRowLayoutRowsWidgetsScorecardTimeSeriesQueryTimeSeriesFilterRatioSecondaryAggregationCrossSeriesReducerEnum.from_proto(
                resource.cross_series_reducer
            ),
            group_by_fields=Primitive.from_proto(resource.group_by_fields),
        )


class DashboardRowLayoutRowsWidgetsScorecardTimeSeriesQueryTimeSeriesFilterRatioSecondaryAggregationArray(
    object
):
    @classmethod
    def to_proto(self, resources):
        if not resources:
            return resources
        return [
            DashboardRowLayoutRowsWidgetsScorecardTimeSeriesQueryTimeSeriesFilterRatioSecondaryAggregation.to_proto(
                i
            )
            for i in resources
        ]

    @classmethod
    def from_proto(self, resources):
        return [
            DashboardRowLayoutRowsWidgetsScorecardTimeSeriesQueryTimeSeriesFilterRatioSecondaryAggregation.from_proto(
                i
            )
            for i in resources
        ]


class DashboardRowLayoutRowsWidgetsScorecardTimeSeriesQueryTimeSeriesFilterRatioPickTimeSeriesFilter(
    object
):
    def __init__(
        self,
        ranking_method: str = None,
        num_time_series: int = None,
        direction: str = None,
    ):
        self.ranking_method = ranking_method
        self.num_time_series = num_time_series
        self.direction = direction

    @classmethod
    def to_proto(self, resource):
        if not resource:
            return None

        res = (
            dashboard_pb2.MonitoringBetaDashboardRowLayoutRowsWidgetsScorecardTimeSeriesQueryTimeSeriesFilterRatioPickTimeSeriesFilter()
        )
        if DashboardRowLayoutRowsWidgetsScorecardTimeSeriesQueryTimeSeriesFilterRatioPickTimeSeriesFilterRankingMethodEnum.to_proto(
            resource.ranking_method
        ):
            res.ranking_method = DashboardRowLayoutRowsWidgetsScorecardTimeSeriesQueryTimeSeriesFilterRatioPickTimeSeriesFilterRankingMethodEnum.to_proto(
                resource.ranking_method
            )
        if Primitive.to_proto(resource.num_time_series):
            res.num_time_series = Primitive.to_proto(resource.num_time_series)
        if DashboardRowLayoutRowsWidgetsScorecardTimeSeriesQueryTimeSeriesFilterRatioPickTimeSeriesFilterDirectionEnum.to_proto(
            resource.direction
        ):
            res.direction = DashboardRowLayoutRowsWidgetsScorecardTimeSeriesQueryTimeSeriesFilterRatioPickTimeSeriesFilterDirectionEnum.to_proto(
                resource.direction
            )
        return res

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return None

        return DashboardRowLayoutRowsWidgetsScorecardTimeSeriesQueryTimeSeriesFilterRatioPickTimeSeriesFilter(
            ranking_method=DashboardRowLayoutRowsWidgetsScorecardTimeSeriesQueryTimeSeriesFilterRatioPickTimeSeriesFilterRankingMethodEnum.from_proto(
                resource.ranking_method
            ),
            num_time_series=Primitive.from_proto(resource.num_time_series),
            direction=DashboardRowLayoutRowsWidgetsScorecardTimeSeriesQueryTimeSeriesFilterRatioPickTimeSeriesFilterDirectionEnum.from_proto(
                resource.direction
            ),
        )


class DashboardRowLayoutRowsWidgetsScorecardTimeSeriesQueryTimeSeriesFilterRatioPickTimeSeriesFilterArray(
    object
):
    @classmethod
    def to_proto(self, resources):
        if not resources:
            return resources
        return [
            DashboardRowLayoutRowsWidgetsScorecardTimeSeriesQueryTimeSeriesFilterRatioPickTimeSeriesFilter.to_proto(
                i
            )
            for i in resources
        ]

    @classmethod
    def from_proto(self, resources):
        return [
            DashboardRowLayoutRowsWidgetsScorecardTimeSeriesQueryTimeSeriesFilterRatioPickTimeSeriesFilter.from_proto(
                i
            )
            for i in resources
        ]


class DashboardRowLayoutRowsWidgetsScorecardGaugeView(object):
    def __init__(self, lower_bound: float = None, upper_bound: float = None):
        self.lower_bound = lower_bound
        self.upper_bound = upper_bound

    @classmethod
    def to_proto(self, resource):
        if not resource:
            return None

        res = (
            dashboard_pb2.MonitoringBetaDashboardRowLayoutRowsWidgetsScorecardGaugeView()
        )
        if Primitive.to_proto(resource.lower_bound):
            res.lower_bound = Primitive.to_proto(resource.lower_bound)
        if Primitive.to_proto(resource.upper_bound):
            res.upper_bound = Primitive.to_proto(resource.upper_bound)
        return res

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return None

        return DashboardRowLayoutRowsWidgetsScorecardGaugeView(
            lower_bound=Primitive.from_proto(resource.lower_bound),
            upper_bound=Primitive.from_proto(resource.upper_bound),
        )


class DashboardRowLayoutRowsWidgetsScorecardGaugeViewArray(object):
    @classmethod
    def to_proto(self, resources):
        if not resources:
            return resources
        return [
            DashboardRowLayoutRowsWidgetsScorecardGaugeView.to_proto(i)
            for i in resources
        ]

    @classmethod
    def from_proto(self, resources):
        return [
            DashboardRowLayoutRowsWidgetsScorecardGaugeView.from_proto(i)
            for i in resources
        ]


class DashboardRowLayoutRowsWidgetsScorecardSparkChartView(object):
    def __init__(self, spark_chart_type: str = None, min_alignment_period: str = None):
        self.spark_chart_type = spark_chart_type
        self.min_alignment_period = min_alignment_period

    @classmethod
    def to_proto(self, resource):
        if not resource:
            return None

        res = (
            dashboard_pb2.MonitoringBetaDashboardRowLayoutRowsWidgetsScorecardSparkChartView()
        )
        if DashboardRowLayoutRowsWidgetsScorecardSparkChartViewSparkChartTypeEnum.to_proto(
            resource.spark_chart_type
        ):
            res.spark_chart_type = DashboardRowLayoutRowsWidgetsScorecardSparkChartViewSparkChartTypeEnum.to_proto(
                resource.spark_chart_type
            )
        if Primitive.to_proto(resource.min_alignment_period):
            res.min_alignment_period = Primitive.to_proto(resource.min_alignment_period)
        return res

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return None

        return DashboardRowLayoutRowsWidgetsScorecardSparkChartView(
            spark_chart_type=DashboardRowLayoutRowsWidgetsScorecardSparkChartViewSparkChartTypeEnum.from_proto(
                resource.spark_chart_type
            ),
            min_alignment_period=Primitive.from_proto(resource.min_alignment_period),
        )


class DashboardRowLayoutRowsWidgetsScorecardSparkChartViewArray(object):
    @classmethod
    def to_proto(self, resources):
        if not resources:
            return resources
        return [
            DashboardRowLayoutRowsWidgetsScorecardSparkChartView.to_proto(i)
            for i in resources
        ]

    @classmethod
    def from_proto(self, resources):
        return [
            DashboardRowLayoutRowsWidgetsScorecardSparkChartView.from_proto(i)
            for i in resources
        ]


class DashboardRowLayoutRowsWidgetsScorecardThresholds(object):
    def __init__(
        self,
        label: str = None,
        value: float = None,
        color: str = None,
        direction: str = None,
    ):
        self.label = label
        self.value = value
        self.color = color
        self.direction = direction

    @classmethod
    def to_proto(self, resource):
        if not resource:
            return None

        res = (
            dashboard_pb2.MonitoringBetaDashboardRowLayoutRowsWidgetsScorecardThresholds()
        )
        if Primitive.to_proto(resource.label):
            res.label = Primitive.to_proto(resource.label)
        if Primitive.to_proto(resource.value):
            res.value = Primitive.to_proto(resource.value)
        if DashboardRowLayoutRowsWidgetsScorecardThresholdsColorEnum.to_proto(
            resource.color
        ):
            res.color = (
                DashboardRowLayoutRowsWidgetsScorecardThresholdsColorEnum.to_proto(
                    resource.color
                )
            )
        if DashboardRowLayoutRowsWidgetsScorecardThresholdsDirectionEnum.to_proto(
            resource.direction
        ):
            res.direction = (
                DashboardRowLayoutRowsWidgetsScorecardThresholdsDirectionEnum.to_proto(
                    resource.direction
                )
            )
        return res

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return None

        return DashboardRowLayoutRowsWidgetsScorecardThresholds(
            label=Primitive.from_proto(resource.label),
            value=Primitive.from_proto(resource.value),
            color=DashboardRowLayoutRowsWidgetsScorecardThresholdsColorEnum.from_proto(
                resource.color
            ),
            direction=DashboardRowLayoutRowsWidgetsScorecardThresholdsDirectionEnum.from_proto(
                resource.direction
            ),
        )


class DashboardRowLayoutRowsWidgetsScorecardThresholdsArray(object):
    @classmethod
    def to_proto(self, resources):
        if not resources:
            return resources
        return [
            DashboardRowLayoutRowsWidgetsScorecardThresholds.to_proto(i)
            for i in resources
        ]

    @classmethod
    def from_proto(self, resources):
        return [
            DashboardRowLayoutRowsWidgetsScorecardThresholds.from_proto(i)
            for i in resources
        ]


class DashboardRowLayoutRowsWidgetsText(object):
    def __init__(self, content: str = None, format: str = None):
        self.content = content
        self.format = format

    @classmethod
    def to_proto(self, resource):
        if not resource:
            return None

        res = dashboard_pb2.MonitoringBetaDashboardRowLayoutRowsWidgetsText()
        if Primitive.to_proto(resource.content):
            res.content = Primitive.to_proto(resource.content)
        if DashboardRowLayoutRowsWidgetsTextFormatEnum.to_proto(resource.format):
            res.format = DashboardRowLayoutRowsWidgetsTextFormatEnum.to_proto(
                resource.format
            )
        return res

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return None

        return DashboardRowLayoutRowsWidgetsText(
            content=Primitive.from_proto(resource.content),
            format=DashboardRowLayoutRowsWidgetsTextFormatEnum.from_proto(
                resource.format
            ),
        )


class DashboardRowLayoutRowsWidgetsTextArray(object):
    @classmethod
    def to_proto(self, resources):
        if not resources:
            return resources
        return [DashboardRowLayoutRowsWidgetsText.to_proto(i) for i in resources]

    @classmethod
    def from_proto(self, resources):
        return [DashboardRowLayoutRowsWidgetsText.from_proto(i) for i in resources]


class DashboardRowLayoutRowsWidgetsBlank(object):
    @classmethod
    def to_proto(self, resource):
        if not resource:
            return None

        res = dashboard_pb2.MonitoringBetaDashboardRowLayoutRowsWidgetsBlank()
        return res

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return None

        return DashboardRowLayoutRowsWidgetsBlank()


class DashboardRowLayoutRowsWidgetsBlankArray(object):
    @classmethod
    def to_proto(self, resources):
        if not resources:
            return resources
        return [DashboardRowLayoutRowsWidgetsBlank.to_proto(i) for i in resources]

    @classmethod
    def from_proto(self, resources):
        return [DashboardRowLayoutRowsWidgetsBlank.from_proto(i) for i in resources]


class DashboardRowLayoutRowsWidgetsLogsPanel(object):
    def __init__(self, filter: str = None, resource_names: list = None):
        self.filter = filter
        self.resource_names = resource_names

    @classmethod
    def to_proto(self, resource):
        if not resource:
            return None

        res = dashboard_pb2.MonitoringBetaDashboardRowLayoutRowsWidgetsLogsPanel()
        if Primitive.to_proto(resource.filter):
            res.filter = Primitive.to_proto(resource.filter)
        if Primitive.to_proto(resource.resource_names):
            res.resource_names.extend(Primitive.to_proto(resource.resource_names))
        return res

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return None

        return DashboardRowLayoutRowsWidgetsLogsPanel(
            filter=Primitive.from_proto(resource.filter),
            resource_names=Primitive.from_proto(resource.resource_names),
        )


class DashboardRowLayoutRowsWidgetsLogsPanelArray(object):
    @classmethod
    def to_proto(self, resources):
        if not resources:
            return resources
        return [DashboardRowLayoutRowsWidgetsLogsPanel.to_proto(i) for i in resources]

    @classmethod
    def from_proto(self, resources):
        return [DashboardRowLayoutRowsWidgetsLogsPanel.from_proto(i) for i in resources]


class DashboardColumnLayout(object):
    def __init__(self, columns: list = None):
        self.columns = columns

    @classmethod
    def to_proto(self, resource):
        if not resource:
            return None

        res = dashboard_pb2.MonitoringBetaDashboardColumnLayout()
        if DashboardColumnLayoutColumnsArray.to_proto(resource.columns):
            res.columns.extend(
                DashboardColumnLayoutColumnsArray.to_proto(resource.columns)
            )
        return res

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return None

        return DashboardColumnLayout(
            columns=DashboardColumnLayoutColumnsArray.from_proto(resource.columns),
        )


class DashboardColumnLayoutArray(object):
    @classmethod
    def to_proto(self, resources):
        if not resources:
            return resources
        return [DashboardColumnLayout.to_proto(i) for i in resources]

    @classmethod
    def from_proto(self, resources):
        return [DashboardColumnLayout.from_proto(i) for i in resources]


class DashboardColumnLayoutColumns(object):
    def __init__(self, weight: int = None, widgets: list = None):
        self.weight = weight
        self.widgets = widgets

    @classmethod
    def to_proto(self, resource):
        if not resource:
            return None

        res = dashboard_pb2.MonitoringBetaDashboardColumnLayoutColumns()
        if Primitive.to_proto(resource.weight):
            res.weight = Primitive.to_proto(resource.weight)
        if DashboardColumnLayoutColumnsWidgetsArray.to_proto(resource.widgets):
            res.widgets.extend(
                DashboardColumnLayoutColumnsWidgetsArray.to_proto(resource.widgets)
            )
        return res

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return None

        return DashboardColumnLayoutColumns(
            weight=Primitive.from_proto(resource.weight),
            widgets=DashboardColumnLayoutColumnsWidgetsArray.from_proto(
                resource.widgets
            ),
        )


class DashboardColumnLayoutColumnsArray(object):
    @classmethod
    def to_proto(self, resources):
        if not resources:
            return resources
        return [DashboardColumnLayoutColumns.to_proto(i) for i in resources]

    @classmethod
    def from_proto(self, resources):
        return [DashboardColumnLayoutColumns.from_proto(i) for i in resources]


class DashboardColumnLayoutColumnsWidgets(object):
    def __init__(
        self,
        title: str = None,
        xy_chart: dict = None,
        scorecard: dict = None,
        text: dict = None,
        blank: dict = None,
        logs_panel: dict = None,
    ):
        self.title = title
        self.xy_chart = xy_chart
        self.scorecard = scorecard
        self.text = text
        self.blank = blank
        self.logs_panel = logs_panel

    @classmethod
    def to_proto(self, resource):
        if not resource:
            return None

        res = dashboard_pb2.MonitoringBetaDashboardColumnLayoutColumnsWidgets()
        if Primitive.to_proto(resource.title):
            res.title = Primitive.to_proto(resource.title)
        if DashboardColumnLayoutColumnsWidgetsXyChart.to_proto(resource.xy_chart):
            res.xy_chart.CopyFrom(
                DashboardColumnLayoutColumnsWidgetsXyChart.to_proto(resource.xy_chart)
            )
        else:
            res.ClearField("xy_chart")
        if DashboardColumnLayoutColumnsWidgetsScorecard.to_proto(resource.scorecard):
            res.scorecard.CopyFrom(
                DashboardColumnLayoutColumnsWidgetsScorecard.to_proto(
                    resource.scorecard
                )
            )
        else:
            res.ClearField("scorecard")
        if DashboardColumnLayoutColumnsWidgetsText.to_proto(resource.text):
            res.text.CopyFrom(
                DashboardColumnLayoutColumnsWidgetsText.to_proto(resource.text)
            )
        else:
            res.ClearField("text")
        if DashboardColumnLayoutColumnsWidgetsBlank.to_proto(resource.blank):
            res.blank.CopyFrom(
                DashboardColumnLayoutColumnsWidgetsBlank.to_proto(resource.blank)
            )
        else:
            res.ClearField("blank")
        if DashboardColumnLayoutColumnsWidgetsLogsPanel.to_proto(resource.logs_panel):
            res.logs_panel.CopyFrom(
                DashboardColumnLayoutColumnsWidgetsLogsPanel.to_proto(
                    resource.logs_panel
                )
            )
        else:
            res.ClearField("logs_panel")
        return res

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return None

        return DashboardColumnLayoutColumnsWidgets(
            title=Primitive.from_proto(resource.title),
            xy_chart=DashboardColumnLayoutColumnsWidgetsXyChart.from_proto(
                resource.xy_chart
            ),
            scorecard=DashboardColumnLayoutColumnsWidgetsScorecard.from_proto(
                resource.scorecard
            ),
            text=DashboardColumnLayoutColumnsWidgetsText.from_proto(resource.text),
            blank=DashboardColumnLayoutColumnsWidgetsBlank.from_proto(resource.blank),
            logs_panel=DashboardColumnLayoutColumnsWidgetsLogsPanel.from_proto(
                resource.logs_panel
            ),
        )


class DashboardColumnLayoutColumnsWidgetsArray(object):
    @classmethod
    def to_proto(self, resources):
        if not resources:
            return resources
        return [DashboardColumnLayoutColumnsWidgets.to_proto(i) for i in resources]

    @classmethod
    def from_proto(self, resources):
        return [DashboardColumnLayoutColumnsWidgets.from_proto(i) for i in resources]


class DashboardColumnLayoutColumnsWidgetsXyChart(object):
    def __init__(
        self,
        data_sets: list = None,
        timeshift_duration: str = None,
        thresholds: list = None,
        x_axis: dict = None,
        y_axis: dict = None,
        chart_options: dict = None,
    ):
        self.data_sets = data_sets
        self.timeshift_duration = timeshift_duration
        self.thresholds = thresholds
        self.x_axis = x_axis
        self.y_axis = y_axis
        self.chart_options = chart_options

    @classmethod
    def to_proto(self, resource):
        if not resource:
            return None

        res = dashboard_pb2.MonitoringBetaDashboardColumnLayoutColumnsWidgetsXyChart()
        if DashboardColumnLayoutColumnsWidgetsXyChartDataSetsArray.to_proto(
            resource.data_sets
        ):
            res.data_sets.extend(
                DashboardColumnLayoutColumnsWidgetsXyChartDataSetsArray.to_proto(
                    resource.data_sets
                )
            )
        if Primitive.to_proto(resource.timeshift_duration):
            res.timeshift_duration = Primitive.to_proto(resource.timeshift_duration)
        if DashboardColumnLayoutColumnsWidgetsXyChartThresholdsArray.to_proto(
            resource.thresholds
        ):
            res.thresholds.extend(
                DashboardColumnLayoutColumnsWidgetsXyChartThresholdsArray.to_proto(
                    resource.thresholds
                )
            )
        if DashboardColumnLayoutColumnsWidgetsXyChartXAxis.to_proto(resource.x_axis):
            res.x_axis.CopyFrom(
                DashboardColumnLayoutColumnsWidgetsXyChartXAxis.to_proto(
                    resource.x_axis
                )
            )
        else:
            res.ClearField("x_axis")
        if DashboardColumnLayoutColumnsWidgetsXyChartYAxis.to_proto(resource.y_axis):
            res.y_axis.CopyFrom(
                DashboardColumnLayoutColumnsWidgetsXyChartYAxis.to_proto(
                    resource.y_axis
                )
            )
        else:
            res.ClearField("y_axis")
        if DashboardColumnLayoutColumnsWidgetsXyChartChartOptions.to_proto(
            resource.chart_options
        ):
            res.chart_options.CopyFrom(
                DashboardColumnLayoutColumnsWidgetsXyChartChartOptions.to_proto(
                    resource.chart_options
                )
            )
        else:
            res.ClearField("chart_options")
        return res

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return None

        return DashboardColumnLayoutColumnsWidgetsXyChart(
            data_sets=DashboardColumnLayoutColumnsWidgetsXyChartDataSetsArray.from_proto(
                resource.data_sets
            ),
            timeshift_duration=Primitive.from_proto(resource.timeshift_duration),
            thresholds=DashboardColumnLayoutColumnsWidgetsXyChartThresholdsArray.from_proto(
                resource.thresholds
            ),
            x_axis=DashboardColumnLayoutColumnsWidgetsXyChartXAxis.from_proto(
                resource.x_axis
            ),
            y_axis=DashboardColumnLayoutColumnsWidgetsXyChartYAxis.from_proto(
                resource.y_axis
            ),
            chart_options=DashboardColumnLayoutColumnsWidgetsXyChartChartOptions.from_proto(
                resource.chart_options
            ),
        )


class DashboardColumnLayoutColumnsWidgetsXyChartArray(object):
    @classmethod
    def to_proto(self, resources):
        if not resources:
            return resources
        return [
            DashboardColumnLayoutColumnsWidgetsXyChart.to_proto(i) for i in resources
        ]

    @classmethod
    def from_proto(self, resources):
        return [
            DashboardColumnLayoutColumnsWidgetsXyChart.from_proto(i) for i in resources
        ]


class DashboardColumnLayoutColumnsWidgetsXyChartDataSets(object):
    def __init__(
        self,
        time_series_query: dict = None,
        plot_type: str = None,
        legend_template: str = None,
        min_alignment_period: str = None,
    ):
        self.time_series_query = time_series_query
        self.plot_type = plot_type
        self.legend_template = legend_template
        self.min_alignment_period = min_alignment_period

    @classmethod
    def to_proto(self, resource):
        if not resource:
            return None

        res = (
            dashboard_pb2.MonitoringBetaDashboardColumnLayoutColumnsWidgetsXyChartDataSets()
        )
        if DashboardColumnLayoutColumnsWidgetsXyChartDataSetsTimeSeriesQuery.to_proto(
            resource.time_series_query
        ):
            res.time_series_query.CopyFrom(
                DashboardColumnLayoutColumnsWidgetsXyChartDataSetsTimeSeriesQuery.to_proto(
                    resource.time_series_query
                )
            )
        else:
            res.ClearField("time_series_query")
        if DashboardColumnLayoutColumnsWidgetsXyChartDataSetsPlotTypeEnum.to_proto(
            resource.plot_type
        ):
            res.plot_type = (
                DashboardColumnLayoutColumnsWidgetsXyChartDataSetsPlotTypeEnum.to_proto(
                    resource.plot_type
                )
            )
        if Primitive.to_proto(resource.legend_template):
            res.legend_template = Primitive.to_proto(resource.legend_template)
        if Primitive.to_proto(resource.min_alignment_period):
            res.min_alignment_period = Primitive.to_proto(resource.min_alignment_period)
        return res

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return None

        return DashboardColumnLayoutColumnsWidgetsXyChartDataSets(
            time_series_query=DashboardColumnLayoutColumnsWidgetsXyChartDataSetsTimeSeriesQuery.from_proto(
                resource.time_series_query
            ),
            plot_type=DashboardColumnLayoutColumnsWidgetsXyChartDataSetsPlotTypeEnum.from_proto(
                resource.plot_type
            ),
            legend_template=Primitive.from_proto(resource.legend_template),
            min_alignment_period=Primitive.from_proto(resource.min_alignment_period),
        )


class DashboardColumnLayoutColumnsWidgetsXyChartDataSetsArray(object):
    @classmethod
    def to_proto(self, resources):
        if not resources:
            return resources
        return [
            DashboardColumnLayoutColumnsWidgetsXyChartDataSets.to_proto(i)
            for i in resources
        ]

    @classmethod
    def from_proto(self, resources):
        return [
            DashboardColumnLayoutColumnsWidgetsXyChartDataSets.from_proto(i)
            for i in resources
        ]


class DashboardColumnLayoutColumnsWidgetsXyChartDataSetsTimeSeriesQuery(object):
    def __init__(
        self,
        time_series_filter: dict = None,
        time_series_filter_ratio: dict = None,
        time_series_query_language: str = None,
        unit_override: str = None,
    ):
        self.time_series_filter = time_series_filter
        self.time_series_filter_ratio = time_series_filter_ratio
        self.time_series_query_language = time_series_query_language
        self.unit_override = unit_override

    @classmethod
    def to_proto(self, resource):
        if not resource:
            return None

        res = (
            dashboard_pb2.MonitoringBetaDashboardColumnLayoutColumnsWidgetsXyChartDataSetsTimeSeriesQuery()
        )
        if DashboardColumnLayoutColumnsWidgetsXyChartDataSetsTimeSeriesQueryTimeSeriesFilter.to_proto(
            resource.time_series_filter
        ):
            res.time_series_filter.CopyFrom(
                DashboardColumnLayoutColumnsWidgetsXyChartDataSetsTimeSeriesQueryTimeSeriesFilter.to_proto(
                    resource.time_series_filter
                )
            )
        else:
            res.ClearField("time_series_filter")
        if DashboardColumnLayoutColumnsWidgetsXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatio.to_proto(
            resource.time_series_filter_ratio
        ):
            res.time_series_filter_ratio.CopyFrom(
                DashboardColumnLayoutColumnsWidgetsXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatio.to_proto(
                    resource.time_series_filter_ratio
                )
            )
        else:
            res.ClearField("time_series_filter_ratio")
        if Primitive.to_proto(resource.time_series_query_language):
            res.time_series_query_language = Primitive.to_proto(
                resource.time_series_query_language
            )
        if Primitive.to_proto(resource.unit_override):
            res.unit_override = Primitive.to_proto(resource.unit_override)
        return res

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return None

        return DashboardColumnLayoutColumnsWidgetsXyChartDataSetsTimeSeriesQuery(
            time_series_filter=DashboardColumnLayoutColumnsWidgetsXyChartDataSetsTimeSeriesQueryTimeSeriesFilter.from_proto(
                resource.time_series_filter
            ),
            time_series_filter_ratio=DashboardColumnLayoutColumnsWidgetsXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatio.from_proto(
                resource.time_series_filter_ratio
            ),
            time_series_query_language=Primitive.from_proto(
                resource.time_series_query_language
            ),
            unit_override=Primitive.from_proto(resource.unit_override),
        )


class DashboardColumnLayoutColumnsWidgetsXyChartDataSetsTimeSeriesQueryArray(object):
    @classmethod
    def to_proto(self, resources):
        if not resources:
            return resources
        return [
            DashboardColumnLayoutColumnsWidgetsXyChartDataSetsTimeSeriesQuery.to_proto(
                i
            )
            for i in resources
        ]

    @classmethod
    def from_proto(self, resources):
        return [
            DashboardColumnLayoutColumnsWidgetsXyChartDataSetsTimeSeriesQuery.from_proto(
                i
            )
            for i in resources
        ]


class DashboardColumnLayoutColumnsWidgetsXyChartDataSetsTimeSeriesQueryTimeSeriesFilter(
    object
):
    def __init__(
        self,
        filter: str = None,
        aggregation: dict = None,
        secondary_aggregation: dict = None,
        pick_time_series_filter: dict = None,
    ):
        self.filter = filter
        self.aggregation = aggregation
        self.secondary_aggregation = secondary_aggregation
        self.pick_time_series_filter = pick_time_series_filter

    @classmethod
    def to_proto(self, resource):
        if not resource:
            return None

        res = (
            dashboard_pb2.MonitoringBetaDashboardColumnLayoutColumnsWidgetsXyChartDataSetsTimeSeriesQueryTimeSeriesFilter()
        )
        if Primitive.to_proto(resource.filter):
            res.filter = Primitive.to_proto(resource.filter)
        if DashboardColumnLayoutColumnsWidgetsXyChartDataSetsTimeSeriesQueryTimeSeriesFilterAggregation.to_proto(
            resource.aggregation
        ):
            res.aggregation.CopyFrom(
                DashboardColumnLayoutColumnsWidgetsXyChartDataSetsTimeSeriesQueryTimeSeriesFilterAggregation.to_proto(
                    resource.aggregation
                )
            )
        else:
            res.ClearField("aggregation")
        if DashboardColumnLayoutColumnsWidgetsXyChartDataSetsTimeSeriesQueryTimeSeriesFilterSecondaryAggregation.to_proto(
            resource.secondary_aggregation
        ):
            res.secondary_aggregation.CopyFrom(
                DashboardColumnLayoutColumnsWidgetsXyChartDataSetsTimeSeriesQueryTimeSeriesFilterSecondaryAggregation.to_proto(
                    resource.secondary_aggregation
                )
            )
        else:
            res.ClearField("secondary_aggregation")
        if DashboardColumnLayoutColumnsWidgetsXyChartDataSetsTimeSeriesQueryTimeSeriesFilterPickTimeSeriesFilter.to_proto(
            resource.pick_time_series_filter
        ):
            res.pick_time_series_filter.CopyFrom(
                DashboardColumnLayoutColumnsWidgetsXyChartDataSetsTimeSeriesQueryTimeSeriesFilterPickTimeSeriesFilter.to_proto(
                    resource.pick_time_series_filter
                )
            )
        else:
            res.ClearField("pick_time_series_filter")
        return res

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return None

        return DashboardColumnLayoutColumnsWidgetsXyChartDataSetsTimeSeriesQueryTimeSeriesFilter(
            filter=Primitive.from_proto(resource.filter),
            aggregation=DashboardColumnLayoutColumnsWidgetsXyChartDataSetsTimeSeriesQueryTimeSeriesFilterAggregation.from_proto(
                resource.aggregation
            ),
            secondary_aggregation=DashboardColumnLayoutColumnsWidgetsXyChartDataSetsTimeSeriesQueryTimeSeriesFilterSecondaryAggregation.from_proto(
                resource.secondary_aggregation
            ),
            pick_time_series_filter=DashboardColumnLayoutColumnsWidgetsXyChartDataSetsTimeSeriesQueryTimeSeriesFilterPickTimeSeriesFilter.from_proto(
                resource.pick_time_series_filter
            ),
        )


class DashboardColumnLayoutColumnsWidgetsXyChartDataSetsTimeSeriesQueryTimeSeriesFilterArray(
    object
):
    @classmethod
    def to_proto(self, resources):
        if not resources:
            return resources
        return [
            DashboardColumnLayoutColumnsWidgetsXyChartDataSetsTimeSeriesQueryTimeSeriesFilter.to_proto(
                i
            )
            for i in resources
        ]

    @classmethod
    def from_proto(self, resources):
        return [
            DashboardColumnLayoutColumnsWidgetsXyChartDataSetsTimeSeriesQueryTimeSeriesFilter.from_proto(
                i
            )
            for i in resources
        ]


class DashboardColumnLayoutColumnsWidgetsXyChartDataSetsTimeSeriesQueryTimeSeriesFilterAggregation(
    object
):
    def __init__(
        self,
        alignment_period: str = None,
        per_series_aligner: str = None,
        cross_series_reducer: str = None,
        group_by_fields: list = None,
    ):
        self.alignment_period = alignment_period
        self.per_series_aligner = per_series_aligner
        self.cross_series_reducer = cross_series_reducer
        self.group_by_fields = group_by_fields

    @classmethod
    def to_proto(self, resource):
        if not resource:
            return None

        res = (
            dashboard_pb2.MonitoringBetaDashboardColumnLayoutColumnsWidgetsXyChartDataSetsTimeSeriesQueryTimeSeriesFilterAggregation()
        )
        if Primitive.to_proto(resource.alignment_period):
            res.alignment_period = Primitive.to_proto(resource.alignment_period)
        if DashboardColumnLayoutColumnsWidgetsXyChartDataSetsTimeSeriesQueryTimeSeriesFilterAggregationPerSeriesAlignerEnum.to_proto(
            resource.per_series_aligner
        ):
            res.per_series_aligner = DashboardColumnLayoutColumnsWidgetsXyChartDataSetsTimeSeriesQueryTimeSeriesFilterAggregationPerSeriesAlignerEnum.to_proto(
                resource.per_series_aligner
            )
        if DashboardColumnLayoutColumnsWidgetsXyChartDataSetsTimeSeriesQueryTimeSeriesFilterAggregationCrossSeriesReducerEnum.to_proto(
            resource.cross_series_reducer
        ):
            res.cross_series_reducer = DashboardColumnLayoutColumnsWidgetsXyChartDataSetsTimeSeriesQueryTimeSeriesFilterAggregationCrossSeriesReducerEnum.to_proto(
                resource.cross_series_reducer
            )
        if Primitive.to_proto(resource.group_by_fields):
            res.group_by_fields.extend(Primitive.to_proto(resource.group_by_fields))
        return res

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return None

        return DashboardColumnLayoutColumnsWidgetsXyChartDataSetsTimeSeriesQueryTimeSeriesFilterAggregation(
            alignment_period=Primitive.from_proto(resource.alignment_period),
            per_series_aligner=DashboardColumnLayoutColumnsWidgetsXyChartDataSetsTimeSeriesQueryTimeSeriesFilterAggregationPerSeriesAlignerEnum.from_proto(
                resource.per_series_aligner
            ),
            cross_series_reducer=DashboardColumnLayoutColumnsWidgetsXyChartDataSetsTimeSeriesQueryTimeSeriesFilterAggregationCrossSeriesReducerEnum.from_proto(
                resource.cross_series_reducer
            ),
            group_by_fields=Primitive.from_proto(resource.group_by_fields),
        )


class DashboardColumnLayoutColumnsWidgetsXyChartDataSetsTimeSeriesQueryTimeSeriesFilterAggregationArray(
    object
):
    @classmethod
    def to_proto(self, resources):
        if not resources:
            return resources
        return [
            DashboardColumnLayoutColumnsWidgetsXyChartDataSetsTimeSeriesQueryTimeSeriesFilterAggregation.to_proto(
                i
            )
            for i in resources
        ]

    @classmethod
    def from_proto(self, resources):
        return [
            DashboardColumnLayoutColumnsWidgetsXyChartDataSetsTimeSeriesQueryTimeSeriesFilterAggregation.from_proto(
                i
            )
            for i in resources
        ]


class DashboardColumnLayoutColumnsWidgetsXyChartDataSetsTimeSeriesQueryTimeSeriesFilterSecondaryAggregation(
    object
):
    def __init__(
        self,
        alignment_period: str = None,
        per_series_aligner: str = None,
        cross_series_reducer: str = None,
        group_by_fields: list = None,
    ):
        self.alignment_period = alignment_period
        self.per_series_aligner = per_series_aligner
        self.cross_series_reducer = cross_series_reducer
        self.group_by_fields = group_by_fields

    @classmethod
    def to_proto(self, resource):
        if not resource:
            return None

        res = (
            dashboard_pb2.MonitoringBetaDashboardColumnLayoutColumnsWidgetsXyChartDataSetsTimeSeriesQueryTimeSeriesFilterSecondaryAggregation()
        )
        if Primitive.to_proto(resource.alignment_period):
            res.alignment_period = Primitive.to_proto(resource.alignment_period)
        if DashboardColumnLayoutColumnsWidgetsXyChartDataSetsTimeSeriesQueryTimeSeriesFilterSecondaryAggregationPerSeriesAlignerEnum.to_proto(
            resource.per_series_aligner
        ):
            res.per_series_aligner = DashboardColumnLayoutColumnsWidgetsXyChartDataSetsTimeSeriesQueryTimeSeriesFilterSecondaryAggregationPerSeriesAlignerEnum.to_proto(
                resource.per_series_aligner
            )
        if DashboardColumnLayoutColumnsWidgetsXyChartDataSetsTimeSeriesQueryTimeSeriesFilterSecondaryAggregationCrossSeriesReducerEnum.to_proto(
            resource.cross_series_reducer
        ):
            res.cross_series_reducer = DashboardColumnLayoutColumnsWidgetsXyChartDataSetsTimeSeriesQueryTimeSeriesFilterSecondaryAggregationCrossSeriesReducerEnum.to_proto(
                resource.cross_series_reducer
            )
        if Primitive.to_proto(resource.group_by_fields):
            res.group_by_fields.extend(Primitive.to_proto(resource.group_by_fields))
        return res

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return None

        return DashboardColumnLayoutColumnsWidgetsXyChartDataSetsTimeSeriesQueryTimeSeriesFilterSecondaryAggregation(
            alignment_period=Primitive.from_proto(resource.alignment_period),
            per_series_aligner=DashboardColumnLayoutColumnsWidgetsXyChartDataSetsTimeSeriesQueryTimeSeriesFilterSecondaryAggregationPerSeriesAlignerEnum.from_proto(
                resource.per_series_aligner
            ),
            cross_series_reducer=DashboardColumnLayoutColumnsWidgetsXyChartDataSetsTimeSeriesQueryTimeSeriesFilterSecondaryAggregationCrossSeriesReducerEnum.from_proto(
                resource.cross_series_reducer
            ),
            group_by_fields=Primitive.from_proto(resource.group_by_fields),
        )


class DashboardColumnLayoutColumnsWidgetsXyChartDataSetsTimeSeriesQueryTimeSeriesFilterSecondaryAggregationArray(
    object
):
    @classmethod
    def to_proto(self, resources):
        if not resources:
            return resources
        return [
            DashboardColumnLayoutColumnsWidgetsXyChartDataSetsTimeSeriesQueryTimeSeriesFilterSecondaryAggregation.to_proto(
                i
            )
            for i in resources
        ]

    @classmethod
    def from_proto(self, resources):
        return [
            DashboardColumnLayoutColumnsWidgetsXyChartDataSetsTimeSeriesQueryTimeSeriesFilterSecondaryAggregation.from_proto(
                i
            )
            for i in resources
        ]


class DashboardColumnLayoutColumnsWidgetsXyChartDataSetsTimeSeriesQueryTimeSeriesFilterPickTimeSeriesFilter(
    object
):
    def __init__(
        self,
        ranking_method: str = None,
        num_time_series: int = None,
        direction: str = None,
    ):
        self.ranking_method = ranking_method
        self.num_time_series = num_time_series
        self.direction = direction

    @classmethod
    def to_proto(self, resource):
        if not resource:
            return None

        res = (
            dashboard_pb2.MonitoringBetaDashboardColumnLayoutColumnsWidgetsXyChartDataSetsTimeSeriesQueryTimeSeriesFilterPickTimeSeriesFilter()
        )
        if DashboardColumnLayoutColumnsWidgetsXyChartDataSetsTimeSeriesQueryTimeSeriesFilterPickTimeSeriesFilterRankingMethodEnum.to_proto(
            resource.ranking_method
        ):
            res.ranking_method = DashboardColumnLayoutColumnsWidgetsXyChartDataSetsTimeSeriesQueryTimeSeriesFilterPickTimeSeriesFilterRankingMethodEnum.to_proto(
                resource.ranking_method
            )
        if Primitive.to_proto(resource.num_time_series):
            res.num_time_series = Primitive.to_proto(resource.num_time_series)
        if DashboardColumnLayoutColumnsWidgetsXyChartDataSetsTimeSeriesQueryTimeSeriesFilterPickTimeSeriesFilterDirectionEnum.to_proto(
            resource.direction
        ):
            res.direction = DashboardColumnLayoutColumnsWidgetsXyChartDataSetsTimeSeriesQueryTimeSeriesFilterPickTimeSeriesFilterDirectionEnum.to_proto(
                resource.direction
            )
        return res

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return None

        return DashboardColumnLayoutColumnsWidgetsXyChartDataSetsTimeSeriesQueryTimeSeriesFilterPickTimeSeriesFilter(
            ranking_method=DashboardColumnLayoutColumnsWidgetsXyChartDataSetsTimeSeriesQueryTimeSeriesFilterPickTimeSeriesFilterRankingMethodEnum.from_proto(
                resource.ranking_method
            ),
            num_time_series=Primitive.from_proto(resource.num_time_series),
            direction=DashboardColumnLayoutColumnsWidgetsXyChartDataSetsTimeSeriesQueryTimeSeriesFilterPickTimeSeriesFilterDirectionEnum.from_proto(
                resource.direction
            ),
        )


class DashboardColumnLayoutColumnsWidgetsXyChartDataSetsTimeSeriesQueryTimeSeriesFilterPickTimeSeriesFilterArray(
    object
):
    @classmethod
    def to_proto(self, resources):
        if not resources:
            return resources
        return [
            DashboardColumnLayoutColumnsWidgetsXyChartDataSetsTimeSeriesQueryTimeSeriesFilterPickTimeSeriesFilter.to_proto(
                i
            )
            for i in resources
        ]

    @classmethod
    def from_proto(self, resources):
        return [
            DashboardColumnLayoutColumnsWidgetsXyChartDataSetsTimeSeriesQueryTimeSeriesFilterPickTimeSeriesFilter.from_proto(
                i
            )
            for i in resources
        ]


class DashboardColumnLayoutColumnsWidgetsXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatio(
    object
):
    def __init__(
        self,
        numerator: dict = None,
        denominator: dict = None,
        secondary_aggregation: dict = None,
        pick_time_series_filter: dict = None,
    ):
        self.numerator = numerator
        self.denominator = denominator
        self.secondary_aggregation = secondary_aggregation
        self.pick_time_series_filter = pick_time_series_filter

    @classmethod
    def to_proto(self, resource):
        if not resource:
            return None

        res = (
            dashboard_pb2.MonitoringBetaDashboardColumnLayoutColumnsWidgetsXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatio()
        )
        if DashboardColumnLayoutColumnsWidgetsXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioNumerator.to_proto(
            resource.numerator
        ):
            res.numerator.CopyFrom(
                DashboardColumnLayoutColumnsWidgetsXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioNumerator.to_proto(
                    resource.numerator
                )
            )
        else:
            res.ClearField("numerator")
        if DashboardColumnLayoutColumnsWidgetsXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioDenominator.to_proto(
            resource.denominator
        ):
            res.denominator.CopyFrom(
                DashboardColumnLayoutColumnsWidgetsXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioDenominator.to_proto(
                    resource.denominator
                )
            )
        else:
            res.ClearField("denominator")
        if DashboardColumnLayoutColumnsWidgetsXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioSecondaryAggregation.to_proto(
            resource.secondary_aggregation
        ):
            res.secondary_aggregation.CopyFrom(
                DashboardColumnLayoutColumnsWidgetsXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioSecondaryAggregation.to_proto(
                    resource.secondary_aggregation
                )
            )
        else:
            res.ClearField("secondary_aggregation")
        if DashboardColumnLayoutColumnsWidgetsXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioPickTimeSeriesFilter.to_proto(
            resource.pick_time_series_filter
        ):
            res.pick_time_series_filter.CopyFrom(
                DashboardColumnLayoutColumnsWidgetsXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioPickTimeSeriesFilter.to_proto(
                    resource.pick_time_series_filter
                )
            )
        else:
            res.ClearField("pick_time_series_filter")
        return res

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return None

        return DashboardColumnLayoutColumnsWidgetsXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatio(
            numerator=DashboardColumnLayoutColumnsWidgetsXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioNumerator.from_proto(
                resource.numerator
            ),
            denominator=DashboardColumnLayoutColumnsWidgetsXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioDenominator.from_proto(
                resource.denominator
            ),
            secondary_aggregation=DashboardColumnLayoutColumnsWidgetsXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioSecondaryAggregation.from_proto(
                resource.secondary_aggregation
            ),
            pick_time_series_filter=DashboardColumnLayoutColumnsWidgetsXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioPickTimeSeriesFilter.from_proto(
                resource.pick_time_series_filter
            ),
        )


class DashboardColumnLayoutColumnsWidgetsXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioArray(
    object
):
    @classmethod
    def to_proto(self, resources):
        if not resources:
            return resources
        return [
            DashboardColumnLayoutColumnsWidgetsXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatio.to_proto(
                i
            )
            for i in resources
        ]

    @classmethod
    def from_proto(self, resources):
        return [
            DashboardColumnLayoutColumnsWidgetsXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatio.from_proto(
                i
            )
            for i in resources
        ]


class DashboardColumnLayoutColumnsWidgetsXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioNumerator(
    object
):
    def __init__(self, filter: str = None, aggregation: dict = None):
        self.filter = filter
        self.aggregation = aggregation

    @classmethod
    def to_proto(self, resource):
        if not resource:
            return None

        res = (
            dashboard_pb2.MonitoringBetaDashboardColumnLayoutColumnsWidgetsXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioNumerator()
        )
        if Primitive.to_proto(resource.filter):
            res.filter = Primitive.to_proto(resource.filter)
        if DashboardColumnLayoutColumnsWidgetsXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioNumeratorAggregation.to_proto(
            resource.aggregation
        ):
            res.aggregation.CopyFrom(
                DashboardColumnLayoutColumnsWidgetsXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioNumeratorAggregation.to_proto(
                    resource.aggregation
                )
            )
        else:
            res.ClearField("aggregation")
        return res

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return None

        return DashboardColumnLayoutColumnsWidgetsXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioNumerator(
            filter=Primitive.from_proto(resource.filter),
            aggregation=DashboardColumnLayoutColumnsWidgetsXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioNumeratorAggregation.from_proto(
                resource.aggregation
            ),
        )


class DashboardColumnLayoutColumnsWidgetsXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioNumeratorArray(
    object
):
    @classmethod
    def to_proto(self, resources):
        if not resources:
            return resources
        return [
            DashboardColumnLayoutColumnsWidgetsXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioNumerator.to_proto(
                i
            )
            for i in resources
        ]

    @classmethod
    def from_proto(self, resources):
        return [
            DashboardColumnLayoutColumnsWidgetsXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioNumerator.from_proto(
                i
            )
            for i in resources
        ]


class DashboardColumnLayoutColumnsWidgetsXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioNumeratorAggregation(
    object
):
    def __init__(
        self,
        alignment_period: str = None,
        per_series_aligner: str = None,
        cross_series_reducer: str = None,
        group_by_fields: list = None,
    ):
        self.alignment_period = alignment_period
        self.per_series_aligner = per_series_aligner
        self.cross_series_reducer = cross_series_reducer
        self.group_by_fields = group_by_fields

    @classmethod
    def to_proto(self, resource):
        if not resource:
            return None

        res = (
            dashboard_pb2.MonitoringBetaDashboardColumnLayoutColumnsWidgetsXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioNumeratorAggregation()
        )
        if Primitive.to_proto(resource.alignment_period):
            res.alignment_period = Primitive.to_proto(resource.alignment_period)
        if DashboardColumnLayoutColumnsWidgetsXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioNumeratorAggregationPerSeriesAlignerEnum.to_proto(
            resource.per_series_aligner
        ):
            res.per_series_aligner = DashboardColumnLayoutColumnsWidgetsXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioNumeratorAggregationPerSeriesAlignerEnum.to_proto(
                resource.per_series_aligner
            )
        if DashboardColumnLayoutColumnsWidgetsXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioNumeratorAggregationCrossSeriesReducerEnum.to_proto(
            resource.cross_series_reducer
        ):
            res.cross_series_reducer = DashboardColumnLayoutColumnsWidgetsXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioNumeratorAggregationCrossSeriesReducerEnum.to_proto(
                resource.cross_series_reducer
            )
        if Primitive.to_proto(resource.group_by_fields):
            res.group_by_fields.extend(Primitive.to_proto(resource.group_by_fields))
        return res

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return None

        return DashboardColumnLayoutColumnsWidgetsXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioNumeratorAggregation(
            alignment_period=Primitive.from_proto(resource.alignment_period),
            per_series_aligner=DashboardColumnLayoutColumnsWidgetsXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioNumeratorAggregationPerSeriesAlignerEnum.from_proto(
                resource.per_series_aligner
            ),
            cross_series_reducer=DashboardColumnLayoutColumnsWidgetsXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioNumeratorAggregationCrossSeriesReducerEnum.from_proto(
                resource.cross_series_reducer
            ),
            group_by_fields=Primitive.from_proto(resource.group_by_fields),
        )


class DashboardColumnLayoutColumnsWidgetsXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioNumeratorAggregationArray(
    object
):
    @classmethod
    def to_proto(self, resources):
        if not resources:
            return resources
        return [
            DashboardColumnLayoutColumnsWidgetsXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioNumeratorAggregation.to_proto(
                i
            )
            for i in resources
        ]

    @classmethod
    def from_proto(self, resources):
        return [
            DashboardColumnLayoutColumnsWidgetsXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioNumeratorAggregation.from_proto(
                i
            )
            for i in resources
        ]


class DashboardColumnLayoutColumnsWidgetsXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioDenominator(
    object
):
    def __init__(self, filter: str = None, aggregation: dict = None):
        self.filter = filter
        self.aggregation = aggregation

    @classmethod
    def to_proto(self, resource):
        if not resource:
            return None

        res = (
            dashboard_pb2.MonitoringBetaDashboardColumnLayoutColumnsWidgetsXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioDenominator()
        )
        if Primitive.to_proto(resource.filter):
            res.filter = Primitive.to_proto(resource.filter)
        if DashboardColumnLayoutColumnsWidgetsXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioDenominatorAggregation.to_proto(
            resource.aggregation
        ):
            res.aggregation.CopyFrom(
                DashboardColumnLayoutColumnsWidgetsXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioDenominatorAggregation.to_proto(
                    resource.aggregation
                )
            )
        else:
            res.ClearField("aggregation")
        return res

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return None

        return DashboardColumnLayoutColumnsWidgetsXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioDenominator(
            filter=Primitive.from_proto(resource.filter),
            aggregation=DashboardColumnLayoutColumnsWidgetsXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioDenominatorAggregation.from_proto(
                resource.aggregation
            ),
        )


class DashboardColumnLayoutColumnsWidgetsXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioDenominatorArray(
    object
):
    @classmethod
    def to_proto(self, resources):
        if not resources:
            return resources
        return [
            DashboardColumnLayoutColumnsWidgetsXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioDenominator.to_proto(
                i
            )
            for i in resources
        ]

    @classmethod
    def from_proto(self, resources):
        return [
            DashboardColumnLayoutColumnsWidgetsXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioDenominator.from_proto(
                i
            )
            for i in resources
        ]


class DashboardColumnLayoutColumnsWidgetsXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioDenominatorAggregation(
    object
):
    def __init__(
        self,
        alignment_period: str = None,
        per_series_aligner: str = None,
        cross_series_reducer: str = None,
        group_by_fields: list = None,
    ):
        self.alignment_period = alignment_period
        self.per_series_aligner = per_series_aligner
        self.cross_series_reducer = cross_series_reducer
        self.group_by_fields = group_by_fields

    @classmethod
    def to_proto(self, resource):
        if not resource:
            return None

        res = (
            dashboard_pb2.MonitoringBetaDashboardColumnLayoutColumnsWidgetsXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioDenominatorAggregation()
        )
        if Primitive.to_proto(resource.alignment_period):
            res.alignment_period = Primitive.to_proto(resource.alignment_period)
        if DashboardColumnLayoutColumnsWidgetsXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioDenominatorAggregationPerSeriesAlignerEnum.to_proto(
            resource.per_series_aligner
        ):
            res.per_series_aligner = DashboardColumnLayoutColumnsWidgetsXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioDenominatorAggregationPerSeriesAlignerEnum.to_proto(
                resource.per_series_aligner
            )
        if DashboardColumnLayoutColumnsWidgetsXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioDenominatorAggregationCrossSeriesReducerEnum.to_proto(
            resource.cross_series_reducer
        ):
            res.cross_series_reducer = DashboardColumnLayoutColumnsWidgetsXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioDenominatorAggregationCrossSeriesReducerEnum.to_proto(
                resource.cross_series_reducer
            )
        if Primitive.to_proto(resource.group_by_fields):
            res.group_by_fields.extend(Primitive.to_proto(resource.group_by_fields))
        return res

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return None

        return DashboardColumnLayoutColumnsWidgetsXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioDenominatorAggregation(
            alignment_period=Primitive.from_proto(resource.alignment_period),
            per_series_aligner=DashboardColumnLayoutColumnsWidgetsXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioDenominatorAggregationPerSeriesAlignerEnum.from_proto(
                resource.per_series_aligner
            ),
            cross_series_reducer=DashboardColumnLayoutColumnsWidgetsXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioDenominatorAggregationCrossSeriesReducerEnum.from_proto(
                resource.cross_series_reducer
            ),
            group_by_fields=Primitive.from_proto(resource.group_by_fields),
        )


class DashboardColumnLayoutColumnsWidgetsXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioDenominatorAggregationArray(
    object
):
    @classmethod
    def to_proto(self, resources):
        if not resources:
            return resources
        return [
            DashboardColumnLayoutColumnsWidgetsXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioDenominatorAggregation.to_proto(
                i
            )
            for i in resources
        ]

    @classmethod
    def from_proto(self, resources):
        return [
            DashboardColumnLayoutColumnsWidgetsXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioDenominatorAggregation.from_proto(
                i
            )
            for i in resources
        ]


class DashboardColumnLayoutColumnsWidgetsXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioSecondaryAggregation(
    object
):
    def __init__(
        self,
        alignment_period: str = None,
        per_series_aligner: str = None,
        cross_series_reducer: str = None,
        group_by_fields: list = None,
    ):
        self.alignment_period = alignment_period
        self.per_series_aligner = per_series_aligner
        self.cross_series_reducer = cross_series_reducer
        self.group_by_fields = group_by_fields

    @classmethod
    def to_proto(self, resource):
        if not resource:
            return None

        res = (
            dashboard_pb2.MonitoringBetaDashboardColumnLayoutColumnsWidgetsXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioSecondaryAggregation()
        )
        if Primitive.to_proto(resource.alignment_period):
            res.alignment_period = Primitive.to_proto(resource.alignment_period)
        if DashboardColumnLayoutColumnsWidgetsXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioSecondaryAggregationPerSeriesAlignerEnum.to_proto(
            resource.per_series_aligner
        ):
            res.per_series_aligner = DashboardColumnLayoutColumnsWidgetsXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioSecondaryAggregationPerSeriesAlignerEnum.to_proto(
                resource.per_series_aligner
            )
        if DashboardColumnLayoutColumnsWidgetsXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioSecondaryAggregationCrossSeriesReducerEnum.to_proto(
            resource.cross_series_reducer
        ):
            res.cross_series_reducer = DashboardColumnLayoutColumnsWidgetsXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioSecondaryAggregationCrossSeriesReducerEnum.to_proto(
                resource.cross_series_reducer
            )
        if Primitive.to_proto(resource.group_by_fields):
            res.group_by_fields.extend(Primitive.to_proto(resource.group_by_fields))
        return res

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return None

        return DashboardColumnLayoutColumnsWidgetsXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioSecondaryAggregation(
            alignment_period=Primitive.from_proto(resource.alignment_period),
            per_series_aligner=DashboardColumnLayoutColumnsWidgetsXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioSecondaryAggregationPerSeriesAlignerEnum.from_proto(
                resource.per_series_aligner
            ),
            cross_series_reducer=DashboardColumnLayoutColumnsWidgetsXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioSecondaryAggregationCrossSeriesReducerEnum.from_proto(
                resource.cross_series_reducer
            ),
            group_by_fields=Primitive.from_proto(resource.group_by_fields),
        )


class DashboardColumnLayoutColumnsWidgetsXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioSecondaryAggregationArray(
    object
):
    @classmethod
    def to_proto(self, resources):
        if not resources:
            return resources
        return [
            DashboardColumnLayoutColumnsWidgetsXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioSecondaryAggregation.to_proto(
                i
            )
            for i in resources
        ]

    @classmethod
    def from_proto(self, resources):
        return [
            DashboardColumnLayoutColumnsWidgetsXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioSecondaryAggregation.from_proto(
                i
            )
            for i in resources
        ]


class DashboardColumnLayoutColumnsWidgetsXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioPickTimeSeriesFilter(
    object
):
    def __init__(
        self,
        ranking_method: str = None,
        num_time_series: int = None,
        direction: str = None,
    ):
        self.ranking_method = ranking_method
        self.num_time_series = num_time_series
        self.direction = direction

    @classmethod
    def to_proto(self, resource):
        if not resource:
            return None

        res = (
            dashboard_pb2.MonitoringBetaDashboardColumnLayoutColumnsWidgetsXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioPickTimeSeriesFilter()
        )
        if DashboardColumnLayoutColumnsWidgetsXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioPickTimeSeriesFilterRankingMethodEnum.to_proto(
            resource.ranking_method
        ):
            res.ranking_method = DashboardColumnLayoutColumnsWidgetsXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioPickTimeSeriesFilterRankingMethodEnum.to_proto(
                resource.ranking_method
            )
        if Primitive.to_proto(resource.num_time_series):
            res.num_time_series = Primitive.to_proto(resource.num_time_series)
        if DashboardColumnLayoutColumnsWidgetsXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioPickTimeSeriesFilterDirectionEnum.to_proto(
            resource.direction
        ):
            res.direction = DashboardColumnLayoutColumnsWidgetsXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioPickTimeSeriesFilterDirectionEnum.to_proto(
                resource.direction
            )
        return res

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return None

        return DashboardColumnLayoutColumnsWidgetsXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioPickTimeSeriesFilter(
            ranking_method=DashboardColumnLayoutColumnsWidgetsXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioPickTimeSeriesFilterRankingMethodEnum.from_proto(
                resource.ranking_method
            ),
            num_time_series=Primitive.from_proto(resource.num_time_series),
            direction=DashboardColumnLayoutColumnsWidgetsXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioPickTimeSeriesFilterDirectionEnum.from_proto(
                resource.direction
            ),
        )


class DashboardColumnLayoutColumnsWidgetsXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioPickTimeSeriesFilterArray(
    object
):
    @classmethod
    def to_proto(self, resources):
        if not resources:
            return resources
        return [
            DashboardColumnLayoutColumnsWidgetsXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioPickTimeSeriesFilter.to_proto(
                i
            )
            for i in resources
        ]

    @classmethod
    def from_proto(self, resources):
        return [
            DashboardColumnLayoutColumnsWidgetsXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioPickTimeSeriesFilter.from_proto(
                i
            )
            for i in resources
        ]


class DashboardColumnLayoutColumnsWidgetsXyChartThresholds(object):
    def __init__(
        self,
        label: str = None,
        value: float = None,
        color: str = None,
        direction: str = None,
    ):
        self.label = label
        self.value = value
        self.color = color
        self.direction = direction

    @classmethod
    def to_proto(self, resource):
        if not resource:
            return None

        res = (
            dashboard_pb2.MonitoringBetaDashboardColumnLayoutColumnsWidgetsXyChartThresholds()
        )
        if Primitive.to_proto(resource.label):
            res.label = Primitive.to_proto(resource.label)
        if Primitive.to_proto(resource.value):
            res.value = Primitive.to_proto(resource.value)
        if DashboardColumnLayoutColumnsWidgetsXyChartThresholdsColorEnum.to_proto(
            resource.color
        ):
            res.color = (
                DashboardColumnLayoutColumnsWidgetsXyChartThresholdsColorEnum.to_proto(
                    resource.color
                )
            )
        if DashboardColumnLayoutColumnsWidgetsXyChartThresholdsDirectionEnum.to_proto(
            resource.direction
        ):
            res.direction = DashboardColumnLayoutColumnsWidgetsXyChartThresholdsDirectionEnum.to_proto(
                resource.direction
            )
        return res

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return None

        return DashboardColumnLayoutColumnsWidgetsXyChartThresholds(
            label=Primitive.from_proto(resource.label),
            value=Primitive.from_proto(resource.value),
            color=DashboardColumnLayoutColumnsWidgetsXyChartThresholdsColorEnum.from_proto(
                resource.color
            ),
            direction=DashboardColumnLayoutColumnsWidgetsXyChartThresholdsDirectionEnum.from_proto(
                resource.direction
            ),
        )


class DashboardColumnLayoutColumnsWidgetsXyChartThresholdsArray(object):
    @classmethod
    def to_proto(self, resources):
        if not resources:
            return resources
        return [
            DashboardColumnLayoutColumnsWidgetsXyChartThresholds.to_proto(i)
            for i in resources
        ]

    @classmethod
    def from_proto(self, resources):
        return [
            DashboardColumnLayoutColumnsWidgetsXyChartThresholds.from_proto(i)
            for i in resources
        ]


class DashboardColumnLayoutColumnsWidgetsXyChartXAxis(object):
    def __init__(self, label: str = None, scale: str = None):
        self.label = label
        self.scale = scale

    @classmethod
    def to_proto(self, resource):
        if not resource:
            return None

        res = (
            dashboard_pb2.MonitoringBetaDashboardColumnLayoutColumnsWidgetsXyChartXAxis()
        )
        if Primitive.to_proto(resource.label):
            res.label = Primitive.to_proto(resource.label)
        if DashboardColumnLayoutColumnsWidgetsXyChartXAxisScaleEnum.to_proto(
            resource.scale
        ):
            res.scale = (
                DashboardColumnLayoutColumnsWidgetsXyChartXAxisScaleEnum.to_proto(
                    resource.scale
                )
            )
        return res

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return None

        return DashboardColumnLayoutColumnsWidgetsXyChartXAxis(
            label=Primitive.from_proto(resource.label),
            scale=DashboardColumnLayoutColumnsWidgetsXyChartXAxisScaleEnum.from_proto(
                resource.scale
            ),
        )


class DashboardColumnLayoutColumnsWidgetsXyChartXAxisArray(object):
    @classmethod
    def to_proto(self, resources):
        if not resources:
            return resources
        return [
            DashboardColumnLayoutColumnsWidgetsXyChartXAxis.to_proto(i)
            for i in resources
        ]

    @classmethod
    def from_proto(self, resources):
        return [
            DashboardColumnLayoutColumnsWidgetsXyChartXAxis.from_proto(i)
            for i in resources
        ]


class DashboardColumnLayoutColumnsWidgetsXyChartYAxis(object):
    def __init__(self, label: str = None, scale: str = None):
        self.label = label
        self.scale = scale

    @classmethod
    def to_proto(self, resource):
        if not resource:
            return None

        res = (
            dashboard_pb2.MonitoringBetaDashboardColumnLayoutColumnsWidgetsXyChartYAxis()
        )
        if Primitive.to_proto(resource.label):
            res.label = Primitive.to_proto(resource.label)
        if DashboardColumnLayoutColumnsWidgetsXyChartYAxisScaleEnum.to_proto(
            resource.scale
        ):
            res.scale = (
                DashboardColumnLayoutColumnsWidgetsXyChartYAxisScaleEnum.to_proto(
                    resource.scale
                )
            )
        return res

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return None

        return DashboardColumnLayoutColumnsWidgetsXyChartYAxis(
            label=Primitive.from_proto(resource.label),
            scale=DashboardColumnLayoutColumnsWidgetsXyChartYAxisScaleEnum.from_proto(
                resource.scale
            ),
        )


class DashboardColumnLayoutColumnsWidgetsXyChartYAxisArray(object):
    @classmethod
    def to_proto(self, resources):
        if not resources:
            return resources
        return [
            DashboardColumnLayoutColumnsWidgetsXyChartYAxis.to_proto(i)
            for i in resources
        ]

    @classmethod
    def from_proto(self, resources):
        return [
            DashboardColumnLayoutColumnsWidgetsXyChartYAxis.from_proto(i)
            for i in resources
        ]


class DashboardColumnLayoutColumnsWidgetsXyChartChartOptions(object):
    def __init__(self, mode: str = None):
        self.mode = mode

    @classmethod
    def to_proto(self, resource):
        if not resource:
            return None

        res = (
            dashboard_pb2.MonitoringBetaDashboardColumnLayoutColumnsWidgetsXyChartChartOptions()
        )
        if DashboardColumnLayoutColumnsWidgetsXyChartChartOptionsModeEnum.to_proto(
            resource.mode
        ):
            res.mode = (
                DashboardColumnLayoutColumnsWidgetsXyChartChartOptionsModeEnum.to_proto(
                    resource.mode
                )
            )
        return res

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return None

        return DashboardColumnLayoutColumnsWidgetsXyChartChartOptions(
            mode=DashboardColumnLayoutColumnsWidgetsXyChartChartOptionsModeEnum.from_proto(
                resource.mode
            ),
        )


class DashboardColumnLayoutColumnsWidgetsXyChartChartOptionsArray(object):
    @classmethod
    def to_proto(self, resources):
        if not resources:
            return resources
        return [
            DashboardColumnLayoutColumnsWidgetsXyChartChartOptions.to_proto(i)
            for i in resources
        ]

    @classmethod
    def from_proto(self, resources):
        return [
            DashboardColumnLayoutColumnsWidgetsXyChartChartOptions.from_proto(i)
            for i in resources
        ]


class DashboardColumnLayoutColumnsWidgetsScorecard(object):
    def __init__(
        self,
        time_series_query: dict = None,
        gauge_view: dict = None,
        spark_chart_view: dict = None,
        thresholds: list = None,
    ):
        self.time_series_query = time_series_query
        self.gauge_view = gauge_view
        self.spark_chart_view = spark_chart_view
        self.thresholds = thresholds

    @classmethod
    def to_proto(self, resource):
        if not resource:
            return None

        res = dashboard_pb2.MonitoringBetaDashboardColumnLayoutColumnsWidgetsScorecard()
        if DashboardColumnLayoutColumnsWidgetsScorecardTimeSeriesQuery.to_proto(
            resource.time_series_query
        ):
            res.time_series_query.CopyFrom(
                DashboardColumnLayoutColumnsWidgetsScorecardTimeSeriesQuery.to_proto(
                    resource.time_series_query
                )
            )
        else:
            res.ClearField("time_series_query")
        if DashboardColumnLayoutColumnsWidgetsScorecardGaugeView.to_proto(
            resource.gauge_view
        ):
            res.gauge_view.CopyFrom(
                DashboardColumnLayoutColumnsWidgetsScorecardGaugeView.to_proto(
                    resource.gauge_view
                )
            )
        else:
            res.ClearField("gauge_view")
        if DashboardColumnLayoutColumnsWidgetsScorecardSparkChartView.to_proto(
            resource.spark_chart_view
        ):
            res.spark_chart_view.CopyFrom(
                DashboardColumnLayoutColumnsWidgetsScorecardSparkChartView.to_proto(
                    resource.spark_chart_view
                )
            )
        else:
            res.ClearField("spark_chart_view")
        if DashboardColumnLayoutColumnsWidgetsScorecardThresholdsArray.to_proto(
            resource.thresholds
        ):
            res.thresholds.extend(
                DashboardColumnLayoutColumnsWidgetsScorecardThresholdsArray.to_proto(
                    resource.thresholds
                )
            )
        return res

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return None

        return DashboardColumnLayoutColumnsWidgetsScorecard(
            time_series_query=DashboardColumnLayoutColumnsWidgetsScorecardTimeSeriesQuery.from_proto(
                resource.time_series_query
            ),
            gauge_view=DashboardColumnLayoutColumnsWidgetsScorecardGaugeView.from_proto(
                resource.gauge_view
            ),
            spark_chart_view=DashboardColumnLayoutColumnsWidgetsScorecardSparkChartView.from_proto(
                resource.spark_chart_view
            ),
            thresholds=DashboardColumnLayoutColumnsWidgetsScorecardThresholdsArray.from_proto(
                resource.thresholds
            ),
        )


class DashboardColumnLayoutColumnsWidgetsScorecardArray(object):
    @classmethod
    def to_proto(self, resources):
        if not resources:
            return resources
        return [
            DashboardColumnLayoutColumnsWidgetsScorecard.to_proto(i) for i in resources
        ]

    @classmethod
    def from_proto(self, resources):
        return [
            DashboardColumnLayoutColumnsWidgetsScorecard.from_proto(i)
            for i in resources
        ]


class DashboardColumnLayoutColumnsWidgetsScorecardTimeSeriesQuery(object):
    def __init__(
        self,
        time_series_filter: dict = None,
        time_series_filter_ratio: dict = None,
        time_series_query_language: str = None,
        unit_override: str = None,
    ):
        self.time_series_filter = time_series_filter
        self.time_series_filter_ratio = time_series_filter_ratio
        self.time_series_query_language = time_series_query_language
        self.unit_override = unit_override

    @classmethod
    def to_proto(self, resource):
        if not resource:
            return None

        res = (
            dashboard_pb2.MonitoringBetaDashboardColumnLayoutColumnsWidgetsScorecardTimeSeriesQuery()
        )
        if DashboardColumnLayoutColumnsWidgetsScorecardTimeSeriesQueryTimeSeriesFilter.to_proto(
            resource.time_series_filter
        ):
            res.time_series_filter.CopyFrom(
                DashboardColumnLayoutColumnsWidgetsScorecardTimeSeriesQueryTimeSeriesFilter.to_proto(
                    resource.time_series_filter
                )
            )
        else:
            res.ClearField("time_series_filter")
        if DashboardColumnLayoutColumnsWidgetsScorecardTimeSeriesQueryTimeSeriesFilterRatio.to_proto(
            resource.time_series_filter_ratio
        ):
            res.time_series_filter_ratio.CopyFrom(
                DashboardColumnLayoutColumnsWidgetsScorecardTimeSeriesQueryTimeSeriesFilterRatio.to_proto(
                    resource.time_series_filter_ratio
                )
            )
        else:
            res.ClearField("time_series_filter_ratio")
        if Primitive.to_proto(resource.time_series_query_language):
            res.time_series_query_language = Primitive.to_proto(
                resource.time_series_query_language
            )
        if Primitive.to_proto(resource.unit_override):
            res.unit_override = Primitive.to_proto(resource.unit_override)
        return res

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return None

        return DashboardColumnLayoutColumnsWidgetsScorecardTimeSeriesQuery(
            time_series_filter=DashboardColumnLayoutColumnsWidgetsScorecardTimeSeriesQueryTimeSeriesFilter.from_proto(
                resource.time_series_filter
            ),
            time_series_filter_ratio=DashboardColumnLayoutColumnsWidgetsScorecardTimeSeriesQueryTimeSeriesFilterRatio.from_proto(
                resource.time_series_filter_ratio
            ),
            time_series_query_language=Primitive.from_proto(
                resource.time_series_query_language
            ),
            unit_override=Primitive.from_proto(resource.unit_override),
        )


class DashboardColumnLayoutColumnsWidgetsScorecardTimeSeriesQueryArray(object):
    @classmethod
    def to_proto(self, resources):
        if not resources:
            return resources
        return [
            DashboardColumnLayoutColumnsWidgetsScorecardTimeSeriesQuery.to_proto(i)
            for i in resources
        ]

    @classmethod
    def from_proto(self, resources):
        return [
            DashboardColumnLayoutColumnsWidgetsScorecardTimeSeriesQuery.from_proto(i)
            for i in resources
        ]


class DashboardColumnLayoutColumnsWidgetsScorecardTimeSeriesQueryTimeSeriesFilter(
    object
):
    def __init__(
        self,
        filter: str = None,
        aggregation: dict = None,
        secondary_aggregation: dict = None,
        pick_time_series_filter: dict = None,
    ):
        self.filter = filter
        self.aggregation = aggregation
        self.secondary_aggregation = secondary_aggregation
        self.pick_time_series_filter = pick_time_series_filter

    @classmethod
    def to_proto(self, resource):
        if not resource:
            return None

        res = (
            dashboard_pb2.MonitoringBetaDashboardColumnLayoutColumnsWidgetsScorecardTimeSeriesQueryTimeSeriesFilter()
        )
        if Primitive.to_proto(resource.filter):
            res.filter = Primitive.to_proto(resource.filter)
        if DashboardColumnLayoutColumnsWidgetsScorecardTimeSeriesQueryTimeSeriesFilterAggregation.to_proto(
            resource.aggregation
        ):
            res.aggregation.CopyFrom(
                DashboardColumnLayoutColumnsWidgetsScorecardTimeSeriesQueryTimeSeriesFilterAggregation.to_proto(
                    resource.aggregation
                )
            )
        else:
            res.ClearField("aggregation")
        if DashboardColumnLayoutColumnsWidgetsScorecardTimeSeriesQueryTimeSeriesFilterSecondaryAggregation.to_proto(
            resource.secondary_aggregation
        ):
            res.secondary_aggregation.CopyFrom(
                DashboardColumnLayoutColumnsWidgetsScorecardTimeSeriesQueryTimeSeriesFilterSecondaryAggregation.to_proto(
                    resource.secondary_aggregation
                )
            )
        else:
            res.ClearField("secondary_aggregation")
        if DashboardColumnLayoutColumnsWidgetsScorecardTimeSeriesQueryTimeSeriesFilterPickTimeSeriesFilter.to_proto(
            resource.pick_time_series_filter
        ):
            res.pick_time_series_filter.CopyFrom(
                DashboardColumnLayoutColumnsWidgetsScorecardTimeSeriesQueryTimeSeriesFilterPickTimeSeriesFilter.to_proto(
                    resource.pick_time_series_filter
                )
            )
        else:
            res.ClearField("pick_time_series_filter")
        return res

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return None

        return DashboardColumnLayoutColumnsWidgetsScorecardTimeSeriesQueryTimeSeriesFilter(
            filter=Primitive.from_proto(resource.filter),
            aggregation=DashboardColumnLayoutColumnsWidgetsScorecardTimeSeriesQueryTimeSeriesFilterAggregation.from_proto(
                resource.aggregation
            ),
            secondary_aggregation=DashboardColumnLayoutColumnsWidgetsScorecardTimeSeriesQueryTimeSeriesFilterSecondaryAggregation.from_proto(
                resource.secondary_aggregation
            ),
            pick_time_series_filter=DashboardColumnLayoutColumnsWidgetsScorecardTimeSeriesQueryTimeSeriesFilterPickTimeSeriesFilter.from_proto(
                resource.pick_time_series_filter
            ),
        )


class DashboardColumnLayoutColumnsWidgetsScorecardTimeSeriesQueryTimeSeriesFilterArray(
    object
):
    @classmethod
    def to_proto(self, resources):
        if not resources:
            return resources
        return [
            DashboardColumnLayoutColumnsWidgetsScorecardTimeSeriesQueryTimeSeriesFilter.to_proto(
                i
            )
            for i in resources
        ]

    @classmethod
    def from_proto(self, resources):
        return [
            DashboardColumnLayoutColumnsWidgetsScorecardTimeSeriesQueryTimeSeriesFilter.from_proto(
                i
            )
            for i in resources
        ]


class DashboardColumnLayoutColumnsWidgetsScorecardTimeSeriesQueryTimeSeriesFilterAggregation(
    object
):
    def __init__(
        self,
        alignment_period: str = None,
        per_series_aligner: str = None,
        cross_series_reducer: str = None,
        group_by_fields: list = None,
    ):
        self.alignment_period = alignment_period
        self.per_series_aligner = per_series_aligner
        self.cross_series_reducer = cross_series_reducer
        self.group_by_fields = group_by_fields

    @classmethod
    def to_proto(self, resource):
        if not resource:
            return None

        res = (
            dashboard_pb2.MonitoringBetaDashboardColumnLayoutColumnsWidgetsScorecardTimeSeriesQueryTimeSeriesFilterAggregation()
        )
        if Primitive.to_proto(resource.alignment_period):
            res.alignment_period = Primitive.to_proto(resource.alignment_period)
        if DashboardColumnLayoutColumnsWidgetsScorecardTimeSeriesQueryTimeSeriesFilterAggregationPerSeriesAlignerEnum.to_proto(
            resource.per_series_aligner
        ):
            res.per_series_aligner = DashboardColumnLayoutColumnsWidgetsScorecardTimeSeriesQueryTimeSeriesFilterAggregationPerSeriesAlignerEnum.to_proto(
                resource.per_series_aligner
            )
        if DashboardColumnLayoutColumnsWidgetsScorecardTimeSeriesQueryTimeSeriesFilterAggregationCrossSeriesReducerEnum.to_proto(
            resource.cross_series_reducer
        ):
            res.cross_series_reducer = DashboardColumnLayoutColumnsWidgetsScorecardTimeSeriesQueryTimeSeriesFilterAggregationCrossSeriesReducerEnum.to_proto(
                resource.cross_series_reducer
            )
        if Primitive.to_proto(resource.group_by_fields):
            res.group_by_fields.extend(Primitive.to_proto(resource.group_by_fields))
        return res

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return None

        return DashboardColumnLayoutColumnsWidgetsScorecardTimeSeriesQueryTimeSeriesFilterAggregation(
            alignment_period=Primitive.from_proto(resource.alignment_period),
            per_series_aligner=DashboardColumnLayoutColumnsWidgetsScorecardTimeSeriesQueryTimeSeriesFilterAggregationPerSeriesAlignerEnum.from_proto(
                resource.per_series_aligner
            ),
            cross_series_reducer=DashboardColumnLayoutColumnsWidgetsScorecardTimeSeriesQueryTimeSeriesFilterAggregationCrossSeriesReducerEnum.from_proto(
                resource.cross_series_reducer
            ),
            group_by_fields=Primitive.from_proto(resource.group_by_fields),
        )


class DashboardColumnLayoutColumnsWidgetsScorecardTimeSeriesQueryTimeSeriesFilterAggregationArray(
    object
):
    @classmethod
    def to_proto(self, resources):
        if not resources:
            return resources
        return [
            DashboardColumnLayoutColumnsWidgetsScorecardTimeSeriesQueryTimeSeriesFilterAggregation.to_proto(
                i
            )
            for i in resources
        ]

    @classmethod
    def from_proto(self, resources):
        return [
            DashboardColumnLayoutColumnsWidgetsScorecardTimeSeriesQueryTimeSeriesFilterAggregation.from_proto(
                i
            )
            for i in resources
        ]


class DashboardColumnLayoutColumnsWidgetsScorecardTimeSeriesQueryTimeSeriesFilterSecondaryAggregation(
    object
):
    def __init__(
        self,
        alignment_period: str = None,
        per_series_aligner: str = None,
        cross_series_reducer: str = None,
        group_by_fields: list = None,
    ):
        self.alignment_period = alignment_period
        self.per_series_aligner = per_series_aligner
        self.cross_series_reducer = cross_series_reducer
        self.group_by_fields = group_by_fields

    @classmethod
    def to_proto(self, resource):
        if not resource:
            return None

        res = (
            dashboard_pb2.MonitoringBetaDashboardColumnLayoutColumnsWidgetsScorecardTimeSeriesQueryTimeSeriesFilterSecondaryAggregation()
        )
        if Primitive.to_proto(resource.alignment_period):
            res.alignment_period = Primitive.to_proto(resource.alignment_period)
        if DashboardColumnLayoutColumnsWidgetsScorecardTimeSeriesQueryTimeSeriesFilterSecondaryAggregationPerSeriesAlignerEnum.to_proto(
            resource.per_series_aligner
        ):
            res.per_series_aligner = DashboardColumnLayoutColumnsWidgetsScorecardTimeSeriesQueryTimeSeriesFilterSecondaryAggregationPerSeriesAlignerEnum.to_proto(
                resource.per_series_aligner
            )
        if DashboardColumnLayoutColumnsWidgetsScorecardTimeSeriesQueryTimeSeriesFilterSecondaryAggregationCrossSeriesReducerEnum.to_proto(
            resource.cross_series_reducer
        ):
            res.cross_series_reducer = DashboardColumnLayoutColumnsWidgetsScorecardTimeSeriesQueryTimeSeriesFilterSecondaryAggregationCrossSeriesReducerEnum.to_proto(
                resource.cross_series_reducer
            )
        if Primitive.to_proto(resource.group_by_fields):
            res.group_by_fields.extend(Primitive.to_proto(resource.group_by_fields))
        return res

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return None

        return DashboardColumnLayoutColumnsWidgetsScorecardTimeSeriesQueryTimeSeriesFilterSecondaryAggregation(
            alignment_period=Primitive.from_proto(resource.alignment_period),
            per_series_aligner=DashboardColumnLayoutColumnsWidgetsScorecardTimeSeriesQueryTimeSeriesFilterSecondaryAggregationPerSeriesAlignerEnum.from_proto(
                resource.per_series_aligner
            ),
            cross_series_reducer=DashboardColumnLayoutColumnsWidgetsScorecardTimeSeriesQueryTimeSeriesFilterSecondaryAggregationCrossSeriesReducerEnum.from_proto(
                resource.cross_series_reducer
            ),
            group_by_fields=Primitive.from_proto(resource.group_by_fields),
        )


class DashboardColumnLayoutColumnsWidgetsScorecardTimeSeriesQueryTimeSeriesFilterSecondaryAggregationArray(
    object
):
    @classmethod
    def to_proto(self, resources):
        if not resources:
            return resources
        return [
            DashboardColumnLayoutColumnsWidgetsScorecardTimeSeriesQueryTimeSeriesFilterSecondaryAggregation.to_proto(
                i
            )
            for i in resources
        ]

    @classmethod
    def from_proto(self, resources):
        return [
            DashboardColumnLayoutColumnsWidgetsScorecardTimeSeriesQueryTimeSeriesFilterSecondaryAggregation.from_proto(
                i
            )
            for i in resources
        ]


class DashboardColumnLayoutColumnsWidgetsScorecardTimeSeriesQueryTimeSeriesFilterPickTimeSeriesFilter(
    object
):
    def __init__(
        self,
        ranking_method: str = None,
        num_time_series: int = None,
        direction: str = None,
    ):
        self.ranking_method = ranking_method
        self.num_time_series = num_time_series
        self.direction = direction

    @classmethod
    def to_proto(self, resource):
        if not resource:
            return None

        res = (
            dashboard_pb2.MonitoringBetaDashboardColumnLayoutColumnsWidgetsScorecardTimeSeriesQueryTimeSeriesFilterPickTimeSeriesFilter()
        )
        if DashboardColumnLayoutColumnsWidgetsScorecardTimeSeriesQueryTimeSeriesFilterPickTimeSeriesFilterRankingMethodEnum.to_proto(
            resource.ranking_method
        ):
            res.ranking_method = DashboardColumnLayoutColumnsWidgetsScorecardTimeSeriesQueryTimeSeriesFilterPickTimeSeriesFilterRankingMethodEnum.to_proto(
                resource.ranking_method
            )
        if Primitive.to_proto(resource.num_time_series):
            res.num_time_series = Primitive.to_proto(resource.num_time_series)
        if DashboardColumnLayoutColumnsWidgetsScorecardTimeSeriesQueryTimeSeriesFilterPickTimeSeriesFilterDirectionEnum.to_proto(
            resource.direction
        ):
            res.direction = DashboardColumnLayoutColumnsWidgetsScorecardTimeSeriesQueryTimeSeriesFilterPickTimeSeriesFilterDirectionEnum.to_proto(
                resource.direction
            )
        return res

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return None

        return DashboardColumnLayoutColumnsWidgetsScorecardTimeSeriesQueryTimeSeriesFilterPickTimeSeriesFilter(
            ranking_method=DashboardColumnLayoutColumnsWidgetsScorecardTimeSeriesQueryTimeSeriesFilterPickTimeSeriesFilterRankingMethodEnum.from_proto(
                resource.ranking_method
            ),
            num_time_series=Primitive.from_proto(resource.num_time_series),
            direction=DashboardColumnLayoutColumnsWidgetsScorecardTimeSeriesQueryTimeSeriesFilterPickTimeSeriesFilterDirectionEnum.from_proto(
                resource.direction
            ),
        )


class DashboardColumnLayoutColumnsWidgetsScorecardTimeSeriesQueryTimeSeriesFilterPickTimeSeriesFilterArray(
    object
):
    @classmethod
    def to_proto(self, resources):
        if not resources:
            return resources
        return [
            DashboardColumnLayoutColumnsWidgetsScorecardTimeSeriesQueryTimeSeriesFilterPickTimeSeriesFilter.to_proto(
                i
            )
            for i in resources
        ]

    @classmethod
    def from_proto(self, resources):
        return [
            DashboardColumnLayoutColumnsWidgetsScorecardTimeSeriesQueryTimeSeriesFilterPickTimeSeriesFilter.from_proto(
                i
            )
            for i in resources
        ]


class DashboardColumnLayoutColumnsWidgetsScorecardTimeSeriesQueryTimeSeriesFilterRatio(
    object
):
    def __init__(
        self,
        numerator: dict = None,
        denominator: dict = None,
        secondary_aggregation: dict = None,
        pick_time_series_filter: dict = None,
    ):
        self.numerator = numerator
        self.denominator = denominator
        self.secondary_aggregation = secondary_aggregation
        self.pick_time_series_filter = pick_time_series_filter

    @classmethod
    def to_proto(self, resource):
        if not resource:
            return None

        res = (
            dashboard_pb2.MonitoringBetaDashboardColumnLayoutColumnsWidgetsScorecardTimeSeriesQueryTimeSeriesFilterRatio()
        )
        if DashboardColumnLayoutColumnsWidgetsScorecardTimeSeriesQueryTimeSeriesFilterRatioNumerator.to_proto(
            resource.numerator
        ):
            res.numerator.CopyFrom(
                DashboardColumnLayoutColumnsWidgetsScorecardTimeSeriesQueryTimeSeriesFilterRatioNumerator.to_proto(
                    resource.numerator
                )
            )
        else:
            res.ClearField("numerator")
        if DashboardColumnLayoutColumnsWidgetsScorecardTimeSeriesQueryTimeSeriesFilterRatioDenominator.to_proto(
            resource.denominator
        ):
            res.denominator.CopyFrom(
                DashboardColumnLayoutColumnsWidgetsScorecardTimeSeriesQueryTimeSeriesFilterRatioDenominator.to_proto(
                    resource.denominator
                )
            )
        else:
            res.ClearField("denominator")
        if DashboardColumnLayoutColumnsWidgetsScorecardTimeSeriesQueryTimeSeriesFilterRatioSecondaryAggregation.to_proto(
            resource.secondary_aggregation
        ):
            res.secondary_aggregation.CopyFrom(
                DashboardColumnLayoutColumnsWidgetsScorecardTimeSeriesQueryTimeSeriesFilterRatioSecondaryAggregation.to_proto(
                    resource.secondary_aggregation
                )
            )
        else:
            res.ClearField("secondary_aggregation")
        if DashboardColumnLayoutColumnsWidgetsScorecardTimeSeriesQueryTimeSeriesFilterRatioPickTimeSeriesFilter.to_proto(
            resource.pick_time_series_filter
        ):
            res.pick_time_series_filter.CopyFrom(
                DashboardColumnLayoutColumnsWidgetsScorecardTimeSeriesQueryTimeSeriesFilterRatioPickTimeSeriesFilter.to_proto(
                    resource.pick_time_series_filter
                )
            )
        else:
            res.ClearField("pick_time_series_filter")
        return res

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return None

        return DashboardColumnLayoutColumnsWidgetsScorecardTimeSeriesQueryTimeSeriesFilterRatio(
            numerator=DashboardColumnLayoutColumnsWidgetsScorecardTimeSeriesQueryTimeSeriesFilterRatioNumerator.from_proto(
                resource.numerator
            ),
            denominator=DashboardColumnLayoutColumnsWidgetsScorecardTimeSeriesQueryTimeSeriesFilterRatioDenominator.from_proto(
                resource.denominator
            ),
            secondary_aggregation=DashboardColumnLayoutColumnsWidgetsScorecardTimeSeriesQueryTimeSeriesFilterRatioSecondaryAggregation.from_proto(
                resource.secondary_aggregation
            ),
            pick_time_series_filter=DashboardColumnLayoutColumnsWidgetsScorecardTimeSeriesQueryTimeSeriesFilterRatioPickTimeSeriesFilter.from_proto(
                resource.pick_time_series_filter
            ),
        )


class DashboardColumnLayoutColumnsWidgetsScorecardTimeSeriesQueryTimeSeriesFilterRatioArray(
    object
):
    @classmethod
    def to_proto(self, resources):
        if not resources:
            return resources
        return [
            DashboardColumnLayoutColumnsWidgetsScorecardTimeSeriesQueryTimeSeriesFilterRatio.to_proto(
                i
            )
            for i in resources
        ]

    @classmethod
    def from_proto(self, resources):
        return [
            DashboardColumnLayoutColumnsWidgetsScorecardTimeSeriesQueryTimeSeriesFilterRatio.from_proto(
                i
            )
            for i in resources
        ]


class DashboardColumnLayoutColumnsWidgetsScorecardTimeSeriesQueryTimeSeriesFilterRatioNumerator(
    object
):
    def __init__(self, filter: str = None, aggregation: dict = None):
        self.filter = filter
        self.aggregation = aggregation

    @classmethod
    def to_proto(self, resource):
        if not resource:
            return None

        res = (
            dashboard_pb2.MonitoringBetaDashboardColumnLayoutColumnsWidgetsScorecardTimeSeriesQueryTimeSeriesFilterRatioNumerator()
        )
        if Primitive.to_proto(resource.filter):
            res.filter = Primitive.to_proto(resource.filter)
        if DashboardColumnLayoutColumnsWidgetsScorecardTimeSeriesQueryTimeSeriesFilterRatioNumeratorAggregation.to_proto(
            resource.aggregation
        ):
            res.aggregation.CopyFrom(
                DashboardColumnLayoutColumnsWidgetsScorecardTimeSeriesQueryTimeSeriesFilterRatioNumeratorAggregation.to_proto(
                    resource.aggregation
                )
            )
        else:
            res.ClearField("aggregation")
        return res

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return None

        return DashboardColumnLayoutColumnsWidgetsScorecardTimeSeriesQueryTimeSeriesFilterRatioNumerator(
            filter=Primitive.from_proto(resource.filter),
            aggregation=DashboardColumnLayoutColumnsWidgetsScorecardTimeSeriesQueryTimeSeriesFilterRatioNumeratorAggregation.from_proto(
                resource.aggregation
            ),
        )


class DashboardColumnLayoutColumnsWidgetsScorecardTimeSeriesQueryTimeSeriesFilterRatioNumeratorArray(
    object
):
    @classmethod
    def to_proto(self, resources):
        if not resources:
            return resources
        return [
            DashboardColumnLayoutColumnsWidgetsScorecardTimeSeriesQueryTimeSeriesFilterRatioNumerator.to_proto(
                i
            )
            for i in resources
        ]

    @classmethod
    def from_proto(self, resources):
        return [
            DashboardColumnLayoutColumnsWidgetsScorecardTimeSeriesQueryTimeSeriesFilterRatioNumerator.from_proto(
                i
            )
            for i in resources
        ]


class DashboardColumnLayoutColumnsWidgetsScorecardTimeSeriesQueryTimeSeriesFilterRatioNumeratorAggregation(
    object
):
    def __init__(
        self,
        alignment_period: str = None,
        per_series_aligner: str = None,
        cross_series_reducer: str = None,
        group_by_fields: list = None,
    ):
        self.alignment_period = alignment_period
        self.per_series_aligner = per_series_aligner
        self.cross_series_reducer = cross_series_reducer
        self.group_by_fields = group_by_fields

    @classmethod
    def to_proto(self, resource):
        if not resource:
            return None

        res = (
            dashboard_pb2.MonitoringBetaDashboardColumnLayoutColumnsWidgetsScorecardTimeSeriesQueryTimeSeriesFilterRatioNumeratorAggregation()
        )
        if Primitive.to_proto(resource.alignment_period):
            res.alignment_period = Primitive.to_proto(resource.alignment_period)
        if DashboardColumnLayoutColumnsWidgetsScorecardTimeSeriesQueryTimeSeriesFilterRatioNumeratorAggregationPerSeriesAlignerEnum.to_proto(
            resource.per_series_aligner
        ):
            res.per_series_aligner = DashboardColumnLayoutColumnsWidgetsScorecardTimeSeriesQueryTimeSeriesFilterRatioNumeratorAggregationPerSeriesAlignerEnum.to_proto(
                resource.per_series_aligner
            )
        if DashboardColumnLayoutColumnsWidgetsScorecardTimeSeriesQueryTimeSeriesFilterRatioNumeratorAggregationCrossSeriesReducerEnum.to_proto(
            resource.cross_series_reducer
        ):
            res.cross_series_reducer = DashboardColumnLayoutColumnsWidgetsScorecardTimeSeriesQueryTimeSeriesFilterRatioNumeratorAggregationCrossSeriesReducerEnum.to_proto(
                resource.cross_series_reducer
            )
        if Primitive.to_proto(resource.group_by_fields):
            res.group_by_fields.extend(Primitive.to_proto(resource.group_by_fields))
        return res

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return None

        return DashboardColumnLayoutColumnsWidgetsScorecardTimeSeriesQueryTimeSeriesFilterRatioNumeratorAggregation(
            alignment_period=Primitive.from_proto(resource.alignment_period),
            per_series_aligner=DashboardColumnLayoutColumnsWidgetsScorecardTimeSeriesQueryTimeSeriesFilterRatioNumeratorAggregationPerSeriesAlignerEnum.from_proto(
                resource.per_series_aligner
            ),
            cross_series_reducer=DashboardColumnLayoutColumnsWidgetsScorecardTimeSeriesQueryTimeSeriesFilterRatioNumeratorAggregationCrossSeriesReducerEnum.from_proto(
                resource.cross_series_reducer
            ),
            group_by_fields=Primitive.from_proto(resource.group_by_fields),
        )


class DashboardColumnLayoutColumnsWidgetsScorecardTimeSeriesQueryTimeSeriesFilterRatioNumeratorAggregationArray(
    object
):
    @classmethod
    def to_proto(self, resources):
        if not resources:
            return resources
        return [
            DashboardColumnLayoutColumnsWidgetsScorecardTimeSeriesQueryTimeSeriesFilterRatioNumeratorAggregation.to_proto(
                i
            )
            for i in resources
        ]

    @classmethod
    def from_proto(self, resources):
        return [
            DashboardColumnLayoutColumnsWidgetsScorecardTimeSeriesQueryTimeSeriesFilterRatioNumeratorAggregation.from_proto(
                i
            )
            for i in resources
        ]


class DashboardColumnLayoutColumnsWidgetsScorecardTimeSeriesQueryTimeSeriesFilterRatioDenominator(
    object
):
    def __init__(self, filter: str = None, aggregation: dict = None):
        self.filter = filter
        self.aggregation = aggregation

    @classmethod
    def to_proto(self, resource):
        if not resource:
            return None

        res = (
            dashboard_pb2.MonitoringBetaDashboardColumnLayoutColumnsWidgetsScorecardTimeSeriesQueryTimeSeriesFilterRatioDenominator()
        )
        if Primitive.to_proto(resource.filter):
            res.filter = Primitive.to_proto(resource.filter)
        if DashboardColumnLayoutColumnsWidgetsScorecardTimeSeriesQueryTimeSeriesFilterRatioDenominatorAggregation.to_proto(
            resource.aggregation
        ):
            res.aggregation.CopyFrom(
                DashboardColumnLayoutColumnsWidgetsScorecardTimeSeriesQueryTimeSeriesFilterRatioDenominatorAggregation.to_proto(
                    resource.aggregation
                )
            )
        else:
            res.ClearField("aggregation")
        return res

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return None

        return DashboardColumnLayoutColumnsWidgetsScorecardTimeSeriesQueryTimeSeriesFilterRatioDenominator(
            filter=Primitive.from_proto(resource.filter),
            aggregation=DashboardColumnLayoutColumnsWidgetsScorecardTimeSeriesQueryTimeSeriesFilterRatioDenominatorAggregation.from_proto(
                resource.aggregation
            ),
        )


class DashboardColumnLayoutColumnsWidgetsScorecardTimeSeriesQueryTimeSeriesFilterRatioDenominatorArray(
    object
):
    @classmethod
    def to_proto(self, resources):
        if not resources:
            return resources
        return [
            DashboardColumnLayoutColumnsWidgetsScorecardTimeSeriesQueryTimeSeriesFilterRatioDenominator.to_proto(
                i
            )
            for i in resources
        ]

    @classmethod
    def from_proto(self, resources):
        return [
            DashboardColumnLayoutColumnsWidgetsScorecardTimeSeriesQueryTimeSeriesFilterRatioDenominator.from_proto(
                i
            )
            for i in resources
        ]


class DashboardColumnLayoutColumnsWidgetsScorecardTimeSeriesQueryTimeSeriesFilterRatioDenominatorAggregation(
    object
):
    def __init__(
        self,
        alignment_period: str = None,
        per_series_aligner: str = None,
        cross_series_reducer: str = None,
        group_by_fields: list = None,
    ):
        self.alignment_period = alignment_period
        self.per_series_aligner = per_series_aligner
        self.cross_series_reducer = cross_series_reducer
        self.group_by_fields = group_by_fields

    @classmethod
    def to_proto(self, resource):
        if not resource:
            return None

        res = (
            dashboard_pb2.MonitoringBetaDashboardColumnLayoutColumnsWidgetsScorecardTimeSeriesQueryTimeSeriesFilterRatioDenominatorAggregation()
        )
        if Primitive.to_proto(resource.alignment_period):
            res.alignment_period = Primitive.to_proto(resource.alignment_period)
        if DashboardColumnLayoutColumnsWidgetsScorecardTimeSeriesQueryTimeSeriesFilterRatioDenominatorAggregationPerSeriesAlignerEnum.to_proto(
            resource.per_series_aligner
        ):
            res.per_series_aligner = DashboardColumnLayoutColumnsWidgetsScorecardTimeSeriesQueryTimeSeriesFilterRatioDenominatorAggregationPerSeriesAlignerEnum.to_proto(
                resource.per_series_aligner
            )
        if DashboardColumnLayoutColumnsWidgetsScorecardTimeSeriesQueryTimeSeriesFilterRatioDenominatorAggregationCrossSeriesReducerEnum.to_proto(
            resource.cross_series_reducer
        ):
            res.cross_series_reducer = DashboardColumnLayoutColumnsWidgetsScorecardTimeSeriesQueryTimeSeriesFilterRatioDenominatorAggregationCrossSeriesReducerEnum.to_proto(
                resource.cross_series_reducer
            )
        if Primitive.to_proto(resource.group_by_fields):
            res.group_by_fields.extend(Primitive.to_proto(resource.group_by_fields))
        return res

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return None

        return DashboardColumnLayoutColumnsWidgetsScorecardTimeSeriesQueryTimeSeriesFilterRatioDenominatorAggregation(
            alignment_period=Primitive.from_proto(resource.alignment_period),
            per_series_aligner=DashboardColumnLayoutColumnsWidgetsScorecardTimeSeriesQueryTimeSeriesFilterRatioDenominatorAggregationPerSeriesAlignerEnum.from_proto(
                resource.per_series_aligner
            ),
            cross_series_reducer=DashboardColumnLayoutColumnsWidgetsScorecardTimeSeriesQueryTimeSeriesFilterRatioDenominatorAggregationCrossSeriesReducerEnum.from_proto(
                resource.cross_series_reducer
            ),
            group_by_fields=Primitive.from_proto(resource.group_by_fields),
        )


class DashboardColumnLayoutColumnsWidgetsScorecardTimeSeriesQueryTimeSeriesFilterRatioDenominatorAggregationArray(
    object
):
    @classmethod
    def to_proto(self, resources):
        if not resources:
            return resources
        return [
            DashboardColumnLayoutColumnsWidgetsScorecardTimeSeriesQueryTimeSeriesFilterRatioDenominatorAggregation.to_proto(
                i
            )
            for i in resources
        ]

    @classmethod
    def from_proto(self, resources):
        return [
            DashboardColumnLayoutColumnsWidgetsScorecardTimeSeriesQueryTimeSeriesFilterRatioDenominatorAggregation.from_proto(
                i
            )
            for i in resources
        ]


class DashboardColumnLayoutColumnsWidgetsScorecardTimeSeriesQueryTimeSeriesFilterRatioSecondaryAggregation(
    object
):
    def __init__(
        self,
        alignment_period: str = None,
        per_series_aligner: str = None,
        cross_series_reducer: str = None,
        group_by_fields: list = None,
    ):
        self.alignment_period = alignment_period
        self.per_series_aligner = per_series_aligner
        self.cross_series_reducer = cross_series_reducer
        self.group_by_fields = group_by_fields

    @classmethod
    def to_proto(self, resource):
        if not resource:
            return None

        res = (
            dashboard_pb2.MonitoringBetaDashboardColumnLayoutColumnsWidgetsScorecardTimeSeriesQueryTimeSeriesFilterRatioSecondaryAggregation()
        )
        if Primitive.to_proto(resource.alignment_period):
            res.alignment_period = Primitive.to_proto(resource.alignment_period)
        if DashboardColumnLayoutColumnsWidgetsScorecardTimeSeriesQueryTimeSeriesFilterRatioSecondaryAggregationPerSeriesAlignerEnum.to_proto(
            resource.per_series_aligner
        ):
            res.per_series_aligner = DashboardColumnLayoutColumnsWidgetsScorecardTimeSeriesQueryTimeSeriesFilterRatioSecondaryAggregationPerSeriesAlignerEnum.to_proto(
                resource.per_series_aligner
            )
        if DashboardColumnLayoutColumnsWidgetsScorecardTimeSeriesQueryTimeSeriesFilterRatioSecondaryAggregationCrossSeriesReducerEnum.to_proto(
            resource.cross_series_reducer
        ):
            res.cross_series_reducer = DashboardColumnLayoutColumnsWidgetsScorecardTimeSeriesQueryTimeSeriesFilterRatioSecondaryAggregationCrossSeriesReducerEnum.to_proto(
                resource.cross_series_reducer
            )
        if Primitive.to_proto(resource.group_by_fields):
            res.group_by_fields.extend(Primitive.to_proto(resource.group_by_fields))
        return res

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return None

        return DashboardColumnLayoutColumnsWidgetsScorecardTimeSeriesQueryTimeSeriesFilterRatioSecondaryAggregation(
            alignment_period=Primitive.from_proto(resource.alignment_period),
            per_series_aligner=DashboardColumnLayoutColumnsWidgetsScorecardTimeSeriesQueryTimeSeriesFilterRatioSecondaryAggregationPerSeriesAlignerEnum.from_proto(
                resource.per_series_aligner
            ),
            cross_series_reducer=DashboardColumnLayoutColumnsWidgetsScorecardTimeSeriesQueryTimeSeriesFilterRatioSecondaryAggregationCrossSeriesReducerEnum.from_proto(
                resource.cross_series_reducer
            ),
            group_by_fields=Primitive.from_proto(resource.group_by_fields),
        )


class DashboardColumnLayoutColumnsWidgetsScorecardTimeSeriesQueryTimeSeriesFilterRatioSecondaryAggregationArray(
    object
):
    @classmethod
    def to_proto(self, resources):
        if not resources:
            return resources
        return [
            DashboardColumnLayoutColumnsWidgetsScorecardTimeSeriesQueryTimeSeriesFilterRatioSecondaryAggregation.to_proto(
                i
            )
            for i in resources
        ]

    @classmethod
    def from_proto(self, resources):
        return [
            DashboardColumnLayoutColumnsWidgetsScorecardTimeSeriesQueryTimeSeriesFilterRatioSecondaryAggregation.from_proto(
                i
            )
            for i in resources
        ]


class DashboardColumnLayoutColumnsWidgetsScorecardTimeSeriesQueryTimeSeriesFilterRatioPickTimeSeriesFilter(
    object
):
    def __init__(
        self,
        ranking_method: str = None,
        num_time_series: int = None,
        direction: str = None,
    ):
        self.ranking_method = ranking_method
        self.num_time_series = num_time_series
        self.direction = direction

    @classmethod
    def to_proto(self, resource):
        if not resource:
            return None

        res = (
            dashboard_pb2.MonitoringBetaDashboardColumnLayoutColumnsWidgetsScorecardTimeSeriesQueryTimeSeriesFilterRatioPickTimeSeriesFilter()
        )
        if DashboardColumnLayoutColumnsWidgetsScorecardTimeSeriesQueryTimeSeriesFilterRatioPickTimeSeriesFilterRankingMethodEnum.to_proto(
            resource.ranking_method
        ):
            res.ranking_method = DashboardColumnLayoutColumnsWidgetsScorecardTimeSeriesQueryTimeSeriesFilterRatioPickTimeSeriesFilterRankingMethodEnum.to_proto(
                resource.ranking_method
            )
        if Primitive.to_proto(resource.num_time_series):
            res.num_time_series = Primitive.to_proto(resource.num_time_series)
        if DashboardColumnLayoutColumnsWidgetsScorecardTimeSeriesQueryTimeSeriesFilterRatioPickTimeSeriesFilterDirectionEnum.to_proto(
            resource.direction
        ):
            res.direction = DashboardColumnLayoutColumnsWidgetsScorecardTimeSeriesQueryTimeSeriesFilterRatioPickTimeSeriesFilterDirectionEnum.to_proto(
                resource.direction
            )
        return res

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return None

        return DashboardColumnLayoutColumnsWidgetsScorecardTimeSeriesQueryTimeSeriesFilterRatioPickTimeSeriesFilter(
            ranking_method=DashboardColumnLayoutColumnsWidgetsScorecardTimeSeriesQueryTimeSeriesFilterRatioPickTimeSeriesFilterRankingMethodEnum.from_proto(
                resource.ranking_method
            ),
            num_time_series=Primitive.from_proto(resource.num_time_series),
            direction=DashboardColumnLayoutColumnsWidgetsScorecardTimeSeriesQueryTimeSeriesFilterRatioPickTimeSeriesFilterDirectionEnum.from_proto(
                resource.direction
            ),
        )


class DashboardColumnLayoutColumnsWidgetsScorecardTimeSeriesQueryTimeSeriesFilterRatioPickTimeSeriesFilterArray(
    object
):
    @classmethod
    def to_proto(self, resources):
        if not resources:
            return resources
        return [
            DashboardColumnLayoutColumnsWidgetsScorecardTimeSeriesQueryTimeSeriesFilterRatioPickTimeSeriesFilter.to_proto(
                i
            )
            for i in resources
        ]

    @classmethod
    def from_proto(self, resources):
        return [
            DashboardColumnLayoutColumnsWidgetsScorecardTimeSeriesQueryTimeSeriesFilterRatioPickTimeSeriesFilter.from_proto(
                i
            )
            for i in resources
        ]


class DashboardColumnLayoutColumnsWidgetsScorecardGaugeView(object):
    def __init__(self, lower_bound: float = None, upper_bound: float = None):
        self.lower_bound = lower_bound
        self.upper_bound = upper_bound

    @classmethod
    def to_proto(self, resource):
        if not resource:
            return None

        res = (
            dashboard_pb2.MonitoringBetaDashboardColumnLayoutColumnsWidgetsScorecardGaugeView()
        )
        if Primitive.to_proto(resource.lower_bound):
            res.lower_bound = Primitive.to_proto(resource.lower_bound)
        if Primitive.to_proto(resource.upper_bound):
            res.upper_bound = Primitive.to_proto(resource.upper_bound)
        return res

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return None

        return DashboardColumnLayoutColumnsWidgetsScorecardGaugeView(
            lower_bound=Primitive.from_proto(resource.lower_bound),
            upper_bound=Primitive.from_proto(resource.upper_bound),
        )


class DashboardColumnLayoutColumnsWidgetsScorecardGaugeViewArray(object):
    @classmethod
    def to_proto(self, resources):
        if not resources:
            return resources
        return [
            DashboardColumnLayoutColumnsWidgetsScorecardGaugeView.to_proto(i)
            for i in resources
        ]

    @classmethod
    def from_proto(self, resources):
        return [
            DashboardColumnLayoutColumnsWidgetsScorecardGaugeView.from_proto(i)
            for i in resources
        ]


class DashboardColumnLayoutColumnsWidgetsScorecardSparkChartView(object):
    def __init__(self, spark_chart_type: str = None, min_alignment_period: str = None):
        self.spark_chart_type = spark_chart_type
        self.min_alignment_period = min_alignment_period

    @classmethod
    def to_proto(self, resource):
        if not resource:
            return None

        res = (
            dashboard_pb2.MonitoringBetaDashboardColumnLayoutColumnsWidgetsScorecardSparkChartView()
        )
        if DashboardColumnLayoutColumnsWidgetsScorecardSparkChartViewSparkChartTypeEnum.to_proto(
            resource.spark_chart_type
        ):
            res.spark_chart_type = DashboardColumnLayoutColumnsWidgetsScorecardSparkChartViewSparkChartTypeEnum.to_proto(
                resource.spark_chart_type
            )
        if Primitive.to_proto(resource.min_alignment_period):
            res.min_alignment_period = Primitive.to_proto(resource.min_alignment_period)
        return res

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return None

        return DashboardColumnLayoutColumnsWidgetsScorecardSparkChartView(
            spark_chart_type=DashboardColumnLayoutColumnsWidgetsScorecardSparkChartViewSparkChartTypeEnum.from_proto(
                resource.spark_chart_type
            ),
            min_alignment_period=Primitive.from_proto(resource.min_alignment_period),
        )


class DashboardColumnLayoutColumnsWidgetsScorecardSparkChartViewArray(object):
    @classmethod
    def to_proto(self, resources):
        if not resources:
            return resources
        return [
            DashboardColumnLayoutColumnsWidgetsScorecardSparkChartView.to_proto(i)
            for i in resources
        ]

    @classmethod
    def from_proto(self, resources):
        return [
            DashboardColumnLayoutColumnsWidgetsScorecardSparkChartView.from_proto(i)
            for i in resources
        ]


class DashboardColumnLayoutColumnsWidgetsScorecardThresholds(object):
    def __init__(
        self,
        label: str = None,
        value: float = None,
        color: str = None,
        direction: str = None,
    ):
        self.label = label
        self.value = value
        self.color = color
        self.direction = direction

    @classmethod
    def to_proto(self, resource):
        if not resource:
            return None

        res = (
            dashboard_pb2.MonitoringBetaDashboardColumnLayoutColumnsWidgetsScorecardThresholds()
        )
        if Primitive.to_proto(resource.label):
            res.label = Primitive.to_proto(resource.label)
        if Primitive.to_proto(resource.value):
            res.value = Primitive.to_proto(resource.value)
        if DashboardColumnLayoutColumnsWidgetsScorecardThresholdsColorEnum.to_proto(
            resource.color
        ):
            res.color = DashboardColumnLayoutColumnsWidgetsScorecardThresholdsColorEnum.to_proto(
                resource.color
            )
        if DashboardColumnLayoutColumnsWidgetsScorecardThresholdsDirectionEnum.to_proto(
            resource.direction
        ):
            res.direction = DashboardColumnLayoutColumnsWidgetsScorecardThresholdsDirectionEnum.to_proto(
                resource.direction
            )
        return res

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return None

        return DashboardColumnLayoutColumnsWidgetsScorecardThresholds(
            label=Primitive.from_proto(resource.label),
            value=Primitive.from_proto(resource.value),
            color=DashboardColumnLayoutColumnsWidgetsScorecardThresholdsColorEnum.from_proto(
                resource.color
            ),
            direction=DashboardColumnLayoutColumnsWidgetsScorecardThresholdsDirectionEnum.from_proto(
                resource.direction
            ),
        )


class DashboardColumnLayoutColumnsWidgetsScorecardThresholdsArray(object):
    @classmethod
    def to_proto(self, resources):
        if not resources:
            return resources
        return [
            DashboardColumnLayoutColumnsWidgetsScorecardThresholds.to_proto(i)
            for i in resources
        ]

    @classmethod
    def from_proto(self, resources):
        return [
            DashboardColumnLayoutColumnsWidgetsScorecardThresholds.from_proto(i)
            for i in resources
        ]


class DashboardColumnLayoutColumnsWidgetsText(object):
    def __init__(self, content: str = None, format: str = None):
        self.content = content
        self.format = format

    @classmethod
    def to_proto(self, resource):
        if not resource:
            return None

        res = dashboard_pb2.MonitoringBetaDashboardColumnLayoutColumnsWidgetsText()
        if Primitive.to_proto(resource.content):
            res.content = Primitive.to_proto(resource.content)
        if DashboardColumnLayoutColumnsWidgetsTextFormatEnum.to_proto(resource.format):
            res.format = DashboardColumnLayoutColumnsWidgetsTextFormatEnum.to_proto(
                resource.format
            )
        return res

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return None

        return DashboardColumnLayoutColumnsWidgetsText(
            content=Primitive.from_proto(resource.content),
            format=DashboardColumnLayoutColumnsWidgetsTextFormatEnum.from_proto(
                resource.format
            ),
        )


class DashboardColumnLayoutColumnsWidgetsTextArray(object):
    @classmethod
    def to_proto(self, resources):
        if not resources:
            return resources
        return [DashboardColumnLayoutColumnsWidgetsText.to_proto(i) for i in resources]

    @classmethod
    def from_proto(self, resources):
        return [
            DashboardColumnLayoutColumnsWidgetsText.from_proto(i) for i in resources
        ]


class DashboardColumnLayoutColumnsWidgetsBlank(object):
    @classmethod
    def to_proto(self, resource):
        if not resource:
            return None

        res = dashboard_pb2.MonitoringBetaDashboardColumnLayoutColumnsWidgetsBlank()
        return res

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return None

        return DashboardColumnLayoutColumnsWidgetsBlank()


class DashboardColumnLayoutColumnsWidgetsBlankArray(object):
    @classmethod
    def to_proto(self, resources):
        if not resources:
            return resources
        return [DashboardColumnLayoutColumnsWidgetsBlank.to_proto(i) for i in resources]

    @classmethod
    def from_proto(self, resources):
        return [
            DashboardColumnLayoutColumnsWidgetsBlank.from_proto(i) for i in resources
        ]


class DashboardColumnLayoutColumnsWidgetsLogsPanel(object):
    def __init__(self, filter: str = None, resource_names: list = None):
        self.filter = filter
        self.resource_names = resource_names

    @classmethod
    def to_proto(self, resource):
        if not resource:
            return None

        res = dashboard_pb2.MonitoringBetaDashboardColumnLayoutColumnsWidgetsLogsPanel()
        if Primitive.to_proto(resource.filter):
            res.filter = Primitive.to_proto(resource.filter)
        if Primitive.to_proto(resource.resource_names):
            res.resource_names.extend(Primitive.to_proto(resource.resource_names))
        return res

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return None

        return DashboardColumnLayoutColumnsWidgetsLogsPanel(
            filter=Primitive.from_proto(resource.filter),
            resource_names=Primitive.from_proto(resource.resource_names),
        )


class DashboardColumnLayoutColumnsWidgetsLogsPanelArray(object):
    @classmethod
    def to_proto(self, resources):
        if not resources:
            return resources
        return [
            DashboardColumnLayoutColumnsWidgetsLogsPanel.to_proto(i) for i in resources
        ]

    @classmethod
    def from_proto(self, resources):
        return [
            DashboardColumnLayoutColumnsWidgetsLogsPanel.from_proto(i)
            for i in resources
        ]


class DashboardGridLayoutWidgetsXyChartDataSetsTimeSeriesQueryTimeSeriesFilterAggregationPerSeriesAlignerEnum(
    object
):
    @classmethod
    def to_proto(self, resource):
        if not resource:
            return resource
        return dashboard_pb2.MonitoringBetaDashboardGridLayoutWidgetsXyChartDataSetsTimeSeriesQueryTimeSeriesFilterAggregationPerSeriesAlignerEnum.Value(
            "MonitoringBetaDashboardGridLayoutWidgetsXyChartDataSetsTimeSeriesQueryTimeSeriesFilterAggregationPerSeriesAlignerEnum%s"
            % resource
        )

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return resource
        return dashboard_pb2.MonitoringBetaDashboardGridLayoutWidgetsXyChartDataSetsTimeSeriesQueryTimeSeriesFilterAggregationPerSeriesAlignerEnum.Name(
            resource
        )[
            len(
                "MonitoringBetaDashboardGridLayoutWidgetsXyChartDataSetsTimeSeriesQueryTimeSeriesFilterAggregationPerSeriesAlignerEnum"
            ) :
        ]


class DashboardGridLayoutWidgetsXyChartDataSetsTimeSeriesQueryTimeSeriesFilterAggregationCrossSeriesReducerEnum(
    object
):
    @classmethod
    def to_proto(self, resource):
        if not resource:
            return resource
        return dashboard_pb2.MonitoringBetaDashboardGridLayoutWidgetsXyChartDataSetsTimeSeriesQueryTimeSeriesFilterAggregationCrossSeriesReducerEnum.Value(
            "MonitoringBetaDashboardGridLayoutWidgetsXyChartDataSetsTimeSeriesQueryTimeSeriesFilterAggregationCrossSeriesReducerEnum%s"
            % resource
        )

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return resource
        return dashboard_pb2.MonitoringBetaDashboardGridLayoutWidgetsXyChartDataSetsTimeSeriesQueryTimeSeriesFilterAggregationCrossSeriesReducerEnum.Name(
            resource
        )[
            len(
                "MonitoringBetaDashboardGridLayoutWidgetsXyChartDataSetsTimeSeriesQueryTimeSeriesFilterAggregationCrossSeriesReducerEnum"
            ) :
        ]


class DashboardGridLayoutWidgetsXyChartDataSetsTimeSeriesQueryTimeSeriesFilterSecondaryAggregationPerSeriesAlignerEnum(
    object
):
    @classmethod
    def to_proto(self, resource):
        if not resource:
            return resource
        return dashboard_pb2.MonitoringBetaDashboardGridLayoutWidgetsXyChartDataSetsTimeSeriesQueryTimeSeriesFilterSecondaryAggregationPerSeriesAlignerEnum.Value(
            "MonitoringBetaDashboardGridLayoutWidgetsXyChartDataSetsTimeSeriesQueryTimeSeriesFilterSecondaryAggregationPerSeriesAlignerEnum%s"
            % resource
        )

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return resource
        return dashboard_pb2.MonitoringBetaDashboardGridLayoutWidgetsXyChartDataSetsTimeSeriesQueryTimeSeriesFilterSecondaryAggregationPerSeriesAlignerEnum.Name(
            resource
        )[
            len(
                "MonitoringBetaDashboardGridLayoutWidgetsXyChartDataSetsTimeSeriesQueryTimeSeriesFilterSecondaryAggregationPerSeriesAlignerEnum"
            ) :
        ]


class DashboardGridLayoutWidgetsXyChartDataSetsTimeSeriesQueryTimeSeriesFilterSecondaryAggregationCrossSeriesReducerEnum(
    object
):
    @classmethod
    def to_proto(self, resource):
        if not resource:
            return resource
        return dashboard_pb2.MonitoringBetaDashboardGridLayoutWidgetsXyChartDataSetsTimeSeriesQueryTimeSeriesFilterSecondaryAggregationCrossSeriesReducerEnum.Value(
            "MonitoringBetaDashboardGridLayoutWidgetsXyChartDataSetsTimeSeriesQueryTimeSeriesFilterSecondaryAggregationCrossSeriesReducerEnum%s"
            % resource
        )

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return resource
        return dashboard_pb2.MonitoringBetaDashboardGridLayoutWidgetsXyChartDataSetsTimeSeriesQueryTimeSeriesFilterSecondaryAggregationCrossSeriesReducerEnum.Name(
            resource
        )[
            len(
                "MonitoringBetaDashboardGridLayoutWidgetsXyChartDataSetsTimeSeriesQueryTimeSeriesFilterSecondaryAggregationCrossSeriesReducerEnum"
            ) :
        ]


class DashboardGridLayoutWidgetsXyChartDataSetsTimeSeriesQueryTimeSeriesFilterPickTimeSeriesFilterRankingMethodEnum(
    object
):
    @classmethod
    def to_proto(self, resource):
        if not resource:
            return resource
        return dashboard_pb2.MonitoringBetaDashboardGridLayoutWidgetsXyChartDataSetsTimeSeriesQueryTimeSeriesFilterPickTimeSeriesFilterRankingMethodEnum.Value(
            "MonitoringBetaDashboardGridLayoutWidgetsXyChartDataSetsTimeSeriesQueryTimeSeriesFilterPickTimeSeriesFilterRankingMethodEnum%s"
            % resource
        )

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return resource
        return dashboard_pb2.MonitoringBetaDashboardGridLayoutWidgetsXyChartDataSetsTimeSeriesQueryTimeSeriesFilterPickTimeSeriesFilterRankingMethodEnum.Name(
            resource
        )[
            len(
                "MonitoringBetaDashboardGridLayoutWidgetsXyChartDataSetsTimeSeriesQueryTimeSeriesFilterPickTimeSeriesFilterRankingMethodEnum"
            ) :
        ]


class DashboardGridLayoutWidgetsXyChartDataSetsTimeSeriesQueryTimeSeriesFilterPickTimeSeriesFilterDirectionEnum(
    object
):
    @classmethod
    def to_proto(self, resource):
        if not resource:
            return resource
        return dashboard_pb2.MonitoringBetaDashboardGridLayoutWidgetsXyChartDataSetsTimeSeriesQueryTimeSeriesFilterPickTimeSeriesFilterDirectionEnum.Value(
            "MonitoringBetaDashboardGridLayoutWidgetsXyChartDataSetsTimeSeriesQueryTimeSeriesFilterPickTimeSeriesFilterDirectionEnum%s"
            % resource
        )

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return resource
        return dashboard_pb2.MonitoringBetaDashboardGridLayoutWidgetsXyChartDataSetsTimeSeriesQueryTimeSeriesFilterPickTimeSeriesFilterDirectionEnum.Name(
            resource
        )[
            len(
                "MonitoringBetaDashboardGridLayoutWidgetsXyChartDataSetsTimeSeriesQueryTimeSeriesFilterPickTimeSeriesFilterDirectionEnum"
            ) :
        ]


class DashboardGridLayoutWidgetsXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioNumeratorAggregationPerSeriesAlignerEnum(
    object
):
    @classmethod
    def to_proto(self, resource):
        if not resource:
            return resource
        return dashboard_pb2.MonitoringBetaDashboardGridLayoutWidgetsXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioNumeratorAggregationPerSeriesAlignerEnum.Value(
            "MonitoringBetaDashboardGridLayoutWidgetsXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioNumeratorAggregationPerSeriesAlignerEnum%s"
            % resource
        )

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return resource
        return dashboard_pb2.MonitoringBetaDashboardGridLayoutWidgetsXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioNumeratorAggregationPerSeriesAlignerEnum.Name(
            resource
        )[
            len(
                "MonitoringBetaDashboardGridLayoutWidgetsXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioNumeratorAggregationPerSeriesAlignerEnum"
            ) :
        ]


class DashboardGridLayoutWidgetsXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioNumeratorAggregationCrossSeriesReducerEnum(
    object
):
    @classmethod
    def to_proto(self, resource):
        if not resource:
            return resource
        return dashboard_pb2.MonitoringBetaDashboardGridLayoutWidgetsXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioNumeratorAggregationCrossSeriesReducerEnum.Value(
            "MonitoringBetaDashboardGridLayoutWidgetsXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioNumeratorAggregationCrossSeriesReducerEnum%s"
            % resource
        )

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return resource
        return dashboard_pb2.MonitoringBetaDashboardGridLayoutWidgetsXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioNumeratorAggregationCrossSeriesReducerEnum.Name(
            resource
        )[
            len(
                "MonitoringBetaDashboardGridLayoutWidgetsXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioNumeratorAggregationCrossSeriesReducerEnum"
            ) :
        ]


class DashboardGridLayoutWidgetsXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioDenominatorAggregationPerSeriesAlignerEnum(
    object
):
    @classmethod
    def to_proto(self, resource):
        if not resource:
            return resource
        return dashboard_pb2.MonitoringBetaDashboardGridLayoutWidgetsXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioDenominatorAggregationPerSeriesAlignerEnum.Value(
            "MonitoringBetaDashboardGridLayoutWidgetsXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioDenominatorAggregationPerSeriesAlignerEnum%s"
            % resource
        )

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return resource
        return dashboard_pb2.MonitoringBetaDashboardGridLayoutWidgetsXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioDenominatorAggregationPerSeriesAlignerEnum.Name(
            resource
        )[
            len(
                "MonitoringBetaDashboardGridLayoutWidgetsXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioDenominatorAggregationPerSeriesAlignerEnum"
            ) :
        ]


class DashboardGridLayoutWidgetsXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioDenominatorAggregationCrossSeriesReducerEnum(
    object
):
    @classmethod
    def to_proto(self, resource):
        if not resource:
            return resource
        return dashboard_pb2.MonitoringBetaDashboardGridLayoutWidgetsXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioDenominatorAggregationCrossSeriesReducerEnum.Value(
            "MonitoringBetaDashboardGridLayoutWidgetsXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioDenominatorAggregationCrossSeriesReducerEnum%s"
            % resource
        )

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return resource
        return dashboard_pb2.MonitoringBetaDashboardGridLayoutWidgetsXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioDenominatorAggregationCrossSeriesReducerEnum.Name(
            resource
        )[
            len(
                "MonitoringBetaDashboardGridLayoutWidgetsXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioDenominatorAggregationCrossSeriesReducerEnum"
            ) :
        ]


class DashboardGridLayoutWidgetsXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioSecondaryAggregationPerSeriesAlignerEnum(
    object
):
    @classmethod
    def to_proto(self, resource):
        if not resource:
            return resource
        return dashboard_pb2.MonitoringBetaDashboardGridLayoutWidgetsXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioSecondaryAggregationPerSeriesAlignerEnum.Value(
            "MonitoringBetaDashboardGridLayoutWidgetsXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioSecondaryAggregationPerSeriesAlignerEnum%s"
            % resource
        )

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return resource
        return dashboard_pb2.MonitoringBetaDashboardGridLayoutWidgetsXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioSecondaryAggregationPerSeriesAlignerEnum.Name(
            resource
        )[
            len(
                "MonitoringBetaDashboardGridLayoutWidgetsXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioSecondaryAggregationPerSeriesAlignerEnum"
            ) :
        ]


class DashboardGridLayoutWidgetsXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioSecondaryAggregationCrossSeriesReducerEnum(
    object
):
    @classmethod
    def to_proto(self, resource):
        if not resource:
            return resource
        return dashboard_pb2.MonitoringBetaDashboardGridLayoutWidgetsXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioSecondaryAggregationCrossSeriesReducerEnum.Value(
            "MonitoringBetaDashboardGridLayoutWidgetsXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioSecondaryAggregationCrossSeriesReducerEnum%s"
            % resource
        )

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return resource
        return dashboard_pb2.MonitoringBetaDashboardGridLayoutWidgetsXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioSecondaryAggregationCrossSeriesReducerEnum.Name(
            resource
        )[
            len(
                "MonitoringBetaDashboardGridLayoutWidgetsXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioSecondaryAggregationCrossSeriesReducerEnum"
            ) :
        ]


class DashboardGridLayoutWidgetsXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioPickTimeSeriesFilterRankingMethodEnum(
    object
):
    @classmethod
    def to_proto(self, resource):
        if not resource:
            return resource
        return dashboard_pb2.MonitoringBetaDashboardGridLayoutWidgetsXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioPickTimeSeriesFilterRankingMethodEnum.Value(
            "MonitoringBetaDashboardGridLayoutWidgetsXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioPickTimeSeriesFilterRankingMethodEnum%s"
            % resource
        )

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return resource
        return dashboard_pb2.MonitoringBetaDashboardGridLayoutWidgetsXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioPickTimeSeriesFilterRankingMethodEnum.Name(
            resource
        )[
            len(
                "MonitoringBetaDashboardGridLayoutWidgetsXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioPickTimeSeriesFilterRankingMethodEnum"
            ) :
        ]


class DashboardGridLayoutWidgetsXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioPickTimeSeriesFilterDirectionEnum(
    object
):
    @classmethod
    def to_proto(self, resource):
        if not resource:
            return resource
        return dashboard_pb2.MonitoringBetaDashboardGridLayoutWidgetsXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioPickTimeSeriesFilterDirectionEnum.Value(
            "MonitoringBetaDashboardGridLayoutWidgetsXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioPickTimeSeriesFilterDirectionEnum%s"
            % resource
        )

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return resource
        return dashboard_pb2.MonitoringBetaDashboardGridLayoutWidgetsXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioPickTimeSeriesFilterDirectionEnum.Name(
            resource
        )[
            len(
                "MonitoringBetaDashboardGridLayoutWidgetsXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioPickTimeSeriesFilterDirectionEnum"
            ) :
        ]


class DashboardGridLayoutWidgetsXyChartDataSetsPlotTypeEnum(object):
    @classmethod
    def to_proto(self, resource):
        if not resource:
            return resource
        return dashboard_pb2.MonitoringBetaDashboardGridLayoutWidgetsXyChartDataSetsPlotTypeEnum.Value(
            "MonitoringBetaDashboardGridLayoutWidgetsXyChartDataSetsPlotTypeEnum%s"
            % resource
        )

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return resource
        return dashboard_pb2.MonitoringBetaDashboardGridLayoutWidgetsXyChartDataSetsPlotTypeEnum.Name(
            resource
        )[
            len("MonitoringBetaDashboardGridLayoutWidgetsXyChartDataSetsPlotTypeEnum") :
        ]


class DashboardGridLayoutWidgetsXyChartThresholdsColorEnum(object):
    @classmethod
    def to_proto(self, resource):
        if not resource:
            return resource
        return dashboard_pb2.MonitoringBetaDashboardGridLayoutWidgetsXyChartThresholdsColorEnum.Value(
            "MonitoringBetaDashboardGridLayoutWidgetsXyChartThresholdsColorEnum%s"
            % resource
        )

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return resource
        return dashboard_pb2.MonitoringBetaDashboardGridLayoutWidgetsXyChartThresholdsColorEnum.Name(
            resource
        )[
            len("MonitoringBetaDashboardGridLayoutWidgetsXyChartThresholdsColorEnum") :
        ]


class DashboardGridLayoutWidgetsXyChartThresholdsDirectionEnum(object):
    @classmethod
    def to_proto(self, resource):
        if not resource:
            return resource
        return dashboard_pb2.MonitoringBetaDashboardGridLayoutWidgetsXyChartThresholdsDirectionEnum.Value(
            "MonitoringBetaDashboardGridLayoutWidgetsXyChartThresholdsDirectionEnum%s"
            % resource
        )

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return resource
        return dashboard_pb2.MonitoringBetaDashboardGridLayoutWidgetsXyChartThresholdsDirectionEnum.Name(
            resource
        )[
            len(
                "MonitoringBetaDashboardGridLayoutWidgetsXyChartThresholdsDirectionEnum"
            ) :
        ]


class DashboardGridLayoutWidgetsXyChartXAxisScaleEnum(object):
    @classmethod
    def to_proto(self, resource):
        if not resource:
            return resource
        return dashboard_pb2.MonitoringBetaDashboardGridLayoutWidgetsXyChartXAxisScaleEnum.Value(
            "MonitoringBetaDashboardGridLayoutWidgetsXyChartXAxisScaleEnum%s" % resource
        )

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return resource
        return dashboard_pb2.MonitoringBetaDashboardGridLayoutWidgetsXyChartXAxisScaleEnum.Name(
            resource
        )[
            len("MonitoringBetaDashboardGridLayoutWidgetsXyChartXAxisScaleEnum") :
        ]


class DashboardGridLayoutWidgetsXyChartYAxisScaleEnum(object):
    @classmethod
    def to_proto(self, resource):
        if not resource:
            return resource
        return dashboard_pb2.MonitoringBetaDashboardGridLayoutWidgetsXyChartYAxisScaleEnum.Value(
            "MonitoringBetaDashboardGridLayoutWidgetsXyChartYAxisScaleEnum%s" % resource
        )

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return resource
        return dashboard_pb2.MonitoringBetaDashboardGridLayoutWidgetsXyChartYAxisScaleEnum.Name(
            resource
        )[
            len("MonitoringBetaDashboardGridLayoutWidgetsXyChartYAxisScaleEnum") :
        ]


class DashboardGridLayoutWidgetsXyChartChartOptionsModeEnum(object):
    @classmethod
    def to_proto(self, resource):
        if not resource:
            return resource
        return dashboard_pb2.MonitoringBetaDashboardGridLayoutWidgetsXyChartChartOptionsModeEnum.Value(
            "MonitoringBetaDashboardGridLayoutWidgetsXyChartChartOptionsModeEnum%s"
            % resource
        )

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return resource
        return dashboard_pb2.MonitoringBetaDashboardGridLayoutWidgetsXyChartChartOptionsModeEnum.Name(
            resource
        )[
            len("MonitoringBetaDashboardGridLayoutWidgetsXyChartChartOptionsModeEnum") :
        ]


class DashboardGridLayoutWidgetsScorecardTimeSeriesQueryTimeSeriesFilterAggregationPerSeriesAlignerEnum(
    object
):
    @classmethod
    def to_proto(self, resource):
        if not resource:
            return resource
        return dashboard_pb2.MonitoringBetaDashboardGridLayoutWidgetsScorecardTimeSeriesQueryTimeSeriesFilterAggregationPerSeriesAlignerEnum.Value(
            "MonitoringBetaDashboardGridLayoutWidgetsScorecardTimeSeriesQueryTimeSeriesFilterAggregationPerSeriesAlignerEnum%s"
            % resource
        )

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return resource
        return dashboard_pb2.MonitoringBetaDashboardGridLayoutWidgetsScorecardTimeSeriesQueryTimeSeriesFilterAggregationPerSeriesAlignerEnum.Name(
            resource
        )[
            len(
                "MonitoringBetaDashboardGridLayoutWidgetsScorecardTimeSeriesQueryTimeSeriesFilterAggregationPerSeriesAlignerEnum"
            ) :
        ]


class DashboardGridLayoutWidgetsScorecardTimeSeriesQueryTimeSeriesFilterAggregationCrossSeriesReducerEnum(
    object
):
    @classmethod
    def to_proto(self, resource):
        if not resource:
            return resource
        return dashboard_pb2.MonitoringBetaDashboardGridLayoutWidgetsScorecardTimeSeriesQueryTimeSeriesFilterAggregationCrossSeriesReducerEnum.Value(
            "MonitoringBetaDashboardGridLayoutWidgetsScorecardTimeSeriesQueryTimeSeriesFilterAggregationCrossSeriesReducerEnum%s"
            % resource
        )

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return resource
        return dashboard_pb2.MonitoringBetaDashboardGridLayoutWidgetsScorecardTimeSeriesQueryTimeSeriesFilterAggregationCrossSeriesReducerEnum.Name(
            resource
        )[
            len(
                "MonitoringBetaDashboardGridLayoutWidgetsScorecardTimeSeriesQueryTimeSeriesFilterAggregationCrossSeriesReducerEnum"
            ) :
        ]


class DashboardGridLayoutWidgetsScorecardTimeSeriesQueryTimeSeriesFilterSecondaryAggregationPerSeriesAlignerEnum(
    object
):
    @classmethod
    def to_proto(self, resource):
        if not resource:
            return resource
        return dashboard_pb2.MonitoringBetaDashboardGridLayoutWidgetsScorecardTimeSeriesQueryTimeSeriesFilterSecondaryAggregationPerSeriesAlignerEnum.Value(
            "MonitoringBetaDashboardGridLayoutWidgetsScorecardTimeSeriesQueryTimeSeriesFilterSecondaryAggregationPerSeriesAlignerEnum%s"
            % resource
        )

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return resource
        return dashboard_pb2.MonitoringBetaDashboardGridLayoutWidgetsScorecardTimeSeriesQueryTimeSeriesFilterSecondaryAggregationPerSeriesAlignerEnum.Name(
            resource
        )[
            len(
                "MonitoringBetaDashboardGridLayoutWidgetsScorecardTimeSeriesQueryTimeSeriesFilterSecondaryAggregationPerSeriesAlignerEnum"
            ) :
        ]


class DashboardGridLayoutWidgetsScorecardTimeSeriesQueryTimeSeriesFilterSecondaryAggregationCrossSeriesReducerEnum(
    object
):
    @classmethod
    def to_proto(self, resource):
        if not resource:
            return resource
        return dashboard_pb2.MonitoringBetaDashboardGridLayoutWidgetsScorecardTimeSeriesQueryTimeSeriesFilterSecondaryAggregationCrossSeriesReducerEnum.Value(
            "MonitoringBetaDashboardGridLayoutWidgetsScorecardTimeSeriesQueryTimeSeriesFilterSecondaryAggregationCrossSeriesReducerEnum%s"
            % resource
        )

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return resource
        return dashboard_pb2.MonitoringBetaDashboardGridLayoutWidgetsScorecardTimeSeriesQueryTimeSeriesFilterSecondaryAggregationCrossSeriesReducerEnum.Name(
            resource
        )[
            len(
                "MonitoringBetaDashboardGridLayoutWidgetsScorecardTimeSeriesQueryTimeSeriesFilterSecondaryAggregationCrossSeriesReducerEnum"
            ) :
        ]


class DashboardGridLayoutWidgetsScorecardTimeSeriesQueryTimeSeriesFilterPickTimeSeriesFilterRankingMethodEnum(
    object
):
    @classmethod
    def to_proto(self, resource):
        if not resource:
            return resource
        return dashboard_pb2.MonitoringBetaDashboardGridLayoutWidgetsScorecardTimeSeriesQueryTimeSeriesFilterPickTimeSeriesFilterRankingMethodEnum.Value(
            "MonitoringBetaDashboardGridLayoutWidgetsScorecardTimeSeriesQueryTimeSeriesFilterPickTimeSeriesFilterRankingMethodEnum%s"
            % resource
        )

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return resource
        return dashboard_pb2.MonitoringBetaDashboardGridLayoutWidgetsScorecardTimeSeriesQueryTimeSeriesFilterPickTimeSeriesFilterRankingMethodEnum.Name(
            resource
        )[
            len(
                "MonitoringBetaDashboardGridLayoutWidgetsScorecardTimeSeriesQueryTimeSeriesFilterPickTimeSeriesFilterRankingMethodEnum"
            ) :
        ]


class DashboardGridLayoutWidgetsScorecardTimeSeriesQueryTimeSeriesFilterPickTimeSeriesFilterDirectionEnum(
    object
):
    @classmethod
    def to_proto(self, resource):
        if not resource:
            return resource
        return dashboard_pb2.MonitoringBetaDashboardGridLayoutWidgetsScorecardTimeSeriesQueryTimeSeriesFilterPickTimeSeriesFilterDirectionEnum.Value(
            "MonitoringBetaDashboardGridLayoutWidgetsScorecardTimeSeriesQueryTimeSeriesFilterPickTimeSeriesFilterDirectionEnum%s"
            % resource
        )

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return resource
        return dashboard_pb2.MonitoringBetaDashboardGridLayoutWidgetsScorecardTimeSeriesQueryTimeSeriesFilterPickTimeSeriesFilterDirectionEnum.Name(
            resource
        )[
            len(
                "MonitoringBetaDashboardGridLayoutWidgetsScorecardTimeSeriesQueryTimeSeriesFilterPickTimeSeriesFilterDirectionEnum"
            ) :
        ]


class DashboardGridLayoutWidgetsScorecardTimeSeriesQueryTimeSeriesFilterRatioNumeratorAggregationPerSeriesAlignerEnum(
    object
):
    @classmethod
    def to_proto(self, resource):
        if not resource:
            return resource
        return dashboard_pb2.MonitoringBetaDashboardGridLayoutWidgetsScorecardTimeSeriesQueryTimeSeriesFilterRatioNumeratorAggregationPerSeriesAlignerEnum.Value(
            "MonitoringBetaDashboardGridLayoutWidgetsScorecardTimeSeriesQueryTimeSeriesFilterRatioNumeratorAggregationPerSeriesAlignerEnum%s"
            % resource
        )

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return resource
        return dashboard_pb2.MonitoringBetaDashboardGridLayoutWidgetsScorecardTimeSeriesQueryTimeSeriesFilterRatioNumeratorAggregationPerSeriesAlignerEnum.Name(
            resource
        )[
            len(
                "MonitoringBetaDashboardGridLayoutWidgetsScorecardTimeSeriesQueryTimeSeriesFilterRatioNumeratorAggregationPerSeriesAlignerEnum"
            ) :
        ]


class DashboardGridLayoutWidgetsScorecardTimeSeriesQueryTimeSeriesFilterRatioNumeratorAggregationCrossSeriesReducerEnum(
    object
):
    @classmethod
    def to_proto(self, resource):
        if not resource:
            return resource
        return dashboard_pb2.MonitoringBetaDashboardGridLayoutWidgetsScorecardTimeSeriesQueryTimeSeriesFilterRatioNumeratorAggregationCrossSeriesReducerEnum.Value(
            "MonitoringBetaDashboardGridLayoutWidgetsScorecardTimeSeriesQueryTimeSeriesFilterRatioNumeratorAggregationCrossSeriesReducerEnum%s"
            % resource
        )

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return resource
        return dashboard_pb2.MonitoringBetaDashboardGridLayoutWidgetsScorecardTimeSeriesQueryTimeSeriesFilterRatioNumeratorAggregationCrossSeriesReducerEnum.Name(
            resource
        )[
            len(
                "MonitoringBetaDashboardGridLayoutWidgetsScorecardTimeSeriesQueryTimeSeriesFilterRatioNumeratorAggregationCrossSeriesReducerEnum"
            ) :
        ]


class DashboardGridLayoutWidgetsScorecardTimeSeriesQueryTimeSeriesFilterRatioDenominatorAggregationPerSeriesAlignerEnum(
    object
):
    @classmethod
    def to_proto(self, resource):
        if not resource:
            return resource
        return dashboard_pb2.MonitoringBetaDashboardGridLayoutWidgetsScorecardTimeSeriesQueryTimeSeriesFilterRatioDenominatorAggregationPerSeriesAlignerEnum.Value(
            "MonitoringBetaDashboardGridLayoutWidgetsScorecardTimeSeriesQueryTimeSeriesFilterRatioDenominatorAggregationPerSeriesAlignerEnum%s"
            % resource
        )

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return resource
        return dashboard_pb2.MonitoringBetaDashboardGridLayoutWidgetsScorecardTimeSeriesQueryTimeSeriesFilterRatioDenominatorAggregationPerSeriesAlignerEnum.Name(
            resource
        )[
            len(
                "MonitoringBetaDashboardGridLayoutWidgetsScorecardTimeSeriesQueryTimeSeriesFilterRatioDenominatorAggregationPerSeriesAlignerEnum"
            ) :
        ]


class DashboardGridLayoutWidgetsScorecardTimeSeriesQueryTimeSeriesFilterRatioDenominatorAggregationCrossSeriesReducerEnum(
    object
):
    @classmethod
    def to_proto(self, resource):
        if not resource:
            return resource
        return dashboard_pb2.MonitoringBetaDashboardGridLayoutWidgetsScorecardTimeSeriesQueryTimeSeriesFilterRatioDenominatorAggregationCrossSeriesReducerEnum.Value(
            "MonitoringBetaDashboardGridLayoutWidgetsScorecardTimeSeriesQueryTimeSeriesFilterRatioDenominatorAggregationCrossSeriesReducerEnum%s"
            % resource
        )

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return resource
        return dashboard_pb2.MonitoringBetaDashboardGridLayoutWidgetsScorecardTimeSeriesQueryTimeSeriesFilterRatioDenominatorAggregationCrossSeriesReducerEnum.Name(
            resource
        )[
            len(
                "MonitoringBetaDashboardGridLayoutWidgetsScorecardTimeSeriesQueryTimeSeriesFilterRatioDenominatorAggregationCrossSeriesReducerEnum"
            ) :
        ]


class DashboardGridLayoutWidgetsScorecardTimeSeriesQueryTimeSeriesFilterRatioSecondaryAggregationPerSeriesAlignerEnum(
    object
):
    @classmethod
    def to_proto(self, resource):
        if not resource:
            return resource
        return dashboard_pb2.MonitoringBetaDashboardGridLayoutWidgetsScorecardTimeSeriesQueryTimeSeriesFilterRatioSecondaryAggregationPerSeriesAlignerEnum.Value(
            "MonitoringBetaDashboardGridLayoutWidgetsScorecardTimeSeriesQueryTimeSeriesFilterRatioSecondaryAggregationPerSeriesAlignerEnum%s"
            % resource
        )

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return resource
        return dashboard_pb2.MonitoringBetaDashboardGridLayoutWidgetsScorecardTimeSeriesQueryTimeSeriesFilterRatioSecondaryAggregationPerSeriesAlignerEnum.Name(
            resource
        )[
            len(
                "MonitoringBetaDashboardGridLayoutWidgetsScorecardTimeSeriesQueryTimeSeriesFilterRatioSecondaryAggregationPerSeriesAlignerEnum"
            ) :
        ]


class DashboardGridLayoutWidgetsScorecardTimeSeriesQueryTimeSeriesFilterRatioSecondaryAggregationCrossSeriesReducerEnum(
    object
):
    @classmethod
    def to_proto(self, resource):
        if not resource:
            return resource
        return dashboard_pb2.MonitoringBetaDashboardGridLayoutWidgetsScorecardTimeSeriesQueryTimeSeriesFilterRatioSecondaryAggregationCrossSeriesReducerEnum.Value(
            "MonitoringBetaDashboardGridLayoutWidgetsScorecardTimeSeriesQueryTimeSeriesFilterRatioSecondaryAggregationCrossSeriesReducerEnum%s"
            % resource
        )

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return resource
        return dashboard_pb2.MonitoringBetaDashboardGridLayoutWidgetsScorecardTimeSeriesQueryTimeSeriesFilterRatioSecondaryAggregationCrossSeriesReducerEnum.Name(
            resource
        )[
            len(
                "MonitoringBetaDashboardGridLayoutWidgetsScorecardTimeSeriesQueryTimeSeriesFilterRatioSecondaryAggregationCrossSeriesReducerEnum"
            ) :
        ]


class DashboardGridLayoutWidgetsScorecardTimeSeriesQueryTimeSeriesFilterRatioPickTimeSeriesFilterRankingMethodEnum(
    object
):
    @classmethod
    def to_proto(self, resource):
        if not resource:
            return resource
        return dashboard_pb2.MonitoringBetaDashboardGridLayoutWidgetsScorecardTimeSeriesQueryTimeSeriesFilterRatioPickTimeSeriesFilterRankingMethodEnum.Value(
            "MonitoringBetaDashboardGridLayoutWidgetsScorecardTimeSeriesQueryTimeSeriesFilterRatioPickTimeSeriesFilterRankingMethodEnum%s"
            % resource
        )

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return resource
        return dashboard_pb2.MonitoringBetaDashboardGridLayoutWidgetsScorecardTimeSeriesQueryTimeSeriesFilterRatioPickTimeSeriesFilterRankingMethodEnum.Name(
            resource
        )[
            len(
                "MonitoringBetaDashboardGridLayoutWidgetsScorecardTimeSeriesQueryTimeSeriesFilterRatioPickTimeSeriesFilterRankingMethodEnum"
            ) :
        ]


class DashboardGridLayoutWidgetsScorecardTimeSeriesQueryTimeSeriesFilterRatioPickTimeSeriesFilterDirectionEnum(
    object
):
    @classmethod
    def to_proto(self, resource):
        if not resource:
            return resource
        return dashboard_pb2.MonitoringBetaDashboardGridLayoutWidgetsScorecardTimeSeriesQueryTimeSeriesFilterRatioPickTimeSeriesFilterDirectionEnum.Value(
            "MonitoringBetaDashboardGridLayoutWidgetsScorecardTimeSeriesQueryTimeSeriesFilterRatioPickTimeSeriesFilterDirectionEnum%s"
            % resource
        )

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return resource
        return dashboard_pb2.MonitoringBetaDashboardGridLayoutWidgetsScorecardTimeSeriesQueryTimeSeriesFilterRatioPickTimeSeriesFilterDirectionEnum.Name(
            resource
        )[
            len(
                "MonitoringBetaDashboardGridLayoutWidgetsScorecardTimeSeriesQueryTimeSeriesFilterRatioPickTimeSeriesFilterDirectionEnum"
            ) :
        ]


class DashboardGridLayoutWidgetsScorecardSparkChartViewSparkChartTypeEnum(object):
    @classmethod
    def to_proto(self, resource):
        if not resource:
            return resource
        return dashboard_pb2.MonitoringBetaDashboardGridLayoutWidgetsScorecardSparkChartViewSparkChartTypeEnum.Value(
            "MonitoringBetaDashboardGridLayoutWidgetsScorecardSparkChartViewSparkChartTypeEnum%s"
            % resource
        )

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return resource
        return dashboard_pb2.MonitoringBetaDashboardGridLayoutWidgetsScorecardSparkChartViewSparkChartTypeEnum.Name(
            resource
        )[
            len(
                "MonitoringBetaDashboardGridLayoutWidgetsScorecardSparkChartViewSparkChartTypeEnum"
            ) :
        ]


class DashboardGridLayoutWidgetsScorecardThresholdsColorEnum(object):
    @classmethod
    def to_proto(self, resource):
        if not resource:
            return resource
        return dashboard_pb2.MonitoringBetaDashboardGridLayoutWidgetsScorecardThresholdsColorEnum.Value(
            "MonitoringBetaDashboardGridLayoutWidgetsScorecardThresholdsColorEnum%s"
            % resource
        )

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return resource
        return dashboard_pb2.MonitoringBetaDashboardGridLayoutWidgetsScorecardThresholdsColorEnum.Name(
            resource
        )[
            len(
                "MonitoringBetaDashboardGridLayoutWidgetsScorecardThresholdsColorEnum"
            ) :
        ]


class DashboardGridLayoutWidgetsScorecardThresholdsDirectionEnum(object):
    @classmethod
    def to_proto(self, resource):
        if not resource:
            return resource
        return dashboard_pb2.MonitoringBetaDashboardGridLayoutWidgetsScorecardThresholdsDirectionEnum.Value(
            "MonitoringBetaDashboardGridLayoutWidgetsScorecardThresholdsDirectionEnum%s"
            % resource
        )

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return resource
        return dashboard_pb2.MonitoringBetaDashboardGridLayoutWidgetsScorecardThresholdsDirectionEnum.Name(
            resource
        )[
            len(
                "MonitoringBetaDashboardGridLayoutWidgetsScorecardThresholdsDirectionEnum"
            ) :
        ]


class DashboardGridLayoutWidgetsTextFormatEnum(object):
    @classmethod
    def to_proto(self, resource):
        if not resource:
            return resource
        return (
            dashboard_pb2.MonitoringBetaDashboardGridLayoutWidgetsTextFormatEnum.Value(
                "MonitoringBetaDashboardGridLayoutWidgetsTextFormatEnum%s" % resource
            )
        )

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return resource
        return (
            dashboard_pb2.MonitoringBetaDashboardGridLayoutWidgetsTextFormatEnum.Name(
                resource
            )[len("MonitoringBetaDashboardGridLayoutWidgetsTextFormatEnum") :]
        )


class DashboardMosaicLayoutTilesWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterAggregationPerSeriesAlignerEnum(
    object
):
    @classmethod
    def to_proto(self, resource):
        if not resource:
            return resource
        return dashboard_pb2.MonitoringBetaDashboardMosaicLayoutTilesWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterAggregationPerSeriesAlignerEnum.Value(
            "MonitoringBetaDashboardMosaicLayoutTilesWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterAggregationPerSeriesAlignerEnum%s"
            % resource
        )

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return resource
        return dashboard_pb2.MonitoringBetaDashboardMosaicLayoutTilesWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterAggregationPerSeriesAlignerEnum.Name(
            resource
        )[
            len(
                "MonitoringBetaDashboardMosaicLayoutTilesWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterAggregationPerSeriesAlignerEnum"
            ) :
        ]


class DashboardMosaicLayoutTilesWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterAggregationCrossSeriesReducerEnum(
    object
):
    @classmethod
    def to_proto(self, resource):
        if not resource:
            return resource
        return dashboard_pb2.MonitoringBetaDashboardMosaicLayoutTilesWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterAggregationCrossSeriesReducerEnum.Value(
            "MonitoringBetaDashboardMosaicLayoutTilesWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterAggregationCrossSeriesReducerEnum%s"
            % resource
        )

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return resource
        return dashboard_pb2.MonitoringBetaDashboardMosaicLayoutTilesWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterAggregationCrossSeriesReducerEnum.Name(
            resource
        )[
            len(
                "MonitoringBetaDashboardMosaicLayoutTilesWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterAggregationCrossSeriesReducerEnum"
            ) :
        ]


class DashboardMosaicLayoutTilesWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterSecondaryAggregationPerSeriesAlignerEnum(
    object
):
    @classmethod
    def to_proto(self, resource):
        if not resource:
            return resource
        return dashboard_pb2.MonitoringBetaDashboardMosaicLayoutTilesWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterSecondaryAggregationPerSeriesAlignerEnum.Value(
            "MonitoringBetaDashboardMosaicLayoutTilesWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterSecondaryAggregationPerSeriesAlignerEnum%s"
            % resource
        )

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return resource
        return dashboard_pb2.MonitoringBetaDashboardMosaicLayoutTilesWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterSecondaryAggregationPerSeriesAlignerEnum.Name(
            resource
        )[
            len(
                "MonitoringBetaDashboardMosaicLayoutTilesWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterSecondaryAggregationPerSeriesAlignerEnum"
            ) :
        ]


class DashboardMosaicLayoutTilesWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterSecondaryAggregationCrossSeriesReducerEnum(
    object
):
    @classmethod
    def to_proto(self, resource):
        if not resource:
            return resource
        return dashboard_pb2.MonitoringBetaDashboardMosaicLayoutTilesWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterSecondaryAggregationCrossSeriesReducerEnum.Value(
            "MonitoringBetaDashboardMosaicLayoutTilesWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterSecondaryAggregationCrossSeriesReducerEnum%s"
            % resource
        )

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return resource
        return dashboard_pb2.MonitoringBetaDashboardMosaicLayoutTilesWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterSecondaryAggregationCrossSeriesReducerEnum.Name(
            resource
        )[
            len(
                "MonitoringBetaDashboardMosaicLayoutTilesWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterSecondaryAggregationCrossSeriesReducerEnum"
            ) :
        ]


class DashboardMosaicLayoutTilesWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterPickTimeSeriesFilterRankingMethodEnum(
    object
):
    @classmethod
    def to_proto(self, resource):
        if not resource:
            return resource
        return dashboard_pb2.MonitoringBetaDashboardMosaicLayoutTilesWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterPickTimeSeriesFilterRankingMethodEnum.Value(
            "MonitoringBetaDashboardMosaicLayoutTilesWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterPickTimeSeriesFilterRankingMethodEnum%s"
            % resource
        )

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return resource
        return dashboard_pb2.MonitoringBetaDashboardMosaicLayoutTilesWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterPickTimeSeriesFilterRankingMethodEnum.Name(
            resource
        )[
            len(
                "MonitoringBetaDashboardMosaicLayoutTilesWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterPickTimeSeriesFilterRankingMethodEnum"
            ) :
        ]


class DashboardMosaicLayoutTilesWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterPickTimeSeriesFilterDirectionEnum(
    object
):
    @classmethod
    def to_proto(self, resource):
        if not resource:
            return resource
        return dashboard_pb2.MonitoringBetaDashboardMosaicLayoutTilesWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterPickTimeSeriesFilterDirectionEnum.Value(
            "MonitoringBetaDashboardMosaicLayoutTilesWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterPickTimeSeriesFilterDirectionEnum%s"
            % resource
        )

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return resource
        return dashboard_pb2.MonitoringBetaDashboardMosaicLayoutTilesWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterPickTimeSeriesFilterDirectionEnum.Name(
            resource
        )[
            len(
                "MonitoringBetaDashboardMosaicLayoutTilesWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterPickTimeSeriesFilterDirectionEnum"
            ) :
        ]


class DashboardMosaicLayoutTilesWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioNumeratorAggregationPerSeriesAlignerEnum(
    object
):
    @classmethod
    def to_proto(self, resource):
        if not resource:
            return resource
        return dashboard_pb2.MonitoringBetaDashboardMosaicLayoutTilesWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioNumeratorAggregationPerSeriesAlignerEnum.Value(
            "MonitoringBetaDashboardMosaicLayoutTilesWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioNumeratorAggregationPerSeriesAlignerEnum%s"
            % resource
        )

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return resource
        return dashboard_pb2.MonitoringBetaDashboardMosaicLayoutTilesWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioNumeratorAggregationPerSeriesAlignerEnum.Name(
            resource
        )[
            len(
                "MonitoringBetaDashboardMosaicLayoutTilesWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioNumeratorAggregationPerSeriesAlignerEnum"
            ) :
        ]


class DashboardMosaicLayoutTilesWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioNumeratorAggregationCrossSeriesReducerEnum(
    object
):
    @classmethod
    def to_proto(self, resource):
        if not resource:
            return resource
        return dashboard_pb2.MonitoringBetaDashboardMosaicLayoutTilesWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioNumeratorAggregationCrossSeriesReducerEnum.Value(
            "MonitoringBetaDashboardMosaicLayoutTilesWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioNumeratorAggregationCrossSeriesReducerEnum%s"
            % resource
        )

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return resource
        return dashboard_pb2.MonitoringBetaDashboardMosaicLayoutTilesWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioNumeratorAggregationCrossSeriesReducerEnum.Name(
            resource
        )[
            len(
                "MonitoringBetaDashboardMosaicLayoutTilesWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioNumeratorAggregationCrossSeriesReducerEnum"
            ) :
        ]


class DashboardMosaicLayoutTilesWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioDenominatorAggregationPerSeriesAlignerEnum(
    object
):
    @classmethod
    def to_proto(self, resource):
        if not resource:
            return resource
        return dashboard_pb2.MonitoringBetaDashboardMosaicLayoutTilesWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioDenominatorAggregationPerSeriesAlignerEnum.Value(
            "MonitoringBetaDashboardMosaicLayoutTilesWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioDenominatorAggregationPerSeriesAlignerEnum%s"
            % resource
        )

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return resource
        return dashboard_pb2.MonitoringBetaDashboardMosaicLayoutTilesWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioDenominatorAggregationPerSeriesAlignerEnum.Name(
            resource
        )[
            len(
                "MonitoringBetaDashboardMosaicLayoutTilesWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioDenominatorAggregationPerSeriesAlignerEnum"
            ) :
        ]


class DashboardMosaicLayoutTilesWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioDenominatorAggregationCrossSeriesReducerEnum(
    object
):
    @classmethod
    def to_proto(self, resource):
        if not resource:
            return resource
        return dashboard_pb2.MonitoringBetaDashboardMosaicLayoutTilesWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioDenominatorAggregationCrossSeriesReducerEnum.Value(
            "MonitoringBetaDashboardMosaicLayoutTilesWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioDenominatorAggregationCrossSeriesReducerEnum%s"
            % resource
        )

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return resource
        return dashboard_pb2.MonitoringBetaDashboardMosaicLayoutTilesWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioDenominatorAggregationCrossSeriesReducerEnum.Name(
            resource
        )[
            len(
                "MonitoringBetaDashboardMosaicLayoutTilesWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioDenominatorAggregationCrossSeriesReducerEnum"
            ) :
        ]


class DashboardMosaicLayoutTilesWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioSecondaryAggregationPerSeriesAlignerEnum(
    object
):
    @classmethod
    def to_proto(self, resource):
        if not resource:
            return resource
        return dashboard_pb2.MonitoringBetaDashboardMosaicLayoutTilesWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioSecondaryAggregationPerSeriesAlignerEnum.Value(
            "MonitoringBetaDashboardMosaicLayoutTilesWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioSecondaryAggregationPerSeriesAlignerEnum%s"
            % resource
        )

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return resource
        return dashboard_pb2.MonitoringBetaDashboardMosaicLayoutTilesWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioSecondaryAggregationPerSeriesAlignerEnum.Name(
            resource
        )[
            len(
                "MonitoringBetaDashboardMosaicLayoutTilesWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioSecondaryAggregationPerSeriesAlignerEnum"
            ) :
        ]


class DashboardMosaicLayoutTilesWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioSecondaryAggregationCrossSeriesReducerEnum(
    object
):
    @classmethod
    def to_proto(self, resource):
        if not resource:
            return resource
        return dashboard_pb2.MonitoringBetaDashboardMosaicLayoutTilesWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioSecondaryAggregationCrossSeriesReducerEnum.Value(
            "MonitoringBetaDashboardMosaicLayoutTilesWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioSecondaryAggregationCrossSeriesReducerEnum%s"
            % resource
        )

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return resource
        return dashboard_pb2.MonitoringBetaDashboardMosaicLayoutTilesWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioSecondaryAggregationCrossSeriesReducerEnum.Name(
            resource
        )[
            len(
                "MonitoringBetaDashboardMosaicLayoutTilesWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioSecondaryAggregationCrossSeriesReducerEnum"
            ) :
        ]


class DashboardMosaicLayoutTilesWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioPickTimeSeriesFilterRankingMethodEnum(
    object
):
    @classmethod
    def to_proto(self, resource):
        if not resource:
            return resource
        return dashboard_pb2.MonitoringBetaDashboardMosaicLayoutTilesWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioPickTimeSeriesFilterRankingMethodEnum.Value(
            "MonitoringBetaDashboardMosaicLayoutTilesWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioPickTimeSeriesFilterRankingMethodEnum%s"
            % resource
        )

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return resource
        return dashboard_pb2.MonitoringBetaDashboardMosaicLayoutTilesWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioPickTimeSeriesFilterRankingMethodEnum.Name(
            resource
        )[
            len(
                "MonitoringBetaDashboardMosaicLayoutTilesWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioPickTimeSeriesFilterRankingMethodEnum"
            ) :
        ]


class DashboardMosaicLayoutTilesWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioPickTimeSeriesFilterDirectionEnum(
    object
):
    @classmethod
    def to_proto(self, resource):
        if not resource:
            return resource
        return dashboard_pb2.MonitoringBetaDashboardMosaicLayoutTilesWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioPickTimeSeriesFilterDirectionEnum.Value(
            "MonitoringBetaDashboardMosaicLayoutTilesWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioPickTimeSeriesFilterDirectionEnum%s"
            % resource
        )

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return resource
        return dashboard_pb2.MonitoringBetaDashboardMosaicLayoutTilesWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioPickTimeSeriesFilterDirectionEnum.Name(
            resource
        )[
            len(
                "MonitoringBetaDashboardMosaicLayoutTilesWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioPickTimeSeriesFilterDirectionEnum"
            ) :
        ]


class DashboardMosaicLayoutTilesWidgetXyChartDataSetsPlotTypeEnum(object):
    @classmethod
    def to_proto(self, resource):
        if not resource:
            return resource
        return dashboard_pb2.MonitoringBetaDashboardMosaicLayoutTilesWidgetXyChartDataSetsPlotTypeEnum.Value(
            "MonitoringBetaDashboardMosaicLayoutTilesWidgetXyChartDataSetsPlotTypeEnum%s"
            % resource
        )

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return resource
        return dashboard_pb2.MonitoringBetaDashboardMosaicLayoutTilesWidgetXyChartDataSetsPlotTypeEnum.Name(
            resource
        )[
            len(
                "MonitoringBetaDashboardMosaicLayoutTilesWidgetXyChartDataSetsPlotTypeEnum"
            ) :
        ]


class DashboardMosaicLayoutTilesWidgetXyChartThresholdsColorEnum(object):
    @classmethod
    def to_proto(self, resource):
        if not resource:
            return resource
        return dashboard_pb2.MonitoringBetaDashboardMosaicLayoutTilesWidgetXyChartThresholdsColorEnum.Value(
            "MonitoringBetaDashboardMosaicLayoutTilesWidgetXyChartThresholdsColorEnum%s"
            % resource
        )

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return resource
        return dashboard_pb2.MonitoringBetaDashboardMosaicLayoutTilesWidgetXyChartThresholdsColorEnum.Name(
            resource
        )[
            len(
                "MonitoringBetaDashboardMosaicLayoutTilesWidgetXyChartThresholdsColorEnum"
            ) :
        ]


class DashboardMosaicLayoutTilesWidgetXyChartThresholdsDirectionEnum(object):
    @classmethod
    def to_proto(self, resource):
        if not resource:
            return resource
        return dashboard_pb2.MonitoringBetaDashboardMosaicLayoutTilesWidgetXyChartThresholdsDirectionEnum.Value(
            "MonitoringBetaDashboardMosaicLayoutTilesWidgetXyChartThresholdsDirectionEnum%s"
            % resource
        )

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return resource
        return dashboard_pb2.MonitoringBetaDashboardMosaicLayoutTilesWidgetXyChartThresholdsDirectionEnum.Name(
            resource
        )[
            len(
                "MonitoringBetaDashboardMosaicLayoutTilesWidgetXyChartThresholdsDirectionEnum"
            ) :
        ]


class DashboardMosaicLayoutTilesWidgetXyChartXAxisScaleEnum(object):
    @classmethod
    def to_proto(self, resource):
        if not resource:
            return resource
        return dashboard_pb2.MonitoringBetaDashboardMosaicLayoutTilesWidgetXyChartXAxisScaleEnum.Value(
            "MonitoringBetaDashboardMosaicLayoutTilesWidgetXyChartXAxisScaleEnum%s"
            % resource
        )

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return resource
        return dashboard_pb2.MonitoringBetaDashboardMosaicLayoutTilesWidgetXyChartXAxisScaleEnum.Name(
            resource
        )[
            len("MonitoringBetaDashboardMosaicLayoutTilesWidgetXyChartXAxisScaleEnum") :
        ]


class DashboardMosaicLayoutTilesWidgetXyChartYAxisScaleEnum(object):
    @classmethod
    def to_proto(self, resource):
        if not resource:
            return resource
        return dashboard_pb2.MonitoringBetaDashboardMosaicLayoutTilesWidgetXyChartYAxisScaleEnum.Value(
            "MonitoringBetaDashboardMosaicLayoutTilesWidgetXyChartYAxisScaleEnum%s"
            % resource
        )

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return resource
        return dashboard_pb2.MonitoringBetaDashboardMosaicLayoutTilesWidgetXyChartYAxisScaleEnum.Name(
            resource
        )[
            len("MonitoringBetaDashboardMosaicLayoutTilesWidgetXyChartYAxisScaleEnum") :
        ]


class DashboardMosaicLayoutTilesWidgetXyChartChartOptionsModeEnum(object):
    @classmethod
    def to_proto(self, resource):
        if not resource:
            return resource
        return dashboard_pb2.MonitoringBetaDashboardMosaicLayoutTilesWidgetXyChartChartOptionsModeEnum.Value(
            "MonitoringBetaDashboardMosaicLayoutTilesWidgetXyChartChartOptionsModeEnum%s"
            % resource
        )

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return resource
        return dashboard_pb2.MonitoringBetaDashboardMosaicLayoutTilesWidgetXyChartChartOptionsModeEnum.Name(
            resource
        )[
            len(
                "MonitoringBetaDashboardMosaicLayoutTilesWidgetXyChartChartOptionsModeEnum"
            ) :
        ]


class DashboardMosaicLayoutTilesWidgetScorecardTimeSeriesQueryTimeSeriesFilterAggregationPerSeriesAlignerEnum(
    object
):
    @classmethod
    def to_proto(self, resource):
        if not resource:
            return resource
        return dashboard_pb2.MonitoringBetaDashboardMosaicLayoutTilesWidgetScorecardTimeSeriesQueryTimeSeriesFilterAggregationPerSeriesAlignerEnum.Value(
            "MonitoringBetaDashboardMosaicLayoutTilesWidgetScorecardTimeSeriesQueryTimeSeriesFilterAggregationPerSeriesAlignerEnum%s"
            % resource
        )

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return resource
        return dashboard_pb2.MonitoringBetaDashboardMosaicLayoutTilesWidgetScorecardTimeSeriesQueryTimeSeriesFilterAggregationPerSeriesAlignerEnum.Name(
            resource
        )[
            len(
                "MonitoringBetaDashboardMosaicLayoutTilesWidgetScorecardTimeSeriesQueryTimeSeriesFilterAggregationPerSeriesAlignerEnum"
            ) :
        ]


class DashboardMosaicLayoutTilesWidgetScorecardTimeSeriesQueryTimeSeriesFilterAggregationCrossSeriesReducerEnum(
    object
):
    @classmethod
    def to_proto(self, resource):
        if not resource:
            return resource
        return dashboard_pb2.MonitoringBetaDashboardMosaicLayoutTilesWidgetScorecardTimeSeriesQueryTimeSeriesFilterAggregationCrossSeriesReducerEnum.Value(
            "MonitoringBetaDashboardMosaicLayoutTilesWidgetScorecardTimeSeriesQueryTimeSeriesFilterAggregationCrossSeriesReducerEnum%s"
            % resource
        )

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return resource
        return dashboard_pb2.MonitoringBetaDashboardMosaicLayoutTilesWidgetScorecardTimeSeriesQueryTimeSeriesFilterAggregationCrossSeriesReducerEnum.Name(
            resource
        )[
            len(
                "MonitoringBetaDashboardMosaicLayoutTilesWidgetScorecardTimeSeriesQueryTimeSeriesFilterAggregationCrossSeriesReducerEnum"
            ) :
        ]


class DashboardMosaicLayoutTilesWidgetScorecardTimeSeriesQueryTimeSeriesFilterSecondaryAggregationPerSeriesAlignerEnum(
    object
):
    @classmethod
    def to_proto(self, resource):
        if not resource:
            return resource
        return dashboard_pb2.MonitoringBetaDashboardMosaicLayoutTilesWidgetScorecardTimeSeriesQueryTimeSeriesFilterSecondaryAggregationPerSeriesAlignerEnum.Value(
            "MonitoringBetaDashboardMosaicLayoutTilesWidgetScorecardTimeSeriesQueryTimeSeriesFilterSecondaryAggregationPerSeriesAlignerEnum%s"
            % resource
        )

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return resource
        return dashboard_pb2.MonitoringBetaDashboardMosaicLayoutTilesWidgetScorecardTimeSeriesQueryTimeSeriesFilterSecondaryAggregationPerSeriesAlignerEnum.Name(
            resource
        )[
            len(
                "MonitoringBetaDashboardMosaicLayoutTilesWidgetScorecardTimeSeriesQueryTimeSeriesFilterSecondaryAggregationPerSeriesAlignerEnum"
            ) :
        ]


class DashboardMosaicLayoutTilesWidgetScorecardTimeSeriesQueryTimeSeriesFilterSecondaryAggregationCrossSeriesReducerEnum(
    object
):
    @classmethod
    def to_proto(self, resource):
        if not resource:
            return resource
        return dashboard_pb2.MonitoringBetaDashboardMosaicLayoutTilesWidgetScorecardTimeSeriesQueryTimeSeriesFilterSecondaryAggregationCrossSeriesReducerEnum.Value(
            "MonitoringBetaDashboardMosaicLayoutTilesWidgetScorecardTimeSeriesQueryTimeSeriesFilterSecondaryAggregationCrossSeriesReducerEnum%s"
            % resource
        )

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return resource
        return dashboard_pb2.MonitoringBetaDashboardMosaicLayoutTilesWidgetScorecardTimeSeriesQueryTimeSeriesFilterSecondaryAggregationCrossSeriesReducerEnum.Name(
            resource
        )[
            len(
                "MonitoringBetaDashboardMosaicLayoutTilesWidgetScorecardTimeSeriesQueryTimeSeriesFilterSecondaryAggregationCrossSeriesReducerEnum"
            ) :
        ]


class DashboardMosaicLayoutTilesWidgetScorecardTimeSeriesQueryTimeSeriesFilterPickTimeSeriesFilterRankingMethodEnum(
    object
):
    @classmethod
    def to_proto(self, resource):
        if not resource:
            return resource
        return dashboard_pb2.MonitoringBetaDashboardMosaicLayoutTilesWidgetScorecardTimeSeriesQueryTimeSeriesFilterPickTimeSeriesFilterRankingMethodEnum.Value(
            "MonitoringBetaDashboardMosaicLayoutTilesWidgetScorecardTimeSeriesQueryTimeSeriesFilterPickTimeSeriesFilterRankingMethodEnum%s"
            % resource
        )

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return resource
        return dashboard_pb2.MonitoringBetaDashboardMosaicLayoutTilesWidgetScorecardTimeSeriesQueryTimeSeriesFilterPickTimeSeriesFilterRankingMethodEnum.Name(
            resource
        )[
            len(
                "MonitoringBetaDashboardMosaicLayoutTilesWidgetScorecardTimeSeriesQueryTimeSeriesFilterPickTimeSeriesFilterRankingMethodEnum"
            ) :
        ]


class DashboardMosaicLayoutTilesWidgetScorecardTimeSeriesQueryTimeSeriesFilterPickTimeSeriesFilterDirectionEnum(
    object
):
    @classmethod
    def to_proto(self, resource):
        if not resource:
            return resource
        return dashboard_pb2.MonitoringBetaDashboardMosaicLayoutTilesWidgetScorecardTimeSeriesQueryTimeSeriesFilterPickTimeSeriesFilterDirectionEnum.Value(
            "MonitoringBetaDashboardMosaicLayoutTilesWidgetScorecardTimeSeriesQueryTimeSeriesFilterPickTimeSeriesFilterDirectionEnum%s"
            % resource
        )

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return resource
        return dashboard_pb2.MonitoringBetaDashboardMosaicLayoutTilesWidgetScorecardTimeSeriesQueryTimeSeriesFilterPickTimeSeriesFilterDirectionEnum.Name(
            resource
        )[
            len(
                "MonitoringBetaDashboardMosaicLayoutTilesWidgetScorecardTimeSeriesQueryTimeSeriesFilterPickTimeSeriesFilterDirectionEnum"
            ) :
        ]


class DashboardMosaicLayoutTilesWidgetScorecardTimeSeriesQueryTimeSeriesFilterRatioNumeratorAggregationPerSeriesAlignerEnum(
    object
):
    @classmethod
    def to_proto(self, resource):
        if not resource:
            return resource
        return dashboard_pb2.MonitoringBetaDashboardMosaicLayoutTilesWidgetScorecardTimeSeriesQueryTimeSeriesFilterRatioNumeratorAggregationPerSeriesAlignerEnum.Value(
            "MonitoringBetaDashboardMosaicLayoutTilesWidgetScorecardTimeSeriesQueryTimeSeriesFilterRatioNumeratorAggregationPerSeriesAlignerEnum%s"
            % resource
        )

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return resource
        return dashboard_pb2.MonitoringBetaDashboardMosaicLayoutTilesWidgetScorecardTimeSeriesQueryTimeSeriesFilterRatioNumeratorAggregationPerSeriesAlignerEnum.Name(
            resource
        )[
            len(
                "MonitoringBetaDashboardMosaicLayoutTilesWidgetScorecardTimeSeriesQueryTimeSeriesFilterRatioNumeratorAggregationPerSeriesAlignerEnum"
            ) :
        ]


class DashboardMosaicLayoutTilesWidgetScorecardTimeSeriesQueryTimeSeriesFilterRatioNumeratorAggregationCrossSeriesReducerEnum(
    object
):
    @classmethod
    def to_proto(self, resource):
        if not resource:
            return resource
        return dashboard_pb2.MonitoringBetaDashboardMosaicLayoutTilesWidgetScorecardTimeSeriesQueryTimeSeriesFilterRatioNumeratorAggregationCrossSeriesReducerEnum.Value(
            "MonitoringBetaDashboardMosaicLayoutTilesWidgetScorecardTimeSeriesQueryTimeSeriesFilterRatioNumeratorAggregationCrossSeriesReducerEnum%s"
            % resource
        )

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return resource
        return dashboard_pb2.MonitoringBetaDashboardMosaicLayoutTilesWidgetScorecardTimeSeriesQueryTimeSeriesFilterRatioNumeratorAggregationCrossSeriesReducerEnum.Name(
            resource
        )[
            len(
                "MonitoringBetaDashboardMosaicLayoutTilesWidgetScorecardTimeSeriesQueryTimeSeriesFilterRatioNumeratorAggregationCrossSeriesReducerEnum"
            ) :
        ]


class DashboardMosaicLayoutTilesWidgetScorecardTimeSeriesQueryTimeSeriesFilterRatioDenominatorAggregationPerSeriesAlignerEnum(
    object
):
    @classmethod
    def to_proto(self, resource):
        if not resource:
            return resource
        return dashboard_pb2.MonitoringBetaDashboardMosaicLayoutTilesWidgetScorecardTimeSeriesQueryTimeSeriesFilterRatioDenominatorAggregationPerSeriesAlignerEnum.Value(
            "MonitoringBetaDashboardMosaicLayoutTilesWidgetScorecardTimeSeriesQueryTimeSeriesFilterRatioDenominatorAggregationPerSeriesAlignerEnum%s"
            % resource
        )

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return resource
        return dashboard_pb2.MonitoringBetaDashboardMosaicLayoutTilesWidgetScorecardTimeSeriesQueryTimeSeriesFilterRatioDenominatorAggregationPerSeriesAlignerEnum.Name(
            resource
        )[
            len(
                "MonitoringBetaDashboardMosaicLayoutTilesWidgetScorecardTimeSeriesQueryTimeSeriesFilterRatioDenominatorAggregationPerSeriesAlignerEnum"
            ) :
        ]


class DashboardMosaicLayoutTilesWidgetScorecardTimeSeriesQueryTimeSeriesFilterRatioDenominatorAggregationCrossSeriesReducerEnum(
    object
):
    @classmethod
    def to_proto(self, resource):
        if not resource:
            return resource
        return dashboard_pb2.MonitoringBetaDashboardMosaicLayoutTilesWidgetScorecardTimeSeriesQueryTimeSeriesFilterRatioDenominatorAggregationCrossSeriesReducerEnum.Value(
            "MonitoringBetaDashboardMosaicLayoutTilesWidgetScorecardTimeSeriesQueryTimeSeriesFilterRatioDenominatorAggregationCrossSeriesReducerEnum%s"
            % resource
        )

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return resource
        return dashboard_pb2.MonitoringBetaDashboardMosaicLayoutTilesWidgetScorecardTimeSeriesQueryTimeSeriesFilterRatioDenominatorAggregationCrossSeriesReducerEnum.Name(
            resource
        )[
            len(
                "MonitoringBetaDashboardMosaicLayoutTilesWidgetScorecardTimeSeriesQueryTimeSeriesFilterRatioDenominatorAggregationCrossSeriesReducerEnum"
            ) :
        ]


class DashboardMosaicLayoutTilesWidgetScorecardTimeSeriesQueryTimeSeriesFilterRatioSecondaryAggregationPerSeriesAlignerEnum(
    object
):
    @classmethod
    def to_proto(self, resource):
        if not resource:
            return resource
        return dashboard_pb2.MonitoringBetaDashboardMosaicLayoutTilesWidgetScorecardTimeSeriesQueryTimeSeriesFilterRatioSecondaryAggregationPerSeriesAlignerEnum.Value(
            "MonitoringBetaDashboardMosaicLayoutTilesWidgetScorecardTimeSeriesQueryTimeSeriesFilterRatioSecondaryAggregationPerSeriesAlignerEnum%s"
            % resource
        )

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return resource
        return dashboard_pb2.MonitoringBetaDashboardMosaicLayoutTilesWidgetScorecardTimeSeriesQueryTimeSeriesFilterRatioSecondaryAggregationPerSeriesAlignerEnum.Name(
            resource
        )[
            len(
                "MonitoringBetaDashboardMosaicLayoutTilesWidgetScorecardTimeSeriesQueryTimeSeriesFilterRatioSecondaryAggregationPerSeriesAlignerEnum"
            ) :
        ]


class DashboardMosaicLayoutTilesWidgetScorecardTimeSeriesQueryTimeSeriesFilterRatioSecondaryAggregationCrossSeriesReducerEnum(
    object
):
    @classmethod
    def to_proto(self, resource):
        if not resource:
            return resource
        return dashboard_pb2.MonitoringBetaDashboardMosaicLayoutTilesWidgetScorecardTimeSeriesQueryTimeSeriesFilterRatioSecondaryAggregationCrossSeriesReducerEnum.Value(
            "MonitoringBetaDashboardMosaicLayoutTilesWidgetScorecardTimeSeriesQueryTimeSeriesFilterRatioSecondaryAggregationCrossSeriesReducerEnum%s"
            % resource
        )

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return resource
        return dashboard_pb2.MonitoringBetaDashboardMosaicLayoutTilesWidgetScorecardTimeSeriesQueryTimeSeriesFilterRatioSecondaryAggregationCrossSeriesReducerEnum.Name(
            resource
        )[
            len(
                "MonitoringBetaDashboardMosaicLayoutTilesWidgetScorecardTimeSeriesQueryTimeSeriesFilterRatioSecondaryAggregationCrossSeriesReducerEnum"
            ) :
        ]


class DashboardMosaicLayoutTilesWidgetScorecardTimeSeriesQueryTimeSeriesFilterRatioPickTimeSeriesFilterRankingMethodEnum(
    object
):
    @classmethod
    def to_proto(self, resource):
        if not resource:
            return resource
        return dashboard_pb2.MonitoringBetaDashboardMosaicLayoutTilesWidgetScorecardTimeSeriesQueryTimeSeriesFilterRatioPickTimeSeriesFilterRankingMethodEnum.Value(
            "MonitoringBetaDashboardMosaicLayoutTilesWidgetScorecardTimeSeriesQueryTimeSeriesFilterRatioPickTimeSeriesFilterRankingMethodEnum%s"
            % resource
        )

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return resource
        return dashboard_pb2.MonitoringBetaDashboardMosaicLayoutTilesWidgetScorecardTimeSeriesQueryTimeSeriesFilterRatioPickTimeSeriesFilterRankingMethodEnum.Name(
            resource
        )[
            len(
                "MonitoringBetaDashboardMosaicLayoutTilesWidgetScorecardTimeSeriesQueryTimeSeriesFilterRatioPickTimeSeriesFilterRankingMethodEnum"
            ) :
        ]


class DashboardMosaicLayoutTilesWidgetScorecardTimeSeriesQueryTimeSeriesFilterRatioPickTimeSeriesFilterDirectionEnum(
    object
):
    @classmethod
    def to_proto(self, resource):
        if not resource:
            return resource
        return dashboard_pb2.MonitoringBetaDashboardMosaicLayoutTilesWidgetScorecardTimeSeriesQueryTimeSeriesFilterRatioPickTimeSeriesFilterDirectionEnum.Value(
            "MonitoringBetaDashboardMosaicLayoutTilesWidgetScorecardTimeSeriesQueryTimeSeriesFilterRatioPickTimeSeriesFilterDirectionEnum%s"
            % resource
        )

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return resource
        return dashboard_pb2.MonitoringBetaDashboardMosaicLayoutTilesWidgetScorecardTimeSeriesQueryTimeSeriesFilterRatioPickTimeSeriesFilterDirectionEnum.Name(
            resource
        )[
            len(
                "MonitoringBetaDashboardMosaicLayoutTilesWidgetScorecardTimeSeriesQueryTimeSeriesFilterRatioPickTimeSeriesFilterDirectionEnum"
            ) :
        ]


class DashboardMosaicLayoutTilesWidgetScorecardSparkChartViewSparkChartTypeEnum(object):
    @classmethod
    def to_proto(self, resource):
        if not resource:
            return resource
        return dashboard_pb2.MonitoringBetaDashboardMosaicLayoutTilesWidgetScorecardSparkChartViewSparkChartTypeEnum.Value(
            "MonitoringBetaDashboardMosaicLayoutTilesWidgetScorecardSparkChartViewSparkChartTypeEnum%s"
            % resource
        )

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return resource
        return dashboard_pb2.MonitoringBetaDashboardMosaicLayoutTilesWidgetScorecardSparkChartViewSparkChartTypeEnum.Name(
            resource
        )[
            len(
                "MonitoringBetaDashboardMosaicLayoutTilesWidgetScorecardSparkChartViewSparkChartTypeEnum"
            ) :
        ]


class DashboardMosaicLayoutTilesWidgetScorecardThresholdsColorEnum(object):
    @classmethod
    def to_proto(self, resource):
        if not resource:
            return resource
        return dashboard_pb2.MonitoringBetaDashboardMosaicLayoutTilesWidgetScorecardThresholdsColorEnum.Value(
            "MonitoringBetaDashboardMosaicLayoutTilesWidgetScorecardThresholdsColorEnum%s"
            % resource
        )

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return resource
        return dashboard_pb2.MonitoringBetaDashboardMosaicLayoutTilesWidgetScorecardThresholdsColorEnum.Name(
            resource
        )[
            len(
                "MonitoringBetaDashboardMosaicLayoutTilesWidgetScorecardThresholdsColorEnum"
            ) :
        ]


class DashboardMosaicLayoutTilesWidgetScorecardThresholdsDirectionEnum(object):
    @classmethod
    def to_proto(self, resource):
        if not resource:
            return resource
        return dashboard_pb2.MonitoringBetaDashboardMosaicLayoutTilesWidgetScorecardThresholdsDirectionEnum.Value(
            "MonitoringBetaDashboardMosaicLayoutTilesWidgetScorecardThresholdsDirectionEnum%s"
            % resource
        )

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return resource
        return dashboard_pb2.MonitoringBetaDashboardMosaicLayoutTilesWidgetScorecardThresholdsDirectionEnum.Name(
            resource
        )[
            len(
                "MonitoringBetaDashboardMosaicLayoutTilesWidgetScorecardThresholdsDirectionEnum"
            ) :
        ]


class DashboardMosaicLayoutTilesWidgetTextFormatEnum(object):
    @classmethod
    def to_proto(self, resource):
        if not resource:
            return resource
        return dashboard_pb2.MonitoringBetaDashboardMosaicLayoutTilesWidgetTextFormatEnum.Value(
            "MonitoringBetaDashboardMosaicLayoutTilesWidgetTextFormatEnum%s" % resource
        )

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return resource
        return dashboard_pb2.MonitoringBetaDashboardMosaicLayoutTilesWidgetTextFormatEnum.Name(
            resource
        )[
            len("MonitoringBetaDashboardMosaicLayoutTilesWidgetTextFormatEnum") :
        ]


class DashboardRowLayoutRowsWidgetsXyChartDataSetsTimeSeriesQueryTimeSeriesFilterAggregationPerSeriesAlignerEnum(
    object
):
    @classmethod
    def to_proto(self, resource):
        if not resource:
            return resource
        return dashboard_pb2.MonitoringBetaDashboardRowLayoutRowsWidgetsXyChartDataSetsTimeSeriesQueryTimeSeriesFilterAggregationPerSeriesAlignerEnum.Value(
            "MonitoringBetaDashboardRowLayoutRowsWidgetsXyChartDataSetsTimeSeriesQueryTimeSeriesFilterAggregationPerSeriesAlignerEnum%s"
            % resource
        )

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return resource
        return dashboard_pb2.MonitoringBetaDashboardRowLayoutRowsWidgetsXyChartDataSetsTimeSeriesQueryTimeSeriesFilterAggregationPerSeriesAlignerEnum.Name(
            resource
        )[
            len(
                "MonitoringBetaDashboardRowLayoutRowsWidgetsXyChartDataSetsTimeSeriesQueryTimeSeriesFilterAggregationPerSeriesAlignerEnum"
            ) :
        ]


class DashboardRowLayoutRowsWidgetsXyChartDataSetsTimeSeriesQueryTimeSeriesFilterAggregationCrossSeriesReducerEnum(
    object
):
    @classmethod
    def to_proto(self, resource):
        if not resource:
            return resource
        return dashboard_pb2.MonitoringBetaDashboardRowLayoutRowsWidgetsXyChartDataSetsTimeSeriesQueryTimeSeriesFilterAggregationCrossSeriesReducerEnum.Value(
            "MonitoringBetaDashboardRowLayoutRowsWidgetsXyChartDataSetsTimeSeriesQueryTimeSeriesFilterAggregationCrossSeriesReducerEnum%s"
            % resource
        )

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return resource
        return dashboard_pb2.MonitoringBetaDashboardRowLayoutRowsWidgetsXyChartDataSetsTimeSeriesQueryTimeSeriesFilterAggregationCrossSeriesReducerEnum.Name(
            resource
        )[
            len(
                "MonitoringBetaDashboardRowLayoutRowsWidgetsXyChartDataSetsTimeSeriesQueryTimeSeriesFilterAggregationCrossSeriesReducerEnum"
            ) :
        ]


class DashboardRowLayoutRowsWidgetsXyChartDataSetsTimeSeriesQueryTimeSeriesFilterSecondaryAggregationPerSeriesAlignerEnum(
    object
):
    @classmethod
    def to_proto(self, resource):
        if not resource:
            return resource
        return dashboard_pb2.MonitoringBetaDashboardRowLayoutRowsWidgetsXyChartDataSetsTimeSeriesQueryTimeSeriesFilterSecondaryAggregationPerSeriesAlignerEnum.Value(
            "MonitoringBetaDashboardRowLayoutRowsWidgetsXyChartDataSetsTimeSeriesQueryTimeSeriesFilterSecondaryAggregationPerSeriesAlignerEnum%s"
            % resource
        )

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return resource
        return dashboard_pb2.MonitoringBetaDashboardRowLayoutRowsWidgetsXyChartDataSetsTimeSeriesQueryTimeSeriesFilterSecondaryAggregationPerSeriesAlignerEnum.Name(
            resource
        )[
            len(
                "MonitoringBetaDashboardRowLayoutRowsWidgetsXyChartDataSetsTimeSeriesQueryTimeSeriesFilterSecondaryAggregationPerSeriesAlignerEnum"
            ) :
        ]


class DashboardRowLayoutRowsWidgetsXyChartDataSetsTimeSeriesQueryTimeSeriesFilterSecondaryAggregationCrossSeriesReducerEnum(
    object
):
    @classmethod
    def to_proto(self, resource):
        if not resource:
            return resource
        return dashboard_pb2.MonitoringBetaDashboardRowLayoutRowsWidgetsXyChartDataSetsTimeSeriesQueryTimeSeriesFilterSecondaryAggregationCrossSeriesReducerEnum.Value(
            "MonitoringBetaDashboardRowLayoutRowsWidgetsXyChartDataSetsTimeSeriesQueryTimeSeriesFilterSecondaryAggregationCrossSeriesReducerEnum%s"
            % resource
        )

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return resource
        return dashboard_pb2.MonitoringBetaDashboardRowLayoutRowsWidgetsXyChartDataSetsTimeSeriesQueryTimeSeriesFilterSecondaryAggregationCrossSeriesReducerEnum.Name(
            resource
        )[
            len(
                "MonitoringBetaDashboardRowLayoutRowsWidgetsXyChartDataSetsTimeSeriesQueryTimeSeriesFilterSecondaryAggregationCrossSeriesReducerEnum"
            ) :
        ]


class DashboardRowLayoutRowsWidgetsXyChartDataSetsTimeSeriesQueryTimeSeriesFilterPickTimeSeriesFilterRankingMethodEnum(
    object
):
    @classmethod
    def to_proto(self, resource):
        if not resource:
            return resource
        return dashboard_pb2.MonitoringBetaDashboardRowLayoutRowsWidgetsXyChartDataSetsTimeSeriesQueryTimeSeriesFilterPickTimeSeriesFilterRankingMethodEnum.Value(
            "MonitoringBetaDashboardRowLayoutRowsWidgetsXyChartDataSetsTimeSeriesQueryTimeSeriesFilterPickTimeSeriesFilterRankingMethodEnum%s"
            % resource
        )

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return resource
        return dashboard_pb2.MonitoringBetaDashboardRowLayoutRowsWidgetsXyChartDataSetsTimeSeriesQueryTimeSeriesFilterPickTimeSeriesFilterRankingMethodEnum.Name(
            resource
        )[
            len(
                "MonitoringBetaDashboardRowLayoutRowsWidgetsXyChartDataSetsTimeSeriesQueryTimeSeriesFilterPickTimeSeriesFilterRankingMethodEnum"
            ) :
        ]


class DashboardRowLayoutRowsWidgetsXyChartDataSetsTimeSeriesQueryTimeSeriesFilterPickTimeSeriesFilterDirectionEnum(
    object
):
    @classmethod
    def to_proto(self, resource):
        if not resource:
            return resource
        return dashboard_pb2.MonitoringBetaDashboardRowLayoutRowsWidgetsXyChartDataSetsTimeSeriesQueryTimeSeriesFilterPickTimeSeriesFilterDirectionEnum.Value(
            "MonitoringBetaDashboardRowLayoutRowsWidgetsXyChartDataSetsTimeSeriesQueryTimeSeriesFilterPickTimeSeriesFilterDirectionEnum%s"
            % resource
        )

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return resource
        return dashboard_pb2.MonitoringBetaDashboardRowLayoutRowsWidgetsXyChartDataSetsTimeSeriesQueryTimeSeriesFilterPickTimeSeriesFilterDirectionEnum.Name(
            resource
        )[
            len(
                "MonitoringBetaDashboardRowLayoutRowsWidgetsXyChartDataSetsTimeSeriesQueryTimeSeriesFilterPickTimeSeriesFilterDirectionEnum"
            ) :
        ]


class DashboardRowLayoutRowsWidgetsXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioNumeratorAggregationPerSeriesAlignerEnum(
    object
):
    @classmethod
    def to_proto(self, resource):
        if not resource:
            return resource
        return dashboard_pb2.MonitoringBetaDashboardRowLayoutRowsWidgetsXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioNumeratorAggregationPerSeriesAlignerEnum.Value(
            "MonitoringBetaDashboardRowLayoutRowsWidgetsXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioNumeratorAggregationPerSeriesAlignerEnum%s"
            % resource
        )

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return resource
        return dashboard_pb2.MonitoringBetaDashboardRowLayoutRowsWidgetsXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioNumeratorAggregationPerSeriesAlignerEnum.Name(
            resource
        )[
            len(
                "MonitoringBetaDashboardRowLayoutRowsWidgetsXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioNumeratorAggregationPerSeriesAlignerEnum"
            ) :
        ]


class DashboardRowLayoutRowsWidgetsXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioNumeratorAggregationCrossSeriesReducerEnum(
    object
):
    @classmethod
    def to_proto(self, resource):
        if not resource:
            return resource
        return dashboard_pb2.MonitoringBetaDashboardRowLayoutRowsWidgetsXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioNumeratorAggregationCrossSeriesReducerEnum.Value(
            "MonitoringBetaDashboardRowLayoutRowsWidgetsXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioNumeratorAggregationCrossSeriesReducerEnum%s"
            % resource
        )

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return resource
        return dashboard_pb2.MonitoringBetaDashboardRowLayoutRowsWidgetsXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioNumeratorAggregationCrossSeriesReducerEnum.Name(
            resource
        )[
            len(
                "MonitoringBetaDashboardRowLayoutRowsWidgetsXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioNumeratorAggregationCrossSeriesReducerEnum"
            ) :
        ]


class DashboardRowLayoutRowsWidgetsXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioDenominatorAggregationPerSeriesAlignerEnum(
    object
):
    @classmethod
    def to_proto(self, resource):
        if not resource:
            return resource
        return dashboard_pb2.MonitoringBetaDashboardRowLayoutRowsWidgetsXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioDenominatorAggregationPerSeriesAlignerEnum.Value(
            "MonitoringBetaDashboardRowLayoutRowsWidgetsXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioDenominatorAggregationPerSeriesAlignerEnum%s"
            % resource
        )

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return resource
        return dashboard_pb2.MonitoringBetaDashboardRowLayoutRowsWidgetsXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioDenominatorAggregationPerSeriesAlignerEnum.Name(
            resource
        )[
            len(
                "MonitoringBetaDashboardRowLayoutRowsWidgetsXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioDenominatorAggregationPerSeriesAlignerEnum"
            ) :
        ]


class DashboardRowLayoutRowsWidgetsXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioDenominatorAggregationCrossSeriesReducerEnum(
    object
):
    @classmethod
    def to_proto(self, resource):
        if not resource:
            return resource
        return dashboard_pb2.MonitoringBetaDashboardRowLayoutRowsWidgetsXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioDenominatorAggregationCrossSeriesReducerEnum.Value(
            "MonitoringBetaDashboardRowLayoutRowsWidgetsXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioDenominatorAggregationCrossSeriesReducerEnum%s"
            % resource
        )

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return resource
        return dashboard_pb2.MonitoringBetaDashboardRowLayoutRowsWidgetsXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioDenominatorAggregationCrossSeriesReducerEnum.Name(
            resource
        )[
            len(
                "MonitoringBetaDashboardRowLayoutRowsWidgetsXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioDenominatorAggregationCrossSeriesReducerEnum"
            ) :
        ]


class DashboardRowLayoutRowsWidgetsXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioSecondaryAggregationPerSeriesAlignerEnum(
    object
):
    @classmethod
    def to_proto(self, resource):
        if not resource:
            return resource
        return dashboard_pb2.MonitoringBetaDashboardRowLayoutRowsWidgetsXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioSecondaryAggregationPerSeriesAlignerEnum.Value(
            "MonitoringBetaDashboardRowLayoutRowsWidgetsXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioSecondaryAggregationPerSeriesAlignerEnum%s"
            % resource
        )

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return resource
        return dashboard_pb2.MonitoringBetaDashboardRowLayoutRowsWidgetsXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioSecondaryAggregationPerSeriesAlignerEnum.Name(
            resource
        )[
            len(
                "MonitoringBetaDashboardRowLayoutRowsWidgetsXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioSecondaryAggregationPerSeriesAlignerEnum"
            ) :
        ]


class DashboardRowLayoutRowsWidgetsXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioSecondaryAggregationCrossSeriesReducerEnum(
    object
):
    @classmethod
    def to_proto(self, resource):
        if not resource:
            return resource
        return dashboard_pb2.MonitoringBetaDashboardRowLayoutRowsWidgetsXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioSecondaryAggregationCrossSeriesReducerEnum.Value(
            "MonitoringBetaDashboardRowLayoutRowsWidgetsXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioSecondaryAggregationCrossSeriesReducerEnum%s"
            % resource
        )

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return resource
        return dashboard_pb2.MonitoringBetaDashboardRowLayoutRowsWidgetsXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioSecondaryAggregationCrossSeriesReducerEnum.Name(
            resource
        )[
            len(
                "MonitoringBetaDashboardRowLayoutRowsWidgetsXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioSecondaryAggregationCrossSeriesReducerEnum"
            ) :
        ]


class DashboardRowLayoutRowsWidgetsXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioPickTimeSeriesFilterRankingMethodEnum(
    object
):
    @classmethod
    def to_proto(self, resource):
        if not resource:
            return resource
        return dashboard_pb2.MonitoringBetaDashboardRowLayoutRowsWidgetsXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioPickTimeSeriesFilterRankingMethodEnum.Value(
            "MonitoringBetaDashboardRowLayoutRowsWidgetsXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioPickTimeSeriesFilterRankingMethodEnum%s"
            % resource
        )

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return resource
        return dashboard_pb2.MonitoringBetaDashboardRowLayoutRowsWidgetsXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioPickTimeSeriesFilterRankingMethodEnum.Name(
            resource
        )[
            len(
                "MonitoringBetaDashboardRowLayoutRowsWidgetsXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioPickTimeSeriesFilterRankingMethodEnum"
            ) :
        ]


class DashboardRowLayoutRowsWidgetsXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioPickTimeSeriesFilterDirectionEnum(
    object
):
    @classmethod
    def to_proto(self, resource):
        if not resource:
            return resource
        return dashboard_pb2.MonitoringBetaDashboardRowLayoutRowsWidgetsXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioPickTimeSeriesFilterDirectionEnum.Value(
            "MonitoringBetaDashboardRowLayoutRowsWidgetsXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioPickTimeSeriesFilterDirectionEnum%s"
            % resource
        )

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return resource
        return dashboard_pb2.MonitoringBetaDashboardRowLayoutRowsWidgetsXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioPickTimeSeriesFilterDirectionEnum.Name(
            resource
        )[
            len(
                "MonitoringBetaDashboardRowLayoutRowsWidgetsXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioPickTimeSeriesFilterDirectionEnum"
            ) :
        ]


class DashboardRowLayoutRowsWidgetsXyChartDataSetsPlotTypeEnum(object):
    @classmethod
    def to_proto(self, resource):
        if not resource:
            return resource
        return dashboard_pb2.MonitoringBetaDashboardRowLayoutRowsWidgetsXyChartDataSetsPlotTypeEnum.Value(
            "MonitoringBetaDashboardRowLayoutRowsWidgetsXyChartDataSetsPlotTypeEnum%s"
            % resource
        )

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return resource
        return dashboard_pb2.MonitoringBetaDashboardRowLayoutRowsWidgetsXyChartDataSetsPlotTypeEnum.Name(
            resource
        )[
            len(
                "MonitoringBetaDashboardRowLayoutRowsWidgetsXyChartDataSetsPlotTypeEnum"
            ) :
        ]


class DashboardRowLayoutRowsWidgetsXyChartThresholdsColorEnum(object):
    @classmethod
    def to_proto(self, resource):
        if not resource:
            return resource
        return dashboard_pb2.MonitoringBetaDashboardRowLayoutRowsWidgetsXyChartThresholdsColorEnum.Value(
            "MonitoringBetaDashboardRowLayoutRowsWidgetsXyChartThresholdsColorEnum%s"
            % resource
        )

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return resource
        return dashboard_pb2.MonitoringBetaDashboardRowLayoutRowsWidgetsXyChartThresholdsColorEnum.Name(
            resource
        )[
            len(
                "MonitoringBetaDashboardRowLayoutRowsWidgetsXyChartThresholdsColorEnum"
            ) :
        ]


class DashboardRowLayoutRowsWidgetsXyChartThresholdsDirectionEnum(object):
    @classmethod
    def to_proto(self, resource):
        if not resource:
            return resource
        return dashboard_pb2.MonitoringBetaDashboardRowLayoutRowsWidgetsXyChartThresholdsDirectionEnum.Value(
            "MonitoringBetaDashboardRowLayoutRowsWidgetsXyChartThresholdsDirectionEnum%s"
            % resource
        )

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return resource
        return dashboard_pb2.MonitoringBetaDashboardRowLayoutRowsWidgetsXyChartThresholdsDirectionEnum.Name(
            resource
        )[
            len(
                "MonitoringBetaDashboardRowLayoutRowsWidgetsXyChartThresholdsDirectionEnum"
            ) :
        ]


class DashboardRowLayoutRowsWidgetsXyChartXAxisScaleEnum(object):
    @classmethod
    def to_proto(self, resource):
        if not resource:
            return resource
        return dashboard_pb2.MonitoringBetaDashboardRowLayoutRowsWidgetsXyChartXAxisScaleEnum.Value(
            "MonitoringBetaDashboardRowLayoutRowsWidgetsXyChartXAxisScaleEnum%s"
            % resource
        )

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return resource
        return dashboard_pb2.MonitoringBetaDashboardRowLayoutRowsWidgetsXyChartXAxisScaleEnum.Name(
            resource
        )[
            len("MonitoringBetaDashboardRowLayoutRowsWidgetsXyChartXAxisScaleEnum") :
        ]


class DashboardRowLayoutRowsWidgetsXyChartYAxisScaleEnum(object):
    @classmethod
    def to_proto(self, resource):
        if not resource:
            return resource
        return dashboard_pb2.MonitoringBetaDashboardRowLayoutRowsWidgetsXyChartYAxisScaleEnum.Value(
            "MonitoringBetaDashboardRowLayoutRowsWidgetsXyChartYAxisScaleEnum%s"
            % resource
        )

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return resource
        return dashboard_pb2.MonitoringBetaDashboardRowLayoutRowsWidgetsXyChartYAxisScaleEnum.Name(
            resource
        )[
            len("MonitoringBetaDashboardRowLayoutRowsWidgetsXyChartYAxisScaleEnum") :
        ]


class DashboardRowLayoutRowsWidgetsXyChartChartOptionsModeEnum(object):
    @classmethod
    def to_proto(self, resource):
        if not resource:
            return resource
        return dashboard_pb2.MonitoringBetaDashboardRowLayoutRowsWidgetsXyChartChartOptionsModeEnum.Value(
            "MonitoringBetaDashboardRowLayoutRowsWidgetsXyChartChartOptionsModeEnum%s"
            % resource
        )

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return resource
        return dashboard_pb2.MonitoringBetaDashboardRowLayoutRowsWidgetsXyChartChartOptionsModeEnum.Name(
            resource
        )[
            len(
                "MonitoringBetaDashboardRowLayoutRowsWidgetsXyChartChartOptionsModeEnum"
            ) :
        ]


class DashboardRowLayoutRowsWidgetsScorecardTimeSeriesQueryTimeSeriesFilterAggregationPerSeriesAlignerEnum(
    object
):
    @classmethod
    def to_proto(self, resource):
        if not resource:
            return resource
        return dashboard_pb2.MonitoringBetaDashboardRowLayoutRowsWidgetsScorecardTimeSeriesQueryTimeSeriesFilterAggregationPerSeriesAlignerEnum.Value(
            "MonitoringBetaDashboardRowLayoutRowsWidgetsScorecardTimeSeriesQueryTimeSeriesFilterAggregationPerSeriesAlignerEnum%s"
            % resource
        )

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return resource
        return dashboard_pb2.MonitoringBetaDashboardRowLayoutRowsWidgetsScorecardTimeSeriesQueryTimeSeriesFilterAggregationPerSeriesAlignerEnum.Name(
            resource
        )[
            len(
                "MonitoringBetaDashboardRowLayoutRowsWidgetsScorecardTimeSeriesQueryTimeSeriesFilterAggregationPerSeriesAlignerEnum"
            ) :
        ]


class DashboardRowLayoutRowsWidgetsScorecardTimeSeriesQueryTimeSeriesFilterAggregationCrossSeriesReducerEnum(
    object
):
    @classmethod
    def to_proto(self, resource):
        if not resource:
            return resource
        return dashboard_pb2.MonitoringBetaDashboardRowLayoutRowsWidgetsScorecardTimeSeriesQueryTimeSeriesFilterAggregationCrossSeriesReducerEnum.Value(
            "MonitoringBetaDashboardRowLayoutRowsWidgetsScorecardTimeSeriesQueryTimeSeriesFilterAggregationCrossSeriesReducerEnum%s"
            % resource
        )

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return resource
        return dashboard_pb2.MonitoringBetaDashboardRowLayoutRowsWidgetsScorecardTimeSeriesQueryTimeSeriesFilterAggregationCrossSeriesReducerEnum.Name(
            resource
        )[
            len(
                "MonitoringBetaDashboardRowLayoutRowsWidgetsScorecardTimeSeriesQueryTimeSeriesFilterAggregationCrossSeriesReducerEnum"
            ) :
        ]


class DashboardRowLayoutRowsWidgetsScorecardTimeSeriesQueryTimeSeriesFilterSecondaryAggregationPerSeriesAlignerEnum(
    object
):
    @classmethod
    def to_proto(self, resource):
        if not resource:
            return resource
        return dashboard_pb2.MonitoringBetaDashboardRowLayoutRowsWidgetsScorecardTimeSeriesQueryTimeSeriesFilterSecondaryAggregationPerSeriesAlignerEnum.Value(
            "MonitoringBetaDashboardRowLayoutRowsWidgetsScorecardTimeSeriesQueryTimeSeriesFilterSecondaryAggregationPerSeriesAlignerEnum%s"
            % resource
        )

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return resource
        return dashboard_pb2.MonitoringBetaDashboardRowLayoutRowsWidgetsScorecardTimeSeriesQueryTimeSeriesFilterSecondaryAggregationPerSeriesAlignerEnum.Name(
            resource
        )[
            len(
                "MonitoringBetaDashboardRowLayoutRowsWidgetsScorecardTimeSeriesQueryTimeSeriesFilterSecondaryAggregationPerSeriesAlignerEnum"
            ) :
        ]


class DashboardRowLayoutRowsWidgetsScorecardTimeSeriesQueryTimeSeriesFilterSecondaryAggregationCrossSeriesReducerEnum(
    object
):
    @classmethod
    def to_proto(self, resource):
        if not resource:
            return resource
        return dashboard_pb2.MonitoringBetaDashboardRowLayoutRowsWidgetsScorecardTimeSeriesQueryTimeSeriesFilterSecondaryAggregationCrossSeriesReducerEnum.Value(
            "MonitoringBetaDashboardRowLayoutRowsWidgetsScorecardTimeSeriesQueryTimeSeriesFilterSecondaryAggregationCrossSeriesReducerEnum%s"
            % resource
        )

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return resource
        return dashboard_pb2.MonitoringBetaDashboardRowLayoutRowsWidgetsScorecardTimeSeriesQueryTimeSeriesFilterSecondaryAggregationCrossSeriesReducerEnum.Name(
            resource
        )[
            len(
                "MonitoringBetaDashboardRowLayoutRowsWidgetsScorecardTimeSeriesQueryTimeSeriesFilterSecondaryAggregationCrossSeriesReducerEnum"
            ) :
        ]


class DashboardRowLayoutRowsWidgetsScorecardTimeSeriesQueryTimeSeriesFilterPickTimeSeriesFilterRankingMethodEnum(
    object
):
    @classmethod
    def to_proto(self, resource):
        if not resource:
            return resource
        return dashboard_pb2.MonitoringBetaDashboardRowLayoutRowsWidgetsScorecardTimeSeriesQueryTimeSeriesFilterPickTimeSeriesFilterRankingMethodEnum.Value(
            "MonitoringBetaDashboardRowLayoutRowsWidgetsScorecardTimeSeriesQueryTimeSeriesFilterPickTimeSeriesFilterRankingMethodEnum%s"
            % resource
        )

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return resource
        return dashboard_pb2.MonitoringBetaDashboardRowLayoutRowsWidgetsScorecardTimeSeriesQueryTimeSeriesFilterPickTimeSeriesFilterRankingMethodEnum.Name(
            resource
        )[
            len(
                "MonitoringBetaDashboardRowLayoutRowsWidgetsScorecardTimeSeriesQueryTimeSeriesFilterPickTimeSeriesFilterRankingMethodEnum"
            ) :
        ]


class DashboardRowLayoutRowsWidgetsScorecardTimeSeriesQueryTimeSeriesFilterPickTimeSeriesFilterDirectionEnum(
    object
):
    @classmethod
    def to_proto(self, resource):
        if not resource:
            return resource
        return dashboard_pb2.MonitoringBetaDashboardRowLayoutRowsWidgetsScorecardTimeSeriesQueryTimeSeriesFilterPickTimeSeriesFilterDirectionEnum.Value(
            "MonitoringBetaDashboardRowLayoutRowsWidgetsScorecardTimeSeriesQueryTimeSeriesFilterPickTimeSeriesFilterDirectionEnum%s"
            % resource
        )

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return resource
        return dashboard_pb2.MonitoringBetaDashboardRowLayoutRowsWidgetsScorecardTimeSeriesQueryTimeSeriesFilterPickTimeSeriesFilterDirectionEnum.Name(
            resource
        )[
            len(
                "MonitoringBetaDashboardRowLayoutRowsWidgetsScorecardTimeSeriesQueryTimeSeriesFilterPickTimeSeriesFilterDirectionEnum"
            ) :
        ]


class DashboardRowLayoutRowsWidgetsScorecardTimeSeriesQueryTimeSeriesFilterRatioNumeratorAggregationPerSeriesAlignerEnum(
    object
):
    @classmethod
    def to_proto(self, resource):
        if not resource:
            return resource
        return dashboard_pb2.MonitoringBetaDashboardRowLayoutRowsWidgetsScorecardTimeSeriesQueryTimeSeriesFilterRatioNumeratorAggregationPerSeriesAlignerEnum.Value(
            "MonitoringBetaDashboardRowLayoutRowsWidgetsScorecardTimeSeriesQueryTimeSeriesFilterRatioNumeratorAggregationPerSeriesAlignerEnum%s"
            % resource
        )

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return resource
        return dashboard_pb2.MonitoringBetaDashboardRowLayoutRowsWidgetsScorecardTimeSeriesQueryTimeSeriesFilterRatioNumeratorAggregationPerSeriesAlignerEnum.Name(
            resource
        )[
            len(
                "MonitoringBetaDashboardRowLayoutRowsWidgetsScorecardTimeSeriesQueryTimeSeriesFilterRatioNumeratorAggregationPerSeriesAlignerEnum"
            ) :
        ]


class DashboardRowLayoutRowsWidgetsScorecardTimeSeriesQueryTimeSeriesFilterRatioNumeratorAggregationCrossSeriesReducerEnum(
    object
):
    @classmethod
    def to_proto(self, resource):
        if not resource:
            return resource
        return dashboard_pb2.MonitoringBetaDashboardRowLayoutRowsWidgetsScorecardTimeSeriesQueryTimeSeriesFilterRatioNumeratorAggregationCrossSeriesReducerEnum.Value(
            "MonitoringBetaDashboardRowLayoutRowsWidgetsScorecardTimeSeriesQueryTimeSeriesFilterRatioNumeratorAggregationCrossSeriesReducerEnum%s"
            % resource
        )

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return resource
        return dashboard_pb2.MonitoringBetaDashboardRowLayoutRowsWidgetsScorecardTimeSeriesQueryTimeSeriesFilterRatioNumeratorAggregationCrossSeriesReducerEnum.Name(
            resource
        )[
            len(
                "MonitoringBetaDashboardRowLayoutRowsWidgetsScorecardTimeSeriesQueryTimeSeriesFilterRatioNumeratorAggregationCrossSeriesReducerEnum"
            ) :
        ]


class DashboardRowLayoutRowsWidgetsScorecardTimeSeriesQueryTimeSeriesFilterRatioDenominatorAggregationPerSeriesAlignerEnum(
    object
):
    @classmethod
    def to_proto(self, resource):
        if not resource:
            return resource
        return dashboard_pb2.MonitoringBetaDashboardRowLayoutRowsWidgetsScorecardTimeSeriesQueryTimeSeriesFilterRatioDenominatorAggregationPerSeriesAlignerEnum.Value(
            "MonitoringBetaDashboardRowLayoutRowsWidgetsScorecardTimeSeriesQueryTimeSeriesFilterRatioDenominatorAggregationPerSeriesAlignerEnum%s"
            % resource
        )

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return resource
        return dashboard_pb2.MonitoringBetaDashboardRowLayoutRowsWidgetsScorecardTimeSeriesQueryTimeSeriesFilterRatioDenominatorAggregationPerSeriesAlignerEnum.Name(
            resource
        )[
            len(
                "MonitoringBetaDashboardRowLayoutRowsWidgetsScorecardTimeSeriesQueryTimeSeriesFilterRatioDenominatorAggregationPerSeriesAlignerEnum"
            ) :
        ]


class DashboardRowLayoutRowsWidgetsScorecardTimeSeriesQueryTimeSeriesFilterRatioDenominatorAggregationCrossSeriesReducerEnum(
    object
):
    @classmethod
    def to_proto(self, resource):
        if not resource:
            return resource
        return dashboard_pb2.MonitoringBetaDashboardRowLayoutRowsWidgetsScorecardTimeSeriesQueryTimeSeriesFilterRatioDenominatorAggregationCrossSeriesReducerEnum.Value(
            "MonitoringBetaDashboardRowLayoutRowsWidgetsScorecardTimeSeriesQueryTimeSeriesFilterRatioDenominatorAggregationCrossSeriesReducerEnum%s"
            % resource
        )

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return resource
        return dashboard_pb2.MonitoringBetaDashboardRowLayoutRowsWidgetsScorecardTimeSeriesQueryTimeSeriesFilterRatioDenominatorAggregationCrossSeriesReducerEnum.Name(
            resource
        )[
            len(
                "MonitoringBetaDashboardRowLayoutRowsWidgetsScorecardTimeSeriesQueryTimeSeriesFilterRatioDenominatorAggregationCrossSeriesReducerEnum"
            ) :
        ]


class DashboardRowLayoutRowsWidgetsScorecardTimeSeriesQueryTimeSeriesFilterRatioSecondaryAggregationPerSeriesAlignerEnum(
    object
):
    @classmethod
    def to_proto(self, resource):
        if not resource:
            return resource
        return dashboard_pb2.MonitoringBetaDashboardRowLayoutRowsWidgetsScorecardTimeSeriesQueryTimeSeriesFilterRatioSecondaryAggregationPerSeriesAlignerEnum.Value(
            "MonitoringBetaDashboardRowLayoutRowsWidgetsScorecardTimeSeriesQueryTimeSeriesFilterRatioSecondaryAggregationPerSeriesAlignerEnum%s"
            % resource
        )

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return resource
        return dashboard_pb2.MonitoringBetaDashboardRowLayoutRowsWidgetsScorecardTimeSeriesQueryTimeSeriesFilterRatioSecondaryAggregationPerSeriesAlignerEnum.Name(
            resource
        )[
            len(
                "MonitoringBetaDashboardRowLayoutRowsWidgetsScorecardTimeSeriesQueryTimeSeriesFilterRatioSecondaryAggregationPerSeriesAlignerEnum"
            ) :
        ]


class DashboardRowLayoutRowsWidgetsScorecardTimeSeriesQueryTimeSeriesFilterRatioSecondaryAggregationCrossSeriesReducerEnum(
    object
):
    @classmethod
    def to_proto(self, resource):
        if not resource:
            return resource
        return dashboard_pb2.MonitoringBetaDashboardRowLayoutRowsWidgetsScorecardTimeSeriesQueryTimeSeriesFilterRatioSecondaryAggregationCrossSeriesReducerEnum.Value(
            "MonitoringBetaDashboardRowLayoutRowsWidgetsScorecardTimeSeriesQueryTimeSeriesFilterRatioSecondaryAggregationCrossSeriesReducerEnum%s"
            % resource
        )

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return resource
        return dashboard_pb2.MonitoringBetaDashboardRowLayoutRowsWidgetsScorecardTimeSeriesQueryTimeSeriesFilterRatioSecondaryAggregationCrossSeriesReducerEnum.Name(
            resource
        )[
            len(
                "MonitoringBetaDashboardRowLayoutRowsWidgetsScorecardTimeSeriesQueryTimeSeriesFilterRatioSecondaryAggregationCrossSeriesReducerEnum"
            ) :
        ]


class DashboardRowLayoutRowsWidgetsScorecardTimeSeriesQueryTimeSeriesFilterRatioPickTimeSeriesFilterRankingMethodEnum(
    object
):
    @classmethod
    def to_proto(self, resource):
        if not resource:
            return resource
        return dashboard_pb2.MonitoringBetaDashboardRowLayoutRowsWidgetsScorecardTimeSeriesQueryTimeSeriesFilterRatioPickTimeSeriesFilterRankingMethodEnum.Value(
            "MonitoringBetaDashboardRowLayoutRowsWidgetsScorecardTimeSeriesQueryTimeSeriesFilterRatioPickTimeSeriesFilterRankingMethodEnum%s"
            % resource
        )

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return resource
        return dashboard_pb2.MonitoringBetaDashboardRowLayoutRowsWidgetsScorecardTimeSeriesQueryTimeSeriesFilterRatioPickTimeSeriesFilterRankingMethodEnum.Name(
            resource
        )[
            len(
                "MonitoringBetaDashboardRowLayoutRowsWidgetsScorecardTimeSeriesQueryTimeSeriesFilterRatioPickTimeSeriesFilterRankingMethodEnum"
            ) :
        ]


class DashboardRowLayoutRowsWidgetsScorecardTimeSeriesQueryTimeSeriesFilterRatioPickTimeSeriesFilterDirectionEnum(
    object
):
    @classmethod
    def to_proto(self, resource):
        if not resource:
            return resource
        return dashboard_pb2.MonitoringBetaDashboardRowLayoutRowsWidgetsScorecardTimeSeriesQueryTimeSeriesFilterRatioPickTimeSeriesFilterDirectionEnum.Value(
            "MonitoringBetaDashboardRowLayoutRowsWidgetsScorecardTimeSeriesQueryTimeSeriesFilterRatioPickTimeSeriesFilterDirectionEnum%s"
            % resource
        )

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return resource
        return dashboard_pb2.MonitoringBetaDashboardRowLayoutRowsWidgetsScorecardTimeSeriesQueryTimeSeriesFilterRatioPickTimeSeriesFilterDirectionEnum.Name(
            resource
        )[
            len(
                "MonitoringBetaDashboardRowLayoutRowsWidgetsScorecardTimeSeriesQueryTimeSeriesFilterRatioPickTimeSeriesFilterDirectionEnum"
            ) :
        ]


class DashboardRowLayoutRowsWidgetsScorecardSparkChartViewSparkChartTypeEnum(object):
    @classmethod
    def to_proto(self, resource):
        if not resource:
            return resource
        return dashboard_pb2.MonitoringBetaDashboardRowLayoutRowsWidgetsScorecardSparkChartViewSparkChartTypeEnum.Value(
            "MonitoringBetaDashboardRowLayoutRowsWidgetsScorecardSparkChartViewSparkChartTypeEnum%s"
            % resource
        )

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return resource
        return dashboard_pb2.MonitoringBetaDashboardRowLayoutRowsWidgetsScorecardSparkChartViewSparkChartTypeEnum.Name(
            resource
        )[
            len(
                "MonitoringBetaDashboardRowLayoutRowsWidgetsScorecardSparkChartViewSparkChartTypeEnum"
            ) :
        ]


class DashboardRowLayoutRowsWidgetsScorecardThresholdsColorEnum(object):
    @classmethod
    def to_proto(self, resource):
        if not resource:
            return resource
        return dashboard_pb2.MonitoringBetaDashboardRowLayoutRowsWidgetsScorecardThresholdsColorEnum.Value(
            "MonitoringBetaDashboardRowLayoutRowsWidgetsScorecardThresholdsColorEnum%s"
            % resource
        )

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return resource
        return dashboard_pb2.MonitoringBetaDashboardRowLayoutRowsWidgetsScorecardThresholdsColorEnum.Name(
            resource
        )[
            len(
                "MonitoringBetaDashboardRowLayoutRowsWidgetsScorecardThresholdsColorEnum"
            ) :
        ]


class DashboardRowLayoutRowsWidgetsScorecardThresholdsDirectionEnum(object):
    @classmethod
    def to_proto(self, resource):
        if not resource:
            return resource
        return dashboard_pb2.MonitoringBetaDashboardRowLayoutRowsWidgetsScorecardThresholdsDirectionEnum.Value(
            "MonitoringBetaDashboardRowLayoutRowsWidgetsScorecardThresholdsDirectionEnum%s"
            % resource
        )

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return resource
        return dashboard_pb2.MonitoringBetaDashboardRowLayoutRowsWidgetsScorecardThresholdsDirectionEnum.Name(
            resource
        )[
            len(
                "MonitoringBetaDashboardRowLayoutRowsWidgetsScorecardThresholdsDirectionEnum"
            ) :
        ]


class DashboardRowLayoutRowsWidgetsTextFormatEnum(object):
    @classmethod
    def to_proto(self, resource):
        if not resource:
            return resource
        return dashboard_pb2.MonitoringBetaDashboardRowLayoutRowsWidgetsTextFormatEnum.Value(
            "MonitoringBetaDashboardRowLayoutRowsWidgetsTextFormatEnum%s" % resource
        )

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return resource
        return dashboard_pb2.MonitoringBetaDashboardRowLayoutRowsWidgetsTextFormatEnum.Name(
            resource
        )[
            len("MonitoringBetaDashboardRowLayoutRowsWidgetsTextFormatEnum") :
        ]


class DashboardColumnLayoutColumnsWidgetsXyChartDataSetsTimeSeriesQueryTimeSeriesFilterAggregationPerSeriesAlignerEnum(
    object
):
    @classmethod
    def to_proto(self, resource):
        if not resource:
            return resource
        return dashboard_pb2.MonitoringBetaDashboardColumnLayoutColumnsWidgetsXyChartDataSetsTimeSeriesQueryTimeSeriesFilterAggregationPerSeriesAlignerEnum.Value(
            "MonitoringBetaDashboardColumnLayoutColumnsWidgetsXyChartDataSetsTimeSeriesQueryTimeSeriesFilterAggregationPerSeriesAlignerEnum%s"
            % resource
        )

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return resource
        return dashboard_pb2.MonitoringBetaDashboardColumnLayoutColumnsWidgetsXyChartDataSetsTimeSeriesQueryTimeSeriesFilterAggregationPerSeriesAlignerEnum.Name(
            resource
        )[
            len(
                "MonitoringBetaDashboardColumnLayoutColumnsWidgetsXyChartDataSetsTimeSeriesQueryTimeSeriesFilterAggregationPerSeriesAlignerEnum"
            ) :
        ]


class DashboardColumnLayoutColumnsWidgetsXyChartDataSetsTimeSeriesQueryTimeSeriesFilterAggregationCrossSeriesReducerEnum(
    object
):
    @classmethod
    def to_proto(self, resource):
        if not resource:
            return resource
        return dashboard_pb2.MonitoringBetaDashboardColumnLayoutColumnsWidgetsXyChartDataSetsTimeSeriesQueryTimeSeriesFilterAggregationCrossSeriesReducerEnum.Value(
            "MonitoringBetaDashboardColumnLayoutColumnsWidgetsXyChartDataSetsTimeSeriesQueryTimeSeriesFilterAggregationCrossSeriesReducerEnum%s"
            % resource
        )

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return resource
        return dashboard_pb2.MonitoringBetaDashboardColumnLayoutColumnsWidgetsXyChartDataSetsTimeSeriesQueryTimeSeriesFilterAggregationCrossSeriesReducerEnum.Name(
            resource
        )[
            len(
                "MonitoringBetaDashboardColumnLayoutColumnsWidgetsXyChartDataSetsTimeSeriesQueryTimeSeriesFilterAggregationCrossSeriesReducerEnum"
            ) :
        ]


class DashboardColumnLayoutColumnsWidgetsXyChartDataSetsTimeSeriesQueryTimeSeriesFilterSecondaryAggregationPerSeriesAlignerEnum(
    object
):
    @classmethod
    def to_proto(self, resource):
        if not resource:
            return resource
        return dashboard_pb2.MonitoringBetaDashboardColumnLayoutColumnsWidgetsXyChartDataSetsTimeSeriesQueryTimeSeriesFilterSecondaryAggregationPerSeriesAlignerEnum.Value(
            "MonitoringBetaDashboardColumnLayoutColumnsWidgetsXyChartDataSetsTimeSeriesQueryTimeSeriesFilterSecondaryAggregationPerSeriesAlignerEnum%s"
            % resource
        )

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return resource
        return dashboard_pb2.MonitoringBetaDashboardColumnLayoutColumnsWidgetsXyChartDataSetsTimeSeriesQueryTimeSeriesFilterSecondaryAggregationPerSeriesAlignerEnum.Name(
            resource
        )[
            len(
                "MonitoringBetaDashboardColumnLayoutColumnsWidgetsXyChartDataSetsTimeSeriesQueryTimeSeriesFilterSecondaryAggregationPerSeriesAlignerEnum"
            ) :
        ]


class DashboardColumnLayoutColumnsWidgetsXyChartDataSetsTimeSeriesQueryTimeSeriesFilterSecondaryAggregationCrossSeriesReducerEnum(
    object
):
    @classmethod
    def to_proto(self, resource):
        if not resource:
            return resource
        return dashboard_pb2.MonitoringBetaDashboardColumnLayoutColumnsWidgetsXyChartDataSetsTimeSeriesQueryTimeSeriesFilterSecondaryAggregationCrossSeriesReducerEnum.Value(
            "MonitoringBetaDashboardColumnLayoutColumnsWidgetsXyChartDataSetsTimeSeriesQueryTimeSeriesFilterSecondaryAggregationCrossSeriesReducerEnum%s"
            % resource
        )

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return resource
        return dashboard_pb2.MonitoringBetaDashboardColumnLayoutColumnsWidgetsXyChartDataSetsTimeSeriesQueryTimeSeriesFilterSecondaryAggregationCrossSeriesReducerEnum.Name(
            resource
        )[
            len(
                "MonitoringBetaDashboardColumnLayoutColumnsWidgetsXyChartDataSetsTimeSeriesQueryTimeSeriesFilterSecondaryAggregationCrossSeriesReducerEnum"
            ) :
        ]


class DashboardColumnLayoutColumnsWidgetsXyChartDataSetsTimeSeriesQueryTimeSeriesFilterPickTimeSeriesFilterRankingMethodEnum(
    object
):
    @classmethod
    def to_proto(self, resource):
        if not resource:
            return resource
        return dashboard_pb2.MonitoringBetaDashboardColumnLayoutColumnsWidgetsXyChartDataSetsTimeSeriesQueryTimeSeriesFilterPickTimeSeriesFilterRankingMethodEnum.Value(
            "MonitoringBetaDashboardColumnLayoutColumnsWidgetsXyChartDataSetsTimeSeriesQueryTimeSeriesFilterPickTimeSeriesFilterRankingMethodEnum%s"
            % resource
        )

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return resource
        return dashboard_pb2.MonitoringBetaDashboardColumnLayoutColumnsWidgetsXyChartDataSetsTimeSeriesQueryTimeSeriesFilterPickTimeSeriesFilterRankingMethodEnum.Name(
            resource
        )[
            len(
                "MonitoringBetaDashboardColumnLayoutColumnsWidgetsXyChartDataSetsTimeSeriesQueryTimeSeriesFilterPickTimeSeriesFilterRankingMethodEnum"
            ) :
        ]


class DashboardColumnLayoutColumnsWidgetsXyChartDataSetsTimeSeriesQueryTimeSeriesFilterPickTimeSeriesFilterDirectionEnum(
    object
):
    @classmethod
    def to_proto(self, resource):
        if not resource:
            return resource
        return dashboard_pb2.MonitoringBetaDashboardColumnLayoutColumnsWidgetsXyChartDataSetsTimeSeriesQueryTimeSeriesFilterPickTimeSeriesFilterDirectionEnum.Value(
            "MonitoringBetaDashboardColumnLayoutColumnsWidgetsXyChartDataSetsTimeSeriesQueryTimeSeriesFilterPickTimeSeriesFilterDirectionEnum%s"
            % resource
        )

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return resource
        return dashboard_pb2.MonitoringBetaDashboardColumnLayoutColumnsWidgetsXyChartDataSetsTimeSeriesQueryTimeSeriesFilterPickTimeSeriesFilterDirectionEnum.Name(
            resource
        )[
            len(
                "MonitoringBetaDashboardColumnLayoutColumnsWidgetsXyChartDataSetsTimeSeriesQueryTimeSeriesFilterPickTimeSeriesFilterDirectionEnum"
            ) :
        ]


class DashboardColumnLayoutColumnsWidgetsXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioNumeratorAggregationPerSeriesAlignerEnum(
    object
):
    @classmethod
    def to_proto(self, resource):
        if not resource:
            return resource
        return dashboard_pb2.MonitoringBetaDashboardColumnLayoutColumnsWidgetsXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioNumeratorAggregationPerSeriesAlignerEnum.Value(
            "MonitoringBetaDashboardColumnLayoutColumnsWidgetsXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioNumeratorAggregationPerSeriesAlignerEnum%s"
            % resource
        )

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return resource
        return dashboard_pb2.MonitoringBetaDashboardColumnLayoutColumnsWidgetsXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioNumeratorAggregationPerSeriesAlignerEnum.Name(
            resource
        )[
            len(
                "MonitoringBetaDashboardColumnLayoutColumnsWidgetsXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioNumeratorAggregationPerSeriesAlignerEnum"
            ) :
        ]


class DashboardColumnLayoutColumnsWidgetsXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioNumeratorAggregationCrossSeriesReducerEnum(
    object
):
    @classmethod
    def to_proto(self, resource):
        if not resource:
            return resource
        return dashboard_pb2.MonitoringBetaDashboardColumnLayoutColumnsWidgetsXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioNumeratorAggregationCrossSeriesReducerEnum.Value(
            "MonitoringBetaDashboardColumnLayoutColumnsWidgetsXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioNumeratorAggregationCrossSeriesReducerEnum%s"
            % resource
        )

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return resource
        return dashboard_pb2.MonitoringBetaDashboardColumnLayoutColumnsWidgetsXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioNumeratorAggregationCrossSeriesReducerEnum.Name(
            resource
        )[
            len(
                "MonitoringBetaDashboardColumnLayoutColumnsWidgetsXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioNumeratorAggregationCrossSeriesReducerEnum"
            ) :
        ]


class DashboardColumnLayoutColumnsWidgetsXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioDenominatorAggregationPerSeriesAlignerEnum(
    object
):
    @classmethod
    def to_proto(self, resource):
        if not resource:
            return resource
        return dashboard_pb2.MonitoringBetaDashboardColumnLayoutColumnsWidgetsXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioDenominatorAggregationPerSeriesAlignerEnum.Value(
            "MonitoringBetaDashboardColumnLayoutColumnsWidgetsXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioDenominatorAggregationPerSeriesAlignerEnum%s"
            % resource
        )

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return resource
        return dashboard_pb2.MonitoringBetaDashboardColumnLayoutColumnsWidgetsXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioDenominatorAggregationPerSeriesAlignerEnum.Name(
            resource
        )[
            len(
                "MonitoringBetaDashboardColumnLayoutColumnsWidgetsXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioDenominatorAggregationPerSeriesAlignerEnum"
            ) :
        ]


class DashboardColumnLayoutColumnsWidgetsXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioDenominatorAggregationCrossSeriesReducerEnum(
    object
):
    @classmethod
    def to_proto(self, resource):
        if not resource:
            return resource
        return dashboard_pb2.MonitoringBetaDashboardColumnLayoutColumnsWidgetsXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioDenominatorAggregationCrossSeriesReducerEnum.Value(
            "MonitoringBetaDashboardColumnLayoutColumnsWidgetsXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioDenominatorAggregationCrossSeriesReducerEnum%s"
            % resource
        )

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return resource
        return dashboard_pb2.MonitoringBetaDashboardColumnLayoutColumnsWidgetsXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioDenominatorAggregationCrossSeriesReducerEnum.Name(
            resource
        )[
            len(
                "MonitoringBetaDashboardColumnLayoutColumnsWidgetsXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioDenominatorAggregationCrossSeriesReducerEnum"
            ) :
        ]


class DashboardColumnLayoutColumnsWidgetsXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioSecondaryAggregationPerSeriesAlignerEnum(
    object
):
    @classmethod
    def to_proto(self, resource):
        if not resource:
            return resource
        return dashboard_pb2.MonitoringBetaDashboardColumnLayoutColumnsWidgetsXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioSecondaryAggregationPerSeriesAlignerEnum.Value(
            "MonitoringBetaDashboardColumnLayoutColumnsWidgetsXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioSecondaryAggregationPerSeriesAlignerEnum%s"
            % resource
        )

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return resource
        return dashboard_pb2.MonitoringBetaDashboardColumnLayoutColumnsWidgetsXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioSecondaryAggregationPerSeriesAlignerEnum.Name(
            resource
        )[
            len(
                "MonitoringBetaDashboardColumnLayoutColumnsWidgetsXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioSecondaryAggregationPerSeriesAlignerEnum"
            ) :
        ]


class DashboardColumnLayoutColumnsWidgetsXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioSecondaryAggregationCrossSeriesReducerEnum(
    object
):
    @classmethod
    def to_proto(self, resource):
        if not resource:
            return resource
        return dashboard_pb2.MonitoringBetaDashboardColumnLayoutColumnsWidgetsXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioSecondaryAggregationCrossSeriesReducerEnum.Value(
            "MonitoringBetaDashboardColumnLayoutColumnsWidgetsXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioSecondaryAggregationCrossSeriesReducerEnum%s"
            % resource
        )

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return resource
        return dashboard_pb2.MonitoringBetaDashboardColumnLayoutColumnsWidgetsXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioSecondaryAggregationCrossSeriesReducerEnum.Name(
            resource
        )[
            len(
                "MonitoringBetaDashboardColumnLayoutColumnsWidgetsXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioSecondaryAggregationCrossSeriesReducerEnum"
            ) :
        ]


class DashboardColumnLayoutColumnsWidgetsXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioPickTimeSeriesFilterRankingMethodEnum(
    object
):
    @classmethod
    def to_proto(self, resource):
        if not resource:
            return resource
        return dashboard_pb2.MonitoringBetaDashboardColumnLayoutColumnsWidgetsXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioPickTimeSeriesFilterRankingMethodEnum.Value(
            "MonitoringBetaDashboardColumnLayoutColumnsWidgetsXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioPickTimeSeriesFilterRankingMethodEnum%s"
            % resource
        )

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return resource
        return dashboard_pb2.MonitoringBetaDashboardColumnLayoutColumnsWidgetsXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioPickTimeSeriesFilterRankingMethodEnum.Name(
            resource
        )[
            len(
                "MonitoringBetaDashboardColumnLayoutColumnsWidgetsXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioPickTimeSeriesFilterRankingMethodEnum"
            ) :
        ]


class DashboardColumnLayoutColumnsWidgetsXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioPickTimeSeriesFilterDirectionEnum(
    object
):
    @classmethod
    def to_proto(self, resource):
        if not resource:
            return resource
        return dashboard_pb2.MonitoringBetaDashboardColumnLayoutColumnsWidgetsXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioPickTimeSeriesFilterDirectionEnum.Value(
            "MonitoringBetaDashboardColumnLayoutColumnsWidgetsXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioPickTimeSeriesFilterDirectionEnum%s"
            % resource
        )

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return resource
        return dashboard_pb2.MonitoringBetaDashboardColumnLayoutColumnsWidgetsXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioPickTimeSeriesFilterDirectionEnum.Name(
            resource
        )[
            len(
                "MonitoringBetaDashboardColumnLayoutColumnsWidgetsXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioPickTimeSeriesFilterDirectionEnum"
            ) :
        ]


class DashboardColumnLayoutColumnsWidgetsXyChartDataSetsPlotTypeEnum(object):
    @classmethod
    def to_proto(self, resource):
        if not resource:
            return resource
        return dashboard_pb2.MonitoringBetaDashboardColumnLayoutColumnsWidgetsXyChartDataSetsPlotTypeEnum.Value(
            "MonitoringBetaDashboardColumnLayoutColumnsWidgetsXyChartDataSetsPlotTypeEnum%s"
            % resource
        )

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return resource
        return dashboard_pb2.MonitoringBetaDashboardColumnLayoutColumnsWidgetsXyChartDataSetsPlotTypeEnum.Name(
            resource
        )[
            len(
                "MonitoringBetaDashboardColumnLayoutColumnsWidgetsXyChartDataSetsPlotTypeEnum"
            ) :
        ]


class DashboardColumnLayoutColumnsWidgetsXyChartThresholdsColorEnum(object):
    @classmethod
    def to_proto(self, resource):
        if not resource:
            return resource
        return dashboard_pb2.MonitoringBetaDashboardColumnLayoutColumnsWidgetsXyChartThresholdsColorEnum.Value(
            "MonitoringBetaDashboardColumnLayoutColumnsWidgetsXyChartThresholdsColorEnum%s"
            % resource
        )

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return resource
        return dashboard_pb2.MonitoringBetaDashboardColumnLayoutColumnsWidgetsXyChartThresholdsColorEnum.Name(
            resource
        )[
            len(
                "MonitoringBetaDashboardColumnLayoutColumnsWidgetsXyChartThresholdsColorEnum"
            ) :
        ]


class DashboardColumnLayoutColumnsWidgetsXyChartThresholdsDirectionEnum(object):
    @classmethod
    def to_proto(self, resource):
        if not resource:
            return resource
        return dashboard_pb2.MonitoringBetaDashboardColumnLayoutColumnsWidgetsXyChartThresholdsDirectionEnum.Value(
            "MonitoringBetaDashboardColumnLayoutColumnsWidgetsXyChartThresholdsDirectionEnum%s"
            % resource
        )

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return resource
        return dashboard_pb2.MonitoringBetaDashboardColumnLayoutColumnsWidgetsXyChartThresholdsDirectionEnum.Name(
            resource
        )[
            len(
                "MonitoringBetaDashboardColumnLayoutColumnsWidgetsXyChartThresholdsDirectionEnum"
            ) :
        ]


class DashboardColumnLayoutColumnsWidgetsXyChartXAxisScaleEnum(object):
    @classmethod
    def to_proto(self, resource):
        if not resource:
            return resource
        return dashboard_pb2.MonitoringBetaDashboardColumnLayoutColumnsWidgetsXyChartXAxisScaleEnum.Value(
            "MonitoringBetaDashboardColumnLayoutColumnsWidgetsXyChartXAxisScaleEnum%s"
            % resource
        )

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return resource
        return dashboard_pb2.MonitoringBetaDashboardColumnLayoutColumnsWidgetsXyChartXAxisScaleEnum.Name(
            resource
        )[
            len(
                "MonitoringBetaDashboardColumnLayoutColumnsWidgetsXyChartXAxisScaleEnum"
            ) :
        ]


class DashboardColumnLayoutColumnsWidgetsXyChartYAxisScaleEnum(object):
    @classmethod
    def to_proto(self, resource):
        if not resource:
            return resource
        return dashboard_pb2.MonitoringBetaDashboardColumnLayoutColumnsWidgetsXyChartYAxisScaleEnum.Value(
            "MonitoringBetaDashboardColumnLayoutColumnsWidgetsXyChartYAxisScaleEnum%s"
            % resource
        )

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return resource
        return dashboard_pb2.MonitoringBetaDashboardColumnLayoutColumnsWidgetsXyChartYAxisScaleEnum.Name(
            resource
        )[
            len(
                "MonitoringBetaDashboardColumnLayoutColumnsWidgetsXyChartYAxisScaleEnum"
            ) :
        ]


class DashboardColumnLayoutColumnsWidgetsXyChartChartOptionsModeEnum(object):
    @classmethod
    def to_proto(self, resource):
        if not resource:
            return resource
        return dashboard_pb2.MonitoringBetaDashboardColumnLayoutColumnsWidgetsXyChartChartOptionsModeEnum.Value(
            "MonitoringBetaDashboardColumnLayoutColumnsWidgetsXyChartChartOptionsModeEnum%s"
            % resource
        )

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return resource
        return dashboard_pb2.MonitoringBetaDashboardColumnLayoutColumnsWidgetsXyChartChartOptionsModeEnum.Name(
            resource
        )[
            len(
                "MonitoringBetaDashboardColumnLayoutColumnsWidgetsXyChartChartOptionsModeEnum"
            ) :
        ]


class DashboardColumnLayoutColumnsWidgetsScorecardTimeSeriesQueryTimeSeriesFilterAggregationPerSeriesAlignerEnum(
    object
):
    @classmethod
    def to_proto(self, resource):
        if not resource:
            return resource
        return dashboard_pb2.MonitoringBetaDashboardColumnLayoutColumnsWidgetsScorecardTimeSeriesQueryTimeSeriesFilterAggregationPerSeriesAlignerEnum.Value(
            "MonitoringBetaDashboardColumnLayoutColumnsWidgetsScorecardTimeSeriesQueryTimeSeriesFilterAggregationPerSeriesAlignerEnum%s"
            % resource
        )

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return resource
        return dashboard_pb2.MonitoringBetaDashboardColumnLayoutColumnsWidgetsScorecardTimeSeriesQueryTimeSeriesFilterAggregationPerSeriesAlignerEnum.Name(
            resource
        )[
            len(
                "MonitoringBetaDashboardColumnLayoutColumnsWidgetsScorecardTimeSeriesQueryTimeSeriesFilterAggregationPerSeriesAlignerEnum"
            ) :
        ]


class DashboardColumnLayoutColumnsWidgetsScorecardTimeSeriesQueryTimeSeriesFilterAggregationCrossSeriesReducerEnum(
    object
):
    @classmethod
    def to_proto(self, resource):
        if not resource:
            return resource
        return dashboard_pb2.MonitoringBetaDashboardColumnLayoutColumnsWidgetsScorecardTimeSeriesQueryTimeSeriesFilterAggregationCrossSeriesReducerEnum.Value(
            "MonitoringBetaDashboardColumnLayoutColumnsWidgetsScorecardTimeSeriesQueryTimeSeriesFilterAggregationCrossSeriesReducerEnum%s"
            % resource
        )

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return resource
        return dashboard_pb2.MonitoringBetaDashboardColumnLayoutColumnsWidgetsScorecardTimeSeriesQueryTimeSeriesFilterAggregationCrossSeriesReducerEnum.Name(
            resource
        )[
            len(
                "MonitoringBetaDashboardColumnLayoutColumnsWidgetsScorecardTimeSeriesQueryTimeSeriesFilterAggregationCrossSeriesReducerEnum"
            ) :
        ]


class DashboardColumnLayoutColumnsWidgetsScorecardTimeSeriesQueryTimeSeriesFilterSecondaryAggregationPerSeriesAlignerEnum(
    object
):
    @classmethod
    def to_proto(self, resource):
        if not resource:
            return resource
        return dashboard_pb2.MonitoringBetaDashboardColumnLayoutColumnsWidgetsScorecardTimeSeriesQueryTimeSeriesFilterSecondaryAggregationPerSeriesAlignerEnum.Value(
            "MonitoringBetaDashboardColumnLayoutColumnsWidgetsScorecardTimeSeriesQueryTimeSeriesFilterSecondaryAggregationPerSeriesAlignerEnum%s"
            % resource
        )

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return resource
        return dashboard_pb2.MonitoringBetaDashboardColumnLayoutColumnsWidgetsScorecardTimeSeriesQueryTimeSeriesFilterSecondaryAggregationPerSeriesAlignerEnum.Name(
            resource
        )[
            len(
                "MonitoringBetaDashboardColumnLayoutColumnsWidgetsScorecardTimeSeriesQueryTimeSeriesFilterSecondaryAggregationPerSeriesAlignerEnum"
            ) :
        ]


class DashboardColumnLayoutColumnsWidgetsScorecardTimeSeriesQueryTimeSeriesFilterSecondaryAggregationCrossSeriesReducerEnum(
    object
):
    @classmethod
    def to_proto(self, resource):
        if not resource:
            return resource
        return dashboard_pb2.MonitoringBetaDashboardColumnLayoutColumnsWidgetsScorecardTimeSeriesQueryTimeSeriesFilterSecondaryAggregationCrossSeriesReducerEnum.Value(
            "MonitoringBetaDashboardColumnLayoutColumnsWidgetsScorecardTimeSeriesQueryTimeSeriesFilterSecondaryAggregationCrossSeriesReducerEnum%s"
            % resource
        )

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return resource
        return dashboard_pb2.MonitoringBetaDashboardColumnLayoutColumnsWidgetsScorecardTimeSeriesQueryTimeSeriesFilterSecondaryAggregationCrossSeriesReducerEnum.Name(
            resource
        )[
            len(
                "MonitoringBetaDashboardColumnLayoutColumnsWidgetsScorecardTimeSeriesQueryTimeSeriesFilterSecondaryAggregationCrossSeriesReducerEnum"
            ) :
        ]


class DashboardColumnLayoutColumnsWidgetsScorecardTimeSeriesQueryTimeSeriesFilterPickTimeSeriesFilterRankingMethodEnum(
    object
):
    @classmethod
    def to_proto(self, resource):
        if not resource:
            return resource
        return dashboard_pb2.MonitoringBetaDashboardColumnLayoutColumnsWidgetsScorecardTimeSeriesQueryTimeSeriesFilterPickTimeSeriesFilterRankingMethodEnum.Value(
            "MonitoringBetaDashboardColumnLayoutColumnsWidgetsScorecardTimeSeriesQueryTimeSeriesFilterPickTimeSeriesFilterRankingMethodEnum%s"
            % resource
        )

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return resource
        return dashboard_pb2.MonitoringBetaDashboardColumnLayoutColumnsWidgetsScorecardTimeSeriesQueryTimeSeriesFilterPickTimeSeriesFilterRankingMethodEnum.Name(
            resource
        )[
            len(
                "MonitoringBetaDashboardColumnLayoutColumnsWidgetsScorecardTimeSeriesQueryTimeSeriesFilterPickTimeSeriesFilterRankingMethodEnum"
            ) :
        ]


class DashboardColumnLayoutColumnsWidgetsScorecardTimeSeriesQueryTimeSeriesFilterPickTimeSeriesFilterDirectionEnum(
    object
):
    @classmethod
    def to_proto(self, resource):
        if not resource:
            return resource
        return dashboard_pb2.MonitoringBetaDashboardColumnLayoutColumnsWidgetsScorecardTimeSeriesQueryTimeSeriesFilterPickTimeSeriesFilterDirectionEnum.Value(
            "MonitoringBetaDashboardColumnLayoutColumnsWidgetsScorecardTimeSeriesQueryTimeSeriesFilterPickTimeSeriesFilterDirectionEnum%s"
            % resource
        )

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return resource
        return dashboard_pb2.MonitoringBetaDashboardColumnLayoutColumnsWidgetsScorecardTimeSeriesQueryTimeSeriesFilterPickTimeSeriesFilterDirectionEnum.Name(
            resource
        )[
            len(
                "MonitoringBetaDashboardColumnLayoutColumnsWidgetsScorecardTimeSeriesQueryTimeSeriesFilterPickTimeSeriesFilterDirectionEnum"
            ) :
        ]


class DashboardColumnLayoutColumnsWidgetsScorecardTimeSeriesQueryTimeSeriesFilterRatioNumeratorAggregationPerSeriesAlignerEnum(
    object
):
    @classmethod
    def to_proto(self, resource):
        if not resource:
            return resource
        return dashboard_pb2.MonitoringBetaDashboardColumnLayoutColumnsWidgetsScorecardTimeSeriesQueryTimeSeriesFilterRatioNumeratorAggregationPerSeriesAlignerEnum.Value(
            "MonitoringBetaDashboardColumnLayoutColumnsWidgetsScorecardTimeSeriesQueryTimeSeriesFilterRatioNumeratorAggregationPerSeriesAlignerEnum%s"
            % resource
        )

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return resource
        return dashboard_pb2.MonitoringBetaDashboardColumnLayoutColumnsWidgetsScorecardTimeSeriesQueryTimeSeriesFilterRatioNumeratorAggregationPerSeriesAlignerEnum.Name(
            resource
        )[
            len(
                "MonitoringBetaDashboardColumnLayoutColumnsWidgetsScorecardTimeSeriesQueryTimeSeriesFilterRatioNumeratorAggregationPerSeriesAlignerEnum"
            ) :
        ]


class DashboardColumnLayoutColumnsWidgetsScorecardTimeSeriesQueryTimeSeriesFilterRatioNumeratorAggregationCrossSeriesReducerEnum(
    object
):
    @classmethod
    def to_proto(self, resource):
        if not resource:
            return resource
        return dashboard_pb2.MonitoringBetaDashboardColumnLayoutColumnsWidgetsScorecardTimeSeriesQueryTimeSeriesFilterRatioNumeratorAggregationCrossSeriesReducerEnum.Value(
            "MonitoringBetaDashboardColumnLayoutColumnsWidgetsScorecardTimeSeriesQueryTimeSeriesFilterRatioNumeratorAggregationCrossSeriesReducerEnum%s"
            % resource
        )

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return resource
        return dashboard_pb2.MonitoringBetaDashboardColumnLayoutColumnsWidgetsScorecardTimeSeriesQueryTimeSeriesFilterRatioNumeratorAggregationCrossSeriesReducerEnum.Name(
            resource
        )[
            len(
                "MonitoringBetaDashboardColumnLayoutColumnsWidgetsScorecardTimeSeriesQueryTimeSeriesFilterRatioNumeratorAggregationCrossSeriesReducerEnum"
            ) :
        ]


class DashboardColumnLayoutColumnsWidgetsScorecardTimeSeriesQueryTimeSeriesFilterRatioDenominatorAggregationPerSeriesAlignerEnum(
    object
):
    @classmethod
    def to_proto(self, resource):
        if not resource:
            return resource
        return dashboard_pb2.MonitoringBetaDashboardColumnLayoutColumnsWidgetsScorecardTimeSeriesQueryTimeSeriesFilterRatioDenominatorAggregationPerSeriesAlignerEnum.Value(
            "MonitoringBetaDashboardColumnLayoutColumnsWidgetsScorecardTimeSeriesQueryTimeSeriesFilterRatioDenominatorAggregationPerSeriesAlignerEnum%s"
            % resource
        )

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return resource
        return dashboard_pb2.MonitoringBetaDashboardColumnLayoutColumnsWidgetsScorecardTimeSeriesQueryTimeSeriesFilterRatioDenominatorAggregationPerSeriesAlignerEnum.Name(
            resource
        )[
            len(
                "MonitoringBetaDashboardColumnLayoutColumnsWidgetsScorecardTimeSeriesQueryTimeSeriesFilterRatioDenominatorAggregationPerSeriesAlignerEnum"
            ) :
        ]


class DashboardColumnLayoutColumnsWidgetsScorecardTimeSeriesQueryTimeSeriesFilterRatioDenominatorAggregationCrossSeriesReducerEnum(
    object
):
    @classmethod
    def to_proto(self, resource):
        if not resource:
            return resource
        return dashboard_pb2.MonitoringBetaDashboardColumnLayoutColumnsWidgetsScorecardTimeSeriesQueryTimeSeriesFilterRatioDenominatorAggregationCrossSeriesReducerEnum.Value(
            "MonitoringBetaDashboardColumnLayoutColumnsWidgetsScorecardTimeSeriesQueryTimeSeriesFilterRatioDenominatorAggregationCrossSeriesReducerEnum%s"
            % resource
        )

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return resource
        return dashboard_pb2.MonitoringBetaDashboardColumnLayoutColumnsWidgetsScorecardTimeSeriesQueryTimeSeriesFilterRatioDenominatorAggregationCrossSeriesReducerEnum.Name(
            resource
        )[
            len(
                "MonitoringBetaDashboardColumnLayoutColumnsWidgetsScorecardTimeSeriesQueryTimeSeriesFilterRatioDenominatorAggregationCrossSeriesReducerEnum"
            ) :
        ]


class DashboardColumnLayoutColumnsWidgetsScorecardTimeSeriesQueryTimeSeriesFilterRatioSecondaryAggregationPerSeriesAlignerEnum(
    object
):
    @classmethod
    def to_proto(self, resource):
        if not resource:
            return resource
        return dashboard_pb2.MonitoringBetaDashboardColumnLayoutColumnsWidgetsScorecardTimeSeriesQueryTimeSeriesFilterRatioSecondaryAggregationPerSeriesAlignerEnum.Value(
            "MonitoringBetaDashboardColumnLayoutColumnsWidgetsScorecardTimeSeriesQueryTimeSeriesFilterRatioSecondaryAggregationPerSeriesAlignerEnum%s"
            % resource
        )

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return resource
        return dashboard_pb2.MonitoringBetaDashboardColumnLayoutColumnsWidgetsScorecardTimeSeriesQueryTimeSeriesFilterRatioSecondaryAggregationPerSeriesAlignerEnum.Name(
            resource
        )[
            len(
                "MonitoringBetaDashboardColumnLayoutColumnsWidgetsScorecardTimeSeriesQueryTimeSeriesFilterRatioSecondaryAggregationPerSeriesAlignerEnum"
            ) :
        ]


class DashboardColumnLayoutColumnsWidgetsScorecardTimeSeriesQueryTimeSeriesFilterRatioSecondaryAggregationCrossSeriesReducerEnum(
    object
):
    @classmethod
    def to_proto(self, resource):
        if not resource:
            return resource
        return dashboard_pb2.MonitoringBetaDashboardColumnLayoutColumnsWidgetsScorecardTimeSeriesQueryTimeSeriesFilterRatioSecondaryAggregationCrossSeriesReducerEnum.Value(
            "MonitoringBetaDashboardColumnLayoutColumnsWidgetsScorecardTimeSeriesQueryTimeSeriesFilterRatioSecondaryAggregationCrossSeriesReducerEnum%s"
            % resource
        )

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return resource
        return dashboard_pb2.MonitoringBetaDashboardColumnLayoutColumnsWidgetsScorecardTimeSeriesQueryTimeSeriesFilterRatioSecondaryAggregationCrossSeriesReducerEnum.Name(
            resource
        )[
            len(
                "MonitoringBetaDashboardColumnLayoutColumnsWidgetsScorecardTimeSeriesQueryTimeSeriesFilterRatioSecondaryAggregationCrossSeriesReducerEnum"
            ) :
        ]


class DashboardColumnLayoutColumnsWidgetsScorecardTimeSeriesQueryTimeSeriesFilterRatioPickTimeSeriesFilterRankingMethodEnum(
    object
):
    @classmethod
    def to_proto(self, resource):
        if not resource:
            return resource
        return dashboard_pb2.MonitoringBetaDashboardColumnLayoutColumnsWidgetsScorecardTimeSeriesQueryTimeSeriesFilterRatioPickTimeSeriesFilterRankingMethodEnum.Value(
            "MonitoringBetaDashboardColumnLayoutColumnsWidgetsScorecardTimeSeriesQueryTimeSeriesFilterRatioPickTimeSeriesFilterRankingMethodEnum%s"
            % resource
        )

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return resource
        return dashboard_pb2.MonitoringBetaDashboardColumnLayoutColumnsWidgetsScorecardTimeSeriesQueryTimeSeriesFilterRatioPickTimeSeriesFilterRankingMethodEnum.Name(
            resource
        )[
            len(
                "MonitoringBetaDashboardColumnLayoutColumnsWidgetsScorecardTimeSeriesQueryTimeSeriesFilterRatioPickTimeSeriesFilterRankingMethodEnum"
            ) :
        ]


class DashboardColumnLayoutColumnsWidgetsScorecardTimeSeriesQueryTimeSeriesFilterRatioPickTimeSeriesFilterDirectionEnum(
    object
):
    @classmethod
    def to_proto(self, resource):
        if not resource:
            return resource
        return dashboard_pb2.MonitoringBetaDashboardColumnLayoutColumnsWidgetsScorecardTimeSeriesQueryTimeSeriesFilterRatioPickTimeSeriesFilterDirectionEnum.Value(
            "MonitoringBetaDashboardColumnLayoutColumnsWidgetsScorecardTimeSeriesQueryTimeSeriesFilterRatioPickTimeSeriesFilterDirectionEnum%s"
            % resource
        )

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return resource
        return dashboard_pb2.MonitoringBetaDashboardColumnLayoutColumnsWidgetsScorecardTimeSeriesQueryTimeSeriesFilterRatioPickTimeSeriesFilterDirectionEnum.Name(
            resource
        )[
            len(
                "MonitoringBetaDashboardColumnLayoutColumnsWidgetsScorecardTimeSeriesQueryTimeSeriesFilterRatioPickTimeSeriesFilterDirectionEnum"
            ) :
        ]


class DashboardColumnLayoutColumnsWidgetsScorecardSparkChartViewSparkChartTypeEnum(
    object
):
    @classmethod
    def to_proto(self, resource):
        if not resource:
            return resource
        return dashboard_pb2.MonitoringBetaDashboardColumnLayoutColumnsWidgetsScorecardSparkChartViewSparkChartTypeEnum.Value(
            "MonitoringBetaDashboardColumnLayoutColumnsWidgetsScorecardSparkChartViewSparkChartTypeEnum%s"
            % resource
        )

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return resource
        return dashboard_pb2.MonitoringBetaDashboardColumnLayoutColumnsWidgetsScorecardSparkChartViewSparkChartTypeEnum.Name(
            resource
        )[
            len(
                "MonitoringBetaDashboardColumnLayoutColumnsWidgetsScorecardSparkChartViewSparkChartTypeEnum"
            ) :
        ]


class DashboardColumnLayoutColumnsWidgetsScorecardThresholdsColorEnum(object):
    @classmethod
    def to_proto(self, resource):
        if not resource:
            return resource
        return dashboard_pb2.MonitoringBetaDashboardColumnLayoutColumnsWidgetsScorecardThresholdsColorEnum.Value(
            "MonitoringBetaDashboardColumnLayoutColumnsWidgetsScorecardThresholdsColorEnum%s"
            % resource
        )

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return resource
        return dashboard_pb2.MonitoringBetaDashboardColumnLayoutColumnsWidgetsScorecardThresholdsColorEnum.Name(
            resource
        )[
            len(
                "MonitoringBetaDashboardColumnLayoutColumnsWidgetsScorecardThresholdsColorEnum"
            ) :
        ]


class DashboardColumnLayoutColumnsWidgetsScorecardThresholdsDirectionEnum(object):
    @classmethod
    def to_proto(self, resource):
        if not resource:
            return resource
        return dashboard_pb2.MonitoringBetaDashboardColumnLayoutColumnsWidgetsScorecardThresholdsDirectionEnum.Value(
            "MonitoringBetaDashboardColumnLayoutColumnsWidgetsScorecardThresholdsDirectionEnum%s"
            % resource
        )

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return resource
        return dashboard_pb2.MonitoringBetaDashboardColumnLayoutColumnsWidgetsScorecardThresholdsDirectionEnum.Name(
            resource
        )[
            len(
                "MonitoringBetaDashboardColumnLayoutColumnsWidgetsScorecardThresholdsDirectionEnum"
            ) :
        ]


class DashboardColumnLayoutColumnsWidgetsTextFormatEnum(object):
    @classmethod
    def to_proto(self, resource):
        if not resource:
            return resource
        return dashboard_pb2.MonitoringBetaDashboardColumnLayoutColumnsWidgetsTextFormatEnum.Value(
            "MonitoringBetaDashboardColumnLayoutColumnsWidgetsTextFormatEnum%s"
            % resource
        )

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return resource
        return dashboard_pb2.MonitoringBetaDashboardColumnLayoutColumnsWidgetsTextFormatEnum.Name(
            resource
        )[
            len("MonitoringBetaDashboardColumnLayoutColumnsWidgetsTextFormatEnum") :
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
