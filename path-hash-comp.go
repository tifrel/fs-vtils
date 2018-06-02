package fsv

import (
	"hash"
	"hash/fnv"
	"io"
	"os"
)

// ----------------–----- comparison & hashing methods ----------–----------- //

// Hash returns a hash.Hash64 after FNV-1a algorithm for the contents of the
// file located at p.
func (p Path) Hash() (hash.Hash64, error) {
	// parallel/ optimized execution because this is an expensive operation
	// io + hashing
	var (
		h   = fnv.New64()
		ch  = make(chan []byte)
		ERR error
		err error
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

// HasHash checks wether the file located as p as the same hash-sum as the h.
func (p Path) HasHash(h hash.Hash64) (bool, error) {
	h2, err := p.Hash()
	if err != nil {
		return false, err
	}
	return h2.Sum64() == h.Sum64(), nil
}

// SameHashAs checks wether p1 and p2 bath have the same.
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

// SameContentsAs checks wether two file have the same content,
// going byte-by-byte and returning early if possible.
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
		buf1 = make([]byte, bufSize)
		buf2 = make([]byte, bufSize)
		err1 error
		err2 error
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

// SameInfoAs compares the files located at p1 and p2 by their os.FileInfo.
//
// Attention: Windows compares by path in this case, so two files pointing
// at the same inode may erroneously be reported as not the same file.
func (p1 Path) SameInfoAs(p2 Path) (bool, error) {
	n1, err := p1.Info()
	if err != nil {
		return false, err
	}

	n2, err := p1.Info()
	if err != nil {
		return false, err
	}

	return os.SameFile(n1, n2), nil
}
