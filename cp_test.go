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

var cpTestLoc = testDir.ExtendStr("Cp")
var cpTests = []testStruct{
	testFsvErr{

		_name: "case01",
		modify: func() error {
			path := cpTestLoc.ExtendStr("case01/file")
			return path.Cp(cpTestLoc.ExtendStr("case01/copied"))
		},
		_expect: nil,
	}, testFsvErr{

		_name: "case02",
		modify: func() error {
			path := cpTestLoc.ExtendStr("case02/filelink")
			return path.Cp(cpTestLoc.ExtendStr("case02/copied"))
		},
		_expect: nil,
	}, testFsvErr{

		_name: "case03",
		modify: func() error {
			path := cpTestLoc.ExtendStr("case03/filelink")
			return path.Cp(cpTestLoc.ExtendStr("case03/copied"), 'd')
		},
		_expect: nil,
	}, testFsvErr{

		_name: "case04",
		modify: func() error {
			path := cpTestLoc.ExtendStr("case04/dir")
			return path.Cp(cpTestLoc.ExtendStr("case04/copied"))
		},
		_expect: fsv.MISSING_REC_FLAG,
	}, testFsvErr{

		_name: "case05",
		modify: func() error {
			path := cpTestLoc.ExtendStr("case05/dir")
			return path.Cp(cpTestLoc.ExtendStr("case05/copied"), 'r')
		},
		_expect: nil,
	}, testFsvErr{

		_name: "case06",
		modify: func() error {
			path := cpTestLoc.ExtendStr("case06/dirlink")
			return path.Cp(cpTestLoc.ExtendStr("case06/copied"), 'd')
		},
		_expect: fsv.MISSING_REC_FLAG,
	}, testFsvErr{

		_name: "case07",
		modify: func() error {
			path := cpTestLoc.ExtendStr("case07/dirlink")
			return path.Cp(cpTestLoc.ExtendStr("case07/copied"), 'd', 'r')
		},
		_expect: nil,
	}, testFsvErr{

		_name: "case08",
		modify: func() error {
			path := cpTestLoc.ExtendStr("case08/dir2")
			return path.Cp(cpTestLoc.ExtendStr("case08/copied"), 'r')
		},
		_expect: nil,
	}, testFsvErr{

		_name: "case09",
		modify: func() error {
			path := cpTestLoc.ExtendStr("case09/dir2")
			return path.Cp(cpTestLoc.ExtendStr("case09/copied"), 'r', 'd')
		},
		_expect: nil,
	}, testFsvErr{

		_name: "case10",
		modify: func() error {
			path := cpTestLoc.ExtendStr("case10/file")
			return path.Cp(cpTestLoc.ExtendStr("case10/dir"))
		},
		_expect: fsv.OCCUPIED_PATH,
	}, testFsvErr{

		_name: "case11",
		modify: func() error {
			path := cpTestLoc.ExtendStr("case11/file")
			return path.Cp(cpTestLoc.ExtendStr("case11/dir"), 'f')
		},
		_expect: nil,
	}, testFsvErr{

		_name: "case12",
		modify: func() error {
			path := cpTestLoc.ExtendStr("case12/file")
			return path.Cp(cpTestLoc.ExtendStr("case12/newdir/copied"))
		},
		_expect: fsv.MISSING_TARGETDIR,
	}, testFsvErr{

		_name: "case13",
		modify: func() error {
			path := cpTestLoc.ExtendStr("case13/file")
			return path.Cp(cpTestLoc.ExtendStr("case13/newdir/copied"), 'p')
		},
		_expect: nil,
	},
}
