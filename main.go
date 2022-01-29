package main

import (
	"bufio"
	"fmt"
	"io"
	"io/fs"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
	"regexp"
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

	removeInvalidFiles(&files)

	if len(files) > 0 {
		fmt.Print("\n\n")
	}

	// for i, fInfo := range files {
	for i := 0; i < len(files); i++ {
		fInfo := files[i]

		fileName := fInfo.Name()

		f, err := os.Open(fileName)
		if err != nil {
			log.Fatal(err)
		}

		name := normalizeName(createValidName(getName(f)))

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

	if len(files) > 0 {
		fmt.Print("\n\n")
	}
}

func removeInvalidFiles(files *[]fs.FileInfo) {
	for i, f := range *files {
		// check to make sure the file can be read
		if !strings.Contains(f.Name(), ".java") && !strings.Contains(f.Name(), ".txt") {
			*files = append((*files)[:i], (*files)[i+1:]...)
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

func createValidName(s string) string {
	re, err := regexp.Compile(`[<>:"/\|?*]`)
	if err != nil {
		log.Fatal(err)
	}

	s = re.ReplaceAllString(s, "")

	s = strings.Trim(s, " .")

	return s
}

func normalizeName(s string) string {
	if strings.HasPrefix(strings.ToLower(s), "name ") {
		s = s[5:]
	}

	re, err := regexp.Compile(` ?[pP]{1} ?\d`)
	if err != nil {
		log.Fatal(err)
	}

	s = re.ReplaceAllString(s, "")

	return s
}
