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
        stub = dashboard_pb2_grpc.MonitoringDashboardServiceStub(channel.Channel())
        request = dashboard_pb2.ApplyMonitoringDashboardRequest()
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

        response = stub.ApplyMonitoringDashboard(request)
        self.name = Primitive.from_proto(response.name)
        self.display_name = Primitive.from_proto(response.display_name)
        self.grid_layout = DashboardGridLayout.from_proto(response.grid_layout)
        self.mosaic_layout = DashboardMosaicLayout.from_proto(response.mosaic_layout)
        self.row_layout = DashboardRowLayout.from_proto(response.row_layout)
        self.column_layout = DashboardColumnLayout.from_proto(response.column_layout)
        self.project = Primitive.from_proto(response.project)
        self.etag = Primitive.from_proto(response.etag)

    def delete(self):
        stub = dashboard_pb2_grpc.MonitoringDashboardServiceStub(channel.Channel())
        request = dashboard_pb2.DeleteMonitoringDashboardRequest()
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

        response = stub.DeleteMonitoringDashboard(request)

    @classmethod
    def list(self, project, service_account_file=""):
        stub = dashboard_pb2_grpc.MonitoringDashboardServiceStub(channel.Channel())
        request = dashboard_pb2.ListMonitoringDashboardRequest()
        request.service_account_file = service_account_file
        request.Project = project

        return stub.ListMonitoringDashboard(request).items

    def to_proto(self):
        resource = dashboard_pb2.MonitoringDashboard()
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

        res = dashboard_pb2.MonitoringDashboardGridLayout()
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

        res = dashboard_pb2.MonitoringDashboardGridLayoutWidgets()
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

        res = dashboard_pb2.MonitoringDashboardGridLayoutWidgetsXyChart()
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

        res = dashboard_pb2.MonitoringDashboardGridLayoutWidgetsXyChartDataSets()
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
            dashboard_pb2.MonitoringDashboardGridLayoutWidgetsXyChartDataSetsTimeSeriesQuery()
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
            dashboard_pb2.MonitoringDashboardGridLayoutWidgetsXyChartDataSetsTimeSeriesQueryTimeSeriesFilter()
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
            dashboard_pb2.MonitoringDashboardGridLayoutWidgetsXyChartDataSetsTimeSeriesQueryTimeSeriesFilterAggregation()
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
            dashboard_pb2.MonitoringDashboardGridLayoutWidgetsXyChartDataSetsTimeSeriesQueryTimeSeriesFilterSecondaryAggregation()
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
            dashboard_pb2.MonitoringDashboardGridLayoutWidgetsXyChartDataSetsTimeSeriesQueryTimeSeriesFilterPickTimeSeriesFilter()
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
            dashboard_pb2.MonitoringDashboardGridLayoutWidgetsXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatio()
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
            dashboard_pb2.MonitoringDashboardGridLayoutWidgetsXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioNumerator()
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
            dashboard_pb2.MonitoringDashboardGridLayoutWidgetsXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioNumeratorAggregation()
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
            dashboard_pb2.MonitoringDashboardGridLayoutWidgetsXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioDenominator()
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
            dashboard_pb2.MonitoringDashboardGridLayoutWidgetsXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioDenominatorAggregation()
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
            dashboard_pb2.MonitoringDashboardGridLayoutWidgetsXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioSecondaryAggregation()
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
            dashboard_pb2.MonitoringDashboardGridLayoutWidgetsXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioPickTimeSeriesFilter()
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

        res = dashboard_pb2.MonitoringDashboardGridLayoutWidgetsXyChartThresholds()
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

        res = dashboard_pb2.MonitoringDashboardGridLayoutWidgetsXyChartXAxis()
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

        res = dashboard_pb2.MonitoringDashboardGridLayoutWidgetsXyChartYAxis()
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

        res = dashboard_pb2.MonitoringDashboardGridLayoutWidgetsXyChartChartOptions()
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

        res = dashboard_pb2.MonitoringDashboardGridLayoutWidgetsScorecard()
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
            dashboard_pb2.MonitoringDashboardGridLayoutWidgetsScorecardTimeSeriesQuery()
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
            dashboard_pb2.MonitoringDashboardGridLayoutWidgetsScorecardTimeSeriesQueryTimeSeriesFilter()
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
            dashboard_pb2.MonitoringDashboardGridLayoutWidgetsScorecardTimeSeriesQueryTimeSeriesFilterAggregation()
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
            dashboard_pb2.MonitoringDashboardGridLayoutWidgetsScorecardTimeSeriesQueryTimeSeriesFilterSecondaryAggregation()
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
            dashboard_pb2.MonitoringDashboardGridLayoutWidgetsScorecardTimeSeriesQueryTimeSeriesFilterPickTimeSeriesFilter()
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
            dashboard_pb2.MonitoringDashboardGridLayoutWidgetsScorecardTimeSeriesQueryTimeSeriesFilterRatio()
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
            dashboard_pb2.MonitoringDashboardGridLayoutWidgetsScorecardTimeSeriesQueryTimeSeriesFilterRatioNumerator()
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
            dashboard_pb2.MonitoringDashboardGridLayoutWidgetsScorecardTimeSeriesQueryTimeSeriesFilterRatioNumeratorAggregation()
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
            dashboard_pb2.MonitoringDashboardGridLayoutWidgetsScorecardTimeSeriesQueryTimeSeriesFilterRatioDenominator()
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
            dashboard_pb2.MonitoringDashboardGridLayoutWidgetsScorecardTimeSeriesQueryTimeSeriesFilterRatioDenominatorAggregation()
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
            dashboard_pb2.MonitoringDashboardGridLayoutWidgetsScorecardTimeSeriesQueryTimeSeriesFilterRatioSecondaryAggregation()
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
            dashboard_pb2.MonitoringDashboardGridLayoutWidgetsScorecardTimeSeriesQueryTimeSeriesFilterRatioPickTimeSeriesFilter()
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

        res = dashboard_pb2.MonitoringDashboardGridLayoutWidgetsScorecardGaugeView()
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
            dashboard_pb2.MonitoringDashboardGridLayoutWidgetsScorecardSparkChartView()
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

        res = dashboard_pb2.MonitoringDashboardGridLayoutWidgetsScorecardThresholds()
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

        res = dashboard_pb2.MonitoringDashboardGridLayoutWidgetsText()
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

        res = dashboard_pb2.MonitoringDashboardGridLayoutWidgetsBlank()
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

        res = dashboard_pb2.MonitoringDashboardGridLayoutWidgetsLogsPanel()
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

        res = dashboard_pb2.MonitoringDashboardMosaicLayout()
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

        res = dashboard_pb2.MonitoringDashboardMosaicLayoutTiles()
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

        res = dashboard_pb2.MonitoringDashboardMosaicLayoutTilesWidget()
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

        res = dashboard_pb2.MonitoringDashboardMosaicLayoutTilesWidgetXyChart()
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

        res = dashboard_pb2.MonitoringDashboardMosaicLayoutTilesWidgetXyChartDataSets()
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
            dashboard_pb2.MonitoringDashboardMosaicLayoutTilesWidgetXyChartDataSetsTimeSeriesQuery()
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
            dashboard_pb2.MonitoringDashboardMosaicLayoutTilesWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilter()
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
            dashboard_pb2.MonitoringDashboardMosaicLayoutTilesWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterAggregation()
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
            dashboard_pb2.MonitoringDashboardMosaicLayoutTilesWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterSecondaryAggregation()
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
            dashboard_pb2.MonitoringDashboardMosaicLayoutTilesWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterPickTimeSeriesFilter()
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
            dashboard_pb2.MonitoringDashboardMosaicLayoutTilesWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatio()
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
            dashboard_pb2.MonitoringDashboardMosaicLayoutTilesWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioNumerator()
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
            dashboard_pb2.MonitoringDashboardMosaicLayoutTilesWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioNumeratorAggregation()
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
            dashboard_pb2.MonitoringDashboardMosaicLayoutTilesWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioDenominator()
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
            dashboard_pb2.MonitoringDashboardMosaicLayoutTilesWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioDenominatorAggregation()
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
            dashboard_pb2.MonitoringDashboardMosaicLayoutTilesWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioSecondaryAggregation()
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
            dashboard_pb2.MonitoringDashboardMosaicLayoutTilesWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioPickTimeSeriesFilter()
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
            dashboard_pb2.MonitoringDashboardMosaicLayoutTilesWidgetXyChartThresholds()
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

        res = dashboard_pb2.MonitoringDashboardMosaicLayoutTilesWidgetXyChartXAxis()
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

        res = dashboard_pb2.MonitoringDashboardMosaicLayoutTilesWidgetXyChartYAxis()
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
            dashboard_pb2.MonitoringDashboardMosaicLayoutTilesWidgetXyChartChartOptions()
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

        res = dashboard_pb2.MonitoringDashboardMosaicLayoutTilesWidgetScorecard()
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
            dashboard_pb2.MonitoringDashboardMosaicLayoutTilesWidgetScorecardTimeSeriesQuery()
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
            dashboard_pb2.MonitoringDashboardMosaicLayoutTilesWidgetScorecardTimeSeriesQueryTimeSeriesFilter()
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
            dashboard_pb2.MonitoringDashboardMosaicLayoutTilesWidgetScorecardTimeSeriesQueryTimeSeriesFilterAggregation()
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
            dashboard_pb2.MonitoringDashboardMosaicLayoutTilesWidgetScorecardTimeSeriesQueryTimeSeriesFilterSecondaryAggregation()
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
            dashboard_pb2.MonitoringDashboardMosaicLayoutTilesWidgetScorecardTimeSeriesQueryTimeSeriesFilterPickTimeSeriesFilter()
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
            dashboard_pb2.MonitoringDashboardMosaicLayoutTilesWidgetScorecardTimeSeriesQueryTimeSeriesFilterRatio()
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
            dashboard_pb2.MonitoringDashboardMosaicLayoutTilesWidgetScorecardTimeSeriesQueryTimeSeriesFilterRatioNumerator()
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
            dashboard_pb2.MonitoringDashboardMosaicLayoutTilesWidgetScorecardTimeSeriesQueryTimeSeriesFilterRatioNumeratorAggregation()
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
            dashboard_pb2.MonitoringDashboardMosaicLayoutTilesWidgetScorecardTimeSeriesQueryTimeSeriesFilterRatioDenominator()
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
            dashboard_pb2.MonitoringDashboardMosaicLayoutTilesWidgetScorecardTimeSeriesQueryTimeSeriesFilterRatioDenominatorAggregation()
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
            dashboard_pb2.MonitoringDashboardMosaicLayoutTilesWidgetScorecardTimeSeriesQueryTimeSeriesFilterRatioSecondaryAggregation()
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
            dashboard_pb2.MonitoringDashboardMosaicLayoutTilesWidgetScorecardTimeSeriesQueryTimeSeriesFilterRatioPickTimeSeriesFilter()
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
            dashboard_pb2.MonitoringDashboardMosaicLayoutTilesWidgetScorecardGaugeView()
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
            dashboard_pb2.MonitoringDashboardMosaicLayoutTilesWidgetScorecardSparkChartView()
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
            dashboard_pb2.MonitoringDashboardMosaicLayoutTilesWidgetScorecardThresholds()
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

        res = dashboard_pb2.MonitoringDashboardMosaicLayoutTilesWidgetText()
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

        res = dashboard_pb2.MonitoringDashboardMosaicLayoutTilesWidgetBlank()
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

        res = dashboard_pb2.MonitoringDashboardMosaicLayoutTilesWidgetLogsPanel()
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

        res = dashboard_pb2.MonitoringDashboardRowLayout()
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

        res = dashboard_pb2.MonitoringDashboardRowLayoutRows()
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

        res = dashboard_pb2.MonitoringDashboardRowLayoutRowsWidgets()
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

        res = dashboard_pb2.MonitoringDashboardRowLayoutRowsWidgetsXyChart()
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

        res = dashboard_pb2.MonitoringDashboardRowLayoutRowsWidgetsXyChartDataSets()
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
            dashboard_pb2.MonitoringDashboardRowLayoutRowsWidgetsXyChartDataSetsTimeSeriesQuery()
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
            dashboard_pb2.MonitoringDashboardRowLayoutRowsWidgetsXyChartDataSetsTimeSeriesQueryTimeSeriesFilter()
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
            dashboard_pb2.MonitoringDashboardRowLayoutRowsWidgetsXyChartDataSetsTimeSeriesQueryTimeSeriesFilterAggregation()
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
            dashboard_pb2.MonitoringDashboardRowLayoutRowsWidgetsXyChartDataSetsTimeSeriesQueryTimeSeriesFilterSecondaryAggregation()
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
            dashboard_pb2.MonitoringDashboardRowLayoutRowsWidgetsXyChartDataSetsTimeSeriesQueryTimeSeriesFilterPickTimeSeriesFilter()
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
            dashboard_pb2.MonitoringDashboardRowLayoutRowsWidgetsXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatio()
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
            dashboard_pb2.MonitoringDashboardRowLayoutRowsWidgetsXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioNumerator()
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
            dashboard_pb2.MonitoringDashboardRowLayoutRowsWidgetsXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioNumeratorAggregation()
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
            dashboard_pb2.MonitoringDashboardRowLayoutRowsWidgetsXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioDenominator()
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
            dashboard_pb2.MonitoringDashboardRowLayoutRowsWidgetsXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioDenominatorAggregation()
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
            dashboard_pb2.MonitoringDashboardRowLayoutRowsWidgetsXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioSecondaryAggregation()
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
            dashboard_pb2.MonitoringDashboardRowLayoutRowsWidgetsXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioPickTimeSeriesFilter()
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

        res = dashboard_pb2.MonitoringDashboardRowLayoutRowsWidgetsXyChartThresholds()
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

        res = dashboard_pb2.MonitoringDashboardRowLayoutRowsWidgetsXyChartXAxis()
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

        res = dashboard_pb2.MonitoringDashboardRowLayoutRowsWidgetsXyChartYAxis()
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

        res = dashboard_pb2.MonitoringDashboardRowLayoutRowsWidgetsXyChartChartOptions()
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

        res = dashboard_pb2.MonitoringDashboardRowLayoutRowsWidgetsScorecard()
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
            dashboard_pb2.MonitoringDashboardRowLayoutRowsWidgetsScorecardTimeSeriesQuery()
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
            dashboard_pb2.MonitoringDashboardRowLayoutRowsWidgetsScorecardTimeSeriesQueryTimeSeriesFilter()
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
            dashboard_pb2.MonitoringDashboardRowLayoutRowsWidgetsScorecardTimeSeriesQueryTimeSeriesFilterAggregation()
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
            dashboard_pb2.MonitoringDashboardRowLayoutRowsWidgetsScorecardTimeSeriesQueryTimeSeriesFilterSecondaryAggregation()
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
            dashboard_pb2.MonitoringDashboardRowLayoutRowsWidgetsScorecardTimeSeriesQueryTimeSeriesFilterPickTimeSeriesFilter()
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
            dashboard_pb2.MonitoringDashboardRowLayoutRowsWidgetsScorecardTimeSeriesQueryTimeSeriesFilterRatio()
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
            dashboard_pb2.MonitoringDashboardRowLayoutRowsWidgetsScorecardTimeSeriesQueryTimeSeriesFilterRatioNumerator()
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
            dashboard_pb2.MonitoringDashboardRowLayoutRowsWidgetsScorecardTimeSeriesQueryTimeSeriesFilterRatioNumeratorAggregation()
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
            dashboard_pb2.MonitoringDashboardRowLayoutRowsWidgetsScorecardTimeSeriesQueryTimeSeriesFilterRatioDenominator()
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
            dashboard_pb2.MonitoringDashboardRowLayoutRowsWidgetsScorecardTimeSeriesQueryTimeSeriesFilterRatioDenominatorAggregation()
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
            dashboard_pb2.MonitoringDashboardRowLayoutRowsWidgetsScorecardTimeSeriesQueryTimeSeriesFilterRatioSecondaryAggregation()
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
            dashboard_pb2.MonitoringDashboardRowLayoutRowsWidgetsScorecardTimeSeriesQueryTimeSeriesFilterRatioPickTimeSeriesFilter()
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

        res = dashboard_pb2.MonitoringDashboardRowLayoutRowsWidgetsScorecardGaugeView()
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
            dashboard_pb2.MonitoringDashboardRowLayoutRowsWidgetsScorecardSparkChartView()
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

        res = dashboard_pb2.MonitoringDashboardRowLayoutRowsWidgetsScorecardThresholds()
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

        res = dashboard_pb2.MonitoringDashboardRowLayoutRowsWidgetsText()
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

        res = dashboard_pb2.MonitoringDashboardRowLayoutRowsWidgetsBlank()
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

        res = dashboard_pb2.MonitoringDashboardRowLayoutRowsWidgetsLogsPanel()
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

        res = dashboard_pb2.MonitoringDashboardColumnLayout()
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

        res = dashboard_pb2.MonitoringDashboardColumnLayoutColumns()
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

        res = dashboard_pb2.MonitoringDashboardColumnLayoutColumnsWidgets()
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

        res = dashboard_pb2.MonitoringDashboardColumnLayoutColumnsWidgetsXyChart()
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
            dashboard_pb2.MonitoringDashboardColumnLayoutColumnsWidgetsXyChartDataSets()
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
            dashboard_pb2.MonitoringDashboardColumnLayoutColumnsWidgetsXyChartDataSetsTimeSeriesQuery()
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
            dashboard_pb2.MonitoringDashboardColumnLayoutColumnsWidgetsXyChartDataSetsTimeSeriesQueryTimeSeriesFilter()
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
            dashboard_pb2.MonitoringDashboardColumnLayoutColumnsWidgetsXyChartDataSetsTimeSeriesQueryTimeSeriesFilterAggregation()
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
            dashboard_pb2.MonitoringDashboardColumnLayoutColumnsWidgetsXyChartDataSetsTimeSeriesQueryTimeSeriesFilterSecondaryAggregation()
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
            dashboard_pb2.MonitoringDashboardColumnLayoutColumnsWidgetsXyChartDataSetsTimeSeriesQueryTimeSeriesFilterPickTimeSeriesFilter()
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
            dashboard_pb2.MonitoringDashboardColumnLayoutColumnsWidgetsXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatio()
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
            dashboard_pb2.MonitoringDashboardColumnLayoutColumnsWidgetsXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioNumerator()
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
            dashboard_pb2.MonitoringDashboardColumnLayoutColumnsWidgetsXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioNumeratorAggregation()
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
            dashboard_pb2.MonitoringDashboardColumnLayoutColumnsWidgetsXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioDenominator()
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
            dashboard_pb2.MonitoringDashboardColumnLayoutColumnsWidgetsXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioDenominatorAggregation()
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
            dashboard_pb2.MonitoringDashboardColumnLayoutColumnsWidgetsXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioSecondaryAggregation()
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
            dashboard_pb2.MonitoringDashboardColumnLayoutColumnsWidgetsXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioPickTimeSeriesFilter()
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
            dashboard_pb2.MonitoringDashboardColumnLayoutColumnsWidgetsXyChartThresholds()
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

        res = dashboard_pb2.MonitoringDashboardColumnLayoutColumnsWidgetsXyChartXAxis()
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

        res = dashboard_pb2.MonitoringDashboardColumnLayoutColumnsWidgetsXyChartYAxis()
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
            dashboard_pb2.MonitoringDashboardColumnLayoutColumnsWidgetsXyChartChartOptions()
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

        res = dashboard_pb2.MonitoringDashboardColumnLayoutColumnsWidgetsScorecard()
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
            dashboard_pb2.MonitoringDashboardColumnLayoutColumnsWidgetsScorecardTimeSeriesQuery()
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
            dashboard_pb2.MonitoringDashboardColumnLayoutColumnsWidgetsScorecardTimeSeriesQueryTimeSeriesFilter()
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
            dashboard_pb2.MonitoringDashboardColumnLayoutColumnsWidgetsScorecardTimeSeriesQueryTimeSeriesFilterAggregation()
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
            dashboard_pb2.MonitoringDashboardColumnLayoutColumnsWidgetsScorecardTimeSeriesQueryTimeSeriesFilterSecondaryAggregation()
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
            dashboard_pb2.MonitoringDashboardColumnLayoutColumnsWidgetsScorecardTimeSeriesQueryTimeSeriesFilterPickTimeSeriesFilter()
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
            dashboard_pb2.MonitoringDashboardColumnLayoutColumnsWidgetsScorecardTimeSeriesQueryTimeSeriesFilterRatio()
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
            dashboard_pb2.MonitoringDashboardColumnLayoutColumnsWidgetsScorecardTimeSeriesQueryTimeSeriesFilterRatioNumerator()
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
            dashboard_pb2.MonitoringDashboardColumnLayoutColumnsWidgetsScorecardTimeSeriesQueryTimeSeriesFilterRatioNumeratorAggregation()
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
            dashboard_pb2.MonitoringDashboardColumnLayoutColumnsWidgetsScorecardTimeSeriesQueryTimeSeriesFilterRatioDenominator()
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
            dashboard_pb2.MonitoringDashboardColumnLayoutColumnsWidgetsScorecardTimeSeriesQueryTimeSeriesFilterRatioDenominatorAggregation()
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
            dashboard_pb2.MonitoringDashboardColumnLayoutColumnsWidgetsScorecardTimeSeriesQueryTimeSeriesFilterRatioSecondaryAggregation()
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
            dashboard_pb2.MonitoringDashboardColumnLayoutColumnsWidgetsScorecardTimeSeriesQueryTimeSeriesFilterRatioPickTimeSeriesFilter()
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
            dashboard_pb2.MonitoringDashboardColumnLayoutColumnsWidgetsScorecardGaugeView()
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
            dashboard_pb2.MonitoringDashboardColumnLayoutColumnsWidgetsScorecardSparkChartView()
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
            dashboard_pb2.MonitoringDashboardColumnLayoutColumnsWidgetsScorecardThresholds()
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

        res = dashboard_pb2.MonitoringDashboardColumnLayoutColumnsWidgetsText()
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

        res = dashboard_pb2.MonitoringDashboardColumnLayoutColumnsWidgetsBlank()
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

        res = dashboard_pb2.MonitoringDashboardColumnLayoutColumnsWidgetsLogsPanel()
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
        return dashboard_pb2.MonitoringDashboardGridLayoutWidgetsXyChartDataSetsTimeSeriesQueryTimeSeriesFilterAggregationPerSeriesAlignerEnum.Value(
            "MonitoringDashboardGridLayoutWidgetsXyChartDataSetsTimeSeriesQueryTimeSeriesFilterAggregationPerSeriesAlignerEnum%s"
            % resource
        )

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return resource
        return dashboard_pb2.MonitoringDashboardGridLayoutWidgetsXyChartDataSetsTimeSeriesQueryTimeSeriesFilterAggregationPerSeriesAlignerEnum.Name(
            resource
        )[
            len(
                "MonitoringDashboardGridLayoutWidgetsXyChartDataSetsTimeSeriesQueryTimeSeriesFilterAggregationPerSeriesAlignerEnum"
            ) :
        ]


