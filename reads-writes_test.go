package fsv_test

import (
	"os"
)

// case01: writing, file is empty => no error
// case02: writing, file has contents => no error
// case03: writing, file doesn't exist => no error
//     TODO: to error, or not to error, that is the question...
// case04: reading, file is empty => no error
// case05: reading, file has contents => no error
// case06: reading, file doesn't exist => error

var rwTestLoc = testDir.AppendStr("rw")
var rwTests = []testStruct{
	testValFn{
		_name: "case01",
		validate: func() (bool, error) {
			err := rwTestLoc.AppendStr("case01/file").WriteString("written")
			return err == nil, err
		},
		_expect: nil,
	},
	testValFn{
		_name: "case02",
		validate: func() (bool, error) {
			err := rwTestLoc.AppendStr("case02/empty").WriteString("written")
			return err == nil, err
		},
		_expect: nil,
	},
	testValFn{
		_name: "case03",
		validate: func() (bool, error) {
			err := rwTestLoc.AppendStr("case03/new").WriteString("written")
			return true, err
		},
		_expect:  os.IsNotExist,
		_example: os.ErrNotExist,
	},
	testValFn{
		_name: "case04",
		validate: func() (bool, error) {
			contents, err := rwTestLoc.AppendStr("case04/file").ReadString()
			return contents == "file-contents\n", err
		},
		_expect: nil,
	},
	testValFn{
		_name: "case05",
		validate: func() (bool, error) {
			contents, err := rwTestLoc.AppendStr("case05/empty").ReadString()
			return contents == "", err
		},
		_expect: nil,
	},
	testValFn{
		_name: "case06",
		validate: func() (bool, error) {
			contents, err := rwTestLoc.AppendStr("case06/yadda").ReadString()
			return contents == "", err
		},
		_expect:  os.IsNotExist,
		_example: os.ErrNotExist,
	},
}
