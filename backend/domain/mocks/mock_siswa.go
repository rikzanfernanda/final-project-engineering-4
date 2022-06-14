// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/rg-km/final-project-engineering-4/backend/domain (interfaces: SiswaRepository,SiswaUseCase)

// Package mocks is a generated GoMock package.
package mocks

import (
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
	domain "github.com/rg-km/final-project-engineering-4/backend/domain"
)

// MockSiswaRepository is a mock of SiswaRepository interface.
type MockSiswaRepository struct {
	ctrl     *gomock.Controller
	recorder *MockSiswaRepositoryMockRecorder
}

// MockSiswaRepositoryMockRecorder is the mock recorder for MockSiswaRepository.
type MockSiswaRepositoryMockRecorder struct {
	mock *MockSiswaRepository
}

// NewMockSiswaRepository creates a new mock instance.
func NewMockSiswaRepository(ctrl *gomock.Controller) *MockSiswaRepository {
	mock := &MockSiswaRepository{ctrl: ctrl}
	mock.recorder = &MockSiswaRepositoryMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockSiswaRepository) EXPECT() *MockSiswaRepositoryMockRecorder {
	return m.recorder
}

// GetByEmail mocks base method.
func (m *MockSiswaRepository) GetByEmail(arg0 string) (*domain.Siswa, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetByEmail", arg0)
	ret0, _ := ret[0].(*domain.Siswa)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetByEmail indicates an expected call of GetByEmail.
func (mr *MockSiswaRepositoryMockRecorder) GetByEmail(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetByEmail", reflect.TypeOf((*MockSiswaRepository)(nil).GetByEmail), arg0)
}

// MockSiswaUseCase is a mock of SiswaUseCase interface.
type MockSiswaUseCase struct {
	ctrl     *gomock.Controller
	recorder *MockSiswaUseCaseMockRecorder
}

// MockSiswaUseCaseMockRecorder is the mock recorder for MockSiswaUseCase.
type MockSiswaUseCaseMockRecorder struct {
	mock *MockSiswaUseCase
}

// NewMockSiswaUseCase creates a new mock instance.
func NewMockSiswaUseCase(ctrl *gomock.Controller) *MockSiswaUseCase {
	mock := &MockSiswaUseCase{ctrl: ctrl}
	mock.recorder = &MockSiswaUseCaseMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockSiswaUseCase) EXPECT() *MockSiswaUseCaseMockRecorder {
	return m.recorder
}