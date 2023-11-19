/*
Copyright 2023.

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

type ServiceSpec struct {
	Name       string `json:"name"`
	Protocol   string `json:"protocol"`
	Port       int32  `json:"port"`
	TargetPort int32  `json:"targetPort"`
	NodePort   int32  `json:"nodePort"`
	Type       string `json:"type"`
}

// EDIT THIS FILE!  THIS IS SCAFFOLDING FOR YOU TO OWN!
// NOTE: json tags are required.  Any new fields you add must have json tags for the fields to be serialized.

// MyPythonAppSpec defines the desired state of MyPythonApp
type MyPythonAppSpec struct {
	// INSERT ADDITIONAL SPEC FIELDS - desired state of cluster
	// Important: Run "make" to regenerate code after modifying this file

	// Foo is an example field of MyPythonApp. Edit mypythonapp_types.go to remove/update
	// These are all spec components for the CR
	CompanyName            string      `json:"companyName,omitempty"`
	ApplicationDescription string      `json:"applicationDescription,omitempty"`
	AppContainerName       string      `json:"appContainerName"`
	AppImage               string      `json:"appImage"`
	AppPort                int32       `json:"appPort"`
	MonitorContainerName   string      `json:"monitorContainerName"`
	MonitorImage           string      `json:"monitorImage"`
	MonitorCommand         string      `json:"monitorCommand"`
	Size                   int32       `json:"size"`
	Service                ServiceSpec `json:"service"`
}

// MyPythonAppStatus defines the observed state of MyPythonApp
type MyPythonAppStatus struct {
	// INSERT ADDITIONAL STATUS FIELD - define observed state of cluster
	// Important: Run "make" to regenerate code after modifying this file
	PodList []string `json:"podList"`
	// the above is used to store the list of podnames or identifiers that are associated with custom resources
	// the actual values in the list will be determined by the controller that managed the MyPythonApp resource.
	// When the MyPythonApp resource is read by the kubernetes client  or controller, it can use the PodList field to determine which Pods are currently associated with the resource and take appropriate action based on this information.

}

//+kubebuilder:object:root=true // marks the Go structure as the root object of the Kubernetes custom resource ***
//+kubebuilder:subresource:status // enables subresource support for the custom resource
// the subresource allows the user to view and updated the status of the custom resource separately from the
// we are not updating anything in the structure

// the above 2 are markers used to generate Kubernetes API objects and enable subresource support spec

// MyPythonApp is the Schema for the mypythonapps API
type MyPythonApp struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   MyPythonAppSpec   `json:"spec,omitempty"`
	Status MyPythonAppStatus `json:"status,omitempty"`
}

//+kubebuilder:object:root=true

// MyPythonAppList contains a list of MyPythonApp
// Not updating anything here, leave as is
type MyPythonAppList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []MyPythonApp `json:"items"`
}

// This init function is used to register the MyPythonApp and MyPythonAppList types with the K8s API machinery, allowing them to be recognized as the valied K8s resources.
// Not updating anything here, leave as is
func init() {
	SchemeBuilder.Register(&MyPythonApp{}, &MyPythonAppList{})
}
