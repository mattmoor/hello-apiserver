
/*
Copyright 2017 The Kubernetes Authors.

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
	"log"

	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/apiserver/pkg/endpoints/request"

	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/util/validation/field"

	"github.com/mattmoor/hello-apiserver/pkg/apis/experimental"
)

// +genclient=true
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// Build
// +k8s:openapi-gen=true
// +resource:path=builds,strategy=BuildStrategy
type Build struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   BuildSpec   `json:"spec,omitempty"`
	Status BuildStatus `json:"status,omitempty"`
}

// BuildSpec defines the desired state of Build
type BuildSpec struct {
    Source SourceSpec `json:"source,omitempty"`

    Steps []StepSpec `json:"steps,omitempty"`

    Images []string `json:"images,omitempty"`
}

// StepSpec defines a step of the Build
type StepSpec struct {
    Name string `json:"name"`
    Env []string `json:"env,omitempty"`
    Dir string `json:"dir,omitempty"`
    Entrypoint string `json:"entrypoint,omitempty"`
    Args []string `json:"args,omitempty"`
    Volumes []VolumeSpec `json:"volumes,omitempty"`
}

// VolumeSpec defines a step of the Build
type VolumeSpec struct {
    Name string `json:"name,omitempty"`
    Path string `json:"path,omitempty"`
}

// SourceSpec defines the input to the Build
type SourceSpec struct {
    StorageSource StorageSourceSpec `json:"storage_source,omitempty"`
    RepoSource RepoSourceSpec `json:"repo_source,omitempty"`
}

// StorageSourceSpec defines the input to the Build
type StorageSourceSpec struct {
    Bucket string `json:"bucket,omitempty"`
    Object string `json:"object,omitempty"`
    Generation int `json:"generation,omitempty"`
}

// RepoSourceSpec defines the input to the Build
type RepoSourceSpec struct {
    ProjectId string `json:"project_id,omitempty"`
    RepoName string `json:"repo_name,omitempty"`
    BranchName string `json:"branch_name,omitempty"`
    TagName string `json:"tag_name,omitempty"`
    CommitSHA string `json:"commit_sha,omitempty"`
    Dir string `json:"dir,omitempty"`
}

// BuildStatus defines the observed state of Build
type BuildStatus struct {
}

// Validate checks that an instance of Build is well formed
func (BuildStrategy) Validate(ctx request.Context, obj runtime.Object) field.ErrorList {
	o := obj.(*experimental.Build)
	log.Printf("Validating fields for Build %s\n", o.Name)
	errors := field.ErrorList{}
	// perform validation here and add to errors using field.Invalid
	return errors
}

// DefaultingFunction sets default Build field values
func (BuildSchemeFns) DefaultingFunction(o interface{}) {
	obj := o.(*Build)
	// set default field values here
	log.Printf("Defaulting fields for Build %s\n", obj.Name)
}