class DashboardGridLayoutWidgetsXyChartDataSetsTimeSeriesQueryTimeSeriesFilterAggregationCrossSeriesReducerEnum(
    object
):
    @classmethod
    def to_proto(self, resource):
        if not resource:
            return resource
        return dashboard_pb2.MonitoringDashboardGridLayoutWidgetsXyChartDataSetsTimeSeriesQueryTimeSeriesFilterAggregationCrossSeriesReducerEnum.Value(
            "MonitoringDashboardGridLayoutWidgetsXyChartDataSetsTimeSeriesQueryTimeSeriesFilterAggregationCrossSeriesReducerEnum%s"
            % resource
        )

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return resource
        return dashboard_pb2.MonitoringDashboardGridLayoutWidgetsXyChartDataSetsTimeSeriesQueryTimeSeriesFilterAggregationCrossSeriesReducerEnum.Name(
            resource
        )[
            len(
                "MonitoringDashboardGridLayoutWidgetsXyChartDataSetsTimeSeriesQueryTimeSeriesFilterAggregationCrossSeriesReducerEnum"
            ) :
        ]


class DashboardGridLayoutWidgetsXyChartDataSetsTimeSeriesQueryTimeSeriesFilterSecondaryAggregationPerSeriesAlignerEnum(
    object
):
    @classmethod
    def to_proto(self, resource):
        if not resource:
            return resource
        return dashboard_pb2.MonitoringDashboardGridLayoutWidgetsXyChartDataSetsTimeSeriesQueryTimeSeriesFilterSecondaryAggregationPerSeriesAlignerEnum.Value(
            "MonitoringDashboardGridLayoutWidgetsXyChartDataSetsTimeSeriesQueryTimeSeriesFilterSecondaryAggregationPerSeriesAlignerEnum%s"
            % resource
        )

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return resource
        return dashboard_pb2.MonitoringDashboardGridLayoutWidgetsXyChartDataSetsTimeSeriesQueryTimeSeriesFilterSecondaryAggregationPerSeriesAlignerEnum.Name(
            resource
        )[
            len(
                "MonitoringDashboardGridLayoutWidgetsXyChartDataSetsTimeSeriesQueryTimeSeriesFilterSecondaryAggregationPerSeriesAlignerEnum"
            ) :
        ]


class DashboardGridLayoutWidgetsXyChartDataSetsTimeSeriesQueryTimeSeriesFilterSecondaryAggregationCrossSeriesReducerEnum(
    object
):
    @classmethod
    def to_proto(self, resource):
        if not resource:
            return resource
        return dashboard_pb2.MonitoringDashboardGridLayoutWidgetsXyChartDataSetsTimeSeriesQueryTimeSeriesFilterSecondaryAggregationCrossSeriesReducerEnum.Value(
            "MonitoringDashboardGridLayoutWidgetsXyChartDataSetsTimeSeriesQueryTimeSeriesFilterSecondaryAggregationCrossSeriesReducerEnum%s"
            % resource
        )

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return resource
        return dashboard_pb2.MonitoringDashboardGridLayoutWidgetsXyChartDataSetsTimeSeriesQueryTimeSeriesFilterSecondaryAggregationCrossSeriesReducerEnum.Name(
            resource
        )[
            len(
                "MonitoringDashboardGridLayoutWidgetsXyChartDataSetsTimeSeriesQueryTimeSeriesFilterSecondaryAggregationCrossSeriesReducerEnum"
            ) :
        ]


class DashboardGridLayoutWidgetsXyChartDataSetsTimeSeriesQueryTimeSeriesFilterPickTimeSeriesFilterRankingMethodEnum(
    object
):
    @classmethod
    def to_proto(self, resource):
        if not resource:
            return resource
        return dashboard_pb2.MonitoringDashboardGridLayoutWidgetsXyChartDataSetsTimeSeriesQueryTimeSeriesFilterPickTimeSeriesFilterRankingMethodEnum.Value(
            "MonitoringDashboardGridLayoutWidgetsXyChartDataSetsTimeSeriesQueryTimeSeriesFilterPickTimeSeriesFilterRankingMethodEnum%s"
            % resource
        )

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return resource
        return dashboard_pb2.MonitoringDashboardGridLayoutWidgetsXyChartDataSetsTimeSeriesQueryTimeSeriesFilterPickTimeSeriesFilterRankingMethodEnum.Name(
            resource
        )[
            len(
                "MonitoringDashboardGridLayoutWidgetsXyChartDataSetsTimeSeriesQueryTimeSeriesFilterPickTimeSeriesFilterRankingMethodEnum"
            ) :
        ]


class DashboardGridLayoutWidgetsXyChartDataSetsTimeSeriesQueryTimeSeriesFilterPickTimeSeriesFilterDirectionEnum(
    object
):
    @classmethod
    def to_proto(self, resource):
        if not resource:
            return resource
        return dashboard_pb2.MonitoringDashboardGridLayoutWidgetsXyChartDataSetsTimeSeriesQueryTimeSeriesFilterPickTimeSeriesFilterDirectionEnum.Value(
            "MonitoringDashboardGridLayoutWidgetsXyChartDataSetsTimeSeriesQueryTimeSeriesFilterPickTimeSeriesFilterDirectionEnum%s"
            % resource
        )

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return resource
        return dashboard_pb2.MonitoringDashboardGridLayoutWidgetsXyChartDataSetsTimeSeriesQueryTimeSeriesFilterPickTimeSeriesFilterDirectionEnum.Name(
            resource
        )[
            len(
                "MonitoringDashboardGridLayoutWidgetsXyChartDataSetsTimeSeriesQueryTimeSeriesFilterPickTimeSeriesFilterDirectionEnum"
            ) :
        ]


class DashboardGridLayoutWidgetsXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioNumeratorAggregationPerSeriesAlignerEnum(
    object
):
    @classmethod
    def to_proto(self, resource):
        if not resource:
            return resource
        return dashboard_pb2.MonitoringDashboardGridLayoutWidgetsXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioNumeratorAggregationPerSeriesAlignerEnum.Value(
            "MonitoringDashboardGridLayoutWidgetsXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioNumeratorAggregationPerSeriesAlignerEnum%s"
            % resource
        )

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return resource
        return dashboard_pb2.MonitoringDashboardGridLayoutWidgetsXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioNumeratorAggregationPerSeriesAlignerEnum.Name(
            resource
        )[
            len(
                "MonitoringDashboardGridLayoutWidgetsXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioNumeratorAggregationPerSeriesAlignerEnum"
            ) :
        ]


class DashboardGridLayoutWidgetsXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioNumeratorAggregationCrossSeriesReducerEnum(
    object
):
    @classmethod
    def to_proto(self, resource):
        if not resource:
            return resource
        return dashboard_pb2.MonitoringDashboardGridLayoutWidgetsXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioNumeratorAggregationCrossSeriesReducerEnum.Value(
            "MonitoringDashboardGridLayoutWidgetsXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioNumeratorAggregationCrossSeriesReducerEnum%s"
            % resource
        )

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return resource
        return dashboard_pb2.MonitoringDashboardGridLayoutWidgetsXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioNumeratorAggregationCrossSeriesReducerEnum.Name(
            resource
        )[
            len(
                "MonitoringDashboardGridLayoutWidgetsXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioNumeratorAggregationCrossSeriesReducerEnum"
            ) :
        ]


class DashboardGridLayoutWidgetsXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioDenominatorAggregationPerSeriesAlignerEnum(
    object
):
    @classmethod
    def to_proto(self, resource):
        if not resource:
            return resource
        return dashboard_pb2.MonitoringDashboardGridLayoutWidgetsXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioDenominatorAggregationPerSeriesAlignerEnum.Value(
            "MonitoringDashboardGridLayoutWidgetsXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioDenominatorAggregationPerSeriesAlignerEnum%s"
            % resource
        )

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return resource
        return dashboard_pb2.MonitoringDashboardGridLayoutWidgetsXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioDenominatorAggregationPerSeriesAlignerEnum.Name(
            resource
        )[
            len(
                "MonitoringDashboardGridLayoutWidgetsXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioDenominatorAggregationPerSeriesAlignerEnum"
            ) :
        ]


class DashboardGridLayoutWidgetsXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioDenominatorAggregationCrossSeriesReducerEnum(
    object
):
    @classmethod
    def to_proto(self, resource):
        if not resource:
            return resource
        return dashboard_pb2.MonitoringDashboardGridLayoutWidgetsXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioDenominatorAggregationCrossSeriesReducerEnum.Value(
            "MonitoringDashboardGridLayoutWidgetsXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioDenominatorAggregationCrossSeriesReducerEnum%s"
            % resource
        )

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return resource
        return dashboard_pb2.MonitoringDashboardGridLayoutWidgetsXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioDenominatorAggregationCrossSeriesReducerEnum.Name(
            resource
        )[
            len(
                "MonitoringDashboardGridLayoutWidgetsXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioDenominatorAggregationCrossSeriesReducerEnum"
            ) :
        ]


class DashboardGridLayoutWidgetsXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioSecondaryAggregationPerSeriesAlignerEnum(
    object
):
    @classmethod
    def to_proto(self, resource):
        if not resource:
            return resource
        return dashboard_pb2.MonitoringDashboardGridLayoutWidgetsXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioSecondaryAggregationPerSeriesAlignerEnum.Value(
            "MonitoringDashboardGridLayoutWidgetsXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioSecondaryAggregationPerSeriesAlignerEnum%s"
            % resource
        )

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return resource
        return dashboard_pb2.MonitoringDashboardGridLayoutWidgetsXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioSecondaryAggregationPerSeriesAlignerEnum.Name(
            resource
        )[
            len(
                "MonitoringDashboardGridLayoutWidgetsXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioSecondaryAggregationPerSeriesAlignerEnum"
            ) :
        ]


class DashboardGridLayoutWidgetsXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioSecondaryAggregationCrossSeriesReducerEnum(
    object
):
    @classmethod
    def to_proto(self, resource):
        if not resource:
            return resource
        return dashboard_pb2.MonitoringDashboardGridLayoutWidgetsXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioSecondaryAggregationCrossSeriesReducerEnum.Value(
            "MonitoringDashboardGridLayoutWidgetsXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioSecondaryAggregationCrossSeriesReducerEnum%s"
            % resource
        )

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return resource
        return dashboard_pb2.MonitoringDashboardGridLayoutWidgetsXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioSecondaryAggregationCrossSeriesReducerEnum.Name(
            resource
        )[
            len(
                "MonitoringDashboardGridLayoutWidgetsXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioSecondaryAggregationCrossSeriesReducerEnum"
            ) :
        ]


class DashboardGridLayoutWidgetsXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioPickTimeSeriesFilterRankingMethodEnum(
    object
):
    @classmethod
    def to_proto(self, resource):
        if not resource:
            return resource
        return dashboard_pb2.MonitoringDashboardGridLayoutWidgetsXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioPickTimeSeriesFilterRankingMethodEnum.Value(
            "MonitoringDashboardGridLayoutWidgetsXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioPickTimeSeriesFilterRankingMethodEnum%s"
            % resource
        )

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return resource
        return dashboard_pb2.MonitoringDashboardGridLayoutWidgetsXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioPickTimeSeriesFilterRankingMethodEnum.Name(
            resource
        )[
            len(
                "MonitoringDashboardGridLayoutWidgetsXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioPickTimeSeriesFilterRankingMethodEnum"
            ) :
        ]


class DashboardGridLayoutWidgetsXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioPickTimeSeriesFilterDirectionEnum(
    object
):
    @classmethod
    def to_proto(self, resource):
        if not resource:
            return resource
        return dashboard_pb2.MonitoringDashboardGridLayoutWidgetsXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioPickTimeSeriesFilterDirectionEnum.Value(
            "MonitoringDashboardGridLayoutWidgetsXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioPickTimeSeriesFilterDirectionEnum%s"
            % resource
        )

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return resource
        return dashboard_pb2.MonitoringDashboardGridLayoutWidgetsXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioPickTimeSeriesFilterDirectionEnum.Name(
            resource
        )[
            len(
                "MonitoringDashboardGridLayoutWidgetsXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioPickTimeSeriesFilterDirectionEnum"
            ) :
        ]


class DashboardGridLayoutWidgetsXyChartDataSetsPlotTypeEnum(object):
    @classmethod
    def to_proto(self, resource):
        if not resource:
            return resource
        return dashboard_pb2.MonitoringDashboardGridLayoutWidgetsXyChartDataSetsPlotTypeEnum.Value(
            "MonitoringDashboardGridLayoutWidgetsXyChartDataSetsPlotTypeEnum%s"
            % resource
        )

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return resource
        return dashboard_pb2.MonitoringDashboardGridLayoutWidgetsXyChartDataSetsPlotTypeEnum.Name(
            resource
        )[
            len("MonitoringDashboardGridLayoutWidgetsXyChartDataSetsPlotTypeEnum") :
        ]


class DashboardGridLayoutWidgetsXyChartThresholdsColorEnum(object):
    @classmethod
    def to_proto(self, resource):
        if not resource:
            return resource
        return dashboard_pb2.MonitoringDashboardGridLayoutWidgetsXyChartThresholdsColorEnum.Value(
            "MonitoringDashboardGridLayoutWidgetsXyChartThresholdsColorEnum%s"
            % resource
        )

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return resource
        return dashboard_pb2.MonitoringDashboardGridLayoutWidgetsXyChartThresholdsColorEnum.Name(
            resource
        )[
            len("MonitoringDashboardGridLayoutWidgetsXyChartThresholdsColorEnum") :
        ]


class DashboardGridLayoutWidgetsXyChartThresholdsDirectionEnum(object):
    @classmethod
    def to_proto(self, resource):
        if not resource:
            return resource
        return dashboard_pb2.MonitoringDashboardGridLayoutWidgetsXyChartThresholdsDirectionEnum.Value(
            "MonitoringDashboardGridLayoutWidgetsXyChartThresholdsDirectionEnum%s"
            % resource
        )

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return resource
        return dashboard_pb2.MonitoringDashboardGridLayoutWidgetsXyChartThresholdsDirectionEnum.Name(
            resource
        )[
            len("MonitoringDashboardGridLayoutWidgetsXyChartThresholdsDirectionEnum") :
        ]


class DashboardGridLayoutWidgetsXyChartXAxisScaleEnum(object):
    @classmethod
    def to_proto(self, resource):
        if not resource:
            return resource
        return dashboard_pb2.MonitoringDashboardGridLayoutWidgetsXyChartXAxisScaleEnum.Value(
            "MonitoringDashboardGridLayoutWidgetsXyChartXAxisScaleEnum%s" % resource
        )

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return resource
        return dashboard_pb2.MonitoringDashboardGridLayoutWidgetsXyChartXAxisScaleEnum.Name(
            resource
        )[
            len("MonitoringDashboardGridLayoutWidgetsXyChartXAxisScaleEnum") :
        ]


class DashboardGridLayoutWidgetsXyChartYAxisScaleEnum(object):
    @classmethod
    def to_proto(self, resource):
        if not resource:
            return resource
        return dashboard_pb2.MonitoringDashboardGridLayoutWidgetsXyChartYAxisScaleEnum.Value(
            "MonitoringDashboardGridLayoutWidgetsXyChartYAxisScaleEnum%s" % resource
        )

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return resource
        return dashboard_pb2.MonitoringDashboardGridLayoutWidgetsXyChartYAxisScaleEnum.Name(
            resource
        )[
            len("MonitoringDashboardGridLayoutWidgetsXyChartYAxisScaleEnum") :
        ]


class DashboardGridLayoutWidgetsXyChartChartOptionsModeEnum(object):
    @classmethod
    def to_proto(self, resource):
        if not resource:
            return resource
        return dashboard_pb2.MonitoringDashboardGridLayoutWidgetsXyChartChartOptionsModeEnum.Value(
            "MonitoringDashboardGridLayoutWidgetsXyChartChartOptionsModeEnum%s"
            % resource
        )

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return resource
        return dashboard_pb2.MonitoringDashboardGridLayoutWidgetsXyChartChartOptionsModeEnum.Name(
            resource
        )[
            len("MonitoringDashboardGridLayoutWidgetsXyChartChartOptionsModeEnum") :
        ]


class DashboardGridLayoutWidgetsScorecardTimeSeriesQueryTimeSeriesFilterAggregationPerSeriesAlignerEnum(
    object
):
    @classmethod
    def to_proto(self, resource):
        if not resource:
            return resource
        return dashboard_pb2.MonitoringDashboardGridLayoutWidgetsScorecardTimeSeriesQueryTimeSeriesFilterAggregationPerSeriesAlignerEnum.Value(
            "MonitoringDashboardGridLayoutWidgetsScorecardTimeSeriesQueryTimeSeriesFilterAggregationPerSeriesAlignerEnum%s"
            % resource
        )

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return resource
        return dashboard_pb2.MonitoringDashboardGridLayoutWidgetsScorecardTimeSeriesQueryTimeSeriesFilterAggregationPerSeriesAlignerEnum.Name(
            resource
        )[
            len(
                "MonitoringDashboardGridLayoutWidgetsScorecardTimeSeriesQueryTimeSeriesFilterAggregationPerSeriesAlignerEnum"
            ) :
        ]


class DashboardGridLayoutWidgetsScorecardTimeSeriesQueryTimeSeriesFilterAggregationCrossSeriesReducerEnum(
    object
):
    @classmethod
    def to_proto(self, resource):
        if not resource:
            return resource
        return dashboard_pb2.MonitoringDashboardGridLayoutWidgetsScorecardTimeSeriesQueryTimeSeriesFilterAggregationCrossSeriesReducerEnum.Value(
            "MonitoringDashboardGridLayoutWidgetsScorecardTimeSeriesQueryTimeSeriesFilterAggregationCrossSeriesReducerEnum%s"
            % resource
        )

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return resource
        return dashboard_pb2.MonitoringDashboardGridLayoutWidgetsScorecardTimeSeriesQueryTimeSeriesFilterAggregationCrossSeriesReducerEnum.Name(
            resource
        )[
            len(
                "MonitoringDashboardGridLayoutWidgetsScorecardTimeSeriesQueryTimeSeriesFilterAggregationCrossSeriesReducerEnum"
            ) :
        ]


