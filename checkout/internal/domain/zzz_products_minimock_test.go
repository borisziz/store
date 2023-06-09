package domain

// Code generated by http://github.com/gojuno/minimock (dev). DO NOT EDIT.

//go:generate minimock -i route256/checkout/internal/domain.ProductServiceCaller -o ./zzz_products_minimock_test.go -n ProductServiceCallerMock

import (
	"context"
	"sync"
	mm_atomic "sync/atomic"
	mm_time "time"

	"github.com/gojuno/minimock/v3"
)

// ProductServiceCallerMock implements ProductServiceCaller
type ProductServiceCallerMock struct {
	t minimock.Tester

	funcGetProduct          func(ctx context.Context, sku uint32) (p1 ProductInfo, err error)
	inspectFuncGetProduct   func(ctx context.Context, sku uint32)
	afterGetProductCounter  uint64
	beforeGetProductCounter uint64
	GetProductMock          mProductServiceCallerMockGetProduct

	funcGetSKUs          func(ctx context.Context) (s1 SKUs, err error)
	inspectFuncGetSKUs   func(ctx context.Context)
	afterGetSKUsCounter  uint64
	beforeGetSKUsCounter uint64
	GetSKUsMock          mProductServiceCallerMockGetSKUs
}

// NewProductServiceCallerMock returns a mock for ProductServiceCaller
func NewProductServiceCallerMock(t minimock.Tester) *ProductServiceCallerMock {
	m := &ProductServiceCallerMock{t: t}
	if controller, ok := t.(minimock.MockController); ok {
		controller.RegisterMocker(m)
	}

	m.GetProductMock = mProductServiceCallerMockGetProduct{mock: m}
	m.GetProductMock.callArgs = []*ProductServiceCallerMockGetProductParams{}

	m.GetSKUsMock = mProductServiceCallerMockGetSKUs{mock: m}
	m.GetSKUsMock.callArgs = []*ProductServiceCallerMockGetSKUsParams{}

	return m
}

type mProductServiceCallerMockGetProduct struct {
	mock               *ProductServiceCallerMock
	defaultExpectation *ProductServiceCallerMockGetProductExpectation
	expectations       []*ProductServiceCallerMockGetProductExpectation

	callArgs []*ProductServiceCallerMockGetProductParams
	mutex    sync.RWMutex
}

// ProductServiceCallerMockGetProductExpectation specifies expectation struct of the ProductServiceCaller.GetProduct
type ProductServiceCallerMockGetProductExpectation struct {
	mock    *ProductServiceCallerMock
	params  *ProductServiceCallerMockGetProductParams
	results *ProductServiceCallerMockGetProductResults
	Counter uint64
}

// ProductServiceCallerMockGetProductParams contains parameters of the ProductServiceCaller.GetProduct
type ProductServiceCallerMockGetProductParams struct {
	ctx context.Context
	sku uint32
}

// ProductServiceCallerMockGetProductResults contains results of the ProductServiceCaller.GetProduct
type ProductServiceCallerMockGetProductResults struct {
	p1  ProductInfo
	err error
}

// Expect sets up expected params for ProductServiceCaller.GetProduct
func (mmGetProduct *mProductServiceCallerMockGetProduct) Expect(ctx context.Context, sku uint32) *mProductServiceCallerMockGetProduct {
	if mmGetProduct.mock.funcGetProduct != nil {
		mmGetProduct.mock.t.Fatalf("ProductServiceCallerMock.GetProduct mock is already set by Set")
	}

	if mmGetProduct.defaultExpectation == nil {
		mmGetProduct.defaultExpectation = &ProductServiceCallerMockGetProductExpectation{}
	}

	mmGetProduct.defaultExpectation.params = &ProductServiceCallerMockGetProductParams{ctx, sku}
	for _, e := range mmGetProduct.expectations {
		if minimock.Equal(e.params, mmGetProduct.defaultExpectation.params) {
			mmGetProduct.mock.t.Fatalf("Expectation set by When has same params: %#v", *mmGetProduct.defaultExpectation.params)
		}
	}

	return mmGetProduct
}

