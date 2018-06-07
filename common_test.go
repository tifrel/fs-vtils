package fsv_test

import (
	"fsv"
	"testing"
)

// ------------------------------ test exports ------------------------------ //

/**** Bash-like functions ****/
func TestCp(t *testing.T) {
	runTests(cpTests, t)
}

func TestLn(t *testing.T) {
	runTests(lnTests, t)
}

func TestMk(t *testing.T) {
	runTests(append(mkDirTests, mkFileTests...), t)
}

func TestMv(t *testing.T) {
	runTests(mvTests, t)
}

func TestRm(t *testing.T) {
	runTests(rmTests, t)
}

/**** Reads & Writes ****/
func TestRW(t *testing.T) {
	runTests(rwTests, t)
}

/**** Information methods ****/

/**** Manipulation methods ****/
func TestManipulation(t *testing.T) {
	runValTests(manipulationTests, t)
}

/**** Hashing and comparison methods ****/
func TestHash(t *testing.T) {
	runValTests(hashTests, t)
}

// ------------------------- test runner functions -------------------------- //
// TODO: refactor, only need one function

const testDir fsv.Path = "/Users/Till/Code/go/src/fsv/testdata"

func runTests(tests []testStruct, t *testing.T) {

	for _, test := range tests {
		if ok, res := test.passes(); !ok {

			t.Error(
				"Test Failed (", test.name(), ")",
				"\nexpected: ", test.expect(),
				"\ngot:      ", res,
				"\n",
			)

		}
	}

}

func runValTests(tests []Tester, t *testing.T) {
	for _, test := range tests {
		if ok, res := test.Passes(); !ok {
			t.Error(
				"Test Failed (", test.Desc(), ")",
				"\nexpected: ", test.Expect(),
				"\ngot:      ", res,
			)
		}
	}
}

// ------------------------- error testing helpers -------------------------- //
// TODO: refactor so that only one of those exists

type testStruct interface {
	name() string
	passes() (bool, error)
	expect() error
}

type testFsvErr struct {
	_name   string
	modify  func() error
	_expect error
}

func (t testFsvErr) name() string {
	return t._name
}

func (t testFsvErr) passes() (bool, error) {
	result := t.modify()
	switch {
	case result == nil && t._expect == nil:
		return true, result
	case result != nil && t._expect == nil:
		return false, result
	case result == nil && t._expect != nil:
		return false, result
	}

	fsve, ok := t.expect().(fsv.Error)
	if !ok {
		return false, result
	}

	return fsve.IsTypeOf(result), result
}

func (t testFsvErr) expect() error {
	return t._expect
}

type testOsErr struct {
	_name    string
	modify   func() error
	_expect  func(error) bool
	_example error
}

func (t testOsErr) name() string {
	return t._name
}

func (t testOsErr) passes() (bool, error) {
	res := t.modify()

	switch {
	case res == nil && t._expect == nil:
		return true, res
	case res != nil && t._expect == nil:
		return false, res
	case res == nil && t._expect != nil:
		return false, res
	}

	return t._expect(res), res
}

func (t testOsErr) expect() error {
	return t._example
}

type testValFn struct {
	_name    string
	validate func() (bool, error)
	_expect  func(error) bool
	_example error
}

func (t testValFn) name() string {
	return t._name
}

func (t testValFn) passes() (bool, error) {
	ok, err := t.validate()
	if !ok {
		return ok, err
	}

	switch {
	case err == nil && t._expect == nil:
		return true, err
	case err != nil && t._expect == nil:
		return false, err
	case err == nil && t._expect != nil:
		return false, err
	}

	return t._expect(err), err
}

func (t testValFn) expect() error {
	return t._example
}

// ----------------------------- value testers ------------------------------ //
type Tester interface {
	Desc() string
	Passes() (bool, string)
	Expect() string
}

type boolTest struct {
	desc   string
	testFn func() (bool, error)
	expect bool
}

func (t boolTest) Passes() (bool, string) {
	res, err := t.testFn()
	if err != nil {
		return false, "ERROR: " + err.Error()
	}
	return t.expect == res, boolToString(res)
}

func (t boolTest) Expect() string {
	return boolToString(t.expect)
}

func (t boolTest) Desc() string {
	return t.desc
}

type intTest struct {
	desc   string
	testFn func() (int, error)
	expect int
}

func (t intTest) Passes() (bool, string) {
	res, err := t.testFn()
	if err != nil {
		return false, "ERROR: " + err.Error()
	}
	return t.expect == res, string(res)
}

func (t intTest) Expect() string {
	return string(t.expect)
}

func (t intTest) Desc() string {
	return t.desc
}

type stringTest struct {
	desc   string
	testFn func() (string, error)
	expect string
}

func (t stringTest) Passes() (bool, string) {
	res, err := t.testFn()
	if err != nil {
		return false, "ERROR: " + err.Error()
	}
	return t.expect == res, string(res)
}

func (t stringTest) Expect() string {
	return string(t.expect)
}

func (t stringTest) Desc() string {
	return t.desc
}

// -------------------- helpers I didn't expect to need --------------------- //
func boolToString(b bool) string {
	if b {
		return "true"
	}
	return "false"
}
