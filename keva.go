package keva

import (
	"io/ioutil"
	"strings"
)

func Open(fileName string) (returnMap map[string]string, err error) {
	fileBytes, err := ioutil.ReadFile(fileName)
	if err != nil {
		err = kevaError{}
		return
	}
	fileContent := string(fileBytes)
	returnMap = parseKevaContent(fileContent)

	return

}

func parseKevaContent(content string) (kevaMap map[string]string) {
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
		if strings.Contains(line, "=") {
			splitAt := strings.Index(line, "=")
			kevaMap[line[:splitAt]] = line[splitAt + 1:]
		}
	}

	return
}
