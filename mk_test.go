package fsv_test

import (
	"fsv"
)

// case01: target exists and is dir => bombä
// case02: target doesnt exist & parent is dir => make dir
//
// case03: target is file & no f flag => error
// case04: target is file & f flag => remove file; make dir
//
// case05: parent doesn't exist & no p flag => error
// case06: parent is file & no p flag => error
// case07: parent is file & p flag & no f flag => error
// case08: parent is file & p flag & f flag
//     => remove parent; make parent; make dir
//
// case09: parent doesn't exist, superparent is file with f and p flags

var mkTestLocdir = testDir.AppendStr("Mkdir")
var mkDirTests = []testStruct{
	testFsvErr{

		_name: "case01",
		modify: func() error {
			path := mkTestLocdir.AppendStr("case01/parent/dir")
			return path.MkDir()
		},
		_expect: nil,
	}, testFsvErr{

		_name: "case02",
		modify: func() error {
			path := mkTestLocdir.AppendStr("case02/parent/new")
			return path.MkDir()
		},
		_expect: nil,
	}, testFsvErr{

		_name: "case03",
		modify: func() error {
			path := mkTestLocdir.AppendStr("case03/parent/file")
			return path.MkDir()
		},
		_expect: fsv.OCCUPIED_PATH,
	}, testFsvErr{

		_name: "case04",
		modify: func() error {
			path := mkTestLocdir.AppendStr("case04/parent/file")
			return path.MkDir('f')
		},
		_expect: nil,
	}, testFsvErr{

		_name: "case05",
		modify: func() error {
			path := mkTestLocdir.AppendStr("case05/parent2/new")
			return path.MkDir()
		},
		_expect: fsv.MISSING_TARGETDIR,
	}, testFsvErr{

		_name: "case06",
		modify: func() error {
			path := mkTestLocdir.AppendStr("case06/parentfile/new")
			return path.MkDir()
		},
		_expect: fsv.MISSING_TARGETDIR,
	}, testFsvErr{

		_name: "case07",
		modify: func() error {
			path := mkTestLocdir.AppendStr("case07/parentfile/new")
			return path.MkDir('p')
		},
		_expect: fsv.OCCUPIED_PATH,
	}, testFsvErr{

		_name: "case08",
		modify: func() error {
			path := mkTestLocdir.AppendStr("case08/parentfile/new")
			return path.MkDir('p', 'f')
		},
		_expect: nil,
	}, testFsvErr{

		_name: "case08",
		modify: func() error {
			path := mkTestLocdir.AppendStr("case09/parentfile/deeply/nested/new/dir")
			return path.MkDir('p', 'f')
		},
		_expect: nil,
	},
}

// case01: target doesnt exist & parent is dir => make file
//
// case02: target is file & no f flag => error
// case03: target is file & f flag => remove file; make file
// case04: target is dir & no f flag => error
// case05: target is dir & f flag => remove file; make file
//
// case06: parent doesn't exist & no p flag => error
// case07: parent doesn't exist & p flag => make parent; make file
// case08: parent is file & no p flag => error
// case09: parent is file & p flag & no f flag => error
// case10: parent is file & f flag & no p flag => error
// case11: parent is file & p flag & f flag
//     => remove parent; make parent; make file
//
// case09: parent doesn't exist, superparent is file with f and p flags

var mkTestLocFile = testDir.AppendStr("MkFile")
var mkFileTests = []testStruct{
	testFsvErr{

		_name: "case01",
		modify: func() error {
			return mkTestLocFile.AppendStr("case01/parent/new").MkFile(0644)
		},
		_expect: nil,
	}, testFsvErr{

		_name: "case02",
		modify: func() error {
			return mkTestLocFile.AppendStr("case02/parent/file").MkFile(0644)
		},
		_expect: fsv.OCCUPIED_PATH,
	}, testFsvErr{

		_name: "case03",
		modify: func() error {
			return mkTestLocFile.AppendStr("case03/parent/file").MkFile(0644, 'f')
		},
		_expect: nil,
	}, testFsvErr{

		_name: "case04",
		modify: func() error {
			return mkTestLocFile.AppendStr("case04/parent/dir").MkFile(0644)
		},
		_expect: fsv.OCCUPIED_PATH,
	}, testFsvErr{

		_name: "case05",
		modify: func() error {
			return mkTestLocFile.AppendStr("case05/parent/dir").MkFile(0644, 'f')
		},
		_expect: nil,
	}, testFsvErr{

		_name: "case06",
		modify: func() error {
			return mkTestLocFile.AppendStr("case06/parent2/new").MkFile(0644)
		},
		_expect: fsv.MISSING_TARGETDIR,
	}, testFsvErr{

		_name: "case07",
		modify: func() error {
			return mkTestLocFile.AppendStr("case07/parent2/new").MkFile(0644, 'p')
		},
		_expect: nil,
	}, testFsvErr{

		_name: "case08",
		modify: func() error {
			return mkTestLocFile.AppendStr("case08/parentfile/new").MkFile(0644)
		},
		_expect: fsv.OCCUPIED_PATH,
	}, testFsvErr{

		_name: "case09",
		modify: func() error {
			return mkTestLocFile.AppendStr("case09/parentfile/new").MkFile(0644, 'p')
		},
		_expect: fsv.OCCUPIED_PATH,
	}, testFsvErr{

		_name: "case10",
		modify: func() error {
			return mkTestLocFile.AppendStr("case10/parentfile/new").MkFile(0644, 'f')
		},
		_expect: fsv.MISSING_TARGETDIR,
	}, testFsvErr{

		_name: "case11",
		modify: func() error {
			return mkTestLocFile.AppendStr("case11/parentfile/new").MkFile(0644, 'f', 'p')
		},
		_expect: nil,
	},
}
