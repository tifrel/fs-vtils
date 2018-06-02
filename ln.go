package fsv

import "os"

// TODO: testing
// TODO: d flag -> dereferencing (recursive)

// Ln creates a symlink at target, pointing to p. Allowed flags:
//		- f (force):    Removes existing files/directories at target.
//		- h (hardlink): Creates a hardlink instead of a symlink.
//		- p (parent):   Creates any dirs necessary to accomodate target.
func (p Path) Ln(target Path, flags ...rune) error {
	_flags := analyzeFlagsLn(flags)
	pnc := recover()
	if pnc != nil {
		err, ok := pnc.(error)
		if ok {
			return err
		}
		panic(pnc)
	}

	return ln(p, target, _flags)
}

func ln(p, target Path, flags lnFlags) error {
	err := prepareLocation(target, flags.toMk())
	if err != nil {
		return err
	}

	if flags.d {
		p, err = p.Target()
	}

	if flags.h {
		return os.Link(string(p), string(target))
	}
	return os.Symlink(string(p), string(target))
}

type lnFlags struct{ d, f, h, p bool }

func (ln lnFlags) toMk() mkFlags {
	return mkFlags{ln.f, ln.p}
}

func analyzeFlagsLn(flagrunes []rune) lnFlags {
	flags := lnFlags{false, false, false, false}

	for _, f := range flagrunes {
		switch f {
		case 'd':
			flags.d = true
		case 'f':
			flags.f = true
		case 'h':
			flags.h = true
		case 'p':
			flags.p = true
		default:
			panic(INVALID_FLAG.new(_PATH_EMPTY, f))
		}
	}

	return flags
}
