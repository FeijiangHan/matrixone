// Code generated by MockGen. DO NOT EDIT.
// Source: ../types.go

// Package mock_frontend is a generated GoMock package.
package mock_frontend

import (
	context "context"
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
	batch "github.com/matrixorigin/matrixone/pkg/container/batch"
	types "github.com/matrixorigin/matrixone/pkg/container/types"
	tree "github.com/matrixorigin/matrixone/pkg/sql/parsers/tree"
	util "github.com/matrixorigin/matrixone/pkg/util"
	process "github.com/matrixorigin/matrixone/pkg/vm/process"
)

// MockComputationRunner is a mock of ComputationRunner interface.
type MockComputationRunner struct {
	ctrl     *gomock.Controller
	recorder *MockComputationRunnerMockRecorder
}

// MockComputationRunnerMockRecorder is the mock recorder for MockComputationRunner.
type MockComputationRunnerMockRecorder struct {
	mock *MockComputationRunner
}

// NewMockComputationRunner creates a new mock instance.
func NewMockComputationRunner(ctrl *gomock.Controller) *MockComputationRunner {
	mock := &MockComputationRunner{ctrl: ctrl}
	mock.recorder = &MockComputationRunnerMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockComputationRunner) EXPECT() *MockComputationRunnerMockRecorder {
	return m.recorder
}

// Run mocks base method.
func (m *MockComputationRunner) Run(ts uint64) (*util.RunResult, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Run", ts)
	ret0, _ := ret[0].(*util.RunResult)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Run indicates an expected call of Run.
func (mr *MockComputationRunnerMockRecorder) Run(ts interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Run", reflect.TypeOf((*MockComputationRunner)(nil).Run), ts)
}

// MockComputationWrapper is a mock of ComputationWrapper interface.
type MockComputationWrapper struct {
	ctrl     *gomock.Controller
	recorder *MockComputationWrapperMockRecorder
}

// MockComputationWrapperMockRecorder is the mock recorder for MockComputationWrapper.
type MockComputationWrapperMockRecorder struct {
	mock *MockComputationWrapper
}

// NewMockComputationWrapper creates a new mock instance.
func NewMockComputationWrapper(ctrl *gomock.Controller) *MockComputationWrapper {
	mock := &MockComputationWrapper{ctrl: ctrl}
	mock.recorder = &MockComputationWrapperMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockComputationWrapper) EXPECT() *MockComputationWrapperMockRecorder {
	return m.recorder
}

// Compile mocks base method.
func (m *MockComputationWrapper) Compile(requestCtx context.Context, u interface{}, fill func(interface{}, *batch.Batch) error) (interface{}, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Compile", requestCtx, u, fill)
	ret0, _ := ret[0].(interface{})
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Compile indicates an expected call of Compile.
func (mr *MockComputationWrapperMockRecorder) Compile(requestCtx, u, fill interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Compile", reflect.TypeOf((*MockComputationWrapper)(nil).Compile), requestCtx, u, fill)
}

// GetAffectedRows mocks base method.
func (m *MockComputationWrapper) GetAffectedRows() uint64 {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetAffectedRows")
	ret0, _ := ret[0].(uint64)
	return ret0
}

// GetAffectedRows indicates an expected call of GetAffectedRows.
func (mr *MockComputationWrapperMockRecorder) GetAffectedRows() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetAffectedRows", reflect.TypeOf((*MockComputationWrapper)(nil).GetAffectedRows))
}

// GetAst mocks base method.
func (m *MockComputationWrapper) GetAst() tree.Statement {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetAst")
	ret0, _ := ret[0].(tree.Statement)
	return ret0
}

// GetAst indicates an expected call of GetAst.
func (mr *MockComputationWrapperMockRecorder) GetAst() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetAst", reflect.TypeOf((*MockComputationWrapper)(nil).GetAst))
}

// GetColumns mocks base method.
func (m *MockComputationWrapper) GetColumns() ([]interface{}, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetColumns")
	ret0, _ := ret[0].([]interface{})
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetColumns indicates an expected call of GetColumns.
func (mr *MockComputationWrapperMockRecorder) GetColumns() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetColumns", reflect.TypeOf((*MockComputationWrapper)(nil).GetColumns))
}

// GetLoadTag mocks base method.
func (m *MockComputationWrapper) GetLoadTag() bool {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetLoadTag")
	ret0, _ := ret[0].(bool)
	return ret0
}