// Inspect accepts an inspector function that has same arguments as the ProductServiceCaller.GetProduct
func (mmGetProduct *mProductServiceCallerMockGetProduct) Inspect(f func(ctx context.Context, sku uint32)) *mProductServiceCallerMockGetProduct {
	if mmGetProduct.mock.inspectFuncGetProduct != nil {
		mmGetProduct.mock.t.Fatalf("Inspect function is already set for ProductServiceCallerMock.GetProduct")
	}

	mmGetProduct.mock.inspectFuncGetProduct = f

	return mmGetProduct
}

// Return sets up results that will be returned by ProductServiceCaller.GetProduct
func (mmGetProduct *mProductServiceCallerMockGetProduct) Return(p1 ProductInfo, err error) *ProductServiceCallerMock {
	if mmGetProduct.mock.funcGetProduct != nil {
		mmGetProduct.mock.t.Fatalf("ProductServiceCallerMock.GetProduct mock is already set by Set")
	}

	if mmGetProduct.defaultExpectation == nil {
		mmGetProduct.defaultExpectation = &ProductServiceCallerMockGetProductExpectation{mock: mmGetProduct.mock}
	}
	mmGetProduct.defaultExpectation.results = &ProductServiceCallerMockGetProductResults{p1, err}
	return mmGetProduct.mock
}

// Set uses given function f to mock the ProductServiceCaller.GetProduct method
func (mmGetProduct *mProductServiceCallerMockGetProduct) Set(f func(ctx context.Context, sku uint32) (p1 ProductInfo, err error)) *ProductServiceCallerMock {
	if mmGetProduct.defaultExpectation != nil {
		mmGetProduct.mock.t.Fatalf("Default expectation is already set for the ProductServiceCaller.GetProduct method")
	}

	if len(mmGetProduct.expectations) > 0 {
		mmGetProduct.mock.t.Fatalf("Some expectations are already set for the ProductServiceCaller.GetProduct method")
	}

	mmGetProduct.mock.funcGetProduct = f
	return mmGetProduct.mock
}

// When sets expectation for the ProductServiceCaller.GetProduct which will trigger the result defined by the following
// Then helper
func (mmGetProduct *mProductServiceCallerMockGetProduct) When(ctx context.Context, sku uint32) *ProductServiceCallerMockGetProductExpectation {
	if mmGetProduct.mock.funcGetProduct != nil {
		mmGetProduct.mock.t.Fatalf("ProductServiceCallerMock.GetProduct mock is already set by Set")
	}

	expectation := &ProductServiceCallerMockGetProductExpectation{
		mock:   mmGetProduct.mock,
		params: &ProductServiceCallerMockGetProductParams{ctx, sku},
	}
	mmGetProduct.expectations = append(mmGetProduct.expectations, expectation)
	return expectation
}

// Then sets up ProductServiceCaller.GetProduct return parameters for the expectation previously defined by the When method
func (e *ProductServiceCallerMockGetProductExpectation) Then(p1 ProductInfo, err error) *ProductServiceCallerMock {
	e.results = &ProductServiceCallerMockGetProductResults{p1, err}
	return e.mock
}

