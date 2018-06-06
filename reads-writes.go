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

// WriteLines (over)writes the contents of the file located at p with a []string.
func (p Path) WriteLines(lines []string) error {
	return p.WriteString(joinWith(lines, "\n"))
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

// AppendString appends its argument to a file
func (p Path) AppendString(contents string) error {
	return p.WriteBytes([]byte(contents))
}

// AppendLines appends its argument to a file
func (p Path) AppendLines(lines []string) error {
	return p.WriteString(joinWith(lines, "\n"))
}

// AppendBytes appends its argument to a file
func (p Path) AppendBytes(bytes []byte) error {
	f, err := os.OpenFile(string(p), os.O_APPEND|os.O_WRONLY, 0600)
	if err != nil {
		return err
	}
	defer f.Close()

	_, err = f.Write(bytes)
	return err
}
