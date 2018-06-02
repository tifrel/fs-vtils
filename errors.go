package fsv

type errNumber uint8

const (
	no_INVALID_FLAG errNumber = 1 + iota
	no_OCCUPIED_PATH
	no_MISSING_TARGETDIR
	no_MISSING_REC_FLAG
	no_MISSING_OS_SUPPORT
	no_UNKNOWN_ERR
)

var (
	INVALID_FLAG       = Err{no_INVALID_FLAG, _PATH_EMPTY, _FLAG_EMPTY}
	OCCUPIED_PATH      = Err{no_OCCUPIED_PATH, _PATH_EMPTY, _FLAG_EMPTY}
	MISSING_TARGETDIR  = Err{no_MISSING_TARGETDIR, _PATH_EMPTY, _FLAG_EMPTY}
	MISSING_REC_FLAG   = Err{no_MISSING_REC_FLAG, _PATH_EMPTY, _FLAG_EMPTY}
	MISSING_OS_SUPPORT = Err{no_MISSING_OS_SUPPORT, _PATH_EMPTY, _FLAG_EMPTY}
	UNKNOWN_ERR        = Err{no_UNKNOWN_ERR, _PATH_EMPTY, _FLAG_EMPTY}
)

const (
	_FLAG_EMPTY rune = 0
	_PATH_EMPTY Path = ""
)

type Err struct {
	Id   errNumber
	Path Path
	Flag rune
}

func (e Err) Error() string {
	switch e.Id {
	case no_INVALID_FLAG:
		return "Invalid flag: " + string(e.Flag)

	case no_OCCUPIED_PATH:
		return "Occupied path: " + string(e.Path)

	case no_MISSING_TARGETDIR:
		return "Inexistent target directory: " + string(e.Path)

	case no_MISSING_REC_FLAG:
		return "Copying/Moving dir requires recursive flag."

	case no_MISSING_OS_SUPPORT:
		return "Operating system does not support this operation."

	case no_UNKNOWN_ERR:
		return "Unkown error."

	default:
		panic("Tried to call Error() on unidentifiable error :(")
	}

}

func (proto Err) new(path Path, flag rune) Err {
	return Err{
		proto.Id,
		path,
		flag,
	}
}

// TODO: errors for PathList.Each, to which the returned errors will be appended

func (proto Err) IsTypeOf(e error) bool {
	switch {
	case e == nil && error(proto) == nil:
		return true
	case e != nil && error(proto) == nil:
		return false
	case e == nil && error(proto) != nil:
		return false
	}

	fsve, ok := e.(Err)
	if !ok {
		return false
	}

	if fsve.Id == proto.Id {
		return true
	} else {
		return false
	}
}
