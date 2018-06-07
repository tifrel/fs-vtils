package fsv_test

import "fsv"

var manipulationTests = []Tester{
	stringTest{
		desc: "Path.Extend",
		testFn: func() (string, error) {
			return string(fsv.Path("a/b/c/d").Extend(fsv.Path("e/f"))), nil
		},
		expect: "a/b/c/d/e/f",
	},
	stringTest{
		desc: "Path.Base",
		testFn: func() (string, error) {
			return string(fsv.Path("a/b/c/d").Base()), nil
		},
		expect: "d",
	},
	stringTest{
		desc: "Path.Dir",
		testFn: func() (string, error) {
			return string(fsv.Path("a/b/c/d").Dir()), nil
		},
		expect: "a/b/c",
	},
	stringTest{
		desc: "Path.RelativeTo",
		testFn: func() (string, error) {
			rel, err := fsv.Path("a/b/c/d").RelativeTo(fsv.Path("a/b"))
			return string(rel), err
		},
		expect: "c/d",
	},
}
