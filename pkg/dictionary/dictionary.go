package dictionary

import (
	"bufio"
	"io"
	"sort"
	"strings"
)

type Dictionary []string

// New takes a io.Reader and add's each line of it as an entry on the dictionary
// returning a pointer to the Dictionary
func New(r io.Reader) (*Dictionary, error) {
	s := bufio.NewScanner(r)

	var d Dictionary

	for s.Scan() {
		if err := s.Err(); err != nil {
			return nil, err
		}

		d = append(d, strings.ToLower(s.Text()))
	}

	sort.Strings(d)

	return &d, nil
}

// HasWord tells if a dictionary has a word or not
func (d Dictionary) HasWord(w string) bool {
	if d == nil {
		return false
	}

	lowercased := strings.ToLower(w)

	index := sort.SearchStrings(d, lowercased)
	if index < len(d) && d[index] == lowercased {
		return true
	}

	return false
}
