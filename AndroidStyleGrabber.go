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

	// find all styles in xml and add them to objects
	for _, xmlFile := range xmlFilePaths {
		xmlFileText, err := os.Open(xmlFile)
		if err != nil {
			log.Fatal(err)
		}

		scanner := bufio.NewScanner(xmlFileText)

		for scanner.Scan() {
			var inputLine = strings.TrimSpace(scanner.Text())
			//var inOpeningTag = false
			//var inXmlElement = false

			if strings.HasPrefix(inputLine, "<") && !strings.HasPrefix(inputLine, "</") {
				var xmlElement = getXmlElement(inputLine)
				if xmlElement == "style" {
					nameValue, propertyFound := getInlineProperty(inputLine, "name")
					if propertyFound {
						fmt.Println(xmlElement + " " + nameValue)
					}
				}
			}
		}
	}

}

func getXmlElement(inputLine string) string {
	var xmlElement = inputLine[1:]
	var elementParts = strings.Split(xmlElement, " ")
	xmlElement = elementParts[0]
	return xmlElement
}

func getInlineProperty(inputLine string, propertyName string) (string, bool) {
	if strings.Contains(inputLine, propertyName) {
		var parts = strings.Split(inputLine, propertyName + "=\"")
		var propertyValue = parts[1]
		propertyValue = strings.Split(propertyValue, "\"")[0]

		return propertyValue, true
	} else {
		return "", false
	}
}