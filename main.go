package main

import (
	"bufio"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
	"strings"
)

func main() {
	dir, err := os.Getwd()
	if err != nil {
		log.Fatal(err)
	}

	files, err := ioutil.ReadDir(dir)
	if err != nil {
		log.Fatal(err)
	}

	// for i, fInfo := range files {
	for i := 0; i < len(files); i++ {
		fInfo := files[i]

		fileName := fInfo.Name()

		// check to make sure the file can be read
		if !strings.Contains(fileName, ".java") && !strings.Contains(fileName, ".txt") {
			continue
		}

		f, err := os.Open(fileName)
		if err != nil {
			log.Fatal(err)
		}

		name := getName(f)

		f.Close()

		err = os.Mkdir(name, 0755)
		if err != nil {
			fmt.Println("Duplicate from: '" + name + "' ... Replacing")
		}

		// move directories
		oldLoc := filepath.Join(dir, fileName)
		newLoc := filepath.Join(dir, name, fileName[:strings.Index(fileName, "(")]+".java")

		fmt.Printf("%d| ~%s | ~%s\n", i+1, fileName, newLoc[len(dir)+1:])

		err = os.Rename(oldLoc, newLoc)
		if err != nil {
			log.Fatal(err)
		}
	}
}

func getName(f io.Reader) string {
	scan := bufio.NewScanner(f)

	for scan.Scan() {
		s := scan.Text()

		if !strings.Contains(s, "package") && !strings.Contains(s, "import") {
			for i := 0; i < len(s); i++ {
				c := strings.ToLower(s)[i]

				if c >= 'a' && c <= 'z' {
					return s[i:]
				}
			}
		}
	}

	return ""
}
