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