// GetProduct implements ProductServiceCaller
func (mmGetProduct *ProductServiceCallerMock) GetProduct(ctx context.Context, sku uint32) (p1 ProductInfo, err error) {
	mm_atomic.AddUint64(&mmGetProduct.beforeGetProductCounter, 1)
	defer mm_atomic.AddUint64(&mmGetProduct.afterGetProductCounter, 1)

	if mmGetProduct.inspectFuncGetProduct != nil {
		mmGetProduct.inspectFuncGetProduct(ctx, sku)
	}

	mm_params := &ProductServiceCallerMockGetProductParams{ctx, sku}

	// Record call args
	mmGetProduct.GetProductMock.mutex.Lock()
	mmGetProduct.GetProductMock.callArgs = append(mmGetProduct.GetProductMock.callArgs, mm_params)
	mmGetProduct.GetProductMock.mutex.Unlock()

	for _, e := range mmGetProduct.GetProductMock.expectations {
		if minimock.Equal(e.params, mm_params) {
			mm_atomic.AddUint64(&e.Counter, 1)
			return e.results.p1, e.results.err
		}
	}

	if mmGetProduct.GetProductMock.defaultExpectation != nil {
		mm_atomic.AddUint64(&mmGetProduct.GetProductMock.defaultExpectation.Counter, 1)
		mm_want := mmGetProduct.GetProductMock.defaultExpectation.params
		mm_got := ProductServiceCallerMockGetProductParams{ctx, sku}
		if mm_want != nil && !minimock.Equal(*mm_want, mm_got) {
			mmGetProduct.t.Errorf("ProductServiceCallerMock.GetProduct got unexpected parameters, want: %#v, got: %#v%s\n", *mm_want, mm_got, minimock.Diff(*mm_want, mm_got))
		}

		mm_results := mmGetProduct.GetProductMock.defaultExpectation.results
		if mm_results == nil {
			mmGetProduct.t.Fatal("No results are set for the ProductServiceCallerMock.GetProduct")
		}
		return (*mm_results).p1, (*mm_results).err
	}
	if mmGetProduct.funcGetProduct != nil {
		return mmGetProduct.funcGetProduct(ctx, sku)
	}
	mmGetProduct.t.Fatalf("Unexpected call to ProductServiceCallerMock.GetProduct. %v %v", ctx, sku)
	return
}

// GetProductAfterCounter returns a count of finished ProductServiceCallerMock.GetProduct invocations
func (mmGetProduct *ProductServiceCallerMock) GetProductAfterCounter() uint64 {
	return mm_atomic.LoadUint64(&mmGetProduct.afterGetProductCounter)
}

// GetProductBeforeCounter returns a count of ProductServiceCallerMock.GetProduct invocations
func (mmGetProduct *ProductServiceCallerMock) GetProductBeforeCounter() uint64 {
	return mm_atomic.LoadUint64(&mmGetProduct.beforeGetProductCounter)
}

// Calls returns a list of arguments used in each call to ProductServiceCallerMock.GetProduct.
// The list is in the same order as the calls were made (i.e. recent calls have a higher index)
func (mmGetProduct *mProductServiceCallerMockGetProduct) Calls() []*ProductServiceCallerMockGetProductParams {
	mmGetProduct.mutex.RLock()

	argCopy := make([]*ProductServiceCallerMockGetProductParams, len(mmGetProduct.callArgs))
	copy(argCopy, mmGetProduct.callArgs)

	mmGetProduct.mutex.RUnlock()

	return argCopy
}

// MinimockGetProductDone returns true if the count of the GetProduct invocations corresponds
// the number of defined expectations
func (m *ProductServiceCallerMock) MinimockGetProductDone() bool {
	for _, e := range m.GetProductMock.expectations {
		if mm_atomic.LoadUint64(&e.Counter) < 1 {
			return false
		}
	}

	// if default expectation was set then invocations count should be greater than zero
	if m.GetProductMock.defaultExpectation != nil && mm_atomic.LoadUint64(&m.afterGetProductCounter) < 1 {
		return false
	}
	// if func was set then invocations count should be greater than zero
	if m.funcGetProduct != nil && mm_atomic.LoadUint64(&m.afterGetProductCounter) < 1 {
		return false
	}
	return true
}

// MinimockGetProductInspect logs each unmet expectation
func (m *ProductServiceCallerMock) MinimockGetProductInspect() {
	for _, e := range m.GetProductMock.expectations {
		if mm_atomic.LoadUint64(&e.Counter) < 1 {
			m.t.Errorf("Expected call to ProductServiceCallerMock.GetProduct with params: %#v", *e.params)
		}
	}

	// if default expectation was set then invocations count should be greater than zero
	if m.GetProductMock.defaultExpectation != nil && mm_atomic.LoadUint64(&m.afterGetProductCounter) < 1 {
		if m.GetProductMock.defaultExpectation.params == nil {
			m.t.Error("Expected call to ProductServiceCallerMock.GetProduct")
		} else {
			m.t.Errorf("Expected call to ProductServiceCallerMock.GetProduct with params: %#v", *m.GetProductMock.defaultExpectation.params)
		}
	}
	// if func was set then invocations count should be greater than zero
	if m.funcGetProduct != nil && mm_atomic.LoadUint64(&m.afterGetProductCounter) < 1 {
		m.t.Error("Expected call to ProductServiceCallerMock.GetProduct")
	}
}

