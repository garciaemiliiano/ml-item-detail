// Code generated by mockery v2.52.3. DO NOT EDIT.

package mocks

import (
	context "context"
	entitiesitems "item-detail-api/src/core/entities/items"
	items "item-detail-api/src/core/usecases/items"

	mock "github.com/stretchr/testify/mock"
)

// GetItem is an autogenerated mock type for the GetItem type
type GetItem struct {
	mock.Mock
}

// Execute provides a mock function with given fields: ctx, config
func (_m *GetItem) Execute(ctx context.Context, config items.GetItemParams) (entitiesitems.Item, error) {
	ret := _m.Called(ctx, config)

	if len(ret) == 0 {
		panic("no return value specified for Execute")
	}

	var r0 entitiesitems.Item
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, items.GetItemParams) (entitiesitems.Item, error)); ok {
		return rf(ctx, config)
	}
	if rf, ok := ret.Get(0).(func(context.Context, items.GetItemParams) entitiesitems.Item); ok {
		r0 = rf(ctx, config)
	} else {
		r0 = ret.Get(0).(entitiesitems.Item)
	}

	if rf, ok := ret.Get(1).(func(context.Context, items.GetItemParams) error); ok {
		r1 = rf(ctx, config)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// NewGetItem creates a new instance of GetItem. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewGetItem(t interface {
	mock.TestingT
	Cleanup(func())
}) *GetItem {
	mock := &GetItem{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
