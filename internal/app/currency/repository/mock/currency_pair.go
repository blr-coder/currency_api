package mock

// Code generated by http://github.com/gojuno/minimock (dev). DO NOT EDIT.

import (
	"context"
	"currency_api/internal/app/currency/models"
	"sync"
	mm_atomic "sync/atomic"
	mm_time "time"

	"github.com/gojuno/minimock/v3"
)

// CurrencyPairMock implements repository.CurrencyPair
type CurrencyPairMock struct {
	t minimock.Tester

	funcCreate          func(ctx context.Context, pair *models.CurrencyPairCreateInput) (cp1 *models.CurrencyPair, err error)
	inspectFuncCreate   func(ctx context.Context, pair *models.CurrencyPairCreateInput)
	afterCreateCounter  uint64
	beforeCreateCounter uint64
	CreateMock          mCurrencyPairMockCreate

	funcGet          func(ctx context.Context, f string, t string) (cp1 *models.CurrencyPair, err error)
	inspectFuncGet   func(ctx context.Context, f string, t string)
	afterGetCounter  uint64
	beforeGetCounter uint64
	GetMock          mCurrencyPairMockGet

	funcList          func(ctx context.Context) (c2 models.CurrencyPairs, err error)
	inspectFuncList   func(ctx context.Context)
	afterListCounter  uint64
	beforeListCounter uint64
	ListMock          mCurrencyPairMockList

	funcUpdatePair          func(ctx context.Context, pair *models.CurrencyPair) (err error)
	inspectFuncUpdatePair   func(ctx context.Context, pair *models.CurrencyPair)
	afterUpdatePairCounter  uint64
	beforeUpdatePairCounter uint64
	UpdatePairMock          mCurrencyPairMockUpdatePair
}

// NewCurrencyPairMock returns a mock for repository.CurrencyPair
func NewCurrencyPairMock(t minimock.Tester) *CurrencyPairMock {
	m := &CurrencyPairMock{t: t}
	if controller, ok := t.(minimock.MockController); ok {
		controller.RegisterMocker(m)
	}

	m.CreateMock = mCurrencyPairMockCreate{mock: m}
	m.CreateMock.callArgs = []*CurrencyPairMockCreateParams{}

	m.GetMock = mCurrencyPairMockGet{mock: m}
	m.GetMock.callArgs = []*CurrencyPairMockGetParams{}

	m.ListMock = mCurrencyPairMockList{mock: m}
	m.ListMock.callArgs = []*CurrencyPairMockListParams{}

	m.UpdatePairMock = mCurrencyPairMockUpdatePair{mock: m}
	m.UpdatePairMock.callArgs = []*CurrencyPairMockUpdatePairParams{}

	return m
}

type mCurrencyPairMockCreate struct {
	mock               *CurrencyPairMock
	defaultExpectation *CurrencyPairMockCreateExpectation
	expectations       []*CurrencyPairMockCreateExpectation

	callArgs []*CurrencyPairMockCreateParams
	mutex    sync.RWMutex
}

// CurrencyPairMockCreateExpectation specifies expectation struct of the CurrencyPair.Create
type CurrencyPairMockCreateExpectation struct {
	mock    *CurrencyPairMock
	params  *CurrencyPairMockCreateParams
	results *CurrencyPairMockCreateResults
	Counter uint64
}

// CurrencyPairMockCreateParams contains parameters of the CurrencyPair.Create
type CurrencyPairMockCreateParams struct {
	ctx  context.Context
	pair *models.CurrencyPairCreateInput
}

// CurrencyPairMockCreateResults contains results of the CurrencyPair.Create
type CurrencyPairMockCreateResults struct {
	cp1 *models.CurrencyPair
	err error
}

// Expect sets up expected params for CurrencyPair.Create
func (mmCreate *mCurrencyPairMockCreate) Expect(ctx context.Context, pair *models.CurrencyPairCreateInput) *mCurrencyPairMockCreate {
	if mmCreate.mock.funcCreate != nil {
		mmCreate.mock.t.Fatalf("CurrencyPairMock.Create mock is already set by Set")
	}

	if mmCreate.defaultExpectation == nil {
		mmCreate.defaultExpectation = &CurrencyPairMockCreateExpectation{}
	}

	mmCreate.defaultExpectation.params = &CurrencyPairMockCreateParams{ctx, pair}
	for _, e := range mmCreate.expectations {
		if minimock.Equal(e.params, mmCreate.defaultExpectation.params) {
			mmCreate.mock.t.Fatalf("Expectation set by When has same params: %#v", *mmCreate.defaultExpectation.params)
		}
	}

	return mmCreate
}

// Inspect accepts an inspector function that has same arguments as the CurrencyPair.Create
func (mmCreate *mCurrencyPairMockCreate) Inspect(f func(ctx context.Context, pair *models.CurrencyPairCreateInput)) *mCurrencyPairMockCreate {
	if mmCreate.mock.inspectFuncCreate != nil {
		mmCreate.mock.t.Fatalf("Inspect function is already set for CurrencyPairMock.Create")
	}

	mmCreate.mock.inspectFuncCreate = f

	return mmCreate
}

