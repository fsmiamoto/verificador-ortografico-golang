package main

import (
	"bufio"
	"compress/gzip"
	"io"
	"os"
	"testing"

	"github.com/fsmiamoto/verificador-ortografico-golang/pkg/dictionary"
)

var dict *dictionary.Dictionary

type testState interface {
	Fatalf(fmt string, args ...interface{})
	Helper()
}

func TestVerificadorOrtografico(t *testing.T) {
	dict := setupDict(t)

	input, _ := os.Open("testes/exemplo.utf8.txt")

	got := VerificadorOrtografico(input, dict)

	expect, _ := os.Open("testes/exemplo_esperado.utf8.txt")
	defer expect.Close()

	assertEqualContent(t, got, expect)
}

func BenchmarkVerificadorOrtografico(b *testing.B) {
	dict := setupDict(b)

	input, err := os.Open("testes/brascubas.utf8.txt")
	if err != nil {
		b.Fatalf("BenchmarkVerificadorOrtografico: %v", err)
	}
	defer input.Close()

	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		VerificadorOrtografico(input, dict)
		input.Seek(0, 0)
	}
}

func BenchmarkVerificadorOrtograficoConcorrente(b *testing.B) {
	dict := setupDict(b)

	input, err := os.Open("testes/brascubas.utf8.txt")
	if err != nil {
		b.Fatalf("BenchmarkVerificadorOrtograficoConcorrente: %v", err)
	}
	defer input.Close()

	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		VerificadorOrtograficoConcorrente(input, dict)
		input.Seek(0, 0)
	}
}

func setupDict(t testState) *dictionary.Dictionary {
	t.Helper()

	if dict != nil {
		return dict
	}

	dictFile, err := os.Open(dictFileName)
	if err != nil {
		t.Fatalf("error: %v", err)
	}

	unziped, err := gzip.NewReader(dictFile)
	if err != nil {
		t.Fatalf("error: %v", err)
	}

	dict, err = dictionary.New(unziped)
	if err != nil {
		t.Fatalf("error: %v", err)
	}

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
