// Code generated by mockery v2.42.2. DO NOT EDIT.

package proto_gen

import (
	context "context"

	grpc "google.golang.org/grpc"

	internal_proto "product_order/internal_proto"

	mock "github.com/stretchr/testify/mock"
)

// OrderServiceClient is an autogenerated mock type for the OrderServiceClient type
type OrderServiceClient struct {
	mock.Mock
}

type OrderServiceClient_Expecter struct {
	mock *mock.Mock
}

func (_m *OrderServiceClient) EXPECT() *OrderServiceClient_Expecter {
	return &OrderServiceClient_Expecter{mock: &_m.Mock}
}

// CreateOrder provides a mock function with given fields: ctx, in, opts
func (_m *OrderServiceClient) CreateOrder(ctx context.Context, in *internal_proto.CreateOrderRequest, opts ...grpc.CallOption) (*internal_proto.CreateOrderResponse, error) {
	_va := make([]interface{}, len(opts))
	for _i := range opts {
		_va[_i] = opts[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, ctx, in)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	if len(ret) == 0 {
		panic("no return value specified for CreateOrder")
	}

	var r0 *internal_proto.CreateOrderResponse
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, *internal_proto.CreateOrderRequest, ...grpc.CallOption) (*internal_proto.CreateOrderResponse, error)); ok {
		return rf(ctx, in, opts...)
	}
	if rf, ok := ret.Get(0).(func(context.Context, *internal_proto.CreateOrderRequest, ...grpc.CallOption) *internal_proto.CreateOrderResponse); ok {
		r0 = rf(ctx, in, opts...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*internal_proto.CreateOrderResponse)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, *internal_proto.CreateOrderRequest, ...grpc.CallOption) error); ok {
		r1 = rf(ctx, in, opts...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// OrderServiceClient_CreateOrder_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'CreateOrder'
type OrderServiceClient_CreateOrder_Call struct {
	*mock.Call
}

// CreateOrder is a helper method to define mock.On call
//   - ctx context.Context
//   - in *internal_proto.CreateOrderRequest
//   - opts ...grpc.CallOption
func (_e *OrderServiceClient_Expecter) CreateOrder(ctx interface{}, in interface{}, opts ...interface{}) *OrderServiceClient_CreateOrder_Call {
	return &OrderServiceClient_CreateOrder_Call{Call: _e.mock.On("CreateOrder",
		append([]interface{}{ctx, in}, opts...)...)}
}

func (_c *OrderServiceClient_CreateOrder_Call) Run(run func(ctx context.Context, in *internal_proto.CreateOrderRequest, opts ...grpc.CallOption)) *OrderServiceClient_CreateOrder_Call {
	_c.Call.Run(func(args mock.Arguments) {
		variadicArgs := make([]grpc.CallOption, len(args)-2)
		for i, a := range args[2:] {
			if a != nil {
				variadicArgs[i] = a.(grpc.CallOption)
			}
		}
		run(args[0].(context.Context), args[1].(*internal_proto.CreateOrderRequest), variadicArgs...)
	})
	return _c
}

func (_c *OrderServiceClient_CreateOrder_Call) Return(_a0 *internal_proto.CreateOrderResponse, _a1 error) *OrderServiceClient_CreateOrder_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *OrderServiceClient_CreateOrder_Call) RunAndReturn(run func(context.Context, *internal_proto.CreateOrderRequest, ...grpc.CallOption) (*internal_proto.CreateOrderResponse, error)) *OrderServiceClient_CreateOrder_Call {
	_c.Call.Return(run)
	return _c
}

// GetOrder provides a mock function with given fields: ctx, in, opts
func (_m *OrderServiceClient) GetOrder(ctx context.Context, in *internal_proto.GetOrderRequest, opts ...grpc.CallOption) (*internal_proto.GetOrderResponse, error) {
	_va := make([]interface{}, len(opts))
	for _i := range opts {
		_va[_i] = opts[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, ctx, in)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	if len(ret) == 0 {
		panic("no return value specified for GetOrder")
	}

	var r0 *internal_proto.GetOrderResponse
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, *internal_proto.GetOrderRequest, ...grpc.CallOption) (*internal_proto.GetOrderResponse, error)); ok {
		return rf(ctx, in, opts...)
	}
	if rf, ok := ret.Get(0).(func(context.Context, *internal_proto.GetOrderRequest, ...grpc.CallOption) *internal_proto.GetOrderResponse); ok {
		r0 = rf(ctx, in, opts...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*internal_proto.GetOrderResponse)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, *internal_proto.GetOrderRequest, ...grpc.CallOption) error); ok {
		r1 = rf(ctx, in, opts...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// OrderServiceClient_GetOrder_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'GetOrder'
type OrderServiceClient_GetOrder_Call struct {
	*mock.Call
}

// GetOrder is a helper method to define mock.On call
//   - ctx context.Context
//   - in *internal_proto.GetOrderRequest
//   - opts ...grpc.CallOption
func (_e *OrderServiceClient_Expecter) GetOrder(ctx interface{}, in interface{}, opts ...interface{}) *OrderServiceClient_GetOrder_Call {
	return &OrderServiceClient_GetOrder_Call{Call: _e.mock.On("GetOrder",
		append([]interface{}{ctx, in}, opts...)...)}
}

func (_c *OrderServiceClient_GetOrder_Call) Run(run func(ctx context.Context, in *internal_proto.GetOrderRequest, opts ...grpc.CallOption)) *OrderServiceClient_GetOrder_Call {
	_c.Call.Run(func(args mock.Arguments) {
		variadicArgs := make([]grpc.CallOption, len(args)-2)
		for i, a := range args[2:] {
			if a != nil {
				variadicArgs[i] = a.(grpc.CallOption)
			}
		}
		run(args[0].(context.Context), args[1].(*internal_proto.GetOrderRequest), variadicArgs...)
	})
	return _c
}

func (_c *OrderServiceClient_GetOrder_Call) Return(_a0 *internal_proto.GetOrderResponse, _a1 error) *OrderServiceClient_GetOrder_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *OrderServiceClient_GetOrder_Call) RunAndReturn(run func(context.Context, *internal_proto.GetOrderRequest, ...grpc.CallOption) (*internal_proto.GetOrderResponse, error)) *OrderServiceClient_GetOrder_Call {
	_c.Call.Return(run)
	return _c
}

// GetOrders provides a mock function with given fields: ctx, in, opts
func (_m *OrderServiceClient) GetOrders(ctx context.Context, in *internal_proto.GetOrdersRequest, opts ...grpc.CallOption) (*internal_proto.GetOrdersResponse, error) {
	_va := make([]interface{}, len(opts))
	for _i := range opts {
		_va[_i] = opts[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, ctx, in)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	if len(ret) == 0 {
		panic("no return value specified for GetOrders")
	}

	var r0 *internal_proto.GetOrdersResponse
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, *internal_proto.GetOrdersRequest, ...grpc.CallOption) (*internal_proto.GetOrdersResponse, error)); ok {
		return rf(ctx, in, opts...)
	}
	if rf, ok := ret.Get(0).(func(context.Context, *internal_proto.GetOrdersRequest, ...grpc.CallOption) *internal_proto.GetOrdersResponse); ok {
		r0 = rf(ctx, in, opts...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*internal_proto.GetOrdersResponse)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, *internal_proto.GetOrdersRequest, ...grpc.CallOption) error); ok {
		r1 = rf(ctx, in, opts...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// OrderServiceClient_GetOrders_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'GetOrders'
type OrderServiceClient_GetOrders_Call struct {
	*mock.Call
}

// GetOrders is a helper method to define mock.On call
//   - ctx context.Context
//   - in *internal_proto.GetOrdersRequest
//   - opts ...grpc.CallOption
func (_e *OrderServiceClient_Expecter) GetOrders(ctx interface{}, in interface{}, opts ...interface{}) *OrderServiceClient_GetOrders_Call {
	return &OrderServiceClient_GetOrders_Call{Call: _e.mock.On("GetOrders",
		append([]interface{}{ctx, in}, opts...)...)}
}

func (_c *OrderServiceClient_GetOrders_Call) Run(run func(ctx context.Context, in *internal_proto.GetOrdersRequest, opts ...grpc.CallOption)) *OrderServiceClient_GetOrders_Call {
	_c.Call.Run(func(args mock.Arguments) {
		variadicArgs := make([]grpc.CallOption, len(args)-2)
		for i, a := range args[2:] {
			if a != nil {
				variadicArgs[i] = a.(grpc.CallOption)
			}
		}
		run(args[0].(context.Context), args[1].(*internal_proto.GetOrdersRequest), variadicArgs...)
	})
	return _c
}

func (_c *OrderServiceClient_GetOrders_Call) Return(_a0 *internal_proto.GetOrdersResponse, _a1 error) *OrderServiceClient_GetOrders_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *OrderServiceClient_GetOrders_Call) RunAndReturn(run func(context.Context, *internal_proto.GetOrdersRequest, ...grpc.CallOption) (*internal_proto.GetOrdersResponse, error)) *OrderServiceClient_GetOrders_Call {
	_c.Call.Return(run)
	return _c
}

// NewOrderServiceClient creates a new instance of OrderServiceClient. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewOrderServiceClient(t interface {
	mock.TestingT
	Cleanup(func())
}) *OrderServiceClient {
	mock := &OrderServiceClient{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