// Return sets up results that will be returned by CurrencyPair.Create
func (mmCreate *mCurrencyPairMockCreate) Return(cp1 *models.CurrencyPair, err error) *CurrencyPairMock {
	if mmCreate.mock.funcCreate != nil {
		mmCreate.mock.t.Fatalf("CurrencyPairMock.Create mock is already set by Set")
	}

	if mmCreate.defaultExpectation == nil {
		mmCreate.defaultExpectation = &CurrencyPairMockCreateExpectation{mock: mmCreate.mock}
	}
	mmCreate.defaultExpectation.results = &CurrencyPairMockCreateResults{cp1, err}
	return mmCreate.mock
}

// Set uses given function f to mock the CurrencyPair.Create method
func (mmCreate *mCurrencyPairMockCreate) Set(f func(ctx context.Context, pair *models.CurrencyPairCreateInput) (cp1 *models.CurrencyPair, err error)) *CurrencyPairMock {
	if mmCreate.defaultExpectation != nil {
		mmCreate.mock.t.Fatalf("Default expectation is already set for the CurrencyPair.Create method")
	}

	if len(mmCreate.expectations) > 0 {
		mmCreate.mock.t.Fatalf("Some expectations are already set for the CurrencyPair.Create method")
	}

	mmCreate.mock.funcCreate = f
	return mmCreate.mock
}

// When sets expectation for the CurrencyPair.Create which will trigger the result defined by the following
// Then helper
func (mmCreate *mCurrencyPairMockCreate) When(ctx context.Context, pair *models.CurrencyPairCreateInput) *CurrencyPairMockCreateExpectation {
	if mmCreate.mock.funcCreate != nil {
		mmCreate.mock.t.Fatalf("CurrencyPairMock.Create mock is already set by Set")
	}

	expectation := &CurrencyPairMockCreateExpectation{
		mock:   mmCreate.mock,
		params: &CurrencyPairMockCreateParams{ctx, pair},
	}
	mmCreate.expectations = append(mmCreate.expectations, expectation)
	return expectation
}

// Then sets up CurrencyPair.Create return parameters for the expectation previously defined by the When method
func (e *CurrencyPairMockCreateExpectation) Then(cp1 *models.CurrencyPair, err error) *CurrencyPairMock {
	e.results = &CurrencyPairMockCreateResults{cp1, err}
	return e.mock
}

// Create implements repository.CurrencyPair
func (mmCreate *CurrencyPairMock) Create(ctx context.Context, pair *models.CurrencyPairCreateInput) (cp1 *models.CurrencyPair, err error) {
	mm_atomic.AddUint64(&mmCreate.beforeCreateCounter, 1)
	defer mm_atomic.AddUint64(&mmCreate.afterCreateCounter, 1)

	if mmCreate.inspectFuncCreate != nil {
		mmCreate.inspectFuncCreate(ctx, pair)
	}

	mm_params := &CurrencyPairMockCreateParams{ctx, pair}

	// Record call args
	mmCreate.CreateMock.mutex.Lock()
	mmCreate.CreateMock.callArgs = append(mmCreate.CreateMock.callArgs, mm_params)
	mmCreate.CreateMock.mutex.Unlock()

	for _, e := range mmCreate.CreateMock.expectations {
		if minimock.Equal(e.params, mm_params) {
			mm_atomic.AddUint64(&e.Counter, 1)
			return e.results.cp1, e.results.err
		}
	}

	if mmCreate.CreateMock.defaultExpectation != nil {
		mm_atomic.AddUint64(&mmCreate.CreateMock.defaultExpectation.Counter, 1)
		mm_want := mmCreate.CreateMock.defaultExpectation.params
		mm_got := CurrencyPairMockCreateParams{ctx, pair}
		if mm_want != nil && !minimock.Equal(*mm_want, mm_got) {
			mmCreate.t.Errorf("CurrencyPairMock.Create got unexpected parameters, want: %#v, got: %#v%s\n", *mm_want, mm_got, minimock.Diff(*mm_want, mm_got))
		}

		mm_results := mmCreate.CreateMock.defaultExpectation.results
		if mm_results == nil {
			mmCreate.t.Fatal("No results are set for the CurrencyPairMock.Create")
		}
		return (*mm_results).cp1, (*mm_results).err
	}
	if mmCreate.funcCreate != nil {
		return mmCreate.funcCreate(ctx, pair)
	}
	mmCreate.t.Fatalf("Unexpected call to CurrencyPairMock.Create. %v %v", ctx, pair)
	return
}

// CreateAfterCounter returns a count of finished CurrencyPairMock.Create invocations
func (mmCreate *CurrencyPairMock) CreateAfterCounter() uint64 {
	return mm_atomic.LoadUint64(&mmCreate.afterCreateCounter)
}

// CreateBeforeCounter returns a count of CurrencyPairMock.Create invocations
func (mmCreate *CurrencyPairMock) CreateBeforeCounter() uint64 {
	return mm_atomic.LoadUint64(&mmCreate.beforeCreateCounter)
}

// Calls returns a list of arguments used in each call to CurrencyPairMock.Create.
// The list is in the same order as the calls were made (i.e. recent calls have a higher index)
func (mmCreate *mCurrencyPairMockCreate) Calls() []*CurrencyPairMockCreateParams {
	mmCreate.mutex.RLock()

	argCopy := make([]*CurrencyPairMockCreateParams, len(mmCreate.callArgs))
	copy(argCopy, mmCreate.callArgs)

	mmCreate.mutex.RUnlock()

	return argCopy
}

