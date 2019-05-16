package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)

func getXMLElement(inputLine string) string {
	var xmlElement = inputLine[1:]
	var elementParts = strings.Split(xmlElement, " ")
	xmlElement = elementParts[0]
	return xmlElement
}

func getInlineProperty(inputLine string, propertyName string) (string, bool) {
	if strings.Contains(inputLine, propertyName) {
		var parts = strings.Split(inputLine, propertyName+"=\"")
		var propertyValue = parts[1]
		propertyValue = strings.Split(propertyValue, "\"")[0]
		return propertyValue, true
	}
	return "", false
}

func getInlineValue(inputLine string) string {
	var parts = strings.Split(inputLine, ">")
	var color = parts[1]
	color = strings.Split(color, "<")[0]
	return color
}

func extractElements(xmlFilePaths []string) (styles []Style, colors []Color) {
	// find all styles in xml and add them to objects
	for _, xmlFile := range xmlFilePaths {
		xmlFileText, err := os.Open(xmlFile)
		if err != nil {
			log.Fatal(err)
		}

		scanner := bufio.NewScanner(xmlFileText)

		var inXmlElement = false

		for scanner.Scan() {
			var inputLine = strings.TrimSpace(scanner.Text())

			if inXmlElement {

			}

			if strings.HasPrefix(inputLine, "<") && !strings.HasPrefix(inputLine, "</") {

				var xmlElement = getXMLElement(inputLine)

				if !strings.HasSuffix(inputLine, "</"+xmlElement+">") && !strings.HasSuffix(inputLine, "/>") {
					inXmlElement = true
				}

				if xmlElement == "style" {
					fmt.Println(inputLine)
					nameValue, nameFound := getInlineProperty(inputLine, "name")
					parentValue, parentFound := getInlineProperty(inputLine, "parent")
					if nameFound && parentFound {
						fmt.Println(xmlElement + "\t| name: " + nameValue + "\t| parent: " + parentValue)
					} else if nameFound {
						fmt.Println(xmlElement + "\t| name: " + nameValue)
					}
				}
				if xmlElement == "color" {
					colorNameValue, colorNameFound := getInlineProperty(inputLine, "name")
					if colorNameFound {
						var colorValue = getInlineValue(inputLine)
						fmt.Println(xmlElement + "\t| name: " + colorNameValue + "\t| color: " + colorValue)
					}
				}
			}
		}
	}

	return nil, nil
}
