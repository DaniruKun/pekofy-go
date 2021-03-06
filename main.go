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
			log.Fatal("Failed to read contents of file: ", err)
		} else {
			text = string(bytes)
		}
	case len(os.Args) == 1 || (len(os.Args) == 2 && hasFileFlag(os.Args)):
		fmt.Println("Usage: ", os.Args[0], "Blob of text to pekofy")
		fmt.Println(os.Args[0], "-f FILENAME")
		return
	default:
		text = os.Args[1]
	}

	err := pekofyText(&text)

	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(text)
}

// pekofyText pekofies a given text string by adding `peko` to the end of each sentence terminated by a dot,
// as well as rhymes words starting with `co` with `peko`
func pekofyText(text *string) error {
	if *text == "" {
		return errors.New("nothing to pekofy")
	}

	if strings.LastIndex(*text, ".") == -1 {
		*text += "."
	}

	*text = strings.ReplaceAll(*text, ".", " peko.")
	*text = strings.ReplaceAll(*text, " co", " peko")
	return nil
}

func hasFileFlag(args []string) bool {
	for _, arg := range args {
		if arg == "-f" || arg == "--file" {
			return true
		}
	}
	return false
}
