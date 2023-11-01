// Code generated by mockery v2.36.0. DO NOT EDIT.

package mocks

import (
	campaign "github.com/RianIhsan/raise-unity/campaign"
	mock "github.com/stretchr/testify/mock"

	transaction "github.com/RianIhsan/raise-unity/transaction"

	user "github.com/RianIhsan/raise-unity/user"
)

// Service is an autogenerated mock type for the Service type
type Service struct {
	mock.Mock
}

// DeleteCampaignById provides a mock function with given fields: id
func (_m *Service) DeleteCampaignById(id int) (campaign.Campaign, error) {
	ret := _m.Called(id)

	var r0 campaign.Campaign
	var r1 error
	if rf, ok := ret.Get(0).(func(int) (campaign.Campaign, error)); ok {
		return rf(id)
	}
	if rf, ok := ret.Get(0).(func(int) campaign.Campaign); ok {
		r0 = rf(id)
	} else {
		r0 = ret.Get(0).(campaign.Campaign)
	}

	if rf, ok := ret.Get(1).(func(int) error); ok {
		r1 = rf(id)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// DeleteUserById provides a mock function with given fields: id
func (_m *Service) DeleteUserById(id int) (user.User, error) {
	ret := _m.Called(id)

	var r0 user.User
	var r1 error
	if rf, ok := ret.Get(0).(func(int) (user.User, error)); ok {
		return rf(id)
	}
	if rf, ok := ret.Get(0).(func(int) user.User); ok {
		r0 = rf(id)
	} else {
		r0 = ret.Get(0).(user.User)
	}

	if rf, ok := ret.Get(1).(func(int) error); ok {
		r1 = rf(id)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetTransactionsPagination provides a mock function with given fields: page, pageSize
func (_m *Service) GetTransactionsPagination(page int, pageSize int) ([]transaction.Transaction, int, int, int, int, error) {
	ret := _m.Called(page, pageSize)

	var r0 []transaction.Transaction
	var r1 int
	var r2 int
	var r3 int
	var r4 int
	var r5 error
	if rf, ok := ret.Get(0).(func(int, int) ([]transaction.Transaction, int, int, int, int, error)); ok {
		return rf(page, pageSize)
	}
	if rf, ok := ret.Get(0).(func(int, int) []transaction.Transaction); ok {
		r0 = rf(page, pageSize)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]transaction.Transaction)
		}
	}

	if rf, ok := ret.Get(1).(func(int, int) int); ok {
		r1 = rf(page, pageSize)
	} else {
		r1 = ret.Get(1).(int)
	}

	if rf, ok := ret.Get(2).(func(int, int) int); ok {
		r2 = rf(page, pageSize)
	} else {
		r2 = ret.Get(2).(int)
	}

	if rf, ok := ret.Get(3).(func(int, int) int); ok {
		r3 = rf(page, pageSize)
	} else {
		r3 = ret.Get(3).(int)
	}

	if rf, ok := ret.Get(4).(func(int, int) int); ok {
		r4 = rf(page, pageSize)
	} else {
		r4 = ret.Get(4).(int)
	}

	if rf, ok := ret.Get(5).(func(int, int) error); ok {
		r5 = rf(page, pageSize)
	} else {
		r5 = ret.Error(5)
	}

	return r0, r1, r2, r3, r4, r5
}

// GetUsersPagination provides a mock function with given fields: page, pageSize
func (_m *Service) GetUsersPagination(page int, pageSize int) ([]user.User, int, int, int, int, error) {
	ret := _m.Called(page, pageSize)

	var r0 []user.User
	var r1 int
	var r2 int
	var r3 int
	var r4 int
	var r5 error
	if rf, ok := ret.Get(0).(func(int, int) ([]user.User, int, int, int, int, error)); ok {
		return rf(page, pageSize)
	}
	if rf, ok := ret.Get(0).(func(int, int) []user.User); ok {
		r0 = rf(page, pageSize)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]user.User)
		}
	}

	if rf, ok := ret.Get(1).(func(int, int) int); ok {
		r1 = rf(page, pageSize)
	} else {
		r1 = ret.Get(1).(int)
	}

	if rf, ok := ret.Get(2).(func(int, int) int); ok {
		r2 = rf(page, pageSize)
	} else {
		r2 = ret.Get(2).(int)
	}

	if rf, ok := ret.Get(3).(func(int, int) int); ok {
		r3 = rf(page, pageSize)
	} else {
		r3 = ret.Get(3).(int)
	}

	if rf, ok := ret.Get(4).(func(int, int) int); ok {
		r4 = rf(page, pageSize)
	} else {
		r4 = ret.Get(4).(int)
	}

	if rf, ok := ret.Get(5).(func(int, int) error); ok {
		r5 = rf(page, pageSize)
	} else {
		r5 = ret.Error(5)
	}

	return r0, r1, r2, r3, r4, r5
}

// SearchTransactionByUsername provides a mock function with given fields: name
func (_m *Service) SearchTransactionByUsername(name string) ([]transaction.Transaction, error) {
	ret := _m.Called(name)

	var r0 []transaction.Transaction
	var r1 error
	if rf, ok := ret.Get(0).(func(string) ([]transaction.Transaction, error)); ok {
		return rf(name)
	}
	if rf, ok := ret.Get(0).(func(string) []transaction.Transaction); ok {
		r0 = rf(name)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]transaction.Transaction)
		}
	}

	if rf, ok := ret.Get(1).(func(string) error); ok {
		r1 = rf(name)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// SearchUserByName provides a mock function with given fields: name
func (_m *Service) SearchUserByName(name string) ([]user.User, error) {
	ret := _m.Called(name)

	var r0 []user.User
	var r1 error
	if rf, ok := ret.Get(0).(func(string) ([]user.User, error)); ok {
		return rf(name)
	}
	if rf, ok := ret.Get(0).(func(string) []user.User); ok {
		r0 = rf(name)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]user.User)
		}
	}

	if rf, ok := ret.Get(1).(func(string) error); ok {
		r1 = rf(name)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// NewService creates a new instance of Service. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewService(t interface {
	mock.TestingT
	Cleanup(func())
}) *Service {
	mock := &Service{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
