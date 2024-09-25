// SPDX-FileCopyrightText: 2023 The Crossplane Authors <https://crossplane.io>
//
// SPDX-License-Identifier: Apache-2.0

/*
Copyright 2022 Upbound Inc.
*/

// Code generated by upjet. DO NOT EDIT.

package v1alpha1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime/schema"

	v1 "github.com/crossplane/crossplane-runtime/apis/common/v1"
)

type FormulaInitParameters struct {

	// (String) The formula string, which is an expression involving named queries.
	// The formula string, which is an expression involving named queries.
	FormulaExpression *string `json:"formulaExpression,omitempty" tf:"formula_expression,omitempty"`
}

type FormulaObservation struct {

	// (String) The formula string, which is an expression involving named queries.
	// The formula string, which is an expression involving named queries.
	FormulaExpression *string `json:"formulaExpression,omitempty" tf:"formula_expression,omitempty"`
}

type FormulaParameters struct {

	// (String) The formula string, which is an expression involving named queries.
	// The formula string, which is an expression involving named queries.
	// +kubebuilder:validation:Optional
	FormulaExpression *string `json:"formulaExpression" tf:"formula_expression,omitempty"`
}

type MetricQueryInitParameters struct {

	// (String) The data source for metrics queries. Defaults to "metrics".
	// The data source for metrics queries. Defaults to `"metrics"`.
	DataSource *string `json:"dataSource,omitempty" tf:"data_source,omitempty"`

	// (String) Name of Datadog service level objective
	// The name of the query for use in formulas.
	Name *string `json:"name,omitempty" tf:"name,omitempty"`

	// (Block List, Max: 1) The metric query of good / total events (see below for nested schema)
	// The metrics query definition.
	Query *string `json:"query,omitempty" tf:"query,omitempty"`
}

type MetricQueryObservation struct {

	// (String) The data source for metrics queries. Defaults to "metrics".
	// The data source for metrics queries. Defaults to `"metrics"`.
	DataSource *string `json:"dataSource,omitempty" tf:"data_source,omitempty"`

	// (String) Name of Datadog service level objective
	// The name of the query for use in formulas.
	Name *string `json:"name,omitempty" tf:"name,omitempty"`

	// (Block List, Max: 1) The metric query of good / total events (see below for nested schema)
	// The metrics query definition.
	Query *string `json:"query,omitempty" tf:"query,omitempty"`
}

type MetricQueryParameters struct {

	// (String) The data source for metrics queries. Defaults to "metrics".
	// The data source for metrics queries. Defaults to `"metrics"`.
	// +kubebuilder:validation:Optional
	DataSource *string `json:"dataSource,omitempty" tf:"data_source,omitempty"`

	// (String) Name of Datadog service level objective
	// The name of the query for use in formulas.
	// +kubebuilder:validation:Optional
	Name *string `json:"name" tf:"name,omitempty"`

	// (Block List, Max: 1) The metric query of good / total events (see below for nested schema)
	// The metrics query definition.
	// +kubebuilder:validation:Optional
	Query *string `json:"query" tf:"query,omitempty"`
}

type QueryInitParameters struct {

	// (String) The sum of the total events.
	// The sum of the `total` events.
	Denominator *string `json:"denominator,omitempty" tf:"denominator,omitempty"`

	// (String) The sum of all the good events.
	// The sum of all the `good` events.
	Numerator *string `json:"numerator,omitempty" tf:"numerator,omitempty"`
}

type QueryObservation struct {

	// (String) The sum of the total events.
	// The sum of the `total` events.
	Denominator *string `json:"denominator,omitempty" tf:"denominator,omitempty"`

	// (String) The sum of all the good events.
	// The sum of all the `good` events.
	Numerator *string `json:"numerator,omitempty" tf:"numerator,omitempty"`
}

