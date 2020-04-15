package main

import (
	"bytes"
	"io"
	"io/ioutil"
	"log"
	"sync"

	"github.com/fsmiamoto/verificador-ortografico-golang/pkg/dictionary"
)

type ChunkResult struct {
	result io.Reader
	id     int
}

func VerificadorOrtograficoConcorrente(r io.Reader, dict *dictionary.Dictionary) io.Reader {
	const chunkSize = 10 // lines

	file, err := ioutil.ReadAll(r)
	if err != nil {
		log.Fatal(err)
	}

	chunks := bytes.SplitAfterN(file, []byte("\n"), chunkSize)

	var wg sync.WaitGroup

	results := make([]io.Reader, len(chunks))

	for i := range chunks {
		b := chunks[i]
		buffer := bytes.NewBuffer(b)
		wg.Add(1)
		go func(id int) {
			defer wg.Done()
			result := VerificadorOrtografico(buffer, dict)
			results[id] = result
		}(i)
	}

	wg.Wait()

	result := &bytes.Buffer{}
	for i := range results {
		result.Write(results[i].(*bytes.Buffer).Bytes())
	}

	return result
}
