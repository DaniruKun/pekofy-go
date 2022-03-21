package main

import (
	"errors"
	"fmt"
	"log"
	"os"
	"strings"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Println("Usage: ", os.Args[0], "Blob of text to pekofy")
		return
	}

	var text string

	if len(os.Args) > 2 {
		text = strings.Join(os.Args[1:], " ")
	} else {
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
