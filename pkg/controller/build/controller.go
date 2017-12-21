
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
	"fmt"
	"log"
	"net/http"
        "encoding/json"
        "time"

	"github.com/kubernetes-incubator/apiserver-builder/pkg/builders"
	"k8s.io/client-go/rest"

	// These are not vendored yet, or modeled by glide.
	// You can fetch them into your GOPATH by "go get"ing the cloudbuild dep.
	"golang.org/x/oauth2"
	"golang.org/x/oauth2/google"
	"google.golang.org/api/cloudbuild/v1"
        "cloud.google.com/go/compute/metadata"

        metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"

	clientset "github.com/mattmoor/hello-apiserver/pkg/client/clientset_generated/clientset/typed/experimental/v1alpha1"
        "github.com/mattmoor/hello-apiserver/pkg/apis/experimental/v1alpha1"
	"github.com/mattmoor/hello-apiserver/pkg/controller/sharedinformers"
	listers "github.com/mattmoor/hello-apiserver/pkg/client/listers_generated/experimental/v1alpha1"
)

// +controller:group=experimental,version=v1alpha1,kind=Build,resource=builds
type BuildControllerImpl struct {
	builders.DefaultControllerFns

	// lister indexes properties about Build
	lister listers.BuildLister

        // experimentalClient allows us to manipulate our status.
        experimentalClient *clientset.ExperimentalV1alpha1Client

	cloudbuild *cloudbuild.Service

        project string
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
        
        // Create a client that we can use for manipulating experimental objects
        experimentalClient, err := clientset.NewForConfig(config)
        if err != nil {
                log.Printf("Could not create experimental clientset: %v", err)
                panic("Could not create experimental clientset")
        }
        c.experimentalClient = experimentalClient
        
	client := &http.Client{
            Transport: &oauth2.Transport{
                // If no account is specified, "default" is used.
                Source: google.ComputeTokenSource(""),
            },
	}

	svc, err := cloudbuild.New(client)
	if err != nil {
		panic(fmt.Sprintf("Unable to initialize cloudbuild client: %v", err))
	}
	c.cloudbuild = svc

        project, err := metadata.ProjectID()
        if err != nil {
                panic(fmt.Sprintf("Unable to determine project-id, are you running on GCE? error: %v", err))
        }
        c.project = project
}

// Reconcile handles enqueued messages
func (c *BuildControllerImpl) Reconcile(u *v1alpha1.Build) error {
	if u.Status.Operation != "" {
		return c.waitUntilDone(u)
	}
	err := c.queueBuild(u)
        if err != nil {
                log.Printf("Unexpected error queuing build: %v", err)
                return err
        }
        return nil
}

func (c *BuildControllerImpl) updateStatus(u *v1alpha1.Build) (*v1alpha1.Build, error) {
    buildClient := c.experimentalClient.Builds(u.Namespace)
    newu, err := buildClient.Get(u.Name, metav1.GetOptions{})
    if err != nil {
        return nil, err
    }
    newu.Status = u.Status
    return buildClient.UpdateStatus(newu)
}

func (c *BuildControllerImpl) updateStatusFromOperation(u *v1alpha1.Build, op *cloudbuild.Operation) error {
    u.Status.Operation = op.Name
    u.Status.Done = op.Done
    if op.Error != nil {
            u.Status.ErrorMessage = op.Error.Message
    }
    _, err := c.updateStatus(u)
    return err
}


func (c *BuildControllerImpl) waitUntilDone(u *v1alpha1.Build) error {
    if u.Status.Done {
            return nil
    }

    go func() {
        for {
                operation, err := c.cloudbuild.Operations.Get(u.Status.Operation).Do()
                if err != nil {
                        panic(fmt.Sprintf("Error fetching build status: %v", err))
                }
                if operation.Done {
                        if err := c.updateStatusFromOperation(u, operation); err != nil {
                                panic(fmt.Sprintf("Error recording final build status: %v", err))
                        }
                }
                time.Sleep(1 * time.Second)
        }
    }()
    return nil
}

func (c *BuildControllerImpl) queueBuild(u *v1alpha1.Build) error {
    b, err := json.Marshal(u.Spec)
    if err != nil {
            return err
    }
    var build cloudbuild.Build
    if err := json.Unmarshal(b, &build); err != nil {
            return err
    }
    operation, err := c.cloudbuild.Projects.Builds.Create(c.project, &build).Do()
    if err != nil {
            return err
    }
    return c.updateStatusFromOperation(u, operation)
}

func (c *BuildControllerImpl) Get(namespace, name string) (*v1alpha1.Build, error) {
	return c.lister.Builds(namespace).Get(name)
}
