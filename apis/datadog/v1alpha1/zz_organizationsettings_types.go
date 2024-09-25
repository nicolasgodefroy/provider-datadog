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

type OrganizationSettingsInitParameters struct {

	// (String) Name for Organization.
	// Name for Organization.
	Name *string `json:"name,omitempty" tf:"name,omitempty"`

	// (List of String) List of emails used for security event notifications from the organization.
	// List of emails used for security event notifications from the organization.
	SecurityContacts []*string `json:"securityContacts,omitempty" tf:"security_contacts,omitempty"`

	// (Block List, Max: 1) Organization settings (see below for nested schema)
	// Organization settings
	Settings []OrganizationSettingsSettingsInitParameters `json:"settings,omitempty" tf:"settings,omitempty"`
}

type OrganizationSettingsObservation struct {

	// (String) Description of the organization.
	// Description of the organization.
	Description *string `json:"description,omitempty" tf:"description,omitempty"`

	// (String) The ID of this resource.
	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	// (String) Name for Organization.
	// Name for Organization.
	Name *string `json:"name,omitempty" tf:"name,omitempty"`

	// (String) The public_id of the organization you are operating within.
	// The `public_id` of the organization you are operating within.
	PublicID *string `json:"publicId,omitempty" tf:"public_id,omitempty"`

	// (List of String) List of emails used for security event notifications from the organization.
	// List of emails used for security event notifications from the organization.
	SecurityContacts []*string `json:"securityContacts,omitempty" tf:"security_contacts,omitempty"`

	// (Block List, Max: 1) Organization settings (see below for nested schema)
	// Organization settings
	Settings []OrganizationSettingsSettingsObservation `json:"settings,omitempty" tf:"settings,omitempty"`
}

type OrganizationSettingsParameters struct {

	// (String) Name for Organization.
	// Name for Organization.
	// +kubebuilder:validation:Optional
	Name *string `json:"name,omitempty" tf:"name,omitempty"`

	// (List of String) List of emails used for security event notifications from the organization.
	// List of emails used for security event notifications from the organization.
	// +kubebuilder:validation:Optional
	SecurityContacts []*string `json:"securityContacts,omitempty" tf:"security_contacts,omitempty"`

	// (Block List, Max: 1) Organization settings (see below for nested schema)
	// Organization settings
	// +kubebuilder:validation:Optional
	Settings []OrganizationSettingsSettingsParameters `json:"settings,omitempty" tf:"settings,omitempty"`
}

type OrganizationSettingsSettingsInitParameters struct {

	// (Boolean) Whether or not the organization users can share widgets outside of Datadog. Defaults to false.
	// Whether or not the organization users can share widgets outside of Datadog. Defaults to `false`.
	PrivateWidgetShare *bool `json:"privateWidgetShare,omitempty" tf:"private_widget_share,omitempty"`

	// (Block List, Min: 1, Max: 1) SAML properties (see below for nested schema)
	// SAML properties
	SAML []SettingsSAMLInitParameters `json:"saml,omitempty" tf:"saml,omitempty"`

	// only user). Allowed enum values: st, adm , ro, ERROR Defaults to "st".
	// The access role of the user. Options are `st` (standard user), `adm` (admin user), or `ro` (read-only user). Allowed enum values: `st`, `adm` , `ro`, `ERROR` Defaults to `"st"`.
	SAMLAutocreateAccessRole *string `json:"samlAutocreateAccessRole,omitempty" tf:"saml_autocreate_access_role,omitempty"`

	// (Block List, Min: 1, Max: 1) List of domains where the SAML automated user creation is enabled. (see below for nested schema)
	// List of domains where the SAML automated user creation is enabled.
	SAMLAutocreateUsersDomains []SettingsSAMLAutocreateUsersDomainsInitParameters `json:"samlAutocreateUsersDomains,omitempty" tf:"saml_autocreate_users_domains,omitempty"`

	// (Block List, Min: 1, Max: 1) Whether or not a SAML identity provider metadata file was provided to the Datadog organization. (see below for nested schema)
	// Whether or not a SAML identity provider metadata file was provided to the Datadog organization.
	SAMLIdpInitiatedLogin []SettingsSAMLIdpInitiatedLoginInitParameters `json:"samlIdpInitiatedLogin,omitempty" tf:"saml_idp_initiated_login,omitempty"`

	// (Block List, Min: 1, Max: 1) Whether or not the SAML strict mode is enabled. If true, all users must log in with SAML. (see below for nested schema)
	// Whether or not the SAML strict mode is enabled. If true, all users must log in with SAML.
	SAMLStrictMode []SettingsSAMLStrictModeInitParameters `json:"samlStrictMode,omitempty" tf:"saml_strict_mode,omitempty"`
}

