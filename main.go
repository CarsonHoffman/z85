package main

import (
	"fmt"
	"io/ioutil"
	"os"

	"github.com/tilinna/z85"
)

func main() {
	transform := z85.Encode
	transformedLength := z85.EncodedLen

	if len(os.Args) >= 2 && (os.Args[1] == "d" || os.Args[1] == "decode") {
		transform = z85.Decode
		transformedLength = z85.DecodedLen
	}

	input, err := ioutil.ReadAll(os.Stdin)
	if err != nil {
		fmt.Fprintln(os.Stderr, "failed to read from stdin")
		os.Exit(1)
	}

	output := make([]byte, transformedLength(len(input)))
	_, err = transform(output, input)
	if err != nil {
		fmt.Fprintln(os.Stderr, "failed to transform input:", err)
		os.Exit(1)
	}

	fmt.Print(string(output))
}
