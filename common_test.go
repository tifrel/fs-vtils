package fsv_test

import (
	"fsv"
	"testing"
)

// TODO: testing of reading/writing/information methods

// ------------------------------ test exports ------------------------------ //

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

// ------------------------------ test helpers ------------------------------ //

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

type testStruct interface {
	name() string
	passes() (bool, error)
	expect() error
}

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
