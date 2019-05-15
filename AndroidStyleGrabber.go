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

	// project path
	var projectPath = "insert path here"

	var xmlFilePaths []string
	var resColors []Color

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
			if strings.HasPrefix(inputLine, "<style") || strings.HasPrefix(inputLine, "<item") {
				if strings.Contains(inputLine, "name=\"") {
					var parts = strings.Split(inputLine, "\"")
					var foundName = false
					var foundValue = false

					if strings.HasPrefix(inputLine, "<style") {
						for _, part := range parts {
							if foundName {
								fmt.Println(part)
								output.WriteString(part + "\n")
								break
							}
							if strings.Contains(part, "name=") {
								foundName = true
							}
						}
					} else {
						for _, part := range parts {
							if foundValue {
								var value = strings.Replace(part, ">", "", -1)
								value = strings.Replace(value, "</item", "", -1)
								fmt.Println("\t\t- " + value)
								output.WriteString("\t\t- " + value + "\n")
								foundValue = false
								break
							}
							if foundName {
								fmt.Print("\t" + strings.Replace(part, "android:", "", -1))
								output.WriteString("\t" + strings.Replace(part, "android:", "", -1))
								foundValue = true
							}
							if strings.Contains(part, "name=") {
								foundName = true
							}
						}
					}
				}

			} else if strings.HasPrefix(inputLine, "<color") {
				if strings.Contains(inputLine, "name=\"") {
					var parts = strings.Split(inputLine, "\"")

					var value = parts[2]
					value = strings.Replace(value, ">", "", -1)
					value = strings.Replace(value, "</color", "", -1)

					resColors = append(resColors, Color{parts[1], value})
				}
			}
		}
	}

	output.WriteString("COLORS\n\n")

	for _, color := range resColors {
		fmt.Println(color)
		//output.WriteString(color)
		//output.WriteString("\n")
	} 

}