type mProductServiceCallerMockGetSKUs struct {
	mock               *ProductServiceCallerMock
	defaultExpectation *ProductServiceCallerMockGetSKUsExpectation
	expectations       []*ProductServiceCallerMockGetSKUsExpectation

	callArgs []*ProductServiceCallerMockGetSKUsParams
	mutex    sync.RWMutex
}

// ProductServiceCallerMockGetSKUsExpectation specifies expectation struct of the ProductServiceCaller.GetSKUs
type ProductServiceCallerMockGetSKUsExpectation struct {
	mock    *ProductServiceCallerMock
	params  *ProductServiceCallerMockGetSKUsParams
	results *ProductServiceCallerMockGetSKUsResults
	Counter uint64
}

// ProductServiceCallerMockGetSKUsParams contains parameters of the ProductServiceCaller.GetSKUs
type ProductServiceCallerMockGetSKUsParams struct {
	ctx context.Context
}

// ProductServiceCallerMockGetSKUsResults contains results of the ProductServiceCaller.GetSKUs
type ProductServiceCallerMockGetSKUsResults struct {
	s1  SKUs
	err error
}

// Expect sets up expected params for ProductServiceCaller.GetSKUs
func (mmGetSKUs *mProductServiceCallerMockGetSKUs) Expect(ctx context.Context) *mProductServiceCallerMockGetSKUs {
	if mmGetSKUs.mock.funcGetSKUs != nil {
		mmGetSKUs.mock.t.Fatalf("ProductServiceCallerMock.GetSKUs mock is already set by Set")
	}

	if mmGetSKUs.defaultExpectation == nil {
		mmGetSKUs.defaultExpectation = &ProductServiceCallerMockGetSKUsExpectation{}
	}

	mmGetSKUs.defaultExpectation.params = &ProductServiceCallerMockGetSKUsParams{ctx}
	for _, e := range mmGetSKUs.expectations {
		if minimock.Equal(e.params, mmGetSKUs.defaultExpectation.params) {
			mmGetSKUs.mock.t.Fatalf("Expectation set by When has same params: %#v", *mmGetSKUs.defaultExpectation.params)
		}
	}

	return mmGetSKUs
}

// Inspect accepts an inspector function that has same arguments as the ProductServiceCaller.GetSKUs
func (mmGetSKUs *mProductServiceCallerMockGetSKUs) Inspect(f func(ctx context.Context)) *mProductServiceCallerMockGetSKUs {
	if mmGetSKUs.mock.inspectFuncGetSKUs != nil {
		mmGetSKUs.mock.t.Fatalf("Inspect function is already set for ProductServiceCallerMock.GetSKUs")
	}

	mmGetSKUs.mock.inspectFuncGetSKUs = f

	return mmGetSKUs
}

// Return sets up results that will be returned by ProductServiceCaller.GetSKUs
func (mmGetSKUs *mProductServiceCallerMockGetSKUs) Return(s1 SKUs, err error) *ProductServiceCallerMock {
	if mmGetSKUs.mock.funcGetSKUs != nil {
		mmGetSKUs.mock.t.Fatalf("ProductServiceCallerMock.GetSKUs mock is already set by Set")
	}

	if mmGetSKUs.defaultExpectation == nil {
		mmGetSKUs.defaultExpectation = &ProductServiceCallerMockGetSKUsExpectation{mock: mmGetSKUs.mock}
	}
	mmGetSKUs.defaultExpectation.results = &ProductServiceCallerMockGetSKUsResults{s1, err}
	return mmGetSKUs.mock
}