// MinimockCreateDone returns true if the count of the Create invocations corresponds
// the number of defined expectations
func (m *CurrencyPairMock) MinimockCreateDone() bool {
	for _, e := range m.CreateMock.expectations {
		if mm_atomic.LoadUint64(&e.Counter) < 1 {
			return false
		}
	}

	// if default expectation was set then invocations count should be greater than zero
	if m.CreateMock.defaultExpectation != nil && mm_atomic.LoadUint64(&m.afterCreateCounter) < 1 {
		return false
	}
	// if func was set then invocations count should be greater than zero
	if m.funcCreate != nil && mm_atomic.LoadUint64(&m.afterCreateCounter) < 1 {
		return false
	}
	return true
}

// MinimockCreateInspect logs each unmet expectation
func (m *CurrencyPairMock) MinimockCreateInspect() {
	for _, e := range m.CreateMock.expectations {
		if mm_atomic.LoadUint64(&e.Counter) < 1 {
			m.t.Errorf("Expected call to CurrencyPairMock.Create with params: %#v", *e.params)
		}
	}

	// if default expectation was set then invocations count should be greater than zero
	if m.CreateMock.defaultExpectation != nil && mm_atomic.LoadUint64(&m.afterCreateCounter) < 1 {
		if m.CreateMock.defaultExpectation.params == nil {
			m.t.Error("Expected call to CurrencyPairMock.Create")
		} else {
			m.t.Errorf("Expected call to CurrencyPairMock.Create with params: %#v", *m.CreateMock.defaultExpectation.params)
		}
	}
	// if func was set then invocations count should be greater than zero
	if m.funcCreate != nil && mm_atomic.LoadUint64(&m.afterCreateCounter) < 1 {
		m.t.Error("Expected call to CurrencyPairMock.Create")
	}
}

type mCurrencyPairMockGet struct {
	mock               *CurrencyPairMock
	defaultExpectation *CurrencyPairMockGetExpectation
	expectations       []*CurrencyPairMockGetExpectation

	callArgs []*CurrencyPairMockGetParams
	mutex    sync.RWMutex
}

// CurrencyPairMockGetExpectation specifies expectation struct of the CurrencyPair.Get
type CurrencyPairMockGetExpectation struct {
	mock    *CurrencyPairMock
	params  *CurrencyPairMockGetParams
	results *CurrencyPairMockGetResults
	Counter uint64
}

// CurrencyPairMockGetParams contains parameters of the CurrencyPair.Get
type CurrencyPairMockGetParams struct {
	ctx context.Context
	f   string
	t   string
}

// CurrencyPairMockGetResults contains results of the CurrencyPair.Get
type CurrencyPairMockGetResults struct {
	cp1 *models.CurrencyPair
	err error
}

// Expect sets up expected params for CurrencyPair.Get
func (mmGet *mCurrencyPairMockGet) Expect(ctx context.Context, f string, t string) *mCurrencyPairMockGet {
	if mmGet.mock.funcGet != nil {
		mmGet.mock.t.Fatalf("CurrencyPairMock.Get mock is already set by Set")
	}

	if mmGet.defaultExpectation == nil {
		mmGet.defaultExpectation = &CurrencyPairMockGetExpectation{}
	}

	mmGet.defaultExpectation.params = &CurrencyPairMockGetParams{ctx, f, t}
	for _, e := range mmGet.expectations {
		if minimock.Equal(e.params, mmGet.defaultExpectation.params) {
			mmGet.mock.t.Fatalf("Expectation set by When has same params: %#v", *mmGet.defaultExpectation.params)
		}
	}

	return mmGet
}

// Inspect accepts an inspector function that has same arguments as the CurrencyPair.Get
func (mmGet *mCurrencyPairMockGet) Inspect(f func(ctx context.Context, f string, t string)) *mCurrencyPairMockGet {
	if mmGet.mock.inspectFuncGet != nil {
		mmGet.mock.t.Fatalf("Inspect function is already set for CurrencyPairMock.Get")
	}

	mmGet.mock.inspectFuncGet = f

	return mmGet
}

// Return sets up results that will be returned by CurrencyPair.Get
func (mmGet *mCurrencyPairMockGet) Return(cp1 *models.CurrencyPair, err error) *CurrencyPairMock {
	if mmGet.mock.funcGet != nil {
		mmGet.mock.t.Fatalf("CurrencyPairMock.Get mock is already set by Set")
	}

	if mmGet.defaultExpectation == nil {
		mmGet.defaultExpectation = &CurrencyPairMockGetExpectation{mock: mmGet.mock}
	}
	mmGet.defaultExpectation.results = &CurrencyPairMockGetResults{cp1, err}
	return mmGet.mock
}

// Set uses given function f to mock the CurrencyPair.Get method
func (mmGet *mCurrencyPairMockGet) Set(f func(ctx context.Context, f string, t string) (cp1 *models.CurrencyPair, err error)) *CurrencyPairMock {
	if mmGet.defaultExpectation != nil {
		mmGet.mock.t.Fatalf("Default expectation is already set for the CurrencyPair.Get method")
	}

	if len(mmGet.expectations) > 0 {
		mmGet.mock.t.Fatalf("Some expectations are already set for the CurrencyPair.Get method")
	}

	mmGet.mock.funcGet = f
	return mmGet.mock
}

