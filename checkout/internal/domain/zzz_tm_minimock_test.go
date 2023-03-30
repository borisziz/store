package domain

// Code generated by http://github.com/gojuno/minimock (dev). DO NOT EDIT.

//go:generate minimock -i route256/checkout/internal/domain.TransactionManager -o ./zzz_tm_minimock_test.go -n TransactionManagerMock

import (
	"context"
	"sync"
	mm_atomic "sync/atomic"
	mm_time "time"

	"github.com/gojuno/minimock/v3"
)

// TransactionManagerMock implements TransactionManager
type TransactionManagerMock struct {
	t minimock.Tester

	funcRunTransaction          func(ctx context.Context, isoLevel string, f func(ctxTX context.Context) error) (err error)
	inspectFuncRunTransaction   func(ctx context.Context, isoLevel string, f func(ctxTX context.Context) error)
	afterRunTransactionCounter  uint64
	beforeRunTransactionCounter uint64
	RunTransactionMock          mTransactionManagerMockRunTransaction
}

// NewTransactionManagerMock returns a mock for TransactionManager
func NewTransactionManagerMock(t minimock.Tester) *TransactionManagerMock {
	m := &TransactionManagerMock{t: t}
	if controller, ok := t.(minimock.MockController); ok {
		controller.RegisterMocker(m)
	}

	m.RunTransactionMock = mTransactionManagerMockRunTransaction{mock: m}
	m.RunTransactionMock.callArgs = []*TransactionManagerMockRunTransactionParams{}

	return m
}

type mTransactionManagerMockRunTransaction struct {
	mock               *TransactionManagerMock
	defaultExpectation *TransactionManagerMockRunTransactionExpectation
	expectations       []*TransactionManagerMockRunTransactionExpectation

	callArgs []*TransactionManagerMockRunTransactionParams
	mutex    sync.RWMutex
}

// TransactionManagerMockRunTransactionExpectation specifies expectation struct of the TransactionManager.RunTransaction
type TransactionManagerMockRunTransactionExpectation struct {
	mock    *TransactionManagerMock
	params  *TransactionManagerMockRunTransactionParams
	results *TransactionManagerMockRunTransactionResults
	Counter uint64
}

// TransactionManagerMockRunTransactionParams contains parameters of the TransactionManager.RunTransaction
type TransactionManagerMockRunTransactionParams struct {
	ctx      context.Context
	isoLevel string
	f        func(ctxTX context.Context) error
}

// TransactionManagerMockRunTransactionResults contains results of the TransactionManager.RunTransaction
type TransactionManagerMockRunTransactionResults struct {
	err error
}

// Expect sets up expected params for TransactionManager.RunTransaction
func (mmRunTransaction *mTransactionManagerMockRunTransaction) Expect(ctx context.Context, isoLevel string, f func(ctxTX context.Context) error) *mTransactionManagerMockRunTransaction {
	if mmRunTransaction.mock.funcRunTransaction != nil {
		mmRunTransaction.mock.t.Fatalf("TransactionManagerMock.RunTransaction mock is already set by Set")
	}

	if mmRunTransaction.defaultExpectation == nil {
		mmRunTransaction.defaultExpectation = &TransactionManagerMockRunTransactionExpectation{}
	}

	mmRunTransaction.defaultExpectation.params = &TransactionManagerMockRunTransactionParams{ctx, isoLevel, f}
	for _, e := range mmRunTransaction.expectations {
		if minimock.Equal(e.params, mmRunTransaction.defaultExpectation.params) {
			mmRunTransaction.mock.t.Fatalf("Expectation set by When has same params: %#v", *mmRunTransaction.defaultExpectation.params)
		}
	}

	return mmRunTransaction
}

// Inspect accepts an inspector function that has same arguments as the TransactionManager.RunTransaction
func (mmRunTransaction *mTransactionManagerMockRunTransaction) Inspect(f func(ctx context.Context, isoLevel string, f func(ctxTX context.Context) error)) *mTransactionManagerMockRunTransaction {
	if mmRunTransaction.mock.inspectFuncRunTransaction != nil {
		mmRunTransaction.mock.t.Fatalf("Inspect function is already set for TransactionManagerMock.RunTransaction")
	}

	mmRunTransaction.mock.inspectFuncRunTransaction = f

	return mmRunTransaction
}