class DashboardGridLayoutWidgetsScorecardTimeSeriesQueryTimeSeriesFilterSecondaryAggregationPerSeriesAlignerEnum(
    object
):
    @classmethod
    def to_proto(self, resource):
        if not resource:
            return resource
        return dashboard_pb2.MonitoringDashboardGridLayoutWidgetsScorecardTimeSeriesQueryTimeSeriesFilterSecondaryAggregationPerSeriesAlignerEnum.Value(
            "MonitoringDashboardGridLayoutWidgetsScorecardTimeSeriesQueryTimeSeriesFilterSecondaryAggregationPerSeriesAlignerEnum%s"
            % resource
        )

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return resource
        return dashboard_pb2.MonitoringDashboardGridLayoutWidgetsScorecardTimeSeriesQueryTimeSeriesFilterSecondaryAggregationPerSeriesAlignerEnum.Name(
            resource
        )[
            len(
                "MonitoringDashboardGridLayoutWidgetsScorecardTimeSeriesQueryTimeSeriesFilterSecondaryAggregationPerSeriesAlignerEnum"
            ) :
        ]


class DashboardGridLayoutWidgetsScorecardTimeSeriesQueryTimeSeriesFilterSecondaryAggregationCrossSeriesReducerEnum(
    object
):
    @classmethod
    def to_proto(self, resource):
        if not resource:
            return resource
        return dashboard_pb2.MonitoringDashboardGridLayoutWidgetsScorecardTimeSeriesQueryTimeSeriesFilterSecondaryAggregationCrossSeriesReducerEnum.Value(
            "MonitoringDashboardGridLayoutWidgetsScorecardTimeSeriesQueryTimeSeriesFilterSecondaryAggregationCrossSeriesReducerEnum%s"
            % resource
        )

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return resource
        return dashboard_pb2.MonitoringDashboardGridLayoutWidgetsScorecardTimeSeriesQueryTimeSeriesFilterSecondaryAggregationCrossSeriesReducerEnum.Name(
            resource
        )[
            len(
                "MonitoringDashboardGridLayoutWidgetsScorecardTimeSeriesQueryTimeSeriesFilterSecondaryAggregationCrossSeriesReducerEnum"
            ) :
        ]


class DashboardGridLayoutWidgetsScorecardTimeSeriesQueryTimeSeriesFilterPickTimeSeriesFilterRankingMethodEnum(
    object
):
    @classmethod
    def to_proto(self, resource):
        if not resource:
            return resource
        return dashboard_pb2.MonitoringDashboardGridLayoutWidgetsScorecardTimeSeriesQueryTimeSeriesFilterPickTimeSeriesFilterRankingMethodEnum.Value(
            "MonitoringDashboardGridLayoutWidgetsScorecardTimeSeriesQueryTimeSeriesFilterPickTimeSeriesFilterRankingMethodEnum%s"
            % resource
        )

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return resource
        return dashboard_pb2.MonitoringDashboardGridLayoutWidgetsScorecardTimeSeriesQueryTimeSeriesFilterPickTimeSeriesFilterRankingMethodEnum.Name(
            resource
        )[
            len(
                "MonitoringDashboardGridLayoutWidgetsScorecardTimeSeriesQueryTimeSeriesFilterPickTimeSeriesFilterRankingMethodEnum"
            ) :
        ]


class DashboardGridLayoutWidgetsScorecardTimeSeriesQueryTimeSeriesFilterPickTimeSeriesFilterDirectionEnum(
    object
):
    @classmethod
    def to_proto(self, resource):
        if not resource:
            return resource
        return dashboard_pb2.MonitoringDashboardGridLayoutWidgetsScorecardTimeSeriesQueryTimeSeriesFilterPickTimeSeriesFilterDirectionEnum.Value(
            "MonitoringDashboardGridLayoutWidgetsScorecardTimeSeriesQueryTimeSeriesFilterPickTimeSeriesFilterDirectionEnum%s"
            % resource
        )

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return resource
        return dashboard_pb2.MonitoringDashboardGridLayoutWidgetsScorecardTimeSeriesQueryTimeSeriesFilterPickTimeSeriesFilterDirectionEnum.Name(
            resource
        )[
            len(
                "MonitoringDashboardGridLayoutWidgetsScorecardTimeSeriesQueryTimeSeriesFilterPickTimeSeriesFilterDirectionEnum"
            ) :
        ]


class DashboardGridLayoutWidgetsScorecardTimeSeriesQueryTimeSeriesFilterRatioNumeratorAggregationPerSeriesAlignerEnum(
    object
):
    @classmethod
    def to_proto(self, resource):
        if not resource:
            return resource
        return dashboard_pb2.MonitoringDashboardGridLayoutWidgetsScorecardTimeSeriesQueryTimeSeriesFilterRatioNumeratorAggregationPerSeriesAlignerEnum.Value(
            "MonitoringDashboardGridLayoutWidgetsScorecardTimeSeriesQueryTimeSeriesFilterRatioNumeratorAggregationPerSeriesAlignerEnum%s"
            % resource
        )

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return resource
        return dashboard_pb2.MonitoringDashboardGridLayoutWidgetsScorecardTimeSeriesQueryTimeSeriesFilterRatioNumeratorAggregationPerSeriesAlignerEnum.Name(
            resource
        )[
            len(
                "MonitoringDashboardGridLayoutWidgetsScorecardTimeSeriesQueryTimeSeriesFilterRatioNumeratorAggregationPerSeriesAlignerEnum"
            ) :
        ]


class DashboardGridLayoutWidgetsScorecardTimeSeriesQueryTimeSeriesFilterRatioNumeratorAggregationCrossSeriesReducerEnum(
    object
):
    @classmethod
    def to_proto(self, resource):
        if not resource:
            return resource
        return dashboard_pb2.MonitoringDashboardGridLayoutWidgetsScorecardTimeSeriesQueryTimeSeriesFilterRatioNumeratorAggregationCrossSeriesReducerEnum.Value(
            "MonitoringDashboardGridLayoutWidgetsScorecardTimeSeriesQueryTimeSeriesFilterRatioNumeratorAggregationCrossSeriesReducerEnum%s"
            % resource
        )

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return resource
        return dashboard_pb2.MonitoringDashboardGridLayoutWidgetsScorecardTimeSeriesQueryTimeSeriesFilterRatioNumeratorAggregationCrossSeriesReducerEnum.Name(
            resource
        )[
            len(
                "MonitoringDashboardGridLayoutWidgetsScorecardTimeSeriesQueryTimeSeriesFilterRatioNumeratorAggregationCrossSeriesReducerEnum"
            ) :
        ]


class DashboardGridLayoutWidgetsScorecardTimeSeriesQueryTimeSeriesFilterRatioDenominatorAggregationPerSeriesAlignerEnum(
    object
):
    @classmethod
    def to_proto(self, resource):
        if not resource:
            return resource
        return dashboard_pb2.MonitoringDashboardGridLayoutWidgetsScorecardTimeSeriesQueryTimeSeriesFilterRatioDenominatorAggregationPerSeriesAlignerEnum.Value(
            "MonitoringDashboardGridLayoutWidgetsScorecardTimeSeriesQueryTimeSeriesFilterRatioDenominatorAggregationPerSeriesAlignerEnum%s"
            % resource
        )

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return resource
        return dashboard_pb2.MonitoringDashboardGridLayoutWidgetsScorecardTimeSeriesQueryTimeSeriesFilterRatioDenominatorAggregationPerSeriesAlignerEnum.Name(
            resource
        )[
            len(
                "MonitoringDashboardGridLayoutWidgetsScorecardTimeSeriesQueryTimeSeriesFilterRatioDenominatorAggregationPerSeriesAlignerEnum"
            ) :
        ]


class DashboardGridLayoutWidgetsScorecardTimeSeriesQueryTimeSeriesFilterRatioDenominatorAggregationCrossSeriesReducerEnum(
    object
):
    @classmethod
    def to_proto(self, resource):
        if not resource:
            return resource
        return dashboard_pb2.MonitoringDashboardGridLayoutWidgetsScorecardTimeSeriesQueryTimeSeriesFilterRatioDenominatorAggregationCrossSeriesReducerEnum.Value(
            "MonitoringDashboardGridLayoutWidgetsScorecardTimeSeriesQueryTimeSeriesFilterRatioDenominatorAggregationCrossSeriesReducerEnum%s"
            % resource
        )

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return resource
        return dashboard_pb2.MonitoringDashboardGridLayoutWidgetsScorecardTimeSeriesQueryTimeSeriesFilterRatioDenominatorAggregationCrossSeriesReducerEnum.Name(
            resource
        )[
            len(
                "MonitoringDashboardGridLayoutWidgetsScorecardTimeSeriesQueryTimeSeriesFilterRatioDenominatorAggregationCrossSeriesReducerEnum"
            ) :
        ]


class DashboardGridLayoutWidgetsScorecardTimeSeriesQueryTimeSeriesFilterRatioSecondaryAggregationPerSeriesAlignerEnum(
    object
):
    @classmethod
    def to_proto(self, resource):
        if not resource:
            return resource
        return dashboard_pb2.MonitoringDashboardGridLayoutWidgetsScorecardTimeSeriesQueryTimeSeriesFilterRatioSecondaryAggregationPerSeriesAlignerEnum.Value(
            "MonitoringDashboardGridLayoutWidgetsScorecardTimeSeriesQueryTimeSeriesFilterRatioSecondaryAggregationPerSeriesAlignerEnum%s"
            % resource
        )

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return resource
        return dashboard_pb2.MonitoringDashboardGridLayoutWidgetsScorecardTimeSeriesQueryTimeSeriesFilterRatioSecondaryAggregationPerSeriesAlignerEnum.Name(
            resource
        )[
            len(
                "MonitoringDashboardGridLayoutWidgetsScorecardTimeSeriesQueryTimeSeriesFilterRatioSecondaryAggregationPerSeriesAlignerEnum"
            ) :
        ]


class DashboardGridLayoutWidgetsScorecardTimeSeriesQueryTimeSeriesFilterRatioSecondaryAggregationCrossSeriesReducerEnum(
    object
):
    @classmethod
    def to_proto(self, resource):
        if not resource:
            return resource
        return dashboard_pb2.MonitoringDashboardGridLayoutWidgetsScorecardTimeSeriesQueryTimeSeriesFilterRatioSecondaryAggregationCrossSeriesReducerEnum.Value(
            "MonitoringDashboardGridLayoutWidgetsScorecardTimeSeriesQueryTimeSeriesFilterRatioSecondaryAggregationCrossSeriesReducerEnum%s"
            % resource
        )

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return resource
        return dashboard_pb2.MonitoringDashboardGridLayoutWidgetsScorecardTimeSeriesQueryTimeSeriesFilterRatioSecondaryAggregationCrossSeriesReducerEnum.Name(
            resource
        )[
            len(
                "MonitoringDashboardGridLayoutWidgetsScorecardTimeSeriesQueryTimeSeriesFilterRatioSecondaryAggregationCrossSeriesReducerEnum"
            ) :
        ]


class DashboardGridLayoutWidgetsScorecardTimeSeriesQueryTimeSeriesFilterRatioPickTimeSeriesFilterRankingMethodEnum(
    object
):
    @classmethod
    def to_proto(self, resource):
        if not resource:
            return resource
        return dashboard_pb2.MonitoringDashboardGridLayoutWidgetsScorecardTimeSeriesQueryTimeSeriesFilterRatioPickTimeSeriesFilterRankingMethodEnum.Value(
            "MonitoringDashboardGridLayoutWidgetsScorecardTimeSeriesQueryTimeSeriesFilterRatioPickTimeSeriesFilterRankingMethodEnum%s"
            % resource
        )

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return resource
        return dashboard_pb2.MonitoringDashboardGridLayoutWidgetsScorecardTimeSeriesQueryTimeSeriesFilterRatioPickTimeSeriesFilterRankingMethodEnum.Name(
            resource
        )[
            len(
                "MonitoringDashboardGridLayoutWidgetsScorecardTimeSeriesQueryTimeSeriesFilterRatioPickTimeSeriesFilterRankingMethodEnum"
            ) :
        ]


class DashboardGridLayoutWidgetsScorecardTimeSeriesQueryTimeSeriesFilterRatioPickTimeSeriesFilterDirectionEnum(
    object
):
    @classmethod
    def to_proto(self, resource):
        if not resource:
            return resource
        return dashboard_pb2.MonitoringDashboardGridLayoutWidgetsScorecardTimeSeriesQueryTimeSeriesFilterRatioPickTimeSeriesFilterDirectionEnum.Value(
            "MonitoringDashboardGridLayoutWidgetsScorecardTimeSeriesQueryTimeSeriesFilterRatioPickTimeSeriesFilterDirectionEnum%s"
            % resource
        )

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return resource
        return dashboard_pb2.MonitoringDashboardGridLayoutWidgetsScorecardTimeSeriesQueryTimeSeriesFilterRatioPickTimeSeriesFilterDirectionEnum.Name(
            resource
        )[
            len(
                "MonitoringDashboardGridLayoutWidgetsScorecardTimeSeriesQueryTimeSeriesFilterRatioPickTimeSeriesFilterDirectionEnum"
            ) :
        ]


class DashboardGridLayoutWidgetsScorecardSparkChartViewSparkChartTypeEnum(object):
    @classmethod
    def to_proto(self, resource):
        if not resource:
            return resource
        return dashboard_pb2.MonitoringDashboardGridLayoutWidgetsScorecardSparkChartViewSparkChartTypeEnum.Value(
            "MonitoringDashboardGridLayoutWidgetsScorecardSparkChartViewSparkChartTypeEnum%s"
            % resource
        )

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return resource
        return dashboard_pb2.MonitoringDashboardGridLayoutWidgetsScorecardSparkChartViewSparkChartTypeEnum.Name(
            resource
        )[
            len(
                "MonitoringDashboardGridLayoutWidgetsScorecardSparkChartViewSparkChartTypeEnum"
            ) :
        ]


class DashboardGridLayoutWidgetsScorecardThresholdsColorEnum(object):
    @classmethod
    def to_proto(self, resource):
        if not resource:
            return resource
        return dashboard_pb2.MonitoringDashboardGridLayoutWidgetsScorecardThresholdsColorEnum.Value(
            "MonitoringDashboardGridLayoutWidgetsScorecardThresholdsColorEnum%s"
            % resource
        )

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return resource
        return dashboard_pb2.MonitoringDashboardGridLayoutWidgetsScorecardThresholdsColorEnum.Name(
            resource
        )[
            len("MonitoringDashboardGridLayoutWidgetsScorecardThresholdsColorEnum") :
        ]


class DashboardGridLayoutWidgetsScorecardThresholdsDirectionEnum(object):
    @classmethod
    def to_proto(self, resource):
        if not resource:
            return resource
        return dashboard_pb2.MonitoringDashboardGridLayoutWidgetsScorecardThresholdsDirectionEnum.Value(
            "MonitoringDashboardGridLayoutWidgetsScorecardThresholdsDirectionEnum%s"
            % resource
        )

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return resource
        return dashboard_pb2.MonitoringDashboardGridLayoutWidgetsScorecardThresholdsDirectionEnum.Name(
            resource
        )[
            len(
                "MonitoringDashboardGridLayoutWidgetsScorecardThresholdsDirectionEnum"
            ) :
        ]


class DashboardGridLayoutWidgetsTextFormatEnum(object):
    @classmethod
    def to_proto(self, resource):
        if not resource:
            return resource
        return dashboard_pb2.MonitoringDashboardGridLayoutWidgetsTextFormatEnum.Value(
            "MonitoringDashboardGridLayoutWidgetsTextFormatEnum%s" % resource
        )

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return resource
        return dashboard_pb2.MonitoringDashboardGridLayoutWidgetsTextFormatEnum.Name(
            resource
        )[len("MonitoringDashboardGridLayoutWidgetsTextFormatEnum") :]


class DashboardMosaicLayoutTilesWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterAggregationPerSeriesAlignerEnum(
    object
):
    @classmethod
    def to_proto(self, resource):
        if not resource:
            return resource
        return dashboard_pb2.MonitoringDashboardMosaicLayoutTilesWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterAggregationPerSeriesAlignerEnum.Value(
            "MonitoringDashboardMosaicLayoutTilesWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterAggregationPerSeriesAlignerEnum%s"
            % resource
        )

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return resource
        return dashboard_pb2.MonitoringDashboardMosaicLayoutTilesWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterAggregationPerSeriesAlignerEnum.Name(
            resource
        )[
            len(
                "MonitoringDashboardMosaicLayoutTilesWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterAggregationPerSeriesAlignerEnum"
            ) :
        ]


class DashboardMosaicLayoutTilesWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterAggregationCrossSeriesReducerEnum(
    object
):
    @classmethod
    def to_proto(self, resource):
        if not resource:
            return resource
        return dashboard_pb2.MonitoringDashboardMosaicLayoutTilesWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterAggregationCrossSeriesReducerEnum.Value(
            "MonitoringDashboardMosaicLayoutTilesWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterAggregationCrossSeriesReducerEnum%s"
            % resource
        )

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return resource
        return dashboard_pb2.MonitoringDashboardMosaicLayoutTilesWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterAggregationCrossSeriesReducerEnum.Name(
            resource
        )[
            len(
                "MonitoringDashboardMosaicLayoutTilesWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterAggregationCrossSeriesReducerEnum"
            ) :
        ]


class DashboardMosaicLayoutTilesWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterSecondaryAggregationPerSeriesAlignerEnum(
    object
):
    @classmethod
    def to_proto(self, resource):
        if not resource:
            return resource
        return dashboard_pb2.MonitoringDashboardMosaicLayoutTilesWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterSecondaryAggregationPerSeriesAlignerEnum.Value(
            "MonitoringDashboardMosaicLayoutTilesWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterSecondaryAggregationPerSeriesAlignerEnum%s"
            % resource
        )

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return resource
        return dashboard_pb2.MonitoringDashboardMosaicLayoutTilesWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterSecondaryAggregationPerSeriesAlignerEnum.Name(
            resource
        )[
            len(
                "MonitoringDashboardMosaicLayoutTilesWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterSecondaryAggregationPerSeriesAlignerEnum"
            ) :
        ]


class DashboardMosaicLayoutTilesWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterSecondaryAggregationCrossSeriesReducerEnum(
    object
):
    @classmethod
    def to_proto(self, resource):
        if not resource:
            return resource
        return dashboard_pb2.MonitoringDashboardMosaicLayoutTilesWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterSecondaryAggregationCrossSeriesReducerEnum.Value(
            "MonitoringDashboardMosaicLayoutTilesWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterSecondaryAggregationCrossSeriesReducerEnum%s"
            % resource
        )

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return resource
        return dashboard_pb2.MonitoringDashboardMosaicLayoutTilesWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterSecondaryAggregationCrossSeriesReducerEnum.Name(
            resource
        )[
            len(
                "MonitoringDashboardMosaicLayoutTilesWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterSecondaryAggregationCrossSeriesReducerEnum"
            ) :
        ]


class DashboardMosaicLayoutTilesWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterPickTimeSeriesFilterRankingMethodEnum(
    object
):
    @classmethod
    def to_proto(self, resource):
        if not resource:
            return resource
        return dashboard_pb2.MonitoringDashboardMosaicLayoutTilesWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterPickTimeSeriesFilterRankingMethodEnum.Value(
            "MonitoringDashboardMosaicLayoutTilesWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterPickTimeSeriesFilterRankingMethodEnum%s"
            % resource
        )

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return resource
        return dashboard_pb2.MonitoringDashboardMosaicLayoutTilesWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterPickTimeSeriesFilterRankingMethodEnum.Name(
            resource
        )[
            len(
                "MonitoringDashboardMosaicLayoutTilesWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterPickTimeSeriesFilterRankingMethodEnum"
            ) :
        ]


class DashboardMosaicLayoutTilesWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterPickTimeSeriesFilterDirectionEnum(
    object
):
    @classmethod
    def to_proto(self, resource):
        if not resource:
            return resource
        return dashboard_pb2.MonitoringDashboardMosaicLayoutTilesWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterPickTimeSeriesFilterDirectionEnum.Value(
            "MonitoringDashboardMosaicLayoutTilesWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterPickTimeSeriesFilterDirectionEnum%s"
            % resource
        )

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return resource
        return dashboard_pb2.MonitoringDashboardMosaicLayoutTilesWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterPickTimeSeriesFilterDirectionEnum.Name(
            resource
        )[
            len(
                "MonitoringDashboardMosaicLayoutTilesWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterPickTimeSeriesFilterDirectionEnum"
            ) :
        ]


class DashboardMosaicLayoutTilesWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioNumeratorAggregationPerSeriesAlignerEnum(
    object
):
    @classmethod
    def to_proto(self, resource):
        if not resource:
            return resource
        return dashboard_pb2.MonitoringDashboardMosaicLayoutTilesWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioNumeratorAggregationPerSeriesAlignerEnum.Value(
            "MonitoringDashboardMosaicLayoutTilesWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioNumeratorAggregationPerSeriesAlignerEnum%s"
            % resource
        )

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return resource
        return dashboard_pb2.MonitoringDashboardMosaicLayoutTilesWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioNumeratorAggregationPerSeriesAlignerEnum.Name(
            resource
        )[
            len(
                "MonitoringDashboardMosaicLayoutTilesWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioNumeratorAggregationPerSeriesAlignerEnum"
            ) :
        ]


class DashboardMosaicLayoutTilesWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioNumeratorAggregationCrossSeriesReducerEnum(
    object
):
    @classmethod
    def to_proto(self, resource):
        if not resource:
            return resource
        return dashboard_pb2.MonitoringDashboardMosaicLayoutTilesWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioNumeratorAggregationCrossSeriesReducerEnum.Value(
            "MonitoringDashboardMosaicLayoutTilesWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioNumeratorAggregationCrossSeriesReducerEnum%s"
            % resource
        )

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return resource
        return dashboard_pb2.MonitoringDashboardMosaicLayoutTilesWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioNumeratorAggregationCrossSeriesReducerEnum.Name(
            resource
        )[
            len(
                "MonitoringDashboardMosaicLayoutTilesWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioNumeratorAggregationCrossSeriesReducerEnum"
            ) :
        ]


class DashboardMosaicLayoutTilesWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioDenominatorAggregationPerSeriesAlignerEnum(
    object
):
    @classmethod
    def to_proto(self, resource):
        if not resource:
            return resource
        return dashboard_pb2.MonitoringDashboardMosaicLayoutTilesWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioDenominatorAggregationPerSeriesAlignerEnum.Value(
            "MonitoringDashboardMosaicLayoutTilesWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioDenominatorAggregationPerSeriesAlignerEnum%s"
            % resource
        )

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return resource
        return dashboard_pb2.MonitoringDashboardMosaicLayoutTilesWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioDenominatorAggregationPerSeriesAlignerEnum.Name(
            resource
        )[
            len(
                "MonitoringDashboardMosaicLayoutTilesWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioDenominatorAggregationPerSeriesAlignerEnum"
            ) :
        ]


class DashboardMosaicLayoutTilesWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioDenominatorAggregationCrossSeriesReducerEnum(
    object
):
    @classmethod
    def to_proto(self, resource):
        if not resource:
            return resource
        return dashboard_pb2.MonitoringDashboardMosaicLayoutTilesWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioDenominatorAggregationCrossSeriesReducerEnum.Value(
            "MonitoringDashboardMosaicLayoutTilesWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioDenominatorAggregationCrossSeriesReducerEnum%s"
            % resource
        )

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return resource
        return dashboard_pb2.MonitoringDashboardMosaicLayoutTilesWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioDenominatorAggregationCrossSeriesReducerEnum.Name(
            resource
        )[
            len(
                "MonitoringDashboardMosaicLayoutTilesWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioDenominatorAggregationCrossSeriesReducerEnum"
            ) :
        ]


class DashboardMosaicLayoutTilesWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioSecondaryAggregationPerSeriesAlignerEnum(
    object
):
    @classmethod
    def to_proto(self, resource):
        if not resource:
            return resource
        return dashboard_pb2.MonitoringDashboardMosaicLayoutTilesWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioSecondaryAggregationPerSeriesAlignerEnum.Value(
            "MonitoringDashboardMosaicLayoutTilesWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioSecondaryAggregationPerSeriesAlignerEnum%s"
            % resource
        )

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return resource
        return dashboard_pb2.MonitoringDashboardMosaicLayoutTilesWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioSecondaryAggregationPerSeriesAlignerEnum.Name(
            resource
        )[
            len(
                "MonitoringDashboardMosaicLayoutTilesWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioSecondaryAggregationPerSeriesAlignerEnum"
            ) :
        ]


class DashboardMosaicLayoutTilesWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioSecondaryAggregationCrossSeriesReducerEnum(
    object
):
    @classmethod
    def to_proto(self, resource):
        if not resource:
            return resource
        return dashboard_pb2.MonitoringDashboardMosaicLayoutTilesWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioSecondaryAggregationCrossSeriesReducerEnum.Value(
            "MonitoringDashboardMosaicLayoutTilesWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioSecondaryAggregationCrossSeriesReducerEnum%s"
            % resource
        )

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return resource
        return dashboard_pb2.MonitoringDashboardMosaicLayoutTilesWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioSecondaryAggregationCrossSeriesReducerEnum.Name(
            resource
        )[
            len(
                "MonitoringDashboardMosaicLayoutTilesWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioSecondaryAggregationCrossSeriesReducerEnum"
            ) :
        ]


class DashboardMosaicLayoutTilesWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioPickTimeSeriesFilterRankingMethodEnum(
    object
):
    @classmethod
    def to_proto(self, resource):
        if not resource:
            return resource
        return dashboard_pb2.MonitoringDashboardMosaicLayoutTilesWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioPickTimeSeriesFilterRankingMethodEnum.Value(
            "MonitoringDashboardMosaicLayoutTilesWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioPickTimeSeriesFilterRankingMethodEnum%s"
            % resource
        )

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return resource
        return dashboard_pb2.MonitoringDashboardMosaicLayoutTilesWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioPickTimeSeriesFilterRankingMethodEnum.Name(
            resource
        )[
            len(
                "MonitoringDashboardMosaicLayoutTilesWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioPickTimeSeriesFilterRankingMethodEnum"
            ) :
        ]


class DashboardMosaicLayoutTilesWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioPickTimeSeriesFilterDirectionEnum(
    object
):
    @classmethod
    def to_proto(self, resource):
        if not resource:
            return resource
        return dashboard_pb2.MonitoringDashboardMosaicLayoutTilesWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioPickTimeSeriesFilterDirectionEnum.Value(
            "MonitoringDashboardMosaicLayoutTilesWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioPickTimeSeriesFilterDirectionEnum%s"
            % resource
        )

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return resource
        return dashboard_pb2.MonitoringDashboardMosaicLayoutTilesWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioPickTimeSeriesFilterDirectionEnum.Name(
            resource
        )[
            len(
                "MonitoringDashboardMosaicLayoutTilesWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioPickTimeSeriesFilterDirectionEnum"
            ) :
        ]


class DashboardMosaicLayoutTilesWidgetXyChartDataSetsPlotTypeEnum(object):
    @classmethod
    def to_proto(self, resource):
        if not resource:
            return resource
        return dashboard_pb2.MonitoringDashboardMosaicLayoutTilesWidgetXyChartDataSetsPlotTypeEnum.Value(
            "MonitoringDashboardMosaicLayoutTilesWidgetXyChartDataSetsPlotTypeEnum%s"
            % resource
        )

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return resource
        return dashboard_pb2.MonitoringDashboardMosaicLayoutTilesWidgetXyChartDataSetsPlotTypeEnum.Name(
            resource
        )[
            len(
                "MonitoringDashboardMosaicLayoutTilesWidgetXyChartDataSetsPlotTypeEnum"
            ) :
        ]


class DashboardMosaicLayoutTilesWidgetXyChartThresholdsColorEnum(object):
    @classmethod
    def to_proto(self, resource):
        if not resource:
            return resource
        return dashboard_pb2.MonitoringDashboardMosaicLayoutTilesWidgetXyChartThresholdsColorEnum.Value(
            "MonitoringDashboardMosaicLayoutTilesWidgetXyChartThresholdsColorEnum%s"
            % resource
        )

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return resource
        return dashboard_pb2.MonitoringDashboardMosaicLayoutTilesWidgetXyChartThresholdsColorEnum.Name(
            resource
        )[
            len(
                "MonitoringDashboardMosaicLayoutTilesWidgetXyChartThresholdsColorEnum"
            ) :
        ]


class DashboardMosaicLayoutTilesWidgetXyChartThresholdsDirectionEnum(object):
    @classmethod
    def to_proto(self, resource):
        if not resource:
            return resource
        return dashboard_pb2.MonitoringDashboardMosaicLayoutTilesWidgetXyChartThresholdsDirectionEnum.Value(
            "MonitoringDashboardMosaicLayoutTilesWidgetXyChartThresholdsDirectionEnum%s"
            % resource
        )

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return resource
        return dashboard_pb2.MonitoringDashboardMosaicLayoutTilesWidgetXyChartThresholdsDirectionEnum.Name(
            resource
        )[
            len(
                "MonitoringDashboardMosaicLayoutTilesWidgetXyChartThresholdsDirectionEnum"
            ) :
        ]


class DashboardMosaicLayoutTilesWidgetXyChartXAxisScaleEnum(object):
    @classmethod
    def to_proto(self, resource):
        if not resource:
            return resource
        return dashboard_pb2.MonitoringDashboardMosaicLayoutTilesWidgetXyChartXAxisScaleEnum.Value(
            "MonitoringDashboardMosaicLayoutTilesWidgetXyChartXAxisScaleEnum%s"
            % resource
        )

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return resource
        return dashboard_pb2.MonitoringDashboardMosaicLayoutTilesWidgetXyChartXAxisScaleEnum.Name(
            resource
        )[
            len("MonitoringDashboardMosaicLayoutTilesWidgetXyChartXAxisScaleEnum") :
        ]


class DashboardMosaicLayoutTilesWidgetXyChartYAxisScaleEnum(object):
    @classmethod
    def to_proto(self, resource):
        if not resource:
            return resource
        return dashboard_pb2.MonitoringDashboardMosaicLayoutTilesWidgetXyChartYAxisScaleEnum.Value(
            "MonitoringDashboardMosaicLayoutTilesWidgetXyChartYAxisScaleEnum%s"
            % resource
        )

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return resource
        return dashboard_pb2.MonitoringDashboardMosaicLayoutTilesWidgetXyChartYAxisScaleEnum.Name(
            resource
        )[
            len("MonitoringDashboardMosaicLayoutTilesWidgetXyChartYAxisScaleEnum") :
        ]


class DashboardMosaicLayoutTilesWidgetXyChartChartOptionsModeEnum(object):
    @classmethod
    def to_proto(self, resource):
        if not resource:
            return resource
        return dashboard_pb2.MonitoringDashboardMosaicLayoutTilesWidgetXyChartChartOptionsModeEnum.Value(
            "MonitoringDashboardMosaicLayoutTilesWidgetXyChartChartOptionsModeEnum%s"
            % resource
        )

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return resource
        return dashboard_pb2.MonitoringDashboardMosaicLayoutTilesWidgetXyChartChartOptionsModeEnum.Name(
            resource
        )[
            len(
                "MonitoringDashboardMosaicLayoutTilesWidgetXyChartChartOptionsModeEnum"
            ) :
        ]


class DashboardMosaicLayoutTilesWidgetScorecardTimeSeriesQueryTimeSeriesFilterAggregationPerSeriesAlignerEnum(
    object
):
    @classmethod
    def to_proto(self, resource):
        if not resource:
            return resource
        return dashboard_pb2.MonitoringDashboardMosaicLayoutTilesWidgetScorecardTimeSeriesQueryTimeSeriesFilterAggregationPerSeriesAlignerEnum.Value(
            "MonitoringDashboardMosaicLayoutTilesWidgetScorecardTimeSeriesQueryTimeSeriesFilterAggregationPerSeriesAlignerEnum%s"
            % resource
        )

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return resource
        return dashboard_pb2.MonitoringDashboardMosaicLayoutTilesWidgetScorecardTimeSeriesQueryTimeSeriesFilterAggregationPerSeriesAlignerEnum.Name(
            resource
        )[
            len(
                "MonitoringDashboardMosaicLayoutTilesWidgetScorecardTimeSeriesQueryTimeSeriesFilterAggregationPerSeriesAlignerEnum"
            ) :
        ]


class DashboardMosaicLayoutTilesWidgetScorecardTimeSeriesQueryTimeSeriesFilterAggregationCrossSeriesReducerEnum(
    object
):
    @classmethod
    def to_proto(self, resource):
        if not resource:
            return resource
        return dashboard_pb2.MonitoringDashboardMosaicLayoutTilesWidgetScorecardTimeSeriesQueryTimeSeriesFilterAggregationCrossSeriesReducerEnum.Value(
            "MonitoringDashboardMosaicLayoutTilesWidgetScorecardTimeSeriesQueryTimeSeriesFilterAggregationCrossSeriesReducerEnum%s"
            % resource
        )

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return resource
        return dashboard_pb2.MonitoringDashboardMosaicLayoutTilesWidgetScorecardTimeSeriesQueryTimeSeriesFilterAggregationCrossSeriesReducerEnum.Name(
            resource
        )[
            len(
                "MonitoringDashboardMosaicLayoutTilesWidgetScorecardTimeSeriesQueryTimeSeriesFilterAggregationCrossSeriesReducerEnum"
            ) :
        ]


class DashboardMosaicLayoutTilesWidgetScorecardTimeSeriesQueryTimeSeriesFilterSecondaryAggregationPerSeriesAlignerEnum(
    object
):
    @classmethod
    def to_proto(self, resource):
        if not resource:
            return resource
        return dashboard_pb2.MonitoringDashboardMosaicLayoutTilesWidgetScorecardTimeSeriesQueryTimeSeriesFilterSecondaryAggregationPerSeriesAlignerEnum.Value(
            "MonitoringDashboardMosaicLayoutTilesWidgetScorecardTimeSeriesQueryTimeSeriesFilterSecondaryAggregationPerSeriesAlignerEnum%s"
            % resource
        )

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return resource
        return dashboard_pb2.MonitoringDashboardMosaicLayoutTilesWidgetScorecardTimeSeriesQueryTimeSeriesFilterSecondaryAggregationPerSeriesAlignerEnum.Name(
            resource
        )[
            len(
                "MonitoringDashboardMosaicLayoutTilesWidgetScorecardTimeSeriesQueryTimeSeriesFilterSecondaryAggregationPerSeriesAlignerEnum"
            ) :
        ]


class DashboardMosaicLayoutTilesWidgetScorecardTimeSeriesQueryTimeSeriesFilterSecondaryAggregationCrossSeriesReducerEnum(
    object
):
    @classmethod
    def to_proto(self, resource):
        if not resource:
            return resource
        return dashboard_pb2.MonitoringDashboardMosaicLayoutTilesWidgetScorecardTimeSeriesQueryTimeSeriesFilterSecondaryAggregationCrossSeriesReducerEnum.Value(
            "MonitoringDashboardMosaicLayoutTilesWidgetScorecardTimeSeriesQueryTimeSeriesFilterSecondaryAggregationCrossSeriesReducerEnum%s"
            % resource
        )

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return resource
        return dashboard_pb2.MonitoringDashboardMosaicLayoutTilesWidgetScorecardTimeSeriesQueryTimeSeriesFilterSecondaryAggregationCrossSeriesReducerEnum.Name(
            resource
        )[
            len(
                "MonitoringDashboardMosaicLayoutTilesWidgetScorecardTimeSeriesQueryTimeSeriesFilterSecondaryAggregationCrossSeriesReducerEnum"
            ) :
        ]


class DashboardMosaicLayoutTilesWidgetScorecardTimeSeriesQueryTimeSeriesFilterPickTimeSeriesFilterRankingMethodEnum(
    object
):
    @classmethod
    def to_proto(self, resource):
        if not resource:
            return resource
        return dashboard_pb2.MonitoringDashboardMosaicLayoutTilesWidgetScorecardTimeSeriesQueryTimeSeriesFilterPickTimeSeriesFilterRankingMethodEnum.Value(
            "MonitoringDashboardMosaicLayoutTilesWidgetScorecardTimeSeriesQueryTimeSeriesFilterPickTimeSeriesFilterRankingMethodEnum%s"
            % resource
        )

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return resource
        return dashboard_pb2.MonitoringDashboardMosaicLayoutTilesWidgetScorecardTimeSeriesQueryTimeSeriesFilterPickTimeSeriesFilterRankingMethodEnum.Name(
            resource
        )[
            len(
                "MonitoringDashboardMosaicLayoutTilesWidgetScorecardTimeSeriesQueryTimeSeriesFilterPickTimeSeriesFilterRankingMethodEnum"
            ) :
        ]


class DashboardMosaicLayoutTilesWidgetScorecardTimeSeriesQueryTimeSeriesFilterPickTimeSeriesFilterDirectionEnum(
    object
):
    @classmethod
    def to_proto(self, resource):
        if not resource:
            return resource
        return dashboard_pb2.MonitoringDashboardMosaicLayoutTilesWidgetScorecardTimeSeriesQueryTimeSeriesFilterPickTimeSeriesFilterDirectionEnum.Value(
            "MonitoringDashboardMosaicLayoutTilesWidgetScorecardTimeSeriesQueryTimeSeriesFilterPickTimeSeriesFilterDirectionEnum%s"
            % resource
        )

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return resource
        return dashboard_pb2.MonitoringDashboardMosaicLayoutTilesWidgetScorecardTimeSeriesQueryTimeSeriesFilterPickTimeSeriesFilterDirectionEnum.Name(
            resource
        )[
            len(
                "MonitoringDashboardMosaicLayoutTilesWidgetScorecardTimeSeriesQueryTimeSeriesFilterPickTimeSeriesFilterDirectionEnum"
            ) :
        ]


class DashboardMosaicLayoutTilesWidgetScorecardTimeSeriesQueryTimeSeriesFilterRatioNumeratorAggregationPerSeriesAlignerEnum(
    object
):
    @classmethod
    def to_proto(self, resource):
        if not resource:
            return resource
        return dashboard_pb2.MonitoringDashboardMosaicLayoutTilesWidgetScorecardTimeSeriesQueryTimeSeriesFilterRatioNumeratorAggregationPerSeriesAlignerEnum.Value(
            "MonitoringDashboardMosaicLayoutTilesWidgetScorecardTimeSeriesQueryTimeSeriesFilterRatioNumeratorAggregationPerSeriesAlignerEnum%s"
            % resource
        )

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return resource
        return dashboard_pb2.MonitoringDashboardMosaicLayoutTilesWidgetScorecardTimeSeriesQueryTimeSeriesFilterRatioNumeratorAggregationPerSeriesAlignerEnum.Name(
            resource
        )[
            len(
                "MonitoringDashboardMosaicLayoutTilesWidgetScorecardTimeSeriesQueryTimeSeriesFilterRatioNumeratorAggregationPerSeriesAlignerEnum"
            ) :
        ]


class DashboardMosaicLayoutTilesWidgetScorecardTimeSeriesQueryTimeSeriesFilterRatioNumeratorAggregationCrossSeriesReducerEnum(
    object
):
    @classmethod
    def to_proto(self, resource):
        if not resource:
            return resource
        return dashboard_pb2.MonitoringDashboardMosaicLayoutTilesWidgetScorecardTimeSeriesQueryTimeSeriesFilterRatioNumeratorAggregationCrossSeriesReducerEnum.Value(
            "MonitoringDashboardMosaicLayoutTilesWidgetScorecardTimeSeriesQueryTimeSeriesFilterRatioNumeratorAggregationCrossSeriesReducerEnum%s"
            % resource
        )

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return resource
        return dashboard_pb2.MonitoringDashboardMosaicLayoutTilesWidgetScorecardTimeSeriesQueryTimeSeriesFilterRatioNumeratorAggregationCrossSeriesReducerEnum.Name(
            resource
        )[
            len(
                "MonitoringDashboardMosaicLayoutTilesWidgetScorecardTimeSeriesQueryTimeSeriesFilterRatioNumeratorAggregationCrossSeriesReducerEnum"
            ) :
        ]


class DashboardMosaicLayoutTilesWidgetScorecardTimeSeriesQueryTimeSeriesFilterRatioDenominatorAggregationPerSeriesAlignerEnum(
    object
):
    @classmethod
    def to_proto(self, resource):
        if not resource:
            return resource
        return dashboard_pb2.MonitoringDashboardMosaicLayoutTilesWidgetScorecardTimeSeriesQueryTimeSeriesFilterRatioDenominatorAggregationPerSeriesAlignerEnum.Value(
            "MonitoringDashboardMosaicLayoutTilesWidgetScorecardTimeSeriesQueryTimeSeriesFilterRatioDenominatorAggregationPerSeriesAlignerEnum%s"
            % resource
        )

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return resource
        return dashboard_pb2.MonitoringDashboardMosaicLayoutTilesWidgetScorecardTimeSeriesQueryTimeSeriesFilterRatioDenominatorAggregationPerSeriesAlignerEnum.Name(
            resource
        )[
            len(
                "MonitoringDashboardMosaicLayoutTilesWidgetScorecardTimeSeriesQueryTimeSeriesFilterRatioDenominatorAggregationPerSeriesAlignerEnum"
            ) :
        ]


