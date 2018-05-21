package fsv

import "io/ioutil"

// --------------------------------- reads  --------------------------------- //

func (p Path) ReadString() (string, error) {
	bytes, err := p.ReadBytes()
	return string(bytes), err
}

func (p Path) ReadLines() ([]string, error) {
	contents, err := p.ReadString()
	return splitRegex(contents, "\n"), err
}

func (p Path) ReadBytes() ([]byte, error) {
	return ioutil.ReadFile(string(p))
}

// -------------------------------- writes  --------------------------------- //

func (p Path) WriteString(contents string) error {
	return p.WriteBytes([]byte(contents))
}

func (p Path) WriteLines(lines []string) error {
	return p.WriteString(joinWith(lines, "\n"))
}

func (p Path) WriteBytes(bytes []byte) error {
	nfo, err := p.Info()
	if err != nil {
		return err
	}

	return ioutil.WriteFile(string(p), bytes, nfo.Mode())
}
