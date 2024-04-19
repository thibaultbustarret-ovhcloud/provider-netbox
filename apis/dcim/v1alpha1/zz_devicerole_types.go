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

type DeviceRoleInitParameters struct {

	// (String)
	ColorHex *string `json:"colorHex,omitempty" tf:"color_hex,omitempty"`

	// (String)
	Description *string `json:"description,omitempty" tf:"description,omitempty"`

	// (String)
	Slug *string `json:"slug,omitempty" tf:"slug,omitempty"`

	// (Set of String)
	Tags []*string `json:"tags,omitempty" tf:"tags,omitempty"`

	// (Boolean) Defaults to true.
	// Defaults to `true`.
	VMRole *bool `json:"vmRole,omitempty" tf:"vm_role,omitempty"`
}

type DeviceRoleObservation struct {

	// (String)
	ColorHex *string `json:"colorHex,omitempty" tf:"color_hex,omitempty"`

	// (String)
	Description *string `json:"description,omitempty" tf:"description,omitempty"`

	// (String) The ID of this resource.
	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	// (String)
	Slug *string `json:"slug,omitempty" tf:"slug,omitempty"`

	// (Set of String)
	Tags []*string `json:"tags,omitempty" tf:"tags,omitempty"`

	// (Boolean) Defaults to true.
	// Defaults to `true`.
	VMRole *bool `json:"vmRole,omitempty" tf:"vm_role,omitempty"`
}

type DeviceRoleParameters struct {

	// (String)
	// +kubebuilder:validation:Optional
	ColorHex *string `json:"colorHex,omitempty" tf:"color_hex,omitempty"`

	// (String)
	// +kubebuilder:validation:Optional
	Description *string `json:"description,omitempty" tf:"description,omitempty"`

	// (String)
	// +kubebuilder:validation:Optional
	Slug *string `json:"slug,omitempty" tf:"slug,omitempty"`

	// (Set of String)
	// +kubebuilder:validation:Optional
	Tags []*string `json:"tags,omitempty" tf:"tags,omitempty"`

	// (Boolean) Defaults to true.
	// Defaults to `true`.
	// +kubebuilder:validation:Optional
	VMRole *bool `json:"vmRole,omitempty" tf:"vm_role,omitempty"`
}

// DeviceRoleSpec defines the desired state of DeviceRole
type DeviceRoleSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     DeviceRoleParameters `json:"forProvider"`
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
	InitProvider DeviceRoleInitParameters `json:"initProvider,omitempty"`
}

// DeviceRoleStatus defines the observed state of DeviceRole.
type DeviceRoleStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        DeviceRoleObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// DeviceRole is the Schema for the DeviceRoles API. From the official documentation https://docs.netbox.dev/en/stable/features/devices/#device-roles: Devices can be organized by functional roles, which are fully customizable by the user. For example, you might create roles for core switches, distribution switches, and access switches within your network.
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,netbox}
type DeviceRole struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.colorHex) || (has(self.initProvider) && has(self.initProvider.colorHex))",message="spec.forProvider.colorHex is a required parameter"
	Spec   DeviceRoleSpec   `json:"spec"`
	Status DeviceRoleStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// DeviceRoleList contains a list of DeviceRoles
type DeviceRoleList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []DeviceRole `json:"items"`
}

// Repository type metadata.
var (
	DeviceRole_Kind             = "DeviceRole"
	DeviceRole_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: DeviceRole_Kind}.String()
	DeviceRole_KindAPIVersion   = DeviceRole_Kind + "." + CRDGroupVersion.String()
	DeviceRole_GroupVersionKind = CRDGroupVersion.WithKind(DeviceRole_Kind)
)

func init() {
	SchemeBuilder.Register(&DeviceRole{}, &DeviceRoleList{})
}
