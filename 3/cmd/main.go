package main

import (
	"os"
	"regexp"
	"strings"

	"dwh/internal/exporter"
)

var reg, _ = regexp.Compile("[^0-9]+")

func csv(dstFile *os.File, str string) {
	str = strings.Trim(str, "[DWH] [\\")
	words := strings.FieldsFunc(str, func(r rune) bool {
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
}

func main() {
	//	fill.Populate()
	exporter.ExportCSV()

	// srcFile, _ := os.Open("./src.txt")
	// dstFile, _ := os.OpenFile("./dst.txt", os.O_CREATE|os.O_WRONLY, os.ModePerm)

	// scanner := bufio.NewScanner(srcFile)
	// for scanner.Scan() {
	// 	str := scanner.Text()
	// 	csv(dstFile, str)
	// }

	// dstFile.Sync()
	// defer srcFile.Close()
	// defer dstFile.Close()

	return
}
