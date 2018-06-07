package fsv

import (
	"os"
)

// PathList is a collection of paths, typically retrieved via Path.Ls().
type PathList []Path

// Filter removes all Paths from a PathList, that do not satisfy the given
// predicate.
func (ps PathList) Filter(pred func(Path) bool) PathList {
	out, i := make([]Path, len(ps)), 0

	for _, p := range ps {
		if pred(p) {
			out[i] = p
			i++
		}
	}

	return out[:i]
}

// Names returns only the basenames of the Paths in a PathList.
func (ps PathList) Names() []string {
	out := make([]string, len(ps))
	for i, p := range ps {
		out[i] = string(p.Base())
	}
	return out
}

// Infos returns a list of os.FileInfos from a PathList.
func (ps PathList) Infos() ([]os.FileInfo, error) {
	var (
		out = make([]os.FileInfo, len(ps))
		nfo os.FileInfo
		err error
	)

	for i, p := range ps {
		nfo, err = p.Info()
		if err != nil {
			return nil, err
		}

		out[i] = nfo
	}
	return out, nil
}

// Each applies a function to each Path in a PathList. Encountered errors
// accumulate and do not abort subsequent actions.
func (ps PathList) Each(fn func(Path) error) error {
	errs := ErrorList(make([]error, len(ps)))

	for i, p := range ps {
		err := fn(p)
		errs[i] = err
	}

	return errs.errorize()
}

// Dirs maps Dir over the PathList.
func (ps PathList) Dirs() []Path {
	ds := make([]Path, len(ps))
	for i, p := range ps {
		ds[i] = p.Dir()
	}
	return ds
}

func (ps PathList) Dir() (Path, error) {
	dir := ps[0].Dir()
	for _, p := range ps[1:] {
		if p.Dir() != dir {
			return _PATH_EMPTY, NO_COMMON_DIR.new(_PATH_EMPTY, _FLAG_EMPTY)
		}
	}
	return dir, nil
}

func (ps PathList) Common() Path {
	splitted := make([][]string, len(ps))
	common := "/"

	for i, p := range ps {
		splitted[i] = dissect(p)
	}

	for n, l := 0, len(splitted[0]); n < l; n++ {
		next := splitted[0][n]

		for _, parts := range splitted[1:] {
			if parts[n] != next {
				return Path(common)
			}
		}

		common += next
	}

	return Path(common)
}

func (ps PathList) String() string {
	s := ""
	for _, p := range ps {
		s += string(os.PathListSeparator) + string(p)
	}
	return s
}
