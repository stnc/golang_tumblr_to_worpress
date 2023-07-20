package main

import (
	"bufio"
	"fmt"
	"io"
	"net/http"
	"os"
	"strings"
)

func main() {

	readFile, err := os.Open("/Users/stnc/go-stnc/selmantunc/test/data.html")

	if err != nil {
		fmt.Println(err)
	}
	fileScanner := bufio.NewScanner(readFile)

	fileScanner.Split(bufio.ScanLines)

	for fileScanner.Scan() {
		fmt.Println(fileScanner.Text())
		url := string(fileScanner.Text())
		err := DownloadFile("down/"+name_(url), url)
		if err != nil {
			panic(err)
		}
		fmt.Println("Downloaded: " + url)
	}

	readFile.Close()
}
func DownloadFile(filepath string, url string) error {

	// Get the data
	resp, err := http.Get(url)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	// Create the file
	out, err := os.Create(filepath)
	if err != nil {
		return err
	}
	defer out.Close()

	// Write the body to file
	_, err = io.Copy(out, resp.Body)
	return err
}
func name_(url string) string {
	stringSlice := strings.Split(url, "/")
	s1 := stringSlice[len(stringSlice)-1]
	return s1
}