// GetLoadTag indicates an expected call of GetLoadTag.
func (mr *MockComputationWrapperMockRecorder) GetLoadTag() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetLoadTag", reflect.TypeOf((*MockComputationWrapper)(nil).GetLoadTag))
}

// GetProcess mocks base method.
func (m *MockComputationWrapper) GetProcess() *process.Process {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetProcess")
	ret0, _ := ret[0].(*process.Process)
	return ret0
}

// GetProcess indicates an expected call of GetProcess.
func (mr *MockComputationWrapperMockRecorder) GetProcess() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetProcess", reflect.TypeOf((*MockComputationWrapper)(nil).GetProcess))
}

// GetServerStatus mocks base method.
func (m *MockComputationWrapper) GetServerStatus() uint16 {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetServerStatus")
	ret0, _ := ret[0].(uint16)
	return ret0
}

// GetServerStatus indicates an expected call of GetServerStatus.
func (mr *MockComputationWrapperMockRecorder) GetServerStatus() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetServerStatus", reflect.TypeOf((*MockComputationWrapper)(nil).GetServerStatus))
}

// GetUUID mocks base method.
func (m *MockComputationWrapper) GetUUID() []byte {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetUUID")
	ret0, _ := ret[0].([]byte)
	return ret0
}

// GetUUID indicates an expected call of GetUUID.
func (mr *MockComputationWrapperMockRecorder) GetUUID() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetUUID", reflect.TypeOf((*MockComputationWrapper)(nil).GetUUID))
}

// RecordExecPlan mocks base method.
func (m *MockComputationWrapper) RecordExecPlan(ctx context.Context) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "RecordExecPlan", ctx)
	ret0, _ := ret[0].(error)
	return ret0
}

// RecordExecPlan indicates an expected call of RecordExecPlan.
func (mr *MockComputationWrapperMockRecorder) RecordExecPlan(ctx interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "RecordExecPlan", reflect.TypeOf((*MockComputationWrapper)(nil).RecordExecPlan), ctx)
}

// Run mocks base method.
func (m *MockComputationWrapper) Run(ts uint64) (*util.RunResult, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Run", ts)
	ret0, _ := ret[0].(*util.RunResult)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Run indicates an expected call of Run.
func (mr *MockComputationWrapperMockRecorder) Run(ts interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Run", reflect.TypeOf((*MockComputationWrapper)(nil).Run), ts)
}

// SetDatabaseName mocks base method.
func (m *MockComputationWrapper) SetDatabaseName(db string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SetDatabaseName", db)
	ret0, _ := ret[0].(error)
	return ret0
}

// SetDatabaseName indicates an expected call of SetDatabaseName.
func (mr *MockComputationWrapperMockRecorder) SetDatabaseName(db interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SetDatabaseName", reflect.TypeOf((*MockComputationWrapper)(nil).SetDatabaseName), db)
}

// MockColumnInfo is a mock of ColumnInfo interface.
type MockColumnInfo struct {
	ctrl     *gomock.Controller
	recorder *MockColumnInfoMockRecorder
}

// MockColumnInfoMockRecorder is the mock recorder for MockColumnInfo.
type MockColumnInfoMockRecorder struct {
	mock *MockColumnInfo
}

// NewMockColumnInfo creates a new mock instance.
func NewMockColumnInfo(ctrl *gomock.Controller) *MockColumnInfo {
	mock := &MockColumnInfo{ctrl: ctrl}
	mock.recorder = &MockColumnInfoMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockColumnInfo) EXPECT() *MockColumnInfoMockRecorder {
	return m.recorder
}

// GetName mocks base method.
func (m *MockColumnInfo) GetName() string {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetName")
	ret0, _ := ret[0].(string)
	return ret0
}

// GetName indicates an expected call of GetName.
func (mr *MockColumnInfoMockRecorder) GetName() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetName", reflect.TypeOf((*MockColumnInfo)(nil).GetName))
}

// GetType mocks base method.
func (m *MockColumnInfo) GetType() types.T {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetType")
	ret0, _ := ret[0].(types.T)
	return ret0
}

// GetType indicates an expected call of GetType.
func (mr *MockColumnInfoMockRecorder) GetType() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetType", reflect.TypeOf((*MockColumnInfo)(nil).GetType))
}

// MockTableInfo is a mock of TableInfo interface.
type MockTableInfo struct {
	ctrl     *gomock.Controller
	recorder *MockTableInfoMockRecorder
}

