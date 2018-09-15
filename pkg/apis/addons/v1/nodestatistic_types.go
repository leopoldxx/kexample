/*
Copyright 2018 leopoldxx.

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
	"k8s.io/api/apps/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// EDIT THIS FILE!  THIS IS SCAFFOLDING FOR YOU TO OWN!
// NOTE: json tags are required.  Any new fields you add must have json tags for the fields to be serialized.

// NodeStatisticSpec defines the desired state of NodeStatistic
type NodeStatisticSpec struct {
	// INSERT ADDITIONAL SPEC FIELDS - desired state of cluster
	Replicas int `json:"replicas"`
}

// NodeStatisticStatus defines the observed state of NodeStatistic
type NodeStatisticStatus struct {
	// INSERT ADDITIONAL STATUS FIELD - define observed state of cluster
	// Important: Run "make" to regenerate code after modifying this file
	Deployment *v1.Deployment `json:"deployment,omitempty"`
}

// +genclient
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// NodeStatistic is the Schema for the nodestatistics API
// +k8s:openapi-gen=true
type NodeStatistic struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   NodeStatisticSpec   `json:"spec,omitempty"`
	Status NodeStatisticStatus `json:"status,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// NodeStatisticList contains a list of NodeStatistic
type NodeStatisticList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []NodeStatistic `json:"items"`
}

func init() {
	SchemeBuilder.Register(&NodeStatistic{}, &NodeStatisticList{})
}
