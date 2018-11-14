// Code generated by protoc-gen-solo-kit. DO NOT EDIT.

package v1alpha1

import (
	"github.com/solo-io/solo-kit/pkg/api/v1/clients"
	"github.com/solo-io/solo-kit/pkg/api/v1/clients/factory"
	"github.com/solo-io/solo-kit/pkg/api/v1/resources"
	"github.com/solo-io/solo-kit/pkg/errors"
)

type ServiceRoleBindingClient interface {
	BaseClient() clients.ResourceClient
	Register() error
	Read(namespace, name string, opts clients.ReadOpts) (*ServiceRoleBinding, error)
	Write(resource *ServiceRoleBinding, opts clients.WriteOpts) (*ServiceRoleBinding, error)
	Delete(namespace, name string, opts clients.DeleteOpts) error
	List(namespace string, opts clients.ListOpts) (ServiceRoleBindingList, error)
	Watch(namespace string, opts clients.WatchOpts) (<-chan ServiceRoleBindingList, <-chan error, error)
}

type serviceRoleBindingClient struct {
	rc clients.ResourceClient
}

func NewServiceRoleBindingClient(rcFactory factory.ResourceClientFactory) (ServiceRoleBindingClient, error) {
	return NewServiceRoleBindingClientWithToken(rcFactory, "")
}

func NewServiceRoleBindingClientWithToken(rcFactory factory.ResourceClientFactory, token string) (ServiceRoleBindingClient, error) {
	rc, err := rcFactory.NewResourceClient(factory.NewResourceClientParams{
		ResourceType: &ServiceRoleBinding{},
		Token:        token,
	})
	if err != nil {
		return nil, errors.Wrapf(err, "creating base ServiceRoleBinding resource client")
	}
	return &serviceRoleBindingClient{
		rc: rc,
	}, nil
}

func (client *serviceRoleBindingClient) BaseClient() clients.ResourceClient {
	return client.rc
}

func (client *serviceRoleBindingClient) Register() error {
	return client.rc.Register()
}

func (client *serviceRoleBindingClient) Read(namespace, name string, opts clients.ReadOpts) (*ServiceRoleBinding, error) {
	opts = opts.WithDefaults()
	resource, err := client.rc.Read(namespace, name, opts)
	if err != nil {
		return nil, err
	}
	return resource.(*ServiceRoleBinding), nil
}

func (client *serviceRoleBindingClient) Write(serviceRoleBinding *ServiceRoleBinding, opts clients.WriteOpts) (*ServiceRoleBinding, error) {
	opts = opts.WithDefaults()
	resource, err := client.rc.Write(serviceRoleBinding, opts)
	if err != nil {
		return nil, err
	}
	return resource.(*ServiceRoleBinding), nil
}

func (client *serviceRoleBindingClient) Delete(namespace, name string, opts clients.DeleteOpts) error {
	opts = opts.WithDefaults()
	return client.rc.Delete(namespace, name, opts)
}

func (client *serviceRoleBindingClient) List(namespace string, opts clients.ListOpts) (ServiceRoleBindingList, error) {
	opts = opts.WithDefaults()
	resourceList, err := client.rc.List(namespace, opts)
	if err != nil {
		return nil, err
	}
	return convertToServiceRoleBinding(resourceList), nil
}

func (client *serviceRoleBindingClient) Watch(namespace string, opts clients.WatchOpts) (<-chan ServiceRoleBindingList, <-chan error, error) {
	opts = opts.WithDefaults()
	resourcesChan, errs, initErr := client.rc.Watch(namespace, opts)
	if initErr != nil {
		return nil, nil, initErr
	}
	serviceRoleBindingsChan := make(chan ServiceRoleBindingList)
	go func() {
		for {
			select {
			case resourceList := <-resourcesChan:
				serviceRoleBindingsChan <- convertToServiceRoleBinding(resourceList)
			case <-opts.Ctx.Done():
				close(serviceRoleBindingsChan)
				return
			}
		}
	}()
	return serviceRoleBindingsChan, errs, nil
}

func convertToServiceRoleBinding(resources resources.ResourceList) ServiceRoleBindingList {
	var serviceRoleBindingList ServiceRoleBindingList
	for _, resource := range resources {
		serviceRoleBindingList = append(serviceRoleBindingList, resource.(*ServiceRoleBinding))
	}
	return serviceRoleBindingList
}