// Return sets up results that will be returned by TransactionManager.RunTransaction
func (mmRunTransaction *mTransactionManagerMockRunTransaction) Return(err error) *TransactionManagerMock {
	if mmRunTransaction.mock.funcRunTransaction != nil {
		mmRunTransaction.mock.t.Fatalf("TransactionManagerMock.RunTransaction mock is already set by Set")
	}

	if mmRunTransaction.defaultExpectation == nil {
		mmRunTransaction.defaultExpectation = &TransactionManagerMockRunTransactionExpectation{mock: mmRunTransaction.mock}
	}
	mmRunTransaction.defaultExpectation.results = &TransactionManagerMockRunTransactionResults{err}
	return mmRunTransaction.mock
}

// Set uses given function f to mock the TransactionManager.RunTransaction method
func (mmRunTransaction *mTransactionManagerMockRunTransaction) Set(f func(ctx context.Context, isoLevel string, f func(ctxTX context.Context) error) (err error)) *TransactionManagerMock {
	if mmRunTransaction.defaultExpectation != nil {
		mmRunTransaction.mock.t.Fatalf("Default expectation is already set for the TransactionManager.RunTransaction method")
	}

	if len(mmRunTransaction.expectations) > 0 {
		mmRunTransaction.mock.t.Fatalf("Some expectations are already set for the TransactionManager.RunTransaction method")
	}

	mmRunTransaction.mock.funcRunTransaction = f
	return mmRunTransaction.mock
}

// When sets expectation for the TransactionManager.RunTransaction which will trigger the result defined by the following
// Then helper
func (mmRunTransaction *mTransactionManagerMockRunTransaction) When(ctx context.Context, isoLevel string, f func(ctxTX context.Context) error) *TransactionManagerMockRunTransactionExpectation {
	if mmRunTransaction.mock.funcRunTransaction != nil {
		mmRunTransaction.mock.t.Fatalf("TransactionManagerMock.RunTransaction mock is already set by Set")
	}

	expectation := &TransactionManagerMockRunTransactionExpectation{
		mock:   mmRunTransaction.mock,
		params: &TransactionManagerMockRunTransactionParams{ctx, isoLevel, f},
	}
	mmRunTransaction.expectations = append(mmRunTransaction.expectations, expectation)
	return expectation
}

// Then sets up TransactionManager.RunTransaction return parameters for the expectation previously defined by the When method
func (e *TransactionManagerMockRunTransactionExpectation) Then(err error) *TransactionManagerMock {
	e.results = &TransactionManagerMockRunTransactionResults{err}
	return e.mock
}

// RunTransaction implements TransactionManager
func (mmRunTransaction *TransactionManagerMock) RunTransaction(ctx context.Context, isoLevel string, f func(ctxTX context.Context) error) (err error) {
	mm_atomic.AddUint64(&mmRunTransaction.beforeRunTransactionCounter, 1)
	defer mm_atomic.AddUint64(&mmRunTransaction.afterRunTransactionCounter, 1)

	if mmRunTransaction.inspectFuncRunTransaction != nil {
		mmRunTransaction.inspectFuncRunTransaction(ctx, isoLevel, f)
	}

	mm_params := &TransactionManagerMockRunTransactionParams{ctx, isoLevel, f}

	// Record call args
	mmRunTransaction.RunTransactionMock.mutex.Lock()
	mmRunTransaction.RunTransactionMock.callArgs = append(mmRunTransaction.RunTransactionMock.callArgs, mm_params)
	mmRunTransaction.RunTransactionMock.mutex.Unlock()

	for _, e := range mmRunTransaction.RunTransactionMock.expectations {
		if minimock.Equal(e.params, mm_params) {
			mm_atomic.AddUint64(&e.Counter, 1)
			return e.results.err
		}
	}

	if mmRunTransaction.RunTransactionMock.defaultExpectation != nil {
		mm_atomic.AddUint64(&mmRunTransaction.RunTransactionMock.defaultExpectation.Counter, 1)
		mm_want := mmRunTransaction.RunTransactionMock.defaultExpectation.params
		mm_got := TransactionManagerMockRunTransactionParams{ctx, isoLevel, f}
		if mm_want != nil && !minimock.Equal(*mm_want, mm_got) {
			mmRunTransaction.t.Errorf("TransactionManagerMock.RunTransaction got unexpected parameters, want: %#v, got: %#v%s\n", *mm_want, mm_got, minimock.Diff(*mm_want, mm_got))
		}

		mm_results := mmRunTransaction.RunTransactionMock.defaultExpectation.results
		if mm_results == nil {
			mmRunTransaction.t.Fatal("No results are set for the TransactionManagerMock.RunTransaction")
		}
		return (*mm_results).err
	}
	if mmRunTransaction.funcRunTransaction != nil {
		return mmRunTransaction.funcRunTransaction(ctx, isoLevel, f)
	}
	mmRunTransaction.t.Fatalf("Unexpected call to TransactionManagerMock.RunTransaction. %v %v %v", ctx, isoLevel, f)
	return
}