class DashboardMosaicLayoutTilesWidgetScorecardTimeSeriesQueryTimeSeriesFilterRatioDenominatorAggregationCrossSeriesReducerEnum(
    object
):
    @classmethod
    def to_proto(self, resource):
        if not resource:
            return resource
        return dashboard_pb2.MonitoringDashboardMosaicLayoutTilesWidgetScorecardTimeSeriesQueryTimeSeriesFilterRatioDenominatorAggregationCrossSeriesReducerEnum.Value(
            "MonitoringDashboardMosaicLayoutTilesWidgetScorecardTimeSeriesQueryTimeSeriesFilterRatioDenominatorAggregationCrossSeriesReducerEnum%s"
            % resource
        )

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return resource
        return dashboard_pb2.MonitoringDashboardMosaicLayoutTilesWidgetScorecardTimeSeriesQueryTimeSeriesFilterRatioDenominatorAggregationCrossSeriesReducerEnum.Name(
            resource
        )[
            len(
                "MonitoringDashboardMosaicLayoutTilesWidgetScorecardTimeSeriesQueryTimeSeriesFilterRatioDenominatorAggregationCrossSeriesReducerEnum"
            ) :
        ]


class DashboardMosaicLayoutTilesWidgetScorecardTimeSeriesQueryTimeSeriesFilterRatioSecondaryAggregationPerSeriesAlignerEnum(
    object
):
    @classmethod
    def to_proto(self, resource):
        if not resource:
            return resource
        return dashboard_pb2.MonitoringDashboardMosaicLayoutTilesWidgetScorecardTimeSeriesQueryTimeSeriesFilterRatioSecondaryAggregationPerSeriesAlignerEnum.Value(
            "MonitoringDashboardMosaicLayoutTilesWidgetScorecardTimeSeriesQueryTimeSeriesFilterRatioSecondaryAggregationPerSeriesAlignerEnum%s"
            % resource
        )

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return resource
        return dashboard_pb2.MonitoringDashboardMosaicLayoutTilesWidgetScorecardTimeSeriesQueryTimeSeriesFilterRatioSecondaryAggregationPerSeriesAlignerEnum.Name(
            resource
        )[
            len(
                "MonitoringDashboardMosaicLayoutTilesWidgetScorecardTimeSeriesQueryTimeSeriesFilterRatioSecondaryAggregationPerSeriesAlignerEnum"
            ) :
        ]


class DashboardMosaicLayoutTilesWidgetScorecardTimeSeriesQueryTimeSeriesFilterRatioSecondaryAggregationCrossSeriesReducerEnum(
    object
):
    @classmethod
    def to_proto(self, resource):
        if not resource:
            return resource
        return dashboard_pb2.MonitoringDashboardMosaicLayoutTilesWidgetScorecardTimeSeriesQueryTimeSeriesFilterRatioSecondaryAggregationCrossSeriesReducerEnum.Value(
            "MonitoringDashboardMosaicLayoutTilesWidgetScorecardTimeSeriesQueryTimeSeriesFilterRatioSecondaryAggregationCrossSeriesReducerEnum%s"
            % resource
        )

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return resource
        return dashboard_pb2.MonitoringDashboardMosaicLayoutTilesWidgetScorecardTimeSeriesQueryTimeSeriesFilterRatioSecondaryAggregationCrossSeriesReducerEnum.Name(
            resource
        )[
            len(
                "MonitoringDashboardMosaicLayoutTilesWidgetScorecardTimeSeriesQueryTimeSeriesFilterRatioSecondaryAggregationCrossSeriesReducerEnum"
            ) :
        ]


class DashboardMosaicLayoutTilesWidgetScorecardTimeSeriesQueryTimeSeriesFilterRatioPickTimeSeriesFilterRankingMethodEnum(
    object
):
    @classmethod
    def to_proto(self, resource):
        if not resource:
            return resource
        return dashboard_pb2.MonitoringDashboardMosaicLayoutTilesWidgetScorecardTimeSeriesQueryTimeSeriesFilterRatioPickTimeSeriesFilterRankingMethodEnum.Value(
            "MonitoringDashboardMosaicLayoutTilesWidgetScorecardTimeSeriesQueryTimeSeriesFilterRatioPickTimeSeriesFilterRankingMethodEnum%s"
            % resource
        )

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return resource
        return dashboard_pb2.MonitoringDashboardMosaicLayoutTilesWidgetScorecardTimeSeriesQueryTimeSeriesFilterRatioPickTimeSeriesFilterRankingMethodEnum.Name(
            resource
        )[
            len(
                "MonitoringDashboardMosaicLayoutTilesWidgetScorecardTimeSeriesQueryTimeSeriesFilterRatioPickTimeSeriesFilterRankingMethodEnum"
            ) :
        ]


class DashboardMosaicLayoutTilesWidgetScorecardTimeSeriesQueryTimeSeriesFilterRatioPickTimeSeriesFilterDirectionEnum(
    object
):
    @classmethod
    def to_proto(self, resource):
        if not resource:
            return resource
        return dashboard_pb2.MonitoringDashboardMosaicLayoutTilesWidgetScorecardTimeSeriesQueryTimeSeriesFilterRatioPickTimeSeriesFilterDirectionEnum.Value(
            "MonitoringDashboardMosaicLayoutTilesWidgetScorecardTimeSeriesQueryTimeSeriesFilterRatioPickTimeSeriesFilterDirectionEnum%s"
            % resource
        )

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return resource
        return dashboard_pb2.MonitoringDashboardMosaicLayoutTilesWidgetScorecardTimeSeriesQueryTimeSeriesFilterRatioPickTimeSeriesFilterDirectionEnum.Name(
            resource
        )[
            len(
                "MonitoringDashboardMosaicLayoutTilesWidgetScorecardTimeSeriesQueryTimeSeriesFilterRatioPickTimeSeriesFilterDirectionEnum"
            ) :
        ]


class DashboardMosaicLayoutTilesWidgetScorecardSparkChartViewSparkChartTypeEnum(object):
    @classmethod
    def to_proto(self, resource):
        if not resource:
            return resource
        return dashboard_pb2.MonitoringDashboardMosaicLayoutTilesWidgetScorecardSparkChartViewSparkChartTypeEnum.Value(
            "MonitoringDashboardMosaicLayoutTilesWidgetScorecardSparkChartViewSparkChartTypeEnum%s"
            % resource
        )

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return resource
        return dashboard_pb2.MonitoringDashboardMosaicLayoutTilesWidgetScorecardSparkChartViewSparkChartTypeEnum.Name(
            resource
        )[
            len(
                "MonitoringDashboardMosaicLayoutTilesWidgetScorecardSparkChartViewSparkChartTypeEnum"
            ) :
        ]


class DashboardMosaicLayoutTilesWidgetScorecardThresholdsColorEnum(object):
    @classmethod
    def to_proto(self, resource):
        if not resource:
            return resource
        return dashboard_pb2.MonitoringDashboardMosaicLayoutTilesWidgetScorecardThresholdsColorEnum.Value(
            "MonitoringDashboardMosaicLayoutTilesWidgetScorecardThresholdsColorEnum%s"
            % resource
        )

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return resource
        return dashboard_pb2.MonitoringDashboardMosaicLayoutTilesWidgetScorecardThresholdsColorEnum.Name(
            resource
        )[
            len(
                "MonitoringDashboardMosaicLayoutTilesWidgetScorecardThresholdsColorEnum"
            ) :
        ]


class DashboardMosaicLayoutTilesWidgetScorecardThresholdsDirectionEnum(object):
    @classmethod
    def to_proto(self, resource):
        if not resource:
            return resource
        return dashboard_pb2.MonitoringDashboardMosaicLayoutTilesWidgetScorecardThresholdsDirectionEnum.Value(
            "MonitoringDashboardMosaicLayoutTilesWidgetScorecardThresholdsDirectionEnum%s"
            % resource
        )

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return resource
        return dashboard_pb2.MonitoringDashboardMosaicLayoutTilesWidgetScorecardThresholdsDirectionEnum.Name(
            resource
        )[
            len(
                "MonitoringDashboardMosaicLayoutTilesWidgetScorecardThresholdsDirectionEnum"
            ) :
        ]


class DashboardMosaicLayoutTilesWidgetTextFormatEnum(object):
    @classmethod
    def to_proto(self, resource):
        if not resource:
            return resource
        return dashboard_pb2.MonitoringDashboardMosaicLayoutTilesWidgetTextFormatEnum.Value(
            "MonitoringDashboardMosaicLayoutTilesWidgetTextFormatEnum%s" % resource
        )

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return resource
        return (
            dashboard_pb2.MonitoringDashboardMosaicLayoutTilesWidgetTextFormatEnum.Name(
                resource
            )[len("MonitoringDashboardMosaicLayoutTilesWidgetTextFormatEnum") :]
        )


class DashboardRowLayoutRowsWidgetsXyChartDataSetsTimeSeriesQueryTimeSeriesFilterAggregationPerSeriesAlignerEnum(
    object
):
    @classmethod
    def to_proto(self, resource):
        if not resource:
            return resource
        return dashboard_pb2.MonitoringDashboardRowLayoutRowsWidgetsXyChartDataSetsTimeSeriesQueryTimeSeriesFilterAggregationPerSeriesAlignerEnum.Value(
            "MonitoringDashboardRowLayoutRowsWidgetsXyChartDataSetsTimeSeriesQueryTimeSeriesFilterAggregationPerSeriesAlignerEnum%s"
            % resource
        )

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return resource
        return dashboard_pb2.MonitoringDashboardRowLayoutRowsWidgetsXyChartDataSetsTimeSeriesQueryTimeSeriesFilterAggregationPerSeriesAlignerEnum.Name(
            resource
        )[
            len(
                "MonitoringDashboardRowLayoutRowsWidgetsXyChartDataSetsTimeSeriesQueryTimeSeriesFilterAggregationPerSeriesAlignerEnum"
            ) :
        ]


class DashboardRowLayoutRowsWidgetsXyChartDataSetsTimeSeriesQueryTimeSeriesFilterAggregationCrossSeriesReducerEnum(
    object
):
    @classmethod
    def to_proto(self, resource):
        if not resource:
            return resource
        return dashboard_pb2.MonitoringDashboardRowLayoutRowsWidgetsXyChartDataSetsTimeSeriesQueryTimeSeriesFilterAggregationCrossSeriesReducerEnum.Value(
            "MonitoringDashboardRowLayoutRowsWidgetsXyChartDataSetsTimeSeriesQueryTimeSeriesFilterAggregationCrossSeriesReducerEnum%s"
            % resource
        )

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return resource
        return dashboard_pb2.MonitoringDashboardRowLayoutRowsWidgetsXyChartDataSetsTimeSeriesQueryTimeSeriesFilterAggregationCrossSeriesReducerEnum.Name(
            resource
        )[
            len(
                "MonitoringDashboardRowLayoutRowsWidgetsXyChartDataSetsTimeSeriesQueryTimeSeriesFilterAggregationCrossSeriesReducerEnum"
            ) :
        ]


class DashboardRowLayoutRowsWidgetsXyChartDataSetsTimeSeriesQueryTimeSeriesFilterSecondaryAggregationPerSeriesAlignerEnum(
    object
):
    @classmethod
    def to_proto(self, resource):
        if not resource:
            return resource
        return dashboard_pb2.MonitoringDashboardRowLayoutRowsWidgetsXyChartDataSetsTimeSeriesQueryTimeSeriesFilterSecondaryAggregationPerSeriesAlignerEnum.Value(
            "MonitoringDashboardRowLayoutRowsWidgetsXyChartDataSetsTimeSeriesQueryTimeSeriesFilterSecondaryAggregationPerSeriesAlignerEnum%s"
            % resource
        )

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return resource
        return dashboard_pb2.MonitoringDashboardRowLayoutRowsWidgetsXyChartDataSetsTimeSeriesQueryTimeSeriesFilterSecondaryAggregationPerSeriesAlignerEnum.Name(
            resource
        )[
            len(
                "MonitoringDashboardRowLayoutRowsWidgetsXyChartDataSetsTimeSeriesQueryTimeSeriesFilterSecondaryAggregationPerSeriesAlignerEnum"
            ) :
        ]


class DashboardRowLayoutRowsWidgetsXyChartDataSetsTimeSeriesQueryTimeSeriesFilterSecondaryAggregationCrossSeriesReducerEnum(
    object
):
    @classmethod
    def to_proto(self, resource):
        if not resource:
            return resource
        return dashboard_pb2.MonitoringDashboardRowLayoutRowsWidgetsXyChartDataSetsTimeSeriesQueryTimeSeriesFilterSecondaryAggregationCrossSeriesReducerEnum.Value(
            "MonitoringDashboardRowLayoutRowsWidgetsXyChartDataSetsTimeSeriesQueryTimeSeriesFilterSecondaryAggregationCrossSeriesReducerEnum%s"
            % resource
        )

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return resource
        return dashboard_pb2.MonitoringDashboardRowLayoutRowsWidgetsXyChartDataSetsTimeSeriesQueryTimeSeriesFilterSecondaryAggregationCrossSeriesReducerEnum.Name(
            resource
        )[
            len(
                "MonitoringDashboardRowLayoutRowsWidgetsXyChartDataSetsTimeSeriesQueryTimeSeriesFilterSecondaryAggregationCrossSeriesReducerEnum"
            ) :
        ]


class DashboardRowLayoutRowsWidgetsXyChartDataSetsTimeSeriesQueryTimeSeriesFilterPickTimeSeriesFilterRankingMethodEnum(
    object
):
    @classmethod
    def to_proto(self, resource):
        if not resource:
            return resource
        return dashboard_pb2.MonitoringDashboardRowLayoutRowsWidgetsXyChartDataSetsTimeSeriesQueryTimeSeriesFilterPickTimeSeriesFilterRankingMethodEnum.Value(
            "MonitoringDashboardRowLayoutRowsWidgetsXyChartDataSetsTimeSeriesQueryTimeSeriesFilterPickTimeSeriesFilterRankingMethodEnum%s"
            % resource
        )

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return resource
        return dashboard_pb2.MonitoringDashboardRowLayoutRowsWidgetsXyChartDataSetsTimeSeriesQueryTimeSeriesFilterPickTimeSeriesFilterRankingMethodEnum.Name(
            resource
        )[
            len(
                "MonitoringDashboardRowLayoutRowsWidgetsXyChartDataSetsTimeSeriesQueryTimeSeriesFilterPickTimeSeriesFilterRankingMethodEnum"
            ) :
        ]


class DashboardRowLayoutRowsWidgetsXyChartDataSetsTimeSeriesQueryTimeSeriesFilterPickTimeSeriesFilterDirectionEnum(
    object
):
    @classmethod
    def to_proto(self, resource):
        if not resource:
            return resource
        return dashboard_pb2.MonitoringDashboardRowLayoutRowsWidgetsXyChartDataSetsTimeSeriesQueryTimeSeriesFilterPickTimeSeriesFilterDirectionEnum.Value(
            "MonitoringDashboardRowLayoutRowsWidgetsXyChartDataSetsTimeSeriesQueryTimeSeriesFilterPickTimeSeriesFilterDirectionEnum%s"
            % resource
        )

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return resource
        return dashboard_pb2.MonitoringDashboardRowLayoutRowsWidgetsXyChartDataSetsTimeSeriesQueryTimeSeriesFilterPickTimeSeriesFilterDirectionEnum.Name(
            resource
        )[
            len(
                "MonitoringDashboardRowLayoutRowsWidgetsXyChartDataSetsTimeSeriesQueryTimeSeriesFilterPickTimeSeriesFilterDirectionEnum"
            ) :
        ]


class DashboardRowLayoutRowsWidgetsXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioNumeratorAggregationPerSeriesAlignerEnum(
    object
):
    @classmethod
    def to_proto(self, resource):
        if not resource:
            return resource
        return dashboard_pb2.MonitoringDashboardRowLayoutRowsWidgetsXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioNumeratorAggregationPerSeriesAlignerEnum.Value(
            "MonitoringDashboardRowLayoutRowsWidgetsXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioNumeratorAggregationPerSeriesAlignerEnum%s"
            % resource
        )

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return resource
        return dashboard_pb2.MonitoringDashboardRowLayoutRowsWidgetsXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioNumeratorAggregationPerSeriesAlignerEnum.Name(
            resource
        )[
            len(
                "MonitoringDashboardRowLayoutRowsWidgetsXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioNumeratorAggregationPerSeriesAlignerEnum"
            ) :
        ]


class DashboardRowLayoutRowsWidgetsXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioNumeratorAggregationCrossSeriesReducerEnum(
    object
):
    @classmethod
    def to_proto(self, resource):
        if not resource:
            return resource
        return dashboard_pb2.MonitoringDashboardRowLayoutRowsWidgetsXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioNumeratorAggregationCrossSeriesReducerEnum.Value(
            "MonitoringDashboardRowLayoutRowsWidgetsXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioNumeratorAggregationCrossSeriesReducerEnum%s"
            % resource
        )

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return resource
        return dashboard_pb2.MonitoringDashboardRowLayoutRowsWidgetsXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioNumeratorAggregationCrossSeriesReducerEnum.Name(
            resource
        )[
            len(
                "MonitoringDashboardRowLayoutRowsWidgetsXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioNumeratorAggregationCrossSeriesReducerEnum"
            ) :
        ]


class DashboardRowLayoutRowsWidgetsXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioDenominatorAggregationPerSeriesAlignerEnum(
    object
):
    @classmethod
    def to_proto(self, resource):
        if not resource:
            return resource
        return dashboard_pb2.MonitoringDashboardRowLayoutRowsWidgetsXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioDenominatorAggregationPerSeriesAlignerEnum.Value(
            "MonitoringDashboardRowLayoutRowsWidgetsXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioDenominatorAggregationPerSeriesAlignerEnum%s"
            % resource
        )

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return resource
        return dashboard_pb2.MonitoringDashboardRowLayoutRowsWidgetsXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioDenominatorAggregationPerSeriesAlignerEnum.Name(
            resource
        )[
            len(
                "MonitoringDashboardRowLayoutRowsWidgetsXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioDenominatorAggregationPerSeriesAlignerEnum"
            ) :
        ]


class DashboardRowLayoutRowsWidgetsXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioDenominatorAggregationCrossSeriesReducerEnum(
    object
):
    @classmethod
    def to_proto(self, resource):
        if not resource:
            return resource
        return dashboard_pb2.MonitoringDashboardRowLayoutRowsWidgetsXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioDenominatorAggregationCrossSeriesReducerEnum.Value(
            "MonitoringDashboardRowLayoutRowsWidgetsXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioDenominatorAggregationCrossSeriesReducerEnum%s"
            % resource
        )

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return resource
        return dashboard_pb2.MonitoringDashboardRowLayoutRowsWidgetsXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioDenominatorAggregationCrossSeriesReducerEnum.Name(
            resource
        )[
            len(
                "MonitoringDashboardRowLayoutRowsWidgetsXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioDenominatorAggregationCrossSeriesReducerEnum"
            ) :
        ]


class DashboardRowLayoutRowsWidgetsXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioSecondaryAggregationPerSeriesAlignerEnum(
    object
):
    @classmethod
    def to_proto(self, resource):
        if not resource:
            return resource
        return dashboard_pb2.MonitoringDashboardRowLayoutRowsWidgetsXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioSecondaryAggregationPerSeriesAlignerEnum.Value(
            "MonitoringDashboardRowLayoutRowsWidgetsXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioSecondaryAggregationPerSeriesAlignerEnum%s"
            % resource
        )

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return resource
        return dashboard_pb2.MonitoringDashboardRowLayoutRowsWidgetsXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioSecondaryAggregationPerSeriesAlignerEnum.Name(
            resource
        )[
            len(
                "MonitoringDashboardRowLayoutRowsWidgetsXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioSecondaryAggregationPerSeriesAlignerEnum"
            ) :
        ]


class DashboardRowLayoutRowsWidgetsXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioSecondaryAggregationCrossSeriesReducerEnum(
    object
):
    @classmethod
    def to_proto(self, resource):
        if not resource:
            return resource
        return dashboard_pb2.MonitoringDashboardRowLayoutRowsWidgetsXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioSecondaryAggregationCrossSeriesReducerEnum.Value(
            "MonitoringDashboardRowLayoutRowsWidgetsXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioSecondaryAggregationCrossSeriesReducerEnum%s"
            % resource
        )

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return resource
        return dashboard_pb2.MonitoringDashboardRowLayoutRowsWidgetsXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioSecondaryAggregationCrossSeriesReducerEnum.Name(
            resource
        )[
            len(
                "MonitoringDashboardRowLayoutRowsWidgetsXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioSecondaryAggregationCrossSeriesReducerEnum"
            ) :
        ]


class DashboardRowLayoutRowsWidgetsXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioPickTimeSeriesFilterRankingMethodEnum(
    object
):
    @classmethod
    def to_proto(self, resource):
        if not resource:
            return resource
        return dashboard_pb2.MonitoringDashboardRowLayoutRowsWidgetsXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioPickTimeSeriesFilterRankingMethodEnum.Value(
            "MonitoringDashboardRowLayoutRowsWidgetsXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioPickTimeSeriesFilterRankingMethodEnum%s"
            % resource
        )

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return resource
        return dashboard_pb2.MonitoringDashboardRowLayoutRowsWidgetsXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioPickTimeSeriesFilterRankingMethodEnum.Name(
            resource
        )[
            len(
                "MonitoringDashboardRowLayoutRowsWidgetsXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioPickTimeSeriesFilterRankingMethodEnum"
            ) :
        ]


class DashboardRowLayoutRowsWidgetsXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioPickTimeSeriesFilterDirectionEnum(
    object
):
    @classmethod
    def to_proto(self, resource):
        if not resource:
            return resource
        return dashboard_pb2.MonitoringDashboardRowLayoutRowsWidgetsXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioPickTimeSeriesFilterDirectionEnum.Value(
            "MonitoringDashboardRowLayoutRowsWidgetsXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioPickTimeSeriesFilterDirectionEnum%s"
            % resource
        )

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return resource
        return dashboard_pb2.MonitoringDashboardRowLayoutRowsWidgetsXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioPickTimeSeriesFilterDirectionEnum.Name(
            resource
        )[
            len(
                "MonitoringDashboardRowLayoutRowsWidgetsXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioPickTimeSeriesFilterDirectionEnum"
            ) :
        ]


class DashboardRowLayoutRowsWidgetsXyChartDataSetsPlotTypeEnum(object):
    @classmethod
    def to_proto(self, resource):
        if not resource:
            return resource
        return dashboard_pb2.MonitoringDashboardRowLayoutRowsWidgetsXyChartDataSetsPlotTypeEnum.Value(
            "MonitoringDashboardRowLayoutRowsWidgetsXyChartDataSetsPlotTypeEnum%s"
            % resource
        )

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return resource
        return dashboard_pb2.MonitoringDashboardRowLayoutRowsWidgetsXyChartDataSetsPlotTypeEnum.Name(
            resource
        )[
            len("MonitoringDashboardRowLayoutRowsWidgetsXyChartDataSetsPlotTypeEnum") :
        ]


class DashboardRowLayoutRowsWidgetsXyChartThresholdsColorEnum(object):
    @classmethod
    def to_proto(self, resource):
        if not resource:
            return resource
        return dashboard_pb2.MonitoringDashboardRowLayoutRowsWidgetsXyChartThresholdsColorEnum.Value(
            "MonitoringDashboardRowLayoutRowsWidgetsXyChartThresholdsColorEnum%s"
            % resource
        )

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return resource
        return dashboard_pb2.MonitoringDashboardRowLayoutRowsWidgetsXyChartThresholdsColorEnum.Name(
            resource
        )[
            len("MonitoringDashboardRowLayoutRowsWidgetsXyChartThresholdsColorEnum") :
        ]


class DashboardRowLayoutRowsWidgetsXyChartThresholdsDirectionEnum(object):
    @classmethod
    def to_proto(self, resource):
        if not resource:
            return resource
        return dashboard_pb2.MonitoringDashboardRowLayoutRowsWidgetsXyChartThresholdsDirectionEnum.Value(
            "MonitoringDashboardRowLayoutRowsWidgetsXyChartThresholdsDirectionEnum%s"
            % resource
        )

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return resource
        return dashboard_pb2.MonitoringDashboardRowLayoutRowsWidgetsXyChartThresholdsDirectionEnum.Name(
            resource
        )[
            len(
                "MonitoringDashboardRowLayoutRowsWidgetsXyChartThresholdsDirectionEnum"
            ) :
        ]


class DashboardRowLayoutRowsWidgetsXyChartXAxisScaleEnum(object):
    @classmethod
    def to_proto(self, resource):
        if not resource:
            return resource
        return dashboard_pb2.MonitoringDashboardRowLayoutRowsWidgetsXyChartXAxisScaleEnum.Value(
            "MonitoringDashboardRowLayoutRowsWidgetsXyChartXAxisScaleEnum%s" % resource
        )

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return resource
        return dashboard_pb2.MonitoringDashboardRowLayoutRowsWidgetsXyChartXAxisScaleEnum.Name(
            resource
        )[
            len("MonitoringDashboardRowLayoutRowsWidgetsXyChartXAxisScaleEnum") :
        ]


class DashboardRowLayoutRowsWidgetsXyChartYAxisScaleEnum(object):
    @classmethod
    def to_proto(self, resource):
        if not resource:
            return resource
        return dashboard_pb2.MonitoringDashboardRowLayoutRowsWidgetsXyChartYAxisScaleEnum.Value(
            "MonitoringDashboardRowLayoutRowsWidgetsXyChartYAxisScaleEnum%s" % resource
        )

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return resource
        return dashboard_pb2.MonitoringDashboardRowLayoutRowsWidgetsXyChartYAxisScaleEnum.Name(
            resource
        )[
            len("MonitoringDashboardRowLayoutRowsWidgetsXyChartYAxisScaleEnum") :
        ]


class DashboardRowLayoutRowsWidgetsXyChartChartOptionsModeEnum(object):
    @classmethod
    def to_proto(self, resource):
        if not resource:
            return resource
        return dashboard_pb2.MonitoringDashboardRowLayoutRowsWidgetsXyChartChartOptionsModeEnum.Value(
            "MonitoringDashboardRowLayoutRowsWidgetsXyChartChartOptionsModeEnum%s"
            % resource
        )

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return resource
        return dashboard_pb2.MonitoringDashboardRowLayoutRowsWidgetsXyChartChartOptionsModeEnum.Name(
            resource
        )[
            len("MonitoringDashboardRowLayoutRowsWidgetsXyChartChartOptionsModeEnum") :
        ]


class DashboardRowLayoutRowsWidgetsScorecardTimeSeriesQueryTimeSeriesFilterAggregationPerSeriesAlignerEnum(
    object
):
    @classmethod
    def to_proto(self, resource):
        if not resource:
            return resource
        return dashboard_pb2.MonitoringDashboardRowLayoutRowsWidgetsScorecardTimeSeriesQueryTimeSeriesFilterAggregationPerSeriesAlignerEnum.Value(
            "MonitoringDashboardRowLayoutRowsWidgetsScorecardTimeSeriesQueryTimeSeriesFilterAggregationPerSeriesAlignerEnum%s"
            % resource
        )

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return resource
        return dashboard_pb2.MonitoringDashboardRowLayoutRowsWidgetsScorecardTimeSeriesQueryTimeSeriesFilterAggregationPerSeriesAlignerEnum.Name(
            resource
        )[
            len(
                "MonitoringDashboardRowLayoutRowsWidgetsScorecardTimeSeriesQueryTimeSeriesFilterAggregationPerSeriesAlignerEnum"
            ) :
        ]


class DashboardRowLayoutRowsWidgetsScorecardTimeSeriesQueryTimeSeriesFilterAggregationCrossSeriesReducerEnum(
    object
):
    @classmethod
    def to_proto(self, resource):
        if not resource:
            return resource
        return dashboard_pb2.MonitoringDashboardRowLayoutRowsWidgetsScorecardTimeSeriesQueryTimeSeriesFilterAggregationCrossSeriesReducerEnum.Value(
            "MonitoringDashboardRowLayoutRowsWidgetsScorecardTimeSeriesQueryTimeSeriesFilterAggregationCrossSeriesReducerEnum%s"
            % resource
        )

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return resource
        return dashboard_pb2.MonitoringDashboardRowLayoutRowsWidgetsScorecardTimeSeriesQueryTimeSeriesFilterAggregationCrossSeriesReducerEnum.Name(
            resource
        )[
            len(
                "MonitoringDashboardRowLayoutRowsWidgetsScorecardTimeSeriesQueryTimeSeriesFilterAggregationCrossSeriesReducerEnum"
            ) :
        ]


class DashboardRowLayoutRowsWidgetsScorecardTimeSeriesQueryTimeSeriesFilterSecondaryAggregationPerSeriesAlignerEnum(
    object
):
    @classmethod
    def to_proto(self, resource):
        if not resource:
            return resource
        return dashboard_pb2.MonitoringDashboardRowLayoutRowsWidgetsScorecardTimeSeriesQueryTimeSeriesFilterSecondaryAggregationPerSeriesAlignerEnum.Value(
            "MonitoringDashboardRowLayoutRowsWidgetsScorecardTimeSeriesQueryTimeSeriesFilterSecondaryAggregationPerSeriesAlignerEnum%s"
            % resource
        )

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return resource
        return dashboard_pb2.MonitoringDashboardRowLayoutRowsWidgetsScorecardTimeSeriesQueryTimeSeriesFilterSecondaryAggregationPerSeriesAlignerEnum.Name(
            resource
        )[
            len(
                "MonitoringDashboardRowLayoutRowsWidgetsScorecardTimeSeriesQueryTimeSeriesFilterSecondaryAggregationPerSeriesAlignerEnum"
            ) :
        ]


class DashboardRowLayoutRowsWidgetsScorecardTimeSeriesQueryTimeSeriesFilterSecondaryAggregationCrossSeriesReducerEnum(
    object
):
    @classmethod
    def to_proto(self, resource):
        if not resource:
            return resource
        return dashboard_pb2.MonitoringDashboardRowLayoutRowsWidgetsScorecardTimeSeriesQueryTimeSeriesFilterSecondaryAggregationCrossSeriesReducerEnum.Value(
            "MonitoringDashboardRowLayoutRowsWidgetsScorecardTimeSeriesQueryTimeSeriesFilterSecondaryAggregationCrossSeriesReducerEnum%s"
            % resource
        )

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return resource
        return dashboard_pb2.MonitoringDashboardRowLayoutRowsWidgetsScorecardTimeSeriesQueryTimeSeriesFilterSecondaryAggregationCrossSeriesReducerEnum.Name(
            resource
        )[
            len(
                "MonitoringDashboardRowLayoutRowsWidgetsScorecardTimeSeriesQueryTimeSeriesFilterSecondaryAggregationCrossSeriesReducerEnum"
            ) :
        ]


class DashboardRowLayoutRowsWidgetsScorecardTimeSeriesQueryTimeSeriesFilterPickTimeSeriesFilterRankingMethodEnum(
    object
):
    @classmethod
    def to_proto(self, resource):
        if not resource:
            return resource
        return dashboard_pb2.MonitoringDashboardRowLayoutRowsWidgetsScorecardTimeSeriesQueryTimeSeriesFilterPickTimeSeriesFilterRankingMethodEnum.Value(
            "MonitoringDashboardRowLayoutRowsWidgetsScorecardTimeSeriesQueryTimeSeriesFilterPickTimeSeriesFilterRankingMethodEnum%s"
            % resource
        )

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return resource
        return dashboard_pb2.MonitoringDashboardRowLayoutRowsWidgetsScorecardTimeSeriesQueryTimeSeriesFilterPickTimeSeriesFilterRankingMethodEnum.Name(
            resource
        )[
            len(
                "MonitoringDashboardRowLayoutRowsWidgetsScorecardTimeSeriesQueryTimeSeriesFilterPickTimeSeriesFilterRankingMethodEnum"
            ) :
        ]


class DashboardRowLayoutRowsWidgetsScorecardTimeSeriesQueryTimeSeriesFilterPickTimeSeriesFilterDirectionEnum(
    object
):
    @classmethod
    def to_proto(self, resource):
        if not resource:
            return resource
        return dashboard_pb2.MonitoringDashboardRowLayoutRowsWidgetsScorecardTimeSeriesQueryTimeSeriesFilterPickTimeSeriesFilterDirectionEnum.Value(
            "MonitoringDashboardRowLayoutRowsWidgetsScorecardTimeSeriesQueryTimeSeriesFilterPickTimeSeriesFilterDirectionEnum%s"
            % resource
        )

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return resource
        return dashboard_pb2.MonitoringDashboardRowLayoutRowsWidgetsScorecardTimeSeriesQueryTimeSeriesFilterPickTimeSeriesFilterDirectionEnum.Name(
            resource
        )[
            len(
                "MonitoringDashboardRowLayoutRowsWidgetsScorecardTimeSeriesQueryTimeSeriesFilterPickTimeSeriesFilterDirectionEnum"
            ) :
        ]


class DashboardRowLayoutRowsWidgetsScorecardTimeSeriesQueryTimeSeriesFilterRatioNumeratorAggregationPerSeriesAlignerEnum(
    object
):
    @classmethod
    def to_proto(self, resource):
        if not resource:
            return resource
        return dashboard_pb2.MonitoringDashboardRowLayoutRowsWidgetsScorecardTimeSeriesQueryTimeSeriesFilterRatioNumeratorAggregationPerSeriesAlignerEnum.Value(
            "MonitoringDashboardRowLayoutRowsWidgetsScorecardTimeSeriesQueryTimeSeriesFilterRatioNumeratorAggregationPerSeriesAlignerEnum%s"
            % resource
        )

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return resource
        return dashboard_pb2.MonitoringDashboardRowLayoutRowsWidgetsScorecardTimeSeriesQueryTimeSeriesFilterRatioNumeratorAggregationPerSeriesAlignerEnum.Name(
            resource
        )[
            len(
                "MonitoringDashboardRowLayoutRowsWidgetsScorecardTimeSeriesQueryTimeSeriesFilterRatioNumeratorAggregationPerSeriesAlignerEnum"
            ) :
        ]


class DashboardRowLayoutRowsWidgetsScorecardTimeSeriesQueryTimeSeriesFilterRatioNumeratorAggregationCrossSeriesReducerEnum(
    object
):
    @classmethod
    def to_proto(self, resource):
        if not resource:
            return resource
        return dashboard_pb2.MonitoringDashboardRowLayoutRowsWidgetsScorecardTimeSeriesQueryTimeSeriesFilterRatioNumeratorAggregationCrossSeriesReducerEnum.Value(
            "MonitoringDashboardRowLayoutRowsWidgetsScorecardTimeSeriesQueryTimeSeriesFilterRatioNumeratorAggregationCrossSeriesReducerEnum%s"
            % resource
        )

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return resource
        return dashboard_pb2.MonitoringDashboardRowLayoutRowsWidgetsScorecardTimeSeriesQueryTimeSeriesFilterRatioNumeratorAggregationCrossSeriesReducerEnum.Name(
            resource
        )[
            len(
                "MonitoringDashboardRowLayoutRowsWidgetsScorecardTimeSeriesQueryTimeSeriesFilterRatioNumeratorAggregationCrossSeriesReducerEnum"
            ) :
        ]


class DashboardRowLayoutRowsWidgetsScorecardTimeSeriesQueryTimeSeriesFilterRatioDenominatorAggregationPerSeriesAlignerEnum(
    object
):
    @classmethod
    def to_proto(self, resource):
        if not resource:
            return resource
        return dashboard_pb2.MonitoringDashboardRowLayoutRowsWidgetsScorecardTimeSeriesQueryTimeSeriesFilterRatioDenominatorAggregationPerSeriesAlignerEnum.Value(
            "MonitoringDashboardRowLayoutRowsWidgetsScorecardTimeSeriesQueryTimeSeriesFilterRatioDenominatorAggregationPerSeriesAlignerEnum%s"
            % resource
        )

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return resource
        return dashboard_pb2.MonitoringDashboardRowLayoutRowsWidgetsScorecardTimeSeriesQueryTimeSeriesFilterRatioDenominatorAggregationPerSeriesAlignerEnum.Name(
            resource
        )[
            len(
                "MonitoringDashboardRowLayoutRowsWidgetsScorecardTimeSeriesQueryTimeSeriesFilterRatioDenominatorAggregationPerSeriesAlignerEnum"
            ) :
        ]


class DashboardRowLayoutRowsWidgetsScorecardTimeSeriesQueryTimeSeriesFilterRatioDenominatorAggregationCrossSeriesReducerEnum(
    object
):
    @classmethod
    def to_proto(self, resource):
        if not resource:
            return resource
        return dashboard_pb2.MonitoringDashboardRowLayoutRowsWidgetsScorecardTimeSeriesQueryTimeSeriesFilterRatioDenominatorAggregationCrossSeriesReducerEnum.Value(
            "MonitoringDashboardRowLayoutRowsWidgetsScorecardTimeSeriesQueryTimeSeriesFilterRatioDenominatorAggregationCrossSeriesReducerEnum%s"
            % resource
        )

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return resource
        return dashboard_pb2.MonitoringDashboardRowLayoutRowsWidgetsScorecardTimeSeriesQueryTimeSeriesFilterRatioDenominatorAggregationCrossSeriesReducerEnum.Name(
            resource
        )[
            len(
                "MonitoringDashboardRowLayoutRowsWidgetsScorecardTimeSeriesQueryTimeSeriesFilterRatioDenominatorAggregationCrossSeriesReducerEnum"
            ) :
        ]


class DashboardRowLayoutRowsWidgetsScorecardTimeSeriesQueryTimeSeriesFilterRatioSecondaryAggregationPerSeriesAlignerEnum(
    object
):
    @classmethod
    def to_proto(self, resource):
        if not resource:
            return resource
        return dashboard_pb2.MonitoringDashboardRowLayoutRowsWidgetsScorecardTimeSeriesQueryTimeSeriesFilterRatioSecondaryAggregationPerSeriesAlignerEnum.Value(
            "MonitoringDashboardRowLayoutRowsWidgetsScorecardTimeSeriesQueryTimeSeriesFilterRatioSecondaryAggregationPerSeriesAlignerEnum%s"
            % resource
        )

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return resource
        return dashboard_pb2.MonitoringDashboardRowLayoutRowsWidgetsScorecardTimeSeriesQueryTimeSeriesFilterRatioSecondaryAggregationPerSeriesAlignerEnum.Name(
            resource
        )[
            len(
                "MonitoringDashboardRowLayoutRowsWidgetsScorecardTimeSeriesQueryTimeSeriesFilterRatioSecondaryAggregationPerSeriesAlignerEnum"
            ) :
        ]


class DashboardRowLayoutRowsWidgetsScorecardTimeSeriesQueryTimeSeriesFilterRatioSecondaryAggregationCrossSeriesReducerEnum(
    object
):
    @classmethod
    def to_proto(self, resource):
        if not resource:
            return resource
        return dashboard_pb2.MonitoringDashboardRowLayoutRowsWidgetsScorecardTimeSeriesQueryTimeSeriesFilterRatioSecondaryAggregationCrossSeriesReducerEnum.Value(
            "MonitoringDashboardRowLayoutRowsWidgetsScorecardTimeSeriesQueryTimeSeriesFilterRatioSecondaryAggregationCrossSeriesReducerEnum%s"
            % resource
        )

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return resource
        return dashboard_pb2.MonitoringDashboardRowLayoutRowsWidgetsScorecardTimeSeriesQueryTimeSeriesFilterRatioSecondaryAggregationCrossSeriesReducerEnum.Name(
            resource
        )[
            len(
                "MonitoringDashboardRowLayoutRowsWidgetsScorecardTimeSeriesQueryTimeSeriesFilterRatioSecondaryAggregationCrossSeriesReducerEnum"
            ) :
        ]


class DashboardRowLayoutRowsWidgetsScorecardTimeSeriesQueryTimeSeriesFilterRatioPickTimeSeriesFilterRankingMethodEnum(
    object
):
    @classmethod
    def to_proto(self, resource):
        if not resource:
            return resource
        return dashboard_pb2.MonitoringDashboardRowLayoutRowsWidgetsScorecardTimeSeriesQueryTimeSeriesFilterRatioPickTimeSeriesFilterRankingMethodEnum.Value(
            "MonitoringDashboardRowLayoutRowsWidgetsScorecardTimeSeriesQueryTimeSeriesFilterRatioPickTimeSeriesFilterRankingMethodEnum%s"
            % resource
        )

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return resource
        return dashboard_pb2.MonitoringDashboardRowLayoutRowsWidgetsScorecardTimeSeriesQueryTimeSeriesFilterRatioPickTimeSeriesFilterRankingMethodEnum.Name(
            resource
        )[
            len(
                "MonitoringDashboardRowLayoutRowsWidgetsScorecardTimeSeriesQueryTimeSeriesFilterRatioPickTimeSeriesFilterRankingMethodEnum"
            ) :
        ]


