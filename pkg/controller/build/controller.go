
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


package build

import (
	"log"

	"github.com/kubernetes-incubator/apiserver-builder/pkg/builders"
	"k8s.io/client-go/rest"

	"github.com/mattmoor/hello-apiserver/pkg/apis/experimental/v1alpha1"
	"github.com/mattmoor/hello-apiserver/pkg/controller/sharedinformers"
	listers "github.com/mattmoor/hello-apiserver/pkg/client/listers_generated/experimental/v1alpha1"
)

// +controller:group=experimental,version=v1alpha1,kind=Build,resource=builds
type BuildControllerImpl struct {
	builders.DefaultControllerFns

	// lister indexes properties about Build
	lister listers.BuildLister
}

// Init initializes the controller and is called by the generated code
// Registers eventhandlers to enqueue events
// config - client configuration for talking to the apiserver
// si - informer factory shared across all controllers for listening to events and indexing resource properties
// queue - message queue for handling new events.  unique to this controller.
func (c *BuildControllerImpl) Init(
	config *rest.Config,
	si *sharedinformers.SharedInformers,
    reconcileKey func(key string) error) {

	// Set the informer and lister for subscribing to events and indexing builds labels
	c.lister = si.Factory.Experimental().V1alpha1().Builds().Lister()
}

// Reconcile handles enqueued messages
func (c *BuildControllerImpl) Reconcile(u *v1alpha1.Build) error {
	// Implement controller logic here
	log.Printf("Running reconcile Build for %s\n", u.Name)
	return nil
}

func (c *BuildControllerImpl) Get(namespace, name string) (*v1alpha1.Build, error) {
	return c.lister.Builds(namespace).Get(name)
}