// RunTransactionAfterCounter returns a count of finished TransactionManagerMock.RunTransaction invocations
func (mmRunTransaction *TransactionManagerMock) RunTransactionAfterCounter() uint64 {
	return mm_atomic.LoadUint64(&mmRunTransaction.afterRunTransactionCounter)
}

// RunTransactionBeforeCounter returns a count of TransactionManagerMock.RunTransaction invocations
func (mmRunTransaction *TransactionManagerMock) RunTransactionBeforeCounter() uint64 {
	return mm_atomic.LoadUint64(&mmRunTransaction.beforeRunTransactionCounter)
}

// Calls returns a list of arguments used in each call to TransactionManagerMock.RunTransaction.
// The list is in the same order as the calls were made (i.e. recent calls have a higher index)
func (mmRunTransaction *mTransactionManagerMockRunTransaction) Calls() []*TransactionManagerMockRunTransactionParams {
	mmRunTransaction.mutex.RLock()

	argCopy := make([]*TransactionManagerMockRunTransactionParams, len(mmRunTransaction.callArgs))
	copy(argCopy, mmRunTransaction.callArgs)

	mmRunTransaction.mutex.RUnlock()

	return argCopy
}

// MinimockRunTransactionDone returns true if the count of the RunTransaction invocations corresponds
// the number of defined expectations
func (m *TransactionManagerMock) MinimockRunTransactionDone() bool {
	for _, e := range m.RunTransactionMock.expectations {
		if mm_atomic.LoadUint64(&e.Counter) < 1 {
			return false
		}
	}

	// if default expectation was set then invocations count should be greater than zero
	if m.RunTransactionMock.defaultExpectation != nil && mm_atomic.LoadUint64(&m.afterRunTransactionCounter) < 1 {
		return false
	}
	// if func was set then invocations count should be greater than zero
	if m.funcRunTransaction != nil && mm_atomic.LoadUint64(&m.afterRunTransactionCounter) < 1 {
		return false
	}
	return true
}

// MinimockRunTransactionInspect logs each unmet expectation
func (m *TransactionManagerMock) MinimockRunTransactionInspect() {
	for _, e := range m.RunTransactionMock.expectations {
		if mm_atomic.LoadUint64(&e.Counter) < 1 {
			m.t.Errorf("Expected call to TransactionManagerMock.RunTransaction with params: %#v", *e.params)
		}
	}

	// if default expectation was set then invocations count should be greater than zero
	if m.RunTransactionMock.defaultExpectation != nil && mm_atomic.LoadUint64(&m.afterRunTransactionCounter) < 1 {
		if m.RunTransactionMock.defaultExpectation.params == nil {
			m.t.Error("Expected call to TransactionManagerMock.RunTransaction")
		} else {
			m.t.Errorf("Expected call to TransactionManagerMock.RunTransaction with params: %#v", *m.RunTransactionMock.defaultExpectation.params)
		}
	}
	// if func was set then invocations count should be greater than zero
	if m.funcRunTransaction != nil && mm_atomic.LoadUint64(&m.afterRunTransactionCounter) < 1 {
		m.t.Error("Expected call to TransactionManagerMock.RunTransaction")
	}
}

// MinimockFinish checks that all mocked methods have been called the expected number of times
func (m *TransactionManagerMock) MinimockFinish() {
	if !m.minimockDone() {
		m.MinimockRunTransactionInspect()
		m.t.FailNow()
	}
}

// MinimockWait waits for all mocked methods to be called the expected number of times
func (m *TransactionManagerMock) MinimockWait(timeout mm_time.Duration) {
	timeoutCh := mm_time.After(timeout)
	for {
		if m.minimockDone() {
			return
		}
		select {
		case <-timeoutCh:
			m.MinimockFinish()
			return
		case <-mm_time.After(10 * mm_time.Millisecond):
		}
	}
}

func (m *TransactionManagerMock) minimockDone() bool {
	done := true
	return done &&
		m.MinimockRunTransactionDone()
}