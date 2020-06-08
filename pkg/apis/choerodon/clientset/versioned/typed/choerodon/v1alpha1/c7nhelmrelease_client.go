/*
Copyright 2018 Jetstack Ltd.

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
	"github.com/choerodon/choerodon-cluster-agent/pkg/apis/choerodon/clientset/versioned/scheme"
	"github.com/choerodon/choerodon-cluster-agent/pkg/apis/choerodon/v1alpha1"
	"k8s.io/client-go/rest"
)

type C7nHelmReleaseGetterV1alpha1Interface interface {
	RESTClient() rest.Interface
	C7nHelmReleaseGetter
}

// C7nHelmReleaseGetterV1alpha1Client is used to interact with features provided by the C7nHelmReleaseGetter.k8s.io group.
type C7nHelmReleaseGetterV1alpha1Client struct {
	restClient rest.Interface
}

func (c *C7nHelmReleaseGetterV1alpha1Client) C7nHelmReleases(namespace string) C7nHelmReleaseInterface {
	return newC7NHelmReleases(c, namespace)
}

// NewForConfig creates a new C7nHelmReleaseGetterV1alpha1Client for the given config.
func NewForConfig(c *rest.Config) (*C7nHelmReleaseGetterV1alpha1Client, error) {
	config := *c
	if err := setConfigDefaults(&config); err != nil {
		return nil, err
	}
	client, err := rest.RESTClientFor(&config)
	if err != nil {
		return nil, err
	}
	return &C7nHelmReleaseGetterV1alpha1Client{client}, nil
}

// NewForConfigOrDie creates a new C7nHelmReleaseGetterV1alpha1Client for the given config and
// panics if there is an error in the config.
func NewForConfigOrDie(c *rest.Config) *C7nHelmReleaseGetterV1alpha1Client {
	client, err := NewForConfig(c)
	if err != nil {
		panic(err)
	}
	return client
}

// New creates a new C7nHelmReleaseGetterV1alpha1Client for the given RESTClient.
func New(c rest.Interface) *C7nHelmReleaseGetterV1alpha1Client {
	return &C7nHelmReleaseGetterV1alpha1Client{c}
}

func setConfigDefaults(config *rest.Config) error {
	gv := v1alpha1.SchemeGroupVersion
	config.GroupVersion = &gv
	config.APIPath = "/apis"
	config.NegotiatedSerializer = scheme.Codecs.WithoutConversion()

	if config.UserAgent == "" {
		config.UserAgent = rest.DefaultKubernetesUserAgent()
	}

	return nil
}

// RESTClient returns a RESTClient that is used to communicate
// with API server by this client implementation.
func (c *C7nHelmReleaseGetterV1alpha1Client) RESTClient() rest.Interface {
	if c == nil {
		return nil
	}
	return c.restClient
}