// When sets expectation for the CurrencyPair.Get which will trigger the result defined by the following
// Then helper
func (mmGet *mCurrencyPairMockGet) When(ctx context.Context, f string, t string) *CurrencyPairMockGetExpectation {
	if mmGet.mock.funcGet != nil {
		mmGet.mock.t.Fatalf("CurrencyPairMock.Get mock is already set by Set")
	}

	expectation := &CurrencyPairMockGetExpectation{
		mock:   mmGet.mock,
		params: &CurrencyPairMockGetParams{ctx, f, t},
	}
	mmGet.expectations = append(mmGet.expectations, expectation)
	return expectation
}

// Then sets up CurrencyPair.Get return parameters for the expectation previously defined by the When method
func (e *CurrencyPairMockGetExpectation) Then(cp1 *models.CurrencyPair, err error) *CurrencyPairMock {
	e.results = &CurrencyPairMockGetResults{cp1, err}
	return e.mock
}

// Get implements repository.CurrencyPair
func (mmGet *CurrencyPairMock) Get(ctx context.Context, f string, t string) (cp1 *models.CurrencyPair, err error) {
	mm_atomic.AddUint64(&mmGet.beforeGetCounter, 1)
	defer mm_atomic.AddUint64(&mmGet.afterGetCounter, 1)

	if mmGet.inspectFuncGet != nil {
		mmGet.inspectFuncGet(ctx, f, t)
	}

	mm_params := &CurrencyPairMockGetParams{ctx, f, t}

	// Record call args
	mmGet.GetMock.mutex.Lock()
	mmGet.GetMock.callArgs = append(mmGet.GetMock.callArgs, mm_params)
	mmGet.GetMock.mutex.Unlock()

	for _, e := range mmGet.GetMock.expectations {
		if minimock.Equal(e.params, mm_params) {
			mm_atomic.AddUint64(&e.Counter, 1)
			return e.results.cp1, e.results.err
		}
	}

	if mmGet.GetMock.defaultExpectation != nil {
		mm_atomic.AddUint64(&mmGet.GetMock.defaultExpectation.Counter, 1)
		mm_want := mmGet.GetMock.defaultExpectation.params
		mm_got := CurrencyPairMockGetParams{ctx, f, t}
		if mm_want != nil && !minimock.Equal(*mm_want, mm_got) {
			mmGet.t.Errorf("CurrencyPairMock.Get got unexpected parameters, want: %#v, got: %#v%s\n", *mm_want, mm_got, minimock.Diff(*mm_want, mm_got))
		}

		mm_results := mmGet.GetMock.defaultExpectation.results
		if mm_results == nil {
			mmGet.t.Fatal("No results are set for the CurrencyPairMock.Get")
		}
		return (*mm_results).cp1, (*mm_results).err
	}
	if mmGet.funcGet != nil {
		return mmGet.funcGet(ctx, f, t)
	}
	mmGet.t.Fatalf("Unexpected call to CurrencyPairMock.Get. %v %v %v", ctx, f, t)
	return
}

// GetAfterCounter returns a count of finished CurrencyPairMock.Get invocations
func (mmGet *CurrencyPairMock) GetAfterCounter() uint64 {
	return mm_atomic.LoadUint64(&mmGet.afterGetCounter)
}

// GetBeforeCounter returns a count of CurrencyPairMock.Get invocations
func (mmGet *CurrencyPairMock) GetBeforeCounter() uint64 {
	return mm_atomic.LoadUint64(&mmGet.beforeGetCounter)
}

// Calls returns a list of arguments used in each call to CurrencyPairMock.Get.
// The list is in the same order as the calls were made (i.e. recent calls have a higher index)
func (mmGet *mCurrencyPairMockGet) Calls() []*CurrencyPairMockGetParams {
	mmGet.mutex.RLock()

	argCopy := make([]*CurrencyPairMockGetParams, len(mmGet.callArgs))
	copy(argCopy, mmGet.callArgs)

	mmGet.mutex.RUnlock()

	return argCopy
}

// MinimockGetDone returns true if the count of the Get invocations corresponds
// the number of defined expectations
func (m *CurrencyPairMock) MinimockGetDone() bool {
	for _, e := range m.GetMock.expectations {
		if mm_atomic.LoadUint64(&e.Counter) < 1 {
			return false
		}
	}

	// if default expectation was set then invocations count should be greater than zero
	if m.GetMock.defaultExpectation != nil && mm_atomic.LoadUint64(&m.afterGetCounter) < 1 {
		return false
	}
	// if func was set then invocations count should be greater than zero
	if m.funcGet != nil && mm_atomic.LoadUint64(&m.afterGetCounter) < 1 {
		return false
	}
	return true
}

