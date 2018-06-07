package fsv

import (
	"io/ioutil"
	"os"
)

// --------------------------------- reads ---------------------------------- //

// ReadString reads the contents of the file at p and returns them as a string.
func (p Path) ReadString() (string, error) {
	bytes, err := p.ReadBytes()
	return string(bytes), err
}

// ReadRunes reads the contents of the file at p and returns them as a []rune.
func (p Path) ReadRunes() ([]rune, error) {
	bytes, err := p.ReadBytes()
	return []rune(string(bytes)), err
}

// ReadLines reads the contents of the file at p and returns them as a []string.
func (p Path) ReadLines() ([]string, error) {
	contents, err := p.ReadString()
	return splitRegex(contents, "\n"), err
}

// ReadBytes reads the contents of the file at p and returns them as a []byte.
func (p Path) ReadBytes() ([]byte, error) {
	return ioutil.ReadFile(string(p))
}

// --------------------------------- writes --------------------------------- //

// WriteString (over)writes the contents of the file located at p with a string.
func (p Path) WriteString(contents string) error {
	return p.WriteBytes([]byte(contents))
}

// WriteRunes (over)writes the contents of the file located at p with a []rune.
func (p Path) WriteRunes(runes []rune) error {
	return p.WriteBytes([]byte(string(runes)))
}

// WriteLines (over)writes the contents of the file located at p with a []string.
// This includes a trailing newline character.
func (p Path) WriteLines(lines []string) error {
	return p.WriteBytes([]byte(joinWith(lines, "\n")))
}

// WriteBytes (over)writes the contents of the file located at p with a []byte.
func (p Path) WriteBytes(bytes []byte) error {
	nfo, err := p.Info()
	if err != nil {
		return err
	}

	return ioutil.WriteFile(string(p), bytes, nfo.Mode())
}

// -------------------------------- appends --------------------------------- //

// AppendString appends a string to the file located at p.
func (p Path) AppendString(contents string) error {
	return p.WriteBytes([]byte(contents))
}

// AppendString appends a []byte to the file located at p.
func (p Path) AppendRunes(runes []rune) error {
	return p.WriteBytes([]byte(string(runes)))
}

// AppendLines appends a []string to the file located at p.
func (p Path) AppendLines(lines []string) error {
	return p.WriteBytes([]byte(joinWith(lines, "\n")))
}

// AppendBytes appends a []rune to the file located at p.
func (p Path) AppendBytes(bytes []byte) error {
	f, err := os.OpenFile(string(p), os.O_APPEND|os.O_WRONLY, 0600)
	if err != nil {
		return err
	}
	defer f.Close()

	_, err = f.Write(bytes)
	return err
}
