package fsv

import (
	"hash"
	"hash/fnv"
	"io"
	"os"
	pathPkg "path"

	"github.com/fsnotify/fsnotify"
)

type Path string

// -------------------------- information methods --------------------------- //

// TODO: Grep(re string) (lineNo int, startByte int, stopByte int, line string, match string)

func (p Path) Exists() bool {
	_, err := os.Stat(string(p))
	return err == nil || os.IsExist(err)
}

func (p Path) Info() (os.FileInfo, error) {
	return os.Lstat(string(p))
}

func (p Path) Mode() (os.FileMode, error) {
	nfo, err := p.Info()
	if err != nil {
		return 0, err
	}
	return nfo.Mode(), nil
}

func (p Path) IsFile() (bool, error) {
	nfo, err := p.Info()
	if err != nil {
		return false, err
	}
	return nfo.Mode().IsRegular(), nil
}

func (p Path) IsSymlink() (bool, error) {
	nfo, err := p.Info()
	if err != nil {
		return false, err
	}
	return (nfo.Mode()&os.ModeSymlink != 0), nil
}

func (p Path) Follow() (Path, error) {
	target, err := os.Readlink(string(p))
	return Path(target), err
}

func (p Path) Target() (Path, error) {
	isLn, err := p.IsSymlink()
	if err != nil {
		return Path(""), err
	}

	for isLn {
		p, err = p.Follow()
		if err != nil {
			return Path(""), err
		}

		isLn, err = p.IsSymlink()
		if err != nil {
			return Path(""), err
		}
	}

	return p, nil
}

func (p Path) IsDir() (bool, error) {
	nfo, err := p.Info()
	if err != nil {
		return false, err
	}
	return nfo.IsDir(), nil
}

func (p Path) Ls() (PathList, error) {
	d, err := os.Open(string(p))
	defer d.Close()
	if err != nil {
		return nil, err
	}

	es, err := d.Readdirnames(-1)
	if err != nil {
		return nil, err
	}

	var ps PathList = make([]Path, len(es))
	for i, e := range es {
		ps[i] = p.AppendStr(e)
	}

	return ps, nil
}

func (p Path) IsOsFile() bool {
	name := string(p.Base())

	for _, n := range []string{
		".DS_Store",
		"___MACOSX",
		"desktop.ini",
		"Thumbs.db",
		"thumbs.db",
	} {
		if name == n {
			return true
		}
	}

	return false

}

func (p Path) IsHidden() bool {
	// only for unix so far => TODO: other OS's
	return string(p.Base())[0] == '.'
}

func (p Path) IsVisible() bool {
	return !p.IsHidden()
}

// ----------------------- path manipulation methods ------------------------ //

func (p Path) Append(x Path) Path {
	return Path(pathPkg.Join(string(p), string(x)))
}
func (p Path) Base() Path {
	return Path(pathPkg.Base(string(p)))
}
func (p Path) Dir() Path {
	return Path(pathPkg.Dir(string(p)))
}

func (p Path) AppendStr(x string) Path {
	return Path(pathPkg.Join(string(p), x))
}
func (p Path) BaseStr() string {
	return pathPkg.Base(string(p))
}
func (p Path) DirStr() string {
	return pathPkg.Dir(string(p))
}

// --------------------------------- makes ---------------------------------- //

func (p Path) MkWatch() (*fsnotify.Watcher, error) {
	w, err := fsnotify.NewWatcher()
	if err != nil {
		return w, err
	}

	w.Add(string(p))

	return w, nil
}

// --------------------------- comparison methods --------------------------- //

func (p Path) Hash() (hash.Hash64, error) {
	// parallel/ optimized execution because this is an expensive operation
	// io + hashing
	var (
		h         = fnv.New64()
		ch        = make(chan []byte)
		ERR error = nil
		err error = nil
	)

	f, err := os.Open(string(p))
	if err != nil {
		return nil, err
	}
	defer f.Close()

	go func() {
		defer close(ch)

		var (
			buf = make([]byte, 8)
			n   int
		)

		for n != 0 && err != io.EOF {
			n, err = f.Read(buf)
			if err != nil {
				ERR = err
				return
			}
			ch <- buf
		}

	}()

	for bytes := range ch {
		h.Write(bytes)
	}

	return h, ERR
}

func (p1 Path) SameHashAs(p2 Path) (bool, error) {
	h1, err := p1.Hash()
	if err != nil {
		return false, err
	}
	h2, err := p2.Hash()
	if err != nil {
		return false, err
	}

	return h1.Sum64() == h2.Sum64(), nil
}

func (p1 Path) SameContentsAs(p2 Path) (bool, error) {
	// low-alloc, buffered because io is expensive operation
	f1, err := os.Open(string(p1))
	if err != nil {
		return false, err
	}
	defer f1.Close()
	f2, err := os.Open(string(p2))
	if err != nil {
		return false, err
	}
	defer f2.Close()

	var (
		buf1       = make([]byte, bufSize)
		buf2       = make([]byte, bufSize)
		err1 error = nil
		err2 error = nil
		n1   int
		n2   int

		i int
		b byte
	)

	// TODO: this loop might need some rework
	for err1 != io.EOF && err2 != io.EOF {

		n1, err1 = f1.Read(buf1)
		if err1 != nil && err1 != io.EOF {
			return false, err1
		}
		n2, err2 = f2.Read(buf2)
		if err2 != nil && err2 != io.EOF {
			return false, err2
		}

		if n1 != n2 {
			// short circuiting + guarding against index-out-of-range panics
			//    in upcoming loop
			return false, nil
		}

		for i, b = range buf1 {
			if b != buf2[i] {
				return false, nil
			}
		}
	}

	return true, nil
}

func (p1 Path) SameInfoAs(p2 Path) (bool, error) {
	// ATTENTION: Windows compares by path, which opens up the possibility of
	//    erroneously false results
	f1, err := os.Open(string(p1))
	if err != nil {
		return false, err
	}
	n1, err := f1.Stat()
	if err != nil {
		return false, err
	}

	f2, err := os.Open(string(p1))
	if err != nil {
		return false, err
	}
	n2, err := f2.Stat()
	if err != nil {
		return false, err
	}

	return os.SameFile(n1, n2), nil
}