// MinimockGetInspect logs each unmet expectation
func (m *CurrencyPairMock) MinimockGetInspect() {
	for _, e := range m.GetMock.expectations {
		if mm_atomic.LoadUint64(&e.Counter) < 1 {
			m.t.Errorf("Expected call to CurrencyPairMock.Get with params: %#v", *e.params)
		}
	}

	// if default expectation was set then invocations count should be greater than zero
	if m.GetMock.defaultExpectation != nil && mm_atomic.LoadUint64(&m.afterGetCounter) < 1 {
		if m.GetMock.defaultExpectation.params == nil {
			m.t.Error("Expected call to CurrencyPairMock.Get")
		} else {
			m.t.Errorf("Expected call to CurrencyPairMock.Get with params: %#v", *m.GetMock.defaultExpectation.params)
		}
	}
	// if func was set then invocations count should be greater than zero
	if m.funcGet != nil && mm_atomic.LoadUint64(&m.afterGetCounter) < 1 {
		m.t.Error("Expected call to CurrencyPairMock.Get")
	}
}

type mCurrencyPairMockList struct {
	mock               *CurrencyPairMock
	defaultExpectation *CurrencyPairMockListExpectation
	expectations       []*CurrencyPairMockListExpectation

	callArgs []*CurrencyPairMockListParams
	mutex    sync.RWMutex
}

// CurrencyPairMockListExpectation specifies expectation struct of the CurrencyPair.List
type CurrencyPairMockListExpectation struct {
	mock    *CurrencyPairMock
	params  *CurrencyPairMockListParams
	results *CurrencyPairMockListResults
	Counter uint64
}

// CurrencyPairMockListParams contains parameters of the CurrencyPair.List
type CurrencyPairMockListParams struct {
	ctx context.Context
}

// CurrencyPairMockListResults contains results of the CurrencyPair.List
type CurrencyPairMockListResults struct {
	c2  models.CurrencyPairs
	err error
}

// Expect sets up expected params for CurrencyPair.List
func (mmList *mCurrencyPairMockList) Expect(ctx context.Context) *mCurrencyPairMockList {
	if mmList.mock.funcList != nil {
		mmList.mock.t.Fatalf("CurrencyPairMock.List mock is already set by Set")
	}

	if mmList.defaultExpectation == nil {
		mmList.defaultExpectation = &CurrencyPairMockListExpectation{}
	}

	mmList.defaultExpectation.params = &CurrencyPairMockListParams{ctx}
	for _, e := range mmList.expectations {
		if minimock.Equal(e.params, mmList.defaultExpectation.params) {
			mmList.mock.t.Fatalf("Expectation set by When has same params: %#v", *mmList.defaultExpectation.params)
		}
	}

	return mmList
}

// Inspect accepts an inspector function that has same arguments as the CurrencyPair.List
func (mmList *mCurrencyPairMockList) Inspect(f func(ctx context.Context)) *mCurrencyPairMockList {
	if mmList.mock.inspectFuncList != nil {
		mmList.mock.t.Fatalf("Inspect function is already set for CurrencyPairMock.List")
	}

	mmList.mock.inspectFuncList = f

	return mmList
}

// Return sets up results that will be returned by CurrencyPair.List
func (mmList *mCurrencyPairMockList) Return(c2 models.CurrencyPairs, err error) *CurrencyPairMock {
	if mmList.mock.funcList != nil {
		mmList.mock.t.Fatalf("CurrencyPairMock.List mock is already set by Set")
	}

	if mmList.defaultExpectation == nil {
		mmList.defaultExpectation = &CurrencyPairMockListExpectation{mock: mmList.mock}
	}
	mmList.defaultExpectation.results = &CurrencyPairMockListResults{c2, err}
	return mmList.mock
}

// Set uses given function f to mock the CurrencyPair.List method
func (mmList *mCurrencyPairMockList) Set(f func(ctx context.Context) (c2 models.CurrencyPairs, err error)) *CurrencyPairMock {
	if mmList.defaultExpectation != nil {
		mmList.mock.t.Fatalf("Default expectation is already set for the CurrencyPair.List method")
	}

	if len(mmList.expectations) > 0 {
		mmList.mock.t.Fatalf("Some expectations are already set for the CurrencyPair.List method")
	}

	mmList.mock.funcList = f
	return mmList.mock
}

// When sets expectation for the CurrencyPair.List which will trigger the result defined by the following
// Then helper
func (mmList *mCurrencyPairMockList) When(ctx context.Context) *CurrencyPairMockListExpectation {
	if mmList.mock.funcList != nil {
		mmList.mock.t.Fatalf("CurrencyPairMock.List mock is already set by Set")
	}

	expectation := &CurrencyPairMockListExpectation{
		mock:   mmList.mock,
		params: &CurrencyPairMockListParams{ctx},
	}
	mmList.expectations = append(mmList.expectations, expectation)
	return expectation
}

// Then sets up CurrencyPair.List return parameters for the expectation previously defined by the When method
func (e *CurrencyPairMockListExpectation) Then(c2 models.CurrencyPairs, err error) *CurrencyPairMock {
	e.results = &CurrencyPairMockListResults{c2, err}
	return e.mock
}