// Set uses given function f to mock the ProductServiceCaller.GetSKUs method
func (mmGetSKUs *mProductServiceCallerMockGetSKUs) Set(f func(ctx context.Context) (s1 SKUs, err error)) *ProductServiceCallerMock {
	if mmGetSKUs.defaultExpectation != nil {
		mmGetSKUs.mock.t.Fatalf("Default expectation is already set for the ProductServiceCaller.GetSKUs method")
	}

	if len(mmGetSKUs.expectations) > 0 {
		mmGetSKUs.mock.t.Fatalf("Some expectations are already set for the ProductServiceCaller.GetSKUs method")
	}

	mmGetSKUs.mock.funcGetSKUs = f
	return mmGetSKUs.mock
}

// When sets expectation for the ProductServiceCaller.GetSKUs which will trigger the result defined by the following
// Then helper
func (mmGetSKUs *mProductServiceCallerMockGetSKUs) When(ctx context.Context) *ProductServiceCallerMockGetSKUsExpectation {
	if mmGetSKUs.mock.funcGetSKUs != nil {
		mmGetSKUs.mock.t.Fatalf("ProductServiceCallerMock.GetSKUs mock is already set by Set")
	}

	expectation := &ProductServiceCallerMockGetSKUsExpectation{
		mock:   mmGetSKUs.mock,
		params: &ProductServiceCallerMockGetSKUsParams{ctx},
	}
	mmGetSKUs.expectations = append(mmGetSKUs.expectations, expectation)
	return expectation
}

// Then sets up ProductServiceCaller.GetSKUs return parameters for the expectation previously defined by the When method
func (e *ProductServiceCallerMockGetSKUsExpectation) Then(s1 SKUs, err error) *ProductServiceCallerMock {
	e.results = &ProductServiceCallerMockGetSKUsResults{s1, err}
	return e.mock
}

// GetSKUs implements ProductServiceCaller
func (mmGetSKUs *ProductServiceCallerMock) GetSKUs(ctx context.Context) (s1 SKUs, err error) {
	mm_atomic.AddUint64(&mmGetSKUs.beforeGetSKUsCounter, 1)
	defer mm_atomic.AddUint64(&mmGetSKUs.afterGetSKUsCounter, 1)

	if mmGetSKUs.inspectFuncGetSKUs != nil {
		mmGetSKUs.inspectFuncGetSKUs(ctx)
	}

	mm_params := &ProductServiceCallerMockGetSKUsParams{ctx}

	// Record call args
	mmGetSKUs.GetSKUsMock.mutex.Lock()
	mmGetSKUs.GetSKUsMock.callArgs = append(mmGetSKUs.GetSKUsMock.callArgs, mm_params)
	mmGetSKUs.GetSKUsMock.mutex.Unlock()

	for _, e := range mmGetSKUs.GetSKUsMock.expectations {
		if minimock.Equal(e.params, mm_params) {
			mm_atomic.AddUint64(&e.Counter, 1)
			return e.results.s1, e.results.err
		}
	}

	if mmGetSKUs.GetSKUsMock.defaultExpectation != nil {
		mm_atomic.AddUint64(&mmGetSKUs.GetSKUsMock.defaultExpectation.Counter, 1)
		mm_want := mmGetSKUs.GetSKUsMock.defaultExpectation.params
		mm_got := ProductServiceCallerMockGetSKUsParams{ctx}
		if mm_want != nil && !minimock.Equal(*mm_want, mm_got) {
			mmGetSKUs.t.Errorf("ProductServiceCallerMock.GetSKUs got unexpected parameters, want: %#v, got: %#v%s\n", *mm_want, mm_got, minimock.Diff(*mm_want, mm_got))
		}

		mm_results := mmGetSKUs.GetSKUsMock.defaultExpectation.results
		if mm_results == nil {
			mmGetSKUs.t.Fatal("No results are set for the ProductServiceCallerMock.GetSKUs")
		}
		return (*mm_results).s1, (*mm_results).err
	}
	if mmGetSKUs.funcGetSKUs != nil {
		return mmGetSKUs.funcGetSKUs(ctx)
	}
	mmGetSKUs.t.Fatalf("Unexpected call to ProductServiceCallerMock.GetSKUs. %v", ctx)
	return
}

