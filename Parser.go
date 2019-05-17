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

		for scanner.Scan() {
			var inputLine = strings.TrimSpace(scanner.Text())
			var xmlElement = getXMLElement(inputLine)
			var styles []Style
			var colors []Color
			
			if strings.HasPrefix(inputLine, "<") && !strings.HasPrefix(inputLine, "</") {

				// if !strings.HasSuffix(inputLine, "</"+xmlElement+">") && !strings.HasSuffix(inputLine, "/>") {
				// 	inXmlElement = true
				// }

				if xmlElement == "item" {
					itemNameValue, itemNameFound := getInlineProperty(inputLine, "name")
					itemValue := getInlineValue(inputLine)
					tempItem := Item{itemNameValue, itemValue}

					if itemNameFound {
						styles[len(styles) - 1].Items = append(styles[len(styles) - 1].Items, tempItem)
					}
				}

				if xmlElement == "style" {
					styleNameValue, styleNameFound := getInlineProperty(inputLine, "name")
					parentValue, parentFound := getInlineProperty(inputLine, "parent")
					if styleNameFound && parentFound {
						styles = append(styles, Style{styleNameValue, parentValue, []Item{}})
						//fmt.Println(xmlElement + "\t| name: " + styleNameValue + "\t| parent: " + parentValue)
					} else if styleNameFound {
						
						fmt.Println(xmlElement + "\t| name: " + styleNameValue)
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
