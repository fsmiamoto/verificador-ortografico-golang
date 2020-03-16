package main

import (
	"bufio"
	"compress/gzip"
	"io"
	"os"
	"testing"

	"github.com/fsmiamoto/verificador-ortografico-golang/pkg/dictionary"
)

func TestVerificadorOrtografico(t *testing.T) {
	dict := setupDict()
	input, _ := os.Open("testes/exemplo.txt")

	got := VerificadorOrtografico(input, dict)

	expect, _ := os.Open("testes/exemplo_esperado.txt")
	defer expect.Close()

	assertEqualContent(t, got, expect)
}

func BenchmarkVerificadorOrtografico(b *testing.B) {
	const n = 10

	dict := setupDict()
	input, _ := os.Open("test-inputs/brascubas.txt")

	b.ResetTimer()
	for i := 0; i < n; i++ {
		VerificadorOrtografico(input, dict)
	}
}

func setupDict() dictionary.Dictionary {
	dictFile, _ := os.Open(dictFileName)
	unziped, _ := gzip.NewReader(dictFile)

	dict := dictionary.New()

	dict.Parse(unziped)

	return dict
}

func assertEqualContent(t *testing.T, got, want io.Reader) {
	scanGot, scanWant := bufio.NewScanner(got), bufio.NewScanner(want)

	for scanGot.Scan() {
		// Different content length
		if !scanWant.Scan() {
			t.Errorf("readers have different length")
		}

		if g, w := scanGot.Text(), scanWant.Text(); g != w {
			t.Errorf("got %q but want %q", g, w)
		}
	}

	// Different content length
	if scanWant.Scan() {
		t.Errorf("readers have different length")
	}
}
