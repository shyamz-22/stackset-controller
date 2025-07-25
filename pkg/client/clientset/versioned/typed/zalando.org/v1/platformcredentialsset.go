/*
Copyright 2025 The Kubernetes Authors.

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

// Code generated by client-gen. DO NOT EDIT.

package v1

import (
	context "context"

	zalandoorgv1 "github.com/zalando-incubator/stackset-controller/pkg/apis/zalando.org/v1"
	scheme "github.com/zalando-incubator/stackset-controller/pkg/client/clientset/versioned/scheme"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	gentype "k8s.io/client-go/gentype"
)

// PlatformCredentialsSetsGetter has a method to return a PlatformCredentialsSetInterface.
// A group's client should implement this interface.
type PlatformCredentialsSetsGetter interface {
	PlatformCredentialsSets(namespace string) PlatformCredentialsSetInterface
}

// PlatformCredentialsSetInterface has methods to work with PlatformCredentialsSet resources.
type PlatformCredentialsSetInterface interface {
	Create(ctx context.Context, platformCredentialsSet *zalandoorgv1.PlatformCredentialsSet, opts metav1.CreateOptions) (*zalandoorgv1.PlatformCredentialsSet, error)
	Update(ctx context.Context, platformCredentialsSet *zalandoorgv1.PlatformCredentialsSet, opts metav1.UpdateOptions) (*zalandoorgv1.PlatformCredentialsSet, error)
	// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().
	UpdateStatus(ctx context.Context, platformCredentialsSet *zalandoorgv1.PlatformCredentialsSet, opts metav1.UpdateOptions) (*zalandoorgv1.PlatformCredentialsSet, error)
	Delete(ctx context.Context, name string, opts metav1.DeleteOptions) error
	DeleteCollection(ctx context.Context, opts metav1.DeleteOptions, listOpts metav1.ListOptions) error
	Get(ctx context.Context, name string, opts metav1.GetOptions) (*zalandoorgv1.PlatformCredentialsSet, error)
	List(ctx context.Context, opts metav1.ListOptions) (*zalandoorgv1.PlatformCredentialsSetList, error)
	Watch(ctx context.Context, opts metav1.ListOptions) (watch.Interface, error)
	Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts metav1.PatchOptions, subresources ...string) (result *zalandoorgv1.PlatformCredentialsSet, err error)
	PlatformCredentialsSetExpansion
}

// platformCredentialsSets implements PlatformCredentialsSetInterface
type platformCredentialsSets struct {
	*gentype.ClientWithList[*zalandoorgv1.PlatformCredentialsSet, *zalandoorgv1.PlatformCredentialsSetList]
}

// newPlatformCredentialsSets returns a PlatformCredentialsSets
func newPlatformCredentialsSets(c *ZalandoV1Client, namespace string) *platformCredentialsSets {
	return &platformCredentialsSets{
		gentype.NewClientWithList[*zalandoorgv1.PlatformCredentialsSet, *zalandoorgv1.PlatformCredentialsSetList](
			"platformcredentialssets",
			c.RESTClient(),
			scheme.ParameterCodec,
			namespace,
			func() *zalandoorgv1.PlatformCredentialsSet { return &zalandoorgv1.PlatformCredentialsSet{} },
			func() *zalandoorgv1.PlatformCredentialsSetList { return &zalandoorgv1.PlatformCredentialsSetList{} },
		),
	}
}