type QueryParameters struct {

	// (String) The sum of the total events.
	// The sum of the `total` events.
	// +kubebuilder:validation:Optional
	Denominator *string `json:"denominator" tf:"denominator,omitempty"`

	// (String) The sum of all the good events.
	// The sum of all the `good` events.
	// +kubebuilder:validation:Optional
	Numerator *string `json:"numerator" tf:"numerator,omitempty"`
}

type QueryQueryInitParameters struct {

	// (Block List, Max: 1) A timeseries formula and functions metrics query. (see below for nested schema)
	// A timeseries formula and functions metrics query.
	MetricQuery []MetricQueryInitParameters `json:"metricQuery,omitempty" tf:"metric_query,omitempty"`
}

type QueryQueryObservation struct {

	// (Block List, Max: 1) A timeseries formula and functions metrics query. (see below for nested schema)
	// A timeseries formula and functions metrics query.
	MetricQuery []MetricQueryObservation `json:"metricQuery,omitempty" tf:"metric_query,omitempty"`
}

type QueryQueryParameters struct {

	// (Block List, Max: 1) A timeseries formula and functions metrics query. (see below for nested schema)
	// A timeseries formula and functions metrics query.
	// +kubebuilder:validation:Optional
	MetricQuery []MetricQueryParameters `json:"metricQuery,omitempty" tf:"metric_query,omitempty"`
}

type ServiceLevelObjectiveInitParameters struct {

	// (String) A description of this service level objective.
	// A description of this service level objective.
	Description *string `json:"description,omitempty" tf:"description,omitempty"`

	// (Boolean) A boolean indicating whether this monitor can be deleted even if it's referenced by other resources (for example, dashboards).
	// A boolean indicating whether this monitor can be deleted even if it's referenced by other resources (for example, dashboards).
	ForceDelete *bool `json:"forceDelete,omitempty" tf:"force_delete,omitempty"`

	// based SLOs
	// A static set of groups to filter monitor-based SLOs
	// +listType=set
	Groups []*string `json:"groups,omitempty" tf:"groups,omitempty"`

	// (Set of Number) A static set of monitor IDs to use as part of the SLO
	// A static set of monitor IDs to use as part of the SLO
	// +listType=set
	MonitorIds []*float64 `json:"monitorIds,omitempty" tf:"monitor_ids,omitempty"`

	// (String) Name of Datadog service level objective
	// Name of Datadog service level objective
	Name *string `json:"name,omitempty" tf:"name,omitempty"`

	// (Block List, Max: 1) The metric query of good / total events (see below for nested schema)
	// The metric query of good / total events
	Query []QueryInitParameters `json:"query,omitempty" tf:"query,omitempty"`

	// (Block List, Max: 1) A map of SLI specifications to use as part of the SLO. (see below for nested schema)
	// A map of SLI specifications to use as part of the SLO.
	SliSpecification []SliSpecificationInitParameters `json:"sliSpecification,omitempty" tf:"sli_specification,omitempty"`

	// (Set of String) A list of tags to associate with your service level objective. This can help you categorize and filter service level objectives in the service level objectives page of the UI. Note: it's not currently possible to filter by these tags when querying via the API
	// A list of tags to associate with your service level objective. This can help you categorize and filter service level objectives in the service level objectives page of the UI. Note: it's not currently possible to filter by these tags when querying via the API
	// +listType=set
	Tags []*string `json:"tags,omitempty" tf:"tags,omitempty"`

	// (Number) The objective's target in (0,100). This must match the corresponding thresholds of the primary time frame.
	// The objective's target in `(0,100)`. This must match the corresponding thresholds of the primary time frame.
	TargetThreshold *float64 `json:"targetThreshold,omitempty" tf:"target_threshold,omitempty"`

	// (Block List, Min: 1) A list of thresholds and targets that define the service level objectives from the provided SLIs. (see below for nested schema)
	// A list of thresholds and targets that define the service level objectives from the provided SLIs.
	Thresholds []ThresholdsInitParameters `json:"thresholds,omitempty" tf:"thresholds,omitempty"`

	// (String) The primary time frame for the objective. The mapping from these types to the types found in the Datadog Web UI can be found in the Datadog API documentation page. Valid values are 7d, 30d, 90d, custom.
	// The primary time frame for the objective. The mapping from these types to the types found in the Datadog Web UI can be found in the Datadog API documentation page. Valid values are `7d`, `30d`, `90d`, `custom`.
	Timeframe *string `json:"timeframe,omitempty" tf:"timeframe,omitempty"`

	// (String) The type of the service level objective. The mapping from these types to the types found in the Datadog Web UI can be found in the Datadog API documentation page. Valid values are metric, monitor, time_slice.
	// The type of the service level objective. The mapping from these types to the types found in the Datadog Web UI can be found in the Datadog API [documentation page](https://docs.datadoghq.com/api/v1/service-level-objectives/#create-a-slo-object). Valid values are `metric`, `monitor`, `time_slice`.
	Type *string `json:"type,omitempty" tf:"type,omitempty"`

	// (Boolean) Whether or not to validate the SLO. It checks if monitors added to a monitor SLO already exist.
	// Whether or not to validate the SLO. It checks if monitors added to a monitor SLO already exist.
	Validate *bool `json:"validate,omitempty" tf:"validate,omitempty"`

	// (Number) The objective's warning value in (0,100). This must be greater than the target value and match the corresponding thresholds of the primary time frame.
	// The objective's warning value in `(0,100)`. This must be greater than the target value and match the corresponding thresholds of the primary time frame.
	WarningThreshold *float64 `json:"warningThreshold,omitempty" tf:"warning_threshold,omitempty"`
}

