package fsv

import (
	"os"

	"github.com/fsnotify/fsnotify"
)

// MkDir creates a directory at p. Allowed flags:
//		- f (force):  Removes existing files/directories at p.
//		- p (parent): Creates any dirs necessary to accomodate target.
func (p Path) MkDir(flags ...rune) error {
	_flags := analyzeFlagsMk(flags)
	pnc := recover()
	if pnc != nil {
		err, ok := pnc.(error)
		if ok {
			return err
		}
		panic(pnc)
	}

	return mkDir(p, _flags)
}

// MkFile creates a (regular) file at p. Allowed flags:
//		- f (force) : Removes existing files/directories at p.
//    - p (parent): Creates any dirs necessary to accomodate target.
func (p Path) MkFile(perm os.FileMode, flags ...rune) error {
	_flags := analyzeFlagsMk(flags)
	pnc := recover()
	if pnc != nil {
		err, ok := pnc.(error)
		if ok {
			return err
		}
		panic(pnc)
	}

	return mkFile(p, perm, _flags)
}

func mkFile(p Path, perm os.FileMode, flags mkFlags) error {
	err := prepareLocation(p, flags)
	if err != nil {
		return err
	}

	f, err := os.Create(string(p))
	if err != nil {
		return err
	}
	defer f.Close()
	return f.Chmod(perm)
}

// TODO: use prepareLocation() if possible
func mkDir(p Path, flags mkFlags) error {
	targetExists := p.Exists()
	targetIsDir, err := p.IsDir()
	if targetExists && err != nil {
		return err
	} else if targetExists && targetIsDir {
		// nothing to do
		return nil
	} else if targetExists && !flags.f {
		return OCCUPIED_PATH.new(p, _FLAG_EMPTY)
	} else if targetExists && flags.f {
		os.RemoveAll(string(p))
	}

	parent := p.Dir()
	parentExists := parent.Exists()
	parentIsDir, err := parent.IsDir()
	if parentExists && err != nil {
		return err
	}
	switch { // switch (as of now) needed for break to work
	case parentExists && parentIsDir:
		break // stop crawling
	case !flags.p:
		return MISSING_TARGETDIR.new(p.Dir(), _FLAG_EMPTY)
	case flags.p:
		// recursive crawl with appropriate 'f' flag
		err = mkDir(parent, flags)
		if err != nil {
			return err
		}
	}

	return os.Mkdir(string(p), 0755)
}

type mkFlags struct{ f, p bool }

func analyzeFlagsMk(flagrunes []rune) mkFlags {
	flags := mkFlags{false, false}

	for _, f := range flagrunes {
		switch f {
		case 'f':
			flags.f = true
		case 'p':
			flags.p = true
		default:
			panic(INVALID_FLAG.new(_PATH_EMPTY, f))
		}
	}

	return flags
}

// --------------------------------- watch ---------------------------------- //

// MkWatch returns a Watch from the fsnotify Package, which can be used to
// efficiently receive notifications for any events occuring on the specified
// file or directory. For further information refer to the fsnotify
// documentation.
func (p Path) MkWatch() (*fsnotify.Watcher, error) {
	w, err := fsnotify.NewWatcher()
	if err != nil {
		return w, err
	}

	w.Add(string(p))

	return w, nil
}
