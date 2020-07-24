/*


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

package v1beta1

import (
	"time"

	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

const (
	// Active shows that the resource is still counting
	Active StatusDescription = "Counting"
	// Disabled shows that the resource has finished counting
	Disabled StatusDescription = "Stopped Counting"
)

// StatusDescription defines the status for the custom resource
type StatusDescription string

// EDIT THIS FILE!  THIS IS SCAFFOLDING FOR YOU TO OWN!
// NOTE: json tags are required.  Any new fields you add must have json tags for the fields to be serialized.

// CounterSpec defines the desired state of Counter
type CounterSpec struct {
	// INSERT ADDITIONAL SPEC FIELDS - desired state of cluster
	// Important: Run "make" to regenerate code after modifying this file

	// MinimumCount is the number where the counter starts counting.
	// It must be smaller than the MaximumCount
	MinimumCount int `json:"minimumCount"`

	// MaximumCount is the number where the counter ends counting
	// It must be bigger than the minimum count
	MaximumCount int `json:"maximumCount"`

	// CounterDelay is the time between the counter increases the count
	CounterDelay time.Duration `json:"counterDelay"`
}

// CounterStatus defines the observed state of Counter
type CounterStatus struct {

	// Current shows the counted value
	Current int `json:"current"`

	// Description indicates, whether the counter is still running
	Description StatusDescription `json:"description"`

	// INSERT ADDITIONAL STATUS FIELD - define observed state of cluster
	// Important: Run "make" to regenerate code after modifying this file
}

// +kubebuilder:object:root=true

// Counter is the Schema for the counters API
// +kubebuilder:printcolumn:name="MinimumCount",type=integer,JSONPath=`.spec.minimumCount`
// +kubebuilder:printcolumn:name="MaximumCount",type=integer,JSONPath=`.spec.maximumCount`
// +kubebuilder:printcolumn:name="Current",type=integer,JSONPath=`.status.current`
// +kubebuilder:printcolumn:name="Description",type=string,JSONPath=`.status.description`
// +kubebuilder:subresource:status
type Counter struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   CounterSpec   `json:"spec,omitempty"`
	Status CounterStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// CounterList contains a list of Counter
type CounterList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []Counter `json:"items"`
}

func init() {
	SchemeBuilder.Register(&Counter{}, &CounterList{})
}
