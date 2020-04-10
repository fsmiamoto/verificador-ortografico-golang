package main

import (
	"bufio"
	"bytes"
	"compress/gzip"
	"fmt"
	"io"
	"log"
	"os"
	"strings"

	"github.com/fsmiamoto/verificador-ortografico-golang/pkg/dictionary"
)

const (
	dictFileName = "brazilian.utf8.gz"
)

func main() {
	f, err := os.Open(dictFileName)
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	unziped, err := gzip.NewReader(f)
	if err != nil {
		log.Fatal(err)
	}
	defer unziped.Close()

	dict, err := dictionary.New(unziped)
	if err != nil {
		log.Fatal(err)
	}

	result := VerificadorOrtografico(os.Stdin, dict)
	fmt.Print(result)
}

func VerificadorOrtografico(r io.Reader, dict *dictionary.Dictionary) io.Reader {
	scanner := bufio.NewScanner(r)
	buffer := &bytes.Buffer{}

	for scanner.Scan() {
		// Lê linha inteira e dá split nas palavras
		words := strings.Split(scanner.Text(), " ")

		// Linha vazia
		if len(words) == 1 {
			buffer.WriteString("\n")
			continue
		}

		for i := range words {
			s := sanitizeWord(words[i])
			if !dict.HasWord(s) {
				words[i] = surroundWord(words[i])
			}
		}

		line := strings.Join(words, " ")

		buffer.WriteString(line)
		buffer.WriteString("\n")
	}

	return buffer
}
