// Code generated by mockery v1.0.0. DO NOT EDIT.

// NOTE: run 'make update-mocks' from payments-api top folder to update this file and generate new ones.

package payments

import mock "github.com/stretchr/testify/mock"

// MockService is an autogenerated mock type for the Service type
type MockService struct {
	mock.Mock
}

// DeletePayment provides a mock function with given fields: req
func (_m *MockService) DeletePayment(req DeletePaymentRequest) (*DeletePaymentResponse, error) {
	ret := _m.Called(req)

	var r0 *DeletePaymentResponse
	if rf, ok := ret.Get(0).(func(DeletePaymentRequest) *DeletePaymentResponse); ok {
		r0 = rf(req)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*DeletePaymentResponse)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(DeletePaymentRequest) error); ok {
		r1 = rf(req)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetListOfPayments provides a mock function with given fields: req
func (_m *MockService) GetListOfPayments(req GetListOfPaymentsRequest) (*GetListOfPaymentsResponse, error) {
	ret := _m.Called(req)

	var r0 *GetListOfPaymentsResponse
	if rf, ok := ret.Get(0).(func(GetListOfPaymentsRequest) *GetListOfPaymentsResponse); ok {
		r0 = rf(req)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*GetListOfPaymentsResponse)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(GetListOfPaymentsRequest) error); ok {
		r1 = rf(req)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetPayment provides a mock function with given fields: req
func (_m *MockService) GetPayment(req GetPaymentRequest) (*GetPaymentResponse, error) {
	ret := _m.Called(req)

	var r0 *GetPaymentResponse
	if rf, ok := ret.Get(0).(func(GetPaymentRequest) *GetPaymentResponse); ok {
		r0 = rf(req)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*GetPaymentResponse)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(GetPaymentRequest) error); ok {
		r1 = rf(req)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// PostPayment provides a mock function with given fields: req
func (_m *MockService) PostPayment(req CreatePaymentRequest) (*CreatePaymentResponse, error) {
	ret := _m.Called(req)

	var r0 *CreatePaymentResponse
	if rf, ok := ret.Get(0).(func(CreatePaymentRequest) *CreatePaymentResponse); ok {
		r0 = rf(req)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*CreatePaymentResponse)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(CreatePaymentRequest) error); ok {
		r1 = rf(req)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// UpdatePayment provides a mock function with given fields: req
func (_m *MockService) UpdatePayment(req UpdatePaymentRequest) (*UpdatePaymentResponse, error) {
	ret := _m.Called(req)

	var r0 *UpdatePaymentResponse
	if rf, ok := ret.Get(0).(func(UpdatePaymentRequest) *UpdatePaymentResponse); ok {
		r0 = rf(req)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*UpdatePaymentResponse)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(UpdatePaymentRequest) error); ok {
		r1 = rf(req)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}
