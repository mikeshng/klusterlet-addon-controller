// IBM Confidential
// OCO Source Materials
// (C) Copyright IBM Corporation 2020 All Rights Reserved
// The source code for this program is not published or otherwise divested of its trade secrets, irrespective of what has been deposited with the U.S. Copyright Office.

package v1beta1

import (
	"github.com/open-cluster-management/endpoint-operator/pkg/image"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// EDIT THIS FILE!  THIS IS SCAFFOLDING FOR YOU TO OWN!
// NOTE: json tags are required.  Any new fields you add must have json tags for the fields to be serialized.

// CertPolicyControllerSpec defines the desired state of CertPolicyController
type CertPolicyControllerSpec struct {
	// INSERT ADDITIONAL SPEC FIELDS - desired state of cluster
	// Important: Run "operator-sdk generate k8s" to regenerate code after modifying this file
	// Add custom validation using kubebuilder tags: https://book-v1.book.kubebuilder.io/beyond_basics/generating_crd.html

	// +kubebuilder:validation:MinLength=1
	FullNameOverride string `json:"fullnameOverride"`

	// +kubebuilder:validation:MinLength=1
	ClusterName string `json:"clusterName"`

	// +kubebuilder:validation:MinLength=1
	ClusterNamespace string `json:"clusterNamespace"`

	// +kubebuilder:validation:MinLength=1
	ConnectionManager string `json:"connectionManager"`

	Image           image.Image `json:"image,omitempty"`
	ImagePullSecret string      `json:"imagePullSecret,omitempty"`
}

// CertPolicyControllerStatus defines the observed state of CertPolicyController
type CertPolicyControllerStatus struct {
	// INSERT ADDITIONAL STATUS FIELD - define observed state of cluster
	// Important: Run "operator-sdk generate k8s" to regenerate code after modifying this file
	// Add custom validation using kubebuilder tags: https://book-v1.book.kubebuilder.io/beyond_basics/generating_crd.html
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// CertPolicyController is the Schema for the certpolicycontrollercontroller API
// +kubebuilder:subresource:status
// +kubebuilder:resource:path=certpoliciescontroller,scope=Namespaced
type CertPolicyController struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   CertPolicyControllerSpec   `json:"spec,omitempty"`
	Status CertPolicyControllerStatus `json:"status,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// CertPolicyControllerList contains a list of CertPolicyController
type CertPolicyControllerList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []CertPolicyController `json:"items"`
}

func init() {
	SchemeBuilder.Register(&CertPolicyController{}, &CertPolicyControllerList{})
}
