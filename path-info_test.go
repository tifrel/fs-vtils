package fsv_test

import "fsv"

var infoTests = []Tester{
	boolTest{
		desc: "Path.Exist, file exists",
		testFn: func() (bool, error) {
			return miscA.Exists(), nil
		},
		expect: true,
	},
	boolTest{
		desc: "Path.Exist, file doesn't exist",
		testFn: func() (bool, error) {
			return fsv.Path("/inexistent/path/to/yadda").Exists(), nil
		},
		expect: false,
	},
	boolTest{
		desc: "Path.IsFile, target is file",
		testFn: func() (bool, error) {
			return miscA.IsFile()
		},
		expect: true,
	},
	boolTest{
		desc: "Path.IsFile, target is dir",
		testFn: func() (bool, error) {
			return miscDir.IsFile()
		},
		expect: false,
	},
	boolTest{
		desc: "Path.IsDir, target is file",
		testFn: func() (bool, error) {
			return miscA.IsDir()
		},
		expect: false,
	},
	boolTest{
		desc: "Path.IsDir, target is dir",
		testFn: func() (bool, error) {
			return miscDir.IsDir()
		},
		expect: true,
	},
	boolTest{
		desc: "Path.IsSymlink, target is file",
		testFn: func() (bool, error) {
			return miscA.IsSymlink()
		},
		expect: false,
	},
	boolTest{
		desc: "Path.IsSymlink, target is dir",
		testFn: func() (bool, error) {
			return miscDir.IsSymlink()
		},
		expect: false,
	},
	boolTest{
		desc: "Path.IsSymlink, target is hardlink",
		testFn: func() (bool, error) {
			return miscD.IsSymlink()
		},
		expect: false,
	},
	boolTest{
		desc: "Path.IsSymlink, target is symlink",
		testFn: func() (bool, error) {
			return miscE.IsSymlink()
		},
		expect: true,
	},
	stringTest{
		desc: "Path.Follow, target of link is file",
		testFn: func() (string, error) {
			p, err := miscE.Follow()
			return string(p), err
		},
		expect: string(miscA),
	},
	stringTest{
		desc: "Path.Follow, target of link is link",
		testFn: func() (string, error) {
			p, err := miscF.Follow()
			return string(p), err
		},
		expect: string(miscE),
	},
	stringTest{
		desc: "Path.Target",
		testFn: func() (string, error) {
			p, err := miscF.Target()
			return string(p), err
		},
		expect: string(miscA),
	},
	// TODO: Path.Ls
	boolTest{
		desc: "Path.IsOsFile, it is",
		testFn: func() (bool, error) {
			return fsv.Path("/yadda/.DS_Store").IsOsFile(), nil
		},
		expect: true,
	},
	boolTest{
		desc: "Path.IsOsFile, it isn't",
		testFn: func() (bool, error) {
			return fsv.Path("/yadda/yadda").IsOsFile(), nil
		},
		expect: false,
	},
	boolTest{
		desc: "Path.IsHidden, it is",
		testFn: func() (bool, error) {
			return fsv.Path("/yadda/.yadda").IsHidden(), nil
		},
		expect: true,
	},
	boolTest{
		desc: "Path.IsHidden, it isn't",
		testFn: func() (bool, error) {
			return fsv.Path("/yadda/yadda").IsHidden(), nil
		},
		expect: false,
	},
	// TODO: Mode and Owner stuff
	// TODO: Size, CountRunes, CountLines
}