type ServiceLevelObjectiveObservation struct {

	// (String) A description of this service level objective.
	// A description of this service level objective.
	Description *string `json:"description,omitempty" tf:"description,omitempty"`

	// (Boolean) A boolean indicating whether this monitor can be deleted even if it's referenced by other resources (for example, dashboards).
	// A boolean indicating whether this monitor can be deleted even if it's referenced by other resources (for example, dashboards).
	ForceDelete *bool `json:"forceDelete,omitempty" tf:"force_delete,omitempty"`

	// based SLOs
	// A static set of groups to filter monitor-based SLOs
	// +listType=set
	Groups []*string `json:"groups,omitempty" tf:"groups,omitempty"`

	// (String) The ID of this resource.
	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	// (Set of Number) A static set of monitor IDs to use as part of the SLO
	// A static set of monitor IDs to use as part of the SLO
	// +listType=set
	MonitorIds []*float64 `json:"monitorIds,omitempty" tf:"monitor_ids,omitempty"`

	// (String) Name of Datadog service level objective
	// Name of Datadog service level objective
	Name *string `json:"name,omitempty" tf:"name,omitempty"`

	// (Block List, Max: 1) The metric query of good / total events (see below for nested schema)
	// The metric query of good / total events
	Query []QueryObservation `json:"query,omitempty" tf:"query,omitempty"`

	// (Block List, Max: 1) A map of SLI specifications to use as part of the SLO. (see below for nested schema)
	// A map of SLI specifications to use as part of the SLO.
	SliSpecification []SliSpecificationObservation `json:"sliSpecification,omitempty" tf:"sli_specification,omitempty"`

	// (Set of String) A list of tags to associate with your service level objective. This can help you categorize and filter service level objectives in the service level objectives page of the UI. Note: it's not currently possible to filter by these tags when querying via the API
	// A list of tags to associate with your service level objective. This can help you categorize and filter service level objectives in the service level objectives page of the UI. Note: it's not currently possible to filter by these tags when querying via the API
	// +listType=set
	Tags []*string `json:"tags,omitempty" tf:"tags,omitempty"`

	// (Number) The objective's target in (0,100). This must match the corresponding thresholds of the primary time frame.
	// The objective's target in `(0,100)`. This must match the corresponding thresholds of the primary time frame.
	TargetThreshold *float64 `json:"targetThreshold,omitempty" tf:"target_threshold,omitempty"`

	// (Block List, Min: 1) A list of thresholds and targets that define the service level objectives from the provided SLIs. (see below for nested schema)
	// A list of thresholds and targets that define the service level objectives from the provided SLIs.
	Thresholds []ThresholdsObservation `json:"thresholds,omitempty" tf:"thresholds,omitempty"`

	// (String) The primary time frame for the objective. The mapping from these types to the types found in the Datadog Web UI can be found in the Datadog API documentation page. Valid values are 7d, 30d, 90d, custom.
	// The primary time frame for the objective. The mapping from these types to the types found in the Datadog Web UI can be found in the Datadog API documentation page. Valid values are `7d`, `30d`, `90d`, `custom`.
	Timeframe *string `json:"timeframe,omitempty" tf:"timeframe,omitempty"`

	// (String) The type of the service level objective. The mapping from these types to the types found in the Datadog Web UI can be found in the Datadog API documentation page. Valid values are metric, monitor, time_slice.
	// The type of the service level objective. The mapping from these types to the types found in the Datadog Web UI can be found in the Datadog API [documentation page](https://docs.datadoghq.com/api/v1/service-level-objectives/#create-a-slo-object). Valid values are `metric`, `monitor`, `time_slice`.
	Type *string `json:"type,omitempty" tf:"type,omitempty"`

	// (Boolean) Whether or not to validate the SLO. It checks if monitors added to a monitor SLO already exist.
	// Whether or not to validate the SLO. It checks if monitors added to a monitor SLO already exist.
	Validate *bool `json:"validate,omitempty" tf:"validate,omitempty"`

	// (Number) The objective's warning value in (0,100). This must be greater than the target value and match the corresponding thresholds of the primary time frame.
	// The objective's warning value in `(0,100)`. This must be greater than the target value and match the corresponding thresholds of the primary time frame.
	WarningThreshold *float64 `json:"warningThreshold,omitempty" tf:"warning_threshold,omitempty"`
}

