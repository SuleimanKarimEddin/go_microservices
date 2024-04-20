// Code generated by mockery v2.42.2. DO NOT EDIT.

package proto_gen

import (
	context "context"
	internal_proto "product_order/internal_proto"

	mock "github.com/stretchr/testify/mock"
)

// ProductServiceServer is an autogenerated mock type for the ProductServiceServer type
type ProductServiceServer struct {
	mock.Mock
}

type ProductServiceServer_Expecter struct {
	mock *mock.Mock
}

func (_m *ProductServiceServer) EXPECT() *ProductServiceServer_Expecter {
	return &ProductServiceServer_Expecter{mock: &_m.Mock}
}

// CreateProduct provides a mock function with given fields: _a0, _a1
func (_m *ProductServiceServer) CreateProduct(_a0 context.Context, _a1 *internal_proto.CreateProductRequest) (*internal_proto.CreateProductResponse, error) {
	ret := _m.Called(_a0, _a1)

	if len(ret) == 0 {
		panic("no return value specified for CreateProduct")
	}

	var r0 *internal_proto.CreateProductResponse
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, *internal_proto.CreateProductRequest) (*internal_proto.CreateProductResponse, error)); ok {
		return rf(_a0, _a1)
	}
	if rf, ok := ret.Get(0).(func(context.Context, *internal_proto.CreateProductRequest) *internal_proto.CreateProductResponse); ok {
		r0 = rf(_a0, _a1)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*internal_proto.CreateProductResponse)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, *internal_proto.CreateProductRequest) error); ok {
		r1 = rf(_a0, _a1)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// ProductServiceServer_CreateProduct_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'CreateProduct'
type ProductServiceServer_CreateProduct_Call struct {
	*mock.Call
}

// CreateProduct is a helper method to define mock.On call
//   - _a0 context.Context
//   - _a1 *internal_proto.CreateProductRequest
func (_e *ProductServiceServer_Expecter) CreateProduct(_a0 interface{}, _a1 interface{}) *ProductServiceServer_CreateProduct_Call {
	return &ProductServiceServer_CreateProduct_Call{Call: _e.mock.On("CreateProduct", _a0, _a1)}
}

func (_c *ProductServiceServer_CreateProduct_Call) Run(run func(_a0 context.Context, _a1 *internal_proto.CreateProductRequest)) *ProductServiceServer_CreateProduct_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(*internal_proto.CreateProductRequest))
	})
	return _c
}

func (_c *ProductServiceServer_CreateProduct_Call) Return(_a0 *internal_proto.CreateProductResponse, _a1 error) *ProductServiceServer_CreateProduct_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *ProductServiceServer_CreateProduct_Call) RunAndReturn(run func(context.Context, *internal_proto.CreateProductRequest) (*internal_proto.CreateProductResponse, error)) *ProductServiceServer_CreateProduct_Call {
	_c.Call.Return(run)
	return _c
}

// GetProduct provides a mock function with given fields: _a0, _a1
func (_m *ProductServiceServer) GetProduct(_a0 context.Context, _a1 *internal_proto.GetProductRequest) (*internal_proto.GetProductResponse, error) {
	ret := _m.Called(_a0, _a1)

	if len(ret) == 0 {
		panic("no return value specified for GetProduct")
	}

	var r0 *internal_proto.GetProductResponse
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, *internal_proto.GetProductRequest) (*internal_proto.GetProductResponse, error)); ok {
		return rf(_a0, _a1)
	}
	if rf, ok := ret.Get(0).(func(context.Context, *internal_proto.GetProductRequest) *internal_proto.GetProductResponse); ok {
		r0 = rf(_a0, _a1)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*internal_proto.GetProductResponse)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, *internal_proto.GetProductRequest) error); ok {
		r1 = rf(_a0, _a1)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// ProductServiceServer_GetProduct_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'GetProduct'
type ProductServiceServer_GetProduct_Call struct {
	*mock.Call
}

// GetProduct is a helper method to define mock.On call
//   - _a0 context.Context
//   - _a1 *internal_proto.GetProductRequest
func (_e *ProductServiceServer_Expecter) GetProduct(_a0 interface{}, _a1 interface{}) *ProductServiceServer_GetProduct_Call {
	return &ProductServiceServer_GetProduct_Call{Call: _e.mock.On("GetProduct", _a0, _a1)}
}

func (_c *ProductServiceServer_GetProduct_Call) Run(run func(_a0 context.Context, _a1 *internal_proto.GetProductRequest)) *ProductServiceServer_GetProduct_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(*internal_proto.GetProductRequest))
	})
	return _c
}

