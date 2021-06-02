/*
Copyright 2021.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package v1alpha1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// DBaaSOperatorInventorySpec defines the desired state of DBaaSInventory
type DBaaSOperatorInventorySpec struct {
	// Important: Run "make" to regenerate code after modifying this file

	// Provider is the name of the database provider that we wish to connect with
	Provider DatabaseProvider `json:"provider"`

	DBaaSInventorySpec
}

//+kubebuilder:object:root=true
//+kubebuilder:subresource:status

// DBaaSInventory is the Schema for the dbaasinventory API
type DBaaSInventory struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   DBaaSOperatorInventorySpec `json:"spec,omitempty"`
	Status DBaaSInventoryStatus       `json:"status,omitempty"`
}

//+kubebuilder:object:root=true

// DBaaSInventoryList contains a list of DBaaSInventories
type DBaaSInventoryList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []DBaaSInventory `json:"items"`
}

func init() {
	SchemeBuilder.Register(&DBaaSInventory{}, &DBaaSInventoryList{})
}
