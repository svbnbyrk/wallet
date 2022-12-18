// Code generated by MockGen. DO NOT EDIT.
// Source: interface.go

// Package usecase_test is a generated GoMock package.
package usecase_test

import (
	context "context"
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
	entity "github.com/svbnbyrk/wallet/internal/entity"
)

// MockTransaction is a mock of Transaction interface.
type MockTransaction struct {
	ctrl     *gomock.Controller
	recorder *MockTransactionMockRecorder
}

// MockTransactionMockRecorder is the mock recorder for MockTransaction.
type MockTransactionMockRecorder struct {
	mock *MockTransaction
}

// NewMockTransaction creates a new mock instance.
func NewMockTransaction(ctrl *gomock.Controller) *MockTransaction {
	mock := &MockTransaction{ctrl: ctrl}
	mock.recorder = &MockTransactionMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockTransaction) EXPECT() *MockTransactionMockRecorder {
	return m.recorder
}

// History mocks base method.
func (m *MockTransaction) History(arg0 context.Context) ([]entity.Transaction, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "History", arg0)
	ret0, _ := ret[0].([]entity.Transaction)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// History indicates an expected call of History.
func (mr *MockTransactionMockRecorder) History(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "History", reflect.TypeOf((*MockTransaction)(nil).History), arg0)
}

// Post mocks base method.
func (m *MockTransaction) Post(arg0 context.Context, arg1 entity.Transaction) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Post", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// Post indicates an expected call of Post.
func (mr *MockTransactionMockRecorder) Post(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Post", reflect.TypeOf((*MockTransaction)(nil).Post), arg0, arg1)
}

// MockTransactionRepository is a mock of TransactionRepository interface.
type MockTransactionRepository struct {
	ctrl     *gomock.Controller
	recorder *MockTransactionRepositoryMockRecorder
}

// MockTransactionRepositoryMockRecorder is the mock recorder for MockTransactionRepository.
type MockTransactionRepositoryMockRecorder struct {
	mock *MockTransactionRepository
}

// NewMockTransactionRepository creates a new mock instance.
func NewMockTransactionRepository(ctrl *gomock.Controller) *MockTransactionRepository {
	mock := &MockTransactionRepository{ctrl: ctrl}
	mock.recorder = &MockTransactionRepositoryMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockTransactionRepository) EXPECT() *MockTransactionRepositoryMockRecorder {
	return m.recorder
}

