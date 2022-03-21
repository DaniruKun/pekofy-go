package main

import (
	"errors"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"strings"
)

func main() {
	var text string

	switch {
	case len(os.Args) == 3 && hasFileFlag(os.Args):
		bytes, err := ioutil.ReadFile(os.Args[2])
		if err != nil {
			log.Fatal("Failed to read contents of file: ", os.Args[2])
		} else {
			text = string(bytes)
		}
	case len(os.Args) == 1 || (len(os.Args) == 2 && hasFileFlag(os.Args)):
		fmt.Println("Usage: ", os.Args[0], "Blob of text to pekofy")
		return
	default:
		text = os.Args[1]
	}

	res, err := pekofyText(text)

	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(res)
}

// pekofyText pekofies a given text string by adding `peko` to the end of each sentence terminated by a dot,
// as well as rhymes words starting with `co` with `peko`
func pekofyText(text string) (string, error) {
	if text == "" {
		return text, errors.New("nothing to pekofy")
	}

	if strings.LastIndex(text, ".") == -1 {
		text += "."
	}

	var pekofied string

	pekofied = strings.ReplaceAll(text, ".", " peko.")
	pekofied = strings.ReplaceAll(pekofied, " co", " peko")
	return pekofied, nil
}

func hasFileFlag(args []string) bool {
	for _, arg := range args {
		if arg == "-f" || arg == "--file" {
			return true
		}
	}
	return false
}
