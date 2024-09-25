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

type IncludedKeywordConfigurationInitParameters struct {

	// (Number) Number of characters before the match to find a keyword validating the match. It must be between 1 and 50 (inclusive).
	// Number of characters before the match to find a keyword validating the match. It must be between 1 and 50 (inclusive).
	CharacterCount *float64 `json:"characterCount,omitempty" tf:"character_count,omitempty"`

	// (List of String) Keyword list that is checked during scanning in order to validate a match. The number of keywords in the list must be lower than or equal to 30.
	// Keyword list that is checked during scanning in order to validate a match. The number of keywords in the list must be lower than or equal to 30.
	Keywords []*string `json:"keywords,omitempty" tf:"keywords,omitempty"`
}

type IncludedKeywordConfigurationObservation struct {

	// (Number) Number of characters before the match to find a keyword validating the match. It must be between 1 and 50 (inclusive).
	// Number of characters before the match to find a keyword validating the match. It must be between 1 and 50 (inclusive).
	CharacterCount *float64 `json:"characterCount,omitempty" tf:"character_count,omitempty"`

	// (List of String) Keyword list that is checked during scanning in order to validate a match. The number of keywords in the list must be lower than or equal to 30.
	// Keyword list that is checked during scanning in order to validate a match. The number of keywords in the list must be lower than or equal to 30.
	Keywords []*string `json:"keywords,omitempty" tf:"keywords,omitempty"`
}

type IncludedKeywordConfigurationParameters struct {

	// (Number) Number of characters before the match to find a keyword validating the match. It must be between 1 and 50 (inclusive).
	// Number of characters before the match to find a keyword validating the match. It must be between 1 and 50 (inclusive).
	// +kubebuilder:validation:Optional
	CharacterCount *float64 `json:"characterCount" tf:"character_count,omitempty"`

	// (List of String) Keyword list that is checked during scanning in order to validate a match. The number of keywords in the list must be lower than or equal to 30.
	// Keyword list that is checked during scanning in order to validate a match. The number of keywords in the list must be lower than or equal to 30.
	// +kubebuilder:validation:Optional
	Keywords []*string `json:"keywords" tf:"keywords,omitempty"`
}

type RuleInitParameters struct {

	// (String) Description of the rule.
	// Description of the rule.
	Description *string `json:"description,omitempty" tf:"description,omitempty"`

	// path of the namespaces array.
	// Attributes excluded from the scan. If namespaces is provided, it has to be a sub-path of the namespaces array.
	ExcludedNamespaces []*string `json:"excludedNamespaces,omitempty" tf:"excluded_namespaces,omitempty"`

	// (String) Id of the scanning group the rule belongs to.
	// Id of the scanning group the rule belongs to.
	GroupID *string `json:"groupId,omitempty" tf:"group_id,omitempty"`

	// argument to true is highly recommended if modifying this field to avoid unexpectedly disabling Sensitive Data Scanner groups. (see below for nested schema)
	// Object defining a set of keywords and a number of characters that help reduce noise. You can provide a list of keywords you would like to check within a defined proximity of the matching pattern. If any of the keywords are found within the proximity check then the match is kept. If none are found, the match is discarded. Setting the `create_before_destroy` lifecycle Meta-argument to `true` is highly recommended if modifying this field to avoid unexpectedly disabling Sensitive Data Scanner groups.
	IncludedKeywordConfiguration []IncludedKeywordConfigurationInitParameters `json:"includedKeywordConfiguration,omitempty" tf:"included_keyword_configuration,omitempty"`

	// (Boolean) Whether or not the rule is enabled.
	// Whether or not the rule is enabled.
	IsEnabled *bool `json:"isEnabled,omitempty" tf:"is_enabled,omitempty"`

	// (String) Name of the rule.
	// Name of the rule.
	Name *string `json:"name,omitempty" tf:"name,omitempty"`

	// (List of String) Attributes included in the scan. If namespaces is empty or missing, all attributes except excluded_namespaces are scanned. If both are missing the whole event is scanned.
	// Attributes included in the scan. If namespaces is empty or missing, all attributes except excluded_namespaces are scanned. If both are missing the whole event is scanned.
	Namespaces []*string `json:"namespaces,omitempty" tf:"namespaces,omitempty"`

	// (String) Not included if there is a relationship to a standard pattern.
	// Not included if there is a relationship to a standard pattern.
	Pattern *string `json:"pattern,omitempty" tf:"pattern,omitempty"`

	// (Number) Priority level of the rule . Used to order sensitive data discovered in the sds summary page. It must be between 1 and 5 (1 being the most important).
	// Priority level of the rule (optional). Used to order sensitive data discovered in the sds summary page. It must be between 1 and 5 (1 being the most important).
	Priority *float64 `json:"priority,omitempty" tf:"priority,omitempty"`

	// (String) Id of the standard pattern the rule refers to. If provided, then pattern must not be provided.
	// Id of the standard pattern the rule refers to. If provided, then pattern must not be provided.
	StandardPatternID *string `json:"standardPatternId,omitempty" tf:"standard_pattern_id,omitempty"`

	// (List of String) List of tags.
	// List of tags.
	Tags []*string `json:"tags,omitempty" tf:"tags,omitempty"`

	// (Block List, Max: 1) Object describing how the scanned event will be replaced. Defaults to type: none (see below for nested schema)
	// Object describing how the scanned event will be replaced. Defaults to `type: none`
	TextReplacement []TextReplacementInitParameters `json:"textReplacement,omitempty" tf:"text_replacement,omitempty"`
}

