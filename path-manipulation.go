package fsv

import (
	pathPkg "path"
	"path/filepath"
)

// ----------------------- path manipulation methods ------------------------ //

// Extend returns p + delimiter + x
func (p Path) Extend(x Path) Path {
	return Path(pathPkg.Join(string(p), string(x)))
}

// Base returns only the last part of a string, e.g.:
// 		- Path("/Users/admin/Documents").Base() == Path("Documents")
func (p Path) Base() Path {
	return Path(pathPkg.Base(string(p)))
}

// Dir returns all parts of a string except the base:
// 		- Path("/Users/admin/Documents").Base() == Path("/Users/admin")
func (p Path) Dir() Path {
	return Path(pathPkg.Dir(string(p)))
}

// ExtendStr is like Extend, but takes a string as argument instead of a Path.
func (p Path) ExtendStr(x string) Path {
	return Path(pathPkg.Join(string(p), x))
}

// BaseStr is like Base, but returns a string instead of a Path.
func (p Path) BaseStr() string {
	return pathPkg.Base(string(p))
}

// DirStr is like Dir, but returns a string instead of a Path.
func (p Path) DirStr() string {
	return pathPkg.Dir(string(p))
}

// RelativeTo returns the Path that p has relative to dir.
func (p Path) RelativeTo(dir Path) (Path, error) {
	rel, err := filepath.Rel(string(dir), string(p))
	return Path(rel), err
}

// RelativeToStr is like RelativeTo, but takes a string as argument instead of a
// Path.
func (p Path) RelativeToStr(dir string) (Path, error) {
	rel, err := filepath.Rel(dir, string(p))
	return Path(rel), err
}

// Dissect returns a []Path of dirnames (and finally a filename) for p.
func (p Path) Dissect() []Path {
	splitted := []Path{p.Base()}
	p = p.Dir()
	for p != "." && p != "/" && p != "" {
		splitted = append([]Path{p.Base()}, splitted...)
		p = p.Dir()
	}
	return splitted
}

// DissectStr is like Dissect, but returns a []string.
func (p Path) DissectStr() []string {
	splitted := []string{p.BaseStr()}
	p = p.Dir()
	for p != "." && p != "/" && p != "" {
		splitted = append([]string{p.BaseStr()}, splitted...)
		p = p.Dir()
	}
	return splitted
}

func (p Path) Split() (dir, base Path) {
	base = p.Base()
	dir = p.Dir()
	return
}

func (p Path) SplitStr() (dir, base string) {
	base = p.BaseStr()
	dir = p.DirStr()
	return
}
