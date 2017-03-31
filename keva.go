package keva

import (
	"strings"
)

func Parse(fileContent string, sep string) (returnMap map[string]string, err error) {

	returnMap = parseKevaContent(fileContent, sep)

	return

}

func parseKevaContent(content, sep string) (kevaMap map[string]string) {
	kevaMap = map[string]string{}

	cut := ""
	if strings.Contains(content, "\r\n") {
		cut = "\r\n"
	} else if strings.Contains(content, "\n") {
		cut = "\n"
	} else if strings.Contains(content, "\r") {
		cut = "\r"
	}

	lines := strings.Split(content, cut)
	for _, line := range lines {
		if strings.Contains(line, sep) {
			splitAt := strings.Index(line, sep)
			kevaMap[trimAndClean(line[:splitAt])] = trimAndClean(line[splitAt + 1:])
		}
	}

	return
}

func trimAndClean(value string) string {

	value = strings.Trim(value, "\n")
	value = strings.Trim(value, "\r")
	value = strings.Trim(value, "\r\n")
	value = strings.Trim(value, " ")

	if strings.HasPrefix(value, "\"") {
		if strings.HasSuffix(value, "\"") {
			value = strings.TrimLeft(value, "\"")
			value = strings.TrimRight(value, "\"")
		}
	}
	value = strings.Trim(value, " ")

	return value
}
