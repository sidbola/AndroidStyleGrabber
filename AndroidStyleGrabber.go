package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"path/filepath"
	"strings"
)

func main() {
	output, err1 := os.Create("output.txt")
	if err1 != nil {
		return
	}
	output.WriteString("STYLES\n\n")

	pathReader := bufio.NewReader(os.Stdin)
	fmt.Print("Enter the path to the project: ")
	text, _ := pathReader.ReadString('\n')

	// project path
	var projectPath = strings.TrimSpace(text)

	var xmlFilePaths []string
	//var resColors []Color

	// find all xml files and add them to a slice
	err := filepath.Walk(projectPath,
		func(path string, info os.FileInfo, err error) error {
			if err != nil {
				return err
			}
			if strings.HasSuffix(path, ".xml") {
				xmlFilePaths = append(xmlFilePaths, path)
			}
			return nil
		})
	if err != nil {
		log.Println(err)
	}

	styles, colors := extractElements(xmlFilePaths)
	if styles == nil {
		return
	} else if colors == nil {
		return
	}

}