type ServiceLevelObjectiveParameters struct {

	// (String) A description of this service level objective.
	// A description of this service level objective.
	// +kubebuilder:validation:Optional
	Description *string `json:"description,omitempty" tf:"description,omitempty"`

	// (Boolean) A boolean indicating whether this monitor can be deleted even if it's referenced by other resources (for example, dashboards).
	// A boolean indicating whether this monitor can be deleted even if it's referenced by other resources (for example, dashboards).
	// +kubebuilder:validation:Optional
	ForceDelete *bool `json:"forceDelete,omitempty" tf:"force_delete,omitempty"`

	// based SLOs
	// A static set of groups to filter monitor-based SLOs
	// +kubebuilder:validation:Optional
	// +listType=set
	Groups []*string `json:"groups,omitempty" tf:"groups,omitempty"`

	// (Set of Number) A static set of monitor IDs to use as part of the SLO
	// A static set of monitor IDs to use as part of the SLO
	// +kubebuilder:validation:Optional
	// +listType=set
	MonitorIds []*float64 `json:"monitorIds,omitempty" tf:"monitor_ids,omitempty"`

	// (String) Name of Datadog service level objective
	// Name of Datadog service level objective
	// +kubebuilder:validation:Optional
	Name *string `json:"name,omitempty" tf:"name,omitempty"`

	// (Block List, Max: 1) The metric query of good / total events (see below for nested schema)
	// The metric query of good / total events
	// +kubebuilder:validation:Optional
	Query []QueryParameters `json:"query,omitempty" tf:"query,omitempty"`

	// (Block List, Max: 1) A map of SLI specifications to use as part of the SLO. (see below for nested schema)
	// A map of SLI specifications to use as part of the SLO.
	// +kubebuilder:validation:Optional
	SliSpecification []SliSpecificationParameters `json:"sliSpecification,omitempty" tf:"sli_specification,omitempty"`

	// (Set of String) A list of tags to associate with your service level objective. This can help you categorize and filter service level objectives in the service level objectives page of the UI. Note: it's not currently possible to filter by these tags when querying via the API
	// A list of tags to associate with your service level objective. This can help you categorize and filter service level objectives in the service level objectives page of the UI. Note: it's not currently possible to filter by these tags when querying via the API
	// +kubebuilder:validation:Optional
	// +listType=set
	Tags []*string `json:"tags,omitempty" tf:"tags,omitempty"`

	// (Number) The objective's target in (0,100). This must match the corresponding thresholds of the primary time frame.
	// The objective's target in `(0,100)`. This must match the corresponding thresholds of the primary time frame.
	// +kubebuilder:validation:Optional
	TargetThreshold *float64 `json:"targetThreshold,omitempty" tf:"target_threshold,omitempty"`

	// (Block List, Min: 1) A list of thresholds and targets that define the service level objectives from the provided SLIs. (see below for nested schema)
	// A list of thresholds and targets that define the service level objectives from the provided SLIs.
	// +kubebuilder:validation:Optional
	Thresholds []ThresholdsParameters `json:"thresholds,omitempty" tf:"thresholds,omitempty"`

	// (String) The primary time frame for the objective. The mapping from these types to the types found in the Datadog Web UI can be found in the Datadog API documentation page. Valid values are 7d, 30d, 90d, custom.
	// The primary time frame for the objective. The mapping from these types to the types found in the Datadog Web UI can be found in the Datadog API documentation page. Valid values are `7d`, `30d`, `90d`, `custom`.
	// +kubebuilder:validation:Optional
	Timeframe *string `json:"timeframe,omitempty" tf:"timeframe,omitempty"`

	// (String) The type of the service level objective. The mapping from these types to the types found in the Datadog Web UI can be found in the Datadog API documentation page. Valid values are metric, monitor, time_slice.
	// The type of the service level objective. The mapping from these types to the types found in the Datadog Web UI can be found in the Datadog API [documentation page](https://docs.datadoghq.com/api/v1/service-level-objectives/#create-a-slo-object). Valid values are `metric`, `monitor`, `time_slice`.
	// +kubebuilder:validation:Optional
	Type *string `json:"type,omitempty" tf:"type,omitempty"`

	// (Boolean) Whether or not to validate the SLO. It checks if monitors added to a monitor SLO already exist.
	// Whether or not to validate the SLO. It checks if monitors added to a monitor SLO already exist.
	// +kubebuilder:validation:Optional
	Validate *bool `json:"validate,omitempty" tf:"validate,omitempty"`

	// (Number) The objective's warning value in (0,100). This must be greater than the target value and match the corresponding thresholds of the primary time frame.
	// The objective's warning value in `(0,100)`. This must be greater than the target value and match the corresponding thresholds of the primary time frame.
	// +kubebuilder:validation:Optional
	WarningThreshold *float64 `json:"warningThreshold,omitempty" tf:"warning_threshold,omitempty"`
}

