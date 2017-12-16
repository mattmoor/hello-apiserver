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

// This file was autogenerated by apiregister-gen. Do not edit it manually!

package v1alpha1

import (
	"github.com/kubernetes-incubator/apiserver-builder/pkg/builders"
	"github.com/mattmoor/hello-apiserver/pkg/apis/experimental"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/apimachinery/pkg/runtime/schema"
)

var (
	experimentalBuildStorage = builders.NewApiResource( // Resource status endpoint
		experimental.InternalBuild,
		BuildSchemeFns{},
		func() runtime.Object { return &Build{} },     // Register versioned resource
		func() runtime.Object { return &BuildList{} }, // Register versioned resource list
		&BuildStrategy{builders.StorageStrategySingleton},
	)
	ApiVersion = builders.NewApiVersion("experimental.mattmoor.io", "v1alpha1").WithResources(
		experimentalBuildStorage,
		builders.NewApiResource( // Resource status endpoint
			experimental.InternalBuildStatus,
			BuildSchemeFns{},
			func() runtime.Object { return &Build{} },     // Register versioned resource
			func() runtime.Object { return &BuildList{} }, // Register versioned resource list
			&BuildStatusStrategy{builders.StatusStorageStrategySingleton},
		))

	// Required by code generated by go2idl
	AddToScheme        = ApiVersion.SchemaBuilder.AddToScheme
	SchemeBuilder      = ApiVersion.SchemaBuilder
	localSchemeBuilder = &SchemeBuilder
	SchemeGroupVersion = ApiVersion.GroupVersion
)

// Required by code generated by go2idl
// Kind takes an unqualified kind and returns a Group qualified GroupKind
func Kind(kind string) schema.GroupKind {
	return SchemeGroupVersion.WithKind(kind).GroupKind()
}

// Required by code generated by go2idl
// Resource takes an unqualified resource and returns a Group qualified GroupResource
func Resource(resource string) schema.GroupResource {
	return SchemeGroupVersion.WithResource(resource).GroupResource()
}

//
// Build Functions and Structs
//
// +k8s:deepcopy-gen=false
type BuildSchemeFns struct {
	builders.DefaultSchemeFns
}

// +k8s:deepcopy-gen=false
type BuildStrategy struct {
	builders.DefaultStorageStrategy
}

// +k8s:deepcopy-gen=false
type BuildStatusStrategy struct {
	builders.DefaultStatusStorageStrategy
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

type BuildList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []Build `json:"items"`
}