class DashboardRowLayoutRowsWidgetsScorecardTimeSeriesQueryTimeSeriesFilterRatioPickTimeSeriesFilterDirectionEnum(
    object
):
    @classmethod
    def to_proto(self, resource):
        if not resource:
            return resource
        return dashboard_pb2.MonitoringDashboardRowLayoutRowsWidgetsScorecardTimeSeriesQueryTimeSeriesFilterRatioPickTimeSeriesFilterDirectionEnum.Value(
            "MonitoringDashboardRowLayoutRowsWidgetsScorecardTimeSeriesQueryTimeSeriesFilterRatioPickTimeSeriesFilterDirectionEnum%s"
            % resource
        )

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return resource
        return dashboard_pb2.MonitoringDashboardRowLayoutRowsWidgetsScorecardTimeSeriesQueryTimeSeriesFilterRatioPickTimeSeriesFilterDirectionEnum.Name(
            resource
        )[
            len(
                "MonitoringDashboardRowLayoutRowsWidgetsScorecardTimeSeriesQueryTimeSeriesFilterRatioPickTimeSeriesFilterDirectionEnum"
            ) :
        ]


class DashboardRowLayoutRowsWidgetsScorecardSparkChartViewSparkChartTypeEnum(object):
    @classmethod
    def to_proto(self, resource):
        if not resource:
            return resource
        return dashboard_pb2.MonitoringDashboardRowLayoutRowsWidgetsScorecardSparkChartViewSparkChartTypeEnum.Value(
            "MonitoringDashboardRowLayoutRowsWidgetsScorecardSparkChartViewSparkChartTypeEnum%s"
            % resource
        )

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return resource
        return dashboard_pb2.MonitoringDashboardRowLayoutRowsWidgetsScorecardSparkChartViewSparkChartTypeEnum.Name(
            resource
        )[
            len(
                "MonitoringDashboardRowLayoutRowsWidgetsScorecardSparkChartViewSparkChartTypeEnum"
            ) :
        ]


class DashboardRowLayoutRowsWidgetsScorecardThresholdsColorEnum(object):
    @classmethod
    def to_proto(self, resource):
        if not resource:
            return resource
        return dashboard_pb2.MonitoringDashboardRowLayoutRowsWidgetsScorecardThresholdsColorEnum.Value(
            "MonitoringDashboardRowLayoutRowsWidgetsScorecardThresholdsColorEnum%s"
            % resource
        )

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return resource
        return dashboard_pb2.MonitoringDashboardRowLayoutRowsWidgetsScorecardThresholdsColorEnum.Name(
            resource
        )[
            len("MonitoringDashboardRowLayoutRowsWidgetsScorecardThresholdsColorEnum") :
        ]


class DashboardRowLayoutRowsWidgetsScorecardThresholdsDirectionEnum(object):
    @classmethod
    def to_proto(self, resource):
        if not resource:
            return resource
        return dashboard_pb2.MonitoringDashboardRowLayoutRowsWidgetsScorecardThresholdsDirectionEnum.Value(
            "MonitoringDashboardRowLayoutRowsWidgetsScorecardThresholdsDirectionEnum%s"
            % resource
        )

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return resource
        return dashboard_pb2.MonitoringDashboardRowLayoutRowsWidgetsScorecardThresholdsDirectionEnum.Name(
            resource
        )[
            len(
                "MonitoringDashboardRowLayoutRowsWidgetsScorecardThresholdsDirectionEnum"
            ) :
        ]


class DashboardRowLayoutRowsWidgetsTextFormatEnum(object):
    @classmethod
    def to_proto(self, resource):
        if not resource:
            return resource
        return (
            dashboard_pb2.MonitoringDashboardRowLayoutRowsWidgetsTextFormatEnum.Value(
                "MonitoringDashboardRowLayoutRowsWidgetsTextFormatEnum%s" % resource
            )
        )

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return resource
        return dashboard_pb2.MonitoringDashboardRowLayoutRowsWidgetsTextFormatEnum.Name(
            resource
        )[len("MonitoringDashboardRowLayoutRowsWidgetsTextFormatEnum") :]


class DashboardColumnLayoutColumnsWidgetsXyChartDataSetsTimeSeriesQueryTimeSeriesFilterAggregationPerSeriesAlignerEnum(
    object
):
    @classmethod
    def to_proto(self, resource):
        if not resource:
            return resource
        return dashboard_pb2.MonitoringDashboardColumnLayoutColumnsWidgetsXyChartDataSetsTimeSeriesQueryTimeSeriesFilterAggregationPerSeriesAlignerEnum.Value(
            "MonitoringDashboardColumnLayoutColumnsWidgetsXyChartDataSetsTimeSeriesQueryTimeSeriesFilterAggregationPerSeriesAlignerEnum%s"
            % resource
        )

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return resource
        return dashboard_pb2.MonitoringDashboardColumnLayoutColumnsWidgetsXyChartDataSetsTimeSeriesQueryTimeSeriesFilterAggregationPerSeriesAlignerEnum.Name(
            resource
        )[
            len(
                "MonitoringDashboardColumnLayoutColumnsWidgetsXyChartDataSetsTimeSeriesQueryTimeSeriesFilterAggregationPerSeriesAlignerEnum"
            ) :
        ]


class DashboardColumnLayoutColumnsWidgetsXyChartDataSetsTimeSeriesQueryTimeSeriesFilterAggregationCrossSeriesReducerEnum(
    object
):
    @classmethod
    def to_proto(self, resource):
        if not resource:
            return resource
        return dashboard_pb2.MonitoringDashboardColumnLayoutColumnsWidgetsXyChartDataSetsTimeSeriesQueryTimeSeriesFilterAggregationCrossSeriesReducerEnum.Value(
            "MonitoringDashboardColumnLayoutColumnsWidgetsXyChartDataSetsTimeSeriesQueryTimeSeriesFilterAggregationCrossSeriesReducerEnum%s"
            % resource
        )

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return resource
        return dashboard_pb2.MonitoringDashboardColumnLayoutColumnsWidgetsXyChartDataSetsTimeSeriesQueryTimeSeriesFilterAggregationCrossSeriesReducerEnum.Name(
            resource
        )[
            len(
                "MonitoringDashboardColumnLayoutColumnsWidgetsXyChartDataSetsTimeSeriesQueryTimeSeriesFilterAggregationCrossSeriesReducerEnum"
            ) :
        ]


class DashboardColumnLayoutColumnsWidgetsXyChartDataSetsTimeSeriesQueryTimeSeriesFilterSecondaryAggregationPerSeriesAlignerEnum(
    object
):
    @classmethod
    def to_proto(self, resource):
        if not resource:
            return resource
        return dashboard_pb2.MonitoringDashboardColumnLayoutColumnsWidgetsXyChartDataSetsTimeSeriesQueryTimeSeriesFilterSecondaryAggregationPerSeriesAlignerEnum.Value(
            "MonitoringDashboardColumnLayoutColumnsWidgetsXyChartDataSetsTimeSeriesQueryTimeSeriesFilterSecondaryAggregationPerSeriesAlignerEnum%s"
            % resource
        )

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return resource
        return dashboard_pb2.MonitoringDashboardColumnLayoutColumnsWidgetsXyChartDataSetsTimeSeriesQueryTimeSeriesFilterSecondaryAggregationPerSeriesAlignerEnum.Name(
            resource
        )[
            len(
                "MonitoringDashboardColumnLayoutColumnsWidgetsXyChartDataSetsTimeSeriesQueryTimeSeriesFilterSecondaryAggregationPerSeriesAlignerEnum"
            ) :
        ]


class DashboardColumnLayoutColumnsWidgetsXyChartDataSetsTimeSeriesQueryTimeSeriesFilterSecondaryAggregationCrossSeriesReducerEnum(
    object
):
    @classmethod
    def to_proto(self, resource):
        if not resource:
            return resource
        return dashboard_pb2.MonitoringDashboardColumnLayoutColumnsWidgetsXyChartDataSetsTimeSeriesQueryTimeSeriesFilterSecondaryAggregationCrossSeriesReducerEnum.Value(
            "MonitoringDashboardColumnLayoutColumnsWidgetsXyChartDataSetsTimeSeriesQueryTimeSeriesFilterSecondaryAggregationCrossSeriesReducerEnum%s"
            % resource
        )

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return resource
        return dashboard_pb2.MonitoringDashboardColumnLayoutColumnsWidgetsXyChartDataSetsTimeSeriesQueryTimeSeriesFilterSecondaryAggregationCrossSeriesReducerEnum.Name(
            resource
        )[
            len(
                "MonitoringDashboardColumnLayoutColumnsWidgetsXyChartDataSetsTimeSeriesQueryTimeSeriesFilterSecondaryAggregationCrossSeriesReducerEnum"
            ) :
        ]


class DashboardColumnLayoutColumnsWidgetsXyChartDataSetsTimeSeriesQueryTimeSeriesFilterPickTimeSeriesFilterRankingMethodEnum(
    object
):
    @classmethod
    def to_proto(self, resource):
        if not resource:
            return resource
        return dashboard_pb2.MonitoringDashboardColumnLayoutColumnsWidgetsXyChartDataSetsTimeSeriesQueryTimeSeriesFilterPickTimeSeriesFilterRankingMethodEnum.Value(
            "MonitoringDashboardColumnLayoutColumnsWidgetsXyChartDataSetsTimeSeriesQueryTimeSeriesFilterPickTimeSeriesFilterRankingMethodEnum%s"
            % resource
        )

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return resource
        return dashboard_pb2.MonitoringDashboardColumnLayoutColumnsWidgetsXyChartDataSetsTimeSeriesQueryTimeSeriesFilterPickTimeSeriesFilterRankingMethodEnum.Name(
            resource
        )[
            len(
                "MonitoringDashboardColumnLayoutColumnsWidgetsXyChartDataSetsTimeSeriesQueryTimeSeriesFilterPickTimeSeriesFilterRankingMethodEnum"
            ) :
        ]


class DashboardColumnLayoutColumnsWidgetsXyChartDataSetsTimeSeriesQueryTimeSeriesFilterPickTimeSeriesFilterDirectionEnum(
    object
):
    @classmethod
    def to_proto(self, resource):
        if not resource:
            return resource
        return dashboard_pb2.MonitoringDashboardColumnLayoutColumnsWidgetsXyChartDataSetsTimeSeriesQueryTimeSeriesFilterPickTimeSeriesFilterDirectionEnum.Value(
            "MonitoringDashboardColumnLayoutColumnsWidgetsXyChartDataSetsTimeSeriesQueryTimeSeriesFilterPickTimeSeriesFilterDirectionEnum%s"
            % resource
        )

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return resource
        return dashboard_pb2.MonitoringDashboardColumnLayoutColumnsWidgetsXyChartDataSetsTimeSeriesQueryTimeSeriesFilterPickTimeSeriesFilterDirectionEnum.Name(
            resource
        )[
            len(
                "MonitoringDashboardColumnLayoutColumnsWidgetsXyChartDataSetsTimeSeriesQueryTimeSeriesFilterPickTimeSeriesFilterDirectionEnum"
            ) :
        ]


class DashboardColumnLayoutColumnsWidgetsXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioNumeratorAggregationPerSeriesAlignerEnum(
    object
):
    @classmethod
    def to_proto(self, resource):
        if not resource:
            return resource
        return dashboard_pb2.MonitoringDashboardColumnLayoutColumnsWidgetsXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioNumeratorAggregationPerSeriesAlignerEnum.Value(
            "MonitoringDashboardColumnLayoutColumnsWidgetsXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioNumeratorAggregationPerSeriesAlignerEnum%s"
            % resource
        )

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return resource
        return dashboard_pb2.MonitoringDashboardColumnLayoutColumnsWidgetsXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioNumeratorAggregationPerSeriesAlignerEnum.Name(
            resource
        )[
            len(
                "MonitoringDashboardColumnLayoutColumnsWidgetsXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioNumeratorAggregationPerSeriesAlignerEnum"
            ) :
        ]


class DashboardColumnLayoutColumnsWidgetsXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioNumeratorAggregationCrossSeriesReducerEnum(
    object
):
    @classmethod
    def to_proto(self, resource):
        if not resource:
            return resource
        return dashboard_pb2.MonitoringDashboardColumnLayoutColumnsWidgetsXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioNumeratorAggregationCrossSeriesReducerEnum.Value(
            "MonitoringDashboardColumnLayoutColumnsWidgetsXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioNumeratorAggregationCrossSeriesReducerEnum%s"
            % resource
        )

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return resource
        return dashboard_pb2.MonitoringDashboardColumnLayoutColumnsWidgetsXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioNumeratorAggregationCrossSeriesReducerEnum.Name(
            resource
        )[
            len(
                "MonitoringDashboardColumnLayoutColumnsWidgetsXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioNumeratorAggregationCrossSeriesReducerEnum"
            ) :
        ]


class DashboardColumnLayoutColumnsWidgetsXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioDenominatorAggregationPerSeriesAlignerEnum(
    object
):
    @classmethod
    def to_proto(self, resource):
        if not resource:
            return resource
        return dashboard_pb2.MonitoringDashboardColumnLayoutColumnsWidgetsXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioDenominatorAggregationPerSeriesAlignerEnum.Value(
            "MonitoringDashboardColumnLayoutColumnsWidgetsXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioDenominatorAggregationPerSeriesAlignerEnum%s"
            % resource
        )

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return resource
        return dashboard_pb2.MonitoringDashboardColumnLayoutColumnsWidgetsXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioDenominatorAggregationPerSeriesAlignerEnum.Name(
            resource
        )[
            len(
                "MonitoringDashboardColumnLayoutColumnsWidgetsXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioDenominatorAggregationPerSeriesAlignerEnum"
            ) :
        ]


class DashboardColumnLayoutColumnsWidgetsXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioDenominatorAggregationCrossSeriesReducerEnum(
    object
):
    @classmethod
    def to_proto(self, resource):
        if not resource:
            return resource
        return dashboard_pb2.MonitoringDashboardColumnLayoutColumnsWidgetsXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioDenominatorAggregationCrossSeriesReducerEnum.Value(
            "MonitoringDashboardColumnLayoutColumnsWidgetsXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioDenominatorAggregationCrossSeriesReducerEnum%s"
            % resource
        )

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return resource
        return dashboard_pb2.MonitoringDashboardColumnLayoutColumnsWidgetsXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioDenominatorAggregationCrossSeriesReducerEnum.Name(
            resource
        )[
            len(
                "MonitoringDashboardColumnLayoutColumnsWidgetsXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioDenominatorAggregationCrossSeriesReducerEnum"
            ) :
        ]


class DashboardColumnLayoutColumnsWidgetsXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioSecondaryAggregationPerSeriesAlignerEnum(
    object
):
    @classmethod
    def to_proto(self, resource):
        if not resource:
            return resource
        return dashboard_pb2.MonitoringDashboardColumnLayoutColumnsWidgetsXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioSecondaryAggregationPerSeriesAlignerEnum.Value(
            "MonitoringDashboardColumnLayoutColumnsWidgetsXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioSecondaryAggregationPerSeriesAlignerEnum%s"
            % resource
        )

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return resource
        return dashboard_pb2.MonitoringDashboardColumnLayoutColumnsWidgetsXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioSecondaryAggregationPerSeriesAlignerEnum.Name(
            resource
        )[
            len(
                "MonitoringDashboardColumnLayoutColumnsWidgetsXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioSecondaryAggregationPerSeriesAlignerEnum"
            ) :
        ]


class DashboardColumnLayoutColumnsWidgetsXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioSecondaryAggregationCrossSeriesReducerEnum(
    object
):
    @classmethod
    def to_proto(self, resource):
        if not resource:
            return resource
        return dashboard_pb2.MonitoringDashboardColumnLayoutColumnsWidgetsXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioSecondaryAggregationCrossSeriesReducerEnum.Value(
            "MonitoringDashboardColumnLayoutColumnsWidgetsXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioSecondaryAggregationCrossSeriesReducerEnum%s"
            % resource
        )

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return resource
        return dashboard_pb2.MonitoringDashboardColumnLayoutColumnsWidgetsXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioSecondaryAggregationCrossSeriesReducerEnum.Name(
            resource
        )[
            len(
                "MonitoringDashboardColumnLayoutColumnsWidgetsXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioSecondaryAggregationCrossSeriesReducerEnum"
            ) :
        ]


class DashboardColumnLayoutColumnsWidgetsXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioPickTimeSeriesFilterRankingMethodEnum(
    object
):
    @classmethod
    def to_proto(self, resource):
        if not resource:
            return resource
        return dashboard_pb2.MonitoringDashboardColumnLayoutColumnsWidgetsXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioPickTimeSeriesFilterRankingMethodEnum.Value(
            "MonitoringDashboardColumnLayoutColumnsWidgetsXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioPickTimeSeriesFilterRankingMethodEnum%s"
            % resource
        )

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return resource
        return dashboard_pb2.MonitoringDashboardColumnLayoutColumnsWidgetsXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioPickTimeSeriesFilterRankingMethodEnum.Name(
            resource
        )[
            len(
                "MonitoringDashboardColumnLayoutColumnsWidgetsXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioPickTimeSeriesFilterRankingMethodEnum"
            ) :
        ]


class DashboardColumnLayoutColumnsWidgetsXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioPickTimeSeriesFilterDirectionEnum(
    object
):
    @classmethod
    def to_proto(self, resource):
        if not resource:
            return resource
        return dashboard_pb2.MonitoringDashboardColumnLayoutColumnsWidgetsXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioPickTimeSeriesFilterDirectionEnum.Value(
            "MonitoringDashboardColumnLayoutColumnsWidgetsXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioPickTimeSeriesFilterDirectionEnum%s"
            % resource
        )

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return resource
        return dashboard_pb2.MonitoringDashboardColumnLayoutColumnsWidgetsXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioPickTimeSeriesFilterDirectionEnum.Name(
            resource
        )[
            len(
                "MonitoringDashboardColumnLayoutColumnsWidgetsXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioPickTimeSeriesFilterDirectionEnum"
            ) :
        ]


class DashboardColumnLayoutColumnsWidgetsXyChartDataSetsPlotTypeEnum(object):
    @classmethod
    def to_proto(self, resource):
        if not resource:
            return resource
        return dashboard_pb2.MonitoringDashboardColumnLayoutColumnsWidgetsXyChartDataSetsPlotTypeEnum.Value(
            "MonitoringDashboardColumnLayoutColumnsWidgetsXyChartDataSetsPlotTypeEnum%s"
            % resource
        )

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return resource
        return dashboard_pb2.MonitoringDashboardColumnLayoutColumnsWidgetsXyChartDataSetsPlotTypeEnum.Name(
            resource
        )[
            len(
                "MonitoringDashboardColumnLayoutColumnsWidgetsXyChartDataSetsPlotTypeEnum"
            ) :
        ]


class DashboardColumnLayoutColumnsWidgetsXyChartThresholdsColorEnum(object):
    @classmethod
    def to_proto(self, resource):
        if not resource:
            return resource
        return dashboard_pb2.MonitoringDashboardColumnLayoutColumnsWidgetsXyChartThresholdsColorEnum.Value(
            "MonitoringDashboardColumnLayoutColumnsWidgetsXyChartThresholdsColorEnum%s"
            % resource
        )

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return resource
        return dashboard_pb2.MonitoringDashboardColumnLayoutColumnsWidgetsXyChartThresholdsColorEnum.Name(
            resource
        )[
            len(
                "MonitoringDashboardColumnLayoutColumnsWidgetsXyChartThresholdsColorEnum"
            ) :
        ]


class DashboardColumnLayoutColumnsWidgetsXyChartThresholdsDirectionEnum(object):
    @classmethod
    def to_proto(self, resource):
        if not resource:
            return resource
        return dashboard_pb2.MonitoringDashboardColumnLayoutColumnsWidgetsXyChartThresholdsDirectionEnum.Value(
            "MonitoringDashboardColumnLayoutColumnsWidgetsXyChartThresholdsDirectionEnum%s"
            % resource
        )

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return resource
        return dashboard_pb2.MonitoringDashboardColumnLayoutColumnsWidgetsXyChartThresholdsDirectionEnum.Name(
            resource
        )[
            len(
                "MonitoringDashboardColumnLayoutColumnsWidgetsXyChartThresholdsDirectionEnum"
            ) :
        ]


class DashboardColumnLayoutColumnsWidgetsXyChartXAxisScaleEnum(object):
    @classmethod
    def to_proto(self, resource):
        if not resource:
            return resource
        return dashboard_pb2.MonitoringDashboardColumnLayoutColumnsWidgetsXyChartXAxisScaleEnum.Value(
            "MonitoringDashboardColumnLayoutColumnsWidgetsXyChartXAxisScaleEnum%s"
            % resource
        )

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return resource
        return dashboard_pb2.MonitoringDashboardColumnLayoutColumnsWidgetsXyChartXAxisScaleEnum.Name(
            resource
        )[
            len("MonitoringDashboardColumnLayoutColumnsWidgetsXyChartXAxisScaleEnum") :
        ]