type OrganizationSettingsSettingsObservation struct {

	// (Boolean) Whether or not the organization users can share widgets outside of Datadog. Defaults to false.
	// Whether or not the organization users can share widgets outside of Datadog. Defaults to `false`.
	PrivateWidgetShare *bool `json:"privateWidgetShare,omitempty" tf:"private_widget_share,omitempty"`

	// (Block List, Min: 1, Max: 1) SAML properties (see below for nested schema)
	// SAML properties
	SAML []SettingsSAMLObservation `json:"saml,omitempty" tf:"saml,omitempty"`

	// only user). Allowed enum values: st, adm , ro, ERROR Defaults to "st".
	// The access role of the user. Options are `st` (standard user), `adm` (admin user), or `ro` (read-only user). Allowed enum values: `st`, `adm` , `ro`, `ERROR` Defaults to `"st"`.
	SAMLAutocreateAccessRole *string `json:"samlAutocreateAccessRole,omitempty" tf:"saml_autocreate_access_role,omitempty"`

	// (Block List, Min: 1, Max: 1) List of domains where the SAML automated user creation is enabled. (see below for nested schema)
	// List of domains where the SAML automated user creation is enabled.
	SAMLAutocreateUsersDomains []SettingsSAMLAutocreateUsersDomainsObservation `json:"samlAutocreateUsersDomains,omitempty" tf:"saml_autocreate_users_domains,omitempty"`

	// (Boolean) Whether or not SAML can be enabled for this organization.
	// Whether or not SAML can be enabled for this organization.
	SAMLCanBeEnabled *bool `json:"samlCanBeEnabled,omitempty" tf:"saml_can_be_enabled,omitempty"`

	// (String) Identity provider endpoint for SAML authentication.
	// Identity provider endpoint for SAML authentication.
	SAMLIdpEndpoint *string `json:"samlIdpEndpoint,omitempty" tf:"saml_idp_endpoint,omitempty"`

	// (Block List, Min: 1, Max: 1) Whether or not a SAML identity provider metadata file was provided to the Datadog organization. (see below for nested schema)
	// Whether or not a SAML identity provider metadata file was provided to the Datadog organization.
	SAMLIdpInitiatedLogin []SettingsSAMLIdpInitiatedLoginObservation `json:"samlIdpInitiatedLogin,omitempty" tf:"saml_idp_initiated_login,omitempty"`

	// (Boolean) Whether or not a SAML identity provider metadata file was provided to the Datadog organization.
	// Whether or not a SAML identity provider metadata file was provided to the Datadog organization.
	SAMLIdpMetadataUploaded *bool `json:"samlIdpMetadataUploaded,omitempty" tf:"saml_idp_metadata_uploaded,omitempty"`

	// (String) URL for SAML logging.
	// URL for SAML logging.
	SAMLLoginURL *string `json:"samlLoginUrl,omitempty" tf:"saml_login_url,omitempty"`

	// (Block List, Min: 1, Max: 1) Whether or not the SAML strict mode is enabled. If true, all users must log in with SAML. (see below for nested schema)
	// Whether or not the SAML strict mode is enabled. If true, all users must log in with SAML.
	SAMLStrictMode []SettingsSAMLStrictModeObservation `json:"samlStrictMode,omitempty" tf:"saml_strict_mode,omitempty"`
}