func (_c *ProductServiceServer_GetProduct_Call) Return(_a0 *internal_proto.GetProductResponse, _a1 error) *ProductServiceServer_GetProduct_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *ProductServiceServer_GetProduct_Call) RunAndReturn(run func(context.Context, *internal_proto.GetProductRequest) (*internal_proto.GetProductResponse, error)) *ProductServiceServer_GetProduct_Call {
	_c.Call.Return(run)
	return _c
}

// GetProducts provides a mock function with given fields: _a0, _a1
func (_m *ProductServiceServer) GetProducts(_a0 context.Context, _a1 *internal_proto.GetProductsRequest) (*internal_proto.GetProductsResponse, error) {
	ret := _m.Called(_a0, _a1)

	if len(ret) == 0 {
		panic("no return value specified for GetProducts")
	}

	var r0 *internal_proto.GetProductsResponse
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, *internal_proto.GetProductsRequest) (*internal_proto.GetProductsResponse, error)); ok {
		return rf(_a0, _a1)
	}
	if rf, ok := ret.Get(0).(func(context.Context, *internal_proto.GetProductsRequest) *internal_proto.GetProductsResponse); ok {
		r0 = rf(_a0, _a1)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*internal_proto.GetProductsResponse)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, *internal_proto.GetProductsRequest) error); ok {
		r1 = rf(_a0, _a1)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// ProductServiceServer_GetProducts_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'GetProducts'
type ProductServiceServer_GetProducts_Call struct {
	*mock.Call
}

// GetProducts is a helper method to define mock.On call
//   - _a0 context.Context
//   - _a1 *internal_proto.GetProductsRequest
func (_e *ProductServiceServer_Expecter) GetProducts(_a0 interface{}, _a1 interface{}) *ProductServiceServer_GetProducts_Call {
	return &ProductServiceServer_GetProducts_Call{Call: _e.mock.On("GetProducts", _a0, _a1)}
}

func (_c *ProductServiceServer_GetProducts_Call) Run(run func(_a0 context.Context, _a1 *internal_proto.GetProductsRequest)) *ProductServiceServer_GetProducts_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(*internal_proto.GetProductsRequest))
	})
	return _c
}

func (_c *ProductServiceServer_GetProducts_Call) Return(_a0 *internal_proto.GetProductsResponse, _a1 error) *ProductServiceServer_GetProducts_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *ProductServiceServer_GetProducts_Call) RunAndReturn(run func(context.Context, *internal_proto.GetProductsRequest) (*internal_proto.GetProductsResponse, error)) *ProductServiceServer_GetProducts_Call {
	_c.Call.Return(run)
	return _c
}

// mustEmbedUnimplementedProductServiceServer provides a mock function with given fields:
func (_m *ProductServiceServer) mustEmbedUnimplementedProductServiceServer() {
	_m.Called()
}

// ProductServiceServer_mustEmbedUnimplementedProductServiceServer_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'mustEmbedUnimplementedProductServiceServer'
type ProductServiceServer_mustEmbedUnimplementedProductServiceServer_Call struct {
	*mock.Call
}

// mustEmbedUnimplementedProductServiceServer is a helper method to define mock.On call
func (_e *ProductServiceServer_Expecter) mustEmbedUnimplementedProductServiceServer() *ProductServiceServer_mustEmbedUnimplementedProductServiceServer_Call {
	return &ProductServiceServer_mustEmbedUnimplementedProductServiceServer_Call{Call: _e.mock.On("mustEmbedUnimplementedProductServiceServer")}
}

func (_c *ProductServiceServer_mustEmbedUnimplementedProductServiceServer_Call) Run(run func()) *ProductServiceServer_mustEmbedUnimplementedProductServiceServer_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run()
	})
	return _c
}

func (_c *ProductServiceServer_mustEmbedUnimplementedProductServiceServer_Call) Return() *ProductServiceServer_mustEmbedUnimplementedProductServiceServer_Call {
	_c.Call.Return()
	return _c
}

func (_c *ProductServiceServer_mustEmbedUnimplementedProductServiceServer_Call) RunAndReturn(run func()) *ProductServiceServer_mustEmbedUnimplementedProductServiceServer_Call {
	_c.Call.Return(run)
	return _c
}

// NewProductServiceServer creates a new instance of ProductServiceServer. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewProductServiceServer(t interface {
	mock.TestingT
	Cleanup(func())
}) *ProductServiceServer {
	mock := &ProductServiceServer{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
