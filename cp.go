package fsv

import (
	"io"
	"os"
)

// TODO: cpDir, cpSymlink, cpFile

// Cp creates a file at target and writes the contents of p to it. Allowed
// flags:
//		- d (dereference): Copies the contents of a symlinks target instead of the
//				link itself. Dereferencing happens recursively until a non-symlink is
//				found for copying.
//		- f (force):       Removes existing files/directories at target.
//		- p (parent):      Creates any dirs necessary to accomodate target.
//		- r (recursive):   Copies directories, including any files, subdirectories
//				etc.
// Failed attempts of writing to the target will trigger deletion of the target.
// If deletion fails, a panic will occur.
// Be aware that the combination of the d and r flags may lead to a circular
// structure, eventually causing a stack and/or drive overflow!
func (p Path) Cp(target Path, flags ...rune) error {
	_flags := analyzeFlagsCp(flags)
	pnc := recover()
	if pnc != nil {
		err, ok := pnc.(error)
		if ok {
			return err
		}
		panic(pnc)
	}

	return cp(p, target, _flags)
}

func cp(p, target Path, flags cpFlags) error {
	err := prepareLocation(target, flags.toMk())
	if err != nil {
		return err
	}

	// handling symlinks
	isSymlink, err := p.IsSymlink()
	if err != nil {
		return err
	} else if isSymlink && !flags.d {

		dereferenced, err := p.Follow()
		if err != nil {
			return err
		}
		return ln(dereferenced, target, flags.toLn())

	} else if isSymlink && flags.d {

		dereferenced, err := p.Target()
		if err != nil {
			return err
		}
		return ln(dereferenced, target, flags.toLn())

	}

	// handling directories
	info, err := p.Info()
	if err != nil {
		return err
	} else if !flags.r && info.IsDir() {
		return MISSING_REC_FLAG.new(_PATH_EMPTY, _FLAG_EMPTY)
	} else if flags.r && info.IsDir() {
		es, err := p.Ls()
		if err != nil {
			return err
		}

		err = mkDir(target, flags.toMk())
		if err != nil {
			return err
		}
		return es.Each(func(e Path) error {
			return cp(e, target.Append(e.Base()), flags)
		})

	} else {
		// handling regular files
		err = target.MkFile(info.Mode())
		if err != nil {
			return err
		}
	}

	return cpContents(p, target)
}

func cpContents(sp Path, tp Path) error {
	// Assumes that destination and target already exist. Takes care of opening
	// and closing the files. Deletes target if operation fails.
	var (
		buf  = make([]byte, bufSize)
		sErr error
		tErr error
	)

	s, sErr := os.OpenFile(string(sp), os.O_RDONLY, 0644)
	if sErr != nil {
		return sErr
	}
	defer s.Close()
	t, tErr := os.OpenFile(string(tp), os.O_WRONLY, 0644)
	if tErr != nil {
		return tErr
	}
	defer t.Close()

	for sErr != io.EOF {

		_, sErr = s.Read(buf)
		if sErr != nil && sErr != io.EOF {

			err := tp.Rm()
			if err != nil && tp.Exists() {
				// TODO: somehow transmit information about tErr when panicking
				panic(err)
			}

			return sErr
		}

		_, tErr = t.Write(buf)
		if tErr != nil {

			err := tp.Rm()
			if err != nil && tp.Exists() {
				// TODO: somehow transmit information about tErr when panicking
				panic(err)
			}

			return tErr
		}

	}

	return nil
}

type cpFlags struct{ d, f, r, p bool }

func (cp cpFlags) toMk() mkFlags {
	return mkFlags{cp.f, cp.p}
}
func (cp cpFlags) toLn() lnFlags {
	return lnFlags{cp.d, cp.f, false, cp.p}
}

func analyzeFlagsCp(flagrunes []rune) cpFlags {
	flags := cpFlags{false, false, false, false}

	for _, f := range flagrunes {
		switch f {
		case 'd': // dereference
			flags.d = true
		case 'f': // force
			flags.f = true
		case 'r': // recursive
			flags.r = true
		case 'p': // parent
			flags.p = true
		default:
			panic(INVALID_FLAG.new(_PATH_EMPTY, f))
		}
	}

	return flags
}
