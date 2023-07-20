package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)

func main() {
	bytesRead, _ := os.ReadFile("/Users/stnc/go-stnc/selmantunc/test/data.html")
	fileContent := string(bytesRead)
	lines := strings.Split(fileContent, "\n")

	bytesReadSed, _ := os.ReadFile("/Users/stnc/go-stnc/selmantunc/test/dataSed.html")
	fileContentSed := string(bytesReadSed)
	linesSED := strings.Split(fileContentSed, "\n")

	for index, element := range lines {

		normalName := name_(element)
		sedName := linesSED[index]

		newName := "http:\\/\\/localhost:8082\\/wp-content\\/uploads\\/2023\\/07\\/" + normalName
		fmt.Println(newName)
		fmt.Println(sedName)
		//
		sed := "sed 's/" + sedName + "/" + newName + "/g' s.xml > changed.txt && mv changed.txt s.xml;"
		Writer(sed)
	}
}

func name_(url string) string {
	stringSlice := strings.Split(url, "/")
	s1 := stringSlice[len(stringSlice)-1]
	return s1
}
func Writer(data string) {
	file, err := os.OpenFile("run.sh", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)

	if err != nil {
		log.Fatalf("failed creating file: %s", err)
	}

	datawriter := bufio.NewWriter(file)

	datawriter.WriteString(data + "\n")

	datawriter.Flush()
	file.Close()
}