type RuleObservation struct {

	// (String) Description of the rule.
	// Description of the rule.
	Description *string `json:"description,omitempty" tf:"description,omitempty"`

	// path of the namespaces array.
	// Attributes excluded from the scan. If namespaces is provided, it has to be a sub-path of the namespaces array.
	ExcludedNamespaces []*string `json:"excludedNamespaces,omitempty" tf:"excluded_namespaces,omitempty"`

	// (String) Id of the scanning group the rule belongs to.
	// Id of the scanning group the rule belongs to.
	GroupID *string `json:"groupId,omitempty" tf:"group_id,omitempty"`

	// (String) The ID of this resource.
	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	// argument to true is highly recommended if modifying this field to avoid unexpectedly disabling Sensitive Data Scanner groups. (see below for nested schema)
	// Object defining a set of keywords and a number of characters that help reduce noise. You can provide a list of keywords you would like to check within a defined proximity of the matching pattern. If any of the keywords are found within the proximity check then the match is kept. If none are found, the match is discarded. Setting the `create_before_destroy` lifecycle Meta-argument to `true` is highly recommended if modifying this field to avoid unexpectedly disabling Sensitive Data Scanner groups.
	IncludedKeywordConfiguration []IncludedKeywordConfigurationObservation `json:"includedKeywordConfiguration,omitempty" tf:"included_keyword_configuration,omitempty"`

	// (Boolean) Whether or not the rule is enabled.
	// Whether or not the rule is enabled.
	IsEnabled *bool `json:"isEnabled,omitempty" tf:"is_enabled,omitempty"`

	// (String) Name of the rule.
	// Name of the rule.
	Name *string `json:"name,omitempty" tf:"name,omitempty"`

	// (List of String) Attributes included in the scan. If namespaces is empty or missing, all attributes except excluded_namespaces are scanned. If both are missing the whole event is scanned.
	// Attributes included in the scan. If namespaces is empty or missing, all attributes except excluded_namespaces are scanned. If both are missing the whole event is scanned.
	Namespaces []*string `json:"namespaces,omitempty" tf:"namespaces,omitempty"`

	// (String) Not included if there is a relationship to a standard pattern.
	// Not included if there is a relationship to a standard pattern.
	Pattern *string `json:"pattern,omitempty" tf:"pattern,omitempty"`

	// (Number) Priority level of the rule . Used to order sensitive data discovered in the sds summary page. It must be between 1 and 5 (1 being the most important).
	// Priority level of the rule (optional). Used to order sensitive data discovered in the sds summary page. It must be between 1 and 5 (1 being the most important).
	Priority *float64 `json:"priority,omitempty" tf:"priority,omitempty"`

	// (String) Id of the standard pattern the rule refers to. If provided, then pattern must not be provided.
	// Id of the standard pattern the rule refers to. If provided, then pattern must not be provided.
	StandardPatternID *string `json:"standardPatternId,omitempty" tf:"standard_pattern_id,omitempty"`

	// (List of String) List of tags.
	// List of tags.
	Tags []*string `json:"tags,omitempty" tf:"tags,omitempty"`

	// (Block List, Max: 1) Object describing how the scanned event will be replaced. Defaults to type: none (see below for nested schema)
	// Object describing how the scanned event will be replaced. Defaults to `type: none`
	TextReplacement []TextReplacementObservation `json:"textReplacement,omitempty" tf:"text_replacement,omitempty"`
}

