package fsv

import "os"

// TODO: flag for wiping hdd space
// https://www.socketloop.com/tutorials/golang-secure-file-deletion-with-wipe-example

// Rm removes the file located at p. Allowed flags:
//		- r (recursive): Deletes directories, including any files, subdirectories
//				etc.
//    - w (wipe): Before deleting, the underlying inode is overwritten with an
//        empty byte-slice
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

func rm(p Path, flags struct{ r, w bool }) error {
	info, err := p.Info()
	if os.IsNotExist(err) {
		return nil
	} else if err != nil {
		return err
	}

	if isDir := info.IsDir(); !flags.r && isDir {
		return MISSING_REC_FLAG.new(_PATH_EMPTY, _FLAG_EMPTY)
	} else if flags.r && isDir {
		return os.RemoveAll(string(p))
	}

	if flags.w {
		size := info.Size()
		emptied := make([]byte, size)
		// TODO: overwrite only once?
		err = p.WriteBytes(emptied)
		if err != nil {
			return err
		}
	}

	return os.Remove(string(p))
}

func analyzeFlagsRm(flagrunes []rune) struct{ r, w bool } {
	flags := struct{ r, w bool }{false, false}

	for _, f := range flagrunes {
		switch f {
		case 'r':
			flags.r = true
		case 'w':
			flags.w = true
		default:
			panic(INVALID_FLAG.new(_PATH_EMPTY, f))
		}
	}

	return flags
}
