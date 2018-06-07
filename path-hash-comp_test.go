package fsv_test

var hashTests = []Tester{
	boolTest{
		desc: "Hash comparison, different contents",
		testFn: func() (bool, error) {
			return miscA.SameHashAs(miscC)
		},
		expect: false,
	},
	boolTest{
		desc: "Byte-by-byte comparison, equal contents",
		testFn: func() (bool, error) {
			return miscA.SameContentsAs(miscB)
		},
		expect: true,
	},
	boolTest{
		desc: "Byte-by-byte comparison, different contents",
		testFn: func() (bool, error) {
			return miscA.SameContentsAs(miscC)
		},
		expect: false,
	},
	boolTest{
		desc: "Cashed hash comparison, equal contents",
		testFn: func() (bool, error) {
			hashA, err := miscA.Hash()
			if err != nil {
				return false, err
			}
			return miscB.HasHash(hashA)
		},
		expect: true,
	},
	boolTest{
		desc: "Cashed hash comparison, different contents",
		testFn: func() (bool, error) {
			hashA, err := miscA.Hash()
			if err != nil {
				return true, err
			}
			return miscC.HasHash(hashA)
		},
		expect: false,
	},
	boolTest{
		desc: "FileInfo comparison, same file",
		testFn: func() (bool, error) {
			return miscA.SameInfoAs(miscA)
		},
		expect: true,
	},
	boolTest{
		desc: "FileInfo comparison, different file",
		testFn: func() (bool, error) {
			return miscA.SameInfoAs(miscB)
		},
		expect: false,
	},
	boolTest{
		desc: "FileInfo comparison, different file",
		testFn: func() (bool, error) {
			return miscA.SameInfoAs(miscC)
		},
		expect: false,
	},
	boolTest{
		desc: "FileInfo comparison, hardlinked files",
		testFn: func() (bool, error) {
			return miscA.SameInfoAs(miscD)
		},
		expect: true,
	},
}
