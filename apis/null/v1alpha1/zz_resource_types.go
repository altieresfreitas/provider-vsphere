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

type ResourceObservation struct {
	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	// A map of arbitrary strings that, when changed, will force the null resource to be replaced, re-running any associated provisioners.
	Triggers map[string]*string `json:"triggers,omitempty" tf:"triggers,omitempty"`
}

type ResourceParameters struct {

	// A map of arbitrary strings that, when changed, will force the null resource to be replaced, re-running any associated provisioners.
	// +kubebuilder:validation:Optional
	Triggers map[string]*string `json:"triggers,omitempty" tf:"triggers,omitempty"`
}

// ResourceSpec defines the desired state of Resource
type ResourceSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     ResourceParameters `json:"forProvider"`
}

// ResourceStatus defines the observed state of Resource.
type ResourceStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        ResourceObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// Resource is the Schema for the Resources API. <no value>
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,vsphere}
type Resource struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              ResourceSpec   `json:"spec"`
	Status            ResourceStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// ResourceList contains a list of Resources
type ResourceList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []Resource `json:"items"`
}

// Repository type metadata.
var (
	Resource_Kind             = "Resource"
	Resource_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: Resource_Kind}.String()
	Resource_KindAPIVersion   = Resource_Kind + "." + CRDGroupVersion.String()
	Resource_GroupVersionKind = CRDGroupVersion.WithKind(Resource_Kind)
)

func init() {
	SchemeBuilder.Register(&Resource{}, &ResourceList{})
}
