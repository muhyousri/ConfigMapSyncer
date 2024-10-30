/*
Copyright 2024.

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

package v1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// EDIT THIS FILE!  THIS IS SCAFFOLDING FOR YOU TO OWN!
// NOTE: json tags are required.  Any new fields you add must have json tags for the fields to be serialized.

// SyncConfigSpec defines the desired state of SyncConfig.
type SyncConfigSpec struct {
	// +kubebuilder:validation:Description=The data to be synced in the ConfigMap
	Data map[string]string `json:"data"`
	// +kubebuilder:validation:Description=This is the target Namespace
	TargetNs string `json:"targetNs"`
}

// SyncConfigStatus defines the observed state of SyncConfig.
type SyncConfigStatus struct {
	// INSERT ADDITIONAL STATUS FIELD - define observed state of cluster
	// Important: Run "make" to regenerate code after modifying this file
	ConfigMapName string `json:"configMapName"`
	State         string `json:"state"`
}

// SyncConfig is the Schema for the syncconfigs API.
// +kubebuilder:resource:path=syncconfigs,scope=Namespaced,shortName=sc
// +kubebuilder:object:root=true
// +kubebuilder:subresource:status
type SyncConfig struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   SyncConfigSpec   `json:"spec,omitempty"`
	Status SyncConfigStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// SyncConfigList contains a list of SyncConfig.
type SyncConfigList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []SyncConfig `json:"items"`
}

func init() {
	SchemeBuilder.Register(&SyncConfig{}, &SyncConfigList{})
}
