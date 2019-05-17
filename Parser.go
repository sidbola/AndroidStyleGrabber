package main

import (
	"bufio"
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
	if len(parts) > 1 {
		var color = parts[1]
		color = strings.Split(color, "<")[0]
		return color
	}
	return ""
}

func extractElements(xmlFilePaths []string) ([]Style, []Color) {
	var styles []Style
	var colors []Color
	// find all styles in xml and add them to objects
	for _, xmlFile := range xmlFilePaths {
		xmlFileText, err := os.Open(xmlFile)
		if err != nil {
			log.Fatal(err)
		}

		scanner := bufio.NewScanner(xmlFileText)

		for scanner.Scan() {
			var inputLine = strings.TrimSpace(scanner.Text())

			if strings.HasPrefix(inputLine, "<") && !strings.HasPrefix(inputLine, "</") && !strings.HasPrefix(inputLine, "<!--") {
				var xmlElement = getXMLElement(inputLine)

				// if !strings.HasSuffix(inputLine, "</"+xmlElement+">") && !strings.HasSuffix(inputLine, "/>") {
				// 	inXmlElement = true
				// }

				if xmlElement == "item" && len(styles) > 0 {
					itemNameValue, itemNameFound := getInlineProperty(inputLine, "name")
					itemValue := getInlineValue(inputLine)
					tempItem := Item{itemNameValue, itemValue}

					if itemNameFound {
						styles[len(styles)-1].Items = append(styles[len(styles)-1].Items, tempItem)
					}
				}

				if xmlElement == "style" {
					styleNameValue, styleNameFound := getInlineProperty(inputLine, "name")
					parentValue, parentFound := getInlineProperty(inputLine, "parent")
					if styleNameFound && parentFound {
						styles = append(styles, Style{styleNameValue, parentValue, []Item{}})
						//fmt.Println(xmlElement + "\t| name: " + styleNameValue + "\t| parent: " + parentValue)
					} else if styleNameFound {
						styles = append(styles, Style{styleNameValue, "", []Item{}})
						//fmt.Println(xmlElement + "\t| name: " + styleNameValue)
					}
				}
				if xmlElement == "color" {
					colorNameValue, colorNameFound := getInlineProperty(inputLine, "name")
					if colorNameFound {
						var colorValue = getInlineValue(inputLine)
						colors = append(colors, Color{colorNameValue, colorValue})
						//fmt.Println(xmlElement + "\t| name: " + colorNameValue + "\t| color: " + colorValue)
					}
				}
			}
		}
	}

	return styles, colors
}