// MockTableInfoMockRecorder is the mock recorder for MockTableInfo.
type MockTableInfoMockRecorder struct {
	mock *MockTableInfo
}

// NewMockTableInfo creates a new mock instance.
func NewMockTableInfo(ctrl *gomock.Controller) *MockTableInfo {
	mock := &MockTableInfo{ctrl: ctrl}
	mock.recorder = &MockTableInfoMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockTableInfo) EXPECT() *MockTableInfoMockRecorder {
	return m.recorder
}

// GetColumns mocks base method.
func (m *MockTableInfo) GetColumns() {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "GetColumns")
}

// GetColumns indicates an expected call of GetColumns.
func (mr *MockTableInfoMockRecorder) GetColumns() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetColumns", reflect.TypeOf((*MockTableInfo)(nil).GetColumns))
}

// MockExecResult is a mock of ExecResult interface.
type MockExecResult struct {
	ctrl     *gomock.Controller
	recorder *MockExecResultMockRecorder
}

// MockExecResultMockRecorder is the mock recorder for MockExecResult.
type MockExecResultMockRecorder struct {
	mock *MockExecResult
}

// NewMockExecResult creates a new mock instance.
func NewMockExecResult(ctrl *gomock.Controller) *MockExecResult {
	mock := &MockExecResult{ctrl: ctrl}
	mock.recorder = &MockExecResultMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockExecResult) EXPECT() *MockExecResultMockRecorder {
	return m.recorder
}

// GetInt64 mocks base method.
func (m *MockExecResult) GetInt64(ctx context.Context, rindex, cindex uint64) (int64, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetInt64", ctx, rindex, cindex)
	ret0, _ := ret[0].(int64)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetInt64 indicates an expected call of GetInt64.
func (mr *MockExecResultMockRecorder) GetInt64(ctx, rindex, cindex interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetInt64", reflect.TypeOf((*MockExecResult)(nil).GetInt64), ctx, rindex, cindex)
}

// GetRowCount mocks base method.
func (m *MockExecResult) GetRowCount() uint64 {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetRowCount")
	ret0, _ := ret[0].(uint64)
	return ret0
}

// GetRowCount indicates an expected call of GetRowCount.
func (mr *MockExecResultMockRecorder) GetRowCount() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetRowCount", reflect.TypeOf((*MockExecResult)(nil).GetRowCount))
}

// GetString mocks base method.
func (m *MockExecResult) GetString(ctx context.Context, rindex, cindex uint64) (string, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetString", ctx, rindex, cindex)
	ret0, _ := ret[0].(string)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetString indicates an expected call of GetString.
func (mr *MockExecResultMockRecorder) GetString(ctx, rindex, cindex interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetString", reflect.TypeOf((*MockExecResult)(nil).GetString), ctx, rindex, cindex)
}

// GetUint64 mocks base method.
func (m *MockExecResult) GetUint64(ctx context.Context, rindex, cindex uint64) (uint64, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetUint64", ctx, rindex, cindex)
	ret0, _ := ret[0].(uint64)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetUint64 indicates an expected call of GetUint64.
func (mr *MockExecResultMockRecorder) GetUint64(ctx, rindex, cindex interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetUint64", reflect.TypeOf((*MockExecResult)(nil).GetUint64), ctx, rindex, cindex)
}

// MockBackgroundExec is a mock of BackgroundExec interface.
type MockBackgroundExec struct {
	ctrl     *gomock.Controller
	recorder *MockBackgroundExecMockRecorder
}

// MockBackgroundExecMockRecorder is the mock recorder for MockBackgroundExec.
type MockBackgroundExecMockRecorder struct {
	mock *MockBackgroundExec
}

// NewMockBackgroundExec creates a new mock instance.
func NewMockBackgroundExec(ctrl *gomock.Controller) *MockBackgroundExec {
	mock := &MockBackgroundExec{ctrl: ctrl}
	mock.recorder = &MockBackgroundExecMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockBackgroundExec) EXPECT() *MockBackgroundExecMockRecorder {
	return m.recorder
}

// ClearExecResultBatches mocks base method.
func (m *MockBackgroundExec) ClearExecResultBatches() {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "ClearExecResultBatches")
}

// ClearExecResultBatches indicates an expected call of ClearExecResultBatches.
func (mr *MockBackgroundExecMockRecorder) ClearExecResultBatches() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ClearExecResultBatches", reflect.TypeOf((*MockBackgroundExec)(nil).ClearExecResultBatches))
}