type RuleParameters struct {

	// (String) Description of the rule.
	// Description of the rule.
	// +kubebuilder:validation:Optional
	Description *string `json:"description,omitempty" tf:"description,omitempty"`

	// path of the namespaces array.
	// Attributes excluded from the scan. If namespaces is provided, it has to be a sub-path of the namespaces array.
	// +kubebuilder:validation:Optional
	ExcludedNamespaces []*string `json:"excludedNamespaces,omitempty" tf:"excluded_namespaces,omitempty"`

	// (String) Id of the scanning group the rule belongs to.
	// Id of the scanning group the rule belongs to.
	// +kubebuilder:validation:Optional
	GroupID *string `json:"groupId,omitempty" tf:"group_id,omitempty"`

	// argument to true is highly recommended if modifying this field to avoid unexpectedly disabling Sensitive Data Scanner groups. (see below for nested schema)
	// Object defining a set of keywords and a number of characters that help reduce noise. You can provide a list of keywords you would like to check within a defined proximity of the matching pattern. If any of the keywords are found within the proximity check then the match is kept. If none are found, the match is discarded. Setting the `create_before_destroy` lifecycle Meta-argument to `true` is highly recommended if modifying this field to avoid unexpectedly disabling Sensitive Data Scanner groups.
	// +kubebuilder:validation:Optional
	IncludedKeywordConfiguration []IncludedKeywordConfigurationParameters `json:"includedKeywordConfiguration,omitempty" tf:"included_keyword_configuration,omitempty"`

	// (Boolean) Whether or not the rule is enabled.
	// Whether or not the rule is enabled.
	// +kubebuilder:validation:Optional
	IsEnabled *bool `json:"isEnabled,omitempty" tf:"is_enabled,omitempty"`

	// (String) Name of the rule.
	// Name of the rule.
	// +kubebuilder:validation:Optional
	Name *string `json:"name,omitempty" tf:"name,omitempty"`

	// (List of String) Attributes included in the scan. If namespaces is empty or missing, all attributes except excluded_namespaces are scanned. If both are missing the whole event is scanned.
	// Attributes included in the scan. If namespaces is empty or missing, all attributes except excluded_namespaces are scanned. If both are missing the whole event is scanned.
	// +kubebuilder:validation:Optional
	Namespaces []*string `json:"namespaces,omitempty" tf:"namespaces,omitempty"`

	// (String) Not included if there is a relationship to a standard pattern.
	// Not included if there is a relationship to a standard pattern.
	// +kubebuilder:validation:Optional
	Pattern *string `json:"pattern,omitempty" tf:"pattern,omitempty"`

	// (Number) Priority level of the rule . Used to order sensitive data discovered in the sds summary page. It must be between 1 and 5 (1 being the most important).
	// Priority level of the rule (optional). Used to order sensitive data discovered in the sds summary page. It must be between 1 and 5 (1 being the most important).
	// +kubebuilder:validation:Optional
	Priority *float64 `json:"priority,omitempty" tf:"priority,omitempty"`

	// (String) Id of the standard pattern the rule refers to. If provided, then pattern must not be provided.
	// Id of the standard pattern the rule refers to. If provided, then pattern must not be provided.
	// +kubebuilder:validation:Optional
	StandardPatternID *string `json:"standardPatternId,omitempty" tf:"standard_pattern_id,omitempty"`

	// (List of String) List of tags.
	// List of tags.
	// +kubebuilder:validation:Optional
	Tags []*string `json:"tags,omitempty" tf:"tags,omitempty"`

	// (Block List, Max: 1) Object describing how the scanned event will be replaced. Defaults to type: none (see below for nested schema)
	// Object describing how the scanned event will be replaced. Defaults to `type: none`
	// +kubebuilder:validation:Optional
	TextReplacement []TextReplacementParameters `json:"textReplacement,omitempty" tf:"text_replacement,omitempty"`
}

type TextReplacementInitParameters struct {

	// (Number) Required if type == 'partial_replacement_from_beginning' or 'partial_replacement_from_end'. It must be > 0.
	// Required if type == 'partial_replacement_from_beginning' or 'partial_replacement_from_end'. It must be > 0.
	NumberOfChars *float64 `json:"numberOfChars,omitempty" tf:"number_of_chars,omitempty"`

	// (String) Required if type == 'replacement_string'.
	// Required if type == 'replacement_string'.
	ReplacementString *string `json:"replacementString,omitempty" tf:"replacement_string,omitempty"`

	// (String) Type of the replacement text. None means no replacement. hash means the data will be stubbed. replacement_string means that one can chose a text to replace the data. partial_replacement_from_beginning allows a user to partially replace the data from the beginning, and partial_replacement_from_end on the other hand, allows to replace data from the end. Valid values are none, hash, replacement_string, partial_replacement_from_beginning, partial_replacement_from_end.
	// Type of the replacement text. None means no replacement. hash means the data will be stubbed. replacement_string means that one can chose a text to replace the data. partial_replacement_from_beginning allows a user to partially replace the data from the beginning, and partial_replacement_from_end on the other hand, allows to replace data from the end. Valid values are `none`, `hash`, `replacement_string`, `partial_replacement_from_beginning`, `partial_replacement_from_end`.
	Type *string `json:"type,omitempty" tf:"type,omitempty"`
}