type OrganizationSettingsSettingsParameters struct {

	// (Boolean) Whether or not the organization users can share widgets outside of Datadog. Defaults to false.
	// Whether or not the organization users can share widgets outside of Datadog. Defaults to `false`.
	// +kubebuilder:validation:Optional
	PrivateWidgetShare *bool `json:"privateWidgetShare,omitempty" tf:"private_widget_share,omitempty"`

	// (Block List, Min: 1, Max: 1) SAML properties (see below for nested schema)
	// SAML properties
	// +kubebuilder:validation:Optional
	SAML []SettingsSAMLParameters `json:"saml" tf:"saml,omitempty"`

	// only user). Allowed enum values: st, adm , ro, ERROR Defaults to "st".
	// The access role of the user. Options are `st` (standard user), `adm` (admin user), or `ro` (read-only user). Allowed enum values: `st`, `adm` , `ro`, `ERROR` Defaults to `"st"`.
	// +kubebuilder:validation:Optional
	SAMLAutocreateAccessRole *string `json:"samlAutocreateAccessRole,omitempty" tf:"saml_autocreate_access_role,omitempty"`

	// (Block List, Min: 1, Max: 1) List of domains where the SAML automated user creation is enabled. (see below for nested schema)
	// List of domains where the SAML automated user creation is enabled.
	// +kubebuilder:validation:Optional
	SAMLAutocreateUsersDomains []SettingsSAMLAutocreateUsersDomainsParameters `json:"samlAutocreateUsersDomains" tf:"saml_autocreate_users_domains,omitempty"`

	// (Block List, Min: 1, Max: 1) Whether or not a SAML identity provider metadata file was provided to the Datadog organization. (see below for nested schema)
	// Whether or not a SAML identity provider metadata file was provided to the Datadog organization.
	// +kubebuilder:validation:Optional
	SAMLIdpInitiatedLogin []SettingsSAMLIdpInitiatedLoginParameters `json:"samlIdpInitiatedLogin" tf:"saml_idp_initiated_login,omitempty"`

	// (Block List, Min: 1, Max: 1) Whether or not the SAML strict mode is enabled. If true, all users must log in with SAML. (see below for nested schema)
	// Whether or not the SAML strict mode is enabled. If true, all users must log in with SAML.
	// +kubebuilder:validation:Optional
	SAMLStrictMode []SettingsSAMLStrictModeParameters `json:"samlStrictMode" tf:"saml_strict_mode,omitempty"`
}

type SettingsSAMLAutocreateUsersDomainsInitParameters struct {

	// (List of String) List of domains where the SAML automated user creation is enabled.
	// List of domains where the SAML automated user creation is enabled.
	Domains []*string `json:"domains,omitempty" tf:"domains,omitempty"`

	// (Boolean) Whether or not SAML is enabled for this organization. Defaults to false.
	// Whether or not the automated user creation based on SAML domain is enabled. Defaults to `false`.
	Enabled *bool `json:"enabled,omitempty" tf:"enabled,omitempty"`
}

type SettingsSAMLAutocreateUsersDomainsObservation struct {

	// (List of String) List of domains where the SAML automated user creation is enabled.
	// List of domains where the SAML automated user creation is enabled.
	Domains []*string `json:"domains,omitempty" tf:"domains,omitempty"`

	// (Boolean) Whether or not SAML is enabled for this organization. Defaults to false.
	// Whether or not the automated user creation based on SAML domain is enabled. Defaults to `false`.
	Enabled *bool `json:"enabled,omitempty" tf:"enabled,omitempty"`
}

type SettingsSAMLAutocreateUsersDomainsParameters struct {

	// (List of String) List of domains where the SAML automated user creation is enabled.
	// List of domains where the SAML automated user creation is enabled.
	// +kubebuilder:validation:Optional
	Domains []*string `json:"domains,omitempty" tf:"domains,omitempty"`

	// (Boolean) Whether or not SAML is enabled for this organization. Defaults to false.
	// Whether or not the automated user creation based on SAML domain is enabled. Defaults to `false`.
	// +kubebuilder:validation:Optional
	Enabled *bool `json:"enabled,omitempty" tf:"enabled,omitempty"`
}

type SettingsSAMLIdpInitiatedLoginInitParameters struct {

	// (Boolean) Whether or not SAML is enabled for this organization. Defaults to false.
	// Whether or not a SAML identity provider metadata file was provided to the Datadog organization. Defaults to `false`.
	Enabled *bool `json:"enabled,omitempty" tf:"enabled,omitempty"`
}

