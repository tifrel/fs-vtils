package fsv_test

import (
	"fsv"
	"os"
)

// case01: p is file => no error
// case02: p is symlink => no error
// case03: p doesn't exist => error
// case04: p is dir & no r flag => error
// case05: p is dir & r flag => no error
//
// case06: target exists & no f flag => error
// case07: target exists & f flag => no error
// case08: targets parent doesn't exist & no p flag => error
// case09: targets parent doesn't exist & p flag => no error

var mvTestLoc = testDir.AppendStr("Mv")
var mvTests = []testStruct{
	testFsvErr{

		_name: "case01",
		modify: func() error {
			path := mvTestLoc.AppendStr("case01/src/file")
			return path.Mv(mvTestLoc.AppendStr("case01/target/moved"))
		},
		_expect: nil,
	}, testFsvErr{

		_name: "case02",
		modify: func() error {
			path := mvTestLoc.AppendStr("case02/src/symlink")
			return path.Mv(mvTestLoc.AppendStr("case02/target/moved"))
		},
		_expect: nil,
	}, testOsErr{

		_name: "case03",
		modify: func() error {
			path := mvTestLoc.AppendStr("case03/src/none")
			return path.Mv(mvTestLoc.AppendStr("case03/target/moved"))
		},
		_expect:  os.IsNotExist,
		_example: os.ErrNotExist,
	}, testFsvErr{

		_name: "case04",
		modify: func() error {
			path := mvTestLoc.AppendStr("case04/src/dir")
			return path.Mv(mvTestLoc.AppendStr("case04/target/moved"))
		},
		_expect: fsv.MISSING_REC_FLAG,
	}, testFsvErr{

		_name: "case05",
		modify: func() error {
			path := mvTestLoc.AppendStr("case05/src/dir")
			return path.Mv(mvTestLoc.AppendStr("case05/target/moved"), 'r')
		},
		_expect: nil,
	}, testFsvErr{

		_name: "case06",
		modify: func() error {
			path := mvTestLoc.AppendStr("case06/src/file")
			return path.Mv(mvTestLoc.AppendStr("case06/target/file"))
		},
		_expect: fsv.OCCUPIED_PATH,
	}, testFsvErr{

		_name: "case07",
		modify: func() error {
			path := mvTestLoc.AppendStr("case07/src/file")
			return path.Mv(mvTestLoc.AppendStr("case07/target/file"), 'f')
		},
		_expect: nil,
	}, testFsvErr{

		_name: "case08",
		modify: func() error {
			path := mvTestLoc.AppendStr("case08/src/file")
			return path.Mv(mvTestLoc.AppendStr("case08/target/nested/moved"))
		},
		_expect: fsv.MISSING_TARGETDIR,
	}, testFsvErr{

		_name: "case09",
		modify: func() error {
			path := mvTestLoc.AppendStr("case09/src/file")
			return path.Mv(mvTestLoc.AppendStr("case09/target/nested/moved"), 'p')
		},
		_expect: nil,
	},
}