// List implements repository.CurrencyPair
func (mmList *CurrencyPairMock) List(ctx context.Context) (c2 models.CurrencyPairs, err error) {
	mm_atomic.AddUint64(&mmList.beforeListCounter, 1)
	defer mm_atomic.AddUint64(&mmList.afterListCounter, 1)

	if mmList.inspectFuncList != nil {
		mmList.inspectFuncList(ctx)
	}

	mm_params := &CurrencyPairMockListParams{ctx}

	// Record call args
	mmList.ListMock.mutex.Lock()
	mmList.ListMock.callArgs = append(mmList.ListMock.callArgs, mm_params)
	mmList.ListMock.mutex.Unlock()

	for _, e := range mmList.ListMock.expectations {
		if minimock.Equal(e.params, mm_params) {
			mm_atomic.AddUint64(&e.Counter, 1)
			return e.results.c2, e.results.err
		}
	}

	if mmList.ListMock.defaultExpectation != nil {
		mm_atomic.AddUint64(&mmList.ListMock.defaultExpectation.Counter, 1)
		mm_want := mmList.ListMock.defaultExpectation.params
		mm_got := CurrencyPairMockListParams{ctx}
		if mm_want != nil && !minimock.Equal(*mm_want, mm_got) {
			mmList.t.Errorf("CurrencyPairMock.List got unexpected parameters, want: %#v, got: %#v%s\n", *mm_want, mm_got, minimock.Diff(*mm_want, mm_got))
		}

		mm_results := mmList.ListMock.defaultExpectation.results
		if mm_results == nil {
			mmList.t.Fatal("No results are set for the CurrencyPairMock.List")
		}
		return (*mm_results).c2, (*mm_results).err
	}
	if mmList.funcList != nil {
		return mmList.funcList(ctx)
	}
	mmList.t.Fatalf("Unexpected call to CurrencyPairMock.List. %v", ctx)
	return
}

// ListAfterCounter returns a count of finished CurrencyPairMock.List invocations
func (mmList *CurrencyPairMock) ListAfterCounter() uint64 {
	return mm_atomic.LoadUint64(&mmList.afterListCounter)
}

// ListBeforeCounter returns a count of CurrencyPairMock.List invocations
func (mmList *CurrencyPairMock) ListBeforeCounter() uint64 {
	return mm_atomic.LoadUint64(&mmList.beforeListCounter)
}

// Calls returns a list of arguments used in each call to CurrencyPairMock.List.
// The list is in the same order as the calls were made (i.e. recent calls have a higher index)
func (mmList *mCurrencyPairMockList) Calls() []*CurrencyPairMockListParams {
	mmList.mutex.RLock()

	argCopy := make([]*CurrencyPairMockListParams, len(mmList.callArgs))
	copy(argCopy, mmList.callArgs)

	mmList.mutex.RUnlock()

	return argCopy
}

// MinimockListDone returns true if the count of the List invocations corresponds
// the number of defined expectations
func (m *CurrencyPairMock) MinimockListDone() bool {
	for _, e := range m.ListMock.expectations {
		if mm_atomic.LoadUint64(&e.Counter) < 1 {
			return false
		}
	}

	// if default expectation was set then invocations count should be greater than zero
	if m.ListMock.defaultExpectation != nil && mm_atomic.LoadUint64(&m.afterListCounter) < 1 {
		return false
	}
	// if func was set then invocations count should be greater than zero
	if m.funcList != nil && mm_atomic.LoadUint64(&m.afterListCounter) < 1 {
		return false
	}
	return true
}

// MinimockListInspect logs each unmet expectation
func (m *CurrencyPairMock) MinimockListInspect() {
	for _, e := range m.ListMock.expectations {
		if mm_atomic.LoadUint64(&e.Counter) < 1 {
			m.t.Errorf("Expected call to CurrencyPairMock.List with params: %#v", *e.params)
		}
	}

	// if default expectation was set then invocations count should be greater than zero
	if m.ListMock.defaultExpectation != nil && mm_atomic.LoadUint64(&m.afterListCounter) < 1 {
		if m.ListMock.defaultExpectation.params == nil {
			m.t.Error("Expected call to CurrencyPairMock.List")
		} else {
			m.t.Errorf("Expected call to CurrencyPairMock.List with params: %#v", *m.ListMock.defaultExpectation.params)
		}
	}
	// if func was set then invocations count should be greater than zero
	if m.funcList != nil && mm_atomic.LoadUint64(&m.afterListCounter) < 1 {
		m.t.Error("Expected call to CurrencyPairMock.List")
	}
}

type mCurrencyPairMockUpdatePair struct {
	mock               *CurrencyPairMock
	defaultExpectation *CurrencyPairMockUpdatePairExpectation
	expectations       []*CurrencyPairMockUpdatePairExpectation

	callArgs []*CurrencyPairMockUpdatePairParams
	mutex    sync.RWMutex
}

// CurrencyPairMockUpdatePairExpectation specifies expectation struct of the CurrencyPair.UpdatePair
type CurrencyPairMockUpdatePairExpectation struct {
	mock    *CurrencyPairMock
	params  *CurrencyPairMockUpdatePairParams
	results *CurrencyPairMockUpdatePairResults
	Counter uint64
}

// CurrencyPairMockUpdatePairParams contains parameters of the CurrencyPair.UpdatePair
type CurrencyPairMockUpdatePairParams struct {
	ctx  context.Context
	pair *models.CurrencyPair
}

// CurrencyPairMockUpdatePairResults contains results of the CurrencyPair.UpdatePair
type CurrencyPairMockUpdatePairResults struct {
	err error
}

