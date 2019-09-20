package exporter

import (
	"bufio"
	"os"
	"regexp"
	"strings"
)

var reg, _ = regexp.Compile("[^0-9]+")

func ExportCSV() {
	threads := 5
	srcFile, _ := os.Open("./src.txt")
	dstFile, _ := os.OpenFile("./dst.txt", os.O_CREATE|os.O_WRONLY, os.ModePerm)

	stringsToConvert := make(chan string)
	noMoreStrings := make(chan bool)

	for i := 0; i < threads; i++ {
		go processString(dstFile, stringsToConvert, noMoreStrings)
	}

	scanner := bufio.NewScanner(srcFile)
	for scanner.Scan() {
		stringsToConvert <- scanner.Text()
	}

	for i := 0; i < threads; i++ {
		noMoreStrings <- true
	}

	dstFile.Sync()

	srcFile.Close()
	dstFile.Close()
}

func processString(dstFile *os.File, stringsToConvert chan string, noMoreStrings chan bool) {
loop:
	for true {
		select {
		case s := <-stringsToConvert:
			s = strings.Trim(s, "[DWH] [\\")
			words := strings.FieldsFunc(s, func(r rune) bool {
				return r == '@' || r == '|'
			})
			words[5] = reg.ReplaceAllString(words[5], "")

			var finalStr string
			for idx, word := range words {
				if idx != len(words)-1 {
					finalStr += "\"" + word + "\","
				} else {
					finalStr += "\"" + word + "\"\n"
				}
			}

			dstFile.WriteString(finalStr)

		case <-noMoreStrings:
			break loop
		}
	}
}
