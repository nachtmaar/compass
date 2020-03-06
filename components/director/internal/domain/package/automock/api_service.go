// Code generated by mockery v1.0.0. DO NOT EDIT.

package automock

import (
	context "context"

	model "github.com/kyma-incubator/compass/components/director/internal/model"
	mock "github.com/stretchr/testify/mock"
)

// APIService is an autogenerated mock type for the APIService type
type APIService struct {
	mock.Mock
}

// GetForPackage provides a mock function with given fields: ctx, id, packageID
func (_m *APIService) GetForPackage(ctx context.Context, id string, packageID string) (*model.APIDefinition, error) {
	ret := _m.Called(ctx, id, packageID)

	var r0 *model.APIDefinition
	if rf, ok := ret.Get(0).(func(context.Context, string, string) *model.APIDefinition); ok {
		r0 = rf(ctx, id, packageID)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*model.APIDefinition)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, string, string) error); ok {
		r1 = rf(ctx, id, packageID)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// ListForPackage provides a mock function with given fields: ctx, packageID, pageSize, cursor
func (_m *APIService) ListForPackage(ctx context.Context, packageID string, pageSize int, cursor string) (*model.APIDefinitionPage, error) {
	ret := _m.Called(ctx, packageID, pageSize, cursor)

	var r0 *model.APIDefinitionPage
	if rf, ok := ret.Get(0).(func(context.Context, string, int, string) *model.APIDefinitionPage); ok {
		r0 = rf(ctx, packageID, pageSize, cursor)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*model.APIDefinitionPage)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, string, int, string) error); ok {
		r1 = rf(ctx, packageID, pageSize, cursor)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}
