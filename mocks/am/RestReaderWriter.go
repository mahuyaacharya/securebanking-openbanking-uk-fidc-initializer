// Code generated by mockery 2.7.4. DO NOT EDIT.

package mocks

import mock "github.com/stretchr/testify/mock"

// RestReaderWriter is an autogenerated mock type for the RestReaderWriter type
type RestReaderWriter struct {
	mock.Mock
}

// Get provides a mock function with given fields: _a0, _a1, _a2
func (_m *RestReaderWriter) Get(_a0 string, _a1 map[string]string, _a2 interface{}) {
	_m.Called(_a0, _a1, _a2)
}

// Patch provides a mock function with given fields: _a0, _a1, _a2
func (_m *RestReaderWriter) Patch(_a0 string, _a1 map[string]string, _a2 interface{}) {
	_m.Called(_a0, _a1, _a2)
}

// Post provides a mock function with given fields: _a0, _a1, _a2
func (_m *RestReaderWriter) Post(_a0 string, _a1 map[string]string, _a2 interface{}) {
	_m.Called(_a0, _a1, _a2)
}
