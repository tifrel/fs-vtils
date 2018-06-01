package fsv_test

import (
	"fsv"
)

// case01: p is file => no error
// case02: p is symlink & no d flag => no error
// case03: p is symlink & d flag => no error
// case04: p is dir & no r flag => error
// case05: p is dir & r flag => no error
// case06: p is symlink to dir & d flag & no r flag => error
// case07: p is symlink to dir & d flag & r flag => no error
// case08: p is dir with symlinks & no d flag => no error
// case09: p is dir with symlinks & d flag => no error
//
// case10: target exists & no f flag => error
// case11: target exists & f flag => no error
// case12: targetdir doesn't exist & no p flag => error
// case13: targetdir doesn't exist & p flag => no error

var testLocCp = testDir.AppendStr("Cp")
var cpTests = []testStruct{
	testFsvErr{

		_name: "case01",
		modify: func() error {
			path := testLocCp.AppendStr("case01/file")
			return path.Cp(testLocCp.AppendStr("case01/copied"))
		},
		_expect: nil,
	}, testFsvErr{

		_name: "case02",
		modify: func() error {
			path := testLocCp.AppendStr("case02/filelink")
			return path.Cp(testLocCp.AppendStr("case02/copied"))
		},
		_expect: nil,
	}, testFsvErr{

		_name: "case03",
		modify: func() error {
			path := testLocCp.AppendStr("case03/filelink")
			return path.Cp(testLocCp.AppendStr("case03/copied"), 'd')
		},
		_expect: nil,
	}, testFsvErr{

		_name: "case04",
		modify: func() error {
			path := testLocCp.AppendStr("case04/dir")
			return path.Cp(testLocCp.AppendStr("case04/copied"))
		},
		_expect: fsv.MISSING_REC_FLAG,
	}, testFsvErr{

		_name: "case05",
		modify: func() error {
			path := testLocCp.AppendStr("case05/dir")
			return path.Cp(testLocCp.AppendStr("case05/copied"), 'r')
		},
		_expect: nil,
	}, testFsvErr{

		_name: "case06",
		modify: func() error {
			path := testLocCp.AppendStr("case06/dirlink")
			return path.Cp(testLocCp.AppendStr("case06/copied"), 'd')
		},
		_expect: nil,
		// _expect: fsv.MISSING_REC_FLAG,
	}, testFsvErr{

		_name: "case07",
		modify: func() error {
			path := testLocCp.AppendStr("case07/dirlink")
			return path.Cp(testLocCp.AppendStr("case07/copied"), 'd', 'r')
		},
		_expect: nil,
	}, testFsvErr{

		_name: "case08",
		modify: func() error {
			path := testLocCp.AppendStr("case08/dir2")
			return path.Cp(testLocCp.AppendStr("case08/copied"), 'r')
		},
		_expect: nil,
	}, testFsvErr{

		_name: "case09",
		modify: func() error {
			path := testLocCp.AppendStr("case09/dir2")
			return path.Cp(testLocCp.AppendStr("case09/copied"), 'r', 'd')
		},
		_expect: nil,
	}, testFsvErr{

		_name: "case10",
		modify: func() error {
			path := testLocCp.AppendStr("case10/file")
			return path.Cp(testLocCp.AppendStr("case10/dir"))
		},
		_expect: fsv.OCCUPIED_PATH,
	}, testFsvErr{

		_name: "case11",
		modify: func() error {
			path := testLocCp.AppendStr("case11/file")
			return path.Cp(testLocCp.AppendStr("case11/dir"), 'f')
		},
		_expect: nil,
	}, testFsvErr{

		_name: "case12",
		modify: func() error {
			path := testLocCp.AppendStr("case12/file")
			return path.Cp(testLocCp.AppendStr("case12/newdir/copied"))
		},
		_expect: fsv.MISSING_TARGETDIR,
	}, testFsvErr{

		_name: "case13",
		modify: func() error {
			path := testLocCp.AppendStr("case13/file")
			return path.Cp(testLocCp.AppendStr("case13/newdir/copied"), 'p')
		},
		_expect: nil,
	},
}