type SettingsSAMLIdpInitiatedLoginObservation struct {

	// (Boolean) Whether or not SAML is enabled for this organization. Defaults to false.
	// Whether or not a SAML identity provider metadata file was provided to the Datadog organization. Defaults to `false`.
	Enabled *bool `json:"enabled,omitempty" tf:"enabled,omitempty"`
}

type SettingsSAMLIdpInitiatedLoginParameters struct {

	// (Boolean) Whether or not SAML is enabled for this organization. Defaults to false.
	// Whether or not a SAML identity provider metadata file was provided to the Datadog organization. Defaults to `false`.
	// +kubebuilder:validation:Optional
	Enabled *bool `json:"enabled,omitempty" tf:"enabled,omitempty"`
}

type SettingsSAMLInitParameters struct {

	// (Boolean) Whether or not SAML is enabled for this organization. Defaults to false.
	// Whether or not SAML is enabled for this organization. Defaults to `false`.
	Enabled *bool `json:"enabled,omitempty" tf:"enabled,omitempty"`
}

type SettingsSAMLObservation struct {

	// (Boolean) Whether or not SAML is enabled for this organization. Defaults to false.
	// Whether or not SAML is enabled for this organization. Defaults to `false`.
	Enabled *bool `json:"enabled,omitempty" tf:"enabled,omitempty"`
}

type SettingsSAMLParameters struct {

	// (Boolean) Whether or not SAML is enabled for this organization. Defaults to false.
	// Whether or not SAML is enabled for this organization. Defaults to `false`.
	// +kubebuilder:validation:Optional
	Enabled *bool `json:"enabled,omitempty" tf:"enabled,omitempty"`
}

type SettingsSAMLStrictModeInitParameters struct {

	// (Boolean) Whether or not SAML is enabled for this organization. Defaults to false.
	// Whether or not the SAML strict mode is enabled. If true, all users must log in with SAML. Defaults to `false`.
	Enabled *bool `json:"enabled,omitempty" tf:"enabled,omitempty"`
}

type SettingsSAMLStrictModeObservation struct {

	// (Boolean) Whether or not SAML is enabled for this organization. Defaults to false.
	// Whether or not the SAML strict mode is enabled. If true, all users must log in with SAML. Defaults to `false`.
	Enabled *bool `json:"enabled,omitempty" tf:"enabled,omitempty"`
}

type SettingsSAMLStrictModeParameters struct {

	// (Boolean) Whether or not SAML is enabled for this organization. Defaults to false.
	// Whether or not the SAML strict mode is enabled. If true, all users must log in with SAML. Defaults to `false`.
	// +kubebuilder:validation:Optional
	Enabled *bool `json:"enabled,omitempty" tf:"enabled,omitempty"`
}

// OrganizationSettingsSpec defines the desired state of OrganizationSettings
type OrganizationSettingsSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     OrganizationSettingsParameters `json:"forProvider"`
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
	InitProvider OrganizationSettingsInitParameters `json:"initProvider,omitempty"`
}

// OrganizationSettingsStatus defines the observed state of OrganizationSettings.
type OrganizationSettingsStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        OrganizationSettingsObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true
// +kubebuilder:subresource:status
// +kubebuilder:storageversion

// OrganizationSettings is the Schema for the OrganizationSettingss API. Provides a Datadog Organization resource. This can be used to manage your Datadog organization's settings.
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,datadog}
type OrganizationSettings struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              OrganizationSettingsSpec   `json:"spec"`
	Status            OrganizationSettingsStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// OrganizationSettingsList contains a list of OrganizationSettingss
type OrganizationSettingsList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []OrganizationSettings `json:"items"`
}

// Repository type metadata.
var (
	OrganizationSettings_Kind             = "OrganizationSettings"
	OrganizationSettings_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: OrganizationSettings_Kind}.String()
	OrganizationSettings_KindAPIVersion   = OrganizationSettings_Kind + "." + CRDGroupVersion.String()
	OrganizationSettings_GroupVersionKind = CRDGroupVersion.WithKind(OrganizationSettings_Kind)
)

func init() {
	SchemeBuilder.Register(&OrganizationSettings{}, &OrganizationSettingsList{})
}