// Expect sets up expected params for CurrencyPair.UpdatePair
func (mmUpdatePair *mCurrencyPairMockUpdatePair) Expect(ctx context.Context, pair *models.CurrencyPair) *mCurrencyPairMockUpdatePair {
	if mmUpdatePair.mock.funcUpdatePair != nil {
		mmUpdatePair.mock.t.Fatalf("CurrencyPairMock.UpdatePair mock is already set by Set")
	}

	if mmUpdatePair.defaultExpectation == nil {
		mmUpdatePair.defaultExpectation = &CurrencyPairMockUpdatePairExpectation{}
	}

	mmUpdatePair.defaultExpectation.params = &CurrencyPairMockUpdatePairParams{ctx, pair}
	for _, e := range mmUpdatePair.expectations {
		if minimock.Equal(e.params, mmUpdatePair.defaultExpectation.params) {
			mmUpdatePair.mock.t.Fatalf("Expectation set by When has same params: %#v", *mmUpdatePair.defaultExpectation.params)
		}
	}

	return mmUpdatePair
}

// Inspect accepts an inspector function that has same arguments as the CurrencyPair.UpdatePair
func (mmUpdatePair *mCurrencyPairMockUpdatePair) Inspect(f func(ctx context.Context, pair *models.CurrencyPair)) *mCurrencyPairMockUpdatePair {
	if mmUpdatePair.mock.inspectFuncUpdatePair != nil {
		mmUpdatePair.mock.t.Fatalf("Inspect function is already set for CurrencyPairMock.UpdatePair")
	}

	mmUpdatePair.mock.inspectFuncUpdatePair = f

	return mmUpdatePair
}

// Return sets up results that will be returned by CurrencyPair.UpdatePair
func (mmUpdatePair *mCurrencyPairMockUpdatePair) Return(err error) *CurrencyPairMock {
	if mmUpdatePair.mock.funcUpdatePair != nil {
		mmUpdatePair.mock.t.Fatalf("CurrencyPairMock.UpdatePair mock is already set by Set")
	}

	if mmUpdatePair.defaultExpectation == nil {
		mmUpdatePair.defaultExpectation = &CurrencyPairMockUpdatePairExpectation{mock: mmUpdatePair.mock}
	}
	mmUpdatePair.defaultExpectation.results = &CurrencyPairMockUpdatePairResults{err}
	return mmUpdatePair.mock
}

// Set uses given function f to mock the CurrencyPair.UpdatePair method
func (mmUpdatePair *mCurrencyPairMockUpdatePair) Set(f func(ctx context.Context, pair *models.CurrencyPair) (err error)) *CurrencyPairMock {
	if mmUpdatePair.defaultExpectation != nil {
		mmUpdatePair.mock.t.Fatalf("Default expectation is already set for the CurrencyPair.UpdatePair method")
	}

	if len(mmUpdatePair.expectations) > 0 {
		mmUpdatePair.mock.t.Fatalf("Some expectations are already set for the CurrencyPair.UpdatePair method")
	}

	mmUpdatePair.mock.funcUpdatePair = f
	return mmUpdatePair.mock
}

// When sets expectation for the CurrencyPair.UpdatePair which will trigger the result defined by the following
// Then helper
func (mmUpdatePair *mCurrencyPairMockUpdatePair) When(ctx context.Context, pair *models.CurrencyPair) *CurrencyPairMockUpdatePairExpectation {
	if mmUpdatePair.mock.funcUpdatePair != nil {
		mmUpdatePair.mock.t.Fatalf("CurrencyPairMock.UpdatePair mock is already set by Set")
	}

	expectation := &CurrencyPairMockUpdatePairExpectation{
		mock:   mmUpdatePair.mock,
		params: &CurrencyPairMockUpdatePairParams{ctx, pair},
	}
	mmUpdatePair.expectations = append(mmUpdatePair.expectations, expectation)
	return expectation
}

// Then sets up CurrencyPair.UpdatePair return parameters for the expectation previously defined by the When method
func (e *CurrencyPairMockUpdatePairExpectation) Then(err error) *CurrencyPairMock {
	e.results = &CurrencyPairMockUpdatePairResults{err}
	return e.mock
}

