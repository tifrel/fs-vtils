package fsv_test

import (
	"os"
)

// case01: writing, file is empty => no error
// case02: writing, file has contents => no error
// case03: writing, file doesn't exist => no error
// case04: reading, file is empty => no error
// case05: reading, file has contents => no error
// case06: reading, file doesn't exist => error
// case07: appending, file is empty => no error
// case08: appending, file has contents => no error
// case09: appending, file doesn't exist => error
// case10: writing multiple lines => no error
// case11: reading multiple lines => no error

var rwTestLoc = testDir.ExtendStr("rw")
var rwTests = []testStruct{
	testValFn{
		_name: "case01",
		validate: func() (bool, error) {
			err := rwTestLoc.ExtendStr("case01/file").WriteString("written")
			return err == nil, err
		},
		_expect: nil,
	},
	testValFn{
		_name: "case02",
		validate: func() (bool, error) {
			err := rwTestLoc.ExtendStr("case02/empty").WriteString("written")
			return err == nil, err
		},
		_expect: nil,
	},
	testValFn{
		_name: "case03",
		validate: func() (bool, error) {
			err := rwTestLoc.ExtendStr("case03/new").WriteString("written")
			return true, err
		},
		_expect:  os.IsNotExist,
		_example: os.ErrNotExist,
	},
	testValFn{
		_name: "case04",
		validate: func() (bool, error) {
			contents, err := rwTestLoc.ExtendStr("case04/file").ReadString()
			return contents == "a\nb\nc\n", err
		},
		_expect: nil,
	},
	testValFn{
		_name: "case05",
		validate: func() (bool, error) {
			contents, err := rwTestLoc.ExtendStr("case05/empty").ReadString()
			return contents == "", err
		},
		_expect: nil,
	},
	testValFn{
		_name: "case06",
		validate: func() (bool, error) {
			contents, err := rwTestLoc.ExtendStr("case06/yadda").ReadString()
			return contents == "", err
		},
		_expect:  os.IsNotExist,
		_example: os.ErrNotExist,
	},
	testValFn{
		_name: "case07",
		validate: func() (bool, error) {
			err := rwTestLoc.ExtendStr("case07/empty").AppendString("appended")
			return err == nil, err
		},
		_expect: nil,
	},
	testValFn{
		_name: "case08",
		validate: func() (bool, error) {
			err := rwTestLoc.ExtendStr("case08/file").AppendString("appended")
			return err == nil, err
		},
		_expect: nil,
	},
	testValFn{
		_name: "case09",
		validate: func() (bool, error) {
			err := rwTestLoc.ExtendStr("case09/yadda").AppendString("appended")
			return err != nil, err
		},
		_expect:  os.IsNotExist,
		_example: os.ErrNotExist,
	},
	testValFn{
		_name: "case10",
		validate: func() (bool, error) {
			lines := []string{"a", "b", "c"}
			err := rwTestLoc.ExtendStr("case10/empty").WriteLines(lines)
			return err == nil, err
		},
		_expect: nil,
	},
	testValFn{
		_name: "case11",
		validate: func() (bool, error) {
			contents, err := rwTestLoc.ExtendStr("case11/file").ReadLines()
			if err != nil {
				return false, err
			}
			ok1 := contents[0] == "a"
			ok2 := contents[1] == "b"
			ok3 := contents[2] == "c"
			return ok1 && ok2 && ok3, err
		},
		_expect: nil,
	},
}
