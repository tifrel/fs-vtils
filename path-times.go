package fsv

import (
	"time"

	"github.com/djherbis/times"
)

// ---------------- modification, creation and access times ----------------- //

// Times returns access, change, modification and birth time (in this order) of
// the file at p. On Plan9 or Windows versions older than and including XP
// change times cannot be retrieved. Birth times can only be retrieved on
// Windows, FreeBSD, NetBSD and Darwin (macOS). If a time cannot be retrieved,
// this method will return Epoch (1970-01-01) for that value.
func (p Path) Times() (atime, ctime, mtime, btime time.Time, err error) {
	// https://github.com/djherbis/times
	timeSpec, err := times.Stat(string(p))
	if err != nil {
		return
	}

	mtime = timeSpec.ModTime()
	atime = timeSpec.AccessTime()
	if timeSpec.HasChangeTime() {
		ctime = timeSpec.ChangeTime()
	}
	if timeSpec.HasBirthTime() {
		btime = timeSpec.BirthTime()
	}

	return
}

// Atime returns the access time of a file. The access time of a file is updated
// when a file is opened.
func (p Path) Atime() (time.Time, error) {
	atime, _, _, _, err := p.Times()
	return atime, err
}

// Ctime returns the change time of a file. The change time of a file is updated
// when the file is modified, including metadata e.g. permissions, owner etc.
// Not supported on Plan9 or Windows version older than (and including) XP.
func (p Path) Ctime() (time.Time, error) {
	_, ctime, _, _, err := p.Times()
	if err != nil {
		return time.Time{}, err
	} else if ctime.IsZero() {
		return ctime, MISSING_OS_SUPPORT.new(_PATH_EMPTY, _FLAG_EMPTY)
	}
	return ctime, err
}

// Mtime returns the mod time of a file. The mod time of a file is updated when
// the file is modified, excluding metadata. (Contents modification only)
func (p Path) Mtime() (time.Time, error) {
	_, _, mtime, _, err := p.Times()
	return mtime, err
}

// Btime returns the birth time of a file. Birth time of a file is never
// updated. Supported only on Windows, NetBSD, FreeBSD and Darwin (macOS).
func (p Path) Btime() (time.Time, error) {
	_, _, _, btime, err := p.Times()
	if err != nil {
		return time.Time{}, err
	} else if btime.IsZero() {
		return btime, MISSING_OS_SUPPORT.new(_PATH_EMPTY, _FLAG_EMPTY)
	}
	return btime, err
}

// SinceAccess return the time.Duration since the access time has changed.
func (p Path) SinceAccess() (time.Duration, error) {
	return since(p.Atime())
}

// SinceChange return the time.Duration since the change time has changed.
func (p Path) SinceChange() (time.Duration, error) {
	return since(p.Ctime())
}

// SinceMod return the time.Duration since the mod time has changed.
func (p Path) SinceMod() (time.Duration, error) {
	return since(p.Mtime())
}

// SinceBirth return the time.Duration since the birth time has changed.
func (p Path) SinceBirth() (time.Duration, error) {
	return since(p.Btime())
}

func since(t time.Time, err error) (time.Duration, error) {
	if err != nil {
		return 0, err
	}

	return time.Now().Sub(t), nil
}