// UpdatePair implements repository.CurrencyPair
func (mmUpdatePair *CurrencyPairMock) UpdatePair(ctx context.Context, pair *models.CurrencyPair) (err error) {
	mm_atomic.AddUint64(&mmUpdatePair.beforeUpdatePairCounter, 1)
	defer mm_atomic.AddUint64(&mmUpdatePair.afterUpdatePairCounter, 1)

	if mmUpdatePair.inspectFuncUpdatePair != nil {
		mmUpdatePair.inspectFuncUpdatePair(ctx, pair)
	}

	mm_params := &CurrencyPairMockUpdatePairParams{ctx, pair}

	// Record call args
	mmUpdatePair.UpdatePairMock.mutex.Lock()
	mmUpdatePair.UpdatePairMock.callArgs = append(mmUpdatePair.UpdatePairMock.callArgs, mm_params)
	mmUpdatePair.UpdatePairMock.mutex.Unlock()

	for _, e := range mmUpdatePair.UpdatePairMock.expectations {
		if minimock.Equal(e.params, mm_params) {
			mm_atomic.AddUint64(&e.Counter, 1)
			return e.results.err
		}
	}

	if mmUpdatePair.UpdatePairMock.defaultExpectation != nil {
		mm_atomic.AddUint64(&mmUpdatePair.UpdatePairMock.defaultExpectation.Counter, 1)
		mm_want := mmUpdatePair.UpdatePairMock.defaultExpectation.params
		mm_got := CurrencyPairMockUpdatePairParams{ctx, pair}
		if mm_want != nil && !minimock.Equal(*mm_want, mm_got) {
			mmUpdatePair.t.Errorf("CurrencyPairMock.UpdatePair got unexpected parameters, want: %#v, got: %#v%s\n", *mm_want, mm_got, minimock.Diff(*mm_want, mm_got))
		}

		mm_results := mmUpdatePair.UpdatePairMock.defaultExpectation.results
		if mm_results == nil {
			mmUpdatePair.t.Fatal("No results are set for the CurrencyPairMock.UpdatePair")
		}
		return (*mm_results).err
	}
	if mmUpdatePair.funcUpdatePair != nil {
		return mmUpdatePair.funcUpdatePair(ctx, pair)
	}
	mmUpdatePair.t.Fatalf("Unexpected call to CurrencyPairMock.UpdatePair. %v %v", ctx, pair)
	return
}

// UpdatePairAfterCounter returns a count of finished CurrencyPairMock.UpdatePair invocations
func (mmUpdatePair *CurrencyPairMock) UpdatePairAfterCounter() uint64 {
	return mm_atomic.LoadUint64(&mmUpdatePair.afterUpdatePairCounter)
}

// UpdatePairBeforeCounter returns a count of CurrencyPairMock.UpdatePair invocations
func (mmUpdatePair *CurrencyPairMock) UpdatePairBeforeCounter() uint64 {
	return mm_atomic.LoadUint64(&mmUpdatePair.beforeUpdatePairCounter)
}

// Calls returns a list of arguments used in each call to CurrencyPairMock.UpdatePair.
// The list is in the same order as the calls were made (i.e. recent calls have a higher index)
func (mmUpdatePair *mCurrencyPairMockUpdatePair) Calls() []*CurrencyPairMockUpdatePairParams {
	mmUpdatePair.mutex.RLock()

	argCopy := make([]*CurrencyPairMockUpdatePairParams, len(mmUpdatePair.callArgs))
	copy(argCopy, mmUpdatePair.callArgs)

	mmUpdatePair.mutex.RUnlock()

	return argCopy
}

// MinimockUpdatePairDone returns true if the count of the UpdatePair invocations corresponds
// the number of defined expectations
func (m *CurrencyPairMock) MinimockUpdatePairDone() bool {
	for _, e := range m.UpdatePairMock.expectations {
		if mm_atomic.LoadUint64(&e.Counter) < 1 {
			return false
		}
	}

	// if default expectation was set then invocations count should be greater than zero
	if m.UpdatePairMock.defaultExpectation != nil && mm_atomic.LoadUint64(&m.afterUpdatePairCounter) < 1 {
		return false
	}
	// if func was set then invocations count should be greater than zero
	if m.funcUpdatePair != nil && mm_atomic.LoadUint64(&m.afterUpdatePairCounter) < 1 {
		return false
	}
	return true
}

// MinimockUpdatePairInspect logs each unmet expectation
func (m *CurrencyPairMock) MinimockUpdatePairInspect() {
	for _, e := range m.UpdatePairMock.expectations {
		if mm_atomic.LoadUint64(&e.Counter) < 1 {
			m.t.Errorf("Expected call to CurrencyPairMock.UpdatePair with params: %#v", *e.params)
		}
	}

	// if default expectation was set then invocations count should be greater than zero
	if m.UpdatePairMock.defaultExpectation != nil && mm_atomic.LoadUint64(&m.afterUpdatePairCounter) < 1 {
		if m.UpdatePairMock.defaultExpectation.params == nil {
			m.t.Error("Expected call to CurrencyPairMock.UpdatePair")
		} else {
			m.t.Errorf("Expected call to CurrencyPairMock.UpdatePair with params: %#v", *m.UpdatePairMock.defaultExpectation.params)
		}
	}
	// if func was set then invocations count should be greater than zero
	if m.funcUpdatePair != nil && mm_atomic.LoadUint64(&m.afterUpdatePairCounter) < 1 {
		m.t.Error("Expected call to CurrencyPairMock.UpdatePair")
	}
}

// MinimockFinish checks that all mocked methods have been called the expected number of times
func (m *CurrencyPairMock) MinimockFinish() {
	if !m.minimockDone() {
		m.MinimockCreateInspect()

		m.MinimockGetInspect()

		m.MinimockListInspect()

		m.MinimockUpdatePairInspect()
		m.t.FailNow()
	}
}

// MinimockWait waits for all mocked methods to be called the expected number of times
func (m *CurrencyPairMock) MinimockWait(timeout mm_time.Duration) {
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

func (m *CurrencyPairMock) minimockDone() bool {
	done := true
	return done &&
		m.MinimockCreateDone() &&
		m.MinimockGetDone() &&
		m.MinimockListDone() &&
		m.MinimockUpdatePairDone()
}