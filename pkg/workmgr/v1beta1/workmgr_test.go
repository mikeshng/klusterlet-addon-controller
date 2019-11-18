// Package v1beta1 of workmgrs Defines the Reconciliation logic and required setup for component operator.
// IBM Confidential
// OCO Source Materials
// (C) Copyright IBM Corporation 2019 All Rights Reserved
// The source code for this program is not published or otherwise divested of its trade secrets, irrespective of what has been deposited with the U.S. Copyright Office.
package v1beta1

import (
	"context"
	"testing"

	"github.com/stretchr/testify/assert"
	extensionsv1beta1 "k8s.io/api/extensions/v1beta1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/apimachinery/pkg/types"
	"k8s.io/client-go/kubernetes/scheme"
	"sigs.k8s.io/controller-runtime/pkg/client/fake"

	multicloudv1beta1 "github.ibm.com/IBMPrivateCloud/ibm-klusterlet-operator/pkg/apis/multicloud/v1beta1"
)

var (
	namespace = "multicluster-endpoint"
)

func newTestDeployment(name string) *extensionsv1beta1.Deployment {
	deployment := &extensionsv1beta1.Deployment{
		ObjectMeta: metav1.ObjectMeta{
			Name:      name,
			Namespace: namespace,
		},
		Status: extensionsv1beta1.DeploymentStatus{
			Conditions: []extensionsv1beta1.DeploymentCondition{extensionsv1beta1.DeploymentCondition{
				Type:   "Available",
				Status: "True",
			}},
		},
	}
	return deployment
}

func newInstance() *multicloudv1beta1.Endpoint {
	instance := &multicloudv1beta1.Endpoint{
		ObjectMeta: metav1.ObjectMeta{
			Name:      "endpoint",
			Namespace: namespace,
		},
		Spec: multicloudv1beta1.EndpointSpec{
			Version: "3.2.1",
		},
	}
	return instance
}

func newInstanceInDeletion() *multicloudv1beta1.Endpoint {
	instance := &multicloudv1beta1.Endpoint{
		ObjectMeta: metav1.ObjectMeta{
			Name:              "endpoint",
			Namespace:         namespace,
			DeletionTimestamp: &metav1.Time{},
			Finalizers:        []string{"endpoint-workmgr"},
		},
		Spec: multicloudv1beta1.EndpointSpec{
			Version:         "3.2.1",
			ImagePullSecret: "image-pull",
		},
	}
	return instance
}

func newWorkmgr() *multicloudv1beta1.WorkManager {
	workmgr := &multicloudv1beta1.WorkManager{
		ObjectMeta: metav1.ObjectMeta{
			Name:      "endpoint-workmgr",
			Namespace: namespace,
		},
	}
	return workmgr
}

func TestCreateReconcile(t *testing.T) {
	instance := newInstance()
	workmgr := &multicloudv1beta1.WorkManager{}

	scheme := scheme.Scheme
	scheme.AddKnownTypes(multicloudv1beta1.SchemeGroupVersion, instance)
	scheme.AddKnownTypes(multicloudv1beta1.SchemeGroupVersion, workmgr)

	objs := []runtime.Object{}
	cl := fake.NewFakeClient(objs...)

	res, err := Reconcile(instance, cl, scheme)
	assert.NoError(t, err, "CREATE workmgr reconcile should success")
	assert.False(t, res, "CREATE workmgr reconcile should return false")

	err = cl.Get(context.TODO(), types.NamespacedName{Name: instance.ObjectMeta.Name + "-workmgr", Namespace: namespace}, workmgr)
	assert.NoError(t, err, "workmgr CR should be created")
}

func TestFinalizeReconcile(t *testing.T) {
	instance := newInstanceInDeletion()
	workmgr := &multicloudv1beta1.WorkManager{}

	scheme := scheme.Scheme
	scheme.AddKnownTypes(multicloudv1beta1.SchemeGroupVersion, instance)
	scheme.AddKnownTypes(multicloudv1beta1.SchemeGroupVersion, workmgr)

	objs := []runtime.Object{}
	cl := fake.NewFakeClient(objs...)

	res, err := Reconcile(instance, cl, scheme)
	assert.NoError(t, err, "FINALIZE workmgr reconcile should success")
	assert.False(t, res, "FINALIZE workmgr reconcile should return false")

	assert.Empty(t, instance.Finalizers, "Finalizer shoule be removed")
}

func TestUpdateReconcile(t *testing.T) {
	instance := newInstance()
	workmgr := newWorkmgr()

	scheme := scheme.Scheme
	scheme.AddKnownTypes(multicloudv1beta1.SchemeGroupVersion, instance)
	scheme.AddKnownTypes(multicloudv1beta1.SchemeGroupVersion, workmgr)

	objs := []runtime.Object{workmgr}
	cl := fake.NewFakeClient(objs...)

	res, err := Reconcile(instance, cl, scheme)
	assert.NoError(t, err, "UPDATE workmgr reconcile should success")
	assert.False(t, res, "UPDATE workmgr reconcile should return false")

	err = cl.Get(context.TODO(), types.NamespacedName{Name: instance.ObjectMeta.Name + "-workmgr", Namespace: namespace}, workmgr)
	assert.NoError(t, err, "workmgr CR should be created")

	assert.Equal(t, workmgr.Spec.FullNameOverride, instance.Name+"-workmgr", "workmgr CR FullNameOverride should be updated")
	assert.Equal(t, workmgr.Spec.ClusterNamespace, instance.Spec.ClusterNamespace, "workmgr CR ClusterNamespace should be updated")
	assert.Equal(t, workmgr.Spec.ImagePullSecret, instance.Spec.ImagePullSecret, "workmgr CR ImagePullSecret should be updated")
}

func TestDeleteReconcile(t *testing.T) {
	instance := newInstanceInDeletion()
	workmgr := newWorkmgr()

	scheme := scheme.Scheme
	scheme.AddKnownTypes(multicloudv1beta1.SchemeGroupVersion, instance)
	scheme.AddKnownTypes(multicloudv1beta1.SchemeGroupVersion, workmgr)

	objs := []runtime.Object{workmgr}
	cl := fake.NewFakeClient(objs...)

	res, err := Reconcile(instance, cl, scheme)
	assert.NoError(t, err, "DELETE workmgr reconcile should success")
	assert.Equal(t, res, true, "DELETE workmgr reconcile should return true")

	err = cl.Get(context.TODO(), types.NamespacedName{Name: instance.ObjectMeta.Name + "-workmgr", Namespace: namespace}, workmgr)
	assert.Error(t, err, "workmgr CR should be deleted")
}

func TestIsReady(t *testing.T) {
	instance := newInstanceInDeletion()
	deployment := newTestDeployment(instance.ObjectMeta.Name + "-workmgr")

	objs := []runtime.Object{deployment}
	cl := fake.NewFakeClient(objs...)

	res, err := IsReady(instance, cl)
	assert.NoError(t, err, "IsReady should success")
	assert.Equal(t, res, true, "IsReady should return true")
}

func TestIsNotReady(t *testing.T) {
	instance := newInstanceInDeletion()

	objs := []runtime.Object{}
	cl := fake.NewFakeClient(objs...)

	res, err := IsReady(instance, cl)
	assert.NoError(t, err, "IsReady should success")
	assert.False(t, res, "IsReady should return false")
}