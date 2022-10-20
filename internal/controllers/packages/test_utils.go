package packages

import (
	"context"

	"github.com/stretchr/testify/mock"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	ctrl "sigs.k8s.io/controller-runtime"
	"sigs.k8s.io/controller-runtime/pkg/client"
	"sigs.k8s.io/controller-runtime/pkg/handler"
)

type errWithStatusError struct {
	errMsg          string
	errStatusReason metav1.StatusReason
}

func (err errWithStatusError) Error() string {
	return err.errMsg
}
func (err errWithStatusError) Status() metav1.Status {
	return metav1.Status{Reason: err.errStatusReason}
}

var _ ownerStrategy = (*mockOwnerStrategy)(nil)

type mockOwnerStrategy struct {
	mock.Mock
}

func (s *mockOwnerStrategy) EnqueueRequestForOwner(
	ownerType client.Object, isController bool,
) handler.EventHandler {
	return nil
}

func (s *mockOwnerStrategy) SetControllerReference(owner, obj metav1.Object) error {
	args := s.Called(owner, obj)
	return args.Error(0)
}

func (s *mockOwnerStrategy) IsOwner(owner, obj metav1.Object) bool {
	return false
}

func (s *mockOwnerStrategy) ReleaseController(obj metav1.Object) {
}

type objectDeploymentStatusReconcilerMock struct {
	mock.Mock
}

type jobReconcilerMock struct {
	mock.Mock
}

func (r *objectDeploymentStatusReconcilerMock) Reconcile(
	ctx context.Context, pkg genericPackage,
) (res ctrl.Result, err error) {
	args := r.Called(ctx, pkg)
	return args.Get(0).(ctrl.Result), args.Error(1)
}

func (r *jobReconcilerMock) Reconcile(
	ctx context.Context, pkg genericPackage,
) (res ctrl.Result, err error) {
	args := r.Called(ctx, pkg)
	return args.Get(0).(ctrl.Result), args.Error(1)
}

func sortInsensitiveStringSlicesMatch(slice1, slice2 []string) bool {
	if len(slice1) != len(slice2) {
		return false
	}
	map1 := map[string]int{}
	map2 := map[string]int{}

	for _, s := range slice1 {
		map1[s]++
	}
	for _, s := range slice2 {
		map2[s]++
	}

	if len(map1) != len(map2) {
		return false
	}

	for k := range map1 {
		if map1[k] != map2[k] {
			return false
		}
	}
	return true
}
