// Code generated by mockery v1.0.0. DO NOT EDIT.

package automock

import (
	graphql "github.com/kyma-incubator/compass/components/director/pkg/graphql"
	mock "github.com/stretchr/testify/mock"

	model "github.com/kyma-incubator/compass/components/director/internal/model"
)

// DocumentConverter is an autogenerated mock type for the DocumentConverter type
type DocumentConverter struct {
	mock.Mock
}

// MultipleToGraphQL provides a mock function with given fields: in
func (_m *DocumentConverter) MultipleToGraphQL(in []*model.Document) []*graphql.Document {
	ret := _m.Called(in)

	var r0 []*graphql.Document
	if rf, ok := ret.Get(0).(func([]*model.Document) []*graphql.Document); ok {
		r0 = rf(in)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]*graphql.Document)
		}
	}

	return r0
}

// ToGraphQL provides a mock function with given fields: in
func (_m *DocumentConverter) ToGraphQL(in *model.Document) *graphql.Document {
	ret := _m.Called(in)

	var r0 *graphql.Document
	if rf, ok := ret.Get(0).(func(*model.Document) *graphql.Document); ok {
		r0 = rf(in)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*graphql.Document)
		}
	}

	return r0
}
