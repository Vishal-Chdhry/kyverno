/*
Copyright 2020 The Kubernetes authors.

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
	policyreportv1alpha2 "github.com/kyverno/kyverno/api/policyreport/v1alpha2"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

type EphemeralReportSpec struct {
	// Owner is a reference to the report owner (e.g. a Deployment, Namespace, or Node)
	Owner metav1.OwnerReference `json:"owner"`

	// PolicyReportSummary provides a summary of results
	// +optional
	Summary policyreportv1alpha2.PolicyReportSummary `json:"summary,omitempty"`

	// PolicyReportResult provides result details
	// +optional
	Results []policyreportv1alpha2.PolicyReportResult `json:"results,omitempty"`
}

// +genclient
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:object:root=true
// +kubebuilder:storageversion
// +kubebuilder:resource:shortName=admr,categories=kyverno
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:printcolumn:name="PASS",type=integer,JSONPath=".spec.summary.pass"
// +kubebuilder:printcolumn:name="FAIL",type=integer,JSONPath=".spec.summary.fail"
// +kubebuilder:printcolumn:name="WARN",type=integer,JSONPath=".spec.summary.warn"
// +kubebuilder:printcolumn:name="ERROR",type=integer,JSONPath=".spec.summary.error"
// +kubebuilder:printcolumn:name="SKIP",type=integer,JSONPath=".spec.summary.skip"
// +kubebuilder:printcolumn:name="GVR",type=string,JSONPath=".metadata.labels['audit\\.kyverno\\.io/resource\\.gvr']"
// +kubebuilder:printcolumn:name="REF",type=string,JSONPath=".metadata.labels['audit\\.kyverno\\.io/resource\\.name']"
// +kubebuilder:printcolumn:name="AGGREGATE",type=string,JSONPath=".metadata.labels['audit\\.kyverno\\.io/report\\.aggregate']",priority=1

// EphemeralReport is the Schema for the EphemeralReports API
type EphemeralReport struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              EphemeralReportSpec `json:"spec"`
}

func (r *EphemeralReport) GetResults() []policyreportv1alpha2.PolicyReportResult {
	return r.Spec.Results
}

func (r *EphemeralReport) SetResults(results []policyreportv1alpha2.PolicyReportResult) {
	r.Spec.Results = results
}

func (r *EphemeralReport) SetSummary(summary policyreportv1alpha2.PolicyReportSummary) {
	r.Spec.Summary = summary
}

// +genclient
// +genclient:nonNamespaced
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:object:root=true
// +kubebuilder:storageversion
// +kubebuilder:resource:scope=Cluster,shortName=cadmr,categories=kyverno
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:printcolumn:name="PASS",type=integer,JSONPath=".spec.summary.pass"
// +kubebuilder:printcolumn:name="FAIL",type=integer,JSONPath=".spec.summary.fail"
// +kubebuilder:printcolumn:name="WARN",type=integer,JSONPath=".spec.summary.warn"
// +kubebuilder:printcolumn:name="ERROR",type=integer,JSONPath=".spec.summary.error"
// +kubebuilder:printcolumn:name="SKIP",type=integer,JSONPath=".spec.summary.skip"
// +kubebuilder:printcolumn:name="GVR",type=string,JSONPath=".metadata.labels['audit\\.kyverno\\.io/resource\\.gvr']"
// +kubebuilder:printcolumn:name="REF",type=string,JSONPath=".metadata.labels['audit\\.kyverno\\.io/resource\\.name']"
// +kubebuilder:printcolumn:name="AGGREGATE",type=string,JSONPath=".metadata.labels['audit\\.kyverno\\.io/report\\.aggregate']",priority=1

// ClusterEphemeralReport is the Schema for the ClusterEphemeralReports API
type ClusterEphemeralReport struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              EphemeralReportSpec `json:"spec"`
}

func (r *ClusterEphemeralReport) GetResults() []policyreportv1alpha2.PolicyReportResult {
	return r.Spec.Results
}

func (r *ClusterEphemeralReport) SetResults(results []policyreportv1alpha2.PolicyReportResult) {
	r.Spec.Results = results
}

func (r *ClusterEphemeralReport) SetSummary(summary policyreportv1alpha2.PolicyReportSummary) {
	r.Spec.Summary = summary
}

// +kubebuilder:object:root=true
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// EphemeralReportList contains a list of EphemeralReport
type EphemeralReportList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []EphemeralReport `json:"items"`
}

// +kubebuilder:object:root=true
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// ClusterEphemeralReportList contains a list of ClusterEphemeralReport
type ClusterEphemeralReportList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []ClusterEphemeralReport `json:"items"`
}
