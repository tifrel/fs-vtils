package fsv_test

import (
	"fsv"
)

// case01: p is file => no error
// case02: p is symlink => no error
// case03: p is dir => no error
// case04: p is file & h flag => no error
// case05: p is file & d flag => no error
// case06: p is symlink & d flag => no error
// case07: p is symlink & d flag & h flag => no error
//
// case08: target exists & no f flag => error
// case09: target exists & f flag => no error
// case10: targets parent doesn't exist & no p flag => error
// case11x: targets parent doesn't exist & p flag => no error

var lnTestLoc = testDir.AppendStr("Ln")
var lnTests = []testStruct{
	testFsvErr{
		_name: "case01",
		modify: func() error {
			path := lnTestLoc.AppendStr("case01/file")
			return path.Ln(lnTestLoc.AppendStr("case01/newlink"))
		},
		_expect: nil,
	},

	testFsvErr{
		_name: "case02",
		modify: func() error {
			path := lnTestLoc.AppendStr("case02/oldlink")
			return path.Ln(lnTestLoc.AppendStr("case02/newlink"))
		},
		_expect: nil,
	},

	testFsvErr{
		_name: "case03",
		modify: func() error {
			path := lnTestLoc.AppendStr("case03/dir")
			return path.Ln(lnTestLoc.AppendStr("case03/newlink"))
		},
		_expect: nil,
	},

	testFsvErr{
		_name: "case04",
		modify: func() error {
			path := lnTestLoc.AppendStr("case04/file")
			return path.Ln(lnTestLoc.AppendStr("case04/newlink"), 'h')
		},
		_expect: nil,
	},

	testFsvErr{
		_name: "case05",
		modify: func() error {
			path := lnTestLoc.AppendStr("case05/file")
			return path.Ln(lnTestLoc.AppendStr("case05/newlink"), 'd')
		},
		_expect: nil,
	},

	testFsvErr{
		_name: "case06",
		modify: func() error {
			path := lnTestLoc.AppendStr("case06/oldlink")
			return path.Ln(lnTestLoc.AppendStr("case06/newlink"), 'd')
		},
		_expect: nil,
	},

	testFsvErr{
		_name: "case07",
		modify: func() error {
			path := lnTestLoc.AppendStr("case07/oldlink")
			return path.Ln(lnTestLoc.AppendStr("case07/newlink"), 'd', 'h')
		},
		_expect: nil,
	},

	testFsvErr{
		_name: "case08",
		modify: func() error {
			path := lnTestLoc.AppendStr("case08/file")
			return path.Ln(lnTestLoc.AppendStr("case08/oldlink"))
		},
		_expect: fsv.OCCUPIED_PATH,
	},

	testFsvErr{
		_name: "case09",
		modify: func() error {
			path := lnTestLoc.AppendStr("case09/file2")
			_ = path.MkFile(0644)
			return path.Ln(lnTestLoc.AppendStr("case09/oldlink"), 'f')
		},
		_expect: nil,
	},

	testFsvErr{
		_name: "case10",
		modify: func() error {
			path := lnTestLoc.AppendStr("case10/file")
			return path.Ln(lnTestLoc.AppendStr("case10/newdir/newlink"))
		},
		_expect: fsv.MISSING_TARGETDIR,
	},

	testFsvErr{
		_name: "case11",
		modify: func() error {
			path := lnTestLoc.AppendStr("case11/file")
			return path.Ln(lnTestLoc.AppendStr("case11/newdir/newlink"), 'p')
		},
		_expect: nil,
	},
}