type SliSpecificationInitParameters struct {

	// (Block List, Min: 1, Max: 1) The time slice condition, composed of 3 parts: 1. The timeseries query, 2. The comparator, and 3. The threshold. Optionally, a fourth part, the query interval, can be provided. (see below for nested schema)
	// The time slice condition, composed of 3 parts: 1. The timeseries query, 2. The comparator, and 3. The threshold. Optionally, a fourth part, the query interval, can be provided.
	TimeSlice []TimeSliceInitParameters `json:"timeSlice,omitempty" tf:"time_slice,omitempty"`
}

type SliSpecificationObservation struct {

	// (Block List, Min: 1, Max: 1) The time slice condition, composed of 3 parts: 1. The timeseries query, 2. The comparator, and 3. The threshold. Optionally, a fourth part, the query interval, can be provided. (see below for nested schema)
	// The time slice condition, composed of 3 parts: 1. The timeseries query, 2. The comparator, and 3. The threshold. Optionally, a fourth part, the query interval, can be provided.
	TimeSlice []TimeSliceObservation `json:"timeSlice,omitempty" tf:"time_slice,omitempty"`
}

type SliSpecificationParameters struct {

	// (Block List, Min: 1, Max: 1) The time slice condition, composed of 3 parts: 1. The timeseries query, 2. The comparator, and 3. The threshold. Optionally, a fourth part, the query interval, can be provided. (see below for nested schema)
	// The time slice condition, composed of 3 parts: 1. The timeseries query, 2. The comparator, and 3. The threshold. Optionally, a fourth part, the query interval, can be provided.
	// +kubebuilder:validation:Optional
	TimeSlice []TimeSliceParameters `json:"timeSlice" tf:"time_slice,omitempty"`
}

