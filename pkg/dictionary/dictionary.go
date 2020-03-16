package dictionary

import (
	"bufio"
	"io"
	"strings"
)

type Dictionary map[string]struct{}

func New() Dictionary {
	d := make(map[string]struct{})
	return Dictionary(d)
}

// Parse takes a io.Reader and add's each line of it as an entry on the dictionary
func (d Dictionary) Parse(r io.Reader) error {
	s := bufio.NewScanner(r)

	for s.Scan() {
		if err := s.Err(); err != nil {
			return err
		}

		word := s.Text()
		d[word] = struct{}{}
	}

	return nil
}

// HasWord tells if a dictionary has a word or not
func (d Dictionary) HasWord(w string) bool {
	_, found := d[strings.ToLower(w)]
	return found
}