class DashboardColumnLayoutColumnsWidgetsXyChartYAxisScaleEnum(object):
    @classmethod
    def to_proto(self, resource):
        if not resource:
            return resource
        return dashboard_pb2.MonitoringDashboardColumnLayoutColumnsWidgetsXyChartYAxisScaleEnum.Value(
            "MonitoringDashboardColumnLayoutColumnsWidgetsXyChartYAxisScaleEnum%s"
            % resource
        )

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return resource
        return dashboard_pb2.MonitoringDashboardColumnLayoutColumnsWidgetsXyChartYAxisScaleEnum.Name(
            resource
        )[
            len("MonitoringDashboardColumnLayoutColumnsWidgetsXyChartYAxisScaleEnum") :
        ]


class DashboardColumnLayoutColumnsWidgetsXyChartChartOptionsModeEnum(object):
    @classmethod
    def to_proto(self, resource):
        if not resource:
            return resource
        return dashboard_pb2.MonitoringDashboardColumnLayoutColumnsWidgetsXyChartChartOptionsModeEnum.Value(
            "MonitoringDashboardColumnLayoutColumnsWidgetsXyChartChartOptionsModeEnum%s"
            % resource
        )

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return resource
        return dashboard_pb2.MonitoringDashboardColumnLayoutColumnsWidgetsXyChartChartOptionsModeEnum.Name(
            resource
        )[
            len(
                "MonitoringDashboardColumnLayoutColumnsWidgetsXyChartChartOptionsModeEnum"
            ) :
        ]


class DashboardColumnLayoutColumnsWidgetsScorecardTimeSeriesQueryTimeSeriesFilterAggregationPerSeriesAlignerEnum(
    object
):
    @classmethod
    def to_proto(self, resource):
        if not resource:
            return resource
        return dashboard_pb2.MonitoringDashboardColumnLayoutColumnsWidgetsScorecardTimeSeriesQueryTimeSeriesFilterAggregationPerSeriesAlignerEnum.Value(
            "MonitoringDashboardColumnLayoutColumnsWidgetsScorecardTimeSeriesQueryTimeSeriesFilterAggregationPerSeriesAlignerEnum%s"
            % resource
        )

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return resource
        return dashboard_pb2.MonitoringDashboardColumnLayoutColumnsWidgetsScorecardTimeSeriesQueryTimeSeriesFilterAggregationPerSeriesAlignerEnum.Name(
            resource
        )[
            len(
                "MonitoringDashboardColumnLayoutColumnsWidgetsScorecardTimeSeriesQueryTimeSeriesFilterAggregationPerSeriesAlignerEnum"
            ) :
        ]


class DashboardColumnLayoutColumnsWidgetsScorecardTimeSeriesQueryTimeSeriesFilterAggregationCrossSeriesReducerEnum(
    object
):
    @classmethod
    def to_proto(self, resource):
        if not resource:
            return resource
        return dashboard_pb2.MonitoringDashboardColumnLayoutColumnsWidgetsScorecardTimeSeriesQueryTimeSeriesFilterAggregationCrossSeriesReducerEnum.Value(
            "MonitoringDashboardColumnLayoutColumnsWidgetsScorecardTimeSeriesQueryTimeSeriesFilterAggregationCrossSeriesReducerEnum%s"
            % resource
        )

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return resource
        return dashboard_pb2.MonitoringDashboardColumnLayoutColumnsWidgetsScorecardTimeSeriesQueryTimeSeriesFilterAggregationCrossSeriesReducerEnum.Name(
            resource
        )[
            len(
                "MonitoringDashboardColumnLayoutColumnsWidgetsScorecardTimeSeriesQueryTimeSeriesFilterAggregationCrossSeriesReducerEnum"
            ) :
        ]


class DashboardColumnLayoutColumnsWidgetsScorecardTimeSeriesQueryTimeSeriesFilterSecondaryAggregationPerSeriesAlignerEnum(
    object
):
    @classmethod
    def to_proto(self, resource):
        if not resource:
            return resource
        return dashboard_pb2.MonitoringDashboardColumnLayoutColumnsWidgetsScorecardTimeSeriesQueryTimeSeriesFilterSecondaryAggregationPerSeriesAlignerEnum.Value(
            "MonitoringDashboardColumnLayoutColumnsWidgetsScorecardTimeSeriesQueryTimeSeriesFilterSecondaryAggregationPerSeriesAlignerEnum%s"
            % resource
        )

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return resource
        return dashboard_pb2.MonitoringDashboardColumnLayoutColumnsWidgetsScorecardTimeSeriesQueryTimeSeriesFilterSecondaryAggregationPerSeriesAlignerEnum.Name(
            resource
        )[
            len(
                "MonitoringDashboardColumnLayoutColumnsWidgetsScorecardTimeSeriesQueryTimeSeriesFilterSecondaryAggregationPerSeriesAlignerEnum"
            ) :
        ]


class DashboardColumnLayoutColumnsWidgetsScorecardTimeSeriesQueryTimeSeriesFilterSecondaryAggregationCrossSeriesReducerEnum(
    object
):
    @classmethod
    def to_proto(self, resource):
        if not resource:
            return resource
        return dashboard_pb2.MonitoringDashboardColumnLayoutColumnsWidgetsScorecardTimeSeriesQueryTimeSeriesFilterSecondaryAggregationCrossSeriesReducerEnum.Value(
            "MonitoringDashboardColumnLayoutColumnsWidgetsScorecardTimeSeriesQueryTimeSeriesFilterSecondaryAggregationCrossSeriesReducerEnum%s"
            % resource
        )

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return resource
        return dashboard_pb2.MonitoringDashboardColumnLayoutColumnsWidgetsScorecardTimeSeriesQueryTimeSeriesFilterSecondaryAggregationCrossSeriesReducerEnum.Name(
            resource
        )[
            len(
                "MonitoringDashboardColumnLayoutColumnsWidgetsScorecardTimeSeriesQueryTimeSeriesFilterSecondaryAggregationCrossSeriesReducerEnum"
            ) :
        ]


class DashboardColumnLayoutColumnsWidgetsScorecardTimeSeriesQueryTimeSeriesFilterPickTimeSeriesFilterRankingMethodEnum(
    object
):
    @classmethod
    def to_proto(self, resource):
        if not resource:
            return resource
        return dashboard_pb2.MonitoringDashboardColumnLayoutColumnsWidgetsScorecardTimeSeriesQueryTimeSeriesFilterPickTimeSeriesFilterRankingMethodEnum.Value(
            "MonitoringDashboardColumnLayoutColumnsWidgetsScorecardTimeSeriesQueryTimeSeriesFilterPickTimeSeriesFilterRankingMethodEnum%s"
            % resource
        )

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return resource
        return dashboard_pb2.MonitoringDashboardColumnLayoutColumnsWidgetsScorecardTimeSeriesQueryTimeSeriesFilterPickTimeSeriesFilterRankingMethodEnum.Name(
            resource
        )[
            len(
                "MonitoringDashboardColumnLayoutColumnsWidgetsScorecardTimeSeriesQueryTimeSeriesFilterPickTimeSeriesFilterRankingMethodEnum"
            ) :
        ]


class DashboardColumnLayoutColumnsWidgetsScorecardTimeSeriesQueryTimeSeriesFilterPickTimeSeriesFilterDirectionEnum(
    object
):
    @classmethod
    def to_proto(self, resource):
        if not resource:
            return resource
        return dashboard_pb2.MonitoringDashboardColumnLayoutColumnsWidgetsScorecardTimeSeriesQueryTimeSeriesFilterPickTimeSeriesFilterDirectionEnum.Value(
            "MonitoringDashboardColumnLayoutColumnsWidgetsScorecardTimeSeriesQueryTimeSeriesFilterPickTimeSeriesFilterDirectionEnum%s"
            % resource
        )

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return resource
        return dashboard_pb2.MonitoringDashboardColumnLayoutColumnsWidgetsScorecardTimeSeriesQueryTimeSeriesFilterPickTimeSeriesFilterDirectionEnum.Name(
            resource
        )[
            len(
                "MonitoringDashboardColumnLayoutColumnsWidgetsScorecardTimeSeriesQueryTimeSeriesFilterPickTimeSeriesFilterDirectionEnum"
            ) :
        ]


class DashboardColumnLayoutColumnsWidgetsScorecardTimeSeriesQueryTimeSeriesFilterRatioNumeratorAggregationPerSeriesAlignerEnum(
    object
):
    @classmethod
    def to_proto(self, resource):
        if not resource:
            return resource
        return dashboard_pb2.MonitoringDashboardColumnLayoutColumnsWidgetsScorecardTimeSeriesQueryTimeSeriesFilterRatioNumeratorAggregationPerSeriesAlignerEnum.Value(
            "MonitoringDashboardColumnLayoutColumnsWidgetsScorecardTimeSeriesQueryTimeSeriesFilterRatioNumeratorAggregationPerSeriesAlignerEnum%s"
            % resource
        )

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return resource
        return dashboard_pb2.MonitoringDashboardColumnLayoutColumnsWidgetsScorecardTimeSeriesQueryTimeSeriesFilterRatioNumeratorAggregationPerSeriesAlignerEnum.Name(
            resource
        )[
            len(
                "MonitoringDashboardColumnLayoutColumnsWidgetsScorecardTimeSeriesQueryTimeSeriesFilterRatioNumeratorAggregationPerSeriesAlignerEnum"
            ) :
        ]


class DashboardColumnLayoutColumnsWidgetsScorecardTimeSeriesQueryTimeSeriesFilterRatioNumeratorAggregationCrossSeriesReducerEnum(
    object
):
    @classmethod
    def to_proto(self, resource):
        if not resource:
            return resource
        return dashboard_pb2.MonitoringDashboardColumnLayoutColumnsWidgetsScorecardTimeSeriesQueryTimeSeriesFilterRatioNumeratorAggregationCrossSeriesReducerEnum.Value(
            "MonitoringDashboardColumnLayoutColumnsWidgetsScorecardTimeSeriesQueryTimeSeriesFilterRatioNumeratorAggregationCrossSeriesReducerEnum%s"
            % resource
        )

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return resource
        return dashboard_pb2.MonitoringDashboardColumnLayoutColumnsWidgetsScorecardTimeSeriesQueryTimeSeriesFilterRatioNumeratorAggregationCrossSeriesReducerEnum.Name(
            resource
        )[
            len(
                "MonitoringDashboardColumnLayoutColumnsWidgetsScorecardTimeSeriesQueryTimeSeriesFilterRatioNumeratorAggregationCrossSeriesReducerEnum"
            ) :
        ]


class DashboardColumnLayoutColumnsWidgetsScorecardTimeSeriesQueryTimeSeriesFilterRatioDenominatorAggregationPerSeriesAlignerEnum(
    object
):
    @classmethod
    def to_proto(self, resource):
        if not resource:
            return resource
        return dashboard_pb2.MonitoringDashboardColumnLayoutColumnsWidgetsScorecardTimeSeriesQueryTimeSeriesFilterRatioDenominatorAggregationPerSeriesAlignerEnum.Value(
            "MonitoringDashboardColumnLayoutColumnsWidgetsScorecardTimeSeriesQueryTimeSeriesFilterRatioDenominatorAggregationPerSeriesAlignerEnum%s"
            % resource
        )

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return resource
        return dashboard_pb2.MonitoringDashboardColumnLayoutColumnsWidgetsScorecardTimeSeriesQueryTimeSeriesFilterRatioDenominatorAggregationPerSeriesAlignerEnum.Name(
            resource
        )[
            len(
                "MonitoringDashboardColumnLayoutColumnsWidgetsScorecardTimeSeriesQueryTimeSeriesFilterRatioDenominatorAggregationPerSeriesAlignerEnum"
            ) :
        ]


class DashboardColumnLayoutColumnsWidgetsScorecardTimeSeriesQueryTimeSeriesFilterRatioDenominatorAggregationCrossSeriesReducerEnum(
    object
):
    @classmethod
    def to_proto(self, resource):
        if not resource:
            return resource
        return dashboard_pb2.MonitoringDashboardColumnLayoutColumnsWidgetsScorecardTimeSeriesQueryTimeSeriesFilterRatioDenominatorAggregationCrossSeriesReducerEnum.Value(
            "MonitoringDashboardColumnLayoutColumnsWidgetsScorecardTimeSeriesQueryTimeSeriesFilterRatioDenominatorAggregationCrossSeriesReducerEnum%s"
            % resource
        )

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return resource
        return dashboard_pb2.MonitoringDashboardColumnLayoutColumnsWidgetsScorecardTimeSeriesQueryTimeSeriesFilterRatioDenominatorAggregationCrossSeriesReducerEnum.Name(
            resource
        )[
            len(
                "MonitoringDashboardColumnLayoutColumnsWidgetsScorecardTimeSeriesQueryTimeSeriesFilterRatioDenominatorAggregationCrossSeriesReducerEnum"
            ) :
        ]


class DashboardColumnLayoutColumnsWidgetsScorecardTimeSeriesQueryTimeSeriesFilterRatioSecondaryAggregationPerSeriesAlignerEnum(
    object
):
    @classmethod
    def to_proto(self, resource):
        if not resource:
            return resource
        return dashboard_pb2.MonitoringDashboardColumnLayoutColumnsWidgetsScorecardTimeSeriesQueryTimeSeriesFilterRatioSecondaryAggregationPerSeriesAlignerEnum.Value(
            "MonitoringDashboardColumnLayoutColumnsWidgetsScorecardTimeSeriesQueryTimeSeriesFilterRatioSecondaryAggregationPerSeriesAlignerEnum%s"
            % resource
        )

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return resource
        return dashboard_pb2.MonitoringDashboardColumnLayoutColumnsWidgetsScorecardTimeSeriesQueryTimeSeriesFilterRatioSecondaryAggregationPerSeriesAlignerEnum.Name(
            resource
        )[
            len(
                "MonitoringDashboardColumnLayoutColumnsWidgetsScorecardTimeSeriesQueryTimeSeriesFilterRatioSecondaryAggregationPerSeriesAlignerEnum"
            ) :
        ]


class DashboardColumnLayoutColumnsWidgetsScorecardTimeSeriesQueryTimeSeriesFilterRatioSecondaryAggregationCrossSeriesReducerEnum(
    object
):
    @classmethod
    def to_proto(self, resource):
        if not resource:
            return resource
        return dashboard_pb2.MonitoringDashboardColumnLayoutColumnsWidgetsScorecardTimeSeriesQueryTimeSeriesFilterRatioSecondaryAggregationCrossSeriesReducerEnum.Value(
            "MonitoringDashboardColumnLayoutColumnsWidgetsScorecardTimeSeriesQueryTimeSeriesFilterRatioSecondaryAggregationCrossSeriesReducerEnum%s"
            % resource
        )

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return resource
        return dashboard_pb2.MonitoringDashboardColumnLayoutColumnsWidgetsScorecardTimeSeriesQueryTimeSeriesFilterRatioSecondaryAggregationCrossSeriesReducerEnum.Name(
            resource
        )[
            len(
                "MonitoringDashboardColumnLayoutColumnsWidgetsScorecardTimeSeriesQueryTimeSeriesFilterRatioSecondaryAggregationCrossSeriesReducerEnum"
            ) :
        ]


class DashboardColumnLayoutColumnsWidgetsScorecardTimeSeriesQueryTimeSeriesFilterRatioPickTimeSeriesFilterRankingMethodEnum(
    object
):
    @classmethod
    def to_proto(self, resource):
        if not resource:
            return resource
        return dashboard_pb2.MonitoringDashboardColumnLayoutColumnsWidgetsScorecardTimeSeriesQueryTimeSeriesFilterRatioPickTimeSeriesFilterRankingMethodEnum.Value(
            "MonitoringDashboardColumnLayoutColumnsWidgetsScorecardTimeSeriesQueryTimeSeriesFilterRatioPickTimeSeriesFilterRankingMethodEnum%s"
            % resource
        )

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return resource
        return dashboard_pb2.MonitoringDashboardColumnLayoutColumnsWidgetsScorecardTimeSeriesQueryTimeSeriesFilterRatioPickTimeSeriesFilterRankingMethodEnum.Name(
            resource
        )[
            len(
                "MonitoringDashboardColumnLayoutColumnsWidgetsScorecardTimeSeriesQueryTimeSeriesFilterRatioPickTimeSeriesFilterRankingMethodEnum"
            ) :
        ]


class DashboardColumnLayoutColumnsWidgetsScorecardTimeSeriesQueryTimeSeriesFilterRatioPickTimeSeriesFilterDirectionEnum(
    object
):
    @classmethod
    def to_proto(self, resource):
        if not resource:
            return resource
        return dashboard_pb2.MonitoringDashboardColumnLayoutColumnsWidgetsScorecardTimeSeriesQueryTimeSeriesFilterRatioPickTimeSeriesFilterDirectionEnum.Value(
            "MonitoringDashboardColumnLayoutColumnsWidgetsScorecardTimeSeriesQueryTimeSeriesFilterRatioPickTimeSeriesFilterDirectionEnum%s"
            % resource
        )

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return resource
        return dashboard_pb2.MonitoringDashboardColumnLayoutColumnsWidgetsScorecardTimeSeriesQueryTimeSeriesFilterRatioPickTimeSeriesFilterDirectionEnum.Name(
            resource
        )[
            len(
                "MonitoringDashboardColumnLayoutColumnsWidgetsScorecardTimeSeriesQueryTimeSeriesFilterRatioPickTimeSeriesFilterDirectionEnum"
            ) :
        ]


class DashboardColumnLayoutColumnsWidgetsScorecardSparkChartViewSparkChartTypeEnum(
    object
):
    @classmethod
    def to_proto(self, resource):
        if not resource:
            return resource
        return dashboard_pb2.MonitoringDashboardColumnLayoutColumnsWidgetsScorecardSparkChartViewSparkChartTypeEnum.Value(
            "MonitoringDashboardColumnLayoutColumnsWidgetsScorecardSparkChartViewSparkChartTypeEnum%s"
            % resource
        )

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return resource
        return dashboard_pb2.MonitoringDashboardColumnLayoutColumnsWidgetsScorecardSparkChartViewSparkChartTypeEnum.Name(
            resource
        )[
            len(
                "MonitoringDashboardColumnLayoutColumnsWidgetsScorecardSparkChartViewSparkChartTypeEnum"
            ) :
        ]


class DashboardColumnLayoutColumnsWidgetsScorecardThresholdsColorEnum(object):
    @classmethod
    def to_proto(self, resource):
        if not resource:
            return resource
        return dashboard_pb2.MonitoringDashboardColumnLayoutColumnsWidgetsScorecardThresholdsColorEnum.Value(
            "MonitoringDashboardColumnLayoutColumnsWidgetsScorecardThresholdsColorEnum%s"
            % resource
        )

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return resource
        return dashboard_pb2.MonitoringDashboardColumnLayoutColumnsWidgetsScorecardThresholdsColorEnum.Name(
            resource
        )[
            len(
                "MonitoringDashboardColumnLayoutColumnsWidgetsScorecardThresholdsColorEnum"
            ) :
        ]


class DashboardColumnLayoutColumnsWidgetsScorecardThresholdsDirectionEnum(object):
    @classmethod
    def to_proto(self, resource):
        if not resource:
            return resource
        return dashboard_pb2.MonitoringDashboardColumnLayoutColumnsWidgetsScorecardThresholdsDirectionEnum.Value(
            "MonitoringDashboardColumnLayoutColumnsWidgetsScorecardThresholdsDirectionEnum%s"
            % resource
        )

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return resource
        return dashboard_pb2.MonitoringDashboardColumnLayoutColumnsWidgetsScorecardThresholdsDirectionEnum.Name(
            resource
        )[
            len(
                "MonitoringDashboardColumnLayoutColumnsWidgetsScorecardThresholdsDirectionEnum"
            ) :
        ]


class DashboardColumnLayoutColumnsWidgetsTextFormatEnum(object):
    @classmethod
    def to_proto(self, resource):
        if not resource:
            return resource
        return dashboard_pb2.MonitoringDashboardColumnLayoutColumnsWidgetsTextFormatEnum.Value(
            "MonitoringDashboardColumnLayoutColumnsWidgetsTextFormatEnum%s" % resource
        )

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return resource
        return dashboard_pb2.MonitoringDashboardColumnLayoutColumnsWidgetsTextFormatEnum.Name(
            resource
        )[
            len("MonitoringDashboardColumnLayoutColumnsWidgetsTextFormatEnum") :
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
