package fsv

import "os"

// later TODO: flag for wiping hdd space
// https://www.socketloop.com/tutorials/golang-secure-file-deletion-with-wipe-example

// TODO: testing

// Rm removes the file located at p. Allowed flags:
//		- r (recursive): Deletes directories, including any files, subdirectories
//				etc.
func (p Path) Rm(flags ...rune) error {
	_flags := analyzeFlagsRm(flags)
	pnc := recover()
	if pnc != nil {
		err, ok := pnc.(error)
		if ok {
			return err
		}
		panic(pnc)
	}

	return rm(p, _flags)
}

func rm(p Path, flags struct{ r bool }) error {
	if !p.Exists() {
		return nil
	}

	isDir, err := p.IsDir()
	if err != nil {
		return err
	} else if !flags.r && isDir {
		return MISSING_REC_FLAG.new(_PATH_EMPTY, _FLAG_EMPTY)
	} else if flags.r {
		return os.RemoveAll(string(p))
	}
	return os.Remove(string(p))
}

func analyzeFlagsRm(flagrunes []rune) struct{ r bool } {
	flags := struct{ r bool }{false}

	for _, f := range flagrunes {
		switch f {
		case 'r':
			flags.r = true
		default:
			panic(INVALID_FLAG.new(_PATH_EMPTY, f))
		}
	}

	return flags
}