// ClearExecResultSet mocks base method.
func (m *MockBackgroundExec) ClearExecResultSet() {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "ClearExecResultSet")
}

// ClearExecResultSet indicates an expected call of ClearExecResultSet.
func (mr *MockBackgroundExecMockRecorder) ClearExecResultSet() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ClearExecResultSet", reflect.TypeOf((*MockBackgroundExec)(nil).ClearExecResultSet))
}

// Close mocks base method.
func (m *MockBackgroundExec) Close() {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "Close")
}

// Close indicates an expected call of Close.
func (mr *MockBackgroundExecMockRecorder) Close() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Close", reflect.TypeOf((*MockBackgroundExec)(nil).Close))
}

// Exec mocks base method.
func (m *MockBackgroundExec) Exec(arg0 context.Context, arg1 string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Exec", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// Exec indicates an expected call of Exec.
func (mr *MockBackgroundExecMockRecorder) Exec(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Exec", reflect.TypeOf((*MockBackgroundExec)(nil).Exec), arg0, arg1)
}

// ExecStmt mocks base method.
func (m *MockBackgroundExec) ExecStmt(arg0 context.Context, arg1 tree.Statement) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ExecStmt", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// ExecStmt indicates an expected call of ExecStmt.
func (mr *MockBackgroundExecMockRecorder) ExecStmt(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ExecStmt", reflect.TypeOf((*MockBackgroundExec)(nil).ExecStmt), arg0, arg1)
}

// GetExecResultBatches mocks base method.
func (m *MockBackgroundExec) GetExecResultBatches() []*batch.Batch {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetExecResultBatches")
	ret0, _ := ret[0].([]*batch.Batch)
	return ret0
}

// GetExecResultBatches indicates an expected call of GetExecResultBatches.
func (mr *MockBackgroundExecMockRecorder) GetExecResultBatches() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetExecResultBatches", reflect.TypeOf((*MockBackgroundExec)(nil).GetExecResultBatches))
}

// GetExecResultSet mocks base method.
func (m *MockBackgroundExec) GetExecResultSet() []interface{} {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetExecResultSet")
	ret0, _ := ret[0].([]interface{})
	return ret0
}

// GetExecResultSet indicates an expected call of GetExecResultSet.
func (mr *MockBackgroundExecMockRecorder) GetExecResultSet() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetExecResultSet", reflect.TypeOf((*MockBackgroundExec)(nil).GetExecResultSet))
}

// MockoutputPool is a mock of outputPool interface.
type MockoutputPool struct {
	ctrl     *gomock.Controller
	recorder *MockoutputPoolMockRecorder
}

// MockoutputPoolMockRecorder is the mock recorder for MockoutputPool.
type MockoutputPoolMockRecorder struct {
	mock *MockoutputPool
}

// NewMockoutputPool creates a new mock instance.
func NewMockoutputPool(ctrl *gomock.Controller) *MockoutputPool {
	mock := &MockoutputPool{ctrl: ctrl}
	mock.recorder = &MockoutputPoolMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockoutputPool) EXPECT() *MockoutputPoolMockRecorder {
	return m.recorder
}

// flush mocks base method.
func (m *MockoutputPool) flush() error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "flush")
	ret0, _ := ret[0].(error)
	return ret0
}

// flush indicates an expected call of flush.
func (mr *MockoutputPoolMockRecorder) flush() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "flush", reflect.TypeOf((*MockoutputPool)(nil).flush))
}

// getEmptyRow mocks base method.
func (m *MockoutputPool) getEmptyRow() ([]interface{}, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "getEmptyRow")
	ret0, _ := ret[0].([]interface{})
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// getEmptyRow indicates an expected call of getEmptyRow.
func (mr *MockoutputPoolMockRecorder) getEmptyRow() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "getEmptyRow", reflect.TypeOf((*MockoutputPool)(nil).getEmptyRow))
}

// reset mocks base method.
func (m *MockoutputPool) reset() {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "reset")
}

// reset indicates an expected call of reset.
func (mr *MockoutputPoolMockRecorder) reset() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "reset", reflect.TypeOf((*MockoutputPool)(nil).reset))
}

// resetLineStr mocks base method.
func (m *MockoutputPool) resetLineStr() {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "resetLineStr")
}

// resetLineStr indicates an expected call of resetLineStr.
func (mr *MockoutputPoolMockRecorder) resetLineStr() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "resetLineStr", reflect.TypeOf((*MockoutputPool)(nil).resetLineStr))
}
