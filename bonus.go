package main

import (
	"archive/zip"
	"fmt"
	"io"
	"log"
	"os"
	"strings"
)

func main() {

	fmt.Println("Enter Directories (by comma): ")

	var input string

	fmt.Scanln(&input)

	s := strings.Split(input, ",")
	fmt.Println()

	for i, path := range s {
		file, err := os.Open(path)
		if err != nil {
			log.Fatal(err)
		} else {
			fmt.Printf("%v th file exists", i)
			fmt.Println()
		}

		defer file.Close()
	}

	fmt.Println("creating zip archive...")
	bonus, err := os.Create("bonus.zip")
	if err != nil {
		log.Fatal(err)
	}
	defer bonus.Close()
	zipWriter := zip.NewWriter(bonus)

	for _, file := range s {
		f, _ := os.Open(file)
		defer f.Close()

		last := file[strings.LastIndex(file, "/")+1:]

		w, err := zipWriter.Create(last)
		if err != nil {
			log.Fatal(err)
		}

		_, err = io.Copy(w, f)
	}

	fmt.Println("closing zip archive...")
	fmt.Println("\nFiles are successfully zipped!")
	zipWriter.Close()

}