type ThresholdsInitParameters struct {

	// (Number) The objective's target in (0,100).
	// The objective's target in `(0,100)`.
	Target *float64 `json:"target,omitempty" tf:"target,omitempty"`

	// (String) The primary time frame for the objective. The mapping from these types to the types found in the Datadog Web UI can be found in the Datadog API documentation page. Valid values are 7d, 30d, 90d, custom.
	// The time frame for the objective. The mapping from these types to the types found in the Datadog Web UI can be found in the Datadog API documentation page. Valid values are `7d`, `30d`, `90d`, `custom`.
	Timeframe *string `json:"timeframe,omitempty" tf:"timeframe,omitempty"`

	// (Number) The objective's warning value in (0,100). This must be greater than the target value.
	// The objective's warning value in `(0,100)`. This must be greater than the target value.
	Warning *float64 `json:"warning,omitempty" tf:"warning,omitempty"`
}

type ThresholdsObservation struct {

	// (Number) The objective's target in (0,100).
	// The objective's target in `(0,100)`.
	Target *float64 `json:"target,omitempty" tf:"target,omitempty"`

	// (String) A string representation of the target that indicates its precision. It uses trailing zeros to show significant decimal places (e.g. 98.00).
	// A string representation of the target that indicates its precision. It uses trailing zeros to show significant decimal places (e.g. `98.00`).
	TargetDisplay *string `json:"targetDisplay,omitempty" tf:"target_display,omitempty"`

	// (String) The primary time frame for the objective. The mapping from these types to the types found in the Datadog Web UI can be found in the Datadog API documentation page. Valid values are 7d, 30d, 90d, custom.
	// The time frame for the objective. The mapping from these types to the types found in the Datadog Web UI can be found in the Datadog API documentation page. Valid values are `7d`, `30d`, `90d`, `custom`.
	Timeframe *string `json:"timeframe,omitempty" tf:"timeframe,omitempty"`

	// (Number) The objective's warning value in (0,100). This must be greater than the target value.
	// The objective's warning value in `(0,100)`. This must be greater than the target value.
	Warning *float64 `json:"warning,omitempty" tf:"warning,omitempty"`

	// (String) A string representation of the warning target (see the description of the target_display field for details).
	// A string representation of the warning target (see the description of the target_display field for details).
	WarningDisplay *string `json:"warningDisplay,omitempty" tf:"warning_display,omitempty"`
}

type ThresholdsParameters struct {

	// (Number) The objective's target in (0,100).
	// The objective's target in `(0,100)`.
	// +kubebuilder:validation:Optional
	Target *float64 `json:"target" tf:"target,omitempty"`

	// (String) The primary time frame for the objective. The mapping from these types to the types found in the Datadog Web UI can be found in the Datadog API documentation page. Valid values are 7d, 30d, 90d, custom.
	// The time frame for the objective. The mapping from these types to the types found in the Datadog Web UI can be found in the Datadog API documentation page. Valid values are `7d`, `30d`, `90d`, `custom`.
	// +kubebuilder:validation:Optional
	Timeframe *string `json:"timeframe" tf:"timeframe,omitempty"`

	// (Number) The objective's warning value in (0,100). This must be greater than the target value.
	// The objective's warning value in `(0,100)`. This must be greater than the target value.
	// +kubebuilder:validation:Optional
	Warning *float64 `json:"warning,omitempty" tf:"warning,omitempty"`
}