type TextReplacementObservation struct {

	// (Number) Required if type == 'partial_replacement_from_beginning' or 'partial_replacement_from_end'. It must be > 0.
	// Required if type == 'partial_replacement_from_beginning' or 'partial_replacement_from_end'. It must be > 0.
	NumberOfChars *float64 `json:"numberOfChars,omitempty" tf:"number_of_chars,omitempty"`

	// (String) Required if type == 'replacement_string'.
	// Required if type == 'replacement_string'.
	ReplacementString *string `json:"replacementString,omitempty" tf:"replacement_string,omitempty"`

	// (String) Type of the replacement text. None means no replacement. hash means the data will be stubbed. replacement_string means that one can chose a text to replace the data. partial_replacement_from_beginning allows a user to partially replace the data from the beginning, and partial_replacement_from_end on the other hand, allows to replace data from the end. Valid values are none, hash, replacement_string, partial_replacement_from_beginning, partial_replacement_from_end.
	// Type of the replacement text. None means no replacement. hash means the data will be stubbed. replacement_string means that one can chose a text to replace the data. partial_replacement_from_beginning allows a user to partially replace the data from the beginning, and partial_replacement_from_end on the other hand, allows to replace data from the end. Valid values are `none`, `hash`, `replacement_string`, `partial_replacement_from_beginning`, `partial_replacement_from_end`.
	Type *string `json:"type,omitempty" tf:"type,omitempty"`
}

type TextReplacementParameters struct {

	// (Number) Required if type == 'partial_replacement_from_beginning' or 'partial_replacement_from_end'. It must be > 0.
	// Required if type == 'partial_replacement_from_beginning' or 'partial_replacement_from_end'. It must be > 0.
	// +kubebuilder:validation:Optional
	NumberOfChars *float64 `json:"numberOfChars,omitempty" tf:"number_of_chars,omitempty"`

	// (String) Required if type == 'replacement_string'.
	// Required if type == 'replacement_string'.
	// +kubebuilder:validation:Optional
	ReplacementString *string `json:"replacementString,omitempty" tf:"replacement_string,omitempty"`

	// (String) Type of the replacement text. None means no replacement. hash means the data will be stubbed. replacement_string means that one can chose a text to replace the data. partial_replacement_from_beginning allows a user to partially replace the data from the beginning, and partial_replacement_from_end on the other hand, allows to replace data from the end. Valid values are none, hash, replacement_string, partial_replacement_from_beginning, partial_replacement_from_end.
	// Type of the replacement text. None means no replacement. hash means the data will be stubbed. replacement_string means that one can chose a text to replace the data. partial_replacement_from_beginning allows a user to partially replace the data from the beginning, and partial_replacement_from_end on the other hand, allows to replace data from the end. Valid values are `none`, `hash`, `replacement_string`, `partial_replacement_from_beginning`, `partial_replacement_from_end`.
	// +kubebuilder:validation:Optional
	Type *string `json:"type" tf:"type,omitempty"`
}

// RuleSpec defines the desired state of Rule
type RuleSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     RuleParameters `json:"forProvider"`
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
	InitProvider RuleInitParameters `json:"initProvider,omitempty"`
}

// RuleStatus defines the observed state of Rule.
type RuleStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        RuleObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true
// +kubebuilder:subresource:status
// +kubebuilder:storageversion

// Rule is the Schema for the Rules API. Provides a Datadog SensitiveDataScannerRule resource. This can be used to create and manage Datadog sensitivedatascanner_rule. Setting the create_before_destroy lifecycle Meta-argument to true is highly recommended if modifying the included_keyword_configuration field to avoid unexpectedly disabling Sensitive Data Scanner groups.
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,datadog}
type Rule struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.groupId) || (has(self.initProvider) && has(self.initProvider.groupId))",message="spec.forProvider.groupId is a required parameter"
	Spec   RuleSpec   `json:"spec"`
	Status RuleStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// RuleList contains a list of Rules
type RuleList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []Rule `json:"items"`
}

// Repository type metadata.
var (
	Rule_Kind             = "Rule"
	Rule_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: Rule_Kind}.String()
	Rule_KindAPIVersion   = Rule_Kind + "." + CRDGroupVersion.String()
	Rule_GroupVersionKind = CRDGroupVersion.WithKind(Rule_Kind)
)

func init() {
	SchemeBuilder.Register(&Rule{}, &RuleList{})
}