// GetSKUsAfterCounter returns a count of finished ProductServiceCallerMock.GetSKUs invocations
func (mmGetSKUs *ProductServiceCallerMock) GetSKUsAfterCounter() uint64 {
	return mm_atomic.LoadUint64(&mmGetSKUs.afterGetSKUsCounter)
}

// GetSKUsBeforeCounter returns a count of ProductServiceCallerMock.GetSKUs invocations
func (mmGetSKUs *ProductServiceCallerMock) GetSKUsBeforeCounter() uint64 {
	return mm_atomic.LoadUint64(&mmGetSKUs.beforeGetSKUsCounter)
}

// Calls returns a list of arguments used in each call to ProductServiceCallerMock.GetSKUs.
// The list is in the same order as the calls were made (i.e. recent calls have a higher index)
func (mmGetSKUs *mProductServiceCallerMockGetSKUs) Calls() []*ProductServiceCallerMockGetSKUsParams {
	mmGetSKUs.mutex.RLock()

	argCopy := make([]*ProductServiceCallerMockGetSKUsParams, len(mmGetSKUs.callArgs))
	copy(argCopy, mmGetSKUs.callArgs)

	mmGetSKUs.mutex.RUnlock()

	return argCopy
}

// MinimockGetSKUsDone returns true if the count of the GetSKUs invocations corresponds
// the number of defined expectations
func (m *ProductServiceCallerMock) MinimockGetSKUsDone() bool {
	for _, e := range m.GetSKUsMock.expectations {
		if mm_atomic.LoadUint64(&e.Counter) < 1 {
			return false
		}
	}

	// if default expectation was set then invocations count should be greater than zero
	if m.GetSKUsMock.defaultExpectation != nil && mm_atomic.LoadUint64(&m.afterGetSKUsCounter) < 1 {
		return false
	}
	// if func was set then invocations count should be greater than zero
	if m.funcGetSKUs != nil && mm_atomic.LoadUint64(&m.afterGetSKUsCounter) < 1 {
		return false
	}
	return true
}

// MinimockGetSKUsInspect logs each unmet expectation
func (m *ProductServiceCallerMock) MinimockGetSKUsInspect() {
	for _, e := range m.GetSKUsMock.expectations {
		if mm_atomic.LoadUint64(&e.Counter) < 1 {
			m.t.Errorf("Expected call to ProductServiceCallerMock.GetSKUs with params: %#v", *e.params)
		}
	}

	// if default expectation was set then invocations count should be greater than zero
	if m.GetSKUsMock.defaultExpectation != nil && mm_atomic.LoadUint64(&m.afterGetSKUsCounter) < 1 {
		if m.GetSKUsMock.defaultExpectation.params == nil {
			m.t.Error("Expected call to ProductServiceCallerMock.GetSKUs")
		} else {
			m.t.Errorf("Expected call to ProductServiceCallerMock.GetSKUs with params: %#v", *m.GetSKUsMock.defaultExpectation.params)
		}
	}
	// if func was set then invocations count should be greater than zero
	if m.funcGetSKUs != nil && mm_atomic.LoadUint64(&m.afterGetSKUsCounter) < 1 {
		m.t.Error("Expected call to ProductServiceCallerMock.GetSKUs")
	}
}

// MinimockFinish checks that all mocked methods have been called the expected number of times
func (m *ProductServiceCallerMock) MinimockFinish() {
	if !m.minimockDone() {
		m.MinimockGetProductInspect()

		m.MinimockGetSKUsInspect()
		m.t.FailNow()
	}
}

// MinimockWait waits for all mocked methods to be called the expected number of times
func (m *ProductServiceCallerMock) MinimockWait(timeout mm_time.Duration) {
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

func (m *ProductServiceCallerMock) minimockDone() bool {
	done := true
	return done &&
		m.MinimockGetProductDone() &&
		m.MinimockGetSKUsDone()
}
