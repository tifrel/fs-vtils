package fsv

import (
	"io"
	"os"
	"regexp"
)

// --------------------------------- regexp --------------------------------- //

func splitRegex(str string, re string) []string {
	is := regexp.MustCompile(re).FindAllStringIndex(str, -1)
	last := 0
	result := make([]string, len(is)+1)
	for i, el := range is {
		result[i] = str[last:el[0]]
		last = el[1]
	}
	result[len(is)] = str[last:len(str)]
	return result
}

func joinWith(strs []string, del string) string {
	// adds a trailing new line, which is useful in this case
	r := ""
	for _, s := range strs {
		r += s + del
	}
	return r
}

// ----------------------------- error handling ----------------------------- //

func panicCheck(e error) {
	if e != nil {
		panic(e)
	}
}

func checkRead(e error) {
	if e != nil && e != io.EOF {
		panic(e)
	}
}

// ----------------------------- file utilities ----------------------------- //

func isOsFileInfo(name string) bool {
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

// ------------------------------ buffer sizes ------------------------------ //

// buffer size 64 KiB => should be good for text files
//    (fits 800 lines at 80 width)
var bufSize uint = 1 << 16

// Sets the buffersize for operations like reading, writing and copying.
func SetBufferSize(n uint) {
	bufSize = n
}

// Restores a buffersize of 65536 bytes (64KiB)
func ResetBufferSize() {
	bufSize = 1 << 16
}

// ------------------------- fs manipulation utils -------------------------- //

// force + existing => delete; not force + existing => error; else => nil
func create(target Path, force, parent bool, perm os.FileMode) error {
	err := prepareBase(target, force)
	if err != nil {
		return err
	}

	err = prepareDir(target, parent)
	if err != nil {
		return err
	}

	return target.MkFile(perm)
}

func prepareBase(target Path, force bool) error {
	switch {
	case !target.Exists():
		return nil
	case force:
		return target.Rm()
	case !force:
		return OCCUPIED_PATH.new(target, _FLAG_EMPTY)
	}

	return UNKNOWN_ERR.new(_PATH_EMPTY, _FLAG_EMPTY)
}

func prepareDir(target Path, parent bool) error {

	// check for parent dir existence and process parent flag
	dir := target.Dir()
	if parent && !dir.Exists() {
		os.MkdirAll(string(dir), 0755)
	} else if !parent && !dir.Exists() {
		return MISSING_TARGETDIR.new(dir, _FLAG_EMPTY)
	}

	return nil
	// later TODO: this fails if there is a file anywhere on the path blocking
	// creation of a dir
}

func prepareLocation(p Path, flags mkFlags) error {
	targetExists := p.Exists()
	if targetExists && !flags.f {
		return OCCUPIED_PATH.new(p, _FLAG_EMPTY)
	} else if targetExists && flags.f {
		os.RemoveAll(string(p))
	}

	parent := p.Dir()
	if !parent.Exists() && !flags.p {
		return MISSING_TARGETDIR.new(parent, _FLAG_EMPTY)
	} else if !parent.Exists() && flags.p {
		err := mkDir(p.Dir(), flags)
		if err != nil {
			return err
		}
	}

	parentIsDir, err := parent.IsDir()
	if err != nil {
		return err
	} else if !parentIsDir && !flags.f {
		return OCCUPIED_PATH.new(parent, _FLAG_EMPTY)
	} else if !parentIsDir && !flags.p {
		return MISSING_TARGETDIR.new(parent, _FLAG_EMPTY)
	}

	return mkDir(parent, flags)
}

func hasFlags(flagstrings []string, name string) bool {
	if len(flagstrings) == 0 {
		return false
	} else if len(flagstrings) > 1 {
		panic(name + " :: Only one flag string allowed")
	} else if flagstrings[0] == "" {
		panic(name + " :: Empty flag string")
	}

	return true
}
