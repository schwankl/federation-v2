/*
Copyright 2018 The Kubernetes Authors.

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

package fake

import (
	v1alpha1 "github.com/kubernetes-sigs/federation-v2/pkg/apis/multiclusterdns/v1alpha1"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	labels "k8s.io/apimachinery/pkg/labels"
	schema "k8s.io/apimachinery/pkg/runtime/schema"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	testing "k8s.io/client-go/testing"
)

// FakeMultiClusterIngressDNSRecords implements MultiClusterIngressDNSRecordInterface
type FakeMultiClusterIngressDNSRecords struct {
	Fake *FakeMulticlusterdnsV1alpha1
	ns   string
}

var multiclusteringressdnsrecordsResource = schema.GroupVersionResource{Group: "multiclusterdns.federation.k8s.io", Version: "v1alpha1", Resource: "multiclusteringressdnsrecords"}

var multiclusteringressdnsrecordsKind = schema.GroupVersionKind{Group: "multiclusterdns.federation.k8s.io", Version: "v1alpha1", Kind: "MultiClusterIngressDNSRecord"}

// Get takes name of the multiClusterIngressDNSRecord, and returns the corresponding multiClusterIngressDNSRecord object, and an error if there is any.
func (c *FakeMultiClusterIngressDNSRecords) Get(name string, options v1.GetOptions) (result *v1alpha1.MultiClusterIngressDNSRecord, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewGetAction(multiclusteringressdnsrecordsResource, c.ns, name), &v1alpha1.MultiClusterIngressDNSRecord{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.MultiClusterIngressDNSRecord), err
}

// List takes label and field selectors, and returns the list of MultiClusterIngressDNSRecords that match those selectors.
func (c *FakeMultiClusterIngressDNSRecords) List(opts v1.ListOptions) (result *v1alpha1.MultiClusterIngressDNSRecordList, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewListAction(multiclusteringressdnsrecordsResource, multiclusteringressdnsrecordsKind, c.ns, opts), &v1alpha1.MultiClusterIngressDNSRecordList{})

	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &v1alpha1.MultiClusterIngressDNSRecordList{}
	for _, item := range obj.(*v1alpha1.MultiClusterIngressDNSRecordList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested multiClusterIngressDNSRecords.
func (c *FakeMultiClusterIngressDNSRecords) Watch(opts v1.ListOptions) (watch.Interface, error) {
	return c.Fake.
		InvokesWatch(testing.NewWatchAction(multiclusteringressdnsrecordsResource, c.ns, opts))

}

// Create takes the representation of a multiClusterIngressDNSRecord and creates it.  Returns the server's representation of the multiClusterIngressDNSRecord, and an error, if there is any.
func (c *FakeMultiClusterIngressDNSRecords) Create(multiClusterIngressDNSRecord *v1alpha1.MultiClusterIngressDNSRecord) (result *v1alpha1.MultiClusterIngressDNSRecord, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewCreateAction(multiclusteringressdnsrecordsResource, c.ns, multiClusterIngressDNSRecord), &v1alpha1.MultiClusterIngressDNSRecord{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.MultiClusterIngressDNSRecord), err
}

// Update takes the representation of a multiClusterIngressDNSRecord and updates it. Returns the server's representation of the multiClusterIngressDNSRecord, and an error, if there is any.
func (c *FakeMultiClusterIngressDNSRecords) Update(multiClusterIngressDNSRecord *v1alpha1.MultiClusterIngressDNSRecord) (result *v1alpha1.MultiClusterIngressDNSRecord, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateAction(multiclusteringressdnsrecordsResource, c.ns, multiClusterIngressDNSRecord), &v1alpha1.MultiClusterIngressDNSRecord{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.MultiClusterIngressDNSRecord), err
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().
func (c *FakeMultiClusterIngressDNSRecords) UpdateStatus(multiClusterIngressDNSRecord *v1alpha1.MultiClusterIngressDNSRecord) (*v1alpha1.MultiClusterIngressDNSRecord, error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateSubresourceAction(multiclusteringressdnsrecordsResource, "status", c.ns, multiClusterIngressDNSRecord), &v1alpha1.MultiClusterIngressDNSRecord{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.MultiClusterIngressDNSRecord), err
}

// Delete takes name of the multiClusterIngressDNSRecord and deletes it. Returns an error if one occurs.
func (c *FakeMultiClusterIngressDNSRecords) Delete(name string, options *v1.DeleteOptions) error {
	_, err := c.Fake.
		Invokes(testing.NewDeleteAction(multiclusteringressdnsrecordsResource, c.ns, name), &v1alpha1.MultiClusterIngressDNSRecord{})

	return err
}

// DeleteCollection deletes a collection of objects.
func (c *FakeMultiClusterIngressDNSRecords) DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error {
	action := testing.NewDeleteCollectionAction(multiclusteringressdnsrecordsResource, c.ns, listOptions)

	_, err := c.Fake.Invokes(action, &v1alpha1.MultiClusterIngressDNSRecordList{})
	return err
}

// Patch applies the patch and returns the patched multiClusterIngressDNSRecord.
func (c *FakeMultiClusterIngressDNSRecords) Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1alpha1.MultiClusterIngressDNSRecord, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewPatchSubresourceAction(multiclusteringressdnsrecordsResource, c.ns, name, data, subresources...), &v1alpha1.MultiClusterIngressDNSRecord{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.MultiClusterIngressDNSRecord), err
}