// GetHistory mocks base method.
func (m *MockTransactionRepository) GetHistory(arg0 context.Context) ([]entity.Transaction, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetHistory", arg0)
	ret0, _ := ret[0].([]entity.Transaction)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetHistory indicates an expected call of GetHistory.
func (mr *MockTransactionRepositoryMockRecorder) GetHistory(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetHistory", reflect.TypeOf((*MockTransactionRepository)(nil).GetHistory), arg0)
}

// Store mocks base method.
func (m *MockTransactionRepository) Store(arg0 context.Context, arg1 entity.Transaction) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Store", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// Store indicates an expected call of Store.
func (mr *MockTransactionRepositoryMockRecorder) Store(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Store", reflect.TypeOf((*MockTransactionRepository)(nil).Store), arg0, arg1)
}

// MockWallet is a mock of Wallet interface.
type MockWallet struct {
	ctrl     *gomock.Controller
	recorder *MockWalletMockRecorder
}

// MockWalletMockRecorder is the mock recorder for MockWallet.
type MockWalletMockRecorder struct {
	mock *MockWallet
}

// NewMockWallet creates a new mock instance.
func NewMockWallet(ctrl *gomock.Controller) *MockWallet {
	mock := &MockWallet{ctrl: ctrl}
	mock.recorder = &MockWalletMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockWallet) EXPECT() *MockWalletMockRecorder {
	return m.recorder
}

// GetWalletsbyUser mocks base method.
func (m *MockWallet) GetWalletsbyUser(arg0 context.Context, arg1 int64) ([]entity.Wallet, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetWalletsbyUser", arg0, arg1)
	ret0, _ := ret[0].([]entity.Wallet)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetWalletsbyUser indicates an expected call of GetWalletsbyUser.
func (mr *MockWalletMockRecorder) GetWalletsbyUser(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetWalletsbyUser", reflect.TypeOf((*MockWallet)(nil).GetWalletsbyUser), arg0, arg1)
}

// Store mocks base method.
func (m *MockWallet) Store(arg0 context.Context, arg1 entity.Wallet) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Store", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// Store indicates an expected call of Store.
func (mr *MockWalletMockRecorder) Store(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Store", reflect.TypeOf((*MockWallet)(nil).Store), arg0, arg1)
}

// MockWalletRepository is a mock of WalletRepository interface.
type MockWalletRepository struct {
	ctrl     *gomock.Controller
	recorder *MockWalletRepositoryMockRecorder
}

// MockWalletRepositoryMockRecorder is the mock recorder for MockWalletRepository.
type MockWalletRepositoryMockRecorder struct {
	mock *MockWalletRepository
}

// NewMockWalletRepository creates a new mock instance.
func NewMockWalletRepository(ctrl *gomock.Controller) *MockWalletRepository {
	mock := &MockWalletRepository{ctrl: ctrl}
	mock.recorder = &MockWalletRepositoryMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockWalletRepository) EXPECT() *MockWalletRepositoryMockRecorder {
	return m.recorder
}

// Get mocks base method.
func (m *MockWalletRepository) Get(arg0 context.Context, arg1 int64) (entity.Wallet, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Get", arg0, arg1)
	ret0, _ := ret[0].(entity.Wallet)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Get indicates an expected call of Get.
func (mr *MockWalletRepositoryMockRecorder) Get(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Get", reflect.TypeOf((*MockWalletRepository)(nil).Get), arg0, arg1)
}

// GetbyUserId mocks base method.
func (m *MockWalletRepository) GetbyUserId(arg0 context.Context, arg1 int64) ([]entity.Wallet, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetbyUserId", arg0, arg1)
	ret0, _ := ret[0].([]entity.Wallet)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetbyUserId indicates an expected call of GetbyUserId.
func (mr *MockWalletRepositoryMockRecorder) GetbyUserId(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetbyUserId", reflect.TypeOf((*MockWalletRepository)(nil).GetbyUserId), arg0, arg1)
}

// Store mocks base method.
func (m *MockWalletRepository) Store(arg0 context.Context, arg1 entity.Wallet) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Store", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// Store indicates an expected call of Store.
func (mr *MockWalletRepositoryMockRecorder) Store(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Store", reflect.TypeOf((*MockWalletRepository)(nil).Store), arg0, arg1)
}

// Update mocks base method.
func (m *MockWalletRepository) Update(arg0 context.Context, arg1 entity.Wallet) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Update", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// Update indicates an expected call of Update.
func (mr *MockWalletRepositoryMockRecorder) Update(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Update", reflect.TypeOf((*MockWalletRepository)(nil).Update), arg0, arg1)
}

// MockUser is a mock of User interface.
type MockUser struct {
	ctrl     *gomock.Controller
	recorder *MockUserMockRecorder
}

// MockUserMockRecorder is the mock recorder for MockUser.
type MockUserMockRecorder struct {
	mock *MockUser
}

// NewMockUser creates a new mock instance.
func NewMockUser(ctrl *gomock.Controller) *MockUser {
	mock := &MockUser{ctrl: ctrl}
	mock.recorder = &MockUserMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockUser) EXPECT() *MockUserMockRecorder {
	return m.recorder
}

// Store mocks base method.
func (m *MockUser) Store(arg0 context.Context, arg1 entity.User) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Store", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// Store indicates an expected call of Store.
func (mr *MockUserMockRecorder) Store(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Store", reflect.TypeOf((*MockUser)(nil).Store), arg0, arg1)
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

// Store mocks base method.
func (m *MockUserRepository) Store(arg0 context.Context, arg1 entity.User) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Store", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// Store indicates an expected call of Store.
func (mr *MockUserRepositoryMockRecorder) Store(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Store", reflect.TypeOf((*MockUserRepository)(nil).Store), arg0, arg1)
}

// Update mocks base method.
func (m *MockUserRepository) Update(arg0 context.Context, arg1 entity.User) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Update", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// Update indicates an expected call of Update.
func (mr *MockUserRepositoryMockRecorder) Update(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Update", reflect.TypeOf((*MockUserRepository)(nil).Update), arg0, arg1)
}

// MockExchangeRepository is a mock of ExchangeRepository interface.
type MockExchangeRepository struct {
	ctrl     *gomock.Controller
	recorder *MockExchangeRepositoryMockRecorder
}

// MockExchangeRepositoryMockRecorder is the mock recorder for MockExchangeRepository.
type MockExchangeRepositoryMockRecorder struct {
	mock *MockExchangeRepository
}

// NewMockExchangeRepository creates a new mock instance.
func NewMockExchangeRepository(ctrl *gomock.Controller) *MockExchangeRepository {
	mock := &MockExchangeRepository{ctrl: ctrl}
	mock.recorder = &MockExchangeRepositoryMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockExchangeRepository) EXPECT() *MockExchangeRepositoryMockRecorder {
	return m.recorder
}

// GetByCurrency mocks base method.
func (m *MockExchangeRepository) GetByCurrency(arg0 context.Context, arg1 string) (entity.Exchange, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetByCurrency", arg0, arg1)
	ret0, _ := ret[0].(entity.Exchange)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetByCurrency indicates an expected call of GetByCurrency.
func (mr *MockExchangeRepositoryMockRecorder) GetByCurrency(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetByCurrency", reflect.TypeOf((*MockExchangeRepository)(nil).GetByCurrency), arg0, arg1)
}