type TimeSliceInitParameters struct {

	// (String) The comparator used to compare the SLI value to the threshold. Valid values are >, >=, <, <=.
	// The comparator used to compare the SLI value to the threshold. Valid values are `>`, `>=`, `<`, `<=`.
	Comparator *string `json:"comparator,omitempty" tf:"comparator,omitempty"`

	// (Block List, Max: 1) The metric query of good / total events (see below for nested schema)
	// A timeseries query, containing named data-source-specific queries and a formula involving the named queries.
	Query []TimeSliceQueryInitParameters `json:"query,omitempty" tf:"query,omitempty"`

	// (Number) The interval used when querying data, which defines the size of a time slice. Valid values are 60, 300. Defaults to 300.
	// The interval used when querying data, which defines the size of a time slice. Valid values are `60`, `300`. Defaults to `300`.
	QueryIntervalSeconds *float64 `json:"queryIntervalSeconds,omitempty" tf:"query_interval_seconds,omitempty"`

	// (Number) The threshold value to which each SLI value will be compared.
	// The threshold value to which each SLI value will be compared.
	Threshold *float64 `json:"threshold,omitempty" tf:"threshold,omitempty"`
}

type TimeSliceObservation struct {

	// (String) The comparator used to compare the SLI value to the threshold. Valid values are >, >=, <, <=.
	// The comparator used to compare the SLI value to the threshold. Valid values are `>`, `>=`, `<`, `<=`.
	Comparator *string `json:"comparator,omitempty" tf:"comparator,omitempty"`

	// (Block List, Max: 1) The metric query of good / total events (see below for nested schema)
	// A timeseries query, containing named data-source-specific queries and a formula involving the named queries.
	Query []TimeSliceQueryObservation `json:"query,omitempty" tf:"query,omitempty"`

	// (Number) The interval used when querying data, which defines the size of a time slice. Valid values are 60, 300. Defaults to 300.
	// The interval used when querying data, which defines the size of a time slice. Valid values are `60`, `300`. Defaults to `300`.
	QueryIntervalSeconds *float64 `json:"queryIntervalSeconds,omitempty" tf:"query_interval_seconds,omitempty"`

	// (Number) The threshold value to which each SLI value will be compared.
	// The threshold value to which each SLI value will be compared.
	Threshold *float64 `json:"threshold,omitempty" tf:"threshold,omitempty"`
}

type TimeSliceParameters struct {

	// (String) The comparator used to compare the SLI value to the threshold. Valid values are >, >=, <, <=.
	// The comparator used to compare the SLI value to the threshold. Valid values are `>`, `>=`, `<`, `<=`.
	// +kubebuilder:validation:Optional
	Comparator *string `json:"comparator" tf:"comparator,omitempty"`

	// (Block List, Max: 1) The metric query of good / total events (see below for nested schema)
	// A timeseries query, containing named data-source-specific queries and a formula involving the named queries.
	// +kubebuilder:validation:Optional
	Query []TimeSliceQueryParameters `json:"query" tf:"query,omitempty"`

	// (Number) The interval used when querying data, which defines the size of a time slice. Valid values are 60, 300. Defaults to 300.
	// The interval used when querying data, which defines the size of a time slice. Valid values are `60`, `300`. Defaults to `300`.
	// +kubebuilder:validation:Optional
	QueryIntervalSeconds *float64 `json:"queryIntervalSeconds,omitempty" tf:"query_interval_seconds,omitempty"`

	// (Number) The threshold value to which each SLI value will be compared.
	// The threshold value to which each SLI value will be compared.
	// +kubebuilder:validation:Optional
	Threshold *float64 `json:"threshold" tf:"threshold,omitempty"`
}

