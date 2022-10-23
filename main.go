package main

import (
	"bytes"
	"fmt"
	"io"
	"os"
)

func main() {
	bs, err := io.ReadAll(os.Stdin)
	if err != nil {
		panic(err)
	}
	type array = []byte
	var transposed [][]array
	maxi := 0
	for j, row := range bytes.Split(bs, []byte("\n")) {
		if len(transposed) < j+1 {
			transposed = append(transposed, make([][]byte, 0))
		}
		for i, cell := range bytes.Split(row, []byte("\t")) {
			if len(transposed[j]) < i+1 {
				transposed[j] = append(transposed[j], make([]byte, 0, maxi))
			}
			transposed[j][i] = cell
			if maxi < i {
				maxi = i
			}
		}

	}
	for i := range transposed[0] {
		for j := range transposed {
			fmt.Printf("%s\t", transposed[j][i])
		}
		fmt.Println()
	}
}
