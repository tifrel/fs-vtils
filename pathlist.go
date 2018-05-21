package fsv

import (
	"errors"
	"os"
)

type PathList []Path

func (in PathList) Filter(pred func(Path) bool) PathList {
	out, i := make([]Path, len(in)), 0

	for _, p := range in {
		if pred(p) {
			out[i] = p
			i++
		}
	}

	return out[:i]
}

func (ps PathList) Names() []string {
	out := make([]string, len(ps))
	for i, p := range ps {
		out[i] = string(p.Base())
	}
	return out
}

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

func (ps PathList) Each(fn func(Path) error) error {
	var ERR error = nil

	for _, p := range ps {
		err := fn(p)

		if err != nil {
			msg := string(p) + " :: " + err.Error()

			if ERR == nil {
				ERR = errors.New(msg)
			} else {
				ERR = errors.New(ERR.Error() + "\n" + msg)
			}
		}

	}

	return ERR
}
