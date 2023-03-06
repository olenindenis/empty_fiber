// Code generated by MockGen. DO NOT EDIT.
// Source: ./internal/core/ports/user.go

// Package mocks is a generated GoMock package.
package mocks

import (
	domain "envs/internal/core/domain"
	dto "envs/internal/dto"
	reflect "reflect"

	fiber "github.com/gofiber/fiber/v2"
	gomock "github.com/golang/mock/gomock"
)

// MockUserHandlers is a mock of UserHandlers interface.
type MockUserHandlers struct {
	ctrl     *gomock.Controller
	recorder *MockUserHandlersMockRecorder
}

// MockUserHandlersMockRecorder is the mock recorder for MockUserHandlers.
type MockUserHandlersMockRecorder struct {
	mock *MockUserHandlers
}

// NewMockUserHandlers creates a new mock instance.
func NewMockUserHandlers(ctrl *gomock.Controller) *MockUserHandlers {
	mock := &MockUserHandlers{ctrl: ctrl}
	mock.recorder = &MockUserHandlersMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockUserHandlers) EXPECT() *MockUserHandlersMockRecorder {
	return m.recorder
}

// Delete mocks base method.
func (m *MockUserHandlers) Delete(ctx *fiber.Ctx) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Delete", ctx)
	ret0, _ := ret[0].(error)
	return ret0
}

// Delete indicates an expected call of Delete.
func (mr *MockUserHandlersMockRecorder) Delete(ctx interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Delete", reflect.TypeOf((*MockUserHandlers)(nil).Delete), ctx)
}

// List mocks base method.
func (m *MockUserHandlers) List(ctx *fiber.Ctx) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "List", ctx)
	ret0, _ := ret[0].(error)
	return ret0
}

// List indicates an expected call of List.
func (mr *MockUserHandlersMockRecorder) List(ctx interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "List", reflect.TypeOf((*MockUserHandlers)(nil).List), ctx)
}

// Show mocks base method.
func (m *MockUserHandlers) Show(ctx *fiber.Ctx) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Show", ctx)
	ret0, _ := ret[0].(error)
	return ret0
}

// Show indicates an expected call of Show.
func (mr *MockUserHandlersMockRecorder) Show(ctx interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Show", reflect.TypeOf((*MockUserHandlers)(nil).Show), ctx)
}

// Update mocks base method.
func (m *MockUserHandlers) Update(ctx *fiber.Ctx) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Update", ctx)
	ret0, _ := ret[0].(error)
	return ret0
}

// Update indicates an expected call of Update.
func (mr *MockUserHandlersMockRecorder) Update(ctx interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Update", reflect.TypeOf((*MockUserHandlers)(nil).Update), ctx)
}

// MockUserService is a mock of UserService interface.
type MockUserService struct {
	ctrl     *gomock.Controller
	recorder *MockUserServiceMockRecorder
}

// MockUserServiceMockRecorder is the mock recorder for MockUserService.
type MockUserServiceMockRecorder struct {
	mock *MockUserService
}

// NewMockUserService creates a new mock instance.
func NewMockUserService(ctrl *gomock.Controller) *MockUserService {
	mock := &MockUserService{ctrl: ctrl}
	mock.recorder = &MockUserServiceMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockUserService) EXPECT() *MockUserServiceMockRecorder {
	return m.recorder
}

// Delete mocks base method.
func (m *MockUserService) Delete(id uint) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Delete", id)
	ret0, _ := ret[0].(error)
	return ret0
}

// Delete indicates an expected call of Delete.
func (mr *MockUserServiceMockRecorder) Delete(id interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Delete", reflect.TypeOf((*MockUserService)(nil).Delete), id)
}

// List mocks base method.
func (m *MockUserService) List(filter dto.ListFilter) ([]domain.User, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "List", filter)
	ret0, _ := ret[0].([]domain.User)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// List indicates an expected call of List.
func (mr *MockUserServiceMockRecorder) List(filter interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "List", reflect.TypeOf((*MockUserService)(nil).List), filter)
}

// Show mocks base method.
func (m *MockUserService) Show(id uint) (domain.User, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Show", id)
	ret0, _ := ret[0].(domain.User)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Show indicates an expected call of Show.
func (mr *MockUserServiceMockRecorder) Show(id interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Show", reflect.TypeOf((*MockUserService)(nil).Show), id)
}

// Update mocks base method.
func (m *MockUserService) Update(user domain.User) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Update", user)
	ret0, _ := ret[0].(error)
	return ret0
}

// Update indicates an expected call of Update.
func (mr *MockUserServiceMockRecorder) Update(user interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Update", reflect.TypeOf((*MockUserService)(nil).Update), user)
}

// MockUserRepository is a mock of UserRepository interface.
type MockUserRepository struct {
	ctrl     *gomock.Controller
	recorder *MockUserRepositoryMockRecorder
}

// MockUserRepositoryMockRecorder is the mock recorder for MockUserRepository.
type MockUserRepositoryMockRecorder struct {
	mock *MockUserRepository
}

// NewMockUserRepository creates a new mock instance.
func NewMockUserRepository(ctrl *gomock.Controller) *MockUserRepository {
	mock := &MockUserRepository{ctrl: ctrl}
	mock.recorder = &MockUserRepositoryMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockUserRepository) EXPECT() *MockUserRepositoryMockRecorder {
	return m.recorder
}

// Delete mocks base method.
func (m *MockUserRepository) Delete(id uint) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Delete", id)
	ret0, _ := ret[0].(error)
	return ret0
}

// Delete indicates an expected call of Delete.
func (mr *MockUserRepositoryMockRecorder) Delete(id interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Delete", reflect.TypeOf((*MockUserRepository)(nil).Delete), id)
}

// Find mocks base method.
func (m *MockUserRepository) Find(id uint) (domain.User, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Find", id)
	ret0, _ := ret[0].(domain.User)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Find indicates an expected call of Find.
func (mr *MockUserRepositoryMockRecorder) Find(id interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Find", reflect.TypeOf((*MockUserRepository)(nil).Find), id)
}

// FindByEmail mocks base method.
func (m *MockUserRepository) FindByEmail(email string) (domain.User, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "FindByEmail", email)
	ret0, _ := ret[0].(domain.User)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// FindByEmail indicates an expected call of FindByEmail.
func (mr *MockUserRepositoryMockRecorder) FindByEmail(email interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "FindByEmail", reflect.TypeOf((*MockUserRepository)(nil).FindByEmail), email)
}

// List mocks base method.
func (m *MockUserRepository) List(filter dto.ListFilter) ([]domain.User, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "List", filter)
	ret0, _ := ret[0].([]domain.User)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// List indicates an expected call of List.
func (mr *MockUserRepositoryMockRecorder) List(filter interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "List", reflect.TypeOf((*MockUserRepository)(nil).List), filter)
}

// Store mocks base method.
func (m *MockUserRepository) Store(name, email, password string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Store", name, email, password)
	ret0, _ := ret[0].(error)
	return ret0
}

// Store indicates an expected call of Store.
func (mr *MockUserRepositoryMockRecorder) Store(name, email, password interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Store", reflect.TypeOf((*MockUserRepository)(nil).Store), name, email, password)
}

// Update mocks base method.
func (m *MockUserRepository) Update(user domain.User) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Update", user)
	ret0, _ := ret[0].(error)
	return ret0
}

// Update indicates an expected call of Update.
func (mr *MockUserRepositoryMockRecorder) Update(user interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Update", reflect.TypeOf((*MockUserRepository)(nil).Update), user)
}