type TimeSliceQueryInitParameters struct {

	// slice SLO. (see below for nested schema)
	// A list that contains exactly one formula, as only a single formula may be used to define a timeseries query for a time-slice SLO.
	Formula []FormulaInitParameters `json:"formula,omitempty" tf:"formula,omitempty"`

	// (Block List, Max: 1) The metric query of good / total events (see below for nested schema)
	// A list of data-source-specific queries that are in the formula.
	Query []QueryQueryInitParameters `json:"query,omitempty" tf:"query,omitempty"`
}

type TimeSliceQueryObservation struct {

	// slice SLO. (see below for nested schema)
	// A list that contains exactly one formula, as only a single formula may be used to define a timeseries query for a time-slice SLO.
	Formula []FormulaObservation `json:"formula,omitempty" tf:"formula,omitempty"`

	// (Block List, Max: 1) The metric query of good / total events (see below for nested schema)
	// A list of data-source-specific queries that are in the formula.
	Query []QueryQueryObservation `json:"query,omitempty" tf:"query,omitempty"`
}

type TimeSliceQueryParameters struct {

	// slice SLO. (see below for nested schema)
	// A list that contains exactly one formula, as only a single formula may be used to define a timeseries query for a time-slice SLO.
	// +kubebuilder:validation:Optional
	Formula []FormulaParameters `json:"formula" tf:"formula,omitempty"`

	// (Block List, Max: 1) The metric query of good / total events (see below for nested schema)
	// A list of data-source-specific queries that are in the formula.
	// +kubebuilder:validation:Optional
	Query []QueryQueryParameters `json:"query" tf:"query,omitempty"`
}

// ServiceLevelObjectiveSpec defines the desired state of ServiceLevelObjective
type ServiceLevelObjectiveSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     ServiceLevelObjectiveParameters `json:"forProvider"`
	// THIS IS A BETA FIELD. It will be honored
	// unless the Management Policies feature flag is disabled.
	// InitProvider holds the same fields as ForProvider, with the exception
	// of Identifier and other resource reference fields. The fields that are
	// in InitProvider are merged into ForProvider when the resource is created.
	// The same fields are also added to the terraform ignore_changes hook, to
	// avoid updating them after creation. This is useful for fields that are
	// required on creation, but we do not desire to update them after creation,
	// for example because of an external controller is managing them, like an
	// autoscaler.
	InitProvider ServiceLevelObjectiveInitParameters `json:"initProvider,omitempty"`
}

// ServiceLevelObjectiveStatus defines the observed state of ServiceLevelObjective.
type ServiceLevelObjectiveStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        ServiceLevelObjectiveObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true
// +kubebuilder:subresource:status
// +kubebuilder:storageversion

// ServiceLevelObjective is the Schema for the ServiceLevelObjectives API. Provides a Datadog service level objective resource. This can be used to create and manage Datadog service level objectives.
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,datadog}
type ServiceLevelObjective struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.name) || (has(self.initProvider) && has(self.initProvider.name))",message="spec.forProvider.name is a required parameter"
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.thresholds) || (has(self.initProvider) && has(self.initProvider.thresholds))",message="spec.forProvider.thresholds is a required parameter"
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.type) || (has(self.initProvider) && has(self.initProvider.type))",message="spec.forProvider.type is a required parameter"
	Spec   ServiceLevelObjectiveSpec   `json:"spec"`
	Status ServiceLevelObjectiveStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// ServiceLevelObjectiveList contains a list of ServiceLevelObjectives
type ServiceLevelObjectiveList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []ServiceLevelObjective `json:"items"`
}

// Repository type metadata.
var (
	ServiceLevelObjective_Kind             = "ServiceLevelObjective"
	ServiceLevelObjective_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: ServiceLevelObjective_Kind}.String()
	ServiceLevelObjective_KindAPIVersion   = ServiceLevelObjective_Kind + "." + CRDGroupVersion.String()
	ServiceLevelObjective_GroupVersionKind = CRDGroupVersion.WithKind(ServiceLevelObjective_Kind)
)

func init() {
	SchemeBuilder.Register(&ServiceLevelObjective{}, &ServiceLevelObjectiveList{})
}
