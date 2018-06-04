package fsv

import "os"

// Mv moves (renames) the file at p to target. Allowed flags:
//		- f (force):     Removes existing files/directories at target.
//		- p (parent):    Creates any dirs necessary to accomodate target.
//		- r (recursive): Moves directories, including any files, subdirectories
//				etc.
func (p Path) Mv(target Path, flags ...rune) error {
	_flags := analyzeFlagsMv(flags)
	pnc := recover()
	if pnc != nil {
		err, ok := pnc.(error)
		if ok {
			return err
		}
		panic(pnc)
	}

	return mv(p, target, _flags)
}

func mv(p Path, target Path, flags mvFlags) error {
	// err := prepareLocation(p, struct{ f, p bool }{flags.f, flags.p})
	err := prepareLocation(target, flags.toMk())
	if err != nil {
		return err
	}

	isDir, err := p.IsDir()
	if err != nil {
		return err
	} else if isDir && !flags.r {
		return MISSING_REC_FLAG.new(_PATH_EMPTY, _FLAG_EMPTY)
	}

	return os.Rename(string(p), string(target))
}

type mvFlags struct{ f, p, r bool }

func (mv mvFlags) toMk() mkFlags {
	return mkFlags{mv.f, mv.p}
}

func analyzeFlagsMv(flagrunes []rune) mvFlags {
	flags := mvFlags{false, false, false}

	for _, f := range flagrunes {
		switch f {
		case 'f':
			flags.f = true
		case 'p':
			flags.p = true
		case 'r':
			flags.r = true
		default:
			panic(INVALID_FLAG.new(_PATH_EMPTY, f))
		}
	}

	return flags
}
